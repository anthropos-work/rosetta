# MC02 — Progress

**Status: `planned`.** Not started. No iters run.

> **This is an ITERATIVE checkpoint milestone — there is no work checklist here.** MC02 builds nothing of
> its own; it **measures**. Progress is the **running ledger** below: what was measured, on which stack,
> against which artifact, and where each RED was routed.
>
> **A RED is a successful measurement.** The admissible outcomes are **PASS** and **work sent back to
> M267 / M269 / M270**. Sending work back is this milestone doing its job. Re-wording a clause so it goes
> green is this milestone failing at it.

## Exit-gate ledger

Each clause is graded **per stack**. **A clause that is PASS on one stack and RED on the other is RED.**
The milestone passes only when all six read PASS on the same pair of stacks in the same pass.

| clause | demo stack | dev stack |
|---|---|---|
| 1 — hero STARTS **and COMPLETES** in chat, code, document; real session rows, before/after counts | _not measured_ | _not measured_ |
| 2 — no hero on **any** seeded org hits `ERROR_JOB_SIMULATION_LIMIT_REACHED`; one start attempted per org | _not measured_ | _not measured_ |
| 3 — `/library/skill-paths` paints a loading affordance then content; browser, cold, cache disabled, time stated | _not measured_ | _not measured_ |
| 4 — batch gate **RUNS and is GREEN** on the **DEFAULT** `/demo-up` path; `skipped` absent from the verdict | _not measured_ | _open — see the dev-side question_ |
| 5 — every clause proven on **both** stacks | _not reached_ | _not reached_ |
| 6 — `playthroughs.md`, `verification.md`, `demopatch-spec.md`, `seeding-spec.md` match the **running** stacks | _not read_ | _not read_ |

**Clause 4 is the one most likely to fail** — `BIND_HOST` / `D-M255-7` has been deferred **three times**
(M255 → M256 → M258, never worked). If it cannot be met because `BIND_HOST` proves larger than M269
scoped it, **the re-scope trigger fires: escalate.** Do not accept a `skipped` gate.

## Running ledger

_Iter closeouts append here, newest last. One entry per iter: what was measured, on which stack, from
which artifact, what came back RED, and where each RED was routed._

_No entries yet._

## Iters

_None._
