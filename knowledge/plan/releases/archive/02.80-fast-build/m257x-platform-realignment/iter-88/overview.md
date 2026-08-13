---
iteration_type: tik
status: archived
opened: 2026-08-05
---

# iter-88 — a fence's own test suite is not fenced

**Type:** tik, under [`TOK-05`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).

## Step 0 — re-survey

Gate **4 of 5** at open. `rosetta 904502c`, `rext 717d565`, both clean, both local == remote.
Platform at origin HEAD `0c91421`, re-fetched at this open — **unchanged**, so the re-scope trigger
stays where iter-87 graded it (occurrence 3, not firing).

## Why this target

iter-87 closed on a finding it did not go looking for. Repairing the corpus was the planned work; the
**rext test suite** turned out to hold a `>= 5` repo-count floor that the platform walked straight
through — `repos.yml` went to 4 and an *anti-vacuity* assert failed for the one reason it must never be
sensitive to. Its lesson was one sentence: **a fence's test suite is a place hand-maintained platform
constants hide, because the suite is not itself fenced.**

And iter-87 only ever ran **`stack-core`**. Four other rext sections read real platform artifacts and
were never re-run at the advanced ref. That is §5 rule 8's territory exactly — iter-04 found a test that
had been RED since iter-02 *because only the newly-written tests were run*.

**Cluster / target:** run the un-run sections at origin HEAD, and treat every skip as a hole to be named
rather than a pass to be counted.

## Hypothesis

The `>= 5` floor is an instance, not a one-off. Other tooling keys on platform service **names** that the
folds have been deleting for three releases, and where it does, the check will have gone quiet rather
than gone red.

## Expected lift

No gate movement — clause 5 is met only by a reading that returns zero, and no reading is taken here.
The deliverable is defects found and the class fenced.

## Phase plan

1. Run `stack-injection`, `stack-verify`, `dev-stack`, `demo-stack` at platform `0c91421`.
2. **Name every skip** (§5 rule 8) — a skip is a hole in the evidence, not a pass.
3. Repair by derivation, never by re-pinning a literal (§2).
4. Generalise into the protocol doc.

## Escalation conditions

- A live defect in a bring-up path → repair it in this iter; do not route a live defect forward.
- A third-party/host-dependent failure → name it, do not chase it.

## Acceptable close-no-lift outcomes

Finding that iter-87's floor was genuinely a one-off would close this iter on a falsification, and would
be worth the run: it would mean the class is bounded and the sweep need not recur.
