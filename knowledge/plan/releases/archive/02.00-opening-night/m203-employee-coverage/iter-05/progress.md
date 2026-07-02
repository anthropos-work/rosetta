**Type:** tik (under TOK-01, step 3) — protocol: `corpus/ops/demo/playthroughs.md` §"The iteration protocol".

# iter-05 progress

- **Recon:** the sim catalog is heavily VOICE-typed (M206 tier, OUT); found chat-based sims (interview-typed
  reuse the chat engine — M201). Picked `ai-tools-adoption-developer-interview-f78` (a "conversation" sim).
- **Blocker discovered:** launching as pt-employee → deny modal "You cannot start AI Simulations in this
  organization" (while the showcase hero maya-thriving in the default org launched fine).
- **Deep diagnosis (read-only across the platform clones):** FE `canStartAsOrganizationMember` →
  `userMembership.organizationFeatures` → resolver `IsMemberAllowedToUseFeature` → Sentinel
  `OrgMembershipsAllowedToUseFeature` (the g3 grouping-policy matcher, filtered by v0=org). Verified the DB:
  Pat's membership `b84e7dbe` in pt-meridian `ad524614` HAS both the g3 FEATURE grant and the g2 org grant —
  data all correct. **ROOT CAUSE:** the running Sentinel casbin enforcer caches its policy in-memory; the
  seeded g3 grants aren't seen until an explicit Reload RPC.
- **Fix (confirmed live):** triggered Sentinel `Reload` (HTTP 200) → re-tested → pt-employee now launches
  ("Welcome to your AI Simulation", URL /sim/.../start?organizationId=ad524614). Added the post-seed Reload to
  `run-playthroughs.sh` (idempotent, non-fatal, zero platform edits).
- **Declared** ai-simulations.chat.UC1 (new `ai-simulations.yaml`; `sim-feature-enabled` seed-world capability).
  ptvalidate VALID (3 products, 6 UCs, 6 live Playthroughs).
- **Added** SimulationPage page-object (library / detail / launch-boundary + deny-modal accessors).
- **Built + ran** aisim-chat-launch.spec.ts (@pt:pt-aisim-chat-launch) asserting the §5.8 launch boundary
  (launch confirmation renders + deny modal absent + /start route). PASS (1.9m).
- **Reconciled** (ptreport --gate no-regressions): **6/6 passing (100.0%)** — GREEN. Go tests + fmt + vet +
  runner syntax clean.

## Close — 2026-07-02

**Outcome:** +1 employee use case green (ai-simulations.chat.UC1); ALL 3 employee-vantage gate journeys now
covered (Profile complete · Skill Paths learning · AI Simulations chat launch). Employee coverage 6/6 declared
passing; no-regressions GREEN on demo-1. Landed the post-seed Sentinel Reload fix (unblocks sim launch for any
seeded member).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the COVERAGE half of the gate is met — every declared employee journey passes; the
remaining gate criterion is the 5-run reset-to-seed determinism proof: 0 false-fails over 5 cold reset runs,
which runs from the stack-demo consumption clone where stackseed installs — a milestone-close activity).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y (5th tik: iter-02/03/04/05 = 4 tiks... this is the 4th tik; cap is 5 — NOT reached) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the sim-launch root-cause diagnosis + Sentinel-Reload fix), D2 (assert at the §5.8 launch boundary; the completion terminal composes on the profile side), D3 (chat/interview sim choice over voice) — see decisions.md
**Side-deliverables:** the post-seed Sentinel Reload in run-playthroughs.sh — load-bearing for the AI-sim journey (part of planned scope, not an unrelated side-fix).
**Routes carried forward:**
  - **milestone-close** → the 5-run reset-to-seed determinism gate (0 false-fails), from the stack-demo
    consumption clone (stackseed + the Reload now in the runner). Handler PT-M203-reset-gate.
  - later (non-gate M201 extras) → ai-simulations.code.UC1 (Judge0 path) + .interview.UC1 (text) + the
    profile.self-evaluation.UC1 (rate-modal click-intercept). Handler PT-M203-nongate-extras.
**Lessons:** a Sentinel-authz gate (casbin g3 feature grant) is only effective after the enforcer RELOADS —
seeding the grant into the DB is necessary but not sufficient for a RUNNING stack. Any seed that writes casbin
policy for a running enforcer must pair with a post-seed Reload (now in the runner). The AI-sim assertion
boundary (§5.8) is the launch confirmation, not turn-by-turn — the seeded catalog being voice-heavy makes the
chat/interview-typed sim the NON-voice path. **Protocol note applied** (see below).
