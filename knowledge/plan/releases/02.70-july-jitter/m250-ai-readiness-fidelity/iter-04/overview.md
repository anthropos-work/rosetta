---
iteration_type: tik
status: closed-fixed
gate: NOT MET
---

# iter-04 — tik: first live reset-to-seed render + MEASURE (measure-first re-survey substitution)

**Active strategy:** TOK-01. **Step 0 re-survey (substitution):** TOK-01 named iter-04 = evidence
distribution, but the protocol is measure→triage→fix. Lanes A+B are landed + SQL-validated but never rendered.
Substituting iter-04 = the first live reset-to-seed MEASURE, to scope the evidence-distribution fix precisely
rather than build it blind. Same TOK strategy; only the named next-target moved.

## What ran
- Built authoring `stackseed`; **reset-to-seed demo-2** (offset 20000, port 25432) with the new code:
  `stackseed --stack demo-2 --reset` then `--seed presets/stories.seed.yaml`. Both exit 0, no errors.
- All seeders GREEN incl. `ai-readiness-config rows=39` (31 skills + 3 sims + 3 steps + 2 cycles),
  **`ai-readiness-sim-skills rows=3`** (my new seeder), `ai-readiness-funnel rows=6965`.

## Gate reading (data-level; browser render is the final confirmation, iter-06)
- **Part 1 (step-1 renders 31): PASS** — 31 config skills for Northwind (was 8); closure **31/31** resolve.
- **Part 2 (step-2 track-keyed named sims + non-empty evaluated skills): PASS** — reproduced the platform's
  `computeSimAssessments` query: simulation/**tech** "Software Engineer Test…" `["Ai Coding", …]`,
  simulation/**business** "Use AI to Turn Survey Data…" `["Critical Thinking…", "Generative AI…", "Prompt
  Engineering"]`, interview/both "AI Readiness Interview" `["Ai Fundamentals"]`. Correct named sims, correct
  tracks, non-empty evaluated-skills lists.
- **Part 3 (completed member's distributed verified skills): FAIL** — **345** sessions against the 3 AI sims but
  only **5** carry a `validation_attempt_results` fan-out. The funnel seeds the step-2/3 SESSIONS but NOT the
  `validation_attempt_results` + `validation_attempt_skill_results` + verified `user_skill_evidences` for the
  sim's evaluated node-ids (tech = K-AICODX-7A9E / K-CRITHI-224F / K-GENAIF-FE2F / K-PROENG-4AFF). **This is the
  remaining build.**
- **Part 4 (manager faithful): PARTIAL** — the manager Step-2 cards use the SAME `computeSimAssessments` (PASS);
  the per-skill dot ratings (`skillRatingsByName`) depend on part 3's distribution.
- **Part 5 (0 invented / closure / frozen-vs-live): largely PASS** — closure **31/31** (real defaults, 0
  fabrication); believable frozen Step-1 spread (1 → 30, ~200 members) at denom 25.0. 0-prod-ejects needs the
  browser render.

## Distance-to-gate
**~2 PASS (1,2) + part 5 largely green; parts 3 (FAIL, scoped) + 4 (partial) gated on the evidence-distribution
fan-out.** The measurement precisely scoped the last build: distribute each completed member's step-2 sim
evaluated node-ids as verified (validation fan-out + user_skill_evidences), which also fills part 4's dots.

## Close — 2026-07-24

**Outcome:** first live reset-to-seed proved Lanes A+B live (31 skills / 3 track-keyed named sims / evaluated
skills all render at the data level; closure 31/31) and produced a definitive 5-part gate reading that scoped
the remaining evidence-distribution build (gate part 3).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (part 3 FAILS; parts 1/2/4/5 need the browser-render final confirmation)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — Outcome: continue (→ iter-05, evidence distribution)
**Decisions:** D11 (measure-first re-survey substitution — legitimate under the protocol, not a re-scope), D12 (drive the reset-to-seed with the authoring stackseed against demo-2's offset DSN — a LOCAL demo consumes MY code without the tag+push dance; billion untouched), D13 (measured via the platform's own SQL read paths — computeSimAssessments etc. — a faithful proxy for the render; browser confirmation deferred to iter-06).
**Routes carried forward:** iter-05 = evidence-distribution build (validation fan-out for completed members' step-2 sim evaluated node-ids, reuse content_stories_write helpers) → flips part 3 + part 4 dots; iter-06 = re-seed + browser-render confirmation of all 5 parts (0 invented, 0 prod-ejects). Still deferred: participants_filter track-tagging + business-sim per-member session routing.
**Lessons:** measure-first pays — one live reset-to-seed converted "3 unbuilt lanes" into a precise reading (2 parts already PASS at the data level; exactly ONE build remains). The "opposite pins" note is fully resolved live: the track LABEL is the name heuristic over simulations.skills, and it lands tech/business correctly.
