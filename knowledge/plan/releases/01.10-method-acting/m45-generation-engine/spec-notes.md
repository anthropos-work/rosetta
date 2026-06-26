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

## Pre-flight audits — iter-01
**`/developer-kit:audit-kb-fidelity --milestone=M45` → GREEN** (2026-06-26; report
[`kb-fidelity-audit.md`](kb-fidelity-audit.md)). All 5 KB-dependency topics PAIRED + ALIGNED; the 2 NEW
docs (`ai-generation-spec.md`, `cache-spec.md`) are intentional DOC-ONLY Delivers targets, not blind
areas. Load-bearing contracts confirmed against code/docs:
- the `ai` library is `github.com/anthropos-work/ai` (pinned `v1.40.1` across consumers, Go 1.25.0);
  interface `ai.AI` (`ChatCompletion`, …); constructors `openai.New`/`NewOpenAI` (direct),
  `openai.NewAzure` (Azure EU). `MetaData.Usage` carries **token counts only** — dollar cost is the
  CONSUMER's job (`app/internal/aiusage`, model→price switch, `Event_AiUsage` over Redis Streams).
- **EU-first routing + cost tracking are NOT in the lib** — they live in each consumer's
  `internal/ai/ai.go`. The M45 `services/ai/` wrapper is exactly such a thin consumer (do not
  reimplement provider internals).
- the resolver seam is real + drop-safe: `jobroleref.go forName()` → zero `jobRole{}` on no-match;
  `skillref_named.go` → empty pools on missing/unmatched taxonomy. A hallucinated name DROPS.
- the Seeder/Registry/DAG contract (`seeder/seeder.go` + `cmd/stackseed/main.go buildRegistry()`) takes
  a new `'generated-batch'` surface cleanly (DependsOn users+taxonomy, PerStackIsolated).
- supply-chain: `grep -rl anthropos-work/ai */go.mod` → empty (0-new-deps streak intact; M45 breaks it
  deliberately).

**Placement decision pending bootstrap tok:** overview's Delivers says `corpus/ops/cache-spec.md`; the
demo-family index convention puts demo docs under `corpus/ops/demo/`. Bootstrap tok resolves
placement (TOK-01).
