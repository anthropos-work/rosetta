# iter-170 — decisions

## `D-M257x-170-1` — the census reports a RUNNER, because the population needs two of them

**Decision:** `suite_census` measures per-module per-runner, defaults to `--runner both`, and prints the
runner-disagreement set explicitly. It **refuses** to run the pytest half when `/usr/bin/python3` is absent
rather than reporting half a population as a whole one.

**Measured.** `/usr/bin/python3` == **3.9.6**, the only interpreter on this box carrying **pytest** — the
fleet runner. The working interpreter is **3.14.6** with **no pytest**. Over 110 modules the two disagree
about **four**, and the disagreements are not noise: two modules `import pytest`, one relied on pytest
putting a test's own directory on `sys.path`, and one binds a server. Booked as `§5` **rule 75**: a green
from one runner is not a green, and *state the runner with every suite number* is the pass/fail twin of
`§8`'s *state the environment with every timing number*.

**Consequence for this milestone's record:** iter-167's family reading (*17 GREEN · 0 RED · 7 not-run*) and
iter-169's rotted assertion both live in this gap. `test_claim_census_guard` — one of the four
waiver-carrying guards iter-166 shipped — **has never executed on the modern interpreter at all.**

## `D-M257x-170-2` — the environment bucket is DECLARED, not sniffed

**Decision:** replace substring-sniffing of error text with an explicit `ENV_GATED` map keyed
`module::test`, and derive only its **completeness** — a declaration whose test no longer exists is
**STALE**, an undeclared failure is **ACTIONABLE**.

**Why the change was forced.** The first draft's sniffer returned **ZERO** against **nine** genuinely
environment-gated failures. The signal was never in the output; it was in what the tests *are*. A third
bucket that never fires is a two-bucket partition wearing three labels — the precise defect `§5` rule 73
names, committed by the instrument written to honour it.

**The repair is rule 73's own prescription** (*keep the partition DECLARED, derive its COMPLETENESS*), and
it ships with controls that fire: each of the four buckets is shown to fire on a synthetic module; the ENV
bucket is shown **not** to swallow an undeclared failure; the staleness check is shown non-vacuous.

**Name-matching was considered and rejected.** The nine all sit in classes like `TestSSRManifestAgainstLiveClone`,
so a `*Live*` pattern would have worked today — and `§5` rules 70/71 are explicit that a fence pinned to a
**spelling** is not pinned to a **property**. A declaration is the honest form.

## `D-M257x-170-3` — the six sha-drift failures are QUANTIFIED and deliberately NOT repaired

**Decision:** record exactly which six tests carry `FIX-M257x-iter145-sha-baseline-drift` and leave them
RED.

**Rationale.** `stack-demo/` clones are **present** on this box, so these are real drift against real files —
not the absence of a clone. That is exactly what the pins are for. Re-pinning them would convert a
**freshness signal into a chore**, which the standing route forbids in as many words. Declaring them keeps
the census honest (`0 ACTIONABLE`) without silencing them: they still print, in their own bucket, on every
run.

**What changed is that "some sha tests are drifting" is now a number.** Six tests, three modules
(`test_demopatch` ×2, `test_ssr_origin_chain` ×3, `test_ant_academy` ×1). iter-168's rule applies —
*measure the hazard, or "the same problem exists elsewhere" is only a mood.*
