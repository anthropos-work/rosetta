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
  tooling there); the demo stack consumes it at a pinned tag as `stack-demo/rosetta-extensions @ <tag>`.
- **The skills (here in rosetta):** [`/demo-up`](../../.claude/skills/demo-up/SKILL.md), `/demo-down`,
  and the generic `/stack-list` drive that tooling (the dev peer is `/dev-up` / `/dev-down`).
- **The secrets:** [`/stack-secrets`](../../.claude/skills/stack-secrets/SKILL.md) provisions the stack's
  per-repo `.env` from one secret source — **values-blind** — and verifies coverage; `/demo-up` runs it as an
  auto-provision step (v1.6 M30) so a demo is self-sourced from the curated secret dir and the prod-write path
  is never re-armed (`DIRECTUS_TOKEN` blank on the non-prod target). Mechanism: [`secrets-spec.md`](secrets-spec.md).
- **The mock it injects:** `rosetta-extensions/clerkenstein/` — see [clerkenstein.md](../services/clerkenstein.md).

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

## Stack workspace layout + the `anthropos-dev` → `stack-dev` back-compat fallback (v1.3b "dress rehearsal", M16)

The tooling resolves the **local dev workspace** to find the platform repos a stack builds from (`<dev>/platform/.env`
for `GH_PAT`, `<dev>/<svc>` for the atlas migration dirs, `<dev>/sentinel/init_policy.sql`). v1.3 "stack party"
renamed that workspace `anthropos-dev/` → **`stack-dev/`** (one of the `stack-*/` family — see CLAUDE.md
*"Working in stack workspaces"*). The rename converged the *whole* family on the `stack-<role>/` convention
(`stack-dev/`, `stack-demo/`, `stack-dev-2/`, …); each holds its cloned platform service repos plus its own
pinned-tag clone of `rosetta-extensions`.

The dev/demo CLIs resolve the workspace with a **single intentional back-compat fallback** — they prefer
`stack-dev/`, and fall back to the legacy `anthropos-dev/` only if `stack-dev/` is absent:

```bash
DEV="$REPO_ROOT/stack-dev"; [ -d "$DEV" ] || DEV="$REPO_ROOT/anthropos-dev"   # prefer stack-dev; legacy fallback
```

This is the **one** place `anthropos-dev` survives — a one-line auto-detect (in `up-injected.sh`, `migrate-demo.sh`,
`rosetta-demo`, `dev-stack`, and the `clone_repos.py` `--dev-root` help) that costs nothing and protects an older
on-disk layout. Everywhere else `stack-dev/` is the documented default. The fallback was the M16 field fix
(shipped in `rosetta-extensions @ dress-rehearsal-m16`): a fresh box that already used `stack-dev/` would otherwise
die at bring-up resolving a non-existent `anthropos-dev/`. **Don't reintroduce bare `anthropos-dev/` references** in
prose or scripts — keep it confined to the fallback line.
