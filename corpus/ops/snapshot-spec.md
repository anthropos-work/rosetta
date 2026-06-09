# Stack Snapshot — Spec

**The reference for `rosetta-extensions/stack-snapshot/`** — how a **public reference surface** (the taxonomy,
the global content library) is **captured once from production safely**, cached outside git, and **replayed into
any stack** — with a tested **tenant-data firewall** (never customer data) and a **measured fidelity** gate.

> **Scope.** This doc covers the **M9a framework** (the capture/serialize/replay contract + portable format, the
> production-safe capture-source policy, the tenant-data firewall, the `.agentspace` manifest-cached pluggable
> store, the `stacksnap` CLI, and the snapshot-fidelity data-DNA extension), the **M9b taxonomy surface** (the
> ~2.1 GB public skiller catalog), **and the M10 Directus content surface** — the public simulation / skill-path
> template library (the 9-table set under the per-surface `directus` public predicate, the per-stack Directus store
> fork, the media refs, the content fidelity gene, and the `sim_id`/`skill_path_id` linkage; see [The Directus
> content surface](#the-directus-content-surface-m10--the-second-real-surface)). With M10 the **last `waived` surface
> is promoted to `snapshot-seeded` → 100% data-DNA coverage** (the v1.2 thesis complete). **M11** curates this into
> the usable product layer — the refreshed presets + the set-dressed `corpus/ops/demo/` recipe family + the
> `/stack-snapshot` skill (the [set-dressing recipe](demo/recipe-snapshot-world.md)). **v1.3 M13** extends the
> mechanism from demo-only to **dev**: a `dev-stack up` bring-up now replays the cached surfaces + stands up a
> per-stack Directus + light-seeds itself by default — see [Dev as a full-fidelity
> peer](#dev-as-a-full-fidelity-peer-m13--local-directus--auto-snapshot--light-seed).
> The snapshot code lives in the gitignored `rosetta-extensions` monorepo (authored + tagged in the authoring copy
> at `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag) — **no platform repo is modified**, and
> snapshot **payloads never enter git**. The read foundation is [`db-access.md`](db-access.md); the write-side
> production-isolation boundary is [`seeding-spec.md`](seeding-spec.md). The cloud/S3 store + AI-generated content +
> shareability are **v1.4** (was v1.3).

## For PMs — what it does

A demo world needs more than an org with users — it needs the **library** behind the product: the ~60K-skill /
18K-role taxonomy and the global content templates. That library is **public reference data** (the same for every
customer), but it lives in production. The snapshot mechanism copies the **public** part of that library out of
prod **once**, in a way that **cannot slow the live product** and **cannot copy any customer's private data**,
stores it locally (never in git — it's gigabytes), and **stamps it into each demo/dev stack** on demand. The
result is a demo world that looks like the real product, built from real public data, with a **measured guarantee**
that what got stamped in faithfully matches what was captured.

## The architecture, in one paragraph

`stacksnap` connects DIRECTLY to a safe prod source (never the platform code), runs a **catalog-only dry-run** to
size the surface (scans nothing), opens a **bounded read-only session**, `COPY`s each table's **public subset**
out, runs the **firewall** (hard-fail on any tenant-scoped row), and writes a small `manifest.json` + per-table
`*.copy` payloads to the **workspace-level `.agentspace/snapshots/<surface>/<schema-version>/`** cache. Replay reads
that cache, **verifies every payload checksum before writing a row**, bulk-`COPY`s each table into a stack in
dependency order, and **rebuilds any pgvector index** that was deliberately not transported. `stack-seeding`
consumes a replayed snapshot at its existing DAG node (`… → taxonomy/content (snapshot) → activity`).

```
[safe prod source] → stacksnap capture (public-only, firewalled)
       → .agentspace/snapshots/<surface>/<schema-ver>/{manifest.json, *.copy}
       → stacksnap replay (per stack: verify → bulk COPY → rebuild pgvector index)
```

## Why capture is its own extension (decoupled from seeding)

Capture is a **privileged prod READ**; seeding is a set of **per-stack WRITES**. They have different blast radii and
different safety contracts, so snapshotting is a **dedicated `stack-snapshot/` section**, a sibling of
`stack-seeding` — not folded into it (M9a-D2). Capture runs **once per release**; replay runs for **every** stack
off the one shared cache. The decoupling makes capture reusable (staging, tests) and keeps the read/write isolation
boundaries distinct: `stack-seeding`'s `AssertClean` guards **writes**; this extension's `AssertPublicOnly` guards
**reads** — together they close both halves.

## The capture-source policy (M9a-D3 — the read-half safety)

A capture must never block the hot primary. The source is **pluggable** and tried in a fixed precedence:

| # | Source | When it applies | Prod impact |
|---|--------|------|-------------|
| 0 | **cache-hit** | the cached manifest's schema version matches the stack | **zero read** |
| 1 | **dump-ingest** *(default)* | a staging prod `pg_dump` exists → restore it into a throwaway Postgres, point `--dsn` at the restore | **zero new prod load** (the restore is the ingest) |
| 2 | **primary-read** *(fallback)* | only a read DSN is available | low — see below |
| 3 | **restore-from-snapshot** *(upgrade)* | once eu-west-1 AWS access is wired | zero (throwaway instance) |
| 4 | **read-replica** *(upgrade)* | once a terraform replica exists | zero (cleanest steady state) |

**Both live sources read over `--dsn`** — there is **no offline pg_dump-FILE reader**. A `pg_dump` is "ingested" by
**restoring it into Postgres and pointing `--dsn` at the restore** (Postgres bulk-load handles the restore well; a
schema-scoped `pg_dump -n skiller` is small to restore). `dump-ingest` and `primary-read` differ only in *what*
`--dsn` addresses — a restored dump vs the prod read endpoint — plus the manifest label + precedence. (A direct
offline file-reader was considered and **dropped**, M9b-D9: it adds no new capability — the produced snapshot is
identical — and no reliable speed gain; restore-then-`--dsn` + the safe primary read cover the need.)

**Why a safe primary read is tolerable (the MVCC correction).** PostgreSQL MVCC means a read-only `SELECT`/`COPY`
**never takes a lock that conflicts with writers** — the only cost is I/O + buffer-cache pressure. So an off-peak,
throttled, public-only, **catalog-sized-first** read of a few-hundred-MB surface is not a scary last resort; it is a
sanctioned fallback. The **bounded read-only session** caps the impact:

```sql
SET TRANSACTION READ ONLY;                          -- structurally unable to write
SET statement_timeout = 1800000;                    -- 30 min: a runaway COPY aborts
SET idle_in_transaction_session_timeout = 60000;    -- a stuck client never pins a snapshot
SET work_mem = '64MB';                              -- modest, no buffer-cache pressure
```

**Infra facts (prod-verified 2026-06-06).** Standalone **RDS PostgreSQL 15.12**, instance
`terraform-20240826114413395400000001`, region **eu-west-1**, terraform-managed. **No read replica today** (not
Aurora; 0 standbys / 0 walsenders / 0 replication slots; `rds.logical_replication=off`). **No local AWS creds**
(`~/.aws/credentials` is empty), so (3)/(4) cannot be driven from this machine — they activate automatically once
the `source` package's `Kind.Available()` flips true. The live default is therefore (1) dump-ingest or (2) safe
primary read. A definitive AWS check (run by someone with creds) lives in the M9a `decisions.md`.

This adds the **read half** the M7a isolation guard lacks — that guard classifies and gates **writes** only.

## The tenant-data firewall (note #3 — the load-bearing safety)

> The firewall + the capture-source policy together are the **read-side half** of the tooling's consolidated
> safety contract, [`safety.md`](safety.md) (the write-side half is in [`seeding-spec.md`](seeding-spec.md)).

`firewall.AssertPublicOnly` is the **read-side analog of seeding's `AssertClean`** — a *concept* enforced by two
real Go gates, `AssertPlan` (plan-time) and `AssertCaptured` (post-capture); see [`safety.md`](safety.md) §1.1.

**The public boundary is per-surface, not one fixed column (M10 generalization, M10-D1).** What "public" means
differs by surface: the **taxonomy** surface uses `organization_id IS NULL`; the **Directus content** surface uses
`private = false AND tenant_id IS NULL AND status = 'published'`. The firewall therefore takes a **`PublicPredicate`**
per surface — the scope column(s) that decide public-vs-customer, plus the SQL `WHERE` that selects the public subset.
The org-only predicate is the **default** (`firewall.DefaultPredicate`), so taxonomy + the reference surfaces are
unchanged; a new surface declares its own. A table is admissible iff **one** of:

- it has **none of the predicate's scope columns** (a pure-reference table — e.g. `skiller.categories` or
  `directus.resource`), captured whole; OR
- it **carries a scope column and is filtered to the public subset** (the predicate's filter — e.g.
  `skiller.skills`: `organization_id IS NULL`, 42,763 public vs 794 customer; or `directus.simulations`:
  `private=false AND tenant_id IS NULL AND status='published'`, 304 public-published of 2,597); OR
- it is **column-less but scoped via a public parent** (embeddings/translations/`sim_tasks` carry no scope column;
  they are public iff their parent is — judged under the surface's predicate). **Multi-level chains** (M10-D4): a
  child whose immediate parent is itself column-less (directus `task_checks → sim_tasks → simulations`) carries a
  `ParentScope.ParentFilter` that chases to the scope-bearing root in one subquery.

The firewall runs **twice, defense in depth**:

1. **PLAN time** (before any read): `AssertPlan` — every table policy must be admissible (a tenant-bearing table
   declares the public filter; a column-less table is pure-reference or public-via). A bad plan refuses **before a
   single byte flows**.
2. **POST-capture** (after the rows are in hand): `AssertCaptured` — a hard re-check that the captured set holds
   **zero** tenant rows. A single leaked row **aborts the capture; nothing is written to the store.**

Prod-proven filters: taxonomy = `organization_id IS NULL`; directus content = `private=false AND tenant_id IS NULL
AND status='published'`; embeddings/translations/content-children via the public parent; the app-Postgres
`cms.studio_*` tables are **100% customer** (`studio_documents`: 0 public / 3,060 customer) → **excluded entirely**.

**The Directus content source — self-resolved (M10-D2).** The spike inferred a "separate self-hosted Directus
Postgres" with no reachable DSN. That was **wrong**: the public content *template* library lives in a **`directus`
schema inside the SAME `postgres` database** the taxonomy capture already reads — reachable read-only via the wired
`postgres` MCP (`marco_read`). So the content surface captures over the **same `--dsn`** as taxonomy, just a different
schema (no new credential). Verified read-only 2026-06-06: `directus.simulations` = 2,597 total / 647 `private=false`
/ **304** strict-public-published; `directus.skill_paths` = 263 / **22** strict.

## The portable format (note #2 — the contract)

A snapshot is a small **`manifest.json`** head plus large, gitignored per-table **`*.copy`** payloads:

| Manifest field | Meaning |
|---|---|
| `format_version` | on-disk schema version; an unknown version is treated as stale |
| `surface` | logical surface name (`taxonomy`, `reference-toy`) |
| `source` | the capture-source kind that produced it (provenance) |
| `schema_version` | the platform schema digest the capture was taken against — **the staleness key** |
| `captured_at` | UTC capture timestamp |
| `tables[]` | per-table: schema, table, captured columns (COPY order), row count, public-only filter, `public_via`, payload file, **SHA-256**, `vector_columns` (index rebuilt on replay) |
| `public_only` | the firewall result — **must be `true`** or the manifest is never replayable |

Tables are listed in **dependency (replay) order** so a bulk COPY never violates an FK. The **schema version** is a
catalog-only digest (`md5` over `information_schema.columns` for the surface's schema) — instant, no table scan;
when the platform schema moves, the digest changes and the cached snapshot is **stale**.

## The `.agentspace` manifest-cached store (note #4 — pluggable, gitignored)

Payloads live under the **workspace-level `.agentspace/snapshots/<surface>/<schema-version>/`** (M9a-D5: **one
shared cache**, captured once + replayed by every stack — `stack-demo` / `stack-dev` / tests — **not per-stack**).
The cache is **gitignored** — GB blobs never enter any git repo. The `manifest.json` drives the **cache-hit vs
stale→refresh** decision (`store.Resolve`):

- **cache-hit** — a cached manifest exists and its `schema_version` matches the target stack → replay, **zero prod
  read**.
- **stale** — the schema moved (or the format version is unknown) → a refresh is required.
- **miss** — no snapshot for the surface → capture it first.

`store.SnapshotStore` is an **interface** with a `localfs` backend now; the **cloud/S3 backend is the named v1.4
swap** (moved from v1.3 with the rest of the cloud/S3/AI-content seeds) — the manifest already addresses payloads
by location, so a remote backend re-implements the same
`PutManifest` / `PutPayload` / `GetManifest` / `GetPayload` / `List` surface with no contract change.

## Embedding capture (M9a-Q3)

pgvector columns are captured **verbatim** (the vectors are in the payload), but the **index is NOT transported** —
for `skiller.skill_embeddings` the index is ~689 MB of a 692 MB total (heap is only ~3 MB). The dry-run flags any
table whose total size dwarfs its heap as **index-rebuild-on-replay**; replay runs `REINDEX` after loading the
table's rows. The **embedding-dimension integrity** fidelity gene then confirms the replayed vectors carry the
captured dimension.

## The `stacksnap` CLI

```bash
stacksnap capture --surface <name> [--source dump-ingest|primary-read] \
                  --dsn <DSN> [--store <root>] [--dry-run]
stacksnap replay  --surface <name> --stack <demo-N|dev-N> [--dsn <base>] \
                  [--schema-version <ver>] [--store <root>]
stacksnap status  [--store <root>]
```

- **`capture`** reads a public surface once **over `--dsn`** (a restored-dump Postgres for `dump-ingest`, the prod
  read endpoint for `primary-read`), firewalls it, and serializes it to the store. `--source` (or the default
  precedence) picks the kind; both kinds read over `--dsn` — there is no `--dump` file path (M9b-D9).
  **`--dry-run`** sizes the surface (catalog-only) and asserts the firewall plan **without reading data** — the
  cheap pre-flight before a real read.
- **`replay`** resolves cache-hit vs stale against the stack's live schema, then loads the cached snapshot via bulk
  COPY + rebuilds any pgvector index. The target is **any stack** — `--stack demo-N` *or* `--stack dev-N`; a dev
  stack is a first-class replay target (the dev set-dressing pass, M13, drives this for `dev-N` cache-first — see
  [Dev as a full-fidelity peer](#dev-as-a-full-fidelity-peer-m13--local-directus--auto-snapshot--light-seed)).
  **Re-run safe (v1.3b M17):** replay **clears every target table** (a per-stack-isolated `TRUNCATE`, child-first)
  before reloading, so a 2nd replay **REPLACES, never appends** — no duplicate-key abort, no silent double. The
  destructive op is fenced to a single-table TRUNCATE on the per-stack offset Postgres only; full contract in
  [`idempotency.md`](idempotency.md).
- **`status`** lists cached snapshots (surface, schema version, rows, source, capture time).

Exit codes: `0` ok · `1` firewall/capture/replay error (e.g. a tenant-data leak aborted capture) · `3` usage error.
The store root defaults to `<workspace>/.agentspace/snapshots` (overridable via `--store` or `STACKSNAP_STORE`).

## The fidelity gate (extends the data-DNA)

The snapshot dimension extends the M7b data-DNA harness (`stack-seeding/dna/`, the `datadna` CLI) — see
[`alignment_testing.md`](../architecture/alignment_testing.md). It adds:

- a **`snapshot-seeded`** surface status that **counts toward coverage** (unlike `waived`): a surface that M7c
  waived (taxonomy + content, the snapshot/shared-store hard line) becomes `snapshot-seeded` once a snapshot fills
  it — **lifting the two waived surfaces to real, measured coverage**, the v1.2 thesis;
- a **snapshot-fidelity gene class** (two-sided: captured source vs replayed stack) — **row-count parity**,
  **structural conformance**, **referential closure**, **embedding-dimension integrity**, and the **public-only /
  provenance** gene (the firewall's measured counterpart: zero tenant rows after replay).

## Proven end-to-end (M9a)

The `reference-toy` surface (four tables exercising every firewall branch + the vector path) proves
capture→store→replay→fidelity end-to-end, independent of any real platform table (the M0 toy-mirror discipline). The
Go tests are **hermetic** (the DB behind small `Capturer` / `Replayer` / `FidelityProbe` interfaces, tested against
fakes); the end-to-end test composes the *real* capture + store + replay packages through one in-memory DB and
asserts row-count parity, the vector rebuild, the stale-cache refresh path, and that **customer data never crosses
the firewall**. A live-run recipe (the DDL in `reference/reference.go`) stands the surface up in a throwaway schema.

## The taxonomy surface (M9b — the first REAL surface)

M9b proves the framework on the live **public skiller taxonomy** — the ~60K-skill / 18K-role library behind the
product (~2.1 GB, prod-measured ~98% public). The surface is enumerated in `stack-snapshot/taxonomy/` (one source of
truth shared by the CLI registry, the fidelity gene, and any live-run recipe): `stacksnap capture --surface taxonomy`.

### The 10 tables, in FK (replay) dependency order

A table is listed AFTER every table it references, so a bulk-COPY replay never violates an FK:

| # | Table | Capture scope | Public rows (2026-06-06) |
|---|-------|---------------|--------------------------|
| 1 | `skiller.categories` | `organization_id IS NULL` | 22 |
| 2 | `skiller.job_role_categories` | **pure-reference** (no org column) — captured whole | 22 |
| 3 | `skiller.specializations` | `organization_id IS NULL` | 1,442 |
| 4 | `skiller.skills` | `organization_id IS NULL` | 42,763 |
| 5 | `skiller.job_roles` | `organization_id IS NULL` | 22,315 |
| 6 | `skiller.skill_embeddings` | public-via `skills` — vector `small_embedding3` dim **1536** | 42,763 |
| 7 | `skiller.job_role_embeddings` | public-via `job_roles` — vector `small_embedding3` dim 1536 | 18,904 |
| 8 | `skiller.skill_translations` | public-via `skills` | 85,491 |
| 9 | `skiller.job_role_translations` | public-via `job_roles` | 43,550 |
| 10 | `skiller.job_role_skills` | public-via **BOTH** `job_roles` AND `skills` | 72,556 |

The FK graph: `skills.parent → specializations`, `specializations.parent → categories`,
`job_roles.category_id → job_role_categories` (a **separate** pure-reference parent — NOT `skiller.categories`),
the embeddings/translations → their parent, and `job_role_skills → {job_roles, skills}`. The public hierarchy is
referentially closed: 0 public skills with a customer/missing specialization parent, 0 public specs with a
customer/missing category, 0 public roles with a missing category.

### The parent-scope capture filter (the M9a-framework extension M9b adds)

The embedding/translation/link tables carry **no `organization_id` of their own** — they are public iff their
parent skill/role is public. M9a recorded the parent name (`PublicVia`) in the manifest but applied an **empty**
capture filter to column-less tables (it would have captured the whole table, including customer-parented rows).
M9b closes this with `TableSpec.ParentScopes` (the FK column + the public-bearing parent) so the capture applies a
real predicate:

```sql
-- one parent (skill_embeddings / skill_translations):
skill_id IN (SELECT id FROM skiller.skills WHERE organization_id IS NULL)
-- two parents, ANDed (job_role_skills — role AND skill must both be public):
job_role_id IN (SELECT id FROM skiller.job_roles WHERE organization_id IS NULL)
  AND skill_id IN (SELECT id FROM skiller.skills WHERE organization_id IS NULL)
```

The post-capture firewall probe (`AssertCaptured`) is parent-aware: for a column-less table it counts rows **within
the captured set** whose parent is a customer row (the capture filter ANDed with the inverse predicate) — 0 by
construction for a correct filter, a hard abort otherwise.

**Why `job_role_skills` needs both endpoints.** Prod has **3** rows where a public role links a customer skill.
Scoping by the role alone would capture those 3 links whose `skill_id` is absent from the public skill set → FK
orphans on replay. Both-endpoints scoping keeps the surface referentially closed (`72,559` public-role links − 3 =
`72,556`).

### The pgvector index rebuild on replay

`skill_embeddings` is 692 MB total but its **heap is only ~3 MB** — ~689 MB is the IVFFLAT index. The vectors are
carried verbatim in the payload; the **index is NOT transported**. Replay loads the rows then `REINDEX TABLE`s to
populate the index DDL the stack's migrated schema already defines (the platform migration owns the index params;
replay does not re-issue `CREATE INDEX`). The **embedding-dimension fidelity gene** then confirms the replayed
vectors carry the captured width (1536, read catalog-only via `atttypmod`).

### Captured columns exclude the generated `ts_search`

The taxonomy tables carry a `ts_search` tsvector that the stack regenerates; it is **not** in the captured column
set (it would be stale on replay). The captured set is the identity + descriptive + scope columns the COPY
serializes in order.

### The fidelity measure (`datadna measure-snapshot`)

The two-sided fidelity gene needs a **source** side (the captured manifest) and a **replay** side (the live stack).
M9b wires both: `dna.CapturedFromManifest` derives the per-table expectations from a real `manifest.json` (refusing
a non-public-only manifest), and `datadna measure-snapshot --stack demo-N --dna <dna> --manifest <taxonomy.json>`
runs the five fidelity operators (row-count / structural / referential / embedding-dim / public-only) against the
replayed stack, exiting non-zero if critical fidelity < 100%.

**Read-side public-only is asymmetric to the capture side** (pinned by the M9b hardening pass — `PgFidelityProbe`):
- A table **with an org column** (categories/specializations/skills/job_roles) is tenant-checked directly — the
  probe counts replayed rows whose `organization_id` is non-null, which must be 0.
- A **column-less** table (embeddings/translations/`job_role_skills`) reports **0 tenant rows by construction** and
  the probe short-circuits without a second query: it has no org column of its own to count, AND the replayed stack
  holds only the public snapshot — there is no customer parent in the stack for a row to reference. The capture-side
  parent-scope leak probe (above) is what actually enforces public-only for these tables; on the replay side there
  is simply nothing customer-scoped left to find. The embedding-dimension probe is likewise catalog-only
  (`atttypmod`): a non-vector column reads `NULL`/`-1` and is rejected as "not a fixed-dimension vector" rather than
  passing the gene as a 0-dim vector.

## The Directus content surface (M10 — the second REAL surface)

M10 captures the **public Directus content library** — the global simulation / skill-path templates behind the
product (the CMS + studio-desk integrate with Directus directly). The surface is enumerated in
`stack-snapshot/directus/` (one source of truth for the CLI registry, the fidelity gene, and the live recipe):
`stacksnap capture --surface directus --dsn <same DSN as taxonomy>`.

### The 9 tables, in FK (replay) dependency order

Directus has **no DB-level FKs** (relations live in the app layer), so closure is by convention — the parent-scope
filters keep the captured child set parent-closed. Public counts prod-verified read-only 2026-06-06:

| # | Table | Capture scope (under the directus predicate) | Public rows |
|---|-------|----------------------------------------------|-------------|
| 1 | `directus.simulations` | scope-bearing root: `private=false AND tenant_id IS NULL AND status='published'` | 304 |
| 2 | `directus.skill_paths` | scope-bearing root (same predicate) | 22 |
| 3 | `directus.resource` | **pure-reference** (global learning-resource library; no tenant column) | 1,543 |
| 4 | `directus.roles` | parent-scoped via `simulations` (col `simulations`) | 953 |
| 5 | `directus.sim_tasks` | parent-scoped via `simulations` (col `simulation`) | 949 |
| 6 | `directus.sequences` | parent-scoped via `simulations` (col `simulation`) | 304 |
| 7 | `directus.task_checks` | **multi-level** via `sim_tasks`→`simulations` (`ParentFilter`) | 2,242 |
| 8 | `directus.task_sub_checks` | **multi-level** via `task_checks`→`sim_tasks`→`simulations` | 2,850 |
| 9 | `directus.sequences_roles` | parent-scoped via **BOTH** `sequences` AND `roles` | 953 |

### The per-stack Directus store fork (M10-D2)

Booting a per-stack Directus needs its `directus_*` **system schema** (collections/fields/permissions DDL) AND the
content rows. The store fork is **bootstrap → replay → boot** (`stack-snapshot/directus/provision.go`):

1. **`directus bootstrap`** creates the `directus_*` system tables + the user-collection table STRUCTURE against the
   stack's empty Postgres `directus` schema. The snapshot **never** carries the `directus_*` system tables — Directus
   owns their version-specific DDL (capturing them would couple the snapshot to a Directus version).
2. **`stacksnap replay --surface directus --stack demo-N`** bulk-`COPY`s the captured content rows into the
   now-existing user-collection tables (the framework's generic `CopyIn(schema, table)` → the `directus` schema,
   class `postgres` = `PerStackIsolated` = always allowed; the shared prod Directus is never written).
3. **Boot the per-stack Directus** pointed at the stack's `directus` schema; CMS / studio-desk for THIS stack point
   `DIRECTUS_BASE_ADDR` at the offset-port container, **not** `content.anthropos.work`. `EnvContract.Validate()`
   hard-rejects any per-stack env that resolves to the prod Directus.

The live container boot is a **documented operational step** (the M9b discipline) — the build proves the contract +
plan hermetically; standing up a per-stack Directus in-build is the operator's recipe.

### Media / blobs (M10-D4) — refs are the floor, blob bytes are S3-gated

Content references media (`simulations.cover`, `skill_paths.{cover,image,video}`, `roles.avatar`,
`sequences.scenario_video`, `resource.{image,file}` — all uuid `directus_files` ids; `sequences.intro_video` is a
varchar embed, NOT a file ref). The **file refs** — the **1,311** `directus_files` rows the public content references
(of 10,340) — are the **floor**: always captured + replayed. The **blob bytes** live in Directus's OWN S3 bucket;
mirroring them needs **S3 read access to that bucket, not confirmed wired here** (`BlobBytesAvailable() == false`).
Until then the per-stack Directus serves refs with a local-storage adapter + placeholder assets (a believable
structural demo); blob-byte mirroring is the operational add (replay ONLY to the per-stack-isolated private bucket,
never the shared prod S3).

### The content fidelity gene + the `sim_id`/`skill_path_id` linkage

The content surface is the data-DNA `content` gene, **promoted `waived-m7c` → `snapshot-seeded-m10`** — the last
waived surface lifted to real coverage. It names `snapshot-row-count` / `-structural` / `-referential` /
`-public-only` (no `-embedding-dim` — content has no vectors). The **public-only gene is measured against the
directus predicate** (not `organization_id`): a `CapturedTable.PublicFilter` carries the per-surface predicate on the
scope-bearing tables, and `PgFidelityProbe.ReplayedNonPublicRows` counts replayed rows `WHERE NOT (<predicate>)` — 0
or it fails. `datadna measure-snapshot … --manifest <directus.json>` runs it.

The **`ContentSnapshotSeeder`** DAG node (sibling of `TaxonomySnapshotSeeder`) verifies the content snapshot was
replayed (counts public `directus.simulations`) and is the ordering anchor the **session/assignment seeders sit
behind**: with content present, the v1.1 seeders' `sim_id` / `skill_path_id` / `resource_id` — previously **free
`deterministicUUID` values with no FK** — resolve against the **real replayed public template ids** (the M10 linkage,
`stack-seeding/seeders/contentref.go`). When no content snapshot is replayed (a structural-only run), the resolver
falls back to the free values (graceful degradation; the snapshot is a prerequisite, not a hard requirement).

## Dev as a full-fidelity peer (M13 — local Directus + auto-snapshot + light seed)

Through v1.2 the snapshot mechanism was demo-facing: a replay set-dressed a **demo** stack (driven by the
skill that v1.3/M14 hard-renamed `/demo-snapshot` → `/stack-snapshot`).
Replay was always **dev-aware in the contract** (`stacksnap replay --stack <demo-N|dev-N>`; `pg.ParseStackN`
parses `dev-3 => 3`), but a **dev** bring-up did none of it — a fresh dev stack had no per-stack Directus (it
pointed at shared prod Directus) and no seeded data. **M13 makes a freshly-built dev stack a full-fidelity peer
of a demo stack**: the `dev-stack up` bring-up runs a **set-dressing pass** (`rosetta-extensions/dev-stack/dev-setdress.sh`),
**default-on**, that gives dev the same world demo gets.

### What the dev set-dressing pass does

After bring-up (and schema migration), `dev-setdress.sh <N>` runs three steps against `dev-N` (offset Postgres,
`5432 + N·10000`):

1. **The per-stack Directus (local Directus on dev).** It emits the M10 store-fork recipe (bootstrap → replay →
   boot) and **firewall-checks the per-stack Directus env** — the dev CMS points its `DIRECTUS_BASE_ADDR` at the
   **per-stack** offset-port Directus, **never** the shared `content.anthropos.work`. Both the recipe and the
   firewall come from one source of truth: the **`stack-snapshot/cmd/provision-plan`** runner, which makes the M10
   `directus.ProvisionPlan` / `EnvContract` / `Validate` contract *executable* (it was library-only through v1.2 —
   #M13-D2).
   `provision-plan --check-env --base-addr … --dsn …` exits non-zero — **hard-aborting the pass before any
   replay** — if the per-stack Directus env ever resolves to the prod Directus. The **live container boot remains a
   documented operational step** (the M9b/M10 discipline): the recipe is printed + the env validated; the operator
   boots the container.
2. **Cache-first auto-snapshot.** It replays the cached **public** surfaces (`taxonomy` then `directus`) into
   `dev-N` via `stacksnap replay` — **cache-first by construction** (replay resolves the cache and **never**
   captures; capture is a separate, privileged release-time prod read). A **stale/missing cache is a warning, not
   a failure**: the dev stack degrades to a structural-only world but **still seeds**; only a real replay/firewall
   error aborts.
3. **The dev-min light seed.** It applies the `dev-min` preset (~1 org + ~10 users) so the stack is **never
   empty** — see [`seeding-spec.md`](seeding-spec.md#the-shipped-presets-stack-seedingpresets).

### Escapes + safety

- **`--no-snapshot`** keeps the seed but skips the heavier per-stack Directus + replay (lean bring-up).
- **`--no-setdress`** skips the whole pass (a bare bring-up — the pre-M13 behaviour).
- **Non-fatal.** The pass is non-fatal on `dev-stack up`: a not-yet-migrated stack still comes **UP**; re-run
  `dev-setdress.sh <N>` by hand after migration (cache-first + idempotent). (default-on-but-non-fatal: #M13-D3)
- **The n=0-dev guard.** The pass **hard-refuses N=0** (the main `anthropos` dev stack) without `--force`, so an
  auto-set-dress can never touch the developer's primary stack (a second layer above `stackseed --reset`'s own
  n=0 refusal — see [`seeding-spec.md`](seeding-spec.md#the-cli)).
- **Prod-safety holds unchanged.** Capture is never run by the dev pass (replay only — a per-stack WRITE to the
  isolated offset-port Postgres + the per-stack Directus schema); the shared prod Directus / prod S3 are never
  written; media stays **refs-only** (blob bytes are v1.4). The read-side `AssertPublicOnly` firewall + the
  write-side isolation guard both still hold.

The net effect: **dev and demo are now the same world built two ways** — the same `stack-snapshot` +
`stack-seeding` machinery, the same per-stack Directus store fork, behind the same firewalls. M12 made dev a peer
for **N-allocation** (the unified registry); M13 makes it a peer for **data**.

## See also
- [`snapshot-cold-start.md`](snapshot-cold-start.md) — the **cold-start runbook** (v1.3b M20): filling the cache once per release on a fresh box (the sanctioned DSN-export / dump-restore path over `--dsn`), why the wired `postgres` MCP is **not** a capture source, and how it slots into the auto-set-dress bring-up.
- [`demo/README.md`](demo/README.md) — the **demo-env family index**: where the snapshot replay (`/stack-snapshot`) sits in the up→snapshot→seed→use→down flow.
- [`demo/recipe-snapshot-world.md`](demo/recipe-snapshot-world.md) — the **set-dressing recipe**: the operator walk-through of capture→replay→a real-catalog/real-content world (this spec is its source of truth).
- [`db-access.md`](db-access.md) — the read foundation + the public/customer boundary + the `/db-query` skill.
- [`seeding-spec.md`](seeding-spec.md) — the write-side production-isolation boundary + the DAG node that consumes a snapshot.
- [`staging_from_dump.md`](staging_from_dump.md) — the full-clone-with-customer-data precedent; the snapshot mechanism is its public-only, low-impact inverse.
- [`alignment_testing.md`](../architecture/alignment_testing.md) — the snapshot-fidelity + public-only gene class alongside the behavioural (v1.0) + structural data-DNA (v1.1 M7b) dimensions.
