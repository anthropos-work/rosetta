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

*   [Security & Compliance](./architecture/security_compliance.md): Data protection, EU compliance, multi-tenancy isolation.
*   [AI Architecture](./architecture/ai_architecture.md): Models, provider routing, voice engine, recording, cost tracking.
*   [Alignment Testing](./architecture/alignment_testing.md): Measuring how faithfully a mirror engine reproduces its source as a 0–100% score.

### [Tools](./tools/)
Registry of development tools and toolchains.

*   **[Tools Index](./tools/README.md): the enumerated list.**
*   [Toolchain Overview](./tools/toolchain_overview.md): Map of tools for setup, dev, and runtime.
*   [Anthropos Labs](./tools/anthropos-labs.md): The internal experiments hub (`anthropos-work/experiments`) — PoCs and prototypes, not part of the platform.

### [Services](./services/)
Individual service documentation and developer maps.

*   **[Services Index](./services/README.md): every service doc, enumerated and grouped — start here rather than guessing a filename.** Covers the core backend tier, the gateway + frontends, the cross-cutting subsystems (AI-readiness, hiring, Clerk, Clerkenstein), and the archived/merged redirects (`skiller`, `chronos`, `intelligence`).
*   **Core Backend Services**: 8 Go microservices (Backend, CMS, Sentinel, etc. — skiller was merged into Backend, July 2026)
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

*   **Disposable stacks (Clerk-free, snapshot-set-dressed, seeded — dev *and* demo, converged in v1.3 "stack party"):**
    *   [Demo Environments — family index](./ops/demo/README.md): **Start here.** The flow (`/dev-up` or `/demo-up` → `/stack-snapshot` → `/stack-seed` → use → `/dev-down` or `/demo-down`) + the index of guides, recipes, and presets.
    *   [Rosetta Demo](./ops/rosetta_demo.md): The lifecycle mechanism — bring-up, the unified first-available-N registry (v1.3/M12), port-offset, Clerkenstein injection, per-stack isolation, teardown.
    *   [Seeding Spec](./ops/seeding-spec.md): The `stack.seed.yaml` blueprint, the dependency-DAG, the **production-isolation boundary**, the data-DNA, the shipped presets (incl. the `dev-min` dev auto-seed).
    *   [Snapshot Spec](./ops/snapshot-spec.md): Capture a **public** reference surface once from a safe prod source, manifest-cache it, replay per-stack — tenant-data firewall + snapshot-fidelity (v1.2). Dev is a full-fidelity peer (v1.3/M13).
    *   [Secrets Spec](./ops/secrets-spec.md): Provision every repo's target `.env` (`dev-N`/`demo-N`) from one secret source (dir/zip) — **values-blind** — verified by the 6-repo/56-gene secret-coverage DNA + the keep-listed gate; the `DIRECTUS_TOKEN` non-rearm safety (v1.6/M27–M30). Driven by `/stack-secrets`.
    *   [DB Access](./ops/db-access.md): Read-only prod DB access + the public-vs-customer boundary (v1.2/M9a).
    *   [Safety & Security](./ops/safety.md): The code-cited safety contract — never reads private data, never touches prod (v1.3/M15).
    *   **Content stories (v2.5 "the playbill")** — real prod sessions, cloned + scrubbed, so a demo shows real played content:
        *   [Content Stories — route map](./ops/demo/content-stories-routes.md): per content product × vantage, the exact result route, classified by prove-by-render (M231).
        *   [Session Clone Spec](./ops/demo/session-clone-spec.md): the write side — the `ContentStorySeeder`, the scrub, and the **accepted residual re-identification risk** (M232).
        *   [Content Stories — manifest + honesty gate](./ops/demo/content-stories-spec.md): the `content-manifest.json` projection the cockpit tab reads — fail-closed, never a fabricated CTA (M233).
    *   Recipes: [snapshot world](./ops/demo/recipe-snapshot-world.md) · [enterprise onboarding](./ops/demo/recipe-enterprise-onboarding.md) · [skill progression](./ops/demo/recipe-skill-progression.md) · [browser login](./ops/demo/recipe-browser-login.md).
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
