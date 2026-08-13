**Type:** tik — under `TOK-08`. Closes the **cross-reference remainder** of
`FIX-M257x-iter135-adjudicated-live-defects` (the sample side), now that iters 138–140 worked the census
side. **5th tik of the session — the cap fires at this close.**

# iter-141 — three cross-references, and a pointer that names a retracted title

## Repaired — 3 sites in 3 files, each re-derived at source before the plan was written

| target | what it said | what is true |
|---|---|---|
| **`ai-readiness.md:18-20`** (`adj-B` P-3) | *"the **only** remaining dependency on `workforce` is the member directory … `LoadMembers`/`LoadMembersByUserIDs`, whose implementations **stayed** in `members.go`"* | The `WorkforceDirectory` interface declares **FOUR** methods (`app/internal/aireadiness/manager.go:40-51` @ `ad9f3c498`) — `LoadMembers` · `LoadMembersByUserIDs` · **`BaseMembers`** · **`LevelsCount`** — and **the source's own doc comment (`:36-39`) already says the seam is *"the active-member directory … and the org's skill-scale setting"***. `LevelsCount` is an **org setting** (`readiness.go:770`), and it **did not stay in `members.go`**: it is at `internal/workforce/manager.go:90` |
| **`clerk-integration.md:126`** | → `ant-academy.md:334`, *"the `DEV_LOGIN_ENABLED` public-route pair"* | **rotted +4** — `:334` is the **AI-proxy** row; the `DEV_LOGIN_ENABLED` row is `:338`. Repaired by **naming the row**, not re-pinning it |
| **`backend.md:13`** | → *"see the **M810 prod teardown is UNEVEN** bullet below"* | that bullet was **retitled at iter-127** to *"The M810 prod teardown has now LANDED for both"* and **its body retracts *"UNEVEN"* in its first sentence** |

**`adj-B`'s P-3 is an absolute quantifier over a coupling seam, refuted by the doc comment on the very
interface it cites** — the same shape as this milestone's four security-surface understatements, on a
different axis.

## The finding: the corpus's retraction idiom is generating the defect it documents

`backend.md:13` is a **new shape** — a cross-reference that names its target **by a title the target
itself retracted**. `D-M257x-137-3` covered quoting a retracted *pin*; this is quoting a retracted *name*,
and it is **harder to catch because nothing breaks**: the pointer is a name, so it still resolves, and
**no anchor fence can see it.** The reader simply arrives at a paragraph that opens by contradicting the
sentence that sent them.

And the second half, which this session **measured rather than asserted** (`D-M257x-141-1`): the house
idiom *"it was `:274` at `<sha>`"* keeps the retracted number **live in the text**, where the next
insertion above its target moves it. **In five iters it turned fences RED three times, in three files,
always on a pin whose own sentence existed to retract it** — `roadrunner.md` (iter-137),
`graphql-wundergraph.md`'s `5050` pointer (which rotted **twice on its own**), and `ai-readiness.md`'s
`:326`/`:274` note (**this iter, caused by this iter's own insertion**).

> **Retract by describing the artifact, not by reproducing it.** *"This doc carried two different line
> numbers for it in successive iters"* says everything the quoted number said and **cannot rot.**

Both are now `§5` **rule 63 (c′)** and **(c″)**.

## Test gates

| gate | result |
|---|---|
| **Guard family** (`--repo-root` + `--platform stack-demo/platform @ 0c91421`) | **18 GREEN · 0 RED · 4 not-run** — *after* a genuine RED caused by this iter's own insertion (above), fixed by **removing** the quoted pin |
| **Scoped fence suites** — chosen by what this iter CHANGED (`D-M257x-138-5`) | **102 passed / 0 failed** — *after* 2 failures on the same self-inflicted defect |
| **Whole suite** | **NOT re-run — §5 rule 60 requires saying so.** Zero `rosetta-extensions` files changed this iter or in any of iters 133–141; iter-132's clean run stands on the same rext tree (`223e4a6`). Stated as a gap, not characterised as covered |
| **Suite wall-time** | not quoted — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands |

**Second session running that the anchor suites caught a self-inflicted defect before the commit** — the
mechanism `D-M257x-138-5` installed after iter-137 shipped one that survived a whole iter.

## Close — 2026-08-08

**Outcome:** the cross-reference remainder of the adjudicated work list is closed — three sites, each
re-derived at source. The headline is a **new defect shape**: a pointer that names its target *by a title
the target retracted*, which **no anchor fence can see because the pointer still resolves**. And the
session's own recurrence is now measured: **the corpus's retraction idiom turned fences RED three times in
five iters**, each time on a pin whose sentence existed to retract it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; no reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 134–141 took no reading, so the metric is UNMEASURED not unmoved — §9's iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y — 5 tiks this session** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-5**
**Decisions:** `D-M257x-141-1` (the retraction idiom is a rot generator — measured, 3 REDs in 5 iters) ·
`D-M257x-141-2` (a cross-reference naming a retracted title is invisible to every anchor fence) ·
`D-M257x-141-3` (`adj-B` P-3 upheld and widened at source).
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter135-adjudicated-live-defects` — **cross-reference half CLOSED this iter.** Remaining, all
  non-citation: `org-repos.md:227`,`:370`,`:43` · `ai_architecture.md:111`,`:224` ·
  `next-web-app.md:17`,`:186` · `external_services.md:368`.
- **`FIX-M257x-iter141-retraction-idiom-sweep` (NEW)** — the idiom is corpus-wide and **generates** rot.
  Sweep the *"it was `:NNN` at …"* / *"this cited `:NN` until …"* forms and convert them to descriptions.
  **Measured motivation: 3 fence REDs in 5 iters.** A grep for the idiom is the enumeration; unlike the
  retracted iter-138 census, **this class's subject is the citing sentence itself, which carries its own
  head** — so by `D-M257x-140-2` it *is* censusable.
- `FIX-M257x-iter140-receipts-not-checkable-here` (13 of 22) · `FIX-M257x-iter140-receipt-fence` ·
  `FIX-M257x-iter138-anchor-rot-fence` (re-specified: head resolution first) ·
  `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` (now with two named consumers) ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
**Lessons:**
1. **A title is a citation.** Retitle a section because its claim was retracted, and every pointer using
   the old title now asserts the retracted claim — in the pointer, where no fence looks.
2. **Retract by describing, never by reproducing.** Three fence REDs in five iters, all from an idiom
   whose purpose is correcting citations.
3. **The doc comment on the interface you cite is evidence.** `adj-B`'s P-3 was refuted by the four lines
   directly above the construct the sentence pointed at — nobody had read down.
