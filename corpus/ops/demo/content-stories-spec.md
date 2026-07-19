# Content-Stories Spec — the content_products manifest + honesty gate + the cockpit render

**The M233 + M234 deliverable (v2.5 "the playbill", Thread B — the manifest half + the render half).** Where
[`session-clone-spec.md`](session-clone-spec.md) (M232) COPIES real production sessions into a demo org, and
[`content-stories-routes.md`](content-stories-routes.md) (M231) proved each result page reads a seedable row,
THIS doc specifies **(M233)** the **manifest** that turns those seeded sessions into a cockpit menu — a
**`content_products[]` projection** (§1–§6) — and **(M234)** the **cockpit render** that turns that manifest
into the 2nd "Content stories" tab + the `content-player-<idx>` seat registration that makes its as-player
CTAs log in (§7). The manifest is a projection from the same fixture the seeder seeds from, honesty-gated so
it can never drift, and fail-closed so it never fabricates a link.

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

`contentProductRegistry()` covers all four products the feature spans; the fixture today carries only
`simulation`, so only it projects (an empty product section is never fabricated). The others are schema-ready
for M234/M235's fixture additions.

| product id | app_base | section icon | player link? | manager surface |
|---|---|---|---|---|
| `simulation` | `web` | `flask` | yes — `/sim/<slug>/result/<sessionId>` | `ai-simulations` (or `interviews` for an interview sim_type) |
| `skill-path-legacy` | `web` | `diagram-project` | yes (M234) | `skill-paths` |
| `skill-path-new` (academy) | `academy` | `graduation-cap` | yes (M234) | — (no academy manager route, M231) |
| `ai-labs` | `web` | `vials` | **no — presence-only** (M231 §5) | — |

Per-**sim_type** row icons: ASSESSMENT `clipboard-check` · TRAINING `dumbbell` · HIRING `user-tie` ·
INTERVIEW `comments`.

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
- **manager_result_path** — `/enterprise/activity-dashboard/<kind>/<simId>/<userId>` where `<kind>` ∈
  {`ai-simulations`, `interviews`, `skill-paths`}, `<simId>` is the sim uuid, `<userId>` is the player member's
  id. Fully offline-derivable (no slug).

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

## 5. Provenance — the source-pins stay in the seed-generation manifest

The prod **source-pins** (which real session each exhibit was cloned from + the anonymization posture) live in
`seed-generation-manifest.yaml`'s **`content_sessions`** block (M232 — [`seed-manifest-spec.md`](seed-manifest-spec.md)
§8, [`session-clone-spec.md`](session-clone-spec.md)). `content-manifest.json` ALSO carries each session's
`source_session_id` so the render projection is self-disclosing, but the auditable source-of-truth disclosure
is the `content_sessions` block. The two are distinct projections of the same fixture — the render MENU
(`content-manifest.json`) and the audit DISCLOSURE (`content_sessions`) — exactly as `cockpit-manifest.json`
(the stories menu) is distinct from `seed-generation-manifest.yaml` (the seed intent).

## 6. Scope boundary — the manifest (M233) vs the render (M234) vs prove-it-lands (M235)

M233 delivered the **manifest** (the schema + the projection + the honesty gate + the fail-closed resolver + the
`--content-export` verb). **M234 (§7) delivers the render half** — the cockpit tab + the seat registration + the
bring-up wiring. Still ahead:

- **The non-simulation product player-path builders + their fixture sessions** (skill-path / academy / ai-labs)
  land with **M235**'s fixture additions (their route fields aren't in the fixture yet; the resolver fail-closes
  on them until then, and the renderer handles them the moment they appear — §7). **Today's fixture is
  simulation-only**, so a real demo renders only the Simulation section.
- **Proving every CTA lands on a non-empty result page** is **M235** (prove-it-lands) — including the academy
  section's exact deep-link + the specific-member academy landing (which also depends on M230's catalog fill).

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

### 7.2 The per-session row + the two-action contract
Each product is a section (product FontAwesome icon + name); each played session is a row: a **per-`sim_type`
FontAwesome icon** (`clipboard-check`/`dumbbell`/`user-tie`/`comments`) + a descriptor (modality pill +
passed/not-passed pill) + **up to two login-and-land CTAs**:

- **As-player** — a fake-FAPI handshake `…/handshake?__clerk_identity=<player_seat>&redirect_url=<base><player_result_path>`,
  rendered iff the session carries a `player_result_path`. `<player_seat>` is the `content-player-<idx>` seat
  M234 registered (§7.4), so the presenter logs in as the exact seeded member who owns the session.
- **As-manager** — the same handshake with the manager hero seat landing on the activity-dashboard result
  surface, **omitted where `has_manager_view=false`** (the `.actions`/two-button layout with omitempty).

### 7.3 Per-product app-base routing + the two special sections
The per-product `app_base` resolves the CTA origin, generalizing the M224 `is_hiring`/`hiring_base` switch
(`content_base`): **`web`→`--app-base` :3000 · `hiring`→`--hiring-base` :3001 · `academy`→`--academy-base`
:3077** (an unset hiring/academy base falls back to `--app-base` — never a dead link). The two M231 special
dispositions:

- **AI-labs = PRESENCE-ONLY (M231 D4).** A session with no result surface (no `player_result_path`, no manager
  view) renders a **muted "Activity & spend only" status note — no CTA** (never a dead button). Data-driven:
  the disposition falls out of the manifest fields, not a flag.
- **Academy (M231 D5).** app_base `academy` → the as-player CTA is a **direct academy-origin link** carrying the
  M53 `e2e_persona=member` cookie seam — the academy is a **separate origin with its own auth, NOT behind the
  fake FAPI** (a FAPI handshake redirect would establish a next-web session the academy can't see). No manager
  CTA (no academy manager route). The **specific-member** academy landing + the exact chapter route are
  finalized in **M235** (depends on M230). *Today's fixture carries no academy session; this path is
  unit-proven and lights up when M235 adds the fixture.*

### 7.4 The seat registration — `content-player-<idx>` in the roster (`roster.go`)
The as-player CTA passes `?__clerk_identity=content-player-<idx>`, so that seat MUST resolve in Clerkenstein's
registry. Pre-M234 the exported roster carried **only heroes**; M234 **appends one identity per DISTINCT
content-player owner slot the projection references** (`contentPlayerSeatsUsed(BuildContentProducts)`) — no
dead seat, no unresolvable CTA. Each identity's claims (auth_id / eid / email / name / picture / org / role)
are derived with the **UsersSeeder's own functions** — the member NAME via the new single-source
`storyPopulationNames` (the UsersSeeder consumes it too, so the exported login identity is **byte-identical to
the seeded `public.users` row**). The seats append **after all heroes**, so `roster[0]` (the default active
seat) stays the first hero. The existing `--roster-export` at bring-up carries them automatically (the roster
is a pure function of the blueprint — no bring-up change beyond §7.5). *The manager seat is the host org's
manager hero (`dan-manager`), already a roster seat — no new registration.*

### 7.5 The bring-up wiring (`up-injected.sh`)
The DEMO_STORIES cockpit block exports `content-manifest.json` via `stackseed --content-export` (a peer of
`--cockpit-export`/`--manifest-export`) and threads it into the cockpit launch as `--content-manifest`.
**NON-FATAL:** a failed or fail-closed export just drops the 2nd tab; the cockpit still serves "Org stories"
(the M18/M19 pattern). No new `/demo-up` flag or `DEMO_*` knob — the content tab is on whenever the storytelling
demo + cockpit are (the existing `DEMO_STORIES` / `DEMO_NO_COCKPIT` gates).

### 7.6 What's proven at M234 (unit) vs left to M235 (runtime)
M234 is **unit-proven, not browser-proven**: `cockpit.py` renders the manifest to correct HTML (per-product
sections, per-session rows, the two CTA hrefs with the right `__clerk_identity`/`redirect_url`, `has_manager_view`
omission, AI-labs presence-only, academy origin), the seats resolve through `roster.go` byte-identically to the
seed, and the export→render pipeline runs end-to-end. **Proving every CTA lands on a non-empty result page (a
live browser on a cold reset-to-seed) is M235; proving it on `billion` is M236.**

## See also
- [`session-clone-spec.md`](session-clone-spec.md) (**M232**) — the seeder that COPIES the real sessions this menu points at + the `content_sessions` source-pins.
- [`content-stories-routes.md`](content-stories-routes.md) (**M231**) — the per-product result-route map + the `has_manager_view` matrix this projection encodes.
- [`seed-manifest-spec.md`](seed-manifest-spec.md) — the seed+gen manifest family this is a peer of (+ the `content_sessions` block §8).
- [`cockpit-spec.md`](cockpit-spec.md) — the presenter cockpit whose 2nd tab (M234) reads `content-manifest.json`.
- [`safety.md`](../safety.md) §3.8 — the VPN/tailnet-scoped read-side exception the copied sessions carry.
