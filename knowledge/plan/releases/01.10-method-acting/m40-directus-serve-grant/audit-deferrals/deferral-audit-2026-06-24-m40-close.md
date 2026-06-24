---
title: "Deferral Audit — milestone M40 (close)"
date: 2026-06-24
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals. The single deferral in scope has a clear, confirmed Fate-2 destination.

## Summary
- Total deferrals in scope: 1
- Single deferrals: 1
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

## Deferral Inventory

```yaml
- id: DEF-M40-01
  item: 'KPI "AI simulations completed" = 0 on the profile/dashboard surface'
  origin_milestone: M40
  first_deferred_on: 2026-06-24
  last_seen_in: m40-directus-serve-grant/decisions.md:14 (M40-D7); progress.md:30-33
  destination: "Fate-2 → M42e (employee coverage) + M42m (manager coverage)"
  reason_recorded: >
    Source public.local_jobsimulation_sessions (=21 seeded) has NO CMS dependency —
    the feed fix is a Directus-serve gap; the KPI reads jobsimulation directly. A
    separate frontend/auth-context concern, genuinely a different surface (frontend
    KPI vs CMS serve). The right home is the coverage sweep that exercises every page,
    not this serve-grant.
  partial_attempted: no
```

Cross-checked sources (all v1.10 milestones in scope):
- M39 (closed): deferral re-audit was GREEN at its own close; its two `Out:` items were Fate-2
  (already owned by M40/M41, confirmed, no edit) — **not deferrals**. Zero defer-language in M39 decisions.md.
- M40 code TODO/FIXME/HACK scan (the 3 M40-touched rext files): **none** (the one grep hit at
  `serve_test.go:217` is a descriptive comment of the bug under test, not a marker).
- M41/M42e/M42m: planned, unbuilt — no accumulated deferrals.

## Repeat-Deferral Patterns
None. DEF-M40-01 has a single occurrence in a single milestone.

## Fate-1 Investigation

### DEF-M40-01 — KPI "AI simulations completed" = 0
- **Fate-1 (land now, complete) feasible:** no
- **Why not Fate-1:** M40 is a `stack-snapshot` Directus serve-grant milestone (replay-side synthesis of
  public-read permissions/relations/fields). The KPI=0 surface is **not** a Directus-serve gap — its data source
  (`public.local_jobsimulation_sessions`, 21 rows already seeded) has **no CMS dependency**, so synthesizing more
  serve-grant rows cannot move the KPI. The residual is a frontend/auth-context rendering question, a different
  subsystem from this milestone's domain. Landing it here would mean reaching outside the serve-grant scope into
  the frontend read path — not a "land a slice of the serve-grant" question.
- **Which fate applies:** **Fate-2** — already owned by M42e + M42m. Their `exit_gate` is a Playwright sweep of
  **EVERY reachable demo page** asserting **non-empty semantic content in the DOM for 100% of pages, 0 failing**.
  A KPI rendering "0" on a reachable page is precisely an empty-section failure the gate catches; M42e's approach
  sketch routes "empty section / missing seed" → `stack-seeding` and "federation / content error" → `stack-snapshot`
  serve-grants. The KPI=0 surface therefore falls inside the gate's commitment as-written.
- **What failed last time:** N/A — first (and only) audit pass for this item; deferred today.

## Recommendations

| Item | Verdict | Mapping |
|------|---------|---------|
| DEF-M40-01 (KPI=0) | **LAND-NEXT** | Fate-2 — already owned by M42e/M42m's exit gate ("every reachable page renders non-empty content, 0 failing"). **No `In:`-list edit needed** — both are iterative; the gate is the commitment and the path emerges per-tik. Fate-3 (annotate) is NOT warranted: the gate as-written already encompasses a non-zero completed-KPI. |

## Applied Changes
- None to the roadmap/overviews. M40-D7 already records the Fate-2 routing; this audit **confirms** it (Fate-2 =
  confirm-don't-edit). No edit to M42e/M42m `overview.md` `In:` lists (iterative gate already encompasses the item).
- This report written to the M40 `audit-deferrals/` subfolder.

## Blocking Items (require user decision)
None. GREEN — close-milestone proceeds.
