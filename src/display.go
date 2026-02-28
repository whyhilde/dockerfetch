package src

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func FormatInfoLines(key string, value string, spaces string, cfg *Config) string {
	line := cfg.KeyCol + cfg.Bold + key + cfg.Sep + spaces + cfg.Reset + value
	return line
}

func CollectDockerInfo(cli *client.Client, ctx context.Context, cfg *Config) []string {
	var lines []string

	if ver, api, os, arch, err := GetDockerVersion(cli, ctx); err == nil {
		lines = append(lines, fmt.Sprint(FormatInfoLines("Version", ver, " ", cfg)))
		lines = append(lines, fmt.Sprint(FormatInfoLines("API", api, "     ", cfg)))
		lines = append(lines, fmt.Sprint(FormatInfoLines("OS", os, "      ", cfg)))
		lines = append(lines, fmt.Sprint(FormatInfoLines("Arch", arch, "    ", cfg)))
	}

	return lines
}

func Display(info []string, cfg *Config) {
	logo := []string{
		"               ###       ##    ",
		"       ### ### ###       ######",
		"   ### ### ### ### ###   ####  ",
		"############################   ",
		"############################   ",
		" #########################     ",
		"  #######################      ",
		"   ###################         ",
		"      ############             ",
	}

	maxLines := len(logo)
	if len(info) > maxLines {
		maxLines = len(info)
	}

	for i := 0; i < maxLines; i++ {
		logoLine := ""
		if i < len(logo) {
			logoLine = cfg.LogoCol + cfg.Bold + logo[i] + cfg.Reset
		}

		infoLine := ""
		if i < len(info) {
			infoLine = info[i]
		}

		fmt.Printf("%s   %s\n", logoLine, infoLine)
	}
}
