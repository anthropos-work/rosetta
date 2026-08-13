# iter-126 — decisions

## `D-M257x-126-1` — UNCLONABLE is not UNRESOLVABLE, and the distinction is what takes a guard off exit 2 without silencing it

`platform_alignment_guard` returned exit 2 because the corpus had started citing `infrastructure` and
`db-backup` — repos no stack clones. **The verdict was correct on a changed subject**, and the subject
changed because iter-123 did the right thing: it cloned the repo that settled four standing questions.

Two conditions were sharing one bucket, and they are not the same thing:

| | what it is | whose defect |
|---|---|---|
| **unresolvable** | a path the guard cannot find and cannot account for — a typo, an unqualified path, a repo that does not exist | **the citation's** |
| **unclonable** | a repo **the map itself documents**, simply not in any stack's clone set | **the substrate's** |

Collapsing the second into the first forced a false choice: silence the guard (forbidden), or stop citing
`infrastructure`. **Decision: widen the subject to the ORG and DISCLOSE the clone-set limit** — unclonable
heads print on every run, ride in `--json`, and **the verdict sentence itself** becomes *"OK OVER ITS
REACH — … N citation(s) into M repo(s) the map documents but no stack clones were NOT checked."*

**The qualifier is IN the verdict for a mechanical reason, not a stylistic one:** `guard_family.run_one`
reports `lines[-1]` for a green member. A qualifier printed anywhere else is **invisible in the family
view**, which is precisely how an over-claiming green survives — the failure `fence_provenance` was built
for, and the one iter-123's own `KNOWN_WEAKNESS`-gated-on-the-wrong-verb bug reproduced.

## `D-M257x-126-2` — the anti-vacuity gate is the whole decision; without it the widening is a silencer in a disclosure's clothes

A head is admitted as `unclonable` **only if the map's own services or census table names that repo.**

Without that gate, `infrastructre/terraform/main.tf:10` — a typo — would launder itself into the excused
bucket, and this change would be indistinguishable from turning the guard off. **An empty mapped-repo set
excuses NOTHING**, which is the safe direction, and it is asserted rather than intended.

**Controls, and each asserts it applied** (§8: a mutant that did not apply and a mutant that survived are
indistinguishable in the output):

| control | what it kills |
|---|---|
| **anti-vacuity** | an undocumented head stays `unresolvable` and the gap is still recorded on the refusal stream |
| **mutation — withdraw the documentation** | delete the map row that documents the repo, and the excuse goes with it |
| **negative control** | a tree citing nothing unclonable still gets the **unqualified** OK, so the qualifier carries information |

**And the gate was killed from the guard's side to prove the controls bind:** removing `head in
mapped_repos` turns **exactly one** test red — a **named kill**, not a suite-wide smear. 51/51 restored.

**The widening immediately bought a finding the exit 2 was masking.** With assertion F able to run to
completion, `messenger`'s row was caught citing `:622`/`:664` bound to `messenger/terraform/main.tf`, a
**121-line file** — they are `infrastructure`'s `services.tf`. **A guard stuck at exit 2 is not a guard
that is being careful; it is a guard that is not looking.**

## `D-M257x-126-3` — a corpus-side derivation over line-pinned sites reads the corpus AT A REF

Re-deriving the backlog's no-sha class **at HEAD** gave **22**; at **`afe58ac`**, the ref iter-123 named,
it gave **7** — reproducing iter-123 exactly. **A 3.1× inflation, manufactured entirely by the substrate.**

`D-M257x-122-4` established this for *platform* clones: *"a stale substrate does not merely fail to
confirm a claim — it produces evidence AGAINST a true one."* **Decision: the rule is not about platform
clones. It is about pins.** A blank-line-delimited block at `security_compliance.md:266` is a different
paragraph today than it was 30 commits ago, and every one of those 15 extra "findings" would have been a
real edit to a citation that was never broken.

**Reproducing the prior iter's number exactly is the control that says the method is right** — and it only
exists because iter-123 published a number this iter could reproduce.

## `D-M257x-126-4` — a pin you have just called unverifiable is REMOVED, not hedged

`B11-020` cited a bare `README.md:21` in a sentence about the shared `ai` library. **Qualifying it made it
worse**: the resolver bound the qualified path to **`studio-desk` @ `41ee357`** — an unrelated repo — and
landed on a blank line. `anchor_construct_guard` went RED **on the qualification, in the session that
wrote it**.

`ai` is a private Go module **no stack clones** (not in `repos.yml`; pulled at Docker build via
`GOPRIVATE`; `ls stack-demo/` has no `ai`). **So no `file:line` into it is verifiable from here at all.**

**Two costs, recorded rather than tidied:**

1. **A manufactured hedge was written and withdrawn.** The first correction said the anchor was *"not
   verifiable from here"* — true — **while leaving the unverifiable pin in place.** The run-80 directive
   forbids manufacturing hedges for facts somebody can measure; **its mirror is equally binding: do not
   keep a pin you have just said nobody can check.** A hedge is not a licence to leave the artifact.
2. **The retraction had to be rewritten in the fence's vocabulary** (§8, iter-98). The first draft
   explained the removal *by quoting the removed pin in backticks*, and the guard parsed the quotation as
   a live citation and stayed RED. **A retraction written in the vocabulary the fence enumerates is
   indistinguishable from the claim it retracts** — the fence is right to refuse it, and the fix is to
   name the document in prose.

## `D-M257x-126-5` — an intra-corpus citation is a FALSE MEMBER of the no-sha class

`B01-021` cites `ops/demo/stories-spec.md:599` — this repository. **A sha does not apply**: it would be a
ref of *this* repo, the citation resolver reads the working tree by design, and both
`corpus_citation_guard` and `anchor_construct_guard` already resolve and grade every intra-corpus anchor
on **every run**.

**Decision: it stays, and the class is narrowed rather than the site repaired.** A standing fence that
grades an anchor continuously is a **stronger** control than a pin, not a weaker one — demanding a sha
here would add ceremony and subtract nothing. The no-sha class should exclude intra-corpus anchors, and
the enumeration says so.
