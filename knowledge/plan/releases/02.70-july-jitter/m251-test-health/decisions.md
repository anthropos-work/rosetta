# M251 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

All code changes land in the **rext authoring copy** (separate gitignored repo); rosetta carries only these
plan artifacts. Consumption tag `july-jitter-m251-test-health` (pushed to origin — rung zero).

## D-1 — run-unit roster: add the 2 orphan unit specs (not a glob)
`content-denominator.unit.spec.ts` + `run-discrete.unit.spec.ts` shipped in v2.6 but were never added to
`run-unit.sh`'s `UNIT_SPECS` roster, so they were COLLECTED (counted toward `MIN_TESTS`) but EXECUTED by
nothing — the exact defect `UnitSpecsAreExecuted` exists to catch. **Decision:** add both to the EXPLICIT
roster (kept an explicit list, NOT a glob — the runner's own header argues a glob would silently re-introduce
the same "unexecuted spec nobody noticed" failure; the two-way disk↔roster check keeps a future omission
loud). 7 → 9 specs; run-unit exits 0 (172 tests); `test_e2e_collection_integrity` 8/8 GREEN.

## D-2 — overlay tests: pin the M218 REMOVAL of the 30s in-flight window (re-point, don't delete)
The cockpit login overlay's localStorage in-flight flag lost its 30s freshness window at M218 (`_OVERLAY_JS`;
cockpit.py:314-319,353-356; documented in `cockpit-spec.md:271-277`): at a 2.4s login the window re-showed a
spinner over an already-loaded workspace. Now ANY return CLEARS the flag (`removeItem` + `pageshow`).
**Decision:** re-point (not delete) the two stale tests to guard the current truth — `test_inflight_window_is_30s`
→ `test_inflight_flag_is_cleared_on_return_no_stale_window` (asserts NO `30000`, key retained, cleared on
return); `test_localstorage_access_is_guarded` → assert the **2** surviving guarded sites (write + clear; the
getItem read-path third block was removed with the window). Preserving the tests as removal-guards means a
silent re-introduction of the stale-verdict window goes RED.

## D-3 — public-host port-13001: add hiring 3001 to `_UI_PORTS` (mirror the generator)
`test_host_prereqs_m215.py::TestF12ServeResetGenerator` compares the generator's serve/reset port set against
a local `_UI_PORTS` mirror. M226 added the REAL apps/hiring 2nd app (base 3001 → 13001 at offset) to the UI
tier's `tailscale serve`/reset fronting (`gen_tailscale_serve.py:UI_BROWSER_FACING`; `tailscale-serve.md`),
but the test's `_UI_PORTS = {3000, 9000, 3077}` never gained 3001, so the generator emitted 13001 and the
assertion's expected set didn't. **Decision:** add `3001` to `_UI_PORTS`. Only the one assertion consumes the
constant; the sibling `test_reset_is_the_exact_inverse_port_set_of_serve` stays consistent (serve & reset both
include 13001). Whole class 7/7 GREEN.

## D-4 — academy tests: the per-hero [Academy] link was REMOVED — invert to removal-guards, keep selector/helper
The per-hero cockpit [Academy] link (M53 F6) was **deliberately removed 2026-07-15** (cockpit.py:858-862:
"LOGIN is the only per-hero CTA"). Ground truth (render `render_page(..., academy_base=ACADEMY)`):
`class="btn academy"`=0, academy origin `33077`=0 — nothing renders a per-hero academy link. The
`academy_url` helper, `_academy_catalog_entry` selector, `a.academy` CSS and persona-cookie JS seam are
RETAINED (dead — no matching element).

**Decision:** the 4 stale RENDER-side tests asserted the removed link. Re-point (not delete) each to guard the
REMOVAL — `test_academy_link_renders_per_hero_when_base_set` → `test_no_per_hero_academy_link_even_when_base_set`;
`test_render_defaults_...` → `test_partial_academy_entry_renders_no_link_and_does_not_crash`;
`test_render_academy_entry_fields_are_escaped` → `test_hostile_academy_entry_produces_no_rendered_link` (the
removal closed the injection surface entirely — no rendered surface interpolates the entry); `test_root_serves_academy_link`
→ `test_root_serves_no_per_hero_academy_link`. The SELECTOR/helper/cookie-seam tests in the same classes were
already GREEN and are LEFT untouched (they cover the retained code). **Alternatives considered:** (a) delete
the 4 render tests — rejected: a removal with no guard lets a silent re-add slip in; (b) fabricate a rendered
academy surface to keep the positive assertions — rejected: that would assert behaviour the code doesn't
exhibit. Chose the truthful removal-guard. This keeps the test count stable and each assertion truthful.

## Surfaced but OUT of M251 scope → M254 (three-fate: Fate 2, already planned)
Running the full demo-stack suite on a **stackless dev box** surfaces additional live/env/docker-gated
failures that are NOT M251's mechanical re-points and were left untouched per the milestone Out list ("any
live-serve-gated assertion → M254"):
- `test_purge::...THE_BUG` (docker-integration — the data-dir purge needs a real container-owned 0700 dir).
- `test_ant_academy.py` launcher/reap (×5) — spawn stub npm/node, bind ports, record/kill PIDs (live
  process-lifecycle; fail here on pidfile timing "no running academy recorded").
- `test_host_isolation::test_mutant_no_term_trap_is_caught` (×1) — a signal-trap mutation test (live process).
- `test_ant_academy_clerk_wiring::test_overlay_has_minted_pk_and_no_real_secret` (×1) — executes a bash
  snippet that sources `detach.sh` with an unbound `$HERE` → exit 127 (env/bash-runtime harness).

**Fate 2 — confirmed covered by M254.** M254's exit gate part **(g)** owns "the ~2 docker/live-gated
test-health tests green" and part **(h)** "the live-browser specs … green", both on a live `billion` box.
**Flag for the M254 driver:** the live-gated failing set on a stackless box is **8**, larger than the "~2"
M254's overview names — the M254 closer should expect to green (or re-triage) the academy-launcher/reap +
host-isolation + clerk-wiring assertions too, not just `test_purge` + one live-serve assertion. NOT annotating
M254's overview from this lane (cross-lane collision avoidance, per the M251 brief); surfaced here + in the
step report for the serialized closer to route.

## Deferred to M247 (optional deliverable)
The optional `corpus/ops/verification.md` demo-stack-python-suite + run-unit-roster index anchor
(overview.md "Delivers →") is DEFERRED to M247, which owns the corpus reground and `verification.md` edits —
authoring it here would collide with M247's concurrent lane on the same doc. Not a blind area (the code it
would index exists + is exercised).
