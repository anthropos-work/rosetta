---
title: "KB Fidelity Audit — M227 the-notes"
date: 2026-07-17
scope: milestone:M227
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Hiring read-model / render-path / gate | `corpus/services/hiring.md` | `app/internal/organization/intelligence.go` (`InsightsByJobSimulations`), `stack-seeding/seeders/hiring_funnel.go` + `hiring_config.go` | PAIRED |
| The seeded funnel shape ("MOST on all 5") | `corpus/ops/seeding-spec.md`, `corpus/ops/demo/stories-spec.md` | `stack-seeding/seeders/hiring_funnel.go` | PAIRED |
| Email domain derivation | `corpus/ops/seeding-spec.md` (G6 org-domain) | `stack-seeding/seeders/userprofile.go` (`emailFor`/`storyEmailDomainFor`), `users.go` | PAIRED |
| Avatar (synthetic-face, deterministic) | `corpus/ops/demo/profile-completeness-spec.md` / coverage-protocol persona self-consistency | `stack-seeding/seeders/avatar.go` + `assets/avatars.go` | PAIRED |
| Coverage sweep (hiring vantage) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/tests/render-hiring-comparison.spec.ts`, `run-hiring-render.sh` | PAIRED |
| Hiring playthrough | `corpus/ops/demo/playthroughs.md` | `playthroughs/e2e/tests/hiring-recruiter.spec.ts`, `manifest/hiring.yaml` | PAIRED |

## Fidelity Findings
1. **hiring.md read-path (mirror table)** — Source `hiring.md` §The comparison read-model. Expected: the scoreboard/list reads `public.local_jobsimulation_sessions`. Actual: `intelligence.go:1472 InsightsByJobSimulations` queries `LocalJobsimulationSession` filtered by org members + org id. **ALIGNED.**
2. **"MOST on all 5 / SOME assigned-only"** — Sources `hiring.md:191`, `seeding-spec.md:361,460`, `stories-spec.md:671`. Actual: `hiring_funnel.go:seedHiringOrgFunnel` loops `for pi, simID := range positions` (all 5) for assessed candidates; ~10% assigned-only. **ALIGNED with current code.** These are exactly the claims M227 fix #3 revises (each candidate → 1 position) — planned revision, not stale-in-error.
3. **Render-probe gate floor `≥40`** — Source `render-hiring-comparison.spec.ts:53` (`RENDER_GATE_FLOOR ?? '40'`), `run-hiring-render.sh:17`, M224 `decisions.md` GATE-DECISION D1, M226 `overview.md` exit_gate (2), M228 `roadmap.md` exit gate (2). Actual: default 40. **ALIGNED with current data (~40/position).** M227 fix #3 retunes to the realistic floor (~6-8) everywhere — planned revision.
4. **Avatar is gender-blind** — Source `avatar.go` doc comment + `profile-completeness-spec`. Actual: `photoAvatarDataURI(seed)` picks by `hash(seed)` over all 12 faces, no gender. **ALIGNED** (the doc doesn't claim gender-consistency). M227 fix #4 adds it (new behavior, documented in Phase 5).
5. **Email is org-domain for all members** — Source `seeding-spec.md` (G6). Actual: `users.go` `domain := storyEmailDomainFor(st)`, role-blind. **ALIGNED.** M227 fix #2 splits candidate→external (new behavior, documented in Phase 5).

## Completeness Gaps
- (incidental, non-blocking) `coverage-protocol.md` persona self-consistency lists "menu==profile real-photo avatar" but not gender-consistency; fix #4 may add a gender-consistency coverage assertion. This is NEW behavior M227 introduces, documented in Phase 5 — not a pre-existing KB gap. No tracking item needed beyond the milestone's `Delivers →` line.

## Applied Fixes
None required — no stale-in-error claims or broken cross-refs found. The `≥40 / all 5 / gender-blind` descriptions correctly document today's behavior; M227 revises them as part of its deliverables (`overview.md` `Delivers →` covers the hiring.md + specs updates).

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. Every topic PAIRED, every audited claim ALIGNED with current code, no blind areas, no stale load-bearing claims. The docs M227 will revise (funnel shape, gate floor) accurately describe present behavior and are enumerated in spec-notes for the fix #3 retune.
