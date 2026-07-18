# M226 "opening night" — Retro

_Iterative, `closed-on-gate` — the FINAL v2.4 "casting call" milestone. Closed 2026-07-17._

## Summary

The acceptance-closer: prove the whole v2.4 hiring release **live on `billion` over the tailnet, default `/demo-up`,
no flags** — the M215/M221 "prove-on-VM" pattern applied to the hiring org. Everything the 7-condition gate checks was
already built + local-proven at M222–M225; M226's only new information was **what breaks on a live cross-machine cold
run**. Five findings surfaced and were fixed live (all tooling/harness/seed, **0 platform edits**), and the gate hit
**7/7 reproducibly across 2 clean cold reset-to-seed cycles** + an independent orchestrator re-verify from this Mac.
Recruiter **p95 click→ACCESS 1.74 s < 5 s** landed as the 3rd measured vantage (after v2.3's employee/manager). The
`billion` demo is **LEFT UP** as the live-proof artifact (the M221 precedent).

## Incidents This Cycle

Five cross-machine findings, each surfaced by the live run and fixed within the milestone (none a platform edit):

- **F1 — serve-hiring (blocks C2/C3/C5 cross-machine).** The `apps/hiring` 2nd app (`:3001+off`, added at M224) was
  never registered in `gen_tailscale_serve.py`'s `UI_BROWSER_FACING`, so it had **no HTTPS listener over the tailnet** —
  the recruiter/candidate vantages worked on the dev box yet were dead cross-machine (the M215/M221 "last breakage is
  cross-machine" lesson, exactly). Fix: add `("hiring", 3001)`.
- **F2 — count 3+47 vs 5+45 (C1).** The 3 hiring heroes occupy population slots 0–2; two candidate-hero overrides
  displaced two would-be-admin slots → 3 admin + 47 candidate deterministically, violating the explicit gate and the
  preset's own stated intent. Fix: a one-line preset bump `role_mix.admin 0.1→0.14` (band=7 → 5 net admins survive the
  2 overrides) — no hero moves, no seeder-logic change.
- **F3 — surgical-orphan (self-resolved).** A `:13001` bind-conflict on the first re-bring-up was an **artifact of
  iter-03's surgical serve re-apply** (it added `:13001` to the live serve but not the on-disk reset plan), NOT the
  default flow. The default bring-up writes a reset plan including `:13001`; subsequent teardown→bring-up cycles are
  clean. Non-gate-relevant.
- **F4 — C2 harness insights-capture race.** On a cold stack the DOM rows hydrate from SSR/RSC before the client
  `insightsByJobSimulations` POST fires, so deriving the 5 sims immediately raced the network capture (~3/5 lost). Fix:
  poll for the LIST capture (bounded 30 s, non-fatal). A measurement-reliability flake, not a product regression.
- **F5 — C2 cold/tailnet probe-budget (P2).** Over the slow tailnet a **cold** compare-drawer first render is ~2.5
  min/sim (the R4 hydration), so 5 sequential drilldowns blew the hardcoded 300 s per-test budget → false-fail. Fix:
  `RENDER_TEST_TIMEOUT_MS` env-tunable (default 300000) so a cold/tailnet C2 measurement can't false-fail.

**Close-phase (deterministic) catch — Phase 4 handbook reconciliation:** the stack-seeding README quoted **832** test
funcs; ground truth is **855** (13 pkgs, 83→86 files). Stale since the v2.1 roll, drifted across the whole v2.4 release
(every milestone touched seeders and none reconciled it). Fixed in-place (rext `7032aea`, tag `casting-call-m226-close`).

**No flakes:** the M226-touched deterministic surface is pure (no network/timing) — 5/5 flake gate clean.

## What Went Well

- **The prove-on-VM discipline held.** Measuring on the environment that hurts (the tailnet peer, cold) surfaced 5
  breakages that four milestones of local green never showed — F1 in particular (localhost-only serve) is invisible on
  a dev box. "State the environment with every number" earned its place again.
- **Attribution before fixing.** Every finding was named to a surface (serve registry / seed slot / surgical-artifact /
  harness race / probe budget) before a fix — so all 5 fixes were one-liners or harness tweaks, none a platform edit.
- **The final harden pass caught the cross-iter deterministic regression** the live loop couldn't see: the iter-02
  hiring-front change added a 7th browser-facing port but three `TestTailscaleServe` exact-count/set assertions still
  pinned 6. The live measurement never re-ran the Python suite — the canonical final-harden catch.
- **R4 was resolved by definition, not by fighting it.** The compare-drawer cold render is genuinely slow, but C2 gates
  on data-present-and-renders and C5 on login→ACCESS — so R4 is a documented warm-up transient with a wide-enough probe
  budget, not a gate violation. The M224 blocks-milestone risk retired honestly.

## What Didn't

- **The C1 count bug should have been caught before the live run.** The preset's stated intent (5+45) never matched its
  math (0.1 admin → 3 admins after 2 hero overrides), and no test asserted the NET distribution — `presets_test.go`
  pinned the ratio as a proxy. Fixed by the RED-proven `hiring_count_harden_test.go`, but it cost a live iter.
- **A README count silently drifted for a whole release.** Four v2.4 closes touched seeder tests; none reconciled the
  handbook count. Caught here only because Phase 4's reconciliation rule fired on the milestone-touched module.
- **Surgical shortcuts leave residue.** iter-03's serve re-apply unblocked measurement fast but left the F3 orphan; the
  default full bring-up is the only thing that proves the default path. The pre-bind reap (Finding-3) would close the
  window but is a live-only change — routed forward.

## Carried Forward

- **Finding-3 — the pre-bind serve reap** (Fate 3, non-gate-blocking): clear stale `tailscale serve` fronts on offset
  ports before bind (the M215 F12 window). A bring-up-path change on a **live-only** surface needing a live re-prove
  (forbidden at close; billion left UP); **self-resolves in the default flow**. → a follow-up build-iter with a live
  re-prove, or the next `prove-on-<VM>` milestone. Recorded in `audit-deferrals/deferral-audit-2026-07-17-m226-close.md`
  (DEF-M226-01).
- **8 pre-existing demo-stack test failures** (6× `test_cockpit.py` + `test_purge` + `test_reap`) — HEAD-identical, in
  files M226 never touched, predate v2.4. → **v2.4 release close** → a future demo-stack test-debt harden pass.
- **M204 `assign-and-track.UC1` assign-WRITE** declared in-manifest TODO (`unimplemented`). → **v2.4 release close**.

## Metrics Delta

- **Go test funcs:** 1887 → **1888** (+1; the RED-proven `hiring_count_harden_test.go` net-5-admin/45-candidate fence).
- **M226 deterministic suites (re-run GREEN at close):** stack-seeding `go test ./...` OK · `test_injection.py` 145 /
  8 skipped (`TestTailscaleServe` 34) · stack-verify/e2e `tsc --noEmit` exit 0 · **flake 5/5**.
- **Recruiter p95 click→ACCESS:** **1.09 s / 2.36 s** (2 cycles), orchestrator re-verify **1.74 s** < 5 s gate — the
  3rd measured vantage.
- **Gate:** **7/7 MET**, reproducible (2 cold cycles + orchestrator re-verify). **Flake 0. Platform-repo edits 0.**
- **rext code-of-record:** `casting-call-m226-close` (`7032aea`). Deferral audit **YELLOW**.

(Full machine-readable delta: [`metrics.json`](metrics.json).)
