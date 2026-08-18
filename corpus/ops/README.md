# Operations Guides

This directory contains guides for operating the Anthropos platform locally.

> ## ⚠️ Monolith merge — read before following any stack runbook
>
> `skiller`, `skillpath`, `jobsimulation` (jobsim-in-app), `cms`
> (cms-in-app v8.0, app v1.360.0), `storage` + `messenger` (v9.0 support-in-app) and
> `customerio-sync` — **seven** — are all **folded into `app`** and run in-process as the single `backend`
> service. **`roadrunner` was listed here as an eighth until M257x iter-137; it was deleted, not folded**
> (no `app/internal/roadrunner/` at any ref) — but for ops the consequence is identical: no container.
> What that changes for ops:
>
> * There is **no `cms`, `jobsimulation`, `skiller`, `skillpath`, `roadrunner`, `storage`,
>   `messenger` or `customerio-sync` container**, profile, port, or subgraph — `838d907`
>   (2026-08-05) deleted the last three. Use `backend` for all of them (`make logs S=backend`,
>   `make dev S=backend`).
> * The federation composes **one** subgraph, and there is no router to compose it with:
>   platform `2adcf71` deleted the `graphql` compose service, its `repos.yml` entry and its
>   clone. GraphQL is served by `backend` itself at `:8082/graphql/query`.
> * `app` is the **only** repo with migrations. All application tables — taxonomy, skill-path
>   sessions, the 23 jobsim run-state tables, the cms similarity/Studio tables — live in
>   **`public`**. The old per-service schemas are legacy and non-authoritative.
> * The in-process cms activates on `DIRECTUS_BASE_ADDR`; `DIRECTUS_WEBHOOK_SECRET` is now
>   **required** (the Directus webhook fails closed without it). The Directus cache lives in
>   `REDIS_CMS_CACHE_INDEX` (5). Judge0 is called directly via `JUDGE0_BASE_URL`.
>
> **The stack runbooks in this directory have NOT all been re-verified against the merged
> stack.** Where one names a per-service container, port, or profile for a folded service,
> read it as `backend`. See [`../services/backend.md`](../services/backend.md) for the
> current shape.

> **Corpus vs. extensions boundary:** rosetta is a read-only doc corpus + dev-env skills; ALL executable tooling that operates a spawned stack lives in rosetta-extensions — authored in `.agentspace/rosetta-extensions/`, tagged, and consumed per-stack via a pinned-tag clone.

## Available Operations

| Guide | Purpose | When to Use |
|-------|---------|-------------|
| **[Staging Bringup](./staging-bringup.md)** | **Full personal-staging onboarding (fresh VM → working Tailscale staging + live prod data + dev Clerk + daily sync). Includes the [colony v2-JWT vendoring recipe](./staging-bringup.md#bringup-quirks-consolidated-as-a-procedural-narrative) (Quirk #11), the [Atlas migrations gap fix](./staging-bringup.md#45-apply-pending-atlas-migrations) (§4.5), and the [known-schema-drifts table](./staging-bringup.md#105-known-schema-drifts-expected-on-staging) (§10.5).** | **New engineer (or AI agent) joining the team — start here** |
| [Staging Sync](./staging-sync.md) | Daily force-reset to `origin/main` + skip-worktree mechanics + recovery. **Note:** Atlas migrations are NOT in the daily run — see [Atlas migrations are NOT run by sync](./staging-sync.md#atlas-migrations-are-not-run-by-sync). | Understanding what the daily 06:00 UTC routine does, recovering clobbered WIP, or remembering to run Atlas periodically |
| [Staging Clerk](./staging-clerk.md) | Dev Clerk app `national-elk-17`, shared cross-engineer test login, the load-bearing `clerk-fetch-fix.js` monkey-patch, and the [v2 session-token anatomy](./staging-clerk.md#anatomy-of-a-v2-session-token) | Setting up auth on a new staging or debugging Clerk symptoms |
| [Platform Setup](./setup_guide.md) | Build the development environment | First time on a new machine (no prod-dump path) |
| [GitHub SSH Setup](./setup_github_guide.md) | Configure GitHub SSH access (single or dual personal+work account) for the `anthropos-work` org — keys, work-account-default for Docker, persistence. Skill `/setup-github`. | Before `make init` can clone the private repos |
| [Personal Staging from a Prod Dump](./staging_from_dump.md) | Restore a prod DB dump, rebind to a dev Clerk app, kill outbound email, apply colony/Clerk patches | Engineer-rebind reference (called from `staging-bringup.md`) |
| [Platform Run](./run_guide.md) | Start the platform locally | Daily development work |
| [Webhook Setup](./webhook_setup.md) | Configure Clerk webhooks for user sync | When you need user/org data locally |
| **[Dev-Stack Identity](./dev-identity.md)** | **Make a real Clerk operator exist to the platform: the `organizations`/`users`/`memberships` rows AND the per-membership Casbin grants (`g2`/`g3`/`p6`) that a dev bring-up never creates.** The main dev stack is deliberately never set-dressed (`dev-setdress.sh` hard-refuses `N=0`), and the two documented paths do not cover the ordinary case — the webhook needs a public tunnel, and `cmd/bootstrap-user` **mints a new Clerk user** so it cannot adopt the account you already sign in with. Failure shape: **authentication succeeds and authorization silently does not** — most sharply as a GraphQL `forbidden` on ONE field while its siblings answer, because the role is read from a `g2` grouping row and NOT from `memberships.role`. `rext dev-stack/dev-identity.sh`. | You signed in fine but every authorized surface is empty or 403s, or one GraphQL field returns `forbidden` |
| [Platform Update](./update_guide.md) | Sync code, deps, and schemas | After being away or before new features (superseded by `staging-sync.md` on staging hosts) |
| [Platform Repo](./platform_repo.md) | The `platform` orchestrator repo — the Makefile entry points, Docker Compose profiles, `repos.yml`, and how `make init`/`up`/`migrate` drive the whole local stack. | Understanding what `make` does, the compose profiles, or the repo layout |
| **[Platform Alignment](./platform-alignment.md)** | **How to detect that the platform moved, follow it, and fence it so the drift cannot silently recur.** The microservices→`app` consolidation is a **program**, not three accidents (v2.0 skiller → v5.0 skillpath → v7.0 jobsim → v8.0 cms → **v9.0 storage+messenger+customerio-sync, landed at `838d907`**), so the next occurrences are already named. Carries the six cheap **detection signals** (the load-bearing pair being `migrations:`/`schema:` in `repos.yml`); the four traps (**`migrations: false` entails nothing** — `sentinel` is `false` and alive with its own schema, so a fence keyed on the flag is wrong; the **declared vs actual** topology can disagree *by design*; the platform's own plan docs lag its code ~9 days; it ships **coordinated multi-repo** changes); why v2.8 was **latent** rather than broken (`migrate-demo.sh`'s **hand-maintained 4-tuple** creates the legacy schemas itself, bypassing `repos.yml` — with `skillpath` already the visible canary and M810 the time bomb); why nobody noticed (**pinning disables drift detection** — 11/11 clones report `behind: null` while the log says "provably fresh"); the **search discipline** (the NUL-byte trap is *folklore*: 3 false absences, 0 caused by NUL bytes — never swallow stderr, always run a positive control); and the **3-layer fence**. (v2.8/M257x) | The platform changed under you, a bring-up fails oddly, or a service is being folded into `app` |
| [Quick Ops](./quick_ops.md) | Common commands reference | When you need a quick command |
| [Demo Stacks](./rosetta_demo.md) | **Disposable, isolated demo stacks (`demo-N`) alongside the dev stack — Clerkenstein-wired, offset ports, killable cleanly, zero platform-repo change. Skills `/demo-up`, `/demo-down`; list via `/stack-list`.** (v1.1/M3) | Spinning up a throwaway demo environment to seed (M4) + show |
| [Seeding Spec](./seeding-spec.md) | Declaratively backfill a stack with structural data (blueprint + DAG + the 3-layer write isolation guard). Skill `/stack-seed` (`dev-N` or `demo-N`). (v1.1/M7) | Populating a demo/dev stack with an org + users + activity |
| [DB Access](./db-access.md) | Read-only prod DB access (the wired `postgres` MCP tool **or** Tailscale + `~/.pgpass`) + the public-vs-customer boundary. Skill `/db-query`. (v1.2/M9a) | Investigating data, sizing a surface, or telling public reference data from customer data |
| [Snapshot Spec](./snapshot-spec.md) | Capture a **public** reference surface once from a safe prod source, manifest-cache it in `.agentspace`, replay per-stack — tenant-data firewall + snapshot-fidelity. `stacksnap` CLI. (v1.2/M9a) | Filling a stack with the real public taxonomy/content library |
| [Secrets Spec](./secrets-spec.md) | Provision every repo's target `.env` (`dev-N`/`demo-N`) from one secret source (dir/zip, default `.agentspace/secrets`) — **values-blind** (no verb reads/echoes a value) — verified by the 6-repo/64-gene secret-coverage DNA + the two-tier keep-listed gate; the source-dir layout contract, alias-family vs distinct-similar rules, the waived class, and the `DIRECTUS_TOKEN` non-rearm safety. `stacksecrets` CLI, skill `/stack-secrets`. (v1.6/M27–M30) | Filling in a stack's `.env` secrets across all six repos + checking coverage |
| [Demo Recipes](./demo/README.md) | The end-to-end demo-env recipe family (up → **snapshot** → seed → use → down) + presets; the `/stack-snapshot` skill *set-dresses* a stack (`dev-N` or `demo-N`) with the real public taxonomy + Directus content (100% catalog). (v1.2/M11) | Running a believable, full-fidelity demo world |
| [Safety & Security](./safety.md) | **The authoritative, code-cited safety contract of the stack tooling: the read-side (the `AssertPlan`/`AssertCaptured` firewall + public predicates + bounded read-only capture) and the write-side (the 3-layer `CheckWrite`/`PreflightEnv`/`AssertClean` isolation guard + never-write shared Directus/prod-S3 + doubled n=0 guards + audit-proven zero pollution).** ⚠️ **Neither side is unqualified, and this row asserted both flatly until M257x close** — the read side carries the v2.5 `cmd/content-capture` exception (§3.8), and the write side is a claim about *the set of pointers the tooling knows to override*, proven incomplete on 2026-08-11 when a demo reached a production S3 bucket outside it. (v1.3/M15) | Understanding *why* snapshot/seed/db-query can't read customer data or pollute production — the safety landing page |
| [Bring-up Re-run Safety](./idempotency.md) | **The idempotency contract: re-running migrate / snapshot-replay / seed is either safe-and-idempotent (converges) or fails loudly with a guard — never silently doubles data.** (v1.3b/M17) | Knowing what happens when you run a bring-up step twice |
| [Observability](./observability.md) | **The tier this corpus documented nowhere until 2026-08-07** (`git grep -i grafana -- corpus/` returned 0 files). What the platform actually emits — **no metrics pipeline at all**, Sentry-protocol traces at a 15 % sample, and an error tier that is a **self-hosted GlitchTip**, not sentry.io — plus `ant-observability`'s live outside-in `product-monitoring/`, which asserts on **body content** because every tier here can return 200 for a failure. Names a **production read path** (asynq → prod ElastiCache) that no safety doc enumerates. | Answering "is production up, and how would I know?", or designing a probe that cannot be fooled by a 200 |
| [Bring-up Verification](./verification.md) | **TWO gates at the tail, with opposite failure semantics. (1) auto-verify — scoped, NON-FATAL: cheap-win `/api/health` + `casbin_rules > 0` asserts (the silent-403 catcher) then the full offset/scope-aware probe set; a verify bug never blocks a good stack (v1.3b/M18). (2) the Playthrough batch gate — LOUD: it drives every seeded hero's journey to completion, emits ONE consolidated red set, and makes the bring-up EXIT NON-ZERO on a non-empty one, while still leaving the stack UP (v2.8/M258). So "UP" means "UP, and every journey verified."** ⚠️ *This row said "NON-FATAL" full stop until the M258 close.* | Understanding what the post-bring-up checks verify, and why one of them can fail your bring-up |
| [Demo Frontend Tier](./demo/frontend-tier.md) | **`/demo-up` brings up the full UI: next-web-app + studio-desk + hiring (per-demo *cached* Docker images from **rext-owned** multi-stage Dockerfiles — build shape 3; the platform clone is a build CONTEXT only) + ant-academy natively (Clerk-free) — 12 GB VM prereq + non-fatal pre-flight, `--no-ui` escape, hard zero-platform-repo-edit line.** (v1.3b/M19; studio-desk moved to shape 3 at v2.8/M258 TIK-A) | Bringing up a demoable UI, sizing the VM, or understanding the per-demo frontend build |
| [Snapshot Cold-Start](./snapshot-cold-start.md) | **Filling the snapshot cache once per release on a fresh box (no `~/.pgpass`): the sanctioned DSN-export / dump-restore path, why the wired `postgres` MCP is *not* a capture source, and how it slots into the auto-set-dress bring-up.** (v1.3b/M20) | Getting the real public catalog onto a fresh machine before a demo |
| [Remote Demo over Tailscale](./demo/tailscale-serve.md) | **Make a demo reachable from another machine on a Tailscale tailnet. NB the two stack families have OPPOSITE defaults since v2.3 M220 (D-DESIGN-3, superseding v2.2's D-DESIGN-1): `/demo-up` is DEFAULT-ON — a bare `/demo-up N` auto-discovers the host and serves remotely; opt out with `--no-public-host`. `/dev-up` stays opt-in via `--public-host <magicdns>`.** The flow: one trusted `tailscale cert` HTTPS origin (Clerk needs a secure context) fronted by per-offset-port `tailscale serve`, `CORS_EXTRA_ORIGINS` + the ant-academy sha-pinned patch, and the fresh-Linux-VM host prereqs (Go + atlas + tailscale operator) the tooling pre-flights/auto-handles/fails-loud on. The FIRST live remote Linux-VM deploy — proven end-to-end, both hero vantages, trusted cert, cold reset-to-seed reproducible; F1–F12 finding set + safety framing.** (v2.2/M212–M215) | Serving a demo to a teammate on your tailnet, or standing one up on a remote Linux VM |
| [Content Stories — route map](./demo/content-stories-routes.md) | **The feasibility spike + per-product result-route map: for each content product × {player, manager}, the exact result route, classified by prove-by-render. Resolves the central unknown — `/sim/<slug>/result/<sessionId>` is a PERSISTED READ, not a live recompute, so a seeded result fan-out renders. Verdicts: Simulation GO; Skill-path GO **player-only** (the manager drill-down renders "Coming soon"); Interview GO behind a flag demo-patch; AI-labs OUT; Academy presence-only. Carries the manager-view MIRROR trap.** (v2.5/M231) | Deciding whether a content product can be proven by render, or debugging a blank manager scoreboard |
| [Session Clone Spec](./demo/session-clone-spec.md) | **The write side of Content stories: the `ContentStorySeeder` COPIES real production job-simulation sessions into a demo org — real LLM feedback, transcripts, submissions, interview reports — scrubbed best-effort of detectable PII, re-tenanted, and source-pinned for deterministic reseed. NOT provably clean: residual re-identification risk is real and ACCEPTED by the data-controller, the control being VPN/tailnet scope (`safety.md` §3.8).** (v2.5/M232) | Understanding where demo content actually comes from, or the PII posture of a demo |
| [Content Stories — manifest + honesty gate](./demo/content-stories-spec.md) | **The manifest half: `stackseed --content-export` projects `content-manifest.json` (the content analog of `cockpit-manifest.json`) that the cockpit's "Content stories" tab reads — per product, the played sessions with player + manager seat keys and result paths. Single-sourced from the same fixture the seeder seeds from, honesty-gated against a checked-in canonical, and fail-closed: a session that can't form a real link is DROPPED with a reason rather than rendering a fabricated CTA.** (v2.5/M233) | Adding a content product, or tracing why a cockpit CTA is missing |
| [Local Directus](./directus-local.md) | **The per-stack local Directus spec: the 11.6.1 bootstrap empirics, the structure-capture model (DDL + PRIMARY KEYs + sequences + serve rows that close the M10 collection-schema gap), the redefined `stacksnap` exit codes, the version-skew rule, the firewall structural-metadata admissibility carve-out, the executed container lifecycle (compose service + offset port + idempotent re-provision + verify probes), and the M23 data-plane cutover (`cms` re-pointed at the in-network instance + referential closure). A `--local-content` stack (demo default; dev opt-in) serves its OWN captured catalog — content-self-contained, asset plane on prod public links.** (v1.5 — structure M21 + lifecycle M22 + cutover M23) | Understanding how a stack serves its own captured content instead of reading live from prod |

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                  First Time Setup + Daily Run                    │
│  /dev-up  or  corpus/ops/setup_guide.md + run_guide.md           │
│  Install tools, clone repos, configure env, start the stack      │
│  (one skill — consolidates the former setup + start)             │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Periodic Updates                              │
│  /stack-update  or  corpus/ops/update_guide.md                   │
│  Pull latest code, install deps, run migrations                  │
└─────────────────────────────────────────────────────────────────┘
```

## Progress Tracking

All operations use Claude's **TodoWrite** tool for real-time progress tracking.

## Ops Reports

When `/dev-up` or `/stack-update` encounter errors or discover improvements, they create **ops reports** in `stack-dev/ops-reports/`:

```
stack-dev/ops-reports/
├── op_20250127_143022_setup_pgvector.md
├── op_20250127_151045_run_port_conflict.md
└── op_20250128_092311_update_migration_fail.md
```

### Report Format

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /dev-up | /stack-update
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Which operation phase]

## Issue Encountered
[Exact error message]

## Context
[What was being done]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in the guides]
```

### Integration Workflow

Ops reports are **not** automatically applied to documentation. Instead:

1. Skills create reports during execution
2. Run `/update-knowledge` to review and apply improvements
3. Human reviews and commits changes

This separates "live execution" from "corpus maintenance".

## Future Operations

This directory may grow to include:
- `deploy_guide.md` - Deployment procedures
- `debug_guide.md` - Debugging and diagnostics
