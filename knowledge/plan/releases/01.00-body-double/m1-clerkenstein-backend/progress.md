# M1 — progress (running ledger)

Iterative milestone. Exit gate: `alignctl run` ≥ 100% critical / ≥95% overall on the Clerk Alignment
DNA. Protocol: `corpus/architecture/alignment_testing.md`. One entry per closed iter.

## Running ledger
- iter-01 (tok/bootstrap): authored TOK-01 (build-order + golden-capture strategy); surfaced D1 (orgclient golden source). closed-fixed.
- iter-02 (tik): authored + validated the Clerk Alignment DNA (`clerk@2.6.0`, **22 genes / 11 caps, 13 critical**); stood up the `clerkenstein` workspace; D1 resolved (hybrid). closed-fixed. Gate NOT MET (score 0% — genome only).

## Score history
| iter | alignment score | note |
|------|-----------------|------|
| (gate) | 100% critical / ≥95% overall | exit gate |
| iter-02 | — (DNA authored; no mirror yet) | 22-gene genome validated |

## Next-iter queue
- **iter-03 (tik)** — build the authn-provider twin (satisfy the cached `colony/authn` Provider interface; local JWT mint/verify; one universal credential) + its `--target source|mirror` runner + locally-captured authn goldens → first real `alignctl run` on the 6 critical authn genes. Resolve iter-01-D1 (colony replace granularity) here. **Feasible:** colony v0.34.1 + clerk-sdk-go confirmed in the module cache.
