---
iter: 159
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-159 — the spelling-pin class: a sharper predicate, proved against the labeled set

**Active strategy reference:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.* A reading SAMPLES; a fence CENSUSES. This iter builds the census instrument for the class
that has produced **seven confirmed instances in five iters** and has never been enumerated.

## Step 0 — re-survey before targeting

`SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` is the largest standing route by instance
count and it is **SIZED but not started**. iter-155 sized it and rejected the sweep with the numbers:
**2,854** string-literal `assertIn`/`assertNotIn` across **109 files** in 5 sections, **766**
expression-shaped under the naive predicate — iter-150's 30-to-1 over-report at 30× the volume. Its
recorded next step is **"a sharper predicate, not the sweep."**

Re-survey confirms the target is untouched: iters 156, 157 and 158 each *added* an instance (7 now) and
none built the predicate. The route is current and still the largest.

**Per iter-158's lesson, the route's own framing is a hypothesis, not a plan.** The route says "sharper
predicate"; it does not say what makes one sharp. That is this iter's question, and it is answered from
evidence, not from the route's wording.

## Cluster / target identified

The class: **a fence pinned to the SPELLING of the code it guards rather than to its PROPERTY**
(`§5` rules 70/71). Every one of the seven was found by *improving the guarded code*, never by review.

## Hypothesis

**The naive predicate looked at the NEEDLE; the property lives in the HAYSTACK.**

Reading the four pre-repair assertions that are still recoverable from git shows the shape:

| instance | pre-repair assertion | haystack |
|---|---|---|
| harden-35 disclosure fence (repaired iter-153) | `assertIn(ui, window, …)` | a slice of `generate.sh` **source** |
| `dev-stack` contract test (repaired iter-154) | `assertIn('[ "$local_content" = 1 ] && verify_svcs="$verify_svcs directus"', self.BODY)` | `open(DEV_STACK).read()` — the script's **source** |
| `test_fence_provenance` (repaired iter-158) | `assertNotIn("force", src)` | `guard_family.py`'s **source** |
| `test_repair_postcondition::test_01` (repaired iter-157) | `assertEqual(on_disk, set(participating) \| set(standalone))` | *not* source text — a derived collection |

Three of the four share one mechanical signature: **the haystack is a repo source file read as raw
text.** That is decidable without interpreting the needle, and it is what separates a spelling-pin from
the 2,854. The fourth is a different sub-signature (frozen-collection equality) and the instrument is
expected to be **blind to it** — a stated blind spot, not a claimed zero (`D-M257x-158-3`'s discipline).

**The refinement that keeps it sharp:** a test may read a file that its subject *produced at run time*
(`open(self.log).read()`, a generated override, subprocess stdout). Those assertions are about
**behaviour** and are entirely legitimate. The discriminator is therefore not "reads a file" but
**"reads a file that is checked into the repo as source."**

## Expected lift

The route moves from **SIZED-and-rejected** to **ENUMERATED with a proven instrument** — a bounded
population with a stated denominator, replacing 2,854 unworkable candidates.

## Phase plan

- **A** — assemble the labeled set: the 7 confirmed instances as (file, pre-repair commit, post-repair commit).
- **B** — author the predicate; **prove the instrument on the labeled set** — it must fire on the
  pre-repair form and NOT on the post-repair form. Recall and blind spots stated as numbers.
- **C** — run the census over the population; report the enumerated count **and its denominator**.
- **D** — fence the instrument (tests, incl. a mutation control and an anti-vacuity control); gates; close.

## Escalation conditions

- **Predicate recall on the labeled set < 50 %** → the predicate is REFUTED. Record the refutation with
  numbers; do not ship a census that cannot find the instances we already know about (`§9` — a census
  must prove its instrument). Close on the falsification; do not substitute a weaker instrument.
- Population turns out unbounded (same 30-to-1 shape) → same refutation path.

## Acceptable close-no-lift outcomes

A measured refutation of the haystack hypothesis, with the recall number on the labeled set, satisfies
this iter. The labeled set is the deliverable that makes either verdict trustworthy.

## Explicitly OUT of scope (tripwire pre-declared)

**Running the population to zero is the SWEEP, and it is not this iter.** This iter delivers the
instrument + its proof + the enumerated population. The sweep routes forward as a named handler.
