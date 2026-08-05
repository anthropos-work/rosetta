# iter-97 adjudication — readings #19 + #20 at platform `0c91421`, corpus `00be1ac`

**Status: COMPLETE.** All 58 booked blockers across 14 seats graded by **four parallel adjudicators**,
one per seat-group, each **re-deriving from the platform/service clones** rather than from any seat's
evidence or any prior verdict.

## Verdict

| | Adj1 · r19-A..D | Adj2 · r19-E..G | Adj3 · r20-A..D | Adj4 · r20-E..G | **total** |
|---|---|---|---|---|---|
| booked | 14 | 13 | 15 | 16 | **58** |
| **UPHELD** | 12 | 12 | 14 | 16 | **54** |
| REJECTED | 2 | 1 | 1 | 0 | **4** |
| in-scope upheld BLOCKERS | 8 | 4 | 6 | 7 | — |

**Upheld rate 93.1 %** — against 92.1 % (iter-80), 93.0 % (iter-84), 92.7 % (iter-95). **Four
adjudications, four rates inside 1.0 points.** The instrument's precision has not drifted in four
readings across five trees and two platform refs.

## **N = 20** distinct in-scope upheld BLOCKER anchors — **17** distinct predicates

| # | anchor | predicate | #19 | #20 |
|---|---|---|---|---|
| 1 | `external_services.md:565` (+`cms.md:223`) | `studio/tools/pdf2md.py` "nothing dispatches it" — `tools/r3.py:139,190,199-206` execs it as pipeline step 2 | ● | |
| 2 | `cms.md:95` | "Go execs only `studio/gen.py`" — `studioManager.go:1045` also execs `studio/postgen.py` | ● | |
| 3 | `backend.md:13` | cms's M810 prod teardown asserted as "still pending" where the corpus rules it unmeasurable | ● | |
| 4 | `platform-migration-status.md:205` | the map's own fence coverage stated as 66 range-only citations; the guard reports **69** | ● | |
| 5 | `messenger.md:17` | messenger's `backend` subscriber has **21** handlers over `flow.go:72-104`, not 20 over `:70-95` | ● | ● |
| 6 | `coursebuilder.md:130-131` | `app`'s **current** version stated as `v1.363.2`; the checkout is `v1.366.0` | ● | ● |
| 7 | `roadrunner.md:114` | the `jobsimulation` repo holds one `roadrunner` mention (5 exact / 14 case-insensitive) | ● | ● |
| 8 | `roadrunner.md:24-26` | (same predicate) "zero hits outside CHANGELOG" / "no other platform repo references it" | ● | |
| 9 | `service_taxonomy.md:3/:128/:129/:364` | GitHub archive state asserted flat, three rows from `:130`'s own retraction of exactly that | ● | |
| 10 | `ant-academy.md:255` | Fonts row names DM Sans + JetBrains Mono; the app loads **Work Sans + Instrument Serif** | ● | ● |
| 11 | `ant-academy.md:322-326` | "the cockpit no longer sets `e2e_persona`" — two live paths set it, and `:331` says so | ● | |
| 12 | `dependency_map.md:59` | the new skiller cell names **two** `app` refs; anchors resolve at one, "one Go occurrence" false at the other | ● | ● |
| 13 | `backend.md:299` | cites a top-level `migrations/` dir deleted from `app` on 2026-06-18 | | ● |
| 14 | `services/README.md:11-13` | `customerio-sync`'s prod teardown asserted where two other docs declare it unmeasured | | ● |
| 15 | `architecture_overview.md:334-336` | `STORAGE_RPC_ADDR` "3 repo-wide, every one a comment" — 29 lines / 18 files at the named ref | | ● |
| 16 | `graphql-wundergraph.md:85-87` | a load-bearing self-citation resolves to a different construct (the claim is at `:192`) | | ● |
| 17 | `hiring.md:267-269` | (same predicate) `:157-159` is the `job_position` bullet; the adjudicated form is `:169-182` | | ● |
| 18 | `shared_libraries.md:77` | "12 Connect-RPC services defined" omits `StorageService`; ≥13, and `storage.md:115` says so | | ● |
| 19 | `external_services.md:614` | (same predicate) in-file `:543` self-anchor **not** re-pointed by the repair that moved it | | ● |
| 20 | `backend.md:33-34` | `app` owns the skiller stream "both producer and consumer" — no publisher at either ref | | ● |

Two predicate collapses: **{16, 17, 19}** are one predicate (a load-bearing intra-corpus citation
resolving to a different construct) and **{7, 8}** are one. 20 anchors → **17 predicates**.

### Per-seat spread of the graded set

| reading | A | B | C | D | E | F | G | total |
|---|---|---|---|---|---|---|---|---|
| **#19** | 2 | 2 | 2 | 2 | 1 | 2 | 1 | **12** |
| **#20** | 0 | 2 | 2 | 1 | 2 | 2 | 4 | **13** |

Matched by both readings: **5**. Union **20**.

### Capture–recapture

`n₁ = 12 · n₂ = 13 · m = 5` → **Chapman N̂ ≈ 29.3**. Per-pass recall **40.9 % / 44.3 %**; union recall
**≈ 68 %**. **≈ 9 in-scope blockers are estimated still unfound**, and N̂ remains a floor for the
reasons iter-95 recorded (heterogeneous detectability biases capture–recapture down; both readings
share a briefing, file set, partition and model).

## The 4 rejections — the ref-discipline class again, now 13 occurrences over four readings

| seat | anchor | why it collapsed |
|---|---|---|
| r19-A B5 | `external_services.md:3` | `graphql-wundergraph` @ `60c229f3` has no commit after 2026-07-30 — the clone is *consistent* with the dated archive, and the fenced map publishes the identical claim |
| r19-B B5 | `sentinel.md:5` | the paragraph names `0dab54d`, where messenger's block exists exactly as described; the currency claim it supports is true at HEAD too |
| r19-G B7 | `platform-alignment.md:209` | the sentence is *"**Measured 2026-08-01**"* — a pinned past-tense census, and the 2026-08-04 evidence post-dates the pin |
| r20-A B7 | `external_services.md:3` | same as r19-A B5, independently |

**Still contributing zero to any graded count**, across four readings. Filtered, not fixed — as directed.

## The pre-registration graded — **3 of 7 held, 4 FAILED**

| # | prediction | band | outcome |
|---|---|---|---|
| 1 | per-reading count | [2,7] both | **FAILED** — 12 and 13, both above |
| 2 | union N | [3,9] | **FAILED** — 20 |
| 3 | zero recurrence of the 12 repaired predicates | exactly 0 | **FAILED** — **3** (#3 = P5, #9 = P4, #20 = P10) |
| 4 | upheld rate | [86 %, 96 %] | **HELD** — 93.1 % |
| 5 | per-pass recall | [30 %, 62 %] | **HELD** — 40.9 % / 44.3 % |
| 6 | platform-move class < 50 % of upheld in-scope blockers | < 50 % | **HELD** — **1 of 20 (5 %)** |
| 7 | repair induces ≥ 2 upheld in-scope blockers | ≥ 2 | **HELD** — exactly **2** genuinely repair-*caused* (#12, #19); 14 of 20 sit in files the repair touched |

**Four of seven failed, and that is the pre-registration working.** iter-95 graded 6 of 6 and recorded
it as a warning; the bands were narrowed until they could fail, and they did — including on the one
prediction that had teeth against my own repair (#3).

## What the three recurrences actually say

They are **not** three of a kind:

- **#9 (`service_taxonomy.md`, P4 archive)** — `D-M257x-96-6` explicitly **routed** this class rather
  than sweeping it, on the ground that no clone can measure the other five repos. The reading found the
  routed item. That is the routing being correct and the scope being too narrow, not a miss.
- **#3 (`backend.md:13`, P5 cms ECS)** — a **third** site of a predicate the repair fixed at two. The
  predicate sweep enumerated `rollback path` and found two; this site says *"still pending"*, a
  paraphrase no string sweep reaches. **iter-93's rule — fencing a document does not fence its
  paraphrases — cost this one.**
- **#20 (`backend.md:33`, P10 skiller stream)** — a genuine miss, and the worst kind: the repair fixed
  `backend.md:127` and left `backend.md:33` standing **in the same file**. The exact dominant class the
  whole predicate-wise method exists to prevent, committed while executing that method.

`claim_twin_guard` is GREEN over all 14 refuted forms, and all three of these are live. **That measures
the fence's reach, not its correctness**: it matches quoted verbatim forms, and every one of these three
is a paraphrase. The fence is necessary and demonstrably not sufficient.

## Propagation the next repair must inherit (routed, NOT repaired)

Adjudicators named unbooked twins while re-deriving:

- the **stale `v1.363.2` currency pin** also at `ai-labs.md:18` and `academy-backend.md:20` (the latter
  says *"currently"*, and is self-contradicted at `:55`)
- the **archive-state assertion** at `platform-migration-status.md:96/:97/:98/:100`, `roadrunner.md:32`,
  `architecture_overview.md:3`, `external_services.md:3/:418/:459/:479`, `graphql-wundergraph.md:8` —
  **10+ in-scope sites**, in the same fenced table that retracts the predicate
- the **truncating `flow.go:70-95` range** also at `services/README.md:39`
- the **"Go execs only `gen.py`"** twin at `studio-room.md:371`
- the correct targets for the wrong-construct citations: `service_taxonomy.md:62`,
  `graphql-wundergraph.md:192` and `:115-116`, `hiring.md:170-175`, `backend.md:241`, `storage.md:115`

## Out-of-scope but real (reported, not counted)

`platform-alignment.md:1236` — **rule 44's "1,178 NULs" is false; the file has 1.** 1,178 is its line
count: `grep -c` counts matching *lines*, and the zsh `$'\x00'` pattern degenerates to empty and matches
every line. **The rule about wrong instruments was written with a wrong instrument**, caught by the
adversarial seat in *both* readings. Also `:1254-1258` — the printed recipe labels a file count as
`hits=` and drops `-i`, returning 2 where the rule's own worked example publishes 22; and `:1260` sits
outside its own loop with `$d` unbound. Plus `platform-alignment.md:1345`, `CLAUDE.md:280`, and
`safety.md:203/:207` (the S3-private row's Class column now disagrees with
`stack-seeding/isolation/isolation.go:106`, which still registers it `PerStackIsolated` — the repair
asserted a code change that did not happen).

**None repaired.** A measuring pass may not contain a repair; all of it routes to
`FIX-M257x-iter97-read-union`.
