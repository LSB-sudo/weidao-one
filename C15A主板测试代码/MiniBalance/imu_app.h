#ifndef __IMU_APP_H
#define __IMU_APP_H

#include "mydefine.h"

void IMU_App_Init(void);
void IMU_App_CalibrateZero(u16 sample_count, u16 delay_ms_per_sample);
void IMU_App_Update(void);
void IMU_App_GetAttitude(ATTITUDE_DATA_t *attitude);
void IMU_App_GetGyro(IMU_DATA_t *imu_data);

#endif
