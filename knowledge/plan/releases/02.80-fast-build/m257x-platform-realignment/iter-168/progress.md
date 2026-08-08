**Type:** tik — under [`TOK-08`](../decisions.md), censusing the class iter-167 repaired one member of.

# iter-168 — the sibling captures, and a class five occurrences deep

## Phase A — census the frozen fixtures and their consumers (denominator stated)

**5 fixture dirs**, **6 consuming assertions.** The predicate that matters is not *"is there a green
fixture"* but *"is a frozen fixture graded against a derivation whose input is LIVE"*:

| consumer | fixture | derivation input | exposed? |
|---|---|---|---|
| `test_claim_twin_guard.py:102` | `claim_twin` (iter-41) | **live** `MILESTONE` | **YES** |
| `test_claim_twin_guard_iter47_answer_key.py:90` | `claim_twin_iter47` | **live** `MILESTONE` | **YES** |
| `test_claim_twin_guard_iter48_answer_key.py:143` | `claim_twin_iter48` | **live** `MILESTONE` | scoped at iter-167 |
| `test_repair_postcondition.py:161,313` | `claim_twin` (iter-41) | **live** `MILESTONE` | **YES — but see Phase E** |
| `test_iter45_mechanical_fences.py` | `mechanical` | the fixture tree only | no |
| `test_repair_leak_guard*.py` | `repair_leak` | a diff + the fixture tree | no |

## Phase B — measure the exposure, do not argue it

Against the live ledger's **264** derived claims (`§9` — a hazard names its denominator too):

| capture | rosetta | claims adjudicated AFTER it |
|---|---|---|
| iter-41's 18 | `48ca53c` | **228 of 264 — 86.4 %** |
| iter-47's 12 | `72298dd` | **216 of 264 — 81.8 %** |
| iter-48's 18 | `cabc3b1` | **198 of 264 — 75.0 %** |

**The one that collided had the SMALLEST post-capture surface.** The other two are green over a
larger one, which settles "green by luck" as a measurement rather than a mood.

## Phase C + D — one shared helper, and every member proves itself

`tests/frozen_capture.py` (new) holds the definition: `assert_capture_silent` (scoped to the
manifest's ledgers, **residual asserted** — every excluded hit must come from a later iter) and
`assert_capture_fires` (the anti-vacuity control — the SAME scoped predicate on the RED fixture must
find every captured site). All three answer keys call it, **including iter-167's, whose inline copies
were deleted here**: three copies of one concept is the condition `FIX-M257x-iter134` names and that
iter-165 lost a whole reading to.

It earned that immediately. **Two coordinate spellings reach `ledger_file`** — `Hit.claim_source`
ends `…C.md:42` (a LINE), `Site.claim_id` ends `…C.md#c-1@21` (an ANCHOR) — and the wrong one yields
a key matching no manifest entry, i.e. a scope that silently excludes **everything**. One definition,
one place to get it right.

## Phase E — the member that does NOT get the pattern

`test_repair_postcondition.py` is structurally identical and **semantically different**: its site set
feeds a **ratchet**, and filtering post-capture sites out of a ratchet's input changes what the
ratchet counts — it could mask a real induced regression, which is the single thing that instrument
exists to catch. Measured rather than assumed (**0 sites** on the green twin today; exposure is
iter-41's 86.4 %) and **routed, not repaired**. iter-158's rule applied to myself: a pattern that fits
three members is a hypothesis about the fourth.

## The class the repair uncovered — five occurrences, one open route

Scoping the suites turned `test_m257x_claim_twin_mutation_battery` RED: its hand-listed `_COPY_FILES`
did not carry `tests/frozen_capture.py`, the staged tree died on `ModuleNotFoundError`, and the
battery reported its **BASELINE RED**. A census found **six** batteries with a hand-listed stage set,
and their own comments record the history — `platform_topology.py` (harden pass 1),
`fence_provenance.py` (iter-111), `corpus_citation_guard.py` (iter-121), `waiver_ledger.py`
(iter-166), `tests/frozen_capture.py` (iter-168).

**iter-111 routed it as `FIX-M257x-iter111-staged-battery-dependency-is-underived` and wrote the fix
down verbatim** — *"a battery that stages a SUBSET carries a dependency contract, and nothing derives
it."* The route stayed open and the next three occurrences were each closed by appending a filename.
That is this milestone's founding sentence — *a recurring class with no written procedure is a class
that will recur* — turned on the milestone itself.

`tests/battery_stage.py` is the procedure, with the data-dependency limit disclosed in the module
rather than left for the sixth occurrence to find.

### And running the "routed" members first turned three of them into repairs

The plan was to migrate 2 and route 4. Running the other batteries before writing that down found
**three of the four already RED — and two of them broken by iter-166**, this run's own previous iter:
`claim_twin_guard` and `repair_reach_guard` grew an `import waiver_ledger` there, and the
hand-listing batteries could not see it. A broken battery is not a route, it is a repair, so
**5 of 6 are migrated**. Only `test_m255_mutation_battery` remains, for a stated reason: it stages
files from `demo-stack/`, `stack-verify/` and `stack-injection/` — outside the guard directory the
derivation walks — so migrating it needs the helper widened.

**iter-166 named its un-run remainder under `§5` rule 60 and that note is the only reason these were
found here rather than by someone else.** The rule usually reads as a formality; this time the
remainder held two real regressions from the iter that wrote the caveat.

## Gates

**Run, green:** `test_claim_twin_guard` · `test_claim_twin_guard_iter47_answer_key` ·
`test_claim_twin_guard_iter48_answer_key` (each now with a net-new `test_02b` control) ·
`test_repair_postcondition` · `test_frozen_expectation_census_m257x` · `test_waiver_ledger_m257x` ·
`test_m257x_claim_twin_mutation_battery` (4) · `test_repair_leak_guard_mutation_battery` ·
`test_m257x_repair_postcondition_mutation_battery` · `test_m257x_mechanical_fences_mutation_battery` ·
`test_m257x_repair_reach_mutation_battery` — the last three as one batch, **16 tests / 603 s, OK**,
after the stage-list derivation; the same batch was **7 failures** before it.

**NOT re-run, named in full (`§5` rule 60):** the rest of `stack-core`, and every other rext section
(`stack-seeding`, `stack-snapshot`, `stack-verify`, `playthroughs`). This iter touched six files, all
under `stack-core/tests/` except none — no guard source changed at all.

## Close — 2026-08-08

**Outcome:** iter-167's repair is confirmed as a **class, measured**: 3 of 6 frozen-fixture consumers
were exposed, with **86.4 % / 81.8 % / 75.0 %** of the live ledger post-dating each capture — and the
one that had already collided held the *smallest* surface. All three answer keys now grade through
**one** shared helper (`tests/frozen_capture.py`), each with its own anti-vacuity control; iter-167's
inline copies were deleted into it. The fourth member was measured and **routed rather than repaired**
because its site set feeds a ratchet. Landing that turned up a second, older class: **six mutation
batteries hand-list their staged dependencies, and that has now failed five times with an open route
from iter-111 naming the exact fix** — 2 of 6 migrated to `tests/battery_stage.py`, 4 routed.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (three consecutive `closed-fixed`/
`-partial` iters; no no-prog streak, and no `N` reading was taken so the metric is UNMEASURED not
unmoved — `§9`) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks this run, not 5) — (6)
protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean:
the iter closed, both repos committed, rext pushed. Two ~10-minute mutation-battery waits inside a
60-minute run are the bulk of the spend, and starting a fourth iter would risk a mid-iter stop)
**Decisions:** `D-M257x-168-1` … `D-M257x-168-4` (see [`decisions.md`](decisions.md))

**Why `closed-fixed-partial` and not `closed-fixed`:** the planned scope — the frozen-capture census
— landed complete. The battery stage-list class was **surfaced by** that work, is five occurrences
deep, and 5 of its 6 members are closed — **three of them because they were already RED, two of those
broken by iter-166**. Grading this `closed-fixed` would let the remaining member and the
ratchet-scope question ride out of the milestone inside a green status line, which is the
mis-classification the protocol's three statuses exist to prevent.

**Side-deliverables:** none. `tests/battery_stage.py` is not a side discovery — the claim_twin
battery went RED on this iter's own change and had to be repaired for the suite to close.
**Routes carried forward:**
- `FIX-M257x-iter168-m255-battery-stages-across-sections` — **NEW, and the last member of
  `FIX-M257x-iter111-staged-battery-dependency-is-underived`**, open since iter-111 and re-encountered
  four times since. `test_m255_mutation_battery` stages from `demo-stack/`, `stack-verify/` and
  `stack-injection/`, outside the guard directory `battery_stage.local_deps` walks. Widening the
  helper is a change with its own failure modes and is not a blind swap.
- `SURVEY-M257x-iter168-ratchet-input-vs-assertion-scope` — **NEW.** `repair_postcondition`'s
  green-twin precondition carries iter-41's 86.4 % exposure, and the scoping pattern must NOT be
  applied to a ratchet's input without deciding what a post-capture site means to a ratchet.
- `FIX-M257x-iter166-stage-derivation-covers-code-not-data` — **sharpened.** Now disclosed in
  `battery_stage.py` itself rather than only in an iter record.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` — **advanced, third consecutive
  iter.** Two shared modules now exist (`waiver_ledger`, `frozen_capture`) where there were none.
- `SURVEY-M257x-iter167-the-other-two-answer-keys-have-the-same-coupling` — **CLOSED by this iter.**
- The standing queue, unchanged.
**Lessons:** **repairing one member of a class and routing "the others probably too" is half a
finding.** iter-167 was right about its siblings and could not say how wrong they were; three
percentages over a stated denominator took ten minutes and inverted the intuition — the collision
happened at the *smallest* exposure, so "the one that broke is the worst one" would have mis-ranked
the queue.

And the harder one: **an open route that names its own fix is worse than no route, if the class keeps
being closed one filename at a time.** iter-111 wrote *"nothing derives it"* and the next three
occurrences each appended a name to the very list the route was about. A route is not a record of
intent; it is a debt, and this one has been paid four times in interest.
