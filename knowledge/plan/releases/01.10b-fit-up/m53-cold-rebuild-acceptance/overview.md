---
milestone: M53
slug: cold-rebuild-acceptance
version: v1.10b "fit-up"
milestone_shape: section
status: planned
created: 2026-06-29
last_updated: 2026-06-29
complexity: medium
delivers: a release acceptance record + the v1.10.1 rext release tag + .agentspace/rext.tag bump
issues: the whole release, proven from cold (the user's "make sure this time it works")
---

# M53 — Cold-rebuild acceptance

## Goal
Prove the **whole release works from cold**. The fixes were iterated against the **single** live demo; this
milestone **destroys** it and **rebuilds from scratch** on a `stack-demo`-only box, then verifies end-to-end — the
single source of acceptance truth (the user's "first fix on live, then destroy + rebuild as the final step").

## Why section
A fixed acceptance checklist — destroy, cold rebuild, assert. No emergent path; any failure routes **back to the
owning milestone's fix**, it does not become new scope here.

## Repo split
- **No new feature code** (acceptance only). Produces the **`v1.10.1`** rext release tag + the `.agentspace/rext.tag`
  bump + a final corpus acceptance note.
- Drives the live `stack-demo` end-to-end via the skills (`/demo-down` → `/demo-up`).

## Scope
- **In:**
  - **Destroy** the live demo — `/demo-down` + image purge (exercises the M49 #6 `demo-down` image cleanup).
  - **Cold rebuild** on a `stack-demo`-only box — a single `/demo-up`, **no manual steps**, at the `.agentspace/rext.tag`
    pin (the consumed `fit-up-m47..m52` tags rolled to the `v1.10.1` release tag).
  - **Assert the acceptance bar:**
    - all backends healthy (the M47 re-synced clones + M49 secrets — no silent `app Exited`);
    - the cold-start MCP-DSN auto-capture filled the snapshot with **no prompt** (M47);
    - set-dress (recaptured snapshot) + seed (**all 3 orgs incl. AI-readiness**) + verify + cockpit all complete
      (no demo-up #7 abort);
    - **both-vantage M42 semantic coverage green** (employee + manager) on the existing orgs (M50);
    - the **AI-readiness dashboard criteria hold** on the 3rd org (M51: enabled, ~80%/3-step, 1 started + 1
      completed);
    - the cockpit **[Download manifest]** returns the **complete inlined** `seed-generation-manifest.yaml` (M52).
  - **Tag the rext** at **`v1.10.1`** (the release tag) + bump `.agentspace/rext.tag` to it.
- **Out:** new fixes — a failed assertion routes back to its owning milestone (M47–M52), not into M53.

## Depends on
**M52** (and transitively the whole chain). **Parallel with:** none (the final, single-demo acceptance).

## Open questions (resolve during build)
- None — this is the verification gate; the bar is the union of the prior milestones' exit conditions.

## KB dependencies (read as contract)
- `corpus/ops/verification.md` (the auto-verify net), `corpus/ops/demo/coverage-protocol.md` (the coverage gate),
  `corpus/ops/rosetta_demo.md` (the lifecycle), `corpus/ops/idempotency.md` (re-run safety).

## Delivers
- **→ rosetta-extensions:** the **`v1.10.1`** release tag (rolls up `fit-up-m47..m52`).
- **→ rosetta:** the `.agentspace/rext.tag` bump + a release acceptance record (feeds `/developer-kit:close-release`).

## Risk
**(blocks-release)** a cold-rebuild surfacing a regression late. *Mitigate:* each prior milestone verified
incrementally on the live demo; M53 is the from-cold confirmation, not the first integration. Failures route back,
never expand M53.
