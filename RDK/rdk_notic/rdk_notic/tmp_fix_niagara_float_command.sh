cd /root/rdk_notic || exit 1
python3 - <<'PY'
from pathlib import Path

path = Path("internal/mqttclient/service.go")
text = path.read_text(encoding="utf-8")

text = text.replace(
    'type NiagaraNumericCommand struct {\n\tValue      int        `json:"value"`\n\tUpdatedAt  *time.Time `json:"updatedAt,omitempty"`\n\tRawPayload string     `json:"rawPayload,omitempty"`\n\tValid      bool       `json:"valid"`\n}\n',
    'type NiagaraNumericCommand struct {\n\tValue      float64    `json:"value"`\n\tUpdatedAt  *time.Time `json:"updatedAt,omitempty"`\n\tRawPayload string     `json:"rawPayload,omitempty"`\n\tValid      bool       `json:"valid"`\n}\n',
)

text = text.replace(
    '\tcase NiagaraTopicMotorLeftSet:\n\t\tvalue, err := strconv.Atoi(raw)\n\t\tif err != nil {\n\t\t\tlog.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)\n\t\t\treturn\n\t\t}\n\t\tp.commands.MotorLeftSetRPM = NiagaraNumericCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}\n\t\tlog.Printf("mqtt: cached niagara command topic=%s value=%d", topic, value)\n\tcase NiagaraTopicMotorRightSet:\n\t\tvalue, err := strconv.Atoi(raw)\n\t\tif err != nil {\n\t\t\tlog.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)\n\t\t\treturn\n\t\t}\n\t\tp.commands.MotorRightSetRPM = NiagaraNumericCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}\n\t\tlog.Printf("mqtt: cached niagara command topic=%s value=%d", topic, value)\n',
    '\tcase NiagaraTopicMotorLeftSet:\n\t\tvalue, err := strconv.ParseFloat(raw, 64)\n\t\tif err != nil {\n\t\t\tlog.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)\n\t\t\treturn\n\t\t}\n\t\tp.commands.MotorLeftSetRPM = NiagaraNumericCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}\n\t\tlog.Printf("mqtt: cached niagara command topic=%s value=%s", topic, formatFloat(value))\n\tcase NiagaraTopicMotorRightSet:\n\t\tvalue, err := strconv.ParseFloat(raw, 64)\n\t\tif err != nil {\n\t\t\tlog.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)\n\t\t\treturn\n\t\t}\n\t\tp.commands.MotorRightSetRPM = NiagaraNumericCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}\n\t\tlog.Printf("mqtt: cached niagara command topic=%s value=%s", topic, formatFloat(value))\n',
)

path.write_text(text, encoding="utf-8")
PY

./.local-go/go/bin/gofmt -w internal/mqttclient/service.go &&
./.local-go/go/bin/go build ./...
