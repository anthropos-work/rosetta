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

_(will accumulate topic → doc → code triples during build)_
