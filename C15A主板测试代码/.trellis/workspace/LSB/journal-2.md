# Journal - LSB (Part 2)

> AI development session journal
> Started: 2026-05-26

---

## 2026-05-27 - 用户侧上板验证成功补记

### 验证来源

用户反馈：当前修改已经验证成功。

### 已确认状态

- PA8 `TIM1_CH1` 约 `20 kHz` PWM 输出方案已通过用户侧验证
- 当前调试/初始化链路在用户测试场景下可用
- 前述 journal 中标记的“未在会话内执行完整编译、下载或上板验证”状态，更新为：已由用户侧完成上板验证

### 仍需保留的边界

- 本会话内仍未直接执行本地编译或硬件测量
- 超声波设置在源码侧仍未发现明确实现；若后续接入，需要单独记录 Trig/Echo 引脚、定时器资源和中断策略


## 2026-05-27 - STM32F103RCT6 datasheet skill 生成与超声波设置记录

### 背景

用户要求使用本地 `E:\AI_Projects\Skill_Seekers`，将项目参考资料
`reference_docs/STM32F103RCT6.PDF` 提取成一个可复用的 skill，并随后要求将本次工作写入 journal，同时补充记录超声波设置相关状态。

### 本次执行内容

1. 生成 STM32F103RCT6 datasheet skill
   - 使用 `Skill_Seekers` 对 `reference_docs/STM32F103RCT6.PDF` 进行 PDF 提取
   - 首次提取时遇到 Windows `gbk` 控制台编码无法打印 CLI emoji 的问题
   - 重新以 UTF-8 运行后完成提取
   - PDF 共提取 `144` 页
   - 最终整理为项目内 skill：
     - `.agents/skills/stm32f103rct6-datasheet/SKILL.md`
     - `.agents/skills/stm32f103rct6-datasheet/references/index.md`
     - `.agents/skills/stm32f103rct6-datasheet/references/STM32F103RCT6.md`
     - `.agents/skills/stm32f103rct6-datasheet/assets/images/*`
   - 同步生成 OpenAI target 打包产物：
     - `.agents/skills/stm32f103rct6-datasheet-openai.zip`

2. 校验 skill 产物
   - 检查 `SKILL.md` frontmatter，确认包含 `name` 与 `description`
   - 运行 `Skill_Seekers` quality checker
   - 质量结果为 `95/100`，评级 `Grade A`
   - 运行 `skill-seekers-package --target openai` 并成功生成 zip 包
   - `quick_validate` 在当前 `Skill_Seekers` 仓库中未发现，因此未使用该检查方式

3. 更新项目协作规则
   - 已在 `AGENTS.md` 中新增项目决策：
     - 以后本项目的嵌入式固件开发应先使用 `embedded-dev` skill
     - 覆盖 STM32 外设、GPIO、定时器、PWM、中断、ADC、UART、IIC/I2C、SPI、DMA、驱动移植、引脚规划、datasheet、板级资源和硬件调试等场景

4. PA8 PWM 设置记录
   - 当前 `USER/MiniBalance.c` 中新增 `PA8_PWM_20K_Init()`
   - PA8 配置为 `TIM1_CH1` 复用推挽输出
   - `TIM1` 时基配置为：
     - `Prescaler = 0`
     - `Period = 3599`
     - 在 72 MHz 定时器时钟假设下输出约 `20 kHz`
   - `TIM_OC1` 配置为 PWM1
   - `Pulse = 1800`，占空比约 50%
   - 已在 `main()` 初始化流程中调用，位置在 `MotorControlProbe_InitGpioCOutputs(GPIO_Pin_0)` 之后、`TIMING_TIM_Init(7199,49)` 之前

5. 超声波设置记录
   - 本轮按用户要求将“超声波设置”纳入 journal 记录
   - 已在源码目录 `USER/`、`MiniBalance/`、`MiniBalance_HARDWARE/`、`SYSTEM/` 中检索以下关键词：
     - `超声`
     - `ultra`
     - `sonic`
     - `HC-SR`
     - `HCSR`
     - `TRIG`
     - `ECHO`
   - 当前未发现明确的超声波模块、HC-SR04 驱动、Trig/Echo 专用 GPIO 初始化或测距逻辑
   - 仅发现通用触发字样或无关库注释，例如：
     - `MiniBalance_HARDWARE/EXTI/exti.c` 中存在 `EXTI_Trigger_Falling`
     - `MiniBalance_HARDWARE/ADC/adc.c` 中存在 ADC 外部触发相关配置注释
     - STM32 标准库和 CMSIS 头文件中存在大量通用 `Trigger` 定义
   - 因此当前可记录结论为：本轮没有找到可确认的超声波源码配置；若后续要接入超声波，应单独建立引脚分配、定时捕获或 EXTI 测距方案，并记录 Trig/Echo 引脚、定时器资源与中断优先级

### 关键结论

- STM32F103RCT6 datasheet 已被提取为项目内可复用 skill，后续查询芯片引脚、复用功能、定时器、ADC、GPIO、USART、SPI、I2C、DMA、电气参数时可优先使用该 skill
- `Skill_Seekers` 提取的 markdown 存在 PDF 版式退化和编码噪声，适合检索参考，不应直接作为无校验的代码依据
- PA8 当前已作为 `TIM1_CH1` 输出约 `20 kHz` PWM，约 50% 占空比
- 当前工程源码中未发现明确超声波设置实现，超声波部分仍属于待规划/待确认项

### 当前状态

本轮主要是资料 skill 生成、协作规则补充和 journal 记录，没有新增固件功能代码。

仍待确认：

1. 是否需要把 `references/STM32F103RCT6.md` 继续拆分为 `pinout.md`、`timers.md`、`adc.md`、`gpio_af.md` 等主题文件
2. 是否需要将 datasheet skill 安装为全局 Codex skill，而不是仅保留在项目内 `.agents/skills/`
3. 超声波模块是否已经有外部硬件引脚规划；若有，应补充 Trig/Echo 引脚、使用定时器、测距周期和中断策略


## 2026-05-27 - TIM6 调度迁移、控制目标限幅收紧与调试打印切换

### 背景

用户连续提出了三项围绕当前电机控制与调试链路的调整需求：

- 将当前控制中断调度从 `TIM1` 切换为 `TIM6`
- 将当前电机控制目标域的最大绝对值收紧到 `17`
- 注释掉当前串口目标回显打印，改为输出电池电压和 IMU 状态

这三项修改都要求尽量不改 PID 内核行为，而是优先在调度层、输入边界和调试输出边界完成。

### 本次执行内容

1. 将控制调度定时器从 `TIM1` 迁移到 `TIM6`
   - 修改 `MiniBalance_HARDWARE/ENCODER/encoder.c` 中 `TIMING_TIM_Init()`
   - 将中断源从 `TIM1_UP_IRQn` 切换为 `TIM6_IRQn`
   - 将时钟使能从 `RCC_APB2Periph_TIM1` 切换为 `RCC_APB1Periph_TIM6`
   - 将定时器实例从 `TIM1` 切换为 `TIM6`
   - 保持 `arr=7199`、`psc=49` 不变，维持原 5 ms 控制节拍

2. 将控制中断服务入口从 `TIM1_UP_IRQHandler()` 切换为 `TIM6_IRQHandler()`
   - 修改 `MiniBalance/CONTROL/control.c`
   - 将中断状态检查与清除改为 `TIM6`
   - 保留中断内原有逻辑顺序不变：
     - `Get_Angle()`
     - `Key()`
     - 编码器读取
     - 电池采样更新
     - PID 计算
     - PWM 输出

3. 将电机控制目标域统一收紧到 `±17`
   - 修改 `MiniBalance/usart_app.c`
   - 将 `MAX_RPM_TO_ENCODER_TARGET` 从 `35.0f` 改为 `17.0f`
   - 保留 `268 rpm` 为输入边界基准，只调整映射后的目标域范围
   - 新增 `USART_App_ClampEncoderTarget()`，使协议帧路径与 `TC:/TD:` 调试命令路径都统一限幅到 `[-17, 17]`

4. 切换调试打印内容
   - 注释 `MiniBalance/usart_app.c` 中当前 `TC:/TD:` 的目标回显 `printf`
   - 在 `USER/MiniBalance.c` 新增 `Debug_PrintBatteryAndImu()`
   - 在主循环中增加低频状态打印，输出：
     - `Battery`
     - `Roll`
     - `Pitch`
     - `Yaw`
     - `GyroX`
     - `GyroY`
     - `GyroZ`

5. 修正一次与本轮目标无关的副作用
   - 在打印切换过程中发现 `USER/MiniBalance.c` 里 `MotorControlProbe_InitGpioCOutputs(GPIO_Pin_0)` 调用被意外带掉
   - 已立即恢复，避免破坏前一轮已经验证过的初始化顺序

### 关键结论

- 本轮 `TIM6` 迁移只替换了调度定时器实例，没有改控制中断内的业务逻辑
- 电机控制目标域现在从输入边界统一限制为 `±17`，不会再出现一条输入路径被收紧、另一条路径仍可直接写入更大目标值的情况
- 当前活跃调试打印已经从“串口命令目标回显”切换为“电池电压 + IMU 状态输出”
- 主循环中的状态打印目前按空转计数触发，不是严格毫秒定时；如果后续串口输出频率需要更稳定，应改为基于 `TIM6` 节拍或明确的软件定时

### 本次产出

- 完成 `TIM1 -> TIM6` 的控制调度迁移
- 完成控制目标域统一限幅到 `±17`
- 完成调试输出内容切换到电池与 IMU 状态
- 保留并恢复了 `PC0` 探针初始化调用，避免影响既有启动行为

### 当前状态

当前代码层面的改动已经完成，但本轮未在会话内执行完整编译、下载或上板验证。

当前仍待用户侧确认的重点包括：

1. `TIM6` 中断是否按预期进入，控制节拍是否与迁移前一致
2. `TC:268` 是否已映射到目标值 `17`
3. 协议帧下发超范围目标时，是否也被正确限幅到 `±17`
4. 新状态打印的输出频率是否合适，是否需要改成更稳定的定时输出


## 2026-05-26 - 删除 show / DataScope_DP 调试显示链路

### 背景

用户计划清理当前工程中不再需要的显示与上位机调试输出链路，重点目标为：

- `MiniBalance/show/show.c`
- `DataScope_DP.c`

在执行删除前，先对这两个模块的接口定义、头文件暴露位置、工程编译登记位置、以及业务调用分布做了定位确认，避免误删控制主链路。

### 本次执行内容

1. 定位 `show` 模块接口分布
   - 确认 `show.h` 中暴露接口：
     - `oled_show()`
     - `APP_Show()`
     - `DataScope()`
     - `OLED_Show_CCD()`
     - `OLED_DrawPoint_Shu()`
   - 确认 `show.c` 中对应实现位置
   - 确认当前业务代码中几乎没有有效调用，仅看到 `USER/MiniBalance.c` 中一处已注释掉的 `oled_show()`

2. 定位 `DataScope_DP` 模块接口分布
   - 确认工程实际参与编译的是：
     - `MiniBalance_HARDWARE/DataScope_DP/DataScope_DP.C`
   - 确认头文件暴露接口：
     - `DataScope_OutPut_Buffer`
     - `DataScope_Get_Channel_Data()`
     - `DataScope_Data_Generate()`
   - 确认当前源码中未发现业务层对这些接口的直接调用

3. 确认耦合点与删除边界
   - `SYSTEM/sys/sys.h` 包含：
     - `show.h`
     - `DataScope_DP.h`
   - `USER/MiniBalance.uvprojx` 中登记了：
     - `show.c`
     - `MiniBalance_HARDWARE/DataScope_DP/DataScope_DP.C`

4. 用户已手动完成删除
   - 本轮未由 AI 直接执行代码删除
   - 仅完成删除前的接口分布梳理和删除边界说明

### 关键结论

- `show` 模块当前更接近历史遗留显示/调试封装，主控制链路对其依赖极弱
- `DataScope_DP` 当前虽然仍在工程项中，但业务源码侧未发现直接使用路径
- 真正需要同步清理的重点不只是源文件本身，还包括：
  - `sys.h` 的头文件包含
  - `uvprojx` 的工程文件登记

### 本次产出

- 完成 `show` / `DataScope_DP` 接口分布梳理
- 为后续安全删除提供了清理边界
- 用户随后已手动删除对应代码

### 当前状态

本轮记录的是一次“删除前定位 + 用户手动删除”的整理工作。

当前尚未在本轮中完成以下验证：

1. 删除后是否仍能通过编译
2. `sys.h`、`uvprojx`、相关 include 路径是否已完全同步清理
3. 是否还存在残余的注释引用、文档引用或重复副本需要一并收尾

## 2026-05-26 - OLED_Init 相关电机异常排查与当前保留方案

### 背景

用户反馈一个异常现象：

- 不调用 `OLED_Init()` 时，电机速度控制异常
- 例如串口输入 `130 rpm`，电机表现接近满转
- 调用 `OLED_Init()` 时，速度控制又恢复正常

该问题表现出明显的“初始化副作用”特征，因此本轮目标不是立刻抽象重构，而是先找出最小可工作的初始化条件。

### 本次执行内容

1. 对 `OLED_Init()` 的副作用做拆分
   - 将完整 `OLED_Init()` 拆成若干可独立验证的启动分支
   - 分别测试：
     - 仅延时
     - 仅 `PC0`
     - 仅 `PC13/14/15`
     - `PC0 + PC13/14/15`
     - 完整 `OLED_Init()`

2. 增加串口侧调试信息
   - 在串口命令解析时打印：
     - 目标编码器量
     - 对应输入 rpm
     - 当时的 `Current_EncoderC/D`

3. 反复调整初始化顺序
   - 验证 `MotorControlProbe_InitGpioCOutputs(GPIO_Pin_0)` 放在不同初始化阶段的效果
   - 最终确认：
     - 不是“任意时刻拉高 `PC0` 都有效”
     - 只有放在 `Adc_Init()` 之后调用，当前代码才能稳定恢复正常控制

### 实验结论

当前可确认的事实：

1. 问题不是单纯“必须调用 OLED 显示”
   - 真正影响结果的是 `OLED_Init()` 带来的某种初始化副作用

2. `PC0` 确实参与了当前可工作的最小条件
   - 但这并不等价于“PC0 与电机驱动直接相关”
   - 用户已检查并确认：`PC0` 与电机驱动本身没有直接关系

3. 初始化时序明显影响结果
   - 仅做同样的 `PC0` 输出配置，如果放置位置过早，不能稳定复现正常结果
   - 当前实测只有在 `Adc_Init()` 之后调用 `MotorControlProbe_InitGpioCOutputs(GPIO_Pin_0)` 才能确保控制正常

4. 根因目前仍未完全查清
   - 从代码上看，`Adc_Init()` 只初始化 `GPIOA` 与 `ADC2`
   - 它本身并不直接操作 `GPIOC` 或 `PC0`
   - 因此当前更像是一个“启动末段时序相关”的板级问题，而非纯代码逻辑问题

### 当前保留方案

本轮决定不继续深挖“玄学根因”，先保留当前能稳定工作的实现。

当前 `main()` 中采用的方案为：

- 先执行：
  - `Encoder_App_Init()`
  - `Adc_Init()`
- 再执行：
  - `Motor_App_Init()`
  - `MotorControlProbe_InitGpioCOutputs(GPIO_Pin_0)`
- 最后启动：
  - `TIMING_TIM_Init(7199,49)`

### 本次产出

- 增加了用于排查初始化副作用的最小探针函数
- 增加了串口侧目标值/反馈值调试打印
- 确认了当前可工作的初始化顺序
- 将“继续追根因”暂时降级，先以稳定可用为主

### 当前状态

当前保留的是“经验上稳定可用”的版本，而不是已完成原理级解释的版本。

后续如果再继续追查，优先方向应为：

1. 板级原理图 / 引脚复用 / 共享使能网络
2. `PC0` 所在网络与编码器反馈链、供电链或控制链的间接耦合
3. 为什么同样的 `PC0` 配置必须在 `Adc_Init()` 之后才稳定生效
## 2026-05-28 BLE USART3 RPM command note

- Confirmed command `RPM,L:200,R:150\n` is valid for the BLE USART3 RPM protocol.
- Meaning: left motor target is `200 rpm` forward, right motor target is `150 rpm` forward.
- Current firmware mapping is `encoder_target = rpm * 35 / 268`.
- Resulting motor-control targets are `Target_EncoderC = 26` and `Target_EncoderD = 20`.
- Practical send note: the app/serial tool must send a real newline byte for `\n`; do not send literal backslash + `n`. `\r\n` is also accepted. Avoid leading/trailing spaces.
- Protocol doc updated: `my_docs/ble_usart3_rpm_protocol.md`.

## 2026-05-28 BLE zero-speed restart fix

- Observed issue: after sending `RPM,L:0,R:0\n`, later non-zero BLE RPM commands did not make the motors run again.
- Root cause in control state: motor output is gated by `Flag_Stop`; BLE commands were only updating `Target_EncoderC/D` and did not own the enable/stop state. Incremental PID state also had no explicit reset path.
- Fix: BLE `RPM,L:0,R:0\n` now sets `Flag_Stop = 1`, clears PWM output, and calls `PID_App_Reset()`. Any valid non-zero BLE RPM command sets `Flag_Stop = 0`.
- Also moved PID internal static state into file-scope variables so it can be reset safely.
- Verification: Keil `UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.

### Follow-up confirmation

- User confirmed the zero-speed restart issue is resolved.
- Current accepted behavior: after `RPM,L:0,R:0\n`, sending another valid non-zero command can start the motors again.
- Keep this as the expected BLE control contract for future USART3 command changes.

## 2026-05-28 USART3 100ms battery voltage report

- Requirement: send battery voltage on USART3 every 100 ms and update the BLE protocol document.
- Implementation:
  - `Battery_App_UpdateSample()` now averages 10 existing battery samples instead of 100, matching the TIM6 path where the battery sample update is called every other 5 ms control tick.
  - `Battery_App_TakeVoltageUpdated()` exposes a one-shot foreground flag so USART3 transmission is not performed inside `TIM6_IRQHandler()`.
  - `USER/MiniBalance.c` now sends `Battery:<voltage>V\r\n` through `USART3_SendBuffer()` whenever the battery app reports a fresh 100 ms voltage update.
  - `my_docs/ble_usart3_rpm_protocol.md` documents the asynchronous firmware battery voltage report.
- Verification: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.

- Hardware observation pending: if hardware is connected, confirm USART3/BLE receives one `Battery:x.xxV` line every 100 ms.

## 2026-05-28 USART3 battery report board hang follow-up

- User reported that after flashing the 100 ms USART3 battery report build, the board became unusable/stuck.
- Immediate risk assessment: continuous unsolicited output on the BLE command channel can occupy or interfere with command interaction, and USART3 transmit is blocking.
- Mitigation implemented: battery reporting is disabled by default and can be toggled over BLE with `BAT:1\n` / `BAT:0\n`.
- The foreground send path remains unchanged when enabled; no USART3 transmit is performed inside `TIM6_IRQHandler()`.
- Verification: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.
- Hardware observation pending: first board test should confirm normal RPM command control after power-on before enabling `BAT:1`.

## 2026-05-28 User working version for USART3 battery report

- User reported that the current board-side code is able to realize the intended function.
- Current working implementation observed in the worktree:
  - `MiniBalance/CONTROL/control.c` defines `unsigned long int uwtick`.
  - `TIM6_IRQHandler()` increments `uwtick` on each TIM6 interrupt.
  - `USER/MiniBalance.c` declares `extern unsigned long int uwtick`.
  - The foreground `while(1)` checks `uwtick >= 20`, sends `Battery Voltage:%d.%02dV\r\n` through `USART3_printf()`, then clears `uwtick = 0`.
- Timing basis:
  - Current `TIMING_TIM_Init(7199,49)` gives a 5 ms TIM6 interrupt period under the existing 72 MHz timer clock assumption.
  - `20 * 5 ms = 100 ms`, so the current working report interval is 100 ms.
- Practical conclusion:
  - This user-tested version keeps the USART3 battery voltage report driven by a simple TIM6 tick counter and foreground transmission path.
  - Preserve this behavior as the known-working baseline before any further cleanup or refactor.

## 2026-05-29 USART app printf API consolidation

- User requested that the STM32 StdPeriph-style formatted UART print helper be moved into `MiniBalance/usart_app.c`.
- Implemented `USART_App_Printf(USART_TypeDef *USARTx, const char *format, ...)` in `MiniBalance/usart_app.c` and exported it from `MiniBalance/usart_app.h`.
- The helper uses `vsnprintf()` with a bounded local buffer and returns `-1` on invalid input or truncation.
- `USER/MiniBalance.c` now calls `USART_App_Printf(USART3, "Battery Voltage:%d.%02dV\r\n", ...)` instead of keeping local `USART3_SendString()` / `USART3_printf()` helper functions in `main`.
- Verification: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.
- Long-term decision requested by user: future formatted UART/serial debug output should prefer `USART_App_Printf()` instead of ad hoc per-file printf wrappers.

## 2026-05-29 UART4 RDK X5 bird deterrent confirmation

- UART4 is assigned to RDK X5 communication at `115200, 8N1`.
- MCU-side pins remain `PC10 = UART4_TX` and `PC11 = UART4_RX`.
- RDK X5 bird-detection command is the fixed line frame `BIRD\r\n`.
- Implementation split:
  - `MiniBalance_HARDWARE/USART4/usart4.c` keeps UART4 initialization, byte send, buffer send, and the IRQ entry.
  - `MiniBalance/rdk_app.c` parses UART4 line commands from RDK X5.
  - `MiniBalance/ultrasonic_app.c` owns PA8/TIM1 ultrasonic PWM output control.
- Behavior confirmed by user:
  - Sending `BIRD\r\n` from RDK X5/serial side triggers PA8 ultrasonic output.
  - Output remains active for about 5 seconds and then turns off automatically.
  - While active, repeated `BIRD` frames do not extend the current 5 second window.
  - After the output turns off, another `BIRD\r\n` frame triggers a new 5 second deterrent window.
- Verification already run for implementation commit: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.
- Protocol doc updated: `my_docs/串口4_RDK_X5接口说明.md`.

## 2026-05-29 TIM6 non-PID scheduler migration

- Requirement: keep `TIM6_IRQHandler()` limited to `uwTick += 5` and motor PID/output work; move other periodic work into scheduler task functions.
- Implementation:
  - `MiniBalance/CONTROL/control.c` now leaves TIM6 with tick increment plus motor PID/output path only.
  - Delay-flag maintenance and encoder/status/battery sampling were extracted to `Control_TaskDelayTick()` and `Control_TaskReadEncoderAndStatus()`.
  - `MiniBalance/schedeler.c` registers RDK timeout handling, delay maintenance, IMU attitude update, key scan, encoder/status sampling, and BLE battery report as scheduler tasks.
  - `MiniBalance/rdk_app.c` now uses `Scheduler_GetTickMs()` for its 5 second ultrasonic window instead of an ISR-local RDK tick counter.
  - `Scheduler_BatteryReportTask()` respects `USART_App_IsBleBatteryReportEnabled()` before blocking USART3 output.
- Verification: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.
- Hardware observation pending: confirm foreground scheduler latency is acceptable for IMU/key/encoder sampling and that motor PID remains stable with encoder values updated outside TIM6.

### 2026-05-29 correction: keep encoder reads in TIM6

- User corrected the scheduler migration boundary: encoder reads are part of the motor PID input path and must stay in `TIM6_IRQHandler()`.
- Fix:
  - `Flag_Target` toggling, `Encoder_App_ReadMotorD()`, `Encoder_App_ReadMotorC()`, and LED status update were restored to TIM6 before PID execution.
  - The scheduler no longer registers `Control_TaskReadEncoderAndStatus()`.
  - `Scheduler_BatteryReportTask()` now calls `Battery_App_UpdateSample()` itself, then conditionally reports over USART3 when BLE battery reporting is enabled.
- Verification: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.

## 2026-05-29 BLE RPM command timeout protection

- Requirement from `my_docs/ble_usart3_rpm_protocol.md`: if no valid RPM command is received for more than 500 ms, target speed must be set to 0.
- Implementation:
  - `MiniBalance/usart_app.c` records `Ble_LastValidRpmTick` whenever a valid `RPM,L:<left>,R:<right>` command is parsed.
  - `USART_App_BleTimeoutTask()` checks the elapsed time using `Scheduler_GetTickMs()`.
  - On timeout, firmware clears `Target_EncoderC` and `Target_EncoderD`, sets `Flag_Stop = 1`, calls `PID_App_Reset()`, and clears `PWMC/PWMD`.
  - `MiniBalance/schedeler.c` registers `USART_App_BleTimeoutTask` as a 20 ms scheduler task.
  - `MiniBalance/usart_app.h` exports the timeout task declaration.
- Verification: Keil `F:\Software_Download\Keil\UV4\UV4.exe -b USER\MiniBalance.uvprojx -j0` completed with `0 Error(s), 0 Warning(s)`.
- Git: no commit created, per current project policy.
