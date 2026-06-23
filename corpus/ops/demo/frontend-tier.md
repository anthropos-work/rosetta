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
| **studio-desk** | per-demo **Docker** image from the unmodified `Dockerfile.dev`, in the `graphql` profile | **single-port 9000** + N×10000 | Clerk-free (minted pk as a build-arg) |
| **ant-academy** | **native** `next dev` (Vercel-native; not dockerized) | **3077** + N×10000 | Clerk-free via `BENCHMARK_VISUAL_BYPASS` (anonymous browse) |

Example: `demo-2` → next-web on `:23000`, studio-desk on `:29000`, ant-academy on `:23077`.

> **studio-desk is single-port (M32).** The studio-desk image (`Dockerfile.dev`) is a **production build**
> (`npm run build:server && build:frontend`, `CMD npm start`, and it even bakes `ENV NODE_ENV=production`): one
> node process serves the built SPA *and* the API on **9000** — the `9100` Vite dev port exists only under
> `npm run dev` and is never in the container, so the demo publishes **9000+offset only** (no dead `9100`).
> **But** the base platform `docker-compose.yml` studio-desk service sets `NODE_ENV=development` +
> `FRONTEND_PORT=9100` in its `environment:` block — and a compose `environment:` value **overrides the image's
> baked `ENV`** (#M32-D4). Because the demo override's per-frontend env block is **additive** (deliberately not
> `!override`, so inherited `PORT`/`VITE_*` survive), that `development` would survive into the demo →
> `src/index.ts` `isProduction=false` → the dev path 302s the browser to the dead `9100`. So the override
> **pins `NODE_ENV=production` (+ `FRONTEND_PORT=9000`)** to win that additive merge back to the production
> `sendFile` path — which serves every dev-block route via `sendFile` + an `express.static(dist/public)` mount +
> an `index.html` SPA fallback, with no route gap (verified by code-read; #M32-D1). Full root-cause: the v1.7 M32
> milestone record.

> **Browser-trusted FAPI cert (M31).** The Clerk-free login routes the browser through Clerkenstein's fake FAPI over
> **HTTPS**; the bring-up mints a **browser-trusted** TLS cert for it via `mkcert` (idempotent `-install` + a leaf
> for `127.0.0.1 localhost ::1`), so a fresh browser renders the signed-in app with **no proceed-anyway**. It
> degrades to an openssl self-signed cert (one-time proceed-anyway) when mkcert is absent or `DEMO_NO_MKCERT=1`. Full
> story + the security/remote-VM/Firefox/expiry caveats: [`recipe-browser-login.md §B step 2`](recipe-browser-login.md).

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
> (The assert + a frontend-build failure are deliberately non-fatal — a soft RAM heuristic must never block an
> otherwise-good bring-up. #M19-D5)

> **Field note — the 12 GB VM does NOT fit on a 16 GB host** (v1.5/M25 field-bake, #M25-D2). Allocating the
> full 12 GB to the Docker Desktop VM on a **16 GB Mac** *fails to boot* the VM (`no route to host
> 192.168.65.7:2376`; `context deadline exceeded`) — macOS + Docker Desktop overhead leaves no room. The
> practical ceiling on a 16 GB box is **~10 GB VM / 2 GB swap** (~9.7 GiB usable), which boots reliably but
> **cannot co-host the full UI tier** (the ~3.7 GB next-web build spike) alongside a backend stack. On a 16 GB
> host, run the UI tier with **only one stack resident**, or use `DEMO_NO_UI=1` and verify the local-Directus
> serve at the **data-plane** level (curl cms + the per-stack Directus — the exact surface a browser calls).
> A 12 GB VM needs a **≥24 GB host** to be comfortable.

## How the pk + URLs are baked (zero platform edit)

| App | URLs | Clerk pk | Context trim |
|-----|------|----------|--------------|
| **next-web** | `--build-arg NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` / `_BACKEND_API_URL` / `_HOSTING_URL` (offset) — ARGs the Dockerfile already declares | **no pk ARG exists** → dropped into a **gitignored `apps/web/.env.local`** in the build context, read by `next build`, removed by a trap after | the repo ships **no** `.dockerignore`, so a **tooling-owned** one (`rosetta-extensions/demo-stack/frontend/next-web.dockerignore`) is applied **transiently** (never clobbers a repo one; trap-removed) to trim the 2.8 GB context (2.5 GB `node_modules`) to <100 MB |
| **studio-desk** | `--build-arg VITE_GRAPHQL_ENDPOINT` (offset) | **`VITE_CLERK_PUBLISHABLE_KEY` IS a declared ARG** → passed straight as a build-arg | the repo **already ships** a `.dockerignore` excluding `node_modules`/`dist`/`.git` — left untouched |

The split — next-web's pk via the gitignored `.env.local` (its Dockerfile declares no pk ARG) vs studio-desk's
pk straight as a build-arg (its Dockerfile *does*) — is dictated by the real, unmodified Dockerfiles (#M19-D3).
The transient tooling-owned `.dockerignore` (non-clobber, trap-removed) keeps next-web's repo byte-clean while
trimming the 2.8 GB context; studio-desk's own `.dockerignore` is sufficient and left untouched (#M19-D4).

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

## Offset-origin CORS (the backend must allow the offset frontends)

The frontends run on **offset origins** (next-web `:13000` for `demo-1`, etc.), but the backend's dev CORS
allowlist (`app/internal/cors/cors.go`) hardcodes the **un-offset** frontend origins
(`localhost:3000/3001/9000/9100`). So out of the box, every **browser → backend** REST call from the offset
origin — `/api/workforce/*` (the Workforce Intelligence dashboards), and any other direct `/api/*` consumer —
is **CORS-blocked**: the pre-flight `OPTIONS` 204s but the actual `GET` carries no `Access-Control-Allow-Origin`,
so the browser drops the response and the data panels render empty (chrome loads, charts don't).

**Decision (zero platform edit).** `cors.go` honors a **`CORS_EXTRA_ORIGINS`** env var in non-production (a
documented runtime hook — *not* a code path the demo adds). The injected override therefore sets it on the
**backend** service to this stack's offset frontend origins:

```
# each entry carries its own scheme+host (e.g. demo-1):
CORS_EXTRA_ORIGINS=http://localhost:13000,http://localhost:13001,http://localhost:19000
```

> **No offset `9100` origin (M32).** The override emits the offset origins for next-web (`3000`/`3001`) +
> studio-desk's **single-port** `9000` — not the dead `9100`. studio-desk is single-port production (the browser
> only ever talks to `9000+offset`), so the un-offset `9100` that `cors.go` still hardcodes is a dead entry the
> override no longer mirrors (#M32-D2).

This is emitted by `gen_injected_override.py` (the `backend` service gets an additive `environment:` block), so it
applies to a stack brought up **through the demo injected override** (`/demo-up`). The **dev** override
(`stack-core/gen_override.py`) does **not** emit it today and the dev bring-up runs no UI tier — so a `dev-N`'s
offset frontends would still be CORS-blocked if you ran them (a known gap, not yet wired on the dev side). It is
**not** the same as next-web's *server-side* SSR `fetch` origin
(that's the build-time `NEXT_PUBLIC_*` URLs above + the absolute-internal-origin item in §"What's out of scope");
CORS is specifically the **browser→backend** allowlist. With it set, the offset origin gets its `ACAO` header
and the REST-backed dashboards load.

## ant-academy — native, Clerk-free, session-detached, with a documented fallback

ant-academy is **Vercel-native** (not in docker-compose) and depends only on Clerk at runtime. `/demo-up`
launches it natively on `:3077+offset` **Clerk-free** using the repo's own `BENCHMARK_VISUAL_BYPASS` (a dev-only,
`NODE_ENV=development` flag that opens `/` and `/chapters/*` to anonymous traffic), paired with
`REQUIRE_ORGANIZATION_MEMBERSHIP=0` to skip the org gate. The per-demo env is a **gitignored `code/.env.local`**
overlay (zero academy-repo edits). Launching it natively (vs only documenting the step) resolved the overview's
open question toward "launch it, fall back if fiddly" — the academy is Vercel-native (not cleanly dockerizable)
and Clerk-only, so the bypass gives anonymous Clerk-free browse with no academy-repo edits (#M19-D6).

> **The native daemon is SESSION-DETACHED (the M33 "dead on a later visit" fix).** ant-academy was previously
> launched with `nohup` alone — which does **not** detach from the launcher's process group. So when a
> backgrounded `/demo-up` task's process tree was reaped on completion (or the launching session ended), the
> academy daemon died with it: the stack looked healthy at bring-up but was **dead on a later visit** (the exact
> M33 hypothesis — now **reproduced and fixed**). The launcher now starts it **session-detached** via a shared
> `demo-stack/detach.sh::launch_detached` (`setsid` where present; a portable `python3 os.setsid` double-fork on
> **macOS**, which has no `setsid`), so the daemon **survives the launching session/task ending**. _(The
> **presenter cockpit** host-native daemon had the identical bug and got the **same** `launch_detached` fix.)_

> **Without the Font Awesome Pro token, ant-academy is an intentional non-fatal SKIP.** Its `npm install`
> needs the team-issued **Font Awesome Pro** token (`FONTAWESOME_NPM_AUTH_TOKEN`), which isn't on every box —
> without it the deps can't install, so `/demo-up` **skips** ant-academy (non-fatal, prints the manual commands)
> and the demo proceeds. That's by design: ant-academy is a **Vercel-deployed, Clerk-only peripheral** surface —
> the **presenter cockpit + next-web + studio-desk carry the demo**. To enable it, provision the token (e.g. via
> [`/stack-secrets`](../secrets-spec.md)) then run `npm install` **once** in `stack-demo/ant-academy/code`.

**Default-on + non-fatal + degrades to a documented step.** If the academy clone isn't present, its deps aren't
installed (`npm install` needs the team-issued **Font Awesome Pro** token — see above), or Node < 22, the tool
prints the exact manual commands and continues — it never aborts a good demo bring-up:

```bash
cd stack-demo/ant-academy/code            # M26: the academy clone lives in the demo's OWN peer set (stack-demo)
cp .env.example .env.local                 # gitignored; keeps the repo clean
#   set REQUIRE_ORGANIZATION_MEMBERSHIP=0, reuse platform/.env's Clerk keys, add FONTAWESOME_NPM_AUTH_TOKEN
npm install
BENCHMARK_VISUAL_BYPASS=1 npm run dev -- --port 23077   # demo-2: Clerk-free anonymous browse
```

`/demo-down N` stops the native academy first (it's a process, not a container, so `compose down` can't reach
it). See [`../../services/ant-academy.md`](../../services/ant-academy.md) for the full app picture.

## Verification covers the UI tier

The M18 [verification net](../verification.md) now covers the frontends: `stack-verify`'s service registry
includes **next-web-app (:3000)** + **studio-desk (:9000)** (single-port; M32), which offset + project-rewrite like every other
service. The bring-up-tail auto-verify is **scoped to the services it started** — so a UI-on demo verifies the
frontends (an HTTP probe; a Clerk-free login redirect is a healthy 2xx/3xx/4xx), and a `--no-ui` demo scopes
them out and never false-`down`s an absent frontend (#M19-D7).

## Where the tooling lives

All of the above is `rosetta-extensions` tooling, authored + tagged in the authoring copy and consumed per-stack
at the **current pinned tag** (`stack-demo/rosetta-extensions @ storytelling-postfix-1` — the M19 UI tier first
shipped at `dress-rehearsal-m19`; the CORS + token-strip items were later, ≥ `dress-rehearsal-m20-fix15`/`fix17`;
the session-detach fix below lands at `storytelling-postfix-1`):

- `stack-injection/gen_injected_override.py` — appends the two frontends to the injected override (offset
  `ports:!override`, `image: demo-N-*` + `build:!reset null` + `pull_policy:never`, `mem_limit:1g`,
  `profiles:!override [graphql]`); `--no-ui` clears the tier. Also sets `CORS_EXTRA_ORIGINS` on the **backend**
  service to the offset frontend origins (see §"Offset-origin CORS"), and **strips the inherited prod
  `DIRECTUS_TOKEN`** (`DIRECTUS_TOKEN=`) on **every** emitted service + both frontends — no prod credential rides
  in a demo container, and studio-desk's prod-Directus *write* path is disarmed (fix16/fix17; see
  [`../safety.md`](../safety.md) §2.3 + §2.2).
- `demo-stack/up-injected.sh` — the per-demo serial-before-up frontend build (offset URLs + minted pk +
  tag-guard), the 12 GB VM pre-flight, the `--no-ui` (`DEMO_NO_UI`) escape, the scoped verify.
- `demo-stack/frontend/next-web.dockerignore` — the tooling-owned context trim for next-web.
- `demo-stack/ant-academy.sh` — the native academy launcher / stopper / documented fallback; launches the
  daemon **session-detached** (via `detach.sh::launch_detached`) and **non-fatally skips** when the Font Awesome
  Pro token / deps are absent.
- `demo-stack/detach.sh` — the shared `launch_detached` helper (`setsid`, or a `python3 os.setsid` double-fork on
  macOS) that session-detaches the host-native daemons (ant-academy **and** the presenter cockpit) so they
  survive the launching `/demo-up` session/task being reaped.

## What's out of scope (the user-owned follow-up)

**True zero-rebuild** — one frontend image that serves every stack with the port/pk switched at *runtime* —
would require **platform-source changes** (runtime rewrites in `next.config.mjs`, an absolute internal origin for
SSR `fetch` in `server.graphql.ts`, a `window.__ENV` shim + explicit `publishableKey` on `<ClerkProvider>`,
optionally `output:'standalone'`). Those are real platform edits with PR/review/prod risk — **forbidden** for the
demo tooling to make locally. It's documented here as an **optional upstream PR you own** (a deferred/unscheduled deploy-CI
precedent), **not built** in M19. The honest residual above (one cached build per new `demo-N`) is the accepted
cost of staying tooling-only.

## Related
- [Demo family index](README.md) · [Lifecycle](../rosetta_demo.md) · [Safety contract](../safety.md) · [Verification](../verification.md)
- [next-web-app](../../services/next-web-app.md) · [studio-desk](../../services/studio-desk.md) · [ant-academy](../../services/ant-academy.md)
