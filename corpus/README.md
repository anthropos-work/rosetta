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

### [Tools](./tools/)
Registry of development tools and toolchains.

*   [Toolchain Overview](./tools/toolchain_overview.md): Map of tools for setup, dev, and runtime.

### [Services](./services/)
Individual service documentation and developer maps.

*   **Core Backend Services**: 9 Go microservices (Backend, CMS, Sentinel, Skiller, etc.)
*   **Studio Services**:
    *   [Studio-Desk](./services/studio-desk.md): Content design service
    *   [Studio-Room](./services/studio-room.md): AI generation service

### [Ops](./ops/)
Operations guides for setting up, running, and updating the platform.

*   [Setup Guide](./ops/setup_guide.md): Step-by-step instructions for macOS and Linux.
*   [Run Guide](./ops/run_guide.md): Start the platform locally.
*   [Update Guide](./ops/update_guide.md): Sync code, dependencies, and schemas.

## Navigation

*   **Getting Started?** → [Setup Guide](./ops/setup_guide.md)
*   **Understanding the System?** → [Architecture Overview](./architecture/architecture_overview.md)
*   **Working on a Service?** → [Services Directory](./services/)
*   **Need Project Context?** → [Root README](../README.md)
*   **Updating This Corpus?** → Use `/anthropos-integrate` skill (see [Root README](../README.md#-updating-documentation-with-claude))
