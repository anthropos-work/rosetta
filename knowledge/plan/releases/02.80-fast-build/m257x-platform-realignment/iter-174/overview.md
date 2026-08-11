---
iter: 174
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-174 — a capability probe that fails OPEN disarms the check it guards

**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

iter-172 routed `SURVEY-M257x-iter172-two-preexisting-actionable-reds`. iter-173 closed one member
(forced by entanglement — its own change added a third unclassified derivation to the same test). **One
member remains, re-confirmed RED at open** in iter-173's whole-suite run:

```
FAILED tests/test_battery_stage.py::TestLocalDepsResolution::test_a_stdlib_shadow_is_refused_not_staged
E       AssertionError: RuntimeError not raised
```

iter-172 characterised it as *"pytest/3.9.6 only — rule-76 shaped"* and did not diagnose it. Re-surveyed:
the target is live, the characterisation is right, and **the cause makes it worse than a stale test.**

## Cluster / target identified

`stack-core/tests/battery_stage.py:98`:

```python
stdlib = getattr(sys, "stdlib_module_names", frozenset())
```

`sys.stdlib_module_names` **landed in Python 3.10**. The box's two interpreters (iter-170) are
`/usr/bin/python3` **3.9.6** — *the only one with pytest* — and `python3` **3.14.6**. So on the runner the
milestone actually uses, the expression evaluates to an **empty frozenset**, `mod in stdlib` is never
true, and the refusal at `:125` **can never fire**.

**This is not a failing test. It is a disarmed safety check, and the test is the only thing that noticed.**
The guard exists to refuse staging a repo file that would shadow a stdlib module inside the staged tree —
"a staged-only failure of precisely the baffling kind this helper exists to end", in its own words. On
3.9.6 it stages the shadow silently.

**The shape:** a `getattr(..., <empty default>)` capability probe turns *"this interpreter cannot tell me"*
into *"the answer is nothing"*. **A check that cannot check must not report OK** — the same failure
direction as `§9` iter-149 (a census returning zero must prove its instrument) and as the M236 green-gate
that aged a stale verdict as fresh west of UTC: **failing open, silently, on half the world.**

## Hypothesis

Deriving the stdlib set on 3.9 — rather than defaulting to empty — arms the refusal on both interpreters
and turns the RED green **for the right reason**. If the set cannot be derived at all, the helper must
**raise**, not proceed: refusing is the only answer that is not a false OK.

## Hazard census, with its denominator (`§8`, iter-168 — *measure the hazard, or "it exists elsewhere" is
only a mood*)

Every `getattr(x, "attr", <empty>)` in `rosetta-extensions`: **13 sites**. Classified by whether the
default decides a **verdict**:

- **1 of 13** — `battery_stage.py:98`, a capability probe on the interpreter whose empty default silently
  disarms a refusal. **The hazard.**
- **12 of 13** — attribute lookups on a test class, an AST node, or a module (`getattr(cls, "started",
  False)`, `getattr(node, "body", [])`, `getattr(module, "postcondition_sites", None)`). The default means
  *"not set"*, which is the true answer, and two of them are immediately asserted on rather than trusted.

So the class is **not** systemic, and this iter says so with the denominator rather than implying a sweep
it did not do.

## Expected lift

No `P`/`N` reading (`§9`: UNMEASURED, not unmoved). Deliverable: the refusal armed on both interpreters,
proven by running the same test under both; a control that would have caught the disarming; the hazard
census with its denominator; and `SURVEY-M257x-iter172-two-preexisting-actionable-reds` **fully closed**.

## Phase plan

1. Derive the stdlib set portably; refuse if it cannot be derived.
2. Controls: the derived set is non-empty and contains known members on **both** interpreters — the
   anti-vacuity assertion whose absence is exactly why this went unnoticed.
3. Run `test_battery_stage` under both runners; then the batteries that consume `local_deps`.
4. Protocol doc: the fail-open capability-probe rule.

## Escalation conditions

- If arming the refusal turns other batteries RED (a real shadow already staged somewhere), that is a
  finding to route, not to suppress by weakening the derivation — iter-158's rule that a proposed
  narrowing which grades broken checks green is a defect, not a fix.

## Acceptable close-no-lift outcomes

Demonstrating that the refusal cannot be armed portably without weakening it, with the falsification
recorded, is a complete iter.
