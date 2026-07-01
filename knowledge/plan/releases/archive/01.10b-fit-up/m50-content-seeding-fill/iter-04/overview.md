---
iter: 4
milestone: M50
iteration_type: tik
status: in-progress
created: 2026-06-30
---

# iter-04 — tik (kill the manager-gate prod-eject + verify)

## Active strategy reference
TOK-01 (sweep-driven seed-fill; fix the highest-leverage cluster + re-sweep). The fix surface here is a demopatch
(not a seeder), but it's the same loop: triage (iter-03) → fix → re-apply → re-sweep.

## Cluster / target identified
The SOLE manager-gate blocker from iter-03's gate-valid baseline: `escapes=1` — the
`/enterprise/activity-dashboard/ai-simulations/<simId>` drill-down links to prod
`https://anthropos.work/library/job-simulations/<slug>/`, built on the hardcoded `PUBLIC_WEBSITE_URL`
constant (no env override) → the "platform-bound escape" routing class → a new demopatch.

## Hypothesis
A new `next-web-public-website-url` demopatch (make `PUBLIC_WEBSITE_URL` read `NEXT_PUBLIC_PUBLIC_WEBSITE_URL`
first, prod hardcode as fallback) + baking the demo's own next-web host into it makes the link demo-local →
`escapes` drops 1→0 → manager gate MET (employee already met → full M50 gate met on warm demo-1).

## Expected lift
manager `escapes` 1 → 0, `failingSections` stays 0, frontier still exhausts → manager gateMet=true.

## Phase plan
B (fix-surface confirmed iter-03) → C (author the demopatch + wire up-injected.sh + rebuild demo-1 frontend +
recreate the next-web container, --no-deps) → D (re-sweep manager) → E (close).

## Escalation conditions
If the rebuilt frontend collapses the crawl (broken auth → reachable≈7 — the M46 inject-loop lesson) →
investigate (grep backend clerk logs), not a content fail. If escapes persist (the bundle didn't bake the env) →
re-check the build_env / .env.local overlay.

## Acceptable close-no-lift outcomes
n/a — the fix is a concrete escape-elimination; the re-sweep proves it.
