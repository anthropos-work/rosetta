---
title: "M235 Carry-Forward — Routes from prove-it-lands"
date: 2026-07-20
status: archived
close_status: closed-incomplete
gate_target: "on a cold reset-to-seed, every in-scope (session × action) logs in on the correct org and lands on a NON-EMPTY result page for BOTH player and manager vantages, 0 ejects; assessment 2-voice/1-code/1-document PASSED present + each type in passed AND not-passed; each product passes or is declared with a documented fate (AI-labs feasibility answered)"
gate_achieved: "offline-buildable scope EXHAUSTED + unit-proven — 13-session simulation matrix + all 3 non-simulation sections; manifest projects all 4 products / 18 sessions, both honesty gates GREEN; AI-labs feasibility answered (presence-only). The LIVE browser (session × action)-lands proof is NOT run here."
gate_distance: "the LIVE cold-reset-to-seed render proof (both vantages, 0 ejects) — routed to M236 by design; a running-stack measurement, not a build gap"
---

## TL;DR

M235 closed **incomplete by design** under the user's pragmatic-close mandate ("build non-sim seeders, then
close", 2026-07-20). Everything the live proof *depends on* is built and unit-proven with **0 platform-repo
edits**: the provably-clean 13-session simulation fixture matrix, all three non-simulation content-story sections
(skill-path / ai-labs / academy), and the manifest projecting all 4 products / 18 sessions behind two GREEN
honesty gates. The one thing that genuinely needs a **running stack** — the live browser proof that every
`(session × action)` lands NON-EMPTY for both vantages — routes to **M236 (prove-on-billion)** in **3 root-cause
clusters, all Fate-3 and ALREADY applied** to M236's `overview.md` `In:` list (iter-08, commit `54eaefe`,
user-authorized).

## Root-cause clusters

All three clusters share one root cause: **the live proof needs a running stack + a live seeded render**, which
is unavailable offline (and was M230-blocked locally). They are separated because they route distinct work into
M236, not because their causes differ.

### Cluster 1: the LIVE (session × action)-lands proof + the new seat-login coverage/Playthrough plumbing
- **Affected items:** the milestone's exit gate itself (every in-scope session × {player, manager} lands NON-EMPTY,
  0 ejects, cold reset-to-seed); the NEW content-stories coverage/Playthrough harness plumbing.
- **Root cause:** USER-BLOCKER-M235-02 — the planned "coverage descriptor" mechanism does not exist. The existing
  coverage harness matches a page descriptor by EXACT normalized path (`coverage-manifest.ts:989-991`) and reaches
  pages by crawling in-app nav from the hero's landing; the content-stories result pages are **dynamic-URL**
  (`/sim/<slug>/result/<sessionId>`) and reached **only** via the cockpit "login as content-player-N" CTA. So they
  need NEW plumbing: a spec that logs in as each `content-player-<idx>` seat (the Playthroughs cockpit seat-switch),
  resolves each session's exact result URL from the seeded `content-manifest.json`, and reuses the shared
  `AISimulationResultContainer` page-object — **authored AND calibrated against a live seeded render** (selectors,
  the mirror-table manager scoreboard, the per-session score/feedback fence). Authoring it blind ships an
  *incorrect* (not merely uncalibrated) load-bearing harness.
- **Estimated scope:** one iterative milestone's worth (M236 is scoped `medium`).
- **Fate:** Fate 3 — **already applied** (M236 `overview.md` `In:` edited iter-08).
- **Target milestone:** M236 (`../m236-prove-on-billion/overview.md`, `In:` — "AUTHOR the new content-stories
  seat-login coverage/Playthrough plumbing" + "Prove EVERY (session × action) lands on a non-empty result").
- **Provenance:** iter-05 (planning; USER-BLOCKER-M235-02), iter-08 (Fate-3 handoff).

### Cluster 2: the per-section live-calibration checklists
- **Affected items:** three product-arm calibrations that only resolve against a live render — **skill-path**
  (`getOrCreateSkillPathSession` version-match [seeder writes version "2"], the active-vs-in_progress status
  vocabulary, the mirror's `(user_id, skill_path_id)` shape); **ai-labs** (the exact `public.lab_sessions` NOT-NULL
  DDL — an app-side table absent from the offline public snapshot); **academy** (the `academy_chapter_progress`
  write via the `app/cmd/academy-seed` platform binary, the progress-bearing chapter route, the M230 catalog fill).
- **Root cause:** these are shapes/DDL/routes the offline environment cannot observe; each is written up as an
  explicit live-calibration item in the owning iter's `decisions.md`.
- **Estimated scope:** small — verification + narrow seeder adjustments during M236's live loop.
- **Fate:** Fate 3 — **already applied**.
- **Target milestone:** M236 (`../m236-prove-on-billion/overview.md`, `In:` — "Work the per-section M235
  live-calibration checklists").
- **Provenance:** iter-05 (skill-path), iter-06 (ai-labs), iter-07 (academy).

### Cluster 3: the M230 carry-forward live items
- **Affected items:** the **ANT_ACADEMY** coverage descriptor (rendered-card count, no Draft chip) consuming rext
  tag `playbill-m230-academy-fs-published`; the **next-web clone re-anchor** (the 2 drifted demopatch manifests) as
  a cold-`/demo-up` prerequisite; the **`getPublicCatalogView` 2nd-manifest** anonymous-routes follow-on.
- **Root cause:** inherited from M230's own pragmatic-close (`../m230-academy-demo-fill/carry-forward.md`); all three
  need the live cold bring-up M236 runs anyway.
- **Estimated scope:** small — folds into M236's cold bring-up + coverage sweep.
- **Fate:** Fate 3 — **already applied**.
- **Target milestone:** M236 (`../m236-prove-on-billion/overview.md`, `In:` — "The M230 carry-forward live items").
- **Provenance:** M230 carry-forward → M235 `overview.md` In: (inherited) → iter-08 (re-routed to M236).

## Projected post-resolution state

If all three clusters resolve as routed, M236 proves the whole feature live on `billion`: both cockpit tabs usable
end-to-end from a 2nd tailnet machine on a cold reset-to-seed, every `(session × action)` landing NON-EMPTY for
both vantages with 0 ejects, p95 click→ACCESS < 5 s, VPN-scoped, 0 platform edits — i.e. M235's exit gate met
live, which is exactly M236's exit gate. **Uncertainty:** the seat-login plumbing is new and calibrated live, so
M236 is a genuine iterative build (measure→triage→fix), not a rubber-stamp; and Cluster 2's `lab_sessions` DDL /
academy-seed catalog dependency could surface a narrow seeder adjustment. None is expected to reopen M235's build.

## Cross-references
- Iter ledger: ../progress.md (Running ledger)
- Decisions: ../decisions.md (TOK-01; USER-BLOCKER-M235-01/02 + resolutions; Adversarial review)
- Deferral audit: audit-deferrals/deferral-audit-2026-07-20-m235-close.md
- M236 (the routing target): ../m236-prove-on-billion/overview.md
- M230 carry-forward (Cluster 3 origin): ../m230-academy-demo-fill/carry-forward.md
- Iteration protocol used: corpus/ops/demo/playthroughs.md + corpus/ops/demo/coverage-protocol.md
