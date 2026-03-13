package docker

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

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

func GetOsName() string {
	switch runtime.GOOS {

	case "linux":
		file, err := os.Open("/etc/os-release")
		if err != nil {
			return "Linux"
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if value, ok := strings.CutPrefix(line, "PRETTY_NAME="); ok {
				value = strings.Trim(value, `"`)
				return value
			}
		}

	case "darwin":
		out, err := exec.Command("sw_vers", "-productVersion").Output()
		if err != nil {
			return "MacOS"
		}

		return "MacOS " + strings.TrimSpace(string(out))

	case "windows":
		out, err := exec.Command("cmd", "/c", "ver").Output()
		if err != nil {
			return "Windows"
		}

		return strings.TrimSpace(string(out))
	}

	return runtime.GOOS
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

func GetDockerInfo(
	cli *client.Client,
	ctx context.Context,
) (cgroup string, driver string, root string, err error) {
	info, err := cli.Info(ctx)
	if err != nil {
		return "", "", "", err
	}

	return info.CgroupDriver, info.Driver, info.DockerRootDir, nil
}

type DockerInfo struct {
	OsName            string
	Version           string
	ContainersTotal   int
	ContainersRunning int
	ContainersStopped int
	Volumes           int
	Images            int
	Networks          int
	CgroupDriver      string
	DockerRoot        string
	Driver            string
}

func FetchDockerInfo(cli *client.Client, ctx context.Context) (*DockerInfo, error) {
	info := &DockerInfo{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	wg.Add(1)
	go func() {
		defer wg.Done()

		version, err := GetDockerVersion(cli, ctx)
		if err != nil {
			mu.Lock()
			errors = append(errors, err)
			mu.Unlock()
			return
		}

		mu.Lock()
		info.Version = version
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		osname := GetOsName()

		mu.Lock()
		info.OsName = osname
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		containersTotal, containersRunning, containersStopped, err := GetContainerStats(cli, ctx)
		if err != nil {
			mu.Lock()
			errors = append(errors, err)
			mu.Unlock()
			return
		}

		mu.Lock()
		info.ContainersTotal = containersTotal
		info.ContainersRunning = containersRunning
		info.ContainersStopped = containersStopped
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		images, err := GetImagesStats(cli, ctx)
		if err != nil {
			mu.Lock()
			errors = append(errors, err)
			mu.Unlock()
			return
		}

		mu.Lock()
		info.Images = images
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		volumes, err := GetVolumesStats(cli, ctx)
		if err != nil {
			mu.Lock()
			errors = append(errors, err)
			mu.Unlock()
			return
		}

		mu.Lock()
		info.Volumes = volumes
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		networks, err := GetNetworksStats(cli, ctx)
		if err != nil {
			mu.Lock()
			errors = append(errors, err)
			mu.Unlock()
			return
		}

		mu.Lock()
		info.Networks = networks
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cgroup, driver, root, err := GetDockerInfo(cli, ctx)
		if err != nil {
			mu.Lock()
			errors = append(errors, err)
			mu.Unlock()
			return
		}

		mu.Lock()
		info.CgroupDriver = cgroup
		info.Driver = driver
		info.DockerRoot = root
		mu.Unlock()
	}()

	wg.Wait()

	if len(errors) > 0 {
		return info, errors[0]
	}

	return info, nil
}
