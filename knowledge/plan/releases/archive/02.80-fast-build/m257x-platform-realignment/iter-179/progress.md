**Type:** tik (standard shape; §9 iter-type refinements consulted, none selected).

# iter-179 — the routed repair was wrong, and the real defect was reach

**Controlling strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase A — re-derive the route before performing it

`FIX-M257x-iter174-accept-registers-one-registry-of-two` has been open since iter-174 and was narrowed at
iter-176 to one member: *"the battery seed list a `--accept` does not write."* Its embedded proposal is
*make `--accept` write it.*

**Refuted before a line was written** (`D-M257x-179-1`). The seed list's honesty is kept by
`test_000_the_copy_list_stages_every_fence_the_baseline_names`, which asserts *staged ⊇ the baseline's
fence names*. A `--accept` writing both registries from one source makes that assertion **compare a set
with itself** — iter-158's shape, already written in the battery's own seed-list comment and read past by
the route. *A routed item's proposed repair is a HYPOTHESIS, not a plan* — fourth application on this
milestone, first to a proposal that was **wrong** rather than over-broad.

## Phase B — what the contract actually is

The staged run's `discover_fences()` derives participation from the `FENCE_KIND` declaration **on the
disk it is pointed at**, and the staged baseline is a byte copy of the real one. So the equality the
staged suite (`test_iter45_mechanical_fences.py::test_21`) asserts is a **conclusion of two conjuncts**:

| conjunct | claim | asserted by | cost |
|---|---|---|---|
| **P1** | real tree: `baseline names == participating` | `test_iter45_mechanical_fences.py::test_21` | sub-second |
| **P2** | `_COPY_FILES ⊇ baseline names` | the battery's `test_000` | sub-second |
| **⟹** | staged tree: `staged-participating == staged baseline` | the staged run of `test_21` | minutes |

`participating ∩ staged = baseline ∩ staged = baseline` whenever `baseline ⊆ staged`. So the *"other
direction"* the route worried about is **entailed, not unfenced** — and **nothing recorded the
entailment.** P1 is not a safe premise to leave unstated: it was a **hard-coded set of four** until
iter-118, so the half P2's sufficiency rests on has already been the weaker one once (`D-M257x-179-2`).

## Phase C — measure the cost claim, do not inherit it

The route says the gap is *"reported only by a ~14-minute battery, one iter after the fact."* Measured at
HEAD `8422706`, units and runners named (`§9` rules 75/76):

| reading | value |
|---|---|
| the whole contract computed standalone | **0.10 s** — `/usr/bin/python3` 3.9.6 **and** `/opt/homebrew/bin/python3` 3.14.6, **agreeing** |
| `test_000` alone | **0.04 s** — pytest 8.4.2 on 3.9.6, `-k test_000` |
| the battery containing it | **718.97 s** measured this iter |

**There was no performance problem.** The check was only reachable by naming a file whose name says
`mutation_battery`, and the standing practice on this milestone — including in its own harden passes — is
scoped runs that exclude exactly those. iter-173's post-fix scoped re-run came back **167 passed, green**
and structurally could not have seen it (`D-M257x-179-3`).

## Phase D — the population, derived; the verdict, declared

Two modules stage the postcondition baseline. **Only one owes anything**, and the decline is *proved*:

| stager | staged participating fences | baseline names | unstaged | verdict |
|---|---|---|---|---|
| `test_m257x_mechanical_fences_mutation_battery.py` | **6** | 6 | 0 | **REQUIRED** — its staged suite carries `test_21` |
| `test_m257x_repair_postcondition_mutation_battery.py` | **1** (`claim_twin_guard`) | 6 | **5** | **DECLINE** — and it is green today with five unstaged |

Membership is derived on every run from each module's `_COPY_FILES` plus `repair_postcondition.
BASELINE_REL` — the baseline's name read **off its owner**, never spelled (§8 iters 70/71). The verdict is
declared per site, because inferring it means deciding *"does this staged suite assert the equality?"*
from source — the wrong-construct guess this milestone spends its iters finding (`D-M257x-179-4`).

## Phase E — what shipped

`stack-core/tests/test_battery_baseline_stage_m257x.py`, **9 tests**:

- **the population** — every baseline-stager classified, **both directions** (unclassified = RED, a
  classification for a module that no longer stages = RED);
- **anti-vacuity** — the derivation must find ≥ 2 stagers and the baseline must name ≥ 4 fences, each
  with the reading that makes a collapse legible as a broken instrument rather than a tidier tree;
- **P1** — asserted here as well as in `test_21`, with the other call site named: one derivation, two
  call sites, never two derivations (§8 iter-175);
- **P2** — for every REQUIRED member, computed by a pure `unstaged_fences()`;
- **four controls**, three synthetic and one negative-side: the finder fires and is silent correctly, the
  file-vs-bare-name comparison is the right way round, an unclassified intruder is named, and a battery
  that declares `_COPY_FILES` **without** the baseline does not enrol.

`test_000` **stays and delegates** to `unstaged_fences()` — its *ordering* is its deliverable (it must
fire before `test_00_`, inside the battery, so the failure reads *"you forgot a file"* and not *"the fence
is broken"*).

### The instrument was proven at the grain of its claim (`§9` iter-149/159)

Synthetic controls prove the arithmetic; they do not prove the wiring. Both arms were fired **live against
the real population**, out-of-band:

- drop `anchor_construct_guard.py` from the real `_COPY_FILES` → `test_04` **FAILS**, naming the file, the
  fence, and the symptom it prevents;
- inject one unclassified stager → `test_01` **FAILS**, naming the intruder.

### Registry tax — measured, not assumed (`D-M257x-179-5`)

iter-178 put its new arm inside an existing guard to avoid the four-registry tax. This iter adds a module
one iter later, so the exemption is evidenced: those four registries are keyed on **fences** (`*_guard.py`
/ `FENCE_KIND`), and a `tests/` module is in none of their populations — `23 passed` on the two registry
fences and `77 passed` on the frozen-expectation census + guard-family suite with the new file present.
**The one registry it does owe is the prose index**, and `stack-core/README.md` had **no row** for it —
nor for **iter-176's own fence**, three iters old. Both rows added.

## Runs — scope stated, and what it did NOT cover

| run | scope | result | wall |
|---|---|---|---|
| batteries + suites touching this change | mechanical-fences battery, `test_iter45_mechanical_fences.py`, registry-population fence, `test_repair_postcondition.py`, repair-postcondition battery | **132 passed · 0 failed** | 718.97 s (11:58) |
| `stack-core` minus all 7 mutation batteries | `tests/` | **1,521 passed · 2 skipped · 0 failed** | 481.28 s (8:01) |

Runner **pytest 8.4.2 on `/usr/bin/python3` 3.9.6** for both; the new module additionally re-run green on
`/opt/homebrew/bin/python3` 3.14.6 under `unittest` (`Ran 9 tests … OK`), so the two-interpreter
disagreement of iter-170 does not touch it.

**Whole-section arithmetic, stated so it can be checked:** the 7 mutation batteries collect **41** tests,
so `stack-core` is **1,562** collected — exactly iter-178's **1,553** plus this iter's **9**. Of those 41,
**11** ran green here (the two batteries this change touches); **30 did not run** (`m220`, `m255`,
`claim_twin`, `repair_reach`, `repair_leak` — none of which stages the postcondition baseline or imports
anything this iter touched). **The four other rext sections were not run.** A scoped green is evidence
about its scope alone (rule 60).

## Close — 2026-08-09

**Outcome:** the open member of `FIX-M257x-iter174-accept-registers-one-registry-of-two` is **closed by
refuting its own proposed repair**, and the defect it actually named turns out to be **reach, not cost** —
the check is 0.04 s and was unreachable from every scoped run this milestone habitually does. Shipped: a
1.4 s fence over a **derived** population of baseline-stagers (2 members, one REQUIRED and one DECLINE
proved by a green battery running with five of six baseline fences unstaged), asserting **both conjuncts**
of the sufficiency pair rather than one direction of a supposed equality — because the premise, not the
bridge, has already been the weaker half once.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (eleventh consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-179-1` … `D-M257x-179-5` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none. (The two README index rows are part of the planned registry-tax step, not a
side discovery.)

**Routes carried forward:**
- `FIX-M257x-iter174-accept-registers-one-registry-of-two` — **CLOSED**, by refutation plus a fence over
  the population its member belonged to. The member's obligation is now *asserted*, not *recorded as
  open*, and the population fence's verdict text for that site is updated to say so.
- `SURVEY-M257x-iter179-readme-indexes-test-modules-unmeasured` — **NEW.** The disclosed
  `16 of 27 / 16 of 26 / 15 of 26` triple measures the README's coverage of **fences**. Its coverage of
  **test modules** has never been measured, and it was missing a three-iter-old fence. A fourth
  derivation would need a row in the checked triple, so this is a design decision, not a corollary.
- `SURVEY-M257x-iter179-thirty-battery-tests-unrun` — **NEW.** 30 of the 41 mutation-battery tests were
  not run this iter. None stages the postcondition baseline; the claim that they are unaffected is an
  argument, not a measurement. Owner = the next harden pass.
- `FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` · `FIX-M257x-iter173-ledger-
  denominator` · `SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` · the observed
  half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a check is only as strong as the runs that reach it — and "too slow" is a measurement.**
Four iters of routing described a 0.04 s set-comparison as a fourteen-minute cost, and the repair that
description implied would have broken the assertion it was meant to strengthen. Two corollaries paid for
directly: when you relocate a check for reach, **keep whatever was load-bearing about its old position**
(here, ordering — so delegate rather than move); and **a sufficient PAIR must be asserted as a pair**,
because an entailment nobody writes down can be broken from the end nobody is watching. Written into
`platform-alignment.md` §8 in this iter's commit.
