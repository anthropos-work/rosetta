**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-188 — the citation resolver's walk prunes 70 % of its candidate files, and said so nowhere

## Phase A — running iter-187's selector rather than re-surveying

iter-187 routed its residual **with a mechanical selector**: *a registry keyed by a CONTAINER whose
justifying reason is a property of the container's CONTENTS.* Run over `stack-core`'s 30 module-level
registries, the strongest hit is `claim_census_guard._SKIP_DIRS` — twelve directory names pruned from
the `os.walk` that builds the **basename index citations are resolved against**. A directory the walk
never enters is a directory no citation can resolve into: the list decides **reach**, not speed.

**The first reading was wrong about the subject and is corrected here, not quietly** (`D-M257x-188-2`).
`clones` falls back to `root / "stack-demo"` (`claim_census_guard.py:768`) when `--clones-root` is
unset, and that tree exists on this box, so the **default** run is the large one:

| roots walked | files indexed | files pruned | share |
|---|---:|---:|---:|
| **`corpus/` + `stack-demo/` — the DEFAULT** | 21,610 | **50,357** | **70.0 %** |
| `corpus/` alone (`stack-demo` absent) | 92 | 0 | 0.0 % |

**Seventy per cent of the resolver's candidate files never enter its index, and no output said so.** Of
the twelve names, **2** did all the pruning (`node_modules` 3 directories, `test-results` 2), **5** were
inert here (`vendor`, `dist`, `build`, `coverage`, `__pycache__`) — and **5 can never fire at all.**

**A registry silently subsumed by an adjacent unnamed rule.** The walk reads
`d not in _SKIP_DIRS and not d.startswith(".")`, so `.git`, `.next`, `.venv`, `.turbo` and `.pnpm-store`
cannot ever be the reason a directory was pruned. A reader auditing *"what do we exclude?"* would have
counted twelve; seven is the answer, and two is what fired.

**Hazard sized, not feared** (`D-M257x-188-5`). The failure mode is a **false** defect — a cited file
under a pruned directory is unfindable, which grades as unresolvable. Over `corpus/**.md` + `CLAUDE.md`:
**2,093** citation-shaped path tokens, of which **1** crosses a pruned name (`.next/static/chunks/*.js`,
a wildcard describing build output, dot-named and therefore pruned by the adjacent rule regardless).
Latent. The pre-registered escalation did not fire.

## Phase B — the repair, and the split that is its point

`_SKIP_DIRS` → `SKIP_DIRS: dict[str, str]`, twelve reasons, with the five subsumed members **kept and
marked** rather than deleted (`D-M257x-188-4`: deleting them would leave a list that is wrong for any
caller pruning without the dot-rule, and would erase the evidence instead of recording it). The fast
membership set is now **derived** from the reasoned one — two literals naming one population is
iter-177's shape, and here it would let the reasons drift from the behaviour while both looked
maintained.

`prune_census()` derives what was actually removed, per member; the guard prints it under **every**
text-mode verb, beside the scope line iter-186 added:

```
  scope            : corpus/services/*.md, corpus/architecture/*.md  ->  41 files
  resolver reach   : 21610 file(s) indexed; 50357 removed by the 12-name prune list
                     (5 of which the adjacent dot-rule already prunes)
                     — pruned: {'node_modules': 3, 'test-results': 2}; inert here: 5
```

**Size is printed, shape is asserted** (`D-M257x-188-3`) — an assertion over 70 % would be vacuous on a
bare checkout and wrong on a populated one, so it would be a fence pinned to an environment.

## Phase C/D — the fence and its mutants

`tests/test_claim_census_skip_registry_m257x.py`, **9 arms in two classes**:

| class | arms |
|---|---|
| `TheShapeIsAsserted` (environment-independent) | every member reasoned · membership set **derived** from the reasoned registry · dot-subsumption marked **both ways** · the walk still applies **both** rules · the registry is the walk's **only** name-based rule |
| `TheSizeIsDerivedAndProvablyNonZero` (`§9`) | `prune_census` can return **non-zero** · it attributes **only** to non-dot members · fired ∪ inert ∪ dot-subsumed **partitions** the registry, pairwise disjoint · the guard **prints** the reach |

**10/10 mutants RED** (`.agentspace/scratch/work-m257x/iter188_mutants.py`, applied to the imported
module so the tree is never edited mid-run):

```
RED ✔ M1  blank a reason              RED ✔ M6  a name hard-coded in the walk
RED ✔ M2  membership set undermined   RED ✔ M7  prune_census can only return zero
RED ✔ M3  unmark a dot member         RED ✔ M8  a dot prune attributed to the registry
RED ✔ M4  mark a non-dot member       RED ✔ M9  the partition leaks
RED ✔ M5  the dot-rule leaves the walk RED ✔ M10 the reach is derived but never printed
```

## Runs — runner and scope named (`§5` r60/75/76)

| scope | runner | result |
|---|---|---|
| `test_claim_census_skip_registry_m257x.py` | unittest 3.14.6 / pytest 8.4.2 (3.9.6) | **9 / 9 passed**, both |
| 6-module guard/registry neighbourhood in `stack-core` | pytest | **120 passed · 0 failed** (27.3 s) |
| same, less the pytest-only module | unittest | **95 passed · 0 failed** (22.3 s) |

**Not covered, stated:** the other Python modules of `stack-core` were not re-run (iter-186's figure
stands, with its scope); the 264 Go tests and 75 TS specs remain **UNMEASURED**; and the 70 % figure is
**this box's** — it is printed, not asserted, precisely because another box will read differently.

## Close — 2026-08-09

**Outcome:** iter-187's selector, run rather than re-surveyed, found the largest silent exclusion in the
milestone's instrument set — the **citation resolver's own walk** prunes **50,357 of 71,967 candidate
files (70.0 %)** on the default run and reported nothing, through a twelve-name list where **five names
can never fire** and **five more were inert**. Reasons supplied, membership derived, subsumption marked
both ways, the reach printed on every verb, 9 arms / 10 mutants RED, and the hazard sized at **1 of
2,093** citation-shaped path tokens.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twentieth consecutive `closed-fixed`; **no
`P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n — (5)
cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-188-1` … `D-M257x-188-6` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter188-the-other-walks-are-unmeasured` — **NEW, with its number.** `_SKIP_DIRS` was one
  of **three** name-based prune rules found in `stack-core` by the same grep: `platform_predicate_guard`
  prunes on `"vendor" in parts or "node_modules" in parts` at **two** sites (`:752`, `:806`) with the
  names **hard-coded inline** — no registry, no reason, no report — and `story_org_count_guard`'s
  `_EXCLUDED_DIRS` is a 4-name frozenset with the same shape. Neither was touched here; both are the
  same class and neither reach is printed.
- `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` — **advanced, not closed.** One member
  worked; the selector is now demonstrated rather than proposed, and the enumeration it runs over (30
  module-level registries in `stack-core`) is the standing denominator.
- `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` (264 Go + 75 TS, still UNMEASURED) ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` (70 collections) ·
  `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open. The standing queue,
  unchanged.

**Lessons:** **a container-keyed exclusion has an environment-dependent SIZE and an environment-independent
SHAPE — print the size, assert the shape.** Fencing the size would have produced a guard that is vacuous
on a bare checkout and wrong on a populated one; refusing to measure it at all is what left a 70 %
exclusion invisible. And the second, which generalises past this list: **a registry can be silently
subsumed by an adjacent unnamed rule** — five of twelve members here could never fire, so the list
over-stated its own work to anyone auditing it. Written into `platform-alignment.md` §8 in this iter's
commit.
