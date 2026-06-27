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

**Status:** `archived` (completed 2026-06-26) — all sections IMPLEMENTED + unit-tested (full rext suite green, `-race`/integration-compile clean).
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

## M44: Final Review

`/developer-kit:close-milestone` cross-cutting review (2026-06-26). Corpus diff is **docs-only** (the
code-of-record lives in the `rosetta-extensions` authoring copy @ tag
`method-acting-m44-profile-completeness-fix2`, a SEPARATE repo — not part of the corpus merge).

### Scope
- [x] All 6 sections checked in `progress.md`; `overview.md` `In:` list fully delivered (§A–§D + Docs).
- [x] Deferral re-audit (Phase 1b, `/developer-kit:audit-deferrals --scope=milestone`) — **GREEN**:
  0 deferrals (M44 `Out:` = Fate-2 to already-planned M45/M46; inherited DEF-M40-01 closed in M42m).
  Report: [`audit-deferrals/deferral-audit-2026-06-26-m44-close.md`](audit-deferrals/deferral-audit-2026-06-26-m44-close.md).

### Code Quality
- [x] Corpus diff docs-only; 0 broken markdown links across the 4 touched docs (link-integrity check).
- [x] rext authoring copy clean @ `fix2`; `gofmt -l seeders/` clean; `go vet`/`-race` clean.

### Documentation
- [x] [should-fix] `seeding-spec.md` `## Status` cited the M44 BUILD tag; the code-of-record advanced
  to `-fix2` (the §D avatar fix + 3 hardening passes). Updated to reference `…-fix2` + the avatar-column
  correction + the hardening delta. (FIXED)
- [x] New `profile-completeness-spec.md` accurate + discoverable (linked from `demo/README.md`,
  `seeding-spec.md`, `stories-spec.md`); §A/§C `stories-spec.md` annotations correct (incl. the
  PopulationEvidenceSeeder-unaffected non-obvious interaction).

### Tests & Benchmarks
- [x] `stack-seeding` full module `go test -race -count=1 ./...` — GREEN. seeders pkg 310 Test/Fuzz funcs;
  module total **567** (was 538 @ `method-acting-m42m-harden-final` tag → **+29**). Harden file
  `profile_completeness_harden_test.go` = 17 tests (reconciles with the hardening ledger). No new benchmark
  surface (deterministic seeders).

### Decision Triage
- [x] D1/D2/D3/D7 — already blended into the spec docs as primary `Delivers →` content (M44 §A/§B/§C/§D
  section anchors are the back-trace); accuracy verified. No re-blend needed.
- [x] D4 (shared ProfileSeeder path) / D5 (surface names) / D6 (superseded by D7) / KB-1 (Phase-0b verdict)
  → archive (maintainer-only / audit record); stay in `decisions.md`.

## M44: Completeness Ledger

Section milestone (`milestone_shape: section`). Every `overview.md` `In:` scope item placed into exactly
one three-fate category.

### Done (Fate 1 — landed in this milestone)
- **(§A) Trajectory-aware self-rating** — `PersonaSeeder` `user_level` branch on `EffectiveSelfRated()`;
  struggling = NULL/incomplete, verified side intact (chart renders). RENDER-VERIFIED on demo-3.
- **(§B1) `CertificatesSeeder`** (surface `'certificates'` → `public.user_certifications`, 2–3/hero) +
  **(§B2) `ProjectsSeeder`** (surface `'projects'` → `public.user_projects`, 3–4/hero) — schema-corrected
  vs the live demo-3 DB. RENDER-VERIFIED (Maya's Certs + Projects sections populated).
- **(§C) Manager personal data** — `IsManager` skips removed; modest flat verified set + manager-track
  timeline + small claimed tail. RENDER-VERIFIED (Dan has his own profile + skills).
- **(§D) Bulk-member depth** — every non-hero gets a shallow career + the avatar-column fix
  (`memberships.picture_url`, the build's render-miss). RENDER-VERIFIED (members list 20/20 photos, was 0/20).
- **Docs** — NEW `profile-completeness-spec.md` (+ the §D avatar GOTCHA) + `seeding-spec.md` /
  `stories-spec.md` / demo `README.md` updates.
- **Hardening** — 3-pass sweep, 17 tests, seeders stmt coverage 96.5%→97.5%, 0 bugs, 0 flakes.
- **Close fix** — `seeding-spec.md` Status line corrected to the final `-fix2` code-of-record tag.

### Confirmed-covered (Fate 2 — already planned in another milestone of this release)
- **LLM-generated profile content** → owned by **M45** (generation engine) + **M46** (org-scale fill) —
  `In:` lists confirm; M44 seeders are deterministic by design. No plan edit.
- **Deep per-fill-member career narratives** → owned by **M46** (LLM-generated richer fill) — `In:` list
  confirms; M44 keeps fill shallow by design (D3). No plan edit.

### Annotated (Fate 3 — attached to a release-milestone at close)
- None. No sibling `overview.md` edited.

### Dropped
- **Platform / next-web UI "% complete" widget** — dropped by design (the user's explicit DATA-DENSITY-only
  choice, `overview.md` `Out:` + goal line). Not a deferral; a deliberate scope exclusion (zero platform
  edits is a v1.10 release invariant).

### Release-scope-breaking deferral (escape hatch — requires user sign-off)
- None.

**Verdict:** all `overview.md` `In:`-list scope items delivered as **Fate 1** (§A–§D + Docs, all
render-verified + hardened). The two `Out:` items are **Fate-2** to already-planned M45/M46; the "% complete"
widget is a by-design **Drop**. Nothing annotated, nothing escape-hatch-deferred. **Zero items require sign-off.**
