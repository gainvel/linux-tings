# Radxa SBC Comparison — Rack / 24-7 Homelab Focus

*Compiled 2026-07-23. Researched against Radxa product briefs and **schematics** (`dl.radxa.com`), Rockchip/Qualcomm SoC datasheets, Jeff Geerling's `sbc-reviews`, Thomas Kaiser's `sbc-bench`, Bret Weber's `sbc.compare`, CNX Software, Liliputing/LinuxGizmos, OpenWrt/Armbian build indexes, and the Radxa + OpenWrt forums.*

***Independently fact-checked 2026-07-23** against primary sources — all five SoC datasheets, the board schematics, every cited GitHub issue, and ARace's full per-variant catalogue. 17 factual errors and 20 sourcing weaknesses were found and corrected; claims that could not be traced to a source were withdrawn rather than softened, and are listed under "Known gaps". Where a correction reverses something this document previously asserted, the old claim is named explicitly so anyone working from an earlier copy can find it.*

**Pricing and availability throughout are from [arace.tech](https://arace.tech), pulled 2026-07-22/23 from per-variant Shopify data**, cross-checked against **[shop.allnetchina.cn](https://shop.allnetchina.cn)** (ALLNET China) on 2026-07-23. Both are **Radxa-only retailers** — neither carries any other brand, so neither is a source for Pi/Orange Pi/x86 comparisons.

### Where the prices come from

| Retailer | Machine-readable? | Use as |
|---|---|---|
| **[arace.tech](https://arace.tech)** | Yes — Shopify, per-variant price + stock exposed | **Primary.** Most Radxa SKUs in stock; cheapest on most lines |
| **[shop.allnetchina.cn](https://shop.allnetchina.cn)** | Yes — Shopify, per-variant price + stock exposed | **Second source — verified working 2026-07-23.** Radxa's long-standing distributor; stocks SKUs ARace has sold out, and vice versa |
| **AliExpress** | **No — blocked** | **Do not cite.** Item URLs bounce `www.aliexpress.com` → `login.aliexpress.com/sync_cookie_read` → `login.aliexpress.us/sync_cookie_write` → back to the item — an endless cookie-sync loop that never serves product data without a browser session. Any AliExpress price here would be a search-snippet rumour, not a read figure |
| ameriDroid, Amazon, Evelta | WARNING: readable, frequently **stale on sold-out SKUs** | Cross-check only |

**Neither ARace nor ALLNET is reliably cheaper — it is per-SKU**, and both post placeholder prices on sold-out variants ($999 at ALLNET, $9,999 at ARace). Check both. Flag key used in the availability tables: [BUY ELSEWHERE] = a cheaper option is **in stock elsewhere, buy there instead** · [RESTOCK WATCH] = a cheaper listing exists elsewhere but is **sold out** (restock target only).

## How to read this document

Every figure is labelled. There are no unlabelled estimates — where nobody has published a number, the cell says **`no data`** rather than a plausible guess.

WARNING: **`no data` means "not found", not "does not exist".** The 2026-07-23 fact-check found eight cells marked `no data` whose values were published — several of them on pages this document already cited for other figures. They have been filled in. Treat a remaining `no data` as a prompt to look again, not as proof the measurement was never made.

| Tag | Meaning |
|---|---|
| **M** | Measured by an independent third party |
| **V** | Vendor claim (datasheet, product brief, marketing) — not a measurement |
| **`no data`** | Nothing published. Do not extrapolate. |
| (AC) | Measured **at the wall** (AC plane — includes adapter loss) |
| (DC) | Measured at the **DC input** (inline meter — excludes adapter loss) |
| [BUY ELSEWHERE] | **A cheaper option is in stock at another Radxa retailer — buy there instead** |
| [RESTOCK WATCH] | A cheaper listing exists at another Radxa retailer but is **sold out** — restock target only |

**The AC/DC distinction matters and the two are not interchangeable.** Geerling and CNX measure at the wall; Bret Weber's `sbc.compare` measures at the DC input with a TC66C. A ~10–20% adapter loss sits between them. For rack/PDU budgeting you want the (AC) figures. No conversion factor has been applied anywhere in this document.

---

## TL;DR for rackers

| If you want… | Pick | Caveat |
|---|---|---|
| Lowest idle watts on a real NIC | (AC) plane: **ROCK 5C** 1.6 W (**$69** 2 GB / **$159** 8 GB) (**CM5** 1.2 W as a module) · (DC) plane: **ROCK 4D** 1.5 W, then **E52C** 1.65 W | The two planes are **not** comparable — see "How to read". E52C and 4D sold out everywhere |
| Best perf-per-watt | **CM5** — 4.86 GFLOPS/W (AC) (module); **Dragon Q6A** 4.79 is the best standalone board (**$169**, 12 GB) | Only six boards have a comparable figure; see §2 |
| Best all-round rack node | **ROCK 5T** (dual 2.5GbE, dual M.2, PoE) — **$179.90** (8 GB) | 12 V DC only, no USB-C power |
| NAS with real SATA | **ROCK 5 ITX** — 4× SATA on a Gen3 ×2 bridge | Sold out; ATX rail stays live in *soft-off* |
| Router / firewall | **E52C** — dedicated PCIe lane per NIC, mainline OpenWrt | Sold out; no M.2 |
| Fastest CPU | **Dragon Q8B** — GB6 1683 ST / 7058 MT | 50 W Linpack transients cooked the PMICs |
| Most expansion | **Orion O6** — PCIe x16 + dual 5GbE | 14.2 W idle, unfixed; x16 **cannot bifurcate** |
| x86 for QuickSync / amd64 images | **Radxa X4** (N100) | 9.1 W idle (AC) — 3–6× an ARM node |
| Dense low-power cluster | **CM5 on DeskPi Super6C** (Radxa-certified, shipping) | No BMC. The denser Xerxes Pi is [X] **unverified** — see §13 |

**Four caveats that shape every decision below:**

1. **PoE is only on SBC-class boards** (ROCK 5 series, ROCK 4D, Dragon Q6A, X4, ZERO 3E, Cubie A5E) via add-on HATs. The Orions, Q8B and the entire E-series have **no PoE at all**.
2. **RK3588S / RK3588S2 / RK3582 have zero PCIe 3.0 lanes.** Any "M.2 NVMe" on ROCK 5A/5C/5C Lite/CM5 is **PCIe 2.1 ×1 (~450 MB/s)**. This is the single most consequential spec-sheet illusion in the lineup — see §4.
3. **There is no BMC, IPMI, or iKVM anywhere in the Radxa lineup.** Nor on any shipping carrier that supports a Radxa module. See §9.
4. **Almost nothing is in stock.** Seven board SKUs plus three accessories are buyable at ARace today. See "Availability & pricing".

---

## Board roster

**ARace column = live per-variant price for SKUs actually buyable on 2026-07-23.** A dash means nothing is in stock — sold-out variants still display prices on the site and several are stale or broken (see [Availability & pricing](#availability--pricing-arace-2026-07-2223)). Pre-orders are listed there too, not here.

| Board | SoC | Class | ARace (in stock) | Why it's in a rack |
|---|---|---|---|---|
| **ROCK 5B** | RK3588 | SBC 100×72 | — | Only ROCK with a true Gen3 ×4 M.2 + dedicated SPI NOR |
| **ROCK 5B+** | RK3588 | SBC 100×72 | **$175.00** (8 GB) | Dual M.2 (×2 each), onboard WiFi 6 |
| **ROCK 5C** | RK3588S2 | SBC 85×56 | **$69.00** (2 GB) · **$159.00** (8 GB) | Lowest measured idle with GbE |
| **ROCK 5C Lite** | RK3582 | SBC 85×56 | — | Same idle, GPU fused off, headless-only |
| **ROCK 5A** | RK3588S | SBC 85×56 | — | Pi-footprint RK3588S; **no M-key, no SATA** |
| **ROCK 5T** | RK3588 | SBC 110×80 | **$179.90** (8 GB Commercial) | Dual 2.5GbE + dual M.2, PoE-capable |
| **ROCK 5 ITX** | RK3588 | Mini-ITX 170² | — | **4× SATA** — the only real Radxa NAS board |
| **ROCK 5 ITX+** | RK3588 | Mini-ITX 170² | — | SATA **removed**, second M.2 added, 32 GB eMMC |
| **ROCK 4D** | RK3576 | SBC ~86×56 | — | Cheapest node with onboard WiFi 6; **no RTC** |
| **Dragon Q6A** | QCS6490 | SBC | **$169.00** (12 GB) · [BUY ELSEWHERE] **8 GB $139.00 at ALLNET** | Best perf/watt; PoE; 3× camera |
| **Dragon Q8B** | SC8280XP | SBC 100×75 | — (pre-order) | Fastest CPU; **standard UEFI**; dual 2.5GbE |
| **Dragon Q5E** | QCS6690 | — | — | **Announced only — never shipped** |
| **Orion O6** | CIX P1 CD8180 | Mini-ITX 170² | — | PCIe x16 + dual 5GbE |
| **Orion O6N** | CIX P1 CD8160 | Nano-ITX 120² | — (pre-order) | O6 compute, lower power, dual 2.5GbE |
| **E20C** | RK3528A | Appliance 66² | — (not carried) | 2× GbE, mainline OpenWrt |
| **E24C** | RK3528A | Appliance | — (not carried) | 4× GbE — **all behind one switch on one RGMII** |
| **E25** | RK3568 (CM3I) | Appliance | — (not carried) | 2× 2.5GbE; only board in OpenWrt 24.10 LTS |
| **E52C** | RK3582 | Appliance 66² | — | Dedicated PCIe lane per NIC + **built-in serial console** |
| **E54C** | RK3582 | Appliance | — | 4× GbE behind a switch; **no mainline OpenWrt** |
| **X4** | Intel N100 | SBC 85×56 | — | x86: QuickSync, amd64 images, UEFI/PXE |
| **X5** | Intel N150 | SBC 85×56 | — | Documented, **no price or ship date** |
| **X2L** | Celeron J4125 | SBC 155×80 | — | Superseded; GbE only, no PoE |
| **ZERO 3E** | RK3566 | SBC 65×30 | **$39.99** (2 GB) | GbE + PoE in a Pi-Zero footprint; **SD-only boot** |
| **ZERO 3W** | RK3566 | SBC 65×30 | **$39.99** (2 GB + 16 GB eMMC, header soldered) | eMMC + WiFi, **no Ethernet** |
| **Cubie A5E** | Allwinner A527 | SBC ~69×56 | — | Dual GbE + NVMe at ~2 W |
| **CM5 / CM5 Lite** | RK3588S2 / RK3582 | Module 55×40 | — | 1.2 W idle (AC) (Lite: 2.4 W (DC)); cluster-carrier density |
| **CM4 (Radxa)** | RK3576 | Module 55×40 | — | RPi-CM4-shaped, WiFi 6 onboard, supported to 2035 |
| **CM3J** | RK3568J | Module 55×40 | — | **Only Radxa module with a true 2×100-pin CM4 connector** |
| **CM3S** | RK3566 | SODIMM | — | RPi CM3-class; currently the most available module |
| **NX5** | RK3588S | SODIMM 260-pin | — | Jetson-shaped; **not Jetson-compatible in software** |

Excluded as non-rack: ZERO 2 Pro (no Ethernet, 30 W class), Cubie A7Z (no Ethernet), ROCK Pi S (100 Mbps, 512 MB ceiling), NIO 12L, SiRider S1.

---

## 1. Compute

| Board | SoC | Process | CPU config | GPU | NPU (V) |
|---|---|---|---|---|---|
| ROCK 5B / 5B+ / 5T / 5 ITX / ITX+ | RK3588 | 8 nm | 4×A76 @2.2–2.4 + 4×A55 @1.8 | Mali-G610 MC4 | 6 TOPS |
| ROCK 5A | RK3588S | 8 nm | 4×A76 + 4×A55 | Mali-G610 MC4 | 6 TOPS |
| ROCK 5C / CM5 | RK3588S2 | 8 nm | 4×A76 + 4×A55 | Mali-G610 MC4 | 6 TOPS |
| ROCK 5C Lite / CM5 Lite / E52C / E54C | RK3582 | 8 nm | **2×A76 + 4×A55** | **none — fused off** | 5 TOPS |
| ROCK 4D / Radxa CM4 | RK3576 | **disputed** ¹ | 4×**A72** @2.2 + 4×A53 @2.0 | Mali-G52 MC3 | 6 TOPS |
| E20C / E24C | RK3528A | — | 4×A53 @2.0 | Mali-450 | none |
| E25 | RK3568 | — | 4×A55 @2.0 | Mali-G52 | 0.8–1 TOPS ² |
| Dragon Q6A | QCS6490 | 6 nm | 1×A78 @2.7 + 3×A78 @2.4 + 4×A55 @1.96 | Adreno 643L | 12 TOPS |
| Dragon Q8B | SC8280XP | 5 nm | **4×Cortex-X1C @3.0** + 4×A78C @2.42 | Adreno 690 | 15 or 29+ ³ |
| Orion O6 / O6N | CIX P1 | 6 nm | 4×A720 @2.6 + 4×A720 @2.4 + 4×A520 @1.8 | Immortalis-G720 MC10 | 30 (45 w/ GPU+CPU) |
| X4 | Intel N100 | Intel 7 | 4×E-core @≤3.4 GHz, 6 W PL1 | UHD 24 EU | **none** ⁴ |
| ZERO 3E / 3W | RK3566 | — | 4×A55 @1.6 | Mali-G52-2EE | 0.8 TOPS |
| Cubie A5E | A527 / T527 | — | 4×A55 @1.8 + 4×A55 @1.4 | Mali-G57 MC1 | 2 TOPS (T527 only) |

¹ **Genuinely unresolved.** CNX indexes RK3576 at 22 nm; ArmSoM says "second-generation 8 nm"; neither Rockchip datasheet states a node.
² Radxa's wiki says 0.8 TOPS, radxa.com says 1 TOPS.
³ `sbc.compare` records 15 TOPS; Radxa marketing says 29+. Both are whole-system AI figures, not NPU-only. The Orion's 30 TOPS *is* NPU-only.
⁴ The N100's "Intel GNA 3.0" is a low-power audio/wake-word DSP, **not** an NPU. Do not treat it as an AI accelerator.

**On the RK3582:** it is an RK3588S die with two A76 cores and the GPU fused off — Radxa themselves call it a binned "lottery" part. For a headless server this is a *good* trade: it keeps **100% of the single-thread performance and 100% of the memory bandwidth**, loses **~24–26%** multi-thread (GB6 2269 vs 2962 = −23%, 2178 vs 2928 = −26%; 7-Zip MT 11639 vs 15681 = −26%), and drops a GPU you were never going to use.

---

## 2. Benchmarks

All figures **M** unless noted. 7-Zip is `sbc-bench` MIPS; memory is `tinymembench` memcpy/memset MB/s. Testers: **JG** = Geerling, **TK** = Kaiser, **BW** = Weber (sbc.compare).

| Board | GB6 ST | GB6 MT | 7-Zip ST | 7-Zip MT | memcpy / memset | Tester |
|---|---|---|---|---|---|---|
| **Dragon Q8B** | **1683** | **7058** | 4467 | **36243** | **18181 / 42179** | BW |
| **Orion O6N** | 1327 | **6954** | 3867 | 33823 | 16366 / 41322 | BW |
| **Orion O6** | 1085–1314 ¹ | 5671–6273 ¹ | 3945 | 32690 | 16880 / **48030** | JG/TK/CNX |
| **X4 (N100)** | 1185–1243 | 2658–2981 | 3459 | **7950–9597** ² | 8139 / 7968 | JG/BW |
| **Dragon Q6A** | **1176–1180** | 3103–3215 | 3829 | 17401 | 8245 / 19482 | JG/BW/TK |
| **NX5** (RK3588S) | 851 | 3090 | `no data` | 16624 | 11892 / 27801 | BW |
| **ROCK 5B** | 748–849 | 2941–3043 | 3113 | 16486 | 10830 / 29220 | JG/BW/TK |
| **ROCK 5C** | 761–834 | 2928–2962 | 3026 | 15681 | 12280 / 29750 | JG/BW/TK |
| **ROCK 5A** | 823 | 2961 | 3029 | 15849 | 9170 / 27080 | BW/TK |
| **ROCK 5T** | 761 | 2995 | 2546 | 15558 | `no data` | BW |
| **ROCK 5 ITX** | 767 | 2969 | 2622 | 15568–15780 | 12540 (A76) ³ | BW/CNX |
| **ROCK 5B+** | 813 | 2944 | 2883 | 15470 | `no data` | BW |
| **CM5** | 768 | 2990 | `no data` | `no data` | 12605 / 29860 | JG |
| **CM5 Lite** (RK3582) | 847–875 | 2178–2269 | 3187 | 11639 | 11944 / 28254 | BW |
| **ROCK 5C Lite** (RK3582) | `no data` | `no data` | 3094 | 11160 | 12410 / 29620 | TK |
| **E52C** (RK3582) | 822 | 2298 | `no data` | `no data` | `no data` | GB browser |
| **ROCK 4D** | 319–325 | 1332–1355 | 1771 | 10920 | 5210 / 15450 | BW/TK |
| **Cubie A5E** | 114–241 ⁴ | 590–1005 ⁴ | 1517 | 8830 | 2710 / 5570 | TK/GB |
| **ZERO 3E** | `no data` | `no data` | `no data` | 4465 | `no data` | BW |
| **ZERO 3W** | `no data` | `no data` | 1155 | 4000 | 2400 / 5580 | TK |
| **E25 / E20C / E24C / E54C** | `no data` | `no data` | `no data` | `no data` | `no data` | — |

**The ROCK 5C Lite and CM5 Lite are separate datasets and are not interchangeable** despite sharing the RK3582 — they were measured by different testers and differ by ~4% on 7-Zip MT. Earlier revisions of this document merged them; that was wrong.

**No benchmark of any kind exists for the Dragon Q5E, X2L, CM3J, CM3S, CM3 or CM3I.** They are absent from every tester's dataset — that is a gap in the world, not an omission here.

¹ **Firmware-dependent.** Geerling re-ran after the April 2025 firmware: GB6 went 1314/6273 (fw 0.2.x) → 1228/6009 (fw 9.0.0). Windows 11 scores lower still (1085/5671). Report as a range.
² **Cooling/PL-dependent, not silicon.** Same N100 scores 13573 on the X4L and 14090 on a generic mini-PC — a **1.8× spread**. Single-thread barely moves, confirming a sustained-power cap.
³ A76 cluster. A55 cluster measures 6548 MB/s. `sbc.compare`'s ITX memory figures (4300/6314) are ~3× below every other RK3588 measurement including the same tester's sibling boards — treat as a bad run.
⁴ **Unresolved 2.1× spread.** Two Armbian Geekbench uploads give 114–115/590–602; a hands-on review gives 241/1005. Likely a cpufreq governor or DVFS cap on the Armbian images.

### Real-workload results (Geerling's identical PTS subset — the only apples-to-apples set)

**Linux kernel compile** (defconfig, seconds, lower better): Orion O6 **776** · X4 **1019** · ROCK 5B **1154** · CM5 **1428** · ROCK 5C **1788** · Dragon Q6A **3554** (!)

WARNING: The Q6A figure is **storage-bound, not a CPU result** — that run had only a microSD (4K random write 0.80 MB/s) and a defconfig build is small-write dominated. Given its 17400 7-Zip MT, its true standing on NVMe should be near or above the ROCK 5B, but **nobody has measured that**.

**phpbench** (higher better): X4 **706,245** · Orion O6 611,861 · Q6A 543,386 · ROCK 5C 413,840 · CM5 373,577 · ROCK 5B 372,255

**This inverts the multi-core ranking.** The 4-core N100 beats 8-core RK3588 by ~1.9×. PHP — and most web/database serving — is latency- and branch-bound, not thread-bound. **If your rack runs web services, core count is the wrong metric.**

**x264 1080p** (fps): Orion O6 **50.8** · Q6A 22.9 · ROCK 5B 22.3 · CM5 22.3 · ROCK 5C 20.9 · X4 20.2

### Perf-per-watt

WARNING: **Two incompatible methodologies exist and must not be ranked against each other.** Geerling divides HPL GFLOPS by watts **at the wall during the run** ((AC)); Weber's `sbc.compare` divides Linpack GFLOPS by **peak** Linpack watts at the **DC input** ((DC)). On the one board measured both ways the answers differ by ~2×: **ROCK 5B is 4.28 by JG's method and 2.17 by BW's.** A single merged ranking is meaningless, so there are two lists.

**JG — HPL GFLOPS/W at the wall (AC):**
**CM5 4.86** ≈ **Q6A 4.79** > ROCK 5B 4.28 > ROCK 5C 4.11 > Orion O6 3.49 > **X4 2.33**

**BW — Linpack GFLOPS / peak Linpack W at the DC input (DC):**
ROCK 4D 3.27 > Q8B 3.10 > ROCK 5B 2.17

The ROCK 4D and Q8B have **no JG figure at all** — Geerling has published HPL for only six Radxa boards (5B, 5C, CM5, X4, O6, Q6A). Earlier revisions of this document ranked the 4D and Q8B inside the JG list; that inverted the true order, since under BW's method both sit *above* the ROCK 5B rather than below it.

### NPU — measured inference, not TOPS

Vendor TOPS ratings are peak-INT8 marketing and are not comparable across vendors. What has actually been measured:

- **RK3588 / RK3588S2** (Rockchip's own `rknn_model_zoo`, single NPU core of 3, INT8): yolov5s **48.4 fps**, resnet50-v2-7 **110.1 fps**, mobilenetv2 450.7 fps, yolov8n 73.5 fps. Independent cross-check: ResNet18 INT8 at **244 fps across 3 cores**.
- **RK3576** (same methodology): yolov5s **57.5 fps**, resnet50 99.0 fps — *per core the RK3576 NPU beats the RK3588's* on most detection models.
- **QCS6490** (independent, QAIRT runtime): YOLOv8 **12.03 ms**, EfficientNet-b0 11.94 ms, ViT 159.76 ms. Precision not stated.
- **CIX P1** (independent, Frigate integration): yolox_m **~40 ms/inference, ~25 fps** on a 1080p stream with the CPU idle.
- **RK3582 (5 TOPS), Allwinner T527, RK3528A, and the Q8B's Hexagon: `no data`.** No per-model figures exist anywhere. The Q8B's "15/29 TOPS" is pure marketing.

### Deliberately excluded

- Radxa's "1.3×/1.15×/2×/2× vs RK3588" Q6A claim — vendor, no methodology.
- The "326.1 / 1693.7" Q6A figure circulating as Geekbench — it is **CPU-Z under Windows 11**.
- A "CM5 with RK3576" entry in one 2025 roundup — that's **ArmSoM's** CM5. Radxa's is RK3588S2.
- cpubenchmark.net/PassMark aggregates — crowd-sourced, unknown board/cooling/OS per sample.
- **sysbench**: essentially absent from this ecosystem. The only figure found anywhere is the Cubie A5E (~2500 events/s). Geerling, Kaiser and Weber have all standardized on 7-Zip/Geekbench/HPL. Don't expect this gap to close.
- **STREAM**: `no data` for every board. All memory figures are tinymembench.

---

## 3. Memory & storage

| Board | RAM | Max | eMMC | SPI NOR | M.2 | SD |
|---|---|---|---|---|---|---|
| ROCK 5B | LPDDR4x | 16 GB | socket | **16 MB dedicated** | 1× M-key **Gen3 ×4** + E-key | yes |
| ROCK 5B+ | LPDDR5 | 32 GB | onboard opt. | **16 MB XT25F128BW** | **2× M-key Gen3 ×2** | yes |
| ROCK 5C / Lite | LPDDR4x | 32 GB | **socket shared with SPI** | shared module | FPC **Gen2 ×1** only | yes |
| ROCK 5A | LPDDR4x | 32 GB | **socket shared with SPI** | shared module | **E-key only** | yes |
| ROCK 5T | LPDDR5 | 32 GB | onboard 16–256 GB | `no data` | **2× M-key Gen3 ×2** | yes |
| ROCK 5 ITX | LPDDR5 5500 | 32 GB | **8 GB soldered** | **16 MB XT25F128BW** | 1× M-key Gen3 ×2 + E-key | yes |
| ROCK 5 ITX+ | LPDDR5 5500 | 32 GB | **32 GB soldered** | claimed 32 MB (unverified) | **2× M-key Gen3 ×2** | yes |
| ROCK 4D | LPDDR5 4800 (**32-bit**) | 16 GB | eMMC **or** UFS module | **16 MB MX25U12832F** | FPC Gen2 ×1 | yes |
| Dragon Q6A | LPDDR5 5500 | 16 GB | eMMC/UFS module | yes | 1× M-key **Gen3 ×2** | yes |
| Dragon Q8B | LPDDR4x-4266 | 32 GB | **none** | UEFI | **Gen3 ×4 + Gen3 ×2** + UFS 3.1 | **no** |
| Orion O6 | LPDDR5 5500 (128-bit) | 64 GB | — | **8 MB, socketed SOP8** | **Gen4 ×4** + E-key + **x16 slot** | yes |
| Orion O6N | LPDDR5 5500 (128-bit) | 64 GB | — | 8 MB, **soldered** | 2× M-key Gen4 + E-key + B-key + UFS | yes |
| E52C | LPDDR4 | 8 GB | 16–64 GB | present | **none** — lanes spent on NICs | yes |
| E54C | LPDDR4 | 32 GB | opt. (retail SKUs: none) | **16 MB** | 1× M-key Gen2 ×1 | yes |
| X4 | LPDDR5 (**32-bit**) | 16 GB | soldered opt. | UEFI | 1× M-key **2230 only**, Gen3 ×4 | **no** |
| ZERO 3E | LPDDR4 | 8 GB | **none** | **none** | none | **yes — only boot media** |
| Cubie A5E | LPDDR4x | 4 GB | 8–32 GB | **16 MB** | M-key 2230 Gen2 ×1 | yes |
| CM5 | LPDDR4X | 32 GB | up to 512 GB | carrier-dependent | 2× Gen2 ×1 (muxed w/ SATA/USB3) | yes |

**Storage takeaways:**

- **Only the ROCK 5B gets a Gen3 ×4 M.2.** Every other RK3588 board splits into ×2. This is hard-wired: on the 5B+, connector lanes 2–3 are **physically unconnected** — no firmware setting recovers ×4.
- **ROCK 5A / 5C / 5C Lite / ROCK 4D share one connector between eMMC and SPI flash.** You get one or the other. Since NVMe-direct boot needs a bootloader in SPI, choosing eMMC on these boards **forecloses NVMe-direct boot**. Both the 5B *and* the 5B+ carry a separate dedicated 16 MB XT25F128BW SPI NOR — that is precisely the 5B/5B+-vs-5C difference that matters for a headless node.
- **ZERO 3E is microSD-only with no fallback** — no eMMC, no SPI. For a 24-7 node that's the classic SBC failure mode with no mitigation.
- **X4's M.2 is 2230 only**, and the NVMe idles above 70 °C in that slot.
- **The Orion's Gen4 M.2 is the fastest interface here on paper**, but the only published measurement negotiated **Gen3 ×4 (8 GT/s)** with a Gen3-class SSD. Whether the slot delivers Gen4 in practice is **unconfirmed**.
- **128-bit LPDDR5 on the Orions** (~48 GB/s measured memset) is roughly 1.7× the RK3588's and 6× the X4's. The **X4's 32-bit bus is a measured handicap** — the same N100 on a 64-bit bus was >2× faster on a kernel compile.

### Measured storage throughput

| Board | Interface | Read | Write |
|---|---|---|---|
| ROCK 5B | NVMe Gen3 ×4 | **3091 MB/s** (JG, fio 1M seq) · 1037 (JG, iozone) · 1401–1468 (BW) ⁑ | 1311 (JG, iozone) · 594–650 (BW) ⁑ |
| ROCK 5 ITX | NVMe Gen3 ×2 | 1.45 GB/s | 1.46 GB/s |
| ROCK 5 ITX | 1× SATA ext4 | 138 MB/s | 142 MB/s |
| ROCK 5 ITX | **4× SATA SSD RAID-0 seq** | — | **~400 MB/s** (!) |
| Dragon Q8B | M.2 (fio) | 1136 MB/s | 1110 MB/s |
| ROCK 5A | M.2 via E-key Gen2 ×1 | 289 MB/s | — |
| ROCK 4D | NVMe via FPC Gen2 ×1 | ~301 MB/s | — |
| Cubie A5E | NVMe Gen2 ×1 | ~366 MB/s | — |
| X4 | NVMe Gen3 ×4 (2230) | 1701 MB/s | 1527 MB/s |
| CM5 | eMMC | 246 MB/s | 210 MB/s |

WARNING: **The ROCK 5 ITX RAID-0 number is the one that should set NAS expectations** — ~400 MB/s across four SSDs, far under the ASM1164's ~1.97 GB/s uplink.

⁑ **Read NVMe figures as method-dependent, not as a board ceiling.** On the *same* ROCK 5B with the same Gen3 ×4 KIOXIA XG6, Geerling's fio 1M sequential read gives **3.09 GB/s — about 78% of the theoretical link** — while his iozone read on the same board gives 1.04 GB/s. A 3× spread between benchmarks on one board means no single quoted number characterises the interface. An earlier revision of this document generalised the low figure into a claim that "the lane map is an upper bound the memory/IOMMU path never reaches"; **Geerling's own fio result refutes that** and it has been removed.

---

## 4. PCIe lane topology

Sourced from Rockchip/Qualcomm datasheets and **net-level tracing of Radxa schematics**. This section exists because "2× M.2" on a spec sheet tells you nothing about what those slots are actually wired to.

### SoC lane budgets

| SoC | PCIe 3.0/4.0 | Combo PHYs (PCIe 2.1 **or** SATA **or** USB3) | Ethernet MACs |
|---|---|---|---|
| **RK3588** | 4× Gen3 (bifurcable ×4 / 2×2 / 4×1) | **3** (PHY2 also does USB3) | 2× RGMII |
| **RK3588S / S2 / RK3582** | **ZERO** | **2** (PHY0, PHY2) | 1× RGMII |
| **RK3576** | **ZERO** | 2 | 2× RGMII |
| **QCS6490** | 3× Gen3 (1×1 + 1×2) | — | **none** |
| **SC8280XP** | 4-lane or 2×2-lane + PCIe4 ctrl | — | none |
| **CIX P1** | `no data` (≥17 lanes visible at board level) | — | none |
| **Intel N100** | 9× Gen3 | — | none |

**Proof of the RK3588 combo-PHY mux, at the ball level:** pin **H30** is `PCIE20_2_TXP / SATA30_2_TXP / USB30_SSTXP` and **H29** the matching `…TXN` leg (RK3588 datasheet V2.2, p.24) — one physical pin, three protocols, one at a time. Radxa's own comparison wiki tabulates this as "3×1 Lane PCIE2.0 / 3× SATA 3.0", which invites exactly the wrong reading.

**RK3588S/S2/RK3582 have no PCIe 3.0 block on the die at all** — confirmed by grepping all three official datasheets for `PCIE30` (no sections, no pins). Max 2 lanes, both Gen2 ×1.

### Per-board lane map

| Board | M.2 #1 | M.2 #2 | Slot | SATA | NIC attachment |
|---|---|---|---|---|---|
| **ROCK 5B** | M-key **Gen3 ×4** | E-key Gen2 ×1 (PHY1) | — | **none** | RTL8125B, Gen2 ×1 (PHY0) |
| **ROCK 5B+** | M-key Gen3 ×2 (Port0) | M-key Gen3 ×2 (Port1) | — | none | RTL8125B Gen2 ×1; WiFi PHY1; B-key = **USB2 only** |
| **ROCK 5T** | M-key Gen3 ×2 | M-key Gen3 ×2 | — | none | 2× 2.5GbE, PHY mapping **unverified** |
| **ROCK 5 ITX** | M-key Gen3 ×2 (Port0) | E-key Gen2 ×1 (PHY1) | — | **4× via ASM1164 on Gen3 ×2** (Port1) | 2× RTL8125BG, Gen2 ×1 each (PHY0+PHY2) |
| **ROCK 5A** | E-key Gen2 ×1 / SATA / USB2 | — | — | via E-key only | **RTL8211F RGMII — no PCIe lane** |
| **ROCK 5C / Lite** | FPC Gen2 ×1 | — | — | none | GbE; attachment `no data` |
| **ROCK 4D** | FPC Gen2 ×1 | — | — | none | **RTL8211F RGMII — no PCIe lane** |
| **CM5** | 2× Gen2 ×1 (one shared w/ USB3+SATA, one w/ SATA) | — | — | on carrier, **mutually exclusive with PCIe** | GMAC RGMII |
| **Dragon Q6A** | M-key **Gen3 ×2** (PCIe1) | — | — | none (no SATA in SoC) | RTL8111K on **Gen3 ×1** |
| **Dragon Q8B** | M-key Gen3 ×4 | M-key Gen3 ×2 + E-key | — | none | **2× 2.5GbE behind one QPS615 switch port** |
| **Orion O6** | M-key **Gen4 ×4** | E-key Gen4 ×2 | **x16 phys / Gen4 ×8 elec, NO bifurcation** | none | 2× RTL8126, **Gen3 ×1 each** (measured) |
| **E52C** | **none** | — | — | none | **2× RTL8125B, dedicated PCIe 2.0 ×1 each** |
| **E54C / E24C** | M-key Gen2 ×1 (**both boards**) | — | — | none | **RTL8367RB switch on ONE RGMII** |
| **E20C** | none | — | — | none | 1× RGMII+RTL8211F, 1× **PCIe 2.0 ×1**+RTL8111H |
| **X4** | M-key Gen3 ×4 (2230) | — | — | none | **Intel I226-V** on PCIe (mainline `igc`) |

### Shared-lane gotchas

1. **On ROCK 5B all three combo PHYs are already spent** (NIC + WiFi + USB3). You cannot add SATA without sacrificing one. The known "switch the E-key to SATA via device-tree overlay" trick works because it re-models PHY1 — and it costs you the WiFi slot.
2. **ROCK 5 ITX trades a USB3 controller for a NIC.** PHY2 goes to the second RTL8125BG, so all four USB3 ports hang off a single **GL3523 hub on one 5 Gbps upstream**. The ITX's aggregate USB3 bandwidth is *lower* than a ROCK 5B's despite having twice the ports.
3. **`USB3OTG_2` has no USB 2.0 companion** on RK3588/S/S2/3582. A port fed by PHY2 must borrow its USB2 pair from another controller, or USB2 devices won't enumerate.
4. **RK3576's combo port 1 couples USB2 to USB3** — repurposing that lane for PCIe/SATA kills a USB2 port too.
5. **Rockchip GMACs cost zero PCIe lanes** (RGMII + external PHY), but **2.5GbE always costs a lane** because RGMII tops out at 1 Gb/s. That's why every 2.5GbE Rockchip board burns combo PHYs.
6. **PCIe 2.1 ×1 is 8b/10b — 500 MB/s raw, ~450 usable.** For a 2.5GbE NIC (312 MB/s) that's only ~1.4× headroom. Contrast the O6's RTL8126 on Gen3 ×1 (985 MB/s for 625 MB/s), which is correctly provisioned — a Gen2 lane would *not* have sufficed for 5GbE.
7. **Orion O6's x16 slot explicitly does not bifurcate.** No ×4×4×4×4 quad-NVMe carriers.
8. **Q8B's dual 2.5GbE are two MAC functions behind the QPS615's third downstream port** — they share one upstream link and are not bandwidth-independent the way the ROCK 5 ITX's two separate RTL8125Bs are. The QPS615 is also a real AVB/TSN bridge (802.1AS-rev/Qav/Qbv/Qbu).
9. **E52C consumes both RK3582 PCIe lanes on its two NICs** — which is exactly why it has no M.2, and exactly why its NICs don't contend.
10. **AICore M.2 NPU accelerators want Gen3 ×4** (DX-M1) but every Radxa compute module offers **Gen2 ×1** — an ~8× link downgrade. Only the CM3I (Gen3 ×2) is better, and it is an RK3568 — four A55s against the CM5's four A76 + four A55.

---

## 5. Networking

| Board | Ethernet | Controller | Attachment | Measured |
|---|---|---|---|---|
| ROCK 5B / 5B+ | 1× 2.5GbE | RTL8125B | PCIe 2.1 ×1 | 2.35 Gbps |
| ROCK 5C | 1× GbE | — | — | 937 (JG) / 942 TX, 941 RX (BW) |
| ROCK 5A / 4D | 1× GbE | RTL8211F | **RGMII (SoC MAC)** | 941 / 936 Mbps |
| ROCK 5T | **2× 2.5GbE** | — | `no data` | **2343 TX / 2338 RX** (BW) |
| ROCK 5 ITX / ITX+ | **2× 2.5GbE** | 2× RTL8125BG | PCIe 2.1 ×1 each | **2.35 TX / 1.3–2.32 RX** (!) |
| CM5 | 1× GbE | — | GMAC RGMII | 940 Mbps (JG) |
| Dragon Q6A | 1× GbE | RTL8111K | PCIe Gen3 ×1 | 943 (JG) / 941 TX, 941 RX (BW) |
| Dragon Q8B | **2× 2.5GbE** | QPS615 TSN bridge | shared switch port | **2342 TX / 2339 RX** (BW) ⁑ |
| Orion O6 | **2× 5GbE** | 2× RTL8126 | Gen3 ×1 each | `no data` above 1 Gbps ³ |
| Orion O6N | **2× 2.5GbE** | — | — | `no data` above 1 Gbps ³ |
| E20C | 2× GbE | RTL8211F + RTL8111H | 1× RGMII, 1× PCIe | 620 Mbps WireGuard ¹ |
| E24C / E54C | **4× GbE** | RTL8367RB switch | **ONE RGMII** | `no data` |
| E25 | 2× 2.5GbE | 2× RTL8125B | PCIe | 2.2–2.3 Gbps (vendor) |
| E52C | **2× 2.5GbE** | 2× RTL8125B | **dedicated PCIe 2.0 ×1 each** | **933 Mbps w/ SQM** ² |
| X4 | 1× 2.5GbE | **Intel I226-V** | PCIe | 2.35 Gbps |
| ZERO 3E | 1× GbE | RTL8211F | RGMII | **942 TX / 941 RX** (BW) |
| Cubie A5E | 2× GbE | 2× Maxio MAE0621A | RGMII (inferred) | `no data` |

¹ CPU/kernel-stack benchmark, not a NIC-path measurement.
² **ISP-bound, not board-bound.** Every E52C tester has a 1 Gbps WAN. No >1 Gbps routed NAT figure exists for it.
³ Both Orions have a published 941/942 Mbps iperf3 result, but it is clearly **test-LAN-limited** — a 1 GbE path measured on a 5GbE/2.5GbE NIC tells you nothing about the interface. Treated as `no data` for the link speeds that matter.
⁑ Aggregate, not per-port independence — both MACs sit behind one QPS615 downstream port (see §4 gotcha 8).

WARNING: **ROCK 5 ITX RX asymmetry is unresolved**: CNX measured 2.32 Gbps RX, Kaiser measured 1.3–1.6 Gbps. Likely driver/kernel dependent.

**For rack use:** Orion O6's **dual 5GbE on correctly-sized Gen3 lanes** is the standout, plus an x16 slot for a real 10GbE NIC. **E52C is the best-architected router node** — a dedicated lane per NIC, no switch in the path, mainline OpenWrt. **The X4's Intel I226-V is the only NIC here using an in-tree mainline driver (`igc`)** rather than an out-of-tree Realtek module — one less thing to break on a rolling kernel.

**Avoid E24C/E54C for routing.** Four ports behind an RTL8367RB on a single RGMII uplink means **~1 Gbps aggregate for anything the CPU routes, NATs or firewalls**. Radxa's own hardware-design page for the E54C claims *"each gigabit port has an independent data channel, avoiding network congestion"* — **their own schematic contradicts this.** Treat that page as marketing.

---

## 6. Power & efficiency

The number that matters most for 24-7. Read the (AC)/(DC) plane markers — they are not interchangeable.

| Board | Idle | Load | Peak | Input | Recommended PSU (V) |
|---|---|---|---|---|---|
| **CM5** (on IO board) | **1.2** (AC) | 8.4 (AC) | 10 (HPL) | 5 V module / 12 V carrier | `no data` |
| **ZERO 3W** | **0.9** (DC) | 2.6 (DC) | 3.9 | USB-C 5 V | 5 V/2 A |
| **ZERO 3E** | **`no data`** | 3.2 (DC) | — | USB-C 5 V, PoE HAT | 5 V/2 A |
| **ROCK 4D** | **1.5** (DC) / 2.7 (AC) ⁷ | 4.8 (DC) / 7 (AC) | 7.0 | USB-C **5 V only** | 5 V/3 A |
| **ROCK 5C** | **1.6** (AC) (2.0 w/ HDMI) / **2.4** (DC) | 9.5 (AC) / 7.4 (DC) | 12.4 (AC) (HPL) / 16.4 (DC) (Linpack) | USB-C 5 V/3 A | 10 W bare / 25 W loaded |
| **ROCK 5C Lite** | 1.63–1.72 (AC) ⁷ | `no data` | — | as 5C | as 5C |
| **E52C** | **1.65 / 1.91 / 2.11** (DC) ¹ | 7.35 (DC) | — | USB-C 5 V/3 A, **no PoE** | 5 V/3 A |
| **Cubie A5E** | ~2 (plane unstated) | ~6 | — | USB-C 5 V | 10 W / 15 W loaded |
| **ROCK 5A** | **2.9** (DC) | 6.5 (DC) | 15.7 (Linpack) | USB-C PD 9–20 V | 24 W / 30 W w/ SSD |
| **Dragon Q6A** | **2.9** (AC) / 2.5 (DC) | 9.3 (AC) / 6.4 (DC) | 10.1 (AC) (HPL) | USB-C PD **12 V** | 30 W |
| **ROCK 5B** | **1.5 / 1.7 ⁷ / 3.6 (AC) / 4.5 (DC)** ² | 9.5 (AC) / 17 (AC) (stress-ng) / 8.6 (DC) | 19.2 | USB-C PD 9–20 V | >30 W |
| **ROCK 5B+** | **`no data`** ³ | 8.6 (DC) | 19.2 | USB-C PD | >40 W |
| **ROCK 5T** | **3.9 (AC) / 4.4 (DC)** | 9.2 (DC) / 11 (AC) | 19.0 | **12 V barrel only** | ≥36 W |
| **ROCK 5 ITX** (bare) | **4.8–5.4** (DC) | 11.5 (DC) | 20.2 | 12 V barrel / ATX / PoE | 36 W (90 W w/ 4 HDD) |
| **ROCK 5 ITX** (4 SSD+10GbE) | **12.20** (DC) | `no data` | — | as above | as above |
| **ROCK 5 ITX** (4-drive NAS) | **27.6–27.9** (AC) | `no data` | `no data` ⁴ | as above | as above |
| **ROCK 5 ITX+** | **`no data`** | **`no data`** | — | as ITX + Molex out | `no data` |
| **Dragon Q8B** | **5.9** (DC) | 25.1 (DC) | **50.3 (Linpack)** ⁵ | USB-C PD, **12 V/20 V switch** | 65 W @ 20 V |
| **X4** (N100) | **9.1** (AC) | 18.5 (AC) / 12.2 (DC) | 21.2 (AC) (GB6-MT) / 22.1 (DC) (Linpack) | USB-C PD **12 V/2.5 A** | 18 W / 25 W loaded |
| **Orion O6N** | **9.3** (DC) (vendor: ~15) | 23.7 (DC) | 39.0 | **12 V barrel / ATX floppy** | ≥60 W |
| **Orion O6** | **14.2** (AC) | 24.9 (AC) | **35.7 (HPL)** | ATX 24-pin / USB-C PD 20 V | ≥65 W |
| E20C / E24C / E25 / E54C | **`no data`** ⁶ | `no data` | — | 5 V USB-C or 12 V barrel | — |

¹ No cable / one 2.5G cable / two cables — the three figures are exactly what the tester reported. Setting the DMC governor to `performance` halves A76 DRAM latency (**230→130 ns**, measured on the E52C); the **power cost on the E52C is `no data`.** The "~half a watt" adder often quoted alongside this was measured by the same author on a **ROCK 5 ITX**, in a different thread, and should not be carried over.
² **Four sources, four answers** (1.5 / 1.7 / 3.6 (AC) / 4.5 (DC)), two from the same author on different rigs. Budget the high figure.
³ **The single most consequential blank in this table.** No published idle measurement exists for the 5B+ anywhere.
⁴ **The ATX PSU keeps running when the board is "off"** — the board cannot fully de-energise its own rail. Independently reproduced by a second user on a different PicoATX supply, so the *behaviour* is solid. **The wattage is not: no source measures the soft-off draw**, and the "1.5 W standby" figure sometimes attributed to Radxa appears in neither forum thread nor the product brief. Earlier revisions of this document quoted "8.0 W in soft-off" — that number is unsourceable and has been withdrawn. Meter it yourself if it matters.
⁵ Transient the board **could not sustain** — the reviewer reported the PMICs "cooked and gave up after 10s each time" even with industrial cooling, and excluded Linpack from his power results.
⁶ E24C vendor claim only: 5 W typical / 12 W max.
⁷ **Unverified.** Could not be traced to a named source during fact-checking. Retained because it is plausible and widely repeated, but do not budget on it — treat these cells as one step below every other **M** figure in this table.

WARNING: **The "Peak" column mixes workloads and planes** and is not a like-for-like ranking: some cells are HPL at the wall, some are Linpack at the DC input, and the X4's is a Geekbench multi-core transient. Each cell names its workload; compare only cells that name the same one.

### Annual energy cost

kWh/yr = W × 8.766. Cost at **$0.15/kWh** — substitute your own rate. **$/yr is computed from the unrounded wattage**, not from the rounded kWh column, so recomputing from the middle column may differ by a cent.

WARNING: **This table is ordered across two measurement planes and is therefore only approximately a ranking.** The plane marker is on every row; a (DC) figure and a (AC) figure one row apart may be the same board consuming the same power, differing only by adapter loss. Compare within a plane.

| Board (idle) | W | kWh/yr | $/yr |
|---|---|---|---|
| ZERO 3W | 0.9 (DC) | 7.9 | $1.18 |
| CM5 | 1.2 (AC) | 10.5 | $1.58 |
| ROCK 4D | 1.5 (DC) | 13.1 | $1.97 |
| ROCK 5C | 1.6 (AC) | 14.0 | $2.10 |
| E52C | 1.65 (DC) | 14.5 | $2.17 |
| Cubie A5E | 2.0 (plane unstated) | 17.5 | $2.63 |
| ROCK 5A | 2.9 (DC) | 25.4 | $3.81 |
| Dragon Q6A | 2.9 (AC) | 25.4 | $3.81 |
| ROCK 5B | 3.6 (AC) | 31.6 | $4.73 |
| ROCK 5T | 3.9 (AC) | 34.2 | $5.13 |
| ROCK 5 ITX (bare) | 4.8 (DC) | 42.1 | $6.31 |
| Dragon Q8B | 5.9 (DC) | 51.7 | $7.76 |
| X4 | 9.1 (AC) | 79.8 | $11.97 |
| Orion O6N | 9.3 (DC) | 81.5 | $12.23 |
| Orion O6 | 14.2 (AC) | 124.5 | **$18.67** |
| ROCK 5 ITX (4-drive NAS) | 27.6 (AC) | 241.9 | **$36.29** |

**Worked cluster comparison:** twelve ROCK 5C nodes idle at **19.2 W combined** — about **1.35× a single Orion O6**, for **48 A76 + 48 A55 cores** against the O6's 12 Armv9 cores. (The RK3588S2 is 4×A76 + 4×A55, so "96 cores" is the honest total but only half of them are big cores — an earlier revision of this document said "96 A76 cores", which overstated the comparison by 2×.) Conversely one Orion O6 costs the same annual electricity as **nine ROCK 5C nodes**. The O6 earns its place as one or two "big iron" nodes, never as a swarm.

### Two documented power anomalies

**Orion O6 — the idle leak is real, unfixed, and you should budget for it.** 14.2 W idle at the wall with an efficient 65 W GaN adapter, one low-power NVMe, nothing else attached. Critically, **that figure did not move across two full firmware generations** (DT firmware 0.2.x → SystemReady 9.0.0/1.0.0-3); only load figures shifted. The one public root-cause analysis — filed 2026-04-28 by the meta-cix BSP maintainer — is that Cix's edk2-platforms **disables PCIe ASPM by default on Sky1 root ports** (`AspmSupport` advertises max `0x02` but the active defaults are all `0x00`). Geerling's `lspci -vvv` corroborates: every root port shows `LnkCap: ASPM L1` but `LnkCtl: ASPM Disabled`. **That issue is still open with zero comments. There is no fix, no fix version, and no post-fix measurement. Do not budget for an improvement.** The same problem is reported on the Minisforum MS-R1, which is essentially the same design.

Separately, the O6's BIOS **does not fully power down cards in the PCIe slot on shutdown** — a reviewer's GPU went to failsafe full-blast fan speed after the board was off, and forum users reported the same.

**ROCK 5 ITX component-level breakdown** (12 V DC plane, from a 10GbE NAS build) — the only board where the deltas are actually isolated:

| Config | W | Delta |
|---|---|---|
| Full NAS idle (4× SATA SSD + 10GbE NIC) | 12.20 | — |
| All SSDs removed | 9.74 | **4× SATA SSD = 2.45 W** |
| + 10GbE link down | 7.16 | **10GbE link = 2.58 W** |
| + 10GbE card removed | 4.88 | **card at link-down = 2.28 W** |

Note the 12 V jack is **5.5/2.5 mm, not the common 5.5/2.1 mm** — likely deliberate, to stop people plugging in a 19 V laptop brick.

**The NVMe power adder is unmeasured on every Radxa board.** No source isolates it. The only vendor hint is the ROCK 5A's "≥24 W bare / ≥30 W with M.2 SSD" — that's PSU headroom, not draw.

---

## 7. PoE

### The HAT lineup

| Product | Standard | Output | Documented compatibility |
|---|---|---|---|
| **Radxa 25W PoE+ HAT** | 802.3af/at | 5 V/4.8 A **or** 12 V/2.1 A | ROCK 3A/3B/3C, 4C+, **5A, 5B**, Cubie A7A |
| **Radxa 25W PoE+ HAT for X4** | **802.3at only** | as above + 12 V terminal block | **X4, Dragon Q6A** |
| **Radxa 20W PoE+ HAT** | 802.3af/at | 5 V/4 A | **Cubie A5E** |
| **Radxa 25W PoE+ Module** | `no data` | 25 W max | **ROCK 5 ITX, ROCK 5T** |
| **ZERO 3E PoE HAT** | `no data` ¹ | `no data` | ZERO 3E |
| Radxa 23W PoE HAT (legacy) | 802.3at | 5 V/4.6 A | ROCK Pi 4 family, ROCK 3A/3C |

¹ Radxa's page says 802.3af/at; retailers say af only. **Radxa publishes no electrical datasheet for it.**

### Does it fit the budget?

WARNING: **Every verdict below assumes a lossless HAT.** **No measurement of any Radxa board running through a Radxa PoE HAT exists**, and no HAT efficiency figure is published for any of them. Real PD front-ends are 85–90% efficient, so actual headroom is smaller than shown. Anything marked "marginal" should be treated as **"assume it fails af"** until metered.

**Plane note — this table deliberately uses different numbers from §6.** Where a DC-plane (DC) figure exists it is used here (ROCK 5C 7.4 (DC) rather than §6's 9.5 (AC); X4 12.2 (DC) rather than 18.5 (AC)), because a PoE HAT delivers DC to the board and the adapter loss in the (AC) figures is not part of the PoE budget. **This means the verdicts below are computed on a more favourable plane than §6 reports.** For sizing a switch's total PoE budget — where the switch's own conversion loss *is* real — use §6's (AC) figures instead and expect the marginal rows to get worse.

| Board | Highest sustained | 802.3af (12.95 W)? | 802.3at (25.5 W)? |
|---|---|---|---|
| ZERO 3W / 3E | 2.6 / 3.2 W | **Yes**, wide margin | Yes |
| ROCK 4D | 4.8–7 W | **Yes** | Yes |
| Cubie A5E | ~6 W | Yes (20 W HAT) | Yes |
| ROCK 5A | 6.5 W (15.7 Linpack) | **Marginal** — peak exceeds af | Yes |
| Dragon Q6A | 6.4–10.1 W | **Marginal** — ~2.8 W margin | Yes |
| ROCK 5C | 7.4 W (DC) (16.4 Linpack (DC); 9.5 (AC)) | **No** | Yes |
| ROCK 5B / 5B+ / 5T | 8.6–11 W (DC) (19 Linpack; **5B hits 17 W (AC) under stress-ng**) | **No** | Yes |
| X4 | 12.2 W (DC) (22.1 Linpack (DC); 18.5 (AC)) | **No** | Yes, ~3 W margin |
| ROCK 5 ITX (board only) | 11.5 W; vendor max 25 W | **No** | **Only just — zero headroom for drives** |
| Dragon Q8B | 25.1 W sustained | **No** | **No** |
| Orion O6 / O6N | 23.7–24.9 W | **No** | **No** — and neither has PoE hardware |

### Field reports worth knowing

- **X4 + PoE resets under load.** At the BIOS-default PL2 of 25 W — and even at 15 W — the board reset when a benchmark started, on a cheap 2.5G PoE switch. A UniFi PoE Lite switch handled 15 W fine. **PL2 = 6 W** was stable on both "and didn't impact performance by much." A separate user on a MikroTik 25 W/port switch sees reboots shortly after kernel load. Note the boot-time transient (~22 W in BIOS) *exceeds* steady-state load — that's what breaks marginal supplies at power-on.
- **X4 PoE HAT compresses the NVMe drive dangerously** at the standoff spacing.
- **ZERO 3E PoE is Mode A only and fails on most midspan injectors.** Multiple users could not power it from injectors (including a Ubiquiti 15 W adapter); a 65 W 8-port PoE **switch** worked. Radxa's docs confirm pins 1/2/3/6 — Mode A. Also: power passes through the **screws and copper pillars**, which must be fully tightened or PoE simply won't work.
- **Q6A + PoE HAT works but does not pass through the GPIO pins**, and the HAT fan DTB didn't work for the tester.
- **ROCK 5T and ROCK 5 ITX put PoE on one specific port only** (on the 5T, the port *not* next to HDMI).
- **ROCK 5B+ is on no documented HAT compatibility list.** Radxa staff said on the forum that the X4 variant works. Unofficial.
- The widely-repeated "5B needs 25 W, 5B+ needs 40 W" is a restatement of **recommended PSU ratings**, not measured draw. Don't use it as a budget input.

---

## 8. I/O & expansion

| Board | USB | Display out | GPIO | Notable |
|---|---|---|---|---|
| ROCK 5B | 2× USB3 + 2× USB2, USB-C | 2× micro-HDMI 8K + DP-alt + DSI | 40-pin | HDMI-**in**, 2× CSI |
| ROCK 5B+ | 2× USB3 + 2× USB2, USB-C PD-in | HDMI + DP-alt | 40-pin | dual M.2 |
| ROCK 5A | 2× USB3 + 2× USB2 (2.8 A total) | 2× micro-HDMI + DSI | 40-pin | ES8316 audio codec |
| ROCK 5T | multiple USB3/USB2 | HDMI + USB-C + DSI, **HDMI-in** | 40-pin | camera connectors |
| ROCK 5 ITX | 4× USB3 (**shared GL3523 hub**) + 2× USB2 + front header | 2× HDMI 2.1 (8K+4K) + DP-alt + 2× DSI + eDP — **4 simultaneous** | **NONE** (!) | HDMI-in, S/PDIF, front-panel hdr |
| ROCK 4D | 2× USB3 + 2× USB2 | HDMI 2.1 4K120 + DSI | 40-pin | 2× CSI; USB-C is power-only |
| Dragon Q6A | USB3 + 3× USB2 + USB-C | mid | 40-pin | **3× CSI** |
| Dragon Q8B | 2× USB-C 3.2 Gen2 + 2× USB-A + 2× USB2 | **3× 4K120** | 40-pin | RTC, fan hdr, 3.5 mm |
| Orion O6 | 2× USB-C DP-alt, USB-A 3.2, USB2 | HDMI 2.0 + DP 1.4 | header | **PCIe x16**, RTC, ATX |
| Orion O6N | USB-C DP-alt, USB-A, USB2 | HDMI + DP | header | dual M.2, RTC |
| X4 | 3× USB3.2 Gen2 + 1× USB2 | 2× micro-HDMI 2.0 dual-4K60 | 40-pin **via RP2040** (!) | no HDMI-CEC |
| E52C | 1× USB3-A (also maskrom) | **none** | **none** | **dedicated USB-C console** |
| E54C | 1× USB3-A + 2× USB2 + USB-C | HDMI 2.1 | 14-pin | CH340C console on Type-C |

WARNING: **ROCK 5 ITX has no GPIO header at all** — absent from the connector list and from the schematic index. If your rack build relies on GPIO for status LEDs, watchdog wiring or an I²C OLED, the ITX cannot do it.

WARNING: **The X4's 40-pin header is not a Pi header in software.** The N100 has **no direct GPIO access**; the header is wired entirely to an onboard RP2040 that the CPU talks to over USB 2.0 and UART. Radxa claims Pi pinout compatibility — that is **electrical only**. `libgpiod`, `gpiozero` and existing HAT drivers will not work; every HAT is a porting project. (The X5 upgrades this to an RP2350.)

**RTC and fan headers** are real rack niceties. RTC confirmed present: **ROCK 5B** and **ROCK 5B+** (both schematics carry an **HYM8563TS** with `VCC_RTC` / `RTC_INT_L` on I2C6), ROCK 5A, 5T, 5 ITX/ITX+ (CR1220 + HYM8563), Orion O6 (CR1220), Dragon Q6A (CR2032, enumerates as `rtc-ds1307`), Dragon Q8B, E24C, E54C, X4 (battery header for BIOS timed wake). **RTC confirmed absent: ROCK 4D** (no coin-cell holder on the schematic) and **E52C on boards before BOM revision v1.20A** — Radxa staff confirmed "v1.20A or later has RTC soldered by default," and one user found U17 empty and hand-soldered a HYM8563TS. **Ask the seller for v1.20A.**

---

## 9. Boot media, console, recovery, power-loss

This section exists because these details decide whether an unattended node recovers on its own, and comparisons almost never cover them.

| Board | Boot media | Serial console | Recovery | Auto power-on |
|---|---|---|---|---|
| **ROCK 5B / 5B+** | **16 MB dedicated SPI NOR** + eMMC + SD + NVMe. **NVMe-direct boot confirmed with no SD/eMMC** | 40-pin pins 8/10, `1500000n8` | Maskrom + Recovery buttons, `rkdeveloptool` | Forum claim only (!) |
| **ROCK 5C / 5A** | eMMC **or** SPI module (shared connector) + SD | 40-pin UART only | 5A: recovery + maskrom **pins**; 5C: power button only | `no data` |
| **ROCK 5T** | onboard eMMC + SD + 2× NVMe | 40-pin UART ×2 | **Power + Recovery + Maskrom buttons** | `no data` |
| **ROCK 5 ITX** | SPI → eMMC → SD; NVMe boot needs EDK2 UEFI in SPI | **Dedicated 3-pin debug header**, 1500000 baud | Maskrom button + recovery header | Yes **Confirmed — ATX rail always live** |
| **ROCK 4D** | 16 MB SPI + eMMC/UFS + SD; SPI updatable via `rsetup` | `no data` | `rkdeveloptool` | `no data` |
| **Dragon Q6A** | SPI NOR + SD + eMMC/UFS + NVMe | 40-pin UART | **EDL button**, QDL `05c6:9008`, `edl-ng` | `no data` (!) |
| **Dragon Q8B** | **Standard UEFI**, unified ARM ISO boot; SD + UFS + 2× NVMe | 40-pin UART | EDL button | `no data` |
| **Orion O6** | **8 MB SPI in a removable SOP8 socket** (EDK2/Tianocore) → eMMC/SD/NVMe/USB | 4× UART, console on UART2 | **Pull the chip, reflash w/ CH341A**; normally USB FAT32 + `startup.nsh` | `no data` |
| **Orion O6N** | same, **SPI soldered down** | `no data` | **Requires physically removing the soldered chip** | `no data` |
| **E52C** | eMMC + SD (SD preferred when present) | Yes **Built-in CH340 on a dedicated USB-C port** | Maskrom pinhole; USB-A **needs an A-to-A cable** | `no data` |
| **E54C** | 16 MB SPI + SD (retail SKUs have no eMMC) | CH340C on the shared OTG/DP Type-C | `no data` | `no data` |
| **X4** | UEFI/BIOS + eMMC + NVMe + SD; **PXE netboot**, WoL, RTC timed wake | RP2040 (no dedicated console) | BIOS file package | `no data` (!) |
| **ZERO 3E** | **microSD only — no fallback** | listed, pinout `no data` | Maskrom button on the back | `no data` |
| **CM5** | SD or onboard eMMC. **Carrier SPI is only 64 KB — too small for firmware**, so no SPI boot | `no data` | Maskrom button | `no data` |

### The three things that matter most here

**1. Auto-power-on after a power cut is undocumented on every board except the ROCK 5 ITX.** Radxa staff confirmed the ITX's ATX rail is designed to stay continuously powered — *"it won't lose power when you turn it off"* — so mains returns and the board comes back. That is exactly the behaviour an unattended rack wants (the trade-off: you can't power it down from software; users resorted to Tasmota smart plugs). **Everything else rests on a single 2024 forum exchange between two community members**, with the original poster noting he'd never seen it documented, while the ROCK 5C power docs describe **only** a manual button press. **Warning:** Bench-test this by pulling the plug before you commit any node to a rack. A board that stays dark after a brownout is a dead node.** Also: do not carry over the "State After G3 / remove the RTC battery" trick — that thread is about the **Rock Pi X**, an older x86 board, not the X4.

**2. There is no BMC, IPMI, or iKVM anywhere in the Radxa lineup — or on any shipping carrier that supports a Radxa module.** Closest substitutes, in descending usefulness: **E52C** (built-in USB console, no adapter needed), **Orion O6/O6N** (standard UEFI + 4 UARTs, serial-concentrator friendly), **X4** (x86 with documented WoL, PXE and RTC timed wake). For real out-of-band you need an external networked serial concentrator plus a switched PDU, or a PiKVM-class device per node.

**3. Serial console cables: budget CH340.** Only the E52C and E54C have onboard USB-UART bridges. Everything else needs an external TTL adapter at **1.5 Mbaud**, and Radxa specifically warns that some CP210x and PL2303x adapters can't reach that rate and some FT232RL parts have power problems. **CH340 cables are the recommended part.**

**Hardware watchdog: present in silicon, off by default, unverified on every board.** RK3588/S device trees define `snps,dw-wdt` exposing `/dev/watchdog` with roughly a 44-second window, but **it is disabled by default and must be enabled in the DTS**. No Radxa documentation confirms the watchdog is wired and enabled on any specific board. Verify on hardware before relying on it.

WARNING: **Dragon Q6A — two reports of a cold-boot DDR-training failure that drops the board into EDL.** Worth knowing about, but read the thread before weighting it heavily: there are **two** users, and Radxa's moderator concluded *"the DDR hardware may be broken according to the log"* and directed the reporter to warranty. That reads as a **possible unit defect handled by RMA**, not an established model-level property. An earlier revision of this document treated it as a headline caveat against the board; on the evidence available that was too strong. It remains a reason to bench-test cold-boot behaviour on arrival — which you should do on any unattended node anyway.

---

## 10. Thermals & cooling

- **RK3588 (ROCK 5 family)** throttles under sustained all-core load without active cooling; a heatsink+fan is effectively mandatory for 24-7. Very manageable — ~10–20 W to dissipate. The ROCK 5 ITX measured **40.7 °C idle, 51.8 °C max at ~35 °C ambient with no throttling**, and its cooler mounting is **75×75 mm Intel LGA-115x compatible**, so ordinary desktop coolers fit.
- **ROCK 5 ITX fan header is 4-pin PWM but has no tachometer** — you cannot read RPM.
- **Radxa X4 requires active cooling, non-negotiably.** Without it the N100 "would overheat to the point where the built-in safety mechanism would power off the system in less than 15 minutes." Cooling is **not included** (~$15 extra), and **the stock fan is always-on at 100% with no PWM control**. Measured: idle 42 °C, single-core OpenSSL peaked **89 °C**, and **raising PL1 to 15 W drove temps to 93–97 °C**. A large third-party heatsink fixes it (a salvaged server heatsink held ~48 °C idle / ~75 °C at full load).
- **Cubie A5E idles at ~70 °C bare-board**, dropping to ~43 °C in its metal case. **Budget for the case on every unit** — 70 °C idle is disqualifying in a dense rack.
- **ZERO 3E passive cooling is officially adequate** for "light-to-moderate workloads" per Radxa's brief — which is exactly the DNS/proxy/monitoring case.
- **E52C runs cool**: 45.3 °C SoC idle at 24 °C ambient, 49–50 °C peak under SQM+AdGuard+Tailscale. Note **OpenWrt runs 4–5 °C hotter at idle than Armbian** because OpenWrt's OPP table doesn't scale below 1 GHz while Armbian goes to 408 MHz.
- **Q8B and the Orions dissipate 25–40 W under load** — treat them as mini-PCs with directed airflow, not passive blocks. The Q8B's PMICs failed under sustained Linpack even with industrial cooling.

---

## 11. Form factor & mounting

| Board | Size | Mounting | Rack fit |
|---|---|---|---|
| ROCK 5B / 5B+ | ~100×72 mm | 4× M2.5 | Pi-style trays, DeskPi/UCTRONICS |
| ROCK 5C / 5A | 85×56 mm | Pi-like ¹ | Densest; standard Pi mounts |
| ROCK 5T | 110×80 mm | 4× standoff | Tray-friendly |
| ROCK 5 ITX / ITX+ | **170×170 Mini-ITX** | ITX standard; **75×75 LGA-115x cooler holes** | 1 per ITX slot / 2U+ |
| ROCK 4D | 86×56 or 87×58 ² | `no data` | Pi-style trays |
| Dragon Q6A | SBC (Pi-ish) | standoffs | Pi-style trays |
| Dragon Q8B | 100×75 mm | standoffs | Custom tray |
| Orion O6 | **Mini-ITX 170×170** | ITX standoffs | 1× ITX slot / 2U+ |
| Orion O6N | **Nano-ITX 120×120** | ITX-family | Smaller ITX / custom |
| E20C / E52C | 66×66 board, 72×72×28.7 case ³ | `no data` | Fanless metal box |
| E24C / E54C | 143×99×25.3 or 130×85×24 ² | `no data` | Wall/rack per vendor |
| X4 | 85×56 mm | WARNING: **NOT Pi-case compatible** ⁴ | Needs custom tray |
| ZERO 3E / 3W | 65×30 (72×30 w/ RJ45) | `no data` | Very dense |
| Cubie A5E | ~69×56 mm ⁵ | `no data` | Pi-ish trays |

¹ **Do not assume the Pi 58×49 mm hole pitch** — the mechanical drawings are vector images with no extractable dimensions. Verify before designing a sled.
² Sources conflict; neither is vendor-authoritative (the ROCK 4D product brief has no mechanical section at all).
³ Radxa docs vs CNX disagree (67×67×15 per CNX).
⁴ The X4 "protrudes by a few extra millimetres along the GPIO edge" and won't fit standard Pi cases. Radxa explicitly fixed this on the X5.
⁵ **Warning:** Two PCB revisions (V1.1/V1.2) have slightly different dimensions** and Radxa warns to verify against current production files before designing an enclosure. Relevant if you 3D-print sleds.

---

## 12. Network appliance tier (E-series)

Purpose-built router/firewall boxes in fanless metal cases. Treated separately because their tradeoffs are unlike the SBCs.

| | E20C | E24C | E25 | **E52C** | E54C |
|---|---|---|---|---|---|
| SoC | RK3528A | RK3528A | RK3568 | **RK3582** | RK3582 |
| Ports | 2× GbE | 4× GbE | 2× 2.5GbE | **2× 2.5GbE** | 4× GbE |
| NIC path | 1 RGMII + 1 PCIe | **switch on 1 RGMII** | PCIe | **dedicated PCIe lane each** | **switch on 1 RGMII** |
| CPU-path ceiling | ~1 Gbps/port | **~1 Gbps total** | ~2.5 Gbps/port | **~2.5 Gbps/port** | **~1 Gbps total** |
| M.2 | — | Gen2 ×1 | 2242 SATA + mPCIe + B-key | **none** | Gen2 ×1 |
| Console | Type-C | Type-C (shared) | Type-C (shared w/ power) | **dedicated, CH340** | Type-C (shared) |
| Idle | `no data` | 5 W (V) | `no data` | **1.65–2.11 W (M) (DC)** | `no data` |
| **Mainline OpenWrt** | Yes 25.12+ | No **no** | Yes **24.10 LTS + 25.12** | Yes 25.12+ | No **no** |
| Armbian | vendor only | vendor + edge | **mainline 6.18** | **mainline 6.18** | vendor only |
| Price (ARace) | not listed | not listed | not listed | $55/$65/$85 | $55/$65/$85 |
| Supply guarantee | 2034 | 2034 / 2030 ¹ | — | **2034** | 2034 |

¹ Brief says Sept 2034; radxa.com says "at least 2030."

**The E52C is the right answer in this family** and it isn't close: dedicated PCIe lane per NIC, measured sub-2.5 W idle, mainline OpenWrt *and* mainline Armbian, a dedicated always-on serial console, and supply guaranteed to 2034. Reported stability is good — 20+ days uptime as a production router.

**Known E52C issues for a 24-7 node:**
1. `kmod-r8125-rss` resets the port under high connection counts (`NETDEV WATCHDOG: transmit queue 1 timed out`), reproducible with `iperf3 -P 30 -t 30`.
2. One user's WAN link flaps at 2.5G against an Adtran ONT on stock `r8169` but is stable on iStoreOS; others report 2.5G fine on 25.12.x.
3. Upgrading across the interface rename needs `sysupgrade -n` — config isn't migratable.
4. **Thin USB-C cables cause brownouts.** One user: *"voltage dropped a lot when the cores were utilized… bought the 'old' RPi 5.1V/3A supply which uses 18 AWG and never looked back."* A Samsung 45 W adapter caused reboots for another. The USB-A port's current budget is undocumented — a 5G dongle bootlooped on it.
5. **RTC only on BOM revision v1.20A and later** (see §8).

**The E25 has the longest OpenWrt track record** — the only E-series board in the 24.10 LTS line — but it's legacy, ~2× the price, and out of stock.

---

## 13. Compute modules & cluster density

### Modules

| Module | SoC | Connector | RPi CM4 compatible? | Measured power |
|---|---|---|---|---|
| **CM5** | RK3588S2 | 3× 100-pin | "somewhat" — **8** documented changes | **1.2 W idle / 8.4 W load** (AC) |
| **CM5 Lite** | RK3582 | 3× 100-pin | as CM5 | 2.4 W (DC) / 6.7 W (DC) |
| **CM4 (Radxa)** | RK3576 | 3× 100-pin | claimed, caveats undocumented | `no data` |
| **CM3J** | RK3568J | **2× 100-pin — true CM4 match** | physically yes; **GPIO alt-functions differ** | `no data` |
| **CM3 / CM3I / CM3S** | RK3566 / RK3568 | 3–4× 100-pin / SODIMM | partial | `no data` |
| **NX5** | RK3588S | 260-pin SODIMM | Jetson-shaped, **not Jetson-compatible** | **3.6 W idle (DC) / 10.9 W load (DC)** (18.8 Linpack) |

**CM5 drop-in compatibility — what the FAQ actually covers.** Radxa's CM5 FAQ documents **8** functional changes, and its subject is **replacing a CM3 with a CM5 on a CM3-derived baseboard** — not, as an earlier revision of this document framed it, running a CM5 on a Raspberry Pi CM4 carrier. Documented: no SPI flash; no 3.5 mm audio (needs an external codec); eDP pins 15/17/21/23/27/29/33/35 unsupported; USB2_HOST3 doesn't exist; MIPI limited to 2-lane. The **only** item in the FAQ that concerns a Raspberry Pi carrier is that **microSD boot on the RPi CM4 IO board requires pin 90 of the third connector to be grounded**.

Three claims that appeared here previously — that the third connector is unavailable, that the CM5 is single-HDMI-only with HDMI "frequently dead in practice", and that an RPi CM4 will not boot on the Radxa CM5 IO Board — are **not in the cited FAQ or any other source reached**, and have been removed rather than left unattributed.

[X] **Radxa's two hard warnings, quoted exactly.** On the CM3 IO board: *"Can Radxa CM5 be used on Radxa CM3 IO board — No, as the 3 pins (94/96/98) (USB_5V_IN) of the current CM3 + Radxa CM3 IO board are connected to 5V, while CM5 is GPIO."* And on custom CM3 baseboards that tie those pins to 5 V: *"…then cannot replace with CM5, otherwise the board will be damaged."* The single-sentence version of this warning circulating elsewhere (including in an earlier revision of this document) splices the two answers together and is not something Radxa wrote.

### Carriers

| Carrier | Nodes | Radxa CM5? | BMC/KVM | Status |
|---|---|---|---|---|
| **Radxa CM5 IO Board** | 1 | native | none | Exists; **ARace does not carry it** — price `no data` ⁑ |
| **Radxa Taco v1.61** | 1 | unconfirmed | none | **In stock, $65** — 5× SATA + NVMe + dual NIC, PCIe 3.0 switch |
| **Xerxes Pi** | 1/board, **12 per 1U** | claimed | none | [X] **UNVERIFIED — vendor could not be confirmed to exist** ⁂ |
| **DeskPi Super6C** | 6 | Yes **Radxa-certified** | none | Shipping; 17 W idle for 6 CM4 nodes |
| **DeskPi Super4C** | 4 (RPi CM5) | `no data` | ESP32 power control only | Shipping; dual NIC + dual PSU per node |
| **Turing Pi 2.5** | 4 | No **USB routing mismatch** | Yes **real BMC + network eMMC flashing** | Shipping, $279 |
| **Compute Blade** | 20 per 1U | not named | none (PoE port cycling) | Shipping |
| **Cerebro** | 4 | claimed | claimed BMC+KVM | [X] **CANCELLED — does not exist** |

⁑ The Radxa CM5 IO Board is a real product, but it is **absent from ARace's entire 103-product catalogue and from ARace's own search**. Since ARace is this document's pricing source of truth, it has no price here. (The only Radxa IO boards ARace stocks are the NX4, CM3 and CM3I.) A "$26.99, in stock" entry in an earlier revision was unsourceable.

⁂ **The Xerxes Pi could not be verified and should not be designed around.** Neither `xerxes-pi.com` nor `xerxespi.com` resolves; it appears in no source in this document's Sources block; and Uptime Industries — cited by name below — ships the Compute Blade, not this. **Everything in the density and cost model that follows is therefore conditional on a product this fact-check could not confirm exists.** It is retained here only because the arithmetic is still useful as an upper bound on what 12 CM5 nodes in 1U would look like. **Confirm the vendor independently before it enters a build plan.** If it doesn't check out, the shipping alternatives are the DeskPi Super6C (6 nodes, Radxa-certified) and Compute Blade (20/1U, CM5 support not stated).

[X] **The Cerebro clusterboard is vaporware.** cerebro-board.com — as archived in April 2026 — **still advertises its "Kickstarter will be launched around April 2025" as a future event**, alongside claims of an onboard KVM + BMC, 4 nodes and compatibility with "4 NVIDIA Jetson NX- and Raspberry Pi CM4/CM5 – and Radxa CM5 modules." The campaign is widely reported to have been **cancelled by the creators in June 2025 at €12,017 of a €1,081,284 goal (24 backers, ~1.1%)**, but **Warning:** those funding figures could not be verified** — Kickstarter blocks automated access. The site's own frozen-in-the-past launch notice is the load-bearing evidence, and it is sufficient: if a rack design depends on the Cerebro's BMC/KVM, it has no path forward.

### Density and the failure-domain tradeoff

**12× CM5 in 1U** (per-node figures from Geerling's CM5-on-IO-board measurements; totals are arithmetic). [X] **The 12-per-1U carrier assumed here is the unverified Xerxes Pi** — read the per-node column as solid and the ×12 column as a hypothetical:

| | Per node | ×12 |
|---|---|---|
| Cores | 4×A76 + 4×A55 | **96 cores** |
| RAM (32 GB SKU) | 32 GB | **384 GB** |
| NPU | 6 TOPS | 72 TOPS |
| HPL | 48.6 GFLOPS | ~583 GFLOPS |
| Idle | 1.2 W | **~14.4 W** |
| Load | 8.4 W | ~101 W |

Cost, **at the unverified Xerxes prices**: ~$198/node (Xerxes **PoE variant $79** + CM5 8/64 $119), ~**$2,376** per 1U plus a switch. The PoE variant is the right one to price, since the whole argument for blades over shared carriers is that out-of-band management degrades to cycling a PoE port. (The $59 non-PoE figure quoted in an earlier revision undercounted by ~$240/U.) **Density is not the constraint here — PoE budget and switch port count are.**

**The architectural choice:**

- **Shared-carrier** (Turing Pi 4, Super6C 6, Super4C 4): one carrier failure kills **every node on it** — a Super6C fault takes out 6 nodes at once, so a 3-node etcd quorum on one carrier is a *total* quorum loss, not a survivable minority failure. One PSU and one switch die feed all nodes. **Upside: these are the only designs with a BMC.** The Super4C is the only one that mitigates, with dual 19 V inputs auto-switching and two independent NICs per node.
- **Blades** (Compute Blade 20/1U; Xerxes Pi 12/1U [X] unverified): failure domain is **one node**. Uptime Industries states this as an explicit design goal for the Compute Blade — the shared component becomes a commodity PoE switch you can keep a cold spare of. Cost: no BMC; out-of-band degrades to "toggle the PoE port and let it netboot."

WARNING: **No confirmed shipping product gives you both a real BMC and validated Radxa CM5 support.** Turing Pi 2.5 has the BMC but not the CM5; Super6C is Radxa-certified but has no management at all; the Xerxes Pi is claimed to have CM5 support and no BMC, but [X] could not be verified to exist. That gap is precisely what the Cerebro was designed to fill.

---

## 14. Software / OS support

| Platform | Maturity |
|---|---|
| **RK3588 / RK3588S / S2** | **Best.** Years of Armbian/Debian/Ubuntu, good mainline progress, huge community, mature vendor kernel. Safest for unattended uptime. |
| **RK3582 (E52C, 5C Lite)** | Same Rockchip base; **E52C has mainline OpenWrt (25.12+) and mainline Armbian 6.18**. E54C has **neither** — vendor kernel only. |
| **RK3576 (ROCK 4D, CM4)** | Newer. Debian 12 + Yocto/Buildroot/**Android 14** (newest Android in the lineup). |
| **Qualcomm QCS6490** | Ubuntu 24.04 with a `-qcom` kernel works well; upstream Qualcomm Linux effort is real but young. |
| **Snapdragon 8cx Gen 3 (Q8B)** | **Standard UEFI + unified ARM ISO boot** — a genuine advantage. Benefits from Windows-on-ARM and 8cx Linux work, but only one independent test dataset exists. |
| **CIX P1 (Orion)** | Newest. Debian 12 / Ubuntu 25.04, active development, **firmware measurably changes performance**; peripheral maturity still catching up. |
| **Intel N100 (X4)** | **Best of all in one sense** — conventional UEFI, stock distro ISOs, Windows, PXE, no device trees or vendor kernels. Mainline `igc` NIC driver. |
| **Allwinner A527 (Cubie A5E)** | WARNING: **Weakest.** Mainline supports "basic headless use" only — **USB 3.0 and the second Ethernet port are still WIP upstream.** Full function requires Radxa's BSP kernel with closed binaries. The dual-NIC feature you'd buy it for is the part mainline doesn't drive. |

**For a set-and-forget rack, RK3588-family boards remain the low-risk choice.** The Qualcomm and CIX boards trade some maturity for substantially more compute. The X4 trades ~3× the idle power for the least exotic software stack in the entire comparison.

**Why an x86 node in an ARM rack** — the concrete, sourced case: **QuickSync**. On the N100, H.264 and HEVC (8- and 10-bit) have **hardware decode *and* encode**, and AV1 has hardware decode (AV1 *encode* is explicitly not supported on N-series — that needs Arc/DG2). Requires kernel ≥6.2 and Intel compute-runtime 23.xx+. A container gotcha: if the container user isn't in the `render` group, transcodes **silently fall back to CPU**. No ARM SBC here has an equivalent; a Pi 5 has no hardware HEVC encode at all. Secondary reasons: amd64-only container images, AES-NI (1230 Mbps OpenSSL), and standard UEFI/PXE.

---

## 15. Recommendations by rack role

NOTE: §16 is the task-first companion to these bullets — it defines every workload named here, plus the resource-tag vocabulary, without reference to any board. **Its scoping rule governs how much a bullet below should say:** name only the workloads a board's hardware genuinely decides, not everything it can run. Where a board is just a competent generalist, say that and list its caveats instead.

- **Dense low-power services (DNS, monitoring, reverse proxy), PoE-fed:** **ROCK 5C** (1.6 W idle (AC), Pi footprint, 802.3at) — **$69.00** 2 GB / **$159.00** 8 GB. Two things weaken this pick: it's on **no documented PoE HAT compatibility list**, and the **Radxa 25W PoE HAT is itself out of stock at ARace** ($19.99) — so verify the HAT before committing to the topology. **ZERO 3E** (**$39.99**, 2 GB) if 1 GB/2 GB is enough and you feed PoE from a **switch, never an injector** (its own $9.99 PoE HAT *is* in stock) — accept SD-only boot with no fallback.
- **Router / firewall:** **E52C** — dedicated PCIe lane per NIC, 1.65 W idle, mainline OpenWrt, built-in console. Ask for BOM rev v1.20A. Avoid E24C/E54C (~1 Gbps aggregate ceiling, no mainline OpenWrt).
- **NAS:** **ROCK 5 ITX** is the only real Radxa option — 4× SATA on a Gen3 ×2 ASM1164, dual independent 2.5GbE, SMB Multichannel measured ~600 MB/s. Set expectations by the **~400 MB/s 4-SSD RAID-0** figure, not the 1.97 GB/s link. Budget 27.6 W idle populated, plus an **unmeasured but non-zero soft-off draw** — the ATX rail stays live and nobody has metered it. Alternative: a CM5 on a **Radxa Taco** ($65, in stock) for 5× SATA on one node.
- **Edge-AI / camera ingest with PoE:** **Dragon Q6A** — best perf/watt of any standalone board here (4.79 GFLOPS/W (AC); only the CM5 *module* edges it at 4.86), 3× CSI, officially supported PoE HAT, and one of the few boards actually in stock at **$169.00 for 12 GB** — the best RAM-per-dollar of any in-stock SKU. Bench-test cold-boot recovery on arrival (see §9), but don't over-weight the EDL reports — two users, and Radxa treated it as an RMA.
- **Heavy compute / build node:** **Dragon Q8B** (fastest ST and MT; standard UEFI) or **Orion O6** (best memory bandwidth, PCIe x16). Budget the O6 at 14.2 W idle permanently.
- **Media transcoding / amd64 workloads:** **Radxa X4**. Accept ~9 W idle, mandatory active cooling, an RP2040-mediated GPIO header, and marginal PoE.
- **Cluster:** **CM5 modules**, on the densest carrier you can actually buy. The **DeskPi Super6C** (6 nodes, Radxa-certified, shipping) is the safe pick, at the cost of a 6-node shared failure domain. The 12-node/1U **Xerxes Pi** would be better on density and failure isolation (~14.4 W total idle, one-node domain, no BMC) — but [X] **this fact-check could not confirm the vendor exists**; verify it independently before planning around it. See §13.
- **Mixed rack (the practical answer):** several ROCK 5C/CM5 nodes for services, one E52C as the router, one ROCK 5 ITX for storage, and — if you need transcoding or x86 images — one X4. Add an Orion O6 only if you specifically need PCIe expansion or memory bandwidth.

---

## 16. Workload catalog — tasks & services

A task-first index of workloads these nodes plausibly run, and the hardware property that actually decides each one. This is the inverse of the device-first tables elsewhere in this document: use it to go from "I need to run X" to "so I need hardware that does Y". It is deliberately hardware-neutral — no device is named here. The tags are the controlled vocabulary the rest of this document should use when describing what a node is for.

**Scoping rule — this list is exhaustive, per-device recommendations are not.** When describing what a given device is for, name only the handful of workloads its hardware genuinely decides: the things it is better at than the boards next to it in the table. Do not enumerate everything it is merely capable of. Nearly every node here can run nearly everything on this list, so a long list carries no information. Where a device is simply a competent generalist, say exactly that in one line and spend the remaining space on its caveats — what it cannot do, where it runs out of headroom, and what it forces you to work around. "Good all-rounder, but single-lane NVMe and no hardware encoder" is worth more to a reader than thirty task names.

### Resource-tag legend

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

### Media & streaming

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

### Vision, camera & sensing

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

### Web scraping & data acquisition

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

### Data stores & state

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

### AI & ML inference

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

### Build, CI & software supply chain

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

### Networking & edge

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

### Storage, backup & sync

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

### Orchestration & virtualization

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

### Observability & operations

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

### Home automation, IoT & physical control

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

### Custom application runtime patterns

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

### Communication & collaboration

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

### Security, identity & privacy

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

### Batch, scientific & specialty compute

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

### Common mis-placements

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

## Availability & pricing (ARace, 2026-07-22/23)

### In stock right now — seven boards

| Board | SKU | ARace | Flag — cross-checked vs ALLNET China, 2026-07-23 |
|---|---|---|---|
| ROCK 5B+ | 8 GB | **$175.00** | [RESTOCK WATCH] cheaper but **sold out**: ALLNET 4 GB $79.99. ALLNET's in-stock 8 GB is **$239.99** — ARace wins by $64.99 |
| ROCK 5C | 2 GB | **$69.00** | ALLNET in stock at $99.99 — ARace wins by $30.99 |
| ROCK 5C | 8 GB | **$159.00** | [RESTOCK WATCH] cheaper but **sold out**: ALLNET **16 GB** $159.99 — double the RAM for the same money if it restocks. ALLNET's in-stock 8 GB is $219.99 |
| ROCK 5T | 8 GB Commercial | **$179.90** | ALLNET in stock at $239.99 — ARace wins by $60.09. [RESTOCK WATCH] ALLNET **also stocks 8 GB + 64 GB eMMC at $329.99**, a config ARace cannot sell at all |
| Dragon Q6A | 12 GB | **$169.00** | [BUY ELSEWHERE] **cheaper option IS in stock elsewhere:** ALLNET **8 GB at $139.00**, $30 under ARace's 12 GB. ALLNET's 12 GB is also $169.00 (tie). [RESTOCK WATCH] 6 GB $119 at ALLNET is sold out |
| ZERO 3E | 2 GB | **$39.99** | [RESTOCK WATCH] cheaper but **sold out**: ALLNET 2 GB $28.99 (no header) / $29.99 (header). **Every** ALLNET ZERO 3E SKU is sold out |
| ZERO 3W | 2 GB + 16 GB eMMC, header soldered | **$39.99** | [RESTOCK WATCH] cheaper but **sold out**: ALLNET 2 GB + 16 GB eMMC $32.90 / $33.90. **Every** ALLNET ZERO 3W SKU is sold out |

**The one live arbitrage:** the **Dragon Q6A 8 GB at ALLNET for $139** is the only case where another retailer beats ARace on a SKU you can actually buy today. It trades 4 GB of RAM against ARace's $169 12 GB — worth the $30 for a RAM-hungry node, not worth it for a service node.

**Plus three accessories in stock:** Radxa Taco carrier **$65.00** · Orion O6 AI PC Kit (case) **$39.00** · ZERO 3E PoE HAT **$9.99**.

WARNING: **The Radxa 25W PoE HAT — the one every ROCK 5-series PoE plan depends on — is out of stock** ($19.99). Any PoE-fed topology in §7 or §15 is gated on sourcing it elsewhere.

**Pre-order:** Dragon Q8B 8/16/32 GB ($209/$329/$569, **ships 31 Jul 2026**) · Orion O6N 16 GB and 32 GB @5000-5500 ($309/$549)

**Sold out at ARace:** ROCK 5B, ROCK 5A, ROCK 5 ITX, ROCK 5 ITX+, ROCK 4D, Orion O6 (board — only the $39 case kit remains), E52C, E54C, X4, Cubie A5E, CM5/CM5 Lite

**Not carried at all:** Dragon Q5E, X2L, E20C, E24C, E25, CM3/CM3I modules

WARNING: **Sold-out variants still display prices, and several are visibly stale or broken:** ROCK 5T 12 GB at $89 but 8 GB at $179.90; ROCK 4D 4 GB ($69) above 8 GB ($60); Orion O6N 32 GB @6000 at $275 below 16 GB @5500 at $309; Industrial SKUs at $9,999 placeholders. **Only the in-stock prices above are live.** Expect upward re-pricing on restock.

**Backdrop:** the 2026 DRAM crisis. CNX's 2024-vs-2026 price comparison — the one sourced figure here — measured **ROCK 5B+ 8 GB at $90 → $129.99 (+44%)** and **Radxa X4 8 GB at $79.96 → $265.99 (+233%)**, attributing it to "the increasing price of RAM… since late 2025 due to AI demand". (!) The widely-circulated specifics (DRAM +90–95% QoQ in Q1 2026, +58–75% forecast for Q2, a Raspberry Pi "seven-fold increase") **appear in none of the sources used here** and have been dropped rather than left uncited — if you need them, go to TrendForce and Raspberry Pi directly.

**On cross-retailer comparison:** do it — but only against a Radxa retailer whose live stock can actually be read, which means **ARace and ALLNET China**. The flags in the table above are the result of checking all seven in-stock ARace SKUs against ALLNET's per-variant data on 2026-07-23. Result: **ARace is cheaper on six of seven**, ALLNET wins the seventh (Q6A 8 GB $139), and on the Q6A 12 GB they tie exactly at $169. That is a per-SKU outcome, not a general rule — an earlier revision of this document drew a general "ARace is cheaper" conclusion from two *sold-out* listings, which was not a valid basis.

What still doesn't work is comparing against sold-out listings anywhere. ameriDroid lists the ROCK 5B at $289.95 against ARace's $99.90, but **both are sold out** — the same trap as ALLNET's $999 and ARace's $9,999 placeholders. One genuine live datapoint for scale: ARace's in-stock **ROCK 5B+ 8 GB at $175** sits **~35% above** CNX's April-2026 reference of $129.99, and ALLNET's live $239.99 for the same SKU sits **~85% above** it.

### Lifecycle

Nothing here is EOL. Radxa publishes supply guarantees: **ROCK 5 family and CM5 to Sept 2032**, **ROCK 4D, ZERO 3E/3W and the entire E-series to Sept 2034**, **Radxa CM4 to 2035**, **Cubie A5E to Sept 2032**. The ROCK 5A and ROCK 4D briefs were both revised 2026-06-26 specifically to update availability language. **The ROCK 5A is functionally de-prioritized** — out of stock everywhere and no longer listed on ARace's ROCK 5 collection page.

### Announced but not shipping

- **Dragon Q5E** (QCS6690, dual 2.5GbE with optional PoE, 65×56 mm; **wireless spec not stated in the announcement** — a "WiFi 7" attribution circulating for it is not in the source) — announced 2026-06-01, **no price, no docs page, not listed anywhere ~7 weeks later.** Do not plan around it.
- **X5** (Intel N150, RP2350, **64-bit memory bus**, Pi-compatible redesign, 12 V DC input) — fully documented with a published schematic, but **no price and no ship date**. It fixes the X4's two biggest measured problems. If your build can wait, it's the better target.
- From Radxa's 22-product 2026 Qualcomm roadmap, two are rack-relevant: **DragonStation** (6-bay M.2 NVMe NAS with 10 GbE) and an unnamed **cluster system on Dragonwing IQ-9075 (200 TOPS)**. Neither has pricing.

---

## Known gaps — what nobody has published

Listed so they aren't mistaken for oversights. **Do not fill these by extrapolation from SoC class or PSU rating.**

**Power:** ROCK 5B+ idle (the single most consequential blank) · ZERO 3E idle · ROCK 5 ITX+ everything · E20C/E24C/E25/E54C everything · ROCK 5C Lite load · CM3/CM3I/CM3J/CM3S/CM4 all power · **ROCK 5 ITX soft-off draw** (the always-live ATX rail is documented; its wattage is not) · E52C DMC-governor power delta · the NVMe adder on every board · **PoE HAT conversion efficiency on all five HATs** (which makes every 802.3af verdict in §7 optimistic by an unquantified margin) · **no measurement of any Radxa board running through any Radxa PoE HAT exists**

**Benchmarks:** E54C/E25/E24C/E20C — no benchmark of any kind · Dragon Q5E, X2L, CM3J, CM3S, CM3, CM3I — nothing from any tester · ROCK 5 ITX+ everything · ROCK 5C Lite Geekbench · NX5 7-Zip single-thread · Q8B has only one independent tester (Geerling's issue #108 is still all TODO) · NPU inference for RK3582, T527, RK3528A and the Q8B's Hexagon · sysbench and STREAM essentially everywhere

**Topology:** CIX P1 total lane count · SC8280XP lane count · QPS615 upstream width on the Q8B · X4 lane allocation beyond its M.2 · ROCK 5T schematic-level PHY mapping · ROCK 5C/5C Lite combo-PHY assignment · Q6A onboard WiFi attachment (its PCIe0 nets are named `_WCN` but terminate at the Ethernet NIC)

**Management:** auto-power-on on every board except the ROCK 5 ITX · hardware watchdog enablement on every board · SPI NOR presence on ROCK 5T and ROCK 5 ITX+

**Throughput:** E52C above 1 Gbps (every tester is ISP-limited) · E24C/E54C anything · routed NAT on E20C · **any >1 Gbps figure for Orion O6/O6N** (both published iperf3 results are test-LAN-limited to ~941 Mbps and say nothing about the 5GbE/2.5GbE path) · **NVMe** on Orion O6/O6N, Q6A, ROCK 5T, ROCK 5C, CM5, ROCK 4D

**Unresolved conflicts:** RK3576 process node (22 nm vs 8 nm; Rockchip states neither) · ROCK 5B idle (1.5/1.7/3.6/4.5 W) · ROCK 5B NVMe read (fio 3.09 GB/s vs iozone 1.04 GB/s vs sbc.compare 1.4 GB/s — same board, same class of drive) · Cubie A5E Geekbench (2.1× spread) · O6N idle (9.3 W measured vs ~15 W vendor) · ROCK 5 ITX PSU sizing with 4 HDDs (65 W docs vs 90 W brief — **use 90 W**) · ROCK 5 ITX RTL8125BG RX (2.32 vs 1.3–1.6 Gbps) · numerous E-series RAM/eMMC/dimension conflicts between Radxa's own briefs and docs

**Claims withdrawn in the 2026-07-23 fact-check** (previously stated here as sourced, now known to be unsupported): the ROCK 5 ITX "8 W in soft-off" figure and the "Radxa claims 1.5 W standby" counterclaim · a merged perf-per-watt ranking mixing two methodologies · "96 A76 cores" for 12× ROCK 5C · a spliced pseudo-quotation of Radxa's CM5/CM3-IO-board warning · three CM5-compatibility items absent from the FAQ · the Xerxes Pi's existence · Cerebro Kickstarter funding totals · DRAM percentage moves · "ARace is materially cheaper" · Dragon Q5E WiFi 7.

**Sources rejected as unreliable during research:** AI-generated content farms circulating fabricated ZERO 3E idle figures (0.8 W, 2.5 W) and a fake "940 Mbps/1.82 Gbit" throughput number · sbcwiki's ROCK 4D page (contaminated with another board's specs) · Boiling Steam's O6N review (republishes Geerling's **O6** power data under an O6N heading; the author states he could not measure) · the "326.1/1693.7" Q6A "Geekbench" figure (it's CPU-Z under Windows). One widely-cited error corrected here: Thomas Kaiser's ROCK 5 ITX article names the boot SPI NOR as a Winbond W25X20CL — per the schematic that part is the **ASM1164's** firmware flash; the actual boot flash is a **16 MB XT25F128BW**.

---

## Sources

**Measurement & benchmarks**
[geerlingguy/sbc-reviews](https://github.com/geerlingguy/sbc-reviews) (issues [#3](https://github.com/geerlingguy/sbc-reviews/issues/3) 5B · [#40](https://github.com/geerlingguy/sbc-reviews/issues/40) CM5 · [#41](https://github.com/geerlingguy/sbc-reviews/issues/41) 5C · [#48](https://github.com/geerlingguy/sbc-reviews/issues/48) X4 · [#62](https://github.com/geerlingguy/sbc-reviews/issues/62) O6 · [#85](https://github.com/geerlingguy/sbc-reviews/issues/85) Q6A · [#108](https://github.com/geerlingguy/sbc-reviews/issues/108) Q8B stub) · [ThomasKaiser/sbc-bench Results](https://github.com/ThomasKaiser/sbc-bench/blob/master/Results.md) · [ThomasKaiser ROCK 5 ITX preview](https://github.com/ThomasKaiser/Knowledge/blob/master/articles/Quick_Preview_of_ROCK_5_ITX.md) · [sbc.compare](https://sbc.compare/) · [bret.dk](https://bret.dk/) · [wtarreau ROCK 5 ITX NAS breakdown](https://wtarreau.blogspot.com/2024/05/an-affordable-10gbe-capable-nas.html) · [LinuxLinks 5T power](https://www.linuxlinks.com/radxa-rock-5t-single-board-computer-running-linux-power-consumption/) · [smarthomecircle Cubie A5E](https://smarthomecircle.com/radxa-cubie-a5e-review-benchmarks-vs-raspberry-pi)

**SoC datasheets & schematics (primary)**
[RK3588 V2.2](https://rockchip.fr/RK3588%20datasheet%20V2.2.pdf) · [RK3588S V1.7](https://rockchip.fr/RK3588S%20datasheet%20V1.7.pdf) · [RK3588S2 V1.3](https://rockchip.fr/RK3588S2%20datasheet%20V1.3.pdf) · [RK3582 V1.1](https://dl.radxa.com/rock5/5c/docs/hw/datasheet/Rockchip%20RK3582%20Datasheet%20V1.1-20230221.pdf) · [RK3576 V1.5](https://rockchip.fr/RK3576%20datasheet%20V1.5.pdf) · [QCS6490 Rev AW](https://dl.radxa.com/q6a/hw/datasheets/80-23889-1_REV_AW_QCS6490_and_QCS5430_Data_Sheet.pdf) · schematics: [ROCK 5B](https://dl.radxa.com/rock5/5b/docs/hw/radxa_rock_5b_v1450_schematic.pdf) · [ROCK 5B+](https://dl.radxa.com/rock5/5b+/docs/hw/radxa_rock5bp_v1.2_schematic.pdf) · [ROCK 5 ITX V1.11](https://dl.radxa.com/rock5/5itx/v1110/radxa_rock_5itx_v1110_schematic.pdf) · [ROCK 5A V1.1](https://dl.radxa.com/rock5/5a/docs/hw/radxa_rock5a_V1.1_sch.pdf) · [ROCK 4D V1.11](https://dl.radxa.com/rock4/4d/docs/hw/Radxa_ROCK_4D_SCH_V1.11.pdf) · [Dragon Q6A V1.21](https://dl.radxa.com/q6a/hw/radxa_dragon_q6a_schematic_v1.21.pdf) · [E52C V1.2](https://dl.radxa.com/e/e52c/hw/radxa_e52c_v1.2_schematic.pdf) · [E54C V1.2](https://dl.radxa.com/e/e54c/hw/radxa_e54c_v1.2_schematic.pdf) · [E20C V1.10](https://dl.radxa.com/e/e20c/v1.10/radxa_e20c_v1100_schematic.pdf) · [E24C V1.2](https://dl.radxa.com/e/e24c/docs/radxa_e24c_v1200_schematic.pdf)

**Radxa product briefs & docs**
[ROCK 5 ITX](https://dl.radxa.com/rock5/5itx/radxa_rock5_itx_product_brief.pdf) · [ROCK 5 ITX+](https://dl.radxa.com/rock5/5itx-plus/docs/rad-doc-0078_radxa_rock5_itx_plus_product_brief__revision_1.2_g148a184.pdf) · [ROCK 5A](https://dl.radxa.com/rock5/5a/docs/radxa_rock5a_product_brief.pdf) · [ROCK 5T](https://dl.radxa.com/rock5/5t/docs/radxa_rock5t_product_brief.pdf) · [ROCK 5C](https://dl.radxa.com/rock5/5c/docs/hw/v1100/radxa_rock5c_product_brief.pdf) · [ROCK 4D](https://dl.radxa.com/rock4/4d/docs/radxa_rock4d_product_brief.pdf) · [Orion O6 manual](https://dl.radxa.com/orion/o6/docs/radxa_orion_o6_user_manual.pdf) · [Orion O6N](https://dl.radxa.com/orion/o6n/docs/radxa_orion_o6n_product_brief_en.pdf) · [E52C](https://dl.radxa.com/e/e52c/radxa_e52c_product_brief_Revision_1.0.pdf) · [E54C](https://dl.radxa.com/e/e54c/docs/radxa_e54c_product_brief_Revision_1.0_g37b8f72.pdf) · [E20C](https://dl.radxa.com/e/e20c/radxa_e20c_product_brief_Revision_1.1.pdf) · [ZERO 3E](https://dl.radxa.com/zero3/docs/hw/3e/radxa_zero_3e_product_brief.pdf) · [Cubie A5E](https://dl.radxa.com/cubie/a5e/radxa_cubie_a5e_product_brief.pdf) · [CM5](https://dl.radxa.com/cm5/radxa_cm5_product_brief.pdf) · [NX5](https://dl.radxa.com/nx5/radxa_nx5_product_brief.pdf) · [CM5 FAQ (CM4 compat)](https://docs.radxa.com/en/som/cm/cm5/faq) · [PoE HAT index](https://docs.radxa.com/en/accessories/poe-hat) · [Rock5 serial console](https://wiki.radxa.com/Rock5/dev/serial-console)

**Press & analysis**
[CNX Software](https://www.cnx-software.com/) — [ROCK 5 ITX review](https://www.cnx-software.com/2024/08/18/radxa-rock-5-itx-rk3588-mini-itx-motherboard-review-building-an-arm-pc-and-nas-with-debian-kde/) · [X4 review](https://www.cnx-software.com/2024/09/29/radxa-x4-review-an-intel-n100-alternative-to-raspberry-pi-5-tested-with-ubuntu-24-04/) · [Orion O6 review](https://www.cnx-software.com/2025/01/29/radxa-orion-o6-review-unboxing-debian-12-installation-and-first-benchmarks/) · [E52C](https://www.cnx-software.com/2024/11/08/radxa-e52c-rockchip-rk3582-router-with-dual-2-5gbe-usb-3-0-port-usb-console-port/) · [E24C/E54C](https://www.cnx-software.com/2025/07/15/radxa-e24c-and-e54c-rockchip-rk3528a-rk3582-network-computers-features-four-gigabit-ethernet-ports/) · [2026 Qualcomm roadmap](https://www.cnx-software.com/2026/06/01/radxa-2026-qualcomm-hardware-dragon-q8b-and-q5e-sbcs-dragonstation-and-dragonbay-nas-systems/) · [SBC prices 2024 vs 2026](https://www.cnx-software.com/2026/04/28/what-a-difference-two-years-make-comparing-sbc-prices-in-2024-and-2026/) · [Jeff Geerling on the Orion O6](https://www.jeffgeerling.com/blog/2025/radxa-orion-o6-brings-arm-midrange-pc/) · [Geerling on DeskPi Super4C](https://www.jeffgeerling.com/blog/2026/deskpi-super4c-sbc-cluster/) · [Liliputing](https://liliputing.com/) · [LinuxGizmos](https://linuxgizmos.com/)

**Firmware / kernel / distro**
[Cix edk2-platforms ASPM issue #1 (OPEN)](https://github.com/cixtech/cix_opensource__release__edk2-platforms/issues/1) · [QPS615 dt-binding](https://lkml.iu.edu/hypermail/linux/kernel/2408.0/02688.html) · [OpenWrt rockchip/armv8 builds](https://downloads.openwrt.org/releases/25.12.5/targets/rockchip/armv8/) · [Armbian board pages](https://www.armbian.com/) · [rknn_model_zoo NPU benchmarks](https://github.com/airockchip/rknn_model_zoo) · [Jellyfin Intel HWA matrix](https://jellyfin.org/docs/general/post-install/transcoding/hardware-acceleration/intel/) · [NixOS on ROCK 5 ITX](https://wiki.nixos.org/wiki/NixOS_on_ARM/Radxa_ROCK_5_ITX)

**Forums (measured user data)**
[ROCK 5 ITX idle power](https://forum.radxa.com/t/rock5-itx-relatively-high-idle-power-consumption/25580) · [ATX PSU always on](https://forum.radxa.com/t/atx-psu-always-on/23992) · [E52C availability + power + RTC](https://forum.radxa.com/t/e52c-is-available-now/24511) · [E52C OpenWrt thread (119 posts)](https://forum.openwrt.org/t/radxa-e52c-a-rockchip-rk3582-router-with-dual-2-5gbe-usb-3-0-port-usb-serial-console-port/215258) · [Q6A cold-boot EDL loop](https://forum.radxa.com/t/dragon-q6a-wont-boot-ddr-rcw-training-fails-on-cold-boot-board-loops-into-edl/30841) · [ZERO 3E PoE failures](https://forum.radxa.com/t/unable-to-power-zero-3e-using-poe/22882) · [Auto power-on (unresolved)](https://forum.radxa.com/t/can-the-5b-or-any-of-the-rocks-auto-power-on/23204)

**Retail**
[arace.tech](https://arace.tech/collections/radxa) (pricing source of truth) · [ALLNET China](https://shop.allnetchina.cn/) · [ameriDroid](https://ameridroid.com/) (cross-check only)
