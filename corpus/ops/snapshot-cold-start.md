# Snapshot Cold-Start — getting the real catalog onto a fresh box

**The workflow for the one case the snapshot mechanism can't shortcut: a fresh box with an empty cache and no
safe `--dsn`.** `stacksnap replay` is almost always a **cache-hit** (zero prod read) — but the cache has to be
filled **once per release** by a `capture`, and a capture needs a **safe Postgres `--dsn`** to read over. On a
brand-new machine there is no cached snapshot *and* no obvious DSN. This doc is the sanctioned path from that
cold start to a real-catalog stack, and it states plainly what is **not** a shortcut (the wired `postgres` MCP).

> **Scope.** This is the operator runbook for ISSUE-13 (the cold-start gap the first real `/demo-up` surfaced).
> It builds on three docs and does not restate them: the **capture-source policy** + the firewall live in
> [`snapshot-spec.md`](snapshot-spec.md); the **read foundation** (Tailscale / `~/.pgpass` / the MCP, the
> public-vs-customer boundary) lives in [`db-access.md`](db-access.md); the **why-it's-safe** contract lives in
> [`safety.md`](safety.md). Everything here is the existing mechanism applied to a fresh box — no new tooling.

## For PMs — the one-paragraph version

A demo or dev stack shows the **real** product catalog (60K skills, the global content templates) by stamping in
a **snapshot** — a public-only copy of the catalog taken from production once per release and cached locally. On a
machine that already has that cache, stamping it in is instant and reads nothing from production. On a **fresh**
machine the cache is empty, so it has to be filled once — and filling it means reading the public catalog out of
a production-shaped database. This doc is the safe, confirmed way to do that one read. Until it's done, a stack
still comes up and is fully usable — it just shows an **empty catalog** (the seeder degrades gracefully); the
catalog goes real the moment the cache is filled and replayed.

## The decision in one picture

```
            ┌─ cache HIT (the steady state) ─────────────────────────────────────┐
stacksnap   │  .agentspace/snapshots/<surface>/<schema-ver>/ exists + matches     │
 replay  ───┤  → replay, ZERO prod read. THIS IS THE NORMAL CASE.                 │
            └────────────────────────────────────────────────────────────────────┘
                                  │ cache MISS / STALE (a fresh box, or schema moved)
                                  ▼
            ┌─ fill the cache ONCE per release (a privileged, confirmed prod READ) ┐
            │  pick a SAFE source for --dsn (precedence, snapshot-spec.md §source): │
            │   1. dump-ingest [default] — restore a staging pg_dump, --dsn at it   │
            │   2. primary-read [fallback] — Tailscale + ~/.pgpass read DSN          │
            │  stacksnap capture --surface taxonomy --dsn <safe>                    │
            │  stacksnap capture --surface directus --dsn <safe>                    │
            └────────────────────────────────────────────────────────────────────┘
                                  │ cache now warm
                                  ▼
                       replay is a cache-hit forever after (per box)
```

## What a fresh box does NOT have (the cold-start symptoms)

The first real `/demo-up` hit all of these at once (ISSUE-13):

- **An empty cache** — `stacksnap status` reports *"no snapshots cached under .agentspace/snapshots"*, so
  `replay` has nothing to stamp in.
- **No `~/.pgpass`** — no read DSN wired for `primary-read`.
- **No staging `pg_dump` on disk** — nothing to restore for `dump-ingest`.
- **The wired `postgres` MCP — which is NOT a capture source.** This is the load-bearing limitation, below.

The stack still comes up. `replay` exits non-zero on a cache miss, the bring-up treats that as a **warning, not a
failure** (the auto-set-dress pass is non-fatal — see [the auto-set-dress chain](#how-this-fits-the-auto-set-dress-bring-up)),
and the seeder degrades to an **empty catalog + free (unlinked) content refs**. So a cold box is never *blocked* —
it just isn't *set-dressed* until the cache is filled.

## Why the wired `postgres` MCP cannot fill the cache (the spike result, #M20-D4)

The obvious-looking shortcut — "we already have a read-only `postgres` MCP pointed at prod, capture through that" —
**does not work, and building an adapter for it is not worth it.** The spike (M20) settled this:

- `stacksnap capture`'s load-bearing primitive is `Capturer.CopyPublic`, which streams a table's public subset out
  via **`COPY (<select>) TO STDOUT`** and returns the **raw Postgres COPY text-format bytes**. Those exact bytes are
  what the manifest's per-table **SHA-256** checksums and the **snapshot-fidelity gate** verify on replay.
- The wired `postgres` MCP (`mcp__postgres__query`) is a **JSON-row query tool** — it returns parsed result rows,
  **not** the COPY wire format, and it cannot run `COPY … TO STDOUT` to a stream.
- An MCP-paging capture adapter would therefore have to **re-serialize JSON rows back into byte-identical COPY text
  format** — the `\N` NULL sentinel, tab/newline/backslash escaping, and per-type formatting (pgvector literals,
  timestamps, the deliberately-excluded `ts_search` tsvector) — *and* page the ~2.1 GB taxonomy + embeddings
  through a query interface. That is fragile, byte-fidelity-critical work for **zero capability gain**: the snapshot
  it produced would be identical to the `--dsn` path (the same reasoning that dropped the offline `pg_dump`-FILE
  reader, M9b-D9).

**So: the MCP is great for *investigating* and *sizing* the surface (catalog-only counts, the public/customer
split — that's exactly what [`db-access.md`](db-access.md) uses it for), but it is not a `stacksnap` capture
source. Capture reads over a real `--dsn`.**

## The sanctioned cold-start path — fill the cache once, over a safe `--dsn`

Pick the **first applicable** source from the capture-source precedence (full table in
[`snapshot-spec.md`](snapshot-spec.md#the-capture-source-policy-m9a-d3--the-read-half-safety)). On a typical box
that means **(1) dump-ingest** if a staging dump is available, otherwise **(2) primary-read** over Tailscale.

> **This is a privileged, prod-touching READ. Confirm before running it.** It is **operator-initiated and
> separate from any bring-up** — `/demo-up` / `/dev-up` never run a capture (they only ever *replay*; see below).
> Every capture session is bounded read-only (`SET TRANSACTION READ ONLY` + `statement_timeout` +
> `idle_in_transaction_session_timeout`) — it is structurally unable to write and cannot run away
> ([`safety.md`](safety.md) §1.4).

### Option 1 — `dump-ingest` (default; zero new prod load)

The team already produces staging `pg_dump`s. Restore one into a **throwaway** local Postgres and point `--dsn` at
the restore — the restore *is* the ingest, so production sees **no new read**:

```bash
# 1. Restore a (public-catalog-bearing) staging dump into a throwaway Postgres.
#    A schema-scoped dump is small to restore: `pg_dump -n skiller -n directus …`.
docker run -d --name snapsrc -e POSTGRES_PASSWORD=x -p 55432:5432 postgres:15
pg_restore -d "postgres://postgres:x@localhost:55432/postgres" path/to/staging.dump

# 2. Capture each public surface OVER --dsn pointed at the restore (dump-ingest is the default source).
SRC="postgres://postgres:x@localhost:55432/postgres"
stacksnap capture --surface taxonomy --dsn "$SRC"
stacksnap capture --surface directus --dsn "$SRC"

# 3. Throwaway gone; the cache is now warm.
docker rm -f snapsrc
```

### Option 2 — `primary-read` (fallback; safe throttled prod read)

If there's no staging dump, read the public surface **directly from the primary** over a read-only DSN. This is
**tolerable** because PostgreSQL MVCC means a read-only `SELECT`/`COPY` never takes a lock that blocks writers — an
off-peak, bounded, public-only, catalog-sized read is a sanctioned fallback, not a last resort
([`db-access.md`](db-access.md) rule 1, [`safety.md`](safety.md) §1.4):

```bash
# Wire the read foundation first (db-access.md §Connecting): Tailscale up, ~/.pgpass with a `<name>_read`
# account, PGSSLMODE=require. Then point --dsn at the prod read endpoint and force the fallback source:
READ_DSN="postgres://<name>_read@<rds-private-ip>:5432/postgres?sslmode=require"
stacksnap capture --surface taxonomy --source primary-read --dsn "$READ_DSN"
stacksnap capture --surface directus --source primary-read --dsn "$READ_DSN"
```

> Both options read **public-only** — the tenant-data firewall (`AssertPublicOnly`) hard-fails the capture on a
> single customer-scoped row, before anything is written to the cache. Customer data **cannot** leave prod through
> this path ([`safety.md`](safety.md) §1.1). The two zero-primary-impact upgrades (`restore-from-snapshot`,
> `read-replica`) activate automatically once eu-west-1 AWS/infra access is wired — they need no doc change here.

### Then: replay is a cache-hit forever after

Once the cache is warm, every stack on this box stamps the real catalog in with **zero** prod read:

```bash
stacksnap status                       # the two surfaces now appear (rows, schema version, source, time)
stacksnap replay --surface taxonomy  --stack demo-1     # or /stack-snapshot N
stacksnap replay --surface directus  --stack demo-1
```

### The Directus replay needs a per-stack Directus (executed on `--local-content`; v1.5)

The `directus` surface is content rows that need a **per-stack Directus on an offset port**, stood up
**bootstrap → apply captured structure → replay rows → boot**. On a **`--local-content` stack** (demo default-on,
dev opt-in) this recipe is **EXECUTED** by the bring-up (v1.5 M22): Directus's `node cli.js bootstrap` creates the
`directus_*` system tables into the stack's `directus` schema [`CREATE SCHEMA` + `DB_SEARCH_PATH=directus`]; then
`stacksnap` **applies the captured user-collection STRUCTURE** (the DDL + primary keys + the `directus_collections`
serve-row registration + public-read permissions — captured atomically with the rows, v1.5 M21) so the schema digest
converges and the cached rows replay in (**`stacksnap` exit 0**); then the container boots pointed at the stack's own
`directus` schema, and `cms` is cut over to it (M23) — never the shared `content.anthropos.work` for the data plane
(the asset plane stays on prod so images stay real). **Without `--local-content`** (or on an `N=0` dev stack) no
per-stack Directus is stood up: the stack reads content **live from prod** and the directus replay is a no-op with
**`stacksnap` exit 4** — the documented prod-read fallback, **not** a cache problem this runbook's capture can fix.
The cold-start case this runbook covers is the prerequisite for the executed path: the cache must be
**structure-bearing** (a capture made with the M21+ tooling) — a rows-only cache makes the local Directus boot
healthy but serve **0** content (the M25 field-bake's headline finding). Full mechanism:
[`snapshot-spec.md`](snapshot-spec.md#the-per-stack-directus-store-fork-m10-d2-recipe-corrected-in-fix16).

## How this fits the auto-set-dress bring-up

Both `/dev-up` (M13) and `/demo-up` (M20) auto-set-dress by default: after migrate they run a **cache-first
replay → light seed** pass. That pass **only ever replays** — it is **cache-first by construction and never
captures** (capture is the separate, privileged, confirmed prod read documented above). So on a cold box the
auto-set-dress pass:

1. tries the replay → **cache miss** → warns (non-fatal), and
2. **still seeds** (the atomicity floor — a stack is never catalog-only-and-403), so the stack comes up usable
   with an empty catalog.

To make the catalog real, fill the cache once (this doc), then re-run the set-dress pass (it's idempotent +
cache-first, and the M17 re-run guards make a retry safe — [`idempotency.md`](idempotency.md)). The catalog goes
real with no teardown.

## See also
- [`snapshot-spec.md`](snapshot-spec.md) — the capture/replay mechanism, the capture-source precedence, the
  tenant-data firewall, the per-stack Directus store fork. (The source of truth this runbook applies.)
- [`db-access.md`](db-access.md) — the read foundation: Tailscale / `~/.pgpass` / the wired `postgres` MCP, and
  the public-vs-customer boundary every read respects.
- [`safety.md`](safety.md) — why the capture read is safe (bounded read-only, public-only firewall) and why a
  bring-up's replay never touches prod.
- [`demo/README.md`](demo/README.md) — the demo-env family flow this slots into.
