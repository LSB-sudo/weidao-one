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
			_ = session.Stop()
			_ = peer.Close()
		}
	})

	peer.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		log.Printf("peer state: %s", state.String())
		if state == webrtc.PeerConnectionStateFailed || state == webrtc.PeerConnectionStateClosed {
			cancelStream()
			_ = session.Stop()
			_ = peer.Close()
		}
	})

	answer, err := peer.CreateAnswer(nil)
	if err != nil {
		cancelStream()
		_ = session.Stop()
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("create answer: %v", err)})
		return
	}

	gatherComplete := webrtc.GatheringCompletePromise(peer)
	if err := peer.SetLocalDescription(answer); err != nil {
		cancelStream()
		_ = session.Stop()
		_ = peer.Close()
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("set local description: %v", err)})
		return
	}
	<-gatherComplete

	local := peer.LocalDescription()
	if local == nil {
		cancelStream()
		_ = session.Stop()
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
			_ = s.stream.Stop()
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
