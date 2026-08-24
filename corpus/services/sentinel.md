# Sentinel Service

> ## ⚠️ Merged into `app` at v11.0 — no longer a standalone service (the 8th fold)
>
> Platform **`766df6c`** (2026-08-11, *"chore(compose): remove sentinel service and related
> configurations"*) — on top of `befca6d` (*"SENTINEL_MODE + SENTINEL_DB_CONNECTION for backend (app
> v11.0)"*) and `48de408` (*"backend is the PDP — drop SENTINEL_MODE and the sentinel dependency"*) —
> **deleted the `sentinel` compose service and its `repos.yml` entry together.** `make init` no longer
> clones the repo; `repos.yml` declares **three** repos (`app`, `next-web-app`, `studio-desk`).
>
> **Where it went.** The Casbin PDP is **`app/internal/sentinel/`**, ported from this repo at tag
> **v0.24.2** (`app/internal/sentinel/doc.go:10`; `f2c46190`, which is this repo's `origin/main` today —
> the port source and the repo head are the same commit). It is wired **exactly once**, at
> `app/main.go:305`, under the source's own comment *"The single wiring point for authorization (v11.0).
> There is no switch and no RPC path: app IS the PDP"*, and `log.Fatalf` on failure. `app/main.go:322`
> logs `authorization: in-process PDP active`. The caller-side facade is `app/internal/authorization/`
> (`NewSentinelManager`).
>
> **The RPC edge is gone, and it took the whole listener with it.** `AUTHORIZATION_ADDRESS` occurs
> **0** times across `docker-compose.yml`, `common.yml` and `repos.yml` at `766df6c`, and in `app`'s Go
> tree exactly once — inside `app/sentinel_wiring_test.go:57`, a test named `TestNoRPCPathSurvives` that
> asserts its **absence**. `app` also deleted its own Connect-RPC server (`app/main.go:1310`, *"NO RPC
> SERVER"* — the port-8081 mux that carried Users / Organizations / Skiller / JobSimulation / CMS / lab),
> so **a local stack now has no cross-process Connect-RPC edge at all.** The handler objects survive and
> are called in-process.
>
> **What did NOT move: the schema.** Unlike the seven folds before it, the policy tables stayed put —
> `sentinel.casbin_rules` is still in the `sentinel` schema, reached through
> `SENTINEL_DB_CONNECTION` (`docker-compose.yml:25`, `search_path=sentinel`) and migrated by `app`'s own
> `make migrations-sentinel` (`app/Makefile:80-81`, `atlas migrate diff --env sentinel`). That is why the
> [fenced migration map](../architecture/platform-migration-status.md) grades production **`mid-fold`**
> and not `merged-into-app`: §1's test requires code-owned **and** tables-in-`public` **and**
> standalone-scaled-to-zero, and two of the three fail.
>
> **Net-new, and documented nowhere before M258 iter-18: cross-replica policy invalidation.** Casbin
> holds the whole policy set in memory, so a write updates the writing process's model and **every other
> replica serves its boot policy forever** — silently, with no error to notice. v11.0 M1102 closes it
> with a casbin `persist.Watcher` over **Redis Pub/Sub**, channel **`sentinel:policy:invalidate`**
> (`app/internal/sentinel/watcher.go:55`). Pub/Sub deliberately, **not** the Watermill consumer-group
> plumbing app uses elsewhere: invalidation must **fan out** to every replica, and an XREADGROUP `>`
> entry goes to exactly one consumer. The enforcer also had to become a `*casbin.SyncedEnforcer` — and
> `SetWatcher` alone was not enough, because it binds the reload callback to the embedded
> **unsynchronised** enforcer (`AttachWatcher` replaces it explicitly).
>
> ⚠️ **This matters to Rosetta's own tooling, and it broke it.** Any tool that writes policy rows in
> **raw SQL** — as the seeders do — bypasses casbin's write path entirely, so **no invalidation is
> published and the in-process enforcer never reloads.** A stale enforcer refuses every org-scoped read
> and write with `forbidden` **at HTTP 200** (the silent-403 class). That is exactly what M258 iter-16
> diagnosed — **every red Playthrough in that batch was org-scoped and every green one was user-scoped**,
> a partition that named the mechanism before any log was read. The fix is
> to publish to `sentinel:policy:invalidate` after any out-of-band policy write.
>
> **The repo is not archived and not deleted** — `origin/main` is `f2c46190` (v0.24.2). Its production
> ECS service was still declared at the last `infrastructure` reading (`module "sentinel_euwest1"`,
> `13c248e6`, 2026-08-07 — see [`org-repos.md` §3](../architecture/org-repos.md)); the platform names
> the teardown **M1103** in `docker-compose.yml:85`. Everything below the line describes the
> **standalone** service and is retained for reading pre-fold source and for the Casbin model, which
> was ported unchanged.

---

## Role & Responsibility

⚠️ **The present tense below is historical from `766df6c` onward — read it as *"the standalone service
did"*.** The behaviour is unchanged; the process boundary is gone.

Sentinel is the **centralized authorization service** of the platform. Its **only** live caller is **`app`** — including the jobsimulation and cms authz call sites it absorbed in-process — which reaches it over Connect-RPC to check permissions before executing operations. (There are no `cms` or `jobsimulation` containers left to receive the address: platform `d11a403` deleted both compose services along with `roadrunner`, so at `0dab54d` `AUTHORIZATION_ADDRESS` is set in exactly **one** block — backend's, `docker-compose.yml:48`.) **`messenger` is not a caller** — and ⚠️ **the evidence clause has to be past tense, because there is no messenger compose block to read (corrected M257x iter-115).** At platform `0c91421d`, `docker-compose.yml` declares **five** services — `sentinel` (`:5`), `backend` (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`), `gotenberg` (`:170`) — and `git grep -n messenger 0c91421d -- docker-compose.yml common.yml repos.yml` returns **only comments**. `838d907` (*"drop the storage, messenger and customerio-sync containers"*, 2026-08-05) deleted it. The sentence read *"its compose block sets no `AUTHORIZATION_ADDRESS` and declares no `depends_on: sentinel`"* in the **present** tense, presupposing a block that does not exist — true at `0dab54d` (where the block began at `docker-compose.yml:156`) and silently expired. **Ten other corpus sites already recorded the deletion, two of them in this same file**, neither framed as a retraction of this sentence — one survivor against ten witnesses. What survives, and was re-derived: messenger's Go source imports no authorization client — **and the receipt is stronger than this page published it.** ⚠️ **Corrected M257x iter-140: the claim was *"returns **one unrelated hit**"*; run verbatim it returns ZERO** — `git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` in `stack-demo/messenger` (`HEAD = fa47850d`, tree clean) exits **1 with 0 lines**, as does the widened whole-tree `git grep -in "authoriz" fa47850d` (exit 1, 0 lines). The **positive control holds in the identical invocation form** — `git grep "colony" fa47850d -- '*.go' go.mod` returns hits in messenger's own `cmd/` package (three files; **the pins are deliberately not published — they are bare cross-repo heads that resolve against the wrong clone**, the defect `§5` rule 63(c) names) — so the form is sound and the absence is real. Independently reproduced by an iter-135 adjudicator and again at iter-140. **The conclusion survives and is strengthened; only the count was wrong** — but a published receipt that does not reproduce teaches a reader to distrust the paragraph around it, which is why it is corrected rather than left as a harmless overstatement; [`clerk-integration.md`](./clerk-integration.md) says the same ("storage, messenger — no auth"). It wraps **Casbin v3** with a PostgreSQL-backed policy store and a single in-memory enforcer that handles all of Anthropos's authorization patterns.

Sentinel does **not** handle authentication — that's Clerk's job. It also does not validate JWTs (the shared `authn` library does that in each consuming service). Sentinel only answers *"is this subject allowed to perform this action on this object?"*.

## Architecture & Code Map

* **Codebase**: `sentinel` (local) — repo `git@github.com:anthropos-work/sentinel`
* **Language**: Go 1.26 (`go.mod:3` `go 1.26.0`; `Dockerfile:2` / `Dockerfile.dev:2` `golang:1.26-bookworm`)
* **Framework**: Connect-RPC, Casbin v3
* **Database**: PostgreSQL `sentinel` schema (single table: `casbin_rules`) — **no Ent ORM**
* **Port**: 8087 (HTTP + Connect-RPC; `PORT=8087` in compose, same on host and inside container). The sentinel binary's default is 8080 (per its own README/CLAUDE.md), but the platform compose overrides it to 8087 explicitly.
* **Profile**: ⚠️ **none — there is no `sentinel` compose service since `766df6c`.** *(Historical: it was always on, declaring no `profiles:` key, so it ran with every `make up`.)*
* **No outbound RPC** to other platform services — sentinel is a leaf
* **No Redis, no GraphQL, no background workers** — stateless request/response only

### Why no Ent / no GraphQL?

Sentinel's data model is exactly one table (Casbin's `casbin_rules`), and it doesn't participate in the federation gateway because its concerns are orthogonal to product data. Keeping it lean makes it cheap to operate (256 CPU / 128 MB on ECS — `terraform/locals.tf:4-5`) and easy to test (all unit tests use in-memory enforcers, no DB fixtures).

### Casbin model

The enforcer defines **6 request types, 6 policy types, 3 role groupings, 6 matchers** to handle the various authorization patterns in one place:

| Matcher | Pattern | Use case |
|---------|---------|----------|
| `m` | User-tier quota | A user passes if they are in the policy's tier (`g(user, tier)`) OR the policy tier is `TIER_FREE` (free-tier policies act as an unconditional baseline, substituted from the proto enum `Tier_TIER_FREE`), AND the requested `count` <= the tier `max`. |
| `m6` | Org-level feature quota | Org-membership check via `g3(org, user)` AND `feat` match AND requested `count` <= the org policy `max` (no tier logic). ⚠️ **And no `'default'` escape** — unlike `m2`/`m3`/`m5`, the `p6` row must name the org id, so there is no one-row way to cover every org. |
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

The `sentinel` schema must exist before sentinel can start — but **you no longer create it by hand**
(corrected M257x iter-130). `68272003` (2026-08-04) added a **second Atlas pipeline, owned by `app`**:
`app/atlas.hcl:50-64` declares `env "sentinel"`, and `app/Makefile:59-60` records that *"`atlas migrate
apply --env sentinel` creates the schema itself, and that is what local/CI actually run"* (@
`ad9f3c498`). The `sentinel` repo itself still has no `atlas.hcl` and no `terraform/migrations/` at
`f2c461903`, which is why `repos.yml` correctly marks it `migrations: false`. The `extensions` schema must also exist (pgvector is required by other migrations, not by sentinel itself — but the platform setup creates both together). See [setup_guide.md §6](../ops/setup_guide.md) for the schema-creation step. Without it, sentinel crash-loops with `pq: no schema has been selected`.

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

> ⚠️ **HISTORICAL — `make init` no longer clones `sentinel` since platform `766df6c`, so there is no
> `sentinel/` directory to `cd` into on a fresh stack.** The requirement did not disappear, it MIGRATED:
> the policy rows still have to exist in the `sentinel` schema, and on a Rosetta stack it is
> `rosetta-extensions/stack-seeding`'s `PolicyGrantsSeeder` that puts them there
> (see [`../ops/seeding-spec.md`](../ops/seeding-spec.md)). ⚠️ **And `platform`'s own `make
> bootstrap-dev` still tries this path and is therefore BROKEN** — it hard-requires
> `../sentinel/init_policy.sql`, then `docker compose restart sentinel`. Clone the repo by hand if you
> need the file.

```bash
cd sentinel        # HISTORICAL — repo not cloned by `make init` since 766df6c
make initdb        # runs init_policy.sql via psql against a HARD-CODED local DSN
```

`make initdb` does NOT read `DB_CONNECTION` — it always targets `postgresql://postgres@localhost:5432/postgres` (sslmode=disable). It works only against a local Postgres on port 5432, and relies on `init_policy.sql` being schema-qualified (`sentinel.casbin_rules`) so the seed lands in the right schema regardless of search_path. For a non-local DB, run psql with your own DSN: `psql "$DB_URL" -f init_policy.sql`. The seed defines the base RBAC rules; "default" org policies apply to all organizations unless overridden by org-specific entries — **but NOT on every matcher, and the exception is load-bearing.** `m2`, `m3` and `m5` each carry an explicit escape (`'default' == p2.org || r2.org == p2.org` at `casbin.go:22`, the same shape at `:23`, and `p5.org == 'default'` at `:44`), and matcher `m` has the templated tier default at `:21`. **`m6` has NO `'default'` escape at all** (`:45`, measured — zero occurrences of `default` in that line), so a `p6` row naming the org `default` authorizes **nobody**: the row must name the real org id. ⚠️ Read as an unqualified rule, this sentence sends you to a one-row fix that silently grants nothing — which is exactly the trap `M267` hit, and why it is qualified here.

### Superadmin / elevated local grants

`init_policy.sql` intentionally omits sensitive capabilities (notably `org:feature:taxonomy:write`, see init_policy.sql:63-66). To grant them locally, apply the on-demand seed:

```bash
psql "$DB_URL" -f local_superadmin_grants.sql
```

This grants org-scoped `taxonomy:write` (p3) to every org admin, and contains a commented-out block of global superadmin rules (p4: impersonation/cross-org reads, global content & taxonomy writes) that you uncomment after substituting a concrete user UUID. **Local-only — never run in staging or production.** After applying, the enforcer must be told: ⚠️ **there is no sentinel process to restart and no `Reload` RPC since v11.0** — publish to the Redis channel `sentinel:policy:invalidate` (`app/internal/sentinel/watcher.go:55`), or restart `backend`. A raw-SQL policy write that skips this leaves the in-process enforcer serving its boot policy and every org-scoped request failing `forbidden` at HTTP 200. **No demo or dev stack ever had this file applied until v2.8 M256** — a demo therefore granted `taxonomy:write` to nobody while production grants it to `admin`; see [`../ops/seeding-spec.md`](../ops/seeding-spec.md) § Status (`PolicyGrantsSeeder` + `stackseed --policy-check`).

## Testing

> ⚠️ **HISTORICAL — the standalone repo's own suite.** The ported package's tests live in `app` now
> (`app/internal/sentinel/*_test.go`, plus `app/sentinel_wiring_test.go`) and run with `go test ./...`
> from `stack-dev/app`. Run the block below only in a hand-made clone.

```bash
cd sentinel        # HISTORICAL — repo not cloned by `make init` since 766df6c
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
