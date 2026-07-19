# Ant Academy

## High-Level Summary (For PMs & Non-Engineers)

**Ant Academy** (a.k.a. *AI Academy*) is the **internal learning portal** for Anthropos employees. It delivers micro-chapters on AI engineering, Claude Code, agent frameworks, and related topics to anyone with an `@anthropos.work` email.

Think of it as **the company's training app**:
- A web portal where employees take short, structured chapters
- A companion **iOS / Android app** (Expo / React Native) that bundles the same content for offline reading
- Authored content lives **inside the repo** as JSON, so curriculum changes ship through normal PRs

It is **not** a platform microservice. It is a standalone product that *uses* the platform's identity provider (Clerk) — and, since **v0.5.1**, **reads its course catalog from the platform's academy backend over GraphQL**. Without that backend it still boots and authenticates, but the catalog **degrades to empty**: the home grid renders **0 cards**. (This is exactly why the academy looks blank in a demo — see [*The Content Model*](#the-content-model--db-authoritative-catalog-v051-m7) below.)

## Technical Deep Dive (For Engineers)

### Service Overview

| Property | Value |
|:---------|:------|
| **Service Type** | Standalone application (Tier 2 — Studio/Standalone) |
| **Technology Stack** | Next.js 16 App Router + React 19.2 (React Compiler enabled), Expo / React Native (mobile) |
| **Deployment** | **Vercel native** (no Docker, no docker-compose entry). Mobile builds via Expo. |
| **Local dev port** | **3077** (web); **8555** (mobile web preview) |
| **Authentication** | Clerk (`@anthropos.work` domain gate + org-membership gate) |
| **Repository** | `git@github.com:anthropos-work/ant-academy.git` → `stack-dev/ant-academy/` |
| **In `repos.yml`** | **No — by design (v1.10b M49 #5).** NOT in `platform/repos.yml`, so `make init` / `make pull` do **not** clone/pull it. M49 did **not** add the entry: `repos.yml` lives in the *ephemeral, gitignored* `stack-demo/platform` clone (editing it is non-durable + a platform-repo edit). Instead, for a **demo**, `ensure-clones.sh` clones ant-academy **explicitly** (phase d2 — the cms/studio submodule-pattern precedent, non-fatal). For **dev**, clone it directly (it's a Vercel-native peripheral). The old "cloned by `make init`" claims are **stale**. |
| **In `docker-compose.yml`** | **No** — runs natively only |

### Role & Responsibility

- **Primary Goal**: Internal-only learning portal that delivers AI-engineering chapters to `@anthropos.work` employees, online and offline (PWA + mobile bundle).
- **Key Functions**:
  - Serve chapter content as a Next.js App Router site at `/chapters/<slug>/`
  - Cache chapters offline via a Serwist-built service worker
  - Bundle the same chapter JSON into the iOS / Android Expo app at build time
  - Provide an opt-in in-app AI assistant ("Cosmo", behind a feature flag) that talks to OpenAI directly from the browser
  - Author / publish / benchmark content via repo-local Claude skills (`.claude/skills/author-chapter`, `author-skill-path`, `author-podcast`, `author-cover`, `benchmark-chapter`, `build-index`, …)

### Repository Layout

```
ant-academy/
├── code/                  # Next.js 16 web app (this is where 99% of work happens)
│   ├── app/               # App Router routes (RSCs + client islands)
│   ├── src/               # Components, hooks, virtual library subsystem
│   ├── public/content/    # Chapter JSON — series / skill-path / chapter
│   ├── ucourses/          # Catalog, chapter engine, Cosmo AI assistant
│   ├── tests/             # Vitest (unit/integration/api) + Playwright e2e
│   ├── tools/             # offline-parity CLI
│   └── package.json       # npm scripts (dev/build/test/validate/...)
├── mobile/                # Expo / React Native app (iOS + Android)
├── knowledge/             # Architecture & authoring docs (project-overview, content-model, ...)
├── tools/course-validator/  # Node CLI: static checks against authoring rules
├── .claude/skills/        # Repo-local authoring/benchmarking skills (separate from platform skills)
├── .env.example           # Repo-root tooling env (NOT for the React app)
└── CLAUDE.md              # In-repo agent guide
```

The **React app's** env lives at `code/.env.example` (Clerk + AI keys); the **repo-root `.env`** is only for authoring-side tooling (`AUTOCONTENT_API_KEY`, `OPENAI_API_KEY` for cover generation).

### How It Fits Into the Platform

Ant Academy is architecturally a **sibling of `studio-desk` and `next-web-app`** — a frontend product that **reuses platform identity** and is a **backend-authoritative read/WRITE GraphQL client** of the platform `app` academy subgraph. It has no backend of its own, but it does call one: it **reads** the catalog (below) and, since **v0.5 "direct line" M2**, **writes** per-user progress to the platform backend (chapter progress, last-activity, bookmarks, certificates, study-time, feedback) — the platform `app internal/academy` store is the sole source of truth (there is NO localStorage/IDB source-of-truth). The earlier "does not call backend services / read-only client" framing is retired (corrected v2.5 M231): progress persists via GraphQL mutations (`upsertChapterProgress[Batch]` / `setLastActivity`, posted from `code/app/api/academy/beacon/route.js`) to Ent tables `academy_chapter_progress` / `academy_last_activity` / … in `app`. This makes a "played academy session" a **seedable server row** (via `app/cmd/academy-seed`).

```mermaid
graph LR
    subgraph External["External / SaaS"]
        Clerk[Clerk Auth]
        OpenAI[OpenAI / Anthropic APIs]
        Vercel[Vercel hosting]
    end

    subgraph Standalone["Standalone Apps (Tier 2)"]
        Academy[Ant Academy<br/>Next 16 + Expo]
        Desk[Studio-Desk]
    end

    subgraph Core["Core Backend (Tier 1, Docker)"]
        App[app]
        CMS[cms]
        Jobsim[jobsimulation]
    end

    Academy --> Clerk
    Academy -->|catalog: GraphQL academy subgraph| App
    Academy -.->|direct, no platform proxy| OpenAI
    Academy ==> Vercel
    Desk --> Clerk
    Desk --> Core
```

**Key contrasts** with the core Go services:
- No PostgreSQL schema of its own, no Atlas migrations
- No Connect-RPC, no Redis Streams
- **Provides** no GraphQL subgraph (it doesn't federate INTO Cosmo Router) — but it **consumes** the platform's
  **academy subgraph** as a GraphQL *client* (see below); "no subgraph" ≠ "no GraphQL"
- Its rendered catalog is **NOT static repo JSON** — since v0.5.1 it is **DB-authoritative**, read from the platform
  academy backend over GraphQL. The committed JSON is the *authoring source* + the dev *draft* layer + a separate
  machine index — not what the authenticated grid renders

The platform-shared concerns are **Clerk** (identity — Ant Academy reuses the platform's Clerk app so engineers log
in with the same identity) **and the academy backend** (`app internal/academy`, queried over GraphQL for the course
catalog).

### The Content Model — DB-authoritative catalog (v0.5.1 M7)

**The home grid does NOT read the committed JSON.** Since ant-academy **v0.5.1 (M7)** the catalog the portal renders is
**DB-authoritative** — queried from the **platform academy subgraph** (served by `app`'s `internal/academy`) over
GraphQL, not from the repo's JSON files. This is the most load-bearing fact about the academy and the root of the demo
"empty grid" symptom below.

**The read chain** (all under `code/`):

```
app/(authed)/page.jsx  →  resolveCatalogView()
    ├─ signed-in → getServerCatalogView()   (src/lib/serverTenant.js)
    └─ anonymous → getPublicCatalogView()    (empty-eid variant)
         ▼
getBackendCatalogView(eids)   (src/lib/backendContent.js)
    →  createServerGraphQLClient()   (src/graphql/server.js — endpoint = NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT)
    →  query academyCatalogSeries + academyCatalogSkillPaths   (tenant-filtered server-side by the user's eids)
```

`getServerCatalogView()` is literally `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView()`, so on
**any** failure the catalog becomes `emptyCatalogView() = { chapters: [], skillPaths: {}, series: [] }` → **0 cards**.
Three failure legs collapse to the same empty grid:
1. **Endpoint unset** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` empty → `createServerGraphQLClient()` throws → `makeClient()`
   returns `null` → `getBackendCatalogView()` returns `null`.
2. **Query error** — network/schema failure → `catch → return null`.
3. **Empty DB** — the query succeeds but the app academy tables have no rows → `[]`.

**Two DIFFERENT "catalog" data paths — do not confuse them:**
- **The grid READS** the catalog from `app internal/academy` **via GraphQL** (the DB-authoritative chain above).
- **`build-catalog.mjs` WRITES** a separate, unauthenticated `code/public/catalog.json` (~2,667 entries) — an
  FS-derived machine index built from the committed content tree, served at `/catalog.json` **for the platform
  backend's Talk-to-Data indexer** (the reverse-ingest noted in `proxy.js`). The grid never reads it. This is why the
  academy can serve `/catalog.json` HTTP 200 with thousands of entries **while the grid renders 0 cards** — unrelated
  sources.

**The committed JSON** (`code/public/content/<series>/<skill-path>/*.json`) is the **authoring source** (published into
the app academy DB via the repo's export path), the input to `build-catalog.mjs`, and the input to the dev **draft**
layer — but it is **not** what the authenticated grid renders.

**The dev DRAFT layer** (`src/lib/draftMode.js` + `src/lib/draftCatalog.js`): `draftsEnabled()` is
`NODE_ENV === 'development' && ACADEMY_SHOW_DRAFTS ∈ {1,true}` — a production **hard-block** (whitelisted on
`'development'`, so it can never leak into a deployed build). When on, `getServerCatalogView()` calls
`mergeDrafts(view, eids)`, which merges the **entire committed FS catalog** into the (possibly empty) backend view and
serves FS chapter bodies unlocked. It is the lightest possible demo-fill lever — but it stamps a visible **"Draft"**
chip on cards, so **v2.5 M230 deliberately chose the production-faithful path instead** (see
[`../ops/demo/frontend-tier.md`](../ops/demo/frontend-tier.md)).

**Why a demo grid is empty** (the v2.4 **F4** carry — long mis-attributed as an ant-academy-repo "client-side render
defect"): the demo launcher `demo-stack/ant-academy.sh` **never sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`** (`.env.example`
ships it empty), the demo runs via `next dev` (so `NODE_ENV=development`) with drafts **off**, and the demo app DB holds
no academy rows — so **failure legs 1 and 3 both hold** and the grid resolves to `emptyCatalogView()`. It is **not** a
render bug in the academy repo; it is the DB-authoritative catalog with **no reachable, populated source**. Filling it
(M230) is a demo-tooling concern — **zero academy-repo edits** (env / a sha-pinned demo-patch on the ephemeral clone).

### Tech Stack

| Layer | Technology |
|:------|:-----------|
| **Framework** | Next.js 16 App Router + React 19.2 (React Compiler enabled, Turbopack default) |
| **Auth** | `@clerk/nextjs` middleware in `proxy.js` (Next 16 renamed `middleware` → `proxy`). `clerkMiddleware()` + org-membership gate; `@anthropos.work` domain restriction is enforced in the Clerk app. Public routes: `/sign-in/*`, `/no-organization`, `/verify/*`, `/api/ai/chat`, `/library`, `/library/*`, `/free`, `/free/*`, `/local-content/*`, `/catalog.json`, `/academy-manifest.json` (other `/api/*` stay gated). The last three are public-by-design: `/local-content/*` for `<audio>` Range requests + cover previews, `/catalog.json` for the external Anthropos backend Talk-to-Data indexer, `/academy-manifest.json` for the PWA manifest (gating any of them 307s the fetch through sign-in and breaks it). |
| **Markdown** | `marked` (client-side rendering) |
| **Styling** | Vanilla CSS with custom properties (dark theme) |
| **Fonts** | DM Sans + Instrument Serif + JetBrains Mono (via `next/font/google`) + Font Awesome Pro **icons self-hosted/vendored in the repo** (`code/public/assets/fontawesome/` — `webfonts/*.woff2` + `css/all.min.css`, used as `<i class="fa-solid …">`; **not** pulled from the FA npm registry, so `npm install` needs no FA token) |
| **PWA** | Serwist 9 (configurator mode); service worker compiled by `serwist build` |
| **Mobile** | Expo SDK 54 / React Native (Expo Router) |
| **Testing** | Vitest (happy-dom + node), Playwright (e2e). 1000+ Vitest tests + ~26 Playwright e2e spec files (tests/e2e/). |
| **Deployment** | Vercel native (minimal `code/vercel.json` — only `{"framework": "nextjs"}`; Next.js handles routing). Mobile builds via Expo. |
| **Node** | `>= 22` (declared in `code/package.json` `engines`) |

### Local Development

#### Prerequisites
- Node **v22+** (from `code/package.json` `engines.node`)
- npm (web app uses npm, not pnpm)
- pnpm — only if you also want to run the mobile app
- Clerk credentials (use the platform's dev tenant — same `@anthropos.work` domain)
- _(vestigial — NOT required)_ A Font Awesome Pro npm token. The FA Pro icons are vendored in the repo (`code/public/assets/fontawesome/`), so a fresh, token-less `npm install` succeeds and the app serves working icons. `FONTAWESOME_NPM_AUTH_TOKEN` survives in `code/.env.example` but is **not** needed to install or run.

#### 1. Clone

ant-academy is **NOT** in `platform/repos.yml` (by design — v1.10b M49 #5), so `make init` does **not** clone it.
This was the demo-up #5 "ant-academy never cloned" gap; **M49 fixed it for a demo by cloning ant-academy
explicitly** in `ensure-clones.sh` (phase d2 — `repos.yml` lives in the ephemeral platform clone, so editing it
is non-durable + a platform-repo edit; the explicit clone mirrors the `make init-studio` exception). **For a
demo, `/demo-up` clones it automatically.** For **dev** (no `ensure-clones.sh`), clone it directly:

```bash
# Demo: ensure-clones.sh does this automatically on /demo-up (phase d2) — into stack-demo/ant-academy/.
# Dev (or a manual clone): clone it directly — it's a Vercel-native peripheral, not in repos.yml.
cd stack-dev
git clone git@github.com:anthropos-work/ant-academy.git
```

> **In a demo the academy is AUTHENTICATED-as-a-member, keyless (v1.10b "fit-up" M53 F6).** `/demo-up` launches
> the academy with **both** halves of its own `e2e_persona` cookie bypass set — the server gate
> `BENCHMARK_VISUAL_BYPASS=1` **and** the client gate `NEXT_PUBLIC_E2E_AUTH=1` — so an `e2e_persona=member`
> cookie drives a **signed-in** context (server RSC `anonymous=false` + entitlement; client Clerk hooks resolve a
> synthetic **`E2E Member`**) with **no real Clerk keys**. ⚠ **CORRECTION (v2.3.2, 2026-07-15): the cockpit's
> per-hero [Academy] link was REMOVED — the cockpit is now login-only** (per user request), so it **no longer
> sets the `e2e_persona` cookie**. Formerly, that link set the cookie browser-side (cookies on `localhost` are
> port-agnostic, so the cockpit origin's cookie was read by the academy origin) then navigated in, and a hero
> landed **authenticated**; reaching the demo academy as a member now needs the cookie set by other means (and
> the academy grid renders **empty** in a demo — the v2.4 **F4** carry, **NOT** a client-side render defect: the
> catalog is [DB-authoritative](#the-content-model--db-authoritative-catalog-v051-m7) and the demo neither sets
> `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` nor holds academy rows → `emptyCatalogView()`. v2.5 **M230** fills it
> production-faithfully, zero academy-repo edits). The **Cosmo AI assistant stays absent** in a demo (its flag +
> OpenAI key are deliberately not provisioned — the AI-keys policy). Zero academy-repo edits: the flags live in the gitignored `code/.env.local`; the cookie is set by the
> standalone cockpit panel. Full mechanics: [`../ops/demo/frontend-tier.md` § ant-academy](../ops/demo/frontend-tier.md).

#### 2. Configure env

The **app's** env file is `code/.env`, not the repo root:

```bash
cd stack-dev/ant-academy/code
cp .env.example .env
# Minimum to boot locally:
#   NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY
#   CLERK_SECRET_KEY
# Vestigial — NOT required (FA Pro icons are vendored in the repo; token-less npm install works):
#   FONTAWESOME_NPM_AUTH_TOKEN
# Needed only for the server-side /api/ai/chat route handler:
#   OPENAI_API_KEY        (server-side)
#   ANTHROPIC_API_KEY     (server-side)
# Optional:
#   NEXT_PUBLIC_STUDIO_URL              (Studio Desk URL bridge)
#   NEXT_PUBLIC_FEATURE_TRAINING_COACH  (1/true to turn on Cosmo — default OFF)
#   REQUIRE_ORGANIZATION_MEMBERSHIP     (0/false to skip the org gate in solo dev)
```

`code/.env.example` is the authoritative, fuller list (it also carries Sentry / Better Stack DSN vars and the `NEXT_PUBLIC_CLERK_SIGN_IN_URL` / `SIGN_UP_URL` family used to keep Clerk on the in-app sign-in page). Reuse the **same Clerk keys** as in `platform/.env` so dev login works across the platform and the academy with a single session.

> **Org-membership gate**: by default, `proxy.js` redirects signed-in users with zero org memberships to `/no-organization`. For solo local dev without an org, set `REQUIRE_ORGANIZATION_MEMBERSHIP=0` in `code/.env`.

#### 3. Install & run (web)

```bash
cd stack-dev/ant-academy/code
npm install
npm run dev          # next dev — port 3077 (3000 is reserved on dev machines)
```

Open <http://localhost:3077>.

#### 4. Install & run (mobile, optional)

```bash
cd stack-dev/ant-academy/mobile
pnpm install
pnpm run dev:web     # web preview at :8555 (Playwright-friendly)
# or run on a real device / simulator with Expo Go
```

The mobile app bundles `code/public/content/` at build time via `pnpm run dev:bundle`.

#### 5. Tests

```bash
cd code
npm test                  # vitest run (unit + integration + api)
npm run test:e2e          # playwright (boots dev server)
npm run validate -- --all # course-validator across all chapters
```

### Repo-Local Claude Skills

`ant-academy/.claude/skills/` ships **its own** set of skills focused on **authoring content** — not to be confused with the platform's `/ant-*` skills in Rosetta:

| Skill | Purpose |
|-------|---------|
| `author-chapter` | Draft a new chapter JSON from an outline |
| `author-skill-path` | Bootstrap a new skill-path directory + path intro |
| `author-podcast` | Generate the `path-intro.mp3` companion audio (uses AutoContent API) |
| `author-cover` | Generate the `cover.{png,webp}` for a skill path (uses OpenAI `gpt-image-2`) |
| `benchmark-chapter` | Drive Playwright through a chapter for visual + content benchmarking |
| `translate-path` | Translate a skill path into another language (i18n pipeline) |
| `benchmark-translation` | Benchmark translation quality/coverage for a translated path |
| `check-plagiarism`, `consolidate-library`, `extend-library`, `build-index`, `publish`, `preview`, `cover-scale` | Other content-pipeline helpers |

Academy is multi-language: content is authored per language, `catalog.json` is emitted per chapter × language, and `proxy.js` propagates `?lang=` into SSR via an `x-locale` header (see `tests/e2e/i18n-language-toggle.spec.js`).

These are **isolated to the ant-academy repo** and are loaded only when working inside it. They share no state with the Rosetta corpus skills.

### Deployment

- **Web**: Pushed to Vercel via `.github/workflows/deploy-academy.yaml`
- **Vercel env sync**: `.github/workflows/sync-vercel-env.yml` mirrors env vars
- **Mobile**: Expo build pipeline (outside platform CI)
- **Coverage CI**: `.github/workflows/sidebar-coverage-tests.yaml`

Releases use **Cocogitto** conventional-commit tagging (`cog.toml`).

### Integration Points

- **Clerk (shared)**: Uses the same Clerk app as the rest of the platform. Domain-gated to `@anthropos.work` so external users cannot enter.
- **OpenAI (direct, browser, opt-in)**: The in-app "Cosmo" assistant — gated behind `NEXT_PUBLIC_FEATURE_TRAINING_COACH` (default OFF) — calls the **OpenAI Responses API** (`gpt-5.2`, `https://api.openai.com/v1/responses`) directly from the browser using a per-user `localStorage('openai_api_key')`. It is OpenAI-only and does **not** route through the platform's shared `ai` library or the `/api/ai/chat` route. (The separate server-side `/api/ai/chat` route handler does support both OpenAI and Anthropic with server keys, but Cosmo does not use it.)
- **Studio Desk (loose link)**: `NEXT_PUBLIC_STUDIO_URL` can deep-link from the academy to the Studio Desk UI; nothing required at runtime.
- **next-web-app (iframe embed):** the Workforce app loads Academy in an iframe with `?embed=anthropos`; `proxy.js` detects this server-side (`Sec-Fetch-Dest=iframe` + an `embed-mode` cookie, or an explicit `?embed=anthropos`), stamps an `x-embed-mode` request header, persists the cookie, and SSRs a light-themed, topbar-less variant (`data-embed` + `data-theme=light`). No data flows back — it is a presentation/cookie coupling only.
- **Platform academy backend (GraphQL, load-bearing):** the home grid **reads its course catalog** from the platform
  **academy subgraph** (`app internal/academy`) over GraphQL at `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` — `getBackendCatalogView()`
  queries `academyCatalogSeries` + `academyCatalogSkillPaths`, tenant-filtered server-side. This is the catalog source of
  truth since v0.5.1 (M7); on failure the grid degrades to `emptyCatalogView()` (0 cards). See
  [*The Content Model*](#the-content-model--db-authoritative-catalog-v051-m7). *(No Connect-RPC and no Redis events — the
  academy talks to the backend over GraphQL only: it **reads** the catalog AND **writes** per-user progress via
  `upsertChapterProgress`/`setLastActivity` mutations since v0.5 M2 — see § How It Fits Into the Platform.)* The reverse ingest also exists: per a
  comment in `proxy.js`, the platform backend's Talk-to-Data indexer pulls Academy's **separate** public `/catalog.json`
  (an FS-derived, metadata-only, per-chapter × language index — **not** the grid's source).

### Why It's Not in `docker-compose.yml`

Ant Academy is deployed to Vercel and runs natively in dev (`npm run dev`) just like `studio-desk` natively or `next-web-app` natively. It has no upstream service it needs to wait on, no migrations to apply, and its container would only duplicate what Vercel already serves. We deliberately mirror the studio-desk pattern: **clone, run natively, skip docker-compose** — though, unlike
studio-desk, ant-academy is **not in `repos.yml`** (by design — v1.10b M49 #5 kept it out, since `repos.yml`
lives in the ephemeral platform clone). So for a **demo**, `ensure-clones.sh` clones it **explicitly** (phase
d2, non-fatal); for **dev**, it's a manual `git clone` — not `make init`.

If you ever need to add a Docker profile (e.g. for an integration-test harness), follow studio-desk's containerized variant as the template.

### Related Documentation
- [Service Taxonomy](../architecture/service_taxonomy.md) — where Ant Academy sits in the three-tier model
- [Architecture Overview](../architecture/architecture_overview.md) — overall platform diagram
- [Frontend Architecture](../architecture/frontend_architecture.md) — sibling frontend (`next-web-app`)
- [Studio-Desk](./studio-desk.md) — closest deployment-shape sibling
- [Run Guide](../ops/run_guide.md) — how to start the academy alongside the rest of the platform
- [External Services](../architecture/external_services.md) — Clerk integration details
