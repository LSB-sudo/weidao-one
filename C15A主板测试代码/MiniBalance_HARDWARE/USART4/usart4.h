/***********************************************
公司：轮趣科技（东莞）有限公司
品牌：WHEELTEC
官网：wheeltec.net
淘宝店铺：shop114407458.taobao.com 
速卖通: https://minibalance.aliexpress.com/store/4455017
版本：
修改时间：2026-05-26

Brand: WHEELTEC
Website: wheeltec.net
Taobao shop: shop114407458.taobao.com 
Aliexpress: https://minibalance.aliexpress.com/store/4455017
Version:
Update：2026-05-26

All rights reserved
***********************************************/
#ifndef __USART4_H
#define __USART4_H

#include "sys.h"

void uart4_init(u32 bound);
void UART4_IRQHandler(void);
void UART4_SendByte(u8 data);
void UART4_SendBuffer(const u8 *buffer, u16 length);

#endif
