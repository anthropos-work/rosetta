---
title: "KB Fidelity Audit — M9a (snapshot framework)"
date: 2026-06-06
scope: milestone:M9a
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Seeding framework + DAG node + isolation guard (the read-side analog) | `corpus/ops/seeding-spec.md` | `rosetta-extensions/stack-seeding/{isolation,seeder,pg}/` | PAIRED |
| Data-DNA harness (the dimension M9a extends) | `corpus/architecture/alignment_testing.md` §"data dimension (M7b)" | `rosetta-extensions/stack-seeding/dna/` + `cmd/datadna` | PAIRED |
| Full-clone-with-customer-data precedent (the contrast) | `corpus/ops/staging_from_dump.md` | — (operational guide) | DOC-ONLY (by design) |
| Prod DB read access + public/customer boundary (read foundation) | `corpus/ops/db-access.md` + `/db-query` skill | live prod (MCP `postgres` tool) | PAIRED (verified live) |
| The snapshot extension + capture/replay contract + store | `corpus/ops/snapshot-spec.md` (M9a *delivers*) | `rosetta-extensions/stack-snapshot/` (M9a *creates*) | DOC-ONLY forward-ref (the milestone's own deliverable) |

No BLIND-AREA: every topic the milestone reads as a contract has a doc anchor. The two DOC-ONLY rows are
intentional — `staging_from_dump.md` is an operational guide (no code), and `snapshot-spec.md` is what M9a authors.

## Fidelity Findings

1. **db-access.md §"two hard rules" — capture-source preference** — Source: `corpus/ops/db-access.md:35`.
   Expected (doc, as written 11:02): "never read the hot primary under load — read a non-primary copy (read replica
   → restore-from-backup)". Actual contract (M9a-D3, finalized in spec-notes after the doc was committed): the
   precedence is **dump-ingest [default] → safe throttled primary read [fallback, MVCC = no write blocking] →
   restore-from-snapshot / replica [upgrades]** — a safe primary read is now a *sanctioned* fallback, not forbidden.
   Verdict: **STALE** (load-bearing — the capture policy reads this as truth). Fix owner: update doc.
   **APPLIED** inline (see below).

2. **db-access.md public/customer split numbers** — Source: `corpus/ops/db-access.md:43-56`. Cross-checked live
   prod 2026-06-06: `skiller.skills` 42,763 public / 794 customer · `categories` 22 public · `cms.studio_documents`
   0 public / 3,060 customer (all-customer → excluded). Verdict: **ALIGNED** (exact match).

3. **db-access.md catalog-only sizing + pgvector index claim** — Source: `corpus/ops/db-access.md:58-71`. Live: the
   catalog-only `pg_class`/`pg_total_relation_size` pattern returns instantly; `skill_embeddings` 692 MB total but
   heap only 3264 kB → ~689 MB is the pgvector index (rebuild-on-replay, don't transport). Verdict: **ALIGNED**.

4. **db-query skill connection model** — Source: `.claude/skills/db-query/SKILL.md`. Live verified: MCP `postgres`
   tool returns `postgres` / `marco_read` / RDS private IP (10.2.22.13), exactly as documented. Verdict: **ALIGNED**.

5. **seeding-spec.md DAG node + waiver** — Source: `corpus/ops/seeding-spec.md:72,150`. The DAG already names the
   `taxonomy/content (snapshot)` node; taxonomy+content are documented `waived` ("the snapshot/shared-store hard
   line, ~v1.2"). Matches the `StatusWaivedPrefix` in `dna/dna.go:73`. Verdict: **ALIGNED** — the milestone's
   surface-status extension (`snapshot-seeded`) slots cleanly alongside `seeded-`/`planned-`/`waived-`.

6. **alignment_testing.md data dimension** — Source: `corpus/architecture/alignment_testing.md:272-302`. The doc
   describes the M7b data-DNA (catalog + measure + diff, 4 structural operators, criticality weights) and matches
   `dna/dna.go` (the same operator set + `Criticality` 3/2/1). The M9a fidelity-gene class extends this cleanly.
   Verdict: **ALIGNED**.

## Completeness Gaps
None critical. The `snapshot-spec.md` doc is a DOC-ONLY forward-ref (4 inbound links from db-access.md / db-query);
those resolve when M9a Phase 5 authors the doc — expected for a milestone whose deliverable is that doc.

## Applied Fixes
- `corpus/ops/db-access.md:32-39` — rewrote the capture-source preference sentence to the canonical M9a-D3
  precedence (dump-ingest default → safe primary-read fallback → restore/replica upgrades), correcting the
  pre-D3 "read replica → restore-from-backup" framing. One-line-class fix; the canonical full policy lands in
  `snapshot-spec.md` (§capture-source) during M9a Phase 5.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. The one stale load-bearing claim (finding 1) was fixed inline; all other doc claims are
ALIGNED and several were verified against live prod. The snapshot-spec.md forward-refs are the milestone's own
deliverable, not blind areas.
</content>
</invoke>
