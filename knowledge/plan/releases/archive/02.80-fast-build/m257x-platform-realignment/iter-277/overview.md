---
iter: 277
milestone: M257x
iteration_type: tik
status: closed-fixed-partial
opened: 2026-08-10
handler: clause-5-reading
---

# iter-277 — read gate clause 5, the last open clause

**Type:** tik

## Step 0 — Re-survey before targeting (mandatory)

1. **Clause 2 is MET as of iter-276** — 30 live / 0 failing / 0 error, twice cold, the second run from
   the stack's own clone at an origin tag. Re-confirmed from iter-276's own logs, not assumed.
2. **Clause 5 is therefore the ONLY open clause.** Clause 1 met (3 consecutive green cycles at the
   shipping pin); clauses 3 and 4 hold; clause 2 met this run.
3. **The clause-5 target is unchanged and still stale.** Last read at **iter-131: `P = 29 / N = 47`**,
   corpus at that iter's ref — and that reading's own headline is that **the test-retest overlap with
   iter-119 was ~0**, i.e. two consecutive readings produced almost disjoint predicate sets. `P` is a
   **floor**, never an estimate of the pool.
4. **The standing KB-fidelity audit record is from 2026-07-31** (`kb-fidelity-audit.md`, iter-01
   Phase 0b): **YELLOW with 3 blocker-severity findings**. That is 275 iters and an enormous volume of
   repair ago. Clause 5's literal text is a verdict over that instrument, so the record is the thing
   that must be refreshed.

## Active strategy reference

**`TOK-08` — *census the mechanical classes; stop sampling them*.** Binding, and it constrains this iter
in a specific way the plan must respect:

> *A reading SAMPLES; a fence CENSUSES.*

TOK-08 also carries a **sealed pre-registered falsification** that **bars authoring a successor
strategy**: after one full mechanical sweep, `P ≥ 19` → *enumeration-first is refuted, STOP, hand back to
the user for a scope decision*; `P ≤ 18` → *the method is working, say so with the number*. iter-131
recorded that no successor strategy is authorable. **This iter therefore measures; it does not invent a
ninth strategy**, whatever the number says.

## Cluster / target identified

Clause 5 as written: *"KB-fidelity audit **GREEN, or YELLOW with 0 blockers**, over `corpus/services/**`
+ `corpus/architecture/**`."*

Two instruments, deliberately distinguished — the milestone has conflated them before:

| instrument | what it is | what it settles |
|---|---|---|
| the **mechanical census** (the `stack-core` fence suite) | 51 guards + 93 test modules run to zero findings | the classes where a claim resolves or does not, no interpretation |
| the **KB-fidelity reading** (`P`) | a semantic reading over the two globs | what a fence cannot decide |

TOK-08's whole point is that the first must be exhausted before the second is trusted, because a reading
at ~60 % per-pass recall cannot enumerate a pool this size.

## Hypothesis

The mechanical classes are at zero (the fences have been run RED-watched for many iters), and the
residual is semantic. The reading will therefore be a **floor over the semantic residual**, and the
honest output of this iter is **a number with its denominator and its method stated**, not a verdict
argued into GREEN.

## Expected lift

A **current** clause-5 reading — the milestone has not had one for 146 iters. Whether it *meets* the
clause is a measurement outcome, not a target: clause 5 is met only by a reading that returns zero, and
the user has ruled four times that it is not re-cut, reinterpreted, narrowed or argued.

## Phase plan

- **A** — run the mechanical census (`stack-core` suite, ~40 min, exit code captured to a file so it
  cannot be lost). This is the TOK-08 instrument.
- **B** — refresh the KB-fidelity reading over `corpus/services/**` + `corpus/architecture/**`.
- **C** — report `P` with its denominator, method and date; grade clause 5 honestly.
- **D** — close, and state the gate position exactly.

## Escalation conditions

- If the census returns RED, that is the finding and it outranks any reading — a corpus with a red fence
  is not GREEN by any argument.
- If `P > 0`, clause 5 is NOT met and the milestone does not close on this run. Say so plainly.

## Acceptable close-no-lift outcomes

A reading that returns `P > 0` is a **complete iter**. The deliverable is the measurement, not a green.
