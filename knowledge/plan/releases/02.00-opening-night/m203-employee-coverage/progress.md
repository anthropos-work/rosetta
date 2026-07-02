# M203 Progress

This is an **`iterative`** milestone — progress accrues per-iter toward the **exit gate** (not a fixed section
checklist). The gate + the iteration protocol live in [`overview.md`](overview.md). Iters are created under
`iter-NN/` by `/developer-kit:build-mstone-iters`. The gate:

> **Every declared EMPLOYEE-vantage use case has a passing Playthrough on a COLD reset-to-seed demo stack (the 3
> employee stories — Skill Paths, AI Simulations NON-voice, Profile), with 0 false-fails over 5 consecutive reset
> runs.**

**Status:** `archived` (completed 2026-07-02) — closed-on-gate. Gate MET at iter-06 (6/6 employee Playthroughs
GREEN on cold reset-to-seed, 5/5 deterministic); merged → `release/02.00-opening-night`; rext tag `opening-night-m203`.

## Running ledger

_(Per-iter entries accrue here as the loop runs: tik/tok, the use case attempted, the failure mode discovered,
the fix surface [page-object registry / dedicated seed / manifest / corpus doc], and the gate delta. The first
entry is the BOOTSTRAP tok — iter-01.)_

| iter | date | use case / surface | outcome | gate delta |
|------|------|--------------------|---------|------------|
| iter-01 (tok/bootstrap) | 2026-07-02 | — (strategy) | TOK-01 authored: deterministic-read-first → mutating → integration boundary | baseline 1/~8 employee UCs — see iter-01/progress.md |
| iter-02 (tik) | 2026-07-02 | Profile: verified + growth (read) | closed-fixed — profile.verified.UC1 + profile.growth.UC1 GREEN on demo-1 (ptreport 3/3, no-regressions) | +2 employee UCs passing — see iter-02/progress.md |
| iter-03 (tik) | 2026-07-02 | Profile: work/education timeline | closed-fixed — profile.timeline.UC1 GREEN; full Profile gate journey covered (ptreport 4/4) | +1 employee UC passing — see iter-03/progress.md |
| iter-04 (tik) | 2026-07-02 | Skill Paths: legacy learn & progress | closed-fixed — skill-paths.legacy.UC1 GREEN (browse→open→start→progress; verify-skill composes P7); 2/3 gate journeys (ptreport 5/5) | +1 employee UC passing — see iter-04/progress.md |
| iter-05 (tik) | 2026-07-02 | AI Simulations: chat launch (NON-voice) | closed-fixed — ai-simulations.chat.UC1 GREEN (§5.8 launch boundary) + post-seed Sentinel Reload fix; ALL 3 gate journeys covered (ptreport 6/6) | +1 employee UC passing — see iter-05/progress.md |
| iter-06 (tik) | 2026-07-02 | 5-run reset-to-seed determinism gate | closed-fixed — **GATE MET**: 5/5 reset runs passed, 0 false-fails (6/6 Playthroughs per run) | **exit gate FULLY MET** — see iter-06/progress.md |

## GATE STATUS: **MET** (2026-07-02)
Every declared employee-vantage use case (6/6 — Profile identity+verified+growth+timeline, Skill Paths legacy,
AI Simulations chat launch) has a PASSING Playthrough on a COLD reset-to-seed demo, with **0 false-fails over 5
consecutive reset runs** (iter-06: 5/5 passed). Next: `/developer-kit:harden-mstone-iters --final` then
`/developer-kit:close-milestone` (which tags `opening-night-m203`).

## M203: Final Review

Consolidated from close-milestone Phases 1-5 (2026-07-02). Deferral re-audit (Phase 1b): **GREEN** —
[`audit-deferrals/deferral-audit-2026-07-02-m203-close.md`](audit-deferrals/deferral-audit-2026-07-02-m203-close.md)
(0 repeat/chronic/aged-out; 4 non-gate UCs → Fate-3 M206, D-CLOSE-1).

### Scope
- [x] Gate MET (iter-06, 6/6 employee UCs, 5/5 cold-reset deterministic) — closed-on-gate, no carry-forward.md needed.
- [x] 4 non-gate edge UCs (code-sim / interview-sim / verify-terminal / self-eval) → Fate-3 → M206 (D-CLOSE-1); roadmap-vision annotated. Academy OUT (by design), voice → M206 (by design) — neither a defer.
- [x] All 6 iters closed-fixed; every commit maps to an iter (7: 1 tok + 5 tiks + 1 harden); ledger complete.

### Code Quality
- [x] [must-fix] F1: inline `\b`-terminal route regexes NOT migrated to the segment-anchor fix — skillpath-legacy.spec.ts:63 (`waitForURL(/…\/chapter\b/)`, load-bearing) + profile-page.ts:69 (`/\/profile\/skills\b/`). Centralize + reuse the anchored patterns.
- [x] [should-fix] F2/F3: dead code — `looksLikeTimelineEntry` (test-only) + 4 speculative self-eval accessors (`roleText`/`growthStat`/`rerateSkillsButton`/`editSkillButtons`, the routed self-eval UC → M206). Remove; give url-shapes a real page-object consumer.
- [x] [should-fix] F5: `/profile/skills` shape duplicated inline in 3 places — centralize in url-shapes.ts (resolved by F1's `SKILLS_TAB_URL`).
- [x] [nice-to-have] F4: `markCompleteButton()` lacks `.first()` (used at skillpath-legacy.spec.ts:70) — add the disambiguation guard.

### Documentation
- [x] DOC1: `playthroughs/README.md` lags M203 — name the 2 new manifest products + the 3 employee page-objects + url-shapes.ts; add an M203 line.
- [x] DOC2: soft M202-era present-tense framing in `corpus/ops/demo/playthroughs.md` ("the one starting surface" / "proof of life") — reconcile to reflect the shipped employee coverage.

### Tests & Benchmarks
- [x] [must-fix-adjacent] TEST-G1: no test asserts the `@pt:` tag-grammar LOCKSTEP between `cmd/ptvalidate/discover.go:20` + `report/playwright.go:18` (duplicated verbatim w/ "change both" comment, unenforced). Add a cross-package agreement test.
- [x] [nice-to-have] TEST-G3: `checkPreconditionCoverage` 96.2% — cover the empty-`seed.world` / blank-precondition arm.
- [x] [nice-to-have] TEST-G4: `discoverRegistry` 94.4% — cover the remaining file-filter branch.

### Decision Triage
- [x] D-CLOSE-1 (the 4 non-gate UCs Fate-3 → M206) → archive (maintainer routing record; the roadmap-vision M206 entry carries the user/future-dev-facing form).
- [x] TOK-01 (deterministic-read-first strategy) → archive (iter-loop strategy record; the shipped manifest + runbook carry the user-facing form).

## M203: Gate Outcome Ledger (Phase 9-iter — closed-on-gate)

**Gate.** Target: every declared employee-vantage use case has a PASSING Playthrough on a COLD reset-to-seed demo,
0 false-fails over 5 consecutive reset runs. Achieved: 6/6 employee Playthroughs pass; 5/5 cold reset runs, 0
false-fails (iter-06). Distance: **0. Status: MET.** → **closed-on-gate; no carry-forward.md (gate fired).**

**Iter ledger.** 6 iters, all closed-fixed: iter-01 tok(bootstrap, TOK-01 strategy) · iter-02 tik (Profile
verified+growth) · iter-03 tik (Profile timeline) · iter-04 tik (Skill Paths legacy) · iter-05 tik (AI Sims
chat launch + Sentinel-Reload) · iter-06 tik (5-run determinism gate MET). 1 tok + 5 tiks, no triggered tok. Every
commit maps to an iter (7 total incl. the final harden). No orphan iters/commits.

**Routes carried forward (Fate-3 → M206; none escape-hatch).** The 4 non-gate edge UCs the iter loop routed —
`ai-simulations.code.UC1`, `ai-simulations.interview.UC1`, the Skill-Paths verify-skill terminal, and
`profile.self-evaluation.UC1` — are additional to the gate (the gate enumerated the 3 core journeys, all GREEN).
All routed Fate-3 → **M206** (roadmap-vision.md annotated; D-CLOSE-1). Fate-1 infeasible at a docs-only close (each
needs a live demo + browser drive; code-sim needs a Judge0 host; self-eval needs live click-intercept iteration).

**Dropped.** None. (Academy skill-path UC = OUT by M201 design, not a defer; voice sims → M206 by design.)

**Protocol evolution.** None — the M202-delivered `corpus/ops/demo/playthroughs.md` protocol held across all 6
iters. The iter-05 post-seed Sentinel-Reload was folded into the runner (documented) + drift-guarded.

**Gate met: 6/6 ≥ target, 5/5 deterministic. No carry-forward; clean closed-on-gate close.**
