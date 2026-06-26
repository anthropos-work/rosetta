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
- iter-05 (tik): **the MANAGER GATE is MET on a FRESH zero-manual `demo-up`** (the authoritative acceptance —
  iter-04 proved it on a re-seeded live stack; this proves the *build* reproduces it). Tore demo-3 down
  `--purge` + removed the next-web image → a fresh `demo-up` applied the Studio demo-patch (SERVED bundle 0×prod
  / 31×`:39000`; clone git-clean), seeded the FeedbackSeeder mirror (162/162 joinable), reconciled the route,
  reloaded Sentinel, replayed the library + seeded the stories — **all automatically, no manual step**. MANAGER
  sweep: `reachable=70/120 (0,0,0,0) EXHAUSTED gateMet=true` (reproduces iter-04). EMPLOYEE re-sweep on the same
  fresh stack: `reachable=59/150 (0,0,0,0) EXHAUSTED gateMet=true` (**no M42e regression**). One clone-cleanliness
  gap fixed in rext (R1b `.dockerignore` sweep); the build-time disk-exhaustion was environmental (pruned + the
  same fresh `demo-up` resumed idempotently). Zero CANONICAL platform edits. `closed-fixed`, **Gate: MET**. rext
  @ `method-acting-m42m-iter05`. v1.10 reproducibly gate-complete (M42e + M42m) — ready for harden+close. See
  iter-05/progress.md + M42m iter-05 D1–D5.

## M42m: Final Review

Close-milestone review (2026-06-26). Iterative shape → Gate Outcome Ledger (Phase 9-iter). Corpus-side diff is
**docs-only** (the M42m CODE lives in the rext authoring repo @ `method-acting-m42m-harden-final`, hardened +
adversarial-recorded there); the corpus review centres on the 2 touched specs + cross-references.

### Scope
- [x] Gate-distance: gate **MET** (manager 70/120 0,0,0,0 EXHAUSTED; employee 59/150 no-regression), fresh
      zero-manual demo-up → close is "shipped on gate", no carry-forward.
- [x] Iter-ledger: 5 iters (iter-01 bootstrap tok + iters 02–05 tiks), each with a closed `iter-NN/` dir + a
      one-line running-ledger entry; every commit maps to an iter (one-commit-per-iter); 0 orphans.
- [x] Carry-forward queue: DEF-M40-01 manager-half LANDED in-milestone (Fate-1); RESCOPE-1 resolved demo-only
      (not a deferral). 0 routes carried forward.
- [x] No code TODO/FIXME routed forward (corpus diff is docs-only).

### Code Quality
- N/A in corpus (zero code in the corpus diff). The rext code-of-record was reviewed + hardened (final pass,
      stabilized) in the `.agentspace/rosetta-extensions` authoring copy; 0 production bug surfaced.

### Documentation
- [x] [must-fix] `coverage-protocol.md` — mangled phrase "`R1` pristine- s a crash-left patch" → "pristine-reverts"
      (the canonical iter-05 phrasing).
- [x] [should-fix] `frontend-tier.md` (listed in M42m `overview.md` `delivers`, but never updated) had **no**
      demopatch / Studio-escape note → added a callout after the pk+URL build-injection block pointing to
      `coverage-protocol.md`'s "Platform-bound escape" routing row.
- [x] `stories-spec.md` (FeedbackSeeder org-feedback mirror note) verified accurate.
- [x] All markdown cross-refs in the 2 touched docs resolve (10 targets checked).

### Tests & Benchmarks
- N/A in corpus (no test runner). rext tests all GREEN per the harden ledger (Go seeders `-race` ok; Python
      demopatch 43 ok; TS unit spec 17 pass; flake gates 3× clean each).

### Decision Triage
- [x] M42m-D1…D6 → **archive (maintainer-only)**. The user-facing mechanisms (the `/enterprise/*` route model,
      the demopatch, the FeedbackSeeder JOIN-mirror fix, the manager SAMPLE_RULES superset) are already blended
      into `coverage-protocol.md` / `stories-spec.md` / (now) `frontend-tier.md` during the iters + this review;
      the decision records hold the diagnosis detail (route-model falsification, the falsified env-rewrite
      hypothesis, the 6-guard internals) which is maintainer-level. Harden ledger "Knowledge backfill: none
      required" confirmed. No further blend needed.
