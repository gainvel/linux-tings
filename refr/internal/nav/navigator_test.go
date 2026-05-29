package nav

import (
	"testing"

	"refr/internal/content"
)

func testTree() *content.Node {
	return &content.Node{
		Name:       "root",
		IsCategory: true,
		Children: []*content.Node{
			{Name: "A", IsCategory: true, Children: []*content.Node{
				{Name: "A1"},
			}},
			{Name: "B"},
		},
	}
}

func TestNewNavigator(t *testing.T) {
	n := New(testTree())
	if n.Current().Name != "root" {
		t.Errorf("current = %q, want %q", n.Current().Name, "root")
	}
}

func TestSelect(t *testing.T) {
	n := New(testTree())
	if err := n.Select(0); err != nil {
		t.Fatalf("select error: %v", err)
	}
	if n.Current().Name != "A" {
		t.Errorf("current = %q, want %q", n.Current().Name, "A")
	}
}

func TestSelectOutOfBounds(t *testing.T) {
	n := New(testTree())
	if err := n.Select(5); err == nil {
		t.Error("expected error for out-of-bounds selection")
	}
}

func TestBack(t *testing.T) {
	n := New(testTree())
	n.Select(0)
	if !n.Back() {
		t.Error("expected Back to return true")
	}
	if n.Current().Name != "root" {
		t.Errorf("current = %q, want %q", n.Current().Name, "root")
	}
}

func TestBackAtRoot(t *testing.T) {
	n := New(testTree())
	if n.Back() {
		t.Error("expected Back at root to return false")
	}
}

func TestAtRoot(t *testing.T) {
	n := New(testTree())
	if !n.AtRoot() {
		t.Error("expected AtRoot to be true")
	}
	n.Select(0)
	if n.AtRoot() {
		t.Error("expected AtRoot to be false after Select")
	}
}

func TestBreadcrumb(t *testing.T) {
	n := New(testTree())
	n.Select(0)
	bc := n.Breadcrumb()
	if bc != "root > A" {
		t.Errorf("breadcrumb = %q, want %q", bc, "root > A")
	}
}

func TestBreadcrumbFor(t *testing.T) {
	n := New(testTree())
	child := testTree().Children[1]
	bc := n.BreadcrumbFor(child)
	if bc != "root > B" {
		t.Errorf("breadcrumb = %q, want %q", bc, "root > B")
	}
}
