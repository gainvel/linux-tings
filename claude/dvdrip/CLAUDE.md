# DVD Ripping → Jellyfin (cmp01)

Everything here concerns **cmp01**, the Jellyfin node at `192.168.1.78` with the DVD drive
plugged into it: ripping discs, encoding, library layout, and the Jellyfin server itself.

**cmp01 is Debian 13 — systemd, apt. The desktop this session runs on is Artix — OpenRC,
pacman.** Commands for one will not work on the other, and the DVD drive is in cmp01.

---

## Critical rules

1. **Probe with `lsdvd -a -s /dev/sr0` before naming any title or track number.** Every
   disc numbers its titles, audio and subtitles differently — a guessed `-t`/`-a`/`-s`
   yields a silently wrong rip discovered 35 minutes later.
2. **Match filters to the source cadence.** NTSC film needs `--detelecine`; PAL does not.
   Detelecining progressive PAL discards real frames.
3. **Write only under `/srv/media/movies` or `/srv/media/shows`, with the exact naming
   scheme below.** Those are the only two paths Jellyfin scans, and its matcher keys off
   `Title (Year)`.
4. **Rip as `node-admin` with no sudo.** `node-admin` is in `cdrom` and owns `/srv/media`.
   Jellyfin's config under `/etc/jellyfin/` is the one thing needing root.
5. **Detach anything outliving a shell** — `tmux new -d -s rip '…'` or `nohup … &`. A
   feature takes ~35 min; a foreground encode dies with the SSH connection.
6. **Confirm title, year and destination before starting an encode.** 35 minutes into the
   wrong path is 35 minutes lost, and Jellyfin will not match it.

---

## The box

| | |
|---|---|
| Host | `cmp01` — `192.168.1.78`, user `node-admin` (sudo needs a password) |
| OS | Debian 13 trixie, kernel 6.12 — `apt`, `systemctl`, `journalctl` |
| CPU | i5-12500T — 6C/12T Alder Lake, 35 W |
| RAM | 7.4 GiB + 7.7 GiB swap — **tight**; Jellyfin's heap grows to ~5 GiB over days of uptime |
| iGPU | Intel UHD 770 → `/dev/dri/renderD128`, VAAPI 1.22 / iHD 25.2.3 |
| Optical | USB PLDS `DVD+-RW DS-8A4S` → `/dev/sr0` (`/dev/cdrom` symlink) |
| Media disk | 1 TB WD_BLACK NVMe → `/srv/media`, `node-admin:media` setgid, ~900 GB free |
| Installed | `HandBrakeCLI` 1.9.2, `lsdvd`, `eject`, `tmux`, `libdvdcss2`, `curl`, `python3` |
| Not installed | `ffmpeg`, `ffprobe`, `mediainfo`, `mkvtoolnix`, `makemkv`, `jq`, `git`, `vim` |

`/usr/lib/jellyfin-ffmpeg/` carries its own `ffmpeg` and `ffprobe` — use those rather than
installing anything.

## Reaching cmp01

`ssh cmp01` fails with *Host key verification failed*: the name resolves through the
`localdomain` search suffix, but `known_hosts` only holds the IP. **Use `192.168.1.78`.**

The desktop key is not in cmp01's `authorized_keys` and `~/.ssh/id_ed25519` is
passphrase-protected, so password auth is the working path today:

```bash
export SSH_ASKPASS=/usr/bin/ksshaskpass SSH_ASKPASS_REQUIRE=force DISPLAY=:0
setsid -w ssh -M -S /tmp/claude-1000/c1.sock -o ControlPersist=4h \
  -o PreferredAuthentications=password -o NumberOfPasswordPrompts=1 \
  -f -N node-admin@192.168.1.78                                   # one GUI prompt
ssh -S /tmp/claude-1000/c1.sock node-admin@192.168.1.78 '<command>'   # free thereafter
```

Keep the control socket path short — a session scratchpad path overflows the 108-byte
`sun_path` limit. For remote sudo, capture the password with `ksshaskpass` and pipe it into
`sudo -S -p "" bash -c "…"`, batching all privileged work into a single call.

Once the pubkey is installed on cmp01, unlock the passphrase the same way:

```bash
eval "$(ssh-agent -s)"
SSH_ASKPASS=/usr/bin/ksshaskpass SSH_ASKPASS_REQUIRE=force ssh-add ~/.ssh/id_ed25519
```

---

## Reading a disc

```bash
lsdvd /dev/sr0            # titles and runtimes
lsdvd -a -s /dev/sr0      # plus audio and subtitle streams
```

**Picking the title.** `Longest track:` at the foot of the output is the feature on a movie
disc. On a TV disc, episodes are the titles inside a runtime window (`~/ripdisc.sh` uses
20–95 min).

| What you see | What it means |
|---|---|
| Dozens of 1–2 second titles | Normal menu/transition padding — ignore |
| Feature split across two long titles | Seamless branching; rip the longer, or rip both and join |
| Every title a near-identical runtime | Playlist obfuscation; pick by chapter count — the real one has the most |
| One title, ~2 h, 40+ chapters | The feature |

**Mapping streams to flags.** `lsdvd`'s audio and subtitle numbers map 1:1 onto `-a` / `-s`:

| `lsdvd` line | Read as |
|---|---|
| `Audio: 1 … Channels: 6` | The 5.1 main mix — nearly always `-a 1` |
| Further `en … Channels: 2` audio | Commentary; name them `Commentary 1`, `Commentary 2`, … |
| `Subtitle: NN … Content: Normal_CC` | Closed captions → `English CC` |
| A second `en … Content: Normal` subtitle | Usually forced/signs → `English Forced` |
| `es` / `fr` tracks | Skipped — this library keeps English only |

Established convention: English 5.1 plus every English commentary; English subtitles
present but off by default.

---

## Encoding

**Video.** NTSC DVDs are 720×480 @ 29.97 fps carrying 23.976 fps film via 3:2 pulldown →
`--detelecine --comb-detect --decomb`. PAL is 720×576 @ 25 fps and usually progressive →
`--comb-detect --decomb`, no detelecine. DVDs are anamorphic — 720-wide frames with a 4:3
or 16:9 display flag — so let HandBrake carry the aspect instead of forcing a resolution.

**Audio.** `-E copy` passes AC3 through untouched; `--audio-fallback av_aac` covers
anything that cannot be copied.

**Subtitles.** DVD subtitles are VOBSUB *bitmaps*, not text — they cannot become SRT
without OCR, and no OCR tool is installed here. `--subtitle-default=none` keeps them
selectable but off.

**Quality.** `-q` is the x265 RF scale, lower being better. DVD source lives in 18–22;
`-q 20 --encoder-preset medium` is the established default. Expect ~1.6 GB and ~35 min for
a 2-hour feature, ~700 MB and ~20 min for a 50-minute episode.

| | CPU x265 (HandBrake) | Hardware HEVC (jellyfin-ffmpeg) |
|---|---|---|
| Command | `HandBrakeCLI -e x265 -q 20 --encoder-preset medium` | `/usr/lib/jellyfin-ffmpeg/ffmpeg -hwaccel vaapi -hwaccel_output_format vaapi -i … -c:v hevc_vaapi` |
| Speed | ~3.5× realtime | Several × faster, effectively disc-read-bound |
| Quality | Best per bit | Larger files for matched quality |
| Pick for | Keepers, anything watched properly | Bulk box sets where throughput wins |

**`HandBrakeCLI` cannot hardware-encode on this box.** Debian's build prints `qsv: not
available on this system` and `Cannot load libnvidia-encode.so.1` at every launch — normal
noise, not a fault to chase. Hardware encoding exists only through
`/usr/lib/jellyfin-ffmpeg/ffmpeg`.

**Copy protection.** `libdvdcss2` (via `libdvd-pkg`) decrypts CSS transparently, caching
keys in `~/.dvdcss/`. A read failure is far more often a dirty or scratched disc than a
missing library. The drive is RPC-2 with a finite region-change counter, so a
region-mismatched disc is a real failure mode.

### Movie

```bash
mkdir -p "/srv/media/movies/TITLE (YEAR)"
HandBrakeCLI -i /dev/sr0 -t 1 \
  -o "/srv/media/movies/TITLE (YEAR)/TITLE (YEAR).mkv" \
  -f av_mkv -e x265 -q 20 --encoder-preset medium \
  --detelecine --comb-detect --decomb \
  -a 1,4 -E copy,copy --audio-fallback av_aac \
  --aname "English 5.1,Commentary" \
  -s 1,2,5 --subtitle-default=none \
  --subname "English CC,English,English Forced" \
  --markers < /dev/null
```

Title, audio and subtitle numbers are **per-disc** — take them from `lsdvd -a -s`, never
from this example. `< /dev/null` stops HandBrake consuming stdin when it is not on a
terminal. `--markers` writes chapter markers.

### TV episode

The same, without commentary, into a season folder:

```bash
-a 1 -E copy --audio-fallback av_aac      # no --aname/--subname needed
-o "/srv/media/shows/SHOW (YEAR)/Season 01/SHOW (YEAR) - S01E01.mkv"
```

`~/ripdisc.sh` on cmp01 batch-rips a season across sequential discs, auto-numbering
episodes from what already exists in the destination. Its `SHOW`/`SEASON`/`DEST_ROOT` are
hardcoded (currently Sopranos S1) — edit the variables rather than writing a new script.

### Running it

```bash
tmux new -d -s rip 'HandBrakeCLI … 2>&1 | tee /tmp/rip.log'
tail -f /tmp/rip.log                      # live progress — the one to hand the user
tmux capture-pane -p -t rip | tail -3     # one-shot progress sample
tmux has-session -t rip && pgrep -a HandBrakeCLI    # is it actually still running?
eject /dev/sr0                            # when the disc is done
```

**Prefer `tail -f /tmp/rip.log` over `tmux attach` for watching an encode.** Both show the
same live output, but Ctrl+C in `tail` stops only the tail, whereas **Ctrl+C in an attached
tmux pane sends SIGINT to HandBrake and kills the encode** — the exit from an attached pane is
`Ctrl+b` then `d`. HandBrake writes to `tee` rather than a terminal, so it emits progress as
periodic newline-terminated lines instead of a rewriting `\r` counter, which is what makes
`tail -f` readable. Re-running the `tmux new` line after a killed encode needs no cleanup;
HandBrake overwrites the output file.

---

## Jellyfin

`http://192.168.1.78:8096` · `jellyfin.service` · v10.11.11

| | |
|---|---|
| Runs as | `jellyfin` (uid 102) in `video`, `render`, `media` |
| Binary / encoder | `/usr/bin/jellyfin` · `/usr/lib/jellyfin-ffmpeg/ffmpeg` |
| Config | `/etc/jellyfin/` — root-only (`encoding.xml`, `network.xml`, `system.xml`) |
| Data & metadata | `/var/lib/jellyfin/` |
| Logs | `journalctl -u jellyfin -b` |
| Libraries | `/srv/media/movies` (Movies), `/srv/media/shows` (Shows) |
| Hardware accel | VAAPI on `/dev/dri/renderD128`, HEVC encode enabled |

**Naming drives matching** — mirror what is already on disk:

```
/srv/media/movies/Blade Runner (1982)/Blade Runner (1982).mkv
/srv/media/movies/Star Wars Episode IV - A New Hope (1977)/Star Wars Episode IV - A New Hope (1977).mkv
/srv/media/shows/The Sopranos (1999)/Season 01/The Sopranos (1999) - S01E01.mkv
```

**Films in a series take the full `FRANCHISE Episode N - SUBTITLE (YEAR)` form**, matching the
existing Episode I and II folders — `Star Wars Episode IV - A New Hope (1977)`.
**`ls /srv/media/movies` is the authority before naming any new rip.**

**Folder names only drive matching — they do not drive display.** Once an item matches, the
Movies view shows TMDB's title, poster and year, not the filename. TMDB's canonical title for
the 1977 film is bare `Star Wars`, so a correctly-matched Episode IV displays as `Star Wars`
alongside `Star Wars: Episode I - The Phantom Menace`. That is a metadata result, not a
mismatch. To override it, edit the item in the UI (Name and Sort Title) and **lock the item**,
or the next library scan reverts it. Sort Titles of the `Star Wars Episode 1` / `… 4` form put
the saga in **episode order** (I, II, IV …) regardless of each film's canonical title. Note that
this is not release order — Episode IV is the 1977 original and the prequels came 1999–2005 — so
for release order, sort the library by Release Date instead of hand-writing Sort Titles.

Alternate cuts: `Blade Runner (1982) {edition-Final Cut}.mkv`. Multi-disc features:
`… - pt1.mkv` / `… - pt2.mkv` in one folder. Show extras live in a `Specials` folder
(season 0).

**Playback ladder:** direct play → direct stream (remux only) → transcode. x265 rips are
small but push older clients into transcoding, which is why VAAPI matters on a 35 W chip.

| Symptom | Cause | Action |
|---|---|---|
| New rip missing from the library | No scan yet | Dashboard → Scan All Libraries, or `POST /Library/Refresh` with an API key |
| Matched to the wrong film | Ambiguous `Title (Year)` | Identify manually in the UI, then lock the item |
| Present but will not play | Permissions | `jellyfin` reads through the `media` group — check group and mode on the file |
| Stutters, CPU pegged | Transcoding | Dashboard → active streams shows the transcode reason |
| Will not start | Config or port | `systemctl status jellyfin`, `journalctl -u jellyfin -b` |

---

## Gotchas

- **`ssh cmp01` fails host key verification** — always `192.168.1.78`.
- **7.4 GiB RAM, and Jellyfin's heap climbs toward 5 GiB the longer it runs.** Check
  `free -h` before encoding in parallel — the OOM killer takes the encode, not the server.
- **No `ffmpeg`/`ffprobe`/`mediainfo` on `PATH`.** Inspect finished files with
  `/usr/lib/jellyfin-ffmpeg/ffprobe`.
- **`sudo` on cmp01 prompts for a password** — batch privileged work into one
  `sudo bash -c '…'`.
- **Jellyfin rewrites `/etc/jellyfin/*.xml` on shutdown**, so stop the service before
  hand-editing config or the change is lost.
- **Eject when a disc finishes** (`eject /dev/sr0`) — discs are swapped by hand, and a
  locked tray reads as a hung rip.
- **Probe before prescribing; detach long encodes.**
