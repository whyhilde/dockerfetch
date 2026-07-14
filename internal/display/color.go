package display

import "runtime"

type Color string

const (
	Reset   Color = "\033[0m"
	Bold    Color = "\033[1m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"
)

var NoColor bool

func colorize(c Color, s string) string {
	if NoColor || runtime.GOOS == "windows" {
		return s
	}
	return string(c) + s + string(Reset)
}
