**Type:** tik (standard shape; §9 iter-type refinements consulted, none selected).

# iter-180 — a prose claim about SETS is a claim a machine can grade

**Controlling strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase A — re-derive before repairing

`FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` was measured at iter-177 and
deliberately not landed there (*"a second line of investigation, and an edit to that file after the suite
ran would have invalidated the run this close reports"*). Re-derived here at HEAD `4b60aa2`,
independently, because a correction is re-derived at source before it is filed (§5 iter-22):

| derivation | size | measured relation to `census` |
|---|---|---|
| `guard_family::union` | **27** | `census ∪ CENSUS_EXCLUSIONS` — **exactly** |
| `guard_family::census` | 26 | — |
| `guard_family::declaring_modules` | **26** | differs by **two** members, in **opposite** directions |
| `repair_postcondition::discover_fences` | **26** | **identical to `declaring_modules`**, and REGISTERED |

`CENSUS_EXCLUSIONS = {guard_family}`; the symmetric difference `declaring_modules Δ census` is exactly
`{guard_family, repair_postcondition}`. iter-177's reading is confirmed in full.

**The shared sentence was true of one entry and false of the other** — and their **counts are equal**, so
every count-based comparison of the two reads green (`D-M257x-180-1`). That cardinality coincidence
already has a mutation control in `test_fence_registry_population_m257x.py`; iter-180 shows it was hiding
a second defect on the other side of the repo.

## Phase B — census the class, then give it a grammar

Rewriting the sentence would fix today's reading and leave the class as rottable as before. So:

**Population, derived not listed: 2 of 76** registry entries name a sibling derivation in backticks —
and they are exactly the two under review. Small enough to close, large enough to be a class.

The claim is made executable rather than editorial:

> `RELATION: <module::attr> == <module::attr> [| <module::attr>]`

graded live by `ARationaleThatAssertsASetRelationIsGRADED`, **both directions**: a sibling-naming
rationale with no clause is RED, and a clause that does not hold is RED with the symmetric difference
printed.

### RED-proven before the repair

Both new arms were run against the **unrepaired** tree and failed, naming the two entries and the missing
relation. A fence first seen green over an already-repaired tree is a fence nobody has watched fail
(`§9` iter-149). After the repair: **37 passed** in that module (was 32).

## Phase C — the pre-registered escalation, reported either way

The iter sealed, before any code: *"if the resolver needs a per-site lookup table, stop and keep the prose
repair only."* **It did not fire.** One generic resolver — `module::attr`, called if callable, then
flattened to a set of `str` — covers all five operands across both clauses, including
`repair_postcondition::discover_fences`, which returns a **pair** of lists. No lookup table, no new
registry (`D-M257x-180-3`).

The flatten has its own control, written against the **quiet** direction: a resolver that took only the
first element of that pair would compare a 6-member set against a 26-member one and report a
*real-looking disagreement* rather than raising.

## Phase D — what is deliberately NOT settled

`SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` (same population, one REGISTERED
and one DECLINED) stays open. iter-180 supplies the fact it rests on — the two return **the same set** — and stops:
resolving it changes what the frozen-expectation census treats as a candidate, which is a scope change to
a live instrument, not a rationale fix (`D-M257x-180-4`). What changed is that the survey now rests on a
measurement rather than on the comment that has just been shown to be wrong about its neighbour.

## Runs — scope stated, and what it did NOT cover

| run | result | wall |
|---|---|---|
| `test_frozen_expectation_census_m257x.py` (the module carrying the new arm) | **37 passed** (was 32) | 2.64 s |
| registry-population + registry-completeness + battery-baseline-stage + `test_repair_postcondition.py` + `test_guard_family.py` | **104 passed · 0 failed** | 34.67 s |
| `guard_family.py` over the corpus | **18 GREEN · 0 RED · 8 not-run** (each needs an input not supplied) | — |

Runner **pytest 8.4.2 on `/usr/bin/python3` 3.9.6**. **Not covered:** the rest of `stack-core` (measured
green at iter-179, 1,521 P · 2 S · 0 F, and this iter changes two rationale STRINGS plus adds tests), all
7 mutation batteries — **none of which stages `derivation_registry.py`, checked by importing every
battery's `_COPY_FILES` rather than by assuming** — and the four other rext sections. A scoped green is
evidence about its scope alone (rule 60).

## Close — 2026-08-09

**Outcome:** the false decline rationale is corrected **and the class it belonged to is fenced**. One
sentence covered two entries and was true of one, false of the other in both directions, undetectable by
count because the two sets are the same size. Now every rationale naming a sibling derivation must carry a
`RELATION:` clause and every clause is resolved and asserted live, both directions — RED-proven against
the unrepaired tree before the repair landed. The pre-registered escalation (*a resolver needing a lookup
table means keep the prose fix only*) was checked and did **not** fire.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twelfth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-180-1` … `D-M257x-180-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` — **CLOSED**, at its class rather
  than its sentence.
- `SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` — **open, and now resting on a
  measurement** (`declaring_modules` and `discover_fences` return the identical set) rather than on the
  comment that was wrong about its neighbour. Deliberately not settled here.
- `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` — **NEW.** The grammar grades `==` and
  a `|` on the right. A rationale wanting `⊂`, `∩` or *"differs by exactly {…}"* has no form to say it
  in and would fall back to prose — i.e. to the class this iter just closed. Population today is 2 and
  both are equalities, so widening now would be speculative; **the point is that the ceiling is written
  down rather than discovered later.**
- `SURVEY-M257x-iter179-readme-indexes-test-modules-unmeasured` · `SURVEY-M257x-iter179-thirty-battery-
  tests-unrun` · `FIX-M257x-iter173-ledger-denominator` · the observed half of
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a prose claim about sets is a claim a machine can grade — so give it a grammar instead of
rewriting it.** And the finding that made this iter cheap: **one rationale covering two entries is two
claims wearing one**, which is exactly where a false half survives, because the true half carries the
sentence's credibility. The counts agreeing is what made it invisible; the *shape* is what made it
possible. Written into `platform-alignment.md` §8 in this iter's commit.
