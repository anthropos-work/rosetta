# M9a — Progress

**Shape:** section · **Status:** built (awaiting close)

## Section checklist (from overview Scope.In)
- [x] (note #1) The dedicated `rosetta-extensions/stack-snapshot/` section + the `stacksnap` CLI (capture / replay / status) — 9 Go packages, tagged `stack-snapshot-m9a`
- [x] Snapshot contract + portable format (per-table COPY payloads + `manifest.json`, schema-version pinned + SHA-256 checksums) — `manifest` package
- [x] (note #2) Production-safe capture-source policy (M9a-D3): cache-hit → (1) prod-`pg_dump` ingest [default] → (2) safe throttled primary read [MVCC, fallback] → (3) restore-from-snapshot / (4) read replica [upgrades]; bounded read-only session + catalog-first dry-run — `source` + `pg` + `capture.BuildPlan`
- [x] (note #3) Tenant-data firewall — `AssertPublicOnly` (plan + post-capture, hard-fail on any tenant row) — `firewall` package; the public-only/provenance gene is in the data-DNA (below)
- [x] (note #4) `.agentspace` manifest-cached store with a pluggable `SnapshotStore` backend (localfs now; cloud/S3 = v1.3); cache-hit vs stale→refresh — `store` package
- [x] Data-DNA extension: `snapshot-seeded` status (counts toward coverage — the v1.2 thesis) + snapshot-fidelity gene class (row-count / structural / referential / embedding-dimension / public-only); `datadna` catalog recognizes snapshot surfaces — `stack-seeding/dna/snapshot.go`
- [x] (note #5) `/db-query` skill (`.claude/skills/db-query/SKILL.md`) + `corpus/ops/db-access.md` (MCP-tool + pgpass/psql paths) — inherited from the release branch; corrected to the M9a-D3 capture-source precedence in Phase 0b
- [x] Tiny reference surface proving capture→store→replay→fidelity-gate end-to-end (no real surface) — `reference` package (hermetic, composes the real packages)
- [x] Delivers: `corpus/ops/snapshot-spec.md` (new) + `corpus/ops/db-access.md` (inherited) + alignment_testing.md snapshot-fidelity + public-only genes

## M9a: Hardening

### Pass 1 — 2026-06-06

**Scope manifest (milestone-touched code, from `git diff release/01.20-set-dressing...stack-snapshot-m9a`):**

Two Go modules in the `rosetta-extensions` monorepo (the code-under-test; rosetta side is docs-only, no test-deepening).

_Module A — `stack-snapshot/` (new, 9 pkgs + `stacksnap` CLI):_
| Package | Source | Test | Baseline cov |
|---|---|---|---|
| `firewall` | firewall.go | firewall_test.go | 100.0% |
| `reference` | reference.go | reference_test.go | 100.0% |
| `manifest` | manifest.go | manifest_test.go | 98.2% |
| `source` | source.go | source_test.go | 93.3% |
| `replay` | replay.go | replay_test.go | 92.3% |
| `capture` | capture.go | capture_test.go | 81.0% |
| `store` | store.go | store_test.go | 80.0% |
| `cmd/stacksnap` | main.go, adapters.go, surfaces.go | main_test.go, adapters_test.go | 40.7% |
| `pg` | pg.go | pg_test.go | 51.7% (rest is live-DB `Conn` methods, exercised via fakes in capture/replay) |

_Module B — `stack-seeding/dna/` (extended):_ `snapshot.go` (new, snapshot-fidelity operators) + `dna.go`/`catalog.go` (snapshot-status `Validate`/`Coverage` branches) + `snapshot_test.go`. Baseline 83.6%.

**New-unit handbook check:** `stack-snapshot/` ships `README.md` (99 lines) + a `knowledge/README.md` nav row — handbook contract satisfied, no Phase 2b gap.

**Highest-priority gaps (testable without a live DB):** `cmd/stacksnap` source-resolution branches + populated-store `status`; `store` error paths + `sanitize`/`safeFilename` traversal fuzzing; each `dna` snapshot operator's probe-error early-return; `capture`/`source`/`manifest`/`replay` error branches; `pg.replaceURLPort` malformed-DSN + `bytesReader` partial reads.

**Coverage delta (milestone-touched files), Pass 1:**
| Package | Before | After | Δ |
|---|---|---|---|
| `store` | 80.0% | 90.0% | +10.0 |
| `cmd/stacksnap` | 40.7% | 58.9% | +18.2 |
| `capture` | 81.0% | 89.7% | +8.7 |
| `replay` | 92.3% | 100.0% | +7.7 |
| `source` | 93.3% | 100.0% | +6.7 |
| `pg` | 51.7% | 52.5% | +0.8 (rest is live-DB `Conn` methods) |
| `dna` (snapshot.go) | 83.6% | 85.0% | +1.4 (snapshot.go itself near-100%) |
| `manifest` / `firewall` / `reference` | 98.2 / 100 / 100 | unchanged | already strong |

**Tests added (Pass 1):**
- `store/store_harden_test.go`: 9 unit (corrupt/invalid manifest, missing payload, vanished root, stray-file skip, traversal guards, Resolve error-propagation, hostile-ref-stays-in-root) + 1 **fuzz** (`FuzzSanitize` — path-traversal invariant).
- `cmd/stacksnap/main_harden_test.go`: 11 unit (source-resolution branches, replay arg-validation, populated/unreadable `status`, `captureDryRun` via fake `Capturer`).
- `capture/capture_harden_test.go`: 6 unit (invalid surface, introspect error, invalid bounds, tenant-probe error, store-write error, index-rebuild threshold).
- `replay/replay_harden_test.go`: 3 unit (manifest-read fault, payload-read fault aborts before COPY, empty-tables no-op).
- `source/source_harden_test.go`: 5 unit (Available default, precedence-order, skip-unavailable, negative-idle, partial-bounds).
- `pg/pg_harden_test.go`: 7 unit (replaceURLPort malformed/no-port-with-query/no-userinfo, ParseStackN edges, QuoteIdent injection-safety, bytesReader chunked + empty).
- `stack-seeding/dna/snapshot_harden_test.go`: 6 unit (every operator's probe-error path, embedding-dim no-vectors vacuous-pass, referential empty-set vs no-FK, stable-order, all-waived coverage).

**Bugs fixed inline:** none — no production-code defect surfaced. All gaps were untested behavior, not incorrect behavior. (One observation pinned: `sanitize("a...b") → "a-.b"`, still safe — the `..`-collapse is single-pass-per-iteration but loops to fixpoint; no traversal run survives.)

**Flakes stabilized:** none observed — 3/3 clean sequential runs on both modules; `FuzzSanitize` ran 6.3M execs with no failure.

### Pass 2 — 2026-06-06
**Coverage delta:** `cmd/stacksnap` 58.9% → 60.1% (+1.2, the `defaultStoreRoot` cwd-fallback); all other packages unchanged. Deltas now <2% across all touched files.

**Tests added:** 1 unit (`TestDefaultStoreRoot_CwdFallback`). Extended `FuzzSanitize` to 20s / 6.3M execs — still no traversal escape.

**Bugs fixed inline / flakes:** none.

### Stop condition
Loop terminated after **Pass 2** (well under the 5-pass cap). All three stop criteria met: the full Step 2b scan found nothing more worth adding (the residual uncovered code — `cmd/stacksnap` adapters, `pg.Conn` methods, `main()`'s `os.Exit` shim, post-`pg.Connect` orchestration — is **irreducibly live-DB-bound**, sitting behind the deliberate `Capturer`/`Replayer`/`SnapshotStore` interface seam that lets all logic be hermetically tested via fakes); coverage deltas <2% on every touched file; zero flakes across the gate. **Knowledge backfill:** none warranted — the snapshot framework's behavioral invariants (firewall double-check, checksum-before-write, cache-staleness key, public-only provenance) are already documented in `corpus/ops/snapshot-spec.md` and the `alignment_testing.md` snapshot dimension; the harden surfaced no new system truth not already captured there.

## Final review
_(filled at close)_
</content>
