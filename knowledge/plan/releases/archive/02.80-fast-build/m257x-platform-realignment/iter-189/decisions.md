# iter-189 — decisions

## `D-M257x-189-1` — the divergence is LATENT, and it was sized twice before anything was repaired

Two independent sizings, both taken before the edit, because *"the same problem exists elsewhere" is only
a mood until it has a number* (`§5`):

1. **Structural:** directories in any local tree whose name *ends* in `vendor` / `node_modules` without
   being exactly that — **0**. So no path on this box is currently classified differently by the two
   rules.
2. **Behavioural, and this is the comparison nothing had ever run for this pair:** both readers were run
   against the real `app` clone at `origin/main@ad9f3c49`. **Both returned `{}`.** They agree today.

The finding is therefore a **latent** silent divergence, not a live corpus error — and it is recorded that
way rather than dressed up. What makes it worth an iter anyway is the direction of the error: on this
guard's *consumer* side, over-exclusion reports **a read that exists as absent**, which is the failure
`app_rpc_reads`'s own docstring is most careful about (*"None is not zero"*).

## `D-M257x-189-2` — the agreement is a ZERO reading, so it gets an instrument (`§9`)

`{} == {}` is exactly what a broken comparison also returns. The differential arm therefore runs against
a **synthetic git tree** carrying 8 paths — 5 that must be seen, 3 that must be excluded — including the
three the old substring rule silently dropped (`cloud-vendor/`, `my-node_modules/`, `vendoring/`).

And the arm is deliberately **two** assertions, not one: *the readers agree* is satisfied by a predicate
that excludes everything, which the mutation run demonstrates (M5 leaves the agreement arm green while
the expected-set arm goes RED). A parity check without an expected set is a parity check that passes on
two identical zeros.

## `D-M257x-189-3` — one predicate, not two reconciled rules

Rejected: keeping both rules and asserting they agree. That fences the symptom and leaves the cause —
two literals expressing one exclusion, which is iter-177's shape and the thing iter-188 had just removed
from `_SKIP_DIRS` one file away. `is_vendored_path()` + `VENDORED_PATH_COMPONENTS` are the single source;
both readers call it; source arms assert that they still do **and** that neither has grown a substring
test again.

The rule they were violating was **already written down in this repo**, by a sibling guard:
`story_org_count_guard.py:125` — *"an exclusion that can swallow the whole repo. **Match components;
never substrings.**"* A rule stated in one module and broken in another is not a documentation gap; it is
an unfenced rule.

## `D-M257x-189-4` — a third finding landed in the same edit, and it is named separately

`_reads_worktree` applied its component test to the **absolute** path (`set(path.parts)` over
`/Users/…/stack-demo/app/…`). That makes the exclusion depend on **where the clone happens to live**: a
checkout under any directory named `vendor` or `node_modules` excludes *every* file and reports a
confident zero for the entire consumer side — the whole-population silent zero this milestone keeps
finding, in the guard that most loudly refuses one.

Sized: **0** such ancestors on this box. Fixed to repo-relative, which is also what makes the two readers
genuinely comparable (`_reads_at_ref` reads repo-relative paths from `git grep`). Fenced by its own arm
and mutation-proven (M2). It is called out separately rather than folded into `D-M257x-189-1` because it
is a **different** defect that happened to be in the same three lines.

## `D-M257x-189-5` — escalation conditions checked; neither fired

`overview.md` pre-registered two. **(a)** *If the two readers disagree today on the real `app` clone, this
is a live defect in a published corpus claim* — measured, they do not (both `{}`). **(b)** *If the
substring form is deliberate, keep it and fence the difference* — it is not: `_reads_at_ref`'s own
docstring calls itself *"same derivation"* as the worktree reader, and no comment anywhere claims a path
shape the component match would miss.

Post-repair the guard's live verdict is unchanged: `G6 8 RPC var(s) graded {'unconfigured': 8}, 0
mid-fold; app consumer side measured @ origin/main@ad9f3c4` → `OK`.

## `D-M257x-189-6` — the route fence fired on this iter's own close, for the defect iter-184 built it to catch

The first draft of the routes block wrote `` `SURVEY-M257x-iter186-…` `` — an **elided** id, used as
shorthand for a route named in full two lines earlier. `route_disposition_guard` went RED with exactly
the right reading:

> `'SURVEY-M257x-iter186-' is not a route id — carried at iter-189` … *a carry-forward that names a SET
> must ENUMERATE it (`§5` rule 73): a glob leaves a truncated stem behind and a hard-wrapped id leaves
> its head, and both read as live backlog in every brief that quotes this queue.*

That is iter-184's founding defect (`SURVEY-M257x-iter181-`, the stem left by iter-182's glob), reproduced
by a different mechanism — an ellipsis rather than a wildcard — and caught. **Repaired by writing the id
in full**, never by relaxing the grammar: the whole value of this fence is that an abbreviation in a
close block is indistinguishable, downstream, from a route that exists.

Worth recording because it is now the **third** time an instrument built in this milestone has fired on
its own author (h42, h44, here). That is the fence working, and it is also the honest answer to *"do these
fences catch anything?"* — the first thing they catch is us.
