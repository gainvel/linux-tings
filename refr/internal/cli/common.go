package cli

import (
	"fmt"
	"io/fs"
	"os"

	"refr/internal/config"
	"refr/internal/content"
)

func buildTree(cfg *config.Config) (*content.Node, error) {
	embeddedFS, err := fs.Sub(content.EmbeddedPages, "pages")
	if err != nil {
		return nil, fmt.Errorf("embedded pages: %w", err)
	}

	tree, err := content.BuildTree(embeddedFS)
	if err != nil {
		return nil, fmt.Errorf("building tree: %w", err)
	}

	if info, statErr := os.Stat(cfg.PagesDir); statErr == nil && info.IsDir() {
		userTree, err := content.BuildTree(os.DirFS(cfg.PagesDir))
		if err == nil {
			tree = content.Overlay(tree, userTree)
		}
	}

	return tree, nil
}
