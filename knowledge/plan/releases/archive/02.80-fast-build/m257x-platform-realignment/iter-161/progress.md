**Type:** tik — under [`TOK-08`](../decisions.md), third clause: *run it to zero, and keep it green.*

# iter-161 — the census reaches zero, and three things resisted on the way

## Phase A — all 10 graded at source, and the overview's own reading was refuted

The `overview.md` recorded a first read of each candidate and said explicitly that those reads were
**hypotheses to be confirmed at source** (iter-158's rule). One of them was wrong, and the confirmation
is what caught it.

It read `test_platform_predicate_guard.py:193` as *"`assertEqual(c.select("core"), {…})` against the
**live** platform"* — which would make it a textbook frozen copy. At source, `self.platform =
write_platform(self.root)`: **a synthetic compose the test writes into a tmpdir.** The literal is the
expected output of a *controlled* input. That is a **golden**, and deriving it would make a parser test
read its answer from the very module it exists to cross-check.

| verdict | count | sites |
|---|---|---|
| **repair — derive it** | **8** | `test_bringup_verify_scope_m257x.py` (6 string args + 2 `write_override` lists) |
| **exempt — golden over a synthetic input** | **2** | `test_platform_predicate_guard.py`, `stack-injection/tests/test_platform_topology.py` |

### The precision limit this measures, rather than assumes

**2 of 10 (20 %) of the census's candidates are goldens over synthetic inputs.** The instrument compares
**values**; it does not know the **provenance of the input** that produced them. A synthetic fixture that
deliberately mirrors the real platform — which is exactly what a good fixture does — will match. That is
a stated precision cost, not a defect to hide, and it is the reason the exemption mechanism exists rather
than a sharper clause: distinguishing the two requires knowing whether a path is a tmpdir the test wrote
or a clone it read, and that is a different instrument.

## Phase B — the repairs

The eight sites now derive from `platform_topology.default_services(PLATFORM)`, bound once at module
level as `DEFAULT_SERVICES` / `PLAT`. The file already required a platform clone (`SKIP_NO_PLATFORM`), so
the repair adds no new precondition.

Both exemptions are declared **at the site, with reasons**, per `D-M257x-159-4`'s declared-never-inferred
rule.

### The marker window bit immediately, and it is right that it is tight

The first exemption did not take: the marker was written at the top of a five-line comment and
`exemption_for` looks back **three** lines. The fix is not to widen the window — **a wide window lets an
exemption drift onto a later, different assertion and silently excuse it.** The convention is therefore
*prose above, marker adjacent to the assertion*, and it is now demonstrated in both exempted files.

## Phase C — zero

```
frozen-expectation-census: 0 unexempted candidate(s) (2 declared exempt) over 9403
multi-token literals in 108 test files, against 3 executed derivation(s)
```

## ⚠ Phase D — the repair broke the proof, and the tempting fixes were both wrong

Running the labeled proof after the repair returned **`MISMATCH — a prediction failed`**. Nothing was
wrong with the instrument: **iter-160's labeled set read the working tree**, because iter-155's instance
was live there. iter-161 repaired it, so the evidence that the census *can fire* disappeared with it.

The two obvious fixes were to **delete the label** or **flip its expectation to BLIND**. Either would
have left a census that now reports **zero** with **no surviving demonstration that it fires at all** —
`§9`'s failure mode arriving through the back door of a *successful repair*, which is the one direction
nobody is watching.

**A labeled set that reads the working tree decays the moment you use it.** Every instance now names the
commit that carries it and the proof reads that blob:

```
✓ iter-155 scope-fence fixture     @4adc595  FIRED at L223,259,271,284,294,295,354,355
✓ iter-157 repair-postcondition    @HEAD     silent
✓ iter-158 traceback fixture       @HEAD     silent
✓ anti-vacuity: a missing platform clone exits 2
```

That is also the ratchet's RED-proof: **8 unexempted at `4adc595`, 0 at HEAD**, same instrument.

**Two fence tests added:** the ratchet (`unexempt == []`, with the offending rows and the repair
instruction in the failure message) and an exemption-quality test — *a declaration with no reason is not
a declaration* — which also fails if the mechanism falls out of use entirely and stops being proven.

## Gates

- `test_frozen_expectation_census_m257x` + `test_spelling_pin_census_m257x` — **39 passed, 0 failed**.
- `stack-core/tests/test_bringup_verify_scope_m257x.py` + `test_platform_predicate_guard.py` (both
  modified) — **194 passed, 0 failed**.
- **`stack-injection` — all 7 test files pass**, including `test_platform_topology`. iter-160 named this
  section as carrying a candidate whose suite it had not run; this iter ran it.

**NOT re-run, named in full (`§5` rule 60):** the stack-core suite in full (~20–35 min — this iter
modified two of its test files and both were run directly), and **demo-stack, dev-stack, stack-verify,
stack-seeding, stack-snapshot, stack-secrets, alignment, playthroughs, clerkenstein** — untouched by this
iter.

## Close — 2026-08-08

**Outcome:** the frozen-expectation class is **run to zero and ratcheted** — 10 graded at source, 8
repaired by derivation, 2 declared exempt with reasons, `0 unexempted over 9,403 literals`. Three things
resisted and each produced a rule: the overview's own first read of a site was **refuted at source**; the
exemption-marker window is **deliberately tight**; and the repair **broke the labeled proof**, whose two
obvious fixes would each have left a zero-returning census with nothing showing it can fire.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (`§9`); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
**y** — Outcome: **exit-7**
**Decisions:** `D-M257x-161-1` … `D-M257x-161-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter160-iter155-fixture-is-frozen-at-8-sites` — **CLOSED by this iter.**
- `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` — **NEW.** 2 of 10 candidates (20 %)
  were goldens over synthetic inputs. Separating them mechanically needs the instrument to know whether
  the input was a tmpdir the test wrote or a clone it read.
- `FIX-M257x-iter161-derivation-registry-is-three-entries` — **NEW.** The census is only as wide as its
  registry: 3 derivations today, so a frozen copy of any *other* derived value is invisible. Note
  `beyond_floor("core") == {"backend","gotenberg"}` sits one line below a candidate and was **not**
  flagged, for exactly this reason.
- `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` — unchanged, and now the larger sibling.
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
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`
**Lessons:** **a successful repair can destroy the evidence that the instrument works, and nobody is
watching that direction.** Every guard in this milestone has been checked for "can it fire"; none had
been checked for "will it still be able to demonstrate that after the tree is clean." Pin the labeled set
to commits. And, smaller but immediate: **write the exemption marker adjacent to what it exempts** — a
window wide enough to be convenient is wide enough to excuse the wrong line.
