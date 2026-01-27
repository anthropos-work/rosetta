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
| `/ant-setup` | Build the dev environment from scratch | `corpus/ops/setup_guide.md` |
| `/ant-setup-github` | Configure GitHub SSH access for the org | `corpus/ops/setup_github_guide.md` |
| `/ant-run` | Start the platform locally | `corpus/ops/run_guide.md` |
| `/ant-update` | Sync code, deps, and schemas | `corpus/ops/update_guide.md` |
| `/ant-document` | Document new evidence across the corpus | N/A (meta-skill) |

### Using the Setup Skill

For building the Anthropos development environment:
```bash
/ant-setup
```

This skill executes `corpus/ops/setup_guide.md` with:
- Verification before/after each step
- User confirmation before destructive operations
- Progress tracking in `anthropos-dev/setup_progress.md`
- Auto-improvement of documentation when issues are found

### Using the GitHub Setup Skill

For configuring GitHub SSH access to contribute to `anthropos-work` repositories:
```bash
/ant-setup-github
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
/ant-run
```

This skill executes `corpus/ops/run_guide.md` with:
- Service health verification
- Proper startup sequence (infra → backend → frontend → studio-desk)
- Port conflict detection and resolution
- Progress tracking via TodoWrite

### Using the Update Skill

For syncing code, dependencies, and database schemas:
```bash
/ant-update
```

This skill executes `corpus/ops/update_guide.md` with:
- Daily/weekly/full update scenarios
- Git conflict handling
- Migration application
- Docker image rebuilding

### Using the Document Skill

For documenting new platform evidence across the Rosetta corpus:
```bash
/ant-document [evidence description]
```

This skill analyzes new evidence and performs a **corpus-wide sweep** to update all relevant documentation:
- Inspects the evidence (repos, features, tools, feedback)
- Checks ALL corpus sections that may need updates
- Updates Claude skills when automation is affected
- Ensures new content is discoverable from parent docs

Example invocations:
- `/ant-document the new studio-analytics repo`
- `/ant-document issues found in setup_progress.md`
- `/ant-document the Redis caching layer isn't documented`

### Working in the Scratchpad

The `anthropos-dev/` directory is a git-ignored workspace for:
- Cloning Anthropos platform repositories
- Building the local development environment
- Inspecting platform code without affecting this documentation repo

All hands-on work with the Anthropos platform should happen in `anthropos-dev/`.

## Architecture Overview

### Three-Tier Service Model

**Core Backend Services (Tier 1)**: 9 Go microservices
- Backend (`app`): Main API gateway and user management
- CMS: Content management and Directus proxy
- Sentinel: Authorization and authentication
- Jobsimulation: Job environments and task simulation
- Skiller: Skill management and assessment
- Skillpath: Skill progression paths
- Storage: File/blob storage management
- Chronos: Scheduling and time-based events
- Intelligence: AI/ML integration layer

**Studio Services (Tier 2)**: Content creation tools
- Studio-Desk (TypeScript/Vite/Express): Design tool for creating simulation blueprints (repo: `studio-desk`)
- Studio-Room (Python/Asyncio): AI-powered content generation pipeline (repo: `anthropos-studio-room`)

**External Services (Tier 3)**: Third-party integrations
- Clerk: User authentication (SaaS)
- Directus: Headless CMS (self-hosted)
- GraphQL/Wundergraph: API gateway and GraphQL federation

**Frontend Applications**: Next.js monorepo
- Next Web App: Main user-facing application
- Hiring App: Recruiting and hiring workflows
- Mobile App: Expo/React Native mobile experience

### Communication Patterns

- **Core Services ↔ Core Services**: HTTP/RPC (Connect RPC) + Redis Streams for async messaging
- **Frontend/Studio → Backend**: GraphQL via Wundergraph (unified gateway)
- **External Integrations**: Clerk SDK + middleware, Directus proxied via CMS service

### Environment Configuration

**Platform services** share a **single centralized `.env` file** in the `platform` repository. Docker-based services do not need their own `.env` files.

**Studio-Desk** requires its own `.env` file (`studio-desk/.env`) with Clerk and OpenAI credentials copied from `platform/.env`.

Critical environment variables:
- `CLERK_SECRET_KEY` & `CLERK_PUBLISHABLE_KEY` (Auth)
- `OPENAI_API_KEY` & `ANTHROPIC_API_KEY` (AI services)
- `DIRECTUS_PUBLIC_BASE_ADDR` (Content)

### Docker Compose Project Naming

Services run with `-p ant-rosetta` flag for isolation:
- Creates containers: `ant-rosetta-postgresql-1`, `ant-rosetta-backend-1`, etc.
- Creates networks: `ant-rosetta_app-network`
- Creates volumes: `ant-rosetta_postgres_data`

This prevents conflicts with other Anthropos environments.

## Key Documentation Locations

### Setup & Onboarding
- `corpus/ops/setup_guide.md`: Complete environment setup instructions
- `corpus/ops/setup_github_guide.md`: GitHub SSH access configuration

### Running the Platform
- `corpus/ops/run_guide.md`: Start the platform locally after setup

### Updating the Platform
- `corpus/ops/update_guide.md`: Sync code, dependencies, and database schemas
- `corpus/ops/update_checklist.md`: Progress tracker for updates

### Architecture Documentation
- `corpus/architecture/architecture_overview.md`: High-level system design
- `corpus/architecture/service_taxonomy.md`: Three-tier service categorization
- `corpus/architecture/frontend_architecture.md`: Next.js monorepo deep dive
- `corpus/architecture/external_services.md`: Clerk, Directus, GraphQL integration patterns
- `corpus/architecture/dependency_map.md`: Service inter-dependency matrix

### Service Documentation
- `corpus/services/`: Individual service documentation following TEMPLATE.md pattern
- Each service doc includes: Role, Architecture, Interface Discovery, Local Development, Testing

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

**Note**: CMS service requires studio-room to be symlinked as `cms/studio` for Docker builds.

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
4. `corpus/ops/update_guide.md` / `update_checklist.md`: Update instructions
5. `.claude/skills/ant-setup/SKILL.md`: Automated setup skill
6. `.claude/skills/ant-setup-github/SKILL.md`: GitHub SSH setup skill
7. `.claude/skills/ant-run/SKILL.md`: Automated run skill
8. `.claude/skills/ant-update/SKILL.md`: Automated update skill
9. `.claude/skills/ant-document/SKILL.md`: Corpus documentation skill

**When to use ant-document**: After discovering new platform elements, receiving setup feedback, or finding documentation gaps. The skill performs a corpus-wide sweep to ensure all relevant sections are updated.

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
├── anthropos-dev/             # Git-ignored scratchpad for platform work
├── .claude/skills/            # Claude Code automation skills
└── README.md                  # Project overview and status
```

## Critical Rules

- **Work in `anthropos-dev/` only** when dealing with actual platform code
- **Never commit `.env` files** to any repository
- **Update documentation immediately** when discovering gaps or better approaches
- **Verify against actual code** - don't assume documentation is 100% correct
- **Maintain dual-level depth** - both PM-friendly and engineer-friendly explanations

## Quick Start for New Developers

1. Read `README.md` for project overview
2. Follow `corpus/ops/setup_guide.md` to build environment (or use `/ant-setup`)
3. Follow `corpus/ops/run_guide.md` to start the platform locally (or use `/ant-run`)
4. Read `corpus/architecture/architecture_overview.md` for system understanding
5. Consult `corpus/services/` for specific service details
