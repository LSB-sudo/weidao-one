#ifndef __MYDEFINE_H
#define __MYDEFINE_H

#include <stdint.h>
#include "sys.h"

#include "bsp_imu.h"
#include "usart4.h"

#include "battery_app.h"
#include "encoder_app.h"
#if !defined(__BSP_IMU_H) || defined(BSP_IMU_TYPES_READY)
#include "imu_app.h"
#endif
#include "motor_app.h"
#include "pid_app.h"
#include "rdk_app.h"
#include "schduler.h"
#include "ultrasonic_app.h"
#include "usart_app.h"

#endif
