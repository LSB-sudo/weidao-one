#ifndef __SCHDULER_H
#define __SCHDULER_H

#include "mydefine.h"

extern volatile uint32_t uwTick;

void scheduler_init(void);
void scheduler_run(void);
uint32_t Scheduler_GetTickMs(void);

#endif
