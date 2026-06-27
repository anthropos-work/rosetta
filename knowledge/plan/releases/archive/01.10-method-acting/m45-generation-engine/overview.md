---
milestone: M45
slug: generation-engine
version: v1.10 "method acting"
milestone_shape: iterative
exit_gate: "On a real batch of N generated members, the engine seeds end-to-end: the cheap model emits valid JSON ≥95% of calls (re-roll on malformed), every role/skill name resolves to a REAL public-taxonomy node-id via the existing resolvers (non-resolving names drop, closure gene stays green), ZERO generated name collides with a hand-curated hero, and total cost lands within the declared --max-cost ceiling. Reproducible: an unchanged batch descriptor re-seeds byte-identical from cache at $0."
iteration_protocol_ref: corpus/ops/demo/ai-generation-spec.md
status: archived
created: 2026-06-26
last_updated: 2026-06-26
complexity: large
delivers: corpus/ops/demo/ai-generation-spec.md (NEW — the gen-acceptance protocol + the AI generation engine: the cmd/gen-batch CLI + GeneratedBatchSeeder + the CODE-owns-structure / AI-owns-content boundary) + corpus/ops/cache-spec.md (NEW, or a Caching section in ai-generation-spec.md — the prompt-hash cache + capture-version invalidation)
depends_on: M44 (reuses the certificate/project + bulk-member surfaces + the trajectory-aware self-rating); soft-depends on a replayed public taxonomy with adequate coverage
spec_ref: .agentspace/scratch/roadmap-research-2026-06-26.md (the v1.10-extend research note; the generation-engine strand)
---

# M45 — Generation engine (LLM batch profile generation + prompt-keyed cache)

## Goal
Stand up the **AI generation engine** — a `cmd/gen-batch` CLI + `GeneratedBatchSeeder` that turns a
high-level **YAML batch descriptor** into realistic **per-member profiles** via a **CHEAP LLM**
(gpt-4o-mini default), with **parallel/throttled** generation, **prompt-hash caching**, and the
**CODE-owns-structure / AI-owns-content** boundary enforced by the existing taxonomy closure gate. This
breaks the seeding module's 0-new-deps streak (the first new third-party dep — a deliberate in-release
decision) and is the **scalable-generation** extension of v1.10. **Entirely in `rosetta-extensions` — zero
platform-repo edits.**

## Exit gate (observable, machine-verifiable)
On a **real batch of N generated members**, the engine seeds end-to-end:
- the cheap model emits **valid JSON ≥95%** of calls (**re-roll** on malformed);
- **every role/skill name resolves to a REAL public-taxonomy node-id** via the existing resolvers
  (non-resolving names **drop**; the closure gene stays **green**);
- **ZERO** generated name collides with a hand-curated hero;
- total cost lands **within the declared `--max-cost` ceiling**.

**Reproducible:** an unchanged batch descriptor **re-seeds byte-identical from cache at $0**.

## Why iterative (not section)
**Greenfield LLM integration** — the first new dep, prompt-engineering hardening, cache-invalidation +
collision-dedup edge cases. The valid-JSON / taxonomy-resolution thresholds are reached by a
**measure→fix→accept loop** on generation quality (prompt + code hardening), not enumerable up front. The
commitment is the **gate**; the fix list emerges per-iter from the run's evidence. **Budget 3–5 iters.**

## Iteration protocol
[`corpus/ops/demo/ai-generation-spec.md`](../../../../corpus/ops/demo/ai-generation-spec.md) — the
milestone **AUTHORS + deepens** this gen-acceptance protocol (the measure→fix→accept loop on generation
quality). Each iter: run a real batch → measure (valid-JSON rate, taxonomy-resolution rate, collision
count, cost vs ceiling) → harden the prompt / code → re-run → close the iter.

**Re-scope trigger:** if the cheap model **cannot reach** the valid-JSON / taxonomy-resolution exit
threshold after **~5 tiks** of prompt + code hardening, escalate to a **user-strategic-replan** (model
upgrade vs scope reduction).

## Scope
**In:**
- A local **`services/ai/` wrapper** around the shared `ai` library (EU-first routing → 429 fallback →
  usage tracking) — the **FIRST new third-party dep** in the seeding module.
- A `blueprint.Batch` type + a `batch[]` field on `Story` / `StackSeed` + `EffectiveBatches()` (Go-template
  per-member-prompt expansion, **NO LLM at parse time**).
- **`cmd/gen-batch`** — gpt-4o-mini default; `--max-concurrent` semaphore (default 5); atomic
  `.tmp`→rename cache writes; a `.lock` fence; a **MANDATORY `--max-cost` ceiling**.
- A **prompt-hash cache** at `.agentspace/.batchcache/batch-${hash}/member-${i}.json`, keyed by the
  **MOTHER prompt**, the cache key **EXTENDED with the taxonomy capture version** (invalidate on
  re-replay).
- **`GeneratedBatchSeeder`** (surface `'generated-batch'`, `DependsOn` users+taxonomy,
  `PerStackIsolated`) — reads cached JSON and drives users/persona/profile rows **DETERMINISTICALLY** via
  the existing resolvers (closure green; non-resolving names **drop** → worst case a **shallower** profile,
  never invalid).
- **Hero-name collision avoidance** — reserved-hero-names in the prompt + a **post-gen collision re-roll**
  fallback.
- `Event_AiUsage` cost emission + a **per-batch cost report**.

**Out:**
- **Org-scale auto-fill** to reach full org size — that's **M46** (M45 proves the engine + cache on a
  **bounded** batch).
- **Deep per-generated-member work-history/education timelines** — kept shallow (name + skills + bio +
  role).
- A **platform-repo secrets store** — key via `.env.local` `OPENAI_API_KEY` env var (git-ignored);
  production-seeding secrets deferred.
- Any **platform-repo edit** — entirely in `rosetta-extensions`.

## Depends on / Parallel with
- **Depends on:** **M44** (reuses the certificate/project + bulk-member surfaces + the trajectory-aware
  self-rating). **Soft-depends** on a replayed public taxonomy with adequate coverage.
- **Parallel with:** none.

## Open questions
- **gpt-4o-mini Azure-EU availability** vs US fallback (the EU-first routing's behaviour for this model).
- **Cache-invalidation granularity** (default coarse: invalidate on taxonomy re-replay via the
  capture-version in the key).
- The **`--max-cost` ceiling value** (mandatory per the user — default a sane per-batch ceiling,
  configurable).
- The **valid-JSON exit threshold** (default 95% + re-roll).
- **Cache location** (default `.agentspace/.batchcache` per-box; a shared/committed store flagged).
- The **production-seeding key story** (deferred; `.env.local` env var for now).

## KB dependencies
M45 reads these corpus docs as contract (it must not contradict them; it extends them):
- `corpus/ops/seeding-spec.md` — the seeding blueprint + the production-isolation boundary; the
  `GeneratedBatchSeeder` is a new surface within that contract.
- `corpus/ops/demo/stories-spec.md` — the verified-skill chain + the resolvers (real public node-ids,
  never fabricated). The CODE-owns-structure / AI-owns-content boundary routes every generated name
  through these resolvers (drop-not-fabricate).
- `corpus/ops/snapshot-spec.md` — the replayed public taxonomy + the **capture version** the cache key is
  extended with (invalidate on re-replay).
- `corpus/architecture/ai_architecture.md` + `corpus/architecture/shared_libraries.md` — the shared `ai`
  library (the EU-first routing the new `services/ai/` wrapper layers on) + the usage-tracking model
  (`Event_AiUsage`).

## Delivers →
- `corpus/ops/demo/ai-generation-spec.md` **(NEW)** — the **gen-acceptance protocol** (measure→fix→accept
  on generation quality) **+** the AI generation engine: the `services/ai/` wrapper, `blueprint.Batch` +
  `EffectiveBatches()`, `cmd/gen-batch` (model + concurrency + `--max-cost` + the cache), the
  `GeneratedBatchSeeder`, the CODE-owns-structure / AI-owns-content boundary, hero-collision avoidance,
  and the cost report.
- `corpus/ops/cache-spec.md` **(NEW, or a Caching section in `ai-generation-spec.md`)** — the prompt-hash
  cache (`.agentspace/.batchcache/batch-${hash}/member-${i}.json`), the MOTHER-prompt key + the
  capture-version extension (invalidate on re-replay), atomic `.tmp`→rename writes, the `.lock` fence.
