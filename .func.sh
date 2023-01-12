#!/bin/sh

run() {
	go run "$WORKSPACE"/main.go
	# ~/.local/share/nvim/mason/bin/golangci-lint run --fix=false --fast --out-format=json --path-prefix "$WORKSPACE"
}

build() {
	cd "$WORKSPACE" && go build
	cd "$OLDPWD" || return 1
}

readme() {
	editor "$WORKSPACE"/README.md
}
