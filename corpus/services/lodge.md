# Lodge (`hyper-studio`)

> **Status: DEPLOYED IN EVERY STACK since v2.10 M272 (2026-08-24).** Until this milestone the corpus
> graded `hyper-studio` *"a CLI suite … **ZERO runtime coupling to the platform today** — no GraphQL, no
> Clerk, no DB, no deployment. KEEP — PRE-INTEGRATION"* ([`org-repos.md`](../architecture/org-repos.md))
> and *"the successor to studio-room, **PRE-INTEGRATION**"*
> ([`platform-migration-status.md`](../architecture/platform-migration-status.md)).
>
> **Those claims were ACCURATE when written and went stale — they were not wrong.** Both were measured
> 2026-08-06/07; hyper-studio's `Dockerfile` landed **2026-08-14** (`717b40e6`, *"the image, the unit,
> and the fifteen-row boot preflight"*) and `deploy/docker-compose.yml` **2026-08-19** (`49216fc3`).
> The repo grew a deployment eight days after the corpus last looked. That is an ordinary staleness
> against a repo averaging ~800 commits a quarter, not a documentation defect — and it is why the rows
> now carry a re-measurement date rather than a retraction.
>
> The narrow part was true then and is true now: it has no *platform* coupling — which is exactly why
> it can be added to every stack without touching the platform's service graph.

## Role & Responsibility

`hyper-studio` (`anthropos-work/hyper-studio`) is the content-creation engine that succeeds studio-room.
**Lodge** is its *deployable*: the service that hosts the HyperForge engine behind an HTTP wire, supervises
jobs as pooled-uid subprocesses, meters their spend, and archives their output.

It is **not** part of the Anthropos platform's service graph. It shares no database, no Redis, no Clerk
tenancy and no GraphQL surface with `backend`. In a rosetta stack it is a self-contained neighbour that
happens to live on the same compose network — which is why integrating it required **zero** platform-repo
edits and adds **no** `depends_on` edge.

## Architecture & Code Map

One Node process, **two listeners**. Both are served from `code/lodge/server.ts`; the container's
entrypoint is `node /app/dist/lodge/lodge-bin.js` (`Dockerfile:391`).

| Path | What lives there |
|---|---|
| `code/lodge/lodge-bin.ts` | process entry — boot, preflight, signal handling |
| `code/lodge/server.ts` | composition root; starts the wire **and** `startPanelServer` (`:2162`) |
| `code/lodge/wire/` | the job API — `routes.ts` (the route table), `wire-server.ts` (`/healthz`, `/readyz` at `:412`) |
| `code/lodge/panel/` | the operator webapp — `panel-server.ts` (bind + auth), `panel-page.ts` (render), `panel-*.ts` |
| `code/lodge/store/`, `archive/`, `exec/` | job store, archive/retention, the pooled-uid supervisor |
| `deploy/` | `docker-compose.yml`, `lodge.env.example`, `lodge.service`, `terraform/` |

### The wire (job API) — `8080` in-container

| Method | Path | Route |
|---|---|---|
| `POST` | `/jobs` | submit |
| `GET` | `/jobs/{id}` | poll |
| `GET` | `/jobs/{id}/result` | result |
| `POST` | `/jobs/{id}:cancel` | cancel |
| `GET` | `/healthz` | liveness — **zero I/O by design**: "this process is up and its event loop is turning" |
| `GET` | `/readyz` | liveness **and** the boot preflight passed. Not a capacity signal |

The four job routes are the table at `code/lodge/wire/routes.ts:177-180`; the two probes are served
outside it (`wire-server.ts:412`), which the source itself flags as a documented past error — the table's
own comment retracts an earlier *"no `/healthz`, no `/readyz`"* claim.

### The operator panel (the "mini lodge webapp") — `7787` in-container

A real server-rendered webapp, not a JSON surface: submit a job, watch it, cancel it, plus live processes,
spend by customer and day, the archive view, and a `/config` readout of every resolved value **with its
provenance**. Measured: ~97 KB of HTML on a stack with no jobs.

## ⚠️ Exposing the panel takes THREE settings, not one

This is the part that is easy to get wrong, and each failure looks like a different bug.

1. **`LODGE_PANEL_HOST=0.0.0.0`.** The shipped default is `127.0.0.1`, which is correct for the production
   posture (reach it over an SSH tunnel) and **unreachable through a published port** — Docker forwards to
   the container's `0.0.0.0` only. This is why hyper-studio's own `deploy/docker-compose.yml` publishes the
   wire and *not* the panel.
2. **`LODGE_PANEL_SECRET`, ≥ 16 characters.** The pair is the opt-in and **half a pair is refused at boot**
   (`panel-server.ts:498`). A short secret on a public interface is refused for the same reason.
3. **`LODGE_PANEL_AUTHORITIES`** — the `Host` allowlist, and **the one that is missing from
   `deploy/lodge.env.example` entirely**. Bound to the wildcard `0.0.0.0` the panel would answer only the
   literal authority `0.0.0.0`, which no browser or proxy ever sends — so on a current image it **REFUSES
   TO BOOT** rather than serve something unreachable: `boot.refused`, **exit 78**, and the whole container
   goes with it, wire included. The knob (hyper-studio v03.01 M48, `D-M48-04`) *replaces* the derivation
   from `host`. Reported upstream: their shipped profile alone cannot expose the panel.

> ⚠️ **This entry described a running panel that 403s everything, and that was WRONG for a current image
> — corrected 2026-08-24 by running it.** That IS what the pre-M48 image on this box did (`lodge:m46-sim2`,
> built 2026-08-19: it booted, challenged, accepted the credential, then refused every request with
> *"this panel is bound to '0.0.0.0' and answers that authority only"*). hyper-studio `a3430baf` (M48 S2)
> turned that unreachable-by-construction state into a **boot refusal**, which is strictly better — loud
> and fail-closed instead of quietly serving 403s. Measured on `lodge:rosetta-0a9eb16b`: wildcard host +
> a valid secret + no authorities → `exited exit=78`, both ports dead.
>
> **A runtime 403 is still reachable — just not this way.** With an authority list present but not
> matching what the client sends (a wrong port, a DNS name, a reverse proxy), the panel boots and then
> refuses that client 403, its own probes included. So: **missing list ⇒ dead container; wrong list ⇒
> live container, 403 for the mismatched caller.** Different failures, different first move.

**The login flow is the browser's own.** A `401` carries
`WWW-Authenticate: Basic realm="Lodge operator panel"` (`panel-server.ts:1390`), so the browser prompts and
stores the credential itself. The secret travels in a header in three spellings — `X-Lodge-Panel-Secret`,
`Authorization: Bearer …`, or `Authorization: Basic …` **with the secret as the password** (the username is
ignored; there are no user accounts). ⚠️ `lodge.env.example:271` still says *"there is no browser login
flow"*; that comment predates `D-LS0821-BROWSER-AUTH` (2026-08-21) and is **stale**.

A 401 and a 403 mean different things here and the difference is the whole debugging story:
**401 = the credential is wrong. 403 = the credential was accepted and the `Host` authority was refused.**

## The boot preflight is FAIL-CLOSED — 16 rows

Lodge refuses to open its listener rather than accept a job it cannot finish. Every row names a failure
that would otherwise surface *after* admission, inside a job a caller is already being billed for. Two
rows bite on a stack:

- **the type corpus** — `code/hyper-artifacts/` must be present *and conformant*. A **version-skewed**
  corpus fails it: an older image's gate against a newer mounted corpus refused with
  `corpus-not-conformant-ai-simulation … (nav.action-over-budget)`.
- **row 16, drain grace** — `LODGE_DRAIN_GRACE_SECONDS` must be **shorter** than the D9 lock's 600 s stale
  window, or a draining task releases holds an arriving task has already judged abandoned. hyper-studio's
  own shipped example sets `900`, which **fails this row**; rosetta overrides it to `30`.

## ⚠️ The upstream `.dockerignore` regression (hyper-studio `62f5c597`, 2026-08-23)

That commit added a blanket `*.md` + `**/*.md` to `.dockerignore`. **The type corpus and the seven engine
templates ARE markdown**, and `Dockerfile:267-268` COPYs both. From that commit on, an image built from a
clean HEAD checkout builds successfully and then **refuses to boot**:

```
[row 10] 7 of 7 engine templates do not resolve … forgedir-mak.tmpl.md, run-context.tmpl.md, …
no type corpus was found under LODGE_TYPE_CORPUS_HOME
```

Rosetta works around it **without touching the repo**: `stack-core/lib/lodge.sh` stages upstream's
Dockerfile verbatim into a scratch dir and writes a *derived* `<dockerfile>.dockerignore` beside it —
BuildKit resolves the ignore file next to the Dockerfile before falling back to the context's. The
derivation reads upstream's current `.dockerignore` and appends only two re-admission lines, so an upstream
fix makes our negations redundant no-ops rather than a drifted fork. **This is a bug to report, not a
policy to carry.**

## Local Development (The "How-To")

### In a rosetta stack — automatic, default-ON

Every `dev-N` and `demo-N` stack deploys lodge. No flag is needed.

```bash
.agentspace/rosetta-extensions/dev-stack/dev-stack up 1     # lodge comes with it
```

The bring-up clones `hyper-studio` into the stack workspace (`stack-dev/`, `stack-demo/`), builds
`lodge:rosetta-<short-sha>` (cached per commit — a second stack on the same commit pays nothing), writes
the per-stack env profile, and emits the compose block.

| Surface | Port | Base + `N`·10000 |
|---|---|---|
| wire (job API) | `8080` | `dev-1` → **18080** |
| operator panel | `7787` | `dev-1` → **17787** |

Open **`http://localhost:17787/`**, any username, the stack's secret as the password.

> ⚠️ **8080 was `customerio-sync`'s host port** before it was retired at platform `838d907`. An answering
> `:8080`-family port on a current stack is lodge, not a zombie.

**Opt out:** `DEV_NO_LODGE=1` (dev) / `DEMO_NO_LODGE=1` (demo).

**⚠️ The main dev stack (`dev-0`, `make up` from `platform/`) does NOT get lodge.** That path is the
platform's own Makefile and rext is not in it; this corpus takes zero platform edits. Lodge is on the
`dev-N` (N ≥ 1) and `demo-N` paths only.

### Live engine — developing lodge against a running stack (v2.10 M273)

A dev stack can serve lodge **natively from a git worktree** instead of from its image, so an edit is
live in ~3 s. `LODGE_TYPE_CORPUS_HOME` then points at the worktree's own `code/hyper-artifacts`, so
type-corpus edits are live too.

```bash
.agentspace/rosetta-extensions/dev-stack/engine-switch.sh 1 live    # native, from the worktree
.agentspace/rosetta-extensions/dev-stack/engine-switch.sh 1 baked   # back to the container
```

⚠️ **The live panel takes NO credential** — native lodge binds loopback, where the secret and Host
allowlist above simply do not apply. Same panel, different exposure. Full contract, including what the
switch does and does not migrate: [`../ops/dev-live-engines.md`](../ops/dev-live-engines.md).

### Standalone, from the checkout (no Docker)

hyper-studio ships its own dev launcher, which is the better tool when you are working *on* lodge:

```bash
cd hyper-studio && .claude/skills/serve-lodge/assets/serve-lodge.sh    # tsx watch, ~1s reload
```

### Teardown — and the data loss is REAL

`dev-down N` / `demo-down N` remove the service with the project, and `down -v` takes its named volume
(`<project>_lodge-data`) with it. **That destroys the stack's entire lodge job store, forgespace, archive
and run history.** Intended for a disposable stack; stated because nothing else will tell you.

`lodge-data` is the **first named volume any rosetta stack has ever had**, and it is deliberately not a
host bind: lodge runs each job under a pooled uid (20000+) with its forgespace chowned `0700`, and a macOS
bind mount virtualises ownership — it would hide that isolation rather than enforce it. The M258 teardown
fence was rewritten at M272 to allowlist it semantically (see Testing).

The built image is **not** removed on teardown — it is shared across stacks and is the whole build cache.
Reclaim manually: `docker rmi $(docker images -q 'lodge:rosetta-*')`.

## Credentials — what a stack does and does not give it

**Booting needs no credentials; running a job does.** A rosetta stack supplies none, deliberately: lodge
boots green, serves both listeners, and answers every probe. A submitted job would fail at its first model
call. Lodge meters spend against `HYPERFORGE_CAP_*` ceilings either way (`MAX_BUDGET_USD`,
`PER_RUN_TOKEN_CAP`, `PER_RUN_WALL_CLOCK_MS`, a rolling global window), all inherited from the shipped
profile — every one a `CHOOSE:` line naming the policy question it answers.

> **Measured scope, stated plainly:** boot, both listeners, all four probe surfaces and the panel render
> were verified live on `dev-1`. **An end-to-end job run was NOT exercised** — no credentials, by design.

## Testing

| Check | Where |
|---|---|
| probe rows (`lodge` wire → `http-200 /healthz`; `lodge-panel` → `http /`) | `stack-verify/lib/services.sh` |
| the **authenticated** panel assert (sends the secret, requires 200) | `dev-stack/dev-stack::lodge_verify` |
| registry ↔ platform-compose fence (`lodge:rext-injected`) | `stack-core/service_registry_guard.py` |
| exposure claim (lodge's 2 published ports are in the denominator) | `stack-injection/exposure_claim_guard.py` |
| the named-volume fence (semantic, allowlisted) | `dev-stack/tests/test_dev_teardown_sweep_m258.py` |

⚠️ **The `lodge-panel` probe row is deliberately weak and must not be read as strong.** `http` accepts any
2xx/3xx/4xx, so it passes on 401 (correct) *and* on 403 — a *wrong* authority list, where the panel boots
and then refuses the mismatched caller. It would catch a *missing* list anyway, since that is a boot
refusal and the probe reports 000. The assert that closes the 403 gap is `lodge_verify`, which presents
the credential and demands 200. Keep both; neither alone grades the panel.

## ⚠️ studio-desk submits to it — the FIRST caller in the field (2026-08-24)

Since 2026-08-24 lodge has a real consumer inside a rosetta stack: **studio-desk fires a second
generation at it every time an author presses generate**, beside the studio-room run it has always
made. One simulation design, compiled twice, so the two artifacts can be read side by side. That is
what lodge is here to be evaluated for.

It changes nothing above. The two services still share no database, no Redis and no Clerk tenancy;
the edge is a plain HTTP POST from a studio-desk **route handler**, and it added no compose
`depends_on` and no platform-repo edit.

**Four server-only variables carry the wiring**, provisioned per stack by
`rext dev-stack/engine-switch.sh` (`lodge_desk_env_apply`) at the stack's **own offset ports**:

| Variable | Value on `dev-5` |
|---|---|
| `LODGE_ENABLED` | `1` — the SERVER half of the gate |
| `LODGE_WIRE_URL` | `http://127.0.0.1:58080` |
| `LODGE_PANEL_URL` | `http://localhost:57787` |
| `LODGE_CUSTOMER` | `studio-desk` |

Four things about this are worth carrying:

1. **The gate is a conjunction across two tiers.** Client-side it is `isSuperAdmin`; server-side it
   is `LODGE_ENABLED`. Both must hold. None of the four is `NEXT_PUBLIC_`, deliberately — that
   prefix is inlined at BUILD time, so a container could not be re-pointed at a different lodge
   without a rebuild.
2. **Pointing at the un-offset `8080`/`7787` is the failure to watch for.** On a box running two
   stacks it submits one stack's designs to the other's lodge, and the jobs land in a panel nobody
   is watching. Nothing errors.
3. **`customer` is `studio-desk` — the calling SERVICE, not the tenant.** It is lodge's only
   isolation key and its only identity axis, and lodge authenticates nothing. The org and the user
   ride as ordinary params (`origin_service`, `origin_user_id`, `origin_user_eid`,
   `origin_org_eid`), stamped from the session inside the route where a user cannot edit them.
4. **A lodge failure is visible and non-blocking, by construction.** The submit is detached and
   catches everything; the studio-room generation is unaffected. Proven live on `dev-5`: with lodge
   stopped the route answers `502 lodge_unreachable` and studio-desk keeps serving.

**The `$0` lever, worth knowing before you test anything here:** blank `brief_mode` in the params
and the engine's `grading-axis-unanswered` throws at **provision**, before any model call. A full
submit-and-read-back round trip then costs nothing — `worked 0s`, `spend ≥ $0.0000` — which is how
every gate in this integration was proven without buying a generation.

**The panel reads the submitted documents back** as of the same date. Open a job and the modal
renders the envelope *and* every bone — the cast, the tasks, the criteria — under its own path and
byte size. Before that it showed the envelope alone, and a malformed cast passed inspection.

The full integration is documented on the caller's side, in `studio-desk` at
`knowledge/engineering/backend/lodge-integration.md`.

## Related

- [`org-repos.md`](../architecture/org-repos.md) — the `hyper-studio` row (corrected at M272)
- [`platform-migration-status.md`](../architecture/platform-migration-status.md) — ditto
- [`../ops/safety.md`](../ops/safety.md) §3 — what a stack publishes, lodge's two ports included
- [`../ops/demo/demo-up-defaults.md`](../ops/demo/demo-up-defaults.md) — `DEMO_NO_LODGE`
- [`studio-room.md`](studio-room.md) — the predecessor lodge succeeds
