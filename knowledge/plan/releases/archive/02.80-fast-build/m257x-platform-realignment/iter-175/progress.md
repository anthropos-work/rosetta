# iter-175 — progress

**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase 1 — the census the route asked for

iter-174 routed a sentence that is a class, not an item: *"five registries are now known; nothing
enumerates them."* The "five" was a **remembered list** — four at iter-173, five at iter-174, each reached
by grepping the tooling for a sibling's name. §2's hand-maintained tuple, one level up.

Two instruments, both recorded (`D-M257x-175-2`):

| instrument | predicate | population |
|---|---|---|
| A (rejected) | a file naming ≥2 fence modules as quoted literals, anywhere | **39** |
| B (used) | a **collection literal** — py `List`/`Tuple`/`Set`/`Dict`-keys/call-args, or a JSON array/object — holding ≥2 fence-module names | **5** |

A measures *mentions*. The claim is about **a set that must track the tree**, so A is graded at the wrong
grain (§9, iter-159) and its 39 rows would have been ~34 declines of a test that happens to name two
guards — volume that reads like rigour and answers a different question.

**Instrument B's five**, fence population derived from `FENCE_KIND` declarations (n = **26**):

| site | names | membership derived by | reported by |
|---|---|---|---|
| `guard_family.py:78` `INVOCATIONS` | 24 | reconciled against `census()`, both ways | **nothing — see below** |
| `repair_postcondition_baseline.json` | 6 | written by `--accept` | ratchet |
| `test_m257x_mechanical_fences_mutation_battery.py:70` `_COPY_FILES` | 6 | hand-maintained | a 14-min battery (`FIX-M257x-iter174-…`) |
| `test_iter45_mechanical_fences.py:384` | 3 | test subject, not a registry | n/a |
| `test_fence_registry_completeness_m257x.py:83` | 2 | the iter-157 regression pin | n/a |

The route's own item is **third by size** and *is* already reported, late but correctly. The top row is
reported by nothing — so the target was substituted under the same strategy (`D-M257x-175-1`), and the
seed-list route stays open.

## Phase 2 — the finding

> **`guard_family.census()` derived the fence family from a FILENAME SPELLING.
> `repair_postcondition.discover_fences()` derives the SAME population from the DECLARATION.
> They disagreed by three members, in both directions, and no fence compared them.**

Reproduced from the **committed** code (`git show HEAD:stack-core/guard_family.py`, run against the live
tree — not from a reconstruction):

```
PRE-FIX census (HEAD code, live tree) n = 25
FENCE_KIND registry            n = 26
registry ∖ census : ['guard_family', 'predicate_enumerator']
census ∖ registry : ['repair_postcondition']
```

**This is iter-157's defect, in the sibling module, still live.** iter-157 measured *"25 modules declared a
`FENCE_KIND`; 23 were enumerated"* in `repair_postcondition`, repaired it, and shipped
`test_fence_registry_completeness_m257x.py` — **which fences that one module's registry.** iter-169's rule
one turn on: *closing a class means fencing its POPULATION, not its last member* — and **the population
here is the set of DERIVATIONS.**

**The consequence is the one `guard_family` exists to prevent.** `predicate_enumerator.py:142` declares
`FENCE_KIND = "standalone"`; the runner whose docstring promises to *"run the WHOLE guard family, and name
every member"* had **never run it and never named it.** Not NOT-RUN in the verdict — **absent from it**,
while the summary line printed an unqualified count. Its own §5 rule 8 is the indictment: *a guard that was
not run reads exactly like a guard that passed.*

## Phase 2 — the repair

**`census := spelled ∪ declared ∪ extra`, minus a declared exclusion table** (`D-M257x-175-3`). The
apparent repair — swap the glob for the declaration, the symmetry with iter-157 — is a **weakening dressed
as a tidy-up**: a `*_guard.py` that declares nothing would then leave the family in silence, and that file
is precisely the one worth catching. A member needs only ONE property, so both gaps close and neither can
lose a member (iter-158). `EXTRA_CENSUS_MEMBERS` survives, demoted to **additive only**.

`declaring_modules()` calls `repair_postcondition.declared_kind` rather than re-implementing the AST read
— **one reader, two consumers**. A private copy would have been a *third* derivation of the same
population, committed while repairing exactly that.

- **`predicate_enumerator` is INVOKED, not excluded** (`D-M257x-175-4`) — `cls: "input"`, ledger-scoped
  like `repair_reach_guard`. It takes no `--range`, and `run_one` derived that requirement from the class,
  so it would have reported **"needs --range, not supplied"** — a NOT-RUN reason naming a flag the guard
  does not accept. `needs_range` is now declared per member, **defaulting to the class rule**, so every
  existing member's behaviour is byte-identical and the one that differs says so.
- **Exclusions are PRINTED, and held to two directions** (`D-M257x-175-5`) — stale (subtracts nothing) and
  ambiguous (also invoked) are both complaints, so the table cannot become a place to park a member.

**Proven live** (`--allow-not-run`, no platform/range/ledger):

```
fence-tree: …/rosetta-extensions @ 17d33e913 · DIRTY …
guard-family: corpus /Users/marco/workspace/anthropos/rosetta @ 5d0fb6619
guard-family: EXCLUDED  guard_family — this runner itself. …
guard-family: 26 member(s) on disk, all placed.            ← was 25
  NOT-RUN   rc=-  predicate_enumerator  [input]  needs --ledger, not supplied
guard-family: 18 GREEN · 0 RED · 0 could-not-check · 8 not-run
```

## Phase 2 — the fence, and its two controls

`+5` tests on `test_guard_family.py` (`128` inserted lines), all pinned to the **property**, never a
spelling (r70/71):

1. **THE arm** — `test_the_two_derivations_of_the_fence_population_AGREE`: `discover_fences()` ∖
   (`census()` ∪ exclusions) must be empty. RED-proofed against the real pre-fix code.
2. **Mutation control** — re-derives the census by the exact pre-iter-175 rule and asserts the arm **would
   have failed**, requiring every missed member to be a declarer not spelled `*_guard`. Without it, a
   future narrowing back toward the spelling makes the arm compare a set with a superset of itself.
3. **Anti-vacuity** (§9) — both derivations must return ≥20 and the exclusion table must be non-empty, or
   the arm is comparing two empty sets.
4. The union's protected direction — a `*_guard.py` declaring **nothing** is still counted.
5. `test_census_finds_every_guard_on_disk`'s fixture already contained `helpers.py` **declaring a
   `FENCE_KIND`** and asserted it was dropped. The expectation was the bug; it now asserts the union.

## Phase 3 — the whole-population run, and what it caught

`FIX-M257x-iter142-whole-suite-owed`, **fourth consecutive iter in which it paid for itself.** Five
sections, pytest/`/usr/bin/python3` 3.9.6 (iter-170's runner):

> **`12 failed · 3355 passed · 3 skipped in 2071.14 s (34:31)`** — **executed = 3370** (iter-172's unit
> rule: `passed` is not `executed`; the three numbers are summed, `deselected` excluded).

**3 of the 12 were introduced by this iter, and every scoped run was green.** Both causes are this iter's
own subject matter, which is the part worth keeping:

| RED | cause | rule it instantiates |
|---|---|---|
| `test_fence_provenance` ×2 | a **sixth** synthetic guard-dir fixture, in a second file, staging a family with no runner in it → `CENSUS_EXCLUSIONS` stale → `exit 2` before the provenance line | **the fixture repair was a grep of ONE file.** §5 rule 73 (*a glob is not a derivation*) extends to greps of a single file |
| `test_frozen_expectation_census_m257x` ×1 | `declaring_modules` + `union` are two new public derivations; `derivation_registry.DECISIONS` classified neither | iter-162's registry-completeness fence **working exactly as designed** — the fifth registry, catching the iter that was written about registries |

The fixture population was then re-derived over the whole tree (every module patching `guard_family.HERE`,
every `main()` caller) rather than grepped again — **exactly 2 files, 6 fixtures**, and the second one is
now repaired from the same single `stage_runner` definition.

**A 4th RED surfaced during the repair's own scoped re-verify** and is a real property, not a fixture
artifact: `test_fence_provenance` asserts the `fence-tree:` line comes **FIRST**, because iter-105 made it
first — it is the input that decides every verdict below it. The `EXCLUDED` lines had been printed beside
the reconcile that computes them, ahead of it. Moved below the stamp.

**9 of the 12 are pre-existing and already routed**, verified by reading their failures rather than
assumed: 6 live-clone sha-pinned demopatch tests (`FIX-M257x-iter145-sha-baseline-drift` — **do NOT
re-pin**) and 3 that need a host Postgres (`FIX-M257x-iter145-migrate-race-needs-a-host-postgres`).

**Post-repair scoped re-verify: `157 passed in 218.71 s`** over the 5 affected modules.

## Close — 2026-08-09

**Outcome:** Two derivations of one population — the fence family — had disagreed by **3 members in both
directions** with nothing comparing them, so the runner that promises to *"run the WHOLE guard family and
name every member"* had never run `predicate_enumerator` and never named it. Repaired by **UNION** (never
by substituting the declaration for the spelling, which would have been a weakening), the missing member
invoked with a NOT-RUN reason that is true, the one deliberate exclusion **printed with its reason** and
held to two directions. Fenced with THE arm + a mutation control + an anti-vacuity control. The
whole-population run then caught 3 self-inflicted REDs — one of them this iter's own fixture repair having
been a **grep of one file** — and a 4th surfaced on re-verify.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (seventh consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean)
**Decisions:** `D-M257x-175-1` … `D-M257x-175-6` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none. The `derivation_registry` classification and the provenance-ordering move are
not side fixes — both are this iter's own change being graded by the fences it had to join.

**Routes carried forward:**
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — **NEW, measured not guessed** (§8 iter-168).
  `stack-core/README.md`'s *"The corpus guards"* table names **16 of the 27** members of the census —
  missing `anchor_offset_guard`, `blocking_state_guard`, `claim_census_guard`, `clone_drift_guard`,
  `corpus_citation_guard`, `guard_family`, `platform_predicate_guard`, `retracted_pin_guard`,
  `service_registry_guard`, `unreadable_repo_claim_guard`, `value_change_guard`. A **sixth** registry: a
  human-facing index of the family, 41 % incomplete, checked by nothing. Whether the table *claims*
  completeness is the first thing to settle — a selective narrative is not a registry.
- `SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` — **NEW.** The two derivations
  of this one population are classified **differently** in `derivation_registry.DECISIONS`:
  `repair_postcondition::discover_fences` is `REGISTERED`, `guard_family::census` is `DECLINE:verdict`.
  This iter's two new halves were classified **with `census`** (the consistent, non-weakening choice) and
  the inconsistency routed rather than resolved, because deciding it changes what the frozen-expectation
  census treats as a candidate — a third line of investigation this iter declined to open.
- `FIX-M257x-iter174-accept-registers-one-registry-of-two` — **unchanged, open.** Deliberately not the
  target; see `D-M257x-175-1`.
- `FIX-M257x-iter173-ledger-denominator` — unchanged; open (owned by the next harden pass).
- The observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — unchanged; open.
- `FIX-M257x-iter145-sha-baseline-drift` / `-migrate-race-needs-a-host-postgres` — re-confirmed live at
  9 REDs this run; unchanged, and **the shas must not be re-pinned**.
- The standing queue, unchanged.

**Lessons:** **two derivations of one population must be COMPARED, or the weaker one is a silent census.**
A completeness fence proves one registry against the tree; it says nothing about a sibling registry
deriving the same set by another rule — and the milestone had shipped exactly that for 18 iters. The
repair direction is the second half: **UNION, never substitution.** Swapping the glob for the declaration
looks like symmetry with the iter that repaired the sibling, and it silently drops the case the glob was
right about. Both are now `§8`. The corollary this iter earned on itself: **a census of one file is not a
census** — the fixture repair looked complete because the file it was written in was the file it was
derived from.
