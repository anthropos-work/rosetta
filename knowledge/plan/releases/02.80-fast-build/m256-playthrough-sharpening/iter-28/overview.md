---
iter: 28
iteration_type: tik
iter_shape: standard
status: closed-fixed-partial
opened: 2026-07-30
---

# iter-28 — the org-prepared onboarding trigger, found by reading the platform

**Active strategy reference:** `TOK-01` move 3's onboarding clause, under iter-27's lesson: *when the question
is "what causes this?" rather than "what does this render?", read the code.*

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| onboarding | **3 of 5** (`completion.UC1`, `…ai-readiness.UC1`, `enterprise-hiring.UC1`) |
| clauses 1 + 2 | MET (196×3, 0 flake; controls 26/28, MUTATES 10, blocked 1/1) |
| `standard.UC1` | still behind the measured CV-upload product defect — iter-18's refusal stands |
| `standard.UC2` | trigger **unidentified** after iter-18 measured FOUR orgs |
| `individual.UC1` | needs a **member-less user** — the large-blast-radius one |

**Target: `standard.UC2`.** Its blocker is an *unknown*, not a cost — and iter-27 just demonstrated that an
unknown of this shape (*what selects this variant?*) yields to a source read in minutes after resisting a probe
sweep for two iters.

## Cluster / target identified

The org-prepared onboarding variant. Both apps mount the **same** `OnboardingUser` component, so whatever
selects the prepared variant is a prop or a query inside it — findable by reading, not by driving more orgs.

## Hypothesis

The trigger is a seedable input. If it is, the UC lands; if it is a runtime-only or platform-flag input, the
verdict upgrades from *"not yet identified"* to a precise, cited reason.

## Expected lift

clause 3: onboarding **3 of 5 → 4 of 5**, with the UC's negative control.

## Phase plan

- **A — read the component** and find the trigger.
- **B — the capability**, with mutants, if the trigger is seedable.
- **C — the seat**, reseeded and PROVEN live to render the prepared variant.
- **D — the Playthrough + its control.**
- **E — the gate.**

## Escalation conditions

- If the trigger is not seedable → stop at a cited verdict; do not build a Playthrough over a variant the seed
  cannot produce.
- **If the completion path beyond the prepared summary is unmeasured when the session budget runs short → stop
  at the proven capability and route the Playthrough.** Asserting an unmeasured multi-step journey is how a
  spec gets written against a surface nobody drove — the failure iter-22 and iter-27 each paid for once.

## Acceptable close-no-lift outcomes

A cited, precise verdict on the trigger is a complete outcome. So is the capability proven live without its
Playthrough, provided the residual is routed with the measurements a successor needs.
