<!-- TRELLIS:START -->
# Trellis Instructions

These instructions are for AI assistants working in this project.

This project is managed by Trellis. The working knowledge you need lives under `.trellis/`:

- `.trellis/workflow.md` — development phases, when to create tasks, skill routing
- `.trellis/spec/` — package- and layer-scoped coding guidelines (read before writing code in a given layer)
- `.trellis/workspace/` — per-developer journals and session traces
- `.trellis/tasks/` — active and archived tasks (PRDs, research, jsonl context)

If a Trellis command is available on your platform (e.g. `/trellis:finish-work`, `/trellis:continue`), prefer it over manual steps. Not every platform exposes every command.

If you're using Codex or another agent-capable tool, additional project-scoped helpers may live in:
- `.agents/skills/` — reusable Trellis skills
- `.codex/agents/` — optional custom subagents

Managed by Trellis. Edits outside this block are preserved; edits inside may be overwritten by a future `trellis update`.

<!-- TRELLIS:END -->


# Project Decisions

## Git Commit Policy

Do not automatically create git commits for future work in this project.

Implementation and verification can still be completed normally, but commits should only be made when the user explicitly requests a commit for that specific change.

When reporting completed work, summarize changed files and verification results instead of committing by default.

## Skill Usage

For future embedded firmware work in this project, use the `embedded-dev` skill before implementation.

Apply this especially to STM32 firmware changes involving peripherals, GPIO, timers, PWM, interrupts, ADC, UART, IIC/I2C, SPI, DMA, driver porting, pin planning, datasheets, board resources, or hardware debugging.

The skill should be used together with the existing Trellis workflow and project layering rules. Keep edits outside the Trellis-managed block so future `trellis update` operations do not overwrite this project decision.

## Architecture Layering

For this STM32 project, use the following long-term layering convention when reading, modifying, or adding code.

### User Logic Control Layer

Directories:

- `USER/`
- `MiniBalance/`

Responsibilities:

- Main program entry and initialization orchestration
- Control logic and control-flow decisions
- Attitude / posture processing and related algorithms
- PID / balance / position control
- Display composition logic
- Debug data organization for host tools

Guidelines:

- Prefer expressing business intent and control decisions here
- Avoid placing direct GPIO / register / bus timing code here unless there is a very strong reason
- Display logic belongs here, but display hardware timing and low-level OLED driving belong in hardware-layer modules

### Hardware Component Library Layer

Directories:

- `MiniBalance_HARDWARE/`
- `SYSTEM/`
- `MiniBalance_COER/`
- `STM32F10x_FWLib/`

Responsibilities:

- Board-level peripheral drivers
- Motor, encoder, ADC, OLED, key, LED, buzzer, USART, IIC, flash access
- Delay, bit-operation, and low-level system support
- Cortex-M3 startup, core support, and STM32 standard peripheral library integration

Guidelines:

- Expose stable interfaces upward to the logic layer
- Encapsulate register details, timing details, and hardware access details here
- Keep each hardware module focused on a single peripheral or hardware concern
- The logic layer should not bypass this layer to manipulate hardware directly

## Current Architecture Understanding

Unless later code changes prove otherwise, treat the current project as a typical:

- interrupt-driven control architecture
- foreground `while(1)` loop for display and debug output

In practice:

- time-critical control work is expected to run in timer / interrupt driven paths
- non-real-time tasks such as OLED display and host-side data output are expected to run in the main loop

## Maintenance Preference

When adding new modules, decide first whether they belong to:

- control / logic / algorithm space
- hardware driver / board support space

Keep this separation stable over time instead of mixing hardware details into control modules.

## UART Print API Convention

For future formatted UART / serial debug output, prefer the shared app-layer API:

- `USART_App_Printf(USART_TypeDef *USARTx, const char *format, ...)`

Use this instead of adding local `USARTx_printf()`, `USARTx_SendString()`, or per-file formatted UART wrappers.

Current intent:

- keep formatting and blocking byte-send behavior centralized in `MiniBalance/usart_app.c`
- allow callers to choose the target USART, such as `USART_App_Printf(USART3, "...")`
- keep low-level USART init and byte-send helpers in `MiniBalance_HARDWARE/USART3/`
- avoid scattering `vsprintf` / `vsnprintf` buffers across user logic files

## RPM Input Convention

For the current serial debug/control path, treat `TC:<value>` as an RPM input command instead of a raw encoder-target command.

Current mapping rule:

- `268 rpm -> 35 encoder_target`
- linear conversion: `encoder_target = rpm * 35 / 268`

Current implementation expectations:

- parse `TC:` payload as `rpm`
- map RPM into the PID target domain before writing `Target_EncoderC`
- keep PID internals unchanged when only the user input unit changes
- clamp mapped target into `[-35, 35]` unless future hardware validation updates this range

Maintenance guidance:

- if future user-facing commands are defined in RPM, keep unit conversion at the user-input boundary
- avoid pushing RPM units directly into existing PID code that currently expects encoder-domain targets
- if mapping constants change after calibration, update both code and this document together

## App Layer Refactor Convention

For motor-control refactoring in this project, prefer moving user-layer control logic into `app` modules under `MiniBalance/`, while keeping low-level peripheral initialization and hardware driver details under `MiniBalance_HARDWARE/`.

Current preferred split:

- `MiniBalance/CONTROL/control.c`
  - keeps timer interrupt entry and control-flow scheduling
  - should trend toward becoming a thin dispatcher rather than a large implementation file

- `MiniBalance/pid_app.c`
  - holds user-layer PID execution logic
  - holds PWM limiting logic closely tied to control policy

- `MiniBalance/encoder_app.c`
  - holds user-layer encoder usage logic
  - may wrap hardware read APIs such as `Read_Encoder()` into motor-channel-specific functions

- `MiniBalance/motor_app.c`
  - holds user-layer motor output orchestration
  - may wrap enable / disable / direction / PWM application logic used by the control loop

What should remain in hardware-layer modules:

- timer encoder-mode initialization such as `Encoder_Init_TIMx()`
- raw encoder counter access helpers such as `Read_Encoder()`
- PWM peripheral setup such as `MiniBalance_PWM_Init()`
- motor GPIO initialization such as `Motor_Init()`

What should move toward app-layer modules:

- motor-channel-specific control orchestration
- PID execution policy
- control-loop-facing encoder read wrappers
- control-loop-facing motor output wrappers

Maintenance guidance:

- when refactoring, do not change behavior and architecture at the same time unless necessary
- prefer first extracting logic into app-layer wrappers, then simplifying `control.c`
- `control.c` should increasingly read like a control-loop schedule, not a bag of hardware operations
