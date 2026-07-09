---
title: "KB Fidelity Audit — M210 Corpus + skills re-ground"
date: 2026-07-08
scope: milestone:M210
invoked-by: build-milestone
---

## Verdict
YELLOW — proceed. Every stale claim found is precisely this milestone's fix-list (tracked in
`progress.md`); none is read as truth by the implementation. Ground truth = M209's landed rext
code, not the stale docs.

## Topic Inventory

| Topic | Knowledge doc | Code paths (ground truth) | Status |
|---|---|---|---|
| Taxonomy schema (`skiller.*`→`public.*`) | rext-facing tooling docs (snapshot-spec, safety, stories-spec, seeding-spec, recipe-snapshot-world, coverage-protocol) | M209 rext (`.agentspace/rosetta-extensions`) — 89 `public.*` query refs, 1 residual skiller.* | PAIRED (STALE — milestone deliverable) |
| Subgraph count 5→4, no skiller container, `SKILLER_RPC_ADDR=backend:8083` | CLAUDE.md, service_taxonomy, architecture_overview, dependency_map, graphql-wundergraph, skill files (dev-up/stack-snapshot/stack-update/db-query) | M208-verified merged compose (no skiller container, 4 subgraphs) | PAIRED (arch half corrected on colleague branch; skill files STALE — milestone deliverable) |
| Merged member surface / profile-completeness | profile-completeness-spec.md | seeders write `public.*` (already correct in doc) | PAIRED (schema ALIGNED) |
| db-access ↔ tooling agreement | db-access.md / db-query SKILL | both must say `public.*` | PAIRED (colleague re-pointed db-access; tooling flip pending — milestone deliverable) |

## Fidelity Findings
1. **33 stale `skiller.<table>` references** across rext-facing docs (snapshot-spec 19, db-query SKILL 8, db-access 7, stories-spec 3, safety 1, directus-local 1) while M209 rext queries `public.*`. Verdict: STALE. Fix owner: update docs (the milestone's core work). NOT a blocker — the milestone overwrites these; it does not read them as truth.
2. **Subgraph/container/RPC drift in skill files** (dev-up/reference, stack-snapshot/SKILL, stack-update/reference, db-query/SKILL) — still reference the pre-merge compose. STALE. Fix owner: docs (Section 5).
3. **profile-completeness-spec.md "43/44" count** (orchestrator END-STATE #2): NO literal `43/44` exists anywhere in the corpus (verified: only `340/341` avatar-coverage narrative at line 72; never in file history via `git log -S`). Schema refs already `public.*`. One prose skiller mention: line 105 "the closure gene governs skiller node-ids" (a permitted historical/conceptual reference, not a tooling-queries-skiller.<table> claim). Verdict: UNVERIFIABLE as literally stated → resolve during Section 2 with evidence-based judgment; document the discrepancy in decisions.md.

## Completeness Gaps
None. Every milestone topic has a doc anchor; no blind areas.

## Applied Fixes
None inline — all findings are the milestone's own section deliverables (fixing them here would pre-empt the section work + its commit discipline).

## Open Items (require user decision)
None blocking. The "43/44" literal (finding 3) is resolved by evidence in Section 2, not a user decision.

## Gate Result
YELLOW: proceed to Phase 1. Findings 1–2 are the milestone's tracked deliverables (progress.md sections 3/5); finding 3 resolves in Section 2. Re-audit at close should be GREEN once the flips land.
