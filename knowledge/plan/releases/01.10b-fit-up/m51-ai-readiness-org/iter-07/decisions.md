# iter-07 — decisions

## D1 — root-cause confirm: the frozen path is CycleID-gated + the FE cycle-preference
**Context:** the user chose the M48 closed-cycle alternative to the iter-06 perf wall. Before implementing,
confirm (read-only, zero platform edit) that a closed cycle actually short-circuits the wall.
**Finding:** `app GetAIReadinessWithOptions` (`ai_readiness.go:283-301`) reaches the fast
`buildResponseFromSnapshots` ONLY when `opts.CycleID != nil && cycle.Status=="closed"`; the DEFAULT GET
(nil CycleID, line 301) is hardcoded to `buildLiveResponse`. The FE (`AIReadinessClient.tsx:131`) prefers
`selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id` and passes it as `?cycle=`. So the strategy is
architecturally sound IFF the FE fires the cycle-scoped GET.
**Choice:** proceed with the closed-cycle seed (necessary + correct), but VERIFY the FE request shape early
(the cheap probe) before assuming the wall clears.

## D2 — the closed-cycle seeder + reset + --reload-sentinel (landed clean, KEPT)
**Context:** implement the closed-cycle showcase in the authoring copy.
**Choice:**
- `ai_readiness_config.go`: cycle `active -> closed` (+ closed_at + representative final_score).
- `ai_readiness_funnel.go`: write a frozen `ai_readiness_snapshots` row per stage>=1 member, scored via the
  platform's own model (computeAIReadiness/classifyV2). Heroes preserved (Aria stage3 / Ben stage1 / Dana none).
- `cmd/stackseed`: the 8 `ai_readiness_*` tables added to `resetTables` (a --reset must clear the OLD active
  cycle before the closed re-seed); `--reload-sentinel` baked (FIX-M51-iter07 — post-seed Reload RPC + docker
  restart fallback).
**Why KEPT despite no metric lift:** the DB is now the CORRECT closed-cycle showcase (199 frozen snapshots,
78.4% stage-3, heroes right, believable archetypes) regardless of how the default FE routes. A presenter who
opens the cycle picker / a `?cycle=` deep-link sees the fast frozen dashboard. The data is honest + reusable;
the only gap is the platform FE's DEFAULT route. Tests + gofmt/vet/build/full-suite GREEN.

## D3 — the falsification + the USER-BLOCKER (the FE default doesn't route to the frozen path)
**Context:** cheap probe (backend log) + focused authenticated network probe.
**Finding:** ZERO `GET /api/workforce/ai-readiness` completions; the FE fires the data GET to the correct offset
`18082` **WITHOUT `?cycle=`** (the LIVE path) and it hangs; the `/cycles` list GET never reaches the backend.
So `latestClosedCycle` is never supplied to the FE, `effectiveCycleId` stays undefined, and the default GET takes
the live-recompute wall. The frozen path EXISTS but the default call doesn't take it.
**Why user-blocker (Phase 5 §4):** closing the gap needs a PLATFORM edit (FE always pass latest-closed, OR
backend default prefer a closed cycle), forbidden by the zero-edit invariant. The session-brief fallback is the
disclosed-presenter-note, which per the coverage-protocol's narrow disclosed-allow needs the user's EXPLICIT
sign-off. So iter-07 surfaces the decision (a=disclose+sign-off / b=escalate-platform / c=new-app-demo-patch)
rather than picking one. Full options in progress.md's USER-BLOCKER block.
