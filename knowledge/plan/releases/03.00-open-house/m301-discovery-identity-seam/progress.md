# M301 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in [`overview.md`](overview.md).

## Section checklist

- [ ] **(1) Resource-catalog registry** — the machine source (OpenAPI 3.1 + `x-anthropos-*` extension), the
  loader, the tests that assert every entry has the required fields.
- [ ] **(2) `Principal` DTO** — the internal type + JSON codec + tests.
- [ ] **(3) `IdentityProvider` adapter port + `ClerkIdentityProvider`** — the port interface, the Clerk
  implementation, the contract tests (given a Clerk JWT, produce a Principal).
- [ ] **(4) `ClerkGuardrails` lint rule** — the static check forbidding `clerk.*` imports in
  `internal/customerapi/` above the adapter package. Wired into CI. **Machine-enforced** P3.
- [ ] **(5) `/v1/access/whoami` smoke handler** — behind an internal-only feature flag; proves the seam
  end-to-end.
- [ ] **Docs** — `spec-notes.md` records the OpenAPI-vs-homegrown decision + rationale; the milestone
  `decisions.md` records the argon2id/bcrypt precedent question if it surfaces here; the pillar spec-progress
  updates its #2 row to `✅`.

**Status:** `planned` — not yet started. Next: `/developer-kit:build-milestone` (M301) after the v3.0 proposal
is reviewed + accepted.

