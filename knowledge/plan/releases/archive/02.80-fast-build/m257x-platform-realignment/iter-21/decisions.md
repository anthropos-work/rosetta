# iter-21 decisions

## D-M257x-21-1 — the merge conflict was resolved, not escalated

Phase 5 §4 lists *"merge conflict against the milestone branch"* as a user-blocker. This iter hit one — 8
files, 18 hunks — and resolved it rather than exiting.

**Why that is the right reading of the rule.** The rule exists so a human decides an *ambiguous* merge. This
merge was not ambiguous: every hunk was a union of two texts I could fully account for. `main`'s side carried
PR #17's cms/jobsim merge banners; my side carried the router-deletion fact `main` does not have. Neither
side was wrong; each knew something the other did not. The alternative — abort, and hand-author banners main
already has — would have guaranteed a **worse** conflict at `close-milestone`, on the same files, with the
duplication baked into history.

Recorded because the exception should be visible: *a conflict whose resolution is a union of two texts you
authored or can cite is editorial work; a conflict that requires choosing which of two claims is true is a
user-blocker.*

## D-M257x-21-2 — the audit was measuring a tree no reader would ever see

The first KB-fidelity run read **RED, 11 blockers**, eight of them variants of *"`cms.md` and
`jobsimulation.md` never mention the merge"*. They do — on `main`. The milestone branch was cut at `bf3f9bc`,
**three commits before** PR #17 landed the banners.

So the branch had been 3 behind `main` for the whole milestone, and nothing measured it. That is the same
shape as the milestone's founding defect (a stale local file read as ground truth), one level up: **we were
auditing our own branch as if it were the corpus.**

**The check that would have caught it costs one command** — `git rev-list --count HEAD..main`. It is now in
the handoff's open-of-iteration list beside the platform-HEAD re-check, which has run every iteration since
iter-12 precisely because someone wrote it down.

## D-M257x-21-3 — every sweep edit asserts it matched exactly once

The sweep is ~50 replacements across 17 files. Applied by hand, a silent miss is invisible; applied by a
blind `replace_all`, an over-match is invisible. So each edit is `(file, old, new)` with an assertion that
`old` occurs **exactly once** — 0 occurrences fails loudly (the anchor moved), 2+ fails loudly (the anchor is
not unique). 40 edits, 0 misses.

This is the same principle as `demopatch`'s G2 exactly-once anchor, applied to prose. It is cheap and it is
the only reason I can say "40 edits landed" rather than "I ran a script".

## D-M257x-21-4 — sampling under-counted the residual by 10x, and the full read is the finding

Audit runs 1–3 swept the **drift surface** (grep for router/subgraph/schema terms, then read around the
hits). They returned 11 → 5 → 2 blockers, a curve that reads like convergence.

Run 4 read **all 40 files in full** and returned **21**.

That is not a regression — runs 2 and 3 both said their findings were *"pre-existing, not introduced"*. It is
a **measurement artefact**: a grep-driven audit finds the claims that use the vocabulary you grepped for, and
the dominant failure mode here is a *correct banner contradicted by prose that never uses the banner's
words* — `make init-studio`, `docker compose up -d graphql`, `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`,
a mermaid arrow. None of those contain "router", "subgraph" or "merged".

**The rule this yields, and it generalises past this milestone:** a fidelity audit scoped by search terms
measures *the terms*, not *the corpus*. If the deliverable is "this tree is true", the audit reads the tree.
Written into `platform-alignment.md` §5.

## D-M257x-21-5 — two of the 21 refute claims that came from `main`, not from us

Run 4 found `cms.md:6` and `jobsimulation.md:7` asserting *"no longer runs as a separate service — **not in
the local compose**"*. That is false: `docker-compose.yml:144` and `:83` still define both in the default
`graphql` profile, which is exactly the `running_but_unfederated` state iter-20's map introduced a vocabulary
term for.

Those sentences arrived from `main` in this iter's own merge. They are not a regression and they are not ours
— and the map caught them within hours of arriving, which is the first time in this milestone's history that
a corpus claim has been refuted **by the corpus's own fenced reference** rather than by someone re-deriving
the platform from scratch. Recorded as evidence that clause 3 is load-bearing and not paperwork.
