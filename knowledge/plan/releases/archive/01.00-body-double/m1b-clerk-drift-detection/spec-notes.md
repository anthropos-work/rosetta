# M1b — spec notes

## Pre-flight audits — S1 (drift-check + CI gate)
**Phase 0b — KB-fidelity (M1b): GREEN.** M1b reuses M0 wholesale; the contract docs all exist —
`corpus/architecture/alignment_testing.md` (the framework + `alignctl dna diff`/`run`),
`corpus/services/clerkenstein.md` (the mirror M1b gates), `corpus/services/clerk-integration.md`. No
blind area; M1b *adds* a drift-runbook section to clerkenstein.md (the Delivers).

## The drift contract (exit codes — the CI-gate interface)
`drift-check.sh [--old PINNED_DNA] --new DNA` →
- **0** — no DNA drift AND the alignment gate is met (mirror still faithful).
- **1** — the DNA moved (`alignctl dna diff` non-empty: added/removed/changed genes) — the source
  surface changed; the mirror + DNA need a human pass (re-author + re-capture goldens).
- **2** — the alignment gate regressed (`alignctl run --gate` failed) — the mirror diverged from the
  (unchanged or updated) source; some genes broke.

`gate.sh` is the inner gate (exit 0 = gate met, 2 = regressed); `drift-check.sh` wraps it with the
DNA-diff step.

## Layout (clerkenstein repo)
```
scripts/gate.sh           build runner + alignctl run --gate
scripts/drift-check.sh    dna diff + gate.sh, combined report + exit
.github/workflows/alignment.yml   CI: push + weekly schedule (honestly-unverified — no GHA runner here)
```
`ALIGN_DIR` env (default `../../test/alignment`) locates rosetta's `alignctl`. The scripts are
runnable locally now; the GHA YAML is the CI mechanization for when clerkenstein is a pushed repo.

## Demonstration (no live Clerk bump available)
S1 proves the mechanism with a **simulated** `clerk@2.7.0` DNA delta (e.g. an added capability + a
changed operator), the way M0's own `dna diff` smoke test did — `drift-check.sh` must exit non-zero
and name the moved genes.
