---
title: "KB Fidelity Audit — M23 Content cutover + referential closure"
date: 2026-06-13
scope: milestone:M23
invoked-by: build-milestone
---

## Verdict
YELLOW

No blind areas; no stale load-bearing claim the implementation would read as truth and act wrongly on. The two
STALE findings are docs lagging *already-landed* (M21/M22) or *about-to-land in this milestone* (§4) work — both
are explicitly retired/updated by M23 §6 (docs) and made true by M23 §4 (code). Tracked as KB-1/KB-2, addressed
in Phase 5.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| DIRECTUS_BASE_ADDR re-point (demo) | safety.md, snapshot-spec.md | stack-injection/gen_injected_override.py | PAIRED |
| DIRECTUS_BASE_ADDR re-point (dev, env-emission) | snapshot-spec.md | stack-core/gen_override.py | PAIRED |
| EnvContract / per-stack BaseAddr | snapshot-spec.md, safety.md | stack-snapshot/directus/provision.go | PAIRED |
| directus_files ref capture | snapshot-spec.md § Media/blobs | stack-snapshot/directus/{media.go,directus.go} | PAIRED (doc overstates) |
| Cross-surface referential closure | snapshot-spec.md § fidelity gate | stack-seeding/dna/snapshot.go | PAIRED (gene is within-surface; cross is new) |
| cms ← Directus dependency truth | corpus/services/cms.md | (platform cms; doc-only) | DOC-ONLY (stale gap claim) |
| jobsimulation/next-web ← Directus path | corpus/services/{jobsimulation,next-web-app}.md | (platform; doc-only) | DOC-ONLY (underspecified) |
| studio-desk ← Directus direct | corpus/services/studio-desk.md | (platform; doc-only) | DOC-ONLY (correct) |

## Fidelity Findings

1. **`directus_files` "always captured + replayed" — STALE/overstated.** snapshot-spec.md ~line 423 claims the
   1,311 file-ref rows are "always captured + replayed (the floor)". CODE: `directus.Surface()` Tables list
   (directus.go) enumerates 9 user tables — **no `directus_files`**; `media.go`'s `FileRefColumns()` /
   `ReferencedFilesFilter()` are defined-but-dead (0 callers). Fix owner: **code wins** — M23 §4 wires it, making
   the doc true. Until then the doc reads ahead of code. Tracked **KB-1**.
2. **cms.md line ~13 — STALE.** Claims the directus replay "skips with stacksnap exit 4 (the M10 collection-schema
   gap)" and cms "reads this public content live from prod" as the *current* state. CODE: M21 auto-provisions the
   structure (provision.go content-schema step "AUTOMATED (M21)"); M22 boots a per-stack Directus. The gap is
   closed / instance booted; only the runtime *cutover re-point* (M23) remains. Fix owner: **doc wins** — M23 §6
   updates cms.md to the M21/M22/M23 truth. Tracked **KB-2**.
3. **Referential-consistency boundary (snapshot-spec ~409-417) — ALIGNED.** Honestly describes the in-progress
   prod-read default + dangling-ref risk; frames M23 as the close (future-tense). No edit needed beyond §6's
   completion update once the cutover lands.
4. **`OpSnapshotReferential` is WITHIN-surface — ALIGNED.** snapshot.go checks every recorded FK's referenced
   table is in the captured set (single-surface closure). The M23 cross-surface gene (content→taxonomy node-id)
   is genuinely NEW code, not an extension of this operator — confirmed, so §5 builds it fresh.
5. **gen_injected (demo) emits `environment:`; gen_override (dev) does not — ALIGNED.** Confirmed asymmetry: the
   demo emitter already builds per-service env blocks (DIRECTUS_TOKEN strip); the dev emitter emits only
   ports/volumes. §1 grows the dev emitter (the one genuinely-new bit of plumbing). No doc claim to fix.
6. **Service-doc dependency chains underspecified.** jobsimulation.md (reads via cms RPC, not direct Directus) and
   next-web-app.md (gateway/cms only, no direct Directus) are accurate but implicit; studio-desk.md (direct
   Directus) is explicit + correct. §6 makes the indirection explicit (env/dependency truth). Incidental.

## Completeness Gaps
- None load-bearing. The `directus_files` wiring (KB-1) is the one real code gap and it is M23 §4's deliverable.

## Applied Fixes
- None inline. Every finding is a deliverable of M23 itself (§4 code → makes KB-1 true; §6 docs → fixes KB-2 +
  finding 6; §5 → the new cross-surface gene). Editing the docs now would pre-empt the milestone's own §6.

## Open Items (require user decision)
- None.

## Gate Result
YELLOW: proceed with tracking. KB-1 (directus_files capture wiring) and KB-2 (cms.md M10-gap staleness) recorded
in decisions.md; both close inside M23 (§4 + §6). No blind area, no implementation-blocking stale claim.
