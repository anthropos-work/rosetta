# M34 — Spec notes

Technical notes accumulate here during build. The authoritative design is
[`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) (§3 the chain, §6 the build plan).
**Port the reference `seed.sql`** (`kb-ant-business/.claude/skills/seed-verified-skill/scripts/seed.sql`) —
don't reinvent the table shapes.

## The 7-table chain (owning schema)
- `jobsimulation`: `sessions` → `validation_attempt_results` → `validation_attempt_skill_results` →
  `validation_criterion_results`.
- `public` (app): `local_jobsimulation_sessions`, `user_skills`, `user_skill_evidences`.
- `skiller.skills` supplies the `node_id` (loose string ref, not a FK).

## Constraint landmines (spec §3 — DB-enforced vs logical)
- **DB-enforced (reject):** `user_skills_check_foreign_keys` CHECK; `idx_unique_job_simulation` partial UNIQUE;
  `user_skill_evidences` UNIQUE(skill_id,user_id); `token` varchar(10) UNIQUE.
- **Logical (insert-but-invisible):** enum/result/sim_type are free-text — write `status='ended'`,
  `completion_status∈{passed,failed}`, `result_status='completed'`, `sim_type='SIMULATION_TYPE_*'`,
  `evaluation_status='passed'`, `competency_level_score>0`; chart needs ≥2 datapoints; levels 0–100 ÷5;
  misspelled `local_jobsimulation_sessions.completition_status`; `job_simulation_id`=sim/Directus UUID.

## O4 — live storage-key names
_(record the `\d` output for `user_skills` / `user_skill_evidences` / `local_jobsimulation_sessions` here
once captured via `/db-query`.)_
