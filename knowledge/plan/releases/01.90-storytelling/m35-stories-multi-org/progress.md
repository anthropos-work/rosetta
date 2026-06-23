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

## M35: Final Review

Close-milestone review (2026-06-23). Deferral re-audit GREEN
([`audit-deferrals/deferral-audit-2026-06-23.md`](audit-deferrals/deferral-audit-2026-06-23.md));
scope GREEN (all 8 sections checked, #M34-D7 landed as D-M35-4, 0 TODOs in the M35 diff, done-bar met
live on demo-3); test suite GREEN (8 packages, `-race` clean).

### Scope
- [x] No scope gaps — #M34-D7 (the M34→M35 routed deferral) landed in full as D-M35-4 (both parts); 0 code TODOs in the diff.

### Code Quality
- [x] [should-fix] `skillPool.at()` (`seeders/taxonomyref.go:63`) did NOT normalize a negative index, unlike its protected sibling `jobRoleRefs.at()` (`jobroleref.go:50`, which the D-M35 negative-modulo Note already hardened). Safe-by-current-usage (its sole call site passes a loop counter), but a latent panic the day a hash is passed — same `int(hashInt) % len` shape the sibling guards. → normalize + regression test.
- [x] [nice-to-have] (no change) module-isolation duplication (`storyEmailDomain`/`slugifyBP` in blueprint vs seeders) is deliberate — blueprint must not import seeders (D-M35 note); the `OrgID`/`orgClerkID` back-compat alias in `org.go` is intentional (M34 test compile). Adversarial Phase 2c: 3 scenarios (StoryOrgID dup-id, array-length mismatch, exhausted flat pool) — all handled by validation/design.

### Documentation
- [x] `stack-seeding/README.md` handbook test-count reconciled to ground truth (326 → 346 non-integration test funcs; "381+ incl. subtests" → 424).

### Tests & Benchmarks
- [x] No coverage gaps — multi-org closure across orgs, name-collision regression, negative-modulo regression, no-fabrication error paths, trajectory fidelity, per-seeder COPY/Exec error propagation all covered (blueprint 100% / seeders 98.0%). Added a `skillPool.at()` negative-index regression test (the should-fix above).

### Decision Triage
- [x] D-M35-1..7 + the negative-modulo Note — all are implementation-mechanics already blended into `stories-spec.md`/`seeding-spec.md` during build (multi-org normalization, the first-story-keeps-default-org rule, the job-role resolver, manager-rides-aggregates, the name-uniqueness invariant). The options-considered detail stays in `decisions.md` as archive. No further blending owed.

_Final review last updated: 2026-06-23._

## M35: Completeness Ledger

Every overview.md In-scope item, fated (three-fate rule). **All delivered as Fate 1.**

### Done (Fate 1 — landed in M35)
- The `stories[]` blueprint (per-hero vantage/trajectory/skills, per-story org/narrative; supersedes `stack.seed.yaml`) — `blueprint/stories.go` + `EffectiveStories()`.
- Multi-org `OrgID`+`orgClerkID` per-story, threaded through all 8 seeders + the Clerkenstein org-claim alignment (data-side, D-M35-2) — `org/users/identity/jobsim_sessions/assignments/persona/activity/skillpath_sessions`.
- `PersonaSeeder` scaled to the 2-story × 3-hero roster (Cervato + Solvantis).
- #M34-D7 (the M34→M35 routed deferral) — BOTH parts landed as D-M35-4: `len(heroes) <= size` validation + declaration-order collision-free slots + warning; short-role-pool flat top-up.
- Trajectory logic (thriving=dense/rising/under-claim vs struggling=sparse/low/over-claim).
- Supporting-population fidelity (real replayed `job_role_id`/name via `jobroleref.go`, ramped `joined_at`, deduped real names).
- Single-org default preserved (legacy blueprint → byte-identical one-story view; full suite green unchanged).
- Docs (`stories-spec.md` + `seeding-spec.md` + `/stack-seed` SKILL.md + the rext README).
- The two close findings (skillPool.at negative-index; handbook count) — landed this close.

### Confirmed-covered (Fate 2 — already owned by another v1.9 milestone)
- The org-aggregate dashboard surfaces (`membership_skills`, tags, target-roles, feedback, succession) → **M36** (its `overview.md` In-list owns them).
- The Clerkenstein multi-identity seat-switch + the presenter cockpit (incl. the literal browser-pixels render of a hero's individual profile via login-AS-a-hero) → **M37/M38**.

### Annotated (Fate 3 — attached to a milestone at close)
- None. (M35 wrote nothing new to route.)

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch)
- None.

**All scope items delivered in this milestone. Nothing routed-new, dropped, or escape-hatch-deferred. Clean close.**

_Completeness ledger: 2026-06-23._
