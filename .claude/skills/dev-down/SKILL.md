---
name: dev-down
description: Tear down an additional dev stack (dev-N, N ≥ 1) cleanly — stop + remove its containers/network/data — without touching the main dev stack or any demo. Use when asked to stop / kill / reclaim an extra dev environment.
argument-hint: [N] [--purge]
---

# Dev Down — tear down an additional dev stack

Stops and removes `dev-N` (N ≥ 1) **only** — its containers, network, and (with `--purge`) its data dir —
leaving the **main `anthropos` dev stack** (N=0) and any demo stacks untouched, and **releasing its slot**
in the unified registry so the N is free again (M12). It is the dev-side peer of `/demo-down`.

> The teardown tooling lives in `rosetta-extensions` (the executable stack tooling), consumed from the dev
> stack's per-stack copy `stack-dev/rosetta-extensions @ <tag>` — its pinned consumption clone. The
> `dev-stack` section name inside `rosetta-extensions` is unchanged.

## Mission
1. **Confirm N** — which dev stack to reclaim (`/stack-list`, or
   `stack-dev/rosetta-extensions/dev-stack/dev-stack status`, lists live `dev-N`). N must be ≥ 1 — the main
   dev stack (N=0) is **not** a `/dev-down` target.
2. **Tear down**:
   ```bash
   DEV=stack-dev/rosetta-extensions/dev-stack/dev-stack
   "$DEV" down N            # stop + remove dev-N's containers/network (releases the registry slot)
   "$DEV" down N --purge    # also remove dev-N's data dir (full reclaim)
   ```
3. **Verify** — `dev-N` is gone, its N is free again in the registry (`/stack-list`), and the **main dev
   stack is still intact** (`docker ps | grep '^anthropos-'` still shows its containers).

## Safety
`down` is hard-scoped `-p dev-N` and **refuses any N that resolves to the main dev project name** (whatever
`platform/.env`'s `COMPOSE_PROJECT_NAME` is) — it can never tear down the main dev stack. Teardown is the
only path that **frees** a registry slot (the race-guard half of the M12 allocator).

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-up` | Build / start / set-dress a dev stack |
| `/demo-up` · `/demo-down` | The demo-stack lifecycle (the peer of dev-up/dev-down) |
| `/stack-list` | List live dev + demo stacks and their freed/used slots |
