# M36 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **`membership_skills`** (mapped) — outnumber verified per skill; `skill_name` set (NOT NULL). _Done: `MembershipSkillsSeeder` + the name-bearing `skillref_named.go` resolver + `EffectiveMapped()`; funnel-believability + AI-narrative + idempotency tests. Funnel matches mapped↔verified by skill NAME (D-M36-1)._
- [x] **`tags` + `membership_tags`** — teams/business-units incl. a `mentor` tag. _Done: `TagsSeeder` — a dozen front-loaded business units (one per member) + the cross-cutting `mentor` tag (the Growth-tab Mentors KPI)._
- [x] **`organization_target_roles` + `user_target_roles`** — gap + two-way mobility. _Done: `TargetRolesSeeder` — role-coherent J- target chosen ≠ the member's current role; both sides a believable subset._
- [x] **Succession feeders** — `validation_attempt_*` (already M34) + `interview_extraction_results`, sized to clear the coverage gate. _Done: `SuccessionSeeder` — interview for >20% of members (+ every hero), trajectory-aware summary (struggling = at-risk); clears `too_sparse`→`full`._
- [x] **`job_simulation_feedbacks`** (~2:1) + **assignments fix** (status mix + due_dates + `organization_assignment_sessions`) + skillpath `completed` share. _Done: `FeedbackSeeder` (2:1, polarity-matched) + the assignments status-mix fix (completed/overdue/in-progress/not-started via status + due_date + sessions; the FK takes the skill-path arm, D-M36-2/3) + skillpath completed ~30% (was ~1%)._
- [x] **Org-scale distributions** — claimed-vs-verified gap (`user_level` vs `anthropos_level`) + AI-readiness skills + growth arc; the two employee heroes as the standout high/low rows. _Done: `PopulationEvidenceSeeder` (population over/under-claimer mix); AI-readiness via the membership-skills AI top-up; growth arc already by jobsim-sessions (M34). Coherence: heroes are the standout rows._
- [x] **Docs** — extend `stories-spec.md` (dashboard surfaces) + `seeding-spec.md`. _Done: stories-spec.md § The Workforce dashboard surfaces (M36) + the scope note; seeding-spec.md M36 paragraph + Status; safety.md PerStackIsolated confirming note._
- [x] **Tests** — `stack-seeding` suite green; the Workforce dashboard renders the seeded story. _Done: full unit suite green `-race` (vet+gofmt clean); the opt-in live integration test (`-tags integration`) seeds the full fleet against demo-3 + asserts every dashboard aggregate (funnel/self-eval/teams/targets/succession/feedback/assignments) resolves, leaving the stack clean. The live browser render is the orchestrator's post-build acceptance (the admin view)._

_Last updated: 2026-06-23 (M36 build complete — 6 seeders + 2 fixes + closure-gene extension + live integration test; ext code on rext `main`, tagged `storytelling-m36`)._

## M36: Hardening

### Pass 1 — 2026-06-23
**Scope manifest** (M36-touched, `git diff storytelling-m35..storytelling-m36` in the rext authoring copy — for close-milestone Phase 4 reuse): the 6 dashboard seeders (`membership_skills.go` + `membership_skills_test.go`, `tags.go` + `tags_test.go`, `target_roles.go` + `target_roles_test.go`, `succession.go` + `succession_test.go`, `feedback.go` + `feedback_test.go`, `population_evidence.go` + `population_evidence_test.go`); the named resolver `skillref_named.go` (no co-located test at build); the `assignments.go` status-mix fix + `assignments_statusmix_test.go`; `skillpath_sessions.go` (completed-share bump); `dna/seed_closure.go` + `dna/fidelity_probe.go` (closure-gene extension, both `_test.go`-covered); `blueprint/blueprint.go` + `cmd/stackseed/main.go` (registration/reset-list, `_test.go`-covered); the `dashboard_integration_test.go` (opt-in `-tags integration`). Baseline coverage: seeders 93.0%, dna 87.7%, blueprint 100%.

**Coverage delta (seeders package, the M36-heavy stack):**
- Statements: 93.0% -> 95.5% (+2.5) across the three passes.

**Tests added:**
- `seeders/distribution_helpers_harden_test.go` (NEW): the pure distribution/pool helpers' boundary branches — `memberInAIShare`/`memberInShare` share<=0/share>=1 short-circuits + determinism/monotonicity; `pickDistinctSkills`/`pickDistinctNodeIDs` n<=0/empty/clamp/dedup/offset-stability; `appendDistinctSkill` empty/dup/fresh; `filterAISkills`; `namedSkillPool.take` bounds + copy-safety; `pickDifferentRole` empty-pool arm. The eight targeted helpers 60–85% -> 100%.

### Pass 2 — 2026-06-23
**Tests added:**
- `seeders/dashboard_errorpath_harden_test.go` (NEW): every dashboard seeder's write-failure path — inject a COPY/Exec fault and assert the operator-facing contract (error propagates, wrapped with the seeder NAME + failing TABLE; two-table seeders return the partial total). Covers `membership_skills` / `feedback` / `succession` single-COPY arms, `tags` + `target_roles` first/second-table arms, `population_evidence` Exec-UPSERT arm. Adds reusable interface-level fault wrappers (`failCopyConn` / `failExecConn`). Every dashboard `Seed`'s error-path coverage up (seeders 94.5% -> 95.2%).

### Pass 3 — 2026-06-23
**Tests added:**
- assignments FK-arm chain (in `dashboard_errorpath_harden_test.go`): the 3 FK-ordered write tables (`organization_assignments` -> `local_skill_path_sessions` -> `organization_assignment_sessions`, D-M36-3) — 2nd/3rd-table error arms with partial-total accounting; `assignments.go` -> **100%**.
- named-resolver last branches (in `distribution_helpers_harden_test.go`): `resolveNamedSkillRefs` blank/duplicate role-name skip (the per-member role feed dedup) + `queryNamedSkills::at` malformed-array degrade (more node_ids than names → empty-lineage pairing, no index panic); `skillref_named.go` -> **100%** (seeders 95.4% -> 95.5%).

**Bugs fixed inline:** none — every hardened path confirmed correct behaviour (the build-phase logic held under the deeper edge/error probes).

**Flakes stabilized:** none observed — the new tests passed 3 consecutive sequential `-race` runs (flake gate clean). Full `stack-seeding` suite green `-race` across all 8 packages; `gofmt`/`vet` clean (incl. `-tags integration` compile).

### Stop condition
Scan clean (only the trivial `Surface`/`DependsOn`/`Isolation` registration accessors + a couple of defensive no-pool guards remain — no behaviour to probe; shallow box-ticking declined per the skill's no-shallow-tests rule) + Pass-3 delta +0.3% (< 2% threshold) + flake gate clean. Three passes, +2.5% statements on the seeders package; the 6 new seeders' load-bearing helpers, the named resolver, and `assignments.go` all at 100%.

_Hardening pass: 2026-06-23 (3 passes; 21 new test functions across 2 new harden test files; no bugs surfaced; ext code on rext `main`, tag `storytelling-m36` moved forward)._
