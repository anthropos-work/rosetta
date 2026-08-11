# iter-249 decisions

## `D-M257x-249-1` — a frozen subject for `stack-core` is a CLONE, never an export

`git archive` at a sha gives a tree with **no `.git`**. A large part of this suite reads history
(`file absent at <sha>^`, answer-key replays of real commits, `fence_provenance`), so an export makes them
refuse — correctly — and the section reads **68 failed**, none of which is about the corpus.

**Decision:** freeze with `git clone --local --shared <repo> <dst> && git -C <dst> checkout <sha>`, for both
repos, laid out so `<dst>/rosetta/.agentspace/rosetta-extensions` mirrors the live path (the guards resolve
their subject as `Path(__file__).resolve().parents[3]`). Measured: **0.33 s, 54 MB** for 172 MB of combined
history. Recorded here rather than as a script because it is two commands and a layout rule; a script would
need its own fence.

## `D-M257x-249-2` — a test that needs untracked local state DECLARES it; skip is right and fail is wrong

The alternative was defensible and was rejected: leave the tests failing, on the grounds that a skip can
hide a regression. Three reasons the skip wins here.

1. **The subject is absent, not wrong.** `frozen_expectation_census` already raises
   `DerivationUnavailable` — *"a zero from this census would be vacuous (§9)"*. The guard refuses; only the
   test converted the refusal into a verdict.
2. **`guard_family` already reports this same requirement as NOT-RUN.** Two consumers of one precondition
   were giving two different answers. They now agree.
3. **The failure text accuses the wrong party.** `a fenced command names a target that does not exist` and
   `the live corpus must resolve every rext path` are statements *about the corpus*, printed because a
   clone is missing. That is worse than silence.

The skip cannot hide a regression **on a box that has the clones** — verified: live run is 116 passed / **0
skipped**. It only fires where the check was never possible.

## `D-M257x-249-3` — a reading records its failure NAMES, not just its count

`PR-2` could not be graded because iter-248 booked *"16 failed"* with no node-ids. Two readings a day apart
are therefore not comparable, and the question `ROUTE-M257x-248` existed to answer — *which of these were
races?* — is now permanently unanswerable for that run.

**Decision:** every suite reading booked in this milestone records the failing node-ids (or a path to the
log that holds them). Routed as `FIX-M257x-249-readings-record-names` so it reaches the protocol doc, not
just this iter.

## `D-M257x-249-4` — the 23 unrepaired failures are routed as ONE class, not thirteen items

They share one cause, one detection method (run the suite against a clone pair; diff against a live run)
and one repair shape (declare the precondition). Splitting them into per-file routes would lose the fact
that the class is **still being manufactured** — 3 of the 15 files were authored this week — and that the
right terminal state is a **fence**, per `TOK-08`, not fifteen edits.

## Side measurement kept although it refuted its own hypothesis

`anchor_offset_guard.citations()` resolves against three different populations depending on how it is
called: working-tree mode globs **every** `*.md` under the repo root (**15,258** on this box, including
`stack-dev/`, `stack-demo/` and `.agentspace/scratch/`), a clean clone sees **2,864**, and `at_rev` mode
uses `git ls-tree` (tracked only). The hypothesis was that this makes the guard's verdict
operator-dependent. Run all three ways at `971cdc4`: **26 targets / 60 citations / 0 ambiguous — identical,
with an empty symmetric difference.** The hazard is latent, not live. Routed
(`ROUTE-M257x-249-anchor-offset-has-three-populations`), not repaired.
