package content

import (
	"testing"
	"testing/fstest"
)

func TestBuildTreeSinglePage(t *testing.T) {
	fsys := fstest.MapFS{
		"test.ref": {Data: []byte("---\ntitle: Test\norder: 1\n---\nHello.\n")},
	}
	root, err := BuildTree(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(root.Children) != 1 {
		t.Fatalf("children = %d, want 1", len(root.Children))
	}
	if root.Children[0].Name != "Test" {
		t.Errorf("child name = %q, want %q", root.Children[0].Name, "Test")
	}
}

func TestBuildTreeNestedCategories(t *testing.T) {
	fsys := fstest.MapFS{
		"cat/_index.ref":   {Data: []byte("---\ntitle: Category\norder: 1\n---\n")},
		"cat/page.ref":     {Data: []byte("---\ntitle: Page One\norder: 1\n---\nContent.\n")},
		"cat/sub/_index.ref": {Data: []byte("---\ntitle: Sub Category\norder: 2\n---\n")},
		"cat/sub/leaf.ref":   {Data: []byte("---\ntitle: Leaf\norder: 1\n---\nLeaf content.\n")},
	}
	root, err := BuildTree(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(root.Children) != 1 {
		t.Fatalf("root children = %d, want 1", len(root.Children))
	}
	cat := root.Children[0]
	if cat.Name != "Category" {
		t.Errorf("cat name = %q, want %q", cat.Name, "Category")
	}
	if len(cat.Children) != 2 {
		t.Fatalf("cat children = %d, want 2", len(cat.Children))
	}
}

func TestBuildTreeSortOrder(t *testing.T) {
	fsys := fstest.MapFS{
		"b.ref": {Data: []byte("---\ntitle: Bravo\norder: 2\n---\n")},
		"a.ref": {Data: []byte("---\ntitle: Alpha\norder: 1\n---\n")},
		"c.ref": {Data: []byte("---\ntitle: Charlie\norder: 1\n---\n")},
	}
	root, err := BuildTree(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(root.Children) != 3 {
		t.Fatalf("children = %d, want 3", len(root.Children))
	}
	names := []string{root.Children[0].Name, root.Children[1].Name, root.Children[2].Name}
	expected := []string{"Alpha", "Charlie", "Bravo"}
	for i, name := range names {
		if name != expected[i] {
			t.Errorf("child[%d] = %q, want %q", i, name, expected[i])
		}
	}
}

func TestBuildTreeNoIndex(t *testing.T) {
	fsys := fstest.MapFS{
		"my-dir/page.ref": {Data: []byte("---\ntitle: Page\n---\n")},
	}
	root, err := BuildTree(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(root.Children) != 1 {
		t.Fatalf("children = %d, want 1", len(root.Children))
	}
	if root.Children[0].Name != "My Dir" {
		t.Errorf("name = %q, want %q", root.Children[0].Name, "My Dir")
	}
}

func TestBuildTreeEmpty(t *testing.T) {
	fsys := fstest.MapFS{}
	root, err := BuildTree(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(root.Children) != 0 {
		t.Errorf("children = %d, want 0", len(root.Children))
	}
}
