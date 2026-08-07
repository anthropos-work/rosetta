# Sentinel Service

## Role & Responsibility

Sentinel is the **centralized authorization service** of the platform. Its **only** live caller is **`app`** — including the jobsimulation and cms authz call sites it absorbed in-process — which reaches it over Connect-RPC to check permissions before executing operations. (There are no `cms` or `jobsimulation` containers left to receive the address: platform `d11a403` deleted both compose services along with `roadrunner`, so at `0dab54d` `AUTHORIZATION_ADDRESS` is set in exactly **one** block — backend's, `docker-compose.yml:48`.) **`messenger` is not a caller** — and ⚠️ **the evidence clause has to be past tense, because there is no messenger compose block to read (corrected M257x iter-115).** At platform `0c91421d`, `docker-compose.yml` declares **five** services — `sentinel` (`:5`), `backend` (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`), `gotenberg` (`:170`) — and `git grep -n messenger 0c91421d -- docker-compose.yml common.yml repos.yml` returns **only comments**. `838d907` (*"drop the storage, messenger and customerio-sync containers"*, 2026-08-05) deleted it. The sentence read *"its compose block sets no `AUTHORIZATION_ADDRESS` and declares no `depends_on: sentinel`"* in the **present** tense, presupposing a block that does not exist — true at `0dab54d` (where the block began at `docker-compose.yml:156`) and silently expired. **Ten other corpus sites already recorded the deletion, two of them in this same file**, neither framed as a retraction of this sentence — one survivor against ten witnesses. What survives, and was re-derived: messenger's Go source imports no authorization client (`git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` returns one unrelated hit, against `colony` present as a positive control); [`clerk-integration.md`](./clerk-integration.md) says the same ("storage, messenger — no auth"). It wraps **Casbin v3** with a PostgreSQL-backed policy store and a single in-memory enforcer that handles all of Anthropos's authorization patterns.

Sentinel does **not** handle authentication — that's Clerk's job. It also does not validate JWTs (the shared `authn` library does that in each consuming service). Sentinel only answers *"is this subject allowed to perform this action on this object?"*.

## Architecture & Code Map

* **Codebase**: `sentinel` (local) — repo `git@github.com:anthropos-work/sentinel`
* **Language**: Go 1.26 (`go.mod:3` `go 1.26.0`; `Dockerfile:2` / `Dockerfile.dev:2` `golang:1.26-bookworm`)
* **Framework**: Connect-RPC, Casbin v3
* **Database**: PostgreSQL `sentinel` schema (single table: `casbin_rules`) — **no Ent ORM**
* **Port**: 8087 (HTTP + Connect-RPC; `PORT=8087` in compose, same on host and inside container). The sentinel binary's default is 8080 (per its own README/CLAUDE.md), but the platform compose overrides it to 8087 explicitly.
* **Profile**: always on (no `profiles:` declared in compose — runs with every `make up`)
* **No outbound RPC** to other platform services — sentinel is a leaf
* **No Redis, no GraphQL, no background workers** — stateless request/response only

### Why no Ent / no GraphQL?

Sentinel's data model is exactly one table (Casbin's `casbin_rules`), and it doesn't participate in the federation gateway because its concerns are orthogonal to product data. Keeping it lean makes it cheap to operate (256 CPU / 128 MB on ECS — `terraform/locals.tf:4-5`) and easy to test (all unit tests use in-memory enforcers, no DB fixtures).

### Casbin model

The enforcer defines **6 request types, 6 policy types, 3 role groupings, 6 matchers** to handle the various authorization patterns in one place:

| Matcher | Pattern | Use case |
|---------|---------|----------|
| `m` | User-tier quota | A user passes if they are in the policy's tier (`g(user, tier)`) OR the policy tier is `TIER_FREE` (free-tier policies act as an unconditional baseline, substituted from the proto enum `Tier_TIER_FREE`), AND the requested `count` <= the tier `max`. |
| `m6` | Org-level feature quota | Org-membership check via `g3(org, user)` AND `feat` match AND requested `count` <= the org policy `max` (no tier logic). |
| `m2` | Org role-based action | "Admins can invite members" |
| `m3` | Org feature access | Role-based gating of insights, workforce, members CRUD, etc. |
| `m4` | Direct user action | Subject-object-action equality |
| `m5` | Membership content action | Org membership + keyMatch on object patterns |

Role groupings:

* `g(user, tier)` — `TIER_FREE` / `TIER_PREMIUM`
* `g2(org, user, role)` — `admin` / `member` / `candidate` / `content_creator` per org (the four `MembershipRole` values in `app/internal/data/ent/enum/membership.go:8-15`; `init_policy.sql` seeds policies for all four, `content_creator` in its own block at `init_policy.sql:88-118` with a dedicated `internal/authorization/casbin_content_creator_test.go`)
* `g3(org, membership)` — enables/disables org memberships for feature access

> **There is no `manager` role.** It appears nowhere in `init_policy.sql`, and only as a fixture string in
> sentinel's own tests (`internal/authorization/casbin_test.go`, `internal/rpcsrv/rpc_test.go`); a live
> stack's `select distinct v2 from sentinel.casbin_rules where p_type='g2'` returns `admin` / `member` /
> `candidate` only. In the demo world "manager" is a **persona** label, not a Casbin role — granting it
> yields a membership with **no policy rows at all**, which is exactly the silent-403 failure mode this
> corpus warns about elsewhere.

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

Consumed via `AUTHORIZATION_ADDRESS=http://sentinel:8087`, set in exactly **one** compose block at platform `0c91421` — **backend**'s, `docker-compose.yml:48` (measured: 1 occurrence across `docker-compose.yml`, `common.yml` and `.env_example`). So the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`**, and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is declared at `:170` and is in the default `core` profile at `:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The correctly-scoped form is the model at [`architecture_overview.md:343`](../architecture/architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*. No other declared service sets `AUTHORIZATION_ADDRESS` — the only ones left to check are `gotenberg`, `studio-desk` and `next-web-app`, and none has the env or a sentinel dependency. The blocks that used to carry it are gone rather than corrected: `jobsimulation`, `cms` and `roadrunner` at `d11a403`, then `storage`, `messenger` and `customerio-sync` at `838d907` — so there is nothing off-path left to hold the address either.

## Dependencies

* **Upstream consumers**: **`app` only** — the sole service that gates requests through Sentinel, and the only compose block that is given its address (`docker-compose.yml:48`). `messenger` and `storage` never called it, and neither is a compose service any more (deleted at `838d907`); `cms`, `jobsimulation` and `roadrunner` went earlier, at `d11a403`
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
make initdb        # runs init_policy.sql via psql against a HARD-CODED local DSN
```

`make initdb` does NOT read `DB_CONNECTION` — it always targets `postgresql://postgres@localhost:5432/postgres` (sslmode=disable). It works only against a local Postgres on port 5432, and relies on `init_policy.sql` being schema-qualified (`sentinel.casbin_rules`) so the seed lands in the right schema regardless of search_path. For a non-local DB, run psql with your own DSN: `psql "$DB_URL" -f init_policy.sql`. The seed defines the base RBAC rules; "default" org policies apply to all organizations unless overridden by org-specific entries.

### Superadmin / elevated local grants

`init_policy.sql` intentionally omits sensitive capabilities (notably `org:feature:taxonomy:write`, see init_policy.sql:63-66). To grant them locally, apply the on-demand seed:

```bash
psql "$DB_URL" -f local_superadmin_grants.sql
```

This grants org-scoped `taxonomy:write` (p3) to every org admin, and contains a commented-out block of global superadmin rules (p4: impersonation/cross-org reads, global content & taxonomy writes) that you uncomment after substituting a concrete user UUID. **Local-only — never run in staging or production.** After applying, restart sentinel or call the `Reload` RPC so the Casbin enforcer picks up the new rows. **No demo or dev stack ever had this file applied until v2.8 M256** — a demo therefore granted `taxonomy:write` to nobody while production grants it to `admin`; see [`../ops/seeding-spec.md`](../ops/seeding-spec.md) § Status (`PolicyGrantsSeeder` + `stackseed --policy-check`).

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
