package src

import (
	"fmt"
)

func Display(info []string) {
	logo := []string {
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
			logoLine = logo[i]
		}

		infoLine := ""
		if i < len(info) {
			infoLine = info[i]
		}

		fmt.Printf("%s   %s\n", logoLine, infoLine)
	}

	fmt.Println()
}
