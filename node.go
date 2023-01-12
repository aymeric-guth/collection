package main

import (
	"fmt"
	"strings"
)

type Node struct {
	name     string
	parent   *Node
	children []*Node
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}
func (n *Node) AddParent(parent *Node) {
	n.parent = parent
}

func AddPath(path string, root *Node) {
	p := strings.Split(path, "/")
	fmt.Println(p)
}
