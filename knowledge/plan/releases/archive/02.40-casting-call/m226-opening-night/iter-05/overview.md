---
iter: 05
milestone: M226
iteration_type: tik
status: closed-fixed
date: 2026-07-17
---

# iter-05 (tik-4) — fix Finding-4 (C2 harness race) + 2nd clean cold cycle → firm reproducibility

**Active strategy reference:** TOK-01 `reprove-hiring-on-billion` (iter-04's close routed this).

**Step 0 — re-survey:** iter-04b hit 7/7 on ONE default cold cycle (provisional per the M221 one-cycle rule),
with C2 needing 4-5 harness retries (Finding-4). The gate's "reproducibly" clause needs a 2nd clean cycle, and a
reliable automated C2 needs the Finding-4 fix. Both fresh.

## Cluster / target identified
Make C2 reliably automated-green (fix Finding-4) + confirm reproducibility with a 2nd clean default cold cycle.

## Fix (Finding-4 — rext `stack-verify` harness robustness)
`render-hiring-comparison.spec.ts`: after the DOM-rows poll, add a bounded poll for the
`insightsByJobSimulations` LIST op to be CAPTURED (up to 30 s) before deriving the 5 sims. The DOM rows hydrate
from SSR/RSC and can settle before the client-side insights POST fires, so deriving immediately raced the network
capture (cold-stack: ~3/5 runs lost). Non-fatal (SSR-only render falls through to the dom-links fallback). Validated
live: **C2 passed 3/3** on the iter-04b demo (was 3/5).

## Hypothesis
The harness fix makes C2 reliably capture the insights + derive the 5 sims; a 2nd clean default cold cycle at the
fixed tag re-measures all 7 → 7/7 reliably automated-green → the gate is firmly MET (2 cold cycles).

## Phase plan
1. Commit + tag the harness fix (`casting-call-m226-c2-race-fix`); consume on billion (the code-of-record).
2. 2nd clean default cold cycle: `down 1 --purge` → `up-injected.sh 1` (NO FLAGS).
3. Re-measure all 7 from this Mac (C2 now reliable) → target 7/7 automated-green.
4. If 7/7 on the 2nd cold cycle → the gate's "reproducibly" clause is MET.

## Escalation conditions
- A new gap on the 2nd cycle → attribute + fix in-bounds (tooling / demo-patch), never a platform edit.

## Acceptable close-no-lift outcomes
If the 2nd cycle surfaces a characterized gap (falsification), that's a complete cycle → closed-no-lift.

## Close
See `progress.md`.
