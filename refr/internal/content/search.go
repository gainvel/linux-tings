package content

import (
	"strings"
)

func Search(root *Node, query string) []*Node {
	query = strings.ToLower(query)
	var results []*Node
	collectMatches(root, query, &results)
	return results
}

func collectMatches(node *Node, query string, results *[]*Node) {
	if !node.IsCategory && node.Page != nil {
		if matchesQuery(node, query) {
			*results = append(*results, node)
		}
	}
	for _, child := range node.Children {
		collectMatches(child, query, results)
	}
}

func matchesQuery(node *Node, query string) bool {
	if strings.Contains(strings.ToLower(node.Page.Meta.Title), query) {
		return true
	}
	for _, tag := range node.Page.Meta.Tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}
	for _, sec := range node.Page.Sections {
		if strings.Contains(strings.ToLower(sec.Content), query) {
			return true
		}
	}
	return false
}
