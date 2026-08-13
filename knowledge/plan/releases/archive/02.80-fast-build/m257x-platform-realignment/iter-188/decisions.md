# iter-188 — decisions

## `D-M257x-188-1` — the target was selected by iter-187's selector, not by a re-survey

iter-187 routed its residual **with a mechanical selector** — *a registry keyed by a CONTAINER whose
justifying reason is a property of the container's CONTENTS* — which is precisely what iter-185's
residual lacked and what iter-186's lesson said to build. This iter **ran** it over the 30 module-level
registries in `stack-core` rather than re-surveying by judgement, and the strongest hit was the one with
the largest consequence: the prune list of the walk that builds the **citation resolver's** index.

That is the class closing as `§8` prescribes — *by an enumeration that keeps running*, not by a repair.

## `D-M257x-188-2` — the first environment reading was WRONG and is corrected in place, not quietly

The iter opened by measuring `corpus/` (0 of 92 pruned) and a stand-in clone tree (537 of 1,718), on the
belief that `corpus/` alone was the default subject. **It is not.** `claim_census_guard.py:768` falls
back to `root / "stack-demo"` when `--clones-root`/`CLONES_ROOT` is unset, and that tree exists on this
box — so the **default** run is the large one: **21,610 indexed / 50,357 pruned / 70.0 %**.

Both the iter `overview.md` and the module comment carry the correction with the reason, per `§5` rule 8
and the retraction discipline. The stand-in figure is dropped rather than kept beside the real one: a
number measured against the wrong subject is not evidence about the right one.

## `D-M257x-188-3` — SIZE is printed, SHAPE is asserted, and the split is the deliverable

Three shapes were available:

1. **Assert the size** (e.g. *"the prune list removes < X %"*). Rejected: vacuous at 0 % on a bare
   checkout, wrong at 70 % on a populated one. It would be a fence pinned to an environment.
2. **Assert nothing and print nothing** — the status quo, and the defect.
3. **Assert the shape, print the size.** Taken. Everything environment-independent (reasons present,
   membership set derived, dot-subsumption marked both ways, the walk having no second name rule) is
   fenced; the size is derived by `prune_census()` and printed by the guard on every text-mode verb.

What a fence *can* prove about a size is proved: that the derivation can return non-zero (`§9`), and
that its attribution **partitions the registry exactly** — fired ∪ inert ∪ dot-subsumed == the registry,
pairwise disjoint. That is iter-186's partition property one layer down.

## `D-M257x-188-4` — the 5 dot-subsumed members are KEPT and MARKED, not deleted

`.git`, `.next`, `.venv`, `.turbo`, `.pnpm-store` can never be the reason a directory was pruned: the
walk reads `d not in _SKIP_DIRS and not d.startswith(".")`. Deleting them was the tempting repair and is
rejected — a future caller pruning **without** the adjacent rule would then get a list that is wrong,
and the deletion would erase the evidence of the subsumption rather than record it.

They are kept, prefixed `[dot-rule] ` in their reason, and the marking is asserted **both ways**: a
marked member that is not dot-named, or a dot-named member that is not marked, both fail. Plus an arm
over `basename_index`'s source, because the subsumption claim is only true while the dot-rule is there.

## `D-M257x-188-5` — the escalation condition was checked with a number, and did not fire

`overview.md` pre-registered: *if any current corpus citation resolves into a pruned directory, that is a
live false defect.* Measured over `corpus/**.md` + `CLAUDE.md`: **2,093** citation-shaped path tokens,
of which **1** crosses a pruned name — `` `.next/static/chunks/*.js` `` in `demo/frontend-tier.md`, a
wildcard describing build output rather than a citation to a file, and dot-named so the dot-rule prunes
it regardless. **Latent, and sized.** The second condition (removing dot members changes the walk) was
not tested by removal because `D-M257x-188-4` chose not to remove them; the equivalent claim — that the
walk applies both rules — is asserted directly instead, and mutation-proven (M5).

## `D-M257x-188-6` — no new `*_guard.py`; the new module is a test, and the fence indexes are unmoved

The deliverable is arms over an existing guard, not a new guard. `test_claim_census_skip_registry_m257x.py`
is a test module; `guard_family`'s census population, `discover_fences`, and the README fence-index
triple are all unchanged — verified green in the neighbourhood run (`test_fence_registry_completeness` +
`test_fence_registry_population` + `test_guard_family`, 95 passed under unittest / 120 under pytest).
