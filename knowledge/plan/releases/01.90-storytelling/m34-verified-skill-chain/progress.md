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

## M34: Final Review (close)

Review found **5 findings**: 0 scope · 1 code-quality (should-fix: 1, nice-to-have: 2 → 1 fix, 2
recorded) · 2 docs · 0 tests · 0 net decision-triage. Deferral re-audit **GREEN** (0 deferrals).
Addressing all fully.

### Scope
- [x] No gaps. All 8 sections checked; 3 surfaced items fated (2× Fate-2 already-owned by M35/M36, 1×
  Fate-1 landed). Done-bar met (orchestrator-verified on live demo-3); literal browser-pixels of Maya's
  individual profile is M37/M38-owned (Fate-2, login-as-a-hero), not an M34 gap.

### Code Quality
- [x] [should-fix] `taxonomyref.go` `take()` comment claimed "the caller falls back to flat (then
  skips) when a role pool is short" — but the caller only falls back to flat on an EMPTY/unresolved role
  pool (via `skillsForRole`), never on a short-but-nonempty one. Corrected the comment to match the
  deliberate behavior: a short role pool yields fewer-but-role-coherent skills (more believable, closure
  stays green); padding-from-flat is a product choice routed to M35's roster work, not M34.
- [x] [nice-to-have → recorded, not changed] persona-index collision / `len(Personas) > Size` has no
  validation guard/warning. Documented-accepted for M34's 1-hero roster; recorded as **D-M34-7** as a
  guard to add when M35 graduates to the multi-hero roster (Fate-3 annotate, not a current-slice gap).
- [x] [nice-to-have → recorded] `splitCSV` hand-rolled splitter (pre-existing, cosmetic, in a touched
  file) — folded into D-M34-7 as a non-action note. Not worth a behavioral change.

### Documentation
- [x] `stack-seeding/README.md` Status reconciled: test count 62→**381** across **8** packages; Status
  line updated from "M7a" to reflect the M34 verified-skill-chain delivery; `seeders/` package-map row
  extended with the M34 seeders (PersonaSeeder / TaxonomyRefs / jobsim-sessions / users patch).
- [x] `dna/README.md` Status test counts reconciled: `dna` 49→**117**, `cmd/datadna` 10→**19**; note
  the M34 seed-side closure gene (`measure-closure`).

### Tests & Benchmarks
- [x] No gaps. 381 tests, `-race` clean, deterministic (3 identical runs), integration test
  double-gated (build tag + `STACKSEED_IT_DSN`), build-caught FK bug regression-guarded
  (`persona_test.go:181-196`). seeders 96.6% / `dna/seed_closure.go` 100%.

### Decision Triage
- [x] D-M34-1..6 → archive (maintainer-only implementation choices). The load-bearing insights (the
  7-table chain, the `user_level` reference omission, the closure gene, the empty-pool skip-don't-
  fabricate, the G14 valid-value class) already flowed into `stories-spec.md` during build; verified
  accurate. Added the `(#M34-D2/#M34-D3)` trace tags to `stories-spec.md` for the hero-index bridge +
  the evidences-UPSERT mechanisms.
- [x] D-M34-7 → new (the deferred multi-hero index-collision guard, Fate-3 annotate to M35).
