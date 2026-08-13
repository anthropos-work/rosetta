# Release Retro: v2.8 "fast build"

**Shipped:** 2026-08-13 · **Tag:** `v2.8` · **Milestones:** M255 · M256 · M257x · M257 · M258
**Thesis:** *from nothing, to live, to provably live, fast.*

## What shipped

A cold demo bring-up went from **450 s to 287 s** on the development host — beating the 360 s gate and its
300 s stretch — and a bring-up now **drives the full journey suite to completion** and leaves the demo
presenter-usable. **11.54 GB** of disk was returned at **zero build-time cost**, with the build cache
deliberately untouched. **Zero platform-repo edits and zero net-new third-party dependencies** across all
five milestones, both verified rather than asserted.

## The headline number, stated honestly

The roadmap promised **666 s → ≤ 360 s**. What was achieved is **286.99 s against `macmini`'s own 449.51 s
baseline — a 36.2 % cut, not the 46 % the framing implies.** `billion`'s 666.29 s was measured on different
hardware and **does not transfer**; M257 re-pointed its gate to the host that exists and met it there. Both
statements are true; only their combination would be misleading, and this release's own standing rule —
*state the environment with every number* — forbids that combination.

**Clause 3 of M258 was never met.** The composed p50 over three cold cycles was measured at **840.01 s with
all three reps instrument-rejected** (peak load1 40/75/52 against a limit of 10). **401.60 s is a projection**
composed from separately-measured halves. A **~290 s warm-cache cycle was deliberately not banked** — it was
too good precisely because it skipped the image-export leg that is 46 % of a cold build. The user ruled the
milestone achieved on its other four clauses; that is a narrowing of the definition, **not a gate that fired**.

## Incidents and defect classes

**P0 — a silent-403 class hid for an entire milestone.** Platform `766df6c` folded **sentinel into `app`**
(the 8th such merge) while three post-seed reload sites still drove the deleted container's RPC and logged
the miss as *"non-fatal"*. That comment was false: an un-reloaded enforcer serves its boot policy forever.
Seeders write grants in raw SQL, off casbin's write path, so **every org-scoped operation refused
`forbidden` at HTTP 200**. Fifteen journeys failed. **The partition was the proof** — all 15 failing were
org-scoped, all 15 passing user-scoped. `batch_seconds` fell **629 → 129**: the suite had been slow *because*
it was broken.

**Recurring class 1 — a fence satisfied by its own comment.** Three instances, the last found at this close,
**four lines from an already-exec-scoped sibling** and in the same file a harden pass had just fixed. The
remedy (`shellInvocationLines`) existed throughout. **No detector was ever built** — carried forward.

**Recurring class 2 — a routing is not a routing until the target's own doc says so.** Fired at least three
times (`BIND_HOST`, M257x's carry-forward, M255's four items to M257), each caught by the *next* close and
never by the routing one. Carried forward as a proposed fence.

**Recurring class 3 — an instrument inside its own subject measures itself.** Nine-plus instances: a test
environment built inside the tree it audits; a benchmark resolving paths relative to its own location; a
census writing its probe into the directory it counts; a restore reading one file from the live stack and
two from a stale clone while printing success.

**Fail-open to GREEN.** A reader graded on what it *printed*, so any crash read as "no failures". A hook that
short-circuited, so a previous run's verdict file — which `--purge` does not clear — was read as the current
one. Both found at M258's close, in code written that same milestone.

## Process findings

- **A code read is not a measurement.** A chronic item three iterations had "verified" by reading was
  refuted by measuring it once.
- **Never read an exit code through a pipe.** `cmd | tail; echo $?` reports the *tail's* status — made twice,
  once nearly reporting a working fence as broken.
- **A column can answer a different question than the one you asked it.** `docker images` SIZE overstates
  reclaimable space ~5× by billing shared layers to every referencing image.
- **A token census finds a wrong value, never an absent one.** This close's own 33-file "live sentinel" grep
  and two of seven flagged doc sites were **false positives** — correct prose describing history. Not edited.
- **An instrument cannot resolve a signal finer than its own sampling interval.** Two attempts missed a
  75-second window while polling every two minutes.

## Metrics

Go test functions **2019 → 2233** (+214, the largest single-release Go growth on record) · TS specs
**292 → 359** · Python demo-stack **910 → 1150**, stack-injection **258 → 343**, stack-verify **154 → 281**.
Coverage: `seeders` **95.7 → 94.4** (−1.3pp, in tolerance; **third consecutive fall**, −1.7pp cumulative —
recorded so it is not re-derived). Flakes **0**.

**Not measured, and stated as such:** the full `stack-core` sweep **has never completed** — it reached 892 of
2419 at ~2 tests/min, with 13 failures already accrued and the census module's own 11 sorting *after* the
slow region. **The full-sweep total is ≥ 24 and has never been established.** Benchmark regression is
**not gradeable** at this close.

## Carry-forward

Named item by item in [`roadmap-vision.md`](../../roadmap-vision.md) § *v2.8 "fast build" carry-forward* —
the destination did not exist until this close, which is itself the finding. Led by
**`FIX-M256-studio-false-green`**, a verified-open false green inside the suite this release's headline rests
on, and **`BIND_HOST` / `D-M255-7`**, which is why the batch gate skips on the default `/demo-up` path.

## Stats

Phase 8c snapshot not taken — the session's subagent budget (300) was exhausted during the review sweeps and
`/developer-kit:project-stats` was not run. Recorded as a gap rather than a silent omission.
