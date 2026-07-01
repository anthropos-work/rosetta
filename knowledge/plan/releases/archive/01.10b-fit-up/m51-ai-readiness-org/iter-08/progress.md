**Type:** tik (tooling-iter shape) — under TOK-01 (coverage-drive strand), applying the user-chosen
run-5 **DEEP-LINK THE DEMO ENTRY** strategy.

# iter-08 — deep-link the demo entry: verify the frozen fast path FIRST → FALSIFIED

## What was attempted
The run-5 brief chose a zero-platform-edit fix for the iter-07 residual: the closed-cycle data is seeded +
correct in demo-1 (199 frozen snapshots, cycle CLOSED, Aria stage3 / Ben stage1 / Dana manager, 78.4%
stage-3), and the platform "serves frozen data FAST via `buildResponseFromSnapshots` — reachable only when
the request carries the closed cycle id (`?cycle=<latest-closed>`)". The fix: point Dana's cockpit `jump_to`
+ the coverage manifest at the dashboard **with `?cycle=<closed-id>`** so the FE fires the cycle-scoped GET →
the fast frozen path → the sections clear.

The brief mandated **VERIFY THE DEEP-LINK HITS THE FROZEN FAST PATH FIRST (cheap probe before a full sweep)**.

### Step 1 — resolve the closed cycle id (DB ground truth)
`ai_readiness_cycles` on demo-1 (offset postgres :15432): the Northwind closed cycle is
**`95d9fc3d-48c0-53ca-82ac-e10713100c97`**, status `closed`, 199 frozen snapshots, 78.4% stage-3, Aria=stage3,
Ben=stage1, Dana no snapshot (manager). Matches the deterministic id
`deterministicUUID("ai-readiness-cycle:" + StoryOrgID("ai-readiness"))`.

### Step 2 — read the FE cycle-selection (confirms `?cycle=` IS read from the URL)
`AIReadinessClient.tsx`: `const selectedCycle = sp.get('cycle') || undefined;` →
`effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id;` → the data query
`useAIReadiness({ cycleId: effectiveCycleId, enabled: featureOn && cyclesQ.isFetched })`. So the FE **does**
read `?cycle=` from the URL — the deep-link IS wired to the FE. BUT the data query is **gated on
`cyclesQ.isFetched`** (waits for the `/cycles` list to settle) — and the backend log showed **0** `/cycles`
completions ever, so a first read of "the FE never fires `?cycle=`" was ambiguous between (a) `/cycles` gate
never clears and (b) the frozen read itself is slow.

### Step 3 — the CHEAP PROBE (authenticated dual-endpoint DIRECT probe) → FALSIFICATION
Authored `probe-aireadiness-deeplink.spec.ts` (a genuine reusable diagnostic): log in as Dana via the cockpit
handshake, lift her bearer token from a real outbound backend request, then hit BOTH endpoints **directly**
via `page.request` (bypassing the FE React Query orchestration) with timing:
- `GET /api/workforce/ai-readiness/cycles` → **200 in 40 ms** (fast — the FE gate is fine; returns the closed
  cycle).
- `GET /api/workforce/ai-readiness?cycle=95d9fc3d-…&includePeople=true` (the frozen path) → **NEVER COMPLETED**
  (timed out the full 180 s test budget) — identical to the live-recompute default.

### Root cause (zero platform edit, source read)
`app buildResponseFromSnapshots` (`ai_readiness.go:512`) reads the frozen SCORES fast (line 513
`ListAIReadinessSnapshots`) but then calls **`loadMembers(orgID, "")`** (line 520) — a **full unbounded
org-member hydration** (`hydrateMembers` with `memberIDs=nil, userIDs=nil` → whole-org tag/skill/sim
aggregation over ~200 members) to re-join CURRENT tags/name/role onto each snapshot. At 200 members that
member-load is the SAME org-scale wall as the live path; the `ai_readiness_refresh` worker (same compute) logs
`context deadline exceeded`. Crucially the cost is NOT the demo-patchable per-object targetRole Sentinel RPC —
`queryBaseMembers` reads `jobRole` from a SQL column (line 473), so the applied `app-targetrole-authz-skip`
patch does nothing for it.

## HYPOTHESIS — FALSIFIED (stronger than iter-07)
iter-07 assumed the frozen branch was fast and only the FE routing was the gap. iter-08 PROVES the frozen READ
itself is org-scale-slow: even a direct authenticated `?cycle=<closed>` GET times out. So the deep-link — the
user-chosen zero-edit fix — **cannot clear the wall even in principle**, because "frozen" froze the *scores*,
not the *response* (the response re-joins live per-member identity → re-incurs the member-load wall). The
deep-link cockpit/manifest edits were therefore NOT made (they would ship an inert `?cycle=` that still hangs)
and the GATED sweep was NOT run (it would only reconfirm `failingSections=5`).

## Close — 2026-07-01

**Outcome:** the user-chosen deep-link strategy is FALSIFIED by a cheap authenticated dual-endpoint probe:
`/cycles` is fast (40 ms) but the frozen data GET `?cycle=<closed>` NEVER completes (180 s), because
`buildResponseFromSnapshots` re-loads the whole org's members (`loadMembers`) — the same org-scale wall,
NOT a demo-patchable authz gate. `failingSections` unchanged at 5 (no sweep run — the deep-link is inert).
The residual is now provably reachable-to-`(0,0)` ONLY via a platform edit (bound `loadMembers` in the
snapshot path / a `frozen_tags` column) OR the disclosed-presenter-note (data proven-correct in the DB,
slow-only) — which needs the user's EXPLICIT sign-off.
**Type:** tik (tooling-iter shape — shipped the reusable deep-link diagnostic + used it for root-cause)
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: y — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D1 (resolve closed cycle id from DB + confirm FE reads `?cycle=` from URL), D2 (the cheap
dual-endpoint direct probe → falsification: frozen read itself is slow via `loadMembers`), D3 (the
USER-BLOCKER — deep-link falsified; the only remaining zero-edit path (disclosed-allow) needs sign-off).
**Side-deliverables (if any):** `probe-aireadiness-deeplink.spec.ts` — a reusable authenticated
dual-endpoint direct probe (lift-token + hit `/cycles` and the frozen data GET directly, with timing). It is
the general diagnostic for "is a cycle-scoped/frozen read actually fast end-to-end" — kept + committed.
**Routes carried forward:**
  - The 2 AI-readiness sections + the 3 workforce-aggregate sections → blocked on a USER DECISION (below).
    Handler once decided: `DISCLOSE-M51-iter09-aireadiness-frozen-slow` (disclosed-presenter-note, needs
    sign-off) OR `PLATFORM-M51-aireadiness-snapshot-loadmembers` (escalated platform edit: bound `loadMembers`
    in `buildResponseFromSnapshots` / add a `frozen_tags` column so the snapshot read needn't re-join live
    members).
**USER-BLOCKER (the decision the user must make — the deep-link is now falsified):** the user-chosen
zero-edit path (deep-link `?cycle=<closed>`) does NOT work: the frozen read is itself org-scale-slow
(`buildResponseFromSnapshots` → `loadMembers` full org hydration → 180 s+, never completes), so a `?cycle=`
deep-link would fire an inert query that still hangs. Every remaining path to `(0,0)`:
  (a) **DISCLOSED residual** (the run-5 brief's stated fallback): the data is PROVEN correct in the DB (199
      frozen snapshots, 78.4% stage-3, heroes right) — the 2 AI-readiness sections (+ the 3 workforce
      aggregates, same `loadMembers` family) are slow-but-correct due to a platform read-path perf wall, not a
      seed gap. Disclose as a presenter-note per the coverage-protocol's NARROW disclosed-allow → the gate
      reaches green-with-disclosure. Per the coverage-protocol + iter-07, this needs the user's EXPLICIT
      sign-off (it is NOT an editorial-citation auto-allow), which the build-iter cannot self-grant without
      loosening the gate to force a pass (the invariant's hard line). **The seeded closed-cycle data STAYS**
      (honest + correct; a presenter who waits out the load, or the platform once fixed, sees the real
      dashboard).
  (b) **ESCALATE a platform edit** (`unimplementable-without-platform-edit`, the milestone Re-scope trigger):
      bound `loadMembers` in the snapshot path (it only needs identity/tags for the ~199 snapshot users, not a
      whole-org unbounded hydration — pass the snapshot user-ids as `restrictUserIDs`), OR add the deferred
      `frozen_tags jsonb` column (M314b) so `buildResponseFromSnapshots` needn't re-join live members at all.
      The invariant forbids doing it here.
  (c) a NEW app read-path demo-patch that bounds `loadMembers` in the snapshot path (the
      `app-targetrole-authz-skip` precedent — a build-scratch source patch of the demo's ephemeral app clone) —
      a substantial new tooling investment NOT in the run-5 brief's chosen scope.
iter-08 surfaces the decision rather than picking one: (a) needs sign-off (can't be self-granted without
loosening the gate); (b)/(c) touch the platform / a new tooling investment. The closed-cycle seed + the
diagnostic probe are KEPT + committed (correct + reusable regardless of the chosen resolution). `fit-up-m51`
is NOT tagged (gate not met).
**Lessons:**
  - "Frozen"/"materialized"/"pre-computed" names a SCORES freeze, not necessarily a RESPONSE freeze. A snapshot
    read that re-joins live per-member identity (tags/name/role) re-incurs the org-scale member-load wall.
    MEASURE the fast branch END-TO-END with a direct authenticated probe of the EXACT request the strategy will
    fire, before betting a strategy on it. iter-07 confirmed only "the FE doesn't route to the frozen branch";
    iter-08 shows the frozen branch itself is slow — a strictly stronger falsification. Added to
    coverage-protocol.md (M51 iter-08 lesson).
  - The right cheap probe for "is this read fast" is a DIRECT authenticated `page.request` of the endpoint
    (bypass the FE React Query gate), not a full page navigation — the navigation conflates the FE's
    cycles-gate orchestration with the backend read cost. probe-aireadiness-deeplink.spec.ts is that probe.
