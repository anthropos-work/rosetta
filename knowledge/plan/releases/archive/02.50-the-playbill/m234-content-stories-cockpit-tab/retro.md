# M234 — Retro

## Summary
M234 delivered the **render half** of Content stories: `cockpit.py`'s 2nd "Content stories" tab reading the M233
`content-manifest.json` (per-product sections, per-session rows, two login-and-land CTAs, AI-labs presence-only,
academy direct-origin link), the `content-player-<idx>` roster-seat registration (single-sourced with the UsersSeeder
via the new `storyPopulationNames`), and the `up-injected.sh` `--content-export`/`--content-manifest` bring-up wiring.
The renderer handles **every** product disposition, unit-proven against the manifest. Section milestone, closed-complete;
merged `--no-ff` into `release/02.50-the-playbill`. **0 platform-repo edits.**

## Incidents This Cycle
- **None.** No regressions, no new test failures, no flakes. The build + harden pass (fd457bf) found 0 bugs; the close
  review added 0 code changes. Flake gate 5/5 both stacks. The 6 pre-existing `test_cockpit.py` failures are the
  chronic demo-stack test-debt carry (unchanged; M234 added 0 new).
- **Process note (not an incident):** the whole-rext Go test-func headline had drifted to a non-reproducible "1954"
  (M233 close). This close re-anchored it to the reproducible `git grep '^func Test'` method (1931→1939) and recorded
  the method in `metrics.json` + the headline — the v1.11 incident-#6 counting-hazard discipline, applied proactively.

## What Went Well
- **The M233 single-source design paid off at render time.** Because the manifest is projected from the same fixture
  the seeder seeds from, the render half had a stable, honesty-gated contract to build against — no manifest/render drift.
- **Harden did the heavy lifting.** Pass 1 deepened coverage to 96% (cockpit.py) / 100% Go func on touched files with
  22 behaviour tests, so the close review surfaced only 2 record-level items and 0 code changes — a genuinely near-clean close.
- **The unit-vs-live boundary held honestly.** M234 claimed exactly "unit-proven, not browser-proven" and homed the live
  proof cleanly in M235 — no over-claiming, no stranded CTAs (the fail-closed resolver + renderer-handles-all-dispositions
  design means new fixtures light up without a renderer change).

## What Didn't
- **Nothing blocking.** The only friction was the test-func-count reconciliation (resolved by re-anchoring to a
  reproducible method). Two record-level close fixes (Adversarial-review subsection was owed by Phase 2c; §7 back-ref
  tags were owed by Phase 5) — both cheap, both landed.

## Carried Forward
- **Non-simulation product fixtures (ai-labs / academy / skill-path) + prove-every-CTA-lands live → M235** (Fate-2;
  verified homed in M235's `In:` + exit_gate; the renderer already handles these dispositions, unit-proven).
- **Specific-member academy landing + exact chapter route → M235** (depends on M230's catalog fill).
- **14 pre-existing demo-stack test failures → v2.5 release-close** re-anchor (REPEAT/CHRONIC, inherited; M234 added 0;
  user-dispositioned, homed — see `audit-deferrals/deferral-audit-2026-07-19-m234-close.md`).

## Metrics Delta (from metrics.json)
- **Go test funcs (whole rext repo, `git grep '^func Test'`):** 1931 (M233 @ c30fee3) → **1939** (M234 @ fd457bf), **+8**.
- **Python demo-stack:** 249 pass / 6 pre-existing fail / **0 new** (test_cockpit.py + test_tooling.py, 255 collected).
- **Coverage:** cockpit.py 92% → **96%**; Go 100% function coverage on M234-touched files (except users.go Seed 97%, pre-existing).
- **Flake:** 0 (5/5 both stacks). **Platform-repo edits:** 0. **Supply chain:** 0 net-new deps. **Alignment:** 100%/100% (untouched).
