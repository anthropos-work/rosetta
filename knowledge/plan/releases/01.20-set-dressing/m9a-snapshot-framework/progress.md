# M9a — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [ ] (note #1) The dedicated `rosetta-extensions/stack-snapshot/` section + the `stacksnap` CLI (capture / replay / status)
- [ ] Snapshot contract + portable format (per-table COPY payloads + `manifest.json`, schema-version pinned)
- [ ] (note #2) Production-safe capture-source policy (M9a-D3): cache-hit → (1) prod-`pg_dump` ingest [default] → (2) safe throttled primary read [MVCC, fallback] → (3) restore-from-snapshot / (4) read replica [upgrades]; bounded read-only session + catalog-first dry-run
- [ ] (note #3) Tenant-data firewall — `AssertPublicOnly` + the public-only/provenance data-DNA gene (hard-fail on any tenant row)
- [ ] (note #4) `.agentspace` manifest-cached store with a pluggable `SnapshotStore` backend (localfs now; cloud/S3 = v1.3); cache-hit vs stale→refresh
- [ ] Data-DNA extension: `snapshot-seeded` status + snapshot-fidelity gene class (incl. embedding-dimension integrity); datadna recognizes snapshot surfaces
- [ ] (note #5) Port `/db-query` skill (`.claude/skills/db-query/SKILL.md`) + `corpus/ops/db-access.md` (MCP-tool + pgpass/psql paths)
- [ ] Tiny reference surface proving capture→store→replay→fidelity-gate end-to-end (no real surface)
- [ ] Delivers: `corpus/ops/snapshot-spec.md` (new) + `corpus/ops/db-access.md` (new) + alignment_testing.md fidelity+public-only genes

## Final review
_(filled at close)_
</content>
