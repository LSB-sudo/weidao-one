# Fix IMU Drift And Settling Issue

## Goal

Eliminate the current IMU attitude drift/slow-settling behavior so the board reaches a stable static attitude quickly after power-on and after motion stops.

## What I Already Know

- The unresolved problem is recorded in `.trellis/workspace/LSB/journal-1.md`.
- The current symptom is slow attitude drift at rest, for example values moving gradually from about `-20` to `-110`.
- `MiniBalance/ICM20948/bsp_imu.c` already changed `SamplePeriod` from `0.002` to `0.005` to match the 5 ms control loop.
- `MiniBalance/imu_app.c` already performs startup zero calibration with `IMU_App_CalibrateZero(200,5)`.
- `MiniBalance/CONTROL/control.c` already adds a lightweight display low-pass filter, but the journal states the drift existed before that filter.

## Requirements

- Keep the current control-loop period alignment at 5 ms.
- Fix attitude fusion behavior in the IMU path rather than relying on display-only smoothing.
- Preserve existing external IMU interfaces unless a small local extension is clearly necessary.
- Keep startup zero calibration behavior, but make it cooperate correctly with the fusion state.

## Acceptance Criteria

- [ ] Static attitude no longer slowly drifts across large angles after startup.
- [ ] After moving the board and placing it still, attitude converges back promptly instead of taking a very long time.
- [ ] Existing IMU read/update call flow still builds cleanly in the current project structure.

## Out Of Scope

- Full IMU driver rewrite
- Magnetometer hard/soft-iron calibration workflow
- PID/control-loop tuning unrelated to attitude estimation

## Technical Notes

- Likely files: `MiniBalance/ICM20948/bsp_imu.c`, `MiniBalance/imu_app.c`, `USER/MiniBalance.c`, `MiniBalance/CONTROL/control.c`
- High-probability root causes under review: fusion state initialization, integral accumulation/reset, and calibration-state interaction.
