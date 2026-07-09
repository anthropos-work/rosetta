---
title: "KB Fidelity Audit — milestone:M208"
date: 2026-07-08
scope: milestone:M208
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Merged skills-taxonomy shape (skiller → app/`public`) | `corpus/services/backend.md`, `corpus/services/skiller.md` | `stack-dev/app/internal/data/ent/schema/{skill,jobrole,category,specialization,*_embeddings,skiller_mixins}.go`; `stack-dev/app/internal/rpc/skillerrpc/` | PAIRED (docs describe PRE-merge world — being corrected this release) |
| Skiller RPC re-point (`SKILLER_RPC_ADDR`) | `corpus/services/backend.md`, `corpus/architecture/dependency_map.md` | `platform@origin/main:docker-compose.yml` (`SKILLER_RPC_ADDR=http://backend:8083`) | PAIRED (stale in current clone / corpus; correct at origin/main) |
| 4-subgraph federation (skiller subgraph removed) | `corpus/services/graphql-wundergraph.md`, `CLAUDE.md` | `platform@origin/main` compose/router config | PAIRED (docs say 5 subgraphs — M210 flips) |
| Stack re-sync ops (`make pull`/`up`/`migrate`) | `corpus/ops/update_guide.md`, `corpus/ops/setup_guide.md`, `corpus/ops/run_guide.md` | `stack-dev/platform/Makefile` | PAIRED (stable; not merge-affected) |
| Merge fact-sheet (the M208 deliverable anchor) | `corpus/services/backend.md` (net-new section) + `corpus/services/skiller.md` (stub) | grounded in the verified app clone + `platform@origin/main` | DOC-ONLY → delivered by this milestone |

## Fidelity Findings

1. **`corpus/services/backend.md` — pre-merge consumer/dependency claims.**
   - Source: `backend.md` §Role & Responsibility ("consumed by skiller, jobsimulation, skillpath, cms"), §Dependencies ("Skiller — taxonomy and matching RPC"), §Redis Streams ("Consumer: cms, skiller events").
   - Expected (doc): skiller is a separate downstream service app calls.
   - Actual (code): the skiller domain (taxonomy, embeddings, AI matching) is merged INTO app (`internal/data/ent/schema/skill.go` et al.; merge commit `1fc00c78`); `SkillerService` RPC served by app; `SKILLER_RPC_ADDR=http://backend:8083`.
   - Verdict: STALE — but **not load-bearing for M208's implementation** (M208 authors the correction; it does not build against these claims). The colleague's `origin/docs/skiller-in-app-merge` already drafts the full body-flip.
   - Fix owner: update doc. **Routed to M210 (full body-flip) — Fate 2.** M208 pins the authoritative fact-sheet anchor that M210/M209/M211 grade against.

2. **`corpus/services/skiller.md` — describes a live standalone service.**
   - Source: whole doc.
   - Expected (doc): skiller is a running Go microservice with its own schema/subgraph/container.
   - Actual (code): decommissioned; merged into app; not in `repos.yml`/compose at `origin/main`.
   - Verdict: STALE — tracked. **Routed to M210 — Fate 2** (colleague branch already drafts the "merged into app" stub). Not read as truth by M208.

3. **`CLAUDE.md` / `graphql-wundergraph.md` — "5 subgraphs".**
   - Expected: 5 subgraphs (app, skiller, jobsimulation, cms, skillpath).
   - Actual: 4 (backend, jobsimulation, cms, skillpath) at `origin/main`.
   - Verdict: STALE — tracked. **Routed to M210 — Fate 2.** M208 pins "4 subgraphs" in the fact-sheet.

## Completeness Gaps

None load-bearing for M208. The merged domain's deep documentation (per-table, per-RPC-method) is M210's body-flip scope; M208 delivers the concise authoritative fact-sheet only.

## Applied Fixes

None applied inline here — the corrective content IS the milestone's Phase-5 deliverable (the merge fact-sheet), not an incidental doc bump. Recording the stale-but-tracked claims as KB items in `decisions.md` (KB-1, KB-2, KB-3) all routed to M210 (Fate 2).

## Open Items (require user decision)

None. The pre-merge corpus staleness is the premise of the release, fully planned: M208 pins the fact-sheet anchor; M210 flips the bodies; M209 re-grounds rext; M211 does bring-up acceptance.

## Gate Result

**YELLOW: proceed with tracking.** No blind area (backend.md/skiller.md exist and M208's `delivers` line makes the fact-sheet a milestone deliverable). No stale claim that M208's own implementation reads as truth — M208 authors the correction, grounded in the verified app clone + `platform@origin/main` + the colleague's docs branch. Stale pre-merge corpus bodies are tracked as KB-1/2/3, all Fate-2 (owned by M210). build-milestone may enter Phase 1.
