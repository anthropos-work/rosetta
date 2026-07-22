# Content-Stories Spec — the content_products manifest + honesty gate + the cockpit render

**The M233 + M234 deliverable (v2.5 "the playbill", Thread B — the manifest half + the render half).** Where
[`session-clone-spec.md`](session-clone-spec.md) (M232) COPIES real production sessions into a demo org, and
[`content-stories-routes.md`](content-stories-routes.md) (M231) proved each result page reads a seedable row,
THIS doc specifies **(M233)** the **manifest** that turns those seeded sessions into a cockpit menu — a
**`content_products[]` projection** (§1–§6) — and **(M234)** the **cockpit render** that turns that manifest
into the 2nd "Content stories" tab + the `content-player-<idx>` seat registration that makes its as-player
CTAs log in (§7). The manifest is a projection from the same fixture the seeder seeds from, honesty-gated so
it can never drift, and fail-closed so it never fabricates a link.

> **v2.6 "sound check" M241 — the EN/IT language axis.** Each session now carries the REAL language it was
> played in (`language` ∈ english|italian), and the cockpit gains an **EN/IT toggle** that swaps which
> language's session each requirement cell logs-in-and-lands on. The pre-M241 seeder hard-coded every clone to
> `english` even though 11 of the 13 pinned sessions were actually Italian; §2's `language` field fixes that
> (each session is consumed in its intended language). The `language`/`lang_toggle` schema is §2, the
> fail-closed language honesty gate is §4, the cockpit toggle is §7.6, and the per-tuple coverage (which cells
> got both languages) is set by a read-only prod pool query — see `session-clone-spec.md` §2.1.

> **Headline — the content menu is a PROJECTION, not hand-authored, and it can't drift.** The cockpit's
> "Content stories" tab reads a `content-manifest.json` that `stackseed --content-export` PROJECTS from the
> exact content-session fixture the `ContentStorySeeder` seeds from + the blueprint that re-tenants it. So the
> projected player seat OWNS the seeded session, and the projected result path names the seeded session id —
> single-sourced by construction (D9). A checked-in canonical + a `CanonicalFileMatchesProjection`-style
> honesty gate keep the checked-in file equal to the live projection, and a **fail-closed resolver** drops
> (never fabricates) any exhibit that can't form a real link.

## For PMs — one paragraph

A "content story" is a real, played session a presenter can log into and see the result of. To drive that, the
presenter cockpit needs a menu: for each kind of content (a simulation, a skill path, an academy course), the
list of sessions, and for each session two buttons — *open this as the employee who did it* and *open this as
their manager*. This doc is the **menu**. It's not written by hand: a tool projects it straight from the same
data the seeder plants, so the menu can never promise a session that wasn't seeded, or a button that leads
nowhere. If any exhibit can't be turned into a real, working link, the tool refuses it loudly rather than ship
a dead button.

---

## 1. Where it sits — a peer manifest, JSON, honesty-gated

The demo already emits two manifests (both single-sourced from the blueprint, the D9 property):

| File | Role | Emitted by |
|---|---|---|
| `cockpit-manifest.json` (the stories MENU) | the org-stories → heroes seat-switch menu (the first cockpit tab) | `stackseed --cockpit-export` |
| `seed-generation-manifest.yaml` (the seed+gen INTENT) | the auditable seed direction — population + generation + the M232 **`content_sessions`** source-pins | `stackseed --manifest-export` |

M233 adds a **third**, the content analog of the stories menu:

| File | Role | Emitted by |
|---|---|---|
| **`content-manifest.json`** (the content MENU) | the `content_products[]` projection — the second "Content stories" cockpit tab (M234) | **`stackseed --content-export`** |

**Why a separate JSON, not a block in `seed-generation-manifest.yaml`** (the M233 open question, resolved —
`#D-M233-1`). The presenter cockpit is **stdlib-only Python: it reads JSON, never YAML** (no PyYAML — the
supply-chain-GREEN posture, [`cockpit.go`](../../../.agentspace/rosetta-extensions/stack-seeding/seeders/cockpit.go)).
So the RENDER surface a tab reads MUST be JSON — exactly like `cockpit-manifest.json`. `content_products[]` is
therefore a **peer manifest artifact** (the content analog of the stories manifest), honesty-gated by its own
checked-in canonical `presets/content-manifest.json` + a `CanonicalFileMatchesProjection`-style test. The M232
**source-pins stay folded in `seed-generation-manifest.yaml`'s `content_sessions` block** (the audit
disclosure — §5). This keeps the **non-fatal bring-up**: a broken content projection drops the tab, never
blocks the cockpit.

## 2. The `content_products[]` schema

`content-manifest.json` (types in `seeders/content_manifest.go`, dual json+yaml tagged):

```jsonc
{
  "stack": "demo-1",
  "products": [
    {
      "id": "simulation",            // simulation | skill-path-legacy | skill-path-new | ai-labs
      "name": "Simulation",          // the section header
      "app_base": "web",             // web | hiring | academy — the origin the CTAs join to
      "icon_key": "flask",           // the section FontAwesome free-solid key (no "fa-" prefix)
      "sessions": [
        {
          "key": "asmt-voice-pass",
          "source_session_id": "…",  // the prod pin, folded into the render projection (auditable)
          "sim_id": "…", "sim_type": "SIMULATION_TYPE_ASSESSMENT",
          "modality": "voice", "passed": true,
          "language": "italian",                 // M241: the REAL language the session was played in (english|italian)
          "lang_toggle": true,                   // M241: true IFF this cell's tuple has BOTH languages (omitempty; solo→absent)
          "icon_key": "clipboard-check",         // per-sim_type row icon
          "player_seat": "content-player-23",    // the owner-member seat key (M234 registers it)
          "player_result_path": "/sim/<slug>/result/<sessionId>",
          "has_manager_view": true,
          "manager_seat": "dan-manager",         // the host org's manager hero (omitted if no manager view)
          "manager_result_path": "/enterprise/activity-dashboard/ai-simulations/<simId>/<userId>"
        }
      ]
    }
  ]
}
```

### The per-product registry (schema-complete; the fixture drives what's projected)

`contentProductRegistry()` covers all four products the feature spans. The **simulation** section projects
from the `contentsession` fixture (M232–M234); the three **non-simulation** sections (skill-path-legacy /
skill-path-new academy / ai-labs) project from a **separate code-owned registry** `nonSimExhibits()`
(`seeders/content_nonsim.go`, M235) appended by `BuildContentProducts` — so the simulation fixture + seeder
stay simulation-shaped and untouched. An empty product section is never fabricated.

| product id | app_base | section icon | player link? | manager surface |
|---|---|---|---|---|
| `simulation` | `web` | `flask` | yes — `/sim/<slug>/result/<sessionId>` | `ai-simulations` (or `interviews` for an interview sim_type) |
| `skill-path-legacy` | `web` | `diagram-project` | yes (M234) | **— none** (the platform surface is unimplemented — M236 iter-07) |
| `skill-path-new` (academy) | `academy` | `graduation-cap` | yes — `/courses/<slug>` (**M236 iter-08**) | — (no academy manager route, M231) |
| `ai-labs` | `web` | `vials` | **no — presence-only** (M231 §5) | — |

> **Two entries in this table were corrected at M236 by driving the routes live; both had been asserted
> offline and defended by green unit tests.**
>
> - **`skill-path-legacy` has no manager surface.** `managerKind` was `skill-paths` through M235. The route
>   renders the literal string **"Coming soon"**: next-web's `InsightsBySkillPathStudentSimulationsContainer`
>   hardcodes `userData = null` and its results table is **commented out**, so no query touches the seeded
>   session. Projecting a CTA there is a fabricated CTA, which §"fail-closed" forbids. Restore `managerKind`
>   the day the platform builds the surface — nothing else needs to change.
> - **The academy player route is `/courses/<slug>`, not `/library/<slug>`.** ant-academy has **no
>   `/library/[slug]` route at all** (only `app/(public)/library/page.jsx`, the index), and the slug M235
>   pinned was not in its catalog either — so that CTA 404'd on every visit. The slug must come from the
>   catalog the **demo** academy serves: its committed **FS** content, because a demo academy runs with
>   `ACADEMY_DEMO_FS_PUBLISHED=1` and **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, so `getBackendCatalogView`
>   always returns null.**
>
> That same missing endpoint is why **`app/cmd/academy-seed` is moot in a demo**: with no backend
> connection, academy progress rows written to the demo DB have nothing that can read them. A
> progress-bearing academy CTA requires wiring the academy to the demo's GraphQL first.

Per-**sim_type** row icons: ASSESSMENT `clipboard-check` · TRAINING `dumbbell` · HIRING `user-tie` ·
INTERVIEW `comments`.

### The language axis (M241 — v2.6 "sound check")

Every SIMULATION row carries two language fields (the non-simulation exhibits carry neither — they are not
real played sessions):

- **`language`** ∈ `english | italian` — the REAL language the source session was played in (the source's
  `jobsimulation.sessions.language`, copied onto the pin and written to the clone's session row so the demo
  renders the story in its intended language). The fixture `Validate` requires a valid language on every pin
  (a missing/unknown one fails loud at load — the seeder never falls back to a wrong label). Before M241 the
  seeder hard-coded every clone to `english`, mislabeling the 11-of-13 Italian pins.
- **`lang_toggle`** (`omitempty`; a missing value = solo/not-toggle-able) — **true IFF the row's requirement
  tuple `(sim_type × modality × pass/fail)` has BOTH an english AND an italian variant in the fixture**, so the
  cockpit's EN/IT switch can actually swap this cell between languages. It is **DERIVED** (`bilingualTuples`
  over the fixture), never hand-authored, so it cannot lie about coverage. A **single-language** cell projects
  `lang_toggle=false`: its row always renders and the toggle is not offered for it (the user-decision
  "toggle hidden/disabled for that tuple" fallback). The one single-language product in the shipped fixture is
  **INTERVIEW** — both interview cells are Italian-only (no believable English interview session exists; the
  pool query found EN interview passes all out-of-band and EN interview fails = 0 — release risk R2).

**Coverage (shipped fixture):** 23 simulation sessions = the 13 base pins (11 italian + 2 english, now
correctly labeled) + 10 EN/IT counterparts, so 11 of the 12 requirement tuples carry both languages; only
INTERVIEW stays Italian-only. The landable-pair denominator moved **29 → 49** accordingly
(`stack-verify/e2e/content-denominator.json`).

## 3. The seat + route model (single-sourced with the seeder)

- **player_seat** — the content-story player is a **non-hero MEMBER** (owner-is-player-vantage, M232). The
  seat key convention is **`content-player-<memberIndex>`**, resolved from the SAME owner derivation the
  seeder uses (`eligiblePlayerOwnerSlots` + the flat-index `slots[idx % len]` assignment), so the projected
  seat OWNS the exact seeded session. M234 REGISTERS these seats in the roster + Clerkenstein (the roster
  today carries only heroes — the non-hero player seats are M234's `roster.go` extension).
  - **The flat index advances through DROPPED sessions** (it is incremented *before* the known-product / drop
    check), so a session in a later product keeps the exact owner the seeder assigned it — the seeder iterates
    the same flattened `Set.Sessions()` and seeds **every** session regardless of product, so the projection's
    index and the seeder's index only agree if drops still consume their slot. A per-product index reset here
    would silently re-own every session after the first drop → dead CTAs. This survives-drops invariant is
    pinned by `TestContentProducts_FlatIndexSurvivesDrops` (M233 harden) — **preserve it when M234 adds
    non-simulation products to the fixture.**
- **manager_seat** — the **host org's manager-vantage hero** (e.g. `dan-manager`), already a roster seat.
  `has_manager_view` follows the M231 matrix, **downgraded to false (fail-closed)** if the org has no manager
  hero — never a promise without a seat.
- **player_result_path** — `/sim/<slug>/result/<sessionId>`. The `[slug]` segment resolves by
  `jobSimulationBySlug` (a **text slug, not the sim uuid**), so the fixture carries a per-session **`sim_slug`**
  (the public sim's slug, resolved read-only from the public catalog at authoring time — public + non-PII;
  `#D-M233-3`). `<sessionId>` is the seeder's own derived id (`contentStorySessionID`), so the link names the
  seeded row.
- **manager_result_path** — `/enterprise/activity-dashboard/<kind>/<simId>/<membershipId>` where `<kind>` ∈
  {`ai-simulations`, `interviews`} (**`skill-paths` was removed at M236 iter-07** — see the table above),
  `<simId>` is the sim uuid. Fully offline-derivable (no slug).

  > **The last segment is a MEMBERSHIP id, not a user id** (M236 iter-05). The page calls
  > `GetMembership(membershipsID)`; hand it a user id and it returns `ent: membership not found` and the
  > whole query **nulls**. It fails *silently*, because the page header is served by a **different** query
  > and renders fine either way — so the scoreboard looks populated while proving nothing. If a manager
  > pair "passes" but you have not seen its **table rows**, assume this bug.

**app_base is `web`, not hiring** (`#D-M233-2`). Content-story sessions are re-tenanted into a **Workforce**
org (`firstNonHiringStory`), so they render in apps/web regardless of the source sim's `sim_type` — M231's
"HIRING → apps/hiring" is the org-ejection rule for genuinely-hiring ORGS (M224), a different feature. The
`academy` app_base is reserved for the future academy product (the offset :3077 origin).

## 4. The honesty gate + the fail-closed resolver

- **Honesty gate (D9).** `TestContentManifest_CanonicalFileMatchesProjection` re-projects from
  `presets/stories.seed.yaml` and asserts the checked-in `presets/content-manifest.json` still equals it — so
  the checked-in menu can never silently drift from the fixture / slugs / seat derivation. A `HasTeeth`
  meta-test mutates a projected field (the manager-hero id) and asserts the gate diverges + carries the
  mutation — proving the byte comparison bites. Regenerate on any intended change:

  ```bash
  stackseed --content-export --seed presets/stories.seed.yaml --content-out presets/content-manifest.json
  ```

- **Fail-closed / no-fabrication (`#D-M233-4`).** `BuildContentProducts` **DROPS** (with a recorded reason,
  never a fabricated link) any session that can't form a real link — no eligible player owner, a missing/blank
  `sim_slug` for a simulation exhibit, an unknown product id, or an unknown sim_type. `ValidateContentManifest`
  (the `--content-export` guard, the analog of `ValidateCockpitManifest`) then **FAILS LOUD** naming any
  dropped exhibit — "a refusal nobody sees never happened" (the D17 / cockpit-guard philosophy). A
  **presence-only** product (AI-labs — no seedable result page) is PROJECTED without a player link, a
  legitimate disposition, not a drop.

- **Empty collections marshal as `[]`, never `null` (the M217 wire-format contract).** The projection
  initializes `products` (and every product's `sessions`) as an empty slice, so an honest empty section
  marshals `"products": []` — the stdlib-Python cockpit reads `products` via `dict.get("products", [])`, and a
  JSON `null` there is `len(None)` → the M217 cockpit-crash class (exit 1). Pinned by
  `TestContentProducts_EmptyMarshalsProductsAsArray` (M233 harden). (`WriteContentManifest` also fails-closed on
  any drop, so an emitted file is never silently holed — but the array-not-null contract is what keeps a valid
  empty projection safe for the JSON consumer.)

- **The route prefixes are a CROSS-LANGUAGE contract, and they are pinned (M236 final harden).** The
  projection emits `player_result_path` / `manager_result_path` as plain strings; the sweep's grader
  (`shapeFor`, in `stack-verify/e2e/lib/content-result-page.ts`) picks which render shape to assert by
  matching **prefixes of those strings**. Nothing else joins them — different language, different section,
  no shared constant. Worse, `shapeFor` **falls through to `player-scored`** for any prefix it does not
  recognise, so renaming `/courses/` here throws nothing, fails no Go test and no TS test, and merely grades
  every academy page against a scored-report shape — reporting a correct render as a failure. That *is* the
  M236 iter-08 defect, and after iter-08 nothing prevented its return: four iters changed this side, four
  changed the grader, and no test covered the join.
  `stack-verify/e2e/tests/content-route-contract.unit.spec.ts` closes it by reading **this checked-in
  canonical manifest** and asserting the grader understands every route in it — per-product expected shape,
  no unexplained fall-through, interview vs simulation manager surfaces kept distinct, manager paths still
  uuid-terminated — and that the projection still yields the pinned landable-pair count (read from
  `content-denominator.json` — **49** since M241's EN/IT counterparts grew the fixture 13 → 23; it was 29 at
  M236). Mutation-verified: a one-character prefix change turns it red. **If you change a route in
  `contentProductRegistry`, expect that spec to fail — updating it is part of the change, not collateral.**

- **The LANGUAGE axis is fail-closed on BOTH sides (M241).** `ValidateLanguageConsistency`
  (`content_manifest.go`, wired into `WriteContentManifest`) refuses to export a manifest whose `lang_toggle`
  disagrees with its own coverage: a **solo cell marked toggle-able** (the cockpit would offer a switch that
  swaps to nothing) or a **bilingual cell marked solo** (a language silently un-reachable) or an **invalid
  language** all FAIL `--content-export`. A `HasTeeth` test mutates each case and asserts the gate bites. The
  TS side mirrors it over the checked-in canonical (`stack-verify/e2e/tests/content-language.unit.spec.ts`):
  valid language label per row, the toggle-able set spans both languages, the same `lang_toggle ⟺ coverage`
  invariant, and the interview cells are Italian-only solo — so a language drift that slips one language is
  caught by the other (the `content-route-contract` cross-language-contract pattern, applied to the language
  axis). The live sweep (`content-stories.spec.ts`) additionally guards that **both** an English and an Italian
  result page are exercised. All assert STRUCTURE / PRESENCE / the language **LABEL** — never a translated
  content **value** (P2 copy-immunity: the test locale is pinned; copy/AI output is forbidden in assertions).

## 5. Provenance — the source-pins stay in the seed-generation manifest

The prod **source-pins** (which real session each exhibit was cloned from + the anonymization posture) live in
`seed-generation-manifest.yaml`'s **`content_sessions`** block (M232 — [`seed-manifest-spec.md`](seed-manifest-spec.md)
§8, [`session-clone-spec.md`](session-clone-spec.md)). `content-manifest.json` ALSO carries each session's
`source_session_id` so the render projection is self-disclosing, but the auditable source-of-truth disclosure
is the `content_sessions` block. The two are distinct projections of the same fixture — the render MENU
(`content-manifest.json`) and the audit DISCLOSURE (`content_sessions`) — exactly as `cockpit-manifest.json`
(the stories menu) is distinct from `seed-generation-manifest.yaml` (the seed intent).

## 6. Scope boundary — the manifest (M233) vs the render (M234) vs prove-it-lands (M236)

M233 delivered the **manifest** (the schema + the projection + the honesty gate + the fail-closed resolver + the
`--content-export` verb). **M234 (§7) delivers the render half** — the cockpit tab + the seat registration + the
bring-up wiring. **M235 (run 3) delivered the non-simulation product sections** (#M235-B2) — built + unit-proven:

- **The three non-simulation product sections** (skill-path-legacy / ai-labs / academy) are built as a
  **separate CODE-OWNED exhibit registry** (`seeders/content_nonsim.go` — `nonSimExhibits()` +
  `ContentStoryNonSimSeeder` + `buildNonSimProducts` appended by `BuildContentProducts`), NOT added to the
  simulation fixture (whose validator + seeder are simulation-shaped). Each has its OWN self-contained
  flat-index owner pairing (single-sourced with the seeder, exactly as the simulation projection). **Skill-path**
  — real progress: a seeded `skillpath.skill_path_sessions` row + the `local_skill_path_sessions` mirror
  (owned by a `content-player` seat, pinned to a REAL public `skill_path_id`), the `/skill-path/<id>` route —
  **player-link-only**: M236 iter-07 proved the per-user *manager* drill-down is UNIMPLEMENTED in next-web
  ("Coming soon", results table commented out, `userData` hardcoded null), so `has_manager_view` is FALSE and
  the mirror row buys no manager surface. **AI-labs** — presence-only (M231 §5): a `lab_sessions` status/spend
  row, NO CTA. **Academy** — app_base=academy, a real `/courses/<slug>` course CTA (direct origin, e2e_persona
  seam; **not** `/library/<slug>` — that route does not exist in ant-academy, M236 iter-08), no manager view.
  **`app/cmd/academy-seed` is MOOT in a demo** (see the §"academy" note above): with no
  `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` the demo academy serves its committed FS catalog, so a seeded
  `academy_chapter_progress` row has no reader. A
  `Label` field carries the believable row title (real course/lab names). Today's demo now renders all four
  product sections. rext tags `playbill-m235-nonsim-{skillpath,ailabs,academy}`.
- **Proving every CTA lands on a non-empty result page (a LIVE browser on a cold reset-to-seed) is M236**
  (prove-on-billion) — M235 unit-proves the seeders + the manifest projection (the sections resolve + the
  cockpit renders them); the live proof needs a running stack. M236 also AUTHORS the new content-stories
  seat-login coverage/Playthrough plumbing (M235's USER-BLOCKER-M235-02: the exact-path/hero-crawl harness
  can't reach the dynamic-URL, cockpit-seat-reached result pages — it must be authored + calibrated live) and
  works the per-section live-calibration checklists (skill-path version-match/status/mirror; ai-labs
  lab_sessions DDL; academy progress-write/route + M230 catalog fill).

## 7. The cockpit render — the 2nd "Content stories" tab (M234)

M234 turns the manifest into a **2nd cockpit tab** beside "Org stories" (`cockpit.py`, still stdlib-only,
standalone-served — the [`cockpit-spec.md`](cockpit-spec.md) panel). It is **the render half of this doc**.

### 7.1 The tabbed model
`cockpit.py --content-manifest <content-manifest.json>` (threaded by the bring-up — §7.5) adds a **client-side
tab toggle** (`_TAB_JS`, a raw string with no manifest data interpolated — the `_OVERLAY_JS` injection-safe
discipline) with two panels: **Org stories** (the heroes menu, default-on) and **Content stories** (the
`content_products[]` render). **Absent or empty `--content-manifest` ⇒ NO tab bar** — the page is
byte-identical to a pre-M234 single-panel cockpit, so an old bring-up is unchanged. The content menu is also
served at **`GET /content-manifest.json`**.

### 7.2 The tuple-regrouped row + the two-action contract (M234 contract, M242 layout)
Each product is a section (product FontAwesome icon + name). **The rows are REGROUPED by requirement tuple
(v2.6 "sound check" M242):** a product's played sessions group by **`(sim_type, modality)`** — a
non-simulation product (skill-path / academy / ai-labs, empty `sim_type` AND `modality`) falls back to
grouping by **`label`** — and each group renders as **ONE row**:

> **`target label` (+ modality pill)  |  passed login options  |  not-passed login options`**

a **per-`sim_type` FontAwesome icon** (`clipboard-check`/`dumbbell`/`user-tie`/`comments`) + the target label
(the sim_type human label, or the believable `label` for a non-sim) + the modality as a title pill, then **two
side-by-side columns** — the `passed:true` session's login options in one, the `passed:false` session's in the
other — so a presenter reads the passing run and the failing run of a requirement *next to each other* instead
of hunting two rows apart. An empty column reads an explicit **"No passing / No failing run"** (a distinct
marker, never a blank cell misread as broken). A **presence-only** group (ai-labs — no result surface, hence no
pass/fail verdict) renders a single inline cells slot, no columns.

Each column holds one **login-options cell** per session of that verdict — the **two-action contract**, and the
atomic unit the M241 EN/IT toggle filters (an EN/IT tuple contributes one cell per language into the same
column; the toggle shows one, hides the other). **The pass/fail moved from a per-session pill to the column
header; the modality from the desc to the tuple title.** The two actions per cell (unchanged from M234):

- **As-player** — a fake-FAPI handshake `…/handshake?__clerk_identity=<player_seat>&redirect_url=<base><player_result_path>`,
  rendered iff the session carries a `player_result_path`. `<player_seat>` is the `content-player-<idx>` seat
  M234 registered (§7.4), so the presenter logs in as the exact seeded member who owns the session.
- **As-manager** — the same handshake with the manager hero seat landing on the activity-dashboard result
  surface, **omitted where `has_manager_view=false`** (the `.sactions`/two-button layout with omitempty). The
  manager CTA is **always** a FAPI handshake (manager surfaces are next-web/hiring, never academy) (#M234-D4).

> **Render helpers (M242).** `render_content_tab` groups by `_content_tuple_key` → `_content_tuple_row`
> (icon + title + columns) → `_content_login_cell` (the per-session cell: language pill + the two CTAs,
> carrying `data-lang`/`lhide` for the toggle). The regroup is **render-layer only — no manifest schema
> change** (every session still carries `sim_type`/`modality`/`passed`/`language` as before).
>
> The "No passing / No failing run" empty marker is rendered server-side for a column that is cell-less at
> render time; because the EN/IT toggle hides cells at *view* time, `_LANG_JS.syncEmpty()` also re-derives the
> marker per column on every toggle (and on load), so an **unbalanced bilingual tuple** (a verdict present in
> only one language) never shows a verdict header over a blank body — the D3 "never a blank misread as broken"
> invariant holds under the toggle too (#M242-D8).

### 7.3 Per-product app-base routing + the two special sections
The per-product `app_base` resolves the CTA origin, generalizing the M224 `is_hiring`/`hiring_base` switch
(`content_base`): **`web`→`--app-base` :3000 · `hiring`→`--hiring-base` :3001 · `academy`→`--academy-base`
:3077** (an unset hiring/academy base falls back to `--app-base` — never a dead link). The two M231 special
dispositions:

- **AI-labs = PRESENCE-ONLY (M231 D4).** A session with no result surface (no `player_result_path`, no manager
  view) renders a **muted "Activity & spend only" status note — no CTA** (never a dead button). Data-driven:
  the disposition falls out of the manifest fields, not a flag (#M234-D5).
- **Academy (M231 D5).** app_base `academy` → the as-player CTA is a **direct academy-origin link** carrying the
  M53 `e2e_persona=member` cookie seam — the academy is a **separate origin with its own auth, NOT behind the
  fake FAPI** (a FAPI handshake redirect would establish a next-web session the academy can't see) (#M234-D4). No
  manager CTA (no academy manager route). The **specific-member** academy landing + the exact chapter route are
  finalized in **M235** (depends on M230), and the CTA went **live at M236 iter-08** as a real
  **`/courses/<slug>`** link into the demo academy's committed FS catalog (see §4 + §7.1).

  > *This previously read: "Today's fixture carries no academy session; this path is unit-proven and lights up
  > when M235 adds the fixture." **Stale** — the academy CTA is live, and it is **not** fixture-driven. A demo
  > academy has no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, so it serves its committed FS catalog and never reads a
  > seeded session; the CTA resolves against that catalog. There is no academy content-session fixture to wait
  > for, and `app/cmd/academy-seed` is moot on a demo.*

### 7.4 The seat registration — `content-player-<idx>` in the roster (`roster.go`)
The as-player CTA passes `?__clerk_identity=content-player-<idx>`, so that seat MUST resolve in Clerkenstein's
registry. Pre-M234 the exported roster carried **only heroes**; M234 **appends one identity per DISTINCT
content-player owner slot the projection references** (`contentPlayerSeatsUsed(BuildContentProducts)`) — no
dead seat, no unresolvable CTA. Each identity's claims (auth_id / eid / email / name / picture / org / role)
are derived with the **UsersSeeder's own functions** — the member NAME via the new single-source
`storyPopulationNames` (the UsersSeeder consumes it too, so the exported login identity is **byte-identical to
the seeded `public.users` row**) (#M234-D1). The seats append **after all heroes**, so `roster[0]` (the default active
seat) stays the first hero (#M234-D2). The existing `--roster-export` at bring-up carries them automatically (the roster
is a pure function of the blueprint — no bring-up change beyond §7.5). *The manager seat is the host org's
manager hero (`dan-manager`), already a roster seat — no new registration.*

### 7.5 The bring-up wiring (`up-injected.sh`)
The DEMO_STORIES cockpit block exports `content-manifest.json` via `stackseed --content-export` (a peer of
`--cockpit-export`/`--manifest-export`) and threads it into the cockpit launch as `--content-manifest`.
**NON-FATAL:** a failed or fail-closed export just drops the 2nd tab; the cockpit still serves "Org stories"
(the M18/M19 pattern). No new `/demo-up` flag or `DEMO_*` knob — the content tab is on whenever the storytelling
demo + cockpit are (the existing `DEMO_STORIES` / `DEMO_NO_COCKPIT` gates).

### 7.6 The EN/IT language toggle (M241 — v2.6 "sound check")
When the content manifest carries **≥ 2 toggle-able languages**, `render_content_tab` prepends an **EN | IT
segmented switch** (`_content_lang_toggle`) above the product sections, and emits `_LANG_JS` — a **raw-string,
injection-free** client filter, exactly the `_TAB_JS`/`_OVERLAY_JS` discipline (no manifest data interpolated).
The mechanics:

- **Per-row surfaces.** Every simulation row shows a **language pill** (`English` / `Italiano`; a solo row is
  tagged `<lang> only`). A **toggle-able** row (`lang_toggle=true`) additionally carries a `data-lang` attribute
  the filter keys on; a **solo** row (the Italian-only interview) carries **no** `data-lang`, so it always
  renders — the toggle skips it (the fallback: a single-language cell is not language-filtered).
- **The swap.** Clicking a language button shows `.session[data-lang=<lang>]` and hides the others (`lhide`) —
  so each bilingual cell's as-player / as-manager CTA re-targets that language's session. The tab opens on the
  **default (English)**: italian toggle-able rows start hidden server-side (`lhide`), so there's no flash before
  `_LANG_JS` runs. English is always present among toggle-able rows (a bilingual tuple has both), so the default
  always has rows.
- **Byte-clean when absent.** A manifest with **no** language axis (a pre-M241 bring-up, or a single-language
  fixture) renders **neither** the toggle **nor** `_LANG_JS` — the tab is unchanged. The cockpit stays
  stdlib-only (no new dep).
- **Proven (unit).** `demo-stack/tests/test_cockpit.py::TestContentLanguageToggle` render-proves the toggle
  structure + the default + the language labels + the solo-row-always-visible rule (STRUCTURE / LABEL only,
  never a translated value). The live click-swap is pure DOM (`_LANG_JS`); the live sweep proves both languages'
  result pages render (§4). *(M242 **delivered** the row-REGROUP by tuple (§7.2): the EN/IT variants of a
  `(sim_type, modality, pass/fail)` cell now share ONE column of ONE tuple row, the toggle filtering between
  them — `TestContentTabTupleRegroup::test_regroup_coexists_with_language_toggle_variants_in_one_column` pins
  the coexistence. M241 delivered the language axis + the global toggle the regroup consumes.)*

### 7.7 What's proven at M234 (unit) vs left to M236 (runtime)
M234 is **unit-proven, not browser-proven**: `cockpit.py` renders the manifest to correct HTML (per-product
sections, per-session rows, the two CTA hrefs with the right `__clerk_identity`/`redirect_url`, `has_manager_view`
omission, AI-labs presence-only, academy origin), the seats resolve through `roster.go` byte-identically to the
seed, and the export→render pipeline runs end-to-end. **Proving every CTA lands on a non-empty result page — a
live browser on a cold reset-to-seed — is M236** (the "prove-it-lands" milestone, run on `billion`). **M235
unit-proves the seeders + the manifest projection**; it does not drive a browser. *(This line previously read
"…is M235; proving it on `billion` is M236", which split one gate across two milestones and credited M235 with a
render proof it did not perform — M235 closed `closed-incomplete`. §4's statement of the same split is the
correct one.)*

## See also
- [`session-clone-spec.md`](session-clone-spec.md) (**M232**) — the seeder that COPIES the real sessions this menu points at + the `content_sessions` source-pins.
- [`content-stories-routes.md`](content-stories-routes.md) (**M231**) — the per-product result-route map + the `has_manager_view` matrix this projection encodes.
- [`seed-manifest-spec.md`](seed-manifest-spec.md) — the seed+gen manifest family this is a peer of (+ the `content_sessions` block §8).
- [`cockpit-spec.md`](cockpit-spec.md) — the presenter cockpit whose 2nd tab (M234) reads `content-manifest.json`.
- [`safety.md`](../safety.md) §3.8 — the VPN/tailnet-scoped read-side exception the copied sessions carry.
