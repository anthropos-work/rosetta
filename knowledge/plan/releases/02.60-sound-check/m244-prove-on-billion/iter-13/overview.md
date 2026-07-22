---
iter: 13
milestone: M244
iteration_type: tik
status: closed-fixed
created: 2026-07-22
---

# iter-13 — land gate (b) 47/47

**Type:** tik · **Active strategy:** TOK-02 (consolidate + sharpen TOK-01; gate-b fix first, then the cheap live gates)

## Active strategy reference
TOK-02 (milestone-root `decisions.md`). Run-6 priority 1: land gate (b) 47/47, then sweep the remaining gates live on the green m244 seed.

## Cluster / target identified
The single gate-(b) residual carried from iter-11: the 2 interview-PLAYER pairs (`intv-voice-pass/player`, `intv-voice-fail/player`) read `mainLen 0` / "no interview acknowledgement text" → live 45/47. TOK-02 named a preferred fix (scope the `next-web-interview-flag-container` FETCH demopatch to `isManagerScope`, mirroring iter-08) with a mandated **probe first**, and an explicit fallback (shape-aware settle-robustness in the sweep runner) if the probe disproves the fetch-cause.

## Hypothesis
Preferred (TOK-02): the unscoped container FETCH-gate makes the player block on a slow extraction fetch → ack ~25 s late. Fallback: the ack simply paints in a later stage than `settle()` waits for, and the sweep early-exits on the chrome plateau.

## Expected lift
Gate (b) 45/47 → **47/47** → metric 3/8 → 4/8.

## Phase plan
Probe (confirm cause) → apply the correct fix → unit test + mutation-verify → re-sweep live on billion → close.

## Escalation conditions
A platform-repo edit required → STOP, SEVERITY=blocker. Neither fix path touches the platform.

## Acceptable close-no-lift outcomes
If the probe had shown the residual needs a platform edit (it did not), close-no-lift with the falsification recorded.
