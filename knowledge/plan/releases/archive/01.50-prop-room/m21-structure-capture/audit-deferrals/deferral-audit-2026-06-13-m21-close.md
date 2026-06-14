---
title: "Deferral Audit — milestone M21"
date: 2026-06-13
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (M21 is the first milestone of v1.5 — nothing inherited from a prior milestone in this release).
- No chronic / aged-out patterns (every item was first deferred 2026-06-11..13 — days old, not months; none crossed a closed destination milestone).
- Every item has a clear single-pass fate decision; the two route-forwards the orchestrator flagged are confirmed legitimately-future and routed (not blocked-on).

## Summary
- Total deferrals in scope: 5 (4 M21-originated + 1 cross-release backlog item carried into v1.5 by design)
- Single deferrals: 5
- Repeat deferrals: 0
- Chronic / aged-out patterns flagged: 0

## Deferral Inventory

```yaml
- id: DEF-M21-01
  item: "HARDEN-M21-AP1-replaycmd-conn-seam — hermetic replayCmd wiring test needs replayCmd refactored to accept an injectable connector"
  origin_milestone: M21
  first_deferred_on: 2026-06-13   # final harden pass
  last_seen_in: hardening-ledger.md:107-111
  destination: "a future replayCmd-seam build iter"
  reason_recorded: "architectural, >50 lines; changes replayCmd signature + main dispatch + touches the load-bearing replay path"
  partial_attempted: no

- id: DEF-M21-02
  item: "HARDEN-M21-serve-live-integration — automated integration harness for the serve-row render SQL needs a live directus Postgres"
  origin_milestone: M21
  first_deferred_on: 2026-06-13   # final harden pass
  last_seen_in: hardening-ledger.md:112-115
  destination: "live-integration backlog"
  reason_recorded: "needs a live directus Postgres (stand the container up); exceeds harden-pass scope. The render is hermetically unit-tested + hand-validated live per iter-08"
  partial_attempted: no   # the render IS hermetically unit-tested; only the LIVE integration harness is routed

- id: DEF-M21-03
  item: "directus_files ref capture — wire the dead media.go (filter/columns) into directus.Surface() for the real-image asset refs"
  origin_milestone: M21
  first_deferred_on: 2026-06-11   # carried iter-02..08
  last_seen_in: progress.md:73, iter-08/progress.md:76
  destination: "M23 (asset plane / content cutover)"
  reason_recorded: "carried each iter as a stage-3/4 sub-task; never blocked the gate (anonymous serve of simulations works without it); it is asset-ref plumbing that belongs with M23's asset-plane work"
  partial_attempted: no

- id: DEF-M21-04
  item: "M23 referential closure of the 20 dangling relations (relations pointing to collections outside the captured subset)"
  origin_milestone: M21
  first_deferred_on: 2026-06-11   # iter-03 finding
  last_seen_in: progress.md:74, decisions.md (iter-03 registry inventory)
  destination: "M23 (referential closure — explicitly owned)"
  reason_recorded: "a booted Directus has dangling relation defs until M23 closes them or they're pruned; not needed for the M21 anonymous-serve gate"
  partial_attempted: no

- id: DEF-M10-01
  item: "S3 media blob BYTES + cloud SnapshotStore (the asset blob plane; refs + prod-link assets are the floor)"
  origin_milestone: M10 (v1.2)
  first_deferred_on: 2026-06-06
  last_seen_in: roadmap.md:107 (re-signed at v1.5 design 2026-06-11)
  destination: "backlog (unscheduled)"
  reason_recorded: "re-signed → backlog at v1.5 design; its user-facing sting removed by v1.5's real-images-via-prod-links posture (asset plane stays on prod public links). NOT touched by M21."
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. No item appears in two distinct milestones' deferral ledgers. DEF-M21-03 and DEF-M21-04 were carried across iters *within* M21 (the iterative-shape next-iter queue), which is intra-milestone iteration, not cross-milestone repeat-deferral — and both are now fated at the milestone boundary.

## Fate-1 Investigation

### DEF-M21-01 — replayCmd conn-seam
- **Fate-1 (land now, complete) feasible:** no
- **If no:** genuinely architectural — a hermetic test requires refactoring `replayCmd` to accept an injectable connector, changing its signature + the `main` dispatch + touching the load-bearing replay path (>50 lines). Doing it as part of a close would be a load-bearing-path refactor unscoped by M21's gate. The replay path is exercised live every iter; the residual is the live-DB connection wrappers (pg 47%, route-forward) that are integration-only by nature. **Fate:** route to a future `replayCmd`-seam build iter (the natural home is M22, which executes the provision path and already touches `replayCmd`/auto-provision wiring — but it is not a hard M22 dependency; recording as a tracked follow-up, not annotating M22's In: since M22 may not need the seam refactor to ship).

### DEF-M21-02 — serve-live-integration harness
- **Fate-1 (land now, complete) feasible:** no
- **If no:** needs a live directus Postgres container stood up — a test-infra build that exceeds a close-pass's inline boundary. The serve-row render is already hermetically unit-tested (serve_test.go + structure_harden_test.go) and hand-validated live in iter-05/iter-08, so the *correctness* is pinned; only the *automated* live harness is routed. **Fate:** route to the live-integration backlog. (M22 boots a live Directus per stack — the live harness becomes nearly free once M22 lands, so it is naturally picked up there or at M25 field-bake.)

### DEF-M21-03 — directus_files ref capture
- **Fate-1 (land now, complete) feasible:** no — and on inspection it should NOT land in M21 at all.
- **If no:** M21's gate is "serve a captured simulation anonymously" — met without `directus_files` (Directus introspects the served collections; the asset refs are a *content-resolution* concern, not a structure-serve concern). The asset plane is M23's explicit domain: v1.5 keeps `DIRECTUS_PUBLIC_BASE_ADDR` on prod public links so browsers fetch real images; `directus_files` rows are what let content rows resolve their image UUIDs to those prod URLs. M23 owns "keep the asset plane on prod" but its `In:` list does not currently name the `directus_files` capture explicitly. **Fate 3 — annotate M23's `In:` to pick it up.** (M21's overview `In:` named it, but the per-iter evidence showed it is orthogonal to the structure-serve gate and belongs with the asset-plane cutover.)

### DEF-M21-04 — M23 referential closure of the 20 dangling relations
- **Fate-1 (land now, complete) feasible:** no — explicitly out of M21 scope (overview `Out:` "the env re-point + referential closure (M23)").
- **If no:** **Fate 2 — already owned by M23.** M23's `In:` names "Referential closure — make the taxonomy capture include every node-id the captured content references … + a cross-surface fidelity gene." The 20 dangling relations are exactly that surface. No plan edit needed; confirm-covered.

### DEF-M10-01 — S3 blob bytes + cloud store
- **Fate-1 (land now, complete) feasible:** no
- **If no:** cross-release backlog, **re-signed fresh at v1.5 design (2026-06-11)** with its user-facing sting removed by the real-images-via-prod-links posture. Not aged (re-signed 2 days ago), not touched by M21, not a v1.5 milestone's scope. **Fate:** KEEP-DEFERRED (already on record at the release level; no fresh M21-close decision required — it is not an M21 item, it predates and is orthogonal to M21).

## Recommendations
- DEF-M21-01 → **LAND-NEXT (tracked follow-up iter)** — route-forward recorded in `hardening-ledger.md` (Fate 3, scope-boundary). Confirmed legitimately-future; no plan edit (no single milestone is the hard owner; it is a tooling-debt follow-up).
- DEF-M21-02 → **LAND-NEXT (live-integration backlog)** — route-forward recorded in `hardening-ledger.md`. Naturally picked up once M22 boots a live Directus / at M25 field-bake.
- DEF-M21-03 → **LAND-NEXT (Fate 3 → annotate M23)** — add `directus_files` ref capture to M23's `In:` list (asset-plane plumbing).
- DEF-M21-04 → **LAND-NEXT (Fate 2 → confirmed covered by M23)** — no plan edit; M23 already owns referential closure.
- DEF-M10-01 → **KEEP-DEFERRED** — release-level backlog item, re-signed fresh 2026-06-11; orthogonal to M21.

## Applied Changes
- **M23 `overview.md`** — added `directus_files` ref capture to the `In: (rosetta-extensions)` list (Fate 3 annotation for DEF-M21-03).
- **M21 `decisions.md`** — recorded the close-audit fate verdicts as M21-D14 (audit trail of the 5 fate decisions).
- No code changed (audit surfaces; the calling skill's build phase implements, and these are all LAND-NEXT/KEEP-DEFERRED — nothing lands in M21).

## Blocking Items (require user decision)
None. Zero repeat deferrals, zero aged-out items. All five items have clear single-pass fates; the two route-forwards are confirmed legitimately-future per the orchestrator's pre-flight expectation. **GREEN — close proceeds.**
