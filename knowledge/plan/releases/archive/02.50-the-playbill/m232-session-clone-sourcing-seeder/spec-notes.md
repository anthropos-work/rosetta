# M232 — Spec notes

## Pre-flight audits — S1 (Prod-session sourcing + fixture)
- **Phase 0b KB-fidelity: GREEN** — `kb-fidelity-audit.md`. All topics PAIRED (session-clone-spec.md is the
  milestone's own deliverable, not a blind area). Contract `content-stories-routes.md` re-verified against live
  code + prod DB: net-new schemas, modality catalog, sourcing pools, interview resolver (`:536`/`:563`), the
  mirror, `internal/runner/` all ALIGNED. 0 stale claims, 10/10 cross-refs resolve.

## Topic → doc → code triples (fast-start for later audits)
- content-stories contract → `corpus/ops/demo/content-stories-routes.md` → jobsim resolvers + next-web render + prod DB
- 7-table fan-out + mirror → `stories-spec.md`/`seeding-spec.md` → `stack-seeding/seeders/{persona,persona_write,hiring_funnel}.go`
- content-ref pinning → `contentref.go` (reserved-tail, type-aware `readHiringSimPool`) — extend to type×modality
- seeder framework → `seeder/seeder.go` (Conn: CopyRows/CopyRowsIdempotent/Exec/QueryRow; Seeder iface), registry `cmd/stackseed/main.go:buildRegistry` (register before line 454)
- fake Conn for tests → `seeders/seeders_test.go` `recordingConn` + `findCopy` + `isolation.NewAuditLog()`
- isolation → `isolation/{isolation,audit}.go` (PerStackIsolated, AuditLog.Record(Entry{...Allowed:true}))

## Net-new result-substrate table schemas (MCP-introspected, live prod)
- `jobsimulation.actors`: id(uuid,gen), user_id(uuid,null), username(varchar NN — **PII name**), alias(varchar NN — **PII**), role_key(varchar NN), stakeholder(bool def false), session_id(uuid NN)
- `jobsimulation.interactions`: id(uuid,gen), timestamp(tz NN), action_type(varchar NN — enum), action_payload(jsonb NN — **transcript**), session_id(uuid NN), target_id(uuid NN → actor), source_id(uuid NN → actor)
- `jobsimulation.validation_check_results`: id, created_at, updated_at, check_id(uuid NN), engine(varchar NN), parameters(jsonb NN), success(bool NN), feedback(varchar NN), validation_criterion_result_id(uuid NN, FK), essential(bool def false)
- `jobsimulation.interview_extraction_results`: id, created_at, updated_at, user_report(jsonb NN), manager_report(jsonb NN), session_id(uuid NN), summary(jsonb null)

## Existing fan-out column helpers (persona_write.go / assignments.go) — REUSE
- `sessionCols()`, `localSessionCols()` (mirror; misspelled `completition_status`), `attemptResultCols()`,
  `skillResultCols()`, `criterionResultCols()`, `userSkillCols()`, `assignmentCols()`, `localSkillPathSessionCols()`,
  `orgAssignmentSessionCols()`, `interviewCols()` (succession — minimal). NET-NEW cols needed: actors,
  interactions, validation_check_results, + full interview user_report/manager_report.

## Sourcing (S1) — public-anchored pools per (type × modality × pass/fail), live prod (MCP)
Public predicate: inner-join `directus.simulations` on `private=false AND tenant_id IS NULL AND status='published'`.
Modality via `directus.sim_tasks.task_type` ∈ {call=VOICE, code=CODE, collaborative_doc/send_attachment=DOCUMENT}.
- ASSESSMENT×call: passed 549 / failed 1049 · ASSESSMENT×code: passed 25 / failed 22 · ASSESSMENT×collab_doc: failed 6
- TRAINING×collab_doc: passed 6 / failed 21 (the passed-document source) · TRAINING×code passed 63 · TRAINING×call p81/f63
- INTERVIEW×call: passed 35 / failed 5 (voice-graded; the ONE public interview sim) · HIRING×call/code present
Pin by `jobsimulation.sessions.id`. Owner-is-player by construction (re-owned to a seeded player member).

## S3 build input — INTERVIEW report is PLAN-DRIVEN (agent-verified, next-web-app)
- Flag gate `AISimulationResultContainer.tsx:498-506`: `isManagerReportEnabled = isManager && isFeatureEnabled('flag_interview_manager_report')`; `isPlayerReportEnabled = !isManager && !hideResult && isFeatureEnabled('flag_interview_player_report')`; `isExtractionEnabled = isInterview && (mgr||player)`. Both flag literals load-bearing.
- Report wire shape: opaque jsonb string → `JSON.parse` into `ExtractionData { results: Record<string, Record<string, unknown>> }` (`InterviewReport/types.ts:65`). Envelope: `{"results": {"<sectionId>": {"<field>": value, "guessed_fields": [], "limited_data": false}}}`.
- **Plan-driven:** UI iterates a SEPARATE CMS `ExtractionPlan` (`simulationExtractionSchema`) and looks up `results[sectionId]`. Believable render needs the seeded section-ids to match the replayed plan; one section MUST have `layout:null` + a `score_grade` field value ∈ {"A","B","C"} (header quality badge). Backend does NO struct enforcement (raw LLM bytes stored) — so the row is insertable with any valid JSON; render fidelity (exact plan match) is M234 coverage's concern. M232 seeds a structurally-valid, believable envelope (score_grade section + narrative sections), capturing the public interview sim's plan section-ids into the fixture where cheap.
- `user_report` = player+unscoped sections; `manager_report` = manager+unscoped (superset, incl. `attention_points`); `summary` = flat JSON or NULL.

## S1 — the 9 PINNED prod sessions (source-pins; deterministic reseed)
Selected via the public-anchoring + non-manager + per-cell query, richest-fan-out-first. `sessions.id → sim_id`:
| key | source_session_id | sim_id | type | modality | passed | score | dur_s | actors | inter |
|---|---|---|---|---|---|---|---|---|---|
| asmt-voice-pass | a9f78ff0-3a10-4d9d-a090-36c52ad8b3a2 | 6d6cdf39-e043-4f94-8a5c-e97116bfe1b2 | ASSESSMENT | voice | yes | 100 | 531 | 2 | 1 |
| asmt-voice-fail | 34d39116-3bfc-436f-b0e8-4560b3b15ddb | f7bc7df0-0303-4914-ba85-4f8731885644 | ASSESSMENT | voice | no | 59 | 668 | 3 | 18 |
| asmt-code-pass | 905dff41-50f7-4278-9693-8081db9436f8 | e0ae482f-3deb-4bfe-8dc8-a4ae4362d9d5 | ASSESSMENT | code | yes | 100 | 1800 | 2 | 25 |
| asmt-code-fail | 4d9f8123-5f76-416e-a474-10815583f269 | 35a191f9-35be-46a3-990f-a5bbc6524f9c | ASSESSMENT | code | no | 57 | 1512 | 2 | 9 |
| asmt-doc-fail | 8b6aba5b-ca3b-4b04-9050-578d66f6d208 | 280f9476-7677-4839-8e7e-4eadf6b343ba | ASSESSMENT | document | no | 47 | 3600 | 3 | 105 |
| train-doc-pass | 886b13b2-1546-4a1d-81a8-6b0e82ce842b | b21a323f-fe25-4702-a3d9-73944ffd9759 | TRAINING | document | yes | 100 | 2482 | 4 | 30 |
| train-chat-fail | 0102476c-a3ae-4417-a123-bcada93f8b3e | 1026cd0c-9531-4601-8d76-06a7aae8dc99 | TRAINING | chat | no | 59 | 2703 | 3 | 51 |
| hire-voice-pass | d25d937b-5f29-4e64-b695-e5cfaff9973c | a9a93976-7c3a-4524-8f86-77c30c3e633d | HIRING | voice | yes | 100 | 971 | 2 | 2 |
| intv-voice-pass | 39c41c70-7056-451c-95e2-8f106a4f9608 | 6d6cdf39-e043-4f94-8a5c-e97116bfe1b2 | INTERVIEW | voice | yes | 100 | 435 | 2 | 1 |
Assessment set = 2 voice + 2 code + 1 document (satisfies "2 voice + 1 code + 1 document"). passed+not-passed both.
NB the sim_id is the REAL public sim (renders the real sim header). The interview + one assessment share public sim 6d6cdf39.

## S3 modality shapes (agent-verified, live prod enum distributions)
- **interactions.action_type**: DB carries ONLY `email` (309744) + `call` (7382). Write email/call only (COPY bypasses Ent — an invalid value would insert-but-be-invisible, the G14 class). VOICE→`call` (actions.Call: elevenlabs_agent_id, elevenlabs_conversation_id, audio_ref{ObjectRef}, audio_duration, status active|ended|userended|error, transcript, transcript_summary, transcript_raw). CHAT/DOCUMENT-upload→`email` (actions.Email: subject, body{text,html}, attachment_refs[ObjectRef], audio_ref, metadata). interactions FK source_id/target_id → actors; CHECK source_id<>target_id.
- **actors**: username/alias (PII → anon identity), role_key(free-text, high-cardinality — synthesize per stakeholder), stakeholder(bool), session_id. NO timestamps (PK mixin only). unique(username,session), unique(user_id,session).
- **CODE**: `code_submissions` (FK session only): status=completed, token(uuid unique), runtime∈{py,java,cpp,go,...}, source_code(base64 zip), stdin, language_id, stdout/stderr/compile_output/time/memory/message, is_test, session_id. Code graded via `collaborative_asset` input_format (editor-content diff), NOT a 'code' input_format (no such enum value).
- **DOCUMENT**: `collaborative_assets` (FK session only): owner_id, asset_id, runtime, content, original_content, file_key, run_instructions(col run_instructions), is_test, session_id. input_format=`collaborative_asset` (editor diff) OR `text_document` (uploaded attachment → {"text_document":"..."} input_data). Upload persisted as an `email` interaction w/ attachment_refs (NOT storage_upload).
- **input_format** enum: chat, collaborative_asset, text_document, call, ai_assistant (NO code). input_data jsonb populated only for text_document ({"text_document":...}); chat proto shape {"chat":...} valid but live-null.
- **G14-valid enums confirmed**: var.evaluation_status passed|failed (the gate), acceptance_status passed|failed, vasr.status=completed, crit.type=evaluation, crit.status=completed, crit.input_format as above, code_submission.status=completed.

## Design decisions (S1/S2/S3 — anonymize-by-construction)
1. **Anonymize BY CONSTRUCTION**: free-text (feedback/input_data/transcript/actor names/reports) SYNTHESIZED, never copied. Sourced = only NON-PII skeleton (source_session_id pin, sim_id, sim_type, modality, score, passed, duration, actor/interaction counts). ⇒ fixture is provably PII-free.
2. **Owner = an existing player-vantage population member** of the target org (real public.users row + player membership → non-manager by construction). Resolves the mint-vs-reuse open question toward REUSE (each must map to a real seeded users row; reuse trivially satisfies it + is more believable than a thin minted seat). Distinct non-hero member per session (keeps hero dashboards clean).
3. **Target org** = the first NON-HIRING EffectiveStory (has player members). Cloned sessions appear in that org's activity (believable; a few extreme scores don't skew a 220-member org). Dedicated-org isolation is an M234 open question.
4. **Fan-out reconstructed** from type-appropriate shape (per M231 doc: ASSESSMENT/TRAINING/HIRING ~1 var + 1-3 skill + 3-5 criterion; INTERVIEW +1 interview_extraction_results) with REAL public taxonomy node-ids drawn from the demo's replayed taxonomy (resolve-or-drop → closure green, never fabricate). Transcript capped (~12-16 interactions) regardless of source count.
5. **Fixture** = go:embed'd curated YAML in a `contentsession` package (the exhibit set is a fixed, source-pinned, code-owned artifact, like the AI-readiness prompts — not per-stack config). Seeder loads the embed; M233 projects the pins into seed-generation-manifest.yaml.
6. **INTERVIEW report** = plan-shaped `{"results":{...}}` envelope with a `layout:null`/`score_grade` "A"/"B"/"C" section + narrative sections; backend does no struct enforcement so it's insertable; exact plan-section match is M234 coverage's concern.
