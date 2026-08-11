# Next Web App (`next-web-app`)

> Service-level / ops map for the main customer-facing frontend. For the
> **monorepo deep dive** (apps, packages, codegen, UX work) see
> [Frontend Architecture](../architecture/frontend_architecture.md). This page is the
> "what is it / how is it built & run" view.

## Role & Responsibility

* **Primary Goal**: The main user-facing frontend — a pnpm + Turborepo monorepo of
  Next.js apps that consume the platform's single GraphQL endpoint (locally **`backend` directly**, since
  platform `2adcf71` deleted the Cosmo router) and authenticate with Clerk.
* **Key Functions**:
  * Ship two **distinct sold products** from one monorepo: **Workforce** (`apps/web`) and **Hiring** (`apps/hiring`). The hiring **org-type** (`is_hiring`) re-skins `apps/web` for a recruiting buyer — **but the recruiter candidate-comparison read-model is NOT reachable in `apps/web` for a *genuine* hiring org** (corrected M257x iter-102; this bullet's conjunction implied it was). A user whose Clerk memberships are **all** `isHiring` orgs is ejected out of `apps/web` into `apps/hiring` — see the ⚠️ note under *Apps* below and [`hiring.md`](hiring.md), which is authoritative for the org-type + read-model.
  * Talk to the backend **only** through the GraphQL endpoint — **`:8082/graphql/query` on `backend` directly** since platform `2adcf71` deleted the router from local dev (it was `:5050/graphql` on the Cosmo Router) — no direct microservice calls. In particular it has **no direct Directus dependency**: content reaches it through **`backend`'s GraphQL endpoint → the cms domain inside `app` → Directus** (it was gateway → CMS subgraph → Directus before the folds; the supergraph has been **one** subgraph since `915da06`, and the local router is gone since `2adcf71`), so the M23 content cutover (re-pointing CMS's `DIRECTUS_BASE_ADDR` at the per-stack Directus) is transparent to next-web — no `DIRECTUS_BASE_ADDR` env on the frontend. (The demo override does strip the inherited prod `DIRECTUS_TOKEN` from next-web too, defence-in-depth, even though it never reads Directus directly.) Browser images still load from the prod asset plane (`DIRECTUS_PUBLIC_BASE_ADDR=content.anthropos.work`), which is why the baked next/image host whitelist needs no rebuild.
  * Enforce auth at the edge via Clerk middleware (all routes protected by default, explicit public allowlist).
  * Deploy per-app to **Vercel**; `apps/web` is the only app with an in-repo Dockerfile and the only frontend in platform compose (a **demo** additionally containerizes `apps/hiring` from a rext-side Dockerfile — note under *Apps*).

## Architecture & Code Map

* **Codebase**: `next-web-app` (local) — repo `git@github.com:anthropos-work/next-web-app`
* **Language / runtime**: **TypeScript**, **Next.js 16.2.12** (App Router, Turbopack), **React 19.2.7**, **Node ≥ 24**, **pnpm 10.30.3**. All four apps (`apps/{web,hiring,integration,maintenance}/package.json`) declare `"next": "~16.2.12"` — **a tilde range, not a caret** — and the lockfile resolves **`16.2.12`**, measured @ `8297c684`. (`apps/mobile` declares no `next`.) **This line read `"^16.2.7"` until M257x iter-108**; `^16.2.7` is the `@next/*` sibling range, not this one. The repo carries `UPGRADE-IMPACT-next16.md`
* **Build system**: Turborepo 2.9.x; `repos.yml` type `node-pnpm`
* **Data layer**: `graphql-request` + **TanStack React Query** (⚠️ **not** Apollo Client) + `@graphql-codegen` client-preset
* **Database**: none (org scoping comes from Clerk session claims; data lives in backend services)

### Apps (`apps/`)

| App | Package | Port | Product / purpose | Dockerized? |
|-----|---------|------|-------------------|-------------|
| **Workforce** | `@anthropos/web-app` | 3000 | Primary product (`app.anthropos.work`): skill paths, AI simulations, org skill management, dashboard, **AI-readiness** (the member 3-step onboarding `components/ai-readiness/` + the manager dashboard `app/.../ai-readiness/`; gates DIFFER by surface — corrected v2.3 M219: the **member** funnel is gated on PostHog `flag_ai_readiness` **and** the org `ai_readiness` setting; the **manager dashboard** is gated on the GraphQL `aiReadinessEnabled` + the `isEnterprise` nav, and does **NOT** read the PostHog flag. Conflating them is the wrong-vantage error M219 spent a section correcting. A demo bakes no PostHog, so the flag resolves `undefined` forever and the member surface needs the `next-web-aireadiness-flag-gate` demo-patch — see [`ai-readiness.md`](ai-readiness.md)) | ✅ (the only one) |
| **Hiring** | `@anthropos/hiring-app` | 3001 | Distinct product (`hiring.anthropos.work`): job ladders, candidate funnels. **The demo's recruiter candidate-comparison scoreboard IS this app** — see the ⚠️ note below the table; the full hiring org-type + read-model is [`hiring.md`](hiring.md) | ❌ not in platform compose — but the **demo containerizes it** (note below) |
| **Integration** | `@anthropos/integration` | 3002 | Public-website embed (WordPress via proxy rewrites, SEO/Prerender) | ❌ Vercel-only |
| **Maintenance** | `@anthropos/maintenance-app` | — | Downtime/outage placeholder UI | ❌ |
| **Mobile** | `@anthropos/mobile` | 3031 (Expo) | Expo / React Native PoC (**paused**); **excluded** from the pnpm workspace, uses `EXPO_PUBLIC_*` | ❌ |

> **⚠️ RETRACTED — "the recruiter scoreboard is an `is_hiring` org-type surface in the dockerized `apps/web`, not
> the Hiring app."** That sentence stood in the Hiring row (and was implied by the *Key Functions* bullet above)
> and is **false in both directions**. Two independent adjudicator readings booked the same anchor; corrected
> M257x iter-102.
>
> * **Why it is false.** `/enterprise/activity-dashboard` exists in *both* apps, so the route's presence in
>   `apps/web` proves nothing. What decides it is a global product-boundary guard: measured @ `next-web-app`
>   **`8297c684`**, `apps/web/src/context/UserStatusContext.tsx:144-148` computes `userHasAllHiringOrgs` from
>   `membership.organization.publicMetadata.isHiring`, and when it holds, `:168-172` sets
>   `window.location.href = buildSwitchHandoffUrl({ targetProduct: 'hiring', … next: '/home' })` — the recruiter
>   is **ejected out of `apps/web`**, on a direct navigation too. So *"the org genuinely reads as hiring"* and
>   *"the scoreboard is reachable in `apps/web`"* are **mutually exclusive**. The screen that actually renders the
>   comparison is `apps/hiring/src/components/containers/InsightsByMembersContainer.tsx:108`, mounted at
>   `apps/hiring/…/enterprise/activity-dashboard/@tabs/ai-simulations/[simId]/page.tsx:14`.
> * **Which half was true.** The scoreboard *is* driven by the `is_hiring` **org-type** and *does* render from
>   seedable data with no platform edit — that half stands. Only the **app** was wrong.
> * **Consequence for the "Dockerized?" column.** `apps/hiring` is still absent from **platform** compose
>   (`platform` `0c91421` `docker-compose.yml` declares `sentinel`, `backend`, `studio-desk`, `next-web-app`,
>   `gotenberg` — the frontend service is `apps/web` only, at `:143`), and the repo ships one
>   `Dockerfile.dev`. But a **demo** builds `apps/hiring` as a **second UI container** from the same unmodified
>   clone using rext's own `demo-stack/frontend/hiring.Dockerfile` (`demo-stack/up-injected.sh:1076-1085`, image
>   `demo-<N>-hiring`, port `3001`+offset) — still zero platform-repo edits. `❌ Vercel-only` was therefore too
>   strong as well.
>
> Authoritative statement, with the render proof: [`hiring.md`](hiring.md) § *The render path* (M224).

### Shared packages (`packages/` + `configs/`)

| Package | Responsibility |
|---------|----------------|
| `@anthropos/graphql` | All data fetching: queries, React Query hooks by domain, server fetchers, **codegen output** (`src/__generated__`), `codegen.ts` |
| `@anthropos/ui` | Shared Ant Design 6 component library |
| `@anthropos/core-js` | Shared constants, types, utils (e.g. `security/printToken`) |
| `configs/*` | Workspace-scoped eslint / prettier / tailwind / tsconfig / **i18n** (8 locales: de, en, es, fr, it, ja, nl, pt) |

## Interface Discovery

* **GraphQL**: single endpoint `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` — compose bakes `http://${PUBLIC_HOST:-localhost}:8082/graphql/query`, as a build arg (`docker-compose.yml:151`) and again in the runtime environment (`:160`), re-anchored at platform `0c91421` (that anchor stood far further down the file at `0dab54d`); the env-var NAME still says wundergraph, the router behind it is gone locally; Clerk bearer token injected via React Query `defaultOptions.queries.meta.getToken`.
* **Auth edge**: **`apps/web/src/proxy.ts`** (and `apps/hiring/src/proxy.ts`, and `apps/integration/src/proxy.ts`) — **not `middleware.ts`**, which exists nowhere in the repo at `next-web-app` **`8297c684`** (re-derived 2026-08-06; the label here read the moving *"origin HEAD"* until M257x iter-102 — a pin is checkable, a moving label rots): **Next 16 renamed the `middleware.ts` convention to `proxy.ts`** (the repo's own `CLAUDE.md:55` says so, verbatim at that ref). `clerkMiddleware` protects every non-public route; public allowlist includes `/login`, `/sign-up`, `/checkout`, `/free-trial`, `/monitoring`, `/print`, `/api/bunny/thumbnail`. `/print` routes are HMAC-gated (`PRINT_ROUTE_SECRET`) for Puppeteer PDF generation.
* **Observability proxies**: `/logpoint/*` → PostHog (EU); `/monitoring` tunnels Sentry/Better Stack events.

## Dependencies

* **Downstream**: the GraphQL endpoint on `backend` (`:8082/graphql/query`; **in prod the router is destroyed — iter-124 — and where the deployed frontends now point is Vercel runtime config this corpus cannot read**), backend `app` API (`:8082`), Clerk, PostHog (EU), Sentry/Better Stack, Stripe (billing), Bunny CDN (thumbnails + Chime recordings, token-signed), Metabase (embedded analytics), Azure OpenAI/OpenAI (server AI routes).
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
pnpm codegen                     # regenerate GraphQL types (needs the endpoint — :8082/graphql/query locally)
pnpm check                       # tsc --noEmit + eslint --fix across the workspace
# pnpm storybook                 # REMOVED — no `storybook` script and no `.storybook/` dir exist at
                                 # next-web-app 8297c684 (re-derived 2026-08-06; this said the moving
                                 # "origin HEAD"); the only trace is configs/tailwind/storybooks.css
```

> Older Node fails with `WARN Unsupported engine`, and pnpm refuses to wipe
> `node_modules` in non-TTY shells (`ERR_PNPM_ABORTED_REMOVE_MODULES_DIR_NO_TTY`).

### Containerized (Workforce only)

```bash
cd platform
make up-frontend                 # builds Dockerfile.dev (web app only), serves :3000
                                 # Makefile:119-120 → --profile core --profile frontend
```

> ⚠️ **`make up PROFILE=frontend` on its own EXITS 1 — it builds nothing.** `next-web-app` declares
> `depends_on: backend` (`docker-compose.yml:165-167`) and `backend` is `profiles: [core, backend, all]`
> (`:110`), which the `frontend` profile does not select — so compose rejects the whole project with
> *"service `next-web-app` depends on undefined service `backend`: invalid compose project."* Use
> `make up-frontend` (which adds `core`), or `make up PROFILE=all`.

The repo's own `Dockerfile.dev` (Node 24 alpine — it is the only Dockerfile in the repo)
builds **only** `@anthropos/web-app`
(`pnpm turbo build --filter=@anthropos/web-app`), and `apps/web` is the only **`next-web-app`** app in
platform compose — ⚠️ **not "the only frontend"**: `studio-desk` is a second compose frontend with its own
browser UI (`platform` `0c91421df`, `docker-compose.yml:112`, `:135`
`VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query`). Integration ships via Vercel; **mobile builds
via EAS** (`apps/mobile/package.json:71-72`) and **maintenance is in no deploy pipeline** — the production
matrix is web / hiring / integration (`.github/workflows/production.yaml` @ `f97ba6599`), which this page's
own Apps table already marked `❌` for both while this sentence said "Vercel only". **Hiring is the
exception a demo makes**: `/demo-up` builds `apps/hiring` into a second UI container from
the same unmodified clone using a Dockerfile that lives in **rext**, not here — see the
⚠️ note under *Apps*. `NEXT_PUBLIC_*` are baked
at **build time**; on a remote VM set `PUBLIC_HOST` in `platform/.env` so the client
bundle resolves the right hostname.

## Environment Variables (high-signal subset)

| Variable | Default | Description |
|----------|---------|-------------|
| `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `CLERK_SECRET_KEY` | — | Clerk auth (client / server) |
| `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` | `http://localhost:8082/graphql/query` | Runtime GraphQL endpoint (baked at build). **Was `:5050/graphql`** until platform `2adcf71` deleted the router from compose |
| `NEXT_PUBLIC_BACKEND_API_URL` | `http://localhost:8082` | Backend (`app`) API base URL |
| `GRAPHQL_SCHEMA_FOR_GEN` | `http://localhost:8082/graphql/query` | ⚠️ **Read by nothing.** Declared in four `.env.example` files and in **no code** — `packages/graphql/codegen.ts:9` hardcodes the endpoint as a literal (`git grep -in SCHEMA_FOR_GEN f97ba6599` → 4 hits, all `.env.example`). Setting it changes nothing. Said *"used by `graphql-codegen`"* until M257x iter-129 |
| `NEXT_PUBLIC_HOSTING_URL` / `PUBLIC_HOST` | `http://localhost:3000` / `localhost` | Public hosting URL; `PUBLIC_HOST` parameterizes baked URLs in compose |
| `NEXT_PUBLIC_POSTHOG_KEY` / `_HOST` · `NEXT_PUBLIC_SENTRY_DSN` / `SENTRY_AUTH_TOKEN` | — | Analytics + error tracking (PostHog EU, Sentry/Better Stack) |
| `STRIPE_*` / `NEXT_PUBLIC_STRIPE_*` | — | Billing/checkout |
| `BUNNY_*` · `METABASE_*` · `PRINT_ROUTE_SECRET` | — | CDN, embedded analytics, signed PDF/print routes |
| `EXPO_PUBLIC_*` | — | Mobile (Expo) variants of the public vars |

## Testing

```bash
pnpm test            # turbo test → jest in apps/web and apps/hiring
                     # (integration & maintenance have no test script)
# E2E: Playwright suite under e2e/ (needs E2E_TEST_EMAIL for the Clerk sign-in ticket — see below)
```

### The two Clerk sign-in-token minting sites in this repo (added M257x iter-121)

**This repo holds 2 of the 5 sign-in-token minting sites in the whole clone set** — a token that buys a
**genuine session as any named user**, so the complete list matters. It is enumerated once, in
[`clerk-integration.md` § Sign-in tokens](./clerk-integration.md#sign-in-tokens--every-minting-site-enumerated);
this section exists so the repo's own page is not silent about it. Both at `next-web-app` `8297c684`:

| site | what it is | gate |
|---|---|---|
| `apps/web/src/app/api/dev/login-as/route.ts:79` | dev *"log in as a real Clerk user"* route → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise. **The same boolean adds `/api/dev/login-as` + `/dev/accept` to the PUBLIC route list** (`apps/web/src/proxy.ts:56`) — it must be reachable before a session exists |
| `e2e/auth.setup.ts:72` | the Playwright auth fixture — it mints a ticket instead of driving the password form | **no `NODE_ENV` gate.** It is a test-runner file, never in an app build, but it runs against a **real Clerk instance** |

**Why `auth.setup.ts` needs no `E2E_TEST_PASSWORD`, stated because the old testing note implied one:** the
e2e account *"enforces 2FA (email_code as second factor); password signin returns `needs_second_factor` and
never produces a session"*, and the ticket path means *"Clerk treats it as fully authenticated and **skips
both factors**"* (`e2e/auth.setup.ts:57-62`, the file's own words). So the suite needs `E2E_TEST_EMAIL` +
`CLERK_SECRET_KEY`, and the token is a **deliberate second-factor bypass** — the platform-side reading of
that is filed in `knowledge/plan/platform-defect-register.md`, not asserted here.

## Notable Gotchas

* **Next.js 16 / React 19** — the repo went 15 → 16 and the corpus missed it for four releases. Its own `CLAUDE.md:15` says *"Next.js 16 App Router"* and is **current** (an older note here claimed it still said 14; it does not). `knowledge/next15-adoption-plan.md` survives as a superseded plan beside `UPGRADE-IMPACT-next16.md`.
* **Only one Dockerfile** (`Dockerfile.dev`) exists at the repo root — the repo `CLAUDE.md` "two Dockerfiles" note is stale.
* **8 locales** on disk (Portuguese added) though some docs say 7.
* Frontend data layer is `graphql-request` + React Query, **not Apollo Client**.
* `npm`/`yarn` are blocked (`engines` + `please-use-pnpm`); conventional commits enforced (commitlint + Husky + cocogitto).

## The clone advanced to `19423a1fb` — what the 12 commits contain (M257x iter-256)

Every ref-pinned claim above is measured at `next-web-app` **`8297c684c`** (`v2.137.0`) and stays true
at that ref. On **2026-08-10** the clone and the canonical demo pin were advanced to `origin/main` =
**`19423a1fb`** (`v2.137.3`), **12 commits** on, under the user's closing condition that the milestone
may only close against the *current* branches.

The advance is **one product theme plus one bug pair**, and it is worth naming because both land on
surfaces this corpus already documents:

* **AI-Readiness ⇄ assignments** — the upskilling-plan tab is reworked around course progress, the
  launch modal leads with the assessment deadline, the auto-assign study window defaults to 3 months,
  and auto-assigned plans fold into a **per-cycle folder** in the assignments list. This is the
  frontend half of the `app` change landing in the same window (`source_ref airx:*` grouping and the
  weekly assignment-reminder cadence), so the two repos moved **together** — the coordinated
  multi-repo shape [`../ops/platform-alignment.md`](../ops/platform-alignment.md) § Trap D warns about.
* **Hiring `/start-sim`** — two fixes keep the invitation token across Clerk auth and set the fallback
  redirect pair on the `SignIn`. Relevant to the hiring Playthrough vantage.

**Not proven here:** that a demo built at `19423a1fb` comes up, or that the AI-Readiness demo surfaces
still render as the seeders expect. The advance is a source-and-pin change; gate clause 1 has **not**
been re-run against it.

## Related Documentation

* [Frontend Architecture](../architecture/frontend_architecture.md) — monorepo deep dive, packages, codegen, recent UX work
* [GraphQL Gateway](./graphql-wundergraph.md) — the federated gateway (**prod-only** since `2adcf71`); locally this app consumes `backend`'s own `:8082/graphql/query`
* [External Services → Clerk](../architecture/external_services.md#clerk-authentication-service)
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Dependency Map](../architecture/dependency_map.md)
