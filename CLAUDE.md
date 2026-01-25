# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

**Project Rosetta** is the documentation corpus for the Anthropos platform. It serves three purposes:
1. **Documentation Repository**: Comprehensive architecture guides for developers
2. **Environment Setup**: Manual for humans and AI agents to build local development environments
3. **Recursive Inspection**: Tool for reverse-engineering and documenting the platform itself

This is NOT the Anthropos platform source code - it's the documentation about it. The actual platform code lives in separate repositories under the `anthropos-work` GitHub organization.

## Development Commands

### Using the Automated Setup Skill

For building the Anthropos development environment:
```bash
/anthropos-setup
```

This skill executes `corpus/setup/setup_guide.md` with:
- Verification before/after each step
- User confirmation before destructive operations
- Progress tracking in `anthropos-dev/setup_progress.md`
- Auto-improvement of documentation when issues are found

### Using the Corpus Integration Skill

For integrating new platform evidence into the Rosetta documentation:
```bash
/anthropos-integrate
```

This skill analyzes new evidence about the Anthropos platform and updates the corpus:
- **New repos/tools**: Clone, inspect, create service documentation
- **New features**: Analyze code, update affected service docs
- **Setup feedback**: Parse issues from setup_progress.md, improve setup guide
- **Missing aspects**: Document undocumented platform elements

Evidence types:
- `A` - New repository, project, or tool
- `B` - New feature/code in existing service
- `C` - New directory in anthropos-dev/
- `D` - Setup feedback (setup_progress.md with issues)
- `E` - Missing platform aspect

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

All services share a **single centralized `.env` file** in the `platform` repository. Individual service repositories do not need their own `.env` files when running via Docker.

Critical environment variables:
- `CLERK_SECRET_KEY` & `CLERK_PUBLISHABLE_KEY` (Auth)
- `OPENAI_API_KEY` & `ANTHROPIC_API_KEY` (AI services)
- `DIRECTUS_PUBLIC_BASE_ADDR` (Content)

### Docker Compose Project Naming

Services run with `-p anthropos-rosetta` flag for isolation:
- Creates containers: `anthropos-rosetta-postgresql-1`, `anthropos-rosetta-backend-1`, etc.
- Creates networks: `anthropos-rosetta_app-network`
- Creates volumes: `anthropos-rosetta_postgres_data`

This prevents conflicts with other Anthropos environments.

## Key Documentation Locations

### Setup & Onboarding
- `corpus/setup/setup_guide.md`: Complete environment setup instructions
- `corpus/setup/setup_checklist_macos.md`: macOS-specific setup checklist
- `corpus/setup/setup_checklist_linux.md`: Linux-specific setup checklist

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
npm run dev    # Runs on localhost:3100
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

When updating `corpus/setup/setup_guide.md`, follow these principles:
1. **Verify Before Install**: Include commands to check if tools exist
2. **Verify After Install**: Include commands to confirm successful installation
3. **Request Confirmation**: Document where user approval is needed
4. **Document Improvements**: Add troubleshooting entries when issues are discovered

### Interconnected Documentation

These files must be maintained together:
1. `corpus/setup/setup_guide.md`: Detailed setup instructions
2. `corpus/setup/setup_checklist_macos.md` / `setup_checklist_linux.md`: OS-specific checklists
3. `.claude/skills/anthropos-setup/SKILL.md`: Automated setup skill
4. `.claude/skills/anthropos-integrate/SKILL.md`: Corpus integration skill

**When to update checklists**: Only when setup structure changes (steps added/removed/reordered), not for clarifications or verification command additions.

**When to use anthropos-integrate**: After discovering new platform elements, receiving setup feedback, or finding documentation gaps. The skill ensures corpus updates follow the modus operandi.

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
│   ├── setup/                 # Environment setup guides
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
2. Follow `corpus/setup/setup_guide.md` to build environment (or use `/anthropos-setup`)
3. Read `corpus/architecture/architecture_overview.md` for system understanding
4. Consult `corpus/services/` for specific service details
