package system

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Info struct {
	OS     string
	OSID   string
	Kernel string
	Uptime string
	Host   string
}

func Get() Info {
	return Info{
		OS:     osName(),
		OSID:   osID(),
		Kernel: kernel(),
		Uptime: uptime(),
		Host:   hostname(),
	}
}

func osID() string {
	switch runtime.GOOS {
	case "linux":
		data, err := os.ReadFile("/etc/os-release")
		if err != nil {
			return "linux"
		}
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "ID=") && !strings.HasPrefix(line, "ID_LIKE=") {
				return strings.TrimSpace(strings.ToLower(strings.TrimPrefix(line, "ID=")))
			}
		}
		return "linux"
	case "darwin":
		return "macos"
	default:
		return runtime.GOOS
	}
}

func osName() string {
	switch runtime.GOOS {
	case "linux":
		data, err := os.ReadFile("/etc/os-release")
		if err == nil {
			for _, line := range strings.Split(string(data), "\n") {
				if strings.HasPrefix(line, "PRETTY_NAME=") {
					return strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
				}
			}
		}
		return "Linux"
	case "darwin":
		return "macOS"
	case "windows":
		return "Windows"
	default:
		return runtime.GOOS
	}
}

func kernel() string {
	data, err := os.ReadFile("/proc/sys/kernel/ostype")
	if err != nil {
		return runtime.GOOS
	}
	release, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return strings.TrimSpace(string(data))
	}
	return fmt.Sprintf("%s %s", strings.TrimSpace(string(data)), strings.TrimSpace(string(release)))
}

func uptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}
	var seconds float64
	fmt.Sscanf(string(data), "%f", &seconds)
	d := time.Duration(seconds) * time.Second
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	return fmt.Sprintf("%dh %dm", h, m)
}

func hostname() string {
	h, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return h
}
