---
iter: 197
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-197 — the 25-test runner gap was reconciled in aggregate and never read per module

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*
*"A reading SAMPLES; a fence CENSUSES."*

## Step 0 — re-survey (mandatory)

`TOK-08` names no specific next target; the controlling backlog is the route registry
(`route_disposition_guard`, 379 ids / 1,316 dispositions / 0 contradictions at iter-194). Re-surveyed the
open set at HEAD `1e0619fe`. The four harden-filed routes from passes 45–47 are the newest, and one of
them is the milestone's only open **`FIX-`** route:

> `FIX-M257x-h44-claim-census-guard-is-single-runner` — *"Convert `test_claim_census_guard.py`'s 25
> pytest-style functions to `TestCase` so both runners collect them… Declared and fenced meanwhile."*
> Pass 45 recorded its live size as *"the 3,502/3,527 baseline gap above… and it reconciles."*

**That sentence is the target, not the conversion.** The evidence that the gap *is* this one module is an
**aggregate arithmetic reconciliation** — 3,527 − 3,502 = 25, and this module has 25 tests. Two modules
with offsetting differences reconcile to the same 25. This milestone already has a name for that shape:
**an agreeing reconstruction is indistinguishable from a reading** (iter-192), and **grade at the grain of
the claim** (`§9`).

Confirmed by reading `suite_census.py` before planning: the both-runner comparison at `:736-742` keys on
`r["verdict"]` only — `GREEN`/`RED`/`ENV-GATED`/`TIMEOUT`. The per-runner **`tests=`** figure at `:733` is
a `sum()` over the whole population. **There is no per-module count comparison anywhere**, so the claim
"the gap is exactly that module" has never been measured at module grain.

## Cluster / target identified

The mechanical class: **test modules whose collected-test count differs between the repo's two runners.**
Mechanically decidable per module (a runner either collects a test or it does not — no sentence needs
interpreting), and censusable cheaply at **collection** grain rather than execution grain.

## Hypothesis

A per-module collection census will either (a) confirm h44's reconstruction as a reading — one member,
`stack-core/tests/test_claim_census_guard.py` — or (b) show the aggregate hid offsetting members. Either
way the class stops being sampled. **The unit is COLLECTED tests, not executed tests** — a different unit
from the `tests=` column (whose unit iter-172 settled as *executed*), and it must be named as such.

## Expected lift

No `P`/`N` reading this iter (clause 5 is read by the adjudicated reading protocol, not by a fence).
The deliverable is a mechanical-class census at zero-or-declared with controls that can fire, plus the
grain correction to h44's evidence.

## Phase plan

1. Add a per-module **collection** census to `suite_census.py` (both runners, no execution), reporting
   every module where the two disagree, and a repo-wide **style** reading (`TestCase` vs bare `def test_`).
2. Run it over all 122 modules. Report the population and **state the denominator**.
3. Fence it: declared-vs-derived on the disagreeing set, with a **mutation control** (`§5` r77 — mtime
   bump, not `__pycache__` removal) and an **anti-vacuity control** that fires on an empty population.
4. Re-state h44 with its evidence corrected from reconstruction to reading.

## Escalation conditions

- Census finds >1 member → the class is larger than h44 declares; report the number, do not widen the
  iter into converting all of them.
- Census cost exceeds ~2 min → drop to a declared subset and say so.

## Acceptable close-no-lift outcomes

The census returning **exactly the declared single member** is a first-class outcome *provided the
instrument is proven able to return more* — a zero (or a one) that cannot move is not a measurement.

## Explicitly NOT in scope

The **conversion itself** (`FIX-M257x-h44-…`). It stays open and is not re-routed by this iter; this iter
gives it a fence and a correctly-grained size. Keeping the two apart is the same separation pass 45 drew
between a population and a verdict.
