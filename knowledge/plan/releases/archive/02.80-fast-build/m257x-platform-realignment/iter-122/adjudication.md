# iter-122 — the CLAIM CENSUS, adjudicated

**Read `pre-registration.md` first.** Everything below is graded against bands and falsifications sealed
at commit `1262ca7`, before a line of the instrument was written. Two of them fired.

**F4 discharged up front: nothing in this file is `P`, is `N`, or is a clause-5 verdict.** Clause 5's
instrument is still the frozen graded read; **the gate did not move and stands at 4 of 5.** No reading was
taken this iter.

---

## 1. The headline: a sealed band FAILED, and it failed on the low side

| sealed band (F5) | measured | verdict |
|---|---|---|
| tier-1 **`DOES-NOT-SUPPORT` rate** on the adjudicated set: **≥ 4 % and ≤ 25 %** | **3 / 427 = 0.70 %** — and **1 / 427 = 0.23 %** after two verdicts were traced to the census's own stale substrate (§4) | **FAILED — low side** |
| tier-2 unevidenced count over the surface: **≥ 200 and ≤ 2,500** | **1,151** | held |
| tier-1 pairs enumerated: **≥ 900 and ≤ 3,000** | **2,603** | held |

The prediction was built from iter-119's *"measured error rate on the judged set is ≥ 13.3 %"* and
iter-120's 8 wrong-construct repairs. **Exhaustive enumeration of the same class does not reproduce it.**

That is not a contradiction, and the explanation is the whole reason the user funded this instrument:
**iter-119's 13.3 % was a rate over claims a READING had selected.** A reading selects by suspicion — it
stops on the sentences that look wrong. The rate among sentences a hunter stopped on is not the rate in
the population, and the gap here is **roughly nineteen-fold**. Every error rate this milestone published
from a graded read is a rate over a hunted sample, and none of them was ever a population rate.

**Stated as plainly as it can be: on the one class where exhaustive verification was achievable, the
corpus's cited claims are far better than four readings implied.** That result is worth more than the
number it replaces, and it could not have been obtained by reading harder.

---

## 2. TIER 1 — what was enumerated, what was adjudicated, what was false

### 2.1 The enumeration (mechanical, complete over the 40 files)

| quantity | value | denominator |
|---|---|---|
| files | **40** | the clause-5 surface, unchanged |
| claiming units (non-heading blocks) | **2,485** | +587 headings, excluded |
| **tier-1 pairs** (claiming unit × citation) | **2,603** | over **932 distinct units** = 37.5 % of 2,485 |
| …carrying a **line pin** | **695** | 26.7 % of 2,603 |
| …of those, **materializable** against the clone set | **525** (first pass) / **484** (after the bare-basename fix) | 75.5 % of 695 |
| …**ambiguous** bare basenames, refused rather than guessed | **209** | 30.1 % of 695 |
| …pin **out of range** in the resolved file | **2** | both sha-qualified to platform `2adcf71` |

Composition of the 2,603 by citation kind: **sha 826 · source-pin 695 · source-file (no pin) 525 ·
doc-link 507 · anchor-link 20 · external-link 19 · corpus-src 11.**

### 2.2 The adjudication — exhaustive over the class it names, and only over that class

**All 525 materializable line-pinned pairs were adjudicated. Not a sample: the split into 12 batches was
for parallelism.** 12 independent adjudicators, blind to one another, briefed by
[`adjudicator-brief.md`](adjudicator-brief.md), verdicts committed **verbatim before aggregation**.
**525 rows returned for 525 items — 12 of 12 seats, 0 lost.**

**41 of the 525 are QUARANTINED** and excluded from every rate below: the census's first draft resolved a
**bare basename** by probing clone directories alphabetically, so `main.go:1276` materialized
`app/main.go` when `sentinel/main.go` satisfies the citation equally. Control `test_22b` caught it after
the batches were already in flight; they are named in
[`raw/QUARANTINE-guessed-resolutions.txt`](raw/QUARANTINE-guessed-resolutions.txt). **Six of the twelve
adjudicators independently flagged the same artifact from the other side**, which is the strongest
evidence available that the quarantine is drawn in the right place and not to flatter a number.

| verdict | all 525 | **clean 484** (quarantine excluded) |
|---|---|---|
| `SUPPORTS` | 436 | **407** |
| `PARTIAL` | 17 | **17** |
| `DOES-NOT-SUPPORT` | 4 | **3** |
| `UNRESOLVABLE` | 68 | **57** |

- **Decidable set = 427** (clean minus `UNRESOLVABLE`).
- **`DOES-NOT-SUPPORT` = 3 / 427 = 0.70 %**; **1 / 427 = 0.23 %** after §4.
- **`DOES-NOT-SUPPORT` + `PARTIAL` = 20 / 427 = 4.7 %** — the broadest defensible reading of "false",
  and the band named the narrow token, so **the band failed on the number it named.**

### 2.3 The dominant non-`SUPPORTS` class is not error — it is UNCHECKABILITY

**57 of 484 (11.8 %) are `UNRESOLVABLE`: the citation is pinned to a commit this tree is not at.** Not
one of them is a claim shown to be wrong; every one is a claim that **cannot be checked from the clone
set as shipped**. Adjudicators repeatedly confirmed the *construct still exists at a different line* —
e.g. `app/main.go`'s handler registrations pinned `@ b948604f` are all live at `ad9f3c49`, ~110 lines
adrift.

This is a **re-pin backlog, not a retraction backlog**, and it is by a factor of nineteen the larger
finding: the corpus's cited claims are mostly right and increasingly **unverifiable from a reader's own
checkout**. A citation nobody can follow is a citation that will rot silently — which is the mechanism
this whole milestone exists to interrupt.

### 2.4 The one genuine `DOES-NOT-SUPPORT`, and it is now repaired

`corpus/services/academy-backend.md:62` cited **`app/main.go:471-472`** for *"constructs the two managers
from `STORAGE_S3_BUCKET` / `STORAGE_S3_PUBLIC_BUCKET`"*. At `ad9f3c49` — and the `app` clone is **level
with its `origin/main`**, so there is no substrate excuse — `:471` is the closing brace of the Bedrock
error branch and `:472` is blank. The construction is at **`:524-525`**, from names read at `:516-517`.
The anchor was **~53 lines short of its own subject**, and `app`'s own `CLAUDE.md` names `:524`/`:525`
independently.

Repaired, with the ref now stated beside the pin. `claim_twin_guard`'s question was asked before the
edit: the claim is stated in exactly one place.

### 2.5 F2 — exhaustive tier-1 adjudication was NOT achievable in run 78, and here is the measurement

Per the sealed rule, this is stated with numbers and **no corpus-correctness percentage is derived from
it**.

- **Adjudicated: 525 pairs of 2,603 = 20.2 %.** Of the line-pinned class: **525 of 695 = 75.5 %**;
  of the *materializable* line-pinned class: **525 of 525 = 100 %.**
- **Cost, measured:** 12 parallel adjudicators, wall time **335–517 s** each (parallel, so ≈ 8.6 min
  end-to-end), **≈ 2.1 M subagent tokens** in total, ≈ 4,000 tokens per pair.
- **Extrapolation to the full 2,603:** ≈ **60 adjudicator runs** and ≈ **10.4 M subagent tokens** — and
  that is the optimistic figure, because the 1,908 unadjudicated pairs are *harder*, not easier: 826 are
  bare shas with no line to read, 525 name a file with no pin (nothing to excerpt), and 507 are doc-links
  whose "construct" is a whole document.
- **Verdict: F2 FIRES for tier 1 as a whole.** It does **not** fire for the class §2.2 names, which was
  covered exhaustively. **The honest form of this result is "100 % of a named 525-pair class", never
  "20 % of the corpus checked".**

---

## 3. TIER 2 — the half no fence in this family had ever looked at

Every prose sentence in the 40 files that asserts something about the platform or the tooling, split by
what it offers a reader as evidence. **Denominator named, per F6.**

| | count | share of 3,292 assertion candidates |
|---|---|---|
| **CITED** — a citation somewhere in its containing block | 2,117 | 64.3 % |
| **UNCITED but HEDGED** — says it is not a measurement | **24** | **0.7 %** |
| **UNCITED and UNHEDGED** — the tier-2 defect | **1,151** | **35.0 %** |

**1,175 factual assertions in the clause-5 surface carry no citation. 1,151 of them — 98.0 % — carry no
hedge either.** The iter-093 principle (*a claim you cannot measure must say so*) is honoured in **24
places out of 1,175 opportunities**.

They live in **975 distinct blocks across 39 of the 40 files** (one file carries zero), enumerated line by
line in [`raw/tier2-unevidenced-assertions.tsv`](raw/tier2-unevidenced-assertions.tsv) with a per-file
breakdown alongside. Heaviest: `external_services.md` 102 · `ai-readiness.md` 76 ·
`alignment_testing.md` 70 · `service_taxonomy.md` 69 · `ant-academy.md` 68.

**Three disclosures that must travel with that number, and they are printed by the guard itself:**

1. **UNEVIDENCED, never FALSE.** Most of these sentences are probably true. What is measured is that a
   reader has nothing to follow.
2. **It is a FLOOR.** Citation and hedge scope is the containing *block*, so one citation anywhere in a
   long paragraph exonerates every sentence in it. The generous direction, chosen deliberately.
3. **The subject token is a PROXY** for "is a factual assertion", and its miss rate is measured in §5
   rather than assumed.

---

## 4. The census's own defects — three, all caught, two by its own controls

**Recorded first and in full, because a census that hides its own error budget is worth less than a
reading that admits one.**

| # | defect | caught by | consequence |
|---|---|---|---|
| 1 | the artifact-name set was derived from the clone set and the **declared** archived/external/rext names were never unioned in — `redis`, `clerk`, `directus`, `skiller`, `stack-seeding` were invisible to the subject-token proxy | **mutation control 11**, before any number was published | tier 2 undercounted by **229 sentences** (922 → 1,151) |
| 2 | a **bare basename** was resolved by probing clone dirs alphabetically — a silent guess | **control `test_22b`**, after the batches were dispatched | **41 of 525** pairs quarantined; independently flagged by 6 of 12 adjudicators |
| 3 | materialization read the clones' **WORKING TREES**, and 6 of 13 are behind their own fetched `origin/main` | **the adjudicators**, not a control | **2 false `DOES-NOT-SUPPORT` verdicts** |

Defect 3 is the one worth reading twice. `storage` is **20** commits behind, `messenger` **7**,
`jobsimulation` **4**, `next-web-app` **4**, `cms` **2**, `rosetta-extensions` **1**. Four independent
adjudicators booked the corpus's M810 claim — *"`6092c6d2` deleted the `module "jobsimulation"` block"* —
as contradicted, because at `462343b0` the block is still there at `:31` with `service_desired_count = 0`
at `:40`.

**The corpus is right.** `6092c6d2` **is an ancestor of `origin/main` `82cb66ec` in the very clone the
census read from**, and at that commit the module is gone and the atlas tracker survives, exactly as the
corpus says. The census read the checkout instead of the ref.

> **A stale substrate does not merely fail to confirm a claim — it manufactures evidence against a true
> one.** That is a new standing rule, and it is the sharpest thing this iter learned. The census cannot
> fix the checkout (the clone set belongs to a live demo stack this milestone may not touch), so it does
> the next honest thing: it **discloses the substrate on every run**, with a staleness table, and
> `KNOWN_WEAKNESS` clause (5) names the failure mode in the guard's own output.

---

## 5. F1 — the blind recall audit. It does not fire.

A 60-line random sample was drawn with **seed 122 and sealed in the pre-registration commit, before the
enumerator existed**. It was adjudicated by an auditor **blind to the census's output**, who confirmed
the blindness held.

| quantity | value |
|---|---|
| sampled prose lines | 60 (of a 9,543-line population) |
| auditor booked **ASSERTION** | **36** · NOT-ASSERTION 24 |
| census placed in **tier 1** (a citation in its block) | 24 |
| census placed in **tier 2** | 10 |
| census placed in **NEITHER** — the misses F1 measures | **2** |
| **enumeration recall** | **34 / 36 = 94.4 %** — floor was **≥ 90 %** |
| miss rate | 5.6 %, against a firing threshold of **> 10 %** |

**F1 does not fire.** The two misses (`alignment_testing.md:447`, `external_services.md:27`) are both
`UNCITED PLAIN` — i.e. both are tier-2 defects the census did not book, which sharpens rather than
softens the FLOOR disclosure: **tier-2 recall against the auditor's own uncited class is 10 of 13 =
76.9 %.** The 1,151 is a floor by a measured margin, not a rhetorical one.

Precision, measured though not pre-registered: **1 of the auditor's 24 NOT-ASSERTION rows** was booked by
the census (`architecture/README.md:3`) — a 4.2 % false-positive rate on that class.

The auditor also recorded three rubric decisions that move the count and are worth carrying: sentence
**continuations** (a two-word line completing a claim above it) are assertions and a line-level enumerator
drops them; **table granularity** (row vs whole table) is the single biggest lever on the cited/uncited
split; and **extensionless code paths** (`alignment/cmd/deployrun`, `apps/web`) are anchors a reader would
follow but the citation rubric does not count.

---

## 6. What the census establishes that a reading cannot — and what it does not

**Establishes:**

1. **A denominator.** 2,603 cited claims, 2,485 claiming units, 3,292 assertion candidates. No reading in
   this milestone ever produced one; every `P` was a numerator with no denominator behind it.
2. **That a hunted sample over-states the population error rate — by about nineteen-fold here.** This is
   a fact about the milestone's own prior measurements, obtainable only by censusing what a reading
   sampled.
3. **Exhaustive coverage of a named class**: 525 of 525 materializable line-pinned pairs, every one
   individually addressed, with the 41 it could not resolve honestly **named** rather than guessed.
4. **An entire defect class that no fence could see**: 1,151 unevidenced assertions, and the fact that
   the corpus's hedge discipline runs at **24 of 1,175**.
5. **A ratchet.** The count cannot rise per-file without a guard going RED — which is the only mechanism
   in this milestone that acts on tier 2 at all.

**Does NOT establish:**

- **That the corpus is correct.** A `SUPPORTS` verdict says the cited source backs the sentence, not that
  the world still matches. §2.3's 57 unresolvables are precisely where that gap lives.
- **That the residual pool is small.** 1,908 tier-1 pairs are unadjudicated and 1,151 tier-2 assertions
  are unverified. The floor is now *larger* and *better named*, not smaller.
- **Anything about clause 5.** The gate is **4 of 5** and clause 5 is met only by a reading that returns
  zero. The census drains a pool; it does not read it.
