---
title: "KB Fidelity Audit — M240 content-stories fidelity"
date: 2026-07-21
scope: milestone:M240
invoked-by: build-milestone
---

## Verdict
YELLOW — proceed with tracking. No un-promoted blind area; the two "stale" doc claims ARE the
defects this milestone fixes (the implementation knows them as defects, not as truth-to-rely-on).

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Content-story sourcing (selection) | `corpus/ops/demo/session-clone-spec.md §2`, `content-stories-routes.md §3.5` | `contentsession/sourcing.go` | PAIRED |
| Capture + scrub pipeline | `session-clone-spec.md §3` | `cmd/content-capture/main.go`, `scrub/`, `contentsession/content.go` | PAIRED |
| Replay / seed write (document) | `session-clone-spec.md §3–4` | `seeders/content_stories_write.go`, `persona_write.go::criterionResultCols` | PAIRED |
| Modality substrate (transcript/code/doc/interview) | `session-clone-spec.md §4` | `seeders/content_stories_modality.go` | PAIRED |
| Read-side safety exception | `corpus/ops/safety.md §3.8` | the read boundary + `scrub` | PAIRED |
| Voice recording / media blob (Chime/S3) | — (Delivers: new media-substrate spec) | `jobsimulation` Recording/ChimeRecording ent + prod S3 | BLIND-AREA (promoted) |

## Fidelity Findings

### KB-1 — criterion `input_data` is captured but DROPPED at seed time (STALE; in-scope = Defect 3)
- **Source:** `session-clone-spec.md §3` (table row "criterion title + candidate submission … `.input_data`")
  and `§4` ("validation_criterion_results (the REAL titles/input_data; input_format per capture)").
- **Expected (doc):** the candidate submission `validation_criterion_results.input_data` is copied AND replayed.
- **Actual (code):** `cmd/content-capture/main.go` DOES capture `input_data` (→ `CapturedCriterion.InputData`),
  but the seeder's criterion write reuses the **shared** `criterionResultCols()` (`persona_write.go:161`) which
  has **no `input_data` column**, and `content_stories_write.go` step 5 never appends `cr.InputData` → the column
  is dropped, so the manager sees no document.
- **Verdict:** STALE — code drifted from the doc-as-contract. **Fix owner: update code (Defect 3)** with a
  content-specific criterion col set (the shared PersonaSeeder set is a landmine), then bump the doc.
- **Not a blocker:** the milestone's whole point is to fix this; the implementation is not misled.

### KB-2 — cell type-match predicate undocumented + unenforced (COMPLETENESS/STALE; in-scope = Defect 1)
- **Source:** `session-clone-spec.md §2` (documents exactly 2 load-bearing predicates: public-anchoring +
  non-manager-played).
- **Expected:** a cell's sourced session should be on a sim whose TYPE matches the cell type.
- **Actual:** `sourcing.go SelectionSQL` filters `s.sim_type` and the modality's task_type, but the public-sim
  CTE admits ANY public sim with a matching task — including the interview sim (`ai-readiness-interview-d62`,
  which has a `call` task) → an assessment-voice cell can source an interview-flavored session (the CQ-1 root
  cause). No 3rd predicate excludes the interview sim from non-interview cells.
- **Verdict:** STALE/incomplete — **Fix owner: update code (Defect 1)** + document the new predicate.
- **Not a blocker:** the defect is the milestone's work.

### KB-3 — media substrate (voice recording + document blob) is a BLIND-AREA, promoted (acceptable)
- **Source:** overview.md `delivers: corpus/ops/demo/ (new media-substrate spec) + safety.md §3.8 (raw-media amendment)`.
- **Expected:** a knowledge doc for capturing/re-hosting the Chime/S3 recording + document blob.
- **Actual:** no such doc exists yet; `safety.md §3.8` covers scrubbed FREE-TEXT only, not raw audio / full
  documents; `content_stories_write.go:72` hardcodes `sessChimeNotAvailable`. The `Recording{recording_id,
  session_id}` reference is in Postgres; the **audio bytes are in prod S3**.
- **Verdict:** BLIND-AREA but PROMOTED to a `Delivers →` line (new media-substrate spec + §3.8 amendment) — per
  the gate rule this is acceptable, not RED.

## Completeness Gaps
- The interview render-fidelity exact-section match is flagged UNPROVEN in `session-clone-spec.md §4` already —
  out of M240 scope (owned by M244's live sweep). No action.

## Environment finding (load-bearing for the gate outcome, not a doc-fidelity issue)
- **S3 media-read is NOT available in this environment** (`~/.aws/credentials` = 0 bytes; no AWS env/profile;
  `aws sts get-caller-identity` → "Unable to locate credentials"). Prod-read (`marco_read`) IS available.
  Consequence: the **blob-port half** of the media-substrate deliverable (Defect 2 voice audio bytes; Defect 3
  doc body IF it is an S3 `storage_upload`) is **S3-blocked = the DEF-M10-01 genuine blocker** the roadmap
  flagged. The safety amendment, Defect 1, Defect 3 (input_data-at-seed-time), and pass-rate are all landable
  on prod-read alone.

## Applied Fixes
- None inline — the two STALE claims are corrected by the milestone's own Defect-1/Defect-3 code changes +
  their doc bumps in Phase 5; KB-3 is delivered by the new media-substrate spec.

## Open Items (require user decision)
- S3 read creds for eu-west-1 (DEF-M10-01) — surfaced as the milestone's genuine blocker for the media-blob port.

## Gate Result
YELLOW: proceed with tracking. KB-1/KB-2 resolved by the milestone's Defect-1/Defect-3 code + doc bumps; KB-3
delivered by the new media-substrate spec + the §3.8 amendment. The S3 blob-port is a separately-surfaced
genuine blocker (DEF-M10-01), not a doc-fidelity RED.
