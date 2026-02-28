package src

import (
	"context"

	"github.com/docker/docker/client"
)

func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func GetDockerVersion(
	cli *client.Client,
	ctx context.Context,
) (string, string, string, string, error) {
	version, err := cli.ServerVersion(ctx)
	if err != nil {
		return "", "", "", "", err
	}

	return version.Version, version.APIVersion, version.Os, version.Arch, nil
}
