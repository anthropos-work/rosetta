# iter-112 — TOK-07 step 1: the enumerator

**Type:** tik · `iter_shape: tooling`. **Active strategy:** `TOK-07`, step 1.

## What landed

**`stack-core/predicate_enumerator.py`** (`FENCE-M257x-iter112`) — a per-predicate, corpus-wide site
sweep that runs **before** any repair and produces the repair's denominator.

- **Judgement/mechanism boundary drawn explicitly** (`D-M257x-112-1`): choosing a form is judgement
  (derived by default, authored where the proposition is prose); enumerating it is mechanical and
  complete, over raw lines **and** `claim_ledger.normalize_document`'s re-flowed text, because corpus
  prose wraps.
- **Seed recall is FAIL-CLOSED** — a form that cannot find the site it was written for is a RED. This
  is the control that fires when the *enumeration* is what is broken.
- **An underivable predicate is exit 2, never 0 sites** — *"0 sites"* is the most convincing possible
  way for a check that skipped to read like a check that passed.
- **The multiplier is printed per predicate, always, and `NO-EXPANSION` is NAMED** rather than left to
  be inferred from equal numbers — iter-108's whole failure was an absent step reading as a satisfied
  one.

**`iter-112/predicate-ledger.json`** — iter-109's **24 predicates**, one row per *proposition* (never
per anchor), seeds = the 29 anchors the reading DETECTED, marked as inputs rather than as the
denominator.

**18 tests**, controls per `TOK-06`'s binding clause, each shown firing against a mutated copy:
deleting the seed-recall check makes a missed seed invisible; restricting the scan to the seed's own
file — *exactly what a detection-bounded repair does* — collapses the multiplier to 1.0; re-adding the
neighbour padding re-imports the adjacent proposition's token. Anti-vacuity: the fixture world is
asserted to actually publish the twin in three files, and the derivation is asserted to return
**nothing** for plain prose — the premise the exit-2 refusal rests on.

Covered by iter-111's machine-mode contract from its first commit (`D-M257x-112-5`): added by name to
`test_fence_provenance.py::_guards_declaring_json`, though deliberately **not** in `guard_family`'s
census, because it is ledger-scoped and cannot run without one.

Invocations, stated with their counts: `pytest tests/test_predicate_enumerator.py -q -p
no:cacheprovider --no-header` → **18 passed** (0.16 s); with `tests/test_fence_provenance.py` →
**52 passed** (84.78 s).

## The fence caught its own derivation, twice, before any number shipped

`D-M257x-112-2`. Run 1: **1 refusal + 5 seed-recall REDs** — the derivation padded each seed by a
neighbouring line and pulled tokens off *adjacent* propositions (`P04`, a speech-model claim, derived
`studio/tools/pdf2md.py`). Fixed by reading the **booked range and not one line more**, plus
first-class range seeds. Run 2: **36.07×** — a number about the English language, because an uncapped
derivation enumerates a seed line's whole vocabulary. Fixed by specificity ranking + a 4-form cap.
**Neither wrong number reached a report.** That is what the controls were written for.

## The measured reach, and it decides how step 2 must be done

`D-M257x-112-3`. **22 of 24 predicates needed an AUTHORED form**, because a large share of this residual
is **prose, not citations** — `TTS v2 HD`, `Cosmo Router`, *"the split is on the endpoint only, not on
the agent name"*. A token-derivation cannot reach a proposition with no literal in it. **This is the
same boundary `claim_twin_guard` already draws** (it matches quoted verbatim forms from an authored
ledger, for exactly this reason) — now **measured** for this residual rather than assumed. `P16` and
`P24` are left derived deliberately, as the control that the derived path still runs.

## The first per-predicate multiplier this milestone has reported — and it is not yet trustworthy

`D-M257x-112-4`. **29 seeds → 211 sites → 7.28×**, seed recall **100 %**, 40 files in scope.
Artifacts: `enumeration.txt`, `enumeration.json`.

The headline is the least useful line on the sheet:

- **12 of 24 read `NO-EXPANSION` (×1.0)** — and by `TOK-07`'s own guard-rail those are **12 verdicts
  against this ledger's forms**, not 12 rare predicates;
- **4 read implausibly broad** (`P16` ×48, `P22` ×37, `P18` ×29, `P24` ×19) — `Cosmo Router` at 37 sites
  is 37 mentions of a deleted component, not 37 publications of the proposition;
- **the credible middle is real**: `P10` ×10, `P09` ×6, `P21` ×6, `P15` ×4, `P06` ×3, `P03`/`P07`/`P20`
  ×2 — **every one a site iter-109 did not book**, i.e. exactly the twin population `D-M257x-109-4`
  predicted and nothing had ever enumerated.

**The instrument lands; the measurement does not.** Step 2 must not repair against this ledger as it
stands — repairing 211 sites of which some are vocabulary would be worse than repairing 46.

## Close — 2026-08-06

**Outcome:** the enumerator ships with its controls **proven firing on real data** — it refused its own
first derivation twice before any number existed — and produces the first per-predicate multiplier in
this milestone: **29 seeds → 211 sites, 7.28×, seed recall 100 %**. The measurement is **explicitly not
yet trustworthy**: 12 predicates read `NO-EXPANSION` and 4 read as vocabulary, both of which
`TOK-07`'s guard-rail scores against the *forms*, not the predicates. Named per predicate rather than
averaged away, and routed.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — (7) budget-exhausted: **y — between iters, both trees committed and pushed** — Outcome: exit-7
**Decisions:** D-M257x-112-1 … D-M257x-112-5
**Side-deliverables:** none
**Routes carried forward:**
- `FIX-M257x-iter112-forms-need-a-second-pass` → **next iter**, and it BLOCKS step 2: per-predicate form
  review with the two failure shapes now named (too-narrow → `NO-EXPANSION`; too-broad → vocabulary)
- `TOK-07` step 2 (repair whole predicates against the enumerated set) → after the form review
- `TOK-07` step 3 (the read) → last, unchanged
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → open
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → open
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → open, de-ranked
- `DEF-M257x-iter101-briefing-rext-tree` → open
**Lessons:**
- **Draw the judgement/mechanism boundary out loud, then fence the judgement half.** "Mechanically
  enumerate a predicate" is not a thing; "derive the form where a literal carries it, author it where
  prose does, and fail closed when the form cannot find its own seed" is.
- **A fence's first live run is the only time its controls are graded against something nobody tuned
  them for.** This one refused twice, and both refusals were correct.
- **Report the number that indicts you.** 12 `NO-EXPANSION` verdicts are the honest headline; 7.28×
  is the flattering one, and it is the average of a good middle, a narrow tail and four forms measuring
  vocabulary.
