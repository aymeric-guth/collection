#!/bin/sh

run() {
	go run "$WORKSPACE"/main.go
}

build() {
	cd "$WORKSPACE" && go build
	cd "$OLDPWD" || return 1
}

readme() {
	editor "$WORKSPACE"/README.md
}
