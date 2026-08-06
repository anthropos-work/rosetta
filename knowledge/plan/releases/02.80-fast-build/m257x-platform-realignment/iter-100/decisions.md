# iter-100 decisions

## `D-M257x-100-1` — a fence's REACH is part of its verdict, and must be gradeable

`anchor_construct_guard` reported *"every resolvable anchor names a construct"* over **360 of 555**
citations. The sentence was true and the coverage was 65 %, and nothing in the output let a reader convert
one into the other. The pre-registration band that caught it (#9, ≤1) worked precisely because it was set
against the guard's *claimed* subject rather than its actual reach.

**Rule:** a fence that skips part of its subject must name the skipped class, and a green must be read as
*green over its reach*. This guard already counted unresolvable heads; what it lacked was anyone treating
that count as a limit on the verdict. Adopted into the protocol as §5 rule 46.

## `D-M257x-100-2` — resolution rules are narrowed by CONSTRUCT, never by answer key

The widened resolver went 360 → 511 anchors and 0 → 23 findings. Triage showed the guard's own documented
failure mode returning: ports resolving as anchors. Four narrowings were applied, and each is justified by a
construct the corpus demonstrably uses rather than by which findings it removes:

1. a complete backticked span (not a prefix) — kills `:8082/graphql/query`;
2. an intervening **address** breaks inheritance — `block_ref`'s existing *"more than one → ambiguous"* rule
   on a second axis;
3. an intervening **filename** breaks inheritance, backticked or bare — the prose switched files;
4. a **superseded quote** (*"it was `:11`"*, *"the anchor said `:489`"*) is not graded.

Rule 4 is load-bearing beyond its count: without it the fence reddens on documented repairs, punishing the
discipline it exists to enforce. **Each narrowing has a mutant, and each mutant is a named kill** — that is
what separates "narrowed for a reason" from Trap A.

## `D-M257x-100-3` — an anchor is defective only if it is wrong at EVERY ref its block names

A block naming two refs is the corpus recording a *move*, not an ambiguity to be resolved by default. The
prior behaviour fell back to the default ladder and graded historical anchors against origin HEAD — reading
a file the document never named for that anchor, which is the exact defect `read_target`'s docstring was
written against, surviving in the one path it did not cover. Six findings were of this class.

Corollary, adopted: **a finding must carry the ref it was graded at.** The run-level `adjudicated at` line
names every ref the pass touched and cannot attribute one to a finding; this iter spent a full derivation
establishing that `app/main.go:1450` is a `}` at `9d00a313` and a constructor call at `2035f9a`.

## `D-M257x-100-4` — intra-document self-contradiction is NOT fenceable, and clause 5's zero is a READING result

Asked what would bound the open-ended residual classes. Scoping errors and self-inflicted model drift both
have derivable legal sets and a working precedent (`platform_predicate_guard` / `demo_knob_guard` for the
first; `repair_reach_guard`'s `--range` for the second) — each is roughly one sibling guard.

**Intra-document self-contradiction is not, and the answer is a plain no rather than a maybe.** It is
quadratic in a document's claim count and the relation is semantic entailment, not string identity —
`frontend-tier.md` carried both readings of the demo-academy auth model nine lines apart, in different
vocabulary, each internally coherent. Detecting it *is* deciding what a sentence claims, which is the line
this entire fence family declares it does not cross (`anchor_construct_guard`'s docstring, on blocker #17,
re-declined again in this iter for 3 of iter-99's 7).

**Consequence, stated rather than hedged:** the fence family can drain the enumerable classes to zero and
the residual will still be whatever two blind readings find, at the recall those readings have. Clause 5's
zero is reachable only through *reading*, which is what clause 5 already demands — the fences shrink the
pool the readings must search, they cannot replace them.

## Carried, untouched, exactly as standing

`FIX-M257x-iter53-union-set` · `FIX-M257x-iter56-assignment-flake` ·
`CHECK-M257x-iter38-ai-act-classification` · RF-2/3/7–14 — **not touched.**
`DEF-M257x-iter80` — **still not resolved**; the false present-tense claim was *withdrawn*, not made true,
and re-classing that store remains the user's open question.

`origin/pr-14` — not read, not merged, not cherry-picked. Verdict stands at DO NOT MERGE.
