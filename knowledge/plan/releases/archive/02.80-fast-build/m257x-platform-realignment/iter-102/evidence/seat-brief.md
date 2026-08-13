# iter-102 REPAIR seat brief — shared by all 10 seats

You are a **REPAIR seat**. You fix documentation defects that four independent adjudicators have **already
graded UPHELD**. You are not re-auditing them and you are not taking a reading.

## Environment

- `export DEVELOPER_DIR=/Library/Developer/CommandLineTools` in **every** shell.
- Repo: `/Users/marco/workspace/anthropos/rosetta`. Absolute paths only.
- **Do NOT `git fetch` anything, ever.** Another lane owns `stack-demo/**` and a fetch moves the ground
  truth under a concurrent reading. Read the clones freely; never write or fetch them.
- **Do NOT commit.** The orchestrator commits. Do not `git add`, `git stash`, `git checkout --`, or
  `git reset`.
- **Zero platform-repo edits.** You may only edit files under `corpus/`.

## Ground truth — settle every claim against these

| repo | ref | note |
|---|---|---|
| `platform` | `0c91421` | `== git ls-remote origin HEAD` |
| `app` | `ad9f3c49` | `== origin/main`. **Moved from `2035f9a4`** — 5 commits |
| `next-web-app` | `8297c684` · `ant-academy` `22df69dd` · `sentinel` `f2c46190` · `studio-desk` `41ee3575` · `cms` `ca50c817` · `jobsimulation` `462343b0` · `messenger` `fa47850d` · `storage` `4ce8ece5` · `roadrunner` `87d8d443` | |
| rext (**per-stack, pinned** — the tree that settles a rext claim) | `stack-demo/rosetta-extensions` | **NOT** `.agentspace/rosetta-extensions`, which is the authoring copy |

> **The settling tree follows the claim's SUBJECT.** A claim about what a **local stack** runs is settled by
> the demo's build pin. A claim about **production infrastructure** is settled by that repo's `origin/main`.
> Getting this backwards produced 4 rejections in one reading and 1 in the next.

## THE SEARCH RULE — §5 rule 44, and it is not optional

**No single search tool is safe.** The environment's recursive `grep` is `ugrep --ignore-files`, so a
`.gitignore` entry **silently hides tracked files**. NUL-bearing source blinds a different instrument;
nested untracked repos blind a third.

- **An absence is established only by `git grep` at a named ref.** Never by `grep -r`.
- For platform claims: `git -C <clone> grep -n '<pat>' <ref> -- '<glob>'`.
- Cross-check any *negative* finding with a second mechanism before you publish it.

## The six binding rules of this repair

1. **REPAIR BY PREDICATE, NOT BY ANCHOR.** The anchor you are given is *one site of a false statement*. Your
   job is to fix **the statement, everywhere it appears in your files.** Measured: iter-96 turned 13 anchors
   into **51** sites; iter-98 turned 20 into **37**. Booked predicate width has run **2** against live widths
   of **7** and **5** — the booking under-counts by roughly 3×. **Assume every predicate is wider than its
   booking and go looking.**

2. **EXPAND ON BOTH AXES — twin AND paraphrase.** A *twin* is the same sentence restated (quoted-verbatim
   forms; `claim_twin_guard` catches these). A *paraphrase* is the same claim in different words — and
   iter-97 measured that **everything that escaped repair was a paraphrase, 3 of 51**. Search for the
   *fact*, not the *phrasing*: the numbers, the identifiers, the service names, the negations.

3. **TRAP A — restate or drop, NEVER re-anchor.** Where the underlying fact was **deleted** rather than
   moved, do not repoint the citation at some other file to make it resolve. **A correctly-cited false
   statement is worse than a stale one.** If a fact is not measurable from the clone set (the
   `infrastructure` repo is in none), say *"not measurable from this repo"* — do not invent a source.

4. **Do not induce new defects — this is measured and it keeps happening.** iter-98 induced 2 (both inside
   prose it rewrote); iter-100 induced at least 1 (its own parenthetical pushed a table down two rows and
   left the numbers pointing at the wrong rows). The rate is ~2/cycle and is the most stable number in this
   milestone. **Therefore:** after every edit, **re-read the whole surrounding block** and check that (a)
   every line number you did not change still points where it did, (b) every count you did not change is
   still right, (c) you did not leave the document contradicting itself a few lines away.

5. **Never weaken a TRUE clause while fixing a false one.** Several of these sentences are a conjunction
   where one half is correct. Fix the false half; keep the true half; say which is which.

6. **A repair pass contains no reading.** Do not book new blockers as findings and do not estimate a
   residual. If you *notice* something outside your assignment, put it in the `## Noticed, not repaired`
   section of your report — do not fix it if it is outside your files.

## Cross-seat predicates — DO NOT word these yourself

Three predicates span several seats. Their wording is derived once, centrally, in
`iter-102/evidence/canonical-repairs.md`. **Read that file.** If one of your anchors is listed there,
**apply the canonical form** — do not compose your own. Everything else you re-derive normally.

## Your report

Write to `iter-102/evidence/seat-<N>.md`. Structure:

```
# seat-<N> report

## Ledger rows  (the shape claim_ledger.py derives from — a table with a claim column and an anchor column)

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| ... | "<VERBATIM quoted false sentence>" | `file.md:LINE` | <what is true, with evidence at a named ref> | <N> |

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |

## Twins outside my files  (REPORT, do not edit)

| predicate | file:line | why it is the same claim |

## Noticed, not repaired
## What I could not settle, and why
```

**The `the false claim` cell must quote the offending sentence VERBATIM between double quotes.** That quote
*is* the pattern `claim_twin_guard` and `repair_postcondition` adopt — it is what makes the claim
un-republishable tree-wide, and it is the only reason the completeness claim is checkable rather than
asserted. A paraphrase in that cell fences nothing.

**Grade your reach honestly.** If you booked 3 anchors, found 5 sites and repaired 4, say so and say why
the fifth was left. A repair that reports 100 % is less trustworthy than one that reports 80 % with the
residue named.
