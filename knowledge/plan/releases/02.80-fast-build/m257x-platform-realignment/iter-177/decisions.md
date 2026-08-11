# iter-177 — decisions

## `D-M257x-177-1` — harden pass 39's retraction of `16 of 27` is RETRACTED, and the correction is the LABEL

Pass 39 recorded, as its Finding 2: *"the disclosed limit said `16 of 27`; it was `15 of 26` the day it
was written… **both operands were wrong at publication**."* Measured at rext `c7f4c3d` and re-derived at
**`5b108d0`** — iter-175's own commit, reconstructed with `git archive`, the same method pass 39 used:

| derivation | owner | rule | at HEAD | at `5b108d0` | named in `README.md` |
|---|---|---|---|---|---|
| `union` | `guard_family.union` | spelled ∪ declared ∪ `EXTRA_CENSUS_MEMBERS` | 27 | 27 | **16** |
| `census` | `guard_family.census` | `union` − `CENSUS_EXCLUSIONS` | 26 | 26 | **16** |
| `declaring` | `repair_postcondition.declared_kind` | declares `FENCE_KIND` | 26 | 26 | **15** |

**`16 of 27` is an exact reading of `union`** — the function iter-175 authored in that very commit — at
both refs. Which set iter-175 enumerated is not an inference: its routed text lists the **11** missing
members by name, and that list contains `guard_family`, which `census` excludes and `declaring`
includes. Only `union` has 27 members with those 11 absent.

**Decision: the figure is restored and the retraction is retracted at the site that carries the live
claim.** The defect iter-175 shipped was real but different — it published a correct number under the
**wrong population name** (*"the census"* for a `union` reading). Pass 39 corrected the *number* onto a
**third** population without noticing there were three, and its replacement `15 of 26` is equally
correct and equally unreadable on its own.

**Why this is not pedantry.** Both surviving errors are the same error: *a count about a population is
unreadable until it names the derivation that produced it.* That is §8's iter-175 rule — **two
derivations of ONE population must be COMPARED** — written one iter earlier and not applied to a number
by the pass that quoted it.

## `D-M257x-177-2` — `FIX-M257x-h39-survey-id-embeds-retracted-figure` is CLOSED AS REFUTED, not renamed

Pass 39 routed a rename of `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` on the ground that the
id *"carries a retracted figure inside the identifier"*. The premise is falsified by `D-M257x-177-1`:
the embedded figure is **correct for `union`**. Renaming it would have propagated the retraction into
two iter progress docs and the hardening ledger, and made the correct reading the one that looks
retracted.

**Decision: close the route as refuted and keep the id.** What the survey *does* still owe is the thing
neither iter stated — **which derivation it intends the README to be complete against.** That is now
written into the route, and it is the only open part.

## `D-M257x-177-3` — the disclosure must publish ONE FIGURE PER DERIVATION, parsed as a labelled construct

The test that exists to keep this limit honest has now been wrong twice about the same sentence, and
each time the assert was weaker than the claim:

| version | assert | what it could not detect |
|---|---|---|
| iter-176 | `len(named) >= 2` | an error of one in either operand |
| harden 39 | one unlabelled `(N, M)` equality | that **three** populations answer to one name |
| **iter-177** | every `**N of M**` carries `(\`derivation\`)`, each checked against **that** derivation, and **all three** must appear | — |

**Decision: the assert is the whole claim.** An unlabelled pair is RED (there is no population to check
it against); a label `_derivations()` does not know is RED; a derivation with no published row is RED.
So a fourth derivation cannot enter the tree the way the third one did — silently, by being nobody's
declared subject.

**Deliberately not done:** collapsing the three into one. Each has a live consumer and a written reason
(`D-M257x-177-4`). Collapsing them would be a narrowing that grades a real distinction green — iter-158's
rule.

## `D-M257x-177-4` — the two 26-member derivations are compared by MEMBERSHIP, because their COUNTS agree

`census` and `declaring` are **both 26** and differ by **one member in each direction**:

* `guard_family` — IN `declaring` (it declares a `FENCE_KIND` so `repair_postcondition` can derive its
  registry from the declaration), OUT of `census` (`CENSUS_EXCLUSIONS` says why: running the family
  runner inside itself recurses, and its verdict IS the family verdict).
* `repair_postcondition` — IN `census` (via `EXTRA_CENSUS_MEMBERS`; the runner invokes it), OUT of
  `declaring` (it declares no `FENCE_KIND` — it is the module that *reads* the declaration, and a rule
  keyed on the declaration cannot enrol its own reader).

**Both dispositions are correct.** The defect is that nothing said they differ, and **every comparison
by count reads green** — which is exactly how one population came to publish `15 of 26`, `16 of 26` and
`16 of 27` simultaneously with nothing going RED.

**Decision: assert the symmetric difference, not the cardinality.** Each difference carries a written
disposition; a third one is RED in either direction. Shipped with the control that makes it
non-obvious — `test_mutation_control_a_COUNT_comparison_reads_GREEN_on_this_disagreement` asserts that
the counts *do* agree today, so if the coincidence ever ends the control goes RED **on purpose** and the
characterisation is re-derived rather than quietly deleted.

## `D-M257x-177-5` — the three derivations are pinned into ONE ALGEBRA

`union == census | declaring`, asserted. Without it the three sets are three opinions and a fourth can
appear between them; with it, any change that breaks the identity names itself instead of moving a count
by one where nobody is looking. Verified true at HEAD and at `5b108d0`.

## `D-M257x-177-6` — the hardening ledger's own entry is ROUTED, never edited here

`hardening-ledger.md` is owned exclusively by `/developer-kit:harden-mstone-iters`; this skill is
forbidden from writing it. Pass 39's Finding 2 and its routed-forward list both carry the retracted
retraction.

**Decision: route it as `FIX-M257x-iter177-ledger-carries-a-retracted-retraction` (Fate 3, handler = the
next harden pass), and put the correction in the live instrument instead.** That is also the honest
shape: an instrument that carries the current reading is worth more than a ledger entry that carries it,
and the ledger's owner corrects its own record. iter-175's and iter-176's `progress.md` are left
untouched for the same reason a harden pass leaves them untouched — an iter's record is what that iter
knew.
