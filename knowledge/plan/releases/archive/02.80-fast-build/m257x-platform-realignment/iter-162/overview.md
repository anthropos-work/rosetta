---
iteration_type: tik
status: closed-fixed
---

# iter-162 — the census's own registry is a hand-maintained tuple

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *build or extend a fence that enumerates every instance in the corpus, run it to zero, and keep it
green.* This iter works the **reach** half of that clause: a census whose registry is hand-listed
enumerates only what somebody remembered to list.

## Step 0 — re-survey (mandatory)

`frozen_expectation_census.py` at HEAD:

```
0 unexempted candidate(s) (2 declared exempt) over 9409 multi-token literals
in 108 test files, against 3 executed derivation(s)
```

Zero, and **3 executed derivations**. The route
`FIX-M257x-iter161-derivation-registry-is-three-entries` is therefore still live and still the
largest named residual on this instrument. Target confirmed, not substituted.

## Cluster / target identified

The census's docstring states its predicate as:

> **A test literal is a candidate iff its token set equals a value some non-test module derives.**

That is a claim about *every* derivation in the tree. Its implementation is `build_derivables()` —
**three hand-written `out.append(...)` calls**. The claim and the reach are different sizes, and
nothing measures the gap.

This is the milestone's own defining defect class turned on the instrument: `§2` of the protocol doc
(*"Why v2.8 was latent — the hand-maintained tuple"*), whose prescribed repair is *derive it at the
point of use*. It is also `§9`'s iter-159 rule — **grade the instrument at the grain of its claim** —
applied to iter-160's instrument rather than iter-159's.

**Measured population (this iter's first act, denominator stated per iter-114):**

| set | count |
|---|---|
| public functions in non-test rext modules returning a collection of `str` | **126** in 36 modules |
| …of those, callable here (0 required args, or path-like required args only) | **53** |
| …of those, in the census registry | **3** |

> **Corrected at close: the top figure is 125, not 126.** This survey used `ast.walk`, which also
> reaches functions nested *inside* functions; the shipped `population()` walks module and class
> bodies only, which is the right subject (a closure is not a module-level derivation). The
> executable-here **53** and the registry **3** are unaffected, and 53 is the denominator every reach
> claim in this iter uses. Recorded rather than overwritten — `§5` rule 22.

Reach is **3 of 53 = 5.7 %** of the executable-here sub-population — against a docstring that says
*every*.

## Hypothesis

Widening the registry from a hand-list to a **declared-and-fenced classification of the enumerated
population** will (a) make the reach checkable rather than assumed, (b) make a newly-added derivation
turn a fence RED until it is classified, and (c) surface frozen copies of derived values that three
entries could not see.

## Expected lift

Primary deliverable is instrument reach, not `P`/`N` (no reading is taken this iter — `§9`: the metric
stays **UNMEASURED**, not unmoved). Success = the population enumerated mechanically, every member
classified with a reason, the registry grown, the census re-run to zero at the wider registry, and the
completeness fenced.

**Falsifiable:** if widening the registry surfaces **zero** new candidates, that is a real and
reportable outcome — it says the class was genuinely small, not that the sweep failed. What would
*refute* the iter is being unable to enumerate the population mechanically at all.

## Phase plan

- **A** — enumerate the population mechanically; classify every member REGISTERED / DECLINED-with-reason.
- **B** — implement the widened registry; execute the registered derivations.
- **C** — re-run the census; grade every new candidate **at source** (iter-158's rule: a proposed
  repair is a hypothesis, not a plan).
- **D** — fence completeness (a new derivation must be classified) + extend the commit-pinned labeled
  set (iter-161's rule: a labeled set that reads the working tree decays the moment you use it).

## Escalation conditions

- A registered derivation needs a live external system (docker, a running stack) → **DECLINE it with
  that reason**; do not weaken the `DerivationUnavailable` fatality that `§9` requires.
- Grading a new candidate turns out to need a platform-repo edit → route forward, never edit.

## Acceptable close-no-lift outcomes

The population is enumerated and every member classified, but **no new candidate fires** — the reach
grew, the class was already at zero, and the iter's deliverable is the *checkable* reach.
