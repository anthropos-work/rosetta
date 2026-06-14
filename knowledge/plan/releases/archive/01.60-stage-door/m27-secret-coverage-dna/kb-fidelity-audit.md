---
title: "KB Fidelity Audit — M27 Secret-coverage DNA + source ingestion"
date: 2026-06-14
scope: milestone:M27
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths (precedent the milestone mirrors) | Status |
|---|---|---|---|
| The DNA framework (gene/score/diff, Criticality 3/2/1, Overall+Critical, vacuous-100 guard) | `corpus/architecture/alignment_testing.md` | `.agentspace/rosetta-extensions/alignment/` + `stack-seeding/dna/{dna,measure,operators}.go` | PAIRED |
| The data-DNA precedent (`datadna` — the one-sided harness to mirror) | `corpus/ops/seeding-spec.md` § "Verifying a seed — datadna" | `.agentspace/rosetta-extensions/stack-seeding/dna/` + `cmd/datadna/main.go` | PAIRED |
| Values-blind / `PreflightEnv` safety discipline | `corpus/ops/safety.md` (lines 156-205) | `.agentspace/rosetta-extensions/stack-seeding/isolation/` (`Guard.PreflightEnv`) | PAIRED |
| The `stack-secrets` section itself (the M27 deliverable) | — (M29 authors `corpus/ops/secrets-spec.md`; `Delivers →` line on M29) | net-new (`stack-secrets/` does not exist yet) | DOC-ONLY (planned) |

## Fidelity Findings

1. **DNA framework — Criticality 3/2/1 → weight.**
   - Source: `corpus/architecture/alignment_testing.md:97`. Expected: critical/standard/optional → weight 3/2/1.
   - Actual: `stack-seeding/dna/dna.go` `Criticality.Weight()` returns 3/2/1. Verdict: ALIGNED.

2. **DNA framework — two-metric score (Overall weighted, Critical unweighted).**
   - Source: `alignment_testing.md:154-156`. Expected: Overall = Σ(weight·aligned)/Σ(weight); Critical = unweighted count ratio.
   - Actual: `stack-seeding/dna/measure.go` `Score{Overall, Critical}`, Overall weighted, Critical unweighted. Verdict: ALIGNED.

3. **DNA framework — vacuous-100% / empty-denominator guard.**
   - Source: `alignment_testing.md:160-166` (zero-critical-genes guard). Expected: `ratio()` returns 1.0 on den==0, but a DNA with no critical gene is treated as vacuous.
   - Actual: `measure.go` `ratio()` returns 1.0 when den==0; this is the anti-vacuous guard the milestone reuses. Verdict: ALIGNED. (M27 must encode the same "at least one critical gene" load-time check — noted in spec-notes.)

4. **data-DNA precedent — `datadna` CLI verbs + 0/1/3 exit codes.**
   - Source: `corpus/ops/seeding-spec.md:189-201`. Expected: catalog/introspect/measure/diff; measure exit 1 if critical<100%; diff exit 1 on drift; usage error.
   - Actual: `stack-seeding/cmd/datadna/main.go` `exitOK=0/exitMoved=1/exitUsage=3`, all four verbs present. Verdict: ALIGNED.

5. **Safety — `PreflightEnv` anchor + values-blind.**
   - Source: `corpus/ops/safety.md:156-205`. Expected: `Guard.PreflightEnv(env, target)` asserts+repairs env before a tool runs; strips Directus write token on non-prod.
   - Actual: anchor present at the cited line range; clause accurate. Verdict: ALIGNED. (M27 verbs are read-only/values-blind — no env write — so this is the discipline M28's `provision` will emit against; M27 honors "no value ever read/echoed".)

## Completeness Gaps
None blocking. The secret-DNA's per-gene fields (`scope`, `source_hint`, alias families, `waived-<reason>`) are net-new vocabulary with no existing doc — but that vocabulary IS the M27 deliverable, and its corpus home (`secrets-spec.md`) is explicitly M29 scope per the roadmap `Delivers →` line. Not a blind area: the *patterns* (gene/score/diff/waived/criticality) are fully documented and the precedent code exists to mirror.

## Applied Fixes
None needed — all PAIRED claims ALIGNED; no stale anchors; no broken cross-references.

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed to build-milestone Phase 1. Every load-bearing claim the milestone reads as truth (Criticality weights, the two-metric score, the empty-denominator/vacuous-100 guard, the 0/1/3 exit contract, the values-blind/PreflightEnv discipline) is verified aligned against the precedent code. The `stack-secrets` section is net-new by design; its corpus doc is M29, not a M27 blind area.
