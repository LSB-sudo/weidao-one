<!-- TRELLIS:START -->
# Trellis Instructions

These instructions are for AI assistants working in this project.

This project is managed by Trellis. The working knowledge you need lives under `.trellis/`:

- `.trellis/workflow.md` ? development phases, when to create tasks, skill routing
- `.trellis/spec/` ? package- and layer-scoped coding guidelines (read before writing code in a given layer)
- `.trellis/workspace/` ? per-developer journals and session traces
- `.trellis/tasks/` ? active and archived tasks (PRDs, research, jsonl context)

If a Trellis command is available on your platform (e.g. `/trellis:finish-work`, `/trellis:continue`), prefer it over manual steps. Not every platform exposes every command.

If you're using Codex or another agent-capable tool, additional project-scoped helpers may live in:
- `.agents/skills/` ? reusable Trellis skills
- `.codex/agents/` ? optional custom subagents

Managed by Trellis. Edits outside this block are preserved; edits inside may be overwritten by a future `trellis update`.

<!-- TRELLIS:END -->

## Engineering Decisions

- 2026-05-27: The first USB camera preview PoC lives in `cmd/rdk-webrtc`.
- USB cameras on the same LAN are the primary target; `/dev/video*` discovery must stay dynamic and must not hardcode `/dev/video0`.
- The board-side signaling/API surface is `GET /health`, `GET /devices`, `GET /viewer`, and `POST /offer`.
- WebRTC is implemented with Go and Pion, using a single outbound H.264 video track.
- The first transport pipeline is `ffmpeg` (V4L2 input, H.264 Annex-B to stdout) feeding a Go NALU parser and Pion track writer.
- No public STUN/TURN is enabled by default; the intended first deployment is same-LAN preview.
- Long-term engineering decisions for this remote workspace should continue to be recorded in this file.
- 2026-05-27: USB camera WebRTC preview tuning updated for lower latency and lower CPU pressure. Default stream profile changed to mjpeg + 640x480 + 15fps + 700kbps; ffmpeg now uses ultrafast, zerolatency, no B-frames, 1-second GOP, reduced buffering, and Annex-B headers. /health now exposes effective ffmpeg args, and /viewer now shows basic WebRTC getStats() diagnostics for decoded frames, dropped frames, bytes, jitter, and resolution.

- 2026-05-27: Annex-B access-unit aggregation was corrected so SPS/PPS stay with the following IDR and multi-slice VCL NALUs are not split into separate samples. The service still feeds full Annex-B access units into Pion `TrackLocalStaticSample` with `packetization-mode=1` baseline H.264.
- 2026-05-27: Browser H.264 decode stability is restored. Based on observed camera orientation, the WebRTC USB preview now defaults to ffmpeg `vflip`; `VIDEO_VFLIP`/`VIDEO_HFLIP` remain available for per-camera adjustment.
- 2026-05-27: H.264 decode is stable again. Based on observed camera orientation, the USB WebRTC preview now defaults to ffmpeg vflip, while VIDEO_VFLIP and VIDEO_HFLIP remain available for per-camera adjustment.
- 2026-05-27: The first Android app for this project lives in `android/RdkCameraViewer` and stays intentionally minimal: a single-Activity Android WebView shell loading `/viewer` over cleartext HTTP on the same LAN. Do not introduce a native Android WebRTC SDK unless stronger app-side control is needed later.
- 2026-06-01: OneNET integration decision for APP data access is fixed. APP-side data retrieval must use OneNET HTTP API or another application-side API and must not reuse the device MQTT identity. The device MQTT ClientID `FieldGuard_I` is exclusively owned by the RDK device session so OneNET property publish, platform replies, and downstream device topics remain stable.
- 2026-06-18: Remote operations for this workspace are standardized on `ssh-skill` and should no longer default to raw `plink` or ad hoc SSH commands. During the migration, three issues were confirmed: the local default `python` resolved to `Python 3.9.1` while `paramiko` was only installed in `Python 3.11`, the local `~/.ssh/config` entry for `192.168.3.142` did not include the `# password: ...` metadata required by `ssh-skill`, and PowerShell argument splitting made long remote command strings unreliable when calling `ssh_execute.py` directly. The accepted fix is: use `C:\Users\Lenovo\AppData\Local\Programs\Python\Python311\python.exe` as the local interpreter for `ssh-skill`, keep the local `Host 192.168.3.142` entry annotated with `# password: root` so `ssh-skill` can authenticate, and use a local helper wrapper when needed to pass multi-line or quote-heavy remote commands safely into `ssh_execute.py`. After this validation, `/root/rdk_notic` was successfully inspected through `ssh-skill`; future remote read/write/check/build operations for this workspace should continue to use `ssh-skill` as the default path unless a later decision recorded in this file explicitly replaces it.


- 2026-06-23: Board serial role mapping is fixed by custom udev aliases rather than transient tty numbering. User-confirmed role mapping is: GPS = USB0 = `/dev/serial/by-id/usb-gps_usb0-port0` -> `/dev/ttyUSB0`, STM32 lower-controller link = USB1 = `/dev/serial/by-id/usb-stm32_ttl_usb1-port0` -> `/dev/ttyUSB1`. The system-generated CH340 alias `usb-1a86_USB_Serial-if00-port0` currently points to `ttyUSB1` and must not be used as the authoritative role identifier for application logic. Future GPS / STM32 serial configuration in this workspace should prefer the custom aliases written by `/etc/udev/rules.d/99-rdk-serial-alias.rules`.


- 2026-06-23: STM32 lower-controller serial integration is now wired into the main RDK service. The default STM32 link uses `/dev/serial/by-id/usb-stm32_ttl_usb1-port0` at `115200 8N1`, enabled via `STM32_ENABLED=true`. The RDK-to-STM32 command frame is fixed as `CMD,left_set_rpm=<float>,right_set_rpm=<float>,boat_run=<0|1>`. The STM32-to-RDK feedback frame is fixed as `FB,left_rpm=<float>,right_rpm=<float>,battery_voltage=<float>`, with tolerant aliases for left/right RPM and battery fields documented in `mydocs/STM32_SERIAL_PROTOCOL.md`. Niagara topics `wd1/boat/sensor/motor_left_rpm`, `wd1/boat/sensor/motor_right_rpm`, and `wd1/boat/sensor/battery` now publish STM32 feedback values instead of fixed placeholders when valid feedback is present.
