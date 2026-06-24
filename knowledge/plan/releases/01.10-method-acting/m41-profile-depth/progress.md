# M41 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist

- [x] **G3 — `ProfileSeeder` (work history)** — new seeder (`profile.go`/`profile_write.go`, surface
  `"profiles"`) writing 3 `public.user_experiences` per END-USER hero (+ a `companies` row per employer): columns
  `[id, created_at, "user", company, title, description, from, to, skills(json), location, location_type,
  job_role_id]`; deterministic UUIDs; backdated within the Activity span; role-aligned titles (resolved
  `jobRoleRefs`); current role `to`=NULL. **Live-schema corrections applied** (company NOT NULL FK, DATE from/to,
  lowercase `location_type` enum, json skills). Managers skipped.
- [x] **G3 — `ProfileSeeder` (education)** — same seeder writing 1 `public.user_educations` per hero: columns
  `[id, created_at, "user", title, from, to, skills(json), institute, field_of_study]`; deterministic UUIDs;
  graduation backdated at/before the earliest job.
- [x] **G5 — bump verified depth** — `verified: 8 → ~30` for the thriving heroes (Maya 30, Sara 28) in
  `stories.seed.yaml` + `stories-maya.seed.yaml`, flowing through `EffectiveVerified → resolveHeroSkills ×
  verifiedSessionsPerSkill=3` ⇒ ~90 `user_skills` + ~30 evidences. Presets re-validated.
- [x] **G5 — claimed-but-unverified tail** — the `ProfileSeeder` seeds ~60 `user_skills`/`user_skill_evidences`
  with `is_verified=false`, no `job_simulation_id` (tied to the experiences via
  `user_skill_experience`/`user_skill_education` for the DB CHECK), `user_level` set, `anthropos_level` NULL —
  widening the visible gap; the `user_level` vs `anthropos_level` mechanic intact (verified side never clobbered).
- [x] **Docs** — `corpus/ops/seeding-spec.md` + `corpus/ops/demo/stories-spec.md` updated with the profile-depth
  layer (the ProfileSeeder fan-out + the depth bump + the unverified tail + the widened gap + the live-schema
  landmines).

**Status:** all 5 sections implemented. Code @ `rosetta-extensions` tag `method-acting-m41`. 9 new unit tests,
full stack-seeding suite green `-race`, vet + gofmt clean, go.mod/go.sum byte-identical, every emitted row
dry-insert-validated against the live demo-3 schema (then rolled back — zero pollution). Ready for
`/developer-kit:close-milestone`.
