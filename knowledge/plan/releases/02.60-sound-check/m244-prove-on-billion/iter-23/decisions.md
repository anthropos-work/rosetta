# iter-23 — decisions

## D1 — the networkidle-goto deadlock (member-done) + the doctrine fix
member-done failed with `page.goto(/home, { waitUntil: 'networkidle' })` timing out at 120s. The page
snapshot showed the shell RENDERED (banner, "Vertex Logistics · Workforce", the member logged in) — so the
goto did NOT fail to load; it never RETURNED, because the ai-readiness surface polls (recomputes a live
snapshot on an interval) and `networkidle` requires a 500ms quiet network gap that never comes. It passed on
localhost (fast, sparse requests settle) and deadlocked over the tailnet (slower, overlapping requests keep
the network "busy"). Fix: `PageObject.goto` `waitUntil: 'networkidle'` → `'domcontentloaded'`, then let each
spec's semantic `toBeVisible()` (auto-retries) prove the content. This is the standing "never gate on
networkidle" doctrine (latency-budget.md / coverage-protocol.md), re-proven live on billion. rext dddef18,
m244 tag moved + pushed. Harness-only (local authoring copy → no billion re-pin). Safe for the 12 passing
specs: `domcontentloaded` returns earlier, and every spec's `toBeVisible` retries until the content appears.

## D2 — the ai-readiness surface-render gap: a real finding, characterized + routed (not fixed this iter)
With the goto deadlock removed, the 4 ai-readiness Playthroughs STILL fail deterministically on the content
assertion (content absent from `main`; not flakiness). Diagnosis against the three classes:
- NOT tailnet-settle (absent after 1.6m + domcontentloaded).
- NOT the manager-MIRROR trap alone (the MEMBER surfaces fail too → org-wide, not a manager scoreboard).
- Every prerequisite is present on billion's pt-world: org `ai_readiness is_enabled=t` on Vertex Logistics
  (the heroes' org), an active + closed cycle, 39 snapshots, 3 steps, 105 user-step-progresses, the
  `next-web-aireadiness-flag-gate` demopatch baked (demopatch.log empty), heroes logging in.
- THE TELL: `ai_readiness_live_snapshots=0`, `ai_readiness_recommendations=0`,
  `ai_readiness_diagnose_narratives=0` are EMPTY, and the manager spec expects "a real diagnostic cycle (not
  the 'no cycles yet' zero-state)". ⇒ the surfaces render the ZERO-STATE because the LIVE diagnostic snapshot
  is not populated: either the pt-world seed does not seed the live snapshot, or the on-load recompute (the
  M219 `app-aireadiness-snapshot-loadmembers` read-path + the ~2.09s live-recompute) is not effective on
  billion. A seed/demopatch tooling fix (a live-snapshot seeder re-seed, or a backend recompute re-bake),
  **NOT a platform edit** and **NOT a SEVERITY=blocker** — heavier than this iter's budget.
- **Routed** to a future run: seed/verify the ai-readiness live snapshot on billion's pt-world, re-run the 4
  ai-readiness Playthroughs → gate c 16/16 → 7/8. The 12 non-aireadiness Playthroughs are green cold.

## D3 — sequencing: playthroughs before BURNIN (TOK-03 refinement); iter-24 = BURNIN despite the tok-floor
Re-ordered the playthroughs ahead of BURNIN under a Step-0 re-survey: TOK-03's "playthroughs LAST" rationale
(the reset destroys the demo seed, which gate-f's demo carries + gate-h needed) is fully DISCHARGED (gate h
done iter-20; gate-f demo carries done iter-21). The demo seed is now spent by the pt-world reset, so the
demo-1 teardown BURNIN needs is free. The 3-no-prog binary-metric floor (iter-21/22/23) nominally fires a
triggered tok at iter-24, but TOK-03 already dispositioned this coarse-binary-per-gate artifact (twice, at
TOK-02 and TOK-03); the strategy is sound + actively progressing, so per TOK-03's standing disposition + the
orchestrator's explicit course-correction, iter-24 proceeds as the BURNIN tik (the substantive final push),
not a redundant fourth re-affirmation tok.
