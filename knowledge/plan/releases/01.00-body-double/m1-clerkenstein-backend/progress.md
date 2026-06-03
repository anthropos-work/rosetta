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

## M1: Final Review (close-milestone, 2026-06-03)

Iterative, **closed-on-gate** (100%/100%, exit 0). Phase 1b deferral re-audit GREEN. In-thread review of the ~250-line mirror (100% covered, adversarially hardened). 0 code must-fixes.

### Scope (the 2 post-gate In-list items → three-fate)
- [x] **`corpus/services/clerkenstein.md`** (the milestone's `Delivers →`) → **Fate 1 DONE**: mirror design + the alignment gate + the injection recipe + disarmed-security properties (M1-D2 blended in).
- [x] **orgclient injection** (fake-Clerk-API-server, M1-D2) → **Fate 3 DONE**: routed to M2 (roadmap.md M2 In-list updated — the fake-API-server is shared with M2's JS side; the authn `go.mod replace` recipe lands in clerkenstein.md).

### Code quality
- 0 must-fix. authn + orgclient independent, no duplication, vet+gofmt clean, 100% covered.

### Adversarial review (Phase 2c — scenarios, recorded in decisions.md)
- **JWT alg not validated** (parse ignores the header `alg`, verifies HMAC regardless) → accepted: security is *disarmed by design*; document as a known property.
- **orgclient Store not thread-safe** (plain maps) → not a gate concern (runner uses a fresh store per gene, single-threaded); relevant only at *injection* time → noted for the injection work (M2-shared).
- **token with no `exp` never expires** → accepted (disarmed; the runner always sets exp for valid tokens).
- No must-fix bug; all are deliberate disarmed-mirror properties → documented in clerkenstein.md.

### Tests & benchmarks
- 100% on authn + orgclient; runner integration-covered by `alignctl run`; 1 fuzz; flake gate clean. No gaps. No perf SLAs (N/A). README quotes no test counts (no drift).

### Decision triage
- **M1-D2** (orgclient injection = fake-API-server, shared w/ M2) → blend into `corpus/services/clerkenstein.md`.
- TOK-01, D1, iter-01-D1, the adversarial scenarios → archive in `decisions.md` (maintainer-only).
