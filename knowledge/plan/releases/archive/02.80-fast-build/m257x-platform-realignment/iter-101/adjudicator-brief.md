# iter-101 adjudicator brief — held fixed at the iter-95/97/99 shape

Four adjudicators, one per seat-group, each **re-deriving from the platform clones** rather than from any
seat's evidence or any prior verdict. This is the half of the instrument that turns booked findings into a
graded number.

## Your job

For every BLOCKER booked by your assigned seats, return **UPHELD** or **REJECTED**, plus scope.

- **UPHELD** — the claim is false or unsupportable against ground truth, or contradicts another corpus
  claim, at the ref the claim itself names.
- **REJECTED** — the booking does not survive re-derivation.

## Re-derive. Do not adjudicate from the seat's evidence.

The seat's citation is a **pointer**, not proof. Open the platform file yourself, at the right ref, and read
around the cited line. A booking that quotes a real line and draws the wrong conclusion is the most common
rejection, and you cannot see it without opening the file.

## The rules that decide the close calls

1. **A claim is settled at the ref the claim itself NAMES** (`platform-alignment.md` §5 rule 33). A pin is a
   **date**, not an excuse: if the claim is true at its named ref it is **TRUE**, however stale. A pin's
   scope is the claim's own block — a table cell, a wrapped sentence. A ref in a neighbouring row does not
   date this row's claim.

2. **THE REF-DISCIPLINE REJECTION CLASS.** This class has now run **17 occurrences across five readings and
   contributed ZERO to any graded count.** It is: a seat books a pinned, past-tense, or dated claim because
   newer evidence contradicts it. That is not a defect — it is the pin working. Expect it; reject it; name
   it as ref-discipline so it stays filtered.

3. **Three instruments, and no single one is safe.** Before upholding any *absence* claim ("returns 0",
   "occurs nowhere", "is read by nothing"), check all three mechanisms:
   - `.gitignore` hides **tracked** files from this shell's `grep` (it is `ugrep --ignore-files`);
   - **NUL-bearing source** is skipped by BOTH `grep -I` and `git grep` (2 such files, **one** NUL byte
     each — count bytes with `tr -dc '\000' < FILE | wc -c`, never `grep -c`);
   - **nested untracked repos** (`stack-demo/app/studio`, `stack-demo/cms/studio`, each own checkout at
     `aeec036`) are invisible to `git grep` at the HOST ref.
   Measured on one predicate the three returned **1 / 0 / 22**, and the **0 was the ref-named `git grep`**.

4. **Verify the PREDICATE, not the arithmetic.** A count can be exactly right while the claim it supports is
   false. Re-derive the SET the arithmetic ranges over, independently, and state its cardinality first.
   *"I re-derived it and it matches"* is the weakest clearance there is.

5. **Self-contradiction is a real finding** even when you cannot tell which side is right — if two passages
   assert incompatible things, that is upheld, and you cite both anchors.

6. **Scope each upheld blocker** as IN-SCOPE (inside `corpus/services/**` or `corpus/architecture/**`) or
   OUT-OF-SCOPE. Only in-scope upheld BLOCKERS enter `N`.

## Two clone sets exist for `rosetta-extensions`

- `stack-demo/rosetta-extensions` — the **pinned per-stack consumption clone**, `ab81527a`.
- `.agentspace/rosetta-extensions` — the **authoring copy**, `09d0607` on `main`.

They are different trees at different refs. Decide which one settles the claim in front of you, on the
claim's own terms, and **say which tree you read**.

## Deduplicate across seats — but only where the PREDICATE is the same

Two seats booking the same underlying false predicate at different anchors is **one** predicate and **two**
anchors. Report both, and say which anchors collapse onto which predicate. Do not collapse two anchors that
merely look similar.

## Output

Write your verdict file to the path you are given. For each booking:

```
<seat> <Bn> | <corpus anchor> | UPHELD|REJECTED | IN-SCOPE|OUT-OF-SCOPE | <predicate in <=20 words>
   evidence: <the repo/file:line YOU opened, and what it says>
   [if REJECTED] class: ref-discipline | wrong-tree | mis-read | already-true | other — <one line>
```

`wrong-tree` means the booking was graded against a different checkout of the same repo than the one that
settles it (including the two `rosetta-extensions` clones and the two nested `studio` checkouts). It is a
**label on a rejection you already decided**, not a reason to reject.

Then a summary line: `BOOKED=<n> UPHELD=<n> REJECTED=<n> IN-SCOPE-UPHELD-BLOCKERS=<n>`.

## Hard bars

- **Do not read `knowledge/plan/**` except this brief, your assigned seat reports, and your own output.**
  Prior audits' ledgers are answer keys; reading them turns adjudication into agreement.
- Read-only. Zero edits outside your one output file. No git state changes.
- `export DEVELOPER_DIR=/Library/Developer/CommandLineTools`; `cd /Users/marco/workspace/anthropos/rosetta` first.
