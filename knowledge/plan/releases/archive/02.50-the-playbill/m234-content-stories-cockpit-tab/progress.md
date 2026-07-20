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

## M234: Final Review

Consolidated findings from `/developer-kit:close-milestone` Phases 1–5 (2026-07-19). Near-clean close:
harden Pass 1 already deepened coverage (22 tests, 0 bugs) and the cross-cutting review surfaced only two
record-level items (no code change). Deferral re-audit (Phase 1b) **YELLOW / 0 blockers** — see
`audit-deferrals/deferral-audit-2026-07-19-m234-close.md`.

### Scope
- [x] All `overview.md` In: items delivered (§1 seat registration / §2 tab render / §3 bring-up wiring); the only
      unchecked progress box (`roadmap/state at close`) is Phase 10's own job — not a gap.
- [x] Both open questions resolved: academy deep-link disposition → D-M234-4 (direct academy-origin link; specific-member
      landing finalized in M235); manager drill-down → D-M234-4 (manager CTA = FAPI handshake on the manifest's manager result path).
- [x] 0 TODO/FIXME/HACK in any M234-touched file.

### Code Quality
- [x] [confirmed] `go build ./...` + `go vet ./...` clean (stack-seeding); harden cross-cutting review found 0 bugs; consistent with existing seeder/cockpit patterns; no dead code / no new unit-without-handbook.

### Documentation
- [x] `content-stories-spec.md` §7 (render half) accurate vs the shipped `cockpit.py`; §6 M234/M235 boundary correct.
- [x] `cockpit-spec.md` `/content-manifest.json` endpoint row + "2nd tab (M234)" section present + accurate.
- [x] No new top-level unit introduced → per-unit-handbook contract N/A.

### Tests & Benchmarks
- [x] Go stack-seeding: all pkgs OK, build+vet clean.
- [x] Python demo-stack: 249 pass / **6 pre-existing fail** (removed-academy-CTA ×4 + overlay-30s ×2) / **0 new** — the M234 contract held. The 6 are the chronic-carry subset homed at v2.5 release-close.
- [x] Handbook test-count reconciliation: no handbook quotes a drifted M234 count (counts live in progress.md's hardening table, matched to runner output).

### Adversarial review (Phase 2c)
- [x] Record the harden-exercised adversarial scenarios in `decisions.md` under an `Adversarial review` subsection (scenario, not just the test) — **fixed** (see decisions.md).

### Decision Triage
- [x] D-M234-1/2/4/5 → blended into `content-stories-spec.md` §7 (single-source seat, append-after-heroes, two-action/academy contract, presence-only-is-data-driven); add `(#M234-DK)` back-ref tags — **fixed**.
- [x] D-M234-3 (hero-scoped `--roster-export` warning) → archive (maintainer-only CLI detail; stays in decisions.md).

## Completeness Ledger (Phase 9 — section variant)

Every `overview.md` `In:` scope item placed into exactly one three-fate category. **Zero escape-hatch deferrals.**

### Done (Fate 1 — delivered in M234)
- Client-side tab toggle in `render_page()` (`_TAB_JS`, stdlib-only, no manifest data in JS) — §2.
- Per-product sections rendering the M233 manifest; per-session rows with per-`sim_type` FontAwesome icons — §2.
- Two fake-FAPI deep-link CTAs per session + `.actions` two-button layout + `has_manager_view` omitempty — §2.
- Per-product app-base routing generalizing the `is_hiring`/`hiring_base` switch (web:3000/hiring:3001/academy:3077 + app-base fallback) — §2.
- Mint/resolve `content-player-<idx>` player seats via `roster.go` + Clerkenstein (single-sourced with UsersSeeder via `storyPopulationNames`) — §1.
- AI-labs section = PRESENCE-ONLY, no CTA (M231 D4) — the RENDER disposition — §2.
- Academy section RENDER disposition = player CTA to the academy origin (M53 cookie seam, not FAPI), no manager CTA (M231 D5 / D-M234-4) — §2.
- Bring-up wiring: `up-injected.sh` `--content-export` + `--content-manifest` (non-fatal) — §3.
- Docs deliverable: `content-stories-spec.md` §7 (render half) + `cockpit-spec.md` 2nd-tab note.

### Confirmed-covered (Fate 2 — already owned by a milestone of this release)
- Non-simulation product **fixtures** (ai-labs / academy / skill-path) + **prove-every-CTA-lands live** → **M235** (`In:` + exit_gate; the renderer already handles all dispositions at M234, unit-proven — so nothing is stranded).
- **Specific-member** academy landing + exact chapter route → **M235** (depends on M230's catalog fill). This is literally M234's overview open question; D-M234-4 records the handoff.

### Annotated (Fate 3 — attached to a release-milestone at close)
- None. M235's `In:` + exit_gate already own the carried items; no roadmap edit required.

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch)
- None.

### Cross-reference (inherited, non-M234-In: — release-level)
- 14 pre-existing demo-stack test failures (6 in `test_cockpit.py`) → v2.5 **release-close** re-anchor (REPEAT/CHRONIC; M234 added 0; user-dispositioned — `audit-deferrals/deferral-audit-2026-07-19-m234-close.md`).
