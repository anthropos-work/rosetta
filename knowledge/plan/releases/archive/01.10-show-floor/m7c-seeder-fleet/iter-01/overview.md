---
iter: 01
type: tok-bootstrap
milestone: M7c
date: 2026-06-04
---

# iter-01 — bootstrap tok: survey + strategy (TOK-01)

The opening strategy iter for the seeder fleet. Surveys the live `demo-1` schema to ground the fleet, corrects
M7b's guessed catalog table names, classifies each planned surface by **reachability**, and authors the
sequencing strategy (TOK-01, recorded in the milestone-root `decisions.md`).

## Survey findings (live demo-1 introspection)

**Corrected catalog table names** (M7b's were guesses — the concrete iter-01 deliverable, applied to
`data-dna.json`):
- `skillpath-sessions`: `skillpath.sessions` → **`skillpath.skill_path_sessions`**
- `assignments`: `public.assignments` → **`public.organization_assignments`**
- `activity`: `jobsimulation.activity` → **`jobsimulation.activity_events`**
- `taxonomy` (`skiller.skills`) + `jobsim-sessions` (`jobsimulation.sessions`) were already correct.

**Reachability classification** (the key strategic finding):

| Surface | Reachable now? | Why |
|---|---|---|
| `jobsim-sessions` (`jobsimulation.sessions`) | ✅ **yes** | **no FKs**; `sim_id`/`sim_type`/`owner_id`/`token` are NOT-NULL but `sim_id` is a free value (references a Directus sim, *not* FK-enforced) |
| `skillpath-sessions` (`skillpath.skill_path_sessions`) | ✅ **yes** | **no FKs**; `skill_path_id` is a free value (not FK-enforced) |
| `assignments` (`public.organization_assignments`) | ✅ **yes** | FKs only to `memberships` + `organizations` (both seeded); `resource_id` is a free value |
| `activity` (`jobsimulation.activity_events`) | ✅ **yes** (after sessions) | FK `session_id`→`sessions` (NOT NULL) — depends on jobsim-sessions seeding first |
| `taxonomy` (`skiller.skills`/`job_roles`) | ⚠️ **blocked** | needs the node-id hierarchy + the **pre-embedded skiller snapshot** (empty in demo-1); the "consume a snapshot" surface |
| `content` (Directus) | ⛔ **blocked** | the **shared** Directus instance — snapshot-replay only (the isolation guard blocks live writes) |

**The unlock:** the *believability core* — time-distributed user **activity** (sessions, assignments, activity
events) — is reachable **without** solving Directus, because the content references (`sim_id`,
`skill_path_id`, `resource_id`) are free values, not DB-enforced FKs. That's the bulk of what makes a demo feel
alive.

## Outcome
Strategy authored (TOK-01). Catalog corrected + validated. Next tik: implement `jobsim-sessions` (the
foundational activity surface; `activity` depends on it) with the **backdated, time-distributed** generator,
promote it planned→seeded, and measure the data-DNA coverage rise.
