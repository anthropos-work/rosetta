# M40 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.
Code lands in the rext authoring copy (`.agentspace/rosetta-extensions`) @ tag `method-acting-m40`;
doc-half in `corpus/ops/snapshot-spec.md`.

## Section checklist
- [x] **(a) `directus_versions` serve-grant** — DONE. Synthesized public-read **+ create** grants on the
  `directus_versions` SYSTEM collection (the full system name, NOT the `versions` API path — verified live: only
  `directus_versions` flips `/versions` from 403→200). `create` lets cms's `GetLatestOrCreateVersion` self-heal the
  empty per-stack version table instead of replaying prod's 539 version rows. Unblocks the entire skill-paths library +
  every detail page. Live: `publicSkillPaths = 22`.
- [x] **(b) library-category collections serve-grant** — DONE. Extended `servedCollections` with the library closure
  (`library_categories`, `library_macro_categories`, the 2 M2M junctions) + synthesized `directus_fields`/`relations`
  for the M2M expansion + a synthesized public-read grant (prod copies only its 5 public collections). Also added
  `resource` + `job_position` (the M2O targets `skill_paths.video` / `simulations.job_position` expand) so cms doesn't
  get an unmarshalable bare FK string. Live: `publicJobSimulations = 50`, no `ToDomain` panic.
- [x] **(c) `simulations.sequences` O2M nested-read serve-grant** — DONE. The O2M is grantable in tooling — it needed
  the `directus_relations` + `directus_fields` registration (the per-stack Directus had 0 of each), NOT a platform
  nil-guard. The library half AND the activity-feed half BOTH ship in tooling (the key-risk fork refuted). Live:
  `jobSimulation(simulationId)` returns `sequences[].scenarioIntro` — the feed federation path no longer panics.
- [x] **Regression test** — DONE. 9 new unit tests in `stack-snapshot/directus/serve_test.go` (closure both-endpoints,
  off-stack-alias drop, self-guarded idempotency, the synth/versions grants, the served-set closure, the six-part
  ordering). Live acceptance on demo-3: all three surfaces serve `> 0` anonymously via cms.

## Notes
- **Root cause was larger than the original a/b/c framing** (recorded M40-D2): the per-stack Directus had
  `directus_relations = 0` / `directus_fields = 0` — the O2M/M2M aliases were UNKNOWN to Directus, not "stripped under
  the public policy". Both library + activity-feed halves ship in tooling (M40-D3).
- **KPI "AI simulations completed" = 0** (overview Open question): its source `public.local_jobsimulation_sessions`
  has no CMS dependency, so it is a separate frontend/auth-context concern — out of scope for this serve-grant, to be
  re-verified during the M42e/M42m coverage sweeps. Not coupled to the feed fix (confirmed: the feed fix is a
  Directus-serve gap; the KPI reads jobsimulation directly).
- Supply-chain GREEN (go.mod/go.sum byte-identical); directus pkg `-race` clean; gofmt/vet clean; zero platform edits.

## M40: Hardening

Code lands in the rext authoring copy (`.agentspace/rosetta-extensions`) on the `directus` package;
the `method-acting-m40` tag was MOVED to the post-harden rext HEAD. The rosetta side carries only this
progress note. 3 passes, stopped on a clean Step 2b scan + flake-gate.

### Pass 1 — 2026-06-24
**Scope manifest (M40-touched code, rext `stack-snapshot/directus`):**
- `structure.go` — `CaptureServeRows` (88.9% baseline; the only sub-100% fn), `assertServeTablesAdmissible`,
  + the new M40 SQL templates/constants (`serveFieldsRowsSQL`, `serveRelationsRowsSQL`,
  `serveVersionsPermissionSQL`, `serveSynthesizedPermissionsSQL`, `servedCollections`, `serveVersionsCollection`).
- Existing tests: `serve_test.go` (9 M40 tests, all string-grep the SQL template text), `structure_harden_test.go`
  (pre-M40 `stagedRunner` — modelled only collections+permissions, did NOT model the 4 new M40 render stages).
- New harden file: `serve_harden_test.go`.

**Coverage delta (directus pkg, milestone-touched files):**
- Statements: 97.0% → 100.0% (+3.0); `CaptureServeRows` 88.9% → 100.0% (+11.1).

**Tests added:** `serve_harden_test.go` — `TestCaptureServeRows_PerStageRenderErrors` (4 sub-cases: fields /
relations / synth-permissions / versions-grant each surface a stage-named error, `errors.Is` the sentinel),
`_RenderOrderHaltsOnFirstError`, `_AllSixStagesRunInOrder`, `_SkipsEmptyM40Stages`, `_OnlyVersionsGrantPresent`.
Added the serve-aware `serveStagedRunner` (models all six render stages + the admissibility probe + per-stage
failure injection — the M40 peer of the pre-M40 `stagedRunner`).

**Bugs fixed inline:** none (Pass 1 surfaced no production bug — it closed the `CaptureServeRows` error-branch gap).

**Knowledge backfill:** no KB-worthy findings (the per-stage-error + render-order contract is internal test
robustness; the serve-grant behavior is already documented in `corpus/ops/snapshot-spec.md`).

### Pass 2 — 2026-06-24
**Coverage delta:** statements 100.0% → 100.0% (Pass 2 deepens SQL-semantics BEHAVIOR at the coverage ceiling —
Go statement coverage cannot measure SQL-template correctness; the grep tests it complements all "ran" already).

**Tests added:** `TestServeRelationsClosure_RequiresBothEndpoints` (asserts the two endpoint clauses are joined
by AND not OR — the both-endpoints closure — and exercises a Go mirror of the predicate over a fixture grid:
both-served / one-off-stack / NULL-junction / directus_-system / both-off-stack), `_ShareClosure` (rel_ok and the
relations render apply the IDENTICAL closure), `TestServeFieldsRelationalSpecialRegex` (compiles the actual
`(o2m|m2m|m2a|files|m2o|file)` alternation from the SQL; every relational special matches/gated, no non-relational
does — the `file`-substring-of-`files` concern), `TestServeGrants_ScopedToPublicPolicyOnly` (no UUID-shaped
over-grant to a non-public policy), `TestServeVersionsGrant_KeyedOnSystemCollectionName` (grant + guard key on
`directus_versions`, read/create dedup independently).

**Bugs fixed inline:** none in production code. Two TEST-fixture bugs surfaced and fixed inline within the pass:
(1) the closure-AND assertion was whitespace-strict — the SQL wraps `) AND (` across an indented newline, so the
assertion now collapses whitespace before checking; (2) the fixture's `served` map omitted the
`simulations_library_categories` junction, making a valid both-endpoints case fail — added the junction (it is
itself a served collection). Both confirm the production closure logic is correct as specified.

**Knowledge backfill:** no KB-worthy findings (the closure predicate's both-endpoints + relational-special-regex
semantics are already documented in the structure.go comments + `corpus/ops/snapshot-spec.md` M40 section).

### Pass 3 — 2026-06-24
**Coverage delta:** statements 100.0% → 100.0% (Dim 4 idempotency + Dim 1 framing depth at the ceiling).

**Tests added:** `TestCaptureServeRows_IsDeterministic` (two captures over the same source produce byte-identical
SQL — the render half of the idempotency contract: a re-replay emits the same guarded INSERTs → with the
WHERE-NOT-EXISTS guards, a true no-op re-apply), `TestCaptureStructure_FullServeFramingAndCount` (the FULL six-part
serve body inside CaptureStructure: one serve-header, serve-after-schema, all six parts present, statement count = 8
— the manifest provenance probe), `TestServeRenders_EveryGuardedRenderIsSelfGuarded` (the universal idempotency
pin: every no-natural-unique-key render — fields/relations/synth-perms/versions — carries a WHERE-NOT-EXISTS guard
keyed on its natural identity and never ON CONFLICT; the COPIED-permissions render correctly stays an unguarded
INSERT…VALUES so the two render styles can't drift).

**Bugs fixed inline:** none (all green first run).

**Knowledge backfill:** no KB-worthy findings.

**Tests-added totals:** 13 new test functions (8 unit/integration on the render machinery, 5 SQL-semantics/
behavioral) across `serve_harden_test.go`, plus the reusable `serveStagedRunner` harness. 0 production bugs;
2 test-fixture bugs fixed inline.

### Stop condition
3 passes. Stopped after the Pass 3 Step 2b scan found nothing new worth adding (the six dimensions are covered:
error paths = all 9 capture query stages; edge cases = empty/lone/NULL/off-stack; idempotency = determinism +
self-guard universality; fuzzing = the existing array-literal escape fuzz still covers the only user-input surface;
perf = N/A, a one-shot capture render with no SLA), coverage held at the 100%-statement ceiling (delta < 2% by
construction), and the flake gate passed (3 consecutive clean sequential runs of the 13 new tests). Well under the
5-pass cap. Supply-chain GREEN throughout (go.mod/go.sum byte-identical); `-race`/gofmt/vet clean; zero platform edits.
