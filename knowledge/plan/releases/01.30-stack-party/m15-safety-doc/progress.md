# M15 — Progress

**Shape:** section · **Status:** built (all sections landed; ready for close)

## Section checklist (from overview Scope.In)
- [x] `corpus/ops/safety.md` — read-side (private-data avoidance: firewall + public predicates + public-only gene) — §1, code-verified
- [x] `corpus/ops/safety.md` — write-side (prod-protection: 3-layer guard + never-write-prod + capture-source policy + n=0 guards) — §2, code-verified
- [x] Cross-link from snapshot-spec / seeding-spec / db-access / security_compliance — back-links added in all four
- [x] Update `rosetta-extensions/knowledge/` for the v1.3 converged model + safety contract — ext repo `main` @ `1d0d2d7` (converged-model + safety-contract sections; stale pre-M14 skill names fixed)
- [x] Refresh root READMEs + demo/ recipes for the unified `stack-*` skills + dev-as-peer — README + CLAUDE.md + demo/README + recipe-snapshot-world (safety.md discoverable; dev-as-peer noted)

## Final review
_(filled at close)_

## Build notes
- Phase 0b KB-fidelity: **GREEN** (`kb-fidelity-audit.md`) — every read/write safety claim verified against the actual extensions code before authoring. Two accuracy guardrails carried in: M15-D3 (no offline file reader), M15-D4 (precise n=0 scope; flagged a pre-existing over-claim in the dev-setdress source comment).
- Decisions: M15-D1 (doc home `corpus/ops/safety.md`), M15-D2 (name the real Go funcs not the conceptual umbrella), M15-D3, M15-D4.
- PR review: 0 findings — predicate strings byte-match `firewall.go`; function names match; all cross-links resolve; no stale skill names.
- Commits — rosetta (`m15/safety-doc`): `da18188` §1+§2, `423a9c8` §3, `9cb3c6f` §5. Extensions (`main`): `1d0d2d7` §4 KB refresh.
