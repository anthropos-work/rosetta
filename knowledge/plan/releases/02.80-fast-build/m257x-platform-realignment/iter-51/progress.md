**Type:** tok (triggered — 3-no-prog streak, iters 48 · 49 · 50)

# iter-51 — the strategy revision

Step 0 re-surveyed and the trigger is **not stale**: platform origin `2adcf714` unchanged, the corpus
unrepaired, the union of readings #9 and #10 standing at 18 anchored blockers with `N̂` ≈ 23.

**Output:** [`TOK-03: repair the UNION, shrink the estimator, make the edits smaller`](../decisions.md#tok-03-repair-the-union-shrink-the-estimator-make-the-edits-smaller--2026-08-03).

## The revision in plain terms

Two strategies have been spent making the **reading** better. iter-50 measured that the reading was never
the binding constraint: a single 7-seat pass finds ~43% of what is there, so **a repair pass can only ever
repair 43% of the pool**, and repair-then-read settles at a fixed point instead of converging. That fixed
point is where the last five readings have been.

So TOK-03 stops optimising the reading and changes four things about the work:

1. **Repair the union of two blind readings, never one** — 18 findings, 78% coverage, instead of 14 at 61%.
2. **Drive the residual estimate `N̂` down first, and take clause 5's reading when it is small.** At
   `N̂` ≈ 23 a zero reading has probability ≈ 10⁻⁵; at `N̂` ≈ 2 it is reachable. **Clause 5 is unchanged** —
   this is about what happens *before* the reading, never about what it must return.
3. **Make the repairs smaller** — prefer deleting a false claim to rewriting it, and count the words each
   repair adds. The two induced classes no fence can reach (paraphrase leak, overshoot in new text) are
   both properties of *rewriting*; the only lever on an unreachable class is to shrink the surface it can
   live on.
4. **Put two blind adversarial readers on the repair diff before the commit**, not after. Seat G has been
   the highest-yield seat in both readings and it has always run one pass too late.

## Close — 2026-08-03

**Outcome:** strategy revised. `TOK-03` authored from a measured recall rather than a conjectured one, and
it names a different optimisation target than either prior strategy: **coverage and repair surface**, not
instrument sharpness.
**Type:** tok (triggered)
**Status:** closed-fixed — the tok's planned deliverable is the revised strategy, and it landed with its
trigger verified non-stale, its evidence measured, and its next-tik direction pre-registered so it can be
refuted.
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: **y** — (3) re-scope: n (platform origin
`2adcf714` unchanged at open and close; trigger stays at occurrence 1 of 2) — (4) user-blocker: n —
(5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: **exit-2**
**Decisions:** `TOK-03` in the milestone-root [`decisions.md`](../decisions.md).
**Side-deliverables:** none.
**Routes carried forward:** all of iter-50's, now folded into TOK-03's ordering —
`FIX-M257x-iter50-union-set` (18, next tik) · `FENCE-M257x-iter50-consecutive-audit-mode` ·
`CHECK-M257x-iter50-audited-zero-is-evidence` · `FENCE-M257x-iter50-paraphrase-leak` (demoted: the class
is live but is not the binding constraint) · `CHECK-M257x-iter49-overshoot-has-no-instrument` (now
addressed by TOK-03 move 4, which answers it with a *reader* exactly as iter-49 said it would have to be)
· rosetta's root `CLAUDE.md`, stale on two claims and **outside the 40-file partition**.

**Lessons:**
1. **Two consecutive strategies optimised the same term without either checking that it was the binding
   one.** The check was cheap — one extra reading — and the protocol had already prescribed it in writing.
   **When a strategy stalls, measure which term it is optimising before authoring the next one.**
2. **A tok authored on a measurement is a different object from a tok authored on a pattern.** TOK-02
   read nine passes of history and inferred a mechanism; TOK-03 has a recall number, a residual estimate
   with a stated bias direction, and two refutable pre-registrations. The first was a good guess. Only the
   second can be wrong in a way that teaches something.
