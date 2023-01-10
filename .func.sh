#!/bin/sh

run() {
	go run "$WORKSPACE"/src/"$PROJECT_NAME"/main.go
}

build() {
	cd "$WORKSPACE"/src/"$PROJECT_NAME" && go build
	cd "$OLDPWD" || return 1
}

readme() {
	editor "$WORKSPACE"/README.md
}
