---
iter: 218
milestone: M257x
iteration_type: tik
status: closed-fixed
created: 2026-08-09
---

# iter-218 — iter-217 taught ONE of six number-matchers to read bold; the other five never learned

**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

iter-217 closed `SURVEY-M257x-h48` by teaching `derivation_registry._MEASURED_RE` this repo's emphasis
idiom. Re-surveyed immediately after: **`_MEASURED_RE` is not the repo's only number-matcher.** A
census by effect over every `re.compile` that pairs a digit with a word returns **six constructs across
four modules**, and the split is the finding — **three already knew about `**` and all three live in one
module**; the other three do not, and one of them is a **live corpus fence**.

This is iters 209–212's class exactly, one construct family over: one rule, several spellings, repaired
one member at a time. **It is also iter-217's own repair, one iter late** — which is why it is the
target rather than a note.

## Cluster / target identified

`platform_predicate_guard._REPO_COUNT` — **G2 repo-count**, the check that grades every corpus sentence
claiming how many repos the clone set has. It is the fence closest to this milestone's own subject:
`repos.yml` membership changed three times inside the window M257x covers (skillpath out, then
`storage` + `messenger` out at `838d907`).

## Hypothesis

The corpus bolds its figures, so G2 cannot see the very claims it exists to grade. Widening it will
surface at least one live repo-count claim that no check has ever read.

## Expected lift

A live fence's blind spot **sized, adjudicated and closed with zero false REDs** — iter-209's
precondition governs, and it is not negotiable: a widening that turns the live corpus RED on a correct
sentence is refused, not landed.

## Phase plan

1. **Seal** V1…V5 in this iter's FIRST commit, before any repair.
2. Widen `_REPO_COUNT` and adjudicate **every** newly-visible match by hand.
3. Size the two remaining blind siblings and dispose of each by name.
4. Arms + a mutation control; verify G2's finding count against the pre-registered stop condition.

## Escalation conditions

- **Any** new false RED on the live corpus → refuse the widening, route it, close on the measurement.

## Acceptable close-no-lift outcomes

A measured zero — the corpus never bolds a repo count — closes this `closed-no-lift`, **provided the
instrument is proven to fire** on a staged fixture (`§9`).
