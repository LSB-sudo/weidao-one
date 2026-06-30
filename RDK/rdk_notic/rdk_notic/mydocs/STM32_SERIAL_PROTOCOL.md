# STM32 Serial Protocol

This document defines the serial link between RDK and STM32 for motor setpoint delivery and actual RPM feedback.

## Serial parameters

- Default RDK device: `/dev/serial/by-id/usb-stm32_ttl_usb1-port0`
- Default baud: `115200`
- Frame format: `8N1`
- Transport: newline-terminated ASCII text lines

## RDK -> STM32 command frame

RDK sends the latest design / setpoint values with a single line:

```text
CMD,left_set_rpm=<float>,right_set_rpm=<float>,boat_run=<0|1>
```

Example:

```text
CMD,left_set_rpm=249.116,right_set_rpm=-65.516,boat_run=1
```

Field meanings:

- `left_set_rpm`: left motor target speed from Niagara / RDK
- `right_set_rpm`: right motor target speed from Niagara / RDK
- `boat_run`: overall run flag from Niagara
  - `1` = run enabled
  - `0` = run disabled

## STM32 -> RDK feedback frame

STM32 should return actual measured RPM and battery voltage with a single line:

```text
FB,left_rpm=<float>,right_rpm=<float>,battery_voltage=<float>
```

Example:

```text
FB,left_rpm=241.5,right_rpm=-61.2,battery_voltage=12.34
```

RDK currently also accepts `STATE` or `RPM` instead of `FB` as the leading token, and it accepts these field aliases:

- left side: `left_rpm`, `left_actual_rpm`, `l`, `left`
- right side: `right_rpm`, `right_actual_rpm`, `r`, `right`
- battery: `battery_voltage`, `battery`, `bat_v`, `voltage`

## Current RDK behavior

- Niagara MQTT topics:
  - subscribe:
    - `wd1/boat/cmd/motor_left_set`
    - `wd1/boat/cmd/motor_right_set`
    - `wd1/boat/cmd/boat_run`
  - publish:
    - `wd1/boat/sensor/motor_left_rpm`
    - `wd1/boat/sensor/motor_right_rpm`
    - `wd1/boat/sensor/battery`
- After RDK receives Niagara setpoint topics, it caches the latest values and sends one `CMD,...` line to STM32.
- After RDK receives one valid STM32 feedback line, it publishes the actual left/right RPM and battery voltage back to Niagara MQTT.

## Implementation note for STM32

- Each frame must end with `\n`
- Keep output one frame per line
- ASCII decimal numbers are sufficient
- If one side has no valid RPM yet, return `0`
