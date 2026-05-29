package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	PagesDir    string      `toml:"pages_dir"`
	Pager       string      `toml:"pager"`
	ShowNumbers bool        `toml:"show_numbers"`
	Theme       ThemeConfig `toml:"theme"`
}

type ThemeConfig struct {
	BG          string       `toml:"bg"`
	Text        string       `toml:"text"`
	Category    string       `toml:"category"`
	Page        string       `toml:"page"`
	Header      string       `toml:"header"`
	Accent      string       `toml:"accent"`
	Border      string       `toml:"border"`
	BorderStyle string       `toml:"border_style"`
	LineNumber  string       `toml:"line_number"`
	Syntax      SyntaxConfig `toml:"syntax"`
}

type SyntaxConfig struct {
	Base     string `toml:"base"`
	Keyword  string `toml:"keyword"`
	String   string `toml:"string"`
	Number   string `toml:"number"`
	Comment  string `toml:"comment"`
	Type     string `toml:"type"`
	Function string `toml:"function"`
	Operator string `toml:"operator"`
	Variable string `toml:"variable"`
	Error    string `toml:"error"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = resolveConfigPath()
	}

	cfg := Default()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg, nil
	}

	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func resolveConfigPath() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "refr", "config.toml")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "refr", "config.toml")
}
