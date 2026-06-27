---
title: "Deferral Audit — M45 (generation engine) close"
date: 2026-06-26
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; the single in-release deferral (org-scale fill) is a Fate-2 route already owned
  by M46's `In:` list. No chronic patterns. No aged-out items (everything dated within the same week).

## Summary
- Total deferrals in scope: 2 (1 in-release Fate-2 route; 1 design-time cross-release scope-out)
- Single deferrals: 2
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

```yaml
- id: DEF-M45-01
  item: "org-scale auto-fill to full org size (fill the remaining N members of a story)"
  origin_milestone: M45
  first_deferred_on: 2026-06-26   # design-time Out: at M45 plan authoring
  last_seen_in: m45-generation-engine/overview.md:74 (Scope — Out)
  destination: "M46 (org-scale fill) — In: list owns it"
  reason_recorded: "M45 proves the engine + cache on a BOUNDED batch; org-scale is M46"
  partial_attempted: no

- id: DEF-M45-02
  item: "production-seeding key story (a platform-repo secrets store for prod-scale gen)"
  origin_milestone: M45
  first_deferred_on: 2026-06-26   # design-time Out: + Open question
  last_seen_in: m45-generation-engine/overview.md:78-79,95 (Out + Open questions)
  destination: "out of v1.10 release scope — .env.local OPENAI_API_KEY env var for now"
  reason_recorded: "production-seeding secrets deferred; key via git-ignored .env.local for now"
  partial_attempted: no
```

Note — the `spec-notes.md` `TODO:` lines (services/ai wrapper, blueprint.Batch, cmd/gen-batch, the cache,
GeneratedBatchSeeder, hero-collision, exit-gate measurement) are **build-plan scaffold markers, NOT
deferrals** — every one was DELIVERED (progress.md iter-02…iter-07 + the gate-met state). They describe done
work; the `TODO:` prefix is stale scaffold text, not an open item. No fate required (a nice-to-have docs
tidy, not a scope concern).

## Repeat-Deferral Patterns
None. Org-scale fill (DEF-M45-01) was designed into M46 in the **same v1.10-extend pass** that authored
M45 — it has never been deferred-then-not-done across milestones. M44's close (deferral-audit-2026-06-26-m44)
already confirmed the same Fate-2 route (LLM-generated content → M45/M46; M44 seeders deterministic by
design). The route is consistent across M44→M45→M46, not a repeat-defer.

## Fate-1 Investigation

### DEF-M45-01 — "org-scale auto-fill to full org size"
- **Fate-1 (land now, complete) feasible:** no
- **If no:** **Fate 2** — already owned by **M46** (`### M46` roadmap block + `m46-org-scale-fill/overview.md`
  `In:` list: "a supporting-population batch (`count: auto-fill to org size`, …) expanding to fill the
  remaining N members of a story so a 220/500/1k org is believable end-to-end"). A complete Fate-1 landing
  in M45 would contradict the milestone's deliberate boundary — M45's gate is explicitly a **bounded** N=20
  batch (proven: $0.0059 ≤ $0.10, 20/20 distinct names, closure GREEN). Population-scale behaviour (dedup at
  scale, taxonomy-clipping under load, throttle/backoff under burst, the believable spread) is M46's
  measure-driven iterative gate. Confirmed Fate-2. No plan edit needed — M46 `In:` already reflects it.

### DEF-M45-02 — "production-seeding key story"
- **Fate-1 (land now, complete) feasible:** no
- **If no:** **out of v1.10 release scope by design** — a platform-repo secrets store for production-scale
  generation is orthogonal to v1.10 ("method acting" = believable demo profiles; tooling + docs only, zero
  platform-repo edits). M45 keys via a git-ignored `.env.local` env var, which fully satisfies the demo/dev
  flow the release is about. This is a standing design-time scope boundary (an Open question logged for a
  future release), not an in-release punt of in-scope work — it does not block close and does not require an
  escape-hatch sign-off (no in-release scope is being broken). Tracked as a standing item, not a v1.10
  deferral.

## Recommendations
- **DEF-M45-01 → LAND-NEXT (Fate 2):** confirmed already owned by M46. No edit. Recorded in M45 close.
- **DEF-M45-02 → standing item (out of release scope by design):** no v1.10 action; remains an Open
  question for a future production-seeding effort.

## Applied Changes
- None to plan files — both routes are pre-existing (M46 `In:` already owns DEF-M45-01; DEF-M45-02 is a
  documented design-time Open question). This audit confirms, it does not re-route.

## Blocking Items (require user decision)
None. GREEN — close proceeds.
