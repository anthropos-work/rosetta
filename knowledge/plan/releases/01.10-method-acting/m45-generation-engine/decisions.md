# M45 Decisions

Implementation decisions with rationale (recorded during build/iters). Design-time decisions live in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| _(intra-iter decisions live in each `iter-NN/decisions.md`)_ | | | |

---

## TOK-01: inside-out fixtures-first build — 2026-06-26

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Build the generation engine **inside-out, in dependency order, fixtures-first** —
each layer unit-tested against a fixture `ai.AI` (no key, no cost) before the real-LLM gate-proving run:
1. `services/ai/` — the thin wrapper over the `ai` lib (`v1.40.1`): EU-first routing → 429 fallback → a
   model→price cost accounting; a fixture-injectable client so the wrapper is testable with no key.
2. `blueprint.Batch` + `batch[]` on Story/StackSeed + `EffectiveBatches()` — pure Go-template per-member
   prompt expansion (NO LLM at parse time), mirroring the existing `EffectiveStories()` normalization.
3. The prompt-hash cache (`.agentspace/.batchcache/batch-${hash}/member-${i}.json`) — key = mother-prompt
   + taxonomy-capture-version; atomic `.tmp`→rename; `.lock` fence.
4. `cmd/gen-batch` — gpt-4o-mini default, `--max-concurrent` (5), the MANDATORY `--max-cost` ceiling,
   re-roll on malformed, the per-batch cost report.
5. `GeneratedBatchSeeder` (surface `'generated-batch'`, DependsOn users+taxonomy, PerStackIsolated) —
   cache → rows DETERMINISTICALLY via the existing resolvers; a non-resolving name DROPS (closure green).
6. Hero-name collision avoidance — reserved-hero roster in the prompt + a post-gen collision re-roll.
7. Gate-proving: a real N=20 capped batch (`OPENAI_KEY` from `stack-demo/platform/.env`, values-blind) +
   the $0 byte-identical re-seed; then re-seed on a fresh demo-3 via the tagged consumption clone.
**Rationale:** the CODE-owns-structure / AI-owns-content boundary (the design's spine) means every layer
EXCEPT the model's English content is deterministic and unit-testable. Building inside-out lets each
deterministic layer be proven in CI against fixtures before spending a single token; the one empirical
unknown (the cheap model's raw valid-JSON / resolution rate) is isolated to the gate-proving run at the
end. The `ai` lib's split (transport in the lib; routing+cost in the consumer) keeps `services/ai/` a thin
consumer — no provider-internal reimplementation.
**Strategy class:** new-direction
**Distance-to-gate context:** Gate = the 5-vector {valid-JSON ≥95% (pre-re-roll), 0-fabrication +
closure-green, 0 hero-collision, cost ≤ `--max-cost`, $0 byte-identical re-seed}. Starting value: the
engine does not exist (0/5). The KB-fidelity gate is GREEN and the `ai` dep is fetchable (`v1.40.1`,
RC=0), so the path is unblocked.
**Next-tik direction:** iter-02 (tik) — build component (1), `services/ai/`: add + license-vet + pin the
`ai` dep (the sanctioned supply-chain inflection), implement routing/fallback + the model→price cost
accounting, ship a fixture `ai.AI` + wrapper unit tests (no real call yet). Measure: wrapper tests green,
cost accounting correct against a fixture usage payload, build/vet clean, dep vetted.
