---
iter: 72
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
---

# iter-72 — the mainline class dissolves, and the guard cannot see 142 citations

**Active strategy reference:** `TOK-05`, step 2 (citations), clearing the last **named** unrepaired
class before step 4. The briefing is explicit: *do not take the graded read while a known class is
unrepaired* — so the remaining routes have to be settled or shown never to have been classes.

## Step 0 — re-survey before targeting

`FIX-M257x-iter58-mainline-shift` has stood at **"21 of 22 outstanding"** since iter-59. iters 69–71
changed the rule underneath it three times (`D-M257x-69-1` the pin rule, `D-M257x-71-1` per-block
refs), so the number cannot be inherited.

## Cluster / target identified

Re-derive the mainline class under the pin rule and settle the route. Whatever the guards **cannot
see** in that class is the more interesting half — iters 68, 69 and 71 each found a reach limit by
reading, and none of them was visible in a GREEN verdict.

## Hypothesis

The mainline class dissolves the way B2 did at iter-69: pinned citations are measurements, and the
structural residual is near zero. If so, the route closes and the read is not held behind it.

## Expected lift

- `FIX-M257x-iter58-mainline-shift` closed with a **derived** verdict rather than a carried number.
- Any reach limit found is **measured and named**, not asserted.

## Phase plan

- **A** — re-derive the mainline class under the pin rule.
- **B** — probe what the guard's own regex and resolver can reach in that class, mechanically.
- **C** — settle the route; design the fix for whatever gap is found; route it with a handler.
- **D** — gates.

## Escalation conditions

A reach gap large enough that closing it is its own build is **routed with a designed fix**, not
half-landed — the scope-creep tripwire, and the same disposition iter-68 used when it measured two
boundary defects unreachable and recorded them rather than patching on speculation.

## Acceptable close-no-lift outcomes

The class dissolving entirely is the expected and complete outcome: the deliverable is then the
derivation plus the route's closure, not edits.
