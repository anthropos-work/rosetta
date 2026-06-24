# M41 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist

- [ ] **G3 — `ProfileSeeder` (work history)** — new seeder in rext `stack-seeding` writing 2–3
  `public.user_experiences` per hero: columns `[id, created_at, "user", company, title, description, from, to,
  skills(json), location, location_type, job_role_id]`; deterministic UUIDs; backdated within the story's
  Activity span; role-aligned titles (reuse `jobRoleRefs`); tied to the verified-skill arc.
- [ ] **G3 — `ProfileSeeder` (education)** — same seeder writing 1–2 `public.user_educations` per hero:
  columns `[id, created_at, "user", title, from, to, skills(json), institute, field_of_study]`; deterministic
  UUIDs; backdated within the Activity span.
- [ ] **G5 — bump verified depth** — raise the preset `verified` count (e.g. `8 → ~30`) in
  `stories.seed.yaml:48` + `stories-maya.seed.yaml:36`, flowing through `blueprint.go:227 EffectiveVerified →
  persona.go:132-178 resolveHeroSkills` × `verifiedSessionsPerSkill=3` ⇒ ~90 `user_skills` + ~30 evidences.
- [ ] **G5 — claimed-but-unverified tail** — seed ~60 `user_skills` / `user_skill_evidences` with
  `is_verified=false`, no `job_simulation_id`, `user_level` set, `anthropos_level` NULL — widening the visible
  claimed-vs-verified gap while keeping the `user_level` vs `anthropos_level` mechanic intact.
- [ ] **Docs** — `corpus/ops/seeding-spec.md` + `corpus/ops/demo/stories-spec.md` updated with the
  profile-depth layer (the ProfileSeeder fan-out + the depth bump + the unverified tail + the widened gap).
