# Frontend Architecture (`next-web-app`)

This document outlines the architecture of the Anthropos **main customer-facing frontend** (`next-web-app`), organized as a **Monorepo** using **Turborepo** and **pnpm** workspaces.

> For the service-level summary (role, ports, profiles, build/run, env vars) see the [Next Web App service doc](../services/next-web-app.md). This page is the monorepo deep dive.

> **Note**: There are **two other frontend products** that live outside this monorepo and have their own architecture:
> - **[Studio-Desk](../services/studio-desk.md)** — Vite + Express, simulation design tool
> - **[Ant Academy](../services/ant-academy.md)** — Next.js 16 + Expo, internal learning portal for `@anthropos.work` employees
>
> Both are pulled by `make init` from `repos.yml` and run natively (not in docker-compose by default). The rest of this document is about `next-web-app` specifically.

## Monorepo Structure

The code is divided into `apps` (deployable applications) and `packages` (shared libraries).

### 1. Applications (`apps/`)

| App Name | Path | Port | Description |
| :--- | :--- | :--- | :--- |
| **Web** | `apps/web` | 3000 | Main user-facing app (dashboard, simulations, skill paths, onboarding, "Talk to Data" UI) |
| **Hiring** | `apps/hiring` | 3001 | Hiring / recruiting product (job ladders, candidate funnels) |
| **Integration** | `apps/integration` | 3002 | Embedded / third-party integration surface (iframe-able views, partner workflows) |
| **Mobile** | `apps/mobile` | 3031 (Expo) | Expo / React Native mobile app — **paused PoC**; **excluded** from the pnpm workspace (`!apps/mobile` in `pnpm-workspace.yaml`), so not installed/built by `pnpm install` / `turbo dev` by default |
| **Maintenance** | `apps/maintenance` | — | Static maintenance / outage pages |

### 2. Core Packages (`packages/`)

| Package Name | Path | Responsibility |
| :--- | :--- | :--- |
| **UI Kit** | `packages/ui` | Shared UI components (Design System). Recent work: `SkillCard` mini-card, `SkillCardCallout`, onboarding `SkillsRefinement` cluster card. |
| **GraphQL** | `packages/graphql` | Shared GraphQL definitions, generated hooks, and types. Populated by `pnpm codegen` against the Cosmo Router supergraph. |
| **Core JS** | `packages/core-js` | Common utilities and helpers. |
| **TSConfig** | `configs/tsconfig` (workspace-scoped as `@anthropos/tsconfig`) | Shared TypeScript configs. |
| **i18n** | `configs/i18n` (workspace-scoped as `@anthropos/i18n`) | Translation messages for 8 locales (de, en, es, fr, it, ja, nl, pt) + next-intl config; consumed as a workspace dependency by all web apps. |

## Data Layer & Communication

The frontend communicates with the backend **exclusively through the federated GraphQL gateway** (Cosmo Router at `:5050/graphql`, env `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`) using `graphql-request` + TanStack React Query, with Clerk bearer tokens injected per-request via `useGraphql` (`Authorization: Bearer <token>`). There are **no** direct Connect/gRPC calls from the frontend.

### 1. GraphQL
*   **Used For**: Content retrieval (via **CMS**), Simulation state, and aggregated data.
*   **Implementation**: `packages/graphql` contains the generated types and hooks.
*   **Code Generation**: Uses `graphql-codegen` to read the supergraph schema from the federated GraphQL endpoint (Cosmo Router, `:5050/graphql`) and emits typed GraphQL documents into `packages/graphql/src/__generated__` via the client-preset (documents sourced from `src/query/**`). React Query hooks are hand-authored on top of these typed documents.

## Key Technologies

* **Framework**: Next.js 15 (App Router, Turbopack), React 19
* **Build System**: Turborepo 2.x
* **Package Manager**: pnpm 10.x (`packageManager: "pnpm@10.30.3"`)
* **Node**: **v24+ required** (`engines.node: ">=24.0.0"` in `package.json`)
* **Lint / format / typecheck**: ESLint + Prettier + TypeScript 5.9, orchestrated by Turbo (`turbo check`, `check:lint`, `check:types`, `check:deprecations`)
* **Commits**: Conventional commits enforced via `@commitlint/cli` + `husky`
* **CHANGELOG**: Generated from conventional commits (cog)

## Development Workflow

### Setup

Make sure Node 24+ is active before installing:

```bash
cd next-web-app
nvm use 24             # or: nvm install 24 && nvm use 24
pnpm install
```

Older Node will fail with `WARN Unsupported engine` and pnpm will refuse to wipe `node_modules` in non-TTY shells (`ERR_PNPM_ABORTED_REMOVE_MODULES_DIR_NO_TTY`).

### Running an app

Top-level `turbo dev` runs all apps. To run a specific app:

```bash
pnpm dev --filter=@anthropos/web-app     # apps/web
pnpm dev --filter=@anthropos/hiring-app  # apps/hiring
pnpm dev --filter=@anthropos/integration # apps/integration
```

### Codegen

GraphQL types are regenerated from the upstream subgraph schemas:

```bash
pnpm codegen           # one-shot
pnpm codegen:watch     # incremental
```

### Cleaning

Several granular targets for cache management; the common one is:

```bash
pnpm clean:quick       # drops node_modules, .next, .turbo, dist
```

## Recent UX Work (May 2026)

The frontend is high-velocity (300+ commits / month). Recent themes from the CHANGELOG:

* **Onboarding redesign**: SkillsRefinement cluster cards, viewport-relative grid heights, swiper layouts, OnboardCard size tuning
* **SkillCard component**: Reusable info-callout mini-card used across import/refinement flows
* **AI Academy**: Iframe header simplification
* **Talk to Data**: SSE-streaming Q&A UI (counterpart to backend's `ask`/`askengine`)
* **Hiring talk-to-data**: Variant scoped to hiring workflows
