---
milestone: M236
iter: 7
iteration_type: tik
status: closed-fixed
created: 2026-07-20
handler: SKILLPATH-M236-iter07-manager-hang
---

# iter-07 — the skill-path manager "hang"

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove) — Phase L, the residual arm.

## Step 0 — re-survey

Target still current: the sweep's last durable ledger (`content-out/demo-1/pairs.jsonl`) shows exactly the
two failures iter-06 characterized — `skill-path-legacy/sp-product-manager-completed/manager`
(`threw: page.goto: Test timeout of 180000ms exceeded`) and `skill-path-new/academy-foundation-of-ai/player`
(`route rendered a not-found`). Stack still UP on `billion` (1 h), same tag. No substitution needed.

## Cluster / target identified

Handler `SKILLPATH-M236-iter07-manager-hang` — 1 pair, routed forward by iter-06 with a hypothesis
attached: heavy 13-chapter instance hangs while its 3-chapter sibling passes → the **per-item fan-out**
signature `latency-budget.md` teaches. Opening move is to falsify or confirm that signature by naming the
stalling leg before reading product code.

## Hypothesis

Per-item fan-out on the completed 13-chapter path — a per-chapter query the light sibling doesn't pay.

## Expected lift

+1 pair (29/31 → 30/31). Residual after this iter would be academy only.

## Phase plan

- **Step 1 — instrument, don't guess.** Navigate with `commit` (never `networkidle` — `latency-budget.md`),
  record every request's wall time and everything still in flight at the deadline. Name the stalling leg.
- **Step 2 — root-cause** from the named leg.
- **Step 3 — fix in tooling** (seeder / manifest / harness). Zero platform edits.
- **Step 4 — re-measure** the full sweep, not just the target pair.

## Escalation conditions

- If the stall is inside the **platform** with no tooling-side seam → route forward; a platform edit is
  forbidden by the gate.
- If the fix changes the **gate denominator** → do not land silently. Record the evidence in
  `decisions.md`, surface it prominently in the iter close and to the orchestrator.

## Acceptable close-no-lift outcomes

A documented falsification of the fan-out hypothesis that names the true cause is a complete iter even if
the pair does not land — the milestone's own lesson is that three of four "defects" so far were wrong
assertions, not product bugs.
