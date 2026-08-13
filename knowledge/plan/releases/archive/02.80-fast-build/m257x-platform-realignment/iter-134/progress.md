**Type:** tik · `iter_shape: audit`

# iter-134 — the conjecture was wrong, and being wrong is the useful part

`FIX-M257x-iter132-marker-fences-cannot-see-retractions` said the blindness *"plausibly affects every
marker-matching fence."* Measured: **it affected exactly one — the one that had it.**

---

## 1. The measurement

Each of the four prose-classifying fences was **imported and interrogated**, not grepped (rule 22 — a
grep of a file is testimony about its text; a loaded module is the thing that runs). Five independent
retraction mechanisms:

| fence | `RETRACTION_MARKERS` | retraction-context predicate | past-tense exclusion | retracted/refuted in vocab | waiver + decay | **score** |
|---|---|---|---|---|---|---|
| `claim_twin_guard` | YES | YES | YES | YES | YES | **5 / 5** |
| `claim_census_guard` | no | no | YES | YES | no | 2 / 5 |
| `platform_predicate_guard` | no | no | YES | no | no | 1 / 5 |
| `unreadable_repo_claim_guard` *(pre-iter-132)* | no | no | no | no | no | **0 / 5** |

**The pre-stated branch:** ≥ 2 blind → pattern, route upheld. **Measured: 1.** → **the conjecture is
REFUTED**, and the route closes on a falsification rather than on four fixes.

**The partial scores are not gaps, and the audit had to check that before calling them coverage.**
`claim_census_guard` measures **unevidenced, never false** — it says so in its own printed banner — so a
retraction carrying a ref is *evidenced* by construction and never enters its numerator; the two
mechanisms it has are the two its job needs. `platform_predicate_guard` grades corpus prose against
platform config and carries an explicit exclusion for `not|never|no longer|formerly|was`
(`platform_predicate_guard.py:900`) plus a documented section on **historical vocabulary** (`:160`) and
historical-sha handling (`:1092`, `:1102`) — the past-tense axis *is* the retraction axis for what it
measures.

## 2. The finding the audit actually produced

**`claim_twin_guard` has shipped the exact capability since M257x iter-48** — `RETRACTION_MARKERS` (14
tokens), `_looks_retracted`, a **320-character** context window, and a waiver file that **decays to
nothing if the retraction is ever deleted**, with its own docstring naming the hazard: *"A site may
legitimately quote a refuted claim in order to retract it."*

**iter-132 hit that exact problem and built a bespoke third bucket instead of reusing it.** It reached
`D-M257x-121-4`'s *disclose-the-floor* ruling independently — which the iter recorded as a virtue — and
the less flattering half is this: **the family already contained a tested, decaying, waiver-backed
solution to the same problem, in a fence that runs beside it in the same runner.** The iter did not
look.

> **This is a fence-family reuse gap, not a fence bug**, and it is why the route closes with a *new*
> route rather than none. `unreadable_repo_claim_guard`'s `mixed` bucket currently means *"a marker plus
> **any** ref-pinned reading in the paragraph"*. `claim_twin_guard`'s test is sharper — *"a marker
> inside an explicit **retraction context**"* — and it decays. **The bucket is honest but coarse, and it
> is coarse because it was invented rather than borrowed.**

## 3. What landed, and what deliberately did not

**Landed:** the audit, its table, and the refutation — recorded here and in `decisions.md`, with the
route closed against a measurement instead of left open on a conjecture.

**Deliberately NOT landed — routed with a named handler:** the refactor that would have
`unreadable_repo_claim_guard` consume `claim_twin_guard`'s retraction predicate. It is the right change
and it is **not a one-line import**: the two fences are independent modules in a family whose runner
loads each standalone, so sharing the predicate means either a cross-fence import (a new coupling in a
family that currently has none) or a shared module (a structural change to the family). **Choosing
between those is a design decision, and this milestone has eight vacuous fences on record from making
one in a hurry** — the same reason iter-133 declined to build a prose fence on the spot.

## 4. Test gates

- **Guard family: 18 GREEN · 0 RED · 4 not-run** (commit-/input-scoped members without
  `--range`/`--ledger`). Not a whole-family green; the runner says so.
- **Zero files changed in `rosetta-extensions` and zero in `corpus/`** — this iter is an audit. Its
  only artifacts are milestone records, so no code-test gate applies and none is claimed.
- **The whole suite was not re-run, and §5 rule 60 requires saying so.** Nothing executable changed
  since iter-132's clean run (`1 failed · 1208 passed`, the 1 being the standing RED) — iters 133 and
  134 modified **no** `rosetta-extensions` file. **Stated as a gap, not characterised as covered.**

---

## Close — 2026-08-08

**Outcome:** the conjecture that iter-132's retraction blindness was a family-wide pattern is
**REFUTED by measurement against a pre-stated branch** — of four prose-classifying fences, **one** was
blind and it is the one already fixed; `claim_twin_guard` scores **5/5** and the other two carry exactly
the mechanisms their subjects need. **The audit's real product is the reuse gap it exposed:**
`claim_twin_guard` has shipped a tested, decaying, waiver-backed retraction predicate since iter-48, and
iter-132 re-invented a coarser one beside it without looking.
**Type:** tik
**Status:** closed-fixed *(the planned deliverable was an audit with a verdict; it landed, and the verdict is a refutation)*
**Gate:** NOT MET — **4 of 5**, unchanged; no reading was taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no successor strategy is authorable — `TOK-08`'s sealed refutation branch bars one; running under the user's direct brief**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-134-1`, `D-M257x-134-2` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter132-marker-fences-cannot-see-retractions` — **CLOSED on a refutation**, superseded by:
- **NEW — `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer`:** the family's fences each
  re-implement prose predicates (retraction detection, past-tense exclusion, hedge vocabularies) with
  **no shared layer and no index of what already exists**, so the next fence author will re-invent the
  next one too. The fix is a design decision (cross-import vs shared module), deliberately not taken
  under time pressure.
- Still open and untouched for **three** consecutive iters: `FIX-M257x-iter131-adjudication-independence`
  — **the oldest unactioned route on the milestone**, and the one iter-131 itself called *"the first
  item the next iter should action."* Named here because three iters have now passed it over.
- `FIX-M257x-iter131-predicate-sets-not-enumerated` · `FIX-M257x-iter131-root-mount-count-underived` ·
  `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it` ·
  `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` · `FIX-M257x-iter133-two-fives-need-a-fence`.
**Lessons:**
1. **A route written as a conjecture is a measurement nobody took.** *"Plausibly affects every…"* cost
   one iter to check and was wrong. It would have cost four iters to act on as written.
2. **Check the family before you invent.** The capability iter-132 built existed, tested and sharper, in
   a fence loaded by the same runner. Neither the fence's docstring nor the protocol pointed to it —
   which is the gap, not the author's diligence.
3. **An audit that refutes its own premise is a first-class deliverable**, and grading it `closed-fixed`
   rather than `closed-no-lift` is correct: the *planned* scope was the verdict, not the fixes.
