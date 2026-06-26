# M44 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist

- [x] **(A) Trajectory-aware self-rating** — `Persona.EffectiveSelfRated()` + `PersonaSeeder` write
  `user_skill_evidences.user_level` only for self-rated heroes; struggling = NULL (incomplete self-rating),
  verified side intact so the chart renders. rext `4614089`.
- [x] **(B1) `CertificatesSeeder`** — NEW `seeders/certificates.go`, surface `'certificates'`. LIVE-SCHEMA
  CORRECTED: `public.user_certifications` (NOT `user_certificates`), 2–3 per end-user hero, `certification`
  NOT NULL, DATE `from`/`to`, NO `created_at`, NO `organization_id`. rext `5db10ac`.
- [x] **(B2) `ProjectsSeeder`** — NEW `seeders/projects.go`, surface `'projects'`. `public.user_projects`,
  3–4 per end-user hero, `title` (NOT `project_name`) NOT NULL, nullable `to` (NOT `end_date`), `skills` json,
  NO `organization_id`. rext `5db10ac`.
- [x] **(C) MANAGER personal data** — removed the `IsManager` skips (`persona.go` + `profile.go`); manager gets
  a flat 3–8 verified set (L1–L2 band, self-rated) + a manager-track timeline (leadership ladder, 3 exp + 1 edu)
  + a small claimed tail (8). rext `56cef82`.
- [x] **(D) BULK-MEMBER depth** — avatar half ALREADY satisfied (`photoAvatarDataURI` on every member since
  M42e P4); `ProfileSeeder` now runs a 2nd pass over non-hero slots (`seedMemberProfile`: 3 short-tenure exp +
  1 edu + flat ≤6 claimed tail). rext `bc5b94c`.
- [x] **Docs** — `corpus/ops/demo/profile-completeness-spec.md` **(NEW)** + updates to `seeding-spec.md` /
  `stories-spec.md` / demo `README.md`. corpus `42ffd53`.

**Status:** all sections IMPLEMENTED + unit-tested (full rext suite green, `-race`/integration-compile clean).
DATA DENSITY only — zero platform / next-web edits (no "% complete" widget). rext tag
`method-acting-m44-profile-completeness`. Live demo-up acceptance pending (the reproducible-mechanism proof).
