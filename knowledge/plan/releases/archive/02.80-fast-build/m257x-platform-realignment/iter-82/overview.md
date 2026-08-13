---
iter: 82
milestone: M257x
type: tik
status: closed-measured
opened: 2026-08-05
tok: TOK-05
iter_shape: measurement
---

# iter-82 — the RE-READ (readings #15 / #16) after the iter-81 repair

**Planned deliverable:** re-run clause 5's instrument against the repaired corpus and report the
number, whatever it is. **No repair lands in this iter.** A repair validated by the same pass that
found the work is the *"reports a state without measuring it"* defect this milestone exists to end;
if N > 0, the composition is booked and the repair is a later iter's job.

## The instrument — held fixed, and PROVEN so rather than asserted

§5 rule 25: an instrument that is *described* rather than *stored* is not frozen. The briefing is
stored, and this iter adds a second proof — the partition **method** is re-executed and shown to
reproduce the prior reading's hand exactly.

| knob | value | how it was verified |
|---|---|---|
| briefing | `instrument/briefing-iter76-AS-RUN.md`, **byte-identical** | sha256 `3858ec53…`; `git log` on the path shows **one** commit (`012edd2`, iter-76) — untouched since it was taken |
| seat count | 7 per reading (A–F full-read + G adversarial diff), **2 readings = 14 seats** | unchanged |
| file set | **the same 40 files** — `corpus/architecture/*.md` + `corpus/services/*.md` | enumerated at HEAD; cardinality 40, identical membership |
| partition method | files sorted by line count **descending**, snake-dealt A→F then F→A | **re-executed at `012edd2` and it reproduced iter-76's hand EXACTLY** — 40 files / 9,544 lines, seat totals 1791 / 1621 / 1537 / 1495 / 1542 / 1558 and per-seat file lists all identical to the recorded partition |
| ground truth | 12 platform clones at the briefing's exact shas | **re-derived at this reading's open: all 12 match the briefing table byte-for-byte.** `app b948604f` · `platform 0dab54df` · `next-web-app bb3313bc` · `sentinel 88bc5592` · `storage 4ce8ece5` · `messenger fa47850d` · `cms ca50c817` · `graphql-wundergraph 60c229f3` · `roadrunner 87d8d443` · `jobsimulation 462343b0` · `studio-desk 14a5442a` · `ant-academy 9c3843cd` |

**The one thing that MOVED, and it moved because the instrument is fixed:** iter-81's repair changed
line counts (+410 / −251 over 33 files), so the fixed method **deals a different hand**. The corpus
is now **40 files / 9,712 lines** (was 9,544). This is the same consequence iter-76 recorded when
its hand differed from #11/#12's, and it is recorded rather than engineered away: freezing the
*output* of the method instead of the *method* would be the drift, not the fix for it.

### The hand, as dealt at HEAD `328ece5`

| seat | lines | files |
|---|---|---|
| **A** | 1801 | external_services · security_compliance · jobsimulation · academy-backend · messenger · skiller · TEMPLATE |
| **B** | 1634 | ai-readiness · ai_architecture · graphql-wundergraph · ai-labs · coursebuilder · customerio-sync · architecture/README |
| **C** | 1576 | alignment_testing · clerkenstein · backend · sentinel · next-web-app · services/README · db-backup |
| **D** | 1552 | studio-room · architecture_overview · cms · roadrunner · clerk-integration · gotenberg · intelligence |
| **E** | 1575 | service_taxonomy · hiring · platform-migration-status · storage · askengine · dependency_map |
| **F** | 1574 | ant-academy · studio-desk · chronos · shared_libraries · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of the iter-81 repair — `f375738..HEAD -- corpus/ CLAUDE.md .claude/`, **33 files / +410 −251** |

Readings #15 and #16 get the **identical** hand. Any disagreement between them is seat variance and
nothing else (§5 rule 23's control).

## PRE-REGISTERED PREDICTIONS — written before any seat returned, before any report was read

The prior read (#13/#14) returned **77 / 75**, adjudicated to **140 upheld / 12 rejected**, deduped
to **11 predicates**, all 11 repaired by iter-81. The naive expectation is therefore zero. It is
**not** my prediction, for one reason that is already measured and is the most load-bearing fact
going in:

> **Recall was below 60%.** #13 and #14 read the identical partition and each booked blockers the
> other missed. Repairing the **union of two readings** cannot repair what **neither** reading saw.
> A fresh pair draws a fresh sample from the same unseen remainder — so a residual is expected on
> statistics alone, entirely independently of whether iter-81's repair was any good.

1. **Neither reading returns zero.** `N₁₅ > 0` **and** `N₁₆ > 0`. *Falsified if either is 0 — which
   is the outcome the gate needs, so this prediction is written against my own interest.*
2. **Per-reading count.** Each of `N₁₅`, `N₁₆` lands in **[10, 45]** — well below the 77/75 of the
   pre-repair read (the 11 predicates were ~140 findings and are gone), well above zero (the recall
   argument above).
3. **Union.** `|#15 ∪ #16| > max(N₁₅, N₁₆)` — recall is still not 100%.
4. **The repair introduced defects.** Seat G books **≥ 1** blocker inside the iter-81 diff. A
   33-file repair that introduces nothing has never happened in this milestone.
5. **The 11 repaired predicates contribute ZERO blockers.** If any seat books a finding inside a
   predicate iter-81 claimed to repair, **the repair was incomplete** and that is this iter's
   headline, not a footnote.
6. **The held carve-out is re-booked.** At least one seat books `storage.md` `:55` / `:154` / `:181`
   (the `/tmp`-sandbox-vs-production-bucket contradiction). It is **held by instruction**
   (`DEF-M257x-iter80-storage-prod-bucket`, escalated, awaiting the user's decision) — so if it
   appears it is **not** a repair failure and must be reported as held, not repaired.

## Escalation condition, written at open

If `N₁₅ ∪ N₁₆` exceeds **60**, the repair-then-re-read loop is not converging and the next move is a
**TOK**, not a twelfth predicate list.
