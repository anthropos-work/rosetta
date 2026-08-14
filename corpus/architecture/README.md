# Architecture Documentation

This directory contains all documentation related to the Anthropos platform architecture.

## Files

*   **[architecture_overview.md](./architecture_overview.md)**: High-level system design, services, and communication patterns. Start here to understand the overall platform structure.

*   **[frontend_architecture.md](./frontend_architecture.md)**: Deep dive into the Next.js monorepo structure, key applications, shared packages, and data fetching patterns.

*   **[dependency_map.md](./dependency_map.md)**: Matrix of service inter-dependencies showing how different components interact with each other.
*   **[org-repos.md](./org-repos.md)**: **The `anthropos-work` org repo register — all 93 repos, measured 2026-08-07, each with a home and an ADVISORY verdict.** The corpus's missing denominator: it documented the ~13 repos a stack clones and had never enumerated the rest, so a repo could be live, load-bearing and invisible at once. Settles the standing **`cms` M810** question from `infrastructure`'s own `services.tf`, and carries the newly-derived homes for `infrastructure`, `directus` (four extensions no stack gets), `judge0` (the one prod box that is not IaC), `metabase`, `AI-Labs`, the five `livekit-agent*` repos, `sim-qa`, `hyper-studio` and `anthropos-knowledge-base` (a **second, contradicting** org corpus). **Records verdicts; deletes nothing.**

*   **[alignment_testing.md](./alignment_testing.md)**: The **alignment test class** (a third class beside unit and integration) and its reusable framework — how we measure, as a 0–100% score, how faithfully a *mirror* engine (e.g. Clerkenstein) reproduces a *source*. Three dimensions: **behavioral** (v1.0 — Clerkenstein vs Clerk), **structural data-DNA** (v1.1 — seeded-data conformance to the live schema), and **snapshot-fidelity** (v1.2 — source-vs-replay for captured public surfaces). Reference implementation: `rosetta-extensions/alignment/` + the `datadna` harness in `stack-seeding/dna/`.

*   **[platform-migration-status.md](./platform-migration-status.md)**: **Where the microservice-into-`app` consolidation actually is** — one row per service the platform has ever had, **two states per row** (production vs a fresh local stack, because those genuinely differ), every claim cited to a sha or `file:line`, plus the **net-new** repos that appear in neither `repos.yml` nor the corpus. **Machine-fenced against the platform's own `repos.yml` in both directions** (`rosetta-extensions/stack-core/platform_alignment_guard.py`) so a service entering *or leaving* the clone set turns a guard RED instead of rotting a doc. The map that [`../ops/platform-alignment.md`](../ops/platform-alignment.md) produces and maintains; read it before trusting any per-service doc's claim about whether that service still runs.

*   **[service_taxonomy.md](./service_taxonomy.md)**: The three-tier service categorization — core backend (Tier 1), studio + internal apps (Tier 2), and external services (Tier 3) — with each service's role, ports, and integration pattern. The "which service is what" reference. Includes the **content-vs-runtime callout**: the cms domain is the *content layer* (it owns skill-path and simulation *content/definitions*), while the *runtime/session engines* reference that content by ID — skill-path engine ≠ skill-path content, jobsimulation engine ≠ simulation content. Since the monolith merge all of these are packages inside `app`, so the boundary is a package boundary rather than a network one.

*   **[external_services.md](./external_services.md)**: The third-party integrations — Clerk (auth), Directus (the production headless CMS the platform reads content from), the WunderGraph Cosmo GraphQL gateway (**deleted from the platform at `2adcf71`** — locally the frontends now hit `backend` directly at `:8082/graphql/query`), the AI providers, LiveKit (voice), and AWS Chime (recording) — how each is configured and consumed.

*   **[taxonomy-canon.md](./taxonomy-canon.md)**: **The taxonomy canon (taxonomy v2)** — the corpus's first doc anchor for the skill/job-role vocabulary's SOURCE, as opposed to its size or its runtime home. The canon is a **checked-in artifact inside `app`** (`app/taxonomy-canon/`), not a dataset repo: **3,562 skills / 706 canonical roles** measured 2026-08-14, consolidated down from **43,584 / 22,511**. Carries the **redirect map** (partial — 12,835 of 39,353 retired skills have a successor; **26,518 have none**), the retired-id 404 contract, and `taxonomyguard`'s closure of the taxonomy to **runtime minting** (get-or-create no longer mints, so a name that does not resolve simply does not exist). **It also settles the "60K / 18K" dispute** — the platform counted 43,584 before removing anything, so "60K skills" is now REFUTED rather than merely unverified, and the ≥42,790 public floor is vindicated by the 794-skill private remainder. Records the five net-new tables that sit OUTSIDE this project's snapshot capture surface.
*   **[shared_libraries.md](./shared_libraries.md)**: The five internal Go libraries — **and that is NOT the set a stack imports.** A service a stack builds imports **five private modules — `analytics-go`, `colony`, `proto`, `storage`, `taxonomy`** (`app/go.mod:14-18` @ `app` `ad9f3c498`, all direct). Two of those (`analytics-go`, `storage`) are not among this doc's five subjects, and two of its subjects (`ai`, `authn`) are not imported — **so the two fives share only a cardinality.** This row said *"only three are imported — colony, proto, taxonomy"* until M257x iter-133. **`ai` is no longer one of them**: `app` folded it in-tree at `1e457fa70` (2026-08-04) and now carries it at `app/internal/ai/`, so no checked-out `go.mod` a stack builds requires `github.com/anthropos-work/ai` — only the frozen `cms` / `jobsimulation` husks still do. **`authn` is not a dependency of any service**: no checked-out `go.mod`/`go.sum` requires `github.com/anthropos-work/authn` (the standalone repo is legacy); it ships **inside colony** and is imported as `github.com/anthropos-work/colony/authn`. The doc covers what each provides and where its responsibilities begin and end (e.g. cost tracking lives in `app`, not the `ai` library).

*   **[ai_architecture.md](./ai_architecture.md)**: The AI plane — models, provider routing (per-consumer wrappers, not the shared `ai` lib; EU-resident by default, and **not** an EU-first fallback ladder), the LiveKit voice engine, AWS Chime recording, and cost tracking (`app/internal/aiusage`).

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
