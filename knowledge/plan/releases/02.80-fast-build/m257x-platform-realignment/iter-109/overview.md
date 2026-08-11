---
iter: 109
milestone: M257x
iteration_type: tik
iter_shape: reading
status: archived
opened: 2026-08-06
---

# iter-109 — `TOK-06` step 4: the read

**Active strategy reference:** [`TOK-06: fence the inflows before repairing again`](../decisions.md#tok-06-fence-the-inflows-before-repairing-again--2026-08-06).
This iter executes its **step 4**, the last of five and the one the other four exist to make meaningful:
*"Read LAST, once the inflows are watched. Read first and the next reading measures the same inflow again
and costs a full cycle to say so."*

**Type selection.** Tik. Phase 0 rule 2's streak was adjudicated at iter-108 (`D-M257x-108-1`, codified in
the protocol doc) and **ratified**: iters 105–107 took no reading at all, so `N` was **UNMEASURED, not
unmoved**, and a delta requires two measurements. The floor is preserved — three tiks that *did* measure and
did not move still fire — and both guard-rails held (each of those iters says in its own close that no `N`
movement is claimed; the sequence was declared in advance by `TOK-06`). Not re-litigated here.

## Step 0 — re-survey before targeting

`TOK-06`'s next-step direction for this slot is *"read LAST"*, and the re-survey confirms it is still the
right target and still unstarted:

- Steps 0–3 are **landed**: iter-105 (guard-tree provenance), iter-106 (drift fence), iter-107 (induction
  checks), iter-108 (the repair — 22 predicates / 23 files, machine reach **46/46 = 100 % of the upheld
  union**).
- The residual has **not** been re-measured since iter-103. The metric is unmeasured by construction.
- **The subject held still**: all 14 platform clones are at the identical sha they were at iter-103's read
  (`ground-truth.md`). This is the cleanest possible conditions for the step-4 question, and those
  conditions are perishable — the next bring-up fetches.

No substitution. The TOK-named target is current.

## Cluster / target identified

Not a cluster — a **measurement**. The whole in-scope corpus (`corpus/services/**` +
`corpus/architecture/**`, 40 files / 10,694 lines) read by **14 blind seats**: 7 seats × 2 independent
readings (#27, #28) of one recomputed partition.

## Hypothesis

**With both inflows fenced and the platform subject provably frozen, the distinct-false-predicate count `P`
falls materially below iter-103's 22.** If it does not, iter-103's 61 %-drift decomposition described the
defect *class* rather than its *arrival*, and repair-and-read is not a converging loop for clause 5.

## Expected lift

**None is claimed, and that is the point.** This is a measuring pass; it takes no repair and moves no
metric by construction. Its deliverable is a **number graded against a rule sealed before the first seat was
dealt** — see [`pre-registration.md`](pre-registration.md), which bands **both** `P` (primary) and `N`
(secondary), plus 12 named bands.

## Phase plan

1. Ground truth re-derived + guard family run **with its own fence tree stated** → [`ground-truth.md`](ground-truth.md).
2. Instrument copied verbatim, sha re-checked **after** copying → [`briefing-AS-DELIVERED.md`](briefing-AS-DELIVERED.md).
3. Pre-registration **sealed in its own commit, before any seat is dealt**.
4. Reading #27 — 7 blind seats. Each committed **verbatim on landing**.
5. Reading #28 — 7 blind seats, identical partition. Same discipline.
6. **Adjudicate before reporting.** Then grade every pre-registered band.

## Escalation conditions

- A clone moving mid-reading → record the fetch, treat the affected refs as MOVED (§5 rule 41a's stated
  limit), do not assert the rule held.
- A guard going RED mid-reading → reproduce it read-only at the same corpus ref from the **authoring** tree
  before drawing any conclusion (`D-M257x-103-7`).
- Anything that would require a repair to settle → routes; it is not fixed inside the measuring pass.

## Acceptable close-no-lift outcomes

**Every branch of the primary rule is an acceptable close.** A `P ≥ 15` is not a failed iter — it is the
finding that changes the milestone's plan, and the pre-registration says it is reported first and loudest.
An iter that reports a number honestly against a rule it could not re-cut has delivered its planned scope.
