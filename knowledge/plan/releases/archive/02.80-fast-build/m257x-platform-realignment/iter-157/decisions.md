# iter-157 — decisions

## `D-M257x-157-1` — a glob is not a derivation

`repair_postcondition.discover_fences()` selected members by `glob("*_guard.py")` while **both** claims
about it were written in terms of the declaration: its own docstring (*"a fence added tomorrow enrols
itself or makes this fail loudly naming its own filename"*) and `guard_family.py:67` (*"`FENCE_KIND` — read
STATICALLY by `repair_postcondition.py`"*). **25 modules declared a kind; 23 were enumerated.** A glob is a
hand-written predicate with a wildcard in it: it fails the way a hand-written list fails, and unlike a list
it *looks* derived — iter-154's "looks-derived" corollary at a different construct. Membership now follows
the declaration.

## `D-M257x-157-2` — the partition had no third bucket, so its gap reported as a pass

A module outside the glob was classified by neither `postcondition` nor `standalone` and never reached the
`undeclared` refusal, so it was **indistinguishable in the output from a module that is not a fence**.
That is iter-150's finding verbatim, one construct up: *a value the mechanism does not classify is treated
as the safe case BY OMISSION*. The repair keeps the partition **declared** and derives its
**completeness** — the same split iter-150 chose, for the same reason.

## `D-M257x-157-3` — the widening kept the requirement it could have quietly dropped

Enrolling by declaration could have replaced the naming requirement instead of adding to it. It does not:
a `*_guard.py` that declares nothing, or declares an illegal kind, is still refused, and there is now a
**new** refusal the widening itself makes necessary — a *non*-guard-named module declaring an illegal kind,
which the new skip branch would otherwise swallow. Four preservation arms, all fixtured.

## `D-M257x-157-4` — the landing test was "does the ratchet verdict move", and it did not

Widening a registry that feeds a ratchet is only safe if the ratchet reads the same. Both newly-enrolled
modules declare `standalone`, so neither is asked for `postcondition_sites()`; the collect path is
untouched and the live verdict is unchanged (`0 site(s) reported`, `OK`). Had either declared
`postcondition`, this would have been routed rather than landed.

## `D-M257x-157-5` — `SURVEY-M257x-iter156-other-reporting-layers` was graded, not assumed

`autoverify.sh:204` captures `verify.sh 2>&1` and derives a probe count from the merged stream — the same
*shape* as iter-156's defect, and **not** the same defect: `verify.sh` writes its probe rows to stderr by
design, so there stderr **is** the subject's own voice, and the site already requires an indented `✗` to
separate rows from the summary. The property is *distinguish the subject's voice*, never *do not merge*.
Recorded so the route is not re-opened on the grep alone; it stays open for the remaining runners.

## `D-M257x-157-6` — two self-inflicted defects in the fence's first draft, and the second is the lesson

The draft called `fence_registry()` (guessed; the real name is `discover_fences`) — 11 of 12 tests failed
on an `AttributeError`, the cheap failure. The second is not cheap: `assertRegex` applies `re.search`
**without** `re.M`, so `^FENCE_KIND` meant *start of file* and the arm reported *"guard_family declares
nothing"* — a **false negative produced by an anchor that meant something other than it looked like**.
That is iter-152's `search()`-vs-`^` defect in mirror image, committed inside the fence written for this
class, one iter after iter-156 re-read that same defect. **Anchors are load-bearing in both directions and
the API decides which**; a helper regex compiled once with `re.M` (as `_DECL` is) and reused is the
structural defence, not care.
