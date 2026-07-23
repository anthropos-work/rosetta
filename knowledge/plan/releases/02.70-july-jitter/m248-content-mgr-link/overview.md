---
milestone_shape: section
milestone: M248
title: "content-stories manager result-link"
status: planned
release: v2.7 "july jitter"
depends_on: [M246]
parallel_with: [M247, M249, M250, M251, M252]
complexity: small
created: 2026-07-23
---

# M248 — content-stories manager result-link

## Goal
The content-stories **manager** CTA jumps to the **per-session manager result view**
(`/sim/<slug>/<userId>/result/<sessionId>`), not the org activity-dashboard scoreboard. Today the manager
seat lands on `/enterprise/activity-dashboard/<kind>/<simId>/<membershipId>` — a scoreboard, not the
individual hero's graded result; this milestone re-points that CTA at the real, already-built per-session
manager result page (which reads the persisted result row the `ContentStorySeeder` already plants).

## Shape (why this shape)
`section`. A bounded, well-understood re-point of one path builder + its downstream graders/docs/manifest —
no exploratory loop, no live measure→fix cycle. The one genuine unknown (does the interview manager report
live on the unified `/sim` route or stay split on `/activity-dashboard/interviews`?) is discharged as a
serial **rung-0 verify-interview read** at the front of the milestone, not an iterative gate. The `/sim`
manager view already exists and reads the seeded persisted row, so this is a projection/grader/doc change,
not a build.

## Scope
### In
- Change the sim `ManagerResultPath` builder (`content_manifest.go:411-423`) to
  `/sim/<slug>/<userId>/result/<sessionId>` (`owner.UserID` is already on the ownerSlot; the per-`sim_type`
  **kind** branching collapses). **verify-interview first** — the rung-0 read: does the interview manager
  report surface on the `/sim` route, or does it stay on `/activity-dashboard/interviews`? (decides whether
  interview keeps a split path).
- Update the e2e grader `content-result-page.ts:459` shape + the `content-route-contract.unit.spec.ts`
  prefixes; regenerate `presets/content-manifest.json` via `stackseed --content-export` (honesty gate);
  update `content-stories-spec.md` + `content-stories-routes.md`.

### Out
- Any platform / next-web edit — the `/sim` manager view **already exists** and reads the persisted result
  row the seeder plants (nothing to build platform-side).

## Dependencies & parallelism
- **depends_on:** M246 (the HARD go/no-go barrier — fan-out worktrees branch from post-M246 HEAD).
- **parallel_with:** M247, M249, M250, M251, M252 (a concurrent fan-out lane off M246).
- **Intra-milestone lanes:** low (~1.2×). The rung-0 **verify-interview** read is a **serial bottleneck** at
  the front; once it resolves, three lanes fan out concurrently — **core-projection** (the `ManagerResultPath`
  builder edit) ∥ **grader** (`content-result-page.ts` + `content-route-contract.unit.spec.ts` + the
  `--content-export` manifest regen) ∥ **docs** (`content-stories-spec.md` + `content-stories-routes.md`).
- **Recommended subagents:** small enough that **a single agent is defensible** — the ~1.2× ceiling doesn't
  repay the coordination overhead of a fan-out.
- **Cross-milestone coordination:** `cmd/stackseed/main.go` is the single seeder registry touched by **both
  M248 and M250** — each edits only its own `MustRegister`/truncate hunk (clean hand-merge). Any net-new
  `*.unit.spec.ts` must be rostered by **M251** (which owns `run-unit.sh` — coordinate the line). The
  `CLAUDE.md` one-line bullet defers to **M247** (sole owner). Merge/close order: M251 → { **M248**, M250 }
  → M249 → M253 → M252 → M247-reconcile → M254.

## KB dependencies
- `corpus/ops/demo/content-stories-spec.md`
- `corpus/ops/demo/content-stories-routes.md`

## Delivers
- `corpus/ops/demo/content-stories-spec.md` — the per-session manager result route as the correct CTA.
- `corpus/ops/demo/content-stories-routes.md` — the per-session manager result route as the correct CTA.

## Open questions
- Does the interview manager report render on the unified `/sim` route (else keep interview split)?
- Confirm `user.externalId` == the seeded `ownerSlot.UserID` at a live render.
