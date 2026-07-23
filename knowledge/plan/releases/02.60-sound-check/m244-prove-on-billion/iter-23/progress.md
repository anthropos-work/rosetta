# iter-23 — progress

**Type:** tik (run 8, under TOK-03 — gate (c) 16 Playthroughs)

Ran the 16 Playthroughs cold on billion (pt-world reset-to-seed). **12/16 green (89/93 test cases);
the 4 ai-readiness Playthroughs fail deterministically** on a real ai-readiness surface-render gap (the
zero-state renders — the live diagnostic snapshot is absent). Shipped a real harness robustness fix
(networkidle deadlock) surfaced by the re-drive. Gate (c) NOT ticked (needs 16/16); the ai-readiness
live-snapshot gap is routed. Metric stays 6/8.

## What landed

### The 16 Playthroughs, cold on billion — 12/16 green (89 passed / 4 failed, 7.2m)
- **pt-world reset-to-seed** on billion (`run-playthroughs.sh 1 --reset-only`, stackseed on PATH): 54509
  rows seeded, **isolation clean (no shared/external writes)**, sentinel reloaded, 30-identity roster
  exported, fake-FAPI ready. The demo seed is now replaced by pt-world.
- **Browser run from the peer** (`PT_HOST=billion.taildc510.ts.net PT_APP_SCHEME=https`, serial): **89
  passed / 4 failed**. The 12 non-aireadiness Playthroughs (activity-drilldown, hiring, assignment-assign,
  skill-path, simulation, profile, etc.) all GREEN cold over the tailnet.
- The **4 failures are exactly the ai-readiness Playthroughs**: manager-dashboard, manager-howwemeasure,
  member-done, member-progress.

### Harness robustness fix — `PageObject.goto` never gates on `networkidle` (shipped)
Root-caused member-done: `page.goto(/home, { waitUntil: 'networkidle' })` **timed out at 120s** — the
ai-readiness surface polls (recomputes a live snapshot on an interval), so `networkidle` never reaches a
500ms quiet gap. It passed on localhost (fast/sparse) and DEADLOCKED over the tailnet. Fixed
`lib/page-object.ts` goto `networkidle` → `domcontentloaded` (the "never gate on networkidle" doctrine,
latency-budget.md / coverage-protocol.md, re-proven on billion) + rationale comment. tsc clean. rext
**dddef18**, m244 tag moved + pushed (peels `^{}` → dddef18 on origin, rung-zero verified). Harness-only
(runs from the local authoring copy → no billion re-pin).

## The ai-readiness gap — CHARACTERIZED + routed (a real finding, NOT flakiness, NOT a platform edit)
Re-ran the 4 ai-readiness Playthroughs WARM with the networkidle fix: they **still fail deterministically**
on the CONTENT assertion — "AI Readiness Diagnostic" / "Your AI Readiness:" / "Discover your AI Readiness"
absent from `main` (15s `toBeVisible`, even member-done after 1.6m). So it is NOT tailnet flakiness (the fix
removed the goto deadlock; the content is genuinely absent). Ruled out the coordinator's classes:
- **Not tailnet-settle**: content absent after 1.6m + domcontentloaded.
- **Not the manager-MIRROR trap alone**: the MEMBER surfaces fail too (member-done/progress), so it is not
  the manager-scoreboard mirror; it is an org-wide ai-readiness render gap.
- **Prerequisites ALL present**: Vertex Logistics (the heroes' org) has `organization_settings`
  `ai_readiness is_enabled=t`, an active + a closed `ai_readiness_cycles`, 39 `ai_readiness_snapshots`,
  3 steps, 105 `ai_readiness_user_step_progresses`; the `next-web-aireadiness-flag-gate` demopatch is baked
  (demopatch.log empty); the heroes log in on Vertex Logistics.
- **The tell**: `ai_readiness_live_snapshots = 0`, `ai_readiness_recommendations = 0`,
  `ai_readiness_diagnose_narratives = 0` are EMPTY, and the manager spec asserts "a real diagnostic cycle
  (not the 'no cycles yet' zero-state)" — so the surfaces render the ZERO-STATE because the LIVE diagnostic
  snapshot is not populated. Root cause = the pt-world seed does not populate the live snapshot AND/OR the
  on-load recompute (the M219 `app-aireadiness-snapshot-loadmembers` backend read-path + the ~2.09s
  live-recompute) is not effective on billion's pt-world. A seed/demopatch (tooling) fix, NOT a platform
  edit — heavier than this iter's budget (a re-seed with a live-snapshot seeder, or a backend re-bake).

## Re-measure
- **Pre-iter metric:** 6/8. **Post-iter metric:** **6/8** (gate c needs 16/16; 12/16 landed). Delta 0 on the
  binary metric; +12 Playthroughs green cold + a shipped harness fix in substance.

## Close — 2026-07-23

**Outcome:** gate (c) 12/16 Playthroughs green cold on billion (89/93 cases) + a shipped networkidle harness
robustness fix (rext dddef18). The 4 ai-readiness Playthroughs fail on a REAL, characterized ai-readiness
live-snapshot render gap (zero-state; prerequisites all present; live_snapshots/recommendations/narratives
empty) — routed, NOT a platform edit, NOT flakiness. Gate (c) NOT ticked. Metric 6/8. 0 platform edits.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (6/8; gate c 12/16, gate f 2/3)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this iter is a tik; NB the 3-no-prog floor iter-21/22/23 nominally fires at iter-24 — see Lessons) — (3) re-scope: n — (4) user-blocker: n (the ai-readiness gap is a routed tooling finding, not a platform-edit blocker) — (5) cap-reached: n (tik 4/5) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1 (the networkidle-goto deadlock + fix); D2 (the ai-readiness live-snapshot gap characterization + routing).
**Side-deliverables:** the networkidle robustness fix (rext dddef18, tag moved+pushed) — a real found-bug fix, doctrine-aligned; does not tick the gate alone.
**Routes carried forward:** the ai-readiness live-snapshot render gap → a future run (seed a live-snapshot / verify the app-aireadiness-snapshot-loadmembers recompute on billion; re-run the 4 ai-readiness Playthroughs → gate c 16/16 → 7/8). iter-24 = BURNIN-M221 completion (per the orchestrator's course-correction; the pt-world reset already spent the demo seed → the demo-1 teardown BURNIN needs is now free).
**Lessons:** (1) `networkidle` is a deadlock on a polling surface over a high-latency link — the doctrine ("never gate on networkidle") is not optional; assert the RENDERED content, never the network. (2) An empty `ai_readiness_live_snapshots` renders the ai-readiness zero-state even with all other prerequisites present — the live diagnostic snapshot is the load-bearing row for every ai-readiness surface. (3) The 3-no-prog binary-metric floor (iter-21/22/23) nominally fires a triggered tok at iter-24, but TOK-03 already dispositioned this coarse-binary-per-gate artifact (twice) — the strategy is sound + actively progressing (gate h complete, gate f 2/3, 12/16 Playthroughs), so per TOK-03's standing disposition + the orchestrator's explicit course-correction, iter-24 proceeds as the BURNIN tik rather than a redundant fourth re-affirmation tok.
