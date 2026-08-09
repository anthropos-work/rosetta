---
iter: 219
milestone: M257x
iteration_type: tik
status: closed-no-lift
created: 2026-08-09
---

# iter-219 — the size-preserving-mutation hazard is PROVEN and has never been SIZED

**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` has been open and re-listed unchanged
since harden pass 42 — through passes 43…53, eleven consecutive re-listings. Re-surveyed at HEAD:
`test_mutation_proof_cache_hazard_m257x.py` **characterises** the hazard in three arms (it is real on
this interpreter; the mtime bump defeats it; clearing `__pycache__` does not) — and **enumerates
nothing**. The mechanism is proven and the population is unknown. `§5` iter-184: *a fence's POPULATION
is a registry too.*

## Cluster / target identified

Every mutation control in `stack-core/tests` that writes a **real file** — the exposed class is a
size-preserving write to a file the running interpreter will re-read.

**And the population may not be selected by NAME.** iter-218 proved that instrument wrong one iter ago:
its own sealed hand-enumeration said six and the by-effect probe said seven. Three static attempts here
returned **54**, **369** and **74** candidates against three different heuristics — *none of them a
population*. So the instrument is a **by-effect** one: patch the write primitives, run the suite, and
record what is actually written.

## Hypothesis

The real-file mutation population is small, nameable, and mostly *not* in the modules whose names say
`mutation` — and some of it is unprotected against the hazard pass 42 proved.

## Expected lift

`SURVEY-M257x-h42` closed with a measured population and a **permanent** arm, so the answer does not
have to be re-derived the next time someone writes a mutation control.

## Phase plan

1. **Seal** the three refuted static heuristics and the by-effect method, before the run.
2. Ship the by-effect recorder; run it over the **whole** `stack-core` suite.
3. Adjudicate every real-file write; classify size-preserving / mtime-bumped / restored.
4. Land the census as an arm with the measured population; re-pin ceilings.

## Escalation conditions

- If the recorder cannot observe writes without changing what the suite does, **refuse it** — an
  instrument that perturbs its subject measures itself.

## Acceptable close-no-lift outcomes

A measured **zero** exposed mutation — every real-file write already mtime-safe — closes this
`closed-no-lift`, **provided the recorder is proven to fire** on a staged write (`§9`).
