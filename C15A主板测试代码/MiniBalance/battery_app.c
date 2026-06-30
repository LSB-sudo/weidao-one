#include "battery_app.h"
#include "adc.h"

static u8 voltage_count;
static float voltage_sum;
static volatile u8 voltage_updated;

int Voltage;

void Battery_App_UpdateSample(void)
{
	voltage_sum += Get_battery_volt();
	if(++voltage_count == 10)
	{
		Voltage = voltage_sum / 10;
		voltage_sum = 0;
		voltage_count = 0;
		voltage_updated = 1;
	}
}

u8 Battery_App_TakeVoltageUpdated(void)
{
	u8 updated;

	INTX_DISABLE();
	updated = voltage_updated;
	voltage_updated = 0;
	INTX_ENABLE();

	return updated;
}

void Battery_App_PrintVoltage(void)
{
	printf("Battery Voltage:%d.%02dV\r\n", Voltage/100, Voltage%100);
}


void Scheduler_BatteryReportTask(void)
{
	Battery_App_UpdateSample();
	USART_App_Printf(USART3, "Battery Voltage:%d.%02dV\r\n", Voltage / 100, Voltage % 100);
}
