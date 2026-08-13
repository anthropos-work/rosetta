---
iter: 25
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-25 — D-v28-5, both halves, proven on one rebuild

**Active strategy reference:** `TOK-01` move 4 (*"close the honesty items last, deliberately"*). D-v28-5
is the one **gate clause** that is neither a Playthrough nor a coverage number: *"the cockpit logout
double-click defect FIXED (no Playthrough added)"*.

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| `D-v28-5-cockpit-logout` | **HALF DONE (iter-16)**; the remaining half specified as **D81**, never implemented |
| the routed instruction | *"prove BOTH halves on ONE rebuild: push the tag, re-pin `stack-demo/rosetta-extensions`, rebuild the `fake-fapi` container, re-run iter-16's four-step browser measurement"* |
| `stack-demo/rosetta-extensions` pin | `fast-build-m256-iter-21` (iter-21 had already re-pinned it — newer than the brief assumed) |
| `handleHandshake` | calls `establishLocked()` **unconditionally** — the mechanism iter-16 named |
| demo-2 | 16 containers Up, 0 exited |

## Cluster / target identified

`D-v28-5` part (b) — **D81**: *an explicit sign-out is sticky until an explicit login.* Chosen over the
onboarding cluster because it is **gate-bearing in its own right** and, unlike the remaining onboarding
UCs, it needed no new seeder capability (iter-24 measured that one).

## Hypothesis

A sticky **sign-OUT** flag — not the negation of `signedIn` — makes the sign-out survive the middleware's
bare handshake without breaking first visit, `autoverify`, or the seat switch.

## Expected lift

The gate's D-v28-5 clause **discharged and proven live**; no Playthrough added (the user's explicit call).

## Phase plan

- **A** — failing-test-first: watch the D81 test go RED on the real defect.
- **B** — implement; mutation-verify every clause of the contract.
- **C** — **prove it LIVE**: tag → push → re-pin the stack clone → rebuild `fake-fapi` from the stack's
  **own** clone (the consumption-copy policy) → re-run iter-16's four-step browser measurement.
- **D** — no-regression gate (every Playthrough logs in through the changed path), fixture restore, close.

## Escalation conditions

- The rebuild leaves the stack unable to log in → **stop and escalate**; `demo-2` is the surface every
  future iter depends on and that outranks this clause.
