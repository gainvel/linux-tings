package render

import (
	"fmt"
	"strings"

	"refr/internal/content"
)

func RenderCategory(node *content.Node, breadcrumb string, showNumbers bool, styles *Styles) string {
	var b strings.Builder

	b.WriteString(styles.Header.Render(breadcrumb))
	b.WriteString("\n\n")

	if len(node.Children) == 0 {
		b.WriteString(styles.Description.Render("  No pages in this category"))
		b.WriteString("\n")
	} else {
		numWidth := len(fmt.Sprintf("%d", len(node.Children)))
		maxName := 0
		for _, child := range node.Children {
			if len(child.Name) > maxName {
				maxName = len(child.Name)
			}
		}

		for i, child := range node.Children {
			nameStyle := styles.Page
			if child.IsCategory {
				nameStyle = styles.Category
			}

			var line strings.Builder
			if showNumbers {
				num := fmt.Sprintf("%*d", numWidth, i+1)
				line.WriteString("  ")
				line.WriteString(styles.Number.Render(num))
				line.WriteString("  ")
			} else {
				line.WriteString("  ")
			}

			paddedName := fmt.Sprintf("%-*s", maxName, child.Name)
			line.WriteString(nameStyle.Render(paddedName))

			if child.Description != "" {
				line.WriteString("    ")
				line.WriteString(styles.Description.Render(child.Description))
			}

			b.WriteString(line.String())
			b.WriteString("\n")
		}
	}

	b.WriteString("\n")
	help := buildHelp(styles, "[number] select", "[b] back", "[q] quit")
	b.WriteString(help)

	return b.String()
}

func RenderPage(page *content.Page, breadcrumb string, styles *Styles, highlighter *Highlighter) string {
	var b strings.Builder

	b.WriteString(styles.Header.Render(breadcrumb))
	b.WriteString("\n")
	b.WriteString(styles.Header.Render(page.Meta.Title))
	b.WriteString("\n\n")

	for _, sec := range page.Sections {
		switch sec.Type {
		case content.SectionText:
			b.WriteString(styles.Text.Render(sec.Content))
			b.WriteString("\n\n")
		case content.SectionCode:
			highlighted := highlighter.Highlight(sec.Content, sec.Language)
			label := styles.CodeLabel.Render(" " + sec.Language + " ")
			b.WriteString(label)
			b.WriteString("\n")
			b.WriteString(styles.CodeBorder.Render(highlighted))
			b.WriteString("\n\n")
		}
	}

	help := buildHelp(styles, "[b] back", "[q] quit")
	b.WriteString(help)

	return b.String()
}

func buildHelp(styles *Styles, items ...string) string {
	sep := styles.Accent.Render(" · ")
	rendered := make([]string, len(items))
	for i, item := range items {
		rendered[i] = styles.HelpBar.Render(item)
	}
	return strings.Join(rendered, sep)
}
