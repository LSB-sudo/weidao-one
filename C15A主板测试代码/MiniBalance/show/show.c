#include "show.h"
#include "battery_app.h"

void oled_show(void)
{
	OLED_ShowString(00,00,"EC:");
	if(Current_EncoderC<0)
	{
		OLED_ShowString(20,00,"-");
		OLED_ShowNumber(35,00,-Current_EncoderC,3,12);
	}
	else
	{
		OLED_ShowString(20,00,"+");
		OLED_ShowNumber(35,00,Current_EncoderC,3,12);
	}

	OLED_ShowString(60,00,"ED:");
	if(Current_EncoderD<0)
	{
		OLED_ShowString(80,00,"-");
		OLED_ShowNumber(95,00,-Current_EncoderD,3,12);
	}
	else
	{
		OLED_ShowString(80,00,"+");
		OLED_ShowNumber(95,00,Current_EncoderD,3,12);
	}

	OLED_ShowString(00,20,"MOTOR");
	if(0==Flag_Stop) OLED_ShowString(50,20,"O N");
	else OLED_ShowString(50,20,"OFF");

	OLED_ShowString(0,50,"P:");
	if(Pitch<0)
	{
		OLED_ShowString(20,50,"-");
		OLED_ShowNumber(30,50,-Pitch,3,12);
	}
	else
	{
		OLED_ShowString(20,50," ");
		OLED_ShowNumber(30,50,Pitch,3,12);
	}
	OLED_ShowString(70,50,"R:");
	if(Roll<0)
	{
		OLED_ShowString(85,50,"-");
		OLED_ShowNumber(95,50,-Roll,5,12);
	}
	else
	{
		OLED_ShowString(85,50," ");
		OLED_ShowNumber(95,50,Roll,5,12);
	}

	OLED_Refresh_Gram();
}

void DataScope(void)
{
	Battery_App_PrintVoltage();
}

void APP_Show(void)
{
}

void OLED_Show_CCD(void)
{
}

void OLED_DrawPoint_Shu(u8 x,u8 y,u8 t)
{
	OLED_DrawPoint(x,y,t);
}
