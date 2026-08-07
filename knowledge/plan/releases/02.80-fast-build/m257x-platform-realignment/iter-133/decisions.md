# iter-133 — decisions

## `D-M257x-133-1` — the route named three anchors; the width search said ten, and the width search wins

`FIX-M257x-iter131-my-three` is a list of three `file:line` anchors, drawn from a reading that samples.
Rule 57's width measurement found **P5 and P6 are one predicate with seven more anchors** — the same
false module set is published in `corpus/README.md`, `corpus/architecture/README.md`,
`service_taxonomy.md` (twice), `dependency_map.md`, `platform_repo.md` and `askengine.md`, none of them
in the route.

**Decision: repair the predicate, and record the substitution in the iter's `overview.md` rather than
silently widening.** This is `TOK-07`'s unit-of-repair rule (*the predicate, not the claim*) surviving
`TOK-07`'s own refutation — the strategy's **premise** about denominators failed at iter-116; its
**mechanism** did not, and `TOK-08`'s own entry says so.

**Not decided:** that every route should be widened. A route is evidence about where a reading looked.
**Grade its scope by measurement, not by deference to the agent who wrote it** — including when that
agent was us.

## `D-M257x-133-2` — the corpus was wrong in BOTH directions, and the cause is a cardinality coincidence

The eight sites do not share an error; they share a **conflation**:

| | members |
|---|---|
| the five **historical shared libraries** (`shared_libraries.md`'s subjects) | `colony`, `authn`, `proto`, `ai`, `taxonomy` |
| the five **private modules a stack imports** (`app/go.mod:14-18` @ `ad9f3c498`) | `analytics-go`, `colony`, `proto`, `storage`, `taxonomy` |

**They overlap in three and share only a cardinality.** So a site could be too *generous* (still listing
`ai`, folded in-tree at `1e457fa70`) or too *stingy* (*"three — colony, proto, taxonomy"*, dropping
`analytics-go` and `storage`, both **direct** requires) — **and several were both**, in one sentence.

**Decision: every repaired site states WHICH five it means, in the sentence itself.** Not a footnote and
not a cross-reference. `CLAUDE.md` already carries the warning — *"do not read this list as `app`'s
dependency set; it is not one"* — and **nine sentences in eight other files did exactly that anyway.**
A warning in the file every agent loads did not stop the error in the files they then read; only naming
the distinction at each site does.

**Why this is a decision:** the cheaper repair was to fix the counts. That leaves the conflation intact
and the sites drift back the moment the module graph moves — which it does: `CLAUDE.md` records the
block shrinking from **seven** to five inside this milestone.

## `D-M257x-133-3` — P7's fence was green over a document that did not define what it enforced

The guard's `ALLOWED_STATES` gained `library-unimported` at iter-130 together with assertion G, and
assertion C's description was updated to say **nine**. **§1's own state table was never given the
row** — so for three iterations the checker enforced nine states, the document defined eight, and
`platform_alignment_guard` was **GREEN throughout**, because it validates rows against its own constant
and never against §1's prose.

**Decision: add the row, and say in it what happened.** The alternative — adding the row silently — would
leave no record that a fence can be green over a definition it does not read.

**And the mirror-image is now recorded on both sides.** iter-131's lesson 1 was *a fence that prints the
right answer does not correct the prose beside it* (assertion G printing the true module set while
`architecture_overview.md:83` contradicted it — repaired in this same iter). P7 is the same defect
running the other way: **the prose that a fence is supposed to implement can fall behind the fence.**
Neither direction is caught by running the fence.

Routed as `FIX-M257x-iter133-two-fives-need-a-fence`: nothing mechanical stops either drift recurring.
**Deliberately not built in this iter** — a fence over prose module-set enumerations is a real design
question (which sentences are in scope? how is "the imported set" parsed from English?), and this
milestone has eight vacuous fences on record from building one in a hurry.
