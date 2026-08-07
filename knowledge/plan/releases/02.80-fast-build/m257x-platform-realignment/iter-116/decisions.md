# iter-116 decisions

## `D-M257x-116-1` — the guard family is RED at this reading's open, and the RED is FALSE

`platform_predicate_guard` check G10 reports `corpus/services/sentinel.md:5` as declaring the wrong
compose-service count, grading it at `d11a403` (8 file-local / 10 effective). **The sentence names
`0c91421d`, where the pair is `(5, 7)` — measured with the guard's own `compose_counts_at` helper — and
all five of its line anchors resolve. The corpus claim is TRUE at the ref it names.**

The defect is in the guard: `sentinel.md:5` is one long wrapped paragraph carrying **two** platform refs,
and G10 takes `_REF_PINNED.search(cell)` — the **first** match in the window — so it dates the claim by
whichever ref appears earliest rather than by the ref the claim names. That is §5 rule 33 violated by the
guard that enforces it.

**Routed as `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`. NOT repaired here** — no repair
is taken inside a measuring pass (`pre-registration.md` binding condition 3). Disclosed to the seats in
the addendum so a seat that notices the same site books what it measured rather than what a guard said.

**It does not block the reading, and the reason is stated rather than assumed:** this iter lands zero code
and zero corpus edits. A RED gate blocks an iter because code must not land on top of one; there is
nothing to land, and the RED's subject was opened at source and holds.

**Second guard in two iters caught resolving a claim against the wrong thing** — the first being
`FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live`. That is a pattern worth naming even though
neither is a corpus defect.

## `D-M257x-116-2` — DENOMINATOR is falsified, VOLUME is the answer; and the same reading EXONERATES the repair

The pre-registration named two stories for why iter-109's pool did not drain, and sealed the cuts that
separate them. They separate cleanly:

- **DENOMINATOR** (the repair only visited detected sites) — **FALSIFIED.** Band #3, blind and on
  predicate, measured **3 re-found of iter-109's 24**. Twenty-one stayed closed. Band #11 measured
  `N`/`P` at **1.108**, down from 1.50 twice — the twins really were closed corpus-wide.
- **VOLUME** (the pool is far larger than a reading samples) — **the answer.** Of the 37, **25 are
  standing pool that no prior reading ever detected**, after four readings and two full repair cycles.

**This is why the refutation is useful.** `TOK-07`'s *mechanism* is vindicated by the same numbers that
refute its *premise*. Grading the strategy as a single thing would have discarded a working enumerator —
which is §9's *grade a refuted strategy LEG BY LEG* (written at iter-110, for `TOK-06`) arriving one
revision later on the strategy that rule was authored for.

**Consequence, and it is not this iter's to take:** `TOK-07` pre-registered `P ≥ 15` as refuting
repair-and-read as a path to clause 5 **under this instrument**, and named the next move as a re-scope
conversation with the user. `P = 37`. **No `TOK-08` is authored.** Clause 5 is untouched — it is met only
by a reading that returns zero, and a tok never revises a gate.

## `D-M257x-116-3` — the induction fences do not SCALE with repair volume

Band #10, measured mechanically by `git blame` of each upheld anchor against the iters-110–115 commit
range: **9 of 41 anchors (22.0 %)** were authored by iter-115's own repair.

| cycle | repair volume | induced share |
|---|---|---|
| ≤ iter-102 | — | ~21 % |
| iter-108 | **+48** lines | **5.6 %** — the lowest in the series, and the reason step 2 was called vindicated |
| iter-115 | **+177** lines | **22.0 %** |

`anchor_offset_guard` and `repair_postcondition` were in production for both, and **fired four times on
iter-115 and were repaired each time**. Nine still landed. The fences are not broken; they do not scale.

**The rule this establishes:** *"the fences are in production"* is a statement about the previous
volume. **A repair must state its expected induction cost as a function of its size**, and a repair large
enough to change the induction rate is a different intervention from a small one, not more of the same.
Routed as `FIX-M257x-iter116-induction-fences-do-not-scale`.

## `D-M257x-116-4` — intra-corpus mis-citation has overtaken platform drift as the largest class

Band #7 predicted ≤ 4 wrong-construct intra-corpus citations among the upheld set and measured **10 of 37
predicates (27 %)** — the single largest class in the reading. Platform/tooling-source drift is **14 of
37 (37.8 %)** but is spread over several mechanisms; the mis-citation class is one mechanism.

**The corpus is now large enough, and repaired often enough, that it mis-cites ITSELF more often than it
mis-describes the platform.** That is a different problem from the one this milestone was opened to
solve, it is mechanical, and it is the one class a machine could close outright. Recorded as evidence for
the re-scope conversation, and routed as
`FIX-M257x-iter116-intra-corpus-miscitation-is-the-largest-class`. **Not repaired here** — no repair is
taken inside a measuring pass.

## `D-M257x-116-5` — the enumeration's completeness warrant is weaker again: a FOURTH measured miss

`FIX-M257x-iter113-adjudication-is-judgement` already carried three measured misses from iter-115. This
reading adds a fourth and it is the largest: adj-3's P6 records the *Ant Academy is internal-only*
predicate as live at **five** sites — `ant-academy.md:31` (+`:5`, +`:298`), `architecture_overview.md:40`
(+`:260`), `service_taxonomy.md:290`, `frontend_architecture.md:9`, `services/README.md:58` — and the
enumeration reached **none** of them, while the same predicate had been upheld at iter-109.

**Only the two anchors a seat actually booked enter `N`.** The other sites are recorded as a reach
finding, not folded into the number — renegotiating the metric to include an adjudicator's incidental
observation would destroy its comparability with three prior readings. The honest statement is that
**`N = 41` is a floor and is known to be one at a named site.**
