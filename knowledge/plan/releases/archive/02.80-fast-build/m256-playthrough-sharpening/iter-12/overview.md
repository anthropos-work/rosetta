---
milestone: M256
iter: 12
iteration_type: tik
status: closed-fixed
created: 2026-07-28
---

# M256 · iter-12 — negative controls, the cheap tail first

**Type:** tik · **Active strategy:** `TOK-01` move 4 · **Handler:** `NEGCTL-M256-cross-vantage`.

## Step 0 — re-survey

- The fence now *reports* the gap on every run: **8 of 24**, with the 16 uncovered ids named. Re-read live,
  not from iter-11's prose: `pt-activity-drilldown`, `pt-aireadiness-manager-dashboard`,
  `pt-aireadiness-manager-howwemeasure`, `pt-aireadiness-member-done`, `pt-aireadiness-member-progress`,
  `pt-hiring-recruiter-compare`, `pt-profile-growth`, `pt-profile-identity`, `pt-profile-timeline`,
  `pt-profile-verified`, `pt-studio-advanced-generate`, `pt-studio-guided-generate`, `pt-workforce-funnel`,
  `pt-workforce-org-feedback`, `pt-workforce-roster`, `pt-workforce-succession`.
- iter-11 left the mechanism proven on one pair, plus a triage rule: **the cheap cases are the ones where the
  two vantages differ by SEEDED STATE.**

**Reading the 16 against that rule changes the shape of the work** — and it surfaced a constraint iter-11's
routing had not stated. Grouped by what their FINAL actually asserts:

| group | what the final asserts | is there a contrast vantage in the seed? |
|---|---|---|
| 4 × AI-readiness | readiness-specific surfaces (`doneHeroTitle`, `progressFunnelTitle`, `dashboardHeading`, `methodHeading`, …) | **YES** — the readiness surfaces exist only for an org with `narrative: ai-readiness`; Orgs A/B are not, and the docs say every readiness surface then answers `ErrAIReadinessDisabled` |
| 1 × hiring recruiter | the apps/hiring Results scoreboard | **YES** — gated on `is_hiring`; a Workforce-org manager is not routed there |
| 1 × `pt-profile-identity` | `heroName('Pat Ellis')` — hero-SPECIFIC | **YES** — any other hero's profile |
| 4 × workforce + 1 × activity-drilldown | **structural** stats (`mappedSkillsStat`, `coverageGap`, `memberRows().first()`, …) | **NO** — those render for *any* populated org, so a different-org manager does not discriminate |
| 3 × profile (verified/growth/timeline) | **structural** stats (`verifiedSkillsStat`, `skillCharts`, `workSection`) | **probe** — `pt-manager` is seeded with no `skills:` block, so she may have 0 verified |
| 2 × studio | blocked behind `FIX-M256-studio-false-green` | **NO — and asserting a control on a known false green would certify it.** Route. |

## Cluster / target identified

The **6 high-confidence seed-state cases** (4 readiness + hiring + profile-identity), plus a probe of the
3 structural profile cases. The 5 structural workforce cases and the 2 studio cases are **routed with a
written reason**, not silently skipped.

## Hypothesis

**H1 — a contrast vantage discriminates when the asserted outcome is ORG- or HERO-specific, and cannot when
it is structural.** A `mappedSkillsStat` renders for every populated org; a `doneHeroTitle` renders only where
readiness is enabled. So the cross-vantage mechanism is not uniformly applicable, and pretending otherwise
would produce 10 controls that pass for any org — the same vacuity iter-07 refuted in the ablation.

**H2 — the control must NOT be paid inside the timed Playthrough.** A second login + navigation inside a
Playthrough would roughly double its duration; the median is 2.282 s against a 2.628 s gate, so 16 in-test
controls would break clause 1 outright. The overview's own note says negative-control runs are *"excluded from
the timed p50 either way"*. So the controls live in their own spec file with **no `@pt:` id** — they are not
Playthroughs, they do not enter the median, and each Playthrough's `@pt-negative-control:` line names the
control that covers it.

**H3 — batching by vantage makes them nearly free.** One login per contrast seat, several absences asserted per
login: ~2 logins covers the 6.

## Expected lift

Negative controls **8 → 14 of 24** (fence-computed), 0 new Playthroughs, clause 1 unchanged within variance
(the controls are outside the Playthrough set). A measured verdict on the 3 structural profile cases.

## Phase plan

- **Phase A — probe** the contrast vantages live: does a non-readiness hero's `/home` really lack the readiness
  surfaces (and stay ALIVE, not dead)? Is a Workforce manager really not served the hiring Results view? Does
  `pt-manager` really have 0 verified skills?
- **Phase B — build** the control spec for whatever discriminated; route the rest with the measurement attached.
- **Phase C — re-measure** (full suite ×3, fence counts, clause 1).
- **Phase D — close.**

## Escalation conditions

- A contrast vantage yields a **dead page** rather than a live-but-empty surface → it cannot discriminate
  (iter-07 D29's refutation); record and route, never ship it.
- More than 3 of the 16 prove to need a platform edit → that is the milestone's re-scope trigger, escalate.

## Acceptable close-no-lift outcomes

A measured finding that the cross-vantage mechanism covers only the org/hero-specific subset, with the
structural subset priced honestly (sharpen each final to name real seeded data — O(tests)). That is a real
deliverable: it converts "16 to go" into two named classes with different costs.
