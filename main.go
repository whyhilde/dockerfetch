package main

import (
	"context"
	"fmt"
	"os"

	"github.com/whyhilde/dockerfetch/src"
)

func main() {
	cli, err := src.NewDockerClient()
	if err != nil {
		fmt.Printf("Can't connect to Docker: %v\n", err)
		os.Exit(1)
	}
	defer cli.Close()

	ctx := context.Background()

	cfg := src.SetOptions()
	info, err := src.FetchDockerInfo(cli, ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching docker info: %v\n", err)
		os.Exit(1)
	}

	lines := src.CollectDockerInfo(info, cfg)
	src.Display(lines, cfg)

	os.Exit(0)
}
