package content

import (
	"testing"
)

func TestParsePageBasic(t *testing.T) {
	input := []byte(`---
title: Test Page
order: 1
tags: [foo, bar]
---

Some intro text.

@code bash
echo "hello"

@text
More explanation.
`)
	page, err := ParsePage(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if page.Meta.Title != "Test Page" {
		t.Errorf("title = %q, want %q", page.Meta.Title, "Test Page")
	}
	if page.Meta.Order != 1 {
		t.Errorf("order = %d, want 1", page.Meta.Order)
	}
	if len(page.Meta.Tags) != 2 {
		t.Fatalf("tags len = %d, want 2", len(page.Meta.Tags))
	}
	if len(page.Sections) != 3 {
		t.Fatalf("sections len = %d, want 3", len(page.Sections))
	}
	if page.Sections[0].Type != SectionText {
		t.Errorf("section[0] type = %v, want SectionText", page.Sections[0].Type)
	}
	if page.Sections[1].Type != SectionCode {
		t.Errorf("section[1] type = %v, want SectionCode", page.Sections[1].Type)
	}
	if page.Sections[1].Language != "bash" {
		t.Errorf("section[1] language = %q, want %q", page.Sections[1].Language, "bash")
	}
	if page.Sections[2].Type != SectionText {
		t.Errorf("section[2] type = %v, want SectionText", page.Sections[2].Type)
	}
}

func TestParsePageNoFrontmatter(t *testing.T) {
	input := []byte("Just some text.\n")
	page, err := ParsePage(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if page.Meta.Title != "" {
		t.Errorf("title = %q, want empty", page.Meta.Title)
	}
	if len(page.Sections) != 1 {
		t.Fatalf("sections len = %d, want 1", len(page.Sections))
	}
	if page.Sections[0].Content != "Just some text." {
		t.Errorf("content = %q, want %q", page.Sections[0].Content, "Just some text.")
	}
}

func TestParsePageFrontmatterOnly(t *testing.T) {
	input := []byte("---\ntitle: Meta Only\n---\n")
	page, err := ParsePage(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if page.Meta.Title != "Meta Only" {
		t.Errorf("title = %q, want %q", page.Meta.Title, "Meta Only")
	}
	if len(page.Sections) != 0 {
		t.Errorf("sections len = %d, want 0", len(page.Sections))
	}
}

func TestParsePageMultipleCodeBlocks(t *testing.T) {
	input := []byte(`---
title: Multi Code
---

Intro.

@code bash
echo "one"

@code python
print("two")

@text
Outro.
`)
	page, err := ParsePage(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(page.Sections) != 4 {
		t.Fatalf("sections len = %d, want 4", len(page.Sections))
	}
	if page.Sections[1].Language != "bash" {
		t.Errorf("section[1] language = %q, want bash", page.Sections[1].Language)
	}
	if page.Sections[2].Language != "python" {
		t.Errorf("section[2] language = %q, want python", page.Sections[2].Language)
	}
}

func TestParsePageCodeWithoutLang(t *testing.T) {
	input := []byte("---\ntitle: Bad\n---\n\n@code\necho hi\n")
	_, err := ParsePage(input)
	if err == nil {
		t.Fatal("expected error for @code without language")
	}
}

func TestParsePageEmpty(t *testing.T) {
	page, err := ParsePage([]byte{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(page.Sections) != 0 {
		t.Errorf("sections len = %d, want 0", len(page.Sections))
	}
}
