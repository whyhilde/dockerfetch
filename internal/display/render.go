package display

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/whyhilde/dockerfetch/internal/docker"
)

type Renderer struct {
	DockerInfo  *docker.DockerInfo
	ASCIIArt    string
	ShortMode   bool
	ShowImages  bool
	ShowVolumes bool
	JSONMode    bool
}

func (r *Renderer) Render() string {
	if r.JSONMode {
		return r.renderJSON()
	}
	return r.renderTable()
}

func (r *Renderer) renderJSON() string {
	data := map[string]interface{}{
		"version":        r.DockerInfo.Version,
		"api_version":    r.DockerInfo.APIVersion,
		"os":             r.DockerInfo.OS,
		"cgroup_driver":  r.DockerInfo.CgroupDriver,
		"storage_driver": r.DockerInfo.StorageDriver,
		"root_dir":       r.DockerInfo.RootDir,
		"cpus":           r.DockerInfo.CPUs,
		"total_memory":   r.DockerInfo.TotalMemory,
		"containers": map[string]int{
			"total":   r.DockerInfo.Containers.Total,
			"running": r.DockerInfo.Containers.Running,
			"paused":  r.DockerInfo.Containers.Paused,
			"stopped": r.DockerInfo.Containers.Stopped,
		},
		"images":   r.DockerInfo.Images,
		"volumes":  r.DockerInfo.Volumes,
		"networks": r.DockerInfo.Networks,
	}
	b, _ := json.MarshalIndent(data, "", "  ")
	return string(b)
}

func (r *Renderer) renderTable() string {
	var infoLines []docker.Line
	if r.ShortMode {
		infoLines = []docker.Line{
			{Key: "OS", Value: r.DockerInfo.OS},
			{Key: "Docker", Value: r.DockerInfo.Version},
			{Key: "API version", Value: r.DockerInfo.APIVersion},
			{Key: "Containers", Value: r.DockerInfo.Containers.String()},
			{Key: "Images", Value: fmtInt(r.DockerInfo.Images)},
			{Key: "Volumes", Value: fmtInt(r.DockerInfo.Volumes)},
		}
	} else {
		infoLines = r.DockerInfo.Lines()
	}

	asciiLines := strings.Split(strings.Trim(r.ASCIIArt, "\n"), "\n")

	return r.renderSideBySide(asciiLines, infoLines)
}

func (r *Renderer) renderSideBySide(ascii []string, info []docker.Line) string {
	maxKey := 0
	for _, l := range info {
		if len(l.Key) > maxKey {
			maxKey = len(l.Key)
		}
	}

	asciiWidth := 0
	for _, line := range ascii {
		clean := stripANSI(line)
		if len(clean) > asciiWidth {
			asciiWidth = len(clean)
		}
	}

	var infoCol []string
	for _, l := range info {
		key := colorize(Bold, l.Key+":")
		key = padRight(key, maxKey+2)
		value := colorize(Cyan, l.Value)
		infoCol = append(infoCol, key+" "+value)
	}

	for len(ascii) < len(infoCol) {
		ascii = append(ascii, "")
	}

	maxLines := len(ascii)
	if len(infoCol) > maxLines {
		maxLines = len(infoCol)
	}

	var out string
	for i := 0; i < maxLines; i++ {
		left := ""
		if i < len(ascii) {
			left = colorize(Blue, ascii[i])
		}
		right := ""
		if i < len(infoCol) {
			right = infoCol[i]
		}
		out += left + "   " + right + "\n"
	}

	return out
}

func padRight(s string, w int) string {
	if len(s) >= w {
		return s
	}
	return s + strings.Repeat(" ", w-len(s))
}

func fmtInt(n int) string {
	return fmt.Sprintf("%d", n)
}

func stripANSI(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\033' && i+1 < len(s) && s[i+1] == '[' {
			for j := i + 2; j < len(s); j++ {
				if s[j] == 'm' {
					i = j
					break
				}
			}
			continue
		}
		b.WriteByte(s[i])
	}
	return b.String()
}
