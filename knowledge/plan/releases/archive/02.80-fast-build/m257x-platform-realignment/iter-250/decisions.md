# iter-250 decisions

## `D-M257x-250-1` — a cited path the tree calls RUNTIME state is a third bucket, decided by `git check-ignore`

`rext_path_guard` asked *"does this path exist in the rext tree"*. Existence is a property of the **box**:
`demo-stack/stacks/registry.json` is matched by `demo-stack/.gitignore:8` (`stacks/`), so it is present on a
box that has run a demo and absent on a clean clone of the same sha. Measured: the guard returned `rc=0`
here and `rc=1` on iter-249's frozen clone, over identical corpus text.

**Decision.** Cited paths are partitioned by `git check-ignore --stdin` against the rext tree — which
decides *pathnames*, existing or not, so the answer comes from the tree:

- **ignored** → RUNTIME. Never a finding, never a silent pass: printed by name on every run, empty or not.
- **not ignored, exists** → resolves.
- **not ignored, absent** → a finding, exactly as before.

**Rejected alternative:** add it to `DISCLOSED_ABSENT`. That list is for citations *whose own sentence says
the file does not exist* — a different claim, and it would have pardoned one instance while leaving the
class. `git check-ignore` closes the class.

**Fail-closed:** when git cannot answer (no `.git`, not a repo), the guard prints `RUNTIME BUCKET
UNDECIDABLE` and states that grading fell back to existence. `D-M257x-248-3`: a guard needing a reference
DECLARES the need. Tested.

## `D-M257x-250-2` — a fence over runnable inputs states WHO can run them

`fence_command_guard.locate` resolves against the repo root and every clone root, so a directory checked
into `rosetta` and one that appears only after `make init` produced the same GREEN. Measured on the live
corpus: **0 of 103 graded `cd` occurrences are reachable from a bare checkout; 102 need a provisioned
`stack-*/` workspace.**

**Decision.** `check()` returns a `reach` partition (`repo` / `tooling` / `workspace` / `unresolved`) that
sums to the graded `cd` count, and `main()` prints it on every run, before the verdict line (so
`guard_family`'s `lines[-1]` contract is untouched). This is a **disclosure, not a gate** — a corpus about
operating a stack is *expected* to be workspace-heavy, and turning the ratio into a threshold would be
inventing a requirement nobody has stated. What was wrong was that the number did not exist.

## `D-M257x-250-3` — the split is derived through the guard's own resolver, never a second parser

The reach census spies on `fence_command_guard.locate` during a real `check()` run rather than re-walking
the fences. iter-209's hand-written slugger was **16× wrong in one direction**; a second parser of the same
substrate is the defect this milestone keeps re-finding. The permanent form of the same rule is that the
tiering lives *inside* `check()`, so it cannot drift from what the guard actually graded — asserted by
`test_the_three_tiers_are_told_apart`, which requires the tiers to sum to the denominator.

## Note — PR-5 was refuted, and the refutation is structural, not incidental

The tier-W count is **25 live, 0 on a fresh clone**: with no `stack-*/` directory the guard refuses those
44 lines as *"workspace not provisioned on this host"* instead of tiering them. A split computed by
resolution cannot survive the loss of the thing being resolved against. Measuring the reach of a corpus
**as a reader would meet it** needs a text-shaped classifier, routed as
`ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace`.
