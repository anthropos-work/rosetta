# iter-176 — progress

**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase 1 — the gap iter-175 left, stated as `TOK-08` states it

`TOK-08`: *"build or extend a fence that **enumerates every instance in the corpus**, run it to zero, and
**keep it green**."* iter-175 did the first clause with a scratch script it deliberately did not check in
— correct for an iter whose deliverable was the union repair, and exactly the gap here.

> **An instrument that ran once and was deleted is a reading wearing a fence's clothes.**

Re-surveyed at `3b5a82d` / rext `5b108d0`: 5 sites, unchanged in shape, one row moved by iter-175's own
repair (`guard_family` 24 → **25** names, `predicate_enumerator` having joined).

**Why the population, not the last member** — the discovery record is the argument:

| iter | registry found | found by |
|---|---|---|
| 173 | `guard_family` · `derivation_registry` · `fence_provenance` | grep for a sibling's name |
| 173 | `repair_postcondition_baseline.json` (4th) | whole-suite run, after the fact |
| 174 | the mechanical-fences battery's fence-seed list (5th) | whole-suite run, one iter later |
| 175 | `derivation_registry` again, and a 6th (prose) by hand | whole-suite run + a hand check |

**Four consecutive discoveries by the most expensive instrument in the milestone.** That is not bad luck;
it is the signature of a population nobody enumerates.

## Phase 2 — the fence

`stack-core/tests/test_fence_registry_population_m257x.py` — **7 tests, 0 findings over 5 sites.**

**Predicate:** a *collection literal* holding ≥2 fence-module names, where the fence population is
derived from `FENCE_KIND` declarations through `repair_postcondition.declared_kind` — the same reader
`discover_fences()` and `guard_family.declaring_modules()` use. **One reader; a private copy here would
have been a fourth derivation of this population**, in the fence written about that.

**Three design calls, each recorded with what it cost:**

- **A test, not a `*_guard.py`** (`D-M257x-176-1`). A new guard must join three registries; this one
  joins **zero**, so shipping the fence that counts registries does not increment the count it reports.
  iter-157 made the same call for its sibling — precedent, not invention.
- **`ast.Call` arguments are in the predicate** (`D-M257x-176-3`). The seed list that started this whole
  thread is not a literal — it is `battery_stage.local_deps(STACK_CORE, "markdown_structure_guard.py", …)`,
  a **call**. A literals-only predicate is tidier, defensible in one sentence, and **structurally blind to
  the registry the fence was commissioned for.**
- **The `fixtures` exclusion was removed after measuring it** (`D-M257x-176-4`) — **5 sites with it, 5
  without.** An exclusion that changes nothing buys nothing and silently forecloses a registry living in
  a fixture. Adding an unmeasured narrowing *inside the fence written against unmeasured narrowings* was
  the one mistake this iter could not afford.

**The five, classified per site with a reason** (`D-M257x-159-4` — explicit, never inferred; keyed by
**path**, never line, `D-M257x-176-2`):

| site | verdict |
|---|---|
| `guard_family.py` | `REGISTRY:reconciled` — `INVOCATIONS`, checked against `census()` both ways; a missing member is exit 2 naming itself |
| `repair_postcondition_baseline.json` | `REGISTRY:ratchet` — `--accept` registers or lowers, never raises |
| `test_m257x_mechanical_fences_mutation_battery.py` | `REGISTRY:hand-maintained` — reported, but only by a ~14-min battery one iter late (`FIX-M257x-iter174-…`, **still open and now recorded IN the fence**) |
| `test_iter45_mechanical_fences.py` | `DECLINE:subject` — the three fences iter-45 shipped; a fourth joining breaks nothing here. (Its sibling `test_21` in the same file **was** a registry and was derived at iter-118 — the distinction is real and was drawn in anger) |
| `test_fence_registry_completeness_m257x.py` | `DECLINE:regression-pin` — the two modules that exposed iter-157's defect, asserting **enrolment** rather than the filename |

**Both arms RED-proofed before the suite, in-process, without editing the tree** (`D-M257x-176-6`):
dropping a `DECISIONS` entry goes RED naming that path; adding a key that holds no such set goes RED
naming it as gone.

**Two mutation controls, each pinned to something that actually happened:**

1. **the JSON arm is load-bearing** — the registry that went four iters unnamed *is* a JSON object and a
   python-only walk cannot see it; deleting the arm must LOSE a site, and every lost site must be JSON.
2. **the ≥2 floor is load-bearing** — at ≥1 the predicate becomes iter-175's rejected **39**-site
   instrument arriving by the back door; raising the floor to 3 must lose a site, or the 2 is decoration.

Plus **anti-vacuity** (§9): ≥20 fences declared and ≥3 sites found, or both arms are comparing empty sets.

**And the limit is disclosed as a TEST, not a sentence** (`D-M257x-176-5`). A prose index is not a
collection literal, so this fence is blind to `stack-core/README.md` — measured by iter-175 at **16 of
27**. `test_the_disclosed_limit_is_STATED_not_assumed` asserts the README really does name several fences
and really is not a site, so a widening of the predicate fails the test rather than quietly falsifying
the docstring.

## Phase 3 — the whole-population run

`FIX-M257x-iter142-whole-suite-owed`. Five sections, pytest / `/usr/bin/python3` 3.9.6:

> **`9 failed · 3365 passed · 3 skipped in 2162.17 s (36:02)`** — **executed = 3377** (iter-172's unit
> rule; the three are summed, `deselected` excluded).

**Zero self-inflicted failures** — the first iter in four where the whole-population run found nothing
this iter had introduced. The delta against iter-175's run reconciles exactly, which is the only reason
it is quoted at all:

```
3355 passed → 3365   =  +7 (this fence's tests)  +3 (iter-175's own REDs, repaired in that iter)
  12 failed →    9   =  −3 (the same three)
```

All **9** remaining are the pre-existing routed pair, re-confirmed by reading the failures rather than
assumed: **6** live-clone sha-pinned demopatch tests (`FIX-M257x-iter145-sha-baseline-drift` — **the shas
must not be re-pinned**) and **3** needing a host Postgres
(`FIX-M257x-iter145-migrate-race-needs-a-host-postgres`). Identical membership to iter-175's run.

## Close — 2026-08-09

**Outcome:** iter-175's instrument is now a **checked-in fence**: every collection literal in rext holding
two or more fence-module names is enumerated and must be classified `REGISTRY:<what keeps it in sync>` or
`DECLINE:<class>: <reason>`, with an unclassified site RED. **5 sites, 0 findings, 7 tests** — two arms
RED-proofed, two mutation controls pinned to real history, one anti-vacuity control, and the prose-index
blind spot disclosed as an executable test rather than a sentence. Discovery of registry #7 moves from a
36-minute suite run four iters later to a sub-second static check, and
`FIX-M257x-iter174-accept-registers-one-registry-of-two` is **closed at its population** — the member it
named stays open *inside the fence*, which is where an open obligation belongs.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (eighth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean)
**Decisions:** `D-M257x-176-1` … `D-M257x-176-6` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-iter174-accept-registers-one-registry-of-two` — **narrowed, not closed.** Its *class* is now
  fenced; its named member (the battery seed list a `--accept` does not write) is unchanged, and is now
  recorded as the verdict text of that site so the fence itself carries the open obligation.
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — unchanged, and now **pinned by a test** as this
  fence's disclosed blind spot.
- `SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` — unchanged; open.
- `FIX-M257x-iter173-ledger-denominator` — unchanged; open (owned by the next harden pass).
- The observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — unchanged; open.
- `FIX-M257x-iter145-sha-baseline-drift` / `-migrate-race-needs-a-host-postgres` — re-confirmed at 9 REDs,
  identical membership to iter-175's run; unchanged.
- The standing queue, unchanged.

**Lessons:** **a class is not closed by a repair; it is closed by an enumeration that keeps running.**
iter-175 repaired the biggest member and *measured* the population — and the population would have gone
back to being remembered the moment the scratch script was deleted. The tell was already in the record:
four consecutive registries found by a 34-minute suite. **When the most expensive instrument you own is
the one making the discoveries, the cheap instrument does not exist yet.** Two corollaries this iter paid
for directly: the predicate must include the shape the motivating case actually has (`ast.Call`, not just
literals — elegance that cannot see its own commissioning case is not elegance), and **an exclusion must
be measured before it ships**, because one that changes nothing is a narrowing with no upside.
