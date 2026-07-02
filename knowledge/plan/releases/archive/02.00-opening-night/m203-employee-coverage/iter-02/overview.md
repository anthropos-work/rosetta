---
iteration: 02
iteration_type: tik
milestone: M203
status: closed-fixed
created: 2026-07-02
---

# iter-02 — TIK (Profile read journeys)

## Active strategy reference
**TOK-01** (Deterministic-read-first, then mutating flows, then integration boundary). This tik executes
TOK-01's step 1 (Profile), starting with the read surfaces.

## Cluster / target identified
Prior close (iter-01, the bootstrap) routed forward: "Profile journeys (verified / growth / self-eval),
starting with `profile-skills.verified.UC1`." Re-survey (Phase 1 Step 0): baseline still 1 UC declared+passing;
a DB probe confirmed pt-employee (Pat Ellis) seeds **24 verified / 36 total** user_skills (≫ the ≥2 Spotlight
datapoint threshold) — so the verified-skill Spotlight WILL render. A live /profile probe found the surface:
tabs Career Profile / Skills / Activities; the Skills tab (/profile/skills) renders "Verified Skills 8",
"Role Skills 10", "Skill Gaps (5)", "Growth", and 46 SVG charts. Target stands: land the two DETERMINISTIC
READ journeys (verified + growth) this iter; route the self-evaluation WRITE flow to iter-03 (its "Re-rate
skills" entry proved finicky under probe — needs its own iter to anchor + assert at the right boundary).

## Hypothesis
The verified-skill + growth surfaces render real data for a hero with a seeded verified-skill history, so a
Playthrough that opens /profile → Skills tab and asserts the verified/claimed/gap stats + a drawn chart is
present will PASS deterministically (pure read, no mutation → no reset needed).

## Expected lift
+2 employee use cases passing (profile.verified.UC1, profile.growth.UC1), no regression on the identity smoke.

## Phase plan
Protocol (`playthroughs.md` §"iteration protocol"): declare the 2 UCs in the manifest (keep ptvalidate green) →
extend ProfilePage with Skills-tab semantic accessors (no seed change — the pt-world seed already provides the
verified-skill capability) → build the 2 specs → run against demo-1 → reconcile with ptreport (no-regressions).

## Escalation conditions
A surface undrivable without a platform edit → `unimplementable-without-platform-edit` (escalate, don't edit).
A verified-skill surface that renders EMPTY despite the seeded datapoints → diagnose seed-vs-platform drift (P6)
before concluding a capability regression.

## Acceptable close-no-lift outcomes
If the verified/growth surface proved un-assertable without pixel/copy coupling (violating P2), that
falsification (recorded) would still be a complete iter — but the probe already found sanctioned semantic anchors.

## Close → see progress.md
