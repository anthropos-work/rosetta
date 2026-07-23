---
iter: 21
milestone: M244
iteration_type: tik
status: closed
run: 8
created: 2026-07-23
---

# iter-21 — gate (f): the 3 v2.3 drift-carries burned in live

**Type:** tik

## Step 0 — Re-survey
Metric = gate parts (a–h) discharged green live on billion = **6/8** (a,b,d,e,g,h). TOK-03's next-target
after gate (h) is gate (f) — still fully OPEN (0/3 carries burned in). Meaningful, untouched. No substitution.

## Active strategy reference
**TOK-03** — final-push step 2: gate (f) 3 drift-carries → ticks 7/8.

## Cluster / target identified
Gate (f) = "the 3 v2.3 drift-carries burned-in live (BURNIN-M221 / F-M220-4 / PROBE-M218-c3)." Two are
demo-side re-checks on the live billion demo (light); one (BURNIN-M221) is a from-scratch remote DEV
`/dev-up --public-host` bring-up (heavy, ≥1h). Planned multi-line shape (declared): PROBE-M218-c3 +
F-M220-4 this tik; assess + route BURNIN-M221.

## Hypothesis
The 2 demo-side carries burn in live on the existing green demo; BURNIN-M221 needs a separate dev bring-up.
Gate (f) ticks only at 3/3, so the coarse metric may stay 6/8 this tik (2/3 landed).

## Expected lift
gate (f) 0/3 → 2/3 (metric stays 6/8 until 3/3); sets up iter-22 for the 7/8 tick.

## Phase plan
`corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md`: (1) PROBE-M218-c3 — probe billion's
Cosmo router content federation for cms/Directus 403s; (2) F-M220-4 — re-run ant-academy.sh on the live
public-host demo, confirm reap-before-launch rebind + serve; (3) BURNIN-M221 — assess feasibility, route.

## Escalation conditions
- BURNIN-M221 genuinely infeasible (RAM below prereq even reduced, or needs a platform edit) → surface as a
  blocker with specifics, NOT a fabricated pass.
- A carry re-check reveals a real unresolved defect that is a platform edit → surface.

## Acceptable close-no-lift outcomes
- A documented falsification that a carry is NOT burned in for a real (routed/blocker) reason.
