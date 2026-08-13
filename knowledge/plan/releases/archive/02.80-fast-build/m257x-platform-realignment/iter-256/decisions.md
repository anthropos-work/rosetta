# iter-256 — decisions

## `D-M257x-256-1` — the user's closing condition is recorded as binding, and it re-orders the milestone

**Supplied by the user on 2026-08-10, and it supersedes "gate 4 of 5" as the definition of done:**

> the milestone closes ONLY when the CURRENT `main`/tagged branches of the still-relevant platform
> repos assemble into a WORKING STACK — **BOTH demo AND dev** — and the corpus reflects that, with the
> **deprecated/removed repos no longer treated as part of the project**.

**Three consequences, each of which changes what an iter may target:**

1. **Advancing the active clones onto current platform code is IN SCOPE AND REQUIRED.** It had been
   routed and deliberately not taken for many iters, on the ground that it changes what a demo builds
   and `demo-1` was up and green with clauses 1–2 proven on it. Under the user's condition that
   reasoning inverts: **a stack built from a stale pin does not count.**
2. **The working-stack proof must cover DEV as well as DEMO.** Every working-stack proof this milestone
   has ever attempted was demo-only. Treat "dev assembles too" as genuinely unproven — iter-243 already
   found the seeder restarting a *demo* container on dev stacks.
3. **"Deprecated repos no longer part of the project" is a census obligation**, not a one-off. The
   phantom removal from the clone pin (5 repos: cms, jobsimulation, storage, messenger, roadrunner) was
   one instance; where the corpus still treats removed repos as live is squarely in scope.

**Not decided here:** whether the physical checkout advance lands in this iter. That is sized against
the measurement, per §7 rule 4c — *you do not have to TAKE an advance to measure it, and you should
measure it first.*

## `D-M257x-256-2` — the fetch is part of the measurement, and it moved a number in the first minute

The brief carried `ant-academy` at **+9** commits. A real `git fetch origin` against the six clones
made it **+10** (`c885dab2` → `249430c3`), while the other five were unchanged. **A remote-tracking ref
is a cache, not a remote** — stated in the brief as a rule and demonstrated on the very repo set the
advance is about, before any other work.

**Standing consequence for this milestone:** any advance arithmetic (`+28 / +12 / +10`) is a reading
with a timestamp, never a property. Every figure in this iter names the fetch that produced it.

## `D-M257x-256-3` — the corpus is a TWO-CLOCK document, and the second clock is our own tooling

Measured, not inferred. `platform-migration-status.md:121`'s seven `app` anchors are exact at
`origin/main` (`3eaadae68`) and wrong in the checkout (`ad9f3c498`); `observability.md:29`'s anchor is
exact in the checkout and wrong at `origin/main` — **one row below a sibling that is `origin/main`-framed.**

**The mechanism is `anchor_construct_guard`'s `CITE_REF=auto` ladder**, which tries `origin/main` before
`HEAD`. That default is right for *grading* — the exit gate names origin HEAD — and it means every anchor
derived or verified through the guard is an `origin/main` number, while every anchor derived by reading a
checkout is a checkout number. Nothing in the corpus marks which, and no fence can tell them apart: both
land on real lines, so the wrong-construct floor `anchor_construct_guard` discloses covers exactly this.

**Decision: the census REPORTS the delta and refuses to attribute it.** `--apply` now requires
`--adjudicated <file>` naming the corpus sites an operator has read. This is `D-M257x-122-5` (a bare
basename is refused, never resolved by proximity) applied unchanged to **refs**: guessing between two
candidates that equally satisfy a citation is the wrong-construct error the census exists to find.

**Why it is recorded as a decision and not a bug:** the rewriter is correct — its inverse restored all 27
sites byte-for-byte against `git diff`. What was wrong was a *premise stated nowhere*, which is the class
this milestone exists to remove.

## `D-M257x-256-4` — the advance is the PIN, and taking it is booked as an unproven state change

> ⚠️ **AMENDED AT iter-257 — the paragraph below was too strong and the strong half is RETRACTED.**
> Measured at iter-257: `DEMO_ADVANCE_CLONES` defaults to `0` (`ensure-clones.sh:212`) and **no other
> file in rext sets it**, so a default bring-up applies **no pin at all** — a fresh box gets each repo's
> default-branch tip from `git clone` + `make init`, and an existing workspace builds whatever its
> clones are checked out at. The pin is a **reproducibility barrier available on request**
> (`DEMO_ADVANCE_CLONES=pinned`), and it reaches the clones through the **workspace copy**, not the
> canonical file. So *"the next `/demo-up` would have re-pinned them back"* is **false** — nothing
> would have. What is true, and is why iter-256's action was still right, is that advancing the
> **checkouts** is what a default bring-up actually consumes, and advancing the **canonical pin** is
> what `pinned` should mean when someone asks for it. Both were needed; neither is the sentence below.

`git merge --ff-only` in three clones changes what is on disk. `rosetta-extensions/demo-stack/clones.pin.json`
changes what a `DEMO_ADVANCE_CLONES=pinned` bring-up builds — `ensure-clones.sh` checks each clone out at
the ref it names. Advancing only the checkouts would have left that mode pointing at the old refs.

Both were advanced (`app` `3eaadae68`, `next-web-app` `19423a1fb`, `ant-academy` `249430c39`), and
`clone_pin_guard` stays GREEN.

**Booked honestly: this is an UNPROVEN state change.** No bring-up was run — the box is not quiet and
`demo-1` has been up three days. The `app` advance carries a **new migration**
(`terraform/migrations/20260804160000_assignment_notification_logs.sql`) and a terraform fix whose own
subject says *"the backend migration pipeline has been a silent no-op since the atlas 0.7.0 bump"*, so
`make migrate` is the first thing the proving iter must watch. It is reversible: three shas in one JSON file.

**What justifies taking it unproven** is `D-M257x-256-1`: under the user's closing condition a stack built
from a stale pin does not count, so the pin had to move before the proof could mean anything. Taking it
also **collapses the two clocks of `D-M257x-256-3` into one** — the residual mis-framed anchors become
measurable at a single ref for the first time.

## `D-M257x-256-5` — a recorded reason that quotes an arrow collides with the ratchet's arrow grammar

`TEST_MODULE_LITERAL_CEILING`'s block is graded by reading the **last** arrow target as the new ceiling.
The reason written for this iter's re-pin quoted the applier's cascade hazard as a pair of renumbering
arrows, so the arm read `530` against a constant of `636` and reported the block contradicting itself.

**Decision: quote a mapping in words inside a ratchet block**, and record the collision in the block
itself rather than silently rewording. Same family as iter-255's *"the COMMENT arrow breached the ceiling
it was written to raise"* — a ratchet's prose is inside the population the ratchet measures, and its
grammar is inside the population its own arm parses.

## `D-M257x-256-6` — SIDE DISCOVERY: the mutation battery has been UNREADABLE to pytest from the rext root since iter-255

Found by this iter's whole-section run — **the first one since iter-252** — and it is not this iter's
defect. `test_m257x_mechanical_fences_mutation_battery.py:49` gained `import suite_census as census` at
iter-255 (`4aadd60`), on a `sys.path` that only carries `tests/`. `suite_census` lives one level up in
`stack-core/`, so the import resolves **only when pytest is invoked from `stack-core/`** — which is how
a human runs it, and is **not** how `suite_census.collected_by_pytest` runs it. From the rext root the
module raised `ModuleNotFoundError` and was recorded **`UNREADABLE (-1)`** against **6** under unittest,
which is the whole of the aggregate gap the census reported (`-1 - 6 = -7`).

**The severity is the subject, not the size.** That module is the harness carrying ~30 of this
milestone's mutation proofs. For four iters it was invisible to one of its two runners, and the
instrument that would have said so — `test_suite_census_collection.py::LiveCollectionReach` — was
itself only run as part of a whole-section sweep that no iter performed. iter-255's own Lesson 1 said
*"a repair can invalidate a proof that was not part of it"*; this is that sentence landing on the
commit that wrote it.

**Decision: fix at the import, not at the invocation** — one `sys.path.insert` of `parents[1]`, with the
failure mode recorded in place so the next person who "tidies" a redundant-looking path insert reads why
it is there. Verified in both directions: the module collects **6** from the rext root and from
`stack-core/`, and `test_suite_census_collection.py` goes **3 failed / 14 passed → 17 passed**.

**It also converts a routed item into a measured one.** `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet`
was a hypothesis that the loop's scoped runs miss things; this is a named, dated instance with a cost
(4 iters of an invisible harness) and it is the second witness that route asked for.
