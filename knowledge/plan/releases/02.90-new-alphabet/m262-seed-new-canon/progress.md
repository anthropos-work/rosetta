# M262 — Progress

**Status: COMPLETE** (structural). 2026-08-15.

- [x] remap seeded refs through the redirect map — **resolved by NAME against the canon's own mapping**
- [x] re-resolve the 8 literal job-role names
- [x] add a per-hero richness floor to the closure gene
- [x] price + re-run the AI profile regeneration — **DONE 2026-08-15, $0.2196**

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

## The AI half — done, and the key was not where I looked

**Both halves landed.**

**Embeddings.** `taxonomy-load` re-run with a key: **4,268 vectors** (3,562 skills + 706 roles), all
non-null, captured and replayed into `demo-5`. The canon is no longer inert for AI matching — that was
M265's clause-4 blocker. Cost: about a cent.

**Profiles.** `gen-batch` regenerated all **364** members against the new canon:

```
calls=934  cache-hits=0  tokens=959,766
cost=$0.2196  ceiling=$5.0000  (4.4%)
valid-JSON rate (pre-re-roll): 100.0%
```

Priced at $0.15–0.25 before spending, per `D-v29-3`; actual **$0.2196**. The cache showed `0/364
already cached`, which is the capture-version key working exactly as designed — a new taxonomy
invalidates every generated member.

### ⚠️ The key was in `studio-desk/.env`, not `app/.env` — and I stopped one file too early

I reported this blocked after testing only `app/.env`, where `OPENAI_KEY` returns
`429 insufficient_quota` and the Azure endpoint has **no chat deployment** (`404
DeploymentNotFound` — it has an *embeddings* deployment, which is why the same key computed the
vectors and failed here). Enumerating **every** secret file found four more AI key sets, and
`studio-desk/.env`'s **`AI_OPENAI_API_KEY` answers `gpt-4o-mini` with HTTP 200**.

**The lesson is the search, not the key:** "the AI key" was treated as one thing when the workspace
holds several, under different names, with different entitlements. `app/.env` and `studio-desk/.env`
both look authoritative; only one works for chat.


_(Resolved — see the section above. The routing to M265 is withdrawn: the dependency is closed.)_
