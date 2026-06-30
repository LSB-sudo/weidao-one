package gps

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

const (
	DefaultDevice        = "/dev/serial/by-id/usb-1a86_USB_Serial-if00-port0"
	DefaultBaud          = 9600
	DefaultStaleAfterSec = 10
)

type Config struct {
	Enabled       bool   `json:"enabled"`
	Device        string `json:"device"`
	Baud          int    `json:"baud"`
	StaleAfterSec int    `json:"staleAfterSec"`
}

type Fix struct {
	Source         string  `json:"source,omitempty"`
	TimestampUTC   string  `json:"timestampUtc,omitempty"`
	Status         string  `json:"status,omitempty"`
	Mode           string  `json:"mode,omitempty"`
	NavStatus      string  `json:"navStatus,omitempty"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	SpeedKnots     float64 `json:"speedKnots"`
	SpeedKPH       float64 `json:"speedKph"`
	Course         float64 `json:"course"`
	AltitudeMeters float64 `json:"altitudeMeters"`
	Satellites     int     `json:"satellites"`
	HDOP           float64 `json:"hdop"`
	FixQuality     int     `json:"fixQuality"`
}

type Snapshot struct {
	Enabled          bool       `json:"enabled"`
	Device           string     `json:"device"`
	Baud             int        `json:"baud"`
	StaleAfterSec    int        `json:"staleAfterSec"`
	Connected        bool       `json:"connected"`
	Valid            bool       `json:"valid"`
	Stale            bool       `json:"stale"`
	LastReadAt       *time.Time `json:"lastReadAt,omitempty"`
	LastSentenceAt   *time.Time `json:"lastSentenceAt,omitempty"`
	LastValidAt      *time.Time `json:"lastValidAt,omitempty"`
	LastSentenceType string     `json:"lastSentenceType,omitempty"`
	LastSentence     string     `json:"lastSentence,omitempty"`
	LastText         string     `json:"lastText,omitempty"`
	AntennaStatus    string     `json:"antennaStatus,omitempty"`
	LastError        string     `json:"lastError,omitempty"`
	Fix              *Fix       `json:"fix,omitempty"`
}

type Summary struct {
	Enabled          bool       `json:"enabled"`
	Device           string     `json:"device"`
	Baud             int        `json:"baud"`
	Connected        bool       `json:"connected"`
	Valid            bool       `json:"valid"`
	Stale            bool       `json:"stale"`
	LastReadAt       *time.Time `json:"lastReadAt,omitempty"`
	LastValidAt      *time.Time `json:"lastValidAt,omitempty"`
	LastSentenceType string     `json:"lastSentenceType,omitempty"`
	AntennaStatus    string     `json:"antennaStatus,omitempty"`
	LastError        string     `json:"lastError,omitempty"`
}

type Service struct {
	cfg       Config
	mu        sync.RWMutex
	snapshot  Snapshot
	fix       Fix
	hasFix    bool
	rmcValid  bool
	ggaValid  bool
	startOnce sync.Once
}

func NewService(cfg Config) *Service {
	cfg = normalizeConfig(cfg)
	return &Service{
		cfg: cfg,
		snapshot: Snapshot{
			Enabled:       cfg.Enabled,
			Device:        cfg.Device,
			Baud:          cfg.Baud,
			StaleAfterSec: cfg.StaleAfterSec,
		},
	}
}

func (s *Service) Start(ctx context.Context) {
	if !s.cfg.Enabled {
		return
	}
	s.startOnce.Do(func() {
		log.Printf("gps: starting serial reader device=%s baud=%d", s.cfg.Device, s.cfg.Baud)
		go s.run(ctx)
	})
}

func (s *Service) Snapshot() Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := cloneSnapshot(s.snapshot)
	out.Stale = s.isStaleLocked(time.Now().UTC())
	return out
}

func (s *Service) Summary() Summary {
	snapshot := s.Snapshot()
	return Summary{
		Enabled:          snapshot.Enabled,
		Device:           snapshot.Device,
		Baud:             snapshot.Baud,
		Connected:        snapshot.Connected,
		Valid:            snapshot.Valid,
		Stale:            snapshot.Stale,
		LastReadAt:       cloneTime(snapshot.LastReadAt),
		LastValidAt:      cloneTime(snapshot.LastValidAt),
		LastSentenceType: snapshot.LastSentenceType,
		AntennaStatus:    snapshot.AntennaStatus,
		LastError:        snapshot.LastError,
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

func (s *Service) staleAfter() time.Duration {
	return time.Duration(s.cfg.StaleAfterSec) * time.Second
}

func (s *Service) isStaleLocked(now time.Time) bool {
	if s.snapshot.LastReadAt == nil {
		return false
	}
	return now.Sub(*s.snapshot.LastReadAt) > s.staleAfter()
}

func (s *Service) currentValidLocked() bool {
	return s.rmcValid || s.ggaValid
}

func (s *Service) run(ctx context.Context) {
	for {
		if err := ctx.Err(); err != nil {
			return
		}

		serial, err := openSerial(s.cfg.Device, s.cfg.Baud)
		if err != nil {
			s.setDisconnected(fmt.Sprintf("open serial %s: %v", s.cfg.Device, err))
			log.Printf("gps: %v", err)
			if !sleepContext(ctx, time.Second) {
				return
			}
			continue
		}

		s.setConnected()
		err = s.readLoop(ctx, serial)
		_ = serial.Close()
		if errors.Is(err, context.Canceled) || ctx.Err() != nil {
			s.setDisconnected("")
			return
		}
		if err != nil {
			log.Printf("gps: reader stopped: %v", err)
			s.setDisconnected(err.Error())
		} else {
			s.setDisconnected("gps reader stopped")
		}
		if !sleepContext(ctx, time.Second) {
			return
		}
	}
}

func (s *Service) readLoop(ctx context.Context, serial io.Reader) error {
	reader := bufio.NewReaderSize(serial, 4096)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return err
			}
			if ctx.Err() != nil {
				return context.Canceled
			}
			return fmt.Errorf("read serial line: %w", err)
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		now := time.Now().UTC()
		msg, parseErr := ParseSentence(trimmed)
		s.recordSentence(now, trimmed, msg, parseErr)
	}
}

func (s *Service) recordSentence(now time.Time, raw string, msg Message, parseErr error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.snapshot.LastReadAt = timePtr(now)
	s.snapshot.LastSentenceAt = timePtr(now)
	s.snapshot.LastSentence = raw
	s.snapshot.LastSentenceType = msg.Formatter
	if s.snapshot.LastSentenceType == "" {
		s.snapshot.LastSentenceType = GuessFormatter(raw)
	}

	if parseErr != nil {
		s.snapshot.LastError = parseErr.Error()
		s.snapshot.Valid = s.currentValidLocked()
		s.snapshot.Stale = s.isStaleLocked(now)
		if s.hasFix {
			snapshotFix := s.fix
			s.snapshot.Fix = &snapshotFix
		}
		return
	}

	s.snapshot.LastError = ""
	if msg.RMC != nil {
		s.applyRMC(now, msg.RMC)
	}
	if msg.GGA != nil {
		s.applyGGA(now, msg.GGA)
	}
	if msg.VTG != nil {
		s.applyVTG(msg.VTG)
	}
	if msg.TXT != nil {
		s.applyTXT(msg.TXT)
	}

	s.snapshot.Valid = s.currentValidLocked()
	s.snapshot.Stale = s.isStaleLocked(now)
	if s.hasFix {
		snapshotFix := s.fix
		s.snapshot.Fix = &snapshotFix
	} else {
		s.snapshot.Fix = nil
	}
}

func (s *Service) applyRMC(now time.Time, data *RMCData) {
	s.hasFix = true
	s.fix.Source = "RMC"
	if data.TimestampUTC != nil {
		s.fix.TimestampUTC = data.TimestampUTC.Format(time.RFC3339Nano)
	}
	if data.Status != "" {
		s.fix.Status = data.Status
	}
	if data.Mode != "" {
		s.fix.Mode = data.Mode
	}
	if data.NavStatus != "" {
		s.fix.NavStatus = data.NavStatus
	}
	if data.HasLatitude {
		s.fix.Latitude = data.Latitude
	}
	if data.HasLongitude {
		s.fix.Longitude = data.Longitude
	}
	if data.HasSpeed {
		s.fix.SpeedKnots = data.SpeedKnots
		s.fix.SpeedKPH = data.SpeedKPH
	}
	if data.HasCourse {
		s.fix.Course = data.Course
	}

	s.rmcValid = strings.EqualFold(data.Status, "A")
	if s.rmcValid {
		s.snapshot.LastValidAt = timePtr(now)
	}
}

func (s *Service) applyGGA(now time.Time, data *GGAData) {
	s.hasFix = true
	if s.fix.Source == "" || data.HasLatitude || data.HasLongitude || data.HasFixQuality {
		s.fix.Source = "GGA"
	}
	if data.HasLatitude {
		s.fix.Latitude = data.Latitude
	}
	if data.HasLongitude {
		s.fix.Longitude = data.Longitude
	}
	if data.HasFixQuality {
		s.fix.FixQuality = data.FixQuality
	}
	if data.HasSatellites {
		s.fix.Satellites = data.Satellites
	}
	if data.HasHDOP {
		s.fix.HDOP = data.HDOP
	}
	if data.HasAltitude {
		s.fix.AltitudeMeters = data.AltitudeMeters
	}

	s.ggaValid = data.HasFixQuality && data.FixQuality > 0
	if s.ggaValid {
		s.snapshot.LastValidAt = timePtr(now)
	}
}

func (s *Service) applyVTG(data *VTGData) {
	s.hasFix = true
	if s.fix.Source == "" {
		s.fix.Source = "VTG"
	}
	if data.HasCourseTrue {
		s.fix.Course = data.CourseTrue
	}
	if data.HasSpeedKnots {
		s.fix.SpeedKnots = data.SpeedKnots
	}
	if data.HasSpeedKPH {
		s.fix.SpeedKPH = data.SpeedKPH
	} else if data.HasSpeedKnots {
		s.fix.SpeedKPH = data.SpeedKnots * 1.852
	}
	if data.Mode != "" {
		s.fix.Mode = data.Mode
	}
}

func (s *Service) applyTXT(data *TXTData) {
	s.snapshot.LastText = data.Text
	status := strings.TrimSpace(data.Text)
	if status != "" {
		s.snapshot.AntennaStatus = status
	}
}

func (s *Service) setConnected() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshot.Connected = true
}

func (s *Service) setDisconnected(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshot.Connected = false
	s.snapshot.Stale = s.isStaleLocked(time.Now().UTC())
	if message != "" {
		s.snapshot.LastError = message
	}
}

func cloneSnapshot(in Snapshot) Snapshot {
	out := in
	out.LastReadAt = cloneTime(in.LastReadAt)
	out.LastSentenceAt = cloneTime(in.LastSentenceAt)
	out.LastValidAt = cloneTime(in.LastValidAt)
	if in.Fix != nil {
		fix := *in.Fix
		out.Fix = &fix
	}
	return out
}

func cloneTime(in *time.Time) *time.Time {
	if in == nil {
		return nil
	}
	value := *in
	return &value
}

func timePtr(value time.Time) *time.Time {
	v := value.UTC()
	return &v
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
