# M41 Spec Notes

## Pre-flight audits — G3 work history (first section)
KB-fidelity audit (`/developer-kit:audit-kb-fidelity --milestone=M41`): **GREEN**. Report:
`kb-fidelity-audit.md`. Both delivery docs (`seeding-spec.md`, `demo/stories-spec.md`) PAIRED + aligned with the
rext seeder code @ `method-acting-m40`; all cited line anchors resolve; M41's new tables are explicit
`Delivers →` deliverables (not a blind area). Reused for all M41 sections (same `stack-seeding` subsystem,
unchanged knowledge docs).


Technical notes accumulate here during build. The authoritative design is
[`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review 2026-06-24;
root-cause workflow w7t4wq2z4). M41 extends the verified-skill chain documented in
[`corpus/ops/demo/stories-spec.md`](../../../../corpus/ops/demo/stories-spec.md) — do not reinvent the
`user_skills` / `user_skill_evidences` shapes.

## G3 — `ProfileSeeder` (rext `stack-seeding`)
Net-new seeder (`profile.go`, surface `"profiles"`, DependsOn `["users","taxonomy","content"]`,
`PerStackIsolated`). Iterates `s.EffectiveStories()` like PersonaSeeder; writes per END-USER hero (managers
skipped — no personal timeline). Reuses `resolveJobRoleRefs` (role node-id + title) + `resolveNamedSkillRefs`
(skill names for the json arrays + the unverified tail's node-ids). Backdating anchored on `st.Activity.Months`.

## LIVE-SCHEMA CORRECTIONS (verified on demo-3 — the overview/spec column-lists were guesses; these win)
- **`user_experiences.company` is `uuid NOT NULL`** with FK → `companies(id)` (NOT nullable as the overview
  claimed). ⇒ the seeder MUST seed a `public.companies` row per distinct (name,domain) and reference it. The
  GraphQL `Company` resolver does `obj.QueryCompany().Only(ctx)` (hard-required), so a NULL/dangling company
  would error the whole timeline. `companies` has UNIQUE `(name, domain)` + NOT-NULL `name`,`domain`.
- **`from`/`to` are `date` (NOT timestamptz).** CHECK `user_experience_check_from_gte_to` = `from <= to OR to
  IS NULL`; CHECK both `>= 1900-01-01`. ⇒ emit `time.Time` truncated to date; a current role leaves `to` NULL.
  Same two CHECKs on `user_educations`.
- **`location_type` stored strings are lowercase `hybrid|fullremote|inoffice`** (ent enum
  `internal/data/ent/enum/location_type.go`) — NOT `INOFFICE/HYBRID`. The GraphQL `LocationType` enum =
  `hybrid|fullremote|inoffice`. A wrong-case value is inserted-but-unmappable (the M34 free-text-enum class).
- **`user_experiences.skills` / `user_educations.skills` are `json` default `'{}'`** — the platform writes a
  json ARRAY of skill names; seed `["name1","name2"]` (the per-experience skill labels). job_role_id is the
  `J-…` node-id form (same as memberships).
- Timeline read is USER-scoped (`TimelineGrouped(userID)` → ent privacy `UserMixin`) — no org column on either
  table; seeding on the hero's `users.id` is sufficient.

## `public.user_experiences` (work history) — final column list
`[id, created_at, "user"(uuid FK→users), company(uuid FK→companies, NOT NULL), title(text), description(varchar),
from(date), to(date,null=current), skills(json array of names), location(varchar), location_type(varchar
lowercase enum), job_role_id(J-… node-id)]`. 2–3 rows/hero, a backdated progression (older role → current role),
deterministic UUIDs; current role's `to` NULL.

## `public.user_educations` (education) — final column list
`[id, created_at, "user", title(text NOT NULL), from(date), to(date), skills(json), institute(varchar NOT NULL),
field_of_study(varchar)]`. 1–2 rows/hero, backdated before the work history; deterministic UUIDs.

## `/profile` timeline read path (read-only — no platform edit)
TODO: confirm the timeline reads `ent.UserExperience` / `ent.UserEducation` via `GET_TIMELINE`; M41 supplies
rows only — no next-web / app / cms / jobsimulation edits.

## G5 — verified-depth bump
TODO: bump the preset `verified` count (e.g. `8 → ~30`) in `stories.seed.yaml:48` + `stories-maya.seed.yaml:36`;
trace `blueprint.go:227 EffectiveVerified → persona.go:132-178 resolveHeroSkills` × `verifiedSessionsPerSkill=3`
(`persona.go:46`) ⇒ ~30 distinct (10 role-coherent + auto-top-up from the flat public pool past the role's 10)
× 3 = ~90 `user_skills` + ~30 evidences.

## G5 — claimed-but-unverified tail
Seed ~60 `user_skills` / `user_skill_evidences` per thriving hero with `is_verified=false`, no
`job_simulation_id`, `user_level` set, `anthropos_level` NULL ⇒ "overall ≈ 90 = ~30 verified + ~60
claimed-unverified"; widens the visible claimed-vs-verified gap. Keep the `user_level` (claimed) vs
`anthropos_level` (verified) mechanic intact.

LANDMINE (DB-enforced): `user_skills_check_foreign_keys` requires ≥1 provenance edge non-NULL. The tail has NO
`job_simulation_id` ⇒ it MUST set another edge. DECIDED: tie the tail to the seeded G3 work-history via
`user_skill_experience`/`user_skill_education` (those FK columns + their partial UNIQUEs exist). This makes the
claimed skills EVIDENCED BY the work history — the natural G3↔G5 join, and what the workExperience.Skills
resolver (`GetUserExperienceSkills` → `userskill.HasExperienceWith`) reads, so the claimed skills render UNDER
each experience too. Partial UNIQUE `idx_unique_user_skill_experience (skill_id, user_skill_user,
user_skill_experience)` ⇒ distinct (skill × experience) per row.

`user_skill_evidences` tail UPSERT differs from the verified one: `is_verified=false`, `anthropos_level` NULL
(stays the gap-renderer), `user_level` set, `jobsimulation_session_id` NULL, `job_simulation_count=0`. Needs a
SEPARATE upsert SQL from the verified `upsertEvidenceSQL` (which hardcodes `is_verified=true`). The "overall"
skill count = distinct `user_skill_evidences` rows per user (verified ∪ unverified, distinct skill_id), so the
~30 verified + ~60 unverified must be DISTINCT skill node-ids (the tail draws from the flat pool past the
hero's verified set).

## Leave-as-is: the `mapped:22` wiring
TODO: confirm `mapped:22` feeds `membership_skills` (the org funnel), NOT `user_skills` — do not touch it;
profile depth comes from `user_skills`.
