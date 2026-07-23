# M244 — Retro (prove on billion)

## Summary
The **iterative terminal closer** of v2.6 "sound check" and its final milestone. Re-proved the whole v2.6 feature +
v2.5's headline **live on the `billion` Tailscale VM, cold reset-to-seed**, over 27 iters (24 tiks / 3 toks) across 10
build-iter runs + 2 harden-iters sessions (Passes 1–9). **Gate MET 8/8, 0 platform edits.** Delivers nothing new (proof
milestone) — it CALIBRATES + PROVES; the value is a live-verified release, extended coverage/playthrough manifests, and
the burned-in carries.

## Incidents This Cycle
- **P1 (real defect, fixed) — interview report rendered EMPTY on `billion` (iter-04/05, gate g).** Root cause:
  `directus.simulations_extraction` (the interview PLAN) was **never captured** by the snapshot (0 rows on the demo) →
  null plan → empty report. Fixed by adding the table to the capture surface (rext `e74e563`); plan load proved
  load-bearing (render 0→520). The plan-section-id alignment assertion then caught a second real drift (v1.3-era session
  vs v1.4 plan), fixed by re-pin (iter-06).
- **P1 (real build defect, fixed durably) — `launched_by` version-skew (iter-25).** `up-injected.sh` built the backend
  `:injected` image from the **highest fetched tag** (v1.351.0, has `ai_readiness_cycles.launched_by`) instead of the
  source's **pinned checkout** (v1.341.0, the migrated schema) → the binary SELECTed a column the schema never created →
  cycles endpoint 500'd → every ai-readiness surface rendered the zero-state (looked like a seed gap). Fixed (build-scratch
  resolves the pinned ref + M217-style preflight; rext `c755370`, +3 regression tests). Blended into `verification.md`
  rung zero.
- **P2 (harness, no product defect) — 3 ai-readiness "sub-render failures" mis-characterized as seed gaps (iter-26 D1),
  falsified iter-27.** Against the live DB + v1.341.0 read paths, all three were **harness locator mismatches** vs a
  correctly-rendering UI (byTeam "Team" not "Tag" / interview panel scoped to a 24-char heading span / dueDate short-date
  with no 4-digit year). Fixed in the e2e lib, 0 seed / 0 platform edits.
- **P2 (flake, stabilized) — assign-WRITE tailnet pick-timing (iter-27).** `pickFirstSkillPath` Assign button stayed
  disabled after keyboard-select over the tailnet; passed on re-run. Hardened (option-rendered gate + 3× retry-until-enabled).
- **Process — the binary-per-gate metric fired 3 toks mechanically.** Gate parts only count when fully green (gate b at
  47/47, gate c at 40/40), so the metric read flat over the milestone's most productive tiks. TOK-01/02/03 all HELD the
  strategy (a coarse-metric artifact, not a stall); no revision. No lost work, but the metric under-reports on this shape.

## What Went Well
- **The gate bound the anti-toothlessness thesis.** Every part was proven by *executing* the check live, not by asserting
  it: ORG-CLEAN ran read-only FIRST, the 40 specs actually executed, DEF-M226-01 was TESTED not assumed, the interview
  assertion caught a real drift. Two real defects (interview capture, launched_by skew) surfaced only because the proof was
  live.
- **0 platform edits held across all 27 iters** — every fix routed to `rosetta-extensions` tooling or a sha-pinned demopatch.
- **The two harden-iters sessions kept the teeth.** Mutation-verify (no line-cov tool wired) confirmed every hardenable
  fix bites; a toothless iter-08 scope test + a dueDate year-less locator that false-passed "24 hours"/"5 days" were both
  found + fixed.

## What Didn't
- **The seed-gap mischaracterization (iter-26) cost a re-survey.** iter-26 recorded a plausible-but-wrong seed-gap story;
  iter-27 had to falsify it against the live DB before landing the harness fixes. Lesson: check the locator against the
  actual rendered UI before hypothesizing a data gap.
- **The coarse binary-per-gate metric.** Fired the tok-trigger 3× on productive tiks. Recorded for future iterative-closer
  design — a per-gate-part fractional metric would read progress more honestly.

## Carried Forward
- **standing demo-stack test debt (8 fails)** — pre-M244 inherited (6 `test_cockpit` academy/overlay + `test_public_host`
  port-13001 + `test_purge` docker-integration), 0 real defects → **`/developer-kit:close-release`** (Fate-3, M244 close D1).
- **DEF-M239-01** — ENOSPC loud-build-fail → **`/developer-kit:close-release`** (Fate-3, needs a real ENOSPC to validate).
- **Resolved, not carried:** reap-17700 (LANDED iter-10), DEF-M240-01 (dispositioned iter-07), DEF-M226-01 (TESTED iter-11).

## Metrics Delta (from metrics.json)
- Gate: **8/8 MET** (closed-on-gate). Iters: **27** (24 tiks / 3 toks). Platform-repo edits: **0**.
- content-stories: **47/47** landed live of the 49-pair denominator (2 voice player presence-only).
- Live-browser specs: **40 EXECUTED GREEN** (24 stack-verify + 16/16 Playthroughs, 96 cases) — the v2.5-deferred execution.
- p95 click→ACCESS (cold on `billion`): employee **1.46 s** / manager **1.31 s** (gate < 5 s).
- Standing demo-stack fails: **9 → 8** (final harden fixed the one M244-introduced M215 fence). Flake: **0**.
- rext HEAD `6feae20 → 498b1a5` (final harden +6 mutation-verified tests + 2 inline fixes); consumption tag on origin.
