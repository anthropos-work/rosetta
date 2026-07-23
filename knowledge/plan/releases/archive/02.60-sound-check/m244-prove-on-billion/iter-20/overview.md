---
iter: 20
milestone: M244
iteration_type: tik
status: closed
run: 8
created: 2026-07-23
---

# iter-20 — gate (h) COMPLETION: academy course-start + language toggle + p95<5s

**Type:** tik

## Step 0 — Re-survey (mandatory)
Primary metric = exit-gate parts (a–h) discharged green live on billion = **5/8** (a,b,d,e,g). Confirmed
current: billion demo-1 UP + green (17 containers, backend force-recreated iter-18 for the Bedrock fix),
autoverify.json green:true but ts 2026-07-22T19:31:10Z — **STALE >12h** vs the 14400s (4h) green-gate.
TOK-03's named next-target (gate-h completion) is still untouched and meaningful: 4/6 v2.6 fixes proven as
iter-18 byproducts (talk-to-data / library / cockpit UX / content-fidelity); remaining = **academy
course-start** + **language toggle (M241 EN/IT)** + **p95 click→ACCESS < 5s** hero vantages. No substitution.

## Active strategy reference
**TOK-03** (milestone `decisions.md`) — HOLD TOK-01/02 + sequence the narrow final push. This tik is push
step (1): gate (h) completion → ticks **6/8**.

## Cluster / target identified
Gate (h) = "every v2.6 fix proven live; p95 click→ACCESS < 5s hero vantages." 4/6 fixes already live-proven
(iter-18). The residual three proofs, in order:
1. **REFRESH billion's stale autoverify FIRST** (STACK_DIR set) — the green-gate the latency run reads.
2. **p95 click→ACCESS < 5s** both hero vantages (`run-latency.sh 1 employee` + `manager`,
   LATENCY_HOST=billion.taildc510.ts.net, LATENCY_SCHEME=https, LATENCY_GATE_MS=5000, fresh autoverify.json).
   Mind the M236 UTC-vs-local age-check bug — the local authoring copy (6aacc32) carries the `TZ=UTC` fix.
3. **academy course-start** (live check on :13077) + **language toggle EN/IT** (M241) render live.

## Hypothesis
Refreshing the autoverify verdict + proving the 2 remaining fixes render live + p95<5s both heroes
discharges gate (h) → metric 5/8 → **6/8**.

## Expected lift
+1 gate part (gate h ticks). 5/8 → 6/8.

## Phase plan
`corpus/ops/verification.md` (autoverify refresh) + `corpus/ops/demo/tailscale-serve.md` +
`latency-budget.md` (run-latency.sh both heroes, https, gate 5000) + live browser checks (academy
course-start + EN/IT toggle). All live-verification against billion from this tailnet peer; 0 platform edits.

## Escalation conditions
- p95 ≥ 5s persistently on a hero vantage after a fresh green gate → investigate the per-leg attribution
  (latency-budget.md arithmetic signatures); if it needs a platform edit or a heavy re-bake → route/surface.
- academy course-start or EN/IT toggle reveals a real defect that is a platform edit → surface (0 platform
  edits is a hard gate line).

## Acceptable close-no-lift outcomes
- A documented falsification that a remaining fix does NOT render live for a real (routed, non-platform)
  reason — that is a complete iter deliverable even if the metric stays 5/8.
