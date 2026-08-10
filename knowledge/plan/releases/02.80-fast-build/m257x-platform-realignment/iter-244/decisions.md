# iter-244 — decisions

## `D-M257x-244-1` — the section set is DERIVED from the rext tree, and the rename limit is STATED, not patched

`rext_path_guard` takes its section set from the rext tree's own top-level directories rather than a
literal list. CLAUDE.md's enumeration of the sections was **measured wrong at iter-129** (it omitted
`stack-secrets` and `playthroughs`), so it is not a source of truth.

**The cost, accepted and named:** a section that is **renamed** silently drops out of the subject instead
of firing. The alternative — scanning any `<token>/<path>.<ext>` and reporting tokens that are *not* live
sections — over-matches every two-segment path in the corpus. So the limitation is disclosed: the derived
set and its count print on **every** run, so the reach is visible rather than assumed (iter-114).

## `D-M257x-244-2` — `knowledge/` is excluded from the section set, on a measurement

`knowledge/` is a real rext directory. Including it produced **26 findings, 23 of them false**, because
`knowledge/…` in this corpus overwhelmingly names a **platform** path (`app/knowledge/architecture.md`,
`infrastructure/knowledge/service-dependencies.md`) or rosetta's own `knowledge/plan/**`.

Recorded because it is a **repeat**: `corpus_citation_guard`'s docstring already carries this exact class
(*"`knowledge/...` paths are NOT rosetta paths … 5 false REDs"*). The lesson was written down in this
repo, in a sibling guard, and re-derived from scratch anyway. A guard family's docstrings are a *corpus*;
reading the neighbours before writing a new one is cheaper than measuring it again.

## `D-M257x-244-3` — the disclosed absence is excluded by **(path, citing file)**, and printed on every run

`stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts` is cited by a sentence whose whole claim is
that the file *was never committed*. That is a correct citation of an absence, not a broken citation.

Two properties make the exclusion safe rather than a hole:

1. **It is keyed on the pair**, not the path. The same path cited from any other document still fires —
   there is a test for exactly that.
2. **It is printed on every run, green or red.** `§5`: a correct exclusion is still a defect while it is
   silent.

## `D-M257x-244-4` — `anchor_construct_guard` is NOT widened

It books an unresolvable anchor as *out of reach* (599 of 1,481) and exits 0. That is **correct for a
guard that cannot know whether the clone is present**, and widening it would make it RED on every box
that simply lacks a checkout — the failure mode `§8` rule 6 says gets a fence disabled.

The right shape is a **second instrument with a stronger precondition**: `rext_path_guard` runs only when
the rext tree is present, and where that precondition fails it exits **2, never 0**. See the protocol
addition in `corpus/ops/platform-alignment.md` §8.

## `D-M257x-244-5` — the inherited RED is repaired to the MEASURED value; the indexing debt is ROUTED, not absorbed

`tests/test_fence_registry_population_m257x.py` publishes a disclosed limit — how many of the fence family
the prose index in `stack-core/README.md` names — and it was **already RED on the tree that opened this
iter**: published `19 of 30`, measured `19 of 32` (`git archive` of `c2d9052`). `skill_invocation_guard`
(iter-239) and `toolchain_floor_guard` (iter-240) each entered the family **without** entering the README
index, and neither iter ran this suite.

**Decision, two halves:**

* The published triple is corrected to the measured one — **20 of 33** (`union`) · **20 of 32** (`census`)
  · **19 of 32** (`declaring`) — because a published limit that is *wrong* is strictly worse than one that
  is *wide*. This is mandatory, not optional: this iter's own guard moves the denominator regardless.
* **Indexing the other two is routed, not done here** (`ROUTE-M257x-244-two-fences-entered-the-family-unindexed`).
  Absorbing it would put the numerator at 22 and encode two other iters' work inside this one's close
  status — the scope-creep tripwire's own rule. The RED is reported as a **side discovery**, and it does
  not upgrade or downgrade this iter's grade.

## `D-M257x-244-6` — 2 of 5 pre-registrations refuted, and the refutations are the yield

`P-244-2` (a non-resolving path would sit in a **runnable position**) and `P-244-3` (**0 of 27** pinned
paths would fail) were both wrong. The second is the valuable one: the single failing pinned path is
precisely the one `anchor_construct_guard` nominally covers, and chasing *why it passed* is what produced
`D-M257x-244-4` and the protocol rule. **A census that only confirms its predictions has told you nothing
you did not already believe.**
