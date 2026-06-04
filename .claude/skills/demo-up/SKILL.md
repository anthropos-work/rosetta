---
name: demo-up
description: Bring up a disposable, isolated demo stack (demo-N) alongside the dev stack — Clerkenstein-wired, on offset ports, killable cleanly. Use when asked to spin up / start a demo environment.
argument-hint: [N] [--profile P] [--services "a b"] [--full]
---

# Demo Up — spin up an isolated demo stack

Brings up `demo-N` as a disposable Anthropos stack **alongside** the dev `anthropos` stack (or other
demos), on offset ports, with its own data, **Clerkenstein-wired by default** — **without touching any
read-only platform repo**. Source of truth: [`corpus/ops/rosetta_demo.md`](../../../corpus/ops/rosetta_demo.md).

## Mission
1. **Read the guide** — `corpus/ops/rosetta_demo.md` (lifecycle, the port-offset scheme, the resource budget).
2. **Resource check** — a full stack is ~10–12 GB. Confirm headroom (`docker info` MemTotal vs running
   stacks). For a tight box, bring up a minimal stack (infra only) or a reduced profile. **Never** exceed
   the box — multiple full stacks need a bigger Docker VM (M3-D5).
3. **Bring it up** via the tooling (the tooling lives in the gitignored `anthropos-demo/rosetta-demo/`):
   ```bash
   DEMO=anthropos-demo/rosetta-demo
   # FULL Clerk-free demo (every Clerk seam injected — real Clerk never used):
   "$DEMO/up-injected.sh" N        # clones + injects the 5 Clerk services (disarmed colony),
                                   # reuses dev images for the rest, runs the fake FAPI/BAPI,
                                   # brings up the full graphql stack -p demo-N. (~15-25 min first build.)
   # minimal stack (infra only — proves isolation, fits a tight box, fast):
   "$DEMO/rosetta-demo" up N --services "postgresql redis"
   ```
4. **Verify** — `"$DS" status`; confirm demo-N is on offset ports and the **dev stack is untouched**.

## Safety
Every op is `-p demo-N`-scoped; the tooling hard-refuses the dev project. The dev stack is never touched.
Resolved per-repo refs are recorded in the registry for reproduction.
