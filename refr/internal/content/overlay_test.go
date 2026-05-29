package content

import (
	"testing"
)

func TestOverlayNilUser(t *testing.T) {
	base := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "A", Path: "a.ref"},
	}}
	result := Overlay(base, nil)
	if len(result.Children) != 1 {
		t.Fatalf("children = %d, want 1", len(result.Children))
	}
	if result.Children[0].Name != "A" {
		t.Errorf("name = %q, want %q", result.Children[0].Name, "A")
	}
}

func TestOverlayReplacePage(t *testing.T) {
	base := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "Page", Path: "page.ref", Page: &Page{Meta: Frontmatter{Title: "Old"}}},
	}}
	user := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "Page", Path: "page.ref", Page: &Page{Meta: Frontmatter{Title: "New"}}},
	}}
	result := Overlay(base, user)
	if result.Children[0].Page.Meta.Title != "New" {
		t.Errorf("title = %q, want %q", result.Children[0].Page.Meta.Title, "New")
	}
}

func TestOverlayExtendTree(t *testing.T) {
	base := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "A", Path: "a.ref"},
	}}
	user := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "B", Path: "b.ref"},
	}}
	result := Overlay(base, user)
	if len(result.Children) != 2 {
		t.Fatalf("children = %d, want 2", len(result.Children))
	}
}

func TestOverlayMergeCategory(t *testing.T) {
	base := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "Cat", Path: "cat", IsCategory: true, Children: []*Node{
			{Name: "A", Path: "a.ref"},
		}},
	}}
	user := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "Cat", Path: "cat", IsCategory: true, Children: []*Node{
			{Name: "B", Path: "b.ref"},
		}},
	}}
	result := Overlay(base, user)
	cat := result.Children[0]
	if len(cat.Children) != 2 {
		t.Fatalf("cat children = %d, want 2", len(cat.Children))
	}
}

func TestOverlayDoesNotMutateBase(t *testing.T) {
	base := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "A", Path: "a.ref"},
	}}
	user := &Node{Name: "root", IsCategory: true, Children: []*Node{
		{Name: "B", Path: "b.ref"},
	}}
	Overlay(base, user)
	if len(base.Children) != 1 {
		t.Errorf("base was mutated: children = %d, want 1", len(base.Children))
	}
}
