**Type:** tik · `iter_shape: adjudication`

# iter-135 — twelve seats, six independent adjudicators, and the deviation ran the OTHER way

Closes `FIX-M257x-iter131-adjudication-independence`, the milestone's oldest unactioned route — passed
over by iters 132, 133 and 134, and named in iter-134's close so it could not lapse quietly.

---

## 1. Method

The twelve seats iter-131's coordinator adjudicated **itself** were dealt to **six independent
adjudicators**, one per seat-letter B–G (matching `adj-1`'s r33+r34 pairing), each running
`iter-131/adjudicator-brief.md` **unmodified**, each under the brief's hard bars: no `knowledge/plan/**`
except the brief and its own two seat reports, no other adjudicator's output, read-only, *a seat's
citation is not evidence until you open it.*

**One instruction was added, and it was load-bearing:** the corpus has been repaired since (iters
132–133), so an adjudicator grades the claim **as booked** and marks `UPHELD (since-repaired)` where the
defect was real and is now fixed. Without it, every claim iters 132–133 repaired would have been
**rejected for having been fixed** — scoring the repair as a refutation.

## 2. The numbers

| adjudicator | claimed | upheld | rejected | class | cannot-settle | wrong-tree | rate |
|---|---|---|---|---|---|---|---|
| `adj-B` | 4 | 4 | 0 | — | 0 | 0 | 100 % |
| `adj-C` | 12 | 12 | 0 | — | 0 | 0 | 100 % |
| `adj-D` | 7 | 6 | 1 | `true-at-its-ref` | 0 | 0 | 85.7 % |
| `adj-E` | 16 | 16 | 0 | — | 0 | 0 | 100 % |
| `adj-F` | 9 | 9 | 0 | — | 0 | 0 | 100 % |
| `adj-G` | 12 | 10 | 2 | `minor-not-blocker` | 0 | 0 | 83.3 % |
| **B–G total** | **60** | **57** | **3** | | **0** | **0** | **95.0 %** |
| `adj-1` (iter-131, already published) | 20 | 19 | 1 | | 0 | 0 | 95.0 % |
| **all 14 seats, independently adjudicated** | **80** | **76** | **4** | | **0** | **0** | **95.0 %** |

**Against the coordinator's own sheet — 80 claimed, 68 upheld, 8 rejected, 4 CANNOT-SETTLE, 89.5 %.**

> ## The deviation ran the OPPOSITE way to the prediction.
>
> This iter's `overview.md` hypothesised that self-adjudication had **inflated** the reading, and that
> independent agents would overturn at a rate above `adj-1`'s single data point. **Measured: independent
> adjudicators uphold MORE than the coordinator did (95.0 % vs 89.5 %), reject fewer (4 vs 8), and
> return ZERO `wrong-tree` — the eighth consecutive reading to do so.** The self-adjudication was
> **conservative, not inflationary.** The hypothesis is refuted and the refutation is the headline.

**And the sharper half: all FOUR of the coordinator's CANNOT-SETTLE were SETTLED.** Not argued away —
opened. `adj-C` settled the `infrastructure` question by locating a genuine clone at exactly
`13c248e6` and reading it (zero `module "cms"` org-wide, positive control 12 `module "` hits in the same
file). `adj-D`, `adj-F` and `adj-G` converged on the root-mount count the coordinator called
*"disputed in three consecutive readings"* and left unsettled.

## 3. What the independence actually bought — the framing corrections

**All six adjudicators returned at least one predicate-framing disagreement**, which is the class
`adj-1` demonstrated at iter-131. Three are load-bearing:

1. **`adj-C` independently reproduced `adj-1`'s premise-vs-inference correction — and found it in the
   brief.** Both its seats booked *"`infrastructure` has never been in a clone set"* as the false
   predicate. It refused: **that premise is TRUE** (15 trees under `stack-demo`, none of them
   `infrastructure`); the falsehood is the **inference**. It then observed that **the brief's own example
   predicate carries the same defect** — so the instrument was teaching the error it exists to catch.
   *Two blind adjudicators, one at iter-131 and one here, converged on the same correction.*
2. **`adj-F` refuted a seat's diagnosis while upholding its number.** The seat named
   `/ai-readiness/unsubscribe/:token` as the missing 8th root route; that route is **already in** the
   corpus's 7-row table. The real omission is **`/v1/labs/:slug/workspace.tar.gz`** (`labs_admin.go:40`)
   — *inside* `internal/web/backend/`, which **inverts the seat's scope argument**. In its own words: *a
   repair driven by the seat's report would fix nothing.* The route matters more than the count — its own
   comment says it is *"OUTSIDE the write group — OPTIONAL auth"*, it is wired unconditionally at
   `backend.go:301`, and it serves a tarball. **The milestone's fourth security-surface understatement.**
3. **`adj-E` overturned a causal story that would have misdirected the remedy.** Seat r33-E called five
   mis-anchors *"introduced by the very repair"*; `adj-E` checked out each authoring commit and found
   **all five were correct when written** and rotted (+2/+3/+8/+14) from unrelated insertions **above**
   them. Remedy changes from *"repair harder"* to *"fence the form"* — and it notes
   `corpus_citation_guard.py` **declares bare `:NN` pins excluded outright**, so all five sit in its
   stated blind spot. It also showed the brief's repair-induced test **cannot** find such a commit
   (`git log -L` on the citing line misses an edit made above it), so **2 of its own 3 positives are
   mis-attributed**.

**Two independent confirmations of repairs made blind to them:** `adj-D` and `adj-F` both surfaced the
private-module defect (3 vs 5) that **iter-133 repaired without seeing their reports**, and `adj-C`
recorded the M810 cluster fix as `UPHELD (since-repaired)`. Blind agreement with a repair is the
strongest evidence this milestone has produced that a repair was aimed correctly.

## 4. Test gates

- **Guard family: 18 GREEN · 0 RED · 4 not-run** (commit-/input-scoped members, no `--range`/`--ledger`).
  Not a whole-family green; the runner says so.
- **Zero `corpus/**` and zero `rosetta-extensions` files changed** — this is an adjudication. No
  code-test gate applies and none is claimed.
- **Whole suite not re-run; §5 rule 60 requires saying so.** Nothing executable has changed since
  iter-132's clean run (`1 failed · 1208 passed`, the 1 being the standing RED). **Stated as a gap.**

---

## Close — 2026-08-08

**Outcome:** the milestone's oldest route is closed by **six independent adjudicators over the twelve
self-adjudicated seats**. **The disclosed deviation ran the opposite way to the prediction** —
independents uphold **95.0 %** against the coordinator's **89.5 %**, reject **4** where it rejected 8,
and return **zero `wrong-tree`** — so the self-adjudication was conservative, not inflationary. **All
four CANNOT-SETTLE blockers were settled by opening the evidence**, including the root-mount count
disputed across three readings. What independence bought was **framing**: all six returned predicate
corrections, one of them locating the defect **in the brief itself**, one refuting a seat's diagnosis
while upholding its number (the 8th root route is a **tarball endpoint with OPTIONAL auth, wired
unconditionally** — the milestone's 4th security-surface understatement), and one overturning a causal
story that would have sent the remedy to the wrong place.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **`P`/`N` are deliberately NOT re-cut**: re-adjudicating an
existing seat set is not a fresh sample, and treating it as one would manufacture a movement out of a
method fix.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no successor strategy is authorable — `TOK-08`'s sealed refutation branch bars one; running under the user's direct brief**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-135-1`, `D-M257x-135-2` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none.
**Routes carried forward:**
- **NEW — `FIX-M257x-iter135-adjudicated-live-defects` (the big one):** the six sheets confirm **a large
  set of blockers still live at HEAD**, untouched by iters 132–133. Named by their adjudicator:
  the **8th root mount** + the 7-row tables (`security_compliance.md:250`,`:265`,`:293`;
  `architecture_overview.md:406`) · `shared_libraries.md:77` (analytics-go wiring cited `main.go:507-508`,
  measured `:494-495`) · `security_compliance.md:156` (→`clerk-integration.md:44`, not `:40`) ·
  `clerk-integration.md:126` (DEV_LOGIN_ENABLED row) · `backend.md:13`'s dangling *UNEVEN bullet*
  cross-ref · roadrunner's prod state (`roadrunner.md:13`,`:30-32`,`:53`,`:74`;
  `architecture_overview.md:228`) · `sentinel.md:5` · `dependency_map.md:9` · `ai-readiness.md:18-20` ·
  `org-repos.md:227`,`:370`,`:43` · `ai_architecture.md:111`,`:224` · `next-web-app.md:17`,`:186` ·
  `external_services.md:368` · and `adj-E`'s five rotted `graphql-wundergraph`/`academy-backend`/
  `ant-academy` anchors. **Each is adjudicated, cited and named — this is a work list, not a sample.**
- **NEW — `FIX-M257x-iter135-brief-teaches-the-error`:** `adj-C` found the adjudicator brief's own
  example predicate carries the premise-vs-inference conflation. **Fix the brief before the next
  reading**, or it will keep teaching it.
- **NEW — `FIX-M257x-iter135-bare-pin-blind-spot`:** `adj-E` showed five mis-anchors sit in
  `corpus_citation_guard.py`'s **declared** exclusion of bare `:NN` pins, and that the brief's
  repair-induced test structurally cannot attribute them. Two instrument gaps, both stated by their own
  code.
- Still open: `FIX-M257x-iter131-predicate-sets-not-enumerated` ·
  `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it` ·
  `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` · `FIX-M257x-iter133-two-fives-need-a-fence` ·
  `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer`.
  **`FIX-M257x-iter131-root-mount-count-underived` CLOSES** — settled at **8** by three independent
  adjudicators.
**Lessons:**
1. **A disclosed deviation is worth measuring, not just confessing.** iter-131 disclosed the
   independence gap honestly and predicted its direction. The disclosure was right and **the prediction
   was wrong** — and only re-running it could tell the difference.
2. **Independence buys FRAMING, not verdicts.** The per-anchor truth barely moved (95.0 % vs 89.5 %,
   both high). Every adjudicator changed how a predicate was *stated* — and one found the error in the
   brief, i.e. in the instrument that trains the next reading.
3. **Blind confirmation of a repair is the strongest signal available here.** Two adjudicators, barred
   from reading iters 132–134, independently surfaced defects those iters had already fixed.
