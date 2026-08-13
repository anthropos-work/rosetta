---
iter: 83
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: closed-fixed
opened: 2026-08-05
closed: 2026-08-05
---

# iter-83 — why a discharged predicate had a surviving member

**Type:** tik, under [`TOK-05`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).
**Shape:** `tooling` (protocol-codified) — step 1 diagnoses the repair machinery's defect, step 2 ships
the instrument that makes the defect impossible to repeat, and uses it within this iter.

## Step 0 — re-survey (mandatory, §Phase 1)

TOK-05's `Next-tik direction` is exhausted: items 1–3 (the sibling guard, the citation-safety half, the
`mid-fold` state) landed at iters 60–62, item 4 (harden) closed at pass 19/iter-79, and item 5 (*"the next
paired reading"*) was **taken at iter-82**. The re-survey therefore substitutes a fresh target from current
evidence under the same strategy — permitted and required by Phase 1 Step 0.

**Current metric, re-derived at this open:** gate **4 of 5**. Clause 5 is the only open clause and is graded
only by a reading that returns **zero**. iter-82's paired reading returned **29 / 30**. Both repos clean
(`rosetta` `8d6bb6c`, `rext` `4e6b64d` on `main`); all six corpus guards exit 0 at open with the correct
refs supplied; platform ground truth `0dab54d`.

**Cluster / target identified.** iter-82 pre-registered *"if any seat books a finding inside a predicate
iter-81 claimed to repair, the repair was incomplete and that is this iter's headline"* — and it was
falsified: `graphql-wundergraph.md:13` sits inside **P4**, one of the eleven predicates iter-81 reports as
**discharged**, and **both** readings booked it independently.

The site is the small part. **The target of this iter is the repair machinery**, because a repair that
reports a predicate discharged while a member survives is this milestone's signature defect — *a check that
reports a state without measuring it* — appearing in the one place it had not yet been found.

## Hypothesis

Four mechanisms were nominated. They are discriminated by **measurement, not by argument**:

| # | candidate mechanism | discriminating measurement |
|---|---|---|
| H1 | the site was outside the assigned seat's file set (partition gap) | does the repair diff touch the FILE? |
| H2 | P4's membership was **estimated** (`~10`) rather than enumerated | do the exact-count predicates (P3 5, P6 3, P8 9, P9 3, P10 4) also have survivors? |
| H3 | a seat marked the predicate done without re-checking its members | per-anchor coverage < 100 % with the file open |
| H4 | the discharge criterion was *"my sites are fixed"*, not *"no member survives"* | per-anchor coverage < 100 % **and** REJECTED anchors preferred over UPHELD ones |

**Predicted (pre-registered, therefore refutable):** H1 and H2 are false; the true mechanism is H3/H4 — the
repair had **no per-anchor post-condition** and its effective unit was *the file swept for a predicate*, not
*the predicate's membership*.

## Expected lift

**None on the gate metric, and that is planned.** This iter deliberately does not repair corpus claims —
the adjudication of iter-82's union is binding-before-repair (iter-80) and is iter-84's job. The deliverable
is the diagnosis plus the instrument that makes the next repair gradeable. `closed-fixed` is earned by the
planned deliverables landing, not by a metric move (Phase 4 Step 0: *"planned scope = what the overview
committed to"*).

## Phase plan — four planned lines (declared, so the tripwire counts against THIS shape)

1. **Settle the overlap arithmetic** from `iter-82/raw/` before any figure derived from it is used.
   `29 + 30 − 41 = 18 ≠ 15`; the recall estimate depends on the overlap, and the recall estimate is what
   says whether a future zero means anything.
2. **Diagnose the mechanism** — per-anchor coverage of all 152 pre-repair booked anchors against the
   iter-81 diff's old-side hunk ranges. Grade H1–H4. State explicitly whether the other ten predicates
   share it.
3. **Ship the post-condition as an instrument** in `rosetta-extensions/stack-core` — the fence that makes
   an unmeasured discharge impossible. Tests + mutants collected before running (§8 rule 5).
4. **Reconstruct iter-81's record** (`FIX-M257x-iter82-iter81-has-no-record`) from `git log` + the diff,
   **every field labelled reconstructed**, anything the diff cannot establish marked unrecoverable.

## Escalation conditions

- If the coverage measurement shows ≥ 95 % per-anchor coverage, H3/H4 are refuted and the mechanism is
  something else — re-open the diagnosis rather than shipping an instrument against the wrong defect.
- If reconstructing iter-81's record requires any field the diff and the log cannot establish, that field is
  marked **unrecoverable** and left empty. Inventing it would be authoring history.
- The storage carve-out (`DEF-M257x-iter80-storage-prod-bucket`) stays **held**; `storage.md:55,:154,:181`
  unchanged.

## Acceptable close-no-lift outcomes

If the coverage measurement refutes H3/H4 and no instrument is warranted, the iter closes `closed-no-lift`
with the falsification recorded — a complete cycle ending in characterization.

## Out of scope, routed

- **Adjudicating iter-82's union** → iter-84 (`FIX-M257x-iter82-reread-union`).
- **Re-deriving the eleven discharge verdicts as membership questions** → iter-84.
- **Repairing by predicate** → iter-85, gated on the adjudication.
- **Re-running the frozen instrument.** Explicitly not done in this run.
