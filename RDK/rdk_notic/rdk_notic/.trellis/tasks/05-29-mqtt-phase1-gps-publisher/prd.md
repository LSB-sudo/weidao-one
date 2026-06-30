# MQTT Phase 1 GPS Publisher

## Goal
Implement a minimal MQTT Phase 1 subset in `/root/rdk_notic` so the RDK X5 periodically publishes the existing GPS snapshot to `broker.emqx.io:1883` on topic `devices/rdk-x5-001/gps`.

## Constraints
- Reuse the existing GPS service snapshot; do not open the serial port a second time.
- Keep WebRTC main flow unchanged.
- MQTT is disabled by default and enabled with environment variables.
- Allowed changes: `internal/mqtt/*.go`, `cmd/rdk-webrtc/main.go`, `internal/server/server.go`, `go.mod`, `go.sum`, optional small docs note.
- No destructive commands; do not revert unrelated remote changes.

## Validation
- `./.local-go/go/bin/go test ./...`
- `./.local-go/go/bin/go build ./...`
- Short runtime check with MQTT enabled and broker connectivity.
