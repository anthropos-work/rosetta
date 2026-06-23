# M37 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §4f
(Clerkenstein single→multi identity, G20). Clerkenstein lives in `.agentspace/rosetta-extensions/clerkenstein/`
(its own `knowledge/` is the source of truth; start at `knowledge/kb-index.md`).

## The constraint
_(single identity today: `clerk-frontend/resources.go:21` "the single universal identity every fake-FAPI
session resolves to"; `SingleSessionMode: true` at `:86`; the `Server` holds one `user DemoUser`,
`NewServer(DefaultDemoUser())` at `cmd/fake-fapi/main.go:18`.)_

## What's cheap vs the real work
_(cheap: per-identity JWT minting already exists — `sessionClaimsFor(u,…)` + `MintRS256`, universal HS256 key,
claim shape `{AuthID,Eid,Org*,OrgRole}`. Real work: (a) the registry, (b) active-user selection [O11].)_

## O11 — seat-switch mechanism
_(spike + record: token-injection [set the cookie/storage clerk-js reads] vs a parameterized FAPI handshake.)_

## Alignment
_(new measured surface — author a DNA + goldens; `/align-run` must keep all 4 existing surfaces at 100%.)_

## Pre-flight audits — Users/orgs registry (first section, 2026-06-23)
Phase 0b KB-fidelity: **GREEN** (report: `kb-fidelity-audit.md`). One completeness gap → `KB-1` in
decisions.md (browser-login handshake doc, folds in via the wip reconcile). Topic→doc→code triples:
- single→multi identity → `clerkenstein/knowledge/scope.md` § The demo identity + `architecture.md` § public API → `clerk-frontend/resources.go` (`DemoUser`/`DefaultDemoUser`), `server.go` (`Server.user`, `establishLocked`, `snapshotLocked`, `sessionClaimsFor`).
- FAPI surface → `architecture.md` § Universal-key JWT → `clerk-frontend/server.go` handlers, `shared/jwt.go` (`Claims`, `Mint`/`MintRS256`/`ParseAny`).
- alignment framework → `corpus/architecture/alignment_testing.md` + `clerkenstein/knowledge/alignment.md` → `clerkenstein/alignment/` (`dna/clerk-js-5.json` is the closest analog; runner `cmd/jsfapirun/main.go`; `scripts/gate.sh`).
- hero/org roster (identity source) → `stack-seeding/presets/stories.seed.yaml` (M35) → each story = 1 org + `heroes[]` with `id/name/role/vantage/login(email)/jump_to`.

## Selection-mechanism deployment surface (from demo-stack inspection)
`fake-fapi` runs as one container `demo-N-fake-fapi` seeded `NewServer(DefaultDemoUser())` (`cmd/fake-fapi/main.go:18`).
The injection layer (`stack-injection/gen_injected_override.py`) mounts its TLS cert + offset port (5400). The
seat-switch must let the **browser** choose which registered identity the FAPI's session resolves to — within
the existing single-`fake-fapi`-container deployment (no per-hero container).
