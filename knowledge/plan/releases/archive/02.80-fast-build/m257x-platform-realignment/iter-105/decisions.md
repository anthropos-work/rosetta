# iter-105 — decisions

## `D-M257x-105-1` — the fence tree is read from where PYTHON LOADED THE MODULE, not from a flag or a cwd

`fence_tree()` defaults its root to `Path(__file__).resolve().parent`. Three alternatives were available and
each reintroduces the defect in a different way:

| alternative | why it fails |
|---|---|
| a `--fence-tree` flag | the one input that decides the verdict becomes the one input an operator can point away from — which is precisely what happened at iter-103, by accident rather than by flag |
| `os.getcwd()` | the family is run with `cwd=repo_root` (the **corpus**), so it would have reported the corpus tree twice and the fence tree never |
| `$ROSETTA_EXTENSIONS_ROOT` or similar | an env var that is usually unset and occasionally wrong; the same class as `.agentspace/rext.tag` being git-ignored — a source of truth that never appears in a diff |

A guard's waivers, baselines and assertion sets are its **siblings on disk**. The tree that settles its
verdict is therefore the tree it was loaded from, and `__file__` is the only witness that cannot disagree
with what Python actually imported.

## `D-M257x-105-2` — DIRTY is disclosed; UNDETERMINABLE is refused

Two different conditions, deliberately graded differently.

**Undeterminable** (no git, no HEAD) → `EXIT 2 — UNMEASURED`, with `--allow-unknown-provenance` to accept
and **record** the gap. This is the family's own existing doctrine — it already exits 2 for a platform clone
with no `origin/main` — applied to itself, and §8's *a guard has three verdicts, not two* rule.

**Dirty** → stated in the reference line **and repeated on the summary line**, and the run proceeds. Two
reasons, and the second is the one that matters:

1. Refusing a dirty tree would make the family unrunnable during exactly the iters that ship fences — the
   iters that most need to run it.
2. The summary line is the line that gets quoted forward. Harden pass-20's finding was that the family
   printed *"OK — every member … returned green"* directly beneath *"N guard(s) NOT RUN and accepted"*, and
   the OK line is the one that travelled. A caveat that lives only in the header has the same fate.

**Live demonstration, unplanned:** every family run inside this iter printed
`fence tree 944fc4a21 is DIRTY — the verdict was taken with uncommitted configuration`, because the fence's
own edits were uncommitted. The fence told the truth about itself on its first run.

## `D-M257x-105-3` — the stamp prints FIRST, and the reason is a defect this milestone has already paid for

`guard_family.run_one()` reports `lines[-1]` for a green member. A stamp printed last would have silently
replaced **every guard's own summary line** in the family view — the identical shape as iter-87's `headline()`
finding, where the family reported whichever assertion happened to sort last and a membership departure went
invisible in the one view that claims to summarise the family.

Printing first also keeps it out of `headline()`'s finding-shaped-line cut (indented, or opening `[`/`-`/`*`)
and out of `_STATED_COUNT`, so it cannot inflate a RED cardinality either. Both were checked, not assumed.

## `D-M257x-105-4` — the conformance check is DERIVED from `guard_family.census()`, and asserted over the AST

The check that every fence stamps is not a list of seventeen filenames. It calls `F.census(HERE)` — the same
derived census the family runner is built on — so a new `*_guard.py` that does not state its tree turns the
test RED **without anyone remembering to add it**. §2's hand-maintained-tuple lesson, applied to the check
rather than to the thing checked.

It parses the module and walks the `if __name__ == "__main__"` block for a call to `fence_provenance.stamp`,
per §8's *a fence over source must assert against a parsed construct, never a whole-file substring*. The
three mutation controls exist because a substring check passes all three: a module that **mentions** the
name in a comment, one that **imports** it and never calls it, and one that stamps at **import time**
(which does not travel with a standalone run and also fires during test collection).

**Anti-vacuity, per §8's iter-94 rule — written against the SUBJECT, not the inputs.** The control does not
merely assert the discovered set is non-empty (one stray file satisfies that). It asserts the discovered set
**is exactly `guard_family.census()`**, so the check and the family cannot drift into disagreeing about what
"the family" is — which is how a fence ends up green over a universe it never examined.

## `D-M257x-105-5` — the 52 prior verdicts are re-graded to *provenance-unstated*, not to *void*

Measured: `grep -rnoE "[0-9]+ GREEN · [0-9]+ RED"` over the milestone dir returns **52 recorded family
verdicts across 26 artifacts**, and **0** state the fence tree. (One line mentions `rext` — it names the
module path `rext stack-core/guard_family`, not a sha.)

None of them is thereby **wrong**. Almost certainly most were taken from the authoring copy, which is where
the work happens. But *"almost certainly"* is the word this milestone exists to remove: they are
**unre-checkable**, which is a strictly weaker thing than a green, and the honest grade is
**provenance-unstated until re-run**.

Deliberately NOT done: retroactively annotating 26 artifacts. That would be 26 edits asserting a provenance
nobody measured — inventing evidence to fix a lack of evidence, in the milestone whose own class is claims
that outrun them. The re-grade is stated once, in §5 rule 50, as a **reading instruction** for anyone
quoting a pre-iter-105 verdict forward.
