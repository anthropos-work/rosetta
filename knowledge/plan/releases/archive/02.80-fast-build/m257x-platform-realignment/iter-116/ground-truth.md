# iter-116 ground truth — re-derived at this reading's open, 2026-08-07

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
TOK-04 P1/P2 discipline. Where it agrees with iter-109's sheet that is a re-derivation, not a copy.

**This is a MEASURING pass. No repair is taken inside it.** The separation between repairing and measuring
is the only reason any of this milestone's numbers mean anything. That includes the false-RED disclosed
below: it is graded, routed, and **not fixed here**.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD at the open | `f581de09a8bad7f71f1d97e572554cb2a1dd997c` (`chore(M257x): state.md → iter-115`) |
| branch | `m257x/platform-realignment` |
| tree at the open | **clean** — `git status --short` empty |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **40 files, 10,871 lines**, 7 seats, greedy longest-processing-time balance (**1533–1576** lines/seat, spread 43) |

## THE PARTITION IS NOT iter-109's — disclosed, for the third consecutive reading

The corpus moved **10,694 → 10,871 lines**, **+177**, from iter-115's 24-predicate / 71-site repair. That
is **+1.66 %** — between iter-102's +3.6 % and iter-108's +0.45 %. As at both prior readings it yields a
**different file-to-seat assignment**, because greedy LPT is chaotic under perturbation: re-ordering the
descending sort cascades every subsequent placement.

**The partitioning ALGORITHM is unchanged and that was proven, not asserted.** The same script
(`.agentspace/scratch/work-m257x/partition109.py`), run over the file sizes at `ac48e5b`, reproduces
iter-109's published partition **exactly** — all seven seats, the same files, the same 1506–1557 loads,
the same 51-line spread. So the instrument is the same; the subject moved under it.

> **iter-116 is NOT a seat-level replicate of iter-109.** It is a fresh reading of a *repaired* corpus on a
> *recomputed* partition. Seat-level numbers are not comparable across the two readings; only
> reading-level numbers (`N`, `P`, upheld rate, per-pass recall, spread) are.

### This reading's partition

| seat | lines | files |
|---|---|---|
| A | 1533 | `external_services.md` · `jobsimulation.md` · `customerio-sync.md` · `askengine.md` · `intelligence.md` |
| B | 1543 | `ai-readiness.md` · `cms.md` · `chronos.md` · `academy-backend.md` · `services/README.md` |
| C | 1574 | `alignment_testing.md` · `ai_architecture.md` · `platform-migration-status.md` · `coursebuilder.md` · `skillpath.md` · `gotenberg.md` |
| D | 1545 | `service_taxonomy.md` · `clerkenstein.md` · `security_compliance.md` · `sentinel.md` · `ai-labs.md` · `db-backup.md` |
| E | 1549 | `studio-room.md` · `backend.md` · `graphql-wundergraph.md` · `messenger.md` · `frontend_architecture.md` · `architecture/README.md` |
| F | 1551 | `ant-academy.md` · `architecture_overview.md` · `storage.md` · `roadrunner.md` · `dependency_map.md` · `TEMPLATE.md` |
| G | 1576 | `hiring.md` · `studio-desk.md` · `shared_libraries.md` · `next-web-app.md` · `clerk-integration.md` · `skiller.md` |

## Platform clones — read with `git rev-parse`, NO fetch (§5 rule 41a)

| repo | checkout HEAD | `origin/main` | last fetched | iter-109 sheet | moved? |
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
| rosetta-extensions (**authoring**) | **`1dc1eb82`** on `main` | `1dc1eb82` | 2026-08-05 23:24 | `680e8529` | yes — ours |

> ### **ZERO of the 14 platform clones have advanced across THREE consecutive readings** (iter-103,
> iter-109, iter-116).
>
> **And iter-109 already established what that does and does not buy.** It removes the *arrival* of new
> drift — and it did, and platform-drift was still **~33 %** of that reading's upheld residual. The
> ARRIVAL-vs-DETECTION question is **closed: DETECTION** (`D-M257x-109-3`). So a frozen subject is not
> a reason to expect fewer drift findings here; it is the reason drift findings here are known to be
> **standing** rather than newly-arrived.

**Fetch times are recorded on BOTH sides.** The most recent platform-side fetch (2026-08-06 12:15)
predates this sheet by half a day and predates any seat being dealt. **The fetch times are re-read at the
close and published there**, so a mid-reading move is detectable rather than merely suspected.

### The known instrument defect, stated but NOT fixed — fifth reading running

`briefing-iter76-AS-RUN.md:37` names `.agentspace/rosetta-extensions` — the **authoring** copy — as "the
tooling". A rext claim about what a stack runs is settled in the **pinned per-stack clone**.

It is delivered **unchanged**, for the fifth reading in a row. Editing it would break the comparability the
series exists to establish. The ADDENDUM names which of the two trees settles a claim (§5 rule 45), plus
the refinement that a claim about a **fence's own verdict or configuration** is settled by the *authoring*
tree (`D-M257x-103-7`). Routed as `DEF-M257x-iter101-briefing-rext-tree`, still open, still
delivered-unfixed. Measured series of the rejection class it produces: **4 → 1 → 1 → 0**.

## The instrument — untouched, and proven so

| | value |
|---|---|
| file | `instrument/briefing-iter76-AS-RUN.md` |
| sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| re-checked | **AFTER** copying to `iter-116/briefing-AS-DELIVERED.md`, not before — both sides printed, and `diff` returned empty |
| `git log --follow` on the FILE | **exactly one commit ever** — `012edd2` (iter-76) |

## Guard family at the open — 14 GREEN · **1 RED** · 4 not-run, and the RED is a FALSE one

```
fence-tree: /Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions @ 1dc1eb824
            describe=fast-build-m257x-iter-101-15-g1dc1eb8
guard-family: corpus /Users/marco/workspace/anthropos/rosetta @ f581de09a
guard-family: 19 member(s) on disk, all placed. platform=stack-demo/platform
guard-family: 14 GREEN · 1 RED · 0 could-not-check · 4 not-run
guard-family: RED — platform_predicate_guard
```

The 4 not-run need `--range`/`--ledger`, which a tree-state run cannot supply — recorded as a gap rather
than hidden (`anchor_offset_guard`, `repair_leak_guard`, `repair_reach_guard`, `value_change_guard`).

### The RED, characterised before it is dismissed

```
[G10 service-count] corpus/services/sentinel.md:5 says compose declares 5 service(s);
at d11a403 `docker-compose.yml` declares 8 and the effective topology (with `include:`) has 10
— 5 is neither
```

**Measured, by opening the source at both refs with the guard's own helper:**

| ref | `(file-local, effective-with-include)` |
|---|---|
| `0c91421d` — **the ref the sentence itself names** | **(5, 7)** |
| `d11a403` — the ref the guard graded it at | (8, 10) |

The corpus sentence reads *"**At platform `0c91421d`**, `docker-compose.yml` declares **five** services —
`sentinel` (`:5`), `backend` (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`), `gotenberg`
(`:170`)"*, and all five line anchors resolve. **The claim is TRUE at the ref it names.**

**The defect is in the guard.** `sentinel.md:5` is a single long wrapped paragraph carrying **two**
platform refs: `d11a403` (dating a *different* proposition — *"platform `d11a403` deleted both compose
services along with `roadrunner`"*) and `0c91421d` (dating this one). G10 builds its window with
`_pin_window(...)` and then takes `_REF_PINNED.search(cell)` — a **`search`, which returns the FIRST match
in the window**. On a multi-pin block it therefore dates the claim by whichever ref happens to appear
earliest, not by the ref the claim names.

That is §5 rule 33 — *a claim is settled at the ref the claim itself names* — and the briefing's own
sentence *"a pin's scope is the claim's own block … a ref named in a neighbouring row does not date this
row's claim"* — **violated by the guard that enforces it.**

> **Routed as `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block` — NOT repaired here.** No
> repair is taken inside a measuring pass. It is disclosed in the seats' addendum too, so a seat that
> notices the same site books what it measured rather than what a guard said.

**Why this does not block the reading, stated rather than assumed.** This iter lands **zero** code and
**zero** corpus edits — it is a measurement. A RED gate blocks an iter because code must not land on top
of one; there is nothing to land. The RED's subject was opened at source and the corpus claim it names
holds. It is recorded as an open FIX against the tooling, in the same tree as
`FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live` — **the second guard in two iters caught
resolving a claim against the wrong thing**, which is a pattern worth naming even though neither is a
corpus defect.

## Reading shape

7 seats × 2 independent readings (**#29, #30**) of the **identical** partition = 14 blind seats. No seat
knows which reading it is in; no seat may read `knowledge/plan/**` beyond its own briefing and its own
output file, so no seat can see a prior audit's answer key or another seat's report.

**Dealt in two batches of 7**, as at iter-103 and iter-109. **Every seat is committed verbatim the moment
it lands**, before adjudication — the discipline that has now bounded the cost of a mid-flight death three
times.

## Test-suite state — a GAP, recorded, not a pass

This is a measuring pass: **no rext code changed, so no suite was run for it.** The standing figure for
`stack-core` remains iter-111's: **`1 failed · 1011 passed` in 1090.88 s**, and the one failure is the
perishable iter-48 answer-key fixture. **That count is quoted with its invocation** (§5 rule 50) and is
**not** re-run here, because nothing in this iter could change it.
