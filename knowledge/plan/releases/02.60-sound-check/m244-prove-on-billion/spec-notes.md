# M244 — Spec notes

Iterative milestone (the closer). This will accumulate iteration-protocol notes — per-iter measurements,
gate-condition (a–h) evidence, and the billion bring-up findings — during the iter loop.

**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` + `coverage-protocol.md` + `playthroughs.md`.

## Gate conditions (a–h) — evidence to accumulate
- (a) `ORG-CLEAN` — 0 surviving source-org tokens (RUN FIRST, read-only)
- (b) content-stories `run-content-stories.sh` green (CQ-1 grader fix + CQ-2 runner wiring + external `EXPECTED_PAIRS`)
- (c) the 39 live-browser specs execute green
- (d) anonymous academy `/library`+`/free` twin renders real cards
- (e) `DEF-M226-01` serve-reap — actively tested or DROPPED
- (f) 3 v2.3 drift-carries burned-in live (`BURNIN-M221` / `F-M220-4` / `PROBE-M218-c3`)
- (g) interview plan-section-id alignment assertion added + green
- (h) every v2.6 fix proven live; p95 click→ACCESS < 5 s hero vantages

_(iteration-protocol notes accumulate here during the iter loop)_
