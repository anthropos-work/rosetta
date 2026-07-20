---
iter: iter-01
milestone: M235
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-19
active_strategy: TOK-01
---

# M235 · iter-01 — bootstrap tok (author TOK-01)

**Type:** tok (bootstrap). Iter-01 of the milestone — authors the FIRST strategy from `overview.md` +
`spec-notes.md` + the protocol (`playthroughs.md` + `coverage-protocol.md`). Does NOT terminate the call
(Phase 5 §2) — the loop continues into iter-02 as a tik under TOK-01.

## Inputs
- Milestone `overview.md` (exit_gate, In/Out), release `roadmap.md` § Active v2.5.
- Protocol: `corpus/ops/demo/playthroughs.md` (function proof — automated-actor-IS-the-user, reset-to-seed,
  4-state map) + `corpus/ops/demo/coverage-protocol.md` (presence sweep, fix-surface routing, textMatch-on-values).
- The Thread-B chain M231→M234: `content-stories-routes.md`, `session-clone-spec.md`, `content-stories-spec.md`.
- M230 carry-forward (`m230-academy-demo-fill/carry-forward.md`).
- Environment probe (this iter): docker up (28.5.1), **no stack running**; both `stack-demo` + `stack-dev`
  present; rext authoring clone present @ `fd457bf` (tags `playbill-m234-content-tab{,-hardened}`); prod DB
  access AVAILABLE (`~/.pgpass` present, tailscale up, `cmd/content-capture` exists).

## Baseline reading (the gate distance)

The gate is a LIVE gate: on a **cold reset-to-seed**, every in-scope (session × action) logs in on the correct
org + lands on a NON-EMPTY result page for BOTH player + manager, 0 ejects; assessment PASSED = 2 voice / 1 code
/ 1 document; each sim type in passed AND not-passed; each product passes or gets a documented fate.

**Two-part distance:**

1. **Fixture readiness (buildable + unit-provable HERE).** Current fixture = 9 sessions, `simulation`-only:
   | requirement | have | gap |
   |---|---|---|
   | assessment voice PASSED ×2 | 1 (`asmt-voice-pass`) | **+1 voice-pass** |
   | assessment code PASSED ×1 | 1 (`asmt-code-pass`) | ✓ |
   | assessment document PASSED ×1 | 0 (`asmt-doc-fail` only) | **+1 doc-pass** |
   | assessment in passed AND not-passed | ✓✓ | ✓ |
   | training in passed AND not-passed | ✓✓ (`train-doc-pass`/`train-chat-fail`) | ✓ |
   | hiring in passed AND not-passed | pass only (`hire-voice-pass`) | **+1 hiring not-passed** |
   | interview in passed AND not-passed | pass only (`intv-voice-pass`) | **+1 interview not-passed** |
   | products: skill-path / academy / ai-labs | absent (`simulation` only) | **+3 product sections** |

2. **Live-proof readiness (needs a running stack).** 0 (session × action) live-proven — no stack is up. A
   cold local `/demo-up` is the M230-blocked path (drifted next-web clone + heavy bring-up + prior docker
   trouble). M236 is the on-`billion` live proof milestone (billion already runs demos).

## Open-question resolutions (from the milestone overview, resolved by M231)
- **Q1 "is /sim/.../result/<sessionId> runtime-blank → composed-outcome fallback / demo-patch?"** —
  **FALSIFIED for simulations.** M231 proved the sim result page is a **PERSISTED READ**
  (`jobsimulation/internal/graph/queries.resolvers.go:70`, plain Ent SELECTs, no recompute on render). A clone
  that INSERTs the result fan-out (+ the manager MIRROR) renders a full result — no composed-outcome fallback,
  no demo-patch for the sim result page. (Interview needs the two flag-gate demopatches — already built M232.)
- **Q2 "does not-passed render a meaningful page or blocked/empty?"** — **meaningful.** A not-passed session
  carries a terminal `evaluation_status` + a persisted sub-threshold score, so the result page renders the
  full breakdown below threshold (not blocked/empty). The "each type in passed AND not-passed" gate is
  therefore renderable once the fixture carries not-passed per type (Track A) + the live proof (Track B).

## The strategy (TOK-01) — see milestone-root decisions.md

**Two tracks.** Track A = build + unit-prove everything that makes the live proof READY (fixture matrix
closure, non-sim product sections, Playthrough + coverage descriptors, M230 clone re-anchor) — the bulk of the
tiks, all Fate-1 here. Track B = the FORMAL live cold-reset-to-seed browser proof; attempt locally once ready,
else document the environment constraint honestly + route the definitive live proof to M236, exiting
`user-blocker` only if a genuine environment DECISION is needed that changes what code lands.

## Next-tik direction (iter-02, first tik)
The **fixture matrix closure** — source + capture + scrub the 4 missing simulation sessions (2nd
asmt-voice-pass, asmt-doc-pass, hiring not-passed, interview not-passed) via `cmd/content-capture` over prod
read-only; fix **KB-1** (the stale `content-sessions.yaml` header) in the same edit; re-project
`content-manifest.json` (honesty gate); unit-prove the seeder fan-out + fixture-cleanliness gate
(`TestEmbeddedContent_NoStructuralPII`) + datadna closure. Highest leverage — it is the substrate the manifest
projection, the descriptors, and the live proof all depend on, and it is fully buildable + unit-provable here.

## Phase plan (bootstrap tok)
Author strategy → resolve open questions → frame baseline → write TOK-01 → close (no metric delta; toks don't
move the gate). Continue into iter-02 (tik) within the same call.
