#include "pid_app.h"

extern int Motor_PwmC;
extern int Motor_PwmD;

static int MotorC_Bias;
static int MotorC_Pwm;
static int MotorC_LastBias;
static int MotorD_Bias;
static int MotorD_Pwm;
static int MotorD_LastBias;

void PID_App_LimitPwm(int amplitude_motor)
{
	if(Motor_PwmC<-amplitude_motor) Motor_PwmC=-amplitude_motor;
	if(Motor_PwmC>amplitude_motor)  Motor_PwmC=amplitude_motor;
	if(Motor_PwmD<-amplitude_motor) Motor_PwmD=-amplitude_motor;
	if(Motor_PwmD>amplitude_motor)  Motor_PwmD=amplitude_motor;
}

void PID_App_Reset(void)
{
	MotorC_Bias = 0;
	MotorC_Pwm = 0;
	MotorC_LastBias = 0;
	MotorD_Bias = 0;
	MotorD_Pwm = 0;
	MotorD_LastBias = 0;
	Motor_PwmC = 0;
	Motor_PwmD = 0;
}

int PID_App_RunMotorC(int encoder, int target)
{
	MotorC_Bias=target-encoder;
	MotorC_Pwm+=100*(MotorC_Bias-MotorC_LastBias)+10*MotorC_Bias;
	MotorC_LastBias=MotorC_Bias;
	if(MotorC_Pwm>7199) MotorC_Pwm = 7199;
	if(MotorC_Pwm<-7199) MotorC_Pwm = -7199;
	return MotorC_Pwm;
}

int PID_App_RunMotorD(int encoder, int target)
{
	MotorD_Bias=target-encoder;
	MotorD_Pwm+=100*(MotorD_Bias-MotorD_LastBias)+10*MotorD_Bias;
	MotorD_LastBias=MotorD_Bias;
	if(MotorD_Pwm>7199) MotorD_Pwm = 7199;
	if(MotorD_Pwm<-7199) MotorD_Pwm = -7199;
	return MotorD_Pwm;
}
