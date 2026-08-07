# Adjudicator brief — M257x iter-131, readings #33 / #34

You are an **adjudicator**, not an auditor. Fourteen blind seats read the corpus and booked claims.
**Your job is to decide, independently and against ground truth, whether each claimed BLOCKER is real.**

A seat's confidence is **not** evidence. A seat's citation is **not** evidence until you open it.

## Repo root

`/Users/marco/workspace/anthropos/rosetta` — **always `cd` there first.**

## GROUND TRUTH — the shas that settle a claim

| repo | path | sha |
|---|---|---|
| app | `stack-demo/app` | `ad9f3c498` |
| platform | `stack-demo/platform` | `0c91421df` |
| next-web-app | `stack-demo/next-web-app` | `8297c684c` |
| sentinel | `stack-demo/sentinel` | `f2c461903` |
| studio-desk | `stack-demo/studio-desk` | `41ee3575d` |
| ant-academy | `stack-demo/ant-academy` | `22df69dd8` — ⚠ **DIRTY working tree; read via `git show 22df69dd8:<path>`** |
| storage | `stack-demo/storage` | `4ce8ece52` |
| messenger | `stack-demo/messenger` | `fa47850d9` |
| cms | `stack-demo/cms` | `ca50c8170` |
| graphql-wundergraph | `stack-demo/graphql-wundergraph` | `60c229f39` |
| roadrunner | `stack-demo/roadrunner` | `87d8d4438` |
| jobsimulation | `stack-demo/jobsimulation` | `462343b05` |
| studio-room | `stack-dev/studio-room` | `aeec036a5` (also nested at `stack-demo/app/studio`, its OWN checkout) |
| **rext — pinned per-stack** | `stack-demo/rosetta-extensions` | `09d06070f` |
| **rext — authoring copy** | `.agentspace/rosetta-extensions` | `f2ea567b3` |

**A claim is settled at the ref the claim itself names** (§5 rule 33). If the passage names a ref, read
THAT ref. If it names none, use the checkout above. A pin is a **date**, not an excuse — if the claim is
true at its named ref, it is TRUE, however stale.

**Which rext tree:** a claim about *what the tooling does on a stack* → the **pinned** clone
(`09d06070`). A claim about *a fence's own verdict or configuration* → the **authoring copy**
(`f2ea567b`). **The two are 33 commits apart at this reading.** A seat that read the wrong one has made
a `wrong-tree` error — see the rejection classes below.

## HARD BARS

1. **You MUST NOT read anything under `knowledge/plan/**` except**: this brief, and the specific
   `iter-131/raw/*.md` seat reports you are assigned. **No other iter dir, no `progress.md`, no
   `decisions.md`, no prior adjudication.** Those are answer keys from earlier readings and using them
   would make you measure agreement instead of truth.
2. **Read-only.** Zero edits to any file except your single output report.
3. **Do not read other adjudicators' outputs.**

## YOUR ASSIGNMENT

You will be told which seat reports to adjudicate. For **every BLOCKER** in them:

### Step 1 — verdict

- **UPHELD** — the claim is true: the corpus text really is false, unsupportable, or self-contradictory
  against ground truth. **You opened the evidence yourself and it holds.**
- **REJECTED** — the claim does not stand. Give the class:
  - `wrong-tree` — the seat read the wrong repo/tree/ref (esp. the two rext trees, or the dirty
    ant-academy working tree instead of its ref). **Label this class explicitly; it is counted
    separately.**
  - `misread` — the seat misread the corpus text or the source.
  - `true-at-its-ref` — the corpus claim is pinned and correct at its own named ref; the seat graded it
    at a different ref.
  - `retraction-not-contradiction` — the corpus says *"X was wrong; the truth is Y"*. That is correct
    prose doing its job, **not** a self-contradiction. Only uphold if Y is itself false, or the
    retraction misdescribes what it retracts.
  - `minor-not-blocker` — real but cosmetic/immaterial; would not mislead a reader doing real work.
  - `not-in-scope` — the anchor is outside `corpus/services/**` + `corpus/architecture/**`.
- **CANNOT-SETTLE** — you genuinely cannot decide (e.g. the subject repo is in no clone set). Say what
  evidence would settle it. **Do not launder a cannot-settle into either a rejection or an uphold.**

### Step 2 — the PREDICATE (this is what the primary metric counts)

For every **UPHELD** blocker, write the **PREDICATE** it falsifies — the general proposition, stated in
one line, independent of where it appears. Examples of the level of abstraction wanted:

- `"cms's production ECS state is UNMEASURABLE **because** infrastructure is in no clone set"`
- `"the ai library is imported as a private Go module by something a stack builds"`
- `"the Cosmo/WunderGraph router still runs in production"`

> ⚠️ **CORRECTED M257x iter-136 — the first example used to read *"cms's production ECS state is
> unmeasurable **/** infrastructure was never in a clone set"*, and that slash was a defect in the
> instrument.** It joins a FALSE proposition to a TRUE one and invites an adjudicator to book the
> conjunction. **`infrastructure` really is in no clone set** — `make init` does not clone it and no
> `stack-*/` holds it. What is false is the **inference** from that to *unmeasurable*, since the repo
> was read at `13c248e6`. **Two independent adjudicators caught this** — `adj-1` at iter-131 (which
> moved the reading's `P` from 30 to 29) and `adj-C` at iter-135, which traced it back to *this line*.
> **A brief that models a conflated predicate teaches the error it exists to catch.** State predicates
> so that **every conjunct is independently false**, or split them.

**Two seats booking the same proposition at two different anchors share ONE predicate.** Two different
propositions at the same anchor are TWO predicates. This distinction is the whole measurement, so state
the predicate deliberately rather than paraphrasing the seat's title.

### Step 3 — classify each upheld blocker

- **class:** one of `intra-corpus-citation` (a corpus→corpus pin that resolves wrong) ·
  `platform-drift` (corpus disagrees with platform source) · `self-contradiction` (two corpus passages
  incompatible) · `arithmetic/count` · `other` (name it)
- **anchor:** `corpus/<file>:<line>`
- **multi-pin block?** yes/no — does the anchor sit in a block carrying several `file:line` pins?
- **repair-induced?** yes/no — **run `git log -L<line>,<line>:<file> --oneline | head -3`** on the anchor.
  If its most recent touching commit is one of iters 120–130 (subjects matching `iter(M257x/12[0-9])`,
  `iter(M257x/130)`, `fix(M257x/12[0-9])`, `probe(M257x/1[23][0-9])`), mark **yes**. State the sha.

## OUTPUT

Write your full report to the exact path you are given. Structure:

```markdown
# Adjudication <id> — seats <list>

## Verdict table
| seat | B# | anchor | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |

## Upheld predicates, deduplicated within my assignment
P-<id> | <one-line predicate> | anchors: <corpus file:line, ...> | class

## Rejections, with the evidence I opened
<one short paragraph each — say what you read and why it does not stand>

## Cannot-settle
<each, with what would settle it>

## Counts
UPHELD=<n> REJECTED=<n> (of which wrong-tree=<n>) CANNOT-SETTLE=<n>
DISTINCT-PREDICATES-IN-MY-SET=<n>
```

Then return, as your final message, ONLY:
```
UPHELD=<n> REJECTED=<n> WRONGTREE=<n> CANNOTSETTLE=<n> PREDICATES=<n>
```
followed by one line per upheld predicate: `P | <predicate> | <anchors>`.

## The standing traps, because they have caught adjudicators before

- **A count can be exactly right while the claim it supports is FALSE.** Verify the predicate, not the
  arithmetic.
- **Re-derive the SET from source, not the sum from the set.** *"I re-derived it and it matches"* is the
  weakest clearance a report can contain.
- **Never let a search's stderr go unread**, and **run a positive control in the same pass** — a pattern
  you know matches. An empty result from a FAILED command is not evidence of absence.
- **`git grep` and bare `grep` disagree** (this shell's `grep` is `ugrep --ignore-files`, which skips
  gitignored files), and **nested repos are invisible to `git grep` at the host ref**. State which trees
  your number covers.
- **A retraction is not a self-contradiction.** See the rejection class above — this is the single most
  common over-booking in this corpus.
- **Do not uphold a blocker because it is well written.** Open the evidence.
