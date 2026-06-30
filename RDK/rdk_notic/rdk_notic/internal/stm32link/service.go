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
	BatteryVoltage float64    `json:"batteryVoltage"`
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
	BatteryVoltage float64    `json:"batteryVoltage"`
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
		cfg:          cfg,
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
		BatteryVoltage: snapshot.BatteryVoltage,
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
	battery, err := parseAnyFloat(values, "battery_voltage", "battery", "bat_v", "voltage")
	if err != nil {
		return fmt.Errorf("battery voltage: %w", err)
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
	s.snapshot.BatteryVoltage = battery
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
