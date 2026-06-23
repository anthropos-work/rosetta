# M36 — Retro

## Summary
M36 delivered the second product Must — the org **Workforce-Intelligence dashboard** renders believably for a
seeded story. Six new `stack-seeding` seeders (`membership_skills` the mapped→verified funnel joined on skill
_name_, `tags`/`membership_tags` teams + a `mentor` tag, `organization_target_roles`+`user_target_roles` gap +
mobility, `succession` interview feeders, `feedback` ~2:1 positive, `population_evidence` the org-scale
claimed-vs-verified gap + AI-readiness) + two fixes (the assignments status-mix on the 3-table skill-path FK
chain, the `skillpath_sessions` completed-share ~1%→~30%) + the closure gene extended to a 4th skill-ref
surface. The done-bar was met live on the `--local-content` demo-3 Cervato org — every aggregate populated +
distributed (funnel drop-off, 172/137 over/under-claimers, succession 29% interviewed → "full", feedback 2.3:1,
70 active / 40 completed assignments). Tooling + docs only — zero platform-repo edits. Close GREEN, 4 findings,
0 blocking, merged into `release/01.90-storytelling`.

## Incidents This Cycle
None. No P2 flakes, no regressions. The 3 harden passes surfaced **zero bugs** — every edge/error path
confirmed correct under deeper probing. The flake gate ran 5/5 clean (shuffled, `-race`). The one code-quality
finding at close (seeder error-prefix dash style) was investigated to a non-defect (uniform convention + tested
error-contract), not a fix.

## What Went Well
- **The risk register's #1 concern (scope creep — "many widgets") never materialized.** The hard scope line
  ("seed the spine for the seeded story, don't chase every widget") held: exactly 6 seeders + 2 fixes landed the
  spine, and the integration test confirmed every aggregate *function* resolves without chasing per-widget
  polish.
- **The O4 live-introspection pass (one `\d` on demo-3) paid off.** Pinning the migrated column/storage-key
  names up front (spec-notes.md §O4) meant every COPY was schema-correct on the first live run — no
  schema-mismatch debugging.
- **The name-vs-node-id funnel join (D-M36-1) was caught at design, not in production.** The non-obvious fact
  that the funnel joins on skill `name` (not node-id) drove the `skillref_named.go` resolver — names match by
  construction because they draw from the same replayed taxonomy as the verified chain.
- **The admin-view insight unblocked the done-bar without M37.** Unlike M34/M35 (which need login-AS-a-hero for
  the individual profile), M36's dashboard is the admin view — viewable by the demo login user now, so the
  browser render acceptance was achievable immediately.

## What Didn't
- **The Phase 2 code-quality agent flagged a false-positive "should-fix"** (the error-prefix dash inconsistency)
  that cost an investigation cycle to confirm as a non-defect. Minor — the investigation was cheap and the
  convention is genuinely uniform.
- **The seeders package coverage dipped** (98.0% M35 → 95.5% M36) because 6 new seeders added a large statement
  surface; the harden passes recovered it but didn't restore the M35 peak. Acceptable — the load-bearing
  helpers, the named resolver, and `assignments.go` are all at 100%; the residual is defensive registration
  glue.

## Carried Forward
- None. M36 introduced no deferrals; the inherited #M34-D7 had already landed as D-M35-4. The 4 v1.9 backlog
  items (M33, DEF-M10-01, DEF-M21-01, M25-D9) remain orthogonal to seeding (re-confirmed GREEN).

## Metrics Delta (from metrics.json)
- **Go test funcs (stack-seeding):** 347 → **406** (+59); 484 incl. subtests.
- **Coverage:** blueprint 100.0% · seeders 93.0%→**95.5%** (3 harden passes, +2.5) · dna 87.7%.
- **Flake count:** 0 (gate 5/5 sequential shuffled `-race`).
- **Supply-chain:** GREEN (0 new third-party deps).
- **Alignment gates:** 100%/100% (untouched — M36 is stack-seeding only).
- **Findings at close:** 4 (1 code-quality should-fix [non-defect], 2 docs, 1 test), 0 blocking.
