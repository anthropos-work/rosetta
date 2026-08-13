**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*
**Shape:** standard tik. Opened 2026-08-09 00:21, `rosetta` `5197f58`.

# iter-174 — the shadow refusal could not fire on the runner that runs it

## What this iter was handed

`SURVEY-M257x-iter172-two-preexisting-actionable-reds`, one member left after iter-173 closed the other:

```
FAILED tests/test_battery_stage.py::TestLocalDepsResolution::test_a_stdlib_shadow_is_refused_not_staged
E       AssertionError: RuntimeError not raised
```

iter-172 characterised it as *"pytest/3.9.6 only — rule-76 shaped"* and did not diagnose it.

## THE FINDING — it was not a failing test, it was a disarmed check

`stack-core/tests/battery_stage.py:98`:

```python
stdlib = getattr(sys, "stdlib_module_names", frozenset())
```

`sys.stdlib_module_names` **landed in Python 3.10.** The box's interpreters (iter-170): `/usr/bin/python3`
**3.9.6** — *the only one with pytest* — and `python3` **3.14.6**. So on the runner the milestone actually
uses, that set is **empty**, `mod in stdlib` is never true, and the refusal at `:125` **cannot fire at any
input**.

The refusal exists to stop staging a repo file that would shadow a stdlib module inside a mutation
battery's staged tree — *"a staged-only failure of precisely the baffling kind this helper exists to
end"*, its own words. **On 3.9.6 it staged the shadow silently.**

**The shape: a capability probe that fails OPEN.** `getattr(x, "cap", <empty>)` turns *"this interpreter
cannot tell me"* into *"the answer is nothing"*. The two directions are not symmetric — an over-broad
stdlib set refuses a legitimate file **loudly**; an empty one permits a shadow **silently**.

### Repair, measured on both interpreters

`stdlib_module_names()` — the native attribute where it exists, a derivation from `sys.builtin_module_names`
plus the `sysconfig` stdlib directory where it does not, and a **`RuntimeError` if neither is available**.
Never an empty default.

| interpreter | names derived | `json` `os` `sys` `re` `sysconfig` | any repo module? |
|---|---|---|---|
| 3.9.6 (pytest runner) | **232** | all present | none |
| 3.14.6 | **297** | all present | none |

## Hazard census, with its denominator (`§8`, iter-168)

Every `getattr(x, "attr", <empty default>)` in `rosetta-extensions`: **13 sites**, classified by whether
the default decides a **verdict**.

| | count | |
|---|---|---|
| capability probe whose empty default disarms a check | **1** | `battery_stage.py:98` — the hazard |
| attribute lookups where the default means *"not set"*, which is the true answer | **12** | `getattr(cls, "started", False)`, `getattr(node, "body", [])`, `getattr(module, "postcondition_sites", None)` — two of them immediately asserted on rather than trusted |
| **total** | **13** | |

**The class is not systemic**, and this iter says so with the denominator instead of implying a sweep it
did not do. `sys.stdlib_module_names` has exactly **one** reader in the repository, now routed through the
derivation.

## The control — it asserts the SET, not the refusal

The old test's failure message was `RuntimeError not raised`: it reads as *"the refusal is broken"* and is
one inference away from *"the set it consults is empty."* That gap is why iter-172 got as far as the
runner and stopped. Four net-new controls assert the **set** directly (non-empty · contains what the
refusal must catch · contains no repo module · equals the native attribute where one exists).

**Mutation-proven.** Restoring the shipped `getattr(..., frozenset())` in place kills **8 assertions** —
the original refusal test *and* the anti-vacuity control. The cause is now named at the point of failure.

## THE MEASUREMENTS — counts, not wall-time

**Both runners** (`§5`, iter-170 — *name the runner*; `§5`, iter-172 — *name the unit*):

| runner | result | executed |
|---|---|---|
| pytest / 3.9.6 | `16 passed, 1 skipped` | **17** |
| unittest / 3.14.6 | `Ran 17 tests · OK` | **17** |

The skip is the native-attribute cross-check, correctly inapplicable on 3.9.6 — **and the two columns
agree only because the unit was named**: `16 passed` and `Ran 17` are the same run.

**Whole `stack-core` suite, stable tree** — `4 failed · 1526 passed · 2 skipped in 853.21 s (0:14:13)`.

⚠️ **The pre-registered escalation condition fired, and it was NOT what it looked like** (`D-M257x-174-3`).
This iter's `overview.md` had pre-registered *"if arming the refusal turns other batteries RED, route it,
do not weaken the derivation."* Four failures landed in the mechanical-fences mutation battery — exactly
that shape. **Read rather than assumed, none of them was the stdlib set.** All four are one cause:
`derived_count_guard.py` — **iter-173's** fence — missing from the battery's fence-seed list, so the staged
tree's baseline named a fence the tree did not contain. **A pre-registration names a symptom; it cannot
name a cause.**

**The fifth registry** (iter-173 enumerated four and found the fourth only by running the suite): the
battery's explicit seed list. `repair_postcondition.py --accept` wrote the fence into the *baseline* and
nothing wrote it here. **Added by name, deliberately NOT derived from the baseline** — `test_000` asserts
*staged ⊇ the baseline's fence names*, so seeding from that baseline would compare a set with itself
(iter-158). Post-repair: **22 passed, 1 skipped** across the battery + the stage-helper suite.

**And the timing is the finding.** iter-173's own post-fix scoped re-run was **167 passed** and green. It
could not see the battery it had just broken. *That is iter-173's own lesson, firing on iter-173's own
commit, one iter later.*

## Close — 2026-08-09

**Outcome:** The last open member of `SURVEY-M257x-iter172-two-preexisting-actionable-reds` was **not a
stale test but a disarmed safety check**: `getattr(sys, "stdlib_module_names", frozenset())` is empty on
Python 3.9.6, which is *the only interpreter here with pytest*, so the staged-tree shadow refusal could
not fire at any input. Repaired by deriving the set (232 names on 3.9.6, 297 on 3.14.6) and **raising**
rather than defaulting when it cannot be derived; four controls now assert the **set** rather than the
refusal, mutation-proven to kill the shipped form with 8 assertions. Hazard censused at **1 of 13**
`getattr`-default sites. The whole-suite run then surfaced a **fifth registry** and a regression from
iter-173's own commit, both repaired.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (sixth consecutive `closed-fixed`; **no `P`/`N`
reading taken, so the metric is UNMEASURED, not unmoved** — `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN
ITERS, tree clean)
**Decisions:** `D-M257x-174-1` … `D-M257x-174-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none. The battery seed-list repair is not a side fix — it is a regression from this
milestone's immediately preceding commit, surfaced by this iter's own suite run.

**Routes carried forward:**
- `SURVEY-M257x-iter172-two-preexisting-actionable-reds` — **CLOSED.** Both members resolved (the
  derivation-registry half at iter-173, this half here).
- `FIX-M257x-iter174-accept-registers-one-registry-of-two` — **NEW.** `repair_postcondition --accept`
  writes a new fence into the postcondition baseline; a second registry (the mechanical-fences battery's
  seed list) must be hand-updated, and the only thing that reports the gap is a 14-minute battery, one
  iter later. Five registries are now known; nothing enumerates them.
- `FIX-M257x-iter173-ledger-denominator` — unchanged; open (owned by the next harden pass).
- The observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a capability probe that fails OPEN disarms the check it guards, and reports it as a test
failure somewhere else.** The symptom (`RuntimeError not raised`) named the refusal; the cause was the set
the refusal consults, one function away, and a whole iter's characterisation stopped at the runner because
nothing asserted the set.

Two that generalise:

1. **When a check depends on an interpreter capability, assert the CAPABILITY, not only the behaviour.**
   Otherwise the check's own test is the only witness, and it reports the wrong thing.
2. **A pre-registered escalation names a symptom, never a cause — diagnose before applying it.** This
   iter's escalation clause matched the observed failure shape exactly and was wrong about all four
   failures. Applied on its face, it would have weakened a correct derivation to silence an unrelated
   regression.
