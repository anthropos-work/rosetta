# M9a — Decisions

_Implementation decisions with rationale. ID scheme: M9a-D1, M9a-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M9a-D1 | Split former M9 → **M9a (framework) + M9b (taxonomy surface)** | The 5 notes (dedicated extension + capture-safety + tenant firewall + `.agentspace` store + db-query port) + the fidelity-DNA make a full framework milestone; the GB taxonomy surface is its own. Mirrors M7a→M7c. | 2026-06-06 (user) |
| M9a-D2 | Snapshotting is a **dedicated `stack-snapshot/` extension**, not folded into `stack-seeding` | Capture = privileged prod **read**; seeding = per-stack **writes**. Decoupling makes capture reusable (staging/tests) and keeps the isolation boundaries distinct. (note #1) | 2026-06-06 (user) |
| M9a-D3 | Capture-source precedence (source-pluggable): **cache-hit → (1) ingest an existing prod `pg_dump` [default] → (2) safe throttled primary read [fallback, MVCC = no write blocking] → (3) restore-from-snapshot / (4) read replica [zero-impact upgrades, once eu-west-1 AWS/infra is wired]** | Investigation (2026-06-06) found no replica + no local AWS creds; dump-ingest = zero new prod load (staging already produces dumps); a read-only bulk read doesn't block prod writes, so (2) always works. (note #2; supersedes the initial "replica-first" framing) | 2026-06-06 (user) |
| M9a-D4 | Snapshots live in a **`.agentspace` manifest-cached store** (gitignored payload), pluggable `SnapshotStore`; **cloud/S3 = v1.3** | No GB blobs in any git repo; manifest drives cache-hit vs refresh; cloud swap is a backend change. (note #4) | 2026-06-06 (user) |
| M9a-D5 | The canonical cache root is the **workspace-level `.agentspace/snapshots/`** (one shared cache, not per-stack) | Captured once, replayed by every stack (stack-demo/dev/tests); matches note #4 literally; gitignored + inspectable alongside the workspace. (Alt considered: `~/.cache/rosetta-snapshots/`.) | 2026-06-06 (user) |

## Resolved during build (2026-06-06)
| ID | Decision | Rationale |
|----|----------|-----------|
| M9a-D6 | The data-DNA snapshot extension lands in the EXISTING `stack-seeding/dna/` harness (`snapshot.go`), NOT a new harness | Fate-2: the M7b data-DNA already owns the gene/score/criticality machinery + the `datadna` CLI + the `waived` status these surfaces carry today; extending it (a `snapshot.go` file + a `snapshot-seeded` status + a snapshot-operator set) is the minimal, cohesive change. A snapshot gene names SNAPSHOT operators; a structural gene names STRUCTURAL operators; `Validate` rejects a cross-wire. `Coverage()` now counts seeded OR snapshot-seeded over the non-waived denominator — so a snapshot-filled formerly-waived surface reads as covered (the v1.2 thesis). |
| M9a-Q2→D | Manifest staleness key = a **catalog-only schema-version digest** (`md5` over `information_schema.columns` for the surface's schema) + the `format_version`. A mismatch (or unknown format) → stale → refresh | The digest is instant (no table scan), changes exactly when the schema moves, and partitions the cache by schema version (`<surface>/<schema-ver>/`) so a moved schema lands in a sibling dir without clobbering a still-valid older snapshot. Checksum corruption is a SEPARATE, replay-time check (needs the payload bytes) — `replay.Run` verifies every payload's SHA-256 before writing a single row. |
| M9a-Q3→D | pgvector: carry the **vectors verbatim** in the payload, **rebuild the index on replay** (`REINDEX TABLE`), do NOT transport the index | Confirmed via live catalog sizing: `skill_embeddings` is 692 MB total but only ~3 MB heap → ~689 MB is the index. The capture dry-run flags any table whose total dwarfs heap as index-rebuild-on-replay; the `snapshot-embedding-dim` fidelity gene confirms the replayed vectors carry the captured dimension. |
| M9a-Q4→D | `SnapshotStore` interface = `PutManifest`/`PutPayload`/`GetManifest`/`GetPayload`/`List` + a separate `Resolve(store, ref, targetSchemaVer)` cache-hit/refresh brain; manifest addresses payloads BY LOCATION (filename) | This makes the **v1.3 cloud/S3 swap** a clean backend re-implementation of the same 5 methods — the manifest already names payload files, so a remote backend stores/fetches them by the same key. localfs is the only M9a backend. |
| M9a-D7 | Each module re-declares the small shared helpers it needs (`ParseStackN`, the offset-DSN math, `QuoteIdent`) rather than importing `stack-seeding` | `stack-snapshot` is a SEPARATE Go module (own `go.mod`) — importing across the two would couple their version graphs. The grammar (`{prefix}-{N}` offset, `5432 + N*10000`) is the documented shared convention; re-declaring ~30 lines is cheaper than a cross-module dependency. |

## Adversarial review (close — 2026-06-06)
Scenarios considered at close (the *scenario* is recorded so future reviewers see what was weighed; each
was verified handled in the shipped code, no fix needed):

- **Partial-capture tenant leak.** Could a tenant row reach the store if a later table in the capture
  loop fails the firewall? No — `capture.Run` **stashes every payload in memory** and runs
  `firewall.AssertCaptured` over ALL tables BEFORE writing a single byte; any non-zero tenant-row count
  aborts with nothing persisted. (Defense-in-depth with the plan-time `AssertPlan`.) Verified by
  `capture/capture_harden_test.go` (store-write-error + tenant-probe-error paths).
- **Corrupt-cache half-replay.** Could a corrupt payload load a partial snapshot into a stack? No —
  `replay.Run` **verifies every payload's SHA-256 against the manifest before the first `CopyIn`**; a
  mismatch aborts before any row is written. Verified by `replay/replay_harden_test.go` (payload-read
  fault aborts before COPY).
- **Path traversal via surface/ref/payload names.** A hostile surface name or payload filename escaping
  the cache root? No — `store.sanitize` (Ref keys, with a `..`-collapse-to-fixpoint) + `safeFilename`
  (payloads, rejecting `/ \ ..`) keep all writes inside the root. Fuzzed (`FuzzSanitize`, 6.3M execs, no
  escape). Pinned observation: `sanitize("a...b") → "a-.b"` — still safe.
- **Replay into prod by a wrong DSN.** `replayCmd`'s `--dsn` defaults to `localhost:5432` and the
  port-offset math derives the per-stack port from the stack name; replay only ever writes to the
  resolved per-stack-isolated DSN (the M7a isolated class). Capture is structurally read-only
  (`SET TRANSACTION READ ONLY` in the bounded session). Both halves stay on their side of the
  read/write boundary by construction.
- **Schema-drift silent replay.** Could an old-schema snapshot replay into a moved-schema stack? No —
  `Resolve` keys the cache by schema-version digest and `IsStaleFor` treats a digest mismatch (or an
  unknown `target` version) as stale → refuse-and-refresh, not silent drift.

No adversarial scenario revealed release-scope-breaking work; all are Fate-1-handled in the shipped code.

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
