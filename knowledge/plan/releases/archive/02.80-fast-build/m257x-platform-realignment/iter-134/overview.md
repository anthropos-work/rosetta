---
iter: 134
milestone: M257x
iteration_type: tik
iter_shape: audit
status: closed-fixed
opened: 2026-08-08
---

# iter-134 — was the retraction blindness a pattern, or an outlier?

## Active strategy reference

**No successor strategy is authorable** — `TOK-08`'s sealed refutation branch bars one; running under
the user's direct brief, as iters 121–133 did.

## Step 0 — re-survey

`FIX-M257x-iter132-marker-fences-cannot-see-retractions` is open, opened by iter-132 two iters ago, and
its wording is a **conjecture, not a finding**: *"the same substring-vs-retraction blindness plausibly
affects every marker-matching fence in the family. Nobody has checked the others."*

**An open route whose text says "plausibly" and "nobody has checked" is a measurement waiting to be
taken, and this milestone's rule is that we take it rather than inherit it** (`D-M257x-121-1`,
re-derive at source before filing; `D-M257x-132-4`, an inherited route is evidence, not instruction).

## Cluster / target identified

The four fences that classify **prose by substring**: `claim_twin_guard`, `platform_predicate_guard`,
`claim_census_guard`, `unreadable_repo_claim_guard`.

## Hypothesis

**Stated so it can fail:** if the blindness is a *pattern*, ≥ 2 of the other three carry no retraction
handling and the route is upheld. If it is an *outlier*, iter-132's fence was the only one missing it
and the route is **refuted** — in which case the interesting question changes from *"fix the others"*
to *"why did iter-132 build a bespoke fix?"*

## Expected lift

**No `N` reading; no `N` movement claimed.** The deliverable is the audit and its verdict either way.
An audit that returns *"the conjecture is false"* is the iter's product, not a failure to produce one.

## Phase plan

1. Enumerate the marker-matching fences.
2. Probe each for five independent retraction mechanisms — **by importing and interrogating the module,
   not by grepping the file** (rule 22: a grep is testimony, the loaded module is evidence).
3. Grade the conjecture against the pre-stated branch.
4. Land whatever the verdict implies; route what it does not.

## Escalation conditions

- If ≥ 2 fences are blind, the scope is larger than one iter → land the audit, route each fence
  separately with a named handler, do not attempt four fixes in one iter.

## Acceptable close-no-lift outcomes

- **The conjecture is refuted.** Recording a refuted conjecture with its measurement is a complete iter
  under this protocol; the milestone has closed nine of them that way.
