# iter-134 — decisions

## `D-M257x-134-1` — a route written as a conjecture gets MEASURED before it gets acted on

`FIX-M257x-iter132-marker-fences-cannot-see-retractions` says the blindness *"plausibly affects every
marker-matching fence in the family. Nobody has checked the others."* Acting on it as written means
four fence fixes. Checking it costs one probe.

**Decision: probe first, against a branch stated before the probe ran** — ≥ 2 blind → pattern, route
upheld; ≤ 1 → refuted. **Measured: 1 of 4, and it is the one already fixed.**

**Why this is a decision and not obvious.** The route was written by this session two iters earlier, in
a progress file that will be read as the milestone's record. **The temptation with one's own route is to
treat it as a finding, because it is written in the same confident register as the findings around it.**
It was a conjecture, and its own text said so twice (*"plausibly"*, *"nobody has checked"*).
`D-M257x-132-4` had just applied this to an inherited route; this applies it to one we wrote.

**The probe is an import, not a grep** (rule 22). Grepping `retract` across the four files would have
returned hits in all four — including in comments *discussing* retraction — and produced the opposite
verdict.

## `D-M257x-134-2` — the reuse gap is recorded against the iter that caused it, and the refactor is NOT taken here

`claim_twin_guard` has carried `RETRACTION_MARKERS`, `_looks_retracted`, a 320-character context window
and a **decaying** waiver file since M257x iter-48, with a docstring naming the exact hazard: *"A site
may legitimately quote a refuted claim in order to retract it."* **iter-132 met that hazard and built a
coarser bucket beside it.**

**Decision: record it against iter-132, and do not refactor under time pressure.**

- **Recorded** because iter-132's own decision entry (`D-M257x-132-2`) presents independent arrival at
  `D-M257x-121-4`'s ruling as a virtue. It is one — and the same paragraph should say that the family
  already held a tested, sharper answer. **A milestone that reports only the flattering half of its own
  pattern-matching is doing the thing it exists to stop.**
- **Not refactored** because sharing the predicate is a *structural* choice — cross-fence import (new
  coupling in a family that has none) or a shared module (a change to how the runner loads members) —
  and the eight vacuous fences on this milestone's record all came from building under pressure.
  Routed as `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` with the two options named,
  which is what makes it actionable rather than a wish.

**Not decided, deliberately:** which option is right. That is the next iter's or the harden pass's work,
and pre-empting it in a decision entry written by the iter that declined to do it would be exactly the
false-confidence this entry criticises.
