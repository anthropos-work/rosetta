**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), working a class in
the order the strategy fixes: descending measured size, denominator stated. See
[`platform-alignment.md` §8](../../../../../corpus/ops/platform-alignment.md).

# iter-183 — the milestone's own backlog is a registry, and it was the only one with no fence

## Phase A — census

`stack-core` ships fences over the corpus index, the platform's own config, the demo knobs, the anchors,
the retractions, the derived counts, the guard family and the guards' own tests. It ships **none** over
the carry-forward **route queue** — the artifact that decides what the next iter works on and the one
every sub-agent brief is assembled from.

| derivation (m257x, 2026-08-09) | value |
|---|---|
| distinct route ids under `iter-*/progress.md` | **206** |
| routes with at least one recorded disposition | **183** |
| total (route × iter) dispositions | **779** |
| routes carrying a FULL closure (grammar-aware) | **32** |
| fences whose subject is this registry | **0** |

**The reading hazard came first, and it is measured.** These ids are long enough that markdown
hard-wraps them mid-token with a trailing hyphen (`SURVEY-M257x-iter179-thirty-battery-` /
`tests-unrun`). A line-scoped reader returns **207 ids where 204 exist** — 3 phantoms, 3 real ids
invisible. Every prior statement this milestone made about "the open routes" was made under it.

## Phase B — adjudicate, one candidate at a time

A first detector that bound a disposition to the whole BULLET returned **6**. Adjudicated individually
(`D-M257x-183-3`), **1 is true and 5 are false**, each falsifiable for a different, nameable reason:
three multi-id bullets whose `CLOSED` belongs to a neighbour, one TABLE row whose verdict is
*superseded*, and one **half** closure that iter-174 correctly scoped to *"the observed half"*.

Booking 6 would have been `D-M257x-122-3` in miniature. A detector rewritten to the tokens the registry
**already writes** returns exactly **1** — the same 1 the hand adjudication reached, independently.

**The true one:** `SURVEY-M257x-iter170-cockpit-runner-dependence`, **CLOSED BY ROOT CAUSE at iter-171**
(`D-M257x-171-1` — a blocking `socket.getfqdn` in shipped `cockpit.py`, 35.005 s on the 3.14 interpreter
against a 2 s bind window), re-published as *"unchanged; open"* by **iter-182**, eleven iters later,
describing the mechanism iter-171 had already attributed and repaired.

**A stale carry-forward and a legitimate re-open are the same sentence**, so the verdict was settled by
re-running the subject, not by reading the record: `unittest`/3.9.6, `unittest`/3.14.6 and
`pytest`/3.9.6 each run **207** `test_cockpit` tests with **0 failures**. `207 = 207 = 207`. It is a
**correction**, landed as an in-place `✅ CORRECTED — iter-183` annotation on iter-182's bullet — never a
rewrite of a closed record, never a re-open the evidence does not support.

## Phase C — fence

[`route_disposition_guard.py`](../../../../../.agentspace/rosetta-extensions/stack-core/) — a route
recorded `CLOSED` at iter K may not be asserted **open** at iter M > K unless the later block says why
in the registry's own grammar (re-open · supersede · half-closure · in-place `CORRECTED`). It sweeps
every planned milestone, is in `guard_family`'s `INVOCATIONS` (so it cannot be dropped in either
direction), stamps its tree, and is indexed in the README with its test module.

**Its docstring's first draft asserted the live registry never writes two ids in one segment. Measured:
28 do, and 2 of those carry a closure verdict** — the claim was plausible, unmeasured, and shipped in
the same commit as the instrument that could have checked it. So the guard applies `D-M257x-122-5` at
segment grain: a multi-id closure is **REFUSED, never resolved by proximity**, and the refusal count
**prints on every run** so the blind spot is visible rather than silent (`D-M257x-183-4`).

**Controls, 20, both families able to fire.** Mutation: the closed→open shape at distance 1 and 8; every
grammar token proven to clear it; the segment rule, the table exclusion and the multi-id refusal each
pinned. Anti-vacuity: written against the subject and **scoped** — zero closures across the default
sweep is exit 2, but one milestone may legitimately close nothing (`m256` and `m257` both do), so the
control names the sweep it applies to (`D-M257x-183-5`). Pure `unittest`, **20/20 under both runners**.

## Phase D — measure

| gate | result |
|---|---|
| `route_disposition_guard --repo-root` | **OK**, 3 milestones, 0 contradictions, 3 wrapped ids rejoined, 2 segments refused |
| `stack-core` suite, `pytest` 3.9.6 | **1,594 passed · 2 failed · 2 skipped**, from 1,553 P at iter-178 |
| the 2 failures | **both this fence's own registration debt**, both repaired → 54/54 green |
| new controls, `unittest` 3.9.6 / 3.14.6 | **20 / 20** |

The two failures are the finding worth keeping: `test_fence_provenance` and
`test_fence_registry_population_m257x` **both fired on the new fence**, correctly — a fence that does
not stamp its tree, and a published index disclosure that is a measurement and therefore moves. Phase 0d
predicted one registration (`INVOCATIONS`); the suite named two more. **A pre-flight that finds one
precondition has not established there is only one.**

## Close — 2026-08-09

**Outcome:** the backlog the whole milestone is steered by now has a fence. 183 routes / 779
dispositions / 32 closures enumerated, **1 live contradiction found and repaired** — a route CLOSED by
root cause at iter-171 and re-published as open at iter-182 — with the verdict settled by re-running the
subject (`207 = 207 = 207` across three runner combinations), not by reading the record. **5 of the 6
raw candidates were refuted**, and the grammar that refutes them is the deliverable.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (fifteenth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` fixes the sweep
order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-183-1` … `D-M257x-183-6` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none. The `stamp_main()` call and the disclosure-triple update are this fence's
own registration cost, not unrelated fixes.

**Routes carried forward:**
- `SURVEY-M257x-iter170-cockpit-runner-dependence` — **CLOSED, again, and this time asserted.** It was
  closed at iter-171; iter-182 re-published it as open; the correction is annotated in place and the
  contradiction class is now fenced, so this cannot recur silently.
- `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` — **NEW.** The grammar binds a
  verdict to its ` · ` segment, so a segment naming several routes is refused rather than resolved. 28
  such segments exist and **2 carry a closure verdict**, which the guard therefore does not see.
  Widening it to bind a verdict to the id it FOLLOWS is a real improvement and a real risk; the ceiling
  is written down rather than discovered later (iter-180's rule).
- `SURVEY-M257x-iter183-only-ONE-registry-property-is-asserted` — **NEW.** This fence asserts
  disposition consistency. It does **not** assert that every route has a birth, that a `NEW` route is
  ever revisited, or that an id in a brief resolves to a live route — the three properties an *orphan*
  in this queue would violate. The population is enumerated; the property set is not.
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — unchanged and open, **but its figure moved**:
  the live triple is now **`17 of 28` (union) · `17 of 27` (census) · `16 of 27` (declaring)**. The id
  is not renamed (iter-177 settled that it embeds a correct reading); what is open is still *which*
  derivation the index should be complete against.
- `SURVEY-M257x-iter179-thirty-battery-tests-unrun` (owner: the next harden pass) ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  the observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **an instrument's own backlog is a registry, and it rots exactly like the ones it watches.**
Eleven layers of fencing over the corpus, the platform and the guards themselves — and the queue
deciding which of them to build next had no assertion at all, while a consumer (the run-17 brief) had
already noticed by hedging. Two corollaries paid for directly: **write the fence in the registry's own
grammar** — 164 blocks were already saying `half CLOSED`, `supersedes`, ` · `, and reading that grammar
cost less than imposing a new one and left the prose readable; and **a pre-flight that finds one
precondition has not established there is only one** — Phase 0d named `INVOCATIONS`, the suite named two
more. Written into `platform-alignment.md` §8 in this iter's commit.
