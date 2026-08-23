# M271 — Spec notes

_Technical notes accumulate here during the build. Nothing below is measured yet — the headers are the
scope's shape, not findings._

## Pre-flight audits

_Not run._

## B4 — the shared-AWS safety question

_Design-time citations (carried from `overview.md`, re-measure before acting):_
`livekit.go:158-180` (room-composite Egress to S3) · `LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test`
· `LIVEKIT_RECORDING_AWS_*` · `secretdna/secret-dna.json:382,396` (`LIVEKIT_API_KEY` / `LIVEKIT_API_SECRET`
at `criticality: standard`).

### Is `anthropos-livekit-test` shared with production?

_Unmeasured._

### The `safety.md` §2.3 framing

_Not written._

### The data-controller decision

_Not taken._

## B1 — self-hosted LiveKit feasibility

_Design-time citation: `stack-demo/platform/docker-compose.yml` declares four services (`backend`,
`studio-desk`, `next-web-app`, `gotenberg`); no LiveKit in the compose or in rext._

_Nothing probed._

## B2 — the demopatch surface for the hardcoded endpoint

_Design-time citations: `docker-compose.yml:38` `LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud` ·
`packages/ui/src/AISimulation/AISimulationCall/AISimulationCallContainer.tsx:58`
`serverUrl='wss://anthropos-pbvktu3v.livekit.cloud'` · `Onboarding/TestRuns/TestRunNetwork.tsx:377`._

### Anchor enumeration

_Not enumerated._

### G1–G7 admissibility

_Not assessed._

## B3 — the out-of-repo agent worker

_Design-time citation: `app/internal/jobsimulation/calls/livekit.go:106-168` (`CreateAgentDispatch`),
`Agents: [{AgentName: "anthropos-agent" | "anthropos-agent-chain" | "anthropos-agent-us"}]`. Five
`livekit-agent*` repos — in no `repos.yml`, no demo clone set, started by no rext script._

### Agent-name → repo mapping

_Unmeasured._

### Obtainability for a demo

_Not probed._

## B5 — headless microphone

_Design-time citation: `TestRunMicrophone.tsx` / `TestRunNetwork.tsx` gate on real device permissions;
candidate flags `--use-fake-device-for-media-stream` + `--use-fake-ui-for-media-stream` in
`e2e/playwright.config.ts`._

_Nothing measured._

## Cost model

_No numbers. A GO without a cost number is not a GO (`overview.md` § Scope)._

## Inherited from the dissolved M206

### recording (Chime)

_No disposition._

### the skill-paths verify-skill terminal

_No disposition._

## Verdict

_Not reached._
