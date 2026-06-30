# Journal - LSB (Part 1)

> AI development session journal
> Started: 2026-05-27

---



## Session 1: USB camera WebRTC PoC

**Date**: 2026-05-27
**Task**: USB camera WebRTC PoC
**Branch**: `master`

### Summary

Implemented minimal USB camera to WebRTC preview service with remote build and API smoke validation.

### Main Changes

- Added Go module and `cmd/rdk-webrtc` service entrypoint.
- Implemented `/health`, `/devices`, `/viewer`, and `/offer` with Pion WebRTC H.264 outbound track handling.
- Implemented dynamic `/dev/video*` discovery with `v4l2-ctl` summaries and capture-node filtering.
- Implemented ffmpeg V4L2 to Annex-B H.264 stdout pipeline plus NALU parsing and Pion sample writing.
- Added `scripts/run_webrtc_usb.sh`, project-local `.local-go` workflow note in README, and engineering decisions in `AGENTS.md`.
- Verified `gofmt`, `go build ./...`, and local HTTP smoke tests for `/health` and `/devices`.
- Current risk: full browser-to-board media preview was not exercised end-to-end in this shell-only session.


### Git Commits

(No commits - planning session)

### Testing

- [OK] (Add test results)

### Status

[OK] **Completed**

### Next Steps

- None - task complete

## [2026-05-27] WebRTC USB 摄像头卡顿调优回访

- 背景：用户反馈连接后非常卡顿，并提供 SDP answer，说明 WebRTC 信令/ICE 基本可用，问题集中在媒体链路。
- 已做调优摘要：默认从 1280x720@30 降到 mjpeg 640x480@15、700kbps；ffmpeg 加强低延迟参数；viewer 增加 stats；health 暴露 ffmpeg 参数。
- 用户反馈：延迟卡顿有一定改善。
- 后续建议：继续观察 viewer stats，若仍卡可降到 640x360@10、500kbps；下一步重点看 framesDropped/jitter/CPU。

## [21:00:08] WebRTC USB 摄像头卡顿调优回访

- 背景：用户反馈连接后非常卡顿，并提供 SDP answer，说明 WebRTC 信令/ICE 基本可用，问题集中在媒体链路。
- 已做调优摘要：默认从 1280x720@30 降到 mjpeg 640x480@15、700kbps；ffmpeg 加强低延迟参数；viewer 增加 stats；health 暴露 ffmpeg 参数。
- 用户反馈：延迟卡顿有一定改善。
- 后续建议：继续观察 viewer stats，若仍卡可降到 640x360@10、500kbps；下一步重点看 framesDropped/jitter/CPU。

## [CST] WebRTC USB 摄像头卡顿调优回访

- 背景：用户反馈连接后非常卡顿，并提供 SDP answer，说明 WebRTC 信令/ICE 基本可用，问题集中在媒体链路。
- 已做调优摘要：默认从 1280x720@30 降到 mjpeg 640x480@15、700kbps；ffmpeg 加强低延迟参数；viewer 增加 stats；health 暴露 ffmpeg 参数。
- 用户反馈：延迟卡顿有一定改善。
- 后续建议：继续观察 viewer stats，若仍卡可降到 640x360@10、500kbps；下一步重点看 framesDropped/jitter/CPU。

  echo ## [2026-05-27 21:00:33 CST] WebRTC USB 摄像头卡顿调优回访
  echo 
- 背景：用户反馈连接后非常卡顿，并提供 SDP answer，说明 WebRTC 信令/ICE 基本可用，问题集中在媒体链路。
- 已做调优摘要：默认从 1280x720@30 降到 mjpeg 640x480@15、700kbps；ffmpeg 加强低延迟参数；viewer 增加 stats；health 暴露 ffmpeg 参数。
- 用户反馈：延迟卡顿有一定改善。
- 后续建议：继续观察 viewer stats，若仍卡可降到 640x360@10、500kbps；下一步重点看 framesDropped/jitter/CPU。

## [2026-05-27 21:00:51 CST] WebRTC USB 摄像头卡顿调优回访

- 背景：用户反馈连接后非常卡顿，并提供 SDP answer，说明 WebRTC 信令/ICE 基本可用，问题集中在媒体链路。
- 已做调优摘要：默认从 1280x720@30 降到 mjpeg 640x480@15、700kbps；ffmpeg 加强低延迟参数；viewer 增加 stats；health 暴露 ffmpeg 参数。
- 用户反馈：延迟卡顿有一定改善。
- 后续建议：继续观察 viewer stats，若仍卡可降到 640x360@10、500kbps；下一步重点看 framesDropped/jitter/CPU。
## [2026-05-27 21:35 CST] WebRTC USB ???????

- ???H.264 ????????????????????????????
- ??????????? `VIDEO_VFLIP=true`???? `VIDEO_HFLIP` ?????????`/health` ????????? ffmpeg ???
- ??????????????????????????????????????????????
## [2026-05-27 21:35 CST] WebRTC USB camera orientation fix

- H.264 decode is working again; the remaining user-visible issue was upside-down video.
- Default behavior now enables VIDEO_VFLIP=true, keeps VIDEO_HFLIP available, and exposes flip settings plus final ffmpeg args in /health.
- The current default assumes the tested camera mount needs vertical flip and can be adjusted per deployment with env vars.


## [\] WebRTC USB 摄像头画面问题解决确认

- 问题回顾：先前 WebRTC USB 摄像头 H.264 画面仅上半部分正常、下半部分出现竖条/灰块；随后在画面完整后，又出现上下颠倒现象。
- 修复摘要：已修复 H.264 access unit 聚合逻辑，避免多 slice 同帧被误拆；保持 Annex-B access unit 输入 Pion；增加 stats/health 诊断；增加 VIDEO_VFLIP / VIDEO_HFLIP，并默认 VIDEO_VFLIP=true。
- 用户反馈：用户确认 ok解决问题了。
- 当前状态：WebRTC USB 摄像头实时预览第一版进入可用状态；后续若更换摄像头或安装方向变化，可通过环境变量调整翻转。

## [2026-05-27 21:49:41 CST] WebRTC USB 摄像头画面问题解决确认

- 问题回顾：先前 WebRTC USB 摄像头 H.264 画面仅上半部分正常、下半部分出现竖条/灰块；随后在画面完整后，又出现上下颠倒现象。
- 修复摘要：已修复 H.264 access unit 聚合逻辑，避免多 slice 同帧被误拆；保持 Annex-B access unit 输入 Pion；增加 stats/health 诊断；增加 VIDEO_VFLIP / VIDEO_HFLIP，默认 VIDEO_VFLIP=true。
- 用户反馈：用户确认 ok解决问题了。
- 当前状态：WebRTC USB 摄像头实时预览第一版进入可用状态；后续若更换摄像头或安装方向变化，可通过环境变量调整翻转。

## [2026-05-27 22:05 CST] Android WebView APK viewer scaffold

- Added `android/RdkCameraViewer`, a minimal Android Studio project for packaging the existing `/viewer` page into an APK.
- The first version uses a single Java Activity plus WebView, with JavaScript, DOM storage, WebChromeClient, WebViewClient, autoplay-friendly media settings, and cleartext HTTP enabled.
- `viewer_url` is centralized in `app/src/main/res/values/strings.xml`; current default is `http://192.168.3.142:8080/viewer`.
- The board-side Go/WebRTC service was left unchanged; Android Studio on a development machine is the intended build path.

## [2026-05-27 22:40:40 CST] Android Viewer APK 构建问题解决确认

- 背景：用户将 Android 工程迁移到 C:\Users\Lenovo\Desktop\android\RdkCameraViewer 后，Android Studio 构建先遇到 AGP 插件仓库缺 google()，后又遇到 Kotlin stdlib duplicate classes。
- 修复摘要：本地 settings.gradle 已补充 pluginManagement.repositories 和 dependencyResolutionManagement.repositories，包含 google()、mavenCentral()、gradlePluginPortal()；本地 gradle.properties 已添加 ndroid.suppressUnsupportedCompileSdk=35；本地 pp/build.gradle 已强制 org.jetbrains.kotlin:kotlin-stdlib* 统一到 1.8.22；ssembleDebug 已成功，APK 生成于 C:\Users\Lenovo\Desktop\android\RdkCameraViewer\app\build\outputs\apk\debug\app-debug.apk。
- 用户反馈：确认 ok 没问题了。
- 同步状态：用户说明远端代码也已经更新。
- 当前状态：Android WebView Viewer 第一版可以打包 APK，用于手机端显示 /viewer。

## [2026-05-27 22:40:40 CST] Android Viewer APK 构建问题解决确认（格式修正）

- 背景：用户将 Android 工程迁移到 C:\Users\Lenovo\Desktop\android\RdkCameraViewer 后，Android Studio 构建先遇到 AGP 插件仓库缺少 google()，后又遇到 Kotlin stdlib duplicate classes。
- 修复摘要：本地 settings.gradle 已补充 pluginManagement.repositories 和 dependencyResolutionManagement.repositories，包含 google()、mavenCentral()、gradlePluginPortal()；本地 gradle.properties 已添加 android.suppressUnsupportedCompileSdk=35；本地 app/build.gradle 已强制 org.jetbrains.kotlin:kotlin-stdlib* 统一到 1.8.22；assembleDebug 已成功，APK 生成于 C:\Users\Lenovo\Desktop\android\RdkCameraViewer\app\build\outputs\apk\debug\app-debug.apk。
- 用户反馈：确认 ok 没问题了。
- 同步状态：用户说明远端代码也已经更新。
- 当前状态：Android WebView Viewer 第一版可以打包 APK，用于手机端显示 /viewer。

## [2026-05-27 23:17:41 CST] RDK 蓝牙与 WebRTC 承载方式讨论

- 主题：RDK 蓝牙与 WebRTC 承载方式讨论。
- 问题：RDK 本身是否有蓝牙，是否可以用 RDK 蓝牙直接完成 WebRTC。
- 结论：
  - RDK 是否有蓝牙需要板上命令确认，不能假设。
  - 即使有蓝牙，也不建议用蓝牙直接承载 WebRTC 实时视频。
  - WebRTC 视频仍应走 Wi-Fi、以太网等 IP 网络。
  - 蓝牙更适合做设备发现、IP 同步、配网、控制命令、状态同步。
- 推荐架构：
  - 视频链路：RDK X5 -> Wi-Fi/WebRTC -> 手机 App。
  - 地址发现/控制链路：RDK X5 -> 串口 -> 32 -> 蓝牙 -> 手机 App；或未来 RDK 自带蓝牙 -> 手机 App。
- 后续建议：如果要确认 RDK 蓝牙能力，可只读检查 rfkill list、hciconfig -a、bluetoothctl list、lsusb、dmesg | grep -Ei ''bluetooth|bt|hci''。

## [2026-05-29 17:10 CST] ATGM336H GPS 串口接入与接口验证记录

- 背景：本次在 /root/rdk_notic 完成 ATGM336H GPS 串口接入，并补齐运行验证、接口观察与使用文档记录。
- 串口接入：当前默认设备为 /dev/serial/by-id/usb-1a86_USB_Serial-if00-port0，默认波特率 9600，运行开关为 GPS_ENABLED=true。
- 服务接口：已接入 GET /gps；GET /health 已增加 gps 摘要信息，便于上层健康检查与状态观察。
- 数据能力：服务已支持解析并缓存 NMEA 数据，当前至少覆盖 RMC、GGA、VTG、TXT、ZDA；接口可返回 connected、valid、stale、antennaStatus、lastSentence、fix 等字段。
- 构建与测试：./.local-go/go/bin/go test ./... 已通过，./.local-go/go/bin/go build ./... 已通过；运行态验证中 /gps 与 /health 均已访问成功。
- 实测状态：当前串口通信正常，已观察到 JSON 状态 connected=true、valid=false、stale=false、lastSentenceType=ZDA、lastSentence=$GNZDA,,,,,,*56。这说明模块串口可读，但当前尚未获得有效定位。
- 使用文档：远端 /root/rdk_notic/GPS_USAGE.md 已新增或更新，内容覆盖启动方式、接口说明、字段解释、排查建议与上层调用建议。
- 访问口径：远端历史上存在多个服务实例，8080 为当前新服务端口，18080 为较早测试实例。浏览器验证时应访问 http://192.168.3.142:8080/gps，不要误用本机 127.0.0.1。
- 当前结论：ATGM336H 接入链路、服务接口、构建测试与基本串口读取均已打通；当前主要未完成项是等待天线在空旷环境下获得有效定位。
- 后续关注：后续需将天线放到空旷处继续观察，确认 valid=true 以及经纬度等定位字段变为有效；如需提升上层可读性，可继续增加更友好的 GPS 状态枚举。

## [2026-05-29 22:13:50 CST] MQTT GPS Phase 1 实现与验证记录

- 背景：本次在 /root/rdk_notic 完成 MQTT Phase 1 最小子集，实现 RDK X5 定时发布 GPS 快照到公共 Broker，供上位机或 MQTT.fx 订阅观察。
- Broker 与 Topic：当前使用公共 Broker broker.emqx.io:1883，发布 Topic 为 devices/rdk-x5-001/gps。
- 默认配置：MQTT_ENABLED=false、MQTT_BROKER_URL=tcp://broker.emqx.io:1883、MQTT_DEVICE_ID=rdk-x5-001、MQTT_CLIENT_ID=rdk-x5-001-rdk、MQTT_PUBLISH_INTERVAL_SEC=5。
- 实现方式：复用现有 gps.Service 的 snapshot 能力，由 MQTT service 定时拉取并发布 GPS payload，不重复打开串口，避免与现有 GPS 读取链路发生资源竞争。
- 相关变更文件：internal/mqttclient/service.go、cmd/rdk-webrtc/main.go、internal/server/server.go、go.mod、go.sum。
- payload 字段：当前发布内容包含 deviceId、ts、source=gps-service、connected、valid、stale、antennaStatus、lastSentenceType、lastSentence、fix。
- 构建与测试：./.local-go/go/bin/go test ./... 已通过，./.local-go/go/bin/go build ./... 已通过。
- 真实验证：远端进程已确认可以连接 broker.emqx.io:1883，并通过短时订阅确认收到 devices/rdk-x5-001/gps 的实际 payload，说明 GPS 串口读取到 MQTT 发布的整条链路已打通。
- 当前 GPS 状态：目前实测 payload 多数仍为 valid=false，主要原因是 GPS 模块尚未获得有效定位；但 connected、串口读取与 MQTT 发布链路均已验证正常。
- 启动方式：先进入 /root/rdk_notic，然后以 MQTT_ENABLED=true、MQTT_BROKER_URL=tcp://broker.emqx.io:1883、MQTT_DEVICE_ID=rdk-x5-001、MQTT_CLIENT_ID=rdk-x5-001-rdk、MQTT_PUBLISH_INTERVAL_SEC=5 启动 ./.local-go/go/bin/go run ./cmd/rdk-webrtc。
- 观察方式：可使用 MQTT.fx 连接 broker.emqx.io:1883，订阅 devices/rdk-x5-001/#，即可观察当前 GPS 发布消息。
- 风险记录：当前 Phase 1 方案基于公共 broker、明文传输、无鉴权、QoS 0，仅适合快速联调与演示验证；后续生产化需要切换到 TLS、ACL、私有 Broker 或云 IoT 方案。
- 环境补充：Go 模块下载阶段曾通过 GOPROXY=https://goproxy.cn,direct 成功，后续若远端依赖拉取再次受限，可优先沿用该配置。

## [2026-06-01 09:00 CST] OneNET APP ??????????

- ????? OneNET ??? MQTT ? APP ?????????????????? MQTT ???? `ClientID=FieldGuard_I`?????????? RDK ? APP/MQTT.fx ????????? RDK ????
- ???APP ?? OneNET ???? HTTP API / ??????????? MQTT ????? MQTT ??? `ClientID=FieldGuard_I` ? RDK ???
- ??????`MQTT_RDK_X5_APP_DESIGN.md` ?????????? RDK ???? MQTT ????APP ?? OneNET HTTP API ??????/????????????????backend?rule engine/data forwarding?WebSocket??
- ???????`AGENTS.md` ??? OneNET APP ?????????????????????????
- ?????`mydocs/ONENET_APP_HTTP_API_ACCESS.md` ???????? APP ?????????????????RDK/APP ?????? API ????????????????????
- ????????????????????????README???? `Agents.md` ??????
- ??????? OneNET ??? HTTP API ??? endpoint??????token ???????????????????????????? `battery_voltage`??????????

## [2026-06-02 21:40 CST] DFPlayer MP3 ??????????

- ???????? USB ? TTL ????? DFPlayer Mini ???????????? 9600 8N1???? TTL_TX->DFPlayer_RX?TTL_RX->DFPlayer_TX?? GND?DFPlayer ?????
- ?????? `/root/rdk_notic` ?????????? Go ????????????? DFPlayer ??????????? 20????? 0001/0002/0003 ??????????? `rdk-webrtc` ????
- ?????`cmd/dfplayer-demo/main.go`?
- ?????
  - ??????????? `/dev/serial/by-id/usb-1a86_USB_Serial-if00-port0`???? `/dev/ttyUSB*`?`/dev/ttyACM*`?
  - ? 9600 8N1 ?????
  - ?????? `7E FF 06 0C 00 00 00 FE EF EF`??? 200ms?
  - ????? 20 ? `7E FF 06 06 00 00 14 FE E1 EF`??? 100ms?
  - ???? 0001/0002/0003 ????????????????????? `7E FF 06 16 00 00 00 FE E5 EF`?
  - ?? `-device`?`-baud`?`-play-duration`?`-cycles` ????????????
- ?????`./.local-go/go/bin/go build ./cmd/dfplayer-demo` ????
- ?????????????????????????????? reset?volume=20?play 0001?stop 0001?play 0002 ???????????????? USB-TTL???????? CH341 USB ????????????? I/O error ??????????????
- ??????? MP3 ??????????????????????????????
- ???????????? HTTP ???????????????????????????? `internal/mp3` ?????????

## [2026-06-18 11:30 CST] ssh-skill remote workflow, static viewer page, and Niagara embedding path

- Background: this session standardized remote board operations on `ssh-skill`, then extended the existing WebRTC camera preview so it can be embedded both through a board-hosted static page and through an external standalone HTML page.
- ssh-skill compatibility findings on the local Windows host:
  - bare `python` resolved to Python 3.9.1 while `paramiko` required by `ssh-skill` was available in Python 3.11;
  - the local `~/.ssh/config` alias needed `# password: ...` metadata for password authentication;
  - direct PowerShell calls to `ssh_execute.py` were unreliable for long or quote-heavy remote commands.
- Accepted local execution path: use `C:\Users\Lenovo\AppData\Local\Programs\Python\Python311\python.exe` with local wrappers to drive `ssh-skill`; keep password metadata in local SSH config; preserve `ssh-skill` as the default remote execution path for this workspace.
- Viewer implementation findings:
  - the original board viewer page is not a standalone disk HTML file;
  - `/viewer` is served from embedded Go source in `internal/server/viewer.go` and returned by `internal/server/server.go`.
- Remote functional additions completed:
  - added board-served static page `web/viewer-static.html`;
  - added route handler `internal/server/viewer_static.go` and route `/viewer-static`;
  - added `internal/server/cors.go` so `/offer`, `/health`, `/devices`, `/viewer`, and `/viewer-static` answer with permissive CORS headers and handle `OPTIONS` preflight.
- External embedding deliverable completed:
  - created local standalone page `viewer_external.html` that accepts a board base URL and connects by absolute URL instead of relying on same-origin `/offer`.
- Verification:
  - remote `./.local-go/go/bin/go build ./...` passed after the CORS and static viewer changes;
  - board-served `/viewer-static` and standalone `viewer_external.html` were both validated as working for camera display;
  - Niagara embedding path was confirmed working with the standalone external HTML approach.
- Current outcome:
  - existing `/viewer` remains intact;
  - board-hosted `/viewer-static` is available as an independent page;
  - Niagara or other external containers can embed the standalone HTML page and connect to the board over WebRTC through absolute URLs.

## [2026-06-23 09:20 CST] Niagara MQTT WD1 topic protocol integration

- Background: the local Niagara MQTT document defines a fixed WD1 topic contract with separate sensor publish topics and control subscribe topics. The existing RDK MQTT code only supported a single publish topic plus an optional single subscribe topic.
- Scope locked for this session:
  - only complete MQTT interaction at the protocol layer;
  - do not connect command topics to real lower-controller actions yet;
  - GPS should be taken from the existing raw GPS snapshot path and reduced to the Niagara-required effective field format.
- Implemented MQTT mode:
  - added `niagara_wd1` payload mode in `internal/mqttclient/service.go`;
  - kept legacy `gps` and `onenet_battery_voltage` modes intact.
- Implemented fixed publish topics:
  - `wd1/boat/sensor/battery`
  - `wd1/boat/sensor/motor_left_rpm`
  - `wd1/boat/sensor/motor_right_rpm`
  - `wd1/boat/sensor/bird_alarm`
  - `wd1/boat/sensor/bird_scare_status`
  - `wd1/boat/sensor/boat_run_status`
  - `wd1/boat/sensor/gps_pos`
- Implemented fixed subscribe topics:
  - `wd1/boat/cmd/motor_left_set`
  - `wd1/boat/cmd/motor_right_set`
  - `wd1/boat/cmd/boat_run`
- Current publish behavior:
  - GPS is published as `longitude,latitude` when `snapshot.Valid=true` and a fix exists;
  - when GPS is invalid or absent, `gps_pos` publishes an empty string;
  - battery, motor rpm, bird alarm, bird scare status, and boat run status currently publish safe placeholder defaults because this session did not yet wire those points to real hardware feedback sources.
- Current subscribe behavior:
  - command topics are parsed, logged, and cached in memory as the last received values with timestamps and raw payloads;
  - no real motor control, no boat run control, and no serial forwarding to the lower controller are performed in this iteration.
- CLI/config update:
  - `cmd/rdk-webrtc/main.go` now documents `niagara_wd1` as a valid `MQTT_PAYLOAD_MODE` option.
- Verification:
  - remote `./.local-go/go/bin/go build ./...` passed after the MQTT changes.
- Follow-up note:
  - a later task should connect cached command topics and non-GPS sensor placeholders to real lower-controller or hardware state sources.


## [2026-06-23 20:45 CST] Niagara MQTT 下行浮点命令兼容修复

- 背景：Niagara 已经可以把控制命令下发到 RDK，但 `wd1/boat/cmd/motor_left_set` 和 `wd1/boat/cmd/motor_right_set` 实际发送的是浮点数字符串，例如 `249.116`、`-65.51599999999999`。
- 问题表现：RDK 端日志出现 `strconv.Atoi` 解析失败，说明 MQTT 订阅链路正常，但下行数值命令被错误地按整数处理。
- 代码修复：已将 `internal/mqttclient/service.go` 中 Niagara 数值命令缓存结构 `NiagaraNumericCommand.Value` 从 `int` 调整为 `float64`，并将两个电机命令的解析从 `strconv.Atoi` 改为 `strconv.ParseFloat(raw, 64)`。
- 日志行为：修复后，电机命令将以浮点值缓存并记录，例如 `mqtt: cached niagara command topic=wd1/boat/cmd/motor_left_set value=249.116`。
- 验证：已在远端执行 `./.local-go/go/bin/gofmt -w internal/mqttclient/service.go` 和 `./.local-go/go/bin/go build ./...`，构建通过。
- 结果：Niagara 到 RDK 的下行控制 Topic 现已兼容浮点数值格式，后续无需强制 Niagara 侧改成整数发布。


## [2026-06-23 21:20 CST] GPS / STM32 USB 串口稳定别名纠正

- 背景：板端同时接入两根 CH340 USB 转串口线，需要把 GPS 与 STM32 下位机通讯口长期稳定地区分开，避免继续依赖会漂移的 `ttyUSB0` / `ttyUSB1` 编号。
- 用户最终确认映射：
  - GPS 对应 USB0，要求固定到 `ttyUSB0`
  - STM32 对应 USB1，要求固定到 `ttyUSB1`
- 只读排查结果：
  - 当前两根串口都表现为 `1a86:7523 QinHeng Electronics CH340 serial converter`
  - 物理口位可区分为：
    - `1-1.2` -> `ttyUSB0`
    - `1-1.1` -> `ttyUSB1`
  - 系统原生别名 `usb-1a86_USB_Serial-if00-port0` 当前指向 `ttyUSB1`，不适合继续用作 GPS/STM32 角色区分依据。
- 远端处理：
  - 新增 `udev` 规则文件 `/etc/udev/rules.d/99-rdk-serial-alias.rules`
  - 规则按用户确认的人为角色进行别名绑定，而不是沿用系统原生别名语义
  - 执行了 `udevadm control --reload-rules` 与 `udevadm trigger`
- 最终稳定别名：
  - GPS：`/dev/serial/by-id/usb-gps_usb0-port0` -> `/dev/ttyUSB0`
  - STM32：`/dev/serial/by-id/usb-stm32_ttl_usb1-port0` -> `/dev/ttyUSB1`
- 决策：
  - 后续工程代码和脚本中，GPS 与 STM32 串口应优先使用上述两个自定义稳定别名
  - 不应再依赖 `ttyUSB0` / `ttyUSB1` 的瞬时编号
  - 也不应再依赖系统原生 `usb-1a86_USB_Serial-if00-port0` 去表达设备角色


## [2026-06-23 22:15 CST] STM32 串口协议接入、转速回传与电池电压上报

- 背景：RDK 已具备 Niagara MQTT 上下行骨架，但此前下行转速命令只停留在内存缓存，实际没有通过串口发给 STM32；同时上行 `motor_left_rpm`、`motor_right_rpm`、`battery` 也还没有真实硬件反馈来源。
- 本次目标：
  - RDK 收到 Niagara 下发的左右设计转速和整机运行位后，通过 STM32 串口下发
  - RDK 接收 STM32 回传的实际左右转速
  - RDK 接收 STM32 回传的电池电压，并发布到 Niagara 现有电池 Topic
- 新增远端实现：
  - 新增 `internal/stm32link/serial_linux.go`，复用 Linux 串口打开与 8N1 配置逻辑
  - 新增 `internal/stm32link/service.go`，负责 STM32 串口连接、命令下发、反馈读回、状态缓存
  - 新增协议文档 `mydocs/STM32_SERIAL_PROTOCOL.md`
  - `cmd/rdk-webrtc/main.go` 新增 STM32 配置入口：
    - `STM32_ENABLED`
    - `STM32_DEVICE`
    - `STM32_BAUD`
    - `STM32_STALE_AFTER_SEC`
  - `internal/server/server.go` 接入 STM32 服务，并在 `/health` 中暴露 `stm32` 状态摘要
  - `internal/mqttclient/service.go` 已接入：
    - Niagara 下行命令缓存后推送到 STM32 串口服务
    - STM32 实际转速与电池电压回填到 Niagara 发布
- 当前协议固定为文本行协议：
  - RDK -> STM32：
    - `CMD,left_set_rpm=<float>,right_set_rpm=<float>,boat_run=<0|1>`
  - STM32 -> RDK：
    - `FB,left_rpm=<float>,right_rpm=<float>,battery_voltage=<float>`
- 当前兼容策略：
  - 反馈帧头兼容 `FB` / `STATE` / `RPM`
  - 左转速字段兼容：
    - `left_rpm`
    - `left_actual_rpm`
    - `l`
    - `left`
  - 右转速字段兼容：
    - `right_rpm`
    - `right_actual_rpm`
    - `r`
    - `right`
  - 电池字段兼容：
    - `battery_voltage`
    - `battery`
    - `bat_v`
    - `voltage`
- Niagara 发布话题已实际接入：
  - `wd1/boat/sensor/motor_left_rpm` <- STM32 实际左转速
  - `wd1/boat/sensor/motor_right_rpm` <- STM32 实际右转速
  - `wd1/boat/sensor/battery` <- STM32 实际电池电压
- 默认 STM32 串口配置：
  - 设备：`/dev/serial/by-id/usb-stm32_ttl_usb1-port0`
  - 波特率：`115200`
  - 格式：`8N1`
- 构建验证：
  - 远端 `./.local-go/go/bin/go build ./...` 已通过
- 当前结论：
  - Niagara -> RDK -> STM32 的串口下发链路代码已接通
  - STM32 -> RDK -> Niagara 的左右实际转速与电池电压回传链路代码已接通
  - 运行时还需要 STM32 端按协议实际发送反馈帧，才能在 MQTT 日志与 Niagara 侧看到真实值
