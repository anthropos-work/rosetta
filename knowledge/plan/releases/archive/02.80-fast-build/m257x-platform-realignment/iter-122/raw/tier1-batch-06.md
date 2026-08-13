# TIER-1 ADJUDICATION BATCH 06 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 06-001
- **id**: `B06-001`
- **corpus site**: `corpus/services/sentinel.md:5-5` (paragraph)
- **citation**: `docker-compose.yml:156`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Sentinel is the **centralized authorization service** of the platform. Its **only** live caller is **`app`** — including the jobsimulation and cms authz call sites it absorbed in-process — which reaches it over Connect-RPC to check permissions before executing operations. (There are no `cms` or `jobsimulation` containers left to receive the address: platform `d11a403` deleted both compose services along with `roadrunner`, so at `0dab54d` `AUTHORIZATION_ADDRESS` is set in exactly **one** block — backend's, `docker-compose.yml:48`.) **`messenger` is not a caller** — and ⚠️ **the evidence clause has to be past tense, because there is no messenger compose block to read (corrected M257x iter-115).** At platform `0c91421d`, `docker-compose.yml` declares **five** services — `sentinel` (`:5`), `backend` (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`), `gotenberg` (`:170`) — and `git grep -n messenger 0c91421d -- docker-compose.yml common.yml repos.yml` returns **only comments**. `838d907` (*"drop the storage, messenger and customerio-sync containers"*, 2026-08-05) deleted it. The sentence read *"its compose block sets no `AUTHORIZATION_ADDRESS` and declares no `depends_on: sentinel`"* in the **present** tense, presupposing a block that does not exist — true at `0dab54d` (where the block began at `docker-compose.yml:156`) and silently expired. **Ten other corpus sites already recorded the deletion, two of them in this same file**, neither framed as a retraction of this sentence — one survivor against ten witnesses. What survives, and was re-derived: messenger's Go source imports no authorization client (`git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` returns one unrelated hit, against `colony` present as a positive control); [`clerk-integration.md`](./clerk-integration.md) says the same ("storage, messenger — no auth"). It wraps **Casbin v3** with a PostgreSQL-backed policy store and a single in-memory enforcer that handles all of Anthropos's authorization patterns.
```

**CITED CONTENT**

```
   153          NEXT_PUBLIC_HOSTING_URL: http://${PUBLIC_HOST:-localhost}:3000
   154      ports:
   155        - "3000:3000"
   156      env_file:
   157        - .env
   158      environment:
   159        - NODE_ENV=production
```

## 06-002
- **id**: `B06-002`
- **corpus site**: `corpus/services/sentinel.md:12-12` (bullet)
- **citation**: `go.mod:3`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
* **Language**: Go 1.26 (`go.mod:3` `go 1.26.0`; `Dockerfile:2` / `Dockerfile.dev:2` `golang:1.26-bookworm`)
```

**CITED CONTENT**

```
     1  module github.com/anthropos-work/app
     2  
     3  go 1.26.4
     4  
     5  require (
     6  	code.sajari.com/docconv v1.3.8
```

## 06-003
- **id**: `B06-003`
- **corpus site**: `corpus/services/sentinel.md:22-22` (paragraph)
- **citation**: `terraform/locals.tf:4-5`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/locals.tf`  (22 lines)

**CLAIMING UNIT**

```md
Sentinel's data model is exactly one table (Casbin's `casbin_rules`), and it doesn't participate in the federation gateway because its concerns are orthogonal to product data. Keeping it lean makes it cheap to operate (256 CPU / 128 MB on ECS — `terraform/locals.tf:4-5`) and easy to test (all unit tests use in-memory enforcers, no DB fixtures).
```

**CITED CONTENT**

```
     1  locals {
     2    tags = {
     3      Environment = "${var.environment}"
     4      Engine      = "terraform"
     5    }
     6    project   = "backend"
     7    port      = 8080
     8    rpc_port  = 8081
```

## 06-004
- **id**: `B06-004`
- **corpus site**: `corpus/services/sentinel.md:40-40` (bullet)
- **citation**: `app/internal/data/ent/enum/membership.go:8-15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/data/ent/enum/membership.go`  (57 lines)

**CLAIMING UNIT**

```md
* `g2(org, user, role)` — `admin` / `member` / `candidate` / `content_creator` per org (the four `MembershipRole` values in `app/internal/data/ent/enum/membership.go:8-15`; `init_policy.sql` seeds policies for all four, `content_creator` in its own block at `init_policy.sql:88-118` with a dedicated `internal/authorization/casbin_content_creator_test.go`)
```

**CITED CONTENT**

```
     5  type MembershipRole string
     6  
     7  const (
     8  	RoleAdmin     MembershipRole = "admin"
     9  	RoleMember    MembershipRole = "member"
    10  	RoleCandidate MembershipRole = "candidate"
    11  	// RoleContentCreator grants a member the ability to create and publish
    12  	// content within the organization plus access to Studio. Like admin and
    13  	// member it originates from Clerk (org role) and is replicated onto the
    14  	// Membership via the Clerk webhook handlers.
    15  	RoleContentCreator MembershipRole = "content_creator"
    16  )
    17  
    18  func (MembershipRole) Values() (kinds []string) {
```

## 06-005
- **id**: `B06-005`
- **corpus site**: `corpus/services/sentinel.md:40-40` (bullet)
- **citation**: `init_policy.sql:88-118`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/init_policy.sql`  (120 lines)

**CLAIMING UNIT**

```md
* `g2(org, user, role)` — `admin` / `member` / `candidate` / `content_creator` per org (the four `MembershipRole` values in `app/internal/data/ent/enum/membership.go:8-15`; `init_policy.sql` seeds policies for all four, `content_creator` in its own block at `init_policy.sql:88-118` with a dedicated `internal/authorization/casbin_content_creator_test.go`)
```

**CITED CONTENT**

```
    85      ('p5','default','member','simulation_type_interview:*','execute','',''),
    86      ('p5','default','member','simulation_type_training:*','execute','',''),
    87  
    88   -- ============================================================
    89   -- content_creator = member capabilities + content authoring/publishing + Studio.
    90   -- A first-class org role (synced from Clerk like admin/member). Keep these rows
    91   -- in sync with the app OrganizationPermissionMap (resolver_queries.go).
    92   -- ============================================================
    93   --   p2 = org, sub_role, obj_role, act
    94      -- content_creator as subject (mirrors member)
    95      ('p2','default','content_creator','admin','org:action:profile:read','',''),
    96      ('p2','default','content_creator','admin','read','',''),
    97      ('p2','default','content_creator','member','org:action:profile:read','',''),
    98      ('p2','default','content_creator','member','read','',''),
    99      ('p2','default','content_creator','content_creator','org:action:profile:read','',''),
   100      ('p2','default','content_creator','content_creator','read','',''),
   101      -- content_creator as object (admins/members see them like members)
   102      ('p2','default','admin','content_creator','org:action:js:sessions:read','',''),
   103      ('p2','default','admin','content_creator','org:action:profile:read','',''),
   104      ('p2','default','admin','content_creator','org:action:sp:sessions:read','',''),
   105      ('p2','default','admin','content_creator','org:action:assignments:write','',''),
   106      ('p2','default','admin','content_creator','read','',''),
   107      ('p2','default','member','content_creator','org:action:profile:read','',''),
   108      ('p2','default','member','content_creator','read','',''),
   109   --   p3 = org, sub_role, feat  (member features + content authoring + studio)
   110      ('p3','default','content_creator','org:feature:search','','',''),
   111      ('p3','default','content_creator','org:feature:taxonomy:read','','',''),
   112      ('p3','default','content_creator','org:feature:content:create','','',''),
   113      ('p3','default','content_creator','org:feature:content:publish','','',''),
   114      ('p3','default','content_creator','org:feature:access:studio','','',''),
   115   --   p5 = org, sub_role
```

## 06-006
- **id**: `B06-006`
- **corpus site**: `corpus/services/sentinel.md:85-85` (paragraph)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Consumed via `AUTHORIZATION_ADDRESS=http://sentinel:8087`, set in exactly **one** compose block at platform `0c91421` — **backend**'s, `docker-compose.yml:48` (measured: 1 occurrence across `docker-compose.yml`, `common.yml` and `.env_example`). So the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`**, and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is declared at `:170` and is in the default `core` profile at `:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The correctly-scoped form is the model at [`architecture_overview.md:335`](../architecture/architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*. No other declared service sets `AUTHORIZATION_ADDRESS` — the only ones left to check are `gotenberg`, `studio-desk` and `next-web-app`, and none has the env or a sentinel dependency. The blocks that used to carry it are gone rather than corrected: `jobsimulation`, `cms` and `roadrunner` at `d11a403`, then `storage`, `messenger` and `customerio-sync` at `838d907` — so there is nothing off-path left to hold the address either.
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

## 06-007
- **id**: `B06-007`
- **corpus site**: `corpus/services/sentinel.md:85-85` (paragraph)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Consumed via `AUTHORIZATION_ADDRESS=http://sentinel:8087`, set in exactly **one** compose block at platform `0c91421` — **backend**'s, `docker-compose.yml:48` (measured: 1 occurrence across `docker-compose.yml`, `common.yml` and `.env_example`). So the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`**, and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is declared at `:170` and is in the default `core` profile at `:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The correctly-scoped form is the model at [`architecture_overview.md:335`](../architecture/architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*. No other declared service sets `AUTHORIZATION_ADDRESS` — the only ones left to check are `gotenberg`, `studio-desk` and `next-web-app`, and none has the env or a sentinel dependency. The blocks that used to carry it are gone rather than corrected: `jobsimulation`, `cms` and `roadrunner` at `d11a403`, then `storage`, `messenger` and `customerio-sync` at `838d907` — so there is nothing off-path left to hold the address either.
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

## 06-008
- **id**: `B06-008`
- **corpus site**: `corpus/services/sentinel.md:85-85` (paragraph)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
Consumed via `AUTHORIZATION_ADDRESS=http://sentinel:8087`, set in exactly **one** compose block at platform `0c91421` — **backend**'s, `docker-compose.yml:48` (measured: 1 occurrence across `docker-compose.yml`, `common.yml` and `.env_example`). So the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`**, and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is declared at `:170` and is in the default `core` profile at `:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The correctly-scoped form is the model at [`architecture_overview.md:335`](../architecture/architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*. No other declared service sets `AUTHORIZATION_ADDRESS` — the only ones left to check are `gotenberg`, `studio-desk` and `next-web-app`, and none has the env or a sentinel dependency. The blocks that used to carry it are gone rather than corrected: `jobsimulation`, `cms` and `roadrunner` at `d11a403`, then `storage`, `messenger` and `customerio-sync` at `838d907` — so there is nothing off-path left to hold the address either.
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

## 06-009
- **id**: `B06-009`
- **corpus site**: `corpus/services/sentinel.md:85-85` (paragraph)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Consumed via `AUTHORIZATION_ADDRESS=http://sentinel:8087`, set in exactly **one** compose block at platform `0c91421` — **backend**'s, `docker-compose.yml:48` (measured: 1 occurrence across `docker-compose.yml`, `common.yml` and `.env_example`). So the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`**, and there are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is declared at `:170` and is in the default `core` profile at `:183`, consumed at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The correctly-scoped form is the model at [`architecture_overview.md:335`](../architecture/architecture_overview.md) — *"the only cross-process RPC edge out of backend on a core stack"*. No other declared service sets `AUTHORIZATION_ADDRESS` — the only ones left to check are `gotenberg`, `studio-desk` and `next-web-app`, and none has the env or a sentinel dependency. The blocks that used to carry it are gone rather than corrected: `jobsimulation`, `cms` and `roadrunner` at `d11a403`, then `storage`, `messenger` and `customerio-sync` at `838d907` — so there is nothing off-path left to hold the address either.
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

## 06-010
- **id**: `B06-010`
- **corpus site**: `corpus/services/sentinel.md:89-89` (bullet)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Upstream consumers**: **`app` only** — the sole service that gates requests through Sentinel, and the only compose block that is given its address (`docker-compose.yml:48`). `messenger` and `storage` never called it, and neither is a compose service any more (deleted at `838d907`); `cms`, `jobsimulation` and `roadrunner` went earlier, at `d11a403`
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

## 06-011
- **id**: `B06-011`
- **corpus site**: `corpus/services/sentinel.md:129-129` (paragraph)
- **citation**: `init_policy.sql:63-66`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/init_policy.sql`  (120 lines)

**CLAIMING UNIT**

```md
`init_policy.sql` intentionally omits sensitive capabilities (notably `org:feature:taxonomy:write`, see init_policy.sql:63-66). To grant them locally, apply the on-demand seed:
```

**CITED CONTENT**

```
    60      ('p3','default','admin','org:feature:members:delete','','',''),
    61      ('p3','default','admin','org:feature:members:assign','','',''),
    62      ('p3','default','admin','org:feature:taxonomy:read','','',''),
    63      -- NOTE: org:feature:taxonomy:write is NOT seeded as a default.
    64      -- It grants org admins the ability to write org-scoped custom taxonomy
    65      -- and is a sensitive capability normally reserved for specific orgs /
    66      -- superadmins. See local_superadmin_grants.sql for on-demand local use.
    67  
    68      ('p3','default','member','org:feature:search','','',''),
    69      ('p3','default','member','org:feature:taxonomy:read','','',''),
```

## 06-012
- **id**: `B06-012`
- **corpus site**: `corpus/services/skillpath.md:3-55` (paragraph)
- **citation**: `app/internal/skillpath/session.go:205-207`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skillpath/session.go`  (1058 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"skillpath-in-app"** program (platform milestones **M502 → M507**), the standalone `skillpath` Go
> microservice has been **merged into the `app` monolith** (the service the platform calls "backend") and then
> **decommissioned**. Skillpath no longer runs as a separate service — not in the local compose, not in the
> supergraph, not in production. This is the same pattern as the earlier [skiller-in-app merge](./skiller.md);
> skillpath was the next runtime engine consolidated into `app`.
>
> **Skillpath was always a runtime/session engine, never a content store** — it tracks per-user progression
> *state* (`SkillPathSession → ChapterSession → StepSession`, progress %, completion). The skill-path **content**
> it tracks against (title, cover, curators, chapters → steps, skills-to-verify, versioning) **remains owned by
> the cms domain inside `app`** ([CMS](./cms.md)) and is read by ID **in-process** — `app/internal/skillpath/session.go:205-207`
> (`// cms-in-app deseam: cms is in-process`) → `contentread.CmsContentReader.GetSkillPathDomain`. It was a
> `CMS_RPC_ADDR` Connect-RPC hop before cms-in-app. The consolidation moved the *engine*
> into `app`; it did not touch the content-vs-runtime split.
>
> Where everything went:
>
> * **Domain / engine** — the session manager + repository (`SessionManager`, the get-or-create + version-upgrade
>   logic, the jobsimulation-event subscriber) now live inside `app`: `app/internal/skillpath/`
>   (`session.go`, `session_domain.go`, `repository/`) and `app/internal/skillpaths/`. Ported at **M502/M503**
>   (manager port, dormant) and the subscriber merged at **M504**.
> * **Data** — runtime session state now lives in the **`public` schema** of the shared PostgreSQL database:
>   `public.skill_path_sessions` (Ent schema `app/internal/data/ent/schema/skill_path_session.go` +
>   `skillpath_mixins.go`). The old `skillpath` DB schema is **legacy — a decommissioned empty husk** (the table
>   was kept but holds 0 rows; runtime state is authoritative in `public`). `askengine` and every other reader was
>   re-pointed `skillpath.skill_path_sessions → public.skill_path_sessions`.
> * **RPC — there is NO `SkillPathSessionService` anywhere.** It was not re-hosted in `app`; it was DROPPED.
>   Measured: **0** occurrences in Go source across the clone set and
```

**CITED CONTENT**

```
   202  }
   203  
   204  func (u *SessionManager) getSkillPath(ctx context.Context, skillPathId uuid.UUID, version *string) (*skillpath.SkillPath, error) {
   205  	// cms-in-app deseam: cms is in-process — read the hydrated domain struct
   206  	// directly (no proto round-trip).
   207  	skillPathDomain, err := u.cms.GetSkillPathDomain(ctx, skillPathId, version)
   208  	if err != nil {
   209  		u.logger.Error("failed to fetch skill path", "error", err)
   210  		return nil, fmt.Errorf("failed to fetch skill path: %w", err)
```

## 06-013
- **id**: `B06-013`
- **corpus site**: `corpus/services/skillpath.md:3-55` (paragraph)
- **citation**: `app/CLAUDE.md:109`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/CLAUDE.md`  (357 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"skillpath-in-app"** program (platform milestones **M502 → M507**), the standalone `skillpath` Go
> microservice has been **merged into the `app` monolith** (the service the platform calls "backend") and then
> **decommissioned**. Skillpath no longer runs as a separate service — not in the local compose, not in the
> supergraph, not in production. This is the same pattern as the earlier [skiller-in-app merge](./skiller.md);
> skillpath was the next runtime engine consolidated into `app`.
>
> **Skillpath was always a runtime/session engine, never a content store** — it tracks per-user progression
> *state* (`SkillPathSession → ChapterSession → StepSession`, progress %, completion). The skill-path **content**
> it tracks against (title, cover, curators, chapters → steps, skills-to-verify, versioning) **remains owned by
> the cms domain inside `app`** ([CMS](./cms.md)) and is read by ID **in-process** — `app/internal/skillpath/session.go:205-207`
> (`// cms-in-app deseam: cms is in-process`) → `contentread.CmsContentReader.GetSkillPathDomain`. It was a
> `CMS_RPC_ADDR` Connect-RPC hop before cms-in-app. The consolidation moved the *engine*
> into `app`; it did not touch the content-vs-runtime split.
>
> Where everything went:
>
> * **Domain / engine** — the session manager + repository (`SessionManager`, the get-or-create + version-upgrade
>   logic, the jobsimulation-event subscriber) now live inside `app`: `app/internal/skillpath/`
>   (`session.go`, `session_domain.go`, `repository/`) and `app/internal/skillpaths/`. Ported at **M502/M503**
>   (manager port, dormant) and the subscriber merged at **M504**.
> * **Data** — runtime session state now lives in the **`public` schema** of the shared PostgreSQL database:
>   `public.skill_path_sessions` (Ent schema `app/internal/data/ent/schema/skill_path_session.go` +
>   `skillpath_mixins.go`). The old `skillpath` DB schema is **legacy — a decommissioned empty husk** (the table
>   was kept but holds 0 rows; runtime state is authoritative in `public`). `askengine` and every other reader was
>   re-pointed `skillpath.skill_path_sessions → public.skill_path_sessions`.
> * **RPC — there is NO `SkillPathSessionService` anywhere.** It was not re-hosted in `app`; it was DROPPED.
>   Measured: **0** occurrences in Go source across the clone set and
```

**CITED CONTENT**

```
   106  
   107  The service runs 4 concurrent servers from `main.go`:
   108  1. **Web server** (port 8080) — Echo HTTP with GraphQL endpoint and REST routes (incl. `POST /api/webhook/directus`, which fails closed without `DIRECTUS_WEBHOOK_SECRET`)
   109  2. **RPC server** (port 8081) — one Connect-RPC mux carrying `BackendUsersService`, `BackendOrganizationsService`, `SkillerService`, `SkillPathSessionService`, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`. Runs with a **60s write timeout** — the ported skiller RAG/LLM calls can exceed the old 10s default
   110  3. **Meta server** (port 8083) — Health checks, version info, Asynq inspector
   111  4. **Workers** — Asynq background job processors: the app worker plus the ported skiller, coursebuilder and cms workers
   112  
```

## 06-014
- **id**: `B06-014`
- **corpus site**: `corpus/services/skillpath.md:3-55` (paragraph)
- **citation**: `app/knowledge/architecture.md:28`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/architecture.md`  (131 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"skillpath-in-app"** program (platform milestones **M502 → M507**), the standalone `skillpath` Go
> microservice has been **merged into the `app` monolith** (the service the platform calls "backend") and then
> **decommissioned**. Skillpath no longer runs as a separate service — not in the local compose, not in the
> supergraph, not in production. This is the same pattern as the earlier [skiller-in-app merge](./skiller.md);
> skillpath was the next runtime engine consolidated into `app`.
>
> **Skillpath was always a runtime/session engine, never a content store** — it tracks per-user progression
> *state* (`SkillPathSession → ChapterSession → StepSession`, progress %, completion). The skill-path **content**
> it tracks against (title, cover, curators, chapters → steps, skills-to-verify, versioning) **remains owned by
> the cms domain inside `app`** ([CMS](./cms.md)) and is read by ID **in-process** — `app/internal/skillpath/session.go:205-207`
> (`// cms-in-app deseam: cms is in-process`) → `contentread.CmsContentReader.GetSkillPathDomain`. It was a
> `CMS_RPC_ADDR` Connect-RPC hop before cms-in-app. The consolidation moved the *engine*
> into `app`; it did not touch the content-vs-runtime split.
>
> Where everything went:
>
> * **Domain / engine** — the session manager + repository (`SessionManager`, the get-or-create + version-upgrade
>   logic, the jobsimulation-event subscriber) now live inside `app`: `app/internal/skillpath/`
>   (`session.go`, `session_domain.go`, `repository/`) and `app/internal/skillpaths/`. Ported at **M502/M503**
>   (manager port, dormant) and the subscriber merged at **M504**.
> * **Data** — runtime session state now lives in the **`public` schema** of the shared PostgreSQL database:
>   `public.skill_path_sessions` (Ent schema `app/internal/data/ent/schema/skill_path_session.go` +
>   `skillpath_mixins.go`). The old `skillpath` DB schema is **legacy — a decommissioned empty husk** (the table
>   was kept but holds 0 rows; runtime state is authoritative in `public`). `askengine` and every other reader was
>   re-pointed `skillpath.skill_path_sessions → public.skill_path_sessions`.
> * **RPC — there is NO `SkillPathSessionService` anywhere.** It was not re-hosted in `app`; it was DROPPED.
>   Measured: **0** occurrences in Go source across the clone set and
```

**CITED CONTENT**

```
    25  | Server | Port | Protocol | Purpose |
    26  |--------|------|----------|---------|
    27  | Web | 8080 | HTTP (Echo) | GraphQL endpoint + REST webhook routes |
    28  | RPC | 8081 | Connect-RPC | One mux: `BackendUsersService`, `BackendOrganizationsService`, `SkillerService`, `SkillPathSessionService`, `JobSimulationService`, `CMSService`, `lab.v1.LabSessionService`. Runs with a **60s write timeout** — the ported skiller RAG/LLM calls can exceed the old 10s default |
    29  | Meta | 8083 | HTTP | Health check, version, Asynq inspector |
    30  | Worker | — | Asynq (Redis) | Background job processor (10 concurrent workers) |
    31  
```

## 06-015
- **id**: `B06-015`
- **corpus site**: `corpus/services/skillpath.md:59-62` (bullet)
- **citation**: `app/internal/skillpath/session.go:205-207`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skillpath/session.go`  (1058 lines)

**CLAIMING UNIT**

```md
* **Content-vs-runtime split (unchanged).** "Skillpath" the engine ≠ skill-path *content*. The content it runs
  against — chapters → steps, curators, the job-simulation steps, skills-to-verify, versioning — is owned by
  the **cms domain inside `app`** ([CMS](./cms.md); the `skill_paths` Directus collection) and read by ID **in-process** — `app/internal/skillpath/session.go:205-207` / `app/internal/skillpaths/skillpaths.go:88-95`. **No `CMS_RPC_ADDR` hop** since cms-in-app — and the variable itself is gone: it survived on the `messenger` block, re-pointed at `backend` by `d11a403` (M809), until `838d907` deleted that block. **No compose file sets any `*_RPC_ADDR` now**, and there is no cms process left to address.
  This is the content-vs-runtime split documented in the [Service Taxonomy](../architecture/service_taxonomy.md).
```

**CITED CONTENT**

```
   202  }
   203  
   204  func (u *SessionManager) getSkillPath(ctx context.Context, skillPathId uuid.UUID, version *string) (*skillpath.SkillPath, error) {
   205  	// cms-in-app deseam: cms is in-process — read the hydrated domain struct
   206  	// directly (no proto round-trip).
   207  	skillPathDomain, err := u.cms.GetSkillPathDomain(ctx, skillPathId, version)
   208  	if err != nil {
   209  		u.logger.Error("failed to fetch skill path", "error", err)
   210  		return nil, fmt.Errorf("failed to fetch skill path: %w", err)
```

## 06-016
- **id**: `B06-016`
- **corpus site**: `corpus/services/skillpath.md:59-62` (bullet)
- **citation**: `app/internal/skillpaths/skillpaths.go:88-95`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skillpaths/skillpaths.go`  (204 lines)

**CLAIMING UNIT**

```md
* **Content-vs-runtime split (unchanged).** "Skillpath" the engine ≠ skill-path *content*. The content it runs
  against — chapters → steps, curators, the job-simulation steps, skills-to-verify, versioning — is owned by
  the **cms domain inside `app`** ([CMS](./cms.md); the `skill_paths` Directus collection) and read by ID **in-process** — `app/internal/skillpath/session.go:205-207` / `app/internal/skillpaths/skillpaths.go:88-95`. **No `CMS_RPC_ADDR` hop** since cms-in-app — and the variable itself is gone: it survived on the `messenger` block, re-pointed at `backend` by `d11a403` (M809), until `838d907` deleted that block. **No compose file sets any `*_RPC_ADDR` now**, and there is no cms process left to address.
  This is the content-vs-runtime split documented in the [Service Taxonomy](../architecture/service_taxonomy.md).
```

**CITED CONTENT**

```
    85  }
    86  
    87  func (u *SkillPathManager) getSkillPath(ctx context.Context, skillPathId string) (*skillpath.SkillPath, error) {
    88  	// cms-in-app deseam: cms is in-process — read the hydrated domain struct
    89  	// directly (no proto round-trip).
    90  	id, err := uuid.Parse(skillPathId)
    91  	if err != nil {
    92  		u.logger.Error("invalid skill path id", "id", skillPathId, "error", err)
    93  		return nil, fmt.Errorf("invalid skill path id %q: %w", skillPathId, err)
    94  	}
    95  	skillPathDomain, err := u.cms.GetSkillPathDomain(ctx, id, nil)
    96  	if err != nil {
    97  		u.logger.Error("failed to fetch skill path", "error", err)
    98  		return nil, fmt.Errorf("failed to fetch skill path: %w", err)
```

## 06-017
- **id**: `B06-017`
- **corpus site**: `corpus/services/skillpath.md:76-81` (bullet)
- **citation**: `app/internal/organization/intelligence.go:1144`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/organization/intelligence.go`  (2296 lines)

**CLAIMING UNIT**

```md
* **The manager view reads the RUNTIME session directly — the mirror is GONE.** The **manager insights**
  surface (`insightsSkillPathByMemberships`, the
  `/enterprise/activity-dashboard/@tabs/skill-paths/[skillPathId]` scoreboard in `apps/web`) reads
  `public.skill_path_sessions` — measured: `InsightsSkillPathByMemberships`
  (`app/internal/organization/intelligence.go:1144`) queries `m.ent.SkillPathSession` filtered by
  `skill_path_id` + `status ∈ {active, completed}` + the tenant predicate (`:1159-1170`).
```

**CITED CONTENT**

```
  1141  		All(ctx)
  1142  }
  1143  
  1144  func (m IntelligenceManager) InsightsSkillPathByMemberships(ctx context.Context, organizationID uuid.UUID, memberships []*ent.Membership, skillPathId uuid.UUID, onlyAssignments bool, options InsightOptions) (*InsightsSkillPathByMembershipsResult, error) {
  1145  	membershipUsersMap := make(map[uuid.UUID]*ent.Membership)
  1146  	userIds := make([]uuid.UUID, 0)
  1147  	for _, m := range memberships {
```

## 06-018
- **id**: `B06-018`
- **corpus site**: `corpus/services/skillpath.md:83-94` (paragraph)
- **citation**: `app/terraform/migrations/20260729133514.sql:63`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/migrations/20260729133514.sql`  (65 lines)

**CLAIMING UNIT**

```md
  > **⚠️ RETRACTION — this bullet previously said the opposite, and the instruction was actively harmful.**
  > It told seeders that the scoreboard reads an `app`-side mirror **`public.local_skill_path_session`** with
  > an Ent schema of its own, and that "the mirror row must be co-written." **Both mirrors were DROPPED** —
  > `DROP TABLE "local_skill_path_sessions"` at
  > `app/terraform/migrations/20260729133514.sql:63` (and `local_jobsimulation_sessions` at `:62`)
  > — and no `local_skill_path_session.go` Ent schema exists. (This note used to call that the **last**
  > migration in the repo; it is not — **three** post-date it at `app` `9d00a313`: `20260731131307.sql`,
  > `20260731154527_academy_chapter_progress_completed_at.sql`,
  > `20260803143844_ai_readiness_recommendation_path.sql`. Corrected M257x.) A seeder following the old
  > text would write to a table that is not there. **Seeding the runtime `skill_path_sessions` row is now both
  > necessary and sufficient for this scoreboard.** The generalized manager-view MIRROR trap described in
  > `content-stories-routes.md` no longer applies to skill-paths.
```

**CITED CONTENT**

```
    60  --    migrations) is cleaned up defensively — the app now maintains
    61  --    memberships.last_activity_date itself on session events.
    62  DROP TABLE "local_jobsimulation_sessions";
    63  DROP TABLE "local_skill_path_sessions";
    64  DROP FUNCTION IF EXISTS on_insert_local_jobsimulation_sessions_update_memberships() CASCADE;
    65  
```

## 06-019
- **id**: `B06-019`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/internal/jobsimulation/recording/recording.go:12`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/recording/recording.go`  (132 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
     9  	"github.com/anthropos-work/app/internal/data/ent"
    10  	"github.com/anthropos-work/app/internal/jobsimulation/repository"
    11  	"github.com/anthropos-work/proto/go/domain/storage/v1"
    12  	storagev1 "github.com/anthropos-work/storage/sdk/storage/v1"
    13  	"github.com/google/uuid"
    14  	jsoniter "github.com/json-iterator/go"
    15  	"github.com/redis/go-redis/v9"
```

## 06-020
- **id**: `B06-020`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/internal/jobsimulation/anticheat/anticheat.go:30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/anticheat/anticheat.go`  (770 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
    27  	"github.com/anthropos-work/proto/go/domain/cms/v1/content/simulation"
    28  	"github.com/anthropos-work/proto/go/domain/jobsimulation/v1/interactions/actions"
    29  	"github.com/anthropos-work/proto/go/domain/storage/v1"
    30  	storagev1 "github.com/anthropos-work/storage/sdk/storage/v1"
    31  	"github.com/google/uuid"
    32  	"github.com/hibiken/asynq"
    33  	jsoniter "github.com/json-iterator/go"
```

## 06-021
- **id**: `B06-021`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/main.go:1048`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
  1045  	askHandler := ask.NewWithCourseBuilder(ent, aiClient, dbConn, orgManager, authz, askEngine, bedrockClient, courseBuilderDeps.Service)
  1046  
  1047  	// jobsim-in-app: the ported engine (jobsimDj / jobsimEngine) is wired FATALLY at the top of the
  1048  	// manager-construction section (right after cmsReaderSw), so every app-internal caller below reads it
  1049  	// in-process. The GraphQL managers are pulled off it here.
  1050  	// jobsim-in-app: bundle the folded-GraphQL managers off the (guaranteed non-nil) jobsim runtime
  1051  	// into one dependency for the graph handler. jobsimEngine (the concrete Simulator) stays a local —
```

## 06-022
- **id**: `B06-022`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `storage/terraform/main.tf:13`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/storage/terraform/main.tf`  (101 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
    10  module "storage" {
    11    source = "github.com/anthropos-work/infrastructure.git//modules/services/base_internal_service?ref=main"
    12  
    13    use_fargate = false
    14  
    15    environment                    = var.environment
    16    tags                           = var.tags
```

## 06-023
- **id**: `B06-023`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `storage/terraform/storage.tf:22-25`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/storage/terraform/storage.tf`  (206 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
    19  #
    20  # KNOWN LIMIT — `prevent_destroy` is read from CONFIGURATION, not state. It does not fire when the
    21  # `module "storage-service_euwest1"` block is deleted, which is exactly what the M907
    22  # decommission does. These guards stop accidental drift and targeted destroys; they do NOT stop
    23  # module removal. That is what M903's `moved` blocks are for — relocate the assets out of this
    24  # module BEFORE M907, and the guards become irrelevant to it rather than load-bearing.
    25  #
    26  # These guards are intentionally load-bearing: a plan that would destroy any of them fails at
    27  # plan time. If a future change genuinely needs to remove one, delete its lifecycle block in a
    28  # separate, deliberate commit that says why.
```

## 06-024
- **id**: `B06-024`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `repos.yml:18-20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
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

## 06-025
- **id**: `B06-025`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/main.go:524`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
   521  			"ephemeral disk", internalstorage.EnvBucket, internalstorage.EnvPublicBucket,
   522  			storageBucket, storagePublicBucket)
   523  	}
   524  	storageManager := internalstorage.NewManager(storageBucket)
   525  	publicStorageManager := internalstorage.NewPublicManager(storagePublicBucket)
   526  	// Prove the task role can actually reach both buckets under these names, now,
   527  	// rather than at the first user upload. See verifyBucketAccess in env_guards.go —
```

## 06-026
- **id**: `B06-026`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/internal/storage/service.go:22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/storage/service.go`  (252 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
    19  const (
    20  	// EnvBucket is the private bucket. Empty means local-dev filesystem mode —
    21  	// main.go refuses to boot on an empty value outside a developer machine.
    22  	EnvBucket = "STORAGE_S3_BUCKET"
    23  	// EnvPublicBucket is the public bucket. Same rule.
    24  	EnvPublicBucket = "STORAGE_S3_PUBLIC_BUCKET"
    25  )
```

## 06-027
- **id**: `B06-027`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/main.go:516`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
   513  	// GetPresignedUrl returns ("", nil) — no error, no log, no metric. The terraform
   514  	// that supplies these names ships from a different repo on a different clock than
   515  	// this code, so that ordering is a real deploy sequence, not a hypothetical.
   516  	storageBucket := os.Getenv(internalstorage.EnvBucket)
   517  	storagePublicBucket := os.Getenv(internalstorage.EnvPublicBucket)
   518  	if deployedEnvironment() && (storageBucket == "" || storagePublicBucket == "") {
   519  		log.Fatalf("storage-in-app: %s and %s are both required outside a developer machine "+
```

## 06-028
- **id**: `B06-028`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/main.go:504`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
   501  		log.Fatalf("can't init authentication manager: %v", err)
   502  	}
   503  	// storage-in-app (v9.0 cutover): the in-process object managers. app IS storage
   504  	// now — the standalone service takes no traffic and STORAGE_RPC_ADDR is gone.
   505  	// One private manager + one public manager, constructed HERE like every other
   506  	// manager and passed explicitly to every consumer (the resource manager, the
   507  	// public clients, cmsStorage and jobsimwiring.Wire). Per-namespace clients are
```

## 06-029
- **id**: `B06-029`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/internal/jobsimwiring/wiring.go:101`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
    98  	// serves both app and the ported engine.
    99  	authz *authorization.SentinelManager,
   100  	// storage-in-app (v9.0 cutover): app's own in-process private object manager,
   101  	// replacing the STORAGE_RPC_ADDR edge. Threaded in rather than constructed here
   102  	// so the namespace literal below stays next to the incident comment that
   103  	// explains it.
   104  	inAppStorage *appstorage.Manager,
```

## 06-030
- **id**: `B06-030`
- **corpus site**: `corpus/services/storage.md:7-40` (paragraph)
- **citation**: `app/internal/storagens/callsites_test.go:189`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/storagens/callsites_test.go`  (298 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terrafor
```

**CITED CONTENT**

```
   186  }
   187  
   188  // TestNoRPCStorageClientsRemain is the cutover's one-way door. app owns the buckets
   189  // in process; the standalone storage service is scaled to 0 and STORAGE_RPC_ADDR is
   190  // gone from the task definition. An sdk/storage.NewClient(addr, ns) reintroduced
   191  // anywhere would construct a client pointed at nothing and fail at the first call,
   192  // in whatever code path happens to reach it first.
```

## 06-031
- **id**: `B06-031`
- **corpus site**: `corpus/services/storage.md:44-44` (paragraph)
- **citation**: `docker-compose.yml:82`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Storage is stateless and owns no database: all state lives in S3 — and since platform `0dab54d` **both** managers are wired to **production** buckets in compose (`docker-compose.yml:82`, `:83` @ `0c91421`, on `backend`), not just the public one. Each manager falls back to local filesystem only when ITS bucket variable is set **empty**; on a stock stack neither is. See the hazard note under "Two storage managers".
```

**CITED CONTENT**

```
    79        # silently writes every upload to the container's ephemeral disk.
    80        - AWS_REGION=eu-west-1
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
```

## 06-032
- **id**: `B06-032`
- **corpus site**: `corpus/services/storage.md:62-62` (paragraph)
- **citation**: `app/internal/storage/storage.go:193-200`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/storage/storage.go`  (363 lines)

**CLAIMING UNIT**

```md
Each manager falls back to local filesystem only when ITS bucket env var is empty (private → `/tmp/anthropos-storage/`, public → `/tmp/anthropos-public-storage/`) — `getKeyPath` branches on `s3Bucket != ""` (`app/internal/storage/storage.go:193-200` @ app `2035f9a`), so any **non-empty** value routes to `s3://…` unconditionally, for both managers.
```

**CITED CONTENT**

```
   190  	return req.URL, nil
   191  }
   192  
   193  func (s *manager) getKeyPath(s3Bucket string, namespace string, key uuid.UUID) string {
   194  	if s3Bucket != "" {
   195  		return fmt.Sprintf("s3://%s/%s/%s", s3Bucket, namespace, key.String())
   196  	} else {
   197  		makeTmpDirIfNotExists(path.Join(s.tmpRoot, namespace))
   198  		return fmt.Sprintf("file://%s", path.Join(s.tmpRoot, namespace, key.String()))
   199  	}
   200  }
   201  
   202  var _ StorageManager = (*manager)(nil)
   203  
```

## 06-033
- **id**: `B06-033`
- **corpus site**: `corpus/services/storage.md:64-94` (paragraph)
- **citation**: `docker-compose.yml:82`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalate
```

**CITED CONTENT**

```
    79        # silently writes every upload to the container's ephemeral disk.
    80        - AWS_REGION=eu-west-1
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
```

## 06-034
- **id**: `B06-034`
- **corpus site**: `corpus/services/storage.md:64-94` (paragraph)
- **citation**: `docker-compose.yml:100`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalate
```

**CITED CONTENT**

```
    97      # jobsim-in-app's Chime/LiveKit recording managers use the AWS SDK default
    98      # credential chain — the mount the standalone jobsimulation container had.
    99      volumes:
   100        - $HOME/.aws/credentials:/root/.aws/credentials:ro
   101      depends_on:
   102        # storage, messenger and customerio-sync are not services any more — this one
   103        # container serves all three in-process.
```

## 06-035
- **id**: `B06-035`
- **corpus site**: `corpus/services/storage.md:64-94` (paragraph)
- **citation**: `README.md:81-87`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/README.md`  (175 lines)

**CLAIMING UNIT**

```md
> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalate
```

**CITED CONTENT**

```
    78  corpus/
    79  ├── architecture/          # System design and service relationships
    80  │   ├── architecture_overview.md    # Start here for the big picture
    81  │   ├── service_taxonomy.md         # Core, Studio, External tiers
    82  │   ├── frontend_architecture.md    # Next.js monorepo details
    83  │   ├── external_services.md        # Clerk, Directus, GraphQL
    84  │   ├── dependency_map.md           # Service inter-dependencies
    85  │   ├── shared_libraries.md         # colony, proto, ai, authn, taxonomy
    86  │   └── alignment_testing.md        # The alignment test class + framework (rosetta-extensions/alignment/)
    87  │
    88  ├── services/              # Individual service documentation
    89  │   ├── backend.md, cms.md, sentinel.md, ...     # Core services
    90  │   ├── graphql-wundergraph.md, next-web-app.md  # Gateway + main frontend
```

## 06-036
- **id**: `B06-036`
- **corpus site**: `corpus/services/storage.md:64-94` (paragraph)
- **citation**: `main.go:518-523`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalate
```

**CITED CONTENT**

```
   515  	// this code, so that ordering is a real deploy sequence, not a hypothetical.
   516  	storageBucket := os.Getenv(internalstorage.EnvBucket)
   517  	storagePublicBucket := os.Getenv(internalstorage.EnvPublicBucket)
   518  	if deployedEnvironment() && (storageBucket == "" || storagePublicBucket == "") {
   519  		log.Fatalf("storage-in-app: %s and %s are both required outside a developer machine "+
   520  			"(got bucket=%q public=%q); an empty bucket silently writes to the container's "+
   521  			"ephemeral disk", internalstorage.EnvBucket, internalstorage.EnvPublicBucket,
   522  			storageBucket, storagePublicBucket)
   523  	}
   524  	storageManager := internalstorage.NewManager(storageBucket)
   525  	publicStorageManager := internalstorage.NewPublicManager(storagePublicBucket)
   526  	// Prove the task role can actually reach both buckets under these names, now,
```

## 06-037
- **id**: `B06-037`
- **corpus site**: `corpus/services/storage.md:64-94` (paragraph)
- **citation**: `app/env_guards.go:37-44`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalate
```

**CITED CONTENT**

```
    34  // "production" to Development. A future ENVIRONMENT=staging would therefore silently
    35  // disarm every guard below. This form fails safe instead — an unrecognised value is
    36  // treated as deployed.
    37  func deployedEnvironment() bool {
    38  	switch strings.ToLower(strings.TrimSpace(os.Getenv("ENVIRONMENT"))) {
    39  	case "", "development", "dev", "local", "test":
    40  		return false
    41  	default:
    42  		return true
    43  	}
    44  }
    45  
    46  // Switches for the folded subsystems that reach OUTSIDE the process on a timer or a
    47  // stream — messenger (sends mail) and customerio-sync (rewrites Brevo contacts).
```

## 06-038
- **id**: `B06-038`
- **corpus site**: `corpus/services/storage.md:64-94` (paragraph)
- **citation**: `docker-compose.yml:56`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalate
```

**CITED CONTENT**

```
    53        - DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
```

## 06-039
- **id**: `B06-039`
- **corpus site**: `corpus/services/storage.md:192-192` (paragraph)
- **citation**: `README.md:81-87`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/README.md`  (175 lines)

**CLAIMING UNIT**

```md
**What the container used to do with its buckets, and what the binary still does with them.** With `STORAGE_S3_BUCKET` empty the PRIVATE manager falls back to `/tmp/anthropos-storage/` automatically, and its presigned URLs return empty strings in that mode (`storage.go:122`). FOOTGUN: the PUBLIC manager is not sandboxed by that fallback — the deleted `storage` compose block hardcoded `STORAGE_S3_PUBLIC_BUCKET` to the production public bucket, so `PutPublicObject`/`GetPublicObject` hit real S3 and failed without AWS credentials. (That parenthetical used to say none were set in `platform/.env` — **no longer true**: platform `README.md:81-87` @ `0c91421` instructs you to put live `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`, and `docker-compose.yml:100` mounts `$HOME/.aws/credentials` into `backend`, so on a current stack the credentials are generally present and the write **succeeds**.) Running the binary by hand, override `STORAGE_S3_PUBLIC_BUCKET` to empty; it then falls back to `/tmp/anthropos-public-storage/` (a separate path from the private fallback).
```

**CITED CONTENT**

```
    78  corpus/
    79  ├── architecture/          # System design and service relationships
    80  │   ├── architecture_overview.md    # Start here for the big picture
    81  │   ├── service_taxonomy.md         # Core, Studio, External tiers
    82  │   ├── frontend_architecture.md    # Next.js monorepo details
    83  │   ├── external_services.md        # Clerk, Directus, GraphQL
    84  │   ├── dependency_map.md           # Service inter-dependencies
    85  │   ├── shared_libraries.md         # colony, proto, ai, authn, taxonomy
    86  │   └── alignment_testing.md        # The alignment test class + framework (rosetta-extensions/alignment/)
    87  │
    88  ├── services/              # Individual service documentation
    89  │   ├── backend.md, cms.md, sentinel.md, ...     # Core services
    90  │   ├── graphql-wundergraph.md, next-web-app.md  # Gateway + main frontend
```

## 06-040
- **id**: `B06-040`
- **corpus site**: `corpus/services/storage.md:192-192` (paragraph)
- **citation**: `docker-compose.yml:100`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
**What the container used to do with its buckets, and what the binary still does with them.** With `STORAGE_S3_BUCKET` empty the PRIVATE manager falls back to `/tmp/anthropos-storage/` automatically, and its presigned URLs return empty strings in that mode (`storage.go:122`). FOOTGUN: the PUBLIC manager is not sandboxed by that fallback — the deleted `storage` compose block hardcoded `STORAGE_S3_PUBLIC_BUCKET` to the production public bucket, so `PutPublicObject`/`GetPublicObject` hit real S3 and failed without AWS credentials. (That parenthetical used to say none were set in `platform/.env` — **no longer true**: platform `README.md:81-87` @ `0c91421` instructs you to put live `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`, and `docker-compose.yml:100` mounts `$HOME/.aws/credentials` into `backend`, so on a current stack the credentials are generally present and the write **succeeds**.) Running the binary by hand, override `STORAGE_S3_PUBLIC_BUCKET` to empty; it then falls back to `/tmp/anthropos-public-storage/` (a separate path from the private fallback).
```

**CITED CONTENT**

```
    97      # jobsim-in-app's Chime/LiveKit recording managers use the AWS SDK default
    98      # credential chain — the mount the standalone jobsimulation container had.
    99      volumes:
   100        - $HOME/.aws/credentials:/root/.aws/credentials:ro
   101      depends_on:
   102        # storage, messenger and customerio-sync are not services any more — this one
   103        # container serves all three in-process.
```

## 06-041
- **id**: `B06-041`
- **corpus site**: `corpus/services/storage.md:228-228` (table-row)
- **citation**: `docker-compose.yml:82`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `STORAGE_S3_BUCKET` | (empty) | Private bucket. The deleted `storage` service block never set it, so **this binary** fell back to `/tmp/anthropos-storage/`. **Do not read that across to the live path:** `backend` sets it to the production private bucket (`docker-compose.yml:82` @ `0c91421`) — see the hazard note under "Two storage managers". |
```

**CITED CONTENT**

```
    79        # silently writes every upload to the container's ephemeral disk.
    80        - AWS_REGION=eu-west-1
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
```

## 06-042
- **id**: `B06-042`
- **corpus site**: `corpus/services/storage.md:231-231` (table-row)
- **citation**: `docker-compose.yml:119`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `ENVIRONMENT` | `development` | Environment name. **The block DID set it** — `- ENVIRONMENT=development` at `docker-compose.yml:119` @ platform `0dab54d` and `:206` @ `2adcf71`. (Corrected M257x iter-102; this cell read *(empty)*, i.e. "never set by compose". The error mattered: `development` is precisely the value that makes `deployedEnvironment()` return false and **disarms** app's boot guards — see the hazard note under "Two storage managers" — so recording it as unset hid the mechanism.) |
```

**CITED CONTENT**

```
   116        ssh: ["default"]
   117        args:
   118          VITE_CLERK_PUBLISHABLE_KEY: ${VITE_CLERK_PUBLISHABLE_KEY}
   119          VITE_GRAPHQL_ENDPOINT: ${VITE_GRAPHQL_ENDPOINT:-http://localhost:8082/graphql/query}
   120          VITE_ENVIRONMENT: ${VITE_ENVIRONMENT:-production}
   121          VERSION: dev
   122      ports:
```

## 06-043
- **id**: `B06-043`
- **corpus site**: `corpus/services/studio-desk.md:21-21` (table-row)
- **citation**: `docker-compose.yml:138-140`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| **Deployment** | Runs natively for dev (`npm run dev`), or containerized via the `studio-desk` docker-compose profile (ports 9000/9100). It `depends_on` **`backend` alone** — `docker-compose.yml:138-140` @ platform `0c91421`, with `profiles: [studio-desk, all]` at `:141` (both re-anchored M257x iter-87; they were `:223-225`/`:226` at `0dab54d`, before `838d907` deleted three service blocks above them). It *also* listed **`cms`** (`:337-341` @ `2adcf71`) until that container was deleted from compose at `d11a403`; there is no `cms` service to depend on now, and it never depended on `graphql`, which is likewise no longer a compose service. Built with `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query`. **⚠️ Asking for `studio-desk` as the only profile exits 1** — the profile selects `studio-desk` but *not* the `backend` it depends on, so compose rejects the whole project (`service "studio-desk" depends on undefined service "backend": invalid compose project`). Use `PROFILE=all`, which selects both. |
```

**CITED CONTENT**

```
   135        - VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query
   136      networks:
   137        - app-network
   138      depends_on:
   139        backend:
   140          condition: service_started
   141      profiles: [studio-desk, all]
   142  
   143    next-web-app:
```

## 06-044
- **id**: `B06-044`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `src/routes/skillpath.ts:374`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/routes/skillpath.ts`  (1492 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
   371  // eval quirk, schema change), a cross-tenant leak would be unacceptable, so
   372  // we trade a tiny bit of CPU for a hard guard.
   373  //
   374  // We do NOT route this through the platform's `privateSkillPaths` GraphQL
   375  // query because that resolver uses a strict Go struct for `chapter_list` and
   376  // crashes on any step shape it doesn't recognize (e.g. an image step whose
   377  // `resource` is a bare uuid string instead of `{key,collection}`), making
```
