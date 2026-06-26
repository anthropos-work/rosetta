# M44 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist

- [ ] **(A) Trajectory-aware self-rating** — branch `PersonaSeeder`'s `user_level` UPSERT on
  `Persona.Trajectory`: THRIVING → completed self-assessment; NON-THRIVING → incomplete/absent (keep ~2–3
  verified skills so the Skill-Spotlight chart still renders).
- [ ] **(B1) `CertificatesSeeder`** — NEW `seeders/certificates.go`, surface `'certificates'`, `DependsOn`
  users+personas: 2–3 `public.user_certificates` rows per end-user hero, backdated in the activity span
  (`cert_name` NOT NULL, DATE types, nullable `organization_id`).
- [ ] **(B2) `ProjectsSeeder`** — NEW `seeders/projects.go`, surface `'projects'`, `DependsOn`
  users+personas: 3–4 `public.user_projects` rows per end-user hero (`project_name` NOT NULL, nullable
  `end_date`, `skills` json array).
- [ ] **(C) MANAGER personal data** — remove the `IsManager` skips (`persona.go:121` + `profile.go:125`);
  write a small verified-skill subset (3–8 skills, L1–L2, flat current-state) + a manager-track timeline
  (3 experiences + 1 education) so the manager's OWN `/profile` is populated.
- [ ] **(D) BULK-MEMBER depth** — extend `photoAvatarDataURI` to EVERY member (avatars on
  `/enterprise/members`) + extend `ProfileSeeder` with a shallow timeline (3 short-tenure experiences, 1
  education, a flat ~6-skill claimed tail) for every member.
- [ ] **Docs** — `corpus/ops/demo/profile-completeness-spec.md` **(NEW)** + updates to `seeding-spec.md`
  / `stories-spec.md` for the new surfaces, the trajectory-aware self-rating, the manager unskip, and the
  bulk-member depth.

**Status:** `planned`. DATA DENSITY only — zero platform / next-web edits (no "% complete" widget).
