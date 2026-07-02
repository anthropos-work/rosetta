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

- iter-01 (tok/bootstrap): authored TOK-01 (manager-surface-per-iter strategy); baseline 0 manager UCs — see iter-01/progress.md
- iter-02 (tik): Workforce funnel + member roster both green + reconciled `[PASS]`; +2 passing manager UCs; side-fix: runner reporter-override (stale-JSON) — see iter-02/progress.md
