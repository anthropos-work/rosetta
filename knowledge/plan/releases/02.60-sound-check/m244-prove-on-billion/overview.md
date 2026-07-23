---
milestone: M244
slug: prove-on-billion
version: v2.6 "sound check"
milestone_shape: iterative
status: archived
created: 2026-07-20
last_updated: 2026-07-23
depends_on: M238, M239, M240, M241, M242, M243
exit_gate: "On a cold reset-to-seed on billion: (a) ORG-CLEAN reports 0 surviving source-org tokens (or each dispositioned) — RUN FIRST, read-only, before the bring-up; (b) content-stories run-content-stories.sh green at the shipped harness with the CQ-1 grader fix + CQ-2 runner wiring + externally-sourced EXPECTED_PAIRS (discharges CLOSE-D3); (c) the 40 live-browser specs (24 stack-verify + 16 Playthroughs, incl. M243's assign-WRITE) execute green (T-3); (d) the anonymous academy /library+/free twin renders real cards (S-1); (e) DEF-M226-01 — the serve-reap self-resolution claim is actively tested or DROPPED; (f) the 3 v2.3 drift-carries burned-in live (BURNIN-M221 / F-M220-4 / PROBE-M218-c3); (g) the interview plan-section-id alignment assertion added + green (S-8/S-9); (h) every v2.6 fix (academy course-start, talk-to-data live answer, library, content fidelity incl. media, language toggle, cockpit UX) proven live; p95 click->ACCESS < 5 s hero vantages. 0 platform edits."
iteration_protocol_ref: corpus/ops/verification.md + corpus/ops/demo/tailscale-serve.md + coverage-protocol.md + playthroughs.md
delivers: none
---

# M244 — prove on billion  [realizes reserved M237]

**Status:** `archived` (completed 2026-07-23)  ·  **Shape:** `iterative` (the closer)  ·  **Complexity:** medium  ·  **Depends on:** M238, M239, M240, M241, M242, M243 (all fixes)

## Goal
Re-prove the whole feature — v2.5's headline `29/29` AND every v2.6 fix — live on the `billion` Tailscale VM, cold reset-to-seed (the house pattern that closed M215/M221/M226/M228/M236). Realizes the reserved `M237` (re-prove-on-billion): the release it re-proves (v2.5) shipped its headline metric unverified-live.

## Exit gate
On a cold reset-to-seed on `billion`:
- **(a)** `ORG-CLEAN` reports **0** surviving source-org tokens (or each dispositioned) — **RUN FIRST**, read-only, before the bring-up.
- **(b)** content-stories `run-content-stories.sh` green at the shipped harness with the CQ-1 grader fix + CQ-2 runner wiring + externally-sourced `EXPECTED_PAIRS` (discharges CLOSE-D3).
- **(c)** the **40 live-browser specs** execute green (T-3) — **24 stack-verify + 16 Playthroughs** (was 39 = 24 + 15; M243 added the 16th Playthrough spec, `assignment-assign.spec.ts` / `pt-assignment-assign`, whose cold re-drive on `billion` M244 owns — proven GREEN + DB-verified at build on demo-1).
- **(d)** the anonymous academy `/library`+`/free` twin renders real cards (S-1).
- **(e)** `DEF-M226-01` — the serve-reap self-resolution claim is **actively tested or DROPPED**.
- **(f)** the 3 v2.3 drift-carries burned-in live (`BURNIN-M221` / `F-M220-4` / `PROBE-M218-c3`).
- **(g)** the interview plan-section-id **alignment assertion** added + green (S-8/S-9).
- **(h)** every v2.6 fix (academy course-start, talk-to-data live answer, library, content fidelity incl. media, language toggle, cockpit UX) proven live; **p95 click→ACCESS < 5 s** hero vantages.
- **0 platform edits.**

**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` + `coverage-protocol.md` + `playthroughs.md`.
**Re-scope trigger:** 5 consecutive toks without a viable strategy → user-strategic-replan.

## Scope
### In
- Bring up a fresh demo on `billion`, cold reset-to-seed; drive every v2.6 fix + v2.5's headline live from a tailnet peer.
- Run the multi-part gate (a–h) above; `ORG-CLEAN` runs FIRST (read-only, before the bring-up).
- Same billion-safety rules (one driver, no detached on-host scripts, assert from a tailnet peer, never kill a mid-build).
- **Inherited from M239 close (2026-07-21, Fate-3):** `DEF-M239-01` — make the demo build **fail loudly on ENOSPC** (the disk-full class that surfaced as a cryptic `redis exited(1)`; M239 fixed the pre-flight to measure the Docker-VM disk, this is the build-time loud-abort follow-on). **9th standing failure** — `test_a_RACED…` hardcodes cockpit port 17700 and collides with a live demo-1; a clean reset-to-seed resolves it, but land the test-isolation fix (fix recipe in M239 `decisions.md` D13). Both non-blocking; discharge during the live re-prove.
- **Inherited from M240 (2026-07-22, Fate-3, user pre-approved):** the content-stories VIDEO exhibit — re-pin a hiring-voice cell to a recorded session + provision the Bunny.net recording signing keys (`BUNNY_RECORDING_CDN_TOKEN_KEY` + `BUNNY_RECORDING_PULL_ZONE_HOST`, absent from the local dev-stack) + wire the exhibit-by-reference render; the posture + gender-coherence contract are pre-blessed (safety.md §3.8.1); land the exhibit live IF the keys are reachable on billion, else keep voice presence-only. Zero byte-port (sign-a-Bunny-URL-at-render).

### Out
- New feature work (all built by M238–M243 — M244 CALIBRATES + PROVES live, it does not re-build).
- Content-seat latency (out of scope, per the v2.5 M236 precedent — the p95 gate is hero-vantages-only).

## Open questions
- None blocking (the multi-part exit gate is the spec). Each carry has a named condition in the gate; the two conditional carries (`DEF-M226-01` "test or DROP"; the interview alignment assertion) resolve inside the milestone.

## Delivers
None (proof milestone; extends the coverage/playthrough manifests + burns in the carries).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
