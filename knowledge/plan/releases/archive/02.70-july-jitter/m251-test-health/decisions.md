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

## Adversarial review (M251 close, Phase 2c — 2026-07-23)
**Scenario — the vacuous removal-guard.** A removal-guard (`assertNotIn('class="btn academy"', page)`,
`assertNotIn("30000", js)`) can pass for the WRONG reason: if a later refactor renamed the token, moved the
render behind a flag these tests don't set, or the assertion targets a string that is trivially absent
regardless of the behaviour, the test stays GREEN while silently no longer guarding the removal it was
re-pointed to protect. A re-added per-hero academy link (or a re-added 30s stale-window) would then slip in
un-caught — the exact failure mode these re-points exist to prevent, inverted into a false sense of safety.
**Response — mutation-verify (recorded in `progress.md § M251: Hardening`).** Each removal-guard was proven
to be a *live* guard: the removed behaviour was temporarily re-introduced in the SOURCE (`cockpit.py`,
`gen_tailscale_serve.py`), the re-pointed test re-run, and confirmed to go **RED**, then the source
guaranteed-restored (git-clean asserted after). **4/4 guards fire:** overlay 30s window re-added → RED; the
`removeItem` try/catch stripped → RED; the per-hero `class="btn academy"` render re-added → all 4 academy
guards RED; the hiring `("hiring", 3001)` dropped from the generator → the port-13001 assertion RED. A guard
that goes RED when its target behaviour returns is not vacuous — it is exactly the check it claims to be.
The remaining live/env/docker-gated demo-stack failures are NOT M251's surface (→ M254, Fate 2) and were not
introduced by these re-points (full-suite run confirms M251's touched files are 207/207 green; the 8
failures live in `test_purge`/`test_ant_academy*`, untouched by M251).

## Inherited-deferral re-audit (M251 close, Phase 1b — 2026-07-23)
The milestone-scope deferral audit re-fated the four M246 drift-ledger rows that named M251 as a *possible*
("/") destination. **None land in M251** — M251's design-confirmed In-list is the run-unit roster + the
`test_cockpit`/`test_public_host` re-points only; each inherited row's authoritative destination is a
DIFFERENT open milestone, and M251 was only ever a secondary "or rext-hygiene / or M254" option (never a
sole commitment). All Fate 2, all confirmed-covered, no sibling-plan edit:
- **DEF-M246-03 / D-03** (`test_injection.py` pins skillpath-as-injected + models already-merged skiller) →
  **Fate 2, M247** — M246's own recommendation routes it through the durable `drift-ledger.md` handoff that
  M247 formally consumes (`depends_on: [M246]`); it lives in `stack-injection/tests/` (a module M251 does
  NOT own — coordination guardrail: M249 owns the demo-stack *patch* tests, M251 the *health/inventory*
  tests, and injection is neither). Inert (no compose service matches → green now).
- **DEF-M246-04 / D-04** (`exposure_claim_guard.py` `_cfg` lists skillpath:8095) → **Fate 2, M247** — mirrors
  D-03 ("update both together"); same drift-ledger→M247 triage. Test-only fixture, inert.
- **DEF-M246-08 / D-08** (fake-FAPI http-vs-TLS cheap-win probe artifact + end-to-end login not re-run) →
  **Fate 2, M254** — gate part (h) owns the live-browser + login re-prove; the http-vs-TLS probe-fix is only
  *verifiable* against a live fake-FAPI container (a live box M251, gated on a stackless local box,
  structurally lacks), so it correctly rides the M254 live billion prove rather than landing blind here.
- **DEF-M246-09 / D-09** (AI Academy peripheral, 2/6 env keys short, not serving) → **Fate 2, M254 /
  standing** — benign peripheral (non-fatal by design); the academy surface is a M254 gate concern.

**Aging check:** all four first-deferred 2026-07-23 (today), across 0 prior milestones; both authoritative
destinations (M247 drift-triage, M254 gate) are OPEN. Not repeat, not aged-out, not chronic. Verdict GREEN.
Full record: `audit-deferrals/deferral-audit-2026-07-23-m251-close.md`.
