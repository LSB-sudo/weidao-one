# STM32F103RCT6

**Pages**: 1-144

---

**📄 Source: PDF Page 1**

This is information on a product in full production. 
November 2015
DocID14611 Rev 12
1/144
STM32F103xC, STM32F103xD,
STM32F103xE
High-density performance line ARM®-based 32-bit MCU with 256 to 512KB
Flash, USB, CAN, 11 timers, 3 ADCs, 13 communication interfaces
Datasheet − production data
Features
•
Core: ARM® 32-bit Cortex®-M3 CPU
–
72 MHz maximum frequency, 1.25 DMIPS/MHz 
(Dhrystone 2.1) performance at 0 wait state 
memory access
–
Single-cycle multiplication and hardware 
division
•
Memories
–
256 to 512 Kbytes of Flash memory
–
up to 64 Kbytes of SRAM
–
Flexible static memory controller with 4 Chip 
Select. Supports Compact Flash, SRAM, 
PSRAM, NOR and NAND memories
–
LCD parallel interface, 8080/6800 modes
•
Clock, reset and supply management
–
2.0 to 3.6 V application supply and I/Os 
–
POR, PDR, and programmable voltage detector 
(PVD)
–
4-to-16 MHz crystal oscillator 
–
Internal 8 MHz factory-trimmed RC
–
Internal 40 kHz RC with calibration 
–
32 kHz oscillator for RTC with calibration
•
Low power
–
Sleep, Stop and Standby modes
–
VBAT supply for RTC and backup registers
•
3 × 12-bit, 1 µs A/D converters (up to 21 
channels)
–
Conversion range: 0 to 3.6 V
–
Triple-sample and hold capability
–
Temperature sensor
•
2 × 12-bit D/A converters
•
DMA: 12-channel DMA controller
–
Supported peripherals: timers, ADCs, DAC, 
SDIO, I2Ss, SPIs, I2Cs and USARTs
•
Debug mode
–
Serial wire debug (SWD) & JTAG interfaces
–
Cortex®-M3 Embedded Trace Macrocell™
•
Up to 112 fast I/O ports
–
51/80/112 I/Os, all mappable on 16 external 
interrupt vectors and almost all 5 V-tolerant
•
Up to 11 timers
–
Up to four 16-bit timers, each with up to 4 
IC/OC/PWM or pulse counter and quadrature 
(incremental) encoder input
–
2 × 16-bit motor control PWM timers with dead-
time generation and emergency stop
–
2 × watchdog timers (Independent and Window)
–
SysTick timer: a 24-bit downcounter
–
2 × 16-bit basic timers to drive the DAC
•
Up to 13 communication interfaces
–
Up to 2 × I2C interfaces (SMBus/PMBus)
–
Up to 5 USARTs (ISO 7816 interface, LIN, IrDA 
capability, modem control)
–
Up to 3 SPIs (18 Mbit/s), 2 with I2S interface 
multiplexed
–
CAN interface (2.0B Active)
–
USB 2.0 full speed interface
–
SDIO interface
•
CRC calculation unit, 96-bit unique ID
•
ECOPACK® packages
         
Table 1.Device summary
Reference
Part number
STM32F103xC
STM32F103RC STM32F103VC 
STM32F103ZC
STM32F103xD
STM32F103RD STM32F103VD 
STM32F103ZD
STM32F103xE
STM32F103RE STM32F103ZE 
STM32F103VE
LQFP64 10 × 10 mm, 
LQFP100 14 × 14 mm, 
LQFP144 20 × 20 mm
LFBGA100 10 × 10 mm
LFBGA144 10 × 10 mm
WLCSP64
www.st.com

---

---

**📄 Source: PDF Page 2**

Contents
STM32F103xC, STM32F103xD, STM32F103xE
2/144
DocID14611 Rev 12
Contents
1
Introduction . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 9
2
Description . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 10
2.1
Device overview . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 11
2.2
Full compatibility throughout the family  . . . . . . . . . . . . . . . . . . . . . . . . . . 14
2.3
Overview  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 15
2.3.1
ARM® Cortex®-M3 core with embedded Flash and SRAM . . . . . . . . . . 15
2.3.2
Embedded Flash memory  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 15
2.3.3
CRC (cyclic redundancy check) calculation unit  . . . . . . . . . . . . . . . . . . 15
2.3.4
Embedded SRAM  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 15
2.3.5
FSMC (flexible static memory controller) . . . . . . . . . . . . . . . . . . . . . . . . 15
2.3.6
LCD parallel interface  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 16
2.3.7
Nested vectored interrupt controller (NVIC) . . . . . . . . . . . . . . . . . . . . . . 16
2.3.8
External interrupt/event controller (EXTI)  . . . . . . . . . . . . . . . . . . . . . . . 16
2.3.9
Clocks and startup . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 16
2.3.10
Boot modes  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 17
2.3.11
Power supply schemes  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 17
2.3.12
Power supply supervisor  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 17
2.3.13
Voltage regulator  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 17
2.3.14
Low-power modes  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 18
2.3.15
DMA . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 18
2.3.16
RTC (real-time clock) and backup registers . . . . . . . . . . . . . . . . . . . . . . 18
2.3.17
Timers and watchdogs . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 19
2.3.18
I²C bus . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 21
2.3.19
Universal synchronous/asynchronous receiver transmitters (USARTs)  21
2.3.20
Serial peripheral interface (SPI) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 21
2.3.21
Inter-integrated sound (I2S) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 21
2.3.22
SDIO  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 22
2.3.23
Controller area network (CAN)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 22
2.3.24
Universal serial bus (USB) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 22
2.3.25
GPIOs (general-purpose inputs/outputs) . . . . . . . . . . . . . . . . . . . . . . . . 22
2.3.26
ADC (analog to digital converter) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 22
2.3.27
DAC (digital-to-analog converter) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 23
2.3.28
Temperature sensor . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 24

---

---

**📄 Source: PDF Page 3**

DocID14611 Rev 12
3/144
STM32F103xC, STM32F103xD, STM32F103xE
Contents
4
2.3.29
Serial wire JTAG debug port (SWJ-DP) . . . . . . . . . . . . . . . . . . . . . . . . . 24
2.3.30
Embedded Trace Macrocell™  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 24
3
Pinouts and pin descriptions . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 25
4
Memory mapping  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 40
5
Electrical characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1
Parameter conditions . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1.1
Minimum and maximum values . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1.2
Typical values . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1.3
Typical curves  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1.4
Loading capacitor  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1.5
Pin input voltage  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
5.1.6
Power supply scheme  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 42
5.1.7
Current consumption measurement  . . . . . . . . . . . . . . . . . . . . . . . . . . . 42
5.2
Absolute maximum ratings . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 43
5.3
Operating conditions  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 44
5.3.1
General operating conditions . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 44
5.3.2
Operating conditions at power-up / power-down  . . . . . . . . . . . . . . . . . . 45
5.3.3
Embedded reset and power control block characteristics  . . . . . . . . . . . 45
5.3.4
Embedded reference voltage . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 46
5.3.5
Supply current characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 46
5.3.6
External clock source characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . 58
5.3.7
Internal clock source characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . 62
5.3.8
PLL characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 64
5.3.9
Memory characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 64
5.3.10
FSMC characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 66
5.3.11
EMC characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 87
5.3.12
Absolute maximum ratings (electrical sensitivity)  . . . . . . . . . . . . . . . . . 88
5.3.13
I/O current injection characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . 89
5.3.14
I/O port characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 90
5.3.15
NRST pin characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 95
5.3.16
TIM timer characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 96
5.3.17
Communications interfaces  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 97
5.3.18
CAN (controller area network) interface . . . . . . . . . . . . . . . . . . . . . . . . 107
5.3.19
12-bit ADC characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 107

---

---

**📄 Source: PDF Page 4**

Contents
STM32F103xC, STM32F103xD, STM32F103xE
4/144
DocID14611 Rev 12
5.3.20
DAC electrical specifications  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 112
5.3.21
Temperature sensor characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . 114
6
Package information . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 115
6.1
LFBGA144 package information  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 115
6.2
LFBGA100 package information  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 118
6.3
WLCSP64 package information . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 121
6.4
LQFP144 package information . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 123
6.5
LQFP100 package information . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 127
6.6
LQFP64 package information . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 130
6.7
Thermal characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 133
6.7.1
Reference document  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 133
6.7.2
Selecting the product temperature range . . . . . . . . . . . . . . . . . . . . . . . 134
7
Part numbering  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 136
8
Revision history  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 137

---

---

**📄 Source: PDF Page 5**

DocID14611 Rev 12
5/144
STM32F103xC, STM32F103xD, STM32F103xE
List of tables
6
List of tables
Table 1.
Device summary . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 1
Table 2.
STM32F103xC, STM32F103xD and STM32F103xE features
and peripheral counts . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 11
Table 3.
STM32F103xx family  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 14
Table 4.
High-density timer feature comparison . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 19
Table 5.
High-density STM32F103xC/D/E pin definitions. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 31
Table 6.
FSMC pin definition  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 38
Table 7.
Voltage characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 43
Table 8.
Current characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 43
Table 9.
Thermal characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 44
Table 10.
General operating conditions . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 44
Table 11.
Operating conditions at power-up / power-down  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 45
Table 12.
Embedded reset and power control block characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . 45
Table 13.
Embedded internal reference voltage. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 46
Table 14.
Maximum current consumption in Run mode, code with data processing
running from Flash . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 47
Table 15.
Maximum current consumption in Run mode, code with data processing
running from RAM. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 47
Table 16.
Maximum current consumption in Sleep mode, code running from Flash or RAM. . . . . . . 49
Table 17.
Typical and maximum current consumptions in Stop and Standby modes  . . . . . . . . . . . . 50
Table 18.
Typical current consumption in Run mode, code with data processing
running from Flash . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 53
Table 19.
Typical current consumption in Sleep mode, code running from Flash or
RAM . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 54
Table 20.
Peripheral current consumption . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 55
Table 21.
High-speed external user clock characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 58
Table 22.
Low-speed external user clock characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 58
Table 23.
HSE 4-16 MHz oscillator characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 60
Table 24.
LSE oscillator characteristics (fLSE = 32.768 kHz) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 61
Table 25.
HSI oscillator characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 62
Table 26.
LSI oscillator characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 63
Table 27.
Low-power mode wakeup timings  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 63
Table 28.
PLL characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 64
Table 29.
Flash memory characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 64
Table 30.
Flash memory endurance and data retention . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 65
Table 31.
Asynchronous non-multiplexed SRAM/PSRAM/NOR read timings . . . . . . . . . . . . . . . . . . 67
Table 32.
Asynchronous non-multiplexed SRAM/PSRAM/NOR write timings . . . . . . . . . . . . . . . . . . 68
Table 33.
Asynchronous multiplexed PSRAM/NOR read timings. . . . . . . . . . . . . . . . . . . . . . . . . . . . 69
Table 34.
Asynchronous multiplexed PSRAM/NOR write timings  . . . . . . . . . . . . . . . . . . . . . . . . . . . 70
Table 35.
Synchronous multiplexed NOR/PSRAM read timings  . . . . . . . . . . . . . . . . . . . . . . . . . . . . 73
Table 36.
Synchronous multiplexed PSRAM write timings. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 75
Table 37.
Synchronous non-multiplexed NOR/PSRAM read timings . . . . . . . . . . . . . . . . . . . . . . . . . 76
Table 38.
Synchronous non-multiplexed PSRAM write timings . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 77
Table 39.
Switching characteristics for PC Card/CF read and write cycles . . . . . . . . . . . . . . . . . . . . 82
Table 40.
Switching characteristics for NAND Flash read and write cycles . . . . . . . . . . . . . . . . . . . . 86
Table 41.
EMS characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 87
Table 42.
EMI characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 88
Table 43.
ESD absolute maximum ratings . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 88

### Code Examples

```unknown
user clock characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 58
```

```unknown
user clock characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 58
```

---

---

**📄 Source: PDF Page 6**

List of tables
STM32F103xC, STM32F103xD, STM32F103xE
6/144
DocID14611 Rev 12
Table 44.
Electrical sensitivities . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 89
Table 45.
I/O current injection susceptibility . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 89
Table 46.
I/O static characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 90
Table 47.
Output voltage characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 92
Table 48.
I/O AC characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 94
Table 49.
NRST pin characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 95
Table 50.
TIMx characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 96
Table 51.
I2C characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 97
Table 52.
SCL frequency (fPCLK1= 36 MHz.,VDD_I2C = 3.3 V) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 98
Table 53.
SPI characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 99
Table 54.
I2S characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 102
Table 55.
SD / MMC characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 104
Table 56.
USB startup time. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 105
Table 57.
USB DC electrical characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 106
Table 58.
USB: full-speed electrical characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 106
Table 59.
ADC characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 107
Table 60.
RAIN max for fADC = 14 MHz. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 108
Table 61.
ADC accuracy - limited test conditions . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 108
Table 62.
ADC accuracy  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 109
Table 63.
DAC characteristics  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 112
Table 64.
TS characteristics . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 114
Table 65.
LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package mechanical data  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 115
Table 66.
LFBGA144 recommended PCB design rules (0.8 mm pitch BGA). . . . . . . . . . . . . . . . . . 116
Table 67.
LFBGA100 - 10 x 10 mm low profile fine pitch ball grid array package
mechanical data . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 118
Table 68.
LFBGA100 recommended PCB design rules (0.8 mm pitch BGA). . . . . . . . . . . . . . . . . . 119
Table 69.
WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm pitch, wafer-level chip-scale
package mechanical data . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 121
Table 70.
WLCSP64 recommended PCB design rules (0.5 mm pitch)  . . . . . . . . . . . . . . . . . . . . . . 122
Table 71.
LQFP144 - 144-pin, 20 x 20 mm low-profile quad flat package 
mechanical data . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 124
Table 72.
LQPF100 – 14 x 14 mm 100-pin low-profile quad flat package 
mechanical data . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 127
Table 73.
LQFP64 – 10 x 10 mm 64 pin low-profile quad flat package mechanical data. . . . . . . . . 130
Table 74.
Package thermal characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 133
Table 75.
Ordering information scheme . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 136

---

---

**📄 Source: PDF Page 7**

DocID14611 Rev 12
7/144
STM32F103xC, STM32F103xD, STM32F103xE
List of figures
8
List of figures
Figure 1.
STM32F103xC, STM32F103xD and STM32F103xE performance line block diagram  . . . 12
Figure 2.
Clock tree . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 13
Figure 3.
STM32F103xC/D/E BGA144 ballout  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 25
Figure 4.
STM32F103xC/D/E performance line BGA100 ballout. . . . . . . . . . . . . . . . . . . . . . . . . . . . 26
Figure 5.
STM32F103xC/D/E performance line LQFP144 pinout . . . . . . . . . . . . . . . . . . . . . . . . . . . 27
Figure 6.
STM32F103xC/D/E performance line LQFP100 pinout . . . . . . . . . . . . . . . . . . . . . . . . . . . 28
Figure 7.
STM32F103xC/D/E performance line LQFP64 pinout . . . . . . . . . . . . . . . . . . . . . . . . . . . . 29
Figure 8.
STM32F103xC/D/E performance line
WLCSP64 ballout, ball side  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 30
Figure 9.
Memory map. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 40
Figure 10.
Pin loading conditions. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
Figure 11.
Pin input voltage . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 41
Figure 12.
Power supply scheme. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 42
Figure 13.
Current consumption measurement scheme . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 42
Figure 14.
Typical current consumption in Run mode versus frequency (at 3.6 V) -
code with data processing running from RAM, peripherals enabled  . . . . . . . . . . . . . . . . . 48
Figure 15.
Typical current consumption in Run mode versus frequency (at 3.6 V)-
code with data processing running from RAM, peripherals disabled  . . . . . . . . . . . . . . . . 48
Figure 16.
Typical current consumption on VBAT with RTC on vs. temperature
at different VBAT values. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 50
Figure 17.
Typical current consumption in Stop mode with regulator in run mode
versus temperature at different VDD values
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 51
Figure 18.
Typical current consumption in Stop mode with regulator in low-power
mode versus temperature at different VDD values  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 51
Figure 19.
Typical current consumption in Standby mode versus temperature at
different VDD values . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 52
Figure 20.
High-speed external clock source AC timing diagram  . . . . . . . . . . . . . . . . . . . . . . . . . . . . 59
Figure 21.
Low-speed external clock source AC timing diagram. . . . . . . . . . . . . . . . . . . . . . . . . . . . . 59
Figure 22.
Typical application with an 8 MHz crystal . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 60
Figure 23.
Typical application with a 32.768 kHz crystal . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 62
Figure 24.
Asynchronous non-multiplexed SRAM/PSRAM/NOR read waveforms . . . . . . . . . . . . . . . 66
Figure 25.
Asynchronous non-multiplexed SRAM/PSRAM/NOR write waveforms . . . . . . . . . . . . . . . 67
Figure 26.
Asynchronous multiplexed PSRAM/NOR read waveforms. . . . . . . . . . . . . . . . . . . . . . . . . 69
Figure 27.
Asynchronous multiplexed PSRAM/NOR write waveforms  . . . . . . . . . . . . . . . . . . . . . . . . 70
Figure 28.
Synchronous multiplexed NOR/PSRAM read timings  . . . . . . . . . . . . . . . . . . . . . . . . . . . . 72
Figure 29.
Synchronous multiplexed PSRAM write timings. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 74
Figure 30.
Synchronous non-multiplexed NOR/PSRAM read timings . . . . . . . . . . . . . . . . . . . . . . . . . 76
Figure 31.
Synchronous non-multiplexed PSRAM write timings . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 77
Figure 32.
PC Card/CompactFlash controller waveforms for common memory read access . . . . . . . 78
Figure 33.
PC Card/CompactFlash controller waveforms for common memory write access. . . . . . . 79
Figure 34.
PC Card/CompactFlash controller waveforms for attribute memory read
access. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 80
Figure 35.
PC Card/CompactFlash controller waveforms for attribute memory write
access. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 81
Figure 36.
PC Card/CompactFlash controller waveforms for I/O space read access . . . . . . . . . . . . . 81
Figure 37.
PC Card/CompactFlash controller waveforms for I/O space write access . . . . . . . . . . . . . 82
Figure 38.
NAND controller waveforms for read access . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 84
Figure 39.
NAND controller waveforms for write access . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 85

---

---

**📄 Source: PDF Page 8**

List of figures
STM32F103xC, STM32F103xD, STM32F103xE
8/144
DocID14611 Rev 12
Figure 40.
NAND controller waveforms for common memory read access . . . . . . . . . . . . . . . . . . . . . 85
Figure 41.
NAND controller waveforms for common memory write access. . . . . . . . . . . . . . . . . . . . . 86
Figure 42.
Standard I/O input characteristics - CMOS port . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 91
Figure 43.
Standard I/O input characteristics - TTL port . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 91
Figure 44.
5 V tolerant I/O input characteristics - CMOS port . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 91
Figure 45.
5 V tolerant I/O input characteristics - TTL port  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 92
Figure 46.
I/O AC characteristics definition . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 95
Figure 47.
Recommended NRST pin protection  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 96
Figure 48.
I2C bus AC waveforms and measurement circuit . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 98
Figure 49.
SPI timing diagram - slave mode and CPHA = 0 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 100
Figure 50.
SPI timing diagram - slave mode and CPHA = 1(1)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 100
Figure 51.
SPI timing diagram - master mode(1) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 101
Figure 52.
I2S slave timing diagram (Philips protocol)(1) . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 103
Figure 53.
I2S master timing diagram (Philips protocol)(1). . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 103
Figure 54.
SDIO high-speed mode  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 104
Figure 55.
SD default mode . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 104
Figure 56.
USB timings: definition of data signal rise and fall time  . . . . . . . . . . . . . . . . . . . . . . . . . . 106
Figure 57.
ADC accuracy characteristics. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 109
Figure 58.
Typical connection diagram using the ADC . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 110
Figure 59.
Power supply and reference decoupling (VREF+ not connected to VDDA). . . . . . . . . . . . . 110
Figure 60.
Power supply and reference decoupling (VREF+ connected to VDDA). . . . . . . . . . . . . . . . 111
Figure 61.
12-bit buffered /non-buffered DAC . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 113
Figure 62.
LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package outline  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 115
Figure 63.
LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package recommended footprint . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 116
Figure 64.
LFBGA144 marking example (package top view)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 117
Figure 65.
LFBGA100 - 10 x 10 mm low profile fine pitch ball grid array package
outline . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 118
Figure 66.
LFBGA100 – 100-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package recommended footprintoutline . . . . . . . . . . . . . . . . . . . . . . . . . . . 119
Figure 67.
LFBGA100 marking example (package top view)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 120
Figure 68.
WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm pitch, wafer-level chip-scale
package outline. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 121
Figure 69.
WLCSP64 - 64-ball, 4.4757 x 4.4049 mm, 0.5 mm pitch wafer level chip scale
package recommended footprint  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 122
Figure 70.
LQFP144 - 144-pin, 20 x 20 mm low-profile quad flat package outline  . . . . . . . . . . . . . . 123
Figure 71.
LQFP144 - 144-pin,20 x 20 mm low-profile quad flat package 
recommended footprint. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 125
Figure 72.
LQFP144 marking example (package top view)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 126
Figure 73.
LQFP100 – 14 x 14 mm 100 pin low-profile quad flat package outline  . . . . . . . . . . . . . . 127
Figure 74.
LQFP100 recommended footprint  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 128
Figure 75.
LQFP100 marking example (package top view)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 129
Figure 76.
LQFP64 – 10 x 10 mm 64 pin low-profile quad flat package outline  . . . . . . . . . . . . . . . . 130
Figure 77.
LQFP64 - 64-pin, 10 x 10 mm low-profile quad flat recommended footprint  . . . . . . . . . . 131
Figure 78.
LQFP64 marking example (package top view)  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 132
Figure 79.
LQFP100 PD max vs. TA  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 135

---

---

**📄 Source: PDF Page 9**

DocID14611 Rev 12
9/144
STM32F103xC, STM32F103xD, STM32F103xE
Introduction
136
1 
Introduction
This datasheet provides the ordering information and mechanical device characteristics of 
the STM32F103xC, STM32F103xD and STM32F103xE high-density performance line 
microcontrollers. For more details on the whole STMicroelectronics STM32F103xC/D/E 
family, please refer to Section 2.2: Full compatibility throughout the family.
The high-density STM32F103xC/D/E datasheet should be read in conjunction with the 
STM32F10xxx reference manual.
For information on programming, erasing and protection of the internal Flash memory 
please refer to the STM32F10xxx Flash programming manual.
The reference and Flash programming manuals are both available from the 
STMicroelectronics website www.st.com.
For information on the Cortex®-M3 core please refer to the Cortex®-M3 Technical Reference 
Manual, available from the www.arm.com website at the following address: 
http://infocenter.arm.com.

---

---

**📄 Source: PDF Page 10**

Description
STM32F103xC, STM32F103xD, STM32F103xE
10/144
DocID14611 Rev 12
2 
Description
The STM32F103xC, STM32F103xD and STM32F103xE performance line family 
incorporates the high-performance ARM® Cortex®-M3 32-bit RISC core operating at a 
72 MHz frequency, high-speed embedded memories (Flash memory up to 512 Kbytes and 
SRAM up to 64 Kbytes), and an extensive range of enhanced I/Os and peripherals 
connected to two APB buses. All devices offer three 12-bit ADCs, four general-purpose 16-
bit timers plus two PWM timers, as well as standard and advanced communication 
interfaces: up to two I2Cs, three SPIs, two I2Ss, one SDIO, five USARTs, an USB and a 
CAN.
The STM32F103xC/D/E high-density performance line family operates in the –40 to 
+105 °C temperature range, from a 2.0 to 3.6 V power supply. A comprehensive set of 
power-saving mode allows the design of low-power applications.
These features make the STM32F103xC/D/E high-density performance line microcontroller 
family suitable for a wide range of applications such as motor drives, application control, 
medical and handheld equipment, PC and gaming peripherals, GPS platforms, industrial 
applications, PLCs, inverters, printers, scanners, alarm systems video intercom, and HVAC.

### Code Examples

```unknown
uses. All devices offer three 12-bit ADCs, four general-purpose 16-
included, the description below gives an overview of the complete range of
```

---

---

**📄 Source: PDF Page 11**

DocID14611 Rev 12
11/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
2.1 
Device overview
The STM32F103xC/D/E high-density performance line family offers devices in six different 
package types: from 64 pins to 144 pins. Depending on the device chosen, different sets of 
peripherals are included, the description below gives an overview of the complete range of 
peripherals proposed in this family.
Figure 1 shows the general block diagram of the device family.
         
Table 2. STM32F103xC, STM32F103xD and STM32F103xE features
and peripheral counts
Peripherals
STM32F103Rx
STM32F103Vx
STM32F103Zx
Flash memory in Kbytes
256
384
512
256
384
512
256
384
512
SRAM in Kbytes
48
64(1)
1.
64 KB RAM for 256 KB Flash are available on devices delivered in CSP packages only.
48
64
48
64
FSMC
No
Yes(2)
2.
For the LQFP100 and BGA100 packages, only FSMC Bank1 and Bank2 are available. Bank1 can only 
support a multiplexed NOR/PSRAM memory using the NE1 Chip Select. Bank2 can only support a 16- or 
8-bit NAND Flash memory using the NCE2 Chip Select. The interrupt line cannot be used since Port G is 
not available in this package.
Yes
Timers
General-purpose
4
Advanced-control
2
Basic
2
Comm
SPI(I2S)(3)
3.
The SPI2 and SPI3 interfaces give the flexibility to work in an exclusive way in either the SPI mode or the 
I2S audio mode.
3(2)
I2C
2
USART
5
USB
1
CAN
1
SDIO
1
GPIOs
51
80
112
12-bit ADC 
Number of channels
3
16
3
16
3
21
12-bit DAC
Number of channels
2
2
CPU frequency
72 MHz
Operating voltage
2.0 to 3.6 V
Operating temperatures
Ambient temperatures: –40 to +85 °C /–40 to +105 °C (see Table 10)
Junction temperature: –40 to + 125 °C (see Table 10)
Package
LQFP64, WLCSP64
LQFP100, BGA100
LQFP144, BGA144

### Code Examples

```unknown
used since Port G is
```

---

---

**📄 Source: PDF Page 12**

Description
STM32F103xC, STM32F103xD, STM32F103xE
12/144
DocID14611 Rev 12
Figure 1. STM32F103xC, STM32F103xD and STM32F103xE performance line block 
diagram
1.
TA = –40 °C to +85 °C (suffix 6, see Table 75) or –40 °C to +105 °C (suffix 7, see Table 75), junction 
temperature up to 105 °C or 125 °C, respectively.
2.
AF = alternate function on I/O port pin.9
0!;=
%84)4
!&
!("
XXB
IT	
7+50
&MAX  -(Z
633
)#
'0 $-!
4)-
4)-
84!, K(Z
&LASH  +BYTES
6$$
"ACKUP INTERFACE
4)-
"US -ATRIX  BIT
24#
2#  -(Z
#ORTEX- #05
$BUS
OBL
&LASH
INTERFACE
53!24
30)  )3
"ACKUP
REG
4)-
)#
28 48 #43 243
53!24
2#  K(Z
3TANDBY
)7$'
 6"!4
0/2  0$2
 6$$!
6"!4  6 TO  6
#+ AS !&
28 48 #43 243
#+ AS !&
.6)#
30)
 INTERFACE
 6$$!
06$
)NT
!0"
!75
4)-
 XXB
IT	
30)  )3
 5!24
2848 AS !&
5!24
2848 AS !&
4)-
0,,
 6$$!
&3-# $!#?/54 AS !&
$!#?/54 AS !&
32!-  +"
'0 $-!
4)-
4)-
.*4234
*4$)
*4#+37#,+
*4-337$)/
*4$/
AS !&
!;=
$;=
#,+
./%
.7%
.%;=
.",;=
.7!)4
., OR .!$6	
AS !&
 CHANNELS
 CHANNELS
'0)/ PORT !
'0)/ PORT "
'0)/ PORT #
'0)/ PORT $
'0)/ PORT %
'0)/ PORT &
'0)/ PORT '
53!24
4EMP SENSOR
BIT !$#
BIT !$#
BIT !$#
)&
)&
)&
0";=
0#;=
0$;=
0%;=
0&;=
0';=
 CHANNELS
 COMPL CHANNELS
"+). %42 AS !&
 CHANNELS
 COMPL CHANNELS
"+). %42 AS !&
-/3) -)3/
3#+ .33 AS !&
28 48 #43
243 #+ AS !&
 !$#?).S
COMMON TO THE  !$#S
 !$#?).S COMMON
TO !$#  !$#
 !$#?).S ON !$#
62%&
62%&n
  6$$!
!0" &MAX   -(Z
!0"
4RACE
CONTROLLER
0BUS
)BUS
3YSTEM
2ESET 
#LOCK
CONTROL
0#,+
0#,+
(#,+
&#,+
0OWER
6OLT REG
 6 TO  6
3UPPLY
SUPERVISION
 6$$
0/2
2ESET
.234
6$$!
633!
/3#?).
/3#?/54
 6$$
84!, /3#
 -(Z
/3#?).
/3#?/54
4!-0%224#
!,!2-3%#/.$ /54
 CHANNELS %42 AS !&
 CHANNELS %42 AS !&
 CHANNELS %42 AS !&
 CHANNELS AS !&
-/3)3$ -)3/
3#+#+ -#+ .3373 AS !&
-/3)3$ -)3/
3#+#+ -#+ .3373 AS !&
3#, 3$! 3-"! AS !&
3#, 3$! 3-"! AS !&
BX#!. DEVICE
53"  &3
DEVICE
53"?$0#!.?48
53"?$-#!.?28
32!-  "
77$'
AIG
!0" &MAX   -(Z
42!#%#,+
42!#%$;=
AS !3
37*4!'
40)5
4RACETRIG
3$)/
$;=
#-$
#+ AS !&
!(" &MAX   -(Z
!("
BIT  $!#
)&
)&
)&
BIT $!# 

---

---

**📄 Source: PDF Page 13**

DocID14611 Rev 12
13/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
Figure 2. Clock tree
1.
When the HSI is used as a PLL clock input, the maximum system clock frequency that can be achieved is 
64 MHz.
2.
For the USB function to be available, both HSE and PLL must be enabled, with the USBCLK at 48 MHz.
3.
To have an ADC conversion time of 1 µs, APB2 must be at 14 MHz, 28 MHz or 56 MHz.
+6(26&
0+]
26&B,1
26&B287
26&B,1
26&B287
/6(26&
N+]
+6,5&
0+]
/6,5&
N+]
WR,QGHSHQGHQW:DWFKGRJ,:'*
3//
[[[
3//08/
+6( +LJK6SHHG([WHUQDOFORFNVLJQDO
/6( /RZ6SHHG([WHUQDOFORFNVLJQDO
/6, /RZ6SHHG,QWHUQDOFORFNVLJQDO
+6, +LJK6SHHG,QWHUQDOFORFNVLJQDO
/HJHQG
0&2
&ORFN2XWSXW
0DLQ
3//;735(

[
$+%
3UHVFDOHU


3//&/.
+6,
+6(
$3%
3UHVFDOHU

$'&
3UHVFDOHU

$'&&/.
3&/.
+&/.
3//&/.
WR$+%EXVFRUH
PHPRU\DQG'0$
86%&/.
WR86%LQWHUIDFH
86%
3UHVFDOHU

WR$'&RU
/6(
/6,
+6,


+6,
+6(
SHULSKHUDOV
WR$3%
3HULSKHUDO&ORFN
(QDEOHELWV
(QDEOHELWV
3HULSKHUDO&ORFN
$3%
3UHVFDOHU

3&/.
7,0	WLPHUV
WR7,0DQG7,0
SHULSKHUDOVWR$3%
3HULSKHUDO&ORFN
(QDEOHELWV
(QDEOHELW
3HULSKHUDO&ORFN
0+]
0+]PD[
0+]
0+]PD[
0+]PD[
WR57&
3//65&
6:
0&2
&66
WR&RUWH[6\VWHPWLPHU

&ORFN
(QDEOHELWV
6<6&/.
PD[
57&&/.
57&6(/>@
7,0[&/.
7,0;&/.
,:'*&/.
6<6&/.
)&/.&RUWH[
IUHHUXQQLQJFORFN

7,0
WR7,0DQG
7R6',2$+%LQWHUIDFH
3HULSKHUDOFORFN
HQDEOH
+&/.
WR)60&
)60&&/.
WR6',2
3HULSKHUDOFORFN
HQDEOH
3HULSKHUDOFORFN
HQDEOH
WR,6
WR,6
3HULSKHUDOFORFN
HQDEOH
3HULSKHUDOFORFN
HQDEOH
,6&/.
,6&/.
6',2&/.
DLE
,I$3%SUHVFDOHU [
HOVH [
,I$3%SUHVFDOHU [
HOVH[
)/,7)&/.
WR)ODVKSURJUDPPLQJLQWHUIDFH

### Code Examples

```typescript
used as a PLL clock input, the maximum system clock frequency that can be achieved is
```

---

---

**📄 Source: PDF Page 14**

Description
STM32F103xC, STM32F103xD, STM32F103xE
14/144
DocID14611 Rev 12
2.2 
Full compatibility throughout the family
The STM32F103xC/D/E is a complete family whose members are fully pin-to-pin, software 
and feature compatible. In the reference manual, the STM32F103x4 and STM32F103x6 are 
identified as low-density devices, the STM32F103x8 and STM32F103xB are referred to as 
medium-density devices and the STM32F103xC, STM32F103xD and STM32F103xE are 
referred to as high-density devices.
Low-density and high-density devices are an extension of the STM32F103x8/B medium-
density devices, they are specified in the STM32F103x4/6 and STM32F103xC/D/E 
datasheets, respectively. Low-density devices feature lower Flash memory and RAM 
capacities, less timers and peripherals. High-density devices have higher Flash memory 
and RAM capacities, and additional peripherals like SDIO, FSMC, I2S and DAC while 
remaining fully compatible with the other members of the family.
The STM32F103x4, STM32F103x6, STM32F103xC, STM32F103xD and STM32F103xE 
are a drop-in replacement for the STM32F103x8/B devices, allowing the user to try different 
memory densities and providing a greater degree of freedom during the development cycle.
Moreover, the STM32F103xx performance line family is fully compatible with all existing 
STM32F101xx access line and STM32F102xx USB access line devices.
         
Table 3. STM32F103xx family
Pinout
Low-density devices
Medium-density devices
High-density devices
16 KB 
Flash
32 KB 
Flash(1)
1.
For orderable part numbers that do not show the A internal code after the temperature range code (6 or 7), 
the reference datasheet for electrical characteristics is that of the STM32F103x8/B medium-density 
devices.
64 KB 
Flash
128 KB 
Flash
256 KB 
Flash
384 KB 
Flash
512 KB 
Flash
6 KB RAM 10 KB RAM 20 KB RAM 20 KB RAM
48 RAM
64 KB RAM 64 KB RAM
144
5 × USARTs
4 × 16-bit timers, 2 × basic timers
3 × SPIs, 2 × I2Ss, 2 × I2Cs
USB, CAN, 2 × PWM timers
3 × ADCs, 2 × DACs, 1 × SDIO
FSMC (100- and 144-pin packages(2))
2.
Ports F and G are not available in devices delivered in 100-pin packages. 
100
3 × USARTs
3 × 16-bit timers
2 × SPIs, 2 × I2Cs, USB, 
CAN, 1 × PWM timer
2 × ADCs
64
2 × USARTs
2 × 16-bit timers
1 × SPI, 1 × I2C, USB, 
CAN, 1 × PWM timer
2 × ADCs
48
36

### Code Examples

```unknown
user to try different
```

---

---

**📄 Source: PDF Page 15**

DocID14611 Rev 12
15/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
2.3 
Overview
2.3.1 
ARM® Cortex®-M3 core with embedded Flash and SRAM
The ARM Cortex®-M3 processor is the latest generation of ARM processors for embedded 
systems. It has been developed to provide a low-cost platform that meets the needs of MCU 
implementation, with a reduced pin count and low-power consumption, while delivering 
outstanding computational performance and an advanced system response to interrupts.
The ARM Cortex®-M3 32-bit RISC processor features exceptional code-efficiency, 
delivering the high-performance expected from an ARM core in the memory size usually 
associated with 8- and 16-bit devices.
With its embedded ARM core, STM32F103xC, STM32F103xD and STM32F103xE 
performance line family is compatible with all ARM tools and software.
Figure 1 shows the general block diagram of the device family.
2.3.2 
Embedded Flash memory
Up to 512 Kbytes of embedded Flash is available for storing programs and data. 
2.3.3 
CRC (cyclic redundancy check) calculation unit
The CRC (cyclic redundancy check) calculation unit is used to get a CRC code from a 32-bit 
data word and a fixed generator polynomial.
Among other applications, CRC-based techniques are used to verify data transmission or 
storage integrity. In the scope of the EN/IEC 60335-1 standard, they offer a means of 
verifying the Flash memory integrity. The CRC calculation unit helps compute a signature of 
the software during runtime, to be compared with a reference signature generated at link-
time and stored at a given memory location.
2.3.4 
Embedded SRAM
Up to 64 Kbytes of embedded SRAM accessed (read/write) at CPU clock speed with 0 wait 
states.
2.3.5 
FSMC (flexible static memory controller)
The FSMC is embedded in the STM32F103xC, STM32F103xD and STM32F103xE 
performance line family. It has four Chip Select outputs supporting the following modes: PC 
Card/Compact Flash, SRAM, PSRAM, NOR and NAND.
Functionality overview:
•
The three FSMC interrupt lines are ORed in order to be connected to the NVIC
•
Write FIFO
•
Code execution from external memory except for NAND Flash and PC Card
•
The targeted frequency, fCLK, is HCLK/2, so external access is at 36 MHz when HCLK 
is at 72 MHz and external access is at 24 MHz when HCLK is at 48 MHz

### Code Examples

```sql
used to get a CRC code from a 32-bit
```

```unknown
used to verify data transmission or
used external oscillator).
```

---

---

**📄 Source: PDF Page 16**

Description
STM32F103xC, STM32F103xD, STM32F103xE
16/144
DocID14611 Rev 12
2.3.6 
LCD parallel interface
The FSMC can be configured to interface seamlessly with most graphic LCD controllers. It 
supports the Intel 8080 and Motorola 6800 modes, and is flexible enough to adapt to 
specific LCD interfaces. This LCD parallel interface capability makes it easy to build cost-
effective graphic applications using LCD modules with embedded controllers or high-
performance solutions using external controllers with dedicated acceleration.
2.3.7 
Nested vectored interrupt controller (NVIC)
The STM32F103xC, STM32F103xD and STM32F103xE performance line embeds a nested 
vectored interrupt controller able to handle up to 60 maskable interrupt channels (not 
including the 16 interrupt lines of Cortex®-M3) and 16 priority levels. 
•
Closely coupled NVIC gives low latency interrupt processing
•
Interrupt entry vector table address passed directly to the core
•
Closely coupled NVIC core interface
•
Allows early processing of interrupts
•
Processing of late arriving higher priority interrupts
•
Support for tail-chaining
•
Processor state automatically saved
•
Interrupt entry restored on interrupt exit with no instruction overhead
This hardware block provides flexible interrupt management features with minimal interrupt 
latency.
2.3.8 
External interrupt/event controller (EXTI)
The external interrupt/event controller consists of 19 edge detector lines used to generate 
interrupt/event requests. Each line can be independently configured to select the trigger 
event (rising edge, falling edge, both) and can be masked independently. A pending register 
maintains the status of the interrupt requests. The EXTI can detect an external line with a 
pulse width shorter than the Internal APB2 clock period. Up to 112 GPIOs can be connected 
to the 16 external interrupt lines.
2.3.9 
Clocks and startup
System clock selection is performed on startup, however the internal RC 8 MHz oscillator is 
selected as default CPU clock on reset. An external 4-16 MHz clock can be selected, in 
which case it is monitored for failure. If failure is detected, the system automatically switches 
back to the internal RC oscillator. A software interrupt is generated if enabled. Similarly, full 
interrupt management of the PLL clock entry is available when necessary (for example with 
failure of an indirectly used external oscillator).
Several prescalers allow the configuration of the AHB frequency, the high speed APB 
(APB2) and the low speed APB (APB1) domains. The maximum frequency of the AHB and 
the high speed APB domains is 72 MHz. The maximum allowed frequency of the low speed 
APB domain is 36 MHz. See Figure 2 for details on the clock tree.

### Code Examples

```unknown
used to generate
```

---

---

**📄 Source: PDF Page 17**

DocID14611 Rev 12
17/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
2.3.10 
Boot modes
At startup, boot pins are used to select one of three boot options:
•
Boot from user Flash: you have an option to boot from any of two memory banks. By 
default, boot from Flash memory bank 1 is selected. You can choose to boot from Flash 
memory bank 2 by setting a bit in the option bytes.
•
Boot from system memory
•
Boot from embedded SRAM
The boot loader is located in system memory. It is used to reprogram the Flash memory by 
using USART1. 
2.3.11 
Power supply schemes
•
VDD = 2.0 to 3.6 V: external power supply for I/Os and the internal regulator. 
Provided externally through VDD pins.
•
VSSA, VDDA = 2.0 to 3.6 V: external analog power supplies for ADC, DAC, Reset 
blocks, RCs and PLL (minimum voltage to be applied to VDDA is 2.4 V when the ADC 
or DAC is used). VDDA and VSSA must be connected to VDD and VSS, respectively.
•
VBAT = 1.8 to 3.6 V: power supply for RTC, external clock 32 kHz oscillator and backup 
registers (through power switch) when VDD is not present.
For more details on how to connect power pins, refer to Figure 12: Power supply scheme.
2.3.12 
Power supply supervisor
The device has an integrated power-on reset (POR)/power-down reset (PDR) circuitry. It is 
always active, and ensures proper operation starting from/down to 2 V. The device remains 
in reset mode when VDD is below a specified threshold, VPOR/PDR, without the need for an 
external reset circuit.
The device features an embedded programmable voltage detector (PVD) that monitors the 
VDD/VDDA power supply and compares it to the VPVD threshold. An interrupt can be 
generated when VDD/VDDA drops below the VPVD threshold and/or when VDD/VDDA is 
higher than the VPVD threshold. The interrupt service routine can then generate a warning 
message and/or put the MCU into a safe state. The PVD is enabled by software. Refer to 
Table 12: Embedded reset and power control block characteristics for the values of 
VPOR/PDR and VPVD.
2.3.13 
Voltage regulator
The regulator has three operation modes: main (MR), low-power (LPR) and power down.
•
MR is used in the nominal regulation mode (Run) 
•
LPR is used in the Stop modes.
•
Power down is used in Standby mode: the regulator output is in high impedance: the 
kernel circuitry is powered down, inducing zero consumption (but the contents of the 
registers and SRAM are lost)
This regulator is always enabled after reset. It is disabled in Standby mode.

### Code Examples

```sql
used to select one of three boot options:
```

```sql
user Flash: you have an option to boot from any of two memory banks. By
```

```unknown
used to reprogram the Flash memory by
```

```unknown
used). VDDA and VSSA must be connected to VDD and VSS, respectively.
```

```unknown
used in the nominal regulation mode (Run)
```

```unknown
used in the Stop modes.
```

```unknown
used in Standby mode: the regulator output is in high impedance: the
```

---

---

**📄 Source: PDF Page 18**

Description
STM32F103xC, STM32F103xD, STM32F103xE
18/144
DocID14611 Rev 12
2.3.14 
Low-power modes
The STM32F103xC, STM32F103xD and STM32F103xE performance line supports three 
low-power modes to achieve the best compromise between low-power consumption, short 
startup time and available wakeup sources:
•
Sleep mode
In Sleep mode, only the CPU is stopped. All peripherals continue to operate and can 
wake up the CPU when an interrupt/event occurs.
•
Stop mode
Stop mode achieves the lowest power consumption while retaining the content of 
SRAM and registers. All clocks in the 1.8 V domain are stopped, the PLL, the HSI RC 
and the HSE crystal oscillators are disabled. The voltage regulator can also be put 
either in normal or in low-power mode. 
The device can be woken up from Stop mode by any of the EXTI line. The EXTI line 
source can be one of the 16 external lines, the PVD output, the RTC alarm or the USB 
wakeup. 
•
Standby mode
The Standby mode is used to achieve the lowest power consumption. The internal 
voltage regulator is switched off so that the entire 1.8 V domain is powered off. The 
PLL, the HSI RC and the HSE crystal oscillators are also switched off. After entering 
Standby mode, SRAM and register contents are lost except for registers in the Backup 
domain and Standby circuitry. 
The device exits Standby mode when an external reset (NRST pin), an IWDG reset, a 
rising edge on the WKUP pin, or an RTC alarm occurs.
Note:
The RTC, the IWDG, and the corresponding clock sources are not stopped by entering Stop 
or Standby mode.
2.3.15 
DMA
The flexible 12-channel general-purpose DMAs (7 channels for DMA1 and 5 channels for 
DMA2) are able to manage memory-to-memory, peripheral-to-memory and memory-to-
peripheral transfers. The two DMA controllers support circular buffer management, 
removing the need for user code intervention when the controller reaches the end of the 
buffer.
Each channel is connected to dedicated hardware DMA requests, with support for software 
trigger on each channel. Configuration is made by software and transfer sizes between 
source and destination are independent.
The DMA can be used with the main peripherals: SPI, I2C, USART, general-purpose, basic 
and advanced-control timers TIMx, DAC, I2S, SDIO and ADC.
2.3.16 
RTC (real-time clock) and backup registers
The RTC and the backup registers are supplied through a switch that takes power either on 
VDD supply when present or through the VBAT pin. The backup registers are forty-two 16-bit 
registers used to store 84 bytes of user application data when VDD power is not present. 
They are not reset by a system or power reset, and they are not reset when the device 
wakes up from the Standby mode.
The real-time clock provides a set of continuously running counters which can be used with 
suitable software to provide a clock calendar function, and provides an alarm interrupt and a

### Code Examples

```julia
user code intervention when the controller reaches the end of the
```

```unknown
used to achieve the lowest power consumption. The internal
```

```unknown
used with the main peripherals: SPI, I2C, USART, general-purpose, basic
```

```unknown
used to store 84 bytes of user application data when VDD power is not present.
used for the time base clock and is by default
```

---

---

**📄 Source: PDF Page 19**

DocID14611 Rev 12
19/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
periodic interrupt. It is clocked by a 32.768 kHz external crystal, resonator or oscillator, the 
internal low-power RC oscillator or the high-speed external clock divided by 128. The 
internal low-speed RC has a typical frequency of 40 kHz. The RTC can be calibrated using 
an external 512 Hz output to compensate for any natural quartz deviation. The RTC features 
a 32-bit programmable counter for long term measurement using the Compare register to 
generate an alarm. A 20-bit prescaler is used for the time base clock and is by default 
configured to generate a time base of 1 second from a clock at 32.768 kHz. 
2.3.17 
Timers and watchdogs
The high-density STM32F103xC/D/E performance line devices include up to two advanced-
control timers, up to four general-purpose timers, two basic timers, two watchdog timers and 
a SysTick timer.
Table 4 compares the features of the advanced-control, general-purpose and basic timers.
         
Table 4. High-density timer feature comparison
Timer
Counter 
resolution
Counter 
type
Prescaler 
factor
DMA request 
generation
Capture/compare 
channels
Complementary
outputs
TIM1, 
TIM8
16-bit
Up, 
down, 
up/down
Any integer 
between 1 
and 65536
Yes
4
Yes
TIM2, 
TIM3, 
TIM4, 
TIM5
16-bit
Up, 
down, 
up/down
Any integer 
between 1 
and 65536
Yes
4
No
TIM6, 
TIM7
16-bit
Up
Any integer 
between 1 
and 65536
Yes
0
No

### Code Examples

```unknown
include up to two advanced-
```

---

---

**📄 Source: PDF Page 20**

Description
STM32F103xC, STM32F103xD, STM32F103xE
20/144
DocID14611 Rev 12
Advanced-control timers (TIM1 and TIM8)
The two advanced-control timers (TIM1 and TIM8) can each be seen as a three-phase 
PWM multiplexed on 6 channels. They have complementary PWM outputs with 
programmable inserted dead-times. They can also be seen as a complete general-purpose 
timer. The 4 independent channels can be used for:
•
Input capture
•
Output compare
•
PWM generation (edge or center-aligned modes)
•
One-pulse mode output
If configured as a standard 16-bit timer, it has the same features as the TIMx timer. If 
configured as the 16-bit PWM generator, it has full modulation capability (0-100%). 
In debug mode, the advanced-control timer counter can be frozen and the PWM outputs 
disabled to turn off any power switch driven by these outputs.
Many features are shared with those of the general-purpose TIM timers which have the 
same architecture. The advanced-control timer can therefore work together with the TIM 
timers via the Timer Link feature for synchronization or event chaining. 
General-purpose timers (TIMx)
There are up to 4 synchronizable general-purpose timers (TIM2, TIM3, TIM4 and TIM5) 
embedded in the STM32F103xC, STM32F103xD and STM32F103xE performance line 
devices. These timers are based on a 16-bit auto-reload up/down counter, a 16-bit prescaler 
and feature 4 independent channels each for input capture/output compare, PWM or one-
pulse mode output. This gives up to 16 input captures / output compares / PWMs on the 
largest packages.
The general-purpose timers can work together with the advanced-control timer via the Timer 
Link feature for synchronization or event chaining. Their counter can be frozen in debug 
mode. Any of the general-purpose timers can be used to generate PWM outputs. They all 
have independent DMA request generation.
These timers are capable of handling quadrature (incremental) encoder signals and the 
digital outputs from 1 to 3 hall-effect sensors.
Basic timers TIM6 and TIM7
These timers are mainly used for DAC trigger generation. They can also be used as a 
generic 16-bit time base.
Independent watchdog
The independent watchdog is based on a 12-bit downcounter and 8-bit prescaler. It is 
clocked from an independent 40 kHz internal RC and as it operates independently from the 
main clock, it can operate in Stop and Standby modes. It can be used either as a watchdog 
to reset the device when a problem occurs, or as a free running timer for application timeout 
management. It is hardware or software configurable through the option bytes. The counter 
can be frozen in debug mode.
Window watchdog
The window watchdog is based on a 7-bit downcounter that can be set as free running. It 
can be used as a watchdog to reset the device when a problem occurs. It is clocked from

### Code Examples

```typescript
used for DAC trigger generation. They can also be used as a
```

```typescript
used either as a watchdog
```

```typescript
used as a watchdog to reset the device when a problem occurs. It is clocked from
```

```unknown
used to generate PWM outputs. They all
```

---

---

**📄 Source: PDF Page 21**

DocID14611 Rev 12
21/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
the main clock. It has an early warning interrupt capability and the counter can be frozen in 
debug mode.
SysTick timer
This timer is dedicated to real-time operating systems, but could also be used as a standard 
down counter. It features:
•
A 24-bit down counter
•
Autoreload capability
•
Maskable system interrupt generation when the counter reaches 0.
•
Programmable clock source 
2.3.18 
I²C bus
Up to two I²C bus interfaces can operate in multimaster and slave modes. They can support 
standard and fast modes. 
They support 7/10-bit addressing mode and 7-bit dual addressing mode (as slave). A 
hardware CRC generation/verification is embedded.
They can be served by DMA and they support SMBus 2.0/PMBus.
2.3.19 
Universal synchronous/asynchronous receiver transmitters (USARTs)
The STM32F103xC, STM32F103xD and STM32F103xE performance line embeds three 
universal synchronous/asynchronous receiver transmitters (USART1, USART2 and 
USART3) and two universal asynchronous receiver transmitters (UART4 and UART5).
These five interfaces provide asynchronous communication, IrDA SIR ENDEC support, 
multiprocessor communication mode, single-wire half-duplex communication mode and 
have LIN Master/Slave capability.
The USART1 interface is able to communicate at speeds of up to 4.5 Mbit/s. The other 
available interfaces communicate at up to 2.25 Mbit/s.
USART1, USART2 and USART3 also provide hardware management of the CTS and RTS 
signals, Smart Card mode (ISO 7816 compliant) and SPI-like communication capability. All 
interfaces can be served by the DMA controller except for UART5.
2.3.20 
Serial peripheral interface (SPI)
Up to three SPIs are able to communicate up to 18 Mbits/s in slave and master modes in 
full-duplex and simplex communication modes. The 3-bit prescaler gives 8 master mode 
frequencies and the frame is configurable to 8 bits or 16 bits. The hardware CRC 
generation/verification supports basic SD Card/MMC modes. 
All SPIs can be served by the DMA controller.
2.3.21 
Inter-integrated sound (I2S)
Two standard I2S interfaces (multiplexed with SPI2 and SPI3) are available, that can be 
operated in master or slave mode. These interfaces can be configured to operate with 16/32 
bit resolution, as input or output channels. Audio sampling frequencies from 8 kHz up to 
48 kHz are supported. When either or both of the I2S interfaces is/are configured in master

### Code Examples

```typescript
used as a standard
```

---

---

**📄 Source: PDF Page 22**

Description
STM32F103xC, STM32F103xD, STM32F103xE
22/144
DocID14611 Rev 12
mode, the master clock can be output to the external DAC/CODEC at 256 times the 
sampling frequency.
2.3.22 
SDIO
An SD/SDIO/MMC host interface is available, that supports MultiMediaCard System 
Specification Version 4.2 in three different databus modes: 1-bit (default), 4-bit and 8-bit.
The interface allows data transfer at up to 48 MHz in 8-bit mode, and is compliant with SD 
Memory Card Specifications Version 2.0.
The SDIO Card Specification Version 2.0 is also supported with two different databus 
modes: 1-bit (default) and 4-bit.
The current version supports only one SD/SDIO/MMC4.2 card at any one time and a stack 
of MMC4.1 or previous.
In addition to SD/SDIO/MMC, this interface is also fully compliant with the CE-ATA digital 
protocol Rev1.1.
2.3.23 
Controller area network (CAN)
The CAN is compliant with specifications 2.0A and B (active) with a bit rate up to 1 Mbit/s. It 
can receive and transmit standard frames with 11-bit identifiers as well as extended frames 
with 29-bit identifiers. It has three transmit mailboxes, two receive FIFOs with 3 stages and 
14 scalable filter banks.
2.3.24 
Universal serial bus (USB)
The STM32F103xC, STM32F103xD and STM32F103xE performance line embed a USB 
device peripheral compatible with the USB full-speed 12 Mbs. The USB interface 
implements a full-speed (12 Mbit/s) function interface. It has software-configurable endpoint 
setting and suspend/resume support. The dedicated 48 MHz clock is generated from the 
internal main PLL (the clock source must use a HSE crystal oscillator).
2.3.25 
GPIOs (general-purpose inputs/outputs)
Each of the GPIO pins can be configured by software as output (push-pull or open-drain), as 
input (with or without pull-up or pull-down) or as peripheral alternate function. Most of the 
GPIO pins are shared with digital or analog alternate functions. All GPIOs are high current-
capable.
The I/Os alternate function configuration can be locked if needed following a specific 
sequence in order to avoid spurious writing to the I/Os registers.
2.3.26 
ADC (analog to digital converter)
Three 12-bit analog-to-digital converters are embedded into STM32F103xC, STM32F103xD 
and STM32F103xE performance line devices and each ADC shares up to 21 external 
channels, performing conversions in single-shot or scan modes. In scan mode, automatic 
conversion is performed on a selected group of analog inputs.
Additional logic functions embedded in the ADC interface allow:
•
Simultaneous sample and hold
•
Interleaved sample and hold
•
Single shunt

### Code Examples

```elixir
use a HSE crystal oscillator).
```

---

---

**📄 Source: PDF Page 23**

DocID14611 Rev 12
23/144
STM32F103xC, STM32F103xD, STM32F103xE
Description
136
The ADC can be served by the DMA controller.
An analog watchdog feature allows very precise monitoring of the converted voltage of one, 
some or all selected channels. An interrupt is generated when the converted voltage is 
outside the programmed thresholds.
The events generated by the general-purpose timers (TIMx) and the advanced-control 
timers (TIM1 and TIM8) can be internally connected to the ADC start trigger and injection 
trigger, respectively, to allow the application to synchronize A/D conversion and timers.
2.3.27 
DAC (digital-to-analog converter)
The two 12-bit buffered DAC channels can be used to convert two digital signals into two 
analog voltage signal outputs. The chosen design structure is composed of integrated 
resistor strings and an amplifier in inverting configuration.
This dual digital Interface supports the following features:
•
two DAC converters: one for each output channel
•
8-bit or 12-bit monotonic output
•
left or right data alignment in 12-bit mode
•
synchronized update capability
•
noise-wave generation
•
triangular-wave generation
•
dual DAC channel independent or simultaneous conversions
•
DMA capability for each channel
•
external triggers for conversion
•
input voltage reference VREF+
Eight DAC trigger inputs are used in the STM32F103xC, STM32F103xD and 
STM32F103xE performance line family. The DAC channels are triggered through the timer 
update outputs that are also connected to different DMA channels.

### Code Examples

```unknown
used to convert two digital signals into two
```

```unknown
used in the STM32F103xC, STM32F103xD and
used to convert the sensor output
```

---

---

**📄 Source: PDF Page 24**

Description
STM32F103xC, STM32F103xD, STM32F103xE
24/144
DocID14611 Rev 12
2.3.28 
Temperature sensor
The temperature sensor has to generate a voltage that varies linearly with temperature. The 
conversion range is between 2 V < VDDA < 3.6 V. The temperature sensor is internally 
connected to the ADC1_IN16 input channel which is used to convert the sensor output 
voltage into a digital value.
2.3.29 
Serial wire JTAG debug port (SWJ-DP)
The ARM SWJ-DP Interface is embedded, and is a combined JTAG and serial wire debug 
port that enables either a serial wire debug or a JTAG probe to be connected to the target. 
The JTAG TMS and TCK pins are shared respectively with SWDIO and SWCLK and a 
specific sequence on the TMS pin is used to switch between JTAG-DP and SW-DP.
2.3.30 
Embedded Trace Macrocell™
The ARM® Embedded Trace Macrocell provides a greater visibility of the instruction and 
data flow inside the CPU core by streaming compressed data at a very high rate from the 
STM32F10xxx through a small number of ETM pins to an external hardware trace port 
analyzer (TPA) device. The TPA is connected to a host computer using USB, Ethernet, or 
any other high-speed channel. Real-time instruction and data flow activity can be recorded 
and then formatted for display on the host computer running debugger software. TPA 
hardware is commercially available from common development tool vendors. It operates 
with third party debugger software tools.

### Code Examples

```unknown
used to switch between JTAG-DP and SW-DP.
```

---

---

**📄 Source: PDF Page 25**

DocID14611 Rev 12
25/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
3 
Pinouts and pin descriptions
Figure 3. STM32F103xC/D/E BGA144 ballout
1.
The above figure shows the package top view. 
AI14798b
VDD_7
PC3
PC2
PF6
VDD_6
VSS_4
PF8
H
VDD_1
D
PG13
PG14
PE6
PE5
C
PG10
PG11
VDD_5
PB8
NRST
B
PG12
PG15
PC15-
OSC32_OUT
PB9
A
8
7
6
5
4
3
2
1
VBAT
OSC_IN
OSC_OUT
VSS_5
G
F
E
PF7
PC0
PF0
PF1
PF2
VSS_10
PG9
PF4
PF3
VSS_3
PF5
VDD_8
VDD_3
VDD_4
VSS_8
PE4
PB5
PB6
BOOT0
PB7
VSS_11
PF10
PC1
VDD_11
VDD_10
PF9
10
9
K
J
VSS_2
PD3
PD4
PD1
PC12
PC11
PD5
PD2
PD0
VDD_9
VSS_9
VDD_2
PG1
PC5
PA5
PE9
PB2/
BOOT1
PC4
PA4
PE10
PG0
PF13
VREF–
PE12
VSSA
PA1
PE13
PA0-WKUP
PD9
PD10
PG4
PD13
12
11
PG8
PA10
NC
PA9
PA11
PA12
PC10
PC9
PA8
PC7
PC6
PC8
PD14
PG3
PG2
PD15
M
L
PF15
PB1
PA7
PE7
PF12
PB0
PA6
PE8
PF14
PF11
VDDA
PE14
VREF+
PA3
PE15
PA2
PB10
PD8
PD12
PB11
PB12
PB14
PB15
PB13
PC13-
TAMPER-RTC
PE3
PE2
PE1
PE0
PB4
JTRST
PB3
JTDO
PD6
PD7
PA15
JTDI
PA14
JTCK
PA13
JTMS
PE11
VSS_6
VSS_7
VSS_1
PG7
PD11
PG5
PG6
PC14-
OSC32_IN

---

---

**📄 Source: PDF Page 26**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
26/144
DocID14611 Rev 12
Figure 4. STM32F103xC/D/E performance line BGA100 ballout
1.
The above figure shows the package top view.
AI14601c
PE10
PC14-
OSC32_IN
PC5
PA5
PC3
PB4
PE15
PB2
PC4
PA4
H
PE14
PE11
PE7
D
PD4
PD3
PB8
PE3
C
PD0
PC12
PE5
PB5
PC0
PE2
B
PC11
PD2
PC15-
OSC32_OUT
PB7
PB6
A
8
7
6
5
4
3
2
1
VSS_5
OSC_IN
OSC_OUT
VDD_5
G
F
E
PC1
VREF–
PC13-
TAMPER-
RTC
PB9
PA15
PB3
PE4
PE1
PE0
VSS_1
PD1
PE6
NRST
PC2
VSS_3
VSS_4
NC
VDD_3
VDD_4
PB15
VBAT
PD5
PD6
BOOT0
PD7
VSS_2
VSSA
PA1
VDD_2
VDD_1
PB14
PA0-WKUP
10
9
K
J
PD10
PD11
PA8
PA9
PA10
PA11
PA12
PC10
PA13
PA14
PC9
PC7
PC6
PD15
PC8
PD14
PE12
PB1
PA7
PB11
PE8
PB0
PA6
PB10
PE13
PE9
VDDA
PB13
VREF+
PA3
PB12
PA2
PD8
PD9
PD13
PD12

---

---

**📄 Source: PDF Page 27**

DocID14611 Rev 12
27/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
Figure 5. STM32F103xC/D/E performance line LQFP144 pinout
1.
The above figure shows the package top view. 
6$$?
633?
0%
0%
0"
0"
"//4
0"
0"
0"
0"
0"
0'
6$$?
633?
0'
0'
0'
0'
0'
0'
0$
0$
6$$?
633?
0$
0$
0$
0$
0$
0$
0#
0#
0#
0!
0!
0%
6$$?
0%
633?
0%
.#
0%
0!
0%
0!
6"!4
0!
0#4!-0%224#
0!
0#/3#?).
0!
0#/3#?/54
0!
0&
0#
0&
0#
0&
0#
0&
0#
0&
6$$?
0&
633?
633?
0'
6$$?
0'
0&
0'
0&
0'
0&
0'
0&
0'
0&
0'
/3#?).
0$
/3#?/54
0$
.234
6$$?
0#
633?
0#
0$
0#
0$
0#
0$
633!
0$
62%&
0$
62%&
0$
6$$!
0"
0!7+50
0"
0!
0"
0!
0"
0!
633?
6$$?
0!
0!
0!
0!
0#
0#
0"
0"
0"
0&
0&
633?
6$$?
0&
0&
0&
0'
0'
0%
0%
0%
633?
6$$?
0%
0%
0%
0%
0%
0%
0"
0"
633?
6$$?




































































































,1&0












































AI

---

---

**📄 Source: PDF Page 28**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
28/144
DocID14611 Rev 12
Figure 6. STM32F103xC/D/E performance line LQFP100 pinout
1.
The above figure shows the package top view. 
AI











































































6$$? 633?  .#  0!   0!   0!   0!   0!   0!   0#  0#  0#  0#  0$  0$  0$  0$  0$  0$  0$  0$  0"  0"  0"  0"  0!
633?
6$$?
0!
0!
0!
0!
0#
0#
0"
0"
0"
0%
0%
0%
0%
0%
0%
0%
0%
0%
0"
0"
633?
6$$?
6$$?  633?  0%  0%  0"  0"  "//4  0"  0"  0"  0"  0"  0$  0$  0$  0$  0$  0$  0$  0$  0#  0#  0#  0!  0! 
























0%
0%
0%
0%
0%
6"!4
0#4!-0%224#
0#/3#?).
0#/3#?/54
633?
6$$?
/3#?).
/3#?/54
.234
0#
0#
0#
0#
633!
62%&
62%&
6$$!
0!7+50
0!
0!
,1&0

---

---

**📄 Source: PDF Page 29**

DocID14611 Rev 12
29/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
Figure 7. STM32F103xC/D/E performance line LQFP64 pinout
1.
The above figure shows the package top view. 
ϲϰ ϲϯ ϲϮ ϲϭ ϲϬ ϱϵ ϱϴ ϱϳ ϱϲ ϱϱ ϱϰ ϱϯ ϱϮ ϱϭ ϱϬ ϰϵ
ϰϴ
ϰϳ
ϰϲ
ϰϱ
ϰϰ
ϰϯ
ϰϮ
ϰϭ
ϰϬ
ϯϵ
ϯϴ
ϯϳ
ϯϲ
ϯϱ
ϯϰ
ϯϯ
ϭϳ ϭϴ ϭϵ ϮϬ Ϯϭ ϮϮ Ϯϯ Ϯϰ
Ϯϵ ϯϬ ϯϭ ϯϮ
Ϯϱ Ϯϲ Ϯϳ Ϯϴ
ϭ
Ϯ
ϯ
ϰ
ϱ
ϲ
ϳ
ϴ
ϵ
ϭϬ
ϭϭ
ϭϮ
ϭϯ
ϭϰ
ϭϱ
ϭϲ
sd
WϭϯͲdDWZͲZd
W ϭϰͲK^ ϯϮͺ/E
W ϭϱͲK^ ϯϮͺKh d
WϬͲK^ ͺ/E
WϭͲK^ ͺKhd
EZ^d
WϬ
Wϭ
WϮ
Wϯ
s^^
s
W ϬͲt< hW
W ϭ
W Ϯ
sͺϯ
s^^ͺϯ
W ϵ
W ϴ
KK dϬ
W ϳ
W ϲ
W ϱ
W ϰ
W ϯ
WϮ
WϭϮ
Wϭϭ
WϭϬ
W ϭϱ
W ϭϰ
sͺϮ
s^^ͺϮ
W ϭϯ
W ϭϮ
W ϭϭ
W ϭϬ
W ϵ
W ϴ
Wϵ
Wϴ
Wϳ
Wϲ
W ϭϱ
W ϭϰ
W ϭϯ
W ϭϮ
W ϯ
s^^ͺϰ
sͺϰ
W ϰ
W ϱ
W ϲ
W ϳ
Wϰ
Wϱ
W Ϭ
W ϭ
W Ϯ
Wϭ Ϭ
Wϭ ϭ
s^^ͺϭ
sͺϭ
>Y&Wϲϰ
DL

---

---

**📄 Source: PDF Page 30**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
30/144
DocID14611 Rev 12
Figure 8. STM32F103xC/D/E performance line
WLCSP64 ballout, ball side
ai15460b
H
D
C
B
A
8
7
6
5
4
3
2
1
G
F
E
VDD_2
PC10
PD2
PB5
PB3
BOOT0
VSS_3
VDD_3
BYPASS/
VSS_2
PA14
PC11
PB4
PB6
PB9
PC15
PC14
PC13
NRST
VBAT
PB7
PC12
PA15
PA12
PA11
OSC_IN
OSC_OUT
PC2
PB8
PA13
PA10
PA9
PC9
PC0
VSSA
PA1
PA5
PA8
PC8
PC7
PC6
PC1
VREF+
PA0-
WKUP
VSS_4
PB1
PB11
PB14
PB15
VDDA
PA3
VDD_4
PA6
PA7
PB10
PB12
PB13
PA2
PA4
PC4
PC5
PB0
PB2
VSS_1
VDD_1

---

---

**📄 Source: PDF Page 31**

DocID14611 Rev 12
31/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
         
Table 5. High-density STM32F103xC/D/E pin definitions
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap
A3
A3
-
-
1
1
PE2
I/O FT
PE2
TRACECK/ FSMC_A23
-
A2
B3
-
-
2
2
PE3
I/O FT
PE3
TRACED0/FSMC_A19
-
B2
C3
-
-
3
3
PE4
I/O FT
PE4
TRACED1/FSMC_A20
-
B3
D3
-
-
4
4
PE5
I/O FT
PE5
TRACED2/FSMC_A21
-
B4
E3
-
-
5
5
PE6
I/O FT
PE6
TRACED3/FSMC_A22
-
C2
B2
C6
1
6
6
VBAT
S
-
VBAT
-
-
A1
A2
C8
2
7
7
PC13-TAMPER-
RTC(5)
I/O
-
PC13(6)
TAMPER-RTC
-
B1
A1
B8
3
8
8
PC14-
OSC32_IN(5)
I/O
-
PC14(6)
OSC32_IN
-
C1
B1
B7
4
9
9
PC15-
OSC32_OUT(5) I/O
-
PC15(6)
OSC32_OUT
-
C3
-
-
-
-
10
PF0
I/O FT
PF0
FSMC_A0
-
C4
-
-
-
-
11
PF1
I/O FT
PF1
FSMC_A1
-
D4
-
-
-
-
12
PF2
I/O FT
PF2
FSMC_A2
-
E2
-
-
-
-
13
PF3
I/O FT
PF3
FSMC_A3
-
E3
-
-
-
-
14
PF4
I/O FT
PF4
FSMC_A4
-
E4
-
-
-
-
15
PF5
I/O FT
PF5
FSMC_A5
-
D2
C2
-
-
10
16
VSS_5
S
-
VSS_5
-
-
D3
D2
-
-
11
17
VDD_5
S
-
VDD_5
-
-
F3
-
-
-
-
18
PF6
I/O
-
PF6
ADC3_IN4/FSMC_NIORD
-
F2
-
-
-
-
19
PF7
I/O
-
PF7
ADC3_IN5/FSMC_NREG
-
G3
-
-
-
-
20
PF8
I/O
-
PF8
ADC3_IN6/FSMC_NIOWR
-
G2
-
-
-
-
21
PF9
I/O
-
PF9
ADC3_IN7/FSMC_CD
-
G1
-
-
-
-
22
PF10
I/O
-
PF10
ADC3_IN8/FSMC_INTR
-
D1
C1
D8
5
12
23
OSC_IN
I
-
OSC_IN
-
-
E1
D1
D7
6
13
24
OSC_OUT
O
-
OSC_OUT
-
-
F1
E1
C7
7
14
25
NRST
I/O
-
NRST
-
-
H1
F1
E8
8
15
26
PC0
I/O
-
PC0
ADC123_IN10
-
H2
F2
F8
9
16
27
PC1
I/O
-
PC1
ADC123_IN11
-

---

---

**📄 Source: PDF Page 32**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
32/144
DocID14611 Rev 12
H3
E2
D6
10
17
28
PC2
I/O
-
PC2
ADC123_IN12
-
H4
F3
-
11
18
29
PC3(7)
I/O
-
PC3
ADC123_IN13
-
J1
G1
E7
12
19
30
VSSA
S
-
VSSA
-
-
K1
H1
-
-
20
31
VREF-
S
-
VREF-
-
-
L1
J1
F7
(8)
-
21
32
VREF+
S
-
VREF+
-
-
M1
K1
G8 13
22
33
VDDA
S
-
VDDA
-
-
J2
G2
F6
14
23
34
PA0-WKUP
I/O
-
PA0
WKUP/USART2_CTS(9)
ADC123_IN0
TIM2_CH1_ETR
TIM5_CH1/TIM8_ETR
-
K2
H2
E6
15
24
35
PA1
I/O
-
PA1
USART2_RTS(9)
ADC123_IN1/
TIM5_CH2/TIM2_CH2(9)
-
L2
J2
H8
16
25
36
PA2
I/O
-
PA2
USART2_TX(9)/TIM5_CH3
ADC123_IN2/ 
TIM2_CH3 (9)
-
M2
K2
G7 17
26
37
PA3
I/O
-
PA3
USART2_RX(9)/TIM5_CH4
ADC123_IN3/TIM2_CH4(9) 
-
G4
E4
F5
18
27
38
VSS_4
S
-
VSS_4
-
-
F4
F4
G6 19
28
39
VDD_4
S
-
VDD_4
-
-
J3
G3
H7
20
29
40
PA4
I/O
-
PA4
SPI1_NSS(9)/
USART2_CK(9) 
DAC_OUT1/ADC12_IN4 
-
K3
H3
E5
21
30
41
PA5
I/O
-
PA5
SPI1_SCK(9)
DAC_OUT2 ADC12_IN5 
-
L3
J3
G5 22
31
42
PA6
I/O
-
PA6
SPI1_MISO(9)
TIM8_BKIN/ADC12_IN6
TIM3_CH1(9) 
TIM1_BKIN
M3
K3
G4 23
32
43
PA7
I/O
-
PA7
SPI1_MOSI(9)/
TIM8_CH1N/ADC12_IN7
TIM3_CH2(9) 
TIM1_CH1N
J4
G4
H6
24
33
44
PC4
I/O
-
PC4
ADC12_IN14 
-
K4
H4
H5
25
34
45
PC5
I/O
-
PC5
ADC12_IN15 
-
Table 5. High-density STM32F103xC/D/E pin definitions (continued)
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap

---

---

**📄 Source: PDF Page 33**

DocID14611 Rev 12
33/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
L4
J4
H4
26
35
46
PB0
I/O
-
PB0
ADC12_IN8/TIM3_CH3
TIM8_CH2N
TIM1_CH2N
M4
K4
F4
27
36
47
PB1
I/O
-
PB1
ADC12_IN9/TIM3_CH4(9)
TIM8_CH3N
TIM1_CH3N
J5
G5
H3
28
37
48
PB2
I/O FT PB2/BOOT1
-
-
M5
-
-
-
-
49
PF11
I/O FT
PF11
FSMC_NIOS16
-
L5
-
-
-
-
50
PF12
I/O FT
PF12
FSMC_A6
-
H5
-
-
-
-
51
VSS_6
S
-
VSS_6
-
-
G5
-
-
-
-
52
VDD_6
S
-
VDD_6
-
-
K5
-
-
-
-
53
PF13
I/O FT
PF13
FSMC_A7
-
M6
-
-
-
-
54
PF14
I/O FT
PF14
FSMC_A8
-
L6
-
-
-
-
55
PF15
I/O FT
PF15
FSMC_A9
-
K6
-
-
-
-
56
PG0
I/O FT
PG0
FSMC_A10
-
J6
-
-
-
-
57
PG1
I/O FT
PG1
FSMC_A11
-
M7
H5
-
-
38
58
PE7
I/O FT
PE7
FSMC_D4
TIM1_ETR
L7
J5
-
-
39
59
PE8
I/O FT
PE8
FSMC_D5
TIM1_CH1N
K7
K5
-
-
40
60
PE9
I/O FT
PE9
FSMC_D6
TIM1_CH1
H6
-
-
-
-
61
VSS_7
S
-
VSS_7
-
-
G6
-
-
-
-
62
VDD_7
S
-
VDD_7
-
-
J7
G6
-
-
41
63
PE10
I/O FT
PE10
FSMC_D7
TIM1_CH2N
H8
H6
-
-
42
64
PE11
I/O FT
PE11
FSMC_D8
TIM1_CH2
J8
J6
-
-
43
65
PE12
I/O FT
PE12
FSMC_D9
TIM1_CH3N
K8
K6
-
-
44
66
PE13
I/O FT
PE13
FSMC_D10
TIM1_CH3
L8
G7
-
-
45
67
PE14
I/O FT
PE14
FSMC_D11
TIM1_CH4
M8
H7
-
-
46
68
PE15
I/O FT
PE15
FSMC_D12
TIM1_BKIN
M9
J7
G3 29
47
69
PB10
I/O FT
PB10
I2C2_SCL/USART3_TX(9)
TIM2_CH3
M10 K7
F3
30
48
70
PB11
I/O FT
PB11
I2C2_SDA/USART3_RX(9)
TIM2_CH4
H7
E7
H2
31
49
71
VSS_1
S
-
VSS_1
-
-
G7
F7
H1
32
50
72
VDD_1
S
-
VDD_1
-
-
Table 5. High-density STM32F103xC/D/E pin definitions (continued)
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap

---

---

**📄 Source: PDF Page 34**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
34/144
DocID14611 Rev 12
M11 K8
G2 33
51
73
PB12
I/O FT
PB12
SPI2_NSS/I2S2_WS/
I2C2_SMBA/
USART3_CK(9)/
TIM1_BKIN(9)
-
M12 J8
G1 34
52
74
PB13
I/O FT
PB13
SPI2_SCK/I2S2_CK
USART3_CTS(9)/
TIM1_CH1N 
-
L11 H8
F2
35
53
75
PB14
I/O FT
PB14
SPI2_MISO/TIM1_CH2N
USART3_RTS(9)/ 
-
L12 G8
F1
36
54
76
PB15
I/O FT
PB15
SPI2_MOSI/I2S2_SD
TIM1_CH3N(9)/ 
-
L9
K9
-
-
55
77
PD8
I/O FT
PD8
FSMC_D13
USART3_TX
K9
J9
-
-
56
78
PD9
I/O FT
PD9
FSMC_D14
USART3_RX
J9
H9
-
-
57
79
PD10
I/O FT
PD10
FSMC_D15
USART3_CK
H9
G9
-
-
58
80
PD11
I/O FT
PD11
FSMC_A16
USART3_CTS
L10 K10
-
-
59
81
PD12
I/O FT
PD12
FSMC_A17
TIM4_CH1 / 
USART3_RTS
K10 J10
-
-
60
82
PD13
I/O FT
PD13
FSMC_A18
TIM4_CH2
G8
-
-
-
-
83
VSS_8
S
-
VSS_8
-
-
F8
-
-
-
-
84
VDD_8
S
-
VDD_8
-
-
K11 H10
-
-
61
85
PD14
I/O FT
PD14
FSMC_D0
TIM4_CH3
K12 G10
-
-
62
86
PD15
I/O FT
PD15
FSMC_D1
TIM4_CH4
J12
-
-
-
-
87
PG2
I/O FT
PG2
FSMC_A12
-
J11
-
-
-
-
88
PG3
I/O FT
PG3
FSMC_A13
-
J10
-
-
-
-
89
PG4
I/O FT
PG4
FSMC_A14
-
H12
-
-
-
-
90
PG5
I/O FT
PG5
FSMC_A15
-
H11
-
-
-
-
91
PG6
I/O FT
PG6
FSMC_INT2
-
H10
-
-
-
-
92
PG7
I/O FT
PG7
FSMC_INT3
-
G11
-
-
-
-
93
PG8
I/O FT
PG8
-
-
G10
-
-
-
-
94
VSS_9
S
-
VSS_9
-
-
F10
-
-
-
-
95
VDD_9
S
-
VDD_9
-
-
Table 5. High-density STM32F103xC/D/E pin definitions (continued)
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap

---

---

**📄 Source: PDF Page 35**

DocID14611 Rev 12
35/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
G12 F10 E1
37
63
96
PC6
I/O FT
PC6
I2S2_MCK/
TIM8_CH1/SDIO_D6
TIM3_CH1
F12 E10 E2
38
64
97
PC7
I/O FT
PC7
I2S3_MCK/
TIM8_CH2/SDIO_D7
TIM3_CH2
F11
F9
E3
39
65
98
PC8
I/O FT
PC8
TIM8_CH3/SDIO_D0
TIM3_CH3
E11 E9
D1
40
66
99
PC9
I/O FT
PC9
TIM8_CH4/SDIO_D1
TIM3_CH4
E12 D9
E4
41
67 100
PA8
I/O FT
PA8
USART1_CK/
TIM1_CH1(9)/MCO 
-
D12 C9
D2
42
68 101
PA9
I/O FT
PA9
USART1_TX(9)/
TIM1_CH2(9)
-
D11 D10 D3
43
69 102
PA10
I/O FT
PA10
USART1_RX(9)/
TIM1_CH3(9)
-
C12 C10 C1
44
70 103
PA11
I/O FT
PA11
USART1_CTS/USBDM
CAN_RX(9)/TIM1_CH4(9) 
-
B12 B10 C2
45
71 104
PA12
I/O FT
PA12
USART1_RTS/USBDP/
CAN_TX(9)/TIM1_ETR(9)
-
A12 A10 D4
46
72 105
PA13
I/O FT
JTMS-
SWDIO
-
PA13
C11 F8
-
-
73 106
Not connected
-
G9
E6
B1
47
74 107
VSS_2
S
-
VSS_2
-
-
F9
F6
A1
48
75 108
VDD_2
S
-
VDD_2
-
-
A11 A9
B2
49
76 109
PA14
I/O FT
JTCK-
SWCLK
-
PA14
A10 A8
C3
50
77 110
PA15
I/O FT
JTDI
SPI3_NSS/
I2S3_WS
TIM2_CH1_ETR 
PA15 / SPI1_NSS
B11 B9
A2
51
78 111
PC10
I/O FT
PC10
UART4_TX/SDIO_D2
USART3_TX
B10 B8
B3
52
79 112
PC11
I/O FT
PC11
UART4_RX/SDIO_D3
USART3_RX
C10 C8
C4
53
80 113
PC12
I/O FT
PC12
UART5_TX/SDIO_CK
USART3_CK
E10 D8
D8
5
81 114
PD0
I/O FT
OSC_IN(10)
FSMC_D2(11)
CAN_RX
D10 E8
D7
6
82 115
PD1
I/O FT OSC_OUT(10)
FSMC_D3(11)
CAN_TX
E9
B7
A3
54
83 116
PD2
I/O FT
PD2
TIM3_ETR/UART5_RX
SDIO_CMD
-
D9
C7
-
-
84 117
PD3
I/O FT
PD3
FSMC_CLK
USART2_CTS
Table 5. High-density STM32F103xC/D/E pin definitions (continued)
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap

---

---

**📄 Source: PDF Page 36**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
36/144
DocID14611 Rev 12
C9
D7
-
-
85 118
PD4
I/O FT
PD4
FSMC_NOE
USART2_RTS
B9
B6
-
-
86 119
PD5
I/O FT
PD5
FSMC_NWE
USART2_TX
E7
-
-
-
-
120
VSS_10
S
-
VSS_10
-
-
F7
-
-
-
-
121
VDD_10
S
-
VDD_10
-
-
A8
C6
-
-
87 122
PD6
I/O FT
PD6
FSMC_NWAIT
USART2_RX
A9
D6
-
-
88 123
PD7
I/O FT
PD7
FSMC_NE1/FSMC_NCE2
USART2_CK
E8
-
-
-
-
124
PG9
I/O FT
PG9
FSMC_NE2/FSMC_NCE3
-
D8
-
-
-
-
125
PG10
I/O FT
PG10
FSMC_NCE4_1/
FSMC_NE3
-
C8
-
-
-
-
126
PG11
I/O FT
PG11
FSMC_NCE4_2
-
B8
-
-
-
-
127
PG12
I/O FT
PG12
FSMC_NE4
-
D7
-
-
-
-
128
PG13
I/O FT
PG13
FSMC_A24
-
C7
-
-
-
-
129
PG14
I/O FT
PG14
FSMC_A25
-
E6
-
-
-
-
130
VSS_11
S
-
VSS_11
-
-
F6
-
-
-
-
131
VDD_11
S
-
VDD_11
-
-
B7
-
-
-
-
132
PG15
I/O FT
PG15
-
-
A7
A7
A4
55
89 133
PB3
I/O FT
JTDO
SPI3_SCK / I2S3_CK/
PB3/TRACESWO
TIM2_CH2 / 
SPI1_SCK
A6
A6
B4
56
90 134
PB4
I/O FT
NJTRST
SPI3_MISO
PB4 / TIM3_CH1
SPI1_MISO
B6
C5
A5
57
91 135
PB5
I/O
-
PB5
I2C1_SMBA/ SPI3_MOSI 
I2S3_SD
TIM3_CH2 / 
SPI1_MOSI
C6
B5
B5
58
92 136
PB6
I/O FT
PB6
I2C1_SCL(9)/ TIM4_CH1(9)
USART1_TX
D6
A5
C5
59
93 137
PB7
I/O FT
PB7
I2C1_SDA(9) / 
FSMC_NADV / 
TIM4_CH2(9)
USART1_RX
D5
D5
A6
60
94 138
BOOT0
I
-
BOOT0
-
-
C5
B4
D5
61
95 139
PB8
I/O FT
PB8
TIM4_CH3(9)/SDIO_D4
I2C1_SCL/
CAN_RX
B5
A4
B6
62
96 140
PB9
I/O FT
PB9
TIM4_CH4(9)/SDIO_D5
I2C1_SDA / 
CAN_TX
Table 5. High-density STM32F103xC/D/E pin definitions (continued)
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap

---

---

**📄 Source: PDF Page 37**

DocID14611 Rev 12
37/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
A5
D4
-
-
97 141
PE0
I/O FT
PE0
TIM4_ETR / FSMC_NBL0
-
A4
C4
-
-
98 142
PE1
I/O FT
PE1
FSMC_NBL1
-
E5
E5
A7
63
99 143
VSS_3
S
-
VSS_3
-
-
F5
F5
A8
64 100 144
VDD_3
S
-
VDD_3
-
-
1.
I = input, O = output, S = supply.
2.
FT = 5 V tolerant.
3.
Function availability depends on the chosen device.
4.
If several peripherals share the same I/O pin, to avoid conflict between these alternate functions only one peripheral should 
be enabled at a time through the peripheral clock enable bit (in the corresponding RCC peripheral clock enable register).
5.
PC13, PC14 and PC15 are supplied through the power switch. Since the switch only sinks a limited amount of current (3 
mA), the use of GPIOs PC13 to PC15 in output mode is limited: the speed should not exceed 2 MHz with a maximum load 
of 30 pF and these IOs must not be used as a current source (e.g. to drive an LED).
6.
Main function after the first backup domain power-up. Later on, it depends on the contents of the Backup registers even 
after reset (because these registers are not reset by the main reset). For details on how to manage these IOs, refer to the 
Battery backup domain and BKP register description sections in the STM32F10xxx reference manual, available from the 
STMicroelectronics website: www.st.com.
7.
In the WCLSP64 package, the PC3 I/O pin is not bonded and it must be configured by software to output mode (Push-pull) 
and writing 0 to the data register in order to avoid an extra consumption during low-power modes.
8.
Unlike in the LQFP64 package, there is no PC3 in the WLCSP package. The VREF+ functionality is provided instead.
9.
This alternate function can be remapped by software to some other port pins (if available on the used package). For more 
details, refer to the Alternate function I/O and debug configuration section in the STM32F10xxx reference manual, 
available from the STMicroelectronics website: www.st.com.
10. For the WCLSP64/LQFP64 package, the pins number 5 and 6 are configured as OSC_IN/OSC_OUT after reset, however 
the functionality of PD0 and PD1 can be remapped by software on these pins. For the LQFP100/BGA100 and 
LQFP144/BGA144 packages, PD0 and PD1 are available by default, so there is no need for remapping. For more details, 
refer to Alternate function I/O and debug configuration section in the STM32F10xxx reference manual.
11. For devices delivered in LQFP64 packages, the FSMC function is not available.
Table 5. High-density STM32F103xC/D/E pin definitions (continued)
Pins
Pin name
Type(1)
I / O Level(2)
Main 
function(3) 
(after reset)
Alternate functions(4)
LFBGA144
LFBGA100
WLCSP64
LQFP64
LQFP100
LQFP144
Default
Remap

### Code Examples

```elixir
use of GPIOs PC13 to PC15 in output mode is limited: the speed should not exceed 2 MHz with a maximum load
```

```elixir
use these registers are not reset by the main reset). For details on how to manage these IOs, refer to the
```

```typescript
used as a current source (e.g. to drive an LED).
```

```unknown
used package). For more
```

---

---

**📄 Source: PDF Page 38**

Pinouts and pin descriptions
STM32F103xC, STM32F103xD, STM32F103xE
38/144
DocID14611 Rev 12
         
Table 6. FSMC pin definition
Pins
FSMC
LQFP100
BGA100(1)
CF
CF/IDE
NOR/PSRAM/
SRAM
NOR/PSRAM Mux
NAND 16 bit
PE2
-
-
A23
A23
-
Yes
PE3
-
-
A19
A19
-
Yes
PE4
-
-
A20
A20
-
Yes
PE5
-
-
A21
A21
-
Yes
PE6
-
-
A22
A22
-
Yes
PF0
A0
A0
A0
-
-
-
PF1
A1
A1
A1
-
-
-
PF2
A2
A2
A2
-
-
-
PF3
A3
-
A3
-
-
-
PF4
A4
-
A4
-
-
-
PF5
A5
-
A5
-
-
-
PF6
NIORD
NIORD
-
-
-
-
PF7
NREG
NREG
-
-
-
-
PF8
NIOWR
NIOWR
-
-
-
-
PF9
CD
CD
-
-
-
-
PF10
INTR
INTR
-
-
-
-
PF11
NIOS16
NIOS16
-
-
-
-
PF12
A6
-
A6
-
-
-
PF13
A7
-
A7
-
-
-
PF14
A8
-
A8
-
-
-
PF15
A9
-
A9
-
-
-
PG0
A10
-
A10
-
-
-
PG1
-
-
A11
-
-
-
PE7
D4
D4
D4
DA4
D4
Yes
PE8
D5
D5
D5
DA5
D5
Yes
PE9
D6
D6
D6
DA6
D6
Yes
PE10
D7
D7
D7
DA7
D7
Yes
PE11
D8
D8
D8
DA8
D8
Yes
PE12
D9
D9
D9
DA9
D9
Yes
PE13
D10
D10
D10
DA10
D10
Yes
PE14
D11
D11
D11
DA11
D11
Yes
PE15
D12
D12
D12
DA12
D12
Yes
PD8
D13
D13
D13
DA13
D13
Yes

---

---

**📄 Source: PDF Page 39**

DocID14611 Rev 12
39/144
STM32F103xC, STM32F103xD, STM32F103xE
Pinouts and pin descriptions
136
PD9
D14
D14
D14
DA14
D14
Yes
PD10
D15
D15
D15
DA15
D15
Yes
PD11
-
-
A16
A16
CLE
Yes
PD12
-
-
A17
A17
ALE
Yes
PD13
-
-
A18
A18
-
Yes
PD14
D0
D0
D0
DA0
D0
Yes
PD15
D1
D1
D1
DA1
D1
Yes
PG2
-
-
A12
-
-
-
PG3
-
-
A13
-
-
-
PG4
-
-
A14
-
-
-
PG5
-
-
A15
-
-
-
PG6
-
-
-
-
INT2
-
PG7
-
-
-
-
INT3
-
PD0
D2
D2
D2
DA2
D2
Yes
PD1
D3
D3
D3
DA3
D3
Yes
PD3
-
-
CLK
CLK
-
Yes
PD4
NOE
NOE
NOE
NOE
NOE
Yes
PD5
NWE
NWE
NWE
NWE
NWE
Yes
PD6
NWAIT
NWAIT
NWAIT
NWAIT
NWAIT
Yes
PD7
-
-
NE1
NE1
NCE2
Yes
PG9
-
-
NE2
NE2
NCE3
-
PG10
NCE4_1
NCE4_1
NE3
NE3
-
-
PG11
NCE4_2
NCE4_2
-
-
-
-
PG12
-
-
NE4
NE4
-
-
PG13
-
-
A24
A24
-
-
PG14
-
-
A25
A25
-
-
PB7
-
-
NADV
NADV
-
Yes
PE0
-
-
NBL0
NBL0
-
Yes
PE1
-
-
NBL1
NBL1
-
Yes
1.
Ports F and G are not available in devices delivered in 100-pin packages.
Table 6. FSMC pin definition (continued)
Pins
FSMC
LQFP100
BGA100(1)
CF
CF/IDE
NOR/PSRAM/
SRAM
NOR/PSRAM Mux
NAND 16 bit

---

---

**📄 Source: PDF Page 40**

Memory mapping
STM32F103xC, STM32F103xD, STM32F103xE
40/144
DocID14611 Rev 12
4 
Memory mapping
The memory map is shown in Figure 9.
Figure 9. Memory map
512-Mbyte
 block 7
Cortex-M3's
internal
peripherals
512-Mbyte
 block 6
Not used
512-Mbyte
 block 5
FSMC register
512-Mbyte
 block 4
FSMC bank 3
& bank4
512-Mbyte
 block 3
FSMC bank1
& bank2
512-Mbyte
 block 2
Peripherals
512-Mbyte
 block 1
SRAM
0x0000 0000
0x1FFF FFFF
0x2000 0000
0x3FFF FFFF
0x4000 0000
0x5FFF FFFF
0x6000 0000
0x7FFF FFFF
0x8000 0000
0x9FFF FFFF
0xA000 0000
0xBFFF FFFF
0xC000 0000
0xDFFF FFFF
0xE000 0000
0xFFFF FFFF
512-Mbyte
 block 0
Code
Flash
0x0808 0000
0x1FFF EFFF
0x1FFF F000- 0x1FFF F7FF
0x1FFF F800 - 0x1FFF F80F
0x0800 0000
0x0807 FFFF
0x0008 0000
0x07FF FFFF
0x0000 0000
0x0007 FFFF
System memory
Reserved
Reserved
Aliased to Flash or system
memory depending on
BOOT pins
SRAM (64 KB aliased
by bit-banding)
Reserved
0x2000 0000
0x2000 FFFF
0x2001 0000
0x3FFF FFFF
TIM2
TIM3
0x4000 0000 - 0x4000 03FF
TIM4
TIM5
TIM6
TIM7
Reserved
0x4000 0400 - 0x4000 07FF
0x4000 0800 - 0x4000 0BFF
0x4000 0C00 - 0x4000 0FFF
0x4000 1000 - 0x4000 13FF
0x4000 1400 - 0x4000 17FF
0x4000 1800 - 0x4000 27FF
RTC
0x4000 2800 - 0x4000 2BFF
WWDG
0x4000 2C00 - 0x4000 2FFF
IWDG
0x4000 3000 - 0x4000 33FF
Reserved
0x4000 3400 - 0x4000 37FF
SPI2/I2S2
0x4000 3800 - 0x4000 3BFF
SPI3/I2S3
0x4000 3C00 - 0x4000 3FFF
Reserved
0x4000 4000 - 0x4000 43FF
USART2
0x4000 4400 - 0x4000 47FF
0x4000 4800 - 0x4000 4BFF
USART3
UART4
0x4000 4C00 - 0x4000 4FFF
UART5
0x4000 5000 - 0x4000 53FF
I2C1
0x4000 5400 - 0x4000 57FF
I2C2
0x4000 5800 - 0x4000 5BFF
Reserved
0x4000 6800 - 0x4000 6BFF
BKP
0x4000 6C00 - 0x4000 6FFF
PWR
0x4000 7000 - 0x4000 73FF
DAC
0x4000 7400 - 0x4000 77FF
Reserved
0x4000 7800 - 0x4000 FFFF
AFIO
0x4001 0000 - 0x4001 03FF
Port A
EXTI
0x4001 0400 - 0x4001 07FF
0x4001 0800 - 0x4001 0BFF
Port B
0x4001 0C00 - 0x4001 0FFF
Port C
0x4001 1000 - 0x4001 13FF
Port D
0x4001 1400 - 0x4001 17FF
Port E
0x4001 1800 - 0x4001 1BFF
Port F
0x4001 1C00 - 0x4001 1FFF
Port G
0x4001 2000 - 0x4001 23FF
ADC1
0x4001 2400 - 0x4001 27FF
0x4001 2800 - 0x4001 2BFF
SPI1
0x4001 3000 - 0x4001 33FF
0x4001 3400 - 0x4001 37FF
USART1
0x4001 3800 - 0x4001 3BFF
Reserved
0x4001 400 - 0x4001 7FFF
DMA1
0x4002 0000 - 0x4002 03FF
DMA2
0x4002 0400 - 0x4002 07FF
Reserved
0x4002 0400 - 0x4002 0FFF
RCC
0x4002 1000 - 0x4002 13FF
Reserved
0x4002 1400 - 0x4002 1FFF
Flash interface
0x4002 2000 - 0x4002 23FF
Reserved
0x4002 2400 - 0x4002 2FFF
CRC
0x4002 3000 - 0x4002 33FF
Reserved
0x4002 4400 - 0x5FFF FFFF
FSMC bank1 NOR/PSRAM 1
0x6000 0000 - 0x63FF FFFF
FSMC bank1 NOR/PSRAM 2
0x6400 0000 - 0x67FF FFFF
FSMC bank1 NOR/PSRAM 3
0x6800 0000 - 0x6BFF FFFF
FSMC bank1 NOR/PSRAM 4
0x6C00 0000 - 0x6FFF FFFF
FSMC bank2 NAND (NAND1)
0x7000 0000 - 0x7FFF FFFF
FSMC bank3 NAND (NAND2)
0x8000 0000 - 0x8FFF FFFF
FSMC bank4 PCCARD
0x9000 0000 - 0x9FFF FFFF
FSMC register
0xA000 0000 - 0xA000 0FFF
Reserved
0xA000 1000 - 0xBFFF FFFF
ai14753d
Option Bytes
TIM8
ADC2
0x4001 8000 - 0x4001 83FF
0x4001 8400 - 0x4001 FFFF
SDIO
Reserved
ADC3
0x4001 3C00 - 0x4001 3FFF
TIM1
0x4001 2C00 - 0x4001 2FFF
USB registers
Shared USB/CAN SRAM 512
bytes
BxCAN
0x4000 5C00 - 0x4000 5FFF
0x4000 6000 - 0x4000 63FF
0x4000 6400 - 0x4000 67FF

---

---

**📄 Source: PDF Page 41**

DocID14611 Rev 12
41/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
5 
Electrical characteristics
5.1 
Parameter conditions
Unless otherwise specified, all voltages are referenced to VSS.
5.1.1 
Minimum and maximum values
Unless otherwise specified the minimum and maximum values are guaranteed in the worst 
conditions of ambient temperature, supply voltage and frequencies by tests in production on 
100% of the devices with an ambient temperature at TA = 25 °C and TA = TAmax (given by 
the selected temperature range).
Data based on characterization results, design simulation and/or technology characteristics 
are indicated in the table footnotes and are not tested in production. Based on 
characterization, the minimum and maximum values refer to sample tests and represent the 
mean value plus or minus three times the standard deviation (mean±3Σ).
5.1.2 
Typical values
Unless otherwise specified, typical data are based on TA = 25 °C, VDD = 3.3 V (for the 
2 V ≤VDD ≤3.6 V voltage range). They are given only as design guidelines and are not 
tested.
Typical ADC accuracy values are determined by characterization of a batch of samples from 
a standard diffusion lot over the full temperature range, where 95% of the devices have an 
error less than or equal to the value indicated (mean±2Σ).
5.1.3 
Typical curves
Unless otherwise specified, all typical curves are given only as design guidelines and are 
not tested.
5.1.4 
Loading capacitor
The loading conditions used for pin parameter measurement are shown in Figure 10.
5.1.5 
Pin input voltage
The input voltage measurement on a pin of the device is described in Figure 11.
         
Figure 10. Pin loading conditions
Figure 11. Pin input voltage
-36
#   P&
-#5 PIN
-36
-#5 PIN
6).

### Code Examples

```unknown
used for pin parameter measurement are shown in Figure 10.
```

---

---

**📄 Source: PDF Page 42**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
42/144
DocID14611 Rev 12
5.1.6 
Power supply scheme
Figure 12. Power supply scheme
Caution:
In Figure 12, the 4.7 µF capacitor must be connected to VDD3.
5.1.7 
Current consumption measurement
Figure 13. Current consumption measurement scheme
ĂŝϭϱϰϬϭ
sϭͬϮͬͬ͘͘͘ϭϭ
ŶĂůŽŐ͗
ZƐ͕W>>͕
͘͘͘
WŽ ǁĞƌƐǁŝƚĐŚ
sd
'W /ͬK Ɛ
K hd
/E
<ĞƌŶĞůůŽŐŝĐ
;Wh͕
ŝŐŝƚĂů
ΘDĞŵŽƌŝĞƐͿ
ĂĐŬƵƉĐŝƌĐƵŝƚƌǇ
;K^ϯϮ<͕Zd͕
ĂĐŬƵƉƌĞŐŝƐƚĞƌƐͿ
tĂŬĞͲƵƉůŽŐŝĐ
ϭϭп ϭϬϬŶ&
нϭп ϰ͘ϳђ&
ϭ͘ϴͲϯ͘ϲs
ZĞŐƵůĂƚŽƌ
s^^ϭͬϮͬͬ͘͘͘ϭϭ
s
sZ&н
sZ&Ͳ
s^^
ͬ

>ĞǀĞůƐŚŝĨƚĞƌ
/K
>ŽŐŝĐ
s
ϭϬŶ&
нϭђ&
sZ&
ϭϬŶ&
нϭђ&
s
AI
6"!4
6$$
6$$!
)$$?6"!4
)$$

---

---

**📄 Source: PDF Page 43**

DocID14611 Rev 12
43/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
5.2 
Absolute maximum ratings
Stresses above the absolute maximum ratings listed in Table 7: Voltage characteristics, 
Table 8: Current characteristics, and Table 9: Thermal characteristics may cause permanent 
damage to the device. These are stress ratings only and functional operation of the device 
at these conditions is not implied. Exposure to maximum rating conditions for extended 
periods may affect device reliability.
          
          
Table 7. Voltage characteristics
Symbol
Ratings
Min
Max
Unit
VDD–VSS
External main supply voltage (including VDDA 
and VDD)(1)
1.
All main power (VDD, VDDA) and ground (VSS, VSSA) pins must always be connected to the external power 
supply, in the permitted range.
–0.3
4.0
V
VIN
(2)
2.
VIN maximum must always be respected. Refer to Table 8: Current characteristics for the maximum 
allowed injected current values. 
Input voltage on five volt tolerant pin
VSS − 0.3
VDD + 4.0 
Input voltage on any other pin
VSS − 0.3
4.0
|ΔVDDx|
Variations between different VDD power pins
-
50
mV
|VSSX − VSS|
Variations between all the different ground 
pins(3)
3.
Include VREF- pin.
-
50
VESD(HBM)
Electrostatic discharge voltage (human body 
model)
see Section 5.3.12: 
Absolute maximum ratings 
(electrical sensitivity)
-
Table 8. Current characteristics
Symbol
Ratings
 Max.
Unit
IVDD
Total current into VDD/VDDA power lines (source)(1)
1.
All main power (VDD, VDDA) and ground (VSS, VSSA) pins must always be connected to the external power 
supply, in the permitted range.
150
mA
IVSS
Total current out of VSS ground lines (sink)(1)
150
IIO
Output current sunk by any I/O and control pin
25
Output current source by any I/Os and control pin
− 25
IINJ(PIN)
(2)
2.
Negative injection disturbs the analog performance of the device. See note 3 below Table 62 on page 109.
Injected current on five volt tolerant pins(3)
3.
Positive injection is not possible on these I/Os. A negative injection is induced by VIN<VSS. IINJ(PIN) must 
never be exceeded. Refer to Table 7: Voltage characteristics for the maximum allowed input voltage 
values.
-5/+0
Injected current on any other pin(4)
4.
 A positive injection is induced by VIN>VDD while a negative injection is induced by VIN<VSS. IINJ(PIN) must 
never be exceeded. Refer to Table 7: Voltage characteristics for the maximum allowed input voltage 
values.
± 5
ΣIINJ(PIN)
Total injected current (sum of all I/O and control pins)(5)
5.
When several inputs are submitted to a current injection, the maximum ΣIINJ(PIN) is the absolute sum of the 
positive and negative injected currents (instantaneous values). 
± 25

---

---

**📄 Source: PDF Page 44**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
44/144
DocID14611 Rev 12
         
5.3 
Operating conditions
5.3.1 
General operating conditions
         
Table 9. Thermal characteristics
Symbol
Ratings
 Value
Unit
TSTG
Storage temperature range
–65 to +150
°C
TJ
Maximum junction temperature
150
°C
Table 10. General operating conditions
Symbol
Parameter
 Conditions
Min
Max
Unit
fHCLK
Internal AHB clock frequency
-
0 
72
MHz
fPCLK1
Internal APB1 clock frequency
-
0 
36
fPCLK2
Internal APB2 clock frequency
-
0 
72
VDD
Standard operating voltage
-
2
3.6
V
VDDA
(1)
1.
When the ADC is used, refer to Table 59: ADC characteristics.
Analog operating voltage
(ADC not used)
Must be the same potential 
as VDD
(2)
2.
It is recommended to power VDD and VDDA from the same source. A maximum difference of 300 mV 
between VDD and VDDA can be tolerated during power-up and operation.
2
3.6
V
Analog operating voltage
(ADC used)
2.4
3.6
VBAT
Backup operating voltage
-
1.8
3.6
V
PD
Power dissipation at TA = 
85 °C for suffix 6 or TA = 
105 °C for suffix 7(3)
3.
If TA is lower, higher PD values are allowed as long as TJ does not exceed TJmax (see Table 6.7: Thermal 
characteristics on page 133).
LQFP144
-
666
mW
LQFP100
-
434
LQFP64
-
444
LFBGA100
-
500
LFBGA144
-
500
WLCSP64
-
400
TA 
Ambient temperature for 6 
suffix version
Maximum power dissipation 
–40
85
°C
Low-power dissipation(4)
4.
In low-power dissipation state, TA can be extended to this range as long as TJ does not exceed TJmax (see 
Table 6.7: Thermal characteristics on page 133).
–40
105
Ambient temperature for 7 
suffix version
Maximum power dissipation 
–40
105
°C
Low-power dissipation(4)
–40
125
TJ 
Junction temperature range
6 suffix version
–40
105
°C
7 suffix version
–40
125

### Code Examples

```unknown
used, refer to Table 59: ADC characteristics.
```

---

---

**📄 Source: PDF Page 45**

DocID14611 Rev 12
45/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
5.3.2 
Operating conditions at power-up / power-down
The parameters given in Table 11 are derived from tests performed under the ambient 
temperature condition summarized in Table 10.
Table 11. Operating conditions at power-up / power-down
5.3.3 
Embedded reset and power control block characteristics
The parameters given in Table 12 are derived from tests performed under ambient 
temperature and VDD supply voltage conditions summarized in Table 10.
          
Symbol
Parameter
Conditions
Min
Max
Unit
tVDD
VDD rise time rate
-
0
∞
µs/V
VDD fall time rate
20
∞
Table 12. Embedded reset and power control block characteristics
Symbol
Parameter
Conditions
Min
Typ 
Max
Unit
VPVD
Programmable voltage 
detector level selection
PLS[2:0]=000 (rising edge)
2.1
2.18
2.26
V
PLS[2:0]=000 (falling edge)
2
2.08
2.16
PLS[2:0]=001 (rising edge)
2.19
2.28
2.37
PLS[2:0]=001 (falling edge)
2.09
2.18
2.27
PLS[2:0]=010 (rising edge)
2.28
2.38
2.48
PLS[2:0]=010 (falling edge)
2.18
2.28
2.38
PLS[2:0]=011 (rising edge)
2.38
2.48
2.58
PLS[2:0]=011 (falling edge)
2.28
2.38
2.48
PLS[2:0]=100 (rising edge)
2.47
2.58
2.69
PLS[2:0]=100 (falling edge)
2.37
2.48
2.59
PLS[2:0]=101 (rising edge)
2.57
2.68
2.79
PLS[2:0]=101 (falling edge)
2.47
2.58
2.69
PLS[2:0]=110 (rising edge)
2.66
2.78
2.9
PLS[2:0]=110 (falling edge)
2.56
2.68
2.8
PLS[2:0]=111 (rising edge)
2.76
2.88
3
PLS[2:0]=111 (falling edge)
2.66
2.78
2.9
VPVDhyst
(2)
PVD hysteresis
-
-
100
-
mV
VPOR/PDR
Power on/power down 
reset threshold
Falling edge
1.8(1)
1.
The product behavior is guaranteed by design down to the minimum VPOR/PDR value.
1.88
1.96
V
Rising edge
1.84
1.92
2.0
VPDRhyst
(2)
PDR hysteresis
-
-
40
-
mV
TRSTTEMPO
(2)
2.
Guaranteed by design.
Reset temporization
-
1
2.5
4.5
ms

---

---

**📄 Source: PDF Page 46**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
46/144
DocID14611 Rev 12
5.3.4 
Embedded reference voltage
The parameters given in Table 13 are derived from tests performed under ambient 
temperature and VDD supply voltage conditions summarized in Table 10.
          
5.3.5 
Supply current characteristics
The current consumption is a function of several parameters and factors such as the 
operating voltage, ambient temperature, I/O pin loading, device software configuration, 
operating frequencies, I/O pin switching rate, program location in memory and executed 
binary code.
The current consumption is measured as described in Figure 13: Current consumption 
measurement scheme.
All Run-mode current consumption measurements given in this section are performed with a 
reduced code that gives a consumption equivalent to Dhrystone 2.1 code.
Maximum current consumption
The MCU is placed under the following conditions:
•
All I/O pins are in input mode with a static value at VDD or VSS (no load)
•
All peripherals are disabled except when explicitly mentioned
•
The Flash memory access time is adjusted to the fHCLK frequency (0 wait state from 0 
to 24 MHz, 1 wait state from 24 to 48 MHz and 2 wait states above)
•
Prefetch in ON (reminder: this bit must be set before clock setting and bus prescaling)
•
When the peripherals are enabled fPCLK1 = fHCLK/2, fPCLK2 = fHCLK
The parameters given in Table 14, Table 15 and Table 16 are derived from tests performed 
under ambient temperature and VDD supply voltage conditions summarized in Table 10.
Table 13. Embedded internal reference voltage
Symbol
Parameter
Conditions
Min
Typ 
Max
Unit
VREFINT
Internal reference voltage
–40 °C < TA < +105 °C
1.16
1.20
1.26
V
–40 °C < TA < +85 °C
1.16
1.20
1.24
TS_vrefint
(1)
1.
Shortest sampling time can be determined in the application by multiple iterations.
ADC sampling time when 
reading the internal reference 
voltage
-
-
5.1
17.1(2)
2.
Guaranteed by design.
µs
VRERINT
(2)
Internal reference voltage 
spread over the temperature 
range
VDD = 3 V ±10 mV
-
-
10
mV
TCoeff
(2)
Temperature coefficient
-
-
-
100
ppm/°C

---

---

**📄 Source: PDF Page 47**

DocID14611 Rev 12
47/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
         CIAO
Table 14. Maximum current consumption in Run mode, code with data processing
running from Flash
Symbol
Parameter
Conditions
fHCLK
Max(1)
1.
Guaranteed by characterization results.
Unit
TA = 85 °C
TA = 105 °C
IDD
Supply current in 
Run mode
External clock(2), all 
peripherals enabled
2.
External clock is 8 MHz and PLL is on when fHCLK > 8 MHz.
72 MHz
69
70
mA
48 MHz
50
50.5
36 MHz
39
39.5
24 MHz
27
28
16 MHz
20
20.5
8 MHz
11
11.5
External clock(2), all 
peripherals disabled
72 MHz
37
37.5
48 MHz
28
28.5
36 MHz
22
22.5
24 MHz
16.5
17
16 MHz
12.5
13
8 MHz
8
8
Table 15. Maximum current consumption in Run mode, code with data processing
running from RAM
Symbol
Parameter
Conditions
fHCLK
Max(1)
1.
Guaranteed by characterization results at VDD max, fHCLK max.
Unit
TA = 85 °C
TA = 105 °C
IDD
Supply current 
in Run mode
External clock(2), all 
peripherals enabled
2.
External clock is 8 MHz and PLL is on when fHCLK > 8 MHz.
72 MHz
66
67
mA
48 MHz
43.5
45.5
36 MHz
33
35
24 MHz
23
24.5
16 MHz
16
18
8 MHz
9
10.5
External clock(2), all 
peripherals disabled
72 MHz
33
33.5
48 MHz
23
23.5
36 MHz
18
18.5
24 MHz
13
13.5
16 MHz
10
10.5
8 MHz
6
6.5

---

---

**📄 Source: PDF Page 48**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
48/144
DocID14611 Rev 12
Figure 14. Typical current consumption in Run mode versus frequency (at 3.6 V) -
code with data processing running from RAM, peripherals enabled 
Figure 15. Typical current consumption in Run mode versus frequency (at 3.6 V)-
code with data processing running from RAM, peripherals disabled
AI













#ONSUMPTION M!	
4EMPERATURE  #	
-(Z
-(Z
-(Z
-(Z
-(Z
-(Z














#ONSUMPTION M!	
4EMPERATURE  #	
-(Z
-(Z
-(Z
-(Z
-(Z
-(Z
AI

---

---

**📄 Source: PDF Page 49**

DocID14611 Rev 12
49/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         Table 16. Maximum current consumption in Sleep mode, code running from Flash or 
RAM
Symbol
Parameter
Conditions
fHCLK
Max(1)
1.
Guaranteed by characterization results at VDD max, fHCLK max with peripherals enabled.
Unit
TA = 85 °C
TA = 105 °C
IDD
Supply current 
in Sleep mode
External clock(2), all 
peripherals enabled
2.
External clock is 8 MHz and PLL is on when fHCLK > 8 MHz.
72 MHz
45
46
mA
48 MHz
31
32
36 MHz
24
25
24 MHz
17
17.5
16 MHz
12.5
13
8 MHz
8
8
External clock(2), all 
peripherals disabled
72 MHz
8.5
9
48 MHz
7
7.5
36 MHz
6
6.5
24 MHz
5
5.5
16 MHz
4.5
5
8 MHz
4
4

---

---

**📄 Source: PDF Page 50**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
50/144
DocID14611 Rev 12
         
Figure 16. Typical current consumption on VBAT with RTC on vs. temperature
at different VBAT values
Table 17. Typical and maximum current consumptions in Stop and Standby modes
Symbol
Parameter
Conditions
Typ(1)
Max
Unit
VDD/VBAT 
= 2.0 V
VDD/VBAT 
= 2.4 V
VDD/VBAT 
= 3.3 V
TA = 
85 °C
TA = 
105 °C
IDD
Supply current 
in Stop mode
Regulator in run mode, low-speed 
and high-speed internal RC 
oscillators and high-speed oscillator 
OFF (no independent watchdog)
-
34.5
35
379
1130
µA
Regulator in low-power mode, low-
speed and high-speed internal RC 
oscillators and high-speed oscillator 
OFF (no independent watchdog)
-
24.5
25
365
1110
Supply current 
in Standby 
mode
Low-speed internal RC oscillator 
and independent watchdog ON
-
3
3.8
-
-
Low-speed internal RC oscillator 
ON, independent watchdog OFF
-
2.8
3.6
-
-
Low-speed internal RC oscillator 
and independent watchdog OFF, 
low-speed oscillator and RTC OFF
-
1.9
2.1
5(2)
6.5(2)
IDD_VBAT
Backup domain 
supply current Low-speed oscillator and RTC ON
1.05
1.1
1.4
2(2)
2.3(2)
1.
Typical values are measured at TA = 25 °C.
2.
Guaranteed by characterization results.






±



7HPSHUDWXUH&
&RQVXPSWLRQ$
9
9
9
9
9
DL

---

---

**📄 Source: PDF Page 51**

DocID14611 Rev 12
51/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 17. Typical current consumption in Stop mode with regulator in run mode
versus temperature at different VDD values
Figure 18. Typical current consumption in Stop mode with regulator in low-power
mode versus temperature at different VDD values







#
#
#
#
#ONSUMPTION ȝ!	
4EMPERATURE  #	
6
6
6
6
6
AI







#
#
#
#
#ONSUMPTION ȝ!	
4EMPERATURE  #	
6
6
6
6
6
AI







#
#
#
#
#ONSUMPTION ȝ!	
4EMPERATURE  #	
6
6
6
6
6
AI

---

---

**📄 Source: PDF Page 52**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
52/144
DocID14611 Rev 12
Figure 19. Typical current consumption in Standby mode versus temperature at
different VDD values 










#
#
#
#
#ONSUMPTION ȝ!	
4EMPERATURE  #	
6
6
6
6
6
AI

---

---

**📄 Source: PDF Page 53**

DocID14611 Rev 12
53/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Typical current consumption
The MCU is placed under the following conditions:
•
All I/O pins are in input mode with a static value at VDD or VSS (no load).
•
All peripherals are disabled except if it is explicitly mentioned.
•
The Flash access time is adjusted to fHCLK frequency (0 wait state from 0 to 24 MHz, 1 
wait state from 24 to 48 MHZ and 2 wait states above).
•
Ambient temperature and VDD supply voltage conditions summarized in Table 10.
•
Prefetch is ON (Reminder: this bit must be set before clock setting and bus prescaling)
When the peripherals are enabled fPCLK1 = fHCLK/4, fPCLK2 = fHCLK/2, fADCCLK = fPCLK2/4
         
Table 18. Typical current consumption in Run mode, code with data processing
running from Flash
Symbol
Parameter
Conditions
fHCLK
Typ(1)
1.
Typical values are measures at TA = 25 °C, VDD = 3.3 V.
Unit
All peripherals 
enabled(2)
2.
Add an additional power consumption of 0.8 mA per ADC for the analog part. In applications, this 
consumption occurs only while the ADC is on (ADON bit is set in the ADC_CR2 register).
All peripherals 
disabled
IDD
Supply 
current in 
Run mode
External clock(3)
3.
External clock is 8 MHz and PLL is on when fHCLK > 8 MHz.
72 MHz
51
30.5
mA
48 MHz
34.6
20.7
36 MHz
26.6
16.2
24 MHz
18.5
11.4
16 MHz
12.8
8.2
8 MHz
7.2
5
4 MHz
4.2
3.1
2 MHz
2.7
2.1
1 MHz
2
1.7
500 kHz
1.6
1.4
125 kHz
1.3
1.2
Running on high 
speed internal RC 
(HSI), AHB 
prescaler used to 
reduce the 
frequency
64 MHz
45
27
mA
48 MHz
34
20.1
36 MHz
26
15.6
24 MHz
17.9
10.8
16 MHz
12.2
7.6
8 MHz
6.6
4.4
4 MHz
3.6
2.5
2 MHz
2.1
1.5
1 MHz
1.4
1.1
500 kHz
1
0.8
125 kHz
0.7
0.6

---

---

**📄 Source: PDF Page 54**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
54/144
DocID14611 Rev 12
         Table 19. Typical current consumption in Sleep mode, code running from Flash or
RAM
Symbol
Parameter
Conditions
fHCLK
Typ(1)
1.
Typical values are measures at TA = 25 °C, VDD = 3.3 V.
Unit
All peripherals 
enabled(2)
2.
Add an additional power consumption of 0.8 mA per ADC for the analog part. In applications, this 
consumption occurs only while the ADC is on (ADON bit is set in the ADC_CR2 register).
All peripherals 
disabled
IDD
Supply 
current in 
Sleep mode
External clock(3)
3.
External clock is 8 MHz and PLL is on when fHCLK > 8 MHz.
72 MHz
29.5
6.4
mA
48 MHz
20
4.6
36 MHz
15.1
3.6
24 MHz
10.4
2.6
16 MHz
7.2
2
8 MHz
3.9
1.3
4 MHz
2.6
1.2
2 MHz
1.85
1.15
1 MHz
1.5
1.1
500 kHz
1.3
1.05
125 kHz
1.2
1.05
Running on high 
speed internal RC 
(HSI), AHB prescaler 
used to reduce the 
frequency
64 MHz
25.6
5.1
48 MHz
19.4
4
36 MHz
14.5
3
24 MHz
9.8
2
16 MHz
6.6
1.4
8 MHz
3.3
0.7
4 MHz
2
0.6
2 MHz
1.25
0.55
1 MHz
0.9
0.5
500 kHz
0.7
0.45
125 kHz
0.6
0.45

### Code Examples

```unknown
used to reduce the
```

---

---

**📄 Source: PDF Page 55**

DocID14611 Rev 12
55/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
On-chip peripheral current consumption
The current consumption of the on-chip peripherals is given in Table 20. The MCU is placed 
under the following conditions:
•
all I/O pins are in input mode with a static value at VDD or VSS (no load)
•
all peripherals are disabled unless otherwise mentioned
•
the given value is calculated by measuring the current consumption
–
with all peripherals clocked off
–
with only one peripheral clocked on
•
ambient operating temperature and VDD supply voltage conditions summarized in 
Table 7
         
Table 20. Peripheral current consumption
Peripheral
Current 
consumption
Unit
AHB (up to 72 MHz)
DMA1
20,42
µA/MHz
DMA2
19,03
FSMC
52,36
CRC
2,36
SDIO
33,33
BusMatrix(1)
9,72

---

---

**📄 Source: PDF Page 56**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
56/144
DocID14611 Rev 12
APB1 (up to 36 MHz)
APB1-Bridge
7,78
µA/MHz
TIM2
33,06
TIM3
31,94
TIM4
31,67
TIM5
31,94
TIM6
8,06
TIM7
8,06
SPI2/I2S2(2)
8,33
SPI3/I2S3(2)
8,33
USART2
12,22
USART3
12,22
UART4
12,22
UART5
12,22
I2C1
10,28
I2C2
10,00
USB
18,06
CAN1
18,33
DAC(3)
8,06
WWDG
3,89
PWR
1,11
BKP
1,11
IWDG
5,28
Table 20. Peripheral current consumption (continued)
Peripheral
Current 
consumption
Unit

---

---

**📄 Source: PDF Page 57**

DocID14611 Rev 12
57/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
APB2 (up to 72 MHz)
APB2-Bridge
4,17
µA/MHz
GPIOA 
8,47
GPIOB
8,47
GPIOC
6,53
GPIOD
8,47
GPIOE
6,53
GPIOF
6,53
GPIOG
6,11
SPI1
4,72
USART1
12,50
TIM1
22,92
TIM8
22,92
ADC1(4)
17,32
ADC2(4)
15,18
ADC3(4)
14,82
1.
The BusMatrix is automatically active when at least one master is ON. (CPU, DMA1 or DMA2).
2.
When the I2S is enabled, a current consumption equal to 0.02 mA must be added.
3.
When DAC_OU1 or DAC_OUT2 is enabled, a current consumption equal to 0.36 mA must be added.
4.
Specific conditions for measuring ADC current consumption: fHCLK = 56 MHz, fAPB1 = fHCLK/2, fAPB2 = 
fHCLK, fADCCLK = fAPB2/4. When ADON bit in the ADCx_CR2 register is set to 1, a current consumption of 
analog part equal to 0.54 mA must be added for each ADC.
Table 20. Peripheral current consumption (continued)
Peripheral
Current 
consumption
Unit

---

---

**📄 Source: PDF Page 58**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
58/144
DocID14611 Rev 12
5.3.6 
External clock source characteristics
High-speed external user clock generated from an external source
The characteristics given in Table 21 result from tests performed using an high-speed 
external clock source, and under ambient temperature and supply voltage conditions 
summarized in Table 10.
         
Low-speed external user clock generated from an external source
The characteristics given in Table 22 result from tests performed using an low-speed 
external clock source, and under ambient temperature and supply voltage conditions 
summarized in Table 10.
         
Table 21. High-speed external user clock characteristics
Symbol
Parameter
Conditions
Min
Typ
Max
Unit
fHSE_ext
User external clock source 
frequency(1)
1.
Guaranteed by design.
-
1
8
25
MHz
VHSEH
OSC_IN input pin high level voltage
0.7VDD
-
VDD
V
VHSEL
OSC_IN input pin low level voltage
VSS
-
0.3VDD
tw(HSE)
tw(HSE)
OSC_IN high or low time(1)
5
-
-
ns
tr(HSE)
tf(HSE)
OSC_IN rise or fall time(1)
-
-
20
Cin(HSE)
OSC_IN input capacitance(1)
-
-
5
-
pF
DuCy(HSE)
Duty cycle
-
45
-
55
%
IL
OSC_IN Input leakage current 
VSS ≤VIN ≤VDD
-
-
±1
µA
Table 22. Low-speed external user clock characteristics
Symbol
Parameter
Conditions
Min
Typ
Max
Unit
fLSE_ext
User External clock source 
frequency(1)
1.
Guaranteed by design.
-
-
32.768
1000
kHz
VLSEH
OSC32_IN input pin high level 
voltage
0.7VDD
-
VDD
V
VLSEL
OSC32_IN input pin low level 
voltage
VSS
-
0.3VDD
tw(LSE)
tw(LSE)
OSC32_IN high or low time(1)
450
-
-
ns
tr(LSE)
tf(LSE)
OSC32_IN rise or fall time(1)
-
-
50
Cin(LSE)
OSC32_IN input capacitance(1)
-
-
5
-
pF
DuCy(LSE) Duty cycle
-
30
-
70
%
IL
OSC32_IN Input leakage current 
VSS ≤VIN ≤VDD
-
-
±1
µA

### Code Examples

```sql
user clock generated from an external source
```

```unknown
user clock characteristics
```

---

---

**📄 Source: PDF Page 59**

DocID14611 Rev 12
59/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 20. High-speed external clock source AC timing diagram
Figure 21. Low-speed external clock source AC timing diagram
AI
/3#?).
%XTERNAL
34-&
CLOCK SOURCE
6(3%(
TF(3%	
T7(3%	
),

 
4(3%
T
TR(3%	
T7(3%	
F(3%?EXT
6(3%,
DL
26&B,1
([WHUQDO
670)
FORFNVRXUFH
9/6(+
WI/6(
W:/6(
,/


7/6(
W
WU/6(
W:/6(
I/6(BH[W
9/6(/

---

---

**📄 Source: PDF Page 60**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
60/144
DocID14611 Rev 12
High-speed external clock generated from a crystal/ceramic resonator
The high-speed external (HSE) clock can be supplied with a 4 to 16 MHz crystal/ceramic 
resonator oscillator. All the information given in this paragraph are based on 
characterization results obtained with typical external components specified in Table 23. In 
the application, the resonator and the load capacitors have to be placed as close as 
possible to the oscillator pins in order to minimize output distortion and startup stabilization 
time. Refer to the crystal resonator manufacturer for more details on the resonator 
characteristics (frequency, package, accuracy).
         
For CL1 and CL2, it is recommended to use high-quality external ceramic capacitors in the 
5 pF to 25 pF range (typ.), designed for high-frequency applications, and selected to match 
the requirements of the crystal or resonator (see Figure 22). CL1 and CL2 are usually the 
same size. The crystal manufacturer typically specifies a load capacitance which is the 
series combination of CL1 and CL2. PCB and MCU pin capacitance must be included (10 pF 
can be used as a rough estimate of the combined pin and board capacitance) when sizing 
CL1 and CL2. Refer to the application note AN2867 “Oscillator design guide for ST 
microcontrollers” available from the ST website www.st.com.
Figure 22. Typical application with an 8 MHz crystal
1.
REXT value depends on the crystal characteristics.
Table 23. HSE 4-16 MHz oscillator characteristics(1)(2)
1.
Resonator characteristics given by the crystal/ceramic resonator manufacturer.
2.
Guaranteed by characterization results.
Symbol
Parameter
Conditions
Min
Typ
Max
Unit
fOSC_IN
Oscillator frequency
-
4
8
16
MHz
RF
Feedback resistor
-
-
200
-
kΩ 
C
Recommended load capacitance 
versus equivalent serial 
resistance of the crystal (RS)(3)
3.
The relatively low value of the RF resistor offers a good protection against issues resulting from use in a 
humid environment, due to the induced leakage and the bias condition change. However, it is 
recommended to take this point into account if the MCU is used in tough humidity conditions.
RS = 30 Ω
-
30
-
pF
i2
HSE driving current
VDD= 3.3 V, VIN = VSS 
with 30 pF load
-
-
1
mA
gm
Oscillator transconductance
Startup
25
-
-
mA/V
tSU(HSE)
(4)
4.
tSU(HSE) is the startup time measured from the moment it is enabled (by software) to a stabilized 8 MHz 
oscillation is reached. This value is measured for a standard crystal resonator and it can vary significantly 
with the crystal manufacturer
Startup time
 VDD is stabilized
-
2
-
ms
DL
26&B287
26&B,1
I+6(
&/
5)
670)
0+]
UHVRQDWRU
5HVRQDWRUZLWK
LQWHJUDWHGFDSDFLWRUV
%LDV
FRQWUROOHG
JDLQ
5(;7
&/

### Code Examples

```elixir
use high-quality external ceramic capacitors in the
```

```typescript
used as a rough estimate of the combined pin and board capacitance) when sizing
```

```unknown
requirements of the crystal or resonator (see Figure 22). CL1 and CL2 are usually the
```

```unknown
used in tough humidity conditions.
```

---

---

**📄 Source: PDF Page 61**

DocID14611 Rev 12
61/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Low-speed external clock generated from a crystal/ceramic resonator
The low-speed external (LSE) clock can be supplied with a 32.768 kHz crystal/ceramic 
resonator oscillator. All the information given in this paragraph are based on 
characterization results obtained with typical external components specified in Table 24. In 
the application, the resonator and the load capacitors have to be placed as close as 
possible to the oscillator pins in order to minimize output distortion and startup stabilization 
time. Refer to the crystal resonator manufacturer for more details on the resonator 
characteristics (frequency, package, accuracy).
         
Note:
For CL1 and CL2, it is recommended to use high-quality ceramic capacitors in the 5 pF to 
15 pF range selected to match the requirements of the crystal or resonator (see Figure 23). 
CL1 and CL2, are usually the same size. The crystal manufacturer typically specifies a load 
capacitance which is the series combination of CL1 and CL2.
Load capacitance CL has the following formula: CL = CL1 x CL2 / (CL1 + CL2) + Cstray where 
Cstray is the pin capacitance and board or trace PCB-related capacitance. Typically, it is 
between 2 pF and 7 pF.
Caution:
To avoid exceeding the maximum value of CL1 and CL2 (15 pF) it is strongly recommended 
to use a resonator with a load capacitance CL ≤ 7 pF. Never use a resonator with a load 
capacitance of 12.5 pF.
Example: if you choose a resonator with a load capacitance of CL = 6 pF, and Cstray = 2 pF, 
then CL1 = CL2 = 8 pF.
Table 24. LSE oscillator characteristics (fLSE = 32.768 kHz)(1)(2)
Symbol
Parameter
Conditions
Min
Typ
Max
Unit
RF
Feedback resistor
-
-
5
-
MΩ 
C(2)
Recommended load capacitance 
versus equivalent serial 
resistance of the crystal (RS)
RS = 30 kΩ
-
-
15
pF
I2
LSE driving current
VDD = 3.3 V, VIN = VSS
-
-
1.4
µA
gm
Oscillator transconductance
-
5
-
-
µA/V
tSU(LSE)
(3) Startup time 
 VDD is 
stabilized
TA = 50 °C
-
1.5
-
s
TA = 25 °C
-
2.5
-
TA = 10 °C
-
4
-
TA = 0 °C
-
6
-
TA = -10 °C
-
10
-
TA = -20 °C
-
17
-
TA = -30 °C
-
32
-
TA = -40 °C
-
60
-
1.
Guaranteed by characterization results.
2.
Refer to the note and caution paragraphs below the table, and to the application note AN2867 “Oscillator design guide for 
ST microcontrollers”.
3.
 tSU(LSE) is the startup time measured from the moment it is enabled (by software) until a stabilized 32.768 kHz oscillation is 
reached. This value is measured for a standard crystal and it can vary significantly with the crystal manufacturer, PCB 
layout and humidity.

### Code Examples

```elixir
use high-quality ceramic capacitors in the 5 pF to
```

```elixir
use a resonator with a load capacitance CL ≤ 7 pF. Never use a resonator with a load
```

```unknown
requirements of the crystal or resonator (see Figure 23).
```

---

---

**📄 Source: PDF Page 62**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
62/144
DocID14611 Rev 12
Figure 23. Typical application with a 32.768 kHz crystal
5.3.7 
Internal clock source characteristics
The parameters given in Table 25 are derived from tests performed under ambient 
temperature and VDD supply voltage conditions summarized in Table 10.
High-speed internal (HSI) RC oscillator
         
DL
26&B287
26&B,1
I/6(
&/
5)
670)
N+]
UHVRQDWRU
5HVRQDWRUZLWK
LQWHJUDWHGFDSDFLWRUV
%LDV
FRQWUROOHG
JDLQ
&/
Table 25. HSI oscillator characteristics(1)
1.
VDD = 3.3 V, TA = –40 to 105 °C unless otherwise specified.
Symbol
Parameter
Conditions
Min
Typ
Max
Unit
fHSI
Frequency
-
-
8
-
MHz 
DuCy(HSI)
Duty cycle
-
45
-
55
% 
ACCHSI
Accuracy of the HSI 
oscillator
User-trimmed with the RCC_CR 
register(2)
2.
Refer to application note AN2868 “STM32F10xxx internal RC oscillator (HSI) calibration” available from 
the ST website www.st.com.
-
-
1(3)
3.
Guaranteed by design.
%
Factory-
calibrated(4)
4.
Guaranteed by characterization results.
TA = –40 to 105 °C
–2
-
2.5
%
TA = –10 to 85 °C
–1.5
-
2.2
%
TA = 0 to 70 °C
–1.3
-
2
%
TA = 25 °C
–1.1
-
1.8
%
tsu(HSI)
(4)
HSI oscillator 
startup time
-
1
-
2
µs
IDD(HSI)
(4)
HSI oscillator power 
consumption
-
-
80
100
µA

---

---

**📄 Source: PDF Page 63**

DocID14611 Rev 12
63/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Low-speed internal (LSI) RC oscillator
         
Wakeup time from low-power mode
The wakeup times given in Table 27 is measured on a wakeup phase with a 8-MHz HSI RC 
oscillator. The clock source used to wake up the device depends from the current operating 
mode:
•
Stop or Standby mode: the clock source is the RC oscillator
•
Sleep mode: the clock source is the clock that was set before entering Sleep mode.
All timings are derived from tests performed under ambient temperature and VDD supply 
voltage conditions summarized in Table 10. 
         
Table 26. LSI oscillator characteristics (1)
1.
VDD = 3 V, TA = –40 to 105 °C unless otherwise specified.
Symbol
Parameter
Min
Typ
Max
Unit
fLSI
(2)
2.
Guaranteed by characterization results.
Frequency 
30
40
60
kHz 
tsu(LSI)
(3)
3.
Guaranteed by design.
LSI oscillator startup time
-
-
85
µs
IDD(LSI)
(3)
LSI oscillator power consumption
-
0.65
1.2
µA
Table 27. Low-power mode wakeup timings
Symbol
Parameter
Typ
Unit
tWUSLEEP
(1)
1.
The wakeup times are measured from the wakeup event to the point in which the user application code 
reads the first instruction.
Wakeup from Sleep mode
1.8
µs
tWUSTOP
(1)
Wakeup from Stop mode (regulator in run mode)
3.6
µs
Wakeup from Stop mode (regulator in low-power mode)
5.4
tWUSTDBY
(1)
Wakeup from Standby mode
50
µs

### Code Examples

```sql
used to wake up the device depends from the current operating
```

```unknown
user application code
```

---

---

**📄 Source: PDF Page 64**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
64/144
DocID14611 Rev 12
5.3.8 
PLL characteristics
The parameters given in Table 28 are derived from tests performed under ambient 
temperature and VDD supply voltage conditions summarized in Table 10.
         
5.3.9 
Memory characteristics
Flash memory
The characteristics are given at TA = –40 to 105 °C unless otherwise specified.
         
Table 28. PLL characteristics
Symbol
Parameter
Value
Unit
Min
Typ
Max(1)
1.
Guaranteed by characterization results.
fPLL_IN
PLL input clock(2)
2.
Take care of using the appropriate multiplier factors so as to have PLL input clock values compatible with 
the range defined by fPLL_OUT.
1
8.0
25
MHz
PLL input clock duty cycle
40
-
60
%
fPLL_OUT
PLL multiplier output clock
16
-
72
MHz
tLOCK
PLL lock time
-
-
200
µs
Jitter
Cycle-to-cycle jitter
-
-
300
ps
Table 29. Flash memory characteristics
Symbol
Parameter
 Conditions
Min
Typ
Max(1)
1.
Guaranteed by design.
Unit
tprog
16-bit programming time
TA = –40 to +105 °C
40
52.5
70
µs
tERASE
Page (2 KB) erase time
TA = –40 to +105 °C
20
-
40
ms
tME
Mass erase time
TA = –40 to +105 °C
20
-
40
ms
IDD
Supply current 
Read mode
fHCLK = 72 MHz with 2 wait 
states, VDD = 3.3 V
-
-
28
mA
Write mode 
fHCLK = 72 MHz, VDD = 3.3 V
-
-
7
mA
Erase mode 
fHCLK = 72 MHz, VDD = 3.3 V
-
-
5
mA
Power-down mode / Halt,
VDD = 3.0 to 3.6 V
-
-
50
µA
Vprog 
Programming voltage
-
2
-
3.6
V

---

---

**📄 Source: PDF Page 65**

DocID14611 Rev 12
65/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
Table 30. Flash memory endurance and data retention
Symbol
Parameter
 Conditions
Value
Unit
Min(1)
1.
Guaranteed by characterization results.
NEND
Endurance
TA = –40 to +85 °C (6 suffix versions)
TA = –40 to +105 °C (7 suffix versions)
10
kcycles
tRET
Data retention
1 kcycle(2) at TA = 85 °C
2.
Cycling performed over the whole temperature range.
30
Years
1 kcycle(2) at TA = 105 °C
10
10 kcycles(2) at TA = 55 °C
20

---

---

**📄 Source: PDF Page 66**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
66/144
DocID14611 Rev 12
5.3.10 
FSMC characteristics
Asynchronous waveforms and timings
Figure 24 through Figure 27 represent asynchronous waveforms and Table 31 through 
Table 34 provide the corresponding timings. The results shown in these tables are obtained 
with the following FSMC configuration:
•
AddressSetupTime = 0
•
AddressHoldTime = 1
•
DataSetupTime = 1
Figure 24. Asynchronous non-multiplexed SRAM/PSRAM/NOR read waveforms
1.
Mode 2/B, C and D only. In Mode 1, FSMC_NADV is not used.
'DWD
)60&B1(
)60&B1%/>@
)60&B'>@
WY%/B1(
W K'DWDB1(
)60&B12(
$GGUHVV
)60&B$>@
WY$B1(
)60&B1:(
WVX'DWDB1(
WZ1(
069
Z12(
W
W Y12(B1(
W K1(B12(
WK'DWDB12(
W K$B12(
W K%/B12(
WVX'DWDB12(
)60&B1$'9
W Y1$'9B1(
WZ1$'9

---

---

**📄 Source: PDF Page 67**

DocID14611 Rev 12
67/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
Figure 25. Asynchronous non-multiplexed SRAM/PSRAM/NOR write waveforms
1.
Mode 2/B, C and D only. In Mode 1, FSMC_NADV is not used.
Table 31. Asynchronous non-multiplexed SRAM/PSRAM/NOR read timings(1) 
1.
CL = 15 pF.
Symbol
Parameter
Min
Max
Unit
tw(NE)
FSMC_NE low time
5tHCLK – 1.5
5tHCLK + 2
ns
tv(NOE_NE)
FSMC_NEx low to FSMC_NOE low
0.5
1.5
ns
tw(NOE)
FSMC_NOE low time
5tHCLK – 1.5
5tHCLK + 1.5
ns
th(NE_NOE)
FSMC_NOE high to FSMC_NE high hold time
–1.5
-
ns
tv(A_NE)
FSMC_NEx low to FSMC_A valid
-
0
ns
th(A_NOE)
Address hold time after FSMC_NOE high
0.1
-
ns
tv(BL_NE)
FSMC_NEx low to FSMC_BL valid
-
0
ns
th(BL_NOE)
FSMC_BL hold time after FSMC_NOE high
0
-
ns
tsu(Data_NE)
Data to FSMC_NEx high setup time
2tHCLK + 25
-
ns
tsu(Data_NOE) Data to FSMC_NOEx high setup time
2tHCLK + 25
-
ns
th(Data_NOE)
Data hold time after FSMC_NOE high
0
-
ns
th(Data_NE)
Data hold time after FSMC_NEx high
0
-
ns
tv(NADV_NE)
FSMC_NEx low to FSMC_NADV low
-
5
ns
tw(NADV)
FSMC_NADV low time
-
tHCLK + 1.5
ns
1%/
'DWD
)60&B1([
)60&B1%/>@
)60&B'>@
WY%/B1(
WK'DWDB1:(
)60&B12(
$GGUHVV
)60&B$>@
WY$B1(
WZ1:(
)60&B1:(
WY1:(B1(
WK1(B1:(
WK$B1:(
WK%/B1:(
WY'DWDB1(
WZ1(
DL
)60&B1$'9
WY1$'9B1(
WZ1$'9

---

---

**📄 Source: PDF Page 68**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
68/144
DocID14611 Rev 12
         Table 32. Asynchronous non-multiplexed SRAM/PSRAM/NOR write timings(1)(2)
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Symbol
Parameter
Min
Max
Unit
tw(NE)
FSMC_NE low time
3tHCLK – 1
3tHCLK + 2
ns
tv(NWE_NE)
FSMC_NEx low to FSMC_NWE low
tHCLK – 0.5
tHCLK + 1.5
ns
tw(NWE)
FSMC_NWE low time
tHCLK – 0.5
tHCLK + 1.5
ns
th(NE_NWE)
FSMC_NWE high to FSMC_NE high hold time
tHCLK
-
ns
tv(A_NE)
FSMC_NEx low to FSMC_A valid
-
7.5
ns
th(A_NWE)
Address hold time after FSMC_NWE high
tHCLK
-
ns
tv(BL_NE)
FSMC_NEx low to FSMC_BL valid
-
0
ns
th(BL_NWE)
FSMC_BL hold time after FSMC_NWE high
tHCLK – 0.5
-
ns
tv(Data_NE)
FSMC_NEx low to Data valid
-
tHCLK + 7
ns
th(Data_NWE)
Data hold time after FSMC_NWE high
tHCLK
-
ns
tv(NADV_NE)
FSMC_NEx low to FSMC_NADV low
-
5.5
ns
tw(NADV)
FSMC_NADV low time
-
tHCLK + 1.5
ns

---

---

**📄 Source: PDF Page 69**

DocID14611 Rev 12
69/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 26. Asynchronous multiplexed PSRAM/NOR read waveforms
         
Table 33. Asynchronous multiplexed PSRAM/NOR read timings(1)(2)
Symbol
Parameter
Min
Max
Unit
tw(NE)
FSMC_NE low time
7tHCLK – 2
7tHCLK + 2
 ns
tv(NOE_NE)
FSMC_NEx low to FSMC_NOE low
3tHCLK – 0.5
3tHCLK + 1.5
 ns
tw(NOE)
FSMC_NOE low time
4tHCLK – 1
4tHCLK + 2
 ns
th(NE_NOE)
FSMC_NOE high to FSMC_NE high hold time
–1
-
 ns
tv(A_NE)
FSMC_NEx low to FSMC_A valid
-
0
 ns
tv(NADV_NE)
FSMC_NEx low to FSMC_NADV low
3
5
 ns
tw(NADV)
FSMC_NADV low time
tHCLK –1.5
tHCLK + 1.5
 ns
th(AD_NADV)
FSMC_AD (address) valid hold time after 
FSMC_NADV high
tHCLK
-
 ns
th(A_NOE)
Address hold time after FSMC_NOE high
tHCLK -2
-
 ns
th(BL_NOE)
FSMC_BL hold time after FSMC_NOE high
0
-
 ns
tv(BL_NE)
FSMC_NEx low to FSMC_BL valid
-
0
 ns
tsu(Data_NE)
Data to FSMC_NEx high setup time
2tHCLK + 24
-
 ns
tsu(Data_NOE)
Data to FSMC_NOE high setup time
2tHCLK + 25
-
 ns
.",
$ATA
&3-#?.",;=
&3-#?!$;=
TV",?.%	
TH$ATA?.%	
!DDRESS
&3-#?!;=
TV!?.%	
&3-#?.7%
T V!?.%	
AIB
!DDRESS
&3-#?.!$6
T V.!$6?.%	
TW.!$6	
TSU$ATA?.%	
TH!$?.!$6	
&3-#?.%
&3-#?./%
TW.%	
T W./%	
TV./%?.%	
T H.%?./%	
TH!?./%	
TH",?./%	
TSU$ATA?./%	
TH$ATA?./%

---

---

**📄 Source: PDF Page 70**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
70/144
DocID14611 Rev 12
Figure 27. Asynchronous multiplexed PSRAM/NOR write waveforms
         
th(Data_NE)
Data hold time after FSMC_NEx high
0
-
 ns
th(Data_NOE)
Data hold time after FSMC_NOE high
0
-
 ns
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Table 34. Asynchronous multiplexed PSRAM/NOR write timings(1)(2)
Symbol
Parameter
Min
Max
Unit
tw(NE)
FSMC_NE low time
5tHCLK – 1
5tHCLK + 2
ns
tv(NWE_NE)
FSMC_NEx low to FSMC_NWE low
2tHCLK
2tHCLK + 1
ns
tw(NWE)
FSMC_NWE low time
2tHCLK – 1
2tHCLK + 2
ns
th(NE_NWE)
FSMC_NWE high to FSMC_NE high hold time
tHCLK – 1
-
ns
tv(A_NE)
FSMC_NEx low to FSMC_A valid
-
7
ns
tv(NADV_NE)
FSMC_NEx low to FSMC_NADV low
3
5
ns
tw(NADV)
FSMC_NADV low time
tHCLK – 1
tHCLK + 1
ns
Table 33. Asynchronous multiplexed PSRAM/NOR read timings(1)(2) (continued)
Symbol
Parameter
Min
Max
Unit
1%/
'DWD
)60&B1([
)60&B1%/>@
)60&B$'>@
WY%/B1(
WK'DWDB1:(
)60&B12(
$GGUHVV
)60&B$>@
WY$B1(
WZ1:(
)60&B1:(
WY1:(B1(
WK1(B1:(
WK$B1:(
WK%/B1:(
WY$B1(
WZ1(
DL%
$GGUHVV
)60&B1$'9
WY1$'9B1(
WZ1$'9
WY'DWDB1$'9
WK$'B1$'9

---

---

**📄 Source: PDF Page 71**

DocID14611 Rev 12
71/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
th(AD_NADV)
FSMC_AD (address) valid hold time after 
FSMC_NADV high
tHCLK – 3
-
ns
th(A_NWE)
Address hold time after FSMC_NWE high
4tHCLK 
-
ns
tv(BL_NE)
FSMC_NEx low to FSMC_BL valid
-
1.6
ns
th(BL_NWE)
FSMC_BL hold time after FSMC_NWE high
tHCLK – 1.5
-
ns
tv(Data_NADV) FSMC_NADV high to Data valid
-
tHCLK + 1.5
ns
th(Data_NWE)
Data hold time after FSMC_NWE high
tHCLK – 5
-
ns
1.
CL = 15 pF.
2.
BGuaranteed by characterization results.
Table 34. Asynchronous multiplexed PSRAM/NOR write timings(1)(2)
Symbol
Parameter
Min
Max
Unit

---

---

**📄 Source: PDF Page 72**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
72/144
DocID14611 Rev 12
Synchronous waveforms and timings
Figure 28 through Figure 31 represent synchronous waveforms and Table 36 through 
Table 38 provide the corresponding timings. The results shown in these tables are obtained 
with the following FSMC configuration:
•
BurstAccessMode = FSMC_BurstAccessMode_Enable;
•
MemoryType = FSMC_MemoryType_CRAM;
•
WriteBurst = FSMC_WriteBurst_Enable;
•
CLKDivision = 1; (0 is not supported, see the STM32F10xxx reference manual)
•
DataLatency = 1 for NOR Flash; DataLatency = 0 for PSRAM
Figure 28. Synchronous multiplexed NOR/PSRAM read timings
&3-#?#,+
&3-#?.%X
&3-#?.!$6
&3-#?!;=
&3-#?./%
&3-#?!$;=
!$;=
$
$
&3-#?.7!)4
7!)4#&'  B 7!)40/,  B	
&3-#?.7!)4
7!)4#&'  B 7!)40/,  B	
TW#,+	
TW#,+	
$ATA LATENCY  
"53452.  
TD#,+,.%X,	
TD#,+,.%X(	
TD#,+,.!$6,	
TD#,+,!6	
TD#,+,.!$6(	
TD#,+,!)6	
TD#,+(./%,	
TD#,+,./%(	
TD#,+,!$6	
TD#,+,!$)6	
TSU!$6#,+(	
TH#,+(!$6	
TSU!$6#,+(	
TH#,+(!$6	
TSU.7!)46#,+(	
TH#,+(.7!)46	
TSU.7!)46#,+(	
TH#,+(.7!)46	
TSU.7!)46#,+(	
TH#,+(.7!)46	
AII
$

---

---

**📄 Source: PDF Page 73**

DocID14611 Rev 12
73/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
Table 35. Synchronous multiplexed NOR/PSRAM read timings(1)(2)
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Symbol
Parameter
Min
Max
Unit
tw(CLK)
FSMC_CLK period
27.7
-
 ns
td(CLKL-NExL)
FSMC_CLK low to FSMC_NEx low (x = 0...2)
-
1.5
 ns
td(CLKL-NExH)
FSMC_CLK low to FSMC_NEx high (x = 0...2)
2
-
 ns
td(CLKL-NADVL)
FSMC_CLK low to FSMC_NADV low
-
4
 ns
td(CLKL-NADVH)
FSMC_CLK low to FSMC_NADV high
5
-
 ns
td(CLKL-AV)
FSMC_CLK low to FSMC_Ax valid (x = 16...25)
-
0
 ns
td(CLKL-AIV)
FSMC_CLK low to FSMC_Ax invalid (x = 16...25)
2
-
 ns
td(CLKL-NOEL)
FSMC_CLK low to FSMC_NOE low
-
1
 ns
td(CLKL-NOEH)
FSMC_CLK low to FSMC_NOE high
 1.5
-
 ns
td(CLKL-ADV)
FSMC_CLK low to FSMC_AD[15:0] valid
-
12
 ns
td(CLKL-ADIV)
FSMC_CLK low to FSMC_AD[15:0] invalid
0
-
 ns
tsu(ADV-CLKH)
FSMC_A/D[15:0] valid data before FSMC_CLK 
high
6
-
 ns
th(CLKH-ADV)
FSMC_A/D[15:0] valid data after FSMC_CLK high
0
-
 ns
tsu(NWAITV-CLKH) FSMC_NWAIT valid before FSMC_CLK high
8
-
 ns
th(CLKH-NWAITV)
FSMC_NWAIT valid after FSMC_CLK high
2
-
 ns

---

---

**📄 Source: PDF Page 74**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
74/144
DocID14611 Rev 12
Figure 29. Synchronous multiplexed PSRAM write timings
&3-#?#,+
&3-#?.%X
&3-#?.!$6
&3-#?!;=
&3-#?.7%
&3-#?!$;=
!$;=
$
$
&3-#?.7!)4
7!)4#&'  B 7!)40/,  B	
TW#,+	
TW#,+	
$ATA LATENCY  
"53452.  
TD#,+,.%X,	
TD#,+,.%X(	
TD#,+,.!$6,	
TD#,+,!6	
TD#,+,.!$6(	
TD#,+,!)6	
TD#,+,.7%(	
TD#,+,.7%,	
TD#,+,.",(	
TD#,+,!$6	
TD#,+,!$)6	
TD#,+,$ATA	
TSU.7!)46#,+(	
TH#,+(.7!)46	
AIG
TD#,+,$ATA	
&3-#?.",

---

---

**📄 Source: PDF Page 75**

DocID14611 Rev 12
75/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
Table 36. Synchronous multiplexed PSRAM write timings(1)(2)
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Symbol
Parameter
Min
Max
Unit
tw(CLK)
FSMC_CLK period
27.7
-
 ns
td(CLKL-NExL)
FSMC_CLK low to FSMC_Nex low (x = 0...2)
-
2
 ns
td(CLKL-NExH)
FSMC_CLK low to FSMC_NEx high (x = 0...2)
2
-
 ns
td(CLKL-NADVL)
FSMC_CLK low to FSMC_NADV low
-
4
 ns
td(CLKL-NADVH)
FSMC_CLK low to FSMC_NADV high
5
-
 ns
td(CLKL-AV)
FSMC_CLK low to FSMC_Ax valid (x = 16...25)
-
0
 ns
td(CLKL-AIV)
FSMC_CLK low to FSMC_Ax invalid (x = 16...25)
2
-
 ns
td(CLKL-NWEL)
FSMC_CLK low to FSMC_NWE low
-
1
 ns
td(CLKL-NWEH)
FSMC_CLK low to FSMC_NWE high
1
-
 ns
td(CLKL-ADV)
FSMC_CLK low to FSMC_AD[15:0] valid
-
12
 ns
td(CLKL-ADIV)
FSMC_CLK low to FSMC_AD[15:0] invalid
3
-
 ns
td(CLKL-Data)
FSMC_A/D[15:0] valid after FSMC_CLK low
-
6
 ns
td(CLKL-NBLH)
FSMC_CLK low to FSMC_NBL high
1
-
 ns
tsu(NWAITV-CLKH)
FSMC_NWAIT valid before FSMC_CLK high
7
-
 ns
th(CLKH-NWAITV)
FSMC_NWAIT valid after FSMC_CLK high
2
-
 ns

---

---

**📄 Source: PDF Page 76**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
76/144
DocID14611 Rev 12
Figure 30. Synchronous non-multiplexed NOR/PSRAM read timings
         
Table 37. Synchronous non-multiplexed NOR/PSRAM read timings(1)(2)
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Symbol
Parameter
Min
Max
Unit
tw(CLK)
FSMC_CLK period
27.7
-
 ns
td(CLKL-NExL)
FSMC_CLK low to FSMC_NEx low (x = 0...2)
-
1.5
 ns
td(CLKL-NExH)
FSMC_CLK low to FSMC_NEx high (x = 0...2)
2
-
 ns
td(CLKL-NADVL)
FSMC_CLK low to FSMC_NADV low
-
4
 ns
td(CLKL-NADVH)
FSMC_CLK low to FSMC_NADV high
5
-
 ns
td(CLKL-AV)
FSMC_CLK low to FSMC_Ax valid (x = 0...25)
-
0
 ns
td(CLKL-AIV)
FSMC_CLK low to FSMC_Ax invalid (x = 0...25)
4
-
 ns
td(CLKL-NOEL)
FSMC_CLK low to FSMC_NOE low
-
1.5
 ns
td(CLKL-NOEH)
FSMC_CLK low to FSMC_NOE high
1.5
-
 ns
tsu(DV-CLKH)
FSMC_D[15:0] valid data before FSMC_CLK high
6.5
-
 ns
th(CLKH-DV)
FSMC_D[15:0] valid data after FSMC_CLK high
7
-
 ns
tsu(NWAITV-CLKH) FSMC_NWAIT valid before FSMC_SMCLK high
7
-
 ns
th(CLKH-NWAITV)
FSMC_NWAIT valid after FSMC_CLK high
2
-
 ns
&3-#?#,+
&3-#?.%X
&3-#?!;=
&3-#?./%
&3-#?$;=
$
$
&3-#?.7!)4
7!)4#&'  B 7!)40/,  B	
&3-#?.7!)4
7!)4#&'  B 7!)40/,  B	
TW#,+	
TW#,+	
$ATA LATENCY  
"53452.  
TD#,+,.%X,	
TD#,+,.%X(	
TD#,+,!6	
TD#,+,!)6	
TD#,+(./%,	
TD#,+,./%(	
TSU$6#,+(	
TH#,+($6	
TSU$6#,+(	
TH#,+($6	
TSU.7!)46#,+(	
TH#,+(.7!)46	
TSU.7!)46#,+(	
T H#,+(.7!)46	
TSU.7!)46#,+(	
TH#,+(.7!)46	
AIH
&3-#?.!$6
TD#,+,.!$6,	
TD#,+,.!$6(	
$

---

---

**📄 Source: PDF Page 77**

DocID14611 Rev 12
77/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 31. Synchronous non-multiplexed PSRAM write timings
         
Table 38. Synchronous non-multiplexed PSRAM write timings(1)(2)
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Symbol
Parameter
Min
Max
Unit
tw(CLK)
 FSMC_CLK period
27.7
-
 ns
td(CLKL-NExL)
 FSMC_CLK low to FSMC_NEx low (x = 0...2)
-
2
 ns
td(CLKL-NExH)
 FSMC_CLK low to FSMC_NEx high (x = 0...2)
2
-
 ns
td(CLKL-NADVL)
 FSMC_CLK low to FSMC_NADV low
-
4
 ns
td(CLKL-NADVH)
 FSMC_CLK low to FSMC_NADV high
5
-
 ns
td(CLKL-AV)
 FSMC_CLK low to FSMC_Ax valid (x = 16...25)
-
0
 ns
td(CLKL-AIV)
 FSMC_CLK low to FSMC_Ax invalid (x = 16...25)
2
-
 ns
td(CLKL-NWEL)
 FSMC_CLK low to FSMC_NWE low
-
1
 ns
td(CLKL-NWEH)
 FSMC_CLK low to FSMC_NWE high
1
-
 ns
td(CLKL-Data)
 FSMC_D[15:0] valid data after FSMC_CLK low
-
6
 ns
td(CLKL-NBLH)
 FSMC_CLK low to FSMC_NBL high
1
-
 ns
tsu(NWAITV-CLKH)
 FSMC_NWAIT valid before FSMC_CLK high
7
-
 ns
th(CLKH-NWAITV)
 FSMC_NWAIT valid after FSMC_CLK high
2
-
 ns
)60&B&/.
)60&B1([
)60&B1$'9
)60&B$>@
)60&B12(
)60&B$'>@
$'>@
'
'
)60&B1:$,7
:$,7&)* E:$,732/E
)60&B1:$,7
:$,7&)* E:$,732/E
WZ&/.
WZ&/.
'DWDODWHQF\ 
%867851 
WG&/./1([/
WG&/./1([+
WG&/./1$'9/
WG&/./$9
WG&/./1$'9+
WG&/./$,9
WG&/.+12(/
WG&/./12(+
WG&/./$'9
WG&/./$',9
WVX$'9&/.+
WK&/.+$'9
WVX$'9&/.+
WK&/.+$'9
WVX1:$,79&/.+
WK&/.+1:$,79
WVX1:$,79&/.+
WK&/.+1:$,79
WVX1:$,79&/.+
WK&/.+1:$,79
DLK

---

---

**📄 Source: PDF Page 78**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
78/144
DocID14611 Rev 12
PC Card/CompactFlash controller waveforms and timings
Figure 32 through Figure 37 represent synchronous waveforms and Table 39 provides the 
corresponding timings. The results shown in this table are obtained with the following FSMC 
configuration:
•
COM.FSMC_SetupTime = 0x04;
•
COM.FSMC_WaitSetupTime = 0x07;
•
COM.FSMC_HoldSetupTime = 0x04;
•
COM.FSMC_HiZSetupTime = 0x00;
•
ATT.FSMC_SetupTime = 0x04;
•
ATT.FSMC_WaitSetupTime = 0x07;
•
ATT.FSMC_HoldSetupTime = 0x04;
•
ATT.FSMC_HiZSetupTime = 0x00;
•
IO.FSMC_SetupTime = 0x04;
•
IO.FSMC_WaitSetupTime = 0x07;
•
IO.FSMC_HoldSetupTime = 0x04;
•
IO.FSMC_HiZSetupTime = 0x00;
•
TCLRSetupTime = 0;
•
TARSetupTime = 0;
Figure 32. PC Card/CompactFlash controller waveforms for common memory read 
access
1.
FSMC_NCE4_2 remains high (inactive during 8-bit access.
)60&B1:(
WZ12(
)60&B12(
)60&B'>@
)60&B$>@
)60&B1&(B
)60&B1&(B
)60&B15(*
)60&B1,2:5
)60&B1,25'
WG1&(B12(
WVX'12(
WK12('
WY1&([$
WG15(*1&([
WG1,25'1&([
WK1&([$,
WK1&([15(*
WK1&([1,25'
WK1&([1,2:5
DLE

---

---

**📄 Source: PDF Page 79**

DocID14611 Rev 12
79/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 33. PC Card/CompactFlash controller waveforms for common memory write 
access
WG1&(B1:(
WZ1:(
WK1:('
WY1&(B$
WG15(*1&(B
WG1,25'1&(B
WK1&(B$,
0(0[+,= 
WY1:('
WK1&(B15(*
WK1&(B1,25'
WK1&(B1,2:5
DLE
)60&B1:(
)60&B12(
)60&B'>@
)60&B$>@
)60&B1&(B
)60&B15(*
)60&B1,2:5
)60&B1,25'
WG1:(1&(B
WG'1:(
)60&B1&(B +LJK

---

---

**📄 Source: PDF Page 80**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
80/144
DocID14611 Rev 12
Figure 34. PC Card/CompactFlash controller waveforms for attribute memory read
access
1.
Only data bits 0...7 are read (bits 8...15 are disregarded).
WG1&(B12(
WZ12(
WVX'12(
WK12('
WY1&(B$
WK1&(B$,
WG15(*1&(B
WK1&(B15(*
DLE
)60&B1:(
)60&B12(
)60&B'>@
)60&B$>@
)60&B1&(B
)60&B1&(B
)60&B15(*
)60&B1,2:5
)60&B1,25'
WG12(1&(B
+LJK

---

---

**📄 Source: PDF Page 81**

DocID14611 Rev 12
81/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 35. PC Card/CompactFlash controller waveforms for attribute memory write
access
1.
Only data bits 0...7 are driven (bits 8...15 remains HiZ).
Figure 36. PC Card/CompactFlash controller waveforms for I/O space read access
WZ1:(
WY1&(B$
WG15(*1&(B
WK1&(B$,
WK1&(B15(*
WY1:('
DLE
)60&B1:(
)60&B12(
)60&B'>@
)60&B$>@
)60&B1&(B
)60&B1&(B
)60&B15(*
)60&B1,2:5
)60&B1,25'
WG1:(1&(B
+LJK
WG1&(B1:(
WG1,25'1&(B
WZ1,25'
WVX'1,25'
WG1,25''
WY1&([$
WK1&(B$,
DL%
)60&B1:(
)60&B12(
)60&B'>@
)60&B$>@
)60&B1&(B
)60&B1&(B
)60&B15(*
)60&B1,2:5
)60&B1,25'

---

---

**📄 Source: PDF Page 82**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
82/144
DocID14611 Rev 12
Figure 37. PC Card/CompactFlash controller waveforms for I/O space write access
         
WG1&(B1,2:5
WZ1,2:5
WY1&([$
WK1&(B$,
WK1,2:5'
$77[+,= 
WY1,2:5'
DLF
)60&B1:(
)60&B12(
)60&B'>@
)60&B$>@
)60&B1&(B
)60&B1&(B
)60&B15(*
)60&B1,2:5
)60&B1,25'
Table 39. Switching characteristics for PC Card/CF read and write cycles(1)(2)
Symbol
Parameter
Min
Max
Unit
tv(NCEx-A) 
tv(NCE4_1-A)
FSMC_NCEx low (x = 4_1/4_2) to FSMC_Ay valid (y = 
0...10) FSMC_NCE4_1 low (x = 4_1/4_2) to FSMC_Ay 
valid (y = 0...10)
-
0
 ns
th(NCEx-AI) 
th(NCE4_1-AI)
FSMC_NCEx high (x = 4_1/4_2) to FSMC_Ax invalid (x = 
0...10) FSMC_NCE4_1 high (x = 4_1/4_2) to FSMC_Ax 
invalid (x = 0...10)
2.5
-
 ns
td(NREG-NCEx) 
td(NREG-NCE4_1)
FSMC_NCEx low to FSMC_NREG valid FSMC_NCE4_1 
low to FSMC_NREG valid
-
5
 ns
th(NCEx-NREG) 
th(NCE4_1-NREG)
FSMC_NCEx high to FSMC_NREG invalid FSMC_NCE4_1 
high to FSMC_NREG invalid
tHCLK + 3
-
 ns
td(NCE4_1-NOE)
FSMC_NCE4_1 low to FSMC_NOE low
-
5tHCLK + 2
 ns
tw(NOE)
FSMC_NOE low width
8tHCLK –1.5
8tHCLK + 1
 ns
td(NOE-NCE4_1
FSMC_NOE high to FSMC_NCE4_1 high
5tHCLK + 2
-
 ns
tsu(D-NOE)
FSMC_D[15:0] valid data before FSMC_NOE high
25
-
 ns
th(NOE-D)
FSMC_D[15:0] valid data after FSMC_NOE high
15
-
 ns
tw(NWE)
FSMC_NWE low width
8tHCLK – 1
8tHCLK + 2
 ns
td(NWE-NCE4_1)
FSMC_NWE high to FSMC_NCE4_1 high
5tHCLK + 2
-
 ns
td(NCE4_1-NWE)
FSMC_NCE4_1 low to FSMC_NWE low
-
5tHCLK + 1.5
 ns
tv(NWE-D)
FSMC_NWE low to FSMC_D[15:0] valid
-
0
 ns
th(NWE-D)
FSMC_NWE high to FSMC_D[15:0] invalid
11tHCLK
-
 ns
td(D-NWE)
FSMC_D[15:0] valid before FSMC_NWE high
13tHCLK
-
 ns

---

---

**📄 Source: PDF Page 83**

DocID14611 Rev 12
83/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
tw(NIOWR)
FSMC_NIOWR low width
8tHCLK + 3
-
 ns
tv(NIOWR-D)
FSMC_NIOWR low to FSMC_D[15:0] valid
-
5tHCLK +1
 ns
th(NIOWR-D)
FSMC_NIOWR high to FSMC_D[15:0]   invalid
11tHCLK
-
 ns
td(NCE4_1-NIOWR) FSMC_NCE4_1 low to FSMC_NIOWR valid
-
5tHCLK+3ns
 ns
th(NCEx-NIOWR) 
th(NCE4_1-NIOWR)
FSMC_NCEx high to FSMC_NIOWR invalid 
FSMC_NCE4_1 high to FSMC_NIOWR invalid
5tHCLK – 5
-
 ns
td(NIORD-NCEx) 
td(NIORD-NCE4_1)
FSMC_NCEx low to FSMC_NIORD valid FSMC_NCE4_1 
low to FSMC_NIORD valid
-
5tHCLK + 2.5
 ns
th(NCEx-NIORD) 
th(NCE4_1-NIORD)
FSMC_NCEx high to FSMC_NIORD invalid 
FSMC_NCE4_1 high to FSMC_NIORD invalid
5tHCLK – 5
-
 ns
tsu(D-NIORD)
FSMC_D[15:0] valid before FSMC_NIORD high
4.5
-
 ns
td(NIORD-D)
FSMC_D[15:0] valid after FSMC_NIORD high
9
-
 ns
tw(NIORD)
FSMC_NIORD low width
8tHCLK + 2
-
 ns
1.
CL = 15 pF.
2.
Guaranteed by characterization results.
Table 39. Switching characteristics for PC Card/CF read and write cycles(1)(2) (continued)
Symbol
Parameter
Min
Max
Unit

---

---

**📄 Source: PDF Page 84**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
84/144
DocID14611 Rev 12
NAND controller waveforms and timings
Figure 38 through Figure 41 represent synchronous waveforms and Table 39 provides the 
corresponding timings. The results shown in this table are obtained with the following FSMC 
configuration:
•
COM.FSMC_SetupTime = 0x01;
•
COM.FSMC_WaitSetupTime = 0x03;
•
COM.FSMC_HoldSetupTime = 0x02;
•
COM.FSMC_HiZSetupTime = 0x01;
•
ATT.FSMC_SetupTime = 0x01;
•
ATT.FSMC_WaitSetupTime = 0x03;
•
ATT.FSMC_HoldSetupTime = 0x02;
•
ATT.FSMC_HiZSetupTime = 0x01;
•
Bank = FSMC_Bank_NAND;
•
MemoryDataWidth = FSMC_MemoryDataWidth_16b;
•
ECC = FSMC_ECC_Enable;
•
ECCPageSize = FSMC_ECCPageSize_512Bytes;
•
TCLRSetupTime = 0;
•
TARSetupTime = 0;
Figure 38. NAND controller waveforms for read access
)60&B1:(
)60&B12(15(
)60&B'>@
WVX'12(
WK12('
DLE
$/()60&B$
&/()60&B$
)60&B1&([ /RZ
WG$/(12(
WK12($/(

---

---

**📄 Source: PDF Page 85**

DocID14611 Rev 12
85/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 39. NAND controller waveforms for write access
Figure 40. NAND controller waveforms for common memory read access
AIC
WK1:('
WY1:('
)60&B1:(
)60&B12(15(
)60&B'>@
$/()60&B$
&/()60&B$
)60&B1&([
WG$/(1:(
WK1:($/(
)60&B1:(
)60&B12(
)60&B'>@
WZ12(
WVX'12(
WK12('
DLE
$/()60&B$
&/()60&B$
)60&B1&([ /RZ
WG$/(12(
WK12($/(

---

---

**📄 Source: PDF Page 86**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
86/144
DocID14611 Rev 12
Figure 41. NAND controller waveforms for common memory write access
         
Table 40. Switching characteristics for NAND Flash read and write cycles(1)
1.
CL = 15 pF.
Symbol
Parameter
Min
Max
Unit
td(D-NWE)
(2)
2.
Guaranteed by characterization results.
FSMC_D[15:0] valid before FSMC_NWE high
5tHCLK + 12
-
ns
tw(NOE)
(2)
FSMC_NWE low width
4tHCLK-1.5
4tHCLK+1.5
ns
tsu(D-NOE)
(2)
FSMC_D[15:0] valid data before 
FSMC_NOE
high
25
-
 ns
th(NOE-D)
(2)
FSMC_D[15:0] valid data after FSMC_NOE high
7
-
-
tw(NWE)
(2)
FSMC_NWE low width
4tHCLK-1
4tHCLK+1
 ns
tv(NWE-D)
(2)
FSMC_NWE low to FSMC_D[15:0] valid
-
0
 ns
th(NWE-D)
(2)
FSMC_NWE high to FSMC_D[15:0] invalid
2tHCLK + 4
-
 ns
td(ALE-NWE)
(3)
3.
Guaranteed by design.
FSMC_ALE valid before FSMC_NWE low
-
3tHCLK + 1.5
 ns
th(NWE-ALE)
(3) FSMC_NWE high to FSMC_ALE invalid
3tHCLK + 4.5
-
 ns
td(ALE-NOE)
(3)
FSMC_ALE valid before FSMC_NOE low
-
3tHCLK+ 2
 ns
th(NOE-ALE)
(3)
FSMC_NWE high to FSMC_ALE invalid
3tHCLK+ 4.5
-
 ns
WZ1:(
WK1:('
WY1:('
DLE
)60&B1:(
)60&B12(
)60&B'>@
WG'1:(
$/()60&B$
&/()60&B$
)60&B1&([ /RZ
WG$/(1:(
WK1:($/(

---

---

**📄 Source: PDF Page 87**

DocID14611 Rev 12
87/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
5.3.11 
EMC characteristics
Susceptibility tests are performed on a sample basis during device characterization.
Functional EMS (electromagnetic susceptibility)
While a simple application is executed on the device (toggling 2 LEDs through I/O ports). 
the device is stressed by two electromagnetic events until a failure occurs. The failure is 
indicated by the LEDs:
•
Electrostatic discharge (ESD) (positive and negative) is applied to all device pins until 
a functional disturbance occurs. This test is compliant with the IEC 61000-4-2 standard.
•
FTB: A Burst of Fast Transient voltage (positive and negative) is applied to VDD and 
VSS through a 100 pF capacitor, until a functional disturbance occurs. This test is 
compliant with the IEC 61000-4-4 standard.
A device reset allows normal operations to be resumed. 
The test results are given in Table 41. They are based on the EMS levels and classes 
defined in application note AN1709.
         
Designing hardened software to avoid noise problems
EMC characterization and optimization are performed at component level with a typical 
application environment and simplified MCU software. It should be noted that good EMC 
performance is highly dependent on the user application and the software in particular.
Therefore it is recommended that the user applies EMC software optimization and 
prequalification tests in relation with the EMC level requested for his application.
Software recommendations
The software flowchart must include the management of runaway conditions such as:
•
Corrupted program counter
•
Unexpected reset
•
Critical Data corruption (control registers...)
Prequalification trials
Most of the common failures (unexpected reset and program counter corruption) can be 
reproduced by manually forcing a low state on the NRST pin or the Oscillator pins for 1 
second.
Table 41. EMS characteristics
Symbol
Parameter
Conditions
Level/
Class
VFESD
Voltage limits to be applied on any I/O pin to 
induce a functional disturbance
VDD = 3.3 V, LQFP144, TA = +25 °C, 
fHCLK = 72 MHz
conforms to IEC 61000-4-2
2B
VEFTB
Fast transient voltage burst limits to be 
applied through 100 pF on VDD and VSS 
pins to induce a functional disturbance
VDD = 3.3 V, LQFP144, TA = +25 
°C, 
fHCLK = 72 MHz
conforms to IEC 61000-4-4
4A

### Code Examples

```unknown
user application and the software in particular.
```

```unknown
user applies EMC software optimization and
```

```unknown
include the management of runaway conditions such as:
required on six parts to assess the latch-up
```

---

---

**📄 Source: PDF Page 88**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
88/144
DocID14611 Rev 12
To complete these trials, ESD stress can be applied directly on the device, over the range of 
specification values. When unexpected behavior is detected, the software can be hardened 
to prevent unrecoverable errors occurring (see application note AN1015).
Electromagnetic Interference (EMI)
The electromagnetic field emitted by the device are monitored while a simple application is 
executed (toggling 2 LEDs through the I/O ports). This emission test is compliant with 
IEC 61967-2 standard which specifies the test board and the pin loading.
         
5.3.12 
Absolute maximum ratings (electrical sensitivity)
Based on three different tests (ESD, LU) using specific measurement methods, the device is 
stressed in order to determine its performance in terms of electrical sensitivity.
Electrostatic discharge (ESD)
Electrostatic discharges (a positive then a negative pulse separated by 1 second) are 
applied to the pins of each sample according to each pin combination. The sample size 
depends on the number of supply pins in the device (3 parts × (n+1) supply pins). This test 
conforms to the JESD22-A114/C101 standard.
          
Static latch-up
Two complementary static tests are required on six parts to assess the latch-up 
performance: 
•
A supply overvoltage is applied to each power supply pin
•
A current injection is applied to each input, output and configurable I/O pin
These tests are compliant with EIA/JESD 78A IC latch-up standard.
Table 42. EMI characteristics
Symbol
Parameter
Conditions
Monitored
frequency band
Max vs. [fHSE/fHCLK]
Unit
8/48 MHz
8/72 MHz
SEMI
Peak level
VDD = 3.3 V, TA = 25 °C,
LQFP144 package
compliant with IEC 
61967-2
0.1 to 30 MHz
8
12
dBµV
30 to 130 MHz
31
21
130 MHz to 1GHz
28
33
SAE EMI Level
4
4
-
Table 43. ESD absolute maximum ratings
Symbol
Ratings
Conditions
Class
Maximum 
value(1)
1.
Guaranteed by characterization results.
Unit
VESD(HBM)
Electrostatic discharge 
voltage (human body model)
TA = +25 °C, conforming 
to JESD22-A114
2
2000
V
VESD(CDM)
Electrostatic discharge 
voltage (charge device 
model)
TA = +25 °C, conforming 
to JESD22-C101
III
500

---

---

**📄 Source: PDF Page 89**

DocID14611 Rev 12
89/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
5.3.13 
I/O current injection characteristics
As a general rule, current injection to the I/O pins, due to external voltage below VSS or 
above VDD (for standard, 3 V-capable I/O pins) should be avoided during normal product 
operation. However, in order to give an indication of the robustness of the microcontroller in 
cases when abnormal injection accidentally happens, susceptibility tests are performed on a 
sample basis during device characterization.
Functional susceptibilty to I/O current injection 
While a simple application is executed on the device,  the device is stressed by injecting 
current into the I/O pins programmed in floating input mode. While current is injected into 
the I/O pin, one at a time, the device is checked for functional failures. 
The failure is indicated by an out of range parameter: ADC error above a certain limit (>5 
LSB TUE), out of spec current injection on adjacent pins or other functional failure (for 
example reset, oscillator frequency deviation). 
The test results are given in Table 45
         
Table 44. Electrical sensitivities
Symbol
Parameter
Conditions
Class
LU
Static latch-up class
TA = +105 °C conforming to JESD78A
II level A
Table 45. I/O current injection susceptibility 
Symbol
Description
Functional susceptibility 
Unit
Negative 
injection
Positive 
injection
IINJ
Injected current on OSC_IN32, 
OSC_OUT32, PA4, PA5, PC13 
-0
+0
mA
Injected current on all FT pins
-5
+0
Injected current on any other pin
-5
+5

---

---

**📄 Source: PDF Page 90**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
90/144
DocID14611 Rev 12
5.3.14 
I/O port characteristics
General input/output characteristics
Unless otherwise specified, the parameters given in Table 46 are derived from tests 
performed under the conditions summarized in Table 10. All I/Os are CMOS and TTL 
compliant.
         
All I/Os are CMOS and TTL compliant (no software configuration required). Their 
characteristics cover more than the strict CMOS-technology or TTL parameters. The 
coverage of these requirements is shown in Figure 42 and Figure 43 for standard I/Os, and 
in Figure 44 and Figure 45 for 5 V tolerant I/Os. 
Table 46. I/O static characteristics
Symbol
Parameter
Conditions
Min
Typ 
Max
Unit
VIL
Standard IO input low 
level voltage
-
–0.3
-
0.28*(VDD-2 V)+0.8 V
V
IO FT(1) input low level 
voltage
–0.3
-
0.32*(VDD-2 V)+0.75 V
V
VIH
Standard IO input high 
level voltage
-
0.41*(VDD-2 V)+1.3 
V
-
VDD+0.3
V
IO FT(1) input high level 
voltage
VDD > 2 V
0.42*(VDD-2 V)+1 V
-
5.5 
V
VDD ≤ 2 V
5.2
Vhys
Standard IO Schmitt 
trigger voltage 
hysteresis(2)
-
200
-
-
mV
IO FT Schmitt trigger 
voltage hysteresis(2)
5% VDD
(3)
-
-
mV
Ilkg
Input leakage current (4)
VSS ≤VIN ≤VDD
Standard I/Os
-
-
±1
µA
VIN= 5 V,
I/O FT
-
-
3
RPU
Weak pull-up equivalent 
resistor(5)
VIN = VSS
30
40
50
kΩ
RPD
Weak pull-down 
equivalent resistor(5)
VIN = VDD
30
40
50
kΩ
CIO
I/O pin capacitance
-
-
5
-
pF
1.
FT = Five-volt tolerant. In order to sustain a voltage higher than VDD+0.3 the internal pull-up/pull-down resistors must be 
disabled.
2.
Hysteresis voltage between Schmitt trigger switching levels. Guaranteed by characterization results.
3.
With a minimum of 100 mV.
4.
Leakage could be higher than max. if negative current is injected on adjacent pins.
5.
Pull-up and pull-down resistors are designed with a true resistance in series with a switchable PMOS/NMOS. This 
MOS/NMOS contribution to the series resistance is minimum (~10% order).

### Code Examples

```unknown
requirements is shown in Figure 42 and Figure 43 for standard I/Os, and
```

```unknown
required). Their
```

---

---

**📄 Source: PDF Page 91**

DocID14611 Rev 12
91/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 42. Standard I/O input characteristics - CMOS port
Figure 43. Standard I/O input characteristics - TTL port
Figure 44. 5 V tolerant I/O input characteristics - CMOS port
AIB
6$$ 6	




)NPUT RANGE NOT GUARANTEED



6)(6$$	


#-/3 STANDARD REQUIREMENT 6)(6$$ 
6)(6), 6	
#-/3 STANDARD REQUIREMENT 6),6$$








7),MAX
7)(MIN
6
$$	
6
),
AI


)NPUT RANGE NOT GUARANTEED
6)(6), 6	






44, REQUIREMENTS  6)(6
6)(6$$	
6),6$$	
44, REQUIREMENTS 6),6
6$$ 6	
7),MAX
7)(MIN
6$$



#-/3 STANDARD REQUIREMENTS 6)(6$$
 #-/3 STANDARD REQUIRMENT 6),6$$














6)(6), 6	
6$$ 6	
)NPUT RANGE NOT GUARANTEED
AIB
6)(6$$	
6),6$$	

---

---

**📄 Source: PDF Page 92**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
92/144
DocID14611 Rev 12
Figure 45. 5 V tolerant I/O input characteristics - TTL port
Output driving current
The GPIOs (general purpose input/outputs) can sink or source up to ±8 mA, and sink or 
source up to ± 20 mA (with a relaxed VOL/VOH) except PC13, PC14 and PC15 which can 
sink or source up to ±3 mA. When using the GPIOs PC13 to PC15 in output mode, the 
speed should not exceed 2 MHz with a maximum load of 30 pF.
In the user application, the number of I/O pins which can drive current must be limited to 
respect the absolute maximum rating specified in Section 5.2:
•
The sum of the currents sourced by all the I/Os on VDD, plus the maximum Run 
consumption of the MCU sourced on VDD, cannot exceed the absolute maximum rating 
IVDD (see Table 8). 
•
The sum of the currents sunk by all the I/Os on VSS plus the maximum Run 
consumption of the MCU sunk on VSS cannot exceed the absolute maximum rating 
IVSS (see Table 8). 
Output voltage levels
Unless otherwise specified, the parameters given in Table 47 are derived from tests 
performed under ambient temperature and VDD supply voltage conditions summarized in 
Table 10. All I/Os are CMOS and TTL compliant.
         





NOT GUARANTEED )NPUT RANGE



44, REQUIREMENT  6)(6
6)(
6$$	
6),
6$$	
44, REQUIREMENTS 6),6 6)(6), 6	
6$$ 6	
7),MAX
7)(MIN
AI
Table 47. Output voltage characteristics
Symbol
Parameter
Conditions
Min
Max
Unit
VOL
(1)
Output low level voltage for an I/O pin 
when 8 pins are sunk at same time
TTL port(3)
IIO = +8 mA
2.7 V < VDD < 3.6 V
-
0.4
V
VOH
(2)
Output high level voltage for an I/O pin 
when 8 pins are sourced at same time
VDD–0.4
-
VOL 
(1)
Output low level voltage for an I/O pin 
when 8 pins are sunk at same time
CMOS port(3)
IIO =+ 8mA
2.7 V < VDD < 3.6 V
-
0.4
V
VOH 
(2)
Output high level voltage for an I/O pin 
when 8 pins are sourced at same time
2.4
-

### Code Examples

```unknown
user application, the number of I/O pins which can drive current must be limited to
```

---

---

**📄 Source: PDF Page 93**

DocID14611 Rev 12
93/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
VOL
(1)(4)
Output low level voltage for an I/O pin 
when 8 pins are sunk at same time
IIO = +20 mA
2.7 V < VDD < 3.6 V
-
1.3
V
VOH
(2)(4) Output high level voltage for an I/O pin 
when 8 pins are sourced at same time
VDD–1.3
-
VOL
(1)(4)
Output low level voltage for an I/O pin 
when 8 pins are sunk at same time 
IIO = +6 mA
2 V < VDD < 2.7 V
-
0.4
V
VOH
(2)(4) Output high level voltage for an I/O pin 
when 8 pins are sourced at same time
VDD–0.4
-
1.
The IIO current sunk by the device must always respect the absolute maximum rating specified in Table 8 
and the sum of IIO (I/O ports and control pins) must not exceed IVSS.
2.
The IIO current sourced by the device must always respect the absolute maximum rating specified in 
Table 8 and the sum of IIO (I/O ports and control pins) must not exceed IVDD.
3.
TTL and CMOS outputs are compatible with JEDEC standards JESD36 and JESD52.
4.
Guaranteed by characterization results.
Table 47. Output voltage characteristics (continued)
Symbol
Parameter
Conditions
Min
Max
Unit

---

---

**📄 Source: PDF Page 94**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
94/144
DocID14611 Rev 12
Input/output AC characteristics
The definition and values of input/output AC characteristics are given in Figure 46 and 
Table 48, respectively.
Unless otherwise specified, the parameters given in Table 48 are derived from tests 
performed under ambient temperature and VDD supply voltage conditions summarized in 
Table 10. 
         
Table 48. I/O AC characteristics(1) 
1.
The I/O speed is configured using the MODEx[1:0] bits. Refer to the STM32F10xxx reference manual for a 
description of GPIO Port configuration register.
MODEx[1:0] 
bit value(1)
Symbol
Parameter
Conditions
Min
Max
Unit
10
fmax(IO)out Maximum frequency(2)
2.
The maximum frequency is defined in Figure 46.
CL = 50 pF, VDD = 2 V to 3.6 V
-
2
MHz
tf(IO)out
Output high to low 
level fall time
CL = 50 pF, VDD = 2 V to 3.6 V
-
125(3)
3.
Guaranteed by design.
ns
tr(IO)out
Output low to high 
level rise time
-
125(3)
01
fmax(IO)out Maximum frequency(2) CL = 50 pF, VDD = 2 V to 3.6 V
-
10
MHz
tf(IO)out
Output high to low 
level fall time
CL = 50 pF, VDD = 2 V to 3.6 V
-
25(3)
ns
tr(IO)out
Output low to high 
level rise time
-
25(3)
11
Fmax(IO)out Maximum frequency(2)
CL = 30 pF, VDD = 2.7 V to 3.6 V
-
50
MHz
CL = 50 pF, VDD = 2.7 V to 3.6 V
-
30
MHz
CL = 50 pF, VDD = 2 V to 2.7 V
-
20
MHz
tf(IO)out
Output high to low 
level fall time
CL = 30 pF, VDD = 2.7 V to 3.6 V
-
5(3)
ns
CL = 50 pF, VDD = 2.7 V to 3.6 V
-
8(3)
CL = 50 pF, VDD = 2 V to 2.7 V
-
12(3)
tr(IO)out
Output low to high 
level rise time
CL = 30 pF, VDD = 2.7 V to 3.6 V
-
5(3)
CL = 50 pF, VDD = 2.7 V to 3.6 V
-
8(3)
CL = 50 pF, VDD = 2 V to 2.7 V
-
12(3)
-
tEXTIpw
Pulse width of external 
signals detected by 
the EXTI controller
-
10
-
ns

---

---

**📄 Source: PDF Page 95**

DocID14611 Rev 12
95/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 46. I/O AC characteristics definition 
5.3.15 
NRST pin characteristics
The NRST pin input driver uses CMOS technology. It is connected to a permanent pull-up 
resistor, RPU (see Table 46).
Unless otherwise specified, the parameters given in Table 49 are derived from tests 
performed under ambient temperature and VDD supply voltage conditions summarized in 
Table 10. 
         
DLG



WU,2RXW
287387
(;7(51$/
21&/
0D[LPXPIUHTXHQF\LVDFKLHYHGLIWUWI7DQGLIWKHGXW\F\FOHLV
ZKHQORDGHGE\&/VSHFLILHGLQWKHWDEOH³,2$&FKDUDFWHULVWLFV´




7
WI,2RXW
Table 49. NRST pin characteristics 
Symbol
Parameter
Conditions
Min
Typ
Max
Unit
VIL(NRST)
(1)
1.
Guaranteed by design.
NRST Input low level voltage
-
–0.5
-
0.8
V
VIH(NRST)
(1)
NRST Input high level voltage
-
2
-
VDD+0.5
Vhys(NRST)
NRST Schmitt trigger voltage 
hysteresis 
-
-
200
-
mV
RPU
Weak pull-up equivalent resistor(2)
2.
The pull-up is designed with a true resistance in series with a switchable PMOS. This PMOS contribution 
to the series resistance must be minimum (~10% order).
VIN = VSS
30
40
50
kΩ
VF(NRST)
(1)
NRST Input filtered pulse
-
-
-
100
ns
VNF(NRST)
(1) NRST Input not filtered pulse
-
300
-
-
ns

### Code Examples

```unknown
uses CMOS technology. It is connected to a permanent pull-up
```

---

---

**📄 Source: PDF Page 96**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
96/144
DocID14611 Rev 12
Figure 47. Recommended NRST pin protection 
1.
The reset network protects the device against parasitic resets.
2.
The user must ensure that the level on the NRST pin can go below the VIL(NRST) max level specified in 
Table 49. Otherwise the reset will not be taken into account by the device.
5.3.16 
TIM timer characteristics
The parameters given in Table 50 are guaranteed by design.
Refer to Section 5.3.14: I/O port characteristics for details on the input/output alternate 
function characteristics (output compare, input capture, external clock, PWM output).
          
DLF
670)
538
1567
9''
)LOWHU
,QWHUQDO5HVHW
)
([WHUQDO
UHVHWFLUFXLW
Table 50. TIMx(1) characteristics
1.
TIMx is used as a general term to refer to the TIM1, TIM2, TIM3 and TIM4 timers.
Symbol
Parameter
Conditions
Min
Max
Unit
tres(TIM)
Timer resolution time
-
1
-
tTIMxCLK
 fTIMxCLK = 72 MHz
13.9
-
ns
fEXT
Timer external clock 
frequency on CH1 to CH4
 -
0
fTIMxCLK/2
MHz
fTIMxCLK = 72 MHz
0
36
MHz
ResTIM
Timer resolution
-
-
16
bit
tCOUNTER
16-bit counter clock period 
when internal clock is 
selected
-
1
65536
tTIMxCLK
 fTIMxCLK = 72 MHz
0.0139
910
µs
tMAX_COUNT
Maximum possible count
-
-
65536 × 65536
tTIMxCLK
 fTIMxCLK = 72 MHz
-
59.6
s

### Code Examples

```javascript
function characteristics (output compare, input capture, external clock, PWM output).
          
DLF
670)
538
1567
9''
)LOWHU
,QWHUQDO5HVHW
)
([WHUQDO
UHVHWFLUFXLW
Table 50. TIMx(1) characteristics
1.
TIMx is used as a general term to refer to the TIM1, TIM2, TIM3 and TIM4 timers.
Symbol
Parameter
Conditions
Min
Max
Unit
tres(TIM)
Timer resolution time
-
1
-
tTIMxCLK
 fTIMxCLK = 72 MHz
13.9
-
ns
fEXT
Timer external clock 
frequency on CH1 to CH4
 -
0
fTIMxCLK/2
MHz
fTIMxCLK = 72 MHz
0
36
MHz
ResTIM
Timer resolution
-
-
16
bit
tCOUNTER
16-bit counter clock period 
when internal clock is 
selected
-
1
65536
tTIMxCLK
 fTIMxCLK = 72 MHz
0.0139
910
µs
tMAX_COUNT
Maximum possible count
-
-
65536 × 65536
tTIMxCLK
 fTIMxCLK = 72 MHz
-
59.6
s
```

```typescript
used as a general term to refer to the TIM1, TIM2, TIM3 and TIM4 timers.
```

```unknown
user must ensure that the level on the NRST pin can go below the VIL(NRST) max level specified in
```

---

---

**📄 Source: PDF Page 97**

DocID14611 Rev 12
97/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
5.3.17 
Communications interfaces
I2C interface characteristics
The STM32F103xC, STM32F103xD and STM32F103xESTM32F103xF and STM32F103xG 
performance line I2C interface meets the requirements of the standard I2C communication 
protocol with the following restrictions: the I/O pins SDA and SCL are mapped to are not 
“true” open-drain. When configured as open-drain, the PMOS connected between the I/O 
pin and VDD is disabled, but is still present.
The I2C characteristics are described in Table 51. Refer also to Section 5.3.14: I/O port 
characteristics for more details on the input/output alternate function characteristics (SDA 
and SCL).
         
Table 51. I2C characteristics 
Symbol
Parameter
Standard mode 
I2C(1)(2)
1.
Guaranteed by design.
Fast mode I2C(1)(2)
2.
fPCLK1 must be at least 2 MHz to achieve standard mode I2C frequencies. It must be at least 4 MHz to 
achieve the fast mode I2C frequencies and it must be a multiple of 10 MHz in order to reach the I2C fast 
mode maximum clock speed of 400 kHz.
Unit
Min
Max
Min
Max
tw(SCLL)
SCL clock low time
4.7
-
1.3 
-
µs
tw(SCLH)
SCL clock high time
4.0
-
0.6 
-
tsu(SDA)
SDA setup time
250
-
100 
-
ns
th(SDA)
SDA data hold time
-
3450(3)
-
900(3)
3.
The device must internally provide a hold time of at least 300ns for the SDA signal in order to bridge the 
undefined region on the falling edge of SCL.
tr(SDA)
tr(SCL)
SDA and SCL rise time
-
1000
-
300
tf(SDA)
tf(SCL)
SDA and SCL fall time
-
300
-
300 
th(STA)
Start condition hold time
4.0
-
0.6
-
µs
tsu(STA)
Repeated Start condition 
setup time
4.7
-
0.6 
-
tsu(STO)
Stop condition setup time
4.0
-
0.6 
-
μs
tw(STO:STA)
Stop to Start condition time 
(bus free)
4.7
-
1.3
-
μs
Cb
Capacitive load for each bus 
line
-
400
-
400
pF
tSP
Pulse width of the spikes 
that are suppressed by the 
analog filter for standard and 
fast mode
0
50(4)
4.
The minimum width of the spikes filtered by the analog filter is above tSP(max).
0
50(4)
μs

### Code Examples

```javascript
function characteristics (SDA 
and SCL).
         
Table 51. I2C characteristics 
Symbol
Parameter
Standard mode 
I2C(1)(2)
1.
Guaranteed by design.
Fast mode I2C(1)(2)
2.
fPCLK1 must be at least 2 MHz to achieve standard mode I2C frequencies. It must be at least 4 MHz to 
achieve the fast mode I2C frequencies and it must be a multiple of 10 MHz in order to reach the I2C fast 
mode maximum clock speed of 400 kHz.
Unit
Min
Max
Min
Max
tw(SCLL)
SCL clock low time
4.7
-
1.3 
-
µs
tw(SCLH)
SCL clock high time
4.0
-
0.6 
-
tsu(SDA)
SDA setup time
250
-
100 
-
ns
th(SDA)
SDA data hold time
-
3450(3)
-
900(3)
3.
The device must internally provide a hold time of at least 300ns for the SDA signal in order to bridge the 
undefined region on the falling edge of SCL.
tr(SDA)
tr(SCL)
SDA and SCL rise time
-
1000
-
300
tf(SDA)
tf(SCL)
SDA and SCL fall time
-
300
-
300 
th(STA)
Start condition hold time
4.0
-
0.6
-
µs
tsu(STA)
Repeated Start condition 
setup time
4.7
-
0.6 
-
tsu(STO)
Stop condition setup time
4.0
-
0.6 
-
μs
tw(STO:STA)
Stop to Start condition time 
(bus free)
4.7
-
1.3
-
μs
Cb
Capacitive load for each bus 
line
-
400
-
400
pF
tSP
Pulse width of the spikes 
that are suppressed by the 
analog filter for standard and 
fast mode
0
50(4)
4.
The minimum width of the spikes filtered by the analog filter is above tSP(max).
0
50(4)
μs
```

```unknown
requirements of the standard I2C communication
used to design the application.
```

---

---

**📄 Source: PDF Page 98**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
98/144
DocID14611 Rev 12
Figure 48. I2C bus AC waveforms and measurement circuit
1.
Measurement points are done at CMOS levels: 0.3VDD and 0.7VDD.
2.
Rs: Series protection resistors.
3.
Rp: Pull-up resistors.
4.
VDD_I2C : I2C bus supply 
         
Table 52. SCL frequency (fPCLK1= 36 MHz.,VDD_I2C = 3.3 V)(1)(2)
1.
RP = External pull-up resistance, fSCL = I2C speed.
2.
For speeds around 200 kHz, the tolerance on the achieved speed is of ±5%. For other speed ranges, the 
tolerance on the achieved speed ±2%. These variations depend on the accuracy of the external 
components used to design the application.
fSCL (kHz)
I2C_CCR value
RP = 4.7 kΩ
400
0x801E
300
0x8028
200
0x803C
100
0x00B4
50
0x0168
20
0x0384
ĂŝϭϰϵϳϵĚ
^ dZ d
^ 
ZW
/ϸďƵƐ
sͺ/Ϯ
^dDϯϮ
^
^>
ƚĨ;^Ϳ
ƚƌ;^Ϳ
^>
ƚŚ;^dͿ
ƚǁ;^>,Ϳ
ƚǁ;^>>Ϳ
ƚƐƵ;^Ϳ
ƚƌ;^>Ϳ
ƚĨ;^>Ϳ
ƚŚ;^Ϳ
^ dZ dZWd
^ dZ d
ƚƐƵ;^dͿ
ƚƐƵ;^dKͿ
^ dKW
ƚǁ;^dK͗^dͿ
sͺ/Ϯ
ZW
Z^
Z^

---

---

**📄 Source: PDF Page 99**

DocID14611 Rev 12
99/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
I2S - SPI characteristics
Unless otherwise specified, the parameters given in Table 53 for SPI or in Table 54 for I2S 
are derived from tests performed under ambient temperature, fPCLKx frequency and VDD 
supply voltage conditions summarized in Table 10.
Refer to Section 5.3.14: I/O port characteristics for more details on the input/output alternate 
function characteristics (NSS, SCK, MOSI, MISO for SPI and WS, CK, SD for I2S).
         
Table 53. SPI characteristics
Symbol
Parameter
Conditions
Min
Max
Unit
fSCK
1/tc(SCK)
SPI clock frequency
Master mode
-
18
MHz
Slave mode
-
 18
tr(SCK)
tf(SCK)
SPI clock rise and fall 
time
Capacitive load: C = 30 pF
-
 8
ns
DuCy(SCK)
SPI slave input clock 
duty cycle
Slave mode
30
70
%
tsu(NSS)
(1)
1.
Guaranteed by characterization results.
NSS setup time 
Slave mode
4tPCLK
-
ns
th(NSS)
(1)
NSS hold time
Slave mode
2tPCLK
-
tw(SCKH)
(1)
tw(SCKL)
(1)
SCK high and low time
Master mode, fPCLK = 36 MHz, 
presc = 4
 50
60
tsu(MI) 
(1)
tsu(SI)
(1)
Data input setup time
Master mode
5
-
Slave mode
5
-
th(MI) 
(1)
Data input hold time
Master mode
5
-
th(SI)
(1)
Slave mode
4
-
ta(SO)
(1)(2)
2.
Min time is for the minimum time to drive the output and the max time is for the maximum time to validate 
the data.
Data output access time
Slave mode, fPCLK = 20 MHz
 0
3tPCLK
tdis(SO)
(1)(3)
3.
Min time is for the minimum time to invalidate the output and the max time is for the maximum time to put 
the data in Hi-Z
Data output disable time
Slave mode
2
10
tv(SO) 
(1)
Data output valid time
Slave mode (after enable edge)
-
25
tv(MO)
(1)
Data output valid time
Master mode (after enable edge)
-
5
th(SO)
(1)
Data output hold time
Slave mode (after enable edge)
15
-
th(MO)
(1)
Master mode (after enable edge)
2
-

### Code Examples

```typescript
function characteristics (NSS, SCK, MOSI, MISO for SPI and WS, CK, SD for I2S).
         
Table 53. SPI characteristics
Symbol
Parameter
Conditions
Min
Max
Unit
fSCK
1/tc(SCK)
SPI clock frequency
Master mode
-
18
MHz
Slave mode
-
 18
tr(SCK)
tf(SCK)
SPI clock rise and fall 
time
Capacitive load: C = 30 pF
-
 8
ns
DuCy(SCK)
SPI slave input clock 
duty cycle
Slave mode
30
70
%
tsu(NSS)
(1)
1.
Guaranteed by characterization results.
NSS setup time 
Slave mode
4tPCLK
-
ns
th(NSS)
(1)
NSS hold time
Slave mode
2tPCLK
-
tw(SCKH)
(1)
tw(SCKL)
(1)
SCK high and low time
Master mode, fPCLK = 36 MHz, 
presc = 4
 50
60
tsu(MI) 
(1)
tsu(SI)
(1)
Data input setup time
Master mode
5
-
Slave mode
5
-
th(MI) 
(1)
Data input hold time
Master mode
5
-
th(SI)
(1)
Slave mode
4
-
ta(SO)
(1)(2)
2.
Min time is for the minimum time to drive the output and the max time is for the maximum time to validate 
the data.
Data output access time
Slave mode, fPCLK = 20 MHz
 0
3tPCLK
tdis(SO)
(1)(3)
3.
Min time is for the minimum time to invalidate the output and the max time is for the maximum time to put 
the data in Hi-Z
Data output disable time
Slave mode
2
10
tv(SO) 
(1)
Data output valid time
Slave mode (after enable edge)
-
25
tv(MO)
(1)
Data output valid time
Master mode (after enable edge)
-
5
th(SO)
(1)
Data output hold time
Slave mode (after enable edge)
15
-
th(MO)
(1)
Master mode (after enable edge)
2
-
```

---

---

**📄 Source: PDF Page 100**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
100/144
DocID14611 Rev 12
Figure 49. SPI timing diagram - slave mode and CPHA = 0
Figure 50. SPI timing diagram - slave mode and CPHA = 1(1)
1.
Measurement points are done at CMOS levels: 0.3VDD and 0.7VDD.
DLF
6&.,QSXW
166LQSXW
W68166
WF6&.
WK166
&3+$ 
&32/ 
&3+$ 
&32/ 
WZ6&.+
WZ6&./
W962
WK62
WU6&.
WI6&.
WGLV62
WD62
0,62
287387
026,
,1387
06%287
%,7287
/6%287
WVX6,
WK6,
06%,1
%,7,1
/6%,1
DL
6&.,QSXW
&3+$ 
026,
,1387
0,62
287 387
&3+$ 
06% 2 87
0 6% ,1
%,7 287
/6% ,1
/6% 287
&32/ 
&32/ 
%,7 ,1
W68166
WF6&.
WK166
WD62
WZ6&.+
WZ6&./
WY62
WK62
WU6&.
WI6&.
WGLV62
WVX6,
WK6,
166LQSXW

---

---

**📄 Source: PDF Page 101**

DocID14611 Rev 12
101/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 51. SPI timing diagram - master mode(1)
1.
Measurement points are done at CMOS levels: 0.3VDD and 0.7VDD.
DLF
6&.2XWSXW
&3+$ 
026,
287387
0,62
,1387
&3+$ 
/6%287
/6%,1
&32/ 
&32/ 
%,7287
166LQSXW
WF6&.
WZ6&.+
WZ6&./
WU6&.
WI6&.
WK0,
+LJK
6&.2XWSXW
&3+$ 
&3+$ 
&32/ 
&32/ 
WVX0,
WY02
WK02
06%,1
%,7,1
06%287

---

---

**📄 Source: PDF Page 102**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
102/144
DocID14611 Rev 12
         
Table 54. I2S characteristics
Symbol
Parameter
Conditions
Min
Max
Unit
DuCy(SCK)
I2S slave input clock duty cycle
Slave mode
30
70
%
fCK
1/tc(CK)
I2S clock frequency
Master mode (data: 16 bits, 
Audio frequency = 48 kHz)
 1.522
 1.525
MHz
Slave mode
0
 6.5
tr(CK)
tf(CK)
I2S clock rise and fall time
Capacitive load CL = 50 pF
-
 8
ns
tv(WS) 
(1)
WS valid time
Master mode
3
-
th(WS) 
(1)
WS hold time
Master mode
I2S2
2
-
I2S3
0
-
tsu(WS) 
(1)
WS setup time
Slave mode
4
-
th(WS) 
(1)
WS hold time
Slave mode
0
-
tw(CKH) 
(1)
CK high and low time
Master fPCLK= 16 MHz, audio 
frequency = 48 kHz
 312.5
-
tw(CKL) 
(1)
 345
-
tsu(SD_MR) 
(1)
Data input setup time
Master receiver
I2S2
 2
-
I2S3
6.5
-
tsu(SD_SR) 
(1)
Data input setup time
Slave receiver
1.5
-
th(SD_MR)
(1)(2)
Data input hold time
Master receiver
 0
-
th(SD_SR) 
(1)(2)
Slave receiver
0.5
-
tv(SD_ST) 
(1)(2)
Data output valid time
Slave transmitter (after enable 
edge) 
-
 18
th(SD_ST) 
(1)
Data output hold time
Slave transmitter (after enable 
edge)
 11
-
tv(SD_MT) 
(1)(2)
Data output valid time
Master transmitter (after enable 
edge)
-
 3
th(SD_MT) 
(1)
Data output hold time
Master transmitter (after enable 
edge)
 0
-
1.
Guaranteed by design and/or characterization results.
2.
Depends on fPCLK. For example, if fPCLK=8 MHz, then TPCLK = 1/fPLCLK =125 ns.

---

---

**📄 Source: PDF Page 103**

DocID14611 Rev 12
103/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 52. I2S slave timing diagram (Philips protocol)(1) 
1.
Measurement points are done at CMOS levels: 0.3 × VDD and 0.7 × VDD.
2.
LSB transmit/receive of the previously transmitted byte. No LSB transmit/receive is sent before the first 
byte.
Figure 53. I2S master timing diagram (Philips protocol)(1)
1.
Guaranteed by characterization results.
2.
LSB transmit/receive of the previously transmitted byte. No LSB transmit/receive is sent before the first 
byte.
&.,QSXW
&32/ 
&32/ 
WF&.
:6LQSXW
6'WUDQVPLW
6'UHFHLYH
WZ&.+
WZ&./
WVX:6
WY6'B67
WK6'B67
WK:6
WVX6'B65
WK6'B65
06%UHFHLYH
%LWQUHFHLYH
/6%UHFHLYH
06%WUDQVPLW
%LWQWUDQVPLW
/6%WUDQVPLW
DLE
/6%UHFHLYH
/6%WUDQVPLW
#+ OUTPUT
#0/,  
#0/,  
TC#+	
73 OUTPUT
3$RECEIVE
3$TRANSMIT
TW#+(	
TW#+,	
TSU3$?-2	
TV3$?-4	
TH3$?-4	
TH73	
TH3$?-2	
-3" RECEIVE
"ITN RECEIVE
,3" RECEIVE
-3" TRANSMIT
"ITN TRANSMIT
,3" TRANSMIT
AIB
TF#+	
TR#+	
TV73	
,3" RECEIVE	
,3" TRANSMIT

---

---

**📄 Source: PDF Page 104**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
104/144
DocID14611 Rev 12
SD/SDIO MMC card host interface (SDIO) characteristics
Unless otherwise specified, the parameters given in Table 55 are derived from tests 
performed under ambient temperature, fPCLKx frequency and VDD supply voltage conditions 
summarized in Table 10.
Refer to Section 5.3.14: I/O port characteristics for more details on the input/output alternate 
function characteristics (D[7:0], CMD, CK).
Figure 54. SDIO high-speed mode
Figure 55. SD default mode
         
Table 55. SD / MMC characteristics
Symbol
Parameter
Conditions
Min
Max
Unit
fPP
Clock frequency in data transfer 
mode
CL ≤ 30 pF
0
48
MHz
tW(CKL)
Clock low time, fPP = 16 MHz
CL ≤ 30 pF
32
-
ns
tW(CKH)
Clock high time, fPP = 16 MHz
CL ≤ 30 pF
30
-
tr
Clock rise time
CL ≤ 30 pF
-
4
tf
Clock fall time
CL ≤ 30 pF
-
5
T7#+(	
#+
$ #-$
OUTPUT	
$ #-$
INPUT	
T#
T7#+,	
T/6
T/(
T)35
T)(
TF
TR
AI
AI
#+
$ #-$
OUTPUT	
T/6$
T/($

### Code Examples

```javascript
function characteristics (D[7:0], CMD, CK).
Figure 54. SDIO high-speed mode
Figure 55. SD default mode
         
Table 55. SD / MMC characteristics
Symbol
Parameter
Conditions
Min
Max
Unit
fPP
Clock frequency in data transfer 
mode
CL ≤ 30 pF
0
48
MHz
tW(CKL)
Clock low time, fPP = 16 MHz
CL ≤ 30 pF
32
-
ns
tW(CKH)
Clock high time, fPP = 16 MHz
CL ≤ 30 pF
30
-
tr
Clock rise time
CL ≤ 30 pF
-
4
tf
Clock fall time
CL ≤ 30 pF
-
5
T7#+(	
#+
$ #-$
OUTPUT	
$ #-$
INPUT	
T#
T7#+,	
T/6
T/(
T)35
T)(
TF
TR
AI
AI
#+
$ #-$
OUTPUT	
T/6$
T/($
```

---

---

**📄 Source: PDF Page 105**

DocID14611 Rev 12
105/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
USB characteristics
The USB interface is USB-IF certified (Full Speed).
         
CMD, D inputs (referenced to CK)
tISU
Input setup time
CL ≤ 30 pF
2
-
ns
tIH
Input hold time
CL ≤ 30 pF
0
-
CMD, D outputs (referenced to CK) in MMC and SD HS mode
tOV
Output valid time
CL ≤ 30 pF
-
6
ns
tOH
Output hold time
CL ≤ 30 pF
0
-
CMD, D outputs (referenced to CK) in SD default mode(1)
tOVD
Output valid default time
CL ≤ 30 pF
-
7
ns
tOHD
Output hold default time
CL ≤ 30 pF
0.5
-
1.
Refer to SDIO_CLKCR, the SDI clock control register to control the CK output.
Table 56. USB startup time
Symbol
Parameter
 Max
 Unit
tSTARTUP
(1)
1.
Guaranteed by design.
USB transceiver startup time
1
µs
Table 55. SD / MMC characteristics
Symbol
Parameter
Conditions
Min
Max
Unit

---

---

**📄 Source: PDF Page 106**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
106/144
DocID14611 Rev 12
         
Figure 56. USB timings: definition of data signal rise and fall time
         
Table 57. USB DC electrical characteristics
Symbol
Parameter
Conditions
Min.(1)
1.
All the voltages are measured from the local ground potential.
Max.(1)
Unit
Input levels
VDD
USB operating voltage(2)
2.
To be compliant with the USB 2.0 full-speed electrical specification, the USB_DP (D+) pin should be pulled 
up with a 1.5 kΩ resistor to a 3.0-to-3.6 V voltage range.
-
3.0(3)
3.
The STM32F103xC/D/E USB functionality is ensured down to 2.7 V but not the full USB electrical 
characteristics which are degraded in the 2.7-to-3.0 V VDD voltage range.
3.6
V
VDI
(4)
4.
Guaranteed by characterization results.
Differential input sensitivity
I(USB_DP, USB_DM) 
0.2
-
V
VCM
(4)
Differential common mode range
Includes VDI range
0.8
2.5
VSE
(4)
Single ended receiver threshold
-
1.3
2.0
Output levels
VOL
Static output level low
RL of 1.5 kΩ to 3.6 V(5)
5.
RL is the load connected on the USB drivers
-
0.3
V
VOH
Static output level high
RL of 15 kΩ to VSS
(5)
2.8
3.6
Table 58. USB: full-speed electrical characteristics
Driver characteristics(1)
1.
Guaranteed by design.
Symbol
Parameter
Conditions
Min
Max
Unit
tr
Rise time(2)
2.
Measured from 10% to 90% of the data signal. For more detailed informations, please refer to USB 
Specification - Chapter 7 (version 2.0).
CL = 50 pF 
4
20
ns
tf
Fall Time(2)
CL = 50 pF
4
20
ns
trfm
Rise/ fall time matching
tr/tf
90
110
%
VCRS
Output signal crossover voltage
-
1.3
2.0
V
DL
WI
66
WU
9
&56
9
'LIIHUHQWLDO
GDWDOLQHV
&URVVRYHU
SRLQWV

---

---

**📄 Source: PDF Page 107**

DocID14611 Rev 12
107/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
5.3.18 
CAN (controller area network) interface
Refer to Section 5.3.14: I/O port characteristics for more details on the input/output alternate 
function characteristics (CAN_TX and CAN_RX).
5.3.19 
12-bit ADC characteristics
Unless otherwise specified, the parameters given in Table 59 are preliminary values derived 
from tests performed under ambient temperature, fPCLK2 frequency and VDDA supply 
voltage conditions summarized in Table 10.
Note:
It is recommended to perform a calibration after each power-up.
         
Table 59. ADC characteristics
Symbol
Parameter
 Conditions
Min
Typ 
Max
Unit
VDDA
Power supply
-
2.4
-
3.6
V
VREF+
Positive reference voltage
-
2.4
-
VDDA
V
VREF-
Negative reference voltage
-
0
V
IVREF
Current on the VREF input 
pin
-
-
160(1)
220
µA
fADC
ADC clock frequency
-
0.6
-
14
MHz
fS
(2)
Sampling rate
-
0.05
-
1
MHz
fTRIG
(2)
External trigger frequency
fADC = 14 MHz
-
-
823
kHz
-
-
-
17
1/fADC
VAIN
Conversion voltage range(3)
-
0 (VSSA or VREF- 
tied to ground)
-
VREF+
V
RAIN
(2)
External input impedance
See Equation 1 and 
Table 60 for details
-
-
50
κΩ
RADC
(2)
Sampling switch resistance
-
-
-
1
κΩ
CADC
(2)
Internal sample and hold 
capacitor
-
-
-
8
pF
tCAL
(2)
Calibration time
fADC = 14 MHz
5.9
µs
-
83
1/fADC
tlat
(2)
Injection trigger conversion 
latency
fADC = 14 MHz
-
-
0.214
µs
-
-
-
3(4)
1/fADC
tlatr
(2)
Regular trigger conversion 
latency
fADC = 14 MHz
-
-
0.143
µs
-
-
-
2(4)
1/fADC
tS
(2)
Sampling time
fADC = 14 MHz
0.107
-
17.1
µs
-
1.5
-
239.5
1/fADC
tSTAB
(2)
Power-up time
-
0
0
1
µs
tCONV
(2)
Total conversion time 
(including sampling time)
fADC = 14 MHz
1
-
18
µs
-
14 to 252 (tS for sampling +12.5 for 
successive approximation)
1/fADC

### Code Examples

```sass
function characteristics (CAN_TX and CAN_RX).
5.3.19 
12-bit ADC characteristics
Unless otherwise specified, the parameters given in Table 59 are preliminary values derived 
from tests performed under ambient temperature, fPCLK2 frequency and VDDA supply 
voltage conditions summarized in Table 10.
Note:
It is recommended to perform a calibration after each power-up.
         
Table 59. ADC characteristics
Symbol
Parameter
 Conditions
Min
Typ 
Max
Unit
VDDA
Power supply
-
2.4
-
3.6
V
VREF+
Positive reference voltage
-
2.4
-
VDDA
V
VREF-
Negative reference voltage
-
0
V
IVREF
Current on the VREF input 
pin
-
-
160(1)
220
µA
fADC
ADC clock frequency
-
0.6
-
14
MHz
fS
(2)
Sampling rate
-
0.05
-
1
MHz
fTRIG
(2)
External trigger frequency
fADC = 14 MHz
-
-
823
kHz
-
-
-
17
1/fADC
VAIN
Conversion voltage range(3)
-
0 (VSSA or VREF- 
tied to ground)
-
VREF+
V
RAIN
(2)
External input impedance
See Equation 1 and 
Table 60 for details
-
-
50
κΩ
RADC
(2)
Sampling switch resistance
-
-
-
1
κΩ
CADC
(2)
Internal sample and hold 
capacitor
-
-
-
8
pF
tCAL
(2)
Calibration time
fADC = 14 MHz
5.9
µs
-
83
1/fADC
tlat
(2)
Injection trigger conversion 
latency
fADC = 14 MHz
-
-
0.214
µs
-
-
-
3(4)
1/fADC
tlatr
(2)
Regular trigger conversion 
latency
fADC = 14 MHz
-
-
0.143
µs
-
-
-
2(4)
1/fADC
tS
(2)
Sampling time
fADC = 14 MHz
0.107
-
17.1
µs
-
1.5
-
239.5
1/fADC
tSTAB
(2)
Power-up time
-
0
0
1
µs
tCONV
(2)
Total conversion time 
(including sampling time)
fADC = 14 MHz
1
-
18
µs
-
14 to 252 (tS for sampling +12.5 for 
successive approximation)
1/fADC
```

---

---

**📄 Source: PDF Page 108**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
108/144
DocID14611 Rev 12
Equation 1: RAIN max formula
The formula above (Equation 1) is used to determine the maximum external impedance 
allowed for an error below 1/4 of LSB. Here N = 12 (from 12-bit resolution).
         
         
1.
 Guaranteed by characterization results.
2.
Guaranteed by design.
3.
VREF+ can be internally connected to VDDA and VREF- can be internally connected to VSSA, depending on the package. 
Refer to Section 3: Pinouts and pin descriptions for further details.
4.
For external triggers, a delay of 1/fPCLK2 must be added to the latency specified in Table 59.
Table 60. RAIN max for fADC = 14 MHz(1)
1.
Guaranteed by design.
Ts (cycles)
tS (µs)
RAIN max (kΩ)
1.5
0.11
0.4
7.5
0.54
5.9
13.5
0.96
11.4
28.5
2.04
25.2
41.5
2.96
37.2
55.5
3.96
50
71.5
5.11
NA
239.5
17.1
NA
Table 61. ADC accuracy - limited test conditions(1)(2)
1.
ADC DC accuracy values are measured after internal calibration.
2.
ADC Accuracy vs. Negative Injection Current: Injecting negative current on any analog input pins should 
be avoided as this significantly reduces the accuracy of the conversion being performed on another analog 
input. It is recommended to add a Schottky diode (pin to ground) to analog pins which may potentially 
inject negative current. 
Any positive injection current within the limits specified for IINJ(PIN) and ΣIINJ(PIN) in Section 5.3.14 does not 
affect the ADC accuracy.
Symbol
Parameter
Test conditions
Typ
Max(3)
3.
Guaranteed by characterization results.
Unit
ET
Total unadjusted error
fPCLK2 = 56 MHz,
fADC = 14 MHz, RAIN < 10 kΩ,
VDDA = 3 V to 3.6 V
TA = 25 °C
Measurements made after 
ADC calibration
VREF+ = VDDA
±1.3
±2
LSB
EO
Offset error
±1
±1.5
EG
Gain error
±0.5
±1.5
ED
Differential linearity error
±0.7
±1
EL
Integral linearity error
±0.8
±1.5
RAIN
TS
fADC
CADC
2
N
2
+
(
)
ln
×
×
---------------------------------------------------------------
RADC
–
<

### Code Examples

```unknown
used to determine the maximum external impedance
```

---

---

**📄 Source: PDF Page 109**

DocID14611 Rev 12
109/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
         
Figure 57. ADC accuracy characteristics
1.
Example of an actual transfer curve.
2.
Ideal transfer curve.
3.
End point correlation line.
4.
ET = Total Unadjusted Error: maximum deviation between the actual and the ideal transfer curves.
EO = Offset Error: deviation between the first actual transition and the first ideal one.
EG = Gain Error: deviation between the last ideal transition and the last actual one.
ED = Differential Linearity Error: maximum deviation between actual steps and the ideal one.
EL = Integral Linearity Error: maximum deviation between any actual transition and the end point 
correlation line.
Table 62. ADC accuracy(1) (2)(3)
1.
ADC DC accuracy values are measured after internal calibration.
2.
Better performance could be achieved in restricted VDD, frequency, VREF and temperature ranges.
3.
ADC Accuracy vs. Negative Injection Current: Injecting negative current on any analog input pins should 
be avoided as this significantly reduces the accuracy of the conversion being performed on another analog 
input. It is recommended to add a Schottky diode (pin to ground) to analog pins which may potentially 
inject negative current. 
Any positive injection current within the limits specified for IINJ(PIN) and ΣIINJ(PIN) in Section 5.3.14 does not 
affect the ADC accuracy.
Symbol
Parameter
Test conditions
Typ
Max(4)
4.
Guaranteed by characterization results.
Unit
ET
Total unadjusted error
fPCLK2 = 56 MHz,
fADC = 14 MHz, RAIN < 10 kΩ,
VDDA = 2.4 V to 3.6 V
Measurements made after 
ADC calibration
±2
±5
LSB
EO
Offset error
±1.5
±2.5
EG
Gain error
±1.5
±3
ED
Differential linearity error
±1
±2
EL
Integral linearity error
±1.5
±3
AIC
%/
%'
, 3")$%!,
















   
	
	
%4
%$
%,
	
6$$!
633!
62%&

OR              DEPENDING ON PACKAGE	=
6$$!

;,3" )$%!,  

---

---

**📄 Source: PDF Page 110**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
110/144
DocID14611 Rev 12
Figure 58. Typical connection diagram using the ADC
1.
Refer to Table 59 for the values of RAIN, RADC and CADC.
2.
Cparasitic represents the capacitance of the PCB (dependent on soldering and PCB layout quality) plus the 
pad capacitance (roughly 7 pF). A high Cparasitic value will downgrade conversion accuracy. To remedy 
this, fADC should be reduced.
General PCB design guidelines
Power supply decoupling should be performed as shown in Figure 59 or Figure 60, 
depending on whether VREF+ is connected to VDDA or not. The 10 nF capacitors should be 
ceramic (good quality). They should be placed them as close as possible to the chip.
Figure 59. Power supply and reference decoupling (VREF+ not connected to VDDA)
1.
VREF+ and VREF– inputs are available only on 100-pin packages. 
DL
670)
9''
$,1[
,/$
9
97
5$,1
&SDUDVLWLF
9$,1
9
97
5$'&
&$'&
ELW
FRQYHUWHU
6DPSOHDQGKROG$'&
FRQYHUWHU
95()
VHHQRWH
670)[[
9''$
966$ 95()±
VHHQRWH
)Q)
)Q)
DLE

---

---

**📄 Source: PDF Page 111**

DocID14611 Rev 12
111/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 60. Power supply and reference decoupling (VREF+ connected to VDDA)
1.
VREF+ and VREF– inputs are available only on 100-pin packages. 
95()9''$
670)[[
)Q)
95()±966$
DL
6HHQRWH
6HHQRWH

---

---

**📄 Source: PDF Page 112**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
112/144
DocID14611 Rev 12
5.3.20 
DAC electrical specifications
         
Table 63. DAC characteristics
Symbol
Parameter
Min
Typ
Max
Unit
Comments
VDDA
Analog supply voltage
2.4
-
3.6 
V
-
VREF+
Reference supply voltage
2.4
-
3.6
V
VREF+ must always be below VDDA
VSSA
Ground
0
-
0
V
-
RLOAD
(1)
Resistive load with buffer ON 5
-
-
kΩ
-
RO
(2)
Impedance output with buffer 
OFF
-
-
15
kΩ
When the buffer is OFF, the Minimum 
resistive load between DAC_OUT 
and VSS to have a 1% accuracy is 
1.5 MΩ
CLOAD
(1) 
Capacitive load
-
-
50
pF
Maximum capacitive load at 
DAC_OUT pin (when the buffer is 
ON).
DAC_OUT 
min(1) 
Lower DAC_OUT voltage 
with buffer ON
0.2 
-
-
V
It gives the maximum output 
excursion of the DAC.
It corresponds to 12-bit input code 
(0x0E0) to (0xF1C) at VREF+ = 3.6 V 
and (0x155) and (0xEAB) at VREF+ = 
2.4 V
DAC_OUT 
max(1) 
Higher DAC_OUT voltage 
with buffer ON
-
-
VDDA – 0.2 
V
DAC_OUT 
min(1) 
Lower DAC_OUT voltage 
with buffer OFF
-
0.5
-
mV
It gives the maximum output 
excursion of the DAC.
DAC_OUT 
max(1)
Higher DAC_OUT voltage 
with buffer OFF
-
-
VREF+ – 1LSB V
IDDVREF+
DAC DC current 
consumption in quiescent 
mode (Standby mode)
-
-
220
µA
With no load, worst code (0xF1C) at 
VREF+ = 3.6 V in terms of DC 
consumption on the inputs
IDDA
DAC DC current 
consumption in quiescent 
mode(3)
-
-
380
µA
With no load, middle code (0x800) on 
the inputs
-
-
480
µA
With no load, worst code (0xF1C) at 
VREF+ = 3.6 V in terms of DC 
consumption on the inputs
DNL(4)
Differential non linearity 
Difference between two 
consecutive code-1LSB)
-
-
±0.5 
LSB Given for the DAC in 10-bit 
configuration
-
-
±2 
LSB Given for the DAC in 12-bit 
configuration 
INL(3)
Integral non linearity 
(difference between 
measured value at Code i 
and the value at Code i on a 
line drawn between Code 0 
and last Code 1023)
-
-
±1
LSB Given for the DAC in 10-bit 
configuration 
-
-
±4
LSB Given for the DAC in 12-bit 
configuration

---

---

**📄 Source: PDF Page 113**

DocID14611 Rev 12
113/144
STM32F103xC, STM32F103xD, STM32F103xE
Electrical characteristics
136
Figure 61. 12-bit buffered /non-buffered DAC
1.
The DAC integrates an output buffer that can be used to reduce the output impedance and to drive external 
loads directly without the use of an external operational amplifier. The buffer can be bypassed by 
configuring the BOFFx bit in the DAC_CR register.
Offset(3)
Offset error
(difference between 
measured value at Code 
(0x800) and the ideal value = 
VREF+/2)
-
-
±10
mV
-
-
-
±3
LSB Given for the DAC in 10-bit at VREF+ 
= 3.6 V
-
-
±12
LSB Given for the DAC in 12-bit at VREF+ 
= 3.6 V
Gain 
error(3)
Gain error
-
-
±0.5
%
Given for the DAC in 12bit 
configuration
tSETTLING
(3)
Settling time (full scale: for a 
10-bit input code transition 
between the lowest and the 
highest input codes when 
DAC_OUT reaches final 
value ±1LSB
-
3
4
µs
CLOAD ≤ 50 pF, RLOAD ≥ 5 kΩ
Update 
rate(3)
Max frequency for a correct 
DAC_OUT change when 
small variation in the input 
code (from code i to i+1LSB)
-
-
1
MS/s CLOAD ≤ 50 pF, RLOAD ≥ 5 kΩ
tWAKEUP
(3)
Wakeup time from off state 
(Setting the ENx bit in the 
DAC Control register)
-
6.5
10
µs
CLOAD ≤ 50 pF, RLOAD ≥ 5 kΩ
input code between lowest and 
highest possible ones.
PSRR+ (1) 
Power supply rejection ratio 
(to VDDA) (static DC 
measurement
-
–67 
–40
dB
No RLOAD, CLOAD = 50 pF
1.
Guaranteed by design.
2.
Guaranteed by characterization.
3.
The quiescent mode corresponds to a state where the DAC maintains a stable output level to ensure that no dynamic 
consumption occurs.
4.
Guaranteed by characterization results.
Table 63. DAC characteristics (continued)
Symbol
Parameter
Min
Typ
Max
Unit
Comments
5 /
& /
%XIIHUHG1RQEXIIHUHG'$&
'$&B287[
%XIIHU
ELW
GLJLWDOWR
DQDORJ
FRQYHUWHU
AI6

### Code Examples

```elixir
use of an external operational amplifier. The buffer can be bypassed by
```

```unknown
used to reduce the output impedance and to drive external
```

---

---

**📄 Source: PDF Page 114**

Electrical characteristics
STM32F103xC, STM32F103xD, STM32F103xE
114/144
DocID14611 Rev 12
5.3.21 
Temperature sensor characteristics
         
Table 64. TS characteristics
Symbol
Parameter
Min
Typ
Max
Unit
TL
VSENSE linearity with temperature
-
±1
±2
°C
Avg_Slope
Average slope
4.0
4.3
4.6
mV/°C
V25
Voltage at 25 °C
1.34
1.43
1.52
V
tSTART
(1)
1.
Guaranteed by design.
Startup time
4
-
10
µs
TS_temp
(2)(1)
2.
Shortest sampling time can be determined in the application by multiple iterations.
ADC sampling time when reading the 
temperature
-
-
17.1
µs

---

---

**📄 Source: PDF Page 115**

DocID14611 Rev 12
115/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
6 
Package information
In order to meet environmental requirements, ST offers these devices in different grades of 
ECOPACK® packages, depending on their level of environmental compliance. ECOPACK® 
specifications, grade definitions and product status are available at: www.st.com. 
ECOPACK® is an ST trademark.
6.1 
LFBGA144 package information
Figure 62. LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package outline
1.
Drawing is not to scale.
         
Table 65. LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Typ
Min
Max
A(2)
-
-
1.700
-
-
0.0669
A1
0.250
0.300
0.350
0.098
0.0118
0.0138
A2
0.810
0.910
1.010
0.0319
0.0358
0.0398
A3
0.225
0.26
0.295
0.0089
0.0102
0.0116
A4
0.585
0.650
0.715
0.0230
0.0256
0.0281
/)%*$B;B0(B9
6HDWLQJSODQH
$
H
)
)
(
0
EEDOOV
$
'
7239,(:
%277209,(:


H
$
$
$
%
&
GGG &
(
'
HHH
& $ %
III


0
0 &
$
$
$%$//
3$'&251(5
$%$//
3$'&251(5

### Code Examples

```unknown
requirements, ST offers these devices in different grades of
```

---

---

**📄 Source: PDF Page 116**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
116/144
DocID14611 Rev 12
Figure 63. LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package recommended footprint
         
b
0.350
0.400
0.450
0.0138
0.0157
0.0177
D
9.900
10.000
10.100
0.3898
0.3937
0.3976
D1
-
8.800
-
-
0.3465
-
E
9.900
10.000
10.100
0.3898
0.3937
0.3976
E1
-
8.800
-
-
0.3465
-
e
-
0.800
-
-
0.0315
-
F
-
0.600
-
-
0.0236
-
ddd
-
-
0.100
-
-
0.0039
eee
-
-
0.150
-
-
0.0059
fff
-
-
0.080
-
-
0.0031
1.
Values in inches are converted from mm and rounded to 4 decimal digits.
2.
STATSChipPAC package dimensions.
Table 66. LFBGA144 recommended PCB design rules (0.8 mm pitch BGA)
Dimension
Recommended values
Pitch
0.8 mm
Dpad
0.400 mm
UBM
0.350 mm
Table 65. LFBGA144 – 144-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package mechanical data (continued)
Symbol
millimeters
inches(1)
Min
Typ
Max
Typ
Min
Max
/)%*$B;B)3B9
'SDG
'VP

---

---

**📄 Source: PDF Page 117**

DocID14611 Rev 12
117/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
Device marking for LFBGA144 package
The following figure gives an example of topside marking orientation versus ball A1 identifier 
location. 
Figure 64. LFBGA144 marking example (package top view)
1.
Parts marked as “ES”, “E” or accompanied by an Engineering Sample notification letter, are not yet 
qualified and therefore not yet ready to be used in production and any consequences deriving from such 
usage will not be at ST charge. In no event, ST will be liable for any customer usage of these engineering 
samples in production. ST Quality has to be contacted prior to any decision to use these Engineering 
samples to run qualification activity.
Dsm
0.470 mm typ. (depends on the solder mask 
registration tolerance)
Stencil opening
0.400 mm
Stencil thickness
Between 0.100 mm to 0.125 mm
Pad trace width
0.120 mm
Ball Diameter
0.400 mm
Table 66. LFBGA144 recommended PCB design rules (0.8 mm pitch BGA) (continued)
Dimension
Recommended values
06Y9
%DOO
$LGHQWLILHU

670)
3URGXFWLGHQWLILFDWLRQ
=&+
'DWHFRGH
< ::

### Code Examples

```elixir
use these Engineering
```

```sql
used in production and any consequences deriving from such
```

---

---

**📄 Source: PDF Page 118**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
118/144
DocID14611 Rev 12
6.2 
LFBGA100 package information
Figure 65. LFBGA100 - 10 x 10 mm low profile fine pitch ball grid array package
outline
1.
Drawing is not to scale.
         
Table 67. LFBGA100 - 10 x 10 mm low profile fine pitch ball grid array package
mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
A
-
-
1.700
-
-
0.0669
A1
0.270
-
-
0.0106
-
-
A2
-
0.300
-
-
0.0118
-
A4
-
-
0.800
-
-
0.0315
b
0.450
0.500
0.550
0.0177
0.0197
0.0217
D
9.850
10.000
10.150
0.3878
0.3937
0.3996
D1
-
7.200
-
-
0.2835
-
E
9.850
10.000
10.150
0.3878
0.3937
0.3996
E1
-
7.200
-
-
0.2835
-
e
-
0.800
-
-
0.0315
-
F
-
1.400
-
-
0.0551
-
ddd
-
-
0.120
-
-
0.0047
+B0(B9
6HDWLQJSODQH
$
$
H
)
)
'
.
HHH
= < ;
III
EEDOOV


$
0
0
(
7239,(:
%277209,(:


H
$
$
=
<
;
=
GGG =
'
(
$EDOO
LGHQWLILHU
$EDOO
LQGH[DUHD

---

---

**📄 Source: PDF Page 119**

DocID14611 Rev 12
119/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
Figure 66. LFBGA100 – 100-ball low profile fine pitch ball grid array, 10 x 10 mm,
0.8 mm pitch, package recommended footprintoutline
eee
-
-
0.150
-
-
0.0059
fff
-
-
0.080
-
-
0.0031
1.
Values in inches are converted from mm and rounded to 4 decimal digits.
Table 68. LFBGA100 recommended PCB design rules (0.8 mm pitch BGA)
Dimension
Recommended values
Pitch
0.8
Dpad
0.500 mm
Dsm
0.570 mm typ. (depends on the soldermask reg-
istration tolerance)
Stencil opening
0.500 mm
Stencil thickness
Between 0.100 mm and 0.125 mm
Pad trace width
0.120 mm
Table 67. LFBGA100 - 10 x 10 mm low profile fine pitch ball grid array package
mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
+B)3B9
'SDG
'VP

---

---

**📄 Source: PDF Page 120**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
120/144
DocID14611 Rev 12
Device marking for LFBGA100 package
The following figure gives an example of topside marking orientation versus ball A1 identifier 
location. 
Figure 67. LFBGA100 marking example (package top view)
1.
Parts marked as “ES”, “E” or accompanied by an Engineering Sample notification letter, are not yet 
qualified and therefore not yet ready to be used in production and any consequences deriving from such 
usage will not be at ST charge. In no event, ST will be liable for any customer usage of these engineering 
samples in production. ST Quality has to be contacted prior to any decision to use these Engineering 
samples to run qualification activity.
06Y9
670)
9+
$GGLWLRQDO
LQIRUPDWLRQ
3URGXFWLGHQWLILFDWLRQ
'DWHFRGH
%DOO$
LQGHQWLILHU
< ::
;

### Code Examples

```elixir
use these Engineering
```

```sql
used in production and any consequences deriving from such
```

---

---

**📄 Source: PDF Page 121**

DocID14611 Rev 12
121/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
6.3 
WLCSP64 package information
Figure 68. WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm pitch, wafer-level chip-scale
package outline
1.
Drawing is not to scale.
2.
Primary datum Z and seating plane are defined by the spherical crowns of the ball.
         Table 69. WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm pitch, wafer-level chip-scale
package mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
A
0.535
0.585
0.635
0.0211
0.0230
0.0250
A1
0.205
0.230
0.255
0.0081
0.0091
0.0100
A2
0.330
0.355
0.380
0.0130
0.0140
0.0150
b(2)
0.290
0.320
0.350
0.0114
0.0126
0.0138
$
%XPS
HHH
'HWDLO$
URWDWHG
6HDWLQJSODQH
E
%XPSVLGH
H
H
H
H
*
)
&5B0(B9
6LGHYLHZ
'HWDLO$
$
$
:DIHUEDFNVLGH
'
(
*
)


$
+
=
EEE =
[
2ULHQWDWLRQ
UHIHUHQFH
$
FFF
GGG
; < =
=
=
DDD
;
<

---

---

**📄 Source: PDF Page 122**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
122/144
DocID14611 Rev 12
Figure 69. WLCSP64 - 64-ball, 4.4757 x 4.4049 mm, 0.5 mm pitch wafer level chip scale
package recommended footprint
         
e
-
0.500
-
-
0.0197
-
e1
-
3.500
-
-
0.1378
-
F
-
0.447
-
-
0.0176
-
G
-
0.483
-
-
0.0190
-
D
4.446
4.466
4.486
0.1750
0.1758
0.1766
E
4.375
4.395
4.415
0.1722
0.1730
0.1738
H
-
0.250
-
-
0.0098
-
L
-
0.200
-
-
0.0079
-
eee
-
0.05
-
-
0.0020
-
aaa
-
0.10
-
-
0.0039
-
Number of balls
64
1.
Values in inches are converted from mm and rounded to 4 decimal digits.
2.
Dimension is measured at the maximum ball diameter parallel to primary datum Z.
Table 70. WLCSP64 recommended PCB design rules (0.5 mm pitch)
Dimension
Recommended values
Pitch
0.5
Dpad
250 µm
Dsm
300 µm
Stencil Opening
320 µm
Stencil Thickness
Between 100 µm to 125 µm
Pad trace width
100 µm
Ball Diameter
320 µm
Table 69. WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm pitch, wafer-level chip-scale
package mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
:/&63B&5B)3B9
'SDG
'VP

---

---

**📄 Source: PDF Page 123**

DocID14611 Rev 12
123/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
6.4 
LQFP144 package information
Figure 70. LQFP144 - 144-pin, 20 x 20 mm low-profile quad flat package outline
1.
Drawing is not to scale. 
H
,'(17,),&$7,21
3,1
*$8*(3/$1(
PP
6($7,1*
3/$1(
'
'
'
(
(
(
.
FFF
&
&








$B0(B9
$
$
$
/
/
F
E
$

---

---

**📄 Source: PDF Page 124**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
124/144
DocID14611 Rev 12
         
Table 71. LQFP144 - 144-pin, 20 x 20 mm low-profile quad flat package 
mechanical data
Symbol
millimeters
inches(1)
1.
Values in inches are converted from mm and rounded to 4 decimal digits.
Min
Typ
Max
Min
Typ
Max
A
-
-
1.600
-
-
0.0630
A1
0.050
-
0.150
0.0020
-
0.0059
A2
1.350
1.400
1.450
0.0531
0.0551
0.0571
b
0.170
0.220
0.270
0.0067
0.0087
0.0106
c
0.090
-
0.200
0.0035
-
0.0079
D
21.800
22.000
22.200
0.8583
0.8661
0.8740
D1
19.800
20.000
20.200
0.7795
0.7874
0.7953
D3
-
17.500
-
-
0.6890
-
E
21.800
22.000
22.200
0.8583
0.8661
0.8740
E1
19.800
20.000
20.200
0.7795
0.7874
0.7953
E3
-
17.500
-
-
0.6890
-
e
-
0.500
-
-
0.0197
-
L
0.450
0.600
0.750
0.0177
0.0236
0.0295
L1
-
1.000
-
-
0.0394
-
k
0°
3.5°
7°
0°
3.5°
7°
ccc
-
-
0.080
-
-
0.0031

---

---

**📄 Source: PDF Page 125**

DocID14611 Rev 12
125/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
Figure 71. LQFP144 - 144-pin,20 x 20 mm low-profile quad flat package 
recommended footprint
1.
Dimensions are expressed in millimeters.








DLH









---

---

**📄 Source: PDF Page 126**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
126/144
DocID14611 Rev 12
Device marking for LQFP144 package
The following figure gives an example of topside marking orientation versus pin 1 identifier 
location. 
Figure 72. LQFP144 marking example (package top view) 
1.
Parts marked as “ES”, “E” or accompanied by an Engineering Sample notification letter, are not yet 
qualified and therefore not yet ready to be used in production and any consequences deriving from such 
usage will not be at ST charge. In no event, ST will be liable for any customer usage of these engineering 
samples in production. ST Quality has to be contacted prior to any decision to use these Engineering 
samples to run qualification activity.
06Y9
'DWHFRGH
3LQLGHQWLILHU
670)=(7
;
3URGXFWLGHQWLILFDWLRQ
5HYLVLRQFRGH
< ::
2SWLRQDOJDWHPDUN

### Code Examples

```elixir
use these Engineering
```

```sql
used in production and any consequences deriving from such
```

---

---

**📄 Source: PDF Page 127**

DocID14611 Rev 12
127/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
6.5 
LQFP100 package information
Figure 73. LQFP100 – 14 x 14 mm 100 pin low-profile quad flat package outline
1.
Drawing is not to scale.
         
E
)$%.4)&)#!4)/.
0). 
'!5'% 0,!.%
 MM
3%!4).' 0,!.%
$
$
$
%
%
%
+
CCC
#
#








,?-%?6
!
!
!
,
,
C
B
!
Table 72. LQPF100 – 14 x 14 mm 100-pin low-profile quad flat package 
mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
A
-
-
1.600
-
-
0.0630
A1
0.050
-
0.150
0.0020
-
0.0059
A2
1.350
1.400
1.450
0.0531
0.0551
0.0571
b
0.170
0.220
0.270
0.0067
0.0087
0.0106
c
0.090
-
0.200
0.0035
-
0.0079
D
15.800
16.000
16.200
0.6220
0.6299
0.6378
D1
13.800
14.000
14.200
0.5433
0.5512
0.5591
D3
-
12.000
-
-
0.4724
-
E
15.800
16.000
16.200
0.6220
0.6299
0.6378
E1
13.800
14.000
14.200
0.5433
0.5512
0.5591
E3
-
12.000
-
-
0.4724
-

---

---

**📄 Source: PDF Page 128**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
128/144
DocID14611 Rev 12
Figure 74. LQFP100 recommended footprint
1.
Dimensions are in millimeters.
e
-
0.500
-
-
0.0197
-
L
0.450
0.600
0.750
0.0177
0.0236
0.0295
L1
-
1.000
-
-
0.0394
-
k
0°
3.5°
7°
0°
3.5°
7°
ccc
-
-
0.08
-
-
0.0031
1.
Values in inches are converted from mm and rounded to 4 decimal digits.
Table 72. LQPF100 – 14 x 14 mm 100-pin low-profile quad flat package 
mechanical data (continued)
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max















AIC

---

---

**📄 Source: PDF Page 129**

DocID14611 Rev 12
129/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
Device marking for LQFP100 package
The following figure gives an example of topside marking orientation versus pin 1 identifier 
location. 
Figure 75. LQFP100 marking example (package top view) 
1.
Parts marked as “ES”, “E” or accompanied by an Engineering Sample notification letter, are not yet 
qualified and therefore not yet ready to be used in production and any consequences deriving from such 
usage will not be at ST charge. In no event, ST will be liable for any customer usage of these engineering 
samples in production. ST Quality has to be contacted prior to any decision to use these Engineering 
samples to run qualification activity.
06Y9
^dDϯϮ&ϭϬϯ
sϴdϲ
y
3URGXFWLGHQWLILFDWLRQ
5HYLVLRQFRGH
tt
z
'DWHFRGH
2SWLRQDOJDWH
PDUN
3LQLGHQWLILHU

### Code Examples

```elixir
use these Engineering
```

```sql
used in production and any consequences deriving from such
```

---

---

**📄 Source: PDF Page 130**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
130/144
DocID14611 Rev 12
6.6 
LQFP64 package information
Figure 76. LQFP64 – 10 x 10 mm 64 pin low-profile quad flat package outline
1.
Drawing is not in scale.
         
:B0(B9
$
$
$
6($7,1*3/$1(
FFF
&
E
&
F
$
/
/
.
,'(17,),&$7,21
3,1
'
'
'
H








(
(
(
*$8*(3/$1(
PP
Table 73. LQFP64 – 10 x 10 mm 64 pin low-profile quad flat package mechanical data
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
A
 -
-
1.600
-
-
0.0630
A1
0.050
-
0.150
0.0020
-
0.0059
A2
1.350
1.400
1.450
0.0531
0.0551
0.0571
b
0.170
0.220
0.270
0.0067
0.0087
0.0106
c
0.090
-
0.200
0.0035
-
0.0079
D
-
12.000
-
-
0.4724
-
D1
-
10.000
-
-
0.3937
-
D3
-
7.500
-
-
0.2953
-
E
-
12.000
-
-
0.4724
-
E1
-
10.000
-
-
0.3937
-

---

---

**📄 Source: PDF Page 131**

DocID14611 Rev 12
131/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
Figure 77. LQFP64 - 64-pin, 10 x 10 mm low-profile quad flat recommended footprint
1.
Dimensions are in millimeters.
E3
-
7.500
-
-
0.2953
-
e
-
0.500
-
-
0.0197
-
θ
0°
3.5°
7°
0°
3.5°
7°
L
0.450
0.600
0.750
0.0177
0.0236
0.0295
L1
-
1.000
-
-
0.0394
-
ccc
-
-
0.080
-
-
0.0031
1.
Values in inches are converted from mm and rounded to 4 decimal digits.
Table 73. LQFP64 – 10 x 10 mm 64 pin low-profile quad flat package mechanical data (continued)
Symbol
millimeters
inches(1)
Min
Typ
Max
Min
Typ
Max
















AIC

---

---

**📄 Source: PDF Page 132**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
132/144
DocID14611 Rev 12
Device marking for LQFP64 package
The following figure gives an example of topside marking orientation versus pin 1 identifier 
location. 
Figure 78. LQFP64 marking example (package top view) 
1.
Parts marked as “ES”, “E” or accompanied by an Engineering Sample notification letter, are not yet 
qualified and therefore not yet ready to be used in production and any consequences deriving from such 
usage will not be at ST charge. In no event, ST will be liable for any customer usage of these engineering 
samples in production. ST Quality has to be contacted prior to any decision to use these Engineering 
samples to run qualification activity.
         
06Y9
^dDϯϮ&ϭϬϯ

z tt
5HYLVLRQFRGH
'DWHFRGH
3LQLGHQWLILHU
3URGXFWLGHQWLILFDWLRQ
Zdϲ

### Code Examples

```elixir
use these Engineering
```

```sql
used in production and any consequences deriving from such
```

---

---

**📄 Source: PDF Page 133**

DocID14611 Rev 12
133/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
6.7 
Thermal characteristics
The maximum chip junction temperature (TJmax) must never exceed the values given in 
Table 10: General operating conditions on page 44.
The maximum chip-junction temperature, TJ max, in degrees Celsius, may be calculated 
using the following equation:
TJ max = TA max + (PD max x ΘJA)
Where:
•
TA max is the maximum ambient temperature in °C,
•
ΘJA is the package junction-to-ambient thermal resistance, in °C/W,
•
PD max is the sum of PINT max and PI/O max (PD max = PINT max + PI/Omax),
•
PINT max is the product of IDD and VDD, expressed in Watts. This is the maximum chip 
internal power.
PI/O max represents the maximum power dissipation on output pins where:
PI/O max = Σ (VOL × IOL) + Σ((VDD – VOH) × IOH),
taking into account the actual VOL / IOL and VOH / IOH of the I/Os at low and high level in the 
application.
         
6.7.1 
Reference document
JESD51-2 Integrated Circuits Thermal Test Method Environment Conditions - Natural 
Convection (Still Air). Available from www.jedec.org
Table 74. Package thermal characteristics
Symbol
Parameter
Value
Unit
ΘJA
Thermal resistance junction-ambient
LFBGA144 - 10 × 10 mm / 0.8 mm pitch
40
°C/W
Thermal resistance junction-ambient
LQFP144 - 20 × 20 mm / 0.5 mm pitch
30
Thermal resistance junction-ambient
LFBGA100 - 10 × 10 mm / 0.8 mm pitch
40
Thermal resistance junction-ambient
LQFP100 - 14 × 14 mm / 0.5 mm pitch
46
Thermal resistance junction-ambient
LQFP64 - 10 × 10 mm / 0.5 mm pitch
45
Thermal resistance junction-ambient
WLCSP64
50

---

---

**📄 Source: PDF Page 134**

Package information
STM32F103xC, STM32F103xD, STM32F103xE
134/144
DocID14611 Rev 12
6.7.2 
Selecting the product temperature range
When ordering the microcontroller, the temperature range is specified in the ordering 
information scheme shown in Table 75: Ordering information scheme.
Each temperature range suffix corresponds to a specific guaranteed ambient temperature at 
maximum dissipation and, to a specific maximum junction temperature.
As applications do not commonly use the STM32F103xC, STM32F103xD and 
STM32F103xE at maximum dissipation, it is useful to calculate the exact power 
consumption and junction temperature to determine which temperature range will be best 
suited to the application.
The following examples show how to calculate the temperature range needed for a given 
application.
Example 1: High-performance application
Assuming the following application conditions:
Maximum ambient temperature TAmax = 82 °C (measured according to JESD51-2), 
IDDmax = 50 mA, VDD = 3.5 V, maximum 20 I/Os used at the same time in output at low 
level with IOL = 8 mA, VOL= 0.4 V and maximum 8 I/Os used at the same time in output 
at low level with IOL = 20 mA, VOL= 1.3 V
PINTmax = 50 mA × 3.5 V= 175 mW
PIOmax = 20 × 8 mA × 0.4 V + 8 × 20 mA × 1.3 V = 272 mW
This gives: PINTmax = 175 mW and PIOmax = 272 mW:
PDmax = 175 + 272 = 447 mW
Thus: PDmax = 447 mW
Using the values obtained in Table 74 TJmax is calculated as follows:
–
For LQFP100, 46 °C/W 
TJmax = 82 °C + (46 °C/W × 447 mW) = 82 °C + 20.6 °C = 102.6 °C
This is within the range of the suffix 6 version parts (–40 < TJ < 105 °C).
In this case, parts must be ordered at least with the temperature range suffix 6 (see 
Table 75: Ordering information scheme).
Example 2: High-temperature application
Using the same rules, it is possible to address applications that run at high ambient 
temperatures with a low dissipation, as long as junction temperature TJ remains within the 
specified range.
Assuming the following application conditions:
Maximum ambient temperature TAmax = 115 °C (measured according to JESD51-2), 
IDDmax = 20 mA, VDD = 3.5 V, maximum 20 I/Os used at the same time in output at low 
level with IOL = 8 mA, VOL= 0.4 V
PINTmax = 20 mA × 3.5 V= 70 mW
PIOmax = 20 × 8 mA × 0.4 V = 64 mW
This gives: PINTmax = 70 mW and PIOmax = 64 mW:
PDmax = 70 + 64 = 134 mW
Thus: PDmax = 134 mW

### Code Examples

```elixir
use the STM32F103xC, STM32F103xD and
```

```unknown
useful to calculate the exact power
```

```unknown
used at the same time in output at low
```

```unknown
used at the same time in output
```

---

---

**📄 Source: PDF Page 135**

DocID14611 Rev 12
135/144
STM32F103xC, STM32F103xD, STM32F103xE
Package information
136
Using the values obtained in Table 74 TJmax is calculated as follows:
–
For LQFP100, 46 °C/W 
TJmax = 115 °C + (46 °C/W × 134 mW) = 115 °C + 6.2 °C = 121.2 °C
This is within the range of the suffix 7 version parts (–40 < TJ < 125 °C).
In this case, parts must be ordered at least with the temperature range suffix 7 (see 
Table 75: Ordering information scheme).
Figure 79. LQFP100 PD max vs. TA
0
100
200
300
400
500
600
700
65
75
85
95
105
115
125
135
TA (°C)
PD (mW)
Suffix 6
Suffix 7

---

---

**📄 Source: PDF Page 136**

Part numbering
STM32F103xC, STM32F103xD, STM32F103xE
136/144
DocID14611 Rev 12
7 
Part numbering
         
For a list of available options (speed, package, etc.) or for further information on any aspect 
of this device, please contact your nearest ST sales office.
Table 75. Ordering information scheme
Example:
STM32
F 103 R
C
T
6
xxx
Device family
STM32 = ARM-based 32-bit microcontroller
Product type
F = general-purpose
Device subfamily
103 = performance line
Pin count
R = 64 pins
V = 100 pins
Z = 144 pins
Flash memory size
C = 256 Kbytes of Flash memory
D = 384 Kbytes of Flash memory
E = 512 Kbytes of Flash memory
Package
H = BGA
T = LQFP
Y = WLCSP64
Temperature range
6 = Industrial temperature range, –40 to 85 °C.
7 = Industrial temperature range, –40 to 105 °C.
Options
xxx = programmed parts
TR = tape and real

---

---

**📄 Source: PDF Page 137**

DocID14611 Rev 12
137/144
STM32F103xC, STM32F103xD, STM32F103xE
Revision history
143
8 
Revision history
         
Table 76.Document revision history
Date
Revision
Changes
07-Apr-2008
1
Initial release.
22-May-2008
2
Document status promoted from Target Specification to Preliminary 
Data.
Section 1: Introduction and Section 2.2: Full compatibility throughout 
the family modified. Small text changes.
Note 2 added in Table 2: STM32F103xC, STM32F103xD and 
STM32F103xE features and peripheral counts on page 11.
LQPF100/BGA100 column added to Table 6: FSMC pin definition on 
page 38.
Values and Figures added to Maximum current consumption on 
page 62 (see Table 18, Table 19, Table 20 and Table 21 and see 
Figure 14, Figure 15, Figure 17, Figure 18 and Figure 19).
Values added to Typical current consumption on page 73 (see 
Table 22, Table 23 and Table 24). Table 19: Typical current 
consumption in Standby mode removed.
Note 4 and Note 1 added to Table 65: USB DC electrical characteristics 
and Table 66: USB: full-speed electrical characteristics on page 129, 
respectively.
VUSB added to Table 65: USB DC electrical characteristics on 
page 129.
Figure 68: Recommended footprint(1) on page 143 corrected.
Equation 1 corrected. Figure 73: LQFP100 PD max vs. TA on page 149 
modified.
Tolerance values corrected in Table 74: LFBGA144 – 144-ball low 
profile fine pitch ball grid array, 10 x 10 mm, 0.8 mm pitch, package 
data on page 139.

---

---

**📄 Source: PDF Page 138**

Revision history
STM32F103xC, STM32F103xD, STM32F103xE
138/144
DocID14611 Rev 12
21-Jul-2008
3
Document status promoted from Preliminary Data to full datasheet.
FSMC (flexible static memory controller) on page 22 modified.
Number of complementary channels corrected in Figure 1: 
STM32F103xF, STM32F103xD and STM32F103xGSTM32F103xF and 
STM32F103xG performance line block diagram.
Power supply supervisor on page 23 modified and VDDA added to 
Table 14: General operating conditions on page 59.
Table notes revised in Section 5: Electrical characteristics.
Capacitance modified in Figure 12: Power supply scheme on page 57.
Table 60: SCL frequency (fPCLK1= 36 MHz.,VDD = 3.3 V) updated.
Table 61: SPI characteristics modified, th(NSS) modified in Figure 49: 
SPI timing diagram - slave mode and CPHA = 0 on page 123.
Minimum SDA and SCL fall time value for Fast mode removed from 
Table 59: I2C characteristics on page 120, note 1 modified.
IDD_VBAT values and some IDD values with regulator in run mode added 
to Table 21: Typical and maximum current consumptions in Stop and 
Standby modes on page 68.
Table 34: Flash memory endurance and data retention on page 87 
updated.
tsu(NSS) modified in Table 61: SPI characteristics on page 122.
EO corrected in Table 70: ADC accuracy on page 132. Figure 58: 
Typical connection diagram using the ADC on page 133 and note below 
corrected.
Typical TS_temp value removed from Table 72: TS characteristics on 
page 137.
Section 6.1: Package mechanical data on page 138 updated.
Small text changes.
Table 76.Document revision history
Date
Revision
Changes

---

---

**📄 Source: PDF Page 139**

DocID14611 Rev 12
139/144
STM32F103xC, STM32F103xD, STM32F103xE
Revision history
143
12-Dec-2008
4
Timers specified on page 1 (motor control capability mentioned).
Section 2.2: Full compatibility throughout the family updated.
Table 6: High-density timer feature comparison added.
General-purpose timers (TIMx) and Advanced-control timers (TIM1 and 
TIM8) on page 27 updated.
Figure 1: STM32F103xF, STM32F103xD and 
STM32F103xGSTM32F103xF and STM32F103xG performance line 
block diagram modified.
Note 10 added, main function after reset and Note 5 on page 44 
updated in Table 8: High-density STM32F103xx pin definitions.
Note 2 modified below Table 11: Voltage characteristics on page 58, 
|DVDDx| min and |DVDDx| min removed.
Note 2 and PD values for LQFP144 and LFBGA144 packages added to 
Table 14: General operating conditions on page 59.
Measurement conditions specified in Section 5.3.5: Supply current 
characteristics on page 62.
Max values at TA = 85 °C and TA = 105 °C updated in Table 21: Typical 
and maximum current consumptions in Stop and Standby modes on 
page 68.
Section 5.3.10: FSMC characteristics on page 87 updated.
Data added to Table 50: EMI characteristics on page 111.
IVREF added to Table 67: ADC characteristics on page 130.
Table 81: Package thermal characteristics on page 146 updated.
Small text changes.
Table 76.Document revision history
Date
Revision
Changes

---

---

**📄 Source: PDF Page 140**

Revision history
STM32F103xC, STM32F103xD, STM32F103xE
140/144
DocID14611 Rev 12
30-Mar-2009
5
I/O information clarified on page 1. Figure 4: STM32F103xC and 
STM32F103xE performance line BGA100 ballout corrected.
I/O information clarified  on page 1.
In Table 5: High-density STM32F103xx pin definitions:
– I/O level of pins PF11, PF12, PF13, PF14, PF15, G0, G1 and G15 
updated
– PB4, PB13, PB14, PB15, PB3/TRACESWO moved from Default 
column to Remap column
PG14 pin description modified in Table 6: FSMC pin definition.
Figure 9: Memory map on page 54 modified.
Note modified in Table 18: Maximum current consumption in Run 
mode, code with data processing  running from Flash and Table 20: 
Maximum current consumption in Sleep mode, code running from Flash 
or RAM.
Figure 17, Figure 18 and Figure 19 show typical curves (titles 
changed).
Table 25: High-speed external user clock characteristics and Table 26: 
Low-speed external user clock characteristics modified. ACCHSI max 
values modified in Table 29: HSI oscillator characteristics.
FSMC configuration modified for Asynchronous waveforms and 
timings. Notes modified below Figure 24: Asynchronous non-
multiplexed SRAM/PSRAM/NOR read waveforms and Figure 25: 
Asynchronous non-multiplexed SRAM/PSRAM/NOR write waveforms.
tw(NADV) values modified in Table 35: Asynchronous non-multiplexed 
SRAM/PSRAM/NOR read timings and Table 39: Asynchronous 
multiplexed PSRAM/NOR write timings. th(Data_NWE) modified in 
Table 36: Asynchronous non-multiplexed SRAM/PSRAM/NOR write 
timings
In Table 41: Synchronous multiplexed PSRAM write timings and 
Table 43: Synchronous non-multiplexed PSRAM write timings:
– tv(Data-CLK) renamed as td(CLKL-Data)
– td(CLKL-Data) min value removed and max value added
– th(CLKL-DV) / th(CLKL-ADV) removed
Figure 28: Synchronous multiplexed NOR/PSRAM read timings, 
Figure 29: Synchronous multiplexed PSRAM write timings and 
Figure 31: Synchronous non-multiplexed PSRAM write timings 
modified.
Figure 52: I2S slave timing diagram (Philips protocol)(1) and Figure 53: 
I2S master timing diagram (Philips protocol)(1) modified.
WLCSP64 package added (see Figure 8: STM32F103xC and 
STM32F103xE performance line WLCSP64 ballout, ball side, Table 8: 
High-density STM32F103xx pin definitions, Figure 65: WLCSP, 64-ball 
4.466 × 4.395 mm, 0.500 mm pitch, wafer-level chip-scale package 
outline and Table 76: WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm 
pitch, wafer-level chip-scale package mechanical data).
Small text changes.
Table 76.Document revision history
Date
Revision
Changes

### Code Examples

```unknown
user clock characteristics and Table 26:
```

```unknown
user clock characteristics modified. ACCHSI max
```

---

---

**📄 Source: PDF Page 141**

DocID14611 Rev 12
141/144
STM32F103xC, STM32F103xD, STM32F103xE
Revision history
143
21-Jul-2009
6
Figure 1: STM32F103xC, STM32F103xD and STM32F103xE 
performance line block diagram updated.
Note 5 updated and Note 4 added in Table 5: High-density 
STM32F103xC/D/E pin definitions.
VRERINT and TCoeff added to Table 13: Embedded internal reference 
voltage.
Table 16: Maximum current consumption in Sleep mode, code running 
from Flash or RAM modified.
fHSE_ext min modified in Table 21: High-speed external user clock 
characteristics.
CL1 and CL2 replaced by C in Table 23: HSE 4-16 MHz oscillator 
characteristics and Table 24: LSE oscillator characteristics (fLSE = 
32.768 kHz), notes modified and moved below the tables.
Note 1 modified below Figure 29: Synchronous multiplexed PSRAM 
write timings. Table 25: HSI oscillator characteristics modified. 
Conditions removed from Table 27: Low-power mode wakeup timings.
Jitter added to Table 28: PLL characteristics.
Figure 47: Recommended NRST pin protection modified.
In Table 31: Asynchronous non-multiplexed SRAM/PSRAM/NOR read 
timings: th(BL_NOE) and th(A_NOE) modified.
In Table 32: Asynchronous non-multiplexed SRAM/PSRAM/NOR write 
timings: th(A_NWE) and th(Data_NWE) modified.
In Table 33: Asynchronous multiplexed PSRAM/NOR read timings: 
th(AD_NADV) and th(A_NOE) modified.
In Table 34: Asynchronous multiplexed PSRAM/NOR write timings: 
th(A_NWE) modified.
In Table 35: Synchronous multiplexed NOR/PSRAM read timings: 
th(CLKH-NWAITV) modified.
In Table 40: Switching characteristics for NAND Flash read and write 
cycles: th(NOE-D) modified.
Table 53: SPI characteristics modified. Values added to Table 54: I2S 
characteristics and Table 55: SD / MMC characteristics.
CADC and RAIN parameters modified in Table 59: ADC characteristics. 
RAIN max values modified in Table 60: RAIN max for fADC = 14 MHz.
Table 71: DAC characteristics modified. Figure 61: 12-bit buffered /non-
buffered DAC added.
Figure 63: LFBGA100 - 10 x 10 mm low profile fine pitch ball grid array 
package outline and Table 75: LFBGA100 - 10 x 10 mm low profile fine 
pitch ball grid array package mechanical data updated.
24-Sep-2009
7
Number of DACs corrected in Table 3: STM32F103xx family.
IDD_VBAT updated in Table 17: Typical and maximum current 
consumptions in Stop and Standby modes.
Figure 16: Typical current consumption on VBAT with RTC on vs. 
temperature  at different VBAT values added.
IEC 1000 standard updated to IEC 61000 and SAE J1752/3 updated to 
IEC 61967-2 in Section 5.3.11: EMC characteristics on page 87.
Table 63: DAC characteristics modified. Small text changes.
Table 76.Document revision history
Date
Revision
Changes

---

---

**📄 Source: PDF Page 142**

Revision history
STM32F103xC, STM32F103xD, STM32F103xE
142/144
DocID14611 Rev 12
19-Apr-2011
8
Updated package choice for 103Rx in Table 2
Updated footnotes below Table 7: Voltage characteristics on page 43 
and Table 8: Current characteristics on page 43
Updated tw min in Table 21: High-speed external user clock 
characteristics on page 58
Updated startup time in Table 24: LSE oscillator characteristics (fLSE = 
32.768 kHz) on page 61
Updated note 2 in Table 51: I2C characteristics on page 97
Updated Figure 48: I2C bus AC waveforms and measurement circuit
Updated Figure 47: Recommended NRST pin protection
Updated Section 5.3.14: I/O port characteristics
Updated Table 35: Synchronous multiplexed NOR/PSRAM read 
timings on page 73
Updated FSMC Figure 26 thru Figure 31
Updated Figure 41.: NAND controller waveforms for common memory 
write access and Figure 48.: I2C bus AC waveforms and measurement 
circuit
Added Section 5.3.13: I/O current injection characteristics
Updated Figure 67 and added Table 69: WLCSP, 64-ball 4.466 × 4.395 
mm, 0.500 mm pitch, wafer-level chip-scale package mechanical data 
on page 121
LQFP64 package mechanical data updated: see Figure 73.: LQFP64 – 
10 x 10 mm 64 pin low-profile quad flat package outline and Table 73: 
LQFP64 – 10 x 10 mm 64 pin low-profile quad flat package mechanical 
data on page 130.
30-Sept-2014
9
Added Note 7 in Table 5: High-density STM32F103xC/D/E pin 
definitions on page 31.
Updated Note 10 in Table 5: High-density STM32F103xC/D/E pin 
definitions on page 31.
Modified Note 2 in Table 62: ADC accuracy on page 109
Modified Note 3 in Table 62: ADC accuracy on page 109
Modified notes in Table 51: I2C characteristics on page 97
Updated Figure 51: SPI timing diagram - master mode(1) on page 101
23-Feb-2015
10
Updated Figure 66.: BGA pad footprint, Figure 70: LQFP144 - 144-pin, 
20 x 20 mm low-profile quad flat package outline, Figure 73.: LQFP100 
– 14 x 14 mm 100 pin low-profile quad flat package outline, Figure 74.: 
LQFP100 recommended footprint, Figure 76.: LQFP64 – 10 x 10 mm 
64 pin low-profile quad flat package outline, Figure 77.: LQFP64 - 64-
pin, 10 x 10 mm low-profile quad flat recommended footprint
Added Figure 72.: LQFP144 marking example (package top view), 
Figure 75.: LQFP100 marking example (package top view), Figure 78.: 
LQFP64 marking example (package top view)
Updated Table 72: LQPF100 – 14 x 14 mm 100-pin low-profile quad flat 
package  mechanical data, Table 73: LQFP64 – 10 x 10 mm 64 pin low-
profile quad flat package mechanical data
Table 76.Document revision history
Date
Revision
Changes

---

---

**📄 Source: PDF Page 143**

DocID14611 Rev 12
143/144
STM32F103xC, STM32F103xD, STM32F103xE
Revision history
143
31-08-2015
11
Replaced USBDP and USBDM by USB_DP and USB_DM in the whole 
document.
Updated:
– Introduction
– Reference standard in Table 43: ESD absolute maximum ratings.
– Updated IDDA description in Table 63: DAC characteristics.
– Section : I2C interface characteristics
– Figure 62: LFBGA144 – 144-ball low profile fine pitch ball grid array, 
10 x 10 mm, 0.8 mm pitch, package outline
– Updated sentence before Figure 78: LQFP64 marking example 
(package top view).
– Figure 65: LFBGA100 - 10 x 10 mm low profile fine pitch ball grid 
array package outline and sentence before Figure 75: LQFP100 
marking example (package top view)
– Figure 68: WLCSP, 64-ball 4.466 × 4.395 mm, 0.500 mm pitch, 
wafer-level chip-scale package outline
– Figure 48: I2C bus AC waveforms and measurement circuit on 
page 98
– Section 6.1: LFBGA144 package information and Section 6.2: 
LFBGA100 package information.
– Table 20: Peripheral current consumption
Added:
– Figure 63: LFBGA144 – 144-ball low profile fine pitch ball grid array, 
10 x 10 mm, 0.8 mm pitch, package recommended footprint
– Figure 64: LFBGA144 marking example (package top view)
– Figure 66: LFBGA100 – 100-ball low profile fine pitch ball grid array, 
10 x 10 mm, 0.8 mm pitch, package recommended footprintoutline
– Figure 69: WLCSP64 - 64-ball, 4.4757 x 4.4049 mm, 0.5 mm pitch 
wafer level chip scale package recommended footprint
– Table 66: LFBGA144 recommended PCB design rules (0.8 mm pitch 
BGA)
– Table 68: LFBGA100 recommended PCB design rules (0.8 mm pitch 
BGA)
– Table 70: WLCSP64 recommended PCB design rules (0.5 mm pitch).
26-Nov-2015
12
Updated:
– Table 59: ADC characteristics
– Table 65: LFBGA144 – 144-ball low profile fine pitch ball grid array, 
10 x 10 mm, 0.8 mm pitch, package mechanical data
– Table 66: LFBGA144 recommended PCB design rules (0.8 mm pitch 
BGA)
Added:
– Note 3 on Table 7: Voltage characteristics
Table 76.Document revision history
Date
Revision
Changes

---

---

**📄 Source: PDF Page 144**

STM32F103xC, STM32F103xD, STM32F103xE
144/144
DocID14611 Rev 12
         
IMPORTANT NOTICE – PLEASE READ CAREFULLY
STMicroelectronics NV and its subsidiaries (“ST”) reserve the right to make changes, corrections, enhancements, modifications, and 
improvements to ST products and/or to this document at any time without notice. Purchasers should obtain the latest relevant information on 
ST products before placing orders. ST products are sold pursuant to ST’s terms and conditions of sale in place at the time of order 
acknowledgement.
Purchasers are solely responsible for the choice, selection, and use of ST products and ST assumes no liability for application assistance or 
the design of Purchasers’ products.
No license, express or implied, to any intellectual property right is granted by ST herein.
Resale of ST products with provisions different from the information set forth herein shall void any warranty granted by ST for such product.
ST and the ST logo are trademarks of ST. All other product or service names are the property of their respective owners.
Information in this document supersedes and replaces information previously supplied in any prior versions of this document.
© 2015 STMicroelectronics – All rights reserved

### Code Examples

```elixir
use of ST products and ST assumes no liability for application assistance or
```

---

