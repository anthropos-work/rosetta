# iter-230 — decisions

## `D-M257x-230-1` — an unresolvable sha is PARTITIONED, never graded false

The census's headline could have been *"14 of 138 corpus shas are broken (10.1 %)"*. It would have been
wrong twice over: eight were the instrument's own missing clones, and the remaining six name repos that no
clone set on this box contains. A sha from an uncloned repo and a sha somebody invented are
**byte-indistinguishable** — `git cat-file -e` returns the same non-zero for both.

So the reported figure is **132 of 132 measurable resolve, 6 UNMEASURED**, with the unmeasured set
enumerated by repo. `§5`'s scoped-green rule in its arithmetic form: **a 0 % error rate over a population
that silently excludes its own largest member is not a rate.**

## `D-M257x-230-2` — no clone was fetched to close the residual

Cloning `infrastructure` would convert 61 UNMEASURED sites into measured ones and is obviously tempting.
It was not done here, for two reasons:

1. **It changes the clone set**, which is the substrate every other guard in this milestone measures
   against — including `clone_pin_guard`, which asserts `clones.pin.json` names *exactly* the `repos.yml`
   repos plus two sanctioned extras. Adding a repo mid-iter would move a fence for reasons unrelated to
   the fence, on the same night `ROUTE-M257x-222-pin-advance-needs-a-reproof` is explicitly holding the
   clone set still until gate clause 1 is proven.
2. **It is a second deliverable** and fires the scope-creep tripwire.

Routed as `ROUTE-M257x-230-82-sites-on-uncloneable-evidence`. The residual is disclosed in the corpus in
the meantime, which is the part that could not wait: a reader who wants to re-check the cms M810 verdict
now learns immediately that no clone here can answer, instead of concluding the sha is wrong.
