# Architecture Documentation

This directory contains all documentation related to the Anthropos platform architecture.

## Files

*   **[architecture_overview.md](./architecture_overview.md)**: High-level system design, services, and communication patterns. Start here to understand the overall platform structure.

*   **[frontend_architecture.md](./frontend_architecture.md)**: Deep dive into the Next.js monorepo structure, key applications, shared packages, and data fetching patterns.

*   **[dependency_map.md](./dependency_map.md)**: Matrix of service inter-dependencies showing how different components interact with each other.

*   **[alignment_testing.md](./alignment_testing.md)**: The **alignment test class** (a third class beside unit and integration) and its reusable framework — how we measure, as a 0–100% score, how faithfully a *mirror* engine (e.g. Clerkenstein) reproduces a *source*. Three dimensions: **behavioral** (v1.0 — Clerkenstein vs Clerk), **structural data-DNA** (v1.1 — seeded-data conformance to the live schema), and **snapshot-fidelity** (v1.2 — source-vs-replay for captured public surfaces). Reference implementation: `rosetta-extensions/alignment/` + the `datadna` harness in `stack-seeding/dna/`.

*   **[service_taxonomy.md](./service_taxonomy.md)**: The three-tier service categorization — core backend (Tier 1), studio + internal apps (Tier 2), and external services (Tier 3) — with each service's role, ports, and integration pattern. The "which service is what" reference.

*   **[external_services.md](./external_services.md)**: The third-party integrations — Clerk (auth), Directus (the production headless CMS the platform reads content from), the WunderGraph Cosmo GraphQL gateway, the AI providers, LiveKit (voice), and AWS Chime (recording) — how each is configured and consumed.

*   **[shared_libraries.md](./shared_libraries.md)**: The five internal Go modules imported as private dependencies (colony, proto, ai, authn, taxonomy) — what each provides and where its responsibilities begin and end (e.g. cost tracking lives in `app`, not the `ai` library).

*   **[ai_architecture.md](./ai_architecture.md)**: The AI plane — models, EU-first provider routing (per-consumer wrappers, not the shared `ai` lib), the LiveKit voice engine, AWS Chime recording, and cost tracking (`app/internal/aiusage`).

*   **[security_compliance.md](./security_compliance.md)**: Security, data protection, EU compliance, and the multi-tenant isolation model (shared DB / shared schema with `organization_id`, the 3-layer isolation: DB + Sentinel authz + Clerk identity).

## Quick Start

1.  Begin with **[Architecture Overview](./architecture_overview.md)** to understand the high-level system design.
2.  Review **[Dependency Map](./dependency_map.md)** to see how services interact.
3.  Dive into **[Frontend Architecture](./frontend_architecture.md)** for UI-specific details.

## For Maintainers

When updating architecture documentation:
*   Keep the **architecture_overview.md** current with any new services or major architectural changes.
*   Update the **dependency_map.md** when service dependencies change.
*   Document frontend changes in **frontend_architecture.md** as the monorepo evolves.
