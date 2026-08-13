# TIER-1 ADJUDICATION BATCH 09 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 09-001
- **id**: `B09-001`
- **corpus site**: `corpus/architecture/dependency_map.md:61-61` (table-row)
- **citation**: `messenger/internal/flow/flow.go:109`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/flow.go`  (135 lines)

**CLAIMING UNIT**

```md
| `cms` | App (+ **Directus webhooks** → `POST /api/webhook/directus`, now authenticated via `DIRECTUS_WEBHOOK_SECRET`) | App (the cms similarity/Studio handlers + the jobsim handlers, merged onto ONE subscriber), the standalone Messenger (`messenger/internal/flow/flow.go:109` — no longer startable since `838d907`), **and app's messenger-in-app subscriber**: the takeover names all three of messenger's streams, this one included — and is **opt-in behind `MESSENGER_ENABLED`, OFF by default**, same as the `backend` row | Content published/updated, translation & clone requests |
```

**CITED CONTENT**

```
   106  		pubsub.EventHandler(h.JobsimulationSessionStartedHandler),
   107  		pubsub.EventHandler(h.JobsimulationSessionEndedHandler),
   108  	))
   109  	h.subServer.AddSubscriber("cms", sub.AddHandler(
   110  		pubsub.EventHandler(h.CmsStudioTaskJobSimulationCompletedHandler),
   111  	))
   112  }
```

## 09-002
- **id**: `B09-002`
- **corpus site**: `corpus/architecture/external_services.md:3-3` (paragraph)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
> **⚠️ Router status, two states (v2.8 M257x).** Platform `b56d731`+`360efd4` (merged **`2adcf71`**, 2026-07-31) **deleted the Cosmo Router from local dev** — no `graphql` compose service, no `repos.yml` entry — and re-pointed the frontends at **`backend` directly, `http://localhost:8082/graphql/query`**. **There is no `:5050` on a local stack.** In *production* the router is still declared (`graphql-wundergraph/terraform/main.tf:20` `= 1`), though **the repo is ARCHIVED on GitHub (2026-07-30)**. And the supergraph is **ONE** subgraph — `backend` — since `915da06` (2026-07-29). The fenced source of truth is [`platform-migration-status.md`](./platform-migration-status.md).
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

## 09-003
- **id**: `B09-003`
- **corpus site**: `corpus/architecture/external_services.md:129-155` (paragraph)
- **citation**: `docker-compose.yml:53`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **The platform `docker-compose.yml` has NO directus service.** A local stack does not run Directus — the cms
> domain in `backend` reaches Directus over the network via `DIRECTUS_BASE_ADDR` / `DIRECTUS_PUBLIC_BASE_ADDR`,
> which point at the **production** instance `https://content.anthropos.work` in the stock compose.
> **⚠️ `backend`'s compose `environment:` block sets exactly ONE of the pair.** At platform `0dab54d` it sets
> `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work` (`docker-compose.yml:53`) and **no
> `DIRECTUS_BASE_ADDR`** — that one `backend` picks up from the shared `env_file: .env`. The standalone `cms`
> service that used to carry both explicitly is **gone from compose** (`d11a403`), so `backend` is the only
> **live** consumer left. **But the re-point tooling declares TWO targets, not one** — `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` (`rosetta-extensions/stack-injection/gen_injected_override.py:86`, and identically in the dev twin `stack-core/gen_override.py:58`, both @ the demo's pinned rext `09d06070`). Only `backend` ever *matches* on a current clone; `cms` is retained deliberately as an **inert key**, and the source says so at `:77-81` — the test *"never matches it on a current clone"* and it is *"kept only so a ROLLBACK/older platform clone that still DEFINES the container gets re-pointed too."* The ⚠️ under *Architecture* below (`:206-211`) states that two-member tuple **in bold**; it **qualifies** this sentence rather than corroborating a one-target reading, and an earlier revision here cited it as corroboration while it said the opposite (booked M257x iter-101, repaired iter-102).
> A freshly-built local stack reads its public content **live from prod**. (Earlier revisions of this doc described a
> `directus/directus:10.10.1` compose service on port 8055 with an `admin@example.com` / `password` admin login
> and an inline `docker-compose.yml` snippet **as if it were CURRENT**, which it is not — there is still no
> Directus service in the platform compose at `0dab54d`.)
>
> **That retraction over-corrected, and this corrects the correction (M257x iter-48).** The twin of this
> paragraph said *"all of that is false; that service **has never existed**"* — repaired at
> [`service_taxonomy.md:350-357`](./service_taxonomy.md) and left standing here. The service **did** exist,
> with exactly that image tag, port and passwo
```

**CITED CONTENT**

```
    50        - CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo
    51        - CMS_STREAM=cms
    52        - COMPANY_LOGO_STORAGE_PATH=company_logos
    53        - DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
```

## 09-004
- **id**: `B09-004`
- **corpus site**: `corpus/architecture/external_services.md:129-155` (paragraph)
- **citation**: `rosetta-extensions/stack-injection/gen_injected_override.py:86`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-injection/gen_injected_override.py`  (892 lines)

**CLAIMING UNIT**

```md
> **The platform `docker-compose.yml` has NO directus service.** A local stack does not run Directus — the cms
> domain in `backend` reaches Directus over the network via `DIRECTUS_BASE_ADDR` / `DIRECTUS_PUBLIC_BASE_ADDR`,
> which point at the **production** instance `https://content.anthropos.work` in the stock compose.
> **⚠️ `backend`'s compose `environment:` block sets exactly ONE of the pair.** At platform `0dab54d` it sets
> `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work` (`docker-compose.yml:53`) and **no
> `DIRECTUS_BASE_ADDR`** — that one `backend` picks up from the shared `env_file: .env`. The standalone `cms`
> service that used to carry both explicitly is **gone from compose** (`d11a403`), so `backend` is the only
> **live** consumer left. **But the re-point tooling declares TWO targets, not one** — `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` (`rosetta-extensions/stack-injection/gen_injected_override.py:86`, and identically in the dev twin `stack-core/gen_override.py:58`, both @ the demo's pinned rext `09d06070`). Only `backend` ever *matches* on a current clone; `cms` is retained deliberately as an **inert key**, and the source says so at `:77-81` — the test *"never matches it on a current clone"* and it is *"kept only so a ROLLBACK/older platform clone that still DEFINES the container gets re-pointed too."* The ⚠️ under *Architecture* below (`:206-211`) states that two-member tuple **in bold**; it **qualifies** this sentence rather than corroborating a one-target reading, and an earlier revision here cited it as corroboration while it said the opposite (booked M257x iter-101, repaired iter-102).
> A freshly-built local stack reads its public content **live from prod**. (Earlier revisions of this doc described a
> `directus/directus:10.10.1` compose service on port 8055 with an `admin@example.com` / `password` admin login
> and an inline `docker-compose.yml` snippet **as if it were CURRENT**, which it is not — there is still no
> Directus service in the platform compose at `0dab54d`.)
>
> **That retraction over-corrected, and this corrects the correction (M257x iter-48).** The twin of this
> paragraph said *"all of that is false; that service **has never existed**"* — repaired at
> [`service_taxonomy.md:350-357`](./service_taxonomy.md) and left standing here. The service **did** exist,
> with exactly that image tag, port and passwo
```

**CITED CONTENT**

```
    83  # studio-desk is wired separately (its own DIRECTUS_BASE_URL — see the FRONTENDS table). The ASSET plane
    84  # (DIRECTUS_PUBLIC_BASE_ADDR) is deliberately NOT re-pointed — it stays on prod public links so
    85  # browser images stay real (the v1.5 data-plane-local / asset-plane-prod split).
    86  DIRECTUS_DATA_CONSUMERS = ("cms", "backend")
    87  
    88  
    89  def skiller_ai_fallback_env():
```

## 09-005
- **id**: `B09-005`
- **corpus site**: `corpus/architecture/external_services.md:129-155` (paragraph)
- **citation**: `stack-core/gen_override.py:58`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-core/gen_override.py`  (286 lines)

**CLAIMING UNIT**

```md
> **The platform `docker-compose.yml` has NO directus service.** A local stack does not run Directus — the cms
> domain in `backend` reaches Directus over the network via `DIRECTUS_BASE_ADDR` / `DIRECTUS_PUBLIC_BASE_ADDR`,
> which point at the **production** instance `https://content.anthropos.work` in the stock compose.
> **⚠️ `backend`'s compose `environment:` block sets exactly ONE of the pair.** At platform `0dab54d` it sets
> `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work` (`docker-compose.yml:53`) and **no
> `DIRECTUS_BASE_ADDR`** — that one `backend` picks up from the shared `env_file: .env`. The standalone `cms`
> service that used to carry both explicitly is **gone from compose** (`d11a403`), so `backend` is the only
> **live** consumer left. **But the re-point tooling declares TWO targets, not one** — `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` (`rosetta-extensions/stack-injection/gen_injected_override.py:86`, and identically in the dev twin `stack-core/gen_override.py:58`, both @ the demo's pinned rext `09d06070`). Only `backend` ever *matches* on a current clone; `cms` is retained deliberately as an **inert key**, and the source says so at `:77-81` — the test *"never matches it on a current clone"* and it is *"kept only so a ROLLBACK/older platform clone that still DEFINES the container gets re-pointed too."* The ⚠️ under *Architecture* below (`:206-211`) states that two-member tuple **in bold**; it **qualifies** this sentence rather than corroborating a one-target reading, and an earlier revision here cited it as corroboration while it said the opposite (booked M257x iter-101, repaired iter-102).
> A freshly-built local stack reads its public content **live from prod**. (Earlier revisions of this doc described a
> `directus/directus:10.10.1` compose service on port 8055 with an `admin@example.com` / `password` admin login
> and an inline `docker-compose.yml` snippet **as if it were CURRENT**, which it is not — there is still no
> Directus service in the platform compose at `0dab54d`.)
>
> **That retraction over-corrected, and this corrects the correction (M257x iter-48).** The twin of this
> paragraph said *"all of that is false; that service **has never existed**"* — repaired at
> [`service_taxonomy.md:350-357`](./service_taxonomy.md) and left standing here. The service **did** exist,
> with exactly that image tag, port and passwo
```

**CITED CONTENT**

```
    55  # (DIRECTUS_PUBLIC_BASE_ADDR) is deliberately NOT re-pointed here — it stays on prod
    56  # public links so browser images stay real (the v1.5 data-plane-local/asset-plane-prod
    57  # split), and re-pointing it would break the baked next/image host whitelist.
    58  DIRECTUS_DATA_CONSUMERS = ("cms", "backend")
    59  
    60  
    61  def resolved_config(platform_dir: str, profiles: str) -> dict:
```

## 09-006
- **id**: `B09-006`
- **corpus site**: `corpus/architecture/external_services.md:168-175` (paragraph)
- **citation**: `docker-compose.yml:110`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
In the **default local posture**, Directus is **not** part of the local stack — `backend` (which hosts the cms
domain since cms-in-app) reaches the **production** Directus over the network. The default `core` profile is
**not** just Postgres + `backend` — but it is far smaller than this page long claimed: it starts **five**
containers. Three are profile-less and so in *every* selection — `postgresql` + `redis` (from the included
`common.yml`) and `sentinel` — and two are the actual `core` members, `backend` (`docker-compose.yml:110`)
and `gotenberg` (`:183`). **There is no `cms`, `jobsimulation` or `roadrunner` container to start** (deleted by `d11a403`, with their `repos.yml` entries), **and no `storage`,
`messenger` or `customerio-sync` one either** (deleted by `838d907`, which also dropped `storage` + `messenger` from `repos.yml`). What survives
is a *production* terraform module — **and not the same one for each**: `cms`'s is still declared **in its own repo** at `service_desired_count = 0` (`cms/terraform/main.tf:39`) — though `6efa1d5` (2026-08-04) deleted cms's build-production workflow saying *"the cms ECR repository is decommissioned (M810)"*, so the two measured facts point opposite ways and the prod-side state is **UNMEASURABLE** (the deletion lands in `infrastructure`, never in a clone set) — while **jobsimulation's ECS service is already destroyed** (`6092c6d2` deleted the `module "jobsimulation"` block — M810 landed for that row) — plus a frozen repo on disk. Neither is a container:
```

**CITED CONTENT**

```
   107          condition: service_healthy
   108        sentinel:
   109          condition: service_started
   110      profiles: [core, backend, all]
   111  
   112    studio-desk:
   113      build:
```

## 09-007
- **id**: `B09-007`
- **corpus site**: `corpus/architecture/external_services.md:168-175` (paragraph)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
In the **default local posture**, Directus is **not** part of the local stack — `backend` (which hosts the cms
domain since cms-in-app) reaches the **production** Directus over the network. The default `core` profile is
**not** just Postgres + `backend` — but it is far smaller than this page long claimed: it starts **five**
containers. Three are profile-less and so in *every* selection — `postgresql` + `redis` (from the included
`common.yml`) and `sentinel` — and two are the actual `core` members, `backend` (`docker-compose.yml:110`)
and `gotenberg` (`:183`). **There is no `cms`, `jobsimulation` or `roadrunner` container to start** (deleted by `d11a403`, with their `repos.yml` entries), **and no `storage`,
`messenger` or `customerio-sync` one either** (deleted by `838d907`, which also dropped `storage` + `messenger` from `repos.yml`). What survives
is a *production* terraform module — **and not the same one for each**: `cms`'s is still declared **in its own repo** at `service_desired_count = 0` (`cms/terraform/main.tf:39`) — though `6efa1d5` (2026-08-04) deleted cms's build-production workflow saying *"the cms ECR repository is decommissioned (M810)"*, so the two measured facts point opposite ways and the prod-side state is **UNMEASURABLE** (the deletion lands in `infrastructure`, never in a clone set) — while **jobsimulation's ECS service is already destroyed** (`6092c6d2` deleted the `module "jobsimulation"` block — M810 landed for that row) — plus a frozen repo on disk. Neither is a container:
```

**CITED CONTENT**

```
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 09-008
- **id**: `B09-008`
- **corpus site**: `corpus/architecture/external_services.md:197-204` (paragraph)
- **citation**: `docker-compose.yml:151`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **Both frontends target `backend`** (`docker-compose.yml:151`/`:160` for next-web-app, `:119`/`:135`
> for studio-desk — build arg then runtime env in each case, all four `:8082/graphql/query`); there is no `cms` service left for them to target
> even if they wanted one. And `backend` does **not** proxy content through a standalone `cms`
> process: `app/cms_reader_switch.go` swaps the cms content reader in-place to the **in-process** cms
> RPC server once Directus is configured, so every content read is *"a DIRECT domain call — no proto round-trip
> … and no internal traffic to a standalone cms."* `backend` requires `DIRECTUS_BASE_ADDR` to boot at all
> (`app/main.go:980-982` `log.Fatalf`s without it — @ `app` `b948604` v1.366.0). The prose two paragraphs above already said this; the
> diagram had not caught up.
```

**CITED CONTENT**

```
   148        # localhost; set it in platform/.env on a remote VM (e.g. PUBLIC_HOST=trillion)
   149        # so the client bundle resolves the user-visible hostname.
   150        args:
   151          NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://${PUBLIC_HOST:-localhost}:8082/graphql/query
   152          NEXT_PUBLIC_BACKEND_API_URL: http://${PUBLIC_HOST:-localhost}:8082
   153          NEXT_PUBLIC_HOSTING_URL: http://${PUBLIC_HOST:-localhost}:3000
   154      ports:
```

## 09-009
- **id**: `B09-009`
- **corpus site**: `corpus/architecture/external_services.md:197-204` (paragraph)
- **citation**: `app/main.go:980-982`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **Both frontends target `backend`** (`docker-compose.yml:151`/`:160` for next-web-app, `:119`/`:135`
> for studio-desk — build arg then runtime env in each case, all four `:8082/graphql/query`); there is no `cms` service left for them to target
> even if they wanted one. And `backend` does **not** proxy content through a standalone `cms`
> process: `app/cms_reader_switch.go` swaps the cms content reader in-place to the **in-process** cms
> RPC server once Directus is configured, so every content read is *"a DIRECT domain call — no proto round-trip
> … and no internal traffic to a standalone cms."* `backend` requires `DIRECTUS_BASE_ADDR` to boot at all
> (`app/main.go:980-982` `log.Fatalf`s without it — @ `app` `b948604` v1.366.0). The prose two paragraphs above already said this; the
> diagram had not caught up.
```

**CITED CONTENT**

```
   977  					cbHandler.SetEnqueuer(workerClient)
   978  				}
   979  				cbHandler.SetUserResolver(authnManager)
   980  				courseBuilderDeps = backend.CourseBuilderDeps{
   981  					Service:       cbSvc,
   982  					Publisher:     cbPublisher,
   983  					AssetUploader: cbAssetSink,
   984  					Notifier:      cbNotifier,
   985  					Handler:       cbHandler,
```

## 09-010
- **id**: `B09-010`
- **corpus site**: `corpus/architecture/external_services.md:206-220` (paragraph)
- **citation**: `rosetta-extensions/stack-injection/gen_injected_override.py:698-699`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-injection/gen_injected_override.py`  (892 lines)

**CLAIMING UNIT**

```md
> **The `--local-content` re-point targets BOTH `cms` and `backend`.** With the v1.5 "prop room" **local
> tooling** (`--local-content` / demo-default) a per-stack `directus` container is added to the stack's
> compose on an offset port, and `rosetta-extensions/stack-injection/gen_injected_override.py:698-699`
> re-points every service in `DIRECTUS_DATA_CONSUMERS`, which is **`("cms", "backend")`** (`:86`) — both @ the demo's **pinned** rext `09d06070`; the same constructs were `:669-670` / `:84` at the prior pin `ab81527a`, which is why an unpinned anchor here rots on every re-pin. `backend`
> is in that tuple because — per the `cms_reader_switch` above — **`backend` is the service that actually
> reads Directus**; re-pointing only `cms` would leave the real reader aimed at production content.
>
> **HISTORICAL — fixed at M257x iter-24 (rext `f9ac72f`).** The tuple originally named `cms` alone, and a
> test (`test_only_cms_is_repointed_not_other_services`) asserted that `backend` must **not** carry the
> re-point — i.e. the suite was *pinning the defect*. Measured on live `demo-1` (2026-08-01) before the fix:
> `cms` had `DIRECTUS_BASE_ADDR=http://directus:8055` while `backend` still had
> `https://content.anthropos.work` with an empty `DIRECTUS_TOKEN`, which surfaced as **96 all-403 lines** in
> `backend`'s log. That test is gone, replaced by `test_backend_the_actual_reader_is_repointed`
> (`stack-injection/tests/test_injection.py:1109`), which asserts the opposite. See
> [`directus-local.md`](../ops/directus-local.md).
```

**CITED CONTENT**

```
   695          # plane (DIRECTUS_PUBLIC_BASE_ADDR) is intentionally left on prod (browser images stay real). On a
   696          # --no-local-content demo there's no local Directus, so we leave DIRECTUS_BASE_ADDR alone (the prod-
   697          # read path the corpus documents) and only the token strip applies.
   698          if with_directus and name in DIRECTUS_DATA_CONSUMERS:
   699              env.append(f"DIRECTUS_BASE_ADDR={DIRECTUS_INNETWORK_ADDR}")
   700          # fix16/fix17: STRIP the inherited prod DIRECTUS_TOKEN from EVERY emitted service. The platform's
   701          # shared `env_file: .env` sprays the prod token stack-wide (audited live on demo-1, 2026-06-11:
   702          # backend, sentinel, storage, roadrunner, skiller, skillpath, cms, jobsimulation + both frontends
```

## 09-011
- **id**: `B09-011`
- **corpus site**: `corpus/architecture/external_services.md:206-220` (paragraph)
- **citation**: `stack-injection/tests/test_injection.py:1109`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-injection/tests/test_injection.py`  (2385 lines)

**CLAIMING UNIT**

```md
> **The `--local-content` re-point targets BOTH `cms` and `backend`.** With the v1.5 "prop room" **local
> tooling** (`--local-content` / demo-default) a per-stack `directus` container is added to the stack's
> compose on an offset port, and `rosetta-extensions/stack-injection/gen_injected_override.py:698-699`
> re-points every service in `DIRECTUS_DATA_CONSUMERS`, which is **`("cms", "backend")`** (`:86`) — both @ the demo's **pinned** rext `09d06070`; the same constructs were `:669-670` / `:84` at the prior pin `ab81527a`, which is why an unpinned anchor here rots on every re-pin. `backend`
> is in that tuple because — per the `cms_reader_switch` above — **`backend` is the service that actually
> reads Directus**; re-pointing only `cms` would leave the real reader aimed at production content.
>
> **HISTORICAL — fixed at M257x iter-24 (rext `f9ac72f`).** The tuple originally named `cms` alone, and a
> test (`test_only_cms_is_repointed_not_other_services`) asserted that `backend` must **not** carry the
> re-point — i.e. the suite was *pinning the defect*. Measured on live `demo-1` (2026-08-01) before the fix:
> `cms` had `DIRECTUS_BASE_ADDR=http://directus:8055` while `backend` still had
> `https://content.anthropos.work` with an empty `DIRECTUS_TOKEN`, which surfaced as **96 all-403 lines** in
> `backend`'s log. That test is gone, replaced by `test_backend_the_actual_reader_is_repointed`
> (`stack-injection/tests/test_injection.py:1109`), which asserts the opposite. See
> [`directus-local.md`](../ops/directus-local.md).
```

**CITED CONTENT**

```
  1106          self.assertEqual(addrs, ["- DIRECTUS_BASE_ADDR=http://directus:8055"],
  1107                           f"cms should re-point at the in-network Directus, got {addrs}")
  1108  
  1109      def test_backend_the_actual_reader_is_repointed(self):
  1110          # M257x iter-24. This test replaces `test_only_cms_is_repointed_not_other_services`, which asserted
  1111          # the OPPOSITE — that `backend` must NOT carry the re-point. That was true under the pre-cms-in-app
  1112          # topology and became false when the cms domain folded into app; the test then PINNED the defect as
```

## 09-012
- **id**: `B09-012`
- **corpus site**: `corpus/architecture/external_services.md:247-253` (paragraph)
- **citation**: `docker-compose.yml:53`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
The only Directus-related platform config is the address `backend` points at — and it arrives by **two** routes,
not one. **`backend`'s compose `environment:` block carries exactly ONE `DIRECTUS_*` variable**,
`DIRECTUS_PUBLIC_BASE_ADDR` (`docker-compose.yml:53` @ platform `0c91421`, inside the block that runs `:46-94`);
the rest arrive through `env_file: .env`, and `.env_example` declares only those (`:91-92`). Compose's
`environment:` **overrides** `env_file:`, so re-pointing the *public* address in `.env` alone is a no-op —
the data-plane address `DIRECTUS_BASE_ADDR` is the one that is genuinely `.env`-settable, which is why the
M23 local-content cutover targets it:
```

**CITED CONTENT**

```
    50        - CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo
    51        - CMS_STREAM=cms
    52        - COMPANY_LOGO_STORAGE_PATH=company_logos
    53        - DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
```

## 09-013
- **id**: `B09-013`
- **corpus site**: `corpus/architecture/external_services.md:300-306` (paragraph)
- **citation**: `app/main.go:980-982`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ This is the cms DOMAIN inside `backend`, not the retired `cms` container.** Since cms-in-app the
> Directus client lives at `app/internal/cms/directus/` and runs in-process in `backend`;
> `app/cms_reader_switch.go` swaps the content reader to the in-process cms server, and
> `app/main.go:980-982` makes `DIRECTUS_BASE_ADDR` a hard boot requirement **of `backend`** (@ `app`
> `b948604` v1.366.0). There is **no `cms` container left to start** — the compose at platform `0c91421` declares
> **five** services (**seven** effective, once `include: common.yml` adds the `postgresql`/`redis` floor; it was
> eight/ten at `0dab54d`, before `838d907` dropped `storage`/`messenger`/`customerio-sync`) — and `cms` is not one of them; every content read is `backend`'s own.
```

**CITED CONTENT**

```
   977  					cbHandler.SetEnqueuer(workerClient)
   978  				}
   979  				cbHandler.SetUserResolver(authnManager)
   980  				courseBuilderDeps = backend.CourseBuilderDeps{
   981  					Service:       cbSvc,
   982  					Publisher:     cbPublisher,
   983  					AssetUploader: cbAssetSink,
   984  					Notifier:      cbNotifier,
   985  					Handler:       cbHandler,
```

## 09-014
- **id**: `B09-014`
- **corpus site**: `corpus/architecture/external_services.md:367-367` (table-row)
- **citation**: `terraform/locals.tf:8`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/locals.tf`  (22 lines)

**CLAIMING UNIT**

```md
| **Port** | **8080** everywhere the router still runs — container and ECS alike (`terraform/locals.tf:8` `port = 8080`, `terraform/main.tf:48-49` maps container 8080 → host 8080; `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`). `5050` was **only** the local compose host mapping (`"5050:8080"`), deleted with the service — **there is no `:5050` on a local stack** |
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

## 09-015
- **id**: `B09-015`
- **corpus site**: `corpus/architecture/external_services.md:367-367` (table-row)
- **citation**: `terraform/main.tf:48-49`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
| **Port** | **8080** everywhere the router still runs — container and ECS alike (`terraform/locals.tf:8` `port = 8080`, `terraform/main.tf:48-49` maps container 8080 → host 8080; `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`). `5050` was **only** the local compose host mapping (`"5050:8080"`), deleted with the service — **there is no `:5050` on a local stack** |
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

## 09-016
- **id**: `B09-016`
- **corpus site**: `corpus/architecture/external_services.md:367-367` (table-row)
- **citation**: `config.prod.yaml:5`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/config.prod.yaml`  (64 lines)

**CLAIMING UNIT**

```md
| **Port** | **8080** everywhere the router still runs — container and ECS alike (`terraform/locals.tf:8` `port = 8080`, `terraform/main.tf:48-49` maps container 8080 → host 8080; `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`). `5050` was **only** the local compose host mapping (`"5050:8080"`), deleted with the service — **there is no `:5050` on a local stack** |
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

## 09-017
- **id**: `B09-017`
- **corpus site**: `corpus/architecture/external_services.md:420-421` (paragraph)
- **citation**: `docker-compose.yml:19-27`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
From `docker-compose.yml` at **`2adcf71^1` (`1e8e754`) — the last mainline state before the drop** — the gateway
`depends_on` named **four** services, each `condition: service_started` (`docker-compose.yml:19-27`):
```

**CITED CONTENT**

```
    16        - .env
    17      environment:
    18        - DB_CONNECTION=postgresql://postgres@postgresql:5432/postgres?search_path=sentinel&sslmode=disable
    19        - ENVIRONMENT=development
    20        - PORT=8087
    21      networks:
    22        - app-network
    23      restart: on-failure
    24      depends_on:
    25        postgresql:
    26          condition: service_healthy
    27  
    28    backend:
    29      build:
    30        # Overridable so a branch checked out elsewhere (e.g. a git worktree) can be built
```

## 09-018
- **id**: `B09-018`
- **corpus site**: `corpus/architecture/external_services.md:441-455` (paragraph)
- **citation**: `graphql-wundergraph/CLAUDE.md:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/CLAUDE.md`  (90 lines)

**CLAIMING UNIT**

```md
> **From February 2026 until its deletion the compose service built from `Dockerfile.dev`, not the production `Dockerfile`** — so the local router **regenerated** `schemas/backend.graphqls` from the sibling `../app` checkout at image-build time, while the production `Dockerfile` (which composes the committed `schemas/` as-is) was the CI/prod path. Three eras, and the middle one is the whole point:
>
> | era | `build:` config | effective Dockerfile |
> |---|---|---|
> | `63d285c` (2024-06-20, then named `wundergraph`) → `719befb` | `context:` only, **no `dockerfile:` key** | Docker's default → the **production `Dockerfile`** |
> | `2c85211` (2026-02-27) → `360efd4` | `dockerfile: Dockerfile.dev`; `67ba772` later raised the context `../graphql-wundergraph` → `..` | **`Dockerfile.dev`** |
> | `360efd4` (2026-07-31), merged as `2adcf71` | block deleted | — |
>
> Verify with `git show 2c85211^:docker-compose.yml` (no `dockerfile:` key) against `git show 2c85211:docker-compose.yml` and `git show 1e8e754:docker-compose.yml` lines 6-8.
>
> **`b56d731` does not end the second era**, though its subject line ("drop the WunderGraph router; point local dev at backend") reads as if it does. It only parked the `graphql` block behind a `wundergraph-deprecated` profile — the block is still there, still `dockerfile: graphql-wundergraph/Dockerfile.dev`. `360efd4`, its sibling in the same PR, is the commit that actually deleted it: `git show b56d731:docker-compose.yml` still has `  graphql:` at `:22`, `git show 360efd4:docker-compose.yml` has no such key (positive control — `  backend:` is at `:28`). The [GraphQL Gateway service doc](../services/graphql-wundergraph.md) states it the same way.
>
> **A caution about how to check this, because it is what made the claim wrong for four releases.** `git log -S "graphql-wundergraph/Dockerfile" -- docker-compose.yml` returns exactly two commits and tempts the conclusion *"it always built from `Dockerfile.dev`."* It cannot see otherwise: that **prefixed** path only came into existence at `67ba772`, so the search is structurally blind to both earlier eras — including the one where no `dockerfile:` key existed at all and Docker silently defaulted to the production file. An absent key is invisible to every search for its value.
>
> **And a caution about the archived repo, which the fence above sends you into.** `graphql-wundergraph/CLAUDE.
```

**CITED CONTENT**

```
    36  
    37  Consequences:
    38  
    39  - **The committed `schemas/backend.graphqls` is now load-bearing locally too.** Since the compose stack builds from the production Dockerfile, a schema change in `app` does **not** appear at `:5050` until you regenerate and commit that file here.
    40  - **For prod**, the committed `schemas/*.graphqls` must match the tagged service version — the Supergraph Update action keeps them honest.
    41  - `Dockerfile.dev` (mirrors `make updatesubg`) defines which files become each subgraph SDL.
    42  - Supergraph configs: `supergraph-config-{compose,dev,prod}.yaml` — each now lists a single `backend` entry.
```

## 09-019
- **id**: `B09-019`
- **corpus site**: `corpus/architecture/external_services.md:554-554` (paragraph)
- **citation**: `app/go.mod:14-18`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
The platform relies on multiple AI providers across backend services, Studio tools, and the simulation engine. Go services access AI behind one `ai.AI` interface (OpenAI, Azure, Anthropic, Bedrock, Mistral) — but ⚠️ **that interface is no longer a shared private module for any service a stack builds, and this sentence said "the shared `ai` library" until M257x iter-115.** `app` folded the library into its own tree at `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*): at `app` `ad9f3c49` neither `app/go.mod` nor `sentinel/go.mod` requires `github.com/anthropos-work/ai` — `app/go.mod:14-18` requires `analytics-go`, `colony`, `proto`, `storage`, `taxonomy` and nothing else — and the code lives at `app/internal/ai/`, imported by 67 `.go` files as `github.com/anthropos-work/app/internal/ai`. The only repos that still *require* the module are the frozen `cms` and `jobsimulation` husks, which `repos.yml` @ `0c91421d` does not list and `make init` does not clone. **This was the unrepaired half of a pair for four readings** — *Unified AI Library* in [`ai_architecture.md`](./ai_architecture.md), [`shared_libraries.md`](./shared_libraries.md#ai) and [`architecture/README.md`](./README.md) all already carried the correction, so a reader who arrived here first got the pre-fold answer with three siblings silently contradicting it. **Provider selection and cost tracking are implemented in the consuming services, not in the `ai` library itself** — see [Shared Libraries → ai](./shared_libraries.md#ai). What that selection actually does is **not** an ordered EU-first ladder; see *Routing: what is actually implemented* below before relying on it for a residency argument.
```

**CITED CONTENT**

```
    11  	github.com/ThreeDotsLabs/watermill v1.5.2
    12  	github.com/ThreeDotsLabs/watermill-redisstream v1.4.5
    13  	github.com/anthropics/anthropic-sdk-go v1.61.0
    14  	github.com/anthropos-work/analytics-go v0.3.1
    15  	github.com/anthropos-work/colony v0.35.2
    16  	github.com/anthropos-work/proto v1.210.0
    17  	github.com/anthropos-work/storage v0.15.2
    18  	github.com/anthropos-work/taxonomy v1.2.0
    19  	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
    20  	github.com/asticode/go-astisub v0.42.0
    21  	github.com/avast/retry-go/v4 v4.7.0
```

## 09-020
- **id**: `B09-020`
- **corpus site**: `corpus/architecture/external_services.md:564-564` (table-row)
- **citation**: `app/internal/jobsimulation/agent/report_agent.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/agent/report_agent.go`  (653 lines)

**CLAIMING UNIT**

```md
| **AWS Bedrock (EU)** | `vendor = AnthropicAws` **or** `Anthropic` — both resolve to the *same* Bedrock client | Jobsimulation domain, Backend (app) | `eu.anthropic.claude-sonnet-4-6` (simulation report agent, `app/internal/jobsimulation/agent/report_agent.go:31`; ask-engine, `app/internal/askengine/bedrock.go:25`) and `eu.anthropic.claude-opus-4-8` / `eu.anthropic.claude-sonnet-4-6` (course-builder author/grader, `app/internal/coursebuilder/bedrock.go:23,29`) |
```

**CITED CONTENT**

```
    28  	// transient draft id that started returning 400 "model identifier
    29  	// invalid" once the profile stabilised. Verified 2026-04-27 against
    30  	// `aws bedrock-runtime invoke-model` with prod IAM grants.
    31  	defaultAgentModel = "eu.anthropic.claude-sonnet-4-6"
    32  
    33  	maxAgentTurns = 10
    34  	maxTokens     = 16_000
```

## 09-021
- **id**: `B09-021`
- **corpus site**: `corpus/architecture/external_services.md:564-564` (table-row)
- **citation**: `app/internal/askengine/bedrock.go:25`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/askengine/bedrock.go`  (787 lines)

**CLAIMING UNIT**

```md
| **AWS Bedrock (EU)** | `vendor = AnthropicAws` **or** `Anthropic` — both resolve to the *same* Bedrock client | Jobsimulation domain, Backend (app) | `eu.anthropic.claude-sonnet-4-6` (simulation report agent, `app/internal/jobsimulation/agent/report_agent.go:31`; ask-engine, `app/internal/askengine/bedrock.go:25`) and `eu.anthropic.claude-opus-4-8` / `eu.anthropic.claude-sonnet-4-6` (course-builder author/grader, `app/internal/coursebuilder/bedrock.go:23,29`) |
```

**CITED CONTENT**

```
    22  // Defaults for the Bedrock client. Both can be overridden via the
    23  // ASK_MODEL_ID and AWS_REGION environment variables.
    24  const (
    25  	DefaultModelID   = "eu.anthropic.claude-sonnet-4-6"
    26  	DefaultRegion    = "eu-west-1"
    27  	DefaultMaxTokens = 4096
    28  )
```

## 09-022
- **id**: `B09-022`
- **corpus site**: `corpus/architecture/external_services.md:564-564` (table-row)
- **citation**: `app/internal/coursebuilder/bedrock.go:23`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
| **AWS Bedrock (EU)** | `vendor = AnthropicAws` **or** `Anthropic` — both resolve to the *same* Bedrock client | Jobsimulation domain, Backend (app) | `eu.anthropic.claude-sonnet-4-6` (simulation report agent, `app/internal/jobsimulation/agent/report_agent.go:31`; ask-engine, `app/internal/askengine/bedrock.go:25`) and `eu.anthropic.claude-opus-4-8` / `eu.anthropic.claude-sonnet-4-6` (course-builder author/grader, `app/internal/coursebuilder/bedrock.go:23,29`) |
```

**CITED CONTENT**

```
    20  	// a 400 "temperature is deprecated for this model". The author
    21  	// path therefore routes through SingleShotNoSampling (see
    22  	// authorClientAdapter below) so no sampling knob reaches Bedrock.
    23  	DefaultAuthorModelID = "eu.anthropic.claude-opus-4-8"
    24  
    25  	// DefaultGraderModelID is the canonical Sonnet grader, held fixed
    26  	// so the ≥ 90 floor means the same thing across runs. Sonnet 4.6
```

## 09-023
- **id**: `B09-023`
- **corpus site**: `corpus/architecture/external_services.md:565-565` (table-row)
- **citation**: `app/internal/cms/studio/markdownManager.go:30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/markdownManager.go`  (144 lines)

**CLAIMING UNIT**

```md
| **Mistral (EU)** | direct client, not via the AI manager | cms domain (Go) + the in-image `studio/tools/` CLI — **never** the generation pipeline | **OCR only** — `mistralocr.New(aiKey)` in `app/internal/cms/studio/markdownManager.go:30` (inside `func NewMarkdownManager`, `:29`; field `ocr *mistralocr.Client` at `:14`, import at `:10`) for studio attachment → markdown. **There is no `mistral.NewMistral(...)` — that symbol is 0-hits repo-wide at `app` `ad9f3c49`**, and the `:19` this row used to cite is a **doc-comment** line, not code; and `from mistralai import Mistral` at `app/studio/tools/pdf2md.py:24` (`mistral-ocr-latest`), a standalone PDF→markdown utility **the generation pipeline never reaches** — `tools/r3.py` DOES dispatch it, as step 2 of the offline chain (`r3.py:139`, `:190`, `:199-206`), so *"nothing dispatches it"* is false; what holds is that no Go caller and no `gen.py` path does (`git -C app/studio grep -i mistral aeec036a`) |
```

**CITED CONTENT**

```
    27  // may grow a failing constructor again — but there is nothing left in here that can
    28  // fail.
    29  func NewMarkdownManager(aiKey string) (*MarkdownManager, error) {
    30  	return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil
    31  }
    32  
    33  func (m *MarkdownManager) OCRProcess(ctx context.Context, documentData []byte) (*string, int, error) {
```

## 09-024
- **id**: `B09-024`
- **corpus site**: `corpus/architecture/external_services.md:565-565` (table-row)
- **citation**: `app/studio/tools/pdf2md.py:24`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/tools/pdf2md.py`  (233 lines)

**CLAIMING UNIT**

```md
| **Mistral (EU)** | direct client, not via the AI manager | cms domain (Go) + the in-image `studio/tools/` CLI — **never** the generation pipeline | **OCR only** — `mistralocr.New(aiKey)` in `app/internal/cms/studio/markdownManager.go:30` (inside `func NewMarkdownManager`, `:29`; field `ocr *mistralocr.Client` at `:14`, import at `:10`) for studio attachment → markdown. **There is no `mistral.NewMistral(...)` — that symbol is 0-hits repo-wide at `app` `ad9f3c49`**, and the `:19` this row used to cite is a **doc-comment** line, not code; and `from mistralai import Mistral` at `app/studio/tools/pdf2md.py:24` (`mistral-ocr-latest`), a standalone PDF→markdown utility **the generation pipeline never reaches** — `tools/r3.py` DOES dispatch it, as step 2 of the offline chain (`r3.py:139`, `:190`, `:199-206`), so *"nothing dispatches it"* is false; what holds is that no Go caller and no `gen.py` path does (`git -C app/studio grep -i mistral aeec036a`) |
```

**CITED CONTENT**

```
    21  from typing import Optional
    22  
    23  import tqdm
    24  from mistralai import Mistral
    25  from dotenv import load_dotenv
    26  
    27  MISTRAL_API_KEY = None
```

## 09-025
- **id**: `B09-025`
- **corpus site**: `corpus/architecture/external_services.md:566-566` (table-row)
- **citation**: `internal/cms/directus/collections/jobsimulation.go:1302`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
| **OpenAI Direct (US)** | **two ways in**: (a) `vendor = Openai` from the caller — including the case where the caller never chose, since a simulation sequence with **`ai_vendor` unset defaults to `openai`** in the cms content layer (`internal/cms/directus/collections/jobsimulation.go:1302`); (b) automatic on **HTTP 429** | (a) any sequence authored without an explicit vendor; (b) the jobsimulation AI manager's retry loop | The 429 retry is the only *automatic fallback* — but it is **not** the only route to US OpenAI. Path (a) gets there on the first attempt. See *Routing* below |
```

**CITED CONTENT**

```
  1299  			aiModel = simulation.SimulationAIModel(*seq.AIModel)
  1300  		}
  1301  
  1302  		aiVendor := simulation.Openai
  1303  		if seq.AIVendor != nil {
  1304  			aiVendor = simulation.SimulationAIVendor(*seq.AIVendor)
  1305  		}
```

## 09-026
- **id**: `B09-026`
- **corpus site**: `corpus/architecture/external_services.md:567-567` (table-row)
- **citation**: `app/internal/coursebuilder/bedrock.go:106-113`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
| **Anthropic Direct (first-party API)** | **presence of `ANTHROPIC_API_KEY`**, not a failure fallback | Course Builder (`app/internal/coursebuilder/bedrock.go:106-113` — key set → first-party API with the model id stripped to its bare form, key unset → Bedrock); Studio-Room (`app/studio/services/ai.py:627-664` `AnthropicProvider`, which `TARGET SERVICE = anthropic` would select — but **no shipped `configs/*.ini` does**: all 30 `*_AI_*_MODEL` lines pin `azure`, so this arm is latent, M257x iter-52) | An either/or **backend switch** for authoring/grading, logged at boot (`app/main.go:770` @ `app` `b948604` v1.366.0, `coursebuilder.ModelBackendName()`) |
```

**CITED CONTENT**

```
   103  }
   104  
   105  // newUnderlyingClient picks the backend for one model role:
   106  // ANTHROPIC_API_KEY present → the first-party Anthropic API (with the
   107  // model id normalized to its bare form); absent → AWS Bedrock, the
   108  // legacy path, byte-for-byte what shipped before the switch existed.
   109  func newUnderlyingClient(ctx context.Context, modelID string) (*askengine.BedrockClient, error) {
   110  	if key := strings.TrimSpace(os.Getenv(AnthropicAPIKeyEnv)); key != "" {
   111  		return askengine.NewAnthropicClientWithModel(key, directModelID(modelID))
   112  	}
   113  	return askengine.NewBedrockClientWithModel(ctx, modelID)
   114  }
   115  
   116  // authorNoSamplingUnderlying is the subset of *askengine.BedrockClient
```

## 09-027
- **id**: `B09-027`
- **corpus site**: `corpus/architecture/external_services.md:567-567` (table-row)
- **citation**: `app/studio/services/ai.py:627-664`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/services/ai.py`  (752 lines)

**CLAIMING UNIT**

```md
| **Anthropic Direct (first-party API)** | **presence of `ANTHROPIC_API_KEY`**, not a failure fallback | Course Builder (`app/internal/coursebuilder/bedrock.go:106-113` — key set → first-party API with the model id stripped to its bare form, key unset → Bedrock); Studio-Room (`app/studio/services/ai.py:627-664` `AnthropicProvider`, which `TARGET SERVICE = anthropic` would select — but **no shipped `configs/*.ini` does**: all 30 `*_AI_*_MODEL` lines pin `azure`, so this arm is latent, M257x iter-52) | An either/or **backend switch** for authoring/grading, logged at boot (`app/main.go:770` @ `app` `b948604` v1.366.0, `coursebuilder.ModelBackendName()`) |
```

**CITED CONTENT**

```
   624          
   625          return self.client.chat.completions.create(**params)
   626  
   627  class AnthropicProvider(AIProvider):
   628  
   629      def _initialize_client(self):
   630          # Approved Anthropic models: Claude 4.5+ tiers (Haiku 4.5, Sonnet 4.6,
   631          # Opus 4.7). Extended-thinking tokens are folded into output_tokens
   632          # by the API and billed at the output rate, so a separate column is
   633          # not needed for billing.
   634          self.cost_table = {
   635              'claude-haiku-4-5': {
   636                  'prompt': 1.00/1000000,
   637                  'completion': 5.00/1000000,
   638              },
   639              'claude-haiku-4-5-20251001': {
   640                  'prompt': 1.00/1000000,
   641                  'completion': 5.00/1000000,
   642              },
   643              'claude-sonnet-4-5': {
   644                  'prompt': 3.00/1000000,
   645                  'completion': 15.00/1000000,
   646              },
   647              'claude-sonnet-4-6': {
   648                  'prompt': 3.00/1000000,
   649                  'completion': 15.00/1000000,
   650              },
   651              'claude-opus-4-5': {
   652                  'prompt': 5.00/1000000,
   653                  'completion': 25.00/1000000,
   654              },
   655              'claude-opus-4-6': {
   656                  'prompt': 5.00/1000000,
   657                  'completion': 25.00/1000000,
   658              },
   659              'claude-opus-4-7': {
   660                  'prompt': 5.00/1000000,
   661                  'completion': 25.00/1000000,
   662              },
   663          }
   664          return Anthropic(api_key=self.api_key)
   665  
   666      def _complete(self, history, top_p=1, temperature=0.2, json_output=False, samples=1):
   667          # Extract system prompt if present
```

## 09-028
- **id**: `B09-028`
- **corpus site**: `corpus/architecture/external_services.md:567-567` (table-row)
- **citation**: `app/main.go:770`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| **Anthropic Direct (first-party API)** | **presence of `ANTHROPIC_API_KEY`**, not a failure fallback | Course Builder (`app/internal/coursebuilder/bedrock.go:106-113` — key set → first-party API with the model id stripped to its bare form, key unset → Bedrock); Studio-Room (`app/studio/services/ai.py:627-664` `AnthropicProvider`, which `TARGET SERVICE = anthropic` would select — but **no shipped `configs/*.ini` does**: all 30 `*_AI_*_MODEL` lines pin `azure`, so this arm is latent, M257x iter-52) | An either/or **backend switch** for authoring/grading, logged at boot (`app/main.go:770` @ `app` `b948604` v1.366.0, `coursebuilder.ModelBackendName()`) |
```

**CITED CONTENT**

```
   767  	appWorker := worker.NewServer(redisAddr)
   768  	workerHandler := tasks.NewHandler(
   769  		ent,
   770  		repo,
   771  		aiClient,
   772  		ocrClient,
   773  		pub,
```

## 09-029
- **id**: `B09-029`
- **corpus site**: `corpus/architecture/external_services.md:597-600` (bullet)
- **citation**: `internal/cms/studio/markdownManager.go:30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/markdownManager.go`  (144 lines)

**CLAIMING UNIT**

```md
5. **Mistral is not in this manager at all.** *Every* use of it in `app` is **OCR** — the cms domain's Go
   client (`internal/cms/studio/markdownManager.go:30`, `studioManager.go:583`) and, in the in-image studio
   tree, `studio/tools/pdf2md.py:24` (a standalone CLI, off the generation pipeline) — so it can neither
   receive nor pass on a simulation request.
```

**CITED CONTENT**

```
    27  // may grow a failing constructor again — but there is nothing left in here that can
    28  // fail.
    29  func NewMarkdownManager(aiKey string) (*MarkdownManager, error) {
    30  	return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil
    31  }
    32  
    33  func (m *MarkdownManager) OCRProcess(ctx context.Context, documentData []byte) (*string, int, error) {
```

## 09-030
- **id**: `B09-030`
- **corpus site**: `corpus/architecture/external_services.md:597-600` (bullet)
- **citation**: `studio/tools/pdf2md.py:24`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/tools/pdf2md.py`  (233 lines)

**CLAIMING UNIT**

```md
5. **Mistral is not in this manager at all.** *Every* use of it in `app` is **OCR** — the cms domain's Go
   client (`internal/cms/studio/markdownManager.go:30`, `studioManager.go:583`) and, in the in-image studio
   tree, `studio/tools/pdf2md.py:24` (a standalone CLI, off the generation pipeline) — so it can neither
   receive nor pass on a simulation request.
```

**CITED CONTENT**

```
    21  from typing import Optional
    22  
    23  import tqdm
    24  from mistralai import Mistral
    25  from dotenv import load_dotenv
    26  
    27  MISTRAL_API_KEY = None
```

## 09-031
- **id**: `B09-031`
- **corpus site**: `corpus/architecture/external_services.md:622-632` (bullet)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:905`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
4. **an authored simulation sequence that simply leaves `ai_vendor` unset** — the easiest of the
   four to miss, because nothing in the AI manager looks like a US default. `ai_vendor` is a
   *nullable* Directus field (`app/internal/cms/directus/collections/jobsimulation.go:905`
   `AIVendor *AIVendor`), and when it is nil the cms content layer supplies `openai` as the
   default (`:1302-1305`, `aiVendor := simulation.Openai`). That value reaches
   `internal/jobsimulation/simulator/ai/ai.go:58-59`, which maps `simulation.Openai` →
   `internalAi.Openai`, and `getClient` resolves that to `a.openaiClient` — the plain
   `openai.NewOpenAI(openaiKey)` client built at `internal/jobsimulation/ai/ai.go:80`, i.e.
   **direct US OpenAI**, on the very first attempt rather than as a 429 retry. (The same switch's
   own `default:` arm at `:114-115` is `internalAi.Openai` too, so an *unrecognized* vendor string
   lands in the same place.)
```

**CITED CONTENT**

```
   902  	ValidationAcceptanceCriteria []ValidationCriterion `json:"validation_acceptance_criteria"`
   903  	ValidationEvaluationCriteria []ValidationCriterion `json:"validation_evaluation_criteria"`
   904  
   905  	AIVendor *AIVendor `json:"ai_vendor"`
   906  	AIModel  *AIModel  `json:"ai_model"`
   907  
   908  	EvaluationSkills  []Skills `json:"evaluation_skills,omitempty"`
```

## 09-032
- **id**: `B09-032`
- **corpus site**: `corpus/architecture/external_services.md:622-632` (bullet)
- **citation**: `internal/jobsimulation/simulator/ai/ai.go:58-59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/simulator/ai/ai.go`  (166 lines)

**CLAIMING UNIT**

```md
4. **an authored simulation sequence that simply leaves `ai_vendor` unset** — the easiest of the
   four to miss, because nothing in the AI manager looks like a US default. `ai_vendor` is a
   *nullable* Directus field (`app/internal/cms/directus/collections/jobsimulation.go:905`
   `AIVendor *AIVendor`), and when it is nil the cms content layer supplies `openai` as the
   default (`:1302-1305`, `aiVendor := simulation.Openai`). That value reaches
   `internal/jobsimulation/simulator/ai/ai.go:58-59`, which maps `simulation.Openai` →
   `internalAi.Openai`, and `getClient` resolves that to `a.openaiClient` — the plain
   `openai.NewOpenAI(openaiKey)` client built at `internal/jobsimulation/ai/ai.go:80`, i.e.
   **direct US OpenAI**, on the very first attempt rather than as a 429 retry. (The same switch's
   own `default:` arm at `:114-115` is `internalAi.Openai` too, so an *unrecognized* vendor string
   lands in the same place.)
```

**CITED CONTENT**

```
    55  
    56  	switch sequence.AIVendor {
    57  	// OpenAI
    58  	case simulation.Openai:
    59  		aiVendor = internalAi.Openai
    60  		switch sequence.AIModel {
    61  		case simulation.GptFourPointOne:
    62  			aiModel = openai.GPT4Dot1
```

## 09-033
- **id**: `B09-033`
- **corpus site**: `corpus/architecture/external_services.md:622-632` (bullet)
- **citation**: `internal/jobsimulation/ai/ai.go:80`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/ai/ai.go`  (355 lines)

**CLAIMING UNIT**

```md
4. **an authored simulation sequence that simply leaves `ai_vendor` unset** — the easiest of the
   four to miss, because nothing in the AI manager looks like a US default. `ai_vendor` is a
   *nullable* Directus field (`app/internal/cms/directus/collections/jobsimulation.go:905`
   `AIVendor *AIVendor`), and when it is nil the cms content layer supplies `openai` as the
   default (`:1302-1305`, `aiVendor := simulation.Openai`). That value reaches
   `internal/jobsimulation/simulator/ai/ai.go:58-59`, which maps `simulation.Openai` →
   `internalAi.Openai`, and `getClient` resolves that to `a.openaiClient` — the plain
   `openai.NewOpenAI(openaiKey)` client built at `internal/jobsimulation/ai/ai.go:80`, i.e.
   **direct US OpenAI**, on the very first attempt rather than as a 429 retry. (The same switch's
   own `default:` arm at `:114-115` is `internalAi.Openai` too, so an *unrecognized* vendor string
   lands in the same place.)
```

**CITED CONTENT**

```
    77  		return nil, fmt.Errorf("can't create Azure Voice AI client: %w", err)
    78  	}
    79  	// openai client
    80  	openaiClient, err := openai.NewOpenAI(openaiKey)
    81  	if err != nil {
    82  		return nil, fmt.Errorf("can't create OpenAI AI client: %w", err)
    83  	}
```

## 09-034
- **id**: `B09-034`
- **corpus site**: `corpus/architecture/external_services.md:633-643` (bullet)
- **citation**: `app/studio/services/ai.py:704-724`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/services/ai.py`  (752 lines)

**CLAIMING UNIT**

```md
5. **Studio-Room's own `openai` `TARGET SERVICE`.** The generation pipeline's provider set is
   `{openai, azure, anthropic}` (`app/studio/services/ai.py:704-724`), and the `openai` arm builds a
   bare `OpenAI(api_key=…)` against **`https://api.openai.com`** (`:383`, `:706-708`;
   `config_template.ini:30-31`) — no Azure endpoint, no EU region. Item 3 above already names
   Studio-Room for its `anthropic` arm, which is what makes leaving this one out an *internal*
   inconsistency rather than merely an omission. Added M257x iter-49.
   ⚠️ **This arm is NOT reachable as shipped, and iter-49 counted it as if it were.** Every
   `*_AI_*_MODEL` line in all three `app/studio/configs/*.ini` pins the service to **`azure`**, and
   `gen.py:41-53` overrides only `*_API_KEY` / `*_ENDPOINT` from the environment — **never the service
   selector**. So the arm exists and nothing selects it: a config edit would arm it, no env var will.
   Count corrected from *five* to *four live + one latent* at M257x iter-52.
```

**CITED CONTENT**

```
   701          return [content.text for content in response.content], consumption
   702  
   703  
   704  def get_client(engine, target_override=None):
   705      providers = {
   706          'openai': OpenAIProvider,
   707          'azure': AzureProvider,
   708          'anthropic': AnthropicProvider
   709      }
   710  
   711      target_engine = engine[(target_override or engine.get('target') or GenMode.DEFAULT).value]
   712  
   713      provider_class = providers.get(target_engine['service'])
   714      if not provider_class:
   715          raise ValueError(f"Unknown AI service: {target_engine['service']}")
   716      
   717      return provider_class(
   718          api_key=target_engine['api_key'],
   719          endpoint=target_engine.get('endpoint'),
   720          model=target_engine.get('model'),
   721          max_tokens=target_engine.get('max_tokens'),
   722          thinking=target_engine.get('thinking'),
   723          transient_retries=target_engine.get('transient_retries', 0)
   724      )
   725  
   726  
   727  def generate_image(engine, prompt):
```

## 09-035
- **id**: `B09-035`
- **corpus site**: `corpus/architecture/external_services.md:668-677` (paragraph)
- **citation**: `app/studio/gen.py:45-48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/gen.py`  (552 lines)

**CLAIMING UNIT**

```md
Re-derive with `command grep -rho '\b<NAME>\b' --include='*.go'` in `app` (`5ba17044`): the four names above
return **5 / 26 / 13 / 13**. `OPENAI_ORG_ID`, `AZURE_OPENAI_ENDPOINT` and `AZURE_OPENAI_DEPLOYMENT` return
**0** and were removed; `OPENAI_API_KEY` returns 2, both the studio subprocess remap. Use `command grep` —
a `grep` aliased to a `.gitignore`-honouring wrapper undercounts, which is how iter-52 first published
12 / 12. **This block covers `app`'s Go only**; `app/studio/gen.py:45-48` reads a separate list of **six**
bare names — `AZURE_API_KEY`, `AZURE_ENDPOINT`, `OPENAI_API_KEY`, `OPENAI_ENDPOINT`, `ANTHROPIC_API_KEY`,
`ANTHROPIC_ENDPOINT` — and **state the tree with the range**: the studio tree is a *nested* checkout at
**`aeec036a`**, not `app` `ad9f3c49`, so `git show ad9f3c49:studio/gen.py` reads the host ref and is the
wrong grep. (Corrected M257x iter-52; the range and the enumeration corrected again at iter-115, which
measured `:45-47` as three of six and cut the list one line short of `ANTHROPIC_*`.)
```

**CITED CONTENT**

```
    42      settings = dict(load_config_section('SERVICES'))
    43  
    44      # Override settings with environment variables
    45      secrets_keys = [
    46          'AZURE_API_KEY', 'AZURE_ENDPOINT', 
    47          'OPENAI_API_KEY', 'OPENAI_ENDPOINT', 
    48          'ANTHROPIC_API_KEY', 'ANTHROPIC_ENDPOINT']
    49  
    50      for key in secrets_keys:
    51          env_key = os.environ.get(key)
```

## 09-036
- **id**: `B09-036`
- **corpus site**: `corpus/architecture/external_services.md:822-823` (bullet)
- **citation**: `app/internal/web/backend/backend.go:130`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/backend.go`  (349 lines)

**CLAIMING UNIT**

```md
- Inspect **backend** logs (`docker compose logs backend`) for `/api/webhook/clerk` errors — Clerk
  user/org sync is app/backend's job (`app/internal/web/backend/backend.go:130`), not Sentinel's
```

**CITED CONTENT**

```
   127  		authnEcho.EchoAuthnMiddleware(authnManager,
   128  			// Skip the following endpoints:
   129  			"/api/webhook/stripe",
   130  			"/api/webhook/clerk",
   131  			"/api/webhook/elevenlabs",
   132  			"/api/webhook/heygen",
   133  			"/api/webhook/chime",
```

## 09-037
- **id**: `B09-037`
- **corpus site**: `corpus/architecture/frontend_architecture.md:7-11` (paragraph)
- **citation**: `docker-compose.yml:141`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **Note**: There are **two other frontend products** that live outside this monorepo and have their own architecture:
> - **[Studio-Desk](../services/studio-desk.md)** — Vite + Express, simulation design tool
> - **[Ant Academy](../services/ant-academy.md)** — Next.js 16 + Expo, internal learning portal for `@anthropos.work` employees
>
> **Studio-Desk** is in `repos.yml` (so `make init` clones it) and *does* have a `studio-desk` compose profile (`docker-compose.yml:141`, `profiles: [studio-desk, all]` — the fact has survived platform `d11a403` **and** `838d907`; only the line number keeps moving — it was 226 before the support containers were deleted). **Ant Academy is deliberately NOT in `repos.yml`** — `make init` never clones it; a demo gets it from `ensure-clones.sh` phase d2, a dev box by hand. At `0c91421` `repos.yml` holds exactly **4** entries — `app`, `sentinel`, `next-web-app`, `studio-desk` (it was 9 at `2adcf71`, before the cms/jobsimulation/roadrunner and storage/messenger drops) — and ant-academy is not one of them. The rest of this document is about `next-web-app` specifically.
```

**CITED CONTENT**

```
   138      depends_on:
   139        backend:
   140          condition: service_started
   141      profiles: [studio-desk, all]
   142  
   143    next-web-app:
   144      build:
```

## 09-038
- **id**: `B09-038`
- **corpus site**: `corpus/architecture/frontend_architecture.md:39-39` (paragraph)
- **citation**: `docker-compose.yml:161`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
The frontend communicates with the backend **primarily — but NOT exclusively — through GraphQL** (**`backend` at `:8082/graphql/query` locally** since platform `2adcf71`; the Cosmo Router at `:8080/graphql` in prod — `:5050` was only ever the deleted LOCAL compose host mapping, never a production port; env `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` either way) using `graphql-request` + TanStack React Query, with Clerk bearer tokens injected per-request via `useGraphql` (`Authorization: Bearer <token>`). There are **no** direct Connect/gRPC calls from the frontend — **but there are direct REST/SSE calls**, spread over **22 non-test source files** that name `NEXT_PUBLIC_BACKEND_API_URL` (`:8082`, `docker-compose.yml:161` runtime env, `:152` build arg): the four `packages/core-js` REST clients (course-builder, credits, Talk-to-Data, workforce), AI-readiness (`useAIReadiness.ts` plus the client and its email-preview modal), invitations (`invite/[token]/page.tsx`), the reminder-unsubscribe page, workforce member analytics, the assignment builder and its ask endpoint (`useAssignmentBuilder.ts`, `useAssignmentAsk.ts` — in **both** `apps/web` and `apps/hiring`), Stripe (`useStripe.tsx` / `useStripe.ts`, plus `useManageSubscription.tsx`), CSV bulk import, onboarding résumé import, the AI-simulation test-run network panel, and the admin backfill tools. *"GraphQL only"* is the wrong mental model for the data layer.
```

**CITED CONTENT**

```
   158      environment:
   159        - NODE_ENV=production
   160        - NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://${PUBLIC_HOST:-localhost}:8082/graphql/query
   161        - NEXT_PUBLIC_BACKEND_API_URL=http://${PUBLIC_HOST:-localhost}:8082
   162        - NEXT_PUBLIC_HOSTING_URL=http://${PUBLIC_HOST:-localhost}:3000
   163      networks:
   164        - app-network
```

## 09-039
- **id**: `B09-039`
- **corpus site**: `corpus/architecture/frontend_architecture.md:41-50` (paragraph)
- **citation**: `packages/core-js/src/talkToData/api.ts:214`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/packages/core-js/src/talkToData/api.ts`  (399 lines)

**CLAIMING UNIT**

```md
> **⚠️ Two quantities, and this passage used to conflate them** (corrected M257x iter-102; it previously said *"**but there are direct REST/SSE calls**, **29 of them across 21 non-test files**"*, with the four core-js clients glossed as *"12 sites between them"*). **29 / 12 were never counts of CALLS** — they were counts of `NEXT_PUBLIC_BACKEND_API_URL` **occurrences**, which is roughly a third of the call surface, so the sentence undercounted the thing its own noun named. Re-derived at `next-web-app` **`8297c684`** (the ref this stack builds, 2026-08-05; 41 commits / 192 files past the `bb3313bc` the old figures came from), over the non-test `.ts`/`.tsx` files, `e2e/` excluded:
>
> | quantity | `bb3313bc` | **`8297c684`** |
> |---|---|---|
> | files naming `NEXT_PUBLIC_BACKEND_API_URL` | 21 | **22** |
> | occurrences of that env var | 29 | **31** |
> | `fetch(` call sites in those files | 43 | **47** |
> | of which, the four `packages/core-js` clients | 12 env / 25 calls | **12 env / 27 calls** |
>
> `new EventSource(` is **0** everywhere — the "SSE" half rides on a streamed `fetch` POST, and `packages/core-js/src/talkToData/api.ts:214` says why: *"uses POST so we can send a JSON body — that rules out EventSource."* Every figure here is **a reading at a ref, never a standing current**: `origin/main` `f97ba659` (4 commits further) reads the same 31 / 22 / 47, but that will move again. The long-standing *"~15 sites"* undercounted by more than half on either measure.
```

**CITED CONTENT**

```
   211  
   212  /**
   213   * Open an SSE connection to /ask/stream and yield parsed events. The fetch
   214   * uses POST so we can send a JSON body — that rules out EventSource and
   215   * forces us to parse the SSE wire format manually.
   216   *
   217   * The generator returns when the stream ends or the signal aborts; the caller
```

## 09-040
- **id**: `B09-040`
- **corpus site**: `corpus/architecture/platform-migration-status.md:60-62` (bullet)
- **citation**: `docker-compose.yml:18`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **`migrations: false` entails nothing on its own.** `sentinel` is `migrations: false` *and* alive *with its
  own `sentinel` schema* (`docker-compose.yml:18`, `search_path=sentinel`). Read the `prod` and
  `fresh local stack` columns, never the flag alone. Live at `0c91421` (`repos.yml:18-20`).
```

**CITED CONTENT**

```
    15      env_file:
    16        - .env
    17      environment:
    18        - DB_CONNECTION=postgresql://postgres@postgresql:5432/postgres?search_path=sentinel&sslmode=disable
    19        - ENVIRONMENT=development
    20        - PORT=8087
    21      networks:
```

## 09-041
- **id**: `B09-041`
- **corpus site**: `corpus/architecture/platform-migration-status.md:60-62` (bullet)
- **citation**: `repos.yml:18-20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
- **`migrations: false` entails nothing on its own.** `sentinel` is `migrations: false` *and* alive *with its
  own `sentinel` schema* (`docker-compose.yml:18`, `search_path=sentinel`). Read the `prod` and
  `fresh local stack` columns, never the flag alone. Live at `0c91421` (`repos.yml:18-20`).
```

**CITED CONTENT**

```
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
    21  
    22    # Frontend
    23    - name: next-web-app
```

## 09-042
- **id**: `B09-042`
- **corpus site**: `corpus/architecture/platform-migration-status.md:75-81` (paragraph)
- **citation**: `docker-compose.yml:185-186`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
**Completeness is measured, not asserted.** The row set is the union of *every name that has ever appeared in `repos.yml`* (`git log -p --follow -- repos.yml` → **14** names: app ·
chronos · cms · graphql-wundergraph · intelligence · jobsimulation · messenger · next-web-app · roadrunner · sentinel · skiller · skillpath · storage · studio-desk — all 14 have
rows) and *every service that has ever appeared in `docker-compose.yml`* (same command on that file → **25** names, including the pre-history the clone set never knew: `nats`,
`web-app`, `chromedp`, `simulator`, `realtime`). **Collect only the keys under `services:`, not every two-space key in the file:** a section-blind pass returns **26**, and the one
extra token is `app-network` — the **network** declared under `networks:` (`docker-compose.yml:185-186` @ `0c91421`), never a service and correctly without a row. **This passage
said 26 until M257x iter-102**, which turned its own audit instruction into a false alarm on every re-run. Re-run those two commands to audit this table; a name they return that
has no row is a gap — three of the 25 have their row filed under the repo name rather than the compose key (`backend` → `app`; `graphql` and `wundergraph` → `graphql-wundergraph`).
```

**CITED CONTENT**

```
   182        - app-network
   183      profiles: [core, backend, all]
   184  
   185  networks:
   186    app-network:
   187  
```

## 09-043
- **id**: `B09-043`
- **corpus site**: `corpus/architecture/platform-migration-status.md:87-87` (table-row)
- **citation**: `app/terraform/main.tf:181`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:14-17` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **seven** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync — each with **its own** wiring call site: skiller `app/main.go:690` (`skiller.NewSkillerManager`), jobsimulation `:721` (`jobsimwiring.Wire`), skillpath `:751` (`skillpath.NewSessionManager`), cms `:1153` (`appcms.Wire`), storage `:524` (`internalstorage.NewManager`), messenger `:1471` (`msgadapters.Wire`), customerio-sync `:395` (`customeriosync.New`) — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync}/`. **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reaso
```

**CITED CONTENT**

```
   178    tags                           = var.tags
   179    aws_region                     = var.aws_region
   180    project                        = local.project
   181    service_desired_count          = 1
   182    service_cpu                    = local.service_cpu
   183    service_memory                 = local.service_memory
   184    service_port                   = local.port
```

## 09-044
- **id**: `B09-044`
- **corpus site**: `corpus/architecture/platform-migration-status.md:87-87` (table-row)
- **citation**: `repos.yml:14-17`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:14-17` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **seven** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync — each with **its own** wiring call site: skiller `app/main.go:690` (`skiller.NewSkillerManager`), jobsimulation `:721` (`jobsimwiring.Wire`), skillpath `:751` (`skillpath.NewSessionManager`), cms `:1153` (`appcms.Wire`), storage `:524` (`internalstorage.NewManager`), messenger `:1471` (`msgadapters.Wire`), customerio-sync `:395` (`customeriosync.New`) — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync}/`. **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reaso
```

**CITED CONTENT**

```
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
    14    - name: app
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
```
