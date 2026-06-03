# M1 — progress (running ledger)

Iterative milestone. Exit gate: `alignctl run` ≥ 100% critical / ≥95% overall on the Clerk Alignment
DNA. Protocol: `corpus/architecture/alignment_testing.md`. One entry per closed iter.

## Running ledger
- iter-01 (tok/bootstrap): authored TOK-01 (build-order + golden-capture strategy); surfaced D1 (orgclient golden source). closed-fixed.
- iter-02 (tik): authored + validated the Clerk Alignment DNA (`clerk@2.6.0`, **22 genes / 11 caps, 13 critical**); stood up the `clerkenstein` workspace; D1 resolved (hybrid). closed-fixed. Gate NOT MET (score 0% — genome only).
- iter-03 (tik): built the **authn twin** (`colony/authn.Provider` drop-in, HS256, offline) → **VerifyToken 4/4**, score **0% → 21.1% overall / 30.8% critical**. iter-01-D1 resolved (injection = replace whole colony). closed-fixed.
- iter-04 (tik): disarmed **critical orgclient** (in-memory store, 4 methods) → **critical 30.8% → 100%**, overall → 68.4%. M1-D2 (orgclient injection finding). closed-fixed.
- iter-05 (tik): **standard orgclient** (invitations + metadata) → **overall 68.4% → 100%, critical 100% — EXIT GATE MET** (`alignctl run --gate-overall 95 --gate-critical 100` exits 0, 22/22). closed-fixed.

## Score history
| iter | overall | critical | note |
|------|---------|----------|------|
| (gate) | ≥95% | 100% | exit gate |
| iter-02 | — | — | 22-gene genome validated (no mirror yet) |
| iter-03 | 21.1% | 30.8% | authn twin: VerifyToken 4/4 |
| iter-04 | 68.4% | **100%** | critical orgclient: critical gate met |
| iter-05 | **100%** | **100%** | standard orgclient: **GATE MET** ✅ |

## Gate status: **MET** (iter-05) — overall 100% / critical 100% (22/22 genes)

## Post-gate scope (for harden + close)
- The alignment gate fired, but two milestone-scope items remain (surfaced by `/developer-kit:close-milestone`): the **injection tik** — `go.mod replace` whole-colony (authn) + the **fake-Clerk-API-server** (orgclient, M1-D2; shared with M2) + skip-worktree — and the **Delivers → `corpus/services/clerkenstein.md`** doc. Whether these are in-scope for M1's close or route to a sibling milestone is the close-milestone decision (the *gate* — alignment fidelity — is met).
