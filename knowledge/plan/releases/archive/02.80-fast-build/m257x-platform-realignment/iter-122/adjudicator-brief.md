# TIER-1 ADJUDICATOR BRIEF — iter-122 claim census

You are adjudicating one batch of the M257x claim census. **You are one of twelve adjudicators working
in parallel; you do not see the others and must not look for their output.** Your verdicts are committed
**verbatim, before anything is aggregated** — do not soften, round, or summarise them.

## The one question

For each item in your batch:

> **Does the CITED CONTENT support the proposition the CLAIMING UNIT asserts?**

Nothing else. Not "is the prose well written", not "is the claim true in the world", not "should the
corpus say this". **Only: does the thing it points at back up what it says.**

## Verdicts — use exactly one of these tokens

| token | meaning |
|---|---|
| `SUPPORTS` | the cited lines do back the proposition the unit asserts about them |
| `DOES-NOT-SUPPORT` | the cited lines are about something else, contradict the unit, or are off by enough that a reader following the citation would not find the claim |
| `PARTIAL` | the citation lands on the right construct but the unit overstates or understates what is there (a quantifier the source does not carry, a scope the source does not cover) |
| `UNRESOLVABLE` | the pin is sha-qualified to a commit this tree is not at, or the cited file has moved such that the question cannot be decided here |
| `NOT-A-CLAIM` | the citation is decorative — the unit is a heading, an index row, or a pointer that asserts nothing about the cited content |

## Rules that are not negotiable

1. **The ±3 lines of context may not be enough. When they are not, GO AND READ THE FILE.** The
   `resolved file` path is absolute; open it and read as much as you need. A verdict guessed from the
   excerpt when the excerpt was insufficient is worse than no verdict.
2. **A citation that stops one line short of its own subject is `DOES-NOT-SUPPORT` or `PARTIAL`, not
   `SUPPORTS`.** This exact class cost M257x its tenth quantifier defect: a repair cited `:230-231` for a
   claim whose subject was the middleware on `:232`.
3. **An absolute quantifier in the unit — "every", "only", "no", "all", "never", "always" — is checked
   against the source.** If the source shows one instance and the unit says "only", that is `PARTIAL` at
   best. Nine of the ten quantifier defects this milestone found were of exactly this shape.
4. **Do not resolve ambiguity in the corpus's favour.** If you cannot tell, say `DOES-NOT-SUPPORT` with a
   reason that says you could not tell. Under-crediting is recoverable; over-crediting is the failure
   this census exists to end.
5. **Read-only.** You edit nothing outside your own verdict file. No corpus edits, no platform edits, no
   git operations.

## Context you need

- The corpus tree root is `/Users/marco/workspace/anthropos/rosetta`.
- The platform / service clones are under `/Users/marco/workspace/anthropos/rosetta/stack-demo/` at
  platform `0c91421`. **They are read-only to you.**
- The corpus frequently cites a source **@ a named sha** (e.g. "``@ platform `2adcf71``"). If the unit
  names a sha and the tree is not at it, that is `UNRESOLVABLE`, **not** `DOES-NOT-SUPPORT` — say which
  sha in the reason.

## Output

Write a TSV to the path you are given, one row per item, **no header**, columns:

```
id <TAB> verdict <TAB> one-line reason (no tabs, no newlines)
```

Every item in your batch gets exactly one row. **A batch that returns fewer rows than it was given items
is a lost seat and will be recorded as one** — count your rows before you finish.

Then report, in your final message: your batch number, the row count, and the tally per verdict token.
