**Type:** tik

# iter-03 — blueprint.Batch + batch[] + EffectiveBatches()

Built component (2) of TOK-01's chain: the generation engine's blueprint INPUT surface + the pure,
LLM-free per-member MOTHER-PROMPT expansion.

## Work done
- **NEW `blueprint/batch.go`** — the `Batch` type (id, count, roles[], seniority[], industry, narrative,
  bias_mix, prompt_template override); `BatchMember` (the resolved per-member unit + the rendered
  MotherPrompt); `EffectiveBatches(reservedNames)` — the batch analog of `EffectiveStories()`, expanding
  each story's batches into per-member units via a Go `text/template` (stdlib — NO new dep, NO LLM at
  parse time); `DefaultBatchPromptTemplate` (the JSON-envelope mother prompt with the reserved-hero-name
  avoidance instruction baked in); deterministic round-robin role/seniority + an INTERLEAVED bias
  distribution (largest-remainder, so a small span still sees the full mix); `MotherPromptHash` (the
  sha256 the cache keys on).
- **`blueprint/blueprint.go`** — added `Batches []Batch` (`yaml:"batch"`) to `StackSeed` (legacy path) +
  `Story` (multi-story path); a `validateBatches` helper (id non-empty+unique, count>=0, bias-mix enum,
  template parses) wired into both `validateLegacy` + `validateStories`. Additive — every existing path
  unchanged.
- **NEW `blueprint/batch_test.go`** — 12 tests: legacy expansion (role/seniority round-robin, prompt
  content), determinism + hash stability, bias distribution (interleave), multi-story, no-batches→0,
  custom-template override, bad-template error, validation errors + valid-passes, YAML `batch:` parse.

## Measurements
- `go build ./...` + `go vet ./...` clean; `gofmt -l blueprint/` clean.
- `go test ./blueprint/...` GREEN; full `go test ./...` GREEN (incl. cmd/stackseed — the `KnownFields`
  parse of every preset unaffected by the additive field). Test funcs 587 -> **599** (+12).

## Close — 2026-06-26

**Outcome:** Component (2) landed — the batch descriptor surface + the pure, deterministic, LLM-free
mother-prompt expansion (the $0-reseed foundation). 12 unit tests; full suite green; no existing path
touched. Gate still 0/5 (the empirical valid-JSON rate needs a real call — components 3-7 + the
gate-proving run).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (0/5 — input surface assembled; the LLM-firing CLI + the seeder + the real run remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (text/template stdlib, no new dep), D2 (interleaved bias distribution)
**Side-deliverables (if any):** none
**Routes carried forward:** iter-04 builds component (3): the prompt-hash cache (key = mother-prompt +
capture-version; atomic .tmp->rename; .lock fence) under TOK-01. (Then cmd/gen-batch is component 4.)
**Lessons:** a block-run weighted sequence bunches biases at the start of a small batch span — use an
INTERLEAVED (largest-remainder) assignment so a 3-5 member batch still shows the declared mix. Keeping the
expansion a PURE function (no I/O, no LLM) is what makes the cache key deterministic + the $0 reseed
possible.
