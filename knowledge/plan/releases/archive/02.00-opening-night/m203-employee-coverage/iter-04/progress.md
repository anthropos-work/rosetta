**Type:** tik (under TOK-01, step 2) — protocol: `corpus/ops/demo/playthroughs.md` §"The iteration protocol".

# iter-04 progress

- **Recon (be-the-human):** the Skill Paths library is `/library/skill-paths`; the catalog is the real
  replayed public content (path slugs like foundation-of-artificial-intelligence). A path detail renders 5
  chapters + a "Start" CTA + "0% complete". Start → the chapter player (`/skill-path/<slug>/chapter?...`) with a
  real step ("What is Artificial Intelligence") + a "Mark complete & continue" advance control. The path's
  assessment (Chapter 05) is a VOICE sim → M206 tier, OUT.
- **Assertion-boundary decision (spec §5.8 + M201 P7):** the drivable, deterministic, NON-voice boundary is
  browse→open→Start→progress. The terminal "complete → skill verifies" EARNS verification via an ASSESSMENT
  (mostly voice this release) and the verify OUTCOME is proven on the profile side (profile.verified.UC1). So
  this UC asserts the learning journey is real + drivable.
- **Declared** skill-paths.legacy.UC1 (new `skill-paths.yaml` product; added `public-catalog` seed-world
  capability — the catalog is the snapshot set-dress, not the roster seed). ptvalidate VALID (2 products, 5 UCs).
- **Added** `SkillPathPage` page-object (library / path-detail / player accessors).
- **Built** skillpath-legacy.spec.ts (@pt:pt-skillpath-legacy). First run FAILED: strict-mode violation — the
  path detail has TWO "Start" buttons (primary CTA + first-chapter inline). **Fixed** by scoping startButton to
  `.first()` (the document-order-first primary CTA) — the locator-discipline disambiguation (§5.2). Re-run PASS.
- **Reconciled** (ptreport --gate no-regressions): **5/5 passing (100.0%)** — GREEN. Go tests + fmt + vet clean.
- **Noted:** `stackseed` (the reset-to-seed dependency) is not on the AUTHORING copy's PATH — the 5-run
  reset-to-seed determinism gate runs from the stack-demo CONSUMPTION clone (where the pinned tools install).
  Sequenced as a milestone-close activity after all journeys are green (route forward).

## Close — 2026-07-02

**Outcome:** +1 employee use case green (skill-paths.legacy.UC1); 2 of the 3 gate journeys covered (Profile
complete + Skill Paths learning). Employee coverage 5/5 declared passing; no-regressions GREEN on demo-1.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the 3rd gate journey — AI Simulations NON-voice — still to land, plus the 5-run
reset-to-seed determinism proof).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (skill-path assertion boundary = browse/open/start/progress; verify-skill composes P7), D2 (Start `.first()` disambiguation), D3 (reset-determinism gate deferred to milestone-close, consumption-clone) — see decisions.md
**Side-deliverables:** none
**Routes carried forward:**
  - iter-05 → **AI Simulations (NON-voice)** per TOK-01 step 3 — the 3rd + final gate journey. The demo catalog
    is heavily VOICE-typed; must find/assert a NON-voice (chat/code/interview-text) sim at the launch/completion
    boundary (§5.8). Handler PT-M203-iter05-aisim.
  - milestone-close → the 5-run reset-to-seed determinism gate (from the stack-demo consumption clone w/
    stackseed installed). Handler PT-M203-reset-gate.
  - later → profile.self-evaluation.UC1 (non-gate M201 extra, rate-modal click-intercept). Handler PT-M203-selfeval.
**Lessons:** the antd skill-path detail renders MULTIPLE "Start" buttons (primary CTA + per-chapter) — a bare
role match is ambiguous; `.first()` (document-order primary CTA) is the stable disambiguation. Skill-path
catalog content is the SNAPSHOT set-dress (shared, `public-catalog`), NOT the roster seed — name the
precondition accordingly. Mutating Playthroughs (Start creates progress) need reset-to-seed for the determinism
gate; the reset dependency (stackseed) lives in the consumption clone, so run that gate at milestone-close.
