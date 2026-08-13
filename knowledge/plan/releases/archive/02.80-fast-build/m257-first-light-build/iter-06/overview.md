---
iter: 6
milestone: M257
iteration_type: tik
status: in-progress
opened: 2026-08-11
---

# iter-06 — `FIX-M257-load1-units-vm`: clause 1 was comparing two machines

**Type:** tik · **Active strategy:** [`TOK-02`](../decisions.md) — step 2, *fix the instrument's units
before trusting its refusal*.

## Step 0 — re-survey

`TOK-02` was authored one iter ago, so its named target cannot have rotted; re-verified anyway, since the
whole reason iter-04 exists is that a route list went stale during a pause:

- `buildbench.py` clause 1 still reads `profile["cores"]` — unchanged.
- The three shipped profiles still declare `cores` per their `kind`; `macmini.json` still has no host core
  count.
- Live: host `os.cpu_count()` = **12**, `sysctl hw.logicalcpu` = **12**, `docker info` NCPU = **8**.

Target confirmed, untouched, and still the thing standing between this milestone and a trustworthy
measurement.

## Cluster / target identified

The gate's HEADROOM clause 1 is *"peak load1 ≤ cores − 2"*. Both live paths read `os.getloadavg()`, which
on this host is the **macOS host's** load across its **12** logical cores. Clause 1 grades that against
`profile["cores"]`, which `macmini.json`'s own `budget_source` declares to be *"the Docker Desktop VM
allocation, NEVER host totals"* — **8**. So the limit computes to **6** where the correct limit is **10**.

Two properties make this worth fixing before, not after, the baseline campaign:

1. **It fails CLOSED.** It produces refusals, not false greens — and a spurious refusal is
   indistinguishable from a real one at the point of use. `TOK-02` step 3 says a clause-1 refusal is a
   *result*; that is only true if the clause is grading the right machine.
2. **The one real refusal on record is from a host where the two counts COINCIDE.** `laptop.json` records
   FAIL clause 1 at load1 10.69 vs `cores-2 = 8` — correct, but only because that machine's VM allocation
   and host core count are both 10. A check that is right by coincidence looks exactly like a check that is
   right, which is how this survived a live refusal.

**The instrument already knows the distinction twice**, which makes this a slip rather than a
misunderstanding: `engine_facts()`'s docstring spells out *"`os.cpu_count()` is 12 and `docker info` NCPU
is 8"*, and `profile_describes_host()` grades cores *"against the quantity the profile's own `kind`
declares it to be."* Clause 1 is the third site and it forgot.

## Hypothesis

Grading `load1` against the logical-core count **of the machine the sample came from** makes clause 1
report the host's real saturation state. Where the basis cannot be established, the clause must **fail**
rather than fall back to the VM number — falling back is the defect itself.

## Expected lift

**Zero on the primary metric, by design** — no lever is touched and no cycle is run. The deliverable is an
instrument whose refusals can be believed, which is `TOK-02`'s stated precondition for the baseline. Graded
on planned deliverables per Phase 4 Step 0.

## Phase plan

1. `load1_core_basis(profile, observed_host_cores=None)` — precedence: observed (live, same-process) →
   declared `host_logical_cores` → `profile["cores"]` **only** for `native-linux` → else `None`.
2. Clause 1 consumes it; a `None` basis with a measured `load1` is a new **fail-closed** clause.
3. Both live call sites pass `os.cpu_count()` from the process that read the loadavg.
4. `kind` becomes a loader-required key (without it the basis is unknowable).
5. Profiles declare `host_logical_cores`, each with its own provenance.
6. Tests: negative control (basis absent ⇒ RED), the regression on the real profile at the real observed
   peak, a still-refuses control, the laptop-refusal-survives control, a wiring test that the kwarg is
   actually **passed** and not merely defined, and a discovered-not-listed sweep over shipped profiles.
7. **Mutation control**: restore the old basis in memory and confirm the new tests go RED.

## Escalation conditions

- If correcting the units turned out to *weaken* the clause in substance rather than re-point it, that is a
  gate change and belongs to the user, not to a tik.
- If the fix could not be proven RED with its precondition absent, it would be a probe that fails open —
  do not ship it.

## Acceptable close-no-lift outcomes

If the units had turned out to be correct as written (e.g. the sampler already ran inside the VM), the
falsification itself would be the deliverable and the iter would close no-lift.

## Deliberately NOT in this iter

`BASELINE-M257-macmini-n3` → **iter-07**. It is a long-running measurement, it is the *consumer* of this
fix, and starting it before the instrument is trustworthy is the ordering `TOK-01` got right and this
milestone has already paid to relearn once.
