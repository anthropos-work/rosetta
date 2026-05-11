# CustomerIO Sync Service

## Role & Responsibility

CustomerIO Sync is a **background data-sync service** that pushes user and organization data from the Anthropos backend database into [Customer.io](https://customer.io/) for marketing automation, lifecycle email campaigns, and product analytics.

It's a one-directional pipeline: PostgreSQL `public` schema → Customer.io API. No inbound traffic, no consumers inside the platform.

## Architecture

* **Repo**: `git@github.com:anthropos-work/customerio-sync` (private)
* **Language**: Go
* **Local port**: 8080 (HTTP — health/metrics)
* **Profile**: `customerio-sync` (NOT in default `graphql` profile)
* **Build pattern**: **unique among Anthropos services** — Docker builds it directly from the GitHub URL, the repo is not cloned locally by `make init`.

### Compose definition

```yaml
customerio-sync:
  build:
    context: git@github.com:anthropos-work/customerio-sync.git#main
    ssh: ["default"]
    args:
      VERSION: dev
      GH_ACCESS_TOKEN: $GH_PAT
  ports: ["8080:8080"]
  environment:
    - DB_CONNECTION_BACKEND=postgresql://postgres@postgresql:5432/postgres?sslmode=disable&search_path=public
  profiles: [customerio-sync, all]
  depends_on:
    postgresql: { condition: service_healthy }
```

Note `context: git@github.com:...#main` — Docker BuildKit clones the repo at build time, no local checkout needed. This works because the build runs inside an SSH-agent-forwarded context (`ssh: ["default"]`) with `$GH_PAT` available.

### Dependencies

* **PostgreSQL** (`public` schema, read access via `DB_CONNECTION_BACKEND`)
* **Customer.io API** (external — credentials live in `platform/.env`)

## Interface Discovery

The service does not expose business APIs to other platform services — it's a pure sync worker. Port 8080 likely serves a health/metrics endpoint.

For protocol and field-mapping details, see the repo:

```bash
gh repo clone anthropos-work/customerio-sync
```

## Local Development

This service is **off by default**. To run it locally:

```bash
cd platform
docker compose --profile customerio-sync up --build -d customerio-sync
```

You'll need Customer.io API credentials in `platform/.env`. For most local-development tasks you do not need this service.

## Production

Runs as an ECS task. Configuration via `platform/.env` and Terraform-managed secrets.

## Related Documentation

* [Backend (app)](./backend.md) — source of user/org data
* [Service Taxonomy](../architecture/service_taxonomy.md) — orchestration profiles
* [External Services](../architecture/external_services.md) — Customer.io as an integrated SaaS

## Notes

The "build from GitHub URL" pattern is intentional: this service is operationally simple and rarely changes, so day-to-day developers do not need it cloned. If you need to iterate on it, clone it as a sibling of `platform/` and add it to `repos.yml` temporarily.
