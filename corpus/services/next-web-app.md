# Next Web App (`next-web-app`)

> Service-level / ops map for the main customer-facing frontend. For the
> **monorepo deep dive** (apps, packages, codegen, UX work) see
> [Frontend Architecture](../architecture/frontend_architecture.md). This page is the
> "what is it / how is it built & run" view.

## Role & Responsibility

* **Primary Goal**: The main user-facing frontend — a pnpm + Turborepo monorepo of
  Next.js apps that consume the federated GraphQL gateway and authenticate with Clerk.
* **Key Functions**:
  * Ship two **distinct sold products** from one monorepo: **Workforce** (`apps/web`) and **Hiring** (`apps/hiring`).
  * Talk to the backend **only** through the GraphQL gateway (`:5050/graphql`) — no direct microservice calls. In particular it has **no direct Directus dependency**: content reaches it through the gateway → the CMS subgraph → Directus, so the M23 content cutover (re-pointing CMS's `DIRECTUS_BASE_ADDR` at the per-stack Directus) is transparent to next-web — no `DIRECTUS_BASE_ADDR` env on the frontend. (The demo override does strip the inherited prod `DIRECTUS_TOKEN` from next-web too, defence-in-depth, even though it never reads Directus directly.) Browser images still load from the prod asset plane (`DIRECTUS_PUBLIC_BASE_ADDR=content.anthropos.work`), which is why the baked next/image host whitelist needs no rebuild.
  * Enforce auth at the edge via Clerk middleware (all routes protected by default, explicit public allowlist).
  * Deploy per-app to **Vercel**; only `apps/web` is also containerizable for local Docker.

## Architecture & Code Map

* **Codebase**: `next-web-app` (local) — repo `git@github.com:anthropos-work/next-web-app`
* **Language / runtime**: **TypeScript**, **Next.js 15.5.16** (App Router, Turbopack), **React 19.2**, **Node ≥ 24**, **pnpm 10.30.3**
* **Build system**: Turborepo 2.9.x; `repos.yml` type `node-pnpm`
* **Data layer**: `graphql-request` + **TanStack React Query** (⚠️ **not** Apollo Client) + `@graphql-codegen` client-preset
* **Database**: none (org scoping comes from Clerk session claims; data lives in backend services)

### Apps (`apps/`)

| App | Package | Port | Product / purpose | Dockerized? |
|-----|---------|------|-------------------|-------------|
| **Workforce** | `@anthropos/web-app` | 3000 | Primary product (`app.anthropos.work`): skill paths, AI simulations, org skill management, dashboard, **AI-readiness** (the member 3-step onboarding `components/ai-readiness/` + the manager dashboard `app/.../ai-readiness/`; gates DIFFER by surface — corrected v2.3 M219: the **member** funnel is gated on PostHog `flag_ai_readiness` **and** the org `ai_readiness` setting; the **manager dashboard** is gated on the GraphQL `aiReadinessEnabled` + the `isEnterprise` nav, and does **NOT** read the PostHog flag. Conflating them is the wrong-vantage error M219 spent a section correcting. A demo bakes no PostHog, so the flag resolves `undefined` forever and the member surface needs the `next-web-aireadiness-flag-gate` demo-patch — see [`ai-readiness.md`](ai-readiness.md)) | ✅ (the only one) |
| **Hiring** | `@anthropos/hiring-app` | 3001 | Distinct product (`hiring.anthropos.work`): job ladders, candidate funnels | ❌ Vercel-only |
| **Integration** | `@anthropos/integration` | 3002 | Public-website embed (WordPress via proxy rewrites, SEO/Prerender) | ❌ Vercel-only |
| **Maintenance** | `@anthropos/maintenance-app` | — | Downtime/outage placeholder UI | ❌ |
| **Mobile** | `@anthropos/mobile` | 3031 (Expo) | Expo / React Native PoC (**paused**); **excluded** from the pnpm workspace, uses `EXPO_PUBLIC_*` | ❌ |

### Shared packages (`packages/` + `configs/`)

| Package | Responsibility |
|---------|----------------|
| `@anthropos/graphql` | All data fetching: queries, React Query hooks by domain, server fetchers, **codegen output** (`src/__generated__`), `codegen.ts` |
| `@anthropos/ui` | Shared Ant Design 6 component library |
| `@anthropos/core-js` | Shared constants, types, utils (e.g. `security/printToken`) |
| `configs/*` | Workspace-scoped eslint / prettier / tailwind / tsconfig / **i18n** (8 locales: de, en, es, fr, it, ja, nl, pt) |

## Interface Discovery

* **GraphQL**: single federated endpoint `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (default `http://localhost:5050/graphql`); Clerk bearer token injected via React Query `defaultOptions.queries.meta.getToken`.
* **Auth edge**: `apps/web/src/middleware.ts` — `clerkMiddleware` protects every non-public route; public allowlist includes `/login`, `/sign-up`, `/checkout`, `/free-trial`, `/monitoring`, `/print`, `/api/bunny/thumbnail`. `/print` routes are HMAC-gated (`PRINT_ROUTE_SECRET`) for Puppeteer PDF generation.
* **Observability proxies**: `/logpoint/*` → PostHog (EU); `/monitoring` tunnels Sentry/Better Stack events.

## Dependencies

* **Downstream**: GraphQL gateway (`:5050/graphql`), backend `app` API (`:8082`), Clerk, PostHog (EU), Sentry/Better Stack, Stripe (billing), Bunny CDN (thumbnails + Chime recordings, token-signed), Metabase (embedded analytics), Azure OpenAI/OpenAI (server AI routes).
* **Upstream**: end users / browsers; WordPress (embeds `apps/integration`); Vercel (prod hosting); platform compose service `next-web-app` (containerized Workforce variant).

## Local Development

### Native (recommended — hot reload)

```bash
cd next-web-app
nvm use 24                       # Node 24+ is required (engines.node ">=24")
pnpm install
cp apps/web/.env.example apps/web/.env   # fill Clerk + GraphQL endpoint; never commit
pnpm dev:web                     # Workforce on :3000  (next dev --turbopack)
pnpm dev:hiring                  # Hiring on :3001
pnpm dev:integration             # Integration on :3002
pnpm codegen                     # regenerate GraphQL types (needs the gateway at :5050)
pnpm check                       # tsc --noEmit + eslint --fix across the workspace
pnpm storybook                   # Storybook on :6006
```

> Older Node fails with `WARN Unsupported engine`, and pnpm refuses to wipe
> `node_modules` in non-TTY shells (`ERR_PNPM_ABORTED_REMOVE_MODULES_DIR_NO_TTY`).

### Containerized (Workforce only)

```bash
cd platform
make up PROFILE=frontend         # builds Dockerfile.dev (web app only), serves :3000
# or: make up-frontend           # next-web-app together with the graphql backend stack
```

`Dockerfile.dev` (Node 24 alpine) builds **only** `@anthropos/web-app`
(`pnpm turbo build --filter=@anthropos/web-app`). Hiring / integration / maintenance /
mobile are **not** containerized — they ship via Vercel only. `NEXT_PUBLIC_*` are baked
at **build time**; on a remote VM set `PUBLIC_HOST` in `platform/.env` so the client
bundle resolves the right hostname.

## Environment Variables (high-signal subset)

| Variable | Default | Description |
|----------|---------|-------------|
| `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `CLERK_SECRET_KEY` | — | Clerk auth (client / server) |
| `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` | `http://localhost:5050/graphql` | Runtime GraphQL gateway endpoint (baked at build) |
| `NEXT_PUBLIC_BACKEND_API_URL` | `http://localhost:8082` | Backend (`app`) API base URL |
| `GRAPHQL_SCHEMA_FOR_GEN` | `http://localhost:5050/graphql` | Schema endpoint used by `graphql-codegen` |
| `NEXT_PUBLIC_HOSTING_URL` / `PUBLIC_HOST` | `http://localhost:3000` / `localhost` | Public hosting URL; `PUBLIC_HOST` parameterizes baked URLs in compose |
| `NEXT_PUBLIC_POSTHOG_KEY` / `_HOST` · `NEXT_PUBLIC_SENTRY_DSN` / `SENTRY_AUTH_TOKEN` | — | Analytics + error tracking (PostHog EU, Sentry/Better Stack) |
| `STRIPE_*` / `NEXT_PUBLIC_STRIPE_*` | — | Billing/checkout |
| `BUNNY_*` · `METABASE_*` · `PRINT_ROUTE_SECRET` | — | CDN, embedded analytics, signed PDF/print routes |
| `EXPO_PUBLIC_*` | — | Mobile (Expo) variants of the public vars |

## Testing

```bash
pnpm test            # turbo test → jest in apps/web and apps/hiring
                     # (integration & maintenance have no test script)
# E2E: Playwright suite under e2e/ (needs E2E_TEST_EMAIL / E2E_TEST_PASSWORD for Clerk login)
```

## Notable Gotchas

* **Next.js 15 / React 19**, not 14 — there's an in-repo `knowledge/next15-adoption-plan.md`; the repo's own `CLAUDE.md` still says 14 (stale).
* **Only one Dockerfile** (`Dockerfile.dev`) exists at the repo root — the repo `CLAUDE.md` "two Dockerfiles" note is stale.
* **8 locales** on disk (Portuguese added) though some docs say 7.
* Frontend data layer is `graphql-request` + React Query, **not Apollo Client**.
* `npm`/`yarn` are blocked (`engines` + `please-use-pnpm`); conventional commits enforced (commitlint + Husky + cocogitto).

## Related Documentation

* [Frontend Architecture](../architecture/frontend_architecture.md) — monorepo deep dive, packages, codegen, recent UX work
* [GraphQL Gateway](./graphql-wundergraph.md) — the federated endpoint this app consumes
* [External Services → Clerk](../architecture/external_services.md#clerk-authentication-service)
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Dependency Map](../architecture/dependency_map.md)
