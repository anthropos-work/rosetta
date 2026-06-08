# M16 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [x] **Push the stranded fixes** — pushed `547de17`+`ed72e94` (+the M16 doc commits) to `origin`; tagged `dress-rehearsal-m16` (local-only `stack-party-devpath-fix` superseded, not deleted); re-consumed `stack-demo/rosetta-extensions`; consumed copy matches.
- [x] **Stack-core rename migration** — `stack-dev` is the documented default everywhere; `anthropos-dev` demoted to the single intentional "legacy alias" (the back-compat fallback line in the scripts + the `clone_repos.py` help).
- [x] **Prose sweep** — `demo-stack/README.md`, `demo-stack/GUIDE.md`, `dev-stack/README.md`, `stack-core/gen_override.py` docstring → `stack-dev/platform`.
- [x] **GUIDE.md header truth** — remote-exists / 13-tests / `/stack-list` / v1.3 (was: no-remote / 78 / `/demo-status` / v1.1·M3).
- [x] **pytest doc fix** — `pytest tests/ -v` + the 3.11/3.12 note (`demo-stack/GUIDE.md`).
- [x] **`rosetta-extensions/knowledge/` KB** — version-jump expectation + per-milestone tag scheme noted (`knowledge/README.md`, ISSUE-5).
- [x] **rosetta corpus** — `corpus/` had 0 stray `anthropos-dev` (sweep = verify no-op); consolidated `stack-dev` layout + back-compat note added to `corpus/ops/rosetta_demo.md` (the doc itself = the cross-link anchor).

## Verification
- [x] `bash -n` / `py_compile` / shellcheck clean on every touched script (4 shell scripts + `gen_override.py`).
- [x] pytest suite green via the documented invocation — `pytest tests/ -v` → **13/13**.
- [x] `grep -rn 'anthropos-dev'` shows only the intentional back-compat-alias mentions (5 script fallback lines + 1 help-text in extensions; the explanatory note in corpus).
- [x] `origin` and the consumed per-stack copy agree at the new tag — origin tag→commit = authoring HEAD = consumed HEAD = `44edc09`.

## Notes
- Extensions-side work lands as commits in the SEPARATE `.agentspace/rosetta-extensions` repo (gitignored from rosetta) — see spec-notes "Publish result". The rosetta `m16/land-fixes` branch carries only the corpus note + these tracking docs.
- PR review surfaced 5 same-class stale facts in `demo-stack/README.md` (missed by the first sweep — they used `anthropos-demo/`); all landed Fate-1 in extensions commit `44edc09`.
- Behavior/idempotency work (ISSUE-11/14) + frontend (ISSUE-8/9) are M17+ — not pulled in.

## M16: Hardening

### Pass 1 — 2026-06-08
**Scope manifest (the milestone-touched code surface):**
- rosetta `m16/land-fixes` diff vs `release/01.3b-dress-rehearsal`: `corpus/ops/rosetta_demo.md` (prose) + the 5 tracking docs + `state.md` — **0 testable code** in the rosetta tree.
- The testable surface is the SEPARATE `.agentspace/rosetta-extensions` repo (main, pushed). M16's functional changes (a31d70b..44edc09): 6 files — `demo-stack/{up-injected.sh, migrate-demo.sh, rosetta-demo, lib/clone_repos.py}`, `dev-stack/dev-stack`, `stack-core/gen_override.py` (the `stack-dev`-default rename + the migrate-race `|| echo 0`). All other touched files are docs.
- Existing co-located tests: `demo-stack/tests/test_tooling.py` (13 at start), `stack-core/tests/`, `stack-injection/tests/`, `dev-stack/tests/`. No untested *new* functions — M16 added no runtime logic.

**Coverage delta (milestone-touched files):** N/A in the line/branch sense — M16 is a docs/publish/rename milestone (shell scripts + prose, no new runtime logic). The meaningful metric is **drift-fence coverage of M16's two functional changes + the corrected doc facts**; both fix-commits (`547de17` rename, `ed72e94` migrate-race) and the GUIDE truth-up are now fenced (0 fences → 4 guard tests this pass).

**Tests added** (extensions `74b53eb`, pushed to `origin/main`):
- `demo-stack/tests/test_tooling.py` — `TestRenameDrift` (3): every workspace-resolving functional script scanned repo-wide; no UNMARKED `anthropos-dev` (the single intentional legacy alias must carry a legacy/fallback/rename marker on its line), `stack-dev` must lead `anthropos-dev` in every resolver, guarded-file-list must exist. `TestGuideDocTruth` (2): the GUIDE's advertised test count pinned to the suite's live collection + the documented `pytest tests/ -v` entrypoint (not `python3 -m pytest`).
- The count guard immediately caught its own +5 drift (13→18); GUIDE updated to match — a live demonstration the fence fails on real drift.

**Bugs fixed inline:** none — M16 introduced no defects (the two fixes were pre-applied + published in build; this pass fences them).

**Flakes stabilized:** none seen (guards are deterministic static-file reads).

**Knowledge backfill:** none warranted — the rename contract (M16-D2) + the migrate-race fix are already documented in `decisions.md`/`spec-notes.md` and the corpus `rosetta_demo.md` layout note; the guards encode that contract, they don't surface new system truths. Recorded the question was asked.

### Pass 2 — 2026-06-08
**Tests added** (extensions `aabbf74`, pushed to `origin/main`):
- `demo-stack/tests/test_tooling.py` — `TestMigrateRaceGuard` (3): the ISSUE-7 regression fence. Asserts `migrate-demo.sh` runs under `set -euo pipefail` (the precondition that makes the fence necessary), the `casbin_rules` COUNT carries the `|| echo 0` set-e resilience, and the `casbin_rules empty?` guard around the one-shot `init_policy.sql` is intact. The LIVE docker race test is routed to M17 (Fate 2 — M17 owns idempotency/race + the harness); this is the static half. Proven via negative case (removing `|| echo 0` fails the guard).
- GUIDE count 18→21 + `TestGuideDocTruth` class-list updated in lockstep.

**Bugs fixed inline:** none. **Flakes stabilized:** none (8 new guard tests, 3/3 consecutive clean runs).

**Knowledge backfill:** none warranted — same rationale as Pass 1.

### Stop condition
**Stopped after Pass 2.** The six-dimension scan found nothing new worth adding: M16's two functional changes (rename, migrate-race) are both fenced; the `git log --grep=fix` surfaced exactly those two fix-commits, both now have regression fences; no parsers/untyped-boundaries (fuzzing N/A); no perf-critical paths (benchmarks N/A); no new public functions. Coverage-delta in the line sense is N/A for a docs/rename milestone. No flakes (3/3 clean). Per the M16 calibration, a rename/docs milestone rewards **contract guards over raw test counts** — that target is met, with no test bloat.

## M16: Hardening — totals
- demo-stack suite: **13 → 21** (+8 guard tests across 3 new classes); full extensions suite **174 → 182**, all green.
- Extensions harden commits (on `main`, pushed to `origin`): `74b53eb` (rename-drift + doc-truth guards), `aabbf74` (migrate-race fence). Tag `dress-rehearsal-m16` deliberately left at `44edc09` — `/developer-kit:close-milestone` reconciles it to final HEAD; per-stack clone (`stack-demo/rosetta-extensions`) fetched the new objects, checkout stays pinned at the tag until reconcile.
- rosetta `m16/land-fixes` branch: only this `progress.md` update (no testable rosetta code).
