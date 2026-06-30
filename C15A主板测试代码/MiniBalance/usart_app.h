#ifndef __USART_APP_H
#define __USART_APP_H

#include "mydefine.h"

void USART_App_IRQHandler(void);
int USART_App_Printf(USART_TypeDef *USARTx, const char *format, ...);
int USART_App_RpmToEncoderTarget(float rpm);
void USART_App_BleByteReceived(u8 data);
u8 USART_App_IsBleBatteryReportEnabled(void);
void USART_App_BleTimeoutTask(void);

#endif
