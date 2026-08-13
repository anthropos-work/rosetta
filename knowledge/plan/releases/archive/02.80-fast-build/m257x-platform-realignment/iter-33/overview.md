---
milestone: M257x
iter: 33
iteration_type: tik
status: closed-fixed
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-33 — clause 5, the full read

## Active strategy reference

`TOK-01: instrument first, then follow`. Clause 5 is the one clause of the five whose instrument has
never been built: it is the only one still measured by *reading*, and iter-21 proved that reading it with
the wrong instrument produces a confidently wrong answer.

## Step 0 — re-survey

- platform origin `2adcf71`, re-fetched at open — unchanged. Re-scope trigger stays at occurrence 1 of 2.
- rosetta @ `6d4641d` (iter-32), clean, 0 behind `main`.
- rext pin unchanged at `fast-build-m257x-iter-31b`; no rext change is expected this iter.
- Clause 5 is confirmed still open and still unmeasured — untouched for three runs, now the
  longest-standing clause and, with clause 2 at `27 / 3 / 1`, plausibly the last one between this
  milestone and its gate.

## Cluster / target identified

**Gate clause 5: "KB-fidelity audit GREEN, or YELLOW with 0 blockers, over `corpus/services/**` +
`corpus/architecture/**`."** 40 files, 8 451 lines.

The target was substituted for the hand-off's named next target (`CHECK-M257x-iter27-drilldown-target-
coupling`, i.e. `pt-activity-drilldown`) under the same TOK, and the substitution has a measured
justification rather than a preference: **iter-32 measured the binding suite at 4 min 50 s, not the
~40 min every hand-off had assumed.** A clause-2 fix iter can now afford to close with its own binding
read, so it no longer needs a whole session to itself — while clause 5 needs a long serial read that
benefits from being given one. The drilldown target is untouched and carried forward unchanged.

## Hypothesis

The corpus contains residual false-at-HEAD claims that **no term-scoped sweep can find**, because they
are wrong without using any of the words a sweep would grep for. Two independent pieces of evidence say
so:

- iter-21's audit reported 11 residual claims, then 5, then 2 — a curve that looks like convergence and
  was not. It was exhausting its own grep vocabulary; a **full read** then found **53**.
- Harden pass 6 found `corpus/services/studio-room.md` reading as a live pipeline for five paragraphs,
  missed by **three consecutive sweeps** because the doc never used the grepped words.

So the only instrument that can close clause 5 is a **full, top-to-bottom read of all 40 files**, with a
`wc -l` positive control per file so an unread file is reported rather than guessed.

## Method

Five read-only sub-agents, ~1 700 lines each, all reading one shared ground-truth brief
(`.agentspace/scratch/work-m257x/iter33-groundtruth.md`) **derived fresh this session** from the platform
clone at `2adcf71` and from `platform-migration-status.md` (whose services table is machine-fenced against
the platform's own `repos.yml`, so it is the authoritative merged/live/gone source).

Launched during iter-32's binding run, on otherwise-idle wall-clock — read-only, no tree mutation, and
recorded there as evidence for *this* iter rather than as work in that one.

**The grading rule, fixed before any finding was read:**

> **BLOCKER = false at platform origin HEAD *and* acting on it would misdirect real work.**

A claim true at HEAD is not a blocker however stale it feels; explicitly-fenced historical or prod-only
content is not a blocker; a merged-service doc that opens with its standing ⚠ banner is *correct*.

## The one thing the ground truth adds that no prior sweep could have had

**`graphql-wundergraph` was deleted from both `repos.yml` and `docker-compose.yml` at `2adcf71`
(2026-07-31) — mid-milestone**, and the repo is archived on GitHub. There is no router in a local stack
any more; local dev points straight at `backend` on `:8082/graphql/query`. **Every clause-5 sweep to date
predates that change**, so it is the newest drift in the corpus and the least likely to have been swept.

## Prediction, recorded BEFORE reading the reports

- The router drift will be the **largest single cluster** of findings.
- Total blockers across all 40 files: **10–25**.
- **Declared acceptable in advance:** a much *lower* count is not a success to be celebrated — given
  iter-21's history it is first evidence that the instrument is under-reading, and the response is to
  check the positive controls, not to declare clause 5 green.

## Phase plan

1. Collect all five audit reports; verify the per-file positive controls cover 40/40 files.
2. **Re-derive a sample of the claimed blockers against the platform myself** — the audits are evidence,
   not verdicts, and this milestone's standing rule is that an inherited claim is verified before it is
   acted on. Thirteen hand-offs have been refuted on re-measurement.
3. Fix the verified blockers by enumerated sweep with exactly-once anchors (the iter-22 harness shape:
   `(file, old, new)`, 0 matches and 2+ matches both fail loudly).
4. Re-run the five corpus guards + the `stack-core` suite against baseline.
5. Grade clause 5.

## Expected lift

Clause 5 **NOT MET → MET** (GREEN, or YELLOW with 0 blockers). Gate **3 of 5 → 4 of 5**.

## Escalation conditions

- Platform origin moves off `2adcf71` → re-scope trigger occurrence 2 → STOP.
- A blocker that can only be fixed by editing the platform repo → route forward. Binding.
- Findings so numerous that fixing them all exceeds the session → fix by evidence rank, re-measure, and
  report clause 5 honestly as still-open rather than quoting a partial sweep as green. This is iter-25's
  own escalation condition, and it has fired in this milestone once already.

## Acceptable close-no-lift outcomes

The audit completing with a **measured, id-level blocker list** and the fixes not all landing — the
milestone's first real measurement of clause 5 is itself the deliverable, and an honest "clause 5 is
NOT met, and here are the N blockers" beats an unmeasured claim of green.
