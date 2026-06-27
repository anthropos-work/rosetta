---
iter: 01
milestone: M45
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-06-26
---

# iter-01 — bootstrap tok (author the initial strategy + the iteration protocol)

## Type
tok (bootstrap) — iter-01 of M45, per Phase 0 rule 1. Authors the FIRST strategy (TOK-01) + the milestone's
iteration protocol doc. Does NOT terminate the call — the loop continues into iter-02 (a tik) under TOK-01.

## Inputs
- `overview.md` (scope, exit gate, In/Out), `spec-notes.md` (the per-component TODOs), the v1.10-extend
  research note.
- The KB-fidelity audit (GREEN) — the load-bearing contracts verified against code/docs.
- A live probe of the shared `ai` library (`v1.40.1`, fetched via GH_PAT) — the real `ai.AI` interface +
  `openai` constructors + `MetaData.Usage`/`CompletionUsage` token fields.
- The existing seeder machinery: `seeder.Seeder`/`Registry`/DAG, the resolvers (`jobroleref.go`,
  `skillref_named.go`), `blueprint.EffectiveStories()` (the normalization layer `EffectiveBatches()`
  mirrors).

## Initial strategy (→ TOK-01)
Build the engine **inside-out, fixtures-first**, in dependency order, so each layer is unit-tested against
a fixture `ai.AI` (no key, no cost) before the real-LLM gate-proving run:
1. `services/ai/` wrapper over the `ai` lib (routing + cost accounting + a fixture-injectable client).
2. `blueprint.Batch` + `batch[]` + `EffectiveBatches()` (pure Go-template expansion, no LLM).
3. The prompt-hash cache (key = mother-prompt + capture-version; atomic writes; `.lock`).
4. `cmd/gen-batch` (model + `--max-concurrent` + the mandatory `--max-cost` + re-roll).
5. `GeneratedBatchSeeder` (cache → rows via the existing resolvers; drop-not-fabricate; closure green).
6. Hero-name collision avoidance (prompt-side roster + post-gen re-roll).
7. The gate-proving real capped batch + the $0 re-seed.

The bootstrap tok itself authors the iteration protocol (`ai-generation-spec.md`) + `cache-spec.md`, runs
the KB-fidelity gate (GREEN), probes the `ai` lib fetchability (RC=0, `v1.40.1`), and resolves the
cache-doc placement (`corpus/ops/demo/cache-spec.md`, for demo-index consistency).

## Distance-to-gate context
Gate metric = the 5-vector {valid-JSON ≥95%, 0-fabrication + closure-green, 0 hero-collision, cost ≤
ceiling, $0 byte-identical re-seed}. Starting value: **engine does not exist yet** (0/5). The bootstrap
tok ships no gate-progress code; iter-02 is the first tik.

## Next-tik direction (iter-02)
Build component (1): the `services/ai/` wrapper — add the `ai` dep (the SANCTIONED supply-chain
inflection: license-vet + pin `v1.40.1`), implement EU-first routing + 429 fallback + the model→price
cost accounting, and a fixture `ai.AI` so the wrapper is unit-testable with no key. Measure: wrapper unit
tests green, cost accounting correct against a fixture usage payload, `go build`/`go vet` clean, dep
license-vetted. (No real-LLM call yet — that's a later tik once the prompt + seeder path exist.)

## Phase plan
Bootstrap tok: author protocol doc + cache doc, run KB-fidelity gate, probe the dep, record TOK-01.
No production engine code lands this iter (that starts iter-02).

## Escalation conditions
- The `ai` lib unreachable / unvettable → user-blocker (the milestone can't proceed without the dep).
  (Probed GREEN — RC=0, MIT/Apache transitive tree pending full vet at the iter-02 add.)

## Acceptable close-no-lift outcomes
N/A — a bootstrap tok closes `closed-fixed` when the strategy + protocol land (no metric to move).
