# Recipe — Set-dress a stack with the real public library (snapshot capture → replay)

**Goal.** Turn a structural demo ("an org with users") into a **set-dressed** world ("an org browsing the real
product catalog") by stamping the real **public** reference library — the ~60K-skill / 18K-role taxonomy and the
global simulation / skill-path content templates — into the stack. This is the v1.2 "set dressing" layer: the
catalog view shows real skills, and seeded sessions/assignments link to real templates instead of placeholder ids.

**Time.** ~seconds on a **cache-hit** (the common case — the snapshot is captured once per release, replayed by
every stack); a few minutes on the rare **capture** (first time this release, or after a platform schema change).

**Source of truth.** [`../snapshot-spec.md`](../snapshot-spec.md) — the capture/replay contract, the tenant-data
firewall, the fidelity gate. This recipe is the operator walk-through.

## The two verbs (almost always you only need `replay`)

```
stacksnap capture …   →  read a PUBLIC surface ONCE from a safe prod source, firewall it, cache it   [maintenance]
stacksnap replay  …   →  stamp the cached snapshot into a stack (verify checksums → bulk COPY)        [per-stack]
stacksnap status      →  list cached snapshots (surface, schema version, rows, source, capture time)
```

`/stack-snapshot` drives all three; `replay` is the headline verb. Because the store is **cache-first**
(`store.Resolve`: a cached manifest whose `schema_version` matches the stack → **zero prod read**), a curator
almost always just replays an existing snapshot. Capture is the rare refresh op.

## A — replay into a stack (the common path)

**Prerequisite.** A stack up (`/demo-up N` **or** `/dev-up N` — dev is a peer; replay works on `dev-N|demo-N`
alike) and migrated (so the `skiller` + `directus` schemas exist as replay targets). Note a `/dev-up N` of a
non-primary dev stack already runs this replay by default; you only call `/stack-snapshot replay` explicitly to
re-run or to set-dress a `demo-N`. The snapshot is **stack-global** public reference data — replay it once per
stack, independent of which org you then `/stack-seed`.

```bash
/stack-snapshot replay 1                         # both surfaces (taxonomy + directus) — the usual call
# or one surface at a time, explicitly:
SN=stack-demo/rosetta-extensions/stack-snapshot
go build -o /tmp/stacksnap "$SN/cmd/stacksnap"
/tmp/stacksnap replay --surface taxonomy --stack demo-1
/tmp/stacksnap replay --surface directus --stack demo-1
```

Replay **resolves cache-hit vs stale** against the stack's live schema, **verifies every payload checksum before
writing a row**, bulk-`COPY`s each table in dependency order, and **rebuilds any pgvector index** that wasn't
transported (the ~689 MB `skill_embeddings` IVFFLAT index is rebuilt via `REINDEX`, not carried). If the cache is
**stale or missing**, replay tells you to capture first (path B).

**Then seed + log in.** With the library in place, seed an org and the seeded sessions link to the real templates:
```bash
/stack-seed 1 --preset mid-500
# log in per recipe-browser-login.md → the catalog + assigned content are real, not placeholder.
```

### The per-stack Directus boot (content surface)
The taxonomy replays straight into the stack's `skiller` Postgres schema and is immediately visible. The Directus
content surface needs its per-stack Directus **booted against the stack's own `directus` schema** (bootstrap →
replay → boot), pointed at the offset-port container — **never** `content.anthropos.work`. The
`EnvContract.Validate()` guard hard-rejects any per-stack env that resolves to the shared prod Directus. See
[`../snapshot-spec.md`](../snapshot-spec.md#the-per-stack-directus-store-fork-m10-d2) for the boot steps. Until S3
blob-byte mirroring is wired (v1.3), the per-stack Directus serves media **refs** with a local-storage adapter +
placeholder assets — a believable structural demo.

## B — capture (the rare maintenance op)

You only capture when `stacksnap status` shows **no cached snapshot** for a surface, or the cache is **stale** (the
platform schema moved — the catalog-only `schema_version` digest changed). Capture reads the public surface **once**
from a safe source and caches it; every stack then replays off that one cache.

```bash
# default precedence picks the source; or name it. Both read over --dsn (see the safety section).
/tmp/stacksnap capture --surface taxonomy --dsn "$SAFE_DSN" --dry-run   # size + assert the firewall plan, NO read
/tmp/stacksnap capture --surface taxonomy --dsn "$SAFE_DSN"             # the real capture
/tmp/stacksnap capture --surface directus --dsn "$SAFE_DSN"            # same DSN — directus is a schema in the same DB
```

`--dry-run` sizes the surface **catalog-only** and asserts the firewall **plan** without reading a single data row —
the cheap pre-flight before a real read. The capture serializes a small `manifest.json` + gitignored per-table
`*.copy` payloads under `.agentspace/snapshots/<surface>/<schema-version>/`.

## Safety — the load-bearing discipline (read this before a capture)

Capture is a **privileged prod READ**. Two guards, both hard-fail:

- **Public-only firewall (read-side `AssertPublicOnly`).** Capture transports **only** public reference rows —
  **never any customer/tenant data**. The boundary is **per-surface**: taxonomy uses `organization_id IS NULL`;
  Directus content uses `private = false AND tenant_id IS NULL AND status = 'published'`. The firewall runs **twice**
  — at PLAN time (before a byte flows) and POST-capture (a hard re-check that zero tenant rows were captured). A
  single leaked row **aborts the capture; nothing is written**.
- **Safe source, never the platform code.** Capture connects DIRECTLY to a safe source over `--dsn` (default: a
  restored staging `pg_dump`; sanctioned fallback: a **throttled, off-peak, read-only** primary read — Postgres
  MVCC means a read-only `COPY` takes no lock that blocks writers). A **bounded read-only session**
  (`SET TRANSACTION READ ONLY`, `statement_timeout`, `idle_in_transaction_session_timeout`) caps the impact. The
  capture **never** runs through the platform services and **never** writes anywhere.

Replay is **per-stack only** — it writes the per-stack-isolated `skiller` / `directus` Postgres (offset-port
container), class `PerStackIsolated` (always allowed); it can **never** write the shared prod Directus, the prod S3
bucket, or live Clerk. The read-side firewall (`AssertPublicOnly`) and the write-side isolation guard
(`AssertClean`, from seeding) together close both halves.

## Verify fidelity (captured source vs replayed stack)
The snapshot extends the data-DNA with a two-sided **fidelity** gene — row-count parity, structural conformance,
referential closure, embedding-dimension integrity, and the public-only/provenance gene (zero tenant rows after
replay). Gate it:
```bash
SS=stack-demo/rosetta-extensions/stack-seeding
SNAP=.agentspace/snapshots
/tmp/datadna measure-snapshot --stack demo-1 --dna "$SS/dna/data-dna.json" \
  --manifest "$SNAP/taxonomy/<ver>/manifest.json" --manifest "$SNAP/directus/<ver>/manifest.json"
# exits non-zero if critical fidelity < 100%.
```
With both surfaces replayed + gated, `datadna catalog` reads **100%** coverage — both formerly-`waived` surfaces
(`taxonomy`, `content`) are `snapshot-seeded`, nothing waived (the v1.2 thesis complete).

## Notes
- **Gigabytes, never in git.** Payloads live in the gitignored `.agentspace/snapshots/` cache (one shared cache,
  captured once + replayed by every stack). The cloud/S3 store is a named **v1.3** swap (same `SnapshotStore`
  interface, no contract change).
- **What's real vs not.** The taxonomy + content **libraries** are real (captured from prod public data). The
  per-session **AI narrative** (transcripts, AI scores, fresh embeddings) is **not** generated — that, plus
  external shareability, is the **v1.3** theme.
