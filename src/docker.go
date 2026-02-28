package src

import (
	"context"
	"fmt"

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

func CollectDockerInfo(cli *client.Client, ctx context.Context) []string {
	var lines []string

	c := "\033[31m"
	b := "\033[1m"
	r := "\033[0m"
	s := ":"

	if ver, api, os, arch, err := GetDockerVersion(cli, ctx); err == nil {
		lines = append(lines, fmt.Sprintf("%s%sVersion%s %s%s", c, b, s, r, ver))
		lines = append(lines, fmt.Sprintf("%s%sApi%s     %s%s", c, b, s, r, api))
		lines = append(lines, fmt.Sprintf("%s%sOs%s      %s%s", c, b, s, r, os))
		lines = append(lines, fmt.Sprintf("%s%sArch%s    %s%s", c, b, s, r, arch))
	}

	return lines
}
