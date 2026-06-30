#ifndef __MOTOR_APP_H
#define __MOTOR_APP_H

#include "mydefine.h"

void Motor_App_Init(void);
void Motor_App_SetPwm(int motor_c, int motor_d);
u8 Motor_App_TurnOff(int voltage);

#endif
