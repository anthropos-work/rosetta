# iter-209 — decisions

## `D-M257x-209-1` — the widening is bounded by a CONTRACT, not by "more is better"

The obvious move once the source set is shown to be 94 of 2,560 is to widen it. The obvious move is
wrong: `knowledge/` is planning, `iter-NN/raw/` is captured evidence quoting corpus text verbatim, and
both would produce findings that are not defects. The guard's own docstring records that three of its
first four runs existed only to kill a false-positive class.

So the widening is **exactly the set `CLAUDE.md` says must be maintained together with the corpus** —
a boundary the repo drew, not one this iter invented — and it was **measured before being adopted**:
20 documents, 135 citations, **0 C1 findings of 133**, 1 C2 finding. Had that pre-measurement produced
false REDs, the correct outcome was to narrow or abandon, and the `overview.md` said so before the
reading.

**The residual exclusion stays, and is now SIZED**: 2,446 documents. `§5` — *a correct exclusion is
still a defect while it is silent*; it is no longer silent.

## `D-M257x-209-2` — the contract list is PARSED from CLAUDE.md, never restated in the test

The arm could have hard-coded the five skill paths. That would pass today and rot the first time the
contract gains a twelfth file — the precise failure this milestone has now recorded a dozen times
(`ENV_GATED`, `LANGUAGE_EXCLUDED_SECTIONS`, the ceiling-arm's own two-entry name list at iter-207).

So `_contract_files()` locates the `Interconnected Documentation` heading and parses the numbered
backticked paths beneath it, and the arm asserts it found **≥ 5** so a changed list shape cannot leave
the arm silently measuring nothing (`§9`). The staged mutation control builds a tree whose `CLAUDE.md`
binds a file outside every source glob, proving the arm can fire on something other than today's tree.

## `D-M257x-209-3` — the independent reading was DISCARDED in favour of the guard's own machinery

The first pass of this iter wrote its own GitHub slugger and reported **97 unresolved of 190**.
Adjudicated, it was the instrument: it collapsed double spaces (an em dash between words leaves two)
and stripped `_`. The shipped `heading_slugs` is right, and deliberately emits two variants — with and
without underscores — which is why `symptom-und_err_connect_timeout-…` resolves.

**16× wrong, entirely in one direction.** iter-201 measured the same asymmetry on a different class
(18 false-RED / 0 false-GREEN) and the lesson repeats: when the question is about a fence's **scope**,
hold its **machinery** fixed. Re-deriving both at once produces a disagreement you cannot attribute —
`§5` iter-175, *two derivations of ONE population must be COMPARED, or the weaker one is a silent
census.* Here the weaker one was this iter's.

The rejected reading is not deleted from the record: its 97 is what makes the 6 meaningful.
