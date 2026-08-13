**Type:** tik · `iter_shape: tooling`
**Active strategy:** `TOK-07` **step 2, first half — rule 4**.

---

# iter-114 — the reach metric names its denominator, or prints no percentage

## What was wrong

`repair_reach_guard` took exactly one input: a `raw/` directory of seat reports — **what one reading
happened to detect**. Per-pass detection recall on that instrument has run **33–83 %**, so that set is a
**sample** of where a predicate is published, never the population. And the guard printed
`reach t/N = P%` over it unconditionally.

That is the instrument that graded **iter-108 at 46/46 = 100 %**. The arithmetic was right. The repair was
right. The next reading found the same propositions still false one file away — because **a predicate's
site list and a reading's detection list are not the same set.** A check reporting a state it did not
measure: the milestone's signature defect, reached into the metric that was supposed to catch it.

## What landed

**One declared denominator, named in the report, and only one of the two may carry a ratio:**

| input | provenance printed | percentage |
|---|---|---|
| `--enumeration <json>` | `corpus-derived-per-predicate` | **yes** |
| `--ledger <dir>` | `prior-reading-detections` | **no** |

**The refusal is the absence of the number, not a caveat beside it.** `TOK-07` rule 4 says such a run must
be *"unable to print a reach percentage at all"*, and a percentage is precisely the thing that gets quoted
in a close — a warning next to it is not. In `--json` the **`reach_pct` key is omitted**, so a machine
consumer raises rather than formatting a figure the run was never entitled to. The counts stay (`109/147`
is a fact about reach); only the ratio goes.

**An enumeration is refused outright (exit 2) when it is not a settled denominator** — seed recall failed,
or a predicate still carries unadjudicated headroom. iter-113's whole point is that a ceiling nobody
settled leaves a **candidate list**, and a candidate list is not a population either. Also refused:
both inputs at once (not a state the report can name), neither (not a measurement), an empty
enumeration, and a malformed site — **refused rather than dropped**, because a silently shrinking
denominator is the one direction a reach number flatters itself in.

## Measured, both paths, live

```
--ledger iter-76/raw      --range 328ece5   152 booked from 14 report(s)
                                            denominator: prior-reading-detections (raw)
                                            reach 109/147 reached — NO PERCENTAGE IS AVAILABLE
--enumeration iter-113    --range 461b547    71 enumerated site(s) over 24 predicate(s)
                                            denominator: corpus-derived-per-predicate
                                            reach 0/71 = 0.0%
```

**The second line is step 2's pre-repair baseline, and it is worth having measured before the repair
starts: `0/71`.** iter-113 touched `knowledge/` and the protocol doc, not the corpus sites, so 0 is
correct — and it means the repair's own grade cannot start from an accidental head start.

**The known-answer fixture is unchanged**, which is the control that proves the extension did not soften
the fence: iter-81's repair against iter-76's ledger still classifies **109 touched · 35 line-unreached ·
3 file-unreached · 4 no-anchor · 1 out-of-tree = 152**, exit 1, `graphql-wundergraph.md:13` still named
from both readings. Only the percentage is gone.

**Tests: 21 → 30 in this file. 60 passed with the enumerator's suite alongside; 46 passed for the reach
mutation battery + `guard_family`** (`/usr/bin/python3 -m pytest ... -q`; 0.78 s and 12.09 s — invocations
stated).

The positive control matters as much as the refusal: `test_02_an_enumeration_denominator_DOES_print_a_percentage`
exists so the refusal cannot be satisfied by a fence that never prints a percentage for anything, which
would measure nothing. And `test_02_the_shipped_iter113_enumeration_is_an_acceptable_denominator` is the
anti-vacuity control — it loads **the artifact iter-113 actually checked in** and asserts 24 predicates /
71 sites, so the whole class cannot be exercising a shape nothing produces.

## What this iter did NOT do

**The repair itself.** `TOK-07` step 2's other half — repair the 24 predicates across the 71 enumerated
sites — is iter-115. **No reading was taken**; `P` is **UNMEASURED, not unmoved** (§9). Gate unchanged at
**4 of 5**.

Zero platform-repo edits · `stack-demo/**` untouched · no clone fetched (§5 rule 41a) ·
`rosetta-extensions` on `main`, no tag cut · clause 5 not re-cut, narrowed or argued.

---

## Close — 2026-08-07

**Outcome:** `repair_reach_guard` can no longer print a completeness figure over a sample. The denominator
is declared and named in every report; `corpus-derived-per-predicate` may carry a ratio and
`prior-reading-detections` may not — the number is **absent**, and the `--json` key **omitted**, rather
than caveated. An unsettled enumeration is refused outright. The iter-81/76 known-answer fixture is
unchanged (109/147, 38 unreached, exit 1), which is what proves the extension did not soften the fence.
Step 2's pre-repair baseline is measured: **0/71**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — (7) budget-exhausted: **y — between iters, both trees committed and pushed** — Outcome: exit-7
**Decisions:** D-M257x-114-1, D-M257x-114-2
**Side-deliverables:** none
**Routes carried forward:**
- `TOK-07` step 2, **second half — the repair itself**: 24 predicates over the 71 enumerated sites,
  graded `--enumeration iter-113/enumeration.json` against the repair commit, from the measured
  baseline **0/71** → next iter
- `TOK-07` step 3 (the read) → last, unchanged
- `FIX-M257x-iter113-adjudication-is-judgement` → open
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → open
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → open
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → open, de-ranked
- `DEF-M257x-iter101-briefing-rext-tree` → open
**Lessons:**
- **Withhold the number, do not annotate it.** A caveat beside a percentage loses to the percentage every
  time, because the percentage is what gets copied into a close. The same asymmetry applies to `--json`:
  omit the key, never null it, or a consumer formats a figure the run was not entitled to.
- **Ship the positive control with the refusal.** "It refuses to print a percentage" is satisfiable by a
  tool that never prints one. Only the pair measures anything.
- **Measure the pre-repair baseline before the repair.** `0/71` costs one command now and makes the
  repair's own grade unarguable later.
