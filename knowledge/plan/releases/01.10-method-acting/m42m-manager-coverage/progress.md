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
