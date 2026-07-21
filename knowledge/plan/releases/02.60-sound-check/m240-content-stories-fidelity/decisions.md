# M240 — Decisions

_(implementation choices with rationale accumulate here during build)_

- **User decision (2026-07-20): media → PORT IT.** Capture + re-host the Chime/S3 voice recording + document blobs. Gated by a HARD internal PII gate (fresh data-controller sign-off + a `safety.md` §3.8 raw-media amendment + a voice/document anonymization decision) that MUST clear before any customer audio lands in a demo — a voice cannot be token-scrubbed. Likely consumes DEF-M10-01. Recorded at design time; carried here for build traceability.

## HARD media-safety gate — CLEARED by explicit user decisions (2026-07-21)

The HARD internal media-safety gate (R1) is **CLEARED** by the following explicit user decisions, handed down 2026-07-21 via the orchestrator. These are the authoritative gate-clearing record; the `safety.md` §3.8 amendment + the fresh dated data-controller sign-off (2026-07-21) MUST land BEFORE the seeder copies any customer audio into a demo (the gate: no customer media in a demo until the posture is documented).

- **Voice → PORT AS-IS; VPN/tailnet scope is the control.** A voice cannot be token-scrubbed; the standing control is the VPN/tailnet scope (`safety.md` §3.8). **NEW HARD CONSTRAINT: the ported voice's apparent GENDER must MATCH the demo persona that owns the re-tenanted session.** The `ContentStorySeeder` re-tenants a real session onto a `content-player-<idx>` owner slot; when that session has a voice recording, the owning persona's (generated account) gender MUST match the voice — no "Maria plays a male voice" mismatches. Detect/label the source voice's gender **at capture time, values-blind, via the tool** (a gender label is not the PII we scrub; derive it in-tool, never listen to audio in-context), and constrain the persona pairing so gender aligns. Believability + coherence requirement.
- **Documents → PORT + SCRUB.** Port the real document body through the SAME best-effort name/PII scrub the transcripts already use (`scrub` package), then re-tenant.
- **Both:** land a `safety.md` §3.8 amendment covering raw audio + full documents (extending the anonymized-real / VPN-scoped posture — audio is NOT token-scrubbable, the control is VPN/tailnet scope) + record a **fresh dated data-controller sign-off (2026-07-21)** here + in the new media-substrate spec. Residual re-identification risk is real and **ACCEPTED by the data-controller (2026-07-21)**; the control is the VPN/tailnet scope.
- **Genuine blockers (surface, don't fake):** S3 media-read access not available (defect 2 can't port the blob — DEF-M10-01 lands here); the 70–95 pass band empty for a requirement cell (can't re-pin believably); the document body is an unreachable blob. These are real "the user must provide X" stops, distinct from the safety gate (which is cleared).

## Phase 0b — KB-fidelity findings (tracked, YELLOW; report `kb-fidelity-audit.md`)
- **KB-1** — `session-clone-spec.md §3–4` claims `validation_criterion_results.input_data` is copied AND replayed; code captures it but the shared `criterionResultCols()` drops it at seed time. STALE; resolved by **Defect 3** (content-specific criterion col set) + a doc bump in Phase 5.
- **KB-2** — `session-clone-spec.md §2` documents only public-anchoring + non-manager-played; no cell-type-match predicate → the interview sim leaks into non-interview voice cells (CQ-1). Resolved by **Defect 1** + a doc bump.
- **KB-3** — media substrate (voice recording + document blob) is a BLIND-AREA, PROMOTED to the overview `Delivers` line (new media-substrate spec + `safety.md §3.8` raw-media amendment). Acceptable per the gate rule.

## Defect 3 investigation — the document body is NOT an S3 blob (2026-07-21)
- Prod-verified (non-PII structural): the three DOCUMENT pins are all `collaborative_doc`. Their body lives in `validation_criterion_results.input_data` under the `text_document` key (input_format `text_document`): asmt-doc-pass **6226 chars** (a real body), asmt-doc-fail 57, train-doc-pass 21. All three have **ZERO `collaborative_assets`**, and jobsimulation has **no `storage_upload`/attachment/file table** — the only doc-blob table is `collaborative_assets`.
- **CONSEQUENCE: Defect 3 is FULLY landable on prod-read alone — no S3 blob to port.** The overview's "if the body is a storage_upload" speculation does NOT apply. The fix is purely the seed-time write (content-specific `contentCriterionResultCols()` + write the placeholder-filled `input_data`), which content-capture already captured + scrubbed. Landed; `TestContentStorySeeder_WritesInputData` fences it. This RESOLVES the "unreachable doc blob" blocker candidate — it is not a blob.

## Defect 1 + pass-rate prod checks (2026-07-21, non-PII)
- Interview-sim discriminator: `directus.simulations.type` = the sim's OWN type; interview sim is the sole `type='SIMULATION_TYPE_INTERVIEW'` (n=1). Fix = `AND d.type = <cell sim_type>` (robust, not slug).
- **Pass-rate 70–95 band is POPULATED for every passed cell** (assessment 253, code 31, doc 12, hiring 31, interview 28, training-doc 4) → the "empty pass band" blocker does NOT fire; the pass-rate feature is landable.

## Pass-rate re-pin picks (identified 2026-07-21; RESUME here — re-pin YAML + `content-capture --only <keys>` + regenerate presets)
CODE landed (rext 0753e48). The fixture re-pin was reverted (machine paused before re-capture). The 5 remaining still-100 passed cells re-pin to these prod-verified 70-95 distinct sessions (non-PII; all richer/believable). `asmt-doc-pass` is already 84 (no re-pin). `asmt-voice-pass` already re-pinned in Defect 1 (score 70).

| key | source_session_id | sim_id | sim_slug | score | dur | actors | inter |
|---|---|---|---|---|---|---|---|
| asmt-voice-pass-2 | e0507f81-e0cb-4075-aded-5d953ccf5fe5 | cbe85f54-68ae-4f4e-bc8c-62f7ffe31705 | rebuild-manager-trust-after-customer-mistake-69f | 74 | 1678 | 3 | 52 |
| asmt-code-pass | e70c5935-2f6a-4b58-bc85-a83a59ae2e73 | 634b9ffd-a6a8-444a-a585-1867c1dc61f4 | who-can-see-this-document-fc0 | 72 | 2650 | 3 | 32 |
| train-doc-pass | 5ef46e63-a1e3-4c14-a601-cd3d878c7173 | 55014cca-6b86-464f-ad3b-9ed501e1e29e | manage-the-migration-to-microsoft-azure-of-a-global-ecommerce | 82 | 5013 | 5 | 167 |
| hire-voice-pass | dbedcc6b-8505-4cf6-a826-ca9cb4bd183f | cfc8f8c6-5973-4414-b2b6-5b5a98e1ce15 | marketing-performance-analysis-and-strategy-proposal-exp-53b | 83 | 445 | 3 | 20 |
| intv-voice-pass | cba53b09-5a94-4adf-850f-c0e54aafbc82 | 6d6cdf39-e043-4f94-8a5c-e97116bfe1b2 | ai-readiness-interview-d62 | 81 | 424 | 2 | 1 |

Resume steps: (1) apply these 5 pins to `content-sessions.yaml`; (2) `go run ./cmd/content-capture --dsn "postgres://marco_read@<pgpass-host>:5432/postgres?sslmode=require" --only asmt-voice-pass-2,asmt-code-pass,train-doc-pass,hire-voice-pass,intv-voice-pass` (counts-only; leak post-condition must pass — if a session leaks, pick the next candidate via the sourcing query); (3) regenerate BOTH canonical presets (`--content-export` → content-manifest.json; `--manifest-export --gen-seed …` → seed-generation-manifest.yaml, preserving the leading comment header); (4) update the `content_manifest_test.go` asmt-voice-pass expectations if any projected path changed; (5) run seeders+contentsession+stackseed suites. Then the media-substrate spec (§Delivers) remains.

## Defect 2 (voice recording) — GENUINE BLOCKER, and the diagnosis is DEEPER than "S3 creds" (2026-07-21)
Prod-verified (non-PII structural), the media port is blocked for THREE independent reasons:
1. **The 7 pinned voice sessions have NO recording.** All have `chime_status='not_available'`, zero `chime_recordings` rows, no `bunny_video_id`, no `media_pipeline_id`. The demo's `not_available` is **FAITHFUL to the source** — there is no code defect for the current pins, and nothing to port.
2. **The media is on Bunny.net CDN, not prod S3.** `ChimeRecording.bunny_video_id` → the frontend `BunnyVideo` plays from `vz-a5a3c33b-037.b-cdn.net` (a Bunny library). So **DEF-M10-01 (eu-west-1 S3 read) is the WRONG/insufficient access** — the served media needs **Bunny.net API access** (to re-host) which is ALSO not provisioned here. (The raw Chime pipeline output may also be in eu-west-1 S3 — still blocked.)
3. **Recorded voice sessions are almost only HIRING** (public-anchored: hiring 13 passed / 22 failed with a bunny video; assessment/training/interview = 0), and **ZERO** sessions in the whole public-anchored voice pool have the render-gating status. The frontend gate is `chimeStatus === 'completed'` (not "available") **AND** a resolvable `bunny_video_id`.

**To land the port, the user must decide + provide:** (a) media access — **Bunny.net API creds** (to download+re-host into the demo storage tier, the overview's intent) and/or eu-west-1 **S3** Chime-output read; (b) accept **re-pinning the voice cells to recorded HIRING sessions** (the only recorded pool) — which serves REAL customer interview **video** (a person's face+voice), the most sensitive/unscrubbable media, under the §3.8.1 raw-media posture (VPN-scoped, signed off 2026-07-21). **SEVERITY=blocker.** The safety posture + the gender-coherence contract are LANDED (§3.8.1 + the media-substrate spec) so the port is pre-blessed the moment access is provided; building the port code now — against media it cannot read or test — would be untestable scaffolding (three-fate rule forbids it).

## Environment: S3 media-read UNAVAILABLE (2026-07-21) — DEF-M10-01 is LIVE here
- `~/.aws/credentials` is 0 bytes; no AWS env/profile; `aws sts get-caller-identity` → "Unable to locate credentials." Prod-read (`marco_read`) IS available.
- Consequence: the **media-blob PORT** (Defect 2 voice audio bytes; Defect 3 doc body IF an S3 `storage_upload`) is **BLOCKED — the user must provide eu-west-1 S3 read creds**. This is the genuine blocker the roadmap pre-flagged (R1's likely DEF-M10-01 consumption), NOT the safety gate (which the 2026-07-21 decisions CLEAR). Everything else (safety amendment, Defect 1, Defect 3 seed-time write, pass-rate) is landable on prod-read alone and WILL be landed this session.
