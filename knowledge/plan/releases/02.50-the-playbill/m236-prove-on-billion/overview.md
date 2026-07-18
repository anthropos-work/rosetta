---
milestone: M236
slug: prove-on-billion
version: v2.5 "the playbill"
milestone_shape: iterative
status: planned
created: 2026-07-19
last_updated: 2026-07-19
depends_on: M235
exit_gate: "Both tabs work live on billion — Content-stories sessions render real content for player + manager vantages, the academy grid renders real cards (Thread A) — reproducibly on a cold reset-to-seed, p95 click->ACCESS < 5 s, 0 platform edits, demo reachable only over the tailnet."
iteration_protocol_ref: corpus/ops/verification.md + corpus/ops/demo/tailscale-serve.md
delivers: none
---

# M236 — prove on billion

**Status:** `planned`  ·  **Shape:** `iterative`  ·  **Complexity:** medium  ·  **Depends on:** M235

## Goal
Re-prove the whole feature live on the billion Tailscale VM (the house pattern that closed M215/M221/M226/M228) — both cockpit tabs usable end-to-end from a 2nd machine on a cold reset-to-seed, VPN-scoped.

## Exit gate
Both tabs work live on billion — Content-stories sessions render real content for player + manager vantages, the academy grid renders real cards (Thread A) — reproducibly on a cold reset-to-seed, p95 click->ACCESS < 5 s, 0 platform edits, demo reachable only over the tailnet.

**Iteration protocol:** `corpus/ops/verification.md + corpus/ops/demo/tailscale-serve.md`

## Scope
### In
- Bring up the demo on billion; drive both the Org-stories and Content-stories tabs remotely
- Prove content-stories sessions render live for player + manager; both academy tabs render (Thread A); reproduce on a cold reset-to-seed; capture p95 click->ACCESS vs the <5 s gate

### Out
- New feature work (built by M235)

## Open questions
- Fold Thread A's live-render prove into this milestone (one combined release -> yes)

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
