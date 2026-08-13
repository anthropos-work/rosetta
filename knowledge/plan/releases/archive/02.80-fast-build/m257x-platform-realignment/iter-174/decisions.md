# iter-174 — decisions

## `D-M257x-174-1` — a capability probe that FAILS OPEN is worse than a missing check

`battery_stage.py:98` read the standard-library module set as:

```python
stdlib = getattr(sys, "stdlib_module_names", frozenset())
```

`sys.stdlib_module_names` **landed in Python 3.10**. This box has two interpreters (iter-170):
`/usr/bin/python3` **3.9.6** — *the only one with pytest, i.e. the runner the milestone actually uses* —
and `python3` **3.14.6**. On 3.9.6 that expression is an **empty frozenset**, so `mod in stdlib` is never
true and the shadow refusal below it **cannot fire at any input**.

**It was reported as a failing test and it was a disarmed safety check.** The guard exists to refuse
staging a repo file that would shadow a stdlib module inside a staged tree — *"a staged-only failure of
precisely the baffling kind this helper exists to end"*, in its own docstring. On the runner that runs it,
it staged the shadow silently.

**Decision: derive the set where the attribute is absent, and RAISE where neither is available.** Never
default to empty. The two failure directions are not symmetric: an over-broad stdlib set refuses a
legitimate file **loudly**; an empty one permits a shadow **silently**. Measured after the repair: **232
names on 3.9.6, 297 on 3.14.6**, both containing `json`/`os`/`sys`/`re`/`sysconfig`/`typing` and neither
containing any repo module.

This is the same failure direction as `§9` iter-149 (*a census returning zero must prove its instrument*)
and as M236's green-gate that parsed a UTC timestamp as local and therefore **aged a stale verdict as
fresh west of UTC**. The general form: **a check that cannot check must not report OK.**

## `D-M257x-174-2` — the control asserts the SET, not the refusal

The old test exercised only the refusal, so its failure message was `RuntimeError not raised` — which
reads as *"the refusal is broken"* and is one inference away from *"the set it consults is empty."*
iter-172's characterisation (*"pytest/3.9.6 only, rule-76 shaped"*) got as far as the runner and stopped,
because nothing in the failure named the set.

**Decision: four net-new controls that assert the derived set directly** — non-empty (>100), contains the
modules the refusal must catch, contains no repo module (the other direction, since a set derived by
listing a directory could pick up anything), and — on an interpreter that HAS the native attribute —
equals it exactly, which is the guarantee that the fast path was not quietly replaced by the fallback.

**Mutation-proven, not asserted.** Restoring the shipped `getattr(..., frozenset())` in place kills **8
assertions**: the original refusal test (the symptom) *and* the anti-vacuity control (the cause). A future
reader gets *"the stdlib set is empty"* instead of *"RuntimeError not raised."*

## `D-M257x-174-3` — the escalation condition fired, and it was NOT what it looked like

iter-174's `overview.md` pre-registered: *"if arming the refusal turns other batteries RED (a real shadow
already staged somewhere), that is a finding to route, not to suppress by weakening the derivation."*

The whole-suite run came back **4 failed · 1,526 passed** with the mechanical-fences mutation battery RED,
which is exactly the shape that condition describes. **It was not that.** Read rather than assumed, the
failure is `derived_count_guard.py` — **iter-173's** fence — missing from the battery's seed list. Nothing
to do with the stdlib set.

**Decision: diagnose before applying the pre-registered response.** A pre-registered escalation names a
*symptom*; it cannot name the cause. Had the condition been applied on its face, this iter would have
weakened a correct derivation to make an unrelated regression go away — iter-169's rule, one turn on:
**a route predicts a cause; it does not certify one**, and neither does a pre-registration.

## `D-M257x-174-4` — `--accept` registers a fence in ONE registry; there is a second, and only a 14-minute battery reports the gap

iter-173 enumerated four registries a new fence must join and found the fourth only by running the whole
suite. **There is a fifth**: the mechanical-fences battery's explicit fence-seed list. iter-173's
`repair_postcondition.py --accept` wrote the new fence into the *baseline*, and the battery then staged a
tree whose baseline named a fence the tree did not contain — a **RED BASELINE**, which the battery's own
`test_000` reports precisely and correctly.

**Decision: add the name; do NOT derive the seed list from the baseline.** Deriving it would make
`test_000`'s assertion — *staged ⊇ the baseline's fence names* — compare a set with itself. That is
iter-158's rule (a narrowing that grades a broken check green is a defect, not a fix), and the battery's
own comment had already reasoned it out. **The recurrence is real and the fix for it is not here**: it is
that `--accept` writes one registry while a second must be hand-updated. Routed as
`FIX-M257x-iter174-accept-registers-one-registry-of-two`.

**And the timing is the point.** iter-173's post-fix scoped re-run was **167 passed**, green, and could not
see this. It took this iter's whole-suite run to surface it — *the same lesson iter-173 itself recorded,
firing on iter-173's own commit one iter later.*
