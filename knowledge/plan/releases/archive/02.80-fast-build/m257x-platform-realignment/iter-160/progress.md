**Type:** tik — under [`TOK-08`](../decisions.md), which directs working the classes **in descending
measured size**. iter-159 split the spelling-pin class in two; this iter takes the half it exposed.

# iter-160 — the value side, and the fixture iter-155 left behind

## Phase 0d — the pre-flight confirmed the hypothesis before a line was written

This iter's predicate depends on **executing a derivation**, so the derivation was dry-run first:

```
platform_topology.default_profile(stack-demo/platform)  → 'core'
platform_topology.default_services(stack-demo/platform) → ['postgresql','redis','sentinel','backend','gotenberg']
```

The frozen literal in iter-155's confirmed instance is `"postgresql redis sentinel backend gotenberg"`.
**Byte for byte.**

## Phase A/B — the predicate and its proof

**A test literal is a candidate iff its token set EQUALS a value some non-test module derives.**

The asymmetry with iter-159's haystack clause is the design, not an inconvenience: a haystack clause is
decided by *reading*; this one can only be decided by **running the derivation and comparing**. That is
what makes a hit actionable — the repair is never a rewording, it is `import the derivation`.

The proof ran three predictions, all **declared in this iter's `overview.md` before it was run**:

| labeled case | predicted | measured |
|---|---|---|
| iter-155 scope-fence fixture (b1) | **FIRE** | ✓ FIRED at L223, 259, 271, 284, 294, 295, 354, 355 |
| iter-157 `assertEqual(on_disk, registry)` (b2) | **BLIND** | ✓ silent |
| iter-158 `stderr="Traceback (most recent call)\n"` | **BLIND** | ✓ silent |
| a missing platform clone | **exit 2** | ✓ exit 2 |

### ⚠ The finding: iter-155's frozen fixture is LIVE AT HEAD, at eight sites

This is not a historical instance recovered from git. **iter-155 re-pointed that test's *expectation* to
a real derivation and left the *fixture* it feeds the subject frozen.** The repair fixed the side that
had gone RED and left the side that had not. So the class survived the iter that was repairing it — in
the same file — and no reading in five subsequent iters saw it, because nothing was enumerating values.

The second declared blind is worth reading as a limitation, not a footnote: **an INEXACT copy cannot
equal a derived value by construction.** Exactness is what this predicate keys on, so near-misses — the
*more* dangerous defect, since they look right — are invisible to it. iter-158's traceback fixture is one.

## Phase C — the census

```
frozen-expectation-census: 10 unexempted candidate(s) (0 declared exempt) over 9370
multi-token literals in 107 test files, against 3 executed derivation(s)
    platform default_services (5 tokens)            ← platform_topology.default_services
    platform default_services + profile (6 tokens)  ← + default_profile (core)
    fence registry (25 tokens)                      ← repair_postcondition.discover_fences()
```

| file | hits |
|---|---|
| `stack-core/tests/test_bringup_verify_scope_m257x.py` | **8** |
| `stack-core/tests/test_platform_predicate_guard.py` | 1 |
| `stack-injection/tests/test_platform_topology.py` | 1 |

**10 over a 9,370-literal denominator** — two orders of magnitude tighter than the haystack census's 961,
and every hit names the derivation it duplicates and that derivation's provenance. At least one (the
topology module's own golden) is expected to be a **declared exemption** rather than a repair; grading
them is the sweep and is pre-declared out of scope. The `fence registry` derivation (25 tokens) matches
iter-157's corrected count and currently has no frozen copy anywhere — an honest zero **from a derivation
proven to execute**, which is the only kind this milestone accepts.

## Phase D — two defects in this iter's own work, both caught by controls

**1. A silent swallow, inside the census built to find silent defects.** The fence-registry derivation
was first written as `except Exception: pass` around a **wrong-arity** call. The derivation vanished, and
the census reported its population against **2** derivations instead of 3 with no indication one was
missing. Now fatal, like the topology entry, and regression-tested by monkeypatching `discover_fences`
to raise.

**2. ⚠ The mutation control was itself pinned to a COUNT — `§5` rule 71, in the fence of the iter
censusing that class.** It asserted that lowering `MIN_TOKENS` to 1 changes the live candidate total. It
does not: every registered derivation has ≥ 2 members and set **equality** already excludes a one-token
literal, so against today's registry the floor is inert and the count is identical — **a correct guard
reading RED.** Re-pointed to the floor's actual property: *if a one-token derivation is ever registered
(`default_profile` alone is `{'core'}`), a bare word must not match it* — asserted in both directions, so
it is green for the guard and RED without it.

**This is the third time a rule-71 defect has been written by someone inside the thread that owns rule
71** (iter-154's author broke it at iter-155; now this). The rule's own prescribed answer is structural,
not vigilance, and it applied here exactly as written.

**Fence: `tests/test_frozen_expectation_census_m257x.py`, 18 tests, all green.** Mutation control:
`MIN_TOKENS = 0` → **2 failures / 18**.

## Gates

- `tests/test_frozen_expectation_census_m257x.py` — **18 passed, 0 failed**; mutation RED 2/18.
- `repair_postcondition.py` — **OK**, registry **unchanged at 5 + 20 = 25** (this iter declares no
  `FENCE_KIND`, per `D-M257x-159-4`).
- `test_fence_registry_completeness_m257x` + `test_spelling_pin_census_m257x` — **31 passed, 0 failed**
  (iter-159's fence re-run and still green alongside the new modules).

**NOT re-run, named in full (`§5` rule 60 + its rule-71 corollary):** the stack-core suite in full (this
iter **added three new files and modified none**), and the suites of **demo-stack, dev-stack,
stack-verify, stack-injection, stack-seeding, stack-snapshot, stack-secrets, alignment, playthroughs and
clerkenstein**. Note `stack-injection` carries one enumerated candidate and its suite was **not** run —
nothing was modified there, but the sweep that touches it must run it.

## Close — 2026-08-08

**Outcome:** the unfenced half of the spelling-pin class now has an instrument that **executes the
derivation and compares** — **10 candidates over a 9,370-literal denominator, from 3 proven-executable
derivations** — and it immediately found that **iter-155's frozen fixture is live at HEAD in 8 places**:
that iter repaired the expectation and left the fixture, so the class survived the repair inside the file
being repaired. All three blind/fire predictions were declared in advance and all three held.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (`§9`); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
n — Outcome: continue
**Decisions:** `D-M257x-160-1` … `D-M257x-160-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none — both defects found this iter were in this iter's own deliverable and are
repaired in it.
**Routes carried forward:**
- `FIX-M257x-iter159-value-side-subsignature-is-unfenced` — **CLOSED for sub-signature (b1)**; (b2)
  remains, below.
- `FIX-M257x-iter160-iter155-fixture-is-frozen-at-8-sites` — **NEW and concrete.** Re-point the fixture
  in `test_bringup_verify_scope_m257x.py` to `platform_topology.default_services()`. Small, but it must
  be done with the stack-injection + stack-core suites, not blind.
- `SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality` — **NEW.** The predicate keys on
  exactness, so a near-miss fixture (iter-158's) cannot be seen. Near-misses are the more dangerous half.
- `FIX-M257x-iter160-b2-over-strict-direction-still-unfenced` — **NEW.** iter-157's sub-signature
  involves no literal and no haystack; neither census can reach it.
- `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` — unchanged, and now with a sibling: the
  10 here should be graded first, being two orders of magnitude smaller and individually attributed.
- Unchanged and still queued: `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` ·
  `SURVEY-M257x-iter156-other-reporting-layers` · `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`
**Lessons:** **a repair fixes the side that went RED, and the side that did not go RED stays broken.**
iter-155 re-pointed an expectation and left the fixture feeding the same subject frozen — the defect
survived, in the file being repaired, for five iters. When a test has an *input* and an *expectation*,
both are values, and only one of them announces itself. And: **declare the blind spots as predictions
before running the instrument** — three were declared here and all three held, which is worth more than
a recall number, because a prediction can fail and a description cannot.
