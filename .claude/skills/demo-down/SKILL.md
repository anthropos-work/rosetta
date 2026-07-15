---
name: demo-down
description: Tear down a demo stack (demo-N) cleanly — stop + remove its containers/network/data — without touching the dev stack or other demos. Use when asked to stop / kill / reclaim a demo environment.
argument-hint: [N] [--purge]
---

# Demo Down — tear down a demo stack

Stops and removes `demo-N` **only** — its containers, network, and (with `--purge`) its data dir — leaving
the dev `anthropos` stack and any other demos untouched. Manual teardown is the only reclaim path (M3-D2).

> The teardown tooling lives in `rosetta-extensions` (the executable stack tooling), consumed from the demo
> stack's per-stack copy `stack-demo/rosetta-extensions @ <tag>` — its pinned consumption clone. (Path renamed:
> `anthropos-demo/` → `stack-demo/`.) The `demo-stack` section name inside `rosetta-extensions` is unchanged.

## Mission
1. **Confirm N** — which demo to reclaim (`stack-demo/rosetta-extensions/demo-stack/rosetta-demo status` lists live demos).
2. **Tear down**:
   ```bash
   DS=stack-demo/rosetta-extensions/demo-stack/rosetta-demo
   "$DS" down N            # stop + remove demo-N's containers/network
   "$DS" down N --purge    # also remove demo-N's data dir (full reclaim)
   ```
3. **Verify** — demo-N is gone and the **dev stack is still intact** (`"$DS" status`;
   `docker ps | grep anthropos-` still shows the dev containers).

## Safety
`down` is hard-scoped `-p demo-N` and **refuses any N that resolves to the dev project name** — it can never
tear down the dev stack. Verified live: demo-1 up→status→down with the dev stack (12 containers) untouched.
Teardown **frees the demo's slot** in the unified dev+demo registry (M12), so its N is re-allocatable.

### The host-native listeners (M217 — check these, `compose down` cannot)

A demo owns **two listeners that are not containers**: the **presenter cockpit** (`7700+N·10000`) and
**ant-academy** (`3077+N·10000`). `docker compose down` cannot reach either. Since M217, `down` **reaps them by
PORT** (identity-checked — a foreign process on the port is reported, never killed) and says so.

> **Before M217 this leaked.** The reap was by PID from a pidfile that `launch_detached` writes *before* the bind
> succeeds and that a re-up *overwrites* — so a leaked cockpit became unreapable while teardown printed *"stopped
> the presenter cockpit"* regardless. A real orphan was found on `billion` still serving an **unauthenticated
> "become any hero" panel on `0.0.0.0:17700`**, pointing at a deleted database, after a `/demo-down` reported
> success.

**When verifying a teardown, check the cockpit port too** — not just the containers:

```bash
ss -tlnp | grep -E ":(7700|3077)"     # + N·10000 for demo-N   (macOS: lsof -nP -iTCP -sTCP:LISTEN)
```

## Related skills

| Skill | Use when |
|-------|----------|
| `/demo-up` | Bring up a demo stack |
| `/dev-up` · `/dev-down` | The **dev** lifecycle — the peer of demo-up/demo-down (`/dev-down` tears down a `dev-N`) |
| `/stack-list` | List live dev + demo stacks and their freed/used slots |
