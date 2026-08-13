# iter-16 — decisions

## D-M257x-16-1: a documented degradation and an unclassified error do not share a verdict

**Decision.** `dev-setdress.sh` splits the replay outcome three ways rather than two: rc=4 (stack
unprovisioned) and rc=5 (cache miss) stay **`skipped(...)` + verdict `set-dressed` + exit 0**; any other
non-zero rc becomes **`FAILED(error)` + verdict `set-dress INCOMPLETE` + exit 3**.

**Why not "any replay failure is a failure".** Because it would be wrong, and would train operators to
ignore the new signal. rc=4 and rc=5 are states the environment legitimately produces (a stack that was
never provisioned; a box whose cache has nothing for this release), each already printing a named operator
fix. Calling them failures makes the common case noisy and the rare case invisible — Trap A's
tune-until-it-catches-nothing, arrived at from the opposite direction.

**Why not exit 1.** `die` already exits 1 and means *aborted — nothing seeded*. Reusing it would erase the
distinction that makes the new code safe to ship: the seed floor **did** run, so a demo built this way still
logs in, and the M13/M18 contract that a bring-up is never aborted by its set-dress pass is intact. 3 says
*completed, with a named surface missing*. A caller can tell the two apart; a boolean cannot, which is
exactly why the `if !` wrapper had to go.

**Alternative considered and rejected:** having `up-injected.sh` inspect `SNAP_SUMMARY` textually for
`FAILED`. It keeps the engine's exit code at 0 and puts the classification in the caller — i.e. a second
place that must know the engine's vocabulary, which is §2's *derive it at the point of use* violated by
construction. The engine knows what happened; it should say so in the channel callers already read.

## D-M257x-16-2: three tests re-pointed, one deliberately left at exit 1

**Decision.** `test_cache_miss_is_non_fatal_seed_still_runs` and
`test_demo_seed_is_the_atomicity_floor_after_a_replay_miss` move from `stacksnap exit 1` to `exit 5`;
`test_replay_error_is_tolerated_seed_is_the_floor` is renamed
`test_replay_error_still_runs_the_seed_floor` and keeps exit 1 with its expected rc changed 0 → 3.
`test_engine_never_captures_even_when_the_replay_misses` **keeps exit 1**.

**Why.** The first two have "miss" in their names and were testing the miss with the code for
*not-a-miss* — a fix16 leftover their own comments acknowledged. The third asserted the tolerant verdict on
the intolerable case, with a rationale that had been true before fix16 and false ever since. The fourth is
the prod-safety invariant (*never reach for a capture to repair a degraded replay*), and the **unclassified
error is the strongest place to pin it** — it is where a future auto-repair would be most tempted — so it
keeps the harsher input and only its expected exit code moves.

**Cost accepted.** These renames make `git log -S` on the old names miss. The names were wrong; carrying a
wrong name to preserve grep-ability is how a test's claim and its content drift apart in the first place.

## D-M257x-16-3: RF-1 is a port, not a redesign

**Decision.** `migrate-dev.sh`'s atlas loop is made byte-for-byte equivalent in *shape* to
`migrate-demo.sh:150-177` (capture, classify, `mig_fail`, refuse to report OK) rather than being
independently improved.

**Why.** The finding is an asymmetry between twins, and the repair for an asymmetry is symmetry. Any
improvement invented here would have re-opened the same gap in the other direction, and the parity fence —
which is the thing meant to notice — reads string equality between the two files. A port keeps the fence
meaningful; a redesign would have needed the fence rewritten in the same commit that was supposed to be
proving it.

**Deliberate non-change:** `atlas` exits 0 with nothing to do, so no service's happy-path output changes.
That the diff is invisible until something breaks is a feature, and it is what makes this safe to land
without a live migrate run.

## D-M257x-16-4: the non-fatality fence asserts the proposition, not the mechanism

**Decision.** `test_setdress_is_non_fatal` no longer string-matches `if ! env `. It asserts, against the
**extracted set-dress block** rather than the whole file, that the invocation's exit is captured
(`|| sd_rc=$?`) and is never converted to `|| exit` / `|| die`.

**Why block-scoped.** `up-injected.sh` legitimately contains `|| exit` for steps that *are* fatal. A
whole-file `assertNotIn` would be asserting something false about a different part of the script — §8
rule 4, and the mirror image of the trap iter-13's compose fence hit (an assertion correct in principle,
evaluated against the wrong enclosing scope).

**Why this counts as landing RF-4 rather than as collateral.** The test went red *because the fix was
correct*: `if !` collapses every non-zero exit into one boolean, and distinguishing rc=3 is the whole
content of RF-4. A fence written against non-fatality-the-claim would not have moved. Repairing it is part
of the change, not a side discovery, so it does not appear as a side-deliverable in the close.

## D-M257x-16-5: routed, not absorbed — the parity fence's hand-maintained list

**Decision.** The four M215 F8 guard strings are added to
`test_parity_with_migrate_demo_on_the_load_bearing_guards`, and the structural finding is routed forward as
`CHECK-M257x-iter16-parity-fence-hand-maintained` rather than fixed here.

**Why route.** Patching the list closes today's hole. The shape — *a parity fence whose scope is a
hand-maintained enumeration of the system's guards* — is the milestone's own §2 class turned on the fences
themselves, and §8's iter-08 finding already says a fence's scope is the worst possible place for a
hand-maintained list, because **a guard nobody added cannot go RED**. That is a design change with its own
derivation question (what *is* the machine-readable set of "load-bearing guards"?), and opening it here
would have been the iter's third line. Tripwire respected.
