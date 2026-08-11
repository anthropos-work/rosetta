---
iter: 173
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
---

# iter-173 — a DERIVED count can be re-derived; census the ones this milestone published

**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

`TOK-08`'s standing direction is to work mechanically-decidable classes exhaustively rather than sample
them. iter-172 closed by routing **`SURVEY-M257x-iter172-published-counts-predate-the-unit-fix`**:

> *"Every test-count this milestone published under the pytest runner before this iter is a **passed**
> count, therefore an undercount of the executed population. iter-170's `3,332` is corrected here; **any
> other quoted pytest count is not**, and nobody has enumerated where they are."*

Re-surveyed at open. The target is live and the route is untouched. **But its premise is a HYPOTHESIS, not
a plan** (the standing rule from iter-158), and the re-survey already qualifies it in two ways:

1. **"Therefore an undercount" does not follow universally.** A pytest `N passed` equals the executed count
   exactly when the run had zero failures and zero skips. The route asserts a defect for a population where
   part of the population is correct. The census has to *classify*, not assume.
2. **The observable half is not reachable by this instrument.** A directly-observed count (`1 failed · 1229
   passed`, read off a runner at a ref that no longer exists) cannot be re-checked without re-running, and
   `§5` rule 51's timing leg plus the 50-minute two-runner census at iter-172 make that the exact cost
   `build-iter` run 13 was told to stop paying.

So this iter scopes to the half that **carries its own evidence** (`§8`, iter-163): a **DERIVED** count — a
published number that is a *function of other published numbers* — can be re-derived from the page it sits
on, with no runner, no host and no clone. That is a census, not a sample, and it costs seconds.

## Cluster / target identified

**Population `P173` — every DERIVED test-count claim in the milestone record + `corpus/`.** Four shapes,
each mechanically recognisable:

| shape | example | how it is re-derived |
|---|---|---|
| `N of M tests` | *"1,280 of ~3,040 tests, 42 %"* | against the per-section table it summarises |
| a `\| **total** \|` row | *"2,978 passed · 22 failed · 11 skipped"* | against the rows above it |
| a delta `A → B (+K)` | *"105 → 121 passed (+16)"* | `B − A == K` |
| a percentage of tests | *"42 %"* | `N / M` |

**Out of scope, and stated up front per `§5` rule 60:** directly-observed counts. This iter's green is
evidence about derived counts alone and says nothing about whether any observed count was correctly read
off its runner. That residual stays routed.

## Hypothesis

The two aggregate denominators visible at open are already false, and one **inherits** the other:

- `hardening-ledger.md:2609` — *"1,230 of **2,989** tests"*. The table two lines above decomposes to
  `1229+1038+151+335+225 = 2978 passed`, `1+9+12 = 22 failed`, `11 skipped` → **executed = 3,011**.
  `2,989 = 2,978 + 11` — the **22 failures were dropped**. Exactly iter-172's defect, one level up: a
  total assembled out of `passed` while being labelled `tests`.
- `hardening-ledger.md:2827` — *"1,280 of **3,040** tests"*. `3,040 = 2,989 + 51`, so it carries the same
  22-test hole forward, and its numerator uses `1280 passed` where the executed figure is `1281`.
- **`corpus/ops/platform-alignment.md:2333`** — the same *"1,280 of ~3,040 tests, 42 %"* is published in the
  protocol doc **as the evidence for `§5` rule 68, the rule that "the whole suite" must name its
  denominator.** The rule about denominators ships a wrong denominator.

Expected: the census finds more of the same, and the fix is a fence that closes the arithmetic rather than
a hand-repair of three lines.

## Expected lift

No `P`/`N` reading is taken this iter — the metric is **UNMEASURED, not unmoved** (`§9` iter-type
refinement). The deliverable is the enumerated population with its denominator stated, every member
adjudicated by re-derivation, the false ones repaired **in place at every publishing site**, and a control
that fires.

## Phase plan

1. Enumerate `P173` mechanically; state the denominator.
2. Re-derive every member from its own page. Classify: `holds` / `false` / `underivable`.
3. Repair the false ones in place at every publishing site — corpus first, since a corpus site is the one a
   future reader quotes.
4. Fence: a check that re-derives the arithmetic, with a mutation control **and** an anti-vacuity control
   that can actually fire (eight vacuous fences have been caught in this milestone).
5. Run it; run the scoped suite for the touched module.

## Escalation conditions

- A member that cannot be re-derived because its decomposition was never published → **`underivable`**, and
  it is counted in the denominator and reported, never quietly dropped (`§9`: grade the cannot-tell).
- If the census returns **zero** findings, it must prove its instrument before that zero is reportable
  (`§9`, iter-149).

## Acceptable close-no-lift outcomes

The census enumerating `P173` and finding every member holds — **provided the instrument is proven to
fire** — is a complete iter. So is a demonstration that the routed premise is false at population scale.
