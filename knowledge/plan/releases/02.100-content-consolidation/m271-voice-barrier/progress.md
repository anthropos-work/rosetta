# M271 — Progress

**Status: `planned`.** Not started. No iters run.

> **The deliverable is a VERDICT.** Every unchecked box below is a probe, not a feature. A box checked
> with the answer **"unresolvable"** is a **completed** probe — the exit gate reads *"resolved **or
> declared unresolvable**"*. A NO-GO reached honestly closes this milestone.

## Checklist

Ordered as written in `overview.md` § Scope. **The order is not fixed** — resolution order depends on what
the previous probe returns, except that B4 is deliberately first because a NO there ends the milestone.

- [ ] **B4 — the data-controller safety decision on shared-AWS writes** *(GATE)*
  - [ ] `anthropos-livekit-test` measured: genuinely isolated, or shared with production?
  - [ ] The question framed as a `safety.md` §2.3 exposure question, in writing
  - [ ] Decision taken and **recorded** — with who took it and on what basis
  - [ ] If NO → jump to the verdict; the re-scope trigger fires
- [ ] **B1 — self-hosted LiveKit feasibility**
  - [ ] Can a LiveKit server exist in a demo stack at all (compose service, rext ownership, port offset)?
  - [ ] Cost of running it
  - [ ] Verdict: resolved / unresolvable
- [ ] **B2 — the demopatch surface for the hardcoded endpoint**
  - [ ] Anchors enumerated (`AISimulationCallContainer.tsx:58`, `TestRunNetwork.tsx:377`,
        `docker-compose.yml:38`)
  - [ ] Assessed against the demopatch contract G1–G7 — or escalated
  - [ ] Verdict: resolved / unresolvable
- [ ] **B3 — the agent-worker question**
  - [ ] Agent name → repo mapping measured across the five `livekit-agent*` repos
  - [ ] Can a registered worker be obtained for a demo (clone set, start path, model cost)?
  - [ ] What "nothing answers" costs a presenter if it cannot
  - [ ] Verdict: resolved / unresolvable
- [ ] **B5 — headless-microphone feasibility**
  - [ ] Fake-device flags measured against the real onboarding gates — not assumed
  - [ ] Second gate checked: does anything downstream require real audio energy?
  - [ ] Verdict: resolved / unresolvable
- [ ] **Cost** — per-run number for a voice demo (LiveKit minutes, egress, agent model calls), **required
      if the verdict trends GO**
- [ ] **The written verdict + the net-new doc** — `corpus/ops/demo/voice-feasibility.md`, GO or NO-GO,
      every blocker marked resolved or declared unresolvable
- [ ] **Inherited from the dissolved M206** — a written disposition for:
  - [ ] recording (Chime)
  - [ ] the skill-paths verify-skill terminal
- [ ] **If and only if the verdict is GO** — the amendment to `corpus/ops/safety.md`

## Exit-gate ledger

| clause | status |
|---|---|
| A written GO or NO-GO verdict for demo voice | _not reached_ |
| B1 resolved or declared unresolvable | _not started_ |
| B2 resolved or declared unresolvable | _not started_ |
| B3 resolved or declared unresolvable | _not started_ |
| B4 resolved or declared unresolvable | _not started_ |
| B5 resolved or declared unresolvable | _not started_ |
| The data-controller safety decision recorded | _not started_ |

## Iters

_None._
