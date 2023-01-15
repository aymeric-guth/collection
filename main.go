package main

import (
	"fmt"
	"node/node"
	"os"
	"strings"
)

func main() {
	vault := os.Getenv("OBSIDIAN_VAULT")
	var root *node.Node
	var previous *node.Node

	for _, p := range strings.Split(vault, "/") {
		if p == "" && root == nil {
			root = node.New("/", nil)
			previous = root
		} else {
			node := node.New(p, previous)
			previous = node
		}
	}

	leaf := root.Walk()
	if leaf != nil {
		fmt.Println(leaf)
	}
}
