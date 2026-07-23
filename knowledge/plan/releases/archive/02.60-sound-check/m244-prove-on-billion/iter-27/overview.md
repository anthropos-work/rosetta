---
iter: 27
milestone: M244
iteration_type: tik
status: closed-fixed
created: 2026-07-23
---

# iter-27 — land the 3 remaining ai-readiness sub-renders (gate c 13/16 → 16/16 = GATE MET)

**Type:** tik (run 10, tik 1)
**Active strategy reference:** TOK-03 (HOLD TOK-01/02 — the final push; gate (c) is the last part, and its
16 Playthroughs are seed-order-forced LAST). This tik completes the gate-(c) Playthrough half by landing the
3 remaining ai-readiness sub-renders (handler FIND-M244-aireadiness-subrenders, routed iter-26).

## Step 0 — re-survey
gate (c) = 13/16 Playthroughs (12 non-aireadiness + ai-readiness member-done green since iter-25). The 3
remaining ai-readiness specs fail on distinct deterministic assertions. **iter-26 D1 characterized all 3 as
SEED gaps; this re-survey FALSIFIES that** against billion's live DB + v1.341.0 read paths + the actual UI:
all three are **HARNESS locator mismatches** vs the correct, fully-seeded, correctly-rendering v1.341.0 UI —
NOT seed gaps, NOT platform gaps.

- **(a) manager-dashboard.UC1 byTeam** — Org C HAS 13 tags + 42 membership_tags; 12 teams of started+tagged
  members; backend ByTeam is rich; UI renders the section titled **"AI Readiness by Team"** (translation
  `aiReadiness.byTag.title`). The harness locator asserts `/AI Readiness by Tag/i` — "Tag" ≠ "Team".
- **(b) manager-dashboard.UC2 interview panel 24 chars** — the seeded `interview_aggregated_reports` row is
  rich (patterns=4/unexpected=4/insights=4/kpis=4/charts=2, session_count 31); `queryInterviewAggregatedReports`
  is keyed by **organization_id ONLY** (NOT sim_ref — the v1.341.0 comment says so explicitly, refuting
  iter-26's sim_ref hypothesis); the 4 findings headings all RENDER (the spec passes lines 82–87). The failing
  `>900` check measures `interviewBreakdown().innerText()`, which scopes to the **24-char heading span**
  ("AI Interview — breakdown"), not the panel CARD that holds the findings (siblings of the heading inside
  the `openStep===3` card div).
- **(c) member-funnel.UC2 deadline** — the active cycle end_date is present; `DueDate.tsx` renders
  **"Due Aug 22 · N days left"** — `toLocaleDateString(locale,{month:'short',day:'numeric'})` (NO year). The
  M219-hardened `dueDate()` locator requires a **4-digit year** in every date alternative → never matches.

## Cluster / target identified
FIND-M244-aireadiness-subrenders — the 3 remaining gate-(c) Playthroughs. TOK-03 named "3 seed-fix cycles";
re-survey shows they are **one class** (harness locators vs v1.341.0 UI) in ONE file
(`playthroughs/e2e/lib/ai-readiness-page.ts` + one spec). Substituting the 3 seed cycles with a single
harness-fix tik under the same TOK-03 gate-(c) push (0 seed changes, 0 platform edits, no billion re-pin/re-seed
— the Playthrough harness runs from the LOCAL authoring copy against billion).

## Hypothesis
Correcting the 3 harness locators to match the actual v1.341.0 UI lands all 3 specs → gate c 16/16 → the
binary gate-parts metric flips 7/8 → **8/8 GATE MET**.

## Phase plan
Fix the 3 locators (a: "Team"; b: add `interviewBreakdownPanel()` card-scope + point the >900 check at it;
c: allow a year-less short-date alternative while KEEPING the due/deadline anchor) → `tsc` + unit gate →
re-run the 4 ai-readiness Playthroughs FOREGROUND from the peer harness against billion → confirm all land →
re-run the FULL 16-Playthrough suite to confirm 16/16 (no regression) → gate c ticks → GATE MET.

## Escalation conditions
If any sub-render turns out to need a PLATFORM edit (the data genuinely isn't rendering) → STOP, SEVERITY=blocker.
If a locator fix makes an assertion toothless (passes without proving the real surface) → do not ship it.

## Acceptable close-no-lift outcomes
If re-run reveals a sub-render is a genuine data/platform gap (not a locator issue) → characterize + route
honestly with the precise falsification, gate c stays 13/16.
