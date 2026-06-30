cd /root/rdk_notic || exit 1
python3 - <<'PY'
from pathlib import Path

path = Path("internal/server/server.go")
text = path.read_text(encoding="utf-8")
text = text.replace("session.Close()", "session.Stop()")
text = text.replace("s.stream.Close()", "s.stream.Stop()")
path.write_text(text, encoding="utf-8")
PY

./.local-go/go/bin/gofmt -w internal/server/server.go &&
./.local-go/go/bin/go build ./...
