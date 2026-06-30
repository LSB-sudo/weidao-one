# Quality Guidelines

> Code quality standards for backend development.

---

## Overview

<!--
Document your project's quality standards here.

Questions to answer:
- What patterns are forbidden?
- What linting rules do you enforce?
- What are your testing requirements?
- What code review standards apply?
-->

(To be filled by the team)

---

## Forbidden Patterns

<!-- Patterns that should never be used and why -->

(To be filled by the team)

---

## Required Patterns

<!-- Patterns that must always be used -->

### TIM6 millisecond tick and foreground scheduler

#### 1. Scope / Trigger
- Use this contract when adding foreground periodic tasks in the STM32 StdPeriph firmware.
- TIM6 remains the 5 ms hardware time base. Application code consumes a millisecond tick instead of raw timer slices.

#### 2. Signatures
```c
extern volatile uint32_t uwTick;
void scheduler_init(void);
void scheduler_run(void);
uint32_t Scheduler_GetTickMs(void);
```

#### 3. Contracts
- `uwTick` is a monotonic millisecond counter produced by `TIM6_IRQHandler()`.
- The 5 ms TIM6 interrupt must update tick time with `uwTick += 5`.
- `scheduler_run()` is called from the foreground `while(1)` loop and must not configure TIM6 or NVIC.
- Scheduler task periods are expressed in milliseconds and should be multiples of 5 ms unless a faster tick source is added.

#### 4. Validation & Error Matrix
- Clearing `uwTick` in foreground code -> breaks monotonic time and scheduler periods.
- Adding blocking work to a scheduler task -> delays all later foreground tasks.
- Moving TIM6 register setup into scheduler code -> violates hardware/application layering.

#### 5. Good/Base/Bad Cases
- Good: `TIM6_IRQHandler()` updates `uwTick += 5`; `scheduler_run()` checks elapsed milliseconds with subtraction.
- Base: a 100 ms battery report task runs from the foreground scheduler.
- Bad: `if (uwTick >= 20) { ...; uwTick = 0; }` mixes 5 ms slice counting with millisecond tick semantics.

#### 6. Tests Required
- Search source files to confirm no foreground code clears `uwTick`.
- Confirm `schedeler.c` is included in the Keil project file before building.
- Build with Keil/ARMCC and verify there are no duplicate tick symbols.

#### 7. Wrong vs Correct
```c
/* Wrong: clears the global time base. */
if(uwTick >= 20)
{
    uwTick = 0;
}

/* Correct: scheduler compares elapsed milliseconds. */
if((uint32_t)(now_time - last_run) >= rate_ms)
{
    last_run = now_time;
    task_func();
}
```

### MiniBalance shared header convention

#### 1. Scope / Trigger
- Use this convention for MiniBalance application-layer headers such as `*_app.h`, `schduler.h`, and `CONTROL/control.h`.

#### 2. Signatures
```c
#include "mydefine.h"
```

#### 3. Contracts
- `mydefine.h` is the shared include entry for MiniBalance application-layer public headers.
- Application-layer headers should include only `mydefine.h` for common project types, STM32 aliases, and app-level cross-module declarations.
- Keep low-level driver internals in their own driver headers; do not move register configuration logic into `mydefine.h`.

#### 4. Validation & Error Matrix
- Directly including `sys.h` from multiple app headers -> repeated dependency churn and harder refactors.
- Including app headers from driver headers before their types are complete -> circular include failures.
- Adding broad behavior or variables into `mydefine.h` -> turns the shared header into hidden logic.

#### 5. Good/Base/Bad Cases
- Good: `battery_app.h`, `encoder_app.h`, and `control.h` include `mydefine.h`.
- Base: `mydefine.h` aggregates stable app-facing headers and guards known hardware-type cycles.
- Bad: each app header independently includes `sys.h`, `bsp_imu.h`, or unrelated driver headers.

#### 6. Tests Required
- Search app-layer headers for direct `sys.h` / `bsp_imu.h` includes after refactors.
- Build with Keil/ARMCC and check for implicit declarations or undefined type warnings.

#### 7. Wrong vs Correct
```c
/* Wrong */
#include "sys.h"

/* Correct */
#include "mydefine.h"
```

---

## Testing Requirements

<!-- What level of testing is expected -->

(To be filled by the team)

---

## Code Review Checklist

<!-- What reviewers should check -->

(To be filled by the team)
