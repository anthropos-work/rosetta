---
title: "KB Fidelity Audit — M24 Docs convergence + hygiene strand"
date: 2026-06-13
scope: milestone:M24
invoked-by: build-milestone
---

## Verdict
GREEN

The stale doc claims this audit confirmed (`external_services.md`, `service_taxonomy.md`, `quick_ops.md`) are
**the milestone's own deliverables** — M24 exists to correct them. Every topic the milestone touches is PAIRED
(doc anchor + code both exist). No blind areas. Nothing the milestone *reads as truth* is stale; the stale claims
are the work, and they are all anchored and verified-against-code.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Local Directus / compose reality | `corpus/architecture/external_services.md`, `service_taxonomy.md`, `corpus/ops/quick_ops.md` | `stack-dev/platform/docker-compose.yml` | PAIRED (docs stale — M24 deliverable) |
| Per-stack Directus end-state | `corpus/ops/directus-local.md` (261 lines, exists), `snapshot-spec.md`, `safety.md` | `rosetta-extensions/stack-snapshot`, `dev-setdress.sh` | PAIRED |
| Snapshot known-state / exit-4 semantics | `corpus/ops/snapshot-spec.md` § known-state (389-413) | `stacksnap` CLI | PAIRED (partly M22-current; M24 finishes the cutover narrative) |
| Safety §2 write-side deltas | `corpus/ops/safety.md` §2 | `stack-core` isolation guard | PAIRED |
| Alignment zero-critical-genes guard | `corpus/architecture/alignment_testing.md` | `alignment/internal/compare/compare.go:247`, `alignment/internal/dna/dna.go:169` | PAIRED (guard absent — hygiene fix) |
| README index-row guard | (new lint — corpus directory READMEs) | `rosetta-extensions` (new) | CODE-ONLY (net-new tool) |
| `/project-stats` scope | `developer-kit/skills/project-stats/stats.sh` | scans `stack-*/` (bug) | CODE-ONLY (hygiene fix) |
| Go toolchain pin | — | `rosetta-extensions` go.mod / toolchain pins | CODE-ONLY (lazy bump) |

## Fidelity Findings

1. **`service_taxonomy.md:242-260` — STALE (load-bearing for the correction, not for impl).** Presents Directus as a
   "Dockerized" local service with `image: directus/directus:10.10.1`, port 8055, dedicated `directus` schema.
   **Actual:** `stack-dev/platform/docker-compose.yml` has **no directus service** — only
   `DIRECTUS_BASE_ADDR=https://content.anthropos.work` + `DIRECTUS_PUBLIC_BASE_ADDR=...` (prod pointers, lines
   237-238). Fix owner: update doc. **This is M24 §1 work.**
2. **`external_services.md` — STALE.** Lines 122 + 181 (`directus/directus:10.10.1`), 270-271 (`admin@example.com` /
   `password`), the `docker-compose.yml directus` snippet (~177-181), and the `docker compose up -d directus` / `ps
   directus` examples (552, 617) — all describe a local Directus service that does not exist in the platform compose.
   Fix owner: update doc. **M24 §1 work.**
3. **`quick_ops.md:162` — STALE/misleading.** Lists `Directus | 8055` among local stack ports; no local Directus
   listens there in the prod-read posture. Fix owner: update doc. **M24 §1 work.**
4. **`snapshot-spec.md` known-state (389-413) — PARTIALLY CURRENT.** Already reflects M22 ("EXECUTED on a
   `--local-content` stack, print-only otherwise"); M24 §2 finishes the narrative to reflect M23's data-plane cutover
   (per-stack Directus now serves real content). Not stale-as-truth; an in-flight cutover narrative to complete.
5. **Alignment zero-critical-genes — CONFIRMED ABSENT (the §6 hygiene fix).** `compare.go:247` `pct(n,d)` returns
   **`100.0` when `d==0`**; `compare.go:92` computes `rep.Critical = pct(critAligned, critTotal)` where `critTotal==0`
   for a DNA with zero critical-class genes. Result: a mirror with **no critical genes scores a perfect 100% critical**
   and passes any critical gate vacuously. `dna.go:169` `Validate()` rejects per-capability invalid criticality
   (`Weight()==0`) but **never** the aggregate "no critical genes at all" condition. Verified at the exact lines the
   overview cites. Fix owner: code (M24 §6).

## Completeness Gaps

1. `corpus/ops/directus-local.md` exists (261 lines, last edited 2026-06-13) — so M24 §2's "finish" is a refinement of
   an existing doc, **not** a net-new authoring task. No blind area.
2. No undocumented load-bearing behavior found in the code paths M24 will touch beyond the enumerated stale claims.

## Applied Fixes
None applied inline — every stale claim is an explicit M24 deliverable (§1–§3) and will be corrected in the build
phase with full context, not piecemeal here. Applying them now would pre-empt the milestone's own sections.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. The stale claims are anchored and are the milestone's chartered work; no blind areas; no
claim the milestone depends on as truth is stale.
