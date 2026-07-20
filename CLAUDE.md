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
  Directus boots on an offset port + `cms` is cut over → content self-contained); without it the stack
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
executable-tooling monorepo that **operates** stacks — sections: `clerkenstein`
(the Clerk mock), `demo-stack`, `dev-stack`, `stack-injection`, `stack-core`,
`stack-seeding`, `stack-snapshot`, `stack-verify`, `alignment`. `rosetta` documents *how the platform works*;
`rosetta-extensions` is *the tooling that spins up, injects, and seeds copies of
it*. It has **two clone roles**:

- **Authoring copy → `.agentspace/rosetta-extensions/`** — the single working clone
  you spawn on demand to **read / build / test** the tooling, then commit and
  **tag**. New tools are developed here.
- **Per-stack consumption copies → `stack-<role>/rosetta-extensions @ <tag>`** —
  each stack consumes the tooling at a **pinned tag** (reproducible). The
  `/dev-*`, `/demo-*`, `/stack-*`, and `/align-*` skills drive a stack's own clone.

**Policy:** all code/scripts that operate the corpus/platform on a spawned stack
live in `rosetta-extensions` — never scattered in the `rosetta` corpus, never
authored ad-hoc inside a stack dir. A new need/tool is built and tested in the
`.agentspace/rosetta-extensions/` authoring copy, tagged, then consumed per-stack
via its tagged clone. See [`corpus/ops/rosetta_demo.md`](corpus/ops/rosetta_demo.md)
and [`corpus/services/clerkenstein.md`](corpus/services/clerkenstein.md).

## Architecture Overview

### Three-Tier Service Model

**Core Backend Services (Tier 1)**: Go microservices

In the default local profile (`graphql`):
- Backend (`app`): Main API gateway and user management; also hosts the **AI-readiness** workforce subsystem (org-level AI-capability diagnostics — see `corpus/services/ai-readiness.md`) **and the skills domain** — taxonomy (60K skills, 18K roles), assessment, AI skill matching, and vector embeddings (RAG), absorbed from the former standalone Skiller service (its Ent models now live in `app`, data in the `public` schema; the old `skiller` DB schema is legacy). The skiller RPC surface (GetSkills, GetSkill, SearchSkill, MatchSkill, GetJobRole) is served by `app` — consumers keep the env var, re-pointed at `SKILLER_RPC_ADDR=http://backend:8083` (local; `http://backend:8081` in prod terraform); the `skiller` git repo still exists but is decommissioned
- CMS: **The content layer** — owns the authored CONTENT / DEFINITIONS (skill paths, simulation blueprints, the content library), wrapping Directus as a proxy + business-logic + cache layer; **and embedded studio-room AI generation pipeline** (`cms/studio/` is the `anthropos-studio-room` repo, cloned via `cd cms && make init-studio` and gitignored — a submodule-style pattern, not a real `.gitmodules` entry). **NB: CMS — not the like-named `skillpath`/`jobsimulation` services — owns skill-path and simulation content** (content-vs-runtime-state split below)
- Sentinel: Authorization only (Casbin RBAC/ABAC) — authentication is Clerk + the `authn` middleware in each service, not Sentinel
- Jobsimulation: **Runtime/session engine** that *runs* AI simulations (voice, chat, code, documents) and emits completion events; the simulation *definition/blueprint* it runs is CONTENT fetched from CMS by ID (`cms.GetSimulation` Connect-RPC). It holds run/session state — not content
- Skillpath: **Runtime/session engine** that tracks per-user progression *state* (`SkillPathSession → ChapterSession → StepSession`, progress %, completion). The skill-path *content* (chapters → steps, curators, skills-to-verify) lives in CMS/Directus and is fetched by ID via `CMS_RPC_ADDR`. It holds no content
- Storage: File/blob storage management
- Roadrunner: Code execution proxy to Judge0 sandbox
- Gotenberg: Office-doc → PDF conversion (third-party image; consumed by `app/internal/converter/gotenberg.go`)

Available in other profiles but NOT started by default:
- Messenger (`messenger` profile): Email notifications via Brevo (Sendinblue)
- CustomerIO Sync (`customerio-sync` profile): Background data sync to Customer.io. Unique build pattern — built directly from GitHub URL, not cloned locally.

Production-only / deployed-only (not in local docker-compose):
- db-backup: Scheduled PostgreSQL backups (every 6h) to S3, Azure, Hetzner

Archived / merged (removed from local orchestration; repo dirs may still exist on disk):
- Chronos (was: scheduling & time-based events) — removed via platform commit `045857c`
- Intelligence (was: background data sync between backend and skiller schemas) — removed via platform commit `fdfa189`
- Skiller (was: skills taxonomy, assessment, embeddings) — **merged into `app`** (July 2026, v2.1 "quick change"); domain now in the `public` schema, `skiller` repo decommissioned, no skiller container/subgraph. See `corpus/services/skiller.md` + the `backend.md` fact-sheet

**Shared Libraries** (imported as private Go modules — **not** cloned by `make init`/`repos.yml`; pulled at Docker build via `GH_PAT`/`GOPRIVATE`). See `corpus/architecture/shared_libraries.md`.
- colony: Platform framework (logging+Sentry, DB, Redis, GraphQL/RPC servers, middleware, pub/sub via Watermill); **also contains `authn`**
- proto: Protobuf definitions (RPC contracts) + hand-written domain types
- ai: AI provider wrapper behind one `ai.AI` interface (OpenAI, Azure, Anthropic, Bedrock, Mistral). NOTE: cost tracking lives in `app/internal/aiusage`, and EU-first routing lives in each consumer's wrapper — **not** in this library
- authn: Clerk JWT authentication — now shipped **inside colony** as `colony/authn` (standalone `authn` repo is legacy)
- taxonomy: **node-id library** (`NodeID` type + ID generation/validation) — **not** a dataset; the 60K-skill/18K-role data lives in `app` (backend — the `public` schema, since the skiller→app merge)

**Studio Services & Standalone Internal Apps (Tier 2)**: Content creation tools + internal-only apps
- Studio-Desk (TypeScript/Vite/Express): Design tool for creating simulation blueprints (repo: `studio-desk`)
- Studio-Room (Python/Asyncio): AI-powered content generation pipeline (repo: `anthropos-studio-room`). **Embedded inside the cms container** as `cms/studio/` via `cd cms && make init-studio`; no longer a standalone deployment.
- Ant Academy (Next.js 16 + Expo): Internal learning portal for `@anthropos.work` employees (repo: `ant-academy`). **Vercel-deployed standalone — not in docker-compose.** **NOT in `repos.yml` (by design — v1.10b M49 #5)** — so `make init` does **not** clone it. For a **demo**, `ensure-clones.sh` clones it **explicitly** (phase d2, non-fatal — `repos.yml` lives in the ephemeral platform clone, so editing it is non-durable + a platform edit); for **dev**, it's a manual `git clone`. Runs natively via `cd ant-academy/code && npm run dev` (port 3077). Auth via Clerk; **since v0.5.1 the course catalog is DB-authoritative** — read from the platform academy subgraph over GraphQL (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`), degrading to an **empty grid** when the endpoint is unset or the academy DB is empty (the demo "empty academy" root cause — the v2.5 M229/M230 thread). See `corpus/services/ant-academy.md`.

**External Services (Tier 3)**: Third-party integrations
- Clerk: User authentication (SaaS)
- Directus: Headless CMS (self-hosted)
- GraphQL/Cosmo Router: Apollo Federation v2 gateway (4 subgraphs: app, jobsimulation, cms, skillpath — the former skiller subgraph was removed; app's subgraph serves its types/queries, and `categoryTree`/`fullCategoryTree` were dropped, not ported)
- AI Providers: OpenAI, Anthropic, Mistral (EU-first routing)
- LiveKit: Real-time voice engine for simulations
- AWS Chime: Video/audio recording

**Frontend Applications**: Next.js 15 monorepo on Vercel (`next-web-app`; see `corpus/services/next-web-app.md`)
- Next Web App: Main user-facing application
- Hiring App: Recruiting and hiring workflows
- Mobile App: Expo/React Native mobile experience

### Communication Patterns

- **Core Services ↔ Core Services**: Connect-RPC + Redis Streams (via Watermill) for async messaging
- **Frontend/Studio → Backend**: GraphQL via Cosmo Router (Apollo Federation v2, 4 subgraphs)
- **External Integrations**: Clerk SDK + JWT middleware (authn library), Directus proxied via CMS service
- **AI**: EU-first routing implemented in each consumer's `internal/ai` wrapper, **not** the shared `ai` library (EU Azure default → US Azure via PostHog flag `flag_use_azure_us` → direct-OpenAI on HTTP 429; Anthropic always Bedrock `eu-west-1`). Cost tracking in `app/internal/aiusage`
- **Multi-tenancy**: Shared DB, shared schema with `organization_id` on every table; 3-layer isolation (DB, Sentinel auth, Clerk identity)

### Environment Configuration

**Platform services** share a **single centralized `.env` file** in the `platform` repository. Docker-based services do not need their own `.env` files.

**Studio-Desk** requires its own `.env` file (`studio-desk/.env`) with Clerk and OpenAI credentials copied from `platform/.env`.

**Ant Academy** requires its own `.env` file at `ant-academy/code/.env.local` (not the repo root — the React app reads only from `code/.env.local`). Reuse `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `CLERK_SECRET_KEY` from `platform/.env`, and add `OPENAI_API_KEY` / `ANTHROPIC_API_KEY` for the `/api/ai/chat` route. Set `REQUIRE_ORGANIZATION_MEMBERSHIP=0` for solo local dev to skip the org-membership gate.

Critical environment variables:
- `GH_PAT` (GitHub Personal Access Token — required for Docker builds to pull private Go modules)
- `CLERK_SECRET_KEY` (Auth — backend services)
- `OPENAI_KEY` (AI services)
- `VITE_CLERK_PUBLISHABLE_KEY` (Studio-Desk via Docker)
- `DIRECTUS_PUBLIC_BASE_ADDR` (Content)

### Makefile-Driven Workflow

The `platform` repository provides a Makefile as the single entry point for all development operations. All service repos are cloned as siblings via `make init` and Docker builds from local code.

```bash
# First-time setup
cd stack-dev/platform
make init              # Clone all repos defined in repos.yml
make up                # Build from local code and start (graphql profile)
make migrate           # Apply all database migrations

# Daily development
make pull              # Pull main on all repos (auto-stash dirty changes)
make status            # Git status across all repos
make up                # Rebuild and start (auto-builds from local code)
make down              # Stop all services
make ps                # Show running containers
make logs S=cms        # Tail logs for a service
make dev S=cms         # Stop container, develop natively
make reset-db          # Wipe DB, restart, re-migrate (WARNING: data loss)
```

Docker Compose profiles control which services start:

| Profile | Services |
|---------|----------|
| `graphql` (default) | All backend + Cosmo Router |
| `backend` | app only |
| `cms` | cms only |
| `frontend` | next-web-app (containerized) |
| `studio-desk` | studio-desk (containerized) |
| `all` | Everything |

Usage: `make up PROFILE=cms`

## Key Documentation Locations

### Setup & Onboarding
- `corpus/ops/setup_guide.md`: Complete environment setup instructions
- `corpus/ops/setup_github_guide.md`: GitHub SSH access configuration

### Running the Platform
- `corpus/ops/run_guide.md`: Start the platform locally after setup
- `corpus/ops/webhook_setup.md`: Configure Clerk webhooks for user/org sync

### Demo Environments (disposable, Clerk-free, seeded + set-dressed — v1.1 "show floor" + v1.2 "set dressing")
- `corpus/ops/safety.md`: **The tooling safety contract** — the consolidated read-side (tenant-data firewall + public predicates + read-only capture) + write-side (3-layer isolation guard + never-write-prod + n=0 guards + audit-proven zero pollution) statement. The *why-it-is-safe* anchor for the whole demo/dev family (v1.3 M15). **v2.3 M220 adds Part 3 — the exposure side**, the third axis: who can *reach* a demo and what they get. It is a **disclosure, not a third "never"** — a demo is an **unauthenticated, authz-weakened build** (Clerk verification disarmed in app/cms/jobsimulation/skillpath; the authz-skip demo-patch default-ON; the presenter cockpit a **password-free "become any seeded hero"** launcher), and **every demo container is published on `0.0.0.0` — all interfaces — on EVERY `demo-up`, flag or no flag** (a claim `tailscale-serve.md` **denied** until M220; the retraction is in place and fenced). What makes that defensible is Parts 1+2: **there is nothing behind the door** — no customer data can be in a demo, and a demo cannot write prod. Part 3 also carries the **supersession of v2.2's D-DESIGN-1** (*"public reach is never default-on"*) by v2.3's **D-DESIGN-3** — **demo path only; `/dev-up` stays opt-in**
- `corpus/ops/demo/README.md`: **The demo-env family index** — the up→snapshot→seed→use→down flow + recipes + presets
- `corpus/ops/demo/demo-up-defaults.md`: **The `/demo-up` defaults contract** (v2.3 "cue to cue" M220) — every knob and flag that controls a bring-up: **all 27 `DEMO_*`/`STACK_PUBLIC_HOST` env knobs + 10 CLI flags**, each with its real default and the exact `file:line` that reads it. **Derived from the parsers and fenced against them in BOTH directions** (`stack-core/demo_knob_guard.py`): a doc-promised flag with no parser entry is a **false promise**; a parser flag with no doc row is **undiscoverable**. Records the fact nobody had written down — **there are TWO entry points**: `up-injected.sh` (what `/demo-up` runs) takes **only** `<N>` + `--public-host` and **hard-errors `exit 1` on anything else**, while `--profile`/`--services` are flags of the separate `rosetta-demo` wrapper (the skill's `argument-hint` conflated them for releases). And the shape: **every feature knob is an opt-OUT (`DEMO_NO_*`, default `0`)** — so a bare `/demo-up N` **already** seeds the 4-org world (3 workforce + the M223 hiring org), the full UI tier, the cockpit, and set-dress. *"Pull all the data + seed the orgs" was always the default; the usual culprit is a **cold snapshot cache**, not a knob.*
- `corpus/ops/rosetta_demo.md`: The demo-stack lifecycle (bring-up, port-offset, Clerkenstein injection, teardown)
- `corpus/ops/seeding-spec.md`: The `stack.seed.yaml` blueprint + the **production-isolation boundary** (write-side) + the data-DNA (now **100%**, nothing waived)
- `corpus/ops/demo/stories-spec.md`: **The verified-skill chain reference** (v1.9 "storytelling" M34) — how a seeded *verified skill* (a hero's profile + Skill Spotlight chart + the claimed-vs-verified gap) is materialized as the **7-table fan-out** the `PersonaSeeder` writes: the DB-enforced vs inserted-but-invisible constraint landmines, the **G14** session-seeder fix (valid `SIMULATION_TYPE_*`/enum/token + continuous growth-arc score), the `user_level` (claimed side) requirement, the `TaxonomyRefs` resolver (real public node-ids, never fabricated), the `users.go` name/avatar/email patch, and the **seed-side closure gene** (`datadna measure-closure`). The believability spine; vertical slice (Maya). M35 adds the full multi-org Stories & Heroes model, **M36 the org Workforce-Intelligence dashboard surfaces** (the mapped→verified funnel + teams + role gap/mobility + succession + feedback + the org-scale claimed-vs-verified gap), M37–M38 the presenter cockpit, and v1.10 "method acting" layers the per-hero **profile depth**: M39 the profile identity (real name/avatar/org-domain email) and **M41 the `ProfileSeeder`** (work-history + education timeline + a claimed-but-unverified `user_skills` tail that widens the visible claimed-vs-verified gap). **v1.10b "fit-up" M51 adds the AI-readiness showcase org as a 3rd story** (org "Northwind Aviation", 200 members, heroes Aria COMPLETED / Ben STARTED / Dana manager) with **four** net-new AI-readiness seeders (`OrgSettingsSeeder` + `AIReadinessConfigSeeder` + `AIReadinessFunnelSeeder` + — net-new at **v2.3 M219** — the **interview-aggregated-report** seeder, without which the manager's four interview-findings blocks render headings with NO content) seeding **both** a `closed` and an `active` cycle (M219 refuted M51's "live-recompute never completes" premise — it takes **2.09 s**) + the `app-aireadiness-snapshot-loadmembers` read-path demo-patch; the seeder contract is `corpus/services/ai-readiness.md`
- `corpus/ops/demo/cockpit-spec.md`: **The presenter-cockpit UX spec** (v1.10 "method acting" M43) — the slick **light** login launcher a demo-giver drives (`rext demo-stack/cockpit.py`, served at `:7700`+offset): the card-per-hero layout + FontAwesome icons (free CDN), the **one unified [Log in as] CTA** per hero (logs in *and* lands on her per-role `jump_to` — no more separate [Jump]), the seed-manifest download, and the staged login-progress overlay — plus the deep-link contract, the standalone-served-panel (zero-platform-edit) model, and the future-feature surface. Graduates the M37/M38 cockpit mechanics scattered across `stories-spec.md` + `clerkenstein.md`
- `corpus/ops/secrets-spec.md`: **The secret-provisioning spec** (v1.6 "stage door" M27–M30) — provision every repo's target `.env` (`dev-N`/`demo-N`) from one secret source (dir/zip, default `.agentspace/secrets`), **values-blind** (no verb reads/echoes a value), verified by the 6-repo/56-gene **secret-coverage DNA** + the two-tier keep-listed gate. The source-dir/zip layout contract (zEnvs defence), the per-repo target-file map, alias-family vs distinct-similar rules, the waived class, N=0 guard + idempotency, the demo-aware check, and the `DIRECTUS_TOKEN` non-rearm safety (the fix16/17 class). Driven by `/stack-secrets`
- `corpus/ops/db-access.md`: **Production DB read access** (read-side) — the `/db-query` skill + the public-vs-customer boundary + the snapshot read foundation (v1.2)
- `corpus/ops/snapshot-spec.md`: The **`stack-snapshot` extension** (v1.2 M9a/M9b/M10) — capture the public taxonomy + Directus content once from a safe prod source, manifest-cache it in `.agentspace`, replay per-stack (`/stack-snapshot`); the tenant-data firewall + the `stacksnap` CLI + the snapshot-fidelity gate
- `corpus/ops/snapshot-cold-start.md`: **The cold-start capture runbook** (v1.3b M20) — the one case the cache can't shortcut: a fresh box with an empty cache + no safe `--dsn`. The sanctioned DSN-export / restore-a-`pg_dump`-then-`--dsn` path to fill the cache once per release (behind the capture-source policy + `AssertPublicOnly`), **why the wired `postgres` MCP is NOT a capture source** (it returns JSON rows, not COPY bytes), and how it slots into the auto-set-dress bring-up (replay-only, never capture)
- `corpus/ops/idempotency.md`: **The bring-up re-run safety contract** (v1.3b M17) — what happens when you re-run migrate / snapshot-replay / seed: each is safe-and-idempotent or fails loudly with a guard (replay TRUNCATE-then-reload, idempotent seed COPY + casbin `WHERE NOT EXISTS`, the fixed `--reset`, the `set -e` first-run-race hardening). The *run-it-twice* companion to snapshot/seeding-spec
- `corpus/ops/verification.md`: **The bring-up auto-verify safety net** (v1.3b M18) — every bring-up ends with a scoped, **non-fatal** `verify live` on the stack's **own offset ports**: cheap-win `/api/health` + `sentinel.casbin_rules > 0` asserts (the silent-403 catcher) then the full offset/project/scope-aware probe set, so "UP" means *verified-working*. Default-on; a verify bug never blocks a good stack. The *is-it-actually-working* companion to `rosetta_demo.md` + the `/test-platform` skill. **v2.5 M236 adds PRE-FLIGHT RUNG ZERO — *tagging is not publishing*:** a remote stack consumes `rosetta-extensions` only at a tag **fetched from origin** (M217 FATAL pin guard), so tooling that exists only in the local authoring copy is **unreachable** to it. M236's first iter found `billion` pinned to the previous release's tag with **0 of 13** `playbill-*` tags on origin — the feature under test could not be obtained at all. Verify the tag is *on origin* before any prove-it-live milestone
- `corpus/ops/demo/frontend-tier.md`: **The demo UI tier** (v1.3b M19) — `/demo-up` brings up next-web-app + studio-desk (per-demo **cached** Docker image from the **unmodified** Dockerfile, offset ports, minted-pk + offset-URL baked) + ant-academy natively (Clerk-free). The 12 GB Docker-VM prereq + non-fatal pre-flight, the honest "one ~3-min cached build per new `demo-N`" residual, the `--no-ui` (`DEMO_NO_UI`) escape, and the hard **zero-platform-repo-edit** line (repo = build context only). The *see-it-in-a-browser* completion of the demo family
- `corpus/ops/demo/demopatch-spec.md`: **The demo-patch mechanism — the sanctioned zero-platform-edit escape hatch** (v2.3 "cue to cue" M217). When a demo needs a fix with **no env/config/compose seam** (the value is baked into platform source), `demopatch` patches the demo's **own ephemeral clone** just before the image build and reverts it after — the *image* carries the fix, the clone is left git-clean, and the canonical `anthropos-work` repos are **never touched**. The **7 guards** (G1 path-assert · G2 drift-refuse + exactly-once anchor · G3 never-commit · G4 idempotent · G5 self-revert · G6 demo-only · **G7 apply post-condition**), the 10-key manifest schema, the **three apply vehicles** (the `app` patches target the build-scratch clone **outside** the workspace, so `demopatch`'s own G1/G6 correctly REFUSE them — two shell helpers re-implement the ladder against the same canonical manifest), the **chain rule** (`next-web-public-website-url`'s `pre_sha256` **IS** studio's `post_sha256` — it reads "DRIFTED" against a pristine file **by design**), and the **self-healing freshness gate**: *the anchor is the contract; the whole-file sha is only a baseline*. **Read it before adding or re-pinning any patch** — a silently-refused perf patch shipped a 76 s members grid for four releases
- `corpus/ops/demo/latency-budget.md`: **The demo's performance budget — what "fast" means, and how it is measured** (v2.3 "cue to cue" M218). Before v2.3 there was **no** perf budget, baseline, gate, or even a **definition of "access"** anywhere in the corpus — while a presenter's click→login actually took **60–120 s**, and the corpus asserted in **four places** that it took *"~2–5 s, which we can't shorten"* (booked as M43-D5 with **zero deferrals recorded**, so it never entered a ledger and was never revisited across four releases). Defines **ACCESS** (the authenticated shell rendered + interactive with the hero's identity present), the **p95 < 5 s gate**, the **per-leg attribution model** (click → handshake/303 → SSR → clerk-js → client-gate → data-query), the baseline (**39.45 s** employee / **38.30 s** manager) and the shipped number (**cold p95 2413 / 1767 ms**, 5 consecutive cold reset-to-seed cycles), the harness contract (`rext stack-verify/e2e/run-latency.sh` — **never** gate on `networkidle`; **always** gate on a *fresh green* `autoverify.json`), and the **arithmetic signatures** that name a bug class before you read a line of code (a *blackholing* address ≈ `3 × 10.5 s + 6 s`; a *fast-failing* fetch ≈ `3 × 33 ms + 6 s`). **State the environment with every number** — the same defect cost ~6 s on a laptop and ~112 s on the tailnet VM. **v2.5 M236 adds:** `LATENCY_SCHEME=https` is mandatory when driving a `--public-host` stack; `autoverify.sh` needs `STACK_DIR`; the green-gate **age check parsed a UTC `ts` as LOCAL time on BSD**, so **west of UTC it aged a stale verdict as fresh** — the guard failed OPEN for half the world (fixed + regression-tested); and all four e2e runners now refuse a non-integer `N`
- `corpus/ops/demo/coverage-protocol.md`: **The demo-coverage iteration protocol** (v1.10 "method acting" M42e) — the **Playwright** sweep + triage + fix loop driving the **semantic believability gate** (real seeded content + per-section cardinality + persona self-consistency [role↔skills, menu==profile real-photo avatar, org name+logo] + 0 prod-eject escapes — supersedes the old `textLen>40` density check). The manifest-driven section model + the fix-surface routing table (empty→`stack-seeding`; content-error→`stack-snapshot` serve-grant; out-of-demo link→injection link-rewriting; runtime-computed→crawl-scope) + the disclosed-presenter-note allow-rule. The Playwright harness lives in rext `stack-verify/e2e/` (the first non-Go rext dev/test dep). Drives the per-vantage coverage gates M42e (employee) + M42m (manager). **v2.5 M236 adds the SECOND sweep it governs — the content-stories `(session × action)` LANDS sweep**: an exact-path visit per (cockpit seat × manifest result path), **six** render shapes selected **by ROUTE, never by keyword**, a **fail-CLOSED** reading (an empty ledger is a FAILURE, not a 0/0 pass) + an `EXPECTED_PAIRS` denominator pin, and a deliberate **reversal** of this protocol's own `skipPaths` `/result/` exclusion (the pages it exists to prove were the pages the rule excluded). 29/29 proven cold on `billion`
- `corpus/ops/demo/playthroughs.md`: **The functional-flow e2e runbook** (the Playthroughs pillar, v2.0 "opening night" M202) — a **Playthrough is an automated actor that IS the user**: it logs in as a seeded hero, plays a real journey end-to-end, and proves the platform delivered the outcome. Proves **function** (the hero can *do* the thing) where `coverage-protocol.md` proves **presence** (every page *shows* real content). The manifest model (Products → Stories → Use Cases → Playthroughs) + the light validator (both-way id integrity + precondition-coverage + the `datadna` closure gate), the per-surface page-object/locator layer (semantic-by-default + a find-only landmark registry; re-pin **O(surfaces), not O(tests)**), the dedicated **decoupled** seed (`pt-world`, test data ≠ demo data, **three** private orgs — Org C, `narrative: ai-readiness`, added at v2.3 M219) + the **reset-to-seed** lifecycle (the real `--reset`, additive re-seed FORBIDDEN, N=0-guarded), the **serial-default** runner (`workers:1`, single shared org-scoped Postgres), and the **4-state reporting map** (`passing`/`failing`/`unimplemented`/`unimplementable-without-platform-edit` — the last escalates, never edits the platform). Reuses the M37 cockpit seat-switch for hero login + the M42 e2e foundation (never forked). Also **the iteration protocol M203/M204** (the coverage milestones) follow. Section `rext playthroughs/` (mixed Go + TS toolchain). Zero platform-repo edits. **Shipped v2.0 (M202 foundation → M203 employee → M204 manager): 10 live Playthroughs (6 employee + 4 manager) GREEN on cold reset-to-seed + 1 declared in-manifest TODO** (the assign-WRITE half). **v2.3 M219 adds the `ai-readiness` product → 14 live Playthroughs** (6 employee + 4 manager + 4 AI-readiness), 1 TODO still outstanding — flagged at the M219 close for an explicit fate at release close. **v2.5 M236 draws the Playthrough-vs-content-story line:** a Playthrough PLAYS forward (the hero performs the actions that produce an outcome); a **content story** is already played (cloned from a real prod session) and what must be proven is that its **result surface renders**. There is nothing to play, so there is no Playthrough — that proof lives in the content-stories sweep in `coverage-protocol.md`
- `corpus/ops/demo/content-stories-routes.md`: **The content-stories feasibility spike + per-product result-route map** (v2.5 "the playbill" M231 — a HARD go/no-go barrier, the one that gates the whole Thread-B "Content stories" build chain M232→M236). For each content product × {player, manager} it enumerates the exact result route and **classifies each by prove-by-render** (renders-from-seed | runtime-computed-blank | needs-demo-patch | no-surface). The central unknown — does `/sim/<slug>/result/<sessionId>` recompute live (unseedable) or read a persisted row? — is **RESOLVED: a PERSISTED READ** (`jobsimulation/internal/graph/queries.resolvers.go:70` does plain Ent SELECTs of `validation_attempt_results` — no engine/LLM recompute on render), so a clone that INSERTs the result fan-out renders a full result. **Verdict:** Simulation (training/assessment/hiring) **GO**; Skill-path **GO PLAYER-ONLY** (M236 iter-07 refuted the manager half — next-web renders "Coming soon" with its results table commented out, so those 2 pairs are not landable and the gate denominator was corrected 31 → **29**); **Interview GO behind a PostHog-flag demo-patch** (`flag_interview_{player,manager}_report`); **AI-labs OUT** (nil client, `grade_result` not GraphQL-exposed, `/labs/[id]` reads live → no seedable result surface → presence-only); **Academy IN** (backend-authoritative since v0.5 M2 → `academy_chapter_progress` seedable via `app/cmd/academy-seed` — but **that binary is MOOT on a demo** (M236 iter-08): a demo academy has no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, so it serves its committed FS catalog and the seeded rows have no reader; the demo CTA is a real `/courses/<slug>` link, and the grid renders 65 real cards with 0 Draft chips). Also carries the **generalized manager-view MIRROR trap** (`local_jobsimulation_sessions` / `local_skill_path_session` — seed the mirror or the manager scoreboard is blank), the **prod-session sourcing + anonymization contract** (pin by `sessions.id`, source only public-anchored `sim_id`, the free-text scrub surface), and the **public-sim-by-modality catalog** (77 voice / 65 code / 30 document public sims). Zero platform-repo edits; the copy is M232
- `corpus/ops/demo/session-clone-spec.md`: **The session-clone / sourcing seeder — the write side of Content stories** (v2.5 "the playbill" M232, consumes M231's contract). The `ContentStorySeeder` (`rosetta-extensions/stack-seeding`) **COPIES real production job-simulation sessions** into a demo org: the REAL result-fan-out CONTENT (LLM feedback, transcript, submission, interview report — the interesting free-text) is **copied** from the pinned session (authoring-time — `cmd/content-capture` reads prod **read-only** via `~/.pgpass`, streaming content prod→scrub→fixture; it never enters an agent's context) and **SCRUBBED best-effort** of detectable PII (real actor names + source org → `<<ACTOR_i>>`/`<<ORG>>` placeholders the seeder fills with the demo persona/org; emails/phones/urls redacted; `package scrub`, tested). **NOT provably clean** — free-text scrubbing is imperfect; **residual re-identification risk is real and ACCEPTED by the data-controller (2026-07-19)**, the control being the **VPN/tailnet scope** (`safety.md` §3.8). **Re-tenanted** into the first Workforce org, **owner = a distinct non-hero MEMBER slot** (owner-is-player-vantage, never a manager seat), **source-pinned** (deterministic reseed; disclosed in `seed-generation-manifest.yaml`'s `content_sessions` block). Replays the full result fan-out (`jobsimulation.sessions` + the `local_jobsimulation_sessions` **MIRROR** + attempt/skill/criterion/**check** results with the REAL skill node-ids + REAL feedback + transcript **actors/interactions** [DB action_type ∈ {email,call} only] + the net-new **CODE** [`code_submissions`] / **DOCUMENT** [`collaborative_assets`] substrate + the **INTERVIEW** `interview_extraction_results` report), all G14-valid. Plus the **two sha-pinned interview-flag-gate demopatches** (`next-web-interview-flag-{container,result}` — the M219 aireadiness twin; no PostHog on a demo ⇒ no rollout gate; wired into up-injected.sh's both frontend builds). The bounded read-side exception `safety.md` §3.8 records. Zero platform-repo edits. (M233 = the manifest projection, M234 = the cockpit tab, M235 = prove-it-lands, **M236 = prove-on-billion — the live proof**)
- `corpus/ops/demo/content-stories-spec.md`: **The content_products manifest + honesty gate — the manifest half of Content stories** (v2.5 "the playbill" M233, consumes M232). `stackseed --content-export` PROJECTS a **`content-manifest.json`** (the content analog of `cockpit-manifest.json`) the 2nd "Content stories" cockpit tab reads: a **`content_products[]`** block — per content product (Simulation / Skill-path legacy / Skill-path new / AI-labs), the played sessions each with a **player + manager seat key**, a **result path** (player `/sim/<slug>/result/<sessionId>`; manager `/enterprise/activity-dashboard/<kind>/<simId>/<membershipId>` — the last segment is a **MEMBERSHIP** id, not a user id (M236 iter-05: `GetMembership(userID)` → `ent: membership not found` → the whole query nulls while the header still renders, so the page *looks* populated and proves nothing)), `has_manager_view`, a per-product **app-base**, and a per-`sim_type` **icon**. **Single-sourced** from the SAME content-session fixture the `ContentStorySeeder` seeds from (the player seat OWNS the seeded session; the path names the seeder's derived session id — no drift): the player seat is the owner **non-hero MEMBER** (`content-player-<idx>`, M234 registers it in the roster); the manager seat is the host org's manager hero; `has_manager_view` downgrades to false (fail-closed) with no manager hero. The player route resolves by **text slug** (`jobSimulationBySlug`, not the sim uuid) → the fixture gained a public **`sim_slug`** (resolved read-only from the public catalog). Content-story sessions render in **apps/web** (host org is Workforce, regardless of sim_type). **Honesty-gated** (a checked-in canonical `presets/content-manifest.json` + a `CanonicalFileMatchesProjection` test, with teeth) and **fail-closed** (a session that can't form a real link is DROPPED with a reason + `ValidateContentManifest` fails loud — never a fabricated CTA; AI-labs presence-only projects no player link). A **SEPARATE JSON** (not a `seed-generation-manifest.yaml` block) because the cockpit reads JSON, not YAML (no PyYAML); the M232 `content_sessions` source-pins stay folded in the YAML manifest. Zero platform-repo edits. (M234 = the cockpit tab render + player-seat registration, M235 = prove-it-lands, **M236 = prove-on-billion**)
- `corpus/ops/demo/profile-completeness-spec.md`: **The whole-roster profile-completeness spec** (v1.10 "method acting" M44) — the DATA-DENSITY layer that bakes EVERY seeded member (and the managers), not just the heroes: trajectory-aware self-ratings (`user_skills`), the `CertificatesSeeder` + `ProjectsSeeder`, manager personal data, and an avatar + career for every fill-member (the `/enterprise/members` `memberships.picture_url` avatar fix, render-verified). Density only — the structural chain stays the `stories-spec.md` 7-table fan-out. Indexed from `demo/README.md` + `seeding-spec.md` + `stories-spec.md`
- `corpus/ops/demo/ai-generation-spec.md`: **The generation engine + gen-acceptance protocol** (v1.10 "method acting" M45) — a cheap LLM (gpt-4o-mini) turns a YAML **batch descriptor** into realistic per-member profiles: the `services/ai/` wrapper (EU-first routing + cost tracking), `blueprint.Batch` + `EffectiveBatches()` (pure Go-template mother-prompt expansion, NO LLM at parse time), `cmd/gen-batch` (mandatory `--max-cost` ceiling + `--max-concurrent` + `--call-timeout` + re-roll-on-malformed + hero-collision re-roll), and the `GeneratedBatchSeeder` — enforcing the **CODE-owns-structure / AI-owns-content** boundary (every generated role/skill name routes through the existing resolvers; non-resolving names **drop**, closure stays green, never fabricated). The measure→fix→accept iteration protocol (5-metric gen-quality gate). **The FIRST new third-party dep in the seeding module** (`ai v1.40.1` — a deliberate, user-acknowledged in-release supply-chain decision). Pairs with `cache-spec.md`
- `corpus/ops/demo/cache-spec.md`: **The prompt-hash cache** (v1.10 M45) — `.agentspace/.batchcache/batch-${hash}/member-${i}.json` keyed by the **MOTHER prompt** + the **taxonomy capture version** (invalidate on re-replay), atomic `.tmp`→rename writes, the `.lock` fence — so an unchanged batch descriptor **re-seeds byte-identical at $0**
- `corpus/ops/demo/seed-manifest-spec.md`: **The consolidated single-auditable seed+generation manifest** (v1.10b "fit-up" M52) — ONE checked-in `seed-generation-manifest.yaml` inlining the whole demo-data intent: the population (all 4 orgs + heroes), the **file-resident** mother prompt (extracted from the Go const to `blueprint/prompts/default_batch_prompt.tmpl`), the batch config (the MANDATORY `max_cost_usd` ceiling + concurrency + re-roll rules), and the snapshot sources — **cache + generated data EXCLUDED**. A PROJECTION of the canonical presets (honesty-gated so it can't drift), emitted by `stackseed --manifest-export`, served by the presenter cockpit's **[Download seed manifest]**. So an auditor reads the entire seed+gen intent in ONE place without reading Go
- `corpus/ops/demo/recipe-snapshot-world.md`: The **set-dressing recipe** — capture→replay the real public library so a demo world's catalog + content templates are real
- `corpus/ops/demo/tailscale-serve.md`: **The remote-access runbook** (v2.2 "panorama" M212–M215; remote reach flipped **default-on for the demo path** at v2.3 M220 — D-DESIGN-3, which SUPERSEDES v2.2's D-DESIGN-1 for that path) — make a demo reachable from **another machine on a Tailscale tailnet** (run it on a Tailscale VM, e.g. `billion.taildc510.ts.net`; a teammate with Tailscale up browses it end-to-end). The remote-reach flow — **default-on for `/demo-up`, opt-out via `--no-public-host`**; **`/dev-up` stays opt-in** via `--public-host <magicdns>` (v2.3 M220 D-DESIGN-3): one trusted **`tailscale cert`** HTTPS origin (Clerk needs a secure context) fronted by per-offset-port `tailscale serve`, the `CORS_EXTRA_ORIGINS` https trio + the ant-academy `allowedDevOrigins` sha-pinned patch, and the **fresh-Linux-VM** host prereqs (Go + atlas + tailscale operator) the tooling pre-flights/auto-handles/fails-loud on. **The FIRST live remote Linux-VM deploy** — proven end-to-end for both hero vantages (employee `maya-thriving` → `/profile`, manager `dan-manager` → `/enterprise/workforce`) on a trusted LE cert, 0 ejects, cold reset-to-seed reproducible; the F1–F12 host-deploy finding set + safety framing. Tooling + docs + a flag only — zero platform-repo edits

### Updating the Platform
- `corpus/ops/update_guide.md`: Sync code, dependencies, and database schemas
- `corpus/ops/update_checklist.md`: Progress tracker for updates

### Architecture Documentation
- `corpus/architecture/architecture_overview.md`: High-level system design
- `corpus/architecture/service_taxonomy.md`: Three-tier service categorization
- `corpus/architecture/frontend_architecture.md`: Next.js monorepo deep dive
- `corpus/architecture/external_services.md`: Clerk, Directus, GraphQL, AI providers, LiveKit, Chime
- `corpus/architecture/dependency_map.md`: Service inter-dependency matrix with Redis Streams events
- `corpus/architecture/shared_libraries.md`: The five internal Go libraries (colony, proto, ai, authn, taxonomy)
- `corpus/architecture/security_compliance.md`: Security, data protection, EU compliance, multi-tenancy
- `corpus/architecture/ai_architecture.md`: AI models, provider routing, voice engine, recording, cost tracking
- `corpus/architecture/alignment_testing.md`: The alignment test class + framework (`rosetta-extensions/alignment/`) — measuring how faithfully a mirror engine (e.g. Clerkenstein) reproduces a source engine as a 0–100% score

### Service Documentation
- `corpus/services/`: Individual service documentation following TEMPLATE.md pattern
- Includes the GraphQL gateway (`graphql-wundergraph.md`) and main frontend (`next-web-app.md`)
- Each service doc includes: Role, Architecture, Interface Discovery, Local Development, Testing
- `corpus/ops/platform_repo.md`: The `platform` orchestrator repo (Make targets, profiles, compose, repos.yml)

### Tools & Development
- `corpus/tools/toolchain_overview.md`: Development tools registry

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
```bash
cd studio-room
pip3 install -r requirements.txt
python3 gen.py --media simulation --template default
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
5. `corpus/ops/update_guide.md` / `update_checklist.md`: Update instructions
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
