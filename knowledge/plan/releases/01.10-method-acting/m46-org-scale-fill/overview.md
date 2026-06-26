---
milestone: M46
slug: org-scale-fill
version: v1.10 "method acting"
milestone_shape: iterative
exit_gate: "A full org (e.g. 500) fills from a single supporting-population descriptor with a believable role/avatar/skill spread (not 90% hollow), the demo-coverage SEMANTIC believability gate (coverage-protocol.md, the M42 Playwright harness) PASSES on the generated population, hero-name collisions stay at 0 under population-scale load, and throughput + cost stay within the declared budget (e.g. ~1k members ≤ a few minutes at --max-concurrent=5)."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
status: planned
created: 2026-06-26
last_updated: 2026-06-26
complexity: medium-large
delivers: updates to corpus/ops/demo/ai-generation-spec.md + corpus/ops/cache-spec.md with the org-scale + preview workflow (the supporting-population batch, per-story distribution, the gen-batch preview/dry-run mode, throughput tuning + 429 backoff, the --gen-batches opt-in fence)
depends_on: M45 (the engine + cache + closure-safe GeneratedBatchSeeder); reuses the M42 Playwright coverage harness
spec_ref: .agentspace/scratch/roadmap-research-2026-06-26.md (the v1.10-extend research note; the org-scale strand)
---

# M46 — Org-scale fill + gen-batch preview CLI

## Goal
Scale the engine from a **bounded batch** to filling an **ENTIRE org** from one descriptor
(supporting-population auto-fill to org size) **+** add a **preview / dry-run CLI** to review a batch
before seeding. M45 proved the engine + cache on a bounded batch; M46 makes a 220/500/1k org **believable
end-to-end** and gives an author a way to inspect a batch before committing it. **Entirely in
`rosetta-extensions` — zero platform-repo edits.**

## Exit gate (observable, machine-verifiable)
- A **full org (e.g. 500)** fills from a **single supporting-population descriptor** with a believable
  role/avatar/skill spread (**not 90% hollow**);
- the **demo-coverage SEMANTIC believability gate** (`coverage-protocol.md`, the **M42 Playwright
  harness**) **PASSES** on the generated population;
- **hero-name collisions stay at 0** under population-scale load;
- **throughput + cost stay within the declared budget** (e.g. ~1k members ≤ a few minutes at
  `--max-concurrent=5`).

## Why iterative (not section)
Population-scale behaviour — dedup at scale, taxonomy-clipping under load, throttle/backoff under burst,
and the semantic-believability spread — is discovered by **measuring the generated org**, not enumerable
up front. The commitment is the **gate** (the M42 semantic-coverage sweep, here measuring the generated
population); the fix list emerges per-iter. **Budget 3–5 iters.**

## Iteration protocol
[`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md) — the
**population-believability gate = the M42 Playwright semantic-coverage sweep**, here measuring the
**generated org**. Each iter: generate + seed a population → run the semantic sweep → triage the
hollow/incoherent spread + any collision/throttle failure → tune the descriptor / engine → re-measure →
close the iter.

**Re-scope trigger:** if **population-scale dedup / taxonomy-clipping / throttle** failures can't
stabilize after **~5 tiks**, escalate to a user-strategic-replan.

## Scope
**In:**
- A **supporting-population batch** (`count: auto-fill to org size`, `roles_mix`, `verified_range`,
  `trajectory_mix`) expanding to fill the **remaining N members** of a story so a 220/500/1k org is
  believable end-to-end.
- **Per-story batch distribution** (story-local, the multi-org Stories model).
- A **`gen-batch` PREVIEW / DRY-RUN mode** — render the expanded per-member prompts + cached generated
  JSON to stdout/file **WITHOUT seeding**, with an **estimated-cost line**, so an author reviews a batch
  before committing it.
- **Throughput tuning** for large pops + **429 backoff verification** under burst.
- An optional **`--gen-batches` opt-in flag** on `stackseed` (fence against silent OpenAI-unreachable
  failures).
- Update `ai-generation-spec.md` + `cache-spec.md` with the **org-scale + preview workflow**.

**Out:**
- The **engine + cache primitives** themselves (M45).
- **Deep timelines** for every generated member at 1k scale (cost-prohibitive — shallow by default unless
  explicitly declared).
- A **GUI / web preview** — the CLI dry-run is the preview surface.
- Any **platform-repo edit**.

## Depends on / Parallel with
- **Depends on:** **M45** (the engine + cache + the closure-safe `GeneratedBatchSeeder`). **Reuses** the
  M42 Playwright coverage harness.
- **Parallel with:** none.

## Open questions
- **Supporting-pop timeline depth** (default shallow: name + skills + bio).
- **Dedup at scale** — pre-gen reserved-names vs post-gen re-roll.
- A **taxonomy-coverage floor** per role before large-batch gen.
- The **curated-vs-batch mix per org** (the product call — default ~3 curated heroes + batch-fill the
  rest).
- Whether the **preview surfaces an estimated cost** before a real run (default: yes).

## KB dependencies
M46 reads these corpus docs as contract (it must not contradict them; it extends them):
- `corpus/ops/demo/ai-generation-spec.md` — the engine + cache + the gen-acceptance protocol (M45 authors
  it); M46 adds the org-scale + preview workflow.
- `corpus/ops/cache-spec.md` — the prompt-hash cache (M45); M46 extends it for population-scale runs.
- `corpus/ops/demo/coverage-protocol.md` — the M42 Playwright semantic-coverage harness + the iteration
  protocol (the population-believability gate reuses it).
- `corpus/ops/demo/stories-spec.md` — the multi-org Stories model the per-story batch distribution fills.
- `corpus/ops/seeding-spec.md` — the seeding blueprint + the production-isolation boundary the supporting
  population is written within.

## Delivers →
- `corpus/ops/demo/ai-generation-spec.md` — the **org-scale + preview workflow**: the supporting-population
  batch (`count`/`roles_mix`/`verified_range`/`trajectory_mix`), per-story distribution, the `gen-batch`
  preview/dry-run mode (expanded prompts + cached JSON + estimated cost, no seeding), throughput tuning +
  429 backoff, and the `--gen-batches` opt-in fence.
- `corpus/ops/cache-spec.md` — the cache behaviour at population scale (the preview reading cached JSON,
  cache reuse across the org-fill).
