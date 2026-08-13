**Type:** tik — under `TOK-08`, working the second mechanical property of the population iter-183
enumerated. See [`platform-alignment.md` §8](../../../../../corpus/ops/platform-alignment.md).

# iter-184 — a carry-forward that names a SET must enumerate it, and the fence that said so was reading 57.8 % of its own subject

## Phase A — the glob

iter-183 routed *"only ONE registry property is asserted."* Probing that residual found a member of the
fence's own population that is not a route: **`SURVEY-M257x-iter181-`**, the stem left behind by
iter-182's `` `SURVEY-M257x-iter181-*` ``. Measured: **iter-181 created zero routes** — its close names
seven ids, all pre-existing. **The glob denotes ∅** and has read as a live backlog item in every brief
that quotes this queue, including run 17's (`D-M257x-184-1`).

## Phase B — repair by withdrawal, not deletion

Deleting it would rewrite a closed iter's published record. The corpus already has the idiom —
`CLAUDE.md` writes `~~ai~~`, `~~authn~~`, `~~GraphQL/Cosmo Router~~`, and this registry used it at
iter-29 — so the glob is **struck**, the correction annotated beneath it, and the guard learns to read
strikethrough as withdrawal (`D-M257x-184-2`). Two safeguards, because a stripper that hides a live
route is this fence's own failure mode: the struck-span count **prints every run**, and the pattern is
**not** `re.S` (a DOTALL non-greedy span would let two unrelated `~~` marks swallow every id between
them). Audited: exactly 2 struck spans exist, withdrawing the glob and one iter-29 `CHECK-` id — **no
live route**.

## Phase C — and then the population itself

Auditing the strikethrough surfaced a `CHECK-` id **that the fence's grammar does not recognise at all**.
Censused inside the carry-forward blocks: **8 kinds are carried that iter-183's declared tuple excluded**
— `CHECK` (76 distinct), `DOC` (28), `FENCE` (15), `MEASURE` (10), `DEF` (3), `HOST`/`REPOINT`/`READ` (2
each) — while **`PROBE` and `TASK`, two of its five declared kinds, never occur.**

**189 of 327 distinct carried ids — 57.8 %.** The fence shipped one iter ago, described as *"the route
registry"*, green, with green controls and a confident census line. It is `§2`'s hand-maintained tuple
returning **inside the fence built to stop registries rotting** (`D-M257x-184-3`).

Repaired by removing the tuple rather than extending it — the kind is now **derived**. And the
well-formedness predicate was narrowed in the same pass: its first form booked `HOST-M257x-toolchain` as
malformed, a legitimate id whose kind is not iter-scoped, so it now asserts **exactly** the defect —
**an id may not end in `-`** (`D-M257x-184-4`). That error was visible only *because* the population had
just been widened to contain a counter-example.

| reading | population | dispositions | closures | contradictions |
|---|---|---|---|---|
| iter-183, kind **declared** | 184 | 789 | 33 | 0 |
| iter-184, kind **derived** | **312** | **1,156** | **36** | **0** |

Both are correct about their own population and neither may be quoted without it (iter-177). Widening
surfaced **no** new contradiction — the consistency property holds on the wider set, so honesty was the
only cost.

## Phase D — measure

| gate | result |
|---|---|
| `route_disposition_guard --repo-root`, pre-repair | **EXIT 1**, naming `'SURVEY-M257x-iter181-' … carried at iter-182` |
| same, post-repair | **OK** — 3 milestones, 0 malformed, 0 contradictions |
| controls, `unittest` 3.9.6 / 3.14.6 | **25 / 25** on both (was 20) |
| fence + both registry guards + guard-family suite | **124 passed · 0 failed** (126.77 s, pytest 8.4.2 / 3.9.6) |
| new `*_guard.py` on disk | **0** — the arm extends the existing guard, so no `INVOCATIONS` entry, no new stamp, and the README index triple **does not move** (Phase 0d's prediction, checked not assumed) |

Not covered: the rest of `stack-core` (1,594 P at iter-183, of which the 2 F were repaired), the 7
batteries (neither touched file is staged by any), the four other rext sections.

## Close — 2026-08-09

**Outcome:** the fence built at iter-183 was measuring **57.8 %** of its own subject and said nothing —
its id KIND was a hand-maintained tuple containing two kinds that do not exist and missing eight that
do, `CHECK` alone accounting for 76 distinct carried ids. Kind now **derived**; the registry reads
**312 routes / 1,156 dispositions / 36 closures / 0 contradictions**. Also closed the property that
found it: a carried member must be a well-formed route id, RED-proven live on `SURVEY-M257x-iter181-` —
a **glob denoting the empty set** that has been in every brief quoting this queue.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (sixteenth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`) — (3) re-scope: n — (4)
user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome:
**continue**
**Decisions:** `D-M257x-184-1` … `D-M257x-184-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter183-only-ONE-registry-property-is-asserted` — **half CLOSED.** Well-formedness is
  now asserted alongside disposition consistency. The other properties it named — every route has a
  birth; a `NEW` route is ever revisited; an id in a brief resolves to a live route — stay open.
- `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` — **NEW, and it carries its
  measurement rather than a mood** (iter-178's rule). **17 of 165** carry-forward blocks discharge the
  remainder with `"The standing queue, unchanged"`, which names no member. Under it, **155 routes have a
  last disposition that is not a closure**, and **117 of those have not been named in 20+ iters** — 3 of
  them silent for **180+**. Whether those are open, satisfied-in-passing, or abandoned is **not decidable
  from the registry**, and a wildcard makes it undecidable by construction. Deliberately not repaired
  here: this is a policy question about how a backlog discharges, not a malformed id.
- `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` — unchanged, **and the count moved
  with the population**: 3 ambiguous segments are now refused, not 2.
- `SURVEY-M257x-iter179-thirty-battery-tests-unrun` (owner: the next harden pass) ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` (live triple now `17/28 · 17/27 · 16/27`) ·
  the observed half of `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a fence's POPULATION is a registry too, and it is the one nobody audits.** iter-183
shipped green, with green mutation and anti-vacuity controls and a printed denominator, while reading
**57.8 %** of its subject — because the property under test (*is this registry consistent?*) was fenced
and the property that decides **what the registry IS** was a tuple typed from memory. Every control
fired correctly on the 57.8 %. Two corollaries paid for directly: **derive the population, never declare
it** — extending the tuple would have shipped the same defect one kind wider; and **a predicate tuned to
the majority shape manufactures findings about the minority shape**, which is only visible once the
population is wide enough to contain one. Written into `platform-alignment.md` §8 in this iter's commit.
