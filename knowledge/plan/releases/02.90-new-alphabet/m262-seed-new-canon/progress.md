# M262 — Progress

**Status: COMPLETE** (structural). 2026-08-15.

- [x] remap seeded refs through the redirect map — **resolved by NAME against the canon's own mapping**
- [x] re-resolve the 8 literal job-role names
- [x] add a per-hero richness floor to the closure gene
- [ ] price + re-run the AI profile regeneration — **BLOCKED on an AI key** (see below)

## The 8 pinned role names: 6 survived, 2 re-pointed by the canon's own mapping

Measured against the loaded canon:

| preset role | canon |
|---|---|
| Account Executive · Backend Developer · Data Analyst · DevOps Engineer · Engineering Manager · Sales Manager | ✅ survive unchanged |
| Business Operations Analyst | → **Business Analyst** (`rematch`, `review=true`) |
| Talent Acquisition Specialist | → **Recruiter** (`rematch`, `review=false` — trusted) |

Both successors come from `role_redirects.csv` — **the canon's own decision, not a guess of ours**.
All 8 now resolve.

> **The redirect TABLE could not have done this.** `job_role_redirects` keys by `old_node_id`; the
> presets pin **names**. Only the bundle CSV carries `old_title`, so a name-pinned preset is resolved
> by consulting the CSV and re-pointing the preset — not by a runtime lookup. That is why this half of
> D-M259-2 landed as a data edit rather than as a resolver.

## The richness floor — and the case that proves it was needed

The closure gene measured whether the refs that EXIST resolve. **It was structurally blind to whether
they exist at all**: when `TaxonomyRefs` returns an empty pool, `PersonaSeeder` skips enrichment
rather than fabricating — right, and silent. Nothing dangles, because nothing was written.

**`demo-5` is a live instance.** Its taxonomy replay failed (M263), the seed ran against an empty
catalogue, and **591 of 591 memberships ended with ZERO skills** — while the gene's own `referenced`
count sat at **59**, carried by evidences and validation results naming skills the personas never had.
Dangling-only called that a closed seed.

The gene now also measures **coverage**: how many seeded memberships hold at least one skill ref, as a
**fraction** (seeds differ in size; a fixed count would miss a small world or fail a large one), with
a deliberately loose floor — the two states it separates are *populated* and *empty*, and nothing
legitimate sits between.

Five controls, including the demo-5 shape that used to pass, **partial degradation (3 of 10 — the case
this milestone was scoped for)**, and two that keep the diagnostics from colliding: an unseeded stack
must still fail the **population witness**, and when refs dangle *and* heroes are empty the dangling
message wins because it names a concrete broken id.

## ⚠️ What is NOT done, and why

**The AI profile regeneration did not run.** `gen-batch` needs an AI key and none is configured on
this machine; the canon itself also loaded **without embeddings** (`vectors not computed: the canon is
loaded but does not take part in matching until this is re-run`). So:

- generated member profiles still reference the OLD taxonomy's skill names;
- **AI skill-matching against the canon does not work at all** until the load is re-run with an
  embedding manager.

Neither is a code gap — both need a key. `D-v29-3` set a $200 ceiling and said price-before-spend;
there is nothing to price until a key exists. **Routed to M265**, which cannot pass its gate without
it: a stack whose heroes have no verified-skill chain is not a proven stack.
