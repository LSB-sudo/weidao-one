# rdk-webrtc

Minimal USB camera to WebRTC preview service for `/root/rdk_notic`.

## What it does

- Discovers `/dev/video*` dynamically and prefers the first capture-capable USB camera.
- Exposes HTTP APIs on the board:
  - `GET /health`
  - `GET /devices`
  - `GET /gps`
  - `GET /viewer`
  - `POST /offer`
- Uses Go + Pion WebRTC.
- Uses `ffmpeg` as the capture pipeline:
  - V4L2 input from a USB camera
  - H.264 Annex-B output to stdout
  - Go parses NALUs into access units and writes them into a Pion H.264 track
- Access-unit boundaries now keep repeated SPS/PPS with the following IDR and only split VCL NALUs on a new slice header, so browser decoders do not receive half-frames or misplaced parameter sets
- Runs a background ATGM336H NMEA reader on the board and exposes the latest cached GPS state over HTTP

## Configuration

Environment variables and matching flags are supported.

| Env | Flag | Default | Meaning |
|---|---|---:|---|
| `LISTEN_ADDR` | `-listen` | `:8080` | HTTP listen address |
| `VIDEO_DEVICE` | `-device` | auto | Preferred `/dev/video*` node |
| `VIDEO_INPUT_FORMAT` | `-input-format` | `mjpeg` | Preferred V4L2 pixel/input format |
| `VIDEO_WIDTH` | `-width` | `640` | Capture width |
| `VIDEO_HEIGHT` | `-height` | `480` | Capture height |
| `VIDEO_FPS` | `-fps` | `15` | Capture frame rate |
| `VIDEO_BITRATE` | `-bitrate` | `700000` | H.264 bitrate in bits/s |
| `VIDEO_VFLIP` | n/a | `true` | Apply ffmpeg `vflip` by default to correct upside-down cameras |
| `VIDEO_HFLIP` | n/a | `false` | Apply optional ffmpeg `hflip` for mirrored image correction |
| `ICE_SERVERS` | `-ice` | empty | Comma-separated ICE URLs |
| `GPS_ENABLED` | `-gps-enabled` | `true` | Enable background GPS serial reader |
| `GPS_DEVICE` | `-gps-device` | `/dev/serial/by-id/usb-1a86_USB_Serial-if00-port0` | GPS serial device path |
| `GPS_BAUD` | `-gps-baud` | `9600` | GPS serial baud rate |
| `GPS_STALE_AFTER_SEC` | `-gps-stale-after-sec` | `10` | Mark GPS data stale when no sentence arrives within this many seconds |

The current USB camera on the board reports these useful modes:

- `MJPG`: up to `1280x720@30`
- `YUYV`: `1280x720@10`, `640x480@30`, `800x600@15`, `848x480@15`

The default profile now prefers low load and lower latency over image quality.

GPS uses a dedicated background reader when enabled. The HTTP handlers only return the latest cached snapshot and never open the serial device per request. On the RDK X5 target the default is enabled; set `GPS_ENABLED=false` when the module is absent or you need to free the serial port.

## Build

If the board has `go` in `PATH`:

```bash
go build ./...
```

If the board uses the project-local toolchain under `.local-go/`:

```bash
./.local-go/go/bin/go build ./...
```

## Run

```bash
./scripts/run_webrtc_usb.sh
```

Then open `http://<board-ip>:8080/viewer` from a browser on the same LAN.

## Android APK Viewer

A minimal Android Studio WebView client now lives in `android/RdkCameraViewer`.

- App title: `RDK Camera Viewer`
- Default URL: `http://192.168.3.142:8080/viewer`
- Intended workflow: open the project in Android Studio on a development machine, sync Gradle there, then build/install the APK
- Phone and board must be on the same LAN, and the board-side service must already be running

See `android/RdkCameraViewer/README.md` for Android Studio usage and IP customization.

## Orientation

- Default is `VIDEO_VFLIP=true`, so the service applies `vflip` unless you disable it.
- If the picture is already upright, start with `VIDEO_VFLIP=false`.
- If the picture is left-right mirrored, add `VIDEO_HFLIP=true`.
- Both can be enabled together when a camera mount needs 180-degree correction.

## Diagnostics

- `GET /health` now includes:
  - detected camera state
  - effective stream config
  - flip settings (`vflip` / `hflip`)
  - ffmpeg argument list and command summary
  - GPS summary: `enabled`, `device`, `baud`, `connected`, `valid`, `stale`, `lastSentenceType`, `antennaStatus`, and `lastError`
- `GET /devices` shows the discovered `/dev/video*` nodes and probe summaries.
- `/viewer` shows basic WebRTC stats from `RTCPeerConnection.getStats()` and warns when dropped frames are high even though packet loss is zero, which usually points to decode pressure or H.264 packetization issues:
  - connection state
  - frames decoded / dropped
  - bytes received
  - jitter
  - current decoded resolution when available

## Stutter Tuning

Start by lowering capture load before changing anything else.

```bash
VIDEO_WIDTH=640 \
VIDEO_HEIGHT=360 \
VIDEO_FPS=15 \
VIDEO_BITRATE=600000 \
./scripts/run_webrtc_usb.sh
```

If the camera behaves better with raw YUYV at lower resolutions, try:

```bash
VIDEO_INPUT_FORMAT=yuyv422 \
VIDEO_WIDTH=640 \
VIDEO_HEIGHT=480 \
VIDEO_FPS=15 \
VIDEO_BITRATE=700000 \
./scripts/run_webrtc_usb.sh
```

If the board CPU is still high:

- lower `VIDEO_FPS` to `10`
- lower `VIDEO_BITRATE` to `500000`
- prefer `640x360` or `640x480`
- check `/health` to confirm the service is actually using the expected ffmpeg parameters
- open browser devtools or the built-in `/viewer` stats panel to watch decoded frames and dropped frames while testing

## Notes

- No public STUN/TURN is configured by default.
- If no USB camera is available, the service still starts and `/health` plus `/devices` report the camera state.
- `POST /offer` returns a clear error when no capture-capable camera is available.

## GPS API

- `GET /gps` returns the latest cached GPS snapshot, including:
  - serial status: `enabled`, `device`, `baud`, `connected`, `stale`, `staleAfterSec`
  - validity: `valid`, `lastValidAt`, `lastError`
  - latest input: `lastSentenceType`, `lastSentence`, `lastText`, `antennaStatus`, `lastReadAt`, `lastSentenceAt`
  - parsed fix fields when available: latitude, longitude, speed in knots and km/h, course, altitude, satellites, HDOP, and fix quality
- `GET /health` includes a compact `gps` summary alongside the existing camera and stream diagnostics.

Current ATGM336H observations on this board can legitimately report `connected=true` and `valid=false` before the antenna has a usable fix. In that state the service still exposes the latest NMEA sentence, TXT antenna status such as `ANTENNA OK`, and any serial or parser error.
