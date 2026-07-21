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

## Environment: S3 media-read UNAVAILABLE (2026-07-21) — DEF-M10-01 is LIVE here
- `~/.aws/credentials` is 0 bytes; no AWS env/profile; `aws sts get-caller-identity` → "Unable to locate credentials." Prod-read (`marco_read`) IS available.
- Consequence: the **media-blob PORT** (Defect 2 voice audio bytes; Defect 3 doc body IF an S3 `storage_upload`) is **BLOCKED — the user must provide eu-west-1 S3 read creds**. This is the genuine blocker the roadmap pre-flagged (R1's likely DEF-M10-01 consumption), NOT the safety gate (which the 2026-07-21 decisions CLEAR). Everything else (safety amendment, Defect 1, Defect 3 seed-time write, pass-rate) is landable on prod-read alone and WILL be landed this session.
