# Auditor briefing — corpus fidelity read (readings #13 / #14, M257x iter-76)

Held fixed at iter-41's instrument on every knob. The **only** changes from
`briefing-iter53-AS-RUN.md` are the ground-truth shas (the clones moved) and one added paragraph
about how to choose the ref for an `app` claim — recorded here rather than applied silently, because
an instrument that is *described* rather than *stored* is not frozen (§5 rule 25).

You are one auditor of a multi-seat fidelity audit of a **documentation corpus** (`rosetta`) that
describes a software platform (`anthropos`). Your job: **read your assigned files in full,
top-to-bottom, and find every claim that is FALSE against the platform's actual source code, or that
contradicts another claim in the corpus.**

## Repo root

`/Users/marco/workspace/anthropos/rosetta` — **always `cd` there first; never assume cwd.**

## GROUND TRUTH — the only thing that settles a claim

Real platform source clones, at these exact shas, **re-derived at this reading's open**. Read them.
Never guess.

| repo | path (relative to repo root) | sha |
|---|---|---|
| app (the backend monolith) | `stack-demo/app` | `b948604f` (v1.366.0) |
| app/studio (studio-room, embedded) | `stack-demo/app/studio` | in-tree under the `app` sha above |
| platform (orchestrator: `repos.yml`, compose, Makefile) | `stack-demo/platform` | `0dab54df` — **level with `origin/main`** |
| next-web-app | `stack-demo/next-web-app` | `bb3313bc` |
| sentinel | `stack-demo/sentinel` | `88bc5592` |
| storage | `stack-demo/storage` | `4ce8ece5` |
| messenger | `stack-demo/messenger` | `fa47850d` |
| cms (legacy, merged into app) | `stack-demo/cms` | `ca50c817` |
| graphql-wundergraph (the deleted Cosmo router) | `stack-demo/graphql-wundergraph` | `60c229f3` |
| roadrunner (legacy, merged) | `stack-demo/roadrunner` | `87d8d443` |
| jobsimulation (legacy, merged) | `stack-demo/jobsimulation` | `462343b0` |
| studio-desk | `stack-demo/studio-desk` | `14a5442a` |
| ant-academy | `stack-demo/ant-academy` | `9c3843cd` |
| rosetta-extensions (the tooling) | `.agentspace/rosetta-extensions` | authoring copy, `main` |

### Which ref settles an `app` claim (new in this reading — read it)

The `app` **checkout** is `b948604f` (v1.366.0). Its **`origin/main` is four commits past
`v1.367.0`**, so the working tree is not the newest thing there is. Both are legitimate references
and they disagree about line numbers — one working day of platform commits has moved this repo's
`main.go` citations before.

The rule, which is the corpus's own (`platform-alignment.md` §5 rule 33): **a claim is settled at
the ref the claim itself names.** So:

- If the passage names a ref (*"at `app` `9d00a313` v1.367.0"*, *"measured at `b948604`"*), read
  **that** ref: `git -C stack-demo/app show <sha>:<path>`. A pin is a **date**, not an excuse — the
  claim is graded as of that date, and if it is true there it is **true**, however stale.
- If the passage names **no** ref, grade it against the **checkout** (`b948604f`) — that is the file
  on disk and the one your `sed`/`grep` will read.
- If a passage's line anchor resolves at neither, **that is a finding**, and say which refs you
  tried.

A pin's scope is **the claim's own block** — a markdown *cell* in a table, a wrapped sentence in
prose. A ref named in a neighbouring row does not date this row's claim.

## HARD BARS — non-negotiable

1. **You MUST NOT read anything under `knowledge/plan/**`.** That directory holds the answer keys of
   prior audits. Reading it makes your pass measure agreement instead of detection, and destroys the
   experiment. Do not `grep` it, do not `ls` it, do not open it.
2. **You MUST NOT read any file under `.agentspace/scratch/`.**
3. **You are one of several seats reading in parallel. Do not look for, or read, any other seat's
   output.**
4. **Read-only.** Make **zero** edits to any file except your single output report. Do not run `git`
   commands that change state. Never `git stash`, `reset`, `checkout --`, `clean`, or `rm`.

## METHOD — read the whole file, do not grep-and-conclude

- For **every** assigned file: run `wc -l <file>` first and state the number in your report. That is
  your **positive control** — it proves you actually opened the file. A file whose line count you
  cannot state was not read, and a partial pass is not a reading.
- Then read the file **top-to-bottom, in full**. Do not sample. Do not skim to headings.
- When a claim cites `file:line` in a platform repo, **open that file and read the lines AROUND the
  cited line** (`sed -n 'N-15,N+5p'`). A constant's meaning lives in its surroundings. Grepping to a
  line and reading only that line is the cheapest way to be confidently wrong about code you have
  "checked".

## THE RULES THAT ACTUALLY CATCH THINGS

Read `corpus/ops/platform-alignment.md` §5 in full before you start — it is the protocol these rules
come from and it is **in scope for you to read** (it is not under `knowledge/plan/`). The ones that
matter most:

- **Never let a search's stderr go unread.** An engine rejection is indistinguishable from "no
  matches".
- **Run a positive control in the same pass** — a pattern you know matches. If it returns 0, your
  pipeline is broken, not the corpus. **An empty result from a FAILED command is not evidence of
  absence.**
- **Check the field name before concluding absence.** The cheapest false absence is a wrong regex.
- **A count can be exactly right while the claim it supports is FALSE.** Verify the **PREDICATE**,
  not just the arithmetic.
- **§5 rule 24 — the one three seats failed simultaneously, so it is a standing instruction.** When a
  corpus passage **shows its own derivation**, the visible arithmetic is an *attractor*: it is
  checkable, it checks out, and checking it feels like auditing the claim. **The incompleteness is
  never in the arithmetic — it is in the SET the arithmetic ranges over.** So: **re-derive the SET
  from source, not the sum from the set.** Enumerate the predicate independently and state its
  cardinality *before* doing any arithmetic. And note: **`grep -c` over source counts commented-out
  code** — exclude what does not compile. Treat *"I re-derived it and it matches"* as the
  **weakest** clearance a report can contain.
- **A claim that was refuted once can come back.** The corpus has been repaired many times;
  corrections have been re-broken and retracted claims re-published verbatim. A confident,
  well-formatted sentence is not evidence.
- **Self-contradiction counts.** If two corpus files (or two passages of one file) assert
  incompatible things, that is a finding even if you cannot tell which side is right — say so and
  cite both anchors.

## PLATFORM CONTEXT you will need (verify it, do not trust it)

The platform team has been **merging microservices back into `app`**, toward a monolith. Services
historically folded in: **skiller**, **skillpath**, **roadrunner**, **jobsimulation**, **cms**, and
most recently **storage** and **messenger**. The **Cosmo/WunderGraph router was deleted outright**
from local dev, so GraphQL is served straight from `backend`. A corpus passage describing any of
these as a live standalone service, container, port, subgraph, or DB schema is suspect — but
**check `stack-demo/app` and `stack-demo/platform/repos.yml` before booking it**, because some are
correctly described as *archived/merged*, and some legacy repos still exist on disk while being
decommissioned. **The reverse error is equally live**: a corpus passage may over-claim the merge,
and a service can be *merged in prod* while still *startable locally under a legacy profile*. Those
are different states and conflating them is a finding in either direction.

## GRADING — every finding is exactly one of these

**BLOCKER** — a claim that is **false or unsupportable against ground truth**, or that
**contradicts another corpus claim**, and that would **mislead a reader doing real work**. Wrong
counts, wrong file paths, wrong line anchors that point at the wrong construct, a named
service/table/schema/env-var/port that does not exist, a mechanism described backwards, a retraction
that contradicts what it retracts.

**MINOR** — cosmetic or non-material: typos, doubled words, formatting damage, a stale-but-harmless
phrasing, a broken markdown structure with no factual consequence.

When in doubt between the two, **book it as a BLOCKER and say why you hesitated.** Under-booking is
the failure mode this instrument exists to fight.

## OUTPUT — write ONE file, then return a COMPACT summary

Write your full report to the exact path you were given. Structure:

```markdown
# Seat <X> — reading #<NN>

## Positive controls
| file | wc -l | read in full? |
(one row per assigned file — every row must say yes)

## Blockers
### B1 — <one-line title>
- **Anchor:** `<corpus file>:<line>` — quote the exact offending text
- **Why it is false:** <the ground-truth evidence, with `repo/file:line` citations you actually opened>
- **Confidence:** high | medium | low
(repeat)

## Minors
- <one line each, with anchor>

## Positively cleared (audited zeros)
For each claim you checked and found CORRECT, state which of these you did — they are NOT equivalent:
- **ENUMERATED** — I enumerated the predicate from source independently and compared cardinalities.
- **RE-DERIVED** — I re-computed the arithmetic the document itself displays.
Only ENUMERATED is a clearance. Label each one.

## What I could not settle
<claims you could not verify, and what evidence would settle them>
```

**Then return, as your final message, ONLY a compact list**: one line per BLOCKER in the form
`B<n> | <corpus file>:<line> | <≤15-word title> | <confidence>`, then a one-line count summary
`BLOCKERS=<n> MINORS=<n>`. Do not paste the full report back — it is on disk.

---

# ADDENDUM — readings #33 / #34 (M257x iter-131), 2026-08-07

**The instrument above is a VERBATIM COPY** of `instrument/briefing-iter76-AS-RUN.md`
(sha256 `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0`, re-checked AFTER copying —
both sides printed in `ground-truth.md`, `diff` returned empty; `git log --follow` shows exactly one
commit ever, `012edd2`). **Nothing above this line was edited.** Ground truth that has moved since it was
frozen is superseded here, in the open, rather than silently.

## SUPERSEDED ground truth — read THESE shas, not the table above

| repo | path | sha at THIS reading |
|---|---|---|
| **platform** | `stack-demo/platform` | **`0c91421d`** — level with `origin/main` |
| **app** | `stack-demo/app` | **`ad9f3c49`** — level with `origin/main` (the table above says `b948604f`; that is stale, though still a reachable ancestor, so a claim *pinned* at `b948604` is still graded there) |
| **next-web-app** | `stack-demo/next-web-app` | **`8297c684`** (checkout); `origin/main` is `f97ba659` |
| **sentinel** | `stack-demo/sentinel` | **`f2c46190`** — level with `origin/main` |
| **studio-desk** | `stack-demo/studio-desk` | **`41ee3575`** — level with `origin/main` |
| **ant-academy** | `stack-demo/ant-academy` | **`22df69dd`** — level with `origin/main` |

**Unchanged and re-verified at this open:** storage `4ce8ece5`, messenger `fa47850d`, cms `ca50c817`,
graphql-wundergraph `60c229f3`, roadrunner `87d8d443`, jobsimulation `462343b0`. For four of these six the
**checkout is behind `origin/main`** (storage `9f8cb532`, messenger `e9421c68`, cms `f38c0c4a`,
jobsimulation `82cb66ec`); both refs are legitimate and the rule above decides which.

⚠ **`stack-demo/ant-academy` has a DIRTY WORKING TREE at this reading** — three modified files
(`code/public/catalog.json`, `code/src/components/UserMenu.jsx`, `code/src/lib/serverChapterBody.js`).
**Read it with `git -C stack-demo/ant-academy show 22df69dd8:<path>`, never `sed`/`cat` on the tree**, or
you will grade the corpus against edits that are not in any commit.

**Nested repos — `stack-demo/app/studio` and `stack-demo/cms/studio` are each their OWN git checkout at
`aeec036a`**, not part of `app`'s tree. `git -C stack-demo/app grep <anything> HEAD -- studio/` returns
**0 for every predicate**, which reads exactly like evidence and is not. Grep them at their own ref.

**No platform clone has advanced since iter-103's reading** — this is the fifth consecutive frozen
reading. That is a property of the reading, **not a reason to relax**: platform-drift was ~33 % of
iter-109's residual and ~38 % of iter-116's, over a subject in which literally nothing changed. **A claim
can be false against a sha that never moved.**

## WHICH TREE SETTLES A `rosetta-extensions` CLAIM — read this, it overrides line 37 above

| role | path | sha |
|---|---|---|
| **per-stack CONSUMPTION clone (pinned)** | `stack-demo/rosetta-extensions` | **`09d06070`** |
| authoring copy | `.agentspace/rosetta-extensions` | **`f2ea567b`** on `main` |

**A corpus claim about what the tooling DOES ON A STACK is settled by the pinned per-stack clone**, because
that is the code a stack executes. **Say which tree you read.** (`platform-alignment.md` §5 rule 45.)

**One refinement, and it cost an earlier reading's coordinator a nearly-published false conclusion:** a
claim about a *fence's own verdict or configuration* is settled by the tree that configuration lives in —
the authoring copy — because a verdict is a measurement taken with that fence's config. Run a fence from
the pinned clone and you measure last release's fence. **The two trees are 33 commits apart at this
reading**, the widest gap any reading has faced, so this distinction bites harder than usual.

## The three mechanisms that hide files from a search — use all three instruments

An absence is **not** established by any single tool. Measured on the `mistralai` predicate, three
instruments returned **1 / 0 / 22**, and the **0 was the ref-named `git grep`**:

1. **`.gitignore` hides TRACKED files from bare `grep`** — this shell's `grep` is `ugrep --ignore-files`.
   `git grep` sees them; bare `grep` does not.
2. **NUL-bearing source is skipped by BOTH `grep -I` and `git grep`.** One byte is enough. Count NUL
   **bytes** with `tr -dc '\000' < FILE | wc -c`; `grep -c` counts LINES.
3. **Nested untracked repos are invisible to `git grep` at the HOST ref** — the two `studio` checkouts.

**State which trees your number covers.** A tree-wide zero that does not name its sub-repos is unproven.

## Scope of THIS reading

`corpus/services/**` + `corpus/architecture/**` — **41 files, 11,922 lines**, partitioned across 7 seats.
Two independent readings of the identical partition (#33 and #34). You are one blind seat: you do not know
which reading you are in, and you must not seek out any other seat's output.

## What HAS changed in the corpus since it was last read — and it is a lot

**This is the opposite of the previous reading's situation, and you should read differently because of
it.** The last reading measured a corpus 5 changed lines from its predecessor. Since then the in-scope
corpus has moved **31 files, +1,248 / −197 lines**, and grown **10,871 → 11,922 lines (+9.7 %)** with one
net-new file (`architecture/org-repos.md`, 466 lines).

That prose was written by **repair passes**, and repair has a measured induction rate in this corpus.
So:

- **Newly written prose is NOT pre-cleared.** Some of the most recently edited paragraphs are the least
  audited ones in your files. A confident, well-cited, freshly written sentence is exactly the shape a
  repair-induced defect takes.
- **Retractions are correct prose, not self-contradictions.** A passage saying *"X was wrong; the truth
  is Y, measured at `<ref>`"* is doing its job. **Book it only if Y is itself false, or if the retraction
  contradicts what it claims to retract.** Do not book the mere presence of a corrected claim.
- **Equally, do not clear a claim because it carries a citation.** Open the citation. A pin that resolves
  to a *different construct* than the sentence describes is a BLOCKER, and line numbers in this corpus
  have moved three times in recorded history.
- **Do not book a claim merely because it looks old.**

**Treat every claim as unaudited, including — especially — the new ones.**
