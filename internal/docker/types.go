package docker

import "fmt"

type DockerInfo struct {
	Version       string
	APIVersion    string
	OS            string
	Arch          string
	CgroupDriver  string
	StorageDriver string
	RootDir       string
	CPUs          int
	TotalMemory   string
	Containers    ContainerSummary
	Images        int
	Volumes       int
	Networks      int
}

type ContainerSummary struct {
	Total   int
	Running int
	Paused  int
	Stopped int
}

type Line struct {
	Key   string
	Value string
}

func (d DockerInfo) Lines() []Line {
	return []Line{
		{"OS", d.OS},
		{"Docker", d.Version},
		{"API version", d.APIVersion},
		{"Cgroup driver", d.CgroupDriver},
		{"Storage driver", d.StorageDriver},
		{"Docker Root", d.RootDir},
		{"CPUs", fmtInt(d.CPUs)},
		{"Total Memory", d.TotalMemory},
		{"Containers", d.Containers.String()},
		{"Images", fmtInt(d.Images)},
		{"Volumes", fmtInt(d.Volumes)},
		{"Networks", fmtInt(d.Networks)},
	}
}

func (c ContainerSummary) String() string {
	s := fmtInt(c.Total)
	var parts []string
	if c.Running > 0 {
		parts = append(parts, fmt.Sprintf("%d running", c.Running))
	}
	if c.Paused > 0 {
		parts = append(parts, fmt.Sprintf("%d paused", c.Paused))
	}
	if c.Stopped > 0 {
		parts = append(parts, fmt.Sprintf("%d stopped", c.Stopped))
	}
	if len(parts) > 0 {
		s += " (" + joinParts(parts) + ")"
	}
	return s
}

func fmtInt(n int) string {
	return fmt.Sprintf("%d", n)
}

func joinParts(parts []string) string {
	s := ""
	for i, p := range parts {
		if i > 0 {
			s += ", "
		}
		s += p
	}
	return s
}
