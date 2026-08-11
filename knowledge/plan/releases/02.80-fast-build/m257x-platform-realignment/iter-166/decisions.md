# iter-166 — decisions

## `D-M257x-166-1` — the accept side is reported BY THE GUARD, never by a second instrument

iter-165 audited the four waiver files with an auditor it wrote alongside them, and withdrew 11 of 11
findings as artifacts of its own preprocessing. The conclusion it reached — *ask the guard* — is
adopted as the design rule here and implemented literally: each guard's `is_waived` now delegates to a
`_waiver_match` that returns the key it matched on, and the report is fed from that return value. **The
suppression decision and the accept-side report are the same decision read twice.** A future change
that reports the accept side from a separately-written predicate re-opens the iter-165 failure mode and
should be refused on sight.

## `D-M257x-166-2` — dormancy has THREE preconditions, not one, and the third was nearly missed

A dormant waiver is evidence that it may be dead **only** when (a) the guard ran, (b) it graded ≥ 1
candidate, and (c) **the run's subject is the population the waiver was written against.**

(c) is net-new and was found the hard way. `repair_reach_waivers.json`'s six entries read **0 of 6
honoured over 152 candidates graded** against iter-76's ledger + `328ece5` — a full, non-vacuous
population — and **6 of 6** against iter-86's ledger + `ae5c1db`. Nothing about the waivers changed.
Its keys are `path:line` coordinates into ONE ledger's anchors; the other three files key on
`path` + a quoted form and are subject-independent. So `WaiverLedger` carries `subject_scoped`, and a
subject-scoped ledger **may not print bare dormancy** — it names its subject and says the caveat in
full, every time. Both readings are pinned to their commits in
`tests/test_waiver_ledger_m257x.py::test_08`/`test_09`.

**This is also new hard evidence for `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer`:**
the three waiver schemas are not a cosmetic inconsistency, they change what a report MEANS.

## `D-M257x-166-3` — a mutation battery that stages a hand-list has stopped measuring its guard

`test_repair_leak_guard_mutation_battery` staged four dependency files by name. The guard grew a fifth
and the staged suite died on ImportError, so the battery reported its **baseline RED** and all five
mutant verdicts alongside it — *uninterpretable, while still looking like real kills*. A stager that
can silently omit a dependency is the iter-162 defect one layer down. `_local_deps()` now derives the
set transitively from the guard's and the suite's own imports; it reproduces the original four and
picks the fifth up unaided.

**Disclosed limit, routed:** the derivation follows `.py` imports only. A *data* dependency (a waiver
JSON, a fixture manifest) would still be missed. Not closed here — `repair_leak_guard.load_waivers`
treats an absent file as "no waivers, the honest default", so the battery is correct today, and
widening the derivation to data files is a separate change with its own failure modes.

## `D-M257x-166-4` — the pre-existing RED was VERIFIED pre-existing, not assumed

`test_claim_twin_guard_iter48_answer_key::test_02` is RED. It would have been easy, and wrong, to
assume it was mine or assume it was not. It was checked: `git archive HEAD stack-core` into a scratch
tree, suite re-run there, **fails identically with none of this iter's code present.** This is the
class the last harden pass named — two fences RED at HEAD with three iters shipped over them. Routed
as `FIX-M257x-iter166-iter48-answer-key-red-at-head`, **not repaired inside this iter**: it is a
different subject, and taking it would be the scope-creep tripwire's third line.
