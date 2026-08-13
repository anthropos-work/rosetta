**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

# iter-150 — the constants that claim a derivation nobody performs

## The census

Parse-level, over the whole tooling monorepo. Module-level `UPPER = <literal>` assignments, with the
contiguous comment block immediately above each read for a derivation claim.

| | |
|---|---|
| Python files parsed | **147** |
| module-level `UPPER` constants | **960** |
| …that are LITERALS whose comment block mentions a derivation | **30** |
| …narrowed to a **self-directed** claim (the comment says *this* value is derived) | **9** |
| already fenced by a named control | 3 — `ARCHIVED_SERVICE_NAMES` (via `platform_alignment_guard`) · `REGISTRY_BASES` (`test_the_test_side_registry_mirrors_services_sh`) · `SECTION_COVERAGE` (`test_every_go_section_declares_its_coverage`) |
| already repaired hours earlier | 1 — `REXT_SECTION_NAMES` (iter-149) |
| **comment is about a neighbouring value, not this one** | 4 — `FRONTENDS` · `_FINDING_PREFIXES` · `RETIRED_PORTS` · `_TOKEN_KEYS` |
| **unfenced claim — the defect** | **1** — `blocking_state_guard.BLOCKING_FIELDS` |

**The 30 → 9 narrowing is the point, and it is why this was hand-graded.** The word *derive* appears in a
constant's comment for at least four unrelated reasons: something else derives FROM the constant (every
`FENCE_KIND` — 9 of the 30), the constant is a fixture for a derivation under test, the comment discusses
derivation as the design choice being rejected, or the sentence is about a value defined nearby. A
token-level count would have reported **30 defects** where there is **1**. `§5` rule 67's shape again: the
same token carries opposite obligations depending on which way the sentence points.

## The defect

`stack-core/blocking_state_guard.py` — the guard that reads every iter's `**Phase 5 grading:**` line and
checks that each blocking `y` is represented in the milestone's deferral audit. It partitions the exit
enum into two hand-typed tuples:

```
BLOCKING_FIELDS     = ("re-scope", "user-blocker", "protocol-stop")
NON_BLOCKING_FIELDS = ("gate-met", "triggered-tok", "cap-reached", "budget-exhausted")
```

and the comment over the first said *"Derived from the iteration protocol's own Phase-5 grading, not from
what any one audit happened to read."* **Nothing derived anything.** The guard already checked one
direction — `phantom`, a blocking name that no grading uses, which would make it green for free — and
never the other: **a field the protocol grades that neither tuple classifies is treated as non-blocking
by omission, which is indistinguishable here from a decision that it is safe.**

**This is not hypothetical, and the evidence is inside the protocol's own text.** `budget-exhausted` was
**added to the exit enum on 2026-08-06** — the skill records why in as many words: three separate sessions
reported a clean budget stop, the enum had no value for it, each was instructed to emit `user-blocker`,
and each was correctly flagged as a mis-grade. Someone hand-added the new name to `NON_BLOCKING_FIELDS`.
Had they not, an entire exit class would have passed through this guard unclassified and unremarked.

## What landed

**The partition's COMPLETENESS is now derived, even though the partition itself cannot be.** Which side of
the line an exit condition falls on is a judgement about what it *means* — no parse can make it, and
pretending otherwise is how the false claim got written. What *is* derivable is the field universe: the
gradings are the only place in this repository where the protocol's enum is written out field by field.
`run()` now takes the set it already computes (`seen`), subtracts both tuples, and reports the remainder.

- A **finding**, not a `could-not-run`. An unclassified field is a real signal that the protocol moved;
  refusing to run would suppress the blocking gradings this guard exists to surface.
- The finding **names where the field first fired** and how many iters grade it, so the reader can act
  without re-deriving.
- The comment is now honest: **declared partition, derived completeness**, with the reason.

**Controls, as a pair** (`ThePartitionIsCheckedAgainstWhatIsActuallyGraded`):
- a fixture grading a net-new `(8) host-quarantine: y` produces exactly one finding, naming `iter-07`;
- the real seven-field grading produces **none** — because a check that flags every field is noise and
  would fire on the live corpus immediately.

**One existing mutation control needed narrowing, and it is disclosed rather than waived.**
`test_MUT_shrinking_BLOCKING_FIELDS_to_user_blocker_LOSES_the_finding` asserted `findings == []` under a
deliberately-broken partition. With a second mechanism writing to the same list, that mutation now *also*
— correctly — reports `re-scope` and `protocol-stop` as classified by neither tuple. The control is scoped
to the representation mechanism it isolates, with the reason in-line; the `(partition)` findings are
asserted by their own pair above, not ignored.

## Tests

- `test_blocking_state_guard.py` — **19/19 green** (17 before, +2).
- **Live run against this milestone**: `OK — every blocking grading is represented in the deferral audit`,
  and **zero partition findings** across 149 graded iters. So the seven classified fields are exactly the
  seven M257x has ever graded — derived, not assumed.
- **Not run:** the rest of `stack-core` (~20–35 min) and the other four sections — `§5` rule 60's scoped
  default. The change is confined to one guard and its own control file.

## Close — 2026-08-08

**Outcome:** 960 module-level constants censused by parse, 30 whose comment mentions a derivation, **9
self-claiming**, and **1 unfenced** — `blocking_state_guard.BLOCKING_FIELDS`, whose comment claimed a
derivation from the iteration protocol that nothing performed, in a guard whose whole subject is the
protocol's exit enum. The partition's **completeness** is now derived from the gradings themselves and
reported as a finding when the enum grows; the partition itself stays declared, because which side a
condition falls on is a judgement and saying otherwise is what produced the false claim.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–150 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**2 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-150-1` (a declared partition can still have a DERIVED completeness — fence the
half that is decidable) · `D-M257x-150-2` (grade a keyword census by hand before publishing it: 30 raw,
1 real) · `D-M257x-150-3` (narrowing an existing control because a new mechanism shares its output list
is disclosed in-line, never silent).
**Side-deliverables:** none.
**Routes carried forward:**
- **`SURVEY-M257x-iter150-partition-completeness-elsewhere` (NEW)** — `blocking_state_guard` was the one
  guard partitioning an enum it does not own. Two others read protocol-shaped grammar
  (`markdown_structure_guard`, `evidence_visibility_guard`); nothing has asked whether either holds a
  closed-world assumption about a vocabulary that can grow.
- `SURVEY-M257x-iter149-declared-lists-unfenced-against-layout` — **CLOSED by this iter.**
- `SURVEY-M257x-iter148-registry-is-hand-maintained` · `SURVEY-M257x-iter147-absent-value-class`
  (`STACK_PROJECT` / `STACK_OFFSET` arm) · `FIX-M257x-iter145-sha-baseline-drift` ·
  `-iter145-migrate-race-needs-a-host-postgres` · `-iter145-green-but-stale-graphql-mentions` ·
  `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **"Derived" in a comment is a claim, and claims in comments are the ones nothing checks.** Two
   consecutive iters found the same shape in two different guards. The tell is syntactic — a literal on
   the right-hand side and the word on the left — which is why the census is a parse and not a read.
1. **When a value genuinely cannot be derived, its COMPLETENESS often still can.** The judgement half of
   `BLOCKING_FIELDS` is irreducible; the coverage half was sitting in the gradings the guard already
   parses. Splitting a claim into its decidable and undecidable halves is usually cheaper than either
   deriving everything or fencing nothing.
2. **A keyword census must be hand-graded before it is published.** 30 → 9 → 1. iter-138's withdrawn
   *"127 rotted pins"* is the standing warning; this iter is the same arithmetic caught before publication
   rather than after.
