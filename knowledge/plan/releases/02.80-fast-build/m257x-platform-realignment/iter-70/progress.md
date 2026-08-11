**Type:** tik — under `TOK-05`, step 2 (citations), clearing the last *named* prerequisite before
step 4, the graded read.

# iter-70 — the class I routed one iteration ago was twelve ports and four citations

## Phase A — the re-survey re-scoped the iter before a single repair

iter-69 routed **`FIX-M257x-iter69-pathless-antecedent` — "23 citations across 14 lines"** and called
it *"a prerequisite for the graded read."* Six were `shared_libraries.md:70`'s and were repaired in
iter-69 itself. Of the remaining 17, the *values* were read rather than counted:

> `:5050` · `:8082` · `:3000` · `:8000` · `:9000` · `:7070`

**Port numbers.** *"GraphQL is served by `backend` itself at `:8082/graphql/query`"* is a URL, not a
citation into `docker-compose.yml`.

## Phase B — one mechanical discriminator, and its own defect reported

**Does the antecedent file even HAVE that many lines?** `repos.yml` has 31. `docker-compose.yml` has
271. `cors.go` has 110.

| | n |
|---|---|
| impossible as a line number | **12** |
| line-plausible | 5 |
| …of which one is a port the heuristic mis-sized | 1 |
| **real** | **4, across 3 lines** |

The heuristic sized the antecedent by `max()` over every file on disk sharing that basename, so
`labsapi/client.go` resolved to a 23,437-line namesake and *"default `:7070`"* came back
line-plausible. **Reported rather than quietly corrected** — a discriminator that flatters its own
recall is the exact failure mode this milestone keeps finding, and the sound version resolves the
antecedent to exactly one file, which `resolve_citation` already does.

**Off by a factor of six, one iteration after routing it.** §5 rule 32 says *re-derive the hand-off's
numbers — including the orchestrator's*, on the evidence of two M257x iterations that corrected an
orchestrator-supplied fact. This is that rule against **my own** hand-off.

## Phase C — the four, adjudicated

| site | verdict |
|---|---|
| `security_compliance.md:71` — `org_membership.go:172-188` | **HELD** @ `9d00a313` — exactly `func (Membership) Policy() ent.Policy`, ending in `privacy.AlwaysDenyRule()` |
| `ai-readiness.md:45` — `` `:33` `` | **HELD — and the detector was wrong.** It binds to the yaml already named on the line, whose `:33` is the `# v2.7 M254 RE-POINT` header the sentence quotes |
| `ai-labs.md:63` — `` `:7070` `` | **not a citation** — a default port |
| `studio-room.md:388` — `` `:36` `` / `` `:261-266` `` *"above"* | **half repaired, half UNMEASURED** |

`studio-room.md:388` said *"consistent with `:36` and `:261-266` **above**"* while `services/ai.py`
is cited **nowhere earlier in that document** — a corpus-internal fact needing no ref and no clone.
The dangling cross-reference is deleted.

Whether those are the right *lines* is **UNMEASURED and stays so**: studio-room lives in
`anthropos-work/anthropos-studio-room`, **not cloned on this box**, and the only copy here is the
untracked CI-pulled `stack-demo/app/studio/` tree that iter-69 established is **in no `app` commit at
any ref**. A verdict read off it would have no ref, and P1 says a number whose ref is a checkout is
not a measurement. **Recorded, not patched on speculation.**

## Phase D — the routed fence, settled by falsification (`D-M257x-70-2`)

iter-69 also routed a fence to teach the parser that a pathless backticked path is an antecedent for
a following bare `:N`. **It must not be built as designed**, on this iter's own evidence: the corpus
uses the port sense **12 times to the line sense 4**, so the rule is a **3-to-1 false-positive
generator** — §4 Trap A. And it would be wrong a third way: in `ai-readiness.md:45` the bare `:33`
binds to the path already on the line, and only *looked* pathless because a line-less path follows.

**Replaced by `FENCE-M257x-iter70-line-or-port`**, which is decidable and has zero false positives on
all twelve: *a bare `:N` is a line citation only if the antecedent resolves to exactly one file and
`N ≤` that file's length at the adjudication ref.*

**iter-69's claim that this class blocked the graded read is retracted.**

## Phase E — gates

| gate | result |
|---|---|
| five corpus guards | **all OK** — alignment · anchor · predicate (8 assertions, G8 8/8) · markdown-structure · corpus-index |
| suites | **not re-run — zero code touched.** One corpus deletion; `git status` is the evidence. iter-69's runs stand: `stack-core` 753/1F · `stack-injection` 332 OK · `dev-stack` 151 OK solo · `demo-stack` 1048/7F, all by IDENTITY |

## Close — 2026-08-04

**Outcome:** the class iter-69 routed as *"23 citations across 14 lines"* and called **a
prerequisite for the graded read** is **4 citations across 3 lines** — **twelve of the seventeen are
PORT NUMBERS** (`repos.yml` has 31 lines; the citation is `:5050`). Three of the four hold. The one
repair is a dangling *"above"* pointing at material that appears nowhere earlier in its document.
The fence iter-69 routed beside them is **falsified as designed** — the corpus uses the port sense
3× more often, so the rule would be a false-positive generator — and replaced by a decidable one.
**Rule 32 against my own hand-off, off by a factor of six, one iteration later.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (2 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-70-1` (the routed class was 23; it is 4 — rule 32 against my own hand-off,
including the discriminator's own reported defect), `D-M257x-70-2` (`FENCE-M257x-iter69-citation-antecedent`
must not be built as designed; replaced by `FENCE-M257x-iter70-line-or-port`), `D-M257x-70-3` (the
four adjudicated; `studio-room.md`'s line numbers UNMEASURED, recorded not guessed).
**Side-deliverables:** none.
**Routes carried forward:**
- `FENCE-M257x-iter70-line-or-port` — *a bare `:N` is a line citation only if the antecedent resolves
  to exactly one file and `N ≤` its length at the adjudication ref.* Replaces the iter-69 design.
- `CHECK-M257x-iter70-studio-room-lines` — `studio-room.md:388`'s `:36` / `:261-266`. **Precondition:
  a clone of `anthropos-work/anthropos-studio-room`**, which this box does not have; the CI-pulled
  in-image copy is in no ref and cannot carry a verdict.
- **Closed here:** `FIX-M257x-iter69-pathless-antecedent` (falsified — 4, not 23; 3 of 4 hold) ·
  `FENCE-M257x-iter69-citation-antecedent` (falsified as designed, replaced).
- Unchanged: `FENCE-M257x-iter68-citation-resolution` · `FIX-M257x-iter58-mainline-shift` ·
  `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
  (**NOT DECIDED**) · `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.

**Lessons:**

1. **Re-derive your OWN hand-off, not just someone else's.** Rule 32 was written about two
   orchestrator-supplied facts. It applies with exactly the same force to a route you wrote
   yesterday, and this one was off by a factor of six.
2. **Read the VALUES, not just the count.** Twelve of seventeen were ports, and one glance at
   `:5050` beside `repos.yml` said so. The count had been carried into a close section, a milestone
   ledger and a commit message before anyone looked at what it was counting.
3. **A construct that is ambiguous in the corpus cannot be fenced by disambiguating it in the
   regex.** The pathless-antecedent rule was 3-to-1 wrong. The version that works adds a
   *decidable* side-condition — `N ≤` the file's length — rather than a better guess.
4. **"Above" is a citation too, and it is the cheapest one to check.** `services/ai.py` appeared
   nowhere earlier in its own document; that needed no clone, no ref, and no tooling.
5. **Report the defect in your own discriminator in the same breath as its result.** The `max()`
   over basenames inflated one antecedent to 23,437 lines and would have quietly passed a port
   through as a citation. A recall figure produced by a flattering instrument is not a measurement.
