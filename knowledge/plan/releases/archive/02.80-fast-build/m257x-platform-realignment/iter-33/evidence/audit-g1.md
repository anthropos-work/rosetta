# iter33 KB-fidelity audit — group 1 (corpus/architecture, 5 files)

Read-only audit. Graded against platform origin HEAD `2adcf71` (clone:
`/Users/marco/workspace/anthropos/rosetta/stack-demo/platform`, sibling repos `stack-demo/app`,
`stack-demo/next-web-app`, `stack-demo/ant-academy`).

## 1. Positive control

| File | `wc -l` | Read to |
|:--|--:|:--|
| `corpus/architecture/external_services.md` | 693 | line 693 (full) |
| `corpus/architecture/service_taxonomy.md` | 406 | line 406 (full) |
| `corpus/architecture/architecture_overview.md` | 319 | line 319 (full) |
| `corpus/architecture/dependency_map.md` | 103 | line 103 (full) |
| `corpus/architecture/README.md` | 38 | line 38 (full) |

Total 1559 / 1559 lines read. No file left unread.

### Ground-truth facts re-verified this session (not taken on trust)

- No `graphql` service in `docker-compose.yml`; services are `sentinel`(5) `backend`(28)
  `jobsimulation`(83) `cms`(144) `storage`(189) `customerio-sync`(220) `messenger`(240)
  `roadrunner`(281) `studio-desk`(311) `next-web-app`(344) `gotenberg`(371). `postgresql`/`redis` in
  `common.yml`. `PROFILE ?= graphql` still the Makefile default (`Makefile:10`) — profile name survives,
  service does not.
- `repos.yml`: app(migrations:true, schema:public), cms, jobsimulation, sentinel, storage, messenger,
  roadrunner, next-web-app, studio-desk. No `graphql-wundergraph`, no skiller/skillpath, no ant-academy.
- `backend` env block = `docker-compose.yml:43-67`; `depends_on` = `:70-80` (sentinel, cms, storage).
- `DIRECTUS_BASE_ADDR`/`DIRECTUS_PUBLIC_BASE_ADDR` set explicitly only on `cms` (`:164-165`). ✔
- `app/cms_reader_switch.go` exists; `app/main.go:971-973` `log.Fatalf`s without `DIRECTUS_BASE_ADDR`. ✔
- `app/internal/cms/directus/`, `app/internal/cms/studio/`, `app/internal/jobsimulation/runner/` exist;
  `app/internal/roadrunner/` does not. `JUDGE0_BASE_URL` read at `internal/jobsimwiring/wiring.go:118`. ✔
- Clerk webhook handled by **app/backend**: `app/internal/web/backend/backend.go:130` `/api/webhook/clerk`.
- `backend` serves the Apollo Sandbox playground itself: `backend.go:315` `/graphql` → `:317` `/graphql/query`. ✔
- next-web-app: `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` only; **zero** occurrences of `NEXT_PUBLIC_GRAPHQL_ENDPOINT`.
  All four apps pin `"next": "^16.2.7"`.
- Ports verified: jobsimulation 8400/8401, cms 8090/8091, storage 8300/8301, roadrunner 10400/10401,
  customerio-sync 8080, sentinel 8087, backend 8081-8083 (META_PORT=8084 unpublished). All doc rows correct.

**Ground-truth facts #4 (router dropped locally) and #5 (one subgraph) are handled correctly in all five
files.** Every one of the four larger files carries the v2.8 M257x router banner or an equivalent
in-place fence; the router sweep is clean. No finding below is a router/subgraph finding. The residue is
a different, older class: **"the CMS *service* is the Directus proxy"** — the studio-room archetype
shape, invisible to a merged-status grep because it never says "cms is live".

---

## 2. Findings

### BLOCKER 1 — `external_services.md:594` (and `:437`)

> ```
> NEXT_PUBLIC_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query   # was :5050/graphql
> ```
> (:437) `endpoint: process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT`

**False at HEAD.** `NEXT_PUBLIC_GRAPHQL_ENDPOINT` does not exist in next-web-app — `grep -rn` over the
whole repo returns nothing. The variable the frontends actually read is
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (`platform/docker-compose.yml:352` build arg, `:361` runtime env;
consumed by `next-web-app/packages/graphql/src/hooks/useGraphql.tsx` and
`packages/graphql/src/server/server.graphql.ts`). This line sits in the **"Environment Variables
Checklist → For Next.js Apps"** — a block a developer copies verbatim into `.env`. Doing so sets a
variable nothing reads and silently leaves the app on its built-in default; the M257x router change makes
this endpoint the single most load-bearing frontend var, so a wrong name here misdirects exactly the
work the release is about.

**Grade: BLOCKER.** *Correction:* rename both occurrences to `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`
(cite `docker-compose.yml:352`/`:361`).

---

### BLOCKER 2 — `external_services.md:643`

> `- Inspect Sentinel logs for sync errors`

**False at HEAD.** Clerk user/org sync is handled by **app/backend** — `POST /api/webhook/clerk` is
registered in `app/internal/web/backend/backend.go:130` (route table at
`internal/web/backend/api/server.gen.go:1744`). Sentinel is authorization-only and does not receive
Clerk webhooks. This file *already states this correctly* at `:76` ("Clerk user/org sync is handled by
the `app`/backend service via Clerk webhooks … not by Sentinel"), so `:643` contradicts its own file.
The failure mode is concrete: a developer debugging "users not syncing" tails `docker compose logs
sentinel`, sees nothing, and concludes the webhook fired.

**Grade: BLOCKER.** *Correction:* `- Inspect backend logs (`docker compose logs backend`) for
`/api/webhook/clerk` errors — sync is app/backend's job, not Sentinel's.`

---

### BLOCKER 3 — `external_services.md:617`

> `- Set up webhooks to production Sentinel endpoint`

**False at HEAD**, same root cause as BLOCKER 2. The production webhook target is the backend's
`/api/webhook/clerk`, not Sentinel — Sentinel exposes no webhook route (it is Casbin authz only, per
`:74-76` of this same file). Being under a *Production Deployment* heading does not make it
prod-fenced-and-therefore-exempt: the claim is false about production too.

**Grade: BLOCKER.** *Correction:* "Set up Clerk webhooks to the production **backend** endpoint
`/api/webhook/clerk`."

---

### BLOCKER 4 — `external_services.md:277-294` (section "CMS Service Integration")

> `### CMS Service Integration`
> `The CMS service connects to Directus via:` … `**Code Integration** (from CMS service):`
> ```go
> // internal/directus/
> // - Client initialization
> ```

**False at HEAD, and unfenced present tense.** There is no CMS service doing this. The Directus client
lives at `app/internal/cms/directus/` (verified on disk: `directus.go`, `collection.go`,
`collections/`) and runs **inside `backend`**; `app/cms_reader_switch.go` swaps the content reader to the
in-process cms RPC server, and `app/main.go:971-973` makes `DIRECTUS_BASE_ADDR` a hard boot requirement
*of `backend`*. The `cms` container is a merged husk that does **not** serve `backend`'s content reads —
which this same file states at `:194-197` and warns about explicitly at `:651-652` ("not `cms` — that
container is a merged husk that no longer serves `backend`'s content reads"). A reader who lands on
this section goes to the frozen `cms` repo's `internal/directus/` instead of `app/internal/cms/directus/`.
This is the studio-room archetype: five correct paragraphs earlier in the file, then an unswept section
that never uses the words a merged-status sweep greps for.

Same-class echoes in this file: `:604` heading "**For CMS Service**" (harmless — the vars *are* still set
on the cms husk) and `:691` "CMS Service — Directus proxy/adapter".

**Grade: BLOCKER.** *Correction:* retitle to "cms-domain Directus integration" and repoint the code map
to `app/internal/cms/directus/` (running inside `backend`), with the husk called out.

---

### BLOCKER 5 — `service_taxonomy.md:280` and `:283`

> ```
> Frontend → CMS Service → Directus API (content.anthropos.work) → PostgreSQL
> ```
> `The **CMS Service** acts as a smart proxy/adapter, adding business logic on top of Directus.`

**False at HEAD.** Frontends do not reach a CMS service: both next-web-app and studio-desk are baked
against `:8082/graphql/query` on **`backend`** (`docker-compose.yml:352`/`:361` and `:318`/`:334`), and
the Directus read happens in-process in `backend`'s cms domain. This is the pre-cms-in-app flow that
`architecture_overview.md:82-83` explicitly labels as the *former* one ("*was `Frontend -> CMS ->
Directus` before cms-in-app*"). Unlike the surrounding paragraphs (`:269-276`, `:285-289`), this diagram
carries **no fence at all** and reads as the current architecture — it is the one diagram in this file a
skimming reader would take away. A content bug chased through this diagram lands on the dead `cms`
container.

**Grade: BLOCKER.** *Correction:* `Frontend/Studio-Desk → backend :8082/graphql/query (cms domain,
app/internal/cms/directus/) → Directus (content.anthropos.work) → PostgreSQL`; the proxy is the cms
**domain inside `backend`**, not the cms container.

---

### minor 1 — `external_services.md:132-133`

> `that block (`:43-77` @ `2adcf71`) has no `DIRECTUS_*` at all`

The **claim is true** (no `DIRECTUS_*` in backend's `environment:`), but the range is wrong: backend's
`environment:` block is `docker-compose.yml:43-67`; `:68-77` is `networks:` + `depends_on:`.
*Correction:* `:43-67`.

---

### minor 2 — `external_services.md:14`

> `The Anthropos platform integrates with **three key external services**:` — followed by a four-item list
(Clerk, Directus, GraphQL/Wundergraph, AI Providers). *Correction:* "four" (or drop the count).

---

### minor 3 — `external_services.md:539`

> `- Configured in `studio-room/configs/*.ini``

The repo is **`anthropos-studio-room`** (ground-truth #13; `service_taxonomy.md:170` has it right) and it
is baked into the `app` image, orchestrated from `app/internal/cms/studio/` (verified on disk). A
`studio-room/` directory is not something a reader will find. *Correction:* `anthropos-studio-room/configs/*.ini`.

---

### minor 4 — `external_services.md:549` and `:565`

> `| **Integration** | Jobsimulation service |` (LiveKit, and again for AWS Chime)

There is no jobsimulation *service* doing this at HEAD — LiveKit/Chime live in
`app/internal/jobsimulation/{calls,recording,agent}/` inside `backend`; the `jobsimulation` container is
an unfederated husk. True-but-misleading (the file establishes the merge at `:344` and `:651`).
*Correction:* "the jobsimulation **domain** in `backend` (`app/internal/jobsimulation/`)".

---

### minor 5 — `service_taxonomy.md:329` and `:333`

> `- **Studio-Room**: Direct integration with CMS service for blueprint retrieval`
> `- **Content Storage**: Directus API (via CMS proxy for core services)`

Same class as BLOCKER 5 but lower-stakes bullets. Studio-Room runs *inside the `app` image*, orchestrated
from `app/internal/cms/studio/` — this file says so itself at `:171` and `:190`. *Correction:* say "the
cms domain in `app`" in both.

---

### minor 6 — `service_taxonomy.md:215`

> `cp .env.example .env   # fill Clerk + AI keys`

The corpus's own instruction elsewhere (root `CLAUDE.md`, `corpus/ops/setup_guide.md`,
`corpus/services/ant-academy.md`) is `code/.env.local`, and states the app "reads only from
`code/.env.local`". The repo ships both `ant-academy/code/.env.example` and `code/.env.local`. Next.js
does load a plain `.env`, so this is unlikely to break — but it is an unexplained divergence from the
corpus's stated rule. *Correction:* `cp .env.example .env.local` (matching `:201` and setup_guide).

---

### minor 7 — `architecture_overview.md:220`

> `| **Next Web App** | Next.js 15 | Main user-facing application (Workforce + Hiring) |`

False at HEAD and self-contradictory: `:29` ("Next.js **16**") and `:45` ("`next: ^16.2.7` across all
four apps") in this same file are correct — `next-web-app/apps/{web,hiring,integration,maintenance}/package.json`
all pin `"next": "^16.2.7"`. *Correction:* Next.js 16.

---

### minor 8 — `architecture_overview.md:237`

> `*   **Directus**: Proxied via CMS service (business logic layer)`

Same class as BLOCKER 5, one bullet, in a file that fences cms extensively at `:145`/`:153-157`.
*Correction:* "Proxied via the cms **domain** inside `backend`".

---

### minor 9 — `architecture_overview.md:18`

> `*   **Roadrunner**: Code execution proxy (via Judge0 sandbox)`

In the PM-level "what runs after `make up`" list, unqualified. The container *does* start
(`docker-compose.yml:281`, `graphql` profile) so listing it is correct — but its stated *role* is false:
nothing calls it, and `backend` reaches Judge0 directly (`app/internal/jobsimwiring/wiring.go:118`,
`getenv("JUDGE0_BASE_URL")`). The file corrects this at `:186` and `:265-267`, so the PM list is the
outlier. Note `:15` gives jobsimulation a husk fence in this very list and `:18` gives roadrunner none.
*Correction:* add "— **orphaned husk**; `backend` calls Judge0 directly" to the bullet.

---

### minor 10 — `architecture_overview.md:33`, `:48`, `:63`

> `:33  *   **GraphQL/Cosmo Router**: API federation gateway`
> `:48  - **APIs**: GraphQL Federation v2 (WunderGraph Cosmo Router), gRPC/Connect-RPC …`
> `:63  3. **External Services**: Clerk, Directus, GraphQL, AI providers, LiveKit, AWS Chime`

Router listed without the prod-only qualifier in three summary lists. **Not a blocker** — the file opens
with the explicit ⚠ router banner at `:3` ("There is no `:5050` on a local stack"), the mermaid node at
`:73` is labelled "PROD ONLY — deleted from local dev", and the Tier-3 table at `:214` and the Request
Flow at `:256-263` are both correct. Flagged only because `service_taxonomy.md:303`/`:406` annotate the
equivalent rows inline and these three do not. *Correction:* append "(prod only)" to each.

---

### minor 11 — `dependency_map.md:13`

> `(compose `depends_on`, `docker-compose.yml:66-80` …)`

The `depends_on:` block is `:70-80`; `:66-67` are the trailing `SUPABASE_DB_CONN`/`COPILOT_DB_CONN` env
entries. Claim (sentinel + cms + storage) is correct. *Correction:* `:70-80`.

---

### minor 12 — `dependency_map.md:3`

> `inferred from configuration files (`docker-compose.yaml`)`

The file is `docker-compose.yml` (as `:7` and every citation in the table correctly write it).
*Correction:* `.yml`.

---

### minor 13 — `dependency_map.md:58`

> `| `backend` | App | CMS | User/org updates |`

Every other stream row in this table was updated to "App … App" post-merge; this one still names CMS as
the consumer. Literally defensible (the `cms` husk does still run with `BACKEND_STREAM=backend`,
`docker-compose.yml:161`), but inconsistent with `:59-62` and with the merge narrative directly below.
*Correction:* `App (cms domain in `app`; the `cms` husk also still subscribes until M810)`.

---

## 3. Per-file verdicts

- **`external_services.md` (693 lines)** — 4 BLOCKERs (1, 2, 3, 4) + 4 minors (1, 2, 3, 4). The router
  content (fact #4/#5) is **clean and well-fenced**: the `:3` banner, the `:374-376` HISTORICAL fence
  covering `:378-413`, and the corrected `:454`/`:461`/`:475`/`:670-682` recipes all check out against
  compose and `backend.go:315-317`. The damage is concentrated in the pre-merge cms-service residue and
  in the Clerk-webhook-goes-to-Sentinel error.
- **`service_taxonomy.md` (406 lines)** — 1 BLOCKER (5) + 2 minors (5, 6). Otherwise the strongest file
  in the group: every port, profile, and compose line number in the Tier-1 and Profiles tables verified
  correct against `docker-compose.yml` @ `2adcf71`, including the "six Go services plus Gotenberg" count,
  the `:381` profile-less base set, and the `:382` "profile name survives, the service does not" note.
- **`architecture_overview.md` (319 lines)** — **0 BLOCKERs**, 4 minors (7, 8, 9, 10). Fact #4 and #5
  handled correctly and in depth (banner `:3`, mermaid `:73`/`:113-115`, two-column Request Flow
  `:247-263`, Tier-3 `:214`). Clean verdict on the router; the minors are a stale framework version,
  one cms-proxy bullet, one unqualified roadrunner bullet, and three unqualified router mentions in
  summary lists.
- **`dependency_map.md` (103 lines)** — **0 BLOCKERs**, 3 minors (11, 12, 13). **Clean on substance.**
  Its `:23` router row and `:31` supergraph note are correct, and — notably — every one of its precise
  citations verified exactly: messenger `:255`/`:256`/`:258`/`:265`, studio-desk `depends_on` `:337-341`,
  husk lines `:83`/`:144`/`:281`, `JUDGE0_BASE_URL` at `wiring.go:118`, `GOTENBERG_URL` at `:51`. The
  three minors are two line/filename slips and one un-updated stream-consumer cell.
- **`README.md` (38 lines)** — **NO FINDINGS. Clean.** The index entries match their targets; `:19`'s
  mention of "the WunderGraph Cosmo GraphQL gateway" accurately describes what
  `external_services.md` still *documents* (as prod-only/historical) and is not a claim about local dev.

## 4. Totals

- **BLOCKERs: 5** — `external_services.md` ×4, `service_taxonomy.md` ×1.
- **minors: 13** — `external_services.md` ×4, `service_taxonomy.md` ×2, `architecture_overview.md` ×4,
  `dependency_map.md` ×3.
- **Files not fully read: none.** 1559/1559 lines read top to bottom.
- **Router / one-subgraph drift (ground truth #4, #5): 0 findings across all 5 files.** That sweep landed.
  The surviving drift is one release older and structurally grep-proof — the standalone-`cms`-as-Directus-proxy
  residue (BLOCKERs 4 and 5, minors 5 and 8) and a Clerk-webhook-to-Sentinel error (BLOCKERs 2 and 3)
  that each contradict a correct statement in their own file.
