package node

import (
	"strings"
)

type INode interface {
	Append(name string)
	Update(path string)
	Merge(node *Node) *Node
	Walk() *Node
}

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

func (n *Node) Walk() *Node {
	q := []*Node{n}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		for _, node := range current.Children {
			if len(node.Children) > 0 {
				q = append(q, node)
			} else {
				return node
			}
		}
	}
	return nil
}

func (n *Node) Merge(node *Node) *Node {
	return nil
}
