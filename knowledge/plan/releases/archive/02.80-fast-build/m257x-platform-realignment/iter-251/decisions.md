# iter-251 decisions

## `D-M257x-251-1` — out-of-subject is decided by the tree, and the verb is DELEGATE, not exempt

`corpus_citation_guard` graded three clauses with `tp.exists()` (`:232`, `:255`, `:262`). This repo's
`.gitignore` excludes `.agentspace/`, so **21** citations of files in the rosetta-extensions clone were a
silent pass on a box that has the clone and **21 findings on a `rosetta`-only checkout** — the guard
reporting correct citations as broken.

**Decision.** Targets this repo ignores are partitioned out by `git check-ignore --stdin` — which decides
pathnames whether or not they exist, so the answer comes from the **tree** — into `census.delegated`,
printed by name on every run, empty or not, and not graded here.

**"Delegated", not "exempt", and the word carries the obligation.** These paths are not being forgiven;
they are being handed to `rext_path_guard`, whose subject is the rext tree and which can actually resolve
them. An exemption asks to be trusted. A delegation names the fence that took over — see the next entry.

**Fail-closed:** no git, not a repo → `census.delegable = False` and the guard prints
`DELEGATION UNDECIDABLE`, stating that grading fell back to existence. `D-M257x-248-3`.

## `D-M257x-251-2` — a delegation is legitimate only if another fence provably owns the target

Checked **before** the repair shipped, and now the live test's answer key: each of the 21 delegated paths
was looked up in `rext_path_guard`'s own collected reference set (`rpg.collect(root,
rpg.reference_re(rpg.sections(rext)))`). **21 delegated · 21 owned · 0 orphans.**

The test asserts the orphan list is empty and says why in its failure message: *a delegated citation that
NO fence owns is worse than the defect this delegation fixed.* That is the failure mode a bare exclusion
would have shipped silently.

## `D-M257x-251-3` — `enumerated` and `graded` are different numbers, and both are printed

The delegation happens **after** the population counter increments, so `C1` stays **1,606** and the header
still reads *"enumerated 1,808 citation(s)"*. The delegated count is printed as its own line beneath the
per-arm breakdown.

**Rejected alternative:** decrement the population so it reads "graded". It would have made the two words
agree at the cost of losing the enumeration — and this guard's anti-vacuity gate (`census.total == 0` →
exit 2) is defined over the enumeration. Two numbers, both stated, is the honest shape; one number that
silently means the other is the class this milestone keeps repairing.

## Note — the class is now three guards wide

`rext_path_guard` (iter-250, 1 instance) · `corpus_citation_guard` (iter-251, 21) ·
**`fence_command_guard.locate` and `clone_drift_guard`, unfixed.** `fence_command_guard` is the one that
matters most: iter-250 measured **102 of its 103 graded `cd` occurrences resident in a `stack-*/`
workspace**. Routed under the existing handler `FIX-M257x-250-existence-is-not-a-tree-property`.
