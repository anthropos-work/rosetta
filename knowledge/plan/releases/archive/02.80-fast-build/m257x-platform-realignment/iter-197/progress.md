**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

# iter-197 — the runner gap was an aggregate; now it is a per-module reading

## What was wrong

Harden pass 45 sized `FIX-M257x-h44-claim-census-guard-is-single-runner` like this:

> *"the 3,502/3,527 baseline gap above is its live size … and it reconciles."*

25 tests missing; that module has 25 tests. **The two facts are consistent and that is all they are.**
`suite_census.main` compares the runners on the module's **verdict** (`:736-742`, keyed on
`r["verdict"]` ∈ GREEN/RED/ENV-GATED/TIMEOUT) and prints each runner's test count as one `sum()` over the
whole population (`:733`). Two modules with **offsetting** differences reconcile to the same 25 and would
be invisible to both. This milestone named the shape at iter-192: **an agreeing reconstruction is
indistinguishable from a reading**, and `§9` says grade at the grain of the claim. The claim is about a
**module**; the evidence was an **integer**.

## What was built

A per-module **collection-reach** census in `suite_census.py` — `collected_by_pytest` /
`collected_by_unittest` / `collection_census` / `collection_disagreements` / `module_test_styles`, a
`RUNNER_COLLECTION_SPLIT` declaration graded both ways, a `--collect` report and a `--runner none` mode so
the reading is not priced at an execution census.

**The unit is COLLECTED tests and it is not the `tests=` column's unit.** That column counts tests that
EXECUTED (iter-172 settled it). Collection is the right grain for a *reach* question — does this runner
see the test at all — and it costs process startup, not a suite run. **Measured: 5.9 s** for 123 modules ×
2 runners across the pool. Cost was never the reason this went unfenced.

## The reading

`suite_census.py --collect --runner none`, this tree, both runners, Python only:

| | value |
|---|---|
| modules censused | **123** (the 5 collected sections; `§5` r60 — not the repo) |
| pytest collects | **3,551** |
| unittest collects | **3,526** |
| aggregate gap | **25** |
| **modules where the two DISAGREE** | **1** |
| the member | `stack-core/tests/test_claim_census_guard.py` — pytest **25**, unittest **0** |
| authoring style | **bare = 1 · testcase = 122** |

**h44's reconstruction is confirmed — as a reading.** One member, and the per-module gaps now sum to the
aggregate by assertion (`test_the_aggregate_gap_is_ACCOUNTED_FOR_by_the_per_module_reading`) rather than
by coincidence. The static half agrees independently: exactly one module of 123 is written as bare `def
test_` functions, and `unittest` collects `TestCase` and nothing else, so its style predicts its
invisibility before any test runs.

**Collected ≠ executed, and the difference is legible:** 3,551/3,526 here against pass 45's 3,527/3,502 —
**+24 on each side**, of which this iter's own new fence is 13. The remainder is the collected-but-not-
executed residue, identical on both runners, which is what a genuine unit difference looks like.

## Three defects, all in code written this iter, all caught by the controls written to prove it fires

1. **A silent zero, in the census built to find silent zeros.** `collected_by_pytest` scraped the summary
   line. On an unimportable module pytest prints **`no tests collected, 1 error`** and exits **2** — so the
   scrape matched `no tests collected` and returned **0**. A module pytest could not read would have been
   recorded as *a module with no tests* and summed into a clean total: `go_census`'s defect, re-created
   within the hour. Repaired by trusting pytest's **exit-code contract** (`PYTEST_RC_READABLE` — 0 = a
   count, 5 = a real zero, everything else UNREADABLE) instead of its prose.
2. **Two blind runners read as consensus.** `collection_disagreements` was `c["pytest"] != c["unittest"]`.
   For a module **neither** runner can read that is `-1 != -1` — **False** — so it left the disagreement
   set and was published inside a *"0 modules disagree"* line. Agreement between two instruments that
   cannot see is not agreement.
3. **A literal that rotted inside its own iter**, twice. `RUNNER_COLLECTION_SPLIT`'s reason read *"the
   repo's other **121** modules are `TestCase`"* — and this iter's own fence made it 122. And
   `test_the_authoring_style_census_agrees…` asserted `len(styles) - len(bare) == len(census) - len(bare)`
   — an **identity**, harden pass 45's finding re-created by hand two days later. Both repaired: the reason
   states only what does not change, and the arm now asserts the two censuses enumerate the **same
   population** (`set(styles) == set(census)`).

`§5` r77 is honoured explicitly: the mutation control is **size-preserving by assertion**
(`assertEqual(len(mutated), len(_BARE))`) and bumps mtime, because CPython invalidates on
(mtime-seconds, size) and a same-length edit is otherwise served from cache — the control would pass
without testing anything.

## What this iter did NOT do

**The conversion.** `FIX-M257x-h44-…` stays open, unchanged, un-re-routed. This iter gave it a correctly
grained size and a fence that fails the day a second member appears; the 25 pytest-fixture functions are
still collected by one runner. That separation is deliberate — the same one pass 45 drew between a
population and a verdict.

## Close — 2026-08-09

**Outcome:** the milestone's only open `FIX-` route was sized by an **aggregate reconciliation** —
3,527 − 3,502 = 25, and that module has 25 — which two offsetting members would satisfy identically. Now
read at module grain: **123 modules · pytest 3,551 · unittest 3,526 · exactly 1 disagrees**, and it is the
declared one (25 / 0), corroborated independently by a static style census (**bare 1 · testcase 122**).
The instrument found **three defects in its own code** on the way, each caught by a control written to
prove it can fire: a silent zero (`no tests collected, 1 error` → `0`), two blind runners scoring as
agreement (`-1 != -1` is False), and a size literal that rotted inside the iter that wrote it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-ninth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt**: this is the **first** tik of run 20 — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-197-1` … `D-M257x-197-5` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **181 passed · 1 skipped** across the
`suite_census` / `derivation` / `fence_registry` / `guard_family` modules; the new
`test_suite_census_collection.py` **13 passed under BOTH runners** (pytest 3.9.6 and unittest 3.9.6),
which is the property it exists to assert. *Scope: `stack-core` only, Python only, changed-code reach
(`§5` r60) — not the whole-section 1,699 of harden pass 47, and the other 10 sections remain unread. No
Go and no TypeScript was run this iter.*

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-h44-claim-census-guard-is-single-runner` — **open, unchanged, and now correctly sized.** Its
  size is a per-module reading rather than an aggregate reconciliation, and a second member can no longer
  appear silently. The conversion itself is untouched.
- `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` — **NEW.**
  `derivation_registry.population()`'s own docstring says *"every public collection-of-str-returning
  function"*, and `unclassified()` returned **0** for this iter's five new derivations because none of
  them returns a collection of `str`. `go_census` and `ts_census` (iters 195/196) are outside it for the
  same reason. The classification fence is therefore about **memberships** and says nothing about
  **counts** — which is harden pass 45's *"the memberships were derived and the SIZES were literals"* one
  level up, at the fence that grades derivations. A green `unclassified() == 0` names a denominator it
  does not print.
- Unchanged and still open: `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` ·
  `SURVEY-M257x-h45-printed-measurement-literals-uncensused` (this iter produced a live instance of the
  class and fixed it by hand, which is the third such hand-fix — the census is still not built) ·
  `SURVEY-M257x-h46-stale-substrate-direction-undeclared` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` ·
  `SURVEY-M257x-iter196-playthroughs-ships-no-node-modules` ·
  `SURVEY-M257x-iter195-the-go-reading-is-a-single-host-single-toolchain-sample` ·
  `SURVEY-M257x-iter194-other-milestones-ledgers-are-unaudited` · and the standing queue.

**Lessons:**
- **An aggregate that reconciles is not a reading of its parts.** One integer can be produced by any
  number of offsetting members; a claim about *which module* needs a per-module measurement. Adding it
  cost six seconds of runtime.
- **A summary line is prose; an exit code is a contract.** Parsing `no tests collected` turned a runner's
  own error into a zero. Where a tool publishes a return-code contract, read that.
- **Two instruments agreeing that they cannot see is not agreement** — and an inequality test over a
  sentinel silently converts blindness into consensus. Compare the sentinel explicitly.
