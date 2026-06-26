# M46 Progress

## Running ledger
Iter closeouts append here (one line each — the tik/tok, what the semantic-coverage sweep measured on the
generated org [believable spread vs hollow, gate PASS/FAIL, collision count, throughput + cost vs budget],
and what was tuned/fixed).

<!-- iter-NN/ dirs are created by /developer-kit:build-mstone-iters on first run. -->

- iter-01 (tok·bootstrap): authored TOK-01 (build 3 deliverables fixtures-first → prove on a real ~500-member org via the M42 semantic sweep); Phase 0b KB-fidelity GREEN — see iter-01/progress.md
- iter-02 (tik): deliverable #1 — auto-fill count (`Batch.Fill`/`resolveBatchCounts`: one descriptor fills a story to its Size, per-story, deterministically); 8 fixtures-first tests + full stack-seeding suite green; gate metric unmoved (proven on the real-run sweep, tik #5) — see iter-02/progress.md
- iter-03 (tik): deliverable #2 — per-story batch distribution (`BatchMember.StoryIndex` + per-member story routing in GeneratedBatchSeeder; was hardcoded stories[0]); composes with the iter-02 fill (each org fills to its OWN Size); 2 new seeder tests + full suite green — see iter-03/progress.md
- iter-04 (tik): deliverable #3 — gen-batch preview mode (`--preview`/`--preview-out`: renders per-member prompts + cached JSON + a per-member + total estimated-cost line WITHOUT seeding; implies --dry-run, no LLM, no key, values-blind); 4 fixtures-first tests + full suite green + a real offline smoke vs the 20-member preset ($0.0062 est) — see iter-04/progress.md
- iter-05 (tik): deliverable #4 — `--gen-batches` opt-in fence on stackseed (a batch[] stack with an empty/incomplete cache fails LOUD before any write; off by default) + 429-backoff verification (ai-lib v1.40.1 DefaultRetryOptions on-by-default + the wrapper EU→direct fallback, locked with tests); exported `seeders.ReservedHeroNames` (one cache-key source of truth); 7 new tests + full suite green; go.mod unchanged — see iter-05/progress.md
- iter-06 (tik): deliverable #5 (part 1) — REAL 614-member Azure gpt-4o-mini gate-proving SURFACED + FIXED 2 org-scale bugs: the multi-batch cache-index collision (lost the whole 2nd story's 117 members) + name-distinctness at scale (58% distinct → a deterministic disambiguator guarantees 100%), both regression-tested. PROVEN on the real org: 0 hero-collisions, 100% valid-JSON, $0 cache-hit reseed, the --max-cost guard aborting 3× at its ceiling. 614/614 cached. The M42 semantic-coverage sweep on the seeded org (the 5th gate face) + the regen/seed/closure tail → tik #6 (next invocation). 5-tik cap reached — see iter-06/progress.md

**Exit gate:** a full org (e.g. 500) fills from a single supporting-population descriptor with a believable
role/avatar/skill spread (not 90% hollow), the demo-coverage SEMANTIC believability gate
(coverage-protocol.md, the M42 Playwright harness) PASSES on the generated population, hero-name collisions
stay at 0 under population-scale load, and throughput + cost stay within budget (e.g. ~1k members ≤ a few
minutes at `--max-concurrent=5`).

**Budget:** 3–5 iters. **Re-scope trigger:** if population-scale dedup/taxonomy-clipping/throttle failures
can't stabilize after ~5 tiks → user-strategic-replan.
