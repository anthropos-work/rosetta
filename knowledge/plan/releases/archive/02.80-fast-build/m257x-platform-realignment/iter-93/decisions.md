# iter-93 — decisions

## `D-M257x-93-1` — fence the HEDGE, tree-wide; hand-repair measurably did not hold

iter-92 repaired six restatements by hand and **leaked twice doing it** (`repair_leak_guard` RED on the iter
commit, RED again on the first repair, GREEN on the second). That is the measured case for TOK-05 read
literally: the unit is the predicate, and the predicate here is *"a claim about a repo in no clone set must
say it is not a measurement."*

`claim_twin_guard` already fences **adjudicated claims** across the tree. Nothing fenced a **hedge** — and a
hedge that lives only in the document a guard reads is worse than no hedge, because its existence implies
the system checked.

## `D-M257x-93-2` — the marker is a SET OF PHRASES, not a mandated token

Deliberate. A fence that requires a magic string teaches authors to type the string, and the corpus would
acquire a ritual phrase that means nothing. The guard accepts any of the phrasings the corpus already uses
naturally (*"not visible to this corpus"*, *"was not measured"*, *"never been in the clone set"*, *"report
both, assert neither"*, …). The property being fenced is **that the reader is told**, not that a token
appears.

## `D-M257x-93-3` — the scope is the PARAGRAPH, and that choice is load-bearing in both directions

A marker anywhere in the file would launder a flat assertion three screens away — a reader lands on a
paragraph, not on a file. So the window is the blank-line-delimited paragraph, with **a wrapped markdown
blockquote counted as one paragraph**, because that is how every one of these claims is actually written; a
fixed ±1-line window would have produced false REDs on the real corpus. Both directions are tested:
`test_the_marker_must_be_in_the_SAME_paragraph_not_merely_in_the_file` and
`test_a_wrapped_blockquote_counts_as_ONE_paragraph`.

## `D-M257x-93-4` — the guard RE-MEASURES ITS OWN PREMISE AND RETIRES ITSELF

The assertion rests on *"we cannot read `infrastructure`"* — a fact about the workspace, not a law. So it is
measured on every run: if an `infrastructure` clone ever appears beside the platform clone, the guard stops
demanding the hedge and prints **PREMISE LIFTED — go and MEASURE those declarations, then retire this
fence.**

This is §8 rule 3 (*pin the mechanism, never the contents*) turned on a fence's own premise. A guard that
kept enforcing a hedge after the hedge became unnecessary would be **pinning the current shape of our
ignorance**, and every future correct change would have to argue with it. Tested in both directions, since
the lifted-premise test alone would pass for the wrong reason.

## `D-M257x-93-5` — the anti-vacuity rung, and the skip that proved it necessary

If the corpus stops mentioning `module.*_euwest1` at all, the guard has checked nothing and **exits 2**
rather than reporting an unearned green (§5 rule 8). That is the rung that stops this fence rotting silently
the day somebody renames the modules.

The need for it was demonstrated inside this iter rather than argued: the live-corpus control was first
written with a hardcoded `parents[3]`, which is `.agentspace/` in the authoring copy — so it **silently
SKIPPED**, and the suite reported `OK (skipped=1)`. It now walks up to whichever ancestor owns
`corpus/architecture`, which is also correct from a per-stack consumption clone at a different depth.
**A check that skips reads exactly like a check that passes — including when it is the check on the guard
that exists to say so.**

## `D-M257x-93-6` — the fence went RED on its own author's protocol note, within minutes, and that is a REAL limitation

Landing the protocol-doc write-up of this very class turned the new guard RED. The offending paragraph
quotes the wrong claim in order to explain that it is wrong:

> Six other documents … stated flatly that `module.cms_euwest1` **is still declared as the rollback path**.

**The guard cannot distinguish a USE of a claim from a MENTION of it.** That is a real limitation, not a
bug, and it is recorded rather than engineered around: adding a "but this one is a quote" escape would give
every future author a one-word way to opt out of the fence, which is precisely how a fence stops fencing.

The chosen fix is the one the fence is asking for: **the paragraph now says the unmeasurable part too**
(*"a thing **not visible to this corpus** at all, since `infrastructure` has never been in the clone set"*).
That is not a workaround — a paragraph that quotes an over-claim is a paragraph a reader can quote onward,
so it should carry the hedge as much as any other.

Third time this session a fence caught the person who built it (iter-92's `main.tf:39` citation,
iter-93's silently-skipping live control, and now this). Recorded as evidence for the practice, not as
apology: **the fences are being written by someone who needs them, and they are catching him.**
