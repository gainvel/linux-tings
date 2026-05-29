package content

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

func BuildTree(fsys fs.FS) (*Node, error) {
	root := &Node{
		Name:       "refr",
		Path:       ".",
		IsCategory: true,
	}

	dirs := map[string]*Node{".": root}
	indexMeta := map[string]*Frontmatter{}

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path == "." {
				return nil
			}
			node := &Node{
				Name:       titleCase(filepath.Base(path)),
				Path:       path,
				IsCategory: true,
			}
			dirs[path] = node
			return nil
		}

		if filepath.Ext(path) != ".ref" {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return nil
		}

		page, err := ParsePage(data)
		if err != nil {
			return nil
		}

		dir := filepath.Dir(path)
		name := filepath.Base(path)

		if name == "_index.ref" {
			indexMeta[dir] = &page.Meta
			return nil
		}

		leaf := &Node{
			Name: page.Meta.Title,
			Path: path,
			Page: page,
		}
		if leaf.Name == "" {
			leaf.Name = titleCase(strings.TrimSuffix(name, ".ref"))
		}
		leaf.Order = page.Meta.Order
		leaf.Description = page.Meta.Description

		parent, ok := dirs[dir]
		if !ok {
			parent = root
		}
		parent.Children = append(parent.Children, leaf)
		return nil
	})
	if err != nil {
		return nil, err
	}

	for path, meta := range indexMeta {
		node, ok := dirs[path]
		if !ok {
			continue
		}
		if meta.Title != "" {
			node.Name = meta.Title
		}
		node.Order = meta.Order
		node.Description = meta.Description
	}

	for path, node := range dirs {
		if path == "." {
			continue
		}
		parentPath := filepath.Dir(path)
		parent, ok := dirs[parentPath]
		if !ok {
			parent = root
		}
		parent.Children = append(parent.Children, node)
	}

	sortChildren(root)

	return root, nil
}

func sortChildren(node *Node) {
	sort.SliceStable(node.Children, func(i, j int) bool {
		a, b := node.Children[i], node.Children[j]
		if a.Order != b.Order {
			return a.Order < b.Order
		}
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)
	})
	for _, child := range node.Children {
		if child.IsCategory {
			sortChildren(child)
		}
	}
}

func titleCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})
	for i, w := range words {
		if len(w) == 0 {
			continue
		}
		runes := []rune(w)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
