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

## Final review
_(filled at close)_
