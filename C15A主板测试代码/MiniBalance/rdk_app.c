#include "rdk_app.h"
#include "battery_app.h"
#include "encoder_app.h"
#include "pid_app.h"
#include "schduler.h"
#include "ultrasonic_app.h"
#include "usart4.h"
#include "usart_app.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define RDK_APP_CMD_BUFFER_SIZE        96
#define RDK_APP_BIRD_COMMAND           "BIRD"
#define RDK_APP_BIRD_ACTIVE_TICKS      1000
#define RDK_APP_FEEDBACK_INTERVAL_MS   100
#define RDK_APP_MAX_ENCODER_TARGET     35

static char RDK_Cmd_Buffer[RDK_APP_CMD_BUFFER_SIZE];
static u8 RDK_Cmd_Count;
static u8 RDK_Bird_Active;
static uint32_t RDK_Bird_StopTick;

static int RDK_App_ClampEncoderTarget(int target)
{
	if(target > RDK_APP_MAX_ENCODER_TARGET)
	{
		return RDK_APP_MAX_ENCODER_TARGET;
	}

	if(target < -RDK_APP_MAX_ENCODER_TARGET)
	{
		return -RDK_APP_MAX_ENCODER_TARGET;
	}

	return target;
}

static void RDK_App_StopBoat(void)
{
	Target_EncoderC = 0;
	Target_EncoderD = 0;
	Flag_Stop = 1;
	PID_App_Reset();
	PWMC = 0;
	PWMD = 0;
}

static void RDK_App_SendString(const char *text)
{
	if(text != 0)
	{
		UART4_SendBuffer((const u8 *)text, (u16)strlen(text));
	}
}

static void RDK_App_SendFeedback(void)
{
	char buffer[128];
	float left_rpm;
	float right_rpm;
	float battery_voltage;
	int length;

	left_rpm = (float)Current_EncoderC * 268.0f / 35.0f;
	right_rpm = (float)Current_EncoderD * 268.0f / 35.0f;
	battery_voltage = (float)Voltage / 100.0f;
	length = snprintf(buffer, sizeof(buffer),
		"FB,left_rpm=%.3f,right_rpm=%.3f,battery_voltage=%.2f\r\n",
		left_rpm, right_rpm, battery_voltage);
	if(length > 0 && length < (int)sizeof(buffer))
	{
		UART4_SendBuffer((const u8 *)buffer, (u16)length);
	}
}

static int RDK_App_ParseBoatRun(const char *cmd_buffer, int *boat_run)
{
	char *boat_run_ptr;
	char *end_ptr;
	long parsed_value;

	boat_run_ptr = strstr(cmd_buffer, "boat_run=");
	if(boat_run_ptr == 0)
	{
		return 0;
	}

	parsed_value = strtol(boat_run_ptr + 9, &end_ptr, 10);
	if(end_ptr == boat_run_ptr + 9 || *end_ptr != '\0')
	{
		return 0;
	}

	if(parsed_value != 0 && parsed_value != 1)
	{
		return 0;
	}

	*boat_run = (int)parsed_value;
	return 1;
}

static int RDK_App_ParseMotorRpmCommand(const char *cmd_buffer, float *left_rpm, float *right_rpm, int *boat_run)
{
	char *left_ptr;
	char *right_ptr;
	char *end_ptr;
	float parsed_left;
	float parsed_right;

	if(strncmp(cmd_buffer, "CMD,", 4) != 0)
	{
		return 0;
	}

	left_ptr = strstr(cmd_buffer, "left_set_rpm=");
	right_ptr = strstr(cmd_buffer, "right_set_rpm=");
	if(left_ptr == 0 || right_ptr == 0)
	{
		return 0;
	}

	parsed_left = (float)strtod(left_ptr + 13, &end_ptr);
	if(end_ptr == left_ptr + 13 || *end_ptr != ',')
	{
		return 0;
	}

	parsed_right = (float)strtod(right_ptr + 14, &end_ptr);
	if(end_ptr == right_ptr + 14 || *end_ptr != ',')
	{
		return 0;
	}

	if(RDK_App_ParseBoatRun(cmd_buffer, boat_run) == 0)
	{
		return 0;
	}

	*left_rpm = parsed_left;
	*right_rpm = parsed_right;
	return 1;
}

static void RDK_App_ProcessBirdCommand(void)
{
	if(RDK_Bird_Active == 0)
	{
		RDK_Bird_Active = 1;
		RDK_Bird_StopTick = Scheduler_GetTickMs() + RDK_APP_BIRD_ACTIVE_TICKS;
		Ultrasonic_App_On();
	}
}

static void RDK_App_ProcessMotorCommand(const char *cmd_buffer)
{
	float left_rpm;
	float right_rpm;
	int boat_run;

	if(RDK_App_ParseMotorRpmCommand(cmd_buffer, &left_rpm, &right_rpm, &boat_run) == 0)
	{
		RDK_App_SendString("ERR,FORMAT\r\n");
		return;
	}

	Target_EncoderC = RDK_App_ClampEncoderTarget(USART_App_RpmToEncoderTarget(left_rpm));
	Target_EncoderD = RDK_App_ClampEncoderTarget(USART_App_RpmToEncoderTarget(right_rpm));
	if(boat_run == 0)
	{
		RDK_App_StopBoat();
	}
	else
	{
		Flag_Stop = 0;
	}
	RDK_App_SendString("OK\r\n");
}

static void RDK_App_ProcessCommand(char *cmd_buffer)
{
	if(strcmp(cmd_buffer, RDK_APP_BIRD_COMMAND) == 0)
	{
		RDK_App_ProcessBirdCommand();
		return;
	}

	if(strncmp(cmd_buffer, "CMD,", 4) == 0)
	{
		RDK_App_ProcessMotorCommand(cmd_buffer);
	}
}

void RDK_App_Init(void)
{
	RDK_Cmd_Count = 0;
	RDK_Bird_Active = 0;
	RDK_Bird_StopTick = 0;
	Ultrasonic_App_Off();
}

void RDK_App_ByteReceived(u8 data)
{
	if(data == '\r' || data == '\n')
	{
		if(RDK_Cmd_Count > 0)
		{
			RDK_Cmd_Buffer[RDK_Cmd_Count] = '\0';
			RDK_App_ProcessCommand(RDK_Cmd_Buffer);
			RDK_Cmd_Count = 0;
		}
	}
	else if(RDK_Cmd_Count < sizeof(RDK_Cmd_Buffer) - 1)
	{
		RDK_Cmd_Buffer[RDK_Cmd_Count++] = (char)data;
	}
	else
	{
		RDK_Cmd_Count = 0;
		RDK_App_SendString("ERR,LEN\r\n");
	}
}

void RDK_App_Task(void)
{
	if(RDK_Bird_Active != 0 && (int32_t)(Scheduler_GetTickMs() - RDK_Bird_StopTick) >= 0)
	{
		RDK_Bird_Active = 0;
		Ultrasonic_App_Off();
	}
}

void RDK_App_FeedbackTask(void)
{
	static uint32_t last_feedback_tick;
	uint32_t now_tick;

	now_tick = Scheduler_GetTickMs();
	if((uint32_t)(now_tick - last_feedback_tick) >= RDK_APP_FEEDBACK_INTERVAL_MS)
	{
		last_feedback_tick = now_tick;
		Battery_App_UpdateSample();
		RDK_App_SendFeedback();
	}
}
