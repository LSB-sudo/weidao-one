#ifndef __BSP_IMU_H
#define __BSP_IMU_H

#include "sys.h"

#define REG_VAL_SELECT_BANK_0 0x00
#define REG_VAL_SELECT_BANK_1 0x10
#define REG_VAL_SELECT_BANK_2 0x20
#define REG_VAL_SELECT_BANK_3 0x30

typedef struct{
	float x;
	float y;
	float z;
}PrivateBuf_t;

typedef struct{
	PrivateBuf_t gyro;
	PrivateBuf_t accel;
	PrivateBuf_t magn;
}IMU_DATA_t;

typedef struct{
	float roll;
	float pitch;
	float yaw;
}ATTITUDE_DATA_t;

#define BSP_IMU_TYPES_READY 1

uint8_t IMU_HW_Init(void);
uint8_t IMU_HW_DeInit(void);
void IMU_HW_Read9Axis(IMU_DATA_t *imudata);
void IMU_HW_SetZeroPoint(const IMU_DATA_t *point);
void IMU_HW_ClearZeroPoint(void);

extern IMU_DATA_t axis_9Val;
extern ATTITUDE_DATA_t AttitudeVal;

#endif
