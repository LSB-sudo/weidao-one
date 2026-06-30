#ifndef __RDK_APP_H
#define __RDK_APP_H

#include "mydefine.h"

void RDK_App_Init(void);
void RDK_App_ByteReceived(u8 data);
void RDK_App_Task(void);
void RDK_App_FeedbackTask(void);

#endif
