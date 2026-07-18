# v2.4 "casting call" — Release Retro

## Summary
The **recruiter-vantage / hiring-org release**: a NET-NEW 4th, hiring demo org (Meridian Talent) where 45 candidates
audition on the same 5 positions and a recruiter compares them side-by-side, distinct from the 3 workforce orgs.
Shipped in 7 milestones (M222→M228), **proven live on billion**, **0 platform-repo edits**. Mid-release RE-OPEN
(M227/M228) added the believability layer from live feedback — the demo now not only *works* but *reads* as real.

## Incidents this cycle
- **[P1] The M227 fix-#1 seeder guard was incomplete (M228 F1/F2/F3).** The first corrected-data cold bring-up on
  billion crashed the seed (SuccessionSeeder FK to a now-skipped session) and leaked a training sim + a 2nd session
  per candidate (FeedbackSeeder mirror rows). The deterministic M227 unit test structurally couldn't catch it — the
  live re-prove did. Fixed M228 iter-03; the regression now enumerates all 8 generic seeders. **The headline lesson:
  live-proving catches a class of bug deterministic tests cannot.**
- **[P2] C2 render "failed" 4× before the real cause surfaced (M228).** The recruiter comparison drawer is a Next.js
  intercepting route (only the first sim per page-load is cleanly detected); several long cold-tailnet render cycles
  were spent on blind hypotheses before reading the trace/screenshots named the cause. Fixed via `RENDER_ONLY_SIM`
  (prove each sim as "the first"). No production impact — a test-harness limitation.
- **[P1, from M226] The billion demo was days-stale between runs** (reset-to-seed hadn't re-run) — surfaced the
  "cold reset-to-seed" hygiene the M228 gate enforced.

## Cross-milestone patterns
- **Live-prove > deterministic-only** for demo believability: M226 and M228 both caught breakages (product-boundary
  eject; incomplete seeder guard) that milestone-local unit tests passed clean. The "prove-on-billion" pattern earned
  its place as the terminal milestone shape for demo releases.
- **The mirror-table read-model** (M222) was load-bearing all the way to M228 — every scoreboard/recruiter surface
  reads `public.local_jobsimulation_sessions`, not `jobsimulation.sessions`.
- **Zero platform edits held across all 7 milestones** — every platform-source wall routed to a sha-pinned demo-patch
  or a config seam.

## Carried forward (accepted carries — NOT v2.4 regressions)
Standing rext demo-stack tooling test-debt + environment preconditions (all in the rext repo, unaffected by the
rosetta docs merge+tag; the delivered tooling is proven working on billion):
1. **6 "against live clone" failures** (test_ssr_origin_chain ×4, test_demopatch ×2) — need a pristine local platform
   clone; none is present (local demo torn down M225). Test logic last touched M219 (v2.3). **Route:** a future rext
   hardening should make these skip-when-no-clone rather than fail (test-robustness), or stage a clone in CI.
2. **6 test_cockpit academy/overlay failures** — pre-existing (ant-academy link + overlay-JS assertions).
3. **test_purge THE_BUG** (M218-era) + **test_host_prereqs_m215 F12** (M220-era) — pre-existing.
4. **pytest not wired into a repeatable env** — the demo-stack suite needed a manual python3.12+pytest to run.
   **Route:** wire a rext CI/dev-env harness (the close-release triple-clean fallback).
5. **M204 assign-WRITE ptvalidate TODO** + **DEF-M226-01 pre-bind serve reap** — inherited standing carries.

These are rext tooling-quality items (not user-facing demo defects); route to a future rext hardening pass or a
dedicated tooling-debt milestone. None blocks the v2.4 docs+tooling ship.

## Metrics delta
- 0 platform edits (7/7 milestones) · supply-chain GREEN (0 CVE, 0 new deps) · rext seeders 96.8% cov · flake 0.
- Live proof: M228 7/7 on billion, recruiter p95 click→ACCESS 1.27 s, render 5/5 (8,8,9,9,8), 2 heroes usable.
- See `metrics.json` + `metrics-history.md`.
