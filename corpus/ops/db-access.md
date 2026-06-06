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
   which are instant and scan nothing. Avoid `COUNT(*)` / full scans on the GB tables (`skiller.skill_embeddings`,
   `skiller.skills`, `public.ai_usages`, `jobsimulation.interactions/validation_*/activity_events`). The snapshot
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

| Surface | public (`org_id IS NULL`) | customer (`org_id` set) | snapshot rule |
|---|---|---|---|
| `skiller.skills` | 42,763 | 794 | capture public |
| `skiller.job_roles` | 22,315 | 2,381 | capture public |
| `skiller.specializations` | 1,442 | 154 | capture public |
| `skiller.categories` | 22 | 42 | capture public |
| `skiller.{skill,job_role}_embeddings` | — (no org col) | — | via public parent; rebuild index on replay |
| `cms.studio_documents` | **0** | 3,060 | **exclude (all customer)** |
| `cms.studio_tasks` | **0** | 2,353 | **exclude (all customer)** |
| `cms.similarities` | 274 | 733 | public only |

The **public content template library** (global simulations/skill-paths) is **not** in the app-Postgres `cms`
schema — it lives in the **separate self-hosted Directus store** (`content.anthropos.work`); that store's
public/global subset is the v1.2 M10 content-snapshot source.

### Sizing the snapshot surfaces (the catalog-only pattern)

```sql
-- Instant, zero table scan: size + approx rows for a schema's tables
SELECT c.relname, to_char(c.reltuples,'FM999,999,999') AS approx_rows,
       pg_size_pretty(pg_total_relation_size(c.oid)) AS total
FROM pg_class c JOIN pg_namespace n ON n.oid=c.relnamespace
WHERE n.nspname='skiller' AND c.relkind='r'
ORDER BY pg_total_relation_size(c.oid) DESC;
```

Prod headline (2026-06-06): `skiller` ≈ **2.1 GB** (the v1.2 taxonomy snapshot surface) — `skill_embeddings` 692 MB
(but heap only 3.3 MB → ~689 MB is the **pgvector index** → rebuild on replay, don't transport it), `skills`
436 MB, `job_roles` 362 MB, `job_role_embeddings` 339 MB, + translations. The `cms` content tables are tens of MB.

## See also
- [`/db-query`](../../.claude/skills/db-query/SKILL.md) — the full schema reference + cross-service relationships.
- [`snapshot-spec.md`](snapshot-spec.md) — the snapshot capture-source policy + tenant firewall (M9a).
- [`seeding-spec.md`](seeding-spec.md) — the write-side production-isolation boundary (the read side is here).
- [`staging_from_dump.md`](staging_from_dump.md) — the full-clone (all-customer-data) precedent; the snapshot
  mechanism is its public-only, low-impact inverse.
</content>
