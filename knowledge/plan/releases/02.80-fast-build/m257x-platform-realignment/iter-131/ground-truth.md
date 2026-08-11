# iter-131 ground truth — re-derived at this reading's open, 2026-08-07

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
`TOK-04` P1/P2 discipline. Where it agrees with iter-119's sheet that is a re-derivation, not a copy.

**This is a MEASURING pass. No repair is taken inside it.** The separation between repairing and
measuring is the only reason any of this milestone's numbers mean anything.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD at the open | `60edbd8408711dc01b086701e25c5f5ae13cbce0` (`iter(M257x/130): tik — close-section + ledger…`) |
| branch | `m257x/platform-realignment` |
| tree at the open | **clean** — `git status --short` empty |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **41 files, 11,922 lines**, 7 seats, greedy longest-processing-time balance (**1678–1718** lines/seat, spread **40**) |

## THE INSTRUMENT — verified on both sides of the copy

| check | result |
|---|---|
| source sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| **sha256 of the copy, re-checked AFTER copying** | `3858ec53…8e52eb0` — **identical** |
| `diff source copy` | **empty** |
| `git log --follow` on the instrument | **exactly one commit ever** — `012edd2` |
| **sha256 of the delivered file's first 171 lines, AFTER the addendum was appended** | `3858ec53…8e52eb0` — **identical**, proving the addendum sits strictly BELOW the frozen text and edited nothing above it |

**The known defect at line 37** (the `rosetta-extensions` row naming only the authoring copy) is
**delivered as-is** and overridden in the addendum, never edited in place. It remains routed, not fixed.

## THIS READING IS **NOT** A SEAT-LEVEL REPLICATE — and that is the headline structural fact

iter-119 was the first true seat-level replicate this milestone ever took, because its corpus was 5
in-place lines from iter-116's. **iter-131 is the opposite**, and the sheet says so before any number:

```
git diff --shortstat 194361e4..HEAD -- corpus/services corpus/architecture
 31 files changed, 1248 insertions(+), 197 deletions(-)
```

| invocation | files | lines | spread | partition |
|---|---|---|---|---|
| `partition109.py` (working tree, **this reading**) | **41** | **11,922** | 40 (1678–1718) | **A…G as published below** |
| `partition109.py 194361e4` (iter-119's ref, the control) | 40 | 10,871 | 43 (1533–1576) | **different — seat for seat** |

**The in-scope corpus GREW 9.7 %** (10,871 → 11,922) and gained one net-new file,
`architecture/org-repos.md` (466 lines, authored at iter-123), which is ~44 % of the growth.

> **Consequences, stated in advance so they cannot be chosen afterwards:**
>
> 1. **Seat-level numbers are NOT comparable to iter-119's.** Band #9 (per-seat spread) tests this
>    reading's partition only.
> 2. **The test-retest band (#3) is measuring something different from iter-119's #3.** iter-119 asked
>    *"with essentially no repair between, how many of the 37 come back?"* — a clean recall measurement.
>    Here, substantial repair HAS happened, so a low overlap is ambiguous between *the repair held* and
>    *the instrument samples a different slice each time*. **The band is kept and the ambiguity is
>    disclosed rather than resolved by assertion.**
> 3. **`P` is measured over a LARGER subject.** A flat `P` against a 9.7 %-larger corpus is not the same
>    result as a flat `P` against an identical one, and the close will say so.

## Clone set — every ref, its `origin/main`, and its fetch time

Re-read at the open. **§5 rule 41a: no clone is fetched while this reading is in flight**; these
timestamps are re-read at the close and published, so a mid-reading move is detectable rather than
suspected.

| repo | checkout | `origin/main` | fetched (local) |
|---|---|---|---|
| ant-academy | `22df69dd8` | `22df69dd8` — level | 2026-08-06T11:18 |
| app | `ad9f3c498` | `ad9f3c498` — level | 2026-08-06T12:15 |
| cms | `ca50c8170` | `f38c0c4a4` — **behind** | 2026-08-05T23:24 |
| graphql-wundergraph | `60c229f39` | `60c229f39` — level | 2026-08-05T23:24 |
| jobsimulation | `462343b05` | `82cb66ecc` — **behind** | 2026-08-05T23:24 |
| messenger | `fa47850d9` | `e9421c68f` — **behind** | 2026-08-05T23:24 |
| next-web-app | `8297c684c` | `f97ba6599` — **behind** | 2026-08-06T12:15 |
| platform | `0c91421df` | `0c91421df` — level | 2026-08-06T12:15 |
| roadrunner | `87d8d4438` | `87d8d4438` — level | 2026-08-05T23:24 |
| sentinel | `f2c461903` | `f2c461903` — level | 2026-08-06T12:15 |
| storage | `4ce8ece52` | `9f8cb5322` — **behind** | 2026-08-05T23:24 |
| studio-desk | `41ee3575d` | `41ee3575d` — level | 2026-08-06T12:15 |
| studio-room (`stack-dev`) | `aeec036a5` | `aeec036a5` — level | 2026-08-06T13:03 |
| **rext — pinned per-stack** | `09d06070f` | `4cb920aac` — behind | 2026-08-06T11:19 |
| **rext — authoring copy** | **`f2ea567b3`** | (local `main`, ahead of origin by this run's 2 commits) | — |

**No platform clone has advanced since iter-103.** Fifth consecutive frozen reading.

⚠ **`stack-demo/ant-academy` has a DIRTY working tree** — 3 modified files. The addendum instructs every
seat to read it via `git show 22df69dd8:<path>`. **This is a live hazard for this reading**, and it is
disclosed rather than cleaned: cleaning a clone mid-reading would itself be a mutation of the substrate.

⚠ **The two rext trees are 33 commits apart** — the widest gap any reading has faced. `09d06070` is what
a stack executes; `f2ea567b` is where this run's fence work landed. The addendum carries the rule.

## Guard family at the open

**18 GREEN · 0 RED · 0 could-not-check · 4 not-run** (the four need `--range`/`--ledger`).
Fence tree printed by the family itself: `.agentspace/rosetta-extensions @ f2ea567b3`, clean.

```
/usr/bin/python3 guard_family.py --repo-root <rosetta> --platform <rosetta>/stack-demo/platform
```

`claim_census_guard`: **1,140** unevidenced assertions over 41 files, baseline 1,164, ratchet holds.

## This reading's partition

| seat | lines | files |
|---|---|---|
| A | 1710 | `external_services.md` · `platform-migration-status.md` · `storage.md` · `askengine.md` · `db-backup.md` |
| B | 1718 | `ai-readiness.md` · `clerkenstein.md` · `chronos.md` · `sentinel.md` · `dependency_map.md` · `services/README.md` |
| C | 1713 | `alignment_testing.md` · `backend.md` · `cms.md` · `clerk-integration.md` · `skillpath.md` · `gotenberg.md` |
| D | 1678 | `service_taxonomy.md` · `security_compliance.md` · `shared_libraries.md` · `ai-labs.md` · `customerio-sync.md` · `intelligence.md` |
| E | 1698 | `ant-academy.md` · `studio-desk.md` · `graphql-wundergraph.md` · `messenger.md` · `academy-backend.md` · `TEMPLATE.md` |
| F | 1709 | `studio-room.md` · `architecture_overview.md` · `jobsimulation.md` · `roadrunner.md` · `frontend_architecture.md` · `skiller.md` |
| G | 1696 | `hiring.md` · `org-repos.md` · `ai_architecture.md` · `next-web-app.md` · `coursebuilder.md` · `architecture/README.md` |

14 seats total: readings **#33** (`r33-A…G`) and **#34** (`r34-A…G`), blind and independent.

## Machine load at the open

`load1 2.38` on 12 cores; one external process (`a8-test`) holding ~1 core throughout. **Counts are
attestable; wall-clock timings are not quoted for anything in this reading**, per the standing rule.
