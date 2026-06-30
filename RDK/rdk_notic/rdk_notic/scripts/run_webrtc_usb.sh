#!/usr/bin/env sh
set -eu

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
ROOT_DIR=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)
GO_BIN=${GO_BIN:-}

if [ -z "$GO_BIN" ] && [ -x "$ROOT_DIR/.local-go/go/bin/go" ]; then
  GO_BIN="$ROOT_DIR/.local-go/go/bin/go"
fi

if [ -z "$GO_BIN" ]; then
  GO_BIN=$(command -v go || true)
fi

if [ -z "$GO_BIN" ]; then
  echo "go toolchain not found. Set GO_BIN or install Go into PATH/.local-go." >&2
  exit 1
fi

mkdir -p "$ROOT_DIR/bin"
cd "$ROOT_DIR"
"$GO_BIN" build -o "$ROOT_DIR/bin/rdk-webrtc" ./cmd/rdk-webrtc
exec "$ROOT_DIR/bin/rdk-webrtc" "$@"
