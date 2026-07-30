---
iter: 24
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-24 — the seat-append rule, pinned; and the "stage-0 seat" that cannot be declared

**Active strategy reference:** `TOK-01` move 3's successor clause — onboarding, ordered second precisely
because *"its cost is a seeder + capability + roster seat, not just specs."* This iter prices the seeder
half before any spec is written.

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| `ONBOARD-M256-seat-append` | still owed; named by iter-18 D89 as the prerequisite for **every** remaining onboarding UC |
| onboarding UCs landed | **1 of 5** (`onboarding.completion.UC1`), 4 `TODO` |
| `personaUserIndexFor` (`seeders/persona.go:510`) | `idx = i + 1`, **declaration order** — so appending is safe and inserting renumbers |
| the cheapest remaining UC's recorded blocker | `enterprise-workforce-ai-readiness.UC1` → *"an Org C stage-0 seat"* |
| `aiReadinessStageFor` (`seeders/ai_readiness_funnel.go:198`) | a hero's `default` branch returns **stage 3** — read, not yet measured |

## Cluster / target identified

`ONBOARD-M256-seat-append`. iter-18 D89 established the constraint in prose (*"append only —
`personaUserIndexFor` indexes by declaration order, so inserting mid-list renumbers existing personas and
breaks every seeded reference"*). **Prose is not a fence.** Nothing stops the next author from inserting a
hero in the natural place — next to the hero it relates to — and silently re-pointing every seeded
reference for the personas below it.

And while pricing the first UC that consumes the mechanism, a second thing surfaced that changes what the
next iter should attempt.

## Hypothesis

1. **Append-only safety is testable and untested.** A test that declares a hero list, appends to it, and
   asserts no pre-existing index moved — plus its mutant (insert mid-list) going RED — converts D89 from a
   comment into a guard.
2. **The routed blocker for `enterprise-workforce-ai-readiness.UC1` is mis-stated.** It reads *"needs an
   Org C stage-0 seat"*, which implies the seat is declarable and merely absent. `aiReadinessStageFor`'s
   hero branch appears to have **no stage-0 outcome for an end-user** — manager → 0, struggling → 1,
   everything else → 3 — so an appended Org C end-user hero would arrive **COMPLETED**, and a Playthrough
   built on it would drive a hero who has already done the thing it means to watch her do. **Measure it.**

## Expected lift

**No gate clause moves this iter.** Planned deliverables:

- the append-only property **pinned by a test**, mutation-verified;
- hypothesis 2 **measured** — confirmed or refuted — and the four remaining onboarding UCs re-priced
  against the result, so iter-25 plans from a measured blocker rather than a paraphrased one.

Deliberately **not** in scope: changing the seeder to support a stage-0 end-user, and editing
`pt-world.seed.yaml`. Both are real work and both would be started on an unmeasured premise if done now —
which is the exact failure iter-21 avoided by probing before shipping a seeder.

## Phase plan

- **A — measure hypothesis 2** with a test that declares an Org-C-shaped story, appends a 4th end-user
  hero, and reports the stage the seeder assigns it.
- **B — pin the append-only rule** (and hold the measured stage-0 gap as a test, so whoever adds stage-0
  support is told what to change).
- **C — re-price the 4 remaining onboarding UCs** in the routing table against A's result.
- **D — verify** (`go test` in `stack-seeding`; no live-stack change, so no gate run is owed — nothing
  shipped touches the harness).

## Escalation conditions

- If hypothesis 2 is **refuted** (an appended end-user hero does land stage 0) then the routed blocker is
  correct as written, the append is genuinely all that is needed, and the honest move is to say so and let
  iter-25 land the UC.

## Acceptable close-no-lift outcomes

Either verdict on hypothesis 2 is a complete outcome. What would not be acceptable is appending a seat to
the real seed on the assumption that stage 0 is what it produces.
