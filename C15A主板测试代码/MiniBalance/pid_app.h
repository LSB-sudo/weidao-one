#ifndef __PID_APP_H
#define __PID_APP_H

#include "mydefine.h"

void PID_App_LimitPwm(int amplitude_motor);
void PID_App_Reset(void);
int PID_App_RunMotorC(int encoder, int target);
int PID_App_RunMotorD(int encoder, int target);

#endif
