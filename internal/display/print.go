package display

import (
	"fmt"

	"github.com/whyhilde/dockerfetch/internal/color"
	"github.com/whyhilde/dockerfetch/internal/config"
)

func PrintWithLogo(infoLines []string, cfg *config.Config) {
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
			logoLine = color.GetColorCode(cfg.LogoC) + color.Bold + logoLines[i] + color.Reset
		}

		infoLine := ""
		if i < len(infoLines) {
			infoLine = infoLines[i]
		}

		fmt.Printf("%s   %s\n", logoLine, infoLine)
	}
}
