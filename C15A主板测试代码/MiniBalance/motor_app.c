#include "motor_app.h"

u32 myabs(long int a);

void Motor_App_Init(void)
{
	MiniBalance_PWM_Init(7199,0);
	Motor_Init();
}

void Motor_App_SetPwm(int motor_c, int motor_d)
{
	if(motor_c<0)
	{
		DIR2=0;
	}
	else
	{
		DIR2=1;
	}

	if(motor_d<0)
	{
		DIR3=0;
	}
	else
	{
		DIR3=1;
	}

	PWMC=myabs(motor_c);
	PWMD=myabs(motor_d);
}

u8 Motor_App_TurnOff(int voltage)
{
	u8 temp;
	voltage = voltage;
	if(Flag_Stop==1)
	{
		Flag_Stop=1;
		temp=1;
		EN2=EN3=0;
		PWMC=0;
		PWMD=0;
	}
	else
	{
		EN2=EN3=1;
		temp=0;
	}
	return temp;
}
