---
title: "KB Fidelity Audit — M2b Clerkenstein repo consolidation"
date: 2026-06-03
scope: milestone:M2b
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Clerkenstein repo (the mirror being reorganized) | `corpus/services/clerkenstein.md` | `anthropos-demo/clerkenstein/{authn,bapi,orgclient,fapi,webhook,cmd,dna,golden,golden-js,scripts}` | PAIRED |
| Alignment framework (how fidelity is measured) | `corpus/architecture/alignment_testing.md` | `test/alignment/` + clerkenstein DNAs/runners | PAIRED |
| Repo's own `knowledge/` base (M2b net-new deliverable) | — (to be authored in S2) | `anthropos-demo/clerkenstein/knowledge/` (not yet exists) | DOC-ONLY-PLANNED |

The third row is a planned **deliverable** of this milestone (overview "Delivers → knowledge: net-new"),
not a blind area — M2b authors it. Not a finding.

## Fidelity Findings

1. **`corpus/services/clerkenstein.md` — flat-structure code map.** Source: clerkenstein.md §"Architecture
   & code map". Expected: dirs `authn/ orgclient/ fapi/ bapi/ webhook/ cmd/{clerkrun,jsfapirun} dna/
   golden/ golden-js/ scripts/`. Actual: confirmed exactly present at HEAD `26b2490` (verified via `ls`).
   Verdict: **ALIGNED** (current truth). S5 will intentionally restructure-and-slim this doc — that is the
   milestone's job, not staleness.

2. **`corpus/services/clerkenstein.md` — script defaults (M1b drift).** Source: §"Drift detection (M1b)".
   Expected: `ALIGN_DIR` default `../../test/alignment`; `gate.sh` defaults `RUNNER_PKG=./cmd/clerkrun`,
   `GOLDEN_DIR=golden`, `DNA=dna/clerk-2.6.0.json`. Actual: confirmed in `scripts/gate.sh` lines 18/21/24.
   Verdict: **ALIGNED**. S1 repoints these (scripts move one level deeper → `../../../test/alignment`).

3. **`corpus/services/clerkenstein.md` — test counts + coverage.** Expected (headline): 7 packages, 112
   test/fuzz funcs (107 tests + 5 fuzz). Actual: `grep` counts 112 (107 + 5) across 7 packages. Verdict:
   **ALIGNED**.

4. **`corpus/services/clerkenstein.md` — green gates.** Expected: Go 22/22 100%/100%, JS 9/9 100%/100%,
   drift ALL PASS. Actual (baseline run at HEAD `26b2490`): Go gate exit 0 `100.0%/100.0% (22/22)`; JS gate
   exit 0 `100.0%/100.0% (9/9)`; drift-test `ALL PASS`; `-race` clean; gofmt/vet/shellcheck clean. Verdict:
   **ALIGNED** — the green-gate safety net is confirmed BEFORE any moves.

5. **`corpus/architecture/alignment_testing.md` — exists + is the measuring contract.** Source: state.md
   references it as the framework doc. Actual: present (13.6KB). Verdict: **ALIGNED** (not modified by M2b;
   the repo's `knowledge/alignment.md` will link to it).

## Completeness Gaps

None blocking. The repo's own `knowledge/` base does not yet exist — but it is the explicit **net-new
deliverable** of S2 (overview §"Delivers → knowledge"), so its absence is planned, not a blind area. The
milestone authors it from the contract docs above + the live repo state.

## Applied Fixes

None needed — all PAIRED topics ALIGNED. No stale claims, no broken cross-refs found in the two
load-bearing corpus docs.

## Open Items (require user decision)

None.

## Gate Result

**GREEN — proceed to Phase 1.** Both load-bearing corpus docs exist and accurately describe the current
flat repo at HEAD `26b2490`; the baseline green-gate (Go 22/22, JS 9/9, drift 9/9, lints clean) is
confirmed, so any post-move gate break is provably a path/import error introduced by the reorg, not a
pre-existing condition. The repo's `knowledge/` base is a planned deliverable, not a blind area.
