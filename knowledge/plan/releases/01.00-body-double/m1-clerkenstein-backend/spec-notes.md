# M1 — spec notes

## Pre-flight audits — iter-01 (bootstrap tok)
**Phase 0b — KB-fidelity (M1): GREEN.** The milestone's contract docs all exist and are accurate:
`corpus/services/clerk-integration.md` (the centralized Clerk integration doc), `alignment_testing.md`
(the iteration protocol + the M0 framework, shipped + verified), `shared_libraries.md` § authn/colony,
`staging-clerk.md`, `webhook_setup.md`. No blind area; `corpus/services/clerkenstein.md` is the one
net-new doc this milestone Delivers (authored as the mirror takes shape, not a precondition). The real
Clerk Go surface is readable from `anthropos-dev/app` for DNA authoring.

## The platform-consumed Clerk Go surface (the DNA's scope — from M0-era research, verified vs anthropos-dev/app)
Two sides, both Go (the JS side is M2):

### A. authn provider (`colony/authn`, local JWT verification — NO live Clerk needed)
- Verifies Clerk session JWTs locally against Clerk's JWKS (`jwt.Verify` + `jwt.Decode`, 1-min leeway).
- Extracts custom claims: `eid` (internal UUID), `email`, `firstname`, `lastname`, `org.eid`, `org_id`, `org_role`.
- **Capability (critical):** token verification + claim extraction. **Golden source:** locally
  mintable — generate a session JWT with a test key + the claim shape; the "source" behavior is the
  Clerk SDK's verify/decode, runnable offline with a local JWKS. This is the easy, offline-capturable side.

### B. orgclient (`app/internal/clerk/orgclient/clerk.go`, networked Clerk API → `https://api.clerk.com/v1`)
Methods the platform calls: `InviteMember`, `CreateMembership`, `CreateOrganization`,
`DeleteOrganizationMembership`, `BulkInviteMembers`, `RevokeInvitation`, `ChangeRole`,
`UpdateMembershipMetadata`, `UpdateUserMetadata`, `UpdateClerkOrganizationWithExternalId`.
- **Capabilities (mostly critical):** org/membership/invitation CRUD + metadata writes.
- **Golden source = the LIVE Clerk API** — this is the record/replay challenge: capturing these
  goldens needs either (a) live Clerk dev-app credentials + network, or (b) hand-authored expected
  shapes derived from the `clerk-sdk-go/v2` response types. **This is the milestone's central open
  decision** (see TOK-01 + decisions.md).

## Injection seam (decided at design; proven prior art)
Build-time `go.mod replace` of the mirror in for `colony/authn` + the orgclient, made invisible
upstream via skip-worktree — the exact mechanism staging already uses for its `vendor-colony/` v2-JWT
patch. Open sub-question (resolve early): stub just `authn`+`orgclient`, or vendor all of `colony`
(authn is a package inside the colony module).

## Score mechanics (M0 framework)
The mirror ships a `--target source|mirror --dna PATH` runner emitting the outcomes protocol; the Clerk
DNA's genes use the operators from M0 (`error_class` for the API error cases, `normalized` for
responses carrying generated ids/timestamps, `exact`/`shape` otherwise). `alignctl run --gate-overall
95 --gate-critical 100` exits 0 when the gate is met.
