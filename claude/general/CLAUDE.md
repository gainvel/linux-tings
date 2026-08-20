# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

This directory (`/home/void/claude/general/`) is a general workspace for system tasks on a riced **Artix Linux gaming desktop**. There is no codebase here — this CLAUDE.md is the system manual. Use it before touching configs, installing packages, or modifying services.

---

## System Identity

- **Distro**: Artix Linux rolling, OpenRC variant (plasma-openrc build)
- **Kernel**: `linux` (vanilla — not zen / cachy / lts / tkg). Check `uname -r` for current version.
- **Init**: **OpenRC** — *not systemd*. `systemctl` does nothing here. Use `rc-service` and `rc-update`.
- **Bootloader**: GRUB (`/boot/grub/`)
- **Filesystem**: ext4. **No snapshots** (no btrfs, no snapper, no timeshift). Destructive ops are unrecoverable without external backup.
- **Swap**: **None.** OOM-killer can fire under memory pressure. If swap is wanted, prefer zram over a swapfile.
- **GPU**: AMD RX 9070 XT (RDNA 4, Navi 48) via `amdgpu` kernel driver + mesa/vulkan-radeon userspace. AMD Raphael iGPU also present (both use `amdgpu`).
- **Audio**: PipeWire (with `pipewire-pulse`, `pipewire-alsa`, `wireplumber`).
- **Display server**: Wayland session active (`echo $XDG_SESSION_TYPE` to confirm). Plasma X11 session also installed as fallback.
- **DE**: KDE Plasma 6, login via SDDM.
- **Shell**: bash. Minimal `~/.bashrc` (only adds `~/.local/bin` to PATH).
- **Terminal**: Alacritty (primary), Konsole (secondary).
- **Network**: NetworkManager (dhcpcd present as fallback).
- **Firewall**: None. **No nftables/ufw/firewalld rules loaded.**
- **CPU governor**: `performance` (fixed — see `/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor`).
- **Package count**: ~1200 total, ~220 explicit. No flatpak, snap, or nix — pure yay/AUR or pacman.
- **Keyboard**: **75% layout** (function row present, no numpad, compact arrow cluster, tight nav column). When designing keybindings: do **not** rely on numpad keys. Everything else is fair game — function row, arrows, Home/End/PgUp/PgDn (standalone or with modifiers) are all fine.

---

## Reliability Rules (read before any change)

1. **Confirm before destructive ops.** No snapshots and no swap means mistakes don't get undone. `rm -rf`, `pacman -Rdd`, `pacman -Rss` on anything touching kernel / display / pipewire / NetworkManager / sddm requires explicit user confirmation.
2. **OpenRC, not systemd — this affects far more than just services.** systemd is completely absent from this system. Do not use or assume the existence of *any* systemd component. This includes but is not limited to:
   - **Services**: No `systemctl`. Use `rc-service` / `rc-update` (see below).
   - **Logging**: No `journalctl`. Logs are plain files in `/var/log/` (see below).
   - **User environment**: No `~/.config/environment.d/` (that's `systemd --user` environment generator). For session-wide env vars, use `~/.config/plasma-workspace/env/*.sh` (sourced by Plasma at login) or `~/.bash_profile` (login shells only).
   - **Timers**: No `systemd-timer`. Use `cronie` (installed and enabled) — `crontab -e` for user cron.
   - **Hostname / time / locale**: No `hostnamectl`, `timedatectl`, `localectl`. Use `/etc/hostname`, `/etc/conf.d/hwclock` + `hwclock`, `/etc/locale.conf` directly.
   - **DNS**: No `systemd-resolved`. DNS goes through NetworkManager (`/etc/resolv.conf` managed by NM).
   - **Tmpfiles**: No `systemd-tmpfiles`. Temp cleanup is via `tmp.d` or manual.
   - **Login sessions**: No `systemd-logind`. Session tracking is via `elogind` (same D-Bus API, separate package). `loginctl` works because elogind provides it.
   - **Boot analysis**: No `systemd-analyze`. Boot profiling isn't available out-of-the-box.
   - **Scope/slice/cgroup management**: No `systemd-run`, `systemd-cgls`, `systemd-cgtop`. Cgroup management is manual or via other tools.
   - **Network**: No `systemd-networkd`. Networking is via NetworkManager.
   - **Mount units**: No `.mount` / `.automount` units. Use `/etc/fstab` and `mount`/`umount`.

   **General rule**: if a tool, path, config format, or mechanism has "systemd" in its name or was introduced as part of the systemd ecosystem, assume it does not exist here and find the traditional/OpenRC/standalone alternative.
3. **GPU is AMD RX 9070 XT.** Driver is in-kernel `amdgpu` + mesa/vulkan-radeon. No NVIDIA packages installed (firmware-nvidia is kept as a hard dep of linux-firmware).
4. **Repos are `[system] [world] [galaxy] [lib32]`.** Don't enable `[universe]`, `chaotic-aur`, or other third-party repos without asking. AUR is available via `yay`.
5. **Plasma rewrites its configs.** Edits to `~/.config/kdeglobals`, `plasmashellrc`, `plasma-org.kde.plasma.desktop-appletsrc` may be clobbered on logout. Prefer Plasma's own GUI tools (`kwriteconfig6`, System Settings) when possible, or edit while the session is logged out.
6. **Verify before pkg removals.** Run `pacman -Rcs --print <pkg>` first to see what will be pulled. Watch for `linux-firmware`, `pipewire`, `wireplumber`, `xorg-*`, `wayland`, `mesa`, `vulkan-radeon`, `qt6-*`, `kf6-*`, `plasma-*` in the cascade.
7. **No firewall; backups are tiered per-folder** (see Storage Layout — `~/.config/storage-backup.list`). Only folders listed there are protected; flag before exposing services to the network and assume no recovery path for any unlisted deleted user data.
8. **Snapshot the explicit-package list before bulk changes**: `pacman -Qqe > /tmp/pkgs-$(date +%F).txt`.

---

## Storage Layout

~10 TB total across three fixed drives:

| Device | Size | Mount | Role |
|---|---|---|---|
| `nvme0n1` (Samsung 990 PRO) | 2 TB | `/` (+ `/boot/efi`) | OS, games, active work (`~/Software`, `~/Repo`, `~/Inbox`) |
| `sda1` "Storage1" | 4 TB | `/mnt/storage1` | **primary** media + Documents |
| `sdb1` "Storage2" | 4 TB | `/mnt/storage2` | Archive + Backups |

**Bulk data lives on the HDDs via `$HOME` symlinks — keep them intact** (apps & XDG write to them by name; deleting/renaming breaks default save locations):
- `~/Music ~/Pictures ~/Videos ~/Documents` → `/mnt/storage1` (`Media/*` + `Documents`, sub-categorized)
- `~/Archive ~/Backups` → `/mnt/storage2` (Archive: `3D`/`Art`/`Design`/`Software`/`Projects-Archive`/`zips`; Backups: tiered mirrors + `Manual` drop)

**Software hub**: `~/Software` (NVMe) is the code hub — **created** projects live in `~/Software/projects/` (was `~/projects`), `~/Repo` holds **downloaded/cloned** repos. Both in the Dolphin Places sidebar. `~/Software/README.md` is a master map of **all** customizable/inspectable software on the machine (apps, daemons, configs, services). Software is backed up at tier-3, Repo at tier-2 (see Backup below).

**Auto-filing** — *"storage sort"/"sorting" tasks mean this system:* `~/.local/bin/storage-sorter` (Go daemon, autostarts at login, watches `~/Inbox`) files dropped files by extension onto the right drive; unknown types → `~/Inbox/_unsorted`. **Routing lives in `~/.config/storage-sorter/routes.json`** (JSON: categories → dirs/exts/rules) and **hot-reloads on write — no restart**. Edit it with the **`storagesort` TUI** (Rust/ratatui config editor, `~/.local/bin/storagesort`, needs a TTY) or by hand; only **Go-code** changes need rebuild+restart (`cd ~/.local/src/storage-sorter && go build -o ~/.local/bin/storage-sorter . && pkill -x storage-sorter; setsid -f ~/.local/bin/storage-sorter watch`). `dir` paths are `$HOME`-relative, never absolute. Dolphin right-click → "Send to category" calls the same daemon. Log: `~/logs/inbox-sort.log`.

**Backup (tiered)**: `~/.local/bin/storage-backup.sh` runs nightly via cron (`0 3 * * *`). Each line of `~/.config/storage-backup.list` is `<path> <tier>`, where tier = the number of **physical drives** holding a copy (incl. the original): `1` = none, `2` = original + one mirror, `3` = all three drives. The script `realpath`s the source to find its origin drive, then mirrors (`rsync -aH --delete`) to the other drives (HDDs preferred) under `<drive>/Backups/<name>/`. It's a `--delete` mirror (not history); only listed paths are protected. To protect a new folder, add a `<path> <tier>` line. Log: `~/logs/storage-backup.log`.

**Full maintenance reference**: `~/.local/src/storage-sorter/README.md`, plus the daemon/TUI contract (routes.json schema + hard rules) in `~/Software/projects/storage-sorter-gui/CLAUDE.md`. Read before changing routing, categories, symlinks, or the Places sidebar (`~/.local/share/user-places.xbel`).

**Placement**: games → SSD; bulky non-game data → HDDs. Watch root headroom (OS + games share the 2 TB NVMe). **Only folders enrolled in the tiered backup list above are protected** — treat everything else on `/mnt/storage*` as primary-only. (Pending: holding folder `~/.storage-migrated-2026-06-18` awaits manual `rm -rf`; the `/mnt/storage2` retired folder is already cleared.)

---
## USB Formatting
1. - **Use ExFat unless directed otherwise** User runs between MacOS and Linux, compatability is needed.
2. - **Use USB naming scheme** User may request you to use a default naming format, on request name the USB "USB_[total_size]GB_[1-99]". The ladder number should be based on any USB drives with the same label that are found when 'lsblk -o NAME,LABEL,MODEL,SIZE' is ran..

---

## Running Commands as Root

Use `SUDO_ASKPASS=/usr/bin/ksshaskpass sudo -A ...` for privileged commands. This pops a KDE password dialog. **Minimize sudo invocations** — each one prompts for a password. Batch multiple root operations into a single `sudo -A sh -c "cmd1 && cmd2 && ..."` call whenever possible.

---

## Package Management (pacman)

```bash
pacman -Syu                    # full system update (run as root)
pacman -Ss <name>              # search repos
pacman -Si <pkg>               # show repo pkg info
pacman -Qs <name>              # search installed
pacman -Qi <pkg>               # show installed pkg info
pacman -Ql <pkg>               # list files owned by pkg
pacman -Qo /path/to/file       # which pkg owns this file
pacman -Qqe                    # explicitly installed (for snapshotting)
pacman -Qdt                    # orphans (unneeded deps)
pacman -Qkk                    # verify file integrity of all installed pkgs
pacman -S <pkg>                # install
pacman -Rcs <pkg>              # remove + unneeded deps + configs (preferred)
pacman -Rcs --print <pkg>      # dry-run (always do this first on system pkgs)
```

**Repos enabled** (see `/etc/pacman.conf`): `[system]`, `[world]`, `[galaxy]`.
**Mirror config**: `/etc/pacman.d/mirrorlist` (Artix mirrors, not Arch mirrors).
**Pacman tweaks**: `ParallelDownloads = 7`, `Color`, `VerbosePkgLists`, `ILoveCandy`, `SigLevel = Required DatabaseOptional`.

**AUR access**: `yay` is installed (built from source). Use `yay -S <pkg>` to install AUR packages, `yay -Ss <pkg>` to search AUR + repos. Prefer official repo packages over AUR equivalents when both exist.

**Arch packages on Artix**: Artix mirrors most of `[extra]` into `[world]` and `[community]` into `[galaxy]`. Most `pacman -S <archname>` commands work as-is. The `[lib32]` repo is the Artix equivalent of Arch's `[multilib]`.

---

## Service Management (OpenRC)

```bash
rc-service <name> status        # check service
rc-service <name> start|stop|restart
rc-update show                  # list runlevel-enabled services
rc-update add <name> default    # enable at boot
rc-update del <name> default    # disable at boot
rc-status                       # status of current runlevel
```

**Logs**: `/var/log/` (no `journalctl`). Syslog daemon is `metalog` (logs to `/var/log/everything/`). Notable: `/var/log/sddm.log`, `/var/log/pacman.log`.

**Currently enabled at boot** (sample): `NetworkManager`, `sddm`, `bluetoothd`, `cupsd`, `cronie`, `acpid`, `dbus`, `elogind`, `cpupower`, `metalog`, `agetty.tty1-6`. PipeWire runs per-user (not via OpenRC). Run `rc-update show default` for the full list.

**Service files live in**: `/etc/init.d/`. Don't hand-write these — use the upstream-provided `*-openrc` companion package (e.g., `networkmanager-openrc`, `sddm-openrc`).

---

## Display, Desktop, Login

- **Login manager**: SDDM (`sddm-openrc`). Theme/config: `/etc/sddm.conf`, `/etc/sddm.conf.d/*`.
- **Session**: Wayland (verify with `echo $XDG_SESSION_TYPE`). X11 fallback available from SDDM session picker.
- **Plasma config (user)**:
  - `~/.config/kdeglobals` — global colors, fonts, recent docs
  - `~/.config/plasmashellrc` — Plasma shell
  - `~/.config/plasma-org.kde.plasma.desktop-appletsrc` — desktop layout, panels, widgets
  - `~/.config/kwinrc` — KWin (window manager) rules, virtual desktops, animations
  - `~/.config/kglobalshortcutsrc` — global shortcuts
- **Konsole config**: `~/.config/konsolerc`, profiles in `~/.local/share/konsole/*.profile`, color schemes in `~/.local/share/konsole/*.colorscheme`.
- **Apply config changes scriptably**: `kwriteconfig6 --file <name> --group <g> --key <k> <value>`, then `kquitapp6 plasmashell && kstart plasmashell` to reload.

---

## GPU & Audio

**GPU stack**:
- AMD RX 9070 XT (RDNA 4, Navi 48) — kernel driver `amdgpu`, userspace `mesa` + `vulkan-radeon`.
- AMD Raphael iGPU also served by `amdgpu` + mesa + `vulkan-radeon`.
- Vulkan loader: `vulkan-icd-loader` + ICDs (`vulkan-radeon`, `vulkan-intel`, `vulkan-swrast`).
- Verify: `vulkaninfo --summary | head -20`, `glxinfo | grep "OpenGL renderer"`.
- No NVIDIA packages installed (only `linux-firmware-nvidia` remains as a hard dep of `linux-firmware`).

**Audio (PipeWire)**:
- `pactl info` — server + sink info (compatibility shim, real backend is PW)
- `pw-cli list-objects` — full PW graph
- `wpctl status` — sinks, sources, default devices
- `wpctl set-default <id>` — change default sink
- Configs: `/etc/pipewire/`, `/etc/wireplumber/`, user overrides in `~/.config/pipewire/` and `~/.config/wireplumber/`.

---

## Gaming Setup

**Current state**: Steam installed with 32-bit graphics (`lib32-mesa`, `lib32-vulkan-radeon`, `lib32-vulkan-icd-loader`). `[lib32]` repo is enabled.

**Common additions** (install as needed): `lutris`, `wine`/`wine-staging`, `gamemode`/`lib32-gamemode`, `gamescope`, `mangohud`/`lib32-mangohud`, `protonup-qt`.

**Proton-GE**: drop tarballs into `~/.steam/root/compatibilitytools.d/`.

---

## Performance & Tuning

- **CPU governor**: currently `performance` (no auto-throttling). Inspect: `cpupower frequency-info`. Change: `cpupower frequency-set -g <governor>`.
- **No swap**. Add zram if memory pressure becomes a problem: `pacman -S zramen` (or write a small OpenRC service). Avoid swapfiles unless requested.
- **sysctl tweaks live in** `/etc/sysctl.d/`. Currently only `99-magic_sysrq-local.conf` (`kernel.sysrq=1`).
- **ulimit / file descriptors**: `/etc/security/limits.d/` is empty. Steam/Wine often want `nofile 1048576` — add a `99-gaming.conf` if needed.
- **`power-profiles-daemon`** and **`cpupower`** services are both running — pick one as the source of truth before editing governor settings.

---

## Key Config Paths

| Area | Path |
|---|---|
| Shell (bash) | `~/.bashrc`, `~/.bash_profile` |
| Terminal (Alacritty) | `~/.config/alacritty/alacritty.toml` |
| Terminal (Konsole) | `~/.config/konsolerc`, `~/.local/share/konsole/` |
| KDE / Plasma | `~/.config/kdeglobals`, `plasmashellrc`, `plasma-org.kde.plasma.desktop-appletsrc`, `kwinrc` |
| GTK 3 | `~/.config/gtk-3.0/{settings.ini,gtk.css,colors.css}` |
| GTK 4 / libadwaita | `~/.config/gtk-4.0/{settings.ini,gtk.css,colors.css,libadwaita.css}` |
| Qt 5 / Qt 6 | `~/.config/qt5ct/`, `~/.config/qt6ct/` |
| Icons / cursors | `~/.icons/`, `/usr/share/icons/` |
| GTK themes | `~/.themes/`, `/usr/share/themes/` |
| Plasma color schemes | `~/.local/share/color-schemes/` |
| Plasma global themes | `~/.local/share/plasma/look-and-feel/` |
| Konsole color schemes | `~/.local/share/konsole/*.colorscheme` |
| Custom scripts | `~/.local/bin/` (`claude` symlink, `storage-sorter` daemon, `storagesort` routing-TUI, `storage-backup.sh`); routing config `~/.config/storage-sorter/routes.json`; sorter source `~/.local/src/storage-sorter/`, TUI source `~/Software/projects/storage-sorter-gui/` |
| Fonts | `~/.local/share/fonts/`, `/usr/share/fonts/` (run `fc-cache -fv` after adds) |
| Pacman | `/etc/pacman.conf`, `/etc/pacman.d/mirrorlist` |
| OpenRC | `/etc/init.d/`, `/etc/conf.d/`, `/etc/rc.conf` |
| Sysctl | `/etc/sysctl.d/` |
| Limits | `/etc/security/limits.d/` |
| SDDM | `/etc/sddm.conf`, `/etc/sddm.conf.d/` |
| Dotfile mgr | None configured. If user wants one, suggest `chezmoi` or GNU `stow`. |

---

## Editors

Both `micro` and `neovim` are installed.

- **micro**: config in `~/.config/micro/` (`settings.json`, `bindings.json`). Colorschemes in `~/.config/micro/colorschemes/<name>.micro`.
- **neovim**: config in `~/.config/nvim/` (`init.lua`). Colorschemes in `~/.config/nvim/colors/<name>.lua` or via plugin.

When the user requests theming, apply the **Forest / Nature Theme** below.

---

## Ricing / Customization Conventions

- **GTK overrides**: drop CSS in `~/.config/gtk-3.0/gtk.css`, `~/.config/gtk-4.0/gtk.css`, and `~/.config/gtk-4.0/libadwaita.css` (libadwaita apps need their own).
- **Plasma color scheme**: write a `.colors` file to `~/.local/share/color-schemes/Forest.colors`, then apply with `plasma-apply-colorscheme Forest`.
- **Plasma global theme** (icons, splash, colors as a bundle): `~/.local/share/plasma/look-and-feel/`, apply with `plasma-apply-lookandfeel -a <id>`.
- **Konsole scheme**: drop a `.colorscheme` file in `~/.local/share/konsole/Forest.colorscheme`, then in Konsole > Settings > Edit Profile > Appearance.
- **Cursor theme**: `~/.icons/<name>/cursors/`, set in `~/.config/gtk-3.0/settings.ini` and `~/.icons/default/index.theme`.
- **Always test in a new session/window** — Plasma can cache aggressively. `kquitapp6 plasmashell && kstart plasmashell` reloads the shell without logging out.

---

## Forest / Nature Theme

Greens-dominant palette with a small brown accent family and grayscale neutrals. No blues / purples / yellows — brown does double duty as warm accent and error/warning family. Greens are stratified by hue (yellow-green / true green / blue-green) and lightness so adjacent tokens never collide. Designed for dark backgrounds.

### Base / UI
| Hex | Name | Role |
|---|---|---|
| `#181818` | Deep Loam | bg | *Use if colorscheme will be shown within a terminal.*
| `#222826` | Damp Bark | bg alt |
| `#2c332f` | Forest Floor | bg light variant |
| `#d8dcc9` | Birch Paper | fg |
| `#8a9085` | Drift Ash | fg muted |
| `#3a4a3e` | Wet Moss Shadow | selection |
| `#c8d4a8` | New Shoot | cursor |
| `#3a423d` | Lichen Edge | border |

### Greens
| Hex | Name | Notes | Role |
|---|---|---|---|
| `#8fb37a` | Moss | warm yellow-green, mid-light | strings |
| `#b5c97a` | New Growth | bright lime, highest chroma | numbers / constants |
| `#6f9a6a` | Fern | true saturated green, anchor | keywords |
| `#a3b89a` | Sage | desaturated cool green | variables |
| `#5d8a78` | Lichen | blue-green, distinctly cooler | functions / types |
| `#3f6b55` | Pine | deep evergreen | status accent |
| `#9ba87a` | Olive | yellow-olive, low chroma | operators / punctuation |
| `#6a7866` | Comment Moss | grey-green, deliberately mute | comments (italic) |

### Browns (accent + diagnostic family)
| Hex | Name | Role |
|---|---|---|
| `#c47a4a` | Hot Bark | errors (warm terracotta) |
| `#a8895a` | Oak | warnings (muted ochre) |
| `#7a5a3e` | Walnut | diff remove / deletion |

### Grayscale
| Hex | Name | Role |
|---|---|---|
| `#0e110f` | True Pitch | ANSI black |
| `#4a514c` | Slate Stone | bright black, dim divider |
| `#6e756f` | River Pebble | mid-grey |
| `#a8aea3` | Ash Grey | light neutral |
| `#e6e8dc` | Bone | off-white |

### Syntax mapping (no new hexes)
- keyword → **Fern**
- string → **Moss**
- comment → **Comment Moss** (italic)
- function → **Lichen**
- type / class → **Lichen** (bold)
- constant / number → **New Growth**
- variable → **Sage**
- operator / punctuation → **Olive**
- error → **Hot Bark** (+ undercurl)
- warning → **Oak** (+ underdash)
- info / hint → **Drift Ash** (+ underdot)
- diff add → **New Growth**
- diff remove → **Walnut**
- status accent → **Pine**
- prompt accent → **New Shoot**

### Terminal ANSI
```
0  #0e110f  True Pitch          8  #4a514c  Slate Stone
1  #c47a4a  Hot Bark    (red)   9  #7a5a3e  Walnut       (br-red)
2  #6f9a6a  Fern        (grn) 10  #b5c97a  New Growth   (br-grn)
3  #a8895a  Oak         (yel) 11  #c8d4a8  New Shoot    (br-yel)
4  #5d8a78  Lichen      (blu) 12  #3f6b55  Pine         (br-blu)
5  #9ba87a  Olive       (mag) 13  #8fb37a  Moss         (br-mag)
6  #a3b89a  Sage        (cyn) 14  #6a7866  Comment Moss (br-cyn)
7  #d8dcc9  Birch Paper       15  #e6e8dc  Bone
```

ANSI mapping logic: red slot → Hot Bark (errors live here), yellow slot → Oak (warnings), blue slot → Lichen (coolest hue available), magenta/cyan → remaining distinct greens. Severity is never color-only — always pair with undercurl/underdash glyphs for accessibility.

---

## Verification & Update Discipline

- **Before pkg changes**: `pacman -Rcs --print <pkg>` (dry-run), `pacman -Qqe > /tmp/pkgs-$(date +%F).txt` (snapshot).
- **Before editing system files**: copy first — `cp /etc/foo.conf{,.bak-$(date +%F)}`.
- **After service changes**: `rc-service <name> status` to confirm, `rc-status` for runlevel view.
- **After font installs**: `fc-cache -fv`.
- **After GTK theme changes**: log out / log back in (Plasma re-applies on session start).
- **After kernel / mesa updates**: reboot before opening anything graphics-heavy. Confirm `vulkaninfo --summary` works post-boot.
- **For risky changes**: tell the user what you're about to do *and* what could break, then wait for confirmation. No snapshots = no undo.
