---
title: "KB Fidelity Audit — M9b (Taxonomy snapshot)"
date: 2026-06-06
scope: milestone:M9b
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Snapshot framework (capture/replay/firewall/manifest/store) | `corpus/ops/snapshot-spec.md` | `.agentspace/rosetta-extensions/stack-snapshot/**` | PAIRED |
| Taxonomy schema (skiller tables, embeddings, translations) | `corpus/services/skiller.md` | live `skiller` schema (prod, catalog-verified) | PAIRED |
| Seeding DAG node + waived taxonomy surface | `corpus/ops/seeding-spec.md` | `stack-seeding/dna/data-dna.json` | PAIRED |
| Snapshot-fidelity + public-only genes | `corpus/architecture/alignment_testing.md` | `stack-seeding/dna/snapshot.go` | PAIRED |

## Fidelity Findings

1. **snapshot-spec.md `PublicVia` ↔ capture filter.** The doc describes column-less tables as "scoped via a
   public parent". The M9a `adapters.go` records `PublicVia` in the manifest but applies an EMPTY capture filter
   to column-less tables (it would capture the whole table). Verdict: **not stale** — the doc states the *intent*;
   the parent-scoped capture filter is exactly M9b's load-bearing deliverable (recorded as M9b-D2). The framework
   was correct for M9a's toy surface (the whole `item_vectors` table is public); the real taxonomy needs the
   filter, which M9b adds. No doc fix needed; the gap is in-scope code work.
2. **skiller.md vector column + dimension.** Doc says `small_embedding3 extensions.vector(1536)` for both embedding
   tables. ALIGNED — catalog-verified (`format_type` → `extensions.vector(1536)`).
3. **skiller.md table set.** Doc names skills/job_roles/specializations/categories + the dedicated embedding +
   translation tables + `jobroleskill`. ALIGNED with the live schema.
4. **alignment_testing.md five fidelity operators.** Doc names snapshot-row-count / -structural / -referential /
   -embedding-dim / -public-only. ALIGNED with `snapshot.go` (`OpSnapshot*`).
5. **seeding-spec.md DAG node + waived surface.** Doc shows `… → taxonomy/content (snapshot) → activity` and the
   two `waived` surfaces (taxonomy + content). ALIGNED with `data-dna.json` (`waived-m7c`).

## Completeness Gaps

1. **skiller.md does not mention `job_role_categories`** (the table `job_roles.category_id` FK-references; NOT
   `skiller.categories`). Classification: **incidental** for skiller.md's purpose (a service doc, not a snapshot-
   surface map). The snapshot-surface FK graph is captured in M9b's spec-notes.md (the authoritative place for the
   capture plan). No skiller.md edit required for M9b's contract; noted for the close pass if the surface map ever
   moves into skiller.md.

## Applied Fixes
- None to knowledge docs (all PAIRED claims ALIGNED). The prod FK/referential findings (10-table surface,
  job_role_categories, both-endpoints job_role_skills, dim 1536) were recorded in M9b's `decisions.md` (M9b-D1…D5)
  and `spec-notes.md` (the 10-table capture plan) — the topic→doc→code triples future audits start from.

## Open Items (require user decision)
- None.

## Gate Result
GREEN: proceed to build-milestone Phase 1. Knowledge docs are the faithful contract; the parent-scoped-filter gap
is in-scope code work (M9b-D2), not a stale doc claim.
