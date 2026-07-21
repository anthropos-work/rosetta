# M237 — Progress

Section milestone. Checklist from the roadmap In-list. **All sections landed (Fate-1).**

## Sections

- [x] **Clone-freshness fix** (`ensure-clones.sh`) — fetch-verified freshness assertion (rc-checked, never suppressed-stderr) + opt-in advance-to-`origin/main`-or-pinned-tag (`DEMO_ADVANCE_CLONES`) + a real 7-state pin model (`clones.pin.json`; pinned vs stale-by-neglect distinguishable in `clones.lock.json`) + `DEMO_FRESHNESS_STRICT`. Tests `TestCloneFreshnessM237` (incl. 12-vs-202 anti-regression); shellcheck clean; demo-knob-guard green.
- [x] **F-M236-CLOSE-2** — R1 pristine sweep now globs `patches/*/*.yaml` (all 14), not the hard-coded 3. Tests `TestR1SweepM237`; harness refactored into `EnsureClonesHarness` mixin (no triple-run).
- [x] **Fresh-clone demo on billion** — re-pinned rext to `sound-check-m237-clean-stage`; §1 freshness + §2 R1(14) dogfooded live; verified-current clones (frontend 0-behind — the "202-behind" premise REFUTED).
- [x] **Confirmed-defect ledger** (decisions.md) — #1 menu **RESOLVED** (hierarchical for managers) · #4 library **does-not-reproduce** (populated, 0 errors) → M239 re-scoped · #2 academy language **SURVIVES** (`/it` 404 on the 5-behind empty academy) → M238. Verified via the e2e cockpit-login harness from the tailnet presenter vantage.
- [x] **Delivers** — `corpus/ops/rosetta_demo.md` (clone-freshness mechanism + 202-refutation correction) + `corpus/ops/demo/demopatch-spec.md` (R1 all-14) + `corpus/ops/demo/demo-up-defaults.md` (the two new knobs).

## rext code-of-record
Tag `sound-check-m237-clean-stage` (origin/main + tag pushed). Commits `8661444` (§1) + `7847473` (§2).

## Rosetta commits (branch `m237/clean-stage`)
`d6aadf1` (knob docs) · `c19c7a6` (ledger) · `230e247` (Delivers docs).

## M237: Hardening

Scope: the two M237-touched rext files — `demo-stack/ensure-clones.sh` (phases d3/e/f) +
`demo-stack/tests/test_tooling.py`. Coverage model = **explicit branch/path enumeration** of the
freshness-assertion + pin-state + advance + R1-sweep logic (no `kcov`/`bashcov`: neither instruments
bash on macOS, and this suite has always used branch-mapping — see D-HARDEN-1). Harden commits land in
**rosetta-extensions** (separate repo): `ed9a6e1` (Pass 1) · `2ef8f43` (Pass 2) · `533c489` (Pass 3).

### Pass 1 — 2026-07-21 (rext `ed9a6e1`)
**Priority — mutation-verified the 12-vs-202 anti-regression has teeth:** removed the `_fetch_ok` gate on
the behind-count (the exact defect: trust a count off a failed fetch) → `test_fetch_failure_…never_fabricated`
went RED (`behind=999` fabricated); restored → green.
**Branch coverage (phase d3/e):** extended `EnsureClonesHarness` with 4 default-preserving params
(`head_ref`, `revlist_rc`, `demopatch_executable`, `demopatch_rc`) + 6 tests: **pinned-detached** (the 7th
pin state — detached HEAD, not stale, behind stays null), **strict-abort on pin-drift** (the 3rd
`_fresh_problems` class the `DEMO_FRESHNESS_STRICT=1` fatal gate covers), **fetch-ok-but-uncountable**
(behind=null, fetch_ok stays true — honesty invariant on the fetch-ok side), and `DEMO_ADVANCE_CLONES`
`=1`/`=pinned-no-file`/`=<garbage>`.

### Pass 2 — 2026-07-21 (rext `2ef8f43`)
**Bug fixed inline:** the "demopatch not found" branch logged `skipping R1/R2`, but **R2 (no-push) + R1b run
unconditionally after that block** — a log that claims R2 was skipped while R2 ran is exactly the
confident-wrong-log class this release kills. Message now names only R1. Pinned by
`test_r1_skipped_when_demopatch_not_executable_but_r2_still_runs` (mutation-verified in Pass 3).
**Branch coverage (phase f + d3/e):** +5 tests — R1 refusing-manifest (non-fatal, sweep continues), R1 empty
patches (`swept 0`, literal-glob guard), R1 demopatch-absent, advance-pinned-uncloned-repo, and a **multi-repo
mixed-state** run (stale-by-neglect + pinned + pin-drift in ONE run — the `_fresh_problems` accumulator counts
exactly the 2 problem repos).

### Pass 3 — 2026-07-21 (rext `533c489`)
**Harness-fidelity bug fixed inline (this release's theme, turned inward):** the sandbox constrained `PATH`
to a stub bindir but never symlinked `grep`/`sed` — so phase **(c-pre)** (the repos.yml stub-sweep, a
`grep … | sed …`) printed `grep: command not found` and ran **INERT in every ensure-clones test**. A sandbox
that "runs the script" while a whole phase proves nothing is the M237 defect class itself. Wired `grep`+`sed`;
the `cms_dir` fixture now carries `.git` (a real clone — a bare no-.git cms is correctly swept by the now-live
c-pre). Locked by a new c-pre sweep test.
**Branch coverage:** +3 tests (+`make_pull_rc` param + git `--verify` sha-resolution) — advance-make-pull-nonzero
warn, **pin-by-exact-sha** match (distinct from pin-by-ref-name), and the c-pre sweep. Also mutation-verified
the Pass-2 honesty fix (old `skipping R1/R2` wording → test red).

**Knowledge backfill:** no rosetta-corpus doc edit warranted — the findings are rext-internal (test-infra
fidelity + a log-string honesty fix) and edge-semantics recorded in `decisions.md` D-HARDEN-2. The
`rosetta_demo.md` §"Clone freshness" narrative + the 7-state model it documents remain accurate.

### Test tally (milestone-touched, rext `demo-stack/tests/test_tooling.py`)
`TestCloneFreshnessM237` 12 → **22** · `TestR1SweepM237` 4 → **7** · `TestEnsureClonesFunctional` 16 → **17**
(the c-pre sweep). +14 tests on the ensure-clones surface. Full `test_tooling.py` 146 → **160**; shellcheck +
demo-knob-guard fences green.

### Stop condition
Stopped at **3 passes**: the core surface (all 7 pin states, both fetch-honesty sides, all 3 strict classes,
every `DEMO_ADVANCE_CLONES` arm, every R1 error path, multi-repo aggregation) is covered and mutation-verified
twice; the remainder is defensive `|| true` / `2>/dev/null` branches (pin.json parse-error, per-repo
checkout-fail warn, dirty-stash, `ref=unknown`) that are non-fatal by construction — testing them would need
stub surgery for near-zero real-world value (the "shallow test to bump a number" the skill forbids). Recorded,
not shallow-tested. `/developer-kit:close-milestone` Phase 4 re-audits independently.
