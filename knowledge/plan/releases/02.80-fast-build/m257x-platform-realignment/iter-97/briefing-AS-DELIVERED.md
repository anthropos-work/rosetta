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
---

# ADDENDUM — SUPERSEDING GROUND TRUTH (iter-97, readings #19/#20)

**Everything above this line is the frozen instrument, byte-identical.** This addendum supersedes the
briefing's sha table ONLY. It was appended after the copy's sha256 was verified against the source.

## Ground truth — re-derived at this reading's open, 2026-08-06

| clone | path (relative to repo root) | checkout | `origin/main` | behind |
|---|---|---|---|---|
| platform | `stack-demo/platform` | **`0c91421d`** | `0c91421d` | 0 |
| app | `stack-demo/app` | `b948604f` | `2035f9a4` | **93** |
| **app/studio** (studio-room, a NESTED repo) | `stack-demo/app/studio` | **`aeec036a`** | — | — |
| **cms/studio** (the SAME repo, second copy) | `stack-demo/cms/studio` | **`aeec036`** | — | — |
| next-web-app | `stack-demo/next-web-app` | `bb3313bc` | `8297c684` | **41** |
| storage | `stack-demo/storage` | `4ce8ece5` | `9f8cb532` | **20** |
| messenger | `stack-demo/messenger` | `fa47850d` | `e9421c68` | **7** |
| ant-academy | `stack-demo/ant-academy` | `9c3843cd` | `22df69dd` | **5** |
| jobsimulation | `stack-demo/jobsimulation` | `462343b0` | `82cb66ec` | **4** |
| cms | `stack-demo/cms` | `ca50c817` | `f38c0c4a` | **2** |
| sentinel | `stack-demo/sentinel` | `88bc5592` | `f2c46190` | **2** |
| studio-desk | `stack-demo/studio-desk` | `14a5442a` | `41ee3575` | **2** |
| roadrunner | `stack-demo/roadrunner` | `87d8d443` | `87d8d443` | 0 |
| graphql-wundergraph | `stack-demo/graphql-wundergraph` | `60c229f3` | `60c229f3` | 0 |
| rosetta-extensions (authoring) | `.agentspace/rosetta-extensions` | `6130bfd8` | `6130bfd8` | 0 |

The corpus under audit is `rosetta` @ **`00be1ac`**, tree clean.

## The two rules that go with the sheet

1. **Grade at the ref the claim names — UNLESS the sentence asserts currency**, in which case no
   neighbouring pin rescues it. A pin is a **date**, not an excuse.
2. **A corpus citation into `rosetta-extensions` grades against the AUTHORING copy** (`6130bfd8`)
   unless the citing block pins a ref. The `stack-demo/` clone is an older clone of the same repo.

## ⚠️ NEW AND LOAD-BEARING — how to establish an ABSENCE in this environment

The shell's `grep` is a **function** that execs `ugrep -G --ignore-files --hidden -I`. **A bare
recursive `grep` cannot establish an absence**, and neither can a naive `git grep`. Three mechanisms,
measured this milestone:

1. **`.gitignore` hides TRACKED files from `grep`.** 12 non-empty text files across the clone set.
   `git grep` sees these; bare `grep` does not.
2. **NUL-bearing source is skipped by BOTH `grep -I` and `git grep`**, and `file(1)` calls it "data".
   Two exist: `next-web-app/apps/web/src/hooks/useCoursebuilder.ts` (50,433 bytes, 1,178 NULs) and
   `ant-academy/code/src/time/store.js`. A zero over a tree containing one is a zero with a hole in it.
3. **Nested untracked repos — where `git grep` at the HOST ref is a GUARANTEED zero.**
   `stack-demo/app/studio` and `stack-demo/cms/studio` are each the `anthropos-studio-room` repo
   (177 files, own HEAD `aeec036`), hidden by `app/.gitignore:79` and `cms/.gitignore:129`.
   `git -C app grep <anything> HEAD -- studio/` returns **0 for every predicate ever**. On one real
   predicate the three instruments returned **1 / 0 / 22** — and the **0** was the ref-named `git grep`.

**So: enumerate the trees first, grep each at ITS OWN ref, and say which trees your number covers.**

```bash
find stack-demo -name .git -maxdepth 4 | sed 's|/\.git$||' | sort   # the search set, nested repos included
git -C <tree> grep -n '<term>' <that tree's ref>                     # one tree, one ref, per tree
```

A tree-wide zero that does not name its sub-repos is an **unproven zero**, and booking a claim as
correct on the strength of one is a false clearance — which has happened inside a reading before.

## One more thing to know about this corpus

The tree you are reading was **repaired immediately before this reading**, at commit `00be1ac`
(23 files, +230 −79). Repairs manufacture defects: a prior iteration of this milestone measured
**9 of 18** findings as induced by the repair that preceded them, and **8 of those 9** were one
mechanical class — *a claim corrected at one site and left standing at another*. **Read for that.**
Do not assume a recently-touched passage is more reliable than an old one; assume the opposite.
