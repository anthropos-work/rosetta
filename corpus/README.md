# Project Rosetta: Anthropos Platform Documentation

Welcome to Project Rosetta. This documentation corpus aims to reverse engineer and document the architecture, product, and tech stack of the **Anthropos** platform.

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

## Structure

*   **[Architecture](./architecture/)**: Complete architecture documentation.
    *   **[Architecture Overview](./architecture/architecture_overview.md)**: High-level system design, services, and communication patterns.
    *   **[Frontend Architecture](./architecture/frontend_architecture.md)**: Deep dive into the Next.js monorepo.
    *   **[Dependency Map](./architecture/dependency_map.md)**: Matrix of service inter-dependencies.
*   **[Tools](./tools/)**: Registry of development tools.
    *   **[Toolchain Overview](./tools/toolchain_overview.md)**: Map of tools for setup, dev, and runtime.
*   **[Services](./services/)**: Individual service documentation and developer maps.
*   **[Setup](./setup/)**: Complete setup documentation for local development environment.
    *   **[Setup Guide](./setup/setup_guide.md)**: Step-by-step instructions for macOS and Linux.
    *   **[Setup Checklists](./setup/)**: OS-specific checklists for tracking setup progress.

## Status

**Current Phase**: Phase 4 - Frontend Architecture Deep Dive.

*   [x] Initial Service Enumeration
*   [x] High-Level Tech Stack Identification
*   [x] Establish Documentation Standards (Modus Operandi)
*   [x] Service Architecture & Developer Map
    *   *Focus: Role, Internal Architecture, Locality, Setup/Test procedures.*
    *   *Note: detailed API endpoints are excluded; we document how to find/consult them.*
*   [x] Frontend Architecture Deep Dive
    *   *Monorepo structure, Key Apps (Hiring), and Data Layer.*
