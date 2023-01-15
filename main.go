package node

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
}

func New(name string, parent *Node) *Node {
	node := new(Node)
	node.Name = name
	node.Parent = parent
	node.Children = []*Node{}
	if parent != nil {
		parent.Children = append(parent.Children, node)
	}
	return node
}

func (n *Node) Append(name string) {
	n.Children = append(n.Children, New(name, n))
}

func (n *Node) Update(path string) {
	previous := n
	for _, p := range strings.Split(path, "/") {
		if len(p) > 0 {
			node := New(p, previous)
			previous = node
		}
	}

}

func main() {
	vault := os.Getenv("OBSIDIAN_VAULT")
	root := New("/", nil)
	previous := root
	for _, p := range strings.Split(vault, "/") {
		if len(p) > 0 {
			node := New(p, previous)
			previous = node
		}
	}

	queue := []*Node{root}
	root.Insert(New("test", root))
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
