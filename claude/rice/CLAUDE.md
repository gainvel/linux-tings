# Desktop Ricing Hub

> **WARNING: This is Artix Linux with OpenRC — NOT systemd.**
> `systemctl`, `journalctl`, `systemd-*`, `.service` units, `~/.config/environment.d/`, systemd timers — **none of these exist here.**
> Use `rc-service`, `rc-update`, `/etc/init.d/` scripts. Logs are plain files in `/var/log/` (syslog via `metalog`).
> For session env vars: `~/.config/plasma-workspace/env/*.sh`. For scheduled tasks: `cronie` (`crontab -e`).
> For autostart: `~/.config/autostart/*.desktop` (XDG autostart). For hostname/locale/time: edit files directly.
> See `/home/void/claude/general/CLAUDE.md` for the full OpenRC reference and what replaces each systemd component.

This directory is the central hub for all desktop ricing projects. Each project lives in its own subdirectory. For exhaustive system details (services, GPU stack, audio, gaming, package management, Forest theme palette), see `/home/void/claude/general/CLAUDE.md`.

---

## System Overview

| Component | Value |
|-----------|-------|
| OS | Artix Linux (OpenRC), rolling release |
| Kernel | `linux` (vanilla), check `uname -r` for version |
| Desktop | KDE Plasma 6.6.5, KWin 6.6.5 |
| Session | **Wayland** (verify: `echo $XDG_SESSION_TYPE`) |
| GPU | AMD RX 9070 XT (RDNA 4, Navi 48) + Raphael iGPU, both `amdgpu` + mesa |
| CPU | AMD Ryzen 7 7800X3D, 64GB RAM |
| Display | Dual 1080p stacked vertically, both 144Hz |
| Panel | 48px, bottom, floating, auto-hide, opaque, on DP-2 |
| Package mgr | pacman + yay (AUR) |
| Shell | bash |
| Keyboard | 75% layout (no numpad) |

## Display Layout

```
+------------------+
|    DP-3 (top)    |  1920x1080 @ 144Hz, geometry: 0,0
+------------------+
|   DP-2 (bottom)  |  1920x1080 @ 144Hz, geometry: 0,1080
+------------------+
         ^-- Panel lives here (bottom edge of DP-2)
```

---

## KDE Customization Paths

### User config (`~/.config/`)
- `kdeglobals` — colors, fonts, widget style
- `kwinrc` — window manager, compositing, effects, tiling
- `plasmarc` — desktop theme selection
- `plasmashellrc` — panel thickness, visibility, opacity
- `plasma-org.kde.plasma.desktop-appletsrc` — panel applets, desktop containments
- `kglobalshortcutsrc` — global keyboard shortcuts
- `kwinrulesrc` — per-window rules
- `kwinoutputconfig.json` — display output config

### User themes (`~/.local/share/`)
- `plasma/desktoptheme/` — desktop themes (Glassy, WhiteSur-dark, sumac-night)
- `plasma/look-and-feel/` — global look-and-feel packages
- `color-schemes/` — color scheme `.colors` files
- `aurorae/themes/` — window decoration themes (WhiteSur variants)
- `icons/` — icon themes (WhiteSur, WhiteSur-dark, WhiteSur-light)
- `wallpapers/` — wallpaper packages

### System resources (`/usr/share/`)
- `plasma/desktoptheme/` — breeze-dark, breeze-light, default, oxygen
- `plasma/look-and-feel/` — breeze, breeze-dark, breezetwilight, oxygen
- `color-schemes/` — system color schemes
- `icons/` — breeze, breeze-dark, oxygen, adwaita
- `themes/` — Breeze, Breeze-Dark, Default

### Cross-toolkit styling
- `~/.config/gtk-3.0/settings.ini` — GTK3 theme, icons, cursor, font
- `~/.config/gtk-4.0/settings.ini` — GTK4
- `~/.gtkrc-2.0` — legacy GTK2
- `~/.config/qt5ct/qt5ct.conf` — Qt5 style (Fusion), icons, color scheme

## Current Theme Stack

- **Desktop theme**: default Breeze (Glassy available)
- **Look-and-feel**: WhiteSur-dark
- **Color scheme**: BreezeDark (WhiteSur, Glassy, Alpha 20 variants available)
- **Window decorations**: Aurorae WhiteSur variants (12 themes)
- **Icons**: breeze-dark (GTK), WhiteSur variants available
- **Cursor**: breeze_cursors 24px
- **GTK**: Artix-dark
- **Qt5**: Fusion
- **Font**: Noto Sans 10pt

---

## Wayland Protocols

These are confirmed available on this compositor — relevant for any overlay/widget work:

- `zwlr_layer_shell_v1` v5 — layer surfaces (panels, overlays, backgrounds)
- `ext_idle_notifier_v1` v2 — system idle detection
- `wp_viewporter` v1 — surface scaling
- `wp_fractional_scale_manager_v1` v1 — fractional scaling
- `zwp_idle_inhibit_manager_v1` v1 — prevent idle

**layer-shell-qt** is installed:
- C++ API: `#include <LayerShellQt/Window>` (CMake: `find_package(LayerShellQt)`)
- QML module: `import org.kde.layershell` (at `/usr/lib/qt6/qml/org/kde/layershell/`)

---

## Development Tools

- **C/C++**: GCC, CMake, qmake
- **Qt**: Qt6 6.11.1, Qt5 5.15.18 — prefer Qt6 for all new work
- **KDE Frameworks**: KF6 full set (KConfig, KWindowSystem, KNotifications, etc.)
- **Rust**: rustc/cargo 1.95.0
- **Python**: 3.14.5 (no PyQt6/PySide6 — install via pip if needed)
- **Node**: v22.22.2
- **Art**: GIMP 3.2.4, Krita 6.0.1, LibreSprite (AUR: `libresprite-git`)

---

## Querying Plasma at Runtime

```bash
# Panel info via DBus
qdbus6 org.kde.plasmashell /PlasmaShell org.kde.PlasmaShell.evaluateScript \
  "print(JSON.stringify(panels()))"

# Display outputs
kscreen-doctor -o

# Write a KDE config key
kwriteconfig6 --file <name> --group <group> --key <key> <value>

# Reload Plasma shell (without logout)
kquitapp6 plasmashell && kstart plasmashell

# Apply color scheme
plasma-apply-colorscheme <name>

# Apply look-and-feel
plasma-apply-lookandfeel -a <id>

# Wayland protocol support
wayland-info | grep "interface:"
```

---

## Projects

- `roaming-cat/` — pixel art cat that roams on top of the taskbar (Qt6/QML + layer-shell-qt)
- *(planned)* Breeze fork — fine-tuned Breeze variant with custom styling

---

## Conventions

- All new GUI work targets **Qt6/KF6 + Wayland**. No X11 assumptions.
- Use `layer-shell-qt` for overlay or panel-adjacent surfaces.
- Test on both monitors; the panel is on the bottom screen (DP-2, geometry 0,1080).
- For services/daemons: write OpenRC init scripts or use XDG autostart — **never** systemd units.
- For config changes: `kwriteconfig6` is safer than direct file edits while Plasma is running (Plasma may overwrite files on logout).
- Run `pacman -Rcs --print <pkg>` before any package removal (no snapshots, no undo).
- The **Forest / Nature theme** palette is defined in `/home/void/claude/general/CLAUDE.md` — reference it for any color work.
