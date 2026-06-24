# M41 Spec Notes

Technical notes accumulate here during build. The authoritative design is
[`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review 2026-06-24;
root-cause workflow w7t4wq2z4). M41 extends the verified-skill chain documented in
[`corpus/ops/demo/stories-spec.md`](../../../../corpus/ops/demo/stories-spec.md) — do not reinvent the
`user_skills` / `user_skill_evidences` shapes.

## G3 — `ProfileSeeder` (rext `stack-seeding`)
TODO: net-new seeder; neither target table is written by any current seeder (confirmed 0 rows DB-wide).

## `public.user_experiences` (work history)
TODO: columns `[id, created_at, "user"(uuid FK), company(nullable), title, description, from, to, skills(json),
location, location_type(enum), job_role_id]`; 2–3 rows/hero; deterministic UUIDs (`deterministicUUID` prefix);
backdated within the story's Activity span; role-aligned titles (reuse the resolved `jobRoleRefs`); tied to the
verified-skill arc.

## `public.user_educations` (education)
TODO: columns `[id, created_at, "user", title, from, to, skills(json), institute, field_of_study]`; 1–2
rows/hero; deterministic UUIDs; backdated within the Activity span.

## `/profile` timeline read path (read-only — no platform edit)
TODO: confirm the timeline reads `ent.UserExperience` / `ent.UserEducation` via `GET_TIMELINE`; M41 supplies
rows only — no next-web / app / cms / jobsimulation edits.

## G5 — verified-depth bump
TODO: bump the preset `verified` count (e.g. `8 → ~30`) in `stories.seed.yaml:48` + `stories-maya.seed.yaml:36`;
trace `blueprint.go:227 EffectiveVerified → persona.go:132-178 resolveHeroSkills` × `verifiedSessionsPerSkill=3`
(`persona.go:46`) ⇒ ~30 distinct (10 role-coherent + auto-top-up from the flat public pool past the role's 10)
× 3 = ~90 `user_skills` + ~30 evidences.

## G5 — claimed-but-unverified tail
TODO: seed ~60 `user_skills` / `user_skill_evidences` with `is_verified=false`, no `job_simulation_id`,
`user_level` set, `anthropos_level` NULL ⇒ "overall ≈ 90 = ~30 verified + ~60 claimed-unverified"; widens the
visible claimed-vs-verified gap. Keep the `user_level` (claimed) vs `anthropos_level` (verified) mechanic intact.

## Leave-as-is: the `mapped:22` wiring
TODO: confirm `mapped:22` feeds `membership_skills` (the org funnel), NOT `user_skills` — do not touch it;
profile depth comes from `user_skills`.
