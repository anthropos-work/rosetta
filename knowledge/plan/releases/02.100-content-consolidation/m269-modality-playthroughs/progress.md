# M269 — Progress

_Not started._

## The boundary move

- [ ] move the assertion boundary past the launch screen — a Playthrough creates a **real**
      `jobsimulation.sessions` row instead of stopping at `/sim/<slug>/start`
      (`aisim-chat-launch.spec.ts` header: *"0 `jobsimulation.sessions` rows were created during the
      probe"*; `/sim/<slug>/start` is `AISimulationStartWithoutSession`, rendered BEFORE
      `handleCreateSession`)
- [ ] record the boundary move as a **decision**, not a test tweak — `ai-simulations.yaml:7-11` states
      the launch-only boundary AS POLICY

## The three non-voice modalities

- [ ] **chat** — pinned to a real replayed catalog sim
- [ ] **code** — `sequences[].sequenceType == SequenceTypeCoding`
- [ ] **document** — a `collaborativeAssets[].filenameDownload` passing `isDocumentFile`
      (`useGetSimulationFlagsAndFeatures.ts:22-68`)
- [ ] retire the single-slug pin `SAMPLE_CHAT_SIM_SLUG` (`e2e/lib/simulation-page.ts:27`) — select by the
      modality predicate, never by keyword

## Inherited from the dissolved M206

- [ ] `ai-simulations.code.UC1`
- [ ] `ai-simulations.interview.UC1`
- [ ] `profile.self-evaluation.UC1`

## Carried deferral — `FIX-M256-studio-false-green`

- [ ] fix the oracle: the studio Playthrough matches EMPTY SCAFFOLDING at +2.1 s, before the LLM draft
      populates (`roadmap-vision.md:493-500`)
- [ ] `NEGCTL-M256-studio-pair`
- [ ] `DOC-M256-llm-lane-premise`
- [ ] **re-run the 2026-08-23 `demo-1` claim** — both studio Playthroughs reported PASS and were cited as
      evidence the migrated studio works

## Carried deferral — `BIND_HOST` / `D-M255-7`

- [ ] `up-injected.sh:146` binds `0.0.0.0` whenever `STACK_PUBLIC_HOST` is set → the batch gate SKIPS on
      the default `/demo-up` path (measured live 2026-08-23: the `demo1` bring-up recorded the Playthrough
      gate as `skipped`, never green; the suite had to be driven manually from a tailnet peer)
- [ ] re-check the premise first — `BIND_HOST` gates only the two host-native servers, and
      `verification.md` records a **second** cause (self-connection to the host's own tailscale IP
      bypasses `tailscale serve`)
- [ ] prove it self-verifying: the new Playthroughs record **green**, not `skipped`, on a bare
      `/demo-up N`

## Docs delivered

- [ ] `corpus/ops/demo/playthroughs.md` — the boundary moves; the live count and the 4-state map change
- [ ] `corpus/ops/verification.md` — the batch gate stops skipping on the default path
