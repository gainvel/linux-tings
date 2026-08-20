# Security Audit — AUR "Atomic Arch" Exposure + Full Vuln Sweep

**Date:** 2026-06-16
**Host:** `woody` (Artix Linux, OpenRC)
**Trigger:** User report of an AUR security vulnerability in the news; AUR is the primary
package manager (39 foreign packages).
**Analyst posture:** Authorized defensive audit. No package was removed, downgraded, or
modified. All findings below are read-only observations + proposed (un-applied) fixes.

---

## 1. The Incident — "Atomic Arch" AUR supply-chain campaign (June 2026)

Disclosed ~June 11, 2026. Attackers used AUR's **orphan-adoption** process to take over
~1,600 abandoned packages, rewrote their PKGBUILDs to pull malicious npm/bun dependencies,
and shipped a Rust infostealer + eBPF rootkit.

| Attribute | Detail |
|---|---|
| Malicious deps | `atomic-lockfile@1.4.2`, `lockfile-js` (npm, Wave 1); `js-digest` (bun, Wave 2) |
| Execution | `preinstall` lifecycle hook `./src/hooks/deps` runs at build/install time |
| Payload | Rust credential stealer (SHA256 `6144D433…`), ELF64 PIE |
| Steals | SSH keys; GitHub/npm/Discord/Slack/Teams/M365/Vault tokens; Docker creds; browser cookies |
| Persistence | systemd service, `Restart=always`, `/etc/systemd/system/` or `~/.config/systemd/user/` |
| Rootkit | eBPF maps `/sys/fs/bpf/hidden_{pids,names,inodes}` (needs CAP_BPF) |
| C2 | Tor onion `olrh4mibs62l6kkuvvjyc5lrercqg5tz543r4lsw3o6mh5qb7g7sneid.onion`, exfil via `temp.sh` |
| Staging binary | `/usr/bin/monero-wallet-gui` |
| Attacker accounts | AUR: `krisztinavarga`, `custodiatovar`, `veramagalhaes`; npm `herbsobering`; GitHub `fardewoak` |
| Unaffected | Official `[core]/[extra]/[multilib]` (Arch) — review-gated. AUR only. |

---

## 2. Exposure Verdict — NOT EXPOSED / NO INFECTION

| Check | Method | Result |
|---|---|---|
| 39 installed AUR pkgs vs. authoritative 1,600+ compromised list | cross-ref `lenucksi/aur-malware-check` package_list | **0 matches** |
| 8 pkgs built in attack window (Jun 9–12) — PKGBUILD/.SRCINFO IOC scan | grep IOC strings in `~/.cache/yay/<pkg>` | **all clean** |
| npm + bun caches for `atomic-lockfile`/`lockfile-js`/`js-digest` | grep `~/.npm` `~/.bun` | **0 hits** |
| eBPF rootkit maps `/sys/fs/bpf/hidden_*` | `ls /sys/fs/bpf` | **absent** |
| C2 staging binary `/usr/bin/monero-wallet-gui` | `test -e` | **absent** |
| Suspicious `/var/lib/<generated>` dirs | `find /var/lib -newermt 2026-06-08` | **none** |
| Tor process / C2 connections (ports 80/8080/9050) | `ss -tnp`, `pgrep tor` | **none** |
| Persistence (cron, autostart, /etc/local.d, systemd units) | enumerated | **all benign / pre-window** |
| Init-system applicability | — | **OpenRC** — the malware's systemd persistence cannot execute here |

**Packages built during the Jun 9–12 window (the only ones with any timing exposure):**
`claude-desktop-bin`, `discord-latest-bin`, `zen-browser-bin`, `visual-studio-code-bin`,
`yay`, `millennium`/`millennium-debug`, `linux-soundboard-git`.
All are high-traffic, actively-maintained packages — never orphaned, never adopted by the
attacker accounts, and all PKGBUILDs scanned clean. The campaign exclusively hijacked
*abandoned* packages, which is structurally why this system was missed.

**Conclusion:** No remediation required for this incident. Nothing to remove or clean.

---

## 3. Broader Vulnerability Sweep (the "while we're here" pass)

### 3a. Arch repo surface — Arch Security Tracker (`security.archlinux.org`)
`arch-audit` could not be installed in this environment (gitlab.com source blocked; crates.io
and the GitHub mirror both pin `alpm` crate versions incompatible with the system's
libalpm 16). Performed the **equivalent check directly** against the same upstream data
(`all.json`), matching installed package versions with `vercmp`.

**24 advisory matches — but ALL have `status=Vulnerable, fixed=-` (no upstream fix released).**
The system is fully current; `pacman -Syu` resolves none of these. Highest-severity, worth
tracking for when fixes land:

| Severity | Package | CVE(s) | AVG |
|---|---|---|---|
| High | `pam 1.7.2-2` | CVE-2025-6020 | AVG-2901 |
| High | `libxml2 2.15.3-1` | CVE-2025-6170, -49794/95/96 | AVG-2898 |
| High | `djvulibre 3.5.30-1` | CVE-2025-53367 | AVG-2907 |
| High | `grub 2:2.14-1` | CVE-2022-2873x (8 CVEs) | AVG-2762 |

Remainder are Low/Medium, mostly long-standing kernel/perl/openssl/libtiff hardening items
with no published fix. **No action available** — none are remediable by upgrade today.

### 3b. Lockfile / language-ecosystem surface — `osv-scanner` v2.3.8
Installed via official Google `go install` (GitHub), v2.3.8. Swept project trees under
`$HOME` (`~/projects`, `~/padloc-src`, `~/Repo`, `~/claude`).

**60+ npm findings — 100% are `(dev)` dependencies in local source checkouts:**
- `padloc-src/packages/tauri/package-lock.json` — build toolchain (webpack 5.52, vite,
  terser, loader-utils, node-forge, ws, etc.). Several high CVSS (webpack GHSA-hc6q… 9.8,
  loader-utils 9.8) but **dev-only, never shipped or run**.
- `projects/Ontairox-Parking-Detection/workbench/package-lock.json` — esbuild/vite/@babel
  dev deps.

Per CLAUDE.md triage (CVSS × **reachability**): these are unused build-time dependencies in
source checkouts, not installed/running software → **near-zero real exposure**. No runtime
dependency, no globally-installed package, and no Go/PyPI/cargo manifest flagged.

---

## 4. Recommendations (NONE applied — awaiting approval per operating principles)

1. **Atomic Arch:** nothing to do. Clean.
2. **Repo High-sev watch items** (`pam`, `libxml2`, `djvulibre`, `grub`): no fix exists yet.
   Run a normal `pacman -Syu` when Arch/Artix publishes patched versions; re-check the
   tracker. Optional: subscribe to `security.archlinux.org`.
3. **npm dev-dep findings:** low priority. If desired, in each project run
   `npm audit fix` (avoid `--force`, it bumps majors) to refresh the dev toolchain. These
   are your own project checkouts, so this is dev maintenance, not host remediation.
4. **Cleanup from this audit:** the failed `arch-audit` build left
   `~/.cache/yay/arch-audit-git/` (empty bare clone) and `/tmp/cargo-install*` artifacts —
   safe to delete. `osv-scanner` (~/go/bin, 111 MB) was kept intentionally as a scanner.

---

## 5. AUR Hygiene — reducing future exposure (AUR is your primary manager)

The Atomic Arch attack worked because AUR runs **arbitrary code as your user at build time**.
You dodged it by luck of package choice; harden the workflow so the next wave also misses:

- **Always review the PKGBUILD/.SRCINFO diff** at the `yay` prompt (don't `--noconfirm`
  blind on untrusted pkgs). Red flags: a *newly-changed maintainer* / recently-adopted
  package, `curl … | bash`, base64 blobs, sources off the real upstream host, npm/pip deps
  pulled inside `prepare()`/`build()`, `.install` hooks that fetch code.
- **Check the AUR page before installing/updating:** orphaned status, last-updated date, and
  whether maintainership just changed hands — that adoption event is this attack's entry point.
- **Periodic audit cadence** (e.g. monthly): re-run the IOC checks in §2, plus
  `osv-scanner scan source -r <project>` and the Arch Tracker match in §3a.
- **Shrink the blast radius of build-time code execution:** keep long-lived secrets
  (GitHub PAT, npm token, SSH keys) scoped and short-lived, so a future build-time stealer
  has less to take. Consider hardware-backed / passphrased SSH keys.

---

## Appendix — Tooling status
- `osv-scanner` v2.3.8 — **installed** (`~/go/bin/osv-scanner`, via official Google go module).
- `arch-audit` — **not installed** (env blockers above); equivalent Arch Tracker check used.
- IOC reference: `github.com/lenucksi/aur-malware-check`; `ioctl.fail/preliminary-analysis-of-aur-malware`.
