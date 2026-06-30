# PRD: Niagara MQTT WD1 protocol integration

## Goal

Add a fixed-topic Niagara MQTT interaction layer for WD1 on the RDK side without replacing existing MQTT, GPS, or WebRTC behavior.

## Requirements

- Support fixed publish topics under `wd1/boat/sensor/...`
- Support fixed subscribe topics under `wd1/boat/cmd/...`
- Keep this iteration at MQTT protocol level only
- Do not connect command topics to real lower-controller actions yet
- Publish Niagara `gps_pos` as `longitude,latitude` when valid


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
