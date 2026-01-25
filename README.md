# Project Rosetta

**The documentation corpus for the Anthropos platform.**

## What Is This?

Project Rosetta is a **documentation repository** - not the Anthropos platform itself. It contains:

- Architecture guides explaining how the platform works
- Setup instructions for building a local development environment
- Service documentation for each microservice
- Tools for reverse-engineering and documenting the platform

**The actual platform code lives in separate repositories** under the `anthropos-work` GitHub organization.

## Who Is This For?

| Audience | What You'll Find |
|----------|------------------|
| **New Engineers** | Setup guides to get running, architecture docs to understand the system |
| **Product Managers** | High-level explanations of what each component does |
| **AI Agents** | Structured, parseable documentation for autonomous operation |

All documentation follows a **dual-level approach**: simplified explanations for PMs alongside technical deep-dives for engineers.

## Project Goals

### 1. The "Recreation Standard"
Can someone use this documentation to **recreate a full development environment from scratch** on a new machine? That's our acceptance criteria.

### 2. Autoconsistent & Discoverable
A new engineer or AI agent should land here, read this README, and know exactly where to find information and what to do next.

### 3. Living Documentation
This corpus evolves with the platform. When you discover gaps or better approaches, update the docs immediately.

## Quick Start

**Setting up for the first time?** Follow these in order:

1. **[Setup Guide](./corpus/ops/setup/setup_guide.md)** - Build your local development environment
2. **[Run Guide](./corpus/ops/run/run_guide.md)** - Start the platform locally
3. **[Architecture Overview](./corpus/architecture/architecture_overview.md)** - Understand how the platform works
4. **[Service Taxonomy](./corpus/architecture/service_taxonomy.md)** - Learn the three-tier service model

### Using Claude Code?

Automate the setup process:
```bash
/anthropos-setup
```

This skill executes the setup guide step-by-step with verification, asks for confirmation before changes, and auto-improves documentation when it finds issues.

## Documentation Structure

```
corpus/
├── architecture/          # System design and service relationships
│   ├── architecture_overview.md    # Start here for the big picture
│   ├── service_taxonomy.md         # Core, Studio, External tiers
│   ├── frontend_architecture.md    # Next.js monorepo details
│   ├── external_services.md        # Clerk, Directus, GraphQL
│   └── dependency_map.md           # Service inter-dependencies
│
├── services/              # Individual service documentation
│   ├── backend.md, cms.md, sentinel.md, ...  # Core services
│   └── studio-desk.md, studio-room.md        # Studio services
│
├── ops/                   # Operations guides
│   ├── setup/             # Environment setup
│   │   ├── setup_guide.md          # The main setup instructions
│   │   ├── setup_checklist_macos.md # macOS progress tracker
│   │   └── setup_checklist_linux.md # Linux progress tracker
│   └── run/               # Running the platform
│       ├── run_guide.md            # Start services locally
│       └── run_checklist.md        # Run progress tracker
│
└── tools/                 # Development tools
    ├── toolchain_overview.md       # Required tools registry
    └── anthropos-labs.md           # Internal experiments hub
```

## The Workspace

The `anthropos-dev/` directory is a **git-ignored scratchpad** for:

- Cloning Anthropos platform repositories
- Building your local development environment
- Inspecting platform code for documentation

All hands-on work happens here. The documentation corpus stays clean.

## Modus Operandi

Project Rosetta follows a strict **iterative approach**:

| Principle | Meaning |
|-----------|---------|
| **Iterative & Goal-Oriented** | One clear goal per iteration. Don't boil the ocean. |
| **Autoconsistent** | Everything you need is in this repo. No hidden dependencies. |
| **Recreation Standard** | Docs must enable full environment recreation from scratch. |
| **Dual-Level** | Every key concept has both PM-friendly and engineer-deep explanations. |

## Updating the Corpus

When you discover new platform elements or documentation gaps:

```bash
/anthropos-integrate
```

This skill:
- Analyzes new evidence (repos, features, setup issues)
- Creates an update plan
- Implements documentation changes following our standards

**Evidence types**: New repos (A), new features (B), new directories (C), setup feedback (D), missing docs (E).

## Platform Architecture (Summary)

Anthropos uses a **three-tier microservices architecture**:

| Tier | Services | Technology |
|------|----------|------------|
| **Core Backend** | Backend, CMS, Sentinel, Skiller, Jobsimulation, Skillpath, Storage, Chronos, Intelligence | Go |
| **Studio** | Studio-Desk (design tool), Studio-Room (AI pipeline) | TypeScript, Python |
| **External** | Clerk (auth), Directus (CMS), GraphQL/Wundergraph (gateway) | SaaS / Docker |

See [Architecture Overview](./corpus/architecture/architecture_overview.md) for the full picture.

## Guidelines for Contributors

- **Verify against code**: Don't assume docs are correct. Check the actual source.
- **Update immediately**: Found a gap? Fix it now, not later.
- **Work in `anthropos-dev/`**: Keep the corpus clean.
- **Never commit secrets**: No `.env` files in any repo.
- **Follow dual-level**: PM summary + engineer deep-dive for key concepts.

## Related Resources

| Resource | Location |
|----------|----------|
| Platform repos | `anthropos-work` GitHub org |
| Claude skills | [.claude/skills/](./.claude/skills/) |
| Setup progress | `anthropos-dev/setup_progress.md` (your local copy) |
