# iter-191 — decisions

## `D-M257x-191-1` — the routed member was re-surveyed and the finding moved one line

`SURVEY-M257x-iter188-the-other-walks-are-unmeasured` named `story_org_count_guard._EXCLUDED_DIRS` as its
third member. Re-surveyed, **the prune list is fine**: component-matched, reasoned, and carrying the story
of the bug it memorialises. Taking it at face value would have produced a cosmetic iter.

One line down, `main()` computes the number it publishes — and the number its **refusal gate** keys on —
from `rglob("*.md")`, while `find_violations` walks `_SCANNED_SUFFIXES` (`.md .sh .yaml .yml`). Measured:
**119 printed, 164 scanned — 72.6 %.** The route's *subject* was right and its *target* was one
expression away, which is why re-survey is a phase and not a formality.

## `D-M257x-191-2` — the unit error runs BOTH ways, and the second direction is the dangerous one

The published count under-states the work. The same expression also gates the refusal: a scope with no
markdown but 45 shell/YAML sites would exit 2 with *"0 markdown file(s) in scope. Nothing was checked;
this is not GREEN."* — a **false CANNOT-RUN** over a scope with 45 files to check. A guard that refuses
when it could have answered is as wrong as one that answers when it could not, and it is harder to
notice because refusing looks conservative.

**And the non-markdown half is not empty of the claim class** — measured before repairing: **25 `.sh` +
99 `.yaml`** lines match the guard's own context pattern. So the markdown-only denominator was not a
harmless label on an all-markdown population.

## `D-M257x-191-3` — the cardinality arm was re-based from a SPELLING to the derivation

`test_the_real_tree_is_still_GREEN_and_states_its_cardinality` asserted `all \d+ scanned doc\(s\) agree`
— a **spelling**, so it passed while the number described 72.6 % of the population and would pass again
after any future narrowing (`§5` r70/71). It now parses the printed cardinality and asserts it **equals**
`sum(in_scope(scan_roots(root)).values())`.

The unit word moved with it: *"doc(s)"* → *"file(s)"*, because `.sh` and `.yaml` are not docs and the
old word is what made the markdown-only count read as correct.

## `D-M257x-191-4` — the absolute-path residual is the OTHER half of a bug this module already tells

`_EXCLUDED_DIRS`' comment memorialises a real RED: the list held `.agentspace`, matched as a **substring
of the absolute path**, and silently skipped both of rext's own sites. The fix addressed the **matching
mode** — *"match components; never substrings"* — and left the **scope**: `_excluded` still tested the
absolute path, so a checkout under any directory named `knowledge` or `node_modules` would exclude
everything beneath it. Both halves were in the same sentence of the original bug report.

Sized before the change (`§5`): **0 difference** on this box — no scan root's absolute path carries an
excluded component — and `main()`'s refusal gate fails closed, so the residual was latent. `root=None`
keeps the unscoped reading for any caller without a root; every call site inside the guard now passes one.

## `D-M257x-191-5` — both escalation conditions checked, with the measurement

`overview.md` pre-registered: **(a)** *if the true denominator changes the guard's VERDICT, this is a
correctness defect rather than a reporting one* — it does not: `OK`, before and after, with `4`/`1` orgs
and 0 violations, the number moving 119 → 164. **(b)** *if scoping `_excluded` changes what is excluded
anywhere, the absolute form was load-bearing* — it does not: `excluded_reach` is **0** under either
reading. So this iter is a **reporting** repair with a latent scope repair attached, and it says so.

## `D-M257x-191-6` — the prune list's reach is printed, and the honest value is zero

Per iter-188's rule, `excluded_reach()` is derived and printed: `0 excluded by 4 pruned dir name(s)`. A
zero — and it is published rather than suppressed, because *"4 names excluded nothing here"* is the fact
a reader needs to size the list. The `§9` obligation is discharged by an arm proving the derivation
returns non-zero over the tooling repo (which has a `node_modules`), not by the guard's own scope.
