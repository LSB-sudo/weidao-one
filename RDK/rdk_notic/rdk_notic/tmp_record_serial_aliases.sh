cd /root/rdk_notic || exit 1
python3 - <<'PY'
from pathlib import Path

journal = Path(".trellis/workspace/LSB/journal-1.md")
agents = Path("AGENTS.md")

journal_append = """

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
"""

agents_append = """

- 2026-06-23: Board serial role mapping is fixed by custom udev aliases rather than transient tty numbering. User-confirmed role mapping is: GPS = USB0 = `/dev/serial/by-id/usb-gps_usb0-port0` -> `/dev/ttyUSB0`, STM32 lower-controller link = USB1 = `/dev/serial/by-id/usb-stm32_ttl_usb1-port0` -> `/dev/ttyUSB1`. The system-generated CH340 alias `usb-1a86_USB_Serial-if00-port0` currently points to `ttyUSB1` and must not be used as the authoritative role identifier for application logic. Future GPS / STM32 serial configuration in this workspace should prefer the custom aliases written by `/etc/udev/rules.d/99-rdk-serial-alias.rules`.
"""

journal.write_text(journal.read_text(encoding="utf-8") + journal_append, encoding="utf-8")
agents.write_text(agents.read_text(encoding="utf-8") + agents_append, encoding="utf-8")
PY
