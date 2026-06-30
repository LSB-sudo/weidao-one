#include "usart_app.h"
#include "pid_app.h"
#include "schduler.h"
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_RPM_TO_ENCODER_RPM       268.0f
#define MAX_RPM_TO_ENCODER_TARGET    35.0f
#define BLE_RPM_MAX                  268
#define BLE_CMD_BUFFER_SIZE          64
#define BLE_RPM_TIMEOUT_MS           500
#define USART_APP_PRINTF_BUFFER_SIZE 128

static char Debug_Cmd_Buffer[32];
static u8 Debug_Cmd_Count;
static char Ble_Cmd_Buffer[BLE_CMD_BUFFER_SIZE];
static u8 Ble_Cmd_Count;
static u8 Ble_Battery_Report_Enabled=1;
static volatile uint32_t Ble_LastValidRpmTick;
static volatile u8 Ble_Rpm_Control_Active;
extern u8 Receive_Data[RECEIVE_DATA_SIZE];
int Cnt_test;

int USART_App_Printf(USART_TypeDef *USARTx, const char *format, ...)
{
	char buffer[USART_APP_PRINTF_BUFFER_SIZE];
	va_list args;
	int length;
	int index;

	if(USARTx == 0 || format == 0)
	{
		return -1;
	}

	va_start(args, format);
	length = vsnprintf(buffer, sizeof(buffer), format, args);
	va_end(args);

	if(length < 0 || length >= (int)sizeof(buffer))
	{
		return -1;
	}

	for(index = 0; index < length; index++)
	{
		while(USART_GetFlagStatus(USARTx, USART_FLAG_TXE) == RESET)
		{
		}
		USART_SendData(USARTx, (u16)buffer[index]);
	}

	return length;
}

int USART_App_RpmToEncoderTarget(float rpm)
{
	float encoder_target;

	encoder_target = rpm * (MAX_RPM_TO_ENCODER_TARGET / MAX_RPM_TO_ENCODER_RPM);

	if(encoder_target > MAX_RPM_TO_ENCODER_TARGET)
	{
		encoder_target = MAX_RPM_TO_ENCODER_TARGET;
	}
	else if(encoder_target < -MAX_RPM_TO_ENCODER_TARGET)
	{
		encoder_target = -MAX_RPM_TO_ENCODER_TARGET;
	}

	if(encoder_target >= 0)
	{
		return (int)(encoder_target + 0.5f);
	}

	return (int)(encoder_target - 0.5f);
}

static int USART_App_ClampEncoderTarget(int encoder_target)
{
	if(encoder_target > (int)MAX_RPM_TO_ENCODER_TARGET)
	{
		return (int)MAX_RPM_TO_ENCODER_TARGET;
	}

	if(encoder_target < -(int)MAX_RPM_TO_ENCODER_TARGET)
	{
		return -(int)MAX_RPM_TO_ENCODER_TARGET;
	}

	return encoder_target;
}

static void USART_App_SendBleString(const char *text)
{
	USART3_SendBuffer((const u8 *)text, (u16)strlen(text));
}

static void USART_App_StopBleRpmControl(void)
{
	Target_EncoderC = 0;
	Target_EncoderD = 0;
	Flag_Stop = 1;
	PID_App_Reset();
	PWMC = 0;
	PWMD = 0;
}

static int USART_App_ParseBleRpmCommand(char *cmd_buffer, int *left_rpm, int *right_rpm)
{
	char *left_ptr;
	char *right_ptr;
	char *end_ptr;
	long parsed_left;
	long parsed_right;

	if(strncmp(cmd_buffer, "RPM,", 4) != 0)
	{
		return 0;
	}

	left_ptr = strstr(cmd_buffer, "L:");
	right_ptr = strstr(cmd_buffer, "R:");
	if(left_ptr == 0 || right_ptr == 0)
	{
		return 0;
	}

	parsed_left = strtol(left_ptr + 2, &end_ptr, 10);
	if(end_ptr == left_ptr + 2 || *end_ptr != ',')
	{
		return 0;
	}

	parsed_right = strtol(right_ptr + 2, &end_ptr, 10);
	if(end_ptr == right_ptr + 2 || *end_ptr != '\0')
	{
		return 0;
	}

	if(parsed_left < -BLE_RPM_MAX || parsed_left > BLE_RPM_MAX ||
	   parsed_right < -BLE_RPM_MAX || parsed_right > BLE_RPM_MAX)
	{
		return -1;
	}

	*left_rpm = (int)parsed_left;
	*right_rpm = (int)parsed_right;
	return 1;
}

static void USART_App_ProcessBleCommand(char *cmd_buffer)
{
	int left_rpm;
	int right_rpm;
	int parse_result;

	if(strcmp(cmd_buffer, "BAT:1") == 0)
	{
		Ble_Battery_Report_Enabled = 1;
		USART_App_SendBleString("OK,BAT:1\r\n");
		return;
	}

	if(strcmp(cmd_buffer, "BAT:0") == 0)
	{
		Ble_Battery_Report_Enabled = 0;
		USART_App_SendBleString("OK,BAT:0\r\n");
		return;
	}

	parse_result = USART_App_ParseBleRpmCommand(cmd_buffer, &left_rpm, &right_rpm);
	if(parse_result == 1)
	{
		Ble_LastValidRpmTick = Scheduler_GetTickMs();
		Target_EncoderC = USART_App_RpmToEncoderTarget((float)left_rpm);
		Target_EncoderD = USART_App_RpmToEncoderTarget((float)right_rpm);
		if(left_rpm == 0 && right_rpm == 0)
		{
			Ble_Rpm_Control_Active = 0;
			USART_App_StopBleRpmControl();
		}
		else
		{
			Ble_Rpm_Control_Active = 1;
			Flag_Stop = 0;
		}
		USART_App_SendBleString("OK\r\n");
	}
	else if(parse_result == -1)
	{
		USART_App_SendBleString("ERR,RANGE\r\n");
	}
	else
	{
		USART_App_SendBleString("ERR,FORMAT\r\n");
	}
}

u8 USART_App_IsBleBatteryReportEnabled(void)
{
	return Ble_Battery_Report_Enabled;
}

void USART_App_BleTimeoutTask(void)
{
	if(Ble_Rpm_Control_Active != 0 &&
	   (uint32_t)(Scheduler_GetTickMs() - Ble_LastValidRpmTick) >= BLE_RPM_TIMEOUT_MS)
	{
		Ble_Rpm_Control_Active = 0;
		USART_App_StopBleRpmControl();
	}
}

static void USART_App_ProcessDebugCommand(char *cmd_buffer)
{
	char *tc_ptr;
	char *td_ptr;
	float target_rpm;

	tc_ptr = strstr(cmd_buffer,"TC:");
	if(tc_ptr)
	{
		target_rpm = (float)atof(tc_ptr + 3);
		Target_EncoderC = USART_App_RpmToEncoderTarget(target_rpm);
		/* printf("Set Target_EncoderC=%d from TC:%.2f rpm, Current_EncoderC=%d\r\n", Target_EncoderC, target_rpm, Current_EncoderC); */
	}

	td_ptr = strstr(cmd_buffer,"TD:");
	if(td_ptr)
	{
		target_rpm = (float)atof(td_ptr + 3);
		Target_EncoderD = USART_App_RpmToEncoderTarget(target_rpm);
		/* printf("Set Target_EncoderD=%d from TD:%.2f rpm, Current_EncoderD=%d\r\n", Target_EncoderD, target_rpm, Current_EncoderD); */
	}
}

static void USART_App_ProcessFrame(u8 *receive_data)
{
	Target_EncoderC = USART_App_ClampEncoderTarget((short)Target_Speed_transition(receive_data[5],receive_data[6]));
	Target_EncoderD = USART_App_ClampEncoderTarget((short)Target_Speed_transition(receive_data[7],receive_data[8]));
}

void USART_App_IRQHandler(void)
{
	int usart_receive;
	static int Count;

	USART_ClearITPendingBit(USART1,USART_IT_RXNE);
	usart_receive = USART1->DR;

	if(usart_receive == '\r' || usart_receive == '\n')
	{
		if(Debug_Cmd_Count > 0)
		{
			Debug_Cmd_Buffer[Debug_Cmd_Count] = '\0';
			USART_App_ProcessDebugCommand(Debug_Cmd_Buffer);
			Debug_Cmd_Count = 0;
		}
	}
	else if(Debug_Cmd_Count < sizeof(Debug_Cmd_Buffer) - 1)
	{
		Debug_Cmd_Buffer[Debug_Cmd_Count++] = (char)usart_receive;
	}
	else
	{
		Debug_Cmd_Count = 0;
	}

	Cnt_test++;
	Receive_Data[Count] = usart_receive;
	if(usart_receive == FRAME_HEADER||Count>0)
		Count++;
	else
		Count=0;

	if(Count==11)
	{
		Count = 0;
		if(Receive_Data[10] == FRAME_TAIL)
		{
			if(Receive_Data[9] == Check_Sum(9))
			{
				USART_App_ProcessFrame(Receive_Data);
			}
		}
	}
}

void USART_App_BleByteReceived(u8 data)
{
	if(data == '\r' || data == '\n')
	{
		if(Ble_Cmd_Count > 0)
		{
			Ble_Cmd_Buffer[Ble_Cmd_Count] = '\0';
			USART_App_ProcessBleCommand(Ble_Cmd_Buffer);
			Ble_Cmd_Count = 0;
		}
	}
	else if(Ble_Cmd_Count < sizeof(Ble_Cmd_Buffer) - 1)
	{
		Ble_Cmd_Buffer[Ble_Cmd_Count++] = (char)data;
	}
	else
	{
		Ble_Cmd_Count = 0;
		USART_App_SendBleString("ERR,LEN\r\n");
	}
}
