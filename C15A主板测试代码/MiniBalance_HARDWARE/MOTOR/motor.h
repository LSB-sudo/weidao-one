/***********************************************
公司：轮趣科技（东莞）有限公司
品牌：WHEELTEC
官网：wheeltec.net
淘宝店铺：shop114407458.taobao.com 
速卖通: https://minibalance.aliexpress.com/store/4455017
版本：5.7
修改时间：2021-04-29

 
Brand: WHEELTEC
Website: wheeltec.net
Taobao shop: shop114407458.taobao.com 
Aliexpress: https://minibalance.aliexpress.com/store/4455017
Version:5.7
Update：2021-04-29

All rights reserved
***********************************************/
#ifndef __MOTOR_H
#define __MOTOR_H
#include <sys.h>	 

//这里是电机输出的PWM          //PWMX_IN1 为PWM输入时，PWMX_IN2 没有PWM输入时，车轮正转快衰竭
                               //PWMX_IN1 为1输入时，PWMX_IN2 有PWM输入时，车轮正转慢衰竭
  
#define PWMD TIM8->CCR3   
#define PWMA TIM8->CCR2  
#define PWMC TIM8->CCR1  
#define PWMB TIM8->CCR4
 
#define DIR2   PAout(5)
#define EN2    PCout(5)
#define DIR3   PBout(14)
#define EN3    PBout(13)
#define DIR1   PCout(4)
#define EN1    PCout(3)
#define AIN1   PAout(2)
#define AIN2   PAout(3)

void MiniBalance_PWM_Init(u16 arr,u16 psc);
void Motor_Init(void);
#endif
