# iter-39 — decisions

## D-M257x-39-1: read all 40 again; weight by repair history, never narrow — 2026-08-02

Run 20 routed a weighting toward the 11 files iter-38 edited. Applied as **weighting** (those files got
double coverage: once in full by their partition owner, once as a diff by a dedicated adversarial reader)
and explicitly **not** as narrowing. Vindicated: ~12 of 37 blockers were in files iter-38 never opened,
against a predicted 3-6. A pass reading 13 of 40 files would have found ~25 and reported the rest clean.

## D-M257x-39-2: repair by CLAIM, not by FILE — 2026-08-02

**Candidate §5 rule 19.** 5 of the 8 self-inflicted defects were cross-file drift: a claim corrected in the
file its owner held while the identical claim survived in a twin file owned by someone else. Disjoint file
ownership is correct for *reading* — it is what produces independent double-finds (§5 rule 18(b), which
worked again here). It is wrong for *repairing*, because a claim does not respect a file boundary.

**Half-repairing a uniformly-wrong corpus is worse than leaving it alone.** A uniformly-wrong corpus is at
least self-consistent; a half-repaired one teaches the reader that the corpus disagrees with itself, and the
next auditor spends its budget adjudicating rather than measuring. Before editing, grep the whole corpus for
the claim and fix every instance in one pass.

## D-M257x-39-3: refuse the `31 of 135` change; route it instead — 2026-08-02

A repairer found that `architecture_overview.md:288`'s "31 schemas auto-filter by organization" is arguably
an undercount by one — `organization.go` declares its own org-filtering `Policy()` — and **declined to change
the number**, on two grounds: the home derivation lives in `security_compliance.md`, which it did not own, so
flipping one site would manufacture a cross-file contradiction; and the **denominator** is itself ambiguous
(135 by `grep -l 'ent.Schema'`, 112 by counting embedded declarations).

Ratified. This is the fence that has been **wrong four times, in both directions**, and its current
generation was established correct only by testing both conjuncts of the predicate per file (§5 rule 17).
A one-site numeric edit is exactly how the previous three generations went wrong. Routed as
`CHECK-M257x-iter39-tenancy-fence-off-by-one`.

## D-M257x-39-4: do not let this iteration settle the EU-AI-Act classification — 2026-08-02

Standing from iter-38 and re-affirmed in every repairer brief. The adversarial pass confirmed neither
`ai_architecture.md` nor `security_compliance.md` now asserts a legal conclusion; both defer explicitly to
counsel. This was the single most dangerous thing the sweep could have done, and it did not.
