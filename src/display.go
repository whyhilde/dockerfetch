package src

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func FormatInfoLines(key string, value string, cfg *Config) string {
	return fmt.Sprintf(
		"%s%-*s %s",
		GetColorCode(cfg.KeyC)+Bold,
		cfg.KeyWidth+1,
		key+cfg.Sep,
		GetColorCode(cfg.ValueC)+value,
	)
}

func CollectDockerInfo(cli *client.Client, ctx context.Context, cfg *Config) []string {
	var lines []string

	if version, err := GetDockerVersion(cli, ctx); err == nil {
		lines = append(lines, FormatInfoLines("Version", version, cfg))
	}

	lines = append(lines, FormatInfoLines("OS", GetOsName(), cfg))

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

	if images, err := GetImagesStats(cli, ctx); err == nil {
		lines = append(lines, FormatInfoLines("Images", fmt.Sprint(images), cfg))
	}

	if volumes, err := GetVolumesStats(cli, ctx); err == nil {
		lines = append(lines, FormatInfoLines("Volumes", fmt.Sprint(volumes), cfg))
	}

	if networks, err := GetNetworksStats(cli, ctx); err == nil {
		lines = append(lines, FormatInfoLines("Networks", fmt.Sprint(networks), cfg))
	}

	if cgroup, root, err := GetDockerInfo(cli, ctx); err == nil {
		lines = append(lines, FormatInfoLines("Cgroup dr", cgroup, cfg))
		lines = append(lines, FormatInfoLines("Root", root, cfg))
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
			logoLine = GetColorCode(cfg.LogoC) + Bold + logo[i] + Reset
		}

		infoLine := ""
		if i < len(info) {
			infoLine = info[i]
		}

		fmt.Printf("%s   %s\n", logoLine, infoLine)
	}
}
