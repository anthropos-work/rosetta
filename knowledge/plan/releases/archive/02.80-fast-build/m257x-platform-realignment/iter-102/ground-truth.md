# iter-102 ground truth — re-derived at this REPAIR pass's open, 2026-08-06

**Nothing here is inherited.** Every value was re-measured at the open. Where it agrees with iter-101's
sheet that is a re-derivation, not a copy.

**This is a REPAIR pass. No reading is taken inside it.** The separation between repairing and measuring is
the only reason any of this milestone's numbers mean anything (iter-95 §, iter-97 §, iter-99 §, iter-101 §).

## Corpus under repair

| | value |
|---|---|
| rosetta HEAD at open | `c716109` (`audit(M257x): pre-stage the blocking deferral gate`) |
| branch | `m257x/platform-realignment` |
| tree at open | clean but for Lane B's untracked `gate-clauses-1-2/raw/` |
| scope | `corpus/**` + `CLAUDE.md` (the repair scope is WIDER than clause 5's read scope — see below) |
| concurrent lanes | **Lane B** owns `stack-demo/**`, git tags, `.agentspace/rext.tag`; **Lane D** owns `deferrals-audit.md` |

**Repair scope vs read scope.** Clause 5's reading is scoped to `corpus/services/**` +
`corpus/architecture/**`. The REPAIR is not so scoped, and deliberately: a predicate refuted inside the read
scope is refuted **everywhere**, and iter-101's finding #2 is exactly the case — the
`sentinel-only-cross-process-edge` predicate is `CLAUDE.md`'s claim **verbatim**, in this repo's own root
instructions, which no reading has ever been allowed to look at. Repairing the three in-scope anchors and
leaving the root instructions asserting the refuted form would republish it into every future agent's
context.

## Platform clones — and the fact that they MOVED under the previous reading

**Five platform clones advanced between iter-101's ground-truth sheet and this open.** Lane B ran a
clone-set refresh at **11:18:16 – 11:20:51**; iter-101's adjudication commit (`a360d66`) landed at
**11:21:55**. See `DEF-M257x-iter101-crosslane-fetch` in `decisions.md`.

| repo | iter-101 sheet | now (HEAD) | Δ commits | Δ files | corpus citations pinned at the OLD sha |
|---|---|---|---|---|---|
| **app** | `b948604f` | **`ad9f3c49`** | **98** | **634** | **17** |
| **next-web-app** | `bb3313bc` | **`8297c684`** | 41 | 192 | 1 |
| **ant-academy** | `9c3843cd` | **`22df69dd`** | 5 | 86 | 0 |
| **sentinel** | `88bc5592` | **`f2c46190`** | 2 | 3 | 4 |
| **studio-desk** | `14a5442a` | **`41ee3575`** | 2 | 9 | 0 |
| platform | `0c91421d` | `0c91421d` | 0 | 0 | — (`== git ls-remote origin HEAD`, re-verified at this open) |
| cms | `ca50c817` | `ca50c817` | 0 | 0 | — (origin/main has moved to `f38c0c4a`; checkout has not) |
| jobsimulation | `462343b0` | `462343b0` | 0 | 0 | — (origin/main `82cb66ec`) |
| messenger | `fa47850d` | `fa47850d` | 0 | 0 | — (origin/main `e9421c68`) |
| storage | `4ce8ece5` | `4ce8ece5` | 0 | 0 | — (origin/main `9f8cb532`) |
| roadrunner | `87d8d443` | `87d8d443` | 0 | 0 | — |
| graphql-wundergraph | `60c229f3` | `60c229f3` | 0 | 0 | — |

`app`'s **`origin/main` moved `2035f9a4 → ad9f3c49`** — **5 commits, 5 files**
(`.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf`,
`terraform/variables.tf`).

### What that move does and does not invalidate

**A pin is a pin.** `2035f9a4` is still a reachable ancestor in the clone, so every citation that *pins*
`2035f9a` still resolves and still means what it meant. What expired is the **LABEL**: the corpus calls
`2035f9a` **`origin/main`**, and origin/main is now `ad9f3c49`. iter-95's own seat report already named this
exactly — *"still holds at 2035f9a; only the 'origin/main' LABEL expired"* (`iter-95/raw/r18-B.md:149`).

**Measured, at this open:**

| | count |
|---|---|
| occurrences of `2035f9a` anywhere in `corpus/` + `CLAUDE.md` | **23** |
| of those, sites that **LABEL it `origin/main`** — the defect class | **17** (15 in `corpus/`, **2 in `CLAUDE.md`**) |
| sites that pin it **without** the label (correct as written) | 6 |

The brief's *"15 corpus sites"* is exactly right for `corpus/`. **`CLAUDE.md` adds two more that no reading
could have seen**, for the same reason as the predicate above: the root instructions are outside every
read scope this milestone has ever used.

## The instrument — NOT consulted, and deliberately

`instrument/briefing-iter76-AS-RUN.md` is not read, copied, or executed in this pass. A repair pass has no
reading in it. Its sha is asserted unchanged only as a tree fact:
`3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0`.

## Guard family at the open

`14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17 members — corpus `c71610922`, platform
`0c91421df` (origin/main in sync, fetched 10 m ago). Identical to iter-101's open. The 3 need
`--range`/`--ledger`, which a tree-state run cannot supply; `guard_family` exits 2 to say so and the run is
accepted with `--allow-not-run`, which **records** the gap rather than hiding it.

`claim_twin_guard` at the open: **134 adjudicated claims**, none published anywhere in the tree.
`anchor_construct_guard`: OK over its (iter-100-widened) subject.

## The subject of this repair — 52 anchors, three unions

| union | anchors | routed at | paid? |
|---|---|---|---|
| `FIX-M257x-iter99-read-union` | **28** | iter-99 | **NO** — iter-100 deliberately did not pay it so iter-101's replicate would run on a fixed subject |
| `FIX-M257x-iter101-read-union` | **24** | iter-101 | **NO** |
| **total** | **52** | | |

`FIX-M257x-iter95-read-union` (13 anchors) was **paid by iter-96** (13 → 51 sites) and
`FIX-M257x-iter97-read-union` (20 anchors) **by iter-98** (20 → 37 sites). The run brief named iter-95 and
iter-99 as the two unpaid unions; **re-derived, the unpaid pair is iter-99 and iter-101**, and their sum is
the 52 the brief specifies. iter-101's own ground-truth sheet states it in terms:
*"iter-99's 28 upheld blockers were NOT repaired — they are routed as `FIX-M257x-iter99-read-union` and
remain unpaid."*

**The reason iter-100 withheld payment has expired.** The replicate is done and its verdict is in.
