# M240 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [ ] **⚠️ HARD media-safety gate (R1) — MUST clear first** — fresh data-controller sign-off + `safety.md` §3.8 raw-media amendment + a voice/document anonymization decision, BEFORE any customer audio lands in a demo
- [ ] **Defect 1 (selection)** — constrain the public sim's type to the cell type in `sourcing.go`; re-pin `content-sessions.yaml`
- [ ] **Defect 3 (document)** — write the dropped `input_data` at seed time + port the document blob if `storage_upload`
- [ ] **Defect 2 (voice)** — port the Chime/S3 recording + re-host audio in the demo storage tier + flip `chime_status` available
- [ ] **Pass-rate (#4)** — score-band in `SelectionSpec` + tiebreak `score ASC`; re-capture
- [ ] **Delivers** — a new media-substrate spec under `corpus/ops/demo/` + a `safety.md` §3.8 raw-media amendment
