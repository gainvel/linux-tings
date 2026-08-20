# Roaming Cat — Implementation Plan

A pixel art cat that roams on top of the KDE Plasma taskbar panel. Built with Qt6/QML using `layer-shell-qt` for precise Wayland overlay positioning.

---

## Technology Stack

**Qt6/C++ + QML + layer-shell-qt**

- `layer-shell-qt` 6.6.5 installed — provides both C++ API (`LayerShellQt::Window`) and QML module (`org.kde.layershell`)
- `zwlr_layer_shell_v1` v5 confirmed on this compositor
- QML `AnimatedSprite` handles sprite sheet animation natively
- Same library stack KDE's own panel uses

**Why not alternatives:**
- *Plasma widget*: confined to widget slot, can't roam freely across panel width
- *Standalone Qt window*: Wayland prevents client-side window positioning
- *Python/PyQt*: no Python bindings for layer-shell-qt
- *Rust/smithay*: viable but more boilerplate for sprite animation vs QML built-ins

---

## Project Structure

```
roaming-cat/
  CMakeLists.txt
  src/
    main.cpp                 -- app entry, QML engine, Wayland input region setup
    catcontroller.h/.cpp     -- state machine logic (C++ QObject exposed to QML)
    paneldetector.h/.cpp     -- DBus panel geometry queries + config fallback
    config.h/.cpp            -- JSON config loading
  qml/
    Main.qml                 -- layer shell window (transparent, anchored above panel)
    CatSprite.qml            -- AnimatedSprite with state-driven source switching
    DebugOverlay.qml         -- optional bounding box / state label overlay
  assets/
    sprites/default/         -- default sprite pack
      manifest.json          -- declares dimensions, animations, FPS
      idle.png
      walk-right.png
      walk-left.png
      turn-r2l.png
      turn-l2r.png
      sit-down.png
      sitting-idle.png
      lay-down.png
      sleeping.png
      wake-up.png
      stretch.png
      groom.png
      alert.png
  scripts/
    install.sh               -- build + install + autostart setup
```

---

## Layer Shell Window

The cat lives on a transparent Wayland layer surface positioned directly above the panel.

```
Layer:          LayerOverlay (above panel and normal windows)
Anchors:        Bottom | Left | Right (spans full screen width)
Bottom margin:  panelHeight (48px) — sits ON TOP of panel
Exclusion zone: 0 (doesn't push other windows)
Keyboard:       None (no keyboard grab)
Input region:   Cat bounding box only (right-click menu; rest is click-through)
Window height:  Sprite height (default 24px)
Scope:          "roaming-cat"
Screen:         DP-2 (bottom monitor, geometry 0,1080)
```

**Click-through**: The `wl_surface` input region is set to the cat's bounding box only (updated each frame as the cat moves). Clicks outside the cat pass through to the panel and desktop. Right-clicking the cat opens a context menu.

**Panel auto-hide**: When the panel hides, adjust `margins.bottom` to 0. When it reappears, restore to `panelHeight`. Monitor via DBus signals or `QFileSystemWatcher` on `plasmashellrc`.

---

## Animation State Machine

```
IDLE (standing still, subtle blink/ear twitch)
  -> WALKING_RIGHT/LEFT    random 3-15s timer
  -> GROOMING              random, low probability
  -> ALERT                 random, brief 1-3s
  -> SITTING               after 10-20s idle

WALKING (moving across panel at configurable speed)
  -> IDLE                  random stop or panel edge reached
  -> TURNING               panel edge reached or random
  -> SITTING               random chance per tick

TURNING (4-frame directional change)
  -> WALKING (opposite)
  -> IDLE

SITTING (stationary, tail swish)
  -> IDLE                  stands back up
  -> LAYING_DOWN           10-30s timer
  -> STRETCHING            random
  -> GROOMING              random

LAYING_DOWN
  -> SLEEPING              10-20s, or system idle detected
  -> WAKING_UP             random or system active

SLEEPING (slow breathing loop)
  -> WAKING_UP             system activity or 60-300s timer

WAKING_UP -> STRETCHING (always)
STRETCHING -> IDLE (always)
GROOMING -> IDLE (after animation completes)
ALERT -> IDLE (after 1-3s)
```

**System-driven triggers:**
- `ext_idle_notifier_v1` idle timeout → bias toward SLEEPING
- `ext_idle_notifier_v1` resume → WAKING_UP if sleeping
- Panel auto-hide → adjust bottom margin

---

## Sprite Pack System

Each sprite pack is a directory with a `manifest.json` — the engine reads dimensions and animation metadata from this file rather than hard-coding anything.

```json
{
  "name": "default",
  "width": 24,
  "height": 24,
  "animations": {
    "idle":         { "frames": 4,  "fps": 3,   "loop": true  },
    "walk-right":   { "frames": 8,  "fps": 10,  "loop": true  },
    "walk-left":    { "frames": 8,  "fps": 10,  "loop": true  },
    "turn-r2l":     { "frames": 4,  "fps": 8,   "loop": false },
    "turn-l2r":     { "frames": 4,  "fps": 8,   "loop": false },
    "sit-down":     { "frames": 4,  "fps": 6,   "loop": false },
    "sitting-idle": { "frames": 4,  "fps": 2,   "loop": true  },
    "lay-down":     { "frames": 4,  "fps": 4,   "loop": false },
    "sleeping":     { "frames": 4,  "fps": 1.5, "loop": true  },
    "wake-up":      { "frames": 4,  "fps": 6,   "loop": false },
    "stretch":      { "frames": 6,  "fps": 8,   "loop": false },
    "groom":        { "frames": 6,  "fps": 6,   "loop": false },
    "alert":        { "frames": 3,  "fps": 6,   "loop": false }
  }
}
```

Each animation is a horizontal sprite strip PNG: `(width * frames) x height`. File names must match the animation keys in the manifest.

Sprite size is configurable — users can create packs at any resolution (32x32, 48x48, non-square). The engine reads `width`/`height` from the manifest and the config can override with `sprite.width`/`sprite.height`.

**Total for default pack: ~63 unique frames across 13 animations.**

---

## Configuration

`~/.config/roaming-cat/config.json` (XDG via `QStandardPaths`):

```json
{
  "sprite": {
    "width": null,
    "height": null,
    "scale": 1.0,
    "theme": "default"
  },
  "behavior": {
    "speed": 2.0,
    "idle_timeout_min": 3,
    "idle_timeout_max": 15,
    "sleep_after_system_idle_sec": 300,
    "edge_behavior": "turn"
  },
  "appearance": {
    "offset_above_panel": 0,
    "monitor": "auto"
  },
  "debug": {
    "show_bounds": false,
    "show_state": false,
    "log_transitions": false
  }
}
```

`sprite.width`/`height` of `null` means "use manifest defaults."

---

## Right-Click Menu

- **Settings** — opens config file in default editor
- **Quit** — graceful shutdown

---

## Panel Detection

1. **Primary**: DBus query at startup via `org.kde.plasmashell` → panel height, location, screen
2. **Fallback**: Parse `~/.config/plasmashellrc` (thickness) + `~/.config/plasma-org.kde.plasma.desktop-appletsrc` (location=4 → bottom, screen)
3. **Reactive**: `QFileSystemWatcher` on plasmashellrc for runtime changes

Current panel: 48px, bottom, floating, auto-hide, opaque, DP-2 (screen 0), 1920px wide.

---

## Autostart

XDG autostart entry at `~/.config/autostart/roaming-cat.desktop`:

```ini
[Desktop Entry]
Type=Application
Name=Roaming Cat
Exec=/path/to/roaming-cat
X-KDE-autostart-phase=2
```

NOT a systemd service. This is an OpenRC system.

---

## Build Dependencies

All already installed:
- `qt6-base`, `qt6-declarative`, `qt6-wayland`
- `layer-shell-qt` (CMake: `LayerShellQt`)
- `extra-cmake-modules`
- `cmake`, `gcc`

---

## Implementation Phases

1. **Skeleton** — layer shell window, transparent, positioned above panel. Single static sprite. Confirm click-through works.
2. **Animation engine** — `AnimatedSprite` with walk-right. Cat moves left/right, bounces at edges.
3. **State machine** — full state machine with all transitions and random timers.
4. **System integration** — idle detection, panel hide/show tracking, multi-monitor awareness.
5. **Polish** — config system, sprite theme swapping, right-click context menu, debug overlay.
