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

	info := src.CollectDockerInfo(cli, ctx)
	src.Display(info)

	os.Exit(0)
}
