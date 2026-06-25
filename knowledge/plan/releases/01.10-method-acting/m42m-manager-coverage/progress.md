# M42m — Progress

**Status:** `planned`. **Shape:** iterative (exit gate in `overview.md`).

> `iter-NN/` dirs are created by `/developer-kit:build-mstone-iters` on first run (iter-01 is the bootstrap tok).

## Running ledger
_Appended after each iter (tik = a standard iter toward the gate; tok = a strategy/retro iter). Iter closeouts
append one line each here._

- iter-01 (tok/bootstrap): authored TOK-01 (reconcile-route + clear-escape + populate-dashboard +
  exhaust-frontier); baseline confirmed `/enterprise/` route prefix + TWO fan-outs + universal Studio eject —
  see iter-01/progress.md
- iter-02 (tik): the Studio-link escape (TOK-01 line 1) — DIAGNOSE → RE-SCOPE TRIGGER. next-web `STUDIO_URL`
  is platform-bound (no `NEXT_PUBLIC_STUDIO_URL` override; only knob is a broad+wrong dev-flip); the sole fix
  is a forbidden platform edit. `closed-no-lift` (falsified). RESCOPE-1 awaits user decision — see
  iter-02/progress.md + RESCOPE-1.
- iter-03 (tik): RESCOPE-1 RESOLVED demo-only via the **demo-patch tool** (the user's pivot). Built `demopatch`
  (6 guards + 18 tests, stdlib-only) + the Studio urls.ts patch + up-injected/ensure-clones wiring (R1/R2). A
  FRESH demo-up applies-then-reverts the patch; the baked bundle (0× prod / 31× :39000) + the LIVE dan-manager
  cockpit click-through both resolve Studio to the demo-local studio-desk (:39000); the manager Studio-escape
  class dropped **139 → 0**. CANONICAL platform repos untouched. `closed-fixed`. Gate NOT MET (the M36 dashboard
  populate + fan-out exhaustion — TOK-01 lines 2-4 — remain a later run). rext @ `method-acting-m42m-iter03`.
  See iter-03/progress.md + M42m-D2/D3.
- iter-04 (tik): TOK-01 lines 2-4 — **the MANAGER GATE is MET**. Diagnosed the real route model
  (`/enterprise/workforce` = ONE tabbed SPA, not `/workforce/*`), reconciled `MANAGER_PAGES` to the 6 real
  `/enterprise/*` routes + calibrated (`calibrated:true`); fixed the ONE empty dashboard page
  (`/enterprise/organization-feedback` "No data") via a `stack-seeding` mirror fix (the FeedbackSeeder now
  writes the `local_jobsimulation_sessions` mirror the org-feedback JOIN reads); added 4 manager sample rules
  (2 fan-outs + the 2 inherited library families) so the frontier EXHAUSTS. Authoritative sweep on the
  re-seeded demo-3: `reachable=70/120 failingSections=0 personaFailures=0 escapes=0 notReached=0
  frontier=EXHAUSTED → GATE MET`. Zero CANONICAL platform edits. `closed-fixed`, **Gate: MET**. rext @
  `method-acting-m42m-iter04`. See iter-04/progress.md + M42m-D4/D5.
