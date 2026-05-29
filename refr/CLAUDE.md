# refr

Personal desktop reference CLI for system commands and knowledge. Keyboard-launched hierarchical browser with syntax-highlighted content pages. **Early development — most rendering and tree-building functions are stubs.**

## Target System

- Artix Linux (rolling, Arch-based), plasma-openrc variant, x86_64
- OpenRC 0.63.1+, Go 1.26.2, KDE Plasma, bash
- Package managers: pacman, yay

## Build & Run

```sh
make build        # go build -o refr ./cmd/refr
make install      # copy to ~/.local/bin/
make test         # go test ./...
make lint         # go vet ./...
make clean        # remove built binary
go run ./cmd/refr # quick dev iteration without installing
```

Single test: `go test -run TestName ./internal/config/`

## Data Flow

```
main.go
  → content.SetEmbeddedPages(refr.Pages)   # must happen before Execute()
  → cli.Execute()
    → cobra parses flags/subcommands
    → config.Load(cfgFile)                  # --config flag > XDG > ~/.config/refr/config.toml
    → content.BuildTree(fs)                 # STUB: returns empty root node
    → content.Overlay(embedded, user)       # STUB: returns base unchanged
    → nav.New(tree)                         # functional: stack-based navigation
    → browse loop: render category → read input → select/back/quit
    → render.RenderPage() for leaf nodes    # STUB: returns title only
```

## Implementation Status

| Package | Status | Notes |
|---------|--------|-------|
| `cmd/refr/main.go` | Working | Entry point, wires embed to content |
| `internal/cli/` | Working | cobra commands: root, browse, search, version. RunE handlers are no-op |
| `internal/config/` | Working | TOML loading, full default config, XDG resolution |
| `internal/nav/` | Working | Stack-based navigator: Select, Back, Breadcrumb |
| `internal/content/page.go` | Types only | Page/Frontmatter structs, no parser |
| `internal/content/section.go` | Types only | Section/SectionType, no parser |
| `internal/content/tree.go` | **Stub** | BuildTree returns empty root node |
| `internal/content/overlay.go` | **Stub** | Overlay returns base unchanged |
| `internal/content/embed.go` | Working | Package-level var + setter for go:embed FS |
| `internal/render/render.go` | **Stub** | RenderCategory/RenderPage return name/title only |
| `internal/render/style.go` | Types only | Styles struct wraps ThemeConfig, no styling logic |
| `internal/render/highlight.go` | **Stub** | Highlight returns code unchanged |
| `internal/render/pager.go` | Working | Pipes to external pager or stdout |

## Project Structure

```
cmd/refr/          entry point
internal/
  cli/             cobra commands (root, browse, search, version)
  config/          TOML config loading, defaults, XDG resolution
  content/         .ref file parsing, tree building, embedded content, overlay
  render/          terminal output: styling, syntax highlighting, pager
  nav/             stack-based navigation state machine
pages/             embedded default .ref content (go:embed)
  _index.ref       root category metadata
```

## Content Format (.ref files)

YAML frontmatter for metadata. Body uses section directives:
- `@text` — plain prose (implicit for content before first directive)
- `@code <lang>` — syntax-highlighted code block
- Directives are parsing markers, never rendered

```
---
title: Example Page
order: 1
tags: [example]
---

Some introductory text.

@code bash
echo "hello"

@text
More explanation here.
```

Category directories use `_index.ref` for metadata (title, order, description). Without it, directory name is title-cased and sorted alphabetically.

## Content Overlay

Embedded pages via `go:embed` provide defaults. User pages at `pages_dir` (default `~/.config/refr/pages/`) overlay the embedded tree:
- Same relative path replaces the embedded version
- New paths extend the tree
- No manifest — directory structure is the tree

## Config

TOML at `~/.config/refr/config.toml`. Resolution: `--config` flag > `$XDG_CONFIG_HOME/refr/config.toml` > `~/.config/refr/config.toml`. **If the file doesn't exist, defaults are returned silently** (no error).

| Key | Purpose |
|-----|---------|
| `pages_dir` | User content overlay directory |
| `pager` | External pager command (e.g. `"less -R"`) |
| `show_numbers` | Show numbered selections in browse view |
| `[theme]` | UI colors: bg, text, category, page, header, accent, border, border_style, line_number |
| `[theme.syntax]` | Per-token syntax highlight colors with chroma base theme fallback |

Token override format: `"[bold] [italic] [underline] #hexcolor"` — modifiers optional.

## Conventions

- Config passed through constructors, no global state — **exception**: `content.EmbeddedPages` is a package-level `embed.FS` var because `go:embed` cannot reference parent directories cross-package. Set once via `SetEmbeddedPages()` in main before anything else runs.
- `fs.FS` interfaces for content loading (testable with `fstest.MapFS`)
- Fail-fast config loading: `config.Load` returns an error on malformed TOML rather than falling back to partial config. Missing file is not an error (returns defaults).
- Adding content = adding .ref files to `pages/` or user overlay dir, no code changes needed

## Gotchas

- **`SetEmbeddedPages()` ordering.** Must be called before `cli.Execute()` — content package reads from the package var during tree building.
- **lipgloss and chroma are planned but not yet in go.mod.** Only cobra, toml, mousetrap, and pflag are current dependencies. Add them when implementing render/style.go and render/highlight.go.
- **No tests exist yet.** `make test` passes (nothing to run) but doesn't validate anything.
- **`cli.rootCmd` uses a package-level var** (`var rootCmd`) — this is standard cobra pattern, not a violation of the no-global-state convention.
