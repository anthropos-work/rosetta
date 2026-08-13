# iter-115 — decisions

## D-M257x-115-1: a pin that has rotted three times is DELETED, not re-derived

**Context.** P08 books `ai-readiness.md:52`'s pin of the `⚠⚠ M51 iter-08/09` block at `:496`. Measured
at this iter: `:496` is the **closing line** of the `> **✅ CORRECTED M219 …**` blockquote (`:476-496`)
— the very blockquote the two sentences above it name as the *wrong* target for the previous
generation. The block opens at `:498`; the quoted parenthetical is at `:500`. Off by two.

**The obvious repair is to write `:498`.** Rejected. The anchor's history is `:458` (iter-46) → `:459`
(iter-100, mechanically re-pointed) → `:496` (iter-102) → and now `:498` would be the fourth, in a
document that grows every time this milestone touches it. **Three generations of one same-file anchor,
each off by a handful of lines, is not three accidents** — it is a measured recurrence rate, and the
passage carrying it is published as a worked example of a repaired anchor.

**Decision.** Delete the numeric pin. Name the construct (the `⚠⚠ M51 iter-08/09` block, identified by
its quoted parenthetical), record all three generations and why the fourth was not written, and keep
the self-heal clause (*"when the two disagree, the name wins"*) — which lowers a reader's cost but was
never capable of making the number true.

**Generalised the same way at two more sites this iter**, both forced by the fence: `service_taxonomy.md`'s
table row numbers (rotted a second time — the Jobsimulation number landed on the table **header**) and
§5 rule 22's own worked example in `platform-alignment.md` (rotted a **fourth** time, moved by this
very repair). In all three the construct name was **already present beside the number**; only the
number was load-bearing enough to break.

## D-M257x-115-2: an unenumerated instance of an enumerated predicate is REPAIRED, and the denominator is NOT renegotiated

**Context.** Three sites publish a predicate in the enumerated set but are absent from it:

- `jobsimulation.md`'s *AI providers* bullet — *"via the shared `ai` library"* — a **third** instance
  of **P02**, which the enumeration books at two.
- `architecture_overview.md:423` — *"public subnets (ALB, Cosmo Router)"* — **P22** verbatim.
  Adjudicator 4 **explicitly named it as a second anchor**; iter-113's ceiling pass then excluded it
  with the reason *"mentions the Cosmo Router's deletion, its routing role or its version; none places
  it in a subnet"*, which is false of this line.
- `cms.md`'s mermaid participant line — **P10**, promoted by the enumerator and correct.

**Two options, and both were considered.** (a) Extend the enumeration to 74 sites and grade against
that. (b) Repair all three, grade against 71, and record the excess.

**Decision: (b).** The denominator is the *instrument's* claim about the corpus, taken before the
repair and checked in. Growing it with sites the repair happened to notice makes the grade
unfalsifiable — the number would then measure how much the repairer looked, which is precisely the
detection-bounded failure `TOK-07` was authored to end. **`TOK-07` rule 3 governs the repair; rule 4
governs the grade; they are allowed to disagree in this direction and not the other** (repairing more
than the denominator is safe; repairing less is the defect).

**And the excess is the finding.** All three are booked against
`FIX-M257x-iter113-adjudication-is-judgement`, which was routed open precisely because **15 of 16
small-class verdicts rest on one agent's readings** with no second witness. That item now carries
measured misses rather than a design argument — including one the ceiling pass excluded *after an
adjudicator had booked it*, which is the strongest available evidence about the class.

## D-M257x-115-3: the corrected form must state its ref IN the sentence that depends on it

**Context.** §5 rule 22's worked example (`platform-alignment.md`) is a post-mortem of one anchor that
rotted three times. **This repair moved it a fourth time** — by repairing the very bullet it points at.
Separately, the fence refused three other commits for anchors this repair's own line shifts broke,
including one I authored *in the same commit that shifted its target*.

**Decision, and it is a standing rule for any repair pass under this protocol, not advice:**

1. **Never author an intra-corpus `file.md:NNN` reference in a commit that moves lines.** Name the
   construct — a section heading, a bullet's bold lead, a table row's name.
2. **Every correction states its ref where the claim sits**, not in a neighbouring sentence. Measured
   twice this iter: `coursebuilder.md:48`'s terraform pins graded at the checkout because the bullet's
   only pin was a parenthetical attached to a *different file*; and `external_services.md:672`'s
   `gen.py` range was read against the host ref when the studio tree is a **nested** checkout at
   `aeec036a`.
3. **A citation of a frozen FIXTURE copy of a corpus file must be construct-named too** —
   `anchor_construct_guard` strips the fixture prefix and resolves against the live corpus, so a
   corpus repair false-REDs a fixture citation that did not move. Routed as
   `FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live`; the guard is the thing to fix, and
   until it is, rule 3 above is the workaround and is recorded as such.
