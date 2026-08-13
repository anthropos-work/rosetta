# iter-124 — decisions

## `D-M257x-124-1` — the triage's default rule is the generous one, and its error rate is MEASURED before the split is published

`triage.py` assigns `cite` when a sentence names no artifact at all (`R4`, 100 of 344 C1 members). That
is a **presumption**, and this milestone has spent a hundred iterations learning what an unaudited
presumption costs. So `R4` was declared generous in the sealed pre-registration, and then **audited**:
30 seeded verdicts, hand-classified, published in [`audit.md`](audit.md).

**Result: `R3` 21/21 = 100 %, `R4` 6/9 = 66.7 %.** The error is not spread across the rules; it is
entirely inside the default, and it runs one way — `cite` where the hand says `drop` or `hedge`.

**Decision: publish the printed split AND the corrected estimate side by side, never one folded into the
other.** The printed split (`cite` 96.2 %) is what the rules say; the corrected estimate (`cite` ≈ 86.6 %)
is what the audit says they mean. Folding the correction in would hide that **when a sentence names an
artifact, "the evidence exists" held without exception** — which is the actually useful finding, and it
is invisible in a single blended number.

## `D-M257x-124-2` — `fix` is a FLOOR and is never quoted as a rate

The triage decides three fates syntactically. **It cannot decide the fourth**: falsity is not a property
of a sentence's shape. `fix` is therefore a hand-adjudicated input (`FIX_SITES`), found by reading in
consequence order — and this milestone measured its own reading's test–retest recall at **~35 %**
(iter-119).

**Decision: report `4 of 344` as a count found, alongside its recall-corrected estimate (≈ 11, ≈ 3.2 %),
and refuse to call either a rate.** The pre-registered `fix ≥ 15 %` branch **did not fire**, and the
honest form of that sentence is *"did not fire"*, not *"is safe"*. This is `D-M257x-122-3` — *a hunted
sample's error rate is not a population rate* — applied to this iter's own number before anybody else
has to apply it for us.

## `D-M257x-124-3` — a correction that reaches one cell is not a correction, and the fix is to enumerate the predicate

iter-123 measured that `graphql-wundergraph`'s production module is **destroyed**, and wrote it into
`graphql-wundergraph.md`'s fold cell and `org-repos.md` § 3. **24 sites across 13 other files kept
asserting the opposite** — including the fenced `platform-migration-status.md` row whose three siblings
(`cms`, `roadrunner`, `messenger`) iter-123 *did* repair in the same pass.

**Why inspection could not have caught it:** every one of the 24 is locally plausible. *"the Cosmo Router
— prod only"* reads as a correct hedge, not as a falsehood, unless you already know the module was
destroyed. The class is only visible when the **predicate** is enumerated across the corpus — `TOK-02`'s
lesson and `TOK-07`'s unit of repair, both vindicated here on a class neither was pointed at.

**Decision: when a measurement retracts a claim, the repair's unit is the PREDICATE, corpus-wide, in the
same iter that measures it — and the enumeration is recorded so the next reader can check the reach.**
The four repos iter-123 settled share one predicate (*a service repo's own `service_desired_count` is
evidence of production state*); three rows got the correction and the fourth did not, which is exactly
the shape a per-site repair produces. → routed as `FIX-M257x-iter124-desired-count-predicate-reach`.

## `D-M257x-124-4` — three fences fired on this iter's own edits, and one of them REFUSED THE COMMIT

Recorded because the milestone's standing question is whether the fences catch the author, not just the
corpus:

| fence | what it caught | would inspection have? |
|---|---|---|
| `claim_census_guard` (ratchet) | `service_taxonomy.md` 68 → 69 — a repair split one uncited sentence into two | no |
| `unreadable_repo_claim_guard` | a new `module.*_euwest1` mention with no ref pin | no |
| **`repair_postcondition`** | **four citations to `architecture_overview.md:335` became blank-line anchors** because this iter's own edits moved the target 8 lines down — **the pre-commit hook rejected the commit** | no |

**All three were repaired by adding the evidence. None was silenced, waived, or baselined away.** The
third is the sharpest: a repair that fixes 24 false sentences and silently breaks 4 true citations is a
net loss nobody would have noticed, and the only reason it was noticed is that the fence runs at commit
time rather than on request.

## `D-M257x-124-5` — the CLAUDE.md `cms` correction was a SIDE-DISCOVERY and it was co-committed, which is a process miss, disclosed

`CLAUDE.md` asserted in two places that M810 *"has **not moved for cms**"* and that its prod state is
*"NOT MEASURABLE from our clone set"*, citing `cms/terraform/main.tf:39`. **iter-123 refuted both.** The
correction had not reached the one file every agent in this repository loads into context.

That is the same predicate as `D-M257x-124-3` but **outside the census SCOPE** (`corpus/services/*.md` +
`corpus/architecture/*.md`), so it is a side-discovery, not planned scope. The scope-creep tripwire says
such a fix lands as a **separate commit** with its own entry.

**It did not — it was staged together with the anchor re-point in `434caa8`.** Recorded here rather than
tidied away: the fix is correct and evidenced, the *process* was not followed, and a decision record that
only lists the times the discipline held is not a record. **It does not change this iter's close status**,
which grades planned scope only.

## `D-M257x-124-6` — 9 sites still hedge about a repo that is now readable, and that is the DIRECTIVE'S forbidden class pointing the other way

`unreadable_repo_claim_guard`'s own closing note, after this iter's edits:

> *22 `module.*_euwest1` mention(s) are satisfied: 9 by an unmeasurable marker, 12 by a REF-PINNED
> reading of `infrastructure`. … a corpus that hedges and measures the same boundary in different files
> is one edit from disagreeing with itself. Reconcile them.*

The run-80 directive forbids **manufacturing hedges for facts somebody can go and measure**. These 9 are
that class arriving from the other direction: hedges that were **correct when written** and became
manufactured the moment iter-123 cloned `infrastructure`. **A hedge has an expiry, and nothing in this
corpus watches it** — the guard reports the split but does not grade the 9 as stale.

**Decision: do NOT bulk-convert them in this iter.** Re-deriving 9 sites needs the `infrastructure` clone,
which is not on disk, and a fourth line of investigation would fire the tripwire. Routed as
`FIX-M257x-iter124-stale-hedges-on-infrastructure` with the guard's own count as the denominator — the
handler must re-clone at a named sha and either measure or re-justify each.
