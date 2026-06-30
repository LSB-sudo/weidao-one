cd /root/rdk_notic || exit 1
python3 - <<'PY'
from pathlib import Path

svc = Path("internal/stm32link/service.go")
text = svc.read_text(encoding="utf-8")

text = text.replace(
    '\tActualLeftRPM  float64    `json:"actualLeftRpm"`\n\tActualRightRPM float64    `json:"actualRightRpm"`\n\tLastCommand    *Command   `json:"lastCommand,omitempty"`\n',
    '\tActualLeftRPM   float64    `json:"actualLeftRpm"`\n\tActualRightRPM  float64    `json:"actualRightRpm"`\n\tBatteryVoltage  float64    `json:"batteryVoltage"`\n\tLastCommand     *Command   `json:"lastCommand,omitempty"`\n',
)

text = text.replace(
    '\tActualLeftRPM  float64    `json:"actualLeftRpm"`\n\tActualRightRPM float64    `json:"actualRightRpm"`\n',
    '\tActualLeftRPM  float64    `json:"actualLeftRpm"`\n\tActualRightRPM float64    `json:"actualRightRpm"`\n\tBatteryVoltage float64    `json:"batteryVoltage"`\n',
)

text = text.replace(
    '\t\tLastError:      snapshot.LastError,\n\t\tActualLeftRPM:  snapshot.ActualLeftRPM,\n\t\tActualRightRPM: snapshot.ActualRightRPM,\n',
    '\t\tLastError:      snapshot.LastError,\n\t\tActualLeftRPM:  snapshot.ActualLeftRPM,\n\t\tActualRightRPM: snapshot.ActualRightRPM,\n\t\tBatteryVoltage: snapshot.BatteryVoltage,\n',
)

text = text.replace(
    '\tright, err := parseAnyFloat(values, "right_rpm", "right_actual_rpm", "r", "right")\n\tif err != nil {\n\t\treturn fmt.Errorf("right rpm: %w", err)\n\t}\n\n\tnow := time.Now().UTC()\n',
    '\tright, err := parseAnyFloat(values, "right_rpm", "right_actual_rpm", "r", "right")\n\tif err != nil {\n\t\treturn fmt.Errorf("right rpm: %w", err)\n\t}\n\tbattery, err := parseAnyFloat(values, "battery_voltage", "battery", "bat_v", "voltage")\n\tif err != nil {\n\t\treturn fmt.Errorf("battery voltage: %w", err)\n\t}\n\n\tnow := time.Now().UTC()\n',
)

text = text.replace(
    '\ts.snapshot.LastError = ""\n\ts.snapshot.ActualLeftRPM = left\n\ts.snapshot.ActualRightRPM = right\n\treturn nil\n',
    '\ts.snapshot.LastError = ""\n\ts.snapshot.ActualLeftRPM = left\n\ts.snapshot.ActualRightRPM = right\n\ts.snapshot.BatteryVoltage = battery\n\treturn nil\n',
)

svc.write_text(text, encoding="utf-8")

mqtt = Path("internal/mqttclient/service.go")
text = mqtt.read_text(encoding="utf-8")
text = text.replace(
    '\tleftRPM := 0.0\n\trightRPM := 0.0\n\tif p.stm32Snapshot != nil {\n\t\tstm32 := p.stm32Snapshot()\n\t\tif stm32.Valid {\n\t\t\tleftRPM = stm32.ActualLeftRPM\n\t\t\trightRPM = stm32.ActualRightRPM\n\t\t}\n\t}\n\n\treturn []NiagaraPublish{\n\t\t{Topic: NiagaraTopicBatteryVoltage, Payload: formatFloat(0)},\n',
    '\tleftRPM := 0.0\n\trightRPM := 0.0\n\tbatteryVoltage := 0.0\n\tif p.stm32Snapshot != nil {\n\t\tstm32 := p.stm32Snapshot()\n\t\tif stm32.Valid {\n\t\t\tleftRPM = stm32.ActualLeftRPM\n\t\t\trightRPM = stm32.ActualRightRPM\n\t\t\tbatteryVoltage = stm32.BatteryVoltage\n\t\t}\n\t}\n\n\treturn []NiagaraPublish{\n\t\t{Topic: NiagaraTopicBatteryVoltage, Payload: formatFloat(batteryVoltage)},\n',
)
mqtt.write_text(text, encoding="utf-8")

doc = Path("mydocs/STM32_SERIAL_PROTOCOL.md")
text = doc.read_text(encoding="utf-8")
text = text.replace(
    'STM32 should return actual measured RPM with a single line:\n\n```text\nFB,left_rpm=<float>,right_rpm=<float>\n```\n\nExample:\n\n```text\nFB,left_rpm=241.5,right_rpm=-61.2\n```\n',
    'STM32 should return actual measured RPM and battery voltage with a single line:\n\n```text\nFB,left_rpm=<float>,right_rpm=<float>,battery_voltage=<float>\n```\n\nExample:\n\n```text\nFB,left_rpm=241.5,right_rpm=-61.2,battery_voltage=12.34\n```\n',
)
text = text.replace(
    '- right side: `right_rpm`, `right_actual_rpm`, `r`, `right`\n',
    '- right side: `right_rpm`, `right_actual_rpm`, `r`, `right`\n- battery: `battery_voltage`, `battery`, `bat_v`, `voltage`\n',
)
text = text.replace(
    '    - `wd1/boat/sensor/motor_left_rpm`\n    - `wd1/boat/sensor/motor_right_rpm`\n',
    '    - `wd1/boat/sensor/motor_left_rpm`\n    - `wd1/boat/sensor/motor_right_rpm`\n    - `wd1/boat/sensor/battery`\n',
)
text = text.replace(
    '- After RDK receives one valid STM32 feedback line, it publishes the actual left/right RPM back to Niagara MQTT.\n',
    '- After RDK receives one valid STM32 feedback line, it publishes the actual left/right RPM and battery voltage back to Niagara MQTT.\n',
)
doc.write_text(text, encoding="utf-8")
PY

./.local-go/go/bin/gofmt -w internal/stm32link/service.go internal/mqttclient/service.go &&
./.local-go/go/bin/go build ./...
