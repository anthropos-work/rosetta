# M251 — Spec notes

Mechanical-fix findings, exact `file:line` anchors, and re-point rationale accumulate here during build.
All code lives in the **rext authoring copy** (`.agentspace/rosetta-extensions/`, a separate gitignored
repo) — committed + tagged + pushed to origin there; the rosetta footprint is only these plan artifacts.

## Pre-flight audits — run-unit roster fix + python re-points (single session)
- **Phase 0b KB-Fidelity: GREEN** (report: `kb-fidelity-audit.md`, sha at audit = post-M246 HEAD `99d597c`).
  Docs describe current truth; the stale artifacts are rext *test code* (M251's subject), not knowledge docs.
  One topic is intentionally DOC-ONLY (the optional `verification.md` demo-stack-suite anchor — deferred to
  M247 per the cross-lane guardrail). No blind areas, no stale load-bearing doc claims.
- Both sections of this session share this verdict (same subsystem, one small diff) per the audit-reuse rule.

## Topic → doc → code triples (fast-start for future audits)
- run-unit roster → `coverage-protocol.md` → `rext stack-verify/e2e/run-unit.sh` + `stack-verify/tests/test_e2e_collection_integrity.py::UnitSpecsAreExecuted`
- overlay 30s window → `cockpit-spec.md:271-277` + `latency-budget.md` → `rext demo-stack/cockpit.py::_OVERLAY_JS` + `tests/test_cockpit.py::TestOverlayJs`
- per-hero academy link (removed) → `cockpit-spec.md` (no render claim) → `rext demo-stack/cockpit.py:858-862` + `tests/test_cockpit.py::TestAcademyLink/TestAcademyCatalogEntryEdges/TestServedPanelWithAcademy`
- hiring 13001 fronting → `tailscale-serve.md` + `safety.md §3` → `rext stack-injection/gen_tailscale_serve.py:44` + `tests/test_host_prereqs_m215.py::TestF12ServeResetGenerator`

## run-unit roster fix (content-denominator + run-discrete unit specs)
**Symptom:** `test_e2e_collection_integrity.py::UnitSpecsAreExecuted::test_every_unit_spec_is_named_by_a_runner`
is RED — 9 `*.unit.spec.ts` on disk but only 7 in the `run-unit.sh` `UNIT_SPECS` roster. The 2 orphans:
`content-denominator.unit.spec.ts` + `run-discrete.unit.spec.ts` (added by v2.6 but never rostered — the exact
COLLECTED-not-EXECUTED defect the guard exists to catch). `run-unit.sh`'s own `unrostered` pre-check also
`exit 2`s on them.
**Fix:** add the 2 specs to `UNIT_SPECS` in `stack-verify/e2e/run-unit.sh` (7 → 9). Verify: `run-unit.sh`
exits 0, prints "9 files", ≥40 tests pass; `UnitSpecsAreExecuted` goes GREEN.

## Mechanical python re-points (test_cockpit + test_public_host)
All are STALE assertions of deliberately-changed behaviour (verified against current code by rendering the
page + reading `_OVERLAY_JS` + the generator port sets). Ground-truth confirmations in-line below.

### Overlay (M218 — the 30s in-flight window was REMOVED; `cockpit.py:314-319,353-356`)
Current `_OVERLAY_JS`: no `30000` constant, `getItem` read-path gone, `removeItem`+`resetOverlayOnReturn`+
`pageshow` clear-on-return present; `try {` count = **2** (setItem write + removeItem clear).
- `TestOverlayJs::test_inflight_window_is_30s` (`tests/test_cockpit.py:997`) — asserts `"30000"` present.
  RE-POINT → assert the window is GONE: no `30000`; the flag (`cockpit.login.inflight`) is CLEARED on return
  (`removeItem` + `pageshow`), not compared to a freshness window. Rename to reflect the removal.
- `TestOverlayJs::test_localstorage_access_is_guarded` (`:1004`) — asserts `try {` count ≥ 3.
  RE-POINT → **2** guarded sites now (write + clear); the getItem read-path was removed with the window.
- Class docstring (`:975-978`) mentions "the 30s in-flight window" → update to the removal.

### Academy per-hero link (REMOVED 2026-07-15 — `cockpit.py:858-862`)
Ground truth: `render_page(..., academy_base=ACADEMY)` yields `class="btn academy"`=0, `33077`=0. Only the
retained CSS (`a.academy`) + JS seam (`a.academy[data-academy-persona]`) remain (dead — no matching element).
The `class="btn academy"` that DOES exist (`cockpit.py:541`) is the content-stories "As player" link, gated on
`content_manifest` (not passed by these tests). All 4 render-side tests assert the removed per-hero link.
- `TestAcademyLink::test_academy_link_renders_per_hero_when_base_set` (`:280`) — RE-POINT → the per-hero
  academy link does NOT render even with `academy_base` + a catalog entry (guards the removal; a re-add is loud).
- `TestAcademyCatalogEntryEdges::test_render_defaults_academy_path_persona_label_when_absent` (`:342`) —
  RE-POINT → a tampered/partial academy entry produces NO rendered link (safe, no crash). Selector tests in
  the same class (`test_selects_academy_entry_among_multiple_externals` etc.) still pass — the SELECTOR is retained.
- `TestAcademyCatalogEntryEdges::test_render_academy_entry_fields_are_escaped` (`:358`) — RE-POINT → a hostile
  entry produces NO rendered academy link ⇒ no injection surface (removal closed the XSS surface entirely).
- `TestServedPanelWithAcademy::test_root_serves_academy_link` (`:570`) — RE-POINT → the served panel does NOT
  serve a per-hero academy link. Update the class docstring (`:556`).

### Public-host port fronting (M226/M220 — hiring 3001 → 13001 joined the UI tier)
`gen_tailscale_serve.py:42-50` UI_BROWSER_FACING now includes `("hiring", 3001)` (M226) → the generator's
serve/reset fronts `--https=13001`. `_UI_PORTS` in the test lags.
- `test_host_prereqs_m215.py::TestF12ServeResetGenerator::test_public_host_emits_per_port_off_for_all_browser_ports`
  (`:438`; the `_UI_PORTS` constant `:418`) — RE-POINT → add `3001` to `_UI_PORTS` (`{3000, 9000, 3077}` →
  `{3000, 3001, 9000, 3077}`). Only the one assertion consumes `_UI_PORTS`, so the change is contained.

## Out of scope (ride M254 — live/docker box; NOT touched here)
`test_purge` (docker-integration) + the live process-lifecycle assertions that also currently fail on a
box with no live stack: `test_ant_academy.py` launcher/reap (5) + `test_host_isolation::test_mutant_no_term_trap_is_caught`
+ `test_ant_academy_clerk_wiring::test_overlay_has_minted_pk_and_no_real_secret` (bash-runtime, exit 127).
These are "any live-serve-gated assertion" per the milestone Out list → M254. Left untouched by design.
