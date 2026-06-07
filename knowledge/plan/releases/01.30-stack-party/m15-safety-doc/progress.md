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

## M15: Hardening

### Pass 1 — 2026-06-07
**Scope manifest (milestone-touched, `git diff release/01.30-stack-party...HEAD`):** 1 net-new deliverable `corpus/ops/safety.md` + 4 cross-linked siblings (`db-access.md`, `snapshot-spec.md`, `seeding-spec.md`, `architecture/security_compliance.md`) + root `README.md`/`CLAUDE.md` + `demo/` recipe family. Extensions KB refresh at ext `main` @ `1d0d2d7`. **No new executable code** — value is accuracy guards + reference integrity (the M14 shape). Baseline: stack-snapshot + stack-seeding `-race` GREEN; ext clone clean.

**Reference-integrity sweep (GREEN, no fixes needed):** every link in safety.md + the 4 siblings resolves both directions; anchors `#multi-tenant-data-isolation` + `§1.1` exist; `demo/` recipe `../safety.md` paths resolve; **0 stale pre-M14 skill names** (`setup-platform`/`start-platform`/`update-platform`/`demo-status`/`demo-seed`/`demo-snapshot`) in the extensions `knowledge/` base.

**Drift guard added (read-side) + M15-D4 fix** — ext `c8a0589`:
- `stack-snapshot/cmd/stacksnap/main_drift_test.go` (extends the M11/M14 `TestDoc…` drift-guard pattern): `TestSafetyDocPredicatesMatchCode` pins the taxonomy + Directus public-predicate literals safety.md quotes to `firewall.PublicFilter` + `directus.Predicate().PublicFilter`; `TestSafetyDocNamesRealFirewallGates` pins the `AssertPlan`/`AssertCaptured` gate names (M15-D2) — symbols referenced so a rename breaks compilation.
- **M15-D4 over-claim fix (Fate-1):** `dev-setdress.sh:20` + the sibling `provision-plan/main_test.go` comment both claimed "stacksnap/stackseed independently refuse N=0 too" — `stacksnap` replay has NO N=0 guard (correct). Corrected both to match the shipped safety.md §2.5. A stale comment contradicting the safety doc is exactly what harden catches.

### Pass 2 — 2026-06-07
**Drift guard added (write-side)** — ext `7f6e53c`. `stack-seeding` had NO docs↔code guard, yet safety.md §2.2 makes load-bearing SECURITY claims about the preflight rejection set. New in-package `stack-seeding/isolation/safety_doc_drift_test.go`: `TestSafetyDocClerkHostsMatchCode` (the complete `realClerkHosts` list), `TestSafetyDocDirectusTokenKeysMatchCode` (the complete `directusTokenKeys` list), `TestSafetyDocForcedBucketOverride` (names `STORAGE_S3_PUBLIC_BUCKET` AND pins the behavior — `PreflightEnv` forces it to `""` on every target), `TestSafetyDocNamesRealGuardSymbols` (`CheckWrite`/`PreflightEnv`/`AssertClean`). In-package so it reads the host/token lists straight off `audit.go`'s vars.

### Pass 3 — 2026-06-07
**Drift guard added (read-safety SQL)** — ext `51ca18b`. `TestSafetyDocBoundedReadSessionMatchesCode` pins safety.md §1.4's quoted capture-session SQL (`SET TRANSACTION READ ONLY` + the timeouts) to `source.DefaultBounds().SetupSQL()` — the single most load-bearing read-safety claim (capture is structurally unable to write), the one quoted code block not yet pinned.

**Knowledge backfill:** none warranted — the new guards PIN safety.md's existing claims to code; they surfaced no new behavioral truth. safety.md was already accurate (Phase 0b KB-fidelity GREEN); the guards make future drift a test failure.

### Pass 4 — stabilization (2026-06-07)
Survey of remaining quoted symbols: every verbatim-literal / code-block / load-bearing-symbol claim in safety.md is now pinned (read-side predicates + gates + bounded-read SQL; write-side rejection set + guard symbols; plus the pre-existing `--source`-kinds + dropped-`--dump` guards). Remaining quoted items (`IsolationClass` String values, `AllowSharedOptIn`, `Record`, `Validate()`) are descriptive prose mentions — pinning them would be brittle string-matching on prose, which this harden deliberately does NOT do.

**Verification:** all 4 Go modules `-race` GREEN (alignment/clerkenstein untouched-pass, stack-snapshot/stack-seeding pass incl. `TestDocSourceSkillRename_M14` + the 7 new guards); `go vet` + `gofmt` clean; `shellcheck dev-setdress.sh` clean; dev-stack pytest 38/38. **Flake gate: 3/3 sequential clean** on the new drift tests (both modules). Both trees clean.

### Stop condition
Scan clean — every load-bearing code-pinned claim in safety.md is now guarded by a fail-closed drift test (each proven to fail on a controlled doc mutation, then restored byte-identical). No new executable code to deepen; remaining unpinned items are prose, not contracts. 4 passes (3 net-new guards + 1 stabilization). Net +7 Go test funcs in the extensions clone, all on ext `main`.
