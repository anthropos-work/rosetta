# Clerkenstein

**Status:** v0.1 (v1.0 "body double" · milestone M1) · **Last updated:** 2026-06-03
**Repo:** `anthropos-demo/clerkenstein` (gitignored demo scratchpad, its own git) · **Measured by:** the [alignment framework](../architecture/alignment_testing.md)

## Role

Clerkenstein is a **drop-in mock of the Clerk library** the platform uses — the *same interface*, with
all security and sync **disarmed**. It exists so **demo** environments can create users / orgs / admins
and log in/out with no Clerk friction (one universal credential, no live API, no webhooks, no rate
limits), while platform repos keep "thinking" they use Clerk with **zero source changes**.

It is the **first mirror produced by the M0 alignment process** (not a hand-built mock): its fidelity
is *measured* as a 0–100% alignment score against a Clerk **Alignment DNA**, and the M1 milestone drove
that score to its gate — **100% critical / 100% overall (22/22 genes)** against `clerk-sdk-go/v2 @
v2.6.0`. The DNA + mirror + goldens + runner live in the clerkenstein repo; the *measuring machinery*
([`test/alignment/`](../../test/alignment/) + `/align-dna` + `/align-run`) lives in rosetta.

## Architecture & code map

The mirror covers the two sides of the platform's **Go** Clerk consumption (the JS side is M2):

### `authn/` — the `colony/authn` provider twin (offline)
Implements the real `colony/authn.Provider` interface (`GetUser(token)`, `GetUserByID(uuid)`,
`Name()`) — a **compile-time `var _ authn.Provider` assertion** guarantees the drop-in. Tokens are
**HS256-signed with one universal key** (`authn/jwt.go`); `GetUser` verifies + extracts the platform
claim set (`eid`→`ID`, `email`, `firstname`, `lastname`, `org.eid`→org `ID`, `org_id`→org `AuthID`,
`org_role`→`AuthRole`) into a `clerkUser` implementing `authn.{User,Organization}` (`authn/user.go`,
`authn/provider.go`). No live Clerk — JWT verify is local.

### `orgclient/` — the Clerk org/membership API twin (disarmed, in-memory)
A small in-memory store (`orgclient/store.go`, `orgclient/invitations.go`) reproducing the
success/error semantics of the 10 consumed methods (CreateOrganization, CreateMembership, ChangeRole,
DeleteOrganizationMembership, InviteMember, BulkInviteMembers, RevokeInvitation, the 3 metadata
writes). No network calls.

### `cmd/clerkrun/` — the alignment runner
`clerkrun --target {source|mirror} --dna PATH` emits the outcomes protocol `alignctl` consumes — the
glue that lets the framework score the mirror. Exercised end-to-end by every `alignctl run`.

## Injection (zero platform-code changes) — two different mechanisms

| Side | Mechanism | Status |
|---|---|---|
| **authn** | `go.mod replace` the whole `colony` module with a Clerkenstein-patched colony (its `authn/provider/clerk` = the disarmed twin), made invisible upstream via **skip-worktree** — the exact pattern staging already uses for its `vendor-colony/` v2-JWT patch. | recipe documented; live wiring is demo-stack work (v1.1) |
| **orgclient** | **Different (M1-D2):** the orgclient is `app`-internal (`app/internal/clerk/orgclient`, not a published module) and calls `api.clerk.com` over HTTP — so it *can't* be `go.mod replace`d. Disarming it needs a **fake-Clerk-API-server** (DNS / base-URL redirect of `api.clerk.com` → a local mock). | **routed to M2** — the same HTTP-interception mechanism M2's JS side needs (a shared component) |

The **alignment gate measures behavior**, which the in-memory twin provides regardless of how injection
eventually wires in — so M1's gate fired without the injection being live.

## Disarmed-security properties (by design — speed + accessibility, not security)

These are deliberate, not bugs (a demo mock, never production):
- **One universal credential** — every token is HS256-signed/verified with a single fixed key.
- **JWT `alg` is not validated** — `parse` verifies the HMAC regardless of the header `alg` (real Clerk
  would reject a mismatch). Acceptable: the mock's job is to *accept* easily.
- **Tokens without `exp` never expire.**
- **The orgclient store is not thread-safe** (plain maps) — irrelevant to the alignment (the runner uses
  a fresh store per gene, single-threaded), but a consideration for the *injection* work if a single
  instance serves concurrent demo requests.

## Local development

```sh
# in anthropos-demo/clerkenstein (builds offline against the cached colony):
GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off go test ./...          # unit tests (authn + orgclient: 100%)
GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off go build -o clerkrun ./cmd/clerkrun

# measure fidelity from rosetta (the alignment gate):
cd test/alignment && go run ./cmd/alignctl run \
  --dna <…>/clerkenstein/dna/clerk-2.6.0.json \
  --runner <…>/clerkenstein/clerkrun \
  --golden-dir <…>/clerkenstein/golden --gate-overall 95 --gate-critical 100
```

## Testing

- **Unit:** `authn` + `orgclient` at **100%** (mint/verify, every error class, tampered tokens, fuzz);
  `cmd/clerkrun` is integration-covered by the alignment run.
- **Alignment:** the gate — `alignctl run` reports 100%/100% over the 22-gene `clerk@2.6.0` DNA. This is
  the milestone's exit criterion and the regression signal **M1b** will CI-gate across Clerk version
  bumps (re-`/align-dna` the new version, re-`/align-run`).

## Drift detection (M1b)

Clerk moves; the mirror must stay aligned. M1b makes a Clerk bump a **flagged, mechanical event** by
reusing M0 wholesale — no new measurement machinery, just two scripts + a CI gate (in the clerkenstein
repo):

- **`scripts/gate.sh`** — the alignment gate: builds the runner + a fresh `alignctl`, runs
  `alignctl run --gate-overall 95 --gate-critical 100`. **Exit 0** = gate met, **2** = the mirror
  regressed. (Uses a built `alignctl` binary, not `go run`, so the exact exit code propagates.)
- **`scripts/drift-check.sh --new <bumped-DNA>`** — wraps the gate with a DNA-diff step. **Exit codes:**
  **0** no drift & gate met · **1** the DNA moved (`alignctl dna diff` shows added/removed/changed
  genes — the Clerk surface changed) · **2** the gate regressed (genes broke).
- **`.github/workflows/alignment.yml`** — runs the gate on push + a **weekly** schedule (the brief's
  "follow platform updates within minutes" cadence), turning a Clerk break into a red build.

**The bump runbook** (when `clerk-sdk-go` / `@clerk/*` updates):
1. `/align-dna` the new version → a new DNA. `drift-check.sh --new …` reports what moved (exit 1).
2. Re-author the changed/added genes in the DNA; **re-capture goldens** for the moved genes
   (`alignctl capture`, or hand-author per the hybrid M1-D1 path).
3. `/align-run` (or `gate.sh`) re-scores the mirror against the bumped surface; close any new
   divergences in the twin until the gate is green again.
4. Re-pin the DNA version. The weekly CI keeps it honest between bumps.

`ALIGN_DIR` (default `../../test/alignment`) locates rosetta's `alignctl`. Verified across all exit
paths against a simulated `clerk@2.7.0` bump.

## See also
- [Alignment Testing](../architecture/alignment_testing.md) — the framework that measures this mirror.
- [Clerk integration](clerk-integration.md) — the real Clerk surface Clerkenstein mirrors.
