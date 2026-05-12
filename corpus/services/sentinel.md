# Sentinel Service

## Role & Responsibility

Sentinel is the **centralized authorization service** of the platform. Every other service (`app`, `skiller`, `jobsimulation`, `cms`, `skillpath`, `messenger`) calls Sentinel via Connect-RPC to check permissions before executing operations. It wraps **Casbin v3** with a PostgreSQL-backed policy store and a single in-memory enforcer that handles all of Anthropos's authorization patterns.

Sentinel does **not** handle authentication — that's Clerk's job. It also does not validate JWTs (the shared `authn` library does that in each consuming service). Sentinel only answers *"is this subject allowed to perform this action on this object?"*.

## Architecture & Code Map

* **Codebase**: `sentinel` (local) — repo `git@github.com:anthropos-work/sentinel`
* **Language**: Go 1.25
* **Framework**: Connect-RPC, Casbin v3
* **Database**: PostgreSQL `sentinel` schema (single table: `casbin_rules`) — **no Ent ORM**
* **Port**: 8087 (HTTP + Connect-RPC; `PORT=8087` in compose, same on host and inside container). The sentinel binary's default is 8080 (per its own README/CLAUDE.md), but the platform compose overrides it to 8087 explicitly.
* **Profile**: always on (no `profiles:` declared in compose — runs with every `make up`)
* **No outbound RPC** to other platform services — sentinel is a leaf
* **No Redis, no GraphQL, no background workers** — stateless request/response only

### Why no Ent / no GraphQL?

Sentinel's data model is exactly one table (Casbin's `casbin_rules`), and it doesn't participate in the federation gateway because its concerns are orthogonal to product data. Keeping it lean makes it cheap to operate (256 CPU / 256 MB on ECS) and easy to test (all unit tests use in-memory enforcers, no DB fixtures).

### Casbin model

The enforcer defines **6 request types, 6 policy types, 3 role groupings, 6 matchers** to handle the various authorization patterns in one place:

| Matcher | Pattern | Use case |
|---------|---------|----------|
| `m`, `m6` | User-tier quota | "Free users can run X simulations/month" |
| `m2` | Org role-based action | "Admins can invite members" |
| `m3` | Org feature access | Role-based gating of insights, workforce, members CRUD, etc. |
| `m4` | Direct user action | Subject-object-action equality |
| `m5` | Membership content action | Org membership + keyMatch on object patterns |

Role groupings:

* `g(user, tier)` — `TIER_FREE` / `TIER_PREMIUM`
* `g2(org, user, role)` — `admin` / `member` / `manager` / `candidate` per org
* `g3(org, membership)` — enables/disables org memberships for feature access

### Key directories

```
cmd/root.go                     Cobra CLI, server bootstrap
internal/
  authorization/
    casbin.go                   Casbin model definition + enforcer factory
    manager.go                  Manager: Check, BulkCheck, org feature credits
    enforcer_conversions.go     Domain types ↔ Casbin enforce requests
    parse.go                    Custom Casbin ParseFloat function
    test.go                     newTestEnforcer() for in-memory tests
  rpcsrv/rpc.go                 Connect-RPC handler (all RPC methods)
init_policy.sql                 DB seed: table creation + default policies
terraform/                      AWS ECS (base_internal_service module)
```

## Interface Discovery

### Connect-RPC (`AuthorizationService`)

| Method | Purpose |
|--------|---------|
| `Check` / `BulkCheck` | Unified check with oneof request types |
| `CheckFeature` | User tier quota check |
| `CheckOrganizationFeature` | Org-level feature quota check |
| `AddUserToTier` / `RemoveUserFromTier` | Manage user tier groupings |
| `GetQuotas` / `GetOrganizationQuotas` | Read policy quotas |
| `OrgAddUserToRole` / `OrgRemoveUserFromRole` / `OrgReplaceUserRole` | Manage org role assignments |
| `OrgClearAll` | Remove all `g2` + `g3` policies for an org |
| `OrgCheckPermission` | Legacy org action check (read-only) |
| `OrgAllowUserToUseFeature` / `OrgDisallowUserToUseFeature` | Manage `g3` membership feature access |
| `OrgMembershipsAllowedToUseFeature` | List memberships with feature access |
| `OrgGetOrganizationFeatureCredits` / `OrgSetOrganizationFeatureCredits` | Manage org feature credit budgets |
| `Reload` | Hot-reload policies from DB |

Consumed via `AUTHORIZATION_ADDRESS=http://sentinel:8087` in every other service's compose env.

## Dependencies

* **Upstream consumers**: every other Anthropos service that gates requests (`app`, `cms`, `skiller`, `skillpath`, `jobsimulation`, `messenger`)
* **Downstream**: PostgreSQL (`sentinel` schema, table `casbin_rules`)
* **No outbound RPC** to other platform services

## Local Development

### First-run schema setup

The `sentinel` schema must exist before sentinel can start. The `extensions` schema must also exist (pgvector is required by other migrations, not by sentinel itself — but the platform setup creates both together). See [setup_guide.md §6](../ops/setup_guide.md) for the schema-creation step. Without it, sentinel crash-loops with `pq: no schema has been selected`.

### Run in Docker

Sentinel is always part of any `make up` (no profile gate). To restart just sentinel:

```bash
cd platform
docker compose restart sentinel
make logs S=sentinel
```

### Run natively

```bash
cd platform
make dev S=sentinel
cd ../sentinel
go run main.go
```

### Seed default policies

```bash
cd sentinel
make initdb        # applies init_policy.sql against the configured DB
```

The `init_policy.sql` seed defines the base RBAC rules. "default" org policies apply to all organizations unless overridden by org-specific entries.

## Testing

```bash
cd sentinel
go test -v ./...
```

All tests use in-memory Casbin enforcers — no PostgreSQL or fixtures required.

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DB_CONNECTION` | yes | — | PostgreSQL DSN with `search_path=sentinel` |
| `PORT` | no | `8087` | HTTP + Connect-RPC port (compose sets this explicitly; binary default is 8080) |
| `ENVIRONMENT` | no | — | Environment name |
| `SERVICE_NAME` | no | `sentinel` | Logging label |
| `SENTRY_DSN` | no | — | Sentry error tracking |

## Operational Notes

* **Hot reload**: changes made directly to the `casbin_rules` table (e.g. via a migration or a manual fix) need a `Reload` RPC call to take effect. Changes made through Sentinel's own RPC methods (`OrgAddUserToRole` etc.) are picked up immediately.
* **Default policies vs org overrides**: most policies live as "default" rows. An org can override behavior by inserting its own rows with the org ID as the policy key.

## Related Documentation

* [External Services](../architecture/external_services.md) — Clerk (auth), Sentinel (authz) split
* [Backend (app)](./backend.md) — biggest consumer
* [Dependency Map](../architecture/dependency_map.md)
* [Security & Compliance](../architecture/security_compliance.md)
