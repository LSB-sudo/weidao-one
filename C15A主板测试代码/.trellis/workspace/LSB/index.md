# Workspace Index - LSB

> Journal tracking for AI development sessions.

---

## Current Status

<!-- @@@auto:current-status -->
- **Active File**: `journal-2.md`
- **Total Sessions**: 15
- **Last Active**: 2026-05-26
<!-- @@@/auto:current-status -->

---

## Active Documents

<!-- @@@auto:active-documents -->
| File | Lines | Status |
|------|-------|--------|
| `journal-1.md` | ~950 | Archived |
| `journal-2.md` | ~145 | Active |
<!-- @@@/auto:active-documents -->

---

## Session History

<!-- @@@auto:session-history -->
| # | Date | Title | Commits | Branch |
|---|------|-------|---------|--------|
| 1 | 2026-05-25 | 工程架构熟悉与分层文档整理 | - | `master` |
| 2 | 2026-05-25 | 串口 RPM 到编码器目标量映射接入 | `6c0169d` | `master` |
| 3 | 2026-05-25 | 电机通道裁剪为仅保留 C / D | - | `master` |
| 4 | 2026-05-25 | 串口增加 D 通道独立/联合控制 | - | `master` |
| 5 | 2026-05-25 | 按重构说明拆分电机控制 app 层第一步 | - | `master` |
| 6 | 2026-05-25 | 串口重构第一步：底层留 SYSTEM，用户逻辑抽到 target_app | - | `master` |
| 7 | 2026-05-25 | 串口重构第二步：保留 uart1_init 在 SYSTEM，中断处理迁入 usart_app | - | `master` |
| 8 | 2026-05-25 | 合并 target_app 到 usart_app | - | `master` |
| 9 | 2026-05-25 | IMU 姿态周期对齐与轻量显示滤波 | - | `master` |
| 10 | 2026-05-25 | IMU 稳定性问题继续排查（问题尚未解决） | - | `master` |
| 11 | 2026-05-25 | IMU P/R/Y 分通道修正与可接受方案收敛 | - | `master` |
| 12 | 2026-05-26 | IMU 分层重构第一步：硬件层与 app 层职责拆分 | - | `master` |
| 13 | 2026-05-26 | 串口3蓝牙接口最小化清理与串口4最小驱动新增 | - | `master` |
| 14 | 2026-05-26 | 删除 show / DataScope_DP 调试显示链路 | - | `master` |
| 15 | 2026-05-26 | OLED_Init 相关电机异常排查与当前保留方案 | - | `master` |
<!-- @@@/auto:session-history -->

---

## Notes

- Sessions are appended to journal files
- New journal file created when current exceeds 2000 lines
- Use `add_session.py` to record sessions
