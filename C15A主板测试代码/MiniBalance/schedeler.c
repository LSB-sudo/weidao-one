#include "schduler.h"
#include "battery_app.h"
#include "control.h"
#include "key.h"
#include "rdk_app.h"
#include "usart_app.h"

typedef struct
{
	void (*task_func)(void);
	uint32_t rate_ms;
	uint32_t last_run;
} task_t;

static uint8_t task_num;

static task_t scheduler_task[] =
{
	{RDK_App_Task, 5, 0},
	{RDK_App_FeedbackTask, 20, 0},
	{Control_TaskDelayTick, 5, 0},
	{Get_Angle, 5, 0},
	{Key, 5, 0},
	{USART_App_BleTimeoutTask, 20, 0},
	{Scheduler_BatteryReportTask, 100, 0},
};

uint32_t Scheduler_GetTickMs(void)
{
	return uwTick;
}

void scheduler_init(void)
{
	task_num = sizeof(scheduler_task) / sizeof(task_t);
}

void scheduler_run(void)
{
	uint8_t i;
	uint32_t now_time;

	for(i = 0; i < task_num; i++)
	{
		now_time = Scheduler_GetTickMs();
		if((uint32_t)(now_time - scheduler_task[i].last_run) >= scheduler_task[i].rate_ms)
		{
			scheduler_task[i].last_run = now_time;
			scheduler_task[i].task_func();
		}
	}
}
