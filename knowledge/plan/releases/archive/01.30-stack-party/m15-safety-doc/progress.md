# M15 — Progress

**Shape:** section · **Status:** `archived` (completed 2026-06-07)

## Section checklist (from overview Scope.In)
- [x] `corpus/ops/safety.md` — read-side (private-data avoidance: firewall + public predicates + public-only gene) — §1, code-verified
- [x] `corpus/ops/safety.md` — write-side (prod-protection: 3-layer guard + never-write-prod + capture-source policy + n=0 guards) — §2, code-verified
- [x] Cross-link from snapshot-spec / seeding-spec / db-access / security_compliance — back-links added in all four
- [x] Update `rosetta-extensions/knowledge/` for the v1.3 converged model + safety contract — ext repo `main` @ `1d0d2d7` (converged-model + safety-contract sections; stale pre-M14 skill names fixed)
- [x] Refresh root READMEs + demo/ recipes for the unified `stack-*` skills + dev-as-peer — README + CLAUDE.md + demo/README + recipe-snapshot-world (safety.md discoverable; dev-as-peer noted)

## M15: Final Review

**Review verdict: GREEN — 1 finding (a self-referential docs-accuracy fix), 0 code/test/scope.** Addressed fully.

### Scope
- [x] All 5 overview Scope.In sections delivered Fate-1; Out: empty (closing milestone). No silently-dropped item.

### Code Quality
- [x] [n/a] No production code path in rosetta. Extensions drift-guard tests idiomatic (in-package for unexported vars; graceful `t.Skip` on unreachable corpus). gofmt+vet clean; `dev-setdress.sh` shellcheck-clean; py_compile clean.

### Documentation
- [x] [must-fix] `decisions.md` M15-D4 text said the n=0 over-claim was "not fixed here" but harden had landed it Fate-1 — corrected the cell to record the Fate-1 harden landing.
- [x] safety.md indexed (CLAUDE.md + root README + demo/ family) + back-linked from all 4 siblings; `#multi-tenant-data-isolation` anchor + §1.1/§2.5 deep-links resolve. Extensions `knowledge/` converged-model + safety sections present; 0 retired skill names.

### Tests & Benchmarks
- [x] All 4 Go modules `-race -count=1` green (713 funcs, +7 vs M14 baseline = the 7 drift guards). 174 Python passed. Flake gate 5/5 on both touched packages (0 flakes).
- [x] Drift guards mutation-proven fail-closed at close (read-side predicate + write-side bucket-override mutations both tripped; safety.md restored byte-identical).

### Decision Triage
- [x] M15-D1…D4 → archive (doc-authoring rationale; the user-facing truth is already IN safety.md). D4 decisions.md text corrected (above).

### Phase 1b deferral re-audit — **GREEN**
`audit-deferrals/deferral-audit-2026-06-07-m15-close.md`. 1 inherited (DEF-M10-01 cloud snapshot store + S3 blob bytes → v1.4, signed escape-hatch, NOT aged out — all 4 aging triggers negative); 0 new, 0 repeat, 0 chronic, 0 aged-out. M15 added zero deferrals (M15-D3 + M15-D4 both Fate-1). As the terminal milestone, run as the de-facto pre-release sweep across M12→M15.

## M15: Completeness Ledger

**Shape:** section. Every overview Scope.In item placed in exactly one three-fate category. No "tracked-for-later" residue.

### Done (Fate 1) — landed completely in M15
- `corpus/ops/safety.md` read-side (§1) — firewall `AssertPlan`/`AssertCaptured`, per-surface public predicates (code-verified), public-only data-DNA gene, bounded read-only capture.
- `corpus/ops/safety.md` write-side (§2) — the 3-layer guard `CheckWrite`/`PreflightEnv`/`AssertClean`, never-write shared Directus/prod-S3, capture-source policy, doubled n=0 guards, audit-proven zero-pollution.
- Cross-links from `db-access.md` / `snapshot-spec.md` / `seeding-spec.md` / `security_compliance.md` — back-links in all four, resolve both directions.
- `rosetta-extensions/knowledge/` refresh — v1.3 converged dev≡demo model + safety-contract section; 0 stale pre-M14 skill names.
- Root `README.md` + `CLAUDE.md` + `demo/` recipe family refresh — unified `stack-*` skills + dev-as-peer + safety.md discoverability.
- (harden, Fate-1) the M15-D4 n=0 over-claim fix in `dev-setdress.sh` + the sibling test comment.
- (close, Fate-1) the decisions.md M15-D4 text correction.

### Confirmed-covered (Fate 2)
- None (closing milestone — empty Out: list).

### Annotated (Fate 3)
- None.

### Dropped
- None new. (M15-D3 ensures safety.md does not resurrect the offline pg_dump-FILE reader claim — that drop was M9b-D9 in v1.2.)

### Release-scope-breaking deferral (escape hatch)
- None introduced by M15. **DEF-M10-01** (cloud snapshot store + S3 media blob bytes) is INHERITED from v1.2, signed to **v1.4**, re-audited GREEN this close (not aged out). safety.md documents only the current refs-only/local-store posture with a labelled Future(v1.4) pointer.

**Ledger verdict: all scope items delivered as Fate 1. Nothing routed, dropped, or escape-hatch-deferred by M15.** → proceed to merge (no sign-off prompt — 0 escape-hatch entries).

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
