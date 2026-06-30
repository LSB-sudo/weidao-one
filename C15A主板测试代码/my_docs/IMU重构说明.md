# IMU 重构说明

## 本次重构目标

将原先混在 `MiniBalance/ICM20948/bsp_imu.c` 中的两类职责拆开：

- 硬件层职责
- 应用层姿态与策略职责

使其符合当前项目已经形成的分层方向：

- `MiniBalance_HARDWARE / 硬件接口`
- `MiniBalance / app 层语义与控制策略`

## 当前重构结果

### 1. `bsp_imu` 现在承担硬件层职责

文件：

- `MiniBalance/ICM20948/bsp_imu.h`
- `MiniBalance/ICM20948/bsp_imu.c`

当前保留内容：

- ICM20948 / AK09916 寄存器初始化
- 设备上下电接口
- 原始 9 轴读取
- 单位换算
- 原始零偏存取

对外接口：

- `IMU_HW_Init()`
- `IMU_HW_DeInit()`
- `IMU_HW_Read9Axis()`
- `IMU_HW_SetZeroPoint()`
- `IMU_HW_ClearZeroPoint()`

### 2. `imu_app` 现在承担应用层职责

文件：

- `MiniBalance/imu_app.h`
- `MiniBalance/imu_app.c`

当前承接内容：

- 姿态融合
- 上电零偏校准
- 在线 `gyro z` bias 修正
- `Yaw` 的受控磁力计慢速回拉
- 姿态状态管理
- 对外导出姿态与 gyro 数据

对外接口：

- `IMU_App_Init()`
- `IMU_App_CalibrateZero()`
- `IMU_App_Update()`
- `IMU_App_GetAttitude()`
- `IMU_App_GetGyro()`

### 3. 调用入口已切换

`USER/MiniBalance.c`

- 启动流程改为：
  - `IMU_HW_Init()`
  - `IMU_App_Init()`
  - `IMU_App_CalibrateZero(...)`

`MiniBalance/CONTROL/control.c`

- `Get_Angle()` 不再直接调用底层 IMU 接口
- 改为：
  - `IMU_App_Update()`
  - `IMU_App_GetAttitude()`
  - `IMU_App_GetGyro()`

## 当前分层边界

### 属于硬件层的

- 传感器寄存器配置
- 原始采样
- 单位换算
- 原始零偏存取

### 属于 app 层的

- 姿态解算
- 零偏校准流程
- 在线 bias 修正
- `Yaw` 策略
- 后续碰撞特征提取

## 后续建议

### 1. 目录层面可继续规范

本次先完成职责拆分，但文件物理位置仍保留在当前工程可工作的路径中。

后续如果继续整理，建议把底层 IMU 文件进一步迁到更明确的硬件目录，例如：

- `MiniBalance_HARDWARE/IMU/`

### 2. Keil 工程文件需要同步检查

如果 `uvprojx` 中仍按旧文件组织引用，后续需要确认：

- 新文件仍被工程包含
- 旧路径没有残留冲突

### 3. `imu_app` 后续可以继续承接

后续适合继续放进 `imu_app.c` 的内容：

- 碰撞特征提取
- 航向稳定性策略
- 船体运动状态判断的 IMU 侧特征接口

## 结论

本次重构已经完成最关键的一步：

- 底层 IMU 读取和寄存器配置，与姿态融合/策略逻辑分开

这使后续无论是继续优化 `Yaw`，还是加入撞障碍特征提取，都不需要再直接改动底层寄存器驱动。
