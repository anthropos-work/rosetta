**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.*

# iter-220 — 32 READMEs cite 97 files and nothing had ever checked that they exist

## The finding

A README is a registry (`§5` iter-184). `rosetta-extensions` has **32** of them, citing **97** distinct
`.py`/`.sh` filenames between them, and **both** directions of that registry were unfenced:

- **Direction B — a cited file exists.** Mechanically decidable, cheap, checked by nothing, in either
  repo. **Landed here.**
- **Direction A — a file that ought to be cited is cited.** `SURVEY-M257x-iter175`'s side, with a
  denominator iter-179 explicitly rejected. **Sized, printed, routed — deliberately not asserted.**

This run supplied the motive: iters 217–219 added **three** test modules and one diagnostic, and none is
referenced anywhere outside itself.

## ⚠️ The instrument's first run produced a false RED, and it is kept as the control

| pool the citation is checked against | dangling names |
|---|---|
| `stack-core/` only | **1** — `exposure_claim_guard.py` |
| the whole repo | **0** |

`exposure_claim_guard.py` is real and lives in **`stack-injection/`**; `stack-core/README.md` cites it
because the fence family spans sections. **The RED belonged to the probe, not to the repo** — iters 209
and 214's class, committed again by the session that has now written it down four times.

It ships as `test_02`, because a census whose only result is a **zero** must prove its instrument in the
same run (`§9`), and the cheapest honest proof available here is *the scope at which it fires*. The arm
asserts **both** halves: that the narrow scope produces a finding (the matcher is not inert) **and**
that the file it names is real (the widening was a correction, not a loosening). A second, independent
fire-proof staged a dangling citation into `stack-core/README.md` — detected, then restored with
`git status` verified clean.

## The mirror direction, sized and NOT asserted

**12 of 76** `stack-core/tests` modules carry a README row. The arm asserts only the **shape** — the
index is a strict subset, and the gap is non-empty — and **prints** the size. Pinning it would pin the
denominator iter-179 measured as wrong (*"the `10 of 63` you get from all test modules"*). *A correct
exclusion is still a defect while it is silent*, and its twin: **a number pinned before its denominator
is justified is worse than no number.**

## Scope, stated rather than implied (`§5` r60)

`/usr/bin/python3 -m pytest` (**pytest 8.4.2 / CPython 3.9.6**), **Python**, `stack-core` only,
changed-code reach: **130 passed / 0 failed** across the new module plus both fence-registry modules and
the frozen-expectation census (35 s). `derivation_registry --ceilings` exits **0**, all three
`exact +0` — **no re-pin needed this iter**, unlike each of the three before it. No whole-section run.
No Go, no TypeScript, no non-`stack-core` Python section.

## Close — 2026-08-09

**Outcome:** a registry class nobody had fenced in either repo now has a running enumeration — every
file citation in every section README, **97 across 32 documents, 0 dangling** — with the instrument
proven to fire two independent ways, and the mirror direction sized, declared and routed rather than
left silent or pinned to a rejected denominator.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — **iter-219 was a `closed-no-lift`; this
one is not, so the streak resets at one and the trigger needs three** — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n — **counted, not felt: iters 217, 218, 219, 220 = four tiks
this run against a cap of five** — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-220-1` … `D-M257x-220-3` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — **unchanged and now SIZED at the other
  scope**: 12 of 76 in `stack-core/tests`. The denominator remains the open question and is why this
  iter did not land it.
- All routes from iters 207–219 unchanged, plus the standing queue.

**Lessons:**
- **Install the fence while it is green.** The value of direction B is not that it found something; it
  is that a dangling row now cannot appear unnoticed.
- **When a census returns zero, the scope at which it returns non-zero is the proof** — and if that
  scope was your own first mistake, keep it as the control instead of deleting it.
- **A number pinned before its denominator is justified is worse than no number.** Print the size,
  assert the shape.
