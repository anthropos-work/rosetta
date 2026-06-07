# Project Rosetta

**The documentation corpus for the Anthropos platform.**

> **v1.0 "body double" + v1.1 "show floor" + v1.2 "set dressing" + v1.3 "stack party" — shipped.** Beyond documentation, Rosetta drives
> the executable stack tooling (in the private `rosetta-extensions` monorepo): an **alignment-testing
> framework** + **Clerkenstein** — a *measured* drop-in mock of Clerk that lets the platform run stacks Clerk-free
> with zero platform-code change (100% on Go · JS/FAPI · `@clerk/express`); disposable, **production-safely-seeded
> stacks**; and the **snapshot mechanism** that *set-dresses* them with the real **public** skills taxonomy +
> Directus content library at **100% data-DNA coverage** — captured read-only, customer data never copied. In v1.3
> **dev and demo stacks converged** — a unified first-available-N registry, dev as a full-fidelity peer (local
> Directus + auto-snapshot + light seed), and one generic `/dev-up` + `/stack-*` skill set. The
> tooling's two safety guarantees — **never reads private/customer data, never touches prod** — are stated
> authoritatively, code-cited, in [`corpus/ops/safety.md`](corpus/ops/safety.md). See also
> [`corpus/architecture/alignment_testing.md`](corpus/architecture/alignment_testing.md),
> [`corpus/services/clerkenstein.md`](corpus/services/clerkenstein.md), and
> [`corpus/ops/snapshot-spec.md`](corpus/ops/snapshot-spec.md).

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

1. **[Setup Guide](./corpus/ops/setup_guide.md)** - Build your local development environment
2. **[Run Guide](./corpus/ops/run_guide.md)** - Start the platform locally
3. **[Architecture Overview](./corpus/architecture/architecture_overview.md)** - Understand how the platform works
4. **[Service Taxonomy](./corpus/architecture/service_taxonomy.md)** - Learn the three-tier service model

### Using Claude Code?

Automate the setup process:
```bash
/dev-up              # First time / daily: build + start the dev environment (one skill — was setup + start)
/dev-up N            # When needed: an additional isolated dev-N stack, set-dressed by default
/demo-up N           # When needed: spin up an isolated, Clerkenstein-wired demo stack (e.g., demo-1, demo-2)
/stack-snapshot N    # Set-dress any stack: replay the real public taxonomy + Directus content (100% catalog; read-only capture)
/stack-seed N        # Then: backfill it with a believable data world (a preset or stack.seed.yaml)
/demo-down N         # When done: tear down a demo stack cleanly (--purge to drop its data); /dev-down N for dev
/stack-list          # Check: list the live dev + demo stacks, their offset ports, and health
```

These skills execute the guides step-by-step with verification, ask for confirmation before changes, and auto-improve documentation when issues are found. The generic `stack-*` ops (`/stack-list`, `/stack-seed`, `/stack-snapshot`, `/stack-update`) work on **any** stack — `dev-N` or `demo-N` — which the unified registry keeps from colliding on ports; each stack is isolated (`-p dev-N` / `-p demo-N`, offset ports) and the seeder is **production-safe** (it cannot write a shared/prod store) — start at the demo-env family index [`corpus/ops/demo/README.md`](./corpus/ops/demo/README.md).

## Documentation Structure

```
corpus/
├── architecture/          # System design and service relationships
│   ├── architecture_overview.md    # Start here for the big picture
│   ├── service_taxonomy.md         # Core, Studio, External tiers
│   ├── frontend_architecture.md    # Next.js monorepo details
│   ├── external_services.md        # Clerk, Directus, GraphQL
│   ├── dependency_map.md           # Service inter-dependencies
│   ├── shared_libraries.md         # colony, proto, ai, authn, taxonomy
│   └── alignment_testing.md        # The alignment test class + framework (rosetta-extensions/alignment/)
│
├── services/              # Individual service documentation
│   ├── backend.md, cms.md, sentinel.md, ...     # Core services
│   ├── graphql-wundergraph.md, next-web-app.md  # Gateway + main frontend
│   ├── studio-desk.md, studio-room.md           # Studio services
│   └── ant-academy.md                           # Internal learning portal (standalone)
│
├── ops/                   # Operations guides
│   ├── platform_repo.md   # The orchestrator repo (Make targets, profiles, compose)
│   ├── setup_guide.md     # Build local development environment
│   ├── run_guide.md       # Start services locally
│   ├── update_guide.md    # Sync code and dependencies
│   ├── safety.md          # The tooling safety contract (never reads customer data / touches prod)
│   ├── snapshot-spec.md   # Capture+replay the public reference library (read-side)
│   ├── seeding-spec.md    # Declarative stack seeding (write-side isolation boundary)
│   └── demo/              # Demo-environment family index + recipes
│
└── tools/                 # Development tools
    ├── toolchain_overview.md       # Required tools registry
    └── anthropos-labs.md           # Internal experiments hub
```

## The Workspace

Hands-on work happens in **stack workspaces** — git-ignored `stack-*/` directories, each spanning one full local stack: its platform service repos **plus its own clone of rosetta-extensions**. The family is `stack-dev` (dev), `stack-demo` (demo), `stack-dev-2` (secondary dev), and future `stack-stage` / `stack-tests`. Each stack-dir is used for:

- Cloning Anthropos platform repositories
- Building your local development environment
- Inspecting platform code for documentation

`rosetta-extensions` — the executable stack tooling — has **two clone roles**: an **authoring** copy at `.agentspace/rosetta-extensions/`, spawned on demand to read/build/test tooling, then committed and **tagged**; and **per-stack consumption** copies, `stack-<role>/rosetta-extensions @ <tag>`, where each stack consumes the tooling at a pinned tag.

All hands-on work happens in these stack-dirs. The documentation corpus stays clean.

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
/update-knowledge
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
| **Core Backend** | Backend, CMS, Sentinel, Skiller, Jobsimulation, Skillpath, Storage, Roadrunner | Go |
| **Studio & Standalone** | Studio-Desk (design tool), Studio-Room (AI pipeline, embedded in CMS), Ant Academy (internal learning portal) | TypeScript, Python, Next.js + Expo |
| **External** | Clerk (auth), Directus (CMS), GraphQL/Wundergraph (gateway) | SaaS / Docker |

See [Architecture Overview](./corpus/architecture/architecture_overview.md) for the full picture.

## Guidelines for Contributors

- **Verify against code**: Don't assume docs are correct. Check the actual source.
- **Update immediately**: Found a gap? Fix it now, not later.
- **Work in a `stack-*/` workspace**: Keep the corpus clean.
- **Tool in rosetta-extensions, not here**: any new stack-operating tool is built and tested in the `.agentspace/rosetta-extensions/` authoring copy and **tagged** — never added to the rosetta repo, never hand-written inside a stack dir. Stacks consume it at a pinned tag.
- **Never commit secrets**: No `.env` files in any repo.
- **Follow dual-level**: PM summary + engineer deep-dive for key concepts.

## Related Resources

| Resource | Location |
|----------|----------|
| Platform repos | `anthropos-work` GitHub org |
| Claude skills | [.claude/skills/](./.claude/skills/) |
| Setup progress | `stack-dev/setup_progress.md` (your local copy) |
