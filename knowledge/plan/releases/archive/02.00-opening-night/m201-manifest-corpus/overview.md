---
milestone: M201
slug: manifest-corpus
version: v2.0 "opening night"
milestone_shape: iterative
exit_gate: "The manifest corpus is comprehensively outlined, validated, and written as prose-intent YAML (one file per product) — covering the platform's products × their must-work user journeys, each use case carrying goal + actor + flow + intermediate/final expectations, structurally valid (the spec §5.3 validator passes, ids unique + both-way) — and the USER signs off the corpus as the complete-enough v2.0 coverage contract."
iteration_protocol_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the capability spec — esp. §2 model, §4 use-case shape, §5.3 manifest format)
status: archived
created: 2026-06-28
last_updated: 2026-06-28
complexity: large
delivers: the prose-intent manifest YAML corpus (one file per product) — the goal-aligned Product → Story → Use Case build+regression contract every coverage milestone (M203/M204) implements against. Lands in the rext `playthroughs` section once M202 exists; until then drafted under this milestone dir / `knowledge/plan/spec-drafts/playthroughs/manifest-draft/`
depends_on: none (the manifest is prose; it can be authored before/parallel to the M202 foundation, which then builds the validator + dedicated seed to match it)
spec_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the consolidated capability spec, v0.3 — esp. §2 model, §4 use-case shape, §5.3 manifest format)
---

# M201 — Manifest corpus

## Goal
**Top-down, user-directed curation of the full goal-aligned Product → Story → Use Case manifest corpus** — the
build + regression contract every coverage milestone (M203 / M204) implements against. The flow per top-down
pass: **outline** (products → stories → use cases) → **validate** (against the real platform surface + the spec's
manifest model, §2/§4/§5.3) → **write the prose-intent manifest YAML** (spec §5.3, **one file per product**).

This is **explicitly NOT bounded by the current minimal / partially-aligned demo stories seed** — it captures
**what the goal says must be proven**, not only what the demo happens to set up today. Where a use case needs
preconditions the demo lacks, that **feeds the M202 dedicated-seed expansion** (note this dependency — **do not
resolve it here**; M202 owns the seed + validator that satisfy the manifest the user signs off).

## Shape — iterative, USER-GUIDED
**`iterative`**, and **driven by the user.** The user directs each top-down pass; this milestone is worked
**conversationally** + via `/developer-kit:work-mstone-iters` — **not autonomously**. The manifest is prose (no
platform code), so it is authorable **before / in parallel with** the M202 foundation; M202 then builds the
validator + the dedicated seed to **match** the corpus the user signs off.

## Exit gate (objective, plus a user sign-off)
> **The manifest corpus is comprehensively outlined, validated, and written as prose-intent YAML (one file per
> product)** — covering the platform's products × their must-work user journeys, each use case carrying **goal +
> actor + flow + intermediate/final expectations**, structurally valid (the spec §5.3 validator passes, **ids
> unique + both-way**) — **and the USER signs off the corpus as the complete-enough v2.0 coverage contract.**

The structural half is machine-checkable (the §5.3 validator: unique ids, both-way id integrity, precondition
resolvability). The **completeness** half is a **deliberate user judgement** — "is this the right, complete-enough
set of journeys to prove for v2.0?" — because there is **no introspectable schema for "what user-facing
capabilities exist"** (spec §5.9, the accepted-risk asymmetry). The user is the owner of *enough*.

## Why iterative (not section)
The corpus is **discovered top-down with the user**, pass by pass — which products, which stories, which use
cases, at what depth — and "complete enough" is a judgement the user converges on, not an enumerable checklist
fixed up front. Each pass outlines a slice → validates it against the real platform surface + the manifest model →
writes it as prose-intent YAML → the user steers the next slice. Build with `/developer-kit:work-mstone-iters`,
worked conversationally.

## Depends on / Parallel with
- **Depends on:** **none** — the manifest is **prose**. It can be authored **before or in parallel with** the
  **M202 foundation** (which then builds the §5.3 validator + the dedicated seed to match it).
- **Parallel with:** **M202** (the foundation). No hard ordering: M201 produces the prose contract; M202 produces
  the machinery that validates + seeds against it. They reconcile when M202's validator runs over the M201 corpus.
- **Feeds → M203 / M204** — the per-vantage coverage milestones implement Playthroughs against the use cases this
  corpus declares.

## The dedicated-seed dependency (note, don't resolve)
A use case may declare a precondition the **current demo seed lacks** (a tier, a private-org-X path, a recruiter
pipeline, an entitlement state). This milestone **records that need in the use case's `seed`/`preconditions`** and
**feeds it to the M202 dedicated-seed expansion** — M202 owns building the decoupled seed (+ the §5.3
precondition-coverage validator) so every declared precondition resolves. M201 **does not** build or extend any
seed; it states what the goal requires.

## Re-scope trigger
If the user signals the corpus should be narrowed/widened mid-flow, that is normal top-down steering (this is a
user-guided milestone) — re-outline the affected slice. A precondition that the demo can't support is **not** a
blocker here: it is captured as a use-case precondition and **handed to M202** (above).

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) — the capability
  spec: **§2** (the Product → Story → Use Case → Playthrough model + vocabulary), **§4** (what a use case
  declares — `goal` / `actor` / `actor.entitlement` / `seed` / `flow` / `outcome` / intermediate + final
  expectations, §4.1–§4.4), **§5.3** (the prose-intent YAML manifest format + the light validator: unique ids,
  both-way id integrity, precondition-coverage).
- [`../m202-foundation/overview.md`](../m202-foundation/overview.md) — the **M202 foundation** (the §5.3 validator
  + the dedicated decoupled seed that satisfy this corpus's preconditions).
- [`corpus/ops/demo/stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md) — the existing demo stories
  & heroes (Maya / Dan) the demo-covered products draw on (the *starting point*, NOT the bound).
- [`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the seeding machinery the M202
  dedicated-seed expansion (which this milestone feeds) builds on.

## Delivers →
- The **prose-intent manifest YAML corpus** — **one file per product** (spec §5.3), each declaring its stories →
  use cases (goal + actor + flow + intermediate/final expectations + `seed`/`preconditions`), structurally valid
  (unique ids, both-way id integrity). **Lands in the rext `playthroughs` section once M202 exists**; until then
  **drafted under this milestone dir** or `knowledge/plan/spec-drafts/playthroughs/manifest-draft/`.
- The user's **sign-off** that the corpus is the complete-enough v2.0 coverage contract.
