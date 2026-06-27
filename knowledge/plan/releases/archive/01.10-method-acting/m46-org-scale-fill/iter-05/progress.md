**Type:** tik (#4, under TOK-01)

# M46 · iter-05 — --gen-batches fence + 429/throughput verification

Implemented deliverable #4: the silent-hollow guard + the burst-resilience confirmation.

## What landed (rext `stack-seeding/`)
- `cmd/stackseed/main.go`: new opt-in `--gen-batches` flag + `assertBatchCacheComplete()`. When set, a
  blueprint that declares a `batch[]` MUST have every generated member present in the prompt-hash cache —
  else the seed FAILS LOUD before any write, naming the missing count + the fix (`run gen-batch first`) +
  the cause (OpenAI/Azure unreachable). A no-batch blueprint is a clean pass (the fence only guards batch[]
  stacks). Off by default → the M45 behavior (silently seed whatever's cached) is unchanged.
- `seeders/generated_batch.go`: exported `ReservedHeroNames` (the curated-hero reserved-name set) as the
  ONE source of truth, so gen-batch + the seeder + the fence all rebuild the IDENTICAL mother prompts the
  cache is keyed on. `reservedHeroNamesForSeed` now delegates to it.

## 429 / throughput verification (inspection + locking tests)
- The ai-lib (`github.com/anthropos-work/ai@v1.40.1`) ships `DefaultRetryOptions` (exponential backoff
  1→512s, **10 attempts**, `retryIfOpenAiError` — retries 429s, skips 401/403/404), applied
  **on-by-default** to every chat completion (`openai/completion.go`: `RetryOptions: append(DefaultRetryOptions, …)`
  + `retry.DoWithData`). So a 429 under burst at `--max-concurrent=5` is first absorbed by the lib's
  backoff-retry.
- The wrapper (`services/ai/ai.go`) layers the documented EU→direct fallback ON TOP: `CompleteJSON` falls
  back to direct OpenAI on a persistent Azure 429 (`is429`). Locked with 3 new tests.

## Tests (fixtures-first, no key/cost), full suite green, go.mod UNCHANGED (supply-chain holds)
- `cmd/stackseed/main_test.go` (4): `GenBatchesFence_NoBatchPasses`, `_EmptyCacheFailsLoud` (names "3 of
  3", "gen-batch", "hollow"), `_CompleteCachePasses`, `_PartialCacheFailsLoud` (names "2 of 3").
- `services/ai/ai_test.go` (3): `CompleteJSON_429FallsBackToDirect` (1 primary 429 + 1 fallback call),
  `_Non429DoesNotFallBack`, `_429NoFallbackPropagates`.

## Re-measure
Gate primary metric (the M42 semantic sweep on a generated org) NOT measured this iter — it moves at the
real-run gate-proving tik (#5). With deliverable #4 done, all FOUR code/workflow deliverables (auto-fill,
per-story distribution, preview, the fence + verified 429 backoff) are landed + unit-proven; the only
remaining gate face is the empirical "does a full org rendered from one descriptor LOOK believable under
the M42 semantic gate" — the real-run sweep (TOK-01 tik #5).

## Close — 2026-06-26

**Outcome:** the `--gen-batches` opt-in fence (silent-hollow guard) landed + the 429-backoff burst
resilience confirmed wired (ai-lib retry on-by-default + the wrapper's EU→direct fallback). 7 new tests +
full suite green; go.mod unchanged (1 vendored dep). All four code deliverables now done; only the
real-run gate-proving remains.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the org-scale gate is proven on the real-run sweep, TOK-01 tik #5 — iter-06)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks this session — iter-02/03/04/05; cap is 5, so tik #5 = the real-run sweep still fits in this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** (none beyond TOK-01)
**Side-deliverables:** none.
**Routes carried forward:** deliverable #5 — the REAL large-batch gate-proving sweep (the 5th + final tik
of this invocation, iter-06): fire a ~500-member Azure gpt-4o-mini batch, seed a stack, run the M42
manager+employee semantic-coverage sweep on the generated org, prove believable-populated + 0 collisions +
closure GREEN + budget.
**Lessons:** Exporting `ReservedHeroNames` as the single source of truth for the reserved-name set is what
makes the fence's expand() produce byte-identical mother prompts to gen-batch + the seeder — three call
sites keyed on the same cache must share ONE prompt-builder, or the fence would check a different cache dir
than the one gen-batch wrote. The 429 resilience was already structurally present (ai-lib default retry +
the wrapper fallback); this iter LOCKED it with tests rather than adding new machinery (the right move —
no new deps, no new retry layer).
