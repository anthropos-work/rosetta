---
title: "KB Fidelity Audit — M209 rext tooling re-ground"
date: 2026-07-08
scope: milestone:M209
invoked-by: build-milestone
---

## Verdict
YELLOW

Pre-flip doc staleness is EXPECTED and chartered to M210 (the lockstep corpus body-flip). No blind areas;
no stale claim that M209's implementation reads as truth (M209 re-grounds against M208's empirically-verified
merged-schema facts + the design workflow's file:line enumeration, not against these doc bodies).

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Taxonomy snapshot capture+replay | corpus/ops/snapshot-spec.md | stack-snapshot/ (taxonomy, capture, store, firewall) | PAIRED (doc STALE → M210) |
| Taxonomy-ref / named-skill seeding | corpus/ops/seeding-spec.md | stack-seeding/seeders/, dna/ | PAIRED (doc prose STALE → M210) |
| Tenant-data firewall / public predicate | corpus/ops/safety.md | stack-snapshot/firewall/firewall.go | PAIRED (1 evidence row STALE → M210) |
| Merged-schema ground truth (public.*) | M208 spec-notes + merge fact-sheet | (contract M209 implements against) | PAIRED — VERIFIED |

## Fidelity Findings

1. **snapshot-spec.md describes the taxonomy surface as `skiller.*` (19 mentions).** STALE vs the post-merge
   `public.*` reality. Verdict: STALE. Fix owner: update doc — but chartered to **M210** (`m210-corpus-reground`
   overview: "snapshot-spec.md — 26 mentions — the taxonomy surface enumeration"). NOT load-bearing for M209's
   code re-ground. → KB-1.
2. **safety.md firewall evidence row `skiller.skills 42,763 public`.** STALE on both the schema (`skiller`→`public`)
   and the count (now ~42,790). Verdict: STALE. Fix owner: update doc → **M210** (overview names it explicitly).
   The firewall CODE (`organization_id IS NULL`) is schema-agnostic and unchanged (design-confirmed). → KB-3.
3. **seeding-spec.md — 0 `skiller.<table>` query refs currently**, but prose ("the public skiller catalog /
   taxonomy") flips in **M210**. Verdict: STALE-prose. → KB-2.

## Completeness Gaps
None new. The merged-schema contract M209 needs is fully documented in M208's spec-notes + merge fact-sheet
(empirically verified: `public.skills/job_roles/...` exist with `organization_id`; vector cols
`extensions.vector(1536)`; trigram `extensions.gin_trgm_ops`; public count 42,790).

## Applied Fixes
None. All findings are chartered to M210's body-flip (do NOT flip corpus doc bodies in M209 — that is M210's
lockstep deliverable per the release plan). Recording KB-1/2/3 as Fate-2 (covered by M210) in decisions.md.

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed with tracking. KB-1/2/3 recorded in decisions.md as Fate-2 (already owned by M210). No blocker.
