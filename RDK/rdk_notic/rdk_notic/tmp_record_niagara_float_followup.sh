cd /root/rdk_notic || exit 1
python3 - <<'PY'
from pathlib import Path

journal = Path(".trellis/workspace/LSB/journal-1.md")
task_prd = Path(".trellis/tasks/06-23-niagara-mqtt-wd1/prd.md")
task_json = Path(".trellis/tasks/06-23-niagara-mqtt-wd1/task.json")

journal_append = """

## [2026-06-23 20:45 CST] Niagara MQTT 下行浮点命令兼容修复

- 背景：Niagara 已经可以把控制命令下发到 RDK，但 `wd1/boat/cmd/motor_left_set` 和 `wd1/boat/cmd/motor_right_set` 实际发送的是浮点数字符串，例如 `249.116`、`-65.51599999999999`。
- 问题表现：RDK 端日志出现 `strconv.Atoi` 解析失败，说明 MQTT 订阅链路正常，但下行数值命令被错误地按整数处理。
- 代码修复：已将 `internal/mqttclient/service.go` 中 Niagara 数值命令缓存结构 `NiagaraNumericCommand.Value` 从 `int` 调整为 `float64`，并将两个电机命令的解析从 `strconv.Atoi` 改为 `strconv.ParseFloat(raw, 64)`。
- 日志行为：修复后，电机命令将以浮点值缓存并记录，例如 `mqtt: cached niagara command topic=wd1/boat/cmd/motor_left_set value=249.116`。
- 验证：已在远端执行 `./.local-go/go/bin/gofmt -w internal/mqttclient/service.go` 和 `./.local-go/go/bin/go build ./...`，构建通过。
- 结果：Niagara 到 RDK 的下行控制 Topic 现已兼容浮点数值格式，后续无需强制 Niagara 侧改成整数发布。
"""

prd_append = """

## Follow-up: Float command compatibility

- Date: 2026-06-23
- Context: Niagara numeric publish points for `wd1/boat/cmd/motor_left_set` and `wd1/boat/cmd/motor_right_set` sent decimal payloads instead of integer strings.
- Failure mode: RDK logs showed `strconv.Atoi` parse failures even though MQTT subscription and delivery were functioning correctly.
- Resolution:
  - changed `NiagaraNumericCommand.Value` from `int` to `float64`
  - changed the two motor command parsers from `strconv.Atoi` to `strconv.ParseFloat(raw, 64)`
  - kept existing topic names and command-cache behavior unchanged
- Verification: `./.local-go/go/bin/go build ./...` passed after the change
- Decision: Niagara-facing numeric downlink topics should remain tolerant of decimal payloads because Niagara numeric points may naturally emit floating-point values.
"""

journal.write_text(journal.read_text(encoding="utf-8") + journal_append, encoding="utf-8")
task_prd.write_text(task_prd.read_text(encoding="utf-8") + prd_append, encoding="utf-8")

task_text = task_json.read_text(encoding="utf-8")
if '"updated_at"' in task_text and "2026-06-23T20:45:00+08:00" not in task_text:
    import re
    task_text = re.sub(
        r'"updated_at"\s*:\s*"[^"]+"',
        '"updated_at": "2026-06-23T20:45:00+08:00"',
        task_text,
        count=1,
    )
task_json.write_text(task_text, encoding="utf-8")
PY
