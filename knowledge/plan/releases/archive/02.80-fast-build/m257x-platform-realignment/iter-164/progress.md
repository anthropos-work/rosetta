**Type:** tik — under [`TOK-08`](../decisions.md), and squarely in `§5` rule 71's family.

# iter-164 — two shared helpers, both pinned to a spelling

## Phase A — the content-free clause knows punctuation, not properties

`anchor_construct_guard` reports an anchor that lands on nothing: blank line, bare closing delimiter,
table separator, table header. The delimiter clause is `_CLOSER_ONLY = ^[\s})\];,]*$` — **a set of
characters**, so it recognises a block terminator only in the C family. A shell script closes a block
with a **word**, and an anchor on a bare `fi` carries exactly as little as one on `})` while reading,
to this guard, as ordinary content.

**Measured before the clause was written**, corpus-wide over **684 resolved in-range anchors**:
**1 instance** — `up-injected.sh:2494`. Reported as one, not generalised.

### The clause nearly shipped a token that would have renamed a class

The first draft was `(fi|esac|done|;;|end)`. **`;;` is punctuation**, already matched by
`_CLOSER_ONLY` — so including it would have made the clause look broader while changing nothing,
**except the class name reported for every `;;` anchor**, which four releases of recorded verdicts
spell `anchor-on-closing-delimiter`. The boundary is now pinned by a test in both directions, and
`finish()` / `ending := true` / `fi := 1` / `endpoint = 3` are asserted to stay content.

### The repair, and the tripwire on it

`corpus/ops/verification.md:662` → `up-injected.sh:2714` (the line that passes `--services`).
`:2494` was the `fi` closing an unrelated Directus-URL-rewrite branch.

**The claim around it was NOT adjudicated, deliberately.** Its other anchor
(`stack-verify/lib/services.sh:43-44`) does not obviously carry the `jobsimulation`/`cms` rows the
sentence attributes to it, and whether those services are inside a `--services` scope *derived from
the platform compose* is a third question. That is the 3rd line of investigation; the pointer is
repaired and the claim is routed.

## ⚠ Phase B — iter-163 blamed the wrong component, and that is retracted here

Two iter-163 exemptions read *"a defect of the shared helper."* `_block_bounds`'s own docstring says
it returns **the PROSE block**, and it was extracted verbatim at iter-100 so that one definition of
"block" serves its markdown callers. **It does what it says; `anchor_subject_census` applied a prose
predicate to Go source**, where a blank line inside a function is ordinary. The caller was wrong.

`enclosing_block(tlines, n, target)` now dispatches on suffix: source → the top-level declaration
(back to a column-0 opener, forward to its column-0 terminator); anything else → the prose block,
untouched.

### The finding is what the loose clause was HIDING

| | iter-163 | iter-164 |
|---|---|---|
| declared exemptions | **9** | **5** |
| …absorbed by mechanism | — | **4** |
| …net-new, surfaced by the sharper block | — | **1** |

The prose block for a 2,700-line shell script ran **line 1 → line 154**, so any literal in the first
154 lines counted as "inside the block." `demo-up-defaults.md:77` was sitting inside that, and it is
a real candidate (graded *not-the-subject*: the anchor's subject is `STACK_PUBLIC_HOST`, which IS the
cited line).

**Every guard in this milestone has been audited for *can it fire*, and since iter-161 for *can it
still show that it fires*. Nobody had audited an ACCEPTANCE clause for over-reach** — a too-generous
"this is fine" produces a green with nothing to look at, which is the same failure as a guard that
cannot fire, arriving from the other side.

## Gates

- `test_anchor_subject_census_m257x` — **21 passed, 0 failed** (4 net-new, incl. the control that the
  prose block would NOT have covered the case — without which the fix is unfalsifiable).
- `test_anchor_construct_denominator` + `test_anchor_offset_guard` + `test_repair_postcondition` —
  **85 passed, 0 failed** (4 net-new terminator tests; `repair_postcondition` re-runs the live-tree
  grading over the repaired doc).
- Combined re-run of all four suites — **106 passed, 0 failed**.
- `anchor-subject-census` — **0 unexempted over 137 adjudicable pairs**, 5 graded exempt.

**NOT re-run, named in full (`§5` rule 60):** the `stack-core` suite in full (~20–35 min — this iter
modified `anchor_construct_guard.py`, `anchor_subject_census.py` and two test files, and every one of
their own suites was run directly), and **demo-stack, dev-stack, stack-injection, stack-verify,
stack-seeding, stack-snapshot, stack-secrets, alignment, playthroughs, clerkenstein** — untouched.

## Close — 2026-08-08

**Outcome:** both shared anchor helpers now express the property they meant. The terminator class is
**1 over 684** and repaired; the census's **declared** exemptions fall **9 → 5** because mechanism
replaced four human declarations — and the sharper acceptance clause **surfaced one candidate the
loose one had been hiding**. iter-163's attribution of two exemptions to `_block_bounds` is
**retracted**: the helper was right and the caller applied a prose predicate to source.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (`§9`); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7)
budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-164-1` … `D-M257x-164-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter163-block-bounds-under-reaches-by-two` — **WITHDRAWN.** The helper was not defective;
  the route would have sent a future iter to widen a helper four guards depend on.
- `FIX-M257x-iter163-anchor-guard-does-not-know-shell-keywords` — **CLOSED**, class size 1, fenced.
- `SURVEY-M257x-iter164-acceptance-clauses-are-unaudited-for-over-reach` — **NEW.** This milestone
  audits guards for *can it fire*; the accept side of every fence in the family is unmeasured, and one
  over-wide clause was hiding a real candidate.
- `SURVEY-M257x-iter164-verification-662-claim-not-adjudicated` — **NEW.** Pointer repaired, claim
  routed: `services.sh:43-44` does not obviously carry the rows the sentence attributes to it, and
  whether `jobsimulation`/`cms` are inside a compose-DERIVED `--services` scope is a second question.
- `FIX-M257x-iter163-block-ref-attaches-the-wrong-sha` — unchanged (1 exemption depends on it).
- `SURVEY-M257x-iter163-anchors-with-no-quoted-literal` ·
  `SURVEY-M257x-iter163-generic-literals-are-unadjudicable` — unchanged.
- Unchanged and still queued: `SURVEY-M257x-iter162-a-literal-has-a-ROLE-the-census-cannot-see` ·
  `SURVEY-M257x-iter162-small-derivations-are-coincidence-prone` ·
  `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` ·
  `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` ·
  `SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality` ·
  `FIX-M257x-iter160-b2-over-strict-direction-still-unfenced` ·
  `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` ·
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
**Lessons:** **audit the ACCEPT side, not just the fire side.** Eleven iters of this milestone have
asked *can this guard fire* and, since iter-161, *can it still demonstrate that*. Nobody asked *does
this guard accept too much* — and a too-generous acceptance clause produces a clean green with no
finding to inspect, which is indistinguishable from a working fence. Second: **an exemption that names
the wrong cause is worse than no exemption**, because it routes a defect against an innocent
component; iter-163's two were withdrawn here before anyone widened a helper four guards depend on.
