# M2b — decisions

## M2b-D2 — Directory scheme = library-named (design-time, user-chosen 2026-06-03)
The repo is reorganized into **one dir per mocked dependency, named after the dependency**:
`authn/` (colony/authn), `clerk-backend/` (clerk-sdk-go/v2), `clerk-frontend/` (@clerk/clerk-js +
@clerk/nextjs), `clerk-webhook/` (svix), plus `shared/` + `alignment/` + `knowledge/`.
**Alternatives rejected:** *surface-named* (`authn/bapi/fapi/webhook` — Clerk's own API-surface
vocabulary, less import churn) and *minimal-move* (keep all current dirs, only add `knowledge/` +
`.agentspace/`). The user chose library-named for explicitness about *which dependency* each dir mocks,
accepting the extra import/rename churn (caught by the green-gate invariant).

## M2b-D1 — Go package identifiers for hyphenated dirs
Go package names can't contain hyphens. Each hyphenated dir declares a clean package: `clerk-backend/` →
`package clerkbackend`, `clerk-frontend/` → `package clerkfrontend`, `clerk-webhook/` → `package
clerkwebhook`. Import paths keep the hyphen (`clerkenstein/clerk-backend`). **To confirm at build** —
fallback is hyphen-free dirs (`clerkbackend/`).

## M2b-D3 — `repo-consolidate` is user-invoked (process constraint, not a choice)
`/singularity-kit:repo-consolidate` is `disable-model-invocation`, so the S4 consolidation run cannot be
model-triggered. The build authors the structure **to** repo-consolidate's standard (S1–S3) so the
user's run is a clean finalize that emits `CLAUDE.md` + `singularity-manifest.md`. Recorded here so the
build doesn't mistake "couldn't auto-run the skill" for a blocker.

<!-- Build/close decisions (M2b-D4+) recorded here as they arise. -->
