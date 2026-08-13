**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), working **class 1,
intra-corpus citation resolution**, `TOK-08`'s own named largest class. Declared 2-step shape (repair,
then probe) — see `overview.md`.

# iter-138 — the anchors rotted, and a machine can see it after all

## Priority 1 — the adjudicated citation set, each target re-derived by opening it

`D-M257x-136-1` applied throughout: the adjudicator's line is a **candidate**, never the answer.
**7 of 7 held up** — recorded, because iter-136's seat did not.

| citing site | said | is | verdict |
|---|---|---|---|
| `shared_libraries.md:77` + **`CLAUDE.md:288`** | `analytics-go` wired at `app/main.go:507-508` | **`:494-495`** — `trackingManager := tracking.New(os.Getenv("BREVO_KEY"))` and the `payments.New(…, trackingManager)` that consumes it: **the file's only two `trackingManager` lines**. `:507-508` is the storage-in-app comment block, and names a *different* construct at all four refs the corpus reads | repaired ×2 |
| `security_compliance.md:156` | the *"only"* precedent lives at `clerk-integration.md:40` | `:40` is the **org-invitations** bullet — no quantifier, no security surface. The subject is the **Sign-in tokens** bullet, whose own text reads *"this bullet used to say "only", and it was false"* | repaired |
| `academy-backend.md:15` → `:85-89` | corroborates the subgraph claim | `:85-89` is **Certificate minting**; the claim is the *Primary — GraphQL, on the `app` subgraph* bullet | repaired |
| `academy-backend.md:136` → `ant-academy.md:82-88` | *"states the opposite in bold"* | `:82-88` is the **store.js / beacon write-path** blockquote; the target is the *no-FS-as-published-fallback* paragraph | repaired |
| `graphql-wundergraph.md:138` → `:116-117` | the struck-through *"`make up` rebuilds `graphql`"* bullet | `:116-117` is the **no-hot-reload** bullet | repaired |
| `graphql-wundergraph.md:138` → `:84` | the *Ports* bullet | `:84` is the **Federation** bullet (`federation_version: =2.3.2`) | repaired |
| `graphql-wundergraph.md:88` → `:193` | *"already said `localhost:5050` refuses the connection"* | `:193` is the **profiles-token warning's opening line** | repaired |

**Every repair names the construct and drops the number** (`D-M257x-138-3`). Re-pinning would have been
faster and would have restarted the clock: **that `5050` pointer has now rotted TWICE** (`:174-176` →
iter-98 → `:193` → iter-138), and **one paragraph of `graphql-wundergraph.md` held three rotted pins.**

**The `handler.go:302-316` half was verified and is exact** — an upheld claim counted as a result.

## Priority 2 — is this class censusable, or only sample-able?

`adj-E` established the fact that reframes the whole class: **all five of its anchors were CORRECT when
written**, invalidated by later unrelated edits inserting lines **above** the target (+2/+3/+8/+14).
So *"repair harder"* has no target — there was no careless author. And
`corpus_citation_guard.py`'s docstring **declares this exact blind spot**: bare `:NN` pins are *"not
mechanically decidable"* and **excluded outright**.

**The exclusion is broader than its evidence.** *Not decidable* is true of the **claim** (does the line
say what the sentence says?) and **false of rot**, which git answers alone — `D-M257x-138-1`.

**Branch pre-registered in `overview.md` before the probe ran:** ≥ 5 rotted → the class is real, route
the fence; ≤ 4 → refute it and do not build for a population of five.

### ⚠️ RETRACTED AT ITER-139 — the numbers below are WRONG. Read this first.

**A stratified 12-case audit at iter-139 classified 12 of 12 as FALSE POSITIVES — precision 0.0 %,
Wilson95 [0.0, 24.3].** The figure *"127 rotted / 57.2 % of 222 decidable"* is **withdrawn**, not
re-qualified. Cause: in this corpus a bare `` `:NN` `` is **overwhelmingly a cross-file continuation pin**
(`` `app/main.go:15`, `:62`, `:63` ``) or a **quoted/historical/negated** pin, not a same-file
self-citation. This iter *disclosed* that failure mode as the cause of its 241-case `out-of-range-then`
bucket — and then reported a number over the 222 it had **not** excluded it from. **Naming a floor is not
bounding it** (`D-M257x-139-2`).

**What stands:** all 9 citation repairs (each re-derived by opening it, none from the probe),
`D-M257x-138-3`, `D-M257x-138-5`. **`FIX-M257x-iter138-anchor-rot-fence` is re-specified** — head
resolution first, no baseline until then. **`FIX-M257x-iter138-127-rotted-pins` is withdrawn.**
Full record: `iter-139/decisions.md`.

### Measured (`rot-probe.py`, in this iter dir) — RETRACTED, retained for the record

| bucket | n |
|---|---|
| bare `:NN` pins in `corpus/**` | **588** |
| `out-of-range-then` — **largely cross-file continuation pins**, the probe's declared floor | 241 |
| `target-too-short-to-be-unique` (< 25 chars: separators, fence markers) | 109 |
| `target-text-gone` (deleted since; rot vs deletion not separable) | 16 |
| **decidable** | **222** |
| ├ **`STABLE`** — *positive control; non-zero, so the probe can return "fine"* | **95** |
| └ **`ROTTED`** | **127 = 57.2 % of the decidable population** |

**Every delta is positive** — `+1` (22 cases) through `+135`. Lines inserted **above** targets, exactly
`adj-E`'s mechanism, reproduced at corpus scale. **The branch fires at 127 against a threshold of 5.**

**The number is published with its floor**, per iter-114: *127* is a claim about the **decidable 222**,
never about the 588, and the 241-strong bucket is the probe admitting it cannot separate a same-file pin
from a continuation pin. A fence must resolve that before it can claim the true population.

**The fence itself is ROUTED, not built here** (`D-M257x-138-4`): it belongs in `rosetta-extensions`
under standing policy, and needs a mutation control plus an anti-vacuity control that can fire. **All
eight vacuous fences on this milestone's record were built under time pressure at the end of an iter.**
The probe's `STABLE=95` controls the *probe*, not the fence.

## Test gates

| gate | result |
|---|---|
| **Guard family** (`--repo-root` + `--platform stack-demo/platform @ 0c91421`) | **18 GREEN · 0 RED · 4 not-run** (commit/input-scoped: `anchor_offset`, `repair_leak`, `repair_reach`, `value_change`). **Not a whole-family green, and the runner says so.** |
| **Scoped fence suites** — chosen by **what this iter CHANGED (anchors)**, not by topic | **102 passed / 0 failed** (`corpus_citation_guard`, `anchor_construct_denominator`, **`anchor_offset_guard`**, `repair_postcondition`, `corpus_index_guard`) — **after a genuine RED**, see below |
| **Whole suite** | **NOT re-run — §5 rule 60 requires saying so.** **Zero `rosetta-extensions` files changed** (the fence is routed, not built); iter-132's clean whole-suite run stands on the same rext tree (`223e4a6`), as it did for iters 133–137. Stated as a gap, not characterised as covered. |
| **Suite wall-time** | not quoted — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands |

## A RED that was iter-137's, and it escaped both of iter-137's gates

`test_anchor_offset_guard.py::TestAntiVacuity` failed on `dependency_map.md:9` — **a line iter-137 wrote
one commit earlier**, citing `` `services/README.md:39` `` when **six files here are named `README.md`**.
Fully qualified to `corpus/services/README.md:39`; `_ambiguous` is `[]` and the suites pass.

**Both of iter-137's gates missed it, in the same direction:** the guard family reported
`anchor_offset_guard` **NOT-RUN** (commit-scoped, no `--range`), and it **was not among iter-137's nine
scoped suites** — which were picked by *topic* (`platform_alignment`, `claim_twin`, …) when what that iter
had actually rewritten, at 29 sites, was **anchors**. `D-M257x-138-5` states the rule: **choose the suites
by what you CHANGED, not by what you were writing ABOUT.** *Disclosed* is not *covered*.

Booked as the loop working, too: the escape survived exactly one iter, was caught by an anti-vacuity
control built for precisely this, and cost one line.

## Close — 2026-08-08

**Outcome:** the adjudicated citation set repaired at **9 sites in 6 files**, every one naming a construct
instead of a fresh line number — and an attempt to census the class that **iter-139 audited and
retracted** (*"127 of 222 (57.2 %)"* was **0-for-12 on audit**; see the banner above and
`iter-139/decisions.md`). The repairs stand; the number does not. The fence
`corpus_citation_guard.py` excluded this class *"as not mechanically decidable"*; that is true of the
claim and **false of rot**, so the milestone's largest measured citation defect sat in a declared blind
spot it did not need to be in.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; no reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 134–138 took no reading, so the metric is UNMEASURED not unmoved — §9's iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**2 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-138-1` (an exclusion is only as narrow as the predicate that justified it) ·
`D-M257x-138-2` (the population, with its denominator and its undecidable buckets) · `D-M257x-138-3`
(name the construct; re-pinning restarts the clock) · `D-M257x-138-4` (the probe is evidence, the fence
is tooling — different repos, and the fence is not built at the end of an iter).
**Side-deliverables:** `dependency_map.md:9`'s ambiguous `README.md` head — **iter-137's own defect**,
caught by this iter's suite choice and fixed here (`D-M257x-138-5`).
**Routes carried forward:**
- **`FIX-M257x-iter138-anchor-rot-fence` (NEW, and it now has a measured denominator)** — build the
  same-file bare-`:NN` rot detector in `rosetta-extensions/stack-core`, with (a) same-file vs
  continuation-pin disambiguation to shrink the 241-strong undecidable bucket, (b) a mutation control,
  (c) an anti-vacuity control that can fire, (d) the 127/222 population as its opening baseline. **This
  supersedes `FIX-M257x-iter135-bare-pin-blind-spot`, which is hereby CLOSED as measured** — the blind
  spot is real, it is 57.2 % defective, and it is fenceable.
- **`FIX-M257x-iter138-127-rotted-pins`** — the 127 measured rotted pins are a *work list*, not a sample.
  9 closed this iter; the residue is enumerable by re-running `rot-probe.py`.
- `FIX-M257x-iter135-adjudicated-live-defects` — remainder: `clerk-integration.md:126` ·
  `backend.md:13`'s dangling *UNEVEN* cross-ref (**and iter-137 made *UNEVEN* itself stale — M810 has now
  landed for both**) · `sentinel.md:5`'s grep receipt · `ai-readiness.md:18-20` · `org-repos.md:227`,
  `:370`, `:43` · `ai_architecture.md:111`,`:224` · `next-web-app.md:17`,`:186` ·
  `external_services.md:368`.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
- **CLOSED this iter:** `FIX-M257x-iter135-bare-pin-blind-spot` (measured, superseded by the fence route).
**Lessons:**
1. **An exclusion is only as narrow as the predicate that justified it.** *"Not mechanically decidable"*
   was established for content and applied to the whole class. Name the predicate you actually tested.
2. **A repair that restores the failing FORM fixes an instance and preserves the class.** Re-pinning is
   the fast move and the wrong one; the durable citation is the construct name plus a substring that
   `grep` returns uniquely.
3. **Measure before building the fence, and publish the undecidable buckets.** 588 pins is not the
   denominator; 222 is. A fence sized against 588 would have been sized against a number that includes
   241 cases it cannot see.
4. **Choose your test suites by what you CHANGED, not by what you were writing ABOUT.** iter-137
   rewrote 29 anchors and ran the topic fences; the anchor fence was NOT-RUN in the family *and* absent
   from its scoped set. Two mechanisms, one blind spot, same commit.
5. **Adjudicator accuracy runs both ways and both should be booked.** iter-136's seat named a wrong
   candidate with a right number; iter-138's seven targets all held when opened.
