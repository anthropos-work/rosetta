---
title: "KB Fidelity Audit — M235 prove-it-lands"
date: 2026-07-19
scope: milestone:M235
invoked-by: build-mstone-iters (Phase 0b, iter-01 bootstrap tok)
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths (rext authoring clone) | Status |
|---|---|---|---|
| Per-product result routes + prove-by-render (M231) | `corpus/ops/demo/content-stories-routes.md` | next-web result pages (platform, read-only) + `stack-seeding/contentsession/` | PAIRED |
| Session-clone / sourcing seeder (M232) | `corpus/ops/demo/session-clone-spec.md` | `stack-seeding/seeders/content_stories{,_write,_modality}.go`, `cmd/content-capture/`, `scrub/` | PAIRED |
| Content-stories manifest + honesty gate (M233) | `corpus/ops/demo/content-stories-spec.md` §1–§6 | `stack-seeding/seeders/content_manifest.go`, `cmd/stackseed` `--content-export` | PAIRED |
| Cockpit "Content stories" tab + content-player seats (M234) | `corpus/ops/demo/content-stories-spec.md` §7 | `demo-stack/cockpit.py`, `stack-seeding/seeders/roster.go`, `storyPopulationNames` (users.go), `presets/content-manifest.json` | PAIRED |
| Playthroughs harness (function proof) | `corpus/ops/demo/playthroughs.md` | `playthroughs/manifest,e2e,seed,report/`, `seed/pt-world.seed.yaml` (Org A–D) | PAIRED |
| Coverage sweep (presence proof) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/` | PAIRED |
| Academy demo-fill (M230 carry-forward) | `corpus/services/ant-academy.md`, `corpus/ops/demo/frontend-tier.md` | rext tag `playbill-m230-academy-fs-published` (academy-fs-published demopatch) + `app/cmd/academy-seed` (platform, read-only) | PAIRED |
| Interview flag demopatches (M232) | `session-clone-spec.md` §5 | `demo-stack/patches/next-web-interview-flag-{container,result}` | PAIRED |
| Safety read-side exception | `corpus/ops/safety.md` §3.8 | (contract doc; enforced by seeder isolation) | PAIRED |

No BLIND-AREA. The net-new M235 surface (Playthrough + coverage descriptors for content-stories; the
non-simulation product player-path builders) is DOC-DECLARED as M235's own deliverable in
`content-stories-spec.md` §6 — expected DOC-ahead-of-code, the doc is the contract M235 honors.

## Fidelity Findings

### KB-1 — `content-sessions.yaml` header comment describes the SUPERSEDED synthesize-first posture (STALE)
- **Source:** `stack-seeding/contentsession/fixture/content-sessions.yaml` header comment (rext clone).
- **Expected (per the authoritative corpus doc `session-clone-spec.md` M232):** the seeder COPIES the real
  prod free-text (LLM feedback / transcript / submission / interview report), SCRUBS best-effort of detectable
  PII, and ships it in `fixture/content/<key>.json`. NOT "provably PII-free"; residual risk accepted,
  VPN/tailnet-scoped (`safety.md` §3.8).
- **Actual:** the header still reads *"NO free-text is captured — every LLM feedback string, every transcript
  line, every candidate submission, every actor name is SYNTHESIZED by the seeder, never copied. So this file
  is provably PII-free."* This is the pre-rework M232 synthesize-first framing; the M232 close reworked it to
  copy-real+scrub (per state.md: *"A synthesize-first build was REWORKED to copy-real per user decision"*), and
  the `fixture/content/*.json` copied blobs + the `scrub/` package are the evidence the rework landed.
- **Verdict:** STALE (code-comment only; the corpus knowledge doc is ALIGNED).
- **Fix owner:** update the code comment. **NOT applied inline** — it lives in the rext fixture file M235's
  first fixture-extension tik edits anyway, so the fix rides that tik's commit (a rext code change belongs in an
  iter commit, not the audit). Routed as the fixture tik's opening housekeeping edit.

## Completeness Gaps
None critical. Every load-bearing M231-M234 symbol the docs cite resolves in the rext clone
(`BuildContentProducts`, `contentProductRegistry`, `ValidateContentManifest`, `WriteContentManifest`,
`--content-export`, `ContentStorySeeder`, `package scrub`, `cmd/content-capture`, the cockpit content tab, the
`content-player-<idx>` roster seats, the two interview flag demopatches, pt-world Org A–D).

## Applied Fixes
None applied inline (see KB-1 routing rationale).

## Open Items (require user decision)
None. KB-1 is a mechanical code-comment fix routed to the fixture tik.

## Gate Result
YELLOW: proceed. No blind areas, no stale load-bearing KNOWLEDGE-doc claim. The single finding (KB-1) is a
stale code-comment in the fixture file the milestone extends first; recorded here + routed to the fixture tik.
Calling skill (build-mstone-iters) may proceed to the bootstrap tok's strategy authoring.
