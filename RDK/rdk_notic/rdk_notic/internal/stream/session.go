package stream

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

type Config struct {
	DevicePath  string `json:"devicePath,omitempty"`
	InputFormat string `json:"inputFormat,omitempty"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FPS         int    `json:"fps"`
	Bitrate     int    `json:"bitrate"`
	VFlip       bool   `json:"vflip"`
	HFlip       bool   `json:"hflip"`
}

type RuntimeInfo struct {
	Config        Config    `json:"config"`
	FFmpegArgs    []string  `json:"ffmpegArgs"`
	FFmpegCommand string    `json:"ffmpegCommand"`
	FrameDuration string    `json:"frameDuration"`
	GeneratedAt   time.Time `json:"generatedAt"`
}

type Session struct {
	cfg      Config
	cmd      *exec.Cmd
	cancel   context.CancelFunc
	done     chan struct{}
	stderr   *tailBuffer
	errMu    sync.Mutex
	runErr   error
	stopOnce sync.Once
}

func Start(parent context.Context, cfg Config, track *webrtc.TrackLocalStaticSample) (*Session, error) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return nil, fmt.Errorf("ffmpeg not available: %w", err)
	}

	cfg = normalizeConfig(cfg)
	ctx, cancel := context.WithCancel(parent)

	cmd := exec.CommandContext(ctx, "ffmpeg", buildFFmpegArgs(cfg)...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, fmt.Errorf("ffmpeg stdout pipe: %w", err)
	}

	stderr := newTailBuffer(16 << 10)
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		cancel()
		return nil, fmt.Errorf("start ffmpeg: %w", err)
	}

	session := &Session{
		cfg:    cfg,
		cmd:    cmd,
		cancel: cancel,
		done:   make(chan struct{}),
		stderr: stderr,
	}

	go session.run(ctx, stdout, track)
	return session, nil
}

func Describe(cfg Config) RuntimeInfo {
	cfg = normalizeConfig(cfg)
	frameDuration := time.Second / time.Duration(max(1, cfg.FPS))
	args := buildFFmpegArgs(cfg)

	return RuntimeInfo{
		Config:        cfg,
		FFmpegArgs:    append([]string(nil), args...),
		FFmpegCommand: "ffmpeg " + strings.Join(args, " "),
		FrameDuration: frameDuration.String(),
		GeneratedAt:   time.Now().UTC(),
	}
}

func (s *Session) run(ctx context.Context, stdout io.ReadCloser, track *webrtc.TrackLocalStaticSample) {
	defer close(s.done)
	defer stdout.Close()

	parser := newAnnexBParser(stdout)
	frameDuration := time.Second / time.Duration(max(1, s.cfg.FPS))
	stats := newH264StreamStats()
	builder := newAccessUnitBuilder(stats)
	defer log.Printf("h264 stream summary: %s", stats.summary())

	writeAccessUnit := func(data []byte) error {
		if len(data) == 0 {
			return nil
		}
		stats.accessUnits++
		return track.WriteSample(media.Sample{Data: data, Duration: frameDuration})
	}

	for {
		nalu, err := parser.NextNALU()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if flushErr := writeAccessUnit(builder.Flush()); flushErr != nil && ctx.Err() == nil {
					s.setErr(flushErr)
				}
				break
			}
			if ctx.Err() == nil {
				s.setErr(fmt.Errorf("read H264 stream: %w", err))
			}
			break
		}
		if len(nalu) == 0 {
			continue
		}

		stats.observeNALU(nalu[0] & 0x1F)
		if sample, ok := builder.Push(nalu); ok {
			if err := writeAccessUnit(sample); err != nil {
				if ctx.Err() == nil {
					s.setErr(fmt.Errorf("write video sample: %w", err))
				}
				break
			}
		}
	}

	waitErr := s.cmd.Wait()
	if waitErr != nil && ctx.Err() == nil {
		s.setErr(fmt.Errorf("ffmpeg exited: %w: %s", waitErr, s.stderr.String()))
	}
}

func (s *Session) Stop() error {
	s.stopOnce.Do(func() {
		s.cancel()
	})
	<-s.done
	return s.Err()
}

func (s *Session) Wait() error {
	<-s.done
	return s.Err()
}

func (s *Session) Err() error {
	s.errMu.Lock()
	defer s.errMu.Unlock()
	return s.runErr
}

func (s *Session) setErr(err error) {
	if err == nil {
		return
	}
	s.errMu.Lock()
	defer s.errMu.Unlock()
	if s.runErr == nil {
		s.runErr = err
	}
}

func normalizeConfig(cfg Config) Config {
	if cfg.InputFormat == "" {
		cfg.InputFormat = "mjpeg"
	}
	if cfg.Width <= 0 {
		cfg.Width = 640
	}
	if cfg.Height <= 0 {
		cfg.Height = 480
	}
	if cfg.FPS <= 0 {
		cfg.FPS = 15
	}
	if cfg.Bitrate <= 0 {
		cfg.Bitrate = 700000
	}
	return cfg
}

func buildVideoFilter(cfg Config) string {
	filters := make([]string, 0, 2)
	if cfg.VFlip {
		filters = append(filters, "vflip")
	}
	if cfg.HFlip {
		filters = append(filters, "hflip")
	}
	return strings.Join(filters, ",")
}

func buildFFmpegArgs(cfg Config) []string {
	fps := max(1, cfg.FPS)
	keyint := fps
	bitrateK := max(300, cfg.Bitrate/1000)
	bufsizeK := max(300, bitrateK/2)
	args := []string{
		"-hide_banner",
		"-loglevel", "error",
		"-fflags", "nobuffer",
		"-flags", "low_delay",
		"-f", "v4l2",
		"-thread_queue_size", "64",
	}
	if cfg.InputFormat != "" {
		args = append(args, "-input_format", cfg.InputFormat)
	}
	args = append(args,
		"-framerate", strconv.Itoa(fps),
		"-video_size", fmt.Sprintf("%dx%d", cfg.Width, cfg.Height),
		"-i", cfg.DevicePath,
		"-an",
	)
	if filter := buildVideoFilter(cfg); filter != "" {
		args = append(args, "-vf", filter)
	}
	args = append(args,
		"-pix_fmt", "yuv420p",
		"-c:v", "libx264",
		"-preset", "ultrafast",
		"-tune", "zerolatency",
		"-profile:v", "baseline",
		"-bf", "0",
		"-g", strconv.Itoa(keyint),
		"-keyint_min", strconv.Itoa(keyint),
		"-sc_threshold", "0",
		"-x264-params", fmt.Sprintf("annexb=1:repeat-headers=1:slices=1:scenecut=0:open-gop=0:bframes=0:sync-lookahead=0:rc-lookahead=0:keyint=%d:min-keyint=%d", keyint, keyint),
		"-b:v", fmt.Sprintf("%dk", bitrateK),
		"-maxrate", fmt.Sprintf("%dk", bitrateK),
		"-bufsize", fmt.Sprintf("%dk", bufsizeK),
		"-f", "h264",
		"pipe:1",
	)
	return args
}

func isVCL(nalType byte) bool {
	return nalType >= 1 && nalType <= 5
}

func isAccessUnitPrefixNALU(nalType byte) bool {
	switch nalType {
	case 6, 7, 8, 9:
		return true
	default:
		return false
	}
}

func joinAnnexB(nalus [][]byte) []byte {
	total := 0
	for _, nalu := range nalus {
		total += len(nalu) + 4
	}

	out := make([]byte, 0, total)
	for _, nalu := range nalus {
		out = append(out, 0x00, 0x00, 0x00, 0x01)
		out = append(out, nalu...)
	}
	return out
}

type accessUnitBuilder struct {
	nalus  [][]byte
	hasVCL bool
	stats  *h264StreamStats
}

func newAccessUnitBuilder(stats *h264StreamStats) *accessUnitBuilder {
	return &accessUnitBuilder{stats: stats}
}

func (b *accessUnitBuilder) Push(nalu []byte) ([]byte, bool) {
	if len(nalu) == 0 {
		return nil, false
	}

	var sample []byte
	if b.shouldStartNewAccessUnit(nalu) {
		sample = b.Flush()
	}

	copied := append([]byte(nil), nalu...)
	b.nalus = append(b.nalus, copied)
	if isVCL(nalu[0] & 0x1F) {
		b.hasVCL = true
	}

	return sample, len(sample) > 0
}

func (b *accessUnitBuilder) Flush() []byte {
	if len(b.nalus) == 0 || !b.hasVCL {
		b.reset()
		return nil
	}
	data := joinAnnexB(b.nalus)
	b.reset()
	return data
}

func (b *accessUnitBuilder) reset() {
	b.nalus = nil
	b.hasVCL = false
}

func (b *accessUnitBuilder) shouldStartNewAccessUnit(nalu []byte) bool {
	if !b.hasVCL || len(nalu) == 0 {
		return false
	}

	nalType := nalu[0] & 0x1F
	if isAccessUnitPrefixNALU(nalType) {
		return true
	}
	if !isVCL(nalType) {
		return false
	}

	startsPicture, err := sliceStartsPicture(nalu)
	if err != nil {
		if b.stats != nil {
			b.stats.sliceHeaderParseErrors++
		}
		return true
	}
	return startsPicture
}

func sliceStartsPicture(nalu []byte) (bool, error) {
	firstMBInSlice, err := parseFirstMBInSlice(nalu)
	if err != nil {
		return false, err
	}
	return firstMBInSlice == 0, nil
}

func parseFirstMBInSlice(nalu []byte) (uint, error) {
	if len(nalu) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	rbsp := make([]byte, 0, len(nalu)-1)
	zeroCount := 0
	for _, b := range nalu[1:] {
		if zeroCount >= 2 && b == 0x03 {
			zeroCount = 0
			continue
		}
		rbsp = append(rbsp, b)
		if b == 0x00 {
			zeroCount++
		} else {
			zeroCount = 0
		}
	}

	reader := bitReader{data: rbsp}
	return reader.readUE()
}

type bitReader struct {
	data []byte
	bit  int
}

func (r *bitReader) readBit() (uint, error) {
	if r.bit >= len(r.data)*8 {
		return 0, io.ErrUnexpectedEOF
	}
	byteIndex := r.bit / 8
	shift := 7 - (r.bit % 8)
	r.bit++
	return uint((r.data[byteIndex] >> shift) & 0x01), nil
}

func (r *bitReader) readBits(n int) (uint, error) {
	var value uint
	for i := 0; i < n; i++ {
		bit, err := r.readBit()
		if err != nil {
			return 0, err
		}
		value = (value << 1) | bit
	}
	return value, nil
}

func (r *bitReader) readUE() (uint, error) {
	leadingZeroBits := 0
	for {
		bit, err := r.readBit()
		if err != nil {
			return 0, err
		}
		if bit == 1 {
			break
		}
		leadingZeroBits++
	}
	if leadingZeroBits == 0 {
		return 0, nil
	}
	info, err := r.readBits(leadingZeroBits)
	if err != nil {
		return 0, err
	}
	return (1 << leadingZeroBits) - 1 + info, nil
}

type h264StreamStats struct {
	accessUnits            uint64
	nalCounts              [32]uint64
	sliceHeaderParseErrors uint64
	lastIDR                time.Time
	lastSPS                time.Time
	lastPPS                time.Time
}

func newH264StreamStats() *h264StreamStats {
	return &h264StreamStats{}
}

func (s *h264StreamStats) observeNALU(nalType byte) {
	if int(nalType) < len(s.nalCounts) {
		s.nalCounts[nalType]++
	}
	now := time.Now().UTC()
	switch nalType {
	case 5:
		s.lastIDR = now
	case 7:
		s.lastSPS = now
	case 8:
		s.lastPPS = now
	}
}

func (s *h264StreamStats) summary() string {
	return fmt.Sprintf(
		"accessUnits=%d nalus{nonIDR=%d idr=%d sei=%d sps=%d pps=%d aud=%d} lastIDR=%s lastSPS=%s lastPPS=%s sliceHeaderParseErrors=%d",
		s.accessUnits,
		s.nalCounts[1],
		s.nalCounts[5],
		s.nalCounts[6],
		s.nalCounts[7],
		s.nalCounts[8],
		s.nalCounts[9],
		formatStatTime(s.lastIDR),
		formatStatTime(s.lastSPS),
		formatStatTime(s.lastPPS),
		s.sliceHeaderParseErrors,
	)
}

func formatStatTime(ts time.Time) string {
	if ts.IsZero() {
		return "never"
	}
	return ts.Format(time.RFC3339)
}

type annexBParser struct {
	r   io.Reader
	buf []byte
	eof bool
}

func newAnnexBParser(r io.Reader) *annexBParser {
	return &annexBParser{r: r, buf: make([]byte, 0, 256*1024)}
}

func (p *annexBParser) NextNALU() ([]byte, error) {
	for {
		start, startLen := findStartCode(p.buf, 0)
		if start >= 0 {
			nextStart, _ := findStartCode(p.buf, start+startLen)
			if nextStart >= 0 {
				nalu := append([]byte(nil), p.buf[start+startLen:nextStart]...)
				p.buf = append([]byte(nil), p.buf[nextStart:]...)
				if len(nalu) == 0 {
					continue
				}
				return nalu, nil
			}
			if p.eof {
				nalu := append([]byte(nil), p.buf[start+startLen:]...)
				p.buf = nil
				if len(nalu) == 0 {
					return nil, io.EOF
				}
				return nalu, nil
			}
		} else if p.eof {
			return nil, io.EOF
		}

		chunk := make([]byte, 32*1024)
		n, err := p.r.Read(chunk)
		if n > 0 {
			p.buf = append(p.buf, chunk[:n]...)
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				p.eof = true
				continue
			}
			return nil, err
		}
	}
}

func findStartCode(buf []byte, from int) (int, int) {
	for i := from; i+3 <= len(buf); i++ {
		if i+4 <= len(buf) && buf[i] == 0x00 && buf[i+1] == 0x00 && buf[i+2] == 0x00 && buf[i+3] == 0x01 {
			return i, 4
		}
		if buf[i] == 0x00 && buf[i+1] == 0x00 && buf[i+2] == 0x01 {
			return i, 3
		}
	}
	return -1, 0
}

type tailBuffer struct {
	mu    sync.Mutex
	limit int
	buf   bytes.Buffer
}

func newTailBuffer(limit int) *tailBuffer {
	return &tailBuffer{limit: limit}
}

func (b *tailBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(p) >= b.limit {
		b.buf.Reset()
		_, _ = b.buf.Write(p[len(p)-b.limit:])
		return len(p), nil
	}

	if b.buf.Len()+len(p) > b.limit {
		keep := b.limit - len(p)
		current := append([]byte(nil), b.buf.Bytes()...)
		b.buf.Reset()
		_, _ = b.buf.Write(current[len(current)-keep:])
	}
	_, _ = b.buf.Write(p)
	return len(p), nil
}

func (b *tailBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return string(bytes.TrimSpace(b.buf.Bytes()))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
