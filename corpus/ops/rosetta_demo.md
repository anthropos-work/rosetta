# Disposable Demo Stacks

**Status:** v1.1 "show floor" / M3 · **Last updated:** 2026-06-04 · **Tooling:** `anthropos-demo/rosetta-demo/` (gitignored, own git, no remote) · **Skills:** `/demo-up`, `/demo-down`, `/demo-status` · **Full injected stack:** LIVE-PROVEN (bring-up → migrate → schema → Clerk-free auth; `/api/health` 200)

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

`rosetta-demo` fixes all three **additively** — it never edits the platform compose:
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
DS=anthropos-demo/rosetta-demo/rosetta-demo

# Full Clerk-free demo — ONE call (measured ~0.9 GB; LIVE-PROVEN co-resident with the dev stack):
anthropos-demo/rosetta-demo/up-injected.sh 1   # clone@tag → inject 4 recipes → build → override → up
anthropos-demo/rosetta-demo/migrate-demo.sh 1  # schemas + atlas migrations → sentinel healthy
                                              # /api/health → 200 (authorized routes need the M4 seed)

# Or the individual rosetta-demo verbs (minimal/manual, e.g. an infra-only proof):
"$DS" clone  1                 # per-demo clones, each repo at its LATEST RELEASE TAG (M3-D3)
"$DS" inject 1 --fapi-host localhost:15400   # wire the 4 Clerkenstein recipes (Clerk-free by default)
"$DS" up     1 --profile graphql            # bring up -p demo-1 on offset ports

# Minimal stack (infra only — proves isolation, fits a tight box):
"$DS" up     1 --services "postgresql redis"

"$DS" status                   # registry + per-demo `ps` + resolved per-repo refs
"$DS" down   1 --purge         # stop + remove demo-1 ONLY (+ its data) — dev stack untouched
```
The `/demo-up`, `/demo-down`, `/demo-status` skills wrap these.

## Full Clerk-free bring-up (M3-proven: `up-injected.sh` + `migrate-demo.sh`)
Two scripts orchestrate the complete flow — the whole point of M3 is that this needs **no real Clerk**.

### `up-injected.sh <N>` — bring up the full stack
One call does the lot:
1. **Clones** the 5 Clerk-consuming Go services (`app`, `skiller`, `cms`, `jobsimulation`, `skillpath`)
   from their dev clones at their latest release tag into `stacks/demo-N/clones/` (cms also copies its
   `studio/` submodule).
2. **Injects** the disarmed vendored colony into each — `inject/apply-authn.sh` clones `colony` at the
   version the service pins, swaps `authn/provider/clerk` for the disarmed twin, and adds a local
   `replace … => ./vendor-colony` (the **authn** recipe, zero app-code change).
3. **Builds** those 5 services as `demo-N-<svc>:injected`, plus the fake FAPI (`fake-fapi`) and fake BAPI
   (`fake-bapi`) from Clerkenstein as tiny demo images. Non-Clerk services reuse the dev images — no rebuild.
4. **Generates** the injected compose override (`lib/gen_injected_override.py`) — wires the fake servers,
   aliases the fake BAPI as `api.clerk.com`, applies the port offset — and the per-demo `.env`
   (`lib/inject.py`, which mints the publishable key + emits the webhook/extra-hosts snippets).
5. **Brings up** `-p demo-N` with the `graphql` profile (all ~13 services), then runs `migrate-demo.sh`.

### `migrate-demo.sh <N>` — initialize the schema
1. Creates the schemas (`extensions`, `sentinel`, `cms`, `jobsimulation`, `skiller`, `skillpath`) and the
   `vector` / `pgcrypto` / `pg_trgm` extensions in demo-N's Postgres (host port `5432 + N·10000`).
2. Runs `atlas migrate apply --env local` for the 5 migration services against demo-N's Postgres.
3. Restarts `demo-N-sentinel-1` + `demo-N-backend-1` so they pick up the migrated schema.

Result: **sentinel goes healthy and `/api/health` returns 200.** Authorized routes still 403 — they need the
M4 declarative seed.

### Tear down
```bash
"$DS" down N --purge      # hard-scoped to -p demo-N; the dev stack is never touched
```

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

### Injection recipe status (M3 — all four PROVEN)
| Recipe | Implemented by | Status |
|---|---|---|
| authn (disarmed `colony/authn` provider) | `inject/apply-authn.sh` | PROVEN — disarmed twin vendored into each service; the real `colony @ v0.34.3` API is satisfied (compiled, not reimplemented) |
| clerk-frontend (publishable-key mint) | `lib/inject.py` (`mint_pk`) | PROVEN — mints `pk_test_<RawStdBase64(host$)>`, byte-identical to Clerkenstein's authoritative Go `cmd/mintpk`; asserted by `test_mint_matches_clerkenstein_source` |
| clerk-backend (BAPI redirect) | `lib/gen_injected_override.py` + `lib/inject.py` | PROVEN — aliases the fake BAPI as `api.clerk.com` (compose network alias) + emits the `extra_hosts: !override` snippet for `app` |
| clerk-webhook (svix-signed injector) | `lib/inject.py` | PROVEN — emits the `NewInjector(endpoint, secret).Inject(payload)` invocation feeding `POST /api/webhook/clerk` |

`up-injected.sh` applies all four in one call — no manual multi-step. The disarmed stack runs entirely on the
**demo identity** seeded by Clerkenstein's `DefaultDemoUser()`: `user_clerkenstein` / `demo@anthropos.test`
(Eid `11111111-…`) as `admin` of `org_clerkenstein` (OrgEid `22222222-…`). The disarmed provider is
identity-agnostic (straight-through claim mapping — it extracts whatever the minted token carries).

## Safety
Every `rosetta-demo` op is scoped `-p demo-N`. `down` **hard-refuses** any N that resolves to the dev
project name (read from the platform `.env`), so it can never tear down the dev stack. **Verified live:**
demo-1 up → status → down with the dev `anthropos` stack (12 containers, postgres healthy) untouched.

## Resource budget (measured reality, M3-proven)
RAM was **never the blocker** the earlier estimate feared. **Measured:** the entire dev `anthropos` stack
(~13 services) sits at **~0.9 GB RAM**, not the old "10-12 GB" guess. A full Clerk-free demo stack is the
same compose profile and service count — stack isolation (the `-p demo-N` project, offset ports, per-demo
data dir) is overhead-free — so a full demo runs **comfortably alongside the dev stack on a 16 GB box**.
Concurrent full stacks (demo-1 + demo-2) fit too; the only practical ceiling is Docker Desktop's configured
VM memory.

### What's proven (2026-06-04)
- **PROVEN LIVE:** the full end-to-end bring-up (`up-injected.sh` → `migrate-demo.sh`) of a full ~13-service
  demo on offset ports, **co-resident with the dev stack, untouched**. All four Clerkenstein injection
  recipes (authn / clerk-frontend / clerk-backend / clerk-webhook), the override/isolation engine, the
  clone-at-release-tag resolution + real clones, the publishable-key mint, the registry, and the full
  up→status→down lifecycle. Sentinel goes healthy; **`/api/health` returns 200**.
- **Pending (the M4 seed):** authorized routes still **403** until the platform is seeded with the demo
  identity's grants. The wiring (disarmed authn accepts the token; authz then rejects pending grants) is in
  place — only the declarative seed (M4) is missing. The end-to-end browser-login (frontend rebuilt with the
  minted key + trusted cert) is built + documented; its full live walk-through follows the same path.

## Testing & verification
The rosetta-demo tooling carries **78 unit tests** (pytest):
- `tests/test_tooling.py` (707 loc, 55 tests) — the override generator, clone resolver, publishable-key
  minter, and registry.
- `tests/test_inject_scripts.py` (439 loc, 23 tests) — the four injection recipes.

```bash
cd anthropos-demo/rosetta-demo && python3 -m pytest tests/ -v
```
The load-bearing one is `test_mint_matches_clerkenstein_source`: it shells out to Clerkenstein's
authoritative Go `cmd/mintpk` and **fails** if the Python `mint_pk` ever diverges from it — keeping the demo
publishable key byte-identical to the alignment-gated contract.

## See also
- [`corpus/services/clerkenstein.md`](../services/clerkenstein.md) — the mock the demos are wired with
  (100% / 100% on all four alignment surfaces, including the deployment/injection dimension).
- `anthropos-demo/clerkenstein/knowledge/kb-index.md` — Clerkenstein's own KB (architecture, alignment,
  injection recipes, coverage); in the gitignored scratchpad alongside the source.
- `anthropos-demo/rosetta-demo/inject/DEPLOYMENT-PROOF.md` — the M3 injection bring-up proof.
- [`corpus/ops/platform_repo.md`](platform_repo.md) — the platform compose/Makefile this overlays.
- [`corpus/ops/run_guide.md`](run_guide.md) · [`corpus/ops/quick_ops.md`](quick_ops.md) — the dev-stack lifecycle.
- M4 (declarative seeding) + M5 (recipes) build on this — see `knowledge/plan/roadmap.md` § In Development — v1.1.
