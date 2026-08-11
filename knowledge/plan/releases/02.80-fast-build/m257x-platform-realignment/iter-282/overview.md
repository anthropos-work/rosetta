---
iter: 282
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-11
---

# iter-282 — census the corpus/tooling prose copies, and grade the copies that disagree

**Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.* A reading SAMPLES;
a fence CENSUSES. Work the classes in descending measured size, report the enumerated population, and
state the denominator.

## Step 0 — re-survey (mandatory; three consecutive iters have opened on a stale route)

`ROUTE-M257x-h70-corpus-and-code-prose-are-copies-with-no-fence` was re-verified **open** before any work
started: `guard_family.py`'s member list, the `stack-core/tests/` module list and a grep for a
prose/copy/verbatim fence all return nothing that grades this class. The route's sizing measurement
(**172 (module, doc) pairs sharing a verbatim 11-word run**, harden pass 70) and its bidirectionality
(pass 71) stand unchallenged and unfenced.

## Cluster / target identified

The route is the **largest remaining structural item on limb 3**, and the milestone has paid for it at
least three times — a §8 passage describing three arms of a four-armed fence, copied out of a docstring
that still said three; a docstring pinning a module count while its own text said floors; a §10 row
asserting *"zero open user questions"* that §8 had retracted 160 iters earlier.

## Hypothesis

**Copying is not the defect; DRIFT is.** 172 shared runs cannot all be repaired and most should not be —
a corpus that quotes its tooling is doing the right thing. What is mechanically decidable, needs no
domain knowledge, and is exactly the failure the three instances share: *where one sentence appears in
both trees and the two copies carry different numbers in the same slot, one of them is stale.* Fence
that, and the readings only have to catch what is genuinely semantic.

The fence must be **direction-blind**: pass 71 measured the coupling running both ways, and a fence that
named an original would close whichever half happened to be found first.

## Expected lift

A `*_guard.py` member enumerating the class corpus-wide, at **zero findings**, with the population and
its denominator stated — plus the real divergences it enumerates repaired at every site, not at the
first one found.

## Phase plan

- **A** — size the population before designing (probe, scratchpad); triage every candidate against source.
- **B** — build `prose_twin_guard.py`; prove it RED on the live tree before any repair.
- **C** — repair; run to zero; controls + regression tests; enrol in the guard family.
- **D** — re-measure; keep the section gate GREEN (it entered this iter at `2240 passed / 3 skipped / 0 failed`).

## Escalation conditions

- If the false-positive rate makes "run it to zero" force a **wrong** repair at any site, the fence is
  wrong, not the prose — narrow the predicate or waive **with a recorded reason**; never edit a true
  sentence to satisfy an instrument.
- If the population turns out to be dominated by legitimate template reuse, say so with the number and
  scope the fence to the sub-population that is real.

## Acceptable close-no-lift outcomes

A measured refutation that the class is not mechanically fenceable — stated with the population, the
false-positive classes and the reason — is a complete iter under this protocol.
