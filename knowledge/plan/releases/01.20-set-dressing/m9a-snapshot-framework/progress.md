# M9a — Progress

**Shape:** section · **Status:** built (awaiting close)

## Section checklist (from overview Scope.In)
- [x] (note #1) The dedicated `rosetta-extensions/stack-snapshot/` section + the `stacksnap` CLI (capture / replay / status) — 9 Go packages, tagged `stack-snapshot-m9a`
- [x] Snapshot contract + portable format (per-table COPY payloads + `manifest.json`, schema-version pinned + SHA-256 checksums) — `manifest` package
- [x] (note #2) Production-safe capture-source policy (M9a-D3): cache-hit → (1) prod-`pg_dump` ingest [default] → (2) safe throttled primary read [MVCC, fallback] → (3) restore-from-snapshot / (4) read replica [upgrades]; bounded read-only session + catalog-first dry-run — `source` + `pg` + `capture.BuildPlan`
- [x] (note #3) Tenant-data firewall — `AssertPublicOnly` (plan + post-capture, hard-fail on any tenant row) — `firewall` package; the public-only/provenance gene is in the data-DNA (below)
- [x] (note #4) `.agentspace` manifest-cached store with a pluggable `SnapshotStore` backend (localfs now; cloud/S3 = v1.3); cache-hit vs stale→refresh — `store` package
- [x] Data-DNA extension: `snapshot-seeded` status (counts toward coverage — the v1.2 thesis) + snapshot-fidelity gene class (row-count / structural / referential / embedding-dimension / public-only); `datadna` catalog recognizes snapshot surfaces — `stack-seeding/dna/snapshot.go`
- [x] (note #5) `/db-query` skill (`.claude/skills/db-query/SKILL.md`) + `corpus/ops/db-access.md` (MCP-tool + pgpass/psql paths) — inherited from the release branch; corrected to the M9a-D3 capture-source precedence in Phase 0b
- [x] Tiny reference surface proving capture→store→replay→fidelity-gate end-to-end (no real surface) — `reference` package (hermetic, composes the real packages)
- [x] Delivers: `corpus/ops/snapshot-spec.md` (new) + `corpus/ops/db-access.md` (inherited) + alignment_testing.md snapshot-fidelity + public-only genes

## Final review
_(filled at close)_
</content>
