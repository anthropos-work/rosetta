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

## Pass 5 — 2026-07-23 — incremental

**Iters hardened this pass:** iter-18 (the one gap Pass 4 routed forward).
**Tiks covered since prior pass:** same batch (iters 13–24); Pass 5 continues Pass 4's scope.

**Mutation-verification outcome:**
- iter-18 run-discrete spec-path normalization — **GAP → FIXED.** Pass 4 covered the guards + the offset/port
  mapping but not the three-branch spec-name normalization (`tests/*` pass-through | bare `*.spec.ts` → `tests/`
  prefix | bare name → `tests/<name>.spec.ts`). A wrong prefix is silent poison: playwright matches 0 specs and
  the discrete sweep reports GREEN over nothing. The `npx` stub now echoes its args, so the composed
  `playwright test <SPEC_PATHS…>` invocation is observable; +1 test asserts all three forms resolve + no doubled
  prefix/suffix. Mutation: drop `.spec.ts` from the bare-name branch → RED.

**Coverage delta on touched files:** stack-verify run-discrete.unit.spec.ts 6 → **7** specs; full stack-verify
unit suite 171 → **172** + tsc clean. No line-coverage tool wired (mutation-verify is the signal).

**Tests added:**
- iter-18 → stack-verify/e2e/tests/run-discrete.unit.spec.ts: +1 (spec-path normalization, all three forms).

**Bugs surfaced + fixed inline:** none new.

**Dimension scan (remaining surface):** the scan is now clean for the hardenable in-scope surface. What is left
un-swept is config/live-sweep robustness whose CORE is already covered: `run-discrete.sh`'s `DISCRETE_STUDIO_BASE`
override + the `render-hiring-comparison` per-spec env defaults (`RENDER_EXPECTED_SIMS`/`COVERAGE_RENDER_GATE`) only
exercise over a LIVE `--public-host` remote sweep — a unit test of those would be gold-plating (same call the prior
harden made for the iter-03 `SETTLE_MS` env parse). No further meaningful unit surface across iter-13/15/16/18/23.

**Flakes stabilized:** none surfaced (flake gate at Pass 6).

**Knowledge backfill:** none (the runner's env contract is documented in its own header + coverage-protocol.md).

**Stop condition:** continue-to-next-pass — Pass 5 added 1 test (the scan found one more gap), so a confirmation
pass is needed to measure the delta as < 2% and run the flake gate on the whole batch of newly-added tests.

## Pass 6 — 2026-07-23 — incremental

**Iters hardened this pass:** confirmation pass over the iter-13..24 scope (no new tests written).
**Tiks covered since prior pass:** same batch.

**Dimension scan (the stabilization check):** nothing new to fix. Every in-scope fix with hardenable code is
covered and mutation-verified — iter-13 (settle, teeth confirmed), iter-15 (public demopatch: chain + scope +
ladder-on-real-clone; + the inventory-drift bug), iter-16 (home card-floor behaviour vs the old token marker),
iter-18 (runner guards + offset mapping + spec-path normalization), iter-23 (goto never-networkidle incl. the
real polling surfaces). iter-14/17/20/21/22/24 shipped no hardenable code (proof / live-infra). The only
un-swept surface is live-sweep-only config robustness (gold-plating, per Pass 5).

**Coverage delta on touched files:** 0% vs Pass 5 (no new tests). stack-verify unit suite **172** + tsc clean;
playthroughs unit **3** + tsc clean; demo-stack python **225** green (incl. the reconciled inventory fence).

**Tests added:** none.

**Bugs surfaced + fixed inline:** none.

**Flakes stabilized:** none surfaced. **Flake gate: 3 consecutive clean runs** of every newly-added test —
stack-verify run-discrete (7) + the section-assert academy-home block (4); playthroughs page-object (3); python
test_academy_fs_published_public (19) + the reconciled test_patch_inventory (5). 3/3 green on all.

**Final green:** stack-verify 172 unit specs + tsc clean; playthroughs 3 unit specs + tsc clean; demo-stack
python 225 green + `test_tooling` green (the patch-set enumerator). Both e2e TS projects typecheck clean.

**Session totals (Passes 4–6):** +33 tests (Python +19 academy-public; TS +4 home card-floor, +7 run-discrete,
+3 page-object) + the inventory fence RED→GREEN; 1 bug fixed inline (iter-15 inventory drift); 6 fixes
mutation-verified to bite (1 already-toothed confirmed, 4 coverage gaps + the drift bug fixed). rext HEAD advanced
`dddef18` → `5d96a57` (3 commits); consumption tag `sound-check-m244-content-sweep-robustness` moved + pushed.

**Stop condition:** stabilized — coverage delta < 2% (0% this pass) AND the dimension scan found nothing new;
every in-scope fix with hardenable code mutation-verified to bite, flake gate 3/3 clean.

---

# Final harden — cumulative scope (post gate-met; before close-milestone)

The `--final` cumulative-scope sweep across all milestone-touched code, run after the gate fired 8/8 (iter-27
`closed-fixed`, gate MET). The two incremental sessions above hardened iters 02–24; the NEW hardenable code
since Pass 6 is the run-9/10 fixes: iter-25 `c755370` (up-injected pinned-ref build fix + M217 anchor-preflight,
+3 test_tooling pins) and iter-27 `6feae20` (4 harness locator/timing fixes). iter-26 `closed-no-lift` shipped
NO code (pure live characterization, routed to iter-27) — nothing to harden. As before, M244's fixes' CODE lives
in `rosetta-extensions` (rext, main); the rosetta-side commit is the ledger only.

## Pass 7 — 2026-07-23 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope); NEW work concentrated on
iter-25 + iter-27 (iter-26 = no code). Prior passes' iter-02..24 surface re-confirmed green (see Verification).
**Tiks covered since prior pass:** all iters in milestone (final).

**Thesis probed (anti-toothlessness — the release thesis):** for every fix that shipped a test, MUTATION-VERIFY
— break the subject, confirm the test goes RED, restore; for the fixes that shipped NO unit test (the iter-27
locators), ADD one that captures the SHIPPED matcher (never a re-typed copy) so a source mutation flows straight
into the assertion. Result: iter-25's 3 pins already bite (confirmed) + 1 gap-fill added; iter-27's 3 locators
were unpinned below the live billion sweep → pinned, and one carried a real toothlessness reintroduction, fixed
inline (Fate 1).

**Mutation-verification outcomes (the crux):**
- iter-25 build-scratch pinned-ref (`test_build_scratch_uses_pinned_checkout_not_highest_tag`) — **TEETH
  CONFIRMED (no change).** Live mutation: reverted the build `tag=` line to the old
  `for-each-ref --sort=-v:refname` → the pin RED (AssertionError: the `describe --exact-match HEAD` form absent).
  The task's exact question — "does a test fail if it reverts to highest-tag?" — answered YES.
- iter-25 resolved-ref fetch (`test_scratch_fetches_the_resolved_ref_so_checkout_cannot_fail`) — **TEETH
  CONFIRMED (no change).** Dropped the `git -C "$dst" fetch --quiet "$src" "$tag"` line → RED.
- iter-25 preflight anchor-tag — **GAP → FIXED.** The fix changed BOTH the build-scratch tag AND the M217
  anchor-preflight `_app_tag` (up-injected.sh:1328) — they must resolve the SAME ref or the demopatch-anchor
  preflight validates a DIFFERENT source than the one that ships — but only the BUILD tag was positively pinned;
  the preflight was covered only by the shared `assertNotIn(for-each-ref…)`. New
  `test_preflight_anchor_tag_resolves_the_pinned_checkout_not_highest_tag`. Mutation C: revert ONLY the preflight
  to a NON-for-each-ref wrong form (`git tag --list "v*" | sort -V | tail -1`) → new test RED **while the
  existing build-tag test stayed GREEN** — proving that wrong-form was previously UNCAUGHT and the gap-fill
  closes a real hole.
- iter-27 byTeam locator — **GAP → PINNED.** New `ai-readiness-locators.unit.spec.ts` captures the shipped
  `getByText` regex via a Recorder fake (the `page-object.unit.spec.ts` pattern). Asserts it matches
  "AI Readiness by Team" (+ suffix) and REJECTS the old "AI Readiness by Tag". Mutation D: revert source to
  "by Tag" → RED.
- iter-27 interviewBreakdownPanel scope — **GAP → PINNED.** The panel must scope to the findings CARD (heading
  AND a findings label, terminal `.last()`), not the 24-char heading span the >900-char density check false-
  measured. The unit test asserts: scopes to a `div`, TWO distinct `hasText` filters (heading + a findings
  label, the second NOT a dup of the heading), the AND rejects the heading-only 24-char text and accepts the
  full card body, and the terminal op is `.last()`. Mutation F: revert to the heading-only `getByText(...).first()`
  form → RED.
- iter-27 dueDate year-less matcher — **TOOTHLESS REINTRODUCTION → FIXED INLINE (Fate 1).** iter-27 added
  year-less date shapes to match the real "Due Aug 22" render, but used an unconstrained `\w{3,9}` for the month
  slot — so a due/deadline anchor beside a RELATIVE-DURATION countdown ("Due in 24 hours", "Deadline 5 days
  left", the degraded "Due · 30 days left" with NO date) matched as if a real date had rendered. That is exactly
  the M219 class (a due-anchor + something-date-ish that isn't a real date) the year-BEARING shapes were hardened
  against. The unit test drove it RED on committed source (the hole), then the fix constrains the two year-less
  shapes' month slot to real month names (`(?:jan…dec)[a-z]*`) — the real render "Due Aug 22 · 30 days left"
  still matches (via "Aug 22"), the year-bearing shapes are untouched (their `\d{4}` already anchors them), and
  the countdown-only strings no longer match. Live billion spec unaffected (Aug IS a month; safe without a
  re-drive). Mutation E: revert the month constraint back to `\w{3,9}` → the relative-duration test RED.
- iter-27 pickFirstSkillPath option-gated retry — **LIVE-SWEEP-ONLY (no unit surface, noted honestly).** A
  timing/retry loop (wait for a real `option` node before ArrowDown/Enter; retry ×3 until submit enables) that
  only exercises over a high-latency tailnet — the same class the prior harden judged gold-plating to unit-test
  (iter-03 SETTLE_MS / tailnet retry). Proven by the live 16/16 billion run (flaked once, passed on re-run,
  hardened). No pure-logic surface; no fabricated coverage.

**Cross-iter integration sanity check (final-mode's defining work):** `playthroughs/e2e/lib/ai-readiness-page.ts`
is touched across MULTIPLE iters — iter-23 (`PageObject.goto` never-networkidle, via the base the AI-readiness
surfaces inherit) and iter-27 (the AI-readiness locators). The two unit specs now cover it end-to-end:
`page-object.unit.spec.ts` (Pass 4) pins that `AIReadinessMemberSurface`/`AIReadinessDashboardPage` inherit
`domcontentloaded`; `ai-readiness-locators.unit.spec.ts` (this pass) pins the same two surfaces' locators — so
both the NAVIGATION contract and the LOCATOR contract of the AI-readiness page objects are pinned below the live
sweep. The iter-27 `aireadiness-manager-howwemeasure.spec.ts` (a live spec) now depends on
`interviewBreakdownPanel()`; the unit test pins that method it relies on. No cross-iter regression surfaced.

**Coverage delta on touched files:** no line-coverage tool wired for TS/Python (as all prior passes);
mutation-verification is the operative signal per the release thesis. Concrete additions: **+6 tests** — Python
+1 (preflight anchor gap-fill; BuildScratchPinnedRef 3→4) + TS +5 (ai-readiness-locators: byTeam, panel,
dueDate ×3). playthroughs unit suite 80 → **85** + tsc clean; demo-stack test_tooling 163 → **164**.

**Tests added:**
- iter-25 → demo-stack/tests/test_tooling.py: +1 (preflight anchor-tag pinned-checkout pin).
- iter-27 → playthroughs/e2e/tests/ai-readiness-locators.unit.spec.ts: +5 (byTeam text; interviewBreakdownPanel
  card-scope; dueDate matches-real / rejects-dateless-anchor / rejects-relative-duration).

**Bugs surfaced + fixed inline:**
- iter-27 dueDate year-less matcher matched relative-duration countdowns (a toothlessness reintroduction of the
  M219 class) — constrained the year-less month slot to real month names (rext, this pass). Fate 1.

**Flakes stabilized:** none surfaced. Flake gate at Pass 8.

**Knowledge backfill:** none (the pinned-ref build contract is documented at the fix + CLAUDE.md's tagging-is-not-
publishing note + verification.md rung-zero; the locator discipline is in page-object.ts's header + coverage-
protocol.md. No protocol-level truth surfaced — the mechanisms were already documented.)

**Stop condition:** continue-to-next-pass — Pass 7 added 1 gap-fill + 5 locator pins and fixed 1 inline
toothlessness bug (the dimension scan found NEW findings), so the stop condition (delta < 2% AND scan clean) is
not yet met; a confirmation pass is needed to measure the delta as < 2% and run the flake gate over the whole
batch of newly-added tests.

## Pass 8 — 2026-07-23 — final

**Iters hardened this pass:** cumulative cross-iter sweep (final mode's defining work) — the BROAD
demo-stack python suite, which spans far beyond any single iter's touched files.
**Tiks covered since prior pass:** all iters in milestone (final).

**Broad-suite cross-iter sweep (the flagship final-mode finding):** ran the WHOLE demo-stack python suite
(`unittest discover`, **865 tests**), not just the milestone-touched-file subset the incremental passes ran.
9 failures. Classified each by root cause against the pre-M244 baseline (`2ef5962`, parent of the first M244
rext commit `a77e89c`):
- **1 M244-INTRODUCED → FIXED INLINE (Fate 1).** `test_host_prereqs_m215.py::test_script_uses_the_no_pipe_
  form_not_git_tag_head` asserted the `for-each-ref --sort=-v:refname …` string is PRESENT in up-injected.sh —
  but iter-25 (`c755370`) correctly REPLACED that exact string with `git describe --tags --exact-match HEAD`
  (the pinned-ref fix) on BOTH the build `tag=` and the M217 preflight `_app_tag` lines. The M215 fence went
  RED at iter-25 and stayed red through iter-27 — it directly CONTRADICTED iter-25's own
  `BuildScratchPinnedRef.assertNotIn(for-each-ref)`, a two-tests-in-one-suite contradiction the iter loop never
  reconciled (the live billion sweep never runs the python suite, so it stayed unseen). The M215 test's real
  INVARIANT — the resolution never uses `git tag | head`, which SIGPIPEs on app's ~337 tags under
  `set -o pipefail` and aborts the bring-up — is still satisfied (`describe` is likewise no-pipe). Reconciled
  the "no-pipe form present" assertion to the current `describe` form (the lock-step-fence pattern, parallel to
  Pass 4's inventory-count drift), keeping the anti-SIGPIPE negatives + the M215 anchor. Mutation-verified:
  inject a `git tag …| head` piped form → the fence RED; remove the `describe` form → the no-pipe-form pin RED.
  The sibling `test_for_each_ref_returns_highest_tag_without_141` (a standalone behavioural proof that
  for-each-ref is SIGPIPE-safe) is untouched and green — the M215 lesson stays documented.
- **8 PRE-M244 INHERITED (Fate 3, out of milestone scope — SUT untouched by M244):**
  - 6 in `test_cockpit.py` — the ledger's already-documented inherited set: academy-link
    (`test_root_serves_academy_link`, `test_academy_link_renders_per_hero_when_base_set`,
    `test_render_academy_entry_fields_are_escaped`, `test_render_defaults_academy_path_persona_label_when_absent`)
    + overlay (`test_inflight_window_is_30s`, `test_localstorage_access_is_guarded`) — the removed-30s-window /
    old-academy-link stale tests (rext `04babf8`, 2026-07-15, ancestor of the M243 end).
  - `test_host_prereqs_m215.py::test_public_host_emits_per_port_off_for_all_browser_ports` — an ASSERTION drift
    (the reset omits port **13001** = the hiring app 3001, M223/v2.4). The tailscale-serve reset generator is
    NOT in the M244 source footprint and NO M244 commit touched this test file → its state is identical pre/post
    M244 → pre-M244 drift between the test's `_UI_PORTS` and the generator. Not M244's.
  - `test_purge.py::test_purges_container_owned_0700_data_THE_BUG` — a docker-integration test (runs
    `docker run alpine:3` to stage a hostile 0700 dir); environmental / pre-M244, purge source untouched by M244.

  Decisive attribution: `git log 2ef5962..HEAD` shows NO M244 commit touched `test_host_prereqs_m215.py`,
  `test_purge.py`, the tailscale-serve reset generator, or any purge/teardown source — so all 8 are pre-M244,
  routed Fate 3 (their root causes are sibling-milestone / environmental territory, per the fixable-inline
  boundary). The milestone-touched python subset (243 tests: test_tooling + the academy-fs/patch-inventory/
  aireadiness-flag/interview-flag suites) is fully GREEN.

**Coverage delta on touched files:** +0 net-new tests (the M215 change is a 1-test RECONCILIATION, red→green,
not an addition). mutation-verification is the operative signal (as all prior passes).

**Tests added:** none (1 reconciled: test_host_prereqs_m215.py no-pipe fence).

**Bugs surfaced + fixed inline:**
- iter-25 left the M215 no-pipe fence asserting the superseded `for-each-ref` string (an M244-introduced stale
  fence contradicting iter-25's own test) — reconciled to the `describe` form (rext 498b1a5). Fate 1.

**Flakes stabilized:** none surfaced. Flake gate at Pass 9.

**Knowledge backfill:** none (the no-pipe/SIGPIPE contract is documented at the fix in up-injected.sh + the M215
test's own reconciled comment; no protocol-level truth surfaced).

**Stop condition:** continue-to-next-pass — Pass 8 surfaced + fixed 1 M244-introduced stale fence (the dimension
scan found a NEW finding), so a confirmation pass is needed to measure the delta as 0 and run the flake gate over
the whole batch of touched tests.

## Pass 9 — 2026-07-23 — final

**Iters hardened this pass:** confirmation pass over the full final cumulative scope (no new tests written).
**Tiks covered since prior pass:** all iters in milestone (final).

**Dimension scan (the stabilization check):** nothing new to fix. The NEW hardenable code since the 2nd
incremental harden (iter-25 `c755370`, iter-27 `6feae20`; iter-26 = no code) is fully covered and
mutation-verified: iter-25 (build-scratch tag / resolved-ref fetch / **preflight anchor gap-fill** — all bite;
the M215 sibling fence reconciled), iter-27 (byTeam text / interviewBreakdownPanel card-scope / dueDate
year-less-month-constrained — all bite; the dueDate toothlessness hole closed; pickFirstSkillPath live-only,
noted). The only un-swept surface is the iter-27 `pickFirstSkillPath` retry timing (live-tailnet-only, no
pure-logic surface — the same gold-plating call the prior harden made for iter-03's SETTLE_MS / tailnet retry)
and the 8 pre-M244 inherited failures (Fate 3, out of milestone scope).

**Cross-iter integration (final):** `playthroughs/e2e/lib/ai-readiness-page.ts` (iter-23 goto + iter-27
locators) is pinned end-to-end by `page-object.unit.spec.ts` (navigation contract) + `ai-readiness-locators.
unit.spec.ts` (locator contract); `demo-stack/up-injected.sh` (iter-25) is pinned by BOTH `test_tooling.
BuildScratchPinnedRef` (the describe form + fetch + preflight) AND the reconciled M215 no-pipe fence (the
anti-SIGPIPE invariant) — the two now agree instead of contradicting. No cross-iter regression remains.

**Coverage delta on touched files:** 0% vs Pass 8 (no new tests). playthroughs unit **85** + tsc clean;
stack-verify unit **172** + tsc clean; demo-stack milestone-touched python subset **243** green + the reconciled
M215 no-pipe test green.

**Tests added:** none.

**Bugs surfaced + fixed inline:** none.

**Flakes stabilized:** none surfaced. **Flake gate: 3 consecutive clean runs** of every touched test —
playthroughs ai-readiness-locators (5), demo-stack BuildScratchPinnedRef (4), demo-stack M215 no-pipe (1
reconciled). 3/3 green on all.

**Final green:** playthroughs 85 unit specs + tsc clean; stack-verify 172 unit specs + tsc clean; demo-stack
milestone-touched subset 243 + the reconciled M215 fence green. Both e2e TS projects typecheck clean.

**Session totals (Passes 7–9, the final harden):** +6 net-new tests (Python +1 preflight-anchor gap-fill; TS +5
ai-readiness locators) + 2 inline fixes (iter-27 dueDate month-constraint toothlessness; the iter-25 M215
no-pipe fence reconciliation); 9 fixes/pins mutation-verified to bite (iter-25 build-tag/fetch confirmed +
preflight gap-fill + M215 fence; iter-27 byTeam/panel/dueDate); the full 865-test broad python suite triaged
(1 M244-introduced stale fence fixed, 8 pre-M244 inherited routed Fate 3). rext HEAD advanced `6feae20` →
`498b1a5` (2 commits); consumption tag `sound-check-m244-content-sweep-robustness` moved + force-pushed to
origin (peels to `498b1a5`) + main pushed.

**Stop condition:** stabilized — coverage delta < 2% (0% this pass) AND the dimension scan found nothing new;
every in-scope M244 fix with hardenable code mutation-verified to bite, the one M244-introduced stale fence
reconciled + mutation-verified, flake gate 3/3 clean. Final-mode entry present (satisfies close-milestone's
iterative-milestone gate).
