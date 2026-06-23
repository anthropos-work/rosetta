# M36 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §3b (the
dashboard surfaces + the aggregates needing distributions), §3c (the claimed-vs-verified aha), §6a #5.

## The dashboard spine (REST `/api/workforce/*` → `app/internal/workforce/*.go`)
_(verified-side: `user_skill_evidences`; mapped-side: `membership_skills` — mapped must outnumber verified per
skill for a believable funnel. Members need `job_role_id`/name + `joined_at` + tags.)_

## Aggregates needing believable distributions
_(self-eval accuracy = `user_level` vs `anthropos_level`; verification funnel; role-readiness [needs
`job_role_skills`]; AI-readiness [AI-named skills]; succession [validation + interview rows]; growth arc
[early-low/late-high]. The two employee heroes = the dashboard's standout high/low rows.)_

## Scope hard line
_(seed the spine for the seeded story; do NOT chase every widget — the milestone's biggest growth risk.)_

## Pre-flight audits — all sections (Phase 0b)
- **Verdict: YELLOW** (report: `kb-fidelity-audit.md`, 2026-06-23). No unmanaged blind area; the dashboard
  org-aggregate surfaces are a tracked `Delivers →` deliverable (extend `stories-spec.md` + `seeding-spec.md`).
  No stale load-bearing claim. Covers the whole M36 subsystem (`stack-seeding` seeders + corpus doc-half) — so
  per the §"Audit reuse" rule this verdict is reused across all M36 sections (same subsystem, knowledge docs
  unchanged between sections).
- **Topic → doc → code triples** (fast-start for future audits):
  - mapped funnel → `stories-spec.md` (NEW M36 §) → `seeders/membership_skills.go` (NEW)
  - tags/teams → `stories-spec.md` (NEW) → `seeders/tags.go` (NEW)
  - target-roles / mobility → `stories-spec.md` (NEW) → `seeders/target_roles.go` (NEW)
  - succession feeders → `stories-spec.md` 7-table chain → `seeders/persona*.go` + interview rows (NEW)
  - feedback ~2:1 → `stories-spec.md` (NEW) → `seeders/feedback.go` (NEW)
  - assignments fix → `seeding-spec.md` assignments surface → `seeders/assignments.go` (FIX)
  - skillpath `completed` share → `seeding-spec.md` → `seeders/skillpath_sessions.go` (verify)
  - org-scale distributions → `stories-spec.md` §3c → `seeders/persona*.go` + distribution logic
  - closure / isolation → `stories-spec.md` "Closure across all orgs" + `safety.md` §2.1 → `dna/seed_closure.go`, `isolation/`

## O4 — MIGRATED column/storage-key names (introspected on demo-3 `demo-3-postgresql-1`, 2026-06-23)
Authoritative source: live migrated tables (`\d`) on the demo-3 DB. Cross-checked against the platform ent
schemas in `stack-dev/{app,jobsimulation}/internal/.../ent/schema/`.

- **`public.membership_skills`** — `id` uuid, `created_at`/`updated_at`, `skill_id` varchar **NOT NULL**,
  `membership_skill_membership` uuid **NOT NULL** (FK→memberships), `membership_skill_organization` uuid
  **NOT NULL** (FK→orgs), `skill_name` varchar (nullable, but the funnel query filters `skill_name IS NOT NULL`
  → MUST set), `specialization_id`/`specialization_name`/`category_id`/`category_name` varchar nullable,
  `skill_level` bigint **NOT NULL**, `match_type` varchar NOT NULL default `'match'`. UNIQUE
  `(skill_id, membership_skill_membership, membership_skill_organization)`.
- **`public.tags`** — `name` varchar NOT NULL, `organization` uuid NOT NULL (FK→orgs). UNIQUE `(name, organization)`.
- **`public.membership_tags`** — `membership_tag_tag` uuid (FK→tags), `membership_tag_membership` uuid
  (FK→memberships), `membership_tag_organization` uuid (FK→orgs); all NOT NULL. No natural UNIQUE beyond `id`
  → dedup on a deterministic `id`.
- **`public.organization_target_roles`** — `target_job_role_id` varchar NOT NULL (a `J-XXXX` node-id),
  `organization_id` uuid NOT NULL, `assignee_id` uuid NOT NULL (FK→memberships), `assigner_id` uuid NOT NULL
  (FK→memberships), `completed_at`/`deleted_at` nullable. UNIQUE
  `(target_job_role_id, assignee_id, organization_id) WHERE deleted_at IS NULL`.
- **`public.user_target_roles`** — `target_job_role_id` varchar NOT NULL, `user_id` uuid nullable (FK→users —
  NOTE keyed off **user**, not membership), `completed_at`/`deleted_at` nullable. UNIQUE
  `(target_job_role_id, user_id) WHERE deleted_at IS NULL`.
- **`public.organization_assignment_sessions`** — `assignment_id` uuid NOT NULL (FK→organization_assignments),
  `organization_id` uuid NOT NULL, `session_id` uuid (FK→local_skill_path_sessions) OR `js_session_id` uuid
  (FK→local_jobsimulation_sessions); CHECK `(session_id IS NOT NULL OR js_session_id IS NOT NULL)`; `progress`
  bigint NOT NULL default 0, `started_at`/`ended_at` nullable. UNIQUE `(assignment_id, js_session_id)` +
  `(assignment_id, session_id)`.
- **`public.job_simulation_feedbacks`** — `simulation_id` uuid NOT NULL, `session_id` uuid NOT NULL,
  `is_positive` bool NOT NULL default false, `option` varchar NOT NULL default `'FUN_AND_USEFUL'` (enum, see
  below), `feedback` varchar nullable, `organization_id` uuid nullable, `user_id` uuid NOT NULL (FK→users).
  UNIQUE `(user_id, session_id)`.
- **`jobsimulation.interview_extraction_results`** — `user_report` jsonb **NOT NULL**, `manager_report` jsonb
  **NOT NULL**, `session_id` uuid NOT NULL (FK→jobsimulation.sessions, UNIQUE), `summary` jsonb nullable. The
  succession query reads `summary->>{sentiment,wellbeing_score,readiness_level,self_awareness,
  signal_count_positive,signal_count_attention}` and requires `summary IS NOT NULL`.
- **`jobsimulation.validation_attempt_skill_results`** (succession feeder — already seeded by the M34 chain;
  M36 sizes it to clear the gate): `skill` varchar NOT NULL, `score`/`competency_level_score` real,
  `validation_attempt_result_id` uuid NOT NULL, `status` varchar NOT NULL default `'pending'`.

### `job_simulation_feedbacks.option` enum (from `app/internal/data/ent/enum/jobsimulation_feedback.go`)
Positive: `FUN_AND_USEFUL`, `HELPFUL_FEEDBACK`, `ENJOYED_CHALLENGE`, `REALISTIC_SCENARIO`, `ENOUGH_TIME`.
Negative: `NOT_ENOUGH_TIME`, `CONFUSING_SCENARIO`, `AI_CHARACTER_ISSUES`, `NOT_USEFUL`.

## Dashboard spine queries (the shape each surface must satisfy — `app/internal/workforce/*.go`)
- **Verification funnel** (`skills.go::aggregateSkills`): mapped = `COUNT(DISTINCT membership_skill_membership)`
  on `membership_skills` GROUP BY `skill_name` (joins live memberships); verified = `COUNT(DISTINCT membership)`
  on `user_skill_evidences WHERE is_verified` joined to `skiller.skills` by `sk.name`. **Funnel matches on
  skill NAME** → `membership_skills.skill_name` must equal the verified skills' `skiller.skills.name`, and
  per skill MAPPED members must OUTNUMBER verified members.
- **Self-eval gap** (`skills_verification.go::GetSelfEvalAccuracy`): reads `user_skill_evidences.user_level`
  vs `anthropos_level` (both NOT NULL) — ALREADY seeded by the M34/M35 PersonaSeeder (D4). M36 surfaces it
  at org scale (the over/under-claimer mix the trajectories already encode).
- **Succession coverage gate** (`succession.go::buildCoverage`): `too_sparse` if `total < 15`; `basic` if
  `<50%` have ≥3 skills; `full` if `>20%` have an `interview_extraction_results` row with a non-null `summary`;
  else `good`. Each story org is sized ≥120 → not sparse; the chain gives heroes skills+sims; M36 adds the
  interview rows (>20% of members) to reach `full`.
- **Role readiness** needs `job_role_skills` (replayed) + members' verified skills (chain) → already present;
  M36 adds `*_target_roles` for the gap+mobility surface.
