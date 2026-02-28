bin := "dockerfetch"
install_path := "/usr/bin"

default:
  @just --list

install:
  go build -o {{bin}}
  sudo mv {{bin}} {{install_path}}

build:
  go build -o {{bin}}

uninstall:
  sudo rm -f {{install_path}}/{{bin}}

clean:
  rm -f {{bin}}
