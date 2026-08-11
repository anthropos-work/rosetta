---
iter: 214
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-09
---

# iter-214 — the claim fences read the corpus; the TOOLING's prose is outside all of them

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*
Class under work: **the population of the already-refuted-claim fence.**

## Cluster / target identified

iter-212 found a **retracted claim living in a rext Python docstring**, acting as a design rationale,
for a whole iter. `claim_twin_guard` (`FENCE-M257x-iter42`) is the fence for *"has an already-refuted
claim come back?"* and its stated scope is *"the published tree — `corpus/**`, `.claude/skills/**`,
`CLAUDE.md`, `README.md`"* under a heading reading **"SCOPE IS TREE-WIDE FROM THE FIRST RUN,
DELIBERATELY"**. Tree-wide means the **rosetta** tree. `rosetta-extensions` — the repo that holds every
fence, every rationale, and `CLAUDE.md`'s own *"all stack-operating tooling lives in
rosetta-extensions"* — is in **no** claim fence's population. Same shape as iter-207 and iter-209, one
repo over.

## Hypothesis

The class is censusable with the guard's OWN machinery against a different source set, and the answer
partitions cleanly enough to decide widening by measurement rather than by argument.

## Pre-registered, sealed in this iter's FIRST commit — before any repair

Measured at corpus `f7b9643` / rext `6daab5e`, guard machinery held fixed (`G.derive`,
`claim_ledger.normalize_document`, `G.find_form`, `G._looks_retracted`); only the source set varies —
iter-209's discipline.

- **T1** — the ledger holds **264** adjudicated claims. The guard's own population is **114** documents.
- **T2** — outside it, in `rosetta-extensions`: **186** Python modules and **211** markdown documents.
- **T3 — the markdown partition, and it decides the iter:**

  | bucket | docs | hits | in a retraction context |
  |---|---:|---:|---:|
  | `tests/fixtures/**` | 138 | **217** | 39 |
  | test modules | 1 | 0 | 0 |
  | **real prose** | **72** | **0** | 0 |

- **T4 — the ZERO proves its own instrument, in the same run.** `§9` requires a census returning zero to
  show the instrument fires. It does: the identical matcher over the identical claim set returns **217**
  on the fixture bucket. The 0 is a measurement, not a silence.
- **T5 — the Python half is NOT clean: 10 hits.** Inspected, the real-prose ones are **fences quoting
  the very claim they exist to catch** (`claim_ledger.py`, `derived_value_guard.py`,
  `value_change_guard.py`). Only 3 of 10 sit in a retraction context. **Widening to `*.py` would
  manufacture ~7 false REDs**, and iter-209's stated precondition for widening a fence's population at
  all is **zero false REDs measured first**.
- **T6 — the honest one, and it limits this iter's own claim.** iter-212's defect **would not have been
  caught by this fence even widened**: that claim was retracted in a *route ledger*, never adjudicated
  in a blocker-ledger, so it is not among the 264. The fence that would have caught it does not exist.

**STOP CONDITION, sealed before the repair:** if the landed arm reports any finding over the 72
real-prose documents, **do not land it as a green assert** — report the findings and route the repair.
A new arm whose first run is RED is a repair task, not a fence.

## Expected lift

The class is closed as a class: a running census over the population no claim fence reads, carrying its
own anti-vacuity proof, plus a **measured** refusal for the half that cannot be widened cleanly.

## Phase plan

A: seal this record. B: land the census as a running arm over rext real prose + its fixture-bucket
anti-vacuity twin. C: record the Python refusal with its measured cost and T6's limit. D: re-run.

## Escalation conditions

Stop condition fires → land the census + the refusal only, route the repair.

## Acceptable close-no-lift outcomes

T3's real-prose zero being an artefact of the matcher rather than a fact about the prose — which T4 is
designed to detect and which would be the deliverable if it fired.
