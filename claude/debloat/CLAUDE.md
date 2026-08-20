# Artix Plasma De-lag & De-brand — System State

This directory documents work done on `woody` on **2026-05-01** to fix UI lag and strip Artix-specific branding from a fresh Plasma + OpenRC install. If the user is reporting "the system feels laggy again" or "something theme-related broke," start here.

---

## Machine

| | |
|---|---|
| Hostname | `woody` |
| Distro | Artix Linux (rolling), `plasma-openrc` variant |
| Init | **OpenRC** (NOT systemd) — use `rc-update`, `rc-service`, configs in `/etc/conf.d/`, never `systemctl` |
| Session | Wayland (KWin) |
| Kernel | `linux` package (was `6.19.14` after work; running kernel may differ until reboot) |
| CPU | AMD Ryzen 7 7800X3D (16 threads) |
| GPU (primary) | **NVIDIA RTX 5070 Ti (Blackwell GB203)** — requires `nvidia-open-dkms`, NOT `nvidia-580xx-dkms` or `nouveau` |
| GPU (iGPU) | AMD Raphael (`amdgpu`) |
| RAM | 64 GiB, no swap |
| Storage | Samsung 990 PRO 2 TB NVMe, ext4 root, EFI/GRUB |
| User | `void` |
| Sudo | Password required. Use `SUDO_ASKPASS=/usr/bin/ksshaskpass sudo -A …` and **batch privileged work into a single `sudo bash -c '…'`** so the user only types the password once |
| Multilib | **Disabled.** Enable in `/etc/pacman.conf` before installing any 32-bit packages (Steam, lib32-nvidia-utils) |

---

## Diagnosis recap (the lag root cause)

If lag returns, this is the order to check:

1. **GPU driver** — primary suspect. The RTX 5070 Ti is Blackwell; `nouveau` falls back to `llvmpipe` software rendering and pegs `kwin_wayland` at 600%+ CPU.
   ```sh
   glxinfo | grep "OpenGL renderer"   # must say "NVIDIA GeForce RTX 5070 Ti"
   lspci -k | grep -A2 NVIDIA          # "Kernel driver in use: nvidia", NOT "nouveau"
   nvidia-smi                          # should list kwin_wayland, plasmashell
   dkms status                         # should show "nvidia/<ver>, <kernel>: installed"
   ```
   If any of these revert, see "Recovering NVIDIA driver" below.

2. **CPU governor**:
   ```sh
   cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor   # should be "performance"
   rc-service cpupower status
   ```

3. **DKMS module not built for current kernel**: After a kernel update, DKMS rebuilds automatically via the pacman hook. If `dkms status` shows "WARNING" or no entry for the running kernel, run `sudo dkms autoinstall`.

---

## Changes made (2026-05-01)

### Driver / performance
- Installed: `nvidia-open-dkms`, `nvidia-utils`, `egl-wayland`, `cpupower`, `cpupower-openrc`.
- `/etc/modprobe.d/nouveau-blacklist.conf` (new):
  ```
  blacklist nouveau
  options nouveau modeset=0
  ```
- `/etc/modprobe.d/nvidia.conf` (new):
  ```
  options nvidia_drm modeset=1 fbdev=1
  ```
- `/etc/mkinitcpio.conf`: `MODULES=(nvidia nvidia_modeset nvidia_uvm nvidia_drm)`. Initramfs regenerated with `mkinitcpio -P`.
- `/etc/conf.d/cpupower`: `START_OPTS="--governor performance"`. Service enabled in `default` runlevel.

### Visual de-branding
- `/etc/sddm.conf` → `[Theme] Current=breeze` (was `artix`).
- `/etc/default/grub` → `GRUB_THEME` line commented out. `grub-mkconfig -o /boot/grub/grub.cfg` regenerated.
- `~/.config/kdeglobals` → rewritten as minimal Breeze-Dark config. **Original preserved at `~/.config/kdeglobals.artix-backup.bak`**.
- `~/.config/plasma-org.kde.plasma.desktop-appletsrc` → removed Kicker `icon=` line and Image/PreviewImage wallpaper keys.

### Packages removed (`pacman -Rns`)
```
artix-gtk-presets        artix-qt-presets        artix-desktop-presets
artix-dark-theme         artix-plasma-splash     artix-sddm-theme
artix-wallpapers         artix-backgrounds       artix-icons
artix-grub-theme         artix-branding-base
```

### Packages **deliberately kept** (do NOT remove)
| Package | Why |
|---|---|
| `artix-keyring` | PGP keys for pacman repos |
| `artix-mirrorlist` | pacman repo definitions |
| `base`, `filesystem` | Core system meta-packages |
| `qt6gtk2` | Qt-to-GTK2 style bridge. Currently unused (widgetStyle=Breeze) but harmless; user may remove with `pacman -R qt6gtk2` if desired |

### Files preserved from `artix-branding-base` (renamed to survive package removal)
- `/etc/sysctl.d/99-magic_sysrq-local.conf` — `kernel.sysrq=1` (magic SysRq emergency keys)
- `/etc/local.d/cleanup-local.start` — clears stale `/var/lib/pacman/db.lck` on boot

### Files dropped from `artix-branding-base` (cosmetic only)
- `/etc/DIR_COLORS` (ls coloring)
- `/etc/bash/bashrc.d/local.bashrc` (colored bash prompt)
- `/etc/local.d/branding.start` (fastfetch into /etc/issue)
- `/etc/local.d/local.start` + `local.stop` (sysv-style `/etc/rc.local` shim — user has no `/etc/rc.local`)
- `/etc/udev/rules.d/80-net-name-slot.rules` (empty mask file; redundant with `net.ifnames=0` on kernel cmdline)

---

## User preferences (durable, apply to future Plasma work on this machine)

1. **Blank visual slate.** Vanilla Breeze everywhere. Do not suggest theming, icon packs, conky, ricing, or opinionated visuals unless the user explicitly asks. The user themes the system themselves.
2. **Functional tuning is welcome and separate** — sysctl, udev, perf governors, kernel options, etc. are not "bloat."
3. **Minimize sudo prompts.** Batch privileged work into one `sudo bash -c '…'` invocation. Each separate sudo call re-prompts ksshaskpass.
4. **OpenRC, not systemd.** Don't propose `systemctl` commands or write `.service` files. Use `/etc/init.d/`, `/etc/conf.d/`, `rc-update`, `rc-service`.

---

## Recovery cheatsheet

### NVIDIA driver broke after a kernel/pacman upgrade
```sh
SUDO_ASKPASS=/usr/bin/ksshaskpass sudo -A bash -c '
  pacman -Syu --noconfirm
  dkms autoinstall
  mkinitcpio -P
'
# reboot
```
Black screen at SDDM after reboot? Drop to TTY (Ctrl+Alt+F2), inspect:
```sh
journalctl -k -b   # or: dmesg | grep -iE "nvidia|nouveau|drm"
lsmod | grep nvidia
cat /etc/modprobe.d/nouveau-blacklist.conf   # should still exist
```

### Plasma theming reset to something Artix-y
The package is gone; nothing should reapply it. If it does, look for stray dotfiles in `/etc/skel/` (those came from `artix-desktop-presets` originally — package is removed, but check anyway). To re-apply blank slate:
```sh
cp ~/.config/kdeglobals ~/.config/kdeglobals.bak
# Restore the minimal version below into ~/.config/kdeglobals, then logout/login.
```
Reference minimal `kdeglobals` is at `~/.config/kdeglobals.artix-backup.bak` — wait, that's the *original Artix* one. The minimal Breeze version is just whatever `kdeglobals` is now (assuming it hasn't been overwritten). If both are gone, Plasma regenerates a default on next login.

### Lag returned
Run the diagnosis block at the top. 99% of the time it's the GPU driver. CPU governor reverting to `powersave` on its own would mean the cpupower OpenRC service stopped — `rc-service cpupower start; rc-update show | grep cpupower`.

### Re-enable an Artix theme (user changed their mind)
```sh
sudo pacman -S artix-dark-theme artix-sddm-theme artix-plasma-splash artix-wallpapers artix-backgrounds artix-icons
# then in System Settings → Global Theme, pick "Artix Dark"
# SDDM: edit /etc/sddm.conf [Theme] Current=artix
```

---

## What this directory is NOT

- Not a place to dump every future change. It is *specifically* the de-bloat work from 2026-05-01.
- Future, unrelated projects belong in their own subdirectory of `~/claude/`.
