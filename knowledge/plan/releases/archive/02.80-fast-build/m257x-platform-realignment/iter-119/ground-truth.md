# iter-119 ground truth — re-derived at this reading's open, 2026-08-07

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
TOK-04 P1/P2 discipline. Where it agrees with iter-116's sheet that is a re-derivation, not a copy.

**This is a MEASURING pass. No repair is taken inside it.** The separation between repairing and measuring
is the only reason any of this milestone's numbers mean anything. That includes the false-RED re-disclosed
below: it is graded, routed, and **not fixed here** — for the second reading running.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD at the open | `194361e473ffa454f2d06fc71f6f73a98a26c835` (`iter(M257x/118): tik — class 2 censused…`) |
| branch | `m257x/platform-realignment` |
| tree at the open | **clean** — `git status --short` empty |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **40 files, 10,871 lines**, 7 seats, greedy longest-processing-time balance (**1533–1576** lines/seat, spread 43) |

## THE PARTITION IS iter-116's — IDENTICALLY, and that is a first

For three consecutive readings this sheet has had to disclose that the partition was *recomputed* and
therefore that seat-level numbers were not comparable. **This reading is the exception, and it is the most
useful structural property on the sheet.**

The in-scope corpus moved **10,871 → 10,871 lines** — **zero net lines**, **5 lines changed in place**, in
**2 files**, all of them `#anchor` fragment repairs from iter-117's citation census:

```
git diff --stat 194361e..f581de09 -- corpus/services corpus/architecture
 corpus/architecture/architecture_overview.md     | 2 +-
 corpus/architecture/platform-migration-status.md | 8 ++++----
 2 files changed, 5 insertions(+), 5 deletions(-)
```

Because greedy LPT partitions on file line-counts and no file's line count changed, **the partition is
bit-identical to iter-116's**. Proven, not asserted — the same script
(`.agentspace/scratch/work-m257x/partition109.py`) run three ways:

| invocation | files | lines | spread | partition |
|---|---|---|---|---|
| `partition109.py` (working tree, **this reading**) | 40 | 10,871 | 43 (1533–1576) | **A…G as published below** |
| `partition109.py f581de09…` (**iter-116's ref**) | 40 | 10,871 | 43 (1533–1576) | **identical, seat for seat** |
| `partition109.py ac48e5b` (iter-109's ref, the older control) | 40 | 10,694 | 51 (1506–1557) | different — as previously published |

> ### **iter-119 IS a seat-level replicate of iter-116.** The first one this milestone has ever taken.
>
> Same 14 clone shas, same partition, same instrument, a corpus differing by 5 in-place lines that changed
> **no proposition**. iter-116 had to write *"NOT a seat-level replicate"*; this sheet gets to write the
> opposite. **That makes test-retest reliability directly measurable for the first time** — see the
> pre-registration's band #3, which is the reason this property is disclosed *before* the number.

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

| repo | checkout HEAD | `origin/main` | last fetched | iter-116 sheet | moved? |
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
| rosetta-extensions (**authoring**) | **`43049308`** on `main` | `43049308` | 2026-08-05 23:24 | `1dc1eb82` | yes — ours |

> ### **ZERO of the 14 platform clones have advanced across FOUR consecutive readings** (iter-103,
> iter-109, iter-116, iter-119).
>
> **And the ARRIVAL-vs-DETECTION question stays closed: DETECTION** (`D-M257x-109-3`). A frozen subject is
> not a reason to expect fewer drift findings here; it is the reason drift findings here are known to be
> **standing** rather than newly-arrived. Two readings have now measured drift at ~33 % and ~38 % of their
> upheld residual over a subject in which nothing moved.

**Fetch times are recorded on BOTH sides.** The most recent platform-side fetch (2026-08-06 12:15)
predates this sheet by more than half a day and predates any seat being dealt. **The fetch times are
re-read at the close and published there**, so a mid-reading move is detectable rather than merely
suspected.

### The known instrument defect, stated but NOT fixed — sixth reading running

`briefing-iter76-AS-RUN.md:37` names `.agentspace/rosetta-extensions` — the **authoring** copy — as "the
tooling". A rext claim about what a stack runs is settled in the **pinned per-stack clone**.

It is delivered **unchanged**, for the sixth reading in a row. Editing it would break the comparability the
series exists to establish. The ADDENDUM names which of the two trees settles a claim (§5 rule 45), plus
the refinement that a claim about a **fence's own verdict or configuration** is settled by the *authoring*
tree (`D-M257x-103-7`). Routed as `DEF-M257x-iter101-briefing-rext-tree`, still open, still
delivered-unfixed. Measured series of the rejection class it produces: **4 → 1 → 1 → 0 → 0**.

## The instrument — untouched, and proven so

| | value |
|---|---|
| file | `instrument/briefing-iter76-AS-RUN.md` |
| sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| re-checked | **AFTER** copying to `iter-119/briefing-AS-DELIVERED.md` **and again after the addendum was appended** — printed both times, and `diff <(head -171 delivered) instrument` returned empty on both |
| `git log --follow` on the FILE | **exactly one commit ever** — `012edd2` (iter-76) |

## Guard family at the open — 15 GREEN · **1 RED** · 4 not-run, and the RED is the SAME FALSE one

```
fence-tree: /Users/marco/workspace/anthropos/rosetta/.agentspace/rosetta-extensions @ 430493087
            describe=fast-build-m257x-iter-101-17-g4304930
guard-family: corpus /Users/marco/workspace/anthropos/rosetta @ 194361e47
guard-family: platform .../stack-demo/platform @ 0c91421df (origin/main 0c91421df, in sync)
guard-family: 20 member(s) on disk, all placed.
guard-family: 15 GREEN · 1 RED · 0 could-not-check · 4 not-run
guard-family: RED — platform_predicate_guard
```

**20 members, up from 19 at iter-116** — `corpus_citation_guard` joined the family at iter-117. The 4
not-run need `--range`/`--ledger`, which a tree-state run cannot supply — recorded as a gap rather than
hidden (`anchor_offset_guard`, `repair_leak_guard`, `repair_reach_guard`, `value_change_guard`).

**Invocation, stated with the count** (§5 rule 50):
`python3 stack-core/guard_family.py --platform <repo-root>/stack-demo/platform`, `ROSETTA_ROOT` set, from
`.agentspace/rosetta-extensions` @ `4304930`.

### The RED is `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`, unchanged and still open

```
[G10 service-count] corpus/services/sentinel.md:5 says compose declares 5 service(s);
at d11a403 `docker-compose.yml` declares 8 and the effective topology (with `include:`) has 10
— 5 is neither
```

It was characterised at iter-116's open by opening the source at both refs: at `0c91421d` — **the ref the
sentence itself names** — the counts are **(5, 7)**, so the claim is TRUE at the ref it names; the guard
takes `_REF_PINNED.search(cell)`, i.e. the **FIRST** ref in a multi-pin block, and dates the claim by
`d11a403` instead. That is §5 rule 33 violated by the guard that enforces it.

**It was routed at iter-116 and deliberately NOT repaired at iter-117 or iter-118** — both were census
iters under `TOK-08` whose scope was citation resolution, and neither touched it. It is re-disclosed here
verbatim, and it is **not repaired in this pass either**: no repair is taken inside a measuring pass.

**Why this does not block the reading, stated rather than assumed.** This iter lands **zero** corpus edits
and **zero** rext code — it is a measurement. A RED gate blocks an iter because code must not land on top
of one; there is nothing to land.

## Reading shape

7 seats × 2 independent readings (**#31, #32**) of the **identical** partition = 14 blind seats. No seat
knows which reading it is in; no seat may read `knowledge/plan/**` beyond its own briefing and its own
output file, so no seat can see a prior audit's answer key or another seat's report.

**Dealt in two batches of 7**, as at iter-103, iter-109 and iter-116. **Every seat is committed verbatim
the moment it lands**, before adjudication — the discipline that has now bounded the cost of a mid-flight
death three times.

## Test-suite state — a GAP, recorded, not a pass

This is a measuring pass: **no rext code changed, so no suite was run for it.** The standing figure for
`stack-core` remains iter-111's: **`1 failed · 1011 passed` in 1090.88 s**, and the one failure is the
perishable iter-48 answer-key fixture. **That count is quoted with its invocation** (§5 rule 50) and is
**not** re-run here, because nothing in this iter could change it.
`FIX-M257x-iter108-stackcore-suite-hangs` remains open, which is why every count in this iter names its
invocation.
