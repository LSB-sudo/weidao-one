# nodehub

这是一个多项目仓库。对于初次使用者，建议优先关注“怎么跑起来”，先启动 RDK 侧 Go 服务并确认网页可访问，再按需处理 Android、STM32 和视觉检测模块。

## 仓库组成

- `RDK/rdk_notic/rdk_notic/`：RDK X5 侧 Go 服务，提供视频预览页面和诊断接口。
- `Niagara_1/`：Android 控制端应用。
- `C15A主板测试代码/`：STM32 主板相关工程，以及一个 BLE Android 调试工程。
- `RDK/rdkx5_img_code/`：Python/ROS2 视觉检测模块。

## 快速开始

如果你的目标是先看到网页视频页面，建议按这个顺序：

1. 在板子上进入 `RDK/rdk_notic/rdk_notic/`
2. 启动 Go 服务
3. 浏览器访问 `http://<板子IP>:8080/viewer`
4. 如需手机控制，再处理 `Niagara_1/`
5. 如需主板联动或视觉检测，再处理 STM32 和 `RDK/rdkx5_img_code/`

## RDK Go 服务运行

目录：

```text
RDK/rdk_notic/rdk_notic/
```

推荐启动方式：

```bash
cd RDK/rdk_notic/rdk_notic
GPS_ENABLED=false ./scripts/run_webrtc_usb.sh
```

访问地址：

```text
http://<板子IP>:8080/viewer
```

诊断接口：

```text
http://<板子IP>:8080/health
http://<板子IP>:8080/devices
```

说明：

- 如果当前没有接 GPS，先用 `GPS_ENABLED=false` 启动。
- `VIDEO_VFLIP`、`VIDEO_HFLIP` 当前只作为环境变量读取，不应当写成命令行参数。

常见场景：

1. 只看视频，不接 GPS：

```bash
cd RDK/rdk_notic/rdk_notic
GPS_ENABLED=false ./scripts/run_webrtc_usb.sh
```

2. 画面方向不对：

- 检查环境变量 `VIDEO_VFLIP`
- 检查环境变量 `VIDEO_HFLIP`

## Niagara_1 Android 应用

目录：

```text
Niagara_1/
```

启动入口：

- `ControlActivity`（launcher Activity）

构建命令：

Windows：

```powershell
cd Niagara_1
.\gradlew.bat assembleDebug
```

Linux / macOS：

```bash
cd Niagara_1
./gradlew assembleDebug
```

APK 输出路径：

```text
Niagara_1/app/build/outputs/apk/debug/app-debug.apk
```

使用前需要修改的地址：

- `viewer_url`
- `niagara_url`

如果不改成你自己的设备地址，App 页面通常无法正常打开。

HTTP 明文访问注意事项：

- 当前页面访问使用的是 HTTP 明文地址。
- Android 侧需要注意 `usesCleartextTraffic` 相关配置是否允许明文流量。
- 如果后续改成 HTTPS 或改为域名访问，需要同步检查网络安全策略，而不只是修改 URL 字符串。

当前应用用途主要包括：

- 扫描 BLE 设备
- 连接 `FFE0/FFE1` 透传类 BLE 模块
- 发送控制协议 `RPM,L:x,R:y\n`
- 打开 RDK viewer 页面
- 打开 Niagara 页面

## STM32 主板工程

Keil 工程路径：

```text
C15A主板测试代码/USER/MiniBalance.uvprojx
```

建议使用 `Keil uVision` 打开、编译，并按你的硬件连接方式完成下载或烧录。

这里不写“已验证”的命令行烧录方式，因为当前任务只整理 README，不声称已在本机验证烧录流程。

## BLE Android 调试工程

目录：

```text
C15A主板测试代码/ble_android/ble_android/
```

这是一个标准 Gradle Android 工程。

可选使用方式：

- 直接用 Android Studio 打开该目录
- 根据目录中的 wrapper 或本机 Gradle 环境尝试构建

能否直接构建成功还取决于本机 Android SDK、JDK 和本地环境配置。

## RDK X5 视觉检测模块

目录：

```text
RDK/rdkx5_img_code/
```

这是一个 Python/ROS2 视觉检测模块，与 Go 服务不是同一条启动链路。

如果你只是想先打开 viewer 页面，不需要先启动这里的内容。

前提说明：

- 下述运行命令仅适用于已经安装 ROS2 与 RDK X5 相关依赖的目标机环境。
- 如果当前机器不是目标环境，需要你自行补齐依赖、设备和模型文件。

参数文件说明：

```text
RDK/rdkx5_img_code/camera_params.yaml
```

`camera_params.yaml` 可以作为参考或通过 ROS2 参数机制挂载使用，但它不是直接执行 Python 脚本时自动加载的配置文件。

可参考的运行入口：

```bash
python3 test.py
python3 running_camera.py
python3 visual.py
python3 bird_deterrent_serial.py
```

这些命令只适用于已安装 ROS2 与 RDK X5 依赖的目标机环境。

## 常见问题排查

### viewer 打不开

先检查：

- `http://<板子IP>:8080/health` 是否正常
- 服务是否真的监听在 `8080`

### devices 为空

通常表示摄像头没有被系统识别，优先检查：

- 摄像头是否接好
- 设备节点是否存在
- `VIDEO_DEVICE` 是否配置正确

### 无 GPS 设备

启动时先关闭 GPS：

```bash
GPS_ENABLED=false ./scripts/run_webrtc_usb.sh
```

### 画面方向异常

检查：

- `VIDEO_VFLIP`
- `VIDEO_HFLIP`

### Android 页面打不开

优先检查：

- `viewer_url`
- `niagara_url`

### BLE 已连接但电机无动作

优先检查：

- Android 权限是否完整
- BLE 服务和特征 UUID 是否仍为 `FFE0/FFE1`
- 下发协议是否符合 `RPM,L:x,R:y\n`
- STM32 固件是否正常运行

## 最小可用启动组合

如果你只想尽快看到系统有响应，可以先做这几步：

1. 在 RDK 板上进入 `RDK/rdk_notic/rdk_notic`
2. 运行 `GPS_ENABLED=false ./scripts/run_webrtc_usb.sh`
3. 浏览器打开 `http://<板子IP>:8080/viewer`
4. 需要手机侧控制时，再进入 `Niagara_1/` 构建 APK
5. 安装前先把 `viewer_url`、`niagara_url` 改成实际地址

## 目录速查

```text
nodehub/
|-- RDK/
|   |-- rdk_notic/rdk_notic/    # RDK X5 Go 服务
|   `-- rdkx5_img_code/         # RDK X5 Python / ROS2 视觉检测模块
|-- Niagara_1/                  # Android 控制 App
`-- C15A主板测试代码/           # STM32 工程与 BLE Android 调试工程
```
