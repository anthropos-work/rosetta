---
iter: 04
milestone: M45
iteration_type: tik
status: closed-fixed
date: 2026-06-26
---

# iter-04 — tik — the prompt-hash cache (`batchcache/`, component 3)

## Active strategy reference
TOK-01 (inside-out fixtures-first build). Third tik; component (3).

## Re-survey
TOK-01 named the cache as iter-04's target; not absorbed. Target current.

## Cluster / target identified
The cache is what makes the gate's reproducibility clause ($0 byte-identical reseed) true, and
`cmd/gen-batch` (component 4, next) writes through it. Building it before the CLI lets the CLI be a thin
orchestrator over a tested cache + the tested wrapper.

## Hypothesis
A `batchcache` package — batch dir keyed by sorted-mother-prompts + capture-version, atomic .tmp→rename
member writes, a `.lock` O_EXCL fence, hit/miss + manifest — gives deterministic, $0-reseed-correct
caching with a clean concurrency story, fully unit-testable on a `t.TempDir()` (no key, no network).

## Expected lift
Infrastructure (no empirical gate move). The reproducibility + cost-control machinery lands, unit-green,
so `cmd/gen-batch` can wire generate→cache and the gate's $0-reseed clause becomes provable.

## Phase plan
Protocol §4b: build (component 3) → unit-test (hash determinism/order-independence/invalidation,
put/get/has atomicity, leftover-.tmp, reproducible reseed, capture-version invalidation, manifest, lock
fence, break-lock) → full-suite regression → close.

## Escalation conditions
- The atomic-write / lock semantics can't be made portable → user-blocker. (Did not happen — os.Rename +
  O_CREATE|O_EXCL are POSIX-atomic on the local FS the cache targets.)

## Acceptable close-no-lift outcomes
N/A — build tik; closes `closed-fixed` on the layer landing green.
