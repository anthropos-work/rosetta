---
milestone: M30
slug: field-bake
version: v1.6 "stage door"
milestone_shape: section
status: archived
created: 2026-06-14
last_updated: 2026-06-14
complexity: medium
delivers: the proven .agentspace/secrets reference dir (built from stack-dev) + the field-bake record; ext tag stage-door-m30
backlog_refs: (none)
---

# M30 — Field-bake: build a compliant secret dir from stack-dev + prove it

## Goal
Prove the whole mechanism on a real stack: assemble a compliant `.agentspace/secrets` dir inferred/pulled from
current stack-dev, run `provision` into a fresh stack, and assert the observable behavior (coverage Critical ==
100%, the stack comes up). This is the user's explicit final step — "build a secret dir compliant with what's
requested, inferring/pulling from current stack-dev to double-check it works."

## Why section
The acceptance gate is objective and known (build dir → provision into a fresh `dev-N`+`demo-N` → `measure`
Critical == 100% → stack UP). Mirrors v1.5's M25 field-bake (an observable-behavior gate, not an iterate-to-a-
fuzzy-bar). Build with `/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`**: any field-fix code the bake surfaces (Fate-1), tagged `stage-door-m30`.
- **`rosetta`**: the field-bake record + the honesty-residual note (in `secrets-spec.md` or the milestone record);
  the proven `.agentspace/secrets` reference dir is gitignored (lives in `.agentspace`, never committed).

## Scope
- **In:**
  - **Assemble a compliant `.agentspace/secrets` dir** from current stack-dev (names-correct, alias-mapped, the
    knowns `waived`). ~85–90% of secrets are present on disk today; the rest correctly `waived`.
  - Run `provision` into a fresh `dev-N` and a fresh `demo-N`; assert `measure` Critical == 100% and the stack
    reaches UP (the **observable-behavior gate**).
  - Fix any real bugs the bake surfaces, Fate-1 (the v1.5 field-bake caught + fixed 4).
  - Document the **honesty residual** — which ~10–15% is `waived` (AWS-mount, optional Bunny/GCloud, profile-gated)
    and why that's correct, not a hole.
- **Out:** new features (the bake is a proving + fix milestone, not a feature milestone).

## Depends on
M29 (the skill + doc the bake exercises).

## Parallel with
None (final milestone of v1.6).

## Estimated complexity
medium.

## Open questions
- Which N to bake on — a throwaway `dev-N≥1` + a `demo-N`, **never N=0**. Settle at bake time per box capacity.

## KB dependencies
- `corpus/ops/secrets-spec.md` — the contract being proven.
- `corpus/ops/rosetta_demo.md` + `corpus/ops/setup_guide.md` — bring-up.
- `corpus/ops/verification.md` — the gate.

## Delivers →
The proven (gitignored) `.agentspace/secrets` reference dir + the field-bake record; ext tag `stage-door-m30`.
