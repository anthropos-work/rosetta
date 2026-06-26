# M45 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md)
(the generation-engine strand). M45 is **entirely in `rosetta-extensions` — zero platform-repo edits**. It
breaks the seeding module's 0-new-deps streak (the FIRST new third-party dep — a deliberate in-release
decision). The iteration protocol is the NEW `corpus/ops/demo/ai-generation-spec.md` (this milestone
authors it).

## `services/ai/` wrapper (the first new third-party dep)
TODO: a local wrapper around the shared `ai` library — EU-first routing → 429 fallback → usage tracking.
Do not reimplement provider routing (the shared `ai` lib + each consumer's wrapper own it); this is the
seeding-module's thin consumer.

## `blueprint.Batch` + `batch[]` + `EffectiveBatches()`
TODO: a `Batch` type + a `batch[]` field on `Story` / `StackSeed`; `EffectiveBatches()` does Go-template
per-member-prompt expansion at PARSE time — NO LLM at parse time (the LLM runs only in `cmd/gen-batch`).

## `cmd/gen-batch` CLI
TODO: gpt-4o-mini default; `--max-concurrent` semaphore (default 5); atomic `.tmp`→rename cache writes; a
`.lock` fence; a MANDATORY `--max-cost` ceiling (abort before exceeding). Emits `Event_AiUsage` + a
per-batch cost report.

## Prompt-hash cache (→ `cache-spec.md` or an ai-generation-spec.md Caching section)
TODO: `.agentspace/.batchcache/batch-${hash}/member-${i}.json`, keyed by the MOTHER prompt; cache key
EXTENDED with the taxonomy capture version (invalidate on re-replay). Atomic writes; the `.lock` fence.
Reproducibility target: an unchanged descriptor re-seeds byte-identical from cache at $0.

## `GeneratedBatchSeeder` (surface `'generated-batch'`)
TODO: `DependsOn` users+taxonomy, `PerStackIsolated`. Reads cached JSON; drives users/persona/profile rows
DETERMINISTICALLY via the existing resolvers (`resolveJobRoleRefs` / `resolveNamedSkillRefs`). The
CODE-owns-structure / AI-owns-content boundary: a non-resolving role/skill name DROPS (worst case a
SHALLOWER profile, never invalid); the closure gene (`datadna measure-closure`) stays green.

## Hero-name collision avoidance
TODO: reserved-hero-names in the prompt (the curated hero roster) + a post-gen collision re-roll fallback.
Exit gate requires ZERO generated name colliding with a hand-curated hero.

## Exit-gate measurement (per iter)
TODO: per-batch metrics — valid-JSON rate (≥95% + re-roll), taxonomy-resolution rate (every name resolves
or drops, closure green), collision count (=0), cost vs `--max-cost`. Plus the byte-identical-from-cache
re-seed at $0.

## Secrets (deferred for production)
`OPENAI_API_KEY` via `.env.local` env var (git-ignored) for now. NO platform-repo secrets store;
production-seeding key story deferred (an Open question).
