---
iter: 256
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
controlling_strategy: TOK-08
---

# iter-256 — measure the advance before taking it

**Type:** tik · **Active strategy:** `TOK-08` (census the mechanical classes; stop sampling them).

## Why this target, and why now — the user's closing condition

The user set a **binding closing condition on 2026-08-10** that supersedes "gate 4 of 5" as the
definition of done:

> the milestone closes ONLY when the CURRENT `main`/tagged branches of the still-relevant platform
> repos assemble into a WORKING STACK — BOTH demo AND dev — and the corpus reflects that, with the
> deprecated/removed repos no longer treated as part of the project.

**The direct consequence: advancing the active clones onto current platform code is now IN SCOPE AND
REQUIRED.** It was routed and deliberately not done for many iters precisely because it changes what a
demo builds. Under the user's condition, a stack built from a stale pin does not count.

## What the protocol says must happen FIRST, and it is not the advance

`corpus/ops/platform-alignment.md` §7 rule 4 and its sub-rules are directly on this target and they
order it:

- **4c — you do not have to TAKE an advance to measure it, and you should measure it first.**
  `git show <ref>:<path>` reads any ref without touching a checkout. iter-68 measured an `app` advance
  of 56 commits and found **25 of 42** corpus citations holding at the pinned ref broke at origin
  HEAD — **60 % in one working day** — and that fifteen minutes of measurement redirected the whole
  iteration.
- **4 — before taking an advance, re-resolve every corpus citation whose path lands in the advancing
  repo, at the new ref, and record moved / dead / held.** The repair belongs to the **advancing iter**.
- **4d — a guard that resolves a citation against a CHECKOUT has no ref, and its verdict is not a
  measurement.** The instrument that answers this already exists: `CITE_REF` (M257x iter-68), read by
  `anchor_construct_guard`, `platform_alignment_guard` and `platform_predicate_guard`.

So this iter's planned scope is the **measurement**, taken through the instrument the milestone already
built for exactly this, and a decision — sized against the measured volume — on whether the repair and
the physical checkout advance land here or in the quiet-window iter that re-proves clause 1.

## Cluster / target identified

`TOK-08` names no next target beyond "work the mechanical classes in descending measured size"; the
user's closing condition names this one, and it is itself a mechanical class — **a citation either
resolves to what it claims at a named ref or it does not.** This is the same census discipline applied
to a *ref delta* rather than to a corpus population.

**Re-survey (mandatory Phase 1 Step 0), taken 2026-08-10 09:16 CEST, by real `git fetch`:**

| repo | HEAD | origin/main (fetched) | behind |
|---|---|---|---|
| `platform` | `0c91421` | `0c91421` | **0** |
| `app` | `ad9f3c498` | `3eaadae68` | **28** |
| `next-web-app` | `8297c684c` | `19423a1fb` | **12** |
| `ant-academy` | `22df69dd` | `249430c3` | **10** |
| `sentinel` | `f2c4619` | `f2c4619` | **0** |
| `studio-desk` | `41ee3575` | `41ee3575` | **0** |

**The fetch already moved one number**: the orchestrator's brief said ant-academy was **+9**, from a
cached remote-tracking ref; the real remote says **+10**. *A remote-tracking ref is a cache, not a
remote* — demonstrated live, in the first minute of the iter, on the very repo set the advance is about.

## Hypothesis

The advance is **not** citation-neutral, and its damage is **much smaller per-commit than iter-68's**,
because iter-68's 56 commits were the `ai`-fold landing in one morning and these 50 are four days of
ordinary work. The measurement is what decides whether the repair + advance fit in this iter.

## Pre-registered claims — sealed in this iter's FIRST commit, before any measurement

Each is stated as a CLAIM plus my PREDICTION, kept separate (iter-255 Lesson 2 — inverting them is a
one-word error that flatters the author).

- **PR-1 — CLAIM: the advance breaks corpus citations at a rate comparable to iter-68's 60 %.**
  PREDICTION: **REFUTED.** Fewer than **20 %** of the citations that resolve into `app` and hold at
  HEAD will fail at `origin/main`.
- **PR-2 — CLAIM: the advance is citation-neutral (zero corpus citations that hold at HEAD fail at
  origin/main).** PREDICTION: **REFUTED — at least 1 fails.** Predicting against a free advance.
- **PR-3 — CLAIM: commit count predicts citation breakage — the repo with the most commits in the
  advance (`app`, 28) also has the highest FRACTION of its citations broken.** PREDICTION: **REFUTED**;
  the ordering by fraction will not match the ordering by commit count.
- **PR-4 — CLAIM: the delta runs both ways — at least one citation is UNRESOLVABLE at HEAD and
  RESOLVES at origin/main (a file born in the advance).** PREDICTION: **REFUTED — zero such
  citations.** 50 commits over four days is small for a net-new cited file.
- **PR-5 — CLAIM: measuring at `CITE_REF=origin/main` turns at least 2 guards RED that are GREEN at
  the default.** PREDICTION: **HELD — at least 2 go RED.**

## Phase plan

- **Phase A** — establish the live baseline at the default `CITE_REF` (the checkout), naming the
  runner, the scope and the language.
- **Phase B** — re-run the ref-aware instruments at `CITE_REF=origin/main`; partition the delta into
  **held / moved / dead / born**.
- **Phase C** — size the repair. If it is completable in-iter, repair; if not, route it with a named
  handler and say so with the number.
- **Phase D** — grade the five pre-registrations; record what the advance would cost.

## Escalation conditions

- If an instrument states its own invalidity (an unresolvable named ref, a collapsed denominator), it
  is **UNMEASURED**, never silently substituted (§5 rule 7) — report it as such rather than reading it.
- Taking the physical checkout advance while `demo-1` is up (3 days) desynchronises a running stack's
  build source from its images. That is **not** a user-blocker — it is a sizing decision this iter
  makes against the measured volume and records.

## Acceptable close-no-lift outcomes

A measurement that comes back **zero broken citations** is a complete iter: it would establish that the
advance is citation-safe and hand the quiet-window iter a cleared precondition. The falsification has to
be *measured*, not assumed — an instrument returning zero must prove it can fire.
