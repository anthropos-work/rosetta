# Hardening Ledger — M244 prove-on-billion

Incremental harden of the closed tik iters since the last harden. M244 is a live-prove milestone;
its fixes' CODE lives in `rosetta-extensions` (rext, main), so the harden deepens the UNIT/mutation
coverage of those fixes there and consumes them via the pinned tag
`sound-check-m244-content-sweep-robustness`. The rosetta-side commits here are the ledger only.

## Pass 1 — 2026-07-22 — incremental

**Iters hardened this pass:** iter-02 … iter-11 (first harden of the milestone; the toks iter-01/iter-12 are out of scope).
**Tiks covered since prior pass:** 10 (iters 02–11).

**Thesis probed (anti-toothlessness):** for each of the six iter fixes that shipped a test, MUTATION-VERIFY — break the subject, confirm the test goes RED, restore. Result: 3 fixes had teeth already; **3 had a toothless test or a coverage gap** that stayed GREEN under the exact regression they exist to catch. All fixed inline (Fate 1), each mutation-verified to bite.

**Mutation-verification outcomes (the crux):**
- iter-06 interview plan-section-id ALIGNMENT — **TEETH CONFIRMED (no change).** Live mutation: dropped a used manager section (`adoption`) from the plan golden → `TestInterviewPlanSectionAlignment` flagged ORPHANED keys for BOTH interview sessions (intv-voice-pass + intv-voice-fail). Catches drift, not just non-emptiness. Unit `TestOrphanKeysDetectsV13Drift` locks the exact v1.3 keys.
- iter-10 reap-17700 test-isolation — **TEETH CONFIRMED (no change).** Held :17700 with a foreign listener; the RACED test PASSES on the free-port default; reverting the default to the hardcoded 17700 makes it FALSE-FAIL (reap refuses the really-held foreign port). The fix genuinely isolates the test from ambient infra.
- iter-08 demopatch scope — **TOOTHLESS → FIXED.** `test_result_demo_widen_is_manager_scoped` used an UNANCHORED `assertRegex(r"isManagerScope\s*&&\s*!\(…")`; `!isManagerScope && !(…)` (the PLAYER-scope leak the fix forbids) also matches as a substring, so the test false-PASSED on the mutation it exists to catch. Strengthened: paren-anchored positive + `assertNotRegex` forbidding a negated-scope widen. Mutation (flip to `!isManagerScope`) → RED via real AssertionError (was GREEN).
- iter-05 simulations_extraction — **GAP → FIXED.** No test looked at the captured COLUMN list; dropping `schema` (the plan JSON the whole fix carries) would leave surface/order/scope tests green while re-rendering the interview report EMPTY. Added `TestSurface_SimulationsExtractionCapturesSchema`. Mutation (drop `schema`) → RED.
- iter-07 player-presence-only — **GAP (both layers) → FIXED.** The iter-07 tests used a presence-only session with NO path, so they couldn't catch a precedence regression. (a) TS `content-pairs`: added a test where the presence-only flag and a real path/seat COEXIST → must stay presence-only, never a landable player pair. Mutation (swap branch order) → RED. (b) Python `cockpit._content_session_actions` gated the as-player CTA on `player_path` alone → a stray path resurrected a landable CTA next to the "unavailable" note. Added a `not player_presence_only` guard + regression test. Mutation (revert guard) → RED.
- iter-11 EN/IT ack grader — **TEETH ADEQUATE (no change).** Italian-ack landing test + a negative "no acknowledgement" test both present and biting. (The iter-11 bash denominator cross-check drift risk is routed to pass 2.)

**Coverage delta on touched files:**
- stack-snapshot/directus: 99.3% → 99.3% stmts (already maximal; the new test is a data assertion on covered code).
- stack-seeding/contentsession: 93.0% → 93.0%; stack-seeding/seeders: 96.1% → 96.1% (unchanged — pass-1 additions are teeth on already-covered code, not new-line coverage).
- demo-stack/cockpit.py: +1 guarded branch (CTA-suppression), covered by the new regression test. TS content-pairs: +1 spec. (No line-coverage tool wired for Python/TS this pass; mutation-verification is the operative signal per the release thesis.)

**Tests added:**
- iter-08 → demo-stack/tests/test_interview_flag_patch_m232.py: strengthened 1 assertion (paren-anchored positive + negated-scope negative).
- iter-05 → stack-snapshot/directus/directus_test.go: +1 unit (schema/jobsimulation column pin).
- iter-07 → stack-verify/e2e/tests/content-pairs.unit.spec.ts: +1 unit (presence-only precedence over a stray path).
- iter-07 → demo-stack/cockpit.py: +1 guarded branch; demo-stack/tests/test_cockpit.py: +1 regression (CTA suppression with a stray path).

**Bugs surfaced + fixed inline:**
- iter-08 scope test false-passed the player-scope leak (toothless regex) — strengthened (rext 2642d6e).
- iter-07 cockpit rendered a landable as-player CTA for a presence-only session when a stray path was present — guarded (rext 2642d6e).

**Flakes stabilized:** none (none surfaced; the reap suite's own TOCTOU retry is intact).

**Knowledge backfill:** none this pass (the mechanisms are already documented in content-stories-routes.md / session-clone-spec.md / demopatch-spec.md; no protocol-level truth surfaced). Re-evaluate at pass close.

**Out-of-scope observation (routed, not fixed):** demo-stack/tests/test_cockpit.py carries **6 inherited pre-M244 failures** (academy-link + overlay-JS stale tests — the overlay `resetOverlayOnReturn` change landed in rext 04babf8 on 2026-07-15, an ancestor of the M243 end; the tests still assert the removed 30s window / old academy link shape). They predate iter-02 and are outside the iter-02..11 diff scope, and are already tracked by the iter-07 "cockpit 159/6" notation. Not fixed here (would expand harden scope past the iter-diff manifest + root cause is pre-M244 territory — Fate 3 by the fixable-inline boundary).

**Stop condition:** continue-to-next-pass — pass 1 surfaced + fixed 3 toothless/gap items (dimension scan found NEW findings), so the stop condition (delta < 2% AND scan clean) is not yet met; iter-11's bash denominator cross-check (the embedded-Python mirror of buildPairs) still lacks a lockstep-drift regression test.
