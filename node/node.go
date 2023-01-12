package node

import "strings"

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
