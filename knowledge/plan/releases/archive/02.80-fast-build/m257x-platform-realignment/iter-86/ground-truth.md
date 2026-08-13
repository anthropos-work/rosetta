# iter-86 ground truth — every clone with BOTH shas

Re-derived at this open, 2026-08-05. **This sheet is authoritative for refs.** A seat that grades a
claim against a ref not on this sheet must say which ref it used and why.

> **The `origin/main` column is new at iter-86, and it is the structural half of
> `CHECK-M257x-iter76-seat-ref-discipline`'s fix** (adjudicator A's recommendation, iter-84). A seat
> given only a checkout sha has no way to see that the checkout is stale — which is precisely how
> occurrences 3, 4 and 5 happened. **6 of 14 clones are behind `origin/main` right now.**

| clone | checkout | `origin/main` | behind |
|---|---|---|---|
| `platform` | `0dab54df` | `0dab54df` | 0 |
| `app` | `b948604f` | `71773747` | **60** ⚠️ |
| `next-web-app` | `bb3313bc` | `ad767d1b` | **26** ⚠️ |
| `storage` | `4ce8ece5` | `63bffc89` | **6** ⚠️ |
| `jobsimulation` | `462343b0` | `82cb66ec` | **4** ⚠️ |
| `messenger` | `fa47850d` | `a0ec933f` | **3** ⚠️ |
| `cms` | `ca50c817` | `f38c0c4a` | **2** ⚠️ |
| `sentinel` | `88bc5592` | `f2c46190` | **2** ⚠️ |
| `ant-academy` | `9c3843cd` | `9c3843cd` | 0 |
| `roadrunner` | `87d8d443` | `87d8d443` | 0 |
| `studio-desk` | `14a5442a` | `14a5442a` | 0 |
| `graphql-wundergraph` | `60c229f3` | `60c229f3` | 0 |
| `rosetta-extensions` (in `stack-demo/`) | `ab81527a` | `ab81527a` | 0 |
| `rosetta-extensions` (**authoring copy**, `.agentspace/`) | `b2019120` | `b2019120` | 0 |

## The two rules that go with the sheet

1. **Grade at the ref the claim names — UNLESS the sentence asserts currency, in which case no
   neighbouring pin rescues it.** (iter-84's amended form of §5 rule 33. Routed, deliberately *not*
   written into the frozen briefing mid-run; see `../iter-86/decisions.md` `D-M257x-86-2` for the
   disposition and its cost.)
2. **A corpus citation into `rosetta-extensions` grades against the AUTHORING copy** (`b2019120`)
   unless the citing block pins a ref. The `stack-demo/` clone is a different clone of the same repo.
