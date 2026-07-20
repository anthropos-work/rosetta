---
milestone: M236
slug: prove-on-billion
version: v2.5 "the playbill"
milestone_shape: iterative
status: planned
created: 2026-07-19
last_updated: 2026-07-20
depends_on: M235
exit_gate: "Both tabs work live on billion — Content-stories sessions render real content for player + manager vantages, the academy grid renders real cards (Thread A) — reproducibly on a cold reset-to-seed, p95 click->ACCESS < 5 s, 0 platform edits, demo reachable only over the tailnet."
iteration_protocol_ref: corpus/ops/verification.md + corpus/ops/demo/tailscale-serve.md
delivers: none
---

# M236 — prove on billion

**Status:** `planned`  ·  **Shape:** `iterative`  ·  **Complexity:** medium  ·  **Depends on:** M235

## Goal
Re-prove the whole feature live on the billion Tailscale VM (the house pattern that closed M215/M221/M226/M228) — both cockpit tabs usable end-to-end from a 2nd machine on a cold reset-to-seed, VPN-scoped.

## Exit gate
Both tabs work live on billion — Content-stories sessions render real content for player + manager vantages, the academy grid renders real cards (Thread A) — reproducibly on a cold reset-to-seed, p95 click->ACCESS < 5 s, 0 platform edits, demo reachable only over the tailnet.

**Iteration protocol:** `corpus/ops/verification.md + corpus/ops/demo/tailscale-serve.md`

## Scope
### In
- Bring up the demo on billion; drive both the Org-stories and Content-stories tabs remotely
- Prove content-stories sessions render live for player + manager; both academy tabs render (Thread A); reproduce on a cold reset-to-seed; capture p95 click->ACCESS vs the <5 s gate
- **AUTHOR the new content-stories seat-login coverage/Playthrough plumbing (Fate-3 from M235's
  USER-BLOCKER-M235-02, user-authorized 2026-07-20).** The M235 pre-flight PROVED the existing exact-path /
  hero-crawl coverage harness structurally cannot reach the content-stories result pages (dynamic-URL +
  cockpit-seat-reached, not hero-linked). M236 must AUTHOR it — a spec that logs in as each
  `content-player-<idx>` seat (the Playthroughs-style cockpit seat-switch, `playthroughs/e2e/lib/hero-login.ts`)
  + resolves each session's EXACT `/sim/<slug>/result/<sessionId>` URL from the seeded `content-manifest.json`
  + reuses the shared `AISimulationResultContainer` page-object — and CALIBRATE it against a LIVE seeded
  render (authoring it blind ships an INCORRECT descriptor into a load-bearing harness).
- **Prove EVERY (session × action) lands on a non-empty result — the SIMULATION sessions AND the M235 NON-sim
  sections (all built offline + unit-proven at M235, tags `playbill-m235-nonsim-{skillpath,ailabs,academy}`).**
  Work the per-section M235 live-calibration checklists (M235 iter-05/06/07 decisions.md): **skill-path**
  (getOrCreateSkillPathSession version-match — the seed writes version "2"; the `active` vs `in_progress`
  status vocabulary; the `local_skill_path_sessions` mirror uniqueness); **ai-labs** (the exact
  `public.lab_sessions` NOT-NULL column set — an app-side table not in the offline snapshot — plus: presence-
  only, no CTA landing to prove, just a clean seed + the presence row rendering); **academy** (wire
  `app/cmd/academy-seed --user-id <the academy content-player owner> --fixture in-progress|completed` into the
  cold bring-up for real `academy_chapter_progress`; calibrate the CTA from the anonymous `/library/<slug>`
  preview to the authed progress-bearing chapter route).
- **The M230 carry-forward live items (Fate-3, `m230-academy-demo-fill/carry-forward.md`) — all need the live
  cold bring-up M236 runs anyway:** the **ANT_ACADEMY** coverage descriptor (rendered-card count, no Draft
  chip) consuming rext tag `playbill-m230-academy-fs-published`; the **next-web clone re-anchor** (the 2 drifted
  demopatch manifests) as a cold-`/demo-up` prerequisite; the **`getPublicCatalogView` 2nd-manifest**
  anonymous-routes follow-on.

### Out
- New feature work (the SEEDERS + the manifest sections are built + unit-proven by M235 — M236 CALIBRATES +
  PROVES them live, it does not re-build them)

## Open questions
- Fold Thread A's live-render prove into this milestone (one combined release -> yes)
- **RESOLVED (2026-07-20, user-authorized Fate-3 from M235):** M236 now OWNS the content-stories seat-login
  coverage/Playthrough plumbing (M235's USER-BLOCKER-M235-02 proved it can't be authored blind — it needs a
  live seeded render) + the live (session×action)-lands proof for BOTH the simulation and the M235 non-sim
  sections + the M230 carry-forward live items. See the expanded `In:` list above and
  `m235-prove-it-lands/decisions.md` (RESOLUTION of USER-BLOCKER-M235-02).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
