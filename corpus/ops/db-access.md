# Production DB Access (read-only) — the snapshot/seeding read foundation

How to read the Anthropos production PostgreSQL safely, and the **public-vs-customer boundary** every read must
respect. This is the corpus anchor for the [`/db-query`](../../.claude/skills/db-query/SKILL.md) skill and the
**read foundation** the v1.2 snapshot capture builds on
([`snapshot-spec.md`](snapshot-spec.md), authored in M9a).

## For PMs — what it is

A read-only window into the live database, used to investigate data, debug, answer product questions, and — for
the demo/snapshot tooling — **size a data surface and tell public reference data apart from customer data**. Access
is **per-engineer, read-only, over the private network** (Tailscale); nobody queries with a write account. The
golden rule: **investigate freely, but never read so much that you slow the live product down, and never treat
customer-scoped rows as if they were shareable reference data.**

## For engineers

### Connecting — two paths

1. **The wired `postgres` MCP tool** (preferred in this workspace) — a read-only `mcp__postgres__query` tool
   already pointed at prod (a `<name>_read` account over Tailscale). Call it directly; verify with
   `SELECT current_database(), current_user, inet_server_addr();` (expect `postgres` / `<name>_read` / the RDS IP).
2. **Tailscale + `~/.pgpass` + `psql`** — `brew install libpq` (keg-only), Tailscale active (the RDS private IP is
   routed via a subnet router), and `~/.pgpass` holding `host:port:database:user:password`. Env vars `PGHOST`,
   `PGPORT`, `PGDATABASE`, `PGUSER=<name>_read`, `PGSSLMODE=require`. See [`/db-query`](../../.claude/skills/db-query/SKILL.md)
   for the full schema map + connection resolution.

### The two hard rules

1. **Read-only + low-impact.** SELECT only; schema-qualified; always `LIMIT`. For sizing/shape prefer
   **catalog-only** queries — `pg_class.reltuples`, `pg_total_relation_size(oid)`, `information_schema.columns` —
   which are instant and scan nothing. Avoid `COUNT(*)` / full scans on the GB tables (`public.skill_embeddings`,
   `public.skills`, `public.ai_usages`, `public.interactions/validation_*/activity_events` — the former `jobsimulation` tables). The snapshot
   **capture-source policy** ([`snapshot-spec.md`](snapshot-spec.md)) generalizes this with a source-pluggable
   precedence (M9a-D3): **ingest an existing prod `pg_dump` [default, zero new prod load]** → **safe throttled
   primary read [fallback]** (MVCC means a read-only `SELECT`/`COPY` never blocks writers — off-peak + chunked +
   bounded is tolerable) → **restore-from-snapshot / read replica [zero-primary-impact upgrades, once AWS/infra is
   wired]**. Whichever the source, **bound the session** (`SET TRANSACTION READ ONLY`, `statement_timeout`).
2. **The public ↔ customer boundary.** `organization_id IS NULL` = **global/public** reference data;
   `organization_id = <uuid>` = **customer-private**. Anything that *leaves* prod (a snapshot) must be **public
   only** — the snapshot tenant-data firewall (`AssertPublicOnly`) hard-fails on any captured row with a non-null
   org scope. Embeddings/translations carry no org column → scope them via the **public parent**.

### The public-vs-customer split (prod-verified 2026-06-06, catalog-grounded)

> *(Note: the monolith merges gave **every** application table a home in the `public` schema — same table names,
> same public/customer split. Taxonomy came in with skiller→app (July 2026); skill-path sessions with
> skillpath-in-app; the 23 simulation run-state tables with jobsim-in-app; the similarity + Studio tables with
> cms-in-app v8.0. The old `skiller`, `skillpath`, `jobsimulation` and `cms` schemas are legacy and no longer
> authoritative **for the platform**. Counts below are the 2026-06-06 verification — i.e. they were taken in the
> pre-merge schemas, so a row prefixed `public.` here is a re-label of a legacy-schema measurement, not a
> re-measurement. **Do not re-label a row whose number is still tied to the legacy schema** — the
> `cms.similarities` row below is exactly that case, and the note under the table says why.)*

| Surface | public (`org_id IS NULL`) | customer (`org_id` set) | snapshot rule |
|---|---|---|---|
| `public.skills` | 42,790 | 794 | capture public |
| `public.job_roles` | 22,315 | 2,381 | capture public |
| `public.specializations` | 1,442 | 154 | capture public |
| `public.categories` | 22 | 42 | capture public |
| `public.{skill,job_role}_embeddings` | — (no org col) | — | via public parent; rebuild index on replay |
| `public.studio_documents` *(was `cms.`)* | **0** | 3,060 | **exclude (all customer)** |
| `public.studio_tasks` *(was `cms.`)* | **0** | 2,353 | **exclude (all customer)** |
| **`cms.similarities`** *(see the attribution note)* | 274 | 733 | public only |

> ⚠️ **The 274 / 733 split was measured in the `cms` schema, and `cms` is still what the tooling captures — do
> not re-label this row `public.`** The measurement is the 2026-06-06 verification, taken before cms-in-app
> moved the table (app's `terraform/migrations/20260724132049_cms_data_model.sql` creates `similarities` in
> `public`, and `scripts/cms-data-sync/sync.sql:46`, `:53-55` copies `cms.similarities` → `public.similarities`;
> both @ app `ad9f3c498`). So **both tables exist**, and only the `cms` one was counted. The snapshot capture
> surface still reads the `cms` one: `stack-snapshot/simembeddings/simembeddings.go:44` @ rext `415240f` is
> `const Schema = "cms"`, and the surface's four tables (`similarities` + `similarity_categories` /
> `_features` / `_skills`, `:85-108`) are all captured from it. Attaching this number to `public.similarities`
> breaks the one link that makes it useful — it is the size of the surface the tooling *captures*, and that
> surface is `cms`. The other two former-`cms` rows above (`studio_documents`, `studio_tasks`) are **0** public
> either way, so their re-label is inert; this one is not.

The **public content template library** (global simulations/skill-paths) is **not** in any of the merged app
tables — it lives in the **`directus` schema inside the SAME `postgres` database** (served at
`content.anthropos.work`, but its rows are reachable read-only via the wired `postgres` MCP / `marco_read`, NOT a
separate Postgres — M10-D2 corrected the spike's "separate store" inference). That `directus` schema's public subset
(predicate `private=false AND tenant_id IS NULL AND status='published'`) is the v1.2 M10 content-snapshot source —
`directus.simulations` 2,597 total / 647 `private=false` / **304** strict-public-published; `directus.skill_paths`
263 / **22** strict.

### Sizing the snapshot surfaces (the catalog-only pattern)

```sql
-- Instant, zero table scan: size + approx rows for a schema's tables
SELECT c.relname, to_char(c.reltuples,'FM999,999,999') AS approx_rows,
       pg_size_pretty(pg_total_relation_size(c.oid)) AS total
FROM pg_class c JOIN pg_namespace n ON n.oid=c.relnamespace
WHERE n.nspname='public' AND c.relkind='r'
ORDER BY pg_total_relation_size(c.oid) DESC;
```

Prod headline (2026-06-06, measured pre-merge in the then-live `skiller` schema; the same tables now live in `public`): the taxonomy surface ≈ **2.1 GB** (the v1.2 taxonomy snapshot surface) — `skill_embeddings` 692 MB
(but heap only 3.3 MB → ~689 MB is the **pgvector index** → rebuild on replay, don't transport it), `skills`
436 MB, `job_roles` 362 MB, `job_role_embeddings` 339 MB, + translations. The former-`cms` content tables (now in `public`) are tens of MB.

## See also
- [`safety.md`](safety.md) — the tooling's consolidated read-side + write-side safety contract (this public-vs-customer
  boundary is the read-side foundation it builds on).
- [`/db-query`](../../.claude/skills/db-query/SKILL.md) — the full schema reference + cross-service relationships.
- [`snapshot-spec.md`](snapshot-spec.md) — the snapshot capture-source policy + tenant firewall (M9a).
- [`snapshot-cold-start.md`](snapshot-cold-start.md) — the cold-start runbook (M20): using this read foundation
  (a safe `--dsn`) to fill the snapshot cache once on a fresh box, and why the MCP is a query tool, not a capture source.
- [`seeding-spec.md`](seeding-spec.md) — the write-side production-isolation boundary (the read side is here).
- [`staging_from_dump.md`](staging_from_dump.md) — the full-clone (all-customer-data) precedent; the snapshot
  mechanism is its public-only, low-impact inverse.
