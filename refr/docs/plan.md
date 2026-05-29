# refr — Implementation Plan

## Overview

`refr` is a keyboard-launched CLI that presents a hierarchical browser for system command references. Users navigate numbered categories, drill into subcategories, and view content pages with mixed plain text and syntax-highlighted code blocks.

## Architecture

```
User runs `refr`
  → cobra parses flags/subcommands
  → config.Load() reads TOML config
  → content.BuildTree() constructs node tree from embedded + user .ref files
  → nav.NewNavigator(tree) initializes browsing state
  → browse loop: render category → read input → select/back/quit
  → render.RenderPage() for leaf nodes (lipgloss text + chroma code)
```

## Content Format (.ref)

YAML frontmatter + section-directive body.

### Directives

| Directive | Purpose |
|---|---|
| `@text` | Plain prose section, styled with lipgloss |
| `@code <lang>` | Syntax-highlighted code via chroma |

- Directives are parsing markers, invisible when rendered
- Content before the first directive is implicitly `@text`
- A new directive ends the previous section

### Page Example

```
---
title: OpenRC Service Management
order: 2
tags: [openrc, services]
---

OpenRC provides service management through the rc-service
and rc-update commands.

@code bash
# Start / stop / restart
sudo rc-service sshd start
sudo rc-service sshd stop
sudo rc-service sshd restart

# Check status
rc-service sshd status

@text
Manage which services start at boot:

@code bash
# Enable at default runlevel
sudo rc-update add sshd default

# Disable
sudo rc-update del sshd default

# List all services by runlevel
rc-update show
```

### Category Metadata (_index.ref)

```
---
title: OpenRC
order: 1
description: Init system commands and configuration
---
```

Optional. Without it, directory name is title-cased and sorted alphabetically.

## Config (`~/.config/refr/config.toml`)

Resolution order: `--config` flag → `$XDG_CONFIG_HOME/refr/config.toml` → `~/.config/refr/config.toml`

```toml
pages_dir    = "~/.config/refr/pages"
pager        = "less -R"
show_numbers = true

[theme]
bg           = "#181818"
text         = "#d8dcc9"
category     = "#5d8a78"
page         = "#8fb37a"
header       = "#6f9a6a"
accent       = "#b5c97a"
border       = "#3a423d"
border_style = "rounded"
line_number  = "#8a9085"

[theme.syntax]
base     = "monokai"
keyword  = "#6f9a6a"
string   = "#8fb37a"
number   = "#b5c97a"
comment  = "italic #6a7866"
type     = "#5d8a78"
function = "#a3b89a"
operator = "#9ba87a"
variable = "#a3b89a"
error    = "bold #c47a4a"
```

Token override format: `"[bold] [italic] [underline] #hexcolor"` — modifiers optional.

## Navigation Flow

1. Launch `refr` (or `refr browse`)
2. Header: breadcrumb showing current location
3. Numbered list of children at current level:
   ```
   refr > OpenRC

   1  Basics                    Getting started with OpenRC
   2  Service Management        Start, stop, enable services
   3  Runlevels                 Default, boot, shutdown runlevels

   [number] select  ·  [b] back  ·  [q] quit
   ```
4. Type number → drill into child (category or page)
5. Pages render inline with styled text and highlighted code
6. `b` goes back, `q` quits

### Search (`refr search <query>`)

Full-text search across page titles, tags, and body content. Results displayed as numbered list; selecting a result renders the page directly.

## Dependencies

| Library | Purpose |
|---|---|
| `github.com/spf13/cobra` | CLI framework, --help, subcommands |
| `github.com/BurntSushi/toml` | Config file parsing |
| `github.com/charmbracelet/lipgloss` | Terminal styling (colors, borders) |
| `github.com/alecthomas/chroma/v2` | Syntax highlighting for @code sections |

## Content Overlay

Embedded pages (via `go:embed`) provide defaults that work out of the box. User pages at `~/.config/refr/pages/` overlay the embedded tree:
- Same relative path → user version replaces embedded
- New paths → extend the tree
- No manifest needed — directory structure is the tree

## Project Structure

```
cmd/refr/main.go                  entry point
internal/
  cli/
    root.go                       cobra root command, global flags
    browse.go                     browse subcommand (default)
    search.go                     search subcommand
    version.go                    version subcommand
  config/
    config.go                     Config/ThemeConfig/SyntaxConfig structs, Load()
    defaults.go                   default config values
  content/
    page.go                       Page struct, frontmatter parsing
    section.go                    Section struct, @text/@code directive parser
    tree.go                       BuildTree() from fs.FS
    node.go                       Node struct (category vs leaf)
    embed.go                      go:embed directive for pages/
    overlay.go                    merge embedded + user trees
  render/
    render.go                     RenderCategory(), RenderPage()
    style.go                      lipgloss styles from ThemeConfig
    highlight.go                  chroma highlighting from SyntaxConfig
    pager.go                      pipe to $PAGER
  nav/
    navigator.go                  Navigator: stack, Select(), Back(), Breadcrumb()
pages/                            embedded .ref content
```

## Implementation Phases

### Phase 1 — Skeleton
1. Go module, cobra root command, main.go
2. Config loading with TOML + defaults
3. Embed a test page, parse .ref format (frontmatter + sections)
4. Build content tree from embedded FS
5. Navigator with stack-based browsing
6. Basic terminal output (no styling)
7. Wire browse loop end-to-end

### Phase 2 — Polish
1. lipgloss styling from theme config
2. chroma syntax highlighting with token-level overrides
3. Pager support for long content
4. User page directory overlay
5. Search subcommand
6. Shell completion via cobra

### Phase 3 — Content
1. Write reference pages for openrc, pacman, yay, networking
2. Decoupled from code — just adding .ref files
