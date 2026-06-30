package device

import (
	"context"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type ToolStatus struct {
	FFmpeg  bool `json:"ffmpeg"`
	V4L2Ctl bool `json:"v4l2Ctl"`
}

type DeviceInfo struct {
	Path           string `json:"path"`
	Name           string `json:"name,omitempty"`
	IsVideoCapture bool   `json:"isVideoCapture"`
	Summary        string `json:"summary,omitempty"`
}

type Discovery struct {
	Requested string       `json:"requested,omitempty"`
	Selected  string       `json:"selected,omitempty"`
	Available bool         `json:"available"`
	Devices   []DeviceInfo `json:"devices"`
	RawList   string       `json:"rawList,omitempty"`
	Error     string       `json:"error,omitempty"`
}

func Tools() ToolStatus {
	return ToolStatus{
		FFmpeg:  hasCommand("ffmpeg"),
		V4L2Ctl: hasCommand("v4l2-ctl"),
	}
}

func Discover(preferred string) Discovery {
	result := Discovery{Requested: preferred}

	paths, err := filepath.Glob("/dev/video*")
	if err != nil {
		result.Error = err.Error()
		return result
	}
	sort.Strings(paths)

	rawList, _ := runCommand(2*time.Second, "v4l2-ctl", "--list-devices")
	result.RawList = strings.TrimSpace(rawList)
	names := parseListDevices(rawList)

	for _, path := range paths {
		info := DeviceInfo{Path: path, Name: names[path]}
		probe, probeErr := runCommand(2*time.Second, "v4l2-ctl", "-d", path, "--all")
		if probeErr == nil {
			info.Summary = summarizeProbe(probe)
			info.IsVideoCapture = looksLikeVideoCapture(probe)
		}
		result.Devices = append(result.Devices, info)
	}

	if preferred != "" {
		for _, info := range result.Devices {
			if info.Path == preferred && info.IsVideoCapture {
				result.Selected = info.Path
				result.Available = true
				return result
			}
		}
		result.Error = "preferred device unavailable or not a capture node"
		return result
	}

	for _, info := range result.Devices {
		if info.IsVideoCapture {
			result.Selected = info.Path
			result.Available = true
			return result
		}
	}

	if len(result.Devices) == 0 && result.Error == "" {
		result.Error = "no /dev/video* devices found"
	}
	return result
}

func hasCommand(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func runCommand(timeout time.Duration, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func parseListDevices(raw string) map[string]string {
	result := make(map[string]string)
	current := ""

	for _, line := range strings.Split(raw, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if !strings.HasPrefix(line, "\t") && strings.HasSuffix(trimmed, ":") {
			current = strings.TrimSuffix(trimmed, ":")
			continue
		}
		if strings.HasPrefix(trimmed, "/dev/video") {
			result[trimmed] = current
		}
	}
	return result
}

func looksLikeVideoCapture(probe string) bool {
	inDeviceCaps := false
	for _, line := range strings.Split(probe, "\n") {
		trimmed := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(trimmed, "Device Caps"):
			inDeviceCaps = true
		case strings.HasPrefix(trimmed, "Media Driver Info:"),
			strings.HasPrefix(trimmed, "Priority:"),
			strings.HasPrefix(trimmed, "Format "),
			strings.HasPrefix(trimmed, "Selection "),
			strings.HasPrefix(trimmed, "Streaming Parameters"),
			strings.HasPrefix(trimmed, "User Controls"),
			strings.HasPrefix(trimmed, "Camera Controls"):
			inDeviceCaps = false
		}
		if inDeviceCaps && (trimmed == "Video Capture" || trimmed == "Video Capture Multiplanar") {
			return true
		}
	}
	return false
}

func summarizeProbe(probe string) string {
	lines := strings.Split(probe, "\n")
	picked := make([]string, 0, 5)
	seen := make(map[string]struct{})
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		switch {
		case strings.HasPrefix(line, "Card type"):
			if _, ok := seen[line]; !ok {
				picked = append(picked, line)
				seen[line] = struct{}{}
			}
		case strings.HasPrefix(line, "Bus info"):
			if _, ok := seen[line]; !ok {
				picked = append(picked, line)
				seen[line] = struct{}{}
			}
		case strings.HasPrefix(line, "Width/Height"):
			if _, ok := seen[line]; !ok {
				picked = append(picked, line)
				seen[line] = struct{}{}
			}
		case strings.HasPrefix(line, "Pixel Format"):
			if _, ok := seen[line]; !ok {
				picked = append(picked, line)
				seen[line] = struct{}{}
			}
		case strings.HasPrefix(line, "Frames per second"):
			if _, ok := seen[line]; !ok {
				picked = append(picked, line)
				seen[line] = struct{}{}
			}
		}
	}
	return strings.Join(picked, "; ")
}
