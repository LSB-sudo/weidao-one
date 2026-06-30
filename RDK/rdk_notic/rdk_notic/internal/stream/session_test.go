package stream

import (
	"bytes"
	"testing"
)

func TestAccessUnitBuilderKeepsRepeatedHeadersWithFollowingIDR(t *testing.T) {
	builder := newAccessUnitBuilder(nil)

	sps1 := []byte{0x67, 0x42, 0xe0, 0x1f}
	pps1 := []byte{0x68, 0xce, 0x06, 0xe2}
	idr1 := []byte{0x65, 0x80}
	sps2 := []byte{0x67, 0x42, 0xe0, 0x1f, 0xaa}
	pps2 := []byte{0x68, 0xce, 0x06, 0xe2, 0xbb}
	idr2 := []byte{0x65, 0x80, 0xcc}

	for _, nalu := range [][]byte{sps1, pps1, idr1} {
		if sample, ok := builder.Push(nalu); ok {
			t.Fatalf("unexpected flush before repeated SPS/PPS: %x", sample)
		}
	}

	sample1, ok := builder.Push(sps2)
	if !ok {
		t.Fatal("expected previous access unit flush when the next SPS arrives")
	}
	want1 := joinAnnexB([][]byte{sps1, pps1, idr1})
	if !bytes.Equal(sample1, want1) {
		t.Fatalf("unexpected first access unit\nwant: %x\n got: %x", want1, sample1)
	}

	for _, nalu := range [][]byte{pps2, idr2} {
		if sample, ok := builder.Push(nalu); ok {
			t.Fatalf("unexpected flush while building second access unit: %x", sample)
		}
	}

	sample2 := builder.Flush()
	want2 := joinAnnexB([][]byte{sps2, pps2, idr2})
	if !bytes.Equal(sample2, want2) {
		t.Fatalf("unexpected second access unit\nwant: %x\n got: %x", want2, sample2)
	}
}

func TestAccessUnitBuilderKeepsMultiSliceFrameTogether(t *testing.T) {
	builder := newAccessUnitBuilder(nil)

	firstSlice := []byte{0x41, 0x80}
	secondSliceSameFrame := []byte{0x41, 0x40}
	nextFrame := []byte{0x41, 0x80, 0xff}

	if sample, ok := builder.Push(firstSlice); ok {
		t.Fatalf("unexpected flush after first slice: %x", sample)
	}
	if sample, ok := builder.Push(secondSliceSameFrame); ok {
		t.Fatalf("unexpected flush inside the same frame: %x", sample)
	}

	sample1, ok := builder.Push(nextFrame)
	if !ok {
		t.Fatal("expected flush when the next frame starts")
	}
	want1 := joinAnnexB([][]byte{firstSlice, secondSliceSameFrame})
	if !bytes.Equal(sample1, want1) {
		t.Fatalf("unexpected multislice access unit\nwant: %x\n got: %x", want1, sample1)
	}

	sample2 := builder.Flush()
	want2 := joinAnnexB([][]byte{nextFrame})
	if !bytes.Equal(sample2, want2) {
		t.Fatalf("unexpected trailing access unit\nwant: %x\n got: %x", want2, sample2)
	}
}

func TestParseFirstMBInSlice(t *testing.T) {
	tests := []struct {
		name string
		nalu []byte
		want uint
	}{
		{name: "first slice", nalu: []byte{0x41, 0x80}, want: 0},
		{name: "non-zero first macroblock", nalu: []byte{0x41, 0x40}, want: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseFirstMBInSlice(tc.nalu)
			if err != nil {
				t.Fatalf("parseFirstMBInSlice() error = %v", err)
			}
			if got != tc.want {
				t.Fatalf("parseFirstMBInSlice() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestBuildVideoFilter(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{name: "default vertical", cfg: Config{VFlip: true}, want: "vflip"},
		{name: "horizontal only", cfg: Config{HFlip: true}, want: "hflip"},
		{name: "both", cfg: Config{VFlip: true, HFlip: true}, want: "vflip,hflip"},
		{name: "disabled", cfg: Config{}, want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := buildVideoFilter(tc.cfg); got != tc.want {
				t.Fatalf("buildVideoFilter() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestBuildFFmpegArgsIncludesFlipFilter(t *testing.T) {
	cfg := Config{
		DevicePath:  "/dev/video0",
		InputFormat: "mjpeg",
		Width:       640,
		Height:      480,
		FPS:         15,
		Bitrate:     700000,
		VFlip:       true,
	}
	args := buildFFmpegArgs(cfg)
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "-vf" && args[i+1] == "vflip" {
			return
		}
	}
	t.Fatalf("expected -vf vflip in args: %v", args)
}
