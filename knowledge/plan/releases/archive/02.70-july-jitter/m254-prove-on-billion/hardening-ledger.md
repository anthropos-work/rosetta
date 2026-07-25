# Hardening Ledger — M254 prove-on-billion

_The milestone's code-of-record is the **rext** repo (tags on origin); M254's iter commits on
`m254/prove-on-billion` carry the journal/state, and the actual tooling/tests live in
`.agentspace/rosetta-extensions/` (main), tagged + pushed to origin. So the harden **test code** lands in rext
(commits `f9ff4af` + `4c1fd90`, tag `july-jitter-m254-harden-final` on origin — rung-zero); this ledger + the
routed-forward decision land on the milestone branch._

## Pass 1 — 2026-07-25 — final

**Iters hardened this pass:** all milestone-touched code (cumulative-scope final pass). Milestone footprint =
the 4 rext commits: `997272b` (aireadiness demopatch re-point) · `cbe9256` (studio-FCP identity) · `dfdd9bc`
(academy no-node host-robustness) · `4f1409e` (studio-builder Playthrough re-tune + /home networkidle
anti-deadlock).
**Tiks covered since prior pass:** all iters in milestone (no prior harden pass — this is the milestone's
single final-mode pass; 9 tiks iter-02..10 + bootstrap tok).

**Coverage delta on touched files** (coverage-as-finder — these tooling stacks carry no wired line-coverage
tool; delta measured as untested-branch surfaces PINNED, mutation-verified where the file's own doctrine
requires it):
- `playthroughs/e2e/lib/studio-builder-page.ts`: the 5 NEW v0.152.1 unified-entry locators
  (`buildEntryHeading` / `scenarioTextbox` / `advancedModeButton` / `advancedDesignerRendered` /
  `guidedInterviewCanvas`) were shipped by iter-10 UNPINNED (0 unit coverage) → now each shipped matcher is
  captured + asserted (role/text, case-insensitivity, meaningfulness negatives). The 5 old matchers were
  already pinned; the 5 new ones closed the gap.
- `playthroughs/e2e/lib/url-shapes.ts` × `studio-builder-page.ts` (cross-iter integration): the re-tune moved
  the page-object nav `path` to the unified entry `/simulation-builder` (+ added `ADVANCED_PATH`). The entry
  had NO classification pin → now pinned as NEITHER advanced NOR guided, and the route constants ↔ predicates
  are locked in lock-step.
- `demo-stack/patches/app-aireadiness-snapshot-loadmembers.yaml`: `test_patch_inventory` proved it VALID but
  nothing pinned it targets the CONSOLIDATED read-path → the iter-02 drift class (stale
  `internal/workforce/ai_readiness.go` path → silent skip → autoverify green:false) is now regression-fenced
  (path/anchor/replacement asserted against the loaded manifest fields, stale symbols asserted GONE).
- Spec-level `/home` hero-login (skillpath-legacy + aisim-chat-launch + aireadiness-member-done/-progress):
  the iter-10 gate-(h) `waitUntil:'domcontentloaded'` fix was protected ONLY by comments (the existing
  `page-object.unit.spec.ts` guard covers `PageObject.goto`, not the spec-level login) → now a source-scan
  invariant pins it across all current + future `/home` specs.

**Tests added:**
- iter-10 (studio re-tune) → `playthroughs/e2e/tests/studio-builder-locators.unit.spec.ts`: +5 new-locator
  capture + 4 cross-surface-separation = 9 unit.
- iter-10 (cross-iter) → `playthroughs/e2e/tests/url-shapes.unit.spec.ts`: +2 unit (entry classification +
  route-constant single-source pin; +`StudioBuilderPage` import).
- iter-10 (gate-h) → `playthroughs/e2e/tests/home-login-networkidle.unit.spec.ts` (NEW): +3 (invariant scan +
  extractor self-test + offender-path teeth).
- iter-03 (demopatch re-point) → `demo-stack/tests/test_aireadiness_snapshot_loadmembers_m254.py` (NEW): +8
  (7 manifest-shape/re-point + 1 bring-up-wiring).
- **Total: 22 tests (14 TS + 8 Python).**

**Verification:** tsc `--noEmit` clean; full playthroughs offline unit-spec set 110→(post-pass) green; Go
`go test -count=1 ./playthroughs/...` all 4 packages green (ptvalidate validates the re-tuned
`studio-builders.yaml`); demopatch-family + new Python test 88 unittest OK. **Flake gate: 3 consecutive clean
runs** of the new/changed tests (92 TS passed + Python OK, each run).

**Iter-shape carve-outs applied:**
- `cbe9256` (studio-FCP identity) — **tooling-iter, no offline surface.** A 1-line default swap
  (maya-thriving→dan-manager); studio-eligibility is enforced LIVE by studio-desk's `checkEnterpriseAndAdmin`
  and there is no in-repo roster→role map to assert offline. Covered by iter-06's live proof (dan/dana/rae
  reach the studio shell; maya bounced by design). Dimensions 2/3/5/6 = no-op. **No new test — legitimate
  carve-out.**
- `dfdd9bc` (academy no-node) — **tooling-iter, dimension-1 verification.** The fix hardens a TEST; its
  node-free-bindir helper is inline in `_run` (no extractable unit surface). Verified host-robust on a 2nd
  host: `test_missing_node_documents` runs GREEN on this macOS box (billion — the discriminating host with
  `/usr/bin/node` v18 — was the iter-07 live proof). Cross-host green stands; no new code.

**Bugs surfaced + fixed inline:** none fixed inline.

**Bugs surfaced + ROUTED FORWARD (Fate 3):**
- **`test_patch_inventory` RED at HEAD — the demopatch inventory fence has drifted** (`EXPECTED_TOTAL=21` /
  `studio-desk:3`, but `patches/` holds **23** manifests / studio-desk **5**). **Root cause = M253** (`b8969c0`)
  added `studio-desk-shell-first-paint` + `studio-desk-no-thirdparty` without bumping the fence — so it has been
  RED since the M253 tip; **not an M254 change** (M254 only changed the aireadiness patch's CONTENT, not the
  patch count). The fix spans a rext constant AND the `demopatch-spec.md §5` corpus table (they move together
  by the fence's own contract) — i.e. cross-subsystem + sibling-milestone territory + a corpus doc edit, all
  outside a harden pass's rext-test scope. Per the fixable-inline boundary ("root cause in a sibling
  milestone's territory" + "final-mode reveals a regression → surface, don't auto-fix"): routed forward as
  `FIX-M254-h-patch-inventory-drift` (carry-forward.md), fated at close-milestone's deferral audit. **Precise
  fix (verified):** `EXPECTED_TOTAL 21→23`, `EXPECTED_BY_REPO["studio-desk"] 3→5`, + the `§5` table.

**Flakes stabilized:** none newly flaky (the pre-existing `/home` networkidle flake was fixed by iter-10;
this pass PINS that fix so it cannot silently regress).

**Knowledge backfill:** none required — the guarded truths are documented in-line by the test files
themselves (the `home-login-networkidle.unit.spec.ts` header records the spec-level-vs-goto guard distinction;
the demopatch re-point test header records the iter-02 drift class). No protocol/subsystem doc gap opened.

**Stop condition:** continue-to-next-pass — pass-1 landed the bulk; a pass-2 orthogonal sweep (cross-surface
separation + a full dimension re-scan) is needed to compute the coverage delta.

## Pass 2 — 2026-07-25 — final

**Scope:** the same cumulative footprint. Orthogonal-dimension sweep + full dimension re-scan.
**Tests added:** the 4 cross-surface-separation tests (folded into the pass-1 count above; landed in the same
`studio-builder-locators.unit.spec.ts` commit) — the entry / advanced-completion / guided-canvas landmark
matrix (each matcher fires ONLY on its own surface's copy).
**Dimension re-scan (all 6 dimensions × the 4 milestone surfaces):** no further meaningful untested branch.
The url-shape predicates already carry an extensive fuzz/edge grid (the M204/M243 blocks); the demopatch
loader error paths are covered by `test_demopatch`; the new studio locators + the networkidle invariant carry
their own negative/teeth cases; studio-FCP + academy-no-node are carve-outs (above). Coverage delta pass-1→
pass-2 = the 4 separation tests; no NEW surface identified.
**Verification:** studio-builder-locators 15 green; Go fresh-green; demopatch-family 88 OK.
**Stop condition:** continue-to-next-pass — pass-2 still added tests; a confirming clean sweep is owed.

## Pass 3 — 2026-07-25 — final

**Scope:** confirmation sweep — no new tests. Full offline unit-spec set + Go + demopatch-family re-run + the
3-consecutive-clean flake gate.
**Coverage delta:** 0 (no tests added; the dimension re-scan found nothing new across all 4 milestone
surfaces). All suites green; flake gate 3/3 clean.
**Stop condition:** **stabilized** — coverage delta < 2% AND the Phase-2 dimension scan found nothing new. The
milestone-touched surfaces are pinned; the single RED at HEAD is the pre-existing M253 inventory-fence drift,
routed forward (Fate 3), not an M254 regression.
