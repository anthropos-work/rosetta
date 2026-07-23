# M240 — Spec notes

Topic → doc → code triples + media-port / selection / pass-rate findings accumulate here during build.

## Defect 1 — selection (wrong type)
- Tighten `rext stack-seeding sourcing.go` to constrain the public sim's type to the cell type (exclude the interview sim from non-interview cells); re-pin `content-sessions.yaml`.

## Defect 3 — document (dropped input_data + blob port)
- Write the dropped `input_data` at seed time (`content_stories_write.go` / a content-specific criterion column set).
- Port the document blob if the body is a `storage_upload`.

## Defect 2 — voice (Chime/S3 recording port)
- Capture the recording reference + re-host the audio in the demo storage tier + flip `chime_status` to available.

## Pass-rate (#4-feature)
- Add a score-band to `SelectionSpec` (`AND s.score BETWEEN 70 AND 95`), flip the tiebreak to `score ASC` (prefer lower), 100% only as fallback; re-capture.

## HARD media-safety gate (R1) — before any customer media lands in a demo
- Fresh data-controller sign-off · `safety.md` §3.8 amendment (raw audio + full documents) · voice/document anonymization decision (a voice cannot be token-scrubbed; control = VPN/tailnet scope).
- Likely consumes DEF-M10-01 (S3 read access).

## Deliverable — new media-substrate spec (corpus/ops/demo/)
- The capture + re-host mechanism for Chime/S3 voice + document blobs into the demo storage tier.

## Pre-flight audits — M240 (all sections; first-section fresh run)
- **Phase 0b KB-fidelity: YELLOW** (proceed with tracking). Report: `kb-fidelity-audit.md`.
- Triples verified:
  - selection → `session-clone-spec.md §2` / `content-stories-routes.md §3.5` → `contentsession/sourcing.go`
  - document → `session-clone-spec.md §3–4` → `seeders/content_stories_write.go` + `persona_write.go::criterionResultCols`
  - capture+scrub → `session-clone-spec.md §3` → `cmd/content-capture/main.go` + `scrub/` + `contentsession/content.go`
  - modality substrate → `session-clone-spec.md §4` → `seeders/content_stories_modality.go`
  - safety → `safety.md §3.8` → read boundary + scrub
  - media (voice/blob) → BLIND-AREA promoted (Delivers: media-substrate spec) → `jobsimulation` Recording/ChimeRecording ent + prod S3
- Findings KB-1 (input_data dropped at seed = Defect 3), KB-2 (cell type-match unenforced = Defect 1), KB-3 (media blind-area promoted). All in-scope; see decisions.md.

## Environment probe (2026-07-21)
- prod-read `marco_read` via `~/.pgpass` over Tailscale: **AVAILABLE** (psql not installed → use `mcp__postgres__query` for non-PII, capture tool `--dsn` for content).
- **S3 media-read (eu-west-1): NOT AVAILABLE** — `~/.aws/credentials` 0 bytes, no AWS env/profile, `aws sts get-caller-identity` fails. → **DEF-M10-01 genuine blocker** for the media-blob port (Defect 2 audio; Defect 3 doc body IF S3 storage_upload).

## Recording schema (Defect 2, from `jobsimulation/internal/ent/schema/recording.go`)
- `Recording{recording_id uuid, session_id uuid (unique)}` — the recording REFERENCE row (Postgres, capturable via prod-read).
- Also `ChimeRecording`, `RealtimeCall{call_id, status, session_id, started_at, completed_at, …}` entities.
- The audio BYTES live in prod S3 (referenced by `recording_id`) → **S3-read required to port the blob** = blocked.

_(will accumulate topic → doc → code triples during build)_
