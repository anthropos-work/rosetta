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

## M2b-D4 — `parse` exported as `shared.Parse` (build, S1)
Extracting `authn/jwt.go` → `shared/jwt.go` split the mint/verify pair across packages: `clerk-frontend`
(now via `shared`) MINTS and `authn` VERIFIES. The verify entry point `parse()` was unexported (package-
local to `authn`); after the move `authn/provider.go` calls it from a *different* package, so it had to
be **exported as `shared.Parse`**. `Mint`/`Claims` were already exported. The unexported helpers
(`b64`, `sign`, `universalSecret`, `errMalformed/errSignature/errExpired`) stay unexported in `shared` —
the JWT-internal tests that reference them moved into `shared/` (co-located), so package-private access
holds. The runner consumes error *strings* via `err.Error()` (`"malformed"/"bad-signature"/"expired"`),
not the sentinel vars, so the gate is unaffected by the var visibility.

## M2b-D5 — script `base` = `alignment/`, ALIGN_DIR depth +1 (build, S1, confirms spec-notes)
The scripts moved `scripts/` → `alignment/scripts/`, one level deeper. Each script resolves
`base="$(cd "$(dirname "$0")/.." && pwd)"` which now points at **`alignment/`** (not the repo root).
Consequence: the runner/DNA/golden defaults (`./cmd/clerkrun`, `dna/…`, `golden`) stay UNCHANGED because
they're relative to `alignment/` — but `ALIGN_DIR` (rosetta's `test/alignment`, reached by walking up out
of the repo) gains one `../`: `../../test/alignment` → **`../../../test/alignment`**. Verified by running
all three scripts from the new location (gate 22/22 + 9/9, drift-test 9/9). The CI YAML sets `ALIGN_DIR`
to an absolute path, so only its script *paths* needed the `alignment/` prefix.

## M2b-D6 — built binaries move to `alignment/`; `.gitignore` re-anchored (build, S1)
The repointed `gate.sh` builds `clerkrun`/`jsfapirun` into `alignment/` (next to `cmd/`), not the repo
root. The old `.gitignore` anchored only `/clerkrun` `/jsfapirun`, so the new `alignment/clerkrun` etc.
would have been committed as multi-MB binaries. Added `/alignment/clerkrun` + `/alignment/jsfapirun`
anchors (kept the root anchors for a manual `go build -o clerkrun …` from root) and fixed the stale
`.agentspace` comment. Confirmed all four paths are `git check-ignore`d; the `cmd/` source stays tracked.

## M2b-D7 — CI gains a JS/FAPI gate step (build, S1)
While repointing `.github/workflows/alignment.yml` to `alignment/scripts/`, added an explicit
"Alignment gate (JS/FAPI surface)" step (env: `RUNNER_PKG=./cmd/jsfapirun … DNA=dna/clerk-js-5.json
GOLDEN_DIR=golden-js`). M2 parameterized `gate.sh` for both surfaces and both DNAs now exist, but CI only
ran the Go gate — the reorg is the natural point to make CI honest about what's gated. Additive; no
behavior change to the mocks.
