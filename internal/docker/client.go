package docker

import (
	"context"

	"github.com/docker/go-units"
	"github.com/moby/moby/client"
)

func Fetch() (*DockerInfo, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx := context.Background()

	sysInfo, err := cli.Info(ctx, client.InfoOptions{})
	if err != nil {
		return nil, err
	}

	info := sysInfo.Info

	return &DockerInfo{
		Version:       info.ServerVersion,
		APIVersion:    cli.ClientVersion(),
		OS:            info.OperatingSystem,
		CgroupDriver:  info.CgroupDriver,
		StorageDriver: info.Driver,
		RootDir:       info.DockerRootDir,
		CPUs:          info.NCPU,
		TotalMemory:   units.BytesSize(float64(info.MemTotal)),
		Containers: ContainerSummary{
			Total:   info.Containers,
			Running: info.ContainersRunning,
			Paused:  info.ContainersPaused,
			Stopped: info.ContainersStopped,
		},
		Images:   info.Images,
		Volumes:  getVolumes(ctx, cli),
		Networks: getNetworks(ctx, cli),
	}, nil
}

func getVolumes(ctx context.Context, cli *client.Client) int {
	result, err := cli.VolumeList(ctx, client.VolumeListOptions{})
	if err != nil {
		return 0
	}
	return len(result.Items)
}

func getNetworks(ctx context.Context, cli *client.Client) int {
	result, err := cli.NetworkList(ctx, client.NetworkListOptions{})
	if err != nil {
		return 0
	}
	count := 0
	for _, n := range result.Items {
		if !n.Ingress {
			count++
		}
	}
	return count
}
