# M244 — Spec notes

Iterative milestone (the closer). This will accumulate iteration-protocol notes — per-iter measurements,
gate-condition (a–h) evidence, and the billion bring-up findings — during the iter loop.

**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` + `coverage-protocol.md` + `playthroughs.md`.

## Gate conditions (a–h) — evidence to accumulate
- (a) `ORG-CLEAN` — 0 surviving source-org tokens (RUN FIRST, read-only)
- (b) content-stories `run-content-stories.sh` green (CQ-1 grader fix + CQ-2 runner wiring + external `EXPECTED_PAIRS`)
- (c) the **40** live-browser specs execute green (**24 stack-verify + 16 Playthroughs**; was "39" pre-M243 — M243 added the 16th Playthrough `assignment-assign.spec.ts`)
- (d) anonymous academy `/library`+`/free` twin renders real cards
- (e) `DEF-M226-01` serve-reap — actively tested or DROPPED
- (f) 3 v2.3 drift-carries burned-in live (`BURNIN-M221` / `F-M220-4` / `PROBE-M218-c3`)
- (g) interview plan-section-id alignment assertion added + green
- (h) every v2.6 fix proven live; p95 click→ACCESS < 5 s hero vantages

## Pre-flight audits — iter-01 (bootstrap tok, 2026-07-22)

**KB-FIDELITY: YELLOW** — report `knowledge/plan/releases/02.60-sound-check/m244-prove-on-billion/kb-fidelity-audit.md`. No blind areas; all gate parts (a–h) have tooling + doc anchors. Two stale-narrative findings, neither blocks (the implementation reads files, not prose):

- **KB-1 — content-stories denominator is now `49`, NOT the historical `29`.** `stack-verify/e2e/content-denominator.json` `expected_pairs=49` (M241 grew simulation sessions 13→23 with EN/IT language counterparts: 23×2 + 2 skill-path-legacy-player + 1 skill-path-new = 49; ai-labs presence-only excluded). `run-content-stories.sh` reads this file as `EXPECTED_PAIRS` and cross-checks the served manifest against it BEFORE the sweep. **⇒ Gate (b) green = 49/49 pairs land, not 29/29.** The "29/29" in `roadmap.md` M244 goal, `state.md`, and `coverage-protocol.md:916-918` is the v2.5/M236 value, now stale; those are prose the live proof does not read — flag for close/harden to reconcile the corpus once the live sweep lands 49/49.
- **KB-2 — spec count is `40` (24 stack-verify + 16 Playthroughs), not `39`.** Verified in code: 16 live-browser playthrough spec files (`playthroughs/e2e/tests/*.spec.ts` minus 2 `.unit`), incl. M243's `assignment-assign.spec.ts`. `overview.md` exit_gate already says 40; `spec-notes.md` (fixed above) and `roadmap.md:462` still said "39" (roadmap flagged for close).

## Pre-flight rung zero — iter-01 (can billion OBTAIN the tooling under test?)
- **billion reachable** (`ssh marco@billion` / `root@billion`, MagicDNS `billion`), but **BARE** as of 2026-07-22: `docker ps` empty, no `~/stack-*`/`~/panorama`, no `rext.tag`, no snapshot cache. A live proof needs a from-scratch cold bring-up (workspace + PAT-HTTPS + secrets + snapshot cache + rext clone).
- **rext tooling published:** the v2.6 tags `sound-check-m237…m243-*` are all on origin. The latest, `sound-check-m243-assign-write-playthrough`, peels to `2ef5962` = the authoring-copy HEAD. **The pin billion must consume is `sound-check-m243-assign-write-playthrough`** (the local `.agentspace/rext.tag` is STALE at `sound-check-m239-enterprise-surfaces` — must be set to m243 on billion).
- **local prereqs present to seed billion:** `.agentspace/secrets` (24K) + `.agentspace/snapshots` (1.4G) exist here to scp.

_(iteration-protocol notes accumulate here during the iter loop)_
