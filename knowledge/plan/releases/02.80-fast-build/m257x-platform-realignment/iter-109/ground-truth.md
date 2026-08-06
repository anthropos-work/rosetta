# iter-109 ground truth — re-derived at this reading's open, 2026-08-06

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
TOK-04 P1/P2 discipline. Where it agrees with iter-103's sheet that is a re-derivation, not a copy.

**This is a MEASURING pass. No repair is taken inside it.** The separation between repairing and measuring
is the only reason any of this milestone's numbers mean anything.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD | `2e3443d316ab1e50e698edf1186bd094c79dc336` (`chore(M257x): advance state.md to iter-108`) |
| branch | `m257x/platform-realignment` |
| tree at the open | **clean** — `git status --short` empty |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **40 files, 10,694 lines**, 7 seats, greedy longest-processing-time balance (**1506–1557** lines/seat, spread 51) |

## THE PARTITION IS NOT iter-103's — disclosed, again, and for a smaller reason than last time

The corpus moved **10,646 → 10,694 lines**, **+48**, from iter-108's 22-predicate repair. That is **+0.45 %**
— an order of magnitude less churn than iter-102's +3.6 %. **It still yields a different file-to-seat
assignment**, because greedy LPT is chaotic under perturbation: a 48-line shift is enough to re-order the
descending sort and cascade every subsequent placement. Small input change ≠ small output change.

**The partitioning ALGORITHM is unchanged and that was proven, not asserted.** The same script
(`.agentspace/scratch/work-m257x/partition109.py`), run over the file sizes at `e6aed2e`, reproduces
iter-103's published partition **exactly** — all seven seats, the same files, the same 1495–1552 spread,
the same 57-line spread. So the instrument is the same; the subject moved under it.

> **iter-109 is NOT a seat-level replicate of iter-103.** It is a fresh reading of a *repaired* corpus on a
> *recomputed* partition. Seat-level numbers are not comparable across the two readings; only
> reading-level numbers (`N`, the predicate count `P`, upheld rate, per-pass recall, spread) are.

### This reading's partition

| seat | lines | files |
|---|---|---|
| A | 1520 | `external_services.md` · `ai_architecture.md` · `coursebuilder.md` · `clerk-integration.md` · `architecture/README.md` |
| B | 1524 | `ai-readiness.md` · `cms.md` · `chronos.md` · `academy-backend.md` · `services/README.md` · `TEMPLATE.md` |
| C | 1514 | `alignment_testing.md` · `platform-migration-status.md` · `jobsimulation.md` · `customerio-sync.md` · `askengine.md` · `intelligence.md` |
| D | 1520 | `service_taxonomy.md` · `clerkenstein.md` · `security_compliance.md` · `messenger.md` · `frontend_architecture.md` · `db-backup.md` |
| E | 1553 | `studio-room.md` · `backend.md` · `graphql-wundergraph.md` · `roadrunner.md` · `dependency_map.md` · `skiller.md` |
| F | 1506 | `ant-academy.md` · `architecture_overview.md` · `storage.md` · `sentinel.md` · `ai-labs.md` |
| G | 1557 | `hiring.md` · `studio-desk.md` · `shared_libraries.md` · `next-web-app.md` · `skillpath.md` · `gotenberg.md` |

## Platform clones — and the single most important fact about this reading

Read at this open with **`git rev-parse`, no fetch** (§5 rule 41a — a reading's ground truth includes the
clone refs, and no lane may fetch while a reading is in flight).

| repo | checkout HEAD | `origin/main` | last fetched | iter-103 sheet | moved? |
|---|---|---|---|---|---|
| **platform** | `0c91421d` | `0c91421d` — in sync | 2026-08-06 12:15 | `0c91421d` | **no** |
| **app** | `ad9f3c49` | `ad9f3c49` — in sync | 2026-08-06 12:15 | `ad9f3c49` | **no** |
| app/studio (nested, own checkout) | `aeec036a` | `aeec036a` | never | `aeec036a` | **no** |
| cms/studio (nested, own checkout) | `aeec036a` | `aeec036a` | never | `aeec036a` | **no** |
| **next-web-app** | `8297c684` | `f97ba659` | 2026-08-06 12:15 | `8297c684` | **no** |
| **sentinel** | `f2c46190` | `f2c46190` — in sync | 2026-08-06 12:15 | `f2c46190` | **no** |
| **studio-desk** | `41ee3575` | `41ee3575` — in sync | 2026-08-06 12:15 | `41ee3575` | **no** |
| **ant-academy** | `22df69dd` | `22df69dd` — in sync | 2026-08-06 11:18 | `22df69dd` | **no** |
| cms | `ca50c817` | `f38c0c4a` | 2026-08-05 23:24 | `ca50c817` | **no** |
| jobsimulation | `462343b0` | `82cb66ec` | 2026-08-05 23:24 | `462343b0` | **no** |
| messenger | `fa47850d` | `e9421c68` | 2026-08-05 23:24 | `fa47850d` | **no** |
| storage | `4ce8ece5` | `9f8cb532` | 2026-08-05 23:24 | `4ce8ece5` | **no** |
| roadrunner | `87d8d443` | `87d8d443` — in sync | 2026-08-05 23:24 | `87d8d443` | **no** |
| graphql-wundergraph | `60c229f3` | `60c229f3` — in sync | 2026-08-05 23:24 | `60c229f3` | **no** |
| rosetta-extensions (**per-stack, pinned**) | `09d06070` | `4cb920aa` | 2026-08-06 11:19 | `09d06070` | **no** |
| rosetta-extensions (**authoring**) | **`680e8529`** on `main` | `680e8529` | 2026-08-05 23:24 | `944fc4a2` | yes — ours |

> ### **ZERO of the 14 platform clones advanced between iter-103's reading and this one.**
>
> This is the fact the whole reading turns on, and it is not luck — §5 rule 41a plus the deliberate absence
> of any bring-up in iters 104–108 kept the subject still. **The clone-advance inflow — 61 % of iter-103's
> `N` — had no opportunity to re-arm.** The only rext tree that moved is our own authoring copy, which
> carries this milestone's fences and is not a subject of the read.
>
> **What that buys, and what it does NOT.** It removes the *arrival* of new drift. It does **not** make an
> old drift defect true, and it does not restore anything iter-103's passes failed to detect: measured
> per-pass recall on this instrument has run **33–83 %**, so an undetected standing residual is fully
> compatible with every number in the series. That distinction — *arrival* versus *detection* — is the
> question the pre-registered rule is built to separate.

**Fetch times are recorded on BOTH sides**, per §5 rule 41a's first corollary. The most recent platform-side
fetch (12:15) predates this sheet by hours and predates any seat being dealt. **The fetch times are re-read
at the close and published there**, so a mid-reading move is detectable rather than merely suspected.

### The known instrument defect, stated but NOT fixed — fourth reading running

`briefing-iter76-AS-RUN.md:37` names `.agentspace/rosetta-extensions` — the **authoring** copy — as "the
tooling". A rext claim about what a stack runs is settled in the **pinned per-stack clone**.

It is delivered **unchanged**, for the fourth reading in a row. Editing it would break the comparability the
series exists to establish. The ADDENDUM names which of the two trees settles a claim (§5 rule 45), and
this reading adds one refinement below it: a claim about a **fence's own verdict or configuration** is
settled by the *authoring* tree, because a verdict is a measurement taken with that fence's config
(`D-M257x-103-7`). Routed as `DEF-M257x-iter101-briefing-rext-tree`, still open, still delivered-unfixed.

## The instrument — untouched, and proven so

| | value |
|---|---|
| file | `instrument/briefing-iter76-AS-RUN.md` |
| sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| re-checked | **AFTER** copying to `iter-109/briefing-AS-DELIVERED.md`, not before — both sides printed |
| `git log --follow` on the FILE | **exactly one commit ever** — `012edd2` (iter-76) |

## Guard family at the open — and it now states its own provenance

```
fence-tree: /Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions @ 680e8529f
            describe=fast-build-m257x-iter-101-6-g680e852
guard-family: corpus @ 2e3443d31 · platform @ 0c91421df (origin/main in sync)
guard-family: 15 GREEN · 0 RED · 0 could-not-check · 4 not-run   [19 members on disk]
```

**`fence-tree:` is `FIX-M257x-iter103-guard-tree-provenance` — TOK-06 step 0 — working in production.** At
iter-103 the family printed the corpus sha and the platform sha and *not its own*, and the coordinator
nearly published two false conclusions off a verdict taken with last release's fence configuration. The one
input that decides a verdict is now the first line of the output.

The family has grown **17 → 19 members** since iter-103 (`anchor_offset_guard`, `clone_drift_guard`,
`fence_provenance` and `repair_postcondition` landed across iters 105–108). The 4 not-run need
`--range`/`--ledger`, which a tree-state run cannot supply — recorded as a gap rather than hidden, and
`guard_family` exits 2 to say so.

## Reading shape

7 seats × 2 independent readings (#27, #28) of the **identical** partition = 14 blind seats. No seat knows
which reading it is in; no seat may read `knowledge/plan/**` beyond its own briefing and its own output
file, so no seat can see a prior audit's answer key or another seat's report.

**Dealt in two batches of 7**, as at iter-103. **Every seat is committed verbatim the moment it lands**, before
adjudication — the discipline that has now bounded the cost of a mid-flight death three times.

---

## CORRECTION, disclosed rather than silently rewritten — the corpus HEAD moved during this open

**Written after the seal (`ac48e5b`), before adjudication. The sheet above said `rosetta HEAD = 2e3443d`
and `tree at the open: clean`. Both were true when measured at 20:10 and neither was true when the seats
were dealt.** A concurrent lane — the previous run, still closing iter-108 — committed **`08cfbd8`** in this
same working tree between the two. `ac48e5b`'s parent is therefore `08cfbd8`, not `2e3443d`.

**This milestone's own class, landing on this milestone's own apparatus for the sixth time**: an evidence
sentence stated more precisely than what was measured, in the direction that made the sheet cleaner. The
conclusion survives; the sentence did not, so the sentence is corrected here rather than quietly patched.

### What the subject actually did — measured, with a firing control

| ref | `corpus/services` tree | `corpus/architecture` tree |
|---|---|---|
| `2e3443d` (sheet's stated HEAD) | `b8d5df873…` | `35953fb68…` |
| `08cfbd8` (the concurrent commit) | `b8d5df873…` | `35953fb68…` |
| `ac48e5b` (the seal — what the seats read) | `b8d5df873…` | `35953fb68…` |
| *negative control* — `e6aed2e` (iter-103's subject) | `1d9da74ac…` — **differs** | — |

`git diff 2e3443d..ac48e5b -- corpus/` is **empty**. `08cfbd8` is +72/−0 across three files, all under
`knowledge/plan/**` (`iter-108/decisions.md`, `iter-108/progress.md`, milestone `progress.md`) — **zero
corpus files**, and `knowledge/plan/**` is barred to every seat.

> **The read scope is byte-identical at all three refs, and the control proves the comparison discriminates
> rather than passing vacuously.** The reading's subject did not move. **The corpus under audit is
> `ac48e5b`**, and the 40-file / 10,694-line partition is unaffected.

### Seal provenance — asked, and answered from evidence

A coordinator flagged that a journal line reading *"an untracked `iter-109/` … NOT mine"* might mean this
seal adopted another lane's artifacts, which would invalidate it. **It does not, and the line is not mine.**
It was written by the *other* lane, about *my* files, and it names them correctly as belonging to "a
concurrent lane preparing TOK-06 step 4" — which is this iter. The file mtimes reconstruct my own tool-call
order exactly: briefing copied 20:14:45 (the same shell call as my 20:14 heartbeat), `ground-truth.md`
20:15:34, `pre-registration.md` 20:16:47, `overview.md` 20:17:11 — authored in that sequence, by me, before
the 20:18 seal. **Case (a): I wrote them; a second observer correctly saw them as not-its-own; the
attribution inverted on the way through.** The seal stands on its own derivation.

**And it is `D-M257x-103-1` / §5 rule 49 again, exactly**: a disagreement between two observers of a
concurrently-mutated surface is *first* evidence that the two observers saw different surfaces — not
evidence that either lied. Third worked example in four days.

## Test-suite state — a GAP, recorded, not a pass

This is a measuring pass: **no rext code changed, so no suite was run for it.** Separately, the `stack-core`
suite is known **not to complete on this host** — a plain `pytest tests/` blocks indefinitely inside
`test_m220_mutation_battery.py::DevWiringMutationBattery`, reproduced twice and proven pre-existing via a
read-only `git archive` of an earlier rext ref. **No full-suite total for `stack-core` exists on this host,
and none is quoted anywhere in this iter** (§5 rule 50 — state the invocation with every count). Routed as
`FIX-M257x-iter108-stackcore-suite-hangs`; second instance of `FIX-M257x-iter100-suite-stall`'s class, now
localised to a named test.
