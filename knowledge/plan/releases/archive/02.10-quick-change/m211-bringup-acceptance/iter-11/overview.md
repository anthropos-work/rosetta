---
iter: 11
milestone: M211
iteration_type: tik
status: closed-fixed-partial
created: 2026-07-08
---

# iter-11 — tik: re-prove under the FIXED tooling → M42e employee coverage GREEN

**Active strategy reference:** TOK-01 move (4). iter-10 fixed the root cause (stale build-scratch); this tik
proves it end-to-end.

## Step 0 — Re-survey
iter-10 landed the build-scratch re-sync fix (rext `0593cff`, tag `quick-change-m211`) + verified it produces a
fresh scratch (app v1.315→v1.334, `Skill.name` 0→5). The consumption clone is re-pinned. demo-1 is still up on
STALE binaries. This tik re-runs the demo-up under the fixed tooling so the injected services rebuild fresh.

## Cluster / target identified
Sub-condition (e) M42e employee coverage — the fix must be exercised end-to-end: re-sync all injected scratches
→ rebuild app/cms/jobsim/skillpath from v1.334.1 → the merged Skill federation entity is present → the router's
`_entities(Skill.name)` resolves → library grids + profile skills populate.

## Hypothesis
`rosetta-demo down 1` (keep images for cache) → `up-injected.sh 1` (fixed tooling re-syncs all scratches +
rebuilds the injected services) → `searchSimulations`/`publicJobSimulations.skills` resolve (no 422) → M42e
employee coverage GREEN (`failingSections 8→0`, `personaFailures 2→0`, `reachable` expands past 7).

## Expected lift
Sub-condition (e) employee half → MET. Also re-confirms the demo stands up green on a consistent, fully-merged
image set.

## Phase plan
1. `rosetta-demo down 1` (keep images — docker cache speeds the unchanged layers).
2. `up-injected.sh 1` under the fixed tooling — reap-safe DETACHED + poll. Only the re-synced-source Go layers
   (app/cms/jobsim/skillpath) rebuild; router + frontends are cached.
3. Spot-check the federation (`publicJobSimulations.skills` resolves, no skiller/backend 422).
4. Re-run the M42e employee coverage sweep → target GREEN.

## Escalation conditions
- If federation STILL 422s after a rebuild from the v1.334.1 scratch → escalate (genuine platform issue).
- Docker reap → resume from the detached log.

## Acceptable close-no-lift outcomes
If coverage surfaces a NEW unrelated content/seed gap after the federation is confirmed fixed, close-fixed-partial
with the federation proven + the new gap routed.
