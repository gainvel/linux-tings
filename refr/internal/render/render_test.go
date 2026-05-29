package render

import (
	"strings"
	"testing"

	"refr/internal/config"
	"refr/internal/content"
)

func defaultStyles() *Styles {
	return NewStyles(config.Default().Theme)
}

func defaultHighlighter() *Highlighter {
	return NewHighlighter(config.Default().Theme.Syntax)
}

func TestNewStylesDoesNotPanic(t *testing.T) {
	_ = defaultStyles()
}

func TestRenderCategoryBasic(t *testing.T) {
	node := &content.Node{
		Name:       "Root",
		IsCategory: true,
		Children: []*content.Node{
			{Name: "Alpha", Description: "First item", IsCategory: true},
			{Name: "Beta", Description: "Second item"},
		},
	}
	out := RenderCategory(node, "refr > Root", true, defaultStyles())
	if !strings.Contains(out, "1") {
		t.Error("expected numbered output")
	}
	if !strings.Contains(out, "Alpha") {
		t.Error("expected child name Alpha")
	}
	if !strings.Contains(out, "Beta") {
		t.Error("expected child name Beta")
	}
	if !strings.Contains(out, "[q] quit") {
		t.Error("expected help bar")
	}
}

func TestRenderCategoryEmpty(t *testing.T) {
	node := &content.Node{Name: "Empty", IsCategory: true}
	out := RenderCategory(node, "refr > Empty", true, defaultStyles())
	if !strings.Contains(out, "No pages") {
		t.Error("expected empty message")
	}
}

func TestRenderPageMixed(t *testing.T) {
	page := &content.Page{
		Meta: content.Frontmatter{Title: "Test Page"},
		Sections: []content.Section{
			{Type: content.SectionText, Content: "Some intro text."},
			{Type: content.SectionCode, Language: "bash", Content: "echo hello"},
			{Type: content.SectionText, Content: "More text."},
		},
	}
	out := RenderPage(page, "refr > Test", defaultStyles(), defaultHighlighter())
	if !strings.Contains(out, "Test Page") {
		t.Error("expected page title")
	}
	if !strings.Contains(out, "bash") {
		t.Error("expected language label")
	}
	if !strings.Contains(out, "[b] back") {
		t.Error("expected help bar")
	}
}

func TestHighlightBash(t *testing.T) {
	h := defaultHighlighter()
	out := h.Highlight("echo \"hello\"", "bash")
	if !strings.Contains(out, "\x1b[") {
		t.Error("expected ANSI escape codes in highlighted output")
	}
}

func TestHighlightUnknownLang(t *testing.T) {
	h := defaultHighlighter()
	input := "some random code"
	out := h.Highlight(input, "nonexistent-lang-xyz")
	if len(out) == 0 {
		t.Error("expected non-empty output for unknown language")
	}
}
