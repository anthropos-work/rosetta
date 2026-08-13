**Type:** tik (under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07))

# iter-248 — the runner that certifies this milestone's iters was overstating its own scope

## Phase A/B — the measurement, and its own invalidity, stated first

The full `stack-core` suite was run for the first time in this run's window:

> **16 failed · 2,064 passed · 3 skipped — 27 m 34 s** (`/usr/bin/python3 -m pytest tests/`, CPython 3.9.6)

**⚠ That reading cannot cleanly separate inherited REDs from self-inflicted ones, and the reason is this
iter's own doing.** The suite started at **04:10** and ran 27 minutes; **iter-248's own edits landed at
04:13–04:14**, three minutes in. **The tree moved under the instrument.** *An instrument that states its
own invalidity must not exit 0* — the milestone's own rule, and it applies to a measurement I took as
much as to one I audit. The 16 is reported and **the inherited/self-inflicted split is explicitly NOT
claimed from it**.

What IS solid is the repair and the re-measure. Every affected file was repaired and re-run:

> **251 passed / 1 failed** (first pass, subset) → after the last repair, **125 passed / 0 failed** across
> `test_fence_provenance` + `test_guard_family` + `test_test_collection_fence`, and the **guard family is
> 29 GREEN · 0 RED · 0 could-not-check · 5 not-run**.

### Two real defects, and the second one proves the route inside a single run

**1. Three test files defined classes AFTER their `if __name__ == "__main__":` guard** — so
`python3 <file>` collects none of them and still prints OK. `test_test_collection_fence` already existed
for exactly this and caught all three:

| file | class | whose |
|---|---|---|
| `test_anchor_construct_denominator.py:526` | `TestRangeBoundsArmM257xIter245` (6 tests) | **iter-245 — mine** |
| `test_guard_family.py:1027` | `ScopeDisclosureM257xIter248` (4 tests) | **iter-248 — mine** |
| `test_platform_alignment_guard.py:1054` | `ReachIsStated` (2 tests) | **inherited** |

All three guards moved to end-of-file. **I appended a class after the guard twice in one session**, which
is precisely the *"a repair to the instance, not the class"* shape iter-246 named — the fence caught it
both times, which is what a fence is for.

**2. `env_absence_guard` (iter-247, the previous iter) exited 2 instead of NOT-RUN.** It passed
`--platform` conditionally rather than declaring `"needs": ("platform",)` like every other
platform-facing guard, so a family run without `--platform` reported **could-not-check** — and
`--allow-not-run` does not cover could-not-check, so `test_fence_provenance`'s escape-hatch test went RED.

**That is `ROUTE-M257x-245`'s thesis demonstrated inside one session.** iter-247 closed reporting
`guard-family: 29 GREEN / 0 RED` — true, and it had just broken a test that `guard_family` does not run.
**Third occurrence: iters 239, 240, and now 247.**

## Phase C/D — the disclosure, so the conflation is unsayable

The tempting repair is to make `guard_family` run the suite. It is the wrong one (`D-M257x-248-1`): the
family is a fast tree-state check an iter runs many times; the suite is a 27-minute run. **One
measurement, two contracts** — the same shape as `D-M255-1`. What was broken was **legibility**.

`guard_family` now prints, on every run, green or red:

```
guard-family: SCOPE — this runs the GUARDS, not their test suite. 84 test file(s) under
stack-core/tests/ are NOT executed here; a green above is a statement about guard verdicts alone. Two
iters (239, 240) each left a different one of those files RED while this line read all-GREEN. Run
`python3 -m pytest tests/` before closing work that touched a guard, a fixture, or a cited corpus line.
```

The file count is **derived from disk** (M220 D1), and the line **names the two iters that paid for it**,
with a test asserting both numbers are present — evidence travels with the rule or a later tidy-up
deletes the rule as boilerplate. **4 net-new tests.**

## Phase E — pre-registrations

| id | prediction | outcome |
|---|---|---|
| **P-248-1** | 0–2 further inherited REDs | **1 identified** (`test_platform_alignment_guard`'s post-guard class) — **but the measurement was raced and the split is not claimed from it.** The number is a floor, not a reading |
| **P-248-2** | passing count ≥ 1,985 baseline + net-new | **HELD — 2,064 passed** |
| **P-248-3** | `guard_family` says nothing today about not being the suite | **CONFIRMED** — and now it does |
| **P-248-4** | the two repaired REDs stay repaired | **HELD** — both green in the targeted re-run |

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-245` closed at its cause. `guard_family` — the runner this milestone uses to
certify its own iters — reported a true verdict that read as a stronger claim than it made, and three
iters (239, 240, and **247, this session's own**) each left a `stack-core` test RED behind an all-GREEN
line. It now states its scope on every run, with a derived file count and the evidence attached. Two real
defects repaired: **3 test files with classes after the `__main__` guard** (two of them mine) and
**iter-247's guard exiting 2 instead of NOT-RUN**. Family **29 GREEN**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5
**Decisions:** `D-M257x-248-1` (the fix is a disclosure, not a runner change — one measurement, two
contracts) · `D-M257x-248-2` (the disclosure names the iters that paid for it, with a test) ·
`D-M257x-248-3` (a guard needing a reference DECLARES the need; it does not run without it and call the
result a check it could not do) · `D-M257x-248-4` (the whole-suite reading is reported and its
inherited/self-inflicted split explicitly NOT claimed — the tree moved under the instrument).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6): whole
section **2,064 passed / 16 failed / 3 skipped**, taken against a **moving tree** and therefore not a
clean verdict; after repair, the affected subset re-ran **125 passed / 0 failed** and the earlier subset
**251 passed / 1 failed → repaired**. Guard family (`--platform`, from repo root): **29 GREEN / 0 RED /
0 could-not-check / 5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-245-guard-family-green-is-not-suite-green` → **CLOSED at its cause** (the disclosure), with
  the third occurrence recorded. What it cannot do is force the run; that stays an operator discipline.
- `ROUTE-M257x-248-whole-suite-was-never-measured-cleanly-this-run` → **new.** The 16-failure reading is
  unusable for attribution because this iter's own edits raced it. **A clean whole-suite run against a
  frozen tree is owed**, and it is the first thing the next session should do — before any new work, so
  the tree is still the one the run describes.
- `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` → open (frontend scripts/ports).
- `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` · `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` ·
  `ROUTE-M257x-h59-range-anchors-are-ungraded` (which-line half) · `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` · `ROUTE-M257x-238-container-vs-native-is-undrawn` ·
  `ROUTE-M257x-237-hardcoded-vs-settable` · `ROUTE-M257x-236-disclosure-scope-is-document-level` ·
  `ROUTE-M257x-235-fence-scope-is-unread` · `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **Do not measure a tree you are editing.** A 27-minute suite run started three minutes before this
   iter's own commits is not a reading of anything; the honest report is the number plus a refusal to
   attribute it.
2. **The route proved itself inside the session that closed it.** iter-247 broke a test and closed
   reporting `29 GREEN / 0 RED`, truthfully. Three occurrences now, and the third was mine.
3. **A guard that needs a reference must DECLARE the need**, not pass it conditionally — otherwise
   "could not check" replaces "not run", and the two have different escape hatches.
4. **The fence that catches you is working even when it catches you twice in one session.** Both
   post-`__main__` classes were mine; `test_test_collection_fence` found both without being asked.
