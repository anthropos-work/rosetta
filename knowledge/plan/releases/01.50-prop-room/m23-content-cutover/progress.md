# M23 — Progress

**Status:** in progress. **Shape:** section.

## Section checklist
_One checkbox per concrete deliverable from `overview.md` § Scope. Code lands in the `rosetta-extensions`
authoring copy (tag `prop-room-m23` at close); docs land in the `rosetta` worktree._

- [x] **§1 Dev env-emission plumbing** (ext) — grow `stack-core/gen_override.py` to emit per-service
  `environment:` blocks (today it emits only ports/volumes — the single genuinely-new bit of plumbing).
  _Landed: `directus_consumer_env()` + env-emission in `build_override`/`to_yaml`; dev `--with-directus`
  re-points cms `DIRECTUS_BASE_ADDR`→in-network instance + strips prod token; asset plane untouched. ext
  `3506976`. +9 tests, 19/19 green._
- [x] **§2 `DIRECTUS_BASE_ADDR` re-point — demo + dev** (ext) — inject the local-Directus data-plane address
  (`http://directus:8055`, the in-network compose service) into the Directus-consuming services; keep
  `DIRECTUS_PUBLIC_BASE_ADDR` on prod (asset plane stays real); extend the prod-token strip to opted-in dev.
  _Landed: dev side in §1 (ext `3506976`); demo side ext `b0f3945` (cms re-point gated on with_directus, asset
  plane untouched, prod-read path preserved on --no-local-content). Both bring-ups already thread the flag.
  82/82 demo + 19/19 dev green._
- [x] **§3 studio-desk local instance + minted admin token** (ext) — point studio-desk's `DIRECTUS_BASE_URL`
  at the per-stack instance + a locally-minted admin token so its skill-path writes target local, never prod.
  _Landed ext `5218953`: `EnvContract.AdminToken` (deterministic) → bootstrap `ADMIN_TOKEN` (stamps the
  admin's static token, verified vs the pinned directus image) → studio-desk `DIRECTUS_BASE_URL` + that token
  on local-content. `ValidateProvisionable` adds the present-token gate; `Validate` stays the prod firewall.
  +12 tests green._
- [x] **§4 `directus_files` ref capture** (ext, Fate-3 from M21/DEF-M21-03) — wire the dead `media.go`
  (FileRefColumns/ReferencedFilesFilter) into `directus.Surface()` via a `directus_files` TableSpec so captured
  content rows resolve their image-asset UUIDs to the prod-public asset-plane URLs (refs only; blob bytes stay
  backlog DEF-M10-01). _Landed ext `2b8e9a0`: new REFERENCED-SUBSET admissibility kind (reverse-reference
  closure, firewall-admitted); `ReferencedFilesColumns()` (26-col, verified vs fresh bootstrap); + the
  replay-clear DELETE/TRUNCATE split for the external directus_settings FK. +13 tests, 12/12 packages green._
- [x] **§5 Referential closure + cross-surface fidelity gene** (ext, DEF-M21-04/NEW-3) — make the taxonomy
  capture referentially closed against the content it serves (full-taxonomy capture, the simple fallback the
  corpus already names) + a measured cross-surface closure gene (no content row references a taxonomy node-id
  the captured subset lacks) + close the 20 dangling relations. _Landed ext `4cb8786`: full-taxonomy capture
  already the state (org_id IS NULL); new `OpSnapshotCrossSurfaceClosure` gene measures content→taxonomy
  node-id closure (standard crit, non-blocking); surfaces the 1 genuine prod residual (K-AIFUNX-E658, a public
  sim → customer-only skill, uncloseable w/o firewall breach). The 20 dangling relations were subsumed by M21's
  26-collection structure capture (M21-D7); M23 owns the external content→taxonomy refs — that's this gene._
- [x] **§6 Docs — env truth + safety + closure gene** (rosetta) — `corpus/services/{cms,studio-desk,
  jobsimulation,next-web-app}.md` (env/dependency truth), `corpus/ops/safety.md` (retire the live-prod-read
  notes; token-strip stays as the write-disarm), `corpus/ops/snapshot-spec.md` (the cross-surface closure gene).
  _Landed rosetta `e364b80`: all 6 docs updated; resolves KB-1 (directus_files now wired+documented) + KB-2
  (cms.md M10-gap retired). cms cutover, studio-desk local token, jobsim-via-RPC, next-web-no-direct-Directus,
  safety retire-live-read, snapshot-spec directus_files-true + closure gene + the 1 named prod residual._

## Build log
_(append per build session)_

## M23: Hardening

### Pass 1 — 2026-06-13

**Scope manifest (M23-touched, diff `e989982..4cb8786`):** 14 source files across 4 stacks —
Python (`stack-core/gen_override.py`, `stack-injection/gen_injected_override.py`), Go
`stack-seeding/dna` (`snapshot.go`, `fidelity_probe.go`), Go `stack-snapshot`
(`directus/{directus,media,provision}.go`, `firewall/firewall.go`, `manifest/manifest.go`,
`replay/replay.go`, `capture/capture.go`, `cmd/stacksnap/adapters.go`,
`cmd/provision-plan/main.go`, `dev-stack/dev-setdress.sh`). Every package ships co-located
tests. Highest-priority gap from the func-level coverage sweep: **`PgFidelityProbe.CrossSurfaceDangling`
shipped at 0%** (the existing `fakeScanner` rejected its TWO-dest `(count, sample)` `QueryRow`
and didn't classify the cross-surface CTE SQL).

**Coverage delta (milestone-touched files):**
- `stack-seeding/dna`: 86.5% → 87.5% (statements); **`CrossSurfaceDangling` 0% → 100%**.
- `stack-snapshot/cmd/stacksnap`: 80.3% → 80.9%; **`CountTenantRows` 93.1% → 100%**.
- `stack-snapshot/firewall`: 100% → 100% (added behavioral pins, not coverage).
- `stack-snapshot/manifest` 98.4%, `replay` 100% (ClearByDelete round-trip + propagation pinned).
- Python: `gen_override.py` 87%, `gen_injected_override.py` 99% — M23 surface already fully covered.

**Tests added:**
- `dna/fidelity_probe_test.go`: +3 (cross-surface closure probe — closed / open+sample / wrapped-error).
- `cmd/stacksnap/adapters_harden_test.go`: +4 (referenced-subset leak surfacing, else-branch
  hascol error, closure-probe count error, the closure-vs-scope-bearing discriminator).
- `firewall/firewall_test.go`: +2 (scope-column+referenced-subset diagnosed via the scope branch;
  whitespace closure filter treated as absent).
- `manifest/manifest_test.go`: +1 (ClearByDelete Marshal→Parse round-trip + omitempty).
- `replay/replay_{test,harden_test}.go`: +1 (ClearByDelete flag propagates manifest.Table → TableRef).

**Bugs fixed inline:** none — every gap was a test gap (no production bug surfaced).

**Flakes stabilized:** none observed.

**Knowledge backfill:** no KB-worthy findings beyond what the build phase already documented —
the cross-surface closure semantics + K-AIFUNX-E658 residual live in `decisions.md` M23-D5 and
`snapshot-spec.md`; the DELETE-before-TRUNCATE ordering in M23-D4 + the `replay.go`/`adapters.go`
docstrings. Hardening confirmed those invariants by test; it surfaced no new behavior to backfill.

### Pass 2 — 2026-06-13

Swept the demo-side §2/§3 surface (`gen_injected_override.py` cms re-point + studio-desk minted
token) — found it already comprehensively covered (token-strip on every service, studio-desk local
token byte-for-byte vs `directus_static_token`, asset-plane-untouched, the no-local-content
prod-disarm path). Closed the one remaining real Go gap:

**Coverage delta:** `stack-snapshot/directus` 99.1% → 100%; **`ValidateProvisionable` 80% → 100%**.

**Tests added:**
- `directus/provision_test.go`: +1 (`ValidateProvisionable` propagates the inner prod-safety
  `Validate` failure BEFORE the present-token gate — a valid local token must not mask a prod BaseAddr).

**Bugs fixed inline:** none. **Flakes:** none. **Knowledge backfill:** none warranted.

### Stop condition
Stopped after Pass 2. All three stop criteria met: (1) the full six-dimension scan found nothing
new worth adding — every M23-touched function is at 100% except the DB-only concrete-conn executors
(`ClearForReplay`, `CopyIn`, provision DDL), which are structurally unit-untestable and whose pure
SQL-builder seams (`clearForReplaySQL`/`truncateForReplaySQL`) ARE 100%; (2) coverage delta on the
remaining unitable surface is <2%; (3) zero flakes across 3 consecutive sequential runs of every
touched suite (171 passed + 8 skipped Python; all Go packages ok ×3).

**Total: +11 tests across 2 passes, 0 inline bugs, 0 flakes.** Ext commits: `ceed313` (cross-surface
probe), `23767da` (directus_files referenced-subset + ClearByDelete), `7e9343a` (ValidateProvisionable).

## M23: Final Review

_Close pass 2026-06-13. Parallel scope/code/adversarial/docs/test scans + the blocking deferral re-audit.
Findings consolidated below; all addressed fully (no partial fixes)._

**Summary: 6 findings — 0 scope · 0 code-quality · 1 docs · 0 tests · 5 decision-triage.** Deferral re-audit
**GREEN** (2 inherited M21 items RESOLVED in-milestone; 1 prod-data residual fated KNOWN-ISSUE). All 4 test
suites green; vet + shellcheck + py_compile CLEAN; flake 0 (5/5).

### Scope
- [x] All 6 sections landed Fate-1; nothing dropped/routed. DEF-M21-03 (directus_files) + DEF-M21-04
  (referential closure) **RESOLVED** here — drop off the ledger. No new scope gaps.

### Code Quality
- [x] Cross-cutting scan (30 ext files, 4 stacks) — **CLEAN, 0 findings.** The two override emitters are
  consistent (shared `http://directus:8055` + `DIRECTUS_DATA_CONSUMERS=("cms",)`); the firewall
  referenced-subset admit-iff branch validates input + is mutually exclusive; the CrossSurfaceDangling probe
  handles nil/empty/error; media.go fully wired (no dead code); ClearByDelete cleanly threaded
  TableSpec→manifest→TableRef; DELETE-before-TRUNCATE correct + scoped. No new TODO/FIXME/HACK.

### Adversarial
- [x] No new fail-mode found — every scenario the adversarial pass probed (filter slipping the admit-iff gate,
  a valid token masking a prod BaseAddr, a two-dest scan on an empty/error result, the directus_settings FK on
  TRUNCATE) is already test-pinned (firewall_test, provision_test, fidelity_probe_test, replay_harden_test). See
  decisions.md § Adversarial review.

### Documentation
- [x] [fix] `snapshot-spec.md` (M13 section, ~L499-501) had a stale future-tense claim — "The M23 *cutover* …
  remains future." M23 landed it. Rewritten: the `--local-content` pass now performs the cutover (re-point cms +
  token strip, asset plane unchanged); the non-`--local-content` prod-read path is the documented fallback.

### Tests & Benchmarks
- [x] Full suites green: stack-core 69 · stack-injection 110 (8 env-gated skip) · stack-snapshot all-pkg ok ·
  stack-seeding all-pkg ok. No test gaps (the harden pass already closed the CrossSurfaceDangling 0% gap +
  ValidateProvisionable). Pre-existing ResourceWarnings (unclosed files) live in test lines M23 did NOT touch
  (`test_stack_registry.py`, pre-M23 `test_injection.py` idioms) — carried as a known pre-existing hygiene smell,
  not an M23 fix (warnings, not failures).

### Decision Triage
- [x] M23-D1 (cms-only cutover) → blended in cms.md + snapshot-spec.md; backref tag `(#M23-D1)` added.
- [x] M23-D2 (studio-desk minted token) → blended in studio-desk.md; backref tag `(#M23-D2)` added.
- [x] M23-D3 (directus_files referenced-subset) → blended in snapshot-spec.md § Media/blobs; tag `(#M23-D3)`.
- [x] M23-D4 (replay-clear DELETE/TRUNCATE split) → blended in snapshot-spec.md; tag `(#M23-D4)`.
- [x] M23-D5 (cross-surface closure gene + K-AIFUNX-E658 residual) → blended in snapshot-spec.md; tag `(#M23-D5)`.
- [x] KB-1/KB-2 → already resolved in §6 build (directus_files wired+documented; cms.md M10-gap retired). Archive.
