# M47 — Re-sync & recapture — retro

## Summary
The foundation milestone of the v1.10b "fit-up" backfill. Scoped as a heavy "re-sync 5-week-stale clones + recapture
+ re-validate." The headline outcome: **the premise was wrong** — the clones were already current, so the heavy
re-sync was a trivial `make pull` and the flagged ⚠ "biggest unknown" evaporated. The real code deliverable was
small: `pg.NormalizeDSN` (sslmode `no-verify→require`) so the wired `postgres` MCP DSN works as a `primary-read`
capture `--dsn` — proven end-to-end by a live dry-run + the full recapture of all 3 snapshot surfaces (digests
unchanged). The AI-readiness feature was confirmed present in current code, resolving the M201 false-negative and
re-pointing the backfill's "stale" framing at its true target: the **corpus** (M48).

## Incidents this cycle
- **P3 (recovered) — accidental background-task kill.** The first taxonomy recapture (~1.4 GB) was stopped
  mid-stream by an accidental UI action. **No harm:** the atomic `.tmp`→rename write left the prior cache intact
  (no corruption, no partial dir), and a clean re-launch completed it. Reinforces that the snapshot store's atomic
  write is doing its job.
- No regressions; no test failures; flake gate 5/5.

## What went well
- **Measured before trusting the doc.** Running `git fetch` + a behind-count (rather than believing the M201
  "115 commits behind" note) surfaced the wrong premise immediately — saved a pointless multi-hour rebuild.
- **The sslmode fix was provable cheaply.** A `--dry-run` capture connected to the live wired DSN and proved
  `NormalizeDSN` end-to-end without streaming a full pull.
- **Scope shrank honestly under user direction.** "No new entry point" collapsed S2 from a planned `up-injected.sh`
  auto-capture mechanism to "reuse the existing capture step + the wired DSN" (D3) — less code, same capability.

## What didn't
- **The milestone was over-scoped on a stale premise.** M47's whole heavy framing (and part of v1.10b's "re-ground"
  thesis, and the v2.0 pause justification) rested on "clones 5 weeks behind," which didn't hold. Lesson: a
  close-discovered staleness claim should be *re-measured* at the next milestone before it drives a release's shape.
- **Early over-design.** The first instinct was to build a new MCP-DSN extraction entry point (`~/.claude.json`
  reader / env var) — the user correctly redirected to the existing process. Reach for the existing mechanism first.

## Carried forward
- **Consumption-clone re-pin** of `stack-demo/rosetta-extensions` to `fit-up-m47` — DEFERRED (push-gated; tracked
  with the release's other pending origin pushes). The fix only affects capture, not the consumed bring-up tooling.
- **Corpus staleness** (the genuine one) → **M48**: document the member-AI-readiness flow + reconcile
  `corpus/architecture` + `corpus/services`.
- **Quoted-keyword `sslmode` value** not normalized — known limitation, accepted (real DSN is URL-form). decisions.md.

## Metrics delta
- rext Go test funcs: stack-snapshot 363 → **364** (+1 `TestNormalizeDSN`, 9 table cases incl. 2 adversarial);
  rext total 1551 → **1552**. Flake **0**. New deps **0**. Snapshot: all 3 surfaces recaptured, digests unchanged.
  (Full: `metrics.json`.)
