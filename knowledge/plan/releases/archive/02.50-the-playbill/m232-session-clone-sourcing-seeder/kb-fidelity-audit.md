---
title: "KB Fidelity Audit — M232 session-clone-sourcing-seeder"
date: 2026-07-19
scope: milestone:M232
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Content-stories sourcing + result-route contract | `corpus/ops/demo/content-stories-routes.md` (M231 deliverable) | jobsimulation resolvers, next-web render, prod DB | PAIRED |
| Seeding 7-table fan-out + mirror trap | `corpus/ops/demo/stories-spec.md` + `corpus/ops/seeding-spec.md` | `stack-seeding/seeders/{persona,persona_write,hiring_funnel,succession}.go` | PAIRED |
| Content-ref pinning (reserved-tail / type-aware) | `content-stories-routes.md` §3.2 + code | `stack-seeding/seeders/contentref.go` | PAIRED |
| jobsimulation service (session/result substrate, runner) | `corpus/services/jobsimulation.md` | `stack-dev/jobsimulation/*` | PAIRED |
| hiring render path (two-app, mirror read-model) | `corpus/services/hiring.md` | `apps/hiring` | PAIRED |
| db-access read boundary | `corpus/ops/db-access.md` | postgres MCP (marco_read) | PAIRED |
| Safety posture (Part 3 exposure) | `corpus/ops/safety.md` | (policy) | PAIRED |
| Demo-patch mechanism (flag enablement seam) | `corpus/ops/demo/demopatch-spec.md` | `stack-seeding`/`demo-stack` | PAIRED |
| **session-clone-spec.md** (the M232 deliverable) | — (Delivers →) | — | DOC-ONLY (milestone produces it — not a blind area) |

## Fidelity Findings

The contract doc `content-stories-routes.md` was authored + merged by M231 (2026-07-19) with code-cites; re-verified
its load-bearing claims against live code + prod DB during recon. All ALIGNED:

1. **Net-new result-substrate table schemas** — Source: content-stories-routes.md §1. Expected: `actors`,
   `interactions`, `validation_check_results`, `interview_extraction_results` present with the fan-out shape.
   Actual (MCP `information_schema`): all four exist with the cited columns (actors.username/alias PII,
   role_key, stakeholder, session_id; interactions.action_type + action_payload jsonb + target_id/source_id;
   validation_check_results.check_id/engine/parameters/success/feedback/essential;
   interview_extraction_results.user_report/manager_report jsonb NOT NULL + summary). **ALIGNED.**
2. **Modality catalog** — Source: §4. Expected: call 77 / code 65 / collaborative_doc ~30 / chat 307 public sims.
   Actual (MCP): chat 307, call 77, code 65, collaborative_doc 29, send_attachment 1 (29+1=30 document).
   **ALIGNED** (collaborative_doc 29 vs doc's 30 is point-in-time prod drift over 1 day, not staleness).
3. **Public-anchored sourcing pools per (type × modality × pass/fail)** — Source: §3.2/§4. Expected: ≥2 voice +
   1 code + 1 document assessment sources + training/hiring/interview pools exist. Actual (MCP): ASSESSMENT×call
   passed 549/failed 1049, ASSESSMENT×code passed 25/failed 22, ASSESSMENT×collab_doc failed 6, TRAINING×collab_doc
   passed 6/failed 21, INTERVIEW×call passed 35/failed 5, HIRING×call/code present. **ALIGNED — GO with margin.**
4. **Interview flag-gated resolvers** — Source: §Interview + §1. Expected: resolvers at `queries.resolvers.go:536`
   (`InterviewExtractionUserReport`) / `:563` (`InterviewExtractionManagerReport`), report is opaque pass-through
   jsonb. Actual: exact match; both return `Report *string` from raw jsonb; manager report admin-gated
   (`OrgActionAssignmentsWrite`). **ALIGNED.**
5. **The manager-view MIRROR** — Source: §mirror trap. Expected: seeding only runtime rows → blank scoreboard;
   must co-write `public.local_jobsimulation_sessions`. Actual: confirmed in `persona_write.go` (localSessionCols,
   misspelled `completition_status`) + `hiring_funnel.go` (the pair-write precedent). **ALIGNED.**
6. **CODE modality engine** — Source: §4 "in-process Judge0 (`jobsimulation/internal/runner/`; standalone
   roadrunner retired)". Actual: `stack-dev/jobsimulation/internal/runner/{runner.go,languages.go}` present.
   **ALIGNED.** (Detailed CodeSubmission/CollaborativeAsset field shapes gathered as S3 build inputs.)

## Completeness Gaps
None load-bearing. The interview `user_report`/`manager_report` JSON internal shape + the CodeSubmission /
CollaborativeAsset / transcript field shapes are **build inputs for S3** (gathered from the live checkouts, not a
KB-doc gap) — they belong in the M232 deliverable `session-clone-spec.md`, not in a pre-existing doc.

## Applied Fixes
None required — no stale claims found. Cross-references in `content-stories-routes.md` all resolve (10/10).

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed to Phase 1. The contract the milestone implements against is accurate and code-verified.
