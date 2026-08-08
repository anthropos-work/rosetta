---
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
closed: 2026-08-08
---

# iter-142 — the retraction idiom, censused

**Active strategy reference:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.* A reading SAMPLES; a fence CENSUSES. Work the classes in descending measured size, build
or extend a fence that enumerates every instance corpus-wide, run it to zero, keep it green, and
**report the enumerated population, how many were already false, and the fence's reach — stating the
denominator** (iter-114's rule).

**Cluster / target identified:** `FIX-M257x-iter141-retraction-idiom-sweep`, opened by the immediately
prior iter and the newest route on the queue. `TOK-08`'s next-tik direction names no specific class
beyond *"descending measured size"*; iter-141 measured this one's motivation directly — **the corpus's
own retraction idiom turned fences RED three times in five iters, in three different files, always on a
pin whose own sentence existed to retract it** (`roadrunner.md`, `graphql-wundergraph.md`'s `5050`
pointer which rotted twice on its own, and `ai-readiness.md`'s note, that one caused by the very iter
repairing the class). It is also **censusable under `D-M257x-140-2`** — *a class is censusable iff its
subject carries its own HEAD* — because the subject here is **the citing sentence itself**, not a
resolved target. That is exactly what iter-138's retracted census lacked.

**Re-survey (Step 0, mandatory).** Confirmed the target is untouched and meaningful before planning:
`rg` over `corpus/ + README.md + CLAUDE.md` with a positive control (981 backticked `:NN` occurrences
across 69 files; 92 files in the search set) returned 46 raw candidate lines for the idiom. Not absorbed
by any prior iter. Pre-iter guard family re-measured GREEN first: **18 GREEN · 0 RED · 4 not-run**,
identical to iter-141's close.

**Hypothesis.** The hazard is not the *retraction* — it is the **token shape**. A retracted pin written
as `` `:274` `` is byte-identical to a live assertion, so every resolver binds it and every insertion
above its target rots it; a fence matching on form *cannot* tell the quotation from the assertion, and
rule 63(c′) says it is right not to. Therefore: **fence the TOKEN, not the digit.** A retraction that
*describes* ("rotted +8", "carried two different numbers in successive iters") keeps every bit of the
evidence and is invisible to every resolver.

**Expected lift.** No `N` reading is planned or claimed this iter — the metric is expensive and
`TOK-08` sequences the read after the mechanical sweep. The iter's deliverable is a **class run to
zero plus the fence that holds it there**, which is `TOK-08`'s unit of progress.

**Phase plan** (three planned lines — a census-then-repair-then-fence shape, declared here so the
scope-creep tripwire counts against the planned shape and not against a single-target tik):

1. **Census.** Author `retracted_pin_guard.py` in the rext authoring copy; enumerate the class
   corpus-wide with a stated denominator; **hand-audit every finding for precision before repairing a
   single line** (§5 rule 63(a) / the iter-138 lesson — a mechanical predicate is not a measurement
   until its precision is measured).
2. **Repair.** Convert every enumerated site from reproduction to description, preserving line counts
   wherever possible so the repair does not induce the rot it is removing.
3. **Fence.** Register the guard in `guard_family.py`; ship it with a **mutation control** and an
   **anti-vacuity control written against the guard's SUBJECT** (§8 iter-94); re-run the family and the
   change-scoped suites (rule 63(d) — choose suites by what you CHANGED).

**Escalation conditions.** If the hand-audit shows precision materially below 100 %, the guard is
retuned or the class is re-specified **before** any repair — under no circumstances is a number
published from an unaudited predicate. If the repair cannot preserve a site's evidence without
reproducing the pin, that site is **disclosed**, not silently deleted.

**Acceptable close-no-lift outcomes.** A falsification that the class is not mechanically separable
from ordinary description (i.e. the precision audit fails) would close this iter no-lift with the
measurement recorded — that is a complete cycle under this protocol, and it would retire rule 63(c′)'s
enforceability claim rather than leave it asserted.
