---
iter: 117
milestone: M257x
iteration_type: tik
iter_shape: census
status: archived
opened: 2026-08-07
---

# iter-117 — `TOK-08` class 1: the corpus stops mis-citing itself

**Active strategy reference:** [`TOK-08`: census the mechanical classes; stop sampling
them](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07) — milestone-root
`decisions.md`, **supplied by the user on 2026-08-07**, not authored by an agent, and **binding**.

`TOK-07`'s own pre-registered falsification fired at iter-116 (`P = 37` against `P ≥ 15`, sealed at
`85f6f1c` before the first seat), routing the milestone to a user re-scope conversation. The user
supplied it. **This iter seals that record in its FIRST commit, before any sweep work**, exactly as
`TOK-07` sealed its own falsification before its first seat was dealt.

## Step 0 — re-survey before targeting (mandatory)

`TOK-08` names the target rather than inheriting one from a stale tok field, so the re-survey question is
whether the named class is still the largest and still untouched:

| check | result |
|---|---|
| Is intra-corpus mis-citation still the largest measured class? | **Yes — 10 of 37 (27 %)**, band #7 of iter-116's sealed pre-registration, measured 2026-08-07 |
| Has any iter repaired it as a class? | **No.** iter-116 explicitly took **no repair inside the measuring pass**; the class is routed open as `FIX-M257x-iter116-intra-corpus-miscitation-is-the-largest-class` |
| Does a corpus-wide citation-resolution fence already exist? | **No.** `stack-core/` has 25 guards; none resolves a corpus→corpus citation. `corpus_index_guard` checks *index membership*, not *resolution* |
| Corpus ref at open | to be recorded in `progress.md` at the census |

No substitution. The `TOK-08`-named target stands.

## The pre-registered class list — fixed here, may only GROW

`TOK-08`'s falsification branch fires *"after one full mechanical sweep."* That phrase is only
falsifiable if the sweep's extent is fixed **before** the sweep, so it is fixed here. **A class may be
added to this list; none may be removed.**

| # | class | measured size at iter-116 | mechanically decidable because |
|---|---|---|---|
| **1** | **intra-corpus citation resolution** | **10 of 37 (27 %)** — band #7, the largest single class | a corpus citation either resolves to what it names or it does not; no sentence is interpreted |
| **2** | **platform-source citation resolution** | the mechanically-decidable subset of the **14 of 37 (37.8 %)** platform-drift class (band #8) | a platform path/pin either exists in the pinned clone set and carries the named literal, or it does not |

Everything else the readings have named is **semantic** — a claim about what the platform *does* — and
stays with the reading instrument. Declaring it mechanical would be the flattering reading, and §5 refuses
those.

**"One full mechanical sweep" = every class above has a fence that (a) enumerates its population
corpus-wide, (b) stands at zero findings, and (c) ships with a mutation control AND an anti-vacuity
control that can actually fire.** Then, and only then, the reading that grades `TOK-08`.

## This iter: class 1

**Hypothesis.** The 10 mis-citations iter-116 sampled are instances of a population a machine can
enumerate completely. If so, the population is far larger than 10 — the same 1.5×–4× multiplier every
prior enumeration in this milestone has measured — and closing it removes a whole class from the pool
rather than 10 members of it.

**Sub-shapes of class 1, all decidable without reading a sentence:**

| id | shape | check |
|---|---|---|
| **1a** | corpus doc names a repo path | the path exists in the rosetta tree |
| **1b** | markdown link carries an `#anchor` fragment | the fragment resolves to a heading in the target doc |
| **1c** | corpus doc pins another corpus doc at `:NN` or `:NN-MM` | the range exists in the target |
| **1d** | corpus doc pins ITSELF at `:NN` / `:NN-MM` | the range exists in the containing file |
| **1e** | a pin is adjacent to a backticked LITERAL naming the construct | the literal occurs inside the pinned range |

**1e is the shape band #7 actually measured** — *"`ai-readiness.md`'s `✅ CORRECTED M219` blockquote is
`:476-496`"*, *"the **Data** bullet of `cms.md`'s merge banner is at `:44-47`"*. 1a–1d are the cheaper
substrate that must be clean before 1e means anything.

**Expected lift.** No `P` movement is claimed and **no reading is taken this iter** — `TOK-08` puts the
reading after the full sweep, and §9's UNMEASURED-is-not-unmoved refinement applies (this iter will say
so in its close, in those words, as that rule's guard-rail 1 requires). The iter's deliverable is the
census + the repair + the fence, reported as `TOK-08` mandates: **enumerated population, how many were
already false, and the fence's reach with its denominator named.**

**Phase plan.** (0) seal `TOK-08` — first commit, done before anything else. (1) census: enumerate the
class corpus-wide and measure how many are already false, before repairing any. (2) repair to zero,
incrementally committed. (3) ship the fence into `rosetta-extensions/stack-core/` with both controls.
(4) close; state the invocation with every count.

**Escalation conditions.** A finding whose repair would require a platform-repo edit → route forward, do
not edit (v2.8 standing constraint). A finding that inverts a shipped security property (the iter-115
`bash -c` class) → grade by consequence and disclose it in the close regardless of class size. If the
census returns a population near the 10 already sampled, that is evidence the **enumeration** is not
working — `TOK-07` rule 2's guard-rail, restated: a multiplier near 1.0× is not proof the class is rare.

**Acceptable close-no-lift outcomes.** The census returning a small population with a documented
derivation is a **finding**, not a failure — it would be early evidence for `TOK-08`'s own refutation
branch and must be reported as such rather than argued away.
