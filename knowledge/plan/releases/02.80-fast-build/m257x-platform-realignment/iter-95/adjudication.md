# iter-95 adjudication — readings #17 + #18 at platform `0c91421`

**Status: COMPLETE.** All 55 booked blockers across 14 seats graded by **four parallel adjudicators**,
one per seat-group, each **re-deriving from the platform/service clones** rather than from any seat's
evidence or any prior verdict.

## Verdict

| | r17-A..D | r17-E..G | r18-A..D | r18-E..G | **total** |
|---|---|---|---|---|---|
| booked | 12 | 15 | 12 | 16 | **55** |
| **UPHELD** | 12 | 14 | 10 | 15 | **51** |
| REJECTED | 0 | 1 | 2 | 1 | **4** |

**Upheld rate 92.7 %** — against iter-80's **92.1 %** and iter-84's **93.0 %**. Three readings, three
adjudications, the same rate. **The instrument is not crying wolf, and it has not drifted.**

## N — the graded number

The gate clause reads: *"KB-fidelity audit **GREEN, or YELLOW with 0 blockers**, over
`corpus/services/**` + `corpus/architecture/**`."* So the graded set is **upheld · BLOCKER-grade ·
in-scope**, deduped across both readings. Minors do not block by the clause's own words; findings in
`CLAUDE.md`, `corpus/ops/**` and `.claude/skills/**` are real but outside the clause's stated scope.

### **N = 13** distinct in-scope upheld BLOCKER anchors — **12** distinct predicates

| # | anchor | predicate | #17 | #18 |
|---|---|---|---|---|
| 1 | `external_services.md:411-414` | the deleted router's `depends_on` set never existed at any ref | ● | |
| 2 | `cms.md:95` | *"`mistralai` imported nowhere"* — 2 hits incl. a live `mistral-ocr` path | ● | |
| 3 | `backend.md:112` vs `:229` | two different production RPC addresses for the same `app` mux | ● | |
| 4 | `service_taxonomy.md:130` | `jobsimulation` ARCHIVED 2026-07-31 vs a merged PR + 4 commits 08-04 | ● | |
| 5 | `service_taxonomy.md:131` | cms ECS asserted as rollback path where the corpus rules *assert neither* | ● | |
| 6 | `platform-migration-status.md:92` | **M903 custody transfer** — wrong anchor, plan never executed | ● | ● |
| 7 | `storage.md:25` | **M903 custody transfer** (same predicate, second propagation site) | ● | ● |
| 8 | `storage.md:58` | PRIVATE storage manager claimed local-FS; compose binds it to the **production bucket** | ● | ● |
| 9 | `shared_libraries.md:42/:70` | *"every live Go service … only those four"* — unpinned currency claim, false at `0c91421` | ● | ● |
| 10 | `dependency_map.md:59` | the `skiller` stream is asserted to have a producer in `app`; there is none | ● | |
| 11 | `external_services.md:248` | *"backend's compose block carries no `DIRECTUS_*`"* — it carries one | | ● |
| 12 | `platform-migration-status.md:110` | *"the 60K-skill dataset"* stated as **fact** in the index of truth | | ● |
| 13 | `alignment_testing.md:25` | the alignment harness *"lives in rosetta"* — it is in rosetta-extensions | | ● |

Items 6 and 7 are **one predicate at two propagation sites** — the offending clause appears verbatim
in both files. By distinct predicate, **N = 12**.

### Per-seat spread of the graded set

| reading | A | B | C | D | E | F | G | total |
|---|---|---|---|---|---|---|---|---|
| **#17** | 3 | 1 | 0 | 2 | 2 | 1 | 1 | **10** |
| **#18** | 1 | 2 | 1 | 0 | 2 | 1 | 0 | **7** |

Matched by both readings: **4**. Union **13**.

### Capture–recapture on the graded set

`n₁ = 10 · n₂ = 7 · m = 4` → **Chapman N̂ ≈ 16.6**. Per-pass recall **60 % (#17) / 42 % (#18)**;
**union recall ≈ 78 %**. Consistent with the milestone's long-standing 43–51 % per-pass figure, which
has now replicated across four grading rules and five trees.

**So ~4 in-scope blockers are likely still unfound**, and N̂ is a **floor** — heterogeneous
detectability biases capture–recapture downward, and both readings share a briefing, a file set, a
partition and a model, so the estimate is optimistic on every axis.

## The 4 rejections

| seat | anchor | why it collapsed |
|---|---|---|
| r17-E B5 | `hiring.md:145` | `role_mix ≈ 0.1/0.9` is **exact** at the M223 ref the sentence names; `5d0297e` moved it to 0.14/0.86 later |
| r18-A B2 | `backend.md:224` | the bullet header names PR #896 (`app 9ecade24`); there `mux.Handle` is exactly 3 and LabSession **is** third |
| r18-D B1 | `dependency_map.md:58` | cell pins `app 9d00a313`, where `MESSENGER_ENABLED` has 0 occurrences and the subscriber **is** unconditional. `:21` vs `:58` are two ref-relative truths, not a contradiction |
| r18-E B3 | `hiring.md:81` | `manager.go:485` **is** `if !org.IsHiring {` at `5ba17044`, the ref the doc's own banner names. The seat's argument (that siblings prove a re-anchoring pass) fails: the siblings resolve identically at both refs |

All four are the **ref-discipline class** the instrument has fought since iter-76 — and all four were
caught by adjudicators applying rule 33, not by the seats. The class remains the dominant
false-positive generator and still contributes **zero** to the graded count.

## Two instrument findings that outrank several of the defects

**1. The environment's recursive `grep` silently hides tracked files.** The repo-local wrapper is
`ugrep --ignore-files`; `studio/.gitignore` lists `tools/`, so `grep -rn mistralai app/studio/`
returns **1** hit while `git grep mistralai <ref>` returns **2** — the second being a live
`from mistralai import Mistral`. Verified independently by the orchestrator.

This produced at least one **false clearance** (a seat cleared the `mistralai` predicate as an audited
zero while another seat correctly booked it), and it is almost certainly how the false claim entered
the corpus in the first place. **Every absence-claim in this milestone taken with a recursive `grep`
is suspect**, and the bias runs toward *under*-counting — which makes N a floor for a second,
independent reason.

**Standing rule this establishes:** an absence is only established by `git grep` **at a named ref**.
A recursive `grep` cannot distinguish *"not in the tree"* from *"ignored by a wrapper"*, and those are
the two states an absence claim exists to separate.

**2. A wrong positive control is a broken instrument, not a typo.** `messenger.md:22` offers
*"`-S SKILLER_RPC` returns 3"* as the control proving the pipeline works. It measures **7** at every
ref, clone and spelling tried. The substantive claim it guards is true, so the finding graded MINOR —
but a control whose whole job is to let the next reader confirm their pipeline manufactures exactly
the doubt it exists to remove.

## The escalation conditions, graded

- **">~15 union blockers → measure and route, do not repair"** — union in-scope blockers is **13**,
  just under. **Nothing was repaired regardless**, per the run's binding instruction: a measuring pass
  may not contain a repair. That separation is the only reason 140 → 43 meant anything.
- **"A seat that cannot state a `wc -l`"** — did not fire. All 14 seats stated per-file positive
  controls; three seats' harnesses blocked their own file writes and their reports were persisted
  verbatim by the orchestrator, which is a delivery failure, not a reading failure.
- **"Any finding inside a class a prior iter closed by adjudication"** — did not fire.

## Propagation the repair must inherit (routed, NOT repaired)

Adjudicators flagged **unbooked twins of upheld predicates** — sites no seat booked because no seat
was assigned them:

- the `mistralai` predicate is live in **four** files (`cms.md:95`, `service_taxonomy.md:209`,
  `dependency_map.md:25`, `studio-room.md:349/:359`); only one was booked
- the prod-RPC-hostname split also sits at `skiller.md:19`
- the `jobsimulation`-archive predicate twins at `architecture_overview.md:20`, `roadrunner.md:40`
- the cms-ECS predicate twins at `architecture_overview.md:222`
- **the M903 clause was positively CLEARED at `platform-migration-status.md:92` by one seat while
  another seat booked it at `storage.md:25`** — the fenced map propagated the false clause and a seat
  cleared it there

**Repair by PREDICATE, never by anchor.** A repair pass that fixes the 13 booked anchors and stops
will leave the same predicates standing in at least 8 other places.
