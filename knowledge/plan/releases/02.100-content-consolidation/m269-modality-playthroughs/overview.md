---
milestone: M269
title: "Modality playthroughs"
milestone_shape: section
status: planned
release: "02.100-content-consolidation"
depends_on: "M267"
parallel_with: "none"
complexity: large
last_updated: "2026-08-23"
---

# M269: Modality playthroughs

**Goal:** A hero actually PLAYS a simulation to completion in each non-voice modality — chat, code,
document — proven by Playthroughs that create **real sessions** rather than stopping at the launch screen.

Serves annotation request **B2** (chat, code, document — **NOT** voice, which is M271).

> ⚠️ **This milestone lands part of M206, a SIX-TIMES-DEFERRED reservation.**
> [`roadmap-vision.md:311-322`](../../roadmap-vision.md) states that a sixth re-reservation *"is not an
> option this file permits"* — and it happened anyway. The v2.10 design run **DISSOLVES M206**: its
> `ai-simulations.code.UC1`, `ai-simulations.interview.UC1` and `profile.self-evaluation.UC1` land
> **HERE**; its voice + recording + skill-paths-verify terminal route to **M271**. **No M206 reservation
> survives this release.** Recorded as `D-M269-1` in [`decisions.md`](decisions.md).

## Scope

**In:**
  - **Move the assertion boundary past the launch screen.** The existing coverage stops **EXACTLY ONE
    CLICK SHORT**, and says so in its own header —
    `playthroughs/e2e/tests/aisim-chat-launch.spec.ts`:
    > `@pt-mutation: READ-ONLY` … *"MEASURED read-only at M256 iter-06: the click reaches
    > `/sim/<slug>/start` and renders the launch confirmation, and 0 `jobsimulation.sessions` rows were
    > created during the probe. The session is written later, past the welcome dialog — on the far side
    > of the §5.8 live-AI boundary."*

    `/sim/<slug>/start` is `AISimulationStartWithoutSession`, which renders **BEFORE**
    `handleCreateSession` is called. A Playthrough that lands here proves the launch screen, not the play.
  - **Record the boundary move as a DECISION, not a test tweak.**
    `playthroughs/manifest/ai-simulations.yaml:7-11` states the non-voice, launch-only boundary **AS
    POLICY** ("ASSERTION BOUNDARY (spec §5.8 + P2) … It asserts at the LAUNCH boundary … the only thing
    provable under P6 with a live LLM in the loop"). **Moving it is the bulk of this milestone's cost**
    and must be argued in `decisions.md`, not slipped in as a locator change.
  - **Three modalities, each pinned to a REAL replayed catalog sim.** Modality is a property of the
    simulation *definition* replayed from the real public catalog, not something we author:
    voice → `simTasks[].taskType ∈ {Call, InterviewCall}`; code → `sequences[].sequenceType ==
    SequenceTypeCoding`; document → a `collaborativeAssets[].filenameDownload` passing `isDocumentFile`
    (`packages/graphql/src/hooks/aiSimulation/useGetSimulationFlagsAndFeatures.ts:22-68`).
    [`content-stories-routes.md`](../../../../corpus/ops/demo/content-stories-routes.md) records
    **77 voice / 65 code / 30 document** public sims, so the source pool is not the constraint.
  - **Retire the single-slug pin.** Today the harness pins **ONE** slug — `SAMPLE_CHAT_SIM_SLUG` at
    `e2e/lib/simulation-page.ts:27`. Three modalities need at least three, selected **by the modality
    predicate above**, never by keyword.
  - **[inherited from the dissolved M206]** land `ai-simulations.code.UC1`,
    `ai-simulations.interview.UC1` and `profile.self-evaluation.UC1`.
  - **[carried deferral — `FIX-M256-studio-false-green`, + `NEGCTL-M256-studio-pair`,
    `DOC-M256-llm-lane-premise`]** the studio Playthrough matches **EMPTY SCAFFOLDING at +2.1 s**, before
    the LLM draft populates. [`roadmap-vision.md:493-500`](../../roadmap-vision.md) calls it *"a false
    green inside the suite v2.8's headline claim rests on"*; deferred M256 → M258, never worked.
    **It matters directly to this release:** on **2026-08-23** both studio Playthroughs reported **PASS**
    on `demo-1` and were cited as evidence the migrated studio works. **Fix the oracle, then re-run that
    claim.**
  - **[carried deferral — `BIND_HOST` / `D-M255-7`, deferred THREE times: M255 → M256 → M258]**
    `up-injected.sh:146` binds `0.0.0.0` whenever `STACK_PUBLIC_HOST` is set, so **the batch gate SKIPS
    on the default `/demo-up` path**. Measured live on **2026-08-23**: the `demo1` bring-up recorded the
    Playthrough gate as **`skipped`**, never green, and the suite had to be driven manually from a tailnet
    peer. **New Playthroughs added here would record `skipped` on the very stack they exist for.**
    Landing this is what makes M269 **self-verifying**.

**Out:**
  - **VOICE and recording** — M271.
  - **Any new simulation CONTENT.** The catalog is **replayed**, not authored. If no public sim in the
    replayed catalog satisfies a modality predicate, that is a finding, not a licence to write one.
  - **Any platform-repo edit.** This corpus takes **ZERO** platform-repo edits. If a modality cannot be
    played without a platform source change, it goes through the sha-pinned **demopatch** mechanism
    ([`demopatch-spec.md`](../../../../corpus/ops/demo/demopatch-spec.md)) or it escalates as
    `unimplementable-without-platform-edit`.

## Depends on

M267 — there is no point proving a start that is still entitlement-gated.

## Parallel with

none — it changes the **shared** Playthrough runner and manifest.

## Open questions

  - **Is "to completion" reachable at all under P6 with a live LLM in the loop, for each of the three
    modalities?** `ai-simulations.yaml:7-11` argues it is not, and that argument has stood since M201.
    This milestone's premise is that the boundary can move **at least as far as a real
    `jobsimulation.sessions` row**; whether it can move all the way to a *result* per modality is
    **NOT measured** and may differ per modality (code and document have deterministic substrate;
    chat does not).
  - **What is the session-creation ORACLE?** A DB row in `jobsimulation.sessions` is the honest signal
    the M256 probe already used to prove the negative — but the suite's page-object layer is
    semantic-by-default and DB-free today. Whether the oracle is a DB assert, a UI landmark past the
    welcome dialog, or both, is undecided.
  - **Does landing `BIND_HOST` actually un-skip the gate?** `roadmap-vision.md` warns:
    *"Re-check the premise before landing: `BIND_HOST` gates only the two host-native servers; compose
    publishes the containers on `0.0.0.0` regardless, so a one-line change may not be sufficient."*
    [`verification.md`](../../../../corpus/ops/verification.md) gives a **second, independent** reason
    for the skip — a connection from the demo host to its own tailscale IP hits the kernel socket and
    **bypasses `tailscale serve`**, which is what terminates TLS. **Two causes, one symptom.** Fixing the
    bind may leave the skip in place.
  - **Do the three modality slugs survive a re-capture?** The pins are catalog slugs; the taxonomy/content
    replay is versioned. Whether a modality pin should be a literal slug or a runtime *query* over the
    replayed catalog is open.
  - **Does the boundary move change the `@pt-mutation` classification of the whole product?** Every
    non-voice AI-sim Playthrough is currently `READ-ONLY`. Sessions are writes; the reset-to-seed
    lifecycle and the serial-default runner were sized against a read-only product.
  - **What does M206's `profile.self-evaluation.UC1` actually require?** It arrives here by dissolution,
    not by design, and this scaffold has **not** sized it. It is the one In: item with no measured
    citation behind it.

## KB dependencies

- [`playthroughs.md`](../../../../corpus/ops/demo/playthroughs.md) — the manifest model, the 4-state
  reporting map, and §5.8 (the live-AI boundary this milestone moves)
- [`content-stories-routes.md`](../../../../corpus/ops/demo/content-stories-routes.md) — the per-modality
  catalog counts (77 voice / 65 code / 30 document)
- [`verification.md`](../../../../corpus/ops/verification.md) — the batch gate contract that `BIND_HOST`
  breaks
- [`demopatch-spec.md`](../../../../corpus/ops/demo/demopatch-spec.md) — the only sanctioned route for a
  platform source change

## Delivers

- `corpus/ops/demo/playthroughs.md` — the assertion boundary moves; **the live-Playthrough count and the
  4-state reporting map change** with it, and §5.8's launch-boundary argument is superseded for the three
  non-voice modalities (and only those).
- `corpus/ops/verification.md` — the batch gate **stops skipping on the default `/demo-up` path**.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** A need that can only be met by a platform edit goes through the sha-pinned
  **demopatch** mechanism ([`demopatch-spec.md`](../../../../corpus/ops/demo/demopatch-spec.md)) or it
  **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged,
  **pushed**, then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
