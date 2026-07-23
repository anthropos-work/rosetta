---
milestone: M240
slug: content-stories-fidelity
version: v2.6 "sound check"
milestone_shape: section
status: archived
created: 2026-07-20
last_updated: 2026-07-22
depends_on: M237
delivers: corpus/ops/demo/ (new media-substrate spec) + corpus/ops/safety.md §3.8 (raw-media amendment)
---

# M240 — content-stories fidelity

**Status:** `archived` (completed 2026-07-22)  ·  **Shape:** `section` (with a HARD media-safety gate)  ·  **Complexity:** large  ·  **Depends on:** M237

## Goal
The cockpit's claim matches the session — right type, playable call, visible document — at a believable pass rate.

## ⚠️ User decision baked in (2026-07-20) — media → PORT IT, behind a HARD internal PII gate
**Capture + re-host the Chime/S3 voice recording + document blobs** so the manager can hear the call / see the document. This **expands the customer-PII surface to raw audio + full documents** — a larger data-controller call than v2.5's scrubbed text; **a voice cannot be token-scrubbed.**

**HARD INTERNAL GATE — MUST clear BEFORE any customer audio lands in a demo:**
1. **Fresh data-controller sign-off** for raw customer voice + full documents in a demo.
2. **A `safety.md` §3.8 amendment** covering raw audio + full documents (extending the anonymized-real / VPN-scoped posture).
3. **A voice/document anonymization decision** — a voice cannot be token-scrubbed; the standing control is the VPN/tailnet scope (`safety.md` §3.8).

This is **R1 (blocks-quality)**. Likely consumes **DEF-M10-01** (S3 read access / media blob bytes).

## Scope
### In
- **Defect 1 (selection):** tighten `rext stack-seeding sourcing.go` to constrain the public sim's type to the cell type (exclude the interview sim from non-interview cells); re-pin `content-sessions.yaml`.
- **Defect 3 (document):** write the dropped `input_data` at seed time (`content_stories_write.go` / a content-specific criterion column set); + **port the document blob** if the body is a `storage_upload`.
- **Defect 2 (voice):** **port the Chime/S3 recording** — capture the recording reference + re-host the audio in the demo storage tier + flip `chime_status` to available.
- **Pass-rate (#4-feature):** add a score-band to `SelectionSpec` (`AND s.score BETWEEN 70 AND 95`), flip the tiebreak to `score ASC` (prefer lower), 100% only as fallback; re-capture.
- **HARD internal media-safety gate** (above) — before any customer media lands in a demo.

### Out
- The language toggle (M241).
- The cockpit-UX regroup (M242).

## Open questions / HARD gate
- Raw-media PII is a larger data-controller call than v2.5's scrubbed text; the internal gate (sign-off + safety amendment + anonymization decision) **must clear before any customer audio lands in a demo**. A voice cannot be token-scrubbed.
- Does DEF-M10-01 (eu-west-1 S3 read access) land here as part of the media port?

## Note
The re-capture needs prod read (`~/.pgpass`).

## Delivers
A **new media-substrate spec** under `corpus/ops/demo/` + a `corpus/ops/safety.md` §3.8 amendment (raw media).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
