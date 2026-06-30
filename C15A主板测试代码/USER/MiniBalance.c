/***********************************************
WHEELTEC MiniBalance user entry
***********************************************/
#include "stm32f10x.h"
#include "mydefine.h"

#define OLED_SIDE_EFFECT_TEST_NONE        0
#define OLED_SIDE_EFFECT_TEST_DELAY       1
#define OLED_SIDE_EFFECT_TEST_PC0         2
#define OLED_SIDE_EFFECT_TEST_PC13_15     3
#define OLED_SIDE_EFFECT_TEST_GPIO_ALL    4
#define OLED_SIDE_EFFECT_TEST_FULL_OLED   5

#ifndef OLED_SIDE_EFFECT_TEST_MODE
#define OLED_SIDE_EFFECT_TEST_MODE OLED_SIDE_EFFECT_TEST_FULL_OLED
#endif

u8 Flag_Stop=1,Flag_Show=0,Flag_Zero=0;
u8 delay_50,delay_flag;
u8 PID_Send;
float Pitch,Roll,Yaw;
float Motor_Balance,Velocity=70;
u8 ADDR[1];
extern int Motor_PwmC;
extern int Cnt_test;

static void MotorControlProbe_InitGpioCOutputs(u16 pins)
{
	GPIO_InitTypeDef GPIO_InitStructure;

	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOC, ENABLE);
	GPIO_InitStructure.GPIO_Pin = pins;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_Init(GPIOC, &GPIO_InitStructure);
	GPIO_SetBits(GPIOC, pins);
}

void Battery_app()
{
	printf("Battery Voltage:%d.%02dV\r\n", Voltage/100, Voltage%100);
}

int main(void)
{
    MY_NVIC_PriorityGroupConfig(2);
	delay_init();
	JTAG_Set(JTAG_SWD_DISABLE);
	JTAG_Set(SWD_ENABLE);
	LED_Init();
	IIC_Init();
	while(IMU_HW_Init())
	{
	}
	IMU_App_Init();
	IMU_App_CalibrateZero(200,5);
	uart1_init(115200);
	uart3_init(9600);
	uart4_init(115200);
	Encoder_App_Init();
	Adc_Init();
	Motor_App_Init();
	MotorControlProbe_InitGpioCOutputs(GPIO_Pin_0);
	Ultrasonic_App_Init();
	RDK_App_Init();
	scheduler_init();
	TIMING_TIM_Init(7199,49);
	for(int i=0;i<100;i++)
	while(1)
	{
		 scheduler_run();
	}
}
