# ble_android 项目说明

`ble_android` 是一个 Android BLE 串口调试 App。它可以扫描附近的 BLE 设备，连接使用 `FFE0/FFE1` 服务与特征 UUID 的透传模块，然后通过手机界面发送文本指令并查看设备返回的数据。

这个项目更像是一个轻量级 BLE 调试工具或硬件联调 Demo，适合用来验证 HM-10、HC-08 以及同类 BLE 串口透传模块是否能正常收发数据。

## 项目概览

- 项目名称：`BleApp`
- 应用包名：`com.example.bleapp`
- 开发语言：Kotlin
- 项目类型：原生 Android App
- 构建系统：Gradle / Android Gradle Plugin
- 最低 Android 版本：API 21
- 目标 Android 版本：API 34
- 核心入口：`app/src/main/java/com/example/bleapp/MainActivity.kt`
- 现有产物：`app/build/outputs/apk/debug/app-debug.apk`

## 目录结构

```text
ble_android/
├── build.gradle
├── settings.gradle
├── gradle.properties
├── local.properties
├── gradle/
│   └── wrapper/
│       └── gradle-wrapper.properties
└── app/
    ├── build.gradle
    └── src/main/
        ├── AndroidManifest.xml
        ├── java/com/example/bleapp/
        │   └── MainActivity.kt
        └── res/
            ├── layout/activity_main.xml
            └── values/themes.xml
```

`app/build/` 和 `.gradle/` 是构建产生的缓存与产物目录，理解源码时可以先忽略。

## 它能做什么

主界面是一个简单的 BLE 串口调试面板，包含：

- 扫描 BLE 设备
- 弹窗选择扫描到的设备
- 建立 GATT 连接
- 查找指定服务和特征
- 开启特征通知，用于接收设备数据
- 输入文本指令并发送
- 在日志区显示系统状态、发送内容和接收内容
- 手动断开连接

界面布局定义在 `activity_main.xml` 中，整体是深色背景，上方显示连接状态和扫描/断开按钮，中间是日志窗口，底部是输入框和发送按钮。

## BLE 通信参数

项目默认适配 HM-10 / HC-08 / 同类 BLE 串口透传模块常见的 UUID：

```kotlin
SERVICE_UUID = 0000ffe0-0000-1000-8000-00805f9b34fb
CHAR_UUID    = 0000ffe1-0000-1000-8000-00805f9b34fb
CCCD_UUID    = 00002902-0000-1000-8000-00805f9b34fb
```

含义大致是：

- `SERVICE_UUID`：BLE 串口服务
- `CHAR_UUID`：用于写入和接收通知的特征
- `CCCD_UUID`：通知开关描述符，用来开启 Notify

如果你的硬件模块不是 `FFE0/FFE1` 协议，需要修改 `MainActivity.kt` 里的这几个 UUID。

## 核心流程

1. 用户点击“扫描”。
2. App 检查并申请蓝牙/定位权限。
3. 调用 `bluetoothLeScanner.startScan()` 扫描 5 秒。
4. 扫描结果按 MAC 地址去重，显示设备名称和地址。
5. 扫描结束后弹出设备选择框。
6. 用户选择设备后调用 `connectGatt()` 建立 GATT 连接。
7. 连接成功后调用 `discoverServices()` 发现服务。
8. 查找 `FFE0` 服务下的 `FFE1` 特征。
9. 找到特征后开启通知，允许设备主动回传数据。
10. 用户输入文本后，App 自动追加 `\r\n` 并写入 BLE 特征。
11. 设备返回的数据按 `\n` 分行显示在日志区。

## 权限处理

代码对不同 Android 版本做了区分：

- Android 12 及以上：申请 `BLUETOOTH_SCAN` 和 `BLUETOOTH_CONNECT`
- Android 11 及以下：申请 `ACCESS_FINE_LOCATION`

`AndroidManifest.xml` 中还声明了：

- `BLUETOOTH`
- `BLUETOOTH_ADMIN`
- `ACCESS_FINE_LOCATION`
- `android.hardware.bluetooth_le`

其中 `android.hardware.bluetooth_le` 被设置为 `required="true"`，表示设备必须支持 BLE 才能安装或正常使用。

## 如何运行

推荐使用 Android Studio：

1. 打开 Android Studio。
2. 选择 `ble_android` 目录作为项目根目录。
3. 等待 Gradle 同步。
4. 连接一台支持 BLE 的 Android 真机。
5. 点击 Run 安装运行。

注意：BLE 扫描和连接通常需要真机，普通模拟器一般不适合做这个项目的完整验证。

也可以用命令行构建，但当前目录里只有 `gradle/wrapper/gradle-wrapper.properties`，没有看到 `gradlew` / `gradlew.bat` 启动脚本。如果本机已经安装 Gradle，可以在 `ble_android` 目录下尝试：

```powershell
gradle assembleDebug
```

构建成功后的 APK 通常位于：

```text
app/build/outputs/apk/debug/app-debug.apk
```

## 关键源码说明

`MainActivity.kt` 承担了几乎全部业务逻辑：

- `neededPermissions()`：根据 Android 版本返回需要申请的权限
- `startScan()`：检查权限并开始扫描
- `doScan()`：执行 BLE 扫描，扫描时长为 5 秒
- `showPicker()`：用弹窗展示扫描到的设备
- `connectDevice()`：连接用户选择的 BLE 设备
- `gattCallback`：处理连接状态、服务发现、数据通知
- `handleRx()`：接收字节数据，转为 UTF-8 文本，并按行显示
- `sendText()`：把输入框内容追加 `\r\n` 后写入特征
- `disconnect()` / `closeGatt()`：断开并释放 GATT 连接
- `appendLog()`：把日志追加到界面，最多保留 200 行

## 当前限制和注意点

- 只适配 `FFE0/FFE1` UUID 的 BLE 透传模块。
- 发送数据默认使用 UTF-8，并自动追加 `\r\n`。
- 接收数据按换行符 `\n` 分包显示，如果设备一直不发送换行，日志可能不会立即显示完整内容。
- 没有实现 MTU 协商，大数据包发送可能需要额外拆包。
- 写入特征时使用了旧版 `characteristic.value + writeCharacteristic()` API，能兼容旧系统，但在较新 Android API 上可以考虑升级为新版写入方式。
- 没有处理 `onCharacteristicWrite()`，所以当前日志里的 `[TX]` 表示已经发起写入，不代表外设一定成功收到。
- 没有检查蓝牙是否关闭、定位服务是否关闭等更细的系统状态。
- 没有扫描过滤器，会列出附近所有 BLE 设备。

## 适合的使用场景

- 调试 BLE 串口模块
- 给单片机、Arduino、ESP32、STM32 等设备发送简单文本指令
- 快速验证 BLE 透传链路是否正常
- 作为 Android BLE 扫描、连接、通知、写入流程的学习样例

## 与卫稻一号 APP 需求的关系

根目录的 `app_requirements.md` 描述的是完整“卫稻一号”机器人控制与管理 APP，目标包含 UniApp 前端、Flask 后端、MQTT、视频流、地图定位、任务规划、历史记录等模块。

本 `ble_android` 工程只覆盖其中的“BLE 近场通信”验证部分，可以视为正式 APP 的通信原型或硬件联调工具。它目前能够证明手机端可以通过 BLE 扫描、连接并向机器人或 BLE 透传模块收发文本指令，但还不是完整的卫稻一号控制端。

和需求文档相比，当前已经具备：

- BLE 设备扫描；
- BLE 设备选择和连接；
- GATT 服务发现；
- `FFE0/FFE1` 串口透传特征读写；
- 通知接收；
- 文本指令发送；
- 基础连接与收发日志。

当前尚未具备：

- 只识别卫稻一号机器人，而不是所有 BLE 设备；
- 批量设备管理；
- 断线自动重连；
- 虚拟摇杆和速度档位；
- 电量、速度、定位、信号强度等结构化状态显示；
- 与正式 UniApp 前端或 Flask/MQTT 后端集成；
- 视频、地图、任务规划、历史作业等业务功能。

后续如果要把它并入正式产品，建议先把 BLE 指令协议固定下来，再把 `MainActivity.kt` 里的 BLE 逻辑拆成独立通信模块。正式 UniApp 端可以选择使用 UniApp 自带蓝牙 API 重新实现，也可以通过 Android 原生插件复用这套 Kotlin BLE 逻辑。
