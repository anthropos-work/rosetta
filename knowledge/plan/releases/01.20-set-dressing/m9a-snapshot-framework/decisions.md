# M9a — Decisions

_Implementation decisions with rationale. ID scheme: M9a-D1, M9a-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M9a-D1 | Split former M9 → **M9a (framework) + M9b (taxonomy surface)** | The 5 notes (dedicated extension + capture-safety + tenant firewall + `.agentspace` store + db-query port) + the fidelity-DNA make a full framework milestone; the GB taxonomy surface is its own. Mirrors M7a→M7c. | 2026-06-06 (user) |
| M9a-D2 | Snapshotting is a **dedicated `stack-snapshot/` extension**, not folded into `stack-seeding` | Capture = privileged prod **read**; seeding = per-stack **writes**. Decoupling makes capture reusable (staging/tests) and keeps the isolation boundaries distinct. (note #1) | 2026-06-06 (user) |
| M9a-D3 | Capture refresh source = **read replica** (auto), fallback **restore-from-backup**; primary only behind `--allow-primary` | No-impact prod read, fully automatic via rosetta-extensions (no manual dump). (note #2) | 2026-06-06 (user) |
| M9a-D4 | Snapshots live in a **`.agentspace` manifest-cached store** (gitignored payload), pluggable `SnapshotStore`; **cloud/S3 = v1.3** | No GB blobs in any git repo; manifest drives cache-hit vs refresh; cloud swap is a backend change. (note #4) | 2026-06-06 (user) |
| M9a-D5 | The canonical cache root is the **workspace-level `.agentspace/snapshots/`** (one shared cache, not per-stack) | Captured once, replayed by every stack (stack-demo/dev/tests); matches note #4 literally; gitignored + inspectable alongside the workspace. (Alt considered: `~/.cache/rosetta-snapshots/`.) | 2026-06-06 (user) |

## Open at design (to resolve during build)
- M9a-Q1: confirm a prod RDS read-replica endpoint is reachable over Tailscale (else restore-from-backup is the default refresh).
  - **Investigated 2026-06-06 (DB-side):** connected to the **primary** (`pg_is_in_recovery=false`, 10.2.22.13);
    **0 connected standbys** (`pg_stat_replication`), **0 walsender backends** (`pg_stat_activity`), **0
    replication slots**. The instance *supports* replicas (`wal_level=replica`, `max_wal_senders=20`) but **none
    exists/streams today**. Caveat: `marco_read` isn't `pg_monitor` (replica-row visibility *could* be filtered),
    but the 0-walsender corroboration triangulates the conclusion. AWS-side `describe-db-instances` **couldn't
    run** — CLI installed (v2.24.4) but no usable credentials (`~/.aws/credentials` has no `[default]` profile, no
    `~/.aws/config`, no `AWS_*` env). **Implication:** with no replica present, **restore-from-backup is the
    de-facto refresh path** unless a replica is provisioned. Re-run the definitive AWS check with creds:
    `aws rds describe-db-instances --region <eu-west-1?> --query 'DBInstances[].{id:DBInstanceIdentifier,replicaSource:ReadReplicaSourceDBInstanceIdentifier,replicas:ReadReplicaDBInstanceIdentifiers,endpoint:Endpoint.Address}' --output table`.
- M9a-Q2: the manifest schema + the cache-staleness rule (schema-version mismatch and/or checksum).
- M9a-Q3: embedding capture — vectors verbatim, **rebuild pgvector index on replay** (don't carry the ~689 MB index); confirm rebuild cost.
- M9a-Q4: the `SnapshotStore` interface shape so the v1.3 cloud swap is a clean backend change.
</content>
