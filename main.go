package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/whyhilde/dockerfetch/internal/ascii"
	"github.com/whyhilde/dockerfetch/internal/display"
	dockerpkg "github.com/whyhilde/dockerfetch/internal/docker"
	"github.com/whyhilde/dockerfetch/internal/system"
)

type flags struct {
	noColor bool
	short   bool
	json    bool
	showImg bool
	showVol bool
}

func main() {
	var f flags
	flag.BoolVar(&f.noColor, "no-color", false, "disable colored output")
	flag.BoolVar(&f.short, "short", false, "show only essential info")
	flag.BoolVar(&f.json, "json", false, "output as JSON")
	flag.BoolVar(&f.showImg, "show-images", false, "show image details")
	flag.BoolVar(&f.showVol, "show-volumes", false, "show volume details")
	flag.Parse()

	if f.noColor {
		display.NoColor = true
	}

	sysInfo := system.Get()

	info, err := dockerpkg.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	info.OS = sysInfo.OS

	renderer := &display.Renderer{
		DockerInfo:  info,
		ASCIIArt:    ascii.Get(sysInfo.OSID),
		ShortMode:   f.short,
		ShowImages:  f.showImg,
		ShowVolumes: f.showVol,
		JSONMode:    f.json,
	}

	fmt.Print(renderer.Render())
}
