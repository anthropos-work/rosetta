# GraphQL Gateway (`graphql-wundergraph`)

> Service-level / developer map for the federated GraphQL gateway. For the
> integration view (how frontends consume it, Clerk/CORS, troubleshooting) see
> [External Services → GraphQL Gateway](../architecture/external_services.md#graphql-gateway--wundergraph-cosmo-router).

## Role & Responsibility

* **Primary Goal**: Federate the platform's four Go GraphQL subgraphs into a single
  Apollo Federation v2 **supergraph**, served by a WunderGraph **Cosmo Router** at one endpoint.
  (The former `skiller` subgraph was removed when skiller merged into `app`, July 2026 — the
  `backend` subgraph now serves the taxonomy types/queries.)
* **Key Functions**:
  * Compose `app` (subgraph name `backend`), `jobsimulation`, `cms`, and `skillpath` into one schema.
  * Serve the unified `/graphql` endpoint that every frontend and Studio-Desk talks to (host `:5050` locally).
  * Carry `jobsimulation` GraphQL **subscriptions** over Server-Sent Events (`sse_post`).
  * Provide a GraphQL **playground + introspection** in dev/compose; both are disabled in production.

> **"WunderGraph" vs "Cosmo Router" — same thing.** Cosmo is WunderGraph's
> Apollo-Federation product. The repo is named `graphql-wundergraph`, the compose
> service is `graphql`, the runtime binary is the Cosmo Router, and the frontend env
> var is `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`. They all refer to this one gateway.

## Architecture & Code Map

* **Codebase**: `graphql-wundergraph` (local) — repo `git@github.com:anthropos-work/graphql-wundergraph`
* **Runtime**: prebuilt Go binary image `ghcr.io/wundergraph/cosmo/router:0.275.0` (pinned)
* **Build tooling**: `wgc@0.104.0` (WunderGraph Cosmo CLI) on a `node:22.11-alpine` build stage
* **Federation**: Apollo Federation v2, `federation_version: =2.3.2` (pinned)
* **Database**: none — stateless gateway (no DB, no Redis)
* **Ports**: host **5050 → container 8080** (router `listen_addr 0.0.0.0:8080`, `graphql_path /graphql`)

> **There is no application source here.** `package.json` is a stub
> (`{"name":"graphql-wundegraph"}` — note the misspelling, carried in the repo).
> The product is configuration + a build pipeline.

### Key files & directories

```
Dockerfile.dev                          Local build: regenerates SDL from SIBLING repos, then wgc compose
Dockerfile                              Prod build: composes from the committed schemas/ dir as-is
config.compose.yaml / .dev / .prod      Router runtime config (playground/introspection/CORS/35MB body)
supergraph-config-compose.yaml / .dev / .prod   Subgraph routing URLs per environment
subgraphs.conf                          Per-subgraph version pins consumed by CI (GitHub Releases path)
schemas/                                Committed concatenated SDL (backend|cms|jobsimulation|skillpath).graphqls
ci/                                     update-subgraph.sh (gh release download), release-supergraph.sh, utils.sh
terraform/                              ECS service "wundergraph" (eu-west-1, port 8080, /health)
.github/workflows/                      release.yml (tag → ECR → infra dispatch), supergraph-update.yml
```

### Build-time, static composition (important)

The supergraph `config.json` is **baked into the image at build time** by
`wgc router compose -i supergraph-config.yaml -o config.json`. The router does
**not** live-introspect running subgraphs. Consequences:

* Adding/changing a subgraph **or a single field** requires re-running `wgc compose`
  and **rebuilding + restarting** the image — there is **no hot reload**.
* `make up` rebuilds `graphql` whenever any subgraph schema changes because the build
  context is the parent dir (`..`) holding all sibling repos.

The two Dockerfiles source schemas differently:

| Dockerfile | Schema source | Used by |
|------------|---------------|---------|
| `Dockerfile.dev` | COPYs SDL fresh from **sibling repos** (`../app`, `../cms`, …) and `awk`-concatenates | local `make up` (compose) |
| `Dockerfile` | Uses the **committed `schemas/*.graphqls`** as-is | production CI build |

## Interface Discovery

| Interface | Kind | Detail |
|-----------|------|--------|
| `/graphql` | GraphQL | Unified federated endpoint. Playground + introspection ON in compose/dev, OFF in prod. |
| `/health` | HTTP | Health path used by the ECS ALB target group. |
| `make run` | repo Makefile | Standalone build+run (README notes aarch-only); reads a local `subgraphs/` checkout — **not** the platform compose flow. |
| `make updatesubg` | repo Makefile | Concatenates `subgraphs/<svc>/…` SDL into `schemas/<svc>.graphqls`. |
| `supergraph-update.yml` | GitHub Action | `workflow_dispatch` bumps subgraph versions in `subgraphs.conf`, re-downloads SDL from GitHub Releases, opens a PR. |

### Subgraph routing (compose mode)

Routing URLs use Docker **service names** on `app-network` (deliberately avoiding
`host.docker.internal`/`extra_hosts` so it works on Docker Desktop *and* native Linux):

| Subgraph | Routing URL (Docker network) | Notes |
|----------|------------------------------|-------|
| `backend` (the `app` service) | `http://backend:8082/graphql/query` | subgraph named `backend`, maps to repo/service `app` (includes the taxonomy queries absorbed from the former `skiller` subgraph) |
| `jobsimulation` | `http://jobsimulation:8400/query` | **subscriptions** via `sse_post` |
| `cms` | `http://cms:8090/query` | |
| `skillpath` | `http://skillpath:8100/query` | |

> `dev` mode uses `host.docker.internal:<port>`; `prod` uses AWS service-discovery
> DNS where all subgraphs share container port **8080**. Use the `-compose` config
> for local dev (there is **no** `-local` variant).

## Dependencies

* **Upstream consumers**: every GraphQL client — `next-web-app`, `studio-desk`, mobile — hits the router at `:5050/graphql`.
* **Downstream (composed subgraphs)**: `app` (as `backend`), `jobsimulation`, `cms`, `skillpath`.
* **Compose `depends_on`** (all `service_started`): `backend`, `jobsimulation`, `cms`, `skillpath`, **`storage`** — note `storage` is **not** a GraphQL subgraph but is in the startup-order list.
* **CI/prod**: GitHub Releases on `anthropos-work/{app,jobsimulation,cms,skillpath}` (schema artifacts) + `anthropos-work/infrastructure` Terraform + `release-service.yml`.

## Local Development

### Run in Docker (normal path)

```bash
cd platform
make up                 # default graphql profile — builds & starts the gateway at :5050
make logs S=graphql     # tail router logs
# Open the playground:
open http://localhost:5050
```

`graphql` lives in the **`graphql`** and **`all`** profiles only. A single-service
profile like `make up PROFILE=cms` does **not** start the gateway — bring up the
`graphql` profile (or `all`) to get a usable federated endpoint.

### Smoke-test the endpoint

```bash
curl -s http://localhost:5050/graphql \
  -H 'content-type: application/json' \
  -d '{"query":"{ __typename }"}'
# → {"data":{"__typename":"Query"}}
```

### Recompose the supergraph manually

```bash
cd graphql-wundergraph
wgc router compose -i supergraph-config-compose.yaml -o config.json
```

## Environment Variables

These are **Docker build args** (not runtime secrets) — they select which config files get baked in.

| Variable | Default (compose) | Description |
|----------|-------------------|-------------|
| `ENVIRONMENT` | `compose` | Picks `config.<ENVIRONMENT>.yaml` → `config.yaml` (router runtime behavior). Also an ECS env var in prod. |
| `ENVIRONMENT_CONFIG` | `compose` | Picks `supergraph-config-<ENVIRONMENT_CONFIG>.yaml` → which routing URLs are baked in. |
| `CONFIG_PATH` | `config.yaml` | Tells the Cosmo router which config file to load. |
| `GH_PAT` / `GH_ACCESS_TOKEN` | — | CI/prod build only (pull private resources). The compose `graphql` service instead uses `build.ssh: [default]`. |
| `VERSION` / `ARCHITECTURE` / `APOLLO_ELV2_LICENSE` | — | Prod CI build args (ECR image tag, `linux/amd64`, accept Apollo ELv2 license). |

The router itself is **stateless** — no DB/Redis/secret env vars at runtime (the compose service mounts `.env` via `env_file` but the router does not consume secrets).

## Testing

There is **no unit/integration test suite** in this repo (only `terraform/tests/`
fixtures). The README documents only `make run`. Schema correctness is enforced at
compose time by `wgc` (federation composition will fail the build on an invalid supergraph).

## Notable Gotchas

* **Composition is static** — no live subgraph discovery; schema changes need a rebuild + restart.
* **Pinned versions**: cosmo router `0.275.0`, `wgc 0.104.0`, federation `2.3.2`.
* **Repo name is misspelled** `graphql-wundegraph` in `package.json` and the repo's own `CLAUDE.md` heading.
* `repos.yml` tags it `type: node-npm`, but there are effectively **no npm deps** — the "node" stage exists only to run `wgc`.
* The repo's own `CLAUDE.md` "Version Tracking" line and `-local.yaml` references are **stale**; `subgraphs.conf` is the real version source of truth and the config variants are `compose`/`dev`/`prod`.

## Related Documentation

* [External Services → GraphQL Gateway](../architecture/external_services.md#graphql-gateway--wundergraph-cosmo-router) — integration view, frontend wiring, troubleshooting
* [Frontend Architecture](../architecture/frontend_architecture.md) — how `next-web-app` consumes the supergraph
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Dependency Map](../architecture/dependency_map.md)
