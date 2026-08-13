**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
*census the mechanical classes; stop sampling them*.

# iter-187 — the exclusion registry sits at SECTION grain; the runner's reach sits at (section × language) grain

## Phase A — the enumeration

iter-186 fixed the section list and, in doing so, made the finer question askable for the first time. The
matrix, measured over `.agentspace/rosetta-extensions` @ `4afbd6c` (2026-08-09):

| section | `test_*.py` | `*_test.py` | `*_test.go` | `*.spec.ts` | iter-186 bucket |
|---|---:|---:|---:|---:|---|
| alignment | 0 | 0 | 21 | 0 | excluded |
| clerkenstein | 0 | 0 | 37 | 0 | excluded |
| demo-stack | 35 | 0 | 0 | 0 | **collected** |
| dev-stack | 5 | 0 | 0 | 0 | **collected** |
| playthroughs | 0 | 0 | 22 | 45 | excluded |
| stack-core | 67 | 0 | 0 | 0 | **collected** |
| stack-injection | 7 | 0 | 0 | 0 | **collected** |
| stack-secrets | 0 | 0 | 20 | 0 | excluded |
| stack-seeding | 0 | 0 | 119 | 0 | excluded |
| stack-snapshot | 0 | 0 | 45 | 0 | excluded |
| **stack-verify** | 5 | 0 | 0 | **30** | **collected** |

**`stack-verify` is COLLECTED and carries 30 Playwright `*.spec.ts`.** Every arm of iter-186's fence is
green over it and each one is *right*: it is in exactly one bucket, the buckets partition the disk, and it
does carry Python tests. The arms check **membership**; the hole is **inside a member**. The printed line
`scope: 5 of 11 sections — Python only` therefore presents 30 specs as read.

Two numbers follow, and only the second is new:

- iter-186's **`264 Go + 45 TS`** is **true and stays** — it is scoped to *the six excluded sections* at
  every site (`D-M257x-187-3`).
- **Nothing published a repo-wide figure**, so the six-section one was the only number available to
  quote. Repo-wide the unread non-Python population is **264 Go + 75 TS**.

The pre-registered escalation was checked and did not fire: no instrument in this milestone reads those
30 specs (`D-M257x-187-5`). The two `stack-core` Python files that even mention Playwright are a keyword
in a vocabulary list and a citation fixture.

## Phase B — the repair, declared not widened

`UNREAD_IN_COLLECTED` (declared: section → reason) + `unread_non_python()` (derived from disk over
`NON_PYTHON_TEST_GLOBS`). iter-150's split, the third use of the idiom in this one file. The scope block
now prints, ahead of any total:

```
  scope: 5 of 11 sections — Python only. 6 section(s) excluded BY LANGUAGE, not by absence:
    …
  within those 5 sections, 30 non-Python test file(s) are ALSO unread — collected by section, not by language:
    - stack-verify: 30 file(s) — 30 Playwright `*.spec.ts` under `e2e/tests/` — driven by `npx playwright test`, …
  Any total below is a statement about the PYTHON tests of those 5 sections and about no others (`§5` rule 60).
```

**What the census runs is unchanged** — widening it was considered and rejected (`D-M257x-187-2`): it
would change what the tool does under cover of fixing what it says.

**Second, latent, and sized before repair.** Harden pass 42's presence arm globbed `test_*.py` **and**
`*_test.py` while the collector globs `test_*.py` alone — a **superset** predicate, so a section whose
Python tests were all spelled `*_test.py` passes the arm and still contributes the silent zero the arm
exists to prevent (`§5` r70/71, the defect re-entering through the spelling). **Hazard size measured, not
asserted: 0 such files across all 11 sections.** Latent. Closed anyway, by re-basing the arm on
`S.modules()` and adding a control that the collector still reads the named `PYTHON_TEST_GLOB` constant
(`D-M257x-187-4`).

## Phase C/D — the fence and its mutants

New class `TheEXCLUSIONIsFinerGrainedThanTheSection` (5 arms) + 1 control arm in the existing class:

| arm | property |
|---|---|
| `…NON_PYTHON_file_in_a_collected_section_is_DECLARED` | undeclared within-section remainder → RED |
| `…no_DECLARATION_outlives_its_subject` | stale declaration → RED |
| `…every_within_section_DECLARATION_carries_a_reason` | `§5` rule 8 |
| `…derived_side_reads_DISK_and_covers_every_non_python_spelling` | **the `§9` instrument control** — both directional arms are satisfied by a derivation that returns `{}`, which is also what a broken one returns |
| `…repo_wide_unread_population_is_STATED_not_implied` | the six-section and repo-wide figures both stay readable |
| `…presence_predicate_IS_the_collector_not_a_restatement` | the collector must keep reading the named glob constant |

**6/6 mutants RED** (`.agentspace/scratch/work-m257x/iter187_mutants.py`, applied to the imported module
so the tree is never edited mid-run — `D-M257x-187-6` states what that does and does not prove):

```
RED ✔ M1 drop the stack-verify declaration          RED ✔ M4 drop *.spec.ts from the glob tuple (2 arms)
RED ✔ M2 declare a section with no subject          RED ✔ M5 collector restates the glob
RED ✔ M3 blank the reason                           RED ✔ M6 collector sees one section only
```

## Runs — with their runner and their scope named (`§5` r60/75/76)

| scope | runner | result |
|---|---|---|
| `test_suite_census_population.py` | unittest (3.14.6) / pytest (fleet 3.9.6) | **13 / 13 passed**, 0 failed, both |
| + `test_suite_census.py` | both | **25 passed**, 0 failed (2.5 s each) |
| 6-module registry/fence neighbourhood in `stack-core` | pytest 8.4.2 / 3.9.6 | **123 passed · 0 failed** (18.45 s) |
| same 6 modules | unittest 3.14.6 | **98 passed · 0 failed · 1 module unloadable** |

The unloadable module is `test_claim_census_guard.py` — `import pytest` under an interpreter that has
none. **Pre-existing and already documented** in `suite_census.py`'s own header as one of the modules only
the fleet runner can load; it is a runner disagreement, not a regression, and the diff does not touch it.

**Not covered, stated rather than implied:** the other 5 sections of `stack-core`'s Python suite were not
re-run (iter-186's figure stands, with its scope); the 264 Go tests and 75 TS specs remain **UNMEASURED**
— this iter names 30 more of them than were nameable before, and reads none.

## Close — 2026-08-09

**Outcome:** the registry iter-186 built to stop a silent exclusion had one **inside a member**:
`stack-verify` is declared collected, carries 5 Python modules, and also carries **30 Playwright specs no
instrument in this milestone reads**. Declared at (section × language) grain, completeness derived from
disk both ways, printed ahead of every total, 6/6 mutants RED — and the repo-wide unread non-Python
population is stated for the first time: **264 Go + 75 TS**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (nineteenth consecutive `closed-fixed`; **no
`P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n — (5)
cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-187-1` … `D-M257x-187-6` (see [`decisions.md`](decisions.md))

**Side-deliverables:** one, and it is a consequence of this iter's own close rather than an unrelated
find. Writing the routes block turned `test_harden_origin_route_visibility_m257x.py` RED: its
`LEDGER_ONLY_DISPOSITIONS` registry carries harden-origin routes *no iter ever cited*, and this close
cites `FIX-M257x-h36-labeled-prover-denominator`. **Dispositioned by removing the entry, not by dropping
the citation** (`D-M257x-187-7`) — suppressing the mention to keep a fence green would re-hide an open
route, which is the defect that module exists to prevent. The route stays OPEN and Fate 3. Fences green
after: **43 passed · 0 failed** over `test_route_disposition_guard` + `test_harden_origin_route_visibility_m257x`
+ `test_suite_census_population`, under **both** runners; `route_disposition_guard` itself `OK` — 182
iters · 365 route ids · 1,202 dispositions · **0 contradictions**.

**Routes carried forward:**
- `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` — **NEW.** This iter found its member by
  asking *"at what grain is this exclusion actually true?"* of one registry. Every other declared
  registry in `stack-core` — `ENV_GATED`, `LANGUAGE_EXCLUDED_SECTIONS`, the fence-family registries — has
  the same question outstanding, and **none has been asked it**. Unlike iter-185's residual this one has
  a mechanical selector: *a registry keyed by a CONTAINER whose justifying reason is a property of the
  container's CONTENTS.*
- `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` — **unchanged and now larger**: the unread
  non-Python population is **264 Go + 75 TS**, not 264 + 45. Still UNMEASURED.
- `SURVEY-M257x-iter185-other-declared-populations-unaudited` — **untouched**; deliberately substituted
  rather than worked (`D-M257x-187-1`). 70 collections still need population-vs-predicate classification.
- `D-M257x-145-3` — unchanged, still the user's to rule; this iter sharpens its second axis (the five
  collected sections are not fully read *either*).
- `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` · `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` ·
  `FIX-M257x-iter173-ledger-denominator` · `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` · `FIX-M257x-h36-labeled-prover-denominator`
  — unchanged; open. The standing queue, unchanged.

**Lessons:** **a membership check cannot see a hole inside a member.** When an exclusion is justified by a
property of the *files* but recorded against a *container*, the registry is coarser than what it claims to
describe — and iter-186's four mutation-proven arms were all green over the counter-example while it sat
in the collected half. Ask the grain question of every registry whose reason is a content property.
And the publishing corollary: **a correctly-scoped number is still a trap when it is the only number in
the room** — publish the total whose denominator a reader will assume. Written into
`platform-alignment.md` §8 in this iter's commit.
