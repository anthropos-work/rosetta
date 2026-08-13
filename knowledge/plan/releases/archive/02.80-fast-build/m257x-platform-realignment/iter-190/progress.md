**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-190 — the census keyed on a SHARED constant, and the broken pair is the one that shares none

## Phase A — replacing a prose selector with a mechanical one

iter-189's routed selector was *a function whose docstring or name claims to be the same derivation as
another*. Re-surveyed: **one hit — iter-189's own founding case.** A selector that finds only the case it
was written from is not an enumeration (`D-M257x-190-1`), so it was replaced with one that needs no prose
and run by AST over all 42 modules of `stack-core`:

> a module-level constant used by BOTH a filesystem-reading function and a git-reading function.

**6 pairs, all in `platform_predicate_guard`:** `_GO_GETENV` (iter-189's), `_TOP_KEY`, `_LIST_ITEM`,
`_INCLUDE_HEAD`, `_REPOS_YML_ENTRY`, `_REF_PINNED`.

**All 6 agree today**, measured against the real clones: `parse_compose` → 7 services;
`compose_counts_at(None)` → `(5, 7)`; `repos_yml_history` ⊇ `parse_repos_yml` with 0 current entries
missing. A ZERO — so `§9` applies, and proving the instrument is what found the defect.

## The finding — the pair a shared-constant census structurally cannot see

`_parse_one_compose` and `compose_counts_at` read **the same construct out of the same file** with
**different regexes**, so they share no constant:

| recogniser | pattern | first character |
|---|---|---|
| `_SVC_KEY` → `parse_compose` (the topology **G1/G7/G8** grade against) | `^  ([A-Za-z][A-Za-z0-9_.-]*):` | **letter only** |
| `_COMPOSE_SERVICE_KEY` → `compose_counts_at` (**G10**) | `^  (?P<name>[A-Za-z0-9_.-]+):` | any of `[A-Za-z0-9_.-]` |

Over a 9-name table they disagree on **5**: `3d-render`, `_internal`, `-legacy`, `.hidden`, `9front`.
Compose's own charset admits a leading digit, so the narrow one is wrong — and the **direction** settles
it independently (`D-M257x-190-3`): a service the topology cannot see is *absent*, so claims about it read
**UNREACHED rather than graded**, while G10 counts it. Under-count on the side that grades.

That is iter-184's rule — *a fence's POPULATION is a registry too* — committed by the instrument written
this iter, which is precisely where iter-184 said to look.

## Phase B — one charset, not two matching literals

`_SVC_NAME = r"[A-Za-z0-9_.-]+"`; both patterns are built from it. Making the two literals agree would
leave iter-177's shape in place: **agreement today is not the property; sharing the source is.**

## Phase C/D — the fence and its mutants

`tests/test_dual_reader_parity_m257x.py`, **10 arms in three classes**:

| class | arms |
|---|---|
| `TheEnumerationIsDeclaredBothWays` | every derived pair **declared** · no declaration **outlives its pair** · every declaration carries a reason · **the derivation fires** (`§9`) |
| `TheRECOGNISERSAgree` (the half no shared-constant census can see) | the two recognisers accept the **same names** · both derived from **one charset** · the charset admits a **leading digit** |
| `ThePairsAreActuallyCompared` | compose topology **vs** compose counts on the effective set · the fixture would **catch** the iter-190 defect · `repos.yml` **file vs history** |

`_REF_PINNED` is declared **NOT A PAIR** with its reason — a shared *predicate*, not a shared population
(`documented_profile_tokens` matches ref pins in prose; `ref_resolves_in` asks git whether a sha
resolves). iter-185's population-vs-predicate split, applied rather than re-derived.

**Fixtures are synthetic on purpose** (`D-M257x-190-5`): a comparison that only runs where a clone
happens to exist stops being run, and would pass by absence on a bare checkout. The compose fixture
carries `9front` deliberately — without a non-letter-initial name every arm is green under **either**
recogniser.

**6/6 mutants RED**, including the one that makes the point:

```
RED ✔ M1 a derived pair left undeclared      RED ✔ M4 the derivation stops firing
RED ✔ M2 a declaration outlives its pair     RED ✔ M5 the narrow recogniser returns (3 arms, 5 subtests)
RED ✔ M3 a declaration loses its reason      RED ✔ M6 narrow recogniser -> the two compose readers
                                                       disagree on the fixture (2 arms)
```

## Runs — runner and scope named (`§5` r60/75/76)

| scope | runner | result |
|---|---|---|
| `test_dual_reader_parity_m257x.py` | unittest 3.14.6 / pytest 8.4.2 (3.9.6) | **10 / 10 passed**, both |
| + `test_rpc_reader_parity` + `test_platform_predicate_guard` + `test_platform_alignment_guard` | pytest | **258 passed · 0 failed** (15.0 s) |
| + `test_guard_family` (less the alignment module) | unittest | **239 passed · 0 failed** (17.7 s) |
| `parse_compose` / `compose_counts_at` against `stack-demo/platform`, before **and** after | — | **7 services, identical set**; counts `(5, 7)` unchanged |

**Not covered, stated:** `story_org_count_guard._EXCLUDED_DIRS` (iter-188's third member) is still
untouched; 264 Go + 75 TS still UNMEASURED; the enumeration covers `stack-core` only — the other ten
sections have not been asked this question.

## Close — 2026-08-09

**Outcome:** the dual-reader question was **enumerated** rather than sampled — 6 pairs by AST across 42
modules, all agreeing — and the `§9` obligation to prove an instrument that returns zero is what surfaced
the defect: **two readers of one compose file recognising a service name with different regexes**,
disagreeing on 5 of 9 candidate names, with the *under-counting* one driving the topology G1/G7/G8 grade
against. One charset now, and an enumeration that keeps running with its population declared both ways.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-second consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-190-1` … `D-M257x-190-5` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven` — **NEW, with its
  denominator.** The AST enumeration runs over `stack-core`'s 42 modules and **no others**; the ten
  remaining sections have never been asked. This is iter-186's scope defect in a fence one week old, and
  it is stated here rather than discovered later.
- `SURVEY-M257x-iter190-one-construct-two-regexes-is-unenumerated` — **NEW.** The defect this iter found
  was reached by proving the instrument, not by the instrument. *Two recognisers of one construct that
  share no constant* has no census at all; the selector would be semantic (do two patterns describe the
  same thing?) and is genuinely harder than the shared-constant one.
- `SURVEY-M257x-iter189-the-parity-question-is-unasked-for-every-other-dual-reader` — **CLOSED by
  enumeration** for `stack-core`: 6 pairs, all declared, all compared, fenced both ways. Its residual is
  the first two routes above, which is the honest split.
- `SURVEY-M257x-iter188-the-other-walks-are-unmeasured` — advanced; `story_org_count_guard._EXCLUDED_DIRS`
  still untouched.
- `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
  `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` · `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open. Standing queue unchanged.

**Lessons:** **a selector that finds only the case it was written from is not an enumeration** — iter-189's
prose selector returned exactly its own founding hit, and replacing it with an AST rule turned one anecdote
into six pairs and a declared population. And the sharper half: **proving the instrument is where the
defect was**, because the pair a shared-constant census cannot see is the pair whose two readers spell the
same construct differently. Written into `platform-alignment.md` §8 in this iter's commit.
