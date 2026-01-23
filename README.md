# Project Rosetta: Anthropos Platform Documentation

Welcome to Project Rosetta. This documentation corpus serves as the cognitive kernel for the **Anthropos** platform. It is designed to support three recursive activities by targeting the source code (conceptually located at `@anthropos`):

1.  **Documentation Repository**: A comprehensive guide for new developers to understand the platform's architecture ("how it works").
2.  **Environment Setup**: A manual for humans and AI agents to build, configure, and verify a local development environment.
3.  **Recursive Inspection**: A tool for "reverse engineering" the platform—inspecting existing code to document, build, and recursively improve this corpus itself.

## Modus Operandi

Project Rosetta follows a strict **iterative approach** to reverse engineering and documenting the Anthropos platform.

1.  **Iterative & Goal-Oriented**: We move in distinct iterations. Each iteration must have a clear, achievable goal (e.g., "Map the high-level architecture", "Document the Backend service"). We do not attempt to boil the ocean in one go.
2.  **Autoconsistent & Discoverable**: This documentation constitutes a self-contained corpus. A new agent or engineer should be able to land here, read this README, and understand exactly where to find information and how to proceed with the next task.
3.  **The "Recreation" Standard**: The ultimate acceptance criteria for this documentation is whether an autonomous agent or a new engineer can use it to **recreate a full development environment** from scratch on a new machine.
4.  **Dual-Level Documentation**: Key documents should provide two layers of depth:
    *   **High-Level / For PMs**: Simplified, conceptual explanations.
    *   **Engineering / Deep Dive**: Technical specifics, configuration details, and inter-service dependencies.

## Guidelines for Agents & Engineers

If you are an AI Agent or Engineer picking up this project:

*   **Start Here**: This README is your source of truth for the project status and structure.
*   **Check `task.md`**: The `task.md` file (in the `brain` directory or mirrored here if applicable) tracks the granular progress.
*   **Respect the Structure**: Maintain the separation between "High-Level" and "Deep Dive" sections.
*   **Verify**: Do not assume the docs are 100% correct. Always verify against the actual code (`docker-compose.yaml`, `go.mod`, source code) when detailed accuracy is required.
*   **Update**: If you discover something new or fix a gap, update these documents immediately. The docs must evolve with the understanding.

## 📂 Documentation Structure

All documentation is located in the **[corpus](./corpus/)** directory:

*   **[Architecture](./corpus/architecture/)**: Complete architecture documentation.
    *   [Architecture Overview](./corpus/architecture/architecture_overview.md): High-level system design, three-tier service model, and communication patterns.
    *   [Service Taxonomy](./corpus/architecture/service_taxonomy.md): Service categorization (Core, Studio, External tiers).
    *   [Frontend Architecture](./corpus/architecture/frontend_architecture.md): Deep dive into the Next.js monorepo.
    *   [External Services](./corpus/architecture/external_services.md): Third-party integrations (Clerk, Directus, GraphQL).
    *   [Dependency Map](./corpus/architecture/dependency_map.md): Matrix of service inter-dependencies.
*   **[Tools](./corpus/tools/)**: Registry of development tools.
    *   [Toolchain Overview](./corpus/tools/toolchain_overview.md): Map of tools for setup, dev, and runtime.
*   **[Services](./corpus/services/)**: Individual service documentation and developer maps.
    *   Core Backend Services: 9 Go microservices (Backend, CMS, Sentinel, Skiller, etc.)
    *   Studio Services: [Studio-Desk](./corpus/services/studio-desk.md) (content design) and [Studio-Room](./corpus/services/studio-room.md) (AI generation)
*   **[Setup](./corpus/setup/)**: Complete setup documentation for local development environment.
    *   [Setup Guide](./corpus/setup/setup_guide.md): Step-by-step instructions for macOS and Linux.
    *   [Setup Checklists](./corpus/setup/): OS-specific checklists for tracking setup progress.

## 🚀 Quick Start

**New to the project?** Start with these documents in order:

1.  **[Setup Guide](./corpus/setup/setup_guide.md)**: Recreate the development environment from scratch.
2.  **[Architecture Overview](./corpus/architecture/architecture_overview.md)**: Understand the system design and service interactions.
3.  **[Service Taxonomy](./corpus/architecture/service_taxonomy.md)**: Learn how services are organized into Core, Studio, and External tiers.

### 🤖 Automated Setup with Claude

If you're using Claude Code, you can use the `/anthropos-setup` skill to automate the environment setup:

```bash
/anthropos-setup
```

This skill:
*   Executes the setup guide step-by-step with verification at each stage
*   Requests confirmation before installing tools or making system changes
*   Tracks progress using the setup checklist
*   **Auto-improves documentation** when it discovers issues or better approaches
*   Follows the STEP RUN guidelines (verify before/after, request confirmation, document improvements)

See [`.claude/skills/anthropos-setup/`](./.claude/skills/anthropos-setup/) for skill details.

## 🏗️ Development Scratchpad

The `anthropos-dev/` directory is an ignored workspace for hands-on activities:

*   **Environment Setup**: Follow the Setup Guide, clone repositories, and build the local development environment.
*   **Recursive Inspection**: Inspect the actual Anthropos platform code to reverse-engineer, document, and recursively improve this corpus.

This keeps the documentation repository clean while providing a dedicated space for working with the actual codebase.

## Status

**Current Phase**: Phase 5 - Studio Services & External Integrations.

*   [x] Initial Service Enumeration
*   [x] High-Level Tech Stack Identification
*   [x] Establish Documentation Standards (Modus Operandi)
*   [x] Service Architecture & Developer Map
    *   *Focus: Role, Internal Architecture, Locality, Setup/Test procedures.*
    *   *Note: detailed API endpoints are excluded; we document how to find/consult them.*
*   [x] Frontend Architecture Deep Dive
    *   *Monorepo structure, Key Apps (Hiring), and Data Layer.*
*   [x] Studio Services & External Integrations
    *   *Documented studio-desk (content design) and studio-room (AI generation)*
    *   *Documented external services: Clerk (auth), Directus (CMS), GraphQL (gateway)*
    *   *Created three-tier service taxonomy (Core, Studio, External)*
