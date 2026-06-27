---
title: "KB Fidelity Audit — milestone:M45 generation-engine"
date: 2026-06-26
scope: milestone:M45
invoked-by: build-mstone-iters
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| The shared `ai` library (interface, constructors, what it does/doesn't do) | `corpus/architecture/shared_libraries.md` §ai + `corpus/architecture/ai_architecture.md` | `github.com/anthropos-work/ai` (external private module — the SANCTIONED new dep) | PAIRED |
| EU-first routing / 429 fallback / usage tracking (the `services/ai/` wrapper contract) | `corpus/architecture/ai_architecture.md` §Provider Routing + shared_libraries.md §ai | each consumer's `internal/ai/ai.go` (read-only ref); `Event_AiUsage` cost model in `app/internal/aiusage` (read-only ref) | PAIRED |
| The seeding blueprint + production-isolation boundary (GeneratedBatchSeeder is a new surface within it) | `corpus/ops/seeding-spec.md` | `stack-seeding/seeder/` (Seeder iface + Registry + DAG), `stack-seeding/isolation/` | PAIRED |
| The verified-skill chain + the resolvers (real public node-ids, drop-not-fabricate) | `corpus/ops/demo/stories-spec.md` | `stack-seeding/seeders/jobroleref.go`, `skillref_named.go`, `taxonomyref.go`, `persona.go`, `profile.go` | PAIRED |
| The replayed public taxonomy + the capture version (the cache-key extension) | `corpus/ops/snapshot-spec.md` | `stack-snapshot/` (capture/replay + manifest version) | PAIRED |
| `corpus/ops/demo/ai-generation-spec.md` (NEW — gen-acceptance protocol + engine) | — (to be authored) | — (to be built) | DOC-ONLY (Delivers target — milestone authors it) |
| `cache-spec.md` (NEW — prompt-hash cache) | — (to be authored) | — (to be built) | DOC-ONLY (Delivers target — milestone authors it) |

## Fidelity Findings

### ai library — module path + interface + version pin
- **Source:** `corpus/architecture/shared_libraries.md` §ai (lines 96–132)
- **Expected:** module `github.com/anthropos-work/ai`, pinned `v1.40.1` across consumers, Go 1.25.0; one interface `ai.AI` (`ChatCompletion`, …); constructors `openai.New`/`NewOpenAI` (direct), `openai.NewAzure` (Azure EU); `MetaData.Usage` carries token counts only.
- **Actual:** the module is external (not yet a rext dep — confirmed `grep -rl anthropos-work/ai */go.mod` returns nothing, the 0-new-deps streak is intact). The doc is the contract the `services/ai/` wrapper compiles against.
- **Verdict:** ALIGNED — doc is the authoritative contract; M45 will pin the same `v1.40.1` (or the consumers' current pin) and import `openai.NewAzure`/`openai.New`.
- **Fix owner:** n/a.

### ai library — EU-first routing + cost tracking live in the CONSUMER, not the lib
- **Source:** `corpus/architecture/ai_architecture.md:42`, `shared_libraries.md:117-125`
- **Expected:** the `ai` library does NOT do EU-first routing or cost tracking; routing/fallback (EU Azure default → US Azure via `flag_use_azure_us` → direct-OpenAI on 429; Anthropic always Bedrock eu-west-1) lives in each consumer's `internal/ai/ai.go`; dollar cost is computed by the consumer (`app/internal/aiusage`, a model→price switch) and emitted via `Event_AiUsage` over Redis Streams.
- **Actual:** matches the overview's design — M45's `services/ai/` wrapper is exactly such a consumer wrapper (EU-first → 429 fallback → usage tracking layered on the bare `ai` lib). The wrapper must NOT reimplement provider internals; it's a thin consumer.
- **Verdict:** ALIGNED — this claim is the load-bearing one for the milestone and it is accurate. The `services/ai/` wrapper owns routing+usage; the lib owns transport.
- **Fix owner:** n/a.

### resolvers — drop-not-fabricate seam
- **Source:** `corpus/ops/demo/stories-spec.md` (TaxonomyRefs resolver — real public node-ids, never fabricated)
- **Expected:** every role/skill name resolves to a real public node-id; a non-resolving name yields a blank/zero result (drop), never a fabricated `J-…`/`K-…`.
- **Actual:** `jobroleref.go` `forName()` returns a zero `jobRole{}` for an unmatched name; `skillref_named.go` resolvers return empty pools on a missing/unmatched taxonomy. Confirmed in code — the seam the AI-generated names feed through is real and drop-safe.
- **Verdict:** ALIGNED — this is WHY a cheap model is safe; a hallucinated skill name just drops, worst case a shallower profile.
- **Fix owner:** n/a.

### seeder framework — the Seeder/Registry/DAG contract for a new surface
- **Source:** `corpus/ops/seeding-spec.md` (the blueprint + production-isolation boundary)
- **Expected:** a new surface declares `Surface()` (unique name), `DependsOn()` (predecessor surfaces), `Isolation()` (class), and `Seed(ctx, Conn, *StackSeed, *AuditLog)`; the DAG topo-sorts + gates every write through the isolation Guard.
- **Actual:** `seeder/seeder.go` defines exactly this interface; `cmd/stackseed/main.go:buildRegistry()` is the static wiring point. `GeneratedBatchSeeder` (surface `'generated-batch'`, DependsOn users+taxonomy, PerStackIsolated) slots in cleanly.
- **Verdict:** ALIGNED.
- **Fix owner:** n/a.

## Completeness Gaps

None critical. The two NEW docs (`ai-generation-spec.md`, the cache section/`cache-spec.md`) are DOC-ONLY by design — they are the milestone's deliverables, authored by the bootstrap tok + deepened across tiks. Their absence now is expected, not a blind area: the milestone `overview.md` `Delivers →` block explicitly makes their production a milestone deliverable.

**Placement note (resolve in bootstrap tok):** the overview's Delivers line says `corpus/ops/cache-spec.md`; the build-iter prompt + CLAUDE.md index convention places demo-family docs under `corpus/ops/demo/`. The bootstrap tok will decide placement (likely `corpus/ops/demo/cache-spec.md` for index-consistency, or a Caching section inside `ai-generation-spec.md` per the overview's stated allowance) and record it as a decision.

## Applied Fixes
None needed — all PAIRED topics ALIGNED; no stale claims, no frontmatter drift in the in-scope docs, no broken cross-refs found.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed. All milestone-scope topics are PAIRED with aligned, load-bearing claims; the two NEW docs are intentional DOC-ONLY deliverables (made explicit in the overview's `Delivers →`), not blind areas. The bootstrap tok authors its strategy against verified knowledge docs.
