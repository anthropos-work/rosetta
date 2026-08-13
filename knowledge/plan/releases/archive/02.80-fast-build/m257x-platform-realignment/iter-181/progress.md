**Type:** tik (standard shape; §9 iter-type refinements consulted, none selected).

# iter-181 — the question was unanswerable until its DENOMINATOR was named

**Controlling strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase A — measure both denominators before answering

`SURVEY-M257x-iter179-readme-indexes-test-modules-unmeasured` asks how completely
`stack-core/README.md` indexes **test modules**. Measured at HEAD `3a8f5b4`:

| denominator | reading | is it the claim? |
|---|---|---|
| all `tests/test_*.py` on disk | **10 of 63** | **NO** |
| **mutation batteries** (`*mutation_battery*.py`) | **6 of 7** | **yes** |

**`10 of 63` would have been a published defect, not a finding** (`D-M257x-181-1`). The index's subject is
the fence family and its batteries; the other 53 are per-guard *behaviour* suites that the index
deliberately does not list — it lists the guard. Answering the survey on its own terms would have put an
84 % gap into a corpus whose entire quarrel is with numbers whose denominators nobody stated. §9 iter-159
— *grade the instrument at the grain of its claim* — turned on a **question**.

## Phase B — at the answerable denominator, a real gap

`test_repair_leak_guard_mutation_battery.py` shipped at **iter-48** (`932554e`, 2026-08-02) and had **no
row** at iter-180 — 20 mutants, 19 kills, and the only proof that the *"did this commit FINISH?"* fence
can fail. **Unindexed for 133 iters.** Nobody misread anything: there was no denominator, so there was
nothing to be short of (`D-M257x-181-2`).

## Phase C — the naive instrument was run, discarded, and kept as a control

The first resolve arm checked README-named `.py` files against `stack-core/` + `stack-core/tests/` and
reported `exposure_claim_guard.py` **missing**. It is not: it lives in `stack-injection/`, and the README
says so in a note under its own table. The arm is scoped to the **repo**, and the discarded version is
kept as a mutation control that goes RED **on purpose** if the cross-section reference ever disappears —
so the narrowing must be re-derived rather than silently adopted (`D-M257x-181-3`). A fence scoped more
narrowly than its subject manufactures findings: the instrument-side form of `D-M257x-122-4`.

## Phase D — what shipped

Into `test_fence_registry_population_m257x.py` — the module that already owns the README's disclosed
limit, so no new module and no new registry tax (iter-178's lesson, third consecutive iter):

- every mutation battery on disk **has a row**, with an anti-vacuity floor on the denominator itself;
- **the other direction** — the index names no battery that is gone;
- every README-named `.py` **resolves somewhere in rext**;
- two mutation controls: the repo-wide scope is load-bearing, and the battery arm fires on an absent row.

**RED-proven against the unrepaired tree first**, naming `test_repair_leak_guard_mutation_battery.py`.

## Phase E — deliberately NOT done

- `repair_leak_guard.py` (the **guard**) also has no row. That is 1 of the 11 members of
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27`, and repairing it would move a **published,
  fenced triple** while answering none of that survey's actual question — *which derivation is the index
  meant to be complete against* (`D-M257x-181-4`).
- `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) was the first candidate
  and is **set aside with a number**: **412 count-mentions across 115 files**, of which the defective
  subset — a `passed` count used *as the executed population* — is not regex-decidable, and iter-173
  already priced it as out of reach without re-running past refs (`D-M257x-181-5`).

## Runs — scope stated, and what it did NOT cover

| run | result | wall |
|---|---|---|
| `test_fence_registry_population_m257x.py` (the module carrying the new arms) | **16 passed** (was 11) | 8.17 s |
| + completeness + frozen-expectation census + battery-baseline-stage + `test_guard_family.py` | **119 passed · 0 failed** | 22.77 s |
| `guard_family.py` over the corpus | **18 GREEN · 0 RED · 8 not-run** | — |

Runner **pytest 8.4.2 on `/usr/bin/python3` 3.9.6**. **Not covered:** the rest of `stack-core` (green at
iter-179 — 1,521 P · 2 S · 0 F over everything but the batteries), the 7 mutation batteries, and the four
other rext sections. This iter adds tests and one README row and changes no guard logic. A scoped green is
evidence about its scope alone (rule 60).

## Close — 2026-08-09

**Outcome:** a one-iter-old survey is closed by **refuting the denominator it implied**. `10 of 63` — the
ratio the question asked for — would have published an 84 % gap that does not exist; the answerable
denominator is the batteries, where the real reading was **6 of 7** and the missing member had been
invisible since iter-48. The row is added and the arm that would have found it ships with it, RED-proven
first, alongside a resolve arm whose naive form was run, refuted by a real cross-section reference, and
kept as the control that keeps the correct scope honest.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirteenth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-181-1` … `D-M257x-181-5` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter179-readme-indexes-test-modules-unmeasured` — **CLOSED**, as ill-posed-until-scoped,
  with both denominators published and the wrong one named as wrong.
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — unchanged, **and one member is now named**:
  `repair_leak_guard.py`. Deliberately not repaired here.
- `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) — unchanged, **and now
  carries a number**: 412 mentions across 115 files, defective subset not regex-decidable.
- `SURVEY-M257x-iter179-thirty-battery-tests-unrun` · `SURVEY-M257x-iter180-relation-grammar-supports-
  only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a coverage question has no answer until it names its denominator — and the obvious
denominator is usually the wrong one.** Three iters running, this milestone has found the same shape:
iter-179's route mis-sized a 0.04 s check as fourteen minutes, iter-180's rationale covered two entries
with one sentence, and here a survey asked for a ratio whose numerator and denominator measure different
things. Each time the honest first move was the same and it was cheap: **measure both candidates, publish
the refuted one, then fence the one that can be answered.** Written into `platform-alignment.md` §8 in
this iter's commit.
