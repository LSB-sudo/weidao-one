#include "encoder_app.h"

void Encoder_App_Init(void)
{
	Encoder_Init_TIM3();
	Encoder_Init_TIM4();
}

int Encoder_App_ReadMotorC(void)
{
	return Read_Encoder(3);
}

int Encoder_App_ReadMotorD(void)
{
	return Read_Encoder(4);
}
