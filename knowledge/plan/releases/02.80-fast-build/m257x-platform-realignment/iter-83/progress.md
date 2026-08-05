# iter-83 — progress

**Type:** tik, shape `tooling` (protocol-codified), under `TOK-05`. Deliverable: **the diagnosis of the
repair machinery, and the fence that closes the hole it found.** No corpus claim is repaired in this
iter — adjudication-before-repair is binding (iter-80) and the adjudication is iter-84's.

---

## THE HEADLINE

**iter-81 reported eleven predicates discharged. Measured against its own input, it reached 109 of 147
booked findings — 74.1 % — and 35 of the 38 it missed were in files it opened and edited.**

The mechanism is not a partition gap, and not estimated membership. It is that **the repair had no
per-anchor post-condition at all**, so "the predicate is discharged" and "I believe the predicate is
discharged" were the same statement. That is this milestone's signature defect — *a check that reports
a state without measuring it* — reached at last into the repair machinery.

---

## 1. The overlap arithmetic, settled before anything was derived from it

iter-82 reported `N₁₅ = 29`, `N₁₆ = 30`, union `41`, and separately *"the two readings share only 15 of
41 anchors."* Those cannot all hold: **29 + 30 − 41 = 18**.

**They cannot hold because they are counts of two different objects.** `29` and `30` are **blocker
blocks** (`### B<n>` headings). `41` was a count of **distinct anchors**. Subtracting one from the other
is the error, and once the units are made consistent the inconsistency disappears.

Re-derived mechanically from the 14 reports on disk:

| quantity | value |
|---|---|
| blocker blocks, reading #15 / #16 | **29 / 30** (total 59 — reproduces iter-82 exactly) |
| **distinct primary anchors**, #15 / #16 | **28 / 29** |
| **union** | **43** |
| **intersection** | **14** |
| check | 28 + 29 − 14 = **43** ✓ |

**Positive control on the extractor** (§5 rule 2): it reproduces iter-82's per-seat table exactly
(A 3/4 · B 3/6 · C 4/5 · D 4/3 · E 6/6 · F 3/2 · G 6/4) and both published totals (152 pre-repair,
59 post-repair). It was not tuned to reproduce the union — the union is what it disagrees on.

**Two corrections to iter-82, both small and both stated:** the union is **43**, not 41; the
intersection is **14**, not 15. iter-82's own prose says 15 while the list it prints has **14 items** —
the mechanical count agrees with its list, not with its sentence.

### What follows: the recall estimate

Two-sample capture–recapture over the settled figures:

| estimator | population `N̂` |
|---|---|
| Lincoln–Petersen `D₁₅·D₁₆/I` | 58.0 |
| **Chapman** (bias-corrected) | **57.0** |

→ **per-pass recall ≈ 49 % / 51 %. Union recall ≈ 75 %.**

This is the third independent measurement in this milestone to land at *roughly half*, and it agrees
with iter-50's paired same-tree experiment (43–48 %) and with the <60 % prior that has held across every
paired reading here.

**The independence assumption is false, and it biases in the unsafe direction.** The two readings share
the briefing, the file set, the partition method and the model. Correlated blind spots **inflate the
intersection**, which **deflates `N̂`** and **flatters recall**. So `57` is a *lower* bound on the true
population and `≈50 %` is an *upper* bound on per-pass recall. Stated here rather than left implicit,
because every claim in §5 below rests on it.

---

## 2. What a zero from this instrument would — and would not — establish

Asked for plainly, and answered plainly, so that if the gate closes it closes on an honest claim.

**What a zero WOULD establish.** That two independent 7-seat blind readings of the 40-file corpus, run
under a frozen briefing against re-derived platform ground truth, each found nothing they judged a
blocker. That is a real and demanding result: the same instrument returned 77/75 in November's tree and
29/30 after one repair pass, so it is demonstrably capable of returning large numbers when there is
something to find. **It is exactly what clause 5 defines as met**, and the user has ruled three times
that this is the grading rule. Nothing here re-cuts it.

**What a zero would NOT establish.** That the corpus is free of defects of this class. At the measured
per-pass recall, a single defect survives *both* readings with probability ≈ **0.25**:

| true residual `R` | P(a paired reading returns zero anyway) |
|---|---|
| 1 | **25 %** |
| 2 | 6.3 % |
| 3 | 1.6 % |
| 5 | 0.1 % |

So a paired zero bounds the residual at roughly **R ≤ 2 at 95 % confidence** — *and even that is
optimistic*, because the correlated-blind-spot bias above makes the true miss probability higher than
0.25. A zero is strong evidence that the residual is **small**. It is not evidence that it is **empty**,
and this milestone should not be read as claiming otherwise.

**The honest one-sentence form, for the close report:** *a zero means two independent readings found
nothing, which bounds the residual at a small number — not at none.*

### Raising recall without touching the frozen instrument — an assessment, not a proposal

Requested as a costing exercise; **not implemented this run**, and it changes *how much* is measured,
never *what* is measured, so the instrument stays frozen either way.

| lever | mechanism | cost | expected effect |
|---|---|---|---|
| **More seats per reading** (7 → 12) over the same partition | each seat reads fewer files, so per-file attention rises | +5 seats × 2 readings = **+10 agent-runs**, ~+70 % tokens on the read, ≈ +25 min wall-clock (seats run parallel) | the strongest single lever; per-seat load is the plausible driver of ~50 % recall |
| **A third independent reading** (#17) | a third sample of the same population | **+7 agent-runs**, ~+50 % read cost, +20 min | union recall ≈ 75 % → ≈ **88 %**; also gives a 3-sample estimator, which is far more robust than 2-sample and would settle the correlation question the table above can only flag |
| **Re-partition the third reading differently** (§5 rule 18(b)) | correlated blind spots are a property of *how the corpus was divided* | free — a different deal of the same method | attacks the specific bias that makes `N̂ = 57` a lower bound |
| **Deal each reading a different hand** | today both readings of a pair share one partition | free | the cheapest lever on the table, and the one most likely to break the correlation |

**Recommendation, unpriced into any plan:** the last two are free and attack the *bias*; the third
reading is the one that would let the milestone state a recall figure it has actually measured rather
than estimated from two correlated samples. Routed as
**`CHECK-M257x-iter83-recall-lift-options`** for the user's decision — not actioned.

---

## 3. The P4 mechanism — measured, with three hypotheses refuted

`corpus/services/graphql-wundergraph.md:13`:

> *"The `graphql` profile name survives in compose and is now simply the default profile — it no longer
> names a router service."*

**Re-derived at platform `0dab54d` (the ground-truth ref), not taken from the briefing:** the token
`graphql` appears in **no** `profiles:` key. The eight that exist are `core`, `backend`, `all`,
`storage-legacy`, `customerio-sync`, `messenger`, `studio-desk`, `frontend`; the only `graphql`
substrings in `docker-compose.yml` are inside `VITE_GRAPHQL_ENDPOINT` / `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`
URL values at `:204`, `:220`, `:236`, `:245`. `Makefile:10` reads `PROFILE ?= core`. **Both halves of
the sentence are false.**

### The measurement that decides it

Every one of the **152** anchors booked pre-repair (`iter-76/raw/`, readings #13/#14) checked against
`328ece5`'s diff at `--unified=0`, on the **old side** — the coordinate system the ledger was written
in:

| outcome | n |
|---|---|
| **touched** (inside a hunk, ±3 lines) | **109** |
| **line-unreached** — the file was edited, the anchor was not | **35** |
| **file-unreached** — the file was never opened | **3** |
| no parseable anchor | 4 |
| path not in the repaired tree | 1 |
| | **152** |

**Reach = 109/147 = 74.1 %.**

### Grading the four candidate mechanisms

| # | candidate | verdict |
|---|---|---|
| **H1** | the site was outside the assigned seat's file set (**partition gap**) | **REFUTED.** 35 of 38 misses are in files the repair *edited*. The 3 exceptions are `corpus/ops/demo/frontend-tier.md:46` and `.claude/skills/dev-up/reference.md:46` — both **outside the read's own 40-file set**, so no seat owned them for repair either |
| **H2** | membership was **estimated** (`~`) rather than enumerated | **REFUTED.** The misses fall on **exact**-count predicates as readily as estimated ones: `external_services.md` ×5 (P8, exact 9) and `storage.md` ×3 (P9, exact 3). The tilde does not predict the misses. It was a reasonable prior; it is not the answer |
| **H3** | a seat marked the predicate done without re-checking its members | **UPHELD** |
| **H4** | the discharge criterion was *"my sites are fixed"*, not *"no member survives"* | **UPHELD, and it is weaker than even that** |

### The decisive single case, and it is mechanical

In **one file**, `graphql-wundergraph.md`, the repair:

- **rewrote `:177`** — *"`make up` starts the `graphql` profile, which since `2adcf71` contains no
  router"* — a finding iter-76's adjudication had **REJECTED** (one of the 12, named in prose as the
  *"pin in a subordinate clause"* mechanism);
- **left `:13`** — booked as **B1**, the *first* blocker, by **both** readings independently, and
  **UPHELD**.

Pinned as a regression test (`test_02_the_rejected_sibling_in_the_same_file_WAS_reached`): at tolerance
3, line 177 **is** covered by a hunk and line 13 **is not**.

**So the criterion was not "no member of P survives" (the strong reading), and not even "every booked
member of P is fixed" (the weak reading). It was "I have swept this file for this predicate."** A
per-file judgment sweep. The two readings had both warned in advance that `:13` would *look* pinned by
the `2adcf71` two sentences above it and is not (§5 rule 33) — and it is exactly the site a sweep
skips.

### Is the mechanism general? **Yes — and therefore all eleven verdicts are unproven.**

The 38 misses spread across **16 files** and hit predicates with exact site counts as readily as
estimated ones. Nothing about the measurement is P4-specific.

> **The other ten predicates' "discharged" verdicts are UNPROVEN and must be re-derived, not trusted.**

Re-derivation is a **membership** question — enumerate each predicate's legal set over the corpus and
check every member — and it does **not** require re-running the frozen instrument. Routed as
**`FIX-M257x-iter83-eleven-discharges-unproven`** → iter-84.

### Two deeper layers, found while measuring

**Layer 2 — the one guard that asks the adjacent question was not run, and it is RED on this commit.**
`repair_leak_guard.py` (*"did this commit FINISH?"*) exits **1** on `328ece5^..328ece5`, naming **3**
sites where a claim the commit rewrote elsewhere still stands. It is **absent** from the six guards the
commit message lists. All three still stand at HEAD:

| site | leaked form |
|---|---|
| `CLAUDE.md:285` | `make up  # Build from local code and start (graphql profile)` — **a live P4 member, in a runnable command block, in the repo's most-read file** |
| `corpus/ops/platform-alignment.md:1249` | *"`STORAGE_RPC_ADDR` is read by `main.go`…"* at `9d00a313` — against P9's own finding of **0** read sites there |
| `corpus/services/messenger.md:122` | the `REDIS_STREAMS_INDEX` row |

Not repaired here (adjudication first). Routed to iter-84.

**Layer 3 — it was not run because the guard list was hand-maintained.** `repair_leak_guard` declares
`FENCE_KIND = "standalone"`, and the **derived** registry in `repair_postcondition.py` selects only
`postcondition`-kind fences. **10 of the 14 guards standing at iter-81 were `standalone`** — 4
`postcondition` — the class you have to *remember*. (11 of 15 once this iter adds one.)
A repair choosing its own guard list by hand is **§2 of this milestone's own protocol doc**: the
hand-maintained tuple nobody updates, which is the defect the whole milestone was opened to end.
Routed as **`CHECK-M257x-iter83-standalone-is-the-forgettable-class`** → iter-85.

---

## 4. The fence — `FENCE-M257x-iter83-repair-reach`

`rosetta-extensions/stack-core/repair_reach_guard.py` (+ 2 test files). The hole it closes is precise:

| fence | question | keyed on |
|---|---|---|
| `repair_postcondition.py` | did the repair **CREATE** a defect? | the tree the commit produced |
| `repair_leak_guard.py` | did the repair **FINISH** a claim? | the prose the commit **removed** |
| **`repair_reach_guard.py`** | did the repair **REACH its INPUT**? | the **ledger**, against the diff |

The first two are keyed on *what the diff contains* and are therefore blind **by construction** to a
finding the repair was handed and never opened: nothing was removed there, nothing was added there, and
the site reads as it did before. There is nothing to key on except the input ledger.

**Watched going RED before being trusted** (§8 rule 5) on a fixture whose answer key is two committed
artifacts, neither authored for the test: `iter-76/raw/` + `328ece5`. **16 behaviour tests + 5 battery
tests, all green.**

- Both positive controls are **fatal** (exit 2, not 0): an empty ledger and an empty diff each mean the
  pipeline broke, and reporting 0 for that is the `|| echo 0` this milestone opened on (§5 rule 8).
- A repair may decline any finding; it may not decline one **silently** — a waiver with an empty reason
  is refused.
- **Mutation battery: 5 RED mutants, 5 kills, ≥3 distinct signatures, plus a declared-GREEN no-op
  control that survives.** Three of the five are **inversions** — removal mutants cannot catch a
  predicate flipped to its opposite. The load-bearing one is `file-level-reach-accepted`, which rewrites
  the classifier so that *"the repairer had the file open"* counts as *"the finding was reached"*:
  **that mutant is not a fault injection, it is a mechanical statement of what iter-81 actually did.**

**A limit recorded rather than discovered:** reach is necessary, not sufficient. A repair can touch an
anchor's line and still leave the claim false. **A green reach report says the repair opened everything
it was given; it does not say the repair was correct.**

**A second limit, and it is a finding of its own:** the denominator is **booked**, not **upheld** —
iter-76's adjudication recorded rejection *mechanisms* and counts but **never per-anchor verdicts**, so
the 12 rejections cannot be subtracted. 74.1 % is therefore a *lower bound* on reach-against-upheld.
Routed as **`FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger`**.

---

## 5. iter-81's record, recovered

`FIX-M257x-iter82-iter81-has-no-record` discharged. `iter-81/progress.md` now exists, **headed by an
explicit non-contemporaneous banner** and with every field tagged `[git]` / `[msg]` / `[iter-83]` /
`UNRECOVERABLE`. Two fields are marked unrecoverable rather than inferred:

- **the per-seat file partition** — `iter-81/raw/` is empty; no ground-truth sheet survives;
- **whether any seat recorded a deliberate skip** — which is the most damaging gap, because it makes an
  omission and a decision indistinguishable at all 38 unreached sites.

The recovered record **re-grades iter-81 from `closed-fixed` to `closed-fixed-partial`**, on the reach
measurement. What landed is not diminished: 109 findings across 33 files, TRAP A held, and the P9
ref-relative ruling all survived iter-82's re-read.

---

## State at close

- **Gate 4 of 5, unchanged.** Clause 5 is graded only by a reading that returns zero; none was taken.
- 6 corpus guards GREEN at open and at close, with correct refs supplied.
- `stack-core` **843** tests (822 → +21), 1 non-green: the known perishable iter-48 fixture.
- Zero platform-repo edits. `storage.md:55,:154,:181` **held unchanged**
  (`DEF-M257x-iter80-storage-prod-bucket`, escalated, awaiting the user).
- Neither `FIX-M257x-iter53-union-set`, `FIX-M257x-iter56-assignment-flake`,
  `CHECK-M257x-iter38-ai-act-classification` nor RF-2/3/7–14 was touched.
- `CHECK-M257x-iter82-commit-message-narration` remains **separate** from `CHECK-iter77`, per
  `D-M257x-82-3`. Not merged.

## Routes carried forward

| handler | to | what |
|---|---|---|
| `FIX-M257x-iter82-reread-union` | iter-84 | adjudicate the **43** (not 41) before any repair |
| `FIX-M257x-iter83-eleven-discharges-unproven` | iter-84 | re-derive all 11 verdicts as membership questions |
| `FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger` | iter-84 | per-anchor verdicts, so reach can be graded against *upheld* |
| `FIX-M257x-iter83-leak-guard-3-sites` | iter-84 | the 3 live leaks, incl. `CLAUDE.md:285` |
| `CHECK-M257x-iter83-standalone-is-the-forgettable-class` | iter-85 | 11 of 15 guards run only if remembered |
| `CHECK-M257x-iter83-recall-lift-options` | user | costed; not actioned |
| `CHECK-M257x-iter76-seat-ref-discipline` | open | 3rd occurrence, unchanged |

## Lessons

1. **A repair unit is not a repair post-condition.** TOK-05 changed the unit to the predicate and that
   was right — but a unit says what to work on, and a post-condition says when you are done. iter-81
   had the first and not the second, and reported completeness on the strength of the first. → §5
   **rule 40**.
2. **Check the units before subtracting.** iter-82's `29 + 30 − 41 = 18 ≠ 15` was not an arithmetic
   slip; it was a category error between blocks and anchors, and it would have propagated into the
   recall estimate that decides what a future zero means.
3. **The forgettable class is the one that gets forgotten.** A derived registry that covers **4** of
   14 guards leaves **10** to memory, and memory is what §2 of this protocol was written about.
4. **The rule was tested on its author within the hour — see `D-M257x-83-9`.** I published *"9 of 14"*
   in five places from a hand count, and the pre-commit hook printed the derived figure on the very
   commit that shipped rule 40. **Derive, else fence, else declare** (TOK-04 P4) is not advice about
   platform facts; it applies to the numbers a milestone states about *itself*.
