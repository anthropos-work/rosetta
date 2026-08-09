---
iter: 220
milestone: M257x
iteration_type: tik
status: closed-fixed
created: 2026-08-09
---

# iter-220 — every section README cites files, and nothing has ever checked that they exist

**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

`SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` is open: `test_fence_registry_population_m257x`
publishes the prose index's reach as a checked triple (**16 of 27** union · **16 of 26** census ·
**15 of 26** declaring) and calls it *its own blind spot* — the index is invisible to all three live
derivations. Re-surveyed at HEAD: **that is the mirror direction, and it has a denominator problem
iter-179 already named.** The *other* direction — does a cited file EXIST — has never been checked in
either repo, by anything.

This run supplies the motive: iters 217–219 added **three** test modules and one diagnostic, and none is
referenced anywhere outside itself.

## Cluster / target identified

Every `README.md` in `rosetta-extensions` — **32 of them**, citing **97** distinct `.py`/`.sh`
filenames. A README is a registry (`§5` iter-184) and this one has no fence in either direction.

## Hypothesis

Direction B (a cited file exists) is mechanically decidable, cheap, and unchecked. It is probably green
today, which is exactly when a fence is worth installing — the point is that nothing would notice when
it stops being.

## Expected lift

A permanent fence over the registry, with the instrument **proven to fire**, and the mirror direction
**sized and declared** rather than left silent.

## Phase plan

1. **Seal** the readings and the scope lesson this iter's own probe produced, before landing anything.
2. Ship the census as an arm, with an anti-vacuity control and a **scope** control.
3. Size the mirror direction, declare it, route its denominator problem.

## Escalation conditions

- If direction B is RED anywhere, adjudicate every hit by hand before landing — a cross-section
  reference is not a dangling one, and this iter's own first probe made exactly that mistake.

## Acceptable close-no-lift outcomes

A measured zero is the *expected* result and closes this **`closed-fixed`**, not `no-lift`, provided the
fence ships and is proven to fire (`§9`) — the deliverable is the enumeration that keeps running, not a
repair.
