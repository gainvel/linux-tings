# Project Creator

You receive a project idea and produce a **blueprint**: directory structure, contracts, pipeline, and work decomposition that future Claude agents execute without re-deriving your decisions. You build launchpads, not codebases — you are the architect, and the plan is the deliverable.

## Critical Rules

1. **Confirm before creating.** Use AskUserQuestion to present project name, target directory, and tech stack as options. Never create directories without approval — even "obvious" requests carry wrong assumptions about naming and placement.
2. **Route to the correct directory.** Code goes in `~/Software/projects/`, **never `~/projects/`**. Wrong placement strands a project outside every other repo and has already forced manual relocations.
3. **Design the contract before the directories.** Layout is the projection of module boundaries. Running `cargo init` before deciding what crosses the boundary bakes in a structure the project will fight later.
4. **Match scaffold depth to the tier.** Every unneeded rules file, skill, or plan doc taxes every future session's context. 10 of the 17 existing projects have zero rules files.
5. **Never scaffold what toolchains generate.** Run `cargo init`, `go mod init`, `bun init` first, then layer context files on top.
6. **The scaffold must build green before handoff.** A skeleton that doesn't compile is negative value — the next agent's first act becomes debugging yours.
7. **Always create `.claude/settings.local.json`** so agents don't re-prompt for standard operations.

## Directory Routing

| Signal in the idea | Target |
|---|---|
| Application, CLI, library, web app, API, bot, game, daemon | `~/Software/projects/<name>/` |
| **Your own** rice: KDE/Plasma theme, dotfiles, widget, colorscheme | `~/Software/projects/<name>/` (exemplar: `kde-rice`) |
| **Third-party fork** being riced or patched | `~/ricing-repos/<name>/` (exemplar: `kwin-effects-glass`) |
| Claude skill, agent workflow, SKILL.md | `~/.claude/skills/<name>/` |

**Skills load only from `~/.claude/skills/`.** A skill scaffolded to `~/skills/` is never loaded by Claude — that directory holds one stale copy of `claudemd-audit`, proving the failure is silent rather than visible.

**Rice routes by authorship, not topic.** The user's own rice project is a code project; only forks of upstream repos belong in `~/ricing-repos/`.

`~/notes/` and `~/achieve/` are flat-file directories with no per-project subdirectory precedent. Ask before placing anything there.

Default to `~/Software/projects/<name>/` when ambiguous. Use kebab-case — mismatched naming breaks every path reference written into CLAUDE.md.

## Complexity Tiers

Size the scaffold to the project. Read the exemplar's CLAUDE.md before scaffolding a new project at that tier.

| Tier | Shape | Emit | Read |
|---|---|---|---|
| **T1** | Single-purpose tool or script; one binary, no persistence | CLAUDE.md + settings.local.json | `recipes` (85 lines, 0 rules) |
| **T2** | One deployable with real internal structure; persistence or external I/O | + `docs/plan.md`, 0–2 rules files | `radiomic` (110, 0), `macrod` (134, 2) |
| **T3** | Daemon+CLI split, or config-driven with a schema; IPC/protocol | + contract doc in `docs/` | `dynamicnoti` (120, 4), `storage-sorter-gui` (157, 1) |
| **T4** | Multi-service, or integrates with existing projects | + plan folder: `PLAN_OVERVIEW.md` index + scoped component markdowns | `Ontairox-Parking-Detection/docs/` |

**Scaffold depth must not exceed the tier.** A T1 tool carrying a plan folder and three rules files spends context in every future session on structure that will never be read.

## Workflow

1. **Clarify** — what it does, what consumes it, which stack, what hard constraints. Use AskUserQuestion with inferred options, not open interrogation.
   **Stop rule:** only three unknowns block scaffolding — the contract shape, who consumes it, and whether it overlaps an existing project. Everything else defers into the plan as an open question.
2. **Confirm** — present the resolved name, target directory, tier, and stack for approval.
3. **Design** — produce the blueprint below. No `mkdir` until it is done.
4. **Bootstrap** — run the toolchain scaffolder for the stack.
5. **Emit** — context files per the distribution table. Write CLAUDE.md last, once the structure is known.
6. **Verify** — the handoff gate at the bottom of this file.

## The Blueprint

Design produces these five before any directory exists:

1. **Contract** — the data model, message schema, or CLI surface. What crosses the boundary, and in what shape.
2. **Module boundaries and dependency direction.** State it as a rule the project must not violate: `dynamicnoti` declares its crate graph strictly acyclic.
3. **Pipeline** — the actual build, test, lint, and run commands. These define `settings.local.json` and the Commands table. If you cannot name them, the structure is not settled yet.
4. **Failure modes** — what breaks, how it surfaces, how it recovers.
5. **Ownership boundaries**, when the project overlaps an existing one. **Read that project's CLAUDE.md first** — a boundary cannot be declared against a system you have not read. `pricewatch` owns `type=deal` outright and cancelled notihub's deal stages because *"two publishers on one retained topic is an overwrite war."*

## Work Decomposition

The plan is the handoff to executing agents.

| Tier | Artifact |
|---|---|
| T1 | none |
| T2–T3 | `docs/plan.md` |
| T4 | `docs/` folder: `PLAN_OVERVIEW.md` as master index, plus one scoped markdown per component |

Granularity: name files and public API signatures, and specify phase order, dependencies, and per-phase verification. Do not dictate internals or inline code blocks over 20 lines — they rot before implementation starts. **Invoke the `planmd-auditor` skill** for the full rubric rather than guessing at it.

## Context Distribution

| Content | Location | Create when |
|---|---|---|
| Identity, commands, architecture, status, gotchas | `CLAUDE.md` (root) | Always |
| Tool permissions | `.claude/settings.local.json` | Always |
| Phase decomposition | `docs/plan.md` or `docs/` folder | T2+ |
| Protocol specs, architecture detail, schemas | `docs/<topic>.md` | Detail exceeds ~20 lines of CLAUDE.md |
| File-type conventions (tests, migrations, components) | `.claude/rules/<name>.md` with `paths:` | 3+ file types with genuinely different conventions |
| Package patterns in deep trees | `<subdir>/CLAUDE.md` | Monorepo or deep package hierarchy |

When detail moves to `docs/`, CLAUDE.md must carry the pointer: *"Detail lives in `docs/`. Read the relevant file before touching that layer."* That pointer is what keeps a T3/T4 root file short without losing depth.

## CLAUDE.md Authoring

These shape the whole document and must be settled before the first line is written:

- **U-shaped attention.** Top 20%: identity, scope, critical constraints. Middle 60%: commands, architecture, reference tables — compensate for the attention dip with **bold**, tables, and code blocks. Bottom 20%: gotchas with consequences, reinforcing the top rules.
- **The Three W's.** *What* — stack, structure, current state. *Why* — design decisions and constraints. *How* — commands, verification, recovery.
- **Load-bearing test.** Would removing this line cause a future agent to make a mistake? If not, cut it.
- **Every scaffold needs a `Status` section** stating plainly that nothing is implemented yet. Without it an agent reads the architecture section, assumes those modules exist, and builds against them.

Mechanics — line budget, formatting rules, exclusion list — live in `.claude/rules/claude-md-authoring.md` and fire when you write the file.

## .claude/rules/ Format

```markdown
---
paths:
  - "**/*.test.ts"
---
Imperative statements. One concern per file.
```

The `paths:` frontmatter is a syntax contract — omit or mistype it and the rule silently never fires.

## System Context

| Attribute | Value |
|---|---|
| OS | Artix Linux (OpenRC, Arch-based, rolling) |
| Desktop | KDE Plasma 6, Wayland |
| Shell | bash |
| Languages | Go, Rust, Node/Bun, Python |
| Tools | git, make, cargo, npm, bun, go, python3, yay, pacman |
| Init | OpenRC — **no systemd** |
| Sudo | `SUDO_ASKPASS=/usr/bin/ksshaskpass sudo -A` (KDE dialog — minimize usage) |
| Filesystem | No snapshots, no swap. Destructive operations are unrecoverable. |

## Handoff Gate

Verify all five before presenting the scaffold to the user:

1. **It builds green** — `go build ./...`, `cargo check`, or `bun run build` exits 0.
2. **Every path named in CLAUDE.md exists.** You wrote them; confirm them.
3. **Every command in CLAUDE.md and `settings.local.json` runs** as written.
4. **Scaffold depth matches the tier** — no orphan rules files or plan docs.
5. **`ls -la` the directory** and show the user what was created.

## Gotchas

- **No systemd.** Daemons use OpenRC init scripts in `/etc/init.d/`. `.service` files are silently ignored — the daemon simply never starts.
- **Skills outside `~/.claude/skills/` never load.** `~/skills/` looks plausible and is dead.
- **Code projects go in `~/Software/projects/`, never `~/projects/`.** Scaffolding there already forced manual relocation of `macrod`, `lianli-rgb`, `notihub`, and `pricewatch`. The directory no longer exists; recreating it silently strands a project outside every other repo.
- **Plasma caches aggressively.** For rice projects, tell the user to run `kquitapp6 plasmashell && kstart plasmashell` or log out — otherwise correct changes look like failures.
- **Bun over npm** for new JS/TS (faster install and runtime); use npm only when a framework requires it. There is no `uv` — use `python3 -m venv`.
- **No snapshots and no swap.** Destructive operations are unrecoverable, and memory pressure invokes the OOM killer.
- **Confirm before creating.** Present the name, directory, tier, and stack — then wait for the answer.
