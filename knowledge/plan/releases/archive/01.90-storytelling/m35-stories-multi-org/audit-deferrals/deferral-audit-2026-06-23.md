---
title: "Deferral Audit — M35 (Stories & Heroes model + multi-org)"
date: 2026-06-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic patterns; no aged-out items requiring re-fate.
- Every inherited deferral routed into M35 LANDED IN FULL; M35 wrote zero new deferrals.

## Summary
- Total deferrals in scope: 1 inherited (M34→M35) + 4 release-level backlog (orthogonal, inherited GREEN at v1.9 design)
- Single deferrals: 0 open (the 1 inherited M34→M35 item was LANDED, not re-deferred)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

### Inherited into M35 (the close re-audit target)

- id: DEF-M34-07 (= #M34-D7)
  item: "Multi-hero index-collision guard (`len(Personas) <= Size` validation + collision warning) + short-role-pool flat top-up"
  origin_milestone: M34
  first_deferred_on: 2026-06-23
  last_seen_in: m34/decisions.md §D-M34-7 (Fate-3 → M35); m34/retro.md §Carried Forward
  destination: M35 (the roster milestone that owns the multi-hero trio + trajectory logic)
  reason_recorded: "benign-by-construction for M34's one-hero slice; reachable once M35 scales to the thriving/struggling/manager trio across stories"
  partial_attempted: no
  **RESOLUTION: LANDED IN FULL as D-M35-4.** Both parts shipped:
  (a) `len(heroes) <= size` blueprint validation (`blueprint/blueprint.go:390`) + a non-fatal residual-clamp
      warning (heroes occupy the first `len(heroes)` population slots IN DECLARATION ORDER — collision-free by
      construction, the strictly-better fix over warning-on-hash); covered by `TestValidate_MultiStory_RosterExceedsSize`.
  (b) short-role-pool flat top-up (`resolveHeroSkills` — role-coherent skills FIRST, flat padding SECOND to hit
      the declared `verified: N`); covered by the `resolveHeroSkills` branch tests in `jobroleref_test.go`.

### Release-level backlog (inherited; confirmed orthogonal at v1.9 design, re-confirmed here)

- DEF-M10-01 (cloud SnapshotStore / S3 media blob bytes) — backlog (unscheduled). Asset-plane work; orthogonal to seeding. No M35 touch.
- DEF-M21-01 — backlog. Snapshot/local-Directus area; orthogonal to seeding. No M35 touch.
- M25-D9 — backlog. Snapshot capture area; orthogonal to seeding. No M35 touch.
- M33 — backlog. Orthogonal to seeding. No M35 touch.

All four were audited GREEN at v1.9 design Phase-0 (roadmap.md §"In Development — v1.9": "the 4 backlog items
… are all orthogonal to seeding; no repeat-deferral dodged"). M35 touched none of their feature areas, so no
aging trigger fired (no substantive area-touch). They remain correctly parked; not in M35's scope.

## Repeat-Deferral Patterns
None. No item appears in ≥2 milestones' deferral ledgers without resolution. DEF-M34-07 appeared once (M34
close) and was resolved at its first destination (M35), the canonical healthy path.

## Aging
No AGED_OUT items. DEF-M34-07 was deferred 2026-06-23 (M34 close) and landed 2026-06-23 (M35) — zero days in
limbo, single milestone hop, destination delivered. The 4 release-backlog items have not aged into M35's scope
(none deferred ≥3mo within this release context; none had a destination milestone close without landing; M35
touched none of their areas).

## Fate-1 Investigation

### DEF-M34-07 — "Multi-hero index-collision guard + short-role-pool top-up"
- **Fate-1 (land now, complete) feasible:** yes — and it WAS landed in full during M35 build.
- **Landing scope (delivered):** D-M35-4 — declaration-order collision-free slot assignment + `len(heroes) <= size`
  validation + residual-clamp warning; `resolveHeroSkills` role-coherent-then-flat top-up. Tests:
  `TestValidate_MultiStory_RosterExceedsSize`, the `resolveHeroSkills` branch suite. No slice / no TODO remainder.

### DEF-M10-01 / DEF-M21-01 / M25-D9 / M33 — release backlog
- **Fate:** KEEP-DEFERRED (escape-hatch, signed at prior closes / v1.9 design). Not re-fated here — no aging
  trigger fired within M35's scope (M35 touched none of these areas; the seeding milestone is orthogonal to the
  snapshot/asset/Directus surfaces these items live in). Correctly out of M35.

## Recommendations
- **DEF-M34-07 → LAND-NOW: DONE.** Already landed in full as D-M35-4. No action.
- **DEF-M10-01 / DEF-M21-01 / M25-D9 / M33 → KEEP-DEFERRED (no change).** Orthogonal to M35; sign-off inherited
  from v1.9 design Phase-0. No fresh fate required (no aging trigger).

## Applied Changes
None. The audit is read-only confirmation: the one inherited deferral was already resolved in build; the
release-backlog items are orthogonal and need no re-fate. No files edited.

## Blocking Items (require user decision)
None. GREEN.
