# GraphQL Gateway (`graphql-wundergraph`)

> ## ⚠️ THE ROUTER IS GONE FROM LOCAL DEV — and its two states differ
>
> | | production | a fresh local stack @ platform origin HEAD |
> |---|---|---|
> | the router | **DESTROYED** (corrected iter-123) — `module.wundergraph_euwest1` is **deleted** from `infrastructure/terraform/production/services.tf` @ `13c248e6`; `:509-517` records that the apply destroyed *"its ECS service, task definition, target group, ALB rule (priority 810), Cloud Map entry, log group, ACM cert and the `wundergraph.anthropos.work` alias"*, leaving only a `removed{}` for the ECR (`:521`) which was hand-deleted **2026-08-05** — *"so production-wundergraph is gone and this block is now inert."* **This cell previously read "still declared — `graphql-wundergraph/terraform/main.tf:20` `service_desired_count = 1`", and that line is ORPHANED DEAD CODE** | **deleted** — no `graphql` compose service, no `repos.yml` entry |
> | the repo | **ARCHIVED on GitHub, 2026-07-30** (read-only) | not cloned by `make init` |
> | what a frontend talks to | the router | **`backend` directly**, `http://localhost:8082/graphql/query` |
>
> Platform `b56d731` + `360efd4` (merged **`2adcf71`**, 2026-07-31) dropped the router from
> `docker-compose.yml` **and** `repos.yml` and re-pointed local dev at `backend`. **There is no `:5050` on a
> local stack.** **The `graphql` profile is gone too:** `0dab54d` (*"rename graphql -> core"*) renamed it,
> so `Makefile:10` reads `PROFILE ?= core` and the token appears in **no `profiles:` key at all** — the
> **five** that exist are `core`, `backend`, `all`, `studio-desk`, `frontend`. (This list read *"the
> eight"* until platform `838d907`, 2026-08-05: `storage-legacy`, `customerio-sync` and `messenger`
> were deleted along with the three services that declared them, so those three tokens are now
> retired exactly as `graphql` is.) Asking for `graphql` therefore **exits 0** and starts only the
> always-on floor (`postgresql`, `redis`, `sentinel`), which is worse than an error.
>
> The supergraph is **ONE** subgraph: `915da06` (2026-07-29) folded the cms subgraph into `backend`
> (cms-in-app v8.0) and deleted the `jobsimulation` entry in the **same commit** — a **3 → 1** step,
> not 2 → 1. `supergraph-config-prod.yaml` lists `backend` alone and `schemas/` holds
> `backend.graphqls` alone.
>
> Everything below the fold describes the gateway **as it still exists in production and in the archived
> repo**. Read [`../architecture/platform-migration-status.md`](../architecture/platform-migration-status.md)
> — the fenced map — before acting on any local-development instruction on this page.

> Service-level / developer map for the federated GraphQL gateway. For the
> integration view (how frontends consume it, Clerk/CORS, troubleshooting) see
> [External Services → GraphQL Gateway](../architecture/external_services.md#graphql-gateway--wundergraph-cosmo-router).

## Role & Responsibility

* **Primary Goal**: Serve the platform's Apollo Federation v2 **supergraph** from a WunderGraph
  **Cosmo Router** at one endpoint. Since **cms-in-app v8.0** the supergraph composes a **single
  subgraph** — `backend`. All four other subgraphs folded into it, though **not one per service
  merge**. Counting entries in `supergraph-config-prod.yaml` at each commit:

  | commit | date | subgraphs | what changed |
  |---|---|---|---|
  | `749dc86~1` | — | **5** | `backend`, `skiller`, `jobsimulation`, `cms`, `skillpath` |
  | `749dc86` | 2026-06-24 | **4** | `skiller` removed |
  | `7c17e63` | 2026-07-21 | **3** | `skillpath` removed ("skillpath-in-app") |
  | `915da06` | 2026-07-29 | **1** | `cms` **and** `jobsimulation` removed together (cms-in-app v8.0) |

  So cms-in-app was the **3 → 1** step. The `jobsimulation` subgraph **outlived jobsim-in-app**: the
  service merged into `app` well before its supergraph entry and `schemas/jobsimulation.graphqls`
  were deleted, and both went at `915da06` alongside cms's (`git show --name-status 915da06` marks
  the two SDL files `D`). `915da06`'s own subject line says *"supergraph 2→1"* and is **wrong** —
  the tree it was committed against lists three. The `backend` subgraph now serves the taxonomy,
  skill-path session, simulation and content types/queries alike.
* **Key Functions**:
  * Compose `app` (subgraph name `backend`) — the only entry in `supergraph-config-*.yaml`; `subgraphs.conf` carries a single `BACKEND=` pin.
  * Serve the unified `/graphql` endpoint that every frontend and Studio-Desk talks to. **In production
    only** — since platform `2adcf71` there is no router in compose, and locally the frontends talk to
    `backend` at `:8082/graphql/query`.
  * ~~Carry the jobsimulation GraphQL **subscriptions** over Server-Sent Events (`sse_post`)~~ —
    **HISTORICAL, and there is nothing to carry today.** The `subscription: protocol: "sse_post"`
    block sat on the `jobsimulation` entry alone and died with that entry at `915da06`; no
    `supergraph-config-*.yaml` carries a `subscription:` block now. Nor is there anything to
    subscribe to on `backend` directly: the composed `schemas/backend.graphqls` declares `type
    Query` and `type Mutation` and **no `type Subscription`** (the `Subscription` substring hits in
    that SDL are Stripe/plan field names — `activeSubscription`, `stripeSubscriptionId`, `type
    PlanSubscription`). On mainline the protocol read `sse_post` for its whole life; `bba862f`
    ("change subscription protocol from sse_post to ws", 2026-02-25) never merged — it exists only
    on `remotes/origin/feat/use-web-socket`. See
    [External Services → What the gateway provides](../architecture/external_services.md#graphql-gateway--wundergraph-cosmo-router).
  * Provide a GraphQL **playground + introspection** in dev/compose; both are disabled in production.

> **"WunderGraph" vs "Cosmo Router" — same thing.** Cosmo is WunderGraph's
> Apollo-Federation product. The repo is named `graphql-wundergraph`, the compose
> service **was** `graphql` (deleted at `2adcf71`), the runtime binary is the Cosmo Router, and the frontend env
> var is `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`. They all refer to this one gateway.

## Architecture & Code Map

* **Codebase**: `graphql-wundergraph` (local) — repo `git@github.com:anthropos-work/graphql-wundergraph`
* **Runtime**: prebuilt Go binary image `ghcr.io/wundergraph/cosmo/router:0.275.0` (pinned)
* **Build tooling**: `wgc@0.104.0` (WunderGraph Cosmo CLI) on a `node:22.11-alpine` build stage
* **Federation**: Apollo Federation v2, `federation_version: =2.3.2` (pinned)
* **Database**: none — stateless gateway (no DB, no Redis)
* **Ports**: **8080 → 8080** (router `listen_addr 0.0.0.0:8080`, `graphql_path /graphql`). **There is no
  `5050` at platform HEAD** — `grep -c 5050 docker-compose.yml` returns 0. The host-`5050` mapping was
  published here in the present tense while **`:193`** of this same doc already said `localhost:5050`
  refuses the connection (this cited `:174-176` until M257x iter-98 — that is the compose line-number
  caveat, a different construct); corrected M257x iter-46 (the same claim iter-40 swept at 8 sites elsewhere)

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
schemas/                                Committed concatenated SDL — now just backend.graphqls
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
* ~~`make up` rebuilds `graphql`~~ — **HISTORICAL: there is no `graphql` compose service since `2adcf71`, so `make up` builds no router at all.** It *used to* rebuild whenever any subgraph schema changed, because the build
  context is the parent dir (`..`) holding all sibling repos.

The two Dockerfiles source schemas differently:

| Dockerfile | Schema source | Used by |
|------------|---------------|---------|
| `Dockerfile.dev` | COPYs SDL fresh from **sibling repos** (`../app`, …) and `awk`-concatenates. **Not used by compose at all any more** — there is no `graphql` compose service since `2adcf71`. It *was* what compose built, from `2c85211` (2026-02-27) right up to the deletion. | (legacy local path) |
| `Dockerfile` | Uses the **committed `schemas/*.graphqls`** as-is | production CI build |

> **Which one compose used, historically.** Verified against `platform`'s `docker-compose.yml`
> history: the router service (`wundergraph`, renamed `graphql` at `d92e84e`) carried **no
> `dockerfile:` key** from its introduction (`63d285c`) through `719befb`, so those builds took the
> default — the **production `Dockerfile`**, committed schemas and all — first from a `git@…` context,
> then from `../graphql-wundergraph` (`a2a3ee6`). `2c85211` (2026-02-27) added
> `dockerfile: Dockerfile.dev`, `67ba772` moved the context to `..` (path becoming
> `graphql-wundergraph/Dockerfile.dev`), and it stayed `Dockerfile.dev` until the service block was
> deleted at `360efd4` (merged as `2adcf71`). So for its **final five months** compose built
> `Dockerfile.dev` — which is why a subgraph SDL change rebuilt the router, as the *Build-time, static composition* bullets at `:114-117` of this doc describe (specifically `:116-117`, the struck-through *"`make up` rebuilds `graphql`"* bullet: *"It **used to** rebuild whenever any subgraph schema changed, because the build context is the parent dir (`..`) holding all sibling repos"*). **Not `:84`**, which is the *Ports* bullet and is about `8080`/`5050` — corrected M257x iter-102.

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

**Historical — compose no longer has a router service to route.** Kept because the archived repo's
configs still contain these rows and a reader will find them there.

| Subgraph | Routing URL (Docker network) | Notes |
|----------|------------------------------|-------|
| `backend` (the `app` service) | `http://backend:8082/graphql/query` | **the only surviving subgraph.** Named `backend`, maps to repo/service `app` — and now also serves everything the four folded subgraphs used to |
| ~~`jobsimulation`~~ | ~~`http://jobsimulation:8400/query`~~ | service folded into `backend` (jobsim-in-app), but **the subgraph entry outlived the merge** — it was deleted at `915da06`, in the same commit as cms's |
| ~~`cms`~~ | ~~`http://cms:8090/query`~~ | folded into `backend` at `915da06` (cms-in-app v8.0) — the step that took the supergraph **3 → 1** |

> All four non-`backend` subgraphs were removed as their services merged into `app`:
> `skiller` (July 2026), `skillpath` ("skillpath-in-app", M502→M507), and then `jobsimulation`
> **and** `cms` together at `915da06` ("cms-in-app v8.0" — the **3 → 1** step; see the ladder under
> *Role & Responsibility*, and note that `915da06`'s subject line's "2→1" is wrong). The `backend`
> subgraph serves all of their types/queries. **Only 1 subgraph remains.**
>
> `dev` mode uses `host.docker.internal:<port>`; `prod` uses AWS service-discovery
> DNS where all subgraphs share container port **8080**. Use the `-compose` config
> for local dev (there is **no** `-local` variant).

## Dependencies

* **Upstream consumers**: every GraphQL client — `next-web-app`, `studio-desk`, mobile. **In production** they
  hit the router; **locally they hit `backend` directly** at `:8082/graphql/query`
  (`docker-compose.yml:135` studio-desk's `VITE_GRAPHQL_ENDPOINT`, `:160` next-web-app's
  `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` — each also baked as a build arg, at `:119` and `:151`), because
  the router service no longer exists in compose. Those line numbers move on every compose
  clean-up — they were `:220`/`:236` at `0dab54d` and `:334`/`:352` at `2adcf71` — so grade the
  construct, not the offset.
* **Downstream (composed subgraphs)**: `app` (as `backend`) — and, historically, `jobsimulation` and `cms`.
* ~~**Compose `depends_on`**~~ — moot: there is no compose service. Historically `backend`, `jobsimulation`, `cms`, **`storage`** (note `storage` was **not** a GraphQL subgraph but was in the startup-order list).
* **CI/prod**: GitHub Releases on **`anthropos-work/app` only** (schema artifacts) + `anthropos-work/infrastructure` Terraform + `release-service.yml`. `ci/update-subgraph.sh:9` carries **exactly one** `gh release download`, `-R anthropos-work/app`; the `jobsimulation` and `cms` downloads were **deleted at `915da06`** when those subgraphs folded into `backend`. (This bullet claimed all three until M257x iter-49 — the two bullets above it already carried their historical fence; this one did not.)

## Local Development

### Run in Docker — **no longer possible; there is no service to start**

`make up` starts the **default** profile, and neither the current default nor the old one contains a router:
`2adcf71` deleted the service while `Makefile:10` still read `PROFILE ?= graphql`, and `0dab54d`
(*"rename graphql -> core"*) then renamed the profile itself — `Makefile:10` now reads `PROFILE ?= core`.
**⚠️ At `0dab54d` the `graphql` token appears in no `profiles:` key at all**, so asking for it does not
error: compose **exits 0** and starts only the always-on floor (`postgresql`, `redis`, `sentinel`) — a
silent no-op that looks like a live stack. Tailing a `graphql` service likewise has nothing to tail, and
`http://localhost:5050` refuses the connection. If you are following an older runbook that says otherwise,
the runbook predates 2026-07-31.

### Smoke-test the endpoint — against `backend`

```bash
# the local GraphQL endpoint IS the backend subgraph now
curl -s -w '\nHTTP=%{http_code}\n' http://localhost:8082/graphql/query \
  -H 'content-type: application/json' \
  -d '{"query":"{ __typename }"}'
# → {"errors":[{"message":"unknown viewer: Forbidden"}],"data":null}
#   HTTP=200
```

**That rejection *is* the healthy response — do not read it as a failure.** The endpoint is
default-closed: `app`'s GraphQL authorization middleware rejects any anonymous operation that isn't
`@public`-annotated, a federation query, or (in local dev only) pure schema introspection
(`app/internal/authorization/gqlauthz/gqlauthz.go:186`). `__typename` is **deliberately excluded**
from the introspection exemption in every environment — the app pins that with its own regression
tests, which drive `__typename` specifically (`gqlauthz_test.go`, `"bare __typename is not exempt"`
and `TestAnonymousRejectionLogsAtWarn`). This is stock behaviour, not something a demo patch does.
Note also that the transport is healthy at **HTTP 200**: GraphQL reports the refusal in the `errors`
array, not in the status code.

So what the probe actually proves is *the server is up, routing, parsing GraphQL and enforcing
authz*. The real failure signals are **connection refused / HTTP 000** (nothing listening — wrong
port; remember a demo stack's ports are offset, e.g. `18082` for `demo-1`), a **non-200 status**, or
an HTML/proxy body instead of JSON.

Two companion probes worth running:

```bash
curl -s http://localhost:8082/api/health                      # → "OK"   (liveness, no GraphQL)

# pure schema introspection IS exempt anonymously — in local dev only
curl -s http://localhost:8082/graphql/query \
  -H 'content-type: application/json' \
  -d '{"query":"{ __schema { queryType { name } } }"}'
# → {"data":{"__schema":{"queryType":{"name":"Query"}}}}
```

The `__schema` form is the one that returns real `data` unauthenticated, and it is what the
frontend's `pnpm codegen` relies on. It is gated on `colony.Development`, so on staging/prod it is
refused like everything else.

Note the path: **`/graphql/query`**, not `/graphql`. On `backend`, `/graphql` serves the Apollo Sandbox UI;
CORS preflight and auth happen at `/query`.

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
| `GH_PAT` / `GH_ACCESS_TOKEN` | — | CI/prod build only (pull private resources). The compose `graphql` service **used to** use `build.ssh: [default]` — that service no longer exists (`2adcf71`). |
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
* `repos.yml` **used to** tag it `type: node-npm` (the entry was deleted at `2adcf71`), but there were effectively **no npm deps** — the "node" stage exists only to run `wgc`.
* ⚠️ **CORRECTED M257x iter-115 — this sentence had a compound subject and only ONE half survives.** The repo's `CLAUDE.md` *"Version Tracking"* section is **NOT stale**: at `graphql-wundergraph` `60c229f3` it reads *"Service versions are tracked in `subgraphs.conf`. There is exactly **one** pin now: `BACKEND=v1.360.0`. The `CMS`, `JOBSIMULATION`, `SKILLER` and `SKILLPATH` entries were removed as each of those services merged into `app`"* — and `subgraphs.conf` at that same ref is the single line `BACKEND=v1.360.0`, a byte match. It **is** the current form of the very claim this bullet offered as its correction, and `git log -- CLAUDE.md` shows the checkout itself last rewrote it; a reader was being told to distrust an accurate section of the ground-truth repo. **What IS stale is the `-local.yaml` reference** — `CLAUDE.md:85` still says `wgc router compose -i supergraph-config-local.yaml`, while `ls supergraph-config-*.yaml` @ `60c229f3` returns `-compose`, `-dev`, `-prod` and no `-local`. So: `subgraphs.conf` is the version source of truth (as that CLAUDE.md already says), and the config variants are `compose`/`dev`/`prod`.

## Related Documentation

* [External Services → GraphQL Gateway](../architecture/external_services.md#graphql-gateway--wundergraph-cosmo-router) — integration view, frontend wiring, troubleshooting
* [Frontend Architecture](../architecture/frontend_architecture.md) — how `next-web-app` consumes the supergraph
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Dependency Map](../architecture/dependency_map.md)
