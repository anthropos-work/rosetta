# M258 — spec notes

_Technical notes accumulate here during the build._

## Pre-flight audits — iter-01

**`/developer-kit:audit-kb-fidelity --milestone=M258` — verdict `YELLOW`**, 2026-08-12.
Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md). No blind areas; every topic this milestone
depends on has a corpus doc. Two classes of finding, both applied inline, plus two measurements the
plan did not carry (below). The bootstrap tok's strategy is authored against the corrected state.

## Topic → doc → code triples (fast start for later audits)

| Topic | Corpus doc | Code (in `rosetta-extensions`) |
|---|---|---|
| bring-up verification / autoverify | `corpus/ops/verification.md` | `stack-verify/live/autoverify.sh`, `stack-verify/live/verify.sh`, `stack-verify/lib/{services,target,readiness}.sh` |
| the Playthrough batch | `corpus/ops/demo/playthroughs.md` | `playthroughs/e2e/run-playthroughs.sh`, `playthroughs/manifest/*.yaml`, `playthroughs/cmd/{ptreport,ptvalidate}` |
| the world contract (reset vs presenter) | `corpus/ops/idempotency.md`, `corpus/ops/demo/playthroughs.md` | `stack-seeding/cmd/stackseed/main.go` (`resetTables` `:44-131`, `TRUNCATE` `:839`), `demo-stack/up-injected.sh` (`STORIES_PRESET` `:254`/`:261`) |
| the build budget / p50 gate | `corpus/ops/demo/build-budget.md` | `demo-stack/buildbench.py`, `demo-stack/hostprofiles/*.json` |
| host topology / public-host | `corpus/ops/demo/tailscale-serve.md` | `playthroughs/e2e/run-playthroughs.sh:92-105` |
| the content-stories sweep | `corpus/ops/demo/coverage-protocol.md`, `.../content-stories-spec.md` | `stack-verify/e2e/run-content-stories.sh`, `stack-verify/e2e/lib/content-pairs.ts`, `stack-verify/e2e/content-denominator.json` |

## ⚠️ Two measurements the plan did not carry (surfaced by the iter-01 audit)

Neither is a stale claim — no doc asserts a wrong number. Both are numbers that **exist in a
predecessor's milestone record and reached neither this plan nor the corpus**, and both bear directly
on whether this milestone's gate is gradeable as written. Recorded here so the bootstrap tok plans
against them rather than rediscovering them at the first campaign.

### 1. The batch half of the composed budget has never been published

This `overview.md` § *Budget honesty* asks M256 to *"measure and report the reset-to-seed leg, so this
composition arithmetic has a third real number instead of two."* **It was not reported.** Neither
`playthroughs.md` nor any other corpus doc publishes a suite wall-clock or a reset-leg figure, and
M256's `Gate Outcome Ledger` grades clause 1 on a **per-test median**, not a wall-clock.

What does exist, in M256's `progress.md` and nowhere else:

| figure | value | provenance |
|---|---|---|
| suite wall-clock median | **56.6 s** | iter-02, n=3, local `demo-2`, **18 specs** — *"reported, not gated"* |
| median per non-studio Playthrough | 3.326 s → **2.014 s** | iter-02 → iter-03 |
| specs at close | **209 passed** | close, 3× cold reset-to-seed |

⚠️ **The 56.6 s is not the shipped suite's number.** It was measured over **18** specs; the suite
closed at **209 passed** across **30 live Playthroughs**. Quoting 56.6 s as the batch half would be
quoting a figure for a suite that is an order of magnitude smaller than the one M258 must run. The
composition therefore has **one measured half (286.99 s bring-up, M257 iter-09) and one unmeasured
half** — measuring the batch half on this host is the first thing a tik must do.

### 2. M256 escalated that its own suite timing is NOT decidable at n=3 on this host

M256 `progress.md` iter-12, verbatim headline: *"clause 1 is **NOT DECIDABLE at n=3 on this host**."*
Six full-suite runs on the same box, the original 16 specs unchanged, and the control subset spanned
**0.5281× → 1.0762× — a 2.04× spread with no trend** (newest 0.529×, oldest 0.528×, the extreme in
between). The measurement stood; the *gate* was re-cut (`D-v28-12` → `D-v28-13`) away from a relative
speed target precisely because the threshold sat **inside its own noise floor**.

**Why this is load-bearing here and not merely historical.** M258's gate is a **p50 over n=3** whose
second half is that same suite on that same class of host. A 2.04× spread on the batch half is not
obviously survivable by three samples, and the lesson M256 published alongside it is the one that
applies: *a relative gate needs its noise floor published next to it or it is not falsifiable.*

**This is evidence, not a re-cut.** It is recorded, not actioned — the gate is unchanged. What it
obliges is that the first campaign **publish the batch half's spread next to its p50**, so the tok
can tell "we are inside the gate" from "we sampled favourably", which is the distinction M256 spent
eleven iters refusing to blur. If the spread turns out to make n=3 undecidable, that is a
user-facing renegotiation with measurements attached, and the milestone's own
`re_scope_trigger` is the declared valve.
