# iter-103 adjudication — readings #25 and #26, graded against the SEALED pre-registration

**Read this line first.** `N = 33`. The pre-registered rule says **`≥ 23` → THE BURN-DOWN LEG DOES NOT
REACH THE RESIDUAL.** That is the verdict, it fired on the branch the pre-registration named as the one
that changes the milestone's plan, and it is reported first and unsoftened, exactly as binding condition 8
requires.

The rule was sealed in its own commit (`04cbcfc`) before any seat was dealt, and it is graded here as
written. It was not re-cut, re-centred, or re-read after the number was known.

---

## Shape of the read as it actually ran

| | value |
|---|---|
| corpus under audit | `e6aed2e` (`iter(M257x/102)`), branch `m257x/platform-realignment` |
| scope | `corpus/services/**` + `corpus/architecture/**` — 40 files, 10,646 lines |
| partition | 7 seats, greedy LPT, 1495–1552 lines/seat (**recomputed**; not iter-101's) |
| reading #25 | **7 of 7 seats** — 23 blockers |
| reading #26 | **7 of 7 seats** — 25 blockers |
| total bookings | **48** |
| adjudicators | 4, grouped by seat LETTER so both readings of one file set land with one grader |
| instrument | `briefing-r26-AS-DELIVERED.md` sha256 `4178b00f0d8e…`; its frozen half byte-identical to `instrument/briefing-iter76-AS-RUN.md` sha256 `3858ec53…`, re-checked AFTER copying |

**Both readings are 7-seat readings.** iter-101 was a 13-seat reading (`r24-D` lost to a spend limit) and
every cross-reading quantity there had to be normalised. Nothing here does. `n₂ = 28` is a **7-seat** figure
and is stated that way wherever it appears below.

### The three re-dealt seats

Run 67 died on a session limit with reading #26 at 4 of 7. `r26-A`, `r26-E`, `r26-G` were re-dealt under the
**identical instrument and the identical recomputed partition** the other four seats were graded under — the
partition is part of the instrument, so it was **not** re-balanced for a 3-seat batch. Every seat in both
readings was committed **verbatim, pre-adjudication**: seven through a network drop, four through a session
limit, three through this run.

---

## The number

| quantity | value | construction |
|---|---|---|
| BOOKED | **48** | 23 (#25) + 25 (#26) |
| UPHELD | **47** | |
| REJECTED | **1** | `r25-G B3`, classed `wrong-tree` |
| **upheld rate, RAW** | **97.9 %** | 47/48 |
| **upheld rate, `wrong-tree` SEPARATED** | **100.0 %** | 47/47 |
| n₁ — reading #25, distinct in-scope upheld anchors, **7 seats** | **25** | |
| n₂ — reading #26, distinct in-scope upheld anchors, **7 seats** | **28** | |
| within-reading overlap `m` | **20** | |
| **`N` — union of distinct in-scope upheld anchors** | **33** | `25 + 28 − 20` |
| distinct false PREDICATES behind those 33 anchors | **22** | 28 adjudicator predicates − 6 cross-adjudicator merges |
| Chapman `N̂` | 34.9 | **and it is not usable — see band #3b** |
| per-pass recall vs the union | #25 **75.8 %** · #26 **84.8 %** | |

**`N` is constructed exactly as iter-101's 24 was** — distinct fully-qualified in-scope corpus anchors,
unioned across the two readings, post-adjudication. A looser extractor that also attaches bare `:NN`
continuations to the preceding file gives `n₁ = 33, n₂ = 34, m = 25, N = 42`. **That looser number is not
compared to 24**, because 24 was not produced that way; it is reported only to size the anchor inflation.

### The comparison that matters, and it is not the one the verdict rule asks for

| | iter-101 (`8f04d3a`) | iter-103 (`e6aed2e`) |
|---|---|---|
| distinct false **predicates** | **22** | **22** |
| distinct **anchors** | **24** | **33** |
| anchors per predicate | 1.09 | **1.50** (2.0 on the loose count) |

**By predicate the pool did not move at all — 22 then, 22 now.** By anchor it rose 38 %. The corpus is not
carrying fewer false propositions after a 98-site repair; it is carrying the same number of them in more
places.

---

## Bands — 4 HELD of 10, and the six failures share ONE cause

| # | prediction | band | measured | verdict |
|---|---|---|---|---|
| 1 | per-reading in-scope upheld count (n₁, n₂) | [5, 18] each | **25** and **28** (7 seats each) | **FAILED HIGH — both** |
| 2 | **union `N`** | [8, 22] | **33** | **FAILED HIGH** |
| 3 | overlap with iter-101's published 24, matched on PREDICATE | [0, 4] | **1** | **HELD** |
| 3b | within-reading overlap `m` | [1, 7] | **20** | **FAILED HIGH, by a factor of 3** |
| 4 | adjudicator upheld rate (raw) | [74 %, 88 %] | **97.9 %** raw · **100.0 %** separated | **FAILED HIGH** |
| 5 | the two passes' recalls differ by | ≥ 15 pts | **9.1 pts** (75.8 / 84.8) | **FAILED** |
| 6 | `wrong-tree` rejections | [0, 4] | **1** | **HELD** |
| 7 | wrong-construct **intra-corpus** citations among upheld | ≤ 5 | **3** | **HELD** |
| 8 | platform-drift share of upheld in-scope blockers | ≤ 10 % | **20 of 33 = 61 %** | **FAILED, hard** |
| 9 | per-seat booked spread over 14 seats | ≤ 8 | **5** (max 7, min 2) | **HELD** |
| 10 | repair-induced — anchor inside prose iter-102 wrote | [1, 6] | **7 of 33** | **FAILED HIGH by one** |

### #3 — the band that says the repair DID work, and it held

**1.** Exactly one of iter-101's 22 predicates is re-found: **`prod-terraform-8081`**, at
`corpus/services/skiller.md:19` — an anchor iter-101 never booked, but which **iter-102's own repair map
listed as a twin of it and flagged `SEAT 9 (?)`**. Every other twin of that predicate was repaired; this one
was not, and the parenthesised question mark in the repair ledger is where it survived.

Two anchors of iter-101's 24 recur — `sentinel.md:85` and `cms.md:196` — and **neither is a predicate
match**. Both carry *different* false propositions now than they did then. They are band #10 findings, not
band #3 findings, and the distinction is the whole reason band #3 is matched on predicate rather than on
anchor.

**So repair efficacy is confirmed by a blind re-read: 21 of iter-101's 22 predicates are gone.** That
corroborates `repair_reach_guard`'s ~100 % reach against the upheld set with an independent instrument. The
repair leg reaches what it aims at. **`N` did not fall anyway**, and the rest of this sheet is why.

**What #3 does not establish.** Absence from this reading is a detection measure, not proof of repair. This
reading demonstrably missed things it was looking straight at — see the `:321` cluster below, where **2 of
the 3 in-scope sites were found and `backend.md:54` was missed by both readings**.

### #3b, #4, #5 and #8 failed together, for one reason: the residual changed COMPOSITION

This is the finding, and it is one finding, not four.

| | iter-101 | iter-103 |
|---|---|---|
| `m` / union — share of the union found by BOTH passes | 4/24 = **17 %** | 20/33 = **61 %** |
| upheld rate | 77.8 % raw / 80.0 % separated | **97.9 %** raw / **100.0 %** separated |
| per-pass recall spread | 83.3 vs 33.3 = **50 pts** | 75.8 vs 84.8 = **9.1 pts** |
| platform-drift share | 1–2 of 24 ≈ **4–8 %** | 20 of 33 = **61 %** |

iter-102 repaired the residual's *subtle* half — those were the defects a reading books. What is left is
dominated by **mechanically checkable drift**: a version literal, a `go.mod` pin, a symbol name, a line
offset. A mechanical defect is found by **every** competent pass (so `m` explodes and the recall spread
collapses) and leaves a seat almost no room to be wrong about it (so precision goes to 98 %).

> **Precision, overlap and inter-pass independence are properties of the RESIDUAL'S COMPOSITION, not of the
> instrument.** The instrument was byte-identical across iter-101 and iter-103. Every one of these four
> numbers moved, and three of them moved in the direction that *flatters* the reading.

That retires a conclusion this milestone has been leaning on. **See item 4 below.**

### #8 — graded FAILED, with all three readings of the class disclosed

| definition of "platform-drift" | count | share | verdict |
|---|---|---|---|
| service / subgraph MEMBERSHIP state only (the two platform guards' fenced class) | 3 of 33 | 9.1 % | would HOLD, at the edge |
| the M257x class as historically counted (membership + a library folded into `app`) | 7 of 33 | 21.2 % | FAILS |
| **the claim was TRUE when written and the platform moved past it** | **20 of 33** | **60.6 %** | **FAILS, and this is the graded figure** |

Two more (`ai-readiness.md:429`, `:436`) are plausibly drift and are **not** counted, because whether those
anchors were correct at `b948604f` was not established. The graded figure is a floor.

**The graded definition is the honest one** because it is the one that describes what the reading found. The
class that exploded is **version-pin and line-anchor drift from the five advanced clones** — `app`
`b948604f → ad9f3c49`, `next-web-app → 8297c684`, `sentinel → f2c46190`, `studio-desk → 41ee3575`,
`ant-academy → 22df69dd`. **Neither platform guard fences it.** `platform_alignment_guard` fences
`repos.yml` membership; `platform_predicate_guard` fences compose profile tokens. Nothing fences *"the
corpus quotes a `package.json` literal that has since been bumped."*

**The pre-registration named this risk in its own #8 commentary** — *"five clones advanced since iter-101's
sheet, so this reading is the first to grade the corpus against `app` at `ad9f3c49`"* — and set the band at
≤ 10 % anyway. That was the right call: the band was allowed to fail, and it did, by 6×.

### #10 — repair-induced, and the mechanical count understates it

**7 of 33** upheld anchors sit inside line ranges iter-102 wrote (`git diff --unified=0 925fabf..e6aed2e --
corpus/`): `backend.md:148-153`, `backend.md:294`, `cms.md:55`, `hiring.md:24-25`, `jobsimulation.md:50`,
`jobsimulation.md:146`, `sentinel.md:85` (`cms.md:196` and `ant-academy.md:63` on the loose count → 9).
Band [1, 6] → **FAILED HIGH by one**.

The ~2-per-cycle induction rate has now held for **six consecutive cycles**. But this cycle it changed
*shape*, and the shape is worse than the count:

**(a) The single largest predicate in the whole reading IS iter-102's canonical repair wording.**
iter-102 closed `prod-terraform-8081` (CANON-2) by replacing the unmeasurable assertion with a sentence
that says the literal has *"one occurrence anywhere in the clone set."* **It has six** — five of them in
`stack-demo/rosetta-extensions`, which the same sentence's own 13-repo / 44-`.tf` denominator counts as one
of its repos. The replacement sentence is **self-refuting against its own stated denominator**, and it was
published at **five anchors** (`cms.md:55`, `cms.md:196`, `jobsimulation.md:50`, `backend.md:148-153`,
`backend.md:294`). Three seats in reading #25 and three in reading #26 found it independently.

> A wide repair that ships one canonical sentence to N sites ships its defects to N sites too. **The
> canonical-wording mechanism converts a single authoring error into an N-anchor defect**, and it is the
> reason anchors-per-predicate went 1.09 → 1.50.

**(b) iter-102 rotted an anchor by inserting prose above it — the identical mechanism iter-101 booked
against iter-100.** At `8f04d3a`, `architecture_overview.md:321` **was** the correct local-stack line
(*"Connect-RPC to sentinel (the only cross-process RPC edge out of backend on a core stack)"*). iter-102
inserted a production-topology block above it; the wording moved to **`:331`** and **every citation to
`:321` stayed put**. `:321` now names *"→ backend (the sole subgraph)"* under the **production Cosmo Router**
— the opposite topology from the one being cited.

Measured, corpus-wide: **4 sites cite `:321`** (`sentinel.md:85`, `jobsimulation.md:146`, `backend.md:54`,
and `CLAUDE.md:282`), and **0 cite `:331`**. The reading found 2 of the 3 in-scope sites and **missed
`backend.md:54`**, which sat inside seat E's own file set in both readings.

> iter-101's finding against iter-100 was: *a two-line parenthetical pushed a table down two rows and left
> the numbers unmoved.* iter-103's finding against iter-102 is the same sentence with different nouns. **The
> repair leg reproduces its own documented induction mechanism one cycle later.** This is not a new class; it
> is an unrepaired one.

### #7 — HELD at 3, and the three are worth naming

Intra-corpus wrong-construct citations (a corpus file citing another corpus file's line, landing on the
wrong construct) — the same definition iter-101 graded its `exactly 4` under:

1. `dependency_map.md:19` → attributes a `depends_on` / `DB_CONNECTION` / `REDIS_ADDR` / `go.mod` statement
   to `storage.md:40,47`. `storage.md` makes that statement **at no line, and never has.**
2. `sentinel.md:85` → `architecture_overview.md:321` (the `:321`/`:331` rot above).
3. `jobsimulation.md:146` → `architecture_overview.md:321`, same rot.

**2 of the 3 are one induced defect.** Held, at 3 of a ≤ 5 band, on a tree that grew +214 citation anchors —
so the repair's *new* citations are, on this measure, as good as the old ones. That is a real, if narrow,
credit to the repair leg.

### #6 — the briefing defect, third measurement: 4 → 1 → **1**

One `wrong-tree` rejection (`r25-G B3` — three `shared_libraries.md` claims about the `ai` library graded
against a tree that does not settle them; all three are true at `v1.40.2`). The class did **not** get worse
when the addendum named the settling tree, and it did not get better either.

**So band #6's question is answered: an addendum CAN carry ground truth a frozen instrument gets wrong,
without editing it.** All fourteen seats stated which `rosetta-extensions` tree they read.
`DEF-M257x-iter101-briefing-rext-tree` stays open and stays undelivered-unfixed — the series is now n=3 and
breaking the instrument now would cost more than the defect does.

### #9 — HELD at 5, and it clears the partition

max 7 (`r26-E`), min 2 (five seats). The recomputed partition did not introduce seat-level variance that
would undercut the reading-level numbers above it. iter-101 got 4 on a different partition.

---

## Composition of `N` — 33 anchors, 22 predicates, six clusters carrying 61 %

| cluster | predicates | anchors | class |
|---|---|---|---|
| the `ai` library folded into `app` (`1e457fa70`) and the corpus still calls it an imported private module | 2 | 6 | drift |
| `backend.internal.anthropos:8081` *"one occurrence anywhere in the clone set"* | 1 | 4 (5 loose) | **iter-102-induced** |
| the Next.js pin — `^16.2.7` quoted where `~16.2.12` is declared | 1 | 3 | drift |
| colony / proto pins — the *"two-way split"* and the *"live skew is two"* | 3 | 3 | drift |
| `architecture_overview.md:321` → the production router line | 1 | 2 | **iter-102-induced** |
| the *"platform academy subgraph"*, asserted live while two sites deny it exists | 1 | 2 (6 loose) | drift |
| — the remaining 13 predicates | 13 | 13 | mixed |

The `ai`-fold cluster deserves its own sentence, because it is the M257x class exactly: **a shared library
was merged into the monolith, and the corpus describes it as external in six places across four files** —
including `corpus/architecture/README.md:21`, the index a new reader starts from. `shared_libraries.md:126`
already states the correct thing in the present tense, so the corpus **contradicts itself** on it.

---

## THE BURN-DOWN VERDICT

> ### `N = 33` → **THE BURN-DOWN LEG DOES NOT REACH THE RESIDUAL.**
> The pre-registered rule's `≥ 23` branch. Flat-to-rising after roughly half the estimated pool was paid in
> one pass.

**Say it precisely, because the precise version is more useful than the loud one.** The repair leg reaches
what it AIMS at — band #3 proves that independently, 21 of 22 predicates closed. What it does not reach is
**the residual**, because the residual is fed by two inflows the repair does not touch:

1. **Clone advance.** 61 % of `N` is drift that did not exist as a defect when iter-102 ran. Five clones
   moved between the two sheets. No guard fences version literals or line offsets, so this inflow is
   invisible until a reading finds it.
2. **The repair's own induction.** 7 of 33 anchors sit in prose iter-102 wrote, and the two largest induced
   clusters are a false canonical sentence multiplied across five sites and an anchor rotted by an insertion
   above it.

**Inflow is comparable to outflow. A loop with that property does not converge, and running it faster does
not help** — which is the thing the pre-registration said would be more important than the number, and it
is.

### What this verdict does NOT establish

- **It does not establish the pool size.** `N` is a **floor** in every branch — the union of two passes whose
  measured per-pass recall on this milestone has run 33–85 %. A reading measures **detection**.
- **It does not establish that repair is futile.** It establishes that repair-alone is. The 21-of-22 closure
  rate is the strongest single-cycle repair result the milestone has.
- **It does not license re-cutting clause 5.** Clause 5 is met only by a reading that returns **zero**. Four
  user rulings. `N = 33` leaves it open and this sheet does not argue it.
- **It does not establish that the two induced clusters are the whole induction.** Band #10 counts anchors
  inside iter-102's diff. An induced defect that landed *outside* those ranges is not counted, and the
  `:321` rot is proof the mechanism can act at a distance from the edit.

---

## The overlap band and what a SECOND independence measurement says about `N̂ ≈ 103`

Band #3b was written to put a second point on the independence question, from *within* one reading, where
the subject is provably identical for both passes.

| | `m` | union | share found by both | Chapman `N̂` |
|---|---|---|---|---|
| iter-101, within-reading | 4 | 24 | **17 %** | — |
| iter-99 × iter-101, cross-reading | 6 | 28 × 24 | — | **≈ 102.6** |
| **iter-103, within-reading** | **20** | **33** | **61 %** | **34.9** |

**The two independence measurements disagree by 3.6×, and the honest conclusion is that neither Chapman
estimate is usable.**

- Chapman assumes the two passes are independent. Positive correlation inflates `m`, which **deflates** `N̂`.
- iter-101 measured near-independence and concluded `N̂ ≈ 103` **as a floor**. That was correct reasoning on
  that reading's evidence.
- iter-103 measures strong positive correlation on the **same instrument**, which means independence is not
  a property of the instrument at all. **It is a property of what is left to find.** Subtle residual →
  independent passes → `N̂` large. Mechanical residual → correlated passes → `N̂` small.

> **`N̂ ≈ 103` is neither corroborated nor refuted. It is now UNESTIMABLE by this method**, because the
> estimator's load-bearing assumption has been measured at both extremes on one unchanged instrument. The
> series **16.7 → 29.4 → 45.2 → ~103** remains four successive corrections to an underestimate rather than
> four measurements of a growing pool — that reading is unchanged and is still the right one — but the
> milestone should stop quoting a point estimate from it.

**What survives is the floor**, and only the floor: **the corpus carried at least 24 blocking falsehoods in
clause 5's scope at `8f04d3a`, and carries at least 33 at `e6aed2e`.** Both are two-pass unions. Both are
floors. Neither is a pool size.

**Retire the Chapman estimator for this milestone.** Two measurements of its central assumption, at 17 % and
61 %, is enough. Track `N` and the predicate count directly; they are floors, they are comparable across
readings, and they need no assumption at all.

---

## The gate

**It did not move. It stays at 4 of 5.**

- Clauses 1 and 2 were closed by **Lane B, not by this reading**, at platform `0c91421`. Clause 2 is
  **MET WITH DISCLOSURE** and the disclosure travels with it forever: a freshly built stack failed the first
  full run **29/1 in 2 of 2 attempts**. It is never recorded as a clean pass.
- Clauses 3 and 4 hold.
- **Clause 5 is the only open one, and `N = 33` leaves it open.** It is met only by a reading that returns
  zero. It was not re-cut, narrowed, reinterpreted or argued in this pass.

---

## Routed

- **`FIX-M257x-iter103-read-union`** — the 22 predicates / 33 anchors above, routed **by claim, not by
  file**, to the next repair iter. Two riders, both learned here:
  - **The canonical-wording mechanism needs a fence.** A sentence published to ≥ 3 sites must be verified
    against its own stated denominator *before* it is multiplied. CANON-2 was not, and cost 5 anchors.
  - **A repair that inserts lines above a cited anchor must re-point the citations.** iter-102 did not, and
    `architecture_overview.md:321` now has 4 stale citers, one of them `CLAUDE.md`. The `anchor_construct_guard`
    reach is ~60 %; intra-corpus anchors rotted by a corpus edit are inside the 40 % it does not resolve.
- **`FIX-M257x-iter103-drift-fence-gap`** — net-new. Neither platform guard fences version literals,
  `go.mod` pins, or line offsets into platform files, and that unfenced class is **61 % of `N`**. This is a
  tooling gap, not a corpus defect, and repairing the 20 anchors without it just re-arms the same class at
  the next clone advance.
- **`DEF-M257x-iter101-briefing-rext-tree`** — stays open, stays delivered-unfixed. n=3 at 4 → 1 → 1.

## Provenance

Four adjudicators, each re-deriving from the platform clones rather than from any seat's evidence or any
prior verdict; each barred from `knowledge/plan/**` beyond its brief, its assigned seat reports and its own
output. No stack was brought up, torn down or reconfigured. No tag was cut. Zero platform-repo edits.

### §5 rule 41a — held, and PROVEN rather than asserted

`ground-truth.md` recorded every clone's HEAD, its `origin/main` and its **fetch time** at the open,
specifically so a mid-reading move would be detectable rather than merely suspected. All three were re-read
at this close with `rev-parse` + `stat` and **no fetch**:

| repo | HEAD | `origin/main` | last fetched | vs the open |
|---|---|---|---|---|
| platform | `0c91421d` | `0c91421d` | 08-06 12:15 | identical |
| app | `ad9f3c49` | `ad9f3c49` | 08-06 12:15 | identical |
| next-web-app | `8297c684` | `f97ba659` | 08-06 12:15 | identical |
| sentinel | `f2c46190` | `f2c46190` | 08-06 12:15 | identical |
| studio-desk | `41ee3575` | `41ee3575` | 08-06 12:15 | identical |
| ant-academy | `22df69dd` | `22df69dd` | 08-06 11:18 | identical |
| cms · jobsimulation · messenger · storage · roadrunner · graphql-wundergraph | `ca50c817` · `462343b0` · `fa47850d` · `4ce8ece5` · `87d8d443` · `60c229f3` | — | 08-05 23:24 | identical |
| rosetta-extensions (pinned per-stack) | `09d06070` | `4cb920aa` | 08-06 11:19 | identical |
| rosetta-extensions (authoring) | `944fc4a2` on `main`, clean | — | — | identical |
| `app/studio` · `cms/studio` (nested, own checkouts) | `aeec036a` | — | — | identical |

**Not one ref and not one fetch timestamp moved between the open and the close.** The reading's ground truth
is therefore not merely believed frozen — it is measured frozen, at both ends, and that is the corollary
`D-M257x-103-4` exists to make load-bearing rather than decorative.
