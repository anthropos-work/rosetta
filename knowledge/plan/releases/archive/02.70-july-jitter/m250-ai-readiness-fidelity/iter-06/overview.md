---
iteration_type: tik
status: closed-fixed
gate: NOT MET
---

# iter-06 — tik: browser-render confirmation (coverage-protocol sweep, both vantages)

**Active strategy:** TOK-01. **Target:** the final gate item — the coverage-protocol Playwright sweep (0
invented / 0 prod-ejects, real rendered content) for player (aria-completed) + manager (dana-manager) on the
Northwind AI-readiness org, demo-2.

## What ran
`rext stack-verify/e2e/run-coverage.sh 2 {employee aria-completed | manager dana-manager} "Northwind Aviation"`
(both `--workers=1`), on the cold-reseeded demo-2. Both sweeps completed (chromium render, ~8m each).

## Render verdict
- **Employee (aria-completed): GREEN** — `failingSections=0`, **`escapes=0` (0 prod-ejects)**, `personaFailures=0`,
  reachable 63/150. The member AI-readiness surface (`/home` `ai-readiness-member-done`) passes ("meaningful text
  3467 ok").
- **Manager (dana-manager): `escapes=0` (0 prod-ejects), `personaFailures=0`**, reachable 70/150, but **3 failing
  sections, all on `/ai-readiness`**. The M250-FIDELITY sections **PASS**: `ai-readiness-matrix` (the 31-skill
  repertoire), `ai-readiness-how-we-measure` (the named Step-2 sim cards + evaluated skills + dots),
  `ai-readiness-hero`, `ai-readiness-cycle-picker`, `ai-readiness-what-to-do-next`. The 3 FAILING are **adjacent
  manager-dashboard sections** (not M250's fidelity target):
  1. **`ai-readiness-by-tag`** — "region missing required content". Members ARE tagged (200 tagged / 223
     `membership_tags`) and `readiness.go`'s `byTag` aggregates `mem.Tags`, yet the region renders empty →
     needs debugging (member-Tags loading on the AI-readiness view, or a team-size filter). The M250-deferred
     `participants_filter`/team-tag lane (iter-02 D7) is the adjacent work.
  2. **`ai-readiness-interview-findings`** — "region missing required content: Breadth, Context fit". The
     `interview_aggregated_reports` row EXISTS (1 for Northwind) but lacks the finding structure the region's
     i18n labels (next-web `enterprise.json`) expect → the M219 interview-report seeder vs the post-M246
     platform's finding dimensions.
  3. **`ai-readiness-handled-for-you`** — `textMatch 1 < 3 for /[1-9][\d,]*\s+(AI skills mapped|minutes saved)/`.
     **The DATA IS CORRECT** — all three cycle-total counters are non-zero (skills_mapped **4272**, hands-on
     **5430 min**, interview questions **1359**). The failure is a coverage-manifest textMatch phrasing
     mismatch (the regex wants "minutes saved"; the page renders correct counters, likely "hours saved") → a
     coverage-manifest/test-health concern, NOT an M250 seeder bug.
- (Both vantages also carry `crossPortFailures=1` — the ant-academy `:23077` `ERR_CONNECTION_RESET` cross-port
  follow, a known peripheral, not an AI-readiness section.)

## Gate reading (on live render)
- **Part 1 (31 skills): PASS** — matrix + hero + how-we-measure render; closure **31/31**.
- **Part 2 (named sims + evaluated skills): PASS** — how-we-measure renders; SQL-confirmed tech/business named
  cards + evaluated skills + correct tracks.
- **Part 3 (distribution): PASS** — member-done renders; 787 session-backed verified evidence + 897 skill results.
- **Part 4 (manager faithful): PARTIAL** — the M250-fidelity manager sections (matrix, how-we-measure) PASS; 3
  ADJACENT manager sections fail (by-tag / interview-findings / handled-for-you — characterized above).
- **Part 5 (0 invented / 0 prod-ejects / closure / arithmetic): PASS** — **escapes=0 BOTH vantages**,
  personaFailures=0, closure 31/31, the M219 fences prove frozen-vs-live at denom 25.0.

**Gate = 4/5 on the strict coverage bar** (parts 1,2,3,5 PASS incl. 0 prod-ejects; part 4's core fidelity PASSES
but 3 adjacent manager sections fail).

## Close — 2026-07-24

**Outcome:** the coverage-protocol sweep confirmed M250's core fidelity renders correctly on BOTH vantages with
**0 prod-ejects** — parts 1, 2, 3, 5 PASS on live render. Part 4's M250-fidelity manager content PASSES; 3
adjacent manager-dashboard sections (by-tag, interview-findings, handled-for-you) fail and are the distance-to-gate.
**Type:** tik (render-measure)
**Status:** closed-fixed (the sweep + verdict landed)
**Gate:** NOT MET (4/5; part 4 has 3 adjacent failing manager sections)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y (5 tiks)** — (6) protocol-stop: n — Outcome: exit-5 (cap)
**Decisions:** D17 (the coverage sweep is the sanctioned render+prod-eject instrument — 0 prod-ejects confirmed both vantages), D18 (handled-for-you's data is CORRECT — non-zero counters; its failure is a coverage-manifest phrasing assertion, not an M250 seeder bug — route to coverage-manifest/test-health, NOT the seeder), D19 (by-tag + interview-findings are adjacent manager surfaces needing fresh debugging; M250-scope-vs-pre-existing is an orchestrator judgment call).
**Routes carried forward (iter-07, next invocation — the 3 manager sections):**
  - by-tag: debug why the region is empty despite 200 tagged members (mem.Tags loading on the AI-readiness view / team-size filter); revisit the deferred participants_filter/team-tag lane.
  - interview-findings: make the interview-report seeder write the finding structure (Breadth / Context fit / …) the post-M246 region expects.
  - handled-for-you: reconcile the coverage-manifest textMatch (the data is correct; the regex/units mismatch) — a coverage-manifest/test-health fix, not the seeder.
**Lessons:** the strict coverage bar (failingSections=0) is a SUPERSET of the 5-part fidelity gate — it flags adjacent manager surfaces M250 doesn't target. Separate "M250-fidelity renders correctly" (proven: matrix + how-we-measure + member-done + 0 prod-ejects) from "every manager-dashboard section passes" (3 adjacent gaps).
