# iter-71 — decisions

## `D-M257x-71-1` — a citation is graded at the ref ITS OWN BLOCK names, not at a process-wide knob

iter-68 gave three guards a ref. It gave each of them **one** — a module-level `CITE_REF` read from
the environment — and that was the right first move against the defect it had just found (three
checkers reading whatever the clone had checked out). But **the corpus does not have one ref.**

Measured on the live corpus at this build: of **125** resolvable citations, **31 sit in a block
naming exactly one ref that resolves in their own clone**, and every one of them was being read at
`origin/main` regardless of what its sentence said. `backend.md:39` pins its mux citations to `app`
`b948604` v1.366.0; iter-69 re-pointed `shared_libraries.md:79` to `9d00a313` v1.367.0. **No single
knob can be right about both**, and iter-68 measured the consequence in the sharpest possible form:
the same corpus is GREEN at origin HEAD and **4-findings RED** at the pinned build ref.

`block_ref()` implements `D-M257x-69-1` — *a ref-pinned citation is a MEASUREMENT and is graded at
the ref it names* — the same reading that turned "64 unrepaired citations" into 5. Three outcomes:

| block names | behaviour |
|---|---|
| exactly one sha that `rev-parse`s **in that citation's clone** | read at it (`block-pinned`) |
| **more than one** | fall back, and **count it as `ambiguous`** |
| none, or none that resolve | the `CITE_REF` ladder, unchanged (`default`) |

**The ambiguous case is a refusal, not an oversight.** A block legitimately names two refs when it
*contrasts* them — `platform-alignment.md` rule 32, as this milestone rewrote it at iter-69, names
the pin **and** the ref that deleted the line. Picking one would be a rule fitted to a sentence
(§4 Trap A). Falling back silently would hide 12 citations inside `default` and let the fence
over-report its own reach, which is the failure this milestone has now found eight times. So it
falls back **and says so**: `ref chosen by default x57, block-pinned x31, no-clone x30, ambiguous x12`.

**A sha must resolve in the citation's OWN clone to count as its pin.** A platform sha quoted beside
an `app` citation has pinned nothing readable; inventing a ref is worse than using the default.

**`CITE_REF=worktree` still overrides every block pin.** It is the escape hatch that asks *what does
the checkout say* — the mode iter-68 used to demonstrate the whole defect — and a per-block pin that
defeated it would remove the only way to ask the question. Verified live: still 1 finding under
`worktree`, still GREEN by default.

**A pin at which the cited FILE does not exist is UNMEASURED**, not clean and not a finding. §5
rule 7: a ref the corpus named and that does not yield the file is a failure to measure, not a
licence to read something else and call it the same thing.

Deliberately **not** reusing `platform_predicate_guard._REF_PINNED`: that regex also matches DATES
(`at 2026-08-04`), which pin a claim but cannot be handed to `git show`. Same predicate, different
output type — sharing it would have meant passing `2026-08-04` to `rev-parse`.

## `D-M257x-71-2` — a mutant SURVIVED, and it was the window bug for the third time

The battery's line-window mutant — narrow `block_ref`'s window from the **block** to `lines[i]` —
**passed the entire suite.** Every fixture I had written put the pin and the citation on the same
line, so the test agreed with the implementation instead of with the corpus.

The corpus does not write that way. `backend.md:39` carries its pin mid-sentence across a wrap;
rule 33 exists *because* pins and the claims they qualify are separated in prose. **Twice already in
this milestone a one-line window has been the bug** — iter-63's retired-token discriminator and
iter-68's negation window, the latter recorded as *"the second window bug of this milestone wearing
a policy's name."* This is the third, and this time it was in the **test**, not the rule.

Two fixtures added, both derived from how the corpus actually reads: a pin **two lines above** its
citation inside one block (must be found), and a pin **one blank line away** in a different block
(must **not** be — rule 33: a pin exempts a claim, never a neighbourhood). Both mutants now die.

**The generalisable rule, and it is the same one harden pass 16 wrote:** a fixture must be derived
from the corpus, not from the pattern. Iter-68's lesson 4 said *"write the lesson down and you will
still not apply it."* Applied here only because the mutation battery refused to let it pass.

## `D-M257x-71-3` — `run()` returns a seven-positional tuple, and it broke four callers

Adding one reach counter grew `anchor_construct_guard.run()`'s return from a 6-tuple to a 7-tuple
and broke **four existing test call sites** — visible only because the battery's own baseline came
back with 4 errors rather than OK, which is exactly why a battery is run against a *stated* baseline
instead of against a remembered one.

Fixed at the call sites (minimal edit). **Recorded rather than refactored**: a positional 7-tuple
that has now grown twice in four iterations is a real fragility, but converting it to a dataclass
touches `main()`, `postcondition_sites()` and every test, which is a third line of work in an iter
that already has two. Routed as `RF-M257x-iter71-run-returns-a-tuple`.
