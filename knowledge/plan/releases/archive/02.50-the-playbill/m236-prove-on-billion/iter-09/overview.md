---
milestone: M236
iter: 9
iteration_type: tik
status: closed-fixed
created: 2026-07-20
handler: LATENCY-M236-iterTBD-hero-p95
---

# iter-09 — the p95 click→ACCESS measurement (hero vantages)

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove) — Phase H, the budget half.

## Step 0 — re-survey

Primary metric **29/29** after iter-08. Two gate components outstanding; this iter takes the first:
**p95 click→ACCESS < 5 s for the HERO vantages only** (user decision B2 — content-seat latency is
explicitly out of scope for v2.5).

## Cluster / target identified

Handler `LATENCY-M236-iterTBD-hero-p95`. The instrument already exists and is the mandated one:
`rext stack-verify/e2e/run-latency.sh` (M218). It must run **from a machine on the tailnet** — the
presenter's actual vantage — which this workstation is. `LATENCY_SCHEME=https` is required against a
`--public-host` demo (F-M221-06b), and the harness refuses to measure a stack whose `autoverify.json`
is not green (the M217 barrier).

## Hypothesis

The measurement is takeable as-is. The shipped M218 numbers were cold p95 2413 ms (employee) / 1767 ms
(manager) on a laptop; `latency-budget.md` warns the same defect cost ~6 s on a laptop and ~112 s on the
tailnet VM, so **the environment must be stated with the number** and the tailnet reading may differ
substantially from the M218 laptop baseline.

## Expected lift

A recorded p95 per hero vantage, graded against the 5000 ms gate. No numerator movement (the primary
metric is already met) — this closes a distinct gate component.

## Phase plan

- **Step 1** — refresh `autoverify.json` to a fresh green (the harness's own precondition).
- **Step 2** — measure the employee vantage (`maya-thriving`), 5 runs, from the tailnet.
- **Step 3** — measure the manager vantage (`dan-manager`), 5 runs.
- **Step 4** — grade against 5000 ms and record with the environment stated.

## Escalation conditions

- **If the budget is blown, check iter-08 D5 FIRST**: `apps/web` carries
  `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql` — the **non-offset** port, dead for
  demo-1 — while `apps/hiring` carries the correct offset origin. A client fetch to a dead address is the
  *fast-failing* arithmetic signature (`≈ 3 × 33 ms + 6 s`) `latency-budget.md` teaches to read before
  attributing any leg. Do not attribute a leg without checking the arithmetic first.
- If a fix needs a platform edit → forbidden by the gate; route forward and surface.

## Acceptable close-no-lift outcomes

A measurement that does not meet the gate, correctly attributed per the budget's per-leg model, with the
environment stated — that is a complete iter even without a fix, since the gate component is then a known
quantity rather than an unknown.
