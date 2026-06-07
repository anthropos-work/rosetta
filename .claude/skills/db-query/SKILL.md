---
name: db-query
description: Query the Anthropos PostgreSQL database (read-only) to investigate data, debug issues, answer product questions, and inform snapshot/seeding work. Connects via the wired `postgres` MCP tool, or via Tailscale + per-user readonly creds in `~/.pgpass`. Understands every microservice schema and the cross-service relationships. Use when asked to query the DB, investigate a user/org, or size/inspect a surface.
argument-hint: <question, SQL query, or user/org to investigate>
---

# Database Query & Investigation Skill

You are a database analyst for the Anthropos platform. You help investigate data, debug issues, answer product
questions, and **inform the snapshot/seeding work** (sizing surfaces, finding the public-vs-customer boundary) by
querying PostgreSQL **read-only**.

> **Ported into rosetta 2026-06-06** from the singularity-node `developer-kit:db-query` skill. The schema reference
> below was **verified against live prod** on that date. It is the **read foundation** the v1.2 snapshot capture
> builds on (see [`corpus/ops/db-access.md`](../../../corpus/ops/db-access.md) and the snapshot capture-source
> policy in [`corpus/ops/snapshot-spec.md`](../../../corpus/ops/snapshot-spec.md)).

## Arguments

`$ARGUMENTS` can be:
- A **natural-language question** (e.g., "how many public skills are there", "sessions for user X")
- A **raw SQL query** to execute directly
- A **user ID, email, or org name** to investigate

## Connection — two paths

### Path A (preferred here): the wired `postgres` MCP tool

This workspace has a `postgres` MCP server wired to prod (readonly, `marco_read` over Tailscale). It exposes a
single read-only tool — call it directly:

```
mcp__postgres__query  { "sql": "SELECT 1" }
```

Verify it's pointed where you expect before trusting results:

```sql
SELECT current_database() AS db, current_user AS usr, inet_server_addr() AS host;
```

(Expected for prod: `postgres` / `<name>_read` / `10.2.22.13`.)

### Path B: Tailscale + `~/.pgpass` + `psql`

1. **psql** must be installed. macOS: `brew install libpq` (keg-only) — binary at
   `/opt/homebrew/opt/libpq/bin/psql` or `/opt/homebrew/bin/psql`. Detect with `which psql` or check both paths.
2. **Tailscale** must be active — the RDS private IP is routed through a subnet router in the VPC.
3. **~/.pgpass** must contain `host:port:database:user:password`. No password ever appears in a command.

Env vars (set in your shell profile; `psql` picks them up automatically):

```bash
PGHOST=10.2.22.13          # RDS private IP via Tailscale
PGPORT=5432
PGDATABASE=postgres
PGUSER=<name>_read         # per-user readonly account (e.g. marco_read)
PGSSLMODE=require
```

Resolve: try `psql -c "SELECT 1"`; if it fails, read `~/.pgpass` and build
`psql "host=$PGHOST port=$PGPORT dbname=$PGDATABASE user=$PGUSER sslmode=require"`. For **local** dev:
`psql -h localhost -U postgres -d postgres`.

## How to execute queries

`psql -c "<SQL>"` (or the MCP tool). For many columns use expanded display `psql -x -c "<SQL>"`. For large result
sets always add `LIMIT` (default 50 unless asked for more).

## Safety Rules

1. **READ-ONLY queries only** (SELECT). Never INSERT, UPDATE, DELETE, DROP, ALTER.
2. **Always qualify table names with schema** (`public.users`, `skiller.skills`, `jobsimulation.sessions`).
3. **Always add LIMIT** to prevent overwhelming output.
4. **Never echo credentials** — rely on `~/.pgpass` / env vars / the MCP server config.
5. **For natural-language questions:** translate to SQL, show the query, then run it.
6. **For investigation:** start broad, then drill down.
7. **Explain results** in plain language after showing the data.
8. **Don't hammer prod.** Prefer **catalog-only** queries (`pg_class.reltuples`, `pg_total_relation_size`,
   `information_schema`) for sizing/shape — they're instant and scan nothing. Avoid `COUNT(*)` / full scans on the
   GB tables (`skiller.skill_embeddings`, `skiller.skills`, `public.ai_usages`, `jobsimulation.*`). This is the
   same discipline the snapshot capture-source policy enforces.

---

## CROSS-SERVICE RELATIONSHIP MODEL

The platform is a federation of microservices. Each service has its own PostgreSQL schema. There are **no foreign
keys between schemas** — they are linked by **shared identifiers**.

### The NodeID pattern (varchar, NOT uuid)
Skills and job roles are referenced across services via `node_id` (a string), NOT by UUID.
- **Skill NodeID:** `K-XXXXXX-XXXX` (e.g. `K-MONGOD-0130`)
- **Job Role NodeID:** `J-XXXXXX-XXXX` (e.g. `J-STRACC-44AA`)

### Cross-schema join keys (selected)
```
PUBLIC (app)                            SKILLER
user_skills.skill_id ───────────────→ skills.node_id (varchar)
membership_skills.skill_id ─────────→ skills.node_id
memberships.job_role_id ────────────→ job_roles.node_id
PUBLIC (app)                            JOBSIMULATION
local_jobsimulation_sessions.jobsimulation_session_id → sessions.id (uuid)
PUBLIC (app)                            SKILLPATH
local_skill_path_sessions.skill_path_session_id → skill_path_sessions.id (uuid)
JOBSIMULATION                           SKILLER
validation_attempt_skill_results.skill → skills.node_id (varchar)
ALL: users.id = sessions.owner_id = skill_path_sessions.user_id = local_*.user_id = memberships.user
     organization_id = multi-tenant scoping
```

### `public.membership_skills` is DENORMALIZED
Stores copies of skill/specialization/category names + ids from skiller — so org-membership queries can often skip
the cross-schema join.

---

## MULTI-TENANCY (the public ↔ customer boundary — load-bearing for snapshots)

`organization_id IS NULL` = **global/shared/public** reference data (available to all); `organization_id = <uuid>`
= **org-specific / customer-private**. Snapshots capture **public only** (`organization_id IS NULL`) and never
customer rows. Prod-verified (2026-06-06):

| Table | public (NULL) | private (org) |
|---|---|---|
| `skiller.skills` | 42,763 | 794 |
| `skiller.job_roles` | 22,315 | 2,381 |
| `skiller.specializations` | 1,442 | 154 |
| `skiller.categories` | 22 | 42 |
| `cms.studio_documents` | **0** | 3,060 |
| `cms.studio_tasks` | **0** | 2,353 |
| `cms.similarities` | 274 | 733 |

Embeddings/translations carry **no** `organization_id` → scope them via the **public parent**
(`… WHERE skill_id IN (SELECT id FROM skiller.skills WHERE organization_id IS NULL)`). Apply `deleted = false` on
skiller queries.

---

## DATABASE SCHEMA REFERENCE

Schemas: **public** (app), **skiller**, **jobsimulation**, **skillpath**, **cms**, **sentinel**.

### PUBLIC (app — users, orgs, profiles, subscriptions)
- **users**: id (uuid), clerk_id, email, firstname, lastname, picture, created_at, updated_at, deleted_at
- **organizations**: id (uuid), name, slug, org_type, is_anonymous, clerk_id, created_at, deleted_at
- **memberships**: id, role (varchar), user (uuid→users), organization_id (uuid), job_role_id (varchar), organization_role_id
- **user_skills**: id, user_id, skill_id (NodeID), competency_level, score, job_simulation_id, skill_path_id, source
- **user_skill_evidences**: id, user_skill_id, jobsimulation_session_id, evaluation_type, score, competency_level, strengths/weaknesses_feedback
- **membership_skills**: id, membership_id, skill_id (NodeID) + denormalized skill/specialization/category name+id, competency_level, score, is_core
- **user_basic_info**, **user_target_roles**, **user_experiences**, **local_jobsimulation_sessions**, **local_skill_path_sessions**, **subscriptions**, **user_resources**, **ai_usages** (large — cost tracking; catalog-only)

### SKILLER (taxonomy, job roles, AI matching — the v1.2 taxonomy snapshot surface, ≈2.1 GB)
- **skills**: id (uuid), node_id (varchar unique), name, description, parent (→specializations), deleted, **organization_id (nullable)**
- **specializations**: id, node_id, name, description, parent (→categories), deleted, organization_id
- **categories**: id, node_id, name, description, deleted, organization_id
- **job_roles**: id, node_id, name, description, deleted, **organization_id (nullable)**
- **job_role_skills**: id, job_role_id (→job_roles), skill_id (→skills), is_core, level
- **skill_embeddings**: id (bigint), small_embedding3 (**vector**), skill_id (uuid→skills.id) — 692 MB, ~689 MB is the pgvector index
- **job_role_embeddings**: id, small_embedding3 (vector), job_role_id (uuid)
- **skill_translations**, **job_role_translations**: i18n, scoped via parent

> Taxonomy hierarchy joins by UUID parent FK: `categories.id ← specializations.parent ← skills.parent`.

### JOBSIMULATION (AI-driven simulations)
- **sessions**: id, sim_id (template uuid), owner_id (=users.id), status, completion_status, result_status, organization_id, timestamps
- **actors**, **interactions** (large), **ai_interactions**, **activity_events** (large), **task_checks** / **task_sub_checks**
- Validation hierarchy: **validation_attempt_results** → **validation_attempt_skill_results** (skill = NodeID) → **validation_criterion_results** → **validation_check_results**; plus **validation_results** (per-interaction)
- Anticheat: **anticheat_results** → **anticheat_evidences**

### SKILLPATH (learning paths)
- **skill_path_sessions** (root, versioned) → **chapter_sessions** → **step_sessions** (step_sessions.simulation_id = template; .last_simulation_session_id = latest attempt)

### CMS (the app-Postgres cms schema — NOT the public Directus template library)
- **studio_documents**, **studio_tasks** — studio-room generated, **100% org-scoped customer data** (excluded from snapshots)
- **similarities** (+ similarity_skills/categories/features join tables)
- The **public** simulation/skill-path **template library** is **not** in this app-Postgres `cms` schema — it
  lives in the **`directus` schema of the same `postgres` database** (served at `content.anthropos.work`, but
  reachable read-only over the same DSN / `marco_read`). Its public subset
  (`private = false AND tenant_id IS NULL AND status = 'published'`) is the v1.2 M10 content-snapshot source.

### SENTINEL (authorization)
Casbin RBAC/ABAC policies (`casbin_rules` / `casbin_rule`). Rarely queried directly — use the sentinel RPC API.
