---
iter: 03
milestone: M45
iteration_type: tik
status: closed-fixed
date: 2026-06-26
---

# iter-03 — tik — `blueprint.Batch` + `batch[]` + `EffectiveBatches()` (component 2)

## Active strategy reference
TOK-01 (inside-out fixtures-first build). Second tik; component (2) of the chain.

## Re-survey
TOK-01 named `blueprint.Batch` + `EffectiveBatches()` as the iter-03 target; not absorbed (the blueprint
package had no batch surface). Target current.

## Cluster / target identified
The batch descriptor is the engine's INPUT surface: `cmd/gen-batch` (component 4) needs the per-member
MOTHER PROMPTS, and the cache (component 3) keys on them. Both depend on `EffectiveBatches()` existing,
so it's the right next layer. It mirrors the existing `EffectiveStories()` normalization the seeder fleet
already iterates.

## Hypothesis
A `Batch` type + a `batch[]` field on Story/StackSeed + `EffectiveBatches()` doing PURE Go-template
per-member prompt expansion (NO LLM, NO I/O) can be added without touching any existing seeder, parses via
the existing `KnownFields(true)` path, validates alongside personas, and is fully deterministic (same
descriptor → byte-identical mother prompts → stable hash) — the $0-reseed foundation.

## Expected lift
Infrastructure (no empirical gate-metric move yet — that needs a real LLM call). The lift: the input
surface + the deterministic mother-prompt expansion land, unit-green, so the cache + CLI tiks have prompts
to key + fire.

## Phase plan
Protocol §4b: build (component 2) → unit-test (expansion, determinism, multi-story, validation, YAML
parse) → full-suite regression → close.

## Escalation conditions
- The `batch[]` field breaks the existing `KnownFields(true)` parse of any preset → user-blocker. (Did not
  happen — additive field; full suite + cmd/stackseed green.)

## Acceptable close-no-lift outcomes
N/A — build tik; closes `closed-fixed` on the layer landing green.
