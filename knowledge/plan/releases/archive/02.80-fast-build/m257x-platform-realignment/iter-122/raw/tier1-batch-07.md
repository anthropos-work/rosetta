# TIER-1 ADJUDICATION BATCH 07 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 07-001
- **id**: `B07-001`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `src/routes/skillpath.ts:44-47`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/routes/skillpath.ts`  (1492 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
    41  }
    42  
    43  function getDirectusConfig(): { url: string; token: string } {
    44    const url = (process.env.DIRECTUS_BASE_URL || '').replace(/\/$/, '');
    45    const token = process.env.DIRECTUS_TOKEN || '';
    46    if (!url || !token) {
    47      throw new Error('Directus not configured (check DIRECTUS_BASE_URL and DIRECTUS_TOKEN in .env)');
    48    }
    49    return { url, token };
    50  }
```

## 07-002
- **id**: `B07-002`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `src/index.ts:303-310`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/index.ts`  (321 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
   300    // every restore point. The check is async + best-effort: a Directus that
   301    // is briefly unreachable on boot shouldn't crash the server, but a
   302    // confirmed permission denial should be loud enough to spot in CI logs.
   303    // Skipped when DIRECTUS_BASE_URL is unset (e.g. unit-test bootstrapping).
   304    if (process.env.DIRECTUS_BASE_URL && process.env.DIRECTUS_TOKEN) {
   305      pingSnapshotCapability().then(ok => {
   306        if (ok) {
   307          console.log('[skillpath] Snapshot capability OK (token can write directus_versions).');
   308        } else {
   309          console.warn(
   310            '[skillpath] Snapshot capability MISSING — the configured DIRECTUS_TOKEN ' +
   311            'cannot create rows in `directus_versions`. Publish/unpublish will keep ' +
   312            'working but no restore points will be saved. Grant the role `create` ' +
   313            'permission on `directus_versions` in Directus admin to fix.'
```

## 07-003
- **id**: `B07-003`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `userService.ts:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/services/userService.ts`  (279 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
    17  
    18  class UserService {
    19    private clerk: Clerk | null = null;
    20    private graphqlClient: GraphQLClient = new GraphQLClient(config.GRAPHQL_ENDPOINT, {
    21      headers: {
    22        'Content-Type': 'application/json',
    23      },
```

## 07-004
- **id**: `B07-004`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `taxonomyService.ts:43`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/services/taxonomyService.ts`  (530 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
    40      private bannedSpecializationIds: string[] = BANNED_SPECIALIZATION_IDS;
    41  
    42      constructor() {
    43          this.graphqlClient = new GraphQLClient(config.GRAPHQL_ENDPOINT, {
    44              headers: {
    45                  'Content-Type': 'application/json',
    46              },
```

## 07-005
- **id**: `B07-005`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `userPreferencesService.js:13`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/services/userPreferencesService.js`  (240 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
    10   */
    11  class UserPreferencesService {
    12      constructor() {
    13          this.client = new GraphQLClient(config.GRAPHQL_ENDPOINT);
    14          this.clerk = null;
    15          this.preferencesCache = null;
    16          this.cacheTimestamp = null;
```

## 07-006
- **id**: `B07-006`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `content/simulationContentService.js:325`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/services/content/simulationContentService.js`  (866 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
   322  class SimulationContentService {
   323  
   324      constructor() {
   325          this.client = new GraphQLClient(config.GRAPHQL_ENDPOINT);
   326          this.lastGeneratedContent = null;
   327          this.clerk = null;
   328      }
```

## 07-007
- **id**: `B07-007`
- **corpus site**: `corpus/services/studio-desk.md:36-53` (bullet)
- **citation**: `app/services/config.ts:6`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/services/config.ts`  (10 lines)

**CLAIMING UNIT**

```md
2. **Backend**: Express.js API server
   - Clerk middleware for route protection
   - ⚠️ **NOT a GraphQL client — corrected M257x iter-115.** At `studio-desk` `41ee3575`,
     `git grep -in graphql -- 'src/*'` returns exactly **two** lines, both comments saying the opposite
     (`src/routes/skillpath.ts:374` *"We do NOT route this through the platform's `privateSkillPaths`
     GraphQL"*, and `:405`); `git grep -n 8082 -- 'src/*'` returns **0**; and `src/index.ts` mounts four
     API routers — `/api/dev` (`:150`), `/api/ai` (`:158`), `/api/skillpath` (`:161`), `/api/youtube`
     (`:164`) — none of them GraphQL. **The Express backend's real remote dependency is Directus over
     REST** (`DIRECTUS_BASE_URL`/`DIRECTUS_TOKEN`, read at `src/routes/skillpath.ts:44-47` and
     `src/index.ts:303-310`). Every `new GraphQLClient(...)` in the repo is in the **frontend**
     (`app/services/{userService.ts:20, taxonomyService.ts:43, userPreferencesService.js:13,
     content/simulationContentService.js:325}`), fed by `app/services/config.ts:6` reading the
     **`VITE_`-prefixed, browser-baked** `VITE_GRAPHQL_ENDPOINT`. This file states it correctly in four
     other places — the Directus integration note, the `app/services/graphql/` example, the
     `VITE_GRAPHQL_ENDPOINT` config line and the env table — so this was a live self-contradiction,
     not a stale leftover
   - Multi-provider AI integration (Azure OpenAI / OpenAI / Anthropic) for Studio Copilot
   - File upload handling
```

**CITED CONTENT**

```
     3    CLERK_SIGN_IN_URL: (import.meta as any).env.VITE_CLERK_SIGN_IN_URL || 'http://localhost:3000/login',
     4    // backend directly, not the Cosmo router — since cms-in-app the router composes
     5    // a single `backend` subgraph, so it is the same schema one hop shorter.
     6    GRAPHQL_ENDPOINT: (import.meta as any).env.VITE_GRAPHQL_ENDPOINT || 'http://localhost:8082/graphql/query',
     7    WEBAPP_URL: (import.meta as any).env.VITE_WEB_APP_URL || 'http://localhost:3000',
     8    ENVIRONMENT: (import.meta as any).env.VITE_ENVIRONMENT || 'development',
     9  };
```

## 07-008
- **id**: `B07-008`
- **corpus site**: `corpus/services/studio-desk.md:121-121` (paragraph)
- **citation**: `src/routes/youtube.ts:43`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/routes/youtube.ts`  (99 lines)

**CLAIMING UNIT**

```md
A builder for learning skill paths, served at `/builder-skill-path` (`app/builder-skill-path` module). Backed by `/api/skillpath` (the largest backend route, ~61KB) and `/api/youtube`. Integrates directly with Directus (`DIRECTUS_BASE_URL` / `DIRECTUS_TOKEN`) and uses `directus_versions` for publish/unpublish snapshot & restore (capability checked at boot via `pingSnapshotCapability`). The skill-path **writes** (create/publish) go to Directus as a `Bearer ${DIRECTUS_TOKEN}` static token (`src/routes/skillpath.ts`). Curates videos from a Bunny CDN library (`BUNNY_LIBRARY_ID` / `BUNNY_LIBRARY_API_KEY`) and searches YouTube via the YouTube Data API v3 through a `YouTubePicker` — the route reads **`YOUTUBE_API_KEY` only** (`src/routes/youtube.ts:43`; with no key it serves a `_mock: true` fallback list). `GCLOUD_SERVICE_ACCOUNT` is declared in `.env.example:**119**` (@ `studio-desk` `41ee3575`) and injected by `terraform/main.tf:129`, but **no code in `src/` reads it** — treat it as vestigial, not a second YouTube credential. (This cited `.env.example:120` until M257x iter-115. The file is 131 lines, so `:120` is **in range and resolves — to a blank line**, which is the failure mode a range check cannot catch: `:117` is `YOUTUBE_API_KEY=`, `:118` the comment, `:119` the declaration, `:120` empty. The other two thirds of the sentence verified exactly.)
```

**CITED CONTENT**

```
    40      return;
    41    }
    42  
    43    const apiKey = process.env.YOUTUBE_API_KEY;
    44    if (!apiKey) {
    45      // Mock fallback
    46      res.json({ data: MOCK_VIDEOS.slice(0, limit), _mock: true });
```

## 07-009
- **id**: `B07-009`
- **corpus site**: `corpus/services/studio-desk.md:121-121` (paragraph)
- **citation**: `terraform/main.tf:129`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
A builder for learning skill paths, served at `/builder-skill-path` (`app/builder-skill-path` module). Backed by `/api/skillpath` (the largest backend route, ~61KB) and `/api/youtube`. Integrates directly with Directus (`DIRECTUS_BASE_URL` / `DIRECTUS_TOKEN`) and uses `directus_versions` for publish/unpublish snapshot & restore (capability checked at boot via `pingSnapshotCapability`). The skill-path **writes** (create/publish) go to Directus as a `Bearer ${DIRECTUS_TOKEN}` static token (`src/routes/skillpath.ts`). Curates videos from a Bunny CDN library (`BUNNY_LIBRARY_ID` / `BUNNY_LIBRARY_API_KEY`) and searches YouTube via the YouTube Data API v3 through a `YouTubePicker` — the route reads **`YOUTUBE_API_KEY` only** (`src/routes/youtube.ts:43`; with no key it serves a `_mock: true` fallback list). `GCLOUD_SERVICE_ACCOUNT` is declared in `.env.example:**119**` (@ `studio-desk` `41ee3575`) and injected by `terraform/main.tf:129`, but **no code in `src/` reads it** — treat it as vestigial, not a second YouTube credential. (This cited `.env.example:120` until M257x iter-115. The file is 131 lines, so `:120` is **in range and resolves — to a blank line**, which is the failure mode a range check cannot catch: `:117` is `YOUTUBE_API_KEY=`, `:118` the comment, `:119` the declaration, `:120` empty. The other two thirds of the sentence verified exactly.)
```

**CITED CONTENT**

```
   126    dir     = "${path.module}/migrations-sentinel?format=atlas"
   127    version = data.atlas_migration.sentinel_migration[0].latest
   128    url     = data.atlas_migration.sentinel_migration[0].url
   129    // NOT "public". The two pipelines must never share a revisions table: one
   130    // atlas_schema_revisions holding two unrelated migration histories breaks the
   131    // integrity check of both. Measured: no atlas_schema_revisions exists in the
   132    // sentinel schema today, so this one is free to be born there.
```

## 07-010
- **id**: `B07-010`
- **corpus site**: `corpus/services/studio-desk.md:154-154` (paragraph)
- **citation**: `docker-compose.yml:119`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
**GraphQL Endpoint**: Configured via `VITE_GRAPHQL_ENDPOINT` — compose bakes `http://localhost:8082/graphql/query` as a build arg (`docker-compose.yml:119`) and again in the runtime environment (`:135`), re-anchored at platform `0c91421` (it was `:204` at `0dab54d`); was `http://localhost:5050/graphql` when the router existed locally
```

**CITED CONTENT**

```
   116        ssh: ["default"]
   117        args:
   118          VITE_CLERK_PUBLISHABLE_KEY: ${VITE_CLERK_PUBLISHABLE_KEY}
   119          VITE_GRAPHQL_ENDPOINT: ${VITE_GRAPHQL_ENDPOINT:-http://localhost:8082/graphql/query}
   120          VITE_ENVIRONMENT: ${VITE_ENVIRONMENT:-production}
   121          VERSION: dev
   122      ports:
```

## 07-011
- **id**: `B07-011`
- **corpus site**: `corpus/services/studio-desk.md:293-301` (bullet)
- **citation**: `app/core/main.ts:105`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/core/main.ts`  (227 lines)

**CLAIMING UNIT**

```md
- **No local datastore — studio-desk has no database of any kind.** `package.json` declares no DB
  driver (no `pg`/`postgres`/`prisma`/`sqlite`/`mysql`/`mongo`/`knex`/`drizzle`/`typeorm`/`sequelize`),
  and nothing in `src/` reads a `DATABASE_URL` or opens a pool/client. All persistence is remote over
  HTTP: skill-path content goes to **Directus** (`DIRECTUS_BASE_URL` / `DIRECTUS_TOKEN`,
  `src/routes/skillpath.ts`), and per-user studio preferences (including the recoverable draft window in
  `app/services/studioDB.js` — a facade, not a datastore) round-trip through the platform **GraphQL** API
  via `GET_USER_STUDIO_PREFERENCES` / `SET_USER_STUDIO_PREFERENCES`. There is no Clerk-user sync job.
  (The repo's only Tailscale-funnel mention is `app/core/main.ts:105` — the public ingest URL of the
  self-hosted **GlitchTip** Sentry endpoint. Error telemetry, unrelated to users or data.)
```

**CITED CONTENT**

```
   102  }> {
   103    if (process.env.NODE_ENV === 'production') {
   104      Sentry.init({
   105        // Self-hosted GlitchTip (public ingest via tailscale funnel). Browser SDK runs
   106        // off-tailnet, so the public :10000 host is required.
   107        dsn: 'https://b86e49a1cd5c45fab992c76a1180e1b2@singularity-obs.taildc510.ts.net:10000/15',
   108        environment: config.ENVIRONMENT,
```

## 07-012
- **id**: `B07-012`
- **corpus site**: `corpus/services/studio-room.md:224-224` (paragraph)
- **citation**: `CLAUDE.md:12-14`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/CLAUDE.md`  (581 lines)

**CLAIMING UNIT**

```md
The repo's own `CLAUDE.md:12-14` gives the real entry point:
```

**CITED CONTENT**

```
     9  2. **Environment Setup**: Manual for humans and AI agents to build local development environments
    10  3. **Recursive Inspection**: Tool for reverse-engineering and documenting the platform itself
    11  
    12  This is NOT the Anthropos platform source code - it's the documentation about it. The actual platform code lives in separate repositories under the `anthropos-work` GitHub organization.
    13  
    14  ## Development Commands
    15  
    16  ### Available Skills
    17  
```

## 07-013
- **id**: `B07-013`
- **corpus site**: `corpus/architecture/ai_architecture.md:33-42` (bullet)
- **citation**: `internal/cms/studio/markdownManager.go:30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/markdownManager.go`  (144 lines)

**CLAIMING UNIT**

```md
5. **Mistral is nowhere in this path.** *Every* use of it in `app` is **OCR**, never generation: Go-side
   in the cms domain (`internal/cms/studio/markdownManager.go:30` — the constructor body
   `return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil` inside `func NewMarkdownManager` at `:29`,
   **@ `app` `ad9f3c49`**; this row cited **`:19`** until M257x iter-115, and `:19` is a **doc-comment**
   line — *"It used to take aiKey and then IGNORE it"* — not code, exactly as
   the *AI Providers* section of [`external_services.md`](external_services.md) already said while this line went on asserting it —
   and `studioManager.go:583`), and Python-side
   in the in-image studio tree at `studio/tools/pdf2md.py:24` (`mistral-ocr-latest`) — a standalone CLI on
   neither the AI manager's path nor the generation pipeline's
   (`git -C app/studio grep -i mistral aeec036a`, 22 hits / 3 files).
```

**CITED CONTENT**

```
    27  // may grow a failing constructor again — but there is nothing left in here that can
    28  // fail.
    29  func NewMarkdownManager(aiKey string) (*MarkdownManager, error) {
    30  	return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil
    31  }
    32  
    33  func (m *MarkdownManager) OCRProcess(ctx context.Context, documentData []byte) (*string, int, error) {
```

## 07-014
- **id**: `B07-014`
- **corpus site**: `corpus/architecture/ai_architecture.md:33-42` (bullet)
- **citation**: `studio/tools/pdf2md.py:24`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/tools/pdf2md.py`  (233 lines)

**CLAIMING UNIT**

```md
5. **Mistral is nowhere in this path.** *Every* use of it in `app` is **OCR**, never generation: Go-side
   in the cms domain (`internal/cms/studio/markdownManager.go:30` — the constructor body
   `return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil` inside `func NewMarkdownManager` at `:29`,
   **@ `app` `ad9f3c49`**; this row cited **`:19`** until M257x iter-115, and `:19` is a **doc-comment**
   line — *"It used to take aiKey and then IGNORE it"* — not code, exactly as
   the *AI Providers* section of [`external_services.md`](external_services.md) already said while this line went on asserting it —
   and `studioManager.go:583`), and Python-side
   in the in-image studio tree at `studio/tools/pdf2md.py:24` (`mistral-ocr-latest`) — a standalone CLI on
   neither the AI manager's path nor the generation pipeline's
   (`git -C app/studio grep -i mistral aeec036a`, 22 hits / 3 files).
```

**CITED CONTENT**

```
    21  from typing import Optional
    22  
    23  import tqdm
    24  from mistralai import Mistral
    25  from dotenv import load_dotenv
    26  
    27  MISTRAL_API_KEY = None
```

## 07-015
- **id**: `B07-015`
- **corpus site**: `corpus/architecture/ai_architecture.md:50-56` (paragraph)
- **citation**: `app/internal/coursebuilder/bedrock.go:106-113`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
**`ANTHROPIC_API_KEY` is a third exit but is NOT within the manager** — an either/or backend switch for
**Course Builder** (`app/internal/coursebuilder/bedrock.go:106-113`), which never touches `AIManager`. An
earlier revision counted it as one of *"exactly three things … within the manager"*, a category error
rather than a miscount. *(It does **not** additionally "flip Studio-Room off Bedrock": Studio-Room was
never on Bedrock — `grep -rin 'bedrock\|boto3' app/studio/` returns **0** hits, and there the key is a
credential while the selector is the ini's `TARGET SERVICE`; see the provider row at
[`external_services.md:567`](external_services.md). Corrected M257x iter-48.)*
```

**CITED CONTENT**

```
   103  }
   104  
   105  // newUnderlyingClient picks the backend for one model role:
   106  // ANTHROPIC_API_KEY present → the first-party Anthropic API (with the
   107  // model id normalized to its bare form); absent → AWS Bedrock, the
   108  // legacy path, byte-for-byte what shipped before the switch existed.
   109  func newUnderlyingClient(ctx context.Context, modelID string) (*askengine.BedrockClient, error) {
   110  	if key := strings.TrimSpace(os.Getenv(AnthropicAPIKeyEnv)); key != "" {
   111  		return askengine.NewAnthropicClientWithModel(key, directModelID(modelID))
   112  	}
   113  	return askengine.NewBedrockClientWithModel(ctx, modelID)
   114  }
   115  
   116  // authorNoSamplingUnderlying is the subset of *askengine.BedrockClient
```

## 07-016
- **id**: `B07-016`
- **corpus site**: `corpus/architecture/ai_architecture.md:58-65` (paragraph)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:905`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
One layer *above* the manager, a sequence whose **`ai_vendor` is unset** reaches **direct US OpenAI on the
first attempt, with no error condition** — the residency-relevant route, and the one nothing in
`app/internal/jobsimulation/ai/ai.go` can show you. The nullable field is on the **Directus DTO**
(`app/internal/cms/directus/collections/jobsimulation.go:905`, `AIVendor *AIVendor`), not on the domain
`simulation.Sequence`; nil is replaced by `simulation.Openai` at `:1302-1305` **before** the sequence is
built at `:1307`, so the value reaching the vendor switch takes **`case simulation.Openai:`**
(`simulator/ai/ai.go:58-59`) — the same arm an explicit caller takes. Full per-line derivation:
[`external_services.md:619-629`](external_services.md) (item 4 of the four live EU-exit routes).
```

**CITED CONTENT**

```
   902  	ValidationAcceptanceCriteria []ValidationCriterion `json:"validation_acceptance_criteria"`
   903  	ValidationEvaluationCriteria []ValidationCriterion `json:"validation_evaluation_criteria"`
   904  
   905  	AIVendor *AIVendor `json:"ai_vendor"`
   906  	AIModel  *AIModel  `json:"ai_model"`
   907  
   908  	EvaluationSkills  []Skills `json:"evaluation_skills,omitempty"`
```

## 07-017
- **id**: `B07-017`
- **corpus site**: `corpus/architecture/ai_architecture.md:92-92` (table-row)
- **citation**: `app/internal/ai/anthropic/completion.go:20-30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/ai/anthropic/completion.go`  (221 lines)

**CLAIMING UNIT**

```md
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:567`](external_services.md)) |
```

**CITED CONTENT**

```
    17  	"github.com/aws/aws-sdk-go-v2/aws"
    18  )
    19  
    20  const (
    21  	Anthropic35SonnetAWS20241022 ai.Model = "eu.anthropic.claude-3-5-sonnet-20240620-v1:0"
    22  	Anthropic35Sonnet20241022    ai.Model = "claude-3-5-sonnet-20241022"
    23  	Anthropic37SonnetAWS20250219 ai.Model = "eu.anthropic.claude-3-7-sonnet-20250219-v1:0"
    24  	Anthropic4SonnetAWS20250514  ai.Model = "eu.anthropic.claude-sonnet-4-20250514-v1:0"
    25  	Anthropic45SonnetAWS20251126 ai.Model = "eu.anthropic.claude-sonnet-4-5-20250929-v1:0"
    26  	// AWS named the Sonnet 4.6 EU inference profile without a date/version suffix
    27  	// (unlike 4.5 and older). The dated "eu.anthropic.claude-sonnet-4-6-20251126-v1:0"
    28  	// form does not exist in Bedrock and returns "model identifier is invalid".
    29  	Anthropic46SonnetAWS20251126 ai.Model = "eu.anthropic.claude-sonnet-4-6"
    30  )
    31  
    32  const (
    33  	anthropicDefaultModel           = Anthropic35SonnetAWS20241022
```

## 07-018
- **id**: `B07-018`
- **corpus site**: `corpus/architecture/ai_architecture.md:92-92` (table-row)
- **citation**: `internal/coursebuilder/bedrock.go:23`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:567`](external_services.md)) |
```

**CITED CONTENT**

```
    20  	// a 400 "temperature is deprecated for this model". The author
    21  	// path therefore routes through SingleShotNoSampling (see
    22  	// authorClientAdapter below) so no sampling knob reaches Bedrock.
    23  	DefaultAuthorModelID = "eu.anthropic.claude-opus-4-8"
    24  
    25  	// DefaultGraderModelID is the canonical Sonnet grader, held fixed
    26  	// so the ≥ 90 floor means the same thing across runs. Sonnet 4.6
```

## 07-019
- **id**: `B07-019`
- **corpus site**: `corpus/architecture/ai_architecture.md:92-92` (table-row)
- **citation**: `internal/askengine/bedrock.go:25`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/askengine/bedrock.go`  (787 lines)

**CLAIMING UNIT**

```md
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:567`](external_services.md)) |
```

**CITED CONTENT**

```
    22  // Defaults for the Bedrock client. Both can be overridden via the
    23  // ASK_MODEL_ID and AWS_REGION environment variables.
    24  const (
    25  	DefaultModelID   = "eu.anthropic.claude-sonnet-4-6"
    26  	DefaultRegion    = "eu-west-1"
    27  	DefaultMaxTokens = 4096
    28  )
```

## 07-020
- **id**: `B07-020`
- **corpus site**: `corpus/architecture/ai_architecture.md:92-92` (table-row)
- **citation**: `internal/jobsimulation/agent/report_agent.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/agent/report_agent.go`  (653 lines)

**CLAIMING UNIT**

```md
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:567`](external_services.md)) |
```

**CITED CONTENT**

```
    28  	// transient draft id that started returning 400 "model identifier
    29  	// invalid" once the profile stabilised. Verified 2026-04-27 against
    30  	// `aws bedrock-runtime invoke-model` with prod IAM grants.
    31  	defaultAgentModel = "eu.anthropic.claude-sonnet-4-6"
    32  
    33  	maxAgentTurns = 10
    34  	maxTokens     = 16_000
```

## 07-021
- **id**: `B07-021`
- **corpus site**: `corpus/architecture/ai_architecture.md:92-92` (table-row)
- **citation**: `internal/coursebuilder/bedrock.go:29`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:567`](external_services.md)) |
```

**CITED CONTENT**

```
    26  	// so the ≥ 90 floor means the same thing across runs. Sonnet 4.6
    27  	// STILL accepts sampling params, so the grader path keeps calling
    28  	// SingleShot with temperature=0 for deterministic scoring.
    29  	DefaultGraderModelID = "eu.anthropic.claude-sonnet-4-6"
    30  
    31  	// AuthorModelEnv overrides DefaultAuthorModelID at process start.
    32  	// The grader has its own ISOLATED env var so a single accidental
```

## 07-022
- **id**: `B07-022`
- **corpus site**: `corpus/architecture/ai_architecture.md:92-92` (table-row)
- **citation**: `internal/ai/openai/config.go:8-26`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/ai/openai/config.go`  (39 lines)

**CLAIMING UNIT**

```md
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:567`](external_services.md)) |
```

**CITED CONTENT**

```
     5  	"github.com/openai/openai-go/v3"
     6  )
     7  
     8  const (
     9  	// 4.1
    10  	GPT4Dot1     ai.Model = ai.Model(openai.ChatModelGPT4_1)
    11  	GPT4Dot1Mini ai.Model = ai.Model(openai.ChatModelGPT4_1Mini)
    12  	// o3
    13  	O3     ai.Model = ai.Model(openai.ChatModelO3)
    14  	O4Mini ai.Model = ai.Model(openai.ChatModelO4Mini)
    15  	// 5
    16  	GPT5     ai.Model = ai.Model(openai.ChatModelGPT5)
    17  	GPT5Mini ai.Model = ai.Model(openai.ChatModelGPT5Mini)
    18  	GPT5Nano ai.Model = ai.Model(openai.ChatModelGPT5Nano)
    19  	// 5.1
    20  	GPT5_1 ai.Model = ai.Model(openai.ChatModelGPT5_1)
    21  	// 5.2
    22  	GPT5_2 ai.Model = ai.Model(openai.ChatModelGPT5_2)
    23  	// 5.4
    24  	GPT5_4      ai.Model = ai.Model(openai.ChatModelGPT5_4)
    25  	GPT5_4_Mini ai.Model = ai.Model(openai.ChatModelGPT5_4Mini)
    26  )
    27  
    28  const (
    29  	gptDefaultModel               = GPT5_4
```

## 07-023
- **id**: `B07-023`
- **corpus site**: `corpus/architecture/ai_architecture.md:94-94` (table-row)
- **citation**: `app/internal/ai/speech.go:9-12`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/ai/speech.go`  (91 lines)

**CLAIMING UNIT**

```md
| **Speech** | GPT-4o Mini TTS (`gpt-4o-mini-tts`) — **the whole set**, and it is also `DefaultModel` | Azure voice client (`CreateSpeech` is Azure-only). Enumerated from `app/internal/ai/speech.go:9-12` **@ `app` `ad9f3c49`**, which is the entire `SpeechModel` const block. **This row also listed "TTS v2 HD, TTS v2" until M257x iter-115.** Those two consts are real, but they belong to the **standalone** `github.com/anthropos-work/ai` module (`speech.go:12-13` @ `v1.40.1`, `TTSV2 = "tts-2"` / `TTSV2HD = "tts-2-hd"`) — a module **no repo a stack builds requires** since the fold at `1e457fa70` (see the **Unified AI Library** section below, which carries that fold); they were dropped in the fold and no caller in the clone set ever referenced them. Both live call sites use `ai.GPT4oMiniTTSS` |
```

**CITED CONTENT**

```
     6  type SpeechVoice string
     7  type SpeechResponseFormat string
     8  
     9  const (
    10  	GPT4oMiniTTSS SpeechModel = "gpt-4o-mini-tts"
    11  	DefaultModel  SpeechModel = GPT4oMiniTTSS
    12  )
    13  
    14  const (
    15  	VoiceAlloy   SpeechVoice = "alloy"
```

## 07-024
- **id**: `B07-024`
- **corpus site**: `corpus/architecture/ai_architecture.md:114-114` (paragraph)
- **citation**: `app/internal/jobsimulation/ai/ai.go:267`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/ai/ai.go`  (355 lines)

**CLAIMING UNIT**

```md
> **Vendor selection/fallback and cost tracking are NOT in the `ai` library** — they live in the consuming services: selection/fallback in each consumer's own wrapper — in `app` @ `5ba17044` that is **`app/internal/jobsimulation/ai/ai.go:267,344`** and **`app/internal/skillerai/ai.go:347`**, *not* a bare `app/internal/ai/ai.go`. **NB that file DOES exist** at `app` `ad9f3c49` — it is the folded `ai.AI` interface (`1e457fa70`), and this parenthetical asserted *"no such file"* until M257x iter-108; what remains true is only that **vendor selection does not live there**, which is the point the sentence is making. The Azure client defaults to EU and swaps to US on the PostHog flag `flag_use_azure_us`; direct OpenAI is the retry target on HTTP 429; Anthropic is always Bedrock `eu-west-1`. **These are three independent mechanisms, not rungs of an ordered ladder** — the ⚠️ under *Provider Routing Strategy* at the head of this file (`:15-17`) retracts that ladder, and this line went on publishing it 68 lines below, in `→` form, until M257x iter-48. Cost tracking is in `app/internal/aiusage/ai_usage.go` (fed by `Event_AiUsage` over Redis Streams). See [Shared Libraries → ai](shared_libraries.md#ai).
```

**CITED CONTENT**

```
   264  		client := a.azureClientEu
   265  		isAzureUsFlagEnabled, err := a.posthogClient.IsFeatureEnabled(
   266  			nil,
   267  			"flag_use_azure_us",
   268  			fflags.WithOnlyEvaluateLocally(true),
   269  		)
   270  		if err != nil {
```

## 07-025
- **id**: `B07-025`
- **corpus site**: `corpus/architecture/ai_architecture.md:114-114` (paragraph)
- **citation**: `app/internal/skillerai/ai.go:347`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skillerai/ai.go`  (372 lines)

**CLAIMING UNIT**

```md
> **Vendor selection/fallback and cost tracking are NOT in the `ai` library** — they live in the consuming services: selection/fallback in each consumer's own wrapper — in `app` @ `5ba17044` that is **`app/internal/jobsimulation/ai/ai.go:267,344`** and **`app/internal/skillerai/ai.go:347`**, *not* a bare `app/internal/ai/ai.go`. **NB that file DOES exist** at `app` `ad9f3c49` — it is the folded `ai.AI` interface (`1e457fa70`), and this parenthetical asserted *"no such file"* until M257x iter-108; what remains true is only that **vendor selection does not live there**, which is the point the sentence is making. The Azure client defaults to EU and swaps to US on the PostHog flag `flag_use_azure_us`; direct OpenAI is the retry target on HTTP 429; Anthropic is always Bedrock `eu-west-1`. **These are three independent mechanisms, not rungs of an ordered ladder** — the ⚠️ under *Provider Routing Strategy* at the head of this file (`:15-17`) retracts that ladder, and this line went on publishing it 68 lines below, in `→` form, until M257x iter-48. Cost tracking is in `app/internal/aiusage/ai_usage.go` (fed by `Event_AiUsage` over Redis Streams). See [Shared Libraries → ai](shared_libraries.md#ai).
```

**CITED CONTENT**

```
   344  		}
   345  		isAzureUsFlagEnabled, err := a.posthogClient.IsFeatureEnabled(
   346  			nil,
   347  			"flag_use_azure_us",
   348  			fflags.WithOnlyEvaluateLocally(true),
   349  		)
   350  		if err != nil {
```

## 07-026
- **id**: `B07-026`
- **corpus site**: `corpus/architecture/ai_architecture.md:130-133` (paragraph)
- **citation**: `studio/configs/production_config.ini:26-36`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/configs/production_config.ini`  (46 lines)

**CLAIMING UNIT**

```md
The Python generation pipeline uses configurable model slots, one per `{MODE}_AI_{BRANCH}_MODEL` key of
`configs/{env}_config.ini` (`service, model, thinking`). Measured at `app` HEAD against the shipping
`studio/configs/production_config.ini:26-36` — **the `stable` and `experimental` branches are currently
identical**, and `development_config.ini:26-36` is identical to both:
```

**CITED CONTENT**

```
    23  [SERVICES]
    24  
    25  # TARGET SERVICE: openai, azure, anthropic # TARGET MODEL: gpt-4o, claude-3-5-sonnet-20241022 # THINKING: none, low, medium, high (optional, only for supported models)
    26  FAST_AI_STABLE_MODEL = azure, gpt-5-mini, none
    27  STRICT_AI_STABLE_MODEL = azure, gpt-5-mini, none
    28  EXECUTION_AI_STABLE_MODEL = azure, gpt-5.4, none
    29  CREATIVE_AI_STABLE_MODEL = azure, gpt-5.4, low
    30  REASONING_AI_STABLE_MODEL = azure, gpt-5.4, medium
    31  
    32  FAST_AI_EXPERIMENTAL_MODEL = azure, gpt-5-mini, none
    33  STRICT_AI_EXPERIMENTAL_MODEL = azure, gpt-5-mini, none
    34  EXECUTION_AI_EXPERIMENTAL_MODEL = azure, gpt-5.4, none
    35  CREATIVE_AI_EXPERIMENTAL_MODEL = azure, gpt-5.4, low
    36  REASONING_AI_EXPERIMENTAL_MODEL = azure, gpt-5.4, medium
    37  
    38  
    39  [SIMULATIONS]
```

## 07-027
- **id**: `B07-027`
- **corpus site**: `corpus/architecture/ai_architecture.md:143-151` (paragraph)
- **citation**: `studio/services/ai.py:356`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/services/ai.py`  (752 lines)

**CLAIMING UNIT**

```md
> **Do not read the slot table off `configs/config_template.ini`.** That file is a non-shipping scaffold
> and still carries the older `gpt-4.1-mini` / `gpt-4.1` / `gpt-4o` / `o3` stable column; the corpus
> asserted a hybrid of it for several releases. **`gpt-5.2` appears in no studio config at all** (only as a
> pricing entry in `studio/services/ai.py:356,508`), and `gpt-4o` appears in no `*_MODEL` slot of any
> **shipping** studio config — it *is* in two slots of the non-shipping scaffold this blockquote opens by
> telling you not to read: `configs/config_template.ini:39` (`EXECUTION_AI_STABLE_MODEL = azure, gpt-4o,
> none`) and `:40` (`CREATIVE_AI_STABLE_MODEL`). The unqualified form contradicted this same blockquote two
> lines earlier, which already says the template *"still carries"* `gpt-4o`; corrected M257x iter-46.
> Agrees with [`studio-room.md`](../services/studio-room.md#ai-service-configuration).
```

**CITED CONTENT**

```
   353                  'completion': 10.00/1000000,
   354                  'type': 'thinking',
   355              },
   356              'gpt-5.2': {
   357                  'prompt': 1.75/1000000,
   358                  'completion': 14.00/1000000,
   359                  'type': 'thinking',
```

## 07-028
- **id**: `B07-028`
- **corpus site**: `corpus/architecture/ai_architecture.md:220-231` (paragraph)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:1079-1085`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
**Engine choice is per SEQUENCE, from the CMS `voice_engine` field** — a 4-member enum on the authored
simulation (`app/internal/cms/directus/collections/jobsimulation.go:1079-1085`); when it is nil the content
layer supplies `gptrealtime` — **`app/internal/cms/directus/collections/jobsimulation.go:1594-1597` @ `app`
`ad9f3c49`**, the whole of `func voiceEngineFromDirectus` down to the closing brace of its nil branch
(`:1595-1596` is `if directusVoiceEngine == nil { return simulation.SimulationVoiceEngineGptrealtime }`).
**Pinned in full, and to `:1597` rather than `:1600`, deliberately (M257x iter-115):** this is one half of a
same-fact-different-pin pair with the **ElevenLabs** bullet under *External Dependencies* in
[`jobsimulation.md`](../services/jobsimulation.md), whose half spelled
the path `cms/directus/collections/…` — a path that exists in **no** clone. Fixing one side and leaving the
other is how the corpus acquires a self-contradiction, so both halves now carry the same path and the same
range. **ElevenLabs remains the active default** for the call/reply
pipeline and transcript improvement, so it is not yet fully replaced.
```

**CITED CONTENT**

```
  1076  	Tasks                   []TaskTranslation `json:"tasks"`
  1077  }
  1078  
  1079  type SimulationVoiceEngine string
  1080  
  1081  const (
  1082  	SimulationVoiceEngineGptrealtime        SimulationVoiceEngine = "gptrealtime"
  1083  	SimulationVoiceEngineElevenlabs         SimulationVoiceEngine = "elevenlabs"
  1084  	SimulationVoiceEngineLivekitgptrealtime SimulationVoiceEngine = "livekitgptrealtime"
  1085  	SimulationVoiceEngineLivekitchain       SimulationVoiceEngine = "livekitchain"
  1086  )
  1087  
  1088  func (s JobSimulation) ToDomain(dc *directus.Client, language *content.Language) *simulation.JobSimulation {
```

## 07-029
- **id**: `B07-029`
- **corpus site**: `corpus/architecture/ai_architecture.md:220-231` (paragraph)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:1594-1597`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
**Engine choice is per SEQUENCE, from the CMS `voice_engine` field** — a 4-member enum on the authored
simulation (`app/internal/cms/directus/collections/jobsimulation.go:1079-1085`); when it is nil the content
layer supplies `gptrealtime` — **`app/internal/cms/directus/collections/jobsimulation.go:1594-1597` @ `app`
`ad9f3c49`**, the whole of `func voiceEngineFromDirectus` down to the closing brace of its nil branch
(`:1595-1596` is `if directusVoiceEngine == nil { return simulation.SimulationVoiceEngineGptrealtime }`).
**Pinned in full, and to `:1597` rather than `:1600`, deliberately (M257x iter-115):** this is one half of a
same-fact-different-pin pair with the **ElevenLabs** bullet under *External Dependencies* in
[`jobsimulation.md`](../services/jobsimulation.md), whose half spelled
the path `cms/directus/collections/…` — a path that exists in **no** clone. Fixing one side and leaving the
other is how the corpus acquires a self-contradiction, so both halves now carry the same path and the same
range. **ElevenLabs remains the active default** for the call/reply
pipeline and transcript improvement, so it is not yet fully replaced.
```

**CITED CONTENT**

```
  1591  	return skills
  1592  }
  1593  
  1594  func voiceEngineFromDirectus(directusVoiceEngine *SimulationVoiceEngine) simulation.SimulationVoiceEngine {
  1595  	if directusVoiceEngine == nil {
  1596  		return simulation.SimulationVoiceEngineGptrealtime
  1597  	}
  1598  
  1599  	switch *directusVoiceEngine {
  1600  	case SimulationVoiceEngineGptrealtime:
```

## 07-030
- **id**: `B07-030`
- **corpus site**: `corpus/architecture/ai_architecture.md:264-264` (bullet)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:983-990`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
2. **Route**: Selected model from the CMS `ai_model` / `ai_vendor` fields per sequence (e.g. `gpt-5`, `gpt-4.1`, `anthropic-45-sonnet-aws` — the enum is `app/internal/cms/directus/collections/jobsimulation.go:983-990`). **There is no single "default model", and `gpt-5` is not a default anywhere.** Three distinct defaults apply at three different points — see below
```

**CITED CONTENT**

```
   980  	Anthropic35SonnetAws AIModel = "anthropic-35-sonnet-aws"
   981  	Anthropic37SonnetAws AIModel = "anthropic-37-sonnet-aws"
   982  	Anthropic4SonnetAws  AIModel = "anthropic-4-sonnet-aws"
   983  	Anthropic45SonnetAws AIModel = "anthropic-45-sonnet-aws"
   984  	Anthropic35Sonnet    AIModel = "anthropic-35-sonnet"
   985  	GptFourPointOne      AIModel = "gpt-4.1"
   986  	GptFourPointOneMini  AIModel = "gpt-4.1-mini"
   987  	Gpt5                 AIModel = "gpt-5"
   988  	Gpt5Mini             AIModel = "gpt-5-mini"
   989  	Gpt5Nano             AIModel = "gpt-5-nano"
   990  	Gpt5_1               AIModel = "gpt-5.1"
   991  )
   992  
   993  type Skills struct {
```

## 07-031
- **id**: `B07-031`
- **corpus site**: `corpus/architecture/ai_architecture.md:277-277` (table-row)
- **citation**: `app/internal/cms/directus/collections/jobsimulation.go:1297`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/directus/collections/jobsimulation.go`  (1742 lines)

**CLAIMING UNIT**

```md
| **Content side** — the CMS `ai_model` / `ai_vendor` field is left unset on a sequence | `openai` | **`gpt-5.1`** | `app/internal/cms/directus/collections/jobsimulation.go:1297` (model), `:1302` (vendor) |
```

**CITED CONTENT**

```
  1294  			})
  1295  		}
  1296  
  1297  		aiModel := simulation.Gpt5Point1
  1298  		if seq.AIModel != nil {
  1299  			aiModel = simulation.SimulationAIModel(*seq.AIModel)
  1300  		}
```

## 07-032
- **id**: `B07-032`
- **corpus site**: `corpus/architecture/ai_architecture.md:278-278` (table-row)
- **citation**: `app/internal/jobsimulation/simulator/ai/ai.go:65-66`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/simulator/ai/ai.go`  (166 lines)

**CLAIMING UNIT**

```md
| **Runtime routing** — `GetAIVendorAndModel` gets a model string it does not recognise | as selected | **`gpt-4.1`** | `app/internal/jobsimulation/simulator/ai/ai.go:65-66` (OpenAI arm), `:82-83` (Azure arm), `:126-127` (unmatched-vendor arm) |
```

**CITED CONTENT**

```
    62  			aiModel = openai.GPT4Dot1
    63  		case simulation.GptFourPointOneMini:
    64  			aiModel = openai.GPT4Dot1Mini
    65  		default:
    66  			aiModel = openai.GPT4Dot1
    67  		}
    68  	// Azure
    69  	case simulation.Azure:
```

## 07-033
- **id**: `B07-033`
- **corpus site**: `corpus/architecture/ai_architecture.md:302-307` (bullet)
- **citation**: `internal/jobsimulation/simulator/validation/v3/validator/validator.go:43`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/simulator/validation/v3/validator/validator.go`  (1202 lines)

**CLAIMING UNIT**

```md
- Each skill has multiple criteria with binary checks (pass/fail), and **most of those verdicts are judged
  by an LLM**. The dispatch is a hardcoded switch, not the `checkerEngines` map — that map is stored and
  **never read** (`internal/jobsimulation/simulator/validation/v3/validator/validator.go:43,60-61,595`),
  so do not cite it as the mechanism. `basevalidator/criterion.go:127` routes LLM checks to `validateLLM`
  → `NewLLMBulkChecker(c.logger)` (`:428`), which sends `basevalidator/templates/checkValidationBulk.tmpl`
  at temperature 0.0 and reads back `{"check_id", "feedback", "success"}`
```

**CITED CONTENT**

```
    40  	logger         *slog.Logger
    41  	storage        storagev1.Service
    42  	manager        manager.Manager
    43  	checkerEngines map[check.Engine]basevalidator.BulkChecker
    44  	skillReader    basevalidator.SkillReader
    45  	aiManager      *localAi.AIManager
    46  }
```

## 07-034
- **id**: `B07-034`
- **corpus site**: `corpus/architecture/ai_architecture.md:314-319` (bullet)
- **citation**: `app/internal/skill/skill.go:617-623`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skill/skill.go`  (707 lines)

**CLAIMING UNIT**

```md
- **There is no 60/65/75/85/95 threshold ladder.** The corpus asserted one for several releases; it does
  not exist in `app`, `cms`, `jobsimulation` or `next-web-app`. The real conversion is
  `calculateCompetencyLevelScore` (`v3/validator/skills.go:40-51`): `20` when `score < 60 && isPassed`,
  `100` at `>= 100`, else `max(0, score*2-100)` — and it carries a `// TODO fix this formula` comment.
  The 0-100 ↔ N-level mapping is a plain division (`app/internal/skill/skill.go:617-623`
  `convertLevelTo100`; frontend `packages/ui/src/Competency/CompetencyReadLevel.tsx:18`)
```

**CITED CONTENT**

```
   614  	return res.Items[0], nil
   615  }
   616  
   617  func (s *SkillManager) convertLevelTo100(score *int, levelsCount int) *int {
   618  	if score == nil {
   619  		return nil
   620  	}
   621  	normalizedScore := 100 * (*score) / levelsCount
   622  	return &normalizedScore
   623  }
   624  
   625  func (s *SkillManager) prepareAddUserSkillsInput(ctx context.Context, userId uuid.UUID, skillId taxonomy.NodeID, experienceId uuid.UUID, experienceType enum.ExperienceType) (repository.AddUserSkillsInput, error) {
   626  	skillInput := repository.AddUserSkillsInput{
```

## 07-035
- **id**: `B07-035`
- **corpus site**: `corpus/architecture/alignment_testing.md:200-200` (table-row)
- **citation**: `gate.sh:69`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/clerkenstein/alignment/scripts/gate.sh`  (88 lines)

**CLAIMING UNIT**

```md
| **What the GATE actually enforces** (`--if-declared`) | ⚠ **Not the same as the bare command — do not conflate them.** `gate.sh:69` calls `alignctl dna coverage --dna "$base/$dna" --if-declared` (its rationale is the comment block at `:58-68`; `:61` — where this used to point — is inside that comment). That flag downgrades **exactly one** case — *"this DNA declares no `consumed_surface` at all"* — from **exit 2** to a **loud warning, exit 0**. A DNA that **does** declare a surface and leaves an endpoint **uncovered** still **fails the gate, exit 2, before a single gene is scored.** So: **a declared hole is fenced; an undeclared surface is only warned about.** The flag exists because a deployment/injection DNA has no HTTP surface to declare, and a hard stop there would be noise. |
```

**CITED CONTENT**

```
    66  # deployment/injection DNA has no HTTP surface to declare). A DNA that DOES declare one and leaves an
    67  # endpoint uncovered fails here — exit 2, before a single gene is scored.
    68  echo "==> capability-coverage check ($dna)"
    69  "$alignctl" dna coverage --dna "$base/$dna" --if-declared
    70  
    71  echo "==> alignment gate ($dna: overall >= ${gate_overall}%, critical >= ${gate_critical}%)"
    72  # `set -e` would swallow the distinction, so take the code explicitly and re-raise it with a verdict.
```

## 07-036
- **id**: `B07-036`
- **corpus site**: `corpus/architecture/alignment_testing.md:358-378` (bullet)
- **citation**: `alignment/cmd/alignctl/run.go:133-136`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/alignment/cmd/alignctl/run.go`  (155 lines)

**CLAIMING UNIT**

```md
- **M2c (`@clerk/express` backend session verification)** exercises the framework a **third** time, on the
  Node backend surface: a *third* DNA — `clerk-express-1` (**5 capabilities / 13 genes**) — with its own runner (`expressrun`)
  and goldens, scored by the same `alignctl` to the same gate. Its runner drives the **genuine
  `@clerk/express`/`@clerk/backend` SDK** (the *verify-against-the-real-library* discipline, the same one
  `clerk-webhook/` uses with `svix`) rather than a reimplementation — so the score measures whether the real
  SDK accepts Clerkenstein's tokens. It added an **additive RS256/JWKS** path beside the existing HS256
  seams (no migration; M1/M2 gates untouched).
  > ⚠ **This surface is DEPENDENCY-GATED, and it USED to be silently UNMEASURED. `TEST-M219-expressrun-dep-gate`
  > is RESOLVED — the behaviour below is historical.** The runner needs `@clerk/express` `node_modules` to
  > build. Without them it once exited **rc=2 — indistinguishable from a regression** — and *nothing in the
  > tooling treated that as a failure*, so on a box lacking the Node modules the gate silently contributed
  > **nothing** while summaries reported "all five surfaces at 100%". The M218 harden pass could re-measure
  > only **4 of the 5** surfaces for this reason (reproduced identically at the pre-pass baseline ⇒
  > pre-existing, not a regression).
  >
  > **Fixed at M219, and the fix was to make the two outcomes different numbers.** `alignment/cmd/alignctl/run.go:133-136`
  > declares `ExitRegressed = 2` and **`ExitUnmeasurable = 3`**; `unmeasurable()` (`:139-153`) returns **rc=3**
  > and prints a boxed banner — *"UNMEASURABLE — the runner could not execute. THIS IS NOT A PASSING SCORE …
  > Do NOT record this run as a pass."* **An absent score is not a passing score**, and it no longer wears a
  > regression's exit code. (Corrected M257x iter-85: this passage described the pre-M219 behaviour in the
  > present tense, and told a reader to read a `2` as a missing Node module.)
```

**CITED CONTENT**

```
   130  //
   131  // THE RULE THIS ENCODES: **absence of a score is not a passing score.** An unmeasurable surface must be
   132  // impossible to mistake for a measured one — so it gets its own code, and a banner that says so.
   133  const (
   134  	ExitRegressed    = 2
   135  	ExitUnmeasurable = 3
   136  )
   137  
   138  // unmeasurable fails LOUD. It never returns a score, and it never returns ExitRegressed.
   139  func unmeasurable(target, runner string, err error) int {
```

## 07-037
- **id**: `B07-037`
- **corpus site**: `corpus/architecture/alignment_testing.md:475-486` (bullet)
- **citation**: `stack-seeding/dna/snapshot.go:62`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/dna/snapshot.go`  (490 lines)

**CLAIMING UNIT**

```md
- a **snapshot-fidelity gene class** (`stack-seeding/dna/snapshot.go:62`) — **six** two-sided operators over a
  `FidelityProbe` (the replayed stack) compared to the captured manifest: **`snapshot-row-count`**
  (source-vs-replay parity), **`snapshot-structural`** (every captured column present after replay),
  **`snapshot-referential`** (the captured surface is referentially closed — every FK's parent table is in the
  captured set), **`snapshot-embedding-dim`**
  (pgvector columns replayed at the captured dimension — the index was rebuilt, the vectors must carry the same
  width), **`snapshot-public-only`** (the **provenance gene** — zero tenant-scoped rows after replay, the
  firewall's measured counterpart), and — added at M23, which is why the count reads five in older passes —
  **`snapshot-cross-surface-closure`**, closure that spans *two* surfaces: every taxonomy node-id the replayed
  content references must resolve to a skill present in the replayed taxonomy. `snapshot-referential` works
  **within** one surface and cannot see that dangle. A snapshot gene names **snapshot** operators; a structural
  gene names **structural** operators — `Validate` rejects a cross-wire so the two classes never mix.
```

**CITED CONTENT**

```
    59  // validSnapshotOperators is the set of snapshot-operator names a snapshot gene
    60  // may name. Kept separate from validOperators so a structural gene cannot name a
    61  // snapshot operator and vice-versa.
    62  var validSnapshotOperators = map[string]bool{
    63  	OpSnapshotRowCount:            true,
    64  	OpSnapshotStructural:          true,
    65  	OpSnapshotReferential:         true,
```

## 07-038
- **id**: `B07-038`
- **corpus site**: `corpus/architecture/architecture_overview.md:3-3` (paragraph)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
> **⚠️ Router status, two states (v2.8 M257x).** Platform `b56d731`+`360efd4` (merged **`2adcf71`**, 2026-07-31) **deleted the Cosmo Router from local dev** — no `graphql` compose service, no `repos.yml` entry — and re-pointed the frontends at **`backend` directly, `http://localhost:8082/graphql/query`**. **There is no `:5050` on a local stack.** In *production* the router is still declared (`graphql-wundergraph/terraform/main.tf:20` `= 1`), though **the repo is ARCHIVED on GitHub (2026-07-30)**. And the supergraph is **ONE** subgraph — `backend` — since `915da06` (2026-07-29). The fenced source of truth is [`platform-migration-status.md`](./platform-migration-status.md).
```

**CITED CONTENT**

```
    17    tags                           = var.tags
    18    aws_region                     = var.aws_region
    19    project                        = local.project
    20    service_desired_count          = 1
    21    service_cpu                    = local.service_cpu
    22    service_memory                 = local.service_memory
    23    service_port                   = local.port
```

## 07-039
- **id**: `B07-039`
- **corpus site**: `corpus/architecture/architecture_overview.md:20-20` (bullet)
- **citation**: `jobsimulation/terraform/main.tf:15-22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

**CLAIMING UNIT**

```md
    *   **Jobsimulation**: runs realistic AI-powered job scenarios with voice, chat, code, and document tasks. (It *runs* the simulation; the simulation *definition* is content owned by the cms domain. **Merged into `app`** — "jobsim-in-app"; **the repo's archive state is not visible to this corpus** — this line asserted a GitHub archive on 2026-07-31, which `origin/main`'s four **2026-08-04** commits (merged PR #439 among them) contradict, an archived GitHub repo being read-only; report both, assert neither — and **M810 has landed for the production ECS service**: `6092c6d2` deleted the `module "jobsimulation"` block outright, destroying the ECS service, task definition and ECR repository, so `service_desired_count` no longer appears in the file at all (`jobsimulation/terraform/main.tf:15-22`). What survives is the module's *other* ownership — the LiveKit and Chime recording buckets `backend` reads by literal name, the `/production/jobsimulation/*` SSM parameters and the atlas tracker; dropping the legacy `jobsimulation` schema is a separate, still-pending M810 step (`:24-40`).)
```

**CITED CONTENT**

```
    12    }
    13  }
    14  
    15  // Inspect the target database and load its state.
    16  // This is used to determine which migration to run.
    17  data "atlas_migration" "jobsimulation_migrations" {
    18    dir = "${path.module}/migrations?format=atlas"
    19    url = "${aws_ssm_parameter.db_connection.value}?search_path=jobsimulation"
    20  }
    21  
    22  // Sync the state of the target database with the migrations directory.
    23  resource "atlas_migration" "jobsimulation_migrations" {
    24    dir              = "${path.module}/migrations?format=atlas"
    25    version          = data.atlas_migration.jobsimulation_migrations.latest # Use latest to run all migrations
```

## 07-040
- **id**: `B07-040`
- **corpus site**: `corpus/architecture/architecture_overview.md:22-23` (bullet)
- **citation**: `app/internal/jobsimwiring/wiring.go:118`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
    *   **Roadrunner**: Judge0 code execution — `backend` reaches Judge0 directly
        (`app/internal/jobsimwiring/wiring.go:118` @ `app` `b948604`), so there is no hop and nothing left to start
```

**CITED CONTENT**

```
   115  	// object's physical address. The jobsim corpus (recordings, conversation clips, interaction
   116  	// audio + attachments, interview report CSVs) was written by the standalone jobsimulation
   117  	// service under "jobsimulation", so it must be read under "jobsimulation". Passing serviceName
   118  	// here ("backend") re-points every read of that historical corpus at a prefix that does not
   119  	// contain it — a silent 404, not an error. See internal/storagens.
   120  	storageV1Client := appstorage.NewClient(inAppStorage, storagens.JobSimulation).V1
   121  	// Judge0 sandbox runner (IN-PROCESS; replaces the removed roadrunner RPC edge — resync to jobsim main
```

## 07-041
- **id**: `B07-041`
- **corpus site**: `corpus/architecture/architecture_overview.md:25-35` (paragraph)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
    **Also domains inside Backend/App, and no longer even opt-in:** **Storage**, **Messenger** (Brevo
    email) and **CustomerIO Sync**. Platform `838d907` (merged **`0c91421`**, 2026-08-05) **deleted all
    three compose services**, with their ports and `depends_on` edges, and dropped `storage` +
    `messenger` from `repos.yml`. The `storage-legacy` / `messenger` / `customerio-sync` profiles are gone
    with them, and asking for one now exits 0 and starts only the always-on floor. `app`
    serves object storage in-process; messenger and customerio-sync ride in the same container but stay
    **OFF** on a developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose
    deliberately does not set (`docker-compose.yml:84-92` says why). Up to that commit the first two were
    kept startable "for rollback comparison"; that escape hatch is gone.
    Archived (removed from local orchestration): Chronos, Intelligence.
    Production-only: **db-backup** (scheduled PostgreSQL backups).
```

**CITED CONTENT**

```
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
    86        # stream or a timer — they send mail and rewrite Brevo contacts — so app gates them
    87        # behind MESSENGER_ENABLED / CUSTOMERIO_SYNC_ENABLED, which default to OFF on a
    88        # developer machine (ENVIRONMENT=development is what makes unset mean off).
    89        # Pinning them to `false` here would override .env and make opting in impossible
    90        # without editing this file. To exercise either one locally, set it in .env — and
    91        # know that messenger then attaches to the LIVE Redis consumer group and
    92        # customerio-sync writes real Brevo contacts.
    93        - SUPABASE_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    94        - COPILOT_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    95      networks:
```

## 07-042
- **id**: `B07-042`
- **corpus site**: `corpus/architecture/architecture_overview.md:221-221` (table-row)
- **citation**: `jobsimulation/terraform/main.tf:15-22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

**CLAIMING UNIT**

```md
| **Jobsimulation** | Merged into Backend/App ("jobsim-in-app") — session engine runs in `app`; the 23 run-state tables moved to `public`. **No local container**: `d11a403` deleted the compose service and the `repos.yml` entry, so at `0dab54d` there is nothing to start. **Prod teardown — M810 has LANDED for the ECS service**: `6092c6d2` deleted the `module "jobsimulation"` block, so the ECS service, task definition and ECR repository are destroyed (`jobsimulation/terraform/main.tf:15-22`). The module file survives owning only the LiveKit/Chime buckets, the SSM parameters and the atlas tracker; the legacy-schema drop is a separate, still-pending M810 step. **Do not read this row onto CMS** | [→](../services/jobsimulation.md) |
```

**CITED CONTENT**

```
    12    }
    13  }
    14  
    15  // Inspect the target database and load its state.
    16  // This is used to determine which migration to run.
    17  data "atlas_migration" "jobsimulation_migrations" {
    18    dir = "${path.module}/migrations?format=atlas"
    19    url = "${aws_ssm_parameter.db_connection.value}?search_path=jobsimulation"
    20  }
    21  
    22  // Sync the state of the target database with the migrations directory.
    23  resource "atlas_migration" "jobsimulation_migrations" {
    24    dir              = "${path.module}/migrations?format=atlas"
    25    version          = data.atlas_migration.jobsimulation_migrations.latest # Use latest to run all migrations
```

## 07-043
- **id**: `B07-043`
- **corpus site**: `corpus/architecture/architecture_overview.md:222-222` (table-row)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
| **CMS** | Merged into Backend/App ("cms-in-app v8.0", app v1.360.0) — content layer + Studio run in `app`; similarity/studio tables moved to `public`; supergraph **3→1** (the same commit, `915da06`, also deleted the `jobsimulation` subgraph — its own commit subject's "2→1" is wrong); the prod ECS module is **not** a settled rollback path — **report both, assert neither**: `cms/terraform/main.tf:39` still declares it at `service_desired_count = 0`, while `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted cms's build-production workflow under *"the cms ECR repository is decommissioned (M810)"*, naming M810's deletion of `module.cms_euwest1`. Whether that has been applied is **not visible to this corpus** — `infrastructure` has never been in any clone set. **No local container**: `d11a403` deleted the compose service and the `repos.yml` entry, and re-pointed `messenger`'s `CMS_RPC_ADDR` at `http://backend:8083` — **M809 has landed** (on **two** variables, this one and `JOBSIMULATION_RPC_ADDR`; not on all four — M257x iter-115). `838d907` (merged `0c91421`) then deleted the `messenger` service itself, so **no compose file sets `CMS_RPC_ADDR` at all** any more; prod teardown is **M810** | [→](../services/cms.md) |
```

**CITED CONTENT**

```
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 07-044
- **id**: `B07-044`
- **corpus site**: `corpus/architecture/architecture_overview.md:265-284` (bullet)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
*   **Synchronous**: Connect-RPC/HTTP endpoints — down to **one Connect-RPC edge on a local stack,
    `backend → sentinel`**. At platform `0c91421` that is the only cross-process **Connect-RPC** address,
    `AUTHORIZATION_ADDRESS=http://sentinel:8087`
    (`docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables**. **It is NOT the only service
    address compose sets, and not the only cross-process edge** — this passage previously said *"compose
    sets exactly one service address"*, which **is false** and is retracted (corrected M257x iter-102).
    The same `backend` block also sets `GOTENBERG_URL=http://gotenberg:3200` (`docker-compose.yml:57` — a
    second container on the **default** `core` profile at `:183`, reached over **plain HTTP**, not
    Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL`
    (`docker-compose.yml:59`) and `REDIS_ADDR` (`docker-compose.yml:66`). **The correctly-scoped form is
    this document's own local-stack diagram below** — *"the only cross-process **RPC** edge out of backend
    on a core stack"* — which was right while this line was wrong, 55 lines apart in one file.
    On the `*_RPC_ADDR` half: the `messenger` block was the last thing
    that set any (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_` — **all four read
    `http://backend:8083`, but `d11a403` moved only the MIDDLE TWO**: `CMS_RPC_ADDR` and
    `JOBSIMULATION_RPC_ADDR`. `BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held that value at
    `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from its introduction
    at `3e85fce` — it only ever moved ports, so there was nothing to re-point. Corrected M257x iter-115),
    and `838d907` deleted that service. The env-var *names* still exist
    in consumer code; no local compose file configures them
```

**CITED CONTENT**

```
    45        - .env
    46      environment:
    47        - AI_USAGE_STREAM=AI
    48        - AUTHORIZATION_ADDRESS=http://sentinel:8087
    49        - AWS_CHIME_SDK_REGION=eu-central-1
    50        - CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo
    51        - CMS_STREAM=cms
```
