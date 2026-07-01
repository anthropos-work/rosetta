# M48 — Corpus re-ground — retro

## Summary
Re-grounded the rosetta corpus against the (M47-confirmed-current) clones. The load-bearing deliverable: a NEW
`corpus/services/ai-readiness.md` documenting the previously-undocumented member-AI-readiness feature — and serving
as the **M51 seeder contract**. Plus material-drift reconciliation (4 docs now name the subsystem) and the
ant-academy stale-claim fix. A 3-agent read-only investigation (AI-readiness backend / frontend / corpus drift) did
the heavy reading; the doc synthesis + the close happened in the main thread. Docs-only — never touched the demo.

## Incidents this cycle
- **P3 (caught in review, not shipped) — doc-accuracy bug in the seed contract.** The first draft of
  `ai-readiness.md` offered a "snapshot-direct" seed strategy (write `ai_readiness_live_snapshots`). Phase 2c
  adversarial review + a code read of `computeOrgBreakdowns` showed the **active-cycle dashboard recomputes from
  signals** (`user_skill_evidences` + jobsim sessions) — so snapshot-direct would leave the live dashboard empty
  (that table is a Talk-to-Data cache). Rewrote the contract as cycle-state-dependent before merge. Had this shipped
  un-reviewed, M51 could have built a seeder against a wrong contract. No regressions, no flakes (docs-only).

## What went well
- **The 3-agent fan-out was the right tool.** Investigating a multi-repo, multi-table feature (backend data model +
  frontend surfaces + corpus drift) in parallel produced a complete, file:line-cited picture fast.
- **Verify-against-code paid off twice:** spot-checking the cited paths/enum/9-tables before writing, then the Phase 2c
  recompute-from-signals check that corrected the seed contract. The doc ships accurate, not plausible.
- **Scope held (material-lag-first).** The corpus was broadly accurate; only the AI-readiness blind spot + the
  ant-academy stale claim + 4 recent-feature omissions needed work. No speculative rewrite.

## What didn't
- **The draft seed contract was plausible-but-wrong** until the adversarial pass. Lesson for doc milestones that
  feed a downstream seeder/implementer: trace the *read path* in code, don't infer the write contract from table
  names alone.

## Carried forward
- **D2 → M51:** the seed STRATEGY pick (active+signals-true vs closed+frozen-snapshot) is M51's call; both documented.
- **D3 → M49 #5:** the ant-academy `repos.yml` code fix (M48 fixed the doc; M49 flips it back to "Yes" after adding the entry).
- Authoritative deep-dive lives in the platform repo's own KB (`app/knowledge/ai-readiness/`) — cross-linked.

## Metrics delta
- Docs-only: NEW `corpus/services/ai-readiness.md` + 6 edited docs. No code/tests → rext test total unchanged (1552).
  Flake n/a. 0 new deps. (Full: `metrics.json`.)
