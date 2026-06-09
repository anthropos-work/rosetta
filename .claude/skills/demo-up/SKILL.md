---
name: demo-up
description: Bring up a disposable, isolated demo stack (demo-N) alongside the dev stack — Clerkenstein-wired, on offset ports, with the full UI tier (next-web + studio-desk + ant-academy), killable cleanly. Use when asked to spin up / start a demo environment.
argument-hint: [N] [--profile P] [--services "a b"] [--full] [--no-ui]
---

# Demo Up — spin up an isolated demo stack

Brings up `demo-N` as a disposable Anthropos stack **alongside** the dev `anthropos` stack (or other
demos), on offset ports, with its own data, the **full UI tier by default** (next-web + studio-desk + ant-academy),
**Clerkenstein-wired** — **without touching any read-only platform repo**. Source of truth:
[`corpus/ops/rosetta_demo.md`](../../../corpus/ops/rosetta_demo.md) (lifecycle) +
[`corpus/ops/demo/frontend-tier.md`](../../../corpus/ops/demo/frontend-tier.md) (the UI tier).

## Mission
1. **Read the guides** — `corpus/ops/rosetta_demo.md` (lifecycle, port-offset, resource budget) +
   `corpus/ops/demo/frontend-tier.md` (the UI tier: per-demo frontend builds, the pk/URL baking, the 12 GB
   VM prereq, ant-academy, `--no-ui`).
2. **Resource check** — a full stack is ~10–12 GB at runtime; the **frontend build** spikes to ~3.7 GB. Set
   the **Docker VM to 12 GB / swap 3 GB** (Settings → Resources) so the per-demo next-web build doesn't
   swap-thrash. `up-injected.sh` runs a **non-fatal** 12 GB pre-flight assert (warns, never blocks;
   `DEMO_VM_MIN_GIB=N` overrides). For a tight box or an API-only demo, use `--no-ui` (`DEMO_NO_UI=1`) or a
   reduced `--services` set. **Never** exceed the box.
3. **Bring it up** via the tooling. The demo stack consumes `demo-stack` tooling from its **OWN**
   `stack-demo/rosetta-extensions` clone pinned at a tag — never edited ad-hoc inside `stack-demo`.
   New or changed tooling is authored + tested in the `.agentspace/rosetta-extensions/` authoring copy
   and tagged first, then consumed per-stack at that pinned tag.
   ```bash
   DEMO=stack-demo/rosetta-extensions/demo-stack
   # FULL Clerk-free demo WITH the UI tier (default): the 5 injected Clerk services + fake FAPI/BAPI +
   # per-demo next-web + studio-desk (offset ports, minted-pk-baked, cached) + ant-academy native.
   "$DEMO/up-injected.sh" N        # ~15-25 min first build; +~3 min/frontend on a NEW demo-N, cached after.
   # backend-only (skip the UI tier — fast, RAM-light, API/QA):
   DEMO_NO_UI=1 "$DEMO/up-injected.sh" N
   # minimal stack (infra only — proves isolation, fits a tight box, fast):
   "$DEMO/rosetta-demo" up N --services "postgresql redis"
   ```
4. **Verify** — the bring-up auto-runs a scoped, non-fatal verify (covers the UI tier when present). Then
   `"$DEMO/rosetta-demo" status`; confirm demo-N is on offset ports (next-web `:3000+`, studio-desk `:9100+`,
   ant-academy `:3077+`) and the **dev stack is untouched**.

## Safety
Every op is `-p demo-N`-scoped; the tooling hard-refuses the dev project. The dev stack is never touched.
Resolved per-repo refs are recorded in the registry for reproduction. N is allocated from the **unified
dev+demo registry** (M12), so a `demo-N` can never collide with a `dev-N` on ports.

## After bring-up

Set-dress + seed the demo with the generic stack-ops (they accept `demo-N` or `dev-N`):
`/stack-snapshot N` (replay the real public catalog + content) → `/stack-seed N` (a believable data world)
→ log in. List live stacks with `/stack-list`.

## Related skills

| Skill | Use when |
|-------|----------|
| `/demo-down` | Tear down a demo stack |
| `/dev-up` · `/dev-down` | The **dev** lifecycle — the peer of demo-up/demo-down, same registry + offset-port model |
| `/stack-snapshot` · `/stack-seed` | Set-dress + seed the demo (generic — any `dev-N \| demo-N`) |
| `/stack-list` | List live dev + demo stacks |
