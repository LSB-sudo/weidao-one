package gps

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Raw       string
	Talker    string
	Formatter string
	RMC       *RMCData
	GGA       *GGAData
	VTG       *VTGData
	TXT       *TXTData
}

type RMCData struct {
	TimestampUTC *time.Time
	Status       string
	Latitude     float64
	HasLatitude  bool
	Longitude    float64
	HasLongitude bool
	SpeedKnots   float64
	HasSpeed     bool
	SpeedKPH     float64
	Course       float64
	HasCourse    bool
	Mode         string
	NavStatus    string
}

type GGAData struct {
	Latitude       float64
	HasLatitude    bool
	Longitude      float64
	HasLongitude   bool
	FixQuality     int
	HasFixQuality  bool
	Satellites     int
	HasSatellites  bool
	HDOP           float64
	HasHDOP        bool
	AltitudeMeters float64
	HasAltitude    bool
}

type VTGData struct {
	CourseTrue        float64
	HasCourseTrue     bool
	CourseMagnetic    float64
	HasCourseMagnetic bool
	SpeedKnots        float64
	HasSpeedKnots     bool
	SpeedKPH          float64
	HasSpeedKPH       bool
	Mode              string
}

type TXTData struct {
	TotalMessages  int
	SentenceNumber int
	TextID         int
	Text           string
}

func ParseSentence(raw string) (Message, error) {
	trimmed := strings.TrimSpace(raw)
	msg := Message{Raw: trimmed}
	if trimmed == "" {
		return msg, fmt.Errorf("empty sentence")
	}
	if trimmed[0] != '$' {
		return msg, fmt.Errorf("sentence must start with $")
	}

	star := strings.LastIndexByte(trimmed, '*')
	if star <= 1 || star+3 != len(trimmed) {
		return msg, fmt.Errorf("sentence missing checksum")
	}

	body := trimmed[1:star]
	expected, err := strconv.ParseUint(trimmed[star+1:], 16, 8)
	if err != nil {
		return msg, fmt.Errorf("invalid checksum encoding: %w", err)
	}
	actual := nmeaChecksum(body)
	if byte(expected) != actual {
		return msg, fmt.Errorf("checksum mismatch: want %02X got %02X", byte(expected), actual)
	}

	parts := strings.Split(body, ",")
	if len(parts) == 0 || parts[0] == "" {
		return msg, fmt.Errorf("missing sentence type")
	}

	id := strings.TrimSpace(parts[0])
	if len(id) < 3 {
		return msg, fmt.Errorf("invalid sentence type %q", id)
	}
	if len(id) > 3 {
		msg.Talker = id[:len(id)-3]
	}
	msg.Formatter = id[len(id)-3:]

	switch msg.Formatter {
	case "RMC":
		parsed, err := parseRMC(parts)
		if err != nil {
			return msg, err
		}
		msg.RMC = parsed
	case "GGA":
		parsed, err := parseGGA(parts)
		if err != nil {
			return msg, err
		}
		msg.GGA = parsed
	case "VTG":
		parsed, err := parseVTG(parts)
		if err != nil {
			return msg, err
		}
		msg.VTG = parsed
	case "TXT":
		parsed, err := parseTXT(parts)
		if err != nil {
			return msg, err
		}
		msg.TXT = parsed
	}

	return msg, nil
}

func GuessFormatter(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if len(trimmed) < 6 || trimmed[0] != '$' {
		return ""
	}
	body := trimmed[1:]
	if star := strings.IndexByte(body, '*'); star >= 0 {
		body = body[:star]
	}
	field := body
	if comma := strings.IndexByte(body, ','); comma >= 0 {
		field = body[:comma]
	}
	if len(field) < 3 {
		return ""
	}
	return field[len(field)-3:]
}

func parseRMC(parts []string) (*RMCData, error) {
	data := &RMCData{
		Status:    field(parts, 2),
		Mode:      field(parts, 12),
		NavStatus: field(parts, 13),
	}

	lat, ok, err := parseCoordinate(field(parts, 3), field(parts, 4))
	if err != nil {
		return nil, fmt.Errorf("parse RMC latitude: %w", err)
	}
	if ok {
		data.Latitude = lat
		data.HasLatitude = true
	}

	lon, ok, err := parseCoordinate(field(parts, 5), field(parts, 6))
	if err != nil {
		return nil, fmt.Errorf("parse RMC longitude: %w", err)
	}
	if ok {
		data.Longitude = lon
		data.HasLongitude = true
	}

	speedKnots, ok, err := parseOptionalFloat(field(parts, 7))
	if err != nil {
		return nil, fmt.Errorf("parse RMC speed: %w", err)
	}
	if ok {
		data.SpeedKnots = speedKnots
		data.SpeedKPH = speedKnots * 1.852
		data.HasSpeed = true
	}

	course, ok, err := parseOptionalFloat(field(parts, 8))
	if err != nil {
		return nil, fmt.Errorf("parse RMC course: %w", err)
	}
	if ok {
		data.Course = course
		data.HasCourse = true
	}

	timestamp, err := parseTimestamp(field(parts, 1), field(parts, 9))
	if err != nil {
		return nil, fmt.Errorf("parse RMC timestamp: %w", err)
	}
	data.TimestampUTC = timestamp

	return data, nil
}

func parseGGA(parts []string) (*GGAData, error) {
	data := &GGAData{}

	lat, ok, err := parseCoordinate(field(parts, 2), field(parts, 3))
	if err != nil {
		return nil, fmt.Errorf("parse GGA latitude: %w", err)
	}
	if ok {
		data.Latitude = lat
		data.HasLatitude = true
	}

	lon, ok, err := parseCoordinate(field(parts, 4), field(parts, 5))
	if err != nil {
		return nil, fmt.Errorf("parse GGA longitude: %w", err)
	}
	if ok {
		data.Longitude = lon
		data.HasLongitude = true
	}

	fixQuality, ok, err := parseOptionalInt(field(parts, 6))
	if err != nil {
		return nil, fmt.Errorf("parse GGA fix quality: %w", err)
	}
	if ok {
		data.FixQuality = fixQuality
		data.HasFixQuality = true
	}

	satellites, ok, err := parseOptionalInt(field(parts, 7))
	if err != nil {
		return nil, fmt.Errorf("parse GGA satellites: %w", err)
	}
	if ok {
		data.Satellites = satellites
		data.HasSatellites = true
	}

	hdop, ok, err := parseOptionalFloat(field(parts, 8))
	if err != nil {
		return nil, fmt.Errorf("parse GGA hdop: %w", err)
	}
	if ok {
		data.HDOP = hdop
		data.HasHDOP = true
	}

	altitude, ok, err := parseOptionalFloat(field(parts, 9))
	if err != nil {
		return nil, fmt.Errorf("parse GGA altitude: %w", err)
	}
	if ok {
		data.AltitudeMeters = altitude
		data.HasAltitude = true
	}

	return data, nil
}

func parseVTG(parts []string) (*VTGData, error) {
	data := &VTGData{Mode: field(parts, 9)}

	courseTrue, ok, err := parseOptionalFloat(field(parts, 1))
	if err != nil {
		return nil, fmt.Errorf("parse VTG true course: %w", err)
	}
	if ok {
		data.CourseTrue = courseTrue
		data.HasCourseTrue = true
	}

	courseMagnetic, ok, err := parseOptionalFloat(field(parts, 3))
	if err != nil {
		return nil, fmt.Errorf("parse VTG magnetic course: %w", err)
	}
	if ok {
		data.CourseMagnetic = courseMagnetic
		data.HasCourseMagnetic = true
	}

	speedKnots, ok, err := parseOptionalFloat(field(parts, 5))
	if err != nil {
		return nil, fmt.Errorf("parse VTG speed knots: %w", err)
	}
	if ok {
		data.SpeedKnots = speedKnots
		data.HasSpeedKnots = true
	}

	speedKPH, ok, err := parseOptionalFloat(field(parts, 7))
	if err != nil {
		return nil, fmt.Errorf("parse VTG speed km/h: %w", err)
	}
	if ok {
		data.SpeedKPH = speedKPH
		data.HasSpeedKPH = true
	}

	return data, nil
}

func parseTXT(parts []string) (*TXTData, error) {
	total, ok, err := parseOptionalInt(field(parts, 1))
	if err != nil {
		return nil, fmt.Errorf("parse TXT total messages: %w", err)
	}
	if !ok {
		total = 0
	}

	number, ok, err := parseOptionalInt(field(parts, 2))
	if err != nil {
		return nil, fmt.Errorf("parse TXT sentence number: %w", err)
	}
	if !ok {
		number = 0
	}

	textID, ok, err := parseOptionalInt(field(parts, 3))
	if err != nil {
		return nil, fmt.Errorf("parse TXT text id: %w", err)
	}
	if !ok {
		textID = 0
	}

	start := min(4, len(parts))
	return &TXTData{
		TotalMessages:  total,
		SentenceNumber: number,
		TextID:         textID,
		Text:           strings.Join(parts[start:], ","),
	}, nil
}

func nmeaChecksum(payload string) byte {
	var checksum byte
	for i := 0; i < len(payload); i++ {
		checksum ^= payload[i]
	}
	return checksum
}

func parseCoordinate(value, hemisphere string) (float64, bool, error) {
	value = strings.TrimSpace(value)
	hemisphere = strings.TrimSpace(strings.ToUpper(hemisphere))
	if value == "" || hemisphere == "" {
		return 0, false, nil
	}

	raw, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false, err
	}
	degrees := math.Floor(raw / 100)
	minutes := raw - (degrees * 100)
	decimal := degrees + (minutes / 60)

	switch hemisphere {
	case "N", "E":
	case "S", "W":
		decimal = -decimal
	default:
		return 0, false, fmt.Errorf("unsupported hemisphere %q", hemisphere)
	}

	return decimal, true, nil
}

func parseTimestamp(timeField, dateField string) (*time.Time, error) {
	if strings.TrimSpace(timeField) == "" || strings.TrimSpace(dateField) == "" {
		return nil, nil
	}
	if len(dateField) != 6 {
		return nil, fmt.Errorf("invalid date %q", dateField)
	}

	day, err := strconv.Atoi(dateField[0:2])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(dateField[2:4])
	if err != nil {
		return nil, err
	}
	year, err := strconv.Atoi(dateField[4:6])
	if err != nil {
		return nil, err
	}
	if year >= 70 {
		year += 1900
	} else {
		year += 2000
	}

	hour, minute, second, nanosecond, err := parseClock(timeField)
	if err != nil {
		return nil, err
	}

	timestamp := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.UTC)
	return &timestamp, nil
}

func parseClock(value string) (int, int, int, int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, 0, 0, 0, nil
	}
	if len(value) < 6 {
		return 0, 0, 0, 0, fmt.Errorf("invalid time %q", value)
	}

	hour, err := strconv.Atoi(value[0:2])
	if err != nil {
		return 0, 0, 0, 0, err
	}
	minute, err := strconv.Atoi(value[2:4])
	if err != nil {
		return 0, 0, 0, 0, err
	}
	second, err := strconv.Atoi(value[4:6])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	nanosecond := 0
	if len(value) > 6 {
		fractional := strings.TrimPrefix(value[6:], ".")
		if fractional != "" {
			if len(fractional) > 9 {
				fractional = fractional[:9]
			}
			for len(fractional) < 9 {
				fractional += "0"
			}
			nanosecond, err = strconv.Atoi(fractional)
			if err != nil {
				return 0, 0, 0, 0, err
			}
		}
	}

	return hour, minute, second, nanosecond, nil
}

func parseOptionalFloat(value string) (float64, bool, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false, nil
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false, err
	}
	return parsed, true, nil
}

func parseOptionalInt(value string) (int, bool, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, false, err
	}
	return parsed, true, nil
}

func field(parts []string, index int) string {
	if index < 0 || index >= len(parts) {
		return ""
	}
	return strings.TrimSpace(parts[index])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
