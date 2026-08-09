---
iter: 230
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-230 — do the corpus's commit shas exist?

## Step 0 — Re-survey

`TOK-08`: *census the mechanical classes; stop sampling them.* The user's redirect ranks **the corpus's
claims about the platform** first. A cited commit sha is the most mechanical claim in the corpus — it either
resolves in the repo the prose attributes it to, or it does not. **No sentence has to be interpreted.**

Nothing has ever censused them. The instruments this milestone has built grade *anchors* (`file:line`) and
*predicates*; `anchor_construct_guard` resolves anchors **at** a ref but never asks whether the ref itself
is real. iter-228 found the corpus's most-cited **current** ref was 28 commits stale; this asks the prior
question of the whole population.

**Instrument feasibility confirmed before sealing** (not a measurement): 15 clones are on disk across
`stack-demo/` and `stack-dev/`.

**Active strategy reference:** `TOK-08`.

## Hypothesis

Some cited shas do not resolve anywhere. A sha that resolves in **no** available clone is either a corpus
error or an uncloned repo — and the census must distinguish those two, per `§9`'s rule that **a clone too
shallow to answer must report UNMEASURED**, never false.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-230-1` | ≥ 150 distinct backticked shas (7–40 hex) are cited across `corpus/**` + `CLAUDE.md` |
| `P-230-2` | ≥ 1 cited sha resolves in **no** available clone |
| `P-230-3` | `app` is the repo the largest number of distinct shas resolve in |
| `P-230-4` | ≥ 1 cited **short** sha resolves in **two or more** different repos — so the sha alone does not identify its repo |

## Expected lift

No `N`/`P` reading. Deliverable: the enumerated population, the resolvable/unresolvable split with its
denominator, and repair of anything found false.

## Escalation conditions

- A sha that fails to resolve **because its repo is not cloned** is `UNMEASURED`, not an error. If the
  unmeasured share dominates, that is the finding and the census must say so rather than publish a rate.
- Repairing prose is in scope; building a standing fence is a **second** deliverable and fires the
  tripwire — route it forward.

## Acceptable close-no-lift outcomes

If `P-230-2` is REFUTED — every cited sha resolves — that is a clean census result and the iter's
deliverable, provided the instrument proves itself (`§9`: *a census returning ZERO must prove its
instrument*) with a known-bad control.
