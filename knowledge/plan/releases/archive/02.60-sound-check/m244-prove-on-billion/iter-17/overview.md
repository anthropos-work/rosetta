---
iter: 17
milestone: M244
iteration_type: tik
status: closed-fixed-partial
created: 2026-07-23
---

# iter-17 — gate (c): the discrete stack-verify specs on billion (run driveable + map the rest)

**Type:** tik · **Active strategy:** TOK-02 (sweep the remaining gates live)

## Active strategy reference
TOK-02. After iter-16 landed the coverage half of gate (c), the remaining gate-(c) stack-verify work is the DISCRETE specs (calibrate*/persona/verify-*/m220/m224/talk-to-data/enterprise-surfaces/render-hiring/cockpit-overlay/smoke).

## Cluster / target identified
Run the discrete stack-verify specs against billion on the demo seed (no reset) and land what's green; the 16 Playthroughs run LAST (pt-world reset — iter-14 D1).

## Hypothesis
The discrete specs run green against billion's tailnet HTTPS from the peer harness.

## Expected lift
Advance gate (c)'s discrete-spec half. (Gate (c) ticks only when ALL 40 pass, so the part likely doesn't tick this tik.)

## Phase plan
Assess each spec's host/env contract → run the remote-driveable ones → characterize the rest → route.

## Escalation conditions
A spec failure requiring a platform edit → STOP. (None hit.)

## Acceptable close-no-lift outcomes
Characterizing which discrete specs are remote-capable vs need a retrofit is a complete measurement even if the gate part doesn't tick.
