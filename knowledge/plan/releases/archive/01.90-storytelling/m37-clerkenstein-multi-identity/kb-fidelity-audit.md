---
title: "KB Fidelity Audit — M37 Clerkenstein multi-identity"
date: 2026-06-23
scope: milestone:M37
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Single→multi identity (the constraint M37 evolves) | `clerkenstein/knowledge/scope.md` § The demo identity; `architecture.md` § public API | `clerk-frontend/resources.go` (`DemoUser`, `DefaultDemoUser`), `clerk-frontend/server.go` (`Server.user`) | PAIRED |
| FAPI surface / session bootstrap | `clerkenstein/knowledge/architecture.md` § Universal-key JWT | `clerk-frontend/server.go` (handlers), `shared/jwt.go` (`Claims`, mint/parse) | PAIRED |
| Alignment framework (DNA / gene / gate) | `corpus/architecture/alignment_testing.md`; `clerkenstein/knowledge/alignment.md` | `clerkenstein/alignment/` (dna/, golden*/, cmd/*run, scripts/) | PAIRED |
| Interactive browser-login handshake (M2d) | `clerkenstein/knowledge/architecture.md` (MISSING — only on `wip/clerkenstein-browser-login`); `corpus/ops/demo/recipe-browser-login.md` § B (exists) | `clerk-frontend/server.go` (`handleHandshake`, `handleDevBrowser`, `handleClerkJSBundle`), `cmd/fake-fapi/main.go` (TLS) | CODE-ONLY (in `architecture.md`) — the wip-note reconciliation M37 owns |
| Corpus identity pointer | `corpus/ops/rosetta_demo.md`, `corpus/services/clerkenstein.md`, `corpus/services/clerk-integration.md` | n/a (doc-only pointer surface) | PAIRED (doc target for M37's pointer-update) |
| Hero/org roster (identity source) | `stack-seeding/presets/stories.seed.yaml` (M35, shipped) | `stack-seeding/blueprint/stories.go` | PAIRED |

## Fidelity Findings

1. **scope.md § The demo identity** — Expected: "one universal credential … `DefaultDemoUser()`"; Actual: `clerk-frontend/resources.go` defines exactly that single `DemoUser`. **ALIGNED.** (This is the current-code contract M37 will EVOLVE to multi-identity — it is accurate today; M37's Phase 5 updates it.)
2. **architecture.md § public API** — Expected: `clerkfrontend.NewServer(DemoUser) *Server`; Actual: `server.go:29` matches. **ALIGNED.**
3. **alignment.md § four DNAs** — Expected: 4 DNAs (Go 22, JS 9, express 9, deploy 7), all 100%/100%; Actual: `alignment/dna/` carries `clerk-2.6.0.json`, `clerk-js-5.json`, `clerk-express-1.json`, `clerk-deploy-1.json`. **ALIGNED.**
4. **alignment_testing.md § DNA format** — Expected: gene = capability×variant, JSON DNA, criticality weights, operator per variant; Actual: `clerk-js-5.json` + `jsfapirun/main.go` confirm the exact shape. **ALIGNED.** (This is the contract the NEW multi-identity DNA must follow.)

## Completeness Gaps

1. **Interactive browser-login handshake (M2d) is in code but not in `clerkenstein/knowledge/architecture.md`.** `server.go` implements the full dev-instance handshake (`handleHandshake` + dev-browser cookie + RS256 `__session` + `sid` claim) and `cmd/fake-fapi/main.go` serves TLS for it — but `architecture.md` on `main` does not document it. The `wip/clerkenstein-browser-login` branch holds a 32-line note that does. **Classification: critical-but-milestone-owned** — M37's "reconcile the wip branch" deliverable folds this note into `architecture.md`, so it lands in Phase 5 of this milestone rather than blocking it. Tracked as `KB-1` in decisions.md.

## Applied Fixes

None inline — the one completeness gap (browser-login handshake doc) is explicitly a M37 deliverable (the wip-branch reconciliation), so it is folded in during the milestone's documentation phase rather than pre-emptively. No stale claims to fix.

## Open Items (require user decision)

None.

## Gate Result

GREEN: proceed. Every milestone-load-bearing topic is PAIRED with accurate, current-code claims. The one completeness gap (M2d browser-login handshake undocumented in `architecture.md`) is a planned M37 deliverable (the wip-branch fold-in), tracked as `KB-1`, not a blocker.
