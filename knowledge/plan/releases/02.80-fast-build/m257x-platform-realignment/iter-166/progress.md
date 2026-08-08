**Type:** tik — under [`TOK-08`](../decisions.md), landing the route iter-165's falsification sharpened.

# iter-166 — the accept side had never been printed, and dormancy is not one question but three

## Phase A — census the accept-side mechanisms (denominator stated)

Not a hand-list. Every consumer declares its waiver file in a module constant, so the population is
readable off the source:

```
grep -nE '^(WAIVERS_REL|WAIVER_FILE)\s*=' stack-core/*.py
```

**4 of 4**, and the grep is the completeness argument: a fifth consumer cannot exist without the
constant. Reach denominator: the 20 `*.py` files under `stack-core/` that mention `waiv|exempt`
(the superset), of which four *own* a waiver file.

| guard | file | entries | scope |
|---|---|---|---|
| `claim_twin_guard` | `claim_twin_waivers.json` | 17 | live corpus |
| `repair_reach_guard` | `repair_reach_waivers.json` | 6 | one named ledger |
| `repair_leak_guard` | `repair_leak_waivers.json` | 1 | a diff |
| `value_change_guard` | `value_change_waivers.json` | 1 | a diff |

## Phase B — the defect, and a promise the code never kept

Three of the four never named a honoured waiver anywhere. `claim_twin_guard` came closest and it is
the instructive case: it printed **`22 acknowledged site(s) skipped`** and emitted `"waived": 22` in
JSON — a **count**, which cannot distinguish a waiver that fires every run from one dead for months.
`repair_reach_guard` carried each honoured disposition into its per-anchor render, but never named a
waiver that matched **nothing**.

And `repair_leak_waivers.json`'s own `_README` has always said:

> "A waiver can only ever make the fence QUIETER, so **every one is reported on each run**"

The guard filtered at `find_leaks` and printed no waiver identity anywhere. **The document described
the property; nothing implemented it** — the `r70/71` shape on a new axis: the *doc* was pinned to
the property, and the code was pinned to nothing.

## Phase C — the repair: one matcher, read twice

`stack-core/waiver_ledger.py` (new). Each guard's `is_waived` now delegates to a `_waiver_match`
returning the KEY it matched on, so the suppression decision and the report **are the same decision**
— never a second predicate written alongside, which is precisely what iter-165 withdrew 11 findings
for. All four wired through an optional `accept_ledger` OUT parameter, so no existing caller's
signature or return shape changed.

## Phase D — the readings, and the near-miss that became the design

**`claim_twin_guard`, live corpus:** `waivers: 17 of 17 honoured (22 candidate finding(s) graded); 0
dormant`. Internally consistent — the per-key honour counts sum to 22, the guard's own long-published
"acknowledged sites skipped" number. This **converts iter-165's withdrawn-and-uncertain "0 provably
dead" into a measured 0**, taken by the only instrument entitled to take it.

**`repair_reach_guard` — and this is the finding.** Run against iter-76's ledger + `328ece5`:

```
waivers: 0 of 6 honoured from repair_reach_waivers.json (152 candidate finding(s) graded); 6 dormant
```

**152 candidates graded.** Not vacuous by any measure this family had. It reads exactly like six dead
waivers, and iter-166 was one edit away from reporting it as such. Run against **iter-86's** ledger +
`ae5c1db` — the run the file's own `_README` says it was written for:

```
waivers: 6 of 6 honoured from repair_reach_waivers.json (46 candidate finding(s) graded); 0 dormant
```

Nothing about the waivers changed. The cause is a **schema difference with no shared layer**:
`claim_twin` / `repair_leak` / `value_change` waivers key on `path` + a quoted form and are
**subject-independent**; `repair_reach` waivers key on `path:line` **coordinates into ONE ledger's
anchors** — a per-run disposition set stored in a standing-config slot. So the ledger grew a fourth
state: a `subject_scoped` guard **names its subject and may never report bare dormancy**.

> **A non-zero candidate count is not enough. Dormancy is evidence about a waiver only when the
> run's SUBJECT is the one the waiver was written for.**

`repair_leak` / `value_change` outside a repair report **VACUOUS** by name — *"the guard graded 0
candidates this run, so dormancy here is evidence about the RUN, not the waivers"* — which is the
`§9` zero-census rule applied to the accept side.

## Phase E — fencing it, and two defects the fencing found in the fences

`tests/test_waiver_ledger_m257x.py` (new, 9 tests, green): the state machine; the refusal to report
before the candidate count is known; the refusal to honour an unregistered key; and — the one worth
the module — **both `repair_reach` readings pinned to their two committed artifacts**, so the
near-miss above cannot be re-made. The four guards are asserted against **their own output**, because
a test that exercised only `waiver_ledger` would stay green while every call site was deleted.

Wiring the guards turned **three** existing fences RED, and none was noise. The third is the one to
note first, because it is the family working exactly as designed: **the derivation-registry
completeness fence caught the new module before any test of mine ran**, naming all three of
`WaiverLedger`'s reading accessors as unclassified. That is iter-162's rule paying out — the registry
is enumerated, so a net-new file cannot join the tree unclassified. Classified `DECLINE:instance-state`
/ `DECLINE:verdict`, with the reason recorded at the entry. The other two were defects:

1. **`test_value_change_guard`'s pass-through mutant took a `TypeError`.** Its double
   re-declares `find_survivors`' signature; widening the real function widened its double too, and
   without that the battery reported **ERROR rather than a mutant verdict** — the harness had stopped
   measuring the guard. Forwarded, with the reason recorded at the site.
2. **`test_repair_leak_guard_mutation_battery` staged its dependencies from a HAND-LIST of four.**
   The guard grew a fifth (`waiver_ledger.py`), the staged suite died on ImportError, and the battery
   reported its **baseline RED** — so all five mutant verdicts became uninterpretable *while still
   looking like real kills*. **This is iter-162's rule one layer down** (*fence a registry's
   completeness, never its contents*), and the fix is the same shape: `_local_deps()` now **derives**
   the stage list transitively from the guard's and suite's own imports. It reproduces the original
   four and picks the fifth up on its own.

The battery's needle-match assertion also fired correctly on the `json-leaks-dropped` mutant when the
JSON face grew its `accept` block — *a mutant whose needle no longer matches is an UNAPPLIED mutation,
not a surviving one*, and the battery already refused to grade the difference. Needle re-pinned.

## Gates

**Run, green:** `test_waiver_ledger_m257x` (9, new) · `test_claim_twin_guard` ·
`test_claim_twin_guard_iter47_answer_key` · `test_value_change_guard` · `test_repair_reach_guard` ·
`test_repair_leak_guard` · `test_repair_leak_guard_mutation_battery` (6, ~120 s) ·
`test_frozen_expectation_census_m257x` · `test_fence_registry_completeness_m257x` ·
`test_guard_family` · `test_fence_provenance` · `test_iter45_mechanical_fences`
(the last five as one 175-test batch).

**RED and NOT caused by this iter — pre-existing at HEAD, verified:**
`test_claim_twin_guard_iter48_answer_key::test_02_the_green_twin_of_every_site_stays_SILENT`. Checked
by staging `git archive HEAD stack-core` into a scratch tree and re-running there: **it fails
identically with none of this iter's code present.** This is the class the last harden named — *two
fences had been RED at HEAD since iters 162/163 and three iters shipped over them.* Routed, not
silently absorbed, and **not repaired here**: it is a different subject and repairing it inside this
iter is the scope-creep tripwire's third line.

**NOT re-run, named in full (`§5` rule 60):** the rest of the `stack-core` suite, and every other
rext section — `stack-seeding`, `stack-snapshot`, `stack-verify`, `playthroughs`. This iter touched
five files, all inside `stack-core`, and nothing outside it imports them; but that is an argument,
not a measurement, and it is stated as one.

## Close — 2026-08-08

**Outcome:** the ACCEPT side of the fence family is reported for the first time. A shared
`waiver_ledger` wired through all four waiver-carrying guards **using each guard's own matcher**;
first readings: `claim_twin` **17 of 17 honoured / 0 dormant** over 22 candidates (a measured
confirmation of iter-165's withdrawn "0 provably dead"), `repair_reach` **6 of 6** against its own
subject. And a fourth state the design did not start with: a **non-vacuous** run can still produce
meaningless dormancy — the same 6 waivers read **0 of 6 over 152 candidates graded** against a
foreign ledger. Three existing fences went RED on the wiring; one was the registry-completeness fence
working as designed, two were real defects in the fences' own harnesses.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (no 3-tik no-prog streak; and no `N`
reading was taken, so the metric is UNMEASURED not unmoved — `§9`, and `TOK-08`'s sealed rule still
forbids a successor strategy) — (3) re-scope: n — (4) user-blocker: n (the pre-existing RED is a
routed Fate-3 item with a named handler, not a decision the user must make) — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-166-1` … `D-M257x-166-4` (see [`decisions.md`](decisions.md))
**Side-deliverables:** the `_local_deps()` derivation in the repair-leak mutation battery and the
`derivation_registry` classifications are in-scope Phase-E work, not side discoveries — both were
required to leave the suite green.
**Routes carried forward:**
- `FIX-M257x-iter166-iter48-answer-key-red-at-head` — **NEW.**
  `test_claim_twin_guard_iter48_answer_key::test_02` is RED at HEAD, verified pre-existing against a
  scratch tree staged from `git archive HEAD`. The harden pass's *"RED at HEAD and shipped over"*
  class, now with a named handler.
- `FIX-M257x-iter166-stage-derivation-covers-code-not-data` — **NEW.** `_local_deps()` follows `.py`
  imports; a data dependency (waiver JSON, fixture manifest) would still be silently omitted.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` — **new hard evidence, second
  consecutive iter.** iter-165 found three waiver schemas cost it a whole reading; iter-166 found the
  schema difference changes what a REPORT MEANS (`D-M257x-166-2`). `waiver_ledger.py` is the first
  piece of that shared layer to actually exist.
- `FIX-M257x-iter165-fences-do-not-report-which-waivers-they-honoured` — **CLOSED by this iter.**
- Unchanged and still queued: `SURVEY-M257x-iter164-verification-662-claim-not-adjudicated` ·
  `FIX-M257x-iter163-block-ref-attaches-the-wrong-sha` ·
  `SURVEY-M257x-iter163-anchors-with-no-quoted-literal` · `-iter163-generic-literals-are-unadjudicable` ·
  `SURVEY-M257x-iter162-a-literal-has-a-ROLE-the-census-cannot-see` ·
  `-iter162-small-derivations-are-coincidence-prone` ·
  `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` ·
  `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` ·
  `SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality` ·
  `FIX-M257x-iter160-b2-over-strict-direction-still-unfenced` ·
  `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` · `SURVEY-M257x-iter156-other-reporting-layers` ·
  `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter133-two-fives-need-a-fence` · `-iter131-predicate-sets-not-enumerated`
**Lessons:** **a fence publishes its FIRE side and hides its ACCEPT side, and only one of those is
watched.** Every reach number in this family answers *what did the guard catch*; nothing answered
*what did it agree to ignore, and does that agreement still apply*. The second question is as
mechanical as the first — but only the guard can answer it, because the answer is a by-product of a
predicate no outside instrument may re-implement.

And the correction that arrived within the iter: **VACUOUS was not enough.** A zero-candidate run is
the obvious way for dormancy to mean nothing; a run that grades 152 candidates against the wrong
*subject* is the non-obvious way, and it produces a number that looks like a finding. Both readings
are now pinned to commits, so the next reader meets the trap already sprung.
