---
iter: 138
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-138 — the rotted anchors, and whether a machine can find them

**Type:** tik
**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* **Class 1
is intra-corpus citation resolution**, `TOK-08`'s own named largest class (10 of 37 at iter-116). This
tik works the citation half of `FIX-M257x-iter135-adjudicated-live-defects` and then asks the question
`TOK-08` actually demands — *can this class be CENSUSED rather than sampled?*

## Step 0 — Re-survey before targeting

Re-surveyed at HEAD (`0a4a43a`, post-iter-137). All still live:

- `shared_libraries.md:77` — the `analytics-go` row still cites `main.go:507-508`.
- `security_compliance.md:156` — still cites `clerk-integration.md:40`.
- `sentinel.md:5` · `clerk-integration.md:126` · `backend.md:13` — live.
- `adj-E`'s five rotted anchors (`academy-backend.md` ×2, `graphql-wundergraph.md` ×3) — live.

Target current. No substitution.

## Cluster / target identified

`adj-E` established the fact that changes what the repair *is*: **every one of the five rotted anchors
was CORRECT when written**, and was invalidated by a later, unrelated edit inserting lines **above** the
target (+2 / +3 / +8 / +14). Verified by the adjudicator at each authoring commit.

> That decides the remedy. *"Repair harder"* assumes a careless author; **there was none.** The class is
> a **form** problem, and `corpus_citation_guard.py`'s own docstring **declares this exact blind spot** —
> bare `:NN` pins are *"excluded outright"* as not mechanically decidable.

## Hypothesis

**The exclusion is broader than the evidence requires, and a sub-class IS mechanically decidable.**

A bare `:NN` pin's *claim* cannot be checked without reading the sentence — granted. But **rot** is a
different predicate, and it is checkable from git alone: if a pin at commit `C` targeted content `X`, and
`X` now lives at a different line in the same file, the pin **rotted** and the new line is the repair.
No sentence has to be interpreted. That is exactly the `TOK-08` shape — *a reading samples; a fence
censuses*.

**Expected lift:** no `N` movement claimed (no reading this iter). Deliverables: the enumerated citation
defects repaired, and the rot question **answered with a number** rather than assumed either way.

## Phase plan (declared 2-step shape — the tripwire counts unplanned lines, not these)

1. **Priority 1 — repair the adjudicated citation set.** Width first (§5 rule 57), each target re-derived
   by opening it, never by trusting the adjudicator's line (`D-M257x-136-1`).
2. **Priority 2 — the rot probe, bounded.** One script: for every same-file bare `:NN` pin in
   `corpus/**`, resolve the target at the pin's authoring commit and at HEAD, and count how many moved.
   **Pre-registered branch, stated before the probe runs:**
   - **≥ 5 rotted pins found** → the class is real and enumerable; **route the fence build** with the
     measured population as its denominator.
   - **≤ 4** → `adj-E`'s five are most of it; **say so, do not build a fence for a population of five**,
     and record the refutation.

## Escalation conditions

- A 3rd unplanned line of investigation → tripwire; land Priority 1, route the rest.
- If the probe needs more than a bounded script to answer, it is **routed, not extended** — eight vacuous
  fences on this milestone's record all came from building under pressure (`D-M257x-134`).

## Acceptable close-no-lift outcomes

The probe returning **≤ 4** is a first-class result: it would refute the fence's value and close
`FIX-M257x-iter135-bare-pin-blind-spot` as *measured and not worth building*, which is a real outcome.
