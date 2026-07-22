**Type:** tik — the gate-(h) cold reset-to-seed + gate (a)/(b)/(g) live re-verify (under TOK-01). Run 5, tik 1.

# iter-11 — progress

## What landed
1. **Cold reset-to-seed on billion at the m244 pin — GREEN.** Rung-zero re-confirmed (tag on origin, peel b38ad75); re-pinned billion's rext clone `stack-demo/rosetta-extensions` to `sound-check-m244-content-sweep-robustness` FROM ORIGIN (`git describe --exact-match` = the tag, HEAD peels b38ad75), updated `.agentspace/rext.tag`; teardown demo-1 `--purge` (serve-reap 7→0, containers 17→0 — gate (e) re-confirmed); cold bring-up `up-injected.sh 1 --public-host billion.taildc510.ts.net` under `setsid nohup` (survives SSH hiccup). **BRINGUP_EXIT=0**, fresh-green `autoverify.json` `{green:true,warnings:0,ts:2026-07-22T19:31:10Z}` (read directly, known provenance — the M236 age-check bug does not apply), 17 containers, all peer origins serving (web 307 / cockpit 200 / hiring 307 / academy 200 / cosmo 200). The re-bake baked the **correct SCOPED** iter-08 interview demopatch (verified in the baked SSR map: `... || (isManagerScope && !(process.env.NEXT_PUBLIC_POSTHOG_KEY && process.env.NEXT_PUBLIC_POSTHOG_HOST))`).

2. **rext sweep-robustness fix #1 — the denominator cross-check (LOAD-BEARING).** `run-content-stories.sh`'s inline cross-check reimplemented buildPairs()'s pair count but was NOT updated when iter-07 added `player_presence_only`, so it counted **49** while buildPairs() counts **47** — and it `exit 2`'d on the denominator mismatch BEFORE the sweep could run (only a LIVE m244 seed carries the 2 presence-only voice cells that expose it; iter-07/08 never live-swept). Fixed to mirror buildPairs (skip `player_presence_only` for the player pair + the path/seat landable checks). Now 47==47. rext `e30df85`, tag moved → 0c514e6/e30df85, pushed to origin.

3. **rext sweep-robustness fix #2 — the player-interview grader is now EN + IT** (multilingual ack) + a regression test (Italian ack in `<body>`, empty `<main>` — the live shape). 77/77 e2e unit specs green. Defensive robustness (see the honest note under gate (b) — the actual residual root cause is timing, not language).

## Gate parts re-verified LIVE on the fresh seed
- **(a) ORG-CLEAN ✓** — contentsession cleanliness + scrub Go suites PASS (count=1) on the b38ad75 fixtures (byte-identical to the seed source ⇒ 0 source-org tokens in the seed).
- **(e) serve-reap ✓** — teardown moved 7 serve ports + 17 containers → 0 (re-confirmed).
- **(g) interview report renders LIVE ✓** — BOTH `intv-voice-{pass,fail}` MANAGER pairs landed live ("interview manager view rendered with an attempt row", mainLen 290); the plan-section-id alignment assertion was discharged iter-06. The manager interview report renders on the fresh seed = the live-render proof gate (g) wanted.

## Gate (b) — 45/47 (NOT green); the 2 residuals FULLY root-caused + routed
Live sweep (cross-check now passes): **45/47**; the 2 misses are the interview **PLAYER** pairs (`intv-voice-{pass,fail}/player`), reason "no interview acknowledgement text", mainLen 0.

**Root cause (probe-confirmed, NOT a language/measurement issue):** the player interview result renders the ack via the full-page `AISimulationUserHiddenResult`, but it paints `<main>` only after **~25 s** — a probe with `settle()` + a 25 s wait rendered the full ack **in English** (the USER's UI locale, not the session's Italian): "Interview completed / responses have been recorded / Thank you for your time" (mainLen 332, matches the ORIGINAL English regex). The likely slowness is the widened interview **FETCH** gate (`next-web-interview-flag-container` widens `isExtractionEnabled` for BOTH scopes) making the player block on a slow interview-extraction fetch it does not need. The sweep's `settle()` returns EARLY on the nav-shell body (126 chars clears the floor) before `<main>` paints, recording a false `mainLen:0`; the reload-based retry never helps because each reload re-incurs the ~25 s. No SSR error, no client console/page error — a genuine slow-render/settle interaction.

**Honest note:** fix #2 (grader EN+IT) was authored on a WRONG hypothesis (that the ack was Italian). The probe showed the ack is ENGLISH (user locale). Fix #2 is kept as a correct *defensive* multilingual improvement, but it is NOT the fix for this residual. The residual needs the timing fix below.

## Close — 2026-07-22

**Outcome:** billion cold reset-to-seed re-established GREEN at the m244 pin (fresh green autoverify); gates (a)/(e)/(g) re-confirmed LIVE on the fresh seed; the LOAD-BEARING cross-check divergence (49→47) fixed so the sweep can run at all; gate (b) measured live **45/47** with the 2 interview-player residuals fully root-caused to a ~25 s-late ack render vs an early-returning settle (routed). Metric stays **3/8** (gate b not green).
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (3/8; gate b 45/47)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the residual is a rext/demopatch fix, NOT a platform edit) — (5) cap-reached: n (tik 1 of run 5) — (6) protocol-stop: n — Outcome: continue (iter-12)
**Decisions:** D1 (re-pin→teardown→bring-up sequence), D2 (setsid-nohup build + peer poller), D3 (teardown re-confirmed gate e), D4 (cross-check mirror-buildPairs fix), D5 (interview-player residual root cause = slow render vs early settle; two fix options routed) — iter-11/decisions.md
**Side-deliverables:** none (the grader EN+IT fix is folded into the gate-b line, not a side discovery).
**Routes carried forward:**
- **(gate b residual, Fate-3) → iter-12: `GATE-M244-iter12-intv-player-render`.** The 2 interview-player pairs render the ack ~25 s late. TWO fix options: **(preferred)** scope the `next-web-interview-flag-container` FETCH demopatch to `isManagerScope` (mirror iter-08's result-gate scope) so the player skips the slow extraction fetch → ack renders fast → existing settle+retry catches it → 47/47 (needs a next-web/hiring re-bake); **(fallback)** a shape-aware settle-robustness fix (wait for `<main>` content on player-into-main shapes, higher ceiling — careful, don't destabilize the portal case). Confirm the slow-fetch cause first.
- Inherited: **DEF-M239-01** (ENOSPC loud-build-fail) still OPEN; gate (d) `/library` demopatch still OPEN (iter-09).
**Lessons:** (1) a diagnostic conclusion is only as good as its confirmation — the "Italian ack" hypothesis was plausible and WRONG (the ack is the user's UI locale, EN); a probe that dumps the actual rendered `<main>`/`<body>` after a long wait is the cheap decisive check, do it BEFORE authoring a fix. (2) `settle()` clearing its floor on nav-shell chrome (126) while `<main>` is still empty is a false-settle for any into-main shape — the retry can't help a consistently-slow page because a reload re-incurs the slowness. (3) a tooling reimplementation of a source-of-truth function (buildPairs) drifts silently unless a live path exercises it — the cross-check over-counted for 2 releases and only a live seed with presence-only cells exposed it.
