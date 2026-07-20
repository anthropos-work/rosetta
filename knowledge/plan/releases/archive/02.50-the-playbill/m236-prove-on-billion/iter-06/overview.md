---
iter: 6
milestone: M236
iteration_type: tik
status: closed-fixed
date: 2026-07-20
metric_pre: "27/31"
metric_post: "29/31"
---

# iter-06 — tik: the residual 4

## Step 0 — re-survey

iter-05's reading stands at 27/31. The residual is 4 pairs across 3 distinct causes — small enough that the
"one arm per iter" framing of TOK-01 Phase L no longer applies; this iter targets the residual set.

## Active strategy reference

**TOK-01 "publish-then-prove", Phase L** — the residual.

## Cluster / target identified

1. **2 interview manager pairs** — fail with *no header at all*, a different mode from the 11 fixed in
   iter-05 (which had a header but an empty table).
2. **1 skill-path manager route** — `page.goto` timeout, persisted through the membership-id fix.
3. **1 academy pair** — `/library/<slug>` not-found; the catalog is unseeded (known since iter-03).

## Hypothesis

(1) is likely a **different surface**, not a defect — the interview manager view may not use the
`<player>'s Results for <sim>` header the shape asserts on. Probe before assuming a bug.

## Expected lift

**+2** if (1) is a shape mis-classification. (3) is out of reach without the academy fill.

## Phase plan

1. Probe the interview manager route; read what it actually renders.
2. If it renders correctly → add the calibrated shape. If not → triage as a defect.
3. Re-measure; characterize whatever remains.
