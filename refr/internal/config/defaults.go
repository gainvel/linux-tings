package config

import (
	"os"
	"path/filepath"
)

func Default() *Config {
	return &Config{
		PagesDir:    defaultPagesDir(),
		Pager:       "",
		ShowNumbers: true,
		Theme: ThemeConfig{
			BG:          "#181818",
			Text:        "#d8dcc9",
			Category:    "#5d8a78",
			Page:        "#8fb37a",
			Header:      "#6f9a6a",
			Accent:      "#b5c97a",
			Border:      "#3a423d",
			BorderStyle: "rounded",
			LineNumber:  "#8a9085",
			Syntax: SyntaxConfig{
				Base:     "monokai",
				Keyword:  "#6f9a6a",
				String:   "#8fb37a",
				Number:   "#b5c97a",
				Comment:  "italic #6a7866",
				Type:     "#5d8a78",
				Function: "#a3b89a",
				Operator: "#9ba87a",
				Variable: "#a3b89a",
				Error:    "bold #c47a4a",
			},
		},
	}
}

func defaultPagesDir() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "refr", "pages")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "refr", "pages")
}
