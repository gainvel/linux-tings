#!/bin/sh
# restore-void-lockscreen.sh
# Re-apply the custom "void" lock screen to the Plasma desktop shell package.
#
# Why this exists: the KDE lock screen greeter (kscreenlocker_greet) loads its QML
# from the plasma desktop SHELL package, not from a user look-and-feel package:
#   /usr/share/plasma/shells/org.kde.plasma.desktop/contents/lockscreen/
# Those files are owned by `plasma-desktop` and are NOT pacman backup files, so every
# `plasma-desktop` upgrade silently overwrites them back to stock (this is what reverted
# the void lock screen to the default Breeze screen). This script restores the void files
# from a stable, pacman-untouched staging dir. It is invoked by a PostTransaction pacman
# hook on plasma-desktop (see void-lockscreen.hook).
SRC=/usr/local/share/void-lockscreen/lockscreen
DEST=/usr/share/plasma/shells/org.kde.plasma.desktop/contents/lockscreen
[ -d "$SRC" ]  || { echo "void-lockscreen: staging $SRC missing, skipping"  >&2; exit 0; }
[ -d "$DEST" ] || { echo "void-lockscreen: target  $DEST missing, skipping" >&2; exit 0; }
cp -a "$SRC/." "$DEST/"
echo "void-lockscreen: restored void lock screen to $DEST"
