# 蓝牙 USART3 BLE 转速控制协议规范

## 1. 目标

本文定义 APP 通过 BLE 蓝牙模块向 STM32 `USART3` 发送左右电机转速指令的数据传输范式。

协议目标：

- 通信链路简单，便于手机 APP、蓝牙调试助手和固件端快速联调
- 明确左右电机转速单位、方向和限幅规则
- 保持 `USART3` 硬件收发层与上层业务解析解耦

## 2. 通信链路

```text
手机 APP
  -> BLE 蓝牙模块
  -> UART TTL 串口
  -> STM32 USART3
  -> 固件命令解析层
  -> 左右电机目标转速
```

## 3. 串口参数

| 参数 | 取值 |
| --- | --- |
| 串口 | USART3 |
| 波特率 | 9600 bps |
| 数据位 | 8 |
| 停止位 | 1 |
| 校验位 | None |
| 流控 | None |
| 编码 | ASCII / UTF-8 兼容 ASCII 子集 |
| 换行 | `\n`，兼容 `\r\n` |

说明：

- BLE 模块对 APP 暴露 BLE 通道，对 STM32 暴露 UART 串口。
- 固件侧只关心 `USART3` 接收到的字节流，不直接处理 BLE GATT 细节。
- APP 每发送一条控制命令，建议以换行符结束，便于固件按行解析。

## 4. 速度单位与方向约定

| 字段 | 含义 | 单位 | 范围 |
| --- | --- | --- | --- |
| `left` | 左电机目标转速 | rpm | `-268` 到 `268` |
| `right` | 右电机目标转速 | rpm | `-268` 到 `268` |

方向约定：

- 正数：前转
- 负数：后转
- `0`：停止或目标转速为零

限幅规则：

- APP 发送值必须在 `[-268, 268]` 范围内。
- 固件解析后仍必须做二次限幅，防止异常输入。
- 超出范围的值建议按错误帧处理；若需要容错，也可以夹紧到边界值。

推荐固件默认策略：

```text
left  < -268 或 left  > 268  -> 拒绝本帧
right < -268 或 right > 268  -> 拒绝本帧
```

## 5. 推荐文本协议

### 5.1 命令格式

```text
RPM,L:<left>,R:<right>\n
```

字段说明：

| 字段 | 说明 |
| --- | --- |
| `RPM` | 命令类型，表示转速控制命令 |
| `L:` | 左电机转速字段 |
| `R:` | 右电机转速字段 |
| `,` | 字段分隔符 |
| `\n` | 命令结束符 |

### 5.2 示例

左右电机都以前进方向 100 rpm 运行：

```text
RPM,L:100,R:100\n
```

左电机前转 120 rpm，右电机后转 120 rpm：

```text
RPM,L:120,R:-120\n
```

左右电机停止：

```text
RPM,L:0,R:0\n
```

固件把 `RPM,L:0,R:0\n` 作为 APP 停车命令处理：

- 左右电机目标值清零。
- 电机输出进入停止状态。
- PID 内部累积状态清零，避免下次启动受到旧 PWM 累积量影响。
- 后续任意合法非零 RPM 命令会重新使能电机输出。

左电机最大后转，右电机最大前转：

```text
RPM,L:-268,R:268\n
```

## 6. 解析规则

固件收到一行数据后按以下顺序处理：

1. 去除行尾 `\r`、`\n`。
2. 判断是否以 `RPM,` 开头。
3. 查找 `L:` 和 `R:` 字段。
4. 将 `left`、`right` 解析为有符号整数。
5. 检查两个值是否都在 `[-268, 268]` 范围内。
6. 通过校验后，更新左右电机目标转速。
7. 校验失败时，丢弃本帧，不更新目标值。

建议：

- 一帧最大长度建议不超过 32 字节。
- 解析缓冲区建议预留 64 字节，防止异常输入导致溢出。
- 接收超时或缓冲区满仍未收到换行符时，应清空当前帧。

## 7. 与现有电机控制目标的关系

APP 下发的是 rpm，不建议直接把 rpm 写入现有 PID 内部目标。

推荐边界转换：

```text
encoder_target = rpm * 35 / 268
```

转换规则：

- `268 rpm -> 35 encoder target`
- `-268 rpm -> -35 encoder target`
- `0 rpm -> 0 encoder target`
- 转换结果夹紧到 `[-35, 35]`

设计原则：

- APP、BLE、串口协议层使用 rpm。
- 固件命令解析层负责把 rpm 转换到电机控制内部目标域。
- PID 内部仍使用当前工程已有的编码器目标域，避免同时修改协议和控制算法。

## 8. APP 发送周期建议

推荐发送策略：

| 场景 | 建议 |
| --- | --- |
| 摇杆/滑杆持续控制 | 20 ms 到 100 ms 发送一次 |
| 按钮控制 | 状态变化时发送一次 |
| 停止按钮 | 立即发送 `RPM,L:0,R:0\n` |
| APP 断开连接 | 固件侧应进入通信超时保护 |

固件侧建议增加通信超时保护：

```text
如果超过 500 ms 未收到有效 RPM 命令，则目标转速置 0  ，已经在32上面实现了 
```

该保护用于避免 APP 断连或 BLE 链路中断后电机继续保持旧目标运行。

## 9. 应答格式

初期联调可以保留回显能力，便于确认 USART3 收发正常。

推荐 ACK：

```text
OK,L:<left>,R:<right>\n
```

示例：

```text
OK,L:100,R:-100\n
```

推荐错误响应：

```text
ERR,<reason>\n
```

常见错误：

| 错误 | 含义 |
| --- | --- |
| `ERR,FORMAT` | 帧格式错误 |
| `ERR,RANGE` | 转速超出范围 |
| `ERR,VALUE` | 数值解析失败 |
| `ERR,LEN` | 帧长度超限 |

如果固件当前阶段只做接收和回显测试，可以暂不实现 ACK/ERR，但协议实现阶段建议加入。

## 10. 非法输入处理示例

| 输入 | 处理结果 |
| --- | --- |
| `RPM,L:300,R:0\n` | 拒绝，返回 `ERR,RANGE` |
| `RPM,L:-269,R:0\n` | 拒绝，返回 `ERR,RANGE` |
| `RPM,L:abc,R:0\n` | 拒绝，返回 `ERR,VALUE` |
| `RPM,L:100\n` | 拒绝，返回 `ERR,FORMAT` |
| `HELLO\n` | 非 RPM 命令，可忽略或返回 `ERR,FORMAT` |

## 11. 固件分层建议

建议保持以下分层：

```text
MiniBalance_HARDWARE/USART3
  只负责 USART3 初始化、字节接收、字节发送

MiniBalance 或 app 层
  负责 BLE 文本命令组帧、解析、限幅、rpm 到控制目标转换

MiniBalance/CONTROL
  只读取已经转换后的目标值，执行原有 PID 控制流程
```

不建议在 `USART3_IRQHandler()` 中直接写复杂业务逻辑。

推荐结构：

```text
USART3_IRQHandler()
  -> 收集字节到行缓冲
  -> 收到换行后通知 app 层解析
  -> app 层解析 RPM 命令
  -> 转换为 Target_EncoderC / Target_EncoderD
```

## 12. 版本记录

| 版本 | 日期 | 说明 |
| --- | --- | --- |
| v1.0 | 2026-05-28 | 定义 BLE USART3 rpm 文本控制协议 |

## 13. 联调确认示例

以下命令是合法的 BLE USART3 转速控制命令：

```text
RPM,L:200,R:150\n
```

含义：

- 左电机目标转速为 `200 rpm`，前转。
- 右电机目标转速为 `150 rpm`，前转。

固件侧目标映射结果：

```text
Target_EncoderC = 200 * 35 / 268 = 26
Target_EncoderD = 150 * 35 / 268 = 20
```

发送注意事项：

- APP 或串口助手必须发送真正的换行符 `\n`，而不是发送反斜杠字符和字母 `n` 两个普通字符。
- 如果工具支持 `CRLF`，发送 `\r\n` 也可以。
- 命令前后不要带空格。
- `R:` 后面的数字后面应直接结束或紧跟换行符。

## 14. Firmware battery voltage report

The firmware can send a battery-voltage status line on `USART3` every 100 ms.

Battery reporting is disabled by default so the BLE command channel is not occupied immediately after power-on.

Enable report:

```text
BAT:1\n
```

Disable report:

```text
BAT:0\n
```

Frame format:

```text
Battery:<voltage>V\r\n
```

Example:

```text
Battery:12.34V\r\n
```

Notes:

- The report is generated from the existing battery ADC sampling path.
- The 100 ms cadence is derived from the TIM6 control tick and sent from the foreground `while(1)` loop, not from inside the interrupt handler.
- The report is only sent while `BAT:1` mode is enabled.
- APP/BLE receivers should treat this as an asynchronous status message that may appear between command ACK/ERR lines.
