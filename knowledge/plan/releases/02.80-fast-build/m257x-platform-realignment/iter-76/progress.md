# iter-76 — progress

**Type:** tik, under `TOK-05`. Planned deliverable: **the paired graded reading**, taken only once
both known-bad citation classes were closed (iter-74, iter-75).

---

## Phase A — the reading, as run

14 blind seats, two readings of the identical partition, all reports on disk under `raw/`.
Every seat returned; every seat stated its per-file `wc -l` positive control.

| seat | #13 | #14 |
|---|---|---|
| A | 9 | 8 |
| B | 4 | 6 |
| C | 13 | 16 |
| D | 6 | 6 |
| E | 9 | 9 |
| F | 14 | 13 |
| G (diff) | 22 | 17 |
| **total** | **77** | **75** |

## Phase B — the pre-registered predictions, graded unsoftened

| # | prediction | outcome |
|---|---|---|
| 1 | each of `N₁₃`, `N₁₄` in **[0, 10]** | **BADLY FALSIFIED** — 77 and 75 |
| 2 | recall below 60% for both | **holds** — the two readings agree on classes but differ substantially site-by-site (e.g. C books 13 vs 16 over identical files; F 14 vs 13 with a different mix) |
| 3 | union > max of the two | **holds** — each reading books blockers the other misses |
| 4 | induced term (seat G's diff scope) below 6 | **FALSIFIED** — G booked 22 / 17 |
| 5 | **the iter-74 and iter-75 classes contribute ZERO blockers** | **HOLDS** — see Phase D |

Prediction 1 was not merely wrong, it was wrong by an order of magnitude. That is the iteration's
finding, not its embarrassment: **a prediction written to be safe is not a prediction**, and this
one was written to be refutable and duly was.

## Phase C — the headline: my guards were GREEN over all of it

Five corpus guards were **OK** at this iter's open and are OK now. The reading found **~150
blockers**. Both statements are true, and the reason is printed by the guard itself on every run:

```
platform_predicate_guard: reach — G1 99 profile site(s) over 8 token(s); … G2 3 repo-count claim(s);
  G5 24 migration claim(s) = 1 enumerated + 21 free prose UNREACHED + 2 ref-pinned; …
```

**`G5` reaches 1 claim of 24. `G2` reaches 3 repo-count claims.** The corpus's false predicates live
in the *unreached* remainder, and the guard has said so, honestly, on every single run. The
milestone had already booked that as *"the declared clause-5 residual with a named owner."* **What
this reading did was measure what is actually inside it** — and the answer is not a long tail of
cosmetics.

> **This is the run's governing sentence, demonstrated rather than quoted: every GREEN verdict is a
> statement about reach.** Five green guards, ~150 real findings, no contradiction between them.

The dominant classes, each booked independently by multiple seats across both readings, and each
**confirmed against the guard's own derived denominators** rather than against another document:

| corpus claims | derived truth (`platform 0dab54d`) |
|---|---|
| the `cms` / `jobsimulation` / `roadrunner` container *"still starts locally"* | **no compose service at all** for any of them |
| *"compose declares nine services"* | **8**, or **10** with the `include` — never 9 |
| *"`make up` … the default `graphql` profile"* | `PROFILE ?= core`; `graphql` is **not one of the 8 legal tokens** |
| roadrunner *"still in `repos.yml`, 1 of 9"* | `repos.yml` has **6**: app · sentinel · storage · messenger · next-web-app · studio-desk |
| *"`STORAGE_RPC_ADDR` is read by `main.go`"* | **0** read sites at the ref the claim names |
| the default profile starts nine containers | `core` selects **5**: backend · gotenberg · postgresql · redis · sentinel |

## Phase D — adjudication started, and it is already earning its keep

**Prediction 5 HOLDS.** No seat booked a blocker inside the classes iter-74 and iter-75 closed. The
citations those iters made reachable are being read correctly.

**And the first adjudicated finding is a FALSE POSITIVE, with a mechanism that will recur.**
`r13-F` B13/B14 and several sibling findings across other seats book
`` `docker-compose.yml:311` `` / `` `:362` `` / `` `:352` `` / `` `:337-341` `` as *"past the end of a
271-line file."* Traced through the guard:

```
frontend_architecture.md:11  cites docker-compose.yml:311
   ref block-pinned pin=2adcf71   file has 387 lines   classify -> None
```

Both citations sit in blocks that **pin `2adcf71`**, where `docker-compose.yml` is **387 lines**.
Line 39 names its pin in words — *"since platform `2adcf71`"*. **§5 rule 33: a claim is settled at
the ref the claim itself names**, which the briefing states in a section of its own. The seats
graded dated claims against today's checkout, and both booked `medium` confidence — they hesitated,
correctly.

**So the routed count is a hypothesis here too.** ~150 raw blockers is an upper bound containing at
least one systematic false-positive class of my own instrument's making (the briefing told seats the
rule; it did not make the rule mechanical for them).

## Phase E — routed, per the pre-registration

The iter's own escalation condition, written before any seat launched: *"More than ~15 union
blockers → measure and route; do not repair inside this iter."* At ~150 it fired decisively.
Repairing 150 findings in the iteration that discovered them is exactly the *"half-repairing a
uniformly-wrong corpus is worse than leaving it"* failure §5 rule 19 exists for.

**Routed to iter-77 as `FIX-M257x-iter76-read-union`**, with three binding conditions taken from
this milestone's own evidence:

1. **Adjudicate every finding before repairing it.** Four routed counts in a row have collapsed on
   adjudication in this milestone (64→5, 23→1, 21→0, 92→0), and this one already has a proven
   false-positive class.
2. **Repair by PREDICATE, not by claim** (`D-M257x-59-1`). Six predicates cover most of the ~150:
   *the husk container still starts* · *nine compose services* · *the `graphql` profile* · *N repos
   in `repos.yml`* · *`STORAGE_RPC_ADDR` is read* · *the default profile's membership*. Each has a
   legal set already derived by `platform_predicate_guard` and printed on every run.
3. **The repair is not done until the reach hole is closed.** G5 reaching 1 of 24 and G2 reaching 3
   is *why* these survived 75 iterations. Repairing the sites without extending the fence leaves the
   next drift equally invisible.

## Close — 2026-08-04

**Outcome:** the paired graded reading was **taken** — 14 blind seats, identical partition, all
reports on disk — and it returned **77 blockers in reading #13 and 75 in #14**, against a
pre-registered ceiling of **10**. **Clause 5 is not met and was never close.** The headline is not
the count but its coexistence with a green board: **five corpus guards were OK throughout, and they
were right to be**, because `platform_predicate_guard` prints on every run that **G5 reaches 1
migration claim of 24 and G2 reaches 3 repo-count claims** — the corpus's false predicates live in
the twenty-one it cannot see. *Every GREEN verdict is a statement about reach*, demonstrated with a
measurement rather than quoted. The dominant classes are each confirmed against the guard's **own
derived denominators**: no compose service exists for cms / jobsimulation / roadrunner while six
documents say the container *"still starts locally"*; `repos.yml` has **6** entries while the corpus
says *"1 of 9"*; `core` selects **5** containers while the corpus says nine; `graphql` is not one of
the 8 legal profile tokens. **Prediction 5 HOLDS** — no seat booked a blocker inside the classes
iter-74 and iter-75 closed. And the first adjudicated finding is a **false positive with a
systematic mechanism**: the *"past the end of a 271-line file"* class sits in blocks pinning
`2adcf71`, where the file is **387 lines** — so ~150 is an **upper bound**, not a work item.
**Type:** tik
**Status:** closed-fixed — the planned deliverable was the reading, and the reading landed, was
graded against its own pre-registrations unsoftened, and was routed per its own pre-registered
escalation.
**Gate:** NOT MET — clause 5 measured for the first time in eighteen iterations and found open by a
wide margin. 4 of 5, unchanged.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (3 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-76-1` (clause 5 is the largest open class, not a residual), `D-M257x-76-2`
(five GREEN guards and ~150 findings are not in contradiction — a fence's GREEN may never again be
quoted without its reach line), `D-M257x-76-3` (the routed count is a hypothesis; a rule stated in
a briefing is not a rule enforced by an instrument), `D-M257x-76-4` (repair by predicate, and not in
the iteration that discovered the class), `D-M257x-76-5` (one denominator disagreed with the tested
parser and is left explicitly unsettled).
**Side-deliverables:** the reading instrument's stored AS-RUN briefing for this reading
(`instrument/briefing-iter76-AS-RUN.md`), whose only deltas from iter-53's are the ground-truth shas
and an explicit ref-selection rule — recorded, not applied silently (§5 rule 25).
**Routes carried forward:**
- **`FIX-M257x-iter76-read-union`** — the ~150-blocker union, to be **adjudicated before repaired**,
  **by predicate not by claim**, and **not closed until the G5/G2 reach hole is closed**. iter-77's
  target.
- **`CHECK-M257x-iter76-seat-ref-discipline`** — seats graded ref-pinned claims against the
  checkout despite a briefing section telling them not to. Give the seats the resolved ref per
  citation, or expect this false-positive class in every future reading.
- **`CHECK-M257x-iter76-compose-service-count`** — 8 vs 9 vs 10; my one-line grep disagreed with the
  tested parser and the disagreement is recorded rather than resolved by assertion.
- Unchanged: `FENCE-M257x-iter70-line-or-port` · `RF-M257x-iter71-run-returns-a-tuple` ·
  `CHECK-M257x-iter70-studio-room-lines` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**)
  · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `FENCE-M257x-iter54-refs-block` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter58-derive-preregistrations` · `CHECK-M257x-iter52-second-ai-manager` ·
  `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` · `-baseline-refs` ·
  RF-2/3/7–13.

**Lessons:**

1. **Quote a fence's GREEN with its reach line or not at all.** *"Five corpus guards OK"* has closed
   every iter since iter-59 and, read alone, implied something it never claimed. G5 reaches 1 of 24.
2. **A rule stated in a briefing is not a rule enforced by an instrument.** The seats were told, in
   a section of its own, that a claim is settled at the ref it names — and graded dated claims
   against the checkout anyway. The same predicate is already mechanical one directory away.
3. **Take the reading before believing the distance.** Eighteen iterations described clause 5 as a
   residual. One reading measured it. The two descriptions are not close, and nothing cheaper than
   the reading would have revealed that.
4. **A prediction written to be safe is not a prediction.** Two of five pre-registrations were
   falsified, one by an order of magnitude, and the falsifications are the most useful output of the
   iteration.
