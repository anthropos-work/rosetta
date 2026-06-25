# M42e Decisions

Implementation decisions with rationale, recorded during the iteration loop (harness home, crawl strategy,
assertion shape, link-rewriting surface, escalations of platform-only blockers).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|

## TOK-01: sweep-then-route-by-leverage — 2026-06-25

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Run the Playwright coverage sweep as the employee hero (`maya-thriving`) against live
demo-3, then iterate **highest-leverage-cluster-first**: each tik runs the sweep (Phase A), triages the
failing pages + escapes by fix surface (Phase B, the routing table in `coverage-protocol.md`), lands the fix
in the routed rext surface (Phase C — `stack-seeding` for empty sections, `stack-snapshot` serve-grants for
content errors, the demo injection/env link-rewriting for escapes, roster/FAPI for identity gaps), re-applies
the affected stack step to the live demo, re-sweeps (Phase D), and closes on whether the targeted cluster
cleared. Drive toward `(failing-pages, escapes) = (0, 0)` over the employee vantage's reachable set.
**Rationale:** The page set + failure modes are discovered by the sweep, not enumerable up front (the reason
the milestone is iterative). M39-M41 already landed the known high-leverage fills (G1 org-name, G2 role
backfill, G3 work/education, G4 avatars, G5 skill depth, G6 library serve-grants), so the sweep chases the
**tail** — the residual empties, the under-investigated G7 activities feed, and any escape links the baked
URLs miss (no studio-host rewrite exists yet). Leverage-first ordering (most pages unblocked per fix) clears
the dominant clusters first; a single serve-grant or seed can light up many pages.
**Strategy class:** new-direction
**Distance-to-gate context:** Gate metric = the coverage report's `(failing-pages, escapes)`; gate = `(0, 0)`
over the employee vantage's reachable pages. Starting value UNMEASURED — the baseline sweep is iter-02 (the
first tik). Known risk areas from `.agentspace/profile_gaps.md`: G7 activities feed (under-investigated), and
escape links (no `NEXT_PUBLIC_STUDIO_URL`-style rewrite baked → a left-menu "Studio" likely escapes to prod).
**Next-tik direction:** iter-02 — run `run-coverage.sh 3 employee maya-thriving` against live demo-3; capture
the baseline `(reachable, failing, escapes)`; triage the highest-leverage failing cluster + pick it as iter-02's
target (or iter-03's if iter-02 establishes the baseline + the first fix in one tik).

