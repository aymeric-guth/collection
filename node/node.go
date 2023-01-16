package node

import (
	"fmt"
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

func (n *Node) Root() *Node {
	for {
		if n.Parent == nil {
			return n
		} else {
			n = n.Parent
		}
	}
}

func (n *Node) Update(path string) {
	q := strings.Split(path, "/")
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == "" {
			continue
		}
		for _, node := range n.Children {
			if node.Name == cur {
				n = node
				goto next
			}
		}
		n = New(cur, n)
	next:
		continue
	}
}

func (n *Node) Walk() *Node {
	q := []*Node{n}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		fmt.Printf("%+v\n", current)
		for _, node := range current.Children {
			if len(node.Children) > 0 {
				q = append(q, node)
			} else {
				fmt.Printf("leaf: %+v\n", node)
			}
		}
	}
	return nil
}

func (n *Node) Merge(node *Node) *Node {
	return nil
}
