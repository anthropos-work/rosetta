# M248 — Retro (content-stories manager result-link)

## Summary
Re-pointed the content-stories MANAGER CTA off the org activity-dashboard scoreboard to the per-session `/sim`
manager RESULT view (`isManagerView`, persisted read). A bounded `section` re-point of one path builder + its
grader/manifest/docs — but with one genuine unknown (does the interview manager report live on `/sim`?) that
turned out to be the milestone's whole story. Shipped: rext code-of-record `july-jitter-m248-harden @ 6e0ed2c`
(176 unit + Go GREEN, mutation-verified, 3× flake-clean); both corpus docs re-pointed; 47/47 manifest pairs;
0 platform-repo edits. Deferral audit YELLOW (0 blockers).

## Incidents This Cycle
- **P2 — a rung-0 STATIC read was overturned by LIVE evidence (D1 → D3).** The pre-flight code read concluded
  the interview manager report renders on the unified `/sim` route (the two existing demopatches force
  `flag_interview_manager_report` on), so the first implementation collapsed the per-`sim_type` branching. The
  LIVE demo-2 render found the `/sim` interview manager surface falls through to a **"Coming Soon"** placeholder
  (`interviewExtractionData` null — flag/data-gated, not reliably fetched on a demo). Corrected in-milestone:
  interview KEEPS its dedicated activity-dashboard route; only the NON-interview family moves to `/sim`. Caught
  by the milestone's own live render-confirm — exactly what the serial rung-0 verify step exists to force.
- **P2 — the manager grader false-FAILed a fully-rendered result.** The `/sim` manager scored view renders in
  the session's LANGUAGE (Italian for IT sessions) with "Evaluated Skills" collapsed behind a "Show Details"
  toggle, so the player-scored English anchors matched nothing on a complete render. Re-calibrated: `manager-scored`
  keys on the SCORE (N/100) — language-agnostic + collapse-proof — with the 400-char floor. Mutation-pinned at harden.

## What Went Well
- The spec's explicit conditional ("keep interview split IF verify-interview says so") pre-authorized the live
  pivot — the design anticipated exactly this unknown, so the overturn was a clean branch, not a re-scope.
- Live render-confirm before finalizing the grader caught both the routing assumption and the language/collapse
  calibration in one pass. Static-read-only would have shipped a broken interview CTA + a false-failing grader.
- Zero platform edits: the `/sim` manager view already existed and reads the seeded persisted row — a projection
  + grader + doc change, not a build.

## What Didn't
- The static rung-0 read gave a confident-but-wrong answer; only the live render disambiguated. Cost one
  implement→revert cycle (D1 collapse → D3 restore of the interview branch). Cheap here, but a reminder that
  flag/data-gated surfaces are not reliably readable from source alone.

## Carried Forward
- **CARRY-M248-01 → M254 (Fate-2):** re-confirm the content-stories manager pairs land on the FRESH `billion`
  reset-to-seed. M254 exit gate (b) "content-stories manager CTA lands on the /sim per-session result view
  (non-empty)" + (h) the live content-stories sweep already own it — no M254 overview edit. The 3 demo-2
  header-only-shell renders (warm M246-era seed) + 1 academy `:23077` env failure are demo-2 host/warm-seed
  artifacts, not M248 code defects (direct drives prove the route renders full results).

## Metrics Delta
- rext unit specs: 174 (build) → **176** (harden +2 mutation-style manager-scored tests); Go seeders GREEN; tsc clean.
- Manifest: 47/47 landable pairs (21 non-interview → /sim, 2 interview → activity-dashboard); honesty gate GREEN.
- Live (demo-2 warm): content-stories sweep 43/47; residual → CARRY-M248-01.
- Flake: 0. Platform-repo edits: 0. Code-of-record: `july-jitter-m248-harden @ 6e0ed2c` (on origin, rung-zero verified).
- Full metrics: `metrics.json`.
