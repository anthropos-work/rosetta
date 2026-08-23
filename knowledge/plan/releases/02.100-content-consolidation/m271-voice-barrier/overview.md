---
milestone: M271
title: "Voice go/no-go barrier"
milestone_shape: iterative
status: planned
release: "02.100-content-consolidation"
exit_gate: "a written GO or NO-GO verdict for demo voice, each of the five blockers below either resolved or declared unresolvable, with the data-controller safety decision recorded — modelled on M231 (corpus/ops/demo/content-stories-routes.md), the go/no-go barrier that gated the whole Thread-B chain."
iteration_protocol_ref: "corpus/ops/demo/content-stories-routes.md"
re_scope_trigger: "If the safety decision on blocker 4 is NO, the verdict is NO-GO and the milestone closes on that verdict — a NO-GO reached honestly IS the deliverable, not a failure."
depends_on: "M269"
parallel_with: "none"
complexity: large
last_updated: "2026-08-23"
---

# M271: Voice go/no-go barrier

**HARD go/no-go barrier** — the deliverable is a **verdict**, not a feature. A NO-GO reached on evidence
closes this milestone successfully.

**Goal:** Decide, **on evidence rather than hope**, whether a demo can run a REAL voice simulation — and if
so at what cost and under what safety posture.

Serves the **voice clause** of annotation request **B2** (`.agentspace/annotations.md`, § *content
consumption* item 2: *"one that also requires voice (the voice pipeline has to work)"*). **The user chose
option (a) — barrier only this release, decide after evidence (2026-08-23).**

> ⚠️ **This milestone MUST NOT assume voice is achievable.** Every sentence below is written so that
> "NO-GO" is an admissible, non-failing outcome. Any planning artifact that presumes a working voice
> pipeline contradicts the milestone's own shape.

## Exit gate

A written **GO or NO-GO verdict for demo voice**, each of the five blockers below either **resolved** or
**declared unresolvable**, with the **data-controller safety decision recorded** — modelled on **M231**
(`corpus/ops/demo/content-stories-routes.md`), the go/no-go barrier that gated the whole Thread-B chain.

## Iteration protocol

[`content-stories-routes.md`](../../../../../corpus/ops/demo/content-stories-routes.md) — the M231
feasibility-spike lineage: enumerate the unknowns, probe each one against real code or a real render,
classify the outcome per unknown, and publish the classification **including the ones that came back
negative**. M231's own value was that it ruled AI-labs **OUT** and cut the landable denominator 31 → 29;
a barrier that only ever says GO is not a barrier.

## Why iterative (not section)

**The deliverable is a VERDICT.** Writing an `In:` list of *deliverables* now would presume the answer.
And **four of the five blockers are unknowns whose resolution order depends on what the previous probe
finds** — e.g. whether blocker 3 (the agent worker) is even worth probing depends on the blocker 4 safety
ruling, and what blocker 2 has to patch depends on whether blocker 1 has a self-hosted target at all.

The `In:` list below therefore enumerates **probes**, not products.

## The five blockers — all measured

These citations are the **milestone contract**. They were measured at design time and are carried verbatim;
do not paraphrase them away, and re-measure before acting on any of them rather than trusting the date.

### B1 — There is no LiveKit container anywhere

`stack-demo/platform/docker-compose.yml` declares **four** services (`backend`, `studio-desk`,
`next-web-app`, `gotenberg`). **No LiveKit in the compose or in rext.**

### B2 — The endpoint is hardcoded in the FRONTEND, not an env var

- `docker-compose.yml:38` — `LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud`
- **AND** `packages/ui/src/AISimulation/AISimulationCall/AISimulationCallContainer.tsx:58` —
  `serverUrl='wss://anthropos-pbvktu3v.livekit.cloud'`
- also `Onboarding/TestRuns/TestRunNetwork.tsx:377`

**Repointing a demo at a self-hosted LiveKit REQUIRES A DEMOPATCH.** This corpus takes **zero
platform-repo edits**; the only sanctioned route for a change baked into platform source with no
env/config/compose seam is the sha-pinned demopatch mechanism
([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)) — patch the demo's own
ephemeral clone before the image build, revert after, canonical repos never touched. A demopatch here is
**two anchors in two files**, not one, and the second is a `packages/ui` component shared by more than the
call container.

### B3 — The AI counterpart is an out-of-repo agent worker

`app/internal/jobsimulation/calls/livekit.go:106-168` (`CreateAgentDispatch`) creates the room with
`Agents: [{AgentName: "anthropos-agent" | "anthropos-agent-chain" | "anthropos-agent-us"}]`.

Those workers live in the **five `livekit-agent*` repos**, which are in **NO `repos.yml`**, **NO demo clone
set**, and are **started by no rext script**. **Without a registered worker in that cloud project, the hero
joins a room and NOTHING ANSWERS.**

### B4 — ⚠️ A voice session writes to SHARED AWS. This is a safety gate, not a footnote

The same call configures **room-composite Egress to S3** (`livekit.go:158-180`) with
`LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test` and `LIVEKIT_RECORDING_AWS_*`.

`stack-secrets` provisions `LIVEKIT_API_KEY` / `LIVEKIT_API_SECRET` at **`criticality: standard`** into a
demo's `platform/.env` (`secretdna/secret-dna.json:382,396`) — **so a demo CAN mint tokens against the
shared project.**

**THIS MAKES "MAKE VOICE WORK" A `safety.md` §2.3 QUESTION — the same class as the production-S3 exposure —
NOT merely an engineering one.** The **data-controller decision is a GATE**, not a footnote. See
[`safety.md`](../../../../../corpus/ops/safety.md) §2.3 (never-write shared Directus / prod-S3) and §3.8
(the bounded read-side exception + the VPN/tailnet scope that is the control).

### B5 — Playwright has no microphone

`TestRunMicrophone.tsx` / `TestRunNetwork.tsx` gate the voice onboarding on **real device permissions**; a
headless run needs `--use-fake-device-for-media-stream` and `--use-fake-ui-for-media-stream` in
`e2e/playwright.config.ts`.

## Prior art — do not re-derive

- `playthroughs/manifest/ai-simulations.yaml:4` **already grades voice out**: *"NON-VOICE only this
  release; voice (LiveKit) + recording (Chime) -> the M206 mirror tier"*.
- **M244 dispositioned the recorded-video exhibit player-presence-only under `DEF-M240-01`**, blocked on
  **BOTH** the Bunny key *values* **AND** a *provisioning path*
  ([`media-substrate-spec.md:118-137`](../../../../../corpus/ops/demo/media-substrate-spec.md);
  **`BUNNY_RECORDING_*` = 0 occurrences in all of rext**).

Both are inputs to this milestone's verdict, not questions for it to reopen.

## Scope

> The `In:` list enumerates **probes and a verdict**, not deliverables. See *Why iterative* above — an
> `In:` list of products would presume the answer this milestone exists to find. Resolution **order** is
> deliberately not fixed: it depends on what the previous probe returns.

**In:**
  - **B4 first — the data-controller safety decision on shared-AWS writes.** Establish, in writing, whether
    a demo may mint tokens against the shared LiveKit project and trigger Egress to
    `anthropos-livekit-test`. Framed as a `safety.md` §2.3 question. **This is a gate: a NO here ends the
    milestone at NO-GO** (see the re-scope trigger).
  - **B1 — self-hosted LiveKit feasibility.** Whether a LiveKit server can exist in a demo stack at all
    (compose service, rext ownership, cost, port offset), or whether the only reachable target is the
    shared cloud project.
  - **B2 — the demopatch surface.** Enumerate the exact anchors a repoint would need
    (`AISimulationCallContainer.tsx:58`, `TestRunNetwork.tsx:377`, `docker-compose.yml:38`) and judge
    whether they satisfy the demopatch contract (G1–G7) or escalate. **No platform edit, ever.**
  - **B3 — the agent-worker question.** Whether a registered worker can be obtained for a demo — clone
    set, start path, model cost, and what "nothing answers" costs a presenter if it cannot.
  - **B5 — headless-microphone feasibility.** Whether the fake-device flags make the voice onboarding
    gates passable in the existing e2e harness, measured rather than assumed.
  - **Cost.** If the verdict trends GO, what a voice demo costs per run (LiveKit minutes, egress, the
    agent's model calls) — the brief asks for *"at what cost"*, so a GO without a number is not a GO.
  - **The written verdict + the net-new doc.** `corpus/ops/demo/voice-feasibility.md`, GO or NO-GO,
    with every blocker marked resolved or declared unresolvable.
  - **Inherited from the dissolved M206** (see below): a written disposition for **recording (Chime)** and
    the **skill-paths verify-skill terminal**. Disposition, not necessarily delivery.

**Out:**
  - **Actually building a mirror voice agent.** That is the **NEXT release**, and **only if this returns
    GO**.
  - Any platform-repo edit. A need that can only be met by a platform edit **escalates**; it does not edit.
  - Re-deriving the prior art above.

## Depends on

**M269** (the non-voice modality harness) — the thing voice would extend. A voice Playthrough, if one is
ever built, is a fifth modality on M269's harness, not a parallel structure; probing voice against a
harness that does not exist yet would measure the harness, not the voice.

## Parallel with

none

## Open questions

Honest uncertainty, recorded here rather than resolved by invention:

- **What does the data controller actually decide on B4, and on what basis?** The precedent set
  (`safety.md` §3.8 / §3.8.1 — the 2026-07-19 text sign-off and the 2026-07-21 VIDEO sign-off) is about
  **reading** production content into a bounded scope. B4 is a **WRITE to a shared bucket from an
  unauthenticated, authz-weakened, all-interfaces-published stack**. Whether the existing precedent
  extends, or whether this is a genuinely new class, is **not known** and must not be assumed either way.
- **Is `anthropos-livekit-test` actually shared with production, or is it already an isolated test
  bucket?** The name suggests test; the name is not evidence. This is measurable and should be measured
  early, because a genuinely isolated bucket materially changes the B4 ruling.
- **Are all five `livekit-agent*` repos live, or are some decommissioned?** Three agent names appear at
  `livekit.go:106-168` and five repos exist (`corpus/architecture/org-repos.md`); the mapping between
  them is unmeasured.
- **Does a self-hosted LiveKit even satisfy the frontend?** `AISimulationCallContainer.tsx:58` is a
  `serverUrl`, but token minting, the agent dispatch API and Egress may each have cloud-only
  assumptions. B1 GO does not imply B2/B3 GO.
- **What is the honest fallback if the verdict is NO-GO?** Presence-only (the `DEF-M240-01` shape), a
  scripted non-live exhibit, or an explicit "voice is not in the demo" disclosure. Naming the fallback is
  part of the verdict; **choosing** it may not be this milestone's call.
- **Does B5 have a second gate behind the first?** `TestRunMicrophone.tsx` / `TestRunNetwork.tsx` gate on
  device permissions; whether anything downstream also gates on **real audio energy** is unmeasured, and
  a fake device emits a synthetic tone, not speech.

## Inherited from the dissolved M206

M206 ("AI-sim mirror tier") was a **reservation, never a milestone** — no `overview.md`, re-reserved
across five consecutive releases before v2.8 M256 re-fated it under `D-v28-4`
(`knowledge/plan/roadmap-vision.md:312-345`). This milestone takes the **voice** half as its subject and
**inherits the rest**:

- **recording (Chime)** — the second half of M206's signature voice/recording journey;
- **the skill-paths verify-skill end-to-end TERMINAL** — one of the non-gate legs M206 absorbed at the
  M203 close.

**M206 no longer exists as a reservation.** Recorded in [`decisions.md`](decisions.md) so the routing
exists **at its destination** — the M258 lesson: *a routing written in a closing milestone's decisions is
not a routing until the target's own doc says so.*

## KB dependencies

- [`content-stories-routes.md`](../../../../../corpus/ops/demo/content-stories-routes.md) — **the barrier
  shape this milestone copies** (M231) and its `iteration_protocol_ref`
- [`safety.md`](../../../../../corpus/ops/safety.md) **§2.3 and §3.8** — the exposure decision (B4)
- [`ai_architecture.md`](../../../../../corpus/architecture/ai_architecture.md) — the voice engine
- [`media-substrate-spec.md`](../../../../../corpus/ops/demo/media-substrate-spec.md) — the recording half
  (and `DEF-M240-01` at `:118-137`)
- [`playthroughs.md`](../../../../../corpus/ops/demo/playthroughs.md) — where a voice Playthrough would
  land if one ever exists
- [`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md) — the only sanctioned route for
  B2's frontend repoint
- [`org-repos.md`](../../../../../corpus/architecture/org-repos.md) — where the five `livekit-agent*`
  repos are registered (added for B3; not in the design brief's list)

**Delivers → `corpus/ops/demo/voice-feasibility.md`** (**NET-NEW** — **demo voice has NO doc anchor
anywhere today**, the one Phase-0b blind area in this release). The doc ships with the verdict **whichever
way it goes**; a NO-GO doc that records *why* is the deliverable, not a placeholder.

**Delivers → an amendment to `corpus/ops/safety.md`** — **only if the verdict is GO.** A GO changes what a
demo may reach and must be written into the safety contract, not left implicit.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** Any platform source change goes through the sha-pinned **demopatch**
  mechanism ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)). A need that cannot
  be met that way **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed
  to origin** (the M236 pre-flight rung zero — *tagging is not publishing*), then consumed per-stack at a
  pinned tag.
- Secrets handled **values-blind** — no verb reads, echoes, logs or commits a value. `LIVEKIT_API_KEY` /
  `LIVEKIT_API_SECRET` and any Bunny key are handled under that rule.
- **Customer media never enters an agent's context** (`media-substrate-spec.md` PII discipline). You
  orchestrate the tooling; you do not view the media.
