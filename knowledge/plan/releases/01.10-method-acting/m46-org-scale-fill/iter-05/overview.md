---
iter: 05
milestone: M46
iteration_type: tik
status: closed-fixed
created: 2026-06-26
---

# M46 · iter-05 (tik #4) — --gen-batches fence + 429/throughput verification (deliverable #4)

**Active strategy reference:** TOK-01. Tik #4 = deliverable #4 (the `--gen-batches` opt-in fence + the
throughput/429 backoff verification).

**Re-survey (Step 0):** TOK-01 named the fence + the throughput verification. The audit confirmed no
`--gen-batches` flag exists; the `GeneratedBatchSeeder` silently writes 0 rows on an empty cache (correct
for a no-batch stack, dangerous for a batch[] stack). Still untouched, still meaningful.

**Cluster / target identified:** (a) the silent-hollow failure mode — a batch[] stack whose cache is empty
(OpenAI/Azure was unreachable when gen-batch ran) seeds a hollow org with no signal; (b) the 429
burst-resilience — verify the ai-lib retry/backoff holds at `--max-concurrent=5`.

**Hypothesis:** (a) add an opt-in `--gen-batches` flag on `stackseed` that asserts the prompt-hash cache is
COMPLETE for a batch[] blueprint before any write, failing loud otherwise. (b) verify (by inspection +
locking tests) the ai-lib's 429 retry + the wrapper's EU→direct fallback are wired.

**Expected lift:** the fence lands (a batch[] stack can't silently seed hollow under `--gen-batches`); the
429 backoff is confirmed on-by-default. (Gate metric on a generated org unmoved until the real-run tik.)

**Phase plan:** code (`--gen-batches` flag + `assertBatchCacheComplete` in stackseed; export
`seeders.ReservedHeroNames` as the one source of truth for the cache key) → 429 verification (inspect the
ai-lib `DefaultRetryOptions` + the wrapper's `CompleteJSON` fallback; lock with tests) → fixtures-first
tests (fence: no-batch passes / empty fails / complete passes / partial fails; 429: fallback fires /
non-429 doesn't / no-fallback propagates) → full suite green + go.mod unchanged.

**Escalation conditions:** if the 429 backoff were NOT wired (a real burst would drop calls), that would be
a real gap needing a wrapper change before the real-run tik. (It IS wired — verified.)

**Acceptable close-no-lift outcomes:** N/A — deterministic code + a verification.
