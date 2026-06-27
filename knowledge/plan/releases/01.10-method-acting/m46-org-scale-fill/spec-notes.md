# M46 Spec Notes

> **Status: `archived` — M46 CLOSED 2026-06-27 (`closed-on-gate`).** The build-time TODO markers below were all
> DELIVERED (iter-02..07); the load-bearing truths live in the corpus docs (`ai-generation-spec.md` §4g,
> `cache-spec.md`, `coverage-protocol.md`, `snapshot-spec.md`). Kept as an intra-build record.

Technical notes accumulate here during build. The authoritative design lives in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md)
(the org-scale strand). M46 is **entirely in `rosetta-extensions` — zero platform-repo edits**. The
iteration protocol REUSES the M42 Playwright semantic-coverage harness
([`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md)) measuring the
GENERATED org.

## Pre-flight audits — iter-01
- **KB-fidelity (Phase 0b, 2026-06-26):** **GREEN**. Report:
  [`kb-fidelity-audit.md`](kb-fidelity-audit.md). 9/9 load-bearing claims across
  `ai-generation-spec.md` + `cache-spec.md` ALIGNED; every M46-scope topic PAIRED (no blind areas). The
  three M46 deliverable gaps confirmed present in code as the correct starting point: (a) `Batch.Count` is
  a fixed int — no org-size auto-fill (`batch.go:32`); (b) `--dry-run` is a count-only summary, no
  per-member render / estimated-cost line (`cmd/gen-batch/main.go:110`); (c) `GeneratedBatchSeeder`
  hardcodes `stories[0]` — no per-story distribution (`seeders/generated_batch.go:118`). Topic→doc→code
  triples are in the audit's Topic Inventory.
- **Path nit (carry):** `overview.md` `delivers`/KB-deps frontmatter cites `corpus/ops/cache-spec.md`; the
  file is at `corpus/ops/demo/cache-spec.md`. Correct during the doc-update phase.

## Supporting-population batch
TODO: `count: auto-fill to org size`, `roles_mix`, `verified_range`, `trajectory_mix`. Expands to fill the
remaining N members of a story so a 220/500/1k org is believable end-to-end (NOT 90% hollow). Per-story
distribution (story-local, the multi-org Stories model).

## `gen-batch` PREVIEW / DRY-RUN mode
TODO: render the expanded per-member prompts + cached generated JSON to stdout/file WITHOUT seeding, with
an estimated-cost line, so an author reviews a batch before committing it. The CLI dry-run IS the preview
surface (no GUI/web preview).

## Throughput tuning + 429 backoff
TODO: tune throughput for large pops; verify 429 backoff under burst. Budget target: ~1k members ≤ a few
minutes at `--max-concurrent=5`.

## `--gen-batches` opt-in fence on `stackseed`
TODO: an optional opt-in flag fencing against silent OpenAI-unreachable failures (a real LLM call is
gated behind the flag, so an absent/unreachable key fails loud, not silent).

## Dedup at population scale
TODO: hero-name collisions stay at 0 under population-scale load. Open: pre-gen reserved-names vs post-gen
re-roll at scale; a taxonomy-coverage floor per role before large-batch gen.

## Curated-vs-batch mix
TODO (product call): default ~3 curated heroes + batch-fill the rest, per org.

## Population-believability gate (the M42 harness, on the generated org)
TODO: run the semantic-coverage sweep (coverage-protocol.md) against the generated population — real
content + substantial per-section cardinality + persona self-consistency + 0 prod-eject escapes — must
PASS on the generated org (not 90% hollow).

## Delivers — updates to ai-generation-spec.md + cache-spec.md
TODO: the org-scale + preview workflow (supporting-population batch, per-story distribution, the dry-run
preview + estimated cost, throughput tuning + 429 backoff, the `--gen-batches` fence); the cache behaviour
at population scale.
