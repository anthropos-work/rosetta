# M1b — progress

B-milestone: automation/config over M0 (`alignctl dna diff` + `alignctl run --gate`). 2 sections — both DONE.

## S1 — Drift-check + CI gate — DONE
- [x] `clerkenstein/scripts/gate.sh` — build runner + a **built `alignctl` binary** (exact exit code, not `go run`-squashed) + `alignctl run --gate-overall 95 --gate-critical 100`; exit 0 met / 2 regressed. Parameterized `ALIGN_DIR`.
- [x] `clerkenstein/scripts/drift-check.sh` — `alignctl dna diff` + gate.sh; exit 0 none / 1 DNA moved / 2 gate regressed (abspath-resolves the DNA args).
- [x] `clerkenstein/.github/workflows/alignment.yml` — CI: build + test + gate on push + weekly schedule (honestly-unverified on a real GHA runner).
- [x] **VERIFIED:** gate.sh → 0 (100%/100%); drift-check no-drift → 0; simulated `clerk@2.7.0` bump → 1 (added/changed genes named); gate regression (impossible threshold) → 2. shellcheck clean.
- Review caught + fixed inline: `go run` squashes exit 2→1 (→ build the binary); relative `--new` path broke under the `cd` (→ abspath).

## S2 — Drift runbook documentation — DONE
- [x] `corpus/services/clerkenstein.md` — "Drift detection (M1b)" section: the scripts, the exit-code contract, and the bump → DNA-diff → re-capture goldens → re-score → CI runbook.
- [x] `corpus/architecture/alignment_testing.md` § "How M1 and M1b consume this" — points to the runbook + scripts.
- [x] cross-refs resolve.

## M1b: Hardening

### Pass 1 — 2026-06-03
**Scope manifest:** the milestone's "code" is 2 shell scripts (`scripts/gate.sh`, `scripts/drift-check.sh`) + a CI YAML — all had **no automated test** (exit paths were verified ad-hoc during build). Single stack (shell); scanned in-thread.

**Coverage:** no `%`-coverage tool for bash (project convention is shell+Playwright, not unit-coverage); the relevant "coverage" is the **exit-code matrix**, now fully covered.

**Tests added:** `scripts/drift-test.sh` — **9 assertions** pinning the full contract:
- gate.sh: met → 0; regression → **exactly 2** (regression test for the built-binary fix, vs the `go run` exit-squash).
- drift-check: no-drift → 0; **reformatted-identical DNA → 0** (canonical no-drift, no spurious flag); **relative `--new` path → 0** (regression test for the abspath fix); bumped DNA → 1; missing/not-found/unknown-arg → 3.
- Wired **shellcheck** + **drift-test** as CI steps in `alignment.yml`.

**Bugs fixed inline:** none new (the 2 build-phase bugs were fixed during S1; this pass *pins* them as regressions). 1 shellcheck SC2164 on the new harness fixed (`cd || exit`).

**Flakes stabilized:** none — 3/3 consecutive clean; shellcheck clean; Go suite green.

**Knowledge backfill:** none KB-worthy — the exit-code contract is documented in `clerkenstein.md` § Drift detection + `spec-notes.md`; the `go run`-squash gotcha is captured inline (gate.sh "built binary, not go run"). Question asked, nothing to propagate.

### Stop condition
Stabilized after Pass 1: the exit-path matrix + edge cases (empty/not-found/unknown-arg) + both regression-pins are covered; dimensions 5 (fuzzing — the DNA parser is alignctl's, fuzzed in M0) and 6 (perf — no SLA) are N/A for a 2-script shell milestone; 0 flakes. Nothing further worth adding.

## M1b: Final Review (close-milestone, 2026-06-03)

Section, all sections + harden complete. Phase 1b deferral re-audit GREEN. In-thread review (2 shell scripts + CI YAML + docs). **0 code must-fixes, 0 scope gaps, 0 doc gaps, 0 decisions to blend.**

### Scope — all Fate-1 delivered
- drift-check + gate scripts, CI workflow (S1), drift runbook doc (S2), drift-test harness + CI shellcheck/drift-test steps (harden). overview In-list fully delivered; "Out" items are M2/v1.1 (correctly scoped).

### Adversarial review (Phase 2c — recorded in decisions.md)
- All script failure modes are exercised by `drift-test.sh`: missing/not-found/unknown-arg (→3), reformatted-identical DNA (→0, no spurious drift), the built-binary exit-2 + abspath regressions. 0 latent issues.

### Decision triage
- M1b's drift contract (exit codes) is already in `corpus/services/clerkenstein.md` § Drift detection + spec-notes — nothing new to blend; decisions.md stays archive.
