# M46 Decisions

Implementation decisions with rationale (recorded during build/iters). Design-time decisions live in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| _(none yet)_ | | | |

---

## TOK-01: build-the-three-deliverables-then-prove-on-a-real-org — 2026-06-26

**Tok type:** bootstrap (iter-01)

**Initial strategy:** M46 is not a "tune a model until it converges" milestone (M45 already proved
gpt-4o-mini emits valid, resolving envelopes). It is a **build-the-three-deliverables-then-empirically-
prove-on-a-real-org** milestone. The M45 seam (CODE owns structure/identity/closure, AI owns content;
non-resolving names drop) is N-invariant, so scaling is a CODE + WORKFLOW problem, not a model-quality
problem. The strategy: land the three In-scope code deliverables **fixtures-first** (so each is unit-proven
without a key or cost), THEN fire ONE real large supporting-population batch via Azure gpt-4o-mini
(values-blind, `--max-cost`-capped) and run the M42 semantic-coverage sweep on the generated org to prove
the gate. Order the tiks by dependency + de-risking:

1. **Auto-fill count (`count: auto-fill`/`fill`).** Teach `Batch` to expand to fill the REMAINING N
   members of its story relative to `Size` (Size − curated heroes − explicit-count batches). Keep the
   prompt expansion pure + deterministic (the cache invariant). Unit-prove the fill math. This is the
   spine — everything else fills a now-org-sized batch.
2. **Per-story batch distribution.** Generalize `GeneratedBatchSeeder` from hardcoded `stories[0]` to
   route each batch's generated members into ITS story's org (the multi-org Stories model). Unit-prove
   the per-story routing (org id / prefix / domain per story).
3. **Preview / dry-run mode** (`gen-batch --preview`/extended `--dry-run`). Render the expanded
   per-member prompts + (cached) generated JSON to stdout/file WITHOUT seeding, with an
   **estimated-cost line** (reuse the existing `estimateUsage` + cost table). Unit-prove the rendered
   plan + the cost estimate.
4. **The `--gen-batches` opt-in fence on `stackseed`** + **throughput/429 verification.** Fence the real
   LLM dependency behind an explicit opt-in (absent/unreachable key fails LOUD, never silent). Verify the
   ai-lib retry/backoff holds under burst at `--max-concurrent=5`. Mostly code + a real-run observation.
5. **The real large-batch gate-proving.** Fire a real ~500-member supporting-population batch via Azure
   gpt-4o-mini (values-blind, `--max-cost ~$0.15`, `--max-concurrent 5`, frequent journal heartbeats),
   seed demo-3 (or a fresh stack), run the M42 manager + employee semantic-coverage sweep on the
   generated org. Gate = the org renders believably populated (the `/enterprise/members` table + org-scale
   surfaces full of believable generated members) + 0 hero-collisions at scale + closure GREEN + cost +
   throughput within budget.

**Fixtures-first discipline (carried from M45 §4c):** build + unit-test each deliverable against fixtures
(no key, no cost) FIRST; the real LLM batch runs LAST, only to prove the empirical believability +
throughput the fixtures can't. Each code deliverable closes `closed-fixed` on its unit proof; the gate
closes on the real-org sweep.

**Rationale:** the gate has five faces (believable spread, semantic-gate PASS, 0 collisions, closure
GREEN, budget) and four are already structurally guaranteed by the M45 seam at any N — the one genuinely
empirical face is **"does a full org rendered from one descriptor LOOK believable, not 90% hollow, under
the M42 semantic gate?"** Building the three deliverables fixtures-first front-loads the deterministic work
(cheap, fast, no key) and isolates the single real-run gate-proof to the end, where a heartbeat-streamed
capped Azure batch answers the empirical question once. This is the opening move because the milestone's
risk is concentrated in the real-org sweep, and we want maximal de-risking (unit-proven deliverables, a
warm demo, a calibrated harness) before we spend the (negligible but capped) real-run budget.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** the gate (M42 semantic sweep on a generated org + 0 collisions + closure
GREEN + budget) has NOT been measured on a generated org yet — M45 proved the engine on a bounded N=20,
not a full org. demo-3 is up (offset +30000); the M42 harness is calibrated for employee (Maya) + manager
(Dan) vantages; the manager `/enterprise/members` page is the headline org-populated surface the gate
reads. Zero of the three code deliverables are built. Budget 3–5 iters.

**Next-tik direction:** iter-02 (the first tik) implements **deliverable #1 — auto-fill count** in
`blueprint/batch.go` (a `Count` that resolves to "fill the remaining N of the story" when set to an
auto-fill sentinel), fixtures-first: unit-prove the fill math (Size − curated − explicit counts), keep the
prompt expansion pure/deterministic so the cache invariant holds, and confirm `EffectiveBatches()` still
expands to the right per-member count. No real LLM call in iter-02 (the fill math is offline).

