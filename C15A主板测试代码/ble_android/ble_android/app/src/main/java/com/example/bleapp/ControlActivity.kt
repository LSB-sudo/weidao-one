package com.example.bleapp

import android.Manifest
import android.annotation.SuppressLint
import android.bluetooth.*
import android.bluetooth.le.ScanCallback
import android.bluetooth.le.ScanResult
import android.content.Context
import android.content.pm.PackageManager
import android.os.Build
import android.os.Bundle
import android.os.Handler
import android.os.Looper
import android.widget.*
import androidx.appcompat.app.AlertDialog
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import java.util.*

/**
 * 电机控制界面 —— 独立的 BLE 扫描 / 连接 / 摇杆控制 / 协议解析。
 *
 * 协议参见：蓝牙控制.md
 *   发送：RPM,L:<left>,R:<right>\n       (-268 ~ 268)
 *   接收：OK,L:<l>,R:<r>\n  |  ERR,<type>\n  |  Battery:<v>V\r\n
 */
class ControlActivity : AppCompatActivity() {

    // ── 常量 ──────────────────────────────────────────────────────────────────

    companion object {
        val SERVICE_UUID: UUID = UUID.fromString("0000ffe0-0000-1000-8000-00805f9b34fb")
        val CHAR_UUID: UUID    = UUID.fromString("0000ffe1-0000-1000-8000-00805f9b34fb")
        val CCCD_UUID: UUID    = UUID.fromString("00002902-0000-1000-8000-00805f9b34fb")

        const val SCAN_DURATION_MS = 5000L
        const val REQ_PERMISSIONS  = 2            // 与 MainActivity 的 1 区分

        const val MAX_RPM    = 268
        const val SEND_INTERVAL_MS = 50L          // 摇杆持续发送间隔
    }

    // ── BLE 对象 ──────────────────────────────────────────────────────────────

    private val bluetoothAdapter: BluetoothAdapter? by lazy {
        (getSystemService(Context.BLUETOOTH_SERVICE) as BluetoothManager).adapter
    }

    private var gatt: BluetoothGatt? = null
    private var txChar: BluetoothGattCharacteristic? = null

    private val scanDevices = mutableListOf<BluetoothDevice>()
    private val scanLabels  = mutableListOf<String>()

    private val rxBuf = StringBuilder()
    private val mainHandler = Handler(Looper.getMainLooper())

    // ── 摇杆状态 ──────────────────────────────────────────────────────────────

    private var currentLeftRpm  = 0
    private var currentRightRpm = 0
    @Volatile private var joystickActive = false     // 手指是否按在摇杆上

    private val sendRunnable = object : Runnable {
        override fun run() {
            if (joystickActive) {
                sendRpm(currentLeftRpm, currentRightRpm)
                mainHandler.postDelayed(this, SEND_INTERVAL_MS)
            }
        }
    }

    // ── UI 控件 ───────────────────────────────────────────────────────────────

    private lateinit var tvStatus: TextView
    private lateinit var btnScan: Button
    private lateinit var btnDisconnect: Button
    private lateinit var tvBattery: TextView
    private lateinit var tvAck: TextView
    private lateinit var joystick: JoystickView
    private lateinit var tvLeftRpm: TextView
    private lateinit var tvRightRpm: TextView
    private lateinit var btnStop: Button
    private lateinit var tvLog: TextView
    private lateinit var scrollLog: ScrollView

    // ══════════════════════════════════════════════════════════════════════════
    //  生命周期
    // ══════════════════════════════════════════════════════════════════════════

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_control)

        tvStatus      = findViewById(R.id.tvCtrlStatus)
        btnScan       = findViewById(R.id.btnCtrlScan)
        btnDisconnect = findViewById(R.id.btnCtrlDisconnect)
        tvBattery     = findViewById(R.id.tvBattery)
        tvAck         = findViewById(R.id.tvAck)
        joystick      = findViewById(R.id.joystick)
        tvLeftRpm     = findViewById(R.id.tvLeftRpm)
        tvRightRpm    = findViewById(R.id.tvRightRpm)
        btnStop       = findViewById(R.id.btnEmergencyStop)
        tvLog         = findViewById(R.id.tvCtrlLog)
        scrollLog     = findViewById(R.id.scrollCtrlLog)

        btnScan.setOnClickListener       { startScan() }
        btnDisconnect.setOnClickListener { disconnect() }
        btnStop.setOnClickListener       { emergencyStop() }

        // ── 摇杆回调 ──────────────────────────────────────────────────────────
        joystick.listener = object : JoystickView.OnJoystickMoveListener {
            override fun onJoystickMove(x: Float, y: Float) {
                if (x == 0f && y == 0f) {
                    // 松手
                    joystickActive = false
                    mainHandler.removeCallbacks(sendRunnable)
                    currentLeftRpm  = 0
                    currentRightRpm = 0
                    updateRpmDisplay()
                    sendRpm(0, 0)         // 立即发送停止
                } else {
                    // 差速模型：y → 基础速度, x → 转向差速
                    val base = (y * MAX_RPM).toInt()
                    val diff = (x * MAX_RPM).toInt()
                    currentLeftRpm  = (base + diff).coerceIn(-MAX_RPM, MAX_RPM)
                    currentRightRpm = (base - diff).coerceIn(-MAX_RPM, MAX_RPM)
                    updateRpmDisplay()

                    if (!joystickActive) {
                        joystickActive = true
                        mainHandler.post(sendRunnable)   // 启动定时发送
                    }
                }
            }
        }

        setUiConnected(false)
    }

    override fun onDestroy() {
        super.onDestroy()
        joystickActive = false
        mainHandler.removeCallbacks(sendRunnable)
        closeGatt()
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  权限
    // ══════════════════════════════════════════════════════════════════════════

    private fun neededPermissions(): Array<String> =
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.S)
            arrayOf(Manifest.permission.BLUETOOTH_SCAN, Manifest.permission.BLUETOOTH_CONNECT)
        else
            arrayOf(Manifest.permission.ACCESS_FINE_LOCATION)

    private fun hasPermissions() = neededPermissions().all {
        ActivityCompat.checkSelfPermission(this, it) == PackageManager.PERMISSION_GRANTED
    }

    override fun onRequestPermissionsResult(
        requestCode: Int, permissions: Array<String>, grantResults: IntArray
    ) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults)
        if (requestCode == REQ_PERMISSIONS &&
            grantResults.all { it == PackageManager.PERMISSION_GRANTED }
        ) doScan()
        else appendLog("[系统] 权限被拒绝", "err")
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  BLE 扫描
    // ══════════════════════════════════════════════════════════════════════════

    private fun startScan() {
        if (!hasPermissions()) {
            ActivityCompat.requestPermissions(this, neededPermissions(), REQ_PERMISSIONS)
            return
        }
        doScan()
    }

    @SuppressLint("MissingPermission")
    private fun doScan() {
        val scanner = bluetoothAdapter?.bluetoothLeScanner
        if (scanner == null) {
            appendLog("[系统] 蓝牙不可用", "err")
            return
        }

        scanDevices.clear()
        scanLabels.clear()
        btnScan.isEnabled = false
        btnScan.text      = "扫描中…"
        appendLog("[系统] 开始扫描（${SCAN_DURATION_MS / 1000}s）…")

        val cb = object : ScanCallback() {
            override fun onScanResult(callbackType: Int, result: ScanResult) {
                val dev = result.device
                if (scanDevices.none { it.address == dev.address }) {
                    scanDevices.add(dev)
                    @SuppressLint("MissingPermission")
                    val label = "${dev.name ?: "未知设备"}  [${dev.address}]"
                    scanLabels.add(label)
                    appendLog("[扫描] $label")
                }
            }
        }

        scanner.startScan(cb)

        mainHandler.postDelayed({
            scanner.stopScan(cb)
            btnScan.isEnabled = true
            btnScan.text      = "扫描"
            if (scanDevices.isEmpty()) appendLog("[系统] 未发现设备", "err")
            else showPicker()
        }, SCAN_DURATION_MS)
    }

    @SuppressLint("MissingPermission")
    private fun showPicker() {
        AlertDialog.Builder(this)
            .setTitle("选择设备")
            .setItems(scanLabels.toTypedArray()) { _, idx -> connectDevice(scanDevices[idx]) }
            .setNegativeButton("取消", null)
            .show()
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  BLE 连接 / GATT 回调
    // ══════════════════════════════════════════════════════════════════════════

    @SuppressLint("MissingPermission")
    private fun connectDevice(device: BluetoothDevice) {
        appendLog("[系统] 正在连接 ${device.name ?: device.address}…")
        tvStatus.text = "● 连接中…"
        tvStatus.setTextColor(0xFFFFCC00.toInt())
        gatt = device.connectGatt(this, false, gattCallback)
    }

    private val gattCallback = object : BluetoothGattCallback() {

        @SuppressLint("MissingPermission")
        override fun onConnectionStateChange(gatt: BluetoothGatt, status: Int, newState: Int) {
            mainHandler.post {
                when (newState) {
                    BluetoothProfile.STATE_CONNECTED -> {
                        appendLog("[系统] GATT 已连接，发现服务中…")
                        gatt.discoverServices()
                    }
                    BluetoothProfile.STATE_DISCONNECTED -> {
                        txChar = null
                        this@ControlActivity.gatt = null
                        setUiConnected(false)
                        appendLog("[系统] 连接断开 (status=$status)")
                    }
                }
            }
        }

        @SuppressLint("MissingPermission")
        override fun onServicesDiscovered(gatt: BluetoothGatt, status: Int) {
            mainHandler.post {
                val char = gatt.getService(SERVICE_UUID)?.getCharacteristic(CHAR_UUID)
                if (char == null) {
                    appendLog("[系统] 未找到 FFE1 特征", "err")
                    closeGatt()
                    return@post
                }

                txChar = char

                // 开启通知
                gatt.setCharacteristicNotification(char, true)
                char.getDescriptor(CCCD_UUID)?.let { desc ->
                    @Suppress("DEPRECATION")
                    desc.value = BluetoothGattDescriptor.ENABLE_NOTIFICATION_VALUE
                    gatt.writeDescriptor(desc)
                }

                setUiConnected(true)
                appendLog("[系统] 已就绪，可以控制电机")
            }
        }

        // API < 33
        @Suppress("DEPRECATION")
        override fun onCharacteristicChanged(
            gatt: BluetoothGatt,
            characteristic: BluetoothGattCharacteristic
        ) {
            handleRx(characteristic.value ?: return)
        }

        // API >= 33
        override fun onCharacteristicChanged(
            gatt: BluetoothGatt,
            characteristic: BluetoothGattCharacteristic,
            value: ByteArray
        ) {
            handleRx(value)
        }
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  接收解析
    // ══════════════════════════════════════════════════════════════════════════

    private fun handleRx(bytes: ByteArray) {
        val text = bytes.toString(Charsets.UTF_8)
        mainHandler.post {
            rxBuf.append(text)
            var nl: Int
            while (rxBuf.indexOf('\n').also { nl = it } != -1) {
                val line = rxBuf.substring(0, nl).trimEnd('\r')
                rxBuf.delete(0, nl + 1)
                if (line.isBlank()) continue

                parseLine(line)
            }
        }
    }

    /**
     * 解析一行固件返回数据。
     *
     * 支持：
     *   Battery:12.34V
     *   OK,L:100,R:-100
     *   ERR,FORMAT / ERR,RANGE / ERR,VALUE / ERR,LEN
     */
    private fun parseLine(line: String) {
        when {
            line.startsWith("Battery:") -> {
                // Battery:12.34V
                val voltage = line.removePrefix("Battery:").removeSuffix("V")
                tvBattery.text = "${voltage}V"
                appendLog("[RX] $line", "rx")
            }

            line.startsWith("OK,") -> {
                // OK,L:100,R:-100
                tvAck.text = line
                tvAck.setTextColor(0xFF4CAF50.toInt())   // 绿色
                appendLog("[RX] $line", "rx")
            }

            line.startsWith("ERR,") -> {
                // ERR,FORMAT / ERR,RANGE / ERR,VALUE / ERR,LEN
                tvAck.text = line
                tvAck.setTextColor(0xFFFF4444.toInt())   // 红色
                appendLog("[RX] $line", "err")
            }

            else -> {
                // 未知数据，原样显示
                appendLog("[RX] $line", "rx")
            }
        }
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  发送
    // ══════════════════════════════════════════════════════════════════════════

    @SuppressLint("MissingPermission")
    private fun sendRpm(left: Int, right: Int) {
        val char = txChar ?: return
        val cmd = "RPM,L:$left,R:$right\n"
        val bytes = cmd.toByteArray(Charsets.UTF_8)

        @Suppress("DEPRECATION")
        char.value = bytes
        @Suppress("DEPRECATION")
        gatt?.writeCharacteristic(char)

        appendLog("[TX] RPM,L:$left,R:$right", "tx")
    }

    private fun emergencyStop() {
        joystickActive = false
        mainHandler.removeCallbacks(sendRunnable)
        currentLeftRpm  = 0
        currentRightRpm = 0
        updateRpmDisplay()
        sendRpm(0, 0)
        appendLog("[系统] ⛔ 紧急停止", "err")
    }

    // ══════════════════════════════════════════════════════════════════════════
    //  辅助
    // ══════════════════════════════════════════════════════════════════════════

    @SuppressLint("MissingPermission")
    private fun disconnect() {
        joystickActive = false
        mainHandler.removeCallbacks(sendRunnable)
        closeGatt()
        appendLog("[系统] 手动断开")
    }

    @SuppressLint("MissingPermission")
    private fun closeGatt() {
        gatt?.disconnect()
        gatt?.close()
        gatt   = null
        txChar = null
        mainHandler.post { setUiConnected(false) }
    }

    private fun setUiConnected(connected: Boolean) {
        btnScan.isEnabled       = !connected
        btnDisconnect.isEnabled = connected
        btnStop.isEnabled       = connected
        joystick.isEnabled      = connected

        tvStatus.text = if (connected) "● 已连接" else "● 未连接"
        tvStatus.setTextColor(
            if (connected) 0xFF4CAF50.toInt() else 0xFF888888.toInt()
        )

        if (!connected) {
            joystickActive = false
            mainHandler.removeCallbacks(sendRunnable)
        }
    }

    private fun updateRpmDisplay() {
        tvLeftRpm.text  = currentLeftRpm.toString()
        tvRightRpm.text = currentRightRpm.toString()
    }

    private fun appendLog(msg: String, type: String = "sys") {
        val color = when (type) {
            "rx"  -> 0xFF4CAF50.toInt()
            "tx"  -> 0xFF7CB9E8.toInt()
            "err" -> 0xFFFF4444.toInt()
            else  -> 0xFFAAAAAA.toInt()
        }
        mainHandler.post {
            val time = java.text.SimpleDateFormat("HH:mm:ss", Locale.getDefault())
                .format(java.util.Date())
            val cur = tvLog.text
            val newText = if (cur.isEmpty()) "[$time] $msg" else "$cur\n[$time] $msg"
            tvLog.text = newText
            val lines = tvLog.text.split('\n')
            if (lines.size > 200) {
                tvLog.text = lines.takeLast(200).joinToString("\n")
            }
            scrollLog.post { scrollLog.fullScroll(ScrollView.FOCUS_DOWN) }
        }
    }
}
