# M44 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md)
(the profile-density strand). M44 is **DATA DENSITY ONLY — zero platform / next-web edits** (no UI
"% complete" widget). It extends the verified-skill chain documented in
[`corpus/ops/demo/stories-spec.md`](../../../../corpus/ops/demo/stories-spec.md) — do not reinvent the
`user_skills` / `user_skill_evidences` / `user_experiences` shapes.

> Note: M41's spec-notes flagged that the overview/spec column-lists are guesses and the **live demo
> schema wins** (e.g. `user_experiences.company` turned out `uuid NOT NULL`, not nullable). Re-verify the
> `user_certificates` / `user_projects` column-lists against the live demo-3 schema before build; the
> column notes below are first-pass and may be corrected by the live schema.

## (A) Trajectory-aware self-rating — `persona.go` `user_level` UPSERT
TODO: branch the `user_level` UPSERT on `Persona.Trajectory`. THRIVING → completed self-assessment;
NON-THRIVING → incomplete/absent self-rating state. Keep ~2–3 verified skills for non-thriving so the
Skill-Spotlight chart still renders (sparse, not broken). Keep the `user_level` (claimed) vs
`anthropos_level` (verified) gap mechanic intact.

## (B1) `CertificatesSeeder` (`seeders/certificates.go`, surface `'certificates'`)
TODO: `DependsOn` users+personas. 2–3 `public.user_certificates` rows per end-user hero, backdated in the
activity span. First-pass columns (VERIFY against live schema): `cert_name` NOT NULL, DATE types, nullable
`organization_id`. Deterministic UUIDs; idempotent COPY.

## (B2) `ProjectsSeeder` (`seeders/projects.go`, surface `'projects'`)
TODO: `DependsOn` users+personas. 3–4 `public.user_projects` rows per end-user hero. First-pass columns
(VERIFY against live schema): `project_name` NOT NULL, nullable `end_date`, `skills` json array.
Deterministic UUIDs; idempotent COPY. Confirm the surface name avoids collision with any existing surface.

## (C) MANAGER personal data — unskip `persona.go:121` + `profile.go:125`
TODO: remove the `IsManager` skips. Write a small verified-skill subset (3–8 skills, L1–L2 band, FLAT
current-state, no growth arc) + a manager-track timeline (3 experiences + 1 education), so the manager's
OWN `/profile` is populated like a real member.

## (D) BULK-MEMBER depth — `photoAvatarDataURI` to EVERY member + shallow `ProfileSeeder`
TODO: extend `photoAvatarDataURI` to every member (not heroes-only) → `/enterprise/members` shows avatars.
Extend `ProfileSeeder` to write a shallow timeline per member: 3 short-tenure experiences, 1 education, a
flat ~6-skill claimed tail. Default ALL members get avatar+career; depth shallow. (Open: every-member
~3K rows/org vs sampled.)

## Reuse (no new mechanics)
- `resolveJobRoleRefs` / `resolveNamedSkillRefs` (role node-id + title; skill names → node-ids).
- `dateOnly` (date truncation), `legacySkillsJSON` (the `{skills:[{level,name}]}` envelope shape).
- `roleTitle` degradation, `CopyRowsIdempotent` (the idempotent COPY).

## Delivers — `corpus/ops/demo/profile-completeness-spec.md` (NEW)
TODO: the "complete profile" rubric — identity + content + semantic layers, per-vantage member vs
manager, each component mapped to its seeding surface + a Playwright acceptance assertion (the M42e/M42m
coverage-protocol gate).

## Pre-flight audits — section A (Trajectory-aware self-rating)
KB-fidelity audit (Phase 0b): **GREEN** — report `kb-fidelity-audit.md` (sha `dafb6ce`-base, run 2026-06-26).
Topic→doc→code triples (verified against the rext authoring copy, all ALIGNED):

| Topic | Doc | Code |
|---|---|---|
| verified-skill chain + `user_level` self-rating | `corpus/ops/demo/stories-spec.md` §claimed-vs-verified | `seeders/persona.go` (`selfEvalLevel`, l.212), `persona_write.go` (`upsertEvidenceSQL`) |
| profile depth (timeline + claimed tail) | `stories-spec.md` §M41 + `corpus/ops/seeding-spec.md` l.281-297 | `seeders/profile.go` |
| seeding isolation + write surfaces | `seeding-spec.md` §isolation-boundary l.84-106 | `stack-seeding/isolation/`, `seeders/users.go` |
| coverage semantic gate | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/` |
| photo avatar / EVERY member | `coverage-protocol.md` persona self-consistency | `seeders/avatar.go` (`photoAvatarDataURI`), `seeders/users.go:156` |

KEY LOAD-BEARING FACTS for the build (all code-verified):
- Manager skips: `persona.go:121` + `profile.go:125` (`if p.IsManager() { continue }`). §C removes both.
- `photoAvatarDataURI(uid)` ALREADY runs for EVERY population user (`users.go:156`, since M42e P4) — §D's
  avatar half is already satisfied; §D = verify + bulk-member shallow career.
- BOTH `PersonaSeeder` + `ProfileSeeder` iterate **only `st.Heroes`** — bulk-member depth (§D) = extend to
  the non-hero population (the `user1..userN` slots `UsersSeeder` already creates).
- Seeder registry: `cmd/stackseed/main.go:239-256` (`reg.MustRegister(...)`). §B1/B2 add two entries.
- Reuse: `resolveJobRoleRefs`/`resolveNamedSkillRefs` (closure gate), `dateOnly`, `legacySkillsJSON`,
  `companyFor`, `deterministicUUID`, `CopyRowsIdempotent("id")`, the `personaRows.flush` COPY pattern,
  `roleTitle` via `jobRoleRefs.forName`, `nullIfEmpty`, the `isolation.PerStackIsolated` + audit pattern.
- `EffectiveMapped()` already returns 0 for a manager — §C must NOT rely on it for the manager's claimed
  tail; §C writes a small explicit verified subset + a manager-track timeline directly.
