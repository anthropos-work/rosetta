# The `anthropos-work` org repo register

**What this document is.** One row for **every** repository in the `anthropos-work` GitHub
organisation — **93 of them**, measured 2026-08-07 — with a *home* for each: a corpus doc, a row in
[`platform-migration-status.md`](platform-migration-status.md), or an explicit **"known, not
documented"** line. It exists because the corpus had no denominator for the org: it documented the
handful a stack has on disk — `repos.yml` lists **four** (`app`, `next-web-app`, `sentinel`,
`studio-desk`), with the frozen merged services sitting beside them — and had never enumerated the
rest, so a repo could be live, load-bearing and invisible at the same time. Several were.

**What this document is NOT.** It is **not a deletion plan.** Every `verdict` below is **advisory**,
and nothing here has been acted on: **no repo was archived, deprecated, or deleted.** Several verdicts
hinge on `infrastructure`, which is not in any stack's clone set; those are marked and left open.

## Provenance — state the invocation with every count

```
GET https://api.github.com/orgs/anthropos-work/repos?per_page=100&page=1..4&sort=pushed
    -> 93 repositories                      (2026-08-07, authenticated read)
```
**72 unarchived · 21 archived · 2 public (`rosetta`, `watermill-redisstream`) · 1 fork
(`watermill-redisstream`).** Every per-repo fact below was **re-derived from a clone at `origin` HEAD
taken on 2026-08-07** — not from the API, and never from a stack's working tree. That distinction is
load-bearing: the M257x claim census read clone *working trees*, 6 of 13 of which sat behind their own
fetched `origin/main`, and **a stale substrate manufactures evidence against a true claim.**

> ### ⚠️ The census that produced this register got several things wrong, and the corrections are the point
>
> A read-only 93-repo census produced a first draft of this table with **no local clone for ~70 of the
> 93**, and said so. Re-deriving each before writing it changed the finding in **five** places. They are
> recorded here rather than quietly fixed, because the pattern — *a repo that looks undocumented, or
> looks load-bearing, and is neither* — is the thing to be careful of:
>
> | Census said | Measured |
> |---|---|
> | `AI-Labs` is a control plane **the corpus does not know exists** | **The corpus knows.** [`ai-labs.md:4-8`](../services/ai-labs.md) names the repo, calls it the live Go control plane, and records that the doc *previously* denied it. Zero contradicted statements found |
> | the five `livekit-agent*` repos are undocumented | **They are enumerated** at [`platform-migration-status.md`](platform-migration-status.md) § census. What was wrong was two corpus sentences claiming *"no corpus document names them"* — **self-refuting**, since they say it while naming them |
> | `watermill-redisstream` is **upstream of the async backbone** | **REFUTED — it is an inert mirror.** `app/go.mod:12` and `colony/go.mod:9` both require **upstream** `github.com/ThreeDotsLabs/watermill-redisstream v1.4.5`, and **neither repo has any `replace`** redirecting it. The fork still declares `module github.com/ThreeDotsLabs/…` (so it is not importable as an org module without one), contains **zero** occurrences of "anthropos", and pins `watermill v1.2.0` against the platform's `v1.5.2`. "Abandoned fork", not "undocumented critical dependency" |
> | `sim-qa` is **a standing write-capable path into production** | **Write-capable: yes. Standing: no. Unmarked: no.** `ls .github` → absent: no workflow, no cron, no scheduler — it is developer-invoked. And its sessions are tagged `is_test=true` **by default** (`src/flow/scenario.ts:156`, `:245`), which `app` honours (`jobsimulations.graphqls:750`, `internal/jobsimulation/simulator/manager/manager.go:441`, column `is_test`). **The stale source of the wrong belief is sim-qa's own README** (`:36-42`, *"There is no `is_test` field… yet"*), contradicted by its own code |
> | `analytics-go` is **absent from the corpus** | Absent from the **library model**, not from the corpus — [`external_services.md:554`](external_services.md) already carried the correct `app/go.mod:14-18` enumeration. **The defect was that two corpus files disagreed and nothing reconciled them** |
>
> What the census got **right and understated**: `ant-observability` — the corpus documents **no
> observability tier at all** (`git grep -i grafana -- corpus/ CLAUDE.md` → **0 files**), and that repo
> holds both the platform's live outside-in monitoring and a **production read path** no safety doc
> enumerates. See [`observability.md`](../ops/observability.md).

---

## 1. The repos this corpus already owns

The stack's clone set and its documented neighbours. Their state lives in
[`platform-migration-status.md`](platform-migration-status.md) (fenced, per-service) and in
[`corpus/services/`](../services/README.md).

`app` · `sentinel` · `next-web-app` · `studio-desk` (the four in `repos.yml`) · `platform` ·
`ant-academy` · `anthropos-studio-room` · `rosetta` · `rosetta-extensions` · **`db-backup`** (added at
M257x close — it has a full service doc, [`db-backup.md`](../services/db-backup.md), and is one of the
**ten instantiated** production modules at `:571`, yet it appeared in **no** section of this register
while §3 measured it live and §12 called it *"one `terraform apply` away from being live again"*. A
doc promising *"one row for **every** repository"* owed it a home) · and the merged/frozen set
`cms`, `jobsimulation`, `storage`, `messenger`, `customerio-sync`, `roadrunner`, `skiller`, `skillpath`,
`chronos`, `intelligence`, `graphql-wundergraph` · plus the libraries `colony`, `proto`, `taxonomy`,
`ai`, `authn` and — **added M257x iter-123** — `analytics-go` (see
[`shared_libraries.md`](shared_libraries.md), whose subject set is **not** `app`'s require set).

---

## 2. Live, load-bearing, and newly given a home

Each of these is **alive**, has a real role in the running platform or its operation, and had **no
corpus row of substance** before 2026-08-07.

| repo | what it is (measured) | verdict | home |
|---|---|---|---|
| **`infrastructure`** (HCL, pushed 2026-08-07) | The **Terraform monorepo** — every AWS resource, account `583848331406`, workloads `eu-west-1`. `terraform/production/services.tf` (666 lines) sources **9** service repos' `//terraform` modules at pinned tags; those modules in turn source `infrastructure.git//modules/services/base_service`. Releases via `upgrade-service.sh`, which `sed`s a `//VERSION:<NAME>:` marker. **`terraform/stage/main.tf` is 0 lines — there is no staging environment.** ⚠️ **Everything this corpus asserts from it was read at `13c248e6` (M257x iter-123), and `infrastructure` is NOT in the standing clone set** — `git ls-remote` confirms `13c248e6` is origin `HEAD` (M257x iter-132), so the readings are sound, but no `stack-*/` holds it and `make init` does not clone it, so the ~31 corpus lines citing it are **not re-derivable in place**. **That is a fact about our habits, not a limit on measurement, and the corpus spent four iterations treating it as the latter** — iter-132 re-cloned the repo in under a minute and settled a second standing hedge from it (the production RPC address; [`backend.md`](../services/backend.md)'s *RPC re-pointed, then un-set* bullet). **When a claim turns on this repo, clone it — do not hedge it.** | **KEEP — and it is the platform's authoritative deployment ledger** | § 3 below + the `cms`/`roadrunner`/`storage` rows of [`platform-migration-status.md`](platform-migration-status.md) |
| **`directus`** (HCL, 2026-06-05) | **Not a fork of Directus.** The CMS container is stock `directus/directus:11.6.1`; this repo builds **four in-house extensions** and copies them into a shared volume (`Dockerfile:32`, `docker-compose.yml:9-19`), which is why `infrastructure` passes it as **`setup_docker_image`** (`services.tf:32`), not `docker_image`. | **KEEP — and a named gap** | § 4 below; [`directus-local.md`](../ops/directus-local.md) |
| **`judge0`** (2026-02-19) | A **vendored copy** of upstream Judge0 CE **v1.13.1** (`judge0-master/` is an unzipped upstream tree — 7 commits all dated 2025-12-09, no upstream remote, no submodule), plus the org's customisations **beside** it. **The only existing record of how the production Judge0 host is built.** | **KEEP — irreplaceable** | § 5 below |
| **`metabase`** (HCL, 2026-05-27) | **Terraform only, no application code**: deploys stock `metabase/metabase:v0.56.8` as a live ECS service (`service_desired_count = 1`) on its own ALB hostname, pinned by `infrastructure` at `v0.2.4`. Connected to the **production Postgres** with deliberately small pools and a 2-minute query timeout (`terraform/main.tf:43-76`). | **KEEP — and it is a security surface** | § 6 below |
| **`AI-Labs`** (2026-08-04) | The **`labs-api` control plane** — two **Go** modules, ~8 kLOC, **stdlib-only** (neither `go.mod` has an external `require`), orchestrating **Firecracker** microVMs behind `anthroposlabs.com`. Its Python is *scenario content*, not platform. | **KEEP** | [`ai-labs.md`](../services/ai-labs.md) — **already correct**; § 7 adds only where it runs |
| **the five `livekit-agent*` repos** | `livekit-agent`, `livekit-agent-chain`, `livekit-agent-azure-eu`, `livekit-agent-azure-eu-fr`, `livekit-agent-azure-us` — all Python, all one LiveKit Cloud project (`subdomain = "anthropos-pbvktu3v"`), five distinct registered agent names. **`app` dispatches exactly three.** | **KEEP 3 · DECIDE 2** | § 8 below |
| **`ant-observability`** (Shell, 2026-08-05) | Two things in one repo: Claude Code usage telemetry (OTel→Prometheus/Loki/Grafana), **and** `product-monitoring/` — the platform's live outside-in monitoring, on a Proxmox VM reachable only over Tailscale. | **KEEP — the corpus's largest single blind spot** | [`observability.md`](../ops/observability.md) |
| **`sim-qa`** (TypeScript, 2026-07-31) | A **developer-invoked** QA harness that drives production GraphQL as a real user, issuing **7 mutations**. Needs a production `sk_live_` Clerk key and mints a JWT against a real user's session. **Sessions are `is_test=true` by default.** | **KEEP — with its scope stated** | § 9 below + a scope note on [`safety.md`](../ops/safety.md) |
| **`hyper-studio`** (TypeScript, 2026-08-06) | The most actively developed repo in the org (827 commits since 2026-06-17). A CLI suite of content-creation agents around **HyperForge**, self-declared *"peer to `studio-desk` and `anthropos-studio-room`"*. **Zero runtime coupling to the platform today** — no GraphQL, no Clerk, no DB, no deployment. | **KEEP — PRE-INTEGRATION** | § 10 below |
| **`analytics-go`** (Go, 2025-02-12) | A two-file Brevo event tracker, **`app/go.mod:14` `v0.3.1`** — a direct compile-time dependency carrying **Stripe subscription-lifecycle events** (`app/internal/payments/handler.go:302-316`). Untouched for ~18 months; `v0.3.1` is its newest tag. | **KEEP — DORMANT AND LOAD-BEARING. Do not delete** | [`shared_libraries.md`](shared_libraries.md) |
| **`anthropos-knowledge-base`** (2026-08-06) | A Claude Code plugin + company knowledge base — **and it contains a second, parallel platform-architecture corpus** covering this project's subject. | **KEEP — and RECONCILE** | § 11 below |
| **`github-runner-config`** (Shell, 2026-06-26) | 2 files / 63 lines. A non-idempotent `init.sh` that bootstraps an Ubuntu box for a self-hosted GitHub Actions runner. Registers no runner, provisions nothing. | **KEEP — known, not documented in depth** | this row. Its value is the *fact* it records: **CI runs on self-hosted EU runners reaching AWS over Tailscale** (`CLAUDE.md:11-12`). **Not a runbook** |
| **`watermill-redisstream`** (Go, 2024-03-21) | The org's **only fork**, public, single commit, **pristine upstream** — `grep -ri anthropos` → empty. **Not a dependency of anything** (see the correction table above). | **DEPRECATE-CANDIDATE — inert** | this row. Recorded so the next reader does not re-spend the hour proving it inert. **Archiving it is safe on the measurement; that call is the owner's** |

---

## 3. `infrastructure` — and the `cms` question it settles

Two facts about the relationship, because the shorthand *"sourced by 9 repos"* is true of **two
different sets of nine** and both are clone-set artifacts:

- **Forward** — `terraform/production/services.tf` sources 9 repos at pinned tags: `sentinel` (`:2`),
  `directus` (`:24`), `storage` (`:131`), `next-web-app` (`:197`), `app` (`:256`), `jobsimulation`
  (`:476`), `studio-desk` (`:529`), `db-backup` (`:571`), `metabase` (`:592`).
- **Reverse** — service modules source `infrastructure.git//modules/services/base_service` or
  `base_internal_service`. Counted over the 12 terraform-bearing repos on disk: **11 do**
  (`next-web-app` is Vercel; `db-backup` composes no module). Nine of those 11 happen to be the
  `stack-demo` clone set, five of which are decommissioned repos.

**The durable statement is the mutual pin, not the number.**

> ### 🔓 `cms` M810: SETTLED — the ECS service is DESTROYED
>
> This corpus recorded a standing *"do not assert either way"*: `cms/terraform/main.tf:39` still reads
> `service_desired_count = 0`, while `6efa1d5` deleted the build workflow saying *"the cms ECR repository
> is decommissioned (M810)"* — two measured facts pointing opposite ways, with the blocker stated as
> *"the destruction itself happens in infrastructure's `services.tf`, which we cannot read."*
>
> **It has now been read.** There is **no `module "cms"` declaration** anywhere in `infrastructure`
> (`grep -rn 'module "cms' terraform/ modules/` → zero). In its place, `services.tf:64-70`:
>
> > *"M810: cms was removed (module block deleted here). cms is folded into the backend (cms-in-app v8.0,
> > app v1.360.0)… the standalone service ran at desired_count = 0 as the rollback path and is now
> > decommissioned. Deleting the module destroys its ECS service, task definition, ECR repository, IAM
> > roles, security group, Cloud Map entry, log group, alarms and the ten `/production/cms/*` SSM
> > parameters."*
>
> plus `:88-94`, a `removed { … lifecycle { destroy = false } }` for the Atlas tracker, and `:85-86`
> noting the legacy `cms` **schema** is untouched — **that drop remains a separate, still-pending M810
> step.** So `cms/terraform/main.tf:39` is not a contradiction; it is **orphaned dead code** — no root
> module instantiates that file. **The deleted workflow was the correct signal.**
>
> The same file settles neighbours in the platform's own words: `storage`'s module **deliberately
> survives** as assets-only because `prevent_destroy` is read from configuration, not state (`:107-129`);
> `customerio-sync` is fully deleted, no `removed` block, owned no data (`:145-164`); the `skiller`,
> `messenger` and `storage` ECR repos were forgotten then hand-deleted **2026-08-05** (`:646-666`).

> ### 🔑 The rule the whole exercise was worth: **a service repo's `service_desired_count` is not evidence of production state**
>
> It is a **module input inside a module nobody may be calling.** It means something only when a root
> module in `infrastructure/terraform/production/services.tf` instantiates that module — and **four of
> the repos this corpus reads declare a count that instantiates nothing.**
>
> Measured at `infrastructure` **`13c248e6`** — `grep -n '^module "' terraform/production/services.tf`
> returns **exactly ten** declarations: `sentinel_euwest1` (`:1`), `directus_euwest1` (`:23`),
> `acm_media_certificate_useast1` (`:96`), `storage-service_euwest1` (`:130`), `next-webapp_euwest1`
> (`:196`), `backend_euwest1` (`:252`), `jobsimulation_euwest1` (`:475`), `studio_desk_euwest1` (`:528`),
> `db-backup-euwest1` (`:570`), `metabase_euwest1` (`:591`). That list **is** the production service set.
>
> | repo, and the line this corpus has been citing | what it describes | measured in `infrastructure` @ `13c248e6` |
> |---|---|---|
> | `cms/terraform/main.tf:39` `= 0` | *"the rollback path"* | **orphaned** — no `module "cms_euwest1"`; `services.tf:64-70` records the deletion and what it destroyed |
> | `roadrunner/terraform/main.tf:19` `= 1` | *"the prod contradiction… still not verified"* | **orphaned** — **`roadrunner` appears in NO terraform in the repo** (7 org-wide hits, all `judge0_*` secret names still labelled "Roadrunner" in two CI workflows, plus one KB line). There is no roadrunner ECS service. **Re-derived independently at M257x iter-137 and UPHELD byte-for-byte** — and the same read supplied the *positive* half this row never had: those secrets feed `TF_VAR_judge0_{api_key,base_url}` (`infrastructure/.github/workflows/wf-terraform-deploy.yml:209-211`) into **`module "backend_euwest1"`** (`infrastructure/terraform/production/services.tf:384-385`), so **production wires Judge0 straight into `backend` under roadrunner-named keys — the fold, visible at the config layer.** `infrastructure/knowledge/service-dependencies.md:119` says it in the platform's own words: *"Judge0 (code execution — called directly now; `roadrunner` is off this path)"* |
> | `graphql-wundergraph/terraform/main.tf:20` `= 1` | *"the router — **still declared**"* in production | **orphaned** — `module.wundergraph_euwest1` **is deleted**; `services.tf:509-517` says the apply destroyed *"its ECS service, task definition, target group, ALB rule (priority 810), Cloud Map entry, log group, ACM cert and the `wundergraph.anthropos.work` alias"*, with only a `removed{}` for the ECR (`:521`), hand-deleted **2026-08-05** — *"so production-wundergraph is gone and this block is now inert"* |
> | `messenger/terraform/main.tf:29` `= 0` | *"scaled to zero as the rollback path"* | **orphaned** — `module.messenger_euwest1` is deleted (`services.tf:622`); only its ECR `removed{}` survives (`:664`) |
>
> **One class, four standing corpus puzzles, all dissolved by the same fact.** This corpus reconstructed
> production from service-repo fragments because it had no access to the root module — and a fragment of
> a module nobody calls reads exactly like live configuration. **Never grade a production claim on a
> service repo's own terraform. Grade it on whether `services.tf` instantiates the module.**
>
> ### And the two that legitimately survive, for the same stated reason
>
> In `infrastructure` @ `13c248e6`, `storage-service_euwest1` (`:130`, `ref=v0.15.8`) and
> `jobsimulation_euwest1` (`:475`, `ref=v0.254.0`)
> **are still declared, deliberately, and are no longer service modules.** Both keep only ASSETS: for
> storage, both S3 buckets, versioning/SSE, the CloudFront distribution + OAI + bucket policy and the
> `media.anthropos.work` CNAME; for jobsimulation, the LiveKit/Chime recording buckets `backend` reuses
> **by literal name**, the Chime SNS topic + S3 notification webhooking into
> `api.anthropos.work/api/webhook/chime`, the `/production/jobsimulation/*` SSM parameters and the atlas
> tracker for the legacy schema. **Both had their ECR repositories destroyed outright** (images
> hand-deleted first) rather than forgotten.
>
> `services.tf:113-118` gives the reason deletion would be wrong, and it is worth carrying verbatim
> because it inverts the intuition: deleting the block *"would destroy every one of those, and the
> `prevent_destroy` guards on them would NOT stop it: `prevent_destroy` is read from **CONFIGURATION**,
> not state, so removing the block removes the guards along with the resources they guard. Keeping the
> assets declared is what keeps those guards load-bearing."*
>
> ⚠️ **`services.tf:122-124` (in `infrastructure` @ `13c248e6`) carries a name-collision warning that
> belongs in any runbook touching this repo:** do not confuse `module.storage-service_euwest1` (the object-storage assets) with
> `module.storage_euwest1` (`modules/core/storage` — **the RDS instance and the ElastiCache replication
> group**). *"A `terraform state` command aimed at the wrong one destroys the platform database. Never
> abbreviate these two module names and never target them with a prefix or wildcard."*

**`services.tf` is where every decommission is justified in prose next to the block that implements
it.** Cite it directly as the production-side source, in preference to reconstructing prod state from
service-repo fragments. *(Its own KB drifts too: `knowledge/architecture.md:41-43` still says core
modules come from a consolidated-away `infrastructure-modules` repo; measured, they are local relative
paths.)*

---

## 4. `directus` — four extensions no stack has ever had

| Extension | Type | What it does |
|---|---|---|
| `directus-extension-ant-skill-path-uuid` v1.0.2 | `interface` (Vue) | A `uuid-generator` field: mints a `uuidv4()` when the field is empty (`src/index.ts:4-12`, `src/interface.vue:18-22`) |
| `directus-extension-image-import` v1.0.0 | `operation` | Flow op `file-import` → `FilesService.importOne(url, {folder})` (`src/api.ts:9-53`) |
| `directus-extension-metalink` v1.2.1 | `bundle` → operation | Fetches **OpenGraph** metadata for a URL (`src/…/api.ts:6-11`) |
| `directus-extension-youtube-meta` v1.0.2 | `bundle` → operation | Resolves **YouTube** metadata, authenticating with `env.GCLOUD_SERVICE_ACCOUNT` (`src/…/api.ts:6-13`). ⚠️ **`:9` is a bare `console.log(env)` at the pin production runs**, so an invocation writes the DB password, Directus signing secret, admin password and Google client secret to the log group — filed as `PLATFORM-M257x-directus-ext-logs-env` |

**They exist nowhere else** (0 hits for each name across `corpus/`, `.claude/`, `stack-demo/`,
`stack-dev/`) and **no stack installs them** — the per-stack Directus emitter
(`rosetta-extensions/stack-injection/gen_injected_override.py:321-364`) pins the stock image and emits
**no `volumes:` key at all**; the dev twin does the same (`dev-stack/dev-setdress.sh:255`).

**This is an authoring-time gap, not a rendering gap** — which is exactly why it has gone unnoticed. A
stack is set-dressed by *replaying already-authored content*, never by authoring it, so demo fidelity is
unaffected. It bites the moment content is authored **in** a stack — notably via `studio-desk`, which
writes skill paths to the per-stack Directus:

1. the `uuid-generator` interface does not resolve → the field degrades to a **plain input**, so a
   record can be saved with an **empty or hand-typed id** instead of a v4 UUID. A silent data-shape
   divergence from prod, **not an error**;
2. any Flow step calling `file-import` / `ant-metalink-operation` / `ant-youtube-operation` **fails to
   resolve**. Replayed content that already holds the fetched metadata renders fine; re-running the flow
   does not.

**Do not author content in a stack's Directus and expect prod-shaped records.**

---

## 5. `judge0` — the one production box that is not infrastructure-as-code

The customisation is **not applied to the vendored source**; it sits beside it.

- `extra/compilers/Dockerfile` layers onto `judge0/compilers:1.6.0-extra`: apt repointed at
  `archive.debian.org` (the base is EOL Debian, `:4-6`), a **Python ML stack** — numpy 1.23.5, pandas
  1.3.5, scikit-learn 1.0.2, scipy 1.9.3, matplotlib 3.5.3, mlxtend 0.21.0, pytest 7.1.3 (`:10-17`) —
  **OpenMPI + `mpi4py`** (`:21-24`) and **.NET SDK 3.1** (`:46-56`).
- A **manual post-deploy SQL statement** repoints Judge0 language id **71** (Python) at that
  interpreter (`README.md`). **Un-automated, non-idempotent, and silently lost on a DB rebuild.**

There is **no Terraform in the repo** and `infrastructure` does not source it. It reaches the platform
as an opaque tfvar: `infrastructure/terraform/production/variables.tf:523` `judge0_base_url` →
`services.tf:385` → `app/terraform/main.tf:638-639` `JUDGE0_BASE_URL`, consumed in-process at
`app/internal/jobsimwiring/wiring.go:123`. Deployment is manual — build locally, `docker save` /
`docker load` onto the VM, `rails db:seed` — onto **Ubuntu 22.04 that must carry
`systemd.unified_cgroup_hierarchy=0` in GRUB**, with both containers `privileged: true`.

> ⚠️ **Caveat when using this repo as a runbook:** its `docker-compose.yml` contradicts its own README on
> image names (`judge0/judge0:latest` + `1.6.0-extra` vs the README's `judge0/anthropos` +
> `judge0/compilers-custom`). The single recorded artifact is internally inconsistent about which images
> the box runs.

**The corpus's existing statement — *"`backend` calls Judge0 directly via `JUDGE0_BASE_URL`"* — is
correct and stops one hop short.** What was never said is what is on the other end of that URL.

---

## 6. `metabase` — a BI console on the production database, outside Sentinel

Live: `service_desired_count = 1`, `use_fargate = false`, port 3000, cpu 1024 / mem 2048, health check
`/api/health`, **its own DNS + ALB listener rule** (`terraform/main.tf:13-33`, `locals.tf:1-5`,
`services.tf:603`). App DB URI in an SSM SecureString (`main.tf:73-75`). Throttled against the platform
DB on purpose: `MB_JDBC_DATA_WAREHOUSE_MAX_CONNECTION_POOL_SIZE=3`,
`MB_APPLICATION_DB_MAX_CONNECTION_POOL_SIZE=3`, `MB_DB_QUERY_TIMEOUT_MINUTES=2` (`:59-69`).

**It reads the production platform database outside the Sentinel/Casbin authorization layer, so
Metabase's own permissions are the only tenant boundary on that surface.** That belongs in the reader's
model alongside [`db-access.md`](../ops/db-access.md) and
[`security_compliance.md`](security_compliance.md), which describe a read surface this one does not
pass through.

---

## 7. `AI-Labs` — where `labs-api` actually runs

[`ai-labs.md`](../services/ai-labs.md) is **correct** about what labs-api is and how `app` consumes it.
What it never said is **where it runs**, and the answer is unusual enough to matter:

- **No Terraform in the repo** (`find . -name '*.tf'` → 0). Deployed by **Ansible + systemd**
  (`ops/ansible/files/{coding-sim-cp,labs-node,cloudflared,ailabs-firewall}.service` + a `Caddyfile`)
  onto a **single tailnet VM** — `STATUS.md:3-6`: host `ailabs`, `100.120.254.65`, **single-worker,
  "Hetzner fleet still pending"**.
- Ingress is **Cloudflare Tunnel + Caddy** (`ops/ansible/setup-worker.yml:2`), with per-session dynamic
  routing via the Caddy admin API (`labs-node/caddy.go:5`; `CP_CADDY_ADMIN_URL` default
  `http://127.0.0.1:2019`).
- The **DNS is** Terraformed — in the other repo:
  `infrastructure/modules/domains/anthroposlabs-com/`.
- Go orchestrates; a **shell script boots the VM**: `labs-node/firecracker.go:52` execs
  `ops/10-boot-one.sh`. VM binaries (firecracker, vmlinux, rootfs) are **not in git**.
- The scenario catalog is **96 templates**, each with a `meta.json` — measured. **Do not copy the repo's
  own README or `STATUS.md:11`, both of which say "15"**; `STATUS.md:3` is dated 2026-04-27 against a
  2026-08-04 HEAD. The GitHub description says 15 too. *A repo's self-description is testimony; the tree
  is evidence.*

---

## 8. The `livekit-agent*` family — one absorbed agent, two dead names

| repo | `livekit.toml` agent | registered name | dispatched by `app`? |
|---|---|---|---|
| `livekit-agent` | `CA_gKnT9pr6QacU` | `anthropos-agent` | **yes** — the production default **and** the EU path |
| `livekit-agent-chain` | `CA_4ncLNCfb2kmA` | `anthropos-agent-chain` | **yes** — the `livekitchain` engine |
| `livekit-agent-azure-us` | `CA_eWTSTw6RkiXZ` | `anthropos-agent-us` | **yes** — `location=us` |
| `livekit-agent-azure-eu` | `CA_wDTotEUgnef8` | `anthropos-agent-eu` | **no** |
| `livekit-agent-azure-eu-fr` | `CA_ypfzhzwVTLvU` | `anthropos-agent-eu-fr` | **no** |

**The absorption is the finding.** Routing moved from *one repo per region* to **one agent whose
endpoint is chosen by dispatch metadata**: `app` sets `metadata["endpoint"]`
(`internal/jobsimulation/calls/livekit.go:148`) — default `azure-eu` (`:111`), a **random** pick from
`euAgentEndpoints` (`:101-104`, `:122-123`), `azure-us` (`:127`), `openai-hosted` (`:143`) — and
`livekit-agent/src/agent.py:268` branches four ways on it (`:283`, `:301`, `:319`, `:337`, else
`raise`). The three azure-* repos hard-wire a single endpoint and **never read that metadata**
(`grep endpoint` in `livekit-agent-azure-eu-fr/src/agent.py` → 0 hits). That is **why** `azure-eu` and
`azure-eu-fr` went dark — they were absorbed, not abandoned.

**And the absorption is incomplete, which is a live inconsistency:** `livekit-agent` *can* serve
`azure-us` (`:319-335`), yet dispatch still routes US to the separate repo. On a US session `app` sets
**both** `agentName="anthropos-agent-us"` **and** `metadata["endpoint"]="azure-us"` (`livekit.go:126-127`)
— and the receiving agent ignores the endpoint entirely. One of the two signals is inert.

**`livekit-agent-chain` is not a successor — it is a different architecture.** Every other agent is
speech-to-speech `openai.realtime.RealtimeModel`; chain is a classic **STT → LLM → TTS** pipeline
(`src/agent.py:607-621`: `inference.LLM("google/gemini-2.5-flash")`,
`inference.TTS("elevenlabs/eleven_multilingual_v2")`). It is selected **by CMS content, not a flag** —
`voiceEngine == SimulationVoiceEngineLivekitchain`, enum `"livekitchain"` at
`app/internal/cms/directus/collections/jobsimulation.go:1085`. It is the actively developed one.

> ⚠️ **The cross-repo coupling is an unfenced string contract.** The platform holds **no** reference to
> any agent repo — no URL, no `CA_*` id, no `livekit.toml`. The whole seam is the agent-name literal plus
> the metadata keys (`endpoint`, `voice`, `prompt`, `first_message`, `call_duration`, `language`,
> `language_code`). **A typo on either side fails at dispatch time, in production, with no build-time
> signal.** A natural candidate for a guard.
>
> ⚠️ **Do not source corpus text from these repos' READMEs.** All three azure-* READMEs (`:8`) call
> `livekit-agent` the *"dev/OpenAI-hosted variant"* — contradicted by `livekit.go:110,120`, where it is
> the production default. `livekit-agent-chain`'s README is still the **unmodified upstream LiveKit
> starter** (`pyproject name = "agent-starter-python"`).

---

## 9. `sim-qa` — a real prod-write path, correctly scoped

**Seven mutations, all resolving in `app`'s current schema**
(`app/internal/web/backend/graphql/graph/schemas/jobsimulations.graphqls`): `startSession` (`:744`),
`sendTextMessage` (`:753`), `sendAIMessage` (`:871`), `completeJobSimulationTask` (`:807`),
`completeSessionWithValidationAttemptResult` (`:762`), `abortSession` (`:761`),
`getJobSimulationLivekitCallToken` (`:821`). These create sessions, drive LLM inference, complete tasks,
trigger evaluation, and abort other sessions.

**Prod is the hardcoded default, not opt-in** — `src/env.ts:24` and `src/cli/abort-active.ts:49` both
fall back to `https://wundergraph.anthropos.work/graphql`, so an empty `.env` still targets production.
It authenticates **as a real production user**: `env.example:2-4` requires an `sk_live_` Clerk secret
key, and `auth/clerk.ts:122-153` resolves an arbitrary user by email and `:96-120` mints a JWT against
their session. **No direct DB access** — every write goes through the application layer.

**Does this contradict [`safety.md`](../ops/safety.md)? No — and the fair statement is that safety.md is
narrower than a reader will assume.** That document is explicitly scoped to Rosetta's own tooling
(`:3`, `:36-37`); every absolute-sounding sentence has *"a demo"* or *"the tooling"* as its subject.
sim-qa is a different team's repo. **Nothing in `safety.md` is falsified.** What *is* worth naming is
that [`db-access.md:12`](../ops/db-access.md) — *"nobody queries with a write account"* — is literally
true (sim-qa holds no DB account) while contemplating only the DB tier; it does not anticipate a prod
write achieved **through the application layer with a real user's JWT**. That is a gap in the threat
model, not a contradiction.

> **Latent, and worth checking before next use:** sim-qa's default endpoint is the **decommissioned**
> Cosmo router — dropped 2026-07-31. Its last commit predates that by two months. Whether that host still
> redirects during a deprecation window is **not measurable from a clone set**; if it does not, sim-qa is
> non-functional out of the box.

---

## 10. `hyper-studio` — pre-integration, and the corpus already borrows from it

A plain TypeScript/Node CLI — no framework, **four runtime deps** (`@anthropic-ai/claude-agent-sdk`,
`dotenv`, `yaml`, `zod`). It is the designated successor to **studio-room** (generation), *not* to
studio-desk — there is no UI code at all. Its own record says why:
*"the legacy is hardcoded because at the time there was no way to make it dynamic. It was built for
**creation only — edit was never possible**. HyperForge exists precisely to close that gap."* A second
agent, **HyperPlay**, is specced to drive content on the platform and is **unbuilt**.

**It does not touch the platform today**: a combined grep for `wundergraph|anthropos.work|/graphql/query|clerk|DATABASE_URL|livekit|directus|judge0|bunny|chime|postgres` over `origin/main` returns **3 hits,
all prose, zero in code**; `grep -inE graphql` → 0 repo-wide. No `.github/`, no Dockerfile, no compose,
no terraform — **no deployment exists.**

**Filing it as a live Tier-2 service would misstate the architecture**; leaving it out entirely already
misstated something else — [`secrets-spec.md:309`](../ops/secrets-spec.md) uses
`../hyper-studio/.env.example` as the template for `app`'s five AWS Bedrock secret genes. **The corpus
borrows a file from a repo it never mentions.**

---

## 11. `anthropos-knowledge-base` — the second corpus, and what it contradicts

A hand-authored Claude Code plugin + company KB (2,434 tracked files). **Six files overlap this
corpus's subject**, all under `knowledge/`: `07-technical-architecture.md`, `07a-architecture-overview.md`,
**`07b-microservices-catalog.md` (376 lines — the direct counterpart)**, `07c-infrastructure-diagram.md`,
`07d-security.md`, `07e-ai-architecture.md` — **≈1,773 lines against this corpus's 34,417 across 90
files.** Of this corpus's 27 service docs: **18 have a full counterpart, 6 product-level only, 3 absent**
(`askengine`, `clerkenstein`, `gotenberg`). AKB has **no** counterpart to `corpus/ops/` or
`corpus/architecture/`.

**Three same-subject comparisons, and they do not all go the same way:**

1. **WunderGraph router — agree on the drop, contradict on residue.** AKB `07b:199` says *"Residue |
   Nothing. The ECS module, ALB rule (priority 810), Cloud Map entry and `wundergraph.anthropos.work`
   were destroyed at retirement"*; this corpus says the module is **still declared** at
   `graphql-wundergraph/terraform/main.tf:20` `service_desired_count = 1`. **RESOLVED AT M257x iter-124,
   IN AKB'S FAVOUR: AKB was right and this corpus was wrong** — § 3 above measured
   `module.wundergraph_euwest1` as **deleted** from `infrastructure/terraform/production/services.tf`
   @ `13c248e6` (`:509-517`), and the count this corpus quoted is orphaned dead code. **This item read
   *"Unresolved"* until iter-125, one screen below the § 3 measurement that had already settled it** —
   the same one-cell-reach failure iter-124 found 24 more of. AKB was better positioned for a structural
   reason worth keeping: **it read the repo this corpus had never cloned.**
2. **The taxonomy figures — direct contradiction, and this is the headline.** AKB asserts *"60,000
   skills… mapped to 18,000 roles"* in **14 places**, **citing no source anywhere**. This corpus marks
   **"18K roles" REFUTED** (measured **22,470** public job roles — public ⊆ total, so 18K is below the
   floor) and **"60K skills" UNVERIFIED** (42,790 public), with dated, reproducible provenance
   ([`shared_libraries.md`](shared_libraries.md#taxonomy-figures)). **This corpus wins on evidence** —
   and the consequential part is that the unsourced figure is load-bearing in **four customer-facing
   competitor-comparison tables**.
3. **How many services folded into `app` — AKB contradicts itself.** Two summary banners say five
   (`07b:7-11`, `07a:13-18`); its own body lists all **eight** (`07b:242-244` for storage, messenger,
   customerio-sync) and its own diagrams agree. **Substance agrees with this corpus at eight**; only the
   banners are stale.

> **AKB's `Last updated` headers are unreliable and must not be trusted** — `07b:3` says "July 2026"
> while its body cites 2026-08-05 events; `07a:3`/`07c:3` say "February 2026" yet carry the 2026-07-31
> router drop.

**How the two relate, stated so a reader can choose.** This corpus is authoritative for **measured
local/runtime state and ops**; AKB is better positioned on **`infrastructure`-derived production state**
and on product/GTM. **Neither cites the other.** This corpus already names AKB in five places
(`setup_guide.md:401`, `staging-bringup.md:125`, `:151`, `staging-sync.md:112`,
`toolchain_overview.md:91`) — **every one of them as a repo to clone or sync, never as a documentation
source** — while `toolchain_overview.md:92-94` records that its plugin injects *"full Anthropos context
(product details, **architecture**…)"* into every engineer's editor. **So a reader following this
corpus's own advice installs a plugin that serves the 60K/18K figure this corpus refutes.** That is the
reconciliation debt, and naming it is the first payment.

---

## 12. Everything else, with a verdict

Nothing below is proposed for deletion; the verdicts are advisory and several are open.

**Already archived on GitHub (21) — no action needed, recorded for completeness:**
`roadrunner` · `skiller` · `skillpath` · `graphql-wundergraph` · `infrastructure-modules` · `realtime` ·
`intelligence` · `skill-gateway` · `chrome-extension` · `gen-ai-pipeline` · `simulator` · `simulator-mx` ·
`web-app` · `landing-page` · `wordpress` · `blog-home` · `resume-parser-go` ·
`anthropos-work.github.io` · `blog` · `chrome-extension-old` · `website`.

**Knowledge bases (live, out of this corpus's scope but named so they are not "unknown"):**
`kb-ant-product` · `kb-ant-business` · `kb-certifications-iso27001` (relevant: `security_compliance.md`
cites no ISO-27001 programme) · `kb-migration-plan` · `kb-domain-singularity` · `ant-singularity`
(the singularity node — see [`staging-bringup.md`](../ops/staging-bringup.md)).

**Live but unexamined — "known, not documented", and each is a real open question:**
`customerio-sync` (frozen, documented) · `auth` (measured at iter-01 as a 3-commit, 8-hour spike marked
*"Not yet deployed"*, all activity stopped 2026-06-18 — **Clerk is not being replaced**) ·
`simulation-form` · `demo-environment` · `clerk` · `experiments` (see
[`anthropos-labs.md`](../tools/anthropos-labs.md)) · `customer-orbyta` (the corpus describes no
per-customer repo pattern) · `Analytics-and-Reports` · `bench-analysis-transcripts` · `transcoder`
(adjacent to `media-substrate-spec.md`'s Bunny path) · `realtime-python` · `studio-tools`.

**Dormant ≥ 18 months, unarchived — DECIDE, and the decision needs `infrastructure` read first:**
`ant-content-extension` · `helm-charts` · `flux` · `mattermost-sentry` · `k8s-infrastructure` ·
`skills-and-job-roles` · `skill-index-data` · `issues` · `taxonomy-generation-tool` · `anthropos-ios` ·
`blueprint-skill-taxonomy` · `documentation` · `flagsmith` · `blueprint-job-simulation` ·
`blueprint-anthropos-studio` · `infrastructure-legacy` · `supabase-edge-functions`.

> **Why "read `infrastructure` first" is not a formality.** `helm-charts`, `flux` and
> `k8s-infrastructure` describe a Kubernetes deployment model the platform no longer uses (it is ECS)
> — but *no longer uses* is an inference from `services.tf`, not from those repos, and three of the
> census's twelve decide-laters turn on exactly that file. **`db-backup` is the standing warning**: it
> looks abandoned, is still pinned by production, and is one `terraform apply` away from being live
> again.

## Related
- [`platform-migration-status.md`](platform-migration-status.md) — per-service prod-vs-local state (fenced)
- [`shared_libraries.md`](shared_libraries.md) — the library model, and why its subject set ≠ `app`'s requires
- [`observability.md`](../ops/observability.md) — the tier this register found missing
- [`corpus/services/README.md`](../services/README.md) — the service-doc index
