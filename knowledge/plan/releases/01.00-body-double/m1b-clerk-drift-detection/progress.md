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
