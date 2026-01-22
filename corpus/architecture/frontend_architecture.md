# Frontend Architecture (`next-web-app`)

This document outlines the architecture of the Anthropos frontend, which is organized as a **Monorepo** using **Turborepo** and **pnpm** workspaces.

## Monorepo Structure

The code is divided into `apps` (deployable applications) and `packages` (shared libraries).

### 1. Applications (`apps/`)

| App Name | Path | Description |
| :--- | :--- | :--- |
| **Hiring** | `apps/hiring` | The main "Hiring" product interface? |
| **Web** | `apps/web` | The core "Web" application (Dashboard, User Profile)? |
| **Mobile** | `apps/mobile` | Mobile-specific views or React Native app? |
| **Maintenance** | `apps/maintenance` | Maintenance mode pages? |

### 2. Core Packages (`packages/`)

| Package Name | Path | Responsibility |
| :--- | :--- | :--- |
| **UI Kit** | `packages/ui` | Shared UI components (Design System). |
| **GraphQL** | `packages/graphql` | Shared GraphQL definitions, hooks, and types (generated from backend). |
| **Core JS** | `packages/core-js` | Common utilities and helpers. |

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
*   **Framework**: Next.js (React)
*   **Build System**: Turborepo
*   **Package Manager**: pnpm
*   **Styling**: (Check: Tailwind? CSS Modules? `packages/ui` will reveal this)
*   **State Management**: (Check: Local state, Context, or external lib?)

## Development Workflow

### Running an App
To run a specific app (e.g., `hiring`):

```bash
cd next-web-app
pnpm dev --filter=hiring
```
Or run all apps:
```bash
pnpm dev
```
