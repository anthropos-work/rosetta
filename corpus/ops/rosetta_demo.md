# Demo stacks — moved to `rosetta-extensions`

> **This is a pointer.** The disposable-demo-stack tooling + its full lifecycle guide now live in the
> **`rosetta-extensions`** monorepo (private: `anthropos-work/rosetta-extensions`), section **`demo-stack/`**.
> rosetta documents *how the platform works*; the tools that spin up demo copies of it live in the extensions repo.

- **What it is:** bring up `demo-N` as isolated, Clerkenstein-wired full stacks alongside the dev stack, on offset
  ports, killable cleanly — **zero read-only-platform change**.
- **The guide:** `rosetta-extensions/demo-stack/GUIDE.md` (lifecycle, port-offset scheme, the 4 injection recipes,
  `migrate-demo.sh`, safety).
- **The tooling (gitignored locally):** `stack-demo/rosetta-extensions/demo-stack/` — the `rosetta-demo` CLI,
  `up-injected.sh`, `migrate-demo.sh`, `inject/`.
- **The clone-role/tag model:** the authoring copy lives at `.agentspace/rosetta-extensions/` (build/test/tag the
  tooling there); the demo stack consumes it at a pinned tag as `stack-demo/rosetta-extensions @ <tag>`
  (current post-v1.9 demo-stack/set-dress tag: **`storytelling-postfix-1`**).
- **The skills (here in rosetta):** [`/demo-up`](../../.claude/skills/demo-up/SKILL.md), `/demo-down`,
  and the generic `/stack-list` drive that tooling (the dev peer is `/dev-up` / `/dev-down`).
- **The secrets:** [`/stack-secrets`](../../.claude/skills/stack-secrets/SKILL.md) provisions the stack's
  per-repo `.env` from one secret source — **values-blind** — and verifies coverage; `/demo-up` runs it as an
  auto-provision step (v1.6 M30) so a demo is self-sourced from the curated secret dir and the prod-write path
  is never re-armed (`DIRECTUS_TOKEN` blank on the non-prod target). Mechanism: [`secrets-spec.md`](secrets-spec.md).
- **The mock it injects:** `rosetta-extensions/clerkenstein/` — see [clerkenstein.md](../services/clerkenstein.md).
- **Multi-identity seat-switch (v1.9 "storytelling" M37):** Clerkenstein's fake FAPI is now **multi-identity** —
  a demo can **switch the active browser identity** among the seeded heroes/orgs (the M35 stories roster), the
  seat-switch the presenter cockpit's "login as" drives. The demo tooling exports a **roster** (`FAKE_FAPI_ROSTER`,
  the seeder-derived clerk claims per hero) the `fake-fapi` loads; the browser selects a seat via
  `?__clerk_identity=<key>` on the handshake (the [Login as] deep-link) or the `/v1/demo/{identities,select}`
  control plane. Server-authoritative, so every surface resolves the same hero. Measured by the `clerk-multi-1`
  Alignment DNA (9 genes, 100%/100%) — see [clerkenstein.md](../services/clerkenstein.md) § Multi-identity.
- **Stories & cockpit are the DEFAULT (post-v1.9 demo-hardening):** a bare `/demo-up N` now seeds the
  multi-org **Stories & Heroes** world (2 orgs × a thriving/struggling/manager hero trio) **and** serves the
  presenter cockpit **by default** — the M38-D4 opt-in flipped to opt-**out**. `DEMO_NO_STORIES=1` (or the
  explicit `DEMO_STORIES=0`) restores the legacy structural **small-200** + single-identity fake-fapi +
  no-cockpit demo (mirroring the `DEMO_NO_*` family). So: stories = default; small-200 = the
  `DEMO_NO_STORIES` fallback. Full stories model: [`demo/stories-spec.md`](demo/stories-spec.md).

## A demo builds from its OWN clone set — self-contained (v1.8 "understudy", M26)

`stack-demo` is a **true peer of `stack-dev`**: it has its **own** platform clone set, and a demo builds
**entirely** from it. A box with **only** `stack-demo/` (no `stack-dev/`) can bring a demo up end-to-end.

**The from-scratch bring-up (what `/demo-up N` does, in order):**
1. **`ensure-clones.sh`** (the first action, before any build) bootstraps `stack-demo`'s peer clone set:
   - bootstrap-clone `stack-demo/platform` from `git@github.com:anthropos-work/platform.git` over SSH if
     absent (fail-loud on a clone failure — **never** falls back to `stack-dev` for the build SOURCE);
   - **seed** the shared secrets: copy `stack-dev/platform/.env` → `stack-demo/platform/.env`
     **copy-if-present** (same Clerk app + same `GH_PAT`, shared by nature; never committed). A box with no
     `stack-dev` **skips this non-fatally** — `/stack-secrets` (M30) provisions the real `.env` from
     `.agentspace/secrets`. This `.env` copy is the **sole** sanctioned `stack-dev` read;
   - `make -C stack-demo/platform init` clones every `repos.yml` repo as a sibling into `stack-demo/`
     (skip-if-present — the platform's own idempotent clone loop), plus `make init-studio` for `cms`;
   - record per-repo `{ref,sha}` provenance into `stack-demo/clones.lock.json`.
2. **build everything from `stack-demo`**: the 5 injected Go services clone their per-demo COPY from
   `stack-demo/<svc>`; the two frontends build from `stack-demo/next-web-app` + `stack-demo/studio-desk`;
   the non-Clerk services (sentinel/storage/roadrunner/graphql) build from `stack-demo`'s clones via the
   compose `build.context` (the compose dir `PLAT` is `stack-demo/platform`, so the relative contexts
   resolve against `stack-demo`). **Dev-image reuse is OFF by default** — a demo never inherits `stack-dev`'s
   built images (which could carry dev WIP), even when dev is up; opt back in with `DEMO_REUSE_DEV_IMAGES=1`.
3. the disarmed-colony injection still mutates **only** the per-demo COPY at `stacks/demo-N/clones/<svc>` —
   the shared `stack-demo/<svc>` clone is the COPY's SOURCE and stays git-clean.

> **The manual `rosetta-demo up` verb** is the minimal/infra-only path — it does **not** call
> `ensure-clones.sh` (the auto `/demo-up` path does). It presupposes a populated `stack-demo` (run
> `up-injected.sh` / a prior `ensure-clones.sh` to bootstrap the peer clone set + seed the shared `.env`).

## Unified stack registry + first-available-N allocation (v1.3 "stack party", M12)

Every isolated stack — **dev** *or* **demo** — maps host port `P → P + N·10000`, so its `N` is what keeps
it off every other stack's ports. Before v1.3 the two kinds tracked `N` separately (demo had a demo-only
registry; dev had none), so `dev-1` and `demo-1` resolved to the **same** offset and collided on every
published port. M12 makes `N` a **single shared resource across both kinds**.

- **One unified registry** (in `rosetta-extensions/stack-core/`, shared by both the `rosetta-demo` and
  `dev-stack` CLIs). One record per live stack, keyed by its docker project `"<type>-<N>"`:
  `{type: dev|demo, n, ports, status, created}`. Pure runtime (gitignored), `flock`-guarded, atomic writes.
- **First-available-N allocation.** Bring-up takes an **explicit `N`** *or* **auto-allocates the lowest
  free `N` across dev+demo**. The allocator reconciles the registry against live `docker ps` (the project
  labels `-p dev-N` / `-p demo-N` are the truth for "this `N` is live") — so a manually-started stack is
  adopted and never double-allocated, and the registry self-heals. A reserved `N` is freed **only** by
  teardown (`down` → release), never by a lagging `docker ps` — which is the race guard that lets a
  just-reserved stack survive the gap before its containers appear.
- **The guarantee:** bringing up `dev, demo, dev, demo, demo` (in any interleaving, from either CLI)
  yields `dev-1, demo-2, dev-3, demo-4, demo-5` — no port collisions, ever. Teardown frees the slot, so the
  next bring-up reclaims the lowest hole.

> **Where it lives / the full model:** `rosetta-extensions/stack-core/README.md` (the registry schema +
> allocator contract), with the demo + dev CLIs documented in `demo-stack/GUIDE.md` and `dev-stack/README.md`.
> The generic `stack-*` skill set that surfaces this (renamed `/demo-*` → `/stack-*`) shipped in M14.

> **The registry's recorded ports back the verify cross-check (v1.3b/M18).** Each record's resolved host
> ports are the **source of truth** the auto-verify (`stack-verify`) reads to confirm it's targeting the
> right offset — never a bare re-computed formula. See [`verification.md`](verification.md) (the auto-verify
> safety net: a scoped, non-fatal `verify live` at every bring-up tail, offset/scope-aware).

> **The per-stack Directus rides the same lifecycle for free (v1.5 "prop room", M22).** A `--local-content`
> stack (demo default; dev opt-in) emits a `directus` **compose service** into its override — offset port
> `8055 + N·10000`, container `<project>-directus-1`, on the stack's `app-network`. Because it's a compose
> service (not a bespoke `docker run`), the existing plumbing covers it with **no new lifecycle code**:
> `demo-down`/`dev-down` (incl. `--purge`) tear it down with the rest of the stack, the unified registry
> records its offset port (`ports_from_override` picks it up), and the auto-verify probes it by the
> `<project>-directus-1` name. A prod-read stack (`DEMO_NO_LOCAL_CONTENT=1` / no `--local-content`) has no
> directus service, so teardown / registry / verify all correctly see nothing. Full lifecycle:
> [`directus-local.md`](directus-local.md) § "Container lifecycle (M22)".

## Auto-verify at the bring-up tail (v1.3b "dress rehearsal", M18)

Every bring-up now ends with an automatic, **non-fatal** verification pass on the stack's own **offset
ports**: the cheap-win `/api/health` + `sentinel.casbin_rules > 0` asserts (the silent-403 catcher), then
the full offset/project/scope-aware `verify live`. So "UP" means *verified-working*, not just
*containers-started*. Default-on (opt out: `DEMO_NO_VERIFY=1` / `DEV_NO_VERIFY=1`); a failing check warns
loudly + points at `/test-platform N` but never aborts a good bring-up. Full contract:
[`verification.md`](verification.md).

> **The per-stack-Directus boot now health-gates so auto-verify can't race it (post-v1.9 demo-hardening).**
> The bring-up's Directus boot step now **waits** for the stack's own offset `/server/health` to answer 200
> (bounded, non-fatal on timeout) before returning, so the verify pass no longer false-reports Directus
> "down" while it re-introspects across its restart. See [`verification.md`](verification.md).

## Stack workspace layout + the `anthropos-dev` → `stack-dev` back-compat fallback (v1.3b "dress rehearsal", M16)

The tooling resolves the **local dev workspace** to find the platform repos a stack builds from (`<dev>/platform/.env`
for `GH_PAT`, `<dev>/<svc>` for the atlas migration dirs, `<dev>/sentinel/init_policy.sql`). v1.3 "stack party"
renamed that workspace `anthropos-dev/` → **`stack-dev/`** (one of the `stack-*/` family — see CLAUDE.md
*"Working in stack workspaces"*). The rename converged the *whole* family on the `stack-<role>/` convention
(`stack-dev/`, `stack-demo/`, `stack-dev-2/`, …); each holds its cloned platform service repos plus its own
pinned-tag clone of `rosetta-extensions`.

The **dev** CLI resolves the dev workspace with a **single intentional back-compat fallback** — it prefers
`stack-dev/`, and falls back to the legacy `anthropos-dev/` only if `stack-dev/` is absent:

```bash
DEV="$REPO_ROOT/stack-dev"; [ -d "$DEV" ] || DEV="$REPO_ROOT/anthropos-dev"   # prefer stack-dev; legacy fallback
```

This back-compat fallback now lives **dev-side only** — `dev-stack` (the dev CLI) + the `clone_repos.py` `--dev-root`
help string. **v1.8 "understudy" (M26) removed it from the demo scripts** (`up-injected.sh`, `migrate-demo.sh`,
`rosetta-demo`, `ant-academy.sh`): a demo now resolves its **own** `stack-demo/` peer clone set (see *"Self-contained
demo stacks"* above) — there is no `stack-dev`/`anthropos-dev` to fall back to on the demo build path. It costs
nothing on the dev side and protects an older on-disk layout. Everywhere else `stack-dev/` is the documented default.
The fallback was the M16 field fix
(shipped in `rosetta-extensions @ dress-rehearsal-m16`): a fresh box that already used `stack-dev/` would otherwise
die at bring-up resolving a non-existent `anthropos-dev/`. **Don't reintroduce bare `anthropos-dev/` references** in
prose or scripts — keep it confined to the fallback line.
