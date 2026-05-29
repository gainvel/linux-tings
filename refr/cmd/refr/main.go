package main

import (
	"os"

	refr "refr"
	"refr/internal/cli"
	"refr/internal/content"
)

func main() {
	content.SetEmbeddedPages(refr.Pages)

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
