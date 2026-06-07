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
| `/setup-platform` | Build the dev environment from scratch | `corpus/ops/setup_guide.md` |
| `/setup-github` | Configure GitHub SSH access for the org | `corpus/ops/setup_github_guide.md` |
| `/start-platform` | Start the platform locally | `corpus/ops/run_guide.md` |
| `/update-platform` | Sync code, deps, and schemas | `corpus/ops/update_guide.md` |
| `/update-knowledge` | Document new evidence across the corpus | N/A (meta-skill) |
| `/test-platform` | Verify a running platform (probes, repo suites, census) | `.claude/skills/test-platform/SKILL.md` |
| `/db-query` | Query the prod Postgres read-only (investigate data, size/inspect surfaces) | `corpus/ops/db-access.md` |
| `/demo-up` | Spin up an isolated demo stack (Clerkenstein-wired, offset ports) | `corpus/ops/rosetta_demo.md` |
| `/demo-down` | Tear down a demo stack cleanly | `corpus/ops/rosetta_demo.md` |
| `/demo-status` | List running demo stacks and their details | `corpus/ops/rosetta_demo.md` |
| `/demo-seed` | Seed a demo stack with realistic structural data (presets or `stack.seed.yaml`) | `corpus/ops/seeding-spec.md` |
| `/demo-snapshot` | Set-dress a stack — replay the real public taxonomy + Directus content into it (or capture/status) | `corpus/ops/snapshot-spec.md` |
| `/align-dna` | Build/update an Alignment DNA for a mirror engine + capture goldens | `corpus/architecture/alignment_testing.md` |
| `/align-run` | Measure a mirror's alignment score vs a source engine | `corpus/architecture/alignment_testing.md` |

### Using the Setup Skill

For building the Anthropos development environment:
```bash
/setup-platform
```

This skill executes `corpus/ops/setup_guide.md` with:
- Verification before/after each step
- User confirmation before destructive operations
- Progress tracking in `stack-dev/setup_progress.md`
- Auto-improvement of documentation when issues are found

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

### Using the Run Skill

For starting the platform locally after setup:
```bash
/start-platform
```

This skill executes `corpus/ops/run_guide.md` with:
- Service health verification
- Proper startup sequence (infra → backend → frontend → studio-desk)
- Port conflict detection and resolution
- Progress tracking via TodoWrite

### Using the Update Skill

For syncing code, dependencies, and database schemas:
```bash
/update-platform
```

This skill executes `corpus/ops/update_guide.md` with:
- Daily/weekly/full update scenarios
- Git conflict handling
- Migration application
- Docker image rebuilding

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
| `stack-demo/` | disposable **demo** stacks (Clerkenstein-wired, offset ports) |
| `stack-dev-2/` | a secondary dev stack |
| `stack-stage/`, `stack-tests/`, … | future stacks, same pattern |

All hands-on platform work happens inside a `stack-*/` dir; the documentation
corpus stays clean. (Setup/run/update progress is tracked in
`stack-dev/setup_progress.md`.)

### `rosetta-extensions` — where stack tooling lives

`rosetta-extensions` (private: `anthropos-work/rosetta-extensions`) is the
executable-tooling monorepo that **operates** stacks — sections: `clerkenstein`
(the Clerk mock), `demo-stack`, `dev-stack`, `stack-injection`, `stack-core`,
`stack-seeding`, `alignment`. `rosetta` documents *how the platform works*;
`rosetta-extensions` is *the tooling that spins up, injects, and seeds copies of
it*. It has **two clone roles**:

- **Authoring copy → `.agentspace/rosetta-extensions/`** — the single working clone
  you spawn on demand to **read / build / test** the tooling, then commit and
  **tag**. New tools are developed here.
- **Per-stack consumption copies → `stack-<role>/rosetta-extensions @ <tag>`** —
  each stack consumes the tooling at a **pinned tag** (reproducible). The
  `/demo-*` and `/align-*` skills drive a stack's own clone.

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
- Backend (`app`): Main API gateway and user management
- CMS: Content management, Directus proxy, **and embedded studio-room AI generation pipeline** (`cms/studio/` is the `anthropos-studio-room` repo, cloned via `cd cms && make init-studio` and gitignored — a submodule-style pattern, not a real `.gitmodules` entry)
- Sentinel: Authorization only (Casbin RBAC/ABAC) — authentication is Clerk + the `authn` middleware in each service, not Sentinel
- Jobsimulation: Job environments and task simulation (voice, chat, code, documents)
- Skiller: Skill management, assessment, taxonomy (60K skills, 18K roles), and vector embeddings (RAG)
- Skillpath: Skill progression paths
- Storage: File/blob storage management
- Roadrunner: Code execution proxy to Judge0 sandbox
- Gotenberg: Office-doc → PDF conversion (third-party image; consumed by `app/internal/converter/gotenberg.go`)

Available in other profiles but NOT started by default:
- Messenger (`messenger` profile): Email notifications via Brevo (Sendinblue)
- CustomerIO Sync (`customerio-sync` profile): Background data sync to Customer.io. Unique build pattern — built directly from GitHub URL, not cloned locally.

Production-only / deployed-only (not in local docker-compose):
- db-backup: Scheduled PostgreSQL backups (every 6h) to S3, Azure, Hetzner

Archived (removed from local orchestration; repo dirs may still exist on disk):
- Chronos (was: scheduling & time-based events) — removed via platform commit `045857c`
- Intelligence (was: background data sync between backend and skiller schemas) — removed via platform commit `fdfa189`

**Shared Libraries** (imported as private Go modules — **not** cloned by `make init`/`repos.yml`; pulled at Docker build via `GH_PAT`/`GOPRIVATE`). See `corpus/architecture/shared_libraries.md`.
- colony: Platform framework (logging+Sentry, DB, Redis, GraphQL/RPC servers, middleware, pub/sub via Watermill); **also contains `authn`**
- proto: Protobuf definitions (RPC contracts) + hand-written domain types
- ai: AI provider wrapper behind one `ai.AI` interface (OpenAI, Azure, Anthropic, Bedrock, Mistral). NOTE: cost tracking lives in `app/internal/aiusage`, and EU-first routing lives in each consumer's wrapper — **not** in this library
- authn: Clerk JWT authentication — now shipped **inside colony** as `colony/authn` (standalone `authn` repo is legacy)
- taxonomy: **node-id library** (`NodeID` type + ID generation/validation) — **not** a dataset; the 60K-skill/18K-role data lives in skiller

**Studio Services & Standalone Internal Apps (Tier 2)**: Content creation tools + internal-only apps
- Studio-Desk (TypeScript/Vite/Express): Design tool for creating simulation blueprints (repo: `studio-desk`)
- Studio-Room (Python/Asyncio): AI-powered content generation pipeline (repo: `anthropos-studio-room`). **Embedded inside the cms container** as `cms/studio/` via `cd cms && make init-studio`; no longer a standalone deployment.
- Ant Academy (Next.js 16 + Expo): Internal learning portal for `@anthropos.work` employees (repo: `ant-academy`). **Vercel-deployed standalone — not in docker-compose.** Cloned by `make init` (in `repos.yml`); runs natively via `cd ant-academy/code && npm run dev` (port 3077). No platform backend dependencies at runtime — only Clerk. See `corpus/services/ant-academy.md`.

**External Services (Tier 3)**: Third-party integrations
- Clerk: User authentication (SaaS)
- Directus: Headless CMS (self-hosted)
- GraphQL/Cosmo Router: Apollo Federation v2 gateway (5 subgraphs: app, skiller, jobsimulation, cms, skillpath)
- AI Providers: OpenAI, Anthropic, Mistral (EU-first routing)
- LiveKit: Real-time voice engine for simulations
- AWS Chime: Video/audio recording

**Frontend Applications**: Next.js 15 monorepo on Vercel (`next-web-app`; see `corpus/services/next-web-app.md`)
- Next Web App: Main user-facing application
- Hiring App: Recruiting and hiring workflows
- Mobile App: Expo/React Native mobile experience

### Communication Patterns

- **Core Services ↔ Core Services**: Connect-RPC + Redis Streams (via Watermill) for async messaging
- **Frontend/Studio → Backend**: GraphQL via Cosmo Router (Apollo Federation v2, 5 subgraphs)
- **External Integrations**: Clerk SDK + JWT middleware (authn library), Directus proxied via CMS service
- **AI**: EU-first routing implemented in each consumer's `internal/ai` wrapper, **not** the shared `ai` library (EU Azure default → US Azure via PostHog flag `flag_use_azure_us` → direct-OpenAI on HTTP 429; Anthropic always Bedrock `eu-west-1`). Cost tracking in `app/internal/aiusage`
- **Multi-tenancy**: Shared DB, shared schema with `organization_id` on every table; 3-layer isolation (DB, Sentinel auth, Clerk identity)

### Environment Configuration

**Platform services** share a **single centralized `.env` file** in the `platform` repository. Docker-based services do not need their own `.env` files.

**Studio-Desk** requires its own `.env` file (`studio-desk/.env`) with Clerk and OpenAI credentials copied from `platform/.env`.

**Ant Academy** requires its own `.env` file at `ant-academy/code/.env` (not the repo root — the React app reads only from `code/.env`). Reuse `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `CLERK_SECRET_KEY` from `platform/.env`, and add `OPENAI_API_KEY` / `ANTHROPIC_API_KEY` for the `/api/ai/chat` route. Set `REQUIRE_ORGANIZATION_MEMBERSHIP=0` for solo local dev to skip the org-membership gate.

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
- `corpus/ops/demo/README.md`: **The demo-env family index** — the up→snapshot→seed→use→down flow + recipes + presets
- `corpus/ops/rosetta_demo.md`: The demo-stack lifecycle (bring-up, port-offset, Clerkenstein injection, teardown)
- `corpus/ops/seeding-spec.md`: The `stack.seed.yaml` blueprint + the **production-isolation boundary** (write-side) + the data-DNA (now **100%**, nothing waived)
- `corpus/ops/db-access.md`: **Production DB read access** (read-side) — the `/db-query` skill + the public-vs-customer boundary + the snapshot read foundation (v1.2)
- `corpus/ops/snapshot-spec.md`: The **`stack-snapshot` extension** (v1.2 M9a/M9b/M10) — capture the public taxonomy + Directus content once from a safe prod source, manifest-cache it in `.agentspace`, replay per-stack (`/demo-snapshot`); the tenant-data firewall + the `stacksnap` CLI + the snapshot-fidelity gate
- `corpus/ops/demo/recipe-snapshot-world.md`: The **set-dressing recipe** — capture→replay the real public library so a demo world's catalog + content templates are real

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
cp .env.example .env   # fill Clerk + AI keys (see corpus/ops/setup_guide.md)
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
6. `.claude/skills/setup-platform/SKILL.md`: Automated setup skill
7. `.claude/skills/setup-github/SKILL.md`: GitHub SSH setup skill
8. `.claude/skills/start-platform/SKILL.md`: Automated run skill
9. `.claude/skills/update-platform/SKILL.md`: Automated update skill
10. `.claude/skills/update-knowledge/SKILL.md`: Corpus documentation skill

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
2. Follow `corpus/ops/setup_guide.md` to build environment (or use `/setup-platform`)
3. Follow `corpus/ops/run_guide.md` to start the platform locally (or use `/start-platform`)
4. Read `corpus/architecture/architecture_overview.md` for system understanding
5. Consult `corpus/services/` for specific service details
