# iter-83 — decisions

## D-M257x-83-1 — the union is 43 and the intersection is 14; iter-82's figures are corrected, not defended

iter-82's `29 + 30 − 41 = 18 ≠ 15` is a **category error, not an arithmetic slip**: `29`/`30` count
blocker **blocks**, `41` counted distinct **anchors**. Re-derived from `iter-82/raw/` with an extractor
whose positive control is that it reproduces iter-82's own per-seat table and both published totals
(152 / 59) exactly: **D₁₅ = 28, D₁₆ = 29, union = 43, intersection = 14**, and 28 + 29 − 14 = 43.

The intersection correction is self-evidencing: iter-82's prose says *"15"* while the list it prints
contains **14 items**. The mechanical count agrees with its list.

**Why this had to be settled before anything else in the run:** the recall estimate is derived from the
overlap, and the recall estimate is what says whether a future zero means anything. Deriving it from
`18` or from `15` gives materially different answers.

## D-M257x-83-2 — the mechanism is the ABSENCE of a per-anchor post-condition, and it is general

Measured, not argued: **109 of 147** gradeable booked findings landed inside a repair hunk (**74.1 %**);
**38** did not; **35 of the 38** are in files the repair opened and edited.

- **H1 (partition gap) REFUTED** — only 3 misses are in unopened files, and all 3 are outside the read's
  own 40-file set.
- **H2 (estimated `~` membership) REFUTED** — misses land on exact-count predicates (`external_services.md`
  ×5 = P8's exact 9; `storage.md` ×3 = P9's exact 3) as readily as on `~` ones.
- **H3/H4 UPHELD**, in the strongest form: in `graphql-wundergraph.md` the repair rewrote `:177` — a
  finding the adjudication **REJECTED** — and left `:13`, booked as **B1** by **both** readings and
  **UPHELD**. Pinned mechanically as a regression test.

**The criterion was "I have swept this file for this predicate."** Not *"no member survives"*, not even
*"every booked member is fixed."*

**Consequence, and it is binding:** the mechanism is not P4-specific — the 38 misses span 16 files and
both count styles — so **all eleven "discharged" verdicts are UNPROVEN**. They are re-derived as
membership questions at iter-84, not trusted. `FIX-M257x-iter83-eleven-discharges-unproven`.

## D-M257x-83-3 — build the reach fence rather than repair the site

The site is one edit. The mechanism produced 38 of them and would produce more at iter-85's repair,
which is *larger* than iter-81's. §8's standing rule — *fence so it cannot silently recur* — and
TOK-05's P4 ordering (**derive, else fence, else declare**) both point the same way, and iter-85 needs a
gradeable repair or it will reproduce iter-81 exactly.

`FENCE-M257x-iter83-repair-reach` — `repair_reach_guard.py`, 16 behaviour tests + a **6-mutant battery
(5 RED kills, ≥3 distinct signatures, 3 of them inversions, + 1 declared-GREEN no-op control that must
survive)**, watched RED on a real answer key (`iter-76/raw/` × `328ece5`) before being trusted.

**Why a new fence and not a widening of an existing one:** `repair_postcondition` is keyed on the tree
the commit *produced* and `repair_leak_guard` on the prose the commit *removed*. Both are blind by
construction to a finding never opened — there is nothing in such a commit to key on. The only artifact
that can see it is the input ledger, which neither fence reads. Same reasoning as `D-M257x-59-2`: extend
the code, never the guard's subject.

## D-M257x-83-4 — `repair_leak_guard` was not run, and the reason is structural

It exits **1** on `328ece5^..328ece5` naming 3 sites, and is **absent** from the six guards iter-81's
commit message lists. The reason is not carelessness: it declares `FENCE_KIND = "standalone"`, so
`repair_postcondition.py`'s **derived** registry — which selects only `postcondition`-kind fences —
never reaches it. **10 of the 14 guards standing at iter-81 were `standalone`** and 4 `postcondition`
(11 of 15 once this iter adds one) — derived in one `ast` pass over `FENCE_KIND`, not counted by hand.

A repair choosing its guard list by hand is **§2 of this milestone's own protocol doc** — the
hand-maintained tuple nobody updates — reproduced inside the machinery built to end it.

**Not fixed this iter** (it is a third line beyond this iter's declared shape, and the right fix touches
how every guard is invoked). Routed `CHECK-M257x-iter83-standalone-is-the-forgettable-class` → iter-85.
The 3 live leaked sites are routed to iter-84 with the adjudication, **not repaired here** —
adjudication-before-repair is binding.

## D-M257x-83-5 — iter-81's record is RECOVERED, and the recovery re-grades it

Written at iter-83 under an explicit non-contemporaneous banner, every field tagged `[git]` / `[msg]` /
`[iter-83]` / `UNRECOVERABLE`. Two fields are marked unrecoverable rather than inferred: the per-seat
file partition, and whether any seat recorded a deliberate skip. The second is the damaging one — it
makes an omission and a decision indistinguishable at all 38 sites.

**The record re-grades iter-81 `closed-fixed` → `closed-fixed-partial`.** Its planned scope was *"repair
the 11 predicates"* and 38 of its own input findings were never written. Re-grading a closed iter from a
later one is unusual; it is done here because the grade was the *message's own*, no plan artifact ever
carried it, and this iter is the first to have measured it. What landed is not diminished.

## D-M257x-83-6 — a reach denominator this milestone cannot currently compute

The fence grades against **booked**, not **upheld**. iter-76's adjudication recorded rejection
*mechanisms* and *counts* (12, by five mechanisms) but **never per-anchor verdicts**, so the 12 cannot be
subtracted from any reach measurement. 74.1 % is a **lower bound** on reach-against-upheld.

This is the same bookkeeping class as iter-81's missing record — a gate-critical result recorded as a
summary rather than as a ledger. Routed `FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger` →
iter-84, which must emit per-anchor verdicts for the 43 so this does not recur one iteration later.

## D-M257x-83-7 — what a zero would establish, stated in the record and not only in the report

Clause 5 is met only by a reading that returns zero. **Not re-cut, not interpreted, not argued.** But
the milestone should not close on an implied claim, so the bound is written down: at the measured
per-pass recall (≈ 50 %), a paired zero bounds the residual at roughly **R ≤ 2 at 95 % confidence** —
and that is *optimistic*, because the two readings share a briefing, a file set, a partition and a
model, and correlated blind spots inflate the intersection, deflate `N̂`, and flatter recall.

**A zero means two independent readings found nothing, which bounds the residual at a small number —
not at none.** Four recall-lift levers were costed (`CHECK-M257x-iter83-recall-lift-options`); two are
free and attack the correlation bias directly (deal each reading of a pair a *different* hand; re-partition
the confirming reading per §5 rule 18(b)). **None implemented — the ask was an assessment.**

## D-M257x-83-8 — §5 gains rule 40

*"A repair unit is not a repair post-condition."* The lesson generalizes past this iter, so the protocol
doc is updated in the same commit (skill § Protocol evolution). It is the corollary TOK-05 was missing:
`D-M257x-59-1` changed the **unit** of repair from claim to predicate and said nothing about **when a
predicate is done**.

## D-M257x-83-9 — rule 40 was tested on its author, within the hour, and it failed

**Recorded rather than quietly corrected**, because a silent fix here would be the same act the whole
iteration is about.

I published *"**9 of 14** guards are `standalone`"* — in `platform-alignment.md` rule 40, in this iter's
`progress.md` (three times), in `decisions.md`, in the milestone ledger, in `iter-81/progress.md`, and in
**both commit messages**. It was a **hand count**, taken from a `grep` I had run and eyeballed.

The correct figure is **10 of 14** (4 `postcondition`) at iter-81, and **11 of 15** once this iter adds
one. It was caught **by the pre-commit hook's own output on the very commit that shipped rule 40** —
`repair-postcondition: 4 participating fence(s) …; 11 standalone` — printed on screen, contradicting the
number in the message being committed.

Three things follow, and the third is why this is a decision and not an erratum:

1. **The figure is now derived, not counted** — one `ast` pass over every `*_guard.py`'s `FENCE_KIND`,
   the same mechanism `repair_postcondition.py` already uses for its registry. Wherever it is stated it
   is stated **with its denominator and its moment** (14 at iter-81, 15 after), because it changes every
   time a guard is added — which is exactly the shape of claim this milestone spends its life repairing.
2. **`D-M257x-59-1`'s ordering — *derive, else fence, else declare* — applies to the numbers a milestone
   states about ITSELF**, not only to platform facts. Every hand-counted scalar in a plan document is a
   `P11` waiting to happen, and P11 (*"false scalars/sets against source"*) is one of the eleven
   predicates this milestone is repairing.
3. **The correction lands as a follow-up commit, not an amend.** The wrong figure is in two commit
   messages that are already objects; rewriting them would erase the evidence that rule 40 caught its
   own author. `platform-alignment.md:74` — *a commit message is testimony, not evidence* — cuts both
   ways: the messages stay, and the **documents** are the record.

**This is the fifth-or-later occurrence in this milestone of *"the author of a newly written rule
violated it while writing it"*** (harden pass 19 counted four). That it was caught in minutes, by
machinery, rather than in two iterations by a blind read, is the only part of it that is progress.

## Unchanged routes

`FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
(**NOT DECIDED**) · `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
`CHECK-M257x-iter77-zsh-modifier` · `CHECK-M257x-iter77-developer-dir` ·
`CHECK-M257x-iter70-studio-room-lines` · `RF-M257x-iter71-run-returns-a-tuple` · RF-2/3/7–14 ·
`CHECK-M257x-iter76-seat-ref-discipline` (3rd occurrence) ·
`CHECK-M257x-iter82-commit-message-narration` (**stays SEPARATE from `CHECK-iter77`**, per
`D-M257x-82-3` — not merged) · `DEF-M257x-iter80-storage-prod-bucket` (**escalated, held; `storage.md`
`:55`/`:154`/`:181` unchanged**).
