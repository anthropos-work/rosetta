**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.*

# iter-217 — the three censuses could not read this repo's own emphasis idiom

## The finding

`derivation_registry._MEASURED_RE` is the ONE matcher all three literal censuses share
(`docstring_measurement_literals`, `comment_measurement_literals`, `excluded_test_module_literals`).
Its separator between the operand and the measured noun was `[\s\-]+`, and its `of`-clause required
plain whitespace after the second operand — so **`**292 of 704** modules`, which is how this repo writes
every figure it considers important, matched nothing at all.**

**Third hole in one regex, one character-class over each time:** iter-205 (the vocabulary was written
lower-case and the regex was not) · harden pass 48 (the separator took whitespace only, so every
attributive compound was invisible) · **this iter** (the separator took unemphasised text only). Each
was invisible to **all three censuses at once** — which is what a shared matcher buys, and what it
costs. The route has been open and re-listed unchanged since pass 48; pass 49 recorded it as *"now the
only open reach hole of its family."*

## Measured — V1…V6 sealed in this iter's FIRST commit (`cac1612`), before any repair

Every reading below holds the shipped machinery fixed (`_MEASURED_NOUNS`, `_measurement_units`,
`_classify_measurement`, `_unit_line`, `_CENSUS_SKIP`) so **the separator is the only variable** —
iter-209's discipline, and the reason iter-209's own hand-written slugger came back 16× wrong.

| | |
|---|---|
| baseline, all three ratchets at `3965790` | **195 · 159 · 492 — `exact +0`**, so nothing below is inherited slack |
| close leg (`**292 of 704**`) — new matches | **39** · 18 non-test · 21 test · 24 `standing` |
| mirror leg (`292 **modules**`) | **0** — shipped anyway, fenced by a **staged** control |
| backtick variant | **+19 further**, every one a code-span tail → **REFUSED** (`D-M257x-217-2`) |
| declared false positive | **1 of 39** — *"only **4** name exactly one corpus doc"*, where `name` is a verb |

**V6 reconciled EXACTLY, and the reconciliation is the proof the delta is a reach fix and not growth.**
Isolated by `git archive`-ing the pre-repair tree and censusing it with the **post-repair** registry —
tree frozen, matcher variable: **12 + 6 + 21 = 39**, with the test leg taking all 21 exactly as
pre-registered. The one admissible shortfall (row de-duplication) did not fire.

## The second finding — a residual is only as wide as the matcher that produces it

`_ANY_NOUN_RE` is the **superset** the vocabulary's reach is measured against. Widening only the
selector would have left it narrower than the set it bounds, so both were widened in one edit
(`D-M257x-217-4`). The consequence arrived immediately: the residual arm — which had reported a **clean
vocabulary, 0 addressable words**, for as long as the hole was open — surfaced **five** words the moment
the separator widened, and **every one of them is bold-emphasised at every site**:

> `**384** refusals` · `**114** sources` · `**0** directories` · `**14 of 1,447** segments` ·
> *"exactly **1** crosses a pruned name"*

**The vocabulary was never what hid them.** Four taken as nouns; `crosses` declined as a verb — plural
shape is not nouniness, for the fifth time. This is the fifth consecutive instance of *the vocabulary's
reach closes on the sentence that widened it*, and the first where the constraint was proven to be one
construct **below** the list rather than in it.

## The ceilings, re-pinned with recorded reasons — and one of them has a FIXPOINT, not a value

| ratchet | before | after | attribution |
|---|---|---|---|
| `DOCSTRING_LITERAL_CEILING` | 195 | **209** | reach only; none of it this iter's prose (frozen-tree isolated) |
| `COMMENT_LITERAL_CEILING` | 159 | **175** | reach **+ this iter's own writing**, itemised below |
| `TEST_MODULE_LITERAL_CEILING` | 492 | **542** | 519 reach (frozen-tree isolated) **+ 23 this iter's own arms** |

`COMMENT_LITERAL_CEILING` was re-pinned **twice in one iter**, and that is the demonstration rather than
an accident: pinning at the first reading breached by exactly 2 on the next, because writing the
paragraph that recorded *why* quoted the two illustrations a second time. All **6** live-minus-frozen
comment rows are itemised rather than netted out — every one is this iter's own note quoting the very
sites the widening surfaced — and the two illustrations land as cleanly as a fixture could:

> `**292 of 704** modules` and `292 **modules**` are **each caught by the leg they illustrate, and by
> nothing else.**

*A provenance note joins the population it explains* — recorded for the third time in this module, and
demonstrated in both legs at once.

## ⚠️ The whole-section run caught a defect in this iter's own re-pin, and it is the harden's class

`TheCeilingProseDoesNotContradictTheCeiling` went RED: `COMMENT_LITERAL_CEILING`'s block handed off
**173** with an arrow while the constant read **175** — the fixpoint paragraph explained the second pin
in prose and left the **arrow** describing the first. That is precisely the shape harden passes 51→52
recorded (*a pass re-pinning a ratchet and the next pass breaking it*), committed **inside one iter this
time**, and the fence that was built for it caught it. Repaired by adding the second arrow rather than
by rewriting the first, so both pins stay readable.

## Scope, stated rather than implied (`§5` r60)

- **Whole-section:** `/usr/bin/python3 -m pytest stack-core/tests` (**pytest 8.4.2 / CPython 3.9.6**),
  **Python**, `stack-core` only — **1,864 passed · 1 failed · 3 skipped in 1,704 s (28 m 24 s)**. The
  single failure is the arrow defect above, **found by this run**, fixed after it.
- **Post-fix, scoped:** the three affected modules — **122 passed / 0 failed** (43 s), and
  `derivation_registry.py --ceilings` exits **0** with all three `exact +0`.
- **NOT re-run whole-section after the fix**, and the delta against the quoted 1,850/4/3 baseline is
  **not attributable** — that figure was taken on a different tree at a different HEAD. What *is*
  measurable here: all three ratchets read `exact +0` at HEAD **before** this iter's first edit (V1), so
  the four baseline failures were not these ratchets at this HEAD.
- **No Go. No TypeScript. No non-`stack-core` Python section.**

## Close — 2026-08-09

**Outcome:** the only open reach hole of the literal-census family is closed. The shared matcher now
reads this repo's own emphasis idiom, worth **39** previously-invisible matches (**12 + 6 + 21**,
reconciled on a frozen tree), which in turn exposed **five** vocabulary words a narrow separator had
been hiding while the residual arm reported a clean vocabulary. All three ratchets re-pinned with
recorded reasons and reading `exact +0`.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-ninth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt: iter 217 = one tik this run against a cap of five** —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-217-1` … `D-M257x-217-5` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-h48-the-censuses-cannot-see-a-bold-wrapped-operand` — **CLOSED.** Sized at 39 matches
  across all three censuses and repaired; the ceilings carry the arrows.
- `SURVEY-M257x-iter217-the-ratchets-have-no-pre-edit-whole-section-reading` — **NEW.** This iter could
  not attribute its whole-section delta because no same-tree pre-run exists; the iter loop has no cheap
  form for one, exactly as the ratchets had none before `--ceilings`.
- All routes from iters 207–216 unchanged, plus the standing queue.

**Lessons:**
- **A residual is only as wide as the matcher that produces it.** A vocabulary-reach metric measured
  through a narrow separator reports a *clean vocabulary* for as long as the hole is open — and it did,
  for four widenings.
- **A ratchet over a population that includes its own explanation has a fixpoint, not a value.**
  Iterate to it and say so; pinning once and moving on breaches by exactly the size of the reason.
- **Isolate a reach fix on a frozen tree.** `git archive` + the post-repair instrument separates *what
  became visible* from *what this commit wrote*, and without it the two arrive as one number nobody can
  attribute.
