# work-milestone M240 — pause/resume handoff

**Paused:** 2026-07-21 (user paused the machine mid-M240-build).
**Resume with:** `/developer-kit:work-milestone for M240` — build-milestone resumes from `progress.md`'s
per-section checklist (done sections stay committed; unchecked sections re-run).
**Branch:** `m240/content-stories-fidelity` (off `release/02.60-sound-check`).

## State at pause

**Pipeline position:** BUILD, mid-flight. Harden + close NOT yet run.

**Sections done + COMMITTED (safe):**
- ✅ **HARD media-safety gate** — `safety.md` §3.8.1 raw-media amendment + fresh dated data-controller
  sign-off (2026-07-21) landed BEFORE any customer media copied. Corpus commit `7020c3f`.
  **User gate decisions (2026-07-21):** voice → **PORT AS-IS, VPN-scope is the control, + the ported
  voice's gender MUST match the generated/demo persona that owns the session**; documents → **PORT + SCRUB**
  (same scrub as transcripts).
- ✅ **Defect 1 (selection)** — `sourcing.go` now constrains the public sim's `d.type` = the cell sim_type
  (robust, not slug-based); re-pinned `asmt-voice-pass` to a real assessment-voice session (score **70** —
  already in the 70–95 band). rext `9e8305a`, corpus `7226a3c`.
- ✅ **Defect 3 (document)** — the dropped `input_data` is now written at seed time via a content-specific
  `contentCriterionResultCols`. **Blob investigation RESOLVED: the document body is
  `input_data.text_document` (a collaborative_doc), NOT an S3 `storage_upload` blob → no blob to port,
  fully landable.** rext `cb64ccd`, corpus `92ae3ed`.

**GENUINE BLOCKER (needs USER decision on resume) — the diagnosis is DEEPER than "S3 creds":**
- ⛔ **Defect 2 (voice recording port)** — recorded at corpus `a20b85c`; full diagnosis in the build's final
  report + `decisions.md`. Three findings that reshape the decision:
  1. **The 7 currently-pinned voice sessions have NO recording at all in prod** — `chime_status='not_available'`
     is FAITHFUL to the source (no code bug for the current pins). Nothing to port from them.
  2. **Recordings live on Bunny.net CDN** (`bunny_video_id`), **not prod S3** — so DEF-M10-01 (S3 read) is the
     WRONG/insufficient access. Bunny.net API access is needed and is also unprovisioned here (`~/.aws/credentials`
     empty, no AWS identity resolves).
  3. **Recorded voice sessions exist almost only on HIRING sims — i.e. real candidate interview VIDEO** (face +
     voice), the most sensitive media in the platform. The render gate is `chime_status='completed'` + a
     resolvable `bunny_video_id`.
  **So the real decision is bigger than "port a voice recording":** to land playable media you must (a) provide
  **Bunny.net (and/or eu-west-1 S3) media access**, AND (b) decide whether to **re-pin the voice cells to recorded
  HIRING sessions — which serves real customer interview VIDEO** (a materially weightier data-controller call than
  the audio-only "port as-is" that was approved 2026-07-20). **Recommended framing for the user on resume:**
  either (i) provide Bunny.net access + consciously approve serving real candidate interview video (fresh
  sign-off — this is more than the audio decision covered), or (ii) **fall back to voice presence-only** (call
  metadata + scrubbed transcript + honest "recording not available in demo" — the 4th design option), which needs
  no media access and no video-PII call. The safety posture + gender-coherence contract are already landed, so a
  port is pre-blessed for AUDIO the moment access arrives — but the video reality warrants an explicit re-decision.

**Sections PENDING (not started / in progress at pause):**
- ⏳ **Pass-rate (#4)** — add a score-band to `SelectionSpec` (`AND s.score BETWEEN 70 AND 95`) + flip
  tiebreak to `score ASC` (prefer lower), 100% only fallback; re-capture. **Needs a prod pool-count check
  first** (confirm each requirement cell has a 70–95 session) + prod read (`~/.pgpass`). NOTE: Defect 1
  already re-pinned one session to score 70, so this is partially demonstrated.
- ⏳ **Delivers** — the new media-substrate spec under `corpus/ops/demo/` (the §3.8 safety amendment is
  already landed as part of the gate).

## Resume checklist
1. **Check both trees clean first.** If `rext` or the corpus branch is dirty from the interrupted pass-rate
   edit (rext showed 2 dirty at pause), INSPECT before proceeding — do not blind-commit. The safe committed
   HEADs at pause: corpus `a20b85c`, rext `cb64ccd`.
2. **Resolve the Defect-2 voice blocker** (user decision above) — that gates whether voice ports or goes
   presence-only.
3. Re-run `/developer-kit:work-milestone for M240` — it resumes the build from the checklist, then harden +
   close as normal. On close, re-pin the consumed rext tag(s) to the hardened HEAD (the M237/M238/M239
   precedent) + push rext.
4. **PII discipline still applies:** customer media flows prod→scrub→fixture through TOOLING, never into an
   agent's context; values-blind on any creds; never commit raw unscrubbed media or a secret.

## Context a fresh session needs
- Two-repo model: fix code in `rext` (`.agentspace/rosetta-extensions`, `main`); docs+records on the rosetta
  `m240/…` branch. Zero platform-repo edits.
- `state.md` was ~15,014 B at M239 close (98% of the 15 KB cap) — M240's close must trim, not grow it.
- The 8 (now 9) standing demo-stack test failures + `DEF-M239-01` are owned by M244 (just added to its
  In-list, commit on the release branch before M240 started).
