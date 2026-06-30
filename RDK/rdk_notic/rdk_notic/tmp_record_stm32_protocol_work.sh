cd /root/rdk_notic || exit 1
python3 - <<'PY'
from pathlib import Path

journal = Path(".trellis/workspace/LSB/journal-1.md")
agents = Path("AGENTS.md")

journal_append = """

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
"""

agents_append = """

- 2026-06-23: STM32 lower-controller serial integration is now wired into the main RDK service. The default STM32 link uses `/dev/serial/by-id/usb-stm32_ttl_usb1-port0` at `115200 8N1`, enabled via `STM32_ENABLED=true`. The RDK-to-STM32 command frame is fixed as `CMD,left_set_rpm=<float>,right_set_rpm=<float>,boat_run=<0|1>`. The STM32-to-RDK feedback frame is fixed as `FB,left_rpm=<float>,right_rpm=<float>,battery_voltage=<float>`, with tolerant aliases for left/right RPM and battery fields documented in `mydocs/STM32_SERIAL_PROTOCOL.md`. Niagara topics `wd1/boat/sensor/motor_left_rpm`, `wd1/boat/sensor/motor_right_rpm`, and `wd1/boat/sensor/battery` now publish STM32 feedback values instead of fixed placeholders when valid feedback is present.
"""

journal.write_text(journal.read_text(encoding="utf-8") + journal_append, encoding="utf-8")
agents.write_text(agents.read_text(encoding="utf-8") + agents_append, encoding="utf-8")
PY
