**Type:** tik

# iter-02 — the services/ai wrapper + the sanctioned `ai` dep

Built component (1) of TOK-01's chain: the stack-seeding module's thin wrapper over the shared `ai`
library, and landed the SANCTIONED first new third-party dep.

## Work done
- **NEW `services/ai/ai.go`** — the thin wrapper: `Completion` interface (`CompleteJSON` + `Model`);
  `Client` with EU-first routing (Azure EU primary when `AZURE_OPENAI_*` set, else direct OpenAI) + a
  429 -> direct-OpenAI fallback; `NewFromEnv` reads the key VALUES-BLIND from env; `usageFromMeta`
  type-asserts the lib's `MetaData.Usage` (a `libai.Usage` = `openai.CompletionUsage`) into plain ints;
  `redactErr`/`scrubSecrets` defend the values-blind contract (any env key value -> [REDACTED]).
- **NEW `services/ai/cost.go`** — the usage-tracking half: a model->price table (gpt-4o-mini load-bearing
  + 3 others), `CostOf`, and a thread-safe `CostTracker` with `WouldExceed` (the mandatory `--max-cost`
  ceiling guard, fails closed on ceiling<=0) + a values-blind per-batch `Report`.
- **NEW `services/ai/fixture.go`** — `FixtureCompletion` (scripted/canned envelopes, recorded prompts +
  seeds, fixed usage) so the whole engine is unit-testable with NO key, NO cost (fixtures-first §4c).
- **NEW `services/ai/ai_test.go`** — 20 tests: usage normalization (value/pointer/nil/unknown), is429,
  redaction (no key leak), NewFromEnv routing (no-key error / direct / Azure+fallback / model default),
  the fixture (scripted-then-fallback, seeds, last-user), the cost table (gpt-mini math, prefix/unknown),
  and the cost tracker (ceiling guard, report, cache-hits, unpriced flag, concurrent-add race).
- **Dep:** `github.com/anthropos-work/ai v1.40.1` added (values-blind via GH_PAT); `go mod tidy`.
  License-vet: the new third-party transitive tree is all permissive (openai-go/v3 Apache; retry-go,
  tidwall gjson/sjson, samber/lo + go-gpt-3-encoder, json-iterator, regexp2 MIT; modern-go/reflect2
  Apache; Azure SDK / MSAL-Go / golang-jwt MIT). `anthropos-work/ai` itself is a FIRST-PARTY Anthropos
  internal module (the same one the platform consumers use), not a third-party license concern.

## Measurements
- `go build ./...` RC=0; `go vet ./services/ai/...` RC=0; `gofmt -l services/ai/` clean.
- `go test ./services/ai/...` GREEN (20 new tests). Full module `go test ./...` GREEN (no pre-existing
  test broken by the dep add). Test funcs 567 -> **587** (+20).

## Close — 2026-06-26

**Outcome:** Component (1) landed — the values-blind EU-first cost-tracking `ai` wrapper + the sanctioned
dep, 20 unit tests green, full suite green, dep license-vetted (all permissive). Gate metric unchanged
(infrastructure tik; the empirical valid-JSON rate needs the prompt + a real call in later tiks).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (0/5 — engine still being assembled; this is the foundation layer)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (dep license-vet record, iter-02 local)
**Side-deliverables (if any):** none
**Routes carried forward:** iter-03 builds component (2): `blueprint.Batch` + `batch[]` + `EffectiveBatches()` (pure Go-template per-member prompt expansion, NO LLM at parse time), under TOK-01.
**Lessons:** the lib sets `MetaData.Usage` to a `libai.Usage` VALUE (not pointer) on the openai path — assert both forms defensively. The lib pulls a heavy transitive tree (Azure SDK, openai-go/v3); all permissive, but the supply-chain footprint grew notably — note for the close-time supply-chain review.
