# M46 Progress

## Running ledger
Iter closeouts append here (one line each — the tik/tok, what the semantic-coverage sweep measured on the
generated org [believable spread vs hollow, gate PASS/FAIL, collision count, throughput + cost vs budget],
and what was tuned/fixed).

<!-- iter-NN/ dirs are created by /developer-kit:build-mstone-iters on first run. -->

- iter-01 (tok·bootstrap): authored TOK-01 (build 3 deliverables fixtures-first → prove on a real ~500-member org via the M42 semantic sweep); Phase 0b KB-fidelity GREEN — see iter-01/progress.md
- iter-02 (tik): deliverable #1 — auto-fill count (`Batch.Fill`/`resolveBatchCounts`: one descriptor fills a story to its Size, per-story, deterministically); 8 fixtures-first tests + full stack-seeding suite green; gate metric unmoved (proven on the real-run sweep, tik #5) — see iter-02/progress.md
- iter-03 (tik): deliverable #2 — per-story batch distribution (`BatchMember.StoryIndex` + per-member story routing in GeneratedBatchSeeder; was hardcoded stories[0]); composes with the iter-02 fill (each org fills to its OWN Size); 2 new seeder tests + full suite green — see iter-03/progress.md

**Exit gate:** a full org (e.g. 500) fills from a single supporting-population descriptor with a believable
role/avatar/skill spread (not 90% hollow), the demo-coverage SEMANTIC believability gate
(coverage-protocol.md, the M42 Playwright harness) PASSES on the generated population, hero-name collisions
stay at 0 under population-scale load, and throughput + cost stay within budget (e.g. ~1k members ≤ a few
minutes at `--max-concurrent=5`).

**Budget:** 3–5 iters. **Re-scope trigger:** if population-scale dedup/taxonomy-clipping/throttle failures
can't stabilize after ~5 tiks → user-strategic-replan.
