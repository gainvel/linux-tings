# Data Center (appid 4170200) — custom shop & rack items via the native mod system

Companion to [`data-center.md`](data-center.md), which covers the **save format**.
This file covers the **mod system**: how to add new buyable items (including rack-mount
gear) to the in-game shop.

**Headline:** despite being a Unity **IL2CPP** build, Data Center ships an *official,
no-code mod loader*. A mod is a folder with a `config.json`, a Wavefront `.obj`, and two
`.png` files. No BepInEx, no MelonLoader, no assembly patching, no asset bundles.

> **None of this is documented publicly.** There is no Steam guide, no dev announcement,
> no GitHub repo describing it. Everything below was reverse-engineered from
> `global-metadata.dat` and confirmed against a live `[ModLoader]` run in `Player.log`.
> Re-verify after game updates. Recovered against build `24593929`, Unity `6000.4.12f1`.

---

## The two mod layers

| Layer | What it is | Needs | Can do | Workshop tag |
|---|---|---|---|---|
| **Native** (this doc) | Game's own `ModLoader`, reads `config.json` | nothing | Buyable / carryable / placeable items built from OBJ + PNG. **Cosmetic — no network or compute behaviour** | `standalone` |
| **Code** | MelonLoader 0.7.x + Il2CppInterop + Harmony | MelonLoader under Proton | Real gameplay: functional servers, ports, cabling, UI | `mod`, `melonloader` |

Real-world confirmation of the split:

- **"Data Center — Wandregale"** ships native mod folders *plus* an optional MelonLoader
  DLL, and states that without the DLL the shelves are still buyable and placeable —
  *"reine Deko"* (pure decoration).
- **"Backplane Boost Servers"** adds *functional* rack servers with real IOPS and
  SFP/QSFP ports, and requires MelonLoader.

`config.json` has a `dlls` array (`DllEntry { fileName, entryClass }`) and the loader has
`LoadDll` / `IModPlugin` (`modFolderPath`, `OnModLoad`, `OnModUnload`). **Assume this is
dead in the IL2CPP build** — managed `Assembly.Load` has no JIT to run against. It is
almost certainly why the entire community codes against MelonLoader instead. Do not plan
around it without testing it first.

---

## Paths

`STEAMAPPS` = `~/.local/share/Steam/steamapps`
`PFX` = `STEAMAPPS/compatdata/4170200/pfx/drive_c/users/steamuser`

| What | Path |
|---|---|
| Install | `STEAMAPPS/common/Data Center/` |
| Data dir | `STEAMAPPS/common/Data Center/Data Center_Data/` |
| **Live mods dir** | `…/Data Center_Data/StreamingAssets/Mods/` |
| Workshop staging | `STEAMAPPS/workshop/content/4170200/<publishedFileId>/` |
| IL2CPP metadata | `…/Data Center_Data/il2cpp_data/Metadata/global-metadata.dat` |
| Rack/server assets | `…/Data Center_Data/sharedassets1.assets` + `.resS` |
| **Log (only feedback channel)** | `PFX/AppData/LocalLow/WASEKU/Data Center/Player.log` |
| Saves | `PFX/AppData/LocalLow/WASEKU/Data Center/saves/` |
| Mod source of truth | `~/claude/modding/work/data-center/mods/` |

The live mods dir is inside the install dir, so **Steam updates and *Verify integrity of
game files* delete it**. Always rebuild from `work/`.

---

## Loader flow

`Assets/Scripts/Mods/ModLoader.cs`:

```
SyncWorkshopThenLoadAll()          coroutine, runs on entering a game
  ├─ enumerate subscribed Steam UGC items for 4170200
  ├─ SteamUGC.GetItemInstallInfo  ->  STEAMAPPS/workshop/content/4170200/<id>/
  ├─ CopyDirectory(sourceDir, destDir)
  │     dest = <StreamingAssets>/Mods/workshop_<publishedFileId>/
  │     writes .workshop_timestamp to detect staleness
  ├─ waits for downloads (has a timeout)
  └─ LoadAllMods()
        └─ for each subdirectory of <StreamingAssets>/Mods/:
             LoadModPack(folderPath)
               ├─ read config.json  ->  ModPackConfig  (Newtonsoft.Json)
               ├─ for each shopItems[]   -> LoadShopItem
               │     ├─ CreateShopTemplate(folderName, modID)
               │     │     ├─ LoadMesh(modelFile)      via ObjImporter.ImportOBJ
               │     │     ├─ CreateMaterial(textureFile)
               │     │     └─ LoadIcon(iconFile)       via LoadTexture
               │     └─ CreateShopButton               instantiates modShopButtonPrefab
               │                                       under modShopItemsParent
               ├─ for each staticItems[] -> LoadStaticItem -> CreateStaticInstance
               └─ for each dlls[]        -> LoadDll        (see caveat above)
```

Loader state: `modTemplates`, `modTemplatesByFolder`, `staticInstances`, `loadedPlugins`,
`nextModID`. Lookup helpers `GetModPrefab` / `GetModPrefabByFolder` exist — these are the
handles a MelonLoader plugin would use to attach behaviour to a native item.

### Log lines to grep for

```
[ModLoader] Loading mod pack: <path>            ok
[ModLoader] Shop item loaded: <name>            ok
[ModLoader] Static item loaded: <name>          ok
[ModLoader] No config.json found in <path>      folder ignored
[ModLoader] Failed to parse config in <path>    JSON shape wrong
[ModLoader] Failed to load model: <file>        OBJ rejected
[ModLoader] Workshop item {0} synced to {1}
[ModLoader] Workshop item {0} already up-to-date.
[ModLoader] Timed out waiting for Workshop downloads.
[ModShopItem] Attempting to buy mod item: {0} for {1} $
```

---

## `config.json` schema

The JSON keys are the C# field names verbatim (Newtonsoft.Json, no custom contract
resolver observed).

```
ModPackConfig     { modName, shopItems[], staticItems[], dlls[] }

ShopItemConfig    { itemName, price, xpToUnlock, sizeInU, mass, modelScale,
                    colliderSize, colliderCenter, modelFile, textureFile,
                    iconFile, objectType }

StaticItemConfig  { itemName, modelScale, colliderSize, colliderCenter,
                    modelFile, textureFile, position, rotation, isKinematic }

DllEntry          { fileName, entryClass }
```

Types: `price` / `xpToUnlock` / `sizeInU` are `int`; `mass` / `modelScale` are `float`;
`isKinematic` is `bool`; `colliderSize`, `colliderCenter`, `position`, `rotation` are
Vector3-shaped (`{"x":…,"y":…,"z":…}`).

`staticItems` place a prop into the world automatically at a fixed `position`/`rotation`
with no shop entry — that is how the dev's own "Waseku Poster" works.

### Example

```json
{
  "modName": "My Rack Gear",
  "shopItems": [
    {
      "itemName": "PWM Fan Hub 1U",
      "price": 250,
      "xpToUnlock": 0,
      "sizeInU": 1,
      "mass": 4.0,
      "modelScale": 1.0,
      "colliderSize":   { "x": 0.4826, "y": 0.04445, "z": 0.6 },
      "colliderCenter": { "x": 0.0,    "y": 0.0,     "z": 0.0 },
      "modelFile":   "fanhub.obj",
      "textureFile": "fanhub.png",
      "iconFile":    "fanhub_icon.png",
      "objectType":  "ModItem"
    }
  ],
  "staticItems": [],
  "dlls": []
}
```

### `objectType` enum

Shared by `ShopItemConfig.objectType`, `Item.itemType` and the player's
`objectInHand` / `objectInHandType`. Newtonsoft accepts the name **or** the integer.

| Value | Name |
|---:|---|
| 0 | `None` |
| 1 | `Server1U` |
| 2 | `Server2U` |
| 3 | `Server3U` |
| 4 | `Switch` |
| 5 | `Rack` |
| 6 | `CableSpinner` |
| 7 | `PatchPanel` |
| 8 | `SFPModule` |
| 9 | `SFPBox` |
| **10** | **`ModItem`** ← the dedicated mod-item type; start here |
| 11 | `Router` |
| 12 | `Firewall` |

**Open question:** whether declaring `Server1U` / `Server3U` on a *mod* item grants real
rack-slot semantics. `RackPosition` carries typed slots (`server`, `networkSwitch`,
`patchPanel`) alongside a generic `uo` (`UsableObject`), and `IsAllowedItem` /
`InsertItemInRack` gate insertion. `ShopItemConfig` having a `sizeInU` field at all is
strong evidence rack mounting is intended for native mod items — but it is unproven.
Test it; record the answer in the [Findings log](#findings-log).

---

## Model and texture constraints

`Assets/Scripts/Mods/ObjImporter.cs` is a hand-rolled Wavefront OBJ parser:
`ImportOBJ`, `ProcessFaceVertex`, `ParseFloat`, producing `outVerts`, `outNorms`, `outUVs`.

- **Positions, normals and UVs are supported.**
- **No `mtllib` / `usemtl` handling** — no material library is parsed. One mesh, one
  material, one texture per item. Bake everything into a single albedo.
- **Triangulate on export.** An n-gon path was not observed; do not rely on it.
- Textures load through `LoadTexture` (Unity `ImageConversion`) — use **PNG**.
- Stock rack-item albedos are **512×512**; match that unless there's a reason not to.
- `iconFile` is the shop-card image. Note the stock cards show a *live 3D render* plus an
  EOL timer; mod cards use a flat icon with name and price only (`ModShopItem` has
  `itemIcon`, `txtName`, `txtPrice` — and no EOL field), so mod items will look slightly
  different from the stock ones in the shop grid.

### Scale

Model to **real 19" rack dimensions**, assuming 1 Unity unit = 1 metre:

| Dimension | Real | Unity units |
|---|---|---|
| 1U height | 44.45 mm | `0.04445` |
| Rack rail width | 482.6 mm (19") | `0.4826` |
| Typical body depth | 600–800 mm | `0.6`–`0.8` |

Then correct with `modelScale` after eyeballing the first item next to a stock 1U server.
This avoids needing an asset extractor. `colliderSize` is a `BoxCollider` size in the same
units; `colliderCenter` offsets it.

---

## Reference assets (existing rack models)

Everything rack-related lives in **`Data Center_Data/sharedassets1.assets`** (77.5 MB),
with mesh and texture bytes streamed from **`sharedassets1.assets.resS`** (657 MB).

Ruled out by inspection: `resources.assets` (UMA characters + localisation strings only —
its 261 "rack" hits are German/Dutch/Polish UI text), `globalgamemanagers.assets`
(MonoScript type table), `sharedassets0.assets` / `level0` (main menu), `level1` (the
`BaseScene`, which *instances* prefabs but contains zero Mesh objects),
`StreamingAssets/aa/**` (253 Addressables bundles, **100 % UMA character overlays** — zero
rack keywords in `catalog.bin`), and `StreamingAssets/EntityScenes` + `ContentArchives`
(one 85 KB DOTS subscene, `PacketsSubScene`, used only for the network-packet
visualisation). **The world, racks and rack items are plain GameObject/MonoBehaviour.**

### The rack

Prefab `Rack_Lanberg_47U` — `[Transform, AudioSource, MonoBehaviour:Rack, LODGroup]`

```
Rack_Lanberg_47U
└─ Rack_lanberg_47U        mesh '8847 Lanberg'                  <- LOD0 is the root itself
   ├─ FrontDoor            mesh 'FrontDoor'  [BoxCollider, RackDoor, Outlinable]
   ├─ Inside               mesh 'Inside'
   ├─ InsideDetails        mesh 'InsideDetails'
   ├─ RackPosition.054 … RackPosition.100   (47x, 768 B each)
   │                       [MeshFilter, MeshRenderer, BoxCollider, RackPosition, Outlinable]
   ├─ LOD1                 mesh '8847 Lanberg'
   ├─ LOD2                 mesh '8847 Lanberg'
   └─ LOD99                mesh 'LOD99_rack'   768 B box impostor
```

LOD convention is **child GameObjects literally named `LOD1` / `LOD2` / `LOD99`**, with
LOD0 being the root's own MeshFilter — *not* `_LOD0` suffixes.

Materials: `BrushedAluminiumRack`, `BoxedRack`, `RackDisolve`, `Transparent Rack Door`,
`DecalMaterial_RackShippingLabel`, `DecalMaterial_Landberg logo`.
Textures: `lanberg` (512), `ShippingLabelRack` (1024), `BoxedRack_Bake1_PBR_Diffuse` /
`_Normal` / `_PackedTex` (1024), `Tutorial_Rack`.

### Rack-mounted items

| Prefab | Mesh | Texture | Script |
|---|---|---|---|
| `Server.Blue1` | `1U.004` | `Server_Blue1` | `Server` |
| `Server.Blue2` | `1U.007` | `Server_Blue2` | `Server` |
| `Server.Purple1` | `1U.Base.002` | `Server_purple1` | `Server` |
| `Server.Purple2` | `1U.Base.004` | `Server_pourple2` *(sic)* | `Server` |
| `Server.Yellow1` | `3U.Yellow` | `Server_Yellow1` | `Server` |
| `Server.Yellow2` | `7U.Yellow` | `Server_Yellow2` | `Server` |
| `Server.Green1` | `GPU Server 3U` | `Server_Green1` | `Server` |
| `Server.Green2` | `GPU Server 7U.001` | `Server_Green2` | `Server` |
| `Switch16CU` | `1U 16x RJ45 switch.002` | `NetworkSwitch16_rj45` | `NetworkSwitch` |
| `Switch4xSFP` | `1U 4x SFP switch` | `NetworkSwitch 4xSFP` | `NetworkSwitch` |
| `Switch32xQSFP` | `1U 32x QSFP switch.001` | `NetworkSwitch 32xQSFP` | `NetworkSwitch` |
| `Switch4xQSXP16xSFP` | `1U 16x SFP 4x QSFP switch` | `NetworkSwitch 4xQSFP16xSFP` | `NetworkSwitch` |
| `Router4xQSXP16xSFP 1` | *(same mesh)* | `Router 4xQSFP16xSFP` | `Router` |
| `Firewall4xQSXP16xSFP` | *(same mesh)* | `Firewall 4xQSFP16xSFP` | `Firewall` |
| `PatchPanel` | `PatchPanel.001` | `PatchPanel` | `PatchPanel` |
| `PatchPanel_fiber` | `PatchPanel.Fiber` | `PatchPanel_Fiber` | `PatchPanel` |
| `PatchPanel_combo` | `PatchPanel.RJ45Fiber` | `PatchPanel_Combo` | `PatchPanel` |

**The switch, router and firewall all share one mesh and differ only by texture.** That is
the cheapest possible recipe for a convincing custom 1U item: one generic 1U chassis mesh,
several albedos.

Shape of a stock rack item (for reference — a *native* mod item gets nothing like this
component stack, only mesh + material + collider + rigidbody):

```
Switch4xQSXP16xSFP  [MeshFilter, MeshRenderer, BoxCollider, Rigidbody, LODGroup,
                     Outlinable, TargetStateListener, NetworkSwitch]
├─ ButtonEdit.014 / ButtonPower.014
├─ SFP_Slot1.001 … SFP_Slot4.016  (16x)  [CableLink]
├─ QSFP_port.01 … .04             (4x)   [CableLink]
└─ LOD1
```

`CableLink` is the port abstraction. Ports are child GameObjects.

### Stock shop definitions

ScriptableObjects `ShopItemSO_*` in the same file, 29 of them, present nowhere else:
`ShopItemSO_Rack47U`, `…_Rack47U custom color`, `…_Server_{Blue1,Blue2,Green1,Green2,Purple1,Purple2,Yellow1,Yellow2}`,
`…_Switch_{16RJ,4SFP,32QSFP,4QSFP_16SFP}`, `…_Router_4QSFP_16SFP 1`,
`…_Firewall_4QSFP_16SFP 2`, `…_PatchPanel_{18RJ,18 Fiber,18 Combo}`,
`…_SFP_{Fiber10,Fiber25,Fiber40,RJ45}`, `…_Cable_*` (8).

`ShopItemSO` fields: `itemDisplayName`, `itemName`, `sprite`, `xpToUnlock`, `price`,
`itemType`, `itemID`, `eol`.

### Extracting them

Only needed if you want the actual geometry as a modelling base. **No extractor is
installed** — ask before installing (workspace rule 6).

- **AssetRipper** — best choice. Serialized files are v22 with **type trees stripped**
  (`enableTypeTree = 0`) on Unity 6000.4.12f1, so the tool must ship a class DB for that
  exact version. Reconstructs prefabs, LODGroups and MonoBehaviour fields (needed to read
  `ShopItemSO_*`).
- **AssetStudio** — pulls raw Mesh/Texture2D/Material and resolves `.resS` automatically,
  but will not rebuild prefab hierarchies.
- **UnityPy** (`pip install UnityPy`) — scriptable, fine for meshes/textures; needs a
  TypeTreeGenerator dump from `GameAssembly.dll` + `global-metadata.dat` for
  `ShopItemSO_*` field values.

222 of 256 meshes stream from `.resS` (~49 MB of vertex/index data at offsets ~599–649 MB).

---

## Persistence

Saves are .NET `BinaryFormatter` (MS-NRBF), `SaveData.version = 8` — see
[`data-center.md`](data-center.md) for the full layout. Two members matter here:

- `SaveData.modItemData` — `List<ModItemSaveData>`, where
  `ModItemSaveData = { modFolderName, saveValue, saveIntArray, saveIntArray2 }`
- `SaveData.shopItemUnlockStates` — `Dictionary<string,bool>`. **Stock** items are keyed
  by GUID (31 entries; `ShopItem` has a `guid` field). How mod items key into unlock state
  is unconfirmed — `ModShopItem` carries `modID` and a `config` reference rather than a
  guid.

### The folder-name trap

`ModItemSaveData.modFolderName` is the key that ties a placed item back to its mod. The
Workshop copy lands in `Mods/workshop_<publishedFileId>/`, so once a mod is published its
folder name **changes** from whatever you used locally. Every instance you placed while
testing under `mods/dc-testitem/` is orphaned in that save. Test on a throwaway save, and
never rename a mod folder after release.

---

## Local test loop

```bash
MODS="$HOME/.local/share/Steam/steamapps/common/Data Center/Data Center_Data/StreamingAssets/Mods"
PFX="$HOME/.local/share/Steam/steamapps/compatdata/4170200/pfx/drive_c/users/steamuser"
LOG="$PFX/AppData/LocalLow/WASEKU/Data Center/Player.log"

~/claude/modding/work/data-center/install-testmod.sh    # work/ -> live Mods/
steam -applaunch 4170200                                 # load a save, check the shop, quit
grep -E '\[ModLoader\]|\[ModShopItem\]' "$LOG"
```

The loader runs on **entering a game**, not at the main menu — you must load or start a
save for mods to load. `Player.log` is truncated each launch; the previous run survives as
`Player-prev.log`.

---

## Publishing to the Workshop (deferred)

The game **subscribes and downloads** UGC but has **no in-game uploader** — there are no
`CreateItem` / `SubmitItemUpdate` call sites, only `GetSubscribedItems`,
`GetItemInstallInfo` and `DownloadItem`. Publishing therefore goes through
**SteamCMD `+workshop_build_item <item.vdf>`** with `appid 4170200`.

Existing Workshop items confirm the tag convention: **`standalone`** for native
(no-loader) mods, **`mod` / `melonloader` / `qol`** for code mods.

---

## Verified vs unverified

**Verified** (observed in `Player.log` on this machine, or read directly out of metadata):
loader exists and runs; workshop sync path and `workshop_<id>` naming; `.workshop_timestamp`;
`config.json` is required per folder; full field lists of `ModPackConfig` / `ShopItemConfig` /
`StaticItemConfig` / `DllEntry` / `ModShopItem` / `ModItemSaveData`; the `objectType` enum
and its ordering; OBJ importer handles verts/normals/UVs and has no MTL support; asset
locations; save layout.

**Unverified — settle by testing:**

1. Does a folder with an **arbitrary name** (not `workshop_*`) load? (Strongly implied —
   "Wandregale" ships three arbitrarily-named folders — but untested here.)
2. Are `colliderSize` / `colliderCenter` really Vector3-shaped in JSON? The IL2CPP type
   index table has aliased entries that made static confirmation unreliable. A wrong shape
   produces `[ModLoader] Failed to parse config in`.
3. Does `objectType: Server1U` (or any non-`ModItem` value) let the item mount in a rack
   slot? Does `sizeInU` reserve multiple U?
4. Does `xpToUnlock > 0` actually gate a mod item, and how is its unlock state keyed?
5. Does the loader ever **delete** unknown folders from `Mods/` during workshop sync?
6. Is `dlls` genuinely dead under IL2CPP?

---

## How the API was recovered (reproducible)

`global-metadata.dat` is unencrypted; strings are plaintext. Metadata version `39`,
magic `AF1B B1FA`. Header is a run of `(offset, size, count)` triples from byte `0x08`.

| Region | Offset | Size |
|---|---|---|
| String literal blob | `0x176BC` | `0xB91D9` |
| **Identifier string blob** | `0xD0898` | `0x2A6C73` |
| **Field table** (`{nameIndex, typeIndex, token}`, 12 B/entry) | `0xBCCF74` | `0xFB6B8` |

Method: find an identifier's offset in the blob, subtract `0xD0898` to get its
`nameIndex`, byte-search the file for that `uint32`, and confirm you've hit the field
table by checking the next entry 12 bytes later is the expected sibling field. Then walk
backwards and forwards — a class's fields are contiguous and in declaration order.

```bash
DC="$HOME/.local/share/Steam/steamapps/common/Data Center/Data Center_Data"
strings -n 4 "$DC/il2cpp_data/Metadata/global-metadata.dat" | grep -F 'Assets\Scripts\Mods'
# \Assets\Scripts\Mods\IModPlugin.cs   ModItemSaveData.cs   ModLoader.cs
#                      ModPackConfig.cs  ModShopItem.cs      ObjImporter.cs
```

Caveat: `typeIndex` values are **aliased** — the same type can have several indices, and
cross-referencing indices between unrelated fields gave contradictory results. Field
*names* and *ordering* are reliable; inferred field *types* are not, which is why item 2
in the unverified list exists.

---

## Gotchas

- **Steam update or *Verify integrity* wipes `StreamingAssets/Mods/`.** Rebuild from `work/`.
- **Nothing in `~/claude/modding` is backed up** — the nightly rsync skips it. Copy
  anything irreplaceable to `~/Backups/Manual/`.
- **Back up saves before testing.** Mod items are written into `modItemData`; a bad
  round-trip can break a load. Steam Cloud is **enabled** for this appid and will revert
  local save edits — disable it for 4170200 before touching saves.
- **Mods load on entering a game**, not at the menu.
- **Never rename a released mod folder** (see the folder-name trap).
- The recovered API is **version-specific**. Re-run the recovery after each game patch.

---

## Findings log

*(append test results here — dates, what was tried, what the log said)*

**2026-08-14** — API recovered from metadata; no in-game testing performed yet.
Subscribed Workshop item present: `workshop_3775738163` (Cart Item Stacker, a MelonLoader
DLL mod) — it has no `config.json`, so every launch logs
`[ModLoader] No config.json found in …\Mods\workshop_3775738163`. That warning is expected
and harmless. MelonLoader is **not** installed on this machine, so that DLL is inert.
