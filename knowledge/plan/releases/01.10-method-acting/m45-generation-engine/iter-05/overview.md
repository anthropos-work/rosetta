---
iter: 05
milestone: M45
iteration_type: tik
status: closed-fixed
date: 2026-06-26
---

# iter-05 — tik — `cmd/gen-batch` (the generation CLI, component 4)

## Active strategy reference
TOK-01 (inside-out fixtures-first build). Fourth tik (5th iter); component (4).

## Re-survey
TOK-01 named `cmd/gen-batch` as iter-05's target; not absorbed. Target current.

## Cluster / target identified
`cmd/gen-batch` is the orchestrator that wires the three landed layers (blueprint expansion → cache →
services/ai) into the actual generation flow with the mandatory `--max-cost` guard. It's the last piece
before the seeder (component 5) + the real gate-proving run, and it's where the valid-JSON-rate +
collision + cost gate metrics are MEASURED.

## Hypothesis
A CLI that expands `EffectiveBatches` → checks the cache (hit?$0) → generates the misses via `services/ai`
under a `--max-concurrent` semaphore with re-roll-on-malformed + post-gen hero-collision re-roll, aborting
BEFORE the `--max-cost` ceiling, and emitting the cost report + valid-JSON rate — can be driven end-to-end
by a FIXTURE completion (no key) to prove every path: cache reseed at $0, re-roll, drop-after-budget,
collision re-roll, ceiling abort, dry-run, lock fence.

## Expected lift
Infrastructure + the MEASUREMENT harness (the gate's valid-JSON-rate / cost / collision metrics are now
computable on a batch). No empirical valid-JSON-rate move yet (that's the REAL run in a later tik), but the
machinery that measures it lands + is fixture-proven.

## Phase plan
Protocol §4b: build (component 4) → unit-test every orchestration path against fixtures (mandatory-cost
validation, dry-run-no-LLM, generate+$0-reseed, re-roll, drop, collision re-roll, ceiling abort, lock
fence) → full-suite regression → close.

## Escalation conditions
- The ceiling guard can't be made to abort BEFORE a breach under concurrency → user-blocker. (Did not
  happen — the pre-launch `WouldExceed` check + the serialized tracker abort the run; test proven.)

## Acceptable close-no-lift outcomes
N/A — build tik; closes `closed-fixed` on the CLI landing green.

## Note (cap)
This is the 5th tik of the session (iter-02..05 under the bootstrap-tok-led first call). After it closes,
the 5-tik cap fires (Phase 5 §5) → the call exits `cap-reached`. The next call resumes with the
GeneratedBatchSeeder (component 5), then hero-collision-in-seeder + the cost-emission + the REAL
gate-proving batch.
