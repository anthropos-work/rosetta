# iter-257 — decisions

## `D-M257x-257-1` — a fence's SUBJECT and its mechanism's subject were two different files, and only one was watched

`clone_pin_guard` (FENCE-M257x-iter222) fences `rosetta-extensions/demo-stack/clones.pin.json`. The
mechanism it exists to protect — `DEMO_ADVANCE_CLONES=pinned` — reads
`stack-*/clones.pin.json`, a **copy** `ensure-clones.sh` seeds **copy-if-absent** and never reconciles.
Measured on this box: the copy named **11** repos to the canonical's **6**, the five phantoms iter-222
removed among them (`cms`, `jobsimulation`, `storage`, `messenger`, `roadrunner`), **all five with
directories present on disk carrying git checkouts** — so the entries were not inert.

**Decision: add arm D rather than change the copy semantics.** The copy-if-absent rule is deliberate —
`ensure-clones.sh:373` names the workspace file an *optional operator pin declaration*, and clobbering a
deliberate declaration would be worse than the staleness. So the asymmetry is encoded instead:

- a **phantom key** in a copy is a **FINDING** (it names a repo the platform removed, and no operator
  declaration can make that legitimate);
- a **value difference** is a **DISCLOSURE**, printed and never failed (an operator pinning a stack to
  an older ref is a state the tooling itself calls `pinned`);
- **no `stack-*/` at all is NOT-RUN with its reason printed**, never a silent pass.

Watched going RED and back: **5 findings + 3 disclosed drifts → 0 findings** after this box's copy was
reconciled. Six new tests including a mutation control that asserts both directions (the mutant GREEN
*and* the un-mutated tree still RED), so the control cannot pass by describing a broken guard.

## `D-M257x-257-2` — iter-256's claim about what the pin controls is RETRACTED, in place, at all three sites

iter-256 wrote that `demo-stack/clones.pin.json` *"changes what a **cold bring-up on any box** builds"*
and that a checkout-only advance would be *"undone by the next `/demo-up`"*. Measured here, both halves
are false:

- `DEMO_ADVANCE_CLONES` defaults to **`0`** (`ensure-clones.sh:220`) and **no file in rext outside
  `ensure-clones.sh` sets it** — a default bring-up applies **no pin at all**.
- A fresh box gets each repo from `git clone` + `make init`, i.e. the **default-branch tip**; an
  existing workspace builds **whatever its clones are checked out at**.
- When `pinned` *is* requested it reads the **workspace copy**, not the canonical file.

**Decision: retract the sentence, keep the action.** Advancing the checkouts was right — that is what a
bring-up consumes. Advancing the canonical pin was right — that is what `pinned` should mean. What was
wrong is the *reason given*, and a right action with a wrong reason is the thing that gets repeated for
the wrong cases. The retraction is landed at all three publishing sites (`iter-256/decisions.md`,
`iter-256/progress.md`, `corpus/ops/platform-alignment.md` §5 rule 79) rather than only where it was
first written — *a retraction that reaches one site has not landed.*

## `D-M257x-257-3` — my own edit moved cited lines, and three fences caught it before commit

Inserting the disclosure branch into `ensure-clones.sh` shifted its line numbers, and
`corpus/ops/demo/demo-up-defaults.md` cites that file twice: `DEMO_ADVANCE_CLONES` at `:212` and
`DEMO_FRESHNESS_STRICT` at `:467`. Three guards went RED from the one cause —
`anchor_construct_guard` (`:467` now lands on a closing `}`), `demo_knob_guard` (STALE ANCHOR, parser
reads `:220`), and `repair_postcondition`. Re-pointed to `:220` and `:475`; family back to
**29 GREEN / 0 RED / 5 not-run**.

**Recorded because it is §7 rule 4 applied to OUR OWN repo.** The rule was written about advancing a
platform clone; the identical hazard exists whenever a rext edit moves a line the corpus cites, and here
the fences supplied the citer list for free. *The iter that moves a cited line re-points it, in that
iter.*

## `D-M257x-257-4` — the noun census read a line-number citation followed by a verb as a measurement

`clone_pin_guard`'s new comment wrote *"`<citation>` calls the workspace file an optional operator pin
declaration"*. The measurement-noun census sees digits followed by a plural-shaped word and books
`calls` as an uncovered measurement noun — a **false positive of the census, caused by ordinary prose**.

**Decision: reword the comment, do not widen the census.** The census's job is to catch numbers nobody
derives; teaching it to except every verb that can follow a citation would cost more reach than it
returns. What the comment now does instead is *describe* the shape without spelling a live example —
because the first attempt to explain the defect **reproduced it inside the explanation**, which is how
the second occurrence was found.

## `D-M257x-257-5` — the new test class was appended AFTER the `__main__` guard, and only a fence saw it

`test_clone_pin_guard.py` ends with `if __name__ == "__main__": unittest.main()`. Arm D's six tests were
appended below it, so `pytest` collected them (it imports the module) while **`python3
test_clone_pin_guard.py` did not — and printed `OK`**. Caught by `test_test_collection_fence.py` in the
whole-section run, not by any scoped run in this iter.

**Decision: move the guard to EOF and record the shape, not just the fix.** This is iter-254's *"a test
that stops running looks exactly like a test that passed"* arriving through a third mechanism — after
`skipUnless` preconditions (iter-254) and a memoised precondition reading `False` (iter-254's `PR-5`),
now file ordering. All three make a real assertion invisible while the runner reports success, and all
three were invisible to the diff.

**And it is the second time in two iters that the whole-section run paid for itself** — iter-256 found
`D-M257x-256-6` the same way. Two findings in two runs, against a routed item
(`ROUTE-M257x-253-the-iter-loop-runs-no-ratchet`) that was a hypothesis three iters ago.
