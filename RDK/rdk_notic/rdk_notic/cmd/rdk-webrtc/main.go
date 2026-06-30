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
