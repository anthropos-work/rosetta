**Type:** tik — under `TOK-08`. Third consecutive member of iter-184's class, and the first reached by
**enumerating** the class rather than sweeping it. See
[`platform-alignment.md` §8](../../../../../corpus/ops/platform-alignment.md).

# iter-186 — "the whole-population baseline" was 5 of 11 sections

## Phase A — enumerate the class

iter-185 routed the honest gap: two members of *a fence's population is a registry too* had been found
by judgement and **nobody knew how many existed.** Enumerated: **70 module-level string collections** in
`stack-core`. The selector that worked is iter-177's shape — **two literals naming one population with
different cardinality**:

| literal | n | against the filesystem |
|---|---|---|
| `claim_census_guard.REXT_SECTION_NAMES` | 11 | exactly the 11 non-`knowledge` dirs ✔ |
| **`suite_census.SECTIONS`** | **5** | **6 omitted** |

## Phase B — what the six contain

| section | non-Python tests |
|---|---|
| `stack-seeding` | **119** `*_test.go` |
| `stack-snapshot` · `clerkenstein` | 45 · 37 |
| `playthroughs` | 22 `*_test.go` + **45** `*.spec.ts` |
| `alignment` · `stack-secrets` | 21 · 20 |

**264 Go test files and 45 TypeScript specs.** `suite_census.py` is the instrument behind this
milestone's **`3,369 passed · 9 failed · 4 skipped`** figure — quoted in run 17's brief and in iter
closes as *"the whole-population baseline"*. It is the whole **Python** population over **5 of 11**
sections, and the tuple said none of that (`D-M257x-186-2`).

The exclusion is **correct** — no Python runner collects a Go test — so the repair is a derivation plus
named reasons, never a widening.

## Phase C — derive, name, fence

`SECTIONS` is now derived from disk; the six are excluded **by name with a reason each** (`§5` rule 8);
and every run prints its scope before any total:

```
scope: 5 of 11 sections — Python only. 6 section(s) excluded BY LANGUAGE, not by absence:
  - alignment: Go module — 21 `*_test.go`; no Python runner collects it
  …
Any total below is a statement about those 5 sections and about no others (`§5` rule 60).
```

Four arms + a derivation control, **two mutants RED-proven with three failures each**: drop an exclusion
→ the section becomes invisible to both halves and the partition breaks; add an exclusion whose subject
has no tests → the stale-exclusion and partition arms fire (`D-M257x-186-4`).

## Phase D — measure

| gate | result |
|---|---|
| new population fence, both runners | **6 / 6** under `unittest` 3.9.6 **and** 3.14.6 |
| mutation controls | **2 mutants, 3 arms each, RED** |
| `suite_census` + `predicate_enumerator` + the iter-183/184 fence | **65 passed** (1.45 s) |
| existing `test_suite_census` + both registry guards | **84 passed · 0 failed** (122.23 s) |
| new `*_guard.py` | **0** — README fence-index triple does not move (Phase 0d, checked) |

Not covered: the rest of `stack-core` (1,594 P at iter-183), the 7 batteries, and — now stated rather
than implied — **the 264 Go tests and 45 TS specs in the six excluded sections, which no reading in this
milestone has ever included.**

## Close — 2026-08-09

**Outcome:** the instrument producing this milestone's headline suite number was reading **5 of the
repo's 11 sections** and calling it the whole population; the six it omits carry **264 Go test files and
45 TypeScript specs**. The omission is right and was unstated — `SECTIONS` is now derived from disk, the
six excluded by name with a reason each, the scope printed with every total, and the partition asserted
against the filesystem in both directions.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (eighteenth consecutive `closed-fixed`; **no
`P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n — (5)
cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN
ITERS, tree clean)
**Decisions:** `D-M257x-186-1` … `D-M257x-186-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter185-other-declared-populations-unaudited` — **narrowed, not closed.** The class is
  now **enumerated** (70 module-level string collections in `stack-core`) rather than swept, and one
  member repaired from that enumeration. What is still open is the **classification**: which of the 70
  are populations and which are predicates. There is no syntactic marker, so the split is judgement —
  and that is precisely why the class was invisible.
- `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` — **NEW, with its number.** No reading in
  this milestone has included the six non-Python sections: **264 `*_test.go` + 45 `*.spec.ts`**. They
  are not green and not red; they are **UNMEASURED**, and until iter-186 that was not visible from any
  published total. Bears directly on `D-M257x-145-3`'s second axis.
- `D-M257x-145-3` — **unchanged and explicitly NOT ruled on**, but it now has the denominator it
  lacked: *"all five sections"* is five of **eleven**. The ruling remains the user's.
- `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-only-ONE-registry-property-is-asserted` (half closed) ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` (owner: the next harden pass) ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  the observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **the third member of a class is the one you should stop hand-sweeping for** — iters 184
and 185 found theirs by judgement, and enumerating instead cost one command and delivered a sharper
target than either sweep. The selector that found it generalises: **two literals naming one population
with different cardinality**, which is iter-177's shape used as a *search key* rather than as a
post-mortem. And the corollary this iter is really about: **a correct exclusion is still a defect while
it is silent.** Nothing about omitting Go tests from a Python runner is wrong; calling the remainder
*"the whole population"* for sixteen iters was. Print the scope beside the number, every run. Written
into `platform-alignment.md` §8 in this iter's commit.
