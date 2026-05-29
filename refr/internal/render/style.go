package render

import (
	"refr/internal/config"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	cfg         config.ThemeConfig
	Header      lipgloss.Style
	Category    lipgloss.Style
	Page        lipgloss.Style
	Description lipgloss.Style
	Number      lipgloss.Style
	HelpBar     lipgloss.Style
	CodeBorder  lipgloss.Style
	CodeLabel   lipgloss.Style
	LineNumber  lipgloss.Style
	Text        lipgloss.Style
	Accent      lipgloss.Style
}

func NewStyles(theme config.ThemeConfig) *Styles {
	s := &Styles{cfg: theme}

	s.Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Header)).
		Bold(true)

	s.Category = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Category)).
		Bold(true)

	s.Page = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Page))

	s.Description = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Text)).
		Faint(true)

	s.Number = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Accent))

	s.HelpBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Border)).
		Faint(true)

	border := resolveBorder(theme.BorderStyle)
	s.CodeBorder = lipgloss.NewStyle().
		Border(border).
		BorderForeground(lipgloss.Color(theme.Border)).
		Padding(0, 1)

	s.CodeLabel = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Border)).
		Faint(true)

	s.LineNumber = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.LineNumber))

	s.Text = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Text))

	s.Accent = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Accent))

	return s
}

func resolveBorder(style string) lipgloss.Border {
	switch style {
	case "rounded":
		return lipgloss.RoundedBorder()
	case "thick":
		return lipgloss.ThickBorder()
	case "double":
		return lipgloss.DoubleBorder()
	default:
		return lipgloss.NormalBorder()
	}
}
