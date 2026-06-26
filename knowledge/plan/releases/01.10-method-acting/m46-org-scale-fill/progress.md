# M46 Progress

## Running ledger
Iter closeouts append here (one line each — the tik/tok, what the semantic-coverage sweep measured on the
generated org [believable spread vs hollow, gate PASS/FAIL, collision count, throughput + cost vs budget],
and what was tuned/fixed).

<!-- iter-NN/ dirs are created by /developer-kit:build-mstone-iters on first run. -->

- iter-01 (tok·bootstrap): authored TOK-01 (build 3 deliverables fixtures-first → prove on a real ~500-member org via the M42 semantic sweep); Phase 0b KB-fidelity GREEN — see iter-01/progress.md

**Exit gate:** a full org (e.g. 500) fills from a single supporting-population descriptor with a believable
role/avatar/skill spread (not 90% hollow), the demo-coverage SEMANTIC believability gate
(coverage-protocol.md, the M42 Playwright harness) PASSES on the generated population, hero-name collisions
stay at 0 under population-scale load, and throughput + cost stay within budget (e.g. ~1k members ≤ a few
minutes at `--max-concurrent=5`).

**Budget:** 3–5 iters. **Re-scope trigger:** if population-scale dedup/taxonomy-clipping/throttle failures
can't stabilize after ~5 tiks → user-strategic-replan.
