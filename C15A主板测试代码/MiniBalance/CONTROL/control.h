#ifndef __CONTROL_H
#define __CONTROL_H
#include "mydefine.h"
extern float Accel_Y,Accel_Z,Accel_X,Gyrox,Gyroy,Gyroz;
#define PI 3.14159265
#define ZHONGZHI 0
#define DIFFERENCE 100
int EXTI15_10_IRQHandler(void);
void Kinematic_Analysis(float Vx,float Vy,float Vz);
void Encoder_Analysis(float Va,float Vb,float Vc);
void Key(void);
float Balance_Control(float Angle,float Gyro);
float Position_Control(int encoder);
u32 myabs(long int a);
float target_limit_float(float insert,float low,float high);
void Get_Angle(void);
void Control_TaskDelayTick(void);
#endif
