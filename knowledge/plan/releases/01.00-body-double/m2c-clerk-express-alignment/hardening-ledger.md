# Hardening Ledger — M2c @clerk/express backend session verification

## Pass 1 — 2026-06-03 — final

**Iters hardened this pass:** all milestone-touched code (iter-02…05, cumulative)
**Tiks covered since prior pass:** all iters in milestone (first harden pass)
**Coverage delta on touched files:**
- `clerk-backend`: 83.8% → **97.2%** stmts (the new `getOrganization` / `getOrganizationMembershipList` handlers + `Store.GetOrganization`/`ListMemberships` — were 0%, integration-only)
- `shared`: 91.8% → **95.9%** stmts (`mustParseRSA` panic paths + `FuzzMintRS256`)
- `alignment/cmd/expressrun`: 44.8% → 49.1% stmts (`identityVal`/`str`; the mirror path is gate-covered — see note)
**Tests added:**
- `clerk-backend/reads_test.go`: 6 unit (2 store + 4 handler, incl. the 404 error path + the empty-org edge)
- `shared/rsa_test.go`: 1 unit (`mustParseRSA` rejects bad PEM/DER) + 1 fuzz (`FuzzMintRS256` — 17.9k execs, 0 crashers, the mint-then-verify invariant holds for any claim set)
- `alignment/cmd/expressrun/main_test.go`: 1 unit (`identityVal` extraction + the `str` helper, partial-claims edge)
**Bugs surfaced + fixed inline:** none (M2c's crux verified first try; no build-phase bugs to regress)
**Flakes stabilized:** none
**Knowledge backfill:** the `expressrun` **mirror path** (`emitMirror`/`verifyViaNode`) requires Node + `@clerk/express` (`EXPRESSRUN_NODE_PATH`) — it is **integration-covered by the express alignment gate** (stable 3/3), not a unit-test gap. This is why the package's *statement* coverage (49%) understates its real coverage. (Same pattern as `clerkrun`/`jsfapirun` being integration-covered by their gates.)
**Stop condition:** continue-to-next-pass — confirm no meaningful residual after the big Pass-1 deltas

## Pass 2 — 2026-06-03 — final

**Iters hardened this pass:** all milestone-touched code (cumulative re-scan)
**Tiks covered since prior pass:** 0 (confirmation pass)
**Coverage delta on touched files:** 0 (no shallow tests added — the cap-discipline guard)
**Tests added:** none — the 6-dimension scan found no meaningful gap. The remaining uncovered is:
  (1) the `expressrun` mirror path (Node-dependent, integration-covered by the express gate, stable 3/3);
  (2) defensive branches — `mustParseRSA`'s non-RSA-key path (needs an EC PKCS#8 fixture; a guard, not a real path) and `MintRS256`'s `SignPKCS1v15` error (unreachable with a valid key). Writing tests for these would be shallow coverage-box-ticking, which the harden discipline rejects.
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none — flake gate 3/3 clean on the new tests
**Stop condition:** stabilized — coverage delta < 2% AND the dimension scan found nothing meaningful new

## Post-harden gate verification (all green, additive)
Go 22/22 · JS 9/9 · **Express 9/9** · drift ALL PASS · `-race` suite 8/8 · gofmt/vet clean.
