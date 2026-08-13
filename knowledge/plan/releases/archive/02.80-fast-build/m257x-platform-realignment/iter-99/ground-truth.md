# iter-99 ground truth — re-derived at this reading's open, 2026-08-06

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
TOK-04 P1/P2 discipline. Where it agrees with a prior iter that is a re-derivation, not a copy.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD | `e858fd455c6781d00a1f6b08c5e45569ffe3d259` (`iter(M257x/98)`) |
| branch | `m257x/platform-realignment` |
| tree | clean at the open |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **40 files, 10,276 lines**, 7 seats, greedy longest-processing-time balance (1431–1506 lines/seat) |

The corpus grew 10,108 → 10,276 lines since iter-97, which is iter-98's repair (+167 −80) plus this iter's
own dirs. **A different hand is dealt than iter-97's** because the partition is recomputed from current
sizes — the method working, not drifting.

## Platform clones — the only thing that settles a claim

| repo | sha | note |
|---|---|---|
| **platform** | `0c91421dfdb08dc75f17f1aabfb61394070e770b` | **== `git ls-remote origin HEAD`, verified at this open** |
| app | `b948604ff86125a4e83516fbe356f210ddfc3809` | v1.366.0; origin/main is `2035f9a4` (v1.369.0) |
| app/studio | `aeec036` | **own nested checkout**, invisible to `git grep` at app's ref |
| cms/studio | `aeec036` | same |
| ant-academy | `9c3843cd35018c9c396fe7d511898d61dd7d260d` | |
| cms | `ca50c8170fefe1122d680efe54f7e56798a79d82` | |
| graphql-wundergraph | `60c229f39adcbbe75c84cd58f0f45052b5423372` | |
| jobsimulation | `462343b05c4f796513a43327d4d8d62d99128c4f` | |
| messenger | `fa47850d9c507d1928da7a38f7b37bac1bb8fabc` | |
| next-web-app | `bb3313bc0133ee5728ce83fda485e95bfea1a6c6` | |
| roadrunner | `87d8d44382ef07a9f165869530cbac9e5e0a4332` | |
| sentinel | `88bc55929dde7ba43913966ec3fc36372e4ff32a` | |
| storage | `4ce8ece52adb7c095e792e235da4a8913214d190` | |
| studio-desk | `14a5442a23d38860c1042e47641b4208782680c0` | |
| rosetta-extensions (per-stack) | `ab81527ae2ebfe4406bc4f1048f6c42056cd90d3` | |
| rosetta-extensions (authoring) | `5fb0915` on `main` | |

**Only `platform` has moved since the instrument was frozen** (`0dab54d → 0c91421`). Every other checkout is
byte-identical to iter-97's sheet. That is stated in the delivered briefing as a **marked addendum**, never
as an edit to the instrument.

## The instrument — untouched, and proven so

| | value |
|---|---|
| file | `instrument/briefing-iter76-AS-RUN.md` |
| sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| re-checked | **AFTER** copying to `iter-99/briefing-AS-DELIVERED.md`, not before |
| `git log --follow` on the FILE | **exactly one commit ever** — `012edd2` (iter-76) |

## Guard family at the open

`14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17 members (the 3 need `--range`/`--ledger`, which
a tree-state run cannot supply — recorded as a gap rather than hidden, and `guard_family` exits 2 to say so).

## Reading shape

7 seats × 2 independent readings (#21, #22) of the **identical** partition = 14 blind seats. No seat knows
which reading it is in; no seat may read `knowledge/plan/**` beyond its own briefing and output, so no seat
can see a prior audit's answer key or another seat's report.
