# The demo frontend tier — the UI of a demo stack

**Purpose.** Make a `demo-N` (or `dev-N`) stack **actually demoable**: bring up the user-facing **UI tier** —
**next-web-app** (the Workforce app) + **studio-desk** at offset ports, plus **ant-academy** natively — so a
stakeholder lands on a real, clickable, Clerk-free UI, not just a running backend. This is the v1.3b M19
deliverable that completes the [demo family](README.md): up → snapshot → seed → **see it in a browser** → down.

> **The demo-patch mechanism is specified in [`demopatch-spec.md`](demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

> **Read [`../safety.md`](../safety.md) first** for *why* this is safe, and [`../rosetta_demo.md`](../rosetta_demo.md)
> for the stack lifecycle this extends. This page is the **frontend-specific** "how the UI tier is built and run".

> **The hard line (non-negotiable).** **Zero platform-repo edits.** next-web-app, studio-desk, and ant-academy
> stay **byte-for-byte pristine** — their repos are used only as a Docker **build context**, and every
> per-demo difference rides a **gitignored** overlay (`.env.local`), a **sha-pinned demo-patch** applied to the
> demo's own ephemeral clone and reverted after ([`demopatch-spec.md`](demopatch-spec.md)), or a
> **tooling-owned file in `rosetta-extensions`**. Nothing the demo tooling does touches a tracked platform file.

> ### There are THREE build shapes, not two — and the third is already in production
>
> This page used to say the platform Dockerfiles are *"consumed UNMODIFIED"*, full stop, which reads as an
> exhaustive statement and is not one. **A tooling-owned Dockerfile in `rosetta-extensions` is a fully
> sanctioned third shape, and the demo has shipped one since M224:**
> `rosetta-extensions/demo-stack/frontend/hiring.Dockerfile` builds `apps/hiring` from the **same unmodified
> next-web-app clone** the web app builds from, because the platform's `Dockerfile.dev` hardcodes the WEB app
> end-to-end (`--filter=@anthropos/web-app`, `start:web`, `EXPOSE 3000`) and cannot be reused verbatim.
>
> | shape | who owns the Dockerfile | platform repo is | example |
> |---|---|---|---|
> | 1 | the platform | build context, Dockerfile consumed as-is | ⚠️ **no member left.** `next-web` moved to shape 3 at M257 iter-09 and **`studio-desk` (`Dockerfile.dev`) followed at v2.8 M258 TIK-A** — this cell named studio-desk as *the* example until the M258 close |
> | 2 | the platform | build context, **source** patched in the ephemeral clone + reverted | the demo-patches — **23** on disk (`ls demo-stack/patches/*/*.yaml \| wc -l` at rext `415240f`), of which **13** are image-baked (11 `next-web-app` · 2 `app`) and 5 are `ant-academy`. ⚠️ **The 5 `studio-desk` manifests are still on disk but RETIRED and no longer applied** — they are sha-pinned to `app/core/{main.ts,scaffold/*}`, which the Next migration deletes, so demopatch refuses them at G2. They stay in the build's cache fingerprint deliberately, so a pre-migration image still rebuilds, patched-then-reverted around a **native** `next dev` rather than a build. *(This cell read "the 11 demo-patches" — the M224-era distinct-manifest total, four milestones stale. The authoritative inventory is [`demopatch-spec.md` §5](demopatch-spec.md), directory-fenced by `TestPatchInventory`.)* |
> | 3 | **`rosetta-extensions`** | **build context only** — rext supplies the Dockerfile | **`hiring`** (`frontend/hiring.Dockerfile`), **`next-web`** (`frontend/next-web.Dockerfile`, M257 iter-09) **and `studio-desk`** (`frontend/studio-desk.Dockerfile`). ⚠️ **The studio-desk figure moved TWICE:** M258 TIK-A multi-staged the Vite/Express image **1.7 GB → 1.35 GB** with a prune-and-copy; the **Next migration** then replaced that shape entirely with `output: 'standalone'` at **119.6 MB** (cold build 50 s, arm64 Mac / containerd) — ~11× smaller again, and studio-desk is now the SMALLEST of the three. The prune argument is moot: standalone TRACES node_modules, so the `@clerk/clerk-js` wallet tree M258 measured as irreducible is simply never copied. All three of the demo's Docker-built frontends now live here |
>
> Shape 3 is *stronger* on the hard line than shape 2, not weaker: nothing in the platform repo is touched at
> all, not even transiently. **It is the shape v2.8's largest speed lever uses — and as of M257 iter-09 that
> is no longer a forecast.** L1 landed: both Next images are multi-stage `.next/standalone` builds, and
> `next-web` moved from shape 1 to shape 3 to get there, because making the platform's own `Dockerfile.dev`
> multi-stage would have been a platform-repo edit. See [`build-budget.md`](build-budget.md). **The
> "out of scope / forbidden upstream PR" section at the end of this page predates shape 3 and lists
> `output:'standalone'` among the things only an upstream PR could deliver. That is no longer true** —
> M255 proved it needs **zero** source edits and **zero** demo-patches via `ENV NEXT_PRIVATE_STANDALONE=1`
> in a tooling-owned Dockerfile, and M257 iter-09 shipped exactly that. **That section has now been rewritten
> with the achieved numbers at the M257 close** (D121: one rewrite, not two) — see §"What's out of scope" at
> the end of this page for the two-host table.
>
> **One flag in that Dockerfile is load-bearing and easy to drop: `turbo … --env-mode=loose`.** Turbo 2
> defaults to `strict` and forwards only the variables named in `turbo.json`'s `globalEnv`/task `env`;
> `NEXT_PRIVATE_STANDALONE` is **not** among the 44 names there. Without the flag the variable never reaches
> `next build`, `output` stays undefined, **the build still exits 0**, and the image is the old ~4 GB one.
> Both Dockerfiles therefore assert `test -d apps/<app>/.next/standalone` and fail loudly rather than trust
> it — and at the M257 final harden that assert stopped being trusted too. It had been fenced by *string
> presence*: two tests checked the text was in the file, which is exactly as strong as the reader's shell
> grammar. `test_the_standalone_assert_ACTUALLY_FAILS_a_build_not_merely_appears_in_one` now extracts the
> shipped `RUN` line, runs it under `sh` against a tree with no standalone output, and requires a non-zero
> exit naming `NEXT_PRIVATE_STANDALONE`. Mutation control: losing the `exit 1` makes the fragment exit 0 and
> every string assert still passes.
>
> **⚠️ The next-web image carries `apps/web/.env` — a real-Clerk publishable key and `CLERK_SECRET_KEY` —
> and it is loaded at runtime, not merely carried** (measured at the M257 final harden on the real post-L1
> image: 19,087 bytes at `/app/apps/web/.env`, with **no** `.env.local` beside it; the hiring image has
> neither). It reaches the runner stage inside `.next/standalone`, and standalone's `server.js` calls
> `loadEnvConfig` at boot. What keeps a demo honest is that `@next/env` never overwrites an already-set
> variable and the injected override sets the four `CLERK_*` explicitly — so the residual is the **set
> difference**: anything in that file compose does not name. The cause is one missing pattern in a
> **tooling-owned** file: `demo-stack/frontend/next-web.dockerignore` excludes `.env*`, Docker matches
> `.dockerignore` patterns from the **context root**, and that rule therefore covers `./.env` and nothing
> nested — while every other rule in the file is deliberately paired with a `**/` twin. **Do not take the
> one-line fix:** `**/.env*` also excludes `apps/web/.env.local`, the overlay carrying the *minted* key into
> `next build`, so the tidy repair bakes the **real** Clerk key — the M218 iter-03 incident, re-created by
> its own fix. Routed as `FIX-M257-dockerignore-env-pattern-unpaired`; the net under it is the ISOLATION
> gate clause, which books a non-minted key in the bundle as `foreign_pk` and reds the campaign.

## What `/demo-up` brings up (UI tier)

| App | How it runs | Port (base + offset) | Auth in the demo |
|-----|-------------|----------------------|------------------|
| **next-web-app** (Workforce) | per-demo **Docker** image from the **rext-owned** `frontend/next-web.Dockerfile` (build shape 3 above), built from the unmodified `next-web-app` clone, in the demo's `core` profile. *Built from the platform's own `Dockerfile.dev` until M257 iter-09; L1 moved it so the image could be multi-stage without a platform-repo edit.* | **3000** + N×10000 | Clerk-free (Clerkenstein-minted pk baked into the bundle) |
| **hiring** (the real `apps/hiring`) | per-demo **Docker** image from the **rext-owned** `frontend/hiring.Dockerfile` (build shape 3 above), built from the same unmodified `next-web-app` clone; a **net-new** compose service `hiring-app` with `profiles: [<derived-default>]` | **3001** + N×10000 | Clerk-free (minted pk baked; `CLERK_API_URL` → the fake BAPI alias) |
| **studio-desk** | per-demo **Docker** image from the **rext-owned** `frontend/studio-desk.Dockerfile` (build shape 3 above), built from the unmodified `studio-desk` clone, in the demo's `core` profile — **same as next-web**. Since the **Next migration** it is an `output: 'standalone'` build: **119.6 MB**, `CMD ["node","server.js"]`. *(Platform `Dockerfile.dev` → M258 TIK-A prune-and-copy 1.7 GB → 1.35 GB → standalone 119.6 MB.)* | **single-port 9000** + N×10000 *(9100 was the Vite dev port and never existed in the container)* | Clerk-free (minted pk as a **`NEXT_PUBLIC_*`** build-arg — the `VITE_*` names are retired, and docker discards an undeclared build-arg **silently**) |
| **ant-academy** | **native** `next dev` (Vercel-native; not dockerized) | **3077** + N×10000 | **Clerkenstein-wired (v2.3 M220)** — the demo's minted pk + the disarmed fake BAPI, read from `<stack>/.env.demo-N`. It **shares the demo's session**: a hero who clicks through from next-web arrives at the academy **signed in as herself**. *Was keyless via the `e2e_persona` bypass — see the box below; that is now removed.* |

Example: `demo-2` → next-web on `:23000`, hiring on `:23001`, studio-desk on `:29000`, ant-academy on `:23077`.

> **The `hiring` row is net-new at M257x — this table listed THREE apps and called itself "what `/demo-up`
> brings up (UI tier)".** The demo has run **four** since v2.4 M224 ("the callback" — the two-app demo): the
> real `apps/hiring` is emitted as a `hiring-app` compose service by
> `stack-injection/gen_injected_override.py:459-509` (called at `:736-737`) on port `3001 + offset` (`:484`, `:504`), and
> the rest of this page already referred to its image (`demo-N-hiring`) and its build shape without it ever
> appearing here. **An enumeration presented as complete is a claim.**

> ### ⚠️ NEITHER containerized frontend rides its base-compose profile in a demo — the injection **overrides** both to `core`
>
> The studio-desk row used to read *"in the `frontend`/`studio-desk` profiles"*, which was wrong twice over.
>
> **In the base compose** (`stack-demo/platform` @ `0c91421df`) the two frontends sit in **different**
> profiles, and `frontend` is **not** one of studio-desk's: `next-web-app` declares
> `profiles: [frontend, all]` (`docker-compose.yml:168`) and `studio-desk` declares
> `profiles: [studio-desk, all]` (`docker-compose.yml:141`). Neither is in `core`, the platform's default
> (`Makefile:10`, `PROFILE ?= core`). *(Selecting either token **alone** also exits 1 — both declare
> `depends_on: backend` (`docker-compose.yml:165-167` and `:138-140`) which those profiles do not select, so
> compose rejects the project as invalid.)*
>
> **In a demo neither of those profiles is ever selected.** `up-injected.sh` derives the platform's default
> profile from the platform's own compose and brings the stack up with **that one only**
> (`demo-stack/up-injected.sh:2174` → `stack-injection/platform_topology.py profile`, then `:2179-2180`
> `--profile "$COMPOSE_PROFILE"`). The two frontends reach it because the generated override **re-declares**
> their profile: `stack-injection/gen_injected_override.py:425` emits
> `profiles: !override [<derived-default>]` for **each** entry of `FRONTENDS`, which is why the generator's own
> comment (`:170-174`) says they *"live in the `frontend`/`studio-desk`/`all` profiles in the platform compose,
> NOT the demo's own default profile — so they DON'T appear in the resolved cfg this generator walks."*
> So on a current clone **both** rows read `core`, and both read it because rext put them there — not because
> the platform did. (`hiring` is a rext-owned shape-3 service and is emitted the same way.)

> **⚠️ UNVERIFIED against the renamed profile.** Platform `0dab54d` (the v9.0 "support-in-app"
> commit) **renamed the default compose profile `graphql` → `core`**; there is no `graphql`
> profile in `docker-compose.yml` any more. The two `graphql` tokens in the table above describe
> what `rosetta-extensions`' `up-injected.sh` passes to compose, and that repo was **not** checked
> when this note was written. If the bring-up still passes `--profile graphql`, compose selects
> **nothing** and the UI tier silently does not start — `docker compose --profile <unknown>` exits
> 0. Verify with `docker compose --profile graphql config --services` against a current platform
> clone before trusting the table, and re-pin the token in `rosetta-extensions` if it is stale.

> ### 🔴 The academy used to **POISON the demo session** — and one click destroyed a live demo (v2.3 M220 S5/i)
>
> This is the single most damaging defect the demo family has shipped, and it hid behind *"the port answers"*
> for four releases.
>
> `ant-academy.sh` built the academy's `.env.local` by **grepping `platform/.env`** for `CLERK_*`. That file
> carries **11 matching lines**; all 11 were written, and in a dotenv file **the last one wins** — and it is not
> the demo's minted key. So `@clerk/nextjs` found **no usable publishable key** and fell into **keyless mode**,
> whose middleware answered every request on `:3077+offset` with:
>
> ```
> Set-Cookie: __session=;     Expires=Thu, 01 Jan 1970 00:00:00 GMT    ← DELETES the demo's session
> Set-Cookie: __client_uat=0; Domain=<tailnet>                         ← DOMAIN-wide, not port-scoped
> ```
>
> **Cookies scope by HOST, not by PORT.** The academy on `:3077+offset` therefore **clobbered the session
> next-web holds on `:3000+offset`**. Two measured consequences:
>
> 1. **A presenter who clicked "AI Academy" was LOGGED OUT of their own live demo**, into
>    `ERR_TOO_MANY_REDIRECTS`. The blank academy page was the *lesser* half of the bug.
> 2. **Every employee coverage sweep aborted** at that link — so the employee vantage had **no runnable sweep
>    at all**, which is itself an absence-read-as-success risk.
>
> It was also a **safety** defect of the `DIRECTUS_TOKEN` fix16/17 class: the grep copied the **REAL Clerk app's
> `CLERK_SECRET_KEY`** — a production secret — into a demo process, which [`safety.md`](../safety.md) forbids
> outright.
>
> **The fix: wire the academy to Clerkenstein**, exactly as studio-desk has been since v1.10. It reads the demo's
> minted pk + the disarmed fake BAPI + the fixed RS256 public key out of `<stack>/.env.demo-N` (never
> `platform/.env`), so keyless mode never engages, **no cookie is ever deleted, and the academy SHARES the demo's
> session** — the hero arrives signed in as herself.
>
> **The `e2e_persona` bypass (`BENCHMARK_VISUAL_BYPASS` + `NEXT_PUBLIC_E2E_AUTH`) is REMOVED.** It existed only to
> fake an authenticated session on a *keyless* academy. Kept alongside real keys it is worse than either alone:
> `proxy.js` short-circuits on the persona cookie **before** it resolves the real session, so the academy would
> render a generic **"E2E Member"** to a presenter logged in as **Maya** — a persona self-consistency defect
> shipped by our own launcher.
>
> **The fake BAPI is now published on `127.0.0.1:5401+offset`** — the demo's **first loopback-bound** port. The
> academy is the demo's one **host-native** frontend (`next dev`, never dockerized), so it cannot use the compose
> alias `api.clerk.com` every container reaches the BAPI through; without a published port its only reachable
> `CLERK_API_URL` would be the default — **real `api.clerk.com`**. Loopback, not `0.0.0.0`: its only consumer is
> a process on the same box, and a disarmed BAPI (it ignores the bearer entirely) is the last thing that should
> be ambient on a tailnet.
>
> **The DoD is not "it paints".** It is the controlled A/B — *log in → click the academy → go back → **still
> logged in***. Proven on `billion` from a tailnet peer in a real browser: `__session` present throughout,
> `__client_uat` a live timestamp (never `0`), and `/profile` still rendering **"Maya Chen"** after the visit.
> Direct `curl` at `:13077` now returns **zero `Set-Cookie` headers**.

> ### ⚠️ A **stale academy from a previous demo** can keep serving on your port (v2.3 M220)
>
> Measured on `billion`, on M220's own bring-up: an academy `next dev` from an earlier demo — started **11½
> hours** before — was **still bound to `:13077`**. Teardown reaps the academy (by pidfile *and* by port), but
> **nothing reaped before a LAUNCH**. So the freshly-wired academy died instantly with
> `EADDRINUSE 0.0.0.0:13077`, the bring-up moved on — **and the ORPHAN kept answering**.
>
> That is the worst possible shape of this bug: *the port answers*, so a render-probe polls the **orphan** and
> can go green, and the presenter (and this milestone's own session-survival proof) would have been talking to
> the **old, keyless** academy — the very process whose cookie deletion M220 exists to remove. **A stale artifact
> outliving the thing it describes, then read as evidence.**
>
> `ant-academy.sh` now **reaps its own port before launching**, using the same identity-checked reaper teardown
> uses (it refuses an empty pattern and matches the academy clone path / this stack's port / `next-server`, so a
> co-resident demo's academy or an unrelated Next app is never touched).

> ### ⚠️ ant-academy needs **Node ≥ 22**, and "started" must mean **the port answers** (F-13, M219)
>
> For four releases the bring-up reported the academy as **started** while the demo's **AI Academy** link served
> a bare **502 for the entire life of the stack**. Two defects, both on the bring-up's *own reporting path*:
>
> 1. **The node check tested EXISTENCE, not the version.** `command -v node` passed on the demo VM — with
>    **Node 18**. ant-academy declares `"engines": { "node": ">=22" }`, ships Next 16, and its toolchain imports
>    `node:util`'s `styleText` (Node ≥ 20), so `next dev` died at import **every single time**:
>    `SyntaxError: The requested module 'node:util' does not provide an export named 'styleText'`.
>    **A version requirement that is only checked for existence is not checked at all.** (The VM had v22.22.1
>    installed under nvm and unused — a session-detached daemon never sources `nvm.sh`.) The launcher now
>    resolves a satisfying node under `~/.nvm`, or **fails loud with the remedy**.
> 2. **The liveness probe polled `kill -0 $pid` for 3 s**, then printed *"started (pid N)"*. The pid was alive at
>    t=1 s and gone by t=5 s. **A probe that cannot outlive the thing it probes is not a probe.** It now polls the
>    **port** (bounded by `ACADEMY_READY_TIMEOUT_S`, default **120 s** — `next dev` genuinely cold-compiles for
>    30–60 s on a slow VM), watching the pid so a crash fails fast, and prints the log tail when the port never
>    answers. The only success phrasing is **`started + SERVING on :PORT`**.
>
> The demo-coverage sweep catches this independently as a **cross-port failure** (`:3077+offset` → HTTP 502) —
> which is how it was confirmed. **The launcher's own test fixtures were the bug**: the npm stub was
> `run) exec sleep 30` (alive, serving nothing) and the node stub answered `-v` with silence, so the suite was
> green *against the broken launcher*. The stubs now model a real academy, and both failure shapes are fenced
> (the liveness one proven RED — the old check cheerfully printed "started" for a daemon serving nothing).

> ### ⚠️ The academy demo-config must survive the LIFECYCLE, not just a fresh up (M245, v2.6 "sound check")
>
> A **fresh `/demo-up` always** brings the academy up fully configured — `ant-academy.sh` writes `code/.env.local`
> (minted pk + `REQUIRE_ORGANIZATION_MEMBERSHIP=0`), applies the four FS-published/dev-origins demo-patches to the
> **ephemeral clone**, and launches `next dev` with `ACADEMY_DEMO_FS_PUBLISHED=1`. But that config lives **in-process
> and in-clone**, and on `billion` (M244) a re-up left the academy **alive but serving a stock, empty, Clerk-gated
> catalog** — **0 course cards on `/library`** over an HTTP 200. Two lifecycle vectors, both leaving the pid alive:
> 1. **`.env.local` dewire.** `stacksecrets provision` also targets `ant-academy/code/.env.local` and runs on **every**
>    `/demo-up` (before the academy launch), copy-if-absent-**appending** the **real** Clerk app's keys. A fresh up's
>    truncating write beats them; the old **"already running" early-exit** skipped that write on a re-up, so the
>    appended real key sat **last** and — dotenv last-wins — **dewired** Clerkenstein.
> 2. **Clone-patch revert.** A `--stop` / reset-to-seed reverts the FS patches from the clone; next dev recompiles the
>    pristine (empty) resolver, and no re-up re-applied them.
>
> **The fix (tooling-only):** `ant-academy.sh`'s "already running" branch no longer early-exits — it **reconciles the
> durable config in place on every invocation** (rewrites `.env.local` authoritatively — it is the single authoritative
> writer, always beating a stacksecrets append — and re-applies the clone patches, idempotent; next dev HMR heals a
> reverted clone), then **verifies the running process actually renders `/library`**, keeping it (no cold restart) if
> so or relaunching if not. And a **standing check** — demo autoverify assert *(f)* now asserts the academy renders its
> catalog at the bring-up tail — so this **regresses loudly** next time instead of silently serving an empty portal.
> Proven on `billion`: reverting the patches + polluting `.env.local` (0 cards reproduced), then one re-run healed it
> back to a rendering catalog with **no restart**.

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

> **studio-desk is a CLERKENSTEIN-authenticated demo surface (v1.10 "method acting" postfix).** The
> manager's **"Anthropos Studio" left-nav** opens the demo's own studio-desk on `:9000+offset`, where the
> logged-in hero authenticates **through Clerkenstein** (the demo's fake FAPI/BAPI) exactly like every Go
> service — it is the **actual logged-in hero**, not a mock-auth bypass. (An earlier postfix used `MOCK_CLERK`
> to render the surface by skipping auth; that was reverted — studio-desk must be the authenticated hero.)
> The production image applies `clerkMiddleware()` + `requireAuth` + `checkEnterpriseAndAdmin` to **all**
> routes; the wiring that makes that pass in a demo:
> - **The FAPI handshake (per-app, no cross-port cookie).** studio-desk's `clerkMiddleware()` 302s an
>   unauthenticated browser to the **fake FAPI** `/v1/client/handshake`, which bounces back a
>   `__clerk_handshake` RS256 token (kid `clerkenstein-rs256-demo`) that `@clerk/express` verifies
>   **networklessly** via `CLERK_JWT_KEY`. **Each app drives its OWN handshake** against the demo's
>   single fake FAPI (which holds the active-seat selection server-side), so the per-port `__session`
>   cookie is **not** needed — the split-port topology is a non-issue. The minted **pk** is baked
>   (`VITE_CLERK_PUBLISHABLE_KEY`) so the SPA derives the same fake-FAPI host the backend talks to.
> - **The admin gate (`checkEnterpriseAndAdmin`).** Once authenticated, studio-desk calls the **fake BAPI**
>   `getOrganizationMembershipList({userId})` and requires a membership with a Studio-eligible role
>   (`admin`/`content_creator`). The fake BAPI is **roster-aware**: `cmd/fake-bapi` reads the **same**
>   `FAKE_FAPI_ROSTER` the fake FAPI loads and seeds each seeded hero's `(org, user) → org_role`, so a
>   **manager** (Dan/Leah = `admin`) **passes** the gate and an **employee** (`member`) is correctly
>   redirected off Studio — the real role-gated behaviour. Without the roster seed the BAPI knows only the
>   universal `user_clerkenstein`, so a logged-in hero's membership list is empty and they bounce to
>   `WEB_APP_URL`.
> - **The requireAuth fallback.** The injected override pins `CLERK_SIGN_IN_URL`/`WEB_APP_URL` at the demo's
>   **own offset** next-web (`:3000+offset`, which HAS a `/login` route) — so the unauthenticated/non-admin
>   fallbacks land somewhere **live**, never the dead un-offset `:3000` (`ERR_TOO_MANY_REDIRECTS`).
>
> > **The auth path needs no demopatch, no mock bundle.** studio-desk's **auth** is the unmodified
> > production code, driven entirely by the **runtime** `CLERK_*` env + the baked pk + the roster-aware fake
> > BAPI. (Scoped to auth: studio-desk **is** a first-class demopatch target now — the **5** M249+M253 source
> > patches for "Back to Cockpit" / prod-eject / first-paint; see the bake table below + [`demopatch-spec.md`](demopatch-spec.md).)
> > (Clerkenstein itself — the fake FAPI/BAPI in `rosetta-extensions` — is tooling-owned and freely
> > edited; the platform repos are untouched.) A `demo-N-studio-desk` image with a **stale pk/offset** is
> > reused by the tag-guard, so clearing it (`docker image rm demo-N-studio-desk`) forces a fresh Clerkenstein
> > bake; the roster-aware BAPI re-seeds on every re-up.

> **v2.7 "july jitter" M252 — the AI-provider `env_file` (the auth model is UNCHANGED).** M252 did **not** touch
> the auth posture above: the demo studio stays the **Clerkenstein-authenticated hero** — a logged-in org-admin
> hero (the manager) 302s through the fake-FAPI handshake and passes `checkEnterpriseAndAdmin` exactly as
> described. (A raw *unauthenticated* `curl` 302s to `/login` — that is the production `clerkMiddleware()`
> catch-all behaving as designed for a browser with **no** session, **not** an unreachable studio.) What M252
> fixes is a distinct **AI-provider** gap: studio-desk is a **base-compose** service, so in a demo it inherited
> **only `platform/.env`** — which carries **no AI-provider keys** — so `POST /api/ai/completion` 500'd. M252
> wires the studio-desk clone's own `.env` into the container via an existence-guarded `env_file` (the
> `gen_injected_override.py` bullet in §"Where the tooling lives" below), supplying the studio's own AI-provider
> keys (`AI_OPENAI_API_KEY` + `AI_ANTHROPIC_API_KEY`). **No `MOCK_CLERK`, no auth change** — a `MOCK_CLERK=true`
> line would regress the demo to the legacy bypass and **fail** the pinned regression tests
> (`test_studio_desk_env_clerkenstein_no_mock_and_offset_sign_in` /
> `test_studio_desk_block_shape_single_port_clerkenstein_wired`, which assert **no** `MOCK_CLERK` line in the
> studio-desk block). This is what makes `/api/ai/completion` callable **by the logged-in hero** in a Playthrough
> ([`playthroughs.md`](playthroughs.md): `pt-studio-advanced-generate` / `pt-studio-guided-generate`) — see
> [`../../services/studio-desk.md`](../../services/studio-desk.md) § Demo AI wiring.

> **Browser-trusted FAPI cert (M31; M213 remote path).** The Clerk-free login routes the browser through
> Clerkenstein's fake FAPI over **HTTPS**; the bring-up mints a **browser-trusted** TLS cert for it. For a **local**
> demo (default) that's `mkcert` (idempotent `-install` + a leaf for `127.0.0.1 localhost ::1`), degrading to an
> openssl self-signed cert (one-time proceed-anyway) when mkcert is absent or `DEMO_NO_MKCERT=1`. For a **remote /
> tailnet** demo (`/demo-up --public-host <magicdns>`, M213/v2.2), the cert is minted via **`tailscale cert`** — a
> real Let's Encrypt cert **trusted tailnet-wide with no per-machine CA install**, so a teammate's browser trusts it
> with no proceed-anyway (falls back to the local mkcert/openssl path if `tailscaled` isn't up). Same output paths
> either way, so the mount is unchanged. Full story + the security/remote-VM/Firefox/expiry/renewal caveats:
> [`recipe-browser-login.md §B step 2`](recipe-browser-login.md).

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
- **The residual (honest) — rewritten with ACHIEVED numbers at the M257 close (D-v28-10 / D121).** This bullet
  read *"one ~3-minute, ~3.7 GB cached build per frontend"*. **Both halves are retracted, and it named no
  host** — which is the defect v2.8's own standing rule (*state the environment with every number*) exists to
  catch. **Measured on `macmini`** — the local M4 Pro Mac mini, **arm64 with the containerd image store**, a
  permanently-contended workstation — at M257 iter-09, n=3 cold cycles: a brand-new `demo-N` costs
  **104.60 s for the whole UI tier**, which is **THREE images, not "per frontend"** (`ui_next_web` 53.31 s ·
  `ui_hiring` 44.21 s · `ui_studio_desk` 7.08 s, sub-phase p50s). The two Next images weigh **417 MB**
  (next-web) and **380 MB** (hiring), against **4.04 GB** and **3.94 GB** before L1; the same block on the same
  host cost **246.23 s** pre-lever. So the retracted *"~3.7 GB"* was an **image size** (roughly right for one
  image, of three) and the retracted *"~3 minutes"* was per-image wall-clock on a slower host, no longer a
  figure for anything here. **Seconds do not transfer between
  hosts** — the pre-L1 UI tier is **436.1 s on `billion`** (x86_64), and that number stays `billion`'s.
  What is unchanged is *why* a per-demo build exists at all: zero-platform-edit + per-stack pk/URL baking.
  *True* zero-rebuild would need runtime-configurable URLs + pk in the platform source — see §"What's out of
  scope", **which is itself re-cut at the end of this page**: `output:'standalone'` is no longer on it.
- **Built serially, before `compose up`.** The two frontend builds run **one at a time, before** the stack
  starts — kept out of the parallel Go-service fan-out so the build RAM spike never overlaps anything else.
- **Non-fatal — actually true now (v1.10b M49 #7).** A frontend build failure **warns** but never aborts the
  backend bring-up. The build step was always non-fatal, but `compose up` would still try to **start** a
  frontend whose image is **absent** (a failed/skipped build) and abort the whole bring-up under
  `set -euo pipefail` — so backend + set-dress + verify + cockpit never ran. Now an absent frontend image is
  **scaled to 0 replicas** at `compose up` (`--scale next-web-app=0` / `--scale studio-desk=0`), so the rest of
  the stack comes up and the demo is usable (API + cockpit); re-run to retry the UI, or `DEMO_NO_UI=1` to skip
  it entirely (under `--no-ui` the injected override omits the frontends, so there's nothing to scale).

## The 12 GB Docker-VM prerequisite

**Runtime is cheap** — measured **~0.66 GiB for BOTH stacks** (dev + demo, 27 containers). The only spike is the
**build**: a next-web compile lane whose **measured** heap peak is **3,116 MiB on `macmini`** (**3,900 MiB** on
`billion`, **4,223 MiB** on the retired M1 Pro `laptop`). On an undersized Docker VM already holding the dev
stack, that spike **swap-thrashes**.

> **⚠️ Retracted at the M257 close (D-v28-10): this prerequisite was reasoned from an IMAGE SIZE used as a
> MEMORY figure, and the diagnosis beside it was wrongly exclusive.** The paragraph said *"a ~3.7 GB next-web
> compile"*. **3.7 GB was never a memory measurement** — it was an image size, and post-L1 the next-web *image*
> is **417 MB on `macmini`**, so the number is now wrong in both readings. The figure this floor actually rests
> on is the per-lane heap peak in `stack-core/hostprofiles/*.json` (`lane_heap_measured_peak_mib`), which is
> **measured per host and differs by host**: 3,116 / 3,900 / 4,223 MiB. That band **brackets 3.7 GB**, which is
> precisely why the wrong number survived four releases — it was approximately right for a reason nobody had
> checked. **The 12 GB floor itself stands**; it is the derivation that was broken, not the conclusion.
>
> The second half — *"the original 'the build takes an hour' symptom was pure memory starvation, **not** a slow
> build"* — is **half true and wrongly exclusive.** Swap-thrashing on an undersized VM holding a second stack
> is real, reproducible, and still the thing this section exists to prevent. But M255 measured a cold cycle on
> a box under **no** memory pressure (`billion`, peak load1 **4.90 of 8**) that still spent **288.4 s** in
> image export/unpack alone: the build genuinely *was* slow, independently of memory. **Both were true and the
> sentence asserted only one.** M257's L1 removed the I/O half — `exporting to image` for the two Next images
> is **136.4 s → 3.8 s** combined on `macmini` — which leaves memory as the dominant remaining build risk, so
> this prerequisite matters *more* after L1, not less.

> **Set the Docker Desktop VM to 12 GB / swap 3 GB** (Settings → Resources). `/demo-up` runs a **non-fatal
> pre-flight assert**: below 12 GB it prints a clear warning (raise the VM, or run `DEMO_NO_UI=1`) but continues
> — a smaller VM may still build fine if no other stack is up. Override the floor with `DEMO_VM_MIN_GIB=N`.
> (The assert + a frontend-build failure are deliberately non-fatal — a soft RAM heuristic must never block an
> otherwise-good bring-up. #M19-D5)

> **Field note — the 12 GB VM does NOT fit on a 16 GB host** (v1.5/M25 field-bake, #M25-D2). Allocating the
> full 12 GB to the Docker Desktop VM on a **16 GB Mac** *fails to boot* the VM (`no route to host
> 192.168.65.7:2376`; `context deadline exceeded`) — macOS + Docker Desktop overhead leaves no room. The
> practical ceiling on a 16 GB box is **~10 GB VM / 2 GB swap** (~9.7 GiB usable), which boots reliably but
> **cannot co-host the full UI tier** (a next-web build lane peaking at **~3.0–4.2 GiB depending on host** —
> see the retraction above; this parenthetical used to cite the ~3.7 GB *image* size) alongside a backend
> stack. On a 16 GB
> host, run the UI tier with **only one stack resident**, or use `DEMO_NO_UI=1` and verify the local-Directus
> serve at the **data-plane** level (curl cms + the per-stack Directus — the exact surface a browser calls).
> A 12 GB VM needs a **≥24 GB host** to be comfortable.

### Disk headroom — a second non-fatal pre-flight (v1.10b M49 #6)

Alongside the RAM check, `/demo-up` runs a **disk-headroom pre-flight** (mirrors the RAM assert: a warning,
never a gate). Each demo's images — `demo-N-next-web`, `demo-N-hiring`, `demo-N-studio-desk`, the
`demo-N-<svc>:injected` Go services, `demo-N-fake-fapi`/`-fake-bapi` — plus the BuildKit layer cache
**accumulate**, and dead demo stacks used to leave their images behind, so a box could slowly fill until a
build hit `ENOSPC` mid-stream.

> **⚠️ v2.8 M255 re-sized the floor, because the numbers it was reasoned from were an order of magnitude
> stale.** This section used to say *"the ~3.7 GB build cache"* and set the floor at ~20 GB. **Measured on
> `billion`, 2026-07-27: the build cache is 105.4 GB**, one cold-images cycle **peaks at 18 GiB of
> consumption** (the export leg stages layers before it unpacks them, so peak ≫ the ~12.6 GiB of resident
> image bytes) and leaves **~2 GiB** behind. The floor is now **25 GiB** = the measured 18 GiB + a 7 GiB
> reserve, and `stack-core/hostprofiles/billion.json` carries the same arithmetic
> (`projected_image_gib` 18 + `disk_floor_gib` 7) so the operator warning and the release gate cannot drift.
> Full derivation + the campaign protocol: [`build-budget.md`](build-budget.md).

> **Below 25 GiB free, `/demo-up` warns + offers a reclaim** (override the floor with `DEMO_DISK_MIN_GIB=N`).
> It never blocks the bring-up. Prefer `docker builder prune -f --filter until=24h` over
> `docker system prune -af`: `until=` drops the cache records **not used within** the window — which is
> exactly the once-only `pnpm install` / `COPY . .` entries (16 of them, 4.029 GB each, **61 % of the whole
> cache**) — while keeping the base layers every build reuses. **The companion fix:
> `rosetta-demo down <N> --purge` removes that stack's images** (`demo-N-*`, scoped so it never touches
> another demo or a dev/base image) — so tearing a demo down with `--purge` reclaims its disk. A **plain
> `down`** still *keeps* the images (a fast re-up); `--purge` is the "I'm done, reclaim everything" path (it
> already dropped volumes + the data dir).
>
> **⚠️ Volumes: every teardown path now passes `-v`, and it took two fixes to be true (v2.8 M258).** The
> bitnami Postgres image declares **three** `VOLUME`s while compose binds only `/bitnami/postgresql`, so
> **every container start mints two anonymous volumes** that any teardown without `-v` orphans —
> **measured at 178 dangling volumes / 5.297 GB over five days across three stacks.** `--purge` was
> exonerated by measurement (it already ran `down -v --remove-orphans`); the **plain `down`** was the
> producer, fixed at M258 iter-14 and carried to the `dev-stack` twin at iter-17.
>
> The M258 close closed the third path: the **label sweep**. When compose *refuses* — a stale generated
> override naming deleted services makes the merged project invalid, so compose touches nothing and every
> container survives — the containers are removed by `docker rm` instead, and **`docker rm` without `-v`
> orphans the anonymous volumes just the same.** That is the branch the sweep exists for, so the leak was
> still fully open on exactly the failure mode that motivated the fix, and now invisible, because the
> plain-down path was believed fixed. Both twins now run `docker rm -fv`.
>
> **Why `-v` is not destructive here, stated because a bare flag reads as one:** a live census of every
> mount in a demo project found the **only** volume-type mounts in a whole stack are those two anonymous
> ones — there are **no named volumes to lose** — and the database is a **host bind mount**, which `-v`
> cannot touch. Fenced in both directions (`dev-stack/tests/test_dev_teardown_sweep_m258.py`), so the day
> the platform adds a named volume the decision re-opens instead of silently becoming destructive.

> **The free-space signal measures the Docker VM's INTERNAL disk, not host `/` (v2.6 M239-F1 correction).**
> On Docker Desktop the engine runs in a Linux VM with its **own fixed-size virtual disk** — the filesystem a
> full-cold build actually `ENOSPC`s. Host root `/` is a *different*, usually-huge filesystem that does **not**
> reflect it. The original pre-flight `df -Pk /`'d host root and so read ~200 GB "free" while the VM's ~59 GB
> disk filled and the build died — staying **GREEN through the exact failure it exists to catch**, which then
> surfaced as a **cryptic downstream `redis exited (1)`** (redis was just the first container to write). It now
> probes the VM disk via a throwaway `busybox df` (the container root == the VM overlay), falls back to host
> `/` only when Docker/`df` is unreachable, and the warn **names the redis mis-attribution** so the cause reads
> as *disk*, not *redis*. (Measured live: 25 GiB VM-disk free vs 212 GiB host-`/` free on the same box — the
> old proxy's blind spot.) `DEMO_DISK_AVAIL_KB` still short-circuits the probe (test/operator override).

## How the pk + URLs are baked (zero platform edit)

| App | URLs | Clerk pk | Context trim |
|-----|------|----------|--------------|
| **next-web** | `--build-arg NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` / `_BACKEND_API_URL` / `_HOSTING_URL` (offset) — ARGs the Dockerfile already declares | **no pk ARG exists** → dropped into a **gitignored `apps/web/.env.local`** in the build context, read by `next build`, removed by a trap after | the repo ships **no** `.dockerignore`, so a **tooling-owned** one (`rosetta-extensions/demo-stack/frontend/next-web.dockerignore`) is applied **transiently** (never clobbers a repo one; trap-removed) to trim the 2.8 GB context (2.5 GB `node_modules`) to <100 MB |
| **studio-desk** | ⚠️ **THIS ROW DESCRIBED THE PRE-MIGRATION BUILD AND EVERY MECHANISM IN IT IS GONE (corrected 2026-08-23, when the Next migration merged to `main`).** It is now **nine `NEXT_PUBLIC_*` `--build-arg`s** against the **rext-owned** `frontend/studio-desk.Dockerfile`, which re-declares them in both the build and runner stages (`up-injected.sh` passes all nine). The `VITE_*` names are retired — and docker discards an undeclared build-arg **silently**, which is why the rename had to move the Dockerfile and the invocation together. **The five source patches are RETIRED, not applied**: they are sha-pinned to `app/core/{main.ts,scaffold/*}`, which the migration deletes, so demopatch refuses them at G2 and `build_frontend_studio_desk` writes RETIRED evidence instead of running the ladder. The prod-eject they fixed is re-landed in SOURCE via `app/_lib/externalUrls.ts`. **The `.env.production.local` overlay, its transient `.dockerignore` re-include and the RETURN trap are all deleted** — the pk and the cockpit URL are real declared ARGs now, and a `--build-arg` cannot be stranded the way an overlay file could. See [`demopatch-spec.md` §8](demopatch-spec.md). | **`NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` IS a declared ARG** → the minted pk passed straight as a build-arg (so the app derives the same fake-FAPI host the backend talks to) | the repo **already ships** a `.dockerignore` excluding `node_modules`/`dist`/`.git` — left untouched; nothing transient is added to it any more |

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

> **Baked URLs with no per-URL override → the demo-patch tool (M42m).** The build-arg / `.env.local` injection
> above rewrites a baked URL only when next-web exposes a per-URL `NEXT_PUBLIC_<thing>_URL` knob for it (as
> `ACADEMY_URL` does via `NEXT_PUBLIC_ACADEMY_URL`). The left-nav **Studio** link has none — `STUDIO_URL` is a
> `NEXT_PUBLIC_NODE_ENV` ternary (`localhost:9000` | prod), wrong-port + side-effecting on flip — so it baked
> `studio.anthropos.work` into the manager nav (a prod-eject escape, 139×). The fix keeps the zero-platform-edit
> line: a **tooling-owned demo-patch** (`rosetta-extensions/demo-stack/patches/demopatch` + a content-anchored
> manifest) source-patches the demo's **EPHEMERAL gitignored next-web clone** to read `NEXT_PUBLIC_STUDIO_URL`
> (a behavior-identical fallback ternary kept) **before** the image build, then **trap-reverts** after it bakes —
> CANONICAL repos never touched (**7** guards — G7, the apply post-condition, was made real at the M217 close: demo-clone-only path-assert, drift-refuse, never-commit, idempotent,
> self-owned reversal, demo-only). Wired into `up-injected.sh` (apply-before-build + RETURN-trap revert, exactly
> like the pk overlay), with the offset value `http://localhost:39000` in the `.env.local` overlay; default-on +
> non-fatal (`DEMO_NO_PATCH=1` opts out). The Studio escape resolved demo-only (139→0); the served bundle carries
> 0× prod / 31× `:39000`. Full mechanism + the failure-mode routing table (the "Platform-bound escape" row):
> [`coverage-protocol.md`](coverage-protocol.md).
>
> **Re-anchored to the current source (v1.10b M49 #8).** The manifest's `pre_sha256`/`post_sha256` pin the
> whole-file hash, and the M47 re-sync moved next-web to **v2.89.0** — so the hashes (pinned to the v1.10 ref)
> would have made G2 **drift-refuse** the pristine current file, leaving the Studio link prod-baked. The
> `STUDIO_URL` hunk itself is byte-identical to v1.10; only the file-level hashes moved (sibling exports drifted
> — `AI_READINESS_URL`, the `/enterprise/*` URLs, the member-profile regexes). M49 recomputed both hashes from
> the v2.89.0 source and verified the apply→revert cycle against it.

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

> **Remote / `--public-host` demo (v2.2 M214).** On a Tailscale-reachable demo the browser's origin is
> `https://$HOST:<offsetport>` (M213's per-port `tailscale serve` fronts each offset port with HTTPS,
> **preserving the port**), so the override **appends** the HTTPS MagicDNS origins while **keeping** the
> `localhost` trio for on-host use:
> ```
> CORS_EXTRA_ORIGINS=http://localhost:13000,http://localhost:13001,http://localhost:19000,\
>                    https://billion.taildc510.ts.net:13000,https://billion.taildc510.ts.net:13001,https://billion.taildc510.ts.net:19000
> ```
> A **single scheme predicate** (`browser_scheme` in `gen_injected_override.py`, mirrored by `$SCHEME` in
> `up-injected.sh`/`ant-academy.sh`) drives the http→https flip for **every** browser-facing surface — the CORS
> origins, the studio-desk `CLERK_SIGN_IN_URL`/`WEB_APP_URL` requireAuth fallback, all the baked
> `NEXT_PUBLIC_*`/`VITE_*` endpoints, and the cross-surface links — so there is **no plain-http browser call**
> under HTTPS-everywhere (mixed-content clean; the asset plane stays prod-HTTPS). Unset host ⇒ byte-identical to
> the localhost block above. studio-desk's sign-in gets a `NEXT_PUBLIC_CLERK_SIGN_IN_URL` bake — a **real declared
> Dockerfile ARG** since the Next migration, so the gitignored `.env.production.local` overlay this line used to
> describe is gone (and studio-desk is no longer an SPA) — and ant-academy's `next dev` `allowedDevOrigins` admits
> the MagicDNS host via the `ant-academy-dev-origins` sha-pinned patch. The full recipe + topology:
> [`tailscale-serve.md`](tailscale-serve.md).

This is emitted by `gen_injected_override.py` (the `backend` service gets an additive `environment:` block), so it
applies to a stack brought up **through the demo injected override** (`/demo-up`). The **dev** override
(`stack-core/gen_override.py`) does **not** emit it today and the dev bring-up runs no UI tier — so a `dev-N`'s
offset frontends would still be CORS-blocked if you ran them (a known gap, not yet wired on the dev side). It is
**not** the same as next-web's *server-side* SSR `fetch` origin
(that's the build-time `NEXT_PUBLIC_*` URLs above + the absolute-internal-origin item in §"What's out of scope");
CORS is specifically the **browser→backend** allowlist. With it set, the offset origin gets its `ACAO` header
and the REST-backed dashboards load.

## ant-academy — native, Clerkenstein-wired, session-detached, with a documented fallback

ant-academy is **Vercel-native** (not in docker-compose). `/demo-up`
launches it natively on `:3077+offset` **Clerkenstein-wired** — the demo's minted publishable key + the
disarmed fake BAPI, read from `<stack>/.env.demo-N` — paired with `REQUIRE_ORGANIZATION_MEMBERSHIP=0` to skip
the org gate. The per-demo env is a **gitignored `code/.env.local`** overlay (zero academy-repo edits).
Launching it natively (vs only documenting the step) resolved the overview's open question toward "launch it,
fall back if fiddly" — the academy is Vercel-native, not cleanly dockerizable (#M19-D6).

> ### 🔻 RETRACTED — *"depends only on Clerk at runtime" / "Clerk-only"* (M257x)
>
> Both spellings stood in this section, and both were **false** — and self-contradicted three paragraphs
> later, where this same file already says the catalog is *"DB-authoritative [read from the platform academy
> subgraph over GraphQL]"* and names the missing endpoint as the **root cause** of the empty grid. One
> document, two readings.
>
> **Measured at `stack-demo/ant-academy` @ `22df69dd8`:** the academy has a **SECOND** runtime dependency —
> the platform GraphQL endpoint, read from `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`:
> - `code/src/graphql/server.js:14` reads it and **throws** *"NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT is not set"*
>   when absent (`:18`); `code/src/graphql/useGraphql.js:18`/`:33` is the client-side twin.
> - The **course catalog** goes through it: `code/src/lib/serverTenant.js:145`
>   `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView()`, and
>   `code/src/lib/backendContent.js:94-100` is the GraphQL request pair.
> - So do the **chapter body** (`backendContent.js:146-149`), the **skill path**
>   (`:175-178`), the beacon route (`code/app/api/academy/beacon/route.js:36`) and certificate verification
>   (`code/app/api/verify/[certId]/route.js:48`).
>
> **This is not a footnote — it is the mechanism of the demo's own F4 defect.** A demo sets that variable
> **0×**, so every one of those reads returns null and the grid falls to `emptyCatalogView()`. The
> `academy-fs-published-*` demo-patch family below exists *because* the academy is not Clerk-only. Also
> retracted downstream in [`content-stories-routes.md:379`](content-stories-routes.md), which already calls the
> premise stale. *(A third runtime dependency, the AI keys behind `/api/ai/chat`, is documented in
> [`../../services/ant-academy.md`](../../services/ant-academy.md).)*

> **The demo academy is AUTHENTICATED via real Clerkenstein keys (v2.3 M220 S5/i).**
> ⚠ **This section described the superseded KEYLESS model until M257x iter-98**, three paragraphs after `:48`
> and `:84` of this same file had already recorded its removal — one document holding both readings.
> Measured at rext `main`, the launcher sets **neither** `BENCHMARK_VISUAL_BYPASS=1` **nor**
> `NEXT_PUBLIC_E2E_AUTH=1` (`demo-stack/ant-academy.sh:576-583`, fenced by two tests). A hero who clicks
> through from next-web arrives **signed in as herself**, not as a synthetic `E2E Member` — which is the
> whole point of the change: `proxy.js` short-circuits on the persona cookie *before* resolving the real
> session, so keeping the bypass alongside real keys would render "E2E Member" over a presenter logged in as
> Maya.
> > **The cockpit still SETS the cookie, at two paths** (`demo-stack/cockpit.py:855` client-side and `:1539`
> > as a `Set-Cookie` on the `/go` 302) for the content-stories academy deep-link — but with the launch-env
> > bypass gone it is **inert** on a stock demo. The retracted claim that the cockpit "no longer sets" it was
> > false in both directions: the cookie is set, and it is not honoured. (The academy grid rendering **empty**
> > in a demo is the v2.4 **F4** carry — **NOT** a
> > client-side render defect: the catalog is **DB-authoritative** [read from the platform academy subgraph over
> > GraphQL], and a demo neither sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` nor holds academy rows → `emptyCatalogView()`
> > = 0 cards. Root-cause + read-chain: [`../../services/ant-academy.md` § The Content Model](../../services/ant-academy.md#the-content-model--db-authoritative-catalog-v051-m7).
> > **v2.5 M230 fills it production-faithfully, zero academy-repo edits.**)
>
> The `e2e_persona=member` cookie (formerly set browser-side by the cockpit's [Academy] link before it
> navigated to the academy origin) drove the authenticated context. Cookies on `localhost` are
> **port-agnostic** (RFC 6265 ignores the port), so the cookie the cockpit origin (`:7700+offset`) set was
> read by the academy origin (`:3077+offset`) — **no academy-side route + no academy-repo edit**. So a hero
> who walked in from the cockpit landed **authenticated as a member** (a non-anonymous academy session), not
> as an anonymous visitor. Without
> the cookie the portal still opens for anonymous browse (the flags enable the bypass; the cookie chooses the
> persona). The academy identity is the synthetic `E2E Member`, **not** the exact seeded platform hero (the
> academy's only platform-backend link is the **GraphQL catalog read** — tenant-filtered, not identity-resolving — so
> it can't map the Clerk session to a seeded platform user) — the F6 bar is "authenticated, not anonymous", which
> `member` (signed-in + org + entitled) satisfies.

### The empty grid is FILLED — the `academy-fs-published-fallback` demo-patch (v2.5 "the playbill" M230)

> **⚠ Scope of this closure (M236, proven live).** The **authenticated** grid is filled and verified cold on
> `billion`: **65 course cards, 483 chapter links, 0 Draft chips**. The **anonymous** `/library` and `/free`
> routes still render **0 cards** — `getPublicCatalogView`'s `new Set()` branch is not covered by this patch
> (the patch manifest names the gap itself). Tracked as `ACADEMY-M236-iter08-public-catalog-twin` → v2.5
> release close. Read "the empty grid is FILLED" as *signed-in*, not *everywhere*.

The F4 carry — a demo academy grid rendering **0 cards** — is closed by **Option C**: a sha-pinned rext
demo-patch (`demo-stack/patches/academy-fs-published-fallback`) on the demo's **own ephemeral ant-academy clone**
that restores an **FS-as-PUBLISHED catalog fallback**. **Zero canonical-repo edits.** (Option B — a firewalled
academy-content snapshot surface: prod-capture the public academy rows + replay into the demo app DB + wire
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` + compose the academy subgraph into the demo router — was weighed and NOT
chosen: it needs a prod-DB read + a new snapshot surface + subgraph composition, i.e. far more cold-`/demo-up`
infra risk. Option C is the least-infra-risk path to the gate.)

**Why the grid was empty (the F4 root cause, code-confirmed — NOT a client render defect).** Since ant-academy
v0.5.1 (M7) the home grid's catalog is DB-authoritative: `serverTenant.js::getServerCatalogView()` is
`const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView(); return draftsEnabled() ? mergeDrafts(view, eids) : view`.
A demo sets **no** `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (`ant-academy.sh` sets it **0×**) and the demo app DB holds
no academy rows, so `getBackendCatalogView` returns null → the grid resolves to `emptyCatalogView()` → 0 cards.
The M7 cutover deliberately **removed** a pre-existing FS-as-published fallback at exactly the `?? emptyCatalogView()`
expression.

**The fill.** The demo-patch restores that fallback, **env-gated on `ACADEMY_DEMO_FS_PUBLISHED`** so it is
behavior-identical to pristine when unset (safe even if upstreamed). When set, the fallback reuses the
already-tested `mergeDrafts(emptyCatalogView(), eids)` — which returns the full public FS catalog — and **strips
the `_draft`/`_origin` tags** from chapters/series/skillPaths, so the cards render through the **unchanged** RSC →
`AcademyClient` → `SkillPathCard` chain as **PUBLISHED — NO "Draft" chip** (the chip is `skillPath?._draft === true`;
Option A — the `ACADEMY_SHOW_DRAFTS` draft layer — was REJECTED precisely because it stamps that chip). This is the
gate's sanctioned **"faithful equivalent"** of the DB-authoritative GraphQL path.

**How it's applied** — like `ant-academy-dev-origins`, by a native rext helper
(`stack-injection/apply-academy-fs-published.sh apply|revert`), **NOT** the image-baked `demopatch` tool, because
ant-academy runs natively via `next dev`: **apply-before-launch, revert-on-`--stop`**, idempotent + drift-refuse +
single-occurrence + post-condition re-check. **Default-on** for every demo; **`DEMO_NO_ACADEMY_FILL=1`** opts out;
**non-fatal** (a refused patch leaves the grid empty — the documented residual). The canonical
`anthropos-work/ant-academy` repo is never touched (only the demo's ephemeral gitignored clone). Shipped in rext
at tag `playbill-m230-academy-fs-published`.

**Proof.** Runtime-proven standalone (M230 iter-02): the patched academy served the home grid with **59
skill-path cards** (real catalog names — Claude Code, AI Foundations, Agent SDK, AI Engineering, Business, …),
**0 `draft-ribbon` / 0 `data-draft="true"`**, and the clone reverted **byte-clean**. The **formal gate** — the
coverage sweep's `ANT_ACADEMY` rendered-card count on a **cold `/demo-up`** ([`coverage-protocol.md`](coverage-protocol.md))
— is the remaining release-close verification.

### The BODY half — `academy-fs-published-chapter-body` (v2.6 "sound check" M238)

The catalog patch above got the **grid + course landing** to render, but clicking **"Start the course"** still
**404'd** — because chapter **bodies** are backend-authoritative too, and the catalog patch only touches
`serverTenant.js` (the catalog), not `serverChapterBody.js` (the body). A demo's null backend → `notFound()` → the
"You wandered off the trail" 404 (see [`../../services/ant-academy.md`](../../services/ant-academy.md) §"The chapter
BODY is backend-authoritative too"). **M238 adds the BODY half:** the `academy-fs-published-chapter-body` demopatch
serves the committed FS chapter body (locale-aware, unlocked, un-chipped) at the backend-null branch — gated on the
**same** `ACADEMY_DEMO_FS_PUBLISHED` env var + **same** `DEMO_NO_ACADEMY_FILL` opt-out, applied together by
`ant-academy.sh` via the sibling native helper `stack-injection/apply-academy-fs-published-body.sh`
(apply-before-launch / revert-on-`--stop`; behavior-identical when the env is unset). So the two halves are one
coherent FS-as-published behavior: **the grid renders FS cards, and clicking one renders the FS body.** Shipped in
rext at tag `sound-check-m238-ant-academy-reliability`. **Proven live on `billion`** (demo-1): a chapter that
returned **HTTP 404 "Not Found"** now returns **HTTP 200** with the real chapter title + body, and
`/chapters/<slug>/?lang=it` also renders (the language switch on a chapter reader — the same backend-null path). The
coverage sweep now also fences the **chapter body** + the **`?lang=it` re-render** (`ANT_ACADEMY_CHAPTER_SECTION`,
[`coverage-protocol.md`](coverage-protocol.md)), not just the home grid. (#M238-D1)

> **⚠️ All five native-run academy patches are applied by SHELL HELPERS, and until v2.8 M258 those helpers
> hard-refused on any whole-file drift** — so an ant-academy version bump silently disarmed them while the
> bring-up reported success. That is how a demo shipped an academy that rendered perfectly and could not
> hydrate (`ant-academy-dev-origins`), and it is one bump away from an empty grid (`academy-fs-published-*`).
> They now share one **self-healing** ladder, `stack-injection/live_patch_ladder.py`: the *anchor* is the
> contract, the whole-file sha only a baseline. Read
> [`demopatch-spec.md`](demopatch-spec.md) § "The gate" before adding a sixth helper — and do not copy a
> ladder again. Live-proven on `billion`: all five apply in chain order, each re-applies as a no-op, all
> revert in reverse, and the clone comes back byte-identical.

> **Known limitation — the five native-run academy patches share one clone (concurrent-demo teardown).** All five
> `ant-academy` patches (`ant-academy-dev-origins`, `academy-fs-published-fallback`, `academy-fs-published-public`,
> `academy-fs-published-chapter-body`, `ant-academy-back-to-cockpit`)
> are applied to the **shared** `stack-demo/ant-academy` working tree — its path is `N`-independent (only the port +
> pidfile are per-`demo-N`). So `ant-academy.sh N --stop` reverts the shared source files unconditionally: tearing
> down `demo-1` while `demo-2`'s native `next dev` is still live reverts the patched files out from under `demo-2`,
> and its next HMR recompile re-404s the chapter route. This is a **pre-existing property of the native-run academy
> pattern** (not introduced by M238 — the M238 body patch merely follows it); it only bites when **multiple demos run
> concurrently against the same box**, the uncommon case. A proper fix (per-demo academy clone, or an applied-refcount
> before revert) is routed to the standing backlog (M238-D6); the single-demo path is unaffected.

> **The academy AI chat (Cosmo) is absent in the demo — by design (M53 F6, per the AI-keys policy).** The
> academy's Cosmo assistant is gated behind `NEXT_PUBLIC_FEATURE_TRAINING_COACH` (default **OFF**) **and** a
> per-user `localStorage('openai_api_key')`. The demo launcher sets **neither** the flag nor any OpenAI key —
> the demo provisions **no** AI keys (the same AI-keys policy that keeps the `/api/ai/chat` route unexercised) —
> so Cosmo is genuinely **absent** in a demo academy. This is intentional: the F6 acceptance makes **no**
> `/api/ai/chat` assertion. Course content + the authenticated browse experience are the demo surface; the AI
> assistant needs keys the demo deliberately doesn't carry.

> **The native daemon is SESSION-DETACHED (the M33 "dead on a later visit" fix).** ant-academy was previously
> launched with `nohup` alone — which does **not** detach from the launcher's process group. So when a
> backgrounded `/demo-up` task's process tree was reaped on completion (or the launching session ended), the
> academy daemon died with it: the stack looked healthy at bring-up but was **dead on a later visit** (the exact
> M33 hypothesis — now **reproduced and fixed**). The launcher now starts it **session-detached** via a shared
> `demo-stack/detach.sh::launch_detached` (`setsid` where present; a portable `python3 os.setsid` double-fork on
> **macOS**, which has no `setsid`), so the daemon **survives the launching session/task ending**. _(The
> **presenter cockpit** host-native daemon had the identical bug and got the **same** `launch_detached` fix.)_

> **ant-academy auto-installs its deps — no token needed (the storytelling-postfix-2 "down in the demo" fix).**
> ant-academy **does** use **Font Awesome Pro** icons, but the FA Pro assets are **self-hosted / vendored in the
> repo** (`code/public/assets/fontawesome/webfonts/*.woff2` + `css/all.min.css`, used as `<i class="fa-solid …">`)
> — they are **not** pulled from the Font Awesome npm registry, so `npm install` needs **no** token (a fresh
> token-less install, no `.npmrc`, completes in ~14 s / 750 pkgs and the launched app serves HTTP 200 with working
> FA icons). The `FONTAWESOME_NPM_AUTH_TOKEN` in `code/.env.example` is **vestigial** — not required to install or
> run. The real "ant-academy down in the demo" cause was a **blocked clone**: an empty `stack-demo/ant-academy/`
> stub (holding only a gitignored `code/.env.local`) tripped `make init`'s skip-if-present, so the source never
> landed. **Fixed at `storytelling-postfix-2`:** `ensure-clones.sh` now **sweeps incomplete sibling stubs** (any
> `repos.yml` repo dir with no `.git`) before `make init`, and `ant-academy.sh` **auto-runs `npm install`** (no
> token) when `node_modules` is absent — so a fresh `/demo-up` now brings ant-academy up **automatically** (proven
> live on `:33077`).
>
> ⚠️ **This closed ONE failure of skip-if-present, not the class — do not read it as "skip-if-present is now
> safe" (`F-M236-CLOSE-1`).** The stub-sweep fixes the case where a repo dir exists with **no `.git`**
> (nothing was ever cloned). The far commoner case — a **complete, healthy clone that is simply out of
> date** — still passes the sweep untouched, because `make init` never fetches, pulls, or checks out an existing
> clone and the bring-up never calls `make pull` unless you opt in. **What changed at v2.6 M237:** the bring-up
> now runs a **fetch-verified freshness assertion** that MEASURES this and warns loud (advisory; `DEMO_FRESHNESS_STRICT=1`
> escalates to fatal), and an opt-in `DEMO_ADVANCE_CLONES` advances the clones — so the visibility gap is closed.
> (The v2.5 draft cited `app` **249** / `next-web-app` **202** commits behind "identically on both boxes" — that
> reading was itself the **suppressed-fetch artifact** M237 eliminated; the verified measurement on `billion` was
> 0–2 behind, frontend current.) See
> [`../rosetta_demo.md` § Clone freshness](../rosetta_demo.md#clone-freshness--the-fetch-verified-assertion-f-m236-close-1-closed-v26-m237)
> — **that gap is now closed** (visibility, not auto-advance), and a stale-looking demo is what an un-advanced clone usually is.

**Default-on + non-fatal + degrades to a documented step.** A fresh `/demo-up` clones the academy (via the
`storytelling-postfix-2` stub-sweep) and auto-runs the token-less `npm install` (see above), so it comes up
automatically. If anything still trips it — Node < 22, or a genuinely unavailable clone — the tool prints the
exact manual commands and continues, never aborting a good demo bring-up:

```bash
cd stack-demo/ant-academy/code            # M26: the academy clone lives in the demo's OWN peer set (stack-demo)
cp .env.example .env.local                 # gitignored; keeps the repo clean
#   set REQUIRE_ORGANIZATION_MEMBERSHIP=0 and NEXT_PUBLIC_E2E_AUTH=1 (M53 F6 authenticated-member session);
#   Clerk keys are optional (keyless works — no FA token needed either, FA Pro is vendored)
npm install
BENCHMARK_VISUAL_BYPASS=1 NEXT_PUBLIC_E2E_AUTH=1 npm run dev -- --port 23077   # demo-2: keyless
#   then set the e2e_persona=member cookie (the cockpit [Academy] link does this) to land authenticated.
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
at the **pinned tag recorded in `.agentspace/rext.tag`** (the single source-of-truth, M49 #1 — see
[`rosetta_demo.md`](../rosetta_demo.md) *"The pin is a file"*; current v1.10b "fit-up" pin: `fit-up-m49`).
*Landing provenance (which historical tag first shipped each piece):* the M19 UI tier first shipped at
`dress-rehearsal-m19`; the CORS + token-strip items were later, ≥ `dress-rehearsal-m20-fix15`/`fix17`; the
session-detach fix below landed at `storytelling-postfix-1`; the academy stub-sweep + token-less auto-install
landed at `storytelling-postfix-2`:

- `stack-injection/gen_injected_override.py` — appends the two frontends to the injected override (offset
  `ports:!override`, `image: demo-N-*` + `build:!reset null` + `pull_policy:never`, `mem_limit:1g`,
  `profiles:!override [<profile>]` — the profile is **derived** from the platform clone via `profile_for()`
  (`gen_injected_override.py:420`), never a literal: `core` at platform `0dab54d`); `--no-ui` clears the
  tier. Also sets `CORS_EXTRA_ORIGINS` on the **backend**
  service to the offset frontend origins (see §"Offset-origin CORS"), and **strips the inherited prod
  `DIRECTUS_TOKEN`** (`DIRECTUS_TOKEN=`) on **every** emitted service + both frontends — no prod credential rides
  in a demo container, and studio-desk's prod-Directus *write* path is disarmed (fix16/fix17; see
  [`../safety.md`](../safety.md) §2.3 + §2.2). **(v2.7 M252, the F8 gap.)** studio-desk is a **base-compose**
  service, so in a demo it inherits **only `platform/.env`** — which carries **no AI-provider keys**, so the
  studio backend 500'd `POST /api/ai/completion`. `frontend_lines()` now also emits an **existence-guarded
  `env_file: [<clone>/studio-desk/.env]`** on the studio service (the studio-desk clone sits beside `platform/`),
  layering the clone's own `.env` over `platform/.env` to supply the studio's own **AI-provider keys**
  (`AI_OPENAI_API_KEY` + `AI_ANTHROPIC_API_KEY`) so `/api/ai/completion` no longer 500s. **No `MOCK_CLERK`, no
  auth change** — the demo studio's auth model is unchanged (Clerkenstein; the injected override even asserts the
  studio-desk block carries **no** `MOCK_CLERK` line — see the studio-desk block above). Precedence is
  preserved: the explicit `environment:` block still wins (the Clerkenstein `CLERK_*`, the stripped
  `DIRECTUS_TOKEN`, `NODE_ENV=production`); `env_file` lists **concatenate**, so studio-desk/.env keys win over
  platform/.env.
- `demo-stack/up-injected.sh` — the per-demo serial-before-up frontend build (offset URLs + minted pk +
  tag-guard), the 12 GB VM pre-flight, the `--no-ui` (`DEMO_NO_UI`) escape, the scoped verify.
- `demo-stack/frontend/next-web.dockerignore` — the tooling-owned context trim for next-web.
- `demo-stack/ant-academy.sh` — the native academy launcher / stopper / documented fallback; **auto-runs the
  token-less `npm install`** when `node_modules` is absent (FA Pro is vendored — no token needed), launches the
  daemon **session-detached** (via `detach.sh::launch_detached`), and **non-fatally** prints the manual commands
  only if a genuine blocker remains (e.g. Node < 22).
- `demo-stack/detach.sh` — the shared `launch_detached` helper (`setsid`, or a `python3 os.setsid` double-fork on
  macOS) that session-detaches the host-native daemons (ant-academy **and** the presenter cockpit) so they
  survive the launching `/demo-up` session/task being reaped.
- `demo-stack/lib/rext_tag.sh` — the shared reader for the `.agentspace/rext.tag` consumption-tag source-of-truth
  (v1.10b M49 #1); both `/demo-up` and `ensure-clones.sh` source it to resolve the pinned rext tag. Picks the first
  non-comment / non-blank token, strips a trailing CR so a CRLF-edited pin still resolves as a clean git ref.

## What's out of scope (the user-owned follow-up)

**True zero-rebuild** — one frontend image that serves every stack with the port/pk switched at *runtime* —
would require **platform-source changes** (runtime rewrites in `next.config.mjs`, an absolute internal origin for
SSR `fetch` in `server.graphql.ts`, a `window.__ENV` shim + explicit `publishableKey` on `<ClerkProvider>`).
Those are real platform edits with PR/review/prod risk — **forbidden** for the demo tooling to make locally.
It's documented here as an **optional upstream PR you own** (a deferred/unscheduled deploy-CI precedent),
**not built** in M19. The honest residual above (one cached build per new `demo-N`) is the accepted cost of
staying tooling-only.

> **⚠️ Two of the four items above have since been delivered WITHOUT an upstream PR, and the list is now
> re-cut — with ACHIEVED numbers (M257 close, D-v28-10 / D121: one rewrite, not two).**
> - **The SSR origin** landed at v2.3 **M218** as the sha-pinned `next-web-ssr-graphql-origin` demo-patch —
>   shape 2, not a PR. It was worth 37.5 s of every authenticated render.
> - **`output:'standalone'`** was listed here as PR-only. **It is not, and it is now SHIPPED for both Next
>   apps.** Next 16's frozen `defaultConfig` reads
>   `output: !!process.env.NEXT_PRIVATE_STANDALONE ? 'standalone' : undefined`, and no app `next.config` sets
>   `output` — so **`ENV NEXT_PRIVATE_STANDALONE=1` in a tooling-owned Dockerfile (shape 3) is sufficient**,
>   with zero source edits and zero demo-patches. M257 iter-09 landed it as multi-stage builds in
>   `demo-stack/frontend/{next-web,hiring}.Dockerfile`, `next-web` moving shape 1 → shape 3 to get there.
>   (Gotcha, still load-bearing: turbo 2 defaults to `--env-mode=strict`, which filters the variable out
>   before `next build` sees it, so the flag silently no-ops without `--env-mode=loose` and the build stays
>   green with the old image.)
>
> **The achieved numbers — two hosts, stated separately, because seconds do not transfer.**
>
> | | `billion` (x86_64, containerd) | `macmini` (arm64, containerd) |
> |---|---|---|
> | when / what | **M255 spike (a)** — `hiring.Dockerfile`, a proof of the mechanism | **M257 iter-09** — both Next apps, shipped in the gated bring-up |
> | hiring image | **4.84 GB → 379 MB** | **3.94 GB → 380 MB** |
> | next-web image | not measured there | **4.04 GB → 417 MB** |
> | export leg | **146.8 s → 2.9 s** (hiring alone) | **136.4 s → 3.8 s** (both Next apps combined) |
> | UI-tier phase | 436.1 s pre-lever (n=3 p50) | **246.23 s → 104.60 s** (n=3 p50, **−141.63 s**) |
>
> **What that bought at the cycle level, on `macmini` only:** the cold `demo-down --purge` + `demo-up` p50 went
> **449.51 s → 286.99 s** (n=3, min 280.99 / max 303.44), under M257's **360 s** gate *and* its **300 s**
> stretch, with `autoverify green:true / 0 warnings`, HEADROOM OK and ISOLATION OK on all three reps, host
> identity `match` ×3, **0 platform-repo edits** and 0 refused demo-patches. Both images were proven
> **behaviourally identical** to the ones they replace before the numbers were believed — hiring's `/login`
> is byte-for-byte **426,914** bytes in both. `billion`'s own post-L1 cycle has **not** been measured; do not
> infer it from these.
>
> **And the ranking moved underneath the plan.** With the UI tier collapsed, the largest single phase is now
> **`set_dress` at 82.04 s = 28.6 %** of the cycle — a lever (L5) that had been priced at ~30–50 s and ranked
> **fifth**. Routed as `LEVER-M257-L5-setdress`. Full derivation, the per-phase attribution table and the
> campaign protocol: [`build-budget.md`](build-budget.md).

## Related
- [Demo family index](README.md) · [Lifecycle](../rosetta_demo.md) · [Safety contract](../safety.md) · [Verification](../verification.md)
- [next-web-app](../../services/next-web-app.md) · [studio-desk](../../services/studio-desk.md) · [ant-academy](../../services/ant-academy.md)
