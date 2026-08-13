# iter-168 — decisions

## `D-M257x-168-1` — the exposure is 86.4 % / 81.8 % / 75.0 %, not "in principle"

iter-167 argued that its sibling captures had the same coupling. iter-168 measured it, against the
live ledger's **264** derived claims:

| capture | rosetta | claims adjudicated AFTER it |
|---|---|---|
| iter-41's 18 | `48ca53c` | **228 of 264 — 86.4 %** |
| iter-47's 12 | `72298dd` | **216 of 264 — 81.8 %** |
| iter-48's 18 | `cabc3b1` | **198 of 264 — 75.0 %** |

The one that collided was the one with the SMALLEST post-capture surface. The other two were green
over a larger one. **`§9`'s rule that a reach metric names its denominator applies to a hazard as
much as to a finding** — "the same coupling" is an argument; three percentages over a stated
denominator is a measurement, and it is what makes "green by luck" a claim rather than a mood.

## `D-M257x-168-2` — one shared helper, not three copies

iter-167 wrote `_ledger_file` / `_iter_number` / the scoped assertion inline. Applying that to two
more modules by copy would have created the exact condition `FIX-M257x-iter134` names and that
iter-165 lost a whole reading to: **three schemas for one concept, none of them the definition.**
`tests/frozen_capture.py` is the definition; all three answer keys now call it, including iter-167's,
whose local copies were deleted in this iter.

It also earned its keep immediately. Two coordinate spellings reach `ledger_file` —
`Hit.claim_source` ends `…C.md:42` (a LINE) and `Site.claim_id` ends `…C.md#c-1@21` (an ANCHOR) —
and the wrong one yields a key matching no manifest entry, i.e. a scope that silently excludes
**everything**. In three inline copies that would have been three chances to get it wrong.

## `D-M257x-168-3` — the fourth census member was NOT repaired, and the reason is the rule

`test_repair_postcondition.py::_sites` calls `rp.collect(["claim_twin_guard"], repo, MILESTONE)` on
the same frozen `claim_twin` fixture and asserts `before == []`. Structurally identical; **semantically
not.** There the site set feeds a RATCHET, and filtering post-capture sites out of a ratchet's input
changes what the ratchet counts — it could mask a real induced regression, which is the one thing that
instrument exists to catch.

Measured rather than assumed: the green twin currently yields **0 sites**, so it is not RED today, and
its exposure is iter-41's 86.4 %. **Routed, not repaired** —
`SURVEY-M257x-iter168-ratchet-input-vs-assertion-scope`. iter-158's rule, applied to myself: the
pattern that fits three members is a hypothesis about the fourth, not a plan for it.

## `D-M257x-168-4` — the battery stage-list class is FIVE occurrences deep with an open route naming the fix

Scoping the suites broke `test_m257x_claim_twin_mutation_battery`: its `_COPY_FILES` hand-list did not
carry `tests/frozen_capture.py`, so the staged tree died on `ModuleNotFoundError` and the battery
reported its **BASELINE RED**. Censusing the batteries found **six** with a hand-listed stage set, and
their own comments record the history:

| when | dependency added | what the battery reported |
|---|---|---|
| harden pass 1 | `platform_topology.py` | baseline RED, **no attributable test** |
| iter-111 | `fence_provenance.py` | **RED BASELINE** |
| iter-121 | `corpus_citation_guard.py` | **RED BASELINE**, unseen for four iters |
| iter-166 | `waiver_ledger.py` | **RED BASELINE**, 5 mutant verdicts uninterpretable |
| iter-168 | `tests/frozen_capture.py` | **RED BASELINE** |

**iter-111 routed it as `FIX-M257x-iter111-staged-battery-dependency-is-underived` and stated the fix
verbatim** — *"a battery that stages a SUBSET carries a dependency contract, and nothing derives it."*
The route stayed open; the next three occurrences were each closed by appending a filename. This is
the milestone's founding sentence turned on the milestone itself: *a recurring class with no written
procedure is a class that will recur.*

`tests/battery_stage.py` is the procedure. Two of six migrated here (the one that broke and iter-166's,
whose inline derivation was promoted). **The remaining four are routed, not silently left** — the
scope-creep tripwire's third line, and each one is a live battery whose stage list has to be verified
rather than assumed.

## `D-M257x-168-5` — iter-166 shipped with two batteries RED, and its own honesty note is what found it

iter-166 ran a scoped suite set and named what it had NOT run, per `§5` rule 60. iter-168 ran two of
those un-run members and **both were RED, broken by iter-166's own change**: `claim_twin_guard` and
`repair_reach_guard` grew an `import waiver_ledger`, and the
`repair_postcondition` / `mechanical_fences` batteries hand-list their staged files, so their staged
trees died on `ModuleNotFoundError` and reported **BASELINE RED**. `repair_reach`'s battery was in the
same state for the same reason.

**This is `§5` rule 60 paying out in the direction it is usually quoted against.** The rule is normally
read as a caveat — *"a scoped green is evidence about its scope alone"* — and it is usually a
formality. Here the un-run remainder contained two real regressions from the very iter that wrote the
caveat, and the caveat is the only reason they were looked for two iters later rather than found by
someone else.

The correction to iter-168's own routing follows: **5 of 6 batteries are migrated**, not 2. Three of
them HAD to be — they were broken, and a broken battery is not a route, it is a repair. Only
`test_m255_mutation_battery` remains, and for a stated reason: it stages files from `demo-stack/`,
`stack-verify/` and `stack-injection/`, i.e. **outside the guard directory the helper's derivation
walks**. Migrating it needs the helper widened, which is a change with its own failure modes.
