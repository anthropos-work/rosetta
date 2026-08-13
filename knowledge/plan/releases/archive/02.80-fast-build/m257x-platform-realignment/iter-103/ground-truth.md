# iter-103 ground truth — re-derived at this reading's open, 2026-08-06

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
TOK-04 P1/P2 discipline. Where it agrees with iter-101's or iter-102's sheet that is a re-derivation, not
a copy.

**This is a MEASURING pass. No repair is taken inside it.** The separation between repairing and measuring
is the only reason any of this milestone's numbers mean anything.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD | `e6aed2e822648e6a2fb501967aa718b5d6d63c9d` (`iter(M257x/102)`) |
| branch | `m257x/platform-realignment` |
| tree at the open | clean but for Lane B's untracked `gate-clauses-1-2/raw/` (11 files, not ours, left alone) |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **40 files, 10,646 lines**, 7 seats, greedy longest-processing-time balance (**1495–1552** lines/seat, spread 57) |

## THE PARTITION IS NOT iter-101's — and that is a disclosed discontinuity, not a smoothing

The corpus moved **10,278 → 10,646 lines**, **+368**, from iter-102's 98-site repair. That is **+3.6 %**,
and it is far more than the `+2` that let iter-101 be a replicate on a fixed subject. A greedy LPT balance
over changed file sizes yields a **different file-to-seat assignment**.

**The partitioning ALGORITHM is unchanged and that was proven, not asserted.** The same script, run over
the file sizes at `8f04d3a`, reproduces iter-101's published partition **exactly** — all seven seats, the
same files, the same 1431–1506 line spread. So the instrument is the same; the subject moved under it.

**The binding consequence, stated once and honoured everywhere below:**

> **iter-103 is NOT a replicate of iter-101.** It is a fresh reading of a *repaired and grown* corpus on a
> *recomputed* partition. Seat-level numbers are not comparable across the two readings; only
> reading-level numbers (`N`, upheld rate, per-pass recall, spread) are, and even those carry the changed
> subject.

### This reading's partition

| seat | lines | files |
|---|---|---|
| A | 1519 | `external_services.md` · `security_compliance.md` · `roadrunner.md` · `dependency_map.md` · `TEMPLATE.md` |
| B | 1541 | `ai-readiness.md` · `cms.md` · `chronos.md` · `ai-labs.md` · `services/README.md` · `skiller.md` |
| C | 1511 | `alignment_testing.md` · `platform-migration-status.md` · `jobsimulation.md` · `customerio-sync.md` · `askengine.md` · `intelligence.md` |
| D | 1515 | `service_taxonomy.md` · `clerkenstein.md` · `graphql-wundergraph.md` · `messenger.md` · `frontend_architecture.md` · `db-backup.md` |
| E | 1513 | `studio-room.md` · `backend.md` · `ai_architecture.md` · `next-web-app.md` · `clerk-integration.md` · `architecture/README.md` |
| F | 1495 | `ant-academy.md` · `architecture_overview.md` · `storage.md` · `sentinel.md` · `academy-backend.md` |
| G | 1552 | `hiring.md` · `studio-desk.md` · `shared_libraries.md` · `coursebuilder.md` · `skillpath.md` · `gotenberg.md` |

## What iter-102 changed, and what it did not

iter-102 paid **both** outstanding read unions in one pass — `FIX-M257x-iter99-read-union` (28 anchors) and
`FIX-M257x-iter101-read-union` (24 anchors) — **52 anchors → 76 assignments → 98 sites found, 94 repaired**,
over a 10-seat fan-out. Repair scope was `corpus/**` + `CLAUDE.md`, **wider than clause 5's read scope**,
deliberately.

**Machine-graded reach** (`repair_reach_guard` against each reading's own raw seat reports): **37/46 = 80.4 %**
against iter-99's ledger, **29/36 = 80.6 %** against iter-101's — and the entire unreached residue is the
adjudicator-**REJECTED** findings, which must not be repaired because a rejection is a claim that turned out
true. Against the **UPHELD** set the reach is effectively 100 %.

**So this reading's central question is one no prior reading could ask:** *both* unions have now been paid.
If repair reaches the residual, `N` must fall materially. If it does not, that is the finding.

## Platform clones — the only thing that settles a claim

Read at this open with **`git rev-parse`, no fetch** (§5 rule 41a — a reading's ground truth includes the
clone refs, and no lane may fetch while a reading is in flight).

| repo | checkout HEAD | `origin/main` | last fetched | iter-101 sheet |
|---|---|---|---|---|
| **platform** | `0c91421d` | `0c91421d` — **in sync** | 12:15 | `0c91421d` |
| **app** | **`ad9f3c49`** | `ad9f3c49` — **in sync** | 12:15 | `b948604f` |
| app/studio (nested, own checkout) | `aeec036a` | — | — | `aeec036a` |
| cms/studio (nested, own checkout) | `aeec036a` | — | — | `aeec036a` |
| **next-web-app** | **`8297c684`** | `f97ba659` | 12:15 | `bb3313bc` |
| **sentinel** | **`f2c46190`** | `f2c46190` — in sync | 12:15 | `88bc5592` |
| **studio-desk** | **`41ee3575`** | `41ee3575` — in sync | 12:15 | `14a5442a` |
| **ant-academy** | **`22df69dd`** | `22df69dd` — in sync | 11:18 | `9c3843cd` |
| cms | `ca50c817` | `f38c0c4a` | 2026-08-05 23:24 | `ca50c817` |
| jobsimulation | `462343b0` | `82cb66ec` | 2026-08-05 23:24 | `462343b0` |
| messenger | `fa47850d` | `e9421c68` | 2026-08-05 23:24 | `fa47850d` |
| storage | `4ce8ece5` | `9f8cb532` | 2026-08-05 23:24 | `4ce8ece5` |
| roadrunner | `87d8d443` | `87d8d443` — in sync | 2026-08-05 23:24 | `87d8d443` |
| graphql-wundergraph | `60c229f3` | `60c229f3` — in sync | 2026-08-05 23:24 | `60c229f3` |
| rosetta-extensions (**per-stack, pinned**) | **`09d06070`** | `4cb920aa` | 11:19 | `ab81527a` |
| rosetta-extensions (**authoring**) | **`944fc4a2`** on `main` | — | — | `09d06070` |

**Fetch times are recorded on BOTH sides**, per §5 rule 41a's first corollary. The most recent fetch
(12:15, five clones) landed **10 minutes before this sheet** and **before any seat was dealt**. The
ground-truth sheet is written first, the pre-registration is sealed in its own commit, and only then are
seats dealt — so a mid-reading move is detectable rather than merely suspected. **The fetch times are
re-read at the close and published there.**

### The known instrument defect, stated but NOT fixed — third reading running

`briefing-iter76-AS-RUN.md:37` names `.agentspace/rosetta-extensions` — the **authoring** copy — as "the
tooling". A rext claim about what a stack runs is settled in the **pinned per-stack clone**.

It is delivered **unchanged**, for the third reading in a row. Editing it would break the comparability the
series exists to establish. **What the ADDENDUM does — and this is the only change from iter-101's
delivery — is add a section that names which of the two trees settles a claim**, per §5 rule 45, which was
written after iter-100 precisely because a briefing that names the wrong tree manufactures false bookings by
construction. The frozen text above the line is untouched; the addendum supersedes ground truth in the open,
which is what the addendum is for. **Band #6 measures whether that helped.**

Routed as `DEF-M257x-iter101-briefing-rext-tree`, still open.

## The instrument — untouched, and proven so

| | value |
|---|---|
| file | `instrument/briefing-iter76-AS-RUN.md` |
| sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| re-checked | **AFTER** copying to `iter-103/briefing-AS-DELIVERED.md`, not before |
| `git log --follow` on the FILE | **exactly one commit ever** — `012edd2` (iter-76) |

## Guard family at the open

`14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17 members, corpus `e6aed2e82`, platform
`0c91421df` (origin/main in sync, fetched 9 m ago). **Identical to iter-101's and iter-102's opens.**
The 3 need `--range`/`--ledger`, which a tree-state run cannot supply — recorded as a gap rather than
hidden, and `guard_family` exits 2 to say so.

### `anchor_construct_guard` reach — measured at BOTH corpus refs with the SAME binary

| corpus | resolved | unresolvable | reach |
|---|---|---|---|
| `8f04d3a` (iter-101's subject) | 526 | 338 | **60.9 %** of 864 |
| `e6aed2e` (this reading's subject) | **645** | 433 | **59.8 %** of 1078 |

**0 findings at both.** iter-102's repair added **+214 anchors** and **+119 resolved** — the reach ratio is
flat, so the repair did not dilute the fence. ⚠ **This is not the same denominator as iter-101's sheet,
which quoted *"528 of 555 citations"*.** That figure is not reproducible from the current flag set and is
**not** reconciled here; the two numbers above were both taken today, with one binary, and are comparable
to each other. Naming the discrepancy is the honest move; smoothing it would be exactly the class this
milestone exists to catch.

## Reading shape

7 seats × 2 independent readings (#25, #26) of the **identical** partition = 14 blind seats. No seat knows
which reading it is in; no seat may read `knowledge/plan/**` beyond its own briefing and output, so no seat
can see a prior audit's answer key or another seat's report.

**Dealt in two batches of 7** — a deliberate change of *operations*, not of instrument. iter-101 dealt all
14 at once and lost `r24-D` to a spend limit, making that reading a 13-seat union. Batching does not touch
the briefing, the partition, the grading rule or the scope.
