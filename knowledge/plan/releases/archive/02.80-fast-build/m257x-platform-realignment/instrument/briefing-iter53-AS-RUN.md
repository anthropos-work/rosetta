# Auditor briefing — corpus fidelity read

You are one auditor of a multi-seat fidelity audit of a **documentation corpus** (`rosetta`) that describes
a software platform (`anthropos`). Your job: **read your assigned files in full, top-to-bottom, and find
every claim that is FALSE against the platform's actual source code, or that contradicts another claim in
the corpus.**

## Repo root

`/Users/marco/workspace/anthropos/rosetta` — **always `cd` there first; never assume cwd.**

## GROUND TRUTH — the only thing that settles a claim

Real platform source clones, at these exact shas. Read them. Never guess.

| repo | path (relative to repo root) | sha |
|---|---|---|
| app (the backend monolith) | `stack-demo/app` | `5ba17044` |
| app/studio (studio-room, embedded) | `stack-demo/app/studio` | `aeec036a` |
| platform (orchestrator: `repos.yml`, compose, Makefile) | `stack-demo/platform` | `2adcf714` |
| next-web-app | `stack-demo/next-web-app` | `bb3313bc` |
| sentinel | `stack-demo/sentinel` | `88bc5592` |
| storage | `stack-demo/storage` | `4ce8ece5` |
| messenger | `stack-demo/messenger` | `fa47850d` |
| cms (legacy, merged into app) | `stack-demo/cms` | `ca50c817` |
| graphql-wundergraph (Cosmo router) | `stack-demo/graphql-wundergraph` | `60c229f3` |
| roadrunner (legacy, merged) | `stack-demo/roadrunner` | `87d8d443` |
| jobsimulation (legacy, merged) | `stack-demo/jobsimulation` | `462343b0` |
| studio-desk | `stack-demo/studio-desk` | `14a5442a` |
| ant-academy | `stack-demo/ant-academy` | `9c3843cd` |
| rosetta-extensions (the tooling) | `.agentspace/rosetta-extensions` | (authoring copy, `main`) |

## HARD BARS — non-negotiable

1. **You MUST NOT read anything under `knowledge/plan/**`.** That directory holds the answer keys of prior
   audits. Reading it makes your pass measure agreement instead of detection, and destroys the experiment.
   Do not `grep` it, do not `ls` it, do not open it.
2. **You MUST NOT read any file under `.agentspace/scratch/`.**
3. **You are one of several seats reading in parallel. Do not look for, or read, any other seat's output.**
4. **Read-only.** Make **zero** edits to any file except your single output report. Do not run `git` commands
   that change state. Never `git stash`, `reset`, `checkout --`, `clean`, or `rm`.

## METHOD — read the whole file, do not grep-and-conclude

- For **every** assigned file: run `wc -l <file>` first and state the number in your report. That is your
  **positive control** — it proves you actually opened the file. A file whose line count you cannot state
  was not read, and a partial pass is not a reading.
- Then read the file **top-to-bottom, in full**. Do not sample. Do not skim to headings.
- When a claim cites `file:line` in a platform repo, **open that file and read the lines AROUND the cited
  line** (`sed -n 'N-15,N+5p'`). A constant's meaning lives in its surroundings. Grepping to a line and
  reading only that line is the cheapest way to be confidently wrong about code you have "checked".

## THE RULES THAT ACTUALLY CATCH THINGS

Read `corpus/ops/platform-alignment.md` §5 in full before you start — it is the protocol these rules come
from and it is **in scope for you to read** (it is not under `knowledge/plan/`). The ones that matter most:

- **Never let a search's stderr go unread.** An engine rejection is indistinguishable from "no matches".
- **Run a positive control in the same pass** — a pattern you know matches. If it returns 0, your pipeline
  is broken, not the corpus. **An empty result from a FAILED command is not evidence of absence.**
- **Check the field name before concluding absence.** The cheapest false absence is a wrong regex.
- **A count can be exactly right while the claim it supports is FALSE.** Verify the **PREDICATE**, not just
  the arithmetic.
- **§5 rule 24 — the one three seats failed simultaneously, so it is promoted to a standing instruction.**
  When a corpus passage **shows its own derivation**, the visible arithmetic is an *attractor*: it is
  checkable, it checks out, and checking it feels like auditing the claim. **The incompleteness is never in
  the arithmetic — it is in the SET the arithmetic ranges over.** So: **re-derive the SET from source, not
  the sum from the set.** Enumerate the predicate independently and state its cardinality *before* doing any
  arithmetic. And note: **`grep -c` over source counts commented-out code** — exclude what does not compile.
  Treat *"I re-derived it and it matches"* as the **weakest** clearance a report can contain.
- **A claim that was refuted once can come back.** The corpus has been repaired many times; corrections have
  been re-broken and retracted claims re-published verbatim. A confident, well-formatted sentence is not
  evidence.
- **Self-contradiction counts.** If two corpus files (or two passages of one file) assert incompatible
  things, that is a finding even if you cannot tell which side is right — say so and cite both anchors.

## PLATFORM CONTEXT you will need (verify it, do not trust it)

The platform team has been **merging microservices back into `app`**, toward a monolith. Services
historically folded in: **skiller**, **skillpath**, **roadrunner**, **jobsimulation**, **cms**. A corpus
passage describing any of these as a live standalone service, container, port, subgraph, or DB schema is
suspect — but **check `stack-demo/app` and `stack-demo/platform/repos.yml` before booking it**, because some
of them are correctly described as *archived/merged* and some legacy repos still exist on disk while being
decommissioned. The reverse error is equally live: a corpus passage may over-claim the merge.

## GRADING — every finding is exactly one of these

**BLOCKER** — a claim that is **false or unsupportable against ground truth**, or that **contradicts another
corpus claim**, and that would **mislead a reader doing real work**. Wrong counts, wrong file paths, wrong
line anchors that point at the wrong construct, a named service/table/schema/env-var/port that does not
exist, a mechanism described backwards, a retraction that contradicts what it retracts.

**MINOR** — cosmetic or non-material: typos, doubled words, formatting damage, a stale-but-harmless
phrasing, a broken markdown structure with no factual consequence.

When in doubt between the two, **book it as a BLOCKER and say why you hesitated.** Under-booking is the
failure mode this instrument exists to fight.

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
