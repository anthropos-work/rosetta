---
name: demo-down
description: Tear down a demo stack (demo-N) cleanly — stop + remove its containers/network/data — without touching the dev stack or other demos. Use when asked to stop / kill / reclaim a demo environment.
argument-hint: [N] [--purge]
---

# Demo Down — tear down a demo stack

Stops and removes `demo-N` **only** — its containers, network, and (with `--purge`) its data dir — leaving
the dev `anthropos` stack and any other demos untouched. Manual teardown is the only reclaim path (M3-D2).

## Mission
1. **Confirm N** — which demo to reclaim (`anthropos-demo/demo-stacks/demo-stack status` lists live demos).
2. **Tear down**:
   ```bash
   DS=anthropos-demo/demo-stacks/demo-stack
   "$DS" down N            # stop + remove demo-N's containers/network
   "$DS" down N --purge    # also remove demo-N's data dir (full reclaim)
   ```
3. **Verify** — demo-N is gone and the **dev stack is still intact** (`"$DS" status`;
   `docker ps | grep anthropos-` still shows the dev containers).

## Safety
`down` is hard-scoped `-p demo-N` and **refuses any N that resolves to the dev project name** — it can never
tear down the dev stack. Verified live: demo-1 up→status→down with the dev stack (12 containers) untouched.
