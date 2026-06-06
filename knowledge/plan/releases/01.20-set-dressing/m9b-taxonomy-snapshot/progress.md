# M9b — Progress

**Shape:** section · **Status:** built (ready for harden/close)

## Section checklist (from overview Scope.In)
- [x] Public taxonomy capture: `skiller.{categories,specializations,skills,job_roles}` filtered `org_id IS NULL` — plus `job_role_categories` (a separate pure-reference parent of job_roles, surfaced + landed Fate 1)
- [x] Parent-scoped capture: `skill_embeddings` / `job_role_embeddings` (vectors only) + `skill_translations` / `job_role_translations` + `job_role_skills` via the public-parent join — `TableSpec.ParentScopes` + `firewall.ParentScopeFilter` build a real predicate (the M9a empty-filter gap closed, M9b-D2); job_role_skills both-endpoints (M9b-D3)
- [x] Bulk-`COPY` replay per-stack (M7a perf path; per-stack-isolated skiller Postgres only) — single streamed COPY (M9b-D4)
- [x] pgvector index **rebuild on replay** (carry vectors verbatim, don't transport the ~689 MB index); embedding-dimension gene green — `REINDEX TABLE`, dim 1536 in manifest (M9b-D5)
- [x] Taxonomy fidelity + public-only genes wired; coverage `waived → taxonomy-seeded` — promoted `waived-m7c → snapshot-seeded-m9b`, all 5 operators; `CapturedFromManifest` + `MeasureSnapshot` + `PgFidelityProbe` + `datadna measure-snapshot` (M9b-D6); catalog 100%
- [x] Taxonomy snapshot wired into the `stack-seeding` DAG node — `TaxonomySnapshotSeeder` verification/ordering node; `activity` orders behind it (M9b-D7)
- [x] Delivers: extend `corpus/ops/snapshot-spec.md` (taxonomy path) + update `corpus/ops/seeding-spec.md` (taxonomy promoted) — + `alignment_testing.md` wired-to-real-surface note

## Build state at exit
- **Tests:** stack-snapshot 128→147, stack-seeding 164→181 funcs; both modules `-race -shuffle` green; gofmt + go vet clean.
- **Extensions commits** (clone `.agentspace/rosetta-extensions`, on `main`): `59c6a0d` (impl), `0404760` (review fix). Tag `stack-snapshot-m9b` to be set at close (per M9a pattern).
- **Rosetta commits** (`m9b/taxonomy-snapshot`): `10f7f6b` (docs + records).
- **PR-review finding (fixed):** parent-scope leak probe must AND the capture filter (was scanning the whole table → false abort). Fixed + regression-tested.

## M9b: Hardening

### Pass 1 — 2026-06-06 (stack-seeding: the fidelity measure)
**Scope manifest (M9b-touched, the two Go modules):** `stack-snapshot/{taxonomy,firewall,capture,reference,cmd/stacksnap}` + `stack-seeding/{dna,seeders,cmd/datadna}` — every package already had a co-located test; gaps were error/edge branches behind a `*pg.Conn` seam and the fidelity-measure wiring M9a left unconstructed.

**Coverage delta (milestone-touched files):**
- `dna`: 81.8% → 87.3% (+5.5) — `fidelity_probe.go` 0% → **100%**, `CapturedFromManifest` 90.5% → 100%.
- `cmd/datadna`: 46% → 52.6% (+6.6) — `printSnapshotScore` 0% → 100%, `measureSnapshotCmd` 66.7% → 69.4%.

**Tests added:** `fidelity_probe_test.go` (14: the concrete `PgFidelityProbe` via a fake `scanner` — column-less short-circuit, nil/negative-`atttypmod` guard, empty-CSV→nil, every error wrap, + the `PgFidelityProbe→RunSnapshotOperators` integration seam); `snapshot_harden_m9b_test.go` (4 unit + 1 **fuzz** `FuzzSplitCSV` — `CapturedFromManifest` empty-surface/no-tables/malformed-JSON/empty-vector-list); `cmd/datadna/main_harden_test.go` (4: manifest-file-missing + missing-DNA usage paths, `printSnapshotScore` PASS/FAIL rendering). Commit `eb20183`.

### Pass 2 — 2026-06-06 (stack-snapshot: the capture adapter's tenant-probe dispatch)
**Coverage delta:** `cmd/stacksnap`: 59.6% → 72.3% (+12.7).

**Refactor (behaviour-preserving):** extracted a narrow `captureConn` interface so the capture adapter's branch logic is unit-testable with a fake; `*pg.Conn` satisfies it verbatim (build + all prior tests green).

**Tests added:** `cmd/stacksnap/adapters_harden_test.go` (12 unit + 1 **fuzz**) — `CountTenantRows`' three-way dispatch (PARENT-SCOPED issues the leak probe; PURE-REFERENCE short-circuits via one `HasColumn`; ORG-BEARING runs the org-not-null count, incl. the empty-filter-but-has-org-column edge), `CopyPublic` row-count + build-before-IO + error wrap, `FuzzBuildParentLeakProbe` (adversarial identifiers → balanced quotes, the captured-set AND, the NOT-IN leak clause never dropped — 766K execs clean), the `captureAdapter→capture.BuildPlan` seam. Commit `016991c`.

### Pass 3 — 2026-06-06 (stack-snapshot: capture orchestrator error paths)
**Coverage delta:** `capture`: 91.0% → 95.5% (+4.5) — `BuildPlan` → 100%.

**Tests added:** `capture/capture_harden_test.go` (3) — `BuildPlan` SchemaVersion + SizeSchema resolution failures abort (the ref is keyed on the schema version), `Run` BeginBoundedSession failure aborts before any COPY and writes nothing. Extended `fakeCapturer` with three fail flags. Commit `f9fabc3`.

**Bugs fixed inline:** none — the build phase + its PR-review fix (the parent-leak-probe AND) had already closed the real defects; hardening surfaced no new bug. The fidelity-probe nil/negative-`atttypmod` guard and the `CountTenantRows` empty-filter-but-has-org-column edge were already correct; the new tests PIN them.

**Flakes stabilized:** none — both modules `-race -shuffle=on` green across all passes; two fuzz targets (`FuzzSplitCSV`, `FuzzBuildParentLeakProbe`) found no crashers (1.7M + 766K execs).

**Knowledge backfill:** see Phase 3b below — 1 invariant blended into `corpus/ops/snapshot-spec.md` (the read-side public-only guard for column-less tables: the probe ANDs the capture filter + a column-less replayed table reports 0 tenant rows by construction). No other KB-worthy findings.

**Test funcs:** stack-snapshot 147 → **165** (+18), stack-seeding 181 → **204** (+23).

### Stop condition
Stopped at Pass 3. The Phase-2b scan found no further pure-logic behaviour to add — the meaningful M9b surfaces (the fidelity probe, the manifest bridge, the parent-scope-leak dispatch, the capture error paths, the CLI rendering) are now 100% or near. All remaining uncovered statements are DB-pass-through plumbing (CLI `main`/`connect`/`*Cmd` dial bodies, `replayAdapter` concrete-`*pg.Conn` methods, the `pg` driver wrappers) behind a live-Postgres seam — covering them needs a running PG, not more unit tests. Next-pass delta on real logic would be < 2%.

## Final review
_(filled at close)_
</content>
