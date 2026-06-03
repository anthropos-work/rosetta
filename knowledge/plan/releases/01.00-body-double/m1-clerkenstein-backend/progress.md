# M1 — progress (running ledger)

Iterative milestone. Exit gate: `alignctl run` ≥ 100% critical / ≥95% overall on the Clerk Alignment
DNA. Protocol: `corpus/architecture/alignment_testing.md`. One entry per closed iter.

## Running ledger
- iter-01 (tok/bootstrap): authored TOK-01 (build-order + golden-capture strategy); surfaced D1 (orgclient golden source). closed-fixed.
- iter-02 (tik): authored + validated the Clerk Alignment DNA (`clerk@2.6.0`, **22 genes / 11 caps, 13 critical**); stood up the `clerkenstein` workspace; D1 resolved (hybrid). closed-fixed. Gate NOT MET (score 0% — genome only).
- iter-03 (tik): built the **authn twin** (`colony/authn.Provider` drop-in, HS256, offline) → **VerifyToken 4/4**, score **0% → 21.1% overall / 30.8% critical**. iter-01-D1 resolved (injection = replace whole colony). closed-fixed.

## Score history
| iter | overall | critical | note |
|------|---------|----------|------|
| (gate) | ≥95% | 100% | exit gate |
| iter-02 | — | — | 22-gene genome validated (no mirror yet) |
| iter-03 | **21.1%** | **30.8%** | authn twin: VerifyToken 4/4 |

## Next-iter queue
- **iter-04 (tik)** — the **critical orgclient** methods (CreateOrganization, CreateMembership, ChangeRole, DeleteOrganizationMembership) + an in-memory disarmed store + goldens → drive **critical → 100%** (9 critical orgclient genes remaining).
- **iter-05 (tik)** — the **standard orgclient** methods (InviteMember, BulkInviteMembers, RevokeInvitation, Update×3) → overall ≥95% → **gate**.
- then: the injection tik (`go.mod replace` whole-colony + skip-worktree) + `corpus/services/clerkenstein.md`.
