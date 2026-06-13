---
title: "Deferral Audit — milestone M24"
date: 2026-06-13
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- **M24 introduced ZERO new deferrals.** All four hygiene items the deferral audit had surfaced (NEW-5 Go pin,
  NEW-6 README index-row guard, NEW-11 zero-critical-genes guard, NEW-14 `/project-stats` scope fix) **landed
  Fate-1 in M24** — the milestone exists to *retire* aged-out hygiene backlog, and it did. The corpus truth-up
  (§1–§3) is all Fate-1 doc work, nothing routed.
- **No repeat deferrals.** No item appears unresolved in two distinct milestones' ledgers. M24 carried nothing
  forward of its own.
- **The three standing items are unchanged and not aged by M24.** DEF-M21-01, DEF-M21-02, DEF-M10-01 stand exactly
  as the M23-close audit fated them; M24 touched none of their code areas (verified: M24's ext diff `7e9343a..d01a3ee`
  touches `alignment/` + `stack-core/` + the 4 go.mod `toolchain` pins — NOT `stack-snapshot/replay.go` (DEF-M21-01's
  seam) nor the serve-row render path (DEF-M21-02)). No destination milestone has closed. No aging trigger fires.
- **One M24 observation is correctly NOT a deferral.** M24-D3 records a pre-existing developer-kit `stats.sh`
  doc-counter limitation (rosetta's docs live under `corpus/`, the counter only looks at `knowledge/`) — surfaced
  out of §7 scope, recorded so it isn't silently lost, explicitly *not* a v1.5 deliverable. It is an observation
  about a foreign shared tool, not a punt of M24 work.

## Summary
- Total deferrals in scope: 3 standing (2 inherited M21 follow-ups + 1 cross-release backlog)
- New deferrals from M24: 0
- Resolved/retired this milestone: 4 aged-out hygiene items (NEW-5/6/11/14 — all landed Fate-1; these were backlog
  hygiene items, not formal DEF-ledger entries, but M24's charter was to clear them and it did)
- Single deferrals: 3
- Repeat deferrals: 0
- Chronic / aged-out patterns flagged: 0

## Deferral Inventory

```yaml
# --- STANDING (inherited, unchanged by M24, not aged) ---
- id: DEF-M21-01
  item: "replayCmd conn-seam — hermetic replayCmd wiring test needs an injectable connector"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  destination: "tooling-debt follow-up (replayCmd-seam build iter)"
  reason_recorded: "architectural, >50 lines; touches the load-bearing replay path"
  m24_status: "NOT touched by M24 — M24's ext work is alignment + stack-core + go.mod pins; replay.go untouched.
    Stands unchanged, not aged (the seam follow-up has no closed destination milestone; tracked tooling-debt)."
  partial_attempted: no

- id: DEF-M21-02
  item: "serve-live-integration harness — automated integration harness for serve-row render SQL"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  destination: "M25 field-bake (live observable-behavior gate)"
  reason_recorded: "needs a live directus Postgres; render is hermetically unit-tested + hand-validated live"
  m24_status: "Unchanged. M25 (its confirmed Fate-2 home) is the next + final milestone, still open. M24 is docs +
    hygiene — it does not touch serve-row render. Not aged."
  partial_attempted: no

- id: DEF-M10-01
  item: "S3 media blob BYTES + cloud SnapshotStore"
  origin_milestone: M10 (v1.2)
  first_deferred_on: 2026-06-06
  destination: "backlog (unscheduled)"
  reason_recorded: "user-facing sting removed by v1.5's real-images-via-prod-links posture; asset plane on prod"
  m24_status: "Unchanged. M24 touches no asset-plane code. The refs-only posture (re-affirmed by M23-D3) stands;
    cross-release backlog, re-signed 2026-06-11. No fresh decision required."
  partial_attempted: no

# --- M24 OBSERVATION (NOT a deferral — recorded for completeness) ---
- id: M24-OBS-statsdoc
  item: "developer-kit stats.sh doc-counter reports 0 for rosetta (it only scans knowledge/, rosetta docs live
    under corpus/); the eval-find PRUNE is fragile under some shells"
  origin_milestone: M24 (surfaced by §7, recorded M24-D3)
  first_seen_on: 2026-06-13
  destination: "NONE — a foreign shared-tool limitation, explicitly not a v1.5 deliverable; noted so it isn't lost"
  reason_recorded: "Independent of the §7 stack-*/ scanning bug (which WAS fixed). A pre-existing developer-kit
    plugin doc-counter limitation, outside M24's chartered scope. Not a punt of M24 work — M24 §7 landed in full."
  partial_attempted: "n/a — not an M24 task"
```

## Repeat-Deferral Patterns
None. No item appears in two distinct milestones' ledgers unresolved. M24 re-deferred nothing and resolved the
four hygiene-backlog items it was chartered to clear (NEW-5/6/11/14, all Fate-1).

## Fate-1 Investigation

### DEF-M21-01 — replayCmd conn-seam
- **Fate-1 (land now in M24) feasible:** no — and correctly so. M24 is a docs + hygiene milestone; the seam
  refactor is a >50-line architectural change to the load-bearing replay path, outside M24's chartered scope and
  not incidentally unblocked by anything M24 did (M24 didn't touch `replay.go`).
- **Fate:** KEEP (tracked tooling-debt follow-up). Unchanged; not unblocked, not required by M24.

### DEF-M21-02 — serve-live-integration harness
- **Fate-1 (land now in M24) feasible:** no — needs a live Directus Postgres + CI infra; that is precisely what
  M25 field-bake provides. M24 has no live-stack surface.
- **Fate:** LAND-NEXT (M25, Fate-2 confirmed-owned). Unchanged. M25 is the next + final milestone, still open —
  destination intact.

### DEF-M10-01 — S3 blob bytes + cloud store
- **Fate-1 (land now in M24) feasible:** no — cross-release backlog; orthogonal to M24's docs/hygiene scope. The
  v1.5 release posture is deliberately refs-only (real images via prod links), re-affirmed M23-D3.
- **Fate:** KEEP-DEFERRED (release-level backlog, re-signed 2026-06-11). No fresh M24 decision required.

## Recommendations
- DEF-M21-01 → **KEEP (tracked follow-up)** — unchanged; not unblocked or required by M24.
- DEF-M21-02 → **LAND-NEXT (M25 field-bake)** — unchanged; destination (next + final milestone) intact.
- DEF-M10-01 → **KEEP-DEFERRED** — release-level backlog, re-signed 2026-06-11; M24 touches no asset-plane code.
- M24-OBS-statsdoc → **NOT A DEFERRAL** — a recorded observation about a foreign shared tool; M24 §7 landed in
  full. No ledger action.

## Applied Changes
- None required to plan files. M24 introduced no new deferrals; the three standing items are already correctly
  recorded (M23-close audit) and need no re-fating or re-routing. The four hygiene items M24 was chartered to
  clear are reflected as landed Fate-1 in M24 progress.md §4–§7 + decisions.md (M24-D1/D2/D3). The M24-D3
  observation is recorded in decisions.md. The audit surfaces; nothing needed changing.

## Blocking Items (require user decision)
None. Zero new deferrals, zero repeat deferrals, zero aged-out items. The four chartered hygiene items landed
Fate-1; the three standing inherited items are unchanged with intact destinations and were not touched (so not
aged) by M24. **GREEN — close proceeds.**
