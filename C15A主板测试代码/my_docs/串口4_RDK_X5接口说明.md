# 串口4 RDK X5 接口说明

## 1. 接口用途

`UART4` 用于 STM32 主控与 RDK X5 之间的串口通信。

当前定位：

- STM32 侧接口：`UART4`
- 对端设备：RDK X5
- 通信类型：TTL UART 串口通信
- 当前驱动路径：`MiniBalance_HARDWARE/USART4/`

## 2. 串口参数

| 参数 | 取值 |
| --- | --- |
| 串口 | UART4 |
| 波特率 | 115200 bps |
| 数据位 | 8 |
| 停止位 | 1 |
| 校验位 | None |
| 流控 | None |
| 通信格式 | 8N1 |

说明：

- `MiniBalance_HARDWARE/USART4/usart4.c` 中的 `uart4_init(u32 bound)` 使用传入的 `bound` 作为实际波特率。
- RDK X5 联调时，STM32 初始化应使用 `uart4_init(115200)`，保证双方波特率一致。

## 3. MCU 侧引脚

| 信号 | STM32 引脚 | 方向 | 说明 |
| --- | --- | --- | --- |
| UART4_TX | PC10 | STM32 -> RDK X5 | STM32 发送，连接 RDK X5 的 RX |
| UART4_RX | PC11 | RDK X5 -> STM32 | STM32 接收，连接 RDK X5 的 TX |
| GND | GND | 共地 | STM32 与 RDK X5 必须共地 |

接线时注意：

- `PC10 / UART4_TX` 接 RDK X5 串口 `RX`
- `PC11 / UART4_RX` 接 RDK X5 串口 `TX`
- STM32 与 RDK X5 必须连接公共地线
- 确认双方 UART 电平兼容，避免 5V 电平直接接入不耐受的 3.3V 串口

## 4. 当前固件实现状态

相关文件：

- `MiniBalance_HARDWARE/USART4/usart4.c`
- `MiniBalance_HARDWARE/USART4/usart4.h`
- `MiniBalance/rdk_app.c`
- `MiniBalance/rdk_app.h`
- `MiniBalance/ultrasonic_app.c`
- `MiniBalance/ultrasonic_app.h`

当前驱动已提供：

- `uart4_init(u32 bound)`：UART4 初始化
- `UART4_SendByte(u8 data)`：发送单字节
- `UART4_SendBuffer(const u8 *buffer, u16 length)`：发送缓冲区
- `UART4_IRQHandler(void)`：UART4 接收中断，只负责把接收字节转交给 app 层

当前 RDK X5 侧命令由 `MiniBalance/rdk_app.c` 解析。RDK X5 检测到鸟类后发送固定字符串：

```text
BIRD\r\n
```

STM32 接收到完整 `BIRD` 命令后，会打开 PA8 的超声波 PWM 输出，持续 5 秒后自动关闭。5 秒驱鸟期间重复收到 `BIRD` 不会延长当前触发窗口；关闭后再次收到 `BIRD\r\n` 才会再次触发。

## 5. 初始化要求

若启用 RDK X5 通信，应在系统初始化阶段调用：

```c
uart4_init(115200);
```

初始化后，RDK X5 侧串口参数也应配置为：

```text
115200, 8N1, no flow control
```

超声波输出由 app 层初始化：

```c
Ultrasonic_App_Init();
RDK_App_Init();
```

主循环中需要周期调用：

```c
RDK_App_Task();
```

## 6. 联调检查项

1. 确认 STM32 `PC10` 与 RDK X5 `RX` 交叉连接。
2. 确认 STM32 `PC11` 与 RDK X5 `TX` 交叉连接。
3. 确认 STM32 与 RDK X5 共地。
4. 确认双方波特率均为 `115200 bps`。
5. RDK X5 发送 `BIRD\r\n`，确认 PA8 超声波输出开启。
6. 确认 PA8 输出持续约 5 秒后自动关闭。
7. 确认关闭后再次发送 `BIRD\r\n` 可以再次触发。

## 7. 风险与备注

- 本文确认的是 STM32 MCU 侧软件引脚配置和联调串口参数。
- `PC10/PC11` 在板级硬件上具体接到哪个接口座，需要结合原理图或板子丝印确认。
- 当前 UART4 驱动只负责收发和中断入口，RDK X5 业务协议解析放在 `MiniBalance/rdk_app.c`。
- PA8 超声波 PWM 控制放在 `MiniBalance/ultrasonic_app.c`，避免把硬件控制细节直接留在主函数中。

## 8. 联调记录

### 2026-05-29

- 串口4基础通信已确认可用。
- RDK X5 侧发送实际 `BIRD\r\n` 帧后，STM32 能通过 UART4 接收并触发 PA8 超声波输出。
- PA8 超声波输出会持续约 5 秒后自动关闭。
- 关闭后再次发送 `BIRD\r\n` 可以再次触发驱鸟动作。
- 当前行为符合约定：5 秒触发窗口内重复收到 `BIRD` 不延长本次输出，必须等本次关闭后再次收到 `BIRD\r\n` 才会重新触发。
