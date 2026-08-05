# iter-81 — progress ⚠️ **RECONSTRUCTED RECORD, NOT CONTEMPORANEOUS**

> **Read this header before any figure below.** This file was written at **iter-83 (2026-08-05)**, two
> iterations after the work it describes. iter-81 closed leaving **no iteration record at all** — an
> empty `raw/`, no `overview.md`, no `progress.md`, no `decisions.md`, and zero mentions in the
> milestone `progress.md`. iter-82 found the gap while closing and rightly declined to author the
> record retrospectively, routing it as **`FIX-M257x-iter82-iter81-has-no-record`**.
>
> This is the discharge of that route, and it is a **recovery, not a reconstruction of intent**. Every
> field is tagged with its provenance:
>
> | tag | meaning |
> |---|---|
> | **[git]** | derived from the commit object or its diff — as hard as any other fact in this milestone |
> | **[msg]** | quoted from the commit message. **Testimony, not evidence** — `platform-migration-status.md:74` and `D-M257x-82-3` both say why, and iter-83 measured one figure in this very message to be wrong |
> | **[iter-83]** | measured at iter-83 against the artifacts, and labelled as such |
> | **UNRECOVERABLE** | the diff and the log cannot establish it. Left empty. Inferring it would be authoring history |
>
> **Nothing below is presented as what iter-81 knew or decided at the time**, except where tagged
> **[msg]** — and a `[msg]` tag is a record that the message *said* it, not that it was so.

**Type:** tik (**[git]** — the commit uses the `iter(M257x/81):` closing prefix, the protocol's iter
boundary marker). **Shape: UNRECOVERABLE** — no `overview.md` was written, so no `iter_shape` was ever
declared.

**Status: `closed-fixed`** as the commit presents it **[msg]**. **iter-83 re-grades this to
`closed-fixed-partial`** — see *"The grade this record cannot leave standing"* below.

---

## Commit **[git]**

| field | value |
|---|---|
| sha | `328ece5e770fa8e58cd7d0157f99eebe98953daa` |
| parent | `f375738` (iter-80's closing commit) |
| authored | 2026-08-05T13:10:22+02:00 |
| subject | *"iter(M257x/81): repair the 11 predicates — 33 files, 7 parallel seats, all 5 corpus guards GREEN"* |
| diffstat | **33 files changed, +410 / −251** |
| scope | `corpus/**` (11 architecture · 21 services) + `CLAUDE.md`. **Nothing outside the corpus** — verified by `git show --name-only`, so the *"zero platform-repo edits"* claim holds independently of the message asserting it |

**Note the subject line says "all 5 corpus guards GREEN" and the body lists SIX.** **[iter-83]** The
body's six are `platform_predicate_guard`, `platform_alignment_guard`, `anchor_construct_guard`,
`markdown_structure_guard`, `corpus_index_guard`, `claim_twin_guard`. The 5-vs-6 discrepancy is
internal to the message and is not resolvable from any artifact — but which guards were *not* run is,
and it turns out to be the load-bearing fact of this iteration. See below.

## Input: the eleven predicates **[msg]**

Quoted verbatim from the commit message, lines 19–29. **This list exists nowhere else** — that is why
`FIX-M257x-iter82-iter81-has-no-record` had to be discharged before the repair could be audited.

| # | predicate | claimed sites |
|---|---|---|
| P1 | dead cms/jobsimulation/roadrunner containers described as live | ~47 |
| P2 | `repos.yml` shape (9 entries / `:17-19` / roadrunner still listed) | ~12 |
| P3 | *"compose declares nine services"* | 5 |
| P4 | *"`graphql` is a live profile / the default"* | ~10 |
| P5 | *"`core` starts nine containers / six Go services"* | ~8 |
| P6 | *"`storage` is in the default set"* | 3 |
| P7 | stale compose anchors in blocks that pin nothing | ~14 |
| P8 | `external_services.md` re-point short by ~9 lines | 9 |
| P9 | *"`STORAGE_RPC_ADDR` is read by main.go at `9d00a313`"* (load-bearing) | 3 |
| P10 | wrong commit attribution (`d11a403` vs `2adcf71`) | 4 |
| P11 | false scalars/sets against source | ~12 |

Its source is [`iter-76/adjudication.md`](../iter-76/adjudication.md), which derived the same eleven
from the 140 upheld findings — that document IS contemporaneous and is the better citation for the
predicates themselves. What the commit message uniquely carries is **which of them iter-81 believed it
had discharged: all eleven.**

## Method **[msg]**

> *"seven parallel seats over disjoint file sets, each working from one re-derived ground-truth sheet,
> each re-deriving every anchor from the platform clones rather than from any prior note."*

**The per-seat file partition is UNRECOVERABLE.** `iter-81/raw/` is empty; no ground-truth sheet
survives on disk. `D-M257x-82-3` already recorded the consequence — a wrong `repos.yml:17-19` figure in
the commit message could not be traced to a sheet, because there is no sheet.

**Whether any seat recorded a deliberate skip is UNRECOVERABLE**, and this is the single most damaging
gap: it makes an omission and a decision indistinguishable at every one of the 38 sites below.

## Two rulings **[msg]**

- **TRAP A held.** Where the underlying FACT was deleted rather than moved, the claim was restated or
  dropped, never re-anchored. Re-pointing `dc:70-80` or `dc:337-341` would have produced a
  correctly-cited false statement.
- **A commit message is testimony, not evidence.** `d11a403`'s own message says roadrunner's
  `repos.yml` entry *"was already gone"*; `git show d11a403 -- repos.yml` shows that very commit
  deleting all three entries. **[iter-83]** The irony is now double-edged: this ruling is stated inside
  a commit message that is itself the sole record of a gate-critical work item, and that contains at
  least one measurably wrong figure (`D-M257x-82-3`).

---

## The grade this record cannot leave standing **[iter-83]**

Reconstructing the record made the iteration auditable for the first time, and the audit changes the
grade. Measured by `repair_reach_guard.py` — built at iter-83 for exactly this question — against this
commit and the ledger it was given (`iter-76/raw/`, 152 booked findings):

| | |
|---|---|
| booked findings carrying a gradeable corpus anchor | **147** |
| landed inside a repair hunk | **109 (74.1 %)** |
| **never reached** | **38** |
| of those, in files the repair **opened and edited** | **35** |
| in files never opened | 3 |

**`closed-fixed` is not supportable.** The iteration's planned scope was *"repair the 11 predicates"*
and 38 of its own input findings were not written. The honest status is **`closed-fixed-partial`**, and
it is recorded here rather than left as the message's own grade.

**This does not diminish what landed.** 109 findings repaired across 33 files, six guards green, TRAP A
held under pressure, and the P9 ref-relative ruling that moved `storage` off `mid-fold` are all real
and all survived iter-82's re-read. The defect is not the repair; it is that **the repair reported a
completeness it never measured** — and the mechanism is iter-83's subject, not a footnote here.

## What iter-83 measured about *how* it happened **[iter-83]**

Three findings, in increasing order of depth. Full argument in
[`../iter-83/progress.md`](../iter-83/progress.md).

1. **No per-anchor post-condition existed.** The discharge criterion was a per-file judgment sweep.
   In `graphql-wundergraph.md` the repair rewrote `:177` — a finding the adjudication had **REJECTED** —
   and left `:13`, which **both** readings booked as their **B1** and which was **UPHELD**.
2. **`repair_leak_guard.py` was not run, and it goes RED on this commit** (rc=1, 3 sites: `CLAUDE.md:285`,
   `platform-alignment.md:1249`, `messenger.md:122` — all three still standing at HEAD). It is the one
   guard in the family whose stated question is *"did this commit FINISH?"*, and it is absent from the
   six the message lists.
3. **It was absent because the guard list was hand-maintained.** `repair_leak_guard` declares
   `FENCE_KIND = "standalone"`, so the DERIVED registry in `repair_postcondition.py` — which selects
   only `postcondition`-kind fences — never reaches it. A repair choosing its guards by hand is §2 of
   this milestone's own protocol doc: **the hand-maintained tuple nobody updates.**

## Routed forward (recorded at iter-83, not by iter-81)

- **`FIX-M257x-iter83-eleven-discharges-unproven`** — all eleven verdicts re-derived as membership
  questions. → iter-84
- **`FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger`** — iter-76 recorded rejection
  *mechanisms* and counts, never per-anchor verdicts, so the 12 rejections cannot be subtracted from
  any reach measurement. → iter-84
- **`CHECK-M257x-iter83-standalone-is-the-forgettable-class`** — 9 of 14 guards are `standalone`;
  nothing runs them by derivation. → iter-85
