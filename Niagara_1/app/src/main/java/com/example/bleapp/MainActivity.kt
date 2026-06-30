package com.example.bleapp

import android.Manifest
import android.annotation.SuppressLint
import android.bluetooth.BluetoothAdapter
import android.bluetooth.BluetoothGatt
import android.bluetooth.BluetoothGattCallback
import android.bluetooth.BluetoothGattCharacteristic
import android.bluetooth.BluetoothGattDescriptor
import android.bluetooth.BluetoothManager
import android.bluetooth.BluetoothProfile
import android.bluetooth.le.ScanCallback
import android.bluetooth.le.ScanResult
import android.content.Context
import android.content.pm.PackageManager
import android.os.Build
import android.os.Bundle
import android.os.Handler
import android.os.Looper
import android.text.Html
import android.view.inputmethod.EditorInfo
import android.widget.Button
import android.widget.EditText
import android.widget.ScrollView
import android.widget.TextView
import androidx.appcompat.app.AlertDialog
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import java.util.UUID

class MainActivity : AppCompatActivity() {

    companion object {
        val SERVICE_UUID: UUID = UUID.fromString("0000ffe0-0000-1000-8000-00805f9b34fb")
        val CHAR_UUID: UUID = UUID.fromString("0000ffe1-0000-1000-8000-00805f9b34fb")
        val CCCD_UUID: UUID = UUID.fromString("00002902-0000-1000-8000-00805f9b34fb")

        const val SCAN_DURATION_MS = 5000L
        const val REQ_PERMISSIONS = 1
    }

    private val bluetoothAdapter: BluetoothAdapter? by lazy {
        (getSystemService(Context.BLUETOOTH_SERVICE) as BluetoothManager).adapter
    }

    private var gatt: BluetoothGatt? = null
    private var txChar: BluetoothGattCharacteristic? = null

    private val scanDevices = mutableListOf<android.bluetooth.BluetoothDevice>()
    private val scanLabels = mutableListOf<String>()
    private val rxBuf = StringBuilder()
    private val mainHandler = Handler(Looper.getMainLooper())

    private lateinit var btnScan: Button
    private lateinit var btnDisconnect: Button
    private lateinit var tvStatus: TextView
    private lateinit var tvLog: TextView
    private lateinit var scrollLog: ScrollView
    private lateinit var etSend: EditText
    private lateinit var btnSend: Button

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        btnScan = findViewById(R.id.btnScan)
        btnDisconnect = findViewById(R.id.btnDisconnect)
        tvStatus = findViewById(R.id.tvStatus)
        tvLog = findViewById(R.id.tvLog)
        scrollLog = findViewById(R.id.scrollLog)
        etSend = findViewById(R.id.etSend)
        btnSend = findViewById(R.id.btnSend)

        btnScan.setOnClickListener { startScan() }
        btnDisconnect.setOnClickListener { disconnect() }
        btnSend.setOnClickListener { sendText() }
        etSend.setOnEditorActionListener { _, action, _ ->
            if (action == EditorInfo.IME_ACTION_SEND) {
                sendText()
                true
            } else {
                false
            }
        }

        setUiConnected(false)
    }

    override fun onDestroy() {
        super.onDestroy()
        closeGatt()
    }

    private fun neededPermissions(): Array<String> =
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.S) {
            arrayOf(Manifest.permission.BLUETOOTH_SCAN, Manifest.permission.BLUETOOTH_CONNECT)
        } else {
            arrayOf(Manifest.permission.ACCESS_FINE_LOCATION)
        }

    private fun hasPermissions() = neededPermissions().all {
        ActivityCompat.checkSelfPermission(this, it) == PackageManager.PERMISSION_GRANTED
    }

    override fun onRequestPermissionsResult(
        requestCode: Int,
        permissions: Array<String>,
        grantResults: IntArray
    ) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults)
        if (requestCode == REQ_PERMISSIONS && grantResults.all { it == PackageManager.PERMISSION_GRANTED }) {
            doScan()
        } else {
            appendLog("[系统] 权限被拒绝", "err")
        }
    }

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
            appendLog("[系统] 蓝牙不可用，请确认已开启蓝牙", "err")
            return
        }

        scanDevices.clear()
        scanLabels.clear()
        btnScan.isEnabled = false
        btnScan.text = "扫描中"
        appendLog("[系统] 开始扫描（${SCAN_DURATION_MS / 1000}s）")

        val cb = object : ScanCallback() {
            override fun onScanResult(callbackType: Int, result: ScanResult) {
                val dev = result.device
                if (scanDevices.none { it.address == dev.address }) {
                    scanDevices.add(dev)
                    val label = "${dev.name ?: "未知设备"} [${dev.address}]"
                    scanLabels.add(label)
                    appendLog("[扫描] $label")
                }
            }
        }

        scanner.startScan(cb)

        mainHandler.postDelayed({
            scanner.stopScan(cb)
            btnScan.isEnabled = true
            btnScan.text = "扫描"
            if (scanDevices.isEmpty()) {
                appendLog("[系统] 未发现 BLE 设备", "err")
            } else {
                showPicker()
            }
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

    @SuppressLint("MissingPermission")
    private fun connectDevice(device: android.bluetooth.BluetoothDevice) {
        appendLog("[系统] 正在连接 ${device.name ?: device.address}")
        tvStatus.text = "● 连接中"
        tvStatus.setTextColor(0xFFFFCC00.toInt())
        gatt = device.connectGatt(this, false, gattCallback)
    }

    private val gattCallback = object : BluetoothGattCallback() {
        @SuppressLint("MissingPermission")
        override fun onConnectionStateChange(gatt: BluetoothGatt, status: Int, newState: Int) {
            mainHandler.post {
                when (newState) {
                    BluetoothProfile.STATE_CONNECTED -> {
                        appendLog("[系统] GATT 已连接，正在发现服务")
                        gatt.discoverServices()
                    }
                    BluetoothProfile.STATE_DISCONNECTED -> {
                        txChar = null
                        this@MainActivity.gatt = null
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
                    appendLog("[系统] 未找到 UUID=FFE1 特征", "err")
                    closeGatt()
                    return@post
                }

                txChar = char
                gatt.setCharacteristicNotification(char, true)
                char.getDescriptor(CCCD_UUID)?.let { desc ->
                    @Suppress("DEPRECATION")
                    desc.value = BluetoothGattDescriptor.ENABLE_NOTIFICATION_VALUE
                    gatt.writeDescriptor(desc)
                }

                setUiConnected(true)
                appendLog("[系统] 已就绪，可以收发数据")
            }
        }

        @Suppress("DEPRECATION")
        override fun onCharacteristicChanged(
            gatt: BluetoothGatt,
            characteristic: BluetoothGattCharacteristic
        ) {
            handleRx(characteristic.value ?: return)
        }

        override fun onCharacteristicChanged(
            gatt: BluetoothGatt,
            characteristic: BluetoothGattCharacteristic,
            value: ByteArray
        ) {
            handleRx(value)
        }
    }

    private fun handleRx(bytes: ByteArray) {
        val text = bytes.toString(Charsets.UTF_8)
        mainHandler.post {
            rxBuf.append(text)
            while (true) {
                val nl = rxBuf.indexOf("\n")
                if (nl == -1) break
                val line = rxBuf.substring(0, nl).trimEnd('\r')
                rxBuf.delete(0, nl + 1)
                if (line.isNotBlank()) appendLog("[RX] $line", "rx")
            }
        }
    }

    @SuppressLint("MissingPermission")
    private fun sendText() {
        val text = etSend.text.toString().trim()
        if (text.isEmpty()) return

        val char = txChar
        if (char == null) {
            appendLog("[系统] 尚未连接设备", "err")
            return
        }

        val payload = "$text\r\n".toByteArray(Charsets.UTF_8)
        @Suppress("DEPRECATION")
        char.value = payload
        @Suppress("DEPRECATION")
        gatt?.writeCharacteristic(char)
        appendLog("[TX] $text", "tx")
        etSend.text.clear()
    }

    private fun disconnect() {
        closeGatt()
        appendLog("[系统] 手动断开")
    }

    @SuppressLint("MissingPermission")
    private fun closeGatt() {
        gatt?.disconnect()
        gatt?.close()
        gatt = null
        txChar = null
        mainHandler.post { setUiConnected(false) }
    }

    private fun setUiConnected(connected: Boolean) {
        btnScan.isEnabled = !connected
        btnDisconnect.isEnabled = connected
        btnSend.isEnabled = connected
        etSend.isEnabled = connected
        tvStatus.text = if (connected) "● 已连接" else "● 未连接"
        tvStatus.setTextColor(if (connected) 0xFF4CAF50.toInt() else 0xFF888888.toInt())
    }

    private fun appendLog(msg: String, type: String = "info") {
        mainHandler.post {
            val color = when (type) {
                "tx" -> "#88ff88"
                "rx" -> "#88ccff"
                "err" -> "#ff8888"
                else -> "#cccccc"
            }
            val line = "<font color=\"$color\">$msg</font><br/>"
            tvLog.append(Html.fromHtml(line, Html.FROM_HTML_MODE_LEGACY))
            scrollLog.post { scrollLog.fullScroll(ScrollView.FOCUS_DOWN) }
        }
    }
}
