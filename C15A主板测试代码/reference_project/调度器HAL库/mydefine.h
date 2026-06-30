#include "stdio.h"
#include "string.h"
#include "stdarg.h"
#include "stdint.h"
#include "stdlib.h"

#include "main.h"
#include "math.h"
#include "scheduler.h"

//应用层头文件
#include "key_app.h"
#include "usart_app.h"
#include "hwt101_app.h"
#include "motor_app.h"
#include "encoder_app.h"
#include "pid_app.h"
#include "gray_app.h"
#include "pid_tuner.h"

//组件库头文件
#include "ringbuffer.h"
#include "ebtn.h"
#include "pid.h"

//硬件驱动库头文件
#include "motor_driver.h"
#include "hwt101_driver.h"
#include "encoder_driver.h"
//#include "btn_driver.h"
#include "hardware_iic.h"


extern UART_HandleTypeDef huart5;
extern UART_HandleTypeDef huart2;
extern DMA_HandleTypeDef hdma_uart5_rx;
extern HWT101_t hwt101;
extern Motor_t left_motor; 
extern Motor_t right_motor;
extern I2C_HandleTypeDef hi2c2;
extern  float g_line_position_error;
extern PID_T pid_line;        // 循迹环

