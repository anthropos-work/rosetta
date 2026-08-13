# iter-183 — decisions

## `D-M257x-183-1` — the milestone's own BACKLOG is a registry, and it was the only one with no fence

**Measured, not argued.** 183 routes carry at least one recorded disposition across `m257x`'s 164
`**Routes carried forward:**` blocks — **779** (route × iter) dispositions in total, **34** of them
closures. Every one has been maintained by eye. `stack-core` ships fences over the corpus, the
platform's config, the demo knobs, the anchors, the retractions, the derived counts, the guard family
itself and the guards' own tests — and **zero** over the queue that decides what the next iter works on.

**Why that is not a cosmetic gap.** The queue is the substrate every sub-agent on this milestone is
briefed from. The consumer noticed before any instrument did: run 17's orchestrator brief listed
`FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` as open with the hedge *"may already
be closed by iter-180 — verify before working it."* It was closed. A brief that hedges is a brief whose
reader spends the first minutes of an iter re-deriving the backlog.

**Decision:** fence it. `route_disposition_guard.py` asserts one property — a route recorded `CLOSED` at
iter K may not be asserted **open** at iter M > K unless the later block says why — sweeps every planned
milestone, and is registered in `guard_family`'s `INVOCATIONS`, which makes it un-droppable in both
directions (`guard_family.py:303`).

## `D-M257x-183-2` — the one live contradiction is a CORRECTION, and that verdict was measured, not assumed

`SURVEY-M257x-iter170-cockpit-runner-dependence` was **CLOSED BY ROOT CAUSE at iter-171**
(`D-M257x-171-1` — not a harness assumption but a blocking `socket.getfqdn` in shipped `cockpit.py`,
35.005 s on the 3.14 interpreter against a 2 s bind window). **iter-182 re-published it as *"unchanged;
open"*** eleven iters later, describing the very mechanism iter-171 had attributed and repaired.

**A stale carry-forward and a legitimate re-open are the same sentence**, so the disposition was settled
by re-running the subject rather than by reading the record:

| runner | `test_cockpit` |
|---|---|
| `unittest` / `/usr/bin/python3` 3.9.6 | **207 run, 0 failures** |
| `unittest` / `python3.14` 3.14.6 | **207 run, 0 failures** |
| `pytest` / `/usr/bin/python3` 3.9.6 | **207 passed** |

`207 = 207 = 207`. iter-171's repair holds; there is no live disagreement to route. So the repair is an
**in-place `✅ CORRECTED — iter-183` annotation** on iter-182's bullet — never a rewrite of a closed
iter's record, and never a re-open the evidence does not support.

## `D-M257x-183-3` — a crude reader returns SIX findings here and FIVE are false; the grammar is the deliverable

The first detector attached a disposition to a whole BULLET. It returned 6. Adjudicated one at a time:

| candidate | verdict |
|---|---|
| `SURVEY-M257x-iter170-cockpit-runner-dependence` | **TRUE** — closed iter-171, open iter-182 |
| `FIX-M257x-iter145-sha-baseline-drift` | false — iter-149's `CLOSED by this iter` belongs to `SURVEY-M257x-iter146-*` in the same bullet |
| `SURVEY-M257x-iter163-anchors-with-no-quoted-literal` | false — same shape; the closure is `FIX-M257x-iter138-anchor-rot-fence`'s |
| `FIX-M257x-iter138-anchor-rot-fence` | false — same shape; the closure is `FIX-M257x-iter135-bare-pin-blind-spot`'s |
| `FIX-M257x-iter27-succession-hero-not-rendered` | false — the disposition is a TABLE row, and the verdict is *superseded* |
| `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` | false — iter-173 wrote **half CLOSED** and iter-174 scoped to *"the observed half"* |

**1 of 6.** Booking the raw 6 would have been `D-M257x-122-3`'s error in miniature: a hunted count is a
statement about the hunt. The grammar-aware detector, written to the tokens the registry *already
writes*, returns exactly 1 — the same 1 the hand adjudication reached, independently.

**The transferable half:** the registry was already writing a grammar (` · ` segments, `half CLOSED`,
*supersedes*, tables-are-not-bullets). Nothing had ever read it. Giving a fence the registry's own
grammar cost less than rewriting 164 blocks into a machine format, and it left the prose readable.

## `D-M257x-183-4` — a multi-id closure is REFUSED, never resolved by proximity — and the docstring that said otherwise was written before it was measured

The first draft of the guard's docstring asserted the live registry never writes two ids in one segment.
**Measured: 28 multi-id segments, 2 of them carrying a closure verdict.** The claim was false and it was
made the way this milestone has caught itself making claims eleven times: stated because it was
plausible, in the same commit as the instrument that could have checked it.

A proximity-resolver books both ids and — here — happens not to fire. That is luck, not a fence. So the
guard applies `D-M257x-122-5` at segment grain: a multi-id segment carrying a closure is **ambiguous**,
closes nothing, and **the refusal count prints on every run**. The cost is stated rather than hidden — a
genuine closure written that way does not register — and widening the grammar is routed, not assumed.

## `D-M257x-183-5` — the anti-vacuity control names its SCOPE, because an unstated scope is how a control goes vacuous

Zero closures across the sweep is exit 2, never 0: a broken parser returns the same zero a clean registry
does. But a *single* milestone may legitimately close nothing — `m256` and `m257` both do — so the
control is scoped to the **default sweep**, the invocation `guard_family` makes, and says so in its own
comment. A `--milestone-dir` caller names its own subject and owns its own floor; the live-registry
ratchet in the tests is that floor for the real registry. This is `§9`'s *grade at the grain of the
claim*, applied to a control rather than to a reading.

## `D-M257x-183-6` — adding one fence moved a PUBLISHED TRIPLE, and the same pair now means two things

Two registry guards went RED on the new fence, both correctly, and both are the cost of adding a fence
rather than a defect: `test_fence_provenance` (the guard did not stamp the tree its configuration lives
in) and `test_fence_registry_population_m257x` (the README fence-index disclosure is a **measurement**,
so it moves when the population does). Repaired: the `stamp_main()` call, and the triple
`16/27 · 16/26 · 15/26` → **`17/28 · 17/27 · 16/27`**.

**And in moving it, the docstring acquired a collision.** `16 of 27` now appears twice in that one file
meaning two different things: a **`union`** reading at `5b108d0` (the historical figure harden pass 39
retracted and iter-177 un-retracted) and a **`declaring`** reading at HEAD. Nothing separates them but
the label and the ref. That is iter-177's rule landing in the file that already knew it — so both sites
now name the derivation **and** the ref, and the collision is stated in the docstring rather than left
for the next reader to trip over.

**Booked, not silently absorbed:** `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` names a figure
this iter moved. The route is **not** renamed — iter-177 settled that the id embeds a correct reading —
but a reader arriving at it now finds a triple that no longer matches, so the route's live figure is
restated in this iter's close.
