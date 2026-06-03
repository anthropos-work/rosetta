# M2b — spec notes

## The green-gate invariant (the safety net for the whole milestone)
M2b moves code; it must not change behavior. After **every section that touches code or scripts**, all
three must still pass (offline, cached modules — `GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off`):
- Go gate → `Score: overall 100.0% critical 100.0% (22/22)`, exit 0.
- JS gate → `Score: overall 100.0% critical 100.0% (9/9)`, exit 0.
- `drift-test.sh` → `ALL PASS` (9/9).
- `go test -race ./...` → all packages ok; `gofmt`/`go vet`/`shellcheck` clean.

If a move breaks a gate, the move is wrong — fix the import/path, don't touch the mock logic.

## Restructure map (current → target) — S1
| Current | Target | Package | Notes |
|---|---|---|---|
| `authn/jwt.go` (universalSecret, `Claims`, mint/parse) | `shared/jwt.go` | `shared` | shared because `clerk-frontend` mints + `authn` verifies with the same key |
| `authn/provider.go`, `authn/user.go` + tests | `authn/` | `authn` | imports `clerkenstein/shared`; mocks `colony/authn` |
| `bapi/{server,resources,doc}.go` + tests | `clerk-backend/` | `clerkbackend` | the Clerk **B**ackend-API mock |
| `orgclient/{store,invitations}.go` + tests | `clerk-backend/` | `clerkbackend` | **merged in** — the in-memory store behind the BAPI server |
| `fapi/{server,resources,key}.go` + tests | `clerk-frontend/` | `clerkfrontend` | the Clerk **F**rontend-API mock; imports `shared` |
| `webhook/{injector,events}.go` + tests | `clerk-webhook/` | `clerkwebhook` | the svix injector |
| `cmd/clerkrun/` | `alignment/cmd/clerkrun/` | `main` | imports `authn` + `clerkbackend` |
| `cmd/jsfapirun/` | `alignment/cmd/jsfapirun/` | `main` | imports `clerkfrontend` (+ `shared`) |
| `dna/`, `golden/`, `golden-js/` | `alignment/{dna,golden,golden-js}/` | — | data assets |
| `scripts/{gate,drift-check,drift-test}.sh` | `alignment/scripts/` | — | repoint defaults (below) |
| `.github/workflows/alignment.yml` | (same path) | — | repoint script + ALIGN_DIR paths |
| `README.md` | `README.md` (slimmed) | — | → points to `knowledge/` |

## M2b-D1 — Go package naming under the library-named scheme
Go package **identifiers** can't contain hyphens. Each hyphenated dir declares a clean package:
`clerk-backend/` → `package clerkbackend`, `clerk-frontend/` → `package clerkfrontend`,
`clerk-webhook/` → `package clerkwebhook`. Import paths are `clerkenstein/clerk-backend` etc. (hyphens are
legal in import *paths*); call sites read `clerkbackend.X`. **Confirm at build** — alternative is
dropping hyphens entirely (`clerkbackend/`), but the user picked the hyphenated library-named look.

## Script repointing (S1) — so the DEFAULT invocation still works
`alignment/scripts/gate.sh` defaults change:
- `RUNNER_PKG` `./cmd/clerkrun` → `./alignment/cmd/clerkrun`; `DNA` `dna/clerk-2.6.0.json` →
  `alignment/dna/clerk-2.6.0.json`; `GOLDEN_DIR` `golden` → `alignment/golden`.
- JS gate env: `RUNNER_PKG=./alignment/cmd/jsfapirun DNA=alignment/dna/clerk-js-5.json GOLDEN_DIR=alignment/golden-js`.
- `ALIGN_DIR` (locates rosetta's `alignctl`) default `../../test/alignment` → now `../../../test/alignment`
  (one level deeper, since scripts moved from `scripts/` to `alignment/scripts/`). **Verify the relative
  depth at build** and update `drift-check.sh` + the CI YAML to match.
- `drift-check.sh` / `drift-test.sh` likewise repoint DNA + golden + ALIGN_DIR.

## knowledge/ base outline (S2) — per `repo-consolidate` code-repo standard + templates/code/
- `knowledge/kb-index.md` — entry/index (code-repo KB index template).
- `knowledge/scope.md` — what Clerkenstein is + why; disarmed-by-design properties (speed + accessibility,
  not security); the M0-mirror provenance (it's *measured*, not hand-built).
- `knowledge/architecture.md` — the 4 library mocks + `shared` + the `alignment` harness; per-dir
  responsibilities; the universal-key JWT flow (frontend mints → authn verifies).
- `knowledge/alignment.md` — how fidelity is validated: the M0 framework (`test/alignment/` in rosetta),
  the two DNAs (`clerk-2.6.0` Go = 11 caps/22 genes; `clerk-js-5` JS = 6 caps/9 genes), the gate
  (100%/100%), **version pinning** + drift detection (M1b). Links to `corpus/architecture/alignment_testing.md`.
- `knowledge/injection.md` — the four per-library recipes: `go.mod replace` whole-colony + skip-worktree
  (authn); `api.clerk.com` DNS/`/etc/hosts` redirect to the fake BAPI (clerk-backend); config-only minted
  publishable-key encoding the fake FAPI host (clerk-frontend); direct svix-signed POST to
  `/api/webhook/clerk` (clerk-webhook). Note which are recipe-only vs spike-proven vs built+gated.
- `knowledge/coverage-index.md` — per-package test coverage map (code-repo coverage-index template).
- Per-library `README.md`: `authn/`, `clerk-backend/`, `clerk-frontend/`, `clerk-webhook/`, `shared/`, `alignment/`.

## repo-consolidate integration (S4)
`/singularity-kit:repo-consolidate code` will detect clerkenstein as **code / library / `go-library`**,
load base + code-repo + asset-hygiene standards, audit, and create/fix: `CLAUDE.md` (slim, from the
audit), `knowledge/kb-index.md` (if not already substantive), a coverage index, and a root
`singularity-manifest.md`. Asset-hygiene will check `.gitignore` completeness (the built `clerkrun` /
`jsfapirun` binaries are already gitignored at root + anchored) and confirm no tracked secrets/transient
files. **Because S2/S3 pre-build to its standard, the run should report mostly-compliant** — it's a
finalize + manifest step, **user-invoked** (the skill is `disable-model-invocation`).

## Pre-flight audits
**Phase 0b — KB-fidelity** runs at build (S1). The load-bearing contract docs all exist
(`corpus/services/clerkenstein.md`, `corpus/architecture/alignment_testing.md`); M2b *adds* the repo's own
`knowledge/` base + slims the corpus doc. Expect GREEN.

### Pre-flight audits — S1 (restructure)
**Verdict: GREEN** (2026-06-03, sha `26b2490`, report `kb-fidelity-audit.md`). Both load-bearing corpus
docs PAIRED + ALIGNED with the current flat repo; baseline green-gate confirmed before any move (Go 22/22,
JS 9/9, drift ALL PASS, `-race`/gofmt/vet/shellcheck clean). The repo's own `knowledge/` base is a planned
S2 deliverable, not a blind area. No applied fixes, no open items.

Topic → doc → code triples (for fast re-audit of later sections):
- Clerkenstein mirror → `corpus/services/clerkenstein.md` → `anthropos-demo/clerkenstein/{authn,bapi,orgclient,fapi,webhook,cmd,dna,golden,golden-js,scripts}`
- Alignment framework → `corpus/architecture/alignment_testing.md` → `test/alignment/`
- Repo `knowledge/` base → (authored in S2) → `anthropos-demo/clerkenstein/knowledge/`
