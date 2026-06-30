package gps

import (
	"math"
	"testing"
	"time"
)

func TestParseSentenceRejectsChecksumMismatch(t *testing.T) {
	_, err := ParseSentence("$GPTXT,01,01,01,ANTENNA OK*00")
	if err == nil {
		t.Fatal("expected checksum error")
	}
}

func TestParseSentenceInvalidRMC(t *testing.T) {
	msg, err := ParseSentence("$GNRMC,,V,,,,,,,,,,N,V*37")
	if err != nil {
		t.Fatalf("ParseSentence() error = %v", err)
	}
	if msg.Formatter != "RMC" || msg.RMC == nil {
		t.Fatalf("expected RMC payload, got %+v", msg)
	}
	if msg.RMC.Status != "V" {
		t.Fatalf("status = %q, want V", msg.RMC.Status)
	}
	if msg.RMC.TimestampUTC != nil {
		t.Fatalf("timestamp = %v, want nil", msg.RMC.TimestampUTC)
	}
	if msg.RMC.HasLatitude || msg.RMC.HasLongitude {
		t.Fatalf("expected empty coordinates, got lat=%v lon=%v", msg.RMC.Latitude, msg.RMC.Longitude)
	}
}

func TestParseSentenceValidRMC(t *testing.T) {
	msg, err := ParseSentence("$GPRMC,123519,A,4807.038,N,01131.000,E,022.4,084.4,230394,003.1,W*6A")
	if err != nil {
		t.Fatalf("ParseSentence() error = %v", err)
	}
	if msg.Formatter != "RMC" || msg.RMC == nil {
		t.Fatalf("expected RMC payload, got %+v", msg)
	}
	if msg.RMC.Status != "A" {
		t.Fatalf("status = %q, want A", msg.RMC.Status)
	}
	if !msg.RMC.HasLatitude || !almostEqual(msg.RMC.Latitude, 48.1173) {
		t.Fatalf("latitude = %v, want 48.1173", msg.RMC.Latitude)
	}
	if !msg.RMC.HasLongitude || !almostEqual(msg.RMC.Longitude, 11.5166666667) {
		t.Fatalf("longitude = %v, want 11.5166666667", msg.RMC.Longitude)
	}
	if !msg.RMC.HasSpeed || !almostEqual(msg.RMC.SpeedKnots, 22.4) {
		t.Fatalf("speedKnots = %v, want 22.4", msg.RMC.SpeedKnots)
	}
	if !almostEqual(msg.RMC.SpeedKPH, 41.4848) {
		t.Fatalf("speedKph = %v, want 41.4848", msg.RMC.SpeedKPH)
	}
	if !msg.RMC.HasCourse || !almostEqual(msg.RMC.Course, 84.4) {
		t.Fatalf("course = %v, want 84.4", msg.RMC.Course)
	}
	wantTime := time.Date(1994, time.March, 23, 12, 35, 19, 0, time.UTC)
	if msg.RMC.TimestampUTC == nil || !msg.RMC.TimestampUTC.Equal(wantTime) {
		t.Fatalf("timestamp = %v, want %v", msg.RMC.TimestampUTC, wantTime)
	}
}

func TestParseSentenceGGA(t *testing.T) {
	msg, err := ParseSentence("$GPGGA,123520,4807.038,N,01131.000,E,1,08,0.9,545.4,M,46.9,M,,*4D")
	if err != nil {
		t.Fatalf("ParseSentence() error = %v", err)
	}
	if msg.Formatter != "GGA" || msg.GGA == nil {
		t.Fatalf("expected GGA payload, got %+v", msg)
	}
	if !msg.GGA.HasLatitude || !almostEqual(msg.GGA.Latitude, 48.1173) {
		t.Fatalf("latitude = %v, want 48.1173", msg.GGA.Latitude)
	}
	if !msg.GGA.HasLongitude || !almostEqual(msg.GGA.Longitude, 11.5166666667) {
		t.Fatalf("longitude = %v, want 11.5166666667", msg.GGA.Longitude)
	}
	if msg.GGA.FixQuality != 1 {
		t.Fatalf("fixQuality = %d, want 1", msg.GGA.FixQuality)
	}
	if msg.GGA.Satellites != 8 {
		t.Fatalf("satellites = %d, want 8", msg.GGA.Satellites)
	}
	if !almostEqual(msg.GGA.HDOP, 0.9) {
		t.Fatalf("hdop = %v, want 0.9", msg.GGA.HDOP)
	}
	if !almostEqual(msg.GGA.AltitudeMeters, 545.4) {
		t.Fatalf("altitude = %v, want 545.4", msg.GGA.AltitudeMeters)
	}
}

func TestParseSentenceTXT(t *testing.T) {
	msg, err := ParseSentence("$GPTXT,01,01,01,ANTENNA OK*35")
	if err != nil {
		t.Fatalf("ParseSentence() error = %v", err)
	}
	if msg.Formatter != "TXT" || msg.TXT == nil {
		t.Fatalf("expected TXT payload, got %+v", msg)
	}
	if msg.TXT.Text != "ANTENNA OK" {
		t.Fatalf("text = %q, want %q", msg.TXT.Text, "ANTENNA OK")
	}
}

func TestGuessFormatter(t *testing.T) {
	if got := GuessFormatter("$GNRMC,,V,,,,,,,,,,N,V*37"); got != "RMC" {
		t.Fatalf("GuessFormatter() = %q, want RMC", got)
	}
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.000001
}
