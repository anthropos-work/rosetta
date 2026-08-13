# iter-119 adjudicator brief — held fixed at the iter-95/97/99/101/103/109/116 shape

Four adjudicators, one per seat-group, each **re-deriving from the platform clones** rather than from any
seat's evidence or any prior verdict. This is the half of the instrument that turns booked findings into a
graded number.

## Your job

For every BLOCKER booked by your assigned seats, return **UPHELD** or **REJECTED**, plus scope, **plus a
one-line statement of the underlying PREDICATE**.

- **UPHELD** — the claim is false or unsupportable against ground truth, or contradicts another corpus
  claim, at the ref the claim itself names.
- **REJECTED** — the booking does not survive re-derivation.

> ### The PREDICATE line is not optional, and here is why
>
> The primary metric of this reading is **`P`, the count of distinct false PREDICATES**, not the count of
> anchors. Two readings ago anchors rose 24 → 33 while predicates held 22 → 22 — *the same falsehoods, in
> more places*. **A verdict without a predicate line cannot be counted.** State the predicate as a
> proposition that is false, in ≤ 20 words, phrased so that two seats booking the same falsehood at
> different anchors yield the *same* sentence.

## Re-derive. Do not adjudicate from the seat's evidence.

The seat's citation is a **pointer**, not proof. Open the platform file yourself, at the right ref, and read
around the cited line. A booking that quotes a real line and draws the wrong conclusion is the most common
rejection, and you cannot see it without opening the file.

## GROUND TRUTH — the clone refs, read at this reading's open with NO fetch

| repo | path | checkout HEAD | `origin/main` |
|---|---|---|---|
| platform | `stack-demo/platform` | `0c91421d` | `0c91421d` — in sync |
| app | `stack-demo/app` | `ad9f3c49` | `ad9f3c49` — in sync |
| app/studio (nested, own checkout) | `stack-demo/app/studio` | `aeec036a` | — |
| cms/studio (nested, own checkout) | `stack-demo/cms/studio` | `aeec036a` | — |
| next-web-app | `stack-demo/next-web-app` | `8297c684` | `f97ba659` |
| sentinel | `stack-demo/sentinel` | `f2c46190` | in sync |
| studio-desk | `stack-demo/studio-desk` | `41ee3575` | in sync |
| ant-academy | `stack-demo/ant-academy` | `22df69dd` | in sync |
| cms | `stack-demo/cms` | `ca50c817` | `f38c0c4a` |
| jobsimulation | `stack-demo/jobsimulation` | `462343b0` | `82cb66ec` |
| messenger | `stack-demo/messenger` | `fa47850d` | `e9421c68` |
| storage | `stack-demo/storage` | `4ce8ece5` | `9f8cb532` |
| roadrunner | `stack-demo/roadrunner` | `87d8d443` | in sync |
| graphql-wundergraph | `stack-demo/graphql-wundergraph` | `60c229f3` | in sync |
| rosetta-extensions (**pinned per-stack**) | `stack-demo/rosetta-extensions` | `09d06070` | `4cb920aa` |
| rosetta-extensions (**authoring**) | `.agentspace/rosetta-extensions` | `43049308` on `main` | — |

**Every one of the 14 platform clones is at the same sha it was at the previous THREE readings.** That is
context, not licence: the previous two readings measured **~33 % and ~38 % platform-drift over a subject in which nothing
moved**. A claim can be false against a sha that never moved, and here that is not a hypothetical.

**Do not `git fetch` or `git pull` anything.** A reading's ground truth includes the clone refs
(§5 rule 41a); a fetch mid-reading moves the very refs the citation guards resolve against and makes the
number unprovable.

## The rules that decide the close calls

1. **A claim is settled at the ref the claim itself NAMES** (`platform-alignment.md` §5 rule 33). A pin is a
   **date**, not an excuse: if the claim is true at its named ref it is **TRUE**, however stale. A pin's
   scope is the claim's own block — a table cell, a wrapped sentence. A ref in a neighbouring row does not
   date this row's claim. **⚠️ And a ref EARLIER IN THE SAME PARAGRAPH does not date a later sentence that
   names its own ref** — a guard in our own tooling got exactly this wrong at the previous reading’s open, is still open, and
   produced a false RED. If a block carries two refs, work out which one dates *this* proposition.

2. **THE REF-DISCIPLINE REJECTION CLASS.** This class ran **17 occurrences across five readings and
   contributed ZERO to any graded count**. It is: a seat books a pinned, past-tense, or dated claim because
   newer evidence contradicts it. That is not a defect — it is the pin working. Expect it; reject it; name
   it as ref-discipline so it stays filtered.

3. **Three instruments, and no single one is safe.** Before upholding any *absence* claim ("returns 0",
   "occurs nowhere", "is read by nothing"), check all three mechanisms:
   - `.gitignore` hides **tracked** files from this shell's `grep` (it is `ugrep --ignore-files`);
   - **NUL-bearing source** is skipped by BOTH `grep -I` and `git grep` (2 such files, **one** NUL byte
     each — count bytes with `tr -dc '\000' < FILE | wc -c`, never `grep -c`);
   - **nested untracked repos** (`stack-demo/app/studio`, `stack-demo/cms/studio`, each own checkout at
     `aeec036a`) are invisible to `git grep` at the HOST ref.
   Measured on one predicate the three returned **1 / 0 / 22**, and the **0 was the ref-named `git grep`**.

4. **Verify the PREDICATE, not the arithmetic.** A count can be exactly right while the claim it supports is
   false. Re-derive the SET the arithmetic ranges over, independently, and state its cardinality first.
   *"I re-derived it and it matches"* is the weakest clearance there is.

5. **Self-contradiction is a real finding** even when you cannot tell which side is right — if two passages
   assert incompatible things, that is upheld, and you cite both anchors. **But a RETRACTION is not a
   self-contradiction.** A passage saying *"X was asserted and is now refuted; the truth is Y"* is correct
   prose. Only book it if the retraction and the thing it retracts are both asserted as live.

6. **A DERIVED-ONCE-AND-POINTED-AT value is one predicate, not many.** A cardinality derived in exactly one
   place, which other documents *point at*, is one assertion. A pointer that correctly names its source is
   **not** a separate assertion of the value. If the derived value is wrong, the predicate lives at the
   derivation site; the pointers are not additional anchors.

7. **A HISTORICAL anchor is not a rotted one.** A passage that says *"what a prior audit found at line N"*
   is a record of where something once was. It is correct prose even when line N now holds something else.
   **The corpus was just repaired in a way that deliberately DELETED several same-file line pins and
   replaced them with construct names** — a passage that names a construct instead of a line is doing the
   right thing, not being vague.

8. **Scope each upheld blocker** as IN-SCOPE (inside `corpus/services/**` or `corpus/architecture/**`) or
   OUT-OF-SCOPE. Only in-scope upheld BLOCKERS enter `N` and `P`.

## Two clone sets exist for `rosetta-extensions`

- `stack-demo/rosetta-extensions` — the **pinned per-stack consumption clone**, `09d06070`. **A corpus
  claim about what the tooling DOES ON A STACK is settled here**, because that is the code a stack runs.
- `.agentspace/rosetta-extensions` — the **authoring copy**, `43049308` on `main`. Where the next tag is
  written; not what any stack executes. **But a claim about a FENCE'S OWN VERDICT OR CONFIGURATION is
  settled HERE**, because a verdict is a measurement taken with that fence's config.

⚠ **The frozen instrument's line 37 names only the authoring copy.** That is a known, deliberately
delivered-unfixed instrument defect (`DEF-M257x-iter101-briefing-rext-tree`); the seats' addendum names
the rule above, but a seat may still have graded against the wrong tree. **Say which tree you read**, and
label such a rejection `wrong-tree`.

## Deduplicate across seats — but only where the PREDICATE is the same

Two seats booking the same underlying false predicate at different anchors is **one** predicate and **two**
anchors. Report both, and say which anchors collapse onto which predicate. Do not collapse two anchors that
merely look similar.

## Output

Write your verdict file to the path you are given. For each booking:

```
<seat> <Bn> | <corpus anchor> | UPHELD|REJECTED | IN-SCOPE|OUT-OF-SCOPE | PREDICATE: <false proposition, <=20 words>
   evidence: <the repo/file:line YOU opened, and what it says>
   [if REJECTED] class: ref-discipline | wrong-tree | mis-read | already-true | pointer-not-assertion | historical-anchor | other — <one line>
```

Then, before the summary, a **PREDICATE ROLL-UP** section listing each distinct upheld in-scope predicate
once, with every anchor that collapses onto it:

```
P1 | <predicate> | anchors: <seat Bn @ file:line>, <seat Bn @ file:line>, ...
```

Then a summary line:
`BOOKED=<n> UPHELD=<n> REJECTED=<n> IN-SCOPE-UPHELD-BLOCKERS=<n> DISTINCT-IN-SCOPE-PREDICATES=<n>`

## Hard bars

- **Do not read `knowledge/plan/**` except this brief, your assigned seat reports, and your own output.**
  Prior audits' ledgers are answer keys; reading them turns adjudication into agreement.
- Read-only. Zero edits outside your one output file. No git state changes. **No fetch.**
- `export DEVELOPER_DIR=/Library/Developer/CommandLineTools`; `cd /Users/marco/workspace/anthropos/rosetta` first.
