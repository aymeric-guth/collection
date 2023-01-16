package node

import (
	"fmt"
	"strings"
)

type INode interface {
	Update(path string)
	DFS()
	BFS()
	Search(string) bool
}

type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
	Distance int
}

func New(name string, parent *Node) *Node {
	node := new(Node)
	node.Name = name
	node.Parent = parent
	node.Children = []*Node{}
	if parent != nil {
		parent.Children = append(parent.Children, node)
	}
	n := node
	for n.Parent != nil {
		n = n.Parent
		node.Distance++
	}
	return node
}

func NewTree() *Node {
	return New("/", nil)
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

func (n *Node) Search(path string) bool {
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
		return false
	next:
		continue
	}
	return true
}

func (n *Node) DFS() {
	q := []*Node{n}
	for len(q) > 0 {
		current := q[len(q)-1]
		q = q[:len(q)-1]
		padding := strings.Repeat(" ", current.Distance*4)
		fmt.Printf("%s%s\n", padding, current.Name)
		q = append(q, current.Children...)
	}
}

func (n *Node) BFS() {
	q := []*Node{n}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		padding := strings.Repeat(" ", current.Distance)
		fmt.Printf("%s%s\n", padding, current.Name)
		q = append(q, current.Children...)
	}
}
