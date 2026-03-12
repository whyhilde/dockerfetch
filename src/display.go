package src

import (
	"fmt"
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

func CollectDockerInfo(info *DockerInfo, cfg *Config) []string {
	var lines []string

	lines = append(lines, FormatInfoLines("Version", info.Version, cfg))

	lines = append(lines, FormatInfoLines("OS", info.OsName, cfg))

	lines = append(
		lines,
		FormatInfoLines(
			"Containers",
			fmt.Sprint(
				info.ContainersTotal,
				" (running: ",
				info.ContainersRunning,
				", stopped: ",
				info.ContainersStopped,
				")",
			),
			cfg,
		),
	)

	lines = append(lines, FormatInfoLines("Images", fmt.Sprint(info.Images), cfg))

	lines = append(lines, FormatInfoLines("Volumes", fmt.Sprint(info.Volumes), cfg))

	lines = append(lines, FormatInfoLines("Networks", fmt.Sprint(info.Networks), cfg))

	lines = append(lines, FormatInfoLines("Cgroup driver", info.CgroupDriver, cfg))

	lines = append(lines, FormatInfoLines("Storage driver", info.Driver, cfg))

	lines = append(lines, FormatInfoLines("Root dir", info.DockerRoot, cfg))

	return lines
}

func Display(infoLines []string, cfg *Config) {
	logoLines := []string{
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

	maxLines := len(logoLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}

	for i := 0; i < maxLines; i++ {
		logoLine := ""
		if i < len(logoLines) {
			logoLine = GetColorCode(cfg.LogoC) + Bold + logoLines[i] + Reset
		}

		infoLine := ""
		if i < len(infoLines) {
			infoLine = infoLines[i]
		}

		fmt.Printf("%s   %s\n", logoLine, infoLine)
	}
}
