# nodehub

This is a multi-project repository. For first-time users, it is recommended to prioritize "getting it up and running": start the RDK-side Go service and confirm the web page is accessible first, then work with the Android, STM32, and vision detection modules as needed.

## Repository Structure

- `RDK/rdk_notic/rdk_notic/`: RDK X5-side Go service, providing video preview pages and diagnostic interfaces.
- `Niagara_1/`: Android control application.
- `C15A主板测试代码/`: STM32 mainboard related projects, including a BLE Android debugging project.
- `RDK/rdkx5_img_code/`: Python/ROS2 vision detection module.

## Quick Start

If your goal is to see the video web page first, follow this order:

1. On the board, navigate to `RDK/rdk_notic/rdk_notic/`
2. Start the Go service
3. Access `http://<board-IP>:8080/viewer` in a browser
4. If mobile control is needed, proceed to `Niagara_1/`
5. If mainboard integration or vision detection is needed, work with STM32 and `RDK/rdkx5_img_code/`

## RDK Go Service Operation

Directory:

```text
RDK/rdk_notic/rdk_notic/
```

Recommended startup command:

```bash
cd RDK/rdk_notic/rdk_notic
GPS_ENABLED=false ./scripts/run_webrtc_usb.sh
```

Access address:

```text
http://<board-IP>:8080/viewer
```

Diagnostic endpoints:

```text
http://<board-IP>:8080/health
http://<board-IP>:8080/devices
```

Notes:

- If no GPS is currently connected, start with `GPS_ENABLED=false`.
- `VIDEO_VFLIP` and `VIDEO_HFLIP` are currently only read as environment variables, not intended to be written as command-line arguments.

Common scenarios:

1. View video only, no GPS connected:

```bash
cd RDK/rdk_notic/rdk_notic
GPS_ENABLED=false ./scripts/run_webrtc_usb.sh
```

2. Incorrect video orientation:

- Check environment variables `VIDEO_VFLIP`
- Check environment variables `VIDEO_HFLIP`

## Niagara_1 Android Application

Directory:

```text
Niagara_1/
```

Launch entry:

- `ControlActivity` (launcher Activity)

Build commands:

Windows:

```powershell
cd Niagara_1
.\gradlew.bat assembleDebug
```

Linux / macOS:

```bash
cd Niagara_1
./gradlew assembleDebug
```

APK output path:

```text
Niagara_1/app/build/outputs/apk/debug/app-debug.apk
```

Addresses to modify before use:

- `viewer_url`
- `niagara_url`

If not changed to your device's address, the App pages will usually not open properly.

HTTP plaintext access note:

- Current page access uses HTTP plaintext addresses.
- On the Android side, check if `usesCleartextTraffic` allows plaintext traffic.
- If later switching to HTTPS or domain access, check network security policies accordingly, not just URL strings.

Current application functions include:

- Scanning BLE devices
- Connecting to `FFE0/FFE1` transparent transmission BLE modules
- Sending control protocol `RPM,L:x,R:y\n`
- Opening the RDK viewer page
- Opening the Niagara page

## STM32 Mainboard Project

Keil project path:

```text
C15A主板测试代码/USER/MiniBalance.uvprojx
```

Recommended to open and compile with `Keil uVision`, and complete download or flashing according to your hardware connections.

No "verified" command-line flashing methods are listed here, as this README only organizes documentation and does not claim verified flashing workflows on local machines.

## BLE Android Debugging Project

Directory:

```text
C15A主板测试代码/ble_android/ble_android/
```

This is a standard Gradle Android project.

Optional usage:

- Open the directory directly in Android Studio
- Attempt to build using the wrapper or local Gradle environment

Successful build depends on local Android SDK, JDK, and environment configuration.

## RDK X5 Vision Detection Module

Directory:

```text
RDK/rdkx5_img_code/
```

This is a Python/ROS2 vision detection module, not part of the same startup chain as the Go service.

If you only want to open the viewer page first, you do not need to start this module first.

Prerequisites:

- The following commands only apply to target environments with ROS2 and RDK X5 dependencies installed.
- If your current machine is not the target environment, you need to install dependencies, devices, and model files yourself.

Parameter file note:

```text
RDK/rdkx5_img_code/camera_params.yaml
```

`camera_params.yaml` can be used as a reference or mounted via ROS2 parameter mechanisms, but it is not automatically loaded when running Python scripts directly.

Reference entry points:

```bash
python3 test.py
python3 running_camera.py
python3 visual.py
python3 bird_deterrent_serial.py
```

These commands only apply to target environments with ROS2 and RDK X5 dependencies installed.

## Troubleshooting

### Viewer page not opening

Check first:

- Is `http://<board-IP>:8080/health` responding normally?
- Is the service actually listening on port `8080`?

### Devices list empty

Usually means the camera is not recognized by the system. Prioritize checking:

- Camera connections
- Existence of device nodes
- Correct configuration of `VIDEO_DEVICE`

### No GPS device

Start with GPS disabled:

```bash
GPS_ENABLED=false ./scripts/run_webrtc_usb.sh
```

### Incorrect video orientation

Check:

- `VIDEO_VFLIP`
- `VIDEO_HFLIP`

### Android page not opening

Prioritize checking:

- `viewer_url`
- `niagara_url`

### BLE connected but no motor movement

Prioritize checking:

- Complete Android permissions
- BLE service and characteristic UUIDs still set to `FFE0/FFE1`
- Control protocol matches `RPM,L:x,R:y\n`
- STM32 firmware is running normally

## Minimum Working Startup Flow

If you just want to see the system respond quickly, follow these steps:

1. On the RDK board, navigate to `RDK/rdk_notic/rdk_notic`
2. Run `GPS_ENABLED=false ./scripts/run_webrtc_usb.sh`
3. Open `http://<board-IP>:8080/viewer` in a browser
4. When mobile control is needed, proceed to build the APK in `Niagara_1/`
5. Before installation, change `viewer_url` and `niagara_url` to the actual device addresses

## Directory Overview

```text
nodehub/
|-- RDK/
|   |-- rdk_notic/rdk_notic/    # RDK X5 Go service
|   `-- rdkx5_img_code/         # RDK X5 Python / ROS2 vision detection module
|-- Niagara_1/                  # Android control App
`-- C15A主板测试代码/           # STM32 project & BLE Android debugging project
```