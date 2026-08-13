# iter-242 — decisions

## `D-M257x-242-1` — the substitution NOTE is non-fatal

Refusing a `--platform` whose `.resolve()` moves the clone set would be the stricter call, and it is the
wrong one: a symlinked workspace is a legitimate setup (containers, worktrees, shared caches all produce
them), and a hard error would break runs that are asking the right question.

**The defect was never the resolution — it was that the resolution was SILENT while changing the
subject.** So the runner names both clone sets, says a symlinked alternative *does not take effect*, and
continues. `§5` — a verdict states the tree it was taken with.

## `D-M257x-242-2` — `service_registry_guard`'s CANNOT-RUN on a restricted set is CORRECT

It flips GREEN → rc=2 on the fresh clone set because it cannot read the compose file it grades. That is
not a bug to fix: it is the fail-closed contract working, and `guard_family` reports it as
`could-not-check` rather than folding it into the green count.

Recorded because the reflex on seeing a member "break" under a new fixture is to make it pass again, and
making it pass would mean making it grade nothing.
