package main

import (
	"fmt"
	"goplayground/node"
	"os"
	"strings"
)

func main() {
	vault := os.Getenv("OBSIDIAN_VAULT")
	root := node.New("/", nil)
	previous := root
	for _, p := range strings.Split(vault, "/") {
		if len(p) > 0 {
			node := node.New(p, previous)
			previous = node
		}
	}

	queue := []*node.Node{root}
	root.Insert(node.New("test", root))
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, node := range current.Children {
			fmt.Printf("%v+\n", current)
			if len(node.Children) > 0 {
				queue = append(queue, node)
			}
		}
	}
}
