**Type:** tik — under TOK-01 (active-cycle signals-true). The coverage-drive strand.

# iter-06 — re-sweep with the corrected harness; root-cause the residual AI-readiness perf wall

## What was attempted
Per the overview Phase plan, against LIVE demo-1 (UP @ fit-up-m50, AI-readiness showcase org seeded):

- **Phase A (re-sweep #1):** GATED manager sweep with the committed iter-05 harness in effect (report
  generatedAt 21:29Z) → `failingSections=5, escapes=0, persona=0, reachable=66, frontier EXHAUSTED`. The iter-05
  org-agnostic email correction WORKED — the 2 prior `cervato-systems.com` false-fails (members-roster,
  assign-roster) are GONE. The 5 residual: verification-funnel + talent-languages (/enterprise/workforce),
  ai-readiness-org-score + ai-readiness-funnel (the iter-05 AI-readiness descriptor — now a real asserted
  residual), activity-table (/enterprise/activity-dashboard).

- **Phase B (triage):** screenshots + DOM + DB + the protocol's diagnostic probe. The AI-readiness page renders
  the title + Northwind org + Dana chrome but a skeleton body; the Workforce Intelligence page renders ALL
  structure (title, Live badge, every tab + card header) but skeleton data regions. DB confirms the data is
  present (1 active cycle, 8 skills, 2 sims, 3 steps, 532 completed user_step_progresses — a 199→177→156 funnel;
  live_snapshots=0). Diagnosed PERF-not-content (D1).

- **Phase C (fix attempt):** D2 — extended `warmHeavyGrids` (added `/enterprise/workforce/ai-readiness` to the
  warm set + deepened the hydration check from first-real-row to skeletons-CLEARED, so a multi-card page's
  chrome doesn't bail the warm while its data cards are still cold). Cache-priming, NOT poll-widening (the
  protocol bans the latter). Authored in the authoring copy + transpile-verified.

- **Phase D (re-sweep #2):** GATED manager sweep with the deepened warm (report generatedAt 21:53Z) → STILL
  `failingSections=5` UNCHANGED, same 5. The warm did not clear the AI-readiness page.

- **Phase B′ (deep root-cause, D3):** the decisive evidence — `GET /api/workforce/ai-readiness` **NEVER
  completes** (0 completions in the entire backend log; only OPTIONS preflights); the `ai_readiness_refresh`
  background worker times out `context deadline exceeded` (4×, why live_snapshots is empty); ALL the
  AI-readiness SQL is ms-fast (queryReadinessSimScores EXPLAIN = 1.4ms, jobsimulation.sessions=1579 rows fully
  indexed; NOT the M46 index-bound wall); the default dashboard handler ALWAYS recomputes live
  (`buildLiveResponse`, no snapshot short-circuit); skiller logs show a per-skill `_entities` translation
  storm (`get skill translation <uuid>/english … context canceled`). The residual is a **platform-side
  AI-readiness response-build perf wall** (live recompute + per-skill federated translation fan-out exceed the
  request deadline at 200 members) — same N+1 class as the M46 per-object Sentinel RPC, in the translation path.
  NOT seed, NOT harness, NOT index, NOT demo-snapshot-warmable.

## Close — 2026-06-30

**Outcome:** the corrected harness CONFIRMED clean (the 2 cervato false-fails cleared; the AI-readiness
descriptor now a real asserted residual); the deepened skeletons-cleared warm landed clean but did NOT lift
the metric (5→5) because the residual is a PLATFORM-side server-deadline perf wall, not a cold-cache the warm
can prime. Root-caused with a complete falsification (the warm hypothesis is FALSIFIED for this residual; the
true cause is the live-recompute + per-skill federated translation fan-out exceeding the request deadline).
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: y — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D1 (perf-not-content triage), D2 (the cache-priming warm-deepening — landed clean, kept), D3 (the root-cause platform perf wall + the escalation).
**Side-deliverables (if any):** the deepened `warmHeavyGrids` (skeletons-cleared hydration + the ai-readiness route in the warm set) is a genuine harness STRENGTHENING — kept (it helps any heavy grid whose wall is cold-cache rather than a hard server deadline). It did not lift THIS metric but it is correct + committed.
**Routes carried forward:**
  - The 5 residual sections (the AI-readiness + workforce-aggregate perf wall) → blocked on a USER DECISION
    (see USER-BLOCKER below). Handler: `PERF-M51-iter07-platform-aireadiness-wall` once the decision lands.
  - **Bake Sentinel-reload-after-seed into the seeding flow** (the real fix surfaced in iter-05) — STILL
    carried (not done this iter; iter-06 went deep on the perf wall). Handler:
    `FIX-M51-iter07-sentinel-reload-after-seed`.
**USER-BLOCKER:** the AI-readiness dashboard GET never completes in-budget at 200 members — a platform-side
response-build perf wall (live recompute + per-skill federated translation fan-out > request deadline). The
gate cannot reach `(0,0)` until this is addressed, and EVERY addressing path is a user/architectural decision
that changes what code lands:
  (a) a NEW app read-path demo-patch (batch/relax the per-skill translation resolution, OR make the default
      dashboard call read the materialized `ai_readiness_live_snapshots` so it skips the live recompute) — the
      `app-targetrole-authz-skip`/`next-web-members-pagination` precedent class, but a substantial new tooling
      investment (a new demo-patch + its inject-loop wiring + a re-seed that populates live_snapshots);
  (b) ESCALATE a platform fix (`unimplementable-without-platform-edit` per the milestone Re-scope trigger — the
      M46 "platform needs a DataLoader / batch RPC" finding's twin);
  (c) accept a DISCLOSED residual: the dashboard renders correctly given enough time (data all present +
      correct) but exceeds the harness budget at 200 members → a disclosed-presenter-note (data proven correct
      in the DB; the section is slow-but-correct) — but per the protocol's narrow disclosed-allow this needs
      the user's explicit sign-off (it is NOT an editorial-citation case the rule auto-covers).
The invariant forbids a unilateral platform edit; (a) is a major scope decision; (c) needs sign-off. So
iter-06 surfaces the decision rather than picking one. Full evidence in decisions.md D3.
**Lessons:**
  - DIAGNOSE a never-completing GET at the SERVER before assuming a cold-cache the warm can prime: a `0
    completions in the backend log` + a timing-out background worker means the server can't produce the
    response in-budget — a warm (which only primes a result that WOULD complete) cannot help. Check the backend
    request log for completions BEFORE investing in a warm-deepening.
  - A materialized-snapshot mirror only helps if the read path CONSULTS it. Here the default dashboard call
    always recomputes live (the snapshot path is gated behind a closed CycleID), so populating
    ai_readiness_live_snapshots from the seed would NOT short-circuit the default GET — the snapshot
    short-circuit itself would be a platform-read-path change.
  - The per-skill federated translation fan-out (`withSkillerLang` → skiller `_entities`) is an N+1 in the
    same family as the M46 per-object Sentinel RPC — a recurring org-scale platform pattern worth a protocol
    note (added to coverage-protocol.md).
