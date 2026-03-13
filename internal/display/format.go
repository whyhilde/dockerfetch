package display

import (
	"fmt"

	"github.com/whyhilde/dockerfetch/internal/color"
	"github.com/whyhilde/dockerfetch/internal/config"
	"github.com/whyhilde/dockerfetch/internal/docker"
)

func FormatInfoLine(key string, value string, cfg *config.Config) string {
	return fmt.Sprintf(
		"%s%-*s %s",
		color.GetColorCode(cfg.KeyC)+color.Bold,
		cfg.KeyWidth+1,
		key+cfg.Sep,
		color.GetColorCode(cfg.ValueC)+value,
	)
}

func CollectDockerInfo(info *docker.DockerInfo, cfg *config.Config) []string {
	var lines []string

	lines = append(lines, FormatInfoLine("Version", info.Version, cfg))

	lines = append(lines, FormatInfoLine("OS", info.OsName, cfg))

	lines = append(
		lines,
		FormatInfoLine(
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

	lines = append(lines, FormatInfoLine("Images", fmt.Sprint(info.Images), cfg))

	lines = append(lines, FormatInfoLine("Volumes", fmt.Sprint(info.Volumes), cfg))

	lines = append(lines, FormatInfoLine("Networks", fmt.Sprint(info.Networks), cfg))

	lines = append(lines, FormatInfoLine("Cgroup driver", info.CgroupDriver, cfg))

	lines = append(lines, FormatInfoLine("Storage driver", info.Driver, cfg))

	lines = append(lines, FormatInfoLine("Root dir", info.DockerRoot, cfg))

	return lines
}
