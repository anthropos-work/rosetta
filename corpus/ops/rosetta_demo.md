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
> The generic `stack-*` skill set that surfaces this (renamed `/demo-*` → `/stack-*`) lands in M14.
