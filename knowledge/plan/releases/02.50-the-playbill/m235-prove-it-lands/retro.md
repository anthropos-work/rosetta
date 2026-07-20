# M235 — Retro

## Summary
M235 built + unit-proved everything the live "prove-it-lands" proof depends on, **0 platform-repo edits**, all in rext
`stack-seeding` + the rosetta corpus: the provably-clean **13-session simulation matrix** (assessment PASSED = 2 voice /
1 code / 1 document; every type present in passed AND not-passed) and all **3 non-simulation content-story sections**
(skill-path-legacy real progress + the `local_skill_path_sessions` mirror; ai-labs presence-only; academy
`/library/<slug>` CTA) via a separate code-owned registry (`seeders/content_nonsim.go`) — so the manifest now projects
all **4 products / 18 sessions**, both honesty gates GREEN. Iterative, **closed-incomplete** under the user's
pragmatic-close mandate: the LIVE `(session × action)`-lands browser proof needs a running stack and routes to **M236**
by design. Merged `--no-ff` into `release/02.50-the-playbill`.

## Incidents This Cycle
- **No regressions, no new test failures, no flakes.** Harden Pass 1+2 (`--final`, same-day) found 0 bugs; the close
  review added 0 rext code changes; flake gate 5/5. The 6 `test_cockpit.py` failures are the chronic demo-stack
  test-debt carry (unchanged; M235 added 0 new).
- **USER-BLOCKER-M235-01 (the real find, P1-severity data-safety):** the anonymization scrub was removing **zero** names
  — 8/9 shipped M232 fixtures leaked a real customer first name. Root cause: the capture built its scrub map only from
  `jobsimulation.actors.username/.alias` (empty for these sessions), never from the session **owner's `public.users`
  identity**, which is where the candidate's first name actually lives (threaded through the LLM feedback). Resolved by
  the user's "fix scrub + re-capture" ruling → owner-identity sourcing + token-split + word-boundary matching + a
  capture-time `SurvivingToken` fail-closed post-condition + a standing CI cleanliness tripwire; re-captured 9 fixtures
  provably clean (0 leaked names, 545 placeholders). **A status artifact (the M232 fixtures) had shipped a claim the
  code did not enforce — the D17 hazard again; the fix binds it with an executable fail-closed gate.**
- **USER-BLOCKER-M235-02 (a planning-assumption miss, not a defect):** TOK-01's planned "coverage descriptor" mechanism
  doesn't exist — the exact-path/hero-crawl coverage harness structurally cannot reach the dynamic-URL,
  cockpit-seat-reached content-stories result pages; that needs NEW seat-login sweep plumbing authored + calibrated
  against a live render. Resolved by the user's "build non-sim seeders, then close" ruling; the plumbing + live proof
  route to M236.

## What Went Well
- **Both blockers surfaced BEFORE any wrong code landed.** M235-01 was caught in an iter-02 re-survey *before* extending
  the fixture with 4 more real sessions (which would have expanded the PII footprint); M235-02 in iter-05 planning
  *before* a single line of a mis-shaped descriptor. Surfacing-then-deciding (not building-then-discovering) is what
  kept the close clean.
- **The offline/live split was drawn honestly.** M235 landed 100% of what's provable without a browser and unit-proved
  it (fixtures, non-sim seeders, manifest projection, cross-arm owner invariant); it never faked a live render and never
  platform-edited to force one. The gate is "not met by design", not "not met and papered over".
- **The non-sim registry mirrors the proven simulation design.** `content_nonsim.go` reuses the same single-source
  flat-index owner pairing + fail-closed drop discipline as the M232/M233 simulation path, so the projection and seeder
  can't diverge for the shipped registry (pinned by the cross-arm owner-consistency test).

## What Didn't
- **Nothing blocking.** The friction was the two user-blockers — both genuine decisions (a data-controller call + a
  scope/sequencing call), both resolved by the user, both landed. The close itself was near-clean: one adversarial
  scenario recorded (unreachable), two back-ref tags, no rext code change.
- **A latent design smell noted (not a bug):** the seeder doesn't re-run the projection's drop-gate, so a *future*
  config-driven registry with a malformed exhibit could seed an orphan row the manifest dropped. Unreachable with the
  compile-time registry; recorded in the Adversarial-review subsection with the one-line fix for whoever makes the
  registry dynamic.

## Carried Forward
All three clusters are Fate-3 → **M236**, already applied to M236's `overview.md` `In:` (iter-08, commit `54eaefe`,
user-authorized). See `carry-forward.md`.
- **The LIVE `(session × action)`-lands proof + the new content-stories seat-login coverage/Playthrough plumbing** → M236.
- **The per-section live-calibration checklists** (skill-path version/status/mirror; ai-labs `lab_sessions` DDL; academy
  progress-write via `academy-seed` + route + M230 catalog fill) → M236.
- **The M230 carry-forward live items** (ANT_ACADEMY coverage descriptor + next-web clone re-anchor +
  `getPublicCatalogView` 2nd manifest) → M236.
- **14 pre-existing demo-stack test failures** (M235's slice = 6 `test_cockpit.py`) → v2.5 release-close re-anchor
  (REPEAT/CHRONIC, user-dispositioned; M235 added 0 new).

## Metrics Delta (from metrics.json)
- **Go test funcs (whole rext repo, `git grep '^func Test'`):** 1939 (M234 @ fd457bf) → **1974** (M235 @ 60eff14), **+35**.
- **Coverage (harden, touched files):** scrub 94.6% → **100%**; contentsession 93.7% → **94.7%**; seeders 95.9% →
  **96.1%**; cmd/content-capture 0% → **28.7%** (pure helpers; the DB-bound capture path is the M236 live surface).
- **Fixtures:** 9 → **13 sessions**, provably clean (0 leaked names, 545 `<<ACTOR_0>>` placeholders). Manifest: **4
  products / 18 sessions**, both honesty gates GREEN.
- **Flake:** 0 (5/5 Go). **Platform-repo edits:** 0. **Supply chain:** 0 net-new deps. **Alignment:** 100%/100% (untouched).
