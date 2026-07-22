# M240 — Retro

## Summary
M240 was the **third post-barrier fix** in v2.6 "sound check" (content-stories fidelity), a `section` milestone with a
**HARD media-safety gate cleared first**. It made the cockpit's claim match the session at a believable pass rate:
**Defect 1 (selection)** — a `d.type = <cell sim_type>` **public-CTE** predicate stops the sole public interview sim from
leaking into non-interview voice cells (CQ-1); **Defect 3 (document)** — the dropped criterion `input_data` is written at
seed time via a content-specific `contentCriterionResultCols` (the body is inline `text_document`, **not** an S3 blob, so
there is nothing to port); **pass-rate** — a `ScoreMin/ScoreMax` band + `score ASC` tiebreak with the 5 still-100% cells
re-pinned to real 70–95% sessions, re-captured values-blind. **Voice (Defect 2) → DELIVERED = voice presence-only** (user
decision 2026-07-22): the faithful `chime_status='not_available'` state IS the v2.6 deliverable. The **real-video exhibit**
is fully documented + **pre-blessed** (exhibit-**by-reference**: `bunny_video_id` + a read-only Bunny CDN signing key, no
media byte ever moves) but the Bunny recording signing keys are genuinely absent from this box's dev-stack, so it routes to
**M244** (`DEF-M240-01`, Fate-3, user pre-approved). **Delivers:** new `media-substrate-spec.md` + `safety.md` §3.8.1
(raw-media/VIDEO amendment + fresh 2026-07-21 data-controller sign-off) + `session-clone-spec.md` §2/§4 bumps. Complete
6-section close; **0 platform-repo edits**; **customer media never entered agent context.**

## Incidents This Cycle
- **None.** Harden (1 pass) mutation-verified all 3 fixes and found **0 bugs inline** — every fix was already correct at
  build. The close's code-quality + adversarial + doc + test reviews found **0 fix-required findings**. Flake gate 5/5
  (one transient "0/5 FAIL" during the close was a **zsh word-split false alarm** in the harness command — unquoted
  `$PKGS` was passed as a single arg, `matched no packages`; re-run with literal paths → 5/5 green; not a product flake).
- **No regressions.** M240 touched only the rext `stack-seeding` module (not the demo-stack test harness), so the standing
  8/9 host-state demo-stack failures do not reproduce from M240's work — 0 M240 regressions.
- **One latent risk pinned (not a live bug):** `append(criterionResultCols(), "input_data")` would alias the shared slice
  if it ever gained spare capacity — today `cap==len` (fresh literal), so it is safe, and the harden pass added an
  append-aliasing probe to fence it.

## What Went Well
- **The gate ordering held.** The media-safety posture (`safety.md` §3.8.1 + the fresh 2026-07-21 VIDEO sign-off + the
  gender-coherence contract) landed **before** any media-exhibit code, exactly as the gate rule requires — so when the
  media port turned out to be key-blocked, the honest fallback was already fully specified.
- **Honest scoping beat manufactured completion.** The real-video exhibit could not be provisioned or playback-verified
  on this box (Bunny keys genuinely absent, verified values-blind). Rather than flip `chime_status='completed'` and ship a
  broken 500 player, the milestone shipped the faithful `not_available` state as the deliverable and routed the exhibit to
  M244 as one atomic unit — the three-fate rule's "no untestable scaffolding" discipline applied correctly.
- **The DEF-M10-01 speculation was resolved, not carried.** The document body was proven to be inline
  `input_data.text_document` (not an S3 `storage_upload`), so the long-flagged S3-blob-port blocker **does not apply** to
  the document facet — a real reduction in outstanding risk, not a deferral.
- **PII discipline held end-to-end.** Customer media (video/audio/docs) never entered agent context; the video exhibit is
  by-reference (no byte-port); the Bunny keys were searched + are handled values-blind; the re-capture ran read-only prod
  through the tool (counts-only), and the scrub leak gates stayed green.
- **Zero platform edits held** — the whole milestone is rext tooling (`stack-seeding`) + corpus docs.

## What Didn't
- **The media port was assumed landable at design time; reality was a hard external blocker.** The roadmap pre-flagged
  "media → PORT IT" with a likely `DEF-M10-01` (S3) consumption; the actual dependency turned out to be **Bunny.net
  recording signing keys** (a different, unprovisioned credential), and the recorded pool is almost entirely HIRING
  interview VIDEO (a weightier data-controller call than the approved audio-only framing). The resolution (presence-only +
  route the exhibit to the live-billion milestone where the keys may be reachable) was the right call, but it took a
  pause/resume + an explicit fresh user decision to reach — worth pre-checking media-credential availability earlier in a
  future media milestone.

## Carried Forward
- **`DEF-M240-01` → M244** (Fate-3, user pre-approved 2026-07-22): the content-stories real-video exhibit — provision the
  Bunny recording signing keys + re-pin a hiring-voice cell to a recorded session + wire the exhibit-by-reference render;
  land live IF the keys are reachable on `billion`, else keep voice presence-only. Zero byte-port. Added to M244's In-list.
- **(Inherited, confirmed) → M244:** the 8/9 standing demo-stack test failures (Fate-2, M238-D5 / M239-D13 reap-17700) +
  `DEF-M239-01` (Fate-3, M239-D12) — already homed; M244 is the named expiry.

## Metrics Delta
- **Tests:** rext whole-repo Go test funcs 1976 (v2.5 baseline) → **1999** (release-cumulative); M240 added **6** deepening
  tests (harden). Full `stack-seeding` module GREEN (16 ok packages, 0 fail). Flake **0**.
- **Deferral audit:** YELLOW (1 tracked repeat = the standing demo-stack debt, fated Fate-2 → M244; DEF-M240-01 Fate-3 →
  M244 user-pre-approved). 0 RED blockers.
- **Platform-repo edits:** 0. **PII/media into context:** 0. (Full machine-readable delta: `metrics.json`.)
