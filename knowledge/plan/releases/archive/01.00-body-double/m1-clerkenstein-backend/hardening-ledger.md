# Hardening Ledger — M1 Clerkenstein backend mirror

## Pass 1+2 — 2026-06-03 — final

**Scope:** cumulative milestone-touched code. M1's production code lives in the gitignored
`anthropos-demo/clerkenstein` repo (rosetta's `m1/...` branch holds only iter docs + `.gitignore`), so
the harden target is the mirror's Go logic: `authn/` (JWT mint/verify + claim extraction +
`colony/authn` impl) and `orgclient/` (the disarmed in-memory store). The mirror had **zero unit
tests** — the alignment score validated behavior end-to-end, but the internals (tampered tokens, wrong
formats, every error branch) were unhardened.

**Iters covered:** all milestone-touched code (final mode) — iter-03 (authn) + iter-04/05 (orgclient).

**Coverage delta on touched files:**
- `clerkenstein/authn`: 0.0% → **100.0%** stmts
- `clerkenstein/orgclient`: 0.0% → **100.0%** stmts
- `clerkenstein/cmd/clerkrun`: 0.0% (unit) — exercised end-to-end by the `alignctl run` integration
  (out-of-process, not unit-instrumented; same pattern as M0's `cmd/alignctl`). The runner is CLI glue
  dispatching to the now-100%-covered packages; dedicated unit tests waived as low-value.

**Tests added:**
- `authn/authn_test.go`: 14 — mint/parse round-trip, expired, malformed (6 forms: garbage / 2-part /
  4-part / empty / "..." / "a..c"), tampered signature, tampered body, bad-base64 body, non-JSON body,
  claim extraction (with-org + no-org), `GetUserByID`, `Emails` empty, `parseUUID` invalid, provider
  name, CreatedAt-zero.
- `authn/fuzz_test.go`: 1 fuzz (`FuzzParse` — attacker-controlled token → no panic).
- `orgclient/store_test.go`: 12 — every method × every error class (missing-name, unknown-org,
  already-member, invalid-role, not-a-member, duplicate, already-revoked), bulk all/empty, metadata
  writes, fresh-seed isolation, new-org member-map init.
- Pass 2 closed the only two residual branches (`parse` base64/JSON-decode error paths;
  `CreateMembership` members-map init) → 96.7% → 100%.

**Bugs surfaced + fixed inline:** none — the mirror was correct (the alignment gate already proved
behavior; the unit tests confirm the internals + error paths). 2 test-authoring fixes (Go composite-
literal-in-`if` parse + an inline `&clerkUser{}`), no production change.

**Flakes stabilized:** none — flake gate 3/3 consecutive clean (`go test` authn + orgclient).

**Knowledge backfill:** none KB-worthy — the harden confirmed documented behavior (the JWT claim
contract + the operators in `corpus/architecture/alignment_testing.md`); nothing new to propagate. The
iteration protocol doc (`alignment_testing.md`) needs no update.

**Stop condition:** `stabilized` — both logic packages at 100%; the 6-dimension scan found nothing
further (the runner's surface is integration-covered); 0 flakes. Verification: full suite green, gate
re-checked 100%/100% (exit 0), vet+gofmt clean. Cleanup: 0 orphan processes, 0 temp clutter.
