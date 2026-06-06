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

## Final review
_(filled at close)_
</content>
