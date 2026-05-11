# Frontend Architecture (`next-web-app`)

This document outlines the architecture of the Anthropos frontend, which is organized as a **Monorepo** using **Turborepo** and **pnpm** workspaces.

## Monorepo Structure

The code is divided into `apps` (deployable applications) and `packages` (shared libraries).

### 1. Applications (`apps/`)

| App Name | Path | Port | Description |
| :--- | :--- | :--- | :--- |
| **Web** | `apps/web` | 3000 | Main user-facing app (dashboard, simulations, skill paths, onboarding, "Talk to Data" UI) |
| **Hiring** | `apps/hiring` | 3001 | Hiring / recruiting product (job ladders, candidate funnels) |
| **Integration** | `apps/integration` | 3002 | Embedded / third-party integration surface (iframe-able views, partner workflows) |
| **Mobile** | `apps/mobile` | — | Expo / React Native mobile app |
| **Maintenance** | `apps/maintenance` | — | Static maintenance / outage pages |

### 2. Core Packages (`packages/`)

| Package Name | Path | Responsibility |
| :--- | :--- | :--- |
| **UI Kit** | `packages/ui` | Shared UI components (Design System). Recent work: `SkillCard` mini-card, `SkillCardCallout`, onboarding `SkillsRefinement` cluster card. |
| **GraphQL** | `packages/graphql` | Shared GraphQL definitions, generated hooks, and types. Populated by `pnpm codegen` against the Cosmo Router supergraph. |
| **Core JS** | `packages/core-js` | Common utilities and helpers. |
| **TSConfig** | `configs/tsconfig` (workspace-scoped as `@anthropos/tsconfig`) | Shared TypeScript configs. |

## Data Layer & Communication

The frontend communicates with the backend services through two primary methods:

### 1. GraphQL
*   **Used For**: Content retrieval (via **CMS**), Simulation state, and aggregated data.
*   **Implementation**: `packages/graphql` contains the generated types and hooks.
*   **Code Generation**: Likely uses `graphql-codegen` to read schemas from `cms` and generate React hooks.

### 2. Connect RPC
*   **Used For**: Direct service calls (e.g., `Backend`, `Sentinel`).
*   **Implementation**: Services expose gRPC/Connect handlers. The frontend likely uses `@connectrpc/connect-web`.

## Key Technologies

* **Framework**: Next.js 14 (App Router), React
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
