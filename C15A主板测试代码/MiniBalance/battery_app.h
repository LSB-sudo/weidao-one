#ifndef __BATTERY_APP_H
#define __BATTERY_APP_H

#include "mydefine.h"

void Battery_App_UpdateSample(void);
u8 Battery_App_TakeVoltageUpdated(void);
void Battery_App_PrintVoltage(void);
void Scheduler_BatteryReportTask(void);

#endif
