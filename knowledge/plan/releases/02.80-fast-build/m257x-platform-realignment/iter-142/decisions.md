# iter-142 — decisions

## `D-M257x-142-1` — audit the predicate BEFORE the repair, not after the publication

iter-138 ran a mechanical predicate corpus-wide, **published 127 rotted of 222 decidable (57.2 %)**, and
iter-139 audited it on a pre-registered stratified sample to **0 of 12 (Wilson95 [0.0, 24.3])**. The
figure was retracted at three sites.

iter-142 ran a comparably mechanical predicate over a comparable population and inverted the order: the
census was taken, **every one of its findings was read in context**, the predicate was corrected on what
that reading showed, and only then was a line of prose touched. Result: **44 of 44 true positives — the
whole population, not a sample.**

**The predicates were not obviously different in quality. The ordering was.** Booked here because the
milestone has now run this experiment in both directions within eight iters, which is as close to a
controlled comparison as this work gets.

## `D-M257x-142-2` — a ref-qualified historical pin is IN the class

*"it was `:100` at `0dab54d`"* is a **true** statement: `0dab54d` is immutable, and at that ref the
construct genuinely was on line 100. The temptation is to exempt it — a true statement about a frozen
ref cannot rot.

**Rejected, and the reason is the mechanism, not the semantics.** The thing that goes RED is a resolver,
and **a resolver does not read the qualification.** It sees `` `:100` `` and binds it against HEAD. That
is exactly rule 63(c′)'s *"a fence matching on FORM cannot tell the quotation from the assertion, and it
is right not to"*, stated from the other side: the hazard is the **token**, independent of the truth of
the sentence wrapped around it.

Consequence for the fence: no ref-qualification exemption, and the census counts these with the rest.

## `D-M257x-142-3` — fence the TOKEN, not the digit

The obvious objection to (c′) is that retractions carry information — *"iter-102 re-derived them by
adding +23 and +16 to the old numbers instead of re-measuring"* is a finding, and it needs numbers.

**It does not need pin tokens.** The class is `` `:NN` `` written in the corpus's citation form; prose
numbers (*"+23 and +16"*, *"ten lines earlier"*, *"rotted +8"*) are outside it and always were. So the
repair is **never** a choice between evidence and hygiene. Every one of the 44 sites kept whatever the
retracted pin was making a point with.

This also fixes the fence's scope: it must match the **backticked digits-only token**, never any number
near a retraction verb. A guard that matched digits would have forced exactly the evidence-destroying
repair the objection fears.

## `D-M257x-142-4` — a sweep against anchor rot must be line-count flat, and that is checkable

iter-141 repaired three sites of this class **and turned a fence RED doing it**, because its own
insertion moved a pin below it. A sweep whose purpose is removing anchor rot must not induce anchor rot.

Every one of this iter's 17 files was rewritten **in place**: 36 lines changed, **added minus removed = 0
for every file**, checked with `git diff --numstat`. It is a one-command post-condition and it belongs on
any repair whose subject is line numbers.

## `D-M257x-142-5` — a duplicate test run was killed, and is disclosed rather than dropped

The scoped suite was launched twice: once piped to `tail` (which produced no readable output), then again
to a file. Both were alive together for roughly a minute. The first was **killed** so two concurrent runs
could not confound each other's temporary state — the same reasoning iter-111 and iter-121 applied to
their confounded runs, applied preventively instead of retrospectively.

**The reported 209 passed / 0 failed is the file-based run taken alone.** Recorded because this milestone
does not let an unreported process share a measurement window.

## `D-M257x-142-6` — a fence's regex IS its denominator, and only another axis can see the omission

The first pass of this iter matched **bare** pins only, enumerated **922**, repaired 44, and reported the
class **clean at 0**. Every number was true of the population the regex matched. **None of them was true
of the class**, which also contains path-qualified pins (`` `main.go:507-508` ``) that rot identically.

**Nothing inside the fence could have caught this.** Its mutation control fired correctly (on bare
pins). Its anti-vacuity arm passed correctly (the enumeration was large and non-empty). Both controls
answer *"is this guard working?"* and the guard **was** working — on the wrong denominator. A census's
controls validate its mechanism; they cannot validate its scope.

What caught it: **`repair_leak_guard`, run over this iter's own commit**, went RED on
`shared_libraries.md:77` — the twin of a `CLAUDE.md` site the new fence had just flagged and repaired,
differing only in that its pin carried a path. A guard on a **different axis** (does this commit leave a
twin of what it just rewrote?) saw what a guard on the same axis structurally could not.

Two rules follow, and they are cheap:

1. **A reach claim must name the FORMS the predicate matches, not just the count it returns.** *"922
   pins, 0 residual"* is not falsifiable by a reader; *"922 pins **of the form** `` `:NN` ``"* is — and
   invites exactly the question that found this.
2. **Run the commit-scoped guards over a sweep's own commit, always.** They are reported `not-run` by
   the family (`anchor_offset`, `repair_leak`, `value_change`, `repair_reach`) and it is tempting to
   leave them there. On this iter they returned **3 RED, all real**: a rotted cross-doc pin, the
   path-qualified twin, and a value-change twin. They were the gate.

This is the same shape as `D-M257x-140-2` (*a class is censusable iff its subject carries its own head*)
one level up: **a class is censused iff the predicate's population and the class's population are the
same set — and that equality is an assumption until something outside the predicate tests it.**


## `D-M257x-142-7` — an article is not a value; disclose a false RED, never obey it

`value_change_guard` went RED claiming `storage.md:29` still published a value this commit corrected in
`skillpath.md`. The "value": **`the` → `that`**. The collision: both docs carry the house idiom *"a pin
is a pin, a branch name is not"*, and this iter's rewording changed the article that follows it in one
of them.

Measured before acting, which is the point:

* `rg "CLAUDE\.md (line|anchor) (was|stood)"` returns **exactly two** sites, `backend.md` and
  `skillpath.md`, and **both already carry the repaired form**. There is no third twin.
* `storage.md:29` makes a different point entirely — that a **sha is a pin and a branch label is not**,
  about `origin/main` — and carries no CLAUDE.md-line claim to correct.

**The tempting fix was to edit `storage.md` so the guard would go quiet. That would have been
vandalism**: changing correct prose to satisfy an instrument is the failure mode this whole fence family
exists to prevent, one level up. The actual fix was at source — keep the article in this iter's own new
wording, which says the same thing and removes the collision.

Routed forward as `FIX-M257x-iter142-value-change-articles`: a change confined to **function words**
should not register as a value change. Until then the behaviour is disclosed rather than assumed benign.
