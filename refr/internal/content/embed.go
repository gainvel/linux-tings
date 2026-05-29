package content

import "embed"

// EmbeddedPages holds the embedded filesystem. It is set by the main
// package via SetEmbeddedPages since go:embed cannot reference parent directories.
var EmbeddedPages embed.FS

func SetEmbeddedPages(fs embed.FS) {
	EmbeddedPages = fs
}
