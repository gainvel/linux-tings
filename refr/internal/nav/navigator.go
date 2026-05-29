package nav

import (
	"fmt"
	"strings"

	"refr/internal/content"
)

type Navigator struct {
	root  *content.Node
	stack []*content.Node
}

func New(root *content.Node) *Navigator {
	return &Navigator{
		root:  root,
		stack: []*content.Node{root},
	}
}

func (n *Navigator) Current() *content.Node {
	return n.stack[len(n.stack)-1]
}

func (n *Navigator) Select(index int) error {
	cur := n.Current()
	if index < 0 || index >= len(cur.Children) {
		return fmt.Errorf("invalid selection: %d", index)
	}
	n.stack = append(n.stack, cur.Children[index])
	return nil
}

func (n *Navigator) Back() bool {
	if len(n.stack) <= 1 {
		return false
	}
	n.stack = n.stack[:len(n.stack)-1]
	return true
}

func (n *Navigator) AtRoot() bool {
	return len(n.stack) <= 1
}

func (n *Navigator) Breadcrumb() string {
	names := make([]string, len(n.stack))
	for i, node := range n.stack {
		names[i] = node.Name
	}
	return strings.Join(names, " > ")
}

func (n *Navigator) BreadcrumbFor(child *content.Node) string {
	return n.Breadcrumb() + " > " + child.Name
}
