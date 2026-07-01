---
iter: 5
milestone: M50
iteration_type: tik
status: in-progress
created: 2026-06-30
---

# iter-05 — tik (the residual content-field escape: stack-snapshot content-URL rewrite)

## Active strategy reference
TOK-01 (sweep-driven seed-fill; fix the escape cluster + re-sweep).

## Cluster / target identified
iter-04's residual: the manager's SOLE escape is the replayed Directus content
`directus.simulations.public_landing_page_url` (+ `read_more_link`) carrying a hardcoded
`https://anthropos.work/library/job-simulations/<slug>` (replayed from prod). The activity-dashboard sim
drill-down renders it → prod-eject. NOT a JS constant (the iter-04 demopatch can't reach DB content).

## Hypothesis
A post-replay content-URL rewrite (rewrite any `anthropos.work` host in those Directus fields → the demo's own
next-web host) makes the link demo-local → `escapes` 1→0 → manager gate MET → (employee already met) full M50
gate met on warm demo-1.

## Phase plan
C (author the rewrite in up-injected.sh + tests; apply to demo-1 + clear cms Redis sim cache) → D (re-sweep
manager) → E (close).

## Escalation conditions
If a DIFFERENT escape host surfaces (another content field / another host) → broaden the rewrite + re-sweep
(same iter, planned-scope). If a NON-content escape surfaces → triage to its surface, route forward.

## Acceptable close-no-lift outcomes
n/a — the rewrite is a concrete escape-elimination; the re-sweep proves it. (If the gate STILL doesn't clear due
to an unrelated newly-surfaced issue, record the falsification + route it.)

## Fix landed (Phase C)
The content-URL rewrite is a post-replay idempotent UPDATE in `up-injected.sh`'s NO_SETDRESS block (regex over
`https?://[a-z0-9.-]*anthropos\.work` → `http://localhost:$((3000+OFFSET))`, demo-local, non-fatal,
`DEMO_NO_CONTENT_URL_REWRITE=1` opt-out) — the same mechanism class as the M46 FK indexes / Directus column
backfill. +4 static-fence tests (frontend-build 55→59 GREEN). Applied to demo-1 (0 anthropos.work URLs remain,
incl the 5 staging-host rows the broadened regex caught) + cleared 161 cms Redis DB-5 sim cache entries.
