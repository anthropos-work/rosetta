**Type:** tik — under TOK-01 (coverage-drive strand), applying the user-chosen M48 CLOSED-CYCLE strategy.

# iter-07 — closed-cycle / frozen-snapshot strategy (the platform perf-wall alternative)

## What was attempted
Per the overview Phase plan, the M48-documented closed-cycle alternative to the iter-06 platform perf wall:

- **Root-cause CONFIRMED via platform source read (zero platform edit):**
  - `app GetAIReadinessWithOptions` (`ai_readiness.go:283-301`): the FAST frozen-snapshot branch
    (`buildResponseFromSnapshots` — no live recompute, no `withSkillerLang` translation N+1) is reached ONLY
    when `opts.CycleID != nil` AND `cycle.Status == "closed"`. The DEFAULT GET (`CycleID == nil`, line 301)
    ALWAYS calls `buildLiveResponse` (the slow wall).
  - FE `AIReadinessClient.tsx:131`: `effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id`
    -> `useAIReadiness({ cycleId })` passes `?cycle=<id>`. So with NO active cycle + a closed cycle, the FE
    *should* pick the closed cycle -> the fast path.

- **Implemented (authoring copy):**
  - `ai_readiness_config.go` -- cycle status `active -> closed` (+ `closed_at` + a representative `final_score`).
  - `ai_readiness_funnel.go` -- write a frozen `ai_readiness_snapshots` row per stage>=1 member, scored via the
    platform's OWN model (computeAIReadiness/classifyV2: step1/2/3, knowledge/usage axes, archetype/buckets,
    stage-gated). Heroes preserved: Aria stage3 COMPLETED / Ben stage1 STARTED / Dana (manager) no snapshot.
  - `cmd/stackseed/main.go` -- added the 8 `ai_readiness_*` tables to `resetTables` (so a `--reset` provably
    clears the OLD active-cycle row before the closed re-seed) + baked `--reload-sentinel`
    (FIX-M51-iter07-sentinel-reload-after-seed): a post-seed Sentinel Reload RPC (offset 8087) with a
    docker-restart fallback, so a live-stack re-seed's g2/g3 grants take effect without a manual restart.
  - Tests: config test -> asserts closed + closed_at + final_score; new funnel test -> asserts frozen snapshots
    (heroes at correct stages, score=sum-of-steps, stage-gating, valid enum values). gofmt/vet/build/full-suite GREEN.

- **Re-seeded demo-1** (`--reset` + seed + `--reload-sentinel`). DB VERIFIED: cycle **closed**, **199 frozen
  snapshots** (78.4% stage-3 -- the ~80% target), Aria=stage3/champion, Ben=stage1/standby, Dana no snapshot,
  archetype spread believable. `--reload-sentinel` worked (RPC unreachable on host -> docker-restart fallback OK).

- **Cheap probe FIRST (the iter-06 lesson) + GATED sweep:**
  - GATED manager sweep held at `failingSections=5, escapes=0, persona=0, reachable=66, frontier EXHAUSTED`.
    The 2 AI-readiness sections (`ai-readiness-org-score`, `ai-readiness-funnel`) STILL render **skeleton**.
  - Backend-log probe: **ZERO** `GET /api/workforce/ai-readiness` completions across the whole window (the
    iter-06 never-completing-GET signature).
  - Focused AUTHENTICATED network probe (login as Dana, watch the dashboard's requests): the data GET fires to
    the correct offset `18082` **WITHOUT `?cycle=`** (the LIVE path) and NEVER completes (hangs -> `networkidle`
    times out); the `/cycles` list GET is **never fired to the backend** at all.

## HYPOTHESIS -- PARTIALLY FALSIFIED
Closing the cycle + populating frozen snapshots is CORRECT and PROVEN in the DB, but it does NOT clear the wall,
because **the platform FE's default dashboard query does not route to the frozen path in this demo**: it fires the
live-recompute GET without `?cycle=` (-> `buildLiveResponse` -> the translation N+1 wall), and the `/cycles` list
query that would supply `latestClosedCycle.id` never reaches the backend. The frozen path is reachable in code
(`buildResponseFromSnapshots`) but the DEFAULT call never takes it. Closing that gap needs a PLATFORM edit -- make
the FE always pass the latest-closed cycle, OR make the backend default prefer a closed cycle when no active cycle
exists -- both forbidden by the zero-platform-edit invariant.

## Close -- 2026-07-01

**Outcome:** the closed-cycle showcase is correct + proven in the DB (199 frozen snapshots, 78.4% stage-3, heroes
right), and the seeder/reset/reload machinery landed clean + GREEN -- but the GATED sweep held at 5 (unchanged): the
platform FE's default GET fires the LIVE path (no `?cycle=`) and never completes, so the 2 AI-readiness sections
stay skeleton. Hypothesis PARTIALLY FALSIFIED with full evidence (the frozen path exists but the default call
doesn't take it; the fix is platform-bound).
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: y -- (5) cap-reached: n -- (6) protocol-stop: n -- Outcome: exit-4
**Decisions:** D1 (root-cause confirm -- the closed path is CycleID-gated + the FE cycle-preference), D2 (the
closed-cycle seeder + reset + --reload-sentinel -- landed clean, KEPT: correct data-in-DB regardless of the FE
route), D3 (the falsification + the user-blocker: the FE default doesn't route to the frozen path; the disclosed
fallback needs sign-off).
**Side-deliverables (if any):** `--reload-sentinel` (FIX-M51-iter07-sentinel-reload-after-seed) is a genuine
seeding-flow strengthening -- a live-stack re-seed now reloads the Sentinel policy itself (no manual restart). Kept.
**Routes carried forward:**
  - The 2 AI-readiness sections + the 3 workforce-aggregate sections -> blocked on a USER DECISION (below).
    Handler once decided: `DISCLOSE-M51-iter08-aireadiness-frozen-fe-route` (disclosed-presenter-note) OR
    `PLATFORM-M51-aireadiness-default-cycle` (escalated platform edit).
**USER-BLOCKER (the decision the user must make):** the closed-cycle strategy produced correct, fast-readable data
(the frozen snapshots), but the platform FE never requests the frozen path (it fires the live-recompute GET with no
`?cycle=`, which is the wall). Every path to `(0,0)` now needs a user/architectural decision:
  (a) **DISCLOSED residual** (the session-brief fallback): the data is PROVEN correct in the DB (199 frozen
      snapshots, funnel, heroes) -- the section is slow-but-correct due to a platform FE/read-path routing behavior,
      not a seed gap. Disclose it as a presenter-note per the coverage-protocol's narrow disclosed-allow -- but that
      rule needs the user's EXPLICIT sign-off (it's not an editorial-citation auto-allow). The gate then reaches
      green-with-disclosure. **The seeded closed-cycle data STAYS** (it's the honest, correct showcase; a demo
      presenter who opens the cycle picker / a `?cycle=` deep-link sees the fast frozen dashboard).
  (b) **ESCALATE a platform edit** (`unimplementable-without-platform-edit`, the milestone Re-scope trigger): make
      next-web's default dashboard query pass the latest-closed cycle id, OR make `app`'s default GET prefer a
      closed cycle when no active cycle exists (route to `buildResponseFromSnapshots`). The invariant forbids doing
      it here.
  (c) a NEW app read-path demo-patch that batches/relaxes the live translation N+1 (the `app-targetrole-authz-skip`
      precedent) -- a substantial new tooling investment; option B of the session brief, NOT chosen.
iter-07 surfaces the decision rather than picking one (the invariant forbids a unilateral platform edit; (a) needs
sign-off).
**Lessons:**
  - A cycle-scoped fast read-path only helps if the DEFAULT client call SELECTS it. Confirm BOTH sides -- the
    server branch AND the FE's request shape (does it pass the scoping param by default?) -- before assuming a
    closed cycle short-circuits the wall. Here `buildResponseFromSnapshots` exists but the default GET is
    hardcoded to `buildLiveResponse` (nil CycleID) and the FE's default fires without `?cycle=`.
  - Diagnose the FE request shape with an AUTHENTICATED network probe (log the outbound request URL + query
    params + whether it completes), not just the backend completion log -- the backend log tells you IF a request
    completed; the FE probe tells you WHICH variant the client fired (live vs cycle-scoped). Added to
    coverage-protocol.md.
