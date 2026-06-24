---
title: "KB Fidelity Audit — M40 Directus serve-grant"
date: 2026-06-24
scope: milestone:M40
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Per-stack Directus serve rows (collections + public-read permissions) | `corpus/ops/snapshot-spec.md` (§ structure extension, l.380-430) | `rext stack-snapshot/directus/structure.go` (serve-row capture), `firewall/firewall.go` (AssertStructuralMetadata) | PAIRED |
| Anonymous (token-less) public read on the per-stack Directus | `corpus/ops/safety.md` §2.9 + l.146/193/220 | cms Directus client (Bearer only when token!=""), `DIRECTUS_TOKEN` blank | PAIRED |
| The up→snapshot→seed→use→down flow (fresh /demo-up acceptance) | `corpus/ops/demo/README.md` | `/demo-up`, `stacksnap replay --surface directus` | PAIRED |
| Directus relational metadata (directus_relations / directus_fields) for nested O2M/M2M reads | — | `rext stack-snapshot/directus/structure.go` (NOT YET captured) | **CODE-ONLY (the M40 deliverable)** |

## Fidelity Findings

1. **safety.md §2.9 / l.193 — anonymous read mechanism — ALIGNED.** Doc: "cms omits the `Authorization` header when the token is empty; prod Directus serves … anonymously." Verified live: `docker exec demo-3-cms-1` → `DIRECTUS_TOKEN` is EMPTY; the public policy is the operative permission set. Doc is the correct contract. (Refutes the build-agent's initial service-token hypothesis — recorded M40-D1.)
2. **snapshot-spec.md l.380-382 — serve-row capture — ALIGNED.** Doc: "captures the structure (DDL + PKs + sequences + serve rows)". Code `CaptureServeRows` does exactly this for `directus_collections` + `directus_permissions`. The M40 extension (relations/fields) is promised by the milestone's `Delivers →` line, so it is a planned doc-production deliverable, not a blind area.
3. **firewall TenantScopeColumns — ALIGNED + extensible.** `directus_relations` and `directus_fields` columns (verified via sanctioned prod structural read) carry NONE of `TenantScopeColumns` → admissible under the existing `AssertStructuralMetadata` carve-out. The M40 work extends the carve-out's table set without loosening the predicate.

## Completeness Gaps

1. **Relational metadata not yet documented as a serve-row class.** snapshot-spec.md documents collections + permissions serve rows but not `directus_relations`/`directus_fields`. This is the M40 `Delivers →` target — authored in Phase 5, not a pre-existing gap.

## Applied Fixes
None required pre-build. The triples are recorded in spec-notes (below) for fast re-audit.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. No blind areas, no stale load-bearing claims. The relations/fields synthesis is the milestone's own deliverable (CODE-ONLY → becomes PAIRED when Phase 5 authors the snapshot-spec extension).
