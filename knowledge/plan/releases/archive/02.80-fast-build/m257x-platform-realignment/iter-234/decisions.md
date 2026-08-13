# iter-234 — decisions

## `D-M257x-234-1` — the co-quotation containment census is REFUSED, not tuned

The instrument was corrected three times (metadata literals → pointer literals → nearest-anchor
association) and each correction was correct on its own terms. A fourth correction was available
(sentence-scoped polarity detection, paraphrase-tolerant matching) and is **declined**.

**Rationale.** The 5/5 hand-verification showed the residual failures are not a long tail of edge cases —
they are the corpus's *normal, correct* citation practice. Tuning the matcher until the number looked good
would be fitting a rule to a sample (`§5` Trap A) and would produce an instrument whose green means
"my regex agrees with itself." The honest output is the refusal, written into `platform-alignment.md` § 5
with its taxonomy, so the class is closed rather than re-attempted.

**What was kept instead:** the *two-clock* reading (line-existence + stated-sha resolution), which needs no
text matching at all and produced the iter's one durable positive — 15 correct-and-dated sites. Routed as
`ROUTE-M257x-234-two-clocks-is-fenceable`.

## `D-M257x-234-2` — zero corpus prose repaired, deliberately

The census surfaced 24 MISMATCH candidates on its valid subset. **None was repaired.** Every candidate
inspected was the corpus being correct, so every "repair" would have introduced a defect into a correct
sentence — the induction hazard `TOK-08` and `§5` rule 107 already name, here with a measured 100 %
false-positive rate on the inspected set.

The 19 uninspected candidates are **not** asserted clean. They are asserted **ungraded**, by an instrument
now on record as unable to grade them.

## `D-M257x-234-3` — the clone set was read, never written

No `git fetch`, `reset`, `checkout` or `clean` ran in any clone. `stack-demo`'s set stays exactly as
iter-233 censused it (`app` 28 / `next-web-app` 12 / `ant-academy` 9 behind origin/main), so this iter's
readings are comparable with iters 230–233 and with each other.
