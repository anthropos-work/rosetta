# M204 Progress

This is an **`iterative`** milestone — progress accrues per-iter toward the **exit gate** (not a fixed section
checklist). The gate + the iteration protocol live in [`overview.md`](overview.md). Iters are created under
`iter-NN/` by `/developer-kit:build-mstone-iters`. The gate:

> **Same shape as M203, manager-vantage: every declared manager-vantage use case has a passing Playthrough on a
> COLD reset-to-seed demo stack (Dan's Workforce funnel + member roster, member drill-down via the
> activity-dashboard, succession/at-risk via the Growth tab), with 0 false-fails over 5 consecutive reset runs.**

**Status:** `archived` (completed 2026-07-02) — iterative, closed-on-gate at iter-05 (4/4 manager UCs GREEN on
cold reset-to-seed, 0 false-fails / 5). Ran **in parallel with M203** — the shared M202
page-object layer was an additive merge surface.

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

## M204: Final Review (close-milestone, 2026-07-02)

Consolidated from close-milestone Phases 1-5. Iterative **closed-on-gate**. Deferral re-audit (Phase 1b):
**GREEN** — [`audit-deferrals/deferral-audit-2026-07-02-m204-close.md`](audit-deferrals/deferral-audit-2026-07-02-m204-close.md)
(0 repeat/chronic/aged-out; the assign-WRITE UC1 is a single in-manifest Fate-2 tracked gap — D-CLOSE-1).

### Scope
- [x] All 5 iters closed-fixed (1 bootstrap tok + 4 tiks); running ledger complete; 5run-gate.log present. Gate MET (4/4 manager UCs, 0 false-fails/5).
- [x] `assignment-monitoring.assign-and-track.UC1` (assign-WRITE half) → **Fate-2**, tracked in-manifest as `playthrough: TODO` (D-CLOSE-1). Not a repeat-defer, no plan edit. Prior v2.0 close audits (M202/M203) both GREEN.

### Code Quality
- [x] [should-fix] Manager predicate-API intent — the 5 manager `isOn*`/`isIn*` functions have no page-object consumer, BUT neither do the M203 `isOnSkillsTab`/`looksLikeTimelineEntry`. They form a deliberate symmetric pattern↔predicate API pinned by the single-source-agreement block → documented the intent in `url-shapes.ts` (kept, not pruned — the M203-close F3 dead-code posture applies to live-Page METHOD accessors, not this pure predicate library). rext commit c81c6dd.
- [x] Confirmed clean: all 4 manager page objects extend `PageObject`, find-only landmarks, single-sourced `/enterprise/*` route shapes, no bare-`\b`, `@pt` lockstep intact, reporter-override fix correct.

### Documentation
- [x] [must-fix] `playthroughs.md` line ~97 — flipped stale "M204 adds" future-voice → M204 (manager vantage) LANDED; corpus now 10 live Playthroughs, 1 TODO (the assign-WRITE UC1).
- [x] `playthroughs.md` header (~line 13) — "grow" → "grew … to 10 live Playthroughs".
- [x] `playthroughs.md` — added the M204 manager page-object-layer bullet (workforce/members/activity-dashboard/succession + the `/enterprise/*` single-sourced route shapes).
- [x] `README.md` index row — flip M203/M204-pending framing → both landed, "10 live Playthroughs, 1 TODO", (v2.0 M202–M204).
- [x] Confirmed no M204 update needed for `coverage-protocol.md` / `stories-spec.md` (M204 is a consumer, not a change); all cross-refs resolve; no new top-level unit (no new handbook/index row).

### Tests & Benchmarks
- [x] Go `go test ./...` 4/4 pkgs PASS (105 test+fuzz funcs) + vet clean; TS unit 58/58; flake gate 5/5 clean. No gaps — the 3-pass final harden already saturated (100% manifest stmts + drift/presence pins).

### Decision Triage
- [x] D-CLOSE-1 (assign-WRITE UC1 → Fate-2) → blended into `decisions.md` + the audit report (user/dev-facing: explains why the sole manifest TODO is tracked, not dropped). No further knowledge-doc blend needed — the manifest header prose + `playthroughs.md` already carry it.

## M204: Gate Outcome Ledger (iterative — Phase 9-iter)

- **Gate:** target = every declared manager-vantage UC has a PASSING Playthrough on a cold reset-to-seed demo, 0 false-fails/5 runs. **Achieved = 4/4 manager UCs pass, 5/5 reset runs, 0 false-fails (iter-05).** Distance = **0**. **Status: MET → `closed-on-gate`.**
- **Iter ledger:** 5 iters (iter-01 bootstrap tok + iter-02..05 tiks), all closed-fixed. 0 triggered toks, 0 orphan iters, 0 orphan commits (one-commit-per-iter). Every iter has a running-ledger row.
- **Routes carried forward:** none (closed-on-gate — no carry-forward.md).
- **Dropped:** none.
- **Protocol evolution:** none — followed the M203 measure→declare→page-object→play→diagnose→re-measure loop verbatim on the shared M202 foundation (the manager surfaces = an additive merge with M203's employee surfaces, no collision).
- **Sign-off:** not required (closed-on-gate; the gate firing IS the success signal). No escape-hatch entries.
