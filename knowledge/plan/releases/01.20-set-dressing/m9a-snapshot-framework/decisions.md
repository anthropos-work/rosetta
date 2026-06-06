# M9a — Decisions

_Implementation decisions with rationale. ID scheme: M9a-D1, M9a-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M9a-D1 | Split former M9 → **M9a (framework) + M9b (taxonomy surface)** | The 5 notes (dedicated extension + capture-safety + tenant firewall + `.agentspace` store + db-query port) + the fidelity-DNA make a full framework milestone; the GB taxonomy surface is its own. Mirrors M7a→M7c. | 2026-06-06 (user) |
| M9a-D2 | Snapshotting is a **dedicated `stack-snapshot/` extension**, not folded into `stack-seeding` | Capture = privileged prod **read**; seeding = per-stack **writes**. Decoupling makes capture reusable (staging/tests) and keeps the isolation boundaries distinct. (note #1) | 2026-06-06 (user) |
| M9a-D3 | Capture is **source-pluggable**; default = **prod-dump ingest** or **safe throttled primary read** (MVCC = no write blocking); **restore-from-snapshot / read replica** are zero-primary-impact upgrades once eu-west-1 AWS/infra access is wired | Investigation (2026-06-06) found no replica + no local AWS creds; a read-only bulk read doesn't block prod writes, so a non-replica path is viable now. (note #2; supersedes the initial "replica-first" framing) | 2026-06-06 |
| M9a-D4 | Snapshots live in a **`.agentspace` manifest-cached store** (gitignored payload), pluggable `SnapshotStore`; **cloud/S3 = v1.3** | No GB blobs in any git repo; manifest drives cache-hit vs refresh; cloud swap is a backend change. (note #4) | 2026-06-06 (user) |
| M9a-D5 | The canonical cache root is the **workspace-level `.agentspace/snapshots/`** (one shared cache, not per-stack) | Captured once, replayed by every stack (stack-demo/dev/tests); matches note #4 literally; gitignored + inspectable alongside the workspace. (Alt considered: `~/.cache/rosetta-snapshots/`.) | 2026-06-06 (user) |

## Open at design (to resolve during build)
- M9a-Q1: which capture source to use (read replica / restore-from-snapshot / safe primary read).
  - **Investigated 2026-06-06, double-checked (user push).** Infra facts (from the wired `postgres` MCP DSN in
    `~/.claude.json`): standalone **RDS PostgreSQL 15.12**, instance
    **`terraform-20240826114413395400000001`**, region **`eu-west-1`**, terraform-managed; the dev stack
    (`stack-dev/platform`) uses a **local Docker Postgres**, not prod (its `.env` carries no AWS/RDS creds — only a
    `DIRECTUS_TOKEN`, relevant to M10).
  - **No read replica exists today**, now rigorously confirmed: (a) **not Aurora** (`aurora_version()` absent — so
    `pg_stat_replication` is *not* blind to readers the way it is for Aurora; the streaming-replica check is valid
    here); (b) **0 standbys** / **0 walsender backends** / **0 replication slots**; (c) `rds.logical_replication=off`.
  - **No AWS access from this machine:** `~/.aws/credentials` is **0 bytes**, no `~/.aws/config`, no `AWS_*` env, no
    terraform repo cloned locally — so neither `describe-db-instances` nor a snapshot-restore can be driven from
    here. Those need whoever holds the eu-west-1 AWS/terraform access.
  - **Correction (over-conservative earlier):** a **read-only** bulk read does **not block prod writes** —
    PostgreSQL MVCC means a `SELECT`/`COPY` never takes a lock that conflicts with writers. The only cost is I/O +
    buffer-cache pressure. So reading the primary is *not* a scary last resort; off-peak + throttled/chunked +
    `statement_timeout` + public-only (data, not indexes) makes a once-per-release ~few-hundred-MB read tolerable.
  - **Viable capture sources, ranked (no replica + no local AWS needed for 1 & 3):**
    1. **Ingest an existing prod `pg_dump`** — the team already produces these for staging
       ([`staging_from_dump.md`](../../../../corpus/ops/staging_from_dump.md)); the snapshot tool reads the dump
       offline, filters public-only → **zero new prod load**.
    2. **Restore-from-snapshot** to a throwaway instance — zero primary impact; needs eu-west-1 AWS access.
    3. **Safe throttled primary read** via the existing `marco_read` access (MVCC = no write blocking), off-peak.
    4. **Provision a read replica** — cleanest steady-state; a one-time terraform change (`aws_db_instance` with
       `replicate_source_db`), then the replica path activates.
  - **Updated stance:** the `SnapshotStore`/capture layer stays source-pluggable; the **default source is (1)
    dump-ingest or (3) safe primary read** today, with (2)/(4) as the zero-primary-impact upgrades once AWS/infra
    access is wired. Definitive AWS check, run by someone with creds:
    `aws rds describe-db-instances --region eu-west-1 --query 'DBInstances[].{id:DBInstanceIdentifier,replicaSource:ReadReplicaSourceDBInstanceIdentifier,replicas:ReadReplicaDBInstanceIdentifiers,endpoint:Endpoint.Address}' --output table`.
- M9a-Q2: the manifest schema + the cache-staleness rule (schema-version mismatch and/or checksum).
- M9a-Q3: embedding capture — vectors verbatim, **rebuild pgvector index on replay** (don't carry the ~689 MB index); confirm rebuild cost.
- M9a-Q4: the `SnapshotStore` interface shape so the v1.3 cloud swap is a clean backend change.
</content>
