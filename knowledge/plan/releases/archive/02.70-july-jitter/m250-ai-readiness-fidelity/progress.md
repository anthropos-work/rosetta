# M250 — Progress

## Running ledger

_(Per-iter progress — tik/tok entries, distance-to-gate, and gate-part (1–5) evidence — accumulates here during
the iter loop. `iter-NN/` dirs are created by `/developer-kit:build-mstone-iters` on its first invocation; there
are NO iter dirs at scaffold.)_

- iter-01 (tok/bootstrap): authored TOK-01 (arithmetic-spine → set-dress → distribute → render loop) — gate 0/5 baseline — see iter-01/progress.md
- iter-02 (tik): Lane A — arithmetic-spine atomic edit landed (config 31 defaults + 3 track-keyed sims; funnel pool + pinned refs; started-hero 3→9; ALL fences re-derived green at 25.0) — full stack-seeding module GREEN — gate at live render still 0/5 (render deferred to iter-05) — see iter-02/progress.md

- iter-03 (tik): Lane B — net-new Directus set-dress seeder (`AIReadinessSimSkillsSeeder`) landed + registered + 4 tests green + LIVE-validated on demo-2 (3 sims resolve real evaluated-skill names; track heuristic → tech/business correct) — gate at full render still 0/5 — see iter-03/progress.md

- iter-04 (tik, measure-first): first live reset-to-seed of demo-2 with Lanes A+B — parts 1+2 PASS (data), part 5 largely green (closure 31/31), part 3 FAIL (345 AI-sim sessions, only 5 with validation fan-out → the evidence-distribution gap), part 4 partial — see iter-04/progress.md

- iter-05 (tik): evidence-distribution build LANDED — new ai_readiness_evidence.go (validation_attempt_results + skill_results + session-backed verified user_skill_evidences for completed members' evaluated node-ids, closure-safe read from directus); re-seed demo-2 → gate part 3 PASS (var 5→345, vasr 897, session-backed verified evidence 787) + part 4 PASS (manager dots render, avg 73-74); all 3 build lanes complete — gate ~4.5/5 (browser-render 0-invented/0-prod-eject confirmation remains) — see iter-05/progress.md

- iter-06 (tik, render-measure): coverage-protocol Playwright sweep BOTH vantages on demo-2 — employee aria-completed GREEN (failingSections=0, escapes=0); manager dana-manager escapes=0/personaFailures=0 but 3 failing sections all on /ai-readiness (by-tag, interview-findings, handled-for-you). M250-fidelity manager sections PASS (matrix=31 skills, how-we-measure=named sims+evaluated skills+dots). Gate 4/5: parts 1,2,3,5 PASS on render (0 prod-ejects, closure 31/31); part 4 = 3 adjacent manager sections — see iter-06/progress.md
- iter-07 (tik, drift reconcile): the 3 adjacent manager sections re-surveyed at source → all post-M246 platform DRIFT, not data gaps — 2 coverage-manifest phrasing reconciles (by-tag "…by Team", handled-for-you "Hours saved") + 1 seeder KPI-id (interview-findings avg_adoption/…/avg_ownership). Fixed + committed + data-confirmed + unit-green (rext @ july-jitter-m250-iter07 584f1fe). LIVE manager-sweep confirmation (~150-page crawl, times out locally) → M254 (CARRY-M250-01, Fate 2). Milestone closes on the user pragmatic-close mandate — see iter-07/overview.md

## Gate Outcome Ledger

### Gate
- **Target:** on a cold reset-to-seed, for a completed Northwind AI-readiness member — (1) step-1 renders the 31 default readiness skills (19 core + 12 enabling); (2) step-2 shows the track-keyed named sim (tech=who-can-see-this-document-fc0 / business=use-ai-to-turn-survey-data-into-a-leadership-email) + interview with a non-empty evaluated-skills list of real node-ids; (3) the profile carries the completed sim's distributed verified skills; (4) the manager view shows the same faithfully; (5) 0 invented values, 0 prod-ejects, closure green, frozen-vs-live arithmetic agrees at 31.
- **Achieved:** parts **1/2/3/5 LIVE-GREEN both vantages** (employee `aria-completed` + manager `dana-manager`, Northwind, demo-2, cold reset-to-seed, escapes=0): 31 real defaults (closure 31/31) · 2 track-keyed named sims + non-empty evaluated-skills · verified-skill distribution to completed members · 0 prod-ejects + arithmetic re-derived green (25.0). Part **4 core fidelity sections LIVE-GREEN** (31-skill matrix + how-we-measure named-sims/evaluated-skills/dots). Part 4's **3 ADJACENT** manager-dashboard sections (`by-tag`/`interview-findings`/`handled-for-you`) drift-fixed + data-confirmed + unit-green; **live** sweep → M254.
- **Distance:** core gate **MET + live-confirmed**; residual = the live confirmation sweep of the 3 adjacent drift-fix sections (not a fidelity gap).
- **Status:** `closed-incomplete` (user pragmatic-close mandate, 2026-07-24).

### Iter ledger summary
- **Total iters:** 7 (toks: 1 [iter-01 bootstrap TOK-01], tiks: 6 [iter-02…07]).
- **Duration:** 2026-07-24 → 2026-07-24 (single day; the marquee's serial render loop).
- **Decisions accumulated:** D1–D18 (iter-level D1–D16 in iter progress; D17/D18 at close). 1 tok strategy (TOK-01, never revised — no triggered tok, no re-scope).
- **Hardening passes embedded:** the arithmetic fences (m219/harden re-derived at 25.0) + closure gate ran at every build lane; unit-green at the code-of-record tag (seeders + dna). No separate `harden-mstone-iters --final` pass — the code-of-record is a separate gitignored rext repo, hardened + unit-green at iter-07; the pragmatic-close mandate governs.

### Routes carried forward — three-fate dispositions

#### Fate 2 — Already planned in another release-milestone
- **CARRY-M250-01** (live manager-coverage-sweep confirmation of the 3 adjacent drift-fix sections) — **Owned by:** M254 → knowledge/plan/releases/02.70-july-jitter/m254-*/overview.md:85 exit gate (d) + :89 (h). See carry-forward.md.

#### Fate 3 — Annotated/attached to a milestone of this release
- None.

#### Escape-hatch — release-scope-breaking deferral
- None.

### Dropped
- **DEF-M250-01** — participants_filter track-tagging + per-member business-sim session routing — **Reason:** non-gate believability nicety; its only gate-relevant suspicion (empty by-tag) was falsified at iter-07 as a manifest copy drift; track label is driven by the landed name-heuristic set-dress — **Decision:** D18.

### Protocol evolution
- The `coverage-protocol.md` + `verification.md` measure→triage→fix→re-render loop was applied to a new surface (AI-readiness manager+employee). iter-04's **measure-first re-survey** (read the platform's own SQL read-paths as a faithful render proxy before spending a live reset-to-seed) and iter-07's **re-survey-at-source** (drift-vs-data-gap triage) are the reusable refinements; no protocol-doc edit needed (both are legitimate protocol moves already sanctioned).

## M250: Final Review

### Scope
- [x] Gate-distance + iter-ledger audit (Phase 1) — 7 iters accounted; core gate met; 1 Fate-2 carry-forward.
- [x] Deferral re-audit (Phase 1b) — GREEN.

### Documentation (Delivers)
- [x] corpus/services/ai-readiness.md — seeding contract updated to the 31-skill default + set-dress + evidence distribution + re-derived arithmetic + interview-findings KPI ids.
- [x] corpus/ops/seeding-spec.md — v2.7 M250 changelog paragraph.

### Decision Triage
- [x] D1/D2/D4/D14-16/D17 → blended into ai-readiness.md (§ Seeding contract 31-skill fidelity + FILLED-ness items 1/3/4/5).
- [x] D5/D6 → blended into ai-readiness.md (FILLED-ness arithmetic).
- [x] D11/D12/D13 → archived (maintainer-only iter-loop mechanics).

### Deliverables (iterative)
- [x] Gate Outcome Ledger (this section).
- [x] carry-forward.md (CARRY-M250-01 → M254, Fate 2).
- [x] metrics.json.

## Next-iter queue
_(Closed — no further iters. The single carry-forward route CARRY-M250-01 is owned by M254; see carry-forward.md.)_
- iter-05 (tik): evidence-distribution build — validation fan-out (validation_attempt_results + skill_results + verified user_skill_evidences) for completed members' step-2 sim evaluated node-ids (reuse content_stories_write helpers); flips part 3 + part 4 dots.
- iter-07 (tik): the 3 adjacent manager sections — (a) by-tag empty-despite-tags debug, (b) interview-findings finding-structure seed, (c) handled-for-you coverage-manifest phrasing reconcile (data is correct).
- iter-05+ (tik): live reset-to-seed render loop on demo-2, measure 5 gate parts.
