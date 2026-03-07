package src

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func GetDockerVersion(cli *client.Client, ctx context.Context) (string, error) {
	version, err := cli.ServerVersion(ctx)
	if err != nil {
		return "", err
	}

	return version.Version, nil
}

func GetContainerStats(cli *client.Client, ctx context.Context) (int, int, int, error) {
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return 0, 0, 0, err
	}

	total := len(containers)
	running := 0
	stopped := 0

	for _, container := range containers {
		if container.State == "running" {
			running++
		} else {
			stopped++
		}
	}

	return total, running, stopped, nil
}

func GetImagesStats(cli *client.Client, ctx context.Context) (int, error) {
	images, err := cli.ImageList(ctx, image.ListOptions{All: true})
	if err != nil {
		return 0, err
	}

	return len(images), nil
}

func GetVolumesStats(cli *client.Client, ctx context.Context) (int, error) {
	volumes, err := cli.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(volumes.Volumes), nil
}

func GetNetworksStats(cli *client.Client, ctx context.Context) (int, error) {
	networks, err := cli.NetworkList(ctx, network.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(networks), nil
}
