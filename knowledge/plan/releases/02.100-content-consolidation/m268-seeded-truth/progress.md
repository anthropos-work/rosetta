# M268 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in
[`overview.md`](overview.md); technical notes in [`spec-notes.md`](spec-notes.md).

**Status: PLANNED — not started.**

## Section checklist

- [ ] **(1) `completionForScore(score)` extracted** — one derivation of verdict from score at threshold
  **60**, routing `jobsim_sessions.go`, `feedback.go` and `hiring_funnel.go` through it. The pattern
  already exists at `hiring_funnel.go:346-350`; `hiring_funnel.go:74` declares
  `hiringPassThreshold = 60`, and seeded `success_threshold` is `60.0` at `persona.go:303` /
  `content_stories_write.go:310`.
- [ ] **(2) The two-hash split killed** (`jobsim_sessions.go:142-157`) — **both** directions:
  - [ ] (a) a **FAILED** session is clamped down (today `passed == false` leaves the score untouched,
        so a failed row can carry a **95**)
  - [ ] (b) the **[55,60)** hole closed (today the nudge floor is **55**, so a "passing" **57** stays 57)
- [ ] **(3) Duplicate session id resolved** — `jobsim_sessions.go:136` and `feedback.go:153` are
  byte-identical for `j == 0`; `CopyRowsIdempotent(..., "id")` is `ON CONFLICT DO NOTHING` and
  `cmd/stackseed/main.go:525` registers the **inconsistent** writer before `:540`, so it wins and
  `feedback.go:196-199`'s score-consistent rule is silently discarded.
- [ ] **(4) Remaining session writers CHECKED, not assumed**
  - [ ] `persona.go:286` — unconditional `passed`; score needs a `>= 60` floor
  - [ ] `ai_readiness_funnel.go:586` — unconditional `passed`; score needs a `>= 60` floor
  - [ ] `content_stories_write.go:57-59` — takes **both** from the source fixture; disposition recorded
        as a decision (a real session that disagrees is truth, not a bug)
- [ ] **(5) The fence** — a data-DNA gene / fence test asserting score↔verdict agreement over **EVERY**
  seeded session row. *The bug survived because nothing asserted this.*
- [ ] **(6) C1 — Programs seeded at 3–5 plans per org** across completed / just-started / in-progress /
  almost-done, with per-plan progress consistent with member counts. Extends the existing seeder
  (`assignment_plans.go:56-88`, `:180-215`; `assignments.go:174-179`, `:210-224`), which today writes
  exactly **one `status=active` plan per org**. Surface: `/enterprise/assignments-list`.
- [ ] **(7) Docs** — `corpus/ops/seeding-spec.md` (the invariant as a documented gene) +
  `corpus/ops/demo/stories-spec.md` (Programs as a seeded story surface).

## Pre-flight

- [ ] `pt-assignment-assign` / `pt-assignments-nav-v2` still green before any change
- [ ] seed-manifest honesty gate (`CanonicalFileMatchesProjection`) checked against the new plan counts

## Verification

- [ ] Live on a demo stack: a rendered session's score and its red/green verdict agree, both directions
- [ ] Live on a demo stack: the Programs section is non-empty and shows distinct stages
- [ ] Full `stack-seeding` suite green; rext tagged **and pushed to origin**
