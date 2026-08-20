# ThinkCentre / ThinkStation Tiny — 1 L Node Reference

*Companion document to `nodes.md`. Scope: **Lenovo Think-branded 1-litre "Tiny" and sub-1 L "Nano"
desktops only** — the boxes that actually fit a 10-inch rack. Built 2026-07-24.*

`nodes.md` carries three Lenovo rows (M720q, M920q, M75q Gen 2) and an **Appendix C** processor catalog.
This document is the expansion: **every** Tiny/Nano model Lenovo has shipped, every silicon variant found
in them, what each is physically capable of in a rack, and what jobs each is and isn't suited to.

---

## Contents

- [How to read this document](#how-to-read-this-document)
- [Scope — what counts, what's excluded](#scope--what-counts-whats-excluded)
- [The generation decoder — the biggest buyer traps](#the-generation-decoder--the-biggest-buyer-traps)
- [Chassis platform eras — what physically changed](#chassis-platform-eras--what-physically-changed)
- **[Part 1 — Device roster](#part-1--device-roster)**
  - [1.0 At-a-glance — every model, one row](#10-at-a-glance--every-model-one-row)
  - [1.1 Legacy Tiny (2012–2016)](#11-legacy-tiny-20122016)
  - [1.2 Intel Tiny4 / Tiny5 — the classic homelab boxes](#12-intel-tiny4--tiny5--the-classic-homelab-boxes)
  - [1.3 AMD Tiny, first wave](#13-amd-tiny-first-wave)
  - [1.4 Modern Intel Tiny — M70q / M80q / M90q](#14-modern-intel-tiny--m70q--m80q--m90q)
  - [1.5 Modern AMD Tiny — M75q Gen 5](#15-modern-amd-tiny--m75q-gen-5)
  - [1.6 Budget / mobile-silicon Tiny — neo 50q](#16-budget--mobile-silicon-tiny--neo-50q)
  - [1.7 Nano — sub-1 L](#17-nano--sub-1-l)
  - [1.8 ThinkStation Tiny — the GPU-capable 1 L](#18-thinkstation-tiny--the-gpu-capable-1-l)
  - [1.9 Excluded and warning rows](#19-excluded-and-warning-rows)
  - [**1.10 Used-market pricing and availability**](#110-used-market-pricing-and-availability)
- **[Part 2 — The chip table](#part-2--the-chip-table)**
- **[Part 3 — 10-inch rack fit, mounting and hardware](#part-3--10-inch-rack-fit-mounting-and-hardware)**
- **[Part 4 — Task ↔ silicon suitability](#part-4--task--silicon-suitability)**
- **[Part 5 — Workload catalog](#part-5--workload-catalog)** — 398 tasks in 15 domains, task-first and hardware-neutral; the resource-tag vocabulary and the scoping rule
- [Known gaps](#known-gaps--what-this-document-does-not-have)
- [Sources](#sources)

---

## How to read this document

Same tag convention as `nodes.md`, so the two can be read together:

| Tag | Meaning |
|---|---|
| **M** | Measured by an independent third party (tester named) |
| **V** | Vendor claim — Lenovo PSREF, Intel ARK, AMD product page. Not a measurement. |
| **`~`** | Estimate / unverified |
| **`no data`** | Nothing found. **"Not found" ≠ "does not exist."** Look again before concluding. |
| (AC) | Measured at the wall — includes adapter loss. This is what you budget a PDU against. |
| [X] | Could not be verified to exist |

**Three provenance warnings specific to this document:**

1. **GB6 scores are aggregator medians, not primary reads.** `browser.geekbench.com` was not read this
   session. Figures marked **M** come from cpu-monkey / nanoreview / cputronic / chipversus, which mirror
   the same Geekbench corpus. Treat as **±5–10 %**. `*` marks a multicore median that is implausibly low
   (tiny, throttled sample pool) and should not be used for ranking.
2. **TDP is a package limit, not wall power.** A 35 W `T` chip in a 1 L box does not draw 35 W at the wall,
   and a 65 W chip in the same box does not sustain 65 W. Use the **measured (AC)** figures in the device
   rows, which come mostly from ServeTheHome's *Project TinyMiniMicro* series.
3. **Watts are per-machine, not per-CPU.** Nobody publishes per-CPU wall power for these boxes. A per-CPU
   watt column would be `no data` end-to-end, so it is omitted from the chip table on purpose.

**The `**Good at:**` / `**Weak at:**` convention in Part 1.** Every model section carries a short pair of
these blocks. Per Part 5's scoping rule, `**Good at:**` lists only what that machine's hardware genuinely
decides — the things it does better than the box in the next section — not everything it is capable of.
Nearly every machine in this document can run nearly anything in the Part 5 catalog, so a long list carries
no information. Where a model is simply a competent generalist, that is what `**Good at:**` should say, in
one line, with the space spent on `**Weak at:**` instead.

---

## Scope — what counts, what's excluded

**In scope:** the 1 L "Tiny" chassis (`q`-suffix models, plus the pre-`q` M700/M900 and the 2012–2015
Tiny boxes), the sub-1 L "Nano" (`n`), and the ThinkStation `P…Tiny` workstation variants of the same
chassis. All are ~179 × 183 mm and 34.5–37 mm tall — comfortably inside a 10" rack's ~222 mm usable width
at 1U (44.45 mm).

**Excluded, and why:**

| Excluded | Why |
|---|---|
| **SFF (`s`) and Tower** | 8–18 L. Do not fit a 10" rack in any orientation. |
| **`x` variants — M910x, M920x** | ~1.35 L, **taller** chassis (approx. 45 mm vs 34.5 mm) to fit a second M.2 + GPU. They *fit the width* but need 2U and defeat the density argument. Kept as warning rows in §1.9. |
| **ThinkStation P3 Ultra / P3 Ultra SFF** | ~3.6 L. Genuinely excellent hardware, genuinely does not fit. §1.9. |
| **Tiny-in-One (TIO) monitors** | The Tiny is the compute; the TIO is a display shell. Irrelevant to a rack. |
| **ThinkEdge SE10 / SE30** | Adjacent Think-branded small boxes, but DIN-rail/edge industrial form factors on Atom/Core-U silicon, not the 1 L Tiny chassis. Noted in §1.9 for completeness. |

**Cost basis — used market only.** Every price in this document is a **used / refurbished street price**.
New and MSRP pricing is deliberately absent: it is not a purchase channel under consideration, so quoting
it would only distort the comparisons. The practical consequence is stated up front because it reshapes
the whole recommendation set: **the newest and most capable models in Part 1 are not currently buyable.**
M90q Gen 6, P3 Tiny Gen 2, neo 50q Gen 5 and M70q Gen 5 are all still in their first corporate lease
cycle. They are documented here so the roster is complete and so you know what to watch for in 2027–29 —
not as things to buy today. See **§1.10** for availability tiers and prices.

**One nuance that trips people up:** the **M90q** is the only 1 L Tiny Lenovo officially qualifies for
**65 W non-`T` desktop chips** (with the performance heatsink). They physically fit. They also power-throttle.
Those SKUs are included and flagged. The same applies to the AMD `G` (65 W) parts on the M75q Gen 5.

---

## The generation decoder — the biggest buyer traps

**Trap 1: Lenovo "Gen N" ≠ Intel "Nth generation."** Two generations ship two microarchitectures at once.

| Lenovo model | Intel gen / µarch | Notes |
|---|---|---|
| M700 / M900 Tiny | 6th — Skylake | `T` only |
| M710q / M910q | 7th — Kaby Lake (+ some 6th carryover) | `T` only |
| M720q / M920q | 8th Coffee Lake **and** 9th Coffee Lake-R | Identical CPU menu; M920q differs by **Q370 vPro** chipset vs M720q's **B360** |
| M70q **Gen 1** | 10th — Comet Lake | `T` only |
| M70q **Gen 2** | **SPLIT:** Celeron/Pentium/i3 + i5-104/105xxT = **10th** Comet Lake; i5-11xxx/i7/i9 = **11th** Rocket Lake | dual-µarch |
| M70q **Gen 3** | 12th — Alder Lake | first hybrid P+E in a Tiny |
| M70q **Gen 4** | 13th — Raptor Lake | no i9 (tops at i7-13700T) |
| M70q **Gen 5** | **13th + 14th** Raptor Lake / Raptor Lake-R, plus **Intel 300T** | DDR5, Q670, **no PCIe riser** |
| M80q **(Gen 1)** | 10th — Comet Lake | PSREF lists it as plain "M80q" |
| M80q **Gen 2** | **DOES NOT EXIST** | Lenovo skipped the label |
| M80q **Gen 3 / Gen 4** | 12th / 13th | `T` only |
| M90q **(Gen 1)** | 10th — Comet Lake | `T` **and** 65 W non-`T` |
| M90q **Gen 2** | **SPLIT:** low-end **10th** Comet Lake + i5/i7/i9 **11th** Rocket Lake | dual-µarch |
| M90q **Gen 3 / Gen 4** | 12th / 13th | `T` and 65 W non-`T` |
| M90q **Gen 5** | **SPLIT: 13th + 14th** Raptor Lake / Raptor Lake-R | all UHD 770 — **not** Core Ultra |
| M90q **Gen 6** | **Core Ultra Series 2** — Arrow Lake-S, LGA1851, Q870 | the first genuinely new platform in years |
| M75q **Gen 1 / Gen 2 / Gen 5** | Picasso / (Renoir + Cezanne) / Phoenix (Zen 4) | there is **no** M75q Gen 3 or Gen 4 |
| neo 50q **Gen 4** | 13th-gen **H-series mobile** (BGA) | plus a separate *Thin Client* variant on Celeron/12th i3 |
| neo 50q **Gen 5** | **Raptor Lake-H rebrand** (Core 5 210H / Core 7 240H) | see Trap 3 |

**Trap 2: "Gen 5" means different silicon depending on the model.** M70q Gen 5 = 13/14th gen desktop `T`.
M75q Gen 5 = Zen 4 Ryzen 8000. M90q Gen 5 = 13/14th gen. M90q **Gen 6** = Core Ultra. These are not aligned.

**Trap 3: "Core 5 / Core 7 200H" is not Core Ultra and is not new silicon.** The neo 50q Gen 5's
<cite index="140-1">Core 7 240H is Raptor Lake-H — 6 P-cores + 4 E-cores, 10C/16T, 2.5 GHz base, 5.2 GHz boost, 24 MB L3, 64-EU integrated graphics, 45 W base TDP with 115 W max turbo</cite>. Retailers list it as
"Series 2," which invites confusion with Core Ultra Series 2 (Arrow Lake). It has **no NPU**. Buy it for
cores-per-dollar, not for AI features.

**Trap 4: the PCIe riser is not universal.** In the 10th-gen-and-later era, **only the M90q has a PCIe
expansion slot — the M70q and M80q do not.** ServeTheHome flagged this explicitly during the M90q Gen 1
review. If you want a 10GbE NIC or a quad-port NIC in a modern Tiny, the model letter matters more than
the generation number. (Full riser detail in §3.4.)

**Trap 5: Lenovo `T` and AMD `GE` are the rack parts.** `T` = 35 W Intel low-power desktop; `GE` = 35 W AMD
low-power APU. The 65 W `non-T` / `G` chips exist in M90q and M75q Gen 5 SKUs, fit physically, and throttle
under sustained all-core load. Buy them for bursty single-thread work, never for a sustained-MT rack role.

---

## Chassis platform eras — what physically changed

Lenovo's internal chassis naming ("Tiny5", "Tiny6") is what riser and accessory sellers use, so it is worth
knowing. All eras share the same ~179 × 183 mm footprint; **height and internal layout are what changed.**

| Era | Models | H (mm) | Storage layout | PCIe riser | Notes for a rack |
|---|---|---|---|---|---|
| Tiny1–2 (2012–14) | M72e, M73, M83, M93/M93p | ~34.5 | 2.5" bay + mSATA/M.2 (late) | none | DDR3, USB 2.0-heavy, dead-end |
| Tiny3 (2015–16) | M53, M600, M700, M900 | ~34.5 | 2.5" bay + M.2 | none | M900's **rubber feet are rotated 90°** vs everything else — breaks generic mounts |
| Tiny4 (2017) | M710q, M910q, M910x, M625q, M715q, P320 Tiny | 34.5 (x = ~45) | 2.5" bay + 1× M.2 | M910x / P320 only | first Tiny with a usable NVMe slot |
| **Tiny5 (2018–19)** | **M720q, M920q**, M920x, P330 Tiny, M75q-1 | 34.5 (x = ~45) | 2.5" bay + 1× M.2 | **x8 via 01AJ940** on Intel | **the enthusiast sweet spot** — cheap, riser-capable, huge used supply |
| Tiny6 (2020–22) | M70q/M80q/M90q Gen 1–2, M75q Gen 2, P340/P350 Tiny | 34.5–37 | 2.5" bay + 1–2× M.2 | **M90q and P-series only** | riser disappears from the mid-tier |
| Tiny7 (2022–24) | M70q/M80q/M90q Gen 3–5, P360/P3 Tiny, neo 50q | 35–37 | 1–2× M.2, bay dropped on most | M90q / P-series only | DDR5, PCIe 4.0, 2.5GbE punch-out appears |
| Tiny8 (2025–26) | **M90q Gen 6**, **P3 Tiny Gen 2**, M75q Gen 5 | 36.5–37 | **3× M.2** (P3 G2: one is Gen 5) | x8 / x16(x8) riser returns | LGA1851, Q870, discrete GPU option returns |

**The single most consequential trend:** the 2.5" SATA bay is gone from modern Tinies, and NVMe slots
replaced it. For a NAS-style role that is a downgrade (no cheap bulk SATA); for everything else it is an
upgrade (2–3 NVMe on real PCIe 4.0 x4 links).

---

# Part 1 — Device roster

## 1.0 At-a-glance — every model, one row

Idle figures are **(AC)**, measured at the wall on 120 V, mostly by ServeTheHome. `~` = vendor/estimated.

| Model | Years | Silicon | Top SKU | RAM max | Storage | PCIe riser | Stock NIC | Idle W | Verdict for a 10" rack |
|---|---|---|---|---|---|---|---|---|---|
| M72e Tiny | 2012–13 | Ivy Bridge `T` | i5-3470T `~` | 16 GB DDR3 | 2.5" + mSATA | no | GbE | `no data` | Museum piece. Skip. |
| M73 Tiny | 2014 | Haswell `T` | i7-4765T `~` | 16 GB DDR3L | 2.5" + M.2 | no | GbE (I217) | `no data` | Free-only. |
| M83 Tiny | 2014 | Haswell `T` | i7-4785T `~` | 16 GB DDR3L | 2.5" + M.2 | no | GbE + vPro | `no data` | Free-only; has AMT. |
| M93 / M93p Tiny | 2013–14 | Haswell `T` | i7-4785T `~` | 16 GB DDR3L | 2.5" + mSATA | no | GbE + vPro | `~10–14` | Best of the legacy tier; still slow. |
| M53 Tiny | 2015 | Bay Trail-D `~` | `no data` | 8 GB `~` | 2.5" + M.2 | no | GbE | `~6–8` | Genuinely low-power, genuinely too slow. |
| M600 Tiny | 2015–16 | Braswell `~` | Pentium N3700 `~` | 8 GB | 2.5" + M.2 | no | GbE | `~6–9` | Lowest-idle Tiny ever. 4 slow Airmont cores. |
| M700 Tiny | 2016 | Skylake `T` | i7-6700T | 32 GB DDR4 | 2.5" + M.2 | no | GbE | `~9–12` | First "modern-feeling" Tiny. |
| M900 Tiny | 2016 | Skylake `T` | i7-6700T | 32 GB DDR4 | 2.5" + M.2 | no | GbE + vPro | `~9–12` | ⚠ rotated feet — check mount fit. |
| M710q | 2017 | Kaby Lake `T` | i7-7700T | 32 GB DDR4 | 2.5" + M.2 | no | GbE | `~9–12` | Cheap 4C/8T. No expansion. |
| M910q | 2017 | Kaby Lake `T` | i7-7700T | 32 GB DDR4 | 2.5" + M.2 | no | GbE + vPro | `~9–12` | vPro AMT is the reason to pick it. |
| **M720q** | 2018–19 | Coffee Lake / -R `T` | i9-9900T | 64 GB DDR4 | 2.5" + M.2 | **x8 (01AJ940)** | GbE (I219) | **~8–10 M** | **The definitive used rack node.** |
| **M920q** | 2018–19 | Coffee Lake / -R `T` | i9-9900T | 64 GB DDR4 | 2.5" + M.2 (2× on some) | **x8 (01AJ940)** | GbE I219-**LM** vPro | **~8–12 M** | M720q + vPro AMT. Worth the premium in a rack. |
| M625q | 2018 | AMD Stoney Ridge | A9-9420e | 8 GB DDR4 | 2.5" + M.2 | no | GbE | `~5–7` | 6 W chip. Almost useless, almost free. |
| M715q Gen 1 / 2 | 2018–19 | AMD Bristol Ridge / Raven Ridge | R5 PRO 2400GE | 32 GB DDR4 | 2.5" + M.2 | no | GbE (Realtek) | **14–17 M** | Higher idle than Intel peers. Skip. |
| M75q-1 (Gen 1) | 2019–20 | AMD Picasso (Zen+) | R5 PRO 3400GE | 32 GB DDR4 | 2.5" + M.2 | no | GbE (Realtek) | **~11 M** | Half its USB is 2.0. Fine budget node. |
| **M75q Gen 2** | 2020–21 | AMD Renoir / Cezanne | R7 PRO 5750GE | 64 GB DDR4 | 2.5" + M.2 | **none** | GbE (Realtek) | **~12 M** | **Best threads-per-dollar used.** No 10GbE path, ever. |
| M70q Gen 1 | 2020 | Comet Lake `T` | i9-10900T | 64 GB DDR4 | 2.5" + M.2 | no | GbE | `~9–12` | 10C/20T at 35 W, no expansion. |
| M70q Gen 2 | 2021 | Comet **+ Rocket** Lake `T` | i9-11900T | 64 GB DDR4 | 2.5" + M.2 | no | GbE | `~10–13` | Rocket Lake SKUs have **AVX-512**. |
| M70q Gen 3 | 2022 | Alder Lake `T` | i9-12900T | 64 GB DDR4/5 | 2.5" + M.2 | no | GbE | `~10–13` | First hybrid. Big ST jump. |
| M70q Gen 4 | 2023 | Raptor Lake `T` | i7-13700T | 64 GB DDR5 | M.2 ×1–2 | no | GbE | `~10–14` | No i9 option. |
| **M70q Gen 5** | 2024–26 | Raptor Lake / -R `T` + Intel 300T | i7-14700T | 64 GB DDR5 | **2× M.2 Gen4 ×4** | **no** | GbE I219-LM (+2.5GbE opt.) | `~10–14` | 20C/28T at 35 W, **but no riser**. RAID 0/1 on NVMe. |
| M80q (Gen 1) | 2020 | Comet Lake `T` | i9-10900T | 64 GB DDR4 | 2.5" + M.2 | no | GbE + vPro | `~9–12` | M70q + vPro. No riser. |
| M80q Gen 3 | 2022 | Alder Lake `T` | i9-12900T | 64 GB DDR5 | 2.5" + M.2 | no | GbE + vPro | `~10–13` | STH reviewed; solid, unexciting. |
| M80q Gen 4 | 2023 | Raptor Lake `T` | i9-13900T | 64 GB DDR5 | M.2 ×1–2 | no | GbE + vPro | `~11–15` | 24C/32T at 35 W with vPro, no riser. |
| **M90q (Gen 1)** | 2020 | Comet Lake `T` + 65 W | i9-10900T / i9-10900 | 64 GB DDR4 | 2.5" + M.2 | **yes** | GbE + vPro | `~10–13` | First of the modern riser-capable tier. |
| **M90q Gen 2** | 2021 | Comet **+ Rocket** `T` + 65 W | i9-11900T | 64 GB DDR4 | 2× M.2 (Gen3) | **yes** | GbE + vPro | **12–14 M** | PCIe **3.0** storage only. AVX-512 on RKL. |
| **M90q Gen 3** | 2022 | Alder Lake `T` + 65 W | i9-12900T / i9-12900 | 64 GB DDR5 | 2× M.2 **Gen4** | **yes** | GbE + vPro, **2.5GbE opt.** | `~11–14` | First Tiny with a factory 2.5GbE option. |
| **M90q Gen 4** | 2023 | Raptor Lake `T` + 65 W | i9-13900T / i9-13900 | 64 GB DDR5 | 2× M.2 Gen4 | **yes** | GbE + vPro | `~12–16` | 24C/32T at 35 W + a PCIe slot. |
| **M90q Gen 5** | 2024 | Raptor **+ Raptor-R** `T` + 65 W | i9-14900T / i9-14900 | 64 GB DDR5 | 2× M.2 Gen4 | **yes** | GbE + vPro | `~12–16` | Peak of the pre-Ultra line. |
| **M90q Gen 6** | 2025–26 | **Core Ultra Series 2 (Arrow Lake-S)** | Ultra 9 285 / 285T | **64 GB DDR5-5600** | **3× M.2 PCIe 4.0 ×4** | **PCIe 4.0 x8 LP** | GbE I219-LM + vPro | `no data` | **Most capable 1 L ever**: Arc A310 dGPU option, I350-T4 quad-GbE option, TB4 punch-out, 13-TOPS NPU. |
| neo 50q Gen 4 | 2023–25 | Raptor Lake-**H** (BGA) | i7-13620H | 64 GB DDR5 | 2× M.2 | no | GbE | `no data` | Mobile silicon in a Tiny shell. Cheap cores. |
| neo 50q Gen 4 *Thin Client* | 2023–26 | Celeron / 12th i3 (BGA) | i3-1215U `~` | 16 GB `~` | 1× M.2 | no | GbE | `~6–9` | Thin-client BIOS; verify it will boot a general OS. |
| neo 50q Gen 5 | 2025–26 | **Raptor Lake-H rebrand** | Core 7 240H | 64 GB DDR5-5600 | 2× M.2 Gen4 | no | GbE | `no data` | 10C/16T new-in-box for very little. **No NPU, no vPro, no riser.** |
| **M75q Gen 5** | 2024–26 | **Zen 4 Phoenix / Phoenix2** | R7 PRO 8700GE | 64 GB official / **128 GB M** | 2× M.2 | no | GbE (Realtek RTL8111FP) | **5–15 M** | STH: *"one of the best project tiny mini micro nodes we've ever reviewed."* 16-TOPS NPU. |
| M90n-1 Nano | 2019–21 | Whiskey Lake-U | i7-8665U | 16 GB `~` | 1× M.2 | no | GbE + vPro (i7/i5-8365U) | `~5–8` | 0.35 L. Lowest-power Think node with real cores. |
| M75n Nano | 2020–21 | Picasso-U / Dali | R5 PRO 3500U | 16 GB `~` | 1× M.2 | no | GbE | `~5–8` | Same idea, AMD. |
| P320 Tiny | 2017 | Kaby Lake `T` | i7-7700T | 32 GB DDR4 | 2.5" + M.2 | **yes (GPU)** | GbE | `~12–16` | Quadro **P600**. No top vent — stackable. |
| P330 Tiny Gen 1 | 2018 | Coffee Lake `T` | i7-8700T | 32 GB DDR4 | 2× M.2 | **yes (GPU)** | GbE + vPro | `~12–16` | Quadro **P620 / P1000**, or **I350-T4** in the slot. |
| P330 Tiny Gen 2 | 2019 | Coffee Lake-R `T` | i9-9900T | 64 GB DDR4 | 2× M.2 | **yes** | GbE + vPro | `~13–17` | 8C/16T + a GPU in 1 L. |
| P340 Tiny | 2020 | Comet Lake `T` | i9-10900T | 64 GB DDR4 | 2× M.2 | **yes** | GbE I219-LM | `~13–18` | Quadro P620 / P1000 / **T1000 6 GB**. **No SATA bay** — that's what buys the slot. |
| P350 Tiny | 2021 | Rocket Lake `T` | i9-11900T | 64 GB DDR4 | 2× M.2 | **yes** | GbE + vPro | `~14–18` | **T400 / T600 / T1000**. AVX-512. Short-lived. |
| P360 Tiny | 2022–23 | Alder Lake `T` | i9-12900T | 64 GB DDR5-4800 | 2× M.2 Gen4, to 4 TB | **yes (PCIe 3.0 x8)** | GbE + vPro | `~14–19` | T400 4 GB / T600 4 GB / **T1000 8 GB**. |
| P3 Tiny | 2023–25 | Raptor Lake `T` | i9-13900T | 64 GB DDR5 | 2–3× M.2 | **yes** | GbE + vPro | `~14–20` | RTX **A400 / A1000**, T1000 8 GB. |
| **P3 Tiny Gen 2** | 2025–26 | **Core Ultra Series 2** | Ultra 9 285 | **128 GB DDR5** (2× 64 GB CSODIMM) | **3× M.2, one PCIe 5.0 ×4** | **PCIe 4.0 x16 (x8 link)** | GbE + up to 2× extra 2.5GbE | `no data` | The most expandable 1 L box that exists. Up to a **330 W** brick. |

*Prices are deliberately absent from this table — see **§1.10** for used-market cost and availability,
which is a different axis from capability and does not track it.*

---

## 1.1 Legacy Tiny (2012–2016)

**M72e · M73 · M83 · M93/M93p · M53 · M600 · M700 · M900**

*Silicon:* Ivy Bridge → Haswell → Bay Trail/Braswell → Skylake, all 35 W `T` desktop parts except the
Atom-derived M53/M600. **Appendix C of `nodes.md` starts at Skylake; the pre-2016 CPU menus here were not
verified this session — confirm on PSREF before quoting.**

*Platform:* DDR3/DDR3L (through M93p), DDR4 from M700/M900. 2.5" SATA bay plus mSATA or a single M.2.
No PCIe riser on any of them. USB 3.0 at best, often half USB 2.0. GbE via Intel I217/I219.

*Where they still make sense:*
- **M900 / M700 (Skylake)** are the floor of "usable in 2026." An i7-6700T is 4C/8T with **HD 530** —
  which still does H.264 and HEVC 8-bit QuickSync, so a single-stream Jellyfin box is viable.
- **M93p** has vPro/AMT, which is a genuine rack feature (out-of-band KVM) on a machine that costs almost
  nothing. If you want to *learn* AMT before trusting it, this is the cheap way.
- **M53 / M600** idle around 6–9 W and are the lowest-power Think boxes with a real x86 chipset. Four
  Airmont cores at ~2 GHz is Raspberry-Pi-4-class throughput on a machine with SATA, DDR3L SO-DIMM and a
  real BIOS.

**Good at:** DNS/DHCP (Pi-hole, Unbound, Technitium), NTP, syslog collector, PXE/netboot server, print
server, Uptime-Kuma-class monitoring, a serial console box, a Tor middle relay, an rsync/Borg backup
target on the SATA bay, a lab machine you're willing to brick.

**Weak at / avoid for:** anything with 2020s single-thread expectations — Home Assistant with a big
automation set, Immich, Nextcloud, modern JS-heavy scraping, any container stack that assumes AVX2 on
Bay Trail/Braswell (**M53/M600 have no AVX2 at all** — this breaks some prebuilt binaries outright).
No AV1, no VP9 hardware decode before Skylake. DDR3 caps you at 16 GB realistically.

**Repairability note:** these are the *most* repairable Tinies — all-screw construction, standard
2.5" SATA, DDR3L SO-DIMM, and Lenovo still publishes Hardware Maintenance Manuals with exploded
diagrams and FRU numbers. Fans and heatsinks are cheap and plentiful. If a box outlives its usefulness
it's because the silicon aged out, not because you couldn't fix it.

---

## 1.2 Intel Tiny4 / Tiny5 — the classic homelab boxes

### M710q / M910q (Kaby Lake, 2017)

4C/8T i7-7700T ceiling, 32 GB DDR4, one M.2 + one 2.5" bay, **no riser**. The M910q adds **vPro/AMT** and
an I219-LM. Cheap, quiet, thoroughly documented.

**Good at:** Proxmox/Docker node for 5–15 light containers, Home Assistant, Frigate with 1–3 cameras using
**UHD 630 QuickSync** for decode plus a USB Coral for detection, Jellyfin single/dual stream, Vaultwarden,
reverse proxy, Wireguard/Tailscale exit node, Gitea, a small Postgres.

**Weak at:** thread-heavy work (4C/8T ceiling), anything needing more than 1 GbE, ZFS with big ARC
(32 GB cap), AV1 anything.

---

### M720q / M920q (Coffee Lake + Coffee Lake-R, 2018–19) — **the reference used node**

This pair is why the Tiny format has a homelab following at all.

*Silicon:* Celeron G4900T through **i9-9900T (8C/16T, 35 W)**. Identical menus. `nodes.md` Appendix C.1
has the full list with GB6.

*Platform:* B360 (M720q) vs **Q370 + vPro** (M920q). 64 GB DDR4-2666 across 2 SO-DIMMs. One 2.5" SATA bay
+ one M.2 2280 NVMe (some M920q SKUs carry a second M.2). GbE I219 / I219-LM.

*The riser — the whole point:* Lenovo's **01AJ940** is a PCIe x16-connector, **x8-electrical** riser
(sometimes sold with the required plastic **baffle/air duct**); **01AJ929** is a x4 variant. Both are
~$9–25 on the grey market. This is the path to 10GbE (Mellanox ConnectX-3, Intel X520/X710), a quad-port
I350-T4, an HBA, or an NVMe carrier.

*Third-party risers worth knowing about:* the motherboard's riser connector actually carries **two** PCIe
links — an **x8 from the CPU** (intended for GPU SKUs) and an **x4 from the PCH** (intended for
Thunderbolt/NIC SKUs). Lenovo's own risers only expose one. **TinyRiser** (FairywrenTech, Tindie) and the
open-source **PowerRiser** (NandFarm) expose both, letting you run a low-profile PCIe card *and* an extra
M.2 NVMe simultaneously. PowerRiser uses an open-ended x8 slot so a physically-x16 card will seat
(still x8 electrically); it requires removing the front Bluetooth-antenna bracket.

*Power:* **~8–10 W idle (AC)**, ~65 W peak. External 65/90/135 W Lenovo slim-tip brick.

**Good at, in rough order of how well the hardware fits:**
Proxmox / XCP-ng / ESXi node · Docker or Podman host (20–40 containers on an i7-8700T) · k3s/k8s worker ·
**Jellyfin/Plex with UHD 630 QuickSync** (H.264 + HEVC 8/10-bit + VP9 decode — still one of the best
transcoders per dollar in existence) · Frigate NVR · **pfSense/OPNsense/VyOS with a real Intel NIC in the
riser** · 10GbE storage frontend · Postgres/MySQL/Redis · Gitea + CI runners · Home Assistant · Immich ·
Nextcloud · Paperless · game servers (Minecraft, Valheim, Satisfactory — 9900T handles several) ·
headless-browser scraping workers · Ansible/Terraform control node · NetBox · Vaultwarden · Wazuh/Graylog
node · Suricata/Zeek sensor (with the NIC) · MinIO/Garage on NVMe · PXE/MAAS provisioning host ·
Sunshine/Moonlight streaming host (weakly).

**Weak at / avoid for:**
- **Bulk storage.** One 2.5" bay and one M.2. A 4-bay NAS this is not; use it as the *head* in front of a
  DAS/JBOD, not as the array.
- **AV1.** Nothing before 11th gen decodes AV1. If your library is going AV1, this is a dead end.
- **Sustained all-core compile.** 35 W and a 1 L cooler. An i9-9900T will complete a kernel build; it will
  do it at ~60 % of the same chip in a tower, and the fan will tell you about it.
- **ECC.** None. Consumer chipset, non-ECC SO-DIMM. Do not build a ZFS "data integrity" story on it.
- **Anything AVX-512.** Coffee Lake has none.

**Buy signal:** M920q over M720q *if* you will actually use AMT for out-of-band power/KVM in a rack —
that's worth the ~$30 premium and removes the need for a JetKVM per node. Otherwise M720q, and put the
difference into the riser and a NIC.

---

## 1.3 AMD Tiny, first wave

### M625q (2018)

A9-9420e, **6 W**, 2 cores, no SMT, DDR4 single-channel in practice. It exists. It idles at ~5–7 W.
**Good at:** a Tor relay, a DNS sinkhole, an NTP server, a serial console, an rsync target.
**Weak at:** everything else, including modern TLS termination at any rate.

### M715q Gen 1 / Gen 2 (2018–19)

Bristol Ridge (Excavator) and Raven Ridge (Zen 1). ServeTheHome measured **14–17 W idle** on a quad-core
unit — *higher* than contemporaneous Intel Tinies. Realtek NIC. Half USB 2.0.
**Verdict:** the worst watts-per-thread in the whole roster. Only worth it free.

### M75q-1 / M75q Gen 1 (2019–20)

Picasso (Zen+). Top SKU **Ryzen 5 PRO 3400GE** (4C/8T, Vega 11, 35 W). **~11 W idle (M)**, ~50 W max.
Half the USB ports are 2.0, which STH called out as inexcusable for a 2019 machine.
**Good at:** Proxmox node, Docker host, general containers, light 1080p transcode via VAAPI.
**Weak at:** anything USB-throughput-bound (SDR arrays, USB DAS, many cameras), 10GbE (no riser, ever).

### M75q Gen 2 (2020–21) — **best used threads-per-dollar**

Renoir (Zen 2) and Cezanne (Zen 3), up to **Ryzen 7 PRO 5750GE — 8C/16T at 35 W**. 64 GB DDR4-3200.
**~12 W idle (M)**, ~55 W max.

**Good at:** VM density (8C/16T for ~$150 used is unmatched in this roster), compile/build workers,
CI runners, container hosts, Kubernetes workers, Postgres, scraping worker pools, game servers.
Zen 2/3 have **SHA-NI**, which meaningfully speeds ZFS SHA-256 checksums and TLS.

**Weak at:**
- **No PCIe riser at any generation.** This is an Intel-only feature on the Tiny line. Your networking
  ceiling is 1 GbE + USB. That single fact disqualifies it from firewall, 10GbE-frontend and NIC-heavy roles.
- **Realtek NIC** — fine on Linux, historically the weak spot on BSD-based firewall distros.
- **Transcoding.** Vega VCN via VAAPI works but is less universally supported than QuickSync in the
  Jellyfin/Plex ecosystem, and Plex hardware transcode support for AMD on Linux has been the rougher path.
  If media is the job, buy Intel.
- **No AMT.** AMD DASH exists on PRO parts but tooling is thinner.

---

## 1.4 Modern Intel Tiny — M70q / M80q / M90q

The three-tier split is the thing to internalise:

| Tier | vPro | PCIe riser | Who it's for |
|---|---|---|---|
| **M70q** | optional on higher SKUs (Gen 4/5) | **never** | cheapest cores |
| **M80q** | yes | **never** | cores + remote management |
| **M90q** | yes | **yes** | cores + management + **expansion** |

**In a rack, the M90q is usually the only one of the three worth paying a premium for**, because the riser
is what turns a 1 L box into a router, a 10GbE node, or a quad-NIC appliance.

### M70q Gen 1–4 / M80q Gen 1, 3, 4 (2020–2023)

Comet Lake → Rocket Lake → Alder Lake → Raptor Lake, all 35 W `T`. Ceilings: **i9-10900T (10C/20T)**,
**i9-11900T (8C/16T, AVX-512)**, **i9-12900T (16C/24T)**, **i9-13900T (24C/32T)** — the last of which is
genuinely remarkable: 24 cores in a 1 L box drawing ~12–16 W idle.

**Good at:** dense virtualization, container hosts, k8s workers, build/CI farms, batch scraping, video
transcoding (UHD 730/770 on 12th gen+ adds **AV1 decode**), Immich with ML, local LLM inference on CPU
(a 12900T with 64 GB will run a 30 B-class quantized model slowly but usefully).

**Weak at:** networking beyond 1 GbE (no riser — this is the recurring theme), sustained all-core work
(a 13900T at 35 W spends most of its life throttled — you're buying *E-cores for parallel throughput*,
not sustained P-core turbo), ESXi on hybrid P+E silicon (Proxmox and modern Linux handle it; older
hypervisors do not schedule it well).

⚠ **13th/14th-gen instability:** the Raptor Lake Vmin-shift degradation issue primarily affected high-power
K-series parts. The 35 W `T` parts are far lower risk by construction, but if you buy a 13th/14th-gen Tiny,
**update the BIOS to a build carrying microcode 0x12B or later before putting it into 24/7 service.**

### M70q Gen 5 (2024–26)

13th/14th-gen `T` plus the entry **Intel 300T** (2C/4T Raptor Lake). Q670, DDR5 to 64 GB.
**Two M.2 2280 slots, both PCIe 4.0 ×4, with RAID 0/1** — and **no 2.5" bay, no PCIe riser**. Optional
factory **2.5GbE (Realtek RTL8125BGS)** via punch-out; optional 100 W PD-in USB-C.

**Good at:** a fast, quiet, all-NVMe container/VM node with mirrored boot; anything wanting an i7-14700T's
20C/28T at 35 W; a 2.5GbE-connected worker.
**Weak at:** any role needing a real NIC or an HBA. Buy an M90q instead.

⚠ **Used pricing is not competitive.** The only observed refurb ask is **$799.99** (i7-14700T · 16 GB ·
512 GB, manufacturer-certified). For that money you can buy **three M75q Gen 2s** (24C/48T total) or an
M920q plus a riser, a 10GbE NIC and a spare machine. Revisit in 2027–28.

### M90q Gen 1–5 (2020–2024)

Same silicon ladder as M70q/M80q **plus 65 W non-`T` SKUs** and — critically — **the PCIe slot**.
Gen 2 measured **12–14 W idle (M)**. Gen 3 introduced DDR5, PCIe 4.0 NVMe, and a **factory 2.5GbE option**;
Gen 2 and earlier are PCIe 3.0 storage only.

**Good at:** everything the M70q is good at, plus **firewall/router with an Intel NIC**, **10GbE storage
frontend**, quad-NIC lab switch/router, HBA-attached JBOD head, IDS/IPS sensor with a real capture NIC,
and any role where AMT out-of-band management matters.

**Weak at:** the 65 W `non-T` SKUs under sustained load — their 148–224 W PL2 is fantasy in a 1 L box.
GB6 MT figures for those chips reflect a well-cooled desktop, not what the Tiny sustains.

### M90q Gen 6 (2025–26) — **the most capable 1 L Lenovo has built**

Verified from PSREF (April/June 2026):

- **Core Ultra Series 2 (Arrow Lake-S), LGA1851, Q870 chipset.** Ultra 5 225 / 235 / 245 / 235T / 245T,
  Ultra 7 265 / 265T, Ultra 9 285 / 285T. All carry **Intel AI Boost NPU up to 13 TOPS**.
- **64 GB DDR5-5600** (2× 32 GB SODIMM).
- **Three M.2 2280 slots, all PCIe 4.0 ×4**, up to 2 TB each, RAID 0/1 (RAID preset special-bid only),
  plus a fourth M.2 for WLAN.
- **One PCIe 4.0 x8 low-profile slot.**
- **Optional discrete Intel Arc A310 4 GB GDDR6** (2× DP 2.0) — a real dGPU in a 1 L Tiny.
- **Optional Intel I350-T4 quad GbE (PCIe x4)**, or 2.5GbE Realtek via punch-out, or **Thunderbolt 4**
  via punch-out port 2, or 4× serial via a PCIe x1 card.
- 179 × 182.9 × **36.5 mm**, ~1.34 kg. Adapters: 65 / 90 / 135 / 230 / 245 W.
- MIL-STD-810H, ENERGY STAR 9.0, TCO gen 10.

**Good at:** basically the full task list. Three NVMe on independent Gen4 ×4 links makes it a credible
all-flash ZFS/Ceph node; the x8 slot still takes 10GbE; Arc A310 brings **AV1 encode and decode** plus a
usable transcode/compute target; the NPU handles Frigate/Immich/Whisper-class inference without a Coral
or Hailo; TB4 opens external GPU/NVMe enclosures.

**Weak at:** the same 1 L thermal ceiling — an Ultra 9 285 (65 W) still throttles, so prefer the **285T**
for a rack. No ECC. And note the **Arc A310 occupies the riser slot**, so you choose between a dGPU and a
NIC, not both.

⚠ **Availability: effectively zero on the used market.** Released 2025–26 and still in its first
deployment cycle. Nothing here is a buying recommendation — it is a watch-list entry for roughly 2028,
when the first three-year lease returns land. See §1.10.

---

## 1.5 Modern AMD Tiny — M75q Gen 5

*Silicon:* Zen 4 — **Phoenix2** (hybrid Zen4 + Zen4c, Radeon 740M, **no NPU**) on the 8300GE/8500GE/8505GE,
and **monolithic Phoenix** (Radeon 760M/780M, **~16 TOPS Ryzen AI NPU**) on the 8600GE/8700GE/8705GE.
65 W `G` SKUs also qualified; they throttle.

*Platform:* DDR5-5200, 2 SO-DIMM slots. **Lenovo says 64 GB max; ServeTheHome verified 96 GB and 128 GB
working with Crucial 2× 48 GB and 2× 64 GB kits.** Two M.2 slots. **GbE via Realtek RTL8111FP with AMD DASH.**
Memory and both M.2 slots moved to the top side with toolless clips — a real serviceability improvement.

*Power (M, ServeTheHome, 8700GE):* **5–15 W idle**, **63–66 W sustained** under load, holding the power
level rather than dropping — at the cost of **47–48 dBA** fan noise.

**Good at:** the highest sustained multi-thread throughput of any 35 W 1 L box on the AMD side; VM density
with 128 GB of RAM (STH's point: this can replace an older single-socket Xeon E5 virtualization host with
2× the per-core performance); compile farms; **local inference** on the 780M iGPU + NPU; **AV1 decode and
encode** via RDNA 3 VCN — the only AMD option in this roster that does AV1 encode.

**Weak at:** **no PCIe riser** (still), so 1 GbE is the ceiling without USB or Thunderbolt tricks; Realtek
NIC; **loud under sustained load** — 47–48 dBA is a real consideration if the rack is in a living space;
AMD's transcode/compute stack is still the less-travelled path in Jellyfin/Plex/Frigate documentation.

**Repairability:** best-in-class for the modern era. Toolless RAM and both M.2 slots, standard SO-DIMM,
standard M.2 2280, Lenovo-published HMM, and a common slim-tip brick shared across the whole ThinkPad
and Tiny line. This is the one modern Tiny where a ten-year service life is a reasonable expectation —
which makes it the model most worth waiting for on the used market.

⚠ **Availability: very thin.** Shipped from 2024, so only the earliest lease returns are circulating and
asks (`~$350–550`) still sit close to retail. **The right machine at the wrong moment.** Watch it from
2027; by then it should land in the $200–300 band, and at that price it is the successor to the M720q's
"just buy this one" status.

---

## 1.6 Budget / mobile-silicon Tiny — neo 50q

Lenovo's SMB line. The important structural difference: **these use BGA mobile H-series silicon soldered
to the board.** You cannot change the CPU, there is no riser, and there is no vPro.

### neo 50q Gen 4 (2023–25)
13th-gen Raptor Lake-**H**: i5-13420H (8C/12T), i7-13620H (10C/16T), and similar. DDR5 SODIMM ×2, 2× M.2.

### neo 50q Gen 4 **Thin Client** (2023–26)
A separate PSREF product: **Intel Celeron or 12th-gen Core i3**. Ships with a thin-client image.
⚠ Verify BIOS/boot behaviour before assuming it will install a general-purpose OS cleanly.

### neo 50q Gen 5 (2025–26)
**Core 5 210H** (8C/12T, 2.2/4.8 GHz, 12 MB L3, UHD Xe G4 48 EU) or **Core 7 240H** (10C/16T, 2.5/5.2 GHz,
24 MB L3, 64 EU). Raptor Lake-H silicon under a "Series 2" marketing name. DDR5-5600 to 64 GB, 2× M.2
PCIe 4.0. 179 × 182.9 × 34.5 mm bare / 36.5 mm with feet — **the shortest modern Tiny**, which matters at 1U.
Wi-Fi 6/6E/7 options, Bluetooth 5.4. GbE only.

Aggregator GB6 for the 240H: **~2632 ST / ~12604 MT**; Notebookcheck placed it roughly level with an
i7-12800H / i7-13700H.

**Good at:** 10 modern cores in a small box; container host; CI runner; compile worker; general Proxmox
node; Jellyfin (the 64-EU Xe-derived iGPU does H.264/HEVC/**AV1 decode**).

**Weak at:** rack-grade manageability (**no vPro/AMT** — budget a JetKVM), expansion (**no riser, no 2.5GbE
option**), CPU upgrades (soldered), and the 45 W base / 115 W turbo power envelope means it burns more at
the wall under load than a 35 W `T` part for similar throughput.

⚠ **Availability: Gen 5 is unobtainable used; Gen 4 is thin (`~$180–330`).** Even at that, the Gen 4 is
hard to justify against an M75q Gen 2 at $120–220 — you get the same core-count class, more RAM headroom,
and a socketed CPU. **The neo line's whole appeal was its new-unit price, and that appeal does not
survive the move to a used-only budget.**

---

## 1.7 Nano — sub-1 L

**M90n-1** (Intel) and **M75n** (AMD), ~0.35 L, both withdrawn. Mobile `U`/`Y` silicon: Celeron 4205U,
i3-8145U, i5-8265U, **i5-8365U / i7-8665U with vPro** on the Intel side; **Ryzen 3 PRO 3300U / Ryzen 5 PRO
3500U**, Athlon Silver 3050e on the AMD side. One M.2, no bay, no riser, GbE.

**Good at:** the lowest-power Think node that still has real out-of-order cores and (on the i5-8365U/
i7-8665U) **AMT out-of-band management** — an unusual combination at ~5–8 W idle. Excellent for DNS,
reverse proxy, monitoring, VPN gateway, a Tor relay, a serial/console server, an always-on jump host,
or as a cluster of tiny k3s control-plane nodes.

**Weak at:** RAM ceiling (~16 GB), single M.2, 15 W thermal envelope, and — for a rack — **you'll fit two
or three side by side in 1U, which raises the question of why not just use one bigger box.** Density
without a mount that exploits it is wasted.

---

## 1.8 ThinkStation Tiny — the GPU-capable 1 L

Same chassis, different board: the riser slot is populated at the factory with a **low-profile, single-slot
professional GPU** fed by a heat-pipe from the main cooler. In exchange, **the 2.5" SATA bay is deleted** on
most of them — that's the trade that makes the slot possible.

| Model | CPU gen | GPU options | Memory | Notes |
|---|---|---|---|---|
| P320 Tiny | 7th Kaby Lake `T` | Quadro **P600 2 GB** | 32 GB DDR4 | Every SKU shipped with the P600. **No top vent** — the only Tiny you can safely stack. |
| P330 Tiny Gen 1 | 8th Coffee Lake `T` | Quadro **P620 2 GB**, **P1000 4 GB**, or **Intel I350-T4** in the slot | 32 GB DDR4 | Riser-connected; dual M.2 |
| P330 Tiny Gen 2 | 9th Coffee Lake-R `T` | P620 / P1000 | 64 GB DDR4 | Up to i9-9900T |
| P340 Tiny | 10th Comet Lake `T` | P620 2 GB, P1000 4 GB, **T1000 6 GB** | 64 GB DDR4 | Adds a **top vent** (no longer stackable). ⚠ P330 risers are **not** interchangeable. |
| P350 Tiny | 11th Rocket Lake `T` | **T400 / T600 / T1000** | 64 GB DDR4 | AVX-512. Short production run. |
| P360 Tiny | 12th Alder Lake `T` | T400 4 GB, T600 4 GB, **T1000 8 GB** (all **PCIe 3.0 ×8**) | 64 GB DDR5-4800 | Up to 4 TB M.2 PCIe 4.0; 6 displays |
| P3 Tiny | 13th Raptor Lake `T` | **RTX A400 4 GB**, **RTX A1000 8 GB**, T1000 8 GB | 64 GB DDR5 | Withdrawn on PSREF |
| **P3 Tiny Gen 2** | **Core Ultra Series 2** | T1000 8 GB, **RTX A1000 8 GB**, **RTX A400 4 GB** (all 50 W, single slot, 4× miniDP) | **128 GB DDR5** (2× 64 GB CSODIMM, DDR5-5600 SODIMM or DDR5-6400 CSODIMM) | **3× M.2 — two Gen 4, one PCIe 5.0 ×4**, NVMe RAID 0/1/5; riser is **PCIe 4.0 x16 (x8 link)** for GPU **or** PCIe 4.0 x4 for NIC/serial; **I350-T2/T4, BCM5719/5720, or 2.5GbE via BTB**; up to **330 W** adapter; Q870 |

**Good at, specifically because of the GPU:**
- **NVENC transcoding.** T400/T600/T1000/A400/A1000 all carry NVIDIA's encoder. Turing-generation
  T-series does H.264 + HEVC; **Ampere RTX A400/A1000 add AV1 *decode***. This is the Jellyfin/Plex
  "many simultaneous streams" answer in 1 L — and unlike consumer GeForce, **Quadro/RTX-pro cards have no
  NVENC concurrent-session limit**, which is the entire reason people put them in transcode boxes.
- **Local LLM / vision inference.** An RTX A1000 with 8 GB GDDR6 runs 7–8 B quantized models comfortably
  and does Frigate/DeepStack/CodeProject.AI detection with room to spare.
- **Whisper / ASR**, Stable Diffusion (slowly, 8 GB), video encoding pipelines, CUDA-dependent
  scientific work, remote CAD/VDI via Sunshine or Parsec.
- **P3 Tiny Gen 2 as a NIC box instead:** swap the GPU option for I350-T4 and you have a quad-GbE 1 L
  router with 128 GB of RAM and three NVMe. That is a serious appliance.

**Weak at:**
- **Watts and noise.** A 50 W GPU inside a 1 L box means a ~14–20 W idle floor and an audible fan under
  load. These are the least efficient nodes in the roster per idle watt.
- **Price.** Even used, P-series Tinies command a premium over the M-series equivalent for the same CPU —
  typically 40–80 % more. The GPU is most of that delta.
- **Availability skews old.** P320/P330/P340 are plentiful; **P3 Tiny Gen 2 is unobtainable used** and
  P3 Tiny is thin at `~$450–750`. The used-market sweet spot in this family is the **P340 Tiny with a
  T1000** at `~$180–330`.
- **The GPU is not really upgradeable** in the way a desktop's is — the cards use custom Lenovo brackets
  and a heat-pipe interface. Community swaps to T600/T1000 work; anything larger is a project. Risers are
  **not** cross-generation compatible (documented failures moving P330 risers into a P340).
- **No SATA bay** on most, so bulk local storage is off the table.

---

## 1.9 Excluded and warning rows

| Device | Why it's here | Verdict |
|---|---|---|
| **M910x / M920x** | ~1.35 L, taller chassis, second M.2 + PCIe slot, shares the **01AJ940** riser with Tiny5 | **Fits 10" width, needs 2U.** If you were going to use 2U anyway, the x-variant gives you the riser *and* a second NVMe without giving up the SATA bay. |
| **ThinkStation P3 Ultra / P3 Ultra SFF (Gen 2)** | ~3.6 L; W680/Q870; up to 192 GB DDR5 ECC (Ultra SFF Gen 1); **RTX A1000/A400/T1000 or Intel Arc Pro B50 16 GB ECC**; two GPUs on Gen 2 | **DOES NOT FIT** a 10" rack. Listed because it is the obvious "next step up" and people mistake it for a Tiny. |
| **ThinkCentre M-series SFF (`s`) / Tower** | 8–18 L | **DOES NOT FIT.** |
| **ThinkEdge SE10 / SE30** | Think-branded compact edge boxes on Atom/Core-U class silicon, DIN-rail and industrial mounting, wide temperature range | Not the 1 L Tiny chassis; different mounting ecosystem entirely. Mentioned for completeness — if you need −20 °C to 60 °C operation these are the Think answer, not a Tiny. |
| **Tiny-in-One (TIO) 22/24/27** | Monitor shells that a Tiny slots into | Irrelevant to a rack, but note: **a Tiny bought "for TIO" may ship with a higher-wattage brick**, which is a bonus. |


---

## 1.10 Used-market pricing and availability

**All prices in this document are used / refurbished. New pricing is not quoted anywhere, by design.**

### Provenance — read this before trusting a number

| | |
|---|---|
| **Source** | eBay US listings and completed/sold indicators, read 2026-07-24, plus the July-2026 verification pass carried in `nodes.md`'s own pricing table. |
| **What the numbers are** | **Asks, not sold prices.** eBay's sold-listing archive was not readable this session. An ask is an upper bound on what a patient buyer pays; assume **10–25 % below ask** is achievable on Best Offer for anything that isn't scarce. |
| **Tag** | Observed asks are tagged **V** (a seller's claim). Ranges without an observed anchor are tagged **`~`** — they are inferred from market shape, **not measured**. |
| **Region** | US. European and AU pricing runs meaningfully higher on the older tiers and lower on ex-lease corporate stock; adjust locally. |
| **Volatility** | **High, and directional.** See the DRAM note immediately below. Re-check before buying. |

### ⚠ The 2025–26 DRAM shortage has broken the normal pricing logic

`nodes.md` documents the shortage (Raspberry Pi 5 16 GB $120 → **$305**; Radxa X4 8 GB $79.96 → **$265.99**,
+233 %). On the used 1 L market it produces a specific and counter-intuitive distortion:

**Used Tiny prices now track installed RAM and SSD far more closely than they track the CPU.**

Four observed asks from the same week make the point:

| Listing (observed ask, **V**) | Price | What it tells you |
|---|---|---|
| M920q · **i5-8500T** · 8 GB · 256 GB | **$160.99** | A vPro-capable, riser-capable 6-core box |
| M720q · **i5-8400T** · 8 GB · 120 GB | **$164.85** | Same class, slightly worse CPU, *same price* |
| M720q · **i5-8500T** · **32 GB** · 120 GB | **$318.02** | **+24 GB of DDR4 ≈ +$157** |
| M910q · **i7-6700T** · **32 GB** · 512 GB | **$299.99** | A *worse, older, riser-less* machine at ~2× the M920q, purely on RAM content |

**The buying rules that fall out of this:**

1. **Buy the RAM, take whatever CPU comes with it.** Across the 6th–9th-gen tier the CPU delta is worth
   $20–40; the RAM delta is worth $100–200. Filling an 8 GB box to 32 GB costs more than the box.
2. **Avoid barebones listings.** "No RAM / no SSD / no charger" units (common for M90q Gen 1 and M720q)
   look like a bargain and are currently a trap — you are opting into the most inflated part of the BOM.
   The exception is if you already have spare DDR4/DDR5 SO-DIMMs on the shelf.
3. **Prefer models whose generation shares your existing RAM type.** A rack of DDR4 Tinies lets you shuffle
   SO-DIMMs between nodes. Mixing DDR4 (≤ M90q Gen 2) and DDR5 (M90q Gen 3+) doubles your spares problem.
4. **The SSD is the cheap part to ignore.** Most listings ship a 120–512 GB boot drive; NVMe pricing has
   been far less affected than DRAM. Buy on RAM, replace the SSD later if you care.
5. **Buy in multiples where sellers offer it.** One observed M910q listing dropped from $299.99 to $287.99
   at quantity 4. Refurbisher stock is lot-based; bulk pricing is normal and worth asking for.

### Availability tiers

The single most important axis for a used-only buyer, and it does **not** correlate with the capability
rankings in Parts 1–4.

| Tier | Meaning | Models |
|---|---|---|
| **Abundant** | Hundreds of listings; buy any day, name your price | M700 · M900 · M710q · M910q · **M720q · M920q** · M715q · M75q-1 · M70q Gen 1 · M80q Gen 1 · M90q Gen 1 · P320 · P330 Gen 1/2 |
| **Plentiful** | Dozens of listings; mild patience needed for a specific CPU/RAM combo | **M75q Gen 2** · M70q Gen 2 · M90q Gen 2 · M625q · P340 |
| **Moderate** | Present but you take what's listed; expect to wait for a good config | M70q Gen 3 · M80q Gen 3 · M90q Gen 3 · P350 · P360 · M90n-1 Nano |
| **Thin** | Occasional listings, prices near new, little negotiating room | M70q Gen 4 · M80q Gen 4 · M90q Gen 4 · P3 Tiny · neo 50q Gen 4 · M75n Nano |
| **Very thin** | First lease returns only; refurb asks still close to retail | M90q Gen 5 · **M70q Gen 5** · **M75q Gen 5** |
| **Effectively unobtainable used** | Still in first deployment. Not a purchase option today. | **M90q Gen 6** · **P3 Tiny Gen 2** · **neo 50q Gen 5** |

**The uncomfortable implication:** the three most capable machines in Part 1 — M90q Gen 6, P3 Tiny Gen 2
and (barely) M75q Gen 5 — are **not available to a used-only buyer at a sane price.** Everything below is
scoped to what you can actually get.

### Used price table

`~` = inferred range, not observed. **V** = observed ask, 2026-07-24. "Working config" = a unit that
arrives with RAM, an SSD and the power brick — the only configuration worth costing.

| Model | Availability | Working config `~` | Observed asks (**V**) | Cost notes |
|---|---|---|---|---|
| M72e / M73 / M83 Tiny | Abundant | `~$25–60` | no data | Below the price of the RAM inside them. Free-tier only |
| M93 / M93p Tiny | Abundant | `~$40–85` | no data | Cheapest box in the roster with **vPro/AMT** |
| M53 / M600 Tiny | Plentiful | `~$35–75` | no data | Cheapest way to a ~7 W always-on x86 node |
| M700 Tiny | Abundant | `~$45–100` | no data | DDR4, so RAM is shareable with the M7xx/M9xx tier |
| M900 Tiny | Abundant | `~$50–110` | no data | ⚠ rotated feet — mount fit |
| M710q | Abundant | `~$55–120` | no data | 4C/8T i7-7700T configs at the top of the range |
| M910q | Abundant | `~$70–150` | **$299.99** (i7-6700T · 32 GB · 512 GB); $287.99 at qty 4 | The observed ask is a **RAM price, not a machine price** |
| **M720q** | Abundant | **`~$110–200`** | **$164.85** (i5-8400T · 8 GB · 120 GB) · **$161.49** (i5-8400T · 12 GB) · **$249.99** (i5-9500T · 16 GB · 256 GB) · **$318.02** (i5-8500T · 32 GB) | `nodes.md` (2026-07-16): ~$100–200 typical, asks to ~$380. **Best value in the roster** |
| **M920q** | Abundant | **`~$130–230`** | **$160.99** (i5-8500T · 8 GB · 256 GB, reduced from $239.99) | `nodes.md`: ~$130–220. vPro for ~$0 premium at the moment — take it |
| M625q | Plentiful | `~$35–75` | no data | |
| M715q Gen 1 / 2 | Abundant | `~$55–110` | no data | Cheap for a reason. Worst watts-per-thread here |
| M75q-1 | Plentiful | `~$85–150` | no data | |
| **M75q Gen 2** | Plentiful | **`~$120–220`** (4650GE/16 GB) · `~$200–320` (5750GE 8C/16T) | no data | `nodes.md` (2026-07-16): ~$120–200. **Cheapest 8C/16T in the roster** |
| M70q Gen 1 | Abundant | `~$110–190` | no data | i9-10900T configs at the top |
| M70q Gen 2 | Plentiful | `~$130–230` | no data | The **only** used route to AVX-512 alongside M90q Gen 2 / P350 |
| M70q Gen 3 | Moderate | `~$180–300` | no data | |
| M70q Gen 4 | Thin | `~$250–420` | no data | |
| M70q Gen 5 | Very thin | `~$550–800` | **$799.99** (i7-14700T · 16 GB · 512 GB, certified refurb, mfr warranty to 2028) | Refurb ask is ~62 % of list. **Not competitive per dollar** — three M75q Gen 2s cost less and give more threads |
| M80q Gen 1 | Abundant | `~$120–200` | no data | M70q + vPro, same price band |
| M80q Gen 3 | Moderate | `~$200–330` | no data | |
| M80q Gen 4 | Thin | `~$300–470` | no data | 24C/32T at 35 W with vPro — the value pick *if* you don't need a riser |
| **M90q Gen 1** | Abundant | **`~$140–260`** | barebones "no RAM / SSD / charger" listings common | ⚠ barebones is a trap under current DRAM pricing. **Cheapest riser-capable modern Tiny** |
| M90q Gen 2 | Plentiful | `~$180–300` | no data | PCIe **3.0** storage only |
| M90q Gen 3 | Moderate | `~$280–450` | listings observed at i5-12500 / 16 GB DDR5 | First 2.5GbE + PCIe 4.0 + DDR5 |
| M90q Gen 4 | Thin | `~$400–620` | no data | |
| M90q Gen 5 | Very thin | `~$550–900` | no data | |
| **M90q Gen 6** | **Unobtainable used** | — | — | Watch from ~2028 |
| neo 50q Gen 4 | Thin | `~$180–330` | no data | Mobile BGA silicon; cheap cores, no riser, no vPro |
| neo 50q Gen 5 | **Unobtainable used** | — | — | |
| **M75q Gen 5** | Very thin | `~$350–550` | no data | First lease returns only. Excellent machine, wrong moment |
| M90n-1 Nano | Moderate | `~$80–170` | no data | i7-8665U/i5-8365U configs carry **vPro at ~6 W idle** — a genuine bargain if you find one |
| M75n Nano | Thin | `~$80–150` | no data | |
| P320 Tiny | Abundant | `~$90–170` | no data | Every unit has a Quadro P600 |
| P330 Tiny Gen 1 | Abundant | `~$120–220` | no data | P620/P1000 |
| P330 Tiny Gen 2 | Plentiful | `~$150–270` | no data | Up to i9-9900T **with** a GPU |
| P340 Tiny | Plentiful | `~$180–330` | no data | T1000 6 GB configs at the top |
| P350 Tiny | Moderate | `~$220–390` | no data | Short production run → thinner supply than neighbours |
| P360 Tiny | Moderate | `~$300–520` | no data | T1000 8 GB configs at the top |
| P3 Tiny | Thin | `~$450–750` | no data | RTX A400/A1000 |
| **P3 Tiny Gen 2** | **Unobtainable used** | — | — | |

### Parts, accessories and consumables — used / grey-market

The machine is often not the expensive part. These are the line items that decide a build's real cost.

| Item | Used / grey-market price | Notes |
|---|---|---|
| **Lenovo 01AJ940 riser** (x16 connector, x8 electrical — Tiny5) | **$8.99–$25 V** | Observed: $8.99 (31 sold) with **"OEM with baffle" / "OEM no baffle"** variants; $16.98 elsewhere. **Get the baffle version** |
| **Lenovo 01AJ929 riser** (x4) | **~$17 V** | For NIC/serial, not GPU |
| **TinyRiser / PowerRiser** (dual-link community risers) | `~$40–70` | Small-run, new-only from the makers. The only way to get a PCIe card **and** an extra NVMe simultaneously |
| P-series factory GPU riser | `~$25–60` | **Not cross-generation compatible** — buy the one for your exact model |
| **Intel I350-T4 quad GbE (low profile)** | `~$25–45` | The classic riser card. Ubiquitous ex-server stock |
| **Mellanox ConnectX-3 SFP+** | `~$15–35` | Cheapest 10GbE. Confirmed by a buyer as fitting the M720q riser with the chassis side screw |
| Intel X520-DA2 / X710 | `~$30–70` | |
| **10" printed rack mount** (filament cost, self-printed) | **~$3–8** | ~130 g PETG per the M900-specific design. **This is the cheapest part of the build** |
| **10" commercial mount** (3drackmounts) | **$39.99 V** | Injection/print, 1U, ThinkCentre-specific |
| 10" JetKVM mount (3drackmounts) | **$24.99–39.99 V** | Relevant only for the non-vPro models |
| **Lenovo slim-tip brick, 65 W, used** | `~$10–20` | Interchangeable across the whole ThinkPad + Tiny line. Buy spares; they are the most common failure |
| Slim-tip brick, 135 W / 170 W+ | `~$25–50` | Required for USB-C punch-out and GPU/NIC configs |
| **DDR4 SO-DIMM 16 GB, used** | ⚠ **shortage-inflated** — re-check | The dominant cost variable. See the DRAM note above |
| **DDR5 SO-DIMM 32 GB** | ⚠ **shortage-inflated** — re-check | Worse than DDR4. A reason to prefer the DDR4 generations right now |
| NVMe 1 TB, used | `~$40–70` | Comparatively sane |
| 2.5" SATA SSD 2–4 TB | `~$70–180` | Only relevant on pre-Tiny7 chassis |

### Worked build costs (used, mid-2026)

Illustrative, using the midpoints above. **Excludes RAM upgrades**, because under current DRAM pricing
you should be buying units that already have the RAM rather than upgrading them.

| Build | Parts | `~` total |
|---|---|---|
| **Minimum viable rack node** | M720q (i5-8500T, 16 GB, 256 GB) + printed mount | **`~$160–200`** |
| **The reference node** | M920q (i5-8500T/i7-8700T, 16 GB) + 01AJ940 riser w/ baffle + I350-T4 + printed mount + spare brick | **`~$220–290`** |
| **10GbE storage frontend** | M920q (i7-9700T, 32 GB) + riser + ConnectX-3 + 2× NVMe + printed mount | **`~$380–500`** |
| **Cheap thread farm, 3 nodes** | 3× M75q Gen 2 (R5 PRO 4650GE, 16 GB) + 3 printed mounts | **`~$400–650`** for 18C/36T |
| **Media node with real NVENC** | P340 Tiny (i7-10700T, 16 GB, T1000) + printed mount | **`~$200–340`** |
| **Always-on service tier, 2 nodes** | 2× M90n-1 Nano (i5-8365U, vPro) + mounts | **`~$180–350`** at ~12 W combined |
| **Modern flagship** | M90q Gen 6 or P3 Tiny Gen 2 | **not purchasable used** |

### Where the money actually goes

- **Under $200, the M720q/M920q is unbeaten** and nothing in the roster changes that. Every "better"
  machine costs 2–4× for 1.5–2× the throughput.
- **Watts are a rounding error at this price level.** `nodes.md` runs the arithmetic: a ~13 W x86 node
  costs ~$17/yr at $0.15/kWh, a 1.6 W ARM node ~$2/yr. A $150 price difference between two Tinies pays
  for roughly **a decade** of the watt gap. Buy on capability and availability, not idle power — unless
  you are running six of them, at which point it starts to matter.
- **The riser is the highest-leverage $10–25 you will spend** on this platform. It is the difference
  between a compute node and an appliance, and it is cheaper than a single stick of RAM.
- **Print your own mounts.** $3–8 of PETG versus $40 commercial, for a part with no reliability
  consequence beyond "print it in PETG, not PLA."
- **Buy one spare machine, not spare parts.** At $110–200 for a whole M720q, a cold spare is cheaper
  than sourcing a board or a cooler, and gives you a donor for the brick, the fan, the WLAN card and the
  bottom plate. This is the single best reliability move available on this platform.

---

# Part 2 — The chip table

**Every processor found in a Think-branded 1 L Tiny or sub-1 L Nano**, in one table. Sorted Intel-first by
microarchitecture, then AMD, then mobile/Nano parts.

### How to read this table

- **Tier** — `T` = 35 W Intel low-power desktop · `non-T` = 65 W desktop (M90q / M75q Gen 5 only, **throttles
  in 1 L**) · `GE` = 35 W AMD low-power APU · `G` = 65 W AMD APU · `U`/`H` = mobile BGA (Nano, neo 50q).
- **C/T (P+E)** — physical cores / threads. Hybrid parts show the P+E split.
- **TDP** — package power limit. `(cTDP-down)` where Intel publishes one; `(PL2 n)` = turbo power ceiling.
  **This is not wall power.** See §3.3.
- **iGPU (NPU)** — integrated graphics, with NPU TOPS where present.
- **Mgmt** — `vPro` = Intel vPro (AMT out-of-band KVM — a genuine rack feature) · `vPro Ent` = vPro
  Enterprise · `DASH` = AMD PRO manageability · `—` = none.
- **GB6** — Geekbench 6 single / multi. **Aggregator medians, not primary reads** (cpu-monkey, nanoreview,
  cputronic, chipversus). ±5–10 %. `~` = additionally low-sample / thermally noisy. `*` = median
  implausibly low, do not rank on it. `no data` = no exact-SKU score found — **not guessed**.
- **ECC:** none of these are usable with ECC in a Tiny. Consumer chipsets, non-ECC SO-DIMM. Omitted as a
  column because the answer is "no" for all 150 rows.
- **Model shorthand:** `M70q1…M70q5`, `M80q1/3/4`, `M90q1…M90q6`, `M75q1/2/5`, `M90n-1`, `M75n`,
  `neo50q4/5`, `P320/P330-1/P330-2/P340/P350/P360/P3T/P3T2`.

| CPU | µarch (gen) | Tier | C/T (P+E) | Base / Boost | TDP (cTDP / PL2) | iGPU (NPU) | Mgmt | GB6 ST / MT | Found in (1 L) |
|---|---|---|---|---|---|---|---|---|---|
| Pentium G4400T | Skylake (6th) | T | 2/2 | 2.9 / — | 35 W (25) | HD 510 | — | 599 / 1058 | M700, M900, M710q |
| Pentium G4500T | Skylake (6th) | T | 2/2 | 3.0 / — | 35 W (25) | HD 530 | — | 701 / 1260 | M700, M900 |
| Core i3-6100T | Skylake (6th) | T | 2/4 | 3.2 / — | 35 W (25) | HD 530 | — | 1041 / 2071 | M700, M900, M710q |
| Core i3-6300T | Skylake (6th) | T | 2/4 | 3.3 / — | 35 W (25) | HD 530 | — | 1094 / 2327 | M700, M900 |
| Core i5-6400T | Skylake (6th) | T | 4/4 | 2.2 / 2.8 | 35 W (25) | HD 530 | — | 937 / 2552 | M700, M900 |
| Core i5-6500T | Skylake (6th) | T | 4/4 | 2.5 / 3.1 | 35 W (25) | HD 530 | vPro | 1042 / 2837 | M700, M900, M710q, M910q |
| Core i5-6600T | Skylake (6th) | T | 4/4 | 2.7 / 3.5 | 35 W (25) | HD 530 | vPro | 1168 / 3144 | M700, M900 |
| Core i7-6700T | Skylake (6th) | T | 4/8 | 2.8 / 3.6 | 35 W (25) | HD 530 | vPro | 1185 / 3613 | M700, M900, M710q, M910q |
| Pentium G4560T | Kaby Lake (7th) | T | 2/4 | 2.9 / — | 35 W (25) | HD 610 | — | no data | M710q |
| Pentium G4600T | Kaby Lake (7th) | T | 2/4 | 3.0 / — | 35 W (25) | HD 630 | — | no data | M710q |
| Core i3-7100T | Kaby Lake (7th) | T | 2/4 | 3.4 / — | 35 W (25) | HD 630 | — | no data | M710q, P320 |
| Core i5-7400T | Kaby Lake (7th) | T | 4/4 | 2.4 / 3.0 | 35 W (25) | HD 630 | — | no data / 2463 | M710q |
| Core i5-7500T | Kaby Lake (7th) | T | 4/4 | 2.7 / 3.3 | 35 W (25) | HD 630 | vPro | 1105 / 3056 | M710q, M910q, P320 |
| Core i5-7600T | Kaby Lake (7th) | T | 4/4 | 2.8 / 3.7 | 35 W (25) | HD 630 | vPro | 1235 / 3404 | M710q, M910q |
| Core i7-7700T | Kaby Lake (7th) | T | 4/8 | 2.9 / 3.8 | 35 W (25) | HD 630 | vPro | 1261 / 3643 | M710q, M910q, P320 |
| Celeron G4900T | Coffee Lake (8th) | T | 2/2 | 2.9 / — | 35 W (25) | UHD 610 | — | no data | M720q, M920q |
| Pentium Gold G5400T | Coffee Lake (8th) | T | 2/4 | 3.1 / — | 35 W (25) | UHD 610 | — | no data | M720q, M920q |
| Pentium Gold G5500T | Coffee Lake (8th) | T | 2/4 | 3.2 / — | 35 W (25) | UHD 630 | — | no data | M720q, M920q |
| Pentium Gold G5600T | Coffee Lake (8th) | T | 2/4 | 3.3 / — | 35 W (25) | UHD 630 | — | no data | M720q, M920q |
| Core i3-8100T | Coffee Lake (8th) | T | 4/4 | 3.1 / — | 35 W (25) | UHD 630 | — | no data | M720q, M920q, P330-1 |
| Core i3-8300T | Coffee Lake (8th) | T | 4/4 | 3.2 / — | 35 W (25) | UHD 630 | — | no data | M720q, M920q |
| Core i5-8400T | Coffee Lake (8th) | T | 6/6 | 1.7 / 3.3 | 35 W (25) | UHD 630 | — | no data | M720q, M920q, P330-1 |
| **Core i5-8500T** | Coffee Lake (8th) | T | 6/6 | 2.1 / 3.5 | 35 W (25) | UHD 630 | vPro | 1278 / 4480 | M720q, M920q, P330-1 |
| Core i5-8600T | Coffee Lake (8th) | T | 6/6 | 2.3 / 3.7 | 35 W (25) | UHD 630 | vPro | no data | M720q, M920q |
| **Core i7-8700T** | Coffee Lake (8th) | T | 6/12 | 2.4 / 4.0 | 35 W | UHD 630 | vPro | 1337 / 5578 | M720q, M920q, P330-1 |
| Celeron G4930T | Coffee Lake-R (9th) | T | 2/2 | 3.0 / — | 35 W (25) | UHD 610 | — | no data | M720q, M920q |
| Pentium Gold G5420T | Coffee Lake-R (9th) | T | 2/4 | 3.2 / — | 35 W (25) | UHD 610 | — | no data | M720q, M920q |
| Core i3-9100T | Coffee Lake-R (9th) | T | 4/4 | 3.1 / 3.7 | 35 W (25) | UHD 630 | — | 1245 / 3425 | M720q, M920q, P330-2 |
| Core i3-9300T | Coffee Lake-R (9th) | T | 4/4 | 3.2 / 3.8 | 35 W (25) | UHD 630 | — | no data | M720q, M920q |
| Core i5-9400T | Coffee Lake-R (9th) | T | 6/6 | 1.8 / 3.4 | 35 W (25) | UHD 630 | — | no data | M720q, M920q, P330-2 |
| Core i5-9500T | Coffee Lake-R (9th) | T | 6/6 | 2.2 / 3.7 | 35 W (25) | UHD 630 | vPro | 1157 / 4625 | M720q, M920q, P330-2 |
| Core i5-9600T | Coffee Lake-R (9th) | T | 6/6 | 2.3 / 3.9 | 35 W (25) | UHD 630 | vPro | 1265 / 4561 | M720q, M920q |
| **Core i7-9700T** | Coffee Lake-R (9th) | T | 8/8 | 2.0 / 4.3 | 35 W | UHD 630 | vPro | ~1450 / ~5700 | M720q, M920q, P330-2 |
| **Core i9-9900T** | Coffee Lake-R (9th) | T | 8/16 | 2.1 / 4.4 | 35 W | UHD 630 | vPro | 1447 / 6315 | M720q, M920q, P330-2 |
| Celeron G5900T | Comet Lake (10th) | T | 2/2 | 3.2 / — | 35 W (25) | UHD 610 | — | no data | M70q1, M80q1, M90q1 |
| Celeron G5905T | Comet Lake (10th) | T | 2/2 | 3.3 / — | 35 W (25) | UHD 610 | — | 721 / 1279 | M70q1-2, M80q1, M90q1-2 |
| Pentium G6400T | Comet Lake (10th) | T | 2/4 | 3.4 / — | 35 W (25) | UHD 610 | — | 734 / 1515 | M70q1, M80q1, M90q1 |
| Pentium G6405T | Comet Lake (10th) | T | 2/4 | 3.5 / — | 35 W (25) | UHD 610 | — | no data | M70q2 |
| Pentium G6500T | Comet Lake (10th) | T | 2/4 | 3.5 / — | 35 W (25) | UHD 630 | — | 762 / 1680 | M70q1, M80q1, M90q1 |
| Pentium G6505T | Comet Lake (10th) | T | 2/4 | 3.6 / — | 35 W (25) | UHD 630 | — | 788 / 1745 | M70q2, M90q2 |
| Core i3-10100T | Comet Lake (10th) | T | 4/8 | 3.0 / 3.8 | 35 W (25) | UHD 630 | — | ~1300 / ~4050 | M70q1, M80q1, M90q1, P340 |
| Core i3-10105T | Comet Lake (10th) | T | 4/8 | 3.0 / 3.9 | 35 W (25) | UHD 630 | — | ~1320 / ~4100 | M70q2, M90q2 |
| Core i3-10300T | Comet Lake (10th) | T | 4/8 | 3.0 / 3.9 | 35 W (25) | UHD 630 | — | ~1340 / ~4290 | M70q1, M80q1, M90q1 |
| Core i3-10305T | Comet Lake (10th) | T | 4/8 | 3.0 / 4.0 | 35 W (25) | UHD 630 | — | ~1320 / ~4400 | M70q2, M90q2 |
| Core i5-10400T | Comet Lake (10th) | T | 6/12 | 2.0 / 3.6 | 35 W (25) | UHD 630 | — | ~1200 / ~4400 | M70q1-2, M80q1, M90q1, P340 |
| **Core i5-10500T** | Comet Lake (10th) | T | 6/12 | 2.3 / 3.8 | 35 W (25) | UHD 630 | vPro | ~1290 / ~4650 | M70q1-2, M80q1, M90q1, P340 |
| Core i5-10600T | Comet Lake (10th) | T | 6/12 | 2.4 / 4.0 | 35 W (25) | UHD 630 | vPro | ~1353 / ~4820 | M70q1, M80q1, M90q1 |
| **Core i7-10700T** | Comet Lake (10th) | T | 8/16 | 2.0 / 4.5 | 35 W (25) | UHD 630 | vPro | ~1450 / ~6000 | M70q1, M80q1, M90q1, P340 |
| **Core i9-10900T** | Comet Lake (10th) | T | 10/20 | 1.9 / 4.6 | 35 W (25) | UHD 630 | vPro | ~1560 / ~7000 | M70q1, M80q1, M90q1, P340 |
| Core i3-10100 | Comet Lake (10th) | non-T | 4/8 | 3.6 / 4.3 | 65 W | UHD 630 | — | 1423 / 4224 | M90q1 |
| Core i3-10105 | Comet Lake (10th) | non-T | 4/8 | 3.7 / 4.4 | 65 W | UHD 630 | — | 1470 / 4896 | M90q2 |
| Core i3-10300 | Comet Lake (10th) | non-T | 4/8 | 3.7 / 4.4 | 65 W | UHD 630 | — | 1530 / 5098 | M90q1 |
| Core i3-10320 | Comet Lake (10th) | non-T | 4/8 | 3.8 / 4.6 | 65 W | UHD 630 | — | 1555 / 5024 | M90q1 |
| Core i3-10325 | Comet Lake (10th) | non-T | 4/8 | 3.9 / 4.7 | 65 W | UHD 630 | — | 1582 / 5245 | M90q2 |
| Core i5-10400 | Comet Lake (10th) | non-T | 6/12 | 2.9 / 4.3 | 65 W | UHD 630 | — | 1422 / 5460 | M90q1 |
| Core i5-10500 | Comet Lake (10th) | non-T | 6/12 | 3.1 / 4.5 | 65 W | UHD 630 | vPro | 1517 / 5430 | M90q1 |
| Core i5-10600 | Comet Lake (10th) | non-T | 6/12 | 3.3 / 4.8 | 65 W | UHD 630 | vPro | 1577 / 6155 | M90q1 |
| Core i7-10700 | Comet Lake (10th) | non-T | 8/16 | 2.9 / 4.8 | 65 W | UHD 630 | vPro | 1567 / 6946 | M90q1 |
| Core i9-10900 | Comet Lake (10th) | non-T | 10/20 | 2.8 / 5.2 | 65 W | UHD 630 | vPro | 1690 / 8379 | M90q1 |
| Core i5-11400T | Rocket Lake (11th) | T | 6/12 | 1.3 / 3.7 | 35 W (25 / PL2 84) | UHD 730 | — | ~1650 / ~5800 | M70q2, M90q2, P350 |
| Core i5-11500T | Rocket Lake (11th) | T | 6/12 | 1.5 / 3.9 | 35 W (25 / PL2 84) | UHD 750 | vPro | ~1750 / ~6000 | M70q2, M90q2, P350 |
| Core i5-11600T | Rocket Lake (11th) | T | 6/12 | 1.7 / 4.1 | 35 W (25 / PL2 84) | UHD 750 | vPro | ~1840 / ~6600 | M70q2, M90q2 |
| Core i7-11700T | Rocket Lake (11th) | T | 8/16 | 1.4 / 4.6 | 35 W (25 / PL2 115) | UHD 750 | vPro | ~1990 / ~6600 | M70q2, M90q2, P350 |
| **Core i9-11900T** | Rocket Lake (11th) | T | 8/16 | 1.5 / 4.9 | 35 W (25 / PL2 115) | UHD 750 | vPro | ~2130 / ~8800 | M70q2, M90q2, P350 |
| Core i5-11400 | Rocket Lake (11th) | non-T | 6/12 | 2.6 / 4.4 | 65 W (154) | UHD 730 | — | 1978 / 7895 | M90q2 |
| Core i5-11500 | Rocket Lake (11th) | non-T | 6/12 | 2.7 / 4.6 | 65 W (154) | UHD 750 | vPro | 2042 / 8215 | M90q2 |
| Core i5-11600 | Rocket Lake (11th) | non-T | 6/12 | 2.8 / 4.8 | 65 W (154) | UHD 750 | vPro | 2127 / 7996 | M90q2 |
| Core i7-11700 | Rocket Lake (11th) | non-T | 8/16 | 2.5 / 4.9 | 65 W (224) | UHD 750 | vPro | 2196 / 9387 | M90q2 |
| Core i9-11900 | Rocket Lake (11th) | non-T | 8/16 | 2.5 / 5.2 | 65 W (224) | UHD 750 | vPro | 2148 / 10040 | M90q2 |
| Celeron G6900T | Alder Lake (12th) | T | 2/2 (2P) | 2.8 / — | 35 W | UHD 710 | — | 1325 / 2306 | M70q3, M80q3, M90q3 |
| Pentium G7400T | Alder Lake (12th) | T | 2/4 (2P) | 3.1 / — | 35 W | UHD 710 | — | no data | M70q3, M80q3, M90q3 |
| Core i3-12100T | Alder Lake (12th) | T | 4/8 (4P) | 2.2 / 4.1 | 35 W (PL2 69) | UHD 730 | — | 2102 / 6532 | M70q3, M80q3, M90q3, P360 |
| Core i3-12300T | Alder Lake (12th) | T | 4/8 (4P) | 2.3 / 4.2 | 35 W (PL2 69) | UHD 730 | — | 2119 / 7086 | M70q3, M80q3, M90q3 |
| Core i5-12400T | Alder Lake (12th) | T | 6/12 (6P) | 1.8 / 4.2 | 35 W (PL2 74) | UHD 730 | — | 2041 / 6750 | M70q3, M80q3, M90q3, P360 |
| **Core i5-12500T** | Alder Lake (12th) | T | 6/12 (6P) | 2.0 / 4.4 | 35 W (PL2 74) | UHD 770 | vPro Ent | ~2335 / ~8850 | M70q3, M80q3, M90q3, P360 |
| Core i5-12600T | Alder Lake (12th) | T | 6/12 (6P) | 2.1 / 4.6 | 35 W (PL2 74) | UHD 770 | vPro Ent | 2312 / 8839 | M70q3, M80q3, M90q3 |
| **Core i7-12700T** | Alder Lake (12th) | T | 12/20 (8P+4E) | 1.4 / 4.7 | 35 W (PL2 99) | UHD 770 | vPro Ent | ~2330 / ~10900 | M70q3, M80q3, M90q3, P360 |
| **Core i9-12900T** | Alder Lake (12th) | T | 16/24 (8P+8E) | 1.4 / 4.9 | 35 W (PL2 106) | UHD 770 | vPro Ent | ~2430 / ~12800 | M70q3, M80q3, M90q3, P360 |
| Core i3-12100 | Alder Lake (12th) | non-T | 4/8 (4P) | 3.3 / 4.3 | 60 W (89) | UHD 730 | — | 2157 / 7428 | M90q3 |
| Core i3-12300 | Alder Lake (12th) | non-T | 4/8 (4P) | 3.5 / 4.4 | 60 W (89) | UHD 730 | — | 2235 / 8288 | M90q3 |
| Core i5-12400 | Alder Lake (12th) | non-T | 6/12 (6P) | 2.5 / 4.4 | 65 W (117) | UHD 730 | — | 2181 / 8820 | M90q3 |
| Core i5-12500 | Alder Lake (12th) | non-T | 6/12 (6P) | 3.0 / 4.6 | 65 W (117) | UHD 770 | vPro Ent | 2365 / 10401 | M90q3 |
| Core i5-12600 | Alder Lake (12th) | non-T | 6/12 (6P) | 3.3 / 4.8 | 65 W (117) | UHD 770 | vPro Ent | 2386 / 10379 | M90q3 |
| Core i7-12700 | Alder Lake (12th) | non-T | 12/20 (8P+4E) | 2.1 / 4.9 | 65 W (180) | UHD 770 | vPro Ent | 2497 / 12448 | M90q3 |
| Core i9-12900 | Alder Lake (12th) | non-T | 16/24 (8P+8E) | 2.4 / 5.1 | 65 W (202) | UHD 770 | vPro Ent | 2637 / 15367 | M90q3 |
| **Intel 300T** | Raptor Lake (13th) | T | 2/4 (2P) | 3.4 / — | 35 W | UHD 710 | — | no data | M70q5 |
| Core i3-13100T | Raptor Lake (13th) | T | 4/8 (4P) | 2.5 / 4.2 | 35 W (PL2 69) | UHD 730 | — | 2134 / 6604 | M70q4-5, M80q4, M90q4 |
| Core i5-13400T | Raptor Lake (13th) | T | 10/16 (6P+4E) | 1.3 / 4.4 | 35 W (PL2 82) | UHD 730 | vPro Ess | no data | M70q4-5, M80q4, M90q4 |
| Core i5-13500T | Raptor Lake (13th) | T | 14/20 (6P+8E) | 1.6 / 4.6 | 35 W (PL2 92) | UHD 770 | vPro Ent | 2375 / no data* | M70q4-5, M80q4, M90q4-5, P3T |
| Core i5-13600T | Raptor Lake (13th) | T | 14/20 (6P+8E) | 1.8 / 4.8 | 35 W (PL2 92) | UHD 770 | vPro Ent | ~2378 / ~10830 | M70q4, M80q4, M90q4-5, P3T |
| **Core i7-13700T** | Raptor Lake (13th) | T | 16/24 (8P+8E) | 1.4 / 4.9 | 35 W (PL2 106) | UHD 770 | vPro Ent | ~2520 / ~13400 | M70q4-5, M80q4, M90q4-5, P3T |
| **Core i9-13900T** | Raptor Lake (13th) | T | 24/32 (8P+16E) | 1.1 / 5.3 | 35 W (PL2 106) | UHD 770 | vPro Ent | ~2618 / ~14600* | M80q4, M90q4-5, P3T — **not** in M70q4 |
| Core i5-13500 | Raptor Lake (13th) | non-T | 14/20 (6P+8E) | 2.5 / 4.8 | 65 W (154) | UHD 770 | vPro Ent | 2451 / 12573 | M90q4-5 |
| Core i5-13600 | Raptor Lake (13th) | non-T | 14/20 (6P+8E) | 2.7 / 5.0 | 65 W (154) | UHD 770 | vPro Ent | 2488 / 14965 | M90q4-5 |
| Core i7-13700 | Raptor Lake (13th) | non-T | 16/24 (8P+8E) | 2.1 / 5.2 | 65 W (219) | UHD 770 | vPro Ent | 2675 / 16079 | M90q4-5 |
| Core i9-13900 | Raptor Lake (13th) | non-T | 24/32 (8P+16E) | 2.0 / 5.6 | 65 W (219) | UHD 770 | vPro Ent | 2876 / 19752 | M90q4-5 |
| Core i3-14100T | Raptor Lake-R (14th) | T | 4/8 (4P) | 2.7 / 4.4 | 35 W (PL2 69) | UHD 730 | — | no data | M70q5 |
| Core i5-14400T | Raptor Lake-R (14th) | T | 10/16 (6P+4E) | 1.5 / 4.5 | 35 W (PL2 82) | UHD 730 | — | no data | M70q5 |
| Core i5-14500T | Raptor Lake-R (14th) | T | 14/20 (6P+8E) | 1.7 / 4.8 | 35 W (PL2 92) | UHD 770 | vPro Ent | 2561 / 12754 | M70q5, M90q5 |
| Core i5-14600T | Raptor Lake-R (14th) | T | 14/20 (6P+8E) | 1.8 / 5.1 | 35 W (PL2 92) | UHD 770 | vPro Ent | 2626 / 15354 | M70q5, M90q5 |
| **Core i7-14700T** | Raptor Lake-R (14th) | T | 20/28 (8P+12E) | 1.3 / 5.2 | 35 W (PL2 106) | UHD 770 | vPro Ent | 2723 / 18390 | M70q5, M90q5 |
| **Core i9-14900T** | Raptor Lake-R (14th) | T | 24/32 (8P+16E) | 1.1 / 5.5 | 35 W (PL2 106) | UHD 770 | vPro Ent | 2823 / ~16750* | M90q5 |
| Core i5-14500 | Raptor Lake-R (14th) | non-T | 14/20 (6P+8E) | 2.6 / 5.0 | 65 W (154) | UHD 770 | vPro Ent | 2611 / 15155 | M90q5 |
| Core i5-14600 | Raptor Lake-R (14th) | non-T | 14/20 (6P+8E) | 2.7 / 5.2 | 65 W (154) | UHD 770 | vPro Ent | 2688 / 16101 | M90q5 |
| Core i7-14700 | Raptor Lake-R (14th) | non-T | 20/28 (8P+12E) | 2.1 / 5.4 | 65 W (219) | UHD 770 | vPro Ent | 2829 / 18874 | M90q5 |
| **Core i9-14900** | Raptor Lake-R (14th) | non-T | 24/32 (8P+16E) | 2.0 / 5.8 | 65 W (219) | UHD 770 | vPro Ent | 2947 / 20495 | M90q5 — **fastest pre-Ultra chip in the range** |
| Core Ultra 5 225 | Arrow Lake-S (Ultra S2) | non-T | 10/10 (6P+4E) | 3.3 / 2.7 → 4.9 / 4.4 | 65 W | Intel Gfx ~4 TOPS (**NPU 13**) | — | no data | M90q6, P3T2 |
| Core Ultra 5 225T | Arrow Lake-S (Ultra S2) | T | 10/10 (6P+4E) | 2.5 / 1.9 → 4.9 / 4.4 | 35 W | Intel Gfx ~4 TOPS (**NPU 13**) | — | no data | P3T2 |
| Core Ultra 5 235 | Arrow Lake-S (Ultra S2) | non-T | 14/14 (6P+8E) | 3.4 / 2.9 → 5.0 / 4.4 | 65 W | Intel Gfx ~6 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| Core Ultra 5 235A | Arrow Lake-S (Ultra S2) | non-T | 14/14 (6P+8E) | 3.4 / 2.9 → 5.0 / 4.4 | 65 W | Intel Gfx ~6 TOPS (**NPU 13**) | vPro Ent | no data | P3T2 |
| Core Ultra 5 235T | Arrow Lake-S (Ultra S2) | T | 14/14 (6P+8E) | 2.2 / 1.6 → 5.0 / 4.4 | 35 W | Intel Gfx ~6 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| Core Ultra 5 235TA | Arrow Lake-S (Ultra S2) | T | 14/14 (6P+8E) | 2.2 / 1.6 → 5.0 / 4.4 | 35 W | Intel Gfx ~6 TOPS (**NPU 13**) | vPro Ent | no data | P3T2 |
| Core Ultra 5 245 | Arrow Lake-S (Ultra S2) | non-T | 14/14 (6P+8E) | 3.5 / 3.0 → 5.1 / 4.5 | 65 W | Intel Gfx ~8 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| Core Ultra 5 245T | Arrow Lake-S (Ultra S2) | T | 14/14 (6P+8E) | 2.2 / 1.7 → 5.1 / 4.5 | 35 W | Intel Gfx ~8 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| **Core Ultra 7 265** | Arrow Lake-S (Ultra S2) | non-T | 20/20 (8P+12E) | 2.4 / 1.8 → 5.3 max | 65 W | Intel Gfx ~8 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| **Core Ultra 7 265T** | Arrow Lake-S (Ultra S2) | T | 20/20 (8P+12E) | 1.5 / 1.2 → 5.3 max | 35 W | Intel Gfx ~8 TOPS (**NPU 13**) | vPro Ent | 2954 / 16455 | M90q6, P3T2 |
| **Core Ultra 9 285** | Arrow Lake-S (Ultra S2) | non-T | 24/24 (8P+16E) | 2.5 / 1.9 → 5.6 max | 65 W | Intel Gfx ~8 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| **Core Ultra 9 285T** | Arrow Lake-S (Ultra S2) | T | 24/24 (8P+16E) | 1.4 / 1.2 → 5.4 max | 35 W | Intel Gfx ~8 TOPS (**NPU 13**) | vPro Ent | no data | M90q6, P3T2 |
| Core i5-13420H | Raptor Lake-H (13th, BGA) | H | 8/12 (4P+4E) | 2.1 / 4.6 | 45 W (PL2 95) | UHD (48 EU) | — | no data | neo50q4 |
| Core i7-13620H | Raptor Lake-H (13th, BGA) | H | 10/16 (6P+4E) | 2.4 / 4.9 | 45 W (PL2 115) | UHD (64 EU) | — | no data | neo50q4 |
| Core 5 210H | Raptor Lake-H ("Series 2", BGA) | H | 8/12 (4P+4E) | 2.2 / 4.8 | 45 W (35–45 / 115) | UHD Xe G4 (48 EU) | — | no data | neo50q5 |
| **Core 7 240H** | Raptor Lake-H ("Series 2", BGA) | H | 10/16 (6P+4E) | 2.5 / 5.2 | 45 W (35–45 / 115) | Intel Gfx (64 EU) | — | ~2632 / ~12604 | neo50q5 |
| AMD A9-9420e | Excavator (Stoney Ridge) | GE | 2/2 | 1.8 / 2.7 | 6 W | Radeon R5 | — | no data | M625q |
| AMD PRO A6-8570E | Excavator (Bristol Ridge) | GE | 2/2 | 3.0 / 3.4 | 35 W | Radeon R5 | DASH | no data | M715q1, M715q2 |
| AMD PRO A6-9500E | Excavator (Bristol Ridge) | GE | 2/2 | 3.0 / 3.4 | 35 W | Radeon R5 | DASH | 611 / 908 | M715q1, M715q2 |
| AMD PRO A10-8770E | Excavator (Bristol Ridge) | GE | 4/4 | 2.8 / 3.5 | 35 W | Radeon R7 | DASH | no data | M715q1, M715q2 |
| AMD PRO A10-9700E | Excavator (Bristol Ridge) | GE | 4/4 | 3.0 / 3.5 | 35 W | Radeon R7 | DASH | no data | M715q1, M715q2 |
| AMD PRO A12-8870E | Excavator (Bristol Ridge) | GE | 4/4 | 2.9 / 3.8 | 35 W | Radeon R7 | DASH | no data | M715q1, M715q2 |
| AMD PRO A12-9800E | Excavator (Bristol Ridge) | GE | 4/4 | 3.1 / 3.8 | 35 W | Radeon R7 | DASH | 646 / 1552 | M715q1, M715q2 |
| AMD Athlon 200GE | Zen (Raven Ridge) | GE | 2/4 | 3.2 / — | 35 W | Vega 3 | — | 915 / 1894 | M715q2 |
| AMD Athlon PRO 200GE | Zen (Raven Ridge) | GE | 2/4 | 3.2 / — | 35 W | Vega 3 | DASH | 932 / 1900 | M715q2, M75q1 |
| AMD Ryzen 3 2200GE | Zen (Raven Ridge) | GE | 4/4 | 3.2 / 3.6 | 35 W | Vega 8 | — | no data | M715q2 |
| AMD Ryzen 3 PRO 2200GE | Zen (Raven Ridge) | GE | 4/4 | 3.2 / 3.6 | 35 W | Vega 8 | DASH | no data | M715q2, M75q1 |
| AMD Ryzen 5 2400GE | Zen (Raven Ridge) | GE | 4/8 | 3.2 / 3.8 | 35 W | Vega 11 | — | no data | M715q2 |
| AMD Ryzen 5 PRO 2400GE | Zen (Raven Ridge) | GE | 4/8 | 3.2 / 3.8 | 35 W | Vega 11 | DASH | no data | M715q2, M75q1 |
| AMD Athlon PRO 300GE | Zen+ (Picasso) | GE | 2/4 | 3.4 / — | 35 W (cTDP 12–35) | Vega 3 | DASH | 934 / 1892 | M75q1 |
| AMD Ryzen 3 PRO 3200GE | Zen+ (Picasso) | GE | 4/4 | 3.3 / 3.8 | 35 W (cTDP 12–35) | Vega 8 | DASH | 1049 / 2952 | M75q1 |
| **AMD Ryzen 5 PRO 3400GE** | Zen+ (Picasso) | GE | 4/8 | 3.3 / 3.9 | 35 W (cTDP 12–35) | Vega 11 | DASH | 1092 / 3290 | M75q1 — **Gen 1 flagship, NOT a Gen 2 chip** |
| AMD Ryzen 3 PRO 4350GE | Renoir (Zen 2) | GE | 4/8 | 3.5 / 4.0 | 35 W (cTDP 12–35) | Vega 6 | DASH | 1446 / 4606 | M75q2 |
| **AMD Ryzen 5 PRO 4650GE** | Renoir (Zen 2) | GE | 6/12 | 3.3 / 4.2 | 35 W (cTDP 12–35) | Vega 7 | DASH | 1470 / 5595 | M75q2 |
| **AMD Ryzen 7 PRO 4750GE** | Renoir (Zen 2) | GE | 8/16 | 3.1 / 4.3 | 35 W (cTDP 12–35) | Vega 8 | DASH | 1612 / 6466 | M75q2 |
| AMD Ryzen 3 PRO 5350GE | Cezanne (Zen 3) | GE | 4/8 | 3.6 / 4.2 | 35 W (cTDP 12–35) | Vega 6 | DASH | 1736 / 5064 | M75q2 |
| **AMD Ryzen 5 PRO 5650GE** | Cezanne (Zen 3) | GE | 6/12 | 3.4 / 4.4 | 35 W (cTDP 12–35) | Vega 7 | DASH | 1938 / 7053 | M75q2 |
| **AMD Ryzen 7 PRO 5750GE** | Cezanne (Zen 3) | GE | 8/16 | 3.2 / 4.6 | 35 W (cTDP 12–35) | Vega 8 | DASH | 1910 / 6996 | M75q2 |
| AMD Ryzen 3 8300GE | Zen 4 Phoenix2 (1×Zen4 + 3×Zen4c) | GE | 4/8 | 3.5 / 4.9 | 35 W | Radeon 740M (**no NPU**) | — | no data | M75q5 |
| AMD Ryzen 3 PRO 8300GE | Zen 4 Phoenix2 | GE | 4/8 | 3.5 / 4.9 | 35 W | Radeon 740M (**no NPU**) | DASH | 2410 / 6410 | M75q5 |
| AMD Ryzen 5 8500GE | Zen 4 Phoenix2 (2×Zen4 + 4×Zen4c) | GE | 6/12 | 3.4 / 5.0 | 35 W | Radeon 740M (**no NPU**) | — | no data | M75q5 |
| **AMD Ryzen 5 PRO 8500GE** | Zen 4 Phoenix2 | GE | 6/12 | 3.4 / 5.0 | 35 W | Radeon 740M (**no NPU**) | DASH | 2661 / 9732 | M75q5 |
| AMD Ryzen 5 8505GE | Zen 4 Phoenix2 | GE | 6/12 | 3.4 / 5.0 | 35 W | Radeon 740M (**no NPU**) | — | no data | M75q5 (DDR5-5200 variant) |
| **AMD Ryzen 5 PRO 8600GE** | Zen 4 Phoenix (monolithic) | GE | 6/12 | 3.9 / 5.0 | 35 W | Radeon 760M (**NPU ~16 TOPS**) | DASH | 2627 / 9145 | M75q5 |
| **AMD Ryzen 7 PRO 8700GE** | Zen 4 Phoenix (monolithic) | GE | 8/16 | 3.6 / 5.1 | 35 W | Radeon 780M (**NPU ~16 TOPS**) | DASH | 2676 / 11969 | M75q5 — **35 W flagship** |
| AMD Ryzen 7 PRO 8705GE | Zen 4 Phoenix (monolithic) | GE | 8/16 | 3.6 / 5.1 | 35 W | Radeon 780M (**NPU**) | DASH | no data | M75q5 (DDR5-5200 variant) |
| AMD Ryzen 3 8300G / PRO 8300G | Zen 4 Phoenix2 | G | 4/8 | 3.4 / 4.9 | 65 W (45–65) | Radeon 740M (no NPU) | DASH (PRO) | ~2410 / ~7600 | M75q5 — throttles in 1 L |
| AMD Ryzen 5 8500G / PRO 8500G | Zen 4 Phoenix2 | G | 6/12 | 3.5 / 5.0 | 65 W (45–65) | Radeon 740M (no NPU) | DASH (PRO) | ~2670 / ~10000 | M75q5 — throttles in 1 L |
| AMD Ryzen 7 8700G / PRO 8700G | Zen 4 Phoenix (monolithic) | G | 8/16 | 4.2 / 5.1 | 65 W (45–65) | Radeon 780M (**NPU**) | DASH (PRO) | 2720 / 14326 | M75q5 — use 8700GE as the 1 L proxy |
| Intel Celeron 4205U | Gemini Lake | U | 2/2 | 1.8 / — | 15 W | UHD 610 | — | no data | M90n-1 (IoT) |
| Intel Celeron 4305UE | Gemini Lake-R | U | 2/2 | 2.0 / — | 15 W | UHD 610 | — | no data | M90n-1 (IoT) |
| Intel Core i3-8145U | Whiskey Lake-U | U | 2/4 | 2.1 / 3.9 | 15 W | UHD 620 | — | 1218 / 2338 | M90n-1 |
| Intel Core i3-8145UE | Whiskey Lake-U | U | 2/4 | 2.2 / 3.9 | 15 W | UHD 620 | — | no data | M90n-1 (IoT) |
| Intel Core i5-8265U | Whiskey Lake-U | U | 4/8 | 1.6 / 3.9 | 15 W | UHD 620 | — | 1286 / 3267 | M90n-1 |
| Intel Core i5-8365U | Whiskey Lake-U | U | 4/8 | 1.6 / 4.1 | 15 W | UHD 620 | **vPro** | 1309 / 3335 | M90n-1 |
| **Intel Core i7-8665U** | Whiskey Lake-U | U | 4/8 | 1.9 / 4.8 | 15 W | UHD 620 | **vPro** | 1437 / 3732 | M90n-1 — top Intel Nano |
| AMD Ryzen 3 PRO 3300U | Zen+ (Picasso) | U | 4/4 | 2.1 / 3.5 | 15 W (cTDP 12–35) | Vega 6 | DASH | 971 / 2577 | M75n |
| AMD Ryzen 5 PRO 3500U | Zen+ (Picasso) | U | 4/8 | 2.1 / 3.7 | 15 W (cTDP 12–35) | Vega 8 | DASH | no data | M75n — top AMD Nano |
| AMD Athlon Silver 3050e | Zen (Dali) | U | 2/4 | 1.4 / 2.8 | 6 W | Vega 3 | — | 758 / 1494 | M75n (IoT) |

### Long-tail SKUs not tabled individually

- **Intel 58 W `non-T`, M90q Gen 1/2 only, mostly `no data` on GB6:** Pentium G6400 / G6500 / G6600 / G6605,
  Celeron G5900 / G5905 / G5920 / G5925.
- **AMD M75q Gen 2 non-PRO and PRO-refresh:** Renoir Ryzen 3 4300GE / 5 4600GE / 7 4700GE; Cezanne Ryzen 3
  5300GE / 5 5600GE / 7 5700GE; PRO-refresh Cezanne **5355GE / 5655GE / 5755GE** (≈ the 535x/565x/575xGE
  rows above); Picasso carryovers Ryzen 3 3200GE, Ryzen 5 PRO 3350GE, Athlon Gold PRO 3150GE, Athlon
  Silver 3050GE / PRO 3125GE. **Verify the exact SKU on PSREF before quoting specs.**
- **Pre-Skylake (M72e / M73 / M83 / M93p / M53 / M600):** not enumerated. Ivy Bridge and Haswell `T` parts
  (G1610T, G2020T, i3-3220T, i5-3470T, G1820T, G3220T, i3-4130T, i5-4570T, i5-4590T, i7-4765T, i7-4785T)
  plus Bay Trail-D and Braswell Celeron/Pentium. **Confirm on PSREF — these were not verified this session.**

### Silicon capability cheat-sheet (what the generation actually buys you)

| Capability | First available in this roster | Why it matters in a rack |
|---|---|---|
| **AES-NI** | all (Skylake onward) | LUKS, TLS, VPN throughput |
| **AVX2** | Skylake+ / Zen+ (**not** Bay Trail / Braswell) | many modern prebuilt binaries assume it |
| **SHA-NI** | **Zen (2018) on AMD; Ice Lake/Alder Lake on Intel** | ZFS SHA-256 checksums, TLS handshakes — a real AMD advantage on 8th/9th-gen-era comparisons |
| **AVX-512** | **Rocket Lake only** (i5/i7/i9-11xxx, `T` and non-T) | fused off from Alder Lake onward. If a workload wants it, 11th gen is the *only* option in a Tiny |
| **QuickSync H.264 + HEVC 8-bit** | HD 530 (Skylake) | baseline Jellyfin/Plex transcode |
| **HEVC 10-bit + VP9 decode** | UHD 630 (Coffee Lake) | the M720q/M920q transcode sweet spot |
| **AV1 decode** | UHD 730/770 (Alder Lake, 12th gen); Radeon 780M; NVIDIA Ampere (RTX A400/A1000) | future-proofing a media library |
| **AV1 encode** | **Intel Arc A310** (M90q Gen 6 option); **Radeon 780M / RDNA 3** (M75q Gen 5) | streaming/re-encode pipelines |
| **NVENC (no session cap)** | Quadro/NVIDIA-pro cards in the P-series Tiny | many-simultaneous-stream transcoding |
| **NPU** | **Ryzen AI ~16 TOPS** (8600GE/8700GE, M75q Gen 5); **Intel AI Boost 13 TOPS** (Core Ultra, M90q Gen 6 / P3 Tiny Gen 2) | local inference without a Coral/Hailo stick |
| **PCIe 4.0 NVMe** | M90q Gen 3 / M70q Gen 5 onward | real NVMe throughput; Gen 2 and earlier are PCIe 3.0 |
| **PCIe 5.0 NVMe** | **P3 Tiny Gen 2 only** (one of three slots) | one very fast scratch/DB drive |
| **DDR5** | M90q Gen 3 / P360 Tiny onward | bandwidth for iGPU inference and many-VM hosts |
| **ECC** | **never** | do not design an integrity story around a Tiny |

---

# Part 3 — 10-inch rack fit, mounting and hardware

## 3.1 The dimensional facts

| | Value | Consequence |
|---|---|---|
| Chassis W × D | **179 × 182.9 mm** (unchanged since ~2012) | Fits the ~222 mm usable interior of a 10" rack with ~43 mm to spare — enough for a printed shell wall plus one or two keystone jacks beside the PC |
| Height, bare | **34.5 mm** (most), **35–37 mm** (Tiny7/8) | 1U = 44.45 mm. A Tiny is a **true sub-1U device** — but only just |
| Height, with rubber feet | **36.5–37 mm** | **Most 10" mounts require removing the rubber feet.** Budget the difference before printing |
| Weight | 1.25–1.4 kg | Trivial for a rail-mounted tray; a front-only mount will sag without a rear support |
| Depth budget | 183 mm chassis + ~40–60 mm for the power barrel, Ethernet boot and DP/HDMI heads | **Plan for ≥250 mm of rack depth.** Right-angle barrel and RJ45 adapters buy back ~25 mm |
| Orientation | Horizontal (flat) in a tray, or vertical | **Vertical wastes 10" width** — two Tinies laid flat side-by-side do not fit (2 × 179 = 358 mm), so one Tiny per U flat is the practical density unless you go vertical, where two *can* stand side by side in ~2U |

**The density math nobody states out loud:** a Tiny occupies about one full U and roughly 80 % of the
10" width. Compared to a stack of SBCs (four Pi-class boards per U is routine), a rack of Tinies is
**low-density but high-capability-per-slot**. That's the trade: one M920q with 64 GB and a 10GbE NIC
replaces four ARM boards and one 1U of space, but you can't run seven of them in a 5U rack.

## 3.2 Mounting — what actually exists

The ThinkCentre Tiny is one of the best-served devices in the 10" ecosystem, because it has been the
same size for over a decade.

**Free / printable (all confirmed to exist as of 2026-07):**

| Design | Notes |
|---|---|
| **Tim — "Lenovo ThinkCentre Tiny 10" Rack Mount, now with optional Keystone Sockets"** (Printables 1040412) | The most-used design. **Up to two standard keystone jacks** on the faceplate to bring Ethernet/USB/HDMI to the rack front. Author recommends **PETG, ABS or ASA** — *"PLA works, but may get soft in enclosed racks (>50 °C)"* — and printing front-face-down. |
| **r3vo — "Unified 10" Rack – Lenovo Thinkcentre Tiny Mount"** (Printables 1215391) | Designed around M720q/M920q; **has a seat for the rubber feet**, so you don't have to remove them. Pairs with the author's separate 1U rear support bracket (Printables 1215562). |
| **owlish — 10" mount for M920q/M720q** (Printables 1384009) | Remix; simpler geometry. |
| **Emplar — "Lenovo Tiny 10 inch rack mount (yet another)"** (Printables 1041164) | Tested on **M900, M910q, M93p** — useful for the legacy tier. |
| **Hermann S — "Lenovo Tiny 10 inch Rackmount with Keystone slots"** (Printables 877502) | M720-class. |
| **Smelliot — M900-specific mount** (Printables 1341050) | Exists **because the M900's rubber feet are rotated 90°** and foul generic mounts. Locks to the chassis screw holes with 2× M3×6. ~130 g of filament. |
| **just_actual_kev — "10" Rackmount for ThinkCentre 1L"** (MakerWorld 1141511) | **Fully enclosed** design, 2 keystone openings with three keystone-orientation options. Claims fitment across M600, M710q, M720q, M900, M910q/x, M920q/x, M70q, M75q, M90q, P320, P330, P340, P350, P360, P3 Tiny. |

**Commercial:**

| Vendor | Notes |
|---|---|
| **MyElectronics.nl** | Sells injection-moulded 10" and 19" mounts for ThinkCentre Tiny; the reference commercial option in the mini-rack community. |
| **3drackmounts (eBay store)** | Sells both a **10" ThinkCentre M720q/M920q/M710/M10q mount (~$40)** and a 19" modular version; also does a **10" JetKVM mount** worth pairing with it. |
| **Hive Tech Solutions (hivets.au)** | 10" 1U mount, PC slides in from the rear with a rear retaining bracket. **Explicitly notes the rubber feet must be removed.** |
| **Etsy sellers** | Multiple PETG 1U mounts listing M720q/M715q/M920q/M910q/M600/M92p/P330 Tiny compatibility. |

**Practical mounting advice:**

1. **Print in PETG/ASA, not PLA.** An enclosed 10" rack with three or four Tinies will exceed 50 °C
   ambient in summer, and a PLA faceplate will creep and drop the machine.
2. **Use a rear support.** A front-faceplate-only mount cantilevers 1.3 kg off two rack screws. Either
   pick a design with a full tray/shell, or add a separate 1U rear support bracket.
3. **Use the keystone variants.** The Tiny's I/O is all on the rear; bringing one RJ45 and one USB
   keystone to the front turns a rack-rear cable dive into a front-panel operation. This is the single
   biggest quality-of-life win available for this form factor.
4. **Check the feet.** M900 feet are rotated; Tiny7/8 feet add 2 mm. Measure before printing.
5. **Leave the side vents clear.** Intake is on the side/bottom, exhaust is rear. A fully-enclosed printed
   shell that blocks the side intake will cook the machine — prefer perforated or open-sided designs, or
   verify the shell has matching vents.

## 3.3 The power brick problem — the real 10" rack headache

This is the thing that bites people, and it is worse than the mounting.

- Every Tiny uses an **external Lenovo "slim tip" barrel adapter** (the same family as ThinkPad chargers).
  Wattages across the roster: **65 W · 90 W · 135 W · 170 W · 230 W · 245 W · 300 W · 330 W** (the top
  three are P3 Tiny Gen 2 territory).
- **135 W is mandatory if the optional USB-C punch-out port is fitted** on M70q Gen 5 / M90q Gen 6.
  A 65 W or 90 W brick will not enumerate it.
- A brick is roughly the size of the machine's own footprint. **Four Tinies in a rack means four bricks
  to hide.** They have no mounting provision at all.

**Mitigations, in order of how well they work:**

| Approach | Notes |
|---|---|
| **Lenovo "power adapter cage"** | A factory option on M70q Gen 5, M90q Gen 6 and others. Straps the brick to the chassis. Cheap, official, and the least-fuss answer. |
| **Printed 1U brick shelf** | Several exist in the 10" ecosystem; put 2–4 bricks in one U behind or below the nodes. Costs a rack unit but keeps everything serviceable. |
| **Single DC bus + slim-tip pigtails** | A 20 V supply on DIN rail feeding slim-tip barrel pigtails. Removes N bricks. **Caveat: Lenovo slim-tip carries an ID pin**, and the BIOS will warn (and in some cases power-limit) on a non-Lenovo adapter. Test before committing a whole rack. |
| **USB-C PD in** | **M70q Gen 5 and M90q Gen 6 support 100 W PD-in on punch-out port 1.** If your rack already has a PD source, this is the cleanest modern answer — but it is a special-bid/optional port, so verify the specific SKU has it. |
| **PoE** | **Not available on any Tiny.** No PoE HAT, no PoE splitter path that Lenovo supports. If your rack's power design is PoE-first, Tinies are the exception you have to plan around. |

## 3.4 The PCIe riser — part numbers, fitment, and the hidden second link

Only some Tinies have a riser slot. Where present it is the difference between a compute node and an
appliance.

| Model family | Riser | Electrical |
|---|---|---|
| M910x / M920x / **M720q / M920q** / P330 Tiny | **01AJ940** (PCIe x16 connector), **01AJ929** (x4 variant) | **x8 from CPU** (01AJ940) or x4 from PCH (01AJ929) |
| P320 / P330 / P340 / P350 / P360 / P3 Tiny | factory GPU riser (**not cross-generation compatible** — documented failures moving a P330 riser into a P340) | PCIe 3.0 x8 (P360 era) |
| M90q Gen 1–5 | Lenovo riser | x8 |
| **M90q Gen 6** | **5C50W00933 / 5C50W00910** family | **PCIe 4.0 x8 low profile** |
| **P3 Tiny Gen 2** | same family | **PCIe 4.0 x16 connector, x8 link** for GPU, or **x4** for NIC/serial |
| M70q (any gen), M80q (any gen), M75q (any gen), neo 50q, Nano | **none, ever** | — |

**The hidden second link.** On Tiny5 boards the riser connector carries **two independent PCIe links**:
an **x8 from the CPU** (intended for the GPU SKUs) and an **x4 from the PCH** (intended for the
Thunderbolt/NIC SKUs). Lenovo's own risers expose one or the other. Two community risers expose both:

- **TinyRiser** (FairywrenTech, Tindie) — low-profile PCIe card **plus** an extra M.2 NVMe simultaneously.
- **PowerRiser** (NandFarm, open-source, on Lectronz/GitHub) — adds a 2230/2242 NVMe while preserving the
  x8 link through an **open-ended x8 connector**, so a physically-x16 card will seat (still x8
  electrically). Requires removing the front Bluetooth-antenna bracket.

Both are compatible with **M720q, M920q, M920x and ThinkStation P330 Tiny** only.

**Fitment gotchas:**
- The riser usually needs the **plastic baffle / air duct** to direct airflow over the card. Sellers list
  "with baffle" and "no baffle" variants — get the baffle.
- Half-height, half-length, **single-slot** cards only.
- Passively-cooled 10GbE NICs (Mellanox ConnectX-3, Intel X520/X710) run **hot** in a 1 L box with no
  dedicated airflow. Many people add a 30–40 mm fan on the heatsink. Plan for it.
- A NIC or GPU in the slot pushes you from a 65 W brick to 90 W or 135 W.

## 3.5 Thermals, noise and airflow

`nodes.md` §10 already says it: *the 1 L used tiny PCs are the best-behaved thermal citizens here — they
were engineered for a 35–65 W TDP in a sealed office box with a real blower.* That holds, with caveats:

- **Intake is side/bottom, exhaust is rear.** In a rack this is ideal — front-to-back is not what these
  do, but side-to-rear works fine with 10–15 mm of clearance and an open-sided mount.
- **Sustained load is audible.** ServeTheHome measured the M75q Gen 5 at **47–48 dBA** holding
  63–66 W over two minutes. That is a design choice (hold performance, accept noise) and it is the right
  one for a rack — but not for a rack in a bedroom.
- **The 65 W `non-T` and `G` SKUs are the noisy ones.** A 35 W `T` part in the same chassis is
  substantially quieter for similar sustained throughput.
- **Dust is the long-term enemy.** STH notes idle power creeping up as dust accumulates in the fan and
  heatsink. Lenovo sells an **optional dust filter** on modern SKUs; in a rack, take it. Otherwise plan
  an annual blow-out — the coolers come apart with a screwdriver on every generation.
- **ICE (Intelligent Cooling Engine)** BIOS modes let you trade acoustics against sustained clocks.
  On a rack node, "Better Thermal Performance" is usually the right setting; "Better Acoustic
  Performance" caps sustained turbo noticeably.
- **P-series (GPU) Tinies run hottest.** A 50 W GPU plus a 35 W CPU in 1 L is the thermal ceiling of the
  format. These need the most clearance and will be the loudest node in the rack.

## 3.6 Networking upgrade paths, ranked

1. **PCIe riser + Intel NIC** (M720q/M920q/M90q/P-series only) — 10GbE SFP+, 10GBase-T, or quad GbE.
   The only path that gives you a real multi-queue NIC. **Requires the riser + baffle + a bigger brick.**
2. **Factory options on modern SKUs** — 2.5GbE Realtek RTL8125BGS punch-out (M90q Gen 3+, M70q Gen 5);
   **Intel I350-T4 quad GbE** (M90q Gen 6, P3 Tiny Gen 2, P330 Tiny); Broadcom BCM5719/5720 (P3 Tiny Gen 2);
   two extra 2.5GbE via board-to-board on P3 Tiny Gen 2.
3. **Thunderbolt 4** punch-out (M90q Gen 6, P3 Tiny Gen 2, some P-series) — external NIC/GPU/NVMe enclosures.
4. **USB 3 NIC** — works everywhere, and STH explicitly suggests it for the riser-less AMD boxes. 2.5GbE
   USB adapters (RTL8156) are cheap and adequate for a second interface; do not expect line rate or low
   CPU overhead.
5. **Wi-Fi M.2 slot repurposing** — the WLAN slot is usually A/E-key and short; it is not a practical
   NIC or NVMe expansion path.

⚠ **For BSD-based firewalls (pfSense/OPNsense), prefer Intel silicon.** The Realtek RTL8125 2.5GbE parts
and the Realtek GbE on the AMD Tinies have historically been the weakest driver story on FreeBSD. On
Linux (VyOS, OpenWrt, plain nftables) this matters far less.

## 3.7 Out-of-band management and unattended behaviour

This is where Tinies quietly beat most SBCs, and it is worth stating plainly.

- **"After Power Loss" BIOS setting** — every Tiny has it. Set to **"Power On"** and the node comes back
  by itself after an outage. `nodes.md` §9 flags this as a live argument for x86 over ARM in a rack, and
  it applies to the entire Think Tiny roster including the 2013 machines.
- **Wake-on-LAN** — universal, including from S5 on most generations.
- **RTC / scheduled wake** — present on most; useful for duty-cycled nodes.
- **Smart Power On** — a designated rear USB port wakes the machine from a keyboard press (Tiny7/8).
- **Intel vPro / AMT** — the big one. On vPro SKUs (M83, M93p, M910q, **M920q**, M80q, M90q any gen,
  P-series, plus the **i5-8365U / i7-8665U Nano**) you get **out-of-band KVM, remote power control,
  remote boot media, and serial-over-LAN — over the onboard NIC, with the OS down.** In a rack that
  replaces a per-node JetKVM (~$70 each) and a switched PDU port.
  - Caveats: AMT must be provisioned (it ships unconfigured), the web UI is dated, and AMT has had
    serious CVEs historically — **put the management interface on a separate VLAN.**
- **AMD DASH** on PRO APUs — comparable capability, thinner tooling and documentation.
- **No vPro, no AMT:** M70q (Gen 1–3), M75q (all), neo 50q (all), M625q/M715q, and the non-vPro CPU SKUs
  of otherwise-capable models. **The CPU determines vPro, not just the chassis** — an M920q with an
  i5-8400T has no vPro because the 8400 tier doesn't carry it. Check the chip, not the badge.

## 3.8 Storage layout by era — what you can actually attach

| Era | Bays | M.2 | Practical maximum |
|---|---|---|---|
| Legacy → Tiny5 | 1× 2.5" SATA (7 mm/9.5 mm) | 1× M.2 2280 (NVMe from Tiny4) | ~2 TB NVMe + 4 TB SATA SSD |
| Tiny6 (M90q Gen 2, M75q Gen 2) | 1× 2.5" | 1–2× M.2 (Gen 3) | ~2 TB × 2 + 4 TB SATA |
| Tiny7 (Gen 3–5) | **bay dropped on most** | 2× M.2 Gen 4 ×4, RAID 0/1 | ~4 TB × 2 |
| **Tiny8 (M90q Gen 6)** | none | **3× M.2 Gen 4 ×4** + WLAN slot | ~2 TB × 3, RAID 0/1 |
| **Tiny8 (P3 Tiny Gen 2)** | none | **3× M.2 — 2× Gen 4, 1× Gen 5 ×4**, NVMe RAID 0/1/5 | ~2 TB × 3 |

**For bulk storage, none of these is the answer.** The Tiny is a *head*, not an array. The realistic
patterns are: (a) NVMe-only, treat the node as stateless/compute; (b) NVMe boot + a USB 3 DAS enclosure
(cheap, adequate, avoid USB-attached ZFS for anything you care about); (c) riser HBA + external SAS
(M720q/M920q/M90q only, and the HBA will be hot); (d) point the Tiny at a real NAS over the network,
which is what most people end up doing.

---

# Part 4 — Task ↔ silicon suitability

## 4.1 Pick the node by the bottleneck, not the badge

Before the table: almost every "which Tiny for job X" question reduces to **which one of six things the
job is actually limited by.**

| Bottleneck | What to look for | Which models have it |
|---|---|---|
| **Thread count** | 8C/16T+ at 35 W | i9-9900T (M720q/M920q) · R7 PRO 4750GE/5750GE (M75q Gen 2) · i9-12900T (16C) · i9-13900T / i9-14900T (24C) · i7-14700T (20C) · R7 PRO 8700GE · Ultra 9 285T (24C) |
| **Single-thread** | high boost, recent µarch | Ultra 9 285T (5.4 GHz) · i9-14900T (5.5) · R7 PRO 8700GE (5.1) · i9-13900T (5.3). Old 6th–9th gen tops out ~4.4 GHz and a much lower IPC |
| **PCIe / NIC** | **a riser slot** | M720q · M920q · M920x · M90q (all gens) · all ThinkStation P Tinies. **Nothing else.** |
| **iGPU / media engine** | QuickSync gen, or a dGPU | UHD 630 (8th/9th) = baseline · UHD 770 (12th+) = +AV1 decode · Arc A310 (M90q Gen 6) = AV1 **encode** · Radeon 780M (M75q Gen 5) = AV1 encode · NVIDIA T/A-series (P Tinies) = unlimited NVENC sessions |
| **RAM** | DDR gen and ceiling | 32 GB (pre-2018) → 64 GB (most) → **128 GB verified on M75q Gen 5 (M)** and officially on **P3 Tiny Gen 2** |
| **Watts / silence** | 35 W `T`/`GE` parts, low core count | M600/M53 (~6–9 W) · Nano (~5–8 W) · M75q Gen 5 (**5–15 W M**) · M720q (~8–10 W M) |
| **Budget / availability** | used supply depth, and how much RAM comes in the box | Abundant + cheap: M720q · M920q · M75q Gen 2 · M90q Gen 1 · P340. Unobtainable used: M90q Gen 6 · P3 Tiny Gen 2 · neo 50q Gen 5. **See §1.10** |

**Rules of thumb that hold across the whole roster:**

- **The `T` and `GE` parts are the rack parts.** Any 65 W `non-T`/`G` SKU in a 1 L box is buying cores you
  cannot sustain. It will benchmark well and throttle in production.
- **More E-cores ≠ more sustained performance in this chassis.** A 24-core i9-13900T at 35 W is a
  throughput machine for many small parallel tasks (containers, compile jobs, encode queues) and a poor
  match for a few latency-sensitive threads.
- **Buy the NIC path before the CPU.** A 6-core M920q with a 10GbE card beats a 16-core M70q Gen 3 for
  any storage, firewall or capture role, because the second machine physically cannot get there.
- **vPro is worth ~$40 of premium in a rack** and displaces a per-node IP-KVM.
- **No ECC anywhere.** Every "data integrity" workload on a Tiny is doing checksums in software on
  non-ECC memory. That's fine for backups and media; think twice for a primary array.
- **On a used-only budget, availability beats capability.** Several "best fit" entries in the table below
  are models you cannot currently buy at any sane price. Where that happens the **Also fine** column is
  the real recommendation — and §4.4 is written entirely against what is actually purchasable.

## 4.2 The task table

**Legend:** ✅ excellent · 🟡 workable with caveats · ❌ actively wrong choice for the hardware.

NOTE: Part 5 is the hardware-neutral superset of this table — 398 tasks across 15 domains, and the source of the resource-tag vocabulary behind the `Bottleneck` column here. **Its scoping rule governs how much any one row should say:** name only the workloads a machine's hardware genuinely decides, not everything it can run.

### Virtualization, orchestration and infrastructure

| Task | Hardware bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Proxmox VE / XCP-ng host** | threads, RAM | M75q Gen 5 (128 GB) · M90q Gen 4/5/6 · M75q Gen 2 | M720q/M920q i9-9900T · M70q Gen 3–5 | M625q, M53/M600, Nano | The canonical Tiny job. VT-x/AMD-V on everything since Skylake |
| **VMware ESXi** | driver support, hybrid scheduling | M920q, M90q Gen 1/2 (homogeneous cores, Intel NIC) | 🟡 M70q Gen 5 | ❌ hybrid P+E (12th gen+) and Realtek-NIC AMD boxes | ESXi's HCL and hybrid-core handling are the constraints, not raw speed |
| **Hyper-V / Windows Server lab** | RAM, vPro | M90q any gen | M80q, M920q | Nano | vPro AMT pairs naturally with Windows tooling |
| **k3s / k8s worker node** | threads, RAM | M75q Gen 2 (8C/16T cheap) · M70q Gen 5 | M720q i7-8700T · neo 50q Gen 5 | M625q | Cheapest cores-per-dollar wins here |
| **k3s / k8s control plane** | etcd write latency → **NVMe IOPS** | any NVMe-equipped Tiny; Nano is fine | M910q | ❌ anything booting from SATA HDD | etcd hates slow fsync. Put it on NVMe, always |
| **Ceph OSD node** | NVMe count, network | ✅ **M90q Gen 6** (3× Gen4 NVMe + x8 slot for 10GbE) · P3 Tiny Gen 2 | 🟡 M90q Gen 3–5 (2 NVMe) | ❌ single-M.2 models | Ceph on 1 GbE is a mistake; you need the riser |
| **ZFS host** | RAM, SHA throughput, **no ECC** | M75q Gen 2/Gen 5 (SHA-NI) | M90q Gen 3+ | ❌ treat as primary array | Use `fletcher4`, accept non-ECC, keep real backups elsewhere |
| **PXE / netboot / MAAS / Foreman** | almost nothing | any, incl. legacy | M600, M710q | — | A 2016 Tiny is genuinely sufficient |
| **Ansible / Terraform control node** | single-thread, RAM | any modern Tiny | Nano | — | Latency-sensitive, not throughput-sensitive |
| **NetBox / IPAM / DCIM** | Postgres single-thread | M720q+ | Nano | — | |
| **Jump host / bastion** | nothing | Nano, M625q, M600 | any | — | Ideal use for the low-power tier |
| **Blue-green / staging environment** | threads, RAM | M75q Gen 2, M70q Gen 4/5 | — | — | |

### Containers, apps and web

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Docker / Podman host (20–50 containers)** | threads, RAM, NVMe | M75q Gen 2 · M70q Gen 5 · M90q Gen 4+ | M720q i7-8700T | M625q, M53 | |
| **Reverse proxy (Traefik / Caddy / NPM)** | AES-NI, single-thread | anything Skylake+ | Nano, M600 | — | TLS termination is AES-NI-bound; every chip here has it |
| **Web / app hosting** | single-thread, RAM | Ultra/Zen 4 tier | M720q+ | — | |
| **Postgres / MySQL / MariaDB** | **single-thread + NVMe fsync latency** | Ultra 9 285T · R7 PRO 8700GE · i9-14900T | M920q i9-9900T | ❌ HDD-booted, ❌ Atom-class | Prefer high boost over high core count |
| **Redis / Valkey / memcached** | single-thread, RAM bandwidth | DDR5 models (M90q Gen 3+, M75q Gen 5) | any | — | Single-threaded by design |
| **ClickHouse / DuckDB / analytics** | cores + RAM bandwidth + NVMe | M75q Gen 5 (128 GB) · P3 Tiny Gen 2 | M90q Gen 5/6 | single-M.2 models | Vectorized engines like DDR5 and many cores |
| **MinIO / Garage / SeaweedFS (object store)** | NVMe count, network | M90q Gen 6 (3 NVMe) | M90q Gen 3–5 | ❌ 1-NVMe models | |
| **Vaultwarden / password manager** | nothing | Nano, M600, M625q | any | — | |
| **Nextcloud / Immich / Paperless** | threads + iGPU (Immich ML) + storage | M75q Gen 5 (NPU) · M90q Gen 6 (NPU) | M920q + Coral | ❌ pre-AVX2 (M53/M600) | Immich's ML container wants AVX2 at minimum |
| **Mail server** | reliability, uptime | M920q (vPro) | any | — | The vPro out-of-band story matters most on things you can't afford to walk to |
| **Matrix / XMPP / Mumble / Jitsi** | single-thread; Jitsi wants cores | M70q Gen 4/5 for Jitsi | M720q for Matrix | — | Jitsi video routing is CPU-hungry |
| **Git server + CI runners (Gitea/Forgejo)** | cores, NVMe | M75q Gen 2/5 · M70q Gen 5 | M920q | — | |

### Storage, backup and NAS

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **NAS (SMB/NFS) head** | **network + drive attach** | 🟡 M920q/M90q + riser HBA or 10GbE | 🟡 M720q + USB DAS | ❌ any riser-less model as a *primary* NAS | **This is the Tiny's weakest role.** One bay, one or two M.2. Use it as a head, not an array |
| **All-flash NVMe share** | NVMe count + network | ✅ **M90q Gen 6** / P3 Tiny Gen 2 + 10GbE | M90q Gen 3–5 | — | Three Gen4 ×4 slots is a genuinely good small all-flash node |
| **Backup target (Borg / Restic / PBS)** | sequential I/O, capacity | M720q/M920q (2.5" bay + USB DAS) | legacy Tiny with a 4 TB SATA SSD | Tiny8 (no bay) | Ironically the *older* Tinies are better at this, because they still have a SATA bay |
| **Proxmox Backup Server** | NVMe IOPS for chunk store + capacity | M90q Gen 6 | M90q Gen 3+ | — | PBS chunk verification is CPU+IOPS heavy |
| **Seedbox** | storage, sustained network | M720q + DAS | any | — | |
| **Bitcoin/Ethereum full node** | NVMe IOPS + capacity + sustained I/O | M90q Gen 6 (2 TB × 3) · P3 Tiny Gen 2 | M70q Gen 5 | ❌ single-M.2, ❌ SATA-only | Chain data outgrows a 1 L box fast; plan for external |
| **Ceph / Longhorn / distributed storage member** | NVMe + 10GbE | M90q Gen 6 | M90q + riser NIC | riser-less models | |

### Networking, firewall and security

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **pfSense / OPNsense firewall** | **Intel NIC + riser**, AES-NI | ✅ **M920q + I350-T4 or X520 in the riser** · M90q any gen | 🟡 M720q + riser | ❌ **all AMD Tinies, M70q, M80q, neo 50q, Nano** — no riser, Realtek NIC | `nodes.md` already lists this as the classic homelab mod. The riser is non-negotiable |
| **VyOS / OpenWrt / Linux router** | NIC count | M920q/M90q + riser | 🟡 any + USB 2.5GbE NIC | — | Linux is far more forgiving of Realtek than BSD |
| **10GbE routing / inter-VLAN** | riser + CPU single-thread | M90q Gen 5/6 | M920q i9-9900T + X520 | riser-less | Software routing at 10G is single-thread-bound |
| **IDS/IPS (Suricata / Zeek)** | **cores + a real capture NIC** | M90q Gen 4/5/6 + Intel NIC | M920q i9-9900T + NIC | ❌ riser-less, ❌ Realtek | Suricata scales with cores; needs RSS queues a Realtek won't give you |
| **DNS (Pi-hole / AdGuard / Unbound / Technitium)** | almost nothing | Nano, M600, M625q | any | — | The single best use for the low-power tier |
| **DHCP / NTP / PTP** | nothing | any legacy Tiny | — | — | For PTP/GPS-disciplined time you want a serial/GPIO path — the **serial punch-out** on modern SKUs helps |
| **VPN server / WireGuard / Tailscale exit node** | AES-NI + single-thread + NIC | M90q + 2.5GbE | M920q, Nano | — | WireGuard is single-thread-per-tunnel; boost clock matters more than cores |
| **Tor relay / bridge / I2P** | network-bound, near-zero CPU | Nano, M600, M625q, M710q | any | — | Cheapest node you own. Bandwidth, not compute |
| **Wazuh / Graylog / SIEM node** | RAM + NVMe + cores | M75q Gen 5 (128 GB) · M90q Gen 5/6 | M75q Gen 2 | ❌ 16 GB models | Elastic/OpenSearch is RAM-hungry |
| **Certificate authority / Vault / step-ca** | nothing much, but wants TPM | any modern Tiny (**discrete TPM 2.0** standard) | — | — | Every modern Tiny ships discrete TPM 2.0, FIPS 140-2 certified |
| **Malware sandbox / detonation VM** | isolation, snapshot speed, cores | M90q Gen 4+ | M75q Gen 2 | — | Keep it on its own VLAN and its own physical box |
| **Honeypot / canary** | nothing | M600, M625q, Nano | any | — | |

### Media

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Jellyfin / Plex / Emby — 1–4 streams** | QuickSync generation | ✅ **M720q / M920q (UHD 630)** — best price/performance transcoder in the roster | M910q (HD 630), M70q Gen 3+ (UHD 770) | 🟡 AMD Tinies (VAAPI works, less well-trodden) | UHD 630 does H.264, HEVC 8/10-bit, VP9 decode |
| **Jellyfin / Plex — many simultaneous streams** | **encoder session limits** | ✅ **ThinkStation P-series with T400/T600/T1000/A400/A1000** — pro NVENC has **no session cap** | M90q Gen 6 + Arc A310 | consumer GPUs (capped) | This is the reason to buy a P-series Tiny |
| **AV1 library (decode)** | media engine | M70q/M80q/M90q Gen 3+ (UHD 730/770) · M75q Gen 5 (780M) · P3 Tiny (A400/A1000) | — | ❌ everything 11th gen and older | |
| **AV1 encode / re-encode pipeline** | media engine | ✅ **M90q Gen 6 + Arc A310** · M75q Gen 5 (RDNA 3) | — | ❌ all others | Only two paths in the whole roster |
| **Live streaming / SRT relay / restreamer** | encoder + network | P-series (NVENC) · M90q Gen 6 (Arc) | M920q (QSV) | — | |
| **Music (Navidrome) / audiobooks / ebooks** | nothing | Nano, M600 | any | — | |
| **Photo library with ML tagging** | NPU or iGPU or dGPU | M75q Gen 5 (16 TOPS) · M90q Gen 6 (13 TOPS) · P3 Tiny (RTX A1000) | M920q + Coral USB | ❌ pre-AVX2 | |
| **Digital signage / kiosk** | display outputs | any with punch-out DP/HDMI (up to 4 displays on modern) | P-series (6 displays) | — | |

### AI, ML and inference

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Local LLM, small (3–8 B quantized)** | RAM bandwidth + iGPU/NPU | ✅ **M75q Gen 5 (780M + 128 GB DDR5)** · P3 Tiny / P3 Tiny Gen 2 (RTX A1000 8 GB) | M90q Gen 6 (NPU + DDR5) | ❌ DDR4 models — bandwidth-starved | 8 GB VRAM is the practical ceiling for the dGPU path |
| **Local LLM, large (30 B+ quantized, CPU)** | RAM capacity + bandwidth | M75q Gen 5 @ 128 GB · P3 Tiny Gen 2 @ 128 GB | M90q Gen 5/6 @ 64 GB | — | Slow but genuinely usable for batch work |
| **Object detection (Frigate, CodeProject.AI)** | NPU / dGPU / Coral | M75q Gen 5 (NPU) · P-series (NVENC decode + CUDA detect) | M720q + Coral USB (the classic cheap combo) | — | Frigate wants **iGPU for decode** and something else for detect |
| **Whisper / ASR transcription** | dGPU or many cores | P3 Tiny (RTX A1000) | M75q Gen 5, i9-13900T | Nano | |
| **Stable Diffusion / image generation** | **VRAM** | 🟡 P3 Tiny / P3 Tiny Gen 2 (8 GB A1000) | 🟡 M75q Gen 5 (iGPU, slow) | ❌ everything else | 8 GB and 50 W is the ceiling. This format is not a good SD box |
| **Training anything** | ❌ | — | — | ❌ all | Don't. Rent it. |
| **Voice assistant (Wyoming / Piper / Rhasspy)** | modest CPU + NPU nice-to-have | M75q Gen 5, M90q Gen 6 | M720q | M53/M600 | |

### Development, build and data

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Heavy compiling (kernel, LLVM, Chromium)** | **sustained all-core** — the 1 L thermal ceiling bites here | 🟡 i7-14700T / i9-14900T (M90q Gen 5) · R7 PRO 8700GE (M75q Gen 5) · Ultra 9 285T | i9-9900T, R7 PRO 5750GE | ❌ 65 W `non-T`/`G` (throttle) | **Honest verdict:** any 1 L Tiny is a *mediocre* compile box relative to its core count. Buy for cores-per-watt, and use `distcc`/`sccache` across several nodes rather than one big one |
| **Distributed build farm (distcc, Nix, Bazel remote)** | cores across nodes | several M75q Gen 2 (cheap 8C/16T) | M70q Gen 5 | — | **This is the right answer** to the compile problem in this format |
| **CI runners (GitHub Actions, GitLab, Woodpecker)** | cores, NVMe, RAM | M75q Gen 2/5 · M70q Gen 5 | M920q | — | |
| **Cross-compile / Yocto / Android ROM builds** | cores + RAM + NVMe | M75q Gen 5 @ 128 GB | i9-13900T models | 16 GB models | Yocto will use every GB you give it |
| **Remote dev box (VS Code Server, JetBrains Gateway)** | single-thread + RAM | Ultra tier, R7 PRO 8700GE | M720q i7-8700T | Nano | Interactive latency = single-thread |
| **Constant unattended mass scraping** | **RAM + NVMe write endurance + threads**; per-worker single-thread if headless-browser | M75q Gen 5 (128 GB) · M75q Gen 2 (cheap threads) · M70q Gen 5 | M920q i9-9900T | ❌ 8–16 GB models | Headless browsers are memory hogs and leak. Cap containers, recycle contexts, and keep the churny DB on local NVMe — network filesystems will punish you |
| **ETL / Airflow / data pipelines** | cores + RAM + NVMe | M90q Gen 5/6 · M75q Gen 5 | M75q Gen 2 | — | |
| **Distributed computing (BOINC / Folding)** | sustained all-core, watts | 🟡 i9-13900T at 35 W is efficient | — | ❌ 65 W SKUs | Perf-per-watt favours `T` parts strongly |

### Home, physical and IoT

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Home Assistant** | single-thread + USB (Zigbee/Z-Wave sticks) | M720q, M70q Gen 3+ | Nano | ❌ M53/M600 (too slow for large automation sets) | Watch USB 2.0 port counts on M715q/M75q-1 |
| **Frigate NVR** | iGPU decode + detect accelerator + NVMe write endurance | M920q + Coral · M75q Gen 5 (NPU) · P-series (CUDA) | M720q | pre-Skylake | Recordings will destroy a cheap NVMe — use a high-TBW drive or a SATA SSD in the bay |
| **3D printer host (OctoPrint / Klipper)** | **USB latency + a few reliable cores** | Nano, M600, M710q | any | — | Klipper's MCU timing wants a machine that isn't also doing something bursty |
| **SDR / ADS-B / marine AIS / Meshtastic gateway** | **USB 3 bandwidth and port count** | M720q/M920q (6× USB 3.1) · M90q Gen 6 | M70q Gen 5 | ❌ M715q / M75q-1 (half USB 2.0) · ❌ Nano (few ports) | Multiple RTL-SDR dongles saturate USB controllers; check port topology, not just count |
| **Ham / APRS / packet gateway** | serial + USB | modern Tiny with **serial punch-out** | any + USB-serial | — | Lenovo offers up to **4× serial via a PCIe x1 card** on M90q Gen 6 / P3 Tiny Gen 2 — unusual and useful |
| **LinuxCNC / realtime motion control** | **realtime latency**, not throughput | 🟡 older non-hybrid Intel (M710q/M720q with HT off) | — | ❌ hybrid P+E (12th gen+), ❌ AMD SMT-heavy | Realtime kernels and hybrid schedulers do not mix well |
| **ROS / robotics host** | cores + USB + GPU | P-series Tiny | M70q Gen 5 | — | |

### Monitoring, observability and ops

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Prometheus + Grafana + Loki** | NVMe write IOPS + RAM | M90q Gen 3+ | M720q | ❌ HDD | TSDB compaction is IOPS-heavy |
| **Zabbix / Netdata / LibreNMS** | modest | any modern Tiny | Nano | — | |
| **Uptime Kuma / status page** | nothing | Nano, M600, M625q | any | — | |
| **Syslog / log aggregation** | disk write + capacity | M720q + SATA bay | any | — | |
| **Speedtest tracker / network probes** | NIC | any | Nano | — | |

### Games, streaming and desktop-in-a-rack

| Task | Bottleneck | Best fit | Also fine | Avoid | Notes |
|---|---|---|---|---|---|
| **Dedicated game servers (Minecraft, Valheim, Palworld, Satisfactory)** | **single-thread** above all | Ultra 9 285T · i9-14900T · R7 PRO 8700GE | M920q i9-9900T | ❌ Atom-class, ❌ many-E-core-few-P-core if running one big server | Most game server ticks are single-threaded. One high-boost core beats sixteen slow ones |
| **Multiple small game servers** | cores | i9-13900T, R7 PRO 5750GE | M75q Gen 2 | — | Now core count wins |
| **CS2 / source-engine servers** | single-thread + network jitter | Ultra tier | — | — | |
| **Game streaming host (Sunshine / Moonlight / Parsec)** | **encoder + GPU** | P-series Tiny (NVENC) · M90q Gen 6 (Arc) | 🟡 M75q Gen 5 (780M) | ❌ iGPU-only pre-12th-gen | |
| **Retro emulation host** | single-thread + iGPU | M75q Gen 5 (780M) | M70q Gen 4/5 | — | |
| **VDI / thin-client backend** | RAM, vPro, GPU | P3 Tiny Gen 2 · M90q Gen 6 | M90q Gen 4/5 | — | Lenovo sells the neo 50q as a thin-*client*; the M90q is the thin-client *host* |
| **Bare-metal Windows workstation in the rack** | anything | M90q Gen 6, P3 Tiny Gen 2 | any | — | The reason "After Power Loss = On" plus vPro matters |

## 4.3 Traits and superlatives — the "which one is the most X" list

| Superlative | Winner | Runner-up |
|---|---|---|
| **Lowest idle watts, real x86** | M600 / M53 (`~6–9 W`) | Nano M90n-1 / M75n (`~5–8 W`) |
| **Lowest idle with modern cores** | **M75q Gen 5 — 5–15 W (M)** | M720q — 8–10 W (M) |
| **Best value, used, all-round** | **M720q + i5-8500T or i7-8700T** | M75q Gen 2 + R5 PRO 4650GE |
| **Most threads per dollar, used** | **M75q Gen 2 + R7 PRO 4750GE / 5750GE** (8C/16T) | M720q/M920q + i9-9900T |
| **Most cores at 35 W, any era** | **i9-13900T / i9-14900T / Ultra 9 285T — 24 cores** | i7-14700T (20C/28T) |
| **Best sustained MT at 35 W** | i7-14700T (best measured MT of the `T` parts) | R7 PRO 8700GE |
| **Best single-thread** | Ultra 9 285T / i9-14900T | R7 PRO 8700GE |
| **Best transcoder per dollar** | **M720q / M920q — UHD 630** | M70q Gen 3 (UHD 770, +AV1 decode) |
| **Best transcoder, period** | **ThinkStation P-series + NVIDIA T1000/A1000** (no NVENC session cap) | M90q Gen 6 + Arc A310 (AV1 encode) |
| **Only AV1 *encode* paths** | M90q Gen 6 + Arc A310 · M75q Gen 5 (RDNA 3) | — |
| **Most expandable** | **P3 Tiny Gen 2** — 3× M.2 (one Gen 5), x16(x8) riser, 128 GB, 2× extra 2.5GbE | M90q Gen 6 |
| **Most NVMe** | M90q Gen 6 / P3 Tiny Gen 2 — **3 slots** | M90q Gen 3–5 — 2 slots |
| **Most RAM** | **128 GB** — P3 Tiny Gen 2 (official), M75q Gen 5 (verified by STH) | 64 GB on nearly everything since 2018 |
| **Only PCIe 5.0 storage** | P3 Tiny Gen 2 | — |
| **Best out-of-band management** | M920q / M90q any gen / P-series (vPro AMT) | Nano i7-8665U (vPro at 15 W) |
| **Quietest under load** | 35 W `T` parts in any modern chassis, ICE set to acoustic | — |
| **Loudest under load** | M75q Gen 5 at sustained 63–66 W (**47–48 dBA M**), P-series with GPU | — |
| **Most repairable / longest service life** | **M75q Gen 5** (toolless RAM + both M.2, standard parts, published HMM) | M720q/M920q (huge parts supply, cheap risers, standard slim-tip bricks) |
| **Best NIC options from the factory** | M90q Gen 6 (I350-T4 quad GbE, 2.5GbE, TB4) | P3 Tiny Gen 2 (I350-T2/T4, BCM5719/5720, 2× 2.5GbE BTB) |
| **Worst watts-per-thread** | M715q Gen 1/2 (14–17 W idle for 4 slow cores) | M625q |
| **Best cost-per-thread, used** | **M75q Gen 2 + R7 PRO 5750GE** — 8C/16T at ~$200–320 | M720q / M920q + i9-9900T |
| **Best capability-per-dollar, used** | **M920q** — vPro + riser + UHD 630 from **$161 V** | M90q Gen 1 (~$140–260) |
| **Best machine you cannot buy used** | M90q Gen 6 · P3 Tiny Gen 2 | M75q Gen 5 (very thin, ~$350–550) |
| **Worst value on the used market** | **M70q Gen 5** — **$799.99 V** for a riser-less node | neo 50q Gen 4 against an M75q Gen 2 |
| **Cheapest vPro/AMT in the roster** | M93p Tiny (~$40–85) | M90n-1 Nano i5-8365U (~$80–170, ~6 W idle) |

## 4.4 Quick picks — used market only

Prices are the working-config midpoints from §1.10. Everything here is something you can actually buy
today; the models excluded for unavailability are named at the end so the omission is deliberate, not
an oversight.

**The default answer, and it is not close:**

> **M920q · i7-8700T or i9-9900T · 16–32 GB · + 01AJ940 riser with baffle · + Intel I350-T4 or
> ConnectX-3 · + a printed PETG mount. `~$220–290` all in.**
>
> vPro AMT for out-of-band power and KVM, UHD 630 for transcode, a real PCIe NIC, 8–12 W idle, a decade
> of parts availability, and a $10 riser that no other price-comparable model has. Buy the M720q instead
> if you find one with more RAM for less money — the vPro premium is currently near zero, so let RAM
> content decide.

| If you want… | Buy | `~$` | Why |
|---|---|---|---|
| **One node that does everything** | M920q (i7-8700T/i9-9900T) + riser + NIC | 220–290 | See above |
| **The cheapest capable node** | M720q (i5-8500T, 16 GB) | 110–200 | Nothing beats it under $200 |
| **Most threads for the money** | M75q Gen 2 (R7 PRO 5750GE, 8C/16T) | 200–320 | Zen 3, SHA-NI, 64 GB ceiling |
| **A cheap cluster** | 3× M75q Gen 2 (R5 PRO 4650GE) | 400–650 | 18C/36T. The right answer to "I need to compile things" in this format |
| **Firewall / router** | M920q or M90q Gen 1 + riser + I350-T4 | 220–330 | **Intel NIC via riser is mandatory.** Never an AMD Tiny |
| **10GbE storage frontend** | M920q (i7-9700T, 32 GB) + riser + ConnectX-3 + 2× NVMe | 380–500 | ConnectX-3 is the cheapest 10GbE that fits |
| **Media node, many streams** | P340 Tiny (i7-10700T) with T1000 | 180–340 | Pro NVENC has **no session cap**. Best value in the P-series |
| **Media node, few streams** | M720q / M920q (UHD 630) | 110–230 | QuickSync is enough below ~4 streams; save the money |
| **Lowest-watt always-on tier** | 2× M90n-1 Nano (i5-8365U — has vPro) | 180–350 | ~6 W each, out-of-band management included |
| **Absolute cheapest always-on** | M600 or M93p Tiny | 35–85 | M93p for vPro, M600 for ~7 W idle |
| **A GPU for inference** | P340 or P360 Tiny with T1000 8 GB | 180–520 | 8 GB VRAM ceiling; fine for 7–8 B models and Frigate detection |
| **AVX-512 (the only route)** | M70q Gen 2 / M90q Gen 2 / P350 Tiny, i5-11500T or above | 130–390 | Fused off from 12th gen onward — there is no other option |
| **A cold spare / donor** | any second M720q | 110–200 | Cheaper than sourcing a board, cooler, brick or bottom plate individually |

**Deliberately not recommended, and why:**

| Model | Reason |
|---|---|
| **M90q Gen 6 · P3 Tiny Gen 2 · neo 50q Gen 5** | **Unobtainable used.** Genuinely the best hardware in the roster; revisit ~2028 |
| **M75q Gen 5** | Very thin supply, `~$350–550`. The right machine at the wrong moment — the price should halve by 2027 |
| **M70q Gen 5** | Observed at **$799.99 V**. Three M75q Gen 2s cost less and give 24C/48T with a spare left over |
| **neo 50q Gen 4** | `~$180–330` for soldered mobile silicon with no vPro and no riser. An M75q Gen 2 is cheaper and better |
| **M715q Gen 1 / 2** | 14–17 W idle for four slow cores. Worst watts-per-thread here even at $55 |
| **M625q** | 6 W and 2 cores. Only if free |
| **Barebones listings** (any model) | Under current DRAM pricing, filling an empty box costs more than buying a full one |
| **Any 65 W `non-T` / `G` SKU** | Throttles in 1 L. You pay for cores you cannot sustain |

**The one non-obvious buying rule, restated:** across the whole abundant tier, **the RAM in the box is
worth more than the CPU in the box.** Sort listings by installed memory, not by processor model, and the
right machine falls out.

---

# Part 5 — Workload catalog

A task-first index of workloads these nodes plausibly run, and the hardware property that actually decides each one. This is the inverse of the device-first tables elsewhere in this document: use it to go from "I need to run X" to "so I need hardware that does Y". It is deliberately hardware-neutral — no device is named here. The tags are the controlled vocabulary the rest of this document should use when describing what a node is for.

**Scoping rule — this list is exhaustive, per-device recommendations are not.** When describing what a given device is for, name only the handful of workloads its hardware genuinely decides: the things it is better at than the boards next to it in the table. Do not enumerate everything it is merely capable of. Nearly every node here can run nearly everything on this list, so a long list carries no information. Where a device is simply a competent generalist, say exactly that in one line and spend the remaining space on its caveats — what it cannot do, where it runs out of headroom, and what it forces you to work around. "Good all-rounder, but single-lane NVMe and no hardware encoder" is worth more to a reader than thirty task names.

## Resource-tag legend

| Tag | Means | Why it changes the pick |
|---|---|---|
| `ST` | single-thread bound | clock and IPC; extra cores add nothing |
| `MT` | all-core bound | core count scales it near-linearly |
| `MEM` | RAM capacity bound | it fits or it thrashes — most node builds have no swap |
| `BW` | memory bandwidth bound | channel count and DDR speed; where wide x86 pulls away from ARM |
| `IOPS` | random I/O bound | NVMe vs SATA vs SD is an order of magnitude |
| `FSYNC` | durable-write-latency bound | the specific SSD's sync-write behaviour; DRAM-less QLC falls off a cliff |
| `CAP` | disk capacity bound | bays, lane budget, whether it can host spinning disks at all |
| `WEAR` | write-endurance bound | TBW; kills SD cards and cheap flash in months |
| `NET` | network throughput bound | NIC speed, and whether the SoC's lanes can actually feed it |
| `PPS` | packet-rate bound | small-packet forwarding, NIC quality, offloads — not GHz |
| `CRYPTO` | crypto-throughput bound | AES-NI / ARMv8 crypto extensions swing this by 10x or more |
| `VPU-ENC` | fixed-function encoder | presence is binary; replaces an entire CPU's worth of work |
| `VPU-DEC` | fixed-function decoder | same on the input side; matters most for concurrent streams |
| `GPU` | 3D or general GPU compute | discrete or a capable iGPU; rules out most headless SBCs |
| `NPU` | NN accelerator | also binary — with it, watts; without it, a whole node |
| `USB` | USB bandwidth or topology | port count, controller sharing, stable enumeration for attached devices |
| `PCIE` | needs lanes or a slot | HBAs, extra NICs, accelerators, passthrough |
| `GPIO` | needs pins or a low-level bus | GPIO, I2C, SPI, UART, PPS, CAN — a hard filter on the hardware list |
| `VIRT` | virtualization extensions | full VMs; nested virt and IOMMU narrow it further |
| `RT` | latency-deterministic | jitter matters more than throughput; hates noisy co-tenants |
| `THERM` | sustained-load bound | cooling and power limits, not peak boost |
| `IDLE` | effectively free | the constraint is uptime and reliability, not performance |
| `x86` | needs amd64 | ARM unsupported, or a standing compatibility tax |

## Media & streaming

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Live media transcoding** | on-the-fly re-encode per client | `VPU-ENC` `VPU-DEC` | the fixed-function encoder is the whole story. Concurrent streams scale with encoder blocks, not cores; CPU-only transcoding is a last resort |
| **Batch archival re-encode** | overnight library recompression to a denser codec | `MT` `THERM` `VPU-ENC` | quality-per-bit favours slow CPU presets, so sustained all-core plus cooling. Hardware encoders collapse the runtime at some quality cost |
| **Remux / container rewrite** | stream copy, no re-encode | `CAP` `NET` | disk and NIC only — runs fine on the weakest node in the rack |
| **Live ingest and restream** | receive one feed, fan out variants | `NET` `VPU-ENC` `RT` | upstream bandwidth plus one encoder pass per variant; jitter buffers dislike noisy neighbours |
| **Streaming packager and segmenter** | produce DASH/HLS segments and manifests | `IOPS` `NET` `CAP` | small-file write rate; segment churn is harder on the disk than the bitrate suggests |
| **Audio transcode and loudness normalization** | format conversion, EBU R128 scanning | `MT` | embarrassingly parallel, tiny RAM — good filler work for otherwise idle cores |
| **Audio fingerprinting and acoustic ID** | match audio against a reference corpus | `MT` `MEM` | the reference index wants to stay resident |
| **Stem separation and source isolation** | split a mix into instrument tracks | `GPU` `MEM` `THERM` | model inference per track; CPU-only runs are hours, not minutes |
| **Podcast and audiobook processing** | chapterize, trim silence, normalize | `MT` `IOPS` | cheap and parallel |
| **Image thumbnailing and derivatives** | resize and re-encode for libraries | `MT` `IOPS` | thousands of small reads; NVMe over SD or spinning disk by a wide margin |
| **Video preview sprite generation** | trickplay and scrub-bar thumbnails | `MT` `VPU-DEC` `IOPS` `CAP` | a full decode pass over the library, then a lot of small files |
| **Perceptual hashing and dedup sweep** | find near-duplicate media | `MT` `IOPS` | read-bound full sweep; SIMD width helps |
| **Chapter and scene detection** | keyframe and shot-boundary extraction | `MT` `VPU-DEC` `IOPS` | decode throughput, not analysis cost |
| **Ad detection and removal** | identify and cut commercial segments | `MT` `VPU-DEC` `THERM` | a decode-and-analyse pass per recording; queues up fast on a busy DVR |
| **HDR tone-mapping pass** | convert HDR to SDR for incompatible clients | `GPU` `VPU-ENC` | software tone-mapping is punishingly slow — a GPU path or don't offer it |
| **Subtitle and transcript generation** | speech recognition over a media library | `MT` `GPU` `NPU` | a GPU is 10-50x. CPU is fine if it can run overnight; impractical on a few weak cores |
| **Live caption overlay and burn-in** | render captions into the video stream | `VPU-ENC` `MT` `RT` | forces a re-encode, so it inherits the transcoding constraint |
| **Music and radio streaming** | serve or live-mix audio | `IDLE` `RT` `NET` | near-free; live mixing wants no scheduling stalls |
| **Multi-room audio sync** | synchronized playback across zones | `RT` `NET` `IDLE` | clock discipline, not throughput. Milliseconds of drift are audible |
| **Media library scan and metadata refresh** | walk the library, fetch and update metadata | `IOPS` `MT` `NET` | directory-walk speed on large libraries; the API calls are the slow part |
| **Media library management** | indexers, renaming, download automation | `IDLE` | many small always-on services, each costing almost nothing |
| **Photogrammetry and offline 3D render** | mesh reconstruction, ray tracing | `MT` `MEM` `GPU` `THERM` | RAM-hungry, and ARM tooling is thin — x86 plus a real GPU |
| **Desktop and game streaming host** | encode a live interactive session | `VPU-ENC` `GPU` `RT` | low-latency encoder and a GPU; not a headless job |
| **DVR and broadcast tuner capture** | scheduled recording from tuner hardware | `PCIE` `USB` `CAP` `RT` | tuner attachment method, and uninterrupted write throughput during recording |

## Vision, camera & sensing

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **NVR continuous recording** | 24/7 capture from many cameras | `NET` `CAP` `WEAR` `VPU-DEC` | it writes forever, so endurance and capacity come first. Recording camera sub-streams avoids decode entirely |
| **Motion-triggered clip extraction** | cut and retain only events of interest | `VPU-DEC` `IOPS` `CAP` | decode cost per triggered clip; retention policy decides the disk |
| **Camera stream restreaming and proxy** | one pull from the camera, many consumers | `NET` `IDLE` | pure fan-out; keeps cheap cameras from collapsing under multiple clients |
| **Real-time object detection on video** | per-frame inference on live feeds | `NPU` `GPU` `VPU-DEC` | an accelerator turns a four-camera CPU meltdown into a few watts. Decode must also be hardware or it eats the budget by itself |
| **Face and person clustering** | batch embedding then grouping | `MEM` `MT` `NPU` | the first-import burst is the sizing event; steady state is near-idle |
| **Plate and text extraction from video** | ALPR and scene text | `NPU` `MT` `RT` | frame rate times model cost; almost always accelerator-gated |
| **Document OCR ingest** | scanned pages to searchable text | `MT` `THERM` | ingest arrives in bursts that pin every core. Queue depth matters, latency does not |
| **Barcode and QR decode** | symbol reading from a camera feed | `MT` `NPU` `RT` | trivial per frame, but the frame rate sets the floor |
| **Wildlife and trail-cam classification** | filter thousands of empty frames | `NPU` `CAP` | almost all the value is discarding non-events cheaply |
| **Timelapse capture and assembly** | long-interval stills into video | `CAP` `MT` `USB` | capture reliability over months; assembly is a one-off encode |
| **Astrophotography capture and stacking** | long sessions, then heavy registration | `USB` `CAP` `MT` `MEM` `RT` | sustained USB for the camera during capture, then a RAM-hungry stacking pass |
| **Microscopy and scientific imaging** | instrument capture to storage | `USB` `CAP` `RT` | driver support and dropped-frame behaviour |
| **Depth camera and stereo processing** | disparity and depth maps | `USB` `MT` `BW` `RT` | bus bandwidth for two streams, then bandwidth-bound matching |
| **Thermal and non-visible-spectrum pipelines** | IR, UV, multispectral capture | `USB` `RT` `MT` | driver availability more than compute |
| **Machine-vision inspection rig** | camera, trigger, classify, actuate | `RT` `NPU` `USB` `GPIO` | deterministic latency and hardware trigger lines — a co-tenancy decision as much as a hardware one |
| **Point-cloud and LiDAR processing** | SLAM, mapping, registration | `MT` `MEM` `BW` | working-set size; bandwidth-bound once it fits |
| **Radar and mmWave presence sensing** | occupancy without cameras | `GPIO` `USB` `RT` `IDLE` | sensor attachment and placement; compute is negligible |
| **Audio event detection** | glass break, alarm, bark, gunshot | `NPU` `MT` `RT` | always-on inference on a low-rate stream — cheap with an accelerator |
| **SDR capture and decode** | radio to decoded data stream | `USB` `RT` | sustained USB bandwidth with no dropped samples. USB controller topology matters more than CPU |
| **LoRa and mesh radio gateway** | long-range radio to message bus | `IDLE` `USB` `GPIO` | radio attachment and antenna placement decide node placement |
| **Sensor telemetry ingest** | field protocols into a time-series store | `IDLE` `WEAR` | trivial CPU, but constant small writes destroy SD cards |
| **Vibration and predictive maintenance** | continuous high-rate sampling and FFT | `GPIO` `RT` `MT` | sample rate and no gaps; analysis is cheap by comparison |
| **Gas, air-quality and radiation sensing** | slow analogue and digital sensors | `GPIO` `IDLE` `USB` | bus availability; calibration matters more than hardware |
| **Environmental and scientific logging** | long-run continuous capture | `IDLE` `CAP` `RT` | uptime and clock discipline over speed — gaps are unrecoverable |

## Web scraping & data acquisition

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **API polling** | scheduled authenticated pulls | `IDLE` `NET` | effectively free — rate limits and credentials are the constraint, not hardware |
| **Feed ingestion** | conditional GETs against feeds and sitemaps | `IDLE` | hundreds of feeds fit comfortably on the smallest node available |
| **Static fetch and parse** | HTML to structured rows | `MT` `NET` | parsing is cheap CPU; concurrency is capped by file descriptors and politeness, not cores |
| **Headless-browser scraping** | JS-rendered pages and login flows | `MEM` `MT` | roughly 300-700 MB per live context and leak-prone — the RAM ceiling sets worker count. Memory-cap the container and recycle contexts aggressively |
| **Fingerprint-managed sessions** | evade bot detection at scale | `MEM` `NET` | the above plus proxy egress. Isolate it — one node, blast radius contained |
| **CAPTCHA and human-in-the-loop queue** | park blocked requests for resolution | `NET` `IDLE` | external service latency; the node does nothing |
| **Session and cookie-jar management** | thousands of authenticated identities | `MEM` `IOPS` | state size, and durability if sessions are expensive to re-establish |
| **Proxy pool and egress rotation** | outbound IP management | `NET` `IDLE` | connection count, not throughput |
| **Crawl frontier and scheduler** | dedup, politeness, priority queueing | `MEM` `IOPS` | seen-sets and bloom filters in RAM, queue state on real SSD |
| **Sitemap, robots and politeness enforcement** | respect crawl budgets and rules | `IDLE` `NET` | correctness, not hardware |
| **Scraper scheduling and backoff** | orchestrate hundreds of jobs with retry | `IDLE` | reliability of the scheduler node |
| **Broad archival crawl and WARC capture** | full-fidelity page snapshots | `CAP` `NET` `MEM` | write throughput and raw capacity. Browser-driven capture inherits the RAM problem |
| **Change detection and diffing** | fetch, normalize, diff, alert | `IDLE` `IOPS` | store deltas rather than snapshots, or capacity growth becomes the failure mode |
| **Scheduled snapshot and provenance archiving** | keep what was seen and when | `CAP` `WEAR` | append-only growth; plan retention up front |
| **Bulk media and file harvesting** | mass download plus optional remux | `NET` `CAP` | NIC and disk. CPU only enters if re-encoding |
| **Email and inbox ingestion** | harvest structured data from mailboxes | `IOPS` `IDLE` | many small message files |
| **Torrent and DHT crawling** | harvest metadata from peer networks | `NET` `MEM` | connection count and in-memory routing tables |
| **Structured document extraction** | parse PDF, spreadsheets, EPUB, DOCX | `MT` `MEM` | pathological single documents are real OOM risks — cap memory per job, not per worker |
| **LLM-assisted extraction** | model converts messy input to schema | `MEM` `BW` `GPU` | either a local inference node or an outbound API. Decide which before sizing anything else |
| **Translation and language detection** | normalize a multilingual corpus | `MT` `NPU` `MEM` | model size; detection is nearly free, translation is not |
| **Geocoding and address normalization** | text to coordinates and canonical form | `MT` `MEM` `IOPS` | a local geocoder wants the whole gazetteer resident or on fast disk |
| **Full-text index build and serve** | inverted index over a scraped corpus | `MEM` `IOPS` `MT` | index build is a burst; serving wants the index resident in page cache |
| **Dataset normalization, join and dedup** | clean and reconcile harvested data | `MT` `MEM` `IOPS` | columnar engines turn this into a memory-bandwidth problem rather than a core-count one |
| **Entity resolution and record linkage** | fuzzy matching across sources | `MT` `MEM` `BW` | pairwise comparison explodes quadratically — the blocking strategy beats any hardware choice |
| **Data quality and schema-drift monitoring** | detect when a source silently changes | `IDLE` `MT` | scan frequency; the alert matters more than the speed |
| **Historical backfill from archives** | one-off bulk import of the past | `NET` `CAP` `MT` | sustained download and parse; a burst that sizes nothing permanent |

## Data stores & state

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **OLTP relational database** | transactional application backend | `FSYNC` `IOPS` `MEM` `ST` | commit latency is the SSD's sync-write behaviour. Never on SD or USB flash; DRAM-less QLC drives fall off a cliff under sustained writes |
| **Embedded SQLite-backed app** | single-file database inside the app | `IOPS` `FSYNC` | fast on local NVMe, slow and corruption-prone over NFS or SMB — keep it local, always |
| **Read replica and standby** | streaming replication target | `NET` `IOPS` `CAP` | a cheap node, but it still needs a real disk |
| **Connection pooling and proxying** | multiplex client connections | `ST` `IDLE` `NET` | single-threaded event loops, so clock speed — and put it near the database |
| **Sharding and routing proxy** | route queries across shards | `ST` `NET` | per-query overhead on the hot path |
| **Time-series database** | metrics retention and query | `WEAR` `CAP` `MEM` | constant compaction writes. Series cardinality drives RAM, retention drives disk and endurance |
| **In-memory KV and cache** | hot data store | `MEM` `ST` | the dataset must fit RAM, and single-core speed sets operations per second |
| **Cache warming and invalidation** | keep the cache useful | `NET` `MEM` `IDLE` | coordination, not compute |
| **Message broker and durable queue** | reliable event transport | `FSYNC` `NET` `MEM` | persisted queues are fsync workloads. Log-structured brokers expect real disks and real RAM |
| **Object storage** | S3-compatible store | `CAP` `NET` `IOPS` | erasure coding adds CPU; NIC and spindle count set throughput |
| **Blob and attachment front-end** | serve user-uploaded files | `CAP` `NET` | bandwidth and capacity only |
| **Vector database** | approximate nearest-neighbour search | `MEM` `BW` | index resident in RAM; SIMD width matters for distance computation |
| **Columnar and OLAP analytics** | scan-heavy aggregation | `BW` `MT` `CAP` | memory bandwidth first, core count second — the clearest case for wide x86 over ARM |
| **Graph database** | traversal-heavy queries | `MEM` `ST` | pointer chasing; latency-bound and parallelizes poorly |
| **Geospatial database** | spatial indexes and queries | `MT` `MEM` `IOPS` | index residency, and geometry operations are CPU-heavy |
| **Distributed search cluster node** | one shard of a larger index | `MEM` `IOPS` `NET` | JVM heap requirements make RAM the gate on most stacks |
| **Append-only ledger and event store** | immutable event log | `FSYNC` `CAP` | write durability, then unbounded growth |
| **Change-data-capture tap** | stream database changes downstream | `NET` `MEM` `FSYNC` | replication-slot lag will fill the source's disk if the consumer stalls |
| **Materialized view and rollup maintenance** | precompute expensive queries | `MT` `IOPS` `MEM` | refresh window versus refresh cost |
| **Schema migration and backfill** | long-running table rewrites | `IOPS` `MT` `CAP` | on-disk size can double mid-rewrite — headroom is the plan |
| **Retention and deletion sweeps** | enforce data lifecycle | `IOPS` `MT` | delete-heavy passes fragment and amplify writes |
| **Cross-store sync and reconciliation** | keep two systems in agreement | `NET` `MT` | comparison cost across the full key space |
| **Database backup and PITR archiving** | log shipping and logical dumps | `CAP` `NET` `MT` | compression is the CPU cost; restore time is the real metric |
| **Data catalog and lineage tracking** | what exists, where it came from | `MEM` `IOPS` `IDLE` | a small database workload wearing a large name |

## AI & ML inference

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Local LLM serving** | chat and completion endpoint | `MEM` `BW` `GPU` | memory *bandwidth* sets tokens per second and memory *capacity* sets which model fits. Core count is nearly irrelevant, and SBC accelerators rarely help — do not size this by advertised TOPS |
| **Local model gateway and routing** | one endpoint in front of several models | `ST` `NET` `IDLE` | request overhead only |
| **Prompt and result caching** | avoid recomputing identical work | `MEM` `IOPS` | cache hit rate; often the cheapest large speedup available |
| **Embedding generation** | text or image to vector, at scale | `MT` `NPU` `BW` | batchable, and small models are acceptable on CPU |
| **RAG retrieval pipeline** | chunk, embed, retrieve, rerank | `MEM` `BW` `IOPS` | the vector index residency dominates; retrieval quality is a software problem |
| **Reranking and classification** | short model passes per item | `NPU` `MT` `ST` | latency per item, not throughput, on interactive paths |
| **Batch speech-to-text** | media library to transcripts | `MT` `GPU` `MEM` | a GPU is 10-50x; CPU is fine if it can run overnight |
| **Realtime speech-to-text** | live captioning and voice input | `RT` `NPU` `GPU` | the latency budget, not aggregate throughput |
| **Speaker diarization and voice ID** | who spoke when | `MT` `GPU` | usually a second pass over already-transcribed audio |
| **Text-to-speech** | synthesized audio | `MT` `GPU` | lightweight engines are nearly free; neural voices want acceleration |
| **Vision detection and segmentation** | bounding boxes and masks | `NPU` `GPU` | accelerator presence is binary — with one it costs watts, without one it costs a whole node |
| **Multimodal document understanding** | layout-aware extraction from scans | `GPU` `MEM` `NPU` | model size; these are large and unfriendly to small nodes |
| **OCR post-correction with language models** | fix recognition errors in context | `MEM` `BW` | inference cost per page, applied to a whole archive |
| **Image and video generation** | diffusion sampling | `GPU` `MEM` `THERM` | discrete GPU with real VRAM. Not a rack-SBC workload |
| **Video AI post-processing** | upscaling and frame interpolation | `GPU` `MEM` `THERM` | sustained GPU load for hours per title |
| **Music and audio generation** | synthesis and stem models | `GPU` `MEM` `THERM` | as above |
| **Recommendation and scoring service** | small model per request | `ST` `MEM` | tail latency; keep the model resident |
| **Content moderation and safety classification** | flag unwanted content | `NPU` `MT` `RT` | throughput at ingest, latency at post time |
| **Anomaly detection on telemetry** | streaming model over metrics | `MT` `MEM` | window size and series cardinality |
| **Time-series forecasting service** | predict future values | `MT` `MEM` | number of series times model cost |
| **Agent and tool-calling runtime** | orchestrate multi-step model calls | `MEM` `NET` `IDLE` | mostly waiting on I/O — a cheap node unless inference is local |
| **Model quantization and conversion** | reformat weights for deployment | `MT` `MEM` `CAP` | a one-off job that briefly needs several times the model's size in RAM and disk |
| **Fine-tuning and LoRA training** | short supervised training runs | `GPU` `MEM` `THERM` | listed mainly so it can be excluded — this is workstation or rented-GPU work, not a rack node |
| **Model and weight distribution** | registry for multi-GB artifacts | `CAP` `NET` | pull bandwidth. Keep it on the same switch as the inference nodes |
| **Prompt and eval harness runs** | batch evaluation over a test set | `MT` `NET` | usually API-bound, occasionally GPU-bound |

## Build, CI & software supply chain

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **CI runner** | executes pipelines | `MT` `MEM` `IOPS` `THERM` | all-core bursts on a duty cycle. Job concurrency is a RAM-per-job calculation, and the disk churn destroys cheap flash |
| **Native-architecture container builds** | build arm64 on arm64, amd64 on amd64 | `MT` `IOPS` `x86` | cross-architecture emulation runs 10-40x slower. Keeping a matching-arch node is often the *entire* justification for that node existing |
| **Multi-arch image assembly** | manifest lists across architectures | `NET` | cheap, once each per-arch build is native |
| **Cross-compilation and emulation host** | build for a foreign architecture | `MT` `MEM` | the slow fallback when no native node exists |
| **Compile-heavy builds** | long compile and link trees | `MT` `MEM` `THERM` | RAM per core, because linking spikes; sustained clocks beat peak boost |
| **Distribution and firmware image builds** | full image assembly from source | `MT` `CAP` `IOPS` | tens of GB of intermediates — capacity is the surprise, not CPU |
| **Golden-image bakery** | produce reusable machine images | `VIRT` `CAP` `MT` | needs virtualization to boot and provision what it bakes |
| **Compile cache and artifact server** | shared build cache | `NET` `IOPS` `CAP` | a fast NIC converts a fleet's rebuilds into cache hits |
| **Distributed compile for a fleet** | farm out compilation units | `MT` `NET` `THERM` | network round-trip versus compile time; only wins on large translation units |
| **Package registry mirror and pull-through cache** | local mirror of upstream | `CAP` `NET` | disk, plus enough uptime that builds do not fall back to the WAN |
| **Container image registry and garbage collection** | store and prune images | `CAP` `NET` `IOPS` | layer churn; GC passes are IOPS-heavy |
| **Git forge** | repository hosting plus web UI | `MEM` `IOPS` | the specific forge changes the answer by an order of magnitude — lightweight ones fit anywhere, heavyweight ones want several GB before they are comfortable |
| **Static site and documentation generation** | site builds from source | `MT` `IOPS` | thousands of small file operations |
| **API-spec and reference generation** | generate docs from code | `MT` `IOPS` | proportional to codebase size |
| **Frontend bundling** | JS toolchain builds | `MT` `MEM` | Node-based build steps are RAM-hungry per worker |
| **Binary and asset optimization** | minify, subset fonts, recompress images | `MT` | parallel and cheap |
| **Localization pipeline builds** | extract, merge and compile translations | `MT` `IOPS` | file count, not size |
| **Static analysis and type checking at scale** | whole-program analysis | `MT` `MEM` | large projects hold the whole type graph in memory |
| **Coverage aggregation and reporting** | merge and render coverage data | `MT` `IOPS` | report size on large test suites |
| **Benchmark and regression runner** | measure performance over time | `THERM` `RT` `MT` | it must be *uncontended* or the results are noise. A dedicated node with a fixed governor, or do not bother |
| **Fuzzing and property-based testing** | long randomized runs | `MT` `THERM` `CAP` | sustained load for days; corpus growth eats disk |
| **Reproducible-build verification** | rebuild and compare bit-for-bit | `MT` `CAP` | double the build cost, by definition |
| **Test and device farm host** | drives physically attached devices under test | `USB` `RT` `GPIO` | port count and stable enumeration, not compute |
| **Nightly and canary environments** | full stacks stood up on a schedule | `MEM` `VIRT` `CAP` | how much of production has to fit |
| **Dependency and vulnerability scanning** | scheduled scans of images and repos | `MT` `NET` `CAP` | metadata pulls dominate; bursty and noisy |
| **License and compliance scanning** | audit dependency licensing | `MT` `NET` | same shape, different database |
| **Release signing, SBOM and provenance** | sign and attest build outputs | `CRYPTO` `IDLE` | a small workload that wants an isolated, minimal-trust node — a placement decision, not a performance one |

## Networking & edge

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Recursive DNS and filtering** | LAN resolver plus blocklists | `IDLE` `RT` | computationally trivial. Its real requirement is that it never goes down — put it on the most boring node available, and run two |
| **Authoritative DNS** | serve your own zones | `IDLE` `NET` | query rate; almost always negligible |
| **Split-horizon and conditional forwarding** | different answers by client or zone | `IDLE` | configuration correctness only |
| **Dynamic DNS updater** | keep a changing address published | `IDLE` | nothing |
| **DNS-layer threat filtering** | block known-bad domains | `IDLE` `MEM` | blocklist size resident in memory |
| **Reverse proxy and TLS termination** | ingress front door | `CRYPTO` `NET` `ST` | handshake cost is the load. Hardware crypto extensions make it a non-issue; their absence makes it the bottleneck |
| **WAF and application-layer filtering** | inspect and block HTTP payloads | `MT` `NET` `CRYPTO` | rule count times request rate, on top of TLS |
| **Load balancer and L4 proxy** | traffic distribution and health checking | `PPS` `NET` `ST` | connection rate and packet rate |
| **CDN edge and caching proxy** | cache hot objects near clients | `MEM` `CAP` `NET` | cache hit rate, which is a RAM question first and a disk question second |
| **Router and firewall** | L3 forwarding, NAT and policy | `PPS` `NET` `x86` | small-packet rate, NIC quality and offload support — not clock speed. The NIC matters more than the CPU here |
| **IDS/IPS and deep packet inspection** | inspect every packet in path | `MT` `PPS` `MEM` | roughly a core per few hundred Mbps inspected — the usual reason a firewall node ends up oversized |
| **DDoS scrubbing and rate shaping** | absorb and drop abusive traffic | `PPS` `NET` | drop rate at the earliest possible point in the stack |
| **QoS and traffic shaping** | prioritize latency-sensitive flows | `PPS` `RT` | queue discipline accuracy under load |
| **VPN gateway and mesh** | encrypted transit | `CRYPTO` `ST` `NET` | throughput is crypto-per-core. Modern kernel VPNs are fast where the extensions exist and painful where they do not |
| **Mesh control plane and coordination** | key distribution and ACLs for a mesh | `IDLE` `MEM` | node count; the data plane is elsewhere |
| **Tunnel broker and NAT-traversal relay** | reach clients behind NAT | `NET` `CRYPTO` | relayed bandwidth and a reachable address |
| **Anonymity network relay** | volunteer transport | `NET` `IDLE` | bandwidth-bound with near-zero CPU — the canonical cheapest-node job |
| **SSH bastion and jump host** | single audited entry point | `IDLE` `CRYPTO` | almost nothing; ideal for the lowest tier of hardware |
| **Multi-WAN, failover and bonding** | link aggregation and failover | `NET` `PPS` | NIC count above all |
| **BGP speaker and anycast node** | participate in routing | `MEM` `RT` `PPS` | full tables need real RAM; a default-only speaker needs none |
| **Flow collection and traffic analytics** | who talked to whom, retained | `MT` `CAP` | flow rate and retention window |
| **Full packet capture** | line-rate capture to disk | `CAP` `NET` `WEAR` | writing at line rate is the hard part; ring buffer sizing is the design |
| **Bandwidth and path monitoring** | continuous reachability and latency probes | `NET` `IDLE` | must be independent of the path it measures |
| **WiFi controller** | manage access points centrally | `MEM` `IDLE` `IOPS` | some controllers are JVM applications and want real RAM despite doing almost nothing |
| **Netboot and provisioning** | boots the rest of the rack | `IDLE` `CAP` | it must be up *before* everything else — dependency ordering beats specifications |
| **DHCP, IPAM and NTP** | core LAN services | `IDLE` `RT` | reliability only |
| **Stratum-1 time source** | GPS-disciplined local clock | `RT` `GPIO` | needs a PPS-capable pin or serial line plus antenna placement — a hard hardware filter, not a preference |
| **Captive portal and guest isolation** | segmented visitor access | `IDLE` `NET` | VLAN support in the path |
| **Network device config backup** | pull and version switch and router configs | `IDLE` | trivial; the value is in the versioning |

## Storage, backup & sync

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **File serving** | SMB and NFS shares | `NET` `CAP` `IOPS` | saturating a fast NIC requires both the NIC *and* an array behind it that can feed it |
| **Protocol translation gateway** | re-export one protocol as another | `NET` `MT` | translation overhead on every operation |
| **Checksummed pool host** | integrity, snapshots, send and receive | `MEM` `CAP` `PCIE` | RAM per TB for caching, scrubs are sustained reads, and HBA attachment needs lanes |
| **Storage tiering and caching** | SSD in front of spinning disk | `IOPS` `WEAR` | cache device endurance, and what happens when it fails |
| **Encrypted volume host** | at-rest encryption | `CRYPTO` `IOPS` | without hardware crypto this halves throughput |
| **Deduplicating backup repository** | chunked, compressed, encrypted backups | `MT` `MEM` `CAP` `CRYPTO` | chunking, compression and encryption are all CPU. Prune and check passes are heavier than the backups themselves |
| **Immutable and WORM backup target** | append-only, tamper-resistant copies | `CAP` `FSYNC` | the guarantee is the point; capacity follows |
| **Air-gapped rotation target** | offline copies rotated by hand | `CAP` `IDLE` | physical process, not hardware |
| **Bare-metal image backup and restore** | whole-machine images | `NET` `CAP` `MT` | restore speed is the metric nobody measures until it matters |
| **Offsite replication** | push copies elsewhere | `NET` `CAP` `CRYPTO` | upstream bandwidth and encryption throughput |
| **Cloud gateway and object tiering** | spill cold data to remote object storage | `NET` `CAP` `CRYPTO` | egress cost and recall latency, not local hardware |
| **Restore drills and verification** | periodically prove restores work | `MT` `NET` `CAP` | the step everyone skips — size a node that can actually *hold* a full restore |
| **Snapshot and retention orchestration** | schedule, prune, verify | `IDLE` `IOPS` | correctness, not speed |
| **File sync server** | multi-device sync | `IOPS` `MEM` `CAP` | many small files means metadata IOPS. Full-featured suites also need a real database behind them |
| **Version-controlled large-file store** | binary assets under version control | `CAP` `NET` `IOPS` | history grows without bound unless pruned |
| **Content-addressed store** | hash-addressed blocks | `IOPS` `NET` `CAP` | block count more than byte count |
| **Peer-to-peer transfer client** | high-connection-count transfer | `NET` `IOPS` `CAP` | thousands of concurrent connections, so raise descriptor limits. Repair and extraction are CPU spikes |
| **Cold archive and spin-down array** | rarely-read bulk storage | `CAP` | capacity, acoustics and idle power |
| **Duplicate finder and reclamation** | find and remove redundant copies | `MT` `IOPS` | a full-corpus hash sweep |
| **Media integrity verification** | parity and checksum validation | `MT` `CRYPTO` `IOPS` | read throughput times corpus size |
| **Filesystem indexing and search** | make local files findable | `IOPS` `MEM` `MT` | index build is the burst; queries are free |
| **Automated file routing and sorting** | watch folders and file by rule | `IDLE` `IOPS` | event latency; the work itself is trivial |
| **Quota and capacity reporting** | who is using what | `IDLE` `IOPS` | directory-walk cost on large trees |
| **Disk health and surface monitoring** | SMART, scrubs, badblock sweeps | `IDLE` | scheduling; the sweep runs in the background |
| **Bulk data migration** | one-off multi-TB moves | `NET` `CAP` `MT` | sustained throughput at both ends simultaneously |

## Orchestration & virtualization

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Cluster control plane** | the cluster's brain and its consensus store | `FSYNC` `RT` `MEM` | the consensus store is fsync-latency-bound and *will* fail on SD cards and cheap USB media. This is the single most common homelab cluster mistake |
| **Cluster worker node** | runs the workloads | `MT` `MEM` | RAM sets pod density. Leave scheduling headroom or eviction storms follow |
| **Plain container host** | no orchestrator, just containers | `MEM` `CAP` `IOPS` | image layers consume disk faster than anyone plans for |
| **Full virtualization host** | VMs with real isolation | `VIRT` `MEM` `CAP` | virtualization extensions plus RAM to overcommit. Nested virtualization narrows the hardware list much further |
| **GPU passthrough and vGPU host** | hand a GPU to a guest | `VIRT` `PCIE` `GPU` `MEM` | IOMMU grouping decides feasibility, and that is a chipset property no spec sheet lists. Verify on the exact board |
| **USB and PCIe device passthrough** | hand peripherals to a guest | `VIRT` `PCIE` `USB` | controller topology; devices sharing a group move together or not at all |
| **Windows and legacy application host** | run software that will not port | `VIRT` `MEM` `x86` | architecture is a hard requirement, not a preference |
| **Thin-client and VDI backend** | remote desktops for users | `VIRT` `MEM` `GPU` `NET` | RAM per session, and encode capacity if the sessions are graphical |
| **Multi-tenant sandbox** | run untrusted code | `VIRT` `MEM` `RT` | isolation is a hardware question — microVMs or equivalent, on a node you are willing to lose |
| **Distributed cluster storage** | replicated block or filesystem | `NET` `IOPS` `MEM` `WEAR` | replication multiplies both network traffic and write volume — the usual reason a homelab cluster feels slow |
| **Service discovery and coordination** | cluster consensus and lookup | `FSYNC` `RT` | the same fsync story as the control plane |
| **Ingress controller** | cluster edge routing | `NET` `CRYPTO` | TLS throughput |
| **Policy engine and admission control** | evaluate rules on every request | `ST` `MEM` `RT` | it sits on the critical path of every deployment |
| **Autoscaling controller** | add and remove capacity | `IDLE` `NET` | decision latency, not compute |
| **Node lifecycle and image rollout** | update the fleet's OS | `NET` `CAP` | image serving bandwidth during a rollout |
| **GitOps and multi-cluster controller** | reconcile declared state continuously | `MEM` `NET` `IOPS` | repository size times reconcile frequency |
| **Cluster state backup and DR** | snapshot and restore cluster objects | `NET` `CAP` | restore correctness; the data is small |
| **Configuration management control node** | push state to the fleet | `MT` `NET` | bursty fan-out, typically a process per target host |
| **Infrastructure-as-code runner** | plan and apply pipelines | `MEM` `NET` | state locking wants a reliable backend more than a fast CPU |
| **Bare-metal provisioning** | image and enrol new machines | `NET` `CAP` | image serving bandwidth |
| **Workflow and DAG orchestration** | scheduled and event-driven pipelines | `MEM` `FSYNC` | the scheduler is small; the workers are the actual load |
| **Function and serverless runtime** | scale-to-zero handlers | `MEM` `ST` | cold-start latency is single-thread-bound |
| **Secrets injection and workload identity** | short-lived credentials for workloads | `IDLE` `CRYPTO` | availability — everything stalls when it is down |
| **Scheduler and cron host** | the fleet's crontab | `IDLE` `RT` | clock accuracy and uptime |

## Observability & operations

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Metrics scraping and storage** | pull, store, downsample | `MEM` `WEAR` `CAP` | active series cardinality drives RAM; retention drives disk and endurance |
| **Log aggregation** | ingest, index and retain | `MT` `CAP` `WEAR` | compression is the CPU cost; ingest rate is a sustained write workload |
| **Log parsing and enrichment** | structure and annotate log lines | `MT` `MEM` | regex cost per line times line rate |
| **Distributed tracing** | span collection and query | `CAP` `MEM` | very high volume per unit of value — sample aggressively |
| **Continuous profiling** | always-on CPU and memory profiles | `MT` `CAP` `MEM` | profile volume; the agents themselves are cheap |
| **eBPF-based observability** | kernel-level instrumentation | `MT` `MEM` `RT` | kernel version and build options are the gate, not the CPU |
| **Dashboards and query frontend** | visualize and alert | `IDLE` `MEM` | light — the datasource does the real work |
| **Uptime and synthetic monitoring** | probe endpoints from outside | `IDLE` `NET` | it must be *independent* of what it watches. A deliberate second node, never a co-tenant |
| **Real-user and browser monitoring** | measure from a real client stack | `MEM` `NET` | headless browsers again, with their RAM profile |
| **Alert routing and notification** | dedupe, route, escalate | `IDLE` | reliability only |
| **On-call scheduling and escalation** | who gets woken up | `IDLE` | availability during incidents |
| **Status page** | externally visible health | `IDLE` `NET` | it must survive the outage it is reporting |
| **SLO and error-budget accounting** | track reliability targets | `IDLE` `MEM` | query cost against the metrics store |
| **Capacity forecasting** | project growth and exhaustion | `MT` `MEM` | history length times series count |
| **Cost and power accounting** | attribute spend and watts to workloads | `IDLE` | metering granularity |
| **SNMP, IPMI and PDU polling** | infrastructure telemetry | `IDLE` | trivial |
| **Out-of-band management** | console and video access when the OS is down | `USB` `RT` `GPIO` | a hard hardware requirement — USB device-mode plus video capture, or a serial fan-out. Not a preference |
| **Power and environmental monitoring** | UPS state, temperature, current | `IDLE` `USB` `GPIO` | the attachment method for the sensors |
| **Physical status display** | at-a-glance rack readout | `IDLE` `GPIO` | needs an actual display bus — SPI, I2C, DSI or HDMI |
| **Configuration drift detection** | compare running state to declared state | `IDLE` `MT` | scan frequency |
| **Post-incident data capture** | auto-snapshot state when an alert fires | `IOPS` `CAP` | write burst at the worst possible moment |
| **Audit trail collection and retention** | durable record of who did what | `CAP` `FSYNC` | write durability, then retention |
| **Asset and inventory tracking** | the source of truth for what exists | `MEM` `IOPS` | a small database workload |
| **Firmware and BIOS update orchestration** | keep the fleet's firmware current | `IDLE` `NET` | vendor tooling availability, which often means x86 |
| **Runbook automation and incident tooling** | scripted response | `IDLE` | availability during incidents |
| **Chaos and failure injection** | deliberately break things | `IDLE` | blast-radius control |

## Home automation, IoT & physical control

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Home automation hub** | the central automation engine | `RT` `USB` `MEM` | latency-sensitive *and* radio-attached. USB passthrough for radio dongles constrains the node choice and complicates virtualization |
| **Radio coordinator bridges** | Zigbee, Z-Wave, Thread and Matter to a message bus | `USB` `RT` `IDLE` | USB stability and the physical placement of the stick matter far more than CPU |
| **MQTT broker** | pub/sub bus for devices | `IDLE` `NET` | connection count at large scale, otherwise free |
| **Rules and flow engine** | event-driven automation logic | `IDLE` `MEM` | trivial |
| **Device firmware build host** | compiles firmware for edge devices | `MT` `CAP` | a genuine compile workload hiding inside a home-automation stack — usually the only heavy component |
| **OTA firmware distribution** | push updates to fleets of devices | `NET` `CAP` `IDLE` | concurrent download fan-out |
| **Device provisioning and onboarding** | enrol new devices securely | `IDLE` `NET` | process, not hardware |
| **Local voice pipeline** | wake word, recognition, intent, speech | `RT` `NPU` `MEM` | end-to-end latency budget across four models. The weakest stage sets the experience |
| **Voice assistant satellite** | microphone endpoint feeding the pipeline | `RT` `NPU` `USB` | microphone hardware and network latency |
| **Presence and BLE tracking** | detect who and what is nearby | `IDLE` `USB` | radio placement drives node placement |
| **Media and scene control bridges** | IR, HDMI-CEC and RF control | `GPIO` `USB` `RT` | emitter and receiver attachment |
| **Local dashboard and wall panel** | always-on display of house state | `IDLE` `GPU` `VPU-DEC` | needs display output; a browser in kiosk mode is heavier than it looks |
| **Digital signage and kiosk** | full-screen looping display | `VPU-DEC` `GPU` `IDLE` | needs display output and hardware decode — disqualifies headless nodes outright |
| **Print and scan server** | shared peripherals | `USB` `IDLE` | driver support and port count |
| **3D printer host** | motion planning plus web UI | `RT` `USB` `VPU-ENC` | step generation wants deterministic latency, and webcam streaming adds an encoder requirement most people forget |
| **CNC, laser and plotter control** | stream motion commands to hardware | `RT` `USB` `GPIO` | jitter causes physical defects — an uncontended node |
| **Robotics middleware node** | sensor fusion and control loops | `RT` `MT` `USB` `GPIO` | a realtime-capable kernel, CAN and GPIO availability, deterministic scheduling |
| **Hardware bring-up and flashing station** | JTAG, SWD, serial and chip programmers | `USB` `RT` `GPIO` | port count and reliable enumeration under repeated replug |
| **Access control and physical security** | doors, gates and relays | `GPIO` `RT` `IDLE` | fail-safe behaviour on power loss is the design question |
| **Alarm and siren control** | local-first alerting | `GPIO` `RT` | it must work when the network is down |
| **Irrigation, climate and greenhouse control** | scheduled actuation with sensor feedback | `GPIO` `RT` `IDLE` | reliability — a missed cycle has physical consequences |
| **Aquarium, incubator and lab control** | closed-loop environmental control | `GPIO` `RT` | watchdog behaviour and failure modes |
| **Pet and livestock monitoring** | feeding, cameras and sensors | `GPIO` `USB` `IDLE` | attachment variety |
| **Energy monitoring** | solar, current clamps and inverters | `IDLE` `USB` `WEAR` | sample rate times retention |
| **Smart meter reading** | utility meter interfaces | `GPIO` `USB` `IDLE` `WEAR` | serial or optical attachment; constant small writes |
| **Vehicle telemetry and OBD logging** | capture from vehicle buses | `USB` `GPIO` `CAP` `IDLE` | CAN interface availability and power behaviour on ignition cycles |

## Custom application runtime patterns

The shapes custom software takes, independent of what it does. Most bespoke services are one or more of these, and this is usually the more useful lens for sizing a node than the name of the product running on it.

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Compiled API service** | HTTP or gRPC backend in a compiled language | `ST` `NET` | tens of MB of RAM and thousands of connections — the cheapest thing in the rack to co-locate |
| **Interpreted application server** | request handling in a scripting runtime | `MEM` `ST` | RAM-per-worker times worker count is the entire sizing formula. Interpreter concurrency models make single-thread speed matter more than core count |
| **Managed-runtime service** | long-lived JVM or .NET application | `MEM` `BW` | heap sizing dominates everything else |
| **Background job workers** | asynchronous work off the request path | `MT` `MEM` | horizontally scalable — the archetypal "just add another cheap node" workload |
| **Scheduled batch and ETL jobs** | periodic pipelines | `MT` `IOPS` `MEM` | duty-cycled, so peak matters and average is irrelevant |
| **Stream processing** | continuous transform of an event stream | `MT` `MEM` `NET` | consumer lag and backpressure, not raw record throughput |
| **WebSocket and realtime fan-out** | many long-lived connections | `MEM` `NET` | file-descriptor and ephemeral-port limits are hit long before CPU — tune ulimits first |
| **Long-poll and SSE endpoints** | held-open HTTP connections | `MEM` `NET` | connection count, same as above |
| **Webhook receiver and event relay** | inbound HTTP into a durable queue | `IDLE` `NET` | reachability and durability, not speed |
| **Persistent bot and protocol clients** | always-connected sessions | `IDLE` `MEM` | a pure uptime workload |
| **API gateway and rate limiting** | quota and policy enforcement | `ST` `MEM` `NET` | shared counters want a fast store on the same node or the same switch |
| **Session and auth token service** | issue and validate credentials | `ST` `CRYPTO` `MEM` | signature verification rate on every request |
| **Feature flag and dynamic config** | low-latency config reads | `ST` `MEM` | tail latency; everything depends on it |
| **Idempotency and deduplication layer** | make retries safe | `MEM` `FSYNC` | the dedup window size, and whether it must survive a restart |
| **Saga and transaction coordinator** | multi-step distributed transactions | `FSYNC` `RT` `MEM` | durable state transitions — an fsync workload wearing application clothes |
| **Audit and event-sourcing store** | append-only record of everything | `FSYNC` `CAP` | write durability, then unbounded growth |
| **Payment, billing and metering** | usage capture and invoicing | `FSYNC` `IDLE` `MEM` | correctness and durability; volume is usually trivial |
| **File upload and ingestion endpoint** | accept large inbound files | `NET` `CAP` `IOPS` | concurrent upload count times file size |
| **On-demand media processing** | resize, watermark and convert per request | `MT` `MEM` `VPU-ENC` | cache aggressively or this becomes a transcoding node by accident |
| **Headless render and report generation** | server-side documents, charts, screenshots | `MEM` `MT` | usually a headless browser underneath, so it inherits that RAM profile |
| **Scheduled report and export generation** | batch exports on a timer | `MT` `MEM` `CAP` | result-set size held in memory during generation |
| **Notification and digest fan-out** | build and send many personalized messages | `MT` `NET` `IOPS` | per-recipient work times recipients |
| **Feed generation and timeline assembly** | fan-out on write or on read | `MEM` `IOPS` `MT` | which fan-out model you chose — it inverts the bottleneck completely |
| **Search and recommendation backend** | query, rank and return | `MEM` `ST` `IOPS` | index residency |
| **Realtime collaboration backend** | merge concurrent edits | `MEM` `RT` `ST` | per-document state held in memory times concurrent documents |
| **Geospatial processing and tile serving** | spatial queries and map tiles | `MT` `MEM` `IOPS` `CAP` | render-on-demand is CPU, pre-rendered tiles are capacity. Pick one before sizing |
| **Localization and content delivery** | serve per-locale content | `MEM` `NET` | catalogue size in memory |
| **Multi-tenant application backend** | isolation between customers | `MEM` `IOPS` `VIRT` | the isolation model — shared process, shared database, or shared nothing — changes the answer entirely |
| **Plugin and extension sandbox host** | run third-party code safely | `MEM` `ST` `VIRT` | per-instance memory floor times tenant count |
| **Experiment assignment and metrics** | A/B bucketing and measurement | `ST` `MEM` | assignment is on the hot path; analysis is not |
| **Deployment controller** | canary and blue-green rollouts | `IDLE` `NET` | coordination only |
| **Shadow-traffic dual-run** | run old and new side by side | `MT` `NET` `MEM` | doubles the load of whatever it is shadowing, by definition |
| **Rate-limited outbound integration worker** | talk to slow third-party APIs | `IDLE` `NET` | almost entirely waiting; a cheap node with good uptime |
| **Data validation and integrity sweeps** | verify a corpus against rules | `MT` `IOPS` | read throughput |
| **Bulk cryptographic work** | hashing, signing and checksumming at scale | `CRYPTO` `MT` | dedicated crypto instructions swing this by an order of magnitude |
| **Compression and recompression at scale** | bulk archive or asset compression | `MT` `THERM` `CAP` | sustained all-core, and algorithm choice changes it by 10x |
| **Simulation and numerical solvers** | numeric batch compute | `MT` `BW` `THERM` | vector width and sustained clocks; memory bandwidth caps the scaling |
| **Parameter sweeps and evolutionary search** | embarrassingly parallel exploration | `MT` `THERM` | an ideal fleet workload — scales by adding nodes and needs no fast interconnect |
| **Time-critical scheduling** | auctions, bidding and market data | `RT` `ST` `NET` | tail latency is the product. Dedicate the node and pin the cores |
| **Long-running stateful session host** | in-memory state per connected user | `MEM` `RT` | state per session times sessions, and what happens on restart |

## Communication & collaboration

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Federated chat homeserver** | protocol server plus federation | `MEM` `IOPS` `NET` | implementation choice swings this enormously — the reference server is RAM- and database-hungry, lightweight alternatives fit on a small node |
| **Lightweight chat services** | small always-on chat infrastructure | `IDLE` | free |
| **Livestream chat and moderation** | high-rate message handling | `MEM` `NET` `ST` | message rate |
| **Community moderation queues** | review and action reported content | `MEM` `MT` | queue depth; ML-assisted triage adds inference cost |
| **Mail server** | send, receive and store mail | `CAP` `IOPS` `IDLE` | deliverability is an address-reputation and DNS problem, not a hardware one. The hardware load is storage and small-file IOPS |
| **Mail filtering and antivirus** | spam scoring and malware scanning | `MEM` `MT` | signature databases alone can want a gigabyte or more resident |
| **Bulk and transactional sending** | queued outbound mail at volume | `NET` `IDLE` `IOPS` | queue durability and send-rate limits |
| **Push notification gateway** | relay to mobile and desktop endpoints | `IDLE` `NET` | connection count to upstream services |
| **SMS and telephony gateway** | modem or SIP trunk interface | `USB` `GPIO` `RT` `IDLE` | hardware attachment for a modem; nothing for a trunk |
| **VoIP and PBX** | call routing and media handling | `RT` `NET` | jitter and packet loss. Codec transcoding adds real CPU |
| **Video conferencing SFU** | route participant streams | `NET` `MT` `RT` | bandwidth first — these forward rather than transcode by design, which is why they scale on modest CPUs |
| **Call and meeting recording** | capture and store sessions | `CAP` `VPU-ENC` `NET` | concurrent recordings times bitrate |
| **Meeting transcription and translation** | live or post-hoc text from audio | `MT` `GPU` `RT` | realtime raises it from a batch job to an accelerator job |
| **Screen-sharing and remote desktop gateway** | broker interactive sessions | `VPU-ENC` `NET` `RT` | encode capacity per concurrent session |
| **Groupware** | calendar and contact sync | `IDLE` | free |
| **Scheduling and booking service** | availability and reservations | `IDLE` `IOPS` | small |
| **Collaborative editing** | realtime shared documents | `MEM` `RT` | concurrent document count |
| **Wiki and knowledge base** | documentation hosting | `IDLE` `IOPS` | small; search indexing is the only real cost |
| **Forum and community platform** | threaded discussion hosting | `MEM` `IOPS` | full-featured platforms expect a real database plus a cache and several GB of RAM |
| **Fediverse instance** | federated social server | `MEM` `IOPS` `NET` | federation fan-out is chatty and unpredictable. Lightweight implementations exist and change the sizing entirely |
| **Team file and link sharing** | pastebins, file drops, short links | `CAP` `IDLE` | retention policy |
| **Feedback, helpdesk and ticketing** | track and resolve requests | `MEM` `IOPS` | a standard small web application |
| **Broadcast and announcement fan-out** | one message to many recipients | `NET` `MT` | recipient count |
| **Multiplayer game server** | authoritative game simulation | `ST` `MEM` `RT` | tick rate is single-thread-bound. High clocks beat many cores, essentially always |

## Security, identity & privacy

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Identity provider and SSO** | central authentication | `MEM` `ST` `RT` | every request touches it, so latency and uptime. Implementations range from trivial to JVM-heavy |
| **Directory service** | the user and group database | `MEM` `IOPS` `RT` | lookup latency under load |
| **MFA and hardware token backend** | second-factor verification | `CRYPTO` `IDLE` `RT` | verification rate; negligible until it is on every request |
| **Network access control** | port-level authentication | `IDLE` `RT` `CRYPTO` | authentication latency at link-up |
| **Zero-trust and identity-aware proxy** | authenticate every request to every service | `CRYPTO` `NET` `ST` | it sits in front of everything, so it inherits every service's traffic |
| **Guest and temporary credential issuance** | short-lived access | `IDLE` | nothing |
| **Password and secret vault for humans** | credential storage | `IDLE` `IOPS` | featherweight to run. Back it up as though it is the only thing that matters |
| **Machine secrets and PKI** | issue and rotate machine credentials | `FSYNC` `CRYPTO` `IDLE` | seal and unseal behaviour, and audit durability. A TPM or secure element is a genuine differentiator |
| **Key ceremony and offline signing** | root key operations | `IDLE` `CRYPTO` | the requirement is physical and procedural — an air-gapped node, deliberately underpowered |
| **Certificate automation** | keep TLS valid everywhere | `IDLE` | reliability only |
| **Certificate transparency and domain monitoring** | watch public logs for your names | `NET` `CAP` | a high-volume feed with tiny compute cost |
| **Attack surface monitoring** | scan yourself from outside | `NET` `MT` | scan breadth |
| **Vulnerability scanning** | active probing of hosts and applications | `MT` `NET` | bursty and noisy — isolate its egress or it will trip your own defences |
| **Secret and credential scanning** | sweep repositories and artifacts | `MT` `IOPS` | corpus size |
| **Data loss prevention scanning** | find sensitive data where it should not be | `MT` `IOPS` | full-corpus content inspection |
| **Threat intelligence ingestion and matching** | consume feeds and match against traffic | `MEM` `MT` `NET` | indicator set resident in memory |
| **SIEM and security monitoring** | correlate events, detect, alert | `MEM` `CAP` `MT` | one of the heaviest self-hosted stacks — size it as a log system with detection CPU on top |
| **Endpoint and host intrusion detection** | file integrity and audit rules | `IDLE` `IOPS` | audit event volume |
| **Honeypot and deception node** | deliberately exposed bait | `IDLE` `VIRT` | assume compromise: the cheapest node, an isolated segment, nothing else on it |
| **Malware sandbox and detonation** | run hostile samples safely | `VIRT` `MEM` `RT` | nested virtualization and fast snapshot rollback. Isolation is a hardware property here |
| **Forensic imaging and analysis** | acquire and examine evidence | `CAP` `MT` `IOPS` `x86` | tooling is largely x86, and images are large |
| **Security lab and training range** | attacker and target machines | `VIRT` `MEM` `CAP` | how many VMs run at once |
| **Privacy relay and anonymizing proxy** | strip identifiers in transit | `NET` `CRYPTO` | throughput per core |
| **Egress-isolated processing** | handle sensitive data offline | `IDLE` `VIRT` | the selection criterion is network topology and physical control, not specifications |
| **Tamper-evident audit archive** | append-only retention with proofs | `CAP` `FSYNC` | write durability |
| **Compliance evidence collection** | gather and retain proof of controls | `IDLE` `CAP` | retention period |
| **Backup and recovery of security state** | keys, policies and identity data | `CAP` `FSYNC` `CRYPTO` | the recovery path matters more than the backup; test it |

## Batch, scientific & specialty compute

| Task | What it is | Stresses | What actually decides it |
|---|---|---|---|
| **Volunteer distributed computing** | donate spare cycles | `MT` `THERM` | thermals and the power bill, nothing else |
| **Blockchain full node** | chain sync and validation | `IOPS` `CAP` `WEAR` `NET` | initial sync is brutally random-I/O heavy and genuinely wears out consumer SSDs. Capacity grows monotonically and forever |
| **Proof-of-work hashing** | mining | `GPU` `THERM` | listed for completeness; the economics rarely work in a home rack |
| **Offline render farm node** | frame or tile rendering | `MT` `GPU` `MEM` `THERM` | frame-parallel, so it scales linearly across cheap nodes — one of the best fleet workloads there is |
| **Distributed batch frameworks** | map, shuffle and reduce across nodes | `MEM` `NET` `CAP` | shuffle is a network and disk problem, not a CPU one |
| **Data-science notebooks** | interactive analysis | `MEM` `BW` | interactive RAM spikes are unpredictable by nature — overprovision or hard-cap |
| **Statistical model training** | classical ML and gradient boosting | `MT` `MEM` `BW` | dataset residency |
| **Optimization and constraint solvers** | scheduling, routing and planning | `ST` `MEM` | many solvers are largely single-threaded, so clock speed |
| **Financial backtesting and simulation** | replay markets against strategies | `MT` `MEM` `IOPS` | tick data volume |
| **Climate and weather model runs** | numerical forecasting | `MT` `BW` `MEM` `THERM` | memory bandwidth, then interconnect past one node |
| **Computational chemistry and physics** | solver runs over physical systems | `MT` `BW` `MEM` `THERM` | as above |
| **Protein folding and molecular dynamics** | simulate molecular behaviour | `GPU` `MT` `BW` `THERM` | GPU-accelerated in practice |
| **Bioinformatics pipelines** | alignment, assembly and variant calling | `MEM` `MT` `CAP` | per-job RAM in the tens of GB — capacity, not core count, is the gate |
| **Astronomical data reduction** | calibrate and stack survey data | `MT` `MEM` `CAP` | image count times frame size |
| **Graph analytics at scale** | centrality, clustering, traversal | `MEM` `BW` `MT` | the graph must fit; partitioning is painful |
| **Text corpus processing and NLP** | tokenize, annotate and index at scale | `MT` `MEM` `IOPS` | corpus size and pipeline depth |
| **Signal and audio DSP batch** | filtering, transforms, feature extraction | `MT` `BW` | vectorizable and bandwidth-bound |
| **Mesh generation and CAD batch** | geometry processing | `MT` `MEM` | model complexity |
| **Synthetic data generation** | manufacture training data | `MT` `GPU` `CAP` | output volume |
| **Archive analysis at scale** | run models over a whole media library | `MT` `NPU` `CAP` `IOPS` | a full read of everything, once, per model version |
| **Password auditing and cryptanalysis** | authorized credential strength testing | `GPU` `MT` `THERM` | GPU throughput; a legitimate defensive task that looks like mining on a graph |
| **Archival format migration** | mass reformat of an archive | `MT` `CAP` `THERM` | runs for days — throughput times corpus size |
| **Digital preservation** | fixity checking and format normalization | `MT` `CAP` `CRYPTO` | periodic full-corpus verification |
| **Long-horizon data collection** | multi-year continuous capture | `IDLE` `CAP` `WEAR` | the node must outlive the experiment — reliability over speed |
| **Retro emulation and game library host** | emulator plus optional streaming | `ST` `GPU` `VPU-ENC` | single-thread speed for accurate emulation, encoder for streaming |
| **Reproducible research environments** | pinned, rebuildable analysis stacks | `CAP` `MT` | image size and rebuild frequency |
| **Hardware-in-the-loop test rigs** | real hardware under automated test | `RT` `USB` `GPIO` `x86` | deterministic timing and interface availability |

## Common mis-placements

The pairings that keep recurring, stated as the thing to avoid.

| Mis-placement | Why it fails | Do instead |
|---|---|---|
| Consensus stores or relational databases on SD or USB flash | `FSYNC` and `WEAR` — sync-write latency is terrible and the media wears out in months | real NVMe or SATA SSD, always |
| An embedded single-file database over NFS or SMB | file locking over a network share is unreliable and slow | keep the database local; sync the exports, not the database |
| Headless browsers without a memory cap | `MEM` — they leak, then the OOM killer takes an unrelated service with them | hard-cap the container, recycle contexts, isolate on one node |
| Transcoding on a node with no fixed-function encoder | `VPU-ENC` — a CPU-only encode consumes the whole node per stream | choose hardware for the encoder, or pre-transcode and direct-play |
| Sizing model inference by advertised accelerator TOPS | `BW` and `MEM` decide tokens per second, not headline numbers | size by memory bandwidth and capacity |
| Cross-architecture builds under emulation | a 10-40x slowdown turns a five-minute build into hours | keep one native-architecture builder per target |
| Monitoring hosted on the thing it monitors | it goes down exactly when you need it | a separate, deliberately boring node |
| Checksummed storage pools on minimal RAM | caching and integrity work need memory headroom | budget RAM per TB, prefer ECC |
| Realtime workloads co-tenanted with bursty builds | `RT` — jitter from the neighbour causes physical or perceptual defects | dedicate the node, or pin and isolate cores |
| Benchmark runners sharing a node with anything | results become noise, silently, and you trust them anyway | a dedicated node, a fixed governor, no co-tenants |
| Cheap consumer NICs for packet-rate-heavy routing | `PPS` — offload gaps and driver quality dominate | choose the NIC first and the CPU second |
| Replicated cluster storage on a single slow network | replication multiplies both traffic and writes | a fast dedicated storage network, or do not replicate |
| Time-series and log stores without a retention policy | `CAP` and `WEAR` — unbounded growth ends in a full disk at 3am | set retention and downsampling on day one |
| GPU or device passthrough planned without checking IOMMU groups | feasibility is a chipset property, not a spec-sheet one | verify the groups on the exact board before buying |
| Radio-attached services virtualized without a passthrough plan | `USB` — dongle enumeration breaks across reboots and migrations | bare metal, or pin the device by a stable path |
| Anything stateful on a node with no power-loss protection | a hard cut mid-write corrupts the store | a UPS plus clean shutdown, or accept the risk explicitly |
| One node running both the router and everything else | any reboot takes the whole rack offline | keep the network layer on its own node |

---

## Corrections and additions to fold back into `nodes.md`

Reading `nodes.md` Appendix C against PSREF on 2026-07-24 turned up four things worth patching:

1. **Appendix C's Lenovo⇄Intel gen map says the M70q "tops at i7-13700T" at Gen 4.**
   **`ThinkCentre M70q Gen 5` exists** (PSREF last modified 2026-06-09): 13th/14th-gen `T` **plus the
   Intel 300T**, Q670, DDR5 to 64 GB, **two M.2 PCIe 4.0 ×4 with RAID 0/1, no 2.5" bay, no PCIe riser**,
   optional 2.5GbE punch-out, optional 100 W USB-C PD-in. Top SKU **i7-14700T (20C/28T)**.

2. **Appendix C's map ends at M90q Gen 5.**
   **`ThinkCentre M90q Gen 6` exists** (PSREF 2026-04-13 / 2026-06-09): **Core Ultra Series 2 (Arrow
   Lake-S), LGA1851, Q870**, 64 GB DDR5-5600, **3× M.2 PCIe 4.0 ×4**, **one PCIe 4.0 x8 low-profile slot**,
   **optional Intel Arc A310 4 GB dGPU**, **optional Intel I350-T4 quad GbE**, TB4 punch-out, 13-TOPS NPU,
   179 × 182.9 × 36.5 mm.

3. **Appendix C's scope note excludes ThinkStation P-series from the 1 L set.** It shouldn't — the
   **P320 / P330 Gen 1–2 / P340 / P350 / P360 / P3 Tiny / P3 Tiny Gen 2** are all **1 L, 179 × 183 mm**,
   and are the only Tinies that ship a discrete GPU. **P3 Tiny Gen 2** takes **128 GB DDR5** and has a
   **PCIe 5.0 ×4 M.2**. They belong in the roster.

4. **`nodes.md`'s `Mini PC` category carries three Lenovo rows.** On the evidence here it should carry at
   least six more: **M90q (any gen)** as the riser-capable Intel tier, **M75q Gen 5** (STH: *"one of the
   best project tiny mini micro nodes that we've ever reviewed"*, **5–15 W idle**, 128 GB verified),
   **M90q Gen 6**, **P3 Tiny Gen 2**, **neo 50q Gen 5**, and the **M90n-1 Nano** as the sub-1 L option.

Also worth noting against `nodes.md` §11: the ThinkCentre Tiny mount ecosystem is **larger** than the
single Printables link currently cited — see §3.2 above for seven printable designs and four commercial
vendors.

---

## Known gaps — what this document does not have

Stated explicitly rather than papered over, in the spirit of `nodes.md`'s own gaps section.

- **No primary Geekbench reads.** `browser.geekbench.com` was not read this session. Every GB6 figure is an
  aggregator median (cpu-monkey, nanoreview, cputronic, chipversus) or carried forward from `nodes.md`
  Appendix C, which had the same limitation. **Re-verify from an unblocked network before ranking on any
  single value.**
- **Arrow Lake `T` benchmarks are almost entirely `no data`.** Only the **Ultra 7 265T (2954 / 16455)** was
  found. The Ultra 5 225/225T/235/235A/235T/235TA/245/245T, Ultra 7 265 and Ultra 9 285/285T have no
  exact-SKU GB6 figure here. They were **not estimated.**
- **No measured wall power for M90q Gen 6, P3 Tiny Gen 2, neo 50q Gen 4/5 or M70q Gen 5.** The idle
  figures for those rows are `no data`. ServeTheHome's *Project TinyMiniMicro* series is the place to look
  when they publish.
- **Pre-Skylake CPU menus (M72e / M73 / M83 / M93p / M53 / M600) were not verified.** They are described
  by architecture family with `~` markers. Confirm on PSREF before quoting a specific SKU.
- **neo 50q Gen 4 Thin Client exact CPU list not read.** PSREF says "Intel Celeron or 12th-generation
  Core i3"; the specific SKUs are `no data` here.
- **No measured PCIe-riser thermal data.** "A 10GbE NIC runs hot in a 1 L box" is the consistent community
  report, but nobody has published temperature-vs-airflow numbers. Treat the add-a-fan advice as folklore
  that happens to be widely repeated, not as measurement.
- **No measured throughput for the riser slot.** Whether a 10GbE card actually hits line rate on a 35 W
  `T` part with software routing is `no data` in every source read here.
- **Third-party riser (TinyRiser / PowerRiser) reliability is `no data`.** Both are small-run community
  products. Their claim to expose the hidden x4 PCH link is well-documented; long-term field reliability
  is not.
- **Non-Lenovo DC supply behaviour is `no data`.** The slim-tip ID pin is known to exist; exactly which
  BIOS versions warn versus power-limit versus refuse is not documented anywhere read here. **Test one
  before wiring a whole rack to a DIN-rail supply.**
- **Pricing is asks, not sold prices.** eBay's completed/sold archive was not readable this session, so
  every figure in §1.10 is an **asking price** (tagged **V**) or an inferred range (tagged `~`). Sold
  prices typically run **10–25 % below ask** on Best Offer. Treat the table as a relative ranking that
  happens to have dollar signs on it, not as a price list.
- **Most per-model price ranges are inferred, not observed.** Only ten concrete asks were read this
  session — M720q (four configs), M920q, M910q, M70q Gen 5, plus riser and mount pricing. Everything
  else in the §1.10 table is `~`, anchored to those observations and to the July-2026 pass in `nodes.md`.
  **They were not made up, but they were not read either.**
- **Pricing is volatile and directional.** The 2025–26 DRAM shortage is still moving used 1 L prices
  month to month, and it moves them by RAM content rather than by model. Re-verify before buying; a
  figure here that is three months old is probably wrong.
- **No non-US pricing.** All observations are eBay US. European and AU markets behave differently,
  particularly on ex-lease corporate stock.
- **Availability tiers are judgement calls.** "Abundant" versus "plentiful" comes from listing counts
  visible in search results, not from a systematic count. The two ends of the scale (abundant, and
  unobtainable) are reliable; the middle is soft.

---

## Sources

**Lenovo PSREF (primary, per-model spec + processor lists):**
`psref.lenovo.com/syspool/Sys/PDF/ThinkCentre/…` and `…/ThinkStation/…` —
**ThinkCentre M90q Gen 6** (read 2026-07-24, doc dated 2026-04-13 / 2026-06-09) ·
**ThinkCentre M70q Gen 5** (read 2026-07-24, doc dated 2026-06-09) ·
**ThinkStation P3 Tiny Gen 2** (read 2026-07-24, doc dated 2026-07-14/15) ·
ThinkCentre M70q / M70q Gen 2 · ThinkCentre neo 50q Gen 4 and neo 50q Gen 4 Thin Client ·
ThinkCentre M75q Gen 5 · ThinkStation P330 Tiny · ThinkStation P360 Tiny · ThinkStation P3 Ultra SFF Gen 2.
Per-model PSREF PDFs for M700/M900/M710/M910/M720/M920 Tiny, M70q/M80q/M90q Gen 1–5, M625/M715 Tiny,
M75q Tiny/Gen 2, and M90n-1/M75n Nano are the authoritative source for the older CPU menus.

**Measured power, thermals and teardowns — ServeTheHome, *Project TinyMiniMicro*:**
M75q Tiny Gen 5 review (5–15 W idle, 63–66 W sustained, 47–48 dBA, 96 GB and 128 GB DDR5 verified) ·
M75q Tiny Gen 2 review (~12 W idle) · M75q-1 Tiny review (~11 W idle, ~50 W max) ·
M715q Tiny CE review (14–17 W idle) · M90q Tiny Gen 2 review (12–14 W idle, PCIe 3.0 storage) ·
M90q Tiny review (the "only the M90q has a PCIe slot" note) · M80q Tiny Gen 3 review ·
ThinkStation P330 Tiny review · ThinkStation P340 Tiny review (GPU riser, no SATA bay, top vent change) ·
"M90q Gen 3 quietly released with 2.5GbE option".

**Per-CPU specifications:** Intel ARK · AMD product pages · cross-checked against `nodes.md` Appendix C.

**Benchmarks (aggregator medians — see the caveat in Part 2):** cpu-monkey · nanoreview · cputronic ·
chipversus · Notebookcheck (Core 7 240H placement) · PassMark cpubenchmark.net.

**Risers and expansion:** Lenovo FRU **01AJ940** (x16-connector / x8-electrical, Tiny5) and **01AJ929**
(x4) listings · **5C50W00933 / 5C50W00910** (M90q Gen 6 / P3 Tiny Gen 2 x8 riser) ·
**TinyRiser** by FairywrenTech (Tindie) · **PowerRiser** by NandFarm (Lectronz / open-source on GitHub).

**Used-market pricing (asks, read 2026-07-24):** eBay US search and item pages for ThinkCentre M720q,
M920q, M910q, M710q, M900, M75q Gen 2, M90q (all gens), M70q Gen 5, ThinkStation P340 / P360 Tiny ·
riser listings for **01AJ940** and **01AJ929** · **3drackmounts** eBay store (10" ThinkCentre and JetKVM
mounts) · cross-checked against the pricing table in `nodes.md`, verified 2026-07-16.

**Mounts:** Printables **1040412** (Tim, keystone version) · **1215391** + **1215562** (r3vo, unified mount
+ rear support) · **1384009** (owlish) · **1041164** (Emplar) · **877502** (Hermann S) · **1341050**
(Smelliot, M900-specific) · MakerWorld **1141511** (just_actual_kev, enclosed + keystones) ·
MyElectronics.nl 10" rack mounts · 3drackmounts (eBay) · hivets.au.

**Companion document:** `nodes.md` — *Mini-Rack Compute Node Reference*, §6 (power), §9 (out-of-band and
power-loss), §10 (thermals), §11 (mounting), Appendix A (per-device pros/cons), Appendix C (1 L processor
catalog).

---

*Built 2026-07-24. Provenance discipline follows `nodes.md`: **`no data` means "not found this session,"
not "does not exist."** Treat every remaining `no data` as a prompt to look again.*
