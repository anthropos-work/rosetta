# iter-122 — PRE-REGISTRATION of the CLAIM CENSUS

**Sealed 2026-08-07, before a line of the instrument was written.** Committed on its own, ahead of any
code, exactly as `TOK-07` (iter-116) and `TOK-08` (iter-117) were. Nothing below may be edited after this
file is committed; a correction is an appended, dated addendum that says what it corrects.

---

## 0. What the user ruled, and what this iter is therefore for

The user was given the honest choice — close at 4 of 5 with clause 5 documented as measured-and-unmet, or
fund a materially different instrument. **They chose the instrument.** So:

> **Enumerate and verify every claim in `corpus/services/**` + `corpus/architecture/**` instead of
> sampling them.**

**Clause 5 is NOT re-cut, reinterpreted, narrowed or argued.** It is met only by a reading that returns
zero, over the same 40 files. **The graded read remains clause 5's instrument.** The census built here is
how the *pool gets drained* so that a reading can honestly return zero — it is **not** the grader, and
§F4 below makes that falsifiable rather than merely promised.

## 1. The subject, measured before the instrument existed

Taken 2026-08-07 at `rosetta` `ad1b82d`, by two throwaway scripts, and stated here as the census's
denominator inputs:

| quantity | value |
|---|---|
| files in scope (`corpus/services/*.md` + `corpus/architecture/*.md`) | **40** |
| total lines | **11,042** |
| **prose lines** (outside ``` fences) | **9,543** |
| markdown links `[..](..)` in prose | **566** |
| source-file tokens (`*.go`/`*.ts`/`*.yml`/…) in prose | **1,695** |
| …of which carry a `:NN` / `:NN-MM` line pin | **650** |
| sha-shaped tokens (`[0-9a-f]{7,40}`) in prose | **1,122** |

These are *candidate* counts from throwaway regexes, not the census's own output. The census will publish
its own, and where they differ from these the difference is itself reported.

## 2. The design problem, stated before it is solved

**A claim is not syntactically marked.** There is no token that means "this sentence asserts something
about the platform." Every previous instrument in this milestone either (a) sampled sentences and let a
reader decide (the graded read — measured at **~35 % test-retest recall**, iter-119), or (b) checked a
*syntactic* property that is decidable without reading the sentence (`corpus_citation_guard`: does the
reference **resolve**; `anchor_construct_guard`: does the anchor land on a **non-blank** line).

Neither reaches the two questions that matter:

- **Does the cited source support the proposition the sentence asserts?** — `anchor_construct_guard`
  explicitly cannot answer this (it detects *resolves-to-blank*, not *resolves-to-the-wrong-construct*),
  and iter-121 measured why a regex will not close the gap: only **28 of 511** backticked `file:line`
  citations supply their own expected content, and binding quote→anchor is **≥ 20 % false-positive** even
  on those.
- **Is a factual assertion carrying any evidence at all?** — no fence in the family has ever looked at an
  *uncited* sentence.

So the census is built in two tiers, and **each tier's reach and its blind side are declared here, before
either is built.**

---

## 3. TIER 1 — cited claims

### 3.1 What it enumerates

Every **(claiming unit, citation)** pair in the 40 files, where

- a **claiming unit** is the smallest containing block of prose — a bullet, a table row, or a
  blank-line-delimited paragraph — that contains the citation, and
- a **citation** is a reference to a source *outside the claiming unit*: an intra-corpus doc path or
  markdown link, a platform/tooling source token `path/to/file.ext[:NN[-MM]]`, or a commit sha.

### 3.2 What adjudication asks — ONE question per pair

> **Does the cited content support the proposition the claiming unit asserts?**

Verdict ∈ **`SUPPORTS` · `DOES-NOT-SUPPORT` · `UNRESOLVABLE` · `NOT-A-CLAIM`**. The census *materializes*
the cited content (the target lines from the clone set at `stack-demo/` @ platform `0c91421`, or the
target lines of the corpus doc) and presents it next to the claiming unit, so the adjudicator never has to
go looking. **The materialization is mechanical; the verdict is judgement.** That split is the whole
design: iter-121 measured that the verdict half cannot be regexed, so the census does not pretend to.

### 3.3 Declared blind sides of tier 1

1. **A `SUPPORTS` verdict does not mean the sentence is TRUE.** It means the cited source supports it. A
   sentence whose citation supports it and whose *world* has since moved is `SUPPORTS` here and a defect
   in the world. This census does not close that.
2. **A claiming unit with no citation is not in tier 1 at all** — that is tier 2's subject, and tier 2
   cannot verify truth either (§4.3).
3. **Bare `:NN` pins are excluded**, for the reason `corpus_citation_guard` already measured: of 387 lines
   carrying one, only 4 name exactly one corpus doc; an early draft resolved them against the last-named
   document and reported **256 findings, all false**.

---

## 4. TIER 2 — uncited factual assertions

### 4.1 The principle, and its precedent

The iter-093 hedge fence (`unreadable_repo_claim_guard`) requires a claim about a repo in no clone set to
**say it is not a measurement**. Tier 2 generalises exactly that:

> **Every factual assertion about the platform either carries a citation, carries a hedge, or is a
> defect.**

That is mechanically checkable in a way *"is this sentence true"* is not — and it is the half of the
corpus no fence in this family has ever reached.

### 4.2 What it enumerates

A prose sentence is a **tier-2 defect** iff all of:

- it is outside a ``` fence, is not a heading, is not a table separator, and is **indicative** (not an
  imperative instruction, not a question);
- it contains at least one **subject token** — a token naming a platform/tooling artifact (a service or
  repo name, a source path, an env var, a port, a schema/table, a container, a compose profile, a version
  or a sha);
- **no citation is in scope** — where "in scope" is the sentence's own line *plus* its smallest containing
  block (bullet / table row / paragraph). Block scope is **deliberately generous to the corpus**: it
  UNDER-flags rather than false-REDs, which iter-117 established is the correct direction for a fence;
- **no hedge marker is in scope** — the marker set extends `unreadable_repo_claim_guard.MARKERS`, and is
  a set of phrases the corpus already uses, never one mandated magic token.

### 4.3 Declared blind sides of tier 2

1. **Tier 2 does not measure truth.** An uncited-unhedged assertion may be perfectly true. Tier 2 measures
   **unevidenced**, and reports it as that word.
2. **Subject-token detection is a proxy for "is a factual assertion".** It will admit some
   non-assertions (a navigational sentence that happens to name a service) and miss some assertions
   (a claim made entirely in English with no artifact token). §F1 measures that miss rate rather than
   assuming it away.
3. **Block scope means a citation anywhere in a long paragraph exonerates every sentence in it.** That is
   the generous direction, chosen deliberately; the resulting count is therefore a **floor**, and must be
   printed as one.

---

## 5. THE PRE-REGISTERED FALSIFICATION

The census is **falsified**, and this iter reports the measurement rather than a number, if any of the
following fires. Each is checked and its result published whichever way it lands.

### F1 — enumeration incompleteness (tier 2's recall)

A **blind hand-audit** of a random sample of **60 prose lines**, drawn with a seed fixed in this file
(`seed = 122`, `random.Random(122).sample(range(len(prose_lines)), 60)`, prose lines in file-then-line
order over the 40 files) **before the enumerator is run**, adjudicated by an agent that is **not shown the
enumerator's verdicts**. If the audit finds factual assertions about the platform that the census placed
in **neither** tier, and those misses exceed **10 %** of the audited assertions, then **the census is a
sample wearing a census's name and this iter says exactly that.**

> Recall floor pre-registered: **≥ 90 %**. Reported either way, with numerator and denominator.

### F2 — adjudication cost (tier 1's reach)

Throughput is measured on a **pilot of the first 40 tier-1 pairs in census order**, then extrapolated. If
the extrapolation exceeds this run's bounded effort, the iter reports:

> *"exhaustive tier-1 adjudication was not achievable in run 78"*

**with the measured numerator, denominator and extrapolation**, and **publishes no percentage of corpus
correctness** derived from the adjudicated subset. A partial adjudication may report *its own* false rate
against *its own* denominator, named as such, and nothing else.

### F3 — vacuity

If **any** tier's enumeration is empty, or if a mutation control does not fire, the run is **void**
(exit 2), never a pass. Ten vacuity-class defects have been caught in this milestone; the most recent was
`fence_provenance` blind to its own subject with a silent suite.

**Every mutation asserts it applied before its result is read** (§5 rule 53) — three of harden pass 26's
mutations silently failed to apply and each read as *"the controls survive."*

### F4 — the enumerator must not quietly become the grader

**No number this instrument produces is reported as `P`, as `N`, or as a clause-5 verdict.** Clause 5's
instrument remains the frozen graded read (sha `3858ec53…`). If any deliverable of this iter contains a
sentence asserting clause 5 is met without a graded read that returned zero, **that sentence is a defect
of this iter** and is to be booked as one.

### F5 — the direction of surprise, declared in advance

Bands, sealed now, each of which can fail:

| quantity | pre-registered band | basis |
|---|---|---|
| tier-1 `DOES-NOT-SUPPORT` rate on the adjudicated set | **≥ 4 % and ≤ 25 %** | iter-119 measured ≥ 13.3 % error on its judged set; iter-120 closed 8 wrong-construct citations out of ~22 judged |
| tier-2 uncited-unhedged count over 9,543 prose lines | **≥ 200 and ≤ 2,500** | no prior measurement exists — this is a genuine prediction, and the widest band here is the honest one |
| tier-1 pairs enumerated | **≥ 900 and ≤ 3,000** | 566 links + 650 line-pinned tokens + shas, minus de-duplication |

**A band that fails is the headline**, exactly as `P = 37` vs `P ≥ 15` was TOK-07's headline and
`P = 22` vs `P ≥ 19` was TOK-08's.

### F6 — a reach metric names its denominator or prints no percentage

iter-114's rule, applied to this instrument's own output: every percentage the census prints carries the
denominator it is over, or it is not printed.

---

## 6. CONTROLS shipped with the fence

- **Anti-vacuity control** — the census over a fixture corpus with a known planted population must return
  the planted counts exactly; an empty enumeration exits 2. This control fires today (it is asserted
  against a fixture, not against the live corpus, so it cannot be satisfied by the live corpus happening
  to be non-empty).
- **Mutation controls**, each asserting `applied == 1` before its result is read:
  1. **strip a citation** from a cited claiming unit → that unit must move **tier 1 → tier 2 defect**;
  2. **add a hedge** to an uncited assertion → it must **leave** the tier-2 defect list;
  3. **remove the hedge vocabulary** → the tier-2 count must **rise**;
  4. **blank the subject-token table** → tier 2 must **collapse to zero and the guard must exit 2**, not
     report a flattering clean sweep.

## 7. What the census establishes that a reading cannot — stated as a claim to be judged, not a boast

A reading at ~35 % recall cannot demonstrate the **absence** of what it did not see. A census can
demonstrate that **every member of an enumerated class was individually addressed** — and, where it could
not address them, it can say **how many it did not**, by name, with the denominator. Whether that is
enough to let a subsequent graded read honestly return zero is **not decided by this file** and is not
claimed here.

---

_Sealed at `rosetta` `ad1b82d`, `rext` `8b41cd2`, platform `0c91421`. Nothing below this line existed when
this file was committed._
