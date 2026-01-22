# Architecture Documentation

This directory contains all documentation related to the Anthropos platform architecture.

## Files

*   **[architecture_overview.md](./architecture_overview.md)**: High-level system design, services, and communication patterns. Start here to understand the overall platform structure.

*   **[frontend_architecture.md](./frontend_architecture.md)**: Deep dive into the Next.js monorepo structure, key applications, shared packages, and data fetching patterns.

*   **[dependency_map.md](./dependency_map.md)**: Matrix of service inter-dependencies showing how different components interact with each other.

## Quick Start

1.  Begin with **[Architecture Overview](./architecture_overview.md)** to understand the high-level system design.
2.  Review **[Dependency Map](./dependency_map.md)** to see how services interact.
3.  Dive into **[Frontend Architecture](./frontend_architecture.md)** for UI-specific details.

## For Maintainers

When updating architecture documentation:
*   Keep the **architecture_overview.md** current with any new services or major architectural changes.
*   Update the **dependency_map.md** when service dependencies change.
*   Document frontend changes in **frontend_architecture.md** as the monorepo evolves.
