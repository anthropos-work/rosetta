# iter-277 — decisions

## D-M257x-277-1 — the census was contaminated by the act of running it, and that was measured rather than argued

**Context.** The full `stack-core` census returned `rc=1`, 12 failed / 2211 passed. Six failures were
literal-ceiling and derivation-registry assertions over the `rosetta-extensions` tree.

**The plausible story was wrong.** iter-276 had just added long doc comments to `jobroleref.go`
containing a numeric step function (`1 → 68`, `2 → 54`, `3–4 → 45`), so "the new comments breached a
measurement-literal ceiling" fitted the evidence and would have been accepted by a reader.

**Falsified by measurement.** The identical counter was run against three trees, the two commits
extracted with `git archive` so only **tracked** files were present:

| tree | COMMENT (ceiling 236) | DOCSTRING (ceiling 240) |
|---|---:|---:|
| `2833a64` (pre-iter-276) | 236 | 240 |
| `0a8674e` (iter-276) | **236** | **240** |
| working tree, venv present | 279 | 282 |

iter-276's comments added **zero**. The +43 / +42 were `stack-core/.venv-check/lib/**/site-packages`
— pip's vendored `urllib3`, `pygments`, `rich`, `tomli` — installed there to run the census.

**Decision.** The venv is relocated **outside** the scanned tree and the six failures are dispositioned
**artifact**. The finding is routed as
`ROUTE-M257x-277-the-census-cannot-be-run-from-inside-its-own-tree`: `stack-core` ships no environment,
the natural place to create one is inside the tree, and doing so breaches two ratchets **that sit exactly
at their ceilings**. Every future census run reproduces this, at 53 minutes per discovery.

**The generalisable form, and it is this milestone's own subject:** *an instrument installed inside its
subject measures itself.* The tell was the coincidence — **both** ratchets sitting *exactly* at their
ceilings beforehand. Exactness like that is evidence about the ratchets being healthy, not about the
breach being real.

## D-M257x-277-2 — the elided route id, and the repair that re-created it

`route_disposition_guard` went RED on iter-276's own close: the closed handler was written with an
ellipsis mid-id, which the guard parses as a truncated stem and — per `§5` rule 73 — *"reads as live
backlog in every brief that quotes this queue."* A **landed** fix would have shown up as **open** backlog
in every downstream brief.

**The first repair re-tripped the guard.** The id was written in full, but the footnote explaining the
rule *quoted the elided form* — putting the same stem back on the same line, inside the sentence
describing why it must not be there. Repaired again, this time describing the elision without spelling
it, with an explicit warning in the note because the next reader's instinct will be identical.

**The generalisable form:** *guards read prose, including prose about guards.* A worked example of a
defect is an instance of the defect unless the example is neutered. This is the second time this
milestone has produced a demonstration that was itself the thing demonstrated.

## D-M257x-277-3 — no `P` is claimed, and that is the disciplined answer rather than the evasive one

Clause 5's semantic half was **not measured**, and this iter publishes **no `P`**.

**Why not a quick reading.** iter-131 — the last one — used 14 seats and its headline finding was that
its **test-retest overlap with iter-119 was ~0**: two consecutive readings produced almost disjoint
predicate sets. A reading at that recall cannot enumerate the pool, which is precisely what `TOK-08`
concluded when it replaced sampling with census. A single-session scoped reading would have produced a
number, and the number would have been evidence about its scope alone — the standing failure mode this
milestone has already been caught by.

**What was measured instead, completely:** the two clause-5 defects iter-131 routed forward were checked
over their **named population** (a census, not a sample) and both are repaired — the `infrastructure`
hedge survives only as retraction prose, and `architecture_overview.md` §4 now lists the correct five
modules with `ai` explicitly excluded.

**Decision.** Report the census, report the two closed routes, and state plainly that clause 5 is **NOT
met** and its semantic half remains unmeasured since iter-131. Clause 5 is met only by a reading that
returns zero; the user has ruled four times that it is not re-cut, reinterpreted, narrowed or argued, and
declining to publish a weak number is the same discipline as declining to re-cut the clause.

## D-M257x-277-4 — shipping the tooling put the corpus out of date, by construction

`clone_drift_guard` is RED because **iter-276 made it so**: `rosetta-extensions` is now at `0a8674e74`,
**22 commits past the nearest of 12 shas the corpus cites across 44 sites**.

This is not a fault in the fix and not a fault in the guard. It is structural: this milestone's two
halves are coupled, so **a green clause 2 mechanically costs clause 5 something**. Routed as
`FIX-M257x-277-corpus-cites-a-rext-sha-that-no-longer-exists` and named the next iter's highest-value
target, because it sits directly under the user's third limb — *"and the corpus reflects that."*

**Worth carrying forward as practice:** the citation debt should be paid in the iter that creates it,
not discovered by a fence 50 minutes afterwards.
