---
iter: 182
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-182 — the third way a test file hides its tests: it needs a runner

**Type:** tik. **Active strategy: `TOK-08`** — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey before targeting (mandatory)

Re-surveyed at HEAD `20ea7cc` (iter-181's commit), rext `b69f008`, trees clean modulo the user's
`.claude/settings.json`.

`FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` — open since iter-170, and the
route states its own cost: *"converting them is a real change with its own failure modes, not a one-line
fix — and per iter-158 the conversion must be shown not to weaken what they assert."*

**Sized before committing to it**, because a two-member route is not necessarily two equal members:

| module | lines | tests | pytest surface |
|---|---|---|---|
| `test_gen_override_home_binds.py` | 165 | 16 | one `@fixture`, one `@parametrize`, one `monkeypatch` |
| `test_claim_census_guard.py` | 481 | 25 | `@fixture` **plus** `tmp_path`, `monkeypatch`, `capsys` |

**They are not the same job.** The first is a mechanical conversion; the second needs three builtin
shims re-implemented.

## Cluster / target identified

Convert the tractable one **completely**, and close the *class* with an enumeration so the other is an
open member **inside a running fence** rather than a line in a markdown route (§8 iter-176 / iter-179).

## Hypothesis

The conversion's risk is not correctness, it is **silent narrowing**: `@parametrize` with 5 cases
collects as 5 tests, and the obvious `subTest` translation collects as 1. That would take the module
16 → 12 while every assertion still ran — a count moving during a change whose whole contract is that
nothing moves. **Predicted and to be checked: five named methods, 16 before and 16 after, on both
runners.**

## Expected lift

No `P`/`N` reading; **no clause-5 movement claimed** (`§9`). Deliverables: one module runnable on the
modern interpreter with its assertions untouched, and a derived enumeration of modules that still need
pytest, each classified, ratcheted downward.

## Phase plan (planned multi-step shape — declared, per the scope-creep carve-out)

1. Size both members; record why they are not one job.
2. Convert the tractable one assertion-for-assertion; prove the count on **both** interpreters.
3. Census the class repo-wide; state the denominator.
4. Fence it where the sibling arms already live (no new module — the tax iters 178–181 declined).
5. Controls: a ratchet floor, a regression arm on the converted module, and a control proving the
   predicate is the DECORATOR and not the import.

## Escalation conditions

- If the conversion changes the collected count in either direction, **stop and route** — a conversion
  that moves a number is not a conversion.
- If the census finds the class is large, land the conversion only and route the fence.

## Acceptable close-no-lift outcomes

Measuring that the two members are different-sized jobs, and landing only the enumeration, is a complete
iter provided the sizing is written down.
