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
	cfg            Config
	snapshot       func() gps.Snapshot
	stm32Snapshot  func() stm32link.Snapshot
	applyCommand   func(stm32link.Command)
	startOnce      sync.Once
	publishMu      sync.Mutex
	commandMu      sync.RWMutex
	onenetSequence int
	messageID      uint64
	commands       NiagaraCommandState
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
	batteryVoltage := 0.0
	if p.stm32Snapshot != nil {
		stm32 := p.stm32Snapshot()
		if stm32.Valid {
			leftRPM = stm32.ActualLeftRPM
			rightRPM = stm32.ActualRightRPM
			batteryVoltage = stm32.BatteryVoltage
		}
	}

	return []NiagaraPublish{
		{Topic: NiagaraTopicBatteryVoltage, Payload: formatFloat(batteryVoltage)},
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
