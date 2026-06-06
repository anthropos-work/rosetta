# M10 — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [x] Per-stack content-store decision resolved + built (the defining fork — bootstrap→replay→boot per-stack Directus, M10-D2)
- [x] **Public** content capture from the **`directus` schema** (in-same-DB, read-only via `marco_read` — M10-D2 corrected the spike's "separate store"; public-published templates only; tenant-firewall generalized to the directus predicate; audited read-only)
- [x] Content replay seeder wired into M9a framework + the seeder DAG (`ContentSnapshotSeeder`), isolation-clean (per-stack directus schema)
- [x] Content fidelity + public-only genes in the data-DNA (4 ops, public-only measured against the directus predicate)
- [x] sim_id / skill_path_id / resource_id linkage → session/assignment refs resolve against real **public** templates (free-value fallback when no snapshot)
- [x] Coverage → 100% of the full catalog; content surface promoted waived→snapshot-seeded (nothing left waived)
- [x] Delivers: snapshot-spec.md (Directus path + store fork + self-resolved capture source) + seeding-spec.md (content surface update) + db-access.md (source-location fix) + alignment_testing.md (content surface)

## Build notes
- §1 firewall generalized to a per-surface PublicPredicate (the spike-flagged architectural gap) — taxonomy unchanged.
- §2 9-table directus surface (FK order, parent-scopes incl. multi-level chains via ParentScope.ParentFilter).
- §3 store fork (provision.go) + media refs (1,311 files); blob bytes S3-gated → MediaCaveat in PENDING.
- KB-1/KB-2 (Phase 0b findings) resolved in Phase 5 docs.

## M10: Hardening

### Pass 1–5 — 2026-06-06
5 passes (cap reached cleanly; stop = full Step-2b scan clean + coverage stable +
0 flakes). All deepening in the **extensions clone** (`stack-snapshot` +
`stack-seeding`); rosetta side docs-only (this record).

**Coverage delta (M10-touched code):**
- `stack-snapshot/firewall`: 95.9% → **100%** (statements)
- `stack-snapshot/directus`: 100% (held)
- `stack-snapshot/capture`: 93.8% → **98.8%** (residual = manifest-invalid-by-
  construction, unreachable from a validated plan)
- `stack-snapshot/cmd/stacksnap` — the M10 capture-adapter funcs
  (`CountTenantRows`, `buildParentLeakProbe`, `CopyPublic`, `buildPublicSelect`):
  80%/95% → **100%** (package total 76.6%→77.7%; the rest is pre-existing
  replay-adapter pass-throughs, not M10-touched)
- `stack-seeding/seeders` — all M10-touched files (`contentref`, `content_snapshot`,
  `jobsim_sessions`, `skillpath_sessions`, `assignments`): **100%** (package
  total 92.2%→95.2%)
- `stack-seeding/dna` — `fidelity_probe.go` `ReplayedNonPublicRows`: 0% → **100%**;
  `snapshot.go` referential dispatch closed (package total 86.2%→87.2%; the rest
  is pre-existing `introspect.go` DB plumbing, not M10-touched)

**Tests added (all in the extensions clone):**
- Pass 1 — firewall multi-level `ParentFilter` chain (the grandparent-exclusion
  SQL: a child whose grandparent sim is private is excluded), zero-predicate
  back-compat, malformed-predicate hard-fail at `AssertPlan`; linkage resolver
  error-path (missing directus schema → empty pool → free-value fallback).
- Pass 2 — directus per-surface-predicate surface E2E through the capture engine
  (`predicate()` non-default branch, manifest predicate label, multi-level leak
  abort); provision env-guard vs **disguised** prod targets (creds-embedded host,
  uppercase, prod-base+local-db); `sanitizeEnv` + Run BuildPlan/manifest-write
  error paths.
- Pass 3 — the directus public-only fidelity gene's **concrete probe**
  (`ReplayedNonPublicRows`, the firewall's measured replay-side counterpart) +
  the gene end-to-end through `PgFidelityProbe` + the referential dispatch.
- Pass 4 — the **CLI-adapter** multi-level leak probe (read-side
  grandparent-exclusion: `NOT IN (public-parent ids)` via `ParentFilter`) + a
  185K-exec multi-level fuzz (0 crashers).
- Pass 5 — linkage-seeder empty-seed/copy-error guards + `assignments.Isolation`
  contract.

**Bugs fixed inline:** none — the M10 build was sound; hardening surfaced no
production-code defects, only uncovered (correct) paths.

**Flakes stabilized:** none — flake gate **0/0** (3 consecutive sequential clean
runs per module; both modules pass `-race`; fuzz corpora 0 crashers).

**Knowledge backfill:** no KB-worthy findings — every invariant the new tests pin
(the per-surface predicate, the multi-level grandparent exclusion, the public-only
gene, the env-guard) is already documented in `snapshot-spec.md` / `seeding-spec.md`
/ `decisions.md` (M10-D1..D4). The hardening confirmed the docs, surfaced nothing new.

## Final review
_(filled at close)_
