---
title: "Deferral Audit — milestone M22"
date: 2026-06-13
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- **No repeat deferrals.** No item appears in two distinct milestones' deferral ledgers. M22 introduced
  **zero new deferrals** of its own (all 6 sections landed Fate-1; KB-1 was *resolved* in §6, not deferred;
  `backlog_refs: NEW-2` is the milestone's executed deliverable, not a punt).
- **No chronic / aged-out patterns.** The four inherited M21 items + DEF-M10-01 were all first recorded
  2026-06-06..13 (days old, not months); none crossed a *closed* destination milestone — M23 (the destination
  for the two content/asset items) is still `planned`, so its window is open, not aged-out.
- Every inherited item has a still-valid single-pass fate; M22 did not re-defer or re-scope any of them.

## Summary
- Total deferrals in scope: 5 (0 M22-originated + 4 inherited from M21 + 1 cross-release backlog item)
- Single deferrals: 5
- Repeat deferrals: 0
- Chronic / aged-out patterns flagged: 0

## Deferral Inventory

M22 originated **no** deferrals. The standing ledger (carried from the M21 close audit
`../../m21-structure-capture/audit-deferrals/deferral-audit-2026-06-13-m21-close.md`) is re-walked here:

```yaml
- id: DEF-M21-01
  item: "replayCmd conn-seam — hermetic replayCmd wiring test needs replayCmd refactored to accept an injectable connector"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  last_seen_in: m21 hardening-ledger.md:107-111
  destination: "future replayCmd-seam build iter (tooling-debt follow-up)"
  reason_recorded: "architectural, >50 lines; touches the load-bearing replay path"
  partial_attempted: no
  m22_status: "NOT touched by M22 — M22 executes the provision path via the existing CLIs (stacksnap/provision-plan) and did not refactor the replayCmd signature; the seam follow-up stands unchanged, not aged (5 days old)"

- id: DEF-M21-02
  item: "serve-live-integration harness — automated integration harness for the serve-row render SQL needs a live directus Postgres"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  last_seen_in: m21 hardening-ledger.md:112-115
  destination: "live-integration backlog (naturally picked up at M25 field-bake)"
  reason_recorded: "needs a live directus Postgres; render is hermetically unit-tested + hand-validated live"
  partial_attempted: no
  m22_status: "M22 BOOTS a live per-stack Directus (the precondition the M21 audit predicted would make this nearly free) but did NOT itself author the automated live-integration harness — M22's tests stub docker/CLIs (appropriate for the authoring harness). The harness lands at M25 field-bake (live observable-behavior gate). Unchanged, not aged."

- id: DEF-M21-03
  item: "directus_files ref capture — wire the dead media.go into directus.Surface() for the real-image asset refs"
  origin_milestone: M21
  first_deferred_on: 2026-06-11
  last_seen_in: M23 overview.md:45 (Fate-3 annotation landed at M21 close)
  destination: "M23 (asset plane / content cutover) — explicitly named in M23 In:"
  reason_recorded: "asset-ref plumbing, orthogonal to M21's structure-serve gate"
  partial_attempted: no
  m22_status: "Confirmed still owned by M23 (overview In: line 45, with backref to the M21 audit). M23 status=planned (open window). Not in M22 scope. Untouched."

- id: DEF-M21-04
  item: "referential closure of the 20 dangling relations"
  origin_milestone: M21
  first_deferred_on: 2026-06-11
  last_seen_in: M23 overview.md:42-44
  destination: "M23 (referential closure — explicitly owned, Fate-2)"
  reason_recorded: "a booted Directus has dangling relation defs until M23 closes them; not needed for the M21 gate"
  partial_attempted: no
  m22_status: "Confirmed still owned by M23 (overview In: lines 42-44). M23 status=planned. Not in M22 scope. Untouched."

- id: DEF-M10-01
  item: "S3 media blob BYTES + cloud SnapshotStore"
  origin_milestone: M10 (v1.2)
  first_deferred_on: 2026-06-06
  last_seen_in: roadmap.md:107 (re-signed at v1.5 design 2026-06-11)
  destination: "backlog (unscheduled)"
  reason_recorded: "user-facing sting removed by v1.5's real-images-via-prod-links posture; asset plane stays on prod public links"
  partial_attempted: no
  m22_status: "Orthogonal to M22 (M22 is the data-plane lifecycle; the asset blob plane is untouched). Re-signed 2 days before M22 close — not aged. No M22-close action."
```

## Repeat-Deferral Patterns
None. No item appears in two distinct milestones' deferral ledgers. M22 carried nothing forward of its own
and re-deferred none of the inherited items — it merely passed through, with M21-02's downstream destination
(M25) now further unblocked by M22 booting a live Directus, exactly as the M21 audit anticipated.

## Fate-1 Investigation

### DEF-M21-01 — replayCmd conn-seam
- **Fate-1 (land now in M22 close) feasible:** no
- **If no:** still genuinely architectural (>50 lines, load-bearing replay-path signature change). M22 executed
  the provision path *around* the existing CLIs without needing the seam refactor, so it neither landed nor
  required it. **Fate:** unchanged — tooling-debt follow-up. Not an M22-close item.

### DEF-M21-02 — serve-live-integration harness
- **Fate-1 (land now in M22 close) feasible:** no
- **If no:** an automated live-integration harness needs a live Directus container *and* CI infra to stand it up —
  test-infra build work that exceeds a close-pass's inline boundary, and belongs to the live observable-behavior
  gate (M25 field-bake) by design. M22 made it cheaper (a per-stack Directus now boots) but the harness itself is
  M25 scope. **Fate:** unchanged — LAND-NEXT, naturally picked up at M25.

### DEF-M21-03 — directus_files ref capture
- **Fate-1 feasible:** no — out of M22 scope (M22 is lifecycle/boot; asset refs are M23's content-resolution domain).
- **Fate:** confirmed **Fate-3 already landed** — M23 `overview.md:45` names it under `In:` with a backref. No action.

### DEF-M21-04 — referential closure of the 20 dangling relations
- **Fate-1 feasible:** no — explicitly M23-owned (overview `Out:` "referential closure (M23)").
- **Fate:** confirmed **Fate-2 covered by M23** (overview `In:` lines 42-44). No plan edit.

### DEF-M10-01 — S3 blob bytes + cloud store
- **Fate-1 feasible:** no — cross-release backlog, re-signed fresh 2026-06-11, orthogonal to M22's data-plane work.
- **Fate:** KEEP-DEFERRED (release-level record; no fresh M22-close decision required — not an M22 item).

## Recommendations
- DEF-M21-01 → **KEEP (tracked follow-up)** — unchanged; not unblocked or required by M22. No plan edit.
- DEF-M21-02 → **LAND-NEXT (M25 field-bake)** — unchanged; M22 furthered the precondition, M25 is the home.
- DEF-M21-03 → **LAND-NEXT (Fate-3, already annotated to M23)** — confirmed, no action.
- DEF-M21-04 → **LAND-NEXT (Fate-2, confirmed covered by M23)** — no plan edit.
- DEF-M10-01 → **KEEP-DEFERRED** — release-level backlog, re-signed 2026-06-11; orthogonal to M22.

## Applied Changes
- None. M22 originated no deferrals; all inherited items retain valid fates and correct destinations
  (verified live against M23 `overview.md`). The audit surfaces; nothing needed re-fating or re-routing.

## Blocking Items (require user decision)
None. Zero repeat deferrals, zero aged-out items, zero M22-originated deferrals. **GREEN — close proceeds.**
