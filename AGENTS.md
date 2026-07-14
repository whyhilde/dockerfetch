# dockerfetch

Neofetch-style CLI for Docker info.

## Build & Verify

```bash
go build -o dockerfetch .
go run .                    # full output
go run . --short            # essential only
go run . --json             # JSON output
go run . --no-color         # monochrome
go vet ./...
go test ./...
gofmt -s -w .               # format
```

## Package layout

```
main.go                     # flags, orchestration
internal/docker/client.go   # Docker SDK — Fetch() returns *DockerInfo
internal/docker/types.go    # DockerInfo, ContainerSummary, Line
internal/ascii/logo.go      # ASCII art constants (Fedora, Ubuntu, Debian, Arch, Alpine, macOS)
internal/display/render.go  # side-by-side table + JSON
internal/display/color.go   # ANSI wrappers, NoColor flag
internal/system/info.go     # host OS, kernel, uptime, OSID
```

## Key decisions

- **ASCII default:** OS logo detected from `/etc/os-release`; falls back to Ubuntu
- **Connection:** Unix socket + `DOCKER_HOST` env (via Docker SDK's `FromEnv`)
- **Flags:** `--short`, `--no-color`, `--json`
- **Colors:** ANSI by default; `--no-color` disables
- **Field order:** OS → Docker → API → Cgroup → Storage → Root → CPUs → Memory → Containers → Images → Volumes → Networks
- **Container statuses:** running / paused / stopped — from `docker info` directly
- **Networks filter:** built-in ingress networks excluded
- **OS detection:** `system.Get()` reads `/etc/os-release` for `ID=` field; logo selected by `ascii.Get(osID)`
- **Logos sourced from:** [fastfetch](https://github.com/fastfetch-cli/fastfetch) — backtick characters replaced with `'` to avoid Go raw string conflicts

## Notes

- Requires running Docker daemon
- Docker SDK pulled via `go mod tidy` (github.com/moby/moby/client + api)
- ANSI codes use `\033` literals (not `\e`)
- The `internal/system/info.go` reads `/proc` (Linux only); Darwin/Windows fall back to `runtime.GOOS`
