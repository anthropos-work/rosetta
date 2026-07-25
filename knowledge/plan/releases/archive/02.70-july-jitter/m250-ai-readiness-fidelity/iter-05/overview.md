---
iteration_type: tik
status: closed-fixed
gate: NOT MET
---

# iter-05 — tik: evidence-distribution build (gate part 3 + 4)

**Active strategy:** TOK-01. **Target:** the last build lane — distribute each completed member's completed-sim
evaluated skills as VERIFIED (the validation fan-out + session-backed user_skill_evidences), which iter-04's
measurement pinpointed as the sole remaining gap (345 AI-sim sessions, only 5 with a fan-out).

## What landed
- **`seeders/ai_readiness_evidence.go`** (net-new):
  - `readAIReadinessSimEvaluatedNodeIDs` — reads a sim's evaluated node-ids from
    `directus.sequences.evaluation_skills` JOINed to `public.skills` (closure-safe: unresolved ids drop, never
    fabricated); empty when no per-stack Directus (fan-out skipped).
  - `appendSimValidation` — per completed member's ended Step-2/3 session: ONE `validation_attempt_results` +
    ONE `validation_attempt_skill_results` per evaluated node-id (the score `computePerformanceBySkill` reads
    for the dots) + ONE session-backed verified `user_skill_evidences` per node-id (the profile verified skill).
    Deterministic ids (keyed off the session) → re-seed no-op.
- **`ai_readiness_funnel.go`** — accumulators (`attemptResults`/`skillResults`), read the sim + interview eval
  node-ids once in `Seed`, thread `aiReadinessEvalNodes` through, call `appendSimValidation` for Step-2 (Tech
  sim) + Step-3 (interview), flush var+vasr in FK order, and pass the session id to the evidence upsert
  (session-backed verified, not a self-map).
- **Tests:** `ai_readiness_evidence_test.go` (4) + fixed the empty-org-guard signature. Full module GREEN.

## Measurement (re-seed demo-2, cold reset-to-seed)
- Seed GREEN (reset+seed exit 0; `ai-readiness-funnel rows=9071` [was 6965; +2106 = the fan-out]; audit
  "isolation: clean — no shared/external writes").
- **Gate part 3: PASS** — `validation_attempt_results` for AI sims **5 → 345**; `validation_attempt_skill_results`
  **897**; session-backed verified `user_skill_evidences` **787**.
- **Gate part 4: PASS** — the manager `computePerformanceBySkill` reproduction returns the evaluated skills with
  real avg scores + dot ratings: Ai Fundamentals 74 (n159), Prompt Engineering 74 (n181), Ai Coding 73,
  Critical Thinking Fundamentals 73, Generative AI Fundamentals 73 (n181 each).

## Gate reading now (data-level)
Parts **1, 2, 3, 4 PASS**; part 5 largely PASS (closure 31/31, fences prove frozen-vs-live). **The only
remaining item is the BROWSER-RENDER confirmation** (0 invented values + 0 prod-ejects proven by rendering the
page per coverage-protocol.md) → iter-06.

## Close — 2026-07-24

**Outcome:** the evidence-distribution lane landed — completed members now carry their completed-sim evaluated
skills as verified (validation fan-out + session-backed evidence), flipping gate parts 3 AND 4 to PASS at the
data level. All 3 build lanes are now complete.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (parts 1-4 PASS at data level; part 5 needs the browser-render 0-invented/0-prod-eject confirmation)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — Outcome: continue (→ iter-06, browser-render confirmation)
**Decisions:** D14 (extend the FUNNEL, not a separate seeder — the funnel holds the session ids in-memory; Conn has no multi-row Query to read them back), D15 (read eval node-ids once from directus, closure-safe via the public.skills JOIN — same truthful source as the set-dress), D16 (reuse the existing evidence upsert with a session id for the verified sim-signal; overlapping config∩eval skills [K-CRITHI] converge cleanly).
**Routes carried forward:** iter-06 = browser-render confirmation of all 5 parts (player + manager, 0 invented, 0 prod-ejects) per coverage-protocol.md. Still deferred: participants_filter track-tagging + business-sim per-member session routing (both non-blocking for the gate as measured; revisit if the render shows a gap).
**Lessons:** measure-first (iter-04) made this build surgical — one precise gap (the fan-out), one file, and the re-measure flipped two gate parts at once (part 3 profile + part 4 dots share the validation_attempt_skill_results source).
