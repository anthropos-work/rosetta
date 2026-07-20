# M231 — Spec notes

Topic → doc → code triples + prove-by-DB / prove-by-render findings accumulate here during build.

## Pre-flight audits — S1 (first section)
- Phase 0b KB-fidelity: invoked `/developer-kit:audit-kb-fidelity --milestone=M231`. Report: `kb-fidelity-audit.md`. Verdict recorded there.
- Topic → doc → code:
  - Sim session/result read-model → `corpus/services/jobsimulation.md` → `stack-dev/{jobsimulation,next-web-app,app}`
  - Hiring render path → `corpus/services/hiring.md` → `stack-dev/next-web-app/apps/{web,hiring}`, `stack-dev/app`
  - Skillpath session/result → `corpus/services/skillpath.md` → `stack-dev/skillpath`, `stack-dev/next-web-app`
  - Assessment modalities → `corpus/architecture/ai_architecture.md` → `stack-dev/{jobsimulation,roadrunner,app}`
  - AI-labs → `corpus/services/backend.md` (labsession) → `stack-dev/app/internal/labsession`
  - Academy session → `corpus/services/ant-academy.md` → `stack-dev/ant-academy`
  - Prod read path → `corpus/ops/db-access.md` (verified: `postgres`/`marco_read`/`10.2.22.13`, read-only) — ALIGNED
  - Exposure/anon posture → `corpus/ops/safety.md` Part 3 — read, current

## DB scouting (prod, read-only, catalog-only + bounded — honors db-access.md)

### Session read-model = a PERSISTED fan-out (not ephemeral) — S1 DB-side evidence
Every `jobsimulation.sessions` row carries persisted `score` (real), `result_status` (`completed`/`pending`/`waiting_validation`), `status`, `completion_status`, `ended_at`. The completed-session result substrate, verified per-type by pinned-id lookup:

| sim_type | completed (prod) | per-session child rows (verified) |
|---|---|---|
| SIMULATION_TYPE_ASSESSMENT | 5,172 | 1 validation_attempt_result · 3 actors · 1 mirror (score=0 no-shows: 0 var) |
| SIMULATION_TYPE_TRAINING | 1,799 | 1 var · 1–3 skill_results · 4–5 criterion_results · 2–3 actors · 1 mirror |
| SIMULATION_TYPE_HIRING | 1,679 | 1 var · 3 skill_results · 6 criterion_results · 2 actors · 1 mirror |
| SIMULATION_TYPE_INTERVIEW | 488 | 1 var · 1 skill_result · 4–5 criterion_results · 2 actors · **1 interview_extraction_results (user_report+manager_report)** · 1 mirror |

Also present (conversation-modality sims, mostly `result_status=pending`): CHAT_CONVERSATION 1,001 · EMAIL_CONVERSATION 57 · FASHION_STORE_CONVERSATION 61.

### The MIRROR trap — confirmed in the data (M219/M222)
`public.local_jobsimulation_sessions` = 19,870 rows ≈ `jobsimulation.sessions` = 19,873 → **~1:1 mirror**, 1 mirror row per completed session confirmed. Columns: `jobsimulation_session_id` (FK), `user_id`, `organization_id`, `tenant_id`, `score`, `status`, `completition_status` [sic], `interaction`, `anticheat_{score,summary,tagline}`. This is the **manager/comparison read-model** (hiring.md: "the score is a MIRROR table in app"). Seeding a result WITHOUT this mirror row → manager view blank. Skillpath has the analogous `public.local_skill_path_sessions` (8,183 ≈ `skillpath.skill_path_sessions` 8,260).

Table sizes (all scout-safe): sessions 17 MB/19,873 · validation_attempt_results 34 MB/14,022 · **interactions 284 MB/315,836 (avoid full scans)** · actors 23 MB/62,980 · interview_extraction_results 5 MB/367 · local_jobsimulation_sessions 11 MB · skillpath.{skill_path 3.6MB/8,260, chapter 11MB/37,378, step 42MB/135,303}.

### S2 — sourcing + anonymization surface (columns classified WITHOUT reading values)
Pin a source by **`jobsimulation.sessions.id` (uuid)** — the deterministic source-pin.

**Scrub-clean structured (re-key / re-tenant deterministically):** all `*_id` (owner_id, organization_id, tenant_id, sim_id, session_id, target/source_id, timer_id…), `token`.
**Keep as-is (non-PII structured):** enums (sim_type, status, completion_status, result_status, acceptance_status, evaluation_status, chime_status, language, criterion type/input_format), numerics (score, success_threshold, competency_level_score, interactions_progress, validation_version, criterion_index), timestamps, booleans, skill node names, role_key.
**FREE-TEXT needing handling (LLM feedback + candidate work-product + transcript + names):**
- `actors.username`, `actors.alias` → **direct-PII names** (candidate + stakeholders); scrub to the anonymized player identity.
- `validation_attempt_results.{explanation_summary, personal_explanation_summary, quick_summary}`
- `validation_attempt_skill_results.{strengths_feedback, weaknesses_feedback, personal_strengths_feedback, personal_weaknesses_feedback, quick_summary}`
- `validation_criterion_results.{title, explanation_summary, personal_explanation_summary, strengths/weaknesses_feedback, personal_*_feedback, quick_summary}` + **`input_data` (jsonb) = candidate's raw submission** (sharpest edge) + `skills` (jsonb)
- `interactions.action_payload` (jsonb) = **the transcript** (candidate's own words — highest PII risk; 284MB table)
- `interview_extraction_results.{user_report, manager_report, summary}` (jsonb) = LLM reports (may name/quote candidate)
- `local_jobsimulation_sessions.anticheat_{summary,tagline}` (free-text)
Mechanism CONFIRMED viable (read → filter per type → pin by id). The COPY is M232; the release amends safety.md Part 3 (anonymized-real, VPN-scoped).

### S3 — public-sim-by-modality catalog: GO with huge margin
Modality lives in `directus.sim_tasks.task_type`. Public-published predicate (`private=false AND tenant_id IS NULL AND status='published'`), distinct public source sims per modality:

| task_type | modality | distinct public sims | engine |
|---|---|---|---|
| `call` | **VOICE** | **77** | LiveKit + AWS Chime |
| `code` | **CODE** | **65** | Judge0 / Roadrunner |
| `collaborative_doc` (+ `send_attachment` 1) | **DOCUMENT** | **30** | Gotenberg / upload |
| `chat` | chat/text | 307 | — |

Requirement (≥2 voice + 1 code + 1 document SOURCES): **satisfied 77 / 65 / 30**. Public sim catalog by purpose-`type`: TRAINING 121 · ASSESSMENT 98 · HIRING 87 · INTERVIEW **1** (only ONE public interview sim → the interview content-story must pin that one). sim_tasks is snapshot-captured (public predicate + parent-scope) → already replayable.

## S1 — per-product result-route map + prove-by-render classification (code discovery)

### THE CENTRAL UNKNOWN — RESOLVED: simulation result = PERSISTED READ (seedable), NOT live-recompute
- Player route: `next-web-app/apps/web/src/app/(authenticated)/(verified)/sim/[slug]/result/[sessionId]/page.tsx`; manager sibling `.../sim/[slug]/[userId]/result/[sessionId]/page.tsx`; hiring mirror routes under `apps/hiring/...`. All render via `packages/ui/src/AISimulation/AISimulationResultContainer.tsx`.
- Render query `GET_SIMULATION_RESULT` / `jobSimulationResult(sessionId)` → resolver `jobsimulation/internal/graph/queries.resolvers.go:70` → plain Ent SELECTs of `validation_attempt_results.evaluation_status` (+ skill/criterion/check results, anticheat). **NO engine call, NO LLM grading, NO replay on render.** Grading is an async Asynq worker that runs at submit-time; the read path never invokes it.
- Behaviour: `evaluationStatus===Pending` → FE polls every 2s (perpetual spinner = blank if unseeded); any terminal status stops polling and renders. `passed` derived client-side (`evaluationStatus===Passed`). `useRecalculateEvaluationResult` is ONLY on the retry button (user-initiated), never on render.
- Corroboration: the existing `seed-verified-skill` skill already lights up this exact UI by direct INSERT; and DB shows completed sessions carry persisted score + result_status + full child fan-out.

### Route-classification matrix (player + manager), code-cited
| Product | Player route + read-model | Manager route + read-model | Classification |
|---|---|---|---|
| **Sim TRAINING** | `apps/web/.../sim/[slug]/result/[sessionId]` → jobsimulation validation/eval rows | `apps/web/.../@tabs/ai-simulations/[simId]` (+`[userId]`) → `insightsJobSimulationByMemberships` → `app.public.local_jobsimulation_sessions` MIRROR (`intelligence.go:1692`, gate `OrgFeatureInsights`) | **renders-from-seed** (both; manager needs co-written mirror + `OrgFeatureInsights` grant) |
| **Sim ASSESSMENT** | same player route | same manager route/mirror | **renders-from-seed** |
| **Sim HIRING** | `apps/hiring/.../sim/[slug]/result/[sessionId]` (`isHiring`, `HiringResult` org-setting kill-switch) | `apps/hiring/.../@tabs/ai-simulations/[simId]` (+`[userId]`) → same mirror. **Genuine hiring org EJECTS from apps/web → apps/hiring** (`UserStatusContext.tsx:168-169`) → served as 2nd UI container (M224) | **renders-from-seed** (blank iff org disables `HiringResult`) |
| **Sim INTERVIEW** | `.../sim/[slug]/result/[sessionId]` → `interviewExtractionUserReport` (`queries.resolvers.go:536`, gate `CheckSessionReadPermission`) | `.../@tabs/interviews/[simId]/[userId]` → `interviewExtractionManagerReport` (`queries.resolvers.go:563`, admin gate `OrgActionAssignmentsWrite`) | **renders-from-seed + FLAG-GATED** → effectively **needs-demo-patch**: both surfaces require `isExtractionEnabled = posthog.isFeatureEnabled('flag_interview_{player,manager}_report')` (`AISimulationResultContainer.tsx:499-506`). Seeded `interview_extraction_results` row is necessary but NOT sufficient — the PostHog flag must be ON in the demo (enable via demo PostHog bootstrap, or a demo-patch forcing the flag). |
| **Skill-path legacy** | `apps/web/.../skill-path/[skillPathId]/page.tsx` → `getOrCreateSkillPathSession` (skillpath runtime) — **get-OR-create auto-materializes a blank pending session** | `apps/web/.../@tabs/skill-paths/[skillPathId]` (+`[userId]`) → `insightsSkillPathByMemberships` → `app.public.local_skill_path_session` MIRROR (`intelligence.go:997/1142`, gate `OrgFeatureInsights`). **apps/hiring = no-surface** | Player: **runtime-computed-blank** (seedable via persisted skillpath `skill_path_session/chapter_session/step_session` rows); Manager: **renders-from-seed** (needs mirror row) |

Key structural correction: **NO Next.js intercepting routes exist** anywhere in apps/ (verified: 0 `(.)`/`(..)` dirs). The recruiter comparison "drawer" is a plain Ant `<Drawer>` (`InsightsByMembersContainer.tsx:359`) on the ordinary `[simId]/page.tsx` leaf. → hiring.md § M228 "intercepting-route-aware" is STALE (fixed inline).

### Manager-view eligibility (which products HAVE a manager result route)
- Sim TRAINING/ASSESSMENT/HIRING: YES (activity-dashboard ai-simulations tab → local_jobsimulation_sessions mirror).
- INTERVIEW: YES (interviews tab → manager_report), flag-gated.
- Skill-path legacy: YES in apps/web only (skill-paths tab → local_skill_path_session mirror); NO in apps/hiring.

### Doc-fidelity findings (audit)
- KB-1 (YELLOW): `jobsimulation.md` is SILENT on the session/result read-model + the mirror (content lives in seeding-spec.md/hiring.md). M231's `content-stories-routes.md` becomes the consolidation home; add a pointer from jobsimulation.md.
- KB-2 (YELLOW, tangential): `jobsimulation.md` ports "8400/8401" conflict with repo CLAUDE.md "8080/8081" (offset ports notwithstanding) — track, not load-bearing for M231.
- KB-3 (STALE, fixed inline): `hiring.md` § "The render probe is intercepting-route-aware (M228)" — no intercepting route exists; it's a plain Drawer on the leaf route. Corrected v2.5 M231.
- KB-4 (STALE/incomplete): `skillpath.md` omits that the manager insights surface reads the `app`-side `local_skill_path_session` MIRROR (same trap as hiring) — add a pointer.

## S4 — AI-labs feasibility + academy "session" verdict (code discovery)

### AI-LABS — VERDICT: OUT (no result-render surface today)
- `LabsAPIClient` is nil whenever `LABS_API_URL` unset (`app/main.go:462-465`). Create then persists a `lab_sessions` row via `idGen()` (12-char hex, **no VM**, status defaults `"booting"`, no `ide_url`/`preview_url`) — `app/internal/labs/session/manager.go:164-219`.
- `grade_result` (Ent col `lab_session.go:122-127`) is written only by `ReportEvent("grade")`, a call FROM labs-api back to app — nil client → no VM → never graded. AND the **GraphQL `LabSession` type exposes NO grade field** (`labs.graphqls:10-24`), so no FE can read it even if seeded directly.
- Only per-session page `/labs/[id]/page.tsx` reads LIVE from labs-api (`lib/labs-api.ts:81-83` throws if `LABS_API_URL` unset) → cannot render without a live worker. Dashboards `/labs` + `/enterprise/labs` read `mySessions`/`labSessions` (activity + spend) → a seeded row shows only as a `status="booting"` spend line, no result.
- Package path is `internal/labs/session` (not `internal/labsession` per backend.md — STALE).
- **Verdict: rule AI-labs OUT of the content-stories result-render matrix.** A lab "result" cannot render from a seed without wiring a live labs-api worker (out of the zero-platform-edit envelope; a grade GraphQL field would be a platform edit → escalate). The content-stories tab MAY still list a lab session presence-only in the activity/spend dashboards, but NOT as a played result.

### ACADEMY SESSION — VERDICT: IN (seedable server row; NOT presence-only)
- The corpus premise ("Clerk-only, no backend at runtime") is STALE. Since ant-academy **v0.5 "direct line" M2** the academy is backend-authoritative: `code/app/api/academy/beacon/route.js` posts `UPSERT_CHAPTER_PROGRESS`/`SET_LAST_ACTIVITY` over GraphQL to the platform `app` academy subgraph.
- Store lives in platform `app`: `app/internal/data/ent/schema/academy_chapter_progress.go` (unique `user_id+chapter_slug`) + `academy_last_activity`/`_chapter_time`/`_bookmark`/`_certificate`/`_feedback`. GraphQL: `academyProgress`/`academyLastActivity` queries; `upsertChapterProgress[Batch]`/`setLastActivity` mutations (`academy.graphqls`).
- **Purpose-built to seed:** `app/cmd/academy-seed/main.go` — fixtures `starter`/`in-progress`/`completed`, targeted by `--user-email`/`--user-id`, idempotent, seeds THROUGH the academy Manager.
- **Verdict: the academy "session" = the per-user `academy_chapter_progress` + `academy_last_activity` rows** → seedable → the academy content-product section can render REAL played progress. Caveat: progress is keyed by `chapter_slug`, decoupled from catalog rows → the chapters it points at still need catalog rows to render (the M230 academy-fill dependency). So the academy content-story DEPENDS ON M230's demo-fill (Fate-2, already in the release).

### S2 bridge — real sessions vs PUBLIC (already-replayable) sims
The clone must pin sessions whose `sim_id` is a **public-published** sim so the demo already has the sim content (snapshot-replayed). Completed org-scoped sessions whose `sim_id ∈ public-published`:

| sim_type | completed (org-scoped) | sim_id is public-published | distinct public sims |
|---|---|---|---|
| ASSESSMENT | 5,064 | 2,427 | 79 |
| TRAINING | 1,707 | 549 | 66 |
| HIRING | 1,679 | 395 | 36 |
| INTERVIEW | 488 | 41 | **1** (the sole public interview sim) |

→ Ample public-sim-anchored real sessions per product to pin (INTERVIEW limited to the 1 public interview sim). The M232 sourcing query MUST inner-join `directus.simulations` on the public predicate to guarantee `sim_id` resolves in the demo (else clone the private sim too — out of scope; prefer public-anchored sources).

## Modality taxonomy (S3 code-side, complements the directus catalog)
Modality is `SimFeature{Chat,Voice,Code,Doc,CollaborativeDoc}` + `InteractionMode{chat,call,code,send_attachment,collaborative_doc}` (proto `cms/v1/.../job_simulations.go`) + per-interaction `ActionType{email,chat_message,storage_upload,validation_request,call}`. `SimTask.TaskType` is a free string; the directus `sim_tasks.task_type` values used in content happen to carry the InteractionMode tokens (so the catalog query is valid). Backing jobsim entities: voice → `RealtimeCall`+`ChimeRecording` (LiveKit + Chime); code → `CodeSubmission` (**in-process Judge0 via `jobsimulation/internal/runner/`, formerly the standalone roadrunner service**); document → `CollaborativeAsset` + `storage_upload`; chat → `Interaction`/`ai_interaction`. `ai_architecture.md` code="via the Roadrunner service" is STALE (KB-5); `roadrunner.md` is orphaned (KB-6); `ant-academy.md` "no backend writes" is STALE (KB-7); `backend.md` labs package path STALE (KB-8).
