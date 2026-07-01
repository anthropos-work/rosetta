---
iter: 3
milestone: M50
iteration_type: tik
status: in-progress
created: 2026-06-30
---

# iter-03 — tik (manager baseline to completion + verify member-field fix)

## Active strategy reference
TOK-01 (sweep-driven seed-fill; re-seed-to-iterate).

## Step 0 — Re-survey (CORRECTS iter-02's "manager sweep will cap" assumption)
iter-02 stopped the manager sweep at page 39 (q=172) ASSUMING it would cap at the 150-page `COVERAGE_MAX_PAGES`.
Re-surveying `crawl.ts`: the sample logic caps **VISITED** pages (`pages.length < maxPages`); sampled-out paths
only inflate the QUEUE (`q`), not `pages.length` — they're skipped fast at dequeue (`continue`, no `page.goto`).
The manager SAMPLE_RULES DO match the fan-outs (`/user/<uuid>` sample 8, `/enterprise/activity-dashboard/<type>/<uuid>`
sample 8, `/sim` 20, `/skill-path` 12 — regex-verified). So the manager visits ~55-65 pages then DRAINS the
large queue via fast sampled-out skips → it would **EXHAUST, not cap**. iter-02's cap assumption was likely
wrong; the sweep was stopped prematurely.

## Cluster / target identified
Re-run the manager baseline **to completion** (with `COVERAGE_MAX_PAGES=300` as a safety margin) to get the
REAL verdict. If it exhausts (likely): a gate-valid manager baseline → triage the actual failing sections. If it
genuinely caps: iter-04 becomes a tooling-iter (tighten SAMPLE_RULES). Either way, also confirm whether the
member-field fix needs re-seeding (it fills NULL cols the current manifest doesn't assert, so it likely won't
move the CURRENT metric — the D4/F1 manifest-strengthening is the path that makes it count).

## Hypothesis
The manager baseline sweep exhausts under a 300-page cap and reports a real `(failingSections, escapes)` — the
true manager starting point, which the iter-02 premature stop denied.

## Expected lift
A gate-valid manager baseline measurement (the metric the manager half of the gate is read against). Not a fix
landing per se — this tik recovers the measurement iter-02 missed + triages the real manager gaps.

## Phase plan
A (manager sweep to completion, cap=300) → B (triage real failing sections) → C/D/E as the verdict dictates
(fix the highest-leverage cluster if the sweep is fast + a clear gap surfaces; else close on the recovered
baseline + route the triaged fixes forward).

## Escalation conditions
A failing section needing a platform edit → re-scope-trigger. A genuine cap (frontier > 300) → iter-04
tooling-iter (SAMPLE_RULES). A surfaced unrelated bug → route forward.

## Acceptable close-no-lift outcomes
Recovering a gate-valid manager baseline + triaging the real failing sections is a valid tik outcome even
without a fix landing this iter (it corrects the iter-02 measurement gap + sets the real target).
