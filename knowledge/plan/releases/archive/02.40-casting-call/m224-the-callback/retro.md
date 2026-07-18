# M224 "the callback" — retro

## Summary
The render-risk payoff milestone: click **[Log in as]** the recruiter → land on the candidate-comparison Results
scoreboard with ~43 comparable, non-junk candidates (20/page, faithful) on **each** of the 5 shared hiring sims;
the org reads as hiring; the 2 candidate heroes render usable, differentiated self-views. Driven as an iterative
render loop (13 iters, single day) to a **closed-on-gate** close. **Zero platform-repo edits** — code in the rext
authoring clone (tag `casting-call-m224-harden`), docs in rosetta.

## Incidents This Cycle
- **P1 (strategy pivot, caught by measure-first): the M222 `apps/web` premise was false for a *genuine* hiring
  org.** iter-05 attributed the render wall to a **product-boundary eject** (`apps/web` ejects an all-hiring-orgs
  user to the standalone Hiring product), not a render-gate — so TOK-01's single-eject-demo-patch fix path was
  falsified. TOK-02 pivoted to running the **real `apps/hiring`** as a second container. Lesson: front-loading a
  real baseline render (the M219-trap discipline the milestone is named against) turned a would-be wrong fix into
  a correct pivot.
- **P2 (probe render-race, caught by the independent re-verify): R5.** The trustworthy probe (iter-12) still had a
  drawer-hydration poll that sampled cells before they rendered → intermittent false-junk on the last sim
  (1/6 readings). The 3-cold-run had *passed*; the independent orchestrator re-verify (D17 discipline) caught it,
  and a bounded first-row-text wait fixed it → 4/4 clean. **Lesson: only an executable probe binds; a green run is
  not proof the probe is trustworthy** (the D17 keeper, applied at the probe layer).
- **P2 (believability, surfaced by the fixed probe): a hidden Studio prod-eject.** The trustworthy probe surfaced a
  `studio.anthropos.work` eject in the hiring NavBar (shared `packages/ui`, baked `STUDIO_URL`) that the earlier
  broken probe had hidden → killed iter-13 by chaining the shared `urls.ts` pair onto `build_frontend_hiring`.
- 0 code regressions. rext suites GREEN (go all pkgs, python 650/8-pre-existing, tsc clean).

## What Went Well
- **Reuse over re-skin.** TOK-02 exercises the platform's *genuine* product routing (a hiring recruiter really uses
  the separate Hiring site) and reads the *same* seeded `local_jobsimulation_sessions` over the *same* Cosmo
  backend — materially more faithful than a workforce re-skin, and it built entirely from the untouched clone.
- **Conditional-emit kept alignment byte-perfect.** The new FAPI `isHiring` field emits only for hiring orgs, so
  every non-hiring org's `public_metadata` stayed byte-identically `{eid}` — `/align-run` 100/100 on both surfaces,
  no golden re-capture.
- **The close scope review earned its keep.** It caught three overlooked `Delivers → corpus` sections
  (`cockpit-spec.md`, `clerkenstein.md`, `hiring.md` render-path) that had **zero** M224 content pre-close — landed
  Fate-1 rather than shipped as a silent gap.

## What Didn't
- **The gate's "≥40 rows" wording predated knowing the surface paginates at 20.** Resolved by GATE-DECISION D1
  (keep the faithful 20/page; re-interpret the gate as "≥40 comparable seeded + reachable + rendering per sim").
  A cleaner authored gate would have said "seeded + reachable," not "painted on one page."
- **8 pre-existing demo-stack test failures** (test_cockpit.py academy/overlay + purge/reap) were only *observed*
  at M224's harden pass — M222/M223 never ran that suite. Carried to the standing test-debt backlog (D6), not this
  milestone's to fix.

## Carried Forward
- **Standing test-debt backlog:** the 8 pre-existing demo-stack failures → a future demo-stack test-debt harden pass.
- **M225 (dress the set):** the hiring-vantage coverage sweep + one GREEN hiring playthrough + `pt-world` wiring.
- **M226 (opening night):** the live `billion`/tailnet proof of the 7-condition gate, incl. recruiter p95
  click→ACCESS < 5 s (the 3rd measured vantage) + re-proving whatever M224 pinned at final code.

## Metrics Delta
- Go test funcs: 1857 → **1885** (+28). Python demo-stack 650 pass / 8 pre-existing fail; stack-injection 255/8-skip;
  tsc clean. Flake (milestone-owned) **0**. Platform edits **0**. Deferral audit **GREEN**. See `metrics.json`.
