---
iteration_type: tik
iter_shape: standard
status: closed-no-lift
---

# iter-07 — closed-cycle / frozen-snapshot strategy (clear the AI-readiness perf wall)

**Type:** tik — under TOK-01 (coverage-drive strand), applying the user-chosen M48-documented
CLOSED-CYCLE alternative to the iter-06-root-caused platform perf wall.

## Active strategy reference
TOK-01 (active-cycle signals-true) — the coverage-drive strand (strand 4). This iter applies the
**M48 closed-cycle refinement** the user chose this session: iter-06 root-caused the residual 5
fails as a PLATFORM perf wall (`GET /api/workforce/ai-readiness` never completes in-budget at 200
members — live-recompute + per-skill federated translation N+1). The M48 contract documents the
alternative: seed the cycle CLOSED with frozen per-member snapshots so the dashboard reads
pre-computed data instead of recomputing live.

## Re-survey (Step 0)
Baseline re-confirmed from iter-06 close: `failingSections=5, escapes=0, persona GREEN, reachable=65`,
frontier-exhausted. Demo-1 UP @ fit-up-m50, Northwind seeded with an ACTIVE cycle (156/199 stage-3 =
78.4%), snapshots=0, live_snapshots=0. Target still meaningful (5 residual, all the AI-readiness +
workforce-aggregate perf-wall sections).

## Cluster / target identified — the closed-cycle read-path pivot
**Root-cause CONFIRMED via platform source read (zero platform edit, read-only):**
- `GetAIReadinessWithOptions` (`app/internal/workforce/ai_readiness.go:283`): `if opts.CycleID != nil`
  AND `cycle.Status == "closed"` → `buildResponseFromSnapshots` (FAST — reads frozen
  `ai_readiness_snapshots`, NO `withSkillerLang`, NO live recompute, NO translation N+1). Otherwise →
  `buildLiveResponse` (the slow wall).
- FE (`AIReadinessClient.tsx:131`): `effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id`
  → `useAIReadiness({ cycleId: effectiveCycleId })` → passes `?cycle=<id>`. So with **NO active cycle**
  but a **closed** cycle present, the FE passes the closed cycle's id → the backend takes the FAST
  frozen path.

So: seed the cycle **closed** (no active cycle) + populate `ai_readiness_snapshots` frozen rows for
all 200 members → the dashboard reads frozen → loads fast → the 5 skeleton-fails clear.

## Hypothesis
Flipping the Northwind cycle to `closed` + writing frozen `ai_readiness_snapshots` (per-member, from
the same stage/step signals the funnel seeder already derives) makes `GET /api/workforce/ai-readiness`
complete fast → the 5 residual sections hydrate → `failingSections 5 → 0`.

## Expected lift
`failingSections 5 → 0`, escapes 0, persona GREEN, frontier-exhausted. (The AI-readiness org-score +
funnel + the workforce-aggregate grids all read the now-fast response.)

## Phase plan (coverage-protocol A–E)
- **Phase A — cheap probe FIRST (the iter-06 lesson):** after seeding the closed cycle, curl
  `GET /api/workforce/ai-readiness?cycle=<id>` and grep the backend log for a COMPLETED GET (not just
  OPTIONS). If it completes fast → hypothesis holds → proceed. If STILL slow → hypothesis FALSIFIED →
  the disclosed-presenter-note fallback (documented, user-surfaced) per the session brief.
- **Phase B — implement:** config seeder → cycle `closed` (+ closed_at + final_score); funnel seeder →
  write frozen snapshots for every stage≥1 member (heroes preserved: Aria stage3 COMPLETED / Ben
  stage1 STARTED). Unit tests. Re-seed demo-1 from the authoring copy (+ Sentinel reload).
- **Phase C — carried item:** bake Sentinel-reload-after-seed into the seeding flow
  (`FIX-M51-iter07-sentinel-reload-after-seed`).
- **Phase D — GATED sweep** (manager vantage, Northwind) → record `(failing, escapes)` delta.
- **Phase E — close** with the mandatory Gate field.

## Escalation conditions
If the closed cycle STILL recomputes / times out (hypothesis falsified) → do NOT do the heavy
read-path demo-patch (option B not chosen) → apply the disclosed-presenter-note fallback + DOCUMENT
prominently + surface to the user. If closing the cycle needs a platform edit → re-scope-trigger.

## Acceptable close-no-lift outcomes
Hypothesis-falsified-with-full-evidence (the closed cycle also recomputes) → closed-no-lift +
disclosed fallback, surfaced to user.
