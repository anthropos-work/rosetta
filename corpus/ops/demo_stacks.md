# Disposable Demo Stacks

**Status:** v1.1 "show floor" / M3 · **Last updated:** 2026-06-03 · **Tooling:** `anthropos-demo/demo-stacks/` (gitignored) · **Skills:** `/demo-up`, `/demo-down`, `/demo-status`

> Spin up `demo-1`, `demo-2`, … as **isolated, full Anthropos stacks on one box**, each Clerkenstein-wired
> so it runs **without real Clerk**, killable cleanly — **without modifying a single read-only platform
> file**. The dev `anthropos` stack (and other demos) are never touched.

## Why this exists
Sales / CS / PM want to boot an isolated demo in the morning, seed it to a use-case (M4), demo it, and kill
it at night — no Clerk friction, no shared-staging contention. M3 is the disposable-stack layer; M4 fills
it with data; M5 makes it repeatable.

## The collision problem (and the additive fix)
The platform compose (`anthropos-dev/platform`) was built for **one** stack:
- **24 hard-coded host ports** (no env indirection) → a 2nd stack collides on every one.
- **one fixed `COMPOSE_PROJECT_NAME=anthropos`** → container/network name clashes.
- **one relative Postgres bind-mount** (`./data/postgresql`) → two stacks share one database.

`demo-stack` fixes all three **additively** — it never edits the platform compose:
1. **`-p demo-N`** isolates container/network/volume names.
2. A **generated compose override** (`docker-compose.demo.yml`) remaps every published host port to
   `host + N·OFFSET` and repoints Postgres's data bind to a per-demo dir. It uses Compose's **`!override`**
   tag so the inherited `ports`/`volumes` sequences are *replaced*, not appended (a plain merge would leave
   the base port bound → collide with the dev stack — the single most important detail here).
3. A **per-demo `.env.demo-N`** carries only the overrides; the platform `.env` (passed first) supplies
   creds/secrets — no secret duplication.

### Port-offset scheme
`demo-N → base_port + N·100` (default `OFFSET=100`, override with `DEMO_OFFSET`). So demo-1's Postgres is
`5532` (base `5432`), redis `6479` (base `6379`), graphql `5150` (base `5050`), etc. — none collide with
the dev stack's base ports. The registry assigns N and records the ports each demo owns.

## Lifecycle
```bash
DS=anthropos-demo/demo-stacks/demo-stack

# Full demo (RAM permitting — a full stack is ~10-12 GB):
"$DS" clone  1                 # per-demo clones, each repo at its LATEST RELEASE TAG (M3-D3)
"$DS" inject 1 --fapi-host localhost:5500   # wire the 4 Clerkenstein recipes (Clerk-free by default)
"$DS" up     1 --profile graphql            # bring up -p demo-1 on offset ports

# Minimal stack (infra only — proves isolation, fits a tight box):
"$DS" up     1 --services "postgresql redis"

"$DS" status                   # registry + per-demo `ps` + resolved per-repo refs
"$DS" down   1 --purge         # stop + remove demo-1 ONLY (+ its data) — dev stack untouched
```
The `/demo-up`, `/demo-down`, `/demo-status` skills wrap these.

### Clone at the release tag, not `main` (M3-D3)
`clone` checks out each platform service repo at its **most recent release tag** (semver, `v`-prefixed or
bare — `app → v1.282.0`, `platform → 0.1.0`), so a demo runs a *released*, reproducible version. Override
per call: `--ref main` (global) or `--ref app=v1.281.0` (per-repo); falls back to the default branch only
if a repo is untagged. The resolved ref per repo is recorded in the registry for reproduction.

### Clerkenstein wired in by default
`inject` applies the four recipes (see [`clerkenstein/knowledge/injection.md`](../services/clerkenstein.md)):
| Recipe | What `inject` does |
|---|---|
| clerk-frontend | mints `pk_test_<base64(host$)>` for the demo's fake FAPI → `NEXT_PUBLIC_/VITE_CLERK_PUBLISHABLE_KEY` (byte-identical to the gated `MintPublishableKey`) |
| clerk-backend | emits an `app` `extra_hosts: api.clerk.com:<fake-BAPI>` override; the trusted-cert step is documented |
| authn | emits the `go.mod replace` directive for the per-demo app clone (throwaway clone → no skip-worktree) |
| clerk-webhook | emits the svix-signed injector invocation feeding `POST /api/webhook/clerk` |

## Safety
Every `demo-stack` op is scoped `-p demo-N`. `down` **hard-refuses** any N that resolves to the dev
project name (read from the platform `.env`), so it can never tear down the dev stack. **Verified live:**
demo-1 up → status → down with the dev `anthropos` stack (12 containers, postgres healthy) untouched.

## Resource budget (important)
A full 12-service stack is **~10-12 GB RAM**; Docker Desktop's VM is often capped (~8 GB) and the dev
stack already fills most of it. On a 16 GB host you can run **one minimal demo alongside the dev stack**,
not two full stacks. `/demo-up` should resource-check before a full bring-up. **Running several full demo
stacks concurrently needs a bigger Docker VM / host** (M3-D5).

### What's proven vs. resource-gated (this hardware)
- **Proven live (16 GB box):** the override/isolation engine, the clone-at-release-tag resolution + real
  clones, the publishable-key mint, and the full up→status→down lifecycle of a minimal demo-1
  (postgres+redis) on offset ports with its own data, **co-resident with the dev stack, untouched**.
- **Resource-gated (→ bigger Docker VM):** a full 12-service single stack, two+ concurrent full stacks, and
  the end-to-end Clerkenstein browser-login (rebuild-with-replace + trusted cert + frontend rebuilt with the
  minted key). The wiring for all of these is built + documented; only the full live verification awaits the
  hardware.

## See also
- [`corpus/services/clerkenstein.md`](../services/clerkenstein.md) — the mock the demos are wired with.
- [`corpus/ops/platform_repo.md`](platform_repo.md) — the platform compose/Makefile this overlays.
- [`corpus/ops/run_guide.md`](run_guide.md) · [`corpus/ops/quick_ops.md`](quick_ops.md) — the dev-stack lifecycle.
- M4 (declarative seeding) + M5 (recipes) build on this — see `knowledge/plan/roadmap.md` § In Development — v1.1.
