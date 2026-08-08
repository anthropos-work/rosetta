**Type:** tik — under [`TOK-08`](../decisions.md), closing the RED iter-166 verified and routed.

# iter-167 — a frozen fixture, a live ledger, and a fence that was right all along

## Phase A — reproduce, and name the firing claim exactly

Two hits on the GREEN twin, both the same claim:

```
FIRED: corpus/04.md | claim_id: …/iter-49/raw/C.md#c-1@57
FIRED: corpus/05.md | claim_id: …/iter-49/raw/C.md#c-1@57
  matched: ent privacy policies auto-filter by organization on only 31 of 135 schemas …
```

Both fixtures are captures of `corpus/architecture/architecture_overview.md:298`, adjudicated by
**iter-48**. The claim firing on them was adjudicated by **iter-49** — *after the capture* — and it
is a different refuted form about the same region: the *"only 31 of 135 schemas"* count, which
iter-49 measured as **32** (`app/internal/data/ent/schema/organization.go:56`, `:94-97`).

## Phase B — which of the three subjects is at fault

| candidate | verdict |
|---|---|
| the fence is brittle | **NO.** It fired on a sentence that IS a refuted claim. It was right. |
| the corpus regressed | **NO.** The live tree is clean — `claim-twin-guard: OK — none of the 264 adjudicated claims is published anywhere in the tree`. The text lives only in the frozen fixture. |
| **the fixture's expectation is stale** | **YES.** |

The fixture is pinned at rosetta `cabc3b1`; the ledger is **re-derived on every run** from the
milestone's blocker-ledgers and has grown by ~117 iters of adjudications since. So the capture's
"green" meant *carries no refuted claim **known then***, and the assertion was written as though it
meant *carries no refuted claim*. Those came apart the moment a later iter re-adjudicated the same
region on a different form — which is not an exotic event in this milestone, it is the milestone's
normal behaviour.

> **A frozen expectation and a live derivation are two clocks. The assertion between them must state
> which one it is read against, or it is only true until the other moves.**

This is the standing route `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance`, arriving
as a concrete failure rather than a survey.

## Phase C — repair, without spending the answer key

Two repairs were available and rejected:

- **Edit the fixture** so the green twin no longer carries iter-49's sentence. Forbidden — `§5` rule
  21: the capture is perishable, exists to support a claim about the *instrument* (a seven-auditor
  read missed these while they sat in its own assigned file sets), and cannot be re-taken.
- **Delete or soften the assertion.** That discards the only control distinguishing a discriminating
  fence from a brittle one.

Landed instead: `test_02` grades against **the capture's own denominator** — the adjudicating ledger
files named in the fixture manifest, `CAPTURED_LEDGERS`, **derived from the manifest, never listed**.
A hit from one of those on the green twin is brittleness and still fails. A hit from a later ledger
is the fence doing its job.

**And the residual is asserted, not swallowed.** Every excluded hit must come from an iter **newer
than the capture** (`CAPTURE_ITER = 48`). Without that clause the scope would quietly absorb a real
in-capture miss that happened to arrive by a different path.

## Phase D — prove the narrowing is not vacuous

The escalation condition this iter opened with, discharged two ways:

1. **`test_02b`, new and permanent.** The *same scoped predicate*, applied to the **RED** fixture,
   must find the captured claim at **every** site. If the fence ever stops detecting them, `test_02`
   would pass vacuously and `test_02b` fails instead. Green.
2. **A mutation, run now.** Raising `CAPTURE_ITER` to 50 makes iter-49's hits inadmissible, and the
   residual clause fires exactly as designed:
   `AssertionError: 49 not greater than 50 : corpus/04.md:1 fired from iter-49, which is NOT newer
   than the capture`. **The clause is live, not decorative.**

## Gates

**Run, green:** `test_claim_twin_guard_iter48_answer_key` (7, was 6 — `test_02b` is net-new) ·
`test_claim_twin_guard` · `test_claim_twin_guard_iter47_answer_key` · `test_waiver_ledger_m257x` ·
`test_frozen_expectation_census_m257x` — **72 tests, 0 failures.**

**The family, at HEAD:** `guard_family.py` → **17 GREEN · 0 RED · 0 could-not-check · 7 not-run**
(each not-run named, each for want of a `--range`/`--platform`/`--ledger` input the tree cannot
supply). Combined with iter-166's fix to the value-change mutant, **the "RED at HEAD since iters
162/163" class is now clear** — and it is worth stating that this was checked with a run rather than
inferred from two repairs.

**NOT re-run, named in full (`§5` rule 60):** the rest of `stack-core`, and `stack-seeding`,
`stack-snapshot`, `stack-verify`, `playthroughs`. This iter touched exactly one test module.

## Close — 2026-08-08

**Outcome:** the RED at HEAD is closed, and the fence is exonerated. `test_02`'s green-twin assertion
joined a **frozen fixture** (rosetta `cabc3b1`) to a **live, re-derived ledger**; iter-49
re-adjudicated the same corpus region on a different form, so the fence fired correctly on a real
refuted claim and the *assertion* was what had rotted. Scoped to the capture's own denominator
(derived from the manifest), residual asserted rather than swallowed, and the narrowing proved
non-vacuous two ways. Family at HEAD re-measured: **17 GREEN · 0 RED**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this iter and iter-166 both closed
`closed-fixed`; no no-prog streak, and no `N` reading was taken so the metric is UNMEASURED not
unmoved — `§9`) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n
— (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-167-1` … `D-M257x-167-4` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none.
**Routes carried forward:**
- `SURVEY-M257x-iter167-the-other-two-answer-keys-have-the-same-coupling` — **NEW.** iter-41's 18 and
  iter-47's 7 are frozen against the same live ledger. Neither is RED today; `D-M257x-167-1` says
  that is a matter of which regions later iters happened to re-adjudicate, not of design. Their
  green-twin assertions should be scoped the same way **before** one of them goes RED and is shipped
  over.
- `FIX-M257x-iter166-iter48-answer-key-red-at-head` — **CLOSED by this iter.**
- `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` — **substantially advanced**, not
  closed. This iter supplies the concrete failure the survey was queued to look for, and a working
  pattern; the survey's scope is every frozen artifact in the suite, which is wider.
- `FIX-M257x-iter166-stage-derivation-covers-code-not-data` · `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer`
  and the standing queue, unchanged.
**Lessons:** **when a fence and a fixture disagree, there are THREE subjects, not two** — the fence,
the corpus, and the *assertion joining them* — and this milestone's reflex has been to suspect the
first two. Here the fence was right, the corpus was clean, and the assertion had silently changed
meaning because one of its two inputs kept moving. Naming all three before repairing is what kept a
one-minute fixture edit from destroying a perishable answer key.
