# CLAUDE.md — Game Modding Workspace

Workspace for researching and modifying **offline / single-player games** installed through Steam on this Artix Linux desktop. This file is the map: where each game's files, saves, and mod targets actually live, and the rules for touching them.

System-level rules (OpenRC, pacman, GPU, storage) live in `/home/void/claude/general/CLAUDE.md` — that file stays authoritative; this one never restates it.

---

## Critical Rules

**1. Never write to an anti-cheat / online title.** For the six games in the [Blocked table](#blocked-titles--read-only-study-only), reading, dumping, parsing, and documenting is encouraged; writing is not. Zero modifications to their install dirs, prefixes, saves, configs, or cloud data — and never launch them with an injected library, patched binary, or altered pak. EAC / BattlEye / VAC / miHoYo AC flag altered files and ban the **account**, which is permanent and unrecoverable.

**2. Follow this order for every save edit** — skipping a step is how saves get lost:
   1. Fully close the game **and** Steam-Cloud-sync it (quit to desktop, wait for the sync toast).
   2. Disable Steam Cloud for that appid: Steam → game → Properties → General → uncheck *Keep game saves in the Steam Cloud*.
   3. Copy the original into `backups/<game>/<YYYY-MM-DD>/` — pristine, never edited.
   4. Copy again into `work/<game>/` and edit **that** copy.
   5. Install the edited copy over the live path, then launch and verify in-game.

**3. Windows games read from their Proton prefix, not `$HOME`.** A game's `Documents` is `compatdata/<appid>/pfx/drive_c/users/steamuser/Documents` — a real directory that is **not** `~/Documents` (a symlink to `/mnt/storage1/Documents`). Editing the `$HOME` one changes a file the game never opens, and the "fix" silently does nothing.

**4. Treat `compatdata/<appid>/` as save data, not cache.** Deleting a prefix, or switching a game to a different Proton version, gives the game a **new empty prefix** — every save inside the old one disappears from the game's view. Copy the prefix to `backups/` before any Proton-version experiment.

**5. Rebuild every mod from `work/` — never only from the live game dir.** Steam updates and *Verify integrity of game files* overwrite anything under `steamapps/common/<game>/`. Anything not reproducible from `work/` + `backups/` is one update away from gone.

**6. Confirm before installing tooling.** Check the [Tooling table](#tooling) first; if something is genuinely missing, ask before running `yay -S` / `pip install`. This box has no snapshots, so package churn is unwound by hand.

**7. Nothing in this workspace is backed up.** `~/claude/modding` is absent from `~/.config/storage-backup.list`, so the nightly tiered rsync skips it entirely. Irreplaceable artifacts go to `~/Backups/Manual/` (mirrored) — losing `backups/` means losing the only clean copy of a save.

---

## Workspace Layout

| Path | Purpose |
|---|---|
| `CLAUDE.md` | This file — paths, rules, per-game map |
| `INVENTORY.md` | **Generated.** Current installed-games table; regenerate, never hand-edit |
| `scan-library.sh` | Read-only regenerator for `INVENTORY.md` |
| `work/<game>/` | Scratch: extractions, decompiles, edits in progress. Disposable |
| `backups/<game>/<date>/` | Pristine originals copied before any edit. Never edited in place |
| `notes/<game>.md` | Research write-ups: file formats, offsets, what worked |

Refresh the inventory after installing or removing anything: `./scan-library.sh`

---

## Steam Paths

One library folder only — everything is under `~/.local/share/Steam`. Below, **`STEAMAPPS`** = `~/.local/share/Steam/steamapps` and **`PFX(id)`** = `STEAMAPPS/compatdata/<id>/pfx/drive_c/users/steamuser`.

| What | Path |
|---|---|
| Steam root | `~/.local/share/Steam` (`~/.steam/root`, `~/.steam/steam` are symlinks to it) |
| Game installs | `STEAMAPPS/common/<installdir>/` |
| Install manifests | `STEAMAPPS/appmanifest_<appid>.acf` (name, installdir, size, buildid) |
| Library index | `STEAMAPPS/libraryfolders.vdf` — the authority on what is installed |
| Proton prefixes | `STEAMAPPS/compatdata/<appid>/pfx/` (`drive_c/users/steamuser/…`) |
| Steam Cloud staging | `~/.local/share/Steam/userdata/488557534/<appid>/` |
| Custom Proton builds | `~/.local/share/Steam/compatibilitytools.d/` (`GE-Proton10-27`, `GE-Proton9-13`) |
| Compat tool mapping | `~/.local/share/Steam/config/config.vdf` → `CompatToolMapping` |
| Steam client logs | `~/.local/share/Steam/logs/` |

**Proton assignment**: the global default is `proton_11` (applies to every game with no override); the only per-game override is appid `1076160` (Command: Modern Operations) → `GE-Proton10-27`. See the memory note on CMO before touching that prefix.

---

## Moddable Offline Games

Save paths are the **external** state — most of what a mod or trainer changes is here, not in the install dir. `PFX` = that game's prefix (shorthand above). Everything else is under `STEAMAPPS/common/`.

| AppID | Game | Engine | Install dir | Saves & config | Mod target |
|---|---|---|---|---|---|
| 673610 | Airport CEO | Unity Mono | `Airport CEO` | *never launched* — created on first run | `Airport CEO_Data/Managed/Assembly-CSharp.dll`; in-game mod loader |
| 550320 | art of rally | Unity Mono, **native** | `artofrally` | `~/.config/unity3d/Funselektor Labs/art of rally/` | `artofrally_Data/Managed/` |
| 1489630 | Carrier Command 2 | Custom (Lua + XML) | `Carrier Command 2` | `PFX(1489630)/AppData/Roaming/Carrier Command 2/` → `saved_games/`, `mods.xml`, `settings.xml` | `Carrier Command 2/rom_0/` (game_objects, behaviour_trees); `mod_dev_kit/` ships official tooling |
| 953490 | CARRION | Custom .NET, **native** | `Carrion` | `~/.local/share/Phobia/Carrion/` | `Carrion/Content/`; `Carrion.dll` (.NET) |
| 1076160 | Command: Modern Operations | Custom .NET (WinForms) | `Command - Modern Operations` | `PFX(1076160)/Documents/My Games/Command Modern Operations/` | Scenario `.scen` files + Lua; `AttachmentRepo/`. **Pinned to GE-Proton10-27** |
| 4170200 | Data Center | Unity IL2CPP | `Data Center` | `PFX(4170200)/AppData/LocalLow/WASEKU/Data Center/` (`saves/`, `Player.log`) | **Official no-code mod loader** — `Data Center_Data/StreamingAssets/Mods/<folder>/config.json` + `.obj` + `.png` builds new shop items (see `notes/data-center-modding.md`). Wiped by Steam update / *Verify*. Code mods use MelonLoader, not the config's `dlls` |
| 2300 | DOOM II | id Tech 1 (DOSBox) | `Doom 2` | `PFX(2300)/Saved Games/id Software/DOOM 2/` | WAD/PWAD files in `Doom 2/base/`; launch `.bat` wrappers |
| 1239080 | Door Kickers 2 | Custom (Lua + XML) | `DoorKickers2` | `PFX(1239080)/AppData/Local/KillHouseGames/DoorKickers2/` | `DoorKickers2/data/` — plain-text Lua/XML, very mod-friendly |
| 1943950 | Escape the Backrooms | Unreal 4.27 | `EscapeTheBackrooms` | `PFX(1943950)/AppData/Local/EscapeTheBackrooms/Saved/` | `.pak` in `<Game>/Content/Paks/~mods/` |
| 427520 | Factorio | Custom (Lua), **native** | `Factorio` | `~/.factorio/` → `saves/`, `mods/`, `config/`, `player-data.json` | `~/.factorio/mods/<name>_<ver>/` (zip or dir + `info.json`) |
| 2483190 | Forza Horizon 6 | ForzaTech | `ForzaHorizon6` | `PFX(2483190)/AppData/Local/ForzaHorizon6/` — **PlayFab, server-authoritative** | Local edits to `PlayFabSaveStorage_*` are overwritten by the server; treat progression as untouchable and stay offline |
| 16900 | GROUND BRANCH | Unreal 4.27 | `Ground Branch` | `PFX(16900)/AppData/Local/GroundBranch/Saved/` **and** `PFX(16900)/Documents/GroundBranch/` (`Loadouts`, `Presets`, `ItemBuilds`) | `GroundBranch/Content/Paks/`; Lua mission scripts |
| 1705180 | Gunner, HEAT, PC! | Unity Mono | `Gunner HEAT PC` | `PFX(1705180)/AppData/LocalLow/Radian Simulations LLC/GHPC/` **and** `PFX(1705180)/Documents/My Games/GHPC/` (`Data`, `Editor`) | `Documents/My Games/GHPC/Data` holds user content; `Bin/` for assemblies |
| 2726490 | Hamster Hunter | Unity IL2CPP, **native** | `Hamster Hunter` | `~/.config/unity3d/Hamuno/Hamster Hunter/` | `hh_Data/`; IL2CPP binary |
| 394360 | Hearts of Iron IV | Clausewitz, **native** | `Hearts of Iron IV` | `~/.local/share/Paradox Interactive/Hearts of Iron IV/` → `save games/`, `settings.txt`, `dlc_load.json` | `…/Hearts of Iron IV/mod/<name>/` + `<name>.mod` descriptor (dir is created on first mod); base data is plain text under the install dir |
| 219150 | Hotline Miami | GameMaker, **native** | `hotline_miami` | *never launched* — created on first run | `HotlineMiami_GL.wad` (asset archive) |
| 4300500 | IRON NEST demo | Unity IL2CPP | `IRON NEST Heavy Turret Simulator Demo` | `PFX(4300500)/AppData/LocalLow/Iron Nest/` | Demo build — expect no mod support |
| 3278310 | LANESPLIT | Unreal 5.6 | `LANESPLIT` | `PFX(3278310)/AppData/Local/LANESPLIT/Saved/` | `LANESPLIT/…/Content/Paks/` |
| 4704690 | MECCHA CHAMELEON | Unreal 5.6 | `MECCHA CHAMELEON` | `PFX(4704690)/AppData/Local/Chameleon/Saved/` | `…/Content/Paks/` |
| 1129580 | Medieval Dynasty | Unreal 4.27 | `Medieval Dynasty` | `PFX(1129580)/AppData/Local/Medieval_Dynasty/Saved/` (`SaveGames/`, `Config/`) | `Medieval_Dynasty/Content/Paks/~mods/` |
| 2168680 | Nuclear Option | Unity Mono | `Nuclear Option` | `PFX(2168680)/AppData/LocalLow/Shockfront/NuclearOption/` | **BepInEx already installed** (`BepInEx/`, `doorstop_config.ini`) — drop plugins in `BepInEx/plugins/` |
| 1913370 | OPERATOR | Unity IL2CPP | `OPERATOR` | `PFX(1913370)/AppData/LocalLow/VECTOR INTERACTIVE/OPERATOR/` (+ `AppData/LocalLow/Unity/VECTOR INTERACTIVE_OPERATOR/`) | IL2CPP — save/config editing is far cheaper than code patching |
| 1623730 | Palworld | Unreal 5.1 | `Palworld` | `PFX(1623730)/AppData/Local/Pal/Saved/SaveGames/` (`.sav`, GVAS) | `Pal/Content/Paks/~mods/`; `.sav` files are GVAS — decode before editing |
| 108600 | Project Zomboid | Java + Lua, **native** | `ProjectZomboid` | `~/Zomboid/` → `Saves/`, `mods/`, `Lua/`, `options.ini`, `Sandbox Presets/` | `~/Zomboid/mods/<name>/` (`mod.info` + `media/`); base Lua under `ProjectZomboid/projectzomboid/media/lua/` |
| 1144200 | Ready or Not | Unreal 5.3 | `Ready Or Not` | `PFX(1144200)/AppData/Local/ReadyOrNot/Saved/` | `PFX(1144200)/AppData/Local/mod.io/3791/` is the mod.io cache; loose paks → `ReadyOrNot/Content/Paks/~mods/` |
| 2067050 | Squirrel with a Gun | Unreal 5.2 | `SquirrelWithAGun` | `PFX(2067050)/AppData/Local/SquirrelGun/Saved/` (+ `CloudSaveable/`) | `…/Content/Paks/` |
| 573090 | Stormworks | Custom (Lua + XML) | `Stormworks` | `PFX(573090)/AppData/Roaming/Stormworks/` → `saves/`, `save.xml`, `mods.xml`, `settings.xml`, `persistent_data.xml` | `Stormworks/rom/` (definitions, scripts) — all plain text |
| 1593030 | Terra Nil | Unity Mono, **native** | `Terra Nil` | `~/.config/unity3d/Free Lives/Terra Nil/` | `Terra Nil_Data/Managed/` |
| 518790 | theHunter: Call of the Wild | Apex (Avalanche) | `theHunterCotW` | `PFX(518790)/Documents/Avalanche Studios/theHunter Call of the Wild/` **and** `…/Avalanche Studios/COTW/` | `archives_win64/*.tab`/`*.arc` — proprietary; extract to `work/` before touching |
| 2198150 | Tiny Glade | Custom (Rust), **native** | `Tiny Glade` | `~/.local/share/Tiny Glade/` | No mod API; asset/save research only |
| 3115220 | Town To City | Unreal 5.7 | `Town To City` | `PFX(3115220)/AppData/Local/TownToCity/Saved/` | `…/Content/Paks/` |
| 1611600 | WARNO | Eugen (custom) | `WARNO` | `PFX(1611600)/Saved Games/EugenSystems/WARNO/` | `WARNO/Mods/` — official tooling: `CreateNewMod.bat`, `ModData/`, `ModRevisionList.ini` |
| 424030 | War of Rights | CryEngine | `War of Rights` | `PFX(424030)/Saved Games/warofrights/` | `War of Rights/Assets/` (`.pak`). **MP-focused**: study offline; never join a server with a modified client |
| 228380 | Wreckfest | Bugbear (custom) | `Wreckfest` | `PFX(228380)/AppData/LocalLow/THQNordic/Wreckfest/` | `Wreckfest/data/` + `BagEdit` (official archive/mod tool that ships with the game) |

`1174860` **Tacview** is an analysis tool, not a game — its data is `PFX(1174860)/AppData/Roaming/Tacview/`. Proton runtimes, Steam Linux Runtimes, soundtracks, and Steamworks redistributables are omitted; `INVENTORY.md` lists them.

---

## Blocked Titles — Read-Only Study Only

Study freely: dump archives, parse formats, read configs, write up findings in `notes/`. Produce **zero** writes to these paths, and never launch them with anything injected or patched.

| AppID | Game | Anti-cheat | Evidence on disk |
|---|---|---|---|
| 393380 | Squad | EasyAntiCheat | `Squad/EasyAntiCheat/`, `UninstallAntiCheat.bat`, `PFX/AppData/Roaming/EasyAntiCheat` |
| 1604270 | Broken Arrow | EasyAntiCheat (EOS) | `broken_arrow/EasyAntiCheat/EasyAntiCheat_EOS_Setup.exe` |
| 1874880 | Arma Reforger | BattlEye | `Arma Reforger/battleye/` |
| 236390 | War Thunder | BattlEye | `War Thunder/BattlEye/` |
| 730 | Counter-Strike 2 | VAC (server-side, no local binary) | — |
| 4162040 | Zenless Zone Zero | miHoYo AC + server-authoritative account | `PFX/AppData/LocalLow/miHoYo/` |

Copy anything you want to inspect into `work/<game>/` first and analyze the copy — that keeps file mtimes and integrity checks in the install dir untouched.

---

## Engine → Technique

| Engine | Identify by | Where the moddable content is | Approach |
|---|---|---|---|
| Unity Mono | `UnityPlayer.dll` + `<Game>_Data/Managed/Assembly-CSharp.dll` | `Managed/` assemblies; `<Game>_Data/*.assets`, `*.bundle` | BepInEx/Harmony plugin (no assembly rewrite); UnityPy for assets |
| Unity IL2CPP | `GameAssembly.dll`/`.so` + `<Game>_Data/il2cpp_data/Metadata/global-metadata.dat` | Native binary — no C# assemblies | Much harder: BepInEx-IL2CPP + Il2CppDumper. Prefer save/config edits |
| Unreal 4/5 | `<Game>/Content/Paks/*.pak`, `Binaries/Win64/` | Paks; `Saved/Config/WindowsNoEditor/*.ini`; `Saved/SaveGames/*.sav` | Loose `.pak` into `Content/Paks/~mods/` (loads after base paks). `.ini` tweaks need no packing. `.sav` = GVAS binary |
| Clausewitz (HOI4) | `clausewitz_rev.txt`, plain-text `common/` tree | Everything is text | Copy a base file into `mod/<name>/` mirroring its relative path; add a `.mod` descriptor |
| Lua/XML data-driven (Door Kickers 2, Carrier Command 2, Stormworks) | `data/` or `rom/` of plain text | Text definitions | Edit text directly; keep diffs in `work/` so updates can be re-applied |
| Java + Lua (Project Zomboid) | `projectzomboid.sh`, `media/lua/` | `media/lua/{client,server,shared}` | Ship a `~/Zomboid/mods/<name>/` overlay; never edit base `media/` |
| id Tech 1 (DOOM II) | `.wad`, DOSBox `.bat` wrappers | WAD lumps | PWAD alongside the IWAD |
| Custom archives (theHunter `.tab/.arc`, Wreckfest, War of Rights `.pak`) | Proprietary containers | Archive contents | Extract into `work/` first; use the game's own tool when it ships one (Wreckfest `BagEdit`, WARNO `CreateNewMod.bat`) |

---

## Tooling

| Available now | Notes |
|---|---|
| `python3`, `pip`, `node`, `go`, `cargo`/`rustc` | Scripting + writing custom parsers |
| `7z`, `unzip` | Zip/7z archives, many `.pak` containers |
| `xxd`, `strings`, `gdb`, `jq` | Binary inspection, save-format reverse engineering, JSON saves |

| Not installed | Get it with (ask first) |
|---|---|
| UnityPy (Unity asset read/write) | `pip install UnityPy` |
| protontricks | Flatpak only on this box: `flatpak run com.github.Matoking.protontricks` |
| unrar, radare2 | `yay -S unrar radare2` |
| .NET/Mono decompiler (ILSpy, dnSpy) | `yay -S avaloniailspy` or use `ilspycmd` via `dotnet tool` |
| UE pak tools (repak, UnrealPak) | `cargo install repak_cli`, or the engine's `UnrealPak` |

There is no standalone `wine`/`winetricks` — Windows tooling runs through a game's Proton prefix (`STEAM_COMPAT_DATA_PATH=STEAMAPPS/compatdata/<appid>` + the Proton build under `compatibilitytools.d/`).

---

## Common Operations

```bash
# Refresh the installed-games inventory
~/claude/modding/scan-library.sh

# Find where an unlisted/newly launched game actually writes (files touched since last launch)
APPID=1129580; PFX=~/.local/share/Steam/steamapps/compatdata/$APPID/pfx/drive_c/users/steamuser
find "$PFX" -type f -newermt '-2 hours' -not -path '*/Temp/*' | head -50

# Back up a save before editing (step 3 of the save workflow)
cp -a "$PFX/AppData/Local/Medieval_Dynasty/Saved" \
      ~/claude/modding/backups/medieval-dynasty/$(date +%F)/

# Back up an entire prefix before a Proton-version change
cp -a ~/.local/share/Steam/steamapps/compatdata/$APPID{,.bak-$(date +%F)}

# Launch a game to verify a change
steam -applaunch $APPID

# Read a game's own log (Unreal / Unity)
ls "$PFX"/AppData/Local/*/Saved/Logs/            # Unreal
ls "$PFX"/AppData/LocalLow/*/*/Player.log        # Unity

# Inspect an unknown save format
xxd -l 256 save.sav; strings -n 6 save.sav | head -40
```

To force a specific Proton build for a one-off run, set it in Steam → Properties → Compatibility rather than by hand — **and read Gotchas below first**, because switching builds creates a fresh prefix.

---

## Gotchas

**Steam Cloud silently reverts save edits.** It is the single most common cause of "my edit didn't take": the client re-uploads (or re-downloads) `userdata/488557534/<appid>/` on launch and quit, overwriting the local file. Disable cloud for the appid *before* editing, and re-enable only after verifying in-game.

**Changing a game's Proton version creates a new, empty prefix.** The old `compatdata/<appid>` contents stay on disk but the game no longer sees them, which reads exactly like "all my saves are gone". Back the prefix up first; to recover, switch back to the original build. The existing `compatdata/1076160.good` and `.proton11-bak-*` folders are prior manual backups of exactly this situation — leave them alone.

**`common/` contains folders for games that are not installed** (Cyberpunk 2077, Satisfactory, Sea Power, The Outlast Trials, Hunt Showdown, Rainbow Six Siege, Black Myth Wukong, Horizon Forbidden West, Last Drone War Demo). They are uninstall leftovers, not a library. Check `libraryfolders.vdf` or `INVENTORY.md` before assuming a game is playable.

**Verifying game files reverts every install-dir mod.** Steam updates do the same. Loose-file mods under `Content/Paks/~mods/` usually survive; edited base files never do.

**No swap on this machine.** Bulk extraction of large paks/bundles (Forza 159 GiB, theHunter 119 GiB, Palworld, WARNO) can trigger the OOM killer and take the desktop session with it. Extract in chunks, stream to disk, and keep an eye on `free -h`.

**A game that has never been launched has no save path yet.** Airport CEO and Hotline Miami have no prefix and no config dir — launch once, then use the `find -newermt` recipe above to discover where they actually write.

**Reminder — the six Blocked titles stay read-only.** Study their files all you like; write nothing, inject nothing, launch nothing modified. A ban is permanent and cannot be undone.
