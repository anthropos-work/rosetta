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

## O4 — live storage-key names (RESOLVED via postgres MCP structural read, 2026-06-23)
Confirmed against the prod schema (M21-sanctioned structural-only read; the per-stack stack has the same
migrated schema). NOT-NULL-no-default columns are the load-bearing ones a raw COPY must supply.

- **`public.user_skills`** — NOT NULL: `level`(bigint), `skill_id`(varchar), `user_skill_user`(uuid — the
  user FK storage key), `acquired_at`. Defaults: `is_verified`=false, `match_type`='match'. CHECK
  `user_skills_check_foreign_keys` = ≥1 of {experience/education/cert/project/content/volunteering edge,
  skill_path_id, **job_simulation_id**} non-NULL → set `job_simulation_id` (the sim/Directus template UUID).
  Partial UNIQUE `idx_unique_job_simulation (skill_id, job_simulation_id, user_skill_user) WHERE
  job_simulation_id IS NOT NULL` → a **distinct sim_id per (skill,user) verified row**.
- **`public.user_skill_evidences`** — NOT NULL: `skill_id`(varchar), `acquired_at`, `user_id`, all `*_count`
  (bigint, default 0), `is_verified`(default false). Nullable LEVEL columns (all bigint, 0–100 by UI
  convention, NO DB CHECK): `level`, `anthropos_level`, **`user_level`** (the seed.sql OMISSION — the
  claimed side of the claimed-vs-verified widget), `manager_level`, `peers_level`. UNIQUE
  `idx_unique_user_skill_evidence (skill_id, user_id)` → **UPSERT** (INSERT … ON CONFLICT (skill_id,user_id)
  DO UPDATE), NOT a blind COPY.
- **`public.local_jobsimulation_sessions`** — NOT NULL: `jobsimulation_id`(=sim_id), `jobsimulation_session_id`
  (=session id), `status`, **`completition_status`** (sic — the misspelled column), `session_created_at`,
  `session_updated_at`, `user_id`. `score` default 0, `validation_version` default 1.
- **`jobsimulation.validation_attempt_results`** — NOT NULL: `session_id`, `timestamp_reference`,
  `acceptance_status`(='passed'), `evaluation_status`(='passed'), `success_threshold`(real),
  `explanation_summary`, `personal_explanation_summary`. `score`,`quick_summary` nullable.
- **`jobsimulation.validation_attempt_skill_results`** — NOT NULL: `skill`(varchar NodeID),
  `validation_attempt_result_id`, `status`(='completed'), `is_qualitative`(=false). `score`,
  **`competency_level_score`** (>0 — the chart datapoint), feedback nullable.
- **`jobsimulation.validation_criterion_results`** — NOT NULL: `criterion_id`, `type`(='evaluation'),
  `title`, `skills`(jsonb [NodeID]), `success_threshold`, `input_format`(='chat'),
  `validation_attempt_result_id`, `criterion_index`(bigint), `status`(='completed').
  `validation_attempt_skill_result_id` nullable (set it to link).

## O6 — usable public roles (RESOLVED)
Maya's spec role "Backend Engineer" is NOT a public job_role; "Backend Developer", "Software Engineer",
"Engineering Manager" ARE (10 core public skills each). → the persona's role string must resolve via
`skillsByRole` OR fall back to the flat pool (the empty-pool fallback is load-bearing). M34 uses a role that
resolves so Maya's verified skills are role-coherent (D3).

## Pre-flight audits — G14 fix (first section)

## Pre-flight audits — G14 fix (first section)
- **Phase 0b KB-fidelity:** YELLOW (proceed). Report: `kb-fidelity-audit.md`. SHA at audit: `02af617`.
  - PAIRED + ALIGNED: seeding-spec.md (M7a/M7c scope boundary), snapshot-spec.md (M23 closure gene — the
    pattern the new seed-side closure gene mirrors), safety.md (PerStackIsolated classes — new surfaces
    inherit it).
  - Blind area (verified-skill-chain reference doc) is a tracked milestone deliverable (overview `Delivers`
    → NEW `corpus/ops/demo/stories-spec.md`), so it does NOT block Phase 1.
  - data-DNA.json does not yet enumerate `user_skills`/`user_skill_evidences`/`validation_attempt_*` — these
    are PersonaSeeder's net-new surfaces; the closure gene is a new data-DNA gene (expected new work).

## Topic → doc → code triples (fast-start for future audits)
- seeding framework + fleet → `corpus/ops/seeding-spec.md` → `stack-seeding/{seeders,seeder,blueprint}/`
- G14 session bug → `seeding-spec.md` + NEW `stories-spec.md` → `seeders/jobsim_sessions.go`
- taxonomy/content-ref resolver → `seeding-spec.md` + `snapshot-spec.md` → `seeders/contentref.go`
- cross-surface closure gene → `snapshot-spec.md` §M23 → `dna/snapshot.go`, `dna/fidelity_probe.go`
- isolation/safety → `corpus/ops/safety.md` → `isolation/`

## Reference impl shapes (PORTED from seed.sql — verified against live schema)
The 7-table chain, in seed order, with the columns the reference `seed.sql` writes (+ the two it OMITS):
1. `jobsimulation.sessions` — id, owner_id, sim_id, sim_type=`SIMULATION_TYPE_ASSESSMENT|_HIRING`,
   status=`ended`, completion_status=`passed`, score, validation_version, language, **result_status=`completed`**,
   chime_status=`not_available`, token=`substr(md5(id),1,7)`, started_at, ended_at, interactions_progress=100.
2. `public.local_jobsimulation_sessions` — jobsimulation_id (=sim_id), jobsimulation_session_id (=session id),
   status=`ended`, **completition_status** (sic)=`passed`, session_*_at, score, validation_version, user_id.
3. `jobsimulation.validation_attempt_results` — id, session_id, timestamp_reference, acceptance_status=`passed`,
   evaluation_status=`passed`, success_threshold=60, score, quick/explanation/personal summaries.
4. `jobsimulation.validation_attempt_skill_results` — id, validation_attempt_result_id, skill=NodeID,
   is_qualitative=false, status=`completed`, score, **competency_level_score>0**, feedback fields.
5. `jobsimulation.validation_criterion_results` — 3/session; id, criterion_id, criterion_index, status=`completed`,
   type=`evaluation`, title, skills=jsonb[NodeID], success_threshold, input_format=`chat`, score, feedback,
   validation_attempt_result_id, validation_attempt_skill_result_id.
6. `public.user_skills` — id, user_skill_user (=user uuid), skill_id=NodeID, level (1–5 convention),
   is_verified=true, **job_simulation_id=sim_id** (the CHECK + partial-UNIQUE column — distinct sim per row),
   acquired_at, match_type=`match`.
7. `public.user_skill_evidences` — UPSERT on (skill_id, user_id); skill_id, acquired_at, years, *_count,
   level (0–100), anthropos_level (0–100), **user_level (0–100 — the seed.sql OMISSION; per self_eval_bias)**,
   is_verified=true, verification_date, user_id, jobsimulation_session_id (=session uuid, NOT sim_id).
