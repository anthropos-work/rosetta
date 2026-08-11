---
iter: 181
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-181 — the question was unanswerable until its DENOMINATOR was named

**Type:** tik. **Active strategy: `TOK-08`** — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey before targeting (mandatory)

Re-surveyed at HEAD `3a8f5b4` (iter-180's commit), rext `c54c733`, trees clean modulo the user's
`.claude/settings.json`.

| route | state |
|---|---|
| `SURVEY-M257x-iter179-readme-indexes-test-modules-unmeasured` | **open, opened one iter ago** — the disclosed `16 of 27 / 16 of 26 / 15 of 26` triple measures the README's coverage of **fences**; its coverage of **test modules** has never been measured, and iter-179 found a three-iter-old fence with no row. |
| `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) | open — **probed first and set aside**: 412 count-mentions across 115 files, but the defect is only where a `passed` count is used *as the executed population*, which is a semantic distinction. iter-173 already priced this as out of reach without re-running past refs. Recorded so the set-aside is a measurement, not a mood. |

**Target: the README-index census** — bounded, mechanical, and one iter old with fresh evidence.

## Cluster / target identified — and the first measurement REFUTES the question's shape

Measured before anything was written:

| denominator | reading |
|---|---|
| all `tests/test_*.py` on disk | **10 of 63 indexed** |
| **mutation batteries** (`*mutation_battery*.py`) | **6 of 7 indexed** |

**`10 of 63` is the wrong ratio and publishing it would be a defect**, not a finding. The index's subject
is the fence family and its batteries; the other 53 are per-guard *behaviour* suites, which the index
deliberately does not list — it lists the guard. A survey asking *"how well does the README cover test
modules?"* has no answer until it says which test modules it means.

The right denominator yields a real gap: **`test_repair_leak_guard_mutation_battery.py` has no row**, and
nothing would ever have said so.

## Hypothesis

The gap is not the point; the *absence of an enumeration* is. Fence it in both directions — every battery
on disk has a row, and every file the README names exists **somewhere in rext** (the second direction is
non-trivial: `exposure_claim_guard.py` is named here and lives in `stack-injection/`, so a check scoped to
`stack-core/` would report a false missing — which is exactly the naive instrument this iter ran first and
had to discard).

## Expected lift

No `P`/`N` reading; **no clause-5 movement claimed** (`§9`). Deliverable: the missing row, the fence, and
the denominator question answered in writing so the survey closes rather than being re-asked.

## Phase plan

1. Measure both denominators; record the refutation of the survey's implied one. *(done above)*
2. Add the missing battery row.
3. Fence: batteries ⊆ indexed (both directions) + every README-named `.py` resolves within rext.
4. Controls: anti-vacuity + a mutation control per arm.
5. Protocol doc §8 rule if it generalises.

## Escalation conditions

- If "which test modules does the index owe a row?" turns out to need a hand list, **land only the row +
  the resolve-arm** and route the rest — a fence that is itself a registry is the tax iters 178–180 all
  declined to pay.

## Acceptable close-no-lift outcomes

Publishing the correct denominator and the refutation of `10 of 63` — with the survey closed as
ill-posed-until-scoped — is a complete iter even if only one row lands.
