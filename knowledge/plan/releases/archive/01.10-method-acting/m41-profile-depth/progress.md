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

**Status:** `archived` (completed 2026-06-25). All 5 sections implemented + hardened + closed. Code @
`rosetta-extensions` tag `method-acting-m41` → `0346113`. Full stack-seeding suite green `-race`, vet + gofmt
clean, go.mod/go.sum byte-identical, every emitted row dry-insert-validated against the live demo-3 schema (then
rolled back — zero pollution). Closed via `/developer-kit:close-milestone` (GREEN, 0 blocking — see § M41: Final
Review + § M41: Completeness Ledger below).

## M41: Hardening

### Pass 1+2 — 2026-06-25
Scope manifest (M41-touched, single Go stack): `stack-seeding/seeders/profile.go` +
`stack-seeding/seeders/profile_write.go` (both covered by `profile_test.go`); `cmd/stackseed/main.go`
(ProfileSeeder registration only — trivial); `presets/stories*.seed.yaml` (data, no test surface). No
new-unit-without-handbook finding (the `profiles` surface is an added seeder in the existing `stack-seeding`
tree, not a new top-level unit; `stack-seeding/README.md` already documents the surface family).

**Coverage delta (per-function, profile.go + profile_write.go):**
- 75.9% avg → **100.0%** avg (+24.1) — **zero uncovered statements** on both M41 files.
- Functions lifted to 100%: `companyFor` (66.7), `fieldOfStudyFor` (57.1), `degreeTitle` (66.7),
  `experienceTitle` (85.7), `combinedNamedPool` (86.7), `pickSkillNames` (90.9), `addCompany` (83.3),
  `flush` (90.0), `seedHeroProfile` (97.7), `seedClaimedTail` (91.3→100 via the direct blank-node-id test).

**Tests added** (`profile_harden_test.go`, 24 tests, all `-race`-clean across 3 sequential flake runs):
- profile_harden_test.go: 6 error-path/regression + 6 invariant + 4 boundary + 8 value-helper-branch.
- Error paths: `flush()` companies-COPY error (wrapped, names table, 0 rows), mid-chain user_skills-COPY
  error (partial total), per-row claimed-evidence Exec error.
- Invariants: the claimed-evidence UPSERT `WHERE is_verified=false` guard NEVER clobbers a verified row
  (the gap mechanic's safety-critical SQL) + idempotent re-upsert; deterministic byte-identity across runs
  (deterministic fields only — wall-clock audit cols excluded by design); the backdated date progression
  (strictly-older jobs, ≥1900, exactly one open-ended current role, education before earliest job);
  date-only UTC truncation; role-aligned title arc + current-role-at-story-org; cross-hero company dedup;
  the per-stack-isolated audit contract.
- Boundary: small-pool graceful tail; exactly-verified → no tail; empty-org-name fallback; json escaping
  of quote/backslash skill names; `seedClaimedTail`'s blank-node-id no-fabrication skip (direct unit, since
  `combinedNamedPool` strips blanks upstream → the in-seeder guard is defense-in-depth, unreachable via `Seed`).

**Bugs fixed inline:** none — no production bug surfaced. The build code was correct; hardening pinned its
invariants. (One test self-correction: the determinism test initially over-asserted by comparing the
wall-clock `created_at`/`updated_at`/`acquired_at` columns — refined to compare only the deterministic
fields, which is the actual reproducibility contract.)

**Flakes stabilized:** none — zero flakes across 3 sequential runs.

**Live validation:** the full company→experience(current+closed)→education→claimed-user_skill→guarded
claimed-evidence-UPSERT chain dry-inserted against the live demo-3 schema inside a rolled-back transaction —
all shapes satisfy the live CHECKs/FKs/enums; rollback verified clean (0 probe rows leaked).

**Knowledge backfill:** no KB-worthy findings — the hardening pinned invariants already documented in
`spec-notes.md` (the LIVE-SCHEMA CORRECTIONS), `decisions.md` (M41-D2/D3/D4 — the company FK, the provenance
edge, the never-clobber guard), and `corpus/ops/seeding-spec.md` / `corpus/ops/demo/stories-spec.md` (the
profile-depth layer). No new system behavior was discovered; nothing to blend.

**Rext commits:** `63bcceb` (`harden(M41): deepen ProfileSeeder tests …`); the `method-acting-m41` tag
moved `18c4edb → 63bcceb`. go.mod/go.sum byte-identical.

### Stop condition
Pass 2 closed the last 2 uncovered blocks → 0 uncovered statements on both M41 files (100% per-function).
A Pass 3 Step-2b scan finds nothing new worth adding; coverage delta would be 0. Loop terminated (scan clean
+ full coverage + no flakes), well within the 5-pass cap.

## M41: Final Review

Review found **8 findings**: 0 scope · 1 code-quality (nice-to-have) · 1 adversarial · 5 docs · 1
decision-triage. All addressed fully (no partial fixes). Deferral re-audit GREEN
([`audit-deferrals/deferral-audit-2026-06-25-m41-close.md`](audit-deferrals/deferral-audit-2026-06-25-m41-close.md)).

### Scope
- [x] No scope gaps — all 5 section checkboxes checked; 3 open-questions resolved in build (M41-D5/D7 + the
  3-exp/1-edu envelope); 3 `Out:` items are release-sibling partitions / leave-as-is, not dropped scope. 0
  TODO/FIXME/HACK in M41 source.

### Code Quality
- [x] [nice-to-have] `combinedNamedPool` uses `rolePool.take(rolePool.len())` (alloc+copy where iteration would
  do) — left as-is: trivial, not load-bearing, the slice is tiny (≤role-pool size). vet + gofmt clean;
  consistent with persona/users seeders (interface, deterministicUUID, COPY/flush, DependsOn). No must/should-fix.

### Adversarial (Phase 2c)
- [x] AR-1 — empty-`eduIDs` round-robin modulo-by-zero: added `TestSeedClaimedTail_EmptyEducationsNoPanic`
  (the in-seeder `len(eduIDs) > 0` guard handles it; code was already correct). AR-2 (never-clobber) + AR-3
  (partial-UNIQUE under round-robin) already pinned by the harden pass. Recorded in `decisions.md`.

### Documentation
- [x] `seeding-spec.md` — completed the `from<=to` CHECK to `from<=to OR to IS NULL` (the current-role NULL case).
- [x] `seeding-spec.md` — added the claimed-tail never-clobber UPSERT guard sentence (was only in stories-spec).
- [x] `CLAUDE.md` — extended the stories-spec.md description to cover v1.10 M39 (identity) + M41 (ProfileSeeder).
- [x] `stack-seeding/README.md` — named `ProfileSeeder` in the `seeders/` row.
- [x] `stack-seeding/README.md` — Status section: added the v1.10 profile-depth layer + reconciled the test
  count **406 → 496** (ground-truth `grep` of Test/Fuzz funcs across 8 packages, post-AR-1).

### Decision Triage
- [x] M41-D2 (company NOT NULL FK) → blend into stories-spec.md (already present; added `(#M41-D2)` tag).
- [x] M41-D3 (provenance edge tie) → blend into stories-spec.md (already present; added `(#M41-D3)` tag).
- [x] M41-D4 (separate guarded UPSERT) → blend into stories-spec.md + seeding-spec.md (added `(#M41-D4)` tag).
- [x] M41-D5 (thriving-heroes verified split) → blend into stories-spec.md (added `(#M41-D5)` tag).
- [x] M41-D6 (combined-pool offset) → blend into stories-spec.md (added `(#M41-D6)` tag).
- [x] M41-D1 (separate ProfileSeeder) → archive (maintainer-only structural choice).
- [x] M41-D7 (heroes-only) → archive (maintainer-only scope choice).

## M41: Completeness Ledger (Phase 9, section variant)

### Done (Fate 1 — delivered in this milestone)
- **G3 work history** — `ProfileSeeder` writes 3 `user_experiences` + a `companies` row per employer per
  end-user hero (deterministic, backdated, role-aligned, current role `to`=NULL; live-schema-correct).
- **G3 education** — 1 `user_educations` per hero, backdated before the work history.
- **G5 verified depth bump** — `verified: 8 → ~30` for thriving heroes (Maya 30, Sara 28) in the presets.
- **G5 claimed-but-unverified tail** — ~60 `user_skills`/`user_skill_evidences` (`is_verified=false`, no
  `job_simulation_id`, tied to experiences/educations, `user_level` set, `anthropos_level` NULL), guarded
  UPSERT never clobbers the verified side.
- **Docs** — `seeding-spec.md` + `stories-spec.md` carry the profile-depth layer (+ the close-fix completions
  and #M41-Dx tags); `CLAUDE.md` + `stack-seeding/README.md` updated.
- **Close-added** — the AR-1 adversarial empty-`eduIDs` guard test (rext `0346113`).

### Confirmed-covered (Fate 2 — already owned by another release-milestone)
- **DEF-M40-01 — KPI "AI simulations completed" = 0** → owned by **M42e + M42m** (M40-D7). Inherited single
  deferral; the per-vantage coverage exit gate already encompasses a non-zero completed-KPI. No plan edit.
  Not aged-out (both destinations open). Confirmed at this close's deferral re-audit (GREEN).

### Annotated (Fate 3 — attached to another release-milestone at this close)
- None.

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch — requires user sign-off)
- None.

**Verdict:** all G3+G5 In-list scope items delivered as **Fate 1**; the one open deferral is the **inherited**
DEF-M40-01 (Fate-2, owned by M42e/M42m). Nothing annotated, dropped, or escape-hatch-deferred. **Zero items
require sign-off.** The 3 overview `Out:` items are release-sibling partitions / explicit leave-as-is; the 3
open-questions were resolved in build (M41-D5/D7 + the 3-exp/1-edu envelope).
