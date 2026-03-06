package src

import (
	"fmt"
	"strconv"
	"strings"
)

var colorMap = map[string]string{
	// base colors
	"normal":  "\033[0m",
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",

	// bright colors
	"bright-black":   "\033[90m",
	"bright-red":     "\033[91m",
	"bright-green":   "\033[92m",
	"bright-yellow":  "\033[93m",
	"bright-blue":    "\033[94m",
	"bright-magenta": "\033[95m",
	"bright-cyan":    "\033[96m",
	"bright-white":   "\033[97m",

	// aliases
	"gray":   "\033[90m",
	"grey":   "\033[90m",
	"orange": "\033[38;5;208m",
	"purple": "\033[38;5;93m",
	"pink":   "\033[38;5;205m",
}

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
)

func GetColorCode(color string) string {
	color = strings.ToLower(strings.TrimSpace(color))

	if code, ok := colorMap[color]; ok {
		return code
	}

	if strings.HasPrefix(color, "#") {
		return hexToANSI(color)
	}

	if strings.Contains(color, ",") {
		return rgbToANSI(color)
	}

	return ""
}

func hexToANSI(hex string) string {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) != 6 {
		return ""
	}

	r, err1 := strconv.ParseInt(hex[0:2], 16, 64)
	g, err2 := strconv.ParseInt(hex[2:4], 16, 64)
	b, err3 := strconv.ParseInt(hex[4:6], 16, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func rgbToANSI(rgb string) string {
	parts := strings.Split(rgb, ",")
	if len(parts) != 3 {
		return ""
	}

	r, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	g, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	b, err3 := strconv.Atoi(strings.TrimSpace(parts[2]))

	if err1 != nil || err2 != nil || err3 != nil {
		return ""
	}

	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
