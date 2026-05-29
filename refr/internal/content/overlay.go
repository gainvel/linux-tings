package content

import (
	"sort"
	"strings"
)

func Overlay(base, user *Node) *Node {
	if user == nil {
		return base
	}
	merged := deepCopy(base)
	mergeInto(merged, user)
	return merged
}

func mergeInto(base, user *Node) {
	index := map[string]int{}
	for i, child := range base.Children {
		index[child.Path] = i
	}

	for _, uc := range user.Children {
		if i, ok := index[uc.Path]; ok {
			existing := base.Children[i]
			if existing.IsCategory && uc.IsCategory {
				if uc.Description != "" {
					existing.Description = uc.Description
				}
				if uc.Name != "" {
					existing.Name = uc.Name
				}
				if uc.Order != 0 {
					existing.Order = uc.Order
				}
				mergeInto(existing, uc)
			} else {
				base.Children[i] = deepCopy(uc)
			}
		} else {
			base.Children = append(base.Children, deepCopy(uc))
		}
	}

	sort.SliceStable(base.Children, func(i, j int) bool {
		a, b := base.Children[i], base.Children[j]
		if a.Order != b.Order {
			return a.Order < b.Order
		}
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)
	})
}

func deepCopy(n *Node) *Node {
	cp := *n
	if n.Page != nil {
		pageCopy := *n.Page
		pageCopy.Sections = make([]Section, len(n.Page.Sections))
		copy(pageCopy.Sections, n.Page.Sections)
		cp.Page = &pageCopy
	}
	if len(n.Children) > 0 {
		cp.Children = make([]*Node, len(n.Children))
		for i, child := range n.Children {
			cp.Children[i] = deepCopy(child)
		}
	}
	return &cp
}
