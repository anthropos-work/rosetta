# Corpus Directory

This directory contains all Project Rosetta documentation. For the full project overview, modus operandi, and guidelines, see the [root README](../README.md).

## Directory Structure

### [Architecture](./architecture/)
Complete architecture documentation for the Anthropos platform.

*   [Architecture Overview](./architecture/architecture_overview.md): High-level system design, three-tier service model, and communication patterns.
*   [Service Taxonomy](./architecture/service_taxonomy.md): Service categorization (Core, Studio, External tiers).
*   [Frontend Architecture](./architecture/frontend_architecture.md): Deep dive into the Next.js monorepo.
*   [External Services](./architecture/external_services.md): Third-party integrations (Clerk, Directus, GraphQL).
*   [Dependency Map](./architecture/dependency_map.md): Matrix of service inter-dependencies.
*   [Shared Libraries](./architecture/shared_libraries.md): The five internal Go libraries (colony, proto, ai, authn, taxonomy).

### [Tools](./tools/)
Registry of development tools and toolchains.

*   [Toolchain Overview](./tools/toolchain_overview.md): Map of tools for setup, dev, and runtime.

### [Services](./services/)
Individual service documentation and developer maps.

*   **Core Backend Services**: 9 Go microservices (Backend, CMS, Sentinel, Skiller, etc.)
*   **Gateway & Frontend**:
    *   [GraphQL Gateway](./services/graphql-wundergraph.md): WunderGraph Cosmo Router (Apollo Federation v2)
    *   [Next Web App](./services/next-web-app.md): Main customer-facing frontend (Workforce + Hiring)
*   **Integrations**:
    *   [Clerk Integration](./services/clerk-integration.md): Identity / authentication / organizations — what it's used for, dependent repos, SDKs
*   **Studio Services & Standalone Internal Apps**:
    *   [Studio-Desk](./services/studio-desk.md): Content design tool
    *   [Studio-Room](./services/studio-room.md): AI generation pipeline (embedded in CMS)
    *   [Ant Academy](./services/ant-academy.md): Internal learning portal for `@anthropos.work` employees (Next.js 16 + Expo, Vercel)

### [Ops](./ops/)
Operations guides for setting up, running, and updating the platform.

*   **Demo environments (disposable, Clerk-free, seeded stacks — v1.1 "show floor"):**
    *   [Demo Environments — family index](./ops/demo/README.md): **Start here.** The 4-step flow (`/demo-up` → `/stack-seed` → use → `/demo-down`) + the index of guides, recipes, and presets.
    *   [Rosetta Demo](./ops/rosetta_demo.md): The lifecycle mechanism — bring-up, port-offset, Clerkenstein injection, per-stack isolation, teardown.
    *   [Seeding Spec](./ops/seeding-spec.md): The `stack.seed.yaml` blueprint, the dependency-DAG, the **production-isolation boundary**, the data-DNA.
    *   Recipes: [enterprise onboarding](./ops/demo/recipe-enterprise-onboarding.md) · [skill progression](./ops/demo/recipe-skill-progression.md) · [browser login](./ops/demo/recipe-browser-login.md).
*   **Personal staging (full onboarding for new engineers + AI agents):**
    *   [Staging Bringup](./ops/staging-bringup.md): The spine doc — fresh VM → Tailscale-attached staging with live prod data + dev Clerk login + daily sync. **Start here if you're new.**
    *   [Staging Sync](./ops/staging-sync.md): Daily force-reset to `origin/main`, skip-worktree mechanics, recovery from clobbered WIP.
    *   [Staging Clerk](./ops/staging-clerk.md): Shared dev Clerk app, cross-engineer test login, the load-bearing `clerk-fetch-fix.js` monkey-patch.
*   **General platform operations:**
    *   [Platform Repo Reference](./ops/platform_repo.md): The orchestrator repo — Make targets, profiles, docker-compose, repos.yml.
    *   [Setup Guide](./ops/setup_guide.md): Step-by-step instructions for macOS and Linux.
    *   [Run Guide](./ops/run_guide.md): Start the platform locally.
    *   [Update Guide](./ops/update_guide.md): Sync code, dependencies, and schemas.
    *   [Staging from Dump](./ops/staging_from_dump.md): Engineer-rebind reference (creates Clerk users + remaps DB).

## Navigation

*   **Getting Started?** → [Setup Guide](./ops/setup_guide.md)
*   **Understanding the System?** → [Architecture Overview](./architecture/architecture_overview.md)
*   **Working on a Service?** → [Services Directory](./services/)
*   **Need Project Context?** → [Root README](../README.md)
*   **Updating This Corpus?** → Use `/update-knowledge` skill (see [Root README](../README.md#updating-the-corpus))
