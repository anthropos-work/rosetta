# iter-119 decisions

## `D-M257x-119-1` — the fall in `P` is SAMPLING, not drainage; the floor rose while `P` fell

`P` went **37 → 22 (−40.5 %)**. The tempting reading is that the sweep nearly worked. **It is false, and
this reading was built to be able to say so.**

The corpus changed by **5 in-place lines, zero net**, and **none of iter-116's 37 predicates was
repaired** (verified against the 8 census repairs: 4 in `platform-migration-status.md`, 1 in
`architecture_overview.md:48`, 3 out of scope — no iter-116 predicate anchors at any). A pool cannot drain
40 % when nothing was removed from it.

Band #3, pre-registered before the number: **13 of iter-116's 37 re-found (35.1 %)**, plus **9 net-new**.
The 24 unseen were adjudicated UPHELD against the same clones; they did not stop being false because a
second panel missed them.

> **Union floor: ≥ 46 at `194361e4`** — higher than any single reading has returned, and the first floor
> on this milestone derived from two readings of one corpus rather than one.

**Consequence:** `P` movements across this milestone's five readings (**22, 22, 24, 37, 22**) are
substantially instrument, not corpus. No branch of any past reading that rested on a `P` *delta* should be
re-used without this correction.

## `D-M257x-119-2` — `P` is a WEAK floor, and now quantifiably so

Three independent measurements of the instrument's precision, all from this reading:

| measurement | value |
|---|---|
| test-retest recall against a prior panel's adjudicated set | **35.1 %** |
| per-pass recall against the two passes' OWN union | **63.6 % · 77.3 %** |
| adjudicator-granularity sensitivity | **±2** (iter-116 split `ai_architecture.md`'s three stale self-citations into 3 predicates; iter-119's panel collapsed them into 1 — had it split, `P = 24`) |

**`P` is comparable across readings to roughly ±2 on granularity alone, before sampling variance.** This
is disclosed in `adjudication.md` § Provenance rather than buried, because it is the second independent
reason the series' movements are smaller than they look.

**This does not soften the refutation** — `P = 22` clears `P ≥ 19` by 3 and `P = 24` clears it by 5; every
granularity reading of this sheet fires the same branch.

## `D-M257x-119-3` — the census and the reader work OPPOSITE HALVES of the citation class

`D-M257x-117-2` predicted this at iter-117's close, in writing, before either census finished. It is now
measured:

| | iter-116 | iter-119 |
|---|---|---|
| intra-corpus citation defects, share of `P` | 10 of 37 = **27 %** | 8 of 22 = **36 %** |

**The census closed the RESOLUTION half at 100 % reach and 0 findings. The CONSTRUCT half grew as a share
of what remains.** A machine can prove a link resolves to a file that exists; it cannot prove that the
lines a pin names hold the **Data** bullet rather than the **Domain** bullet. **Five of the eight were
booked at iter-116 and are still false.**

**Rule this establishes:** *say which half of a class your fence reaches before you claim the class.* A
fence at 100 % reach over its enumerated population can leave the class's dominant half untouched, and the
reach metric will not tell you — it names its denominator, and its denominator is the half it can see.

## `D-M257x-119-4` — a census cannot promote a small-class verdict

The two censuses enumerate **citations**. A `SMALL-CLASS-*` verdict is a claim about a **ceiling** — *no
other site in the corpus publishes this predicate*. **Neither census measures a ceiling**, so neither can
convert a judged verdict into a proven one.

**Measured, in answer to the standing question on `FIX-M257x-iter113-adjudication-is-judgement`:**

- **0 of the 16** small-class verdicts were promoted from judged to proven by `TOK-08`'s sweep.
- **P20 remains the only `SMALL-CLASS-PROVEN`**, and it was proven at iter-113 by its own zero-headroom
  ceiling.
- This reading **re-books 2 of the 15 judged** (P23 *Ant Academy internal-only*, P08 *the `⚠⚠ M51`
  anchor* — fifth generation) and **0 of the 1 proven** → measured error rate of the judged verdicts
  **≥ 13.3 %**. P23's exclusion list dismissed 12 candidates as publishing a different proposition; at
  least four publish exactly this one.

The asymmetry — the judged fail, the proven holds — is exactly what the open FIX predicted, and it is now
measurement rather than suspicion.

## `D-M257x-119-5` — no successor strategy is authored, and that is the instruction, not a judgement

`TOK-08`'s sealed rule (milestone-root `decisions.md`, `577446b`) reads: *"`P ≥ 19` → … **STOP. Do NOT
author a successor strategy.** Report the refutation and hand the milestone back for a scope decision from
the user. The user has stated in advance that this is a legitimate outcome and that they will carry it."*

`P = 22`. **The branch fires. No `TOK-09` is written.** Phase 5 grades this `exit-3`
(`re-scope-trigger`), not `tok-fired` — a triggered tok would author the very thing the rule forbids.

**Both agent-authored and user-authored strategies have now been refuted by their own pre-registered
arithmetic** — `TOK-07` at iter-116 (`P = 37` vs `P ≥ 15`), `TOK-08` here (`P = 22` vs `P ≥ 19`). The
milestone's next move is a **user scope decision**, and this iter deliberately does not pre-empt it.
