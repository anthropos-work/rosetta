**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*
**Shape:** standard tik (no protocol-codified shape). Opened 2026-08-08 23:30, `rosetta` `9c118ae`.

# iter-173 — the rule about denominators was published with a wrong denominator

## What this iter was handed

iter-172 closed by routing `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix`: every test-count
this milestone published under the pytest runner before that iter is a **`passed`** count, *"and nobody
has enumerated where they are."*

## Step 0 — the re-survey qualified the route before accepting it

Two qualifications, both recorded as `D-M257x-173-1`:

1. **The premise is not universal.** A pytest `N passed` equals the executed count exactly when the run
   had zero failures and zero skips. This iter's own controls are the demonstration: **12 tests**, printed
   as `Ran 12 tests` by unittest/3.14.6 and `12 passed` by pytest/3.9.6 — **the two units coincide because
   there is nothing to drop.** The defect is not that `passed` was published. It is that `passed` was
   published **under the word `tests`**, and then *summed*.
2. **Half the class is out of reach at any price this iter can pay.** So the class was split by
   **derivability**:

| | re-checkable here? | how |
|---|---|---|
| **observed** — read off a runner (`1 failed · 1229 passed`) | **no** — needs a re-run at a ref that may not exist; `§5` rule 51's timing leg is unusable on this host, and iter-172's two-runner census cost 50 min | — |
| **derived** — a function of other published numbers on the same page | **yes** | arithmetic, on the page itself |

The derived half is a census that runs in **under a second**. The observed half stays routed, and this
iter's green is evidence about the derived half **alone** (`§5` rule 60).

## THE FINDING

The harden ledger summarises its own five-section whole-suite table as:

> *"one section of five, **1,230 of 2,989** tests"*

Six lines above it, that table says `2,978 passed · 22 failed · 11 skipped`.

```
2,989 = 2,978 passed + 11 skipped          ← the 22 FAILURES were dropped
3,011 = 2,978 passed + 22 failed + 11 skipped   ← executed
```

**The denominator had silently changed unit** — from *executed* to *passed-and-skipped*. It is iter-172's
defect one level up, and where it ended up is the point: the next ledger entry carried the hole forward as
`2,989 + 51 = 3,040`, and from there the figure reached **`corpus/ops/platform-alignment.md`, where it is
published as the EVIDENCE FOR `§5` rule 68 — the rule that "the whole suite" must name its denominator.**

### Re-derived, and cross-checked against a second run

| | passed | failed | skipped | **executed** |
|---|---|---|---|---|
| `stack-core` (pass 32) | 1,229 | 1 | 0 | **1,230** |
| the four other sections (pass 32 table) | 1,749 | 21 | 11 | **1,781** |
| the four other sections — **iter-145's independent re-run** | 1,758 | 21 | 2 | **1,781** |
| whole suite @ pass 32 | 2,978 | 22 | 11 | **3,011** |
| `stack-core` after the `+51` | 1,280 | 1 | 0 | **1,281** |
| **whole suite, the corrected figure** | | | | **3,062** |

**Two independent runs agree on 1,781 executed** with different pass/skip splits (9 tests moved from
skipped to passed between them). That agreement is what lets **1,281 of 3,062** be published as a number
rather than an estimate.

**And why it survived 28 iters unnoticed: `1,280/3,040` and `1,281/3,062` are both 42 %.** The headline
conclusion — *"'the whole suite' meant 42 % of it"* — was true throughout. **A percentage can survive an
error its operands do not**, which makes a ratio the most durable place for a wrong count to live.

## The census — population, with its denominator

`stack-core/derived_count_guard.py` (net-new), three zero-inference arms over **698 markdown files** of
the milestone record + `corpus/` (**240 excluded**: `raw/` and `evidence/` hold pre-adjudication seat
artifacts and captured run logs — inputs committed verbatim by protocol, not claims; editing one to
satisfy a fence would corrupt an evidence artifact, so the exclusion is printed rather than left silent).

| arm | what it re-derives | sites evaluated | findings |
|---|---|---|---|
| A — table total | a `total` row vs the component-wise sum of its own table's rows, matched as a *bag* (`passed`/`failed`/`skipped`/`error`) so a total in one vocabulary is never compared against another | 16 | 0 |
| B — explicit delta | `A → B … (+K)` ⟹ `B − A == ±K` | 5 | 0 |
| C — percent triple | `X of ~Y …, Z %` ⟹ `round(100·X/Y) == Z` | 4 | 0 |
| **total** | | **25** | **0** |

⚠️ **That 25 is pinned to the tree the census RAN on — and the number is already 27.** The guard's subject
is the published record, this report is part of the published record, and the census table above is itself
two arm-A sites (one per numeric column). **A census whose subject includes its own report has a
population that moves as it publishes** (`D-M257x-173-6`), which is why the figure is stated with the tree
it was taken at rather than as a standing fact — `§5` rule 51(b)'s discipline, arrived at from the other
direction. The 27-site run is green too; the two added sites are the row you are reading.

**NOT REACHED, and the guard prints this on every run including green:** the `N of M` prose shape itself.
`M` names no source, and attributing it to a nearby table is an *inference* — precisely what
`D-M257x-117-2` and iter-119's refutation established this milestone cannot afford. So the guard fences
the **operands** instead (`D-M257x-173-4`): the prose repair stays a hand repair, but it now rests on a
fenced ground truth rather than on a second reading.

### Arm B's first draft fired on six real lines, and all six meant something else

The first version paired any `A → B` with any following `(+K)`. Against the real record that produced
**six findings, zero of them true**, in four distinct shapes — a percentage delta (`37 → 22 (−40.5 %)`), a
decimal percentage (`10,278 → 10,646 lines (+3.6 %)`), a diffstat pair (`(+167 −80)`), and a prose
parenthetical (`16 → 21 (+1 existing test updated …)`). Arm B now requires a **parsed construct**: a
digit-free gap between the arrow and the parenthesis, and a parenthetical that is *exactly* a signed count
(§8 — assert against a parsed construct, never a substring that merely contains one). **All six are pinned
as literal text in the controls**, so a later loosening fails in the tests rather than in the report.

## The repair — three sites in place, two routed

| site | disposition |
|---|---|
| `corpus/ops/platform-alignment.md` (rule 68's evidence) | **repaired** → `1,281 of 3,062 tests`, *tests = executed = passed + failed + skipped*, with the retraction stating how the old figure was assembled |
| milestone `progress.md` (iter-145's ledger entry) | **repaired** — and it is self-refuting: the per-section figures that disprove the total are **in the same sentence** |
| `iter-145/decisions.md` | **repaired** (operands), the ~42 % explicitly noted as unchanged |
| `iter-145/overview.md` — a **block quote** of the ledger | **annotated beneath, quote left verbatim** (`D-M257x-173-3`): editing a quote to fix its source makes the record cite a sentence nobody wrote |
| `hardening-ledger.md` ×2 | **ROUTED** — that file is owned exclusively by `/developer-kit:harden-mstone-iters`; this skill reads it and must not write it. Derivation pre-computed so the next pass applies an answer (`D-M257x-173-2`) |

## The fence was graded by the family before it was trusted — and failed three times

The guard's own 12 controls were green while the guard was still non-conformant with the family it was
joining. `test_fence_provenance` caught two (**no tree stamp on direct execution**; `--json` not
self-describing), `repair_postcondition` caught a third (**`FENCE_KIND = "postcondition"` declared with no
`postcondition_sites()`** — a *copied* declaration, the registry-rot shape this milestone keeps finding),
and the family view surfaced a fourth: the guard **printed nothing when green**, so `guard_family`'s
`lines[-1]` rendered it as a blank summary. All four fixed; the one-line verdict now prints **last** and
carries its reach clause with it. Full table in `D-M257x-173-5`.

## Controls — the green is not the deliverable, the firing is

Eight vacuous fences have been caught in this milestone, so:

- **mutation control per arm** — a correct fixture broken in exactly one place; each arm fires. Arm A's
  mutant is *the iter-173 defect in miniature*: a total assembled out of `passed` alone.
- **a NO-OP control REQUIRED TO SURVIVE** — a battery in which every mutant dies is measuring the harness
  (`§5` rule 7).
- **anti-vacuity written against the guard's SUBJECT, not its inputs** (iter-94): the guard must be shown
  to have **evaluated a site in each of the three files the census adjudicated by hand** — a count alone
  cannot tell *checked and passed* from *never looked at*. The guard therefore records every evaluated
  `path:line`, and the control asserts membership.
- **a refusal control** — an empty subject exits **2**, never 0 (`§9` iter-149: a census returning ZERO
  must prove its instrument).
- **six false-positive pins** taken from real corpus lines.

## Protocol evolution

`corpus/ops/platform-alignment.md` `§8` gains **"A DERIVED number is censusable; an OBSERVED one is not —
split the class before you scope the fence"**, with three sub-rules: *a percentage can survive an error
its operands do not, so never audit the ratio in place of the operands*; *a fence that cannot reach the
claim is still worth building if it proves the claim's operands, provided it prints what it does not
reach*; and *two independent runs agreeing on a total is stronger evidence than either*. Plus the
file-ownership note that shaped the repair.

## THE MEASUREMENTS — counts, not wall-time (`§5` rule 51's timing leg is unusable on this host)

**The new fence, both runners** (`§5`, iter-170 — *name the runner*):

| runner | invocation | result |
|---|---|---|
| pytest / `/usr/bin/python3` **3.9.6** | `pytest tests/test_derived_count_guard.py -q` | **12 passed** |
| unittest / `python3` **3.14.6** | `python3 -m unittest tests.test_derived_count_guard` | **Ran 12 tests · OK** |

The two units coincide here **because there is nothing to drop** — 0 failures, 0 skips. That is the
demonstration behind `D-M257x-173-1`, not a coincidence worth passing over.

**Fence-conformance set** (`test_fence_provenance` + `test_guard_family` + both `repair_postcondition`
suites + the new controls): **143 passed** in 170.64 s, after four conformance fixes.

**The whole `stack-core` suite, on a stable tree** — no edit was made between launch and result
(eight runs have been discarded as confounded for exactly that):

```
3 failed · 1525 passed  in  1368.14 s (0:22:48)     /usr/bin/python3 3.9.6, rext 99e8aec + this iter
```

**And it caught a regression this iter had just introduced** — `test_21_the_shipped_baseline_records_
EVERY_participating_fence`, the new fence missing from the repair-postcondition **ratchet baseline**. The
scoped runs were all green. Full grading of all three failures, and the reason the scope enumeration
missed the fourth registry, in `D-M257x-173-7`.

**After the repairs:** ratchet baseline registered · all three unclassified derivations classified ·
`test_frozen_expectation_census` + `test_iter45_mechanical_fences` + both `repair_postcondition` suites
re-run together → **167 passed**. One failure remains and it is **not this iter's**:
`test_battery_stage::test_a_stdlib_shadow_is_refused_not_staged` (`RuntimeError not raised`, no reference
to this iter's module), pre-existing and routed.

**The four registries a new fence had to join** — the enumeration is the deliverable, because getting it
by grep got it wrong:

| registry | how it is kept | found by |
|---|---|---|
| `guard_family.py` GUARDS | hand-maintained dict | grep for a sibling's name |
| `derivation_registry.py` DECISIONS | hand-maintained dict | grep |
| `stack-core/README.md` fence table | hand-maintained markdown | grep |
| **`repair_postcondition_baseline.json`** | `--accept`-written JSON ratchet | **only the whole-suite run** — it names no module in any `.py`, so a source grep is structurally blind to it |

## Close — 2026-08-09

**Outcome:** The milestone's own rule that *"the whole suite" must name its denominator* was **published
with a wrong denominator**. `1,230 of 2,989` drops 22 failures from the very table it summarises
(`2,989 = 2,978 passed + 11 skipped`; executed is **3,011**); the next entry carried the hole forward as
`2,989 + 51 = 3,040`, and it reached `corpus/ops/platform-alignment.md` as the **evidence for `§5` rule
68**. Corrected to **1,281 of 3,062**, cross-checked by two independent runs that agree on the
four-section subtotal of **1,781 executed**. iter-172's routed class was **split by derivability** —
derived counts are censusable on the page, observed ones are not — and the derived half is now fenced by
net-new `stack-core/derived_count_guard.py` (**25 sites, 0 findings, 3 arms**, with a NOT-REACHED clause
printed on every run). **And the 42 % held under both readings**, which is why the error survived 28
iters.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (fifth consecutive `closed-fixed`; no
no-prog streak, and **no `P`/`N` reading was taken, so the metric is UNMEASURED, not unmoved** — `§9`) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
n — Outcome: **continue**
**Decisions:** `D-M257x-173-1` … `D-M257x-173-7` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none. The two pre-existing `derivation_registry` classifications were not a side
fix but a forced consequence of this iter's own entry (`D-M257x-173-7`).

**Routes carried forward:**
- `FIX-M257x-iter173-ledger-denominator` — **NEW.** `hardening-ledger.md` lines carrying *"1,230 of
  2,989"* and *"1,280 of 3,040"*, to be corrected by the next harden pass (which owns that file) to
  **1,230 of 3,011** and **1,281 of 3,062**. The `59 %` on the following line holds under both readings
  (58.85 % → 59.15 %) and needs no change. Derivation pre-computed; nothing to re-derive.
- `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — **half CLOSED.** The **derived** half is
  censused and fenced. The **observed** half (a `passed` count read off a runner at a past ref) stays
  open and is out of reach of any instrument that does not re-run a suite.
- `SURVEY-M257x-iter172-two-preexisting-actionable-reds` — **half CLOSED.** The
  `derivation_registry` member is closed (forced by entanglement); `test_battery_stage::
  test_a_stdlib_shadow_is_refused_not_staged` remains open, pytest/3.9.6-only, rule-76 shaped.
- `SURVEY-M257x-iter173-derived-count-guard-reach` — **NEW.** The guard reaches table totals, explicit
  deltas and percent-triples. It does **not** reach the `N of M` prose shape, `A → B` with no stated
  delta, or any prose total. Population of the unreached shapes is **not measured**.
- The standing queue, unchanged.

**Lessons:** **a DERIVED number is censusable; an OBSERVED one is not — split the class before you scope
the fence.** iter-170 earned *name the runner*, iter-171 *an unexplained disagreement is a defect*,
iter-172 *name the unit*; this iter is the fourth turn and it moves up a level: the offender was not a
column but a **conclusion**, and the arithmetic under it had been quoted four times without once being
re-added.

Three that generalise:

1. **A percentage can survive an error its operands do not.** `1,280/3,040` and `1,281/3,062` are both
   42 %. Auditing the ratio would have passed forever. **Audit operands, never conclusions.**
2. **A fence that cannot reach the claim is still worth building if it proves the claim's operands** —
   provided it prints what it does not reach, on every run, green included.
3. **The scope of a change is not derivable by grepping for a sibling's name.** Three registries were
   found that way; the fourth is a JSON ratchet that names no module in any source file, and only the
   whole-suite run saw it. `§5` rule 73 — *a glob is not a derivation* — extends to greps.
