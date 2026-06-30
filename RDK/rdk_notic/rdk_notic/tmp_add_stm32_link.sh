cd /root/rdk_notic || exit 1
mkdir -p internal/stm32link mydocs

cat > internal/stm32link/serial_linux.go <<'EOF'
//go:build linux

package stm32link

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func openSerial(device string, baud int) (*os.File, error) {
	speed, err := baudAttr(baud)
	if err != nil {
		return nil, err
	}

	serial, err := os.OpenFile(device, os.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, err
	}
	if err := configureSerial(serial.Fd(), speed); err != nil {
		_ = serial.Close()
		return nil, err
	}
	if err := syscall.SetNonblock(int(serial.Fd()), false); err != nil {
		_ = serial.Close()
		return nil, err
	}
	return serial, nil
}

func configureSerial(fd uintptr, speed uint32) error {
	termios := &syscall.Termios{}
	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(termios)), 0, 0, 0); errno != 0 {
		return errno
	}

	termios.Iflag = 0
	termios.Oflag = 0
	termios.Lflag = 0
	termios.Cflag &^= syscall.CSIZE | syscall.PARENB | syscall.CSTOPB
	termios.Cflag |= syscall.CREAD | syscall.CLOCAL | syscall.CS8
	termios.Ispeed = speed
	termios.Ospeed = speed
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0

	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(termios)), 0, 0, 0); errno != 0 {
		return errno
	}
	return nil
}

func baudAttr(baud int) (uint32, error) {
	switch baud {
	case 4800:
		return syscall.B4800, nil
	case 9600:
		return syscall.B9600, nil
	case 19200:
		return syscall.B19200, nil
	case 38400:
		return syscall.B38400, nil
	case 57600:
		return syscall.B57600, nil
	case 115200:
		return syscall.B115200, nil
	default:
		return 0, fmt.Errorf("unsupported baud rate %d", baud)
	}
}
EOF

cat > internal/stm32link/service.go <<'EOF'
package stm32link

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	DefaultDevice        = "/dev/serial/by-id/usb-stm32_ttl_usb1-port0"
	DefaultBaud          = 115200
	DefaultStaleAfterSec = 5
)

type Config struct {
	Enabled       bool   `json:"enabled"`
	Device        string `json:"device"`
	Baud          int    `json:"baud"`
	StaleAfterSec int    `json:"staleAfterSec"`
}

type Command struct {
	LeftSetRPM  float64 `json:"leftSetRpm"`
	RightSetRPM float64 `json:"rightSetRpm"`
	BoatRun     bool    `json:"boatRun"`
}

type Snapshot struct {
	Enabled        bool       `json:"enabled"`
	Device         string     `json:"device"`
	Baud           int        `json:"baud"`
	StaleAfterSec  int        `json:"staleAfterSec"`
	Connected      bool       `json:"connected"`
	Valid          bool       `json:"valid"`
	Stale          bool       `json:"stale"`
	LastReadAt     *time.Time `json:"lastReadAt,omitempty"`
	LastWriteAt    *time.Time `json:"lastWriteAt,omitempty"`
	LastFrame      string     `json:"lastFrame,omitempty"`
	LastError      string     `json:"lastError,omitempty"`
	ActualLeftRPM  float64    `json:"actualLeftRpm"`
	ActualRightRPM float64    `json:"actualRightRpm"`
	LastCommand    *Command   `json:"lastCommand,omitempty"`
}

type Summary struct {
	Enabled        bool       `json:"enabled"`
	Device         string     `json:"device"`
	Baud           int        `json:"baud"`
	Connected      bool       `json:"connected"`
	Valid          bool       `json:"valid"`
	Stale          bool       `json:"stale"`
	LastReadAt     *time.Time `json:"lastReadAt,omitempty"`
	LastWriteAt    *time.Time `json:"lastWriteAt,omitempty"`
	LastError      string     `json:"lastError,omitempty"`
	ActualLeftRPM  float64    `json:"actualLeftRpm"`
	ActualRightRPM float64    `json:"actualRightRpm"`
}

type Service struct {
	cfg          Config
	startOnce    sync.Once
	commandPulse chan struct{}

	mu       sync.RWMutex
	snapshot Snapshot
}

func NewService(cfg Config) *Service {
	cfg = normalizeConfig(cfg)
	return &Service{
		cfg: cfg,
		commandPulse: make(chan struct{}, 1),
		snapshot: Snapshot{
			Enabled:       cfg.Enabled,
			Device:        cfg.Device,
			Baud:          cfg.Baud,
			StaleAfterSec: cfg.StaleAfterSec,
		},
	}
}

func normalizeConfig(cfg Config) Config {
	if cfg.Device == "" {
		cfg.Device = DefaultDevice
	}
	if cfg.Baud <= 0 {
		cfg.Baud = DefaultBaud
	}
	if cfg.StaleAfterSec <= 0 {
		cfg.StaleAfterSec = DefaultStaleAfterSec
	}
	return cfg
}

func (s *Service) Start(ctx context.Context) {
	if !s.cfg.Enabled {
		return
	}
	s.startOnce.Do(func() {
		log.Printf("stm32link: starting serial link device=%s baud=%d", s.cfg.Device, s.cfg.Baud)
		go s.run(ctx)
	})
}

func (s *Service) Snapshot() Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := s.snapshot
	out.Stale = s.isStaleLocked(time.Now().UTC())
	out.LastReadAt = cloneTime(out.LastReadAt)
	out.LastWriteAt = cloneTime(out.LastWriteAt)
	out.LastCommand = cloneCommand(out.LastCommand)
	return out
}

func (s *Service) Summary() Summary {
	snapshot := s.Snapshot()
	return Summary{
		Enabled:        snapshot.Enabled,
		Device:         snapshot.Device,
		Baud:           snapshot.Baud,
		Connected:      snapshot.Connected,
		Valid:          snapshot.Valid,
		Stale:          snapshot.Stale,
		LastReadAt:     cloneTime(snapshot.LastReadAt),
		LastWriteAt:    cloneTime(snapshot.LastWriteAt),
		LastError:      snapshot.LastError,
		ActualLeftRPM:  snapshot.ActualLeftRPM,
		ActualRightRPM: snapshot.ActualRightRPM,
	}
}

func (s *Service) ApplyCommand(cmd Command) {
	s.mu.Lock()
	s.snapshot.LastCommand = cloneCommand(&cmd)
	s.mu.Unlock()

	select {
	case s.commandPulse <- struct{}{}:
	default:
	}
}

func (s *Service) run(ctx context.Context) {
	for {
		if err := ctx.Err(); err != nil {
			return
		}

		serial, err := openSerial(s.cfg.Device, s.cfg.Baud)
		if err != nil {
			s.setDisconnected(fmt.Sprintf("open serial %s: %v", s.cfg.Device, err))
			log.Printf("stm32link: %v", err)
			if !sleepContext(ctx, time.Second) {
				return
			}
			continue
		}

		s.setConnected()
		select {
		case s.commandPulse <- struct{}{}:
		default:
		}

		err = s.session(ctx, serial)
		_ = serial.Close()
		if errors.Is(err, context.Canceled) || ctx.Err() != nil {
			s.setDisconnected("")
			return
		}
		if err != nil {
			log.Printf("stm32link: session stopped: %v", err)
			s.setDisconnected(err.Error())
		} else {
			s.setDisconnected("stm32 session stopped")
		}
		if !sleepContext(ctx, time.Second) {
			return
		}
	}
}

func (s *Service) session(ctx context.Context, serial *os.File) error {
	sessionCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, 2)
	go func() {
		errCh <- s.readLoop(sessionCtx, serial)
	}()
	go func() {
		errCh <- s.writeLoop(sessionCtx, serial)
	}()

	select {
	case <-ctx.Done():
		return context.Canceled
	case err := <-errCh:
		cancel()
		if err == nil {
			return errors.New("stm32 session ended")
		}
		return err
	}
}

func (s *Service) readLoop(ctx context.Context, serial io.Reader) error {
	reader := bufio.NewReaderSize(serial, 4096)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if ctx.Err() != nil {
				return context.Canceled
			}
			return fmt.Errorf("read serial line: %w", err)
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if err := s.handleFrame(trimmed); err != nil {
			log.Printf("stm32link: ignored frame=%q err=%v", trimmed, err)
		}
	}
}

func (s *Service) writeLoop(ctx context.Context, serial *os.File) error {
	for {
		select {
		case <-ctx.Done():
			return context.Canceled
		case <-s.commandPulse:
			frame, ok := s.currentCommandFrame()
			if !ok {
				continue
			}
			if _, err := serial.Write([]byte(frame)); err != nil {
				return fmt.Errorf("write command frame: %w", err)
			}
			s.recordWrite(frame)
			log.Printf("stm32link: sent frame=%s", strings.TrimSpace(frame))
		}
	}
}

func (s *Service) currentCommandFrame() (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.snapshot.LastCommand == nil {
		return "", false
	}
	cmd := *s.snapshot.LastCommand
	return formatCommand(cmd), true
}

func formatCommand(cmd Command) string {
	run := 0
	if cmd.BoatRun {
		run = 1
	}
	return fmt.Sprintf(
		"CMD,left_set_rpm=%s,right_set_rpm=%s,boat_run=%d\n",
		formatFloat(cmd.LeftSetRPM),
		formatFloat(cmd.RightSetRPM),
		run,
	)
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (s *Service) handleFrame(frame string) error {
	fields := strings.Split(frame, ",")
	if len(fields) == 0 {
		return errors.New("empty frame")
	}
	kind := strings.ToUpper(strings.TrimSpace(fields[0]))
	if kind != "FB" && kind != "STATE" && kind != "RPM" {
		return fmt.Errorf("unsupported frame type %q", kind)
	}

	values := map[string]string{}
	for _, field := range fields[1:] {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		parts := strings.SplitN(field, "=", 2)
		if len(parts) != 2 {
			continue
		}
		values[strings.ToLower(strings.TrimSpace(parts[0]))] = strings.TrimSpace(parts[1])
	}

	left, err := parseAnyFloat(values, "left_rpm", "left_actual_rpm", "l", "left")
	if err != nil {
		return fmt.Errorf("left rpm: %w", err)
	}
	right, err := parseAnyFloat(values, "right_rpm", "right_actual_rpm", "r", "right")
	if err != nil {
		return fmt.Errorf("right rpm: %w", err)
	}

	now := time.Now().UTC()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshot.Connected = true
	s.snapshot.Valid = true
	s.snapshot.Stale = false
	s.snapshot.LastReadAt = cloneTime(&now)
	s.snapshot.LastFrame = frame
	s.snapshot.LastError = ""
	s.snapshot.ActualLeftRPM = left
	s.snapshot.ActualRightRPM = right
	return nil
}

func parseAnyFloat(values map[string]string, keys ...string) (float64, error) {
	for _, key := range keys {
		raw, ok := values[key]
		if !ok {
			continue
		}
		return strconv.ParseFloat(raw, 64)
	}
	return 0, fmt.Errorf("missing any of keys %v", keys)
}

func (s *Service) recordWrite(frame string) {
	now := time.Now().UTC()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshot.LastWriteAt = cloneTime(&now)
	s.snapshot.LastError = ""
}

func (s *Service) setConnected() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshot.Connected = true
	s.snapshot.LastError = ""
}

func (s *Service) setDisconnected(lastErr string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshot.Connected = false
	s.snapshot.Stale = s.isStaleLocked(time.Now().UTC())
	if lastErr != "" {
		s.snapshot.LastError = lastErr
	}
}

func (s *Service) staleAfter() time.Duration {
	return time.Duration(s.cfg.StaleAfterSec) * time.Second
}

func (s *Service) isStaleLocked(now time.Time) bool {
	if s.snapshot.LastReadAt == nil {
		return false
	}
	return now.Sub(*s.snapshot.LastReadAt) > s.staleAfter()
}

func cloneTime(in *time.Time) *time.Time {
	if in == nil {
		return nil
	}
	t := *in
	return &t
}

func cloneCommand(in *Command) *Command {
	if in == nil {
		return nil
	}
	cmd := *in
	return &cmd
}

func sleepContext(ctx context.Context, d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}
EOF

cat > mydocs/STM32_SERIAL_PROTOCOL.md <<'EOF'
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

STM32 should return actual measured RPM with a single line:

```text
FB,left_rpm=<float>,right_rpm=<float>
```

Example:

```text
FB,left_rpm=241.5,right_rpm=-61.2
```

RDK currently also accepts `STATE` or `RPM` instead of `FB` as the leading token, and it accepts these field aliases:

- left side: `left_rpm`, `left_actual_rpm`, `l`, `left`
- right side: `right_rpm`, `right_actual_rpm`, `r`, `right`

## Current RDK behavior

- Niagara MQTT topics:
  - subscribe:
    - `wd1/boat/cmd/motor_left_set`
    - `wd1/boat/cmd/motor_right_set`
    - `wd1/boat/cmd/boat_run`
  - publish:
    - `wd1/boat/sensor/motor_left_rpm`
    - `wd1/boat/sensor/motor_right_rpm`
- After RDK receives Niagara setpoint topics, it caches the latest values and sends one `CMD,...` line to STM32.
- After RDK receives one valid STM32 feedback line, it publishes the actual left/right RPM back to Niagara MQTT.

## Implementation note for STM32

- Each frame must end with `\n`
- Keep output one frame per line
- ASCII decimal numbers are sufficient
- If one side has no valid RPM yet, return `0`
EOF

cat > internal/mqttclient/service.go <<'EOF'
package mqttclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"rdk_notic/internal/gps"
	"rdk_notic/internal/stm32link"
)

const (
	DefaultBrokerURL          = "tcp://broker.emqx.io:1883"
	DefaultDeviceID           = "rdk-x5-001"
	DefaultClientID           = "rdk-x5-001-rdk"
	DefaultPublishIntervalSec = 5
	DefaultSource             = "gps-service"
	DefaultPayloadMode        = PayloadModeGPS

	PayloadModeGPS                  = "gps"
	PayloadModeOneNETBatteryVoltage = "onenet_battery_voltage"
	PayloadModeNiagaraWd1           = "niagara_wd1"

	NiagaraTopicBatteryVoltage = "wd1/boat/sensor/battery"
	NiagaraTopicMotorLeftRPM   = "wd1/boat/sensor/motor_left_rpm"
	NiagaraTopicMotorRightRPM  = "wd1/boat/sensor/motor_right_rpm"
	NiagaraTopicBirdAlarm      = "wd1/boat/sensor/bird_alarm"
	NiagaraTopicBirdScare      = "wd1/boat/sensor/bird_scare_status"
	NiagaraTopicBoatRunStatus  = "wd1/boat/sensor/boat_run_status"
	NiagaraTopicGPSPos         = "wd1/boat/sensor/gps_pos"

	NiagaraTopicMotorLeftSet  = "wd1/boat/cmd/motor_left_set"
	NiagaraTopicMotorRightSet = "wd1/boat/cmd/motor_right_set"
	NiagaraTopicBoatRunCtrl   = "wd1/boat/cmd/boat_run"
)

type Config struct {
	Enabled            bool   `json:"enabled"`
	BrokerURL          string `json:"brokerUrl"`
	DeviceID           string `json:"deviceId"`
	ClientID           string `json:"clientId"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"-"`
	PublishTopic       string `json:"publishTopic,omitempty"`
	SubscribeTopic     string `json:"subscribeTopic,omitempty"`
	PayloadMode        string `json:"payloadMode,omitempty"`
	PublishIntervalSec int    `json:"publishIntervalSec"`
}

type Publisher struct {
	cfg          Config
	snapshot     func() gps.Snapshot
	stm32Snapshot func() stm32link.Snapshot
	applyCommand func(stm32link.Command)
	startOnce    sync.Once
	publishMu    sync.Mutex
	commandMu    sync.RWMutex
	onenetSequence int
	messageID    uint64
	commands     NiagaraCommandState
}

type payload struct {
	DeviceID         string   `json:"deviceId"`
	TS               string   `json:"ts"`
	Source           string   `json:"source"`
	Connected        bool     `json:"connected"`
	Valid            bool     `json:"valid"`
	Stale            bool     `json:"stale"`
	AntennaStatus    string   `json:"antennaStatus,omitempty"`
	LastSentenceType string   `json:"lastSentenceType,omitempty"`
	LastSentence     string   `json:"lastSentence,omitempty"`
	Fix              *gps.Fix `json:"fix,omitempty"`
}

type oneNETPropertyPost struct {
	ID      string                    `json:"id"`
	Version string                    `json:"version"`
	Params  map[string]oneNETProperty `json:"params"`
}

type oneNETProperty struct {
	Value float64 `json:"value"`
}

type NiagaraCommandState struct {
	MotorLeftSetRPM  NiagaraNumericCommand `json:"motorLeftSetRpm"`
	MotorRightSetRPM NiagaraNumericCommand `json:"motorRightSetRpm"`
	BoatRunCtrl      NiagaraBoolCommand    `json:"boatRunCtrl"`
}

type NiagaraNumericCommand struct {
	Value      float64    `json:"value"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	RawPayload string     `json:"rawPayload,omitempty"`
	Valid      bool       `json:"valid"`
}

type NiagaraBoolCommand struct {
	Value      bool       `json:"value"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	RawPayload string     `json:"rawPayload,omitempty"`
	Valid      bool       `json:"valid"`
}

type NiagaraPublish struct {
	Topic   string
	Payload string
}

func NormalizeConfig(cfg Config) Config {
	if cfg.BrokerURL == "" {
		cfg.BrokerURL = DefaultBrokerURL
	}
	if cfg.DeviceID == "" {
		cfg.DeviceID = DefaultDeviceID
	}
	if cfg.ClientID == "" {
		cfg.ClientID = DefaultClientID
	}
	if cfg.PublishIntervalSec <= 0 {
		cfg.PublishIntervalSec = DefaultPublishIntervalSec
	}
	cfg.PayloadMode = normalizePayloadMode(cfg.PayloadMode)
	if cfg.PayloadMode == PayloadModeGPS && cfg.PublishTopic == "" {
		cfg.PublishTopic = Topic(cfg.DeviceID)
	}
	return cfg
}

func Topic(deviceID string) string {
	return fmt.Sprintf("devices/%s/gps", deviceID)
}

func New(cfg Config, snapshot func() gps.Snapshot, stm32Snapshot func() stm32link.Snapshot, applyCommand func(stm32link.Command)) *Publisher {
	return &Publisher{
		cfg:           NormalizeConfig(cfg),
		snapshot:      snapshot,
		stm32Snapshot: stm32Snapshot,
		applyCommand:  applyCommand,
	}
}

func (p *Publisher) Start(ctx context.Context) {
	if !p.cfg.Enabled || p.snapshot == nil {
		return
	}
	p.startOnce.Do(func() {
		go p.run(ctx)
	})
}

func (p *Publisher) run(ctx context.Context) {
	opts := mqtt.NewClientOptions().
		AddBroker(p.cfg.BrokerURL).
		SetClientID(p.cfg.ClientID).
		SetUsername(p.cfg.Username).
		SetPassword(p.cfg.Password).
		SetProtocolVersion(4).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval(3 * time.Second).
		SetKeepAlive(30 * time.Second).
		SetPingTimeout(10 * time.Second).
		SetWriteTimeout(10 * time.Second).
		SetOrderMatters(false)

	opts.OnConnect = func(client mqtt.Client) {
		log.Printf("mqtt: connected broker=%s mode=%s publishTopic=%s clientId=%s", p.cfg.BrokerURL, p.cfg.PayloadMode, p.cfg.PublishTopic, p.cfg.ClientID)
		p.subscribe(client)
	}
	opts.OnConnectionLost = func(_ mqtt.Client, err error) {
		log.Printf("mqtt: connection lost: %v", err)
	}
	opts.OnReconnecting = func(_ mqtt.Client, _ *mqtt.ClientOptions) {
		log.Printf("mqtt: reconnecting broker=%s", p.cfg.BrokerURL)
	}

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if ok := token.WaitTimeout(15 * time.Second); !ok {
		log.Printf("mqtt: initial connect timeout broker=%s", p.cfg.BrokerURL)
	} else if err := token.Error(); err != nil {
		log.Printf("mqtt: initial connect failed broker=%s err=%v", p.cfg.BrokerURL, err)
	}

	ticker := time.NewTicker(time.Duration(p.cfg.PublishIntervalSec) * time.Second)
	defer ticker.Stop()
	defer client.Disconnect(250)

	p.publish(client)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.publish(client)
		}
	}
}

func (p *Publisher) subscribe(client mqtt.Client) {
	if p.cfg.PayloadMode == PayloadModeNiagaraWd1 {
		p.subscribeNiagara(client)
		return
	}
	if strings.TrimSpace(p.cfg.SubscribeTopic) == "" {
		return
	}

	token := client.Subscribe(p.cfg.SubscribeTopic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		log.Printf("mqtt: received topic=%s payload=%s", msg.Topic(), string(msg.Payload()))
	})
	if ok := token.WaitTimeout(10 * time.Second); !ok {
		log.Printf("mqtt: subscribe timeout topic=%s", p.cfg.SubscribeTopic)
		return
	}
	if err := token.Error(); err != nil {
		log.Printf("mqtt: subscribe failed topic=%s err=%v", p.cfg.SubscribeTopic, err)
		return
	}

	log.Printf("mqtt: subscribed topic=%s", p.cfg.SubscribeTopic)
}

func (p *Publisher) subscribeNiagara(client mqtt.Client) {
	topics := []string{
		NiagaraTopicMotorLeftSet,
		NiagaraTopicMotorRightSet,
		NiagaraTopicBoatRunCtrl,
	}
	for _, topic := range topics {
		topic := topic
		token := client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
			p.handleNiagaraCommand(msg.Topic(), string(msg.Payload()))
		})
		if ok := token.WaitTimeout(10 * time.Second); !ok {
			log.Printf("mqtt: subscribe timeout topic=%s", topic)
			continue
		}
		if err := token.Error(); err != nil {
			log.Printf("mqtt: subscribe failed topic=%s err=%v", topic, err)
			continue
		}
		log.Printf("mqtt: subscribed topic=%s", topic)
	}
}

func (p *Publisher) publish(client mqtt.Client) {
	if !client.IsConnected() {
		log.Printf("mqtt: publish skipped connected=false mode=%s", p.cfg.PayloadMode)
		return
	}

	if p.cfg.PayloadMode == PayloadModeNiagaraWd1 {
		p.publishNiagara(client)
		return
	}

	body, detail, err := p.buildPayload()
	if err != nil {
		log.Printf("mqtt: build payload failed: %v", err)
		return
	}

	token := client.Publish(p.cfg.PublishTopic, 0, false, body)
	if ok := token.WaitTimeout(10 * time.Second); !ok {
		log.Printf("mqtt: publish timeout topic=%s", p.cfg.PublishTopic)
		return
	}
	if err := token.Error(); err != nil {
		log.Printf("mqtt: publish failed topic=%s err=%v", p.cfg.PublishTopic, err)
		return
	}

	log.Printf("mqtt: published topic=%s bytes=%d %s", p.cfg.PublishTopic, len(body), detail)
}

func (p *Publisher) publishNiagara(client mqtt.Client) {
	publishes := p.buildNiagaraPublishes()
	for _, item := range publishes {
		token := client.Publish(item.Topic, 0, false, item.Payload)
		if ok := token.WaitTimeout(10 * time.Second); !ok {
			log.Printf("mqtt: publish timeout topic=%s", item.Topic)
			continue
		}
		if err := token.Error(); err != nil {
			log.Printf("mqtt: publish failed topic=%s err=%v", item.Topic, err)
			continue
		}
		log.Printf("mqtt: published topic=%s payload=%s", item.Topic, item.Payload)
	}
}

func (p *Publisher) buildNiagaraPublishes() []NiagaraPublish {
	snapshot := p.snapshot()
	gpsPos := buildNiagaraGPSPos(snapshot)

	leftRPM := 0.0
	rightRPM := 0.0
	if p.stm32Snapshot != nil {
		stm32 := p.stm32Snapshot()
		if stm32.Valid {
			leftRPM = stm32.ActualLeftRPM
			rightRPM = stm32.ActualRightRPM
		}
	}

	return []NiagaraPublish{
		{Topic: NiagaraTopicBatteryVoltage, Payload: formatFloat(0)},
		{Topic: NiagaraTopicMotorLeftRPM, Payload: formatFloat(leftRPM)},
		{Topic: NiagaraTopicMotorRightRPM, Payload: formatFloat(rightRPM)},
		{Topic: NiagaraTopicBirdAlarm, Payload: strconv.FormatBool(false)},
		{Topic: NiagaraTopicBirdScare, Payload: strconv.FormatBool(false)},
		{Topic: NiagaraTopicBoatRunStatus, Payload: strconv.FormatBool(false)},
		{Topic: NiagaraTopicGPSPos, Payload: gpsPos},
	}
}

func buildNiagaraGPSPos(snapshot gps.Snapshot) string {
	if !snapshot.Valid || snapshot.Fix == nil {
		return ""
	}
	return fmt.Sprintf("%.6f,%.6f", snapshot.Fix.Longitude, snapshot.Fix.Latitude)
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (p *Publisher) handleNiagaraCommand(topic, raw string) {
	raw = strings.TrimSpace(raw)
	now := time.Now().UTC()
	p.commandMu.Lock()
	defer p.commandMu.Unlock()

	switch topic {
	case NiagaraTopicMotorLeftSet:
		value, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			log.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)
			return
		}
		p.commands.MotorLeftSetRPM = NiagaraNumericCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}
		log.Printf("mqtt: cached niagara command topic=%s value=%s", topic, formatFloat(value))
	case NiagaraTopicMotorRightSet:
		value, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			log.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)
			return
		}
		p.commands.MotorRightSetRPM = NiagaraNumericCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}
		log.Printf("mqtt: cached niagara command topic=%s value=%s", topic, formatFloat(value))
	case NiagaraTopicBoatRunCtrl:
		value, err := strconv.ParseBool(strings.ToLower(raw))
		if err != nil {
			log.Printf("mqtt: invalid niagara command topic=%s payload=%q err=%v", topic, raw, err)
			return
		}
		p.commands.BoatRunCtrl = NiagaraBoolCommand{Value: value, UpdatedAt: &now, RawPayload: raw, Valid: true}
		log.Printf("mqtt: cached niagara command topic=%s value=%t", topic, value)
	default:
		log.Printf("mqtt: unhandled niagara command topic=%s payload=%s", topic, raw)
		return
	}

	p.pushSTM32CommandLocked()
}

func (p *Publisher) pushSTM32CommandLocked() {
	if p.applyCommand == nil {
		return
	}
	cmd := stm32link.Command{}
	if p.commands.MotorLeftSetRPM.Valid {
		cmd.LeftSetRPM = p.commands.MotorLeftSetRPM.Value
	}
	if p.commands.MotorRightSetRPM.Valid {
		cmd.RightSetRPM = p.commands.MotorRightSetRPM.Value
	}
	if p.commands.BoatRunCtrl.Valid {
		cmd.BoatRun = p.commands.BoatRunCtrl.Value
	}
	p.applyCommand(cmd)
}

func (p *Publisher) buildPayload() ([]byte, string, error) {
	switch p.cfg.PayloadMode {
	case PayloadModeOneNETBatteryVoltage:
		return p.buildOneNETBatteryVoltagePayload()
	case PayloadModeGPS:
		fallthrough
	default:
		return p.buildGPSPayload()
	}
}

func (p *Publisher) buildGPSPayload() ([]byte, string, error) {
	snapshot := p.snapshot()
	msg := buildPayload(p.cfg.DeviceID, snapshot)
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, "", err
	}
	return body, fmt.Sprintf("mode=%s valid=%t stale=%t connected=%t", PayloadModeGPS, msg.Valid, msg.Stale, msg.Connected), nil
}

func (p *Publisher) buildOneNETBatteryVoltagePayload() ([]byte, string, error) {
	p.publishMu.Lock()
	value := float64(p.onenetSequence)
	p.onenetSequence = (p.onenetSequence + 1) % 16
	p.messageID++
	id := fmt.Sprintf("%d", p.messageID)
	p.publishMu.Unlock()

	msg := oneNETPropertyPost{
		ID:      id,
		Version: "1.0",
		Params: map[string]oneNETProperty{
			"battery_voltage": {Value: value},
		},
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, "", err
	}
	return body, fmt.Sprintf("mode=%s battery_voltage=%.0f id=%s", PayloadModeOneNETBatteryVoltage, value, id), nil
}

func buildPayload(deviceID string, snapshot gps.Snapshot) payload {
	return payload{
		DeviceID:         deviceID,
		TS:               time.Now().UTC().Format(time.RFC3339Nano),
		Source:           DefaultSource,
		Connected:        snapshot.Connected,
		Valid:            snapshot.Valid,
		Stale:            snapshot.Stale,
		AntennaStatus:    snapshot.AntennaStatus,
		LastSentenceType: snapshot.LastSentenceType,
		LastSentence:     snapshot.LastSentence,
		Fix:              cloneFix(snapshot.Fix),
	}
}

func cloneFix(in *gps.Fix) *gps.Fix {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func normalizePayloadMode(mode string) string {
	mode = strings.TrimSpace(strings.ToLower(mode))
	if mode == "" {
		return DefaultPayloadMode
	}
	return mode
}
EOF

cat > cmd/rdk-webrtc/main.go <<'EOF'
package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"rdk_notic/internal/gps"
	mqttclient "rdk_notic/internal/mqttclient"
	"rdk_notic/internal/server"
	"rdk_notic/internal/stm32link"
	"rdk_notic/internal/stream"
)

func main() {
	listenAddr := flag.String("listen", envString("LISTEN_ADDR", ":8080"), "HTTP listen address")
	devicePath := flag.String("device", envString("VIDEO_DEVICE", ""), "preferred video device path, e.g. /dev/video0")
	inputFormat := flag.String("input-format", envString("VIDEO_INPUT_FORMAT", "mjpeg"), "preferred V4L2 input format, e.g. mjpeg or yuyv422")
	width := flag.Int("width", envInt("VIDEO_WIDTH", 640), "video width")
	height := flag.Int("height", envInt("VIDEO_HEIGHT", 480), "video height")
	fps := flag.Int("fps", envInt("VIDEO_FPS", 15), "video frames per second")
	bitrate := flag.Int("bitrate", envInt("VIDEO_BITRATE", 700000), "video bitrate in bits per second")
	iceServers := flag.String("ice", envString("ICE_SERVERS", ""), "comma-separated ICE server URLs")
	gpsEnabled := flag.Bool("gps-enabled", envBool("GPS_ENABLED", true), "enable GPS serial reader")
	gpsDevice := flag.String("gps-device", envString("GPS_DEVICE", gps.DefaultDevice), "GPS serial device path")
	gpsBaud := flag.Int("gps-baud", envInt("GPS_BAUD", gps.DefaultBaud), "GPS serial baud rate")
	gpsStaleAfter := flag.Int("gps-stale-after-sec", envInt("GPS_STALE_AFTER_SEC", gps.DefaultStaleAfterSec), "seconds before GPS data is considered stale")
	stm32Enabled := flag.Bool("stm32-enabled", envBool("STM32_ENABLED", false), "enable STM32 serial link")
	stm32Device := flag.String("stm32-device", envString("STM32_DEVICE", stm32link.DefaultDevice), "STM32 serial device path")
	stm32Baud := flag.Int("stm32-baud", envInt("STM32_BAUD", stm32link.DefaultBaud), "STM32 serial baud rate")
	stm32StaleAfter := flag.Int("stm32-stale-after-sec", envInt("STM32_STALE_AFTER_SEC", stm32link.DefaultStaleAfterSec), "seconds before STM32 feedback is considered stale")
	mqttEnabled := flag.Bool("mqtt-enabled", envBool("MQTT_ENABLED", false), "enable MQTT publisher")
	mqttBrokerURL := flag.String("mqtt-broker-url", envString("MQTT_BROKER_URL", mqttclient.DefaultBrokerURL), "MQTT broker URL")
	mqttDeviceID := flag.String("mqtt-device-id", envString("MQTT_DEVICE_ID", mqttclient.DefaultDeviceID), "MQTT device identifier")
	mqttClientID := flag.String("mqtt-client-id", envString("MQTT_CLIENT_ID", mqttclient.DefaultClientID), "MQTT client identifier")
	mqttUsername := flag.String("mqtt-username", envString("MQTT_USERNAME", ""), "MQTT username")
	mqttPassword := flag.String("mqtt-password", envString("MQTT_PASSWORD", ""), "MQTT password or token")
	mqttPublishTopic := flag.String("mqtt-publish-topic", envString("MQTT_PUBLISH_TOPIC", ""), "MQTT publish topic")
	mqttSubscribeTopic := flag.String("mqtt-subscribe-topic", envString("MQTT_SUBSCRIBE_TOPIC", ""), "MQTT subscribe topic")
	mqttPayloadMode := flag.String("mqtt-payload-mode", envString("MQTT_PAYLOAD_MODE", mqttclient.DefaultPayloadMode), "MQTT payload mode: gps, onenet_battery_voltage, niagara_wd1")
	mqttPublishInterval := flag.Int("mqtt-publish-interval-sec", envInt("MQTT_PUBLISH_INTERVAL_SEC", mqttclient.DefaultPublishIntervalSec), "MQTT publish interval in seconds")
	flag.Parse()

	cfg := server.Config{
		ListenAddr: *listenAddr,
		ICEServers: splitCSV(*iceServers),
		Stream: stream.Config{
			DevicePath:  *devicePath,
			InputFormat: *inputFormat,
			Width:       *width,
			Height:      *height,
			FPS:         *fps,
			Bitrate:     *bitrate,
			VFlip:       envBool("VIDEO_VFLIP", true),
			HFlip:       envBool("VIDEO_HFLIP", false),
		},
		GPS: gps.Config{
			Enabled:       *gpsEnabled,
			Device:        *gpsDevice,
			Baud:          *gpsBaud,
			StaleAfterSec: *gpsStaleAfter,
		},
		STM32: stm32link.Config{
			Enabled:       *stm32Enabled,
			Device:        *stm32Device,
			Baud:          *stm32Baud,
			StaleAfterSec: *stm32StaleAfter,
		},
		MQTT: mqttclient.Config{
			Enabled:            *mqttEnabled,
			BrokerURL:          *mqttBrokerURL,
			DeviceID:           *mqttDeviceID,
			ClientID:           *mqttClientID,
			Username:           *mqttUsername,
			Password:           *mqttPassword,
			PublishTopic:       *mqttPublishTopic,
			SubscribeTopic:     *mqttSubscribeTopic,
			PayloadMode:        *mqttPayloadMode,
			PublishIntervalSec: *mqttPublishInterval,
		},
	}

	log.Printf("starting rdk-webrtc on %s", cfg.ListenAddr)
	log.Printf("default stream config: %+v", stream.Describe(cfg.Stream).Config)
	if cfg.GPS.Enabled {
		log.Printf("gps config: enabled device=%s baud=%d staleAfterSec=%d", cfg.GPS.Device, cfg.GPS.Baud, cfg.GPS.StaleAfterSec)
	} else {
		log.Printf("gps config: disabled")
	}
	if cfg.STM32.Enabled {
		log.Printf("stm32 config: enabled device=%s baud=%d staleAfterSec=%d", cfg.STM32.Device, cfg.STM32.Baud, cfg.STM32.StaleAfterSec)
	} else {
		log.Printf("stm32 config: disabled")
	}
	if cfg.MQTT.Enabled {
		log.Printf("mqtt config: enabled broker=%s mode=%s publishTopic=%s subscribeTopic=%s clientId=%s username=%t password=%t intervalSec=%d", cfg.MQTT.BrokerURL, cfg.MQTT.PayloadMode, cfg.MQTT.PublishTopic, cfg.MQTT.SubscribeTopic, cfg.MQTT.ClientID, cfg.MQTT.Username != "", cfg.MQTT.Password != "", cfg.MQTT.PublishIntervalSec)
	} else {
		log.Printf("mqtt config: disabled")
	}

	srv := server.New(cfg)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := srv.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func envString(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func envBool(key string, fallback bool) bool {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func envInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func splitCSV(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}
EOF

cat > internal/server/server.go <<'EOF'
package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/pion/webrtc/v4"

	"rdk_notic/internal/device"
	"rdk_notic/internal/gps"
	mqttclient "rdk_notic/internal/mqttclient"
	"rdk_notic/internal/stm32link"
	"rdk_notic/internal/stream"
)

type Config struct {
	ListenAddr string            `json:"listenAddr"`
	ICEServers []string          `json:"iceServers,omitempty"`
	Stream     stream.Config     `json:"stream"`
	GPS        gps.Config        `json:"gps"`
	STM32      stm32link.Config  `json:"stm32"`
	MQTT       mqttclient.Config `json:"mqtt"`
}

type Server struct {
	cfg   Config
	gps   *gps.Service
	stm32 *stm32link.Service
	mqtt  *mqttclient.Publisher
}

type errorResponse struct {
	Error string `json:"error"`
}

const h264FmtpLine = "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f"

func New(cfg Config) *Server {
	gpsSvc := gps.NewService(cfg.GPS)
	stm32Svc := stm32link.NewService(cfg.STM32)
	return &Server{
		cfg:   cfg,
		gps:   gpsSvc,
		stm32: stm32Svc,
		mqtt:  mqttclient.New(cfg.MQTT, gpsSvc.Snapshot, stm32Svc.Snapshot, stm32Svc.ApplyCommand),
	}
}

func (s *Server) Run(ctx context.Context) error {
	s.gps.Start(ctx)
	s.stm32.Start(ctx)
	s.mqtt.Start(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/devices", s.handleDevices)
	mux.HandleFunc("/gps", s.handleGPS)
	mux.HandleFunc("/viewer", s.handleViewer)
	mux.HandleFunc("/viewer-static", s.handleViewerStatic)
	mux.HandleFunc("/offer", s.handleOffer)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/viewer", http.StatusTemporaryRedirect)
	})

	srv := &http.Server{
		Addr:    s.cfg.ListenAddr,
		Handler: mux,
	}

	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return srv.Shutdown(shutdownCtx)
	case err := <-errCh:
		return err
	}
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if handlePreflight(w, r) {
		return
	}
	applyCORS(w)
	discovery := device.Discover(s.cfg.Stream.DevicePath)
	runtimeCfg := s.cfg.Stream
	if discovery.Selected != "" {
		runtimeCfg.DevicePath = discovery.Selected
	}
	runtime := stream.Describe(runtimeCfg)

	writeJSON(w, http.StatusOK, map[string]any{
		"status":        "ok",
		"time":          time.Now().Format(time.RFC3339),
		"tools":         device.Tools(),
		"camera":        discovery,
		"config":        s.cfg,
		"gps":           s.gps.Summary(),
		"stm32":         s.stm32.Summary(),
		"streamRuntime": runtime,
		"webrtcCodec": map[string]any{
			"mime": webrtc.MimeTypeH264,
			"fmtp": h264FmtpLine,
		},
		"webrtcMime": webrtc.MimeTypeH264,
	})
}

func (s *Server) handleDevices(w http.ResponseWriter, r *http.Request) {
	if handlePreflight(w, r) {
		return
	}
	applyCORS(w)
	writeJSON(w, http.StatusOK, device.Discover(s.cfg.Stream.DevicePath))
}

func (s *Server) handleGPS(w http.ResponseWriter, r *http.Request) {
	if handlePreflight(w, r) {
		return
	}
	applyCORS(w)
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, s.gps.Snapshot())
}

func (s *Server) handleViewer(w http.ResponseWriter, r *http.Request) {
	if handlePreflight(w, r) {
		return
	}
	applyCORS(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(viewerHTML))
}

func (s *Server) handleOffer(w http.ResponseWriter, r *http.Request) {
	if handlePreflight(w, r) {
		return
	}
	applyCORS(w)
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}

	var offer webrtc.SessionDescription
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid offer json"})
		return
	}
	if offer.Type != webrtc.SDPTypeOffer {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "expected SDP offer"})
		return
	}

	discovery := device.Discover(s.cfg.Stream.DevicePath)
	if !discovery.Available {
		msg := "camera unavailable"
		if discovery.Error != "" {
			msg = fmt.Sprintf("camera unavailable: %s", discovery.Error)
		}
		writeJSON(w, http.StatusServiceUnavailable, errorResponse{Error: msg})
		return
	}

	peerCfg := webrtc.Configuration{}
	if len(s.cfg.ICEServers) > 0 {
		peerCfg.ICEServers = []webrtc.ICEServer{{URLs: s.cfg.ICEServers}}
	}

	peer, err := webrtc.NewPeerConnection(peerCfg)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("create peer: %v", err)})
		return
	}

	track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType:    webrtc.MimeTypeH264,
		ClockRate:   90000,
		SDPFmtpLine: h264FmtpLine,
	}, "video", "usbcam")
	if err != nil {
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("create track: %v", err)})
		return
	}

	sender, err := peer.AddTrack(track)
	if err != nil {
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("add track: %v", err)})
		return
	}
	go drainRTCP(sender)

	if err := peer.SetRemoteDescription(offer); err != nil {
		_ = peer.Close()
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: fmt.Sprintf("set remote description: %v", err)})
		return
	}

	streamCfg := s.cfg.Stream
	streamCfg.DevicePath = discovery.Selected
	runtime := stream.Describe(streamCfg)
	log.Printf("starting ffmpeg stream: %s", runtime.FFmpegCommand)
	streamCtx, cancelStream := context.WithCancel(context.Background())
	session, err := stream.Start(streamCtx, streamCfg, track)
	if err != nil {
		cancelStream()
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("start camera stream: %v", err)})
		return
	}

	peer.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log.Printf("ice state: %s", state.String())
		if state == webrtc.ICEConnectionStateDisconnected || state == webrtc.ICEConnectionStateFailed || state == webrtc.ICEConnectionStateClosed {
			cancelStream()
			_ = session.Close()
			_ = peer.Close()
		}
	})

	peer.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		log.Printf("peer state: %s", state.String())
		if state == webrtc.PeerConnectionStateFailed || state == webrtc.PeerConnectionStateClosed {
			cancelStream()
			_ = session.Close()
			_ = peer.Close()
		}
	})

	answer, err := peer.CreateAnswer(nil)
	if err != nil {
		cancelStream()
		_ = session.Close()
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("create answer: %v", err)})
		return
	}

	gatherComplete := webrtc.GatheringCompletePromise(peer)
	if err := peer.SetLocalDescription(answer); err != nil {
		cancelStream()
		_ = session.Close()
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("set local description: %v", err)})
		return
	}
	<-gatherComplete

	local := peer.LocalDescription()
	if local == nil {
		cancelStream()
		_ = session.Close()
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "local description missing"})
		return
	}

	writeJSON(w, http.StatusOK, local)
}

func drainRTCP(sender *webrtc.RTPSender) {
	buffer := make([]byte, 1500)
	for {
		if _, _, err := sender.Read(buffer); err != nil {
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

type videoSession struct {
	peer   *webrtc.PeerConnection
	stream *stream.Session
	cancel context.CancelFunc
	once   sync.Once
}

func (s *videoSession) close() {
	s.once.Do(func() {
		if s.cancel != nil {
			s.cancel()
		}
		if s.stream != nil {
			_ = s.stream.Close()
		}
		if s.peer != nil {
			_ = s.peer.Close()
		}
	})
}

func closeIfErr(target *videoSession, err error) error {
	if err != nil && target != nil {
		target.close()
	}
	return err
}

func safeError(message string, err error) errorResponse {
	if err == nil {
		return errorResponse{Error: message}
	}
	return errorResponse{Error: fmt.Sprintf("%s: %v", message, err)}
}

func ignoreClosed(err error) error {
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
EOF

./.local-go/go/bin/gofmt -w internal/stm32link/*.go internal/mqttclient/service.go cmd/rdk-webrtc/main.go internal/server/server.go &&
./.local-go/go/bin/go build ./...
