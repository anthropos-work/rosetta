---
iter: 07
milestone: M46
iteration_type: tik
status: closed
outcome: closed-diagnostic + re-scope-trigger
created: 2026-06-26
closed: 2026-06-26
---

# M46 · iter-07 — the FINAL gate face: seed the populated org + the M42 semantic-coverage sweep

**Type:** tik (#6, under TOK-01)

## Active strategy reference
**TOK-01** (build-the-three-deliverables-then-prove-on-a-real-org). Run 1's iter-06 landed the engine
fixes (multi-batch cache-index collision + the org-scale name disambiguator) and proved 4 of the 5 gate
faces empirically on a real 614-member Azure gpt-4o-mini batch. This tik closes the **5th and final gate
face** — the M42 Playwright semantic-coverage sweep on the SEEDED generated org — which is exactly TOK-01
step 5's "the real large-batch gate-proving … run the M42 manager + employee semantic-coverage sweep on
the generated org."

## Cluster / target identified
iter-06's close routed forward (verbatim): *"tik #6 (next invocation) — the gate-proof TAIL: regenerate the
now-distinct 614-member cache, seed demo-3 with `gen-batch-org-fill.seed.yaml` (`--reset --gen-batches`), run
the M42 manager + employee semantic-coverage sweep on the generated org + `datadna measure-closure`."*

**Phase 0d pre-flight finding (the load-bearing correction to the routed plan):** the gen-time disambiguator
(`genOne`, the cache-MISS path) does NOT fire on a `$0` cache-hit reseed — the `GeneratedBatchSeeder` reads
`env.Name` verbatim from the cache, and the existing complete 614-cache is only **57.7% distinct** (Zara
Al-Mansoori ×20, Jin-Soo Park ×18, …) because the disambiguator landed AFTER the cache was complete. A `$0`
reseed of THIS cache seeds DUPLICATE names → FAILS the persona/believability bar. So the right tik-6 fix is
to apply the SAME deterministic disambiguator at **SEED time** (in `GeneratedBatchSeeder`) — a `$0`,
reproducible, cost-free transform of cache → distinct rows — which makes the `$0` cache-hit reseed genuinely
distinct WITHOUT a re-generation. This is strictly better than re-generating (no real budget spent; the
existing 614-cache is reused) and it is the correct home (the seeder is "the deterministic transform of
cache → rows"; identity-distinctness is a CODE-owned property, never the AI's).

## Hypothesis
Applying the deterministic disambiguator at seed time guarantees 100% distinct generated identities from the
existing complete cache at `$0`; the seeded org then renders believably populated and PASSES the M42
semantic-coverage gate for both the manager (Dan) and employee (Maya) vantages, with closure GREEN and 0
hero-collisions at scale — closing the 5th gate face → all 5 faces pass → gate-met.

## Expected lift
The gate's primary metric `(failingSections, escapes, personaFailures, notReached, crossPortFails)` reaches
`(0,0,0,0,0)` over a frontier-exhausted crawl on the manager vantage (the headline `/enterprise/members`
org-populated surface) AND the employee vantage, on the populated org.

## Phase plan
1. **Phase 0d fix (fixtures-first):** move the deterministic disambiguation into the `seeders` package as the
   ONE source of truth (`DisambiguateGeneratedName` + the surname pool), apply it in `GeneratedBatchSeeder`
   at seed time, and have `cmd/gen-batch` call the same exported primitive (so gen-time + seed-time agree).
   Unit-prove: a cache full of duplicate names seeds 100% distinct identities, deterministically (same cache
   → same distinct names across reseeds). Full stack-seeding suite green.
2. **Bump the consumption clone** `stack-demo/rosetta-extensions` to the iter-07 tag.
3. **Seed demo-3** from the complete 614-cache via `gen-batch-org-fill.seed.yaml` (`--reset --gen-batches`,
   `--cache-root` = the authoring-copy cache) — a `$0` cache-hit reseed (no LLM). Confirm Cervato ~500 +
   Solvantis ~120 members, distinct names, avatars, role-coherent skills, closure GREEN, 0 collisions.
4. **Phase A/D — the M42 sweep:** run the `stack-verify/e2e` M42 Playwright semantic-coverage harness on the
   populated org, manager (dan-manager) + employee (maya-thriving) vantages; assert the believable-populated
   bar (`/enterprise/members` full of believable distinct members, avatars, role-coherent skills, persona
   consistency, 0 prod-eject escapes). Heartbeat every ~60–90s (it's long).
5. **Phase E — close:** if the populated org renders believably + the M42 sweep PASSES + closure GREEN + 0
   collisions + cost/throughput within budget → all 5 gate faces pass → gate-met.

## Escalation conditions
- An unstabilizable believability gap after hardening that is NOT seedable / fixable in-rext, and is
  platform-bound (a `NEXT_PUBLIC_*`-less escape or a runtime-computed surface) → re-scope-trigger (the
  coverage-protocol zero-edit line).
- The seed-time disambiguation breaks an unrelated stack-seeding test suite → user-blocker.

## Acceptable close-no-lift outcomes
If the sweep surfaces a real seed/snapshot gap whose fix is itself the deliverable (a fixtures-proven seeder
fix that lands and is re-swept), that is `closed-fixed`. A documented falsification (e.g. the existing cache
is provably the only blocker and the seed-time fix resolves it) is a complete iter.
