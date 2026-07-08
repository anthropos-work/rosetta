---
iter: 10
milestone: M211
iteration_type: tik
status: closed-fixed-partial
created: 2026-07-08
---

# iter-10 — tik: the TRULY-cold /demo-up (purge the stale mishmash) + M42 coverage GREEN

**Active strategy reference:** TOK-01 move (4) "prove COLD + M42 coverage". This tik executes the honest
cold-demo proof that iter-09's diagnosis demands.

## Step 0 — Re-survey
iter-09 proved demo-1 was a **stale-image mishmash** spanning the merge (pre-merge router with a dead `skiller`
subgraph; stale `app:injected` rejecting `Skill.name`) — iter-08's "cold /demo-up GREEN" was image-WARM. The
clones + composition are correct (app v1.334.1 + `backend.graphqls` both define `Skill.name`; 4 subgraphs, no
skiller). So a **truly-cold rebuild** (purge the built images + rebuild all from the merged clones) is the
honest gate proof and will produce a consistent merged platform.

## Cluster / target identified
Sub-condition (e) M42 coverage — blocked by the stale image set, NOT by content/seed (data verified present:
`public.skills`=42,790, `user_skills`=3,884). A clean rebuild aligns router (4-subgraph) ↔ backend (Skill.name)
→ `searchSimulations`/`publicJobSimulations.skills` resolve → library grids + profile populate.

## Hypothesis
`rosetta-demo down 1 --purge` (remove demo-1's built images) → truly-cold `/demo-up` (rebuild ALL from the
merged clones: router recomposes 4 subgraphs, `app:injected` gets `Skill.name`, frontends offset-correct) →
the merged platform stands up consistent → M42e employee coverage GREEN (the 8 failing sections + 2 persona
failures clear).

## Expected lift
Sub-condition (e) employee half: `failingSections 8→0`, `personaFailures 2→0`, `reachable` expands past 7.
This ALSO re-proves (a)-(d),(f) COLD on a truly-clean image set (correcting iter-08's image-warm claim).

## Phase plan
1. **Purge:** `rosetta-demo down 1 --purge` (scoped to demo-1 images; dev + other demos untouched).
2. **Truly-cold `/demo-up`:** `up-injected.sh 1` (consumption clone @ quick-change-m211) — reap-safe
   DETACHED + poll. Rebuilds every backend Go service + router + frontends + clerkenstein from the merged
   clones. Watch for the auto-verify tail + the sub-condition markers.
3. **Re-measure:** re-run the M42e employee coverage sweep vs the clean demo-1 → target GREEN. Spot-check the
   `searchSimulations` federation (no skiller error, backend resolves `Skill.name`).

## Escalation conditions
- If, AFTER a truly-clean rebuild, `searchSimulations`/`Skill.name` STILL fails → that would be a genuine
  platform version-skew (graphql-wundergraph composition vs app clone) → escalate
  `unimplementable-without-platform-edit` (never edit the platform).
- If a docker reap kills the cold build → resume (detached log + `.done` marker).

## Acceptable close-no-lift outcomes
If the cold build proves the platform stands up but coverage surfaces an unrelated new content/seed gap that's
a separate iter's work, close-fixed-partial with the cold-rebuild proven + the new gap routed.
