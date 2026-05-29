package content

import (
	"testing"
)

func buildTestTree() *Node {
	return &Node{
		Name:       "root",
		IsCategory: true,
		Children: []*Node{
			{
				Name: "OpenRC Services",
				Page: &Page{
					Meta:     Frontmatter{Title: "OpenRC Services", Tags: []string{"openrc", "init"}},
					Sections: []Section{{Type: SectionText, Content: "rc-service start stop"}},
				},
			},
			{
				Name: "Pacman Install",
				Page: &Page{
					Meta:     Frontmatter{Title: "Pacman Install", Tags: []string{"pacman", "packages"}},
					Sections: []Section{{Type: SectionText, Content: "pacman -S package"}},
				},
			},
			{
				Name: "Git Basics",
				Page: &Page{
					Meta:     Frontmatter{Title: "Git Basics", Tags: []string{"git", "vcs"}},
					Sections: []Section{{Type: SectionText, Content: "git init clone"}},
				},
			},
		},
	}
}

func TestSearchByTitle(t *testing.T) {
	results := Search(buildTestTree(), "openrc")
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
	if results[0].Name != "OpenRC Services" {
		t.Errorf("name = %q, want %q", results[0].Name, "OpenRC Services")
	}
}

func TestSearchByTag(t *testing.T) {
	results := Search(buildTestTree(), "vcs")
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
	if results[0].Name != "Git Basics" {
		t.Errorf("name = %q, want %q", results[0].Name, "Git Basics")
	}
}

func TestSearchByContent(t *testing.T) {
	results := Search(buildTestTree(), "pacman -S")
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
}

func TestSearchNoMatch(t *testing.T) {
	results := Search(buildTestTree(), "nonexistent")
	if len(results) != 0 {
		t.Errorf("results = %d, want 0", len(results))
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	results := Search(buildTestTree(), "OPENRC")
	if len(results) != 1 {
		t.Errorf("results = %d, want 1", len(results))
	}
}
