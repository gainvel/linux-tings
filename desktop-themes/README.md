# Desktop themes — `void` SDDM login + matching KDE lock screen

Master copies of the custom **void** theming on `woody` (Artix, KDE Plasma 6), plus the
greeter monitor-layout fix. Kept here so future changes are clear and so the system can be
rebuilt if a package update wipes the in-place files (which has happened — see below).

## What's here / where it installs

| Source in this repo | Installs to | Enabled by |
|---|---|---|
| `sddm-void/` | `/usr/share/sddm/themes/void/` | `/etc/sddm.conf` → `[Theme] Current=void` |
| `lockscreen-void/` | `/usr/share/plasma/shells/org.kde.plasma.desktop/contents/lockscreen/` **and** staged at `/usr/local/share/void-lockscreen/lockscreen/` | loaded automatically by `kscreenlocker_greet` from the plasma **shell** package; the pacman hook re-applies it after `plasma-desktop` upgrades |
| `sddm-Xsetup` | `/usr/share/sddm/scripts/Xsetup` | run automatically by SDDM at X11 greeter startup |
| `pacman-hook/void-lockscreen.hook` | `/etc/pacman.d/hooks/void-lockscreen.hook` | pacman runs it after every `plasma-desktop` install/upgrade |
| `pacman-hook/restore-void-lockscreen.sh` | `/usr/local/bin/restore-void-lockscreen.sh` | invoked by the hook |

## Why the lock screen kept reverting (important)

The KDE lock screen greeter loads its QML from the **plasma desktop shell package**
(`org.kde.plasma.desktop`), *not* from a user `look-and-feel` package — so the void
lock screen has to live under `/usr/share/plasma/shells/.../contents/lockscreen/`.

Those files are owned by `plasma-desktop` and are **not** pacman *backup* files, so a
`plasma-desktop` upgrade silently overwrites them back to stock (this is exactly what
reverted the lock screen to the default Breeze screen on 2026-05-12).

The fix for that recurrence is the **pacman hook**: after any `plasma-desktop` transaction,
pacman runs `restore-void-lockscreen.sh`, which copies the void files from the stable staging
dir `/usr/local/share/void-lockscreen/lockscreen/` back into place. To change the lock screen,
edit the staging copy (or this repo, then re-stage) so the hook keeps applying your version.

The SDDM `Xsetup` script, by contrast, **is** a pacman backup file, so edits there survive
sddm upgrades and need no hook.

## The lock screen vs. the SDDM theme

The lock screen (`lockscreen-void/`) is a QML clone of the SDDM `void` theme. The only
intended difference: the SDDM theme's 4th bottom-bar control (the **session / DE switcher**)
is omitted, since you can't switch session on a lock screen. The power bar therefore has just
suspend / reboot / shutdown.

## Re-applying from scratch

```sh
ASK="SUDO_ASKPASS=/usr/bin/ksshaskpass sudo -A"
R=~/Repo/desktop-themes

# SDDM login theme
$ASK cp -a "$R/sddm-void" /usr/share/sddm/themes/void

# Lock screen: stage + install + hook
$ASK sh -c '
  install -d /usr/local/share/void-lockscreen &&
  cp -a "'$R'/lockscreen-void" /usr/local/share/void-lockscreen/lockscreen &&
  cp -a /usr/local/share/void-lockscreen/lockscreen/. \
        /usr/share/plasma/shells/org.kde.plasma.desktop/contents/lockscreen/ &&
  install -Dm755 "'$R'/pacman-hook/restore-void-lockscreen.sh" /usr/local/bin/restore-void-lockscreen.sh &&
  install -Dm644 "'$R'/pacman-hook/void-lockscreen.hook" /etc/pacman.d/hooks/void-lockscreen.hook
'

# Greeter monitor layout
$ASK cp "$R/sddm-Xsetup" /usr/share/sddm/scripts/Xsetup
$ASK chmod 755 /usr/share/sddm/scripts/Xsetup
```

> Note: `lockscreen-void/images/wallpaper.jpg` is ~12 MB. Committing it bloats the `~/Repo`
> git history — consider a dedicated repo or gitignoring the binary if that matters.
