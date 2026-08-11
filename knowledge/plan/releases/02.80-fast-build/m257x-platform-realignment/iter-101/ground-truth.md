# iter-101 ground truth — re-derived at this reading's open, 2026-08-06

**Nothing here is inherited.** Every value below was re-measured at the open of this reading, per the
TOK-04 P1/P2 discipline. Where it agrees with iter-99 that is a re-derivation, not a copy.

## Corpus under audit

| | value |
|---|---|
| rosetta HEAD | `8f04d3ae3b76de359f990b80b6b47eced7c0b31c` (`iter(M257x/100)`) |
| branch | `m257x/platform-realignment` |
| tree | clean at the open (`git status --porcelain` → 0 lines) |
| scope | `corpus/services/**` + `corpus/architecture/**` |
| partition | **40 files, 10,278 lines**, 7 seats, greedy longest-processing-time balance (**1431–1506** lines/seat) |

The corpus moved **10,276 → 10,278 lines** since iter-99 — iter-100's citation repair, `+2` net across 5
in-scope files. **The partition is recomputed from current sizes and comes out identical to iter-99's**
(same 7 seats, same file-to-seat assignment, same 1431–1506 spread), because a 2-line change cannot move a
greedy LPT balance. That makes this reading a **replicate on a fixed subject**, which is the first time the
instrument has been run twice over materially the same pool.

## What iter-100 changed in scope, and what it did not

| file | change |
|---|---|
| `corpus/architecture/service_taxonomy.md` | bare `:137`/`:138` anchors given their file name |
| `corpus/services/ai-readiness.md` | `:458 → :459`; `useAIReadiness.ts:274 → :326` |
| `corpus/services/hiring.md` | `manager.go:448 → :450`, `:485 → :537`, `siminvitationlink.go:62 → :63` |
| `corpus/services/messenger.md` | the consumer row's bare `9d00a313` ref re-stated against `2035f9a` |

**iter-99's 28 upheld blockers were NOT repaired** — they are routed as `FIX-M257x-iter99-read-union` and
remain unpaid. At most ~4 of the 28 are touched by the anchor repairs above.

## Platform clones — the only thing that settles a claim

| repo | sha | note |
|---|---|---|
| **platform** | `0c91421dfdb08dc75f17f1aabfb61394070e770b` | **== `git ls-remote origin HEAD`, verified at this open** |
| app | `b948604ff86125a4e83516fbe356f210ddfc3809` | v1.366.0 |
| app/studio | `aeec036a51c8a4ae0c5b8f7d5d21cfa7086b658e` | **own nested checkout**, invisible to `git grep` at app's ref |
| cms/studio | `aeec036a51c8a4ae0c5b8f7d5d21cfa7086b658e` | same |
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
| rosetta-extensions (**per-stack, pinned**) | `ab81527ae2ebfe4406bc4f1048f6c42056cd90d3` | `stack-demo/rosetta-extensions` |
| rosetta-extensions (**authoring**) | `09d06070fd99c742d7a671c468abf93074278575` on `main` | `.agentspace/rosetta-extensions` |

**Nothing has moved since iter-99's sheet except the rext authoring copy** (`5fb0915 → 09d0607`, iter-100's
guard repair). `platform` is unchanged at `0c91421` and re-verified level with origin.

### The known instrument defect, stated but NOT fixed

`briefing-iter76-AS-RUN.md:37` names `.agentspace/rosetta-extensions` — the **authoring** copy — as "the
tooling". A rext claim in the corpus is settled in the **pinned per-stack clone** `ab81527a`. In iter-99
**two independent seats followed line 37 correctly and both bookings were rejected**, i.e. the instrument
manufactures identical false bookings by construction (4 of that reading's 10 rejections).

It is delivered **unchanged**. Editing it would break the comparability this replicate exists to establish,
and the class is caught by adjudication, so `N` is unaffected. The cost lands on the upheld rate, which is
why band #6 measures it and why the rate is reported twice. Routed as `DEF-M257x-iter101-briefing-rext-tree`.

## The instrument — untouched, and proven so

| | value |
|---|---|
| file | `instrument/briefing-iter76-AS-RUN.md` |
| sha256 | `3858ec536ea613ee00dfa8b383f858486564e23815581ef5135e85ec18e52eb0` |
| re-checked | **AFTER** copying to `iter-101/briefing-AS-DELIVERED.md`, not before |
| `git log --follow` on the FILE | **exactly one commit ever** — `012edd2` (iter-76) |

## Guard family at the open

`14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17 members, corpus `8f04d3ae3`, platform
`0c91421df` (origin/main in sync). The 3 need `--range`/`--ledger`, which a tree-state run cannot supply —
recorded as a gap rather than hidden, and `guard_family` exits 2 to say so.

`anchor_construct_guard` reach at this open: **528 of 555** citations adjudicated, **0 findings** — the
iter-100 repair, in place and green over the subject it now reaches.

## Reading shape

7 seats × 2 independent readings (#23, #24) of the **identical** partition = 14 blind seats. No seat knows
which reading it is in; no seat may read `knowledge/plan/**` beyond its own briefing and output, so no seat
can see a prior audit's answer key or another seat's report.

## The partition

| seat | lines | files |
|---|---|---|
| A | 1506 | `external_services.md` · `chronos.md` · `messenger.md` · `frontend_architecture.md` · `gotenberg.md` |
| B | 1471 | `ai-readiness.md` · `cms.md` · `storage.md` · `coursebuilder.md` · `services/README.md` · `TEMPLATE.md` |
| C | 1493 | `alignment_testing.md` · `backend.md` · `ai_architecture.md` · `customerio-sync.md` · `skillpath.md` · `skiller.md` |
| D | 1467 | `studio-room.md` · `platform-migration-status.md` · `jobsimulation.md` · `sentinel.md` · `askengine.md` · `architecture/README.md` |
| E | 1449 | `service_taxonomy.md` · `clerkenstein.md` · `security_compliance.md` · `ai-labs.md` · `clerk-integration.md` · `intelligence.md` |
| F | 1461 | `ant-academy.md` · `architecture_overview.md` · `shared_libraries.md` · `roadrunner.md` · `dependency_map.md` · `db-backup.md` |
| G | 1431 | `studio-desk.md` · `hiring.md` · `graphql-wundergraph.md` · `academy-backend.md` · `next-web-app.md` |
