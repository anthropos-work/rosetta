---
milestone: M257x
iter: 38
iteration_type: tik
status: in-progress
opened: 2026-08-02
---

# iter-38 — `MEASURE-M257x-iter36-clause5-fourth-pass`

**Active strategy:** `TOK-01`. The last outstanding gate clause. Clauses 1, 2, 3 and 4 hold as of iter-37.

## Step 0 — re-survey

- `git diff --name-only 5e37bb1 2c9befe -- corpus/` returns **9 files**, exactly as the hand-off said —
  re-derived, not trusted. **8 of the 9 are in clause-5 scope** (`corpus/services/**` +
  `corpus/architecture/**`); the 9th, `corpus/ops/platform-alignment.md`, is not.
- `git diff --name-only 2c9befe HEAD -- corpus/` returns **nothing**: iters 35–37 changed no corpus file,
  so the corpus under audit is byte-identical to what iter-34 left. The measurement is therefore about the
  same text, and nothing since has perturbed it.
- Scope: **40 files / 8 624 lines**; the 8 repaired ones are **2 686 lines (31 %)**.
- Platform origin re-fetched at open: `2adcf71`, unchanged.

## The design decision this iter had to make, and why it went the way it did

The routed item said *"scope the next pass to the 9 changed files."* §5 rule 18 supports that: iter-34
measured a **~9×** density difference (0.69 blockers/file in the 13 repaired vs 0.074 in the 27 untouched).

**It is being run WIDE anyway, and the reason is the clause's own wording.** Clause 5 asks for an audit
*"over `corpus/services/**` + `corpus/architecture/**`"*. A pass that reads 8 of 40 files cannot return a
verdict about 40, however well-targeted it is — and iter-21 is this milestone's own precedent for a scoped
audit converging on a number that a full read then multiplied by five. What rule 18 licenses is
**weighting**, not **narrowing**: the repaired files get a dedicated auditor each (or a pair), and the
other 32 get a third full read at lower expected yield. Every in-scope file is read top-to-bottom.

**Re-partitioned, per rule 18(b).** iter-33 and iter-34 both used five auditors over the same 40 files.
This pass uses **six**, with different boundaries, and the split is by *risk* rather than by directory:
`security_compliance.md` + `architecture_overview.md` get an auditor to themselves (the multi-tenancy
fence between them has now been wrong **four times, in both directions**, each time by verifying the
denominator of a conjunction and not the predicate — §5 rule 17). No auditor inherits a prior boundary.

## PRE-REGISTERED PREDICTION (written before any auditor reported)

**2–5 blockers total; at least 3 of them inside the 8 repaired files; 0–1 across the 32 untouched.**
Therefore **clause 5 is expected NOT to be met by this pass** — a fourth pass that fixes what it finds is
not a zero-blocker reading, and iters 33/34 both refused to claim the clause on that basis.

If the pass returns **zero**, clause 5 is met and the gate is 5 of 5.

## Escalation conditions

- A blocker requiring a platform edit to resolve → STOP (binding constraint).
- A count that reproduces but whose claim is a conjunction → verify every conjunct before editing (rule 17).

## Acceptable close-no-lift outcomes

A measured non-zero blocker count, enumerated with exact anchors and fixed, is a complete iter even though
the clause stays open — the clause is met by a **reading**, not by a repair.
