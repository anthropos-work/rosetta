---
milestone: M268
title: "Seeded truth"
milestone_shape: section
status: planned
release: "02.100-content-consolidation"
depends_on: "none"
parallel_with: "M266, M270"
complexity: large
last_updated: "2026-08-23"
---

# M268: Seeded truth

**Goal:** What the demo shows about a session is internally consistent — a score and its verdict never
disagree — and the Programs surface is populated rather than empty.

Serves annotation requests **C2** (*"sometimes you show a simulation with score above 60 as red/failed,
and viceversa below 60 as passed/green"*) and **C1** (*"each program section … show at least 3 to 5
programs at different stage"*).

## Scope

**In:**

  - **(1) Extract `completionForScore(score)`** — ONE derivation of verdict from score, at the platform's
    own threshold of **60** — and route `jobsim_sessions.go`, `feedback.go` and `hiring_funnel.go`
    through it. **The correct pattern already exists in this codebase**: `hiring_funnel.go:346-350`
    derives completion FROM the score. This is an extraction, not an invention.

    **60 is the platform's own threshold, not ours.** Seeded
    `validation_attempt_results.success_threshold` is written `60.0` at **`persona.go:303`** and
    **`content_stories_write.go:310`**, and **`hiring_funnel.go:74`** declares
    `hiringPassThreshold = 60 // score ≥ this ⇒ completition_status=passed, else failed`.

  - **(2) Kill the two-hash split at `jobsim_sessions.go:142-157`.** Score and verdict are derived from
    **TWO DIFFERENT, INDEPENDENT HASHES**:

    ```go
    score  := growthArcScore(prefix, i, frac)            // continuous 0-100, from key+":t"
    passed := int(hashInt(key+":p")%100) < passThreshold // a DIFFERENT hash. Independent.
    if passed && score < 55 { score = 55 + score/3 }
    ```

    Two defects fall out of it, in **opposite** directions — which is exactly what the reviewer saw:
      - **(a)** A **FAILED** session is **never clamped down** — `passed == false` leaves `score`
        untouched, so a failed row can carry a **95**. That is *"score above 60 shown as red/failed"*.
      - **(b)** The nudge floor is **55, not 60**, and it does not fire in **[55,60)** — a "passing"
        **57** stays 57. That is *"below 60 shown as passed/green"*.

  - **(3) Resolve the duplicate-session-id collision — the THIRD defect, and the reason this is not a
    one-liner.** TWO SEEDERS WRITE THE SAME SESSION ID UNDER DIFFERENT RULES:

    ```
    jobsim_sessions.go:136  key     := fmt.Sprintf("%s:session:%d:%d", prefix, i, j)
    feedback.go:153         sessKey := fmt.Sprintf("%s:session:%d:%d", prefix, i, 0)
    ```

    For `j == 0` these are **BYTE-IDENTICAL**, so `deterministicUUID` yields the same
    `public.job_simulation_sessions.id`. Both compute the same score; **`feedback.go:196-199` uses the
    score-consistent rule (`score >= 55`)** while **`jobsim_sessions.go` uses the coin flip**. Both COPY
    with `CopyRowsIdempotent(..., "id")` = `ON CONFLICT DO NOTHING`, so **FIRST WRITER WINS** — and
    `cmd/stackseed/main.go` registers `JobsimSessionsSeeder` at **`:525`** *before* `FeedbackSeeder` at
    **`:540`**, so **THE INCONSISTENT WRITER WINS** and the consistent one is silently discarded.

  - **(4) CHECK the remaining session writers — do not assume.**
      - **`persona.go:286`** and **`ai_readiness_funnel.go:586`** write an **unconditional**
        `sessCompletionPassed`; their scores need a **`>= 60` floor** for the invariant to hold.
      - **`content_stories_write.go:57-59`** takes **both** score and `passed` **from the source
        fixture** (a real, cloned production session) — it is not a derivation at all, and how it is
        treated is an open question below, not a settled scope item.

  - **(5) THE DELIVERABLE BEYOND THE FIX — a data-DNA / fence test asserting the invariant over EVERY
    seeded session row**, so this cannot regress silently. **The bug survived because nothing asserted
    score-verdict agreement.** A fix without the fence re-opens the same hole on the next seeder.

  - **(6) C1 — Programs seeding.** A plan seeder **ALREADY EXISTS**:
    `stack-seeding/seeders/assignment_plans.go:56-88` and **`:180-215`**, with `assignments.go:174-179`
    and **`:210-224`**. The gap is that it writes exactly **ONE `status=active` plan per org**
    (`assignments.go:174-179`, *"One program per org"*; `assignment_plans.go:180-215` writes the literal
    `"active"`). Extend to **3–5 plans per org** spanning **completed / just-started / in-progress /
    almost-done**, with **per-plan progress consistent with member counts**. The surface is
    `/enterprise/assignments-list` ("Programs").

  - **(7) Docs** — the score-verdict invariant written up as a documented gene in
    [`seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md), and Programs written up as a seeded
    story surface in [`stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md).

**Out:**

  - **The Programs product UI itself** — that is platform (`next-web-app`) code. This milestone seeds
    the data the existing surface reads; it does not change the surface.
  - **New assignment SEMANTICS beyond stage variety** — no new resource types, no new enrollment model,
    no scheduling/window behaviour. More plans, at more stages, is the whole of C1.

## Depends on

**none.** M268 is independent of M267 — different tables. It touches `rosetta-extensions/stack-seeding`
only, which is why it starts cold.

## Parallel with

**M266, M270.** M266 is `demo-stack/cockpit.py` only; M270 is `next-web-app` via demopatch. No file and
no repo pair is shared with either.

## Open questions

These are genuinely open. Each is a place where the measurement stops and a design decision starts.

  - **Does deriving the verdict FROM the score preserve each story's `pass_rate`?** Today the dependency
    runs the other way: `passed` is driven by the story's `pass_rate` and the score is nudged
    afterwards (`jobsim_sessions.go:142-157`). Routing everything through `completionForScore(score)`
    **inverts** that, so `pass_rate` either becomes an input to the score **band** or it stops being
    honoured. Which one is not decided, and the choice changes every seeded org's visible pass mix.
  - **Who owns the shared `%s:session:%d:%d` key for `j == 0`?** Either `feedback.go` stops writing the
    session row and takes a dependency on `JobsimSessionsSeeder`, or the key namespace changes.
    **Changing the key changes the deterministic UUID**, i.e. a reseed is no longer byte-identical with
    a prior one — whether that is acceptable is a decision, not a measurement.
  - **Do `persona.go` / `ai_readiness_funnel.go` scores already sit above 60?** Not measured. If they do
    not, a `>= 60` floor changes the hero narrative — a "struggling" hero whose every verified session
    now reads as passed is a different story than the one `stories-spec.md` documents.
  - **What happens when a REAL cloned session disagrees with the invariant?**
    `content_stories_write.go:57-59` copies score and `passed` from a real production session. A real
    session that disagrees is **truth, not a bug**. Normalising it would be falsifying a real record;
    leaving it would leave one class of row outside the fence. Undecided — and the fence's scope
    depends on the answer.
  - **Where does the fence live?** A `datadna` gene, a Go unit test in `stack-seeding`, or an
    `autoverify` SQL assert are all plausible and are not equivalent (one grades a blueprint, one
    grades the code, one grades a live stack). The **documented** gene lands in `seeding-spec.md`
    either way.
  - **(C1) Is a program's stage a stored `status` value or computed from its assignments?** NOT
    measured. `assignment_plans.go:180-215` writes the literal `"active"` and its own comment names only
    `active` and `draft` as `enum.PlanStatus` values. Whether "completed" / "almost done" is a plan
    status, or is derived by the surface from enrollment + assignment completion, decides whether C1 is
    a status-variety change or a progress-fan-out change.
  - **(C1) Which orgs get 3–5 plans?** All four (3 workforce + the M223 hiring org) or workforce only —
    the generic activity seeders already skip a hiring org via `hiring_scope.go`.
  - **(C1) Does "almost done" need more than one cycle per plan?** The current design is deliberately
    **one cycle, `ordered=false`** (`assignment_plans.go:180-215` explains why: an ordered cycle would
    lock every step but the first and render the program unstartable). A multi-stage program may need
    more, and that would be re-opening a decision that was made on purpose.
  - **Does multiplying plans trip the seed-manifest honesty gate?** The manifest is a *projection* of
    the canonical presets and is honesty-gated (`CanonicalFileMatchesProjection`). If plan counts are
    preset-visible, the canonical file must be regenerated in the same commit.
  - **Do any live Playthroughs go red?** `pt-assignment-assign` and `pt-assignments-nav-v2` both drive
    assignment surfaces. Not checked — it is a pre-flight for this milestone, not an assumption.

## KB dependencies

- [`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the **data-DNA** and the
  **production-isolation boundary** (write-side). The new fence gene lands inside the data-DNA, and every
  extra Programs write stays inside the isolation contract.
- [`corpus/ops/demo/stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md) — the **7-table
  verified-skill fan-out** and the **G14 session-seeder rules** (valid `SIMULATION_TYPE_*` / enum /
  token + continuous growth-arc score). C2 changes how a seeded session's verdict is derived; G14 is the
  contract it must not break.
- [`corpus/ops/demo/seed-manifest-spec.md`](../../../../../corpus/ops/demo/seed-manifest-spec.md) — the
  single auditable `seed-generation-manifest.yaml` projection + its honesty gate, which is what a change
  in seeded population (more plans per org) has to stay consistent with.

## Delivers →

- [`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the **score-verdict
  invariant becomes a documented gene**: one threshold (60), one derivation, asserted over every seeded
  session row, with the three-defect history recorded so the next seeder does not re-introduce it.
- [`corpus/ops/demo/stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md) — **Programs joins
  the seeded story surface**: 3–5 plans per org at distinct stages, and what a Program's stage is
  materialized from.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** M268's work is entirely inside `rosetta-extensions/stack-seeding`, so no
  platform source change is expected. If one turns out to be needed — e.g. the Programs surface reads a
  column no seeder can populate — it goes through the **sha-pinned `demopatch` mechanism**
  ([`corpus/ops/demo/demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)), never a repo
  edit, and an un-patchable surface **escalates**.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged,
  **pushed to origin**, then consumed per-stack at a pinned tag. *Tagging is not publishing.*
- Secrets handled values-blind.
