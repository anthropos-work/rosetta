---
iter: 13
milestone: M211
iteration_type: tik
status: closed-fixed
created: 2026-07-08
---

# iter-13 — tik: teach the cross-port hook the ant-academy destination (crossPortFollowFails 1→0)

**Active strategy reference:** TOK-01 move (4). Clears the crossPortFollowFail iter-12 exposed; assesses the
last employee residual.

## Step 0 — Re-survey
iter-12 fixed escapes (40→0) but the demo-local academy link made the crawler FOLLOW :13077 and the
studio-desk-only cross-port hook false-FAILED it (crossPortFollowFails=1). Employee coverage: failingSections=1
(sim-embeddings) + crossPortFollowFails=1.

## Hypothesis
Branch `onCrossPortFollow` by destination port (studio-desk :9000+offset vs ant-academy :3077+offset), asserting
each app's own home markers → the academy follow passes → crossPortFollowFails 1→0.

## Outcome
Landed the hook branch + `ANT_ACADEMY_HOME_SECTION` (rext commit, tag `quick-change-m211` moved; compiles + 42
unit specs pass). Re-ran M42e: `cross-port follow ok … ant-academy home OK (localhost:13077, HTTP 200, marker
present)` → **crossPortFollowFails 1→0**.

## Re-measurement (M42e employee coverage — final this session)
| metric | iter-12 | iter-13 |
|---|---|---|
| reachable | 62 | 62 |
| failingSections | 1 | **1** (sim-embeddings — the SOLE residual) |
| personaFailures | 0 | 0 |
| escapes | 0 | 0 |
| crossPortFollowFails | 1 | **0** ✅ |
| GATE | NOT MET | NOT MET (1 section short) |

**Employee coverage is now ONE section short** — everything green except `/library/ai-simulations`
`sim-card-grid` (empty). Root: `public.simulation_embeddings` is ABSENT (only the taxonomy `skill_embeddings`
+ `job_role_embeddings` loaded); the AI-sims grid's `searchSimulations` needs simulation embeddings for its
vector search. This is a DATA/snapshot gap (no cached sim-embeddings snapshot + no local capture source — the
iter-08 rc=5 cache-miss), not a code/federation bug.

## Routes carried forward (Fate-3 → next session)
- **`TOOLING-M211-sim-embeddings`:** fill the AI-sims grid. Three candidate surfaces: (a) a sim-embeddings
  snapshot cache-fill (needs a capture source — the cold-start class, may be unfillable locally); (b) run the
  platform's local embedding generation over the seeded sims; (c) a frontend demopatch so the empty-query
  AI-sims default uses `publicJobSimulations` (which resolves) instead of `searchSimulations`. NOT a platform
  bug — fillable via tooling/data, just non-trivial.
- **M42m manager coverage** (+ re-pin the drifted `next-web-studio-url`/`next-web-public-website-url`
  demopatch manifest hashes to the re-synced v2.1 source), **v2.0 Playthroughs**, **cold `/dev-up`**.
