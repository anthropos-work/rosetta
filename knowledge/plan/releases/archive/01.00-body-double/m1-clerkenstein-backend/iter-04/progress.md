**Type:** tik

# iter-04 — critical orgclient

## Close — 2026-06-03

**Outcome:** disarmed in-memory orgclient store for the 4 critical methods (fresh-seeded per gene) + runner dispatch + 9 goldens. **`alignctl run`: critical 30.8% → 100.0%, overall 21.1% → 68.4% (13/22).** All 13 critical genes align.
**Type:** tik
**Status:** closed-fixed (planned scope — critical orgclient + goldens — landed)
**Gate:** NOT MET (overall 68.4% < 95%; critical 100% ✓)
**Metric delta:** critical 30.8 → **100.0%** (+69.2); overall 21.1 → **68.4%** (+47.3). +9 genes aligned.
**Phase 5 grading:** (1) gate-met: n (overall < 95%) — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — **Outcome: continue → iter-05**
**Decisions:** M1-D2 (orgclient injection — see milestone `decisions.md`).
**Routes carried forward:** iter-05 — the 9 standard orgclient genes → overall ≥95%.
**Lessons:** the disarmed store seeds the DNA's test scenarios (org o_1, member u_1) and the runner re-seeds a fresh store per gene so genes don't interfere — clean isolation without per-gene teardown.
