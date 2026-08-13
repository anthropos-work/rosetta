---
iter: 169
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-169 — close the hand-listed-stage class, and fence its population

## Step 0 — re-survey before targeting

`TOK-08`'s standing direction is *census the mechanical classes; stop sampling them.* iter-168's
close routed **one** named residual with a warning attached — `FIX-M257x-iter168-m255-battery-stages-across-sections`,
*"the last member of `FIX-M257x-iter111-staged-battery-dependency-is-underived`, open since iter-111 and
re-encountered four times since… widening the helper is a change with its own failure modes and is not a
blind swap."*

Re-survey (this iter, before targeting): the target is **still open and still meaningful**.
`stack-core/tests/test_m255_mutation_battery.py:64-81` still carries a 12-entry `_COPY_FILES` literal, and
`grep -l battery_stage stack-core/tests/*mutation_battery*.py` returns 5 of the 7 battery files. Nothing
absorbed it.

**And the re-survey found the sixth occurrence, already live.** Running the derivation against m255's own
seeds returns one `.py` the hand-list does not carry: **`stack-core/fence_provenance.py`**. That is the
same file iter-111 added by hand to a different battery. The class is not "five past occurrences"; it is
five past occurrences **and one standing**.

## Active strategy reference

`TOK-08` — *census the mechanical classes; stop sampling them.* A battery's stage list is mechanically
decidable (the imports either resolve or they do not), and the class's population is enumerable. This iter
does the census, not another sample.

## Cluster / target identified

The staged-dependency class, in full:

| member | today |
|---|---|
| `test_m257x_claim_twin_mutation_battery.py` | derived (iter-168) |
| `test_m257x_repair_reach_mutation_battery.py` | derived (iter-168) |
| `test_m257x_repair_postcondition_mutation_battery.py` | derived (iter-168) |
| `test_m257x_mechanical_fences_mutation_battery.py` | derived (iter-168) |
| `test_repair_leak_guard_mutation_battery.py` | derived (iter-166/168) |
| **`test_m255_mutation_battery.py`** | **hand-listed — this iter** |
| `test_m220_mutation_battery.py` | **claimed out of class by iter-168 — verify, do not inherit** |

iter-168 asserted the population was **six**. That number was never derived, and this iter grades it.

## Hypothesis

Three things, in order:

1. **The population claim is checkable and one member's exclusion is load-bearing.** m220 was excluded
   from the class without a recorded measurement. Either it stages a file SET (in class, and the class is
   seven with two open) or it does not (out of class, and the exclusion earns a reason).
2. **The helper widens by resolving an import against the importing file's OWN directory** — which is what
   Python does at runtime — rather than by widening the search root. Sibling-first resolution handles the
   cross-section case (`stack-injection/gen_injected_override.py` → `stack-injection/platform_topology.py`)
   and is a no-op for the single-section batteries already migrated.
3. **A class closed member-by-member re-opens on member seven.** The residual is not the last hand-list; it
   is that nothing stops the next battery from carrying one. The deliverable is a fence over the
   *population*, per iter-162's registry rule: fence the completeness, not the contents.

## Expected lift

Not an `N`/`P` reading (no reading is taken this iter — `§9`: the metric stays UNMEASURED, not unmoved).
The lift is class closure with a stated denominator: the battery population enumerated and graded, 6-or-7
of them derived, one exemption carrying a *measured* reason, and a fence that fails on a new hand-list.

## Phase plan

- **A — census the population.** Enumerate every mutation battery mechanically; grade each as
  derived / hand-listed / structurally-exempt. Settle m220 by reading what it stages, not by inheriting.
- **B — widen `battery_stage.local_deps` and migrate m255.** Sibling-relative import resolution + the
  named failure mode (a repo `.py` shadowing a stdlib name in the staged tree) refused rather than
  silently staged. Prove the derivation reproduces the hand-list AND adds `fence_provenance.py`.
- **C — fence the population.** Every battery either derives its stage set or is on an exemption list
  whose entries must *prove* they stage no file set. Mutation control + anti-vacuity control, per the
  standing `TOK-08` requirement.
- **D — verify.** m255 battery baseline GREEN before and after (the battery's own anti-theatre #1); the
  five already-migrated batteries unaffected; scoped suite green with an honest statement of what it did
  NOT cover (`§5` rule 60).

## Escalation conditions

- If widening the helper changes the derived set for **any already-migrated battery**, that is a
  regression in five green fences: stop, do not paper over, land the narrower form and route the rest.
- If the m255 baseline is RED *before* this iter's change, that is a pre-existing finding to be reported
  as such, never as this iter's breakage.

## Acceptable close-no-lift outcomes

- The population census returns **seven** with m220 in class and its migration not landable this iter —
  the census is the deliverable and the residual routes with a named handler.
- Sibling-first resolution proves insufficient for m255 (e.g. an import that only resolves through a
  `sys.path` mutation) — record the falsification, keep the hand-list, and say what the helper cannot do.
