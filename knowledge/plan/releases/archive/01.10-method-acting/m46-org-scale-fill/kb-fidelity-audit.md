---
title: "KB Fidelity Audit — M46 (org-scale fill + gen-batch preview CLI)"
date: 2026-06-26
scope: milestone:M46
invoked-by: build-mstone-iters (Phase 0b, iter-01 bootstrap-tok pre-flight)
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| AI generation engine (Batch descriptor, gen-batch CLI, GeneratedBatchSeeder, ai wrapper) | `corpus/ops/demo/ai-generation-spec.md` | `stack-seeding/blueprint/batch.go`, `cmd/gen-batch/main.go`, `services/ai/*.go`, `seeders/generated_batch.go` | PAIRED |
| Prompt-hash batch cache | `corpus/ops/demo/cache-spec.md` | `stack-seeding/batchcache/cache.go` + `lock.go` | PAIRED |
| Demo-coverage Playwright semantic gate (the M46 iteration protocol) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/` (lib/crawl.ts, section-assert.ts, persona-assert.ts, coverage-manifest.ts) | PAIRED |
| Multi-org Stories & Heroes model + supporting-population fidelity | `corpus/ops/demo/stories-spec.md` | `stack-seeding/blueprint/stories.go`, `seeders/users.go`, `seeders/population_evidence.go` | PAIRED |
| Seeding blueprint + production-isolation boundary | `corpus/ops/seeding-spec.md` | `stack-seeding/isolation/`, `seeders/*` | PAIRED |

No BLIND-AREA topics: every M46-scope topic has both a doc anchor and code.

## Fidelity Findings (Phase 2)

All 9 audited load-bearing claims across the two M46-extended docs (`ai-generation-spec.md`,
`cache-spec.md`) verified **ALIGNED** (9/9):

1. Batch descriptor fields (count/roles/seniority/industry/narrative/bias_mix/prompt_template) — ALIGNED (`batch.go` Batch struct).
2. `EffectiveBatches()` pure, expands Batch → MOTHER prompts, no LLM at parse time — ALIGNED (`batch.go`).
3. gen-batch: `--max-cost` MANDATORY + aborts before breach; gpt-4o-mini default; `--max-concurrent` default 5; bounded re-roll; atomic cache + `.lock` fence — ALIGNED (`cmd/gen-batch/main.go`).
4. GeneratedBatchSeeder Surface/DependsOn/Isolation + cache→rows via resolvers + drop-not-fabricate — ALIGNED (`seeders/generated_batch.go`).
5. Hero-name collision avoidance: prompt-side reserved roster + code-side post-gen re-roll — ALIGNED.
6. services/ai EU-first Azure routing → 429 fallback → direct OpenAI; gpt-4o-mini; JSON mode; fixed seed; values-blind — ALIGNED (`services/ai/ai.go` + `cost.go`).
7. Cache layout (batch-${hash}/member-${i}.json + .lock + manifest.json) — ALIGNED.
8. Cache key = MOTHER prompt + taxonomy capture version; re-replay invalidates whole batch (coarse) — ALIGNED.
9. Atomic writes (.tmp→rename), .lock fence, cache hit = $0 — ALIGNED.

Cross-references in the M46-scope docs resolve. No stale numeric line-anchor citations found in the
audited docs (the docs cite symbols/sections, not `file:line`).

## Completeness Gaps (Phase 3) — the three M46 deliverable gaps (correctly OUT of M45, documented as M46 scope)

These are NOT stale findings — the docs explicitly mark them as M46 scope (`ai-generation-spec.md §5
"What's OUT"`). The code confirms they are not yet implemented, which is exactly correct for the M46
starting point:

- **(a) No org-size auto-fill.** `Batch.Count` is a fixed `int` (`batch.go:32`); there is no `Size`-aware
  "auto-fill to org size" expansion. M46 In-scope deliverable #1.
- **(b) `--dry-run` is a count-only summary.** `cmd/gen-batch --dry-run` prints `"%d to generate, %d
  cached"` to stdout (`main.go:110-113`); it does NOT render per-member prompts/cached JSON to a
  file/stdout nor an estimated-cost line. M46 In-scope deliverable #3 (the PREVIEW mode).
- **(c) GeneratedBatchSeeder hardcodes `stories[0]`.** Lines 118-124 resolve only the first story's
  org/prefix/domain; all generated members seed into the first org. The code comment itself says
  "per-story batch routing is an M46 concern." M46 In-scope deliverable #2 (per-story distribution).

Additionally noted (not a gap, a context confirmation): the existing structural supporting-population path
(`UsersSeeder` + `population_evidence.go`, per `stories-spec.md §Supporting-population fidelity`) fills a
population with role-coherent members + a ~55% claimed-vs-verified gap. M46's GENERATED supporting
population layers believable per-member English content (real names/bios/skill-names from the LLM) on top
of that same drop-not-fabricate / PerStackIsolated / closure-GREEN contract — it extends, does not
contradict, the documented seeding-isolation boundary.

## Applied Fixes
None required — all claims ALIGNED, all M46-scope gaps correctly documented as M46 deliverables.

(The overview.md `delivers` frontmatter cites `corpus/ops/cache-spec.md`; the file actually lives at
`corpus/ops/demo/cache-spec.md`. This is a 1-token path nit in the milestone's own overview frontmatter,
not a corpus-doc fidelity issue — left for the milestone's doc-update phase to correct alongside the
cache-spec edits, tracked here so it is not lost.)

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed. Every M46-scope topic is PAIRED, all 9 audited load-bearing claims ALIGNED, the three
completeness gaps are the M46 deliverables themselves (correctly documented as M46 scope, not stale).
The bootstrap tok may author its strategy against verified knowledge docs.
