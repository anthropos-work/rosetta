---
iteration_type: tik
milestone: M254
iter: 05
status: closed-fixed-partial
---

# M254 · iter-05 — read-only measurement cluster (h-latency solo + c-render)

**Type:** tik · **Active strategy:** TOK-01 (cluster-per-tik live re-prove) — clusters 3 (latency solo) + 2 (coverage c-render).

## Step 0 — Re-survey (done at plan time)
Primary metric = count of a–h parts GREEN cold on billion. Standing: a,b,d MET; c prod-eject side proven.
billion demo UP (16 containers 10h, 0 skillpath), fresh-green autoverify (06:25:41Z, 13 min old, <4h),
5 tailnet origins reachable, cockpit seats maya-thriving/dan-manager present. The two remaining READ-ONLY
gate parts (h-latency, c-render) are bankable on THIS fresh-green demo with no re-bring-up. TOK-01 named
this cluster as clusters 2+3; target still current.

## Cluster / target
- **(h)-latency:** p95 click→ACCESS < 5 s hero vantages (employee maya-thriving + manager dan-manager),
  SOLO / quiet system, `LATENCY_SCHEME=https`, `LATENCY_HOST=billion.taildc510.ts.net`,
  `LATENCY_AUTOVERIFY_JSON`=scratch copy of billion's green verdict, `LATENCY_GATE_MS=5000`, 5 cold runs.
- **(c)-render:** the render side of gate (c) — "← Back to Cockpit" item renders in all 4 apps + studio
  logo/back/logout resolve to the stack app (0 prod-ejects). Prod-eject side already proven iter-04
  (0 escapes/133 pages); this closes the render half.

## Hypothesis
The fresh-green consolidated demo already satisfies both — measure to confirm and bank the gate parts.

## Expected lift
(h)-latency confirmed p95 < 5 s both vantages; (c) fully MET (render side + prod-eject side). Gate ~4/8 → ~5–6/8.

## Phase plan
verification.md measure→confirm: latency runner SOLO (no concurrent load), then coverage runner c-render check.

## Escalation conditions
- latency p95 ≥ 5 s → characterize per-leg, route fix-forward (not a blocker unless structural).
- c-render item missing / resolves off-stack → route to injection link-rewrite fix-iter.

## Acceptable close-no-lift
n/a — these are expected green on the fresh demo; a measured-green is the deliverable.
