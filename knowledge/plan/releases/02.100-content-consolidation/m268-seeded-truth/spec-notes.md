# M268 — Spec notes

Technical notes accumulate here during build. The authoritative scope lives in
[`overview.md`](overview.md); the release narrative in
[`../../../roadmap.md`](../../../roadmap.md) § v2.10.

M268 is **SEED-DATA ONLY — zero platform-repo edits** (a platform-source wall routes to a sha-pinned
`demopatch`, never a repo edit). Everything below is in `rosetta-extensions/stack-seeding`.

---

## Measured ground (design-time, carried verbatim — do not paraphrase away)

Every line below was measured. Re-verify against the authoring copy before editing, but do not
re-derive from documentation.

| Anchor | What it says |
|---|---|
| `seeders/jobsim_sessions.go:142-157` | score and `passed` from **two different, independent hashes** |
| `seeders/jobsim_sessions.go:136` | `key := fmt.Sprintf("%s:session:%d:%d", prefix, i, j)` |
| `seeders/feedback.go:153` | `sessKey := fmt.Sprintf("%s:session:%d:%d", prefix, i, 0)` — **byte-identical for `j==0`** |
| `seeders/feedback.go:196-199` | the score-consistent rule (`score >= 55`) — the writer that **loses** |
| `cmd/stackseed/main.go:525` / `:540` | `JobsimSessionsSeeder` registered **before** `FeedbackSeeder` |
| `seeders/hiring_funnel.go:346-350` | **the correct pattern** — completion derived FROM the score |
| `seeders/hiring_funnel.go:74` | `hiringPassThreshold = 60` |
| `seeders/persona.go:303` | seeded `success_threshold` = `60.0` |
| `seeders/content_stories_write.go:310` | seeded `success_threshold` = `60.0` |
| `seeders/persona.go:286` | unconditional `sessCompletionPassed` — **CHECK, do not assume** |
| `seeders/ai_readiness_funnel.go:586` | unconditional `sessCompletionPassed` — **CHECK, do not assume** |
| `seeders/content_stories_write.go:57-59` | score **and** `passed` taken from the source fixture |
| `seeders/assignment_plans.go:56-88`, `:180-215` | the plan seeder that already exists; writes literal `"active"` |
| `seeders/assignments.go:174-179`, `:210-224` | *"One program per org"*; the assignment status/bucket switch |

---

## (1) `completionForScore(score)` — the single derivation

TODO. Extraction, not invention — `hiring_funnel.go:346-350` is the shape. Decide where it lives
(a shared helper in `seeders/`) and what it returns (`sessCompletionPassed` / `sessCompletionFailed`).

## (2) The two-hash split in `jobsim_sessions.go`

TODO. Both directions must close: the un-clamped **failed** row (a failed 95) and the **[55,60)** hole
(a passing 57). See Open question 1 in `overview.md` — what happens to the story's `pass_rate` is the
design decision here, not the arithmetic.

## (3) The duplicate session id (`j == 0`)

TODO. `CopyRowsIdempotent(..., "id")` = `ON CONFLICT DO NOTHING` ⇒ first-writer-wins, and registration
order decides the winner. Two candidate resolutions (key change vs. seeder dependency) with different
determinism costs — see Open question 2.

## (4) The remaining session writers

TODO — `persona.go:286`, `ai_readiness_funnel.go:586`, `content_stories_write.go:57-59`. Measure the
actual score distributions first; the fixture-sourced case may not belong inside the fence at all.

## (5) The fence — the score↔verdict invariant over EVERY seeded session row

TODO. **The bug survived because nothing asserted score-verdict agreement.** Vehicle undecided
(`datadna` gene / Go unit test / `autoverify` SQL assert) — see Open question 5.

## (6) C1 — Programs: 3–5 plans per org across four stages

TODO. Extend the existing materializer rather than adding a seeder. Unresolved before design:
whether a stage is a stored `enum.PlanStatus` or computed from enrollment + assignment completion;
which orgs; whether "almost done" needs more than the single `ordered=false` cycle the current design
chose on purpose.

## (7) Docs

TODO — `seeding-spec.md` (the gene) and `stories-spec.md` (Programs as a seeded story surface).

---

## Pre-flight (before any code)

TODO — do the live Playthroughs that drive assignment surfaces (`pt-assignment-assign`,
`pt-assignments-nav-v2`) still pass, and does the seed-manifest honesty gate
(`CanonicalFileMatchesProjection`) still hold once plan counts change?
