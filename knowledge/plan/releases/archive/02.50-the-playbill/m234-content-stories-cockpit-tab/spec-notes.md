# M234 — Spec notes

## Pre-flight audits — §1/§2/§3 (whole milestone)
KB-fidelity audit (2026-07-19): **GREEN**. Report: `kb-fidelity-audit.md`. Every M234 topic anchored + fresh
(M231/M232/M233 landed the specs same-day); render half + seat registration correctly scoped to M234.

## Topic → doc → code triples
- content-manifest projection → `corpus/ops/demo/content-stories-spec.md` §2/§3 → `stack-seeding/seeders/content_manifest.go` + `presets/content-manifest.json`
- result-route dispositions (AI-labs presence-only, academy IN, has_manager_view matrix) → `content-stories-routes.md` §2/§5/§6 → `content_manifest.go` `contentProductRegistry`
- cockpit render → `cockpit-spec.md` + (M234 render half) `content-stories-spec.md` §6 → `demo-stack/cockpit.py`
- content-player seat + roster → `content-stories-spec.md` §3 + `corpus/services/clerkenstein.md` → `stack-seeding/seeders/roster.go` (+ `content_stories.go` `eligiblePlayerOwnerSlots`, `users.go`, `userprofile.go` name derivation)
- bring-up wiring → `cockpit-spec.md` "Bring it up" → `demo-stack/up-injected.sh` (§2b roster export + §4b cockpit launch)

## Scope boundary (from roadmap M234 vs M235)
- **M234 (this, `section`):** the cockpit render half + `content-player-<idx>` seat registration + bring-up
  `--content-export`/`--content-manifest` wiring + docs. Renderer handles ALL product dispositions driven by
  manifest fields; unit-proven against synthetic manifests (no live browser here).
- **M235 (`iterative`):** populate the tab with INTERESTING real-shaped sessions (fixture additions — academy
  depends on M230) + prove every CTA lands on a non-empty result page. The academy/skill-path player-path
  builders + their fixture sessions are M235; today's fixture is simulation-only, so a real M234 demo shows the
  Simulation section (the AI-labs/academy render paths are exercised by unit tests + light up when M235 adds
  the fixtures).

## Test baselines (pre-M234)
- `demo-stack/tests/test_cockpit.py`: 83 tests, **6 pre-existing failures** (removed academy-CTA ×4 + overlay
  30s-window ×2) — the standing carry routed to the v2.5 release close. M234 must add **0** new failures.
- `stack-seeding` go build + `seeders` tests: clean.
