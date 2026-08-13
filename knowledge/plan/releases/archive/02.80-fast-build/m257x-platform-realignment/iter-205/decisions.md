# iter-205 — decisions

## `D-M257x-205-1` — the third site-kind, and one classifier with two callers

A `#` comment is a **token**, discarded before an AST exists, so neither existing census could ever have
seen one. `comment_measurement_literals` reads them with `tokenize`: **118 measurement-shaped numbers**,
of which **95 are `standing`** — a markedly higher standing share than either sibling (docstrings 73 of
164, prints 0 of 10), which fits what comments are for: this repo writes its *"Measured: N of M"*
provenance there.

**The classification rule is now `_classify_measurement`, one function, two callers.** Writing it twice
was the available shortcut and it is the shape iter-202 paid for — two reconstructions of one rule that
agreed until they did not, and disagreed **16 against 19** on their first joint run. Fenced by an arm
that asserts both censuses call it.

The three site-kinds **partition**: a fence stages the same sentence as a docstring, a comment and a
`print`, and requires each census to see exactly its own.

## `D-M257x-205-2` — the vocabulary was CASE-SENSITIVE, and that blinded all three censuses at once

The residual arm `noun_vocabulary_reach` gained at iter-204 came back with exactly one entry:
`playthroughs::1`. But `playthroughs?` **is** in the vocabulary. The occurrence is
`demo-stack/cockpit.py:229` — *"23 Playthroughs stayed green"* — and `_MEASURED_RE` was
case-**sensitive** while `_MEASURED_NOUNS` is written lower-case.

**Every capitalised measurement noun was invisible to all three censuses simultaneously**, and this repo
capitalises the nouns it treats as proper: Playthroughs, Stories, Heroes. Repaired with `re.IGNORECASE`
and regression-tested on the exact sentence.

**The find is the fence's, not a reader's.** Nobody would have caught this by re-reading the word list —
the word was *in* it. It took a residual arm that compares the vocabulary against what the tree actually
writes, which is what makes a reach audit different from a longer list.

## `D-M257x-205-3` — the reach audit had the blind spot it exists to find

`noun_vocabulary_reach` scanned **string literals only** for exactly one iter — the identical hole it was
built to expose in its consumers, in the function that exposes it. Extended to comments in the same pass:
the superset is now **799** number+word occurrences (up from 483), reach **38.4 %**, residual **zero**.

Recorded as its own decision because the shape recurs: *an instrument inherits the blind spot of whatever
it reads, including when the instrument's whole job is to find that blind spot elsewhere.*

## `D-M257x-205-4` — both ceilings taken from the census, for the third iter running

| ratchet | set at | from |
|---|---|---|
| `COMMENT_LITERAL_CEILING` | **118** | the census, after both of this iter's vocabulary changes |
| `DOCSTRING_LITERAL_CEILING` | **164** | the census, re-taken because the same changes moved it |

The first scan of the comment site-kind read **101 across 89 sites**. Quoting that would have set a
ceiling through the narrower lens — the third repeat in three iters of `D-M257x-203-2` (*a census's
denominator must come from the census*), after `D-M257x-204-4` caught the second. The pattern is now
explicit enough to state as a rule: **take the number after the last change, from the instrument, every
time — a ceiling measured mid-change is a ceiling for a tree that no longer exists.**

## `D-M257x-205-5` — five consecutive REDs on the registry's own author

Adding `comment_measurement_literals` turned `test_every_executable_derivation_is_classified` RED inside
a minute — the **fifth** time in four iters (`printed_arithmetic_totals` iter-193,
`printed_measurement_literals` iter-199, `docstring_measurement_literals` iter-203,
`noun_vocabulary_reach` iter-204, this one).

Five REDs on five different days, each on the table's own author, is the strongest evidence available
that `unclassified()` has **not** silently fallen behind the tree — which is the entire claim it was
built to make. Recorded as a result, not as a chore.
