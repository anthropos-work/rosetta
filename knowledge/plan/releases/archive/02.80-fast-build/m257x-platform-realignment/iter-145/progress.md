**Type:** tik — under [`TOK-08`](../decisions.md) (*census the mechanical classes; stop sampling them*).
A test suite is the most mechanical census this milestone has: every instance is enumerated by
construction and each one is decided by running it. This iter censuses the one class `TOK-08` never
reached, because nobody ran it.

# iter-145 — the 21 graded, and 12 of them are OURS

## Phase A — run the four sections nothing in this milestone has ever run

Five sections exist. Every "whole suite" number in 144 iters and 32 harden passes has been
`stack-core` alone. Run at `rosetta-extensions` `aee2bde` (working tree), 2026-08-08:

| section | result | wall |
|---|---|---|
| `demo-stack` | **9 failed** · 1,047 passed · 2 skipped | 3:36 |
| `dev-stack` | 151 passed | 1:39 |
| `stack-injection` | 335 passed | 0:07 |
| `stack-verify` | **12 failed** · 225 passed | 6:02 |
| **total** | **21 failed** · 1,758 passed · 2 skipped | ~11:24 |

**21 — the routed count, reproduced exactly.** So the population is real and stable, and the question
is the one nobody asked: *what are they?*

### The measurement cost two runs before it took, and the fact was already written down

`python3` on this shell is homebrew **3.14 and has no pytest**; the only interpreter on this host that
has it is **`/usr/bin/python3` (3.9.6)**. That is recorded — twice — in `hardening-ledger.md`
(`:2116`, `:2250`) and **nowhere in `platform-alignment.md`**, which is the doc an iter reads. A
measurement precondition that lives only in the harden ledger is a precondition the iter loop
rediscovers. `§9` gains it.

## Phase B — grade every one of the 21

`D-M257x-144-2`, landed one iter ago: **grade a survey arm's findings before treating its count as a
backlog — a routed count is an estimate of WORK, and quoting it makes it a measurement of DEFECTS.**
The 21 has been routed three times (passes 30, 31, 32) carrying one characterisation:

> *"They are provably not ours. `git diff --name-only 6ad8866..HEAD` returns 5 files, all `stack-core`
> … The failures are live-clone / live-container assertions … **Pre-existing, environment-coupled**,
> and out of this pass's iter-diff scope."* — `hardening-ledger.md:2612`

That was derived from a **diff scope**, and a diff scope cannot distinguish a defect from an
environment. Graded individually, the 21 is **three populations, not one**:

| cause | n | how it was decided (not by name) |
|---|---|---|
| **A real defect, and OURS** — the test-side mirror of a platform service this milestone itself deleted | **12** | every one in `stack-verify`; single root cause; repaired below |
| **Whole-file sha baseline drift** against an advanced but **clean** clone | **6** | `git status --short` on `stack-demo/next-web-app` + `studio-desk` is **empty**, so not a dirty clone; and the sibling **anchor** assertions in the same classes are **GREEN** |
| **Host environment** — no live postgres socket for the harness | **3** | `pg_isready … /var/run/postgresql:5432 - no response`; `migrate-demo.sh` aborts before seeding |

**The routed characterisation is falsified for 12 of 21 (57 %).** They are ours, they are one defect,
and they are not environment-coupled at all — six of the twelve are pure table arithmetic that touches
no container and no clone.

### The 12: this milestone's own iter-13, invisible for 132 iters

Platform `2adcf71` (2026-07-31, PR #23) deleted the WunderGraph/Cosmo router. **M257x iter-13**
(`4414527`, 2026-08-01) re-pointed the tooling off it and dropped the `graphql` row from
`stack-verify/lib/services.sh` — correctly, with a comment saying so (`services.sh:53`). It left the
**test side's copy of that row** in place: a `BASES` map at `test_verify.py:788` still carrying
`"graphql": 5050`, plus **six independent count literals** (13 · 13 · 12 · 10 · 13 · 14) restating how
many rows the table has.

Twelve tests went RED that day. They stayed RED for **132 iters**, because no iter close and no harden
pass in this milestone has ever executed the `stack-verify` section.

**iter-13's own commit message names the defect it was fixing:**

> *"six hand-written 5050 sites collapsed into ONE derived `browser_graphql_endpoint` — **six copies of
> a platform fact is the hand-maintained-tuple defect M257x exists to end**"*

It then left a seventh copy in the test file, and the seventh is the one nothing watched. The
milestone's central thesis, reproduced inside the milestone's own tooling, by the iter written to end
it.

**The failure messages are why it stayed invisible even when read.** `12 != 13`, `9 != 10`,
`'all 14 expected container(s) running' not found in '… all 13 …'`, and one bare
`TypeError: 'NoneType' object is not subscriptable`. Not one of them names the row, the file, or the
service. A reader skimming that list sees arithmetic noise; the routed one-liner called it
*"live-clone / live-container assertions"* and six of the twelve touch neither.

## Phase C — repair, and fence the class

`stack-verify/tests/test_verify.py`, +85/−31, **one file**:

1. **ONE literal registry.** `REGISTRY_BASES` promoted to module scope — the `(name, base-port)` map for
   all 12 rows, `graphql` removed. It stays **hand-written on purpose**: it is the anti-vacuity control
   for the offset sweep (`§8` — a port expectation read out of the table under test asserts nothing).
   What it must not be is one of seven copies of *how many rows there are*.
2. **Every count DERIVED from it** — `len(REGISTRY_BASES)`, `len(BACKEND_INFRA_ROWS)`,
   `len(BACKEND_INFRA_ROWS) + 1 + len(INJECTED_DEMO_CONTAINERS)`. Six literals → zero. The two
   `--services` scope strings are now joined from the same tuples instead of being retyped.
3. **The fence** — `test_the_test_side_registry_mirrors_services_sh` asserts
   `set(REGISTRY_BASES) == {r["name"] for r in emit_rows({})}`, and its failure message **names the
   drifted rows in both directions**. This is the assertion that would have caught iter-13 the day it
   landed.
4. Two assertions that were *about* `graphql` re-pointed at a row that still exists rather than
   deleted: the offset check now uses `roadrunner` (10400 → 40400, the high-base row that crosses the
   5-digit boundary — the arithmetic actually worth asserting), and the created-but-never-started
   container check uses `storage`.
5. The `"14 of 16 containers Up"` figure M256 iter-15 measured is **kept attached to its ref** rather
   than restated: the comment records that the arithmetic is now `12 + 3 = 15` and was `13 + 3 = 16`
   *before `2adcf71`*.

**Result: `stack-verify` 12 failed · 225 passed → 0 failed · 238 passed.**

### The anti-vacuity control on the new fence — run, not assumed

`§8`: *write the anti-vacuity control against the guard's SUBJECT.* A copy of the repaired file with
`"graphql": 5050` put back into `REGISTRY_BASES` was run as `zz_drift_control.py` and removed:

```
FAILED …::test_the_test_side_registry_mirrors_services_sh
  AssertionError: Items in the first set but not the second:   ← NAMES the row
FAILED …::test_every_base_shifts_by_exactly_N_times_10000
  AssertionError: 12 != 13 : all 13 services present at demo-1  ← names nothing
```

The fence fires, and the pair is the point: **both assertions detect the drift; only one of them says
what drifted.** The old file had six of the second kind and none of the first.

## Phase D — the scope call, stated as an assumption

The orchestrator's standing instruction is that an iter needing a position states the assumption and
records it. The position, `D-M257x-145-3`:

> **"The suite" for M257x means all five `rosetta-extensions` sections, not `stack-core` alone.**

Not as a preference — as the finding. Widening it by one run surfaced a real defect that the narrow
definition had hidden for 132 iters, and the narrow definition is *why* it hid. The **user has not
ruled**, and this does not pre-empt that; it records what the widening measured so the ruling has
evidence instead of a count.

## Close — 2026-08-08

**Outcome:** the never-run sections were run and the **21 is graded: 12 real defects (OURS) · 6 sha
baseline drift · 3 host environment** — falsifying the *"provably not ours / environment-coupled"*
characterisation for **57 %** of it. All 12 are one cause: **M257x iter-13 deleted the `graphql` row
from `services.sh` and left the test side's copy plus six count literals**, RED for **132 iters**
because nothing ever ran that section. Repaired to **0 failed · 238 passed**, the seven copies
collapsed to one literal registry + derived counts, and a **named-drift fence** added and proven to
fire. `§5` gains **rule 68**; `§9` gains the interpreter precondition.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–145 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n (**the 12 are repaired, not escalated: they are a test-side mirror, not a write path clause 4 covers**) — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-145-1` (a diff SCOPE cannot grade a failure — *"not ours"* derived from
`git diff --name-only` over a 5-iter window mis-classified a defect this milestone introduced 132
iters earlier) · `D-M257x-145-2` (keep the MEMBERSHIP literal, derive the COUNT, fence the two — the
count literal is the copy that rots, the literal set is the anti-vacuity control and must stay) ·
`D-M257x-145-3` (the scope call, stated as an assumption pending the user's ruling).
**Side-deliverables:** none.
**Routes carried forward:**
- **`FIX-M257x-iter145-sha-baseline-drift` (NEW)** — the 6 sha failures. The **anchor** assertions in
  the same classes are GREEN, so by `demopatch-spec.md`'s own model (*the anchor is the contract; the
  whole-file sha is only a baseline*) the patches still land. What is unresolved is whether a test
  asserting the baseline should be RED on every box whose clones have advanced, or should grade the
  baseline separately from the contract. **Do not "fix" it by re-pinning the shas** — that converts a
  freshness signal into a maintenance chore and the next real anchor move would land in the same
  commit unnoticed.
- **`FIX-M257x-iter145-migrate-race-needs-a-host-postgres` (NEW)** — the 3 `test_migrate_race_live`
  failures are a genuine host-environment coupling (`pg_isready … no response` on the host socket).
  They should SKIP loudly with the coverage-hole wording `test_ssr_origin_chain` already uses
  (*"SKIPPED, NOT PASSED"*), not fail — a failure that everyone learns to ignore is worse than a skip
  that says it is a hole.
- **`FIX-M257x-iter145-green-but-stale-graphql-mentions` (NEW)** — `graphql` still appears in ~9
  **passing** `test_verify.py` sites (stub readiness scopes, `:890`/`:897`/`:904`/`:1693`/`:1699`/
  `:1900`/`:1906`/`:1970`). They pass because those tests use their own fixture scope, so the name is
  inert — which is exactly the `§5` iter-23 class (*a named-consumer list survives the merge that
  moved the consumer*). Inert today, a false landmark for the next reader.
- `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` (⚠️ **this iter appends rule 68, so the
  offset grows again — third consecutive iter to widen it and say so**) ·
  `FIX-M257x-iter142-value-change-articles` · `-iter142-path-arm-window` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
- **CLOSED by this iter:** `FIX-M257x-h30-nonstackcore-suite` — run, graded, 12 of 21 repaired, the
  other 9 re-routed as two *named* classes instead of one unexamined count.
**Lessons:**
0. **A suite you never run is not a green suite — it is an unmeasured one.** 132 iters of *"the whole
   suite passes"* were true of 42 % of the suite. The milestone that exists to catch denominators
   quoted one that omitted the section holding its own defect.
1. **"Not ours" is a claim about authorship and a diff scope cannot make it.** `git diff 6ad8866..HEAD`
   asked *"did the last five iters touch this?"* and answered no, correctly. The defect was iter-13's,
   132 iters upstream, and no window that starts after the breaking change can ever see it. **Bisect
   the failure, don't scope the diff.**
2. **The count is the copy that rots; the set is the control.** Six literals restating the size of a
   table are six things to forget. One literal set plus derived counts is one — and a fence between
   the set and the table turns silent arithmetic into a named row.
3. **A failure message that names nothing gets ignored, and being ignored is how it survives.**
   `12 != 13` was visible on any run for four months of iters. It said nothing about `graphql`,
   `services.sh` or the platform, so even a reader who ran it had no reason to look.
