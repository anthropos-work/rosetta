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

## Pass 2 — 2026-07-22 — incremental

**Iters hardened this pass:** iter-07, iter-11 (the two gaps pass 1 routed forward).
**Tiks covered since prior pass:** same batch (iters 02–11); pass 2 continues pass 1's scope.

**Mutation-verification outcomes:**
- iter-11 bash denominator cross-check (`run-content-stories.sh` embedded `_PIN_PY`) — **GAP → FIXED.** The embedded Python that cross-checks the denominator is a hand-copy of `buildPairs()` and hand-copies DRIFT — this WAS the iter-11 bug (the presence-only branch wasn't mirrored → over-counted 49 vs 47). The runtime cross-check catches it only during a LIVE sweep. New `tests/content-denominator.unit.spec.ts` EXTRACTS the real `_PIN_PY` out of the shell heredoc (never a re-typed copy) and runs it: a presence-only cell counts as its MANAGER pair only, an over-count exit-2s, a 0/absent pin is refused. Mutation: strip the `not player_presence_only` guard from the embedded program → count becomes 2 → RED.
- iter-07 disclosure invariant (ERROR path) — **GAP → FIXED.** `contentsession.Validate()` rejects `player_result_unavailable` with no `player_unavailable_reason`; iter-07 added the rule but nothing drove it. New `TestValidate_PlayerResultUnavailableRequiresReason` exercises both sides. Mutation: remove the check → RED.

**Coverage delta on touched files:**
- stack-seeding/contentsession: 93.0% → **93.6%** stmts (+0.6% — the Validate error-path branch is now exercised).
- stack-seeding/seeders 96.1%, stack-snapshot/directus 99.3% — unchanged.
- stack-verify/e2e: +3 specs (content-denominator.unit.spec.ts); 56 → 59 content unit specs; tsc clean.

**Tests added:**
- iter-11 → stack-verify/e2e/tests/content-denominator.unit.spec.ts: +3 unit (presence-only excluded; over-count exit-2; 0/absent pin refused).
- iter-07 → stack-seeding/contentsession/contentsession_test.go: +1 unit (Validate reason-required, both sides).

**Bugs surfaced + fixed inline:** none new (both items were coverage gaps on correct code — the fix is the test that pins it).

**Flakes stabilized:** none.

**Knowledge backfill:** none (the denominator cross-check + disclosure invariant are already documented in coverage-protocol.md / session-clone-spec.md; no protocol-level truth surfaced).

**Stop condition:** continue-to-next-pass — pass 2 fixed 2 more gaps (dimension scan still found items), so not yet stabilized; a confirmation pass is needed to measure the delta and verify the scan is clean.

## Pass 3 — 2026-07-22 — incremental

**Iters hardened this pass:** confirmation pass over the iter-02..11 scope (no new tests written).
**Tiks covered since prior pass:** same batch.

**Dimension scan (the stabilization check):** nothing new to fix. The remaining un-swept iter surface is iter-03's `content-result-page.ts` (`SETTLE_MS` env-override) + the per-pair warm-retry in `content-stories.spec.ts` — both are config/live-sweep robustness whose CORE (`settle()`) is already covered by the content-result-page settle tests; a unit test of the env parse or the tailnet retry would be gold-plating (the retry only exercises over a live `--host` sweep). The six named probes + the two pass-2 gaps + the iter-07 error path are all covered and mutation-verified.

**Field validation (unplanned):** this box is running a LIVE ambient demo-1 presenter cockpit on :17700 (pid 42277, detached since 2026-07-21). That is exactly the ambient infrastructure iter-10's test-isolation fix exists to coexist with — its live presence independently confirms the fix's necessity (and its teeth were mutation-verified against a held :17700 in pass 1). Left running (legitimate; never touched).

**Coverage delta on touched files:** 0% vs pass 2 (no new tests) — stack-seeding/contentsession 93.6%, seeders 96.1%, stack-snapshot/directus 99.3%; TS 63 content unit specs + tsc clean.

**Tests added:** none.

**Bugs surfaced + fixed inline:** none.

**Flakes stabilized:** none surfaced. Flake gate: 3 consecutive clean runs of every newly-added test (Go schema-col + Validate; TS precedence + 3× denominator; Python strengthened-scope + cockpit-guard) — 3/3 green on all.

**Final green:** Go stack-seeding + stack-snapshot green + vet clean; 63 TS content unit specs + tsc clean; demo-stack reap 41/41, interview-flag 12/12, cockpit 160 pass. The 6 cockpit failures are the inherited pre-M244 academy/overlay stale tests (rext 04babf8, the "159/6" notation) — out of the iter-02..11 diff scope, unchanged by this harden.

**Stop condition:** stabilized — coverage delta < 2% (0% this pass) AND the dimension scan found nothing new; all six named-probe fixes mutation-verified to bite (3 already-toothed confirmed, 1 toothless test + 4 gaps fixed across passes 1–2), flake gate clean.

---

# Second incremental harden — iters 13–24 (toks 12/19 excluded)

The 2nd incremental harden, scoped to the 11 tik iters closed since Pass 3 (last-harden terminating
commit `51ebd76`). M244's fixes' CODE lives in `rosetta-extensions` (rext, main), so the harden deepens
the UNIT/mutation coverage of those fixes there and consumes them via the pinned tag
`sound-check-m244-content-sweep-robustness`; the rosetta-side commits are the ledger + one paired doc fix.

The hardenable code surface is 5 rext commits: iter-13 `2bb0473` (shape-aware settle), iter-15 `8391843`
(academy anon-view public demopatch), iter-16 `2a71e08` (coverage marker recalibration), iter-18 `6aacc32`
(run-discrete.sh runner), iter-23 `dddef18` (PageObject.goto networkidle→domcontentloaded). iter-14, iter-17,
iter-20, iter-21, iter-22, iter-24 landed **no new hardenable code** — they are live-prove / live-infra iters
(gate-c measure, discrete-spec mapping [became iter-18], gate-h proof on the pre-existing latency harness, the
3 drift-carry burn-ins, and the BURNIN dev-up cycle); noted honestly, no coverage fabricated.

## Pass 4 — 2026-07-23 — incremental

**Iters hardened this pass:** iter-13, iter-15, iter-16, iter-18, iter-23 (the 5 with hardenable code).
**Tiks covered since prior pass:** 11 (iters 13–24; toks 12/19 excluded).

**Thesis probed (anti-toothlessness):** MUTATION-VERIFY each fix — break the subject, confirm the test goes
RED, restore. Result: iter-13 already toothed (confirmed, no change); iter-15/16/18/23 each had a UNIT-coverage
GAP (the fix was correct but nothing pinned it below the live billion sweep) → all filled + mutation-verified;
and iter-15 surfaced a real **bug** (the inventory fence RED on committed main), fixed inline (Fate 1).

**Mutation-verification outcomes (the crux):**
- iter-13 shape-aware settle — **TEETH CONFIRMED (no change).** Neutralized `contentReady` (→ always true, the
  "settle early-exits on the nav-shell chrome plateau" mutation) → exactly **2 of 4** shape-aware settle tests
  RED ("does NOT early-exit on the 128-char plateau" + "ack-never-paints polls to the deadline"), matching the
  iter-13 commit's own claim. The blast-radius pin (player-scored settles ack-blind) + the fast-path stay green.
- iter-15 academy anon-view public demopatch — **GAP (ZERO coverage) → FIXED.** The `-public` variant had **no
  test at all** (the sibling test never touched it). New `demo-stack/tests/test_academy_fs_published_public.py`
  (19): THE CHAIN (`public.pre_sha256 == sibling.post_sha256`; distinct `new Set()`/`eids` anchors), THE
  PUBLIC-ONLY SCOPE (`new Set()` never `eids` → no tenant leak; draft-strip → no chip), the apply/revert LADDER
  (the pinned clone IS present here, so it ran for real: chain realized, roundtrip, idempotent, the **chain-order
  guard REFUSES** applying `-public` on a pristine/unchained file, drift-refuse), and the ant-academy.sh ordering
  (apply AFTER / revert BEFORE FSPUB). Mutation: corrupt `pre_sha256` → TestChain RED; swap `new Set()`→`eids` in
  the replacement → TestScope RED (2 assertions).
- iter-15 inventory drift — **BUG (RED on committed main) → FIXED INLINE.** iter-15 added the **16th** patch
  manifest (`academy-fs-published-public`) but never bumped the mirrored inventory count, so
  `test_patch_inventory` (`EXPECTED_TOTAL`/`EXPECTED_BY_REPO`) had been RED since dddef18 (16≠15; ant-academy
  4≠3). This is exactly the drift the fence exists to catch; the iter loop's per-symptom tests never swept it.
  Reconciled the rext constants (16 / 4 ant-academy) **and** `corpus/ops/demo/demopatch-spec.md` §5 (the fence
  mandates the two move together) — header count, apply-vehicle table, reconciliation blockquote, live-fence
  parenthetical, current-total line, + a new inventory-table row. The fence is the regression pin (RED→GREEN).
- iter-16 academy-home card-floor marker — **GAP (behaviour un-pinned) → FIXED.** The iter-16 unit test pinned
  the descriptor SHAPE only; nothing proved the recalibrated `ANT_ACADEMY_HOME_SECTION` is *stronger* — i.e. that
  it FAILs a page the OLD token marker PASSED. New section-assert block (4) runs the SHIPPED descriptor through
  the SHIPPED `assertSection`: a 0-card "AI Academy" header FAILs (the exact gate-(d) false-empty), a real 12-card
  grid PASSES, an 11-card thin catalog FAILs (the floor is a real threshold), a wordmark-less grid FAILs (the
  both-kind text half). Mutation: revert the descriptor to the old `{kind:'text', mustInclude:['AI Academy']}`
  form → **3 of 4** RED, incl. the load-bearing 0-card false-empty flipping to a (wrong) PASS.
- iter-18 run-discrete.sh runner — **GAP (ZERO coverage) → FIXED.** The runner had no test. New
  `run-discrete.unit.spec.ts` (6, subprocess-driven — no re-implementation): the N-integer guard (non-numeric →
  exit 2, the wrong-stack guard the runner's own comment calls out), the SCHEME guard, and the OFFSET/port
  composition (N*10000 + per-surface ports, npm/npx PATH-stubbed to read the banner without a real browser run).
  Mutation: `3000 + OFFSET`→`3000 + N` → mapping test RED; neutralize the N-guard case pattern → guard test RED.
- iter-23 PageObject.goto — **GAP (doctrine un-pinned) → FIXED.** Only a comment stood between a revert to
  `networkidle` and the tailnet re-deadlock iter-23 already paid for once. New playthroughs
  `page-object.unit.spec.ts` (3): the base goto passes `domcontentloaded`, NEVER `networkidle`, and composes the
  URL — plus the REAL AI-readiness polling surfaces (`/home`, `/ai-readiness`, the ones that deadlocked) inherit
  it. Mutation: revert goto to `networkidle` → 2 RED.

**Coverage delta on touched files:** no line-coverage tool wired for TS/Python this batch (as the prior passes);
mutation-verification is the operative signal per the release thesis. Concrete additions: **+32 tests** — Python
+19 (academy public demopatch) + inventory fence RED→GREEN; TS +4 (section-assert home block) +6 (run-discrete)
+3 (playthroughs page-object). Full stack-verify unit suite 165→**171** + tsc clean; demo-stack python **225**
green (was 2 RED on the inventory fence).

**Tests added:**
- iter-15 → demo-stack/tests/test_academy_fs_published_public.py: +19 (manifest/chain/scope/ladder/wiring).
- iter-15 → demo-stack/tests/test_patch_inventory.py + corpus/ops/demo/demopatch-spec.md §5: inventory reconcile.
- iter-16 → stack-verify/e2e/tests/section-assert.unit.spec.ts: +4 (shipped-descriptor home card-floor behaviour).
- iter-18 → stack-verify/e2e/tests/run-discrete.unit.spec.ts: +6 (guards + offset/port mapping).
- iter-23 → playthroughs/e2e/tests/page-object.unit.spec.ts: +3 (goto never-networkidle contract).

**Bugs surfaced + fixed inline:**
- iter-15 inventory count drift: `test_patch_inventory` RED on committed main (16 patches, doc/constants said 15).
  Reconciled rext constants + demopatch-spec.md §5 (rext a0eb684 + rosetta 770b595). Fate 1.

**Flakes stabilized:** none surfaced (flake gate at Pass 6).

**Out-of-scope observation (routed, not fixed):** `playthroughs/e2e/lib/skill-path-page.ts::gotoPath` and
`simulation-page.ts::gotoSim` STILL use a blocking `waitUntil:'networkidle'` — the same class iter-23 fixed on
the base goto. They are OUTSIDE the iter-23 diff (page-object.ts only), and iter-23's live billion re-drive did
NOT find them deadlocking (skillpath-legacy + aisim-chat-launch were among the 12/16 GREEN), so changing them
would be speculative scope-expansion that risks currently-green specs. Noted for a future iter (Fate 3 by the
fixable-inline boundary), not fixed here. (The activity-dashboard/assignments/profile `waitForLoadState(
'networkidle').catch(()=>{})` calls are NON-blocking — the `.catch` bounds them — so they are safe by construction.)

**Knowledge backfill:** demopatch-spec.md §5 updated as part of the Fate-1 inventory fix (the lock-step fence
demands it). No protocol-level truth surfaced.

**Stop condition:** continue-to-next-pass — Pass 4 filled 4 coverage gaps + fixed 1 inline bug across 5 iters
(the dimension scan found NEW findings), so the stop condition (delta < 2% AND scan clean) is not yet met; the
run-discrete spec-path normalization branch (`tests/*` vs `*.spec.ts` vs bare) is still un-swept → Pass 5.
