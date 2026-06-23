# M34 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **G14 fix** — `jobsim_sessions.go`: valid `status`/`completion_status`/`result_status`/token + full `SIMULATION_TYPE_*` + continuous mid-skewed score + per-user growth arc + ASSESSMENT/HIRING share *(rext `c0872be`)*
- [x] **`TaxonomyRefs` resolver** — real `skiller.skills.node_id` + `skillsByRole`, empty-pool fallback (mirror `contentref.go`) *(rext `87e5377`)*
- [x] **`PersonaSeeder`** — the 7-table chain per (hero × skill), incl. `user_level` + `result_status` (the `seed.sql` omissions) *(rext `7817553`; FK fix `dad8c72`)*
- [x] **`users.go` patch** — real names + avatars + org-domain emails *(rext `8e1b57a`)*
- [x] **Closure gene** — data-DNA: 0 dangling node_ids, `datadna measure-closure` (mirror the M23 cross-surface gene) *(rext `1e323e7`)*
- [x] **Maya proof** — integration test (the automated half) seeds the chain against demo-3: profile (18 verified skills) + Spotlight chart (18 datapoints) + the claimed-vs-verified gap render + closure green. *Live browser render = the orchestrator's post-build step.* *(rext `dad8c72`)*
- [x] **Docs** — NEW `corpus/ops/demo/stories-spec.md` + `seeding-spec.md` / `safety.md` / demo `README.md` / `CLAUDE.md` updates *(rosetta worktree)*
- [x] **Tests** — `stack-seeding` suite green (`-race` clean); integration test opt-in (`-tags integration`); zero platform-repo edits

_Last updated: 2026-06-23 (all sections landed; rext tagged `storytelling-m34`)._

## M34: Hardening

Scope manifest (Phase 1, `git diff` c0872be~1..storytelling-m34 in the rext authoring copy — M34's
code lives there; the rosetta worktree holds only docs): 14 touched source files in one Go stack
(`stack-seeding`), all under `seeders/` + `dna/`. Highest-priority gaps at baseline were the seeder
error paths (`flush`, `users.Seed`), the boundary clamps (`selfEvalLevel`, `growthArcScore`,
`hexToken`), and two load-bearing invariants with no seeder-level test (the `personaIndexMap`
collision rule + the D2 hero-rides-on-population-index bridge). Coverage tooling: Go native
`-cover` (no instrumentation needed); no golangci-lint configured — `go vet` is the project's check.

### Pass 1 — 2026-06-23
**Coverage delta (seeders package):** Statements 93.7% → 96.2% (+2.5). (dna 87.7%, untouched — its M34
file `seed_closure.go` was already 100%.)

**Tests added** (`persona_harden_test.go`, `seeders_harden_test.go`):
- persona: 3 error-path (flush COPY-error propagation + partial-total + evidences-UPSERT-error), 6 edge
  (selfEvalLevel clamps, personaUserIndex Size≤0, personaIndexMap collision, short-role-pool, unresolved-role
  flat fallback, audit-records-every-write)
- helpers/jobsim/users: 4 edge (hexToken length clamps, growthArcScore score clamps, stackHost empty
  fallback, UsersSeeder isolation contract)

**Bugs fixed inline:** none (the build + integration test already caught the one real bug — the
`validation_attempt_result_id` FK — in the build phase; harden found no new defects).

**Flakes stabilized:** none (suite is deterministic by construction — seeded hashes, no time/random).

**Knowledge backfill:** no KB-worthy findings this pass (the edge/error behaviors confirmed the
documented spec invariants — the chain, the clamps, the no-fabrication fallback — rather than
surfacing new truths). The doc-half (`stories-spec.md` et al.) already captures the chain semantics.

### Pass 2 — 2026-06-23
**Coverage delta (seeders package):** Statements 96.2% → 96.6% (+0.4). `users.go` Seed 91.9% → 97.3%.

**Tests added:**
- `TestUsersSeeder_HeroRidesPopulationIndex` — the D2 bridge: the hero's real name + org-domain email
  land at exactly her derived population index (the index the PersonaSeeder verifies against), and
  exactly one row carries the hero name (no hero-branch leak).
- `TestSeedVerifiedSkill_PerSessionScoreClamps` — score/comp stay in [0,100]/[1,100] across the shipped
  COPY rows.
- 2 users error-path (`users` COPY error → 0 rows + "copy users"; `memberships` COPY error → partial
  users-written count + "copy memberships").

**Bugs fixed inline:** none.

**Flakes stabilized:** none.

**Knowledge backfill:** recorded **D-M34-6** in `decisions.md` — two residual uncovered blocks
(`seedVerifiedSkill` score/comp clamps + the `users.go` casbin-grant error wrapper) are
unreachable-by-construction / already-tested-upstream defensive code; no compensating test owed.
(An audit trail so a future close/audit pass doesn't re-flag them.)

### Stop condition
Loop terminated after Pass 2: the Step 2b scan found nothing new worth adding (the two remaining
uncovered blocks are documented defensive dead-code, not behavioral gaps — testing them would be
shallow line-bumping, disallowed by the three-fate rule), the Pass-2 delta (+0.4%) is below the 2%
threshold, and zero flakes appeared across the run (3 clean sequential runs of the new tests in
Phase 4). Final: seeders **96.6%** statements, full `stack-seeding` suite `-race` green.
