---
milestone_shape: iterative
milestone: M258
title: "proven-live build (the closer)"
status: planned
release: v2.8 "fast build"
exit_gate: "One cold command on billion brings the stack up AND drives the full Playthrough batch to completion with ZERO standing red, at total p50 <= 480 s across 3 consecutive cold reset-to-seed cycles, reproducible, 0 platform-repo edits. Batch-gate semantics (D-v28-3): the suite always runs to completion — never halts at first red, never retries to mask a flake — and emits ONE consolidated red set at batch end; a non-empty set escalates to the user for renegotiation. The stack is left UP regardless; the bring-up exits non-zero and says so loudly."
iteration_protocol_ref: corpus/ops/verification.md
re_scope_trigger: "If the composed p50 exceeds 600 s after 3 tiks, split the suite into a fast smoke lane gating the bring-up + a full lane run after, and renegotiate the gate with the user."
depends_on: [M256, M257]
parallel_with: []
complexity: large
created: 2026-07-27
last_updated: 2026-07-27
---

# M258 — proven-live build  (`iterative`, the closer)

**Status:** `planned` · **Shape:** `iterative` (the closer) · **Complexity:** large · **Release:** v2.8 "fast build"
**Depends on:** M256, M257

## Goal

A demo stack comes up **and proves itself**. One cold command, and what you get is not "UP" but
**"UP, and every journey verified"** — fast enough that this is the normal way to bring a stack up, not a
ceremony reserved for a release gate.

## What this composes

- **M257's** restructured bring-up (parallel UI tier, multi-stage images, overlapped `compose up`).
- **M256's** restructured suite (read-only parallel lane, per-seat session reuse, negative controls, the
  landed onboarding + org-admin coverage).
- The existing **`autoverify`** net (`corpus/ops/verification.md`) — which proves the stack is *reachable and
  healthy*. This milestone adds the layer above it: the stack is *functionally correct*.

## Batch-gate behaviour (D-v28-3) — the design's core

The suite **always runs to completion**. It never halts at the first red. It never retries to mask a flake
(`retries: 0` stays). At batch end it emits **one consolidated red set**.

- **Empty red set** → the gate.
- **Non-empty red set** → **escalates to the user for renegotiation**, once, at batch end. Each item is either
  fixed or given an **explicit written disposition**. Red **never accumulates silently across runs**.
- **The stack is left UP regardless** — the `autoverify` precedent: a test bug must never cost a good demo.
- **The bring-up exits non-zero and says so loudly** when the red set is non-empty. Loud, not fatal.

This is the answer to *"just make sure there aren't accumulating red playthroughs — all has to work, or it
needs to be renegotiated with me at each full batch run (don't stop at every step)."*

## Shape (why iterative)

The **composition** is the unknown. A bring-up just restructured for speed meets a suite just restructured for
parallelism, on a box whose headroom is now budgeted. The interactions — a parallel build changing image
readiness ordering, a parallel test lane contending with set-dress, a headroom cap starving the browser
workers — surface only live. A fixed checklist would be speculative.

## Iteration protocol

The **prove-on-billion** lineage (M221 → M236 → M244 → M254):

- **fresh agent per run** — context does not survive an 11-minute foreground op cleanly;
- sub-agents **foreground-poll** long operations — **never background-and-yield** (the documented stall trap);
- the coordinator watchdog is **never stood down** (the 7.5 h lesson);
- **pre-flight rung zero** (`corpus/ops/verification.md`): *tagging is not publishing* — verify the rext tag is
  on **origin** (`git ls-remote --tags origin`) before any live prove. M236 lost its entire first iteration to
  a tag that existed only in the local authoring copy.
- **state the environment with every number** (`latency-budget.md`).

## Open questions

- Does the batch run **on the demo host** or **from a tailnet peer**? A `--public-host` demo can only be
  *browsed* from a peer (docker-proxy binds `0.0.0.0`, so a connection from the host to its own tailscale IP
  bypasses `tailscale serve` and dies with `ERR_SSL_PROTOCOL_ERROR`). `run-playthroughs.sh --reset-only`
  already splits the DB half from the browser half — the composed command must respect that split.
- Should the baked-in batch be the **full** suite or a **gate subset**, if M256's suite lands well above 120 s?
  (The re-scope trigger covers this; the preference is full.)

## Hard constraints

Zero platform-repo edits · all tooling in `rosetta-extensions`, tagged, **pushed to origin** · reset-to-seed
reproducible (the real `--reset`; additive re-seed is FORBIDDEN) · N=0 dev-stack guard honored.

## KB dependencies

`corpus/ops/verification.md` (the iteration protocol) · `corpus/ops/demo/playthroughs.md` ·
`corpus/ops/demo/build-budget.md` (M255) · `corpus/ops/demo/tailscale-serve.md` · `corpus/ops/rosetta_demo.md`
· `corpus/ops/idempotency.md` · `corpus/ops/demo/latency-budget.md`

**Delivers → `corpus/ops/verification.md`** (the bring-up now ends in a functional batch gate, not only
`autoverify`)
**Delivers → `corpus/ops/demo/playthroughs.md`** (the baked-in lifecycle)
