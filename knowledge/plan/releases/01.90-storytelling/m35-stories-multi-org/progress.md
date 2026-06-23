# M35 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **`stories[]` blueprint** — per-hero `vantage`/`trajectory`/`skills`, per-story `org`/`narrative`; supersedes `stack.seed.yaml` for demo stacks
- [x] **Multi-org parameterization** — per-story `OrgID` + `orgClerkID` threaded through the 4 seeders (org/users/identity/jobsim-sessions/assignments) + Clerkenstein org-claim alignment _(first story keeps the Clerkenstein default org; all 8 seeders threaded, not just 4)_
- [x] **`PersonaSeeder` roster scaling** — the 2-stories × 3-heroes v1 roster (Cervato + Solvantis) _(+ #M34-D7: collision-free declaration-order hero slots + warning, short-pool flat top-up)_
- [x] **Trajectory logic** — thriving (dense/rising/under-claim) vs struggling (sparse/low/over-claim)
- [x] **Supporting-population fidelity** — `job_role_id`+name (real replayed via `jobroleref.go`), ramped `joined_at`, names on non-hero members
- [x] **Single-org default preserved** — existing `dev-min`/preset path still passes (the legacy blueprint resolves to a byte-identical one-story view; full M34/M7 suite green unchanged)
- [x] **Docs** — extended `stories-spec.md` (the model) + `seeding-spec.md` (blueprint supersession) + `/stack-seed` SKILL.md + the rext module README; M35 `decisions.md`
- [x] **Tests** — `stack-seeding` suite green (326 test funcs, +24); closure gene green across all orgs (both integration tests pass live against demo-3); `-race` clean

_Last updated: 2026-06-23 (all sections complete; ext tag `storytelling-m35`)._

## M35: Hardening

Test-deepening pass over the M35-touched rext `stack-seeding` code (16 source + 8 test files;
all work in the rext authoring copy on `main`, tag `storytelling-m35` moved forward to the new
HEAD). 3 passes; stopped on the stabilization condition (delta < 2% + scan clean + 0 flakes).

### Pass 1 — 2026-06-23
**Scope manifest:** `blueprint/{blueprint.go,stories.go}` + `seeders/{persona,users,userprofile,
jobroleref,identity,org,activity,taxonomyref,assignments,jobsim_sessions,skillpath_sessions,
helpers}.go`. Baseline coverage: blueprint 90.9%, seeders 95.6% (`-race`+vet clean).

**Coverage delta (milestone-touched files):**
- blueprint: 90.9% → 100.0% statements
- seeders: 95.6% → 96.7% statements

**Tests added:** blueprint — `EffectiveArc` (was 0%), the story-level `Validate` error branches,
the default-admin-email derivation, `storyEmailDomain`/`slugifyBP` degenerate fallbacks (5 funcs).
seeders — direct unit tests for the pure trajectory helpers (`clampScore`/`clampComp` incl. v>100,
`trajectoryArcGain` incl. declining, `trajectoryVerifiedCount` density ordering, `selfEvalLevel`
bias+clamp, `trajectoryLevelBand` ordering) + the production-isolation contract sweep (8 M35
seeders each declare `PerStackIsolated`; `ActivitySeeder.Isolation()` was 0%).

**Bugs fixed inline:**
- **Supporting-population name collision** (orchestrator finding: a live multi-org seed produced
  two extra members also named "Leah Donovan", a hero). `nameForIndexAvoiding` re-rolls
  deterministically off a reserved set (hero names + names already emitted), index-disambiguating
  the surname only if the bank is exhausted. The first roll is **byte-identical** to the old
  `nameForIndex`, so the hero-free legacy/dev-min path is unchanged (suite stays green). Commit
  `8165372`. Regression test pins: no name twice in a 50-person org, each hero name exactly once,
  byte-identical first roll, exhausted-bank fallback. (Proven to FAIL on pre-fix code.)

**Flakes stabilized:** none seen.

**Knowledge backfill:** the name-uniqueness invariant blended into `seeding-spec.md` (see below).

### Pass 2-3 — 2026-06-23
**Coverage delta:** seeders 96.7% → 98.0% statements (blueprint held at 100%).

**Tests added (error paths + edges):** the jobroleref no-fabrication ERROR path (failed read →
unavailable pool, never fabricated) + the defensive array-length-mismatch pairing; `resolveHeroSkills`
directly across all branches incl. the malformed-flat-entry (`""`/dup) skip; the casbin g2-grant
Exec-error propagation (bulk + single + empty-list no-op) — the Members-page-403 failure mode;
`storyAssignerMembership` first-vs-later-story (distinct per-org assigners); `personaUserIndexFor`
defensive past-size clamp + not-in-roster fallback. `storyConn` gained `jobRoleQueryErr` /
`mismatchNames` / `casbinExecErr` injection hooks. Commit `764ec07`.

**Bugs fixed inline:** none (Pass 2-3 are pure deepening).

**Flakes stabilized:** none. Flake gate: 3 consecutive clean sequential runs of all new tests.

### Stop condition
Pass 3 delta +0.7% statements (< 2%); the qualitative scan found nothing new worth adding; 0 flakes.
The residual sub-100% functions are uniform DB-write COPY-error wrappers inside seeder `Seed()`
methods (structurally identical to the persona `flush()` paths M34 already covers) — left for
`/developer-kit:close-milestone` Phase 4 defense-in-depth rather than gold-plating here.

_Hardening last updated: 2026-06-23._
