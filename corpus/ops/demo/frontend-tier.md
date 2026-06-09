# The demo frontend tier — the UI of a demo stack

**Purpose.** Make a `demo-N` (or `dev-N`) stack **actually demoable**: bring up the user-facing **UI tier** —
**next-web-app** (the Workforce app) + **studio-desk** at offset ports, plus **ant-academy** natively — so a
stakeholder lands on a real, clickable, Clerk-free UI, not just a running backend. This is the v1.3b M19
deliverable that completes the [demo family](README.md): up → snapshot → seed → **see it in a browser** → down.

> **Read [`../safety.md`](../safety.md) first** for *why* this is safe, and [`../rosetta_demo.md`](../rosetta_demo.md)
> for the stack lifecycle this extends. This page is the **frontend-specific** "how the UI tier is built and run".

> **The hard line (non-negotiable).** **Zero platform-repo edits.** next-web-app, studio-desk, and ant-academy
> stay **byte-for-byte pristine** — their repos are used only as a Docker **build context** (their Dockerfiles
> consumed UNMODIFIED), and every per-demo difference rides a **gitignored** overlay (`.env.local`) or a
> tooling-owned file in `rosetta-extensions`. Nothing the demo tooling does touches a tracked platform file.

## What `/demo-up` brings up (UI tier)

| App | How it runs | Port (base + offset) | Auth in the demo |
|-----|-------------|----------------------|------------------|
| **next-web-app** (Workforce) | per-demo **Docker** image from the unmodified `Dockerfile.dev`, in the demo's `graphql` profile | **3000** + N×10000 | Clerk-free (Clerkenstein-minted pk baked into the bundle) |
| **studio-desk** | per-demo **Docker** image from the unmodified `Dockerfile.dev`, in the `graphql` profile | **9100** (frontend) + **9000** (backend), each + N×10000 | Clerk-free (minted pk as a build-arg) |
| **ant-academy** | **native** `next dev` (Vercel-native; not dockerized) | **3077** + N×10000 | Clerk-free via `BENCHMARK_VISUAL_BYPASS` (anonymous browse) |

Example: `demo-2` → next-web on `:23000`, studio-desk on `:29100`/`:29000`, ant-academy on `:23077`.

**Default-on, skippable.** The UI tier is built + brought up by default. `DEMO_NO_UI=1 /demo-up N` (or the
`--no-ui` equivalent) brings up a **backend-only** demo — no frontend build, no academy, and the verify net is
scoped so it never warns about the absent UI. Use it for a fast API-only stack or a RAM-tight box.

## Why per-demo builds (and the honest residual)

The frontends inline their backend/router URLs **and** the Clerk publishable key into the client bundle **at
build time** (empirically confirmed — the pk literal lives in `.next/static/chunks/*.js`). A demo runs on
**offset ports** with its **own minted pk**, so the bundle is demo-specific: each new `demo-N` needs its own
image. The tooling makes this cheap-where-it-can:

- **Built once per `demo-N`, then cached.** The build is **tag-guarded** (`docker image inspect demo-N-next-web`):
  a re-up of the same demo reuses the cached image in **seconds**. Only a **brand-new** `demo-N` (or a frontend
  code/dep change) pays the build.
- **The residual (honest):** a new `demo-N` costs **one ~3-minute, ~3.7 GB cached build per frontend**
  (next-web is the heavy one; studio-desk is light). That's the price of zero-platform-edit + per-stack pk/URL
  baking. *True* zero-rebuild would need runtime-configurable URLs + pk in the platform source — an **optional
  upstream PR you'd own**, explicitly **out of scope** here (it edits platform repos → forbidden). See §"What's
  out of scope".
- **Built serially, before `compose up`.** The two frontend builds run **one at a time, before** the stack
  starts — kept out of the parallel Go-service fan-out so the build RAM spike never overlaps anything else.
- **Non-fatal.** A frontend build failure **warns** but never aborts the backend bring-up; re-run to retry, or
  `DEMO_NO_UI=1` to skip.

## The 12 GB Docker-VM prerequisite

**Runtime is cheap** — measured **~0.66 GiB for BOTH stacks** (dev + demo, 27 containers). The only spike is the
**build**: a ~3.7 GB next-web compile. On an undersized Docker VM already holding the dev stack, that spike
**swap-thrashes** (the original "the build takes an hour" symptom was pure memory starvation, not a slow build).

> **Set the Docker Desktop VM to 12 GB / swap 3 GB** (Settings → Resources). `/demo-up` runs a **non-fatal
> pre-flight assert**: below 12 GB it prints a clear warning (raise the VM, or run `DEMO_NO_UI=1`) but continues
> — a smaller VM may still build fine if no other stack is up. Override the floor with `DEMO_VM_MIN_GIB=N`.

## How the pk + URLs are baked (zero platform edit)

| App | URLs | Clerk pk | Context trim |
|-----|------|----------|--------------|
| **next-web** | `--build-arg NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` / `_BACKEND_API_URL` / `_HOSTING_URL` (offset) — ARGs the Dockerfile already declares | **no pk ARG exists** → dropped into a **gitignored `apps/web/.env.local`** in the build context, read by `next build`, removed by a trap after | the repo ships **no** `.dockerignore`, so a **tooling-owned** one (`rosetta-extensions/demo-stack/frontend/next-web.dockerignore`) is applied **transiently** (never clobbers a repo one; trap-removed) to trim the 2.8 GB context (2.5 GB `node_modules`) to <100 MB |
| **studio-desk** | `--build-arg VITE_GRAPHQL_ENDPOINT` (offset) | **`VITE_CLERK_PUBLISHABLE_KEY` IS a declared ARG** → passed straight as a build-arg | the repo **already ships** a `.dockerignore` excluding `node_modules`/`dist`/`.git` — left untouched |

The minted pk comes from the demo's Clerkenstein injection (`inject.py` mints `pk_test_<base64(fapi-host$)>` and
prints it); the build bakes that exact value, so the browser SDK talks to the demo's fake FAPI, never real Clerk.

> **The cleanup is `RETURN`-scoped, so it fires on the failure/abort path too.** The trap that removes the pk
> `.env.local` overlay and the transient `.dockerignore` is bound to the build function's `RETURN`, not its
> success — so a **failed** (or aborted) `docker build` leaves the repo just as byte-clean as a successful one.
> The load-bearing proof is a guard test that stands up a real git repo as the build context and asserts
> `git status` stays empty after the (stubbed) build, plus a `git check-ignore` fence that the pk overlay path
> is covered by a `.gitignore` rule (so it can never be tracked even mid-build). _(M19 harden — surfaced when
> the failed-build and real-git-status invariants were pinned: `test_next_web_failed_build_still_removes_*`,
> `TestZeroPlatformRepoEdit` in `demo-stack/tests/test_frontend_build.py`.)_

## ant-academy — native, Clerk-free, with a documented fallback

ant-academy is **Vercel-native** (not in docker-compose) and depends only on Clerk at runtime. `/demo-up`
launches it natively on `:3077+offset` **Clerk-free** using the repo's own `BENCHMARK_VISUAL_BYPASS` (a dev-only,
`NODE_ENV=development` flag that opens `/` and `/chapters/*` to anonymous traffic), paired with
`REQUIRE_ORGANIZATION_MEMBERSHIP=0` to skip the org gate. The per-demo env is a **gitignored `code/.env.local`**
overlay (zero academy-repo edits).

**Default-on + non-fatal + degrades to a documented step.** If the academy clone isn't present, its deps aren't
installed (`npm install` needs the team-issued **Font Awesome Pro** token), or Node < 22, the tool prints the
exact manual commands and continues — it never aborts a good demo bring-up:

```bash
cd stack-dev/ant-academy/code
cp .env.example .env.local                 # gitignored; keeps the repo clean
#   set REQUIRE_ORGANIZATION_MEMBERSHIP=0, reuse platform/.env's Clerk keys, add FONTAWESOME_NPM_AUTH_TOKEN
npm install
BENCHMARK_VISUAL_BYPASS=1 npm run dev -- --port 23077   # demo-2: Clerk-free anonymous browse
```

`/demo-down N` stops the native academy first (it's a process, not a container, so `compose down` can't reach
it). See [`../../services/ant-academy.md`](../../services/ant-academy.md) for the full app picture.

## Verification covers the UI tier

The M18 [verification net](../verification.md) now covers the frontends: `stack-verify`'s service registry
includes **next-web-app (:3000)** + **studio-desk (:9100)**, which offset + project-rewrite like every other
service. The bring-up-tail auto-verify is **scoped to the services it started** — so a UI-on demo verifies the
frontends (an HTTP probe; a Clerk-free login redirect is a healthy 2xx/3xx/4xx), and a `--no-ui` demo scopes
them out and never false-`down`s an absent frontend.

## Where the tooling lives

All of the above is `rosetta-extensions` tooling, authored + tagged in the authoring copy and consumed per-stack
at the pinned tag (`stack-demo/rosetta-extensions @ dress-rehearsal-m19`):

- `stack-injection/gen_injected_override.py` — appends the two frontends to the injected override (offset
  `ports:!override`, `image: demo-N-*` + `build:!reset null` + `pull_policy:never`, `mem_limit:1g`,
  `profiles:!override [graphql]`); `--no-ui` clears the tier.
- `demo-stack/up-injected.sh` — the per-demo serial-before-up frontend build (offset URLs + minted pk +
  tag-guard), the 12 GB VM pre-flight, the `--no-ui` (`DEMO_NO_UI`) escape, the scoped verify.
- `demo-stack/frontend/next-web.dockerignore` — the tooling-owned context trim for next-web.
- `demo-stack/ant-academy.sh` — the native academy launcher / stopper / documented fallback.

## What's out of scope (the user-owned follow-up)

**True zero-rebuild** — one frontend image that serves every stack with the port/pk switched at *runtime* —
would require **platform-source changes** (runtime rewrites in `next.config.mjs`, an absolute internal origin for
SSR `fetch` in `server.graphql.ts`, a `window.__ENV` shim + explicit `publishableKey` on `<ClerkProvider>`,
optionally `output:'standalone'`). Those are real platform edits with PR/review/prod risk — **forbidden** for the
demo tooling to make locally. It's documented here as an **optional upstream PR you own** (the v1.4 deploy-CI
precedent), **not built** in M19. The honest residual above (one cached build per new `demo-N`) is the accepted
cost of staying tooling-only.

## Related
- [Demo family index](README.md) · [Lifecycle](../rosetta_demo.md) · [Safety contract](../safety.md) · [Verification](../verification.md)
- [next-web-app](../../services/next-web-app.md) · [studio-desk](../../services/studio-desk.md) · [ant-academy](../../services/ant-academy.md)
