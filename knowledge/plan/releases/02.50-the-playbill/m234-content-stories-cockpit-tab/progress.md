# M234 — Progress

## Sections

Derived from `overview.md` + `roadmap.md` § M234. M234 = the **cockpit-UX render half** of Content
stories (the manifest half is M233; the fixture-population + prove-it-lands is M235). Three implementation
sections + docs folded into each section's Phase 5.

### §1 — Content-player seat registration (roster.go) ✅ (rext 6f7001b)
The as-player CTA (`?__clerk_identity=content-player-<idx>`) must authenticate as the real seeded member
who owns the session. Extend the exported roster so those seats resolve, single-sourced with the seeder's
own id/claim derivation.
- [x] Extract `storyPopulationNames` (pure replay of the UsersSeeder name assignment) + refactor UsersSeeder to consume it (true single-source, no drift)
- [x] `contentPlayerRosterIdentities(s)` — one roster identity per DISTINCT content-player owner slot the content-manifest projection references (auth_id / eid / email / name / picture / org-claims / role=member single-sourced)
- [x] Append content-player seats to `BuildRoster` output (after all heroes → default seat unchanged)
- [x] Tests: claims match UsersSeeder's writes; seat key == manifest `player_seat`; first roster entry still a hero; no duplicate keys; Clerkenstein-decodable shape (roster↔registry lockstep) — 6 new tests, full module green

### §2 — Content stories cockpit tab render (cockpit.py) ✅ (rext cbca13c)
The 2nd "Content stories" tab beside "Org stories": per-product sections rendering the M233 manifest.
- [x] `--content-manifest` arg + shape validation (byte-identical page when absent — old bring-up safe)
- [x] Client-side tab toggle (Org stories | Content stories) reusing the stdlib `_OVERLAY_JS`-style pattern; no manifest data interpolated into JS
- [x] Per-product sections (product FontAwesome icon + name); per-session rows with per-`sim_type` row icons + descriptor (modality / passed vs not-passed)
- [x] Two fake-FAPI deep-link CTAs per session (as-player / as-manager), `.actions` two-button layout + `has_manager_view` omission
- [x] Per-product app-base routing generalizing the `is_hiring`/`hiring_base` switch (web :3000 / hiring :3001 / academy :3077)
- [x] AI-labs section = PRESENCE-ONLY (no player path → status/spend line, no CTAs) — M231 D4
- [x] Academy section = player CTA to the academy origin (app_base=academy), no manager CTA — M231 D5 (direct academy link + M53 cookie seam; specific-member landing → M235)
- [x] Serve `/content-manifest.json` endpoint
- [x] Python unit tests: 23 new — dispositions, CTA hrefs, omission, presence-only, academy origin, tab toggle, served endpoint, shape validation, escaping; 0 new failures (106 tests, 6-fail baseline unchanged)

### §3 — Bring-up wiring (up-injected.sh) ✅ (rext 7f55eb4)
- [x] Export `content-manifest.json` via `stackseed --content-export` at bring-up (non-fatal, parallel to `--cockpit-export`; fail-closed export just drops the tab)
- [x] Thread `--content-manifest` into the `cockpit.py` launch (+alt guard, set -u safe)
- [x] Verify the export + launch wiring — new `StorytellingCockpitWiring` test (10/10) + proven end-to-end (real export → cockpit → probe: tab renders 9 sessions × 2 CTAs, `/content-manifest.json` served)

### Docs (Phase 5, folded) ✅
- [x] Extend `corpus/ops/demo/content-stories-spec.md` with the render half — new §7 (tabbed model, two-action contract, per-type icon map, per-product base routing, AI-labs presence-only, academy origin, seat registration, bring-up wiring, unit-vs-M235 boundary); header + §6 updated
- [x] Note in `cockpit-spec.md` — the `/content-manifest.json` endpoint row + the "2nd tab (M234)" section pointing to `content-stories-spec.md` §7
- [ ] roadmap/state at close-of-milestone (deferred to `/developer-kit:close-milestone` per its contract)

## M234: Hardening

### Pass 1 — 2026-07-19
One comprehensive pass; the M234-touched-code scan was exhausted in a single sweep (a Pass 2 would only
reach pre-existing/out-of-scope lines), so the loop stopped after Pass 1.

**Scope manifest (milestone-touched, rext clone @ `playbill-m234-content-tab` 7f55eb4):**
- Python (demo-stack): `cockpit.py` (+302, the content-tab render) · `up-injected.sh` (+16, bring-up wiring)
  — tests `tests/test_cockpit.py`, `tests/test_tooling.py`
- Go (stack-seeding): `seeders/roster.go` (+82) · `seeders/userprofile.go` (+30, `storyPopulationNames`) ·
  `seeders/content_manifest.go` (+32) · `seeders/users.go` (+29, refactor to consume the shared names) ·
  `cmd/stackseed/main.go` (+12, the hero-scoped warning) — tests `content_player_roster_test.go`,
  `roster_test.go`, `main_test.go`, `cockpit_test.go`
- Every touched source file had co-located tests (no new-unit-without-handbook gap; no new package/tool).

**Coverage delta (milestone-touched files):**
- `cockpit.py` statements: 92% → **96%** (+4). Every M234-touched content-tab path now covered (493/499
  render_content_tab dispositions, 809/815 content-validator raises, 872-878/901-902 main() --content-manifest).
  The residual 10 missed lines are **all pre-existing / out-of-scope**: `_validate_manifest_shape` raises
  (790-799, M217), the bind-error branch (890-896, M217), `except KeyboardInterrupt` + `__main__` boilerplate
  (911-912, 919). Not M234-touched → left for the release-close re-anchor (same class as the 6 pre-existing fails).
- Go (both packages measured): **100% function coverage** on all M234-touched files except `users.go Seed`
  (97% — a pre-existing DB-write error path needing a live DB, not M234-introduced). `RosterHeroCount` was
  already line-covered by the CLI test, but its **D-M234-3 invariant was not behaviourally pinned** — now it is.

**Tests added (22 total — behaviour, not lines; 0 shallow):**
- `test_cockpit.py` (+19, 6 classes):
  - `TestContentTabRenderEdges` (5) — empty-product skip (siblings still render), all-empty friendly note
    (never a blank panel), has_manager_view **promised without a seat** omits the manager CTA, an **academy
    session with a manager view** routes the manager CTA to `app_base` (not the academy origin), `content_base`
    hiring/academy key falls back to `app_base` when the base is unset (--no-ui).
  - `TestContentAcademyJsGate` (4) — `_content_uses_academy` true/false grid + a web-only tab emits **no inert
    academy-cookie JS** (byte-clean).
  - `TestContentTabJsGate` (2) — no `_TAB_JS` without content; present + brace/paren-balanced with content.
  - `TestContentManifestShapeValidationDepth` (3) — not-a-dict + product-not-object rejections; null-sessions accepted.
  - `TestContentTabByteIdenticalFallback` (3) — absent / empty-products / empty-dict content manifest each
    render **byte-identical** to the pre-M234 single-panel page (the back-compat contract, stronger than markup-absence).
  - `TestContentTabMainWiring` (2) — `main()` `--content-manifest` end-to-end (served tab + `/content-manifest.json`);
    a broken content manifest is **non-fatal** (single panel, endpoint 404s).
- `test_tooling.py` (+1) — the content-export failure is a **non-fatal warning** (arg-array defaults empty
  before the on-success set → the `+alt` guard drops the tab, never aborts the bring-up).
- `content_player_roster_test.go` (+2) — the **D-M234-3 `RosterHeroCount` invariant**: it is the HERO
  partition (strictly < the total), and a 0-hero structural Workforce org still projects content-players yet
  `RosterHeroCount` stays 0 (the CLI "0 heroes" signal isn't masked by the inflated total).

**Bugs fixed inline:** none — every M234 code path behaved correctly under the deepened tests (pure
test-only deepening).

**Flakes stabilized:** none observed. Flake gate 3/3 clean on both stacks (Python 20 new; Go 2 new).

**Knowledge backfill:** no KB-worthy findings — the hardening confirmed already-documented behaviours
(the render_content_tab dispositions, the byte-identical fallback, the D-M234-3 invariant) rather than
surfacing new invariants; all are already stated in `content-stories-spec.md` §7 + `decisions.md`.

**rext code-of-record:** committed `fd457bf` in the rext authoring clone, re-tagged
**`playbill-m234-content-tab-hardened`** (the original `playbill-m234-content-tab` @ 7f55eb4 stays pinned).
Test-only; **0 platform-repo edits**; 0 net-new deps.

### Stop condition
Loop stopped after Pass 1: the Step 2b scan for M234-touched code was exhausted (a further pass would only
reach pre-existing M217 / boilerplate lines, out of scope), coverage on touched files is complete, and the
flake gate is clean (3/3). Not chased (documented pre-existing, release-close re-anchor): the 6 `test_cockpit.py`
failures (removed academy CTA ×4 + overlay-JS ×2) + the pre-existing `_validate_manifest_shape` / bind-error
coverage gaps. `/developer-kit:close-milestone` Phase 4 runs independently as defense-in-depth.
