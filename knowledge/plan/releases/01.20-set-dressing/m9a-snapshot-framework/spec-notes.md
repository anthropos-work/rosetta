# M9a — Spec notes

Technical notes accumulate here during build. Prod evidence (2026-06-06):
[`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../../../../.agentspace/scratch/roadmap-research-2026-06-06.md).

## The `stack-snapshot/` extension (note #1)
_Dedicated rosetta-extensions section: capture + serialize + store + replay + the `stacksnap` CLI. `stack-seeding`
consumes it at the DAG `taxonomy/content (snapshot)` node._

## Snapshot contract + format
_Per-table COPY payloads + `manifest.json`; schema-version pinned; per-stack addressing._

## Production-safe capture-source policy (note #2)
_Precedence: cache-hit (no prod read) → read-replica refresh → restore-from-backup fallback → `--allow-primary`
last resort. Bounded read-only session (`SET TRANSACTION READ ONLY`, statement_timeout, chunked COPY streamed to
disk). Catalog-first dry-run sizes the read. Adds the read half the M7a guard lacks._

## Tenant-data firewall (note #3)
_`AssertPublicOnly`: every captured table has no org column OR is filtered `organization_id IS NULL`; hard-fail on
any captured tenant row. Public-only/provenance gene in the data-DNA. Prod-proven filters: taxonomy = `org_id IS
NULL`; embeddings/translations via public parent; app-Postgres `cms.studio_*` = 100% tenant → excluded._

## `.agentspace` manifest-cached store (note #4)
_Payload (gitignored) at the **workspace-level** `.agentspace/snapshots/<surface>/<schema-ver|hash>/…` (M9a-D5 —
one shared cache, captured once, replayed by every stack; not per-stack); `manifest.json` (source, schema ver, row
counts, public-only result, checksum, location, format ver). `SnapshotStore` interface — localfs now, cloud/S3
v1.3 (alt root considered: `~/.cache/rosetta-snapshots/`)._

Flow: `read replica → stacksnap capture (public-only) → .agentspace/snapshots/<surface>/<ver>/{manifest.json,
*.copy.gz} → stacksnap replay (per stack, bulk COPY + rebuild pgvector index)`. The manifest decides cache-hit
(replay, no prod read) vs stale (re-capture)._

## Data-DNA fidelity + public-only gene class
_source-vs-replay: row-count / structural conformance / referential integrity / embedding-dimension integrity +
public-only/provenance. `snapshot-seeded` surface status._

## `/db-query` port (note #5)
_`.claude/skills/db-query/SKILL.md` + `corpus/ops/db-access.md`; documents the `mcp__postgres__query` tool AND the
`~/.pgpass`+Tailscale+psql path. Schema map verified against live prod 2026-06-06._

## Tiny reference surface
_Proves capture→store→replay→fidelity-gate end-to-end, independent of taxonomy (M0 toy-mirror discipline)._
</content>
