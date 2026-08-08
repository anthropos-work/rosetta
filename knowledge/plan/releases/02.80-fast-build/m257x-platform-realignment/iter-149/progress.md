**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

# iter-149 — the emitter census, widened from one retired service to all twelve

## What was measured

iter-146 censused **one** retired platform fact (the GraphQL router platform `2adcf71` deleted): 84
references, 82 correct, 2 emitter defects. It routed the same question at the other retired names and
warned against widening its fence un-audited. This iter ran the full grid.

**Population.** The retired set is **not** a thirteenth hand list — it is
`claim_census_guard.ARCHIVED_SERVICE_NAMES` (12 names), already the row set of
`platform-migration-status.md`, itself fenced against the platform's `repos.yml` in both directions by
`platform_alignment_guard`. The **arms** are the four ways a dead service gets named in content a machine
or an operator consumes: base **port**, **container** name, `host:port` **address**, `*_RPC_ADDR`
variable. Base ports were **derived**, not recalled — from the tooling's own probe registry at the ref
before the merges stripped its rows (`stack-verify/lib/services.sh` @ `c95bce4`): skiller 8085 ·
skillpath 8100 · cms 8090 · jobsimulation 8400 · storage 8300 · roadrunner 10400 · graphql 5050.

**The reading, over the whole monorepo:**

| | |
|---|---|
| references in the 12 × 4 grid | **354** |
| in executable-position code | **17** |
| …Python **docstring** prose inside four guards (`platform_predicate_guard`, `anchor_construct_guard`, `retracted_pin_guard`, `corpus_citation_guard`) | 9 |
| …the four retained `stack-verify/lib/services.sh` probe-registry rows (4 rows × {port, container}) | 8 |
| **emitters** | **0** |
| retired names with **zero** references in every arm | chronos · intelligence · messenger · customerio-sync · db-backup |
| retired names appearing **only** as an `*_RPC_ADDR` mention, all of it doc/test/comment | skiller (20) · skillpath (4) |

**The census's own first reading was wrong, and that is the finding.** The raw run returned **50**
executable-content hits, not 17. **33 of them — 66 % of the signal — were `.m220-mutant-*` files**: staged
copies of `dev-stack` that the M220 mutation battery writes *beside* the real subject and deletes in
`tearDown`, left behind by interrupted runs and **accumulating since 2026-08-04**. Every one of them
carried, verbatim, the exact defect line iter-146 repaired. They are untracked, correctly gitignored, and
nothing invokes them — but a stale mutant is a **perfect forgery of the bug its own battery exists to
catch**, and it is the worst thing to leave lying in a tree that gets censused. `.gitignore:32` already
described the leak (*"this is the belt for an interrupted run"*) — iter-148's lesson 0 again: **a file
that discloses a hazard in a comment is a file nobody fixed.**

## What landed

1. **The leak is self-healing.** `ShellMutationHarness.setUpClass` now sweeps abandoned
   `.m220-mutant-*` siblings before staging. **Age-gated at 1 h**, not unconditional: a mutant belongs to
   a live run for as long as that run holds it, and this battery must not delete a concurrent sibling's
   staging mid-measurement — a failure that would read as a mutation result rather than as a missing file.
   Proven **on the real leftovers before the fixture was spent** (`§5` rule 21): 33 → 0, repo-wide.
   Paired control (`StaleMutantsAreSweptAndLiveOnesAreNot`): an abandoned mutant is reclaimed, a young one
   and the real subject beside it are not.
2. **The fence generalised** — `test_deleted_router_endpoints.py` → `test_retired_service_endpoints.py`,
   from one hard-coded port to the imported retired set × 3 arms, over the same emitter allowlist.
   iter-146's comment carve-out and its both-directions pair control are kept intact.
   - **RED-proofed on a real answer key:** run against the pre-iter-146 `dev-stack` (`1a44b97^`) it
     returns **1 hit — the original defect line**; against the current file, **0**.
   - **Every arm proven to fire.** The census returned zero, so the tree cannot demonstrate the container
     and address arms work; each is shown tripping on synthetic content written for it. An arm proven only
     by a green tree is an arm that may be matching nothing.
   - **The `services.sh` carve-out is declared, not silent.** Those four rows DECLARE a probe target that
     a scoped run filters out; they emit to nobody. iter-148 measured the cost of reading that table
     unscoped (6 of 20 probes failed vs 1 of 14 scoped) and fixed the caller; the table is owned by
     `SURVEY-M257x-iter148-registry-is-hand-maintained`. Matching it here would be a standing RED with a
     known owner, which is how a fence gets switched off. An anti-vacuity control asserts the carve-out
     names a file the fence actually scans.

## Side discovery (separate commit, does not change this iter's status)

`claim_census_guard.REXT_SECTION_NAMES` described itself as *"derived from the monorepo's own layout"*.
It was **declared**, and it had **drifted**: `stack-secrets` was missing — 10 names for 11 sections on
disk. So every claim naming the section behind the `/stack-secrets` skill resolved to no known artifact
and left the census silently. **This is exactly the enumeration defect M257x iter-129 repaired in
`CLAUDE.md`, one layer down, in a guard whose job is to notice omissions.** Repaired, the comment made
honest about being declared and why (this module is pure data with no repo-root notion and is imported
from copies whose layout differs), and the property the old comment merely asserted is now fenced by a
test that compares the tuple against the on-disk layout in both directions. RED-proofed: the pre-fix
tuple fails it, the post-fix tuple passes.

## Tests

- `test_retired_service_endpoints.py` — **5/5 green** (incl. the arm RED-proof and the carve-out pair).
- `test_m220_mutation_battery.py` — **12/12 green**, 105 s, with the sweep in `setUpClass`.
- `test_claim_census_guard.py` — **25/25 green** under `/usr/bin/python3 -m pytest`, including the new
  control 34. **This nearly became a false finding, and the near-miss is worth recording:** `python3` on
  this shell is homebrew 3.14 with no pytest, the run failed `No module named pytest`, and the first draft
  of this close booked a new route for it. `§9`'s *Measurement preconditions* already states the fact in
  its first bullet — *the suite's interpreter is `/usr/bin/python3`, and it is the only one on this host
  with pytest; iter-145 lost two full runs to it.* Reading the protocol section cost one lookup and
  deleted a fabricated route. **The section works; the failure mode is not reading it.**
- **Not run:** the rest of `stack-core` (~20–35 min) and the other four sections. `§5` rule 60's scoped
  default. Only two files in the monorepo consume the data this iter changed, and both are named above.

## Close — 2026-08-08

**Outcome:** the emitter class censused across **all twelve** retired services × 4 arms — **354
references, 17 in executable position, 0 emitters** — closing iter-146's route with a falsification
rather than a repair. The defect the census found was in the **census's own denominator**: 33 abandoned
mutation-battery stagings, 66 % of the raw signal, each a verbatim copy of the bug iter-146 fixed. Swept,
self-healed, fenced; the emitter fence generalised and RED-proofed on real pre-fix content.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–149 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-149-1` (a stale mutation staging is a perfect forgery of the bug its battery
catches — sweep it, age-gated, never unconditionally) · `D-M257x-149-2` (bind a fence's subject set to
an existing fenced one; a thirteenth hand list is the defect) · `D-M257x-149-3` (when a census returns
zero, the arms must be RED-proofed on synthetic content — a green tree cannot tell a working arm from a
decorative one) · `D-M257x-149-4` (a carve-out with a named owning route is a waiver; one without is a
hole).
**Side-deliverables:** `REXT_SECTION_NAMES` was missing `stack-secrets` (10 declared / 11 on disk) — the
iter-129 class inside a guard. Fixed + fenced against the layout, separate commit.
**Routes carried forward:**
- **`SURVEY-M257x-iter149-declared-lists-unfenced-against-layout` (NEW)** — `REXT_SECTION_NAMES` claimed
  a derivation it did not perform. Nothing has asked how many other declared tuples in the fence family
  make that claim. The check is one comparison per list.
- `SURVEY-M257x-iter148-registry-is-hand-maintained` (now also this fence's declared carve-out) ·
  `SURVEY-M257x-iter147-absent-value-class` (`STACK_PROJECT` / `STACK_OFFSET` arm still open) ·
  `SURVEY-M257x-iter146-other-retired-services-unaudited` — **CLOSED by this iter** ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter144-correction-vs-retraction-unfenced` · `FIX-M257x-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `FIX-M257x-iter142-value-change-articles` · `-iter142-path-arm-window` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` (**narrowed** — this fence imports its subject set
  instead of re-declaring it, the first seam in the family to do so) · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **Sweep the substrate before reading the census, or the substrate IS the census.** Two thirds of this
   iter's raw signal was test staging that no longer belonged to any run, and every one of those hits was
   a byte-identical copy of a defect already repaired. iter-148 said *state the substrate before booking a
   failure*; this iter says the stronger form — **a measurement over a working tree must first establish
   what in that tree is a measurement subject and what is exhaust.**
1. **A census that returns zero must prove its instrument.** Every arm that finds nothing is
   indistinguishable, from a green tree alone, from an arm that matches nothing. The RED-proof on real
   pre-fix content plus the per-arm synthetic trip is what makes a zero worth publishing.
2. **A list that calls itself derived and is not, drifts — and it drifts inside the guards too.**
   `REXT_SECTION_NAMES` had the word "derived" in its own comment. The repair is not to derive it (this
   module cannot), it is to **fence the property the comment asserted**.
