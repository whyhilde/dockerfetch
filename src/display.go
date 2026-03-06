package src

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func FormatInfoLines(key string, value string, cfg *Config) string {
	return fmt.Sprintf(
		"%s%-*s %s",
		cfg.KeyC+cfg.Bold,
		cfg.KeyWidth+1,
		key+cfg.Sep,
		cfg.ValueC+value,
	)
}

func CollectDockerInfo(cli *client.Client, ctx context.Context, cfg *Config) []string {
	var lines []string

	if ver, api, os, arch, err := GetDockerVersion(cli, ctx); err == nil {
		lines = append(lines, FormatInfoLines("Version", ver, cfg))
		lines = append(lines, FormatInfoLines("API", api, cfg))
		lines = append(lines, FormatInfoLines("OS", os, cfg))
		lines = append(lines, FormatInfoLines("Arch", arch, cfg))
	}

	if total, running, stopped, err := GetContainerStats(cli, ctx); err == nil {
		lines = append(
			lines,
			FormatInfoLines(
				"Containers",
				fmt.Sprint(total, " (running: ", running, ", stopped: ", stopped, ")"),
				cfg,
			),
		)
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
			logoLine = cfg.LogoC + cfg.Bold + logo[i] + cfg.Reset
		}

		infoLine := ""
		if i < len(info) {
			infoLine = info[i]
		}

		fmt.Printf("%s   %s\n", logoLine, infoLine)
	}
}
