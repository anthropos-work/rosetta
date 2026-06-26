# M45 Progress

## Running ledger
Iter closeouts append here (one line each — the tik/tok, what the batch run measured against the exit
gate [valid-JSON rate, taxonomy-resolution, collisions, cost vs ceiling, the $0 re-seed], and what was
fixed/hardened).

- iter-01 (tok/bootstrap): authored the iteration protocol (`ai-generation-spec.md`) + `cache-spec.md`; KB-fidelity GREEN; `ai` dep (`v1.40.1`) fetchable; TOK-01 (inside-out fixtures-first build) recorded — see iter-01/progress.md
- iter-02 (tik): component (1) — the values-blind EU-first cost-tracking `services/ai/` wrapper + the sanctioned `ai` dep (v1.40.1, all-permissive tree); 20 unit tests; 567→587; full suite green — see iter-02/progress.md
- iter-03 (tik): component (2) — `blueprint.Batch` + `batch[]` + `EffectiveBatches()` (pure Go-template mother-prompt expansion, NO LLM at parse time, deterministic $0-reseed foundation); 12 tests; 587→599; full suite green — see iter-03/progress.md
- iter-04 (tik): component (3) — the `batchcache/` prompt-hash cache (atomic .tmp→rename, .lock fence, capture-version invalidation; $0 byte-identical reseed proven in unit); 14 tests; 599→613; full suite green — see iter-04/progress.md
- iter-05 (tik): component (4) — `cmd/gen-batch` (the generation CLI: mandatory --max-cost ceiling, --max-concurrent semaphore, re-roll-on-malformed, hero-collision re-roll, $0 cache reseed, dry-run, lock fence; ALL fixture-proven, no key); 10 tests; 613→623; full suite green — see iter-05/progress.md
- iter-06 (tik): component (5) — the `GeneratedBatchSeeder` (cache → users/memberships/claimed-skills via the existing resolvers; the CODE-vs-AI drop-not-fabricate boundary, unit-proven; registered in the DAG); 9 tests; 623→632; full suite + race green. **Engine CODE-COMPLETE** — next call = the REAL gate-proving batch — see iter-06/progress.md
- iter-07 (tik): the REAL gate-proving run (EU-first Azure / Sweden, gpt-4o-mini, N=20) — **gate 0/5 → 5/5 PASS**. valid-JSON 100% (33/33) · 47/47 skills + 20/20 roles → real `skiller.node_id`, `datadna measure-closure --stack demo-3` = `[PASS]` (closure GREEN, 0 fabrication) · 0 hero-collisions · cost $0.0059 ≤ $0.10 · $0 byte-identical re-seed. Surfaced + fixed 3 issues: per-call `--call-timeout` (the stall class), intra-batch name dedup (→ 20/20 distinct multicultural names via prompt+re-roll-hint), the `user_skills` CHECK 23514 (→ ONE current-role experience per gen member). +1 test; rext tag `method-acting-m45-iter07-gate`, demo-3 consumption clone bumped. **GATE MET** — see iter-07/progress.md

**Exit gate:** on a real batch of N — valid JSON ≥95% (re-roll on malformed), every role/skill name
resolves to a real public-taxonomy node-id (non-resolving drop, closure green), ZERO generated name
collides with a hand-curated hero, total cost within `--max-cost`; reproducible byte-identical from cache
at $0.

**Budget:** 3–5 iters. **Re-scope trigger:** if the cheap model can't reach the valid-JSON /
taxonomy-resolution threshold after ~5 tiks of prompt+code hardening → user-strategic-replan (model
upgrade vs scope reduction).
