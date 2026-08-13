# TIER-1 ADJUDICATION BATCH 12 — 41 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 12-001
- **id**: `B12-001`
- **corpus site**: `corpus/architecture/service_taxonomy.md:222-224` (paragraph)
- **citation**: `studio-desk.md:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/studio-desk.md`  (450 lines)

**CLAIMING UNIT**

```md
`package.json`, 0 `.tsx`/`.jsx` in the repo). *"React"* was published here and contradicted by
[`studio-desk.md:20`](../services/studio-desk.md) (*"vanilla TS frontend, no framework"*); corrected
M257x iter-46 |
```

**CITED CONTENT**

```
    17  | Property | Value |
    18  |:---------|:------|
    19  | **Service Type** | Custom Application (Tier 2 - Studio Services) |
    20  | **Technology Stack** | TypeScript, Vite, Express.js (vanilla TS frontend, no framework) |
    21  | **Deployment** | Runs natively for dev (`npm run dev`), or containerized via the `studio-desk` docker-compose profile (ports 9000/9100). It `depends_on` **`backend` alone** — `docker-compose.yml:138-140` @ platform `0c91421`, with `profiles: [studio-desk, all]` at `:141` (both re-anchored M257x iter-87; they were `:223-225`/`:226` at `0dab54d`, before `838d907` deleted three service blocks above them). It *also* listed **`cms`** (`:337-341` @ `2adcf71`) until that container was deleted from compose at `d11a403`; there is no `cms` service to depend on now, and it never depended on `graphql`, which is likewise no longer a compose service. Built with `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query`. **⚠️ Asking for `studio-desk` as the only profile exits 1** — the profile selects `studio-desk` but *not* the `backend` it depends on, so compose rejects the whole project (`service "studio-desk" depends on undefined service "backend": invalid compose project`). Use `PROFILE=all`, which selects both. |
    22  | **Port(s)** | 9100 (frontend), 9000 (backend) - configurable via `.env` |
    23  | **Authentication** | Clerk |
```

## 12-002
- **id**: `B12-002`
- **corpus site**: `corpus/architecture/service_taxonomy.md:295-295` (table-row)
- **citation**: `code/src/graphql/server.js:14`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/graphql/server.js`  (24 lines)

**CLAIMING UNIT**

```md
| **Platform dependencies** | **A GraphQL client of the platform `app` (`backend`) at runtime** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (`code/src/graphql/server.js:14,18` — it **throws** when unset). **There is no separate "academy subgraph"**: the supergraph declares exactly **one** subgraph, `backend` — all three of `graphql-wundergraph`'s configs at `60c229f3` (`supergraph-config-prod.yaml`, `-dev`, `-compose`) carry a single `- name: backend` entry, and `schemas/` holds one file, `backend.graphqls` — and the academy types are **one SDL file inside it** — `app/internal/web/backend/graphql/graph/schemas/academy.graphqls`, 1 of 43 files in that directory at `app` `ad9f3c49`. Locally there is no router at all (deleted at platform `2adcf71`), so the endpoint resolves straight to `backend` `:8082/graphql/query`. Same statement, same words, at [`academy-backend.md`](../services/academy-backend.md) (*"There is no separate 'academy subgraph'"*), and this file says it again below — *"**`backend` alone (1)**"*. Reads: the course catalog is **DB-authoritative**, not the committed FS tree (`code/src/lib/backendContent.js:36,102-103`; `code/src/lib/serverTenant.js:145`). Writes: per-user progress, bookmarks, certificates and feedback POST through `code/app/api/academy/beacon/route.js:36,41-55` (`UPSERT_CHAPTER_PROGRESS`, `SET_LAST_ACTIVITY`, …). Also reuses platform Clerk; AI calls go straight to the providers (never through the platform `ai` library). No Connect-RPC, no Redis. |
```

**CITED CONTENT**

```
    11  
    12  import { GraphQLClient } from 'graphql-request'
    13  
    14  const endpoint = process.env.NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT
    15  
    16  export function createServerGraphQLClient({ token, additionalHeaders } = {}) {
    17      if (!endpoint) {
```

## 12-003
- **id**: `B12-003`
- **corpus site**: `corpus/architecture/service_taxonomy.md:295-295` (table-row)
- **citation**: `code/src/lib/backendContent.js:36`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/backendContent.js`  (190 lines)

**CLAIMING UNIT**

```md
| **Platform dependencies** | **A GraphQL client of the platform `app` (`backend`) at runtime** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (`code/src/graphql/server.js:14,18` — it **throws** when unset). **There is no separate "academy subgraph"**: the supergraph declares exactly **one** subgraph, `backend` — all three of `graphql-wundergraph`'s configs at `60c229f3` (`supergraph-config-prod.yaml`, `-dev`, `-compose`) carry a single `- name: backend` entry, and `schemas/` holds one file, `backend.graphqls` — and the academy types are **one SDL file inside it** — `app/internal/web/backend/graphql/graph/schemas/academy.graphqls`, 1 of 43 files in that directory at `app` `ad9f3c49`. Locally there is no router at all (deleted at platform `2adcf71`), so the endpoint resolves straight to `backend` `:8082/graphql/query`. Same statement, same words, at [`academy-backend.md`](../services/academy-backend.md) (*"There is no separate 'academy subgraph'"*), and this file says it again below — *"**`backend` alone (1)**"*. Reads: the course catalog is **DB-authoritative**, not the committed FS tree (`code/src/lib/backendContent.js:36,102-103`; `code/src/lib/serverTenant.js:145`). Writes: per-user progress, bookmarks, certificates and feedback POST through `code/app/api/academy/beacon/route.js:36,41-55` (`UPSERT_CHAPTER_PROGRESS`, `SET_LAST_ACTIVITY`, …). Also reuses platform Clerk; AI calls go straight to the providers (never through the platform `ai` library). No Connect-RPC, no Redis. |
```

**CITED CONTENT**

```
    33   */
    34  
    35  import { auth } from '@clerk/nextjs/server'
    36  import { createServerGraphQLClient } from '../graphql/server.js'
    37  import {
    38    GET_ACADEMY_CATALOG_SERIES,
    39    GET_ACADEMY_CATALOG_SKILL_PATHS,
```

## 12-004
- **id**: `B12-004`
- **corpus site**: `corpus/architecture/service_taxonomy.md:295-295` (table-row)
- **citation**: `code/src/lib/serverTenant.js:145`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/serverTenant.js`  (337 lines)

**CLAIMING UNIT**

```md
| **Platform dependencies** | **A GraphQL client of the platform `app` (`backend`) at runtime** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (`code/src/graphql/server.js:14,18` — it **throws** when unset). **There is no separate "academy subgraph"**: the supergraph declares exactly **one** subgraph, `backend` — all three of `graphql-wundergraph`'s configs at `60c229f3` (`supergraph-config-prod.yaml`, `-dev`, `-compose`) carry a single `- name: backend` entry, and `schemas/` holds one file, `backend.graphqls` — and the academy types are **one SDL file inside it** — `app/internal/web/backend/graphql/graph/schemas/academy.graphqls`, 1 of 43 files in that directory at `app` `ad9f3c49`. Locally there is no router at all (deleted at platform `2adcf71`), so the endpoint resolves straight to `backend` `:8082/graphql/query`. Same statement, same words, at [`academy-backend.md`](../services/academy-backend.md) (*"There is no separate 'academy subgraph'"*), and this file says it again below — *"**`backend` alone (1)**"*. Reads: the course catalog is **DB-authoritative**, not the committed FS tree (`code/src/lib/backendContent.js:36,102-103`; `code/src/lib/serverTenant.js:145`). Writes: per-user progress, bookmarks, certificates and feedback POST through `code/app/api/academy/beacon/route.js:36,41-55` (`UPSERT_CHAPTER_PROGRESS`, `SET_LAST_ACTIVITY`, …). Also reuses platform Clerk; AI calls go straight to the providers (never through the platform `ai` library). No Connect-RPC, no Redis. |
```

**CITED CONTENT**

```
   142   */
   143  export async function getServerCatalogView() {
   144    const eids = await getUserEids()
   145    const view = (await getBackendCatalogView(eids)) ?? (process.env.ACADEMY_DEMO_FS_PUBLISHED === '1' ? (v => ({ ...v, chapters: (v.chapters ?? []).map(({ _draft, _origin, ...c }) => c), series: (v.series ?? []).map(({ _draft, _origin, ...s }) => s), skillPaths: Object.fromEntries(Object.entries(v.skillPaths ?? {}).map(([k, { _draft, _origin, ...p }]) => [k, p])) }))(mergeDrafts(emptyCatalogView(), eids)) : emptyCatalogView())
   146    return draftsEnabled() ? mergeDrafts(view, eids) : view
   147  }
   148  
```

## 12-005
- **id**: `B12-005`
- **corpus site**: `corpus/architecture/service_taxonomy.md:295-295` (table-row)
- **citation**: `code/app/api/academy/beacon/route.js:36`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/app/api/academy/beacon/route.js`  (96 lines)

**CLAIMING UNIT**

```md
| **Platform dependencies** | **A GraphQL client of the platform `app` (`backend`) at runtime** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (`code/src/graphql/server.js:14,18` — it **throws** when unset). **There is no separate "academy subgraph"**: the supergraph declares exactly **one** subgraph, `backend` — all three of `graphql-wundergraph`'s configs at `60c229f3` (`supergraph-config-prod.yaml`, `-dev`, `-compose`) carry a single `- name: backend` entry, and `schemas/` holds one file, `backend.graphqls` — and the academy types are **one SDL file inside it** — `app/internal/web/backend/graphql/graph/schemas/academy.graphqls`, 1 of 43 files in that directory at `app` `ad9f3c49`. Locally there is no router at all (deleted at platform `2adcf71`), so the endpoint resolves straight to `backend` `:8082/graphql/query`. Same statement, same words, at [`academy-backend.md`](../services/academy-backend.md) (*"There is no separate 'academy subgraph'"*), and this file says it again below — *"**`backend` alone (1)**"*. Reads: the course catalog is **DB-authoritative**, not the committed FS tree (`code/src/lib/backendContent.js:36,102-103`; `code/src/lib/serverTenant.js:145`). Writes: per-user progress, bookmarks, certificates and feedback POST through `code/app/api/academy/beacon/route.js:36,41-55` (`UPSERT_CHAPTER_PROGRESS`, `SET_LAST_ACTIVITY`, …). Also reuses platform Clerk; AI calls go straight to the providers (never through the platform `ai` library). No Connect-RPC, no Redis. |
```

**CITED CONTENT**

```
    33  export const runtime = 'nodejs'
    34  export const dynamic = 'force-dynamic'
    35  
    36  const endpoint = process.env.NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT
    37  
    38  // op → { doc, varsFrom } — the supergraph mutation + how to shape its variables
    39  // from the beacon body. The body carries the SAME input the seam built (already
```

## 12-006
- **id**: `B12-006`
- **corpus site**: `corpus/architecture/service_taxonomy.md:298-298` (bullet)
- **citation**: `code/src/lib/serverTenant.js:115-145`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/serverTenant.js`  (337 lines)

**CLAIMING UNIT**

```md
- Static chapter *bodies* as JSON in `code/public/content/<series>/<skill-path>/` — but **the catalog that decides what is visible is read from the platform over GraphQL, not from this tree**. With `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` unset or the academy tables empty, the read degrades to an **empty grid**; it does *not* back-fill from the committed FS content (`code/src/lib/serverTenant.js:115-145` — *"there is NO FS-as-published fallback … not reversible-on-error"*). This is the "empty academy" demo symptom, and a **demo** only shows a populated grid because a rext demo-patch (`demo-stack/patches/academy-fs-published-fallback`) restores that fallback on the demo's ephemeral clone — it is not the shipped behaviour
```

**CITED CONTENT**

```
   112   * tenant metadata and is the one piece not yet modeled in the backend catalog,
   113   * so it passes through verbatim (same as the backend adapter does).
   114   */
   115  function emptyCatalogView() {
   116    return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }
   117  }
   118  
   119  /**
   120   * The catalog the current (authenticated) user is allowed to receive: public +
   121   * the user's org tenants. This is what authed RSCs pass to the client instead
   122   * of importing the raw CHAPTERS array — so tenant content never ships to a
   123   * browser that can't see it.
   124   *
   125   * DB-authoritative (M7): the catalog is read UNCONDITIONALLY from the academy
   126   * backend (`getBackendCatalogView(eids)` → already tenant-filtered server-side,
   127   * with a chapter-level eid refinement in the adapter). DB presence gates
   128   * visibility — a path/chapter renders only if the backend returns it; there is
   129   * NO FS-as-published fallback. The eids are threaded in for the adapter's
   130   * chapter-level tenant refinement. A null backend result (not-composed /
   131   * outage) resolves to the EMPTY view rather than the committed FS catalog: the
   132   * cutover is intentional, not reversible-on-error.
   133   *
   134   * DEV DRAFT LAYER (M8): when `draftsEnabled()` (dev + the ACADEMY_SHOW_DRAFTS
   135   * opt-in; NEVER in production — hard-blocked), FS-only content (present-on-FS ∧
   136   * absent-from-DB) is merged on top of the DB view tagged `_draft: true`, so an
   137   * author previews locally-committed-but-not-yet-exported content. Off (incl. all
   138   * production) → the DB view passes through verbatim, zero behavior change. The
   139   * merge runs server-side; the client still receives only the threaded view
   140   * (catalog-client-boundary intact). The eids thread through so a tenant-scoped
   141   * draft stays tenant-gated.
   142   */
   143  export async function getServerCatalogView() {
   144    const eids = await getUserEids()
   145    const view = (await getBackendCatalogView(eids)) ?? (process.env.ACADEMY_DEMO_FS_PUBLISHED === '1' ? (v => ({ ...v, chapters: (v.chapters ?? []).map(({ _draft, _origin, ...c }) => c), series: (v.series ?? []).map(({ _draft, _origin, ...s }) => s), skillPaths: Object.fromEntries(Object
```

## 12-007
- **id**: `B12-007`
- **corpus site**: `corpus/architecture/service_taxonomy.md:299-299` (bullet)
- **citation**: `code/tests/unit/next-scaffold.test.js:106`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/tests/unit/next-scaffold.test.js`  (225 lines)

**CLAIMING UNIT**

```md
- **No service worker / no offline caching** — the Serwist 9 layer was REMOVED (v0.5 M1). `code/package.json` has no `serwist`/`workbox` dependency, no `sw.*` is emitted, `RegisterServiceWorker.jsx` is now a kill-switch that *unregisters* any surviving worker, and the repo regression-fences the removal (`code/tests/unit/next-scaffold.test.js:106,111`; `react-compiler-config.test.js:41`). **The web-app MANIFEST survives** (`public/academy-manifest.json`, `display: standalone`, declared at `code/app/layout.jsx:132`), so the app is still installable — it is simply online-only. Offline chapter bundling survives only in the Expo mobile app
```

**CITED CONTENT**

```
   103      expect(pkg.scripts.start).toBe('next start')
   104    })
   105  
   106    it('v0.5 M1: build is plain "next build" — no serwist SW step', () => {
   107      expect(pkg.scripts.build).toBe('next build')
   108      expect(pkg.scripts.build).not.toMatch(/serwist/)
   109    })
```

## 12-008
- **id**: `B12-008`
- **corpus site**: `corpus/architecture/service_taxonomy.md:299-299` (bullet)
- **citation**: `react-compiler-config.test.js:41`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/tests/unit/react-compiler-config.test.js`  (44 lines)

**CLAIMING UNIT**

```md
- **No service worker / no offline caching** — the Serwist 9 layer was REMOVED (v0.5 M1). `code/package.json` has no `serwist`/`workbox` dependency, no `sw.*` is emitted, `RegisterServiceWorker.jsx` is now a kill-switch that *unregisters* any surviving worker, and the repo regression-fences the removal (`code/tests/unit/next-scaffold.test.js:106,111`; `react-compiler-config.test.js:41`). **The web-app MANIFEST survives** (`public/academy-manifest.json`, `display: standalone`, declared at `code/app/layout.jsx:132`), so the app is still installable — it is simply online-only. Offline chapter bundling survives only in the Expo mobile app
```

**CITED CONTENT**

```
    38      expect(config.turbopack?.root).toBe(codeRoot)
    39      // M6-D1: Serwist must NOT wrap next.config (configurator mode runs separately)
    40      const src = readFileSync(resolve(codeRoot, 'next.config.js'), 'utf8')
    41      expect(src).not.toMatch(/from '@serwist\/next'/)
    42    })
    43  })
    44  
```

## 12-009
- **id**: `B12-009`
- **corpus site**: `corpus/architecture/service_taxonomy.md:401-411` (paragraph)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
> **⚠️ Not a local service.** Platform `2adcf71` (2026-07-31) deleted the `graphql` compose service **and** the
> `graphql-wundergraph` `repos.yml` entry; the GitHub repo was **archived 2026-07-30** (a dated snapshot — see
> the archive-state note above the *Archived / merged* table — the **⚠️ *"Two different fates shared this
> table"*** blockquote, **named, not pinned** (this said `:142`, which at M257x iter-120 was a middle line of that blockquote, not its start); the clone is consistent with it, no commit after that date). **There is no `:5050` on
> a local stack** — the frontends and studio-desk hit `backend` at `:8082/graphql/query`. The table below
> describes the router as it still exists **in production** (`graphql-wundergraph/terraform/main.tf:20` `= 1`)
> and in the archived repo; **do not follow it as a local-development instruction.** Consistent with the
> *"**There is no `graphql` profile**, and no cms / …"* sentence in the **Tier 1** deep-dive section above —
> **named, not pinned** (this said `:67-68`, which at M257x iter-120 was the *Communication* + *Database*
> characteristic bullets, not that sentence).
> Fenced source of truth: [`platform-migration-status.md`](./platform-migration-status.md).
```

**CITED CONTENT**

```
    17    tags                           = var.tags
    18    aws_region                     = var.aws_region
    19    project                        = local.project
    20    service_desired_count          = 1
    21    service_cpu                    = local.service_cpu
    22    service_memory                 = local.service_memory
    23    service_port                   = local.port
```

## 12-010
- **id**: `B12-010`
- **corpus site**: `corpus/architecture/service_taxonomy.md:416-416` (table-row)
- **citation**: `terraform/locals.tf:8`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/locals.tf`  (22 lines)

**CLAIMING UNIT**

```md
| **Port** | **8080** everywhere the router still runs — container and ECS alike (`terraform/locals.tf:8` `port = 8080`; `terraform/main.tf:48-49` maps container 8080 → host 8080; `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`). **`5050` was never a production port** — it was only the LOCAL compose host mapping `"5050:8080"`, deleted with the service at `2adcf71` |
```

**CITED CONTENT**

```
     5    }
     6    project   = "backend"
     7    port      = 8080
     8    rpc_port  = 8081
     9    meta_port = 8083
    10    # Bumped for the cms-in-app merge (v8.0): the app task now runs jobsimulation,
    11    # skiller, skillpath AND cms in-process (skiller-in-app PR #958 set the prior
```

## 12-011
- **id**: `B12-011`
- **corpus site**: `corpus/architecture/service_taxonomy.md:416-416` (table-row)
- **citation**: `terraform/main.tf:48-49`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
| **Port** | **8080** everywhere the router still runs — container and ECS alike (`terraform/locals.tf:8` `port = 8080`; `terraform/main.tf:48-49` maps container 8080 → host 8080; `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`). **`5050` was never a production port** — it was only the LOCAL compose host mapping `"5050:8080"`, deleted with the service at `2adcf71` |
```

**CITED CONTENT**

```
    45  // all on schema sentinel — no USAGE, no CREATE, no SELECT, no INSERT — so an
    46  // apply through it cannot even create the revisions table.
    47  //
    48  // What replaces it is the connection the standalone sentinel service is ALREADY
    49  // running on — the same DSN, unchanged, held by infrastructure as
    50  // var.sentinel_db_connection_euwest1 and stored at SSM
    51  // /<env>/sentinel/db_connection. The same one the in-process PDP will open its
    52  // own pool from in v11.0. "Dedicated" here means a purpose-built DSN, not a
```

## 12-012
- **id**: `B12-012`
- **corpus site**: `corpus/architecture/service_taxonomy.md:416-416` (table-row)
- **citation**: `config.prod.yaml:5`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/config.prod.yaml`  (64 lines)

**CLAIMING UNIT**

```md
| **Port** | **8080** everywhere the router still runs — container and ECS alike (`terraform/locals.tf:8` `port = 8080`; `terraform/main.tf:48-49` maps container 8080 → host 8080; `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`). **`5050` was never a production port** — it was only the LOCAL compose host mapping `"5050:8080"`, deleted with the service at `2adcf71` |
```

**CITED CONTENT**

```
     2  
     3  # Path to the previous generated file
     4  router_config_path: config.json
     5  listen_addr: 0.0.0.0:8080
     6  graph:
     7    # Result of `wgc router token create`. Can be omitted for local testing.
     8    token: ""
```

## 12-013
- **id**: `B12-013`
- **corpus site**: `corpus/architecture/service_taxonomy.md:437-437` (bullet)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Synchronous**: at platform `0c91421` the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set a single service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The last four `*_RPC_ADDR` (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_`) were `messenger`'s and all read `http://backend:8083`. ⚠️ **They were not all re-pointed by `d11a403`, and this sentence said they were until M257x iter-115.** That commit changed exactly two values on the messenger block — `CMS_RPC_ADDR` (`http://cms:8091` → `http://backend:8083`) and `JOBSIMULATION_RPC_ADDR` (`http://jobsimulation:8401` → `http://backend:8083`). At `d11a403^` the other two **already** read `http://backend:8083`, and `BACKEND_USERS_RPC_ADDR` never addressed anything else from its introduction at `3e85fce` — it only ever moved ports, so there was nothing to re-point. The end-state (*all four reach `backend`*) is true; **the agentive form is the false one**, and the clause *"— the M809 re-point landed —"* is what forced it. Root `CLAUDE.md` states the precise version (*"`d11a403` had re-pointed the **middle two**"*), so the corpus knew the distinction and this file stated it wrong. The M809 re-point did land — and `838d907` deleted the `messenger` service, taking all four with it. The env-var names survive in consumer code; no compose file configures them. The correctly-scoped model form is [`architecture_overview.md`](./architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*
```

**CITED CONTENT**

```
    45        - .env
    46      environment:
    47        - AI_USAGE_STREAM=AI
    48        - AUTHORIZATION_ADDRESS=http://sentinel:8087
    49        - AWS_CHIME_SDK_REGION=eu-central-1
    50        - CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo
    51        - CMS_STREAM=cms
```

## 12-014
- **id**: `B12-014`
- **corpus site**: `corpus/architecture/service_taxonomy.md:437-437` (bullet)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Synchronous**: at platform `0c91421` the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set a single service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The last four `*_RPC_ADDR` (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_`) were `messenger`'s and all read `http://backend:8083`. ⚠️ **They were not all re-pointed by `d11a403`, and this sentence said they were until M257x iter-115.** That commit changed exactly two values on the messenger block — `CMS_RPC_ADDR` (`http://cms:8091` → `http://backend:8083`) and `JOBSIMULATION_RPC_ADDR` (`http://jobsimulation:8401` → `http://backend:8083`). At `d11a403^` the other two **already** read `http://backend:8083`, and `BACKEND_USERS_RPC_ADDR` never addressed anything else from its introduction at `3e85fce` — it only ever moved ports, so there was nothing to re-point. The end-state (*all four reach `backend`*) is true; **the agentive form is the false one**, and the clause *"— the M809 re-point landed —"* is what forced it. Root `CLAUDE.md` states the precise version (*"`d11a403` had re-pointed the **middle two**"*), so the corpus knew the distinction and this file stated it wrong. The M809 re-point did land — and `838d907` deleted the `messenger` service, taking all four with it. The env-var names survive in consumer code; no compose file configures them. The correctly-scoped model form is [`architecture_overview.md`](./architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*
```

**CITED CONTENT**

```
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
```

## 12-015
- **id**: `B12-015`
- **corpus site**: `corpus/architecture/service_taxonomy.md:437-437` (bullet)
- **citation**: `docker-compose.yml:183`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Synchronous**: at platform `0c91421` the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set a single service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The last four `*_RPC_ADDR` (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_`) were `messenger`'s and all read `http://backend:8083`. ⚠️ **They were not all re-pointed by `d11a403`, and this sentence said they were until M257x iter-115.** That commit changed exactly two values on the messenger block — `CMS_RPC_ADDR` (`http://cms:8091` → `http://backend:8083`) and `JOBSIMULATION_RPC_ADDR` (`http://jobsimulation:8401` → `http://backend:8083`). At `d11a403^` the other two **already** read `http://backend:8083`, and `BACKEND_USERS_RPC_ADDR` never addressed anything else from its introduction at `3e85fce` — it only ever moved ports, so there was nothing to re-point. The end-state (*all four reach `backend`*) is true; **the agentive form is the false one**, and the clause *"— the M809 re-point landed —"* is what forced it. Root `CLAUDE.md` states the precise version (*"`d11a403` had re-pointed the **middle two**"*), so the corpus knew the distinction and this file stated it wrong. The M809 re-point did land — and `838d907` deleted the `messenger` service, taking all four with it. The env-var names survive in consumer code; no compose file configures them. The correctly-scoped model form is [`architecture_overview.md`](./architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*
```

**CITED CONTENT**

```
   180        - "3200:3200"
   181      networks:
   182        - app-network
   183      profiles: [core, backend, all]
   184  
   185  networks:
   186    app-network:
```

## 12-016
- **id**: `B12-016`
- **corpus site**: `corpus/architecture/service_taxonomy.md:437-437` (bullet)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
- **Synchronous**: at platform `0c91421` the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set a single service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The last four `*_RPC_ADDR` (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_`) were `messenger`'s and all read `http://backend:8083`. ⚠️ **They were not all re-pointed by `d11a403`, and this sentence said they were until M257x iter-115.** That commit changed exactly two values on the messenger block — `CMS_RPC_ADDR` (`http://cms:8091` → `http://backend:8083`) and `JOBSIMULATION_RPC_ADDR` (`http://jobsimulation:8401` → `http://backend:8083`). At `d11a403^` the other two **already** read `http://backend:8083`, and `BACKEND_USERS_RPC_ADDR` never addressed anything else from its introduction at `3e85fce` — it only ever moved ports, so there was nothing to re-point. The end-state (*all four reach `backend`*) is true; **the agentive form is the false one**, and the clause *"— the M809 re-point landed —"* is what forced it. Root `CLAUDE.md` states the precise version (*"`d11a403` had re-pointed the **middle two**"*), so the corpus knew the distinction and this file stated it wrong. The M809 re-point did land — and `838d907` deleted the `messenger` service, taking all four with it. The env-var names survive in consumer code; no compose file configures them. The correctly-scoped model form is [`architecture_overview.md`](./architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*
```

**CITED CONTENT**

```
    28  		return nil, fmt.Errorf("gotenberg: can't finalize multipart body: %w", err)
    29  	}
    30  
    31  	req, err := http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)
    32  	if err != nil {
    33  		return nil, fmt.Errorf("gotenberg: can't create request: %w", err)
    34  	}
```

## 12-017
- **id**: `B12-017`
- **corpus site**: `corpus/architecture/service_taxonomy.md:437-437` (bullet)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Synchronous**: at platform `0c91421` the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set a single service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The last four `*_RPC_ADDR` (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_`) were `messenger`'s and all read `http://backend:8083`. ⚠️ **They were not all re-pointed by `d11a403`, and this sentence said they were until M257x iter-115.** That commit changed exactly two values on the messenger block — `CMS_RPC_ADDR` (`http://cms:8091` → `http://backend:8083`) and `JOBSIMULATION_RPC_ADDR` (`http://jobsimulation:8401` → `http://backend:8083`). At `d11a403^` the other two **already** read `http://backend:8083`, and `BACKEND_USERS_RPC_ADDR` never addressed anything else from its introduction at `3e85fce` — it only ever moved ports, so there was nothing to re-point. The end-state (*all four reach `backend`*) is true; **the agentive form is the false one**, and the clause *"— the M809 re-point landed —"* is what forced it. Root `CLAUDE.md` states the precise version (*"`d11a403` had re-pointed the **middle two**"*), so the corpus knew the distinction and this file stated it wrong. The M809 re-point did land — and `838d907` deleted the `messenger` service, taking all four with it. The env-var names survive in consumer code; no compose file configures them. The correctly-scoped model form is [`architecture_overview.md`](./architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*
```

**CITED CONTENT**

```
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
    61        - LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud
    62        - LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test
```

## 12-018
- **id**: `B12-018`
- **corpus site**: `corpus/architecture/shared_libraries.md:57-57` (table-row)
- **citation**: `app/go.mod:15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
| **Version pin** | **ONE pin across the live services: `v0.35.2`.** `app` (`app/go.mod:15` @ `ad9f3c49`) and `sentinel` (`sentinel/go.mod:8` @ `f2c46190`) — the only two `type: go` entries in `repos.yml` — now agree. **The long-standing "split" is CLOSED**, and the closing event is dated: `sentinel`'s `88036d7` *"chore(deps): update dependencies to latest versions"* took it `v0.34.3 → v0.35.2`, two commits past the `88bc5592` this row used to cite. The frozen repos keep their own pins and are **not** part of this reading — `storage` `v0.34.3` (`4ce8ece5`), `messenger` `v0.35.2` (`fa47850d`), archived `chronos` `v0.30.1`; the `v0.35.1` third pin went with the `cms` + `jobsimulation` husk containers at `d11a403`. Measured from each repo's `go.mod` at the ref stated beside it. |
```

**CITED CONTENT**

```
    12  	github.com/ThreeDotsLabs/watermill-redisstream v1.4.5
    13  	github.com/anthropics/anthropic-sdk-go v1.61.0
    14  	github.com/anthropos-work/analytics-go v0.3.1
    15  	github.com/anthropos-work/colony v0.35.2
    16  	github.com/anthropos-work/proto v1.210.0
    17  	github.com/anthropos-work/storage v0.15.2
    18  	github.com/anthropos-work/taxonomy v1.2.0
```

## 12-019
- **id**: `B12-019`
- **corpus site**: `corpus/architecture/shared_libraries.md:57-57` (table-row)
- **citation**: `sentinel/go.mod:8`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/go.mod`  (54 lines)

**CLAIMING UNIT**

```md
| **Version pin** | **ONE pin across the live services: `v0.35.2`.** `app` (`app/go.mod:15` @ `ad9f3c49`) and `sentinel` (`sentinel/go.mod:8` @ `f2c46190`) — the only two `type: go` entries in `repos.yml` — now agree. **The long-standing "split" is CLOSED**, and the closing event is dated: `sentinel`'s `88036d7` *"chore(deps): update dependencies to latest versions"* took it `v0.34.3 → v0.35.2`, two commits past the `88bc5592` this row used to cite. The frozen repos keep their own pins and are **not** part of this reading — `storage` `v0.34.3` (`4ce8ece5`), `messenger` `v0.35.2` (`fa47850d`), archived `chronos` `v0.30.1`; the `v0.35.1` third pin went with the `cms` + `jobsimulation` husk containers at `d11a403`. Measured from each repo's `go.mod` at the ref stated beside it. |
```

**CITED CONTENT**

```
     5  require (
     6  	connectrpc.com/connect v1.20.0
     7  	github.com/Blank-Xu/sql-adapter v1.2.1
     8  	github.com/anthropos-work/colony v0.35.2
     9  	github.com/anthropos-work/proto v1.210.0
    10  	github.com/casbin/casbin/v3 v3.10.0
    11  	github.com/google/uuid v1.6.0
```

## 12-020
- **id**: `B12-020`
- **corpus site**: `corpus/architecture/shared_libraries.md:58-58` (table-row)
- **citation**: `roadrunner/main.go:7`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/roadrunner/main.go`  (19 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Every live Go service — `app` and `sentinel`, and only those two, at platform `0c91421`.** `repos.yml` there lists four entries (app, sentinel, next-web-app, studio-desk) of which two are `type: go`. The four-service reading (app, sentinel, **storage**, **messenger**) was true at `0dab54d` and **`838d907` ended it**: that commit deleted the `storage` and `messenger` clone entries *and* their compose services, so `make init` now clones app, sentinel, next-web-app, studio-desk. Both repos still import colony in their own `go.mod` (storage `v0.34.3` @ `4ce8ece5`, messenger `v0.35.2` @ `fa47850d`) but nothing clones or builds them — they join the `cms`, `jobsimulation` and `roadrunner` repos, **gone from compose** at `0dab54d` (`d11a403`); there is no profile that starts them and no `graphql` profile at all. Their domains run inside `app`, and the three repos are **frozen legacy** — still on GitHub as the pre-merge reference, but with no compose service, no `repos.yml` entry and nothing that starts them, so they are not importers of anything a stack runs. The `roadrunner` repo's own import was minimal but real while it lasted: `roadrunner/main.go:7` imports `colony` for `NewVersionConfig` (`roadrunner/go.mod:7` pins `v0.34.3`) |
```

**CITED CONTENT**

```
     4  	"cmp"
     5  	"os"
     6  
     7  	"github.com/anthropos-work/colony"
     8  	"github.com/anthropos-work/roadrunner/cmd"
     9  	_ "github.com/joho/godotenv/autoload"
    10  )
```

## 12-021
- **id**: `B12-021`
- **corpus site**: `corpus/architecture/shared_libraries.md:58-58` (table-row)
- **citation**: `roadrunner/go.mod:7`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/roadrunner/go.mod`  (66 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Every live Go service — `app` and `sentinel`, and only those two, at platform `0c91421`.** `repos.yml` there lists four entries (app, sentinel, next-web-app, studio-desk) of which two are `type: go`. The four-service reading (app, sentinel, **storage**, **messenger**) was true at `0dab54d` and **`838d907` ended it**: that commit deleted the `storage` and `messenger` clone entries *and* their compose services, so `make init` now clones app, sentinel, next-web-app, studio-desk. Both repos still import colony in their own `go.mod` (storage `v0.34.3` @ `4ce8ece5`, messenger `v0.35.2` @ `fa47850d`) but nothing clones or builds them — they join the `cms`, `jobsimulation` and `roadrunner` repos, **gone from compose** at `0dab54d` (`d11a403`); there is no profile that starts them and no `graphql` profile at all. Their domains run inside `app`, and the three repos are **frozen legacy** — still on GitHub as the pre-merge reference, but with no compose service, no `repos.yml` entry and nothing that starts them, so they are not importers of anything a stack runs. The `roadrunner` repo's own import was minimal but real while it lasted: `roadrunner/main.go:7` imports `colony` for `NewVersionConfig` (`roadrunner/go.mod:7` pins `v0.34.3`) |
```

**CITED CONTENT**

```
     4  
     5  require (
     6  	connectrpc.com/connect v1.19.2
     7  	github.com/anthropos-work/colony v0.34.3
     8  	github.com/anthropos-work/proto v1.196.0
     9  	github.com/gorilla/websocket v1.5.3
    10  	github.com/hibiken/asynq v0.25.1
```

## 12-022
- **id**: `B12-022`
- **corpus site**: `corpus/architecture/shared_libraries.md:85-85` (table-row)
- **citation**: `app/go.mod:16`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
| **Version pin** | **ONE value across the live services: `v1.210.0`.** `app` (`app/go.mod:16` @ `ad9f3c49`) and `sentinel` (`sentinel/go.mod:9` @ `f2c46190`) agree, so **the live skew is ZERO, not two** — closed by the same `sentinel` `88036d7` dependency bump that closed the colony split (`v1.200.0 → v1.210.0`). At platform `0dab54d` the reading was four repos (app/messenger `v1.210.0`, sentinel `v1.200.0`, storage `v1.196.0`); `838d907` deleted the `storage` and `messenger` clone entries, so `storage v1.196.0` (`4ce8ece5`) and `messenger v1.210.0` (`fa47850d`) are frozen alongside the husks below and are not part of the live reading. The frozen repos keep their own `go.mod` and therefore their own pins (cms `v1.207.0` @ `ca50c817`, jobsimulation `v1.205.0` @ `462343b0`, roadrunner `v1.196.0` @ `87d8d44`), but **nothing clones or builds them any more**: `d11a403` deleted all three `repos.yml` entries *and* their compose services in one commit, so those pins compile nowhere. Reading them as part of the platform's skew was the error this row used to make |
```

**CITED CONTENT**

```
    13  	github.com/anthropics/anthropic-sdk-go v1.61.0
    14  	github.com/anthropos-work/analytics-go v0.3.1
    15  	github.com/anthropos-work/colony v0.35.2
    16  	github.com/anthropos-work/proto v1.210.0
    17  	github.com/anthropos-work/storage v0.15.2
    18  	github.com/anthropos-work/taxonomy v1.2.0
    19  	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
```

## 12-023
- **id**: `B12-023`
- **corpus site**: `corpus/architecture/shared_libraries.md:85-85` (table-row)
- **citation**: `sentinel/go.mod:9`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/go.mod`  (54 lines)

**CLAIMING UNIT**

```md
| **Version pin** | **ONE value across the live services: `v1.210.0`.** `app` (`app/go.mod:16` @ `ad9f3c49`) and `sentinel` (`sentinel/go.mod:9` @ `f2c46190`) agree, so **the live skew is ZERO, not two** — closed by the same `sentinel` `88036d7` dependency bump that closed the colony split (`v1.200.0 → v1.210.0`). At platform `0dab54d` the reading was four repos (app/messenger `v1.210.0`, sentinel `v1.200.0`, storage `v1.196.0`); `838d907` deleted the `storage` and `messenger` clone entries, so `storage v1.196.0` (`4ce8ece5`) and `messenger v1.210.0` (`fa47850d`) are frozen alongside the husks below and are not part of the live reading. The frozen repos keep their own `go.mod` and therefore their own pins (cms `v1.207.0` @ `ca50c817`, jobsimulation `v1.205.0` @ `462343b0`, roadrunner `v1.196.0` @ `87d8d44`), but **nothing clones or builds them any more**: `d11a403` deleted all three `repos.yml` entries *and* their compose services in one commit, so those pins compile nowhere. Reading them as part of the platform's skew was the error this row used to make |
```

**CITED CONTENT**

```
     6  	connectrpc.com/connect v1.20.0
     7  	github.com/Blank-Xu/sql-adapter v1.2.1
     8  	github.com/anthropos-work/colony v0.35.2
     9  	github.com/anthropos-work/proto v1.210.0
    10  	github.com/casbin/casbin/v3 v3.10.0
    11  	github.com/google/uuid v1.6.0
    12  	github.com/joho/godotenv v1.5.1
```

## 12-024
- **id**: `B12-024`
- **corpus site**: `corpus/architecture/shared_libraries.md:86-86` (table-row)
- **citation**: `app/main.go:1187`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |
```

**CITED CONTENT**

```
  1184  		JobSimulationClient: jobsimDj.RPCServer,
  1185  		Studio:              cmsManagers.Studio,
  1186  		Asynq:               cmsAsynq,
  1187  		Pub:                 cmsPub,
  1188  		Storage:             cmsStorage,
  1189  		AiVideo:             cmsManagers.AiVideo,
  1190  	}
```

## 12-025
- **id**: `B12-025`
- **corpus site**: `corpus/architecture/shared_libraries.md:86-86` (table-row)
- **citation**: `app/main.go:1188`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |
```

**CITED CONTENT**

```
  1185  		Studio:              cmsManagers.Studio,
  1186  		Asynq:               cmsAsynq,
  1187  		Pub:                 cmsPub,
  1188  		Storage:             cmsStorage,
  1189  		AiVideo:             cmsManagers.AiVideo,
  1190  	}
  1191  	// cms-in-app M807/M809: one cms RPC server backs BOTH the served CMSService handler
```

## 12-026
- **id**: `B12-026`
- **corpus site**: `corpus/architecture/shared_libraries.md:86-86` (table-row)
- **citation**: `app/main.go:1196`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |
```

**CITED CONTENT**

```
  1193  	// caller cutover). It satisfies both the connect Handler and Client interfaces.
  1194  	cmsRPCServer = cmsrpcsrv.NewRPCServer(cmsManagers.Directus, cmsManagers.AiVideo)
  1195  	// cms-in-app: the inbound Directus webhook now lands on APP (POST /api/webhook/directus) —
  1196  	// at release the Directus Flows re-point here so cache-clear / re-index / clone / ai-video
  1197  	// creation fire in-process (no traffic to the standalone cms). Authenticated by
  1198  	// DIRECTUS_WEBHOOK_SECRET (M809b M-2, fail-closed).
  1199  	cmsWebhookHandler = cmswebhooks.Handler(os.Getenv("DIRECTUS_WEBHOOK_SECRET"), cmsManagers.Directus, cmsPub, cmsManagers.AiVideo)
```

## 12-027
- **id**: `B12-027`
- **corpus site**: `corpus/architecture/shared_libraries.md:86-86` (table-row)
- **citation**: `app/main.go:1204`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |
```

**CITED CONTENT**

```
  1201  	// read cms via the in-process RPC server instead of over the wire — no traffic to the
  1202  	// standalone cms. Active whenever the Directus edge is configured (the release sets it);
  1203  	// the external client the switch was seeded with is only the construction-time placeholder.
  1204  	cmsReaderSw.set(cmsRPCServer)
  1205  	// M805: consume the cms studio + ai_video Asynq queue in-process (the app is the sole
  1206  	// consumer post-release — the standalone cms takes no traffic). The consumer polls the SAME
  1207  	// DB index the enqueue client writes to (audit R2). The studio gen.py/postgen.py pipeline
```

## 12-028
- **id**: `B12-028`
- **corpus site**: `corpus/architecture/shared_libraries.md:86-86` (table-row)
- **citation**: `app/main.go:1228`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |
```

**CITED CONTENT**

```
  1225  		intelligenceManager,
  1226  		userManager,
  1227  		academyContentManager,
  1228  		aiReadinessManager,
  1229  		profileHistoryManager,
  1230  		subscriptionManager,
  1231  		pub,
```

## 12-029
- **id**: `B12-029`
- **corpus site**: `corpus/architecture/shared_libraries.md:86-86` (table-row)
- **citation**: `app/main.go:1212-1214`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |
```

**CITED CONTENT**

```
  1209  	cmsWorker := cmsworker.NewServer(redisAddr, cmsWorkerIndex, logger)
  1210  	wg.Go(func() {
  1211  		defer cancelServerContext()
  1212  		if err := cmsWorker.Start(serverContext, cmsManagers.Studio, cmsManagers.AiVideo); err != nil {
  1213  			logger.Info("shutting down the cms worker", "error", err)
  1214  		}
  1215  	})
  1216  
  1217  	graphHandler := graph.NewHandler(
```

## 12-030
- **id**: `B12-030`
- **corpus site**: `corpus/architecture/shared_libraries.md:93-104` (paragraph)
- **citation**: `storage.md:129`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/storage.md`  (249 lines)

**CLAIMING UNIT**

```md
**At least 13** Connect-RPC services are defined — **this is a floor, not a count**: `proto` is a private Go
module and **is in no clone set**, so the list below is hand-enumerated from consumers and cannot be
verified against the source of truth. It omitted `StorageService` until M257x iter-98, which
[`storage.md:129`](../services/storage.md) documents in full. The named ones:
`UsersService`, `OrganizationsService`,
`CMSService`, `JobSimulationService`, `SkillerService` (all served by app since the merges — one RPC mux),
`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. skillpath-in-app M506 *removed* the RPC rather than re-hosting it; `app/internal/skillpaths/skillpaths.go:27-31` calls its replacement "the drop-in for the **removed** skillpath RPC client". Likewise roadrunner: `backend` calls Judge0 over plain HTTP — `jsrunner.NewRunnerManager` at `app/internal/jobsimwiring/wiring.go:123` @ `app` `9d00a313` v1.367.0 — and `ROADRUNNER_RPC_ADDR` is read by no Go code in `app` **and is not in the platform compose at all** — 0 occurrences in either, at `app` `9d00a313` / platform `0dab54d`. (This line long cited `docker-compose.yml:118` for it; that line sets `AWS_REGION`.)),
`LabSessionService` (served by app — the AI Labs domain, see `../services/ai-labs.md`),
`AuthorizationService` (Sentinel), `MessengerService`, `RoadRunnerService`,
`RealtimeService`, `ChronosService` (archived service, contract still present), and
`StorageService` (the storage surface — served in-process by `app` since the v9.0 fold). Plus
`events`/`flags`/`ai` message-only protos used over Redis Streams pub/sub.
```

**CITED CONTENT**

```
   126  
   127  ## Interface Discovery
   128  
   129  ### Connect-RPC (`StorageService`)
   130  
   131  Private:
   132  
```

## 12-031
- **id**: `B12-031`
- **corpus site**: `corpus/architecture/shared_libraries.md:93-104` (paragraph)
- **citation**: `app/internal/skillpaths/skillpaths.go:27-31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skillpaths/skillpaths.go`  (204 lines)

**CLAIMING UNIT**

```md
**At least 13** Connect-RPC services are defined — **this is a floor, not a count**: `proto` is a private Go
module and **is in no clone set**, so the list below is hand-enumerated from consumers and cannot be
verified against the source of truth. It omitted `StorageService` until M257x iter-98, which
[`storage.md:129`](../services/storage.md) documents in full. The named ones:
`UsersService`, `OrganizationsService`,
`CMSService`, `JobSimulationService`, `SkillerService` (all served by app since the merges — one RPC mux),
`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. skillpath-in-app M506 *removed* the RPC rather than re-hosting it; `app/internal/skillpaths/skillpaths.go:27-31` calls its replacement "the drop-in for the **removed** skillpath RPC client". Likewise roadrunner: `backend` calls Judge0 over plain HTTP — `jsrunner.NewRunnerManager` at `app/internal/jobsimwiring/wiring.go:123` @ `app` `9d00a313` v1.367.0 — and `ROADRUNNER_RPC_ADDR` is read by no Go code in `app` **and is not in the platform compose at all** — 0 occurrences in either, at `app` `9d00a313` / platform `0dab54d`. (This line long cited `docker-compose.yml:118` for it; that line sets `AWS_REGION`.)),
`LabSessionService` (served by app — the AI Labs domain, see `../services/ai-labs.md`),
`AuthorizationService` (Sentinel), `MessengerService`, `RoadRunnerService`,
`RealtimeService`, `ChronosService` (archived service, contract still present), and
`StorageService` (the storage surface — served in-process by `app` since the v9.0 fold). Plus
`events`/`flags`/`ai` message-only protos used over Redis Streams pub/sub.
```

**CITED CONTENT**

```
    24  	"github.com/google/uuid"
    25  )
    26  
    27  // sessionReader is the narrow read the loopback consumer needs from the in-process
    28  // skillpath SessionManager — the drop-in for the removed skillpath RPC client
    29  // (skillpath-in-app M506). *skillpath.SessionManager satisfies it.
    30  type sessionReader interface {
    31  	GetSessionDomainByID(ctx context.Context, id uuid.UUID) (*skillpathsession.SkillPathSession, error)
    32  }
    33  
    34  type SkillPathManager struct {
```

## 12-032
- **id**: `B12-032`
- **corpus site**: `corpus/architecture/shared_libraries.md:93-104` (paragraph)
- **citation**: `app/internal/jobsimwiring/wiring.go:123`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
**At least 13** Connect-RPC services are defined — **this is a floor, not a count**: `proto` is a private Go
module and **is in no clone set**, so the list below is hand-enumerated from consumers and cannot be
verified against the source of truth. It omitted `StorageService` until M257x iter-98, which
[`storage.md:129`](../services/storage.md) documents in full. The named ones:
`UsersService`, `OrganizationsService`,
`CMSService`, `JobSimulationService`, `SkillerService` (all served by app since the merges — one RPC mux),
`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. skillpath-in-app M506 *removed* the RPC rather than re-hosting it; `app/internal/skillpaths/skillpaths.go:27-31` calls its replacement "the drop-in for the **removed** skillpath RPC client". Likewise roadrunner: `backend` calls Judge0 over plain HTTP — `jsrunner.NewRunnerManager` at `app/internal/jobsimwiring/wiring.go:123` @ `app` `9d00a313` v1.367.0 — and `ROADRUNNER_RPC_ADDR` is read by no Go code in `app` **and is not in the platform compose at all** — 0 occurrences in either, at `app` `9d00a313` / platform `0dab54d`. (This line long cited `docker-compose.yml:118` for it; that line sets `AWS_REGION`.)),
`LabSessionService` (served by app — the AI Labs domain, see `../services/ai-labs.md`),
`AuthorizationService` (Sentinel), `MessengerService`, `RoadRunnerService`,
`RealtimeService`, `ChronosService` (archived service, contract still present), and
`StorageService` (the storage surface — served in-process by `app` since the v9.0 fold). Plus
`events`/`flags`/`ai` message-only protos used over Redis Streams pub/sub.
```

**CITED CONTENT**

```
   120  	storageV1Client := appstorage.NewClient(inAppStorage, storagens.JobSimulation).V1
   121  	// Judge0 sandbox runner (IN-PROCESS; replaces the removed roadrunner RPC edge — resync to jobsim main
   122  	// v0.253.0, which deleted chronos + realtime + the roadrunner-submission event).
   123  	runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))
   124  
   125  	// --- Asynq producer client (task-type / queue name strings are frozen — M705 contract).
   126  	workerIndex, _ := strconv.Atoi(getenv("REDIS_WORKER_INDEX"))
```

## 12-033
- **id**: `B12-033`
- **corpus site**: `corpus/architecture/shared_libraries.md:93-104` (paragraph)
- **citation**: `docker-compose.yml:118`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
**At least 13** Connect-RPC services are defined — **this is a floor, not a count**: `proto` is a private Go
module and **is in no clone set**, so the list below is hand-enumerated from consumers and cannot be
verified against the source of truth. It omitted `StorageService` until M257x iter-98, which
[`storage.md:129`](../services/storage.md) documents in full. The named ones:
`UsersService`, `OrganizationsService`,
`CMSService`, `JobSimulationService`, `SkillerService` (all served by app since the merges — one RPC mux),
`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. skillpath-in-app M506 *removed* the RPC rather than re-hosting it; `app/internal/skillpaths/skillpaths.go:27-31` calls its replacement "the drop-in for the **removed** skillpath RPC client". Likewise roadrunner: `backend` calls Judge0 over plain HTTP — `jsrunner.NewRunnerManager` at `app/internal/jobsimwiring/wiring.go:123` @ `app` `9d00a313` v1.367.0 — and `ROADRUNNER_RPC_ADDR` is read by no Go code in `app` **and is not in the platform compose at all** — 0 occurrences in either, at `app` `9d00a313` / platform `0dab54d`. (This line long cited `docker-compose.yml:118` for it; that line sets `AWS_REGION`.)),
`LabSessionService` (served by app — the AI Labs domain, see `../services/ai-labs.md`),
`AuthorizationService` (Sentinel), `MessengerService`, `RoadRunnerService`,
`RealtimeService`, `ChronosService` (archived service, contract still present), and
`StorageService` (the storage surface — served in-process by `app` since the v9.0 fold). Plus
`events`/`flags`/`ai` message-only protos used over Redis Streams pub/sub.
```

**CITED CONTENT**

```
   115        dockerfile: Dockerfile.dev
   116        ssh: ["default"]
   117        args:
   118          VITE_CLERK_PUBLISHABLE_KEY: ${VITE_CLERK_PUBLISHABLE_KEY}
   119          VITE_GRAPHQL_ENDPOINT: ${VITE_GRAPHQL_ENDPOINT:-http://localhost:8082/graphql/query}
   120          VITE_ENVIRONMENT: ${VITE_ENVIRONMENT:-production}
   121          VERSION: dev
```

## 12-034
- **id**: `B12-034`
- **corpus site**: `corpus/architecture/shared_libraries.md:126-126` (table-row)
- **citation**: `cms/go.mod:9`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/go.mod`  (127 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **No repo a stack builds** (corrected M257x iter-102; this row previously said *"`app` alone among the services a stack runs"*). `app` **dropped** the module at `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*): `git show ad9f3c49:go.mod` has no `anthropos-work/ai` line and `go.sum` has zero, while `app/internal/ai/` carries the library in-tree — with a one-way door, `internal/ai/module_import_guard_test.go`, whose own comment records that the repo *"was deliberately left in place because at least one consumer outside this codebase (anthropos-work/rosetta-extensions/stack-seeding) pins it."* `sentinel` never required it. The frozen `cms` and `jobsimulation` repos still require it directly (`cms/go.mod:9` @ `ca50c817`, `jobsimulation/go.mod:11` @ `462343b0`, both `v1.40.2`), but neither has a compose service or a `repos.yml` entry, so nothing builds or starts them — a `go.mod` require in a repo nothing compiles is not a live import. Go services only — **not** Studio-Desk, which is TypeScript, and **not** roadrunner, whose only shared-lib requires are colony + proto |
```

**CITED CONTENT**

```
     6  	connectrpc.com/connect v1.20.0
     7  	entgo.io/ent v0.14.6
     8  	github.com/99designs/gqlgen v0.17.94
     9  	github.com/anthropos-work/ai v1.40.2
    10  	github.com/anthropos-work/colony v0.35.1
    11  	github.com/anthropos-work/proto v1.207.0
    12  	github.com/anthropos-work/storage v0.15.2
```

## 12-035
- **id**: `B12-035`
- **corpus site**: `corpus/architecture/shared_libraries.md:126-126` (table-row)
- **citation**: `jobsimulation/go.mod:11`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/go.mod`  (221 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **No repo a stack builds** (corrected M257x iter-102; this row previously said *"`app` alone among the services a stack runs"*). `app` **dropped** the module at `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*): `git show ad9f3c49:go.mod` has no `anthropos-work/ai` line and `go.sum` has zero, while `app/internal/ai/` carries the library in-tree — with a one-way door, `internal/ai/module_import_guard_test.go`, whose own comment records that the repo *"was deliberately left in place because at least one consumer outside this codebase (anthropos-work/rosetta-extensions/stack-seeding) pins it."* `sentinel` never required it. The frozen `cms` and `jobsimulation` repos still require it directly (`cms/go.mod:9` @ `ca50c817`, `jobsimulation/go.mod:11` @ `462343b0`, both `v1.40.2`), but neither has a compose service or a `repos.yml` entry, so nothing builds or starts them — a `go.mod` require in a repo nothing compiles is not a live import. Go services only — **not** Studio-Desk, which is TypeScript, and **not** roadrunner, whose only shared-lib requires are colony + proto |
```

**CITED CONTENT**

```
     8  	entgo.io/ent v0.14.6
     9  	github.com/99designs/gqlgen v0.17.92
    10  	github.com/anthropics/anthropic-sdk-go v1.51.1
    11  	github.com/anthropos-work/ai v1.40.2
    12  	github.com/anthropos-work/colony v0.35.1
    13  	github.com/anthropos-work/proto v1.205.0
    14  	github.com/anthropos-work/storage v0.15.2
```

## 12-036
- **id**: `B12-036`
- **corpus site**: `corpus/architecture/shared_libraries.md:213-213` (table-row)
- **citation**: `app/go.mod:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |
```

**CITED CONTENT**

```
    17  	github.com/anthropos-work/storage v0.15.2
    18  	github.com/anthropos-work/taxonomy v1.2.0
    19  	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
    20  	github.com/asticode/go-astisub v0.42.0
    21  	github.com/avast/retry-go/v4 v4.7.0
    22  	github.com/aws/aws-sdk-go-v2 v1.43.0
    23  	github.com/aws/aws-sdk-go-v2/config v1.32.31
```

## 12-037
- **id**: `B12-037`
- **corpus site**: `corpus/architecture/shared_libraries.md:213-213` (table-row)
- **citation**: `sentinel/go.mod:21`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/go.mod`  (54 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |
```

**CITED CONTENT**

```
    18  require (
    19  	github.com/99designs/gqlgen v0.17.94 // indirect
    20  	github.com/agnivade/levenshtein v1.2.1 // indirect
    21  	github.com/anthropos-work/taxonomy v1.2.0 // indirect
    22  	github.com/bmatcuk/doublestar/v4 v4.10.0 // indirect
    23  	github.com/brunoscheufler/aws-ecs-metadata-go v0.0.0-20221221133751-67e37ae746cd // indirect
    24  	github.com/casbin/govaluate v1.10.0 // indirect
```

## 12-038
- **id**: `B12-038`
- **corpus site**: `corpus/architecture/shared_libraries.md:213-213` (table-row)
- **citation**: `storage/go.mod:25`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/storage/go.mod`  (71 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |
```

**CITED CONTENT**

```
    22  require (
    23  	github.com/99designs/gqlgen v0.17.90 // indirect
    24  	github.com/agnivade/levenshtein v1.2.1 // indirect
    25  	github.com/anthropos-work/taxonomy v1.2.0 // indirect
    26  	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.4 // indirect
    27  	github.com/aws/aws-sdk-go-v2/credentials v1.19.9 // indirect
    28  	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.17 // indirect
```

## 12-039
- **id**: `B12-039`
- **corpus site**: `corpus/architecture/shared_libraries.md:213-213` (table-row)
- **citation**: `messenger/go.mod:9`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/go.mod`  (77 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |
```

**CITED CONTENT**

```
     6  	connectrpc.com/connect v1.20.0
     7  	github.com/anthropos-work/colony v0.35.2
     8  	github.com/anthropos-work/proto v1.210.0
     9  	github.com/anthropos-work/taxonomy v1.2.0
    10  	github.com/dustin/go-humanize v1.0.1
    11  	github.com/getbrevo/brevo-go v1.1.3
    12  	github.com/google/uuid v1.6.0
```

## 12-040
- **id**: `B12-040`
- **corpus site**: `corpus/architecture/shared_libraries.md:213-213` (table-row)
- **citation**: `cms/go.mod:13`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/go.mod`  (127 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |
```

**CITED CONTENT**

```
    10  	github.com/anthropos-work/colony v0.35.1
    11  	github.com/anthropos-work/proto v1.207.0
    12  	github.com/anthropos-work/storage v0.15.2
    13  	github.com/anthropos-work/taxonomy v1.2.0
    14  	github.com/asticode/go-astisub v0.42.0
    15  	github.com/gabriel-vasile/mimetype v1.4.13
    16  	github.com/go-playground/validator/v10 v10.30.3
```

## 12-041
- **id**: `B12-041`
- **corpus site**: `corpus/architecture/shared_libraries.md:213-213` (table-row)
- **citation**: `jobsimulation/go.mod:15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/go.mod`  (221 lines)

**CLAIMING UNIT**

```md
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |
```

**CITED CONTENT**

```
    12  	github.com/anthropos-work/colony v0.35.1
    13  	github.com/anthropos-work/proto v1.205.0
    14  	github.com/anthropos-work/storage v0.15.2
    15  	github.com/anthropos-work/taxonomy v1.2.0
    16  	github.com/avast/retry-go/v4 v4.7.0
    17  	github.com/aws/aws-sdk-go-v2 v1.42.0
    18  	github.com/aws/aws-sdk-go-v2/config v1.32.25
```
