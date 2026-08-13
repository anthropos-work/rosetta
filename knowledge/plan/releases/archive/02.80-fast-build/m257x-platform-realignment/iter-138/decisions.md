# iter-138 — decisions

## `D-M257x-138-1` — anchor ROT is mechanically decidable even though the anchor's CLAIM is not, and the fence excluded both

`corpus_citation_guard.py`'s docstring excludes **bare `:NN` pins** *"outright"* as *"not mechanically
decidable."* That is **true of the claim** — *does this line say what the sentence says it says?* needs the
sentence read. **It is false of rot**, which is a different predicate and needs no sentence at all:

> If the citing line was authored at commit `C`, and the text that stood at the cited line at `C` now
> stands at a **different** line, the pin **rotted** — and the line it now stands at is the repair.

`adj-E` proved this reasoning sound on five hand-checked anchors. iter-138 ran it corpus-wide.

> **Rule.** An exclusion is only as narrow as the predicate that justified it. *"Not mechanically
> decidable"* was established for **content** and applied to **the whole class** — so the largest
> measured citation defect on this milestone sat inside a fence's declared blind spot **and did not need
> to.** When you exclude a class, name the predicate you tested; the next predicate over may be free.

## `D-M257x-138-2` — the measured population, with its denominator and its undecidable buckets

Pre-registered branch (sealed in `overview.md` before the probe ran): **≥ 5 rotted → route the fence
build; ≤ 4 → refute it.** Measured:

| bucket | n | what it is |
|---|---|---|
| bare `:NN` pins in `corpus/**` | **588** | the raw population |
| `out-of-range-then` | 241 | target line did not exist in **this** file at the citing line's blame commit — **largely cross-file continuation pins** (`` `main.go:507`, `:509` ``), which the probe reads as same-file. **Its declared floor** |
| `target-too-short-to-be-unique` (< 25 chars) | 109 | table separators, fence markers — excluded, not judged |
| `target-text-gone` | 16 | content deleted since; rot vs deletion is not separable here |
| **decidable** | **222** | unique, non-trivial target text present at HEAD |
| ├ **`STABLE`** *(positive control)* | **95** | pin still lands on its content — **non-zero, so the probe can return "fine"** |
| └ **`ROTTED`** | **127** | **57.2 % of the decidable population** |

**Every delta is positive** — the target moved *down*, i.e. lines were inserted **above** it. That is
`adj-E`'s causal story reproduced at scale, and it is why the remedy is *fence the form*, not *repair
harder*: **no author could have prevented any of them.** Deltas run `+1` (22 cases) to `+135`.

> **The number is published with its floor stated, per iter-114's rule.** *"127 rotted"* is a claim about
> the **decidable 222**, not about the 588 — and the 241-strong `out-of-range-then` bucket is the probe's
> honest admission that it cannot tell a same-file pin from a continuation pin. A fence must resolve that
> before it can enumerate the true population.

## `D-M257x-138-3` — the repairs NAME the construct instead of re-pinning it

Every citation repaired this iter replaces a line number with a **construct name plus a durable
substring**, and records the old pin as retracted. Re-pinning would have been faster and would have
restarted the exact clock that produced the defect — `graphql-wundergraph.md`'s `5050` pointer has now
rotted **twice** (`:174-176` → iter-98 → `:193` → iter-138), and one paragraph of that file had **three**
rotted pins in it.

> **Rule.** A repair that restores the failing *form* has fixed an instance and preserved the class.
> `grep -n "refuses the connection"` returns the citing line and its target and nothing else — **that
> substring is the citation; the line number never was.**

Consistent with iter-133 (*the robust re-derivation is the shared substring*) and iter-137's
`D-M257x-137-3` (*never quote a retracted pin*).

## `D-M257x-138-4` — the probe is EVIDENCE; the fence is TOOLING, and they go in different repos

`rot-probe.py` is a one-shot measurement and lives in **this iter's evidence dir**. The durable fence it
justifies belongs in **`rosetta-extensions`** under the standing policy (all stack-operating tooling in
rext, built in the authoring copy, tagged, pushed). **Routed, not built here** — a fence needs a mutation
control and an anti-vacuity control that can actually fire, and **all eight vacuous fences on this
milestone's record were built under time pressure at the end of an iter** (`D-M257x-134`). The probe's
`STABLE=95` is a control for the *probe*; it is not the fence's control.

## Upheld claims counted as results

- **`payments/handler.go:302-316`** — the `analytics-go` half the corpus got right. Opened at `app`
  `ad9f3c498`: `:302` is `m.analyticsManager.Track(analytics.Event{`, `:316` is its closing `})`.
  **Exact.** Only the `main.go` half was wrong.
- **`adj-D` and `adj-E` were correct on every anchor this iter re-derived** — 7 of 7. Unlike iter-136,
  where the seat's number was right and its candidate wrong, these adjudicators named targets that hold
  up when opened. Recorded because the milestone books adjudicator accuracy in both directions.

## `D-M257x-138-5` — iter-137's own repair introduced a citation defect that escaped BOTH of its gates

`test_anchor_offset_guard.py::TestAntiVacuity` went **RED** on this iter's scoped run, on
`corpus/architecture/dependency_map.md:9` — **a line iter-137 wrote one commit earlier**. The repair
cited `` `services/README.md:39` ``, and **six files in this corpus are named `README.md`**, so the
resolver could not bind it. Fixed by fully qualifying the head to `corpus/services/README.md:39`;
`G.citations(LIVE)["_ambiguous"]` now returns `[]`.

**It escaped both gates iter-137 ran, and for two different reasons:**

| gate iter-137 ran | why it missed this |
|---|---|
| the 22-member guard family | `anchor_offset_guard` is **commit-scoped** — it needs `--range`, was not given one, and reported **NOT-RUN**. It sat in the *accepted-gap* bucket iter-137 disclosed |
| 9 scoped fence suites | `test_anchor_offset_guard.py` **was not one of the nine.** The nine were chosen by "which fences' subjects did I edit" — and this fence's subject is *citation form*, which every repair touches |

> **Rule.** **A repair that rewrites citations must run the CITATION fences, not the fences for the
> subject it was writing about.** iter-137 picked its suites by topic (`platform_alignment`,
> `claim_twin`, …) when the thing it had actually changed, at 29 sites, was **anchors**. The
> not-run bucket is disclosed honestly every iter — but *disclosed* is not *covered*, and here the two
> mechanisms failed in the same direction on the same commit.

**Booked as a success of the milestone's discipline, not only as a defect:** the escape survived exactly
**one iter**, was caught by a control (`TestAntiVacuity`) whose entire purpose is this, and cost one line
to fix. That is the loop working.
