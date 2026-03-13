package main

import (
	"context"
	"fmt"
	"os"

	"github.com/whyhilde/dockerfetch/internal/config"
	"github.com/whyhilde/dockerfetch/internal/display"
	"github.com/whyhilde/dockerfetch/internal/docker"
)

func main() {
	cli, err := docker.NewDockerClient()
	if err != nil {
		fmt.Printf("Can't connect to Docker: %v\n", err)
		os.Exit(1)
	}
	defer cli.Close()

	ctx := context.Background()

	cfg := config.SetOptions()
	info, err := docker.FetchDockerInfo(cli, ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching docker info: %v\n", err)
		os.Exit(1)
	}

	lines := display.CollectDockerInfo(info, cfg)
	display.PrintWithLogo(lines, cfg)

	os.Exit(0)
}
