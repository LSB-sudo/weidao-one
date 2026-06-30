#include "control.h"
#include "battery_app.h"
#include "encoder_app.h"
#include "imu_app.h"
#include "motor_app.h"
#include "pid_app.h"

#define ATTITUDE_DISPLAY_LPF_ALPHA  0.15f

float Accel_Y,Accel_Z,Accel_X,Gyrox,Gyroy,Gyroz;
u8 Flag_Target,Flag_Change,Flag_Velocity,Flag_All,DIR=0;
float Single;
int j,Balance_Pwm_X,Balance_Pwm_Y,Balance_Pwm_Z,Velocity_Pwm_X,Velocity_Pwm_Y,Velocity_Pwm_Z,Max_Pwm=3500;
u16 time_cnt;
int Target_EncoderC=15,Target_EncoderD=15,Current_EncoderC,Current_EncoderD;
int Motor_PwmC=0,Motor_PwmD=0;

static float LowPassFilter(float current_value, float last_value)
{
	return last_value + ATTITUDE_DISPLAY_LPF_ALPHA * (current_value - last_value);
}

volatile uint32_t uwTick;

void TIM6_IRQHandler(void)
{
	if(TIM_GetITStatus(TIM6, TIM_IT_Update) != RESET )
	{
		TIM_ClearITPendingBit(TIM6,TIM_IT_Update);
		uwTick += 5;
		Flag_Target=!Flag_Target;
		if(Flag_Target==1)
		{
			Current_EncoderD=Encoder_App_ReadMotorD();
			Current_EncoderC=Encoder_App_ReadMotorC();
			if(Flag_Show) Led_Flash(100);
			else Led_Flash(0);
		}
		if(Motor_App_TurnOff(Voltage)==0)
		{
			Motor_PwmC = PID_App_RunMotorC(Current_EncoderC,Target_EncoderC);
			Motor_PwmD = PID_App_RunMotorD(Current_EncoderD,Target_EncoderD);
			PID_App_LimitPwm(7199);
			Motor_App_SetPwm(Motor_PwmC,Motor_PwmD);
		}
	}
}

void Control_TaskDelayTick(void)
{
	if(delay_flag==1)
	{
		if(++delay_50==10)
		{
			delay_50=0;
			delay_flag=0;
		}
	}
}

u32 myabs(long int a)
{
	u32 temp;
	if(a<0) temp=-a;
	else temp=a;
	return temp;
}

float target_limit_float(float insert,float low,float high)
{
	if(insert < low)
		return low;
	else if(insert > high)
		return high;
	else
		return insert;
}

void Get_Angle(void)
{
	static float roll_filtered;
	static float pitch_filtered;
	static float yaw_filtered;
	ATTITUDE_DATA_t attitude;
	IMU_DATA_t imu_data;
	float roll_raw;
	float pitch_raw;
	float yaw_raw;

	IMU_App_Update();
	IMU_App_GetAttitude(&attitude);
	IMU_App_GetGyro(&imu_data);

	roll_raw = attitude.roll*57.2958f;
	pitch_raw = attitude.pitch*57.2958f;
	yaw_raw = attitude.yaw*57.2958f;

	roll_filtered = LowPassFilter(roll_raw, roll_filtered);
	pitch_filtered = LowPassFilter(pitch_raw, pitch_filtered);
	yaw_filtered = LowPassFilter(yaw_raw, yaw_filtered);

	Roll = roll_filtered;
	Pitch = pitch_filtered;
	Yaw = yaw_filtered;
	Gyrox = imu_data.gyro.x*57.2958f;
	Gyroy = imu_data.gyro.y*57.2958f;
	Gyroz = imu_data.gyro.z*57.2958f;
//	USART_App_Printf(USART1,"R %f;P:%f;Y:%f",Roll,Pitch,Yaw);
}
