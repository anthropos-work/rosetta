# iter-241 — decisions

## `D-M257x-241-1` — the repair is DISCLOSURE, not a bigger clone set

The tempting fix is to clone the six frozen-legacy repos into `stack-demo` so the guard checks 27
citations instead of excusing them. It is the wrong move twice over:

* it makes the number look better **without checking anything a fresh box could check** — the gate's
  subject is a cold bring-up, and a cold bring-up will never have them;
* it is `§8`'s fetch rule one level out — a clone set curated to make a fence green is *measuring a
  memory* of the platform, not the platform.

So the guard now **states its reach** (numerator, denominator, and the clone-set roster it read) and the
six leftovers stay untouched. A reader comparing two runs can now see that one checked 98 citations and
the other 82.

## `D-M257x-241-2` — the six leftover clones are NOT deleted

`stack-demo/` holds `cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner` and
`storage` — clones a fresh `make init` + `ensure-clones.sh` does not create. They are pre-`838d907`
leftovers and they are **useful**: 107 corpus citations point into them, and having them locally is how
several of this milestone's readings were possible at all.

Deleting them would destroy read access to the pre-merge source that `repos.yml`'s own header comment
tells you to obtain by hand. **The defect was never their presence — it was that nothing said the
verdict depended on it.** They stay; the dependence is now printed.

## `D-M257x-241-3` — the citation TOTAL must be clone-set-invariant, and that is tested

The new verdict carries `N of M`. `N` (unchecked) is a property of the box; `M` (total citations in the
map) is a property of the **map** and must not move when the clone set does — otherwise the fraction
would be two moving numbers and would say nothing.

Pinned by `test_a_smaller_clone_set_reports_a_LARGER_unchecked_count`, which asserts both directions at
once: the numerator rises on a restricted set **and** the denominator stays put (109 on both). A
regression that made `M` clone-set-dependent would pass a naive "the number got bigger" test and fail
this one.

## `D-M257x-241-4` — the wider citation surface is disclosed, not fenced, in this iter

The guard's subject is one file (the migration map, 109 citations). Corpus-wide there are **107** more
citations into the same six repos, graded by nothing. Fencing those is a real piece of work — it needs
the base-directory model that iter-241's path census showed is not mechanical — and it is not this iter's
planned scope.

Recorded in `platform-alignment.md` §8 as a **stated gap** with its number, and routed. `§5` — *a
CORRECT exclusion is still a defect while it is silent*; it is no longer silent.
