# iter-176 — decisions

## `D-M257x-176-1` — the fence is a TEST, not a `*_guard.py`, and the reason is this iter's own subject

A new `*_guard.py` must join **three** registries (`guard_family.INVOCATIONS`, the postcondition
registry's declaration contract, and — if participating — the ratchet baseline), and the last two iters
are a record of what that costs when one is missed.

**Decision: ship it as `tests/test_fence_registry_population_m257x.py`.** Its subject is the *set of
registries*, not the corpus; it needs no CLI, no `FENCE_KIND`, and no invocation entry. It joins **zero**
registries, so shipping the fence that counts registries does not increment the count it reports. This
is the same call iter-157 made for `test_fence_registry_completeness_m257x.py`, and it is the precedent
rather than a new idea.

## `D-M257x-176-2` — the key is the PATH, never the line

Line numbers drift on every edit, and a classification keyed on one rots into a chore. The obligation
the table records is per-file — *"if a fence is added, must this file be touched?"* — so the key is the
repo-relative path and every site in that file shares its verdict. The battery already proves the shape
matters: it holds **two** literals (`:70` and `:81`) and one obligation.

## `D-M257x-176-3` — `ast.Call` arguments are IN the predicate, and excluding them would have blinded the fence to its own motivating case

The seed list that started this thread is not a literal at all — it is
`battery_stage.local_deps(STACK_CORE, "markdown_structure_guard.py", …)`, a **call**. A predicate over
list/tuple/set/dict literals only is elegant, defensible in a sentence, and structurally blind to the
registry `FIX-M257x-iter174-…` is about.

**Decision: call arguments count.** The cost is a wider net; the alternative is a fence that cannot see
the case it was commissioned for — iter-158's rule stated as a design constraint rather than a review
finding.

## `D-M257x-176-4` — the `fixtures` exclusion was REMOVED after measuring it at zero

The first draft skipped `fixtures/` alongside `.git`, `__pycache__` and `node_modules`, on the reflex
that fixture trees are noise. Measured before shipping: the site count with and without it is
**identical — 5 either way**.

**Decision: remove it.** An exclusion that changes nothing buys nothing, and it silently forecloses a
registry living in a fixture. A narrowing whose effect nobody measured is this milestone's defining
defect class; adding one *inside the fence written against that class* would have been the joke telling
itself.

## `D-M257x-176-5` — the disclosed limit is a TEST, not a sentence in a docstring

A **prose** index — a markdown table of the guards — is not a collection literal, so this fence cannot
reach `stack-core/README.md`, which iter-175 measured at **16 of 27**.

**Decision: state the limit as an executable fact.** `test_the_disclosed_limit_is_STATED_not_assumed`
asserts that the README really does name several fences and really is not a site — so if the predicate
ever widens to reach it, the test fails and the docstring's NOT-REACHED clause must be rewritten rather
than quietly becoming false. Same idiom as `derived_count_guard` printing its NOT-REACHED clause on
every run, green included: **a fence that hides what it does not reach is worse than one that reaches
less.**

## `D-M257x-176-6` — both arms RED-proofed before the suite, and both mutation controls are about a specific history

Arms proven to fire, in-process, without editing the tree:

| perturbation | result |
|---|---|
| drop `repair_postcondition_baseline.json` from `DECISIONS` | RED, naming that path |
| add a `DECISIONS` key that holds no such set | RED, naming it as gone |

And the two mutation controls are pinned to things that actually happened, not to plausible faults:

- **the JSON arm is load-bearing** — the registry that went four iters unnamed *is* a JSON object, and a
  python-only walk cannot see it. Deleting the arm must LOSE a site.
- **the ≥2 floor is load-bearing** — at ≥1 the predicate becomes iter-175's rejected 39-site instrument
  arriving by the back door. Raising the floor to 3 must lose a site, or the number 2 is decoration.
