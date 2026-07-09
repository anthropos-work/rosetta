---
iter: 09
milestone: M211
iteration_type: tik
status: closed-no-lift
created: 2026-07-08
---

# iter-09 — tik: M42e employee coverage GREEN on the merged demo (fix the stale-offset UI-tier image)

**Active strategy reference:** TOK-01 "Warm-first cache-migrate, then cold-prove both stacks" — move (4)
"prove cold + M42 coverage + v2.0 Playthroughs". This tik targets gate sub-condition **(e)** (the M42 coverage
presence gate), employee vantage.

## Step 0 — Re-survey
The run-2 partial M42e sweep (killed mid-iter at the pause) left a report: `reachable=7 failingSections=8
personaFailures=2 escapes=0 notReached=0 → GATE NOT MET`. Re-surveyed the live demo-1 (still UP, 16 containers):
the target (employee coverage GREEN) is still untouched and still the right next thing under TOK-01.

## Cluster / target identified — root cause (diagnosed live)
The 8 failing sections + 2 persona failures are ALL one root cause: the demo's **next-web-app image bakes the
WRONG (non-offset) browser-facing URLs**. The running `demo-1-next-web` image has
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql` (+ BACKEND :8082, HOSTING :3000) — the
Dockerfile.dev **defaults** — instead of the demo-1 offset (:15050 / :18082 / :13000). Server-side/lightweight
queries use the internal docker URL (greeting, org, activity-stats, menu avatar → render fine), but **client-side
grid/profile queries fetch the dead :5050** (dev stack torn down) → perpetual skeleton loaders → empty library
grids + empty profile identity/skills → the 8 failing sections + the 2 persona failures (no role title, no
profile avatar). Verified: the router + backend + CMS all serve the data correctly (`publicJobSimulations`
returns 6, `libraryCategories` real, `public.skills`=42,790, `public.user_skills`=3,884); it is purely a
client-URL baking defect. Confirmed the studio-desk image is CORRECT (`VITE_GRAPHQL_ENDPOINT=...15050`) — only
next-web is stale.

**Why the tooling shipped a stale image:** `up-injected.sh`'s image-reuse guard (`if docker image inspect $img;
then reuse`) reuses a cached `demo-N-<frontend>` image **without validating its baked offset**. A pre-fix /
different-offset cached image is silently reused, baking wrong client URLs. This is the tooling bug (Fate-1,
fixable in rext — the platform Dockerfile is correct and untouched).

## Hypothesis
Harden the reuse guard to read the cached image's baked client endpoint (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` for
next-web, `VITE_GRAPHQL_ENDPOINT` for studio-desk) and **rebuild if it doesn't match this stack's offset**; then
rebuild demo-1's next-web with the correct offset URLs → the client grid/profile queries resolve → the 8 failing
sections + 2 persona failures clear → M42e employee GREEN.

## Expected lift
Gate sub-condition (e) employee half: `failingSections 8→0`, `personaFailures 2→0`, `escapes` stays 0,
`reachable` expands past 7 (populated grids give the crawler links to follow). Overall gate 5/6 → 5/6 with (e)
employee half proven (manager half → iter-10).

## Phase plan (per verification.md + coverage-protocol.md)
1. **Fix (rext, authoring copy):** `up-injected.sh` — add baked-offset validation to BOTH frontend reuse guards
   (rebuild-on-mismatch; reuse when matching or unverifiable → no regression). + unit tests in
   `test_frontend_build.py` (stale→rebuild, match→reuse) + the docker stub enhancement.
2. **Land:** run the frontend-build test suite; commit; tag `quick-change-m211`; re-pin the stack-demo
   consumption clone + `.agentspace/rext.tag`.
3. **Rebuild + prove:** rebuild `demo-1-next-web` via the fixed guard (lib-only source of the consumption
   clone's `up-injected.sh` → `build_frontend_next_web`, faithful pk + demopatches + offset args; reap-safe
   detached+poll) → recreate the next-web container → re-run the M42e employee coverage sweep → measure.

## Escalation conditions
- If, after the rebuild, coverage surfaces a NEW real content/seed gap → Fate-1 fix in seed/snapshot if tractable
  in-iter, else route forward.
- If a failure needs a platform-repo edit → escalate `unimplementable-without-platform-edit` (never edit the
  platform).

## Acceptable close-no-lift outcomes
If the rebuild proves the URL fix works but the sweep surfaces an unrelated new blocker that needs a separate
iter, close-fixed-partial with the URL fix landed + the new blocker routed forward.
