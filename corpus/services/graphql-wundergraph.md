# GraphQL Gateway (`graphql-wundergraph`) — RETIRED

> ## ⛔ Retired 2026-07-31 — the router is gone, not folded in
>
> The WunderGraph/Cosmo federation router was **decommissioned on 2026-07-31**. This is not a
> merge banner: unlike [cms](./cms.md), [jobsimulation](./jobsimulation.md),
> [messenger](./messenger.md), [storage](./storage.md) and [customerio-sync](./customerio-sync.md)
> — which were folded *into* `app` and still have a domain to point at — the router had **nothing to
> fold in**. It was configuration and a build pipeline in front of one subgraph, so retiring it
> deleted the hop rather than relocating it.
>
> **Clients call `backend`'s own gqlgen endpoint directly:**
>
> | | Endpoint |
> |---|---|
> | Production | `https://gql.anthropos.work/graphql/query` |
> | Local compose | `http://localhost:8082/graphql/query` |
>
> **There is no supergraph and no composition step any more.** `cms-in-app v8.0` had already taken
> the supergraph from 2 subgraphs to 1 (`backend`), which made the router a pure extra hop — the
> whole point of a federation router is to *federate*, and there was nothing left to federate. A
> schema change now needs **no** `wgc compose`, no gateway rebuild and no gateway restart; `backend`
> serves its own schema.
>
> What was destroyed:
>
> * **Compose** — there is **no `graphql` service** in `platform/docker-compose.yml` and **no
>   `graphql` profile**. The default profile was renamed `graphql` → **`core`** at platform
>   `0dab54d` (`PROFILE ?= core`). **Host port `:5050` is free — nothing listens there.**
>   Beware: `docker compose --profile <unknown>` exits **0** and selects nothing, so a stale
>   `--profile graphql` in an old runbook starts **nothing** while reporting success.
> * **Infrastructure** — `module.wundergraph_euwest1` was **deleted**. Its ECS service, task
>   definition, target group, ALB rule (**priority 810**, left free), Cloud Map entry, log group and
>   ACM cert are destroyed, along with the **`wundergraph.anthropos.work`** Route53 alias. The ECR
>   repository was carried through the removal with `removed { destroy = false }` (the module sets no
>   `force_delete`, so destroying a repo holding images would have aborted the apply) and then
>   **hand-deleted on 2026-08-05** — `production-wundergraph` no longer exists in ECR and the
>   `removed` block in `services.tf` is now inert.
> * **Repo** — `anthropos-work/graphql-wundergraph` is **archived** (read-only on GitHub). It is
>   gone from `platform/repos.yml`, so `make init` does not clone it. Nothing should reference it.
> * **The env var names survive on purpose.** `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`
>   (next-web-app, ant-academy), `GRAPHQL_SCHEMA_FOR_GEN` (codegen) and `VITE_GRAPHQL_ENDPOINT`
>   (studio-desk) were **not** renamed — the names are historical, the values point at `backend`.
>   Renaming them is a coordinated code + deploy change across three frontends and their deploy
>   configs; do not "fix" the name in a doc or a `.env`. See
>   [Retired, but the names live on](#retired-but-the-names-live-on) below.
>
> For the current GraphQL surface see **[Backend (`app`)](./backend.md)** and
> [External Services → GraphQL endpoint](../architecture/external_services.md#graphql-endpoint--backends-own-gqlgen-server).

---

## Retired, but the names live on

Three things outlived the router and still carry its name. All three are **live and correct** —
they are fossils, not bugs.

| Name | Where | What it means today |
|------|-------|---------------------|
| `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` | `next-web-app` (`apps/web`, `apps/hiring`, `apps/integration`), `ant-academy` | The GraphQL endpoint the browser/SSR client calls. Value is **`backend`**: `http://localhost:8082/graphql/query` locally, `https://gql.anthropos.work/graphql/query` in production. Read in `packages/graphql/src/{server/server.graphql.ts,hooks/useGraphql.tsx}` and several route handlers. |
| `GRAPHQL_SCHEMA_FOR_GEN` | `next-web-app` codegen | The schema endpoint `graphql-codegen` introspects. Also `http://localhost:8082/graphql/query` — so **`backend` must be running** for `pnpm codegen`. |
| `VITE_GRAPHQL_ENDPOINT` | `studio-desk` | Same endpoint, Vite-side. Defaulted to `http://localhost:8082/graphql/query` in `platform/docker-compose.yml`. |

One more fossil, in Terraform: `infrastructure/terraform/production/services.tf` still passes
`wundergraph_endpoint = ""` to the `next-web-app` module. It is a **dead input** — every project
reads `backend_gql_endpoint` instead — but the variable has no default in next-web-app `v2.133.0`,
so it must still be supplied. It can be deleted once next-web-app's own drop-wundergraph change
ships and the module ref is bumped past it.

> **Do not blank `backend_gql_endpoint` as a "rollback".** There is nothing to roll back to: the
> projects would fall through to the empty `wundergraph_endpoint`.

---

## Historical record — what the gateway was

**Everything below this line describes the retired service.** It is kept as a record of what ran
until 2026-07-31; none of it is a live instruction. Do not follow the commands.

### Role & responsibility (as it was)

* Served the platform's Apollo Federation v2 **supergraph** from a WunderGraph **Cosmo Router** at
  one endpoint. By the end it composed a **single** subgraph — `backend`. All four other subgraphs
  had folded into it in sequence: `skiller` (July 2026), `skillpath` ("skillpath-in-app",
  M502→M507), `jobsimulation` ("jobsim-in-app"), and `cms` ("cms-in-app v8.0" — the 2→1 step).
* Served the unified `/graphql` endpoint that every frontend and Studio-Desk talked to (host
  `:5050` locally).
* Carried the jobsimulation GraphQL **subscriptions** over Server-Sent Events (`sse_post`) — those
  are served by `backend` now.
* Provided a GraphQL **playground + introspection** in dev/compose; both were disabled in
  production. (On `gql.anthropos.work` today, introspection and the Apollo Sandbox playground are
  disabled at the **app** layer in staging/production.)

> **"WunderGraph" vs "Cosmo Router" — same thing.** Cosmo was WunderGraph's Apollo-Federation
> product. The repo was named `graphql-wundergraph`, the compose service was `graphql`, the runtime
> binary was the Cosmo Router, and the frontend env var is `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`. They
> all referred to this one gateway. (Only the env var still exists.)

### Architecture & code map (as it was)

* **Codebase**: `graphql-wundergraph` — repo `git@github.com:anthropos-work/graphql-wundergraph`
  (**archived**)
* **Runtime**: prebuilt Go binary image `ghcr.io/wundergraph/cosmo/router:0.275.0` (pinned)
* **Build tooling**: `wgc@0.104.0` (WunderGraph Cosmo CLI) on a `node:22.11-alpine` build stage
* **Federation**: Apollo Federation v2, `federation_version: =2.3.2` (pinned)
* **Database**: none — stateless gateway (no DB, no Redis)
* **Ports**: host **5050 → container 8080** (router `listen_addr 0.0.0.0:8080`, `graphql_path
  /graphql`)

> **There was no application source here.** `package.json` was a stub
> (`{"name":"graphql-wundegraph"}` — note the misspelling, carried in the repo). The product was
> configuration + a build pipeline. That is exactly why retiring it removed a hop rather than
> relocating a domain.

```
Dockerfile.dev                          Local build: regenerated SDL from SIBLING repos, then wgc compose
Dockerfile                              Prod build: composed from the committed schemas/ dir as-is
config.compose.yaml / .dev / .prod      Router runtime config (playground/introspection/CORS/35MB body)
supergraph-config-compose.yaml / .dev / .prod   Subgraph routing URLs per environment
subgraphs.conf                          Per-subgraph version pins consumed by CI (GitHub Releases path)
schemas/                                Committed concatenated SDL — by the end just backend.graphqls
ci/                                     update-subgraph.sh (gh release download), release-supergraph.sh, utils.sh
terraform/                              ECS service "wundergraph" (eu-west-1, port 8080, /health)
.github/workflows/                      release.yml (tag → ECR → infra dispatch), supergraph-update.yml
```

### Why build-time composition mattered

The supergraph `config.json` was **baked into the image at build time** by
`wgc router compose -i supergraph-config.yaml -o config.json`. The router did **not** live-introspect
running subgraphs. Consequences, while it ran:

* Adding/changing a subgraph **or a single field** required re-running `wgc compose` and
  **rebuilding + restarting** the image — there was **no hot reload**.
* `make up` rebuilt `graphql` whenever any subgraph schema changed, because the build context was
  the parent dir (`..`) holding all sibling repos.

**None of that applies now.** `backend` serves its own gqlgen schema; a `.graphqls` change is picked
up by rebuilding `backend` alone, and `pnpm codegen` introspects `backend` directly.

### Dependencies (as they were)

* **Upstream consumers**: every GraphQL client — `next-web-app`, `studio-desk`, mobile — hit the
  router at `:5050/graphql`. They now hit `backend` directly.
* **Downstream (composed subgraphs)**: by the end, `app` (as `backend`) alone.
* **CI/prod**: GitHub Releases on `anthropos-work/app` (schema artifacts) +
  `anthropos-work/infrastructure` Terraform + `release-service.yml`.

### Notable gotchas (historical)

* **Composition was static** — no live subgraph discovery; schema changes needed a rebuild + restart.
* **Pinned versions**: cosmo router `0.275.0`, `wgc 0.104.0`, federation `2.3.2`.
* **Repo name is misspelled** `graphql-wundegraph` in `package.json` and the repo's own `CLAUDE.md`
  heading.
* `repos.yml` tagged it `type: node-npm`, but there were effectively **no npm deps** — the "node"
  stage existed only to run `wgc`.

## Related Documentation

* [Backend (`app`)](./backend.md) — the service that serves GraphQL now
* [External Services → GraphQL endpoint](../architecture/external_services.md#graphql-endpoint--backends-own-gqlgen-server) — integration view, frontend wiring, troubleshooting
* [Frontend Architecture](../architecture/frontend_architecture.md) — how `next-web-app` consumes the schema
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Dependency Map](../architecture/dependency_map.md)
* [Platform repo](../ops/platform_repo.md) — profiles and Make targets after the rename
