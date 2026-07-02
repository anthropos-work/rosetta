# M204 Progress

This is an **`iterative`** milestone — progress accrues per-iter toward the **exit gate** (not a fixed section
checklist). The gate + the iteration protocol live in [`overview.md`](overview.md). Iters are created under
`iter-NN/` by `/developer-kit:build-mstone-iters`. The gate:

> **Same shape as M203, manager-vantage: every declared manager-vantage use case has a passing Playthrough on a
> COLD reset-to-seed demo stack (Dan's Workforce funnel + member roster, member drill-down via the
> activity-dashboard, succession/at-risk via the Growth tab), with 0 false-fails over 5 consecutive reset runs.**

**Status:** `planned` — not yet started. Next (after M202 lands): `/developer-kit:build-mstone-iters` (the
BOOTSTRAP iter authors the first coverage strategy as iter-01). Runs **in parallel with M203** — the shared M202
page-object layer is an additive merge surface.

## Running ledger

_(Per-iter entries accrue here as the loop runs: tik/tok, the use case attempted, the failure mode discovered,
the fix surface [page-object registry / dedicated seed / manifest / corpus doc], and the gate delta. The first
entry is the BOOTSTRAP tok — iter-01.)_

| iter | date | use case / surface | outcome | gate delta |
|------|------|--------------------|---------|------------|
| iter-01 | 2026-07-02 | (bootstrap tok) | TOK-01 authored: manager-surface-per-iter | 0 declared / 0 passing (baseline) |
| iter-02 | 2026-07-02 | Workforce funnel + member roster | +2 passing (`skills-funnel.UC1` + `roster.UC1` green) | 0 → 2 passing manager UCs |
| iter-03 | 2026-07-02 | Member drill-down (activity-dashboard) | +1 passing (`assign-and-track.UC2` green) | 2 → 3 passing manager UCs |
| iter-04 | 2026-07-02 | Succession / at-risk (last surface) | +1 passing (`talent-pool.UC1` green) | 3 → 4 passing manager UCs — all journeys functionally green |
| iter-05 | 2026-07-02 | 5-run reset-to-seed determinism gate | **GATE MET** — 0 false-fails / 5, all 4 manager UCs green each reset run | **exit gate FULLY MET** |

- iter-01 (tok/bootstrap): authored TOK-01 (manager-surface-per-iter strategy); baseline 0 manager UCs — see iter-01/progress.md
- iter-02 (tik): Workforce funnel + member roster both green + reconciled `[PASS]`; +2 passing manager UCs; side-fix: runner reporter-override (stale-JSON) — see iter-02/progress.md
- iter-03 (tik): activity-dashboard per-member drill-down green + `[PASS]`; +1 passing manager UC; fixed a SPA-URL race + an out-of-main table scope (D1/D2) — see iter-03/progress.md
- iter-04 (tik): succession/at-risk green + `[PASS]`; +1 passing manager UC → all 3 declared manager journeys / 4 UCs functionally green (no seed expansion needed); DETERMINISM half (5-run gate) next — see iter-04/progress.md
- iter-05 (tik): the 5-run reset-to-seed determinism gate — **0 false-fails / 5, all 4 manager UCs green each run → M204 exit gate FULLY MET** — see iter-05/progress.md + 5run-gate.log
