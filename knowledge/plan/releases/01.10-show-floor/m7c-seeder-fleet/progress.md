# M7c — progress (iterative running ledger)

**Milestone:** M7c — The seeder fleet, to a coverage gate · **Shape:** iterative · **Status:** planned

**Exit gate:** a `stack.seed.yaml` (1k-user org, N months activity) → demo identity logs in **200** · data-DNA
**coverage ≥ 90% / critical 100%** · seeding **< 2 min** · seeding audit **zero shared/prod writes**.

**Iteration protocol:** `corpus/architecture/alignment_testing.md` (measure → implement/deepen a seeder →
re-measure data-DNA coverage), via M7b's data dimension. Built by `/developer-kit:build-mstone-iters`.

## Running ledger

| Iter | Type | Surface(s) | Coverage | Notes |
|------|------|-----------|----------|-------|
| 01 | tok-bootstrap | — | 40% (4/10, baseline) | TOK-01 strategy; reachability survey; corrected 3 guessed catalog table names. Believability core (activity surfaces) reachable w/o Directus. |
| 02 | tik | jobsim-sessions | **50% (5/10)** | The believability core: `JobsimSessionsSeeder` — 1980 backdated, pass/fail-weighted sessions (deterministic, time-distributed over the activity span). Promoted planned→seeded; `measure` [PASS] 100%/Critical 100% live on demo-1; isolation clean. Caught + fixed a harness bug: `introspect` couldn't load a freshly-promoted surface (Validate-before-populate). seeders 14 tests. |
| 03 | tik | skillpath-sessions | **60% (6/10)** | `SkillpathSessionsSeeder` — 1024 backdated skill-path sessions (0–2/user). Caught a real UNIQUE constraint `(user_id, skill_path_id, version)` → fixed (skill_path_id indexed by session number so a user's paths are distinct). [PASS]. |
| 04 | tik | assignments | **70% (7/10)** | `AssignmentsSeeder` — 505 org assignments (admin→members, ~½ of members), 3 FKs (assignee/assigner→memberships, org→organizations) all referentially valid (fk-valid [PASS]). |
| 05 | tik | activity | **80% (8/10)** | `ActivitySeeder` — 4000 backdated activity events (4/user), FK session_id→sessions valid (anchored to each user's session :0). **Full believability core complete.** seeders **20 tests**. Full 8-seeder seed (~8500 rows) runs in **0.69s** (< 2 min gate). measure Overall/Critical 100%. |

## Gate scorecard (after iter-05)
| Gate condition | Status |
|---|---|
| (a) demo identity logs in → **200** | ✅ HTTP 200 (`membershipsCount: 1001`) live on demo-1 |
| (b) data-DNA coverage **≥ 90% / critical 100%** | ⚠️ **80% (8/10)**, critical 100% — the 90% needs taxonomy or content (both snapshot-blocked) → **re-scope/waiver decision** |
| (c) seeding **< 2 min** | ✅ **0.69s** (~8500 rows) |
| (d) isolation audit **zero shared/prod writes** | ✅ clean |

## Routes carried forward / re-scope decision (iter-05)
- **taxonomy** (`skiller.skills`/`job_roles`) + **content** (shared Directus) are the 2 remaining surfaces, both
  **snapshot-blocked**: taxonomy needs the pre-embedded skiller node-hierarchy snapshot (empty in demo-1); content
  is the shared Directus instance (the isolation guard blocks live writes — snapshot-replay only). Both are the
  **hard line** (M7c seeds structural data; it does not author the 60K-skill taxonomy nor write shared Directus).
- **Re-scope decision (awaiting user):** the 90% gate is unreachable without these two → either **(A) waive
  them** per the Re-scope trigger and close at 80% coverage / 100%-over-reachable (the believable-demo subset is
  complete), or **(B) keep M7c open** to solve the snapshot (taxonomy import + Directus snapshot-replay — heavy,
  arguably v1.2 / out of M7c's structural-data scope).
