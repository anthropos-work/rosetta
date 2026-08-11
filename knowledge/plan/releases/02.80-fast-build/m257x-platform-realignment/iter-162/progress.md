**Type:** tik — under [`TOK-08`](../decisions.md), the reach half of *build a fence that enumerates
every instance*.

# iter-162 — the census's own registry was a hand-maintained tuple

## Phase A — the population, measured before anything was changed

`frozen_expectation_census.py` claims its subject is *"a value **some non-test module derives**"*.
Its implementation was **three hand-written `out.append(...)` calls**. Nothing measured the gap, so
the first act of this iter was to measure it — by AST, over every non-test rext module.

| set | count |
|---|---|
| public functions returning a collection of `str`, in non-test rext modules | **125** in 36 modules |
| …*executable here* (0 required args, or path-like required args only) | **53** |
| …in the registry | **2 sites → 3 executed derivations** |

**Reach was 3 of 53 = 5.7 %** against a docstring that says *every*. This is `§2`'s hand-maintained
tuple — the milestone's own defining defect class — living inside the instrument built to find it,
and `§9`'s iter-159 rule (*grade the instrument at the grain of its claim*) pointed at iter-160's
instrument rather than iter-159's.

## Phase B — the repair is a DECLARED registry whose COMPLETENESS is derived

The registry could not simply become *"execute everything"*: some of those 53 shell out to docker,
walk the whole git tree, or return an audit **verdict** rather than a reference set. So
`stack-core/derivation_registry.py` classifies **every** executable-here site as `REGISTERED` or
`DECLINE:<class>: <reason>` — explicit per site, never inferred (`D-M257x-159-4`) — under seven named
decline classes (`verdict` · `tree-scan` · `self` · `instance-state` · `waiver-file` · `external` ·
`history`), each stating *why the value cannot be a frozen copy worth censusing*.

**What is derived is the completeness, not the contents.** `unclassified()` returns executable-here
sites with no decision, `stale_decisions()` returns decisions for sites that no longer exist, and
both are asserted empty. A derivation added to rext tomorrow turns the fence RED with its own id.

| | before | after |
|---|---|---|
| registered **sites** | 2 | **13** |
| **executed derivations** | 3 | **24** |
| executable-here population | 53 | **59** (+6 this module's own) |
| whole population | 125 | **130** |

The 24 cross-check against the corpus without being told to: floor **3**, legal profile tokens **5**,
`repos.yml` clone set **4**, demo env knobs **31**, demo CLI flags **10**, fence registry **25** —
every one the number `CLAUDE.md` and `demo-up-defaults.md` publish.

### The fence fired twice during its own iter, which is the best evidence it works

1. `frozen_expectation_census.py::build_derivables` **entered the population** the moment it was
   rewritten as a comprehension — caught, and classified `self`.
2. Seven `predicate_enumerator` declines read `"DECLINE:instance-state."` with **no reason after the
   class**. The decline-quality test rejected them. *A declaration with no reason is not a
   declaration* — iter-161's rule, re-earned one layer up.

## Phase C — 4 new candidates, and 4 of 4 are FALSE POSITIVES

Widening 3 → 24 surfaced **4 candidates the narrow registry could not see**, all in
`test_platform_predicate_guard.py`. Every one was graded **at source** before any repair was written
(iter-158: *a routed item's proposed repair is a hypothesis, not a plan*), and none is a frozen copy:

| site | matched | verdict |
|---|---|---|
| `:182` `assertEqual(c.floor, {"postgresql","redis","sentinel"})` | `compose.floor` | **golden over the SYNTHETIC compose** `write_platform()` builds — iter-161's declared precision limit, at a new line |
| `:1485` `"backend,,gotenberg"` | `beyond_floor('backend')` | a **fuzz INPUT** in a `CELLS` list. It equals a derived set only because the census splits on commas |
| `:1639` `{"backend","gotenberg"}` | `beyond_floor('backend')` | a **tokenizer contract** — the expected tokenization of a hand-written input |
| `:1639` `"backend, gotenberg"` | `beyond_floor('backend')` | that input |

**The stated outcome, not a hidden one: the class was already at zero, and the widening bought reach
rather than defects.** That was pre-registered in this iter's `overview.md` as an acceptable result.

Two of the four name a precision limit the instrument did not have a word for before. iter-161 found
that the census cannot see the **provenance of an input** (synthetic vs real). These add the adjacent
one: it cannot see the **role of a literal** — expectation, or argument. `"backend,,gotenberg"` is
fed *to* the code under test, and it matched. **Coincidence lives in the small derivations**: 7 of
the 24 carry ≤ 3 tokens, and both false positives here matched a **2**-token set.

## Phase D — the fences

Six tests in `TheRegistryIsEnumeratedNotListed`: completeness · staleness · decline-quality ·
the reach ratchet (a **number**, 13, because *"more than before"* is unfalsifiable once the before is
forgotten) · a registered derivation that cannot run is **FATAL** (the silent-swallow guard, carried
across when the registry moved out of the census) · and a **planted-derivation anti-vacuity control**
written against the subject (a module on disk), not against `DECISIONS`.

The labeled proof gains a **4th commit-pinned instance** — iter-161's rule applied prospectively,
since this iter exempts all four sites in the working tree:

```
✓ iter-155 scope-fence fixture     @4adc595  FIRED at L223,259,271,284,294,295,354,355
✓ iter-157 repair-postcondition    @HEAD     silent
✓ iter-158 traceback fixture       @HEAD     silent
✓ iter-162 registry widening       @c083819  FIRED at L1485,182,1639,1639 (+2 exempt)
✓ anti-vacuity: a missing platform clone exits 2
```

It also changed how the proof **grades**: on `unexempt`, not on raw candidates. A file that is
entirely declared-exempt would otherwise read as a firing instrument forever. The exempt count still
prints, because it is the evidence the match happened.

## ⚠ Side discovery — a rot detector keyed on a DELIMITER measures luck

Adding one comment line to `test_platform_predicate_guard.py` turned `repair_postcondition` RED:
`corpus/services/backend.md:182` cites `…test_platform_predicate_guard.py:435`, and `:435` had just
become a bare `)`.

**The anchor was already wrong before this iter, by 5 lines, and green.** At `c083819` the cited
subject (`CMS_RPC_ADDR=http://backend.internal.anthropos:8081`) sat at `:440`, while `:435` was
`corpus = write_corpus(self.root, body)` — a plausible-looking line, so nothing fired. A **+1** shift
landed a closing delimiter there and the guard went RED instantly.

So the detector's threshold is doing the work: **a rot onto a delimiter is caught; a rot onto ordinary
code is invisible.** Repaired here (`:435 → :441`, verified against the subject, not the offset);
the general case is evidence for the standing route `FIX-M257x-iter138-anchor-rot-fence`.

## Gates

- `test_frozen_expectation_census_m257x` — **26 passed, 0 failed** (6 net-new).
- `test_platform_predicate_guard` + `test_spelling_pin_census_m257x` + `test_repair_postcondition` —
  **225 + 27 passed, 0 failed** (the last re-run after the anchor repair; RED → green).
- `frozen_expectation_census --labeled-set` — **4/4 labels + anti-vacuity**, verdict as predicted.

**NOT re-run, named in full (`§5` rule 60):** the `stack-core` suite in full (~20–35 min — this iter
modified three of its files and all three were run directly), and **demo-stack, dev-stack,
stack-verify, stack-seeding, stack-snapshot, stack-secrets, alignment, playthroughs, clerkenstein,
stack-injection** — untouched by this iter (`platform_topology` and `exposure_claim_guard` are
*imported* by the new registry, not modified).

## Close — 2026-08-08

**Outcome:** the frozen-expectation census's registry stops being a hand-list. Population enumerated
by AST (**130**, of which **59** executable-here), **every** executable-here site classified
REGISTERED or DECLINE-with-reason, executed derivations **3 → 24**, registered sites **2 → 13**, and
the registry's **completeness** — not its contents — is now fenced in both directions. The widening
surfaced **4** candidates the narrow registry could not see; all **4 graded false positive at
source**, so the class stays at **0 unexempted over 9,409 literals**. The completeness fence fired
twice on this iter's own work before the iter closed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (`§9`); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7)
budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-162-1` … `D-M257x-162-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** the `backend.md:182` anchor repair (`:435 → :441`) — a **pre-existing** 5-line
rot, not caused by this iter, surfaced by it.
**Routes carried forward:**
- `FIX-M257x-iter161-derivation-registry-is-three-entries` — **CLOSED by this iter.**
- `SURVEY-M257x-iter162-a-literal-has-a-ROLE-the-census-cannot-see` — **NEW.** 2 of 4 candidates were
  tokenizer *inputs*, not expectations. Separating them needs AST context (is the literal an argument
  to the call under test, or the expected value?), the sibling of iter-161's input-provenance limit.
- `SURVEY-M257x-iter162-small-derivations-are-coincidence-prone` — **NEW.** 7 of 24 derivations carry
  ≤ 3 tokens and both false positives matched a **2**-token set. Raising `MIN_TOKENS` would delete the
  `beyond_floor` class outright, so the honest options are a per-derivation floor or a declared cost.
- `FIX-M257x-iter138-anchor-rot-fence` — **new evidence, unchanged route.** A 5-line rot sat green for
  as long as `:435` happened to be ordinary code; a 6-line rot went RED the same second. The detector
  is keyed on landing-on-a-closing-delimiter, which measures luck.
- `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` — unchanged, and this iter added a
  4th instance of it (`:182`).
- `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` — unchanged, still the larger sibling.
- `SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality` ·
  `FIX-M257x-iter160-b2-over-strict-direction-still-unfenced` · unchanged.
- Unchanged and still queued: `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` ·
  `SURVEY-M257x-iter156-other-reporting-layers` · `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter134-fence-family-has-no-shared-predicate-layer` ·
  `-iter133-two-fives-need-a-fence` · `-iter131-predicate-sets-not-enumerated`
**Lessons:** **an instrument's REACH is a measurement, and until somebody takes it, a hand-list reads
exactly like a census.** Three entries and twenty-four entries print the same sentence — *"0
unexempted candidates"* — and only the second one earns it. The generalizable form: **fence the
COMPLETENESS of a registry, never its contents.** Contents are a judgement that changes; completeness
is derivable, and it is the half that silently falls behind the tree. Corollary earned twice in this
iter's own commit: a completeness fence catches the derivation you add *in the same commit as the
fence*, which is the only kind a reviewer would never think to look for.
