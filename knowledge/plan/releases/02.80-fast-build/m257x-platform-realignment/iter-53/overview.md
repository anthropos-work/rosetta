---
iter: 53
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-03
closed: 2026-08-03
---

# iter-53 — the PAIRED READING (#11 + #12): two blind full passes, no repair

**Active strategy:** [`TOK-03: repair the UNION, shrink the estimator, make the edits smaller`](../decisions.md#tok-03-repair-the-union-shrink-the-estimator-make-the-edits-smaller--2026-08-03)
— this iteration executes **move 1** (*repair the union of two independent readings, never one*) in its
measurement half, and produces the input for **move 2** (*drive `N̂` down first*). It is the tik
TOK-03's own **Next-tik direction** names: *"iter-53 = the paired reading #11 + #12, blind, frozen
instrument, and recompute `N̂`."*

**This is a MEASUREMENT, not a gate attempt.** `N̂ ≈ 14` (floor) at open. §5 rule 23's arithmetic says a
single pass missing an entire residual of that size has probability ≈ 10⁻⁵ — so no reading taken here can
meet clause 5, and none is being represented as trying to. The deliverable is the union, the overlap, the
per-reading recall, and a re-derived `N̂`.

## Step 0 — re-survey before targeting (mandatory, done before this plan settled)

- **Platform origin re-fetched at open: `2adcf714`, unchanged.** Re-scope trigger stays at **occurrence 1
  of 2**.
- **The corpus HAS changed since reading #10** — iter-52 repaired the union of 18 (`0586167`). So this is
  *not* a same-tree replication of iter-50; it is a fresh paired reading of the **post-repair** tree. The
  partition method is held fixed and therefore deals a **different hand** (file sizes moved), which is the
  honest consequence of a fixed method over a changed corpus and is recorded rather than engineered away.
- **Ground-truth clones byte-identical to iter-50's**, re-read at open: `app 5ba17044` ·
  `app/studio aeec036a` · `platform 2adcf714` · `next-web-app bb3313bc` · `sentinel 88bc5592` ·
  `storage 4ce8ece5` · `messenger fa47850d` · `cms ca50c817` · `graphql-wundergraph 60c229f3` ·
  `roadrunner 87d8d443` · `jobsimulation 462343b0` · `studio-desk 14a5442a` · `ant-academy 9c3843cd`.
- **Scope unchanged: 40 files, 9,395 lines** (was 9,326 at #9/#10 — the repair added 69 net lines).

## Instrument — HELD FIXED at iter-41's, on every knob

- **Seven auditors per reading** — six full-read partitions (A–F) + one adversarial diff seat (G).
- **Same briefing** — every §5 rule that bears on reading, the blocker/minor grading rule, and §5 rule 24's
  *enumerate the SET, not the sum* clause promoted to a standing instruction after three seats failed it.
- **Same partition METHOD** — files sorted by line count descending, snake-dealt A→F then F→A.
- **All 40 files read in full, top-to-bottom**, under a per-file `wc -l` positive control.
- **Seat G reads iter-52's own repair diff** — `1255998..0e35b1a -- corpus/ CLAUDE.md .claude/`, 11 files.
- **Fresh seats, blind.** No seat is told what any prior reading found, that a prior reading exists, or
  that a *parallel* reading is running. Every seat is barred from `knowledge/plan/**` — which holds the
  answer keys — and from the other reading's raw dir.

### Partition (computed, 40 files / 9,395 lines)

| auditor | lines | files |
|---|---|---|
| **A** | 1775 | external_services · security_compliance · backend · academy-backend · coursebuilder · skiller · TEMPLATE |
| **B** | 1609 | ai-readiness · ai_architecture · graphql-wundergraph · ai-labs · messenger · customerio-sync · architecture/README |
| **C** | 1532 | alignment_testing · architecture_overview · cms · sentinel · clerk-integration · services/README · db-backup |
| **D** | 1481 | studio-room · clerkenstein · chronos · roadrunner · next-web-app · gotenberg · intelligence |
| **E** | 1497 | service_taxonomy · hiring · shared_libraries · storage · askengine · dependency_map |
| **F** | 1501 | ant-academy · studio-desk · jobsimulation · platform-migration-status · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of iter-52's repair, 11 files |

**Readings #11 and #12 get the IDENTICAL partition** — same hand, same diff. Any disagreement between them
is seat variance and nothing else (§5 rule 23's control).

## PRE-REGISTERED PREDICTIONS — written before any seat is launched, before any report is read

1. **Per-reading count.** Each of `N₁₁`, `N₁₂` lands in **[5, 15]**.
2. **Overlap.** Recall is **below 60% for both readings** (i.e. `m/N₁₁ < 0.6` and `m/N₁₂ < 0.6`).
3. **Union.** `|#11 ∪ #12| > max(N₁₁, N₁₂)` — each reading books at least one blocker the other misses.
4. **`N̂` (Chapman).** Lands in **[6, 20]**.
5. **Induced term** — findings attributable to iter-52's repair diff — is **below 8** (iter-52 measured ~8).

### TOK-03's own pre-registrations, carried here to be adjudicated unsoftened

TOK-03 pre-registered, for **this** iteration: **`N̂` below 12**, and **the induced term below 4**.
iter-52 recorded both as *heading toward refutation* and did not soften them. This iteration reports the
verdict on each as measured, whichever way it falls.

## Phase plan

| step | work | done when |
|---|---|---|
| 1 | Re-survey; recompute the partition from current line counts | done at Step 0 |
| 2 | Launch reading **#11** — 7 blind seats at the frozen instrument | 7 raw reports in `iter-53/raw/r11-*.md` |
| 3 | Launch reading **#12** — 7 blind seats, identical hand, blind to #11 | 7 raw reports in `iter-53/raw/r12-*.md` |
| 4 | Adjudicate each reading to a blocker ledger; compute overlap, recall, union, Chapman `N̂` | `blocker-ledger.md` + `variance.md` |
| 5 | Capture the union as the perishable fixture for iter-54's repair | `fixture-union.md` |

## Escalation conditions

- A seat whose positive control fails → re-run that seat; a partial pass is not a reading (§5 rule 8).
- A platform commit landing mid-iteration → `EXIT_REASON: re-scope-trigger` (occurrence 2 of 2).
- **No repair happens in this iteration under any circumstance.** A repair destroys the measurement and
  spends the fixture (§5 rule 21's perishability clause).

## Acceptable close-no-lift outcomes

**This iteration repairs nothing and cannot move the blocker count.** Pre-declared so the close cannot be
graded as a shortfall: the deliverable is the paired measurement. `closed-fixed` if both readings are
taken and adjudicated and `N̂` is re-derived; `closed-no-lift` only if a reading cannot be completed.
