# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

**Project Rosetta** is the documentation corpus for the Anthropos platform. It serves three purposes:
1. **Documentation Repository**: Comprehensive architecture guides for developers
2. **Environment Setup**: Manual for humans and AI agents to build local development environments
3. **Recursive Inspection**: Tool for reverse-engineering and documenting the platform itself

This is NOT the Anthropos platform source code - it's the documentation about it. The actual platform code lives in separate repositories under the `anthropos-work` GitHub organization.

## Development Commands

### Available Skills

| Skill | Purpose | Guide |
|-------|---------|-------|
| `/dev-up` | Build / start / set-dress a dev stack (consolidates the former setup-platform + start-platform; drives the M13 dev set-dress flow) | `corpus/ops/setup_guide.md` + `corpus/ops/run_guide.md` |
| `/dev-down` | Tear down an additional dev stack (`dev-N`, N ≥ 1) — frees its registry slot | `corpus/ops/rosetta_demo.md` |
| `/dev-for-dummies` | One-command LIVE-development setup (or resume): pins rosetta + rosetta-extensions to clean main, brings up a remote-accessible demo stack (`demo-N` over Tailscale `--public-host`), and runs each TARGET repo natively with hot-reload from an isolated `feat/<name>`/`fix/<scope>` worktree | `.claude/skills/dev-for-dummies/SKILL.md` |
| `/setup-github` | Configure GitHub SSH access for the org | `corpus/ops/setup_github_guide.md` |
| `/update-knowledge` | Document new evidence across the corpus | N/A (meta-skill) |
| `/test-platform` | Verify a running platform (probes, repo suites, census) | `.claude/skills/test-platform/SKILL.md` |
| `/db-query` | Query the prod Postgres read-only (investigate data, size/inspect surfaces) | `corpus/ops/db-access.md` |
| `/demo-up` | Spin up an isolated demo stack (Clerkenstein-wired, offset ports, full UI tier + auto-set-dressed — the M20 demo set-dress flow, mirroring `/dev-up`) | `corpus/ops/rosetta_demo.md` + `corpus/ops/demo/README.md` |
| `/demo-down` | Tear down a demo stack cleanly | `corpus/ops/rosetta_demo.md` |
| `/stack-list` | List the live stacks — every `dev-N` and `demo-N` — from the unified registry | `corpus/ops/rosetta_demo.md` |
| `/stack-secrets` | Provision a stack's `.env` secrets (`dev-N` or `demo-N`) from one source + verify coverage — **values-blind** (no verb reads/echoes a secret value) | `corpus/ops/secrets-spec.md` |
| `/stack-seed` | Seed a stack (`dev-N` or `demo-N`) with realistic structural data (presets or `stack.seed.yaml`) | `corpus/ops/seeding-spec.md` |
| `/stack-snapshot` | Set-dress a stack (`dev-N` or `demo-N`) — replay the real public taxonomy + Directus content into it (or capture/status) | `corpus/ops/snapshot-spec.md` |
| `/stack-update` | Sync a stack's code, deps, and schemas (the dev side — demo = teardown + bring-up at a tag) | `corpus/ops/update_guide.md` |
| `/align-dna` | Build/update an Alignment DNA for a mirror engine + capture goldens | `corpus/architecture/alignment_testing.md` |
| `/align-run` | Measure a mirror's alignment score vs a source engine | `corpus/architecture/alignment_testing.md` |

> **The skill set converged in v1.3 "stack party" (M14, hard-rename, no aliases):** the dev lifecycle
> (`/dev-up`, `/dev-down`) mirrors the demo lifecycle (`/demo-up`, `/demo-down`); one generic stack-ops
> set (`/stack-list`, `/stack-seed`, `/stack-snapshot`, `/stack-update`) works on **any** `dev-N | demo-N`.
> `/dev-up` consolidates the former `setup-platform` + `start-platform`; `/stack-update` ← `update-platform`;
> `/stack-list` ← `demo-status`; `/stack-seed` ← `demo-seed`; `/stack-snapshot` ← `demo-snapshot`.

### Using the Dev-Up Skill

For building, starting, or set-dressing the Anthropos development environment:
```bash
/dev-up           # the main dev stack (N=0): first-time build (or resume) + start
/dev-up 2         # an additional isolated dev-2 stack, set-dressed by default
```

`/dev-up` consolidates the former `setup-platform` + `start-platform`. It executes
`corpus/ops/setup_guide.md` (first-time build) + `corpus/ops/run_guide.md` (start + health) with:
- Verification before/after each step + user confirmation before destructive operations
- Progress tracking via TodoWrite
- For an additional `dev-N`: the M13 set-dress pass (cache-first snapshot replay + a light `dev-min` seed
  + the per-stack-Directus firewall check), default-on + non-fatal. The per-stack Directus itself is
  **opt-in for dev** via `--local-content` (v1.5 M22/M23): with it the recipe is EXECUTED (a per-stack
  Directus boots on an offset port + `backend`'s cms domain is cut over → content self-contained); without it the stack
  reads content live from prod (the documented fallback)
- Auto-improvement of documentation when issues are found (ops-reports → `/update-knowledge`)

Tear an additional dev stack down with `/dev-down N` (mirrors `/demo-down`).

### Using the GitHub Setup Skill

For configuring GitHub SSH access to contribute to `anthropos-work` repositories:
```bash
/setup-github
```

This skill executes `corpus/ops/setup_github_guide.md` with:
- Support for single account or dual account (personal + work) setups
- SSH key generation and configuration
- Ensuring work account is the default (critical for Docker compatibility)
- Key persistence across terminal/computer restarts
- Progress tracking via TodoWrite

### Using the Stack-Update Skill

For syncing a stack's code, dependencies, and database schemas:
```bash
/stack-update           # the main dev stack
/stack-update dev-2     # a named additional dev stack
```

This skill (← the former `update-platform`) executes `corpus/ops/update_guide.md` with:
- Daily/weekly/full update scenarios
- Git conflict handling
- Migration application
- Docker image rebuilding

(Demo stacks aren't updated in place — they're disposable; re-create with `/demo-down` + `/demo-up` at the
desired refs.)

### Using the Document Skill

For documenting new platform evidence across the Rosetta corpus:
```bash
/update-knowledge [evidence description]
```

This skill analyzes new evidence and performs a **corpus-wide sweep** to update all relevant documentation:
- Inspects the evidence (repos, features, tools, feedback)
- Checks ALL corpus sections that may need updates
- Updates Claude skills when automation is affected
- Ensures new content is discoverable from parent docs

Example invocations:
- `/update-knowledge the new studio-analytics repo`
- `/update-knowledge issues found in setup_progress.md`
- `/update-knowledge the Redis caching layer isn't documented`

### Working in stack workspaces

Hands-on work with the Anthropos platform happens in a **stack workspace** — a
git-ignored `stack-*/` directory that "spans" one full local stack. Each holds its
cloned platform service repos **plus its own clone of the `rosetta-extensions`
tooling monorepo**:

| Workspace | Stack |
|-----------|-------|
| `stack-dev/` | the local **dev** stack (platform repos + its dev tooling clone) |
| `stack-demo/` | disposable **demo** stacks (Clerkenstein-wired, offset ports) — a **true peer of `stack-dev`** with its **own** platform clone set (v1.8 "understudy" M26) |
| `stack-dev-2/` | a secondary dev stack |
| `stack-stage/`, `stack-tests/`, … | future stacks, same pattern |

All hands-on platform work happens inside a `stack-*/` dir; the documentation
corpus stays clean. (Setup/run/update progress is tracked in
`stack-dev/setup_progress.md`.)

> **`stack-demo` is self-contained (v1.8 "understudy" M26).** A demo now builds **entirely from its own
> `stack-demo/` clone set** — `/demo-up`'s `ensure-clones.sh` bootstrap-clones `stack-demo/platform` from
> GitHub + `make init`s the peer repos, so a box with **only** `stack-demo/` (no `stack-dev/`) can bring a
> demo up end-to-end. The sole sanctioned `stack-dev` read is `ensure-clones.sh` seeding the shared
> `platform/.env` copy-if-present (same Clerk app + GH_PAT; non-fatal if absent — `/stack-secrets` provisions
> the real one). It never borrows `stack-dev`'s repos or built images for the build SOURCE.

### `rosetta-extensions` — where stack tooling lives

`rosetta-extensions` (private: `anthropos-work/rosetta-extensions`) is the
executable-tooling monorepo that **operates** stacks — **eleven** sections: `clerkenstein`
(the Clerk mock), `demo-stack`, `dev-stack`, `stack-injection`, `stack-core`,
`stack-seeding`, `stack-secrets`, `stack-snapshot`, `stack-verify`, `alignment`, `playthroughs`
(plus a non-code `knowledge/`). ⚠️ **This list omitted `stack-secrets` and `playthroughs` until M257x
iter-129** — both are full sections with their own `go.mod`, `stack-secrets` is the one the
`/stack-secrets` row below *already depends on*, and `playthroughs` is described at length elsewhere in
this same file. **This is the enumeration every session loads**, so an agent looking for either tool
concluded it did not exist. `rosetta` documents *how the platform works*;
`rosetta-extensions` is *the tooling that spins up, injects, and seeds copies of
it*. It has **two clone roles**:

- **Authoring copy → `.agentspace/rosetta-extensions/`** — the single working clone
  you spawn on demand to **read / build / test** the tooling, then commit, **tag,
  and `git push --tags` to origin**. New tools are developed here.
- **Per-stack consumption copies → `stack-<role>/rosetta-extensions @ <tag>`** —
  each stack consumes the tooling at a **pinned tag** (reproducible). The
  `/dev-*`, `/demo-*`, `/stack-*`, and `/align-*` skills drive a stack's own clone.

**Policy:** all code/scripts that operate the corpus/platform on a spawned stack
live in `rosetta-extensions` — never scattered in the `rosetta` corpus, never
authored ad-hoc inside a stack dir. A new need/tool is built and tested in the
`.agentspace/rosetta-extensions/` authoring copy, tagged, **the tag pushed to
origin**, then consumed per-stack via its tagged clone.

> **⚠️ Tagging is not publishing.** A stack — especially a remote one — clones
> `rosetta-extensions` **from origin** at a pinned tag (the M217 FATAL pin guard).
> A tag that exists only in your local authoring copy is **unreachable** to it, and
> the failure looks like a missing feature rather than a missing tag. **`git push
> --tags` is part of shipping a tool, not an afterthought.** M236 lost its entire
> first iteration to this: `billion` was pinned to the previous release's tag with
> **0 of 13** `playbill-*` tags on origin, so the feature under test could not be
> obtained at all. Before any prove-it-live milestone, verify the tag is *on origin*
> (`git ls-remote --tags origin`) — see
> [`corpus/ops/verification.md`](corpus/ops/verification.md) pre-flight rung zero.

See [`corpus/ops/rosetta_demo.md`](corpus/ops/rosetta_demo.md)
and [`corpus/services/clerkenstein.md`](corpus/services/clerkenstein.md).

## Architecture Overview

### Three-Tier Service Model

**Core Backend Services (Tier 1)**: Go microservices

> **⚠️ `app` is the backend monolith.** Eight services below left compose — **SEVEN of them were folded
> into `app`** and run in-process as the single `backend` service: **skiller** (July 2026), **skillpath**
> ("skillpath-in-app", M502→M507), **jobsimulation** ("jobsim-in-app"),
> **cms** ("cms-in-app v8.0", app **v1.360.0** — the step that took the supergraph **3→1**), and —
> since the v9.0 "support-in-app" program, whose containers platform `838d907` (merged `0c91421`,
> 2026-08-05) deleted outright — **storage**, **messenger** and **customerio-sync**.
> **⚠️ `roadrunner` is the eighth and it was NOT folded — this banner listed it among the seven until
> M257x iter-137.** `app/internal/roadrunner/` **exists at no ref and was never added** (`git log --all
> --diff-filter=A -- internal/roadrunner` → **0 commits, ever**, in a full 6,728-ref clone at `app`
> `ad9f3c498`; positive control `jobsimwiring` → 3 paths). The seven that WERE folded are
> `app/internal/{cms,customeriosync,jobsimulation,messenger,skiller,skillpath,storage}/`, and the fenced
> map has said **seven** since iter-102. Roadrunner was **deleted and its job replaced**, in-process,
> **inside the jobsimulation domain** — `app/internal/jobsimwiring/wiring.go:123` wires
> `jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))` against
> `internal/jobsimulation/runner`, and the source's own comment there says it *"replaces the removed
> roadrunner RPC edge"*. **Deleted-and-replaced is not merged-and-undeployed**, and only one of the two
> predicts an `app/internal/` package.
> There is no cms / jobsimulation / skiller / skillpath / roadrunner / storage / messenger /
> customerio-sync container, profile, port
> or subgraph. Every application table lives in the **`public`** schema. Standalone-deployment
> teardown is **M810** — and it has now **LANDED for both**: for jobsimulation at `6092c6d2` (which
> deleted the ECS service *and* the ECR repository), and **for cms too — corrected M257x iter-124**.
> **This banner said M810 *"has not moved for cms"* on the strength of `cms/terraform/main.tf:39`
> `service_desired_count = 0`, and that reading is retracted**: iter-123 cloned `infrastructure`
> (`13c248e6`) and found **no `module "cms"` anywhere in it**, with
> `terraform/production/services.tf:64-70` recording the destruction in the platform's own words. **A
> service repo's own `service_desired_count` is not evidence of production state** — it is an input to
> a module that must be *instantiated* by that `services.tf`, which declares exactly ten; four repos
> this corpus quotes declare a count that instantiates nothing (`corpus/architecture/org-repos.md` § 3).
> What is still pending for cms is the legacy **schema** drop, a separate M810 step. The fenced
> statement is
> [`corpus/architecture/platform-migration-status.md`](corpus/architecture/platform-migration-status.md).
> skillpath's teardown is **M507**.

In the default local profile (`core` — renamed from `graphql` at platform `0dab54d`):
- Backend (`app`): Main API gateway and user management; also hosts the **AI-readiness** workforce subsystem (org-level AI-capability diagnostics — see `corpus/services/ai-readiness.md`) **and the skills domain** — taxonomy (**≥42,790 skills / ≥22,470 job roles**, the measured *public* subset; the long-quoted "60K skills / 18K roles" is [not a measurement — 18K is **refuted**, 60K **unverified**](corpus/architecture/shared_libraries.md#taxonomy-figures)), assessment, AI skill matching, and vector embeddings (RAG), absorbed from the former standalone Skiller service (its Ent models now live in `app`, data in the `public` schema; the old `skiller` DB schema is legacy). The skiller RPC surface (GetSkills, GetSkill, SearchSkill, MatchSkill, GetJobRole) is served by `app` — but **`SKILLER_RPC_ADDR` is now set nowhere in compose and read by nothing**: its last consumer was the `messenger` container, deleted at `838d907`, and app's own reader collapses that client onto the in-app skiller RPC server in-process (`app/internal/messenger/adapters/skiller.go:11`). The `skiller` git repo still exists but is decommissioned. **Also hosts the skill-path progression engine** (per-user `SkillPathSession → ChapterSession → StepSession` state) — absorbed from the former standalone **Skillpath** service ("skillpath-in-app", platform M502→M507); session state now lives in `public.skill_path_sessions` (the old `skillpath` DB schema is a legacy husk), and the skill-path session GraphQL types are served by `app`'s `backend` subgraph. **There is NO `SkillPathSessionService` RPC** — M506 *removed* it rather than re-hosting it (measured: 0 occurrences in Go source); the engine is reached in-process and over GraphQL. And the **newer app-owned domains**: course-builder (`corpus/services/coursebuilder.md`), AI Labs + credits (`corpus/services/ai-labs.md`), ask-engine / Talk-to-Data (`corpus/services/askengine.md`), and the server-owned academy store (`corpus/services/academy-backend.md`)
  **Plus the cms and jobsimulation domains** (see the merge banner above — **there is no roadrunner
  domain**; Judge0 is reached from inside the jobsimulation domain):
  - **cms domain** (`app/internal/cms/`): **the content layer** — owns the authored CONTENT / DEFINITIONS (skill paths, simulation blueprints, the content library), wrapping Directus as a proxy + business-logic + cache layer; **and the embedded studio-room AI generation pipeline** (the `anthropos-studio-room` repo is pulled into the `app` image by CI). Directus itself stays external at `content.anthropos.work`. **NB: the cms domain — not the skill-path or jobsimulation engines — owns skill-path and simulation content** (content-vs-runtime-state split below)
  - **jobsimulation domain** (`app/internal/jobsimulation/`, wired by `internal/jobsimwiring/`): **runtime/session engine** that *runs* AI simulations (voice, chat, code, documents) and emits completion events; the simulation *definition/blueprint* it runs is CONTENT read from the cms domain by ID — **in-process** now, no `cms.GetSimulation` RPC hop. It holds run/session state — not content
  - **Judge0 code execution** — reached directly via `JUDGE0_BASE_URL`, wired **inside the jobsimulation
    domain** (`internal/jobsimwiring/wiring.go:123` → `internal/jobsimulation/runner`). **This line read
    *"roadrunner domain"* until M257x iter-137; there is no such package** (see the banner above)
- Sentinel: Authorization only (Casbin RBAC/ABAC) — authentication is Clerk + the `authn` middleware in each service, not Sentinel
- Gotenberg: Office-doc → PDF conversion (third-party image; consumed by `app/internal/converter/gotenberg.go`)

> **⚠️ Storage, Messenger and CustomerIO-Sync are ALL FOLDED INTO `app`.** The v9.0 program landed
> 2026-08-04 (platform `0dab54d` / `app` `9d00a313` v1.367.0); one day later `838d907` (PR #26,
> *"drop the support-service containers"*) **deleted all three compose services outright** and took
> `storage` + `messenger` out of `repos.yml`, so `make init` no longer clones them. `core` starts
> **five** containers — `backend`, `gotenberg` and the always-on floor (`postgresql`, `redis`,
> `sentinel`). There is **no local `storage` / `messenger` / `customerio-sync` container at all** — not
> even a rollback path, which is how `storage-legacy` and `messenger` read for the two commits between
> `0dab54d` (which introduced both tokens) and `838d907` (which removed both); the old *"dangerous to
> run alongside `backend`"* warning (two writers on one bucket, two consumers on one Redis group) is moot
> because there is no second writer left to start. The two that reach outside the process are gated
> in-app by `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, **unset = OFF on a developer machine**
> (`app/main.go:285`, `:286`; compose states the rationale where the variables would have gone,
> `docker-compose.yml:84-92`). In production, messenger's ECS service is scaled to zero as the
> rollback path (`messenger/terraform/main.tf:29`, `service_desired_count = 0`) — **but that line
> instantiates nothing: `module.messenger_euwest1` is DELETED from `infrastructure` @ `13c248e6`
> (`services.tf:622`), so it is orphaned dead code and the "rollback path" comment is an intent
> production no longer holds (M257x iter-123; the rule and its three siblings are in
> `corpus/architecture/org-repos.md` § 3) — while **storage's ECS
> service block is DELETED, not scaled** — its module survives only to keep the buckets, the
> CloudFront distribution and their `prevent_destroy` guards in configuration
> (`storage/terraform/main.tf`, 18 lines). `app` serves object storage in-process
> (`app/main.go:524`, `:525`) and has **taken over messenger's own Redis consumer group**
> (`:1384`, `:1450` — every anchor here is at `app` **`2035f9a`**, which was `origin/main` on 2026-08-06;
> origin/main was `ad9f3c49` when this was written — it is **`3eaadae6`** as of 2026-08-09, and `ad9f3c49` is **28 commits behind** it (corrected M257x iter-228; the observability and shared-libraries docs both already recorded the newer ref, so this file disagreed with the corpus, not just with the platform), and all five anchors resolve **identically** there — the 5 intervening
> commits touched no Go source. Cite the sha, never the moving label). All of it is cited on both sides in the fenced map —
> `corpus/architecture/platform-migration-status.md`, the `storage`, `messenger` **and
> `customerio-sync`** rows (the last one net-new at iter-87: it was never in `repos.yml`, so no
> membership assertion could have caught its state change).

(No Tier-1 service sits in a non-default profile any more — the last three that did were deleted at
`838d907`; see the archived list below.)

Production-only / deployed-only (not in local docker-compose):
- db-backup: a **43-line Bash script** (not Go) that `pg_dump`s Postgres to **S3 + a Hetzner Storage Box** — **two** destinations, never Azure. **Its schedule and trigger have been commented out since `7dd1b80` (2025-05-29) and production pins that commit**, so the task definition is deployed and nothing fires it. The long-quoted *"every 6 h to S3, Azure, Hetzner"* was wrong on five counts and **"6 h" never had a source at all** (the disabled value was `rate(12 hours)`). RDS multi-AZ + an hourly AWS Backup plan with PITR still cover durability; what is lost is the **offsite, non-AWS** leg. See `corpus/services/db-backup.md`

Archived / merged (removed from local orchestration; repo dirs may still exist on disk):
- Chronos (was: scheduling & time-based events) — removed via platform commit `045857c`
- Intelligence (was: background data sync between backend and skiller schemas) — removed via platform commit `fdfa189`
- Skiller (was: skills taxonomy, assessment, embeddings) — **merged into `app`** (July 2026, v2.1 "quick change"); domain now in the `public` schema, `skiller` repo decommissioned, no skiller container/subgraph. See `corpus/services/skiller.md` + the `backend.md` fact-sheet
- Skillpath (was: per-user skill-path progression runtime engine) — **merged into `app` then decommissioned** ("skillpath-in-app", platform M502→M507); the engine now runs in `app`, session state moved to `public.skill_path_sessions` (old `skillpath` schema is a legacy husk), no skillpath container/subgraph. The skill-path *content* lives in the cms domain. See `corpus/services/skillpath.md` (redirect) + the `backend.md` fact-sheet
- Roadrunner (was: Judge0 code-execution proxy) — **DELETED, not merged** (corrected M257x iter-137; this row said *"merged into `app`"*). Its compose service and `repos.yml` entry went at `d11a403`; **no `app/internal/roadrunner/` was ever created**, and `backend` calls Judge0 directly via `JUDGE0_BASE_URL` from inside the jobsimulation domain. **In production there is no roadrunner ECS service either** — `infrastructure` @ `13c248e6` declares **ten** service modules in `terraform/production/services.tf` and roadrunner is not among them; `roadrunner/terraform/main.tf:19` `service_desired_count = 1` is an input to an **uninstantiated** module (the `cms`/`messenger`/`wundergraph` class — `corpus/architecture/org-repos.md` § 3). See `corpus/services/roadrunner.md`
- Jobsimulation (was: the AI-simulation runtime/session engine) — **merged into `app`** ("jobsim-in-app"); the engine is `app/internal/jobsimulation/`, its 23 run-state tables moved to `public`, no container/subgraph. **M810 has LANDED for the ECS service**: `6092c6d2` deleted the `module "jobsimulation"` block outright, so there is no ECS service, task definition or ECR repository any more (`jobsimulation/terraform/main.tf:15-22`). The module *file* survives, and deliberately, because it still **owns the LiveKit/Chime recording S3 buckets** that `backend` reuses by literal name, the `/production/jobsimulation/*` SSM parameters, and the atlas tracker for the legacy `jobsimulation` schema — **dropping that schema is a separate, still-pending M810 step** (`:24-40`). **Do not generalise this to cms** (next row). See `corpus/services/jobsimulation.md`
- CMS (was: the content layer + Studio) — **merged into `app`** ("cms-in-app v8.0", app **v1.360.0**); the domain is `app/internal/cms/`, its similarity/Studio tables moved to `public`, and the supergraph went **3 → 1 subgraph** (`915da06` deletes *both* the cms and jobsimulation SDLs from a config that listed three — **do not take the count from that commit's own subject line**, which says "2→1" and is where the wrong figure entered the corpus; see `corpus/services/graphql-wundergraph.md`). Directus stays external. **M810 for cms is RESOLVED — the ECS service is DESTROYED** (corrected M257x iter-124; measured at iter-123). **This entry said the prod state was *"NOT MEASURABLE from our clone set"* because `infrastructure` had never been in one — that was a clone-set limit, not a measurement limit, and cloning the repo settled it.** At `infrastructure` `13c248e6` there is **no `module "cms"` declaration at all**, and `terraform/production/services.tf:64-70` states what the apply destroyed — the ECS service, task definition, ECR repository, IAM roles, security group, Cloud Map entry, log group, alarms and the ten `/production/cms/*` SSM parameters. So the two facts that looked contradictory never were: `6efa1d5` (merged `f38c0c4`, 2026-08-04), which **deleted** the build-production workflow saying *"the cms ECR repository is decommissioned (M810)"*, was the correct signal, and **`cms/terraform/main.tf:39` `service_desired_count = 0` is ORPHANED DEAD CODE** — no root module instantiates that file. The legacy **schema** drop is a separate, still-pending M810 step. See `corpus/services/cms.md`
- Storage (was: file/blob storage management) — **merged into `app`** (v9.0 "support-in-app"); no compose service since platform `838d907`, and out of `repos.yml`. See the v9.0 banner above + `corpus/services/storage.md`
- Messenger (was: email notifications via Brevo/Sendinblue) — **merged into `app`** (v9.0); no compose service since `838d907`, and out of `repos.yml`. See the banner above + `corpus/services/messenger.md`
- CustomerIO Sync (was: background data sync to Customer.io — the one service built straight from a GitHub URL, never cloned locally) — **merged into `app`**; no compose service since `838d907`. See the banner above + `corpus/services/customerio-sync.md`

**Shared Libraries** — **do not read this list as `app`'s dependency set; it is not one.** Measured at
`app` `3eaadae6` (v1.371.1), `app/go.mod:14-18` requires **five** org-private modules, **all direct, zero
`// indirect`, and no org `replace`** — `analytics-go` `v0.3.1`, `colony` `v0.35.2`, `proto` `v1.210.0`,
**`storage` `v0.15.2`**, `taxonomy` `v1.2.0`. **`ai` is NOT among them** and **`authn` never was.** None
are cloned by `make init`/`repos.yml` — they are pulled at Docker build via `GH_PAT`/`GOPRIVATE`. Full
picture: `corpus/architecture/shared_libraries.md`.
- colony: Platform framework (logging+Sentry, DB, Redis, GraphQL/RPC servers, middleware, pub/sub via Watermill); **also contains `authn`**
- proto: Protobuf definitions (RPC contracts) + hand-written domain types
- taxonomy: **node-id library** (`NodeID` type + ID generation/validation) — **not** a dataset; the skill/job-role data (**≥42,790 skills / ≥22,470 job roles** — public subset, measured 2026-06-29; ["60K / 18K" is not a measurement](corpus/architecture/shared_libraries.md#taxonomy-figures)) lives in `app` (backend — the `public` schema, since the skiller→app merge)
- **storage** (`go.mod:17`): a **TYPE SHIM, not an RPC edge.** `app` imports `sdk/storage` (the `Client`/`PublicClient` structs) and `sdk/storage/v1` (the three-method `Service` interface) at 36 import lines across 32 files, then **implements that interface itself** — `internal/storage/service.go:48-56` fills `sdkstorage.Client{V1: NewService(...)}` with its own in-process manager. No SDK RPC client is ever constructed and `STORAGE_RPC_ADDR` occurs in **zero** Go source
- **analytics-go** (`go.mod:14`): a two-file Brevo product-event tracker (`Init`/`Track` fan-out). Wrapped by `app/internal/tracking`, and it carries **Stripe subscription-lifecycle events → Brevo** at `app/internal/payments/handler.go:302-316` (seven event names switched off `entSub.Status`), wired at **`main.go:494-495`** @ `app` `ad9f3c498` (the file's only two `trackingManager` lines; **this pinned the storage-in-app comment block instead, until M257x iter-138**, and `handler.go:302-316` was verified exact). The repo has been untouched since **2025-02-12** and `v0.3.1` **is its newest tag** — dormant, and load-bearing
- ~~ai~~: **folded INTO `app`** at `1e457fa70` (2026-08-04) from tag `v1.40.2`, now `app/internal/ai/` with **84 importing files**, and kept out by `.github/workflows/ai-module-guard.yml` — a **PR-time CI job** running `TestNoExternalAIModuleImports`, *not* a build-time failure, and it **passes** when its merge conflicts (`:96-100`). Cost tracking is `app/internal/aiusage`; **vendor selection** lives in each consumer's wrapper — a caller-supplied switch, not an EU-first fallback ladder
- ~~authn~~: Clerk JWT authentication — shipped **inside colony** as `colony/authn`; in no repo's `go.mod`

> **⚠️ "Merged into `app`" describes the RUNTIME, not the module graph — and the two came apart.**
> `ai` left `go.mod`; **`storage` did not.** A cleanup driven by this file as it read before 2026-08-07
> would have deleted **two** repos `app` cannot build without: `storage` (a direct require, *and* still
> maintained — HEAD 2026-08-05, tags out to `v0.15.8`, six past app's pin) and `analytics-go`, which
> nothing guards at all. **`ai` must not be deleted either**, for a reason that lives in this repo:
> `rosetta-extensions/stack-seeding` pins `ai v1.40.1`, so a doc-driven "app doesn't import it any more"
> deletion breaks **Rosetta's own tooling**. (`app/internal/ai/module_import_guard_test.go:15-17` and
> `app/CLAUDE.md:289-294` both say so.) The org-module block is **actively shrinking** — at `b948604f`
> it was **seven** (`ai` at `:14`, `messenger` at `:17`) — so treat any list here as a dated measurement,
> never as a standing fact.
>
> Since the merges these libraries are imported by **app**, **sentinel**, **storage** and **messenger**
> only — but the last two are frozen legacy repos since `838d907` (no compose service, not in
> `repos.yml`), so only **app** and **sentinel** are still built. **`storage` is a supplier as well as a
> consumer here, and the old wording hid that.** **`customerio-sync` is not in that
> measurement either way**: its repo has never been in the clone set, so nothing here has read its
> `go.mod` — the same blind spot the `customerio-sync` row of
> `corpus/architecture/platform-migration-status.md` records.

**Studio Services & Standalone Apps (Tier 2)**: Content creation tools + standalone apps. **"Internal-only"
is right for Studio-Desk/Studio-Room and WRONG for Ant Academy** — see its row below
- Studio-Desk (TypeScript/Vite/Express): Design tool for creating simulation blueprints (repo: `studio-desk`)
- Studio-Room (Python/Asyncio): AI-powered content generation pipeline (repo: `anthropos-studio-room`). **Embedded inside the `app` (backend) container** since cms-in-app — pulled into the image by CI (`additional_repo`, app v1.360.1); never a standalone deployment.
- Ant Academy (Next.js 16 + Expo): the AI-academy learning product — **a public storefront with an enterprise/org tier, NOT an internal `@anthropos.work`-only portal** (that description was refuted at M257x iter-115 and swept corpus-wide at run 81: it sells a $399/yr subscription to anonymous visitors and its code carries no domain predicate) (repo: `ant-academy`). **Vercel-deployed standalone — not in docker-compose.** **NOT in `repos.yml` (by design — v1.10b M49 #5)** — so `make init` does **not** clone it. For a **demo**, `ensure-clones.sh` clones it **explicitly** (phase d2, non-fatal — `repos.yml` lives in the ephemeral platform clone, so editing it is non-durable + a platform edit); for **dev**, it's a manual `git clone`. Runs natively via `cd ant-academy/code && npm run dev` (port 3077). Auth via Clerk; **since v0.5.1 the course catalog is DB-authoritative** — read from the platform academy subgraph over GraphQL (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`), degrading to an **empty grid** when the endpoint is unset or the academy DB is empty (the demo "empty academy" root cause — the v2.5 M229/M230 thread). See `corpus/services/ant-academy.md`.

**External Services (Tier 3)**: Third-party integrations
- Clerk: User authentication (SaaS)
- Directus: Headless CMS (self-hosted)
- ~~GraphQL/Cosmo Router~~ — **DELETED from the platform** at `2adcf71` (2026-07-31, PR #23 *"drop the WunderGraph router; point local dev at backend"*). There is no `graphql` compose service, no `graphql-wundergraph` entry in `repos.yml`, and no federation gateway. **GraphQL is served directly by `backend`** at `:8082/graphql/query` — note the path moved too (`/graphql` → `/graphql/query`), so a host-only re-point 404s rather than errors. `.env_example` records that the `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` name is now historical. (Measured in M257x iter-12; the prior claim here of *"3 subgraphs"* was itself already stale — the cms-in-app merge had taken the supergraph to one.) `categoryTree`/`fullCategoryTree` were dropped, not ported.
- AI Providers: OpenAI, Anthropic, Mistral — EU-resident clients by default; **there is no ordered EU-first fallback ladder** (`corpus/architecture/external_services.md:579`)
- LiveKit: Real-time voice engine for simulations
- AWS Chime: Video/audio recording

**Frontend Applications**: Next.js 15 monorepo on Vercel (`next-web-app`; see `corpus/services/next-web-app.md`)
- Next Web App: Main user-facing application
- Hiring App: Recruiting and hiring workflows
- Mobile App: Expo/React Native mobile experience

### Communication Patterns

- **Core Services ↔ Core Services**: Connect-RPC + Redis Streams (via Watermill) for async messaging. **The only cross-process *Connect-RPC* edge left in a local stack is `backend → sentinel`** — `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`), and **zero `*_RPC_ADDR` variables anywhere**. **⚠️ It is NOT the only cross-process edge, and compose does NOT set exactly one service address** — both of those stronger forms stood here until M257x iter-102 and both are false. `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the **default `core` profile**, `:183` `profiles: [core, backend, all]`; consumed at `app/internal/converter/gotenberg.go:31`, a plain `http.NewRequestWithContext(…"POST"…)`, not RPC), and reaches Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The qualified wording at `corpus/architecture/architecture_overview.md:343` is the model. The `messenger → backend` edge was not re-pointed, it is *gone*: the `messenger` block was the only thing that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` and `SKILLER_RPC_ADDR` (`d11a403` had re-pointed the middle two at `backend` — M809), and `838d907` deleted the service and all four with it. **backend → storage is no longer mid-fold either** — `STORAGE_RPC_ADDR` is set nowhere *and* read nowhere; `app` serves object storage in-process and says so at `app/main.go:504` **@ `2035f9a4`** (**state that ref** — this sentence also cites `b948604f` below, and at `b948604f` line 504 is an unrelated jobsim-in-app comment; `2035f9a4` was origin/main on 2026-08-06, origin/main was `ad9f3c49` when this was written — it is **`3eaadae6`** as of 2026-08-09, and `ad9f3c49` is **28 commits behind** it (corrected M257x iter-228; the observability and shared-libraries docs both already recorded the newer ref, so this file disagreed with the corpus, not just with the platform), and `:504` is identical there). **`app` is both producer and consumer of FIVE streams**, not three — `backend`, `skillpath`, `jobsimulation`, `cms` and the `ai_usage` usage stream. *(This line named only the middle three until M257x iter-102; it omitted `backend` itself while asserting exhaustiveness.)* Watch the partition when quoting a number: **four** is the application-stream subtotal, **five** the both-ways total, **six** the subscriber count. **`skiller` is the exception and the sixth subscriber:** `app` subscribes to it (`app/main.go:1276` @ `b948604f`) but nothing publishes to it — the producer was the standalone skiller service and was **deleted, not re-hosted**. Full enumeration with both refs: `corpus/services/backend.md`
- **Frontend/Studio → Backend**: GraphQL **straight to `backend`** at `:8082/graphql/query` — the Cosmo router was deleted at platform `2adcf71`, so there is no gateway hop and no federation
- **External Integrations**: Clerk SDK + JWT middleware (authn library), Directus proxied via the cms domain inside `backend`
- **AI**: vendor selection implemented in each consumer's `internal/ai` wrapper, **not** the shared `ai` library. **Not a fallback ladder** (`corpus/architecture/external_services.md:579`): an EU Azure client by DEFAULT, a US Azure client swapped in by the PostHog flag `flag_use_azure_us`, and direct-OpenAI as the RETRY target on HTTP 429 — three independent levers, not three ordered rungs; Anthropic always Bedrock `eu-west-1`, except Course Builder's `ANTHROPIC_API_KEY` path to `api.anthropic.com`. Cost tracking in `app/internal/aiusage`
- **Multi-tenancy**: Shared DB, shared schema, `organization_id` on **most** application tables; 3-layer isolation (DB, Sentinel auth, Clerk identity). **⚠️ "on every table, so no cross-tenant access is possible" is RETRACTED** — the DB layer auto-filters only where an Ent schema declares the privacy policy; the rest are filtered by application code or not at all. The measured split and the derivation command live in `corpus/architecture/security_compliance.md`

### Environment Configuration

**Platform services** share a **single centralized `.env` file** in the `platform` repository. Docker-based services do not need their own `.env` files.

**Studio-Desk** requires its own `.env` file (`studio-desk/.env`) with Clerk and OpenAI credentials copied from `platform/.env`.

**Ant Academy** requires its own `.env` file at `ant-academy/code/.env.local` (not the repo root — the React app reads only from `code/.env.local`). Reuse `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `CLERK_SECRET_KEY` from `platform/.env`, and add `OPENAI_API_KEY` / `ANTHROPIC_API_KEY` for the `/api/ai/chat` route. Set `REQUIRE_ORGANIZATION_MEMBERSHIP=0` for solo local dev to skip the org-membership gate.

Critical environment variables — **each one verified declared in `platform/.env_example` @ `0c91421`**
(M257x iter-237; this list is a hand-maintained tuple and nothing fences it, so treat it as a dated
measurement):
- `GH_PAT` (GitHub Personal Access Token — required for Docker builds to pull private Go modules)
- `CLERK_SECRET_KEY` (Auth — backend services)
- `OPENAI_KEY` (AI services — the name really is `OPENAI_KEY`, **not** `OPENAI_API_KEY`; both exist and mean different things)
- `VITE_CLERK_PUBLISHABLE_KEY` (Studio-Desk via Docker)
- **`DIRECTUS_TOKEN`** (Content — ships **BLANK** at `.env_example:92` and is the one you must fill; a
  missing value is the classic *stack boots, catalog empty* failure). `DIRECTUS_BASE_ADDR` (`:91`) is
  declared with a working default.

> **⚠️ `DIRECTUS_PUBLIC_BASE_ADDR` was listed here and is NOT a variable you set** (corrected M257x
> iter-237). Compose **hardcodes** it —
> `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work` at `docker-compose.yml:53` — and
> `.env_example` does not declare it at all, so putting it in your `.env` does nothing. The corpus already
> said so precisely at `corpus/architecture/external_services.md:133`; only this list disagreed. It named
> the one Directus variable you **cannot** set while omitting the one you **must**.

### Makefile-Driven Workflow

The `platform` repository provides a Makefile as the single entry point for all development operations. The repos `repos.yml` still lists are cloned as siblings via `make init`, and Docker builds from local code.

```bash
# First-time setup
cd stack-dev/platform
make init              # Clone the 4 repos in repos.yml: app, sentinel, next-web-app, studio-desk
make up                # Build from local code and start (core profile — the default)
make migrate           # Apply all database migrations

# Daily development
make pull              # Pull main on all repos (auto-stash dirty changes)
make status            # Git status across all repos
make up                # Rebuild and start (auto-builds from local code)
make down              # Stop all services
make ps                # Show running containers
make logs S=backend    # Tail logs for a service
make dev S=backend     # Stop container, develop natively
make reset-db          # Wipe DB, restart, re-migrate (WARNING: data loss)
```

Docker Compose profiles control which services start:

> **⚠️ `graphql` is no longer a profile, and asking for it does NOT fail.** Platform `0dab54d`
> (*"rename graphql -> core"*) renamed the profile and dropped the standalone storage; the `Makefile`
> now reads `PROFILE ?= core`. **`postgresql`, `redis` and `sentinel` declare no `profiles:` key at
> all**, so they are in *every* selection — which means asking for the old `graphql` token **exits 0
> and starts those three**. Postgres answers, `docker ps` is non-empty, the stack looks alive, and
> the application is simply absent. The retired `cms`, `jobsimulation`, `roadrunner` and `storage`
> tokens behave identically — as do `storage-legacy`, `messenger` and `customerio-sync`, retired at
> `838d907` (2026-08-05). **Only five profile tokens still exist**: `all`, `backend`, `core`,
> `frontend`, `studio-desk`. **Grade a documented compose command on "does it still select
> anything", never on "does it still parse."** Fenced by
> `rosetta-extensions/stack-core/platform_predicate_guard.py` (G1/G3), which is also why no retired
> token is spelled here in runnable form — a copy-pasteable command for a silent no-op is the defect.

| Profile | Services started (besides the always-on floor `postgresql`, `redis`, `sentinel`) |
|---------|----------|
| `core` *(default — `PROFILE ?= core`)* | backend, gotenberg |
| `backend` | backend, gotenberg |
| `all` | backend, gotenberg, next-web-app, studio-desk |
| `frontend` | next-web-app — **but selecting it alone exits 1**: `next-web-app` declares `depends_on: backend` (`docker-compose.yml:165-167`), which this profile does not select, so compose rejects the project as invalid |
| `studio-desk` | studio-desk — **also exits 1**, same `depends_on: backend` reason (`docker-compose.yml:138-140`) |

Usage: `make up PROFILE=core`

## Key Documentation Locations

### Setup & Onboarding
- `corpus/ops/setup_guide.md`: Complete environment setup instructions
- `corpus/ops/setup_github_guide.md`: GitHub SSH access configuration
- `corpus/ops/staging-bringup.md`: **"New engineer (or AI agent) joining the team — start here"** (its own framing, and `corpus/ops/README.md:11` seconds it). The end-to-end staging bring-up narrative; the longest-form onboarding path in the corpus
- `corpus/ops/quick_ops.md`: The short-order cookbook — common one-off ops recipes without reading a full guide

### Running the Platform
- `corpus/ops/run_guide.md`: Start the platform locally after setup
- `corpus/ops/webhook_setup.md`: Configure Clerk webhooks for user/org sync
- `corpus/ops/directus-local.md`: **The per-stack Directus model** — how a stack gets its own Directus (the `--local-content` cutover, offset port, `cms` re-point) instead of reading content live from prod. The mechanism the demo/dev set-dress flow above depends on; read it before debugging any "content is empty / content is prod's" symptom

### Staging Environments
- `corpus/ops/staging-bringup.md`: Stand a staging environment up end-to-end (also the onboarding path above)
- `corpus/ops/staging-clerk.md`: Clerk wiring for staging (real Clerk, not Clerkenstein)
- `corpus/ops/staging-sync.md`: Keep staging in sync with upstream code + schema
- `corpus/ops/staging_from_dump.md`: Build a staging environment from a database dump

### Demo Environments (disposable, Clerk-free, seeded + set-dressed — v1.1 "show floor" + v1.2 "set dressing")
- `corpus/ops/safety.md`: **The tooling safety contract** — the consolidated read-side (tenant-data firewall + public predicates + read-only capture) + write-side (3-layer isolation guard + never-write-prod + n=0 guards + audit-proven zero pollution) statement. The *why-it-is-safe* anchor for the whole demo/dev family (v1.3 M15). **v2.3 M220 adds Part 3 — the exposure side**, the third axis: who can *reach* a demo and what they get. It is a **disclosure, not a third "never"** — a demo is an **unauthenticated, authz-weakened build** (Clerk verification disarmed in app/cms/jobsimulation — skillpath is decommissioned into app; the authz-skip demo-patch default-ON; the presenter cockpit a **password-free "become any seeded hero"** launcher), and **every demo container is published on `0.0.0.0` — all interfaces — on EVERY `demo-up`, flag or no flag** (a claim `tailscale-serve.md` **denied** until M220; the retraction is in place and fenced). What makes that defensible is Parts 1+2: **there is nothing behind the door** — no customer data can be in a demo, and a demo cannot write prod. Part 3 also carries the **supersession of v2.2's D-DESIGN-1** (*"public reach is never default-on"*) by v2.3's **D-DESIGN-3** — **demo path only; `/dev-up` stays opt-in**
- `corpus/ops/demo/README.md`: **The demo-env family index** — the up→snapshot→seed→use→down flow + recipes + presets
- `corpus/ops/demo/demo-up-defaults.md`: **The `/demo-up` defaults contract** (v2.3 "cue to cue" M220) — every knob and flag that controls a bring-up: **all 31 `DEMO_*`/`STACK_PUBLIC_HOST` env knobs + 10 CLI flags**, each with its real default and the exact `file:line` that reads it. **Derived from the parsers and fenced against them in BOTH directions** (`stack-core/demo_knob_guard.py`): a doc-promised flag with no parser entry is a **false promise**; a parser flag with no doc row is **undiscoverable**. Records the fact nobody had written down — **there are TWO entry points**: `up-injected.sh` (what `/demo-up` runs) takes **only** `<N>` + `--public-host` and **hard-errors `exit 1` on anything else**, while `--profile`/`--services` are flags of the separate `rosetta-demo` wrapper (the skill's `argument-hint` conflated them for releases). And the shape: **every feature knob is an opt-OUT (`DEMO_NO_*`, default `0`)** — so a bare `/demo-up N` **already** seeds the 4-org world (3 workforce + the M223 hiring org), the full UI tier, the cockpit, and set-dress. *"Pull all the data + seed the orgs" was always the default; the usual culprit is a **cold snapshot cache**, not a knob.*
- `corpus/ops/rosetta_demo.md`: The demo-stack lifecycle (bring-up, port-offset, Clerkenstein injection, teardown)
- `corpus/ops/seeding-spec.md`: The `stack.seed.yaml` blueprint + the **production-isolation boundary** (write-side) + the data-DNA (now **100%**, nothing waived)
- `corpus/ops/demo/stories-spec.md`: **The verified-skill chain reference** (v1.9 "storytelling" M34) — how a seeded *verified skill* (a hero's profile + Skill Spotlight chart + the claimed-vs-verified gap) is materialized as the **7-table fan-out** the `PersonaSeeder` writes: the DB-enforced vs inserted-but-invisible constraint landmines, the **G14** session-seeder fix (valid `SIMULATION_TYPE_*`/enum/token + continuous growth-arc score), the `user_level` (claimed side) requirement, the `TaxonomyRefs` resolver (real public node-ids, never fabricated), the `users.go` name/avatar/email patch, and the **seed-side closure gene** (`datadna measure-closure`). The believability spine; vertical slice (Maya). M35 adds the full multi-org Stories & Heroes model, **M36 the org Workforce-Intelligence dashboard surfaces** (the mapped→verified funnel + teams + role gap/mobility + succession + feedback + the org-scale claimed-vs-verified gap), M37–M38 the presenter cockpit, and v1.10 "method acting" layers the per-hero **profile depth**: M39 the profile identity (real name/avatar/org-domain email) and **M41 the `ProfileSeeder`** (work-history + education timeline + a claimed-but-unverified `user_skills` tail that widens the visible claimed-vs-verified gap). **v1.10b "fit-up" M51 adds the AI-readiness showcase org as a 3rd story** (org "Northwind Aviation", 200 members, heroes Aria COMPLETED / Ben STARTED / Dana manager) with **four** net-new AI-readiness seeders (`OrgSettingsSeeder` + `AIReadinessConfigSeeder` + `AIReadinessFunnelSeeder` + — net-new at **v2.3 M219** — the **interview-aggregated-report** seeder, without which the manager's four interview-findings blocks render headings with NO content) seeding **both** a `closed` and an `active` cycle (M219 refuted M51's "live-recompute never completes" premise — it takes **2.09 s**) + the `app-aireadiness-snapshot-loadmembers` read-path demo-patch; the seeder contract is `corpus/services/ai-readiness.md`
- `corpus/ops/demo/cockpit-spec.md`: **The presenter-cockpit UX spec** (v1.10 "method acting" M43) — the slick **light** login launcher a demo-giver drives (`rext demo-stack/cockpit.py`, served at `:7700`+offset): the card-per-hero layout + FontAwesome icons (free CDN), the **one unified [Log in as] CTA** per hero (logs in *and* lands on her per-role `jump_to` — no more separate [Jump]), the seed-manifest download, and the staged login-progress overlay — plus the deep-link contract, the standalone-served-panel (zero-platform-edit) model, and the future-feature surface. Graduates the M37/M38 cockpit mechanics scattered across `stories-spec.md` + `clerkenstein.md`
- `corpus/ops/secrets-spec.md`: **The secret-provisioning spec** (v1.6 "stage door" M27–M30) — provision every repo's target `.env` (`dev-N`/`demo-N`) from one secret source (dir/zip, default `.agentspace/secrets`), **values-blind** (no verb reads/echoes a value), verified by the 6-repo/64-gene **secret-coverage DNA** + the two-tier keep-listed gate. The source-dir/zip layout contract (zEnvs defence), the per-repo target-file map, alias-family vs distinct-similar rules, the waived class, N=0 guard + idempotency, the demo-aware check, and the `DIRECTUS_TOKEN` non-rearm safety (the fix16/17 class). Driven by `/stack-secrets`
- `corpus/ops/db-access.md`: **Production DB read access** (read-side) — the `/db-query` skill + the public-vs-customer boundary + the snapshot read foundation (v1.2)
- `corpus/ops/snapshot-spec.md`: The **`stack-snapshot` extension** (v1.2 M9a/M9b/M10) — capture the public taxonomy + Directus content once from a safe prod source, manifest-cache it in `.agentspace`, replay per-stack (`/stack-snapshot`); the tenant-data firewall + the `stacksnap` CLI + the snapshot-fidelity gate
- `corpus/ops/snapshot-cold-start.md`: **The cold-start capture runbook** (v1.3b M20) — the one case the cache can't shortcut: a fresh box with an empty cache + no safe `--dsn`. The sanctioned DSN-export / restore-a-`pg_dump`-then-`--dsn` path to fill the cache once per release (behind the capture-source policy + `AssertPublicOnly`), **why the wired `postgres` MCP is NOT a capture source** (it returns JSON rows, not COPY bytes), and how it slots into the auto-set-dress bring-up (replay-only, never capture)
- `corpus/ops/idempotency.md`: **The bring-up re-run safety contract** (v1.3b M17) — what happens when you re-run migrate / snapshot-replay / seed: each is safe-and-idempotent or fails loudly with a guard (replay TRUNCATE-then-reload, idempotent seed COPY + casbin `WHERE NOT EXISTS`, the fixed `--reset`, the `set -e` first-run-race hardening). The *run-it-twice* companion to snapshot/seeding-spec
- `corpus/ops/verification.md`: **The bring-up auto-verify safety net** (v1.3b M18) — every bring-up ends with a scoped, **non-fatal** `verify live` on the stack's **own offset ports**: cheap-win `/api/health` + `sentinel.casbin_rules > 0` asserts (the silent-403 catcher) then the full offset/project/scope-aware probe set, so "UP" means *verified-working*. Default-on; a verify bug never blocks a good stack. The *is-it-actually-working* companion to `rosetta_demo.md` + the `/test-platform` skill. **v2.5 M236 adds PRE-FLIGHT RUNG ZERO — *tagging is not publishing*:** a remote stack consumes `rosetta-extensions` only at a tag **fetched from origin** (M217 FATAL pin guard), so tooling that exists only in the local authoring copy is **unreachable** to it. M236's first iter found `billion` pinned to the previous release's tag with **0 of 13** `playbill-*` tags on origin — the feature under test could not be obtained at all. Verify the tag is *on origin* before any prove-it-live milestone
- `corpus/ops/demo/frontend-tier.md`: **The demo UI tier** (v1.3b M19) — `/demo-up` brings up next-web-app + studio-desk (per-demo **cached** Docker image from the **unmodified** Dockerfile, offset ports, minted-pk + offset-URL baked) + ant-academy natively (Clerk-free). The 12 GB Docker-VM prereq + non-fatal pre-flight, the honest "one ~3-min cached build per new `demo-N`" residual, the `--no-ui` (`DEMO_NO_UI`) escape, and the hard **zero-platform-repo-edit** line (repo = build context only). The *see-it-in-a-browser* completion of the demo family
- `corpus/ops/demo/demopatch-spec.md`: **The demo-patch mechanism — the sanctioned zero-platform-edit escape hatch** (v2.3 "cue to cue" M217). When a demo needs a fix with **no env/config/compose seam** (the value is baked into platform source), `demopatch` patches the demo's **own ephemeral clone** just before the image build and reverts it after — the *image* carries the fix, the clone is left git-clean, and the canonical `anthropos-work` repos are **never touched**. The **7 guards** (G1 path-assert · G2 drift-refuse + exactly-once anchor · G3 never-commit · G4 idempotent · G5 self-revert · G6 demo-only · **G7 apply post-condition**), the 10-key manifest schema, the **three apply vehicles** (the `app` patches target the build-scratch clone **outside** the workspace, so `demopatch`'s own G1/G6 correctly REFUSE them — two shell helpers re-implement the ladder against the same canonical manifest), the **chain rule** (`next-web-public-website-url`'s `pre_sha256` **IS** studio's `post_sha256` — it reads "DRIFTED" against a pristine file **by design**), and the **self-healing freshness gate**: *the anchor is the contract; the whole-file sha is only a baseline*. **Read it before adding or re-pinning any patch** — a silently-refused perf patch shipped a 76 s members grid for four releases
- `corpus/ops/demo/build-budget.md`: **The bring-up build budget — what "fast" means for a `/demo-down --purge` + `/demo-up` cycle, and the harness that grades it** (v2.8 "fast build" M255). The sibling of `latency-budget.md`, one layer down, and it exists for the same reason: the corpus asserted *"~3 min per frontend"*, *"~3.7 GB first build"* and *"the ~3.7 GB build cache"* — measured, the cache is **105.4 GB** and the free-disk floor was sized against the wrong number. Defines **READY** (`up-injected.sh` exits 0 **and** `autoverify.json` is green), the **per-phase attribution model** (2 top-level + 12 anchor-derived sub-phases that must sum back, fail-closed), the measured baseline (**n=3 p50 666.29 s** on `billion`, 3/3 green — **billion is DEMO-ONLY, and `odysseus` is RETIRED — `D-v28-15` superseded `D-v28-14` the same day (2026-07-31): dev/test is LOCAL to the new Mac. ⚠️ That supersession reached `knowledge/` 35 times and this corpus ZERO times until M257x iter-226, so `build-budget.md` still argues throughout for a gate host that no longer exists. And no profile has ever been measured for the host that replaced it — `hostprofiles/` holds only `billion.json` and a `laptop.json` describing a different, retired Mac, so gate clause 1 is NOT gradeable today (M257x iter-225)**, of which **UI-tier image builds are 65.5 % and image export/unpack ALONE is 46.2 %**), the **headroom contract** (**four** clauses — a **clause zero** (`require_measured`, which a fresh host with an unpopulated sampler hits FIRST) plus CPU, memory, disk — against *measured, checked-in* host profiles; memory uses the **measured per-lane peak**, never the V8 ceiling, because a ceiling is not a reservation and would fail the single lane the host demonstrably runs → **neither host fits two concurrent Next.js build lanes**), the **campaign protocol** (the binding constraint is the ~18 GiB mid-cycle TRANSIENT, not the ~2 GiB a steady rep nets; reclaim with `--filter until=24h`, never `-af`, which would silently make the next rep truly-cold — but `until=24h` is **NOT** a guarantee that rep-touched records survive: being *served* from cache does not reliably refresh a BuildKit record's clock, and one 356.8 MB eviction cost **173 s**, which is why the baseline is a `p50` over `n ≥ 3` and never a mean), and the **union-apply** parallelism rule. Two rules it is worth reading it for alone: **a mid-campaign ENOSPC presents as the cryptic `redis exited (1)` (M239-F1), not as a disk error** — under a speed campaign it reads as "my lever broke the stack"; and **state the environment with every number** — the same Dockerfile and context yield **4.84 GB on `billion` (x86_64, containerd)** and **2.88 GB on an arm64 laptop (overlay2, which pays no unpack leg at all)**. **D-M255-1** is recorded here: the identical assert **hard-fails in `buildbench`** (a gate) and stays **advisory in the bring-up** (an operator) — two consumers, two contracts
- `corpus/ops/demo/latency-budget.md`: **The demo's performance budget — what "fast" means, and how it is measured** (v2.3 "cue to cue" M218). Before v2.3 there was **no** perf budget, baseline, gate, or even a **definition of "access"** anywhere in the corpus — while a presenter's click→login actually took **60–120 s**, and the corpus asserted in **four places** that it took *"~2–5 s, which we can't shorten"* (booked as M43-D5 with **zero deferrals recorded**, so it never entered a ledger and was never revisited across four releases). Defines **ACCESS** (the authenticated shell rendered + interactive with the hero's identity present), the **p95 < 5 s gate**, the **per-leg attribution model** (click → handshake/303 → SSR → clerk-js → client-gate → data-query), the baseline (**39.45 s** employee / **38.30 s** manager) and the shipped number (**cold p95 2413 / 1767 ms**, 5 consecutive cold reset-to-seed cycles), the harness contract (`rext stack-verify/e2e/run-latency.sh` — **never** gate on `networkidle`; **always** gate on a *fresh green* `autoverify.json`), and the **arithmetic signatures** that name a bug class before you read a line of code (a *blackholing* address ≈ `3 × 10.5 s + 6 s`; a *fast-failing* fetch ≈ `3 × 33 ms + 6 s`). **State the environment with every number** — the same defect cost ~6 s on a laptop and ~112 s on the tailnet VM. **v2.5 M236 adds:** `LATENCY_SCHEME=https` is mandatory when driving a `--public-host` stack; `autoverify.sh` needs `STACK_DIR`; the green-gate **age check parsed a UTC `ts` as LOCAL time on BSD**, so **west of UTC it aged a stale verdict as fresh** — the guard failed OPEN for half the world (fixed + regression-tested); and all four e2e runners now refuse a non-integer `N`
- `corpus/ops/demo/coverage-protocol.md`: **The demo-coverage iteration protocol** (v1.10 "method acting" M42e) — the **Playwright** sweep + triage + fix loop driving the **semantic believability gate** (real seeded content + per-section cardinality + persona self-consistency [role↔skills, menu==profile real-photo avatar, org name+logo] + 0 prod-eject escapes — supersedes the old `textLen>40` density check). The manifest-driven section model + the fix-surface routing table (empty→`stack-seeding`; content-error→`stack-snapshot` serve-grant; out-of-demo link→injection link-rewriting; runtime-computed→crawl-scope) + the disclosed-presenter-note allow-rule. The Playwright harness lives in rext `stack-verify/e2e/` (the first non-Go rext dev/test dep). Drives the per-vantage coverage gates M42e (employee) + M42m (manager). **v2.5 M236 adds the SECOND sweep it governs — the content-stories `(session × action)` LANDS sweep**: an exact-path visit per (cockpit seat × manifest result path), **six** render shapes selected **by ROUTE, never by keyword**, a **fail-CLOSED** reading (an empty ledger is a FAILURE, not a 0/0 pass) + an `EXPECTED_PAIRS` denominator pin, and a deliberate **reversal** of this protocol's own `skipPaths` `/result/` exclusion (the pages it exists to prove were the pages the rule excluded). 29/29 proven cold on `billion`. **v2.6 M244 re-proves the sweep LIVE at the grown denominator — 47/47 landed of the 49-pair count (M241's EN/IT growth moved it 29→49; 2 Bunny-absent voice *player* cells held presence-only), cold reset-to-seed on `billion`.**
- `corpus/ops/demo/playthroughs.md`: **The functional-flow e2e runbook** (the Playthroughs pillar, v2.0 "opening night" M202) — a **Playthrough is an automated actor that IS the user**: it logs in as a seeded hero, plays a real journey end-to-end, and proves the platform delivered the outcome. Proves **function** (the hero can *do* the thing) where `coverage-protocol.md` proves **presence** (every page *shows* real content). The manifest model (Products → Stories → Use Cases → Playthroughs) + the light validator (both-way id integrity + precondition-coverage + the `datadna` closure gate), the per-surface page-object/locator layer (semantic-by-default + a find-only landmark registry; re-pin **O(surfaces), not O(tests)**), the dedicated **decoupled** seed (`pt-world`, test data ≠ demo data, **three** private orgs — Org C, `narrative: ai-readiness`, added at v2.3 M219) + the **reset-to-seed** lifecycle (the real `--reset`, additive re-seed FORBIDDEN, N=0-guarded), the **serial-default** runner (`workers:1`, single shared org-scoped Postgres), and the **4-state reporting map** (`passing`/`failing`/`unimplemented`/`unimplementable-without-platform-edit` — the last escalates, never edits the platform). Reuses the M37 cockpit seat-switch for hero login + the M42 e2e foundation (never forked). Also **the iteration protocol M203/M204** (the coverage milestones) follow. Section `rext playthroughs/` (mixed Go + TS toolchain). Zero platform-repo edits. **Shipped v2.0 (M202 foundation → M203 employee → M204 manager): 10 live Playthroughs (6 employee + 4 manager) GREEN on cold reset-to-seed + 1 declared in-manifest TODO** (the assign-WRITE half). **v2.3 M219 adds the `ai-readiness` product** (6 employee + 4 manager + 4 AI-readiness = 14), and **v2.4 M225 adds `hiring.yaml` → 15 live Playthroughs** (`pt-hiring-recruiter-compare`, the recruiter-vantage candidate comparison), and **v2.6 M243 lands the assign-WRITE half** (`pt-assignment-assign` — the manager assigns a skill path to a member and the `organization_assignments` row is written + read back), flipping the last TODO, and **v2.7 M252 adds `studio-builders.yaml`** (Product **"Studio"**, studio-desk's first-ever manifest entry — `pt-studio-advanced-generate` + `pt-studio-guided-generate`, both driven by the org-admin manager hero). **v2.8 M256 "playthrough sharpening" grew it to 30 live Playthroughs + 1 declared TODO** — the **org-admin** product 0 → 4 of 4, the **onboarding** product 0 → **4 of its 5 CURATED use cases landed** (the 5th, the self-import journey, carries a machine-checked `will-not-build` **verdict** instead: its only advancing path scrapes a live third-party profile, so its RED would read as a product regression), plus `workforce-org-feedback`, `skillpath-bookmark` and the suite's first `outcome: blocked` Playthrough. **The corpus stands at 30 live Playthroughs + 1 verdicted TODO across 31 manifest use cases and 10 products** — `playthroughs.md` is authoritative for this count. (Onboarding's manifest share reads 6 declared / 5 live, because M256 also added one **net-new, non-curated** use case; the curated and manifest denominators are deliberately distinct and must not be conflated.) **v2.5 M236 draws the Playthrough-vs-content-story line:** a Playthrough PLAYS forward (the hero performs the actions that produce an outcome); a **content story** is already played (cloned from a real prod session) and what must be proven is that its **result surface renders**. There is nothing to play, so there is no Playthrough — that proof lives in the content-stories sweep in `coverage-protocol.md`
- `corpus/ops/demo/content-stories-routes.md`: **The content-stories feasibility spike + per-product result-route map** (v2.5 "the playbill" M231 — a HARD go/no-go barrier, the one that gates the whole Thread-B "Content stories" build chain M232→M236). For each content product × {player, manager} it enumerates the exact result route and **classifies each by prove-by-render** (renders-from-seed | runtime-computed-blank | needs-demo-patch | no-surface). The central unknown — does `/sim/<slug>/result/<sessionId>` recompute live (unseedable) or read a persisted row? — is **RESOLVED: a PERSISTED READ** (`jobsimulation/internal/graph/queries.resolvers.go:70` does plain Ent SELECTs of `validation_attempt_results` — no engine/LLM recompute on render), so a clone that INSERTs the result fan-out renders a full result. **Verdict:** Simulation (training/assessment/hiring) **GO**; Skill-path **GO PLAYER-ONLY** (M236 iter-07 refuted the manager half — next-web renders "Coming soon" with its results table commented out, so those 2 pairs are not landable and the gate denominator was corrected 31 → **29**); **Interview GO behind a PostHog-flag demo-patch** (`flag_interview_{player,manager}_report`); **AI-labs OUT** (nil client, `grade_result` not GraphQL-exposed, `/labs/[id]` reads live → no seedable result surface → presence-only); **Academy IN** (backend-authoritative since v0.5 M2 → `academy_chapter_progresses` seedable via `app/cmd/academy-seed` — but **that binary is MOOT on a demo** (M236 iter-08): a demo academy has no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, so it serves its committed FS catalog and the seeded rows have no reader; the demo CTA is a real `/courses/<slug>` link, and the grid renders 65 real cards with 0 Draft chips). Also carries the **generalized manager-view MIRROR trap** (`local_jobsimulation_sessions` / `local_skill_path_session` — seed the mirror or the manager scoreboard is blank), the **prod-session sourcing + anonymization contract** (pin by `sessions.id`, source only public-anchored `sim_id`, the free-text scrub surface), and the **public-sim-by-modality catalog** (77 voice / 65 code / 30 document public sims). Zero platform-repo edits; the copy is M232. **v2.6 grows the map: M240 media-fidelity (the recorded call by-reference + the inline document body — `media-substrate-spec.md`), M241's EN/IT language axis moved the landable denominator 29→49, and M244 proved 47/47 land live on `billion` (2 voice player cells presence-only)**
- `corpus/ops/demo/session-clone-spec.md`: **The session-clone / sourcing seeder — the write side of Content stories** (v2.5 "the playbill" M232, consumes M231's contract). The `ContentStorySeeder` (`rosetta-extensions/stack-seeding`) **COPIES real production job-simulation sessions** into a demo org: the REAL result-fan-out CONTENT (LLM feedback, transcript, submission, interview report — the interesting free-text) is **copied** from the pinned session (authoring-time — `cmd/content-capture` reads prod **read-only** via `~/.pgpass`, streaming content prod→scrub→fixture; it never enters an agent's context) and **SCRUBBED best-effort** of detectable PII (real actor names + source org → `<<ACTOR_i>>`/`<<ORG>>` placeholders the seeder fills with the demo persona/org; emails/phones/urls redacted; `package scrub`, tested). **NOT provably clean** — free-text scrubbing is imperfect; **residual re-identification risk is real and ACCEPTED by the data-controller (2026-07-19)**, the control being the **VPN/tailnet scope** (`safety.md` §3.8). **Re-tenanted** into the first Workforce org, **owner = a distinct non-hero MEMBER slot** (owner-is-player-vantage, never a manager seat), **source-pinned** (deterministic reseed; disclosed in `seed-generation-manifest.yaml`'s `content_sessions` block). Replays the full result fan-out (`jobsimulation.sessions` + the `local_jobsimulation_sessions` **MIRROR** + attempt/skill/criterion/**check** results with the REAL skill node-ids + REAL feedback + transcript **actors/interactions** [DB action_type ∈ {email,call} only] + the net-new **CODE** [`code_submissions`] / **DOCUMENT** [`collaborative_assets`] substrate + the **INTERVIEW** `interview_extraction_results` report), all G14-valid. Plus the **two sha-pinned interview-flag-gate demopatches** (`next-web-interview-flag-{container,result}` — the M219 aireadiness twin; no PostHog on a demo ⇒ no rollout gate; wired into up-injected.sh's both frontend builds). The bounded read-side exception `safety.md` §3.8 records. Zero platform-repo edits. (M233 = the manifest projection, M234 = the cockpit tab, M235 = prove-it-lands, **M236 = prove-on-billion — the live proof**). **v2.6 extends the write side: M240 media fidelity (the recorded-call/document media half — `media-substrate-spec.md`) + real per-session language, M241 EN/IT counterparts (denominator 29→49), M244 proved 47/47 live on `billion` (2 voice player cells presence-only)**
- `corpus/ops/demo/media-substrate-spec.md`: **The media substrate — recorded interview VIDEO + document bodies in a content-story demo** (v2.6 "sound check" M240) — the MEDIA half of Content stories (`session-clone-spec.md` copies the *free-text*; this covers the *media* that makes a session playable). Two facets, two shapes: **(1) the recorded call = a Bunny.net CDN REFERENCE, not a blob** — `jobsimulation.sessions.chime_status` is the render gate (`'completed'` = playable, `'not_available'` = no recording, the faithful default) + `ChimeRecording.bunny_video_id` resolves to an MP4 in a **Bunny.net Stream pull-zone**; **what the DEMO serves is the Bunny reference — no media byte is ever copied into a demo, and none lives in the platform DB** — which is why the old `DEF-M10-01` S3 read is neither necessary (the demo streams by reference) nor sufficient (S3 bytes carry no `bunny_video_id` and no `chime_status`, and the render gate reads both). **NB the platform's own copy IS in prod S3** — `jobsimulation` writes the composited MP4 there and reads it back before uploading to Bunny (M257x iter-86 retracted the flat *"never in prod S3"*; S3 is the origin, Bunny the delivery copy). Rendered **by-reference at render time only** (Bunny-CDN → demo-server → browser; no media byte is ever pre-copied). The recorded pool is **almost entirely HIRING interview VIDEO of real candidates**, so v2.6 ships **voice presence-only** (`chime_status='not_available'` IS the deliverable; the real-video exhibit was Bunny-key-blocked → dispositioned player-presence-only at M244, `DEF-M240-01`). **(2) the document body = inline `validation_criterion_results.input_data.text_document`** (a text field, **not** an S3 blob) — copied + scrubbed like the transcript, **fully landed** (M240 Defect 3), which closes the "is the document body a blob?" question: it is not. **PII discipline (load-bearing):** customer media — recorded video, audio, faces, document bodies — must **never** enter an agent's context; you ORCHESTRATE the tooling and never view the media, and Bunny keys are handled values-blind. Safety contract: [`safety.md` §3.8.1](corpus/ops/safety.md) (the raw-media amendment + the 2026-07-21 data-controller VIDEO sign-off). Zero platform-repo edits
- `corpus/ops/demo/content-stories-spec.md`: **The content_products manifest + honesty gate — the manifest half of Content stories** (v2.5 "the playbill" M233, consumes M232). `stackseed --content-export` PROJECTS a **`content-manifest.json`** (the content analog of `cockpit-manifest.json`) the 2nd "Content stories" cockpit tab reads: a **`content_products[]`** block — per content product (Simulation / Skill-path legacy / Skill-path new / AI-labs), the played sessions each with a **player + manager seat key**, a **result path** (player `/sim/<slug>/result/<sessionId>`; manager `/enterprise/activity-dashboard/<kind>/<simId>/<membershipId>` — the last segment is a **MEMBERSHIP** id, not a user id (M236 iter-05: `GetMembership(userID)` → `ent: membership not found` → the whole query nulls while the header still renders, so the page *looks* populated and proves nothing)), `has_manager_view`, a per-product **app-base**, and a per-`sim_type` **icon**. **Single-sourced** from the SAME content-session fixture the `ContentStorySeeder` seeds from (the player seat OWNS the seeded session; the path names the seeder's derived session id — no drift): the player seat is the owner **non-hero MEMBER** (`content-player-<idx>`, M234 registers it in the roster); the manager seat is the host org's manager hero; `has_manager_view` downgrades to false (fail-closed) with no manager hero. The player route resolves by **text slug** (`jobSimulationBySlug`, not the sim uuid) → the fixture gained a public **`sim_slug`** (resolved read-only from the public catalog). Content-story sessions render in **apps/web** (host org is Workforce, regardless of sim_type). **Honesty-gated** (a checked-in canonical `presets/content-manifest.json` + a `CanonicalFileMatchesProjection` test, with teeth) and **fail-closed** (a session that can't form a real link is DROPPED with a reason + `ValidateContentManifest` fails loud — never a fabricated CTA; AI-labs presence-only projects no player link). A **SEPARATE JSON** (not a `seed-generation-manifest.yaml` block) because the cockpit reads JSON, not YAML (no PyYAML); the M232 `content_sessions` source-pins stay folded in the YAML manifest. Zero platform-repo edits. (M234 = the cockpit tab render + player-seat registration, M235 = prove-it-lands, **M236 = prove-on-billion**). **v2.6: M241 added the real per-session `Language` field + the cockpit **EN|IT** toggle (denominator 29→49) behind a fail-closed language-consistency gate; M242 regrouped the rows by requirement tuple; M244 proved 47/47 pairs land live on `billion`**
- `corpus/ops/demo/profile-completeness-spec.md`: **The whole-roster profile-completeness spec** (v1.10 "method acting" M44) — the DATA-DENSITY layer that bakes EVERY seeded member (and the managers), not just the heroes: trajectory-aware self-ratings (`user_skills`), the `CertificatesSeeder` + `ProjectsSeeder`, manager personal data, and an avatar + career for every fill-member (the `/enterprise/members` `memberships.picture_url` avatar fix, render-verified). Density only — the structural chain stays the `stories-spec.md` 7-table fan-out. Indexed from `demo/README.md` + `seeding-spec.md` + `stories-spec.md`
- `corpus/ops/demo/ai-generation-spec.md`: **The generation engine + gen-acceptance protocol** (v1.10 "method acting" M45) — a cheap LLM (gpt-4o-mini) turns a YAML **batch descriptor** into realistic per-member profiles: the `services/ai/` wrapper (EU-first routing + cost tracking), `blueprint.Batch` + `EffectiveBatches()` (pure Go-template mother-prompt expansion, NO LLM at parse time), `cmd/gen-batch` (mandatory `--max-cost` ceiling + `--max-concurrent` + `--call-timeout` + re-roll-on-malformed + hero-collision re-roll), and the `GeneratedBatchSeeder` — enforcing the **CODE-owns-structure / AI-owns-content** boundary (every generated role/skill name routes through the existing resolvers; non-resolving names **drop**, closure stays green, never fabricated). The measure→fix→accept iteration protocol (5-metric gen-quality gate). **The FIRST new third-party dep in the seeding module** (`ai v1.40.1` — a deliberate, user-acknowledged in-release supply-chain decision). Pairs with `cache-spec.md`
- `corpus/ops/demo/cache-spec.md`: **The prompt-hash cache** (v1.10 M45) — `.agentspace/.batchcache/batch-${hash}/member-${i}.json` keyed by the **MOTHER prompt** + the **taxonomy capture version** (invalidate on re-replay), atomic `.tmp`→rename writes, the `.lock` fence — so an unchanged batch descriptor **re-seeds byte-identical at $0**
- `corpus/ops/demo/seed-manifest-spec.md`: **The consolidated single-auditable seed+generation manifest** (v1.10b "fit-up" M52) — ONE checked-in `seed-generation-manifest.yaml` inlining the whole demo-data intent: the population (all 4 orgs + heroes), the **file-resident** mother prompt (extracted from the Go const to `blueprint/prompts/default_batch_prompt.tmpl`), the batch config (the MANDATORY `max_cost_usd` ceiling + concurrency + re-roll rules), and the snapshot sources — **cache + generated data EXCLUDED**. A PROJECTION of the canonical presets (honesty-gated so it can't drift), emitted by `stackseed --manifest-export`, served by the presenter cockpit's **[Download seed manifest]**. So an auditor reads the entire seed+gen intent in ONE place without reading Go
- `corpus/ops/demo/recipe-snapshot-world.md`: The **set-dressing recipe** — capture→replay the real public library so a demo world's catalog + content templates are real
- `corpus/ops/demo/recipe-browser-login.md`: **Interactive browser login (Clerk-free)** — open a browser, log in as the demo user with **no real Clerk**, and land in a seeded org where authorized routes return 200. Carries the `api.clerk.com` cert-redirect + minted-publishable-key injection recipes
- `corpus/ops/demo/recipe-enterprise-onboarding.md`: **Enterprise org onboarding** — a believable enterprise customer org: an admin logs in to a populated workforce (hundreds of members with roles, tiers, months of activity). The "this is what your org looks like in Anthropos" demo
- `corpus/ops/demo/recipe-skill-progression.md`: **Multi-month skill progression** — growth over time: members who ran simulations and skill paths across months, passes and fails, so the workforce-growth and skill-verification views have a real timeline. Believability lives entirely in the **backdated activity**
- `corpus/ops/demo/tailscale-serve.md`: **The remote-access runbook** (v2.2 "panorama" M212–M215; remote reach flipped **default-on for the demo path** at v2.3 M220 — D-DESIGN-3, which SUPERSEDES v2.2's D-DESIGN-1 for that path) — make a demo reachable from **another machine on a Tailscale tailnet** (run it on a Tailscale VM, e.g. `billion.taildc510.ts.net`; a teammate with Tailscale up browses it end-to-end). The remote-reach flow — **default-on for `/demo-up`, opt-out via `--no-public-host`**; **`/dev-up` stays opt-in** via `--public-host <magicdns>` (v2.3 M220 D-DESIGN-3): one trusted **`tailscale cert`** HTTPS origin (Clerk needs a secure context) fronted by per-offset-port `tailscale serve`, the `CORS_EXTRA_ORIGINS` https trio + the ant-academy `allowedDevOrigins` sha-pinned patch, and the **fresh-Linux-VM** host prereqs (Go + atlas + tailscale operator) the tooling pre-flights/auto-handles/fails-loud on. **The FIRST live remote Linux-VM deploy** — proven end-to-end for both hero vantages (employee `maya-thriving` → `/profile`, manager `dan-manager` → `/enterprise/workforce`) on a trusted LE cert, 0 ejects, cold reset-to-seed reproducible; the F1–F12 host-deploy finding set + safety framing. Tooling + docs + a flag only — zero platform-repo edits

### Updating the Platform
- `corpus/ops/update_guide.md`: Sync code, dependencies, and database schemas (daily/weekly/full scenarios, conflict handling, migrations, image rebuilds). Driven by `/stack-update`

### Architecture Documentation
- `corpus/architecture/architecture_overview.md`: High-level system design
- `corpus/architecture/platform-migration-status.md`: **Where the microservice-into-`app` consolidation actually is** — one row per service the platform has ever had, **two states per row** (production vs a fresh local stack), every claim cited to a sha or `file:line`, plus the **net-new** org repos that appear in neither `repos.yml` nor the corpus. **Machine-fenced against the platform's own `repos.yml` in both directions** (`rosetta-extensions/stack-core/platform_alignment_guard.py`, v2.8 M257x) — a service entering *or leaving* the clone set turns a guard RED. **Read it before trusting any per-service claim below about whether that service still runs**; the merge banner in this file is prose and the map is fenced
- `corpus/architecture/org-repos.md`: **The `anthropos-work` org repo register — all 93, measured 2026-08-07, each with a home and an ADVISORY verdict (nothing deleted or archived).** The denominator this corpus never had: it documented the ~13 repos a stack clones and had never enumerated the rest. **It settles the standing `cms` M810 question** (there is no `module "cms"` in `infrastructure`; `cms/terraform/main.tf:39` is orphaned dead code), and it is the home for `infrastructure`, `directus`, `judge0`, `metabase`, the five `livekit-agent*` repos, `sim-qa`, `hyper-studio`, `analytics-go` and **`anthropos-knowledge-base` — a second org corpus that contradicts this one on the taxonomy figures in 14 unsourced places**
- `corpus/ops/observability.md`: **The observability tier this corpus documented NOWHERE until 2026-08-07.** What the platform emits (**no metrics pipeline**; Sentry-protocol traces at 15 %; the error tier is a **self-hosted GlitchTip**) vs what `ant-observability` runs (live outside-in `product-monitoring/` asserting on **body content**, because gqlgen, Next ISR and the LB health check can all return 200 for a failure)
- `corpus/architecture/service_taxonomy.md`: Three-tier service categorization
- `corpus/architecture/frontend_architecture.md`: Next.js monorepo deep dive
- `corpus/architecture/external_services.md`: Clerk, Directus, GraphQL, AI providers, LiveKit, Chime
- `corpus/architecture/dependency_map.md`: Service inter-dependency matrix with Redis Streams events
- `corpus/architecture/shared_libraries.md`: the internal Go libraries. **Its subject set is the five historical "shared libraries" and that is NOT `app`'s require set** — measured at `app` `3eaadae6`, `app/go.mod:14-18` is `analytics-go`, `colony`, `proto`, `storage`, `taxonomy`. `ai` was folded into `app` (`1e457fa70`); `authn` ships inside colony as `colony/authn` and is a dependency of no service
- `corpus/architecture/security_compliance.md`: Security, data protection, EU compliance, multi-tenancy
- `corpus/architecture/ai_architecture.md`: AI models, provider routing, voice engine, recording, cost tracking
- `corpus/architecture/alignment_testing.md`: The alignment test class + framework (`rosetta-extensions/alignment/`) — measuring how faithfully a mirror engine (e.g. Clerkenstein) reproduces a source engine as a 0–100% score

### Service Documentation
- **`corpus/services/README.md`: the enumerated index of all 27 service docs** — start here rather than guessing a filename. Grouped into core backend (Tier 1), frontends & gateway, cross-cutting subsystems / `app`-owned domains (AI-readiness, Course Builder, AI Labs + credits, Ask/Talk-to-Data, Academy backend, hiring, Clerk, Clerkenstein), and archived/merged redirects (`skiller`, `skillpath`, `chronos`, `intelligence`)
- `corpus/services/TEMPLATE.md`: the pattern each doc follows — Role, Architecture & Code Map, Interface Discovery, Local Development, Testing
- `corpus/ops/platform_repo.md`: The `platform` orchestrator repo (Make targets, profiles, compose, repos.yml)

### Tools & Development
- **`corpus/tools/README.md`: the tools index**
- `corpus/tools/toolchain_overview.md`: Development tools registry
- `corpus/tools/anthropos-labs.md`: The internal experiments hub (`anthropos-work/experiments`) — PoCs and prototypes, not part of the platform

## Working with Service Code

### Go Services (Backend, CMS, Sentinel, etc.)

Common development pattern:
```bash
# Setup (first time only)
make setup    # Install tools: mockgen, ent, atlas
make gen      # Generate code from protobuf/ent schemas

# Database migrations (when schema changes)
atlas migrate apply --env local

# Run locally
go run .

# Run tests
go test ./...
```

Key directories in Go services:
- `rpc.go`: Main RPC server implementation (entry point for API)
- `internal/data/ent`: Database schema and ORM code
- `internal/app`: Component wire-up
- Domain-specific folders: `internal/organization`, `internal/user`, etc.

### Frontend (Next.js Monorepo)

```bash
# Install dependencies
pnpm install

# Run development server
pnpm dev

# Build
pnpm build

# Run tests
pnpm test
```

### Studio Services

**Studio-Desk** (TypeScript):
```bash
cd studio-desk
npm install
npm run dev    # Runs on localhost:9100 (frontend) and localhost:9000 (backend)
```

**Studio-Room** (Python):

> **⚠️ There is no `studio-room` directory to `cd` into, and this block contradicted the Tier-2 section
> above until M257x iter-236.** The repo is `anthropos-studio-room`, it is **not in `repos.yml`**, and
> `make init` does not clone it — it is pulled into the `app` **image** by CI (`additional_repo`), which
> is what the Tier-2 entry means by *"never a standalone deployment."* Its root **is `app/studio/`**, and
> that path is **`.gitignore`d in `app` itself** (`app/.gitignore:78-79`, *"pulled at build via
> additional_repo"*) — so it is absent from a fresh `make init` clone and present only on a box where a
> build or a hand-clone populated it. `cd app/studio` is the real path when it is there; see
> `corpus/services/studio-room.md`.

```bash
cd app/studio          # NOT `cd studio-room` — see the warning above
pip3 install -r requirements.txt
python3 gen.py --media simulation
# ⚠️ There is NO --template flag. gen.py's parse_known_args SILENTLY ABSORBS it, so
# `--template default` succeeds and generates something unrelated. See corpus/services/studio-room.md
```

**Note**: Studio-Desk can also run containerized via `make up PROFILE=studio-desk`.

**Ant Academy** (Next.js 16 + Expo — native only, not in docker-compose):
```bash
# Web app
cd ant-academy/code
cp .env.example .env.local   # fill Clerk + AI keys (the app reads code/.env.local; see corpus/ops/setup_guide.md)
npm install
npm run dev            # next dev — port 3077

# Mobile app (optional, separate process)
cd ant-academy/mobile
pnpm install
pnpm run dev:web       # web preview on port 8555
```

See [Ant Academy service doc](corpus/services/ant-academy.md) for the full picture (auth gates, content layout, Cosmo AI assistant, repo-local authoring skills).

## Documentation Maintenance

### STEP RUN Guidelines

When updating `corpus/ops/setup_guide.md`, follow these principles:
1. **Verify Before Install**: Include commands to check if tools exist
2. **Verify After Install**: Include commands to confirm successful installation
3. **Request Confirmation**: Document where user approval is needed
4. **Document Improvements**: Add troubleshooting entries when issues are discovered

### Interconnected Documentation

These files must be maintained together:
1. `corpus/ops/setup_guide.md`: Detailed setup instructions
2. `corpus/ops/setup_github_guide.md`: GitHub SSH access configuration
3. `corpus/ops/run_guide.md`: Platform startup instructions
4. `corpus/ops/webhook_setup.md`: Clerk webhook tunnel configuration
5. `corpus/ops/update_guide.md`: Update instructions
6. `.claude/skills/dev-up/SKILL.md`: The consolidated dev build + start + set-dress skill (← setup-platform + start-platform)
7. `.claude/skills/setup-github/SKILL.md`: GitHub SSH setup skill
8. `.claude/skills/stack-update/SKILL.md`: The stack code/deps/schema sync skill (← update-platform)
9. `.claude/skills/update-knowledge/SKILL.md`: Corpus documentation skill
10. `corpus/ops/secrets-spec.md`: The secret-provisioning spec (the source-of-truth `/stack-secrets` reads) — paired with `setup_guide.md` (which now points to `/stack-secrets` instead of the manual `.env` hand-copy) and `safety.md` (the values-blind / `DIRECTUS_TOKEN`-non-rearm clause)
11. `.claude/skills/stack-secrets/SKILL.md`: The values-blind secret-provisioning skill (drives the `stacksecrets` CLI at its pinned tag)

**When to use update-knowledge**: After discovering new platform elements, receiving setup feedback, or finding documentation gaps. The skill performs a corpus-wide sweep to ensure all relevant sections are updated.

### Modus Operandi

Project Rosetta follows strict iterative reverse engineering:
1. **Iterative & Goal-Oriented**: Clear, achievable goals per iteration
2. **Autoconsistent & Discoverable**: Self-contained corpus, new agents can start from README
3. **Recreation Standard**: Documentation quality measured by ability to recreate full dev environment from scratch
4. **Dual-Level Documentation**: High-level (for PMs) + Deep dive (for engineers)

### Service Documentation Template

Follow `corpus/services/TEMPLATE.md` when documenting services:
- Role & Responsibility
- Architecture & Code Map
- Interface Discovery
- Local Development
- Testing

## Repository Structure

```
rosetta/
├── corpus/                    # All documentation
│   ├── architecture/          # System design docs
│   ├── services/              # Per-service documentation
│   ├── ops/                   # Operations guides (setup, run, update)
│   └── tools/                 # Development tools registry
├── stack-dev/                 # Git-ignored DEV-stack workspace (one of the stack-*/ family)
├── stack-demo/                # Git-ignored DEMO-stack workspace (+ its rosetta-extensions clone)
├── .agentspace/               # Git-ignored: skill output + the rosetta-extensions authoring copy
├── .claude/skills/            # Claude Code automation skills
└── README.md                  # Project overview and status
```

## Critical Rules

- **Work inside a `stack-*/` workspace** (e.g. `stack-dev/`) when dealing with actual platform code — never in the corpus
- **All stack-operating tooling lives in `rosetta-extensions`** — built/tested in the `.agentspace/rosetta-extensions/` authoring copy and tagged, then consumed per-stack via a pinned-tag clone; never scattered in `rosetta`, never authored ad-hoc inside a stack dir
- **Never commit `.env` files** to any repository
- **Update documentation immediately** when discovering gaps or better approaches
- **Verify against actual code** - don't assume documentation is 100% correct
- **Maintain dual-level depth** - both PM-friendly and engineer-friendly explanations

## Quick Start for New Developers

1. Read `README.md` for project overview
2. Follow `corpus/ops/setup_guide.md` to build environment + `corpus/ops/run_guide.md` to start it (or use `/dev-up`, which drives both)
4. Read `corpus/architecture/architecture_overview.md` for system understanding
5. Consult `corpus/services/` for specific service details
