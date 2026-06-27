---
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
---

# iter-01 — bootstrap tok (TOK-01): the manager-coverage strategy

**Type:** tok (bootstrap) — authors the FIRST strategy for M42m. Does NOT terminate the call; the loop
continues into iter-02 (the first tik) under TOK-01.

## Inputs
- `overview.md` (the manager gate — same (a)-(d) bar as M42e, manager vantage, incl. the M36 Workforce
  dashboard) + `corpus/ops/demo/coverage-protocol.md` (the iteration protocol + harness) +
  `corpus/ops/demo/stories-spec.md` §"The Workforce dashboard surfaces (M36)" + the M42e iter-23 close
  (the manager smoke-sweep that calibrated the residual) + the M42e design-plan.
- Phase 0b KB-fidelity: **YELLOW** (docs accurate; 2 load-bearing facts confirmed live — see below).

## Distance-to-gate context (the M42e manager smoke-sweep residual, iter-23)
`dan-manager` vs fresh demo-3: `reachable=150/150 failingSections=0 personaFailures=0 escapes=139
notReached=5 frontier=CAPPED(+79) → GATE NOT MET`.
- **escapes=139** — ALL `studio.anthropos.work`, ONE root cause: a baked left-nav "Studio" prod link the
  manager/enterprise nav renders on every authenticated page (employee nav doesn't → employee had 0).
- **notReached=5** — ALL the manager manifest's `/workforce/*` pages. **KB-fidelity finding: these paths are
  WRONG GUESSES.** The seeded manager `jump_to` is **`/enterprise/workforce?tab=skills-verification`**
  (confirmed in `stories.seed.yaml` + `test_cockpit.py`). The dashboard is ONE tabbed route
  `/enterprise/workforce`, not 5 sub-routes. So notReached=5 is primarily a **manifest route-model error**,
  reconciled by re-authoring MANAGER_PAGES to the tab-query model — THEN any genuine nav/seed gap surfaces.
- **frontier CAPPED(+79)** — the manager's team-roster `/user/<id>` fan-out (per team member) exceeds the cap;
  needs a representative-sample rule (like the employee `/sim/<slug>` rule) + a higher cap to EXHAUST.
- **persona PASS (all 3)** — the M42e identity machinery generalizes to dan-manager for free.
- **failingSections=0 is over the REACHED set only** — the workforce CONTENT gate is UNMEASURED until the
  dashboard route is reached + asserted.

## Two load-bearing facts (Phase 0b, confirmed live)
1. **Manager route = `/enterprise/workforce?tab=skills-verification`** (NOT `/workforce/*`). Reconcile the
   manifest FIRST.
2. **`CORS_EXTRA_ORIGINS` for `/api/workforce/*` is already wired** by the demo injection
   (`gen_injected_override.py:283`, tested) — so an empty dashboard is NOT a CORS gap; check seed/render.

## Initial strategy (TOK-01) — see decisions.md for the full TOK-01 entry
Manager gate = clear escapes + reach-and-populate the (tabbed) Workforce dashboard + exhaust the team-roster
frontier + calibrate the manager manifest. Sequenced by leverage:
1. **Studio-link escape** (clears 139 in one fix IF env-configurable; escalate as re-scope if hardcoded).
2. **Reconcile the manager manifest to the real `/enterprise/workforce` tab route-model** (turns notReached=5
   into a measured content gate) + verify the dashboard renders REAL seeded M36 data (the 6 seeders already
   write it; CORS already wired) — seed/serve-grant only what the live render proves empty.
3. **Team-roster `/user/<id>` sample rule + raise cap** so the frontier EXHAUSTS honestly.
4. **Calibrate** the manager manifest floors/selectors against the live render (`calibrated:false → true`).

## Next-tik direction (iter-02)
Start with the **Studio-link escape** (highest leverage, 139 escapes): diagnose where the baked Studio
left-nav link's host is set in the rendered DOM — is it a `NEXT_PUBLIC_*_URL` the demo injection can rewrite
to the local studio-desk offset port (`:39000`), or hardcoded in next-web (platform, read-only → re-scope
trigger)? Land the rewrite in the demo injection (rext) if configurable; re-build the demo frontend; re-sweep
to confirm escapes→0.

## Baseline sweep finding (this iter, cap 250, no-gate) — TWO exploding fan-outs, not one
The live no-gate manager sweep (demo-3, dan-manager, cap 250) crawled 165 pages then **TIMED OUT** at the
25-min test budget before writing the report — because the manager frontier has **TWO** template-identical
fan-outs (employee had one):
1. **`/user/<id>`** team-roster (+`/skills` +`/activities`) — ~28 members × 3 = the M42e-flagged fan-out.
2. **`/enterprise/activity-dashboard/{ai-simulations,skill-paths,interviews}/<uuid>`** per-activity
   drill-downs (+ a nested `/<uuid>/<uuid>` level) — a SECOND large fan-out the M42e smoke-sweep didn't reach.
The real route prefix is confirmed **`/enterprise/`** (`/enterprise/assignments/*`,
`/enterprise/activity-dashboard/*`) — the manifest's `/workforce/*` guesses match NOTHING (so notReached=5 is
a manifest route error). Nearly every page shows `eject=1` (the baked Studio left-nav link, on every
authenticated manager page). **Implication for TOK-01:** the sample rules are a PRECONDITION to getting a
manager report at all (not just polish) — without them the sweep can't finish the budget. The authoritative
gate sweep must apply sample rules + find the real `/enterprise/workforce` (or the dashboard's actual landing
route) before it can measure the content gate.

## Phase plan
Bootstrap tok: author TOK-01 from the inputs + a fresh baseline reading (no-gate manager sweep on live
demo-3). No fix lands this iter; the deliverable is the strategy + the baseline. **Status:** the baseline
sweep timed out (no report JSON) but the streamed `[crawl]` log gave the load-bearing facts (the `/enterprise`
route prefix, the two fan-outs, the universal Studio eject). The iter-23 smoke-sweep already gave the gate
numbers (escapes=139, notReached=5, frontier CAPPED). Baseline = grounded.
