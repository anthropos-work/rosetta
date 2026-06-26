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

## M44: Hardening

Full hardening ledger: [`hardening-ledger.md`](hardening-ledger.md). Code-of-record in the
`rosetta-extensions` authoring copy (`stack-seeding/seeders/profile_completeness_harden_test.go`,
tests-only → tag bumped to **`method-acting-m44-profile-completeness-fix2`**); this corpus holds
the ledger.

### Pass 1+2+3 — 2026-06-26 (final — stabilized)
Scope manifest (M44-touched, single Go stack): `seeders/persona.go`+`persona_write.go`,
`blueprint.EffectiveSelfRated()`, `seeders/certificates.go`+`projects.go` (NEW), `seeders/profile.go`,
`seeders/users.go` (§D fix1 avatar column). No new-unit-without-handbook finding (two added seeders in
the existing tree, not new top-level units).

**Coverage delta (seeders pkg statements):** 96.5% → **97.5%** (+1.0). All **M44-introduced** seeder
functions now **100%** (`certBankForRole`/`projectBankForRole`/`certFieldOfStudy`/`managerTitle` +
`CertificatesSeeder.Seed`/`ProjectsSeeder.Seed`). Residual sub-100% (`resolveHeroSkills` 95.7%, two
`users.go` casbin grant-error arms) is **pre-existing M42e code**, outside the M44 footprint.

**Tests added** (`profile_completeness_harden_test.go`, **17**, all `-race`-clean across 3 sequential
shuffled runs): §B role-coherence across all 3 cert/project families + cross-family disjointness,
field-of-study branches, the **uint64-modulo bank-index regression** (negative-int-index class),
cert date invariants + the expiry believability mix (some expire, some perpetual), COPY-failure
wraps, the `months<=0` span default; §C the manager cross-seeder-skip-but-gets-profile DAG invariant,
`managerTitle` field branches, the manager **claimed-tail distinctness** (skip
`trajectoryVerifiedCount`, not `EffectiveVerified()=0`); §D member role↔membership coherence, the
offset-0 whole-tail-claimed, the no-pool degradation; and **the highest-value avatar regression** —
the §D fix1 `memberships.picture_url` heal-failure now propagates loudly (atop the existing
`users.picture==memberships.picture_url` cross-column match + backfill idempotency in `seeders_test.go`).

**Bugs fixed inline:** none — the build code (incl. the `fix1` avatar correction) was correct;
hardening pinned its invariants. One Pass-2 test self-correction (the months-default precondition).

**Flakes stabilized:** none — zero flakes (new tests + full suite, 3 sequential shuffled runs + `-race`).

**Knowledge backfill:** no KB-worthy new behavior — the invariants are already in `decisions.md` (D7
avatar column), the seeders' LIVE-SCHEMA doc comments, and `corpus/ops/demo/profile-completeness-spec.md`.

**Stop condition:** stabilized — delta +0.2% on the final pass (< 2%), Step-2b scan clean of
M44-footprint gaps, 0 flakes.
